// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/read_write_splitting/RALStatement.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class RALStatementParser extends Parser {
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
		FROM=51, READWRITE_SPLITTING=52, WRITE_STORAGE_UNIT=53, READ_STORAGE_UNITS=54, 
		TRANSACTIONAL_READ_QUERY_STRATEGY=55, TYPE=56, NAME=57, PROPERTIES=58, 
		RULES=59, RESOURCES=60, STATUS=61, ENABLE=62, DISABLE=63, READ=64, IF=65, 
		EXISTS=66, COUNT=67, ROUND_ROBIN=68, RANDOM=69, WEIGHT=70, NOT=71, FOR_GENERATOR=72, 
		IDENTIFIER_=73, STRING_=74, INT_=75, HEX_=76, NUMBER_=77, HEXDIGIT_=78, 
		BITNUM_=79;
	public static final int
		RULE_alterReadwriteSplittingStorageUnitStatus = 0, RULE_showStatusFromReadwriteSplittingRules = 1, 
		RULE_literal = 2, RULE_algorithmDefinition = 3, RULE_algorithmTypeName = 4, 
		RULE_buildInReadQueryLoadBalanceAlgorithmType = 5, RULE_propertiesDefinition = 6, 
		RULE_properties = 7, RULE_property = 8, RULE_databaseName = 9, RULE_groupName = 10, 
		RULE_storageUnitName = 11;
	private static String[] makeRuleNames() {
		return new String[] {
			"alterReadwriteSplittingStorageUnitStatus", "showStatusFromReadwriteSplittingRules", 
			"literal", "algorithmDefinition", "algorithmTypeName", "buildInReadQueryLoadBalanceAlgorithmType", 
			"propertiesDefinition", "properties", "property", "databaseName", "groupName", 
			"storageUnitName"
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
			"BQ_", "QUESTION_", "AT_", "SEMI_", "JSONSEPARATOR_", "UL_", "WS", "TRUE", 
			"FALSE", "CREATE", "ALTER", "DROP", "SHOW", "RULE", "FROM", "READWRITE_SPLITTING", 
			"WRITE_STORAGE_UNIT", "READ_STORAGE_UNITS", "TRANSACTIONAL_READ_QUERY_STRATEGY", 
			"TYPE", "NAME", "PROPERTIES", "RULES", "RESOURCES", "STATUS", "ENABLE", 
			"DISABLE", "READ", "IF", "EXISTS", "COUNT", "ROUND_ROBIN", "RANDOM", 
			"WEIGHT", "NOT", "FOR_GENERATOR", "IDENTIFIER_", "STRING_", "INT_", "HEX_", 
			"NUMBER_", "HEXDIGIT_", "BITNUM_"
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
	public String getGrammarFileName() { return "RALStatement.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public RALStatementParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class AlterReadwriteSplittingStorageUnitStatusContext extends ParserRuleContext {
		public TerminalNode ALTER() { return getToken(RALStatementParser.ALTER, 0); }
		public TerminalNode READWRITE_SPLITTING() { return getToken(RALStatementParser.READWRITE_SPLITTING, 0); }
		public TerminalNode RULE() { return getToken(RALStatementParser.RULE, 0); }
		public StorageUnitNameContext storageUnitName() {
			return getRuleContext(StorageUnitNameContext.class,0);
		}
		public TerminalNode ENABLE() { return getToken(RALStatementParser.ENABLE, 0); }
		public TerminalNode DISABLE() { return getToken(RALStatementParser.DISABLE, 0); }
		public GroupNameContext groupName() {
			return getRuleContext(GroupNameContext.class,0);
		}
		public TerminalNode FROM() { return getToken(RALStatementParser.FROM, 0); }
		public DatabaseNameContext databaseName() {
			return getRuleContext(DatabaseNameContext.class,0);
		}
		public AlterReadwriteSplittingStorageUnitStatusContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_alterReadwriteSplittingStorageUnitStatus; }
	}

	public final AlterReadwriteSplittingStorageUnitStatusContext alterReadwriteSplittingStorageUnitStatus() throws RecognitionException {
		AlterReadwriteSplittingStorageUnitStatusContext _localctx = new AlterReadwriteSplittingStorageUnitStatusContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_alterReadwriteSplittingStorageUnitStatus);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			match(ALTER);
			setState(25);
			match(READWRITE_SPLITTING);
			setState(26);
			match(RULE);
			setState(28);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER_) {
				{
				setState(27);
				groupName();
				}
			}

			setState(30);
			_la = _input.LA(1);
			if ( !(_la==ENABLE || _la==DISABLE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(31);
			storageUnitName();
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

	public static class ShowStatusFromReadwriteSplittingRulesContext extends ParserRuleContext {
		public TerminalNode SHOW() { return getToken(RALStatementParser.SHOW, 0); }
		public TerminalNode STATUS() { return getToken(RALStatementParser.STATUS, 0); }
		public List<TerminalNode> FROM() { return getTokens(RALStatementParser.FROM); }
		public TerminalNode FROM(int i) {
			return getToken(RALStatementParser.FROM, i);
		}
		public TerminalNode READWRITE_SPLITTING() { return getToken(RALStatementParser.READWRITE_SPLITTING, 0); }
		public TerminalNode RULES() { return getToken(RALStatementParser.RULES, 0); }
		public TerminalNode RULE() { return getToken(RALStatementParser.RULE, 0); }
		public GroupNameContext groupName() {
			return getRuleContext(GroupNameContext.class,0);
		}
		public DatabaseNameContext databaseName() {
			return getRuleContext(DatabaseNameContext.class,0);
		}
		public ShowStatusFromReadwriteSplittingRulesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_showStatusFromReadwriteSplittingRules; }
	}

	public final ShowStatusFromReadwriteSplittingRulesContext showStatusFromReadwriteSplittingRules() throws RecognitionException {
		ShowStatusFromReadwriteSplittingRulesContext _localctx = new ShowStatusFromReadwriteSplittingRulesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_showStatusFromReadwriteSplittingRules);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			match(SHOW);
			setState(37);
			match(STATUS);
			setState(38);
			match(FROM);
			setState(39);
			match(READWRITE_SPLITTING);
			setState(43);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RULES:
				{
				setState(40);
				match(RULES);
				}
				break;
			case RULE:
				{
				setState(41);
				match(RULE);
				setState(42);
				groupName();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(47);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FROM) {
				{
				setState(45);
				match(FROM);
				setState(46);
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

	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode STRING_() { return getToken(RALStatementParser.STRING_, 0); }
		public TerminalNode INT_() { return getToken(RALStatementParser.INT_, 0); }
		public TerminalNode MINUS_() { return getToken(RALStatementParser.MINUS_, 0); }
		public TerminalNode TRUE() { return getToken(RALStatementParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(RALStatementParser.FALSE, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_literal);
		int _la;
		try {
			setState(56);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(49);
				match(STRING_);
				}
				break;
			case MINUS_:
			case INT_:
				enterOuterAlt(_localctx, 2);
				{
				setState(51);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS_) {
					{
					setState(50);
					match(MINUS_);
					}
				}

				setState(53);
				match(INT_);
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(54);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(55);
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
		public TerminalNode TYPE() { return getToken(RALStatementParser.TYPE, 0); }
		public TerminalNode LP_() { return getToken(RALStatementParser.LP_, 0); }
		public TerminalNode NAME() { return getToken(RALStatementParser.NAME, 0); }
		public TerminalNode EQ_() { return getToken(RALStatementParser.EQ_, 0); }
		public AlgorithmTypeNameContext algorithmTypeName() {
			return getRuleContext(AlgorithmTypeNameContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RALStatementParser.RP_, 0); }
		public TerminalNode COMMA_() { return getToken(RALStatementParser.COMMA_, 0); }
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
		enterRule(_localctx, 6, RULE_algorithmDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			match(TYPE);
			setState(59);
			match(LP_);
			setState(60);
			match(NAME);
			setState(61);
			match(EQ_);
			setState(62);
			algorithmTypeName();
			setState(65);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA_) {
				{
				setState(63);
				match(COMMA_);
				setState(64);
				propertiesDefinition();
				}
			}

			setState(67);
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
		public TerminalNode STRING_() { return getToken(RALStatementParser.STRING_, 0); }
		public BuildInReadQueryLoadBalanceAlgorithmTypeContext buildInReadQueryLoadBalanceAlgorithmType() {
			return getRuleContext(BuildInReadQueryLoadBalanceAlgorithmTypeContext.class,0);
		}
		public AlgorithmTypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_algorithmTypeName; }
	}

	public final AlgorithmTypeNameContext algorithmTypeName() throws RecognitionException {
		AlgorithmTypeNameContext _localctx = new AlgorithmTypeNameContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_algorithmTypeName);
		try {
			setState(71);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(69);
				match(STRING_);
				}
				break;
			case ROUND_ROBIN:
			case RANDOM:
			case WEIGHT:
				enterOuterAlt(_localctx, 2);
				{
				setState(70);
				buildInReadQueryLoadBalanceAlgorithmType();
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

	public static class BuildInReadQueryLoadBalanceAlgorithmTypeContext extends ParserRuleContext {
		public TerminalNode ROUND_ROBIN() { return getToken(RALStatementParser.ROUND_ROBIN, 0); }
		public TerminalNode RANDOM() { return getToken(RALStatementParser.RANDOM, 0); }
		public TerminalNode WEIGHT() { return getToken(RALStatementParser.WEIGHT, 0); }
		public BuildInReadQueryLoadBalanceAlgorithmTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_buildInReadQueryLoadBalanceAlgorithmType; }
	}

	public final BuildInReadQueryLoadBalanceAlgorithmTypeContext buildInReadQueryLoadBalanceAlgorithmType() throws RecognitionException {
		BuildInReadQueryLoadBalanceAlgorithmTypeContext _localctx = new BuildInReadQueryLoadBalanceAlgorithmTypeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_buildInReadQueryLoadBalanceAlgorithmType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			_la = _input.LA(1);
			if ( !(((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & ((1L << (ROUND_ROBIN - 68)) | (1L << (RANDOM - 68)) | (1L << (WEIGHT - 68)))) != 0)) ) {
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
		public TerminalNode PROPERTIES() { return getToken(RALStatementParser.PROPERTIES, 0); }
		public TerminalNode LP_() { return getToken(RALStatementParser.LP_, 0); }
		public TerminalNode RP_() { return getToken(RALStatementParser.RP_, 0); }
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
		enterRule(_localctx, 12, RULE_propertiesDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			match(PROPERTIES);
			setState(76);
			match(LP_);
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING_) {
				{
				setState(77);
				properties();
				}
			}

			setState(80);
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
		public List<TerminalNode> COMMA_() { return getTokens(RALStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RALStatementParser.COMMA_, i);
		}
		public PropertiesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_properties; }
	}

	public final PropertiesContext properties() throws RecognitionException {
		PropertiesContext _localctx = new PropertiesContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_properties);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			property();
			setState(87);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(83);
				match(COMMA_);
				setState(84);
				property();
				}
				}
				setState(89);
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
		public TerminalNode EQ_() { return getToken(RALStatementParser.EQ_, 0); }
		public TerminalNode STRING_() { return getToken(RALStatementParser.STRING_, 0); }
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
		enterRule(_localctx, 16, RULE_property);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			((PropertyContext)_localctx).key = match(STRING_);
			setState(91);
			match(EQ_);
			setState(92);
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

	public static class DatabaseNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RALStatementParser.IDENTIFIER_, 0); }
		public DatabaseNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_databaseName; }
	}

	public final DatabaseNameContext databaseName() throws RecognitionException {
		DatabaseNameContext _localctx = new DatabaseNameContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_databaseName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
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

	public static class GroupNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RALStatementParser.IDENTIFIER_, 0); }
		public GroupNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_groupName; }
	}

	public final GroupNameContext groupName() throws RecognitionException {
		GroupNameContext _localctx = new GroupNameContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_groupName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
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

	public static class StorageUnitNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RALStatementParser.IDENTIFIER_, 0); }
		public StorageUnitNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_storageUnitName; }
	}

	public final StorageUnitNameContext storageUnitName() throws RecognitionException {
		StorageUnitNameContext _localctx = new StorageUnitNameContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_storageUnitName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3Qg\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4\f"+
		"\t\f\4\r\t\r\3\2\3\2\3\2\3\2\5\2\37\n\2\3\2\3\2\3\2\3\2\5\2%\n\2\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\5\3.\n\3\3\3\3\3\5\3\62\n\3\3\4\3\4\5\4\66\n\4"+
		"\3\4\3\4\3\4\5\4;\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5D\n\5\3\5\3\5\3\6"+
		"\3\6\5\6J\n\6\3\7\3\7\3\b\3\b\3\b\5\bQ\n\b\3\b\3\b\3\t\3\t\3\t\7\tX\n"+
		"\t\f\t\16\t[\13\t\3\n\3\n\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\2\2\16"+
		"\2\4\6\b\n\f\16\20\22\24\26\30\2\4\3\2@A\3\2FH\2f\2\32\3\2\2\2\4&\3\2"+
		"\2\2\6:\3\2\2\2\b<\3\2\2\2\nI\3\2\2\2\fK\3\2\2\2\16M\3\2\2\2\20T\3\2\2"+
		"\2\22\\\3\2\2\2\24`\3\2\2\2\26b\3\2\2\2\30d\3\2\2\2\32\33\7\61\2\2\33"+
		"\34\7\66\2\2\34\36\7\64\2\2\35\37\5\26\f\2\36\35\3\2\2\2\36\37\3\2\2\2"+
		"\37 \3\2\2\2 !\t\2\2\2!$\5\30\r\2\"#\7\65\2\2#%\5\24\13\2$\"\3\2\2\2$"+
		"%\3\2\2\2%\3\3\2\2\2&\'\7\63\2\2\'(\7?\2\2()\7\65\2\2)-\7\66\2\2*.\7="+
		"\2\2+,\7\64\2\2,.\5\26\f\2-*\3\2\2\2-+\3\2\2\2.\61\3\2\2\2/\60\7\65\2"+
		"\2\60\62\5\24\13\2\61/\3\2\2\2\61\62\3\2\2\2\62\5\3\2\2\2\63;\7L\2\2\64"+
		"\66\7\17\2\2\65\64\3\2\2\2\65\66\3\2\2\2\66\67\3\2\2\2\67;\7M\2\28;\7"+
		".\2\29;\7/\2\2:\63\3\2\2\2:\65\3\2\2\2:8\3\2\2\2:9\3\2\2\2;\7\3\2\2\2"+
		"<=\7:\2\2=>\7\36\2\2>?\7;\2\2?@\7\27\2\2@C\5\n\6\2AB\7$\2\2BD\5\16\b\2"+
		"CA\3\2\2\2CD\3\2\2\2DE\3\2\2\2EF\7\37\2\2F\t\3\2\2\2GJ\7L\2\2HJ\5\f\7"+
		"\2IG\3\2\2\2IH\3\2\2\2J\13\3\2\2\2KL\t\3\2\2L\r\3\2\2\2MN\7<\2\2NP\7\36"+
		"\2\2OQ\5\20\t\2PO\3\2\2\2PQ\3\2\2\2QR\3\2\2\2RS\7\37\2\2S\17\3\2\2\2T"+
		"Y\5\22\n\2UV\7$\2\2VX\5\22\n\2WU\3\2\2\2X[\3\2\2\2YW\3\2\2\2YZ\3\2\2\2"+
		"Z\21\3\2\2\2[Y\3\2\2\2\\]\7L\2\2]^\7\27\2\2^_\5\6\4\2_\23\3\2\2\2`a\7"+
		"K\2\2a\25\3\2\2\2bc\7K\2\2c\27\3\2\2\2de\7K\2\2e\31\3\2\2\2\f\36$-\61"+
		"\65:CIPY";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}