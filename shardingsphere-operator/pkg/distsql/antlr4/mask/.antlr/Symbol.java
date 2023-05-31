// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/mask/Symbol.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Symbol extends Lexer {
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
		SQ_=36, BQ_=37, QUESTION_=38, AT_=39, SEMI_=40, JSONSEPARATOR_=41, UL_=42;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"AND_", "OR_", "NOT_", "TILDE_", "VERTICALBAR_", "AMPERSAND_", "SIGNEDLEFTSHIFT_", 
			"SIGNEDRIGHTSHIFT_", "CARET_", "MOD_", "COLON_", "PLUS_", "MINUS_", "ASTERISK_", 
			"SLASH_", "BACKSLASH_", "DOT_", "DOTASTERISK_", "SAFEEQ_", "DEQ_", "EQ_", 
			"NEQ_", "GT_", "GTE_", "LT_", "LTE_", "POUND_", "LP_", "RP_", "LBE_", 
			"RBE_", "LBT_", "RBT_", "COMMA_", "DQ_", "SQ_", "BQ_", "QUESTION_", "AT_", 
			"SEMI_", "JSONSEPARATOR_", "UL_"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'&&'", "'||'", "'!'", "'~'", "'|'", "'&'", "'<<'", "'>>'", "'^'", 
			"'%'", "':'", "'+'", "'-'", "'*'", "'/'", "'\\'", "'.'", "'.*'", "'<=>'", 
			"'=='", "'='", null, "'>'", "'>='", "'<'", "'<='", "'#'", "'('", "')'", 
			"'{'", "'}'", "'['", "']'", "','", "'\"'", "'''", "'`'", "'?'", "'@'", 
			"';'", "'->>'", "'_'"
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
			"BQ_", "QUESTION_", "AT_", "SEMI_", "JSONSEPARATOR_", "UL_"
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


	public Symbol(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Symbol.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2,\u00bb\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\3"+
		"\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\t"+
		"\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3"+
		"\20\3\21\3\21\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3"+
		"\25\3\26\3\26\3\27\3\27\3\27\3\27\5\27\u008e\n\27\3\30\3\30\3\31\3\31"+
		"\3\31\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37"+
		"\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3"+
		"*\3*\3*\3+\3+\2\2,\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r"+
		"\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33"+
		"\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,\3\2\2\2\u00bb\2\3\3\2"+
		"\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17"+
		"\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2"+
		"\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3"+
		"\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3"+
		"\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2"+
		"=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3"+
		"\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2"+
		"\2\3W\3\2\2\2\5Z\3\2\2\2\7]\3\2\2\2\t_\3\2\2\2\13a\3\2\2\2\rc\3\2\2\2"+
		"\17e\3\2\2\2\21h\3\2\2\2\23k\3\2\2\2\25m\3\2\2\2\27o\3\2\2\2\31q\3\2\2"+
		"\2\33s\3\2\2\2\35u\3\2\2\2\37w\3\2\2\2!y\3\2\2\2#{\3\2\2\2%}\3\2\2\2\'"+
		"\u0080\3\2\2\2)\u0084\3\2\2\2+\u0087\3\2\2\2-\u008d\3\2\2\2/\u008f\3\2"+
		"\2\2\61\u0091\3\2\2\2\63\u0094\3\2\2\2\65\u0096\3\2\2\2\67\u0099\3\2\2"+
		"\29\u009b\3\2\2\2;\u009d\3\2\2\2=\u009f\3\2\2\2?\u00a1\3\2\2\2A\u00a3"+
		"\3\2\2\2C\u00a5\3\2\2\2E\u00a7\3\2\2\2G\u00a9\3\2\2\2I\u00ab\3\2\2\2K"+
		"\u00ad\3\2\2\2M\u00af\3\2\2\2O\u00b1\3\2\2\2Q\u00b3\3\2\2\2S\u00b5\3\2"+
		"\2\2U\u00b9\3\2\2\2WX\7(\2\2XY\7(\2\2Y\4\3\2\2\2Z[\7~\2\2[\\\7~\2\2\\"+
		"\6\3\2\2\2]^\7#\2\2^\b\3\2\2\2_`\7\u0080\2\2`\n\3\2\2\2ab\7~\2\2b\f\3"+
		"\2\2\2cd\7(\2\2d\16\3\2\2\2ef\7>\2\2fg\7>\2\2g\20\3\2\2\2hi\7@\2\2ij\7"+
		"@\2\2j\22\3\2\2\2kl\7`\2\2l\24\3\2\2\2mn\7\'\2\2n\26\3\2\2\2op\7<\2\2"+
		"p\30\3\2\2\2qr\7-\2\2r\32\3\2\2\2st\7/\2\2t\34\3\2\2\2uv\7,\2\2v\36\3"+
		"\2\2\2wx\7\61\2\2x \3\2\2\2yz\7^\2\2z\"\3\2\2\2{|\7\60\2\2|$\3\2\2\2}"+
		"~\7\60\2\2~\177\7,\2\2\177&\3\2\2\2\u0080\u0081\7>\2\2\u0081\u0082\7?"+
		"\2\2\u0082\u0083\7@\2\2\u0083(\3\2\2\2\u0084\u0085\7?\2\2\u0085\u0086"+
		"\7?\2\2\u0086*\3\2\2\2\u0087\u0088\7?\2\2\u0088,\3\2\2\2\u0089\u008a\7"+
		">\2\2\u008a\u008e\7@\2\2\u008b\u008c\7#\2\2\u008c\u008e\7?\2\2\u008d\u0089"+
		"\3\2\2\2\u008d\u008b\3\2\2\2\u008e.\3\2\2\2\u008f\u0090\7@\2\2\u0090\60"+
		"\3\2\2\2\u0091\u0092\7@\2\2\u0092\u0093\7?\2\2\u0093\62\3\2\2\2\u0094"+
		"\u0095\7>\2\2\u0095\64\3\2\2\2\u0096\u0097\7>\2\2\u0097\u0098\7?\2\2\u0098"+
		"\66\3\2\2\2\u0099\u009a\7%\2\2\u009a8\3\2\2\2\u009b\u009c\7*\2\2\u009c"+
		":\3\2\2\2\u009d\u009e\7+\2\2\u009e<\3\2\2\2\u009f\u00a0\7}\2\2\u00a0>"+
		"\3\2\2\2\u00a1\u00a2\7\177\2\2\u00a2@\3\2\2\2\u00a3\u00a4\7]\2\2\u00a4"+
		"B\3\2\2\2\u00a5\u00a6\7_\2\2\u00a6D\3\2\2\2\u00a7\u00a8\7.\2\2\u00a8F"+
		"\3\2\2\2\u00a9\u00aa\7$\2\2\u00aaH\3\2\2\2\u00ab\u00ac\7)\2\2\u00acJ\3"+
		"\2\2\2\u00ad\u00ae\7b\2\2\u00aeL\3\2\2\2\u00af\u00b0\7A\2\2\u00b0N\3\2"+
		"\2\2\u00b1\u00b2\7B\2\2\u00b2P\3\2\2\2\u00b3\u00b4\7=\2\2\u00b4R\3\2\2"+
		"\2\u00b5\u00b6\7/\2\2\u00b6\u00b7\7@\2\2\u00b7\u00b8\7@\2\2\u00b8T\3\2"+
		"\2\2\u00b9\u00ba\7a\2\2\u00baV\3\2\2\2\4\2\u008d\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}