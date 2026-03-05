package types

import "strings"

type Type interface {
	String() string
	Equals(Type) bool
}

type IntTy struct{}

func (i *IntTy) String() string {
	return "int"
}

func (i *IntTy) Equals(t Type) bool {
	_, ok := t.(*IntTy)
	return ok
}

type BoolTy struct{}

func (b *BoolTy) String() string {
	return "bool"
}

func (b *BoolTy) Equals(t Type) bool {
	_, ok := t.(*BoolTy)
	return ok
}

type StructTy struct {
	Name string
}

func (s *StructTy) String() string {
	return s.Name
}

func (s *StructTy) Equals(t Type) bool {
	other, ok := t.(*StructTy)
	return ok && s.Name == other.Name
}

type PtrStructTy struct {
	Base *StructTy
}

func (p *PtrStructTy) String() string {
	return "*" + p.Base.Name
}

func (p *PtrStructTy) Equals(t Type) bool {
	other, ok := t.(*PtrStructTy)
	return ok && p.Base.Name == other.Base.Name
}

type NilTy struct{}

func (s *NilTy) String() string {
	return "nil"
}

func (i *NilTy) Equals(t Type) bool {
	_, ok := t.(*NilTy)
	return ok
}

type UnknownTy struct{}

func (u *UnknownTy) String() string {
	return "unknown"
}

func (u *UnknownTy) Equals(t Type) bool {
	_, ok := t.(*UnknownTy)
	return ok
}

type VoidTy struct{}

func (v *VoidTy) String() string {
	return "void"
}

func (v *VoidTy) Equals(t Type) bool {
	_, ok := t.(*VoidTy)
	return ok
}

type FunctionTy struct {
	Name string
}

func (f *FunctionTy) String() string {
	return f.Name
}

func (f *FunctionTy) Equals(t Type) bool {
	other, ok := t.(*FunctionTy)
	return ok && f.Name == other.Name
}

func StringToType(ty string, name string) Type {
	if strings.HasPrefix(ty, "*") {
		base := strings.TrimPrefix(ty, "*")

		if _, ok := StructTySig[base]; !ok {
			StructTySig[base] = &StructTy{Name: base}
		}

		return &PtrStructTy{Base: StructTySig[base]}
	}

	switch ty {
	case "int":
		return IntTySig
	case "bool":
		return BoolTySig
	case "nil":
		return NilTySig
	case "void":
		return VoidTySig
	case "unknown":
		return UnknownTySig
	case "function":
		if f, ok := FunctionTySig[name]; ok {
			return f
		}

		newTy := &FunctionTy{Name: name}
		FunctionTySig[name] = newTy
		return newTy
	case "struct":
		if s, ok := StructTySig[name]; ok {
			return s
		}

		newTy := &StructTy{Name: name}
		StructTySig[name] = newTy
		return newTy
	default:
		panic("Unknown type encountered!")
	}

}

var IntTySig *IntTy
var BoolTySig *BoolTy
var NilTySig *NilTy
var VoidTySig *VoidTy
var UnknownTySig *UnknownTy
var FunctionTySig map[string]*FunctionTy
var StructTySig map[string]*StructTy

func init() {
	IntTySig = &IntTy{}
	BoolTySig = &BoolTy{}
	NilTySig = &NilTy{}
	VoidTySig = &VoidTy{}
	UnknownTySig = &UnknownTy{}
	FunctionTySig = make(map[string]*FunctionTy)
	StructTySig = make(map[string]*StructTy)
}
