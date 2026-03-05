package ir

import (
	"fmt"
	"golite/types"
	"strings"
)

type Value interface {
	String() string
	GetType() types.Type
}

type Register struct {
	ID              int
	Ty              types.Type
	allocatedPhyReg string
}

func (r *Register) String() string {
	if r.allocatedPhyReg != "" {
		return fmt.Sprintf("%s %%t%d", r.allocatedPhyReg, r.ID)
	}
	return fmt.Sprintf("%%t%d", r.ID)
}

func (r *Register) GetType() types.Type {
	return r.Ty
}

func (r *Register) AllocatedPhyReg() string {
	return r.allocatedPhyReg
}

type Constant struct {
	Val string
	Ty  types.Type
}

func (c *Constant) String() string {
	return c.Val
}

func (c *Constant) GetType() types.Type {
	return c.Ty
}

type PointerType struct {
	Base types.Type
}

func (p *PointerType) String() string {
	return "*" + p.Base.String()
}

func (p *PointerType) Equals(t types.Type) bool {
	other, ok := t.(*PointerType)
	return ok && p.Base.Equals(other.Base)
}

type ArrayType struct {
	Size int
	Base types.Type
}

func (a *ArrayType) String() string {
	return fmt.Sprintf("[%d x %s]", a.Size, LLVMType(a.Base))
}

func (a *ArrayType) GetType() types.Type {
	return a
}

func (a *ArrayType) Equals(t types.Type) bool {
	other, ok := t.(*ArrayType)
	return ok && a.Size == other.Size && a.Base.Equals(other.Base)
}

type I8Ty struct{}

func (i *I8Ty) String() string {
	return "i8"
}

func (i *I8Ty) Equals(t types.Type) bool {
	_, ok := t.(*I8Ty)
	return ok
}

func (i *I8Ty) GetType() types.Type {
	return i
}

type I32Ty struct{}

func (i *I32Ty) String() string {
	return "i32"
}

func (i *I32Ty) Equals(t types.Type) bool {
	_, ok := t.(*I32Ty)
	return ok
}

func (i *I32Ty) GetType() types.Type {
	return i
}

type Instruction interface {
	String() string
}

type BinaryOp struct {
	Op     string
	Result *Register
	Left   Value
	Right  Value
}

func (b *BinaryOp) String() string {
	return fmt.Sprintf("%s = %s %s %s, %s", b.Result.String(), b.Op, LLVMType(b.Result.GetType()), b.Left.String(), b.Right.String())
}

type Alloca struct {
	Result *Register
	Ty     types.Type
}

func (a *Alloca) String() string {
	return fmt.Sprintf("%s = alloca %s", a.Result.String(), LLVMType(a.Ty))
}

type Store struct {
	Src Value
	Dst Value
}

func (s *Store) String() string {
	return fmt.Sprintf("store %s %s, %s %s", LLVMType(s.Src.GetType()), s.Src.String(), LLVMType(s.Dst.GetType()), s.Dst.String())
}

type Load struct {
	Result *Register
	Src    Value
}

func (l *Load) String() string {
	return fmt.Sprintf("%s = load %s, %s %s", l.Result.String(), LLVMType(l.Result.GetType()), LLVMType(l.Src.GetType()), l.Src.String())
}

type Return struct {
	Val Value
}

func (r *Return) String() string {
	if r.Val == nil {
		return "ret void"
	}
	return fmt.Sprintf("ret %s %s", LLVMType(r.Val.GetType()), r.Val.String())
}

type Branch struct {
	Target string
}

func (b *Branch) String() string {
	return fmt.Sprintf("br label %%%s", b.Target)
}

type CondBranch struct {
	Cond   Value
	TrueL  string
	FalseL string
}

func (c *CondBranch) String() string {
	return fmt.Sprintf("br i1 %s, label %%%s, label %%%s", c.Cond.String(), c.TrueL, c.FalseL)
}

type Call struct {
	Result     *Register
	Name       string
	Args       []Value
	RetTy      types.Type
	IsVariadic bool
}

func (c *Call) String() string {
	argStrings := make([]string, len(c.Args))
	for i, arg := range c.Args {
		argStrings[i] = fmt.Sprintf("%s %s", LLVMType(arg.GetType()), arg.String())
	}
	args := strings.Join(argStrings, ", ")

	retTy := LLVMType(c.RetTy)
	signature := retTy
	if c.IsVariadic {

		signature = fmt.Sprintf("%s (i8*, ...)", retTy)
	}

	if c.Result == nil {
		return fmt.Sprintf("call %s @%s(%s)", signature, c.Name, args)
	}
	return fmt.Sprintf("%s = call %s @%s(%s)", c.Result.String(), signature, c.Name, args)
}

type Gep struct {
	Result *Register
	Ty     types.Type
	Ptr    Value
	Index  int
}

func (g *Gep) String() string {

	ptrTyStr := LLVMType(g.Ptr.GetType())

	if ptrTyStr == "i8*" {

	}
	return fmt.Sprintf("%s = getelementptr %s, %s %s, i32 0, i32 %d",
		g.Result.String(), LLVMType(g.Ty), ptrTyStr, g.Ptr.String(), g.Index)
}

type Bitcast struct {
	Result *Register
	Src    Value
}

func (b *Bitcast) String() string {
	return fmt.Sprintf("%s = bitcast %s %s to %s",
		b.Result.String(), LLVMType(b.Src.GetType()), b.Src.String(), LLVMType(b.Result.GetType()))
}

type Icmp struct {
	Cond   string
	Result *Register
	Left   Value
	Right  Value
}

func (i *Icmp) String() string {
	return fmt.Sprintf("%s = icmp %s %s %s, %s",
		i.Result.String(), i.Cond, LLVMType(i.Left.GetType()), i.Left.String(), i.Right.String())
}

type PhiEntry struct {
	Value Value
	Block string
}

type Phi struct {
	Result  *Register
	Entries []PhiEntry
}

func (p *Phi) String() string {
	parts := make([]string, len(p.Entries))
	for i, e := range p.Entries {
		parts[i] = fmt.Sprintf("[ %s, %%%s ]", e.Value.String(), e.Block)
	}
	return fmt.Sprintf("%s = phi %s %s", p.Result.String(), LLVMType(p.Result.GetType()), strings.Join(parts, ", "))
}

type Mov struct {
	Result *Register
	Src    Value
}

func (m *Mov) String() string {
	tyStr := LLVMType(m.Src.GetType())
	switch tyStr {
	case "i64":
		return fmt.Sprintf("%s = add %s %s, 0", m.Result.String(), tyStr, m.Src.String())
	case "i1":
		return fmt.Sprintf("%s = or i1 %s, false", m.Result.String(), m.Src.String())
	default:
		return fmt.Sprintf("%s = getelementptr i8, %s %s, i64 0", m.Result.String(), tyStr, m.Src.String())
	}
}

func LLVMType(t types.Type) string {
	switch ty := t.(type) {
	case *types.IntTy:
		return "i64"
	case *types.BoolTy:
		return "i1"
	case *types.StructTy:
		return "%struct." + ty.Name
	case *types.PtrStructTy:
		return "ptr"
	case *types.VoidTy:
		return "void"
	case *types.NilTy:
		return "ptr"
	case *PointerType:
		return "ptr"
	case *ArrayType:
		return fmt.Sprintf("[%d x %s]", ty.Size, LLVMType(ty.Base))
	case *I8Ty:
		return "i8"
	case *I32Ty:
		return "i32"
	default:
		return "i64"
	}
}
