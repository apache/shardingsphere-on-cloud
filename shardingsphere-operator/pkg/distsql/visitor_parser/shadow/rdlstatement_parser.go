// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 86, 218,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 53, 10, 2, 3, 2, 3, 2, 3, 2,
	7, 2, 58, 10, 2, 12, 2, 14, 2, 61, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 69, 10, 3, 12, 3, 14, 3, 72, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 78, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 83, 10, 4, 12, 4, 14, 4, 86, 11,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 92, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 97,
	10, 5, 12, 5, 14, 5, 100, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 107,
	10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 116, 10, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 137, 10, 9, 12, 9, 14, 9, 140,
	11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 149, 10, 10,
	12, 10, 14, 10, 152, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 5, 17, 173, 10, 17, 3, 17, 3, 17, 3, 17, 5, 17, 178,
	10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 187, 10,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 5, 19, 193, 10, 19, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 5, 21, 200, 10, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 7,
	22, 207, 10, 22, 12, 22, 14, 22, 210, 11, 22, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 2, 2, 25, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 2, 3, 3, 2, 74, 76, 2,
	213, 2, 48, 3, 2, 2, 2, 4, 62, 3, 2, 2, 2, 6, 73, 3, 2, 2, 2, 8, 87, 3,
	2, 2, 2, 10, 101, 3, 2, 2, 2, 12, 110, 3, 2, 2, 2, 14, 117, 3, 2, 2, 2,
	16, 123, 3, 2, 2, 2, 18, 143, 3, 2, 2, 2, 20, 155, 3, 2, 2, 2, 22, 157,
	3, 2, 2, 2, 24, 159, 3, 2, 2, 2, 26, 161, 3, 2, 2, 2, 28, 163, 3, 2, 2,
	2, 30, 166, 3, 2, 2, 2, 32, 177, 3, 2, 2, 2, 34, 179, 3, 2, 2, 2, 36, 192,
	3, 2, 2, 2, 38, 194, 3, 2, 2, 2, 40, 196, 3, 2, 2, 2, 42, 203, 3, 2, 2,
	2, 44, 211, 3, 2, 2, 2, 46, 215, 3, 2, 2, 2, 48, 49, 7, 48, 2, 2, 49, 50,
	7, 52, 2, 2, 50, 52, 7, 54, 2, 2, 51, 53, 5, 30, 16, 2, 52, 51, 3, 2, 2,
	2, 52, 53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 59, 5, 16, 9, 2, 55, 56,
	7, 36, 2, 2, 56, 58, 5, 16, 9, 2, 57, 55, 3, 2, 2, 2, 58, 61, 3, 2, 2,
	2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 3, 3, 2, 2, 2, 61, 59, 3,
	2, 2, 2, 62, 63, 7, 49, 2, 2, 63, 64, 7, 52, 2, 2, 64, 65, 7, 54, 2, 2,
	65, 70, 5, 16, 9, 2, 66, 67, 7, 36, 2, 2, 67, 69, 5, 16, 9, 2, 68, 66,
	3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2,
	71, 5, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 74, 7, 50, 2, 2, 74, 75, 7,
	52, 2, 2, 75, 77, 7, 54, 2, 2, 76, 78, 5, 28, 15, 2, 77, 76, 3, 2, 2, 2,
	77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 84, 5, 46, 24, 2, 80, 81, 7,
	36, 2, 2, 81, 83, 5, 46, 24, 2, 82, 80, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2,
	84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 7, 3, 2, 2, 2, 86, 84, 3, 2,
	2, 2, 87, 88, 7, 50, 2, 2, 88, 89, 7, 52, 2, 2, 89, 91, 7, 62, 2, 2, 90,
	92, 5, 28, 15, 2, 91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93, 3, 2,
	2, 2, 93, 98, 5, 26, 14, 2, 94, 95, 7, 36, 2, 2, 95, 97, 5, 26, 14, 2,
	96, 94, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3,
	2, 2, 2, 99, 9, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 102, 7, 48, 2, 2,
	102, 103, 7, 70, 2, 2, 103, 104, 7, 52, 2, 2, 104, 106, 7, 62, 2, 2, 105,
	107, 5, 30, 16, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108,
	3, 2, 2, 2, 108, 109, 5, 34, 18, 2, 109, 11, 3, 2, 2, 2, 110, 111, 7, 50,
	2, 2, 111, 112, 7, 70, 2, 2, 112, 113, 7, 52, 2, 2, 113, 115, 7, 62, 2,
	2, 114, 116, 5, 28, 15, 2, 115, 114, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2,
	116, 13, 3, 2, 2, 2, 117, 118, 7, 49, 2, 2, 118, 119, 7, 70, 2, 2, 119,
	120, 7, 52, 2, 2, 120, 121, 7, 62, 2, 2, 121, 122, 5, 34, 18, 2, 122, 15,
	3, 2, 2, 2, 123, 124, 5, 46, 24, 2, 124, 125, 7, 30, 2, 2, 125, 126, 7,
	53, 2, 2, 126, 127, 7, 23, 2, 2, 127, 128, 5, 20, 11, 2, 128, 129, 7, 36,
	2, 2, 129, 130, 7, 52, 2, 2, 130, 131, 7, 23, 2, 2, 131, 132, 5, 22, 12,
	2, 132, 133, 7, 36, 2, 2, 133, 138, 5, 18, 10, 2, 134, 135, 7, 36, 2, 2,
	135, 137, 5, 18, 10, 2, 136, 134, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138,
	136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 141, 3, 2, 2, 2, 140, 138,
	3, 2, 2, 2, 141, 142, 7, 31, 2, 2, 142, 17, 3, 2, 2, 2, 143, 144, 5, 24,
	13, 2, 144, 145, 7, 30, 2, 2, 145, 150, 5, 34, 18, 2, 146, 147, 7, 36,
	2, 2, 147, 149, 5, 34, 18, 2, 148, 146, 3, 2, 2, 2, 149, 152, 3, 2, 2,
	2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 153, 3, 2, 2, 2, 152,
	150, 3, 2, 2, 2, 153, 154, 7, 31, 2, 2, 154, 19, 3, 2, 2, 2, 155, 156,
	7, 79, 2, 2, 156, 21, 3, 2, 2, 2, 157, 158, 7, 79, 2, 2, 158, 23, 3, 2,
	2, 2, 159, 160, 7, 79, 2, 2, 160, 25, 3, 2, 2, 2, 161, 162, 7, 79, 2, 2,
	162, 27, 3, 2, 2, 2, 163, 164, 7, 71, 2, 2, 164, 165, 7, 72, 2, 2, 165,
	29, 3, 2, 2, 2, 166, 167, 7, 71, 2, 2, 167, 168, 7, 77, 2, 2, 168, 169,
	7, 72, 2, 2, 169, 31, 3, 2, 2, 2, 170, 178, 7, 80, 2, 2, 171, 173, 7, 15,
	2, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 178, 7, 81, 2, 2, 175, 178, 7, 46, 2, 2, 176, 178, 7, 47, 2, 2, 177,
	170, 3, 2, 2, 2, 177, 172, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 176,
	3, 2, 2, 2, 178, 33, 3, 2, 2, 2, 179, 180, 7, 58, 2, 2, 180, 181, 7, 30,
	2, 2, 181, 182, 7, 59, 2, 2, 182, 183, 7, 23, 2, 2, 183, 186, 5, 36, 19,
	2, 184, 185, 7, 36, 2, 2, 185, 187, 5, 40, 21, 2, 186, 184, 3, 2, 2, 2,
	186, 187, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 7, 31, 2, 2, 189,
	35, 3, 2, 2, 2, 190, 193, 7, 80, 2, 2, 191, 193, 5, 38, 20, 2, 192, 190,
	3, 2, 2, 2, 192, 191, 3, 2, 2, 2, 193, 37, 3, 2, 2, 2, 194, 195, 9, 2,
	2, 2, 195, 39, 3, 2, 2, 2, 196, 197, 7, 60, 2, 2, 197, 199, 7, 30, 2, 2,
	198, 200, 5, 42, 22, 2, 199, 198, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200,
	201, 3, 2, 2, 2, 201, 202, 7, 31, 2, 2, 202, 41, 3, 2, 2, 2, 203, 208,
	5, 44, 23, 2, 204, 205, 7, 36, 2, 2, 205, 207, 5, 44, 23, 2, 206, 204,
	3, 2, 2, 2, 207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 209, 3, 2,
	2, 2, 209, 43, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 212, 7, 80, 2, 2,
	212, 213, 7, 23, 2, 2, 213, 214, 5, 32, 17, 2, 214, 45, 3, 2, 2, 2, 215,
	216, 7, 79, 2, 2, 216, 47, 3, 2, 2, 2, 19, 52, 59, 70, 77, 84, 91, 98,
	106, 115, 138, 150, 172, 177, 186, 192, 199, 208,
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
	"", "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'",
}
var symbolicNames = []string{
	"", "AND_", "OR_", "NOT_", "TILDE_", "VERTICALBAR_", "AMPERSAND_", "SIGNEDLEFTSHIFT_",
	"SIGNEDRIGHTSHIFT_", "CARET_", "MOD_", "COLON_", "PLUS_", "MINUS_", "ASTERISK_",
	"SLASH_", "BACKSLASH_", "DOT_", "DOTASTERISK_", "SAFEEQ_", "DEQ_", "EQ_",
	"NEQ_", "GT_", "GTE_", "LT_", "LTE_", "POUND_", "LP_", "RP_", "LBE_", "RBE_",
	"LBT_", "RBT_", "COMMA_", "DQ_", "SQ_", "BQ_", "QUESTION_", "AT_", "SEMI_",
	"JSONSEPARATOR_", "UL_", "WS", "TRUE", "FALSE", "CREATE", "ALTER", "DROP",
	"SHOW", "SHADOW", "SOURCE", "RULE", "FROM", "RESOURCES", "TABLE", "TYPE",
	"NAME", "PROPERTIES", "RULES", "ALGORITHM", "ALGORITHMS", "SET", "ADD",
	"DATABASE_VALUE", "TABLE_VALUE", "STATUS", "CLEAR", "DEFAULT", "IF", "EXISTS",
	"COUNT", "VALUE_MATCH", "REGEX_MATCH", "SQL_HINT", "NOT", "FOR_GENERATOR",
	"IDENTIFIER_", "STRING_", "INT_", "HEX_", "NUMBER_", "HEXDIGIT_", "BITNUM_",
	"BOOL_",
}

