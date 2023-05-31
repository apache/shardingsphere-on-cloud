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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 81, 192,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 55, 10, 2, 3,
	2, 3, 2, 3, 2, 7, 2, 60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 71, 10, 3, 12, 3, 14, 3, 74, 11, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 80, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 85, 10, 4, 12,
	4, 14, 4, 88, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 95, 10, 5, 3,
	5, 3, 5, 5, 5, 99, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 127, 10, 12, 12,
	12, 14, 12, 130, 11, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 5, 16, 143, 10, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 148, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17,
	157, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 5, 18, 163, 10, 18, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 5, 20, 170, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 7, 21, 177, 10, 21, 12, 21, 14, 21, 180, 11, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 2, 2, 26, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 2, 3, 3, 2, 70, 72, 2, 183, 2, 50, 3, 2, 2, 2, 4, 64, 3,
	2, 2, 2, 6, 75, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2, 10, 102, 3, 2, 2, 2, 12,
	106, 3, 2, 2, 2, 14, 108, 3, 2, 2, 2, 16, 112, 3, 2, 2, 2, 18, 117, 3,
	2, 2, 2, 20, 121, 3, 2, 2, 2, 22, 123, 3, 2, 2, 2, 24, 131, 3, 2, 2, 2,
	26, 133, 3, 2, 2, 2, 28, 136, 3, 2, 2, 2, 30, 147, 3, 2, 2, 2, 32, 149,
	3, 2, 2, 2, 34, 162, 3, 2, 2, 2, 36, 164, 3, 2, 2, 2, 38, 166, 3, 2, 2,
	2, 40, 173, 3, 2, 2, 2, 42, 181, 3, 2, 2, 2, 44, 185, 3, 2, 2, 2, 46, 187,
	3, 2, 2, 2, 48, 189, 3, 2, 2, 2, 50, 51, 7, 48, 2, 2, 51, 52, 7, 54, 2,
	2, 52, 54, 7, 52, 2, 2, 53, 55, 5, 28, 15, 2, 54, 53, 3, 2, 2, 2, 54, 55,
	3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 61, 5, 8, 5, 2, 57, 58, 7, 36, 2, 2,
	58, 60, 5, 8, 5, 2, 59, 57, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3,
	2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 3, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64,
	65, 7, 49, 2, 2, 65, 66, 7, 54, 2, 2, 66, 67, 7, 52, 2, 2, 67, 72, 5, 8,
	5, 2, 68, 69, 7, 36, 2, 2, 69, 71, 5, 8, 5, 2, 70, 68, 3, 2, 2, 2, 71,
	74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 5, 3, 2, 2,
	2, 74, 72, 3, 2, 2, 2, 75, 76, 7, 50, 2, 2, 76, 77, 7, 54, 2, 2, 77, 79,
	7, 52, 2, 2, 78, 80, 5, 26, 14, 2, 79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2,
	2, 80, 81, 3, 2, 2, 2, 81, 86, 5, 12, 7, 2, 82, 83, 7, 36, 2, 2, 83, 85,
	5, 12, 7, 2, 84, 82, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2,
	86, 87, 3, 2, 2, 2, 87, 7, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90, 5, 12,
	7, 2, 90, 91, 7, 30, 2, 2, 91, 94, 5, 10, 6, 2, 92, 93, 7, 36, 2, 2, 93,
	95, 5, 18, 10, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 98, 3, 2,
	2, 2, 96, 97, 7, 36, 2, 2, 97, 99, 5, 32, 17, 2, 98, 96, 3, 2, 2, 2, 98,
	99, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 7, 31, 2, 2, 101, 9, 3,
	2, 2, 2, 102, 103, 5, 14, 8, 2, 103, 104, 7, 36, 2, 2, 104, 105, 5, 16,
	9, 2, 105, 11, 3, 2, 2, 2, 106, 107, 7, 75, 2, 2, 107, 13, 3, 2, 2, 2,
	108, 109, 7, 55, 2, 2, 109, 110, 7, 23, 2, 2, 110, 111, 5, 20, 11, 2, 111,
	15, 3, 2, 2, 2, 112, 113, 7, 56, 2, 2, 113, 114, 7, 30, 2, 2, 114, 115,
	5, 22, 12, 2, 115, 116, 7, 31, 2, 2, 116, 17, 3, 2, 2, 2, 117, 118, 7,
	57, 2, 2, 118, 119, 7, 23, 2, 2, 119, 120, 5, 24, 13, 2, 120, 19, 3, 2,
	2, 2, 121, 122, 5, 48, 25, 2, 122, 21, 3, 2, 2, 2, 123, 128, 5, 48, 25,
	2, 124, 125, 7, 36, 2, 2, 125, 127, 5, 48, 25, 2, 126, 124, 3, 2, 2, 2,
	127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	23, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 76, 2, 2, 132, 25, 3,
	2, 2, 2, 133, 134, 7, 67, 2, 2, 134, 135, 7, 68, 2, 2, 135, 27, 3, 2, 2,
	2, 136, 137, 7, 67, 2, 2, 137, 138, 7, 73, 2, 2, 138, 139, 7, 68, 2, 2,
	139, 29, 3, 2, 2, 2, 140, 148, 7, 76, 2, 2, 141, 143, 7, 15, 2, 2, 142,
	141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 148,
	7, 77, 2, 2, 145, 148, 7, 46, 2, 2, 146, 148, 7, 47, 2, 2, 147, 140, 3,
	2, 2, 2, 147, 142, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 146, 3, 2, 2,
	2, 148, 31, 3, 2, 2, 2, 149, 150, 7, 58, 2, 2, 150, 151, 7, 30, 2, 2, 151,
	152, 7, 59, 2, 2, 152, 153, 7, 23, 2, 2, 153, 156, 5, 34, 18, 2, 154, 155,
	7, 36, 2, 2, 155, 157, 5, 38, 20, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3,
	2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 7, 31, 2, 2, 159, 33, 3, 2, 2,
	2, 160, 163, 7, 76, 2, 2, 161, 163, 5, 36, 19, 2, 162, 160, 3, 2, 2, 2,
	162, 161, 3, 2, 2, 2, 163, 35, 3, 2, 2, 2, 164, 165, 9, 2, 2, 2, 165, 37,
	3, 2, 2, 2, 166, 167, 7, 60, 2, 2, 167, 169, 7, 30, 2, 2, 168, 170, 5,
	40, 21, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 3, 2,
	2, 2, 171, 172, 7, 31, 2, 2, 172, 39, 3, 2, 2, 2, 173, 178, 5, 42, 22,
	2, 174, 175, 7, 36, 2, 2, 175, 177, 5, 42, 22, 2, 176, 174, 3, 2, 2, 2,
	177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179,
	41, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181, 182, 7, 76, 2, 2, 182, 183,
	7, 23, 2, 2, 183, 184, 5, 30, 16, 2, 184, 43, 3, 2, 2, 2, 185, 186, 7,
	75, 2, 2, 186, 45, 3, 2, 2, 2, 187, 188, 7, 75, 2, 2, 188, 47, 3, 2, 2,
	2, 189, 190, 7, 75, 2, 2, 190, 49, 3, 2, 2, 2, 16, 54, 61, 72, 79, 86,
	94, 98, 128, 142, 147, 156, 162, 169, 178,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'&&'", "'||'", "'!'", "'~'", "'|'", "'&'", "'<<'", "'>>'", "'^'",
	"'%'", "':'", "'+'", "'-'", "'*'", "'/'", "'\\'", "'.'", "'.*'", "'<=>'",
	"'=='", "'='", "", "'>'", "'>='", "'<'", "'<='", "'#'", "'('", "')'", "'{'",
	"'}'", "'['", "']'", "','", "'\"'", "'''", "'`'", "'?'", "'@'", "';'",
	"'->>'", "'_'", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'",
}
var symbolicNames = []string{
	"", "AND_", "OR_", "NOT_", "TILDE_", "VERTICALBAR_", "AMPERSAND_", "SIGNEDLEFTSHIFT_",
	"SIGNEDRIGHTSHIFT_", "CARET_", "MOD_", "COLON_", "PLUS_", "MINUS_", "ASTERISK_",
	"SLASH_", "BACKSLASH_", "DOT_", "DOTASTERISK_", "SAFEEQ_", "DEQ_", "EQ_",
	"NEQ_", "GT_", "GTE_", "LT_", "LTE_", "POUND_", "LP_", "RP_", "LBE_", "RBE_",
	"LBT_", "RBT_", "COMMA_", "DQ_", "SQ_", "BQ_", "QUESTION_", "AT_", "SEMI_",
	"JSONSEPARATOR_", "UL_", "WS", "TRUE", "FALSE", "CREATE", "ALTER", "DROP",
	"SHOW", "RULE", "FROM", "READWRITE_SPLITTING", "WRITE_STORAGE_UNIT", "READ_STORAGE_UNITS",
	"TRANSACTIONAL_READ_QUERY_STRATEGY", "TYPE", "NAME", "PROPERTIES", "RULES",
	"RESOURCES", "STATUS", "ENABLE", "DISABLE", "READ", "IF", "EXISTS", "COUNT",
	"ROUND_ROBIN", "RANDOM", "WEIGHT", "NOT", "FOR_GENERATOR", "IDENTIFIER_",
	"STRING_", "INT_", "HEX_", "NUMBER_", "HEXDIGIT_", "BITNUM_",
}

