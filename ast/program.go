package ast

import (
	"bytes"
	"golite/token"
)

type Program struct {
	*token.Token
	Types     []*TypeDecl
	Globals   []*VarDecl
	Functions []*FuncDecl
}

func NewProgram(types []*TypeDecl, globals []*VarDecl, functions []*FuncDecl, token *token.Token) *Program {
	return &Program{token, types, globals, functions}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, types := range p.Types {
		out.WriteString(types.String())
		out.WriteString("\n")
	}
	out.WriteString("\n")

	for _, globals := range p.Globals {
		out.WriteString(globals.String())
		out.WriteString("\n")
	}
	out.WriteString("\n")

	for _, functions := range p.Functions {
		out.WriteString(functions.String())
	}
	out.WriteString("\n")

	return out.String()
}
