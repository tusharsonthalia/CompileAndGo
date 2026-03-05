// Generated from /Users/tusharsonthalia/Documents/Masters/UChicago/Academics/Q2/#1. Compilers/proj-compileandgo/grammars/GoliteParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GoliteParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMMENT=1, TYPE=2, STRUCT=3, FUNCTION=4, PRINT=5, DELETE=6, SCAN=7, FOR=8, 
		IF=9, ELSE=10, RETURN=11, VAR=12, INT=13, BOOL=14, TRUE=15, FALSE=16, 
		NIL=17, NEW=18, OR=19, AND=20, DOUBLEEQ=21, NEQ=22, GEQ=23, LEQ=24, LBRACE=25, 
		RBRACE=26, LPAREN=27, RPAREN=28, SEMICOLON=29, COMMA=30, DOT=31, EXCLAMATION=32, 
		FSLASH=33, ASTERISK=34, PLUS=35, MINUS=36, GT=37, LT=38, EQUALS=39, IDENTIFIER=40, 
		NUMBER=41, STRING=42, WS=43;
	public static final int
		RULE_program = 0, RULE_types = 1, RULE_typeDeclaration = 2, RULE_fields = 3, 
		RULE_decl = 4, RULE_type = 5, RULE_declarations = 6, RULE_declaration = 7, 
		RULE_ids = 8, RULE_functions = 9, RULE_function = 10, RULE_parameters = 11, 
		RULE_returnType = 12, RULE_statements = 13, RULE_statement = 14, RULE_block = 15, 
		RULE_delete = 16, RULE_read = 17, RULE_assignment = 18, RULE_print = 19, 
		RULE_conditional = 20, RULE_loop = 21, RULE_return = 22, RULE_invocation = 23, 
		RULE_arguments = 24, RULE_lvalue = 25, RULE_expression = 26, RULE_boolterm = 27, 
		RULE_equalterm = 28, RULE_relationterm = 29, RULE_simpleterm = 30, RULE_term = 31, 
		RULE_unaryterm = 32, RULE_selectorterm = 33, RULE_factor = 34;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "types", "typeDeclaration", "fields", "decl", "type", "declarations", 
			"declaration", "ids", "functions", "function", "parameters", "returnType", 
			"statements", "statement", "block", "delete", "read", "assignment", "print", 
			"conditional", "loop", "return", "invocation", "arguments", "lvalue", 
			"expression", "boolterm", "equalterm", "relationterm", "simpleterm", 
			"term", "unaryterm", "selectorterm", "factor"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'type'", "'struct'", "'func'", "'printf'", "'delete'", "'scan'", 
			"'for'", "'if'", "'else'", "'return'", "'var'", "'int'", "'bool'", "'true'", 
			"'false'", "'nil'", "'new'", "'||'", "'&&'", "'=='", "'!='", "'>='", 
			"'<='", "'{'", "'}'", "'('", "')'", "';'", "','", "'.'", "'!'", "'/'", 
			"'*'", "'+'", "'-'", "'>'", "'<'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMMENT", "TYPE", "STRUCT", "FUNCTION", "PRINT", "DELETE", "SCAN", 
			"FOR", "IF", "ELSE", "RETURN", "VAR", "INT", "BOOL", "TRUE", "FALSE", 
			"NIL", "NEW", "OR", "AND", "DOUBLEEQ", "NEQ", "GEQ", "LEQ", "LBRACE", 
			"RBRACE", "LPAREN", "RPAREN", "SEMICOLON", "COMMA", "DOT", "EXCLAMATION", 
			"FSLASH", "ASTERISK", "PLUS", "MINUS", "GT", "LT", "EQUALS", "IDENTIFIER", 
			"NUMBER", "STRING", "WS"
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
	public String getGrammarFileName() { return "GoliteParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GoliteParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GoliteParser.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			types();
			setState(71);
			declarations();
			setState(72);
			functions();
			setState(73);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypesContext extends ParserRuleContext {
		public List<TypeDeclarationContext> typeDeclaration() {
			return getRuleContexts(TypeDeclarationContext.class);
		}
		public TypeDeclarationContext typeDeclaration(int i) {
			return getRuleContext(TypeDeclarationContext.class,i);
		}
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_types);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE) {
				{
				{
				setState(75);
				typeDeclaration();
				}
				}
				setState(80);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclarationContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(GoliteParser.TYPE, 0); }
		public TerminalNode IDENTIFIER() { return getToken(GoliteParser.IDENTIFIER, 0); }
		public TerminalNode STRUCT() { return getToken(GoliteParser.STRUCT, 0); }
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public FieldsContext fields() {
			return getRuleContext(FieldsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public TypeDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDeclaration; }
	}

	public final TypeDeclarationContext typeDeclaration() throws RecognitionException {
		TypeDeclarationContext _localctx = new TypeDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_typeDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(81);
			match(TYPE);
			setState(82);
			match(IDENTIFIER);
			setState(83);
			match(STRUCT);
			setState(84);
			match(LBRACE);
			setState(85);
			fields();
			setState(86);
			match(RBRACE);
			setState(87);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FieldsContext extends ParserRuleContext {
		public List<DeclContext> decl() {
			return getRuleContexts(DeclContext.class);
		}
		public DeclContext decl(int i) {
			return getRuleContext(DeclContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(GoliteParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(GoliteParser.SEMICOLON, i);
		}
		public FieldsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fields; }
	}

	public final FieldsContext fields() throws RecognitionException {
		FieldsContext _localctx = new FieldsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_fields);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			decl();
			setState(90);
			match(SEMICOLON);
			setState(96);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(91);
				decl();
				setState(92);
				match(SEMICOLON);
				}
				}
				setState(98);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(GoliteParser.IDENTIFIER, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl; }
	}

	public final DeclContext decl() throws RecognitionException {
		DeclContext _localctx = new DeclContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			match(IDENTIFIER);
			setState(100);
			type();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(GoliteParser.INT, 0); }
		public TerminalNode BOOL() { return getToken(GoliteParser.BOOL, 0); }
		public TerminalNode ASTERISK() { return getToken(GoliteParser.ASTERISK, 0); }
		public TerminalNode IDENTIFIER() { return getToken(GoliteParser.IDENTIFIER, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_type);
		try {
			setState(106);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				match(INT);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				match(BOOL);
				}
				break;
			case ASTERISK:
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				match(ASTERISK);
				setState(105);
				match(IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationsContext extends ParserRuleContext {
		public List<DeclarationContext> declaration() {
			return getRuleContexts(DeclarationContext.class);
		}
		public DeclarationContext declaration(int i) {
			return getRuleContext(DeclarationContext.class,i);
		}
		public DeclarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarations; }
	}

	public final DeclarationsContext declarations() throws RecognitionException {
		DeclarationsContext _localctx = new DeclarationsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(108);
				declaration();
				}
				}
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(GoliteParser.VAR, 0); }
		public IdsContext ids() {
			return getRuleContext(IdsContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(VAR);
			setState(115);
			ids();
			setState(116);
			type();
			setState(117);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IdsContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(GoliteParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(GoliteParser.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public IdsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ids; }
	}

	public final IdsContext ids() throws RecognitionException {
		IdsContext _localctx = new IdsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_ids);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(IDENTIFIER);
			setState(124);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(120);
				match(COMMA);
				setState(121);
				match(IDENTIFIER);
				}
				}
				setState(126);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionsContext extends ParserRuleContext {
		public List<FunctionContext> function() {
			return getRuleContexts(FunctionContext.class);
		}
		public FunctionContext function(int i) {
			return getRuleContext(FunctionContext.class,i);
		}
		public FunctionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functions; }
	}

	public final FunctionsContext functions() throws RecognitionException {
		FunctionsContext _localctx = new FunctionsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(127);
				function();
				}
				}
				setState(132);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionContext extends ParserRuleContext {
		public TerminalNode FUNCTION() { return getToken(GoliteParser.FUNCTION, 0); }
		public TerminalNode IDENTIFIER() { return getToken(GoliteParser.IDENTIFIER, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public ReturnTypeContext returnType() {
			return getRuleContext(ReturnTypeContext.class,0);
		}
		public FunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function; }
	}

	public final FunctionContext function() throws RecognitionException {
		FunctionContext _localctx = new FunctionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_function);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			match(FUNCTION);
			setState(134);
			match(IDENTIFIER);
			setState(135);
			parameters();
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 17179893760L) != 0)) {
				{
				setState(136);
				returnType();
				}
			}

			setState(139);
			match(LBRACE);
			setState(140);
			declarations();
			setState(141);
			statements();
			setState(142);
			match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ParametersContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public List<DeclContext> decl() {
			return getRuleContexts(DeclContext.class);
		}
		public DeclContext decl(int i) {
			return getRuleContext(DeclContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(LPAREN);
			setState(153);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(145);
				decl();
				setState(150);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(146);
					match(COMMA);
					setState(147);
					decl();
					}
					}
					setState(152);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(155);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ReturnTypeContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ReturnTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnType; }
	}

	public final ReturnTypeContext returnType() throws RecognitionException {
		ReturnTypeContext _localctx = new ReturnTypeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_returnType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(157);
			type();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
	}

	public final StatementsContext statements() throws RecognitionException {
		StatementsContext _localctx = new StatementsContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_statements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1099511630816L) != 0)) {
				{
				{
				setState(159);
				statement();
				}
				}
				setState(164);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public ReadContext read() {
			return getRuleContext(ReadContext.class,0);
		}
		public DeleteContext delete() {
			return getRuleContext(DeleteContext.class,0);
		}
		public ConditionalContext conditional() {
			return getRuleContext(ConditionalContext.class,0);
		}
		public LoopContext loop() {
			return getRuleContext(LoopContext.class,0);
		}
		public ReturnContext return_() {
			return getRuleContext(ReturnContext.class,0);
		}
		public InvocationContext invocation() {
			return getRuleContext(InvocationContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_statement);
		try {
			setState(173);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				assignment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
				print();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(167);
				read();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(168);
				delete();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(169);
				conditional();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(170);
				loop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(171);
				return_();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(172);
				invocation();
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

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(LBRACE);
			setState(176);
			statements();
			setState(177);
			match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeleteContext extends ParserRuleContext {
		public TerminalNode DELETE() { return getToken(GoliteParser.DELETE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeleteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delete; }
	}

	public final DeleteContext delete() throws RecognitionException {
		DeleteContext _localctx = new DeleteContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_delete);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			match(DELETE);
			setState(180);
			expression();
			setState(181);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ReadContext extends ParserRuleContext {
		public TerminalNode SCAN() { return getToken(GoliteParser.SCAN, 0); }
		public LvalueContext lvalue() {
			return getRuleContext(LvalueContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ReadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_read; }
	}

	public final ReadContext read() throws RecognitionException {
		ReadContext _localctx = new ReadContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_read);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(SCAN);
			setState(184);
			lvalue();
			setState(185);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ParserRuleContext {
		public LvalueContext lvalue() {
			return getRuleContext(LvalueContext.class,0);
		}
		public TerminalNode EQUALS() { return getToken(GoliteParser.EQUALS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			lvalue();
			setState(188);
			match(EQUALS);
			setState(189);
			expression();
			setState(190);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrintContext extends ParserRuleContext {
		public TerminalNode PRINT() { return getToken(GoliteParser.PRINT, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode STRING() { return getToken(GoliteParser.STRING, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(PRINT);
			setState(193);
			match(LPAREN);
			setState(194);
			match(STRING);
			setState(199);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(195);
				match(COMMA);
				setState(196);
				expression();
				}
				}
				setState(201);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(202);
			match(RPAREN);
			setState(203);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionalContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(GoliteParser.IF, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(GoliteParser.ELSE, 0); }
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			match(IF);
			setState(206);
			match(LPAREN);
			setState(207);
			expression();
			setState(208);
			match(RPAREN);
			setState(209);
			block();
			setState(212);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(210);
				match(ELSE);
				setState(211);
				block();
				}
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

	@SuppressWarnings("CheckReturnValue")
	public static class LoopContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(GoliteParser.FOR, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop; }
	}

	public final LoopContext loop() throws RecognitionException {
		LoopContext _localctx = new LoopContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_loop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(FOR);
			setState(215);
			match(LPAREN);
			setState(216);
			expression();
			setState(217);
			match(RPAREN);
			setState(218);
			block();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ReturnContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(GoliteParser.RETURN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ReturnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return; }
	}

	public final ReturnContext return_() throws RecognitionException {
		ReturnContext _localctx = new ReturnContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_return);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(220);
			match(RETURN);
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3371684036608L) != 0)) {
				{
				setState(221);
				expression();
				}
			}

			setState(224);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InvocationContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(GoliteParser.IDENTIFIER, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public InvocationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_invocation; }
	}

	public final InvocationContext invocation() throws RecognitionException {
		InvocationContext _localctx = new InvocationContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_invocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			match(IDENTIFIER);
			setState(227);
			arguments();
			setState(228);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoliteParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoliteParser.COMMA, i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(230);
			match(LPAREN);
			setState(239);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3371684036608L) != 0)) {
				{
				setState(231);
				expression();
				setState(236);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(232);
					match(COMMA);
					setState(233);
					expression();
					}
					}
					setState(238);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(241);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LvalueContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(GoliteParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(GoliteParser.IDENTIFIER, i);
		}
		public List<TerminalNode> DOT() { return getTokens(GoliteParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(GoliteParser.DOT, i);
		}
		public LvalueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lvalue; }
	}

	public final LvalueContext lvalue() throws RecognitionException {
		LvalueContext _localctx = new LvalueContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_lvalue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			match(IDENTIFIER);
			setState(248);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(244);
				match(DOT);
				setState(245);
				match(IDENTIFIER);
				}
				}
				setState(250);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public List<BooltermContext> boolterm() {
			return getRuleContexts(BooltermContext.class);
		}
		public BooltermContext boolterm(int i) {
			return getRuleContext(BooltermContext.class,i);
		}
		public List<TerminalNode> OR() { return getTokens(GoliteParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(GoliteParser.OR, i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(251);
			boolterm();
			setState(256);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(252);
				match(OR);
				setState(253);
				boolterm();
				}
				}
				setState(258);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BooltermContext extends ParserRuleContext {
		public List<EqualtermContext> equalterm() {
			return getRuleContexts(EqualtermContext.class);
		}
		public EqualtermContext equalterm(int i) {
			return getRuleContext(EqualtermContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(GoliteParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(GoliteParser.AND, i);
		}
		public BooltermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolterm; }
	}

	public final BooltermContext boolterm() throws RecognitionException {
		BooltermContext _localctx = new BooltermContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_boolterm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			equalterm();
			setState(264);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(260);
				match(AND);
				setState(261);
				equalterm();
				}
				}
				setState(266);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class EqualtermContext extends ParserRuleContext {
		public List<RelationtermContext> relationterm() {
			return getRuleContexts(RelationtermContext.class);
		}
		public RelationtermContext relationterm(int i) {
			return getRuleContext(RelationtermContext.class,i);
		}
		public List<TerminalNode> DOUBLEEQ() { return getTokens(GoliteParser.DOUBLEEQ); }
		public TerminalNode DOUBLEEQ(int i) {
			return getToken(GoliteParser.DOUBLEEQ, i);
		}
		public List<TerminalNode> NEQ() { return getTokens(GoliteParser.NEQ); }
		public TerminalNode NEQ(int i) {
			return getToken(GoliteParser.NEQ, i);
		}
		public EqualtermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalterm; }
	}

	public final EqualtermContext equalterm() throws RecognitionException {
		EqualtermContext _localctx = new EqualtermContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_equalterm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(267);
			relationterm();
			setState(272);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOUBLEEQ || _la==NEQ) {
				{
				{
				setState(268);
				_la = _input.LA(1);
				if ( !(_la==DOUBLEEQ || _la==NEQ) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(269);
				relationterm();
				}
				}
				setState(274);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class RelationtermContext extends ParserRuleContext {
		public List<SimpletermContext> simpleterm() {
			return getRuleContexts(SimpletermContext.class);
		}
		public SimpletermContext simpleterm(int i) {
			return getRuleContext(SimpletermContext.class,i);
		}
		public List<TerminalNode> GT() { return getTokens(GoliteParser.GT); }
		public TerminalNode GT(int i) {
			return getToken(GoliteParser.GT, i);
		}
		public List<TerminalNode> LT() { return getTokens(GoliteParser.LT); }
		public TerminalNode LT(int i) {
			return getToken(GoliteParser.LT, i);
		}
		public List<TerminalNode> GEQ() { return getTokens(GoliteParser.GEQ); }
		public TerminalNode GEQ(int i) {
			return getToken(GoliteParser.GEQ, i);
		}
		public List<TerminalNode> LEQ() { return getTokens(GoliteParser.LEQ); }
		public TerminalNode LEQ(int i) {
			return getToken(GoliteParser.LEQ, i);
		}
		public RelationtermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationterm; }
	}

	public final RelationtermContext relationterm() throws RecognitionException {
		RelationtermContext _localctx = new RelationtermContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_relationterm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(275);
			simpleterm();
			setState(280);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 412342026240L) != 0)) {
				{
				{
				setState(276);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 412342026240L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(277);
				simpleterm();
				}
				}
				setState(282);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SimpletermContext extends ParserRuleContext {
		public List<TermContext> term() {
			return getRuleContexts(TermContext.class);
		}
		public TermContext term(int i) {
			return getRuleContext(TermContext.class,i);
		}
		public List<TerminalNode> PLUS() { return getTokens(GoliteParser.PLUS); }
		public TerminalNode PLUS(int i) {
			return getToken(GoliteParser.PLUS, i);
		}
		public List<TerminalNode> MINUS() { return getTokens(GoliteParser.MINUS); }
		public TerminalNode MINUS(int i) {
			return getToken(GoliteParser.MINUS, i);
		}
		public SimpletermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleterm; }
	}

	public final SimpletermContext simpleterm() throws RecognitionException {
		SimpletermContext _localctx = new SimpletermContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_simpleterm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			term();
			setState(288);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PLUS || _la==MINUS) {
				{
				{
				setState(284);
				_la = _input.LA(1);
				if ( !(_la==PLUS || _la==MINUS) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(285);
				term();
				}
				}
				setState(290);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TermContext extends ParserRuleContext {
		public List<UnarytermContext> unaryterm() {
			return getRuleContexts(UnarytermContext.class);
		}
		public UnarytermContext unaryterm(int i) {
			return getRuleContext(UnarytermContext.class,i);
		}
		public List<TerminalNode> ASTERISK() { return getTokens(GoliteParser.ASTERISK); }
		public TerminalNode ASTERISK(int i) {
			return getToken(GoliteParser.ASTERISK, i);
		}
		public List<TerminalNode> FSLASH() { return getTokens(GoliteParser.FSLASH); }
		public TerminalNode FSLASH(int i) {
			return getToken(GoliteParser.FSLASH, i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(291);
			unaryterm();
			setState(296);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FSLASH || _la==ASTERISK) {
				{
				{
				setState(292);
				_la = _input.LA(1);
				if ( !(_la==FSLASH || _la==ASTERISK) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(293);
				unaryterm();
				}
				}
				setState(298);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class UnarytermContext extends ParserRuleContext {
		public SelectortermContext selectorterm() {
			return getRuleContext(SelectortermContext.class,0);
		}
		public TerminalNode EXCLAMATION() { return getToken(GoliteParser.EXCLAMATION, 0); }
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public UnarytermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryterm; }
	}

	public final UnarytermContext unaryterm() throws RecognitionException {
		UnarytermContext _localctx = new UnarytermContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_unaryterm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(300);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EXCLAMATION || _la==MINUS) {
				{
				setState(299);
				_la = _input.LA(1);
				if ( !(_la==EXCLAMATION || _la==MINUS) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(302);
			selectorterm();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SelectortermContext extends ParserRuleContext {
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public List<TerminalNode> DOT() { return getTokens(GoliteParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(GoliteParser.DOT, i);
		}
		public List<TerminalNode> IDENTIFIER() { return getTokens(GoliteParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(GoliteParser.IDENTIFIER, i);
		}
		public SelectortermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorterm; }
	}

	public final SelectortermContext selectorterm() throws RecognitionException {
		SelectortermContext _localctx = new SelectortermContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_selectorterm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			factor();
			setState(309);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(305);
				match(DOT);
				setState(306);
				match(IDENTIFIER);
				}
				}
				setState(311);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FactorContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public TerminalNode IDENTIFIER() { return getToken(GoliteParser.IDENTIFIER, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(GoliteParser.NUMBER, 0); }
		public TerminalNode NEW() { return getToken(GoliteParser.NEW, 0); }
		public TerminalNode TRUE() { return getToken(GoliteParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(GoliteParser.FALSE, 0); }
		public TerminalNode NIL() { return getToken(GoliteParser.NIL, 0); }
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_factor);
		int _la;
		try {
			setState(326);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(312);
				match(LPAREN);
				setState(313);
				expression();
				setState(314);
				match(RPAREN);
				}
				break;
			case IDENTIFIER:
				enterOuterAlt(_localctx, 2);
				{
				setState(316);
				match(IDENTIFIER);
				setState(318);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LPAREN) {
					{
					setState(317);
					arguments();
					}
				}

				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 3);
				{
				setState(320);
				match(NUMBER);
				}
				break;
			case NEW:
				enterOuterAlt(_localctx, 4);
				{
				setState(321);
				match(NEW);
				setState(322);
				match(IDENTIFIER);
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 5);
				{
				setState(323);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 6);
				{
				setState(324);
				match(FALSE);
				}
				break;
			case NIL:
				enterOuterAlt(_localctx, 7);
				{
				setState(325);
				match(NIL);
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

	public static final String _serializedATN =
		"\u0004\u0001+\u0149\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0005"+
		"\u0001M\b\u0001\n\u0001\f\u0001P\t\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003_\b"+
		"\u0003\n\u0003\f\u0003b\t\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005k\b\u0005\u0001"+
		"\u0006\u0005\u0006n\b\u0006\n\u0006\f\u0006q\t\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0005"+
		"\b{\b\b\n\b\f\b~\t\b\u0001\t\u0005\t\u0081\b\t\n\t\f\t\u0084\t\t\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0003\n\u008a\b\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b"+
		"\u0095\b\u000b\n\u000b\f\u000b\u0098\t\u000b\u0003\u000b\u009a\b\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0005\r\u00a1\b\r\n\r"+
		"\f\r\u00a4\t\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00ae\b\u000e\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u00c6\b\u0013\n\u0013"+
		"\f\u0013\u00c9\t\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0003\u0014\u00d5\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0003\u0016\u00df\b\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0005\u0018\u00eb\b\u0018"+
		"\n\u0018\f\u0018\u00ee\t\u0018\u0003\u0018\u00f0\b\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u00f7\b\u0019\n"+
		"\u0019\f\u0019\u00fa\t\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0005"+
		"\u001a\u00ff\b\u001a\n\u001a\f\u001a\u0102\t\u001a\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0005\u001b\u0107\b\u001b\n\u001b\f\u001b\u010a\t\u001b\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0005\u001c\u010f\b\u001c\n\u001c\f\u001c"+
		"\u0112\t\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u0117\b"+
		"\u001d\n\u001d\f\u001d\u011a\t\u001d\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0005\u001e\u011f\b\u001e\n\u001e\f\u001e\u0122\t\u001e\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0005\u001f\u0127\b\u001f\n\u001f\f\u001f\u012a\t\u001f"+
		"\u0001 \u0003 \u012d\b \u0001 \u0001 \u0001!\u0001!\u0001!\u0005!\u0134"+
		"\b!\n!\f!\u0137\t!\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0003"+
		"\"\u013f\b\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0003\"\u0147"+
		"\b\"\u0001\"\u0000\u0000#\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BD\u0000\u0005\u0001"+
		"\u0000\u0015\u0016\u0002\u0000\u0017\u0018%&\u0001\u0000#$\u0001\u0000"+
		"!\"\u0002\u0000  $$\u014c\u0000F\u0001\u0000\u0000\u0000\u0002N\u0001"+
		"\u0000\u0000\u0000\u0004Q\u0001\u0000\u0000\u0000\u0006Y\u0001\u0000\u0000"+
		"\u0000\bc\u0001\u0000\u0000\u0000\nj\u0001\u0000\u0000\u0000\fo\u0001"+
		"\u0000\u0000\u0000\u000er\u0001\u0000\u0000\u0000\u0010w\u0001\u0000\u0000"+
		"\u0000\u0012\u0082\u0001\u0000\u0000\u0000\u0014\u0085\u0001\u0000\u0000"+
		"\u0000\u0016\u0090\u0001\u0000\u0000\u0000\u0018\u009d\u0001\u0000\u0000"+
		"\u0000\u001a\u00a2\u0001\u0000\u0000\u0000\u001c\u00ad\u0001\u0000\u0000"+
		"\u0000\u001e\u00af\u0001\u0000\u0000\u0000 \u00b3\u0001\u0000\u0000\u0000"+
		"\"\u00b7\u0001\u0000\u0000\u0000$\u00bb\u0001\u0000\u0000\u0000&\u00c0"+
		"\u0001\u0000\u0000\u0000(\u00cd\u0001\u0000\u0000\u0000*\u00d6\u0001\u0000"+
		"\u0000\u0000,\u00dc\u0001\u0000\u0000\u0000.\u00e2\u0001\u0000\u0000\u0000"+
		"0\u00e6\u0001\u0000\u0000\u00002\u00f3\u0001\u0000\u0000\u00004\u00fb"+
		"\u0001\u0000\u0000\u00006\u0103\u0001\u0000\u0000\u00008\u010b\u0001\u0000"+
		"\u0000\u0000:\u0113\u0001\u0000\u0000\u0000<\u011b\u0001\u0000\u0000\u0000"+
		">\u0123\u0001\u0000\u0000\u0000@\u012c\u0001\u0000\u0000\u0000B\u0130"+
		"\u0001\u0000\u0000\u0000D\u0146\u0001\u0000\u0000\u0000FG\u0003\u0002"+
		"\u0001\u0000GH\u0003\f\u0006\u0000HI\u0003\u0012\t\u0000IJ\u0005\u0000"+
		"\u0000\u0001J\u0001\u0001\u0000\u0000\u0000KM\u0003\u0004\u0002\u0000"+
		"LK\u0001\u0000\u0000\u0000MP\u0001\u0000\u0000\u0000NL\u0001\u0000\u0000"+
		"\u0000NO\u0001\u0000\u0000\u0000O\u0003\u0001\u0000\u0000\u0000PN\u0001"+
		"\u0000\u0000\u0000QR\u0005\u0002\u0000\u0000RS\u0005(\u0000\u0000ST\u0005"+
		"\u0003\u0000\u0000TU\u0005\u0019\u0000\u0000UV\u0003\u0006\u0003\u0000"+
		"VW\u0005\u001a\u0000\u0000WX\u0005\u001d\u0000\u0000X\u0005\u0001\u0000"+
		"\u0000\u0000YZ\u0003\b\u0004\u0000Z`\u0005\u001d\u0000\u0000[\\\u0003"+
		"\b\u0004\u0000\\]\u0005\u001d\u0000\u0000]_\u0001\u0000\u0000\u0000^["+
		"\u0001\u0000\u0000\u0000_b\u0001\u0000\u0000\u0000`^\u0001\u0000\u0000"+
		"\u0000`a\u0001\u0000\u0000\u0000a\u0007\u0001\u0000\u0000\u0000b`\u0001"+
		"\u0000\u0000\u0000cd\u0005(\u0000\u0000de\u0003\n\u0005\u0000e\t\u0001"+
		"\u0000\u0000\u0000fk\u0005\r\u0000\u0000gk\u0005\u000e\u0000\u0000hi\u0005"+
		"\"\u0000\u0000ik\u0005(\u0000\u0000jf\u0001\u0000\u0000\u0000jg\u0001"+
		"\u0000\u0000\u0000jh\u0001\u0000\u0000\u0000k\u000b\u0001\u0000\u0000"+
		"\u0000ln\u0003\u000e\u0007\u0000ml\u0001\u0000\u0000\u0000nq\u0001\u0000"+
		"\u0000\u0000om\u0001\u0000\u0000\u0000op\u0001\u0000\u0000\u0000p\r\u0001"+
		"\u0000\u0000\u0000qo\u0001\u0000\u0000\u0000rs\u0005\f\u0000\u0000st\u0003"+
		"\u0010\b\u0000tu\u0003\n\u0005\u0000uv\u0005\u001d\u0000\u0000v\u000f"+
		"\u0001\u0000\u0000\u0000w|\u0005(\u0000\u0000xy\u0005\u001e\u0000\u0000"+
		"y{\u0005(\u0000\u0000zx\u0001\u0000\u0000\u0000{~\u0001\u0000\u0000\u0000"+
		"|z\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000}\u0011\u0001\u0000"+
		"\u0000\u0000~|\u0001\u0000\u0000\u0000\u007f\u0081\u0003\u0014\n\u0000"+
		"\u0080\u007f\u0001\u0000\u0000\u0000\u0081\u0084\u0001\u0000\u0000\u0000"+
		"\u0082\u0080\u0001\u0000\u0000\u0000\u0082\u0083\u0001\u0000\u0000\u0000"+
		"\u0083\u0013\u0001\u0000\u0000\u0000\u0084\u0082\u0001\u0000\u0000\u0000"+
		"\u0085\u0086\u0005\u0004\u0000\u0000\u0086\u0087\u0005(\u0000\u0000\u0087"+
		"\u0089\u0003\u0016\u000b\u0000\u0088\u008a\u0003\u0018\f\u0000\u0089\u0088"+
		"\u0001\u0000\u0000\u0000\u0089\u008a\u0001\u0000\u0000\u0000\u008a\u008b"+
		"\u0001\u0000\u0000\u0000\u008b\u008c\u0005\u0019\u0000\u0000\u008c\u008d"+
		"\u0003\f\u0006\u0000\u008d\u008e\u0003\u001a\r\u0000\u008e\u008f\u0005"+
		"\u001a\u0000\u0000\u008f\u0015\u0001\u0000\u0000\u0000\u0090\u0099\u0005"+
		"\u001b\u0000\u0000\u0091\u0096\u0003\b\u0004\u0000\u0092\u0093\u0005\u001e"+
		"\u0000\u0000\u0093\u0095\u0003\b\u0004\u0000\u0094\u0092\u0001\u0000\u0000"+
		"\u0000\u0095\u0098\u0001\u0000\u0000\u0000\u0096\u0094\u0001\u0000\u0000"+
		"\u0000\u0096\u0097\u0001\u0000\u0000\u0000\u0097\u009a\u0001\u0000\u0000"+
		"\u0000\u0098\u0096\u0001\u0000\u0000\u0000\u0099\u0091\u0001\u0000\u0000"+
		"\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a\u009b\u0001\u0000\u0000"+
		"\u0000\u009b\u009c\u0005\u001c\u0000\u0000\u009c\u0017\u0001\u0000\u0000"+
		"\u0000\u009d\u009e\u0003\n\u0005\u0000\u009e\u0019\u0001\u0000\u0000\u0000"+
		"\u009f\u00a1\u0003\u001c\u000e\u0000\u00a0\u009f\u0001\u0000\u0000\u0000"+
		"\u00a1\u00a4\u0001\u0000\u0000\u0000\u00a2\u00a0\u0001\u0000\u0000\u0000"+
		"\u00a2\u00a3\u0001\u0000\u0000\u0000\u00a3\u001b\u0001\u0000\u0000\u0000"+
		"\u00a4\u00a2\u0001\u0000\u0000\u0000\u00a5\u00ae\u0003$\u0012\u0000\u00a6"+
		"\u00ae\u0003&\u0013\u0000\u00a7\u00ae\u0003\"\u0011\u0000\u00a8\u00ae"+
		"\u0003 \u0010\u0000\u00a9\u00ae\u0003(\u0014\u0000\u00aa\u00ae\u0003*"+
		"\u0015\u0000\u00ab\u00ae\u0003,\u0016\u0000\u00ac\u00ae\u0003.\u0017\u0000"+
		"\u00ad\u00a5\u0001\u0000\u0000\u0000\u00ad\u00a6\u0001\u0000\u0000\u0000"+
		"\u00ad\u00a7\u0001\u0000\u0000\u0000\u00ad\u00a8\u0001\u0000\u0000\u0000"+
		"\u00ad\u00a9\u0001\u0000\u0000\u0000\u00ad\u00aa\u0001\u0000\u0000\u0000"+
		"\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ac\u0001\u0000\u0000\u0000"+
		"\u00ae\u001d\u0001\u0000\u0000\u0000\u00af\u00b0\u0005\u0019\u0000\u0000"+
		"\u00b0\u00b1\u0003\u001a\r\u0000\u00b1\u00b2\u0005\u001a\u0000\u0000\u00b2"+
		"\u001f\u0001\u0000\u0000\u0000\u00b3\u00b4\u0005\u0006\u0000\u0000\u00b4"+
		"\u00b5\u00034\u001a\u0000\u00b5\u00b6\u0005\u001d\u0000\u0000\u00b6!\u0001"+
		"\u0000\u0000\u0000\u00b7\u00b8\u0005\u0007\u0000\u0000\u00b8\u00b9\u0003"+
		"2\u0019\u0000\u00b9\u00ba\u0005\u001d\u0000\u0000\u00ba#\u0001\u0000\u0000"+
		"\u0000\u00bb\u00bc\u00032\u0019\u0000\u00bc\u00bd\u0005\'\u0000\u0000"+
		"\u00bd\u00be\u00034\u001a\u0000\u00be\u00bf\u0005\u001d\u0000\u0000\u00bf"+
		"%\u0001\u0000\u0000\u0000\u00c0\u00c1\u0005\u0005\u0000\u0000\u00c1\u00c2"+
		"\u0005\u001b\u0000\u0000\u00c2\u00c7\u0005*\u0000\u0000\u00c3\u00c4\u0005"+
		"\u001e\u0000\u0000\u00c4\u00c6\u00034\u001a\u0000\u00c5\u00c3\u0001\u0000"+
		"\u0000\u0000\u00c6\u00c9\u0001\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000"+
		"\u0000\u0000\u00c7\u00c8\u0001\u0000\u0000\u0000\u00c8\u00ca\u0001\u0000"+
		"\u0000\u0000\u00c9\u00c7\u0001\u0000\u0000\u0000\u00ca\u00cb\u0005\u001c"+
		"\u0000\u0000\u00cb\u00cc\u0005\u001d\u0000\u0000\u00cc\'\u0001\u0000\u0000"+
		"\u0000\u00cd\u00ce\u0005\t\u0000\u0000\u00ce\u00cf\u0005\u001b\u0000\u0000"+
		"\u00cf\u00d0\u00034\u001a\u0000\u00d0\u00d1\u0005\u001c\u0000\u0000\u00d1"+
		"\u00d4\u0003\u001e\u000f\u0000\u00d2\u00d3\u0005\n\u0000\u0000\u00d3\u00d5"+
		"\u0003\u001e\u000f\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000\u00d4\u00d5"+
		"\u0001\u0000\u0000\u0000\u00d5)\u0001\u0000\u0000\u0000\u00d6\u00d7\u0005"+
		"\b\u0000\u0000\u00d7\u00d8\u0005\u001b\u0000\u0000\u00d8\u00d9\u00034"+
		"\u001a\u0000\u00d9\u00da\u0005\u001c\u0000\u0000\u00da\u00db\u0003\u001e"+
		"\u000f\u0000\u00db+\u0001\u0000\u0000\u0000\u00dc\u00de\u0005\u000b\u0000"+
		"\u0000\u00dd\u00df\u00034\u001a\u0000\u00de\u00dd\u0001\u0000\u0000\u0000"+
		"\u00de\u00df\u0001\u0000\u0000\u0000\u00df\u00e0\u0001\u0000\u0000\u0000"+
		"\u00e0\u00e1\u0005\u001d\u0000\u0000\u00e1-\u0001\u0000\u0000\u0000\u00e2"+
		"\u00e3\u0005(\u0000\u0000\u00e3\u00e4\u00030\u0018\u0000\u00e4\u00e5\u0005"+
		"\u001d\u0000\u0000\u00e5/\u0001\u0000\u0000\u0000\u00e6\u00ef\u0005\u001b"+
		"\u0000\u0000\u00e7\u00ec\u00034\u001a\u0000\u00e8\u00e9\u0005\u001e\u0000"+
		"\u0000\u00e9\u00eb\u00034\u001a\u0000\u00ea\u00e8\u0001\u0000\u0000\u0000"+
		"\u00eb\u00ee\u0001\u0000\u0000\u0000\u00ec\u00ea\u0001\u0000\u0000\u0000"+
		"\u00ec\u00ed\u0001\u0000\u0000\u0000\u00ed\u00f0\u0001\u0000\u0000\u0000"+
		"\u00ee\u00ec\u0001\u0000\u0000\u0000\u00ef\u00e7\u0001\u0000\u0000\u0000"+
		"\u00ef\u00f0\u0001\u0000\u0000\u0000\u00f0\u00f1\u0001\u0000\u0000\u0000"+
		"\u00f1\u00f2\u0005\u001c\u0000\u0000\u00f21\u0001\u0000\u0000\u0000\u00f3"+
		"\u00f8\u0005(\u0000\u0000\u00f4\u00f5\u0005\u001f\u0000\u0000\u00f5\u00f7"+
		"\u0005(\u0000\u0000\u00f6\u00f4\u0001\u0000\u0000\u0000\u00f7\u00fa\u0001"+
		"\u0000\u0000\u0000\u00f8\u00f6\u0001\u0000\u0000\u0000\u00f8\u00f9\u0001"+
		"\u0000\u0000\u0000\u00f93\u0001\u0000\u0000\u0000\u00fa\u00f8\u0001\u0000"+
		"\u0000\u0000\u00fb\u0100\u00036\u001b\u0000\u00fc\u00fd\u0005\u0013\u0000"+
		"\u0000\u00fd\u00ff\u00036\u001b\u0000\u00fe\u00fc\u0001\u0000\u0000\u0000"+
		"\u00ff\u0102\u0001\u0000\u0000\u0000\u0100\u00fe\u0001\u0000\u0000\u0000"+
		"\u0100\u0101\u0001\u0000\u0000\u0000\u01015\u0001\u0000\u0000\u0000\u0102"+
		"\u0100\u0001\u0000\u0000\u0000\u0103\u0108\u00038\u001c\u0000\u0104\u0105"+
		"\u0005\u0014\u0000\u0000\u0105\u0107\u00038\u001c\u0000\u0106\u0104\u0001"+
		"\u0000\u0000\u0000\u0107\u010a\u0001\u0000\u0000\u0000\u0108\u0106\u0001"+
		"\u0000\u0000\u0000\u0108\u0109\u0001\u0000\u0000\u0000\u01097\u0001\u0000"+
		"\u0000\u0000\u010a\u0108\u0001\u0000\u0000\u0000\u010b\u0110\u0003:\u001d"+
		"\u0000\u010c\u010d\u0007\u0000\u0000\u0000\u010d\u010f\u0003:\u001d\u0000"+
		"\u010e\u010c\u0001\u0000\u0000\u0000\u010f\u0112\u0001\u0000\u0000\u0000"+
		"\u0110\u010e\u0001\u0000\u0000\u0000\u0110\u0111\u0001\u0000\u0000\u0000"+
		"\u01119\u0001\u0000\u0000\u0000\u0112\u0110\u0001\u0000\u0000\u0000\u0113"+
		"\u0118\u0003<\u001e\u0000\u0114\u0115\u0007\u0001\u0000\u0000\u0115\u0117"+
		"\u0003<\u001e\u0000\u0116\u0114\u0001\u0000\u0000\u0000\u0117\u011a\u0001"+
		"\u0000\u0000\u0000\u0118\u0116\u0001\u0000\u0000\u0000\u0118\u0119\u0001"+
		"\u0000\u0000\u0000\u0119;\u0001\u0000\u0000\u0000\u011a\u0118\u0001\u0000"+
		"\u0000\u0000\u011b\u0120\u0003>\u001f\u0000\u011c\u011d\u0007\u0002\u0000"+
		"\u0000\u011d\u011f\u0003>\u001f\u0000\u011e\u011c\u0001\u0000\u0000\u0000"+
		"\u011f\u0122\u0001\u0000\u0000\u0000\u0120\u011e\u0001\u0000\u0000\u0000"+
		"\u0120\u0121\u0001\u0000\u0000\u0000\u0121=\u0001\u0000\u0000\u0000\u0122"+
		"\u0120\u0001\u0000\u0000\u0000\u0123\u0128\u0003@ \u0000\u0124\u0125\u0007"+
		"\u0003\u0000\u0000\u0125\u0127\u0003@ \u0000\u0126\u0124\u0001\u0000\u0000"+
		"\u0000\u0127\u012a\u0001\u0000\u0000\u0000\u0128\u0126\u0001\u0000\u0000"+
		"\u0000\u0128\u0129\u0001\u0000\u0000\u0000\u0129?\u0001\u0000\u0000\u0000"+
		"\u012a\u0128\u0001\u0000\u0000\u0000\u012b\u012d\u0007\u0004\u0000\u0000"+
		"\u012c\u012b\u0001\u0000\u0000\u0000\u012c\u012d\u0001\u0000\u0000\u0000"+
		"\u012d\u012e\u0001\u0000\u0000\u0000\u012e\u012f\u0003B!\u0000\u012fA"+
		"\u0001\u0000\u0000\u0000\u0130\u0135\u0003D\"\u0000\u0131\u0132\u0005"+
		"\u001f\u0000\u0000\u0132\u0134\u0005(\u0000\u0000\u0133\u0131\u0001\u0000"+
		"\u0000\u0000\u0134\u0137\u0001\u0000\u0000\u0000\u0135\u0133\u0001\u0000"+
		"\u0000\u0000\u0135\u0136\u0001\u0000\u0000\u0000\u0136C\u0001\u0000\u0000"+
		"\u0000\u0137\u0135\u0001\u0000\u0000\u0000\u0138\u0139\u0005\u001b\u0000"+
		"\u0000\u0139\u013a\u00034\u001a\u0000\u013a\u013b\u0005\u001c\u0000\u0000"+
		"\u013b\u0147\u0001\u0000\u0000\u0000\u013c\u013e\u0005(\u0000\u0000\u013d"+
		"\u013f\u00030\u0018\u0000\u013e\u013d\u0001\u0000\u0000\u0000\u013e\u013f"+
		"\u0001\u0000\u0000\u0000\u013f\u0147\u0001\u0000\u0000\u0000\u0140\u0147"+
		"\u0005)\u0000\u0000\u0141\u0142\u0005\u0012\u0000\u0000\u0142\u0147\u0005"+
		"(\u0000\u0000\u0143\u0147\u0005\u000f\u0000\u0000\u0144\u0147\u0005\u0010"+
		"\u0000\u0000\u0145\u0147\u0005\u0011\u0000\u0000\u0146\u0138\u0001\u0000"+
		"\u0000\u0000\u0146\u013c\u0001\u0000\u0000\u0000\u0146\u0140\u0001\u0000"+
		"\u0000\u0000\u0146\u0141\u0001\u0000\u0000\u0000\u0146\u0143\u0001\u0000"+
		"\u0000\u0000\u0146\u0144\u0001\u0000\u0000\u0000\u0146\u0145\u0001\u0000"+
		"\u0000\u0000\u0147E\u0001\u0000\u0000\u0000\u001bN`jo|\u0082\u0089\u0096"+
		"\u0099\u00a2\u00ad\u00c7\u00d4\u00de\u00ec\u00ef\u00f8\u0100\u0108\u0110"+
		"\u0118\u0120\u0128\u012c\u0135\u013e\u0146";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}