var ruleNames = []string{
	"createReadwriteSplittingRule", "alterReadwriteSplittingRule", "dropReadwriteSplittingRule",
	"readwriteSplittingRuleDefinition", "dataSourceDefinition", "ruleName",
	"writeStorageUnit", "readStorageUnits", "transactionalReadQueryStrategy",
	"writeStorageUnitName", "readStorageUnitsNames", "transactionalReadQueryStrategyName",
	"ifExists", "ifNotExists", "literal", "algorithmDefinition", "algorithmTypeName",
	"buildInReadQueryLoadBalanceAlgorithmType", "propertiesDefinition", "properties",
	"property", "databaseName", "groupName", "storageUnitName",
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
	RDLStatementParserEOF                               = antlr.TokenEOF
	RDLStatementParserAND_                              = 1
	RDLStatementParserOR_                               = 2
	RDLStatementParserNOT_                              = 3
	RDLStatementParserTILDE_                            = 4
	RDLStatementParserVERTICALBAR_                      = 5
	RDLStatementParserAMPERSAND_                        = 6
	RDLStatementParserSIGNEDLEFTSHIFT_                  = 7
	RDLStatementParserSIGNEDRIGHTSHIFT_                 = 8
	RDLStatementParserCARET_                            = 9
	RDLStatementParserMOD_                              = 10
	RDLStatementParserCOLON_                            = 11
	RDLStatementParserPLUS_                             = 12
	RDLStatementParserMINUS_                            = 13
	RDLStatementParserASTERISK_                         = 14
	RDLStatementParserSLASH_                            = 15
	RDLStatementParserBACKSLASH_                        = 16
	RDLStatementParserDOT_                              = 17
	RDLStatementParserDOTASTERISK_                      = 18
	RDLStatementParserSAFEEQ_                           = 19
	RDLStatementParserDEQ_                              = 20
	RDLStatementParserEQ_                               = 21
	RDLStatementParserNEQ_                              = 22
	RDLStatementParserGT_                               = 23
	RDLStatementParserGTE_                              = 24
	RDLStatementParserLT_                               = 25
	RDLStatementParserLTE_                              = 26
	RDLStatementParserPOUND_                            = 27
	RDLStatementParserLP_                               = 28
	RDLStatementParserRP_                               = 29
	RDLStatementParserLBE_                              = 30
	RDLStatementParserRBE_                              = 31
	RDLStatementParserLBT_                              = 32
	RDLStatementParserRBT_                              = 33
	RDLStatementParserCOMMA_                            = 34
	RDLStatementParserDQ_                               = 35
	RDLStatementParserSQ_                               = 36
	RDLStatementParserBQ_                               = 37
	RDLStatementParserQUESTION_                         = 38
	RDLStatementParserAT_                               = 39
	RDLStatementParserSEMI_                             = 40
	RDLStatementParserJSONSEPARATOR_                    = 41
	RDLStatementParserUL_                               = 42
	RDLStatementParserWS                                = 43
	RDLStatementParserTRUE                              = 44
	RDLStatementParserFALSE                             = 45
	RDLStatementParserCREATE                            = 46
	RDLStatementParserALTER                             = 47
	RDLStatementParserDROP                              = 48
	RDLStatementParserSHOW                              = 49
	RDLStatementParserRULE                              = 50
	RDLStatementParserFROM                              = 51
	RDLStatementParserREADWRITE_SPLITTING               = 52
	RDLStatementParserWRITE_STORAGE_UNIT                = 53
	RDLStatementParserREAD_STORAGE_UNITS                = 54
	RDLStatementParserTRANSACTIONAL_READ_QUERY_STRATEGY = 55
	RDLStatementParserTYPE                              = 56
	RDLStatementParserNAME                              = 57
	RDLStatementParserPROPERTIES                        = 58
	RDLStatementParserRULES                             = 59
	RDLStatementParserRESOURCES                         = 60
	RDLStatementParserSTATUS                            = 61
	RDLStatementParserENABLE                            = 62
	RDLStatementParserDISABLE                           = 63
	RDLStatementParserREAD                              = 64
	RDLStatementParserIF                                = 65
	RDLStatementParserEXISTS                            = 66
	RDLStatementParserCOUNT                             = 67
	RDLStatementParserROUND_ROBIN                       = 68
	RDLStatementParserRANDOM                            = 69
	RDLStatementParserWEIGHT                            = 70
	RDLStatementParserNOT                               = 71
	RDLStatementParserFOR_GENERATOR                     = 72
	RDLStatementParserIDENTIFIER_                       = 73
	RDLStatementParserSTRING_                           = 74
	RDLStatementParserINT_                              = 75
	RDLStatementParserHEX_                              = 76
	RDLStatementParserNUMBER_                           = 77
	RDLStatementParserHEXDIGIT_                         = 78
	RDLStatementParserBITNUM_                           = 79
)

