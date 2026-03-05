// Code generated from GoliteParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GoliteParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoliteParser struct {
	*antlr.BaseParser
}

var GoliteParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goliteparserParserInit() {
	staticData := &GoliteParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'type'", "'struct'", "'func'", "'printf'", "'delete'", "'scan'",
		"'for'", "'if'", "'else'", "'return'", "'var'", "'int'", "'bool'", "'true'",
		"'false'", "'nil'", "'new'", "'||'", "'&&'", "'=='", "'!='", "'>='",
		"'<='", "'{'", "'}'", "'('", "')'", "';'", "','", "'.'", "'!'", "'/'",
		"'*'", "'+'", "'-'", "'>'", "'<'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "TYPE", "STRUCT", "FUNCTION", "PRINT", "DELETE", "SCAN",
		"FOR", "IF", "ELSE", "RETURN", "VAR", "INT", "BOOL", "TRUE", "FALSE",
		"NIL", "NEW", "OR", "AND", "DOUBLEEQ", "NEQ", "GEQ", "LEQ", "LBRACE",
		"RBRACE", "LPAREN", "RPAREN", "SEMICOLON", "COMMA", "DOT", "EXCLAMATION",
		"FSLASH", "ASTERISK", "PLUS", "MINUS", "GT", "LT", "EQUALS", "IDENTIFIER",
		"NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"program", "types", "typeDeclaration", "fields", "decl", "type", "declarations",
		"declaration", "ids", "functions", "function", "parameters", "returnType",
		"statements", "statement", "block", "delete", "read", "assignment",
		"print", "conditional", "loop", "return", "invocation", "arguments",
		"lvalue", "expression", "boolterm", "equalterm", "relationterm", "simpleterm",
		"term", "unaryterm", "selectorterm", "factor",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 329, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 5, 1, 77, 8, 1, 10, 1, 12, 1, 80, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 95, 8, 3, 10,
		3, 12, 3, 98, 9, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 107,
		8, 5, 1, 6, 5, 6, 110, 8, 6, 10, 6, 12, 6, 113, 9, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 123, 8, 8, 10, 8, 12, 8, 126, 9, 8,
		1, 9, 5, 9, 129, 8, 9, 10, 9, 12, 9, 132, 9, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 138, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 149, 8, 11, 10, 11, 12, 11, 152, 9, 11, 3, 11, 154,
		8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 5, 13, 161, 8, 13, 10, 13, 12,
		13, 164, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 174, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 198, 8, 19, 10, 19, 12, 19, 201, 9,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		3, 20, 213, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 3, 22, 223, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 5, 24, 235, 8, 24, 10, 24, 12, 24, 238, 9, 24, 3,
		24, 240, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 5, 25, 247, 8, 25, 10,
		25, 12, 25, 250, 9, 25, 1, 26, 1, 26, 1, 26, 5, 26, 255, 8, 26, 10, 26,
		12, 26, 258, 9, 26, 1, 27, 1, 27, 1, 27, 5, 27, 263, 8, 27, 10, 27, 12,
		27, 266, 9, 27, 1, 28, 1, 28, 1, 28, 5, 28, 271, 8, 28, 10, 28, 12, 28,
		274, 9, 28, 1, 29, 1, 29, 1, 29, 5, 29, 279, 8, 29, 10, 29, 12, 29, 282,
		9, 29, 1, 30, 1, 30, 1, 30, 5, 30, 287, 8, 30, 10, 30, 12, 30, 290, 9,
		30, 1, 31, 1, 31, 1, 31, 5, 31, 295, 8, 31, 10, 31, 12, 31, 298, 9, 31,
		1, 32, 3, 32, 301, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 5, 33, 308,
		8, 33, 10, 33, 12, 33, 311, 9, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 3, 34, 319, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34,
		327, 8, 34, 1, 34, 0, 0, 35, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 0, 5, 1, 0, 21, 22, 2, 0, 23, 24, 37, 38, 1, 0, 35,
		36, 1, 0, 33, 34, 2, 0, 32, 32, 36, 36, 332, 0, 70, 1, 0, 0, 0, 2, 78,
		1, 0, 0, 0, 4, 81, 1, 0, 0, 0, 6, 89, 1, 0, 0, 0, 8, 99, 1, 0, 0, 0, 10,
		106, 1, 0, 0, 0, 12, 111, 1, 0, 0, 0, 14, 114, 1, 0, 0, 0, 16, 119, 1,
		0, 0, 0, 18, 130, 1, 0, 0, 0, 20, 133, 1, 0, 0, 0, 22, 144, 1, 0, 0, 0,
		24, 157, 1, 0, 0, 0, 26, 162, 1, 0, 0, 0, 28, 173, 1, 0, 0, 0, 30, 175,
		1, 0, 0, 0, 32, 179, 1, 0, 0, 0, 34, 183, 1, 0, 0, 0, 36, 187, 1, 0, 0,
		0, 38, 192, 1, 0, 0, 0, 40, 205, 1, 0, 0, 0, 42, 214, 1, 0, 0, 0, 44, 220,
		1, 0, 0, 0, 46, 226, 1, 0, 0, 0, 48, 230, 1, 0, 0, 0, 50, 243, 1, 0, 0,
		0, 52, 251, 1, 0, 0, 0, 54, 259, 1, 0, 0, 0, 56, 267, 1, 0, 0, 0, 58, 275,
		1, 0, 0, 0, 60, 283, 1, 0, 0, 0, 62, 291, 1, 0, 0, 0, 64, 300, 1, 0, 0,
		0, 66, 304, 1, 0, 0, 0, 68, 326, 1, 0, 0, 0, 70, 71, 3, 2, 1, 0, 71, 72,
		3, 12, 6, 0, 72, 73, 3, 18, 9, 0, 73, 74, 5, 0, 0, 1, 74, 1, 1, 0, 0, 0,
		75, 77, 3, 4, 2, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1,
		0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 3, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81,
		82, 5, 2, 0, 0, 82, 83, 5, 40, 0, 0, 83, 84, 5, 3, 0, 0, 84, 85, 5, 25,
		0, 0, 85, 86, 3, 6, 3, 0, 86, 87, 5, 26, 0, 0, 87, 88, 5, 29, 0, 0, 88,
		5, 1, 0, 0, 0, 89, 90, 3, 8, 4, 0, 90, 96, 5, 29, 0, 0, 91, 92, 3, 8, 4,
		0, 92, 93, 5, 29, 0, 0, 93, 95, 1, 0, 0, 0, 94, 91, 1, 0, 0, 0, 95, 98,
		1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 7, 1, 0, 0, 0,
		98, 96, 1, 0, 0, 0, 99, 100, 5, 40, 0, 0, 100, 101, 3, 10, 5, 0, 101, 9,
		1, 0, 0, 0, 102, 107, 5, 13, 0, 0, 103, 107, 5, 14, 0, 0, 104, 105, 5,
		34, 0, 0, 105, 107, 5, 40, 0, 0, 106, 102, 1, 0, 0, 0, 106, 103, 1, 0,
		0, 0, 106, 104, 1, 0, 0, 0, 107, 11, 1, 0, 0, 0, 108, 110, 3, 14, 7, 0,
		109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111,
		112, 1, 0, 0, 0, 112, 13, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 5,
		12, 0, 0, 115, 116, 3, 16, 8, 0, 116, 117, 3, 10, 5, 0, 117, 118, 5, 29,
		0, 0, 118, 15, 1, 0, 0, 0, 119, 124, 5, 40, 0, 0, 120, 121, 5, 30, 0, 0,
		121, 123, 5, 40, 0, 0, 122, 120, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 17, 1, 0, 0, 0, 126, 124, 1,
		0, 0, 0, 127, 129, 3, 20, 10, 0, 128, 127, 1, 0, 0, 0, 129, 132, 1, 0,
		0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 19, 1, 0, 0, 0,
		132, 130, 1, 0, 0, 0, 133, 134, 5, 4, 0, 0, 134, 135, 5, 40, 0, 0, 135,
		137, 3, 22, 11, 0, 136, 138, 3, 24, 12, 0, 137, 136, 1, 0, 0, 0, 137, 138,
		1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 5, 25, 0, 0, 140, 141, 3, 12,
		6, 0, 141, 142, 3, 26, 13, 0, 142, 143, 5, 26, 0, 0, 143, 21, 1, 0, 0,
		0, 144, 153, 5, 27, 0, 0, 145, 150, 3, 8, 4, 0, 146, 147, 5, 30, 0, 0,
		147, 149, 3, 8, 4, 0, 148, 146, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150,
		148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150,
		1, 0, 0, 0, 153, 145, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 156, 5, 28, 0, 0, 156, 23, 1, 0, 0, 0, 157, 158, 3, 10, 5, 0,
		158, 25, 1, 0, 0, 0, 159, 161, 3, 28, 14, 0, 160, 159, 1, 0, 0, 0, 161,
		164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 27, 1,
		0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 174, 3, 36, 18, 0, 166, 174, 3, 38,
		19, 0, 167, 174, 3, 34, 17, 0, 168, 174, 3, 32, 16, 0, 169, 174, 3, 40,
		20, 0, 170, 174, 3, 42, 21, 0, 171, 174, 3, 44, 22, 0, 172, 174, 3, 46,
		23, 0, 173, 165, 1, 0, 0, 0, 173, 166, 1, 0, 0, 0, 173, 167, 1, 0, 0, 0,
		173, 168, 1, 0, 0, 0, 173, 169, 1, 0, 0, 0, 173, 170, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 29, 1, 0, 0, 0, 175, 176, 5,
		25, 0, 0, 176, 177, 3, 26, 13, 0, 177, 178, 5, 26, 0, 0, 178, 31, 1, 0,
		0, 0, 179, 180, 5, 6, 0, 0, 180, 181, 3, 52, 26, 0, 181, 182, 5, 29, 0,
		0, 182, 33, 1, 0, 0, 0, 183, 184, 5, 7, 0, 0, 184, 185, 3, 50, 25, 0, 185,
		186, 5, 29, 0, 0, 186, 35, 1, 0, 0, 0, 187, 188, 3, 50, 25, 0, 188, 189,
		5, 39, 0, 0, 189, 190, 3, 52, 26, 0, 190, 191, 5, 29, 0, 0, 191, 37, 1,
		0, 0, 0, 192, 193, 5, 5, 0, 0, 193, 194, 5, 27, 0, 0, 194, 199, 5, 42,
		0, 0, 195, 196, 5, 30, 0, 0, 196, 198, 3, 52, 26, 0, 197, 195, 1, 0, 0,
		0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200,
		202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203, 5, 28, 0, 0, 203, 204,
		5, 29, 0, 0, 204, 39, 1, 0, 0, 0, 205, 206, 5, 9, 0, 0, 206, 207, 5, 27,
		0, 0, 207, 208, 3, 52, 26, 0, 208, 209, 5, 28, 0, 0, 209, 212, 3, 30, 15,
		0, 210, 211, 5, 10, 0, 0, 211, 213, 3, 30, 15, 0, 212, 210, 1, 0, 0, 0,
		212, 213, 1, 0, 0, 0, 213, 41, 1, 0, 0, 0, 214, 215, 5, 8, 0, 0, 215, 216,
		5, 27, 0, 0, 216, 217, 3, 52, 26, 0, 217, 218, 5, 28, 0, 0, 218, 219, 3,
		30, 15, 0, 219, 43, 1, 0, 0, 0, 220, 222, 5, 11, 0, 0, 221, 223, 3, 52,
		26, 0, 222, 221, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0,
		224, 225, 5, 29, 0, 0, 225, 45, 1, 0, 0, 0, 226, 227, 5, 40, 0, 0, 227,
		228, 3, 48, 24, 0, 228, 229, 5, 29, 0, 0, 229, 47, 1, 0, 0, 0, 230, 239,
		5, 27, 0, 0, 231, 236, 3, 52, 26, 0, 232, 233, 5, 30, 0, 0, 233, 235, 3,
		52, 26, 0, 234, 232, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0,
		0, 0, 236, 237, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0,
		239, 231, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241,
		242, 5, 28, 0, 0, 242, 49, 1, 0, 0, 0, 243, 248, 5, 40, 0, 0, 244, 245,
		5, 31, 0, 0, 245, 247, 5, 40, 0, 0, 246, 244, 1, 0, 0, 0, 247, 250, 1,
		0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 51, 1, 0, 0,
		0, 250, 248, 1, 0, 0, 0, 251, 256, 3, 54, 27, 0, 252, 253, 5, 19, 0, 0,
		253, 255, 3, 54, 27, 0, 254, 252, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256,
		254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 53, 1, 0, 0, 0, 258, 256, 1,
		0, 0, 0, 259, 264, 3, 56, 28, 0, 260, 261, 5, 20, 0, 0, 261, 263, 3, 56,
		28, 0, 262, 260, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0,
		264, 265, 1, 0, 0, 0, 265, 55, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 272,
		3, 58, 29, 0, 268, 269, 7, 0, 0, 0, 269, 271, 3, 58, 29, 0, 270, 268, 1,
		0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0,
		0, 273, 57, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 280, 3, 60, 30, 0, 276,
		277, 7, 1, 0, 0, 277, 279, 3, 60, 30, 0, 278, 276, 1, 0, 0, 0, 279, 282,
		1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 59, 1, 0,
		0, 0, 282, 280, 1, 0, 0, 0, 283, 288, 3, 62, 31, 0, 284, 285, 7, 2, 0,
		0, 285, 287, 3, 62, 31, 0, 286, 284, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0,
		288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 61, 1, 0, 0, 0, 290, 288,
		1, 0, 0, 0, 291, 296, 3, 64, 32, 0, 292, 293, 7, 3, 0, 0, 293, 295, 3,
		64, 32, 0, 294, 292, 1, 0, 0, 0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0,
		0, 0, 296, 297, 1, 0, 0, 0, 297, 63, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0,
		299, 301, 7, 4, 0, 0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301,
		302, 1, 0, 0, 0, 302, 303, 3, 66, 33, 0, 303, 65, 1, 0, 0, 0, 304, 309,
		3, 68, 34, 0, 305, 306, 5, 31, 0, 0, 306, 308, 5, 40, 0, 0, 307, 305, 1,
		0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0,
		0, 310, 67, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313, 5, 27, 0, 0, 313,
		314, 3, 52, 26, 0, 314, 315, 5, 28, 0, 0, 315, 327, 1, 0, 0, 0, 316, 318,
		5, 40, 0, 0, 317, 319, 3, 48, 24, 0, 318, 317, 1, 0, 0, 0, 318, 319, 1,
		0, 0, 0, 319, 327, 1, 0, 0, 0, 320, 327, 5, 41, 0, 0, 321, 322, 5, 18,
		0, 0, 322, 327, 5, 40, 0, 0, 323, 327, 5, 15, 0, 0, 324, 327, 5, 16, 0,
		0, 325, 327, 5, 17, 0, 0, 326, 312, 1, 0, 0, 0, 326, 316, 1, 0, 0, 0, 326,
		320, 1, 0, 0, 0, 326, 321, 1, 0, 0, 0, 326, 323, 1, 0, 0, 0, 326, 324,
		1, 0, 0, 0, 326, 325, 1, 0, 0, 0, 327, 69, 1, 0, 0, 0, 27, 78, 96, 106,
		111, 124, 130, 137, 150, 153, 162, 173, 199, 212, 222, 236, 239, 248, 256,
		264, 272, 280, 288, 296, 300, 309, 318, 326,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoliteParserInit initializes any static state used to implement GoliteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoliteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteParserInit() {
	staticData := &GoliteParserParserStaticData
	staticData.once.Do(goliteparserParserInit)
}

// NewGoliteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoliteParser(input antlr.TokenStream) *GoliteParser {
	GoliteParserInit()
	this := new(GoliteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoliteParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GoliteParser.g4"

	return this
}

// GoliteParser tokens.
const (
	GoliteParserEOF         = antlr.TokenEOF
	GoliteParserCOMMENT     = 1
	GoliteParserTYPE        = 2
	GoliteParserSTRUCT      = 3
	GoliteParserFUNCTION    = 4
	GoliteParserPRINT       = 5
	GoliteParserDELETE      = 6
	GoliteParserSCAN        = 7
	GoliteParserFOR         = 8
	GoliteParserIF          = 9
	GoliteParserELSE        = 10
	GoliteParserRETURN      = 11
	GoliteParserVAR         = 12
	GoliteParserINT         = 13
	GoliteParserBOOL        = 14
	GoliteParserTRUE        = 15
	GoliteParserFALSE       = 16
	GoliteParserNIL         = 17
	GoliteParserNEW         = 18
	GoliteParserOR          = 19
	GoliteParserAND         = 20
	GoliteParserDOUBLEEQ    = 21
	GoliteParserNEQ         = 22
	GoliteParserGEQ         = 23
	GoliteParserLEQ         = 24
	GoliteParserLBRACE      = 25
	GoliteParserRBRACE      = 26
	GoliteParserLPAREN      = 27
	GoliteParserRPAREN      = 28
	GoliteParserSEMICOLON   = 29
	GoliteParserCOMMA       = 30
	GoliteParserDOT         = 31
	GoliteParserEXCLAMATION = 32
	GoliteParserFSLASH      = 33
	GoliteParserASTERISK    = 34
	GoliteParserPLUS        = 35
	GoliteParserMINUS       = 36
	GoliteParserGT          = 37
	GoliteParserLT          = 38
	GoliteParserEQUALS      = 39
	GoliteParserIDENTIFIER  = 40
	GoliteParserNUMBER      = 41
	GoliteParserSTRING      = 42
	GoliteParserWS          = 43
)

// GoliteParser rules.
const (
	GoliteParserRULE_program         = 0
	GoliteParserRULE_types           = 1
	GoliteParserRULE_typeDeclaration = 2
	GoliteParserRULE_fields          = 3
	GoliteParserRULE_decl            = 4
	GoliteParserRULE_type            = 5
	GoliteParserRULE_declarations    = 6
	GoliteParserRULE_declaration     = 7
	GoliteParserRULE_ids             = 8
	GoliteParserRULE_functions       = 9
	GoliteParserRULE_function        = 10
	GoliteParserRULE_parameters      = 11
	GoliteParserRULE_returnType      = 12
	GoliteParserRULE_statements      = 13
	GoliteParserRULE_statement       = 14
	GoliteParserRULE_block           = 15
	GoliteParserRULE_delete          = 16
	GoliteParserRULE_read            = 17
	GoliteParserRULE_assignment      = 18
	GoliteParserRULE_print           = 19
	GoliteParserRULE_conditional     = 20
	GoliteParserRULE_loop            = 21
	GoliteParserRULE_return          = 22
	GoliteParserRULE_invocation      = 23
	GoliteParserRULE_arguments       = 24
	GoliteParserRULE_lvalue          = 25
	GoliteParserRULE_expression      = 26
	GoliteParserRULE_boolterm        = 27
	GoliteParserRULE_equalterm       = 28
	GoliteParserRULE_relationterm    = 29
	GoliteParserRULE_simpleterm      = 30
	GoliteParserRULE_term            = 31
	GoliteParserRULE_unaryterm       = 32
	GoliteParserRULE_selectorterm    = 33
	GoliteParserRULE_factor          = 34
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Types() ITypesContext
	Declarations() IDeclarationsContext
	Functions() IFunctionsContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ProgramContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *ProgramContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoliteParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GoliteParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoliteParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Types()
	}
	{
		p.SetState(71)
		p.Declarations()
	}
	{
		p.SetState(72)
		p.Functions()
	}
	{
		p.SetState(73)
		p.Match(GoliteParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeDeclaration() []ITypeDeclarationContext
	TypeDeclaration(i int) ITypeDeclarationContext

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) AllTypeDeclaration() []ITypeDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclarationContext); ok {
			tst[i] = t.(ITypeDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *TypesContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *GoliteParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoliteParserRULE_types)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserTYPE {
		{
			p.SetState(75)
			p.TypeDeclaration()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	Fields() IFieldsContext
	RBRACE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_typeDeclaration
	return p
}

func InitEmptyTypeDeclarationContext(p *TypeDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_typeDeclaration
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTYPE, 0)
}

func (s *TypeDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, 0)
}

func (s *TypeDeclarationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRUCT, 0)
}