var ruleNames = []string{
	"createShadowRule", "alterShadowRule", "dropShadowRule", "dropShadowAlgorithm",
	"createDefaultShadowAlgorithm", "dropDefaultShadowAlgorithm", "alterDefaultShadowAlgorithm",
	"shadowRuleDefinition", "shadowTableRule", "source", "shadow", "tableName",
	"algorithmName", "ifExists", "ifNotExists", "literal", "algorithmDefinition",
	"algorithmTypeName", "buildInShadowAlgorithmType", "propertiesDefinition",
	"properties", "property", "ruleName",
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
	RDLStatementParserEOF               = antlr.TokenEOF
	RDLStatementParserAND_              = 1
	RDLStatementParserOR_               = 2
	RDLStatementParserNOT_              = 3
	RDLStatementParserTILDE_            = 4
	RDLStatementParserVERTICALBAR_      = 5
	RDLStatementParserAMPERSAND_        = 6
	RDLStatementParserSIGNEDLEFTSHIFT_  = 7
	RDLStatementParserSIGNEDRIGHTSHIFT_ = 8
	RDLStatementParserCARET_            = 9
	RDLStatementParserMOD_              = 10
	RDLStatementParserCOLON_            = 11
	RDLStatementParserPLUS_             = 12
	RDLStatementParserMINUS_            = 13
	RDLStatementParserASTERISK_         = 14
	RDLStatementParserSLASH_            = 15
	RDLStatementParserBACKSLASH_        = 16
	RDLStatementParserDOT_              = 17
	RDLStatementParserDOTASTERISK_      = 18
	RDLStatementParserSAFEEQ_           = 19
	RDLStatementParserDEQ_              = 20
	RDLStatementParserEQ_               = 21
	RDLStatementParserNEQ_              = 22
	RDLStatementParserGT_               = 23
	RDLStatementParserGTE_              = 24
	RDLStatementParserLT_               = 25
	RDLStatementParserLTE_              = 26
	RDLStatementParserPOUND_            = 27
	RDLStatementParserLP_               = 28
	RDLStatementParserRP_               = 29
	RDLStatementParserLBE_              = 30
	RDLStatementParserRBE_              = 31
	RDLStatementParserLBT_              = 32
	RDLStatementParserRBT_              = 33
	RDLStatementParserCOMMA_            = 34
	RDLStatementParserDQ_               = 35
	RDLStatementParserSQ_               = 36
	RDLStatementParserBQ_               = 37
	RDLStatementParserQUESTION_         = 38
	RDLStatementParserAT_               = 39
	RDLStatementParserSEMI_             = 40
	RDLStatementParserJSONSEPARATOR_    = 41
	RDLStatementParserUL_               = 42
	RDLStatementParserWS                = 43
	RDLStatementParserTRUE              = 44
	RDLStatementParserFALSE             = 45
	RDLStatementParserCREATE            = 46
	RDLStatementParserALTER             = 47
	RDLStatementParserDROP              = 48
	RDLStatementParserSHOW              = 49
	RDLStatementParserSHADOW            = 50
	RDLStatementParserSOURCE            = 51
	RDLStatementParserRULE              = 52
	RDLStatementParserFROM              = 53
	RDLStatementParserRESOURCES         = 54
	RDLStatementParserTABLE             = 55
	RDLStatementParserTYPE              = 56
	RDLStatementParserNAME              = 57
	RDLStatementParserPROPERTIES        = 58
	RDLStatementParserRULES             = 59
	RDLStatementParserALGORITHM         = 60
	RDLStatementParserALGORITHMS        = 61
	RDLStatementParserSET               = 62
	RDLStatementParserADD               = 63
	RDLStatementParserDATABASE_VALUE    = 64
	RDLStatementParserTABLE_VALUE       = 65
	RDLStatementParserSTATUS            = 66
	RDLStatementParserCLEAR             = 67
	RDLStatementParserDEFAULT           = 68
	RDLStatementParserIF                = 69
	RDLStatementParserEXISTS            = 70
	RDLStatementParserCOUNT             = 71
	RDLStatementParserVALUE_MATCH       = 72
	RDLStatementParserREGEX_MATCH       = 73
	RDLStatementParserSQL_HINT          = 74
	RDLStatementParserNOT               = 75
	RDLStatementParserFOR_GENERATOR     = 76
	RDLStatementParserIDENTIFIER_       = 77
	RDLStatementParserSTRING_           = 78
	RDLStatementParserINT_              = 79
	RDLStatementParserHEX_              = 80
	RDLStatementParserNUMBER_           = 81
	RDLStatementParserHEXDIGIT_         = 82
	RDLStatementParserBITNUM_           = 83
	RDLStatementParserBOOL_             = 84
)

