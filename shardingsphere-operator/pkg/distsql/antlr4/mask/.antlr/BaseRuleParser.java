// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/mask/BaseRule.g4 by ANTLR 4.9.2
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
		WS=43, TRUE=44, FALSE=45, CREATE=46, ALTER=47, DROP=48, SHOW=49, RULE=50, 
		FROM=51, MASK=52, TYPE=53, NAME=54, PROPERTIES=55, COLUMN=56, RULES=57, 
		TABLE=58, COLUMNS=59, IF=60, EXISTS=61, COUNT=62, NOT=63, MD5=64, KEEP_FIRST_N_LAST_M=65, 
		KEEP_FROM_X_TO_Y=66, MASK_FIRST_N_LAST_M=67, MASK_FROM_X_TO_Y=68, MASK_BEFORE_SPECIAL_CHARS=69, 
		MASK_AFTER_SPECIAL_CHARS=70, PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE=71, 
		MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE=72, LANDLINE_NUMBER_RANDOM_REPLACE=73, 
		TELEPHONE_RANDOM_REPLACE=74, UNIFIED_CREDIT_CODE_RANDOM_REPLACE=75, GENERIC_TABLE_RANDOM_REPLACE=76, 
		ADDRESS_RANDOM_REPLACE=77, FOR_GENERATOR=78, IDENTIFIER_=79, STRING_=80, 
		INT_=81, HEX_=82, NUMBER_=83, HEXDIGIT_=84, BITNUM_=85;
	public static final int
		RULE_literal = 0, RULE_algorithmDefinition = 1, RULE_algorithmTypeName = 2, 
		RULE_buildInMaskAlgorithmType = 3, RULE_propertiesDefinition = 4, RULE_properties = 5, 
		RULE_property = 6, RULE_ruleName = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"literal", "algorithmDefinition", "algorithmTypeName", "buildInMaskAlgorithmType", 
			"propertiesDefinition", "properties", "property", "ruleName"
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
			null, null, null, "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'"
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
			"BQ_", "QUESTION_", "AT_", "SEMI_", "JSONSEPARATOR_", "UL_", "WS", "TRUE", 
			"FALSE", "CREATE", "ALTER", "DROP", "SHOW", "RULE", "FROM", "MASK", "TYPE", 
			"NAME", "PROPERTIES", "COLUMN", "RULES", "TABLE", "COLUMNS", "IF", "EXISTS", 
			"COUNT", "NOT", "MD5", "KEEP_FIRST_N_LAST_M", "KEEP_FROM_X_TO_Y", "MASK_FIRST_N_LAST_M", 
			"MASK_FROM_X_TO_Y", "MASK_BEFORE_SPECIAL_CHARS", "MASK_AFTER_SPECIAL_CHARS", 
			"PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE", "MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE", 
			"LANDLINE_NUMBER_RANDOM_REPLACE", "TELEPHONE_RANDOM_REPLACE", "UNIFIED_CREDIT_CODE_RANDOM_REPLACE", 
			"GENERIC_TABLE_RANDOM_REPLACE", "ADDRESS_RANDOM_REPLACE", "FOR_GENERATOR", 
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
		public TerminalNode STRING_() { return getToken(BaseRuleParser.STRING_, 0); }
		public BuildInMaskAlgorithmTypeContext buildInMaskAlgorithmType() {
			return getRuleContext(BuildInMaskAlgorithmTypeContext.class,0);
		}
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
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				match(STRING_);
				}
				break;
			case MD5:
			case KEEP_FIRST_N_LAST_M:
			case KEEP_FROM_X_TO_Y:
			case MASK_FIRST_N_LAST_M:
			case MASK_FROM_X_TO_Y:
			case MASK_BEFORE_SPECIAL_CHARS:
			case MASK_AFTER_SPECIAL_CHARS:
			case PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE:
			case MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE:
			case LANDLINE_NUMBER_RANDOM_REPLACE:
			case TELEPHONE_RANDOM_REPLACE:
			case UNIFIED_CREDIT_CODE_RANDOM_REPLACE:
			case GENERIC_TABLE_RANDOM_REPLACE:
				enterOuterAlt(_localctx, 2);
				{
				setState(37);
				buildInMaskAlgorithmType();
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

	public static class BuildInMaskAlgorithmTypeContext extends ParserRuleContext {
		public TerminalNode MD5() { return getToken(BaseRuleParser.MD5, 0); }
		public TerminalNode KEEP_FIRST_N_LAST_M() { return getToken(BaseRuleParser.KEEP_FIRST_N_LAST_M, 0); }
		public TerminalNode KEEP_FROM_X_TO_Y() { return getToken(BaseRuleParser.KEEP_FROM_X_TO_Y, 0); }
		public TerminalNode MASK_FIRST_N_LAST_M() { return getToken(BaseRuleParser.MASK_FIRST_N_LAST_M, 0); }
		public TerminalNode MASK_FROM_X_TO_Y() { return getToken(BaseRuleParser.MASK_FROM_X_TO_Y, 0); }
		public TerminalNode MASK_BEFORE_SPECIAL_CHARS() { return getToken(BaseRuleParser.MASK_BEFORE_SPECIAL_CHARS, 0); }
		public TerminalNode MASK_AFTER_SPECIAL_CHARS() { return getToken(BaseRuleParser.MASK_AFTER_SPECIAL_CHARS, 0); }
		public TerminalNode PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE() { return getToken(BaseRuleParser.PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE() { return getToken(BaseRuleParser.MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode LANDLINE_NUMBER_RANDOM_REPLACE() { return getToken(BaseRuleParser.LANDLINE_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode TELEPHONE_RANDOM_REPLACE() { return getToken(BaseRuleParser.TELEPHONE_RANDOM_REPLACE, 0); }
		public TerminalNode UNIFIED_CREDIT_CODE_RANDOM_REPLACE() { return getToken(BaseRuleParser.UNIFIED_CREDIT_CODE_RANDOM_REPLACE, 0); }
		public TerminalNode GENERIC_TABLE_RANDOM_REPLACE() { return getToken(BaseRuleParser.GENERIC_TABLE_RANDOM_REPLACE, 0); }
		public BuildInMaskAlgorithmTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_buildInMaskAlgorithmType; }
	}

	public final BuildInMaskAlgorithmTypeContext buildInMaskAlgorithmType() throws RecognitionException {
		BuildInMaskAlgorithmTypeContext _localctx = new BuildInMaskAlgorithmTypeContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_buildInMaskAlgorithmType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			_la = _input.LA(1);
			if ( !(((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (MD5 - 64)) | (1L << (KEEP_FIRST_N_LAST_M - 64)) | (1L << (KEEP_FROM_X_TO_Y - 64)) | (1L << (MASK_FIRST_N_LAST_M - 64)) | (1L << (MASK_FROM_X_TO_Y - 64)) | (1L << (MASK_BEFORE_SPECIAL_CHARS - 64)) | (1L << (MASK_AFTER_SPECIAL_CHARS - 64)) | (1L << (PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE - 64)) | (1L << (MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE - 64)) | (1L << (LANDLINE_NUMBER_RANDOM_REPLACE - 64)) | (1L << (TELEPHONE_RANDOM_REPLACE - 64)) | (1L << (UNIFIED_CREDIT_CODE_RANDOM_REPLACE - 64)) | (1L << (GENERIC_TABLE_RANDOM_REPLACE - 64)))) != 0)) ) {
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

	public static class RuleNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(BaseRuleParser.IDENTIFIER_, 0); }
		public RuleNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleName; }
	}

	public final RuleNameContext ruleName() throws RecognitionException {
		RuleNameContext _localctx = new RuleNameContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ruleName);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3WB\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\5\2\25\n\2\3"+
		"\2\3\2\3\2\5\2\32\n\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3#\n\3\3\3\3\3\3\4"+
		"\3\4\5\4)\n\4\3\5\3\5\3\6\3\6\3\6\5\6\60\n\6\3\6\3\6\3\7\3\7\3\7\7\7\67"+
		"\n\7\f\7\16\7:\13\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\2\2\n\2\4\6\b\n\f\16\20"+
		"\2\3\3\2BN\2A\2\31\3\2\2\2\4\33\3\2\2\2\6(\3\2\2\2\b*\3\2\2\2\n,\3\2\2"+
		"\2\f\63\3\2\2\2\16;\3\2\2\2\20?\3\2\2\2\22\32\7R\2\2\23\25\7\17\2\2\24"+
		"\23\3\2\2\2\24\25\3\2\2\2\25\26\3\2\2\2\26\32\7S\2\2\27\32\7.\2\2\30\32"+
		"\7/\2\2\31\22\3\2\2\2\31\24\3\2\2\2\31\27\3\2\2\2\31\30\3\2\2\2\32\3\3"+
		"\2\2\2\33\34\7\67\2\2\34\35\7\36\2\2\35\36\78\2\2\36\37\7\27\2\2\37\""+
		"\5\6\4\2 !\7$\2\2!#\5\n\6\2\" \3\2\2\2\"#\3\2\2\2#$\3\2\2\2$%\7\37\2\2"+
		"%\5\3\2\2\2&)\7R\2\2\')\5\b\5\2(&\3\2\2\2(\'\3\2\2\2)\7\3\2\2\2*+\t\2"+
		"\2\2+\t\3\2\2\2,-\79\2\2-/\7\36\2\2.\60\5\f\7\2/.\3\2\2\2/\60\3\2\2\2"+
		"\60\61\3\2\2\2\61\62\7\37\2\2\62\13\3\2\2\2\638\5\16\b\2\64\65\7$\2\2"+
		"\65\67\5\16\b\2\66\64\3\2\2\2\67:\3\2\2\28\66\3\2\2\289\3\2\2\29\r\3\2"+
		"\2\2:8\3\2\2\2;<\7R\2\2<=\7\27\2\2=>\5\2\2\2>\17\3\2\2\2?@\7Q\2\2@\21"+
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