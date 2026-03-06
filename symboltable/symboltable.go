package symboltable

import (
	"bytes"
	"fmt"
	"golite/token"
	"golite/types"
)

type VarScope int

const (
	GLOBAL VarScope = iota
	LOCAL
)

// TypeEntry is the interface for all vars, funcs and structs
type TypeEntry interface {
	String() string
	GetName() string
	GetToken() *token.Token
	GetType() types.Type
}

// ======================
//      Struct Entry
// ======================

type StructEntry struct {
	*token.Token
	Name    string
	Ty      types.Type
	Scope   VarScope
	Fields  []*VarEntry
	LocalST *SymbolTable[TypeEntry]
}

func NewStructEntry(name string, ty types.Type, scope VarScope, fields []*VarEntry, token *token.Token) *StructEntry {
	st := NewSymbolTable[TypeEntry](nil)
	return &StructEntry{token, name, ty, scope, fields, st}
}

func (s *StructEntry) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("Struct: %s\n", s.Name))
	for _, field := range s.Fields {
		out.WriteString(fmt.Sprintf("  Field: %s\n", field.String()))
	}
	return out.String()
}

func (s *StructEntry) GetName() string        { return s.Name }
func (s *StructEntry) GetToken() *token.Token { return s.Token }
func (s *StructEntry) GetType() types.Type    { return s.Ty }

// ======================
//       Var Entry
// ======================

type VarEntry struct {
	*token.Token
	Name  string
	Ty    types.Type
	Scope VarScope
}

func NewVarEntry(name string, ty types.Type, scope VarScope, token *token.Token) *VarEntry {
	return &VarEntry{token, name, ty, scope}
}

func (v *VarEntry) String() string {
	return fmt.Sprintf("Var: %s [%s] Scope: %d", v.Name, v.Ty, v.Scope)
}

func (v *VarEntry) GetName() string        { return v.Name }
func (v *VarEntry) GetToken() *token.Token { return v.Token }
func (v *VarEntry) GetType() types.Type    { return v.Ty }

// ======================
//       Func Entry
// ======================

type FuncEntry struct {
	*token.Token
	Name     string
	ReturnTy types.Type
	Params   []*VarEntry
	Scope    VarScope
	LocalST  *SymbolTable[TypeEntry]
}

func NewFuncEntry(name string, returnTy types.Type, params []*VarEntry, scope VarScope, parent *SymbolTable[TypeEntry], token *token.Token) *FuncEntry {
	st := NewSymbolTable(parent)
	return &FuncEntry{token, name, returnTy, params, scope, st}
}

func (f *FuncEntry) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("Func: %s (Ret: %s)\n", f.Name, f.ReturnTy))
	for _, param := range f.Params {
		out.WriteString(fmt.Sprintf("  Param: %s\n", param.String()))
	}
	return out.String()
}

func (f *FuncEntry) GetName() string        { return f.Name }
func (f *FuncEntry) GetToken() *token.Token { return f.Token }
func (f *FuncEntry) GetType() types.Type    { return f.ReturnTy }

// ==================================
//        SYMBOL TABLES (Global)
// ==================================

type SymbolTables struct {
	Globals *SymbolTable[TypeEntry]
}

func NewSymbolTables() *SymbolTables {
	return &SymbolTables{NewSymbolTable[TypeEntry](nil)}
}

// ==============================
//      Generic Symbol Table
// ==============================

// SymbolTable is a generic, scope-chained table. Each table points to its parent,
// so lookups walk up the chain (local -> function -> global) until found.
type SymbolTable[T TypeEntry] struct {
	parent *SymbolTable[T]
	table  map[string]T
}

func NewSymbolTable[T TypeEntry](parent *SymbolTable[T]) *SymbolTable[T] {
	return &SymbolTable[T]{
		parent,
		make(map[string]T),
	}
}

// Insert adds an entry to the current scope only. Returns false if already declared
// (redeclaration error -- we don't overwrite, caller decides what to do).
func (st *SymbolTable[T]) Insert(id string, entry T) (T, bool) {
	if existing, ok := st.table[id]; ok {
		return existing, false
	}

	st.table[id] = entry
	var empty T
	return empty, true

}

// Contains walks up the scope chain looking for the identifier.
// This is what gives us lexical scoping: locals shadow globals.
func (st *SymbolTable[T]) Contains(id string) (T, bool) {
	for cur := st; cur != nil; cur = cur.parent {
		if v, ok := cur.table[id]; ok {
			return v, true
		}
	}
	var empty T
	return empty, false
}
