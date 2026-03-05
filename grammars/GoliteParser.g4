parser grammar GoliteParser;

options {
    tokenVocab = GoliteLexer;
}

program
    : types declarations functions EOF
    ;

types
    : typeDeclaration*
    ;

typeDeclaration
    : TYPE IDENTIFIER STRUCT LBRACE fields RBRACE SEMICOLON
    ;

fields
    : decl SEMICOLON (decl SEMICOLON)*
    ;

decl
    : IDENTIFIER type
    ;

type
    : INT
    | BOOL
    | ASTERISK IDENTIFIER
    ;

declarations
    : declaration*
    ;

declaration
    : VAR ids type SEMICOLON
    ;

ids
    : IDENTIFIER (COMMA IDENTIFIER)*
    ;

functions
    : function*
    ;

function
    : FUNCTION IDENTIFIER parameters returnType? LBRACE declarations statements RBRACE
    ;

parameters
    : LPAREN (decl (COMMA decl)*)? RPAREN
    ;

returnType
    : type
    ;

statements
    : statement* 
    ;

statement
    : assignment
    | print
    | read
    | delete
    | conditional
    | loop
    | return
    | invocation
    ;

block
    : LBRACE statements RBRACE
    ;

delete
    : DELETE expression SEMICOLON
    ;

read
    : SCAN lvalue SEMICOLON
    ;

assignment
    : lvalue EQUALS expression SEMICOLON
    ;

print
    : PRINT LPAREN STRING (COMMA expression)* RPAREN SEMICOLON
    ;

conditional
    : IF LPAREN expression RPAREN block (ELSE block)?
    ;

loop
    : FOR LPAREN expression RPAREN block
    ;

return
    : RETURN expression? SEMICOLON
    ;

invocation
    : IDENTIFIER arguments SEMICOLON
    ;

arguments
    : LPAREN (expression (COMMA expression)*)? RPAREN
    ;

lvalue
    : IDENTIFIER (DOT IDENTIFIER)*
    ;

expression
    : boolterm (OR boolterm)*
    ;

boolterm
    : equalterm (AND equalterm)*
    ;

equalterm
    : relationterm ((DOUBLEEQ | NEQ) relationterm)*
    ;

relationterm
    : simpleterm ((GT | LT | GEQ | LEQ) simpleterm)*
    ;

simpleterm
    : term ((PLUS | MINUS) term)*
    ;

term
    : unaryterm ((ASTERISK | FSLASH) unaryterm)*
    ;

unaryterm
    : (EXCLAMATION | MINUS)? selectorterm
    ;

selectorterm
    : factor (DOT IDENTIFIER)*
    ;

factor
    : LPAREN expression RPAREN
    | IDENTIFIER arguments?
    | NUMBER
    | NEW IDENTIFIER
    | TRUE
    | FALSE
    | NIL
    ;