func (s *TypeDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *TypeDeclarationContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *TypeDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *TypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (p *GoliteParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoliteParserRULE_typeDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(GoliteParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(GoliteParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(GoliteParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(GoliteParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Fields()
	}
	{
		p.SetState(86)
		p.Match(GoliteParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDecl() []IDeclContext
	Decl(i int) IDeclContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
	return p
}

func InitEmptyFieldsContext(p *FieldsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserSEMICOLON)
}

func (s *FieldsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, i)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *GoliteParser) Fields() (localctx IFieldsContext) {
	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoliteParserRULE_fields)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Decl()
	}
	{
		p.SetState(90)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserIDENTIFIER {
		{
			p.SetState(91)
			p.Decl()
		}
		{
			p.SetState(92)
			p.Match(GoliteParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
	return p
}

func InitEmptyDeclContext(p *DeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, 0)
}

func (s *DeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *GoliteParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoliteParserRULE_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(GoliteParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	ASTERISK() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(GoliteParserINT, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GoliteParserBOOL, 0)
}

func (s *TypeContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, 0)
}

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *GoliteParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoliteParserRULE_type)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoliteParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Match(GoliteParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(GoliteParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserASTERISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(GoliteParserASTERISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(GoliteParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationsContext is an interface to support dynamic dispatch.
type IDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsDeclarationsContext differentiates from other interfaces.
	IsDeclarationsContext()
}

type DeclarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsContext() *DeclarationsContext {
	var p = new(DeclarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declarations
	return p
}

func InitEmptyDeclarationsContext(p *DeclarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declarations
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declarations

	return p
}

func (s *DeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationsContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclarations(s)
	}
}

func (s *DeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclarations(s)
	}
}

func (p *GoliteParser) Declarations() (localctx IDeclarationsContext) {
	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoliteParserRULE_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserVAR {
		{
			p.SetState(108)
			p.Declaration()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	Ids() IIdsContext
	Type_() ITypeContext
	SEMICOLON() antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GoliteParserVAR, 0)
}

func (s *DeclarationContext) Ids() IIdsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdsContext)
}

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *GoliteParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoliteParserRULE_declaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(GoliteParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Ids()
	}
	{
		p.SetState(116)
		p.Type_()
	}
	{
		p.SetState(117)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdsContext is an interface to support dynamic dispatch.
type IIdsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdsContext differentiates from other interfaces.
	IsIdsContext()
}

type IdsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdsContext() *IdsContext {
	var p = new(IdsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_ids
	return p
}

func InitEmptyIdsContext(p *IdsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_ids
}

func (*IdsContext) IsIdsContext() {}

func NewIdsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdsContext {
	var p = new(IdsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_ids

	return p
}

func (s *IdsContext) GetParser() antlr.Parser { return s.parser }

func (s *IdsContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserIDENTIFIER)
}

func (s *IdsContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, i)
}

func (s *IdsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *IdsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *IdsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterIds(s)
	}
}

func (s *IdsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitIds(s)
	}
}

func (p *GoliteParser) Ids() (localctx IIdsContext) {
	localctx = NewIdsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoliteParserRULE_ids)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GoliteParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(120)
			p.Match(GoliteParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(GoliteParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFunction() []IFunctionContext
	Function(i int) IFunctionContext

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_functions
	return p
}

func InitEmptyFunctionsContext(p *FunctionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_functions
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionsContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *GoliteParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoliteParserRULE_functions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFUNCTION {
		{
			p.SetState(127)
			p.Function()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Parameters() IParametersContext
	LBRACE() antlr.TerminalNode
	Declarations() IDeclarationsContext
	Statements() IStatementsContext
	RBRACE() antlr.TerminalNode
	ReturnType() IReturnTypeContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(GoliteParserFUNCTION, 0)
}

func (s *FunctionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *FunctionContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *FunctionContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *FunctionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *FunctionContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *GoliteParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoliteParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(GoliteParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(GoliteParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Parameters()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17179893760) != 0 {
		{
			p.SetState(136)
			p.ReturnType()
		}

	}
	{
		p.SetState(139)
		p.Match(GoliteParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Declarations()
	}
	{
		p.SetState(141)
		p.Statements()
	}
	{
		p.SetState(142)
		p.Match(GoliteParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllDecl() []IDeclContext
	Decl(i int) IDeclContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ParametersContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *GoliteParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoliteParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserIDENTIFIER {
		{
			p.SetState(145)
			p.Decl()
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(146)
				p.Match(GoliteParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(147)
				p.Decl()
			}

			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(155)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_returnType
	return p
}

func InitEmptyReturnTypeContext(p *ReturnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_returnType
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (p *GoliteParser) ReturnType() (localctx IReturnTypeContext) {
	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoliteParserRULE_returnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *GoliteParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoliteParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1099511630816) != 0 {
		{
			p.SetState(159)
			p.Statement()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Print_() IPrintContext
	Read() IReadContext
	Delete_() IDeleteContext
	Conditional() IConditionalContext
	Loop() ILoopContext
	Return_() IReturnContext
	Invocation() IInvocationContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) Read() IReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatementContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) Return_() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *StatementContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GoliteParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoliteParserRULE_statement)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Print_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Read()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.Delete_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(169)
			p.Conditional()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(170)
			p.Loop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(171)
			p.Return_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(172)
			p.Invocation()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	Statements() IStatementsContext
	RBRACE() antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLBRACE, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *GoliteParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoliteParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(GoliteParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Statements()
	}
	{
		p.SetState(177)
		p.Match(GoliteParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
	return p
}

func InitEmptyDeleteContext(p *DeleteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GoliteParserDELETE, 0)
}

func (s *DeleteContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeleteContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (p *GoliteParser) Delete_() (localctx IDeleteContext) {
	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoliteParserRULE_delete)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(GoliteParserDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Expression()
	}
	{
		p.SetState(181)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SCAN() antlr.TerminalNode
	Lvalue() ILvalueContext
	SEMICOLON() antlr.TerminalNode

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_read
	return p
}

func InitEmptyReadContext(p *ReadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_read
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) SCAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserSCAN, 0)
}

func (s *ReadContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *ReadContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *GoliteParser) Read() (localctx IReadContext) {
	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoliteParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(GoliteParserSCAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Lvalue()
	}
	{
		p.SetState(185)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lvalue() ILvalueContext
	EQUALS() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GoliteParserEQUALS, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *GoliteParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoliteParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Lvalue()
	}
	{
		p.SetState(188)
		p.Match(GoliteParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Expression()
	}
	{
		p.SetState(190)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GoliteParserPRINT, 0)
}

func (s *PrintContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *PrintContext) STRING() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRING, 0)
}

func (s *PrintContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *PrintContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *PrintContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *PrintContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *PrintContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *GoliteParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoliteParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(GoliteParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(GoliteParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(195)
			p.Match(GoliteParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Expression()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
	return p
}

func InitEmptyConditionalContext(p *ConditionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(GoliteParserIF, 0)
}

func (s *ConditionalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ConditionalContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ConditionalContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserELSE, 0)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *GoliteParser) Conditional() (localctx IConditionalContext) {
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoliteParserRULE_conditional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(GoliteParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Expression()
	}
	{
		p.SetState(208)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Block()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserELSE {
		{
			p.SetState(210)
			p.Match(GoliteParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoliteParserFOR, 0)
}

func (s *LoopContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *GoliteParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoliteParserRULE_loop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(GoliteParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Expression()
	}
	{
		p.SetState(217)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_return
	return p
}

func InitEmptyReturnContext(p *ReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_return
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRETURN, 0)
}

func (s *ReturnContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReturnContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *GoliteParser) Return_() (localctx IReturnContext) {
	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoliteParserRULE_return)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(GoliteParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3371684036608) != 0 {
		{
			p.SetState(221)
			p.Expression()
		}

	}
	{
		p.SetState(224)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Arguments() IArgumentsContext
	SEMICOLON() antlr.TerminalNode

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
	return p
}

func InitEmptyInvocationContext(p *InvocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, 0)
}

func (s *InvocationContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *InvocationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (p *GoliteParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoliteParserRULE_invocation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(GoliteParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Arguments()
	}
	{
		p.SetState(228)
		p.Match(GoliteParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GoliteParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoliteParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(GoliteParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3371684036608) != 0 {
		{
			p.SetState(231)
			p.Expression()
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(232)
				p.Match(GoliteParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(233)
				p.Expression()
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(241)
		p.Match(GoliteParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_lvalue
	return p
}

func InitEmptyLvalueContext(p *LvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_lvalue
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserIDENTIFIER)
}

func (s *LvalueContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, i)
}

func (s *LvalueContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserDOT)
}

func (s *LvalueContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, i)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *GoliteParser) Lvalue() (localctx ILvalueContext) {
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoliteParserRULE_lvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(GoliteParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(244)
			p.Match(GoliteParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Match(GoliteParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBoolterm() []IBooltermContext
	Boolterm(i int) IBooltermContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllBoolterm() []IBooltermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBooltermContext); ok {
			len++
		}
	}

	tst := make([]IBooltermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBooltermContext); ok {
			tst[i] = t.(IBooltermContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Boolterm(i int) IBooltermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooltermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooltermContext)
}

func (s *ExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserOR)
}

func (s *ExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserOR, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GoliteParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoliteParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Boolterm()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserOR {
		{
			p.SetState(252)
			p.Match(GoliteParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Boolterm()
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooltermContext is an interface to support dynamic dispatch.
type IBooltermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualterm() []IEqualtermContext
	Equalterm(i int) IEqualtermContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsBooltermContext differentiates from other interfaces.
	IsBooltermContext()
}

type BooltermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooltermContext() *BooltermContext {
	var p = new(BooltermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_boolterm
	return p
}

func InitEmptyBooltermContext(p *BooltermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_boolterm
}

func (*BooltermContext) IsBooltermContext() {}

func NewBooltermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooltermContext {
	var p = new(BooltermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolterm

	return p
}

func (s *BooltermContext) GetParser() antlr.Parser { return s.parser }

func (s *BooltermContext) AllEqualterm() []IEqualtermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualtermContext); ok {
			len++
		}
	}

	tst := make([]IEqualtermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualtermContext); ok {
			tst[i] = t.(IEqualtermContext)
			i++
		}
	}

	return tst
}

func (s *BooltermContext) Equalterm(i int) IEqualtermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualtermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualtermContext)
}

func (s *BooltermContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserAND)
}

func (s *BooltermContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserAND, i)
}

func (s *BooltermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooltermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooltermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolterm(s)
	}
}

func (s *BooltermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolterm(s)
	}
}

func (p *GoliteParser) Boolterm() (localctx IBooltermContext) {
	localctx = NewBooltermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoliteParserRULE_boolterm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Equalterm()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserAND {
		{
			p.SetState(260)
			p.Match(GoliteParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Equalterm()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualtermContext is an interface to support dynamic dispatch.
type IEqualtermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRelationterm() []IRelationtermContext
	Relationterm(i int) IRelationtermContext
	AllDOUBLEEQ() []antlr.TerminalNode
	DOUBLEEQ(i int) antlr.TerminalNode
	AllNEQ() []antlr.TerminalNode
	NEQ(i int) antlr.TerminalNode

	// IsEqualtermContext differentiates from other interfaces.
	IsEqualtermContext()
}

type EqualtermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualtermContext() *EqualtermContext {
	var p = new(EqualtermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_equalterm
	return p
}

func InitEmptyEqualtermContext(p *EqualtermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_equalterm
}

func (*EqualtermContext) IsEqualtermContext() {}

func NewEqualtermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualtermContext {
	var p = new(EqualtermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalterm

	return p
}

func (s *EqualtermContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualtermContext) AllRelationterm() []IRelationtermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationtermContext); ok {
			len++
		}
	}

	tst := make([]IRelationtermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationtermContext); ok {
			tst[i] = t.(IRelationtermContext)
			i++
		}
	}

	return tst
}

func (s *EqualtermContext) Relationterm(i int) IRelationtermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationtermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationtermContext)
}

func (s *EqualtermContext) AllDOUBLEEQ() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserDOUBLEEQ)
}

func (s *EqualtermContext) DOUBLEEQ(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserDOUBLEEQ, i)
}

func (s *EqualtermContext) AllNEQ() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserNEQ)
}

func (s *EqualtermContext) NEQ(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserNEQ, i)
}

func (s *EqualtermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualtermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualtermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualterm(s)
	}
}

func (s *EqualtermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualterm(s)
	}
}

func (p *GoliteParser) Equalterm() (localctx IEqualtermContext) {
	localctx = NewEqualtermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoliteParserRULE_equalterm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Relationterm()
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOUBLEEQ || _la == GoliteParserNEQ {
		{
			p.SetState(268)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoliteParserDOUBLEEQ || _la == GoliteParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(269)
			p.Relationterm()
		}

		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationtermContext is an interface to support dynamic dispatch.
type IRelationtermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimpleterm() []ISimpletermContext
	Simpleterm(i int) ISimpletermContext
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllGEQ() []antlr.TerminalNode
	GEQ(i int) antlr.TerminalNode
	AllLEQ() []antlr.TerminalNode
	LEQ(i int) antlr.TerminalNode

	// IsRelationtermContext differentiates from other interfaces.
	IsRelationtermContext()
}

type RelationtermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationtermContext() *RelationtermContext {
	var p = new(RelationtermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_relationterm
	return p
}

func InitEmptyRelationtermContext(p *RelationtermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_relationterm
}

func (*RelationtermContext) IsRelationtermContext() {}

func NewRelationtermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationtermContext {
	var p = new(RelationtermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relationterm

	return p
}

func (s *RelationtermContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationtermContext) AllSimpleterm() []ISimpletermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpletermContext); ok {
			len++
		}
	}

	tst := make([]ISimpletermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpletermContext); ok {
			tst[i] = t.(ISimpletermContext)
			i++
		}
	}

	return tst
}

func (s *RelationtermContext) Simpleterm(i int) ISimpletermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpletermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpletermContext)
}

func (s *RelationtermContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserGT)
}

