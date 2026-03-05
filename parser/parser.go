package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"golite/ast"
	"golite/context"
	"golite/lexer"
	"golite/token"
	"golite/types"
	"strconv"
)

type Parser interface {
	Parse() *ast.Program
	PrintAST(*ast.Program)
	GetErrors() []*context.CompilerError
}

type parserWrapper struct {
	*antlr.DefaultErrorListener
	*BaseGoliteParserListener
	antlrParser *GoliteParser
	lexer       lexer.Lexer
	errors      []*context.CompilerError
	nodes       map[string]interface{}
}

func NewParser(lexer lexer.Lexer) Parser {
	parser := &parserWrapper{antlr.NewDefaultErrorListener(),
		&BaseGoliteParserListener{}, nil, nil, nil,
		make(map[string]interface{}),
	}
	antlrParser := NewGoliteParser(lexer.GetTokenStream())
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(parser)
	parser.antlrParser = antlrParser
	parser.lexer = lexer

	return parser
}

func (parser *parserWrapper) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	parser.errors = append(parser.errors, &context.CompilerError{
		Line:  line,
		Col:   column,
		Msg:   msg,
		Phase: context.PARSER,
	})
}

func (parser *parserWrapper) GetErrors() []*context.CompilerError {
	return parser.errors
}

func (parser *parserWrapper) Parse() *ast.Program {
	ctx := parser.antlrParser.Program()
	if context.HasErrors(parser.lexer.GetErrors()) || context.HasErrors(parser.GetErrors()) {
		return nil
	}

	antlr.ParseTreeWalkerDefault.Walk(parser, ctx)
	_, _, key := GetTokenInfo(ctx)

	return parser.nodes[key].(*ast.Program)
}

func (parser *parserWrapper) PrintAST(program *ast.Program) {
	fmt.Println(program)
}

func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string) {
	ctx_start := ctx.GetStart()
	line := ctx_start.GetLine()
	column := ctx_start.GetColumn()
	key := fmt.Sprintf("%d,%d", line, column)

	return line, column, key
}

/*
==============================
Implementation of Listeners
==============================
*/

func (parser *parserWrapper) ExitProgram(c *ProgramContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	allTypes := c.Types().AllTypeDeclaration()
	types := make([]*ast.TypeDecl, len(allTypes))
	for i, type_ := range allTypes {
		_, _, typeKey := GetTokenInfo(type_)
		types[i] = parser.nodes[typeKey].(*ast.TypeDecl)
	}

	allGlobals := c.Declarations().AllDeclaration()
	globals := make([]*ast.VarDecl, len(allGlobals))
	for i, decl := range allGlobals {
		_, _, declKey := GetTokenInfo(decl)
		globals[i] = parser.nodes[declKey].(*ast.VarDecl)
	}

	allFunctions := c.Functions().AllFunction()
	functions := make([]*ast.FuncDecl, len(allFunctions))
	for i, function := range allFunctions {
		_, _, functionKey := GetTokenInfo(function)
		functions[i] = parser.nodes[functionKey].(*ast.FuncDecl)
	}

	parser.nodes[key] = ast.NewProgram(
		types,
		globals,
		functions,
		tok,
	)
}

func (parser *parserWrapper) ExitTypeDeclaration(c *TypeDeclarationContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	nameLine := c.IDENTIFIER().GetSymbol().GetLine()
	nameCol := c.IDENTIFIER().GetSymbol().GetColumn()
	nameTok := token.NewToken(nameLine, nameCol)
	name := ast.NewVariable(c.IDENTIFIER().GetText(), types.StringToType("struct", c.IDENTIFIER().GetText()), nameTok)

	allFields := c.Fields().AllDecl()
	fields := make([]*ast.Field, len(allFields))
	for i, field := range allFields {
		_, _, fieldKey := GetTokenInfo(field)
		fields[i] = parser.nodes[fieldKey].(*ast.Field)
	}

	parser.nodes[key] = ast.NewTypeDecl(
		name,
		fields,
		tok,
	)
}

