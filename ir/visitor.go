package ir

import (
	"fmt"
	"golite/ast"
	st "golite/symboltable"
	"golite/types"
	"strings"
)

func (b *builder) visitExpression(expr ast.Expression) Value {
	if expr == nil {
		return nil
	}

	switch e := expr.(type) {
	case *ast.Selector:
		if len(e.Accessors) == 0 {
			return b.visitExpression(e.Target)
		}
		return b.visitSelectorExpression(e)
	case *ast.IntLit:
		return &Constant{Val: fmt.Sprintf("%d", e.Value), Ty: types.IntTySig}
	case *ast.BoolLit:
		val := "0"
		if e.Value {
			val = "1"
		}
		return &Constant{Val: val, Ty: types.BoolTySig}
	case *ast.NilLit:
		return &Constant{Val: "null", Ty: types.NilTySig}
	case *ast.Variable:
		return b.loadVariable(e)
	case *ast.BinOp:
		return b.visitBinOp(e)
	case *ast.UnaryOp:
		return b.visitUnaryOp(e)
	case *ast.Call:
		return b.visitCallExpression(e)
	case *ast.Allocate:
		return b.visitAllocate(e)
	case *ast.LValue:
		if len(e.Values) == 1 {
			return b.visitExpression(e.Values[0])
		}
		addr := b.visitSelectorAddress(e)
		res := b.nextRegister(e.GetType())
		b.currBlock.AddInstruction(&Load{Result: res, Src: addr})
		return res
	default:
		panic(fmt.Sprintf("unsupported expression type: %T", expr))
	}
}

func (b *builder) loadVariable(v *ast.Variable) Value {
	addr := b.getVariableAddress(v.Name, v.GetType())
	res := b.nextRegister(v.GetType())
	b.currBlock.AddInstruction(&Load{Result: res, Src: addr})
	return res
}

func (b *builder) getVariableAddress(name string, ty types.Type) Value {
	addr, ok := b.locals[name]
	if !ok {

		return &Constant{Val: "@" + name, Ty: &PointerType{Base: ty}}
	}
	return addr
}

func (b *builder) visitBinOp(e *ast.BinOp) Value {
	if e.Op == ast.AND || e.Op == ast.OR {
		return b.visitLogicalOp(e)
	}

	left := b.visitExpression(e.LValue)
	right := b.visitExpression(e.RValue)

	var op string
	switch e.Op {
	case ast.PLUS:
		op = "add"
	case ast.MINUS:
		op = "sub"
	case ast.ASTERISK:
		op = "mul"
	case ast.FSLASH:
		op = "sdiv"
	case ast.GT, ast.LT, ast.GEQ, ast.LEQ, ast.DOUBLEEQ, ast.NEQ:
		return b.visitComparison(e, left, right)
	default:
		panic(fmt.Sprintf("unsupported binop: %v", e.Op))
	}

	res := b.nextRegister(e.GetType())
	b.currBlock.AddInstruction(&BinaryOp{Op: op, Result: res, Left: left, Right: right})
	return res
}

func (b *builder) visitComparison(e *ast.BinOp, left, right Value) Value {
	var cond string
	switch e.Op {
	case ast.GT:
		cond = "sgt"
	case ast.LT:
		cond = "slt"
	case ast.GEQ:
		cond = "sge"
	case ast.LEQ:
		cond = "sle"
	case ast.DOUBLEEQ:
		cond = "eq"
	case ast.NEQ:
		cond = "ne"
	}

	res := b.nextRegister(types.BoolTySig)
	b.currBlock.AddInstruction(&Icmp{Cond: cond, Result: res, Left: left, Right: right})
	return res
}

func (b *builder) visitLogicalOp(e *ast.BinOp) Value {
	left := b.visitExpression(e.LValue)
	right := b.visitExpression(e.RValue)
	res := b.nextRegister(types.BoolTySig)
	op := "and"
	if e.Op == ast.OR {
		op = "or"
	}
	b.currBlock.AddInstruction(&BinaryOp{Op: op, Result: res, Left: left, Right: right})
	return res
}