func (s *RelationtermContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserGT, i)
}

func (s *RelationtermContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserLT)
}

func (s *RelationtermContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserLT, i)
}

func (s *RelationtermContext) AllGEQ() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserGEQ)
}

func (s *RelationtermContext) GEQ(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserGEQ, i)
}

func (s *RelationtermContext) AllLEQ() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserLEQ)
}

func (s *RelationtermContext) LEQ(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserLEQ, i)
}

func (s *RelationtermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationtermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationtermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelationterm(s)
	}
}

func (s *RelationtermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelationterm(s)
	}
}

func (p *GoliteParser) Relationterm() (localctx IRelationtermContext) {
	localctx = NewRelationtermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GoliteParserRULE_relationterm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Simpleterm()
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412342026240) != 0 {
		{
			p.SetState(276)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412342026240) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(277)
			p.Simpleterm()
		}

		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpletermContext is an interface to support dynamic dispatch.
type ISimpletermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsSimpletermContext differentiates from other interfaces.
	IsSimpletermContext()
}

type SimpletermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpletermContext() *SimpletermContext {
	var p = new(SimpletermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleterm
	return p
}

func InitEmptySimpletermContext(p *SimpletermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleterm
}

func (*SimpletermContext) IsSimpletermContext() {}

func NewSimpletermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpletermContext {
	var p = new(SimpletermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleterm

	return p
}

func (s *SimpletermContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpletermContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *SimpletermContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpletermContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserPLUS)
}

func (s *SimpletermContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserPLUS, i)
}