// RDLStatementParser rules.
const (
	RDLStatementParserRULE_createReadwriteSplittingRule             = 0
	RDLStatementParserRULE_alterReadwriteSplittingRule              = 1
	RDLStatementParserRULE_dropReadwriteSplittingRule               = 2
	RDLStatementParserRULE_readwriteSplittingRuleDefinition         = 3
	RDLStatementParserRULE_dataSourceDefinition                     = 4
	RDLStatementParserRULE_ruleName                                 = 5
	RDLStatementParserRULE_writeStorageUnit                         = 6
	RDLStatementParserRULE_readStorageUnits                         = 7
	RDLStatementParserRULE_transactionalReadQueryStrategy           = 8
	RDLStatementParserRULE_writeStorageUnitName                     = 9
	RDLStatementParserRULE_readStorageUnitsNames                    = 10
	RDLStatementParserRULE_transactionalReadQueryStrategyName       = 11
	RDLStatementParserRULE_ifExists                                 = 12
	RDLStatementParserRULE_ifNotExists                              = 13
	RDLStatementParserRULE_literal                                  = 14
	RDLStatementParserRULE_algorithmDefinition                      = 15
	RDLStatementParserRULE_algorithmTypeName                        = 16
	RDLStatementParserRULE_buildInReadQueryLoadBalanceAlgorithmType = 17
	RDLStatementParserRULE_propertiesDefinition                     = 18
	RDLStatementParserRULE_properties                               = 19
	RDLStatementParserRULE_property                                 = 20
	RDLStatementParserRULE_databaseName                             = 21
	RDLStatementParserRULE_groupName                                = 22
	RDLStatementParserRULE_storageUnitName                          = 23
)

