package ast

import (
	"bytes"
	"golite/token"
)

type Invocation struct {
	*token.Token
	Name      Expression
	Arguments []Expression
}

func NewInvocation(name Expression, arguments []Expression, token *token.Token) Statement {
	return &Invocation{token, name, arguments}
}

func (i *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(i.Name.String())
	out.WriteString("(")
	for i, expr := range i.Arguments {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(expr.String())
	}
	out.WriteString(")")
	out.WriteString(";")
	out.WriteString("\n")

	return out.String()
}
