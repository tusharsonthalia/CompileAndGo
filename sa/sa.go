package sa

// Package sa implements the Semantic Analysis phase of the compiler.
// It is responsible for:
// 1. Symbol Table mapping: Resolving names and scopes for variables, functions, and types.
// 2. Type checking: Ensuring that operations are valid for their given types.
// If any errors are found during these steps, the compiler returns nil to halt further processing.

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
