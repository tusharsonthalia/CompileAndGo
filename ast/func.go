package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type FuncDecl struct {
	*token.Token
	Name       Expression
	Parameters []*Field
	ReturnType types.Type
	LocalDecl  []*VarDecl
	Stmts      []Statement
}

func NewFuncDecl(name Expression, parameters []*Field,
	returnType types.Type, localDecl []*VarDecl,
	stmts []Statement, token *token.Token) *FuncDecl {
	return &FuncDecl{token, name, parameters, returnType, localDecl, stmts}
}

func (f *FuncDecl) String() string {
	var out bytes.Buffer

	out.WriteString("func " + f.Name.String() + "(")
	for i, param := range f.Parameters {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(param.String())
	}

	out.WriteString(")")
	if f.ReturnType != types.VoidTySig {
		out.WriteString(" ")
		out.WriteString(f.ReturnType.String())
	}

	out.WriteString(" ")
	out.WriteString("{")
	out.WriteString("\n")

	for _, decl := range f.LocalDecl {
		out.WriteString(decl.String())
		out.WriteString("\n")
	}

	for _, stmt := range f.Stmts {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")
	out.WriteString("\n")

	return out.String()
}
