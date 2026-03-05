package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type Variable struct {
	*token.Token
	Name string
	Ty   types.Type
}

func NewVariable(target string, ty types.Type, token *token.Token) Expression {
	return &Variable{token, target, ty}
}

func (v *Variable) GetToken() *token.Token {
	return v.Token
}

func (v *Variable) GetType() types.Type {
	return v.Ty
}

func (v *Variable) String() string {
	var out bytes.Buffer

	out.WriteString(v.Name)

	return out.String()
}
