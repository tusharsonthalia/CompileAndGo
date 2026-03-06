package sa

// Package sa runs semantic analysis in two passes over the AST:
// 1. Name resolution: populates the symbol table and checks all identifiers exist
// 2. Type checking: validates type compatibility for all expressions and statements
// If either pass finds errors, it prints them and returns nil to halt compilation.

import (
	"golite/ast"
	"golite/context"
	st "golite/symboltable"
)

func Execute(program *ast.Program) *st.SymbolTables {
	tables := st.NewSymbolTables()

	nr := NewNameResolver(make([]*context.CompilerError, 0), tables)
	nr.ResolveAST(program)

	if nr.HasErrors() {
		return nil
	}

	checker := NewTypeChecker(make([]*context.CompilerError, 0), tables)
	checker.TypeCheckAST(program)
	if checker.HasErrors() {
		return nil
	}

	return tables
}