func (parser *parserWrapper) ExitDecl(c *DeclContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	nameLine := c.IDENTIFIER().GetSymbol().GetLine()
	nameCol := c.IDENTIFIER().GetSymbol().GetColumn()
	nameTok := token.NewToken(nameLine, nameCol)
	type_ := types.StringToType(c.Type_().GetText(), "")
	name := ast.NewVariable(c.IDENTIFIER().GetText(), type_, nameTok)

	parser.nodes[key] = ast.NewField(
		name,
		type_,
		tok,
	)
}

func (parser *parserWrapper) ExitDeclaration(c *DeclarationContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	allNames := c.Ids().AllIDENTIFIER()
	type_ := types.StringToType(c.Type_().GetText(), "")
	names := make([]ast.Expression, len(allNames))
	for i, name := range allNames {
		nameLine := name.GetSymbol().GetLine()
		nameCol := name.GetSymbol().GetColumn()
		nameTok := token.NewToken(nameLine, nameCol)
		names[i] = ast.NewVariable(name.GetText(), type_, nameTok)
	}

	parser.nodes[key] = ast.NewVarDecl(
		names,
		type_,
		tok,
	)
}

func (parser *parserWrapper) ExitFunction(c *FunctionContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	funcVar := c.IDENTIFIER()
	funcVarLine := funcVar.GetSymbol().GetLine()
	funcVarCol := funcVar.GetSymbol().GetColumn()
	funcVarTok := token.NewToken(funcVarLine, funcVarCol)

	allParameters := c.Parameters().AllDecl()
	parameters := make([]*ast.Field, len(allParameters))
	for i, param := range allParameters {
		_, _, paramKey := GetTokenInfo(param)
		parameters[i] = parser.nodes[paramKey].(*ast.Field)
	}

	allDeclarations := c.Declarations().AllDeclaration()
	declarations := make([]*ast.VarDecl, len(allDeclarations))
	for i, decl := range allDeclarations {
		_, _, declKey := GetTokenInfo(decl)
		declarations[i] = parser.nodes[declKey].(*ast.VarDecl)
	}

	allStatements := c.Statements().AllStatement()
	statements := make([]ast.Statement, len(allStatements))
	for i, statement := range allStatements {
		_, _, statementKey := GetTokenInfo(statement)
		statements[i] = parser.nodes[statementKey].(ast.Statement)
	}

	var returnType types.Type
	if c.ReturnType() != nil {
		returnType = types.StringToType(c.ReturnType().Type_().GetText(), "")
	} else {
		returnType = types.VoidTySig
	}

	type_ := types.StringToType("function", funcVar.GetText())
	parser.nodes[key] = ast.NewFuncDecl(
		ast.NewVariable(funcVar.GetText(), type_, funcVarTok),
		parameters,
		returnType,
		declarations,
		statements,
		tok,
	)
}

/* ---------------- Statements ---------------- */

func (parser *parserWrapper) ExitBlock(c *BlockContext) {
	_, _, key := GetTokenInfo(c)

	allStatements := c.Statements().AllStatement()
	statements := make([]ast.Statement, len(allStatements))

	for i, statement := range allStatements {
		_, _, statementKey := GetTokenInfo(statement)
		statements[i] = parser.nodes[statementKey].(ast.Statement)
	}

	parser.nodes[key] = statements
}

func (parser *parserWrapper) ExitDelete(c *DeleteContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	_, _, exprKey := GetTokenInfo(c.Expression())
	expr := parser.nodes[exprKey].(ast.Expression)

	parser.nodes[key] = ast.NewDelete(
		expr,
		tok,
	)
}

func (parser *parserWrapper) ExitRead(c *ReadContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	_, _, lValueKey := GetTokenInfo(c.Lvalue())
	lValue := parser.nodes[lValueKey].(*ast.LValue)

	parser.nodes[key] = ast.NewRead(
		lValue,
		tok,
	)
}

func (parser *parserWrapper) ExitAssignment(c *AssignmentContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	_, _, lValueKey := GetTokenInfo(c.Lvalue())
	lValue := parser.nodes[lValueKey].(*ast.LValue)

	_, _, exprKey := GetTokenInfo(c.Expression())
	expr := parser.nodes[exprKey].(ast.Expression)

	parser.nodes[key] = ast.NewAssignment(
		lValue,
		expr,
		tok,
	)
}

