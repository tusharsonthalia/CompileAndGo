package sa

import (
	"fmt"
	"golite/ast"
	"golite/context"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	"strings"
)

type TypeChecker interface {
	addError(line, col int, msg string)
	typeCheckHelper(ty types.Type, token *token.Token)
	StmtGuaranteesReturn(stmt ast.Statement) bool
	HasErrors() bool
	TypeCheckLValueChain(ast.Expression, []string, bool, *st.SymbolTable[st.TypeEntry]) types.Type
	TypeCheckExpression(ast.Expression, *st.SymbolTable[st.TypeEntry])
	TypeCheckStatement(ast.Statement, *st.FuncEntry)
	TypeCheckTypeDecl(*ast.TypeDecl)
	TypeCheckVarDecl(*ast.VarDecl)
	TypeCheckFuncDecl(*ast.FuncDecl)
	TypeCheckAST(*ast.Program)
}

type typeChecker struct {
	errors []*context.CompilerError
	tables *st.SymbolTables
}

func NewTypeChecker(errors []*context.CompilerError, tables *st.SymbolTables) TypeChecker {
	return &typeChecker{errors, tables}
}

func (tc *typeChecker) addError(line, col int, msg string) {
	tc.errors = append(tc.errors, &context.CompilerError{
		Line:  line,
		Col:   col,
		Msg:   msg,
		Phase: context.SEMANTIC,
	})
}

func (tc *typeChecker) typeCheckHelper(ty types.Type, token *token.Token) {
	switch ty {
	case types.IntTySig:
		return
	case types.BoolTySig:
		return
	default: // Handling pointer to struct type checking
		ptr, ok := ty.(*types.PtrStructTy)
		if !ok {
			msg := fmt.Sprintf("invalid type defined (%v).", ty.String())
			tc.addError(token.Line, token.Column, msg)
			return
		}

		typeName := ptr.Base.Name
		val, ok := tc.tables.Globals.Contains(typeName)
		if !ok {
			msg := fmt.Sprintf("use of undefined type (%v).", typeName)
			tc.addError(token.Line, token.Column, msg)
			return
		}

		if !val.GetType().Equals(ptr.Base) {
			msg := fmt.Sprintf("(%v) is not a struct type.", typeName)
			tc.addError(token.Line, token.Column, msg)
			return
		}
	}
}

func (tc *typeChecker) BlockGuaranteesReturn(stmts []ast.Statement) bool {
	for _, stmt := range stmts {
		if tc.StmtGuaranteesReturn(stmt) {
			return true
		}
	}
	return false
}

// StmtGuaranteesReturn checks if a statement unconditionally returns.
// Used to verify that non-void functions always return on every path.
func (tc *typeChecker) StmtGuaranteesReturn(stmt ast.Statement) bool {
	switch s := stmt.(type) {
	case *ast.Return:
		return true

	case *ast.Conditional:
		// An if/else guarantees a return only if BOTH branches do
		if len(s.ElseBlock) == 0 {
			return false
		}
		return tc.BlockGuaranteesReturn(s.ThenBlock) && tc.BlockGuaranteesReturn(s.ElseBlock)

	case *ast.Loop:
		// Can't guarantee a loop body executes, so loops don't count
		return false

	default:
		return false
	}
}

func (tc *typeChecker) HasErrors() bool {
	return context.HasErrors(tc.errors)
}

func (tc *typeChecker) TypeCheckLValueChain(baseExpr ast.Expression, fields []string, selectorTerm bool, table *st.SymbolTable[st.TypeEntry]) types.Type {
	tc.TypeCheckExpression(baseExpr, table)
	currField := baseExpr.String()
	currType := baseExpr.GetType()

	for _, field := range fields {
		ptrStruct, ok := currType.(*types.PtrStructTy)
		if !ok {
			msg := fmt.Sprintf("field \"%v\" type's not a defined struct type.", currField)
			tc.addError(baseExpr.GetToken().Line, baseExpr.GetToken().Column, msg)
			return types.UnknownTySig
		}

		s, _ := tc.tables.Globals.Contains(ptrStruct.Base.Name)
		structEntry, ok := s.(*st.StructEntry)
		if !ok {
			msg := fmt.Sprintf("struct type (%v) not defined", ptrStruct.Base.Name)
			tc.addError(baseExpr.GetToken().Line, baseExpr.GetToken().Column, msg)
			return types.UnknownTySig
		}

		found := false
		for _, f := range structEntry.Fields {
			if f.Name == field {
				currField = field
				currType = f.GetType()
				found = true
				break
			}
		}

		if !found {
			msg := ""
			if selectorTerm {
				msg = fmt.Sprintf("Did not find field: %v inside %v struct", field, structEntry.GetType())
			} else {
				msg = fmt.Sprintf("field \"%v\" not defined in \"%v\" definition.", field, structEntry.GetType())
			}
			tc.addError(baseExpr.GetToken().Line, baseExpr.GetToken().Column, msg)
			return types.UnknownTySig
		}

	}

	return currType
}

