package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type TypeDecl struct {
	*token.Token
	Name   Expression
	Fields []*Field
}

func NewTypeDecl(name Expression, fields []*Field, token *token.Token) *TypeDecl {
	return &TypeDecl{token, name, fields}
}

func (t *TypeDecl) String() string {
	var out bytes.Buffer
	out.WriteString("type " + t.Name.String() + " struct {\n")
	for _, field := range t.Fields {
		out.WriteString(field.String())
		out.WriteString(";\n")
	}
	out.WriteString("};")

	return out.String()
}

type Field struct {
	*token.Token
	Name Expression
	Type types.Type
}

func NewField(name Expression, t types.Type, token *token.Token) *Field {
	return &Field{token, name, t}
}

func (f *Field) String() string {
	var out bytes.Buffer
	out.WriteString(f.Name.String())
	out.WriteString(" ")
	out.WriteString(f.Type.String())

	return out.String()
}

type VarDecl struct {
	*token.Token
	Names []Expression
	Type  types.Type
}

func NewVarDecl(names []Expression, type_ types.Type, token *token.Token) *VarDecl {
	return &VarDecl{token, names, type_}
}

func (v *VarDecl) String() string {
	var out bytes.Buffer

	out.WriteString("var ")

	for i, field := range v.Names {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(field.String())
	}

	out.WriteString(" ")
	out.WriteString(v.Type.String())
	out.WriteString(";")

	return out.String()
}
