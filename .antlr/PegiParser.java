// Generated from /home/tema/pegilang/parser/PegiParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class PegiParser extends PegiParserBase {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INTERFACE=1, STRUCT=2, PARSER=3, IMPORT=4, IDENTIFIER=5, L_PAREN=6, R_PAREN=7, 
		L_CURLY=8, R_CURLY=9, L_BRACKET=10, R_BRACKET=11, ASSIGN=12, COMMA=13, 
		SEMI=14, COLON=15, DOT=16, PLUS_PLUS=17, MINUS_MINUS=18, DECLARE_ASSIGN=19, 
		ELLIPSIS=20, LOGICAL_OR=21, LOGICAL_AND=22, EQUALS=23, NOT_EQUALS=24, 
		LESS=25, LESS_OR_EQUALS=26, GREATER=27, GREATER_OR_EQUALS=28, OR=29, DIV=30, 
		EXCLAMATION=31, PLUS=32, UNARY_MINUS=33, MINUS=34, MINUS_IN_NAME=35, STAR=36, 
		DECIMAL_LIT=37, BINARY_LIT=38, OCTAL_LIT=39, HEX_LIT=40, FLOAT_LIT=41, 
		DECIMAL_FLOAT_LIT=42, HEX_FLOAT_LIT=43, RUNE_LIT=44, BYTE_VALUE=45, OCTAL_BYTE_VALUE=46, 
		HEX_BYTE_VALUE=47, LITTLE_U_VALUE=48, BIG_U_VALUE=49, RAW_STRING_LIT=50, 
		INTERPRETED_STRING_LIT=51, WS=52, COMMENT=53, TERMINATOR=54, LINE_COMMENT=55;
	public static final int
		RULE_sourceFile = 0, RULE_parserClause = 1, RULE_importDecl = 2, RULE_importSpec = 3, 
		RULE_importPath = 4, RULE_string_ = 5, RULE_statementList = 6, RULE_statement = 7, 
		RULE_shortVarDecl = 8, RULE_expressionList = 9, RULE_identifierList = 10, 
		RULE_pipe = 11, RULE_pipeUnit = 12, RULE_pipeUnitIdentifers = 13, RULE_pipeUnitExpression = 14, 
		RULE_expression = 15, RULE_primaryExpr = 16, RULE_operand = 17, RULE_literal = 18, 
		RULE_functionCall = 19, RULE_arguments = 20, RULE_basicLit = 21, RULE_operandName = 22, 
		RULE_unaryExpr = 23, RULE_sliceValue = 24, RULE_sliceValueSingleLine = 25, 
		RULE_sliceValueMultiLine = 26, RULE_sliceValueMultiLineLine = 27, RULE_messageValue = 28, 
		RULE_keyedElementKV = 29, RULE_keyedElementIdentifer = 30, RULE_keyedElement = 31, 
		RULE_messageValueSingleLine = 32, RULE_messageValueMultiLine = 33, RULE_lineKeyedElement = 34, 
		RULE_key = 35, RULE_element = 36, RULE_sliceType = 37, RULE_typeT = 38, 
		RULE_typeName = 39, RULE_typeLit = 40, RULE_interfaceType = 41, RULE_fieldDecl = 42, 
		RULE_integer = 43, RULE_eos = 44;
	private static String[] makeRuleNames() {
		return new String[] {
			"sourceFile", "parserClause", "importDecl", "importSpec", "importPath", 
			"string_", "statementList", "statement", "shortVarDecl", "expressionList", 
			"identifierList", "pipe", "pipeUnit", "pipeUnitIdentifers", "pipeUnitExpression", 
			"expression", "primaryExpr", "operand", "literal", "functionCall", "arguments", 
			"basicLit", "operandName", "unaryExpr", "sliceValue", "sliceValueSingleLine", 
			"sliceValueMultiLine", "sliceValueMultiLineLine", "messageValue", "keyedElementKV", 
			"keyedElementIdentifer", "keyedElement", "messageValueSingleLine", "messageValueMultiLine", 
			"lineKeyedElement", "key", "element", "sliceType", "typeT", "typeName", 
			"typeLit", "interfaceType", "fieldDecl", "integer", "eos"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'interface'", "'struct'", "'parser'", "'import'", null, "'('", 
			"')'", "'{'", "'}'", "'['", "']'", "'='", "','", "';'", "':'", "'.'", 
			"'++'", "'--'", "':='", "'...'", "'||'", "'&&'", "'=='", "'!='", "'<'", 
			"'<='", "'>'", "'>='", "'|'", "' / '", "'!'", "' + '", "' -'", "' - '", 
			"'-'", "' * '"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INTERFACE", "STRUCT", "PARSER", "IMPORT", "IDENTIFIER", "L_PAREN", 
			"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", 
			"COMMA", "SEMI", "COLON", "DOT", "PLUS_PLUS", "MINUS_MINUS", "DECLARE_ASSIGN", 
			"ELLIPSIS", "LOGICAL_OR", "LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS", 
			"LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "OR", "DIV", "EXCLAMATION", 
			"PLUS", "UNARY_MINUS", "MINUS", "MINUS_IN_NAME", "STAR", "DECIMAL_LIT", 
			"BINARY_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT", 
			"HEX_FLOAT_LIT", "RUNE_LIT", "BYTE_VALUE", "OCTAL_BYTE_VALUE", "HEX_BYTE_VALUE", 
			"LITTLE_U_VALUE", "BIG_U_VALUE", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT", 
			"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "PegiParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public PegiParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SourceFileContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(PegiParser.EOF, 0); }
		public ParserClauseContext parserClause() {
			return getRuleContext(ParserClauseContext.class,0);
		}
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public List<ImportDeclContext> importDecl() {
			return getRuleContexts(ImportDeclContext.class);
		}
		public ImportDeclContext importDecl(int i) {
			return getRuleContext(ImportDeclContext.class,i);
		}
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public SourceFileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sourceFile; }
	}

	public final SourceFileContext sourceFile() throws RecognitionException {
		SourceFileContext _localctx = new SourceFileContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_sourceFile);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PARSER) {
				{
				setState(90);
				parserClause();
				setState(91);
				eos();
				}
			}

			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IMPORT) {
				{
				{
				setState(95);
				importDecl();
				setState(96);
				eos();
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(104);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERFACE || _la==IDENTIFIER) {
				{
				setState(103);
				statementList();
				}
			}

			setState(106);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParserClauseContext extends ParserRuleContext {
		public String_Context parserPath;
		public TerminalNode PARSER() { return getToken(PegiParser.PARSER, 0); }
		public String_Context string_() {
			return getRuleContext(String_Context.class,0);
		}
		public ParserClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parserClause; }
	}

	public final ParserClauseContext parserClause() throws RecognitionException {
		ParserClauseContext _localctx = new ParserClauseContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_parserClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(PARSER);
			setState(109);
			((ParserClauseContext)_localctx).parserPath = string_();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportDeclContext extends ParserRuleContext {
		public TerminalNode IMPORT() { return getToken(PegiParser.IMPORT, 0); }
		public List<ImportSpecContext> importSpec() {
			return getRuleContexts(ImportSpecContext.class);
		}
		public ImportSpecContext importSpec(int i) {
			return getRuleContext(ImportSpecContext.class,i);
		}
		public TerminalNode L_PAREN() { return getToken(PegiParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(PegiParser.R_PAREN, 0); }
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public ImportDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importDecl; }
	}

	public final ImportDeclContext importDecl() throws RecognitionException {
		ImportDeclContext _localctx = new ImportDeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_importDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			match(IMPORT);
			setState(123);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
			case DOT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
				{
				setState(112);
				importSpec();
				}
				break;
			case L_PAREN:
				{
				setState(113);
				match(L_PAREN);
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IDENTIFIER) | (1L << DOT) | (1L << RAW_STRING_LIT) | (1L << INTERPRETED_STRING_LIT))) != 0)) {
					{
					{
					setState(114);
					importSpec();
					setState(115);
					eos();
					}
					}
					setState(121);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(122);
				match(R_PAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportSpecContext extends ParserRuleContext {
		public Token alias;
		public ImportPathContext importPath() {
			return getRuleContext(ImportPathContext.class,0);
		}
		public TerminalNode DOT() { return getToken(PegiParser.DOT, 0); }
		public TerminalNode IDENTIFIER() { return getToken(PegiParser.IDENTIFIER, 0); }
		public ImportSpecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importSpec; }
	}

	public final ImportSpecContext importSpec() throws RecognitionException {
		ImportSpecContext _localctx = new ImportSpecContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_importSpec);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER || _la==DOT) {
				{
				setState(125);
				((ImportSpecContext)_localctx).alias = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==IDENTIFIER || _la==DOT) ) {
					((ImportSpecContext)_localctx).alias = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(128);
			importPath();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportPathContext extends ParserRuleContext {
		public String_Context string_() {
			return getRuleContext(String_Context.class,0);
		}
		public ImportPathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importPath; }
	}

	public final ImportPathContext importPath() throws RecognitionException {
		ImportPathContext _localctx = new ImportPathContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_importPath);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			string_();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class String_Context extends ParserRuleContext {
		public TerminalNode RAW_STRING_LIT() { return getToken(PegiParser.RAW_STRING_LIT, 0); }
		public TerminalNode INTERPRETED_STRING_LIT() { return getToken(PegiParser.INTERPRETED_STRING_LIT, 0); }
		public String_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string_; }
	}

	public final String_Context string_() throws RecognitionException {
		String_Context _localctx = new String_Context(_ctx, getState());
		enterRule(_localctx, 10, RULE_string_);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			_la = _input.LA(1);
			if ( !(_la==RAW_STRING_LIT || _la==INTERPRETED_STRING_LIT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementListContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
	}

	public final StatementListContext statementList() throws RecognitionException {
		StatementListContext _localctx = new StatementListContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_statementList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(134);
				statement();
				setState(135);
				eos();
				}
				}
				setState(139); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==INTERFACE || _la==IDENTIFIER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public ShortVarDeclContext shortVarDecl() {
			return getRuleContext(ShortVarDeclContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public InterfaceTypeContext interfaceType() {
			return getRuleContext(InterfaceTypeContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(141);
				shortVarDecl();
				}
				break;
			case 2:
				{
				setState(142);
				functionCall();
				}
				break;
			case 3:
				{
				setState(143);
				interfaceType();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ShortVarDeclContext extends ParserRuleContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public TerminalNode DECLARE_ASSIGN() { return getToken(PegiParser.DECLARE_ASSIGN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public PipeContext pipe() {
			return getRuleContext(PipeContext.class,0);
		}
		public ShortVarDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_shortVarDecl; }
	}

	public final ShortVarDeclContext shortVarDecl() throws RecognitionException {
		ShortVarDeclContext _localctx = new ShortVarDeclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_shortVarDecl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			identifierList();
			setState(147);
			match(DECLARE_ASSIGN);
			setState(150);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(148);
				expressionList();
				}
				break;
			case 2:
				{
				setState(149);
				pipe();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionListContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(PegiParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(PegiParser.COMMA, i);
		}
		public ExpressionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionList; }
	}

	public final ExpressionListContext expressionList() throws RecognitionException {
		ExpressionListContext _localctx = new ExpressionListContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_expressionList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			expression(0);
			setState(157);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(153);
					match(COMMA);
					setState(154);
					expression(0);
					}
					} 
				}
				setState(159);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IdentifierListContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(PegiParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(PegiParser.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(PegiParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(PegiParser.COMMA, i);
		}
		public IdentifierListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierList; }
	}

	public final IdentifierListContext identifierList() throws RecognitionException {
		IdentifierListContext _localctx = new IdentifierListContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_identifierList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(IDENTIFIER);
			setState(165);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(161);
					match(COMMA);
					setState(162);
					match(IDENTIFIER);
					}
					} 
				}
				setState(167);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipeContext extends ParserRuleContext {
		public List<PipeUnitContext> pipeUnit() {
			return getRuleContexts(PipeUnitContext.class);
		}
		public PipeUnitContext pipeUnit(int i) {
			return getRuleContext(PipeUnitContext.class,i);
		}
		public PipeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipe; }
	}

	public final PipeContext pipe() throws RecognitionException {
		PipeContext _localctx = new PipeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_pipe);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(168);
			pipeUnit();
			setState(170); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(169);
					pipeUnit();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(172); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipeUnitContext extends ParserRuleContext {
		public PipeUnitIdentifersContext pipeUnitIdentifers() {
			return getRuleContext(PipeUnitIdentifersContext.class,0);
		}
		public PipeUnitExpressionContext pipeUnitExpression() {
			return getRuleContext(PipeUnitExpressionContext.class,0);
		}
		public PipeUnitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipeUnit; }
	}

	public final PipeUnitContext pipeUnit() throws RecognitionException {
		PipeUnitContext _localctx = new PipeUnitContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_pipeUnit);
		try {
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(174);
				pipeUnitIdentifers();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(175);
				pipeUnitExpression();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipeUnitIdentifersContext extends ParserRuleContext {
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public TerminalNode OR() { return getToken(PegiParser.OR, 0); }
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public PipeUnitIdentifersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipeUnitIdentifers; }
	}

	public final PipeUnitIdentifersContext pipeUnitIdentifers() throws RecognitionException {
		PipeUnitIdentifersContext _localctx = new PipeUnitIdentifersContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_pipeUnitIdentifers);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			eos();
			setState(179);
			match(OR);
			setState(180);
			identifierList();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PipeUnitExpressionContext extends ParserRuleContext {
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public TerminalNode OR() { return getToken(PegiParser.OR, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PipeUnitExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipeUnitExpression; }
	}

	public final PipeUnitExpressionContext pipeUnitExpression() throws RecognitionException {
		PipeUnitExpressionContext _localctx = new PipeUnitExpressionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_pipeUnitExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			eos();
			setState(183);
			match(OR);
			setState(184);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public Token mul_op;
		public Token add_op;
		public Token rel_op;
		public PrimaryExprContext primaryExpr() {
			return getRuleContext(PrimaryExprContext.class,0);
		}
		public UnaryExprContext unaryExpr() {
			return getRuleContext(UnaryExprContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode STAR() { return getToken(PegiParser.STAR, 0); }
		public TerminalNode DIV() { return getToken(PegiParser.DIV, 0); }
		public TerminalNode PLUS() { return getToken(PegiParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(PegiParser.MINUS, 0); }
		public TerminalNode EQUALS() { return getToken(PegiParser.EQUALS, 0); }
		public TerminalNode NOT_EQUALS() { return getToken(PegiParser.NOT_EQUALS, 0); }
		public TerminalNode LESS() { return getToken(PegiParser.LESS, 0); }
		public TerminalNode LESS_OR_EQUALS() { return getToken(PegiParser.LESS_OR_EQUALS, 0); }
		public TerminalNode GREATER() { return getToken(PegiParser.GREATER, 0); }
		public TerminalNode GREATER_OR_EQUALS() { return getToken(PegiParser.GREATER_OR_EQUALS, 0); }
		public TerminalNode LOGICAL_AND() { return getToken(PegiParser.LOGICAL_AND, 0); }
		public TerminalNode LOGICAL_OR() { return getToken(PegiParser.LOGICAL_OR, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(187);
				primaryExpr();
				}
				break;
			case 2:
				{
				setState(188);
				unaryExpr();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(208);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(206);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(191);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(192);
						((ExpressionContext)_localctx).mul_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIV || _la==STAR) ) {
							((ExpressionContext)_localctx).mul_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(193);
						expression(6);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(194);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(195);
						((ExpressionContext)_localctx).add_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((ExpressionContext)_localctx).add_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(196);
						expression(5);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(197);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(198);
						((ExpressionContext)_localctx).rel_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << EQUALS) | (1L << NOT_EQUALS) | (1L << LESS) | (1L << LESS_OR_EQUALS) | (1L << GREATER) | (1L << GREATER_OR_EQUALS))) != 0)) ) {
							((ExpressionContext)_localctx).rel_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(199);
						expression(4);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(200);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(201);
						match(LOGICAL_AND);
						setState(202);
						expression(3);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(203);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(204);
						match(LOGICAL_OR);
						setState(205);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(210);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrimaryExprContext extends ParserRuleContext {
		public OperandContext operand() {
			return getRuleContext(OperandContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public PrimaryExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpr; }
	}

	public final PrimaryExprContext primaryExpr() throws RecognitionException {
		PrimaryExprContext _localctx = new PrimaryExprContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_primaryExpr);
		try {
			setState(213);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(211);
				operand();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(212);
				functionCall();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OperandContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public OperandNameContext operandName() {
			return getRuleContext(OperandNameContext.class,0);
		}
		public TerminalNode L_PAREN() { return getToken(PegiParser.L_PAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(PegiParser.R_PAREN, 0); }
		public OperandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operand; }
	}

	public final OperandContext operand() throws RecognitionException {
		OperandContext _localctx = new OperandContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_operand);
		try {
			setState(221);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case L_CURLY:
			case L_BRACKET:
			case DECIMAL_LIT:
			case BINARY_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
			case FLOAT_LIT:
			case RUNE_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
				enterOuterAlt(_localctx, 1);
				{
				setState(215);
				literal();
				}
				break;
			case IDENTIFIER:
				enterOuterAlt(_localctx, 2);
				{
				setState(216);
				operandName();
				}
				break;
			case L_PAREN:
				enterOuterAlt(_localctx, 3);
				{
				setState(217);
				match(L_PAREN);
				setState(218);
				expression(0);
				setState(219);
				match(R_PAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LiteralContext extends ParserRuleContext {
		public BasicLitContext basicLit() {
			return getRuleContext(BasicLitContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_literal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			basicLit();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionCallContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_functionCall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			typeName();
			setState(226);
			arguments();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode L_PAREN() { return getToken(PegiParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(PegiParser.R_PAREN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public TypeTContext typeT() {
			return getRuleContext(TypeTContext.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(PegiParser.ELLIPSIS, 0); }
		public List<TerminalNode> COMMA() { return getTokens(PegiParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(PegiParser.COMMA, i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(L_PAREN);
			setState(243);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INTERFACE) | (1L << IDENTIFIER) | (1L << L_PAREN) | (1L << L_CURLY) | (1L << L_BRACKET) | (1L << EXCLAMATION) | (1L << UNARY_MINUS) | (1L << DECIMAL_LIT) | (1L << BINARY_LIT) | (1L << OCTAL_LIT) | (1L << HEX_LIT) | (1L << FLOAT_LIT) | (1L << RUNE_LIT) | (1L << RAW_STRING_LIT) | (1L << INTERPRETED_STRING_LIT))) != 0)) {
				{
				setState(235);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
				case 1:
					{
					setState(229);
					expressionList();
					}
					break;
				case 2:
					{
					setState(230);
					typeT();
					setState(233);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
					case 1:
						{
						setState(231);
						match(COMMA);
						setState(232);
						expressionList();
						}
						break;
					}
					}
					break;
				}
				setState(238);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ELLIPSIS) {
					{
					setState(237);
					match(ELLIPSIS);
					}
				}

				setState(241);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(240);
					match(COMMA);
					}
				}

				}
			}

			setState(245);
			match(R_PAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BasicLitContext extends ParserRuleContext {
		public IntegerContext integer() {
			return getRuleContext(IntegerContext.class,0);
		}
		public MessageValueContext messageValue() {
			return getRuleContext(MessageValueContext.class,0);
		}
		public SliceValueContext sliceValue() {
			return getRuleContext(SliceValueContext.class,0);
		}
		public String_Context string_() {
			return getRuleContext(String_Context.class,0);
		}
		public TerminalNode FLOAT_LIT() { return getToken(PegiParser.FLOAT_LIT, 0); }
		public TerminalNode RUNE_LIT() { return getToken(PegiParser.RUNE_LIT, 0); }
		public BasicLitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_basicLit; }
	}

	public final BasicLitContext basicLit() throws RecognitionException {
		BasicLitContext _localctx = new BasicLitContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_basicLit);
		try {
			setState(253);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(247);
				integer();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(248);
				messageValue();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(249);
				sliceValue();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(250);
				string_();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(251);
				match(FLOAT_LIT);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(252);
				match(RUNE_LIT);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OperandNameContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(PegiParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(PegiParser.IDENTIFIER, i);
		}
		public TerminalNode DOT() { return getToken(PegiParser.DOT, 0); }
		public OperandNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operandName; }
	}

	public final OperandNameContext operandName() throws RecognitionException {
		OperandNameContext _localctx = new OperandNameContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_operandName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(255);
			match(IDENTIFIER);
			setState(258);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				{
				setState(256);
				match(DOT);
				setState(257);
				match(IDENTIFIER);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UnaryExprContext extends ParserRuleContext {
		public Token unary_op;
		public PrimaryExprContext primaryExpr() {
			return getRuleContext(PrimaryExprContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode UNARY_MINUS() { return getToken(PegiParser.UNARY_MINUS, 0); }
		public TerminalNode EXCLAMATION() { return getToken(PegiParser.EXCLAMATION, 0); }
		public UnaryExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryExpr; }
	}

	public final UnaryExprContext unaryExpr() throws RecognitionException {
		UnaryExprContext _localctx = new UnaryExprContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_unaryExpr);
		int _la;
		try {
			setState(263);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
			case L_PAREN:
			case L_CURLY:
			case L_BRACKET:
			case DECIMAL_LIT:
			case BINARY_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
			case FLOAT_LIT:
			case RUNE_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
				enterOuterAlt(_localctx, 1);
				{
				setState(260);
				primaryExpr();
				}
				break;
			case EXCLAMATION:
			case UNARY_MINUS:
				enterOuterAlt(_localctx, 2);
				{
				setState(261);
				((UnaryExprContext)_localctx).unary_op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==EXCLAMATION || _la==UNARY_MINUS) ) {
					((UnaryExprContext)_localctx).unary_op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(262);
				expression(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SliceValueContext extends ParserRuleContext {
		public SliceValueSingleLineContext sliceValueSingleLine() {
			return getRuleContext(SliceValueSingleLineContext.class,0);
		}
		public SliceValueMultiLineContext sliceValueMultiLine() {
			return getRuleContext(SliceValueMultiLineContext.class,0);
		}
		public SliceValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceValue; }
	}

	public final SliceValueContext sliceValue() throws RecognitionException {
		SliceValueContext _localctx = new SliceValueContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_sliceValue);
		try {
			setState(267);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(265);
				sliceValueSingleLine();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(266);
				sliceValueMultiLine();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SliceValueSingleLineContext extends ParserRuleContext {
		public TerminalNode L_BRACKET() { return getToken(PegiParser.L_BRACKET, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode R_BRACKET() { return getToken(PegiParser.R_BRACKET, 0); }
		public List<TerminalNode> COMMA() { return getTokens(PegiParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(PegiParser.COMMA, i);
		}
		public SliceValueSingleLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceValueSingleLine; }
	}

	public final SliceValueSingleLineContext sliceValueSingleLine() throws RecognitionException {
		SliceValueSingleLineContext _localctx = new SliceValueSingleLineContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_sliceValueSingleLine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(269);
			match(L_BRACKET);
			setState(270);
			expression(0);
			setState(275);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(271);
				match(COMMA);
				setState(272);
				expression(0);
				}
				}
				setState(277);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(278);
			match(R_BRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SliceValueMultiLineContext extends ParserRuleContext {
		public TerminalNode L_BRACKET() { return getToken(PegiParser.L_BRACKET, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public TerminalNode R_BRACKET() { return getToken(PegiParser.R_BRACKET, 0); }
		public List<SliceValueMultiLineLineContext> sliceValueMultiLineLine() {
			return getRuleContexts(SliceValueMultiLineLineContext.class);
		}
		public SliceValueMultiLineLineContext sliceValueMultiLineLine(int i) {
			return getRuleContext(SliceValueMultiLineLineContext.class,i);
		}
		public SliceValueMultiLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceValueMultiLine; }
	}

	public final SliceValueMultiLineContext sliceValueMultiLine() throws RecognitionException {
		SliceValueMultiLineContext _localctx = new SliceValueMultiLineContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_sliceValueMultiLine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(280);
			match(L_BRACKET);
			setState(281);
			eos();
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IDENTIFIER) | (1L << L_PAREN) | (1L << L_CURLY) | (1L << L_BRACKET) | (1L << EXCLAMATION) | (1L << UNARY_MINUS) | (1L << DECIMAL_LIT) | (1L << BINARY_LIT) | (1L << OCTAL_LIT) | (1L << HEX_LIT) | (1L << FLOAT_LIT) | (1L << RUNE_LIT) | (1L << RAW_STRING_LIT) | (1L << INTERPRETED_STRING_LIT))) != 0)) {
				{
				{
				setState(282);
				sliceValueMultiLineLine();
				}
				}
				setState(287);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(288);
			match(R_BRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SliceValueMultiLineLineContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(PegiParser.COMMA, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public SliceValueMultiLineLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceValueMultiLineLine; }
	}

	public final SliceValueMultiLineLineContext sliceValueMultiLineLine() throws RecognitionException {
		SliceValueMultiLineLineContext _localctx = new SliceValueMultiLineLineContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_sliceValueMultiLineLine);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			expression(0);
			setState(291);
			match(COMMA);
			setState(292);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MessageValueContext extends ParserRuleContext {
		public MessageValueSingleLineContext messageValueSingleLine() {
			return getRuleContext(MessageValueSingleLineContext.class,0);
		}
		public MessageValueMultiLineContext messageValueMultiLine() {
			return getRuleContext(MessageValueMultiLineContext.class,0);
		}
		public MessageValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_messageValue; }
	}

	public final MessageValueContext messageValue() throws RecognitionException {
		MessageValueContext _localctx = new MessageValueContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_messageValue);
		try {
			setState(296);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(294);
				messageValueSingleLine();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(295);
				messageValueMultiLine();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyedElementKVContext extends ParserRuleContext {
		public KeyContext key() {
			return getRuleContext(KeyContext.class,0);
		}
		public TerminalNode COLON() { return getToken(PegiParser.COLON, 0); }
		public ElementContext element() {
			return getRuleContext(ElementContext.class,0);
		}
		public KeyedElementKVContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyedElementKV; }
	}

	public final KeyedElementKVContext keyedElementKV() throws RecognitionException {
		KeyedElementKVContext _localctx = new KeyedElementKVContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_keyedElementKV);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			key();
			setState(299);
			match(COLON);
			setState(300);
			element();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyedElementIdentiferContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(PegiParser.IDENTIFIER, 0); }
		public KeyedElementIdentiferContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyedElementIdentifer; }
	}

	public final KeyedElementIdentiferContext keyedElementIdentifer() throws RecognitionException {
		KeyedElementIdentiferContext _localctx = new KeyedElementIdentiferContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_keyedElementIdentifer);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(302);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyedElementContext extends ParserRuleContext {
		public KeyedElementIdentiferContext keyedElementIdentifer() {
			return getRuleContext(KeyedElementIdentiferContext.class,0);
		}
		public KeyedElementKVContext keyedElementKV() {
			return getRuleContext(KeyedElementKVContext.class,0);
		}
		public KeyedElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyedElement; }
	}

	public final KeyedElementContext keyedElement() throws RecognitionException {
		KeyedElementContext _localctx = new KeyedElementContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_keyedElement);
		try {
			setState(306);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(304);
				keyedElementIdentifer();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(305);
				keyedElementKV();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MessageValueSingleLineContext extends ParserRuleContext {
		public TerminalNode L_CURLY() { return getToken(PegiParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(PegiParser.R_CURLY, 0); }
		public List<KeyedElementContext> keyedElement() {
			return getRuleContexts(KeyedElementContext.class);
		}
		public KeyedElementContext keyedElement(int i) {
			return getRuleContext(KeyedElementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(PegiParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(PegiParser.COMMA, i);
		}
		public MessageValueSingleLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_messageValueSingleLine; }
	}

	public final MessageValueSingleLineContext messageValueSingleLine() throws RecognitionException {
		MessageValueSingleLineContext _localctx = new MessageValueSingleLineContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_messageValueSingleLine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(308);
			match(L_CURLY);
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(309);
				keyedElement();
				}
			}

			setState(316);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(312);
				match(COMMA);
				setState(313);
				keyedElement();
				}
				}
				setState(318);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(319);
			match(R_CURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MessageValueMultiLineContext extends ParserRuleContext {
		public TerminalNode L_CURLY() { return getToken(PegiParser.L_CURLY, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public TerminalNode R_CURLY() { return getToken(PegiParser.R_CURLY, 0); }
		public List<LineKeyedElementContext> lineKeyedElement() {
			return getRuleContexts(LineKeyedElementContext.class);
		}
		public LineKeyedElementContext lineKeyedElement(int i) {
			return getRuleContext(LineKeyedElementContext.class,i);
		}
		public MessageValueMultiLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_messageValueMultiLine; }
	}

	public final MessageValueMultiLineContext messageValueMultiLine() throws RecognitionException {
		MessageValueMultiLineContext _localctx = new MessageValueMultiLineContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_messageValueMultiLine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			match(L_CURLY);
			setState(322);
			eos();
			setState(324); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(323);
				lineKeyedElement();
				}
				}
				setState(326); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==IDENTIFIER );
			setState(328);
			match(R_CURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LineKeyedElementContext extends ParserRuleContext {
		public KeyedElementContext keyedElement() {
			return getRuleContext(KeyedElementContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(PegiParser.COMMA, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public LineKeyedElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lineKeyedElement; }
	}

	public final LineKeyedElementContext lineKeyedElement() throws RecognitionException {
		LineKeyedElementContext _localctx = new LineKeyedElementContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_lineKeyedElement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(330);
			keyedElement();
			setState(331);
			match(COMMA);
			setState(332);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(PegiParser.IDENTIFIER, 0); }
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_key);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(334);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ElementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public MessageValueContext messageValue() {
			return getRuleContext(MessageValueContext.class,0);
		}
		public SliceValueContext sliceValue() {
			return getRuleContext(SliceValueContext.class,0);
		}
		public ElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_element; }
	}

	public final ElementContext element() throws RecognitionException {
		ElementContext _localctx = new ElementContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_element);
		try {
			setState(339);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(336);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(337);
				messageValue();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(338);
				sliceValue();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SliceTypeContext extends ParserRuleContext {
		public TerminalNode L_BRACKET() { return getToken(PegiParser.L_BRACKET, 0); }
		public TerminalNode R_BRACKET() { return getToken(PegiParser.R_BRACKET, 0); }
		public TypeTContext typeT() {
			return getRuleContext(TypeTContext.class,0);
		}
		public SliceTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceType; }
	}

	public final SliceTypeContext sliceType() throws RecognitionException {
		SliceTypeContext _localctx = new SliceTypeContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_sliceType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(341);
			match(L_BRACKET);
			setState(342);
			match(R_BRACKET);
			setState(343);
			typeT();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeTContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public TypeLitContext typeLit() {
			return getRuleContext(TypeLitContext.class,0);
		}
		public TypeTContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeT; }
	}

	public final TypeTContext typeT() throws RecognitionException {
		TypeTContext _localctx = new TypeTContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_typeT);
		try {
			setState(347);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(345);
				typeName();
				}
				break;
			case INTERFACE:
				enterOuterAlt(_localctx, 2);
				{
				setState(346);
				typeLit();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeNameContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(PegiParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(PegiParser.IDENTIFIER, i);
		}
		public List<TerminalNode> DOT() { return getTokens(PegiParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(PegiParser.DOT, i);
		}
		public TypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeName; }
	}

	public final TypeNameContext typeName() throws RecognitionException {
		TypeNameContext _localctx = new TypeNameContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_typeName);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			match(IDENTIFIER);
			setState(354);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(350);
					match(DOT);
					setState(351);
					match(IDENTIFIER);
					}
					} 
				}
				setState(356);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeLitContext extends ParserRuleContext {
		public InterfaceTypeContext interfaceType() {
			return getRuleContext(InterfaceTypeContext.class,0);
		}
		public TypeLitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeLit; }
	}

	public final TypeLitContext typeLit() throws RecognitionException {
		TypeLitContext _localctx = new TypeLitContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_typeLit);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(357);
			interfaceType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InterfaceTypeContext extends ParserRuleContext {
		public Token interfaceName;
		public TerminalNode INTERFACE() { return getToken(PegiParser.INTERFACE, 0); }
		public TerminalNode L_CURLY() { return getToken(PegiParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(PegiParser.R_CURLY, 0); }
		public TerminalNode IDENTIFIER() { return getToken(PegiParser.IDENTIFIER, 0); }
		public List<EosContext> eos() {
			return getRuleContexts(EosContext.class);
		}
		public EosContext eos(int i) {
			return getRuleContext(EosContext.class,i);
		}
		public List<FieldDeclContext> fieldDecl() {
			return getRuleContexts(FieldDeclContext.class);
		}
		public FieldDeclContext fieldDecl(int i) {
			return getRuleContext(FieldDeclContext.class,i);
		}
		public List<TypeNameContext> typeName() {
			return getRuleContexts(TypeNameContext.class);
		}
		public TypeNameContext typeName(int i) {
			return getRuleContext(TypeNameContext.class,i);
		}
		public InterfaceTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceType; }
	}

	public final InterfaceTypeContext interfaceType() throws RecognitionException {
		InterfaceTypeContext _localctx = new InterfaceTypeContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_interfaceType);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			match(INTERFACE);
			setState(360);
			((InterfaceTypeContext)_localctx).interfaceName = match(IDENTIFIER);
			setState(361);
			match(L_CURLY);
			setState(370);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,38,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(364);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
					case 1:
						{
						setState(362);
						fieldDecl();
						}
						break;
					case 2:
						{
						setState(363);
						typeName();
						}
						break;
					}
					setState(366);
					eos();
					}
					} 
				}
				setState(372);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,38,_ctx);
			}
			setState(373);
			match(R_CURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FieldDeclContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(PegiParser.IDENTIFIER, 0); }
		public TypeTContext typeT() {
			return getRuleContext(TypeTContext.class,0);
		}
		public FieldDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldDecl; }
	}

	public final FieldDeclContext fieldDecl() throws RecognitionException {
		FieldDeclContext _localctx = new FieldDeclContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_fieldDecl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(375);
			if (!(p.noTerminatorBetween(2))) throw new FailedPredicateException(this, "p.noTerminatorBetween(2)");
			setState(376);
			match(IDENTIFIER);
			setState(377);
			typeT();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IntegerContext extends ParserRuleContext {
		public TerminalNode DECIMAL_LIT() { return getToken(PegiParser.DECIMAL_LIT, 0); }
		public TerminalNode BINARY_LIT() { return getToken(PegiParser.BINARY_LIT, 0); }
		public TerminalNode OCTAL_LIT() { return getToken(PegiParser.OCTAL_LIT, 0); }
		public TerminalNode HEX_LIT() { return getToken(PegiParser.HEX_LIT, 0); }
		public TerminalNode RUNE_LIT() { return getToken(PegiParser.RUNE_LIT, 0); }
		public IntegerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integer; }
	}

	public final IntegerContext integer() throws RecognitionException {
		IntegerContext _localctx = new IntegerContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_integer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DECIMAL_LIT) | (1L << BINARY_LIT) | (1L << OCTAL_LIT) | (1L << HEX_LIT) | (1L << RUNE_LIT))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EosContext extends ParserRuleContext {
		public TerminalNode SEMI() { return getToken(PegiParser.SEMI, 0); }
		public TerminalNode EOF() { return getToken(PegiParser.EOF, 0); }
		public EosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eos; }
	}

	public final EosContext eos() throws RecognitionException {
		EosContext _localctx = new EosContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_eos);
		try {
			setState(385);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(381);
				match(SEMI);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(382);
				match(EOF);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(383);
				if (!(p.lineTerminatorAhead())) throw new FailedPredicateException(this, "p.lineTerminatorAhead()");
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(384);
				if (!(p.checkPreviousTokenText("}"))) throw new FailedPredicateException(this, "p.checkPreviousTokenText(\"}\")");
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 15:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 42:
			return fieldDecl_sempred((FieldDeclContext)_localctx, predIndex);
		case 44:
			return eos_sempred((EosContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		case 1:
			return precpred(_ctx, 4);
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 2);
		case 4:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean fieldDecl_sempred(FieldDeclContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return p.noTerminatorBetween(2);
		}
		return true;
	}
	private boolean eos_sempred(EosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return p.lineTerminatorAhead();
		case 7:
			return p.checkPreviousTokenText("}");
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\39\u0186\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\3\2\3\2\3\2\5\2`\n\2\3\2\3\2\3\2\7\2e\n\2\f\2\16\2h\13"+
		"\2\3\2\5\2k\n\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\7\4x\n\4\f"+
		"\4\16\4{\13\4\3\4\5\4~\n\4\3\5\5\5\u0081\n\5\3\5\3\5\3\6\3\6\3\7\3\7\3"+
		"\b\3\b\3\b\6\b\u008c\n\b\r\b\16\b\u008d\3\t\3\t\3\t\5\t\u0093\n\t\3\n"+
		"\3\n\3\n\3\n\5\n\u0099\n\n\3\13\3\13\3\13\7\13\u009e\n\13\f\13\16\13\u00a1"+
		"\13\13\3\f\3\f\3\f\7\f\u00a6\n\f\f\f\16\f\u00a9\13\f\3\r\3\r\6\r\u00ad"+
		"\n\r\r\r\16\r\u00ae\3\16\3\16\5\16\u00b3\n\16\3\17\3\17\3\17\3\17\3\20"+
		"\3\20\3\20\3\20\3\21\3\21\3\21\5\21\u00c0\n\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\7\21\u00d1\n\21\f\21"+
		"\16\21\u00d4\13\21\3\22\3\22\5\22\u00d8\n\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\5\23\u00e0\n\23\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26"+
		"\5\26\u00ec\n\26\5\26\u00ee\n\26\3\26\5\26\u00f1\n\26\3\26\5\26\u00f4"+
		"\n\26\5\26\u00f6\n\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u0100"+
		"\n\27\3\30\3\30\3\30\5\30\u0105\n\30\3\31\3\31\3\31\5\31\u010a\n\31\3"+
		"\32\3\32\5\32\u010e\n\32\3\33\3\33\3\33\3\33\7\33\u0114\n\33\f\33\16\33"+
		"\u0117\13\33\3\33\3\33\3\34\3\34\3\34\7\34\u011e\n\34\f\34\16\34\u0121"+
		"\13\34\3\34\3\34\3\35\3\35\3\35\3\35\3\36\3\36\5\36\u012b\n\36\3\37\3"+
		"\37\3\37\3\37\3 \3 \3!\3!\5!\u0135\n!\3\"\3\"\5\"\u0139\n\"\3\"\3\"\7"+
		"\"\u013d\n\"\f\"\16\"\u0140\13\"\3\"\3\"\3#\3#\3#\6#\u0147\n#\r#\16#\u0148"+
		"\3#\3#\3$\3$\3$\3$\3%\3%\3&\3&\3&\5&\u0156\n&\3\'\3\'\3\'\3\'\3(\3(\5"+
		"(\u015e\n(\3)\3)\3)\7)\u0163\n)\f)\16)\u0166\13)\3*\3*\3+\3+\3+\3+\3+"+
		"\5+\u016f\n+\3+\3+\7+\u0173\n+\f+\16+\u0176\13+\3+\3+\3,\3,\3,\3,\3-\3"+
		"-\3.\3.\3.\3.\5.\u0184\n.\3.\2\3 /\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\2\t\4\2\7\7\22\22\3\2\64\65"+
		"\4\2  &&\4\2\"\"$$\3\2\31\36\4\2!!##\4\2\'*..\2\u018c\2_\3\2\2\2\4n\3"+
		"\2\2\2\6q\3\2\2\2\b\u0080\3\2\2\2\n\u0084\3\2\2\2\f\u0086\3\2\2\2\16\u008b"+
		"\3\2\2\2\20\u0092\3\2\2\2\22\u0094\3\2\2\2\24\u009a\3\2\2\2\26\u00a2\3"+
		"\2\2\2\30\u00aa\3\2\2\2\32\u00b2\3\2\2\2\34\u00b4\3\2\2\2\36\u00b8\3\2"+
		"\2\2 \u00bf\3\2\2\2\"\u00d7\3\2\2\2$\u00df\3\2\2\2&\u00e1\3\2\2\2(\u00e3"+
		"\3\2\2\2*\u00e6\3\2\2\2,\u00ff\3\2\2\2.\u0101\3\2\2\2\60\u0109\3\2\2\2"+
		"\62\u010d\3\2\2\2\64\u010f\3\2\2\2\66\u011a\3\2\2\28\u0124\3\2\2\2:\u012a"+
		"\3\2\2\2<\u012c\3\2\2\2>\u0130\3\2\2\2@\u0134\3\2\2\2B\u0136\3\2\2\2D"+
		"\u0143\3\2\2\2F\u014c\3\2\2\2H\u0150\3\2\2\2J\u0155\3\2\2\2L\u0157\3\2"+
		"\2\2N\u015d\3\2\2\2P\u015f\3\2\2\2R\u0167\3\2\2\2T\u0169\3\2\2\2V\u0179"+
		"\3\2\2\2X\u017d\3\2\2\2Z\u0183\3\2\2\2\\]\5\4\3\2]^\5Z.\2^`\3\2\2\2_\\"+
		"\3\2\2\2_`\3\2\2\2`f\3\2\2\2ab\5\6\4\2bc\5Z.\2ce\3\2\2\2da\3\2\2\2eh\3"+
		"\2\2\2fd\3\2\2\2fg\3\2\2\2gj\3\2\2\2hf\3\2\2\2ik\5\16\b\2ji\3\2\2\2jk"+
		"\3\2\2\2kl\3\2\2\2lm\7\2\2\3m\3\3\2\2\2no\7\5\2\2op\5\f\7\2p\5\3\2\2\2"+
		"q}\7\6\2\2r~\5\b\5\2sy\7\b\2\2tu\5\b\5\2uv\5Z.\2vx\3\2\2\2wt\3\2\2\2x"+
		"{\3\2\2\2yw\3\2\2\2yz\3\2\2\2z|\3\2\2\2{y\3\2\2\2|~\7\t\2\2}r\3\2\2\2"+
		"}s\3\2\2\2~\7\3\2\2\2\177\u0081\t\2\2\2\u0080\177\3\2\2\2\u0080\u0081"+
		"\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0083\5\n\6\2\u0083\t\3\2\2\2\u0084"+
		"\u0085\5\f\7\2\u0085\13\3\2\2\2\u0086\u0087\t\3\2\2\u0087\r\3\2\2\2\u0088"+
		"\u0089\5\20\t\2\u0089\u008a\5Z.\2\u008a\u008c\3\2\2\2\u008b\u0088\3\2"+
		"\2\2\u008c\u008d\3\2\2\2\u008d\u008b\3\2\2\2\u008d\u008e\3\2\2\2\u008e"+
		"\17\3\2\2\2\u008f\u0093\5\22\n\2\u0090\u0093\5(\25\2\u0091\u0093\5T+\2"+
		"\u0092\u008f\3\2\2\2\u0092\u0090\3\2\2\2\u0092\u0091\3\2\2\2\u0093\21"+
		"\3\2\2\2\u0094\u0095\5\26\f\2\u0095\u0098\7\25\2\2\u0096\u0099\5\24\13"+
		"\2\u0097\u0099\5\30\r\2\u0098\u0096\3\2\2\2\u0098\u0097\3\2\2\2\u0099"+
		"\23\3\2\2\2\u009a\u009f\5 \21\2\u009b\u009c\7\17\2\2\u009c\u009e\5 \21"+
		"\2\u009d\u009b\3\2\2\2\u009e\u00a1\3\2\2\2\u009f\u009d\3\2\2\2\u009f\u00a0"+
		"\3\2\2\2\u00a0\25\3\2\2\2\u00a1\u009f\3\2\2\2\u00a2\u00a7\7\7\2\2\u00a3"+
		"\u00a4\7\17\2\2\u00a4\u00a6\7\7\2\2\u00a5\u00a3\3\2\2\2\u00a6\u00a9\3"+
		"\2\2\2\u00a7\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\27\3\2\2\2\u00a9"+
		"\u00a7\3\2\2\2\u00aa\u00ac\5\32\16\2\u00ab\u00ad\5\32\16\2\u00ac\u00ab"+
		"\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae\u00ac\3\2\2\2\u00ae\u00af\3\2\2\2\u00af"+
		"\31\3\2\2\2\u00b0\u00b3\5\34\17\2\u00b1\u00b3\5\36\20\2\u00b2\u00b0\3"+
		"\2\2\2\u00b2\u00b1\3\2\2\2\u00b3\33\3\2\2\2\u00b4\u00b5\5Z.\2\u00b5\u00b6"+
		"\7\37\2\2\u00b6\u00b7\5\26\f\2\u00b7\35\3\2\2\2\u00b8\u00b9\5Z.\2\u00b9"+
		"\u00ba\7\37\2\2\u00ba\u00bb\5 \21\2\u00bb\37\3\2\2\2\u00bc\u00bd\b\21"+
		"\1\2\u00bd\u00c0\5\"\22\2\u00be\u00c0\5\60\31\2\u00bf\u00bc\3\2\2\2\u00bf"+
		"\u00be\3\2\2\2\u00c0\u00d2\3\2\2\2\u00c1\u00c2\f\7\2\2\u00c2\u00c3\t\4"+
		"\2\2\u00c3\u00d1\5 \21\b\u00c4\u00c5\f\6\2\2\u00c5\u00c6\t\5\2\2\u00c6"+
		"\u00d1\5 \21\7\u00c7\u00c8\f\5\2\2\u00c8\u00c9\t\6\2\2\u00c9\u00d1\5 "+
		"\21\6\u00ca\u00cb\f\4\2\2\u00cb\u00cc\7\30\2\2\u00cc\u00d1\5 \21\5\u00cd"+
		"\u00ce\f\3\2\2\u00ce\u00cf\7\27\2\2\u00cf\u00d1\5 \21\4\u00d0\u00c1\3"+
		"\2\2\2\u00d0\u00c4\3\2\2\2\u00d0\u00c7\3\2\2\2\u00d0\u00ca\3\2\2\2\u00d0"+
		"\u00cd\3\2\2\2\u00d1\u00d4\3\2\2\2\u00d2\u00d0\3\2\2\2\u00d2\u00d3\3\2"+
		"\2\2\u00d3!\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d5\u00d8\5$\23\2\u00d6\u00d8"+
		"\5(\25\2\u00d7\u00d5\3\2\2\2\u00d7\u00d6\3\2\2\2\u00d8#\3\2\2\2\u00d9"+
		"\u00e0\5&\24\2\u00da\u00e0\5.\30\2\u00db\u00dc\7\b\2\2\u00dc\u00dd\5 "+
		"\21\2\u00dd\u00de\7\t\2\2\u00de\u00e0\3\2\2\2\u00df\u00d9\3\2\2\2\u00df"+
		"\u00da\3\2\2\2\u00df\u00db\3\2\2\2\u00e0%\3\2\2\2\u00e1\u00e2\5,\27\2"+
		"\u00e2\'\3\2\2\2\u00e3\u00e4\5P)\2\u00e4\u00e5\5*\26\2\u00e5)\3\2\2\2"+
		"\u00e6\u00f5\7\b\2\2\u00e7\u00ee\5\24\13\2\u00e8\u00eb\5N(\2\u00e9\u00ea"+
		"\7\17\2\2\u00ea\u00ec\5\24\13\2\u00eb\u00e9\3\2\2\2\u00eb\u00ec\3\2\2"+
		"\2\u00ec\u00ee\3\2\2\2\u00ed\u00e7\3\2\2\2\u00ed\u00e8\3\2\2\2\u00ee\u00f0"+
		"\3\2\2\2\u00ef\u00f1\7\26\2\2\u00f0\u00ef\3\2\2\2\u00f0\u00f1\3\2\2\2"+
		"\u00f1\u00f3\3\2\2\2\u00f2\u00f4\7\17\2\2\u00f3\u00f2\3\2\2\2\u00f3\u00f4"+
		"\3\2\2\2\u00f4\u00f6\3\2\2\2\u00f5\u00ed\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6"+
		"\u00f7\3\2\2\2\u00f7\u00f8\7\t\2\2\u00f8+\3\2\2\2\u00f9\u0100\5X-\2\u00fa"+
		"\u0100\5:\36\2\u00fb\u0100\5\62\32\2\u00fc\u0100\5\f\7\2\u00fd\u0100\7"+
		"+\2\2\u00fe\u0100\7.\2\2\u00ff\u00f9\3\2\2\2\u00ff\u00fa\3\2\2\2\u00ff"+
		"\u00fb\3\2\2\2\u00ff\u00fc\3\2\2\2\u00ff\u00fd\3\2\2\2\u00ff\u00fe\3\2"+
		"\2\2\u0100-\3\2\2\2\u0101\u0104\7\7\2\2\u0102\u0103\7\22\2\2\u0103\u0105"+
		"\7\7\2\2\u0104\u0102\3\2\2\2\u0104\u0105\3\2\2\2\u0105/\3\2\2\2\u0106"+
		"\u010a\5\"\22\2\u0107\u0108\t\7\2\2\u0108\u010a\5 \21\2\u0109\u0106\3"+
		"\2\2\2\u0109\u0107\3\2\2\2\u010a\61\3\2\2\2\u010b\u010e\5\64\33\2\u010c"+
		"\u010e\5\66\34\2\u010d\u010b\3\2\2\2\u010d\u010c\3\2\2\2\u010e\63\3\2"+
		"\2\2\u010f\u0110\7\f\2\2\u0110\u0115\5 \21\2\u0111\u0112\7\17\2\2\u0112"+
		"\u0114\5 \21\2\u0113\u0111\3\2\2\2\u0114\u0117\3\2\2\2\u0115\u0113\3\2"+
		"\2\2\u0115\u0116\3\2\2\2\u0116\u0118\3\2\2\2\u0117\u0115\3\2\2\2\u0118"+
		"\u0119\7\r\2\2\u0119\65\3\2\2\2\u011a\u011b\7\f\2\2\u011b\u011f\5Z.\2"+
		"\u011c\u011e\58\35\2\u011d\u011c\3\2\2\2\u011e\u0121\3\2\2\2\u011f\u011d"+
		"\3\2\2\2\u011f\u0120\3\2\2\2\u0120\u0122\3\2\2\2\u0121\u011f\3\2\2\2\u0122"+
		"\u0123\7\r\2\2\u0123\67\3\2\2\2\u0124\u0125\5 \21\2\u0125\u0126\7\17\2"+
		"\2\u0126\u0127\5Z.\2\u01279\3\2\2\2\u0128\u012b\5B\"\2\u0129\u012b\5D"+
		"#\2\u012a\u0128\3\2\2\2\u012a\u0129\3\2\2\2\u012b;\3\2\2\2\u012c\u012d"+
		"\5H%\2\u012d\u012e\7\21\2\2\u012e\u012f\5J&\2\u012f=\3\2\2\2\u0130\u0131"+
		"\7\7\2\2\u0131?\3\2\2\2\u0132\u0135\5> \2\u0133\u0135\5<\37\2\u0134\u0132"+
		"\3\2\2\2\u0134\u0133\3\2\2\2\u0135A\3\2\2\2\u0136\u0138\7\n\2\2\u0137"+
		"\u0139\5@!\2\u0138\u0137\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u013e\3\2\2"+
		"\2\u013a\u013b\7\17\2\2\u013b\u013d\5@!\2\u013c\u013a\3\2\2\2\u013d\u0140"+
		"\3\2\2\2\u013e\u013c\3\2\2\2\u013e\u013f\3\2\2\2\u013f\u0141\3\2\2\2\u0140"+
		"\u013e\3\2\2\2\u0141\u0142\7\13\2\2\u0142C\3\2\2\2\u0143\u0144\7\n\2\2"+
		"\u0144\u0146\5Z.\2\u0145\u0147\5F$\2\u0146\u0145\3\2\2\2\u0147\u0148\3"+
		"\2\2\2\u0148\u0146\3\2\2\2\u0148\u0149\3\2\2\2\u0149\u014a\3\2\2\2\u014a"+
		"\u014b\7\13\2\2\u014bE\3\2\2\2\u014c\u014d\5@!\2\u014d\u014e\7\17\2\2"+
		"\u014e\u014f\5Z.\2\u014fG\3\2\2\2\u0150\u0151\7\7\2\2\u0151I\3\2\2\2\u0152"+
		"\u0156\5 \21\2\u0153\u0156\5:\36\2\u0154\u0156\5\62\32\2\u0155\u0152\3"+
		"\2\2\2\u0155\u0153\3\2\2\2\u0155\u0154\3\2\2\2\u0156K\3\2\2\2\u0157\u0158"+
		"\7\f\2\2\u0158\u0159\7\r\2\2\u0159\u015a\5N(\2\u015aM\3\2\2\2\u015b\u015e"+
		"\5P)\2\u015c\u015e\5R*\2\u015d\u015b\3\2\2\2\u015d\u015c\3\2\2\2\u015e"+
		"O\3\2\2\2\u015f\u0164\7\7\2\2\u0160\u0161\7\22\2\2\u0161\u0163\7\7\2\2"+
		"\u0162\u0160\3\2\2\2\u0163\u0166\3\2\2\2\u0164\u0162\3\2\2\2\u0164\u0165"+
		"\3\2\2\2\u0165Q\3\2\2\2\u0166\u0164\3\2\2\2\u0167\u0168\5T+\2\u0168S\3"+
		"\2\2\2\u0169\u016a\7\3\2\2\u016a\u016b\7\7\2\2\u016b\u0174\7\n\2\2\u016c"+
		"\u016f\5V,\2\u016d\u016f\5P)\2\u016e\u016c\3\2\2\2\u016e\u016d\3\2\2\2"+
		"\u016f\u0170\3\2\2\2\u0170\u0171\5Z.\2\u0171\u0173\3\2\2\2\u0172\u016e"+
		"\3\2\2\2\u0173\u0176\3\2\2\2\u0174\u0172\3\2\2\2\u0174\u0175\3\2\2\2\u0175"+
		"\u0177\3\2\2\2\u0176\u0174\3\2\2\2\u0177\u0178\7\13\2\2\u0178U\3\2\2\2"+
		"\u0179\u017a\6,\7\2\u017a\u017b\7\7\2\2\u017b\u017c\5N(\2\u017cW\3\2\2"+
		"\2\u017d\u017e\t\b\2\2\u017eY\3\2\2\2\u017f\u0184\7\20\2\2\u0180\u0184"+
		"\7\2\2\3\u0181\u0184\6.\b\2\u0182\u0184\6.\t\2\u0183\u017f\3\2\2\2\u0183"+
		"\u0180\3\2\2\2\u0183\u0181\3\2\2\2\u0183\u0182\3\2\2\2\u0184[\3\2\2\2"+
		"*_fjy}\u0080\u008d\u0092\u0098\u009f\u00a7\u00ae\u00b2\u00bf\u00d0\u00d2"+
		"\u00d7\u00df\u00eb\u00ed\u00f0\u00f3\u00f5\u00ff\u0104\u0109\u010d\u0115"+
		"\u011f\u012a\u0134\u0138\u013e\u0148\u0155\u015d\u0164\u016e\u0174\u0183";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}