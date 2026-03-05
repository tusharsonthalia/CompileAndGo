package ast

import (
	"bytes"
	"golite/token"
)

type Delete struct {
	*token.Token
	Target Expression
}

func NewDelete(target Expression, token *token.Token) Statement {
	return &Delete{token, target}
}

func (p *Delete) String() string {
	var out bytes.Buffer

	out.WriteString("delete ")
	out.WriteString(p.Target.String())
	out.WriteString(";\n")

	return out.String()
}
