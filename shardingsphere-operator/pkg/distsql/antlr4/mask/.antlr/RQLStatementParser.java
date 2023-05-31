// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/mask/RQLStatement.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class RQLStatementParser extends Parser {
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
		RULE_showMaskRules = 0, RULE_countMaskRule = 1, RULE_databaseName = 2, 
		RULE_literal = 3, RULE_algorithmDefinition = 4, RULE_algorithmTypeName = 5, 
		RULE_buildInMaskAlgorithmType = 6, RULE_propertiesDefinition = 7, RULE_properties = 8, 
		RULE_property = 9, RULE_ruleName = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"showMaskRules", "countMaskRule", "databaseName", "literal", "algorithmDefinition", 
			"algorithmTypeName", "buildInMaskAlgorithmType", "propertiesDefinition", 
			"properties", "property", "ruleName"
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
	public String getGrammarFileName() { return "RQLStatement.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public RQLStatementParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ShowMaskRulesContext extends ParserRuleContext {
		public TerminalNode SHOW() { return getToken(RQLStatementParser.SHOW, 0); }
		public TerminalNode MASK() { return getToken(RQLStatementParser.MASK, 0); }
		public TerminalNode RULE() { return getToken(RQLStatementParser.RULE, 0); }
		public RuleNameContext ruleName() {
			return getRuleContext(RuleNameContext.class,0);
		}
		public TerminalNode RULES() { return getToken(RQLStatementParser.RULES, 0); }
		public TerminalNode TABLE() { return getToken(RQLStatementParser.TABLE, 0); }
		public TerminalNode FROM() { return getToken(RQLStatementParser.FROM, 0); }
		public DatabaseNameContext databaseName() {
			return getRuleContext(DatabaseNameContext.class,0);
		}
		public ShowMaskRulesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_showMaskRules; }
	}

	public final ShowMaskRulesContext showMaskRules() throws RecognitionException {
		ShowMaskRulesContext _localctx = new ShowMaskRulesContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_showMaskRules);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			match(SHOW);
			setState(23);
			match(MASK);
			setState(25);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TABLE) {
				{
				setState(24);
				match(TABLE);
				}
			}

			setState(30);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RULE:
				{
				setState(27);
				match(RULE);
				setState(28);
				ruleName();
				}
				break;
			case RULES:
				{
				setState(29);
				match(RULES);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(34);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FROM) {
				{
				setState(32);
				match(FROM);
				setState(33);
				databaseName();
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

	public static class CountMaskRuleContext extends ParserRuleContext {
		public TerminalNode COUNT() { return getToken(RQLStatementParser.COUNT, 0); }
		public TerminalNode MASK() { return getToken(RQLStatementParser.MASK, 0); }
		public TerminalNode RULE() { return getToken(RQLStatementParser.RULE, 0); }
		public TerminalNode FROM() { return getToken(RQLStatementParser.FROM, 0); }
		public DatabaseNameContext databaseName() {
			return getRuleContext(DatabaseNameContext.class,0);
		}
		public CountMaskRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_countMaskRule; }
	}

	public final CountMaskRuleContext countMaskRule() throws RecognitionException {
		CountMaskRuleContext _localctx = new CountMaskRuleContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_countMaskRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			match(COUNT);
			setState(37);
			match(MASK);
			setState(38);
			match(RULE);
			setState(41);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FROM) {
				{
				setState(39);
				match(FROM);
				setState(40);
				databaseName();
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

	public static class DatabaseNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RQLStatementParser.IDENTIFIER_, 0); }
		public DatabaseNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_databaseName; }
	}

	public final DatabaseNameContext databaseName() throws RecognitionException {
		DatabaseNameContext _localctx = new DatabaseNameContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_databaseName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
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

	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode STRING_() { return getToken(RQLStatementParser.STRING_, 0); }
		public TerminalNode INT_() { return getToken(RQLStatementParser.INT_, 0); }
		public TerminalNode MINUS_() { return getToken(RQLStatementParser.MINUS_, 0); }
		public TerminalNode TRUE() { return getToken(RQLStatementParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(RQLStatementParser.FALSE, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_literal);
		int _la;
		try {
			setState(52);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(45);
				match(STRING_);
				}
				break;
			case MINUS_:
			case INT_:
				enterOuterAlt(_localctx, 2);
				{
				setState(47);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS_) {
					{
					setState(46);
					match(MINUS_);
					}
				}

				setState(49);
				match(INT_);
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(50);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(51);
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
		public TerminalNode TYPE() { return getToken(RQLStatementParser.TYPE, 0); }
		public TerminalNode LP_() { return getToken(RQLStatementParser.LP_, 0); }
		public TerminalNode NAME() { return getToken(RQLStatementParser.NAME, 0); }
		public TerminalNode EQ_() { return getToken(RQLStatementParser.EQ_, 0); }
		public AlgorithmTypeNameContext algorithmTypeName() {
			return getRuleContext(AlgorithmTypeNameContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RQLStatementParser.RP_, 0); }
		public TerminalNode COMMA_() { return getToken(RQLStatementParser.COMMA_, 0); }
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
		enterRule(_localctx, 8, RULE_algorithmDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(TYPE);
			setState(55);
			match(LP_);
			setState(56);
			match(NAME);
			setState(57);
			match(EQ_);
			setState(58);
			algorithmTypeName();
			setState(61);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA_) {
				{
				setState(59);
				match(COMMA_);
				setState(60);
				propertiesDefinition();
				}
			}

			setState(63);
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
		public TerminalNode STRING_() { return getToken(RQLStatementParser.STRING_, 0); }
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
		enterRule(_localctx, 10, RULE_algorithmTypeName);
		try {
			setState(67);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
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
				setState(66);
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
		public TerminalNode MD5() { return getToken(RQLStatementParser.MD5, 0); }
		public TerminalNode KEEP_FIRST_N_LAST_M() { return getToken(RQLStatementParser.KEEP_FIRST_N_LAST_M, 0); }
		public TerminalNode KEEP_FROM_X_TO_Y() { return getToken(RQLStatementParser.KEEP_FROM_X_TO_Y, 0); }
		public TerminalNode MASK_FIRST_N_LAST_M() { return getToken(RQLStatementParser.MASK_FIRST_N_LAST_M, 0); }
		public TerminalNode MASK_FROM_X_TO_Y() { return getToken(RQLStatementParser.MASK_FROM_X_TO_Y, 0); }
		public TerminalNode MASK_BEFORE_SPECIAL_CHARS() { return getToken(RQLStatementParser.MASK_BEFORE_SPECIAL_CHARS, 0); }
		public TerminalNode MASK_AFTER_SPECIAL_CHARS() { return getToken(RQLStatementParser.MASK_AFTER_SPECIAL_CHARS, 0); }
		public TerminalNode PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE() { return getToken(RQLStatementParser.PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE() { return getToken(RQLStatementParser.MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode LANDLINE_NUMBER_RANDOM_REPLACE() { return getToken(RQLStatementParser.LANDLINE_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode TELEPHONE_RANDOM_REPLACE() { return getToken(RQLStatementParser.TELEPHONE_RANDOM_REPLACE, 0); }
		public TerminalNode UNIFIED_CREDIT_CODE_RANDOM_REPLACE() { return getToken(RQLStatementParser.UNIFIED_CREDIT_CODE_RANDOM_REPLACE, 0); }
		public TerminalNode GENERIC_TABLE_RANDOM_REPLACE() { return getToken(RQLStatementParser.GENERIC_TABLE_RANDOM_REPLACE, 0); }
		public BuildInMaskAlgorithmTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_buildInMaskAlgorithmType; }
	}

	public final BuildInMaskAlgorithmTypeContext buildInMaskAlgorithmType() throws RecognitionException {
		BuildInMaskAlgorithmTypeContext _localctx = new BuildInMaskAlgorithmTypeContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_buildInMaskAlgorithmType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
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
		public TerminalNode PROPERTIES() { return getToken(RQLStatementParser.PROPERTIES, 0); }
		public TerminalNode LP_() { return getToken(RQLStatementParser.LP_, 0); }
		public TerminalNode RP_() { return getToken(RQLStatementParser.RP_, 0); }
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
		enterRule(_localctx, 14, RULE_propertiesDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(PROPERTIES);
			setState(72);
			match(LP_);
			setState(74);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING_) {
				{
				setState(73);
				properties();
				}
			}

			setState(76);
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
		public List<TerminalNode> COMMA_() { return getTokens(RQLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RQLStatementParser.COMMA_, i);
		}
		public PropertiesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_properties; }
	}

	public final PropertiesContext properties() throws RecognitionException {
		PropertiesContext _localctx = new PropertiesContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_properties);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			property();
			setState(83);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(79);
				match(COMMA_);
				setState(80);
				property();
				}
				}
				setState(85);
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
		public TerminalNode EQ_() { return getToken(RQLStatementParser.EQ_, 0); }
		public TerminalNode STRING_() { return getToken(RQLStatementParser.STRING_, 0); }
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
		enterRule(_localctx, 18, RULE_property);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			((PropertyContext)_localctx).key = match(STRING_);
			setState(87);
			match(EQ_);
			setState(88);
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
		public TerminalNode IDENTIFIER_() { return getToken(RQLStatementParser.IDENTIFIER_, 0); }
		public RuleNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleName; }
	}

	public final RuleNameContext ruleName() throws RecognitionException {
		RuleNameContext _localctx = new RuleNameContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_ruleName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3W_\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4\f"+
		"\t\f\3\2\3\2\3\2\5\2\34\n\2\3\2\3\2\3\2\5\2!\n\2\3\2\3\2\5\2%\n\2\3\3"+
		"\3\3\3\3\3\3\3\3\5\3,\n\3\3\4\3\4\3\5\3\5\5\5\62\n\5\3\5\3\5\3\5\5\5\67"+
		"\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6@\n\6\3\6\3\6\3\7\3\7\5\7F\n\7\3\b"+
		"\3\b\3\t\3\t\3\t\5\tM\n\t\3\t\3\t\3\n\3\n\3\n\7\nT\n\n\f\n\16\nW\13\n"+
		"\3\13\3\13\3\13\3\13\3\f\3\f\3\f\2\2\r\2\4\6\b\n\f\16\20\22\24\26\2\3"+
		"\3\2BN\2_\2\30\3\2\2\2\4&\3\2\2\2\6-\3\2\2\2\b\66\3\2\2\2\n8\3\2\2\2\f"+
		"E\3\2\2\2\16G\3\2\2\2\20I\3\2\2\2\22P\3\2\2\2\24X\3\2\2\2\26\\\3\2\2\2"+
		"\30\31\7\63\2\2\31\33\7\66\2\2\32\34\7<\2\2\33\32\3\2\2\2\33\34\3\2\2"+
		"\2\34 \3\2\2\2\35\36\7\64\2\2\36!\5\26\f\2\37!\7;\2\2 \35\3\2\2\2 \37"+
		"\3\2\2\2!$\3\2\2\2\"#\7\65\2\2#%\5\6\4\2$\"\3\2\2\2$%\3\2\2\2%\3\3\2\2"+
		"\2&\'\7@\2\2\'(\7\66\2\2(+\7\64\2\2)*\7\65\2\2*,\5\6\4\2+)\3\2\2\2+,\3"+
		"\2\2\2,\5\3\2\2\2-.\7Q\2\2.\7\3\2\2\2/\67\7R\2\2\60\62\7\17\2\2\61\60"+
		"\3\2\2\2\61\62\3\2\2\2\62\63\3\2\2\2\63\67\7S\2\2\64\67\7.\2\2\65\67\7"+
		"/\2\2\66/\3\2\2\2\66\61\3\2\2\2\66\64\3\2\2\2\66\65\3\2\2\2\67\t\3\2\2"+
		"\289\7\67\2\29:\7\36\2\2:;\78\2\2;<\7\27\2\2<?\5\f\7\2=>\7$\2\2>@\5\20"+
		"\t\2?=\3\2\2\2?@\3\2\2\2@A\3\2\2\2AB\7\37\2\2B\13\3\2\2\2CF\7R\2\2DF\5"+
		"\16\b\2EC\3\2\2\2ED\3\2\2\2F\r\3\2\2\2GH\t\2\2\2H\17\3\2\2\2IJ\79\2\2"+
		"JL\7\36\2\2KM\5\22\n\2LK\3\2\2\2LM\3\2\2\2MN\3\2\2\2NO\7\37\2\2O\21\3"+
		"\2\2\2PU\5\24\13\2QR\7$\2\2RT\5\24\13\2SQ\3\2\2\2TW\3\2\2\2US\3\2\2\2"+
		"UV\3\2\2\2V\23\3\2\2\2WU\3\2\2\2XY\7R\2\2YZ\7\27\2\2Z[\5\b\5\2[\25\3\2"+
		"\2\2\\]\7Q\2\2]\27\3\2\2\2\f\33 $+\61\66?ELU";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}