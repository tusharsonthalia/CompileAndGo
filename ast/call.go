package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type Call struct {
	*token.Token
	Name      Expression
	Arguments []Expression
	Ty        types.Type
}

func NewCall(name Expression, arguments []Expression, token *token.Token) Expression {
	return &Call{token, name, arguments, nil}
}

func (c *Call) GetType() types.Type {
	return c.Ty
}

func (c *Call) GetToken() *token.Token {
	return c.Token
}

func (i *Call) String() string {
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

	return out.String()
}
