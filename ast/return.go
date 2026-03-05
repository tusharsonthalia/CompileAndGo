package ast

import (
	"bytes"
	"golite/token"
)

type Return struct {
	*token.Token
	RValue Expression
}

func NewReturn(rvalue Expression, token *token.Token) Statement {
	return &Return{token, rvalue}
}

func (p *Return) String() string {
	var out bytes.Buffer

	out.WriteString("return")
	if p.RValue != nil {
		out.WriteString(" ")
		out.WriteString(p.RValue.String())
	}
	out.WriteString(";")
	out.WriteString("\n")

	return out.String()
}