func (parser *parserWrapper) ExitPrint(c *PrintContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	allExpressions := c.AllExpression()
	exprs := make([]ast.Expression, len(allExpressions))

	for i, expr := range allExpressions {
		_, _, exprKey := GetTokenInfo(expr)
		exprs[i] = parser.nodes[exprKey].(ast.Expression)
	}

	parser.nodes[key] = ast.NewPrint(
		c.STRING().GetText(),
		exprs,
		tok,
	)
}

func (parser *parserWrapper) ExitConditional(c *ConditionalContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	_, _, exprKey := GetTokenInfo(c.Expression())

	allBlocks := c.AllBlock()

	var ifBlock, elseBlock []ast.Statement

	_, _, ifBlockKey := GetTokenInfo(allBlocks[0])
	ifBlock = parser.nodes[ifBlockKey].([]ast.Statement)

	if len(allBlocks) > 1 {
		_, _, elseBlockKey := GetTokenInfo(allBlocks[1])
		elseBlock = parser.nodes[elseBlockKey].([]ast.Statement)
	}

	parser.nodes[key] = ast.NewConditional(
		parser.nodes[exprKey].(ast.Expression),
		ifBlock,
		elseBlock,
		tok,
	)
}

func (parser *parserWrapper) ExitLoop(c *LoopContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	_, _, exprKey := GetTokenInfo(c.Expression())
	_, _, blockKey := GetTokenInfo(c.Block())

	parser.nodes[key] = ast.NewLoop(
		parser.nodes[exprKey].(ast.Expression),
		parser.nodes[blockKey].([]ast.Statement),
		tok,
	)
}

func (parser *parserWrapper) ExitReturn(c *ReturnContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	target := ast.Expression(nil)

	if c.Expression() != nil {
		_, _, exprKey := GetTokenInfo(c.Expression())
		target = parser.nodes[exprKey].(ast.Expression)
	}

	parser.nodes[key] = ast.NewReturn(
		target,
		tok,
	)
}

func (parser *parserWrapper) ExitInvocation(c *InvocationContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	type_ := types.StringToType("function", c.IDENTIFIER().GetText())
	varNode := ast.NewVariable(c.IDENTIFIER().GetText(), type_, tok)
	_, _, argKey := GetTokenInfo(c.Arguments())

	parser.nodes[key] = ast.NewInvocation(
		varNode,
		parser.nodes[argKey].([]ast.Expression),
		tok,
	)
}

func (parser *parserWrapper) ExitArguments(c *ArgumentsContext) {
	_, _, key := GetTokenInfo(c)

	exprs := c.AllExpression()
	expr := make([]ast.Expression, len(exprs))

	for i, id := range exprs {
		_, _, exprKey := GetTokenInfo(id)
		expr[i] = parser.nodes[exprKey].(ast.Expression)
	}

	parser.nodes[key] = expr
}

func (parser *parserWrapper) ExitLvalue(c *LvalueContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	ids := c.AllIDENTIFIER()
	values := make([]ast.Expression, len(ids))

	for i, id := range ids {
		varLine := id.GetSymbol().GetLine()
		varCol := id.GetSymbol().GetColumn()
		varTok := token.NewToken(varLine, varCol)
		values[i] = ast.NewVariable(id.GetText(), nil, varTok)
	}

	parser.nodes[key] = ast.NewLValue(
		values,
		tok,
	)
}

/* ---------------- Expressions ---------------- */

func (parser *parserWrapper) ExitExpression(c *ExpressionContext) {
	_, _, key := GetTokenInfo(c)

	boolTerms := c.AllBoolterm()
	_, _, unaryKey := GetTokenInfo(boolTerms[0])

	currentNode := parser.nodes[unaryKey].(ast.Expression)

	if len(boolTerms) == 1 {
		parser.nodes[key] = currentNode
		return
	}

	for i := 1; i < len(boolTerms); i++ {
		opNode := c.GetChild(2*i - 1).(antlr.TerminalNode)

		_, _, termKey := GetTokenInfo(boolTerms[i].(*BooltermContext))
		rvalue := parser.nodes[termKey].(ast.Expression)

		opTok := token.NewToken(
			opNode.GetSymbol().GetLine(),
			opNode.GetSymbol().GetColumn(),
		)

		currentNode = ast.NewBinOp(
			currentNode,
			ast.StrToOp(opNode.GetText()),
			rvalue,
			opTok,
		)
	}

	parser.nodes[key] = currentNode
}

