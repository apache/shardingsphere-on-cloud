// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RDLStatement

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 87, 162,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 5, 2, 38, 10, 2, 3, 2, 3, 2, 5, 2, 42, 10, 2, 3, 2, 3, 2, 3, 2,
	7, 2, 47, 10, 2, 12, 2, 14, 2, 50, 11, 2, 3, 3, 3, 3, 3, 3, 5, 3, 55, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 61, 10, 3, 12, 3, 14, 3, 64, 11, 3, 3,
	4, 3, 4, 3, 4, 5, 4, 69, 10, 4, 3, 4, 3, 4, 5, 4, 73, 10, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 78, 10, 4, 12, 4, 14, 4, 81, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 7, 5, 90, 10, 5, 12, 5, 14, 5, 93, 11, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 117, 10, 10, 3,
	10, 3, 10, 3, 10, 5, 10, 122, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 5, 11, 131, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 137,
	10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 144, 10, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 7, 15, 151, 10, 15, 12, 15, 14, 15, 154, 11, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 3, 3, 2, 66, 78, 2,
	162, 2, 34, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6, 65, 3, 2, 2, 2, 8, 82, 3,
	2, 2, 2, 10, 97, 3, 2, 2, 2, 12, 105, 3, 2, 2, 2, 14, 107, 3, 2, 2, 2,
	16, 110, 3, 2, 2, 2, 18, 121, 3, 2, 2, 2, 20, 123, 3, 2, 2, 2, 22, 136,
	3, 2, 2, 2, 24, 138, 3, 2, 2, 2, 26, 140, 3, 2, 2, 2, 28, 147, 3, 2, 2,
	2, 30, 155, 3, 2, 2, 2, 32, 159, 3, 2, 2, 2, 34, 35, 7, 48, 2, 2, 35, 37,
	7, 54, 2, 2, 36, 38, 7, 60, 2, 2, 37, 36, 3, 2, 2, 2, 37, 38, 3, 2, 2,
	2, 38, 39, 3, 2, 2, 2, 39, 41, 7, 52, 2, 2, 40, 42, 5, 16, 9, 2, 41, 40,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 48, 5, 8, 5, 2,
	44, 45, 7, 36, 2, 2, 45, 47, 5, 8, 5, 2, 46, 44, 3, 2, 2, 2, 47, 50, 3,
	2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 3, 3, 2, 2, 2, 50,
	48, 3, 2, 2, 2, 51, 52, 7, 49, 2, 2, 52, 54, 7, 54, 2, 2, 53, 55, 7, 60,
	2, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57,
	7, 52, 2, 2, 57, 62, 5, 8, 5, 2, 58, 59, 7, 36, 2, 2, 59, 61, 5, 8, 5,
	2, 60, 58, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63,
	3, 2, 2, 2, 63, 5, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 7, 50, 2, 2,
	66, 68, 7, 54, 2, 2, 67, 69, 7, 60, 2, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72, 7, 52, 2, 2, 71, 73, 5, 14, 8, 2,
	72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 79, 5,
	32, 17, 2, 75, 76, 7, 36, 2, 2, 76, 78, 5, 32, 17, 2, 77, 75, 3, 2, 2,
	2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 7, 3,
	2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83, 5, 32, 17, 2, 83, 84, 7, 30, 2, 2,
	84, 85, 7, 61, 2, 2, 85, 86, 7, 30, 2, 2, 86, 91, 5, 10, 6, 2, 87, 88,
	7, 36, 2, 2, 88, 90, 5, 10, 6, 2, 89, 87, 3, 2, 2, 2, 90, 93, 3, 2, 2,
	2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94, 3, 2, 2, 2, 93, 91,
	3, 2, 2, 2, 94, 95, 7, 31, 2, 2, 95, 96, 7, 31, 2, 2, 96, 9, 3, 2, 2, 2,
	97, 98, 7, 30, 2, 2, 98, 99, 7, 56, 2, 2, 99, 100, 7, 23, 2, 2, 100, 101,
	5, 12, 7, 2, 101, 102, 7, 36, 2, 2, 102, 103, 5, 20, 11, 2, 103, 104, 7,
	31, 2, 2, 104, 11, 3, 2, 2, 2, 105, 106, 7, 81, 2, 2, 106, 13, 3, 2, 2,
	2, 107, 108, 7, 62, 2, 2, 108, 109, 7, 63, 2, 2, 109, 15, 3, 2, 2, 2, 110,
	111, 7, 62, 2, 2, 111, 112, 7, 65, 2, 2, 112, 113, 7, 63, 2, 2, 113, 17,
	3, 2, 2, 2, 114, 122, 7, 82, 2, 2, 115, 117, 7, 15, 2, 2, 116, 115, 3,
	2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 122, 7, 83, 2,
	2, 119, 122, 7, 46, 2, 2, 120, 122, 7, 47, 2, 2, 121, 114, 3, 2, 2, 2,
	121, 116, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122,
	19, 3, 2, 2, 2, 123, 124, 7, 55, 2, 2, 124, 125, 7, 30, 2, 2, 125, 126,
	7, 56, 2, 2, 126, 127, 7, 23, 2, 2, 127, 130, 5, 22, 12, 2, 128, 129, 7,
	36, 2, 2, 129, 131, 5, 26, 14, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2,
	2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 7, 31, 2, 2, 133, 21, 3, 2, 2, 2,
	134, 137, 7, 82, 2, 2, 135, 137, 5, 24, 13, 2, 136, 134, 3, 2, 2, 2, 136,
	135, 3, 2, 2, 2, 137, 23, 3, 2, 2, 2, 138, 139, 9, 2, 2, 2, 139, 25, 3,
	2, 2, 2, 140, 141, 7, 57, 2, 2, 141, 143, 7, 30, 2, 2, 142, 144, 5, 28,
	15, 2, 143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2,
	145, 146, 7, 31, 2, 2, 146, 27, 3, 2, 2, 2, 147, 152, 5, 30, 16, 2, 148,
	149, 7, 36, 2, 2, 149, 151, 5, 30, 16, 2, 150, 148, 3, 2, 2, 2, 151, 154,
	3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 29, 3, 2,
	2, 2, 154, 152, 3, 2, 2, 2, 155, 156, 7, 82, 2, 2, 156, 157, 7, 23, 2,
	2, 157, 158, 5, 18, 10, 2, 158, 31, 3, 2, 2, 2, 159, 160, 7, 81, 2, 2,
	160, 33, 3, 2, 2, 2, 17, 37, 41, 48, 54, 62, 68, 72, 79, 91, 116, 121,
	130, 136, 143, 152,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'&&'", "'||'", "'!'", "'~'", "'|'", "'&'", "'<<'", "'>>'", "'^'",
	"'%'", "':'", "'+'", "'-'", "'*'", "'/'", "'\\'", "'.'", "'.*'", "'<=>'",
	"'=='", "'='", "", "'>'", "'>='", "'<'", "'<='", "'#'", "'('", "')'", "'{'",
	"'}'", "'['", "']'", "','", "'\"'", "'''", "'`'", "'?'", "'@'", "';'",
	"'->>'", "'_'", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'",
}
var symbolicNames = []string{
	"", "AND_", "OR_", "NOT_", "TILDE_", "VERTICALBAR_", "AMPERSAND_", "SIGNEDLEFTSHIFT_",
	"SIGNEDRIGHTSHIFT_", "CARET_", "MOD_", "COLON_", "PLUS_", "MINUS_", "ASTERISK_",
	"SLASH_", "BACKSLASH_", "DOT_", "DOTASTERISK_", "SAFEEQ_", "DEQ_", "EQ_",
	"NEQ_", "GT_", "GTE_", "LT_", "LTE_", "POUND_", "LP_", "RP_", "LBE_", "RBE_",
	"LBT_", "RBT_", "COMMA_", "DQ_", "SQ_", "BQ_", "QUESTION_", "AT_", "SEMI_",
	"JSONSEPARATOR_", "UL_", "WS", "TRUE", "FALSE", "CREATE", "ALTER", "DROP",
	"SHOW", "RULE", "FROM", "MASK", "TYPE", "NAME", "PROPERTIES", "COLUMN",
	"RULES", "TABLE", "COLUMNS", "IF", "EXISTS", "COUNT", "NOT", "MD5", "KEEP_FIRST_N_LAST_M",
	"KEEP_FROM_X_TO_Y", "MASK_FIRST_N_LAST_M", "MASK_FROM_X_TO_Y", "MASK_BEFORE_SPECIAL_CHARS",
	"MASK_AFTER_SPECIAL_CHARS", "PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE",
	"MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE", "LANDLINE_NUMBER_RANDOM_REPLACE",
	"TELEPHONE_RANDOM_REPLACE", "UNIFIED_CREDIT_CODE_RANDOM_REPLACE", "GENERIC_TABLE_RANDOM_REPLACE",
	"ADDRESS_RANDOM_REPLACE", "FOR_GENERATOR", "IDENTIFIER_", "STRING_", "INT_",
	"HEX_", "NUMBER_", "HEXDIGIT_", "BITNUM_",
}

