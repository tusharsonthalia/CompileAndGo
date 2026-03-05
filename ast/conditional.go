package ast

import (
	"bytes"
	"golite/token"
)

type Conditional struct {
	*token.Token
	Condition Expression
	ThenBlock []Statement
	ElseBlock []Statement
}

func NewConditional(condition Expression, thenBlock []Statement, elseBlock []Statement, token *token.Token) Statement {
	return &Conditional{token, condition, thenBlock, elseBlock}
}

func (p *Conditional) String() string {
	var out bytes.Buffer

	out.WriteString("if (")
	out.WriteString(p.Condition.String())
	out.WriteString(") {\n")
	for _, stmt := range p.ThenBlock {
		out.WriteString(stmt.String())
	}

	for i, stmt := range p.ElseBlock {
		if i == 0 {
			out.WriteString("} else {\n")
		}
		out.WriteString(stmt.String())
	}

	out.WriteString("}")
	out.WriteString("\n")

	return out.String()
}
