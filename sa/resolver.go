package sa

import (
	"fmt"
	"golite/ast"
	"golite/context"
	st "golite/symboltable"
	"golite/types"
)

type NameResolver interface {
	addError(line, col int, msg string)
	HasErrors() bool
	PrintErrors()
	ResolveExpression(astExpr ast.Expression, table *st.SymbolTable[st.TypeEntry])
	ResolveStatement(astStmt ast.Statement, table *st.SymbolTable[st.TypeEntry])
	ResolveTypeDecl(decl *ast.TypeDecl)
	ResolveVarDecl(decl *ast.VarDecl)
	ResolveFuncDecl(decl *ast.FuncDecl)
	ResolveAST(*ast.Program)
}

type nameResolver struct {
	errors []*context.CompilerError
	tables *st.SymbolTables
}

func NewNameResolver(errors []*context.CompilerError, tables *st.SymbolTables) NameResolver {
	return &nameResolver{errors, tables}
}

func (nr *nameResolver) addError(line, col int, msg string) {
	nr.errors = append(nr.errors, &context.CompilerError{
		Line:  line,
		Col:   col,
		Msg:   msg,
		Phase: context.SEMANTIC,
	})
}

func (nr *nameResolver) HasErrors() bool {
	return context.HasErrors(nr.errors)
}

func (nr *nameResolver) PrintErrors() {
	for _, err := range nr.errors {
		fmt.Println(err)
	}
}

// ResolveExpression checks that all identifiers used in an expression are declared
// in the given scope. For calls, we check the global scope since functions are top-level.
func (nr *nameResolver) ResolveExpression(astExpr ast.Expression, table *st.SymbolTable[st.TypeEntry]) {
	switch expr := astExpr.(type) {

	case *ast.Variable:
		if _, contains := table.Contains(expr.Name); !contains {
			msg := fmt.Sprintf("Variable \"%v\" is not defined.", expr.Name)
			nr.addError(expr.Line, expr.Column, msg)
		}
	case *ast.LValue:
		// Only need to resolve the base variable; field access is checked during type checking
		nr.ResolveExpression(expr.Values[0], table)
	case *ast.BinOp:
		nr.ResolveExpression(expr.LValue, table)
		nr.ResolveExpression(expr.RValue, table)
	case *ast.UnaryOp:
		nr.ResolveExpression(expr.RValue, table)
	case *ast.Allocate:
		if _, contains := table.Contains(expr.Name.String()); !contains {
			msg := fmt.Sprintf("new expects expression of type struct but found undefined \"%v\" type", expr.Name.String())
			nr.addError(expr.Line, expr.Column, msg)
		}
	case *ast.Call:
		// Functions are always in the global scope
		nr.ResolveExpression(expr.Name, nr.tables.Globals)
	case *ast.Selector:
		nr.ResolveExpression(expr.Target, table)
	default:
		return
	}
}

func (nr *nameResolver) ResolveStatement(astStmt ast.Statement, table *st.SymbolTable[st.TypeEntry]) {
	switch stmt := astStmt.(type) {

	case *ast.Assignment:
		nr.ResolveExpression(stmt.Target, table)
		nr.ResolveExpression(stmt.RValue, table)
	case *ast.Print:
		for _, expr := range stmt.Expressions {
			nr.ResolveExpression(expr, table)
		}
	case *ast.Read:
		nr.ResolveExpression(stmt.Target, table)
	case *ast.Delete:
		nr.ResolveExpression(stmt.Target, table)
	case *ast.Conditional:
		nr.ResolveExpression(stmt.Condition, table)
		for _, thenStmt := range stmt.ThenBlock {
			nr.ResolveStatement(thenStmt, table)
		}
		for _, elseStmt := range stmt.ElseBlock {
			nr.ResolveStatement(elseStmt, table)
		}
	case *ast.Loop:
		nr.ResolveExpression(stmt.Condition, table)
		for _, loopStmt := range stmt.LoopBlock {
			nr.ResolveStatement(loopStmt, table)
		}
	case *ast.Return:
		nr.ResolveExpression(stmt.RValue, table)
	case *ast.Invocation:
		nr.ResolveExpression(stmt.Name, table)
		for _, expr := range stmt.Arguments {
			nr.ResolveExpression(expr, table)
		}

	default:
		return
	}

}

