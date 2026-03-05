package ast

import (
	"bytes"
	"golite/token"
)

type Assignment struct {
	*token.Token
	Target Expression
	RValue Expression
}

func NewAssignment(target Expression, rvalue Expression, token *token.Token) Statement {
	return &Assignment{token, target, rvalue}
}

func (a *Assignment) String() string {
	var out bytes.Buffer

	out.WriteString(a.Target.String())
	out.WriteString(" = ")
	out.WriteString(a.RValue.String())
	out.WriteString(";\n")

	return out.String()
}