// RDLStatementParser rules.
const (
	RDLStatementParserRULE_createShadowRule             = 0
	RDLStatementParserRULE_alterShadowRule              = 1
	RDLStatementParserRULE_dropShadowRule               = 2
	RDLStatementParserRULE_dropShadowAlgorithm          = 3
	RDLStatementParserRULE_createDefaultShadowAlgorithm = 4
	RDLStatementParserRULE_dropDefaultShadowAlgorithm   = 5
	RDLStatementParserRULE_alterDefaultShadowAlgorithm  = 6
	RDLStatementParserRULE_shadowRuleDefinition         = 7
	RDLStatementParserRULE_shadowTableRule              = 8
	RDLStatementParserRULE_source                       = 9
	RDLStatementParserRULE_shadow                       = 10
	RDLStatementParserRULE_tableName                    = 11
	RDLStatementParserRULE_algorithmName                = 12
	RDLStatementParserRULE_ifExists                     = 13
	RDLStatementParserRULE_ifNotExists                  = 14
	RDLStatementParserRULE_literal                      = 15
	RDLStatementParserRULE_algorithmDefinition          = 16
	RDLStatementParserRULE_algorithmTypeName            = 17
	RDLStatementParserRULE_buildInShadowAlgorithmType   = 18
	RDLStatementParserRULE_propertiesDefinition         = 19
	RDLStatementParserRULE_properties                   = 20
	RDLStatementParserRULE_property                     = 21
	RDLStatementParserRULE_ruleName                     = 22
)

