// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/encrypt/RDLStatement.g4 by ANTLR 4.9.2
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
		RULE_createEncryptRule = 0, RULE_alterEncryptRule = 1, RULE_dropEncryptRule = 2, 
		RULE_encryptRuleDefinition = 3, RULE_resourceDefinition = 4, RULE_resourceName = 5, 
		RULE_encryptColumnDefinition = 6, RULE_columnDefinition = 7, RULE_columnName = 8, 
		RULE_dataType = 9, RULE_plainColumnDefinition = 10, RULE_plainColumnName = 11, 
		RULE_cipherColumnDefinition = 12, RULE_cipherColumnName = 13, RULE_assistedQueryColumnDefinition = 14, 
		RULE_assistedQueryColumnName = 15, RULE_likeQueryColumnDefinition = 16, 
		RULE_likeQueryColumnName = 17, RULE_encryptAlgorithm = 18, RULE_assistedQueryAlgorithm = 19, 
		RULE_likeQueryAlgorithm = 20, RULE_queryWithCipherColumn = 21, RULE_ifExists = 22, 
		RULE_ifNotExists = 23, RULE_literal = 24, RULE_algorithmDefinition = 25, 
		RULE_algorithmTypeName = 26, RULE_buildinAlgorithmTypeName = 27, RULE_propertiesDefinition = 28, 
		RULE_properties = 29, RULE_property = 30, RULE_tableName = 31;
	private static String[] makeRuleNames() {
		return new String[] {
			"createEncryptRule", "alterEncryptRule", "dropEncryptRule", "encryptRuleDefinition", 
			"resourceDefinition", "resourceName", "encryptColumnDefinition", "columnDefinition", 
			"columnName", "dataType", "plainColumnDefinition", "plainColumnName", 
			"cipherColumnDefinition", "cipherColumnName", "assistedQueryColumnDefinition", 
			"assistedQueryColumnName", "likeQueryColumnDefinition", "likeQueryColumnName", 
			"encryptAlgorithm", "assistedQueryAlgorithm", "likeQueryAlgorithm", "queryWithCipherColumn", 
			"ifExists", "ifNotExists", "literal", "algorithmDefinition", "algorithmTypeName", 
			"buildinAlgorithmTypeName", "propertiesDefinition", "properties", "property", 
			"tableName"
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

	public static class CreateEncryptRuleContext extends ParserRuleContext {
		public TerminalNode CREATE() { return getToken(RDLStatementParser.CREATE, 0); }
		public TerminalNode ENCRYPT() { return getToken(RDLStatementParser.ENCRYPT, 0); }
		public TerminalNode RULE() { return getToken(RDLStatementParser.RULE, 0); }
		public List<EncryptRuleDefinitionContext> encryptRuleDefinition() {
			return getRuleContexts(EncryptRuleDefinitionContext.class);
		}
		public EncryptRuleDefinitionContext encryptRuleDefinition(int i) {
			return getRuleContext(EncryptRuleDefinitionContext.class,i);
		}
		public IfNotExistsContext ifNotExists() {
			return getRuleContext(IfNotExistsContext.class,0);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public CreateEncryptRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_createEncryptRule; }
	}

	public final CreateEncryptRuleContext createEncryptRule() throws RecognitionException {
		CreateEncryptRuleContext _localctx = new CreateEncryptRuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_createEncryptRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			match(CREATE);
			setState(65);
			match(ENCRYPT);
			setState(66);
			match(RULE);
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IF) {
				{
				setState(67);
				ifNotExists();
				}
			}

			setState(70);
			encryptRuleDefinition();
			setState(75);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(71);
				match(COMMA_);
				setState(72);
				encryptRuleDefinition();
				}
				}
				setState(77);
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

	public static class AlterEncryptRuleContext extends ParserRuleContext {
		public TerminalNode ALTER() { return getToken(RDLStatementParser.ALTER, 0); }
		public TerminalNode ENCRYPT() { return getToken(RDLStatementParser.ENCRYPT, 0); }
		public TerminalNode RULE() { return getToken(RDLStatementParser.RULE, 0); }
		public List<EncryptRuleDefinitionContext> encryptRuleDefinition() {
			return getRuleContexts(EncryptRuleDefinitionContext.class);
		}
		public EncryptRuleDefinitionContext encryptRuleDefinition(int i) {
			return getRuleContext(EncryptRuleDefinitionContext.class,i);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public AlterEncryptRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_alterEncryptRule; }
	}

	public final AlterEncryptRuleContext alterEncryptRule() throws RecognitionException {
		AlterEncryptRuleContext _localctx = new AlterEncryptRuleContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_alterEncryptRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			match(ALTER);
			setState(79);
			match(ENCRYPT);
			setState(80);
			match(RULE);
			setState(81);
			encryptRuleDefinition();
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(82);
				match(COMMA_);
				setState(83);
				encryptRuleDefinition();
				}
				}
				setState(88);
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

	public static class DropEncryptRuleContext extends ParserRuleContext {
		public TerminalNode DROP() { return getToken(RDLStatementParser.DROP, 0); }
		public TerminalNode ENCRYPT() { return getToken(RDLStatementParser.ENCRYPT, 0); }
		public TerminalNode RULE() { return getToken(RDLStatementParser.RULE, 0); }
		public List<TableNameContext> tableName() {
			return getRuleContexts(TableNameContext.class);
		}
		public TableNameContext tableName(int i) {
			return getRuleContext(TableNameContext.class,i);
		}
		public IfExistsContext ifExists() {
			return getRuleContext(IfExistsContext.class,0);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public DropEncryptRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dropEncryptRule; }
	}

	public final DropEncryptRuleContext dropEncryptRule() throws RecognitionException {
		DropEncryptRuleContext _localctx = new DropEncryptRuleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_dropEncryptRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(DROP);
			setState(90);
			match(ENCRYPT);
			setState(91);
			match(RULE);
			setState(93);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IF) {
				{
				setState(92);
				ifExists();
				}
			}

			setState(95);
			tableName();
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(96);
				match(COMMA_);
				setState(97);
				tableName();
				}
				}
				setState(102);
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

	public static class EncryptRuleDefinitionContext extends ParserRuleContext {
		public TableNameContext tableName() {
			return getRuleContext(TableNameContext.class,0);
		}
		public List<TerminalNode> LP_() { return getTokens(RDLStatementParser.LP_); }
		public TerminalNode LP_(int i) {
			return getToken(RDLStatementParser.LP_, i);
		}
		public TerminalNode COLUMNS() { return getToken(RDLStatementParser.COLUMNS, 0); }
		public List<EncryptColumnDefinitionContext> encryptColumnDefinition() {
			return getRuleContexts(EncryptColumnDefinitionContext.class);
		}
		public EncryptColumnDefinitionContext encryptColumnDefinition(int i) {
			return getRuleContext(EncryptColumnDefinitionContext.class,i);
		}
		public List<TerminalNode> RP_() { return getTokens(RDLStatementParser.RP_); }
		public TerminalNode RP_(int i) {
			return getToken(RDLStatementParser.RP_, i);
		}
		public ResourceDefinitionContext resourceDefinition() {
			return getRuleContext(ResourceDefinitionContext.class,0);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public TerminalNode QUERY_WITH_CIPHER_COLUMN() { return getToken(RDLStatementParser.QUERY_WITH_CIPHER_COLUMN, 0); }
		public TerminalNode EQ_() { return getToken(RDLStatementParser.EQ_, 0); }
		public QueryWithCipherColumnContext queryWithCipherColumn() {
			return getRuleContext(QueryWithCipherColumnContext.class,0);
		}
		public EncryptRuleDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_encryptRuleDefinition; }
	}

	public final EncryptRuleDefinitionContext encryptRuleDefinition() throws RecognitionException {
		EncryptRuleDefinitionContext _localctx = new EncryptRuleDefinitionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_encryptRuleDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			tableName();
			setState(104);
			match(LP_);
			setState(108);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RESOURCE) {
				{
				setState(105);
				resourceDefinition();
				setState(106);
				match(COMMA_);
				}
			}

			setState(110);
			match(COLUMNS);
			setState(111);
			match(LP_);
			setState(112);
			encryptColumnDefinition();
			setState(117);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(113);
				match(COMMA_);
				setState(114);
				encryptColumnDefinition();
				}
				}
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(120);
			match(RP_);
			setState(125);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA_) {
				{
				setState(121);
				match(COMMA_);
				setState(122);
				match(QUERY_WITH_CIPHER_COLUMN);
				setState(123);
				match(EQ_);
				setState(124);
				queryWithCipherColumn();
				}
			}

			setState(127);
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

	public static class ResourceDefinitionContext extends ParserRuleContext {
		public TerminalNode RESOURCE() { return getToken(RDLStatementParser.RESOURCE, 0); }
		public TerminalNode EQ_() { return getToken(RDLStatementParser.EQ_, 0); }
		public ResourceNameContext resourceName() {
			return getRuleContext(ResourceNameContext.class,0);
		}
		public ResourceDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_resourceDefinition; }
	}

	public final ResourceDefinitionContext resourceDefinition() throws RecognitionException {
		ResourceDefinitionContext _localctx = new ResourceDefinitionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_resourceDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(RESOURCE);
			setState(130);
			match(EQ_);
			setState(131);
			resourceName();
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

	public static class ResourceNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public ResourceNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_resourceName; }
	}

	public final ResourceNameContext resourceName() throws RecognitionException {
		ResourceNameContext _localctx = new ResourceNameContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_resourceName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
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

	public static class EncryptColumnDefinitionContext extends ParserRuleContext {
		public TerminalNode LP_() { return getToken(RDLStatementParser.LP_, 0); }
		public ColumnDefinitionContext columnDefinition() {
			return getRuleContext(ColumnDefinitionContext.class,0);
		}
		public List<TerminalNode> COMMA_() { return getTokens(RDLStatementParser.COMMA_); }
		public TerminalNode COMMA_(int i) {
			return getToken(RDLStatementParser.COMMA_, i);
		}
		public CipherColumnDefinitionContext cipherColumnDefinition() {
			return getRuleContext(CipherColumnDefinitionContext.class,0);
		}
		public EncryptAlgorithmContext encryptAlgorithm() {
			return getRuleContext(EncryptAlgorithmContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RDLStatementParser.RP_, 0); }
		public PlainColumnDefinitionContext plainColumnDefinition() {
			return getRuleContext(PlainColumnDefinitionContext.class,0);
		}
		public AssistedQueryColumnDefinitionContext assistedQueryColumnDefinition() {
			return getRuleContext(AssistedQueryColumnDefinitionContext.class,0);
		}
		public LikeQueryColumnDefinitionContext likeQueryColumnDefinition() {
			return getRuleContext(LikeQueryColumnDefinitionContext.class,0);
		}
		public AssistedQueryAlgorithmContext assistedQueryAlgorithm() {
			return getRuleContext(AssistedQueryAlgorithmContext.class,0);
		}
		public LikeQueryAlgorithmContext likeQueryAlgorithm() {
			return getRuleContext(LikeQueryAlgorithmContext.class,0);
		}
		public TerminalNode QUERY_WITH_CIPHER_COLUMN() { return getToken(RDLStatementParser.QUERY_WITH_CIPHER_COLUMN, 0); }
		public TerminalNode EQ_() { return getToken(RDLStatementParser.EQ_, 0); }
		public QueryWithCipherColumnContext queryWithCipherColumn() {
			return getRuleContext(QueryWithCipherColumnContext.class,0);
		}
		public EncryptColumnDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_encryptColumnDefinition; }
	}

	public final EncryptColumnDefinitionContext encryptColumnDefinition() throws RecognitionException {
		EncryptColumnDefinitionContext _localctx = new EncryptColumnDefinitionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_encryptColumnDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(LP_);
			setState(136);
			columnDefinition();
			setState(139);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(137);
				match(COMMA_);
				setState(138);
				plainColumnDefinition();
				}
				break;
			}
			setState(141);
			match(COMMA_);
			setState(142);
			cipherColumnDefinition();
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(143);
				match(COMMA_);
				setState(144);
				assistedQueryColumnDefinition();
				}
				break;
			}
			setState(149);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(147);
				match(COMMA_);
				setState(148);
				likeQueryColumnDefinition();
				}
				break;
			}
			setState(151);
			match(COMMA_);
			setState(152);
			encryptAlgorithm();
			setState(155);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(153);
				match(COMMA_);
				setState(154);
				assistedQueryAlgorithm();
				}
				break;
			}
			setState(159);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(157);
				match(COMMA_);
				setState(158);
				likeQueryAlgorithm();
				}
				break;
			}
			setState(165);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA_) {
				{
				setState(161);
				match(COMMA_);
				setState(162);
				match(QUERY_WITH_CIPHER_COLUMN);
				setState(163);
				match(EQ_);
				setState(164);
				queryWithCipherColumn();
				}
			}

			setState(167);
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
		public TerminalNode NAME() { return getToken(RDLStatementParser.NAME, 0); }
		public List<TerminalNode> EQ_() { return getTokens(RDLStatementParser.EQ_); }
		public TerminalNode EQ_(int i) {
			return getToken(RDLStatementParser.EQ_, i);
		}
		public ColumnNameContext columnName() {
			return getRuleContext(ColumnNameContext.class,0);
		}
		public TerminalNode COMMA_() { return getToken(RDLStatementParser.COMMA_, 0); }
		public TerminalNode DATA_TYPE() { return getToken(RDLStatementParser.DATA_TYPE, 0); }
		public DataTypeContext dataType() {
			return getRuleContext(DataTypeContext.class,0);
		}
		public ColumnDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnDefinition; }
	}

	public final ColumnDefinitionContext columnDefinition() throws RecognitionException {
		ColumnDefinitionContext _localctx = new ColumnDefinitionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_columnDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(NAME);
			setState(170);
			match(EQ_);
			setState(171);
			columnName();
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(172);
				match(COMMA_);
				setState(173);
				match(DATA_TYPE);
				setState(174);
				match(EQ_);
				setState(175);
				dataType();
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

	public static class ColumnNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public ColumnNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnName; }
	}

	public final ColumnNameContext columnName() throws RecognitionException {
		ColumnNameContext _localctx = new ColumnNameContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_columnName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
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

	public static class DataTypeContext extends ParserRuleContext {
		public TerminalNode STRING_() { return getToken(RDLStatementParser.STRING_, 0); }
		public DataTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dataType; }
	}

	public final DataTypeContext dataType() throws RecognitionException {
		DataTypeContext _localctx = new DataTypeContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_dataType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			match(STRING_);
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

	public static class PlainColumnDefinitionContext extends ParserRuleContext {
		public TerminalNode PLAIN() { return getToken(RDLStatementParser.PLAIN, 0); }
		public List<TerminalNode> EQ_() { return getTokens(RDLStatementParser.EQ_); }
		public TerminalNode EQ_(int i) {
			return getToken(RDLStatementParser.EQ_, i);
		}
		public PlainColumnNameContext plainColumnName() {
			return getRuleContext(PlainColumnNameContext.class,0);
		}
		public TerminalNode COMMA_() { return getToken(RDLStatementParser.COMMA_, 0); }
		public TerminalNode PLAIN_DATA_TYPE() { return getToken(RDLStatementParser.PLAIN_DATA_TYPE, 0); }
		public DataTypeContext dataType() {
			return getRuleContext(DataTypeContext.class,0);
		}
		public PlainColumnDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_plainColumnDefinition; }
	}

	public final PlainColumnDefinitionContext plainColumnDefinition() throws RecognitionException {
		PlainColumnDefinitionContext _localctx = new PlainColumnDefinitionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_plainColumnDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			match(PLAIN);
			setState(183);
			match(EQ_);
			setState(184);
			plainColumnName();
			setState(189);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				setState(185);
				match(COMMA_);
				setState(186);
				match(PLAIN_DATA_TYPE);
				setState(187);
				match(EQ_);
				setState(188);
				dataType();
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

	public static class PlainColumnNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public PlainColumnNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_plainColumnName; }
	}

	public final PlainColumnNameContext plainColumnName() throws RecognitionException {
		PlainColumnNameContext _localctx = new PlainColumnNameContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_plainColumnName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(191);
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

	public static class CipherColumnDefinitionContext extends ParserRuleContext {
		public TerminalNode CIPHER() { return getToken(RDLStatementParser.CIPHER, 0); }
		public List<TerminalNode> EQ_() { return getTokens(RDLStatementParser.EQ_); }
		public TerminalNode EQ_(int i) {
			return getToken(RDLStatementParser.EQ_, i);
		}
		public CipherColumnNameContext cipherColumnName() {
			return getRuleContext(CipherColumnNameContext.class,0);
		}
		public TerminalNode COMMA_() { return getToken(RDLStatementParser.COMMA_, 0); }
		public TerminalNode CIPHER_DATA_TYPE() { return getToken(RDLStatementParser.CIPHER_DATA_TYPE, 0); }
		public DataTypeContext dataType() {
			return getRuleContext(DataTypeContext.class,0);
		}
		public CipherColumnDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cipherColumnDefinition; }
	}

	public final CipherColumnDefinitionContext cipherColumnDefinition() throws RecognitionException {
		CipherColumnDefinitionContext _localctx = new CipherColumnDefinitionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_cipherColumnDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			match(CIPHER);
			setState(194);
			match(EQ_);
			setState(195);
			cipherColumnName();
			setState(200);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				{
				setState(196);
				match(COMMA_);
				setState(197);
				match(CIPHER_DATA_TYPE);
				setState(198);
				match(EQ_);
				setState(199);
				dataType();
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

	public static class CipherColumnNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public CipherColumnNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cipherColumnName; }
	}

	public final CipherColumnNameContext cipherColumnName() throws RecognitionException {
		CipherColumnNameContext _localctx = new CipherColumnNameContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_cipherColumnName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
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

	public static class AssistedQueryColumnDefinitionContext extends ParserRuleContext {
		public TerminalNode ASSISTED_QUERY_COLUMN() { return getToken(RDLStatementParser.ASSISTED_QUERY_COLUMN, 0); }
		public List<TerminalNode> EQ_() { return getTokens(RDLStatementParser.EQ_); }
		public TerminalNode EQ_(int i) {
			return getToken(RDLStatementParser.EQ_, i);
		}
		public AssistedQueryColumnNameContext assistedQueryColumnName() {
			return getRuleContext(AssistedQueryColumnNameContext.class,0);
		}
		public TerminalNode COMMA_() { return getToken(RDLStatementParser.COMMA_, 0); }
		public TerminalNode ASSISTED_QUERY_DATA_TYPE() { return getToken(RDLStatementParser.ASSISTED_QUERY_DATA_TYPE, 0); }
		public DataTypeContext dataType() {
			return getRuleContext(DataTypeContext.class,0);
		}
		public AssistedQueryColumnDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assistedQueryColumnDefinition; }
	}

	public final AssistedQueryColumnDefinitionContext assistedQueryColumnDefinition() throws RecognitionException {
		AssistedQueryColumnDefinitionContext _localctx = new AssistedQueryColumnDefinitionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_assistedQueryColumnDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			match(ASSISTED_QUERY_COLUMN);
			setState(205);
			match(EQ_);
			setState(206);
			assistedQueryColumnName();
			setState(211);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(207);
				match(COMMA_);
				setState(208);
				match(ASSISTED_QUERY_DATA_TYPE);
				setState(209);
				match(EQ_);
				setState(210);
				dataType();
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

	public static class AssistedQueryColumnNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public AssistedQueryColumnNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assistedQueryColumnName; }
	}

	public final AssistedQueryColumnNameContext assistedQueryColumnName() throws RecognitionException {
		AssistedQueryColumnNameContext _localctx = new AssistedQueryColumnNameContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_assistedQueryColumnName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
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

	public static class LikeQueryColumnDefinitionContext extends ParserRuleContext {
		public TerminalNode LIKE_QUERY_COLUMN() { return getToken(RDLStatementParser.LIKE_QUERY_COLUMN, 0); }
		public List<TerminalNode> EQ_() { return getTokens(RDLStatementParser.EQ_); }
		public TerminalNode EQ_(int i) {
			return getToken(RDLStatementParser.EQ_, i);
		}
		public LikeQueryColumnNameContext likeQueryColumnName() {
			return getRuleContext(LikeQueryColumnNameContext.class,0);
		}
		public TerminalNode COMMA_() { return getToken(RDLStatementParser.COMMA_, 0); }
		public TerminalNode LIKE_QUERY_DATA_TYPE() { return getToken(RDLStatementParser.LIKE_QUERY_DATA_TYPE, 0); }
		public DataTypeContext dataType() {
			return getRuleContext(DataTypeContext.class,0);
		}
		public LikeQueryColumnDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_likeQueryColumnDefinition; }
	}

	public final LikeQueryColumnDefinitionContext likeQueryColumnDefinition() throws RecognitionException {
		LikeQueryColumnDefinitionContext _localctx = new LikeQueryColumnDefinitionContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_likeQueryColumnDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			match(LIKE_QUERY_COLUMN);
			setState(216);
			match(EQ_);
			setState(217);
			likeQueryColumnName();
			setState(222);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				{
				setState(218);
				match(COMMA_);
				setState(219);
				match(LIKE_QUERY_DATA_TYPE);
				setState(220);
				match(EQ_);
				setState(221);
				dataType();
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

	public static class LikeQueryColumnNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public LikeQueryColumnNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_likeQueryColumnName; }
	}

	public final LikeQueryColumnNameContext likeQueryColumnName() throws RecognitionException {
		LikeQueryColumnNameContext _localctx = new LikeQueryColumnNameContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_likeQueryColumnName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
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

	public static class EncryptAlgorithmContext extends ParserRuleContext {
		public TerminalNode ENCRYPT_ALGORITHM() { return getToken(RDLStatementParser.ENCRYPT_ALGORITHM, 0); }
		public TerminalNode LP_() { return getToken(RDLStatementParser.LP_, 0); }
		public AlgorithmDefinitionContext algorithmDefinition() {
			return getRuleContext(AlgorithmDefinitionContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RDLStatementParser.RP_, 0); }
		public EncryptAlgorithmContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_encryptAlgorithm; }
	}

	public final EncryptAlgorithmContext encryptAlgorithm() throws RecognitionException {
		EncryptAlgorithmContext _localctx = new EncryptAlgorithmContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_encryptAlgorithm);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			match(ENCRYPT_ALGORITHM);
			setState(227);
			match(LP_);
			setState(228);
			algorithmDefinition();
			setState(229);
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

	public static class AssistedQueryAlgorithmContext extends ParserRuleContext {
		public TerminalNode ASSISTED_QUERY_ALGORITHM() { return getToken(RDLStatementParser.ASSISTED_QUERY_ALGORITHM, 0); }
		public TerminalNode LP_() { return getToken(RDLStatementParser.LP_, 0); }
		public AlgorithmDefinitionContext algorithmDefinition() {
			return getRuleContext(AlgorithmDefinitionContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RDLStatementParser.RP_, 0); }
		public AssistedQueryAlgorithmContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assistedQueryAlgorithm; }
	}

	public final AssistedQueryAlgorithmContext assistedQueryAlgorithm() throws RecognitionException {
		AssistedQueryAlgorithmContext _localctx = new AssistedQueryAlgorithmContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_assistedQueryAlgorithm);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			match(ASSISTED_QUERY_ALGORITHM);
			setState(232);
			match(LP_);
			setState(233);
			algorithmDefinition();
			setState(234);
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

	public static class LikeQueryAlgorithmContext extends ParserRuleContext {
		public TerminalNode LIKE_QUERY_ALGORITHM() { return getToken(RDLStatementParser.LIKE_QUERY_ALGORITHM, 0); }
		public TerminalNode LP_() { return getToken(RDLStatementParser.LP_, 0); }
		public AlgorithmDefinitionContext algorithmDefinition() {
			return getRuleContext(AlgorithmDefinitionContext.class,0);
		}
		public TerminalNode RP_() { return getToken(RDLStatementParser.RP_, 0); }
		public LikeQueryAlgorithmContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_likeQueryAlgorithm; }
	}

	public final LikeQueryAlgorithmContext likeQueryAlgorithm() throws RecognitionException {
		LikeQueryAlgorithmContext _localctx = new LikeQueryAlgorithmContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_likeQueryAlgorithm);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			match(LIKE_QUERY_ALGORITHM);
			setState(237);
			match(LP_);
			setState(238);
			algorithmDefinition();
			setState(239);
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

	public static class QueryWithCipherColumnContext extends ParserRuleContext {
		public TerminalNode TRUE() { return getToken(RDLStatementParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(RDLStatementParser.FALSE, 0); }
		public QueryWithCipherColumnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_queryWithCipherColumn; }
	}

	public final QueryWithCipherColumnContext queryWithCipherColumn() throws RecognitionException {
		QueryWithCipherColumnContext _localctx = new QueryWithCipherColumnContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_queryWithCipherColumn);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241);
			_la = _input.LA(1);
			if ( !(_la==TRUE || _la==FALSE) ) {
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
		enterRule(_localctx, 44, RULE_ifExists);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			match(IF);
			setState(244);
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
		enterRule(_localctx, 46, RULE_ifNotExists);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			match(IF);
			setState(247);
			match(NOT);
			setState(248);
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
		enterRule(_localctx, 48, RULE_literal);
		int _la;
		try {
			setState(257);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_:
				enterOuterAlt(_localctx, 1);
				{
				setState(250);
				match(STRING_);
				}
				break;
			case MINUS_:
			case INT_:
				enterOuterAlt(_localctx, 2);
				{
				setState(252);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MINUS_) {
					{
					setState(251);
					match(MINUS_);
					}
				}

				setState(254);
				match(INT_);
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(255);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(256);
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
		enterRule(_localctx, 50, RULE_algorithmDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			match(TYPE);
			setState(260);
			match(LP_);
			setState(261);
			match(NAME);
			setState(262);
			match(EQ_);
			setState(263);
			algorithmTypeName();
			setState(266);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA_) {
				{
				setState(264);
				match(COMMA_);
				setState(265);
				propertiesDefinition();
				}
			}

			setState(268);
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
		public TerminalNode STRING_() { return getToken(RDLStatementParser.STRING_, 0); }
		public AlgorithmTypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_algorithmTypeName; }
	}

	public final AlgorithmTypeNameContext algorithmTypeName() throws RecognitionException {
		AlgorithmTypeNameContext _localctx = new AlgorithmTypeNameContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_algorithmTypeName);
		try {
			setState(272);
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
				setState(270);
				buildinAlgorithmTypeName();
				}
				break;
			case STRING_:
				enterOuterAlt(_localctx, 2);
				{
				setState(271);
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
		public TerminalNode MD5() { return getToken(RDLStatementParser.MD5, 0); }
		public TerminalNode AES() { return getToken(RDLStatementParser.AES, 0); }
		public TerminalNode RC4() { return getToken(RDLStatementParser.RC4, 0); }
		public TerminalNode SM3() { return getToken(RDLStatementParser.SM3, 0); }
		public TerminalNode SM4() { return getToken(RDLStatementParser.SM4, 0); }
		public TerminalNode CHAR_DIGEST_LIKE() { return getToken(RDLStatementParser.CHAR_DIGEST_LIKE, 0); }
		public BuildinAlgorithmTypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_buildinAlgorithmTypeName; }
	}

	public final BuildinAlgorithmTypeNameContext buildinAlgorithmTypeName() throws RecognitionException {
		BuildinAlgorithmTypeNameContext _localctx = new BuildinAlgorithmTypeNameContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_buildinAlgorithmTypeName);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
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
		enterRule(_localctx, 56, RULE_propertiesDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			match(PROPERTIES);
			setState(277);
			match(LP_);
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING_) {
				{
				setState(278);
				properties();
				}
			}

			setState(281);
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
		enterRule(_localctx, 58, RULE_properties);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			property();
			setState(288);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA_) {
				{
				{
				setState(284);
				match(COMMA_);
				setState(285);
				property();
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
		enterRule(_localctx, 60, RULE_property);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(291);
			((PropertyContext)_localctx).key = match(STRING_);
			setState(292);
			match(EQ_);
			setState(293);
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
		public TerminalNode IDENTIFIER_() { return getToken(RDLStatementParser.IDENTIFIER_, 0); }
		public TableNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableName; }
	}

	public final TableNameContext tableName() throws RecognitionException {
		TableNameContext _localctx = new TableNameContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_tableName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3]\u012c\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\3\2\3\2\3\2\3\2\5\2G\n\2\3\2\3\2\3\2\7\2L\n\2\f\2\16\2O\13\2\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\7\3W\n\3\f\3\16\3Z\13\3\3\4\3\4\3\4\3\4\5\4`\n\4\3"+
		"\4\3\4\3\4\7\4e\n\4\f\4\16\4h\13\4\3\5\3\5\3\5\3\5\3\5\5\5o\n\5\3\5\3"+
		"\5\3\5\3\5\3\5\7\5v\n\5\f\5\16\5y\13\5\3\5\3\5\3\5\3\5\3\5\5\5\u0080\n"+
		"\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b\5\b\u008e\n\b\3\b\3"+
		"\b\3\b\3\b\5\b\u0094\n\b\3\b\3\b\5\b\u0098\n\b\3\b\3\b\3\b\3\b\5\b\u009e"+
		"\n\b\3\b\3\b\5\b\u00a2\n\b\3\b\3\b\3\b\3\b\5\b\u00a8\n\b\3\b\3\b\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\5\t\u00b3\n\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\5\f\u00c0\n\f\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\5\16\u00cb\n\16\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u00d6"+
		"\n\20\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u00e1\n\22\3\23"+
		"\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\32\3\32\5\32"+
		"\u00ff\n\32\3\32\3\32\3\32\5\32\u0104\n\32\3\33\3\33\3\33\3\33\3\33\3"+
		"\33\3\33\5\33\u010d\n\33\3\33\3\33\3\34\3\34\5\34\u0113\n\34\3\35\3\35"+
		"\3\36\3\36\3\36\5\36\u011a\n\36\3\36\3\36\3\37\3\37\3\37\7\37\u0121\n"+
		"\37\f\37\16\37\u0124\13\37\3 \3 \3 \3 \3!\3!\3!\2\2\"\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@\2\4\3\2EF\3\2OT\2\u0126"+
		"\2B\3\2\2\2\4P\3\2\2\2\6[\3\2\2\2\bi\3\2\2\2\n\u0083\3\2\2\2\f\u0087\3"+
		"\2\2\2\16\u0089\3\2\2\2\20\u00ab\3\2\2\2\22\u00b4\3\2\2\2\24\u00b6\3\2"+
		"\2\2\26\u00b8\3\2\2\2\30\u00c1\3\2\2\2\32\u00c3\3\2\2\2\34\u00cc\3\2\2"+
		"\2\36\u00ce\3\2\2\2 \u00d7\3\2\2\2\"\u00d9\3\2\2\2$\u00e2\3\2\2\2&\u00e4"+
		"\3\2\2\2(\u00e9\3\2\2\2*\u00ee\3\2\2\2,\u00f3\3\2\2\2.\u00f5\3\2\2\2\60"+
		"\u00f8\3\2\2\2\62\u0103\3\2\2\2\64\u0105\3\2\2\2\66\u0112\3\2\2\28\u0114"+
		"\3\2\2\2:\u0116\3\2\2\2<\u011d\3\2\2\2>\u0125\3\2\2\2@\u0129\3\2\2\2B"+
		"C\7.\2\2CD\7\65\2\2DF\7\63\2\2EG\5\60\31\2FE\3\2\2\2FG\3\2\2\2GH\3\2\2"+
		"\2HM\5\b\5\2IJ\7$\2\2JL\5\b\5\2KI\3\2\2\2LO\3\2\2\2MK\3\2\2\2MN\3\2\2"+
		"\2N\3\3\2\2\2OM\3\2\2\2PQ\7/\2\2QR\7\65\2\2RS\7\63\2\2SX\5\b\5\2TU\7$"+
		"\2\2UW\5\b\5\2VT\3\2\2\2WZ\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y\5\3\2\2\2ZX\3"+
		"\2\2\2[\\\7\60\2\2\\]\7\65\2\2]_\7\63\2\2^`\5.\30\2_^\3\2\2\2_`\3\2\2"+
		"\2`a\3\2\2\2af\5@!\2bc\7$\2\2ce\5@!\2db\3\2\2\2eh\3\2\2\2fd\3\2\2\2fg"+
		"\3\2\2\2g\7\3\2\2\2hf\3\2\2\2ij\5@!\2jn\7\36\2\2kl\5\n\6\2lm\7$\2\2mo"+
		"\3\2\2\2nk\3\2\2\2no\3\2\2\2op\3\2\2\2pq\7?\2\2qr\7\36\2\2rw\5\16\b\2"+
		"st\7$\2\2tv\5\16\b\2us\3\2\2\2vy\3\2\2\2wu\3\2\2\2wx\3\2\2\2xz\3\2\2\2"+
		"yw\3\2\2\2z\177\7\37\2\2{|\7$\2\2|}\7D\2\2}~\7\27\2\2~\u0080\5,\27\2\177"+
		"{\3\2\2\2\177\u0080\3\2\2\2\u0080\u0081\3\2\2\2\u0081\u0082\7\37\2\2\u0082"+
		"\t\3\2\2\2\u0083\u0084\7\62\2\2\u0084\u0085\7\27\2\2\u0085\u0086\5\f\7"+
		"\2\u0086\13\3\2\2\2\u0087\u0088\7W\2\2\u0088\r\3\2\2\2\u0089\u008a\7\36"+
		"\2\2\u008a\u008d\5\20\t\2\u008b\u008c\7$\2\2\u008c\u008e\5\26\f\2\u008d"+
		"\u008b\3\2\2\2\u008d\u008e\3\2\2\2\u008e\u008f\3\2\2\2\u008f\u0090\7$"+
		"\2\2\u0090\u0093\5\32\16\2\u0091\u0092\7$\2\2\u0092\u0094\5\36\20\2\u0093"+
		"\u0091\3\2\2\2\u0093\u0094\3\2\2\2\u0094\u0097\3\2\2\2\u0095\u0096\7$"+
		"\2\2\u0096\u0098\5\"\22\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098"+
		"\u0099\3\2\2\2\u0099\u009a\7$\2\2\u009a\u009d\5&\24\2\u009b\u009c\7$\2"+
		"\2\u009c\u009e\5(\25\2\u009d\u009b\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u00a1"+
		"\3\2\2\2\u009f\u00a0\7$\2\2\u00a0\u00a2\5*\26\2\u00a1\u009f\3\2\2\2\u00a1"+
		"\u00a2\3\2\2\2\u00a2\u00a7\3\2\2\2\u00a3\u00a4\7$\2\2\u00a4\u00a5\7D\2"+
		"\2\u00a5\u00a6\7\27\2\2\u00a6\u00a8\5,\27\2\u00a7\u00a3\3\2\2\2\u00a7"+
		"\u00a8\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00aa\7\37\2\2\u00aa\17\3\2\2"+
		"\2\u00ab\u00ac\7:\2\2\u00ac\u00ad\7\27\2\2\u00ad\u00b2\5\22\n\2\u00ae"+
		"\u00af\7$\2\2\u00af\u00b0\7G\2\2\u00b0\u00b1\7\27\2\2\u00b1\u00b3\5\24"+
		"\13\2\u00b2\u00ae\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3\21\3\2\2\2\u00b4\u00b5"+
		"\7W\2\2\u00b5\23\3\2\2\2\u00b6\u00b7\7X\2\2\u00b7\25\3\2\2\2\u00b8\u00b9"+
		"\7A\2\2\u00b9\u00ba\7\27\2\2\u00ba\u00bf\5\30\r\2\u00bb\u00bc\7$\2\2\u00bc"+
		"\u00bd\7H\2\2\u00bd\u00be\7\27\2\2\u00be\u00c0\5\24\13\2\u00bf\u00bb\3"+
		"\2\2\2\u00bf\u00c0\3\2\2\2\u00c0\27\3\2\2\2\u00c1\u00c2\7W\2\2\u00c2\31"+
		"\3\2\2\2\u00c3\u00c4\7@\2\2\u00c4\u00c5\7\27\2\2\u00c5\u00ca\5\34\17\2"+
		"\u00c6\u00c7\7$\2\2\u00c7\u00c8\7I\2\2\u00c8\u00c9\7\27\2\2\u00c9\u00cb"+
		"\5\24\13\2\u00ca\u00c6\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb\33\3\2\2\2\u00cc"+
		"\u00cd\7W\2\2\u00cd\35\3\2\2\2\u00ce\u00cf\7B\2\2\u00cf\u00d0\7\27\2\2"+
		"\u00d0\u00d5\5 \21\2\u00d1\u00d2\7$\2\2\u00d2\u00d3\7J\2\2\u00d3\u00d4"+
		"\7\27\2\2\u00d4\u00d6\5\24\13\2\u00d5\u00d1\3\2\2\2\u00d5\u00d6\3\2\2"+
		"\2\u00d6\37\3\2\2\2\u00d7\u00d8\7W\2\2\u00d8!\3\2\2\2\u00d9\u00da\7C\2"+
		"\2\u00da\u00db\7\27\2\2\u00db\u00e0\5$\23\2\u00dc\u00dd\7$\2\2\u00dd\u00de"+
		"\7K\2\2\u00de\u00df\7\27\2\2\u00df\u00e1\5\24\13\2\u00e0\u00dc\3\2\2\2"+
		"\u00e0\u00e1\3\2\2\2\u00e1#\3\2\2\2\u00e2\u00e3\7W\2\2\u00e3%\3\2\2\2"+
		"\u00e4\u00e5\7\67\2\2\u00e5\u00e6\7\36\2\2\u00e6\u00e7\5\64\33\2\u00e7"+
		"\u00e8\7\37\2\2\u00e8\'\3\2\2\2\u00e9\u00ea\78\2\2\u00ea\u00eb\7\36\2"+
		"\2\u00eb\u00ec\5\64\33\2\u00ec\u00ed\7\37\2\2\u00ed)\3\2\2\2\u00ee\u00ef"+
		"\79\2\2\u00ef\u00f0\7\36\2\2\u00f0\u00f1\5\64\33\2\u00f1\u00f2\7\37\2"+
		"\2\u00f2+\3\2\2\2\u00f3\u00f4\t\2\2\2\u00f4-\3\2\2\2\u00f5\u00f6\7L\2"+
		"\2\u00f6\u00f7\7M\2\2\u00f7/\3\2\2\2\u00f8\u00f9\7L\2\2\u00f9\u00fa\7"+
		"U\2\2\u00fa\u00fb\7M\2\2\u00fb\61\3\2\2\2\u00fc\u0104\7X\2\2\u00fd\u00ff"+
		"\7\17\2\2\u00fe\u00fd\3\2\2\2\u00fe\u00ff\3\2\2\2\u00ff\u0100\3\2\2\2"+
		"\u0100\u0104\7Y\2\2\u0101\u0104\7E\2\2\u0102\u0104\7F\2\2\u0103\u00fc"+
		"\3\2\2\2\u0103\u00fe\3\2\2\2\u0103\u0101\3\2\2\2\u0103\u0102\3\2\2\2\u0104"+
		"\63\3\2\2\2\u0105\u0106\7\66\2\2\u0106\u0107\7\36\2\2\u0107\u0108\7:\2"+
		"\2\u0108\u0109\7\27\2\2\u0109\u010c\5\66\34\2\u010a\u010b\7$\2\2\u010b"+
		"\u010d\5:\36\2\u010c\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d\u010e\3\2"+
		"\2\2\u010e\u010f\7\37\2\2\u010f\65\3\2\2\2\u0110\u0113\58\35\2\u0111\u0113"+
		"\7X\2\2\u0112\u0110\3\2\2\2\u0112\u0111\3\2\2\2\u0113\67\3\2\2\2\u0114"+
		"\u0115\t\3\2\2\u01159\3\2\2\2\u0116\u0117\7;\2\2\u0117\u0119\7\36\2\2"+
		"\u0118\u011a\5<\37\2\u0119\u0118\3\2\2\2\u0119\u011a\3\2\2\2\u011a\u011b"+
		"\3\2\2\2\u011b\u011c\7\37\2\2\u011c;\3\2\2\2\u011d\u0122\5> \2\u011e\u011f"+
		"\7$\2\2\u011f\u0121\5> \2\u0120\u011e\3\2\2\2\u0121\u0124\3\2\2\2\u0122"+
		"\u0120\3\2\2\2\u0122\u0123\3\2\2\2\u0123=\3\2\2\2\u0124\u0122\3\2\2\2"+
		"\u0125\u0126\7X\2\2\u0126\u0127\7\27\2\2\u0127\u0128\5\62\32\2\u0128?"+
		"\3\2\2\2\u0129\u012a\7W\2\2\u012aA\3\2\2\2\33FMX_fnw\177\u008d\u0093\u0097"+
		"\u009d\u00a1\u00a7\u00b2\u00bf\u00ca\u00d5\u00e0\u00fe\u0103\u010c\u0112"+
		"\u0119\u0122";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}