func (s *SimpletermContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserMINUS)
}

func (s *SimpletermContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, i)
}

func (s *SimpletermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpletermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpletermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleterm(s)
	}
}

func (s *SimpletermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleterm(s)
	}
}

func (p *GoliteParser) Simpleterm() (localctx ISimpletermContext) {
	localctx = NewSimpletermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoliteParserRULE_simpleterm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Term()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserPLUS || _la == GoliteParserMINUS {
		{
			p.SetState(284)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoliteParserPLUS || _la == GoliteParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(285)
			p.Term()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryterm() []IUnarytermContext
	Unaryterm(i int) IUnarytermContext
	AllASTERISK() []antlr.TerminalNode
	ASTERISK(i int) antlr.TerminalNode
	AllFSLASH() []antlr.TerminalNode
	FSLASH(i int) antlr.TerminalNode

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllUnaryterm() []IUnarytermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnarytermContext); ok {
			len++
		}
	}

	tst := make([]IUnarytermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnarytermContext); ok {
			tst[i] = t.(IUnarytermContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Unaryterm(i int) IUnarytermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarytermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarytermContext)
}

func (s *TermContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserASTERISK)
}

func (s *TermContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, i)
}

func (s *TermContext) AllFSLASH() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserFSLASH)
}

