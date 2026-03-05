package lexer

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"golite/context"
)

type Lexer interface {
	GetTokenStream() *antlr.CommonTokenStream
	GetErrors() []*context.CompilerError
	PrintTokens()
}

type lexerWrapper struct {
	*antlr.DefaultErrorListener
	antlrLexer *GoliteLexer
	stream     *antlr.CommonTokenStream
	errors     []*context.CompilerError
}

func NewLexer(inputSourcePath string) Lexer {
	input, err := antlr.NewFileStream(inputSourcePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	lexer := &lexerWrapper{antlr.NewDefaultErrorListener(), nil, nil, nil}
	antlrLexer := NewGoliteLexer(input)
	antlrLexer.RemoveErrorListeners()
	antlrLexer.AddErrorListener(lexer)
	tokenStream := antlr.NewCommonTokenStream(antlrLexer, 0)
	lexer.antlrLexer = antlrLexer
	lexer.stream = tokenStream

	return lexer
}

func (lexer *lexerWrapper) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	lexer.errors = append(lexer.errors, &context.CompilerError{
		Line:  line,
		Col:   column,
		Msg:   msg,
		Phase: context.LEXER,
	})
}

func (lexer *lexerWrapper) GetTokenStream() *antlr.CommonTokenStream {
	return lexer.stream
}

func (lexer *lexerWrapper) GetErrors() []*context.CompilerError {
	return lexer.errors
}

func (lexer *lexerWrapper) PrintTokens() {
	tokens := lexer.GetTokenStream()
	tokens.Fill()
	var allTokens []antlr.Token = tokens.GetAllTokens()

	fmt.Println("Line | Col | TokenType(TokenText)")
	fmt.Println("---------------------------------")

	for _, t := range allTokens {
		tokenType := t.GetTokenType()
		var tokenName string

		switch {
		case tokenType == antlr.TokenEOF:
			tokenName = "EOF"
		case tokenType > 0 && tokenType < len(lexer.antlrLexer.SymbolicNames):
			tokenName = lexer.antlrLexer.SymbolicNames[tokenType]
		default:
			tokenName = "UNKNOWN"
		}

		fmt.Printf(
			"%4d | %3d | %-12s (%q)\n",
			t.GetLine(),
			t.GetColumn(),
			tokenName,
			t.GetText(),
		)
	}
}
