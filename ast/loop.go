package ast

import (
	"bytes"
	"golite/token"
)

type Loop struct {
	*token.Token
	Condition Expression
	LoopBlock []Statement
}

func NewLoop(condition Expression, loopBlock []Statement, token *token.Token) Statement {
	return &Loop{token, condition, loopBlock}
}

func (p *Loop) String() string {
	var out bytes.Buffer

	out.WriteString("for (")
	out.WriteString(p.Condition.String())
	out.WriteString(") {\n")
	for _, stmt := range p.LoopBlock {
		out.WriteString(stmt.String())
	}
	out.WriteString("}")
	out.WriteString("\n")

	return out.String()
}