func (s *TermContext) FSLASH(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserFSLASH, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GoliteParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoliteParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Unaryterm()
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFSLASH || _la == GoliteParserASTERISK {
		{
			p.SetState(292)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoliteParserFSLASH || _la == GoliteParserASTERISK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(293)
			p.Unaryterm()
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnarytermContext is an interface to support dynamic dispatch.
type IUnarytermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Selectorterm() ISelectortermContext
	EXCLAMATION() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsUnarytermContext differentiates from other interfaces.
	IsUnarytermContext()
}

type UnarytermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnarytermContext() *UnarytermContext {
	var p = new(UnarytermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryterm
	return p
}

func InitEmptyUnarytermContext(p *UnarytermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryterm
}

func (*UnarytermContext) IsUnarytermContext() {}

func NewUnarytermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnarytermContext {
	var p = new(UnarytermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_unaryterm

	return p
}

func (s *UnarytermContext) GetParser() antlr.Parser { return s.parser }

func (s *UnarytermContext) Selectorterm() ISelectortermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectortermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectortermContext)
}

func (s *UnarytermContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(GoliteParserEXCLAMATION, 0)
}

func (s *UnarytermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *UnarytermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarytermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnarytermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterUnaryterm(s)
	}
}

func (s *UnarytermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitUnaryterm(s)
	}
}

