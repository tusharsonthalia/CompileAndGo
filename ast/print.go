package ast

import (
	"bytes"
	"golite/token"
)

type Print struct {
	*token.Token
	Target      string
	Expressions []Expression
}

func NewPrint(target string, expressions []Expression, token *token.Token) Statement {
	return &Print{token, target, expressions}
}

func (p *Print) String() string {
	var out bytes.Buffer

	out.WriteString("printf(")
	out.WriteString(p.Target)
	for _, expr := range p.Expressions {
		out.WriteString(", ")
		out.WriteString(expr.String())
	}
	out.WriteString(");\n")

	return out.String()
}
