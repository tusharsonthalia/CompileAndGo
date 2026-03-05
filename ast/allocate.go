package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type Allocate struct {
	*token.Token
	Name Expression
	Ty   types.Type
}

func NewAllocate(target Expression, token *token.Token) Expression {
	return &Allocate{token, target, target.GetType()}
}

func (a *Allocate) GetType() types.Type {
	return a.Ty
}
func (a *Allocate) GetToken() *token.Token {
	return a.Token
}

func (a *Allocate) String() string {
	var out bytes.Buffer

	out.WriteString("new ")
	out.WriteString(a.Name.String())

	return out.String()
}