func (b *builder) visitUnaryOp(e *ast.UnaryOp) Value {
	operand := b.visitExpression(e.RValue)
	res := b.nextRegister(e.GetType())

	switch e.Op {
	case ast.MINUS:
		b.currBlock.AddInstruction(&BinaryOp{Op: "sub", Result: res, Left: &Constant{Val: "0", Ty: types.IntTySig}, Right: operand})
	case ast.EXCLAMATION:
		b.currBlock.AddInstruction(&BinaryOp{Op: "xor", Result: res, Left: operand, Right: &Constant{Val: "1", Ty: types.BoolTySig}})
	}
	return res
}

func (b *builder) visitStatement(stmt ast.Statement) {
	if stmt == nil {
		return
	}

	switch s := stmt.(type) {
	case *ast.Assignment:
		b.visitAssignment(s)
	case *ast.Return:
		b.visitReturn(s)
	case *ast.Conditional:
		b.visitConditional(s)
	case *ast.Loop:
		b.visitLoop(s)
	case *ast.Print:
		b.visitPrint(s)
	case *ast.Read:
		b.visitRead(s)
	case *ast.Delete:
		b.visitDelete(s)
	case *ast.Invocation:
		b.visitInvocation(s)
	default:
		panic(fmt.Sprintf("unsupported statement type: %T", stmt))
	}
}

func (b *builder) visitAssignment(s *ast.Assignment) {
	rhs := b.visitExpression(s.RValue)
	addr := b.visitSelectorAddress(s.Target)
	b.currBlock.AddInstruction(&Store{Src: rhs, Dst: addr})
}

func (b *builder) visitReturn(s *ast.Return) {
	val := b.visitExpression(s.RValue)
	b.currBlock.AddInstruction(&Return{Val: val})
}

func (b *builder) visitConditional(s *ast.Conditional) {
	trueL := b.nextLabel()
	falseL := b.nextLabel()
	joinL := b.nextLabel()

	cond := b.visitExpression(s.Condition)
	b.currBlock.AddInstruction(&CondBranch{Cond: cond, TrueL: trueL, FalseL: falseL})

	b.currBlock = NewBasicBlock(trueL)
	b.currFunction.AddBlock(b.currBlock)
	for _, st := range s.ThenBlock {
		b.visitStatement(st)
	}
	if b.currBlock.Terminator == nil {
		b.currBlock.AddInstruction(&Branch{Target: joinL})
	}

	b.currBlock = NewBasicBlock(falseL)
	b.currFunction.AddBlock(b.currBlock)
	for _, st := range s.ElseBlock {
		b.visitStatement(st)
	}
	if b.currBlock.Terminator == nil {
		b.currBlock.AddInstruction(&Branch{Target: joinL})
	}

	b.currBlock = NewBasicBlock(joinL)
	b.currFunction.AddBlock(b.currBlock)
}

func (b *builder) visitLoop(s *ast.Loop) {
	condL := b.nextLabel()
	bodyL := b.nextLabel()
	exitL := b.nextLabel()

	b.currBlock.AddInstruction(&Branch{Target: condL})

	b.currBlock = NewBasicBlock(condL)
	b.currFunction.AddBlock(b.currBlock)
	cond := b.visitExpression(s.Condition)
	b.currBlock.AddInstruction(&CondBranch{Cond: cond, TrueL: bodyL, FalseL: exitL})

	b.currBlock = NewBasicBlock(bodyL)
	b.currFunction.AddBlock(b.currBlock)
	for _, st := range s.LoopBlock {
		b.visitStatement(st)
	}
	if b.currBlock.Terminator == nil {
		b.currBlock.AddInstruction(&Branch{Target: condL})
	}

	b.currBlock = NewBasicBlock(exitL)
	b.currFunction.AddBlock(b.currBlock)
}