func (tc *typeChecker) TypeCheckExpression(astExpr ast.Expression, table *st.SymbolTable[st.TypeEntry]) {
	switch expr := astExpr.(type) {

	case *ast.Variable:
		entry, ok := table.Contains(expr.Name)
		if !ok {
			msg := fmt.Sprintf("undefined variable of type (%v)", expr.Name)
			tc.addError(expr.Line, expr.Column, msg)
			expr.Ty = types.UnknownTySig
		} else {
			expr.Ty = entry.GetType()
		}
	case *ast.LValue:
		names := make([]string, len(expr.Values)-1)
		for i, v := range expr.Values[1:] {
			names[i] = v.String()
		}
		expr.Ty = tc.TypeCheckLValueChain(expr.Values[0], names, false, table)
	case *ast.BinOp:
		tc.TypeCheckExpression(expr.LValue, table)
		tc.TypeCheckExpression(expr.RValue, table)

		leftType := expr.LValue.GetType()
		rightType := expr.RValue.GetType()

		switch expr.Op {
		case ast.PLUS, ast.MINUS, ast.ASTERISK, ast.FSLASH:
			if leftType.Equals(types.IntTySig) && leftType.Equals(rightType) {
				expr.Ty = types.IntTySig
			} else {
				expr.Ty = types.UnknownTySig
				msg := fmt.Sprintf("left-operand of type \"%v\" does not match right-operand of type \"%v\"", leftType, rightType)
				tc.addError(expr.Line, expr.Column, msg)
			}
		case ast.GT, ast.GEQ, ast.LT, ast.LEQ:
			if leftType.Equals(types.IntTySig) && leftType.Equals(rightType) {
				expr.Ty = types.BoolTySig
			} else {
				expr.Ty = types.UnknownTySig
				msg := fmt.Sprintf("left-operand of type \"%v\" does not match right-operand of type \"%v\"", leftType, rightType)
				tc.addError(expr.Line, expr.Column, msg)
			}
		case ast.DOUBLEEQ, ast.NEQ:
			if leftType.Equals(rightType) {
				expr.Ty = types.BoolTySig
			} else {
				if _, ok := leftType.(*types.PtrStructTy); ok && rightType.Equals(types.NilTySig) {
					expr.Ty = types.BoolTySig
				} else if _, ok := rightType.(*types.PtrStructTy); ok && leftType.Equals(types.NilTySig) {
					expr.Ty = types.BoolTySig
				} else {
					expr.Ty = types.UnknownTySig
					msg := fmt.Sprintf("left-operand of type \"%v\" does not match right-operand of type \"%v\"", leftType, rightType)
					tc.addError(expr.Line, expr.Column, msg)
				}
			}
		case ast.AND, ast.OR:
			if leftType.Equals(types.BoolTySig) && leftType.Equals(rightType) {
				expr.Ty = types.BoolTySig
			} else {
				expr.Ty = types.UnknownTySig
				msg := fmt.Sprintf("left-operand of type \"%v\" does not match right-operand of type \"%v\"", leftType, rightType)
				tc.addError(expr.Line, expr.Column, msg)
			}
		default:
			expr.Ty = types.UnknownTySig
			msg := fmt.Sprintf("unknown binary operator %v", ast.OpToStr(expr.Op))
			tc.addError(expr.Line, expr.Column, msg)

		}
	case *ast.UnaryOp:
		tc.TypeCheckExpression(expr.RValue, table)
		rightType := expr.RValue.GetType()
		op := expr.Op
		if rightType.Equals(types.BoolTySig) && op == ast.EXCLAMATION {
			expr.Ty = types.BoolTySig
		} else if rightType.Equals(types.IntTySig) && op == ast.MINUS {
			expr.Ty = types.IntTySig
		} else {
			expr.Ty = types.UnknownTySig
			msg := fmt.Sprintf("Incompatible Operand (%v) with type (%v).", ast.OpToStr(op), expr.RValue.GetType())
			tc.addError(expr.Line, expr.Column, msg)
		}
	case *ast.Allocate:
		tc.TypeCheckExpression(expr.Name, table)
		expr.Ty = types.StringToType("*"+expr.GetType().String(), "")
	case *ast.Call:
		function, ok := tc.tables.Globals.Contains(expr.Name.String())
		if !ok {
			tc.addError(expr.Line, expr.Column, fmt.Sprintf("undefined function (%v)", expr.Name.String()))
			expr.Ty = types.UnknownTySig
			return
		}

		funcEntry, ok := function.(*st.FuncEntry)
		if !ok {
			msg := fmt.Sprintf("undefined function call named \"(%v)\"", expr.Name.String())
			tc.addError(expr.Line, expr.Column, msg)
			expr.Ty = types.UnknownTySig
			return
		}

		if len(expr.Arguments) != len(funcEntry.Params) {
			msg := fmt.Sprintf("function expects %d arguments, got %d", len(funcEntry.Params), len(expr.Arguments))
			tc.addError(expr.Line, expr.Column, msg)
		}

		for i, arg := range expr.Arguments {
			tc.TypeCheckExpression(arg, table)
			param := funcEntry.Params[i]
			if !arg.GetType().Equals(param.GetType()) {
				msg := fmt.Sprintf("argument %d has type (%v) but function expects type (%v)", i, arg.GetType(), param.GetType())
				tc.addError(expr.Line, expr.Column, msg)
			}
		}
		tc.TypeCheckExpression(expr.Name, tc.tables.Globals)

		expr.Ty = funcEntry.ReturnTy
	case *ast.Selector:
		expr.Ty = tc.TypeCheckLValueChain(expr.Target, expr.Accessors, true, table)
	default:
		return
	}
}

