// Code generated from ./PegiParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // PegiParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePegiParserListener is a complete listener for a parse tree produced by PegiParser.
type BasePegiParserListener struct{}

var _ PegiParserListener = &BasePegiParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePegiParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePegiParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePegiParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePegiParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BasePegiParserListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BasePegiParserListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterParserClause is called when production parserClause is entered.
func (s *BasePegiParserListener) EnterParserClause(ctx *ParserClauseContext) {}

// ExitParserClause is called when production parserClause is exited.
func (s *BasePegiParserListener) ExitParserClause(ctx *ParserClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BasePegiParserListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BasePegiParserListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BasePegiParserListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BasePegiParserListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *BasePegiParserListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BasePegiParserListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BasePegiParserListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BasePegiParserListener) ExitString_(ctx *String_Context) {}

// EnterStatementList is called when production statementList is entered.
func (s *BasePegiParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BasePegiParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePegiParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePegiParserListener) ExitStatement(ctx *StatementContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *BasePegiParserListener) EnterShortVarDecl(ctx *ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *BasePegiParserListener) ExitShortVarDecl(ctx *ShortVarDeclContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasePegiParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasePegiParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BasePegiParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BasePegiParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BasePegiParserListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BasePegiParserListener) ExitPipe(ctx *PipeContext) {}

// EnterPipeUnit is called when production pipeUnit is entered.
func (s *BasePegiParserListener) EnterPipeUnit(ctx *PipeUnitContext) {}

// ExitPipeUnit is called when production pipeUnit is exited.
func (s *BasePegiParserListener) ExitPipeUnit(ctx *PipeUnitContext) {}

// EnterPipeUnitIdentifers is called when production pipeUnitIdentifers is entered.
func (s *BasePegiParserListener) EnterPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) {}

// ExitPipeUnitIdentifers is called when production pipeUnitIdentifers is exited.
func (s *BasePegiParserListener) ExitPipeUnitIdentifers(ctx *PipeUnitIdentifersContext) {}

// EnterPipeUnitExpression is called when production pipeUnitExpression is entered.
func (s *BasePegiParserListener) EnterPipeUnitExpression(ctx *PipeUnitExpressionContext) {}

// ExitPipeUnitExpression is called when production pipeUnitExpression is exited.
func (s *BasePegiParserListener) ExitPipeUnitExpression(ctx *PipeUnitExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePegiParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePegiParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BasePegiParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BasePegiParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BasePegiParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BasePegiParserListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePegiParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePegiParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasePegiParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasePegiParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasePegiParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasePegiParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BasePegiParserListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BasePegiParserListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BasePegiParserListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BasePegiParserListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BasePegiParserListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BasePegiParserListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterSliceValue is called when production sliceValue is entered.
func (s *BasePegiParserListener) EnterSliceValue(ctx *SliceValueContext) {}

// ExitSliceValue is called when production sliceValue is exited.
func (s *BasePegiParserListener) ExitSliceValue(ctx *SliceValueContext) {}

// EnterSliceValueSingleLine is called when production sliceValueSingleLine is entered.
func (s *BasePegiParserListener) EnterSliceValueSingleLine(ctx *SliceValueSingleLineContext) {}

// ExitSliceValueSingleLine is called when production sliceValueSingleLine is exited.
func (s *BasePegiParserListener) ExitSliceValueSingleLine(ctx *SliceValueSingleLineContext) {}

// EnterSliceValueMultiLine is called when production sliceValueMultiLine is entered.
func (s *BasePegiParserListener) EnterSliceValueMultiLine(ctx *SliceValueMultiLineContext) {}

// ExitSliceValueMultiLine is called when production sliceValueMultiLine is exited.
func (s *BasePegiParserListener) ExitSliceValueMultiLine(ctx *SliceValueMultiLineContext) {}

// EnterSliceValueMultiLineLine is called when production sliceValueMultiLineLine is entered.
func (s *BasePegiParserListener) EnterSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) {}