// ICreateReadwriteSplittingRuleContext is an interface to support dynamic dispatch.
type ICreateReadwriteSplittingRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateReadwriteSplittingRuleContext differentiates from other interfaces.
	IsCreateReadwriteSplittingRuleContext()
}

type CreateReadwriteSplittingRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateReadwriteSplittingRuleContext() *CreateReadwriteSplittingRuleContext {
	var p = new(CreateReadwriteSplittingRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createReadwriteSplittingRule
	return p
}

func (*CreateReadwriteSplittingRuleContext) IsCreateReadwriteSplittingRuleContext() {}

func NewCreateReadwriteSplittingRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateReadwriteSplittingRuleContext {
	var p = new(CreateReadwriteSplittingRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createReadwriteSplittingRule

	return p
}

func (s *CreateReadwriteSplittingRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateReadwriteSplittingRuleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateReadwriteSplittingRuleContext) READWRITE_SPLITTING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREADWRITE_SPLITTING, 0)
}

func (s *CreateReadwriteSplittingRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *CreateReadwriteSplittingRuleContext) AllReadwriteSplittingRuleDefinition() []IReadwriteSplittingRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReadwriteSplittingRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IReadwriteSplittingRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReadwriteSplittingRuleDefinitionContext)
		}
	}

	return tst
}

func (s *CreateReadwriteSplittingRuleContext) ReadwriteSplittingRuleDefinition(i int) IReadwriteSplittingRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadwriteSplittingRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReadwriteSplittingRuleDefinitionContext)
}

func (s *CreateReadwriteSplittingRuleContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateReadwriteSplittingRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *CreateReadwriteSplittingRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *CreateReadwriteSplittingRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateReadwriteSplittingRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateReadwriteSplittingRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateReadwriteSplittingRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateReadwriteSplittingRule() (localctx ICreateReadwriteSplittingRuleContext) {
	localctx = NewCreateReadwriteSplittingRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RDLStatementParserRULE_createReadwriteSplittingRule)
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
		p.SetState(48)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(49)
		p.Match(RDLStatementParserREADWRITE_SPLITTING)
	}
	{
		p.SetState(50)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(51)
			p.IfNotExists()
		}

	}
	{
		p.SetState(54)
		p.ReadwriteSplittingRuleDefinition()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(55)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(56)
			p.ReadwriteSplittingRuleDefinition()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlterReadwriteSplittingRuleContext is an interface to support dynamic dispatch.
type IAlterReadwriteSplittingRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterReadwriteSplittingRuleContext differentiates from other interfaces.
	IsAlterReadwriteSplittingRuleContext()
}

type AlterReadwriteSplittingRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterReadwriteSplittingRuleContext() *AlterReadwriteSplittingRuleContext {
	var p = new(AlterReadwriteSplittingRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterReadwriteSplittingRule
	return p
}

func (*AlterReadwriteSplittingRuleContext) IsAlterReadwriteSplittingRuleContext() {}

func NewAlterReadwriteSplittingRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterReadwriteSplittingRuleContext {
	var p = new(AlterReadwriteSplittingRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterReadwriteSplittingRule

	return p
}

func (s *AlterReadwriteSplittingRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterReadwriteSplittingRuleContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterReadwriteSplittingRuleContext) READWRITE_SPLITTING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREADWRITE_SPLITTING, 0)
}

func (s *AlterReadwriteSplittingRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *AlterReadwriteSplittingRuleContext) AllReadwriteSplittingRuleDefinition() []IReadwriteSplittingRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReadwriteSplittingRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IReadwriteSplittingRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReadwriteSplittingRuleDefinitionContext)
		}
	}

	return tst
}

func (s *AlterReadwriteSplittingRuleContext) ReadwriteSplittingRuleDefinition(i int) IReadwriteSplittingRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadwriteSplittingRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReadwriteSplittingRuleDefinitionContext)
}

func (s *AlterReadwriteSplittingRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *AlterReadwriteSplittingRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *AlterReadwriteSplittingRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterReadwriteSplittingRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterReadwriteSplittingRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterReadwriteSplittingRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterReadwriteSplittingRule() (localctx IAlterReadwriteSplittingRuleContext) {
	localctx = NewAlterReadwriteSplittingRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RDLStatementParserRULE_alterReadwriteSplittingRule)
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
		p.SetState(62)
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(63)
		p.Match(RDLStatementParserREADWRITE_SPLITTING)
	}
	{
		p.SetState(64)
		p.Match(RDLStatementParserRULE)
	}
	{
		p.SetState(65)
		p.ReadwriteSplittingRuleDefinition()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(66)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(67)
			p.ReadwriteSplittingRuleDefinition()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropReadwriteSplittingRuleContext is an interface to support dynamic dispatch.
type IDropReadwriteSplittingRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropReadwriteSplittingRuleContext differentiates from other interfaces.
	IsDropReadwriteSplittingRuleContext()
}

type DropReadwriteSplittingRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropReadwriteSplittingRuleContext() *DropReadwriteSplittingRuleContext {
	var p = new(DropReadwriteSplittingRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropReadwriteSplittingRule
	return p
}

func (*DropReadwriteSplittingRuleContext) IsDropReadwriteSplittingRuleContext() {}

func NewDropReadwriteSplittingRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropReadwriteSplittingRuleContext {
	var p = new(DropReadwriteSplittingRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropReadwriteSplittingRule

	return p
}

func (s *DropReadwriteSplittingRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropReadwriteSplittingRuleContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropReadwriteSplittingRuleContext) READWRITE_SPLITTING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREADWRITE_SPLITTING, 0)
}