func (tc *typeChecker) TypeCheckStatement(astStmt ast.Statement, funcEntry *st.FuncEntry) {
	switch stmt := astStmt.(type) {
	case *ast.Assignment:
		tc.TypeCheckExpression(stmt.Target, funcEntry.LocalST)
		tc.TypeCheckExpression(stmt.RValue, funcEntry.LocalST)

		leftType := stmt.Target.GetType()
		rightType := stmt.RValue.GetType()

		if leftType.Equals(rightType) {
			return
		}

		if rightType.Equals(types.NilTySig) {
			if _, ok := leftType.(*types.PtrStructTy); ok {
				return
			}
		}

		msg := fmt.Sprintf("left-hand side value of type \"%v\" does not match its expression of type \"%v\"", stmt.Target.GetType(), stmt.RValue.GetType())
		tc.addError(stmt.Line, stmt.Column, msg)
	case *ast.Print:
		for _, expr := range stmt.Expressions {
			tc.TypeCheckExpression(expr, funcEntry.LocalST)
			if !expr.GetType().Equals(types.IntTySig) {
				msg := fmt.Sprintf("printf expects expression of type int but found %v type", expr.GetType())
				tc.addError(stmt.Line, stmt.Column, msg)
			}
		}
		placeholderCount := strings.Count(stmt.Target, "%d")
		if placeholderCount != len(stmt.Expressions) {
			msg := fmt.Sprintf("found %d placeholders but got only %d arguments", placeholderCount, len(stmt.Expressions))
			tc.addError(stmt.Line, stmt.Column, msg)
		}
	case *ast.Read:
		tc.TypeCheckExpression(stmt.Target, funcEntry.LocalST)
		targetType := stmt.Target.GetType()

		if !targetType.Equals(types.IntTySig) {
			msg := fmt.Sprintf("cannot scan into type (%v)", targetType)
			tc.addError(stmt.Line, stmt.Column, msg)
		}
	case *ast.Delete:
		tc.TypeCheckExpression(stmt.Target, funcEntry.LocalST)
		targetType := stmt.Target.GetType()
		if _, ok := targetType.(*types.PtrStructTy); !ok {
			msg := fmt.Sprintf("delete statement requires expression to be a struct pointer but found \"(%v)\" type.", stmt.Target.GetType())
			tc.addError(stmt.Line, stmt.Column, msg)
		}
		tc.typeCheckHelper(targetType, stmt.Token)
	case *ast.Conditional:
		tc.TypeCheckExpression(stmt.Condition, funcEntry.LocalST)
		if !stmt.Condition.GetType().Equals(types.BoolTySig) {
			msg := fmt.Sprintf("if conditional expression expects bool type but found %v type", stmt.Condition.GetType())
			tc.addError(stmt.Line, stmt.Column, msg)
		}
		for _, thenStmt := range stmt.ThenBlock {
			tc.TypeCheckStatement(thenStmt, funcEntry)
		}
		for _, elseStmt := range stmt.ElseBlock {
			tc.TypeCheckStatement(elseStmt, funcEntry)
		}
	case *ast.Loop:
		tc.TypeCheckExpression(stmt.Condition, funcEntry.LocalST)
		if !stmt.Condition.GetType().Equals(types.BoolTySig) {
			msg := fmt.Sprintf("condition expression must be a boolean. Found %v", stmt.Condition.GetType())
			tc.addError(stmt.Line, stmt.Column, msg)
		}
		for _, loopStmt := range stmt.LoopBlock {
			tc.TypeCheckStatement(loopStmt, funcEntry)
		}
	case *ast.Return:
		if stmt.RValue != nil {
			tc.TypeCheckExpression(stmt.RValue, funcEntry.LocalST)
			returnedType := stmt.RValue.GetType()

			if funcEntry.ReturnTy.Equals(types.VoidTySig) {
				msg := fmt.Sprintf("return statement returns value of type \"%v\" but function does not return a type", returnedType)
				tc.addError(stmt.Line, stmt.Column, msg)
			} else if !funcEntry.ReturnTy.Equals(returnedType) {
				msg := fmt.Sprintf("return statement returns value of type \"%v\" but function returns type:\"%v\"", returnedType, funcEntry.ReturnTy)
				tc.addError(stmt.Line, stmt.Column, msg)
			}
		} else {
			if !funcEntry.ReturnTy.Equals(types.VoidTySig) {
				msg := fmt.Sprintf("return statement requires a return value of type \"%v\"", funcEntry.ReturnTy)
				tc.addError(stmt.Line, stmt.Column, msg)
			}
		}
	case *ast.Invocation:
		tc.TypeCheckExpression(stmt.Name, funcEntry.LocalST)
		for _, expr := range stmt.Arguments {
			tc.TypeCheckExpression(expr, funcEntry.LocalST)
		}
	default:
		return
	}

}

