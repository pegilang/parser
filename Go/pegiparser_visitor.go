// Code generated from ./PegiParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // PegiParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PegiParser.
type PegiParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PegiParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by PegiParser#parserClause.
	VisitParserClause(ctx *ParserClauseContext) interface{}

	// Visit a parse tree produced by PegiParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by PegiParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by PegiParser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}

	// Visit a parse tree produced by PegiParser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by PegiParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by PegiParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by PegiParser#shortVarDecl.
	VisitShortVarDecl(ctx *ShortVarDeclContext) interface{}

	// Visit a parse tree produced by PegiParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by PegiParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by PegiParser#pipe.
	VisitPipe(ctx *PipeContext) interface{}

	// Visit a parse tree produced by PegiParser#pipeUnit.
	VisitPipeUnit(ctx *PipeUnitContext) interface{}

	// Visit a parse tree produced by PegiParser#pipeUnitIdentifers.
	VisitPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) interface{}

	// Visit a parse tree produced by PegiParser#pipeUnitExpression.
	VisitPipeUnitExpression(ctx *PipeUnitExpressionContext) interface{}

	// Visit a parse tree produced by PegiParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PegiParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by PegiParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by PegiParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by PegiParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by PegiParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by PegiParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by PegiParser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by PegiParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by PegiParser#sliceValue.
	VisitSliceValue(ctx *SliceValueContext) interface{}

	// Visit a parse tree produced by PegiParser#sliceValueSingleLine.
	VisitSliceValueSingleLine(ctx *SliceValueSingleLineContext) interface{}

	// Visit a parse tree produced by PegiParser#sliceValueMultiLine.
	VisitSliceValueMultiLine(ctx *SliceValueMultiLineContext) interface{}

	// Visit a parse tree produced by PegiParser#sliceValueMultiLineLine.
	VisitSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) interface{}

	// Visit a parse tree produced by PegiParser#messageValue.
	VisitMessageValue(ctx *MessageValueContext) interface{}

	// Visit a parse tree produced by PegiParser#keyedElementKV.
	VisitKeyedElementKV(ctx *KeyedElementKVContext) interface{}

	// Visit a parse tree produced by PegiParser#keyedElementIdentifer.
	VisitKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) interface{}

	// Visit a parse tree produced by PegiParser#keyedElement.
	VisitKeyedElement(ctx *KeyedElementContext) interface{}

	// Visit a parse tree produced by PegiParser#messageValueSingleLine.
	VisitMessageValueSingleLine(ctx *MessageValueSingleLineContext) interface{}

	// Visit a parse tree produced by PegiParser#messageValueMultiLine.
	VisitMessageValueMultiLine(ctx *MessageValueMultiLineContext) interface{}

	// Visit a parse tree produced by PegiParser#lineKeyedElement.
	VisitLineKeyedElement(ctx *LineKeyedElementContext) interface{}

	// Visit a parse tree produced by PegiParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by PegiParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by PegiParser#sliceType.
	VisitSliceType(ctx *SliceTypeContext) interface{}

	// Visit a parse tree produced by PegiParser#typeT.
	VisitTypeT(ctx *TypeTContext) interface{}

	// Visit a parse tree produced by PegiParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by PegiParser#typeLit.
	VisitTypeLit(ctx *TypeLitContext) interface{}

	// Visit a parse tree produced by PegiParser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by PegiParser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by PegiParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by PegiParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