func (p *GoliteParser) Unaryterm() (localctx IUnarytermContext) {
	localctx = NewUnarytermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoliteParserRULE_unaryterm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserEXCLAMATION || _la == GoliteParserMINUS {
		{
			p.SetState(299)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoliteParserEXCLAMATION || _la == GoliteParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(302)
		p.Selectorterm()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectortermContext is an interface to support dynamic dispatch.
type ISelectortermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsSelectortermContext differentiates from other interfaces.
	IsSelectortermContext()
}

type SelectortermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectortermContext() *SelectortermContext {
	var p = new(SelectortermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorterm
	return p
}

func InitEmptySelectortermContext(p *SelectortermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorterm
}

func (*SelectortermContext) IsSelectortermContext() {}

func NewSelectortermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectortermContext {
	var p = new(SelectortermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorterm

	return p
}

func (s *SelectortermContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectortermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *SelectortermContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserDOT)
}

func (s *SelectortermContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, i)
}

func (s *SelectortermContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserIDENTIFIER)
}

func (s *SelectortermContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, i)
}

func (s *SelectortermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectortermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectortermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorterm(s)
	}
}

func (s *SelectortermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorterm(s)
	}
}

func (p *GoliteParser) Selectorterm() (localctx ISelectortermContext) {
	localctx = NewSelectortermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoliteParserRULE_selectorterm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Factor()
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(305)
			p.Match(GoliteParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(GoliteParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Arguments() IArgumentsContext
	NUMBER() antlr.TerminalNode
	NEW() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NIL() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *FactorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoliteParserIDENTIFIER, 0)
}

func (s *FactorContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FactorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GoliteParserNUMBER, 0)
}

func (s *FactorContext) NEW() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEW, 0)
}

func (s *FactorContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTRUE, 0)
}

func (s *FactorContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserFALSE, 0)
}

func (s *FactorContext) NIL() antlr.TerminalNode {
	return s.GetToken(GoliteParserNIL, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *GoliteParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoliteParserRULE_factor)
	var _la int

	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoliteParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(312)
			p.Match(GoliteParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Expression()
		}
		{
			p.SetState(314)
			p.Match(GoliteParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)
			p.Match(GoliteParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GoliteParserLPAREN {
			{
				p.SetState(317)
				p.Arguments()
			}

		}

	case GoliteParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(320)
			p.Match(GoliteParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserNEW:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(321)
			p.Match(GoliteParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(GoliteParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(323)
			p.Match(GoliteParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(324)
			p.Match(GoliteParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoliteParserNIL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(325)
			p.Match(GoliteParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
