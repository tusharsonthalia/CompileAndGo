package ast

import (
	"golite/token"
	"golite/types"
)

type Operator int

const (
	OR Operator = iota
	AND
	DOUBLEEQ
	NEQ
	GEQ
	LEQ
	EXCLAMATION
	FSLASH
	ASTERISK
	PLUS
	MINUS
	GT
	LT
	EQUALS
)

func StrToOp(op string) Operator {
	switch op {
	case "||":
		return OR
	case "&&":
		return AND
	case "==":
		return DOUBLEEQ
	case "!=":
		return NEQ
	case ">=":
		return GEQ
	case "<=":
		return LEQ
	case "!":
		return EXCLAMATION
	case "/":
		return FSLASH
	case "*":
		return ASTERISK
	case "+":
		return PLUS
	case "-":
		return MINUS
	case ">":
		return GT
	case "<":
		return LT
	case "=":
		return EQUALS
	default:
		panic("Invalid Operator String Encountered")
	}
}

func OpToStr(op Operator) string {
	switch op {
	case OR:
		return "||"
	case AND:
		return "&&"
	case DOUBLEEQ:
		return "=="
	case NEQ:
		return "!="
	case GEQ:
		return ">="
	case LEQ:
		return "<="
	case EXCLAMATION:
		return "!"
	case FSLASH:
		return "/"
	case ASTERISK:
		return "*"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case GT:
		return ">"
	case LT:
		return "<"
	case EQUALS:
		return "="
	default:
		panic("Invalid Operator Encountered")
	}
}

type Expression interface {
	String() string
	GetType() types.Type
	GetToken() *token.Token
}