// ICreateShadowRuleContext is an interface to support dynamic dispatch.
type ICreateShadowRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateShadowRuleContext differentiates from other interfaces.
	IsCreateShadowRuleContext()
}

type CreateShadowRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateShadowRuleContext() *CreateShadowRuleContext {
	var p = new(CreateShadowRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createShadowRule
	return p
}

func (*CreateShadowRuleContext) IsCreateShadowRuleContext() {}

func NewCreateShadowRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateShadowRuleContext {
	var p = new(CreateShadowRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createShadowRule

	return p
}

func (s *CreateShadowRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateShadowRuleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateShadowRuleContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *CreateShadowRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *CreateShadowRuleContext) AllShadowRuleDefinition() []IShadowRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShadowRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IShadowRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShadowRuleDefinitionContext)
		}
	}

	return tst
}

func (s *CreateShadowRuleContext) ShadowRuleDefinition(i int) IShadowRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShadowRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShadowRuleDefinitionContext)
}

func (s *CreateShadowRuleContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateShadowRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *CreateShadowRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *CreateShadowRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateShadowRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateShadowRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateShadowRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateShadowRule() (localctx ICreateShadowRuleContext) {
	localctx = NewCreateShadowRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RDLStatementParserRULE_createShadowRule)
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
		p.SetState(46)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(47)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(48)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(49)
			p.IfNotExists()
		}

	}
	{
		p.SetState(52)
		p.ShadowRuleDefinition()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(53)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(54)
			p.ShadowRuleDefinition()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlterShadowRuleContext is an interface to support dynamic dispatch.
type IAlterShadowRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterShadowRuleContext differentiates from other interfaces.
	IsAlterShadowRuleContext()
}

type AlterShadowRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterShadowRuleContext() *AlterShadowRuleContext {
	var p = new(AlterShadowRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterShadowRule
	return p
}

func (*AlterShadowRuleContext) IsAlterShadowRuleContext() {}

func NewAlterShadowRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterShadowRuleContext {
	var p = new(AlterShadowRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterShadowRule

	return p
}

func (s *AlterShadowRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterShadowRuleContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterShadowRuleContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *AlterShadowRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *AlterShadowRuleContext) AllShadowRuleDefinition() []IShadowRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShadowRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IShadowRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShadowRuleDefinitionContext)
		}
	}

	return tst
}