func (nr *nameResolver) ResolveTypeDecl(decl *ast.TypeDecl) {
	name := decl.Name.String()

	newStruct := st.NewStructEntry(
		name,
		types.StringToType("struct", name),
		st.GLOBAL,
		make([]*st.VarEntry, 0),
		decl.Token,
	)

	if existing, ok := nr.tables.Globals.Insert(name, newStruct); !ok {
		msg := fmt.Sprintf("redeclaration of struct (%v). Other declaration at (%d,%d)", name, existing.GetToken().Line, existing.GetToken().Column)
		nr.addError(decl.Line, decl.Column, msg)
	}

	for _, field := range decl.Fields {
		fieldEntry := st.NewVarEntry(
			field.Name.String(),
			field.Type,
			st.LOCAL,
			field.Token,
		)

		if existing, ok := newStruct.LocalST.Insert(fieldEntry.Name, fieldEntry); !ok {
			msg := fmt.Sprintf("redeclaration of variable (%v). Other declaration at (%d,%d)", field.Name, existing.GetToken().Line, existing.GetToken().Column)
			nr.addError(field.Line, field.Column, msg)
		} else {
			newStruct.Fields = append(newStruct.Fields, fieldEntry)
		}
	}
}

func (nr *nameResolver) ResolveVarDecl(decl *ast.VarDecl) {
	for _, name := range decl.Names {
		newVar := st.NewVarEntry(
			name.String(),
			name.GetType(),
			st.GLOBAL,
			name.GetToken(),
		)

		if existing, ok := nr.tables.Globals.Insert(newVar.Name, newVar); !ok {
			msg := fmt.Sprintf("redeclaration of variable (%v). Other declaration at (%d,%d)", newVar.Name, existing.GetToken().Line, existing.GetToken().Column)
			nr.addError(newVar.Line, newVar.Column, msg)
		}
	}
}

func (nr *nameResolver) ResolveFuncDecl(decl *ast.FuncDecl) {
	name := decl.Name.String()

	newFunc := st.NewFuncEntry(
		name,
		decl.ReturnType,
		make([]*st.VarEntry, 0),
		st.GLOBAL,
		nr.tables.Globals,
		decl.Token,
	)

	if existing, ok := nr.tables.Globals.Insert(name, newFunc); !ok {
		msg := fmt.Sprintf("function with name \"%v\" already defined. Previous declaration at %d:%d", name, existing.GetToken().Line, existing.GetToken().Column)
		nr.addError(decl.Line, decl.Column, msg)
	}

	for _, param := range decl.Parameters {
		paramEntry := st.NewVarEntry(
			param.Name.String(),
			param.Type,
			st.LOCAL,
			param.Token,
		)

		if existing, ok := newFunc.LocalST.Insert(paramEntry.Name, paramEntry); !ok {
			msg := fmt.Sprintf("redeclaration of parameter \"%v\" in function arguments. Other declaration at (%d,%d)", param.Name, existing.GetToken().Line, existing.GetToken().Column)
			nr.addError(param.Line, param.Column, msg)
		} else {
			newFunc.Params = append(newFunc.Params, paramEntry)
		}
	}

	for _, localVar := range decl.LocalDecl {
		for _, name := range localVar.Names {
			newVar := st.NewVarEntry(
				name.String(),
				name.GetType(),
				st.LOCAL,
				name.GetToken(),
			)

			if existing, ok := newFunc.LocalST.Insert(newVar.Name, newVar); !ok {
				msg := fmt.Sprintf("redeclaration of variable (%v). Other declaration at (%d,%d)", newVar.Name, existing.GetToken().Line, existing.GetToken().Column)
				nr.addError(newVar.Line, newVar.Column, msg)
			}
		}
	}

	for _, stmt := range decl.Stmts {
		nr.ResolveStatement(stmt, newFunc.LocalST)
	}
}

// ResolveAST processes declarations in grammar order: types -> globals -> functions.
// This ordering matters because functions can reference types and globals.
func (nr *nameResolver) ResolveAST(program *ast.Program) {
	for _, typeDecl := range program.Types {
		nr.ResolveTypeDecl(typeDecl)
	}

	for _, varDecl := range program.Globals {
		nr.ResolveVarDecl(varDecl)
	}

	for _, funcDecl := range program.Functions {
		nr.ResolveFuncDecl(funcDecl)
	}

	// Language spec requires a main() with no args and no return
	entry, _ := nr.tables.Globals.Contains("main")
	funcEntry, ok := entry.(*st.FuncEntry)
	if !ok {
		msg := "No required function named \"main\" defined."
		nr.addError(0, 0, msg)
	} else {
		if len(funcEntry.Params) > 0 {
			msg := "No required function named \"main\" must not take in any arguments."
			nr.addError(0, 0, msg)
		}
		if !funcEntry.ReturnTy.Equals(types.VoidTySig) {
			msg := "No required function named \"main\" must not return a type."
			nr.addError(0, 0, msg)
		}
	}

}
