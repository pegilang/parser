// Generated from ./PegiParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link PegiParser}.
 */
public interface PegiParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link PegiParser#sourceFile}.
	 * @param ctx the parse tree
	 */
	void enterSourceFile(PegiParser.SourceFileContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#sourceFile}.
	 * @param ctx the parse tree
	 */
	void exitSourceFile(PegiParser.SourceFileContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#parserClause}.
	 * @param ctx the parse tree
	 */
	void enterParserClause(PegiParser.ParserClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#parserClause}.
	 * @param ctx the parse tree
	 */
	void exitParserClause(PegiParser.ParserClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#importDecl}.
	 * @param ctx the parse tree
	 */
	void enterImportDecl(PegiParser.ImportDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#importDecl}.
	 * @param ctx the parse tree
	 */
	void exitImportDecl(PegiParser.ImportDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#importSpec}.
	 * @param ctx the parse tree
	 */
	void enterImportSpec(PegiParser.ImportSpecContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#importSpec}.
	 * @param ctx the parse tree
	 */
	void exitImportSpec(PegiParser.ImportSpecContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#importPath}.
	 * @param ctx the parse tree
	 */
	void enterImportPath(PegiParser.ImportPathContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#importPath}.
	 * @param ctx the parse tree
	 */
	void exitImportPath(PegiParser.ImportPathContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#string_}.
	 * @param ctx the parse tree
	 */
	void enterString_(PegiParser.String_Context ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#string_}.
	 * @param ctx the parse tree
	 */
	void exitString_(PegiParser.String_Context ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#statementList}.
	 * @param ctx the parse tree
	 */
	void enterStatementList(PegiParser.StatementListContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#statementList}.
	 * @param ctx the parse tree
	 */
	void exitStatementList(PegiParser.StatementListContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(PegiParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(PegiParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#shortVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterShortVarDecl(PegiParser.ShortVarDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#shortVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitShortVarDecl(PegiParser.ShortVarDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#expressionList}.
	 * @param ctx the parse tree
	 */
	void enterExpressionList(PegiParser.ExpressionListContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#expressionList}.
	 * @param ctx the parse tree
	 */
	void exitExpressionList(PegiParser.ExpressionListContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#identifierList}.
	 * @param ctx the parse tree
	 */
	void enterIdentifierList(PegiParser.IdentifierListContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#identifierList}.
	 * @param ctx the parse tree
	 */
	void exitIdentifierList(PegiParser.IdentifierListContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#pipe}.
	 * @param ctx the parse tree
	 */
	void enterPipe(PegiParser.PipeContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#pipe}.
	 * @param ctx the parse tree
	 */
	void exitPipe(PegiParser.PipeContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#pipeUnit}.
	 * @param ctx the parse tree
	 */
	void enterPipeUnit(PegiParser.PipeUnitContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#pipeUnit}.
	 * @param ctx the parse tree
	 */
	void exitPipeUnit(PegiParser.PipeUnitContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#pipeUnitIdentifers}.
	 * @param ctx the parse tree
	 */
	void enterPipeUnitIdentifers(PegiParser.PipeUnitIdentifersContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#pipeUnitIdentifers}.
	 * @param ctx the parse tree
	 */
	void exitPipeUnitIdentifers(PegiParser.PipeUnitIdentifersContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#pipeUnitExpression}.
	 * @param ctx the parse tree
	 */
	void enterPipeUnitExpression(PegiParser.PipeUnitExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#pipeUnitExpression}.
	 * @param ctx the parse tree
	 */
	void exitPipeUnitExpression(PegiParser.PipeUnitExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(PegiParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(PegiParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#primaryExpr}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpr(PegiParser.PrimaryExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#primaryExpr}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpr(PegiParser.PrimaryExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterOperand(PegiParser.OperandContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitOperand(PegiParser.OperandContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(PegiParser.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(PegiParser.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(PegiParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(PegiParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#arguments}.
	 * @param ctx the parse tree
	 */
	void enterArguments(PegiParser.ArgumentsContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#arguments}.
	 * @param ctx the parse tree
	 */
	void exitArguments(PegiParser.ArgumentsContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#basicLit}.
	 * @param ctx the parse tree
	 */
	void enterBasicLit(PegiParser.BasicLitContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#basicLit}.
	 * @param ctx the parse tree
	 */
	void exitBasicLit(PegiParser.BasicLitContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#operandName}.
	 * @param ctx the parse tree
	 */
	void enterOperandName(PegiParser.OperandNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#operandName}.
	 * @param ctx the parse tree
	 */
	void exitOperandName(PegiParser.OperandNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#unaryExpr}.
	 * @param ctx the parse tree
	 */
	void enterUnaryExpr(PegiParser.UnaryExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#unaryExpr}.
	 * @param ctx the parse tree
	 */
	void exitUnaryExpr(PegiParser.UnaryExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#sliceValue}.
	 * @param ctx the parse tree
	 */
	void enterSliceValue(PegiParser.SliceValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#sliceValue}.
	 * @param ctx the parse tree
	 */
	void exitSliceValue(PegiParser.SliceValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#sliceValueSingleLine}.
	 * @param ctx the parse tree
	 */
	void enterSliceValueSingleLine(PegiParser.SliceValueSingleLineContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#sliceValueSingleLine}.
	 * @param ctx the parse tree
	 */
	void exitSliceValueSingleLine(PegiParser.SliceValueSingleLineContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#sliceValueMultiLine}.
	 * @param ctx the parse tree
	 */
	void enterSliceValueMultiLine(PegiParser.SliceValueMultiLineContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#sliceValueMultiLine}.
	 * @param ctx the parse tree
	 */
	void exitSliceValueMultiLine(PegiParser.SliceValueMultiLineContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#sliceValueMultiLineLine}.
	 * @param ctx the parse tree
	 */
	void enterSliceValueMultiLineLine(PegiParser.SliceValueMultiLineLineContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#sliceValueMultiLineLine}.
	 * @param ctx the parse tree
	 */
	void exitSliceValueMultiLineLine(PegiParser.SliceValueMultiLineLineContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#messageValue}.
	 * @param ctx the parse tree
	 */
	void enterMessageValue(PegiParser.MessageValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#messageValue}.
	 * @param ctx the parse tree
	 */
	void exitMessageValue(PegiParser.MessageValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#keyedElementKV}.
	 * @param ctx the parse tree
	 */
	void enterKeyedElementKV(PegiParser.KeyedElementKVContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#keyedElementKV}.
	 * @param ctx the parse tree
	 */
	void exitKeyedElementKV(PegiParser.KeyedElementKVContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#keyedElementIdentifer}.
	 * @param ctx the parse tree
	 */
	void enterKeyedElementIdentifer(PegiParser.KeyedElementIdentiferContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#keyedElementIdentifer}.
	 * @param ctx the parse tree
	 */
	void exitKeyedElementIdentifer(PegiParser.KeyedElementIdentiferContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#keyedElement}.
	 * @param ctx the parse tree
	 */
	void enterKeyedElement(PegiParser.KeyedElementContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#keyedElement}.
	 * @param ctx the parse tree
	 */
	void exitKeyedElement(PegiParser.KeyedElementContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#messageValueSingleLine}.
	 * @param ctx the parse tree
	 */
	void enterMessageValueSingleLine(PegiParser.MessageValueSingleLineContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#messageValueSingleLine}.
	 * @param ctx the parse tree
	 */
	void exitMessageValueSingleLine(PegiParser.MessageValueSingleLineContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#messageValueMultiLine}.
	 * @param ctx the parse tree
	 */
	void enterMessageValueMultiLine(PegiParser.MessageValueMultiLineContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#messageValueMultiLine}.
	 * @param ctx the parse tree
	 */
	void exitMessageValueMultiLine(PegiParser.MessageValueMultiLineContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#lineKeyedElement}.
	 * @param ctx the parse tree
	 */
	void enterLineKeyedElement(PegiParser.LineKeyedElementContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#lineKeyedElement}.
	 * @param ctx the parse tree
	 */
	void exitLineKeyedElement(PegiParser.LineKeyedElementContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#key}.
	 * @param ctx the parse tree
	 */
	void enterKey(PegiParser.KeyContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#key}.
	 * @param ctx the parse tree
	 */
	void exitKey(PegiParser.KeyContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#element}.
	 * @param ctx the parse tree
	 */
	void enterElement(PegiParser.ElementContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#element}.
	 * @param ctx the parse tree
	 */
	void exitElement(PegiParser.ElementContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#sliceType}.
	 * @param ctx the parse tree
	 */
	void enterSliceType(PegiParser.SliceTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#sliceType}.
	 * @param ctx the parse tree
	 */
	void exitSliceType(PegiParser.SliceTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#typeT}.
	 * @param ctx the parse tree
	 */
	void enterTypeT(PegiParser.TypeTContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#typeT}.
	 * @param ctx the parse tree
	 */
	void exitTypeT(PegiParser.TypeTContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#typeName}.
	 * @param ctx the parse tree
	 */
	void enterTypeName(PegiParser.TypeNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#typeName}.
	 * @param ctx the parse tree
	 */
	void exitTypeName(PegiParser.TypeNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#typeLit}.
	 * @param ctx the parse tree
	 */
	void enterTypeLit(PegiParser.TypeLitContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#typeLit}.
	 * @param ctx the parse tree
	 */
	void exitTypeLit(PegiParser.TypeLitContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#interfaceType}.
	 * @param ctx the parse tree
	 */
	void enterInterfaceType(PegiParser.InterfaceTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#interfaceType}.
	 * @param ctx the parse tree
	 */
	void exitInterfaceType(PegiParser.InterfaceTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#fieldDecl}.
	 * @param ctx the parse tree
	 */
	void enterFieldDecl(PegiParser.FieldDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#fieldDecl}.
	 * @param ctx the parse tree
	 */
	void exitFieldDecl(PegiParser.FieldDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#integer}.
	 * @param ctx the parse tree
	 */
	void enterInteger(PegiParser.IntegerContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#integer}.
	 * @param ctx the parse tree
	 */
	void exitInteger(PegiParser.IntegerContext ctx);
	/**
	 * Enter a parse tree produced by {@link PegiParser#eos}.
	 * @param ctx the parse tree
	 */
	void enterEos(PegiParser.EosContext ctx);
	/**
	 * Exit a parse tree produced by {@link PegiParser#eos}.
	 * @param ctx the parse tree
	 */
	void exitEos(PegiParser.EosContext ctx);
}