// ExitSliceValueMultiLineLine is called when production sliceValueMultiLineLine is exited.
func (s *BasePegiParserListener) ExitSliceValueMultiLineLine(ctx *SliceValueMultiLineLineContext) {}

// EnterMessageValue is called when production messageValue is entered.
func (s *BasePegiParserListener) EnterMessageValue(ctx *MessageValueContext) {}

// ExitMessageValue is called when production messageValue is exited.
func (s *BasePegiParserListener) ExitMessageValue(ctx *MessageValueContext) {}

// EnterKeyedElementKV is called when production keyedElementKV is entered.
func (s *BasePegiParserListener) EnterKeyedElementKV(ctx *KeyedElementKVContext) {}

// ExitKeyedElementKV is called when production keyedElementKV is exited.
func (s *BasePegiParserListener) ExitKeyedElementKV(ctx *KeyedElementKVContext) {}

// EnterKeyedElementIdentifer is called when production keyedElementIdentifer is entered.
func (s *BasePegiParserListener) EnterKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) {}

// ExitKeyedElementIdentifer is called when production keyedElementIdentifer is exited.
func (s *BasePegiParserListener) ExitKeyedElementIdentifer(ctx *KeyedElementIdentiferContext) {}

// EnterKeyedElement is called when production keyedElement is entered.
func (s *BasePegiParserListener) EnterKeyedElement(ctx *KeyedElementContext) {}

// ExitKeyedElement is called when production keyedElement is exited.
func (s *BasePegiParserListener) ExitKeyedElement(ctx *KeyedElementContext) {}

// EnterMessageValueSingleLine is called when production messageValueSingleLine is entered.
func (s *BasePegiParserListener) EnterMessageValueSingleLine(ctx *MessageValueSingleLineContext) {}

// ExitMessageValueSingleLine is called when production messageValueSingleLine is exited.
func (s *BasePegiParserListener) ExitMessageValueSingleLine(ctx *MessageValueSingleLineContext) {}

// EnterMessageValueMultiLine is called when production messageValueMultiLine is entered.
func (s *BasePegiParserListener) EnterMessageValueMultiLine(ctx *MessageValueMultiLineContext) {}

// ExitMessageValueMultiLine is called when production messageValueMultiLine is exited.
func (s *BasePegiParserListener) ExitMessageValueMultiLine(ctx *MessageValueMultiLineContext) {}

// EnterLineKeyedElement is called when production lineKeyedElement is entered.
func (s *BasePegiParserListener) EnterLineKeyedElement(ctx *LineKeyedElementContext) {}

// ExitLineKeyedElement is called when production lineKeyedElement is exited.
func (s *BasePegiParserListener) ExitLineKeyedElement(ctx *LineKeyedElementContext) {}

// EnterKey is called when production key is entered.
func (s *BasePegiParserListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasePegiParserListener) ExitKey(ctx *KeyContext) {}

// EnterElement is called when production element is entered.
func (s *BasePegiParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasePegiParserListener) ExitElement(ctx *ElementContext) {}

// EnterSliceType is called when production sliceType is entered.
func (s *BasePegiParserListener) EnterSliceType(ctx *SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *BasePegiParserListener) ExitSliceType(ctx *SliceTypeContext) {}

// EnterTypeT is called when production typeT is entered.
func (s *BasePegiParserListener) EnterTypeT(ctx *TypeTContext) {}

// ExitTypeT is called when production typeT is exited.
func (s *BasePegiParserListener) ExitTypeT(ctx *TypeTContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BasePegiParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BasePegiParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *BasePegiParserListener) EnterTypeLit(ctx *TypeLitContext) {}

// ExitTypeLit is called when production typeLit is exited.
func (s *BasePegiParserListener) ExitTypeLit(ctx *TypeLitContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BasePegiParserListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BasePegiParserListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BasePegiParserListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BasePegiParserListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasePegiParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasePegiParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterEos is called when production eos is entered.
func (s *BasePegiParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BasePegiParserListener) ExitEos(ctx *EosContext) {}
