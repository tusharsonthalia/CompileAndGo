package ast

import (
	"golite/token"
	"golite/types"
)

type NilLit struct {
	*token.Token
	Name string
	Ty   types.Type
}

func NewNilLit(name string, token *token.Token) Expression {
	return &NilLit{token, name, types.NilTySig}
}

func (n *NilLit) GetToken() *token.Token {
	return n.Token
}

func (n *NilLit) GetType() types.Type {
	return n.Ty
}

func (n *NilLit) String() string {
	return n.Name
}