func (parser *parserWrapper) ExitBoolterm(c *BooltermContext) {
	_, _, key := GetTokenInfo(c)

	equalTerms := c.AllEqualterm()
	_, _, unaryKey := GetTokenInfo(equalTerms[0])

	currentNode := parser.nodes[unaryKey].(ast.Expression)

	if len(equalTerms) == 1 {
		parser.nodes[key] = currentNode
		return
	}

	for i := 1; i < len(equalTerms); i++ {
		opNode := c.GetChild(2*i - 1).(antlr.TerminalNode)

		_, _, termKey := GetTokenInfo(equalTerms[i].(*EqualtermContext))
		rvalue := parser.nodes[termKey].(ast.Expression)

		opTok := token.NewToken(
			opNode.GetSymbol().GetLine(),
			opNode.GetSymbol().GetColumn(),
		)

		currentNode = ast.NewBinOp(
			currentNode,
			ast.StrToOp(opNode.GetText()),
			rvalue,
			opTok,
		)
	}

	parser.nodes[key] = currentNode
}

func (parser *parserWrapper) ExitEqualterm(c *EqualtermContext) {
	_, _, key := GetTokenInfo(c)

	relationTerms := c.AllRelationterm()
	_, _, unaryKey := GetTokenInfo(relationTerms[0])

	currentNode := parser.nodes[unaryKey].(ast.Expression)

	if len(relationTerms) == 1 {
		parser.nodes[key] = currentNode
		return
	}

	for i := 1; i < len(relationTerms); i++ {
		opNode := c.GetChild(2*i - 1).(antlr.TerminalNode)

		_, _, termKey := GetTokenInfo(relationTerms[i].(*RelationtermContext))
		rvalue := parser.nodes[termKey].(ast.Expression)

		opTok := token.NewToken(
			opNode.GetSymbol().GetLine(),
			opNode.GetSymbol().GetColumn(),
		)

		currentNode = ast.NewBinOp(
			currentNode,
			ast.StrToOp(opNode.GetText()),
			rvalue,
			opTok,
		)
	}

	parser.nodes[key] = currentNode
}

func (parser *parserWrapper) ExitRelationterm(c *RelationtermContext) {
	_, _, key := GetTokenInfo(c)

	simpleTerms := c.AllSimpleterm()
	_, _, unaryKey := GetTokenInfo(simpleTerms[0])

	currentNode := parser.nodes[unaryKey].(ast.Expression)

	if len(simpleTerms) == 1 {
		parser.nodes[key] = currentNode
		return
	}

	for i := 1; i < len(simpleTerms); i++ {
		opNode := c.GetChild(2*i - 1).(antlr.TerminalNode)

		_, _, termKey := GetTokenInfo(simpleTerms[i].(*SimpletermContext))
		rvalue := parser.nodes[termKey].(ast.Expression)

		opTok := token.NewToken(
			opNode.GetSymbol().GetLine(),
			opNode.GetSymbol().GetColumn(),
		)

		currentNode = ast.NewBinOp(
			currentNode,
			ast.StrToOp(opNode.GetText()),
			rvalue,
			opTok,
		)
	}

	parser.nodes[key] = currentNode
}

func (parser *parserWrapper) ExitSimpleterm(c *SimpletermContext) {
	_, _, key := GetTokenInfo(c)

	terms := c.AllTerm()
	_, _, unaryKey := GetTokenInfo(terms[0])

	currentNode := parser.nodes[unaryKey].(ast.Expression)

	if len(terms) == 1 {
		parser.nodes[key] = currentNode
		return
	}

	for i := 1; i < len(terms); i++ {
		opNode := c.GetChild(2*i - 1).(antlr.TerminalNode)

		_, _, termKey := GetTokenInfo(terms[i].(*TermContext))
		rvalue := parser.nodes[termKey].(ast.Expression)

		opTok := token.NewToken(
			opNode.GetSymbol().GetLine(),
			opNode.GetSymbol().GetColumn(),
		)

		currentNode = ast.NewBinOp(
			currentNode,
			ast.StrToOp(opNode.GetText()),
			rvalue,
			opTok,
		)
	}

	parser.nodes[key] = currentNode
}

