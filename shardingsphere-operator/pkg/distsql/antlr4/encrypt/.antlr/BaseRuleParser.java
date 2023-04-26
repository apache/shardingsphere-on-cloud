// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/encrypt/BaseRule.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BaseRuleParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		AND_=1, OR_=2, NOT_=3, TILDE_=4, VERTICALBAR_=5, AMPERSAND_=6, SIGNEDLEFTSHIFT_=7, 
		SIGNEDRIGHTSHIFT_=8, CARET_=9, MOD_=10, COLON_=11, PLUS_=12, MINUS_=13, 
		ASTERISK_=14, SLASH_=15, BACKSLASH_=16, DOT_=17, DOTASTERISK_=18, SAFEEQ_=19, 
		DEQ_=20, EQ_=21, NEQ_=22, GT_=23, GTE_=24, LT_=25, LTE_=26, POUND_=27, 
		LP_=28, RP_=29, LBE_=30, RBE_=31, LBT_=32, RBT_=33, COMMA_=34, DQ_=35, 
		SQ_=36, BQ_=37, QUESTION_=38, AT_=39, SEMI_=40, JSONSEPARATOR_=41, UL_=42, 
		WS=43, CREATE=44, ALTER=45, DROP=46, SHOW=47, RESOURCE=48, RULE=49, FROM=50, 
		ENCRYPT=51, TYPE=52, ENCRYPT_ALGORITHM=53, ASSISTED_QUERY_ALGORITHM=54, 
		LIKE_QUERY_ALGORITHM=55, NAME=56, PROPERTIES=57, COLUMN=58, RULES=59, 
		TABLE=60, COLUMNS=61, CIPHER=62, PLAIN=63, ASSISTED_QUERY_COLUMN=64, LIKE_QUERY_COLUMN=65, 
		QUERY_WITH_CIPHER_COLUMN=66, TRUE=67, FALSE=68, DATA_TYPE=69, PLAIN_DATA_TYPE=70, 
		CIPHER_DATA_TYPE=71, ASSISTED_QUERY_DATA_TYPE=72, LIKE_QUERY_DATA_TYPE=73, 
		IF=74, EXISTS=75, COUNT=76, MD5=77, AES=78, RC4=79, SM3=80, SM4=81, CHAR_DIGEST_LIKE=82, 
		NOT=83, FOR_GENERATOR=84, IDENTIFIER_=85, STRING_=86, INT_=87, HEX_=88, 
		NUMBER_=89, HEXDIGIT_=90, BITNUM_=91;
	public static final int
		RULE_literal = 0, RULE_algorithmDefinition = 1, RULE_algorithmTypeName = 2, 
		RULE_buildinAlgorithmTypeName = 3, RULE_propertiesDefinition = 4, RULE_properties = 5, 
		RULE_property = 6, RULE_tableName = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"literal", "algorithmDefinition", "algorithmTypeName", "buildinAlgorithmTypeName", 
			"propertiesDefinition", "properties", "property", "tableName"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'&&'", "'||'", "'!'", "'~'", "'|'", "'&'", "'<<'", "'>>'", "'^'", 
			"'%'", "':'", "'+'", "'-'", "'*'", "'/'", "'\\'", "'.'", "'.*'", "'<=>'", 
			"'=='", "'='", null, "'>'", "'>='", "'<'", "'<='", "'#'", "'('", "')'", 
			"'{'", "'}'", "'['", "']'", "','", "'\"'", "'''", "'`'", "'?'", "'@'", 
			"';'", "'->>'", "'_'", null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "AND_", "OR_", "NOT_", "TILDE_", "VERTICALBAR_", "AMPERSAND_", 
			"SIGNEDLEFTSHIFT_", "SIGNEDRIGHTSHIFT_", "CARET_", "MOD_", "COLON_", 
			"PLUS_", "MINUS_", "ASTERISK_", "SLASH_", "BACKSLASH_", "DOT_", "DOTASTERISK_", 
			"SAFEEQ_", "DEQ_", "EQ_", "NEQ_", "GT_", "GTE_", "LT_", "LTE_", "POUND_", 
			"LP_", "RP_", "LBE_", "RBE_", "LBT_", "RBT_", "COMMA_", "DQ_", "SQ_", 
			"BQ_", "QUESTION_", "AT_", "SEMI_", "JSONSEPARATOR_", "UL_", "WS", "CREATE", 
			"ALTER", "DROP", "SHOW", "RESOURCE", "RULE", "FROM", "ENCRYPT", "TYPE", 
			"ENCRYPT_ALGORITHM", "ASSISTED_QUERY_ALGORITHM", "LIKE_QUERY_ALGORITHM", 
			"NAME", "PROPERTIES", "COLUMN", "RULES", "TABLE", "COLUMNS", "CIPHER", 
			"PLAIN", "ASSISTED_QUERY_COLUMN", "LIKE_QUERY_COLUMN", "QUERY_WITH_CIPHER_COLUMN", 
			"TRUE", "FALSE", "DATA_TYPE", "PLAIN_DATA_TYPE", "CIPHER_DATA_TYPE", 
			"ASSISTED_QUERY_DATA_TYPE", "LIKE_QUERY_DATA_TYPE", "IF", "EXISTS", "COUNT", 
			"MD5", "AES", "RC4", "SM3", "SM4", "CHAR_DIGEST_LIKE", "NOT", "FOR_GENERATOR", 
			"IDENTIFIER_", "STRING_", "INT_", "HEX_", "NUMBER_", "HEXDIGIT_", "BITNUM_"
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
	public String getGrammarFileName() { return "BaseRule.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public BaseRuleParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode STRING_() { return getToken(BaseRuleParser.STRING_, 0); }
		public TerminalNode INT_() { return getToken(BaseRuleParser.INT_, 0); }
		public TerminalNode MINUS_() { return getToken(BaseRuleParser.MINUS_, 0); }
		public TerminalNode TRUE() { return getToken(BaseRuleParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(BaseRuleParser.FALSE, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_literal);
		int _la;
		try {
			setState(23);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(16);
				match(STRING_);
				}
				break;
			case MINUS_:
			case INT_:
				enterOuterAlt(_localctx, 2);
				{
				setState(18);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS_) {
					{
					setState(17);
					match(MINUS_);
					}
				}

				setState(20);
				match(INT_);
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(21);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(22);
				match(FALSE);
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

	public static class AlgorithmDefinitionContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(BaseRuleParser.TYPE, 0); }
		public TerminalNode LP_() { return getToken(BaseRuleParser.LP_, 0); }
		public TerminalNode NAME() { return getToken(BaseRuleParser.NAME, 0); }
		public TerminalNode EQ_() { return getToken(BaseRuleParser.EQ_, 0); }
		public AlgorithmTypeNameContext algorithmTypeName() {
			return getRuleContext(AlgorithmTypeNameContext.class,0);
		}
		public TerminalNode RP_() { return getToken(BaseRuleParser.RP_, 0); }
		public TerminalNode COMMA_() { return getToken(BaseRuleParser.COMMA_, 0); }
		public PropertiesDefinitionContext propertiesDefinition() {
			return getRuleContext(PropertiesDefinitionContext.class,0);
		}
		public AlgorithmDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_algorithmDefinition; }
	}

	public final AlgorithmDefinitionContext algorithmDefinition() throws RecognitionException {
		AlgorithmDefinitionContext _localctx = new AlgorithmDefinitionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_algorithmDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(TYPE);
			setState(26);
			match(LP_);
			setState(27);
			match(NAME);
			setState(28);
			match(EQ_);
			setState(29);
			algorithmTypeName();
			setState(32);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA_) {
				{
				setState(30);
				match(COMMA_);
				setState(31);
				propertiesDefinition();
				}
			}

			setState(34);
			match(RP_);
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

	public static class AlgorithmTypeNameContext extends ParserRuleContext {
		public BuildinAlgorithmTypeNameContext buildinAlgorithmTypeName() {
			return getRuleContext(BuildinAlgorithmTypeNameContext.class,0);
		}
		public TerminalNode STRING_() { return getToken(BaseRuleParser.STRING_, 0); }
		public AlgorithmTypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_algorithmTypeName; }
	}

	public final AlgorithmTypeNameContext algorithmTypeName() throws RecognitionException {
		AlgorithmTypeNameContext _localctx = new AlgorithmTypeNameContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_algorithmTypeName);
		try {
			setState(38);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MD5:
			case AES:
			case RC4:
			case SM3:
			case SM4:
			case CHAR_DIGEST_LIKE:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				buildinAlgorithmTypeName();
				}
				break;
			case STRING_:
				enterOuterAlt(_localctx, 2);
				{
				setState(37);
				match(STRING_);
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

	public static class BuildinAlgorithmTypeNameContext extends ParserRuleContext {
		public TerminalNode MD5() { return getToken(BaseRuleParser.MD5, 0); }
		public TerminalNode AES() { return getToken(BaseRuleParser.AES, 0); }
		public TerminalNode RC4() { return getToken(BaseRuleParser.RC4, 0); }
		public TerminalNode SM3() { return getToken(BaseRuleParser.SM3, 0); }
		public TerminalNode SM4() { return getToken(BaseRuleParser.SM4, 0); }
		public TerminalNode CHAR_DIGEST_LIKE() { return getToken(BaseRuleParser.CHAR_DIGEST_LIKE, 0); }
		public BuildinAlgorithmTypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_buildinAlgorithmTypeName; }
	}

	public final BuildinAlgorithmTypeNameContext buildinAlgorithmTypeName() throws RecognitionException {
		BuildinAlgorithmTypeNameContext _localctx = new BuildinAlgorithmTypeNameContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_buildinAlgorithmTypeName);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			_la = _input.LA(1);
			if ( !(((((_la - 77)) & ~0x3f) == 0 && ((1L << (_la - 77)) & ((1L << (MD5 - 77)) | (1L << (AES - 77)) | (1L << (RC4 - 77)) | (1L << (SM3 - 77)) | (1L << (SM4 - 77)) | (1L << (CHAR_DIGEST_LIKE - 77)))) != 0)) ) {
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

	public static class PropertiesDefinitionContext extends ParserRuleContext {
		public TerminalNode PROPERTIES() { return getToken(BaseRuleParser.PROPERTIES, 0); }
		public TerminalNode LP_() { return getToken(BaseRuleParser.LP_, 0); }
		public TerminalNode RP_() { return getToken(BaseRuleParser.RP_, 0); }
		public PropertiesContext properties() {
			return getRuleContext(PropertiesContext.class,0);
		}
		public PropertiesDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertiesDefinition; }
	}

	public final PropertiesDefinitionContext propertiesDefinition() throws RecognitionException {
		PropertiesDefinitionContext _localctx = new PropertiesDefinitionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_propertiesDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42);
			match(PROPERTIES);
			setState(43);
			match(LP_);
			setState(45);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING_) {
				{
				setState(44);
				properties();
				}
			}

			setState(47);
			match(RP_);
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

	public static class PropertiesContext extends ParserRuleContext {
		public List<PropertyContext> property() {
			return getRuleContexts(PropertyContext.class);
		}
		public PropertyContext property(int i) {
			return getRuleContext(PropertyContext.class,i);
		}
		public List<TerminalNode> COMMA_() { return getTokens(BaseRuleParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(BaseRuleParser.COMMA_, i);
		}
		public PropertiesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_properties; }
	}

	public final PropertiesContext properties() throws RecognitionException {
		PropertiesContext _localctx = new PropertiesContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_properties);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(49);
			property();
			setState(54);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(50);
				match(COMMA_);
				setState(51);
				property();
				}
				}
				setState(56);
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

	public static class PropertyContext extends ParserRuleContext {
		public Token key;
		public LiteralContext value;
		public TerminalNode EQ_() { return getToken(BaseRuleParser.EQ_, 0); }
		public TerminalNode STRING_() { return getToken(BaseRuleParser.STRING_, 0); }
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public PropertyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_property; }
	}

	public final PropertyContext property() throws RecognitionException {
		PropertyContext _localctx = new PropertyContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_property);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			((PropertyContext)_localctx).key = match(STRING_);
			setState(58);
			match(EQ_);
			setState(59);
			((PropertyContext)_localctx).value = literal();
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

	public static class TableNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(BaseRuleParser.IDENTIFIER_, 0); }
		public TableNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableName; }
	}

	public final TableNameContext tableName() throws RecognitionException {
		TableNameContext _localctx = new TableNameContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_tableName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			match(IDENTIFIER_);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3]B\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\5\2\25\n\2\3"+
		"\2\3\2\3\2\5\2\32\n\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3#\n\3\3\3\3\3\3\4"+
		"\3\4\5\4)\n\4\3\5\3\5\3\6\3\6\3\6\5\6\60\n\6\3\6\3\6\3\7\3\7\3\7\7\7\67"+
		"\n\7\f\7\16\7:\13\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\2\2\n\2\4\6\b\n\f\16\20"+
		"\2\3\3\2OT\2A\2\31\3\2\2\2\4\33\3\2\2\2\6(\3\2\2\2\b*\3\2\2\2\n,\3\2\2"+
		"\2\f\63\3\2\2\2\16;\3\2\2\2\20?\3\2\2\2\22\32\7X\2\2\23\25\7\17\2\2\24"+
		"\23\3\2\2\2\24\25\3\2\2\2\25\26\3\2\2\2\26\32\7Y\2\2\27\32\7E\2\2\30\32"+
		"\7F\2\2\31\22\3\2\2\2\31\24\3\2\2\2\31\27\3\2\2\2\31\30\3\2\2\2\32\3\3"+
		"\2\2\2\33\34\7\66\2\2\34\35\7\36\2\2\35\36\7:\2\2\36\37\7\27\2\2\37\""+
		"\5\6\4\2 !\7$\2\2!#\5\n\6\2\" \3\2\2\2\"#\3\2\2\2#$\3\2\2\2$%\7\37\2\2"+
		"%\5\3\2\2\2&)\5\b\5\2\')\7X\2\2(&\3\2\2\2(\'\3\2\2\2)\7\3\2\2\2*+\t\2"+
		"\2\2+\t\3\2\2\2,-\7;\2\2-/\7\36\2\2.\60\5\f\7\2/.\3\2\2\2/\60\3\2\2\2"+
		"\60\61\3\2\2\2\61\62\7\37\2\2\62\13\3\2\2\2\638\5\16\b\2\64\65\7$\2\2"+
		"\65\67\5\16\b\2\66\64\3\2\2\2\67:\3\2\2\28\66\3\2\2\289\3\2\2\29\r\3\2"+
		"\2\2:8\3\2\2\2;<\7X\2\2<=\7\27\2\2=>\5\2\2\2>\17\3\2\2\2?@\7W\2\2@\21"+
		"\3\2\2\2\b\24\31\"(/8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}