var ruleNames = []string{
	"createMaskRule", "alterMaskRule", "dropMaskRule", "maskRuleDefinition",
	"columnDefinition", "columnName", "ifExists", "ifNotExists", "literal",
	"algorithmDefinition", "algorithmTypeName", "buildInMaskAlgorithmType",
	"propertiesDefinition", "properties", "property", "ruleName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RDLStatementParser struct {
	*antlr.BaseParser
}

func NewRDLStatementParser(input antlr.TokenStream) *RDLStatementParser {
	this := new(RDLStatementParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "RDLStatement.g4"

	return this
}

// RDLStatementParser tokens.
const (
	RDLStatementParserEOF                                     = antlr.TokenEOF
	RDLStatementParserAND_                                    = 1
	RDLStatementParserOR_                                     = 2
	RDLStatementParserNOT_                                    = 3
	RDLStatementParserTILDE_                                  = 4
	RDLStatementParserVERTICALBAR_                            = 5
	RDLStatementParserAMPERSAND_                              = 6
	RDLStatementParserSIGNEDLEFTSHIFT_                        = 7
	RDLStatementParserSIGNEDRIGHTSHIFT_                       = 8
	RDLStatementParserCARET_                                  = 9
	RDLStatementParserMOD_                                    = 10
	RDLStatementParserCOLON_                                  = 11
	RDLStatementParserPLUS_                                   = 12
	RDLStatementParserMINUS_                                  = 13
	RDLStatementParserASTERISK_                               = 14
	RDLStatementParserSLASH_                                  = 15
	RDLStatementParserBACKSLASH_                              = 16
	RDLStatementParserDOT_                                    = 17
	RDLStatementParserDOTASTERISK_                            = 18
	RDLStatementParserSAFEEQ_                                 = 19
	RDLStatementParserDEQ_                                    = 20
	RDLStatementParserEQ_                                     = 21
	RDLStatementParserNEQ_                                    = 22
	RDLStatementParserGT_                                     = 23
	RDLStatementParserGTE_                                    = 24
	RDLStatementParserLT_                                     = 25
	RDLStatementParserLTE_                                    = 26
	RDLStatementParserPOUND_                                  = 27
	RDLStatementParserLP_                                     = 28
	RDLStatementParserRP_                                     = 29
	RDLStatementParserLBE_                                    = 30
	RDLStatementParserRBE_                                    = 31
	RDLStatementParserLBT_                                    = 32
	RDLStatementParserRBT_                                    = 33
	RDLStatementParserCOMMA_                                  = 34
	RDLStatementParserDQ_                                     = 35
	RDLStatementParserSQ_                                     = 36
	RDLStatementParserBQ_                                     = 37
	RDLStatementParserQUESTION_                               = 38
	RDLStatementParserAT_                                     = 39
	RDLStatementParserSEMI_                                   = 40
	RDLStatementParserJSONSEPARATOR_                          = 41
	RDLStatementParserUL_                                     = 42
	RDLStatementParserWS                                      = 43
	RDLStatementParserTRUE                                    = 44
	RDLStatementParserFALSE                                   = 45
	RDLStatementParserCREATE                                  = 46
	RDLStatementParserALTER                                   = 47
	RDLStatementParserDROP                                    = 48
	RDLStatementParserSHOW                                    = 49
	RDLStatementParserRULE                                    = 50
	RDLStatementParserFROM                                    = 51
	RDLStatementParserMASK                                    = 52
	RDLStatementParserTYPE                                    = 53
	RDLStatementParserNAME                                    = 54
	RDLStatementParserPROPERTIES                              = 55
	RDLStatementParserCOLUMN                                  = 56
	RDLStatementParserRULES                                   = 57
	RDLStatementParserTABLE                                   = 58
	RDLStatementParserCOLUMNS                                 = 59
	RDLStatementParserIF                                      = 60
	RDLStatementParserEXISTS                                  = 61
	RDLStatementParserCOUNT                                   = 62
	RDLStatementParserNOT                                     = 63
	RDLStatementParserMD5                                     = 64
	RDLStatementParserKEEP_FIRST_N_LAST_M                     = 65
	RDLStatementParserKEEP_FROM_X_TO_Y                        = 66
	RDLStatementParserMASK_FIRST_N_LAST_M                     = 67
	RDLStatementParserMASK_FROM_X_TO_Y                        = 68
	RDLStatementParserMASK_BEFORE_SPECIAL_CHARS               = 69
	RDLStatementParserMASK_AFTER_SPECIAL_CHARS                = 70
	RDLStatementParserPERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE = 71
	RDLStatementParserMILITARY_IDENTITY_NUMBER_RANDOM_REPLACE = 72
	RDLStatementParserLANDLINE_NUMBER_RANDOM_REPLACE          = 73
	RDLStatementParserTELEPHONE_RANDOM_REPLACE                = 74
	RDLStatementParserUNIFIED_CREDIT_CODE_RANDOM_REPLACE      = 75
	RDLStatementParserGENERIC_TABLE_RANDOM_REPLACE            = 76
	RDLStatementParserADDRESS_RANDOM_REPLACE                  = 77
	RDLStatementParserFOR_GENERATOR                           = 78
	RDLStatementParserIDENTIFIER_                             = 79
	RDLStatementParserSTRING_                                 = 80
	RDLStatementParserINT_                                    = 81
	RDLStatementParserHEX_                                    = 82
	RDLStatementParserNUMBER_                                 = 83
	RDLStatementParserHEXDIGIT_                               = 84
	RDLStatementParserBITNUM_                                 = 85
)

// RDLStatementParser rules.
const (
	RDLStatementParserRULE_createMaskRule           = 0
	RDLStatementParserRULE_alterMaskRule            = 1
	RDLStatementParserRULE_dropMaskRule             = 2
	RDLStatementParserRULE_maskRuleDefinition       = 3
	RDLStatementParserRULE_columnDefinition         = 4
	RDLStatementParserRULE_columnName               = 5
	RDLStatementParserRULE_ifExists                 = 6
	RDLStatementParserRULE_ifNotExists              = 7
	RDLStatementParserRULE_literal                  = 8
	RDLStatementParserRULE_algorithmDefinition      = 9
	RDLStatementParserRULE_algorithmTypeName        = 10
	RDLStatementParserRULE_buildInMaskAlgorithmType = 11
	RDLStatementParserRULE_propertiesDefinition     = 12
	RDLStatementParserRULE_properties               = 13
	RDLStatementParserRULE_property                 = 14
	RDLStatementParserRULE_ruleName                 = 15
)

// ICreateMaskRuleContext is an interface to support dynamic dispatch.
type ICreateMaskRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateMaskRuleContext differentiates from other interfaces.
	IsCreateMaskRuleContext()
}

type CreateMaskRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateMaskRuleContext() *CreateMaskRuleContext {
	var p = new(CreateMaskRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createMaskRule
	return p
}

func (*CreateMaskRuleContext) IsCreateMaskRuleContext() {}

func NewCreateMaskRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateMaskRuleContext {
	var p = new(CreateMaskRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createMaskRule

	return p
}

func (s *CreateMaskRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateMaskRuleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateMaskRuleContext) MASK() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMASK, 0)
}

func (s *CreateMaskRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *CreateMaskRuleContext) AllMaskRuleDefinition() []IMaskRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMaskRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IMaskRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMaskRuleDefinitionContext)
		}
	}

	return tst
}

func (s *CreateMaskRuleContext) MaskRuleDefinition(i int) IMaskRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaskRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMaskRuleDefinitionContext)
}

func (s *CreateMaskRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *CreateMaskRuleContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateMaskRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *CreateMaskRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *CreateMaskRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateMaskRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateMaskRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateMaskRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateMaskRule() (localctx ICreateMaskRuleContext) {
	localctx = NewCreateMaskRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RDLStatementParserRULE_createMaskRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(33)
		p.Match(RDLStatementParserMASK)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserTABLE {
		{
			p.SetState(34)
			p.Match(RDLStatementParserTABLE)
		}

	}
	{
		p.SetState(37)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(38)
			p.IfNotExists()
		}

	}
	{
		p.SetState(41)
		p.MaskRuleDefinition()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(42)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(43)
			p.MaskRuleDefinition()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlterMaskRuleContext is an interface to support dynamic dispatch.
type IAlterMaskRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterMaskRuleContext differentiates from other interfaces.
	IsAlterMaskRuleContext()
}

type AlterMaskRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterMaskRuleContext() *AlterMaskRuleContext {
	var p = new(AlterMaskRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterMaskRule
	return p
}

func (*AlterMaskRuleContext) IsAlterMaskRuleContext() {}

func NewAlterMaskRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterMaskRuleContext {
	var p = new(AlterMaskRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterMaskRule

	return p
}

func (s *AlterMaskRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterMaskRuleContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterMaskRuleContext) MASK() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMASK, 0)
}

func (s *AlterMaskRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *AlterMaskRuleContext) AllMaskRuleDefinition() []IMaskRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMaskRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IMaskRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMaskRuleDefinitionContext)
		}
	}

	return tst
}

