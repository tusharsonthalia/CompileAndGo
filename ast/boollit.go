package ast

import (
	"bytes"
	"fmt"
	"golite/token"
	"golite/types"
)

type BoolLit struct {
	*token.Token
	Value bool
	Ty    types.Type
}

func NewBoolLit(value bool, token *token.Token) Expression {
	return &BoolLit{token, value, types.BoolTySig}
}

func (b *BoolLit) GetType() types.Type {
	return b.Ty
}

func (b *BoolLit) GetToken() *token.Token {
	return b.Token
}

func (b *BoolLit) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("%v", b.Value))

	return out.String()
}
