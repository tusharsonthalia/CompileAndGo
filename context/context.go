package context

import "fmt"

type CompilerPhase int

const (
	LEXER CompilerPhase = iota
	PARSER
	SEMANTIC
)

type CompilerError struct {
	Line  int
	Col   int
	Msg   string
	Phase CompilerPhase
}

func (err *CompilerError) String() string {
	switch err.Phase {
	case LEXER:
		return fmt.Sprintf("lexer error(%d:%d): %s", err.Line, err.Col, err.Msg)
	case PARSER:
		return fmt.Sprintf("syntax error(%d:%d): %s", err.Line, err.Col, err.Msg)
	case SEMANTIC:
		return fmt.Sprintf("semantic error(%d:%d): %s", err.Line, err.Col, err.Msg)
	}

	panic("Invalid Phase found!")
}

func HasErrors(errs []*CompilerError) bool {
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
		return true
	}
	return false
}