func (b *builder) visitPrint(p *ast.Print) {
	format := p.Target
	fmtName := b.getOrCreateFmtString(format)

	var fmtGlobal *Global
	for _, g := range b.globals {
		if g.Name == fmtName {
			fmtGlobal = g
			break
		}
	}

	args := make([]Value, 0)
	fmtReg := b.nextRegister(types.NilTySig)
	b.currBlock.AddInstruction(&Gep{
		Result: fmtReg,
		Ty:     fmtGlobal.Ty,
		Ptr:    &Constant{Val: "@" + fmtName, Ty: &PointerType{Base: fmtGlobal.Ty}},
		Index:  0,
	})
	args = append(args, fmtReg)

	for _, expr := range p.Expressions {
		args = append(args, b.visitExpression(expr))
	}

	b.currBlock.AddInstruction(&Call{Name: "printf", Args: args, RetTy: &I32Ty{}, IsVariadic: true})
}

func (b *builder) getOrCreateFmtString(fmtStr string) string {
	if strings.HasPrefix(fmtStr, "\"") && strings.HasSuffix(fmtStr, "\"") {
		fmtStr = fmtStr[1 : len(fmtStr)-1]
	}

	if name, ok := b.fmtStrings[fmtStr]; ok {
		return name
	}
	name := fmt.Sprintf(".fmt%d", len(b.fmtStrings))
	b.fmtStrings[fmtStr] = name

	val := strings.ReplaceAll(fmtStr, "\\n", "\\0A")
	val = fmt.Sprintf("c\"%s\\00\"", val)

	size := 0
	for i := 0; i < len(fmtStr); i++ {
		if i+1 < len(fmtStr) && fmtStr[i] == '\\' && fmtStr[i+1] == 'n' {
			size++
			i++
		} else {
			size++
		}
	}
	size++

	b.globals = append(b.globals, &Global{
		Name:      name,
		Ty:        &ArrayType{Size: size, Base: &I8Ty{}},
		Value:     val,
		IsConst:   true,
		IsPrivate: true,
	})
	return name
}

func (b *builder) visitRead(r *ast.Read) {
	format := "%ld"
	fmtName := b.getOrCreateFmtString(format)

	var fmtGlobal *Global
	for _, g := range b.globals {
		if g.Name == fmtName {
			fmtGlobal = g
			break
		}
	}

	fmtReg := b.nextRegister(types.NilTySig)
	b.currBlock.AddInstruction(&Gep{
		Result: fmtReg,
		Ty:     fmtGlobal.Ty,
		Ptr:    &Constant{Val: "@" + fmtName, Ty: &PointerType{Base: fmtGlobal.Ty}},
		Index:  0,
	})

	addr := b.visitSelectorAddress(r.Target)
	b.currBlock.AddInstruction(&Call{Name: "scanf", Args: []Value{fmtReg, addr}, RetTy: &I32Ty{}, IsVariadic: true})
}

func (b *builder) visitDelete(s *ast.Delete) {
	val := b.visitExpression(s.Target)
	ptr := b.nextRegister(types.NilTySig)
	b.currBlock.AddInstruction(&Bitcast{Result: ptr, Src: val})
	b.currBlock.AddInstruction(&Call{Name: "free", Args: []Value{ptr}, RetTy: types.VoidTySig})
}

func (b *builder) visitSelectorExpression(e *ast.Selector) Value {
	addr := b.visitSelectorAddress(e)
	res := b.nextRegister(e.GetType())
	b.currBlock.AddInstruction(&Load{Result: res, Src: addr})
	return res
}

