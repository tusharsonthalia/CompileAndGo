package sa

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