func (parser *parserWrapper) ExitTerm(c *TermContext) {
	_, _, key := GetTokenInfo(c)

	unaryTerms := c.AllUnaryterm()
	_, _, unaryKey := GetTokenInfo(unaryTerms[0])

	currentNode := parser.nodes[unaryKey].(ast.Expression)

	if len(unaryTerms) == 1 {
		parser.nodes[key] = currentNode
		return
	}

	for i := 1; i < len(unaryTerms); i++ {
		opNode := c.GetChild(2*i - 1).(antlr.TerminalNode)

		_, _, termKey := GetTokenInfo(unaryTerms[i].(*UnarytermContext))
		rvalue := parser.nodes[termKey].(ast.Expression)

		opTok := token.NewToken(
			opNode.GetSymbol().GetLine(),
			opNode.GetSymbol().GetColumn(),
		)

		currentNode = ast.NewBinOp(
			currentNode,
			ast.StrToOp(opNode.GetText()),
			rvalue,
			opTok,
		)
	}

	parser.nodes[key] = currentNode
}

func (parser *parserWrapper) ExitUnaryterm(c *UnarytermContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	_, _, selectorKey := GetTokenInfo(c.Selectorterm())

	if c.EXCLAMATION() != nil {
		parser.nodes[key] = ast.NewUnaryOp(ast.EXCLAMATION, parser.nodes[selectorKey].(ast.Expression), tok)
	} else if c.MINUS() != nil {
		parser.nodes[key] = ast.NewUnaryOp(ast.MINUS, parser.nodes[selectorKey].(ast.Expression), tok)
	} else {
		parser.nodes[key] = parser.nodes[selectorKey].(ast.Expression)
	}
}

func (parser *parserWrapper) ExitSelectorterm(c *SelectortermContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)

	_, _, factorKey := GetTokenInfo(c.Factor())
	var accessors []string

	for _, id := range c.AllIDENTIFIER() {
		accessors = append(accessors, id.GetText())
	}

	parser.nodes[key] = ast.NewSelector(parser.nodes[factorKey].(ast.Expression), accessors, nil, tok)
}

func (parser *parserWrapper) ExitFactor(c *FactorContext) {
	line, col, key := GetTokenInfo(c)
	tok := token.NewToken(line, col)
	var node interface{}

	if numberFactor := c.NUMBER(); numberFactor != nil {
		intValue, _ := strconv.ParseInt(numberFactor.GetText(), 10, 64)
		node = ast.NewIntLit(int64(intValue), tok)
	} else if trueFactor := c.TRUE(); trueFactor != nil {
		node = ast.NewBoolLit(true, tok)
	} else if falseFactor := c.FALSE(); falseFactor != nil {
		node = ast.NewBoolLit(false, tok)
	} else if nilFactor := c.NIL(); nilFactor != nil {
		node = ast.NewNilLit(nilFactor.GetText(), tok)
	} else if newFactor := c.NEW(); newFactor != nil && c.IDENTIFIER() != nil {
		varLine := c.IDENTIFIER().GetSymbol().GetLine()
		varCol := c.IDENTIFIER().GetSymbol().GetColumn()
		varTok := token.NewToken(varLine, varCol)
		type_ := types.StringToType("struct", c.IDENTIFIER().GetText())
		varDef := ast.NewVariable(c.IDENTIFIER().GetText(), type_, varTok)
		node = ast.NewAllocate(varDef, tok)
	} else if idFactor := c.IDENTIFIER(); idFactor != nil {
		if argFactor := c.Arguments(); argFactor != nil {
			_, _, argKey := GetTokenInfo(c.Arguments())
			type_ := types.StringToType("function", idFactor.GetText())
			varDef := ast.NewVariable(idFactor.GetText(), type_, tok)
			node = ast.NewCall(varDef, parser.nodes[argKey].([]ast.Expression), tok)
		} else {
			node = ast.NewVariable(idFactor.GetText(), nil, tok)
		}
	} else if expressionFactor := c.Expression(); expressionFactor != nil {
		_, _, exprKey := GetTokenInfo(c.Expression())
		node = parser.nodes[exprKey].(ast.Expression)
	}

	parser.nodes[key] = node
}
