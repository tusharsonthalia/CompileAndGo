package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type Selector struct {
	*token.Token
	Target    Expression
	Accessors []string
	Ty        types.Type
}

func NewSelector(target Expression, accessors []string, ty types.Type, token *token.Token) Expression {
	return &Selector{token, target, accessors, ty}
}

func (s *Selector) GetToken() *token.Token {
	return s.Token
}

func (s *Selector) GetType() types.Type {
	return s.Ty
}

func (s *Selector) String() string {
	var out bytes.Buffer

	out.WriteString(s.Target.String())
	for _, id := range s.Accessors {
		out.WriteString(".")
		out.WriteString(id)
	}

	return out.String()
}