func (s *AlterMaskRuleContext) MaskRuleDefinition(i int) IMaskRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaskRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMaskRuleDefinitionContext)
}

func (s *AlterMaskRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *AlterMaskRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *AlterMaskRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *AlterMaskRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterMaskRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterMaskRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterMaskRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterMaskRule() (localctx IAlterMaskRuleContext) {
	localctx = NewAlterMaskRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RDLStatementParserRULE_alterMaskRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(50)
		p.Match(RDLStatementParserMASK)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserTABLE {
		{
			p.SetState(51)
			p.Match(RDLStatementParserTABLE)
		}

	}
	{
		p.SetState(54)
		p.Match(RDLStatementParserRULE)
	}
	{
		p.SetState(55)
		p.MaskRuleDefinition()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(56)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(57)
			p.MaskRuleDefinition()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropMaskRuleContext is an interface to support dynamic dispatch.
type IDropMaskRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropMaskRuleContext differentiates from other interfaces.
	IsDropMaskRuleContext()
}

type DropMaskRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropMaskRuleContext() *DropMaskRuleContext {
	var p = new(DropMaskRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropMaskRule
	return p
}

func (*DropMaskRuleContext) IsDropMaskRuleContext() {}

func NewDropMaskRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropMaskRuleContext {
	var p = new(DropMaskRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropMaskRule

	return p
}

func (s *DropMaskRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropMaskRuleContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropMaskRuleContext) MASK() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMASK, 0)
}

func (s *DropMaskRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *DropMaskRuleContext) AllRuleName() []IRuleNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleNameContext)(nil)).Elem())
	var tst = make([]IRuleNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleNameContext)
		}
	}

	return tst
}

