// Code generated from ./PegiParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // PegiParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PegiParserListener is a complete listener for a parse tree produced by PegiParser.
type PegiParserListener interface {
	antlr.ParseTreeListener

	// EnterSourceFile is called when entering the sourceFile production.
	EnterSourceFile(c *SourceFileContext)

	// EnterParserClause is called when entering the parserClause production.
	EnterParserClause(c *ParserClauseContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterImportSpec is called when entering the importSpec production.
	EnterImportSpec(c *ImportSpecContext)

	// EnterImportPath is called when entering the importPath production.
	EnterImportPath(c *ImportPathContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterShortVarDecl is called when entering the shortVarDecl production.
	EnterShortVarDecl(c *ShortVarDeclContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

	// EnterPipeUnit is called when entering the pipeUnit production.
	EnterPipeUnit(c *PipeUnitContext)

	// EnterPipeUnitIdentifers is called when entering the pipeUnitIdentifers production.
	EnterPipeUnitIdentifers(c *PipeUnitIdentifersContext)

	// EnterPipeUnitExpression is called when entering the pipeUnitExpression production.
	EnterPipeUnitExpression(c *PipeUnitExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterBasicLit is called when entering the basicLit production.
	EnterBasicLit(c *BasicLitContext)

	// EnterOperandName is called when entering the operandName production.
	EnterOperandName(c *OperandNameContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterSliceValue is called when entering the sliceValue production.
	EnterSliceValue(c *SliceValueContext)

	// EnterSliceValueSingleLine is called when entering the sliceValueSingleLine production.
	EnterSliceValueSingleLine(c *SliceValueSingleLineContext)

	// EnterSliceValueMultiLine is called when entering the sliceValueMultiLine production.
	EnterSliceValueMultiLine(c *SliceValueMultiLineContext)

	// EnterSliceValueMultiLineLine is called when entering the sliceValueMultiLineLine production.
	EnterSliceValueMultiLineLine(c *SliceValueMultiLineLineContext)

	// EnterMessageValue is called when entering the messageValue production.
	EnterMessageValue(c *MessageValueContext)

	// EnterKeyedElementKV is called when entering the keyedElementKV production.
	EnterKeyedElementKV(c *KeyedElementKVContext)

	// EnterKeyedElementIdentifer is called when entering the keyedElementIdentifer production.
	EnterKeyedElementIdentifer(c *KeyedElementIdentiferContext)

	// EnterKeyedElement is called when entering the keyedElement production.
	EnterKeyedElement(c *KeyedElementContext)

	// EnterMessageValueSingleLine is called when entering the messageValueSingleLine production.
	EnterMessageValueSingleLine(c *MessageValueSingleLineContext)

	// EnterMessageValueMultiLine is called when entering the messageValueMultiLine production.
	EnterMessageValueMultiLine(c *MessageValueMultiLineContext)

	// EnterLineKeyedElement is called when entering the lineKeyedElement production.
	EnterLineKeyedElement(c *LineKeyedElementContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterSliceType is called when entering the sliceType production.
	EnterSliceType(c *SliceTypeContext)

	// EnterTypeT is called when entering the typeT production.
	EnterTypeT(c *TypeTContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterTypeLit is called when entering the typeLit production.
	EnterTypeLit(c *TypeLitContext)

	// EnterInterfaceType is called when entering the interfaceType production.
	EnterInterfaceType(c *InterfaceTypeContext)

	// EnterFieldDecl is called when entering the fieldDecl production.
	EnterFieldDecl(c *FieldDeclContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitParserClause is called when exiting the parserClause production.
	ExitParserClause(c *ParserClauseContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitImportSpec is called when exiting the importSpec production.
	ExitImportSpec(c *ImportSpecContext)

	// ExitImportPath is called when exiting the importPath production.
	ExitImportPath(c *ImportPathContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitShortVarDecl is called when exiting the shortVarDecl production.
	ExitShortVarDecl(c *ShortVarDeclContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

	// ExitPipeUnit is called when exiting the pipeUnit production.
	ExitPipeUnit(c *PipeUnitContext)

	// ExitPipeUnitIdentifers is called when exiting the pipeUnitIdentifers production.
	ExitPipeUnitIdentifers(c *PipeUnitIdentifersContext)

	// ExitPipeUnitExpression is called when exiting the pipeUnitExpression production.
	ExitPipeUnitExpression(c *PipeUnitExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitBasicLit is called when exiting the basicLit production.
	ExitBasicLit(c *BasicLitContext)

	// ExitOperandName is called when exiting the operandName production.
	ExitOperandName(c *OperandNameContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitSliceValue is called when exiting the sliceValue production.
	ExitSliceValue(c *SliceValueContext)

	// ExitSliceValueSingleLine is called when exiting the sliceValueSingleLine production.
	ExitSliceValueSingleLine(c *SliceValueSingleLineContext)

	// ExitSliceValueMultiLine is called when exiting the sliceValueMultiLine production.
	ExitSliceValueMultiLine(c *SliceValueMultiLineContext)

	// ExitSliceValueMultiLineLine is called when exiting the sliceValueMultiLineLine production.
	ExitSliceValueMultiLineLine(c *SliceValueMultiLineLineContext)

	// ExitMessageValue is called when exiting the messageValue production.
	ExitMessageValue(c *MessageValueContext)

	// ExitKeyedElementKV is called when exiting the keyedElementKV production.
	ExitKeyedElementKV(c *KeyedElementKVContext)

	// ExitKeyedElementIdentifer is called when exiting the keyedElementIdentifer production.
	ExitKeyedElementIdentifer(c *KeyedElementIdentiferContext)

	// ExitKeyedElement is called when exiting the keyedElement production.
	ExitKeyedElement(c *KeyedElementContext)

	// ExitMessageValueSingleLine is called when exiting the messageValueSingleLine production.
	ExitMessageValueSingleLine(c *MessageValueSingleLineContext)

	// ExitMessageValueMultiLine is called when exiting the messageValueMultiLine production.
	ExitMessageValueMultiLine(c *MessageValueMultiLineContext)

	// ExitLineKeyedElement is called when exiting the lineKeyedElement production.
	ExitLineKeyedElement(c *LineKeyedElementContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitSliceType is called when exiting the sliceType production.
	ExitSliceType(c *SliceTypeContext)

	// ExitTypeT is called when exiting the typeT production.
	ExitTypeT(c *TypeTContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitTypeLit is called when exiting the typeLit production.
	ExitTypeLit(c *TypeLitContext)

	// ExitInterfaceType is called when exiting the interfaceType production.
	ExitInterfaceType(c *InterfaceTypeContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
