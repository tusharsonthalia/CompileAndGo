package ast

import (
	"bytes"
	"golite/token"
)

type Read struct {
	*token.Token
	Target Expression
}

func NewRead(target Expression, token *token.Token) Statement {
	return &Read{token, target}
}

func (p *Read) String() string {
	var out bytes.Buffer

	out.WriteString("scan ")
	out.WriteString(p.Target.String())
	out.WriteString(";\n")

	return out.String()
}
