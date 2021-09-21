// Generated from ./PegiParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link PegiParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface PegiParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link PegiParser#sourceFile}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSourceFile(PegiParser.SourceFileContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#parserClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParserClause(PegiParser.ParserClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#importDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportDecl(PegiParser.ImportDeclContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#importSpec}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportSpec(PegiParser.ImportSpecContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#importPath}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImportPath(PegiParser.ImportPathContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#string_}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitString_(PegiParser.String_Context ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#statementList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatementList(PegiParser.StatementListContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatement(PegiParser.StatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#shortVarDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitShortVarDecl(PegiParser.ShortVarDeclContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#expressionList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionList(PegiParser.ExpressionListContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#identifierList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdentifierList(PegiParser.IdentifierListContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#pipe}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPipe(PegiParser.PipeContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#pipeUnit}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPipeUnit(PegiParser.PipeUnitContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#pipeUnitIdentifers}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPipeUnitIdentifers(PegiParser.PipeUnitIdentifersContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#pipeUnitExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPipeUnitExpression(PegiParser.PipeUnitExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression(PegiParser.ExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#primaryExpr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimaryExpr(PegiParser.PrimaryExprContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#operand}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperand(PegiParser.OperandContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLiteral(PegiParser.LiteralContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#functionCall}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionCall(PegiParser.FunctionCallContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#arguments}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArguments(PegiParser.ArgumentsContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#basicLit}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBasicLit(PegiParser.BasicLitContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#operandName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperandName(PegiParser.OperandNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#unaryExpr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnaryExpr(PegiParser.UnaryExprContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#sliceValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSliceValue(PegiParser.SliceValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#sliceValueSingleLine}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSliceValueSingleLine(PegiParser.SliceValueSingleLineContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#sliceValueMultiLine}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSliceValueMultiLine(PegiParser.SliceValueMultiLineContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#sliceValueMultiLineLine}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSliceValueMultiLineLine(PegiParser.SliceValueMultiLineLineContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#messageValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMessageValue(PegiParser.MessageValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#keyedElementKV}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeyedElementKV(PegiParser.KeyedElementKVContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#keyedElementIdentifer}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeyedElementIdentifer(PegiParser.KeyedElementIdentiferContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#keyedElement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKeyedElement(PegiParser.KeyedElementContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#messageValueSingleLine}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMessageValueSingleLine(PegiParser.MessageValueSingleLineContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#messageValueMultiLine}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMessageValueMultiLine(PegiParser.MessageValueMultiLineContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#lineKeyedElement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLineKeyedElement(PegiParser.LineKeyedElementContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#key}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitKey(PegiParser.KeyContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#element}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitElement(PegiParser.ElementContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#sliceType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSliceType(PegiParser.SliceTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#typeT}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeT(PegiParser.TypeTContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#typeName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeName(PegiParser.TypeNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#typeLit}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeLit(PegiParser.TypeLitContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#interfaceType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterfaceType(PegiParser.InterfaceTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#fieldDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFieldDecl(PegiParser.FieldDeclContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#integer}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInteger(PegiParser.IntegerContext ctx);
	/**
	 * Visit a parse tree produced by {@link PegiParser#eos}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEos(PegiParser.EosContext ctx);
}