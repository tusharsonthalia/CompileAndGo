package arm

import (
	"fmt"
	"golite/ir"
	"os"
	"strings"
)

type Builder interface {
	BuildProgram(filename string)
}

type builder struct {
	irBuilder ir.Builder
	dataLines []string
	bssLines  []string
	textLines []string
	isApple   bool
}

func NewBuilder(irBuilder ir.Builder, target string) Builder {
	targetLower := strings.ToLower(target)
	isApple := strings.Contains(targetLower, "apple") || strings.Contains(targetLower, "darwin")

	return &builder{
		irBuilder: irBuilder,
		dataLines: make([]string, 0),
		bssLines:  make([]string, 0),
		textLines: make([]string, 0),
		isApple:   isApple,
	}
}

func (b *builder) emitData(s string) {
	b.dataLines = append(b.dataLines, s)
}

func (b *builder) emitBss(s string) {
	b.bssLines = append(b.bssLines, s)
}

func (b *builder) emitText(s string) {
	b.textLines = append(b.textLines, s)
}

func (b *builder) BuildProgram(filename string) {
	b.emitText("\t.arch armv8-a")

	b.translateGlobals()
	b.translateFunctions()

	outFile := filename + ".s"
	f, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Error creating assembly file %s: %v\n", outFile, err)
		return
	}
	defer f.Close()

	if len(b.dataLines) > 0 {
		f.WriteString("\t.data\n")
		for _, line := range b.dataLines {
			f.WriteString(line + "\n")
		}
		f.WriteString("\n")
	}

	if len(b.bssLines) > 0 {
		f.WriteString("\t.bss\n")
		for _, line := range b.bssLines {
			f.WriteString(line + "\n")
		}
		f.WriteString("\n")
	}

	f.WriteString("\t.text\n")
	for _, line := range b.textLines {
		f.WriteString(line + "\n")
	}
}
