package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type LValue struct {
	*token.Token
	Values []Expression
	Ty     types.Type
}

func NewLValue(target []Expression, token *token.Token) Expression {
	return &LValue{token, target, nil}
}

func (l *LValue) String() string {
	var out bytes.Buffer

	for i, id := range l.Values {
		if i > 0 {
			out.WriteString(".")
		}
		out.WriteString(id.String())
	}

	return out.String()
}

func (l *LValue) GetToken() *token.Token {
	return l.Token
}

func (l *LValue) GetType() types.Type {
	return l.Ty
}
