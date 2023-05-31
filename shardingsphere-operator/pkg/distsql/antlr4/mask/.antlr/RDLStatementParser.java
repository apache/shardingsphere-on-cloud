// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/mask/RDLStatement.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class RDLStatementParser extends Parser {
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
		RULE_createMaskRule = 0, RULE_alterMaskRule = 1, RULE_dropMaskRule = 2, 
		RULE_maskRuleDefinition = 3, RULE_columnDefinition = 4, RULE_columnName = 5, 
		RULE_ifExists = 6, RULE_ifNotExists = 7, RULE_literal = 8, RULE_algorithmDefinition = 9, 
		RULE_algorithmTypeName = 10, RULE_buildInMaskAlgorithmType = 11, RULE_propertiesDefinition = 12, 
		RULE_properties = 13, RULE_property = 14, RULE_ruleName = 15;
	private static String[] makeRuleNames() {
		return new String[] {
			"createMaskRule", "alterMaskRule", "dropMaskRule", "maskRuleDefinition", 
			"columnDefinition", "columnName", "ifExists", "ifNotExists", "literal", 
			"algorithmDefinition", "algorithmTypeName", "buildInMaskAlgorithmType", 
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
	public String getGrammarFileName() { return "RDLStatement.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public RDLStatementParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class CreateMaskRuleContext extends ParserRuleContext {
		public TerminalNode CREATE() { return getToken(RDLStatementParser.CREATE, 0); }
		public TerminalNode MASK() { return getToken(RDLStatementParser.MASK, 0); }
		public TerminalNode RULE() { return getToken(RDLStatementParser.RULE, 0); }
		public List<MaskRuleDefinitionContext> maskRuleDefinition() {
			return getRuleContexts(MaskRuleDefinitionContext.class);
		}
		public MaskRuleDefinitionContext maskRuleDefinition(int i) {
			return getRuleContext(MaskRuleDefinitionContext.class,i);
		}
		public TerminalNode TABLE() { return getToken(RDLStatementParser.TABLE, 0); }
		public IfNotExistsContext ifNotExists() {
			return getRuleContext(IfNotExistsContext.class,0);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public CreateMaskRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_createMaskRule; }
	}

	public final CreateMaskRuleContext createMaskRule() throws RecognitionException {
		CreateMaskRuleContext _localctx = new CreateMaskRuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_createMaskRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			match(CREATE);
			setState(33);
			match(MASK);
			setState(35);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TABLE) {
				{
				setState(34);
				match(TABLE);
				}
			}

			setState(37);
			match(RULE);
			setState(39);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IF) {
				{
				setState(38);
				ifNotExists();
				}
			}

			setState(41);
			maskRuleDefinition();
			setState(46);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(42);
				match(COMMA_);
				setState(43);
				maskRuleDefinition();
				}
				}
				setState(48);
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

	public static class AlterMaskRuleContext extends ParserRuleContext {
		public TerminalNode ALTER() { return getToken(RDLStatementParser.ALTER, 0); }
		public TerminalNode MASK() { return getToken(RDLStatementParser.MASK, 0); }
		public TerminalNode RULE() { return getToken(RDLStatementParser.RULE, 0); }
		public List<MaskRuleDefinitionContext> maskRuleDefinition() {
			return getRuleContexts(MaskRuleDefinitionContext.class);
		}
		public MaskRuleDefinitionContext maskRuleDefinition(int i) {
			return getRuleContext(MaskRuleDefinitionContext.class,i);
		}
		public TerminalNode TABLE() { return getToken(RDLStatementParser.TABLE, 0); }
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public AlterMaskRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_alterMaskRule; }
	}

	public final AlterMaskRuleContext alterMaskRule() throws RecognitionException {
		AlterMaskRuleContext _localctx = new AlterMaskRuleContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_alterMaskRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(49);
			match(ALTER);
			setState(50);
			match(MASK);
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TABLE) {
				{
				setState(51);
				match(TABLE);
				}
			}

			setState(54);
			match(RULE);
			setState(55);
			maskRuleDefinition();
			setState(60);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(56);
				match(COMMA_);
				setState(57);
				maskRuleDefinition();
				}
				}
				setState(62);
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

	public static class DropMaskRuleContext extends ParserRuleContext {
		public TerminalNode DROP() { return getToken(RDLStatementParser.DROP, 0); }
		public TerminalNode MASK() { return getToken(RDLStatementParser.MASK, 0); }
		public TerminalNode RULE() { return getToken(RDLStatementParser.RULE, 0); }
		public List<RuleNameContext> ruleName() {
			return getRuleContexts(RuleNameContext.class);
		}
		public RuleNameContext ruleName(int i) {
			return getRuleContext(RuleNameContext.class,i);
		}
		public TerminalNode TABLE() { return getToken(RDLStatementParser.TABLE, 0); }
		public IfExistsContext ifExists() {
			return getRuleContext(IfExistsContext.class,0);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public DropMaskRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dropMaskRule; }
	}

	public final DropMaskRuleContext dropMaskRule() throws RecognitionException {
		DropMaskRuleContext _localctx = new DropMaskRuleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_dropMaskRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(63);
			match(DROP);
			setState(64);
			match(MASK);
			setState(66);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TABLE) {
				{
				setState(65);
				match(TABLE);
				}
			}

			setState(68);
			match(RULE);
			setState(70);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IF) {
				{
				setState(69);
				ifExists();
				}
			}

			setState(72);
			ruleName();
			setState(77);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(73);
				match(COMMA_);
				setState(74);
				ruleName();
				}
				}
				setState(79);
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

	public static class MaskRuleDefinitionContext extends ParserRuleContext {
		public RuleNameContext ruleName() {
			return getRuleContext(RuleNameContext.class,0);
		}
		public List<TerminalNode> LP_() { return getTokens(RDLStatementParser.LP_); }
		public TerminalNode LP_(int i) {
			return getToken(RDLStatementParser.LP_, i);
		}
		public TerminalNode COLUMNS() { return getToken(RDLStatementParser.COLUMNS, 0); }
		public List<ColumnDefinitionContext> columnDefinition() {
			return getRuleContexts(ColumnDefinitionContext.class);
		}
		public ColumnDefinitionContext columnDefinition(int i) {
			return getRuleContext(ColumnDefinitionContext.class,i);
		}
		public List<TerminalNode> RP_() { return getTokens(RDLStatementParser.RP_); }
		public TerminalNode RP_(int i) {
			return getToken(RDLStatementParser.RP_, i);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public MaskRuleDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_maskRuleDefinition; }
	}

	public final MaskRuleDefinitionContext maskRuleDefinition() throws RecognitionException {
		MaskRuleDefinitionContext _localctx = new MaskRuleDefinitionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_maskRuleDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			ruleName();
			setState(81);
			match(LP_);
			setState(82);
			match(COLUMNS);
			setState(83);
			match(LP_);
			setState(84);
			columnDefinition();
			setState(89);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(85);
				match(COMMA_);
				setState(86);
				columnDefinition();
				}
				}
				setState(91);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(92);
			match(RP_);
			setState(93);
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

	public static class ColumnDefinitionContext extends ParserRuleContext {
		public TerminalNode LP_() { return getToken(RDLStatementParser.LP_, 0); }
		public TerminalNode NAME() { return getToken(RDLStatementParser.NAME, 0); }
		public TerminalNode EQ_() { return getToken(RDLStatementParser.EQ_, 0); }
		public ColumnNameContext columnName() {
			return getRuleContext(ColumnNameContext.class,0);
		}
		public TerminalNode COMMA_() { return getToken(RDLStatementParser.COMMA_, 0); }
		public AlgorithmDefinitionContext algorithmDefinition() {
			return getRuleContext(AlgorithmDefinitionContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RDLStatementParser.RP_, 0); }
		public ColumnDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnDefinition; }
	}

	public final ColumnDefinitionContext columnDefinition() throws RecognitionException {
		ColumnDefinitionContext _localctx = new ColumnDefinitionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_columnDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			match(LP_);
			setState(96);
			match(NAME);
			setState(97);
			match(EQ_);
			setState(98);
			columnName();
			setState(99);
			match(COMMA_);
			setState(100);
			algorithmDefinition();
			setState(101);
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

	public static class ColumnNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public ColumnNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnName; }
	}

	public final ColumnNameContext columnName() throws RecognitionException {
		ColumnNameContext _localctx = new ColumnNameContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_columnName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
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

	public static class IfExistsContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(RDLStatementParser.IF, 0); }
		public TerminalNode EXISTS() { return getToken(RDLStatementParser.EXISTS, 0); }
		public IfExistsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifExists; }
	}

	public final IfExistsContext ifExists() throws RecognitionException {
		IfExistsContext _localctx = new IfExistsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ifExists);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(IF);
			setState(106);
			match(EXISTS);
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

	public static class IfNotExistsContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(RDLStatementParser.IF, 0); }
		public TerminalNode NOT() { return getToken(RDLStatementParser.NOT, 0); }
		public TerminalNode EXISTS() { return getToken(RDLStatementParser.EXISTS, 0); }
		public IfNotExistsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifNotExists; }
	}

	public final IfNotExistsContext ifNotExists() throws RecognitionException {
		IfNotExistsContext _localctx = new IfNotExistsContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ifNotExists);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(IF);
			setState(109);
			match(NOT);
			setState(110);
			match(EXISTS);
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
		public TerminalNode STRING_() { return getToken(RDLStatementParser.STRING_, 0); }
		public TerminalNode INT_() { return getToken(RDLStatementParser.INT_, 0); }
		public TerminalNode MINUS_() { return getToken(RDLStatementParser.MINUS_, 0); }
		public TerminalNode TRUE() { return getToken(RDLStatementParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(RDLStatementParser.FALSE, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_literal);
		int _la;
		try {
			setState(119);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(112);
				match(STRING_);
				}
				break;
			case MINUS_:
			case INT_:
				enterOuterAlt(_localctx, 2);
				{
				setState(114);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS_) {
					{
					setState(113);
					match(MINUS_);
					}
				}

				setState(116);
				match(INT_);
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(117);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(118);
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
		public TerminalNode TYPE() { return getToken(RDLStatementParser.TYPE, 0); }
		public TerminalNode LP_() { return getToken(RDLStatementParser.LP_, 0); }
		public TerminalNode NAME() { return getToken(RDLStatementParser.NAME, 0); }
		public TerminalNode EQ_() { return getToken(RDLStatementParser.EQ_, 0); }
		public AlgorithmTypeNameContext algorithmTypeName() {
			return getRuleContext(AlgorithmTypeNameContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RDLStatementParser.RP_, 0); }
		public TerminalNode COMMA_() { return getToken(RDLStatementParser.COMMA_, 0); }
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
		enterRule(_localctx, 18, RULE_algorithmDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			match(TYPE);
			setState(122);
			match(LP_);
			setState(123);
			match(NAME);
			setState(124);
			match(EQ_);
			setState(125);
			algorithmTypeName();
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA_) {
				{
				setState(126);
				match(COMMA_);
				setState(127);
				propertiesDefinition();
				}
			}

			setState(130);
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
		public TerminalNode STRING_() { return getToken(RDLStatementParser.STRING_, 0); }
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
		enterRule(_localctx, 20, RULE_algorithmTypeName);
		try {
			setState(134);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
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
				setState(133);
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
		public TerminalNode MD5() { return getToken(RDLStatementParser.MD5, 0); }
		public TerminalNode KEEP_FIRST_N_LAST_M() { return getToken(RDLStatementParser.KEEP_FIRST_N_LAST_M, 0); }
		public TerminalNode KEEP_FROM_X_TO_Y() { return getToken(RDLStatementParser.KEEP_FROM_X_TO_Y, 0); }
		public TerminalNode MASK_FIRST_N_LAST_M() { return getToken(RDLStatementParser.MASK_FIRST_N_LAST_M, 0); }
		public TerminalNode MASK_FROM_X_TO_Y() { return getToken(RDLStatementParser.MASK_FROM_X_TO_Y, 0); }
		public TerminalNode MASK_BEFORE_SPECIAL_CHARS() { return getToken(RDLStatementParser.MASK_BEFORE_SPECIAL_CHARS, 0); }
		public TerminalNode MASK_AFTER_SPECIAL_CHARS() { return getToken(RDLStatementParser.MASK_AFTER_SPECIAL_CHARS, 0); }
		public TerminalNode PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE() { return getToken(RDLStatementParser.PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE() { return getToken(RDLStatementParser.MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode LANDLINE_NUMBER_RANDOM_REPLACE() { return getToken(RDLStatementParser.LANDLINE_NUMBER_RANDOM_REPLACE, 0); }
		public TerminalNode TELEPHONE_RANDOM_REPLACE() { return getToken(RDLStatementParser.TELEPHONE_RANDOM_REPLACE, 0); }
		public TerminalNode UNIFIED_CREDIT_CODE_RANDOM_REPLACE() { return getToken(RDLStatementParser.UNIFIED_CREDIT_CODE_RANDOM_REPLACE, 0); }
		public TerminalNode GENERIC_TABLE_RANDOM_REPLACE() { return getToken(RDLStatementParser.GENERIC_TABLE_RANDOM_REPLACE, 0); }
		public BuildInMaskAlgorithmTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_buildInMaskAlgorithmType; }
	}

	public final BuildInMaskAlgorithmTypeContext buildInMaskAlgorithmType() throws RecognitionException {
		BuildInMaskAlgorithmTypeContext _localctx = new BuildInMaskAlgorithmTypeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_buildInMaskAlgorithmType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
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
		public TerminalNode PROPERTIES() { return getToken(RDLStatementParser.PROPERTIES, 0); }
		public TerminalNode LP_() { return getToken(RDLStatementParser.LP_, 0); }
		public TerminalNode RP_() { return getToken(RDLStatementParser.RP_, 0); }
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
		enterRule(_localctx, 24, RULE_propertiesDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			match(PROPERTIES);
			setState(139);
			match(LP_);
			setState(141);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING_) {
				{
				setState(140);
				properties();
				}
			}

			setState(143);
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
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public PropertiesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_properties; }
	}

	public final PropertiesContext properties() throws RecognitionException {
		PropertiesContext _localctx = new PropertiesContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_properties);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(145);
			property();
			setState(150);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(146);
				match(COMMA_);
				setState(147);
				property();
				}
				}
				setState(152);
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
		public TerminalNode EQ_() { return getToken(RDLStatementParser.EQ_, 0); }
		public TerminalNode STRING_() { return getToken(RDLStatementParser.STRING_, 0); }
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
		enterRule(_localctx, 28, RULE_property);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			((PropertyContext)_localctx).key = match(STRING_);
			setState(154);
			match(EQ_);
			setState(155);
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
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public RuleNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleName; }
	}

	public final RuleNameContext ruleName() throws RecognitionException {
		RuleNameContext _localctx = new RuleNameContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_ruleName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(157);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3W\u00a2\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\3\2\3\2\3"+
		"\2\5\2&\n\2\3\2\3\2\5\2*\n\2\3\2\3\2\3\2\7\2/\n\2\f\2\16\2\62\13\2\3\3"+
		"\3\3\3\3\5\3\67\n\3\3\3\3\3\3\3\3\3\7\3=\n\3\f\3\16\3@\13\3\3\4\3\4\3"+
		"\4\5\4E\n\4\3\4\3\4\5\4I\n\4\3\4\3\4\3\4\7\4N\n\4\f\4\16\4Q\13\4\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\7\5Z\n\5\f\5\16\5]\13\5\3\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\5\nu"+
		"\n\n\3\n\3\n\3\n\5\nz\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u0083"+
		"\n\13\3\13\3\13\3\f\3\f\5\f\u0089\n\f\3\r\3\r\3\16\3\16\3\16\5\16\u0090"+
		"\n\16\3\16\3\16\3\17\3\17\3\17\7\17\u0097\n\17\f\17\16\17\u009a\13\17"+
		"\3\20\3\20\3\20\3\20\3\21\3\21\3\21\2\2\22\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \2\3\3\2BN\2\u00a2\2\"\3\2\2\2\4\63\3\2\2\2\6A\3\2\2\2\b"+
		"R\3\2\2\2\na\3\2\2\2\fi\3\2\2\2\16k\3\2\2\2\20n\3\2\2\2\22y\3\2\2\2\24"+
		"{\3\2\2\2\26\u0088\3\2\2\2\30\u008a\3\2\2\2\32\u008c\3\2\2\2\34\u0093"+
		"\3\2\2\2\36\u009b\3\2\2\2 \u009f\3\2\2\2\"#\7\60\2\2#%\7\66\2\2$&\7<\2"+
		"\2%$\3\2\2\2%&\3\2\2\2&\'\3\2\2\2\')\7\64\2\2(*\5\20\t\2)(\3\2\2\2)*\3"+
		"\2\2\2*+\3\2\2\2+\60\5\b\5\2,-\7$\2\2-/\5\b\5\2.,\3\2\2\2/\62\3\2\2\2"+
		"\60.\3\2\2\2\60\61\3\2\2\2\61\3\3\2\2\2\62\60\3\2\2\2\63\64\7\61\2\2\64"+
		"\66\7\66\2\2\65\67\7<\2\2\66\65\3\2\2\2\66\67\3\2\2\2\678\3\2\2\289\7"+
		"\64\2\29>\5\b\5\2:;\7$\2\2;=\5\b\5\2<:\3\2\2\2=@\3\2\2\2><\3\2\2\2>?\3"+
		"\2\2\2?\5\3\2\2\2@>\3\2\2\2AB\7\62\2\2BD\7\66\2\2CE\7<\2\2DC\3\2\2\2D"+
		"E\3\2\2\2EF\3\2\2\2FH\7\64\2\2GI\5\16\b\2HG\3\2\2\2HI\3\2\2\2IJ\3\2\2"+
		"\2JO\5 \21\2KL\7$\2\2LN\5 \21\2MK\3\2\2\2NQ\3\2\2\2OM\3\2\2\2OP\3\2\2"+
		"\2P\7\3\2\2\2QO\3\2\2\2RS\5 \21\2ST\7\36\2\2TU\7=\2\2UV\7\36\2\2V[\5\n"+
		"\6\2WX\7$\2\2XZ\5\n\6\2YW\3\2\2\2Z]\3\2\2\2[Y\3\2\2\2[\\\3\2\2\2\\^\3"+
		"\2\2\2][\3\2\2\2^_\7\37\2\2_`\7\37\2\2`\t\3\2\2\2ab\7\36\2\2bc\78\2\2"+
		"cd\7\27\2\2de\5\f\7\2ef\7$\2\2fg\5\24\13\2gh\7\37\2\2h\13\3\2\2\2ij\7"+
		"Q\2\2j\r\3\2\2\2kl\7>\2\2lm\7?\2\2m\17\3\2\2\2no\7>\2\2op\7A\2\2pq\7?"+
		"\2\2q\21\3\2\2\2rz\7R\2\2su\7\17\2\2ts\3\2\2\2tu\3\2\2\2uv\3\2\2\2vz\7"+
		"S\2\2wz\7.\2\2xz\7/\2\2yr\3\2\2\2yt\3\2\2\2yw\3\2\2\2yx\3\2\2\2z\23\3"+
		"\2\2\2{|\7\67\2\2|}\7\36\2\2}~\78\2\2~\177\7\27\2\2\177\u0082\5\26\f\2"+
		"\u0080\u0081\7$\2\2\u0081\u0083\5\32\16\2\u0082\u0080\3\2\2\2\u0082\u0083"+
		"\3\2\2\2\u0083\u0084\3\2\2\2\u0084\u0085\7\37\2\2\u0085\25\3\2\2\2\u0086"+
		"\u0089\7R\2\2\u0087\u0089\5\30\r\2\u0088\u0086\3\2\2\2\u0088\u0087\3\2"+
		"\2\2\u0089\27\3\2\2\2\u008a\u008b\t\2\2\2\u008b\31\3\2\2\2\u008c\u008d"+
		"\79\2\2\u008d\u008f\7\36\2\2\u008e\u0090\5\34\17\2\u008f\u008e\3\2\2\2"+
		"\u008f\u0090\3\2\2\2\u0090\u0091\3\2\2\2\u0091\u0092\7\37\2\2\u0092\33"+
		"\3\2\2\2\u0093\u0098\5\36\20\2\u0094\u0095\7$\2\2\u0095\u0097\5\36\20"+
		"\2\u0096\u0094\3\2\2\2\u0097\u009a\3\2\2\2\u0098\u0096\3\2\2\2\u0098\u0099"+
		"\3\2\2\2\u0099\35\3\2\2\2\u009a\u0098\3\2\2\2\u009b\u009c\7R\2\2\u009c"+
		"\u009d\7\27\2\2\u009d\u009e\5\22\n\2\u009e\37\3\2\2\2\u009f\u00a0\7Q\2"+
		"\2\u00a0!\3\2\2\2\21%)\60\66>DHO[ty\u0082\u0088\u008f\u0098";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}