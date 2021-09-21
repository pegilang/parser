parser grammar PegiParser;

options {
	tokenVocab = PegiLexer;
	superClass = PegiParserBase;
}

sourceFile:	(parserClause eos)? (importDecl eos)* statementList? EOF;

parserClause: PARSER parserPath = string_;

importDecl:
	IMPORT (importSpec | L_PAREN (importSpec eos)* R_PAREN);

importSpec: alias = (DOT | IDENTIFIER)? importPath;

importPath: string_;



string_: RAW_STRING_LIT | INTERPRETED_STRING_LIT;
//eof_string: EOF_START eos (LINE eos)* EOF_END;
//block: L_CURLY statementList? R_CURLY;
statementList: (statement eos)+;

statement: (
    shortVarDecl
    | functionCall
    | interfaceType
//	| declaration
//	| labeledStmt
//	| goStmt
//	| returnStmt
//	| breakStmt
//	| continueStmt
//	| gotoStmt
//	| fallthroughStmt
//	| block
//	| ifStmt
//	| switchStmt
//	| selectStmt
//	| forStmt
//	| deferStmt;
);

shortVarDecl: identifierList DECLARE_ASSIGN (expressionList | pipe);
expressionList: expression (COMMA expression)*;
identifierList: IDENTIFIER (COMMA IDENTIFIER)*;

pipe: pipeUnit (pipeUnit)+;
pipeUnit: pipeUnitIdentifers | pipeUnitExpression;
pipeUnitIdentifers: eos OR identifierList;
pipeUnitExpression: eos OR expression;

expression:
	primaryExpr
	| unaryExpr
	| expression mul_op = (
		STAR
		| DIV
//		| MOD
//		| LSHIFT
//		| RSHIFT
//		| AMPERSAND
//		| BIT_CLEAR
	) expression
	| expression add_op = (PLUS | MINUS) expression
	| expression rel_op = (
		EQUALS
		| NOT_EQUALS
		| LESS
		| LESS_OR_EQUALS
		| GREATER
		| GREATER_OR_EQUALS
	) expression
	| expression LOGICAL_AND expression
	| expression LOGICAL_OR expression
;


primaryExpr:
	operand
//	| conversion
	| functionCall
;
//methodExpr: receiverType DOT IDENTIFIER;
//receiverType: type_;
operand: literal | operandName | L_PAREN expression R_PAREN;
literal: basicLit; // | compositeLit | functionLit;
functionCall: typeName arguments;

arguments:
	L_PAREN (
		(expressionList | typeT (COMMA expressionList)?) ELLIPSIS? COMMA?
	)? R_PAREN;

basicLit:
//	NIL_LIT
	integer
	| messageValue
	| sliceValue
	| string_
	| FLOAT_LIT
//	| IMAGINARY_LIT
	| RUNE_LIT;

operandName: IDENTIFIER (DOT IDENTIFIER)?;

unaryExpr:
	primaryExpr
	| unary_op = (
	    UNARY_MINUS
	    | EXCLAMATION
	) expression;

sliceValue: sliceValueSingleLine | sliceValueMultiLine;
sliceValueSingleLine: L_BRACKET expression (COMMA expression)* R_BRACKET;
sliceValueMultiLine: L_BRACKET eos sliceValueMultiLineLine* R_BRACKET;
sliceValueMultiLineLine: expression COMMA eos;

messageValue: messageValueSingleLine | messageValueMultiLine;
//structType: STRUCT; // L_CURLY (fieldDecl eos)* R_CURLY;
//elementList: keyedElement (COMMA keyedElement)*;
keyedElementKV: key COLON element;
keyedElementIdentifer: IDENTIFIER;
keyedElement: keyedElementIdentifer | keyedElementKV;

messageValueSingleLine: L_CURLY keyedElement? (COMMA keyedElement)* R_CURLY;
messageValueMultiLine: L_CURLY eos (lineKeyedElement)+ R_CURLY;
lineKeyedElement: keyedElement COMMA eos;

key: IDENTIFIER; // | expression | messageValue;
element: expression | messageValue | sliceValue;

sliceType: L_BRACKET R_BRACKET typeT;
typeT: typeName | typeLit;
//type_: ; // | L_PAREN type_ R_PAREN;
typeName: IDENTIFIER (DOT IDENTIFIER)*;
typeLit:
//	arrayType
//	structType
//	| pointerType
//	| functionType
	interfaceType
//	| sliceType
//	| mapType
//	| channelType;
;
interfaceType:
	INTERFACE interfaceName = IDENTIFIER L_CURLY ((fieldDecl | typeName) eos)* R_CURLY;

fieldDecl: {p.noTerminatorBetween(2)}? IDENTIFIER typeT;

//methodSpec:
//	{noTerminatorAfterParams(2)}? IDENTIFIER parameters result
//	| IDENTIFIER parameters;
//parameters:
//	L_PAREN (parameterDecl (COMMA parameterDecl)* COMMA?)? R_PAREN;

//fieldDecl: (
//		{noTerminatorBetween(2)}? identifierList type_
//		| embeddedField
//	) tag = string_?;

integer:
	DECIMAL_LIT
	| BINARY_LIT
	| OCTAL_LIT
	| HEX_LIT
//	| IMAGINARY_LIT
	| RUNE_LIT;

eos:
	SEMI
	| EOF
	| {p.lineTerminatorAhead()}?
	| {p.checkPreviousTokenText("}")}?;