func (b *builder) visitSelectorAddress(e ast.Expression) Value {
	if s, ok := e.(*ast.Selector); ok && len(s.Accessors) == 0 {
		return b.visitSelectorAddress(s.Target)
	}

	if l, ok := e.(*ast.LValue); ok {
		if len(l.Values) == 1 {
			return b.visitSelectorAddress(l.Values[0])
		}

		currAddr := b.visitSelectorAddress(l.Values[0])
		for i := 1; i < len(l.Values); i++ {
			v, ok := l.Values[i].(*ast.Variable)
			if !ok {
				panic(fmt.Sprintf("LValue part %d is not a variable: %T", i, l.Values[i]))
			}
			currAddr = b.gepField(currAddr, v.Name)
		}
		return currAddr
	}

	switch s := e.(type) {
	case *ast.Selector:
		currAddr := b.visitSelectorAddress(s.Target)
		for _, acc := range s.Accessors {
			currAddr = b.gepField(currAddr, acc)
		}
		return currAddr

	case *ast.Variable:
		return b.getVariableAddress(s.Name, s.GetType())
	default:
		panic(fmt.Sprintf("invalid selector base: %T", e))
	}
}

func (b *builder) gepField(ptr Value, fieldName string) Value {

	ptrTy, ok := ptr.GetType().(*PointerType)
	if !ok {

		if _, ok := ptr.GetType().(*types.PtrStructTy); ok {
			ptrTy = &PointerType{Base: ptr.GetType()}
		} else {
			panic(fmt.Sprintf("gepField on non-pointer: %T", ptr.GetType()))
		}
	}

	currPtr := ptr
	baseTy := ptrTy.Base

	if ps, ok := baseTy.(*types.PtrStructTy); ok {
		res := b.nextRegister(ps)
		b.currBlock.AddInstruction(&Load{Result: res, Src: currPtr})
		currPtr = res
		baseTy = ps.Base
	}

	var structName string
	if st, ok := baseTy.(*types.StructTy); ok {
		structName = st.Name
	} else {
		panic(fmt.Sprintf("gepField: base is not a struct: %T", baseTy))
	}

	entry, _ := b.tables.Globals.Contains(structName)
	structEntry := entry.(*st.StructEntry)
	fieldIdx := -1
	var innerTy types.Type
	for i, field := range structEntry.Fields {
		if field.Name == fieldName {
			fieldIdx = i
			innerTy = field.Ty
			break
		}
	}

	if fieldIdx == -1 {
		panic(fmt.Sprintf("field %s not found in struct %s", fieldName, structName))
	}

	res := b.nextRegister(&PointerType{Base: innerTy})
	b.currBlock.AddInstruction(&Gep{Result: res, Ty: &types.StructTy{Name: structName}, Ptr: currPtr, Index: fieldIdx})
	return res
}

func (b *builder) visitInvocation(s *ast.Invocation) {
	b.visitCallExpression(&ast.Call{Name: s.Name, Arguments: s.Arguments, Ty: types.VoidTySig})
}

func (b *builder) visitCallExpression(e *ast.Call) Value {
	args := make([]Value, len(e.Arguments))
	for i, arg := range e.Arguments {
		args[i] = b.visitExpression(arg)
	}

	var res *Register
	if e.Ty != types.VoidTySig && e.Ty != nil {
		res = b.nextRegister(e.Ty)
	}

	name := e.Name.String()
	b.currBlock.AddInstruction(&Call{Result: res, Name: name, Args: args, RetTy: e.Ty})
	return res
}

func (b *builder) visitAllocate(e *ast.Allocate) Value {

	ptrStructTy := e.GetType().(*types.PtrStructTy)
	structName := ptrStructTy.Base.Name
	entry, _ := b.tables.Globals.Contains(structName)
	structEntry := entry.(*st.StructEntry)

	sizeBytes := len(structEntry.Fields) * 8
	size := &Constant{Val: fmt.Sprintf("%d", sizeBytes), Ty: &I32Ty{}}

	ptr := b.nextRegister(types.NilTySig)
	b.currBlock.AddInstruction(&Call{Result: ptr, Name: "malloc", Args: []Value{size}, RetTy: types.NilTySig})

	res := b.nextRegister(e.GetType())
	b.currBlock.AddInstruction(&Bitcast{Result: res, Src: ptr})
	return res
}
