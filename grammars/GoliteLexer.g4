lexer grammar GoliteLexer;

COMMENT      : '//' ~[\r\n]* -> skip            ;

TYPE         : 'type'                           ;
STRUCT       : 'struct'                         ;
FUNCTION     : 'func'                           ;
PRINT        : 'printf'                         ;
DELETE       : 'delete'                         ;
SCAN         : 'scan'                           ;
FOR          : 'for'                            ;
IF           : 'if'                             ;
ELSE         : 'else'                           ;
RETURN       : 'return'                         ;
VAR          : 'var'                            ;
INT          : 'int'                            ;
BOOL         : 'bool'                           ;
TRUE         : 'true'                           ;
FALSE        : 'false'                          ;
NIL          : 'nil'                            ;
NEW          : 'new'                            ;

OR           : '||'                             ;
AND          : '&&'                             ;
DOUBLEEQ     : '=='                             ;
NEQ          : '!='                             ;
GEQ          : '>='                             ;
LEQ          : '<='                             ;

LBRACE       : '{'                              ;
RBRACE       : '}'                              ;
LPAREN       : '('                              ;
RPAREN       : ')'                              ;
SEMICOLON    : ';'                              ;
COMMA        : ','                              ;
DOT          : '.'                              ;

EXCLAMATION  : '!'                              ;
FSLASH       : '/'                              ;
ASTERISK     : '*'                              ;
PLUS         : '+'                              ;
MINUS        : '-'                              ;

GT           : '>'                              ;
LT           : '<'                              ;
EQUALS       : '='                              ;

IDENTIFIER   : [a-zA-Z][a-zA-Z0-9]*             ;
NUMBER       : '0' | [1-9][0-9]*                ;

STRING       : '"' (~["\\\r\n] | '\\' .)* '"'   ;

WS           : [ \r\t\n]+ -> skip               ;