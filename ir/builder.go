package ir

import (
	"fmt"
	"golite/ast"
	st "golite/symboltable"
	"golite/types"
	"os"
)

type Builder interface {
	addStatement(stmt string)
	printBlocks()
	llvmType(t types.Type) string
	addTypeDecl(*st.StructEntry)
	BuildProgram(string, string, bool)
	SaveBlocks(string)
	Functions() []*Function
	Globals() []*Global
}

type builder struct {
	program      *ast.Program
	tables       *st.SymbolTables
	functions    []*Function
	globals      []*Global
	types        []string
	preamble     []string
	registerID   int
	labelID      int
	currFunction *Function
	currBlock    *BasicBlock
	locals       map[string]Value
	fmtStrings   map[string]string
}

func (b *builder) Functions() []*Function {
	return b.functions
}

func (b *builder) Globals() []*Global {
	return b.globals
}

func NewBuilder(program *ast.Program, tables *st.SymbolTables) Builder {
	return &builder{
		program:    program,
		tables:     tables,
		functions:  make([]*Function, 0),
		globals:    make([]*Global, 0),
		types:      make([]string, 0),
		preamble:   make([]string, 0),
		locals:     make(map[string]Value),
		fmtStrings: make(map[string]string),
	}
}

func (b *builder) nextRegister(ty types.Type) *Register {
	reg := &Register{ID: b.registerID, Ty: ty}
	b.registerID++
	return reg
}

func (b *builder) nextLabel() string {
	label := fmt.Sprintf("L%d", b.labelID)
	b.labelID++
	return label
}

func (b *builder) addStatement(stmt string) {
	b.preamble = append(b.preamble, stmt)
}

func (b *builder) printBlocks() {
	for _, s := range b.preamble {
		fmt.Println(s)
	}
	for _, t := range b.types {
		fmt.Println(t)
	}
	for _, g := range b.globals {
		fmt.Println(g.String())
	}
	for _, fn := range b.functions {
		fmt.Println(fn.String())
	}
}

func (b *builder) SaveBlocks(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Error encountered in creating .ll file: %v", err))
	}
	defer f.Close()

	for _, s := range b.preamble {
		f.WriteString(s + "\n")
	}
	for _, t := range b.types {
		f.WriteString(t + "\n")
	}
	if len(b.globals) > 0 {
		f.WriteString("\n")
		for _, g := range b.globals {
			f.WriteString(g.String() + "\n")
		}
	}
	for _, fn := range b.functions {
		f.WriteString("\n" + fn.String())
	}
}

func (b *builder) llvmType(t types.Type) string {
	return LLVMType(t)
}

func (b *builder) defaultVal(t types.Type) string {
	switch LLVMType(t) {
	case "i64", "i1":
		return "0"
	default:
		return "null"
	}
}

func (b *builder) addTypeDecl(entry *st.StructEntry) {
	stmt := fmt.Sprintf("%%struct.%s = type {", entry.Name)
	for i, field := range entry.Fields {
		if i > 0 {
			stmt += ", "
		}
		stmt += LLVMType(field.Ty)
	}
	stmt += "}"
	b.types = append(b.types, stmt)
}

func (b *builder) addVarDecl(entry *st.VarEntry) {
	g := &Global{
		Name:    entry.Name,
		Ty:      entry.Ty,
		Value:   b.defaultVal(entry.Ty),
		IsConst: false,
	}
	b.globals = append(b.globals, g)
	b.locals[entry.Name] = &Constant{Val: "@" + entry.Name, Ty: &PointerType{Base: entry.Ty}}
}

func (b *builder) addFuncEntry(entry *st.FuncEntry, decl *ast.FuncDecl) {
	b.registerID = 0
	b.labelID = 0
	b.locals = make(map[string]Value)

	fn := NewFunction(entry)
	b.currFunction = fn
	b.functions = append(b.functions, fn)

	entryBlock := NewBasicBlock("entry")
	b.currBlock = entryBlock
	fn.AddBlock(entryBlock)

	for _, param := range entry.Params {
		reg := b.nextRegister(&PointerType{Base: param.Ty})
		b.currBlock.AddInstruction(&Alloca{Result: reg, Ty: param.Ty})
		b.currBlock.AddInstruction(&Store{
			Src: &Constant{Val: "%" + param.Name, Ty: param.Ty},
			Dst: reg,
		})
		b.locals[param.Name] = reg
	}

	for _, localDecl := range decl.LocalDecl {
		for _, nameExpr := range localDecl.Names {
			name := nameExpr.String()
			localVar, ok := entry.LocalST.Contains(name)
			if !ok {
				panic(fmt.Sprintf("local variable %s not found in symbol table", name))
			}
			ty := localVar.GetType()
			reg := b.nextRegister(&PointerType{Base: ty})
			b.currBlock.AddInstruction(&Alloca{Result: reg, Ty: ty})
			b.currBlock.AddInstruction(&Store{
				Src: &Constant{Val: b.defaultVal(ty), Ty: ty},
				Dst: reg,
			})
			b.locals[name] = reg
		}
	}

	for _, stmt := range decl.Stmts {
		b.visitStatement(stmt)
	}

	if b.currBlock.Terminator == nil {
		if entry.ReturnTy == types.VoidTySig || entry.ReturnTy == nil {
			b.currBlock.AddInstruction(&Return{Val: nil})
		} else {
			b.currBlock.AddInstruction(&Return{Val: &Constant{Val: "0", Ty: entry.ReturnTy}})
		}
	}
}

func (b *builder) BuildProgram(filename string, target string, llvmStackMode bool) {
	b.addStatement(fmt.Sprintf("source_filename = \"%v\"", filename))
	b.addStatement(fmt.Sprintf("target triple = \"%v\"", target))

	for _, typeDecl := range b.program.Types {
		entryST, _ := b.tables.Globals.Contains(typeDecl.Name.String())
		structEntry, _ := entryST.(*st.StructEntry)
		b.addTypeDecl(structEntry)
	}

	for _, varDecl := range b.program.Globals {
		for _, singleVarDecl := range varDecl.Names {
			entryST, _ := b.tables.Globals.Contains(singleVarDecl.String())
			varEntry, _ := entryST.(*st.VarEntry)
			b.addVarDecl(varEntry)
		}
	}
	b.addStatement("")

	for _, funcDecl := range b.program.Functions {
		entry, _ := b.tables.Globals.Contains(funcDecl.Name.String())
		funcEntry, _ := entry.(*st.FuncEntry)
		b.addFuncEntry(funcEntry, funcDecl)
	}

	if !llvmStackMode {
		b.Mem2RegPass()
		b.OutOfSSAPass()
		b.LinearScanPass()
	}

	b.addStatement("")

	b.addStatement("declare ptr @malloc(i32)")
	b.addStatement("declare i32 @scanf(ptr, ...)")
	b.addStatement("declare i32 @printf(ptr, ...)")
	b.addStatement("declare void @free(ptr)")

	b.SaveBlocks(filename + ".ll")
}
