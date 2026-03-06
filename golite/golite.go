package main

import (
	"flag"
	"fmt"
	"golite/arm"
	"golite/ir"
	"golite/lexer"
	"golite/parser"
	"golite/sa"
	"os"
	"strings"
)

func main() {
	// Debug flags for inspecting intermediate stages
	lexerFlag := flag.Bool("l", false, "print lexer tokens")
	astFlag := flag.Bool("ast", false, "print AST")
	llvmFlag := flag.String("target", "arm64-apple-macosx14.0.0", "llvm target architecuture")
	// stack mode skips SSA optimization passes and outputs a simple alloca-based .ll
	llvmStackFlag := flag.Bool("llvm-stack", false, "output stack-based LLVM IR instead of register-based SSA")
	armFlag := flag.Bool("S", false, "Compile LLVM to ARM64 assembly")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Error: filename is required")
		fmt.Println("Usage: program [-flags] <filename>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputSourcePath := args[0]

	// Pipeline: lex -> parse -> semantic analysis -> IR -> ARM code generation
	lexer := lexer.NewLexer(inputSourcePath)
	if *lexerFlag {
		lexer.PrintTokens()
	}
	parser := parser.NewParser(lexer)
	program := parser.Parse()
	if *astFlag {
		parser.PrintAST(program)
	}
	// SA populates symbol tables and validates types; nil means errors were found
	tables := sa.Execute(program)
	if tables == nil {
		os.Exit(1)
	}
	builder := ir.NewBuilder(program, tables)
	filename := strings.TrimSuffix(inputSourcePath, ".golite")
	// BuildProgram handles IR generation + optimization passes + writing .ll file
	builder.BuildProgram(filename, *llvmFlag, *llvmStackFlag)

	if *armFlag {
		armBuilder := arm.NewBuilder(builder, *llvmFlag)
		armBuilder.BuildProgram(filename)
	}
}