func (s *DropMaskRuleContext) RuleName(i int) IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *DropMaskRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *DropMaskRuleContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropMaskRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropMaskRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropMaskRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropMaskRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropMaskRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropMaskRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropMaskRule() (localctx IDropMaskRuleContext) {
	localctx = NewDropMaskRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RDLStatementParserRULE_dropMaskRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(64)
		p.Match(RDLStatementParserMASK)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserTABLE {
		{
			p.SetState(65)
			p.Match(RDLStatementParserTABLE)
		}

	}
	{
		p.SetState(68)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(69)
			p.IfExists()
		}

	}
	{
		p.SetState(72)
		p.RuleName()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(73)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(74)
			p.RuleName()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMaskRuleDefinitionContext is an interface to support dynamic dispatch.
type IMaskRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaskRuleDefinitionContext differentiates from other interfaces.
	IsMaskRuleDefinitionContext()
}

type MaskRuleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaskRuleDefinitionContext() *MaskRuleDefinitionContext {
	var p = new(MaskRuleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_maskRuleDefinition
	return p
}

func (*MaskRuleDefinitionContext) IsMaskRuleDefinitionContext() {}

func NewMaskRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaskRuleDefinitionContext {
	var p = new(MaskRuleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_maskRuleDefinition

	return p
}

func (s *MaskRuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MaskRuleDefinitionContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *MaskRuleDefinitionContext) AllLP_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserLP_)
}

func (s *MaskRuleDefinitionContext) LP_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, i)
}

func (s *MaskRuleDefinitionContext) COLUMNS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOLUMNS, 0)
}

func (s *MaskRuleDefinitionContext) AllColumnDefinition() []IColumnDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnDefinitionContext)(nil)).Elem())
	var tst = make([]IColumnDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnDefinitionContext)
		}
	}

	return tst
}

func (s *MaskRuleDefinitionContext) ColumnDefinition(i int) IColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *MaskRuleDefinitionContext) AllRP_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserRP_)
}

func (s *MaskRuleDefinitionContext) RP_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, i)
}

func (s *MaskRuleDefinitionContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *MaskRuleDefinitionContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *MaskRuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaskRuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaskRuleDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitMaskRuleDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) MaskRuleDefinition() (localctx IMaskRuleDefinitionContext) {
	localctx = NewMaskRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RDLStatementParserRULE_maskRuleDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.RuleName()
	}
	{
		p.SetState(81)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(82)
		p.Match(RDLStatementParserCOLUMNS)
	}
	{
		p.SetState(83)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(84)
		p.ColumnDefinition()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(85)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(86)
			p.ColumnDefinition()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(RDLStatementParserRP_)
	}
	{
		p.SetState(93)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IColumnDefinitionContext is an interface to support dynamic dispatch.
type IColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnDefinitionContext differentiates from other interfaces.
	IsColumnDefinitionContext()
}

type ColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionContext() *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_columnDefinition
	return p
}

func (*ColumnDefinitionContext) IsColumnDefinitionContext() {}

func NewColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_columnDefinition

	return p
}

func (s *ColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ColumnDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserNAME, 0)
}

func (s *ColumnDefinitionContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *ColumnDefinitionContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *ColumnDefinitionContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *ColumnDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ColumnDefinition() (localctx IColumnDefinitionContext) {
	localctx = NewColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RDLStatementParserRULE_columnDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(96)
		p.Match(RDLStatementParserNAME)
	}
	{
		p.SetState(97)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(98)
		p.ColumnName()
	}
	{
		p.SetState(99)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(100)
		p.AlgorithmDefinition()
	}
	{
		p.SetState(101)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RDLStatementParserRULE_columnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IIfExistsContext is an interface to support dynamic dispatch.
type IIfExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfExistsContext differentiates from other interfaces.
	IsIfExistsContext()
}

type IfExistsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExistsContext() *IfExistsContext {
	var p = new(IfExistsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_ifExists
	return p
}

func (*IfExistsContext) IsIfExistsContext() {}

func NewIfExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExistsContext {
	var p = new(IfExistsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_ifExists

	return p
}

func (s *IfExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExistsContext) IF() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIF, 0)
}

func (s *IfExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEXISTS, 0)
}

func (s *IfExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExistsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitIfExists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) IfExists() (localctx IIfExistsContext) {
	localctx = NewIfExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RDLStatementParserRULE_ifExists)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(106)
		p.Match(RDLStatementParserEXISTS)
	}

	return localctx
}

// IIfNotExistsContext is an interface to support dynamic dispatch.
type IIfNotExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfNotExistsContext differentiates from other interfaces.
	IsIfNotExistsContext()
}

type IfNotExistsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistsContext() *IfNotExistsContext {
	var p = new(IfNotExistsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_ifNotExists
	return p
}

func (*IfNotExistsContext) IsIfNotExistsContext() {}

func NewIfNotExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistsContext {
	var p = new(IfNotExistsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_ifNotExists

	return p
}

func (s *IfNotExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistsContext) IF() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIF, 0)
}

func (s *IfNotExistsContext) NOT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserNOT, 0)
}

func (s *IfNotExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEXISTS, 0)
}

func (s *IfNotExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitIfNotExists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) IfNotExists() (localctx IIfNotExistsContext) {
	localctx = NewIfNotExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RDLStatementParserRULE_ifNotExists)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(109)
		p.Match(RDLStatementParserNOT)
	}
	{
		p.SetState(110)
		p.Match(RDLStatementParserEXISTS)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *LiteralContext) INT_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserINT_, 0)
}

func (s *LiteralContext) MINUS_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMINUS_, 0)
}

func (s *LiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTRUE, 0)
}

func (s *LiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserFALSE, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RDLStatementParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserMINUS_, RDLStatementParserINT_:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RDLStatementParserMINUS_ {
			{
				p.SetState(113)
				p.Match(RDLStatementParserMINUS_)
			}

		}
		{
			p.SetState(116)
			p.Match(RDLStatementParserINT_)
		}

	case RDLStatementParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(117)
			p.Match(RDLStatementParserTRUE)
		}

	case RDLStatementParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(118)
			p.Match(RDLStatementParserFALSE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAlgorithmDefinitionContext is an interface to support dynamic dispatch.
type IAlgorithmDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlgorithmDefinitionContext differentiates from other interfaces.
	IsAlgorithmDefinitionContext()
}

type AlgorithmDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmDefinitionContext() *AlgorithmDefinitionContext {
	var p = new(AlgorithmDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_algorithmDefinition
	return p
}

func (*AlgorithmDefinitionContext) IsAlgorithmDefinitionContext() {}

func NewAlgorithmDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmDefinitionContext {
	var p = new(AlgorithmDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_algorithmDefinition

	return p
}

func (s *AlgorithmDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmDefinitionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTYPE, 0)
}

func (s *AlgorithmDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *AlgorithmDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserNAME, 0)
}

func (s *AlgorithmDefinitionContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *AlgorithmDefinitionContext) AlgorithmTypeName() IAlgorithmTypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmTypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmTypeNameContext)
}

func (s *AlgorithmDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *AlgorithmDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *AlgorithmDefinitionContext) PropertiesDefinition() IPropertiesDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertiesDefinitionContext)
}

