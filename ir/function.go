package ir

import (
	"fmt"
	"golite/symboltable"
	"strings"
)

type Function struct {
	Entry  *symboltable.FuncEntry
	Blocks []*BasicBlock
}

func NewFunction(entry *symboltable.FuncEntry) *Function {
	return &Function{
		Entry:  entry,
		Blocks: make([]*BasicBlock, 0),
	}
}

func (f *Function) AddBlock(block *BasicBlock) {
	f.Blocks = append(f.Blocks, block)
}

func (f *Function) String() string {
	var out strings.Builder
	out.WriteString(fmt.Sprintf("define %s @%s(", LLVMType(f.Entry.ReturnTy), f.Entry.Name))

	for i, param := range f.Entry.Params {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(fmt.Sprintf("%s %%%s", LLVMType(param.Ty), param.Name))
	}
	out.WriteString(") {\n")

	for _, block := range f.Blocks {
		out.WriteString(block.String())
	}

	out.WriteString("}\n")
	return out.String()
}