func (s *DropReadwriteSplittingRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *DropReadwriteSplittingRuleContext) AllRuleName() []IRuleNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleNameContext)(nil)).Elem())
	var tst = make([]IRuleNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleNameContext)
		}
	}

	return tst
}

func (s *DropReadwriteSplittingRuleContext) RuleName(i int) IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *DropReadwriteSplittingRuleContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropReadwriteSplittingRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropReadwriteSplittingRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropReadwriteSplittingRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropReadwriteSplittingRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropReadwriteSplittingRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropReadwriteSplittingRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropReadwriteSplittingRule() (localctx IDropReadwriteSplittingRuleContext) {
	localctx = NewDropReadwriteSplittingRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RDLStatementParserRULE_dropReadwriteSplittingRule)
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
		p.SetState(73)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(74)
		p.Match(RDLStatementParserREADWRITE_SPLITTING)
	}
	{
		p.SetState(75)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(76)
			p.IfExists()
		}

	}
	{
		p.SetState(79)
		p.RuleName()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(80)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(81)
			p.RuleName()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReadwriteSplittingRuleDefinitionContext is an interface to support dynamic dispatch.
type IReadwriteSplittingRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadwriteSplittingRuleDefinitionContext differentiates from other interfaces.
	IsReadwriteSplittingRuleDefinitionContext()
}

type ReadwriteSplittingRuleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadwriteSplittingRuleDefinitionContext() *ReadwriteSplittingRuleDefinitionContext {
	var p = new(ReadwriteSplittingRuleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_readwriteSplittingRuleDefinition
	return p
}

func (*ReadwriteSplittingRuleDefinitionContext) IsReadwriteSplittingRuleDefinitionContext() {}

func NewReadwriteSplittingRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadwriteSplittingRuleDefinitionContext {
	var p = new(ReadwriteSplittingRuleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_readwriteSplittingRuleDefinition

	return p
}

func (s *ReadwriteSplittingRuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadwriteSplittingRuleDefinitionContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *ReadwriteSplittingRuleDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ReadwriteSplittingRuleDefinitionContext) DataSourceDefinition() IDataSourceDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataSourceDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataSourceDefinitionContext)
}

func (s *ReadwriteSplittingRuleDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ReadwriteSplittingRuleDefinitionContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ReadwriteSplittingRuleDefinitionContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ReadwriteSplittingRuleDefinitionContext) TransactionalReadQueryStrategy() ITransactionalReadQueryStrategyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransactionalReadQueryStrategyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransactionalReadQueryStrategyContext)
}

func (s *ReadwriteSplittingRuleDefinitionContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *ReadwriteSplittingRuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadwriteSplittingRuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadwriteSplittingRuleDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitReadwriteSplittingRuleDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ReadwriteSplittingRuleDefinition() (localctx IReadwriteSplittingRuleDefinitionContext) {
	localctx = NewReadwriteSplittingRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RDLStatementParserRULE_readwriteSplittingRuleDefinition)
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
		p.SetState(87)
		p.RuleName()
	}
	{
		p.SetState(88)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(89)
		p.DataSourceDefinition()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(90)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(91)
			p.TransactionalReadQueryStrategy()
		}

	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(94)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(95)
			p.AlgorithmDefinition()
		}

	}
	{
		p.SetState(98)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IDataSourceDefinitionContext is an interface to support dynamic dispatch.
type IDataSourceDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataSourceDefinitionContext differentiates from other interfaces.
	IsDataSourceDefinitionContext()
}

type DataSourceDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataSourceDefinitionContext() *DataSourceDefinitionContext {
	var p = new(DataSourceDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dataSourceDefinition
	return p
}

func (*DataSourceDefinitionContext) IsDataSourceDefinitionContext() {}

func NewDataSourceDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataSourceDefinitionContext {
	var p = new(DataSourceDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dataSourceDefinition

	return p
}

func (s *DataSourceDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DataSourceDefinitionContext) WriteStorageUnit() IWriteStorageUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWriteStorageUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWriteStorageUnitContext)
}

func (s *DataSourceDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *DataSourceDefinitionContext) ReadStorageUnits() IReadStorageUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadStorageUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReadStorageUnitsContext)
}

func (s *DataSourceDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSourceDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataSourceDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDataSourceDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DataSourceDefinition() (localctx IDataSourceDefinitionContext) {
	localctx = NewDataSourceDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RDLStatementParserRULE_dataSourceDefinition)

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
		p.SetState(100)
		p.WriteStorageUnit()
	}
	{
		p.SetState(101)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(102)
		p.ReadStorageUnits()
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
	p.EnterRule(localctx, 10, RDLStatementParserRULE_ruleName)

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
		p.SetState(104)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IWriteStorageUnitContext is an interface to support dynamic dispatch.
type IWriteStorageUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWriteStorageUnitContext differentiates from other interfaces.
	IsWriteStorageUnitContext()
}

type WriteStorageUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteStorageUnitContext() *WriteStorageUnitContext {
	var p = new(WriteStorageUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_writeStorageUnit
	return p
}

func (*WriteStorageUnitContext) IsWriteStorageUnitContext() {}

func NewWriteStorageUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteStorageUnitContext {
	var p = new(WriteStorageUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_writeStorageUnit

	return p
}

func (s *WriteStorageUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteStorageUnitContext) WRITE_STORAGE_UNIT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserWRITE_STORAGE_UNIT, 0)
}

func (s *WriteStorageUnitContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *WriteStorageUnitContext) WriteStorageUnitName() IWriteStorageUnitNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWriteStorageUnitNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWriteStorageUnitNameContext)
}

func (s *WriteStorageUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteStorageUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WriteStorageUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitWriteStorageUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) WriteStorageUnit() (localctx IWriteStorageUnitContext) {
	localctx = NewWriteStorageUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RDLStatementParserRULE_writeStorageUnit)

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
		p.SetState(106)
		p.Match(RDLStatementParserWRITE_STORAGE_UNIT)
	}
	{
		p.SetState(107)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(108)
		p.WriteStorageUnitName()
	}

	return localctx
}

// IReadStorageUnitsContext is an interface to support dynamic dispatch.
type IReadStorageUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadStorageUnitsContext differentiates from other interfaces.
	IsReadStorageUnitsContext()
}

type ReadStorageUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadStorageUnitsContext() *ReadStorageUnitsContext {
	var p = new(ReadStorageUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_readStorageUnits
	return p
}

func (*ReadStorageUnitsContext) IsReadStorageUnitsContext() {}

func NewReadStorageUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadStorageUnitsContext {
	var p = new(ReadStorageUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_readStorageUnits

	return p
}

func (s *ReadStorageUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadStorageUnitsContext) READ_STORAGE_UNITS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREAD_STORAGE_UNITS, 0)
}

func (s *ReadStorageUnitsContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ReadStorageUnitsContext) ReadStorageUnitsNames() IReadStorageUnitsNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadStorageUnitsNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReadStorageUnitsNamesContext)
}

func (s *ReadStorageUnitsContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ReadStorageUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadStorageUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadStorageUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitReadStorageUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ReadStorageUnits() (localctx IReadStorageUnitsContext) {
	localctx = NewReadStorageUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RDLStatementParserRULE_readStorageUnits)

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
		p.SetState(110)
		p.Match(RDLStatementParserREAD_STORAGE_UNITS)
	}
	{
		p.SetState(111)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(112)
		p.ReadStorageUnitsNames()
	}
	{
		p.SetState(113)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// ITransactionalReadQueryStrategyContext is an interface to support dynamic dispatch.
type ITransactionalReadQueryStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransactionalReadQueryStrategyContext differentiates from other interfaces.
	IsTransactionalReadQueryStrategyContext()
}

type TransactionalReadQueryStrategyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransactionalReadQueryStrategyContext() *TransactionalReadQueryStrategyContext {
	var p = new(TransactionalReadQueryStrategyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_transactionalReadQueryStrategy
	return p
}

func (*TransactionalReadQueryStrategyContext) IsTransactionalReadQueryStrategyContext() {}

func NewTransactionalReadQueryStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransactionalReadQueryStrategyContext {
	var p = new(TransactionalReadQueryStrategyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_transactionalReadQueryStrategy

	return p
}

func (s *TransactionalReadQueryStrategyContext) GetParser() antlr.Parser { return s.parser }

func (s *TransactionalReadQueryStrategyContext) TRANSACTIONAL_READ_QUERY_STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTRANSACTIONAL_READ_QUERY_STRATEGY, 0)
}

func (s *TransactionalReadQueryStrategyContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *TransactionalReadQueryStrategyContext) TransactionalReadQueryStrategyName() ITransactionalReadQueryStrategyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransactionalReadQueryStrategyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransactionalReadQueryStrategyNameContext)
}

func (s *TransactionalReadQueryStrategyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransactionalReadQueryStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransactionalReadQueryStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitTransactionalReadQueryStrategy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) TransactionalReadQueryStrategy() (localctx ITransactionalReadQueryStrategyContext) {
	localctx = NewTransactionalReadQueryStrategyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RDLStatementParserRULE_transactionalReadQueryStrategy)

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
		p.SetState(115)
		p.Match(RDLStatementParserTRANSACTIONAL_READ_QUERY_STRATEGY)
	}
	{
		p.SetState(116)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(117)
		p.TransactionalReadQueryStrategyName()
	}

	return localctx
}

// IWriteStorageUnitNameContext is an interface to support dynamic dispatch.
type IWriteStorageUnitNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWriteStorageUnitNameContext differentiates from other interfaces.
	IsWriteStorageUnitNameContext()
}

type WriteStorageUnitNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteStorageUnitNameContext() *WriteStorageUnitNameContext {
	var p = new(WriteStorageUnitNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_writeStorageUnitName
	return p
}

func (*WriteStorageUnitNameContext) IsWriteStorageUnitNameContext() {}

func NewWriteStorageUnitNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteStorageUnitNameContext {
	var p = new(WriteStorageUnitNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_writeStorageUnitName

	return p
}

func (s *WriteStorageUnitNameContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteStorageUnitNameContext) StorageUnitName() IStorageUnitNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStorageUnitNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStorageUnitNameContext)
}

func (s *WriteStorageUnitNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteStorageUnitNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WriteStorageUnitNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitWriteStorageUnitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) WriteStorageUnitName() (localctx IWriteStorageUnitNameContext) {
	localctx = NewWriteStorageUnitNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RDLStatementParserRULE_writeStorageUnitName)

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
		p.SetState(119)
		p.StorageUnitName()
	}

	return localctx
}

