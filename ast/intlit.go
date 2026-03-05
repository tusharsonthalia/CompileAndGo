package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
	"strconv"
)

type IntLit struct {
	*token.Token
	Value int64
	Ty    types.Type
}

func NewIntLit(value int64, token *token.Token) Expression {
	return &IntLit{token, value, types.IntTySig}
}

func (i *IntLit) GetType() types.Type {
	return i.Ty
}

func (i *IntLit) GetToken() *token.Token {
	return i.Token
}

func (i *IntLit) String() string {
	var out bytes.Buffer

	out.WriteString(strconv.FormatInt(i.Value, 10))

	return out.String()
}
