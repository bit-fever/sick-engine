//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

parser grammar TiqParser;

options {
    tokenVocab = TiqLexer;
}

//=============================================================================

script
    : packageDef? ( constantsDef | variablesDef | functionDef | enumDef | classDef )*
    ;

//=============================================================================
//===
//=== Top entities
//===
//=============================================================================

packageDef
    : PACKAGE IDENTIFIER
    ;

//=============================================================================

constantsDef
    : CONST ( L_CURLY ( constantDef )* R_CURLY | constantDef )
    ;

constantDef
    : IDENTIFIER EQUAL expression
    ;

//=============================================================================

variablesDef
    : VAR ( L_CURLY ( variableDef )* R_CURLY | variableDef )
    ;

variableDef
    : IDENTIFIER EQUAL expression
    ;

//=============================================================================

functionDef
    : FUNC (class)? IDENTIFIER parameters results? block
    ;

class
    : L_PAREN fqIdentifier R_PAREN
    ;

parameters
    : L_PAREN (parameterDecl (COMMA parameterDecl)* )? R_PAREN
    ;

parameterDecl
    : IDENTIFIER type
    ;

results
    : type (COMMA type)*
    ;

//=============================================================================

enumDef
    : ENUM IDENTIFIER L_CURLY (enumItem)* R_CURLY;

enumItem
    : IDENTIFIER ( L_PAREN enumValue R_PAREN )?
    ;

enumValue
    : INT_VALUE | STRING_VALUE;

//=============================================================================

classDef
    : CLASS IDENTIFIER L_CURLY (property)* R_CURLY;

property
    : IDENTIFIER type
    ;

//=============================================================================
//===
//=== Types
//===
//=============================================================================

type
    : INT
    | REAL
    | BOOL
    | STRING
    | TIME
    | DATE
    | timeSeriesType
    | fqIdentifier
    | ERROR
    | listType
    | mapType
    ;

timeSeriesType
    : TIMESERIES ( L_PAREN type R_PAREN )?
    ;

listType
    : LIST L_PAREN type R_PAREN
    ;

mapType
    : MAP L_PAREN keyType COMMA type R_PAREN
    ;

keyType
    : INT
    | STRING
    | TIME
    | DATE
    ;

//=============================================================================
//===
//=== Statements
//===
//=============================================================================

block
    : L_CURLY (statement)* R_CURLY;

//=============================================================================

statement
    : varsDeclaration
    | varsAssignmentOrFunctionCall
    | ifStatement
    | returnStatement
    ;

//=============================================================================

varsDeclaration
    : VAR varDeclaration ( COMMA varDeclaration )*
    ;

varDeclaration
    : IDENTIFIER ( COLON type )? (EQUAL expression)?
    ;

//=============================================================================

varsAssignmentOrFunctionCall
    : chainedExpression ( COMMA chainedExpression)* ( EQUAL expression ( COMMA expression)* )?
    ;

//=============================================================================

ifStatement
    : IF expression block (elseIfBlock)* (elseBlock)?
    ;

elseIfBlock : ELSE IF expression block;

elseBlock : ELSE block;

//=============================================================================

returnStatement
    : RETURN
    | RETURN expression ( COMMA expression)*
    ;

//=============================================================================
//===
//=== Expressions
//===
//=============================================================================

expression
    : unaryExpression
    | expression (STAR | SLASH) expression
    | expression (PLUS | MINUS) expression
    | expression ( EQUAL | NOT_EQUAL | LESS_THAN | LESS_OR_EQUAL | GREATER_THAN | GREATER_OR_EQUAL ) expression
    | NOT expression
    | expression ( AND expression )+
    | expression ( OR  expression )+
    | listExpression
    | mapExpression
    ;

//=============================================================================

unaryExpression
    : (MINUS|PLUS) unaryExpression
    | expressionInParenthesis
    | barExpression
    | constantValueExpression
    | chainedExpression
    ;

//=============================================================================

expressionInParenthesis
    : L_PAREN expression R_PAREN;

//=============================================================================

chainedExpression
    : ( THIS DOT )? chainItem ( DOT chainItem )*
    | NEW fqIdentifier paramsExpression ( DOT chainItem )*
    ;

chainItem
    : IDENTIFIER paramsExpression? accessorExpression?
    ;

paramsExpression
    : L_PAREN (expression (COMMA expression)* )? R_PAREN
    ;

accessorExpression
    : L_BRACKET expression R_BRACKET
    ;

//=============================================================================

barExpression
    : BAR DOT ( OPEN | HIGH | LOW | CLOSE ) accessorExpression?
    ;

//=============================================================================

listExpression
    : L_BRACKET expression ( COMMA expression)* R_BRACKET
    ;

//=============================================================================

mapExpression
    : L_CURLY keyValueCouple ( COMMA keyValueCouple)* R_CURLY
    ;

keyValueCouple
    : keyValue COLON expression
    ;

keyValue
    : INT_VALUE
    | STRING_VALUE
    | timeValue
    | dateValue
    ;

//=============================================================================

constantValueExpression
    : INT_VALUE
    | REAL_VALUE
    | boolValue
    | STRING_VALUE
    | timeValue
    | dateValue
    | errorValue
    ;

timeValue
    : APOS INT_VALUE COLON INT_VALUE APOS;

dateValue
    : APOS INT_VALUE MINUS INT_VALUE MINUS INT_VALUE APOS;

boolValue
    : TRUE | FALSE ;

errorValue
    : ERROR L_PAREN STRING_VALUE R_PAREN;

//=============================================================================

fqIdentifier
    : IDENTIFIER ( DOT IDENTIFIER )?
    ;

//=============================================================================