func (s *AlterShadowRuleContext) ShadowRuleDefinition(i int) IShadowRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShadowRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShadowRuleDefinitionContext)
}

func (s *AlterShadowRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *AlterShadowRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *AlterShadowRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterShadowRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterShadowRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterShadowRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterShadowRule() (localctx IAlterShadowRuleContext) {
	localctx = NewAlterShadowRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RDLStatementParserRULE_alterShadowRule)
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
		p.SetState(60)
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(61)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(62)
		p.Match(RDLStatementParserRULE)
	}
	{
		p.SetState(63)
		p.ShadowRuleDefinition()
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(64)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(65)
			p.ShadowRuleDefinition()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropShadowRuleContext is an interface to support dynamic dispatch.
type IDropShadowRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropShadowRuleContext differentiates from other interfaces.
	IsDropShadowRuleContext()
}

type DropShadowRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropShadowRuleContext() *DropShadowRuleContext {
	var p = new(DropShadowRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropShadowRule
	return p
}

func (*DropShadowRuleContext) IsDropShadowRuleContext() {}

func NewDropShadowRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropShadowRuleContext {
	var p = new(DropShadowRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropShadowRule

	return p
}

func (s *DropShadowRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropShadowRuleContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropShadowRuleContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *DropShadowRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *DropShadowRuleContext) AllRuleName() []IRuleNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleNameContext)(nil)).Elem())
	var tst = make([]IRuleNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleNameContext)
		}
	}

	return tst
}

func (s *DropShadowRuleContext) RuleName(i int) IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *DropShadowRuleContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropShadowRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropShadowRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropShadowRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropShadowRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropShadowRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropShadowRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropShadowRule() (localctx IDropShadowRuleContext) {
	localctx = NewDropShadowRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RDLStatementParserRULE_dropShadowRule)
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
		p.SetState(71)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(72)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(73)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(74)
			p.IfExists()
		}

	}
	{
		p.SetState(77)
		p.RuleName()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(78)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(79)
			p.RuleName()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropShadowAlgorithmContext is an interface to support dynamic dispatch.
type IDropShadowAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropShadowAlgorithmContext differentiates from other interfaces.
	IsDropShadowAlgorithmContext()
}

type DropShadowAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropShadowAlgorithmContext() *DropShadowAlgorithmContext {
	var p = new(DropShadowAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropShadowAlgorithm
	return p
}

func (*DropShadowAlgorithmContext) IsDropShadowAlgorithmContext() {}

func NewDropShadowAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropShadowAlgorithmContext {
	var p = new(DropShadowAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropShadowAlgorithm

	return p
}

func (s *DropShadowAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *DropShadowAlgorithmContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropShadowAlgorithmContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *DropShadowAlgorithmContext) ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALGORITHM, 0)
}

func (s *DropShadowAlgorithmContext) AllAlgorithmName() []IAlgorithmNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAlgorithmNameContext)(nil)).Elem())
	var tst = make([]IAlgorithmNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAlgorithmNameContext)
		}
	}

	return tst
}

func (s *DropShadowAlgorithmContext) AlgorithmName(i int) IAlgorithmNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmNameContext)
}

func (s *DropShadowAlgorithmContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropShadowAlgorithmContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropShadowAlgorithmContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropShadowAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropShadowAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropShadowAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropShadowAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropShadowAlgorithm() (localctx IDropShadowAlgorithmContext) {
	localctx = NewDropShadowAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RDLStatementParserRULE_dropShadowAlgorithm)
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
		p.SetState(85)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(86)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(87)
		p.Match(RDLStatementParserALGORITHM)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(88)
			p.IfExists()
		}

	}
	{
		p.SetState(91)
		p.AlgorithmName()
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(92)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(93)
			p.AlgorithmName()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICreateDefaultShadowAlgorithmContext is an interface to support dynamic dispatch.
type ICreateDefaultShadowAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateDefaultShadowAlgorithmContext differentiates from other interfaces.
	IsCreateDefaultShadowAlgorithmContext()
}

type CreateDefaultShadowAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateDefaultShadowAlgorithmContext() *CreateDefaultShadowAlgorithmContext {
	var p = new(CreateDefaultShadowAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createDefaultShadowAlgorithm
	return p
}

func (*CreateDefaultShadowAlgorithmContext) IsCreateDefaultShadowAlgorithmContext() {}

func NewCreateDefaultShadowAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDefaultShadowAlgorithmContext {
	var p = new(CreateDefaultShadowAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createDefaultShadowAlgorithm

	return p
}

func (s *CreateDefaultShadowAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDefaultShadowAlgorithmContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateDefaultShadowAlgorithmContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDEFAULT, 0)
}

func (s *CreateDefaultShadowAlgorithmContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *CreateDefaultShadowAlgorithmContext) ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALGORITHM, 0)
}

func (s *CreateDefaultShadowAlgorithmContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *CreateDefaultShadowAlgorithmContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateDefaultShadowAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDefaultShadowAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDefaultShadowAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateDefaultShadowAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateDefaultShadowAlgorithm() (localctx ICreateDefaultShadowAlgorithmContext) {
	localctx = NewCreateDefaultShadowAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RDLStatementParserRULE_createDefaultShadowAlgorithm)
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
		p.SetState(99)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(100)
		p.Match(RDLStatementParserDEFAULT)
	}
	{
		p.SetState(101)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(102)
		p.Match(RDLStatementParserALGORITHM)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(103)
			p.IfNotExists()
		}

	}
	{
		p.SetState(106)
		p.AlgorithmDefinition()
	}

	return localctx
}

// IDropDefaultShadowAlgorithmContext is an interface to support dynamic dispatch.
type IDropDefaultShadowAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropDefaultShadowAlgorithmContext differentiates from other interfaces.
	IsDropDefaultShadowAlgorithmContext()
}

type DropDefaultShadowAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropDefaultShadowAlgorithmContext() *DropDefaultShadowAlgorithmContext {
	var p = new(DropDefaultShadowAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropDefaultShadowAlgorithm
	return p
}

func (*DropDefaultShadowAlgorithmContext) IsDropDefaultShadowAlgorithmContext() {}

func NewDropDefaultShadowAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropDefaultShadowAlgorithmContext {
	var p = new(DropDefaultShadowAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropDefaultShadowAlgorithm

	return p
}

func (s *DropDefaultShadowAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *DropDefaultShadowAlgorithmContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropDefaultShadowAlgorithmContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDEFAULT, 0)
}

func (s *DropDefaultShadowAlgorithmContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *DropDefaultShadowAlgorithmContext) ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALGORITHM, 0)
}

func (s *DropDefaultShadowAlgorithmContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropDefaultShadowAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropDefaultShadowAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropDefaultShadowAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropDefaultShadowAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropDefaultShadowAlgorithm() (localctx IDropDefaultShadowAlgorithmContext) {
	localctx = NewDropDefaultShadowAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RDLStatementParserRULE_dropDefaultShadowAlgorithm)
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
		p.SetState(108)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(109)
		p.Match(RDLStatementParserDEFAULT)
	}
	{
		p.SetState(110)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(111)
		p.Match(RDLStatementParserALGORITHM)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(112)
			p.IfExists()
		}

	}

	return localctx
}

// IAlterDefaultShadowAlgorithmContext is an interface to support dynamic dispatch.
type IAlterDefaultShadowAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterDefaultShadowAlgorithmContext differentiates from other interfaces.
	IsAlterDefaultShadowAlgorithmContext()
}

type AlterDefaultShadowAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterDefaultShadowAlgorithmContext() *AlterDefaultShadowAlgorithmContext {
	var p = new(AlterDefaultShadowAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterDefaultShadowAlgorithm
	return p
}

func (*AlterDefaultShadowAlgorithmContext) IsAlterDefaultShadowAlgorithmContext() {}

func NewAlterDefaultShadowAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterDefaultShadowAlgorithmContext {
	var p = new(AlterDefaultShadowAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterDefaultShadowAlgorithm

	return p
}

func (s *AlterDefaultShadowAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterDefaultShadowAlgorithmContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterDefaultShadowAlgorithmContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDEFAULT, 0)
}

func (s *AlterDefaultShadowAlgorithmContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *AlterDefaultShadowAlgorithmContext) ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALGORITHM, 0)
}

func (s *AlterDefaultShadowAlgorithmContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *AlterDefaultShadowAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterDefaultShadowAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterDefaultShadowAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterDefaultShadowAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterDefaultShadowAlgorithm() (localctx IAlterDefaultShadowAlgorithmContext) {
	localctx = NewAlterDefaultShadowAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RDLStatementParserRULE_alterDefaultShadowAlgorithm)

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
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(116)
		p.Match(RDLStatementParserDEFAULT)
	}
	{
		p.SetState(117)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(118)
		p.Match(RDLStatementParserALGORITHM)
	}
	{
		p.SetState(119)
		p.AlgorithmDefinition()
	}

	return localctx
}

// IShadowRuleDefinitionContext is an interface to support dynamic dispatch.
type IShadowRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShadowRuleDefinitionContext differentiates from other interfaces.
	IsShadowRuleDefinitionContext()
}

type ShadowRuleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShadowRuleDefinitionContext() *ShadowRuleDefinitionContext {
	var p = new(ShadowRuleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shadowRuleDefinition
	return p
}

func (*ShadowRuleDefinitionContext) IsShadowRuleDefinitionContext() {}

func NewShadowRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShadowRuleDefinitionContext {
	var p = new(ShadowRuleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shadowRuleDefinition

	return p
}

func (s *ShadowRuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShadowRuleDefinitionContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *ShadowRuleDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ShadowRuleDefinitionContext) SOURCE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSOURCE, 0)
}

func (s *ShadowRuleDefinitionContext) AllEQ_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserEQ_)
}

func (s *ShadowRuleDefinitionContext) EQ_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, i)
}

func (s *ShadowRuleDefinitionContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *ShadowRuleDefinitionContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ShadowRuleDefinitionContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ShadowRuleDefinitionContext) SHADOW() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHADOW, 0)
}

func (s *ShadowRuleDefinitionContext) Shadow() IShadowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShadowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShadowContext)
}

func (s *ShadowRuleDefinitionContext) AllShadowTableRule() []IShadowTableRuleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShadowTableRuleContext)(nil)).Elem())
	var tst = make([]IShadowTableRuleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShadowTableRuleContext)
		}
	}

	return tst
}

func (s *ShadowRuleDefinitionContext) ShadowTableRule(i int) IShadowTableRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShadowTableRuleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShadowTableRuleContext)
}

func (s *ShadowRuleDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ShadowRuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShadowRuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShadowRuleDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShadowRuleDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShadowRuleDefinition() (localctx IShadowRuleDefinitionContext) {
	localctx = NewShadowRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RDLStatementParserRULE_shadowRuleDefinition)
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
		p.RuleName()
	}
	{
		p.SetState(122)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(123)
		p.Match(RDLStatementParserSOURCE)
	}
	{
		p.SetState(124)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(125)
		p.Source()
	}
	{
		p.SetState(126)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(127)
		p.Match(RDLStatementParserSHADOW)
	}
	{
		p.SetState(128)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(129)
		p.Shadow()
	}
	{
		p.SetState(130)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(131)
		p.ShadowTableRule()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(132)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(133)
			p.ShadowTableRule()
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(139)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IShadowTableRuleContext is an interface to support dynamic dispatch.
type IShadowTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShadowTableRuleContext differentiates from other interfaces.
	IsShadowTableRuleContext()
}

type ShadowTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShadowTableRuleContext() *ShadowTableRuleContext {
	var p = new(ShadowTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shadowTableRule
	return p
}

func (*ShadowTableRuleContext) IsShadowTableRuleContext() {}

func NewShadowTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShadowTableRuleContext {
	var p = new(ShadowTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shadowTableRule

	return p
}

func (s *ShadowTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ShadowTableRuleContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *ShadowTableRuleContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ShadowTableRuleContext) AllAlgorithmDefinition() []IAlgorithmDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem())
	var tst = make([]IAlgorithmDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAlgorithmDefinitionContext)
		}
	}

	return tst
}

func (s *ShadowTableRuleContext) AlgorithmDefinition(i int) IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *ShadowTableRuleContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ShadowTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ShadowTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ShadowTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShadowTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShadowTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShadowTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShadowTableRule() (localctx IShadowTableRuleContext) {
	localctx = NewShadowTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RDLStatementParserRULE_shadowTableRule)
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
		p.SetState(141)
		p.TableName()
	}
	{
		p.SetState(142)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(143)
		p.AlgorithmDefinition()
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(144)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(145)
			p.AlgorithmDefinition()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_source
	return p
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitSource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RDLStatementParserRULE_source)

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
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IShadowContext is an interface to support dynamic dispatch.
type IShadowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShadowContext differentiates from other interfaces.
	IsShadowContext()
}

type ShadowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShadowContext() *ShadowContext {
	var p = new(ShadowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shadow
	return p
}

func (*ShadowContext) IsShadowContext() {}

func NewShadowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShadowContext {
	var p = new(ShadowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shadow

	return p
}

func (s *ShadowContext) GetParser() antlr.Parser { return s.parser }

func (s *ShadowContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *ShadowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShadowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShadowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShadow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) Shadow() (localctx IShadowContext) {
	localctx = NewShadowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RDLStatementParserRULE_shadow)

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
		p.SetState(155)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitTableName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RDLStatementParserRULE_tableName)

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

// IAlgorithmNameContext is an interface to support dynamic dispatch.
type IAlgorithmNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlgorithmNameContext differentiates from other interfaces.
	IsAlgorithmNameContext()
}

type AlgorithmNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmNameContext() *AlgorithmNameContext {
	var p = new(AlgorithmNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_algorithmName
	return p
}

func (*AlgorithmNameContext) IsAlgorithmNameContext() {}

func NewAlgorithmNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmNameContext {
	var p = new(AlgorithmNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_algorithmName

	return p
}

func (s *AlgorithmNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *AlgorithmNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlgorithmName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlgorithmName() (localctx IAlgorithmNameContext) {
	localctx = NewAlgorithmNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RDLStatementParserRULE_algorithmName)

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
		p.SetState(159)
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
	p.EnterRule(localctx, 26, RDLStatementParserRULE_ifExists)

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
		p.SetState(161)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(162)
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
	p.EnterRule(localctx, 28, RDLStatementParserRULE_ifNotExists)

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
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(165)
		p.Match(RDLStatementParserNOT)
	}
	{
		p.SetState(166)
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
	p.EnterRule(localctx, 30, RDLStatementParserRULE_literal)
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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserMINUS_, RDLStatementParserINT_:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RDLStatementParserMINUS_ {
			{
				p.SetState(169)
				p.Match(RDLStatementParserMINUS_)
			}

		}
		{
			p.SetState(172)
			p.Match(RDLStatementParserINT_)
		}

	case RDLStatementParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.Match(RDLStatementParserTRUE)
		}

	case RDLStatementParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
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
	p.EnterRule(localctx, 32, RDLStatementParserRULE_algorithmDefinition)
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
		p.SetState(177)
		p.Match(RDLStatementParserTYPE)
	}
	{
		p.SetState(178)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(179)
		p.Match(RDLStatementParserNAME)
	}
	{
		p.SetState(180)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(181)
		p.AlgorithmTypeName()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(182)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(183)
			p.PropertiesDefinition()
		}

	}
	{
		p.SetState(186)
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

func (s *AlgorithmTypeNameContext) BuildInShadowAlgorithmType() IBuildInShadowAlgorithmTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildInShadowAlgorithmTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildInShadowAlgorithmTypeContext)
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
	p.EnterRule(localctx, 34, RDLStatementParserRULE_algorithmTypeName)

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

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserVALUE_MATCH, RDLStatementParserREGEX_MATCH, RDLStatementParserSQL_HINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.BuildInShadowAlgorithmType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuildInShadowAlgorithmTypeContext is an interface to support dynamic dispatch.
type IBuildInShadowAlgorithmTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildInShadowAlgorithmTypeContext differentiates from other interfaces.
	IsBuildInShadowAlgorithmTypeContext()
}

type BuildInShadowAlgorithmTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildInShadowAlgorithmTypeContext() *BuildInShadowAlgorithmTypeContext {
	var p = new(BuildInShadowAlgorithmTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildInShadowAlgorithmType
	return p
}

func (*BuildInShadowAlgorithmTypeContext) IsBuildInShadowAlgorithmTypeContext() {}

func NewBuildInShadowAlgorithmTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildInShadowAlgorithmTypeContext {
	var p = new(BuildInShadowAlgorithmTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildInShadowAlgorithmType

	return p
}

func (s *BuildInShadowAlgorithmTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildInShadowAlgorithmTypeContext) VALUE_MATCH() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserVALUE_MATCH, 0)
}

func (s *BuildInShadowAlgorithmTypeContext) REGEX_MATCH() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREGEX_MATCH, 0)
}

func (s *BuildInShadowAlgorithmTypeContext) SQL_HINT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSQL_HINT, 0)
}

func (s *BuildInShadowAlgorithmTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInShadowAlgorithmTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildInShadowAlgorithmTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildInShadowAlgorithmType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildInShadowAlgorithmType() (localctx IBuildInShadowAlgorithmTypeContext) {
	localctx = NewBuildInShadowAlgorithmTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RDLStatementParserRULE_buildInShadowAlgorithmType)
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
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(RDLStatementParserVALUE_MATCH-72))|(1<<(RDLStatementParserREGEX_MATCH-72))|(1<<(RDLStatementParserSQL_HINT-72)))) != 0) {
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
	p.EnterRule(localctx, 38, RDLStatementParserRULE_propertiesDefinition)
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
		p.SetState(194)
		p.Match(RDLStatementParserPROPERTIES)
	}
	{
		p.SetState(195)
		p.Match(RDLStatementParserLP_)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserSTRING_ {
		{
			p.SetState(196)
			p.Properties()
		}

	}
	{
		p.SetState(199)
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
	p.EnterRule(localctx, 40, RDLStatementParserRULE_properties)
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
		p.SetState(201)
		p.Property()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(202)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(203)
			p.Property()
		}

		p.SetState(208)
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
	p.EnterRule(localctx, 42, RDLStatementParserRULE_property)

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
		p.SetState(209)

		var _m = p.Match(RDLStatementParserSTRING_)

		localctx.(*PropertyContext).key = _m
	}
	{
		p.SetState(210)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(211)

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
	p.EnterRule(localctx, 44, RDLStatementParserRULE_ruleName)

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
		p.SetState(213)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}
