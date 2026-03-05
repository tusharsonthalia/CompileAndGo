package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type BinOp struct {
	*token.Token
	LValue Expression
	Op     Operator
	RValue Expression
	Ty     types.Type
}

func NewBinOp(lvalue Expression, op Operator, rvalue Expression, token *token.Token) Expression {
	return &BinOp{token, lvalue, op, rvalue, nil}
}

func (b *BinOp) GetToken() *token.Token {
	return b.Token
}

func (b *BinOp) GetType() types.Type {
	return b.Ty
}

func (b *BinOp) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(b.LValue.String())
	out.WriteString(" ")
	out.WriteString(OpToStr(b.Op))
	out.WriteString(" ")
	out.WriteString(b.RValue.String())
	out.WriteString(")")

	return out.String()
}