func (tc *typeChecker) TypeCheckTypeDecl(decl *ast.TypeDecl) {
	for _, field := range decl.Fields {
		tc.typeCheckHelper(field.Type, field.Token)
	}
}

func (tc *typeChecker) TypeCheckVarDecl(decl *ast.VarDecl) {
	tc.typeCheckHelper(decl.Type, decl.Token)
}

func (tc *typeChecker) TypeCheckFuncDecl(decl *ast.FuncDecl) {
	for _, param := range decl.Parameters {
		tc.typeCheckHelper(param.Type, param.Token)
	}

	for _, localDecl := range decl.LocalDecl {
		tc.typeCheckHelper(localDecl.Type, localDecl.Token)
	}

	funcSymbolTable, _ := tc.tables.Globals.Contains(decl.Name.String())
	funcEntry := funcSymbolTable.(*st.FuncEntry)
	for _, stmt := range decl.Stmts {
		tc.TypeCheckStatement(stmt, funcEntry)
	}

	if !funcEntry.ReturnTy.Equals(types.VoidTySig) {
		if !tc.BlockGuaranteesReturn(decl.Stmts) {
			msg := fmt.Sprintf("Not all possible control flow paths in \"%v\" return.", decl.Name.String())
			tc.addError(decl.Line, decl.Column, msg)
		}
	}
}

func (tc *typeChecker) TypeCheckAST(program *ast.Program) {
	for _, typeDecl := range program.Types {
		tc.TypeCheckTypeDecl(typeDecl)
	}

	for _, varDecl := range program.Globals {
		tc.TypeCheckVarDecl(varDecl)
	}

	for _, funcDecl := range program.Functions {
		tc.TypeCheckFuncDecl(funcDecl)
	}

}
