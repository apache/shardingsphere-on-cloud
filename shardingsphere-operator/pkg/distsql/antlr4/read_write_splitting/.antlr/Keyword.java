// Generated from /root/workspaces/golang/src/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/antlr4/read_write_splitting/Keyword.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Keyword extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, TRUE=2, FALSE=3, CREATE=4, ALTER=5, DROP=6, SHOW=7, RULE=8, FROM=9, 
		READWRITE_SPLITTING=10, WRITE_STORAGE_UNIT=11, READ_STORAGE_UNITS=12, 
		TRANSACTIONAL_READ_QUERY_STRATEGY=13, TYPE=14, NAME=15, PROPERTIES=16, 
		RULES=17, RESOURCES=18, STATUS=19, ENABLE=20, DISABLE=21, READ=22, IF=23, 
		EXISTS=24, COUNT=25, ROUND_ROBIN=26, RANDOM=27, WEIGHT=28, NOT=29, FOR_GENERATOR=30;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"WS", "TRUE", "FALSE", "CREATE", "ALTER", "DROP", "SHOW", "RULE", "FROM", 
			"READWRITE_SPLITTING", "WRITE_STORAGE_UNIT", "READ_STORAGE_UNITS", "TRANSACTIONAL_READ_QUERY_STRATEGY", 
			"TYPE", "NAME", "PROPERTIES", "RULES", "RESOURCES", "STATUS", "ENABLE", 
			"DISABLE", "READ", "IF", "EXISTS", "COUNT", "ROUND_ROBIN", "RANDOM", 
			"WEIGHT", "NOT", "FOR_GENERATOR", "A", "B", "C", "D", "E", "F", "G", 
			"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", 
			"V", "W", "X", "Y", "Z", "UL_"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "TRUE", "FALSE", "CREATE", "ALTER", "DROP", "SHOW", "RULE", 
			"FROM", "READWRITE_SPLITTING", "WRITE_STORAGE_UNIT", "READ_STORAGE_UNITS", 
			"TRANSACTIONAL_READ_QUERY_STRATEGY", "TYPE", "NAME", "PROPERTIES", "RULES", 
			"RESOURCES", "STATUS", "ENABLE", "DISABLE", "READ", "IF", "EXISTS", "COUNT", 
			"ROUND_ROBIN", "RANDOM", "WEIGHT", "NOT", "FOR_GENERATOR"
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


	public Keyword(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Keyword.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2 \u01d3\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\3\2\6\2w\n\2\r\2\16"+
		"\2x\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\27"+
		"\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3"+
		"#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3."+
		"\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\66"+
		"\3\66\3\67\3\67\38\38\39\39\3:\3:\2\2;\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21"+
		"\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30"+
		"/\31\61\32\63\33\65\34\67\359\36;\37= ?\2A\2C\2E\2G\2I\2K\2M\2O\2Q\2S"+
		"\2U\2W\2Y\2[\2]\2_\2a\2c\2e\2g\2i\2k\2m\2o\2q\2s\2\3\2\35\5\2\13\f\17"+
		"\17\"\"\4\2CCcc\4\2DDdd\4\2EEee\4\2FFff\4\2GGgg\4\2HHhh\4\2IIii\4\2JJ"+
		"jj\4\2KKkk\4\2LLll\4\2MMmm\4\2NNnn\4\2OOoo\4\2PPpp\4\2QQqq\4\2RRrr\4\2"+
		"SSss\4\2TTtt\4\2UUuu\4\2VVvv\4\2WWww\4\2XXxx\4\2YYyy\4\2ZZzz\4\2[[{{\4"+
		"\2\\\\||\2\u01b8\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13"+
		"\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2"+
		"\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2"+
		"!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3"+
		"\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2"+
		"\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\3v\3\2\2\2\5|\3\2\2\2\7\u0081\3\2\2"+
		"\2\t\u0087\3\2\2\2\13\u008e\3\2\2\2\r\u0094\3\2\2\2\17\u0099\3\2\2\2\21"+
		"\u009e\3\2\2\2\23\u00a3\3\2\2\2\25\u00a8\3\2\2\2\27\u00bc\3\2\2\2\31\u00cf"+
		"\3\2\2\2\33\u00e2\3\2\2\2\35\u0104\3\2\2\2\37\u0109\3\2\2\2!\u010e\3\2"+
		"\2\2#\u0119\3\2\2\2%\u011f\3\2\2\2\'\u0129\3\2\2\2)\u0130\3\2\2\2+\u0137"+
		"\3\2\2\2-\u013f\3\2\2\2/\u0144\3\2\2\2\61\u0147\3\2\2\2\63\u014e\3\2\2"+
		"\2\65\u0154\3\2\2\2\67\u0160\3\2\2\29\u0167\3\2\2\2;\u016e\3\2\2\2=\u0172"+
		"\3\2\2\2?\u019d\3\2\2\2A\u019f\3\2\2\2C\u01a1\3\2\2\2E\u01a3\3\2\2\2G"+
		"\u01a5\3\2\2\2I\u01a7\3\2\2\2K\u01a9\3\2\2\2M\u01ab\3\2\2\2O\u01ad\3\2"+
		"\2\2Q\u01af\3\2\2\2S\u01b1\3\2\2\2U\u01b3\3\2\2\2W\u01b5\3\2\2\2Y\u01b7"+
		"\3\2\2\2[\u01b9\3\2\2\2]\u01bb\3\2\2\2_\u01bd\3\2\2\2a\u01bf\3\2\2\2c"+
		"\u01c1\3\2\2\2e\u01c3\3\2\2\2g\u01c5\3\2\2\2i\u01c7\3\2\2\2k\u01c9\3\2"+
		"\2\2m\u01cb\3\2\2\2o\u01cd\3\2\2\2q\u01cf\3\2\2\2s\u01d1\3\2\2\2uw\t\2"+
		"\2\2vu\3\2\2\2wx\3\2\2\2xv\3\2\2\2xy\3\2\2\2yz\3\2\2\2z{\b\2\2\2{\4\3"+
		"\2\2\2|}\5e\63\2}~\5a\61\2~\177\5g\64\2\177\u0080\5G$\2\u0080\6\3\2\2"+
		"\2\u0081\u0082\5I%\2\u0082\u0083\5? \2\u0083\u0084\5U+\2\u0084\u0085\5"+
		"c\62\2\u0085\u0086\5G$\2\u0086\b\3\2\2\2\u0087\u0088\5C\"\2\u0088\u0089"+
		"\5a\61\2\u0089\u008a\5G$\2\u008a\u008b\5? \2\u008b\u008c\5e\63\2\u008c"+
		"\u008d\5G$\2\u008d\n\3\2\2\2\u008e\u008f\5? \2\u008f\u0090\5U+\2\u0090"+
		"\u0091\5e\63\2\u0091\u0092\5G$\2\u0092\u0093\5a\61\2\u0093\f\3\2\2\2\u0094"+
		"\u0095\5E#\2\u0095\u0096\5a\61\2\u0096\u0097\5[.\2\u0097\u0098\5]/\2\u0098"+
		"\16\3\2\2\2\u0099\u009a\5c\62\2\u009a\u009b\5M\'\2\u009b\u009c\5[.\2\u009c"+
		"\u009d\5k\66\2\u009d\20\3\2\2\2\u009e\u009f\5a\61\2\u009f\u00a0\5g\64"+
		"\2\u00a0\u00a1\5U+\2\u00a1\u00a2\5G$\2\u00a2\22\3\2\2\2\u00a3\u00a4\5"+
		"I%\2\u00a4\u00a5\5a\61\2\u00a5\u00a6\5[.\2\u00a6\u00a7\5W,\2\u00a7\24"+
		"\3\2\2\2\u00a8\u00a9\5a\61\2\u00a9\u00aa\5G$\2\u00aa\u00ab\5? \2\u00ab"+
		"\u00ac\5E#\2\u00ac\u00ad\5k\66\2\u00ad\u00ae\5a\61\2\u00ae\u00af\5O(\2"+
		"\u00af\u00b0\5e\63\2\u00b0\u00b1\5G$\2\u00b1\u00b2\5s:\2\u00b2\u00b3\5"+
		"c\62\2\u00b3\u00b4\5]/\2\u00b4\u00b5\5U+\2\u00b5\u00b6\5O(\2\u00b6\u00b7"+
		"\5e\63\2\u00b7\u00b8\5e\63\2\u00b8\u00b9\5O(\2\u00b9\u00ba\5Y-\2\u00ba"+
		"\u00bb\5K&\2\u00bb\26\3\2\2\2\u00bc\u00bd\5k\66\2\u00bd\u00be\5a\61\2"+
		"\u00be\u00bf\5O(\2\u00bf\u00c0\5e\63\2\u00c0\u00c1\5G$\2\u00c1\u00c2\5"+
		"s:\2\u00c2\u00c3\5c\62\2\u00c3\u00c4\5e\63\2\u00c4\u00c5\5[.\2\u00c5\u00c6"+
		"\5a\61\2\u00c6\u00c7\5? \2\u00c7\u00c8\5K&\2\u00c8\u00c9\5G$\2\u00c9\u00ca"+
		"\5s:\2\u00ca\u00cb\5g\64\2\u00cb\u00cc\5Y-\2\u00cc\u00cd\5O(\2\u00cd\u00ce"+
		"\5e\63\2\u00ce\30\3\2\2\2\u00cf\u00d0\5a\61\2\u00d0\u00d1\5G$\2\u00d1"+
		"\u00d2\5? \2\u00d2\u00d3\5E#\2\u00d3\u00d4\5s:\2\u00d4\u00d5\5c\62\2\u00d5"+
		"\u00d6\5e\63\2\u00d6\u00d7\5[.\2\u00d7\u00d8\5a\61\2\u00d8\u00d9\5? \2"+
		"\u00d9\u00da\5K&\2\u00da\u00db\5G$\2\u00db\u00dc\5s:\2\u00dc\u00dd\5g"+
		"\64\2\u00dd\u00de\5Y-\2\u00de\u00df\5O(\2\u00df\u00e0\5e\63\2\u00e0\u00e1"+
		"\5c\62\2\u00e1\32\3\2\2\2\u00e2\u00e3\5e\63\2\u00e3\u00e4\5a\61\2\u00e4"+
		"\u00e5\5? \2\u00e5\u00e6\5Y-\2\u00e6\u00e7\5c\62\2\u00e7\u00e8\5? \2\u00e8"+
		"\u00e9\5C\"\2\u00e9\u00ea\5e\63\2\u00ea\u00eb\5O(\2\u00eb\u00ec\5[.\2"+
		"\u00ec\u00ed\5Y-\2\u00ed\u00ee\5? \2\u00ee\u00ef\5U+\2\u00ef\u00f0\5s"+
		":\2\u00f0\u00f1\5a\61\2\u00f1\u00f2\5G$\2\u00f2\u00f3\5? \2\u00f3\u00f4"+
		"\5E#\2\u00f4\u00f5\5s:\2\u00f5\u00f6\5_\60\2\u00f6\u00f7\5g\64\2\u00f7"+
		"\u00f8\5G$\2\u00f8\u00f9\5a\61\2\u00f9\u00fa\5o8\2\u00fa\u00fb\5s:\2\u00fb"+
		"\u00fc\5c\62\2\u00fc\u00fd\5e\63\2\u00fd\u00fe\5a\61\2\u00fe\u00ff\5?"+
		" \2\u00ff\u0100\5e\63\2\u0100\u0101\5G$\2\u0101\u0102\5K&\2\u0102\u0103"+
		"\5o8\2\u0103\34\3\2\2\2\u0104\u0105\5e\63\2\u0105\u0106\5o8\2\u0106\u0107"+
		"\5]/\2\u0107\u0108\5G$\2\u0108\36\3\2\2\2\u0109\u010a\5Y-\2\u010a\u010b"+
		"\5? \2\u010b\u010c\5W,\2\u010c\u010d\5G$\2\u010d \3\2\2\2\u010e\u010f"+
		"\5]/\2\u010f\u0110\5a\61\2\u0110\u0111\5[.\2\u0111\u0112\5]/\2\u0112\u0113"+
		"\5G$\2\u0113\u0114\5a\61\2\u0114\u0115\5e\63\2\u0115\u0116\5O(\2\u0116"+
		"\u0117\5G$\2\u0117\u0118\5c\62\2\u0118\"\3\2\2\2\u0119\u011a\5a\61\2\u011a"+
		"\u011b\5g\64\2\u011b\u011c\5U+\2\u011c\u011d\5G$\2\u011d\u011e\5c\62\2"+
		"\u011e$\3\2\2\2\u011f\u0120\5a\61\2\u0120\u0121\5G$\2\u0121\u0122\5c\62"+
		"\2\u0122\u0123\5[.\2\u0123\u0124\5g\64\2\u0124\u0125\5a\61\2\u0125\u0126"+
		"\5C\"\2\u0126\u0127\5G$\2\u0127\u0128\5c\62\2\u0128&\3\2\2\2\u0129\u012a"+
		"\5c\62\2\u012a\u012b\5e\63\2\u012b\u012c\5? \2\u012c\u012d\5e\63\2\u012d"+
		"\u012e\5g\64\2\u012e\u012f\5c\62\2\u012f(\3\2\2\2\u0130\u0131\5G$\2\u0131"+
		"\u0132\5Y-\2\u0132\u0133\5? \2\u0133\u0134\5A!\2\u0134\u0135\5U+\2\u0135"+
		"\u0136\5G$\2\u0136*\3\2\2\2\u0137\u0138\5E#\2\u0138\u0139\5O(\2\u0139"+
		"\u013a\5c\62\2\u013a\u013b\5? \2\u013b\u013c\5A!\2\u013c\u013d\5U+\2\u013d"+
		"\u013e\5G$\2\u013e,\3\2\2\2\u013f\u0140\5a\61\2\u0140\u0141\5G$\2\u0141"+
		"\u0142\5? \2\u0142\u0143\5E#\2\u0143.\3\2\2\2\u0144\u0145\5O(\2\u0145"+
		"\u0146\5I%\2\u0146\60\3\2\2\2\u0147\u0148\5G$\2\u0148\u0149\5m\67\2\u0149"+
		"\u014a\5O(\2\u014a\u014b\5c\62\2\u014b\u014c\5e\63\2\u014c\u014d\5c\62"+
		"\2\u014d\62\3\2\2\2\u014e\u014f\5C\"\2\u014f\u0150\5[.\2\u0150\u0151\5"+
		"g\64\2\u0151\u0152\5Y-\2\u0152\u0153\5e\63\2\u0153\64\3\2\2\2\u0154\u0155"+
		"\5a\61\2\u0155\u0156\5[.\2\u0156\u0157\5g\64\2\u0157\u0158\5Y-\2\u0158"+
		"\u0159\5E#\2\u0159\u015a\5s:\2\u015a\u015b\5a\61\2\u015b\u015c\5[.\2\u015c"+
		"\u015d\5A!\2\u015d\u015e\5O(\2\u015e\u015f\5Y-\2\u015f\66\3\2\2\2\u0160"+
		"\u0161\5a\61\2\u0161\u0162\5? \2\u0162\u0163\5Y-\2\u0163\u0164\5E#\2\u0164"+
		"\u0165\5[.\2\u0165\u0166\5W,\2\u01668\3\2\2\2\u0167\u0168\5k\66\2\u0168"+
		"\u0169\5G$\2\u0169\u016a\5O(\2\u016a\u016b\5K&\2\u016b\u016c\5M\'\2\u016c"+
		"\u016d\5e\63\2\u016d:\3\2\2\2\u016e\u016f\5Y-\2\u016f\u0170\5[.\2\u0170"+
		"\u0171\5e\63\2\u0171<\3\2\2\2\u0172\u0173\7F\2\2\u0173\u0174\7Q\2\2\u0174"+
		"\u0175\7\"\2\2\u0175\u0176\7P\2\2\u0176\u0177\7Q\2\2\u0177\u0178\7V\2"+
		"\2\u0178\u0179\7\"\2\2\u0179\u017a\7O\2\2\u017a\u017b\7C\2\2\u017b\u017c"+
		"\7V\2\2\u017c\u017d\7E\2\2\u017d\u017e\7J\2\2\u017e\u017f\7\"\2\2\u017f"+
		"\u0180\7C\2\2\u0180\u0181\7P\2\2\u0181\u0182\7[\2\2\u0182\u0183\7\"\2"+
		"\2\u0183\u0184\7V\2\2\u0184\u0185\7J\2\2\u0185\u0186\7K\2\2\u0186\u0187"+
		"\7P\2\2\u0187\u0188\7I\2\2\u0188\u0189\7.\2\2\u0189\u018a\7\"\2\2\u018a"+
		"\u018b\7L\2\2\u018b\u018c\7W\2\2\u018c\u018d\7U\2\2\u018d\u018e\7V\2\2"+
		"\u018e\u018f\7\"\2\2\u018f\u0190\7H\2\2\u0190\u0191\7Q\2\2\u0191\u0192"+
		"\7T\2\2\u0192\u0193\7\"\2\2\u0193\u0194\7I\2\2\u0194\u0195\7G\2\2\u0195"+
		"\u0196\7P\2\2\u0196\u0197\7G\2\2\u0197\u0198\7T\2\2\u0198\u0199\7C\2\2"+
		"\u0199\u019a\7V\2\2\u019a\u019b\7Q\2\2\u019b\u019c\7T\2\2\u019c>\3\2\2"+
		"\2\u019d\u019e\t\3\2\2\u019e@\3\2\2\2\u019f\u01a0\t\4\2\2\u01a0B\3\2\2"+
		"\2\u01a1\u01a2\t\5\2\2\u01a2D\3\2\2\2\u01a3\u01a4\t\6\2\2\u01a4F\3\2\2"+
		"\2\u01a5\u01a6\t\7\2\2\u01a6H\3\2\2\2\u01a7\u01a8\t\b\2\2\u01a8J\3\2\2"+
		"\2\u01a9\u01aa\t\t\2\2\u01aaL\3\2\2\2\u01ab\u01ac\t\n\2\2\u01acN\3\2\2"+
		"\2\u01ad\u01ae\t\13\2\2\u01aeP\3\2\2\2\u01af\u01b0\t\f\2\2\u01b0R\3\2"+
		"\2\2\u01b1\u01b2\t\r\2\2\u01b2T\3\2\2\2\u01b3\u01b4\t\16\2\2\u01b4V\3"+
		"\2\2\2\u01b5\u01b6\t\17\2\2\u01b6X\3\2\2\2\u01b7\u01b8\t\20\2\2\u01b8"+
		"Z\3\2\2\2\u01b9\u01ba\t\21\2\2\u01ba\\\3\2\2\2\u01bb\u01bc\t\22\2\2\u01bc"+
		"^\3\2\2\2\u01bd\u01be\t\23\2\2\u01be`\3\2\2\2\u01bf\u01c0\t\24\2\2\u01c0"+
		"b\3\2\2\2\u01c1\u01c2\t\25\2\2\u01c2d\3\2\2\2\u01c3\u01c4\t\26\2\2\u01c4"+
		"f\3\2\2\2\u01c5\u01c6\t\27\2\2\u01c6h\3\2\2\2\u01c7\u01c8\t\30\2\2\u01c8"+
		"j\3\2\2\2\u01c9\u01ca\t\31\2\2\u01cal\3\2\2\2\u01cb\u01cc\t\32\2\2\u01cc"+
		"n\3\2\2\2\u01cd\u01ce\t\33\2\2\u01cep\3\2\2\2\u01cf\u01d0\t\34\2\2\u01d0"+
		"r\3\2\2\2\u01d1\u01d2\7a\2\2\u01d2t\3\2\2\2\4\2x\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}