// IReadStorageUnitsNamesContext is an interface to support dynamic dispatch.
type IReadStorageUnitsNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadStorageUnitsNamesContext differentiates from other interfaces.
	IsReadStorageUnitsNamesContext()
}

type ReadStorageUnitsNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadStorageUnitsNamesContext() *ReadStorageUnitsNamesContext {
	var p = new(ReadStorageUnitsNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_readStorageUnitsNames
	return p
}

func (*ReadStorageUnitsNamesContext) IsReadStorageUnitsNamesContext() {}

func NewReadStorageUnitsNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadStorageUnitsNamesContext {
	var p = new(ReadStorageUnitsNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_readStorageUnitsNames

	return p
}

func (s *ReadStorageUnitsNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadStorageUnitsNamesContext) AllStorageUnitName() []IStorageUnitNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStorageUnitNameContext)(nil)).Elem())
	var tst = make([]IStorageUnitNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStorageUnitNameContext)
		}
	}

	return tst
}

func (s *ReadStorageUnitsNamesContext) StorageUnitName(i int) IStorageUnitNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStorageUnitNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStorageUnitNameContext)
}

func (s *ReadStorageUnitsNamesContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ReadStorageUnitsNamesContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ReadStorageUnitsNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadStorageUnitsNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadStorageUnitsNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitReadStorageUnitsNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ReadStorageUnitsNames() (localctx IReadStorageUnitsNamesContext) {
	localctx = NewReadStorageUnitsNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RDLStatementParserRULE_readStorageUnitsNames)
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
		p.StorageUnitName()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(122)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(123)
			p.StorageUnitName()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITransactionalReadQueryStrategyNameContext is an interface to support dynamic dispatch.
type ITransactionalReadQueryStrategyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransactionalReadQueryStrategyNameContext differentiates from other interfaces.
	IsTransactionalReadQueryStrategyNameContext()
}

type TransactionalReadQueryStrategyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransactionalReadQueryStrategyNameContext() *TransactionalReadQueryStrategyNameContext {
	var p = new(TransactionalReadQueryStrategyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_transactionalReadQueryStrategyName
	return p
}

func (*TransactionalReadQueryStrategyNameContext) IsTransactionalReadQueryStrategyNameContext() {}

func NewTransactionalReadQueryStrategyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransactionalReadQueryStrategyNameContext {
	var p = new(TransactionalReadQueryStrategyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_transactionalReadQueryStrategyName

	return p
}

func (s *TransactionalReadQueryStrategyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TransactionalReadQueryStrategyNameContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *TransactionalReadQueryStrategyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransactionalReadQueryStrategyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransactionalReadQueryStrategyNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitTransactionalReadQueryStrategyName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) TransactionalReadQueryStrategyName() (localctx ITransactionalReadQueryStrategyNameContext) {
	localctx = NewTransactionalReadQueryStrategyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RDLStatementParserRULE_transactionalReadQueryStrategyName)

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
		p.SetState(129)
		p.Match(RDLStatementParserSTRING_)
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
	p.EnterRule(localctx, 24, RDLStatementParserRULE_ifExists)

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
		p.SetState(131)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 26, RDLStatementParserRULE_ifNotExists)

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
		p.SetState(134)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(135)
		p.Match(RDLStatementParserNOT)
	}
	{
		p.SetState(136)
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
	p.EnterRule(localctx, 28, RDLStatementParserRULE_literal)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserMINUS_, RDLStatementParserINT_:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RDLStatementParserMINUS_ {
			{
				p.SetState(139)
				p.Match(RDLStatementParserMINUS_)
			}

		}
		{
			p.SetState(142)
			p.Match(RDLStatementParserINT_)
		}

	case RDLStatementParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(143)
			p.Match(RDLStatementParserTRUE)
		}

	case RDLStatementParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)
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
	p.EnterRule(localctx, 30, RDLStatementParserRULE_algorithmDefinition)
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
		p.SetState(147)
		p.Match(RDLStatementParserTYPE)
	}
	{
		p.SetState(148)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(149)
		p.Match(RDLStatementParserNAME)
	}
	{
		p.SetState(150)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(151)
		p.AlgorithmTypeName()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(152)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(153)
			p.PropertiesDefinition()
		}

	}
	{
		p.SetState(156)
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

func (s *AlgorithmTypeNameContext) BuildInReadQueryLoadBalanceAlgorithmType() IBuildInReadQueryLoadBalanceAlgorithmTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildInReadQueryLoadBalanceAlgorithmTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildInReadQueryLoadBalanceAlgorithmTypeContext)
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
	p.EnterRule(localctx, 32, RDLStatementParserRULE_algorithmTypeName)

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

	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserROUND_ROBIN, RDLStatementParserRANDOM, RDLStatementParserWEIGHT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.BuildInReadQueryLoadBalanceAlgorithmType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuildInReadQueryLoadBalanceAlgorithmTypeContext is an interface to support dynamic dispatch.
type IBuildInReadQueryLoadBalanceAlgorithmTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildInReadQueryLoadBalanceAlgorithmTypeContext differentiates from other interfaces.
	IsBuildInReadQueryLoadBalanceAlgorithmTypeContext()
}

type BuildInReadQueryLoadBalanceAlgorithmTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildInReadQueryLoadBalanceAlgorithmTypeContext() *BuildInReadQueryLoadBalanceAlgorithmTypeContext {
	var p = new(BuildInReadQueryLoadBalanceAlgorithmTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildInReadQueryLoadBalanceAlgorithmType
	return p
}

func (*BuildInReadQueryLoadBalanceAlgorithmTypeContext) IsBuildInReadQueryLoadBalanceAlgorithmTypeContext() {
}

func NewBuildInReadQueryLoadBalanceAlgorithmTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildInReadQueryLoadBalanceAlgorithmTypeContext {
	var p = new(BuildInReadQueryLoadBalanceAlgorithmTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildInReadQueryLoadBalanceAlgorithmType

	return p
}

func (s *BuildInReadQueryLoadBalanceAlgorithmTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildInReadQueryLoadBalanceAlgorithmTypeContext) ROUND_ROBIN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserROUND_ROBIN, 0)
}

func (s *BuildInReadQueryLoadBalanceAlgorithmTypeContext) RANDOM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRANDOM, 0)
}

func (s *BuildInReadQueryLoadBalanceAlgorithmTypeContext) WEIGHT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserWEIGHT, 0)
}

func (s *BuildInReadQueryLoadBalanceAlgorithmTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInReadQueryLoadBalanceAlgorithmTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildInReadQueryLoadBalanceAlgorithmTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildInReadQueryLoadBalanceAlgorithmType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildInReadQueryLoadBalanceAlgorithmType() (localctx IBuildInReadQueryLoadBalanceAlgorithmTypeContext) {
	localctx = NewBuildInReadQueryLoadBalanceAlgorithmTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RDLStatementParserRULE_buildInReadQueryLoadBalanceAlgorithmType)
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
		p.SetState(162)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(RDLStatementParserROUND_ROBIN-68))|(1<<(RDLStatementParserRANDOM-68))|(1<<(RDLStatementParserWEIGHT-68)))) != 0) {
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
	p.EnterRule(localctx, 36, RDLStatementParserRULE_propertiesDefinition)
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
		p.SetState(164)
		p.Match(RDLStatementParserPROPERTIES)
	}
	{
		p.SetState(165)
		p.Match(RDLStatementParserLP_)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserSTRING_ {
		{
			p.SetState(166)
			p.Properties()
		}

	}
	{
		p.SetState(169)
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
	p.EnterRule(localctx, 38, RDLStatementParserRULE_properties)
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
		p.SetState(171)
		p.Property()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(172)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(173)
			p.Property()
		}

		p.SetState(178)
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
	p.EnterRule(localctx, 40, RDLStatementParserRULE_property)

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
		p.SetState(179)

		var _m = p.Match(RDLStatementParserSTRING_)

		localctx.(*PropertyContext).key = _m
	}
	{
		p.SetState(180)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(181)

		var _x = p.Literal()

		localctx.(*PropertyContext).value = _x
	}

	return localctx
}

// IDatabaseNameContext is an interface to support dynamic dispatch.
type IDatabaseNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatabaseNameContext differentiates from other interfaces.
	IsDatabaseNameContext()
}

type DatabaseNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabaseNameContext() *DatabaseNameContext {
	var p = new(DatabaseNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_databaseName
	return p
}

func (*DatabaseNameContext) IsDatabaseNameContext() {}

func NewDatabaseNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabaseNameContext {
	var p = new(DatabaseNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_databaseName

	return p
}

func (s *DatabaseNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabaseNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *DatabaseNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabaseNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabaseNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDatabaseName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DatabaseName() (localctx IDatabaseNameContext) {
	localctx = NewDatabaseNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RDLStatementParserRULE_databaseName)

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
		p.SetState(183)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IGroupNameContext is an interface to support dynamic dispatch.
type IGroupNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupNameContext differentiates from other interfaces.
	IsGroupNameContext()
}

type GroupNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupNameContext() *GroupNameContext {
	var p = new(GroupNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_groupName
	return p
}

func (*GroupNameContext) IsGroupNameContext() {}

func NewGroupNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupNameContext {
	var p = new(GroupNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_groupName

	return p
}

func (s *GroupNameContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *GroupNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitGroupName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) GroupName() (localctx IGroupNameContext) {
	localctx = NewGroupNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RDLStatementParserRULE_groupName)

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
		p.SetState(185)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IStorageUnitNameContext is an interface to support dynamic dispatch.
type IStorageUnitNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStorageUnitNameContext differentiates from other interfaces.
	IsStorageUnitNameContext()
}

type StorageUnitNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStorageUnitNameContext() *StorageUnitNameContext {
	var p = new(StorageUnitNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_storageUnitName
	return p
}

func (*StorageUnitNameContext) IsStorageUnitNameContext() {}

func NewStorageUnitNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StorageUnitNameContext {
	var p = new(StorageUnitNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_storageUnitName

	return p
}

func (s *StorageUnitNameContext) GetParser() antlr.Parser { return s.parser }

func (s *StorageUnitNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *StorageUnitNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StorageUnitNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StorageUnitNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitStorageUnitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) StorageUnitName() (localctx IStorageUnitNameContext) {
	localctx = NewStorageUnitNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RDLStatementParserRULE_storageUnitName)

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
		p.SetState(187)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}
