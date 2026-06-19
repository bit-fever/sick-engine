//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

lexer grammar TiqLexer;

//=== Keywords ================================================================

PACKAGE  : 'package';
CONST    : 'const';
VAR      : 'var';
FUNC     : 'func';

IF       : 'if';
ELSE     : 'else';
FOR      : 'for';
TRUE     : 'true';
FALSE    : 'false';
RETURN   : 'return';
AND      : 'and';
OR       : 'or';
NOT      : 'not';
NULL     : 'null';
THIS     : 'this';
NEW      : 'new';

//--- Bar

BAR         : 'bar';
OPEN        : 'open';
HIGH        : 'high';
LOW         : 'low';
CLOSE       : 'close';

//--- Orders

BUY         : 'buy';
SELL        : 'sell';
SELL_SHORT  : 'sellShort';
BUY_TO_COVER: 'buyToCover';
AT          : 'at';
MARKET      : 'market';
STOP        : 'stop';
LIMIT       : 'limit';
CONTRACTS   : 'contracts';

//--- Datatypes

INT        : 'int';
REAL       : 'real';
BOOL       : 'bool';
STRING     : 'string';
TIME       : 'time';
DATE       : 'date';
TIMESERIES : 'timeseries';
ENUM       : 'enum';
CLASS      : 'class';
ERROR      : 'error';
LIST       : 'list';
MAP        : 'map';

//=== Identifiers =============================================================

IDENTIFIER : [a-zA-Z]+ ( '_' | [a-zA-Z] | [0-9] )* ;

//=== Values =================================================================

INT_VALUE : DIGIT+;

REAL_VALUE
    : DIGIT+ '.' DIGIT*
    | '.' DIGIT+
    ;

DIGIT: [0-9];

STRING_VALUE : '"' (~["\\] | ESCAPED_VALUE)* '"' ;

fragment ESCAPED_VALUE: '\\' [bnrt\\"];

//=== Symbols =================================================================

L_PAREN    : '(' ;
R_PAREN    : ')' ;
L_BRACKET  : '[' ;
R_BRACKET  : ']' ;
L_CURLY    : '{' ;
R_CURLY    : '}' ;
EQUAL      : '=' ;

STAR       : '*' ;
SLASH      : '/' ;
PLUS       : '+' ;
MINUS      : '-' ;
COMMA      : ',' ;
DOT        : '.' ;
COLON      : ':' ;
APOS       : '\'';

//=== Relation operators

NOT_EQUAL        : '<>';
LESS_THAN        : '<';
LESS_OR_EQUAL    : '<=';
GREATER_THAN     : '>';
GREATER_OR_EQUAL : '>=';

//=== Comments and white spaces ===============================================

WS           : [ \t]+        -> skip;
NEWLINE      : [\r\n]+       -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

//=============================================================================
