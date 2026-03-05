package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type UnaryOp struct {
	*token.Token
	Op     Operator
	RValue Expression
	Ty     types.Type
}

func NewUnaryOp(op Operator, rvalue Expression, token *token.Token) Expression {
	return &UnaryOp{token, op, rvalue, nil}
}

func (u *UnaryOp) GetToken() *token.Token {
	return u.Token
}

func (u *UnaryOp) GetType() types.Type {
	return u.Ty
}

func (u *UnaryOp) String() string {
	var out bytes.Buffer

	out.WriteString(OpToStr(u.Op))
	out.WriteString(u.RValue.String())

	return out.String()
}
