// Code generated from ./PegiParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // PegiParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePegiParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePegiParserVisitor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitParserClause(ctx *ParserClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitShortVarDecl(ctx *ShortVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitPipe(ctx *PipeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitPipeUnit(ctx *PipeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitPipeUnitExpression(ctx *PipeUnitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitOperandName(ctx *OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitSliceValue(ctx *SliceValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitSliceValueSingleLine(ctx *SliceValueSingleLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitSliceValueMultiLine(ctx *SliceValueMultiLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitMessageValue(ctx *MessageValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitKeyedElementKV(ctx *KeyedElementKVContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitKeyedElement(ctx *KeyedElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitMessageValueSingleLine(ctx *MessageValueSingleLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitMessageValueMultiLine(ctx *MessageValueMultiLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitLineKeyedElement(ctx *LineKeyedElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitSliceType(ctx *SliceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitTypeT(ctx *TypeTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitTypeLit(ctx *TypeLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitFieldDecl(ctx *FieldDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePegiParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