func (s *AlgorithmDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlgorithmDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlgorithmDefinition() (localctx IAlgorithmDefinitionContext) {
	localctx = NewAlgorithmDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RDLStatementParserRULE_algorithmDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(RDLStatementParserTYPE)
	}
	{
		p.SetState(122)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(123)
		p.Match(RDLStatementParserNAME)
	}
	{
		p.SetState(124)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(125)
		p.AlgorithmTypeName()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(126)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(127)
			p.PropertiesDefinition()
		}

	}
	{
		p.SetState(130)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IAlgorithmTypeNameContext is an interface to support dynamic dispatch.
type IAlgorithmTypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlgorithmTypeNameContext differentiates from other interfaces.
	IsAlgorithmTypeNameContext()
}

type AlgorithmTypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmTypeNameContext() *AlgorithmTypeNameContext {
	var p = new(AlgorithmTypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_algorithmTypeName
	return p
}

func (*AlgorithmTypeNameContext) IsAlgorithmTypeNameContext() {}

func NewAlgorithmTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmTypeNameContext {
	var p = new(AlgorithmTypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_algorithmTypeName

	return p
}

func (s *AlgorithmTypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmTypeNameContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *AlgorithmTypeNameContext) BuildInMaskAlgorithmType() IBuildInMaskAlgorithmTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildInMaskAlgorithmTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildInMaskAlgorithmTypeContext)
}

func (s *AlgorithmTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmTypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmTypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlgorithmTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlgorithmTypeName() (localctx IAlgorithmTypeNameContext) {
	localctx = NewAlgorithmTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RDLStatementParserRULE_algorithmTypeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserMD5, RDLStatementParserKEEP_FIRST_N_LAST_M, RDLStatementParserKEEP_FROM_X_TO_Y, RDLStatementParserMASK_FIRST_N_LAST_M, RDLStatementParserMASK_FROM_X_TO_Y, RDLStatementParserMASK_BEFORE_SPECIAL_CHARS, RDLStatementParserMASK_AFTER_SPECIAL_CHARS, RDLStatementParserPERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE, RDLStatementParserMILITARY_IDENTITY_NUMBER_RANDOM_REPLACE, RDLStatementParserLANDLINE_NUMBER_RANDOM_REPLACE, RDLStatementParserTELEPHONE_RANDOM_REPLACE, RDLStatementParserUNIFIED_CREDIT_CODE_RANDOM_REPLACE, RDLStatementParserGENERIC_TABLE_RANDOM_REPLACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.BuildInMaskAlgorithmType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuildInMaskAlgorithmTypeContext is an interface to support dynamic dispatch.
type IBuildInMaskAlgorithmTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildInMaskAlgorithmTypeContext differentiates from other interfaces.
	IsBuildInMaskAlgorithmTypeContext()
}

type BuildInMaskAlgorithmTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildInMaskAlgorithmTypeContext() *BuildInMaskAlgorithmTypeContext {
	var p = new(BuildInMaskAlgorithmTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildInMaskAlgorithmType
	return p
}

func (*BuildInMaskAlgorithmTypeContext) IsBuildInMaskAlgorithmTypeContext() {}

func NewBuildInMaskAlgorithmTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildInMaskAlgorithmTypeContext {
	var p = new(BuildInMaskAlgorithmTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildInMaskAlgorithmType

	return p
}

func (s *BuildInMaskAlgorithmTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildInMaskAlgorithmTypeContext) MD5() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMD5, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) KEEP_FIRST_N_LAST_M() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserKEEP_FIRST_N_LAST_M, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) KEEP_FROM_X_TO_Y() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserKEEP_FROM_X_TO_Y, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) MASK_FIRST_N_LAST_M() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMASK_FIRST_N_LAST_M, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) MASK_FROM_X_TO_Y() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMASK_FROM_X_TO_Y, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) MASK_BEFORE_SPECIAL_CHARS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMASK_BEFORE_SPECIAL_CHARS, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) MASK_AFTER_SPECIAL_CHARS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMASK_AFTER_SPECIAL_CHARS, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserPERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMILITARY_IDENTITY_NUMBER_RANDOM_REPLACE, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) LANDLINE_NUMBER_RANDOM_REPLACE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLANDLINE_NUMBER_RANDOM_REPLACE, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) TELEPHONE_RANDOM_REPLACE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTELEPHONE_RANDOM_REPLACE, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) UNIFIED_CREDIT_CODE_RANDOM_REPLACE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserUNIFIED_CREDIT_CODE_RANDOM_REPLACE, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) GENERIC_TABLE_RANDOM_REPLACE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserGENERIC_TABLE_RANDOM_REPLACE, 0)
}

func (s *BuildInMaskAlgorithmTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInMaskAlgorithmTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildInMaskAlgorithmTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildInMaskAlgorithmType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildInMaskAlgorithmType() (localctx IBuildInMaskAlgorithmTypeContext) {
	localctx = NewBuildInMaskAlgorithmTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RDLStatementParserRULE_buildInMaskAlgorithmType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(RDLStatementParserMD5-64))|(1<<(RDLStatementParserKEEP_FIRST_N_LAST_M-64))|(1<<(RDLStatementParserKEEP_FROM_X_TO_Y-64))|(1<<(RDLStatementParserMASK_FIRST_N_LAST_M-64))|(1<<(RDLStatementParserMASK_FROM_X_TO_Y-64))|(1<<(RDLStatementParserMASK_BEFORE_SPECIAL_CHARS-64))|(1<<(RDLStatementParserMASK_AFTER_SPECIAL_CHARS-64))|(1<<(RDLStatementParserPERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE-64))|(1<<(RDLStatementParserMILITARY_IDENTITY_NUMBER_RANDOM_REPLACE-64))|(1<<(RDLStatementParserLANDLINE_NUMBER_RANDOM_REPLACE-64))|(1<<(RDLStatementParserTELEPHONE_RANDOM_REPLACE-64))|(1<<(RDLStatementParserUNIFIED_CREDIT_CODE_RANDOM_REPLACE-64))|(1<<(RDLStatementParserGENERIC_TABLE_RANDOM_REPLACE-64)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPropertiesDefinitionContext is an interface to support dynamic dispatch.
type IPropertiesDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertiesDefinitionContext differentiates from other interfaces.
	IsPropertiesDefinitionContext()
}

type PropertiesDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertiesDefinitionContext() *PropertiesDefinitionContext {
	var p = new(PropertiesDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_propertiesDefinition
	return p
}

func (*PropertiesDefinitionContext) IsPropertiesDefinitionContext() {}

func NewPropertiesDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertiesDefinitionContext {
	var p = new(PropertiesDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_propertiesDefinition

	return p
}

func (s *PropertiesDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertiesDefinitionContext) PROPERTIES() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserPROPERTIES, 0)
}

func (s *PropertiesDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *PropertiesDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *PropertiesDefinitionContext) Properties() IPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertiesContext)
}

func (s *PropertiesDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertiesDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertiesDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitPropertiesDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) PropertiesDefinition() (localctx IPropertiesDefinitionContext) {
	localctx = NewPropertiesDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RDLStatementParserRULE_propertiesDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(RDLStatementParserPROPERTIES)
	}
	{
		p.SetState(139)
		p.Match(RDLStatementParserLP_)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserSTRING_ {
		{
			p.SetState(140)
			p.Properties()
		}

	}
	{
		p.SetState(143)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IPropertiesContext is an interface to support dynamic dispatch.
type IPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertiesContext differentiates from other interfaces.
	IsPropertiesContext()
}

type PropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertiesContext() *PropertiesContext {
	var p = new(PropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_properties
	return p
}

func (*PropertiesContext) IsPropertiesContext() {}

func NewPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertiesContext {
	var p = new(PropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_properties

	return p
}

func (s *PropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertiesContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *PropertiesContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PropertiesContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *PropertiesContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *PropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertiesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitProperties(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) Properties() (localctx IPropertiesContext) {
	localctx = NewPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RDLStatementParserRULE_properties)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Property()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(146)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(147)
			p.Property()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() ILiteralContext

	// SetValue sets the value rule contexts.
	SetValue(ILiteralContext)

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  ILiteralContext
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) GetKey() antlr.Token { return s.key }

func (s *PropertyContext) SetKey(v antlr.Token) { s.key = v }

func (s *PropertyContext) GetValue() ILiteralContext { return s.value }

func (s *PropertyContext) SetValue(v ILiteralContext) { s.value = v }

func (s *PropertyContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *PropertyContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *PropertyContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RDLStatementParserRULE_property)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)

		var _m = p.Match(RDLStatementParserSTRING_)

		localctx.(*PropertyContext).key = _m
	}
	{
		p.SetState(154)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(155)

		var _x = p.Literal()

		localctx.(*PropertyContext).value = _x
	}

	return localctx
}

// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_ruleName
	return p
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitRuleName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) RuleName() (localctx IRuleNameContext) {
	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RDLStatementParserRULE_ruleName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}
