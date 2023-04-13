// Code generated from encrypt_g4/RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 93, 300,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	3, 2, 3, 2, 3, 2, 5, 2, 71, 10, 2, 3, 2, 3, 2, 3, 2, 7, 2, 76, 10, 2, 12,
	2, 14, 2, 79, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 87, 10,
	3, 12, 3, 14, 3, 90, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 96, 10, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 101, 10, 4, 12, 4, 14, 4, 104, 11, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 111, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	118, 10, 5, 12, 5, 14, 5, 121, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5,
	5, 128, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 142, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 148, 10,
	8, 3, 8, 3, 8, 5, 8, 152, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 158, 10,
	8, 3, 8, 3, 8, 5, 8, 162, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 168, 10,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 179, 10,
	9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 5, 12, 192, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 203, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 214, 10, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 225, 10, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 5, 26, 255, 10, 26, 3, 26, 3, 26, 3,
	26, 5, 26, 260, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	5, 27, 269, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 5, 28, 275, 10, 28, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 5, 30, 282, 10, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 7, 31, 289, 10, 31, 12, 31, 14, 31, 292, 11, 31, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 2, 2, 34, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 2, 4, 3, 2, 69, 70, 3, 2, 79, 84, 2, 294, 2,
	66, 3, 2, 2, 2, 4, 80, 3, 2, 2, 2, 6, 91, 3, 2, 2, 2, 8, 105, 3, 2, 2,
	2, 10, 131, 3, 2, 2, 2, 12, 135, 3, 2, 2, 2, 14, 137, 3, 2, 2, 2, 16, 171,
	3, 2, 2, 2, 18, 180, 3, 2, 2, 2, 20, 182, 3, 2, 2, 2, 22, 184, 3, 2, 2,
	2, 24, 193, 3, 2, 2, 2, 26, 195, 3, 2, 2, 2, 28, 204, 3, 2, 2, 2, 30, 206,
	3, 2, 2, 2, 32, 215, 3, 2, 2, 2, 34, 217, 3, 2, 2, 2, 36, 226, 3, 2, 2,
	2, 38, 228, 3, 2, 2, 2, 40, 233, 3, 2, 2, 2, 42, 238, 3, 2, 2, 2, 44, 243,
	3, 2, 2, 2, 46, 245, 3, 2, 2, 2, 48, 248, 3, 2, 2, 2, 50, 259, 3, 2, 2,
	2, 52, 261, 3, 2, 2, 2, 54, 274, 3, 2, 2, 2, 56, 276, 3, 2, 2, 2, 58, 278,
	3, 2, 2, 2, 60, 285, 3, 2, 2, 2, 62, 293, 3, 2, 2, 2, 64, 297, 3, 2, 2,
	2, 66, 67, 7, 46, 2, 2, 67, 68, 7, 53, 2, 2, 68, 70, 7, 51, 2, 2, 69, 71,
	5, 48, 25, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2,
	2, 72, 77, 5, 8, 5, 2, 73, 74, 7, 36, 2, 2, 74, 76, 5, 8, 5, 2, 75, 73,
	3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2,
	78, 3, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 7, 47, 2, 2, 81, 82, 7,
	53, 2, 2, 82, 83, 7, 51, 2, 2, 83, 88, 5, 8, 5, 2, 84, 85, 7, 36, 2, 2,
	85, 87, 5, 8, 5, 2, 86, 84, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3,
	2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 5, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91,
	92, 7, 48, 2, 2, 92, 93, 7, 53, 2, 2, 93, 95, 7, 51, 2, 2, 94, 96, 5, 46,
	24, 2, 95, 94, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97,
	102, 5, 64, 33, 2, 98, 99, 7, 36, 2, 2, 99, 101, 5, 64, 33, 2, 100, 98,
	3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 7, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 106, 5, 64, 33, 2,
	106, 110, 7, 30, 2, 2, 107, 108, 5, 10, 6, 2, 108, 109, 7, 36, 2, 2, 109,
	111, 3, 2, 2, 2, 110, 107, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 113, 7, 63, 2, 2, 113, 114, 7, 30, 2, 2, 114, 119, 5,
	14, 8, 2, 115, 116, 7, 36, 2, 2, 116, 118, 5, 14, 8, 2, 117, 115, 3, 2,
	2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2,
	120, 122, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 127, 7, 31, 2, 2, 123,
	124, 7, 36, 2, 2, 124, 125, 7, 68, 2, 2, 125, 126, 7, 23, 2, 2, 126, 128,
	5, 44, 23, 2, 127, 123, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3,
	2, 2, 2, 129, 130, 7, 31, 2, 2, 130, 9, 3, 2, 2, 2, 131, 132, 7, 50, 2,
	2, 132, 133, 7, 23, 2, 2, 133, 134, 5, 12, 7, 2, 134, 11, 3, 2, 2, 2, 135,
	136, 7, 87, 2, 2, 136, 13, 3, 2, 2, 2, 137, 138, 7, 30, 2, 2, 138, 141,
	5, 16, 9, 2, 139, 140, 7, 36, 2, 2, 140, 142, 5, 22, 12, 2, 141, 139, 3,
	2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 7, 36, 2,
	2, 144, 147, 5, 26, 14, 2, 145, 146, 7, 36, 2, 2, 146, 148, 5, 30, 16,
	2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149,
	150, 7, 36, 2, 2, 150, 152, 5, 34, 18, 2, 151, 149, 3, 2, 2, 2, 151, 152,
	3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 7, 36, 2, 2, 154, 157, 5, 38,
	20, 2, 155, 156, 7, 36, 2, 2, 156, 158, 5, 40, 21, 2, 157, 155, 3, 2, 2,
	2, 157, 158, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 160, 7, 36, 2, 2, 160,
	162, 5, 42, 22, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 167,
	3, 2, 2, 2, 163, 164, 7, 36, 2, 2, 164, 165, 7, 68, 2, 2, 165, 166, 7,
	23, 2, 2, 166, 168, 5, 44, 23, 2, 167, 163, 3, 2, 2, 2, 167, 168, 3, 2,
	2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 7, 31, 2, 2, 170, 15, 3, 2, 2, 2,
	171, 172, 7, 58, 2, 2, 172, 173, 7, 23, 2, 2, 173, 178, 5, 18, 10, 2, 174,
	175, 7, 36, 2, 2, 175, 176, 7, 71, 2, 2, 176, 177, 7, 23, 2, 2, 177, 179,
	5, 20, 11, 2, 178, 174, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 17, 3, 2,
	2, 2, 180, 181, 7, 87, 2, 2, 181, 19, 3, 2, 2, 2, 182, 183, 7, 88, 2, 2,
	183, 21, 3, 2, 2, 2, 184, 185, 7, 65, 2, 2, 185, 186, 7, 23, 2, 2, 186,
	191, 5, 24, 13, 2, 187, 188, 7, 36, 2, 2, 188, 189, 7, 72, 2, 2, 189, 190,
	7, 23, 2, 2, 190, 192, 5, 20, 11, 2, 191, 187, 3, 2, 2, 2, 191, 192, 3,
	2, 2, 2, 192, 23, 3, 2, 2, 2, 193, 194, 7, 87, 2, 2, 194, 25, 3, 2, 2,
	2, 195, 196, 7, 64, 2, 2, 196, 197, 7, 23, 2, 2, 197, 202, 5, 28, 15, 2,
	198, 199, 7, 36, 2, 2, 199, 200, 7, 73, 2, 2, 200, 201, 7, 23, 2, 2, 201,
	203, 5, 20, 11, 2, 202, 198, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 27,
	3, 2, 2, 2, 204, 205, 7, 87, 2, 2, 205, 29, 3, 2, 2, 2, 206, 207, 7, 66,
	2, 2, 207, 208, 7, 23, 2, 2, 208, 213, 5, 32, 17, 2, 209, 210, 7, 36, 2,
	2, 210, 211, 7, 74, 2, 2, 211, 212, 7, 23, 2, 2, 212, 214, 5, 20, 11, 2,
	213, 209, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 31, 3, 2, 2, 2, 215, 216,
	7, 87, 2, 2, 216, 33, 3, 2, 2, 2, 217, 218, 7, 67, 2, 2, 218, 219, 7, 23,
	2, 2, 219, 224, 5, 36, 19, 2, 220, 221, 7, 36, 2, 2, 221, 222, 7, 75, 2,
	2, 222, 223, 7, 23, 2, 2, 223, 225, 5, 20, 11, 2, 224, 220, 3, 2, 2, 2,
	224, 225, 3, 2, 2, 2, 225, 35, 3, 2, 2, 2, 226, 227, 7, 87, 2, 2, 227,
	37, 3, 2, 2, 2, 228, 229, 7, 55, 2, 2, 229, 230, 7, 30, 2, 2, 230, 231,
	5, 52, 27, 2, 231, 232, 7, 31, 2, 2, 232, 39, 3, 2, 2, 2, 233, 234, 7,
	56, 2, 2, 234, 235, 7, 30, 2, 2, 235, 236, 5, 52, 27, 2, 236, 237, 7, 31,
	2, 2, 237, 41, 3, 2, 2, 2, 238, 239, 7, 57, 2, 2, 239, 240, 7, 30, 2, 2,
	240, 241, 5, 52, 27, 2, 241, 242, 7, 31, 2, 2, 242, 43, 3, 2, 2, 2, 243,
	244, 9, 2, 2, 2, 244, 45, 3, 2, 2, 2, 245, 246, 7, 76, 2, 2, 246, 247,
	7, 77, 2, 2, 247, 47, 3, 2, 2, 2, 248, 249, 7, 76, 2, 2, 249, 250, 7, 85,
	2, 2, 250, 251, 7, 77, 2, 2, 251, 49, 3, 2, 2, 2, 252, 260, 7, 88, 2, 2,
	253, 255, 7, 15, 2, 2, 254, 253, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255,
	256, 3, 2, 2, 2, 256, 260, 7, 89, 2, 2, 257, 260, 7, 69, 2, 2, 258, 260,
	7, 70, 2, 2, 259, 252, 3, 2, 2, 2, 259, 254, 3, 2, 2, 2, 259, 257, 3, 2,
	2, 2, 259, 258, 3, 2, 2, 2, 260, 51, 3, 2, 2, 2, 261, 262, 7, 54, 2, 2,
	262, 263, 7, 30, 2, 2, 263, 264, 7, 58, 2, 2, 264, 265, 7, 23, 2, 2, 265,
	268, 5, 54, 28, 2, 266, 267, 7, 36, 2, 2, 267, 269, 5, 58, 30, 2, 268,
	266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 271,
	7, 31, 2, 2, 271, 53, 3, 2, 2, 2, 272, 275, 5, 56, 29, 2, 273, 275, 7,
	88, 2, 2, 274, 272, 3, 2, 2, 2, 274, 273, 3, 2, 2, 2, 275, 55, 3, 2, 2,
	2, 276, 277, 9, 3, 2, 2, 277, 57, 3, 2, 2, 2, 278, 279, 7, 59, 2, 2, 279,
	281, 7, 30, 2, 2, 280, 282, 5, 60, 31, 2, 281, 280, 3, 2, 2, 2, 281, 282,
	3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 284, 7, 31, 2, 2, 284, 59, 3, 2,
	2, 2, 285, 290, 5, 62, 32, 2, 286, 287, 7, 36, 2, 2, 287, 289, 5, 62, 32,
	2, 288, 286, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290,
	291, 3, 2, 2, 2, 291, 61, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 293, 294, 7,
	88, 2, 2, 294, 295, 7, 23, 2, 2, 295, 296, 5, 50, 26, 2, 296, 63, 3, 2,
	2, 2, 297, 298, 7, 87, 2, 2, 298, 65, 3, 2, 2, 2, 27, 70, 77, 88, 95, 102,
	110, 119, 127, 141, 147, 151, 157, 161, 167, 178, 191, 202, 213, 224, 254,
	259, 268, 274, 281, 290,
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
	"", "", "", "", "", "", "", "", "", "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'",
}
var symbolicNames = []string{
	"", "AND_", "OR_", "NOT_", "TILDE_", "VERTICALBAR_", "AMPERSAND_", "SIGNEDLEFTSHIFT_",
	"SIGNEDRIGHTSHIFT_", "CARET_", "MOD_", "COLON_", "PLUS_", "MINUS_", "ASTERISK_",
	"SLASH_", "BACKSLASH_", "DOT_", "DOTASTERISK_", "SAFEEQ_", "DEQ_", "EQ_",
	"NEQ_", "GT_", "GTE_", "LT_", "LTE_", "POUND_", "LP_", "RP_", "LBE_", "RBE_",
	"LBT_", "RBT_", "COMMA_", "DQ_", "SQ_", "BQ_", "QUESTION_", "AT_", "SEMI_",
	"JSONSEPARATOR_", "UL_", "WS", "CREATE", "ALTER", "DROP", "SHOW", "RESOURCE",
	"RULE", "FROM", "ENCRYPT", "TYPE", "ENCRYPT_ALGORITHM", "ASSISTED_QUERY_ALGORITHM",
	"LIKE_QUERY_ALGORITHM", "NAME", "PROPERTIES", "COLUMN", "RULES", "TABLE",
	"COLUMNS", "CIPHER", "PLAIN", "ASSISTED_QUERY_COLUMN", "LIKE_QUERY_COLUMN",
	"QUERY_WITH_CIPHER_COLUMN", "TRUE", "FALSE", "DATA_TYPE", "PLAIN_DATA_TYPE",
	"CIPHER_DATA_TYPE", "ASSISTED_QUERY_DATA_TYPE", "LIKE_QUERY_DATA_TYPE",
	"IF", "EXISTS", "COUNT", "MD5", "AES", "RC4", "SM3", "SM4", "CHAR_DIGEST_LIKE",
	"NOT", "FOR_GENERATOR", "IDENTIFIER_", "STRING_", "INT_", "HEX_", "NUMBER_",
	"HEXDIGIT_", "BITNUM_",
}

var ruleNames = []string{
	"createEncryptRule", "alterEncryptRule", "dropEncryptRule", "encryptRuleDefinition",
	"resourceDefinition", "resourceName", "encryptColumnDefinition", "columnDefinition",
	"columnName", "dataType", "plainColumnDefinition", "plainColumnName", "cipherColumnDefinition",
	"cipherColumnName", "assistedQueryColumnDefinition", "assistedQueryColumnName",
	"likeQueryColumnDefinition", "likeQueryColumnName", "encryptAlgorithm",
	"assistedQueryAlgorithm", "likeQueryAlgorithm", "queryWithCipherColumn",
	"ifExists", "ifNotExists", "literal", "algorithmDefinition", "algorithmTypeName",
	"buildinAlgorithmTypeName", "propertiesDefinition", "properties", "property",
	"tableName",
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
	RDLStatementParserEOF                      = antlr.TokenEOF
	RDLStatementParserAND_                     = 1
	RDLStatementParserOR_                      = 2
	RDLStatementParserNOT_                     = 3
	RDLStatementParserTILDE_                   = 4
	RDLStatementParserVERTICALBAR_             = 5
	RDLStatementParserAMPERSAND_               = 6
	RDLStatementParserSIGNEDLEFTSHIFT_         = 7
	RDLStatementParserSIGNEDRIGHTSHIFT_        = 8
	RDLStatementParserCARET_                   = 9
	RDLStatementParserMOD_                     = 10
	RDLStatementParserCOLON_                   = 11
	RDLStatementParserPLUS_                    = 12
	RDLStatementParserMINUS_                   = 13
	RDLStatementParserASTERISK_                = 14
	RDLStatementParserSLASH_                   = 15
	RDLStatementParserBACKSLASH_               = 16
	RDLStatementParserDOT_                     = 17
	RDLStatementParserDOTASTERISK_             = 18
	RDLStatementParserSAFEEQ_                  = 19
	RDLStatementParserDEQ_                     = 20
	RDLStatementParserEQ_                      = 21
	RDLStatementParserNEQ_                     = 22
	RDLStatementParserGT_                      = 23
	RDLStatementParserGTE_                     = 24
	RDLStatementParserLT_                      = 25
	RDLStatementParserLTE_                     = 26
	RDLStatementParserPOUND_                   = 27
	RDLStatementParserLP_                      = 28
	RDLStatementParserRP_                      = 29
	RDLStatementParserLBE_                     = 30
	RDLStatementParserRBE_                     = 31
	RDLStatementParserLBT_                     = 32
	RDLStatementParserRBT_                     = 33
	RDLStatementParserCOMMA_                   = 34
	RDLStatementParserDQ_                      = 35
	RDLStatementParserSQ_                      = 36
	RDLStatementParserBQ_                      = 37
	RDLStatementParserQUESTION_                = 38
	RDLStatementParserAT_                      = 39
	RDLStatementParserSEMI_                    = 40
	RDLStatementParserJSONSEPARATOR_           = 41
	RDLStatementParserUL_                      = 42
	RDLStatementParserWS                       = 43
	RDLStatementParserCREATE                   = 44
	RDLStatementParserALTER                    = 45
	RDLStatementParserDROP                     = 46
	RDLStatementParserSHOW                     = 47
	RDLStatementParserRESOURCE                 = 48
	RDLStatementParserRULE                     = 49
	RDLStatementParserFROM                     = 50
	RDLStatementParserENCRYPT                  = 51
	RDLStatementParserTYPE                     = 52
	RDLStatementParserENCRYPT_ALGORITHM        = 53
	RDLStatementParserASSISTED_QUERY_ALGORITHM = 54
	RDLStatementParserLIKE_QUERY_ALGORITHM     = 55
	RDLStatementParserNAME                     = 56
	RDLStatementParserPROPERTIES               = 57
	RDLStatementParserCOLUMN                   = 58
	RDLStatementParserRULES                    = 59
	RDLStatementParserTABLE                    = 60
	RDLStatementParserCOLUMNS                  = 61
	RDLStatementParserCIPHER                   = 62
	RDLStatementParserPLAIN                    = 63
	RDLStatementParserASSISTED_QUERY_COLUMN    = 64
	RDLStatementParserLIKE_QUERY_COLUMN        = 65
	RDLStatementParserQUERY_WITH_CIPHER_COLUMN = 66
	RDLStatementParserTRUE                     = 67
	RDLStatementParserFALSE                    = 68
	RDLStatementParserDATA_TYPE                = 69
	RDLStatementParserPLAIN_DATA_TYPE          = 70
	RDLStatementParserCIPHER_DATA_TYPE         = 71
	RDLStatementParserASSISTED_QUERY_DATA_TYPE = 72
	RDLStatementParserLIKE_QUERY_DATA_TYPE     = 73
	RDLStatementParserIF                       = 74
	RDLStatementParserEXISTS                   = 75
	RDLStatementParserCOUNT                    = 76
	RDLStatementParserMD5                      = 77
	RDLStatementParserAES                      = 78
	RDLStatementParserRC4                      = 79
	RDLStatementParserSM3                      = 80
	RDLStatementParserSM4                      = 81
	RDLStatementParserCHAR_DIGEST_LIKE         = 82
	RDLStatementParserNOT                      = 83
	RDLStatementParserFOR_GENERATOR            = 84
	RDLStatementParserIDENTIFIER_              = 85
	RDLStatementParserSTRING_                  = 86
	RDLStatementParserINT_                     = 87
	RDLStatementParserHEX_                     = 88
	RDLStatementParserNUMBER_                  = 89
	RDLStatementParserHEXDIGIT_                = 90
	RDLStatementParserBITNUM_                  = 91
)

// RDLStatementParser rules.
const (
	RDLStatementParserRULE_createEncryptRule             = 0
	RDLStatementParserRULE_alterEncryptRule              = 1
	RDLStatementParserRULE_dropEncryptRule               = 2
	RDLStatementParserRULE_encryptRuleDefinition         = 3
	RDLStatementParserRULE_resourceDefinition            = 4
	RDLStatementParserRULE_resourceName                  = 5
	RDLStatementParserRULE_encryptColumnDefinition       = 6
	RDLStatementParserRULE_columnDefinition              = 7
	RDLStatementParserRULE_columnName                    = 8
	RDLStatementParserRULE_dataType                      = 9
	RDLStatementParserRULE_plainColumnDefinition         = 10
	RDLStatementParserRULE_plainColumnName               = 11
	RDLStatementParserRULE_cipherColumnDefinition        = 12
	RDLStatementParserRULE_cipherColumnName              = 13
	RDLStatementParserRULE_assistedQueryColumnDefinition = 14
	RDLStatementParserRULE_assistedQueryColumnName       = 15
	RDLStatementParserRULE_likeQueryColumnDefinition     = 16
	RDLStatementParserRULE_likeQueryColumnName           = 17
	RDLStatementParserRULE_encryptAlgorithm              = 18
	RDLStatementParserRULE_assistedQueryAlgorithm        = 19
	RDLStatementParserRULE_likeQueryAlgorithm            = 20
	RDLStatementParserRULE_queryWithCipherColumn         = 21
	RDLStatementParserRULE_ifExists                      = 22
	RDLStatementParserRULE_ifNotExists                   = 23
	RDLStatementParserRULE_literal                       = 24
	RDLStatementParserRULE_algorithmDefinition           = 25
	RDLStatementParserRULE_algorithmTypeName             = 26
	RDLStatementParserRULE_buildinAlgorithmTypeName      = 27
	RDLStatementParserRULE_propertiesDefinition          = 28
	RDLStatementParserRULE_properties                    = 29
	RDLStatementParserRULE_property                      = 30
	RDLStatementParserRULE_tableName                     = 31
)

// ICreateEncryptRuleContext is an interface to support dynamic dispatch.
type ICreateEncryptRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateEncryptRuleContext differentiates from other interfaces.
	IsCreateEncryptRuleContext()
}

type CreateEncryptRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateEncryptRuleContext() *CreateEncryptRuleContext {
	var p = new(CreateEncryptRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createEncryptRule
	return p
}

func (*CreateEncryptRuleContext) IsCreateEncryptRuleContext() {}

func NewCreateEncryptRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateEncryptRuleContext {
	var p = new(CreateEncryptRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createEncryptRule

	return p
}

func (s *CreateEncryptRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateEncryptRuleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateEncryptRuleContext) ENCRYPT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserENCRYPT, 0)
}

func (s *CreateEncryptRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *CreateEncryptRuleContext) AllEncryptRuleDefinition() []IEncryptRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEncryptRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IEncryptRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEncryptRuleDefinitionContext)
		}
	}

	return tst
}

func (s *CreateEncryptRuleContext) EncryptRuleDefinition(i int) IEncryptRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncryptRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEncryptRuleDefinitionContext)
}

func (s *CreateEncryptRuleContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateEncryptRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *CreateEncryptRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *CreateEncryptRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateEncryptRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateEncryptRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateEncryptRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateEncryptRule() (localctx ICreateEncryptRuleContext) {
	localctx = NewCreateEncryptRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RDLStatementParserRULE_createEncryptRule)
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
		p.SetState(64)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(65)
		p.Match(RDLStatementParserENCRYPT)
	}
	{
		p.SetState(66)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(67)
			p.IfNotExists()
		}

	}
	{
		p.SetState(70)
		p.EncryptRuleDefinition()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(71)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(72)
			p.EncryptRuleDefinition()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlterEncryptRuleContext is an interface to support dynamic dispatch.
type IAlterEncryptRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterEncryptRuleContext differentiates from other interfaces.
	IsAlterEncryptRuleContext()
}

type AlterEncryptRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterEncryptRuleContext() *AlterEncryptRuleContext {
	var p = new(AlterEncryptRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterEncryptRule
	return p
}

func (*AlterEncryptRuleContext) IsAlterEncryptRuleContext() {}

func NewAlterEncryptRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterEncryptRuleContext {
	var p = new(AlterEncryptRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterEncryptRule

	return p
}

func (s *AlterEncryptRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterEncryptRuleContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterEncryptRuleContext) ENCRYPT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserENCRYPT, 0)
}

func (s *AlterEncryptRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *AlterEncryptRuleContext) AllEncryptRuleDefinition() []IEncryptRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEncryptRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IEncryptRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEncryptRuleDefinitionContext)
		}
	}

	return tst
}

func (s *AlterEncryptRuleContext) EncryptRuleDefinition(i int) IEncryptRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncryptRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEncryptRuleDefinitionContext)
}

func (s *AlterEncryptRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *AlterEncryptRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *AlterEncryptRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterEncryptRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterEncryptRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterEncryptRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterEncryptRule() (localctx IAlterEncryptRuleContext) {
	localctx = NewAlterEncryptRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RDLStatementParserRULE_alterEncryptRule)
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
		p.SetState(78)
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(79)
		p.Match(RDLStatementParserENCRYPT)
	}
	{
		p.SetState(80)
		p.Match(RDLStatementParserRULE)
	}
	{
		p.SetState(81)
		p.EncryptRuleDefinition()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(82)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(83)
			p.EncryptRuleDefinition()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropEncryptRuleContext is an interface to support dynamic dispatch.
type IDropEncryptRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropEncryptRuleContext differentiates from other interfaces.
	IsDropEncryptRuleContext()
}

type DropEncryptRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropEncryptRuleContext() *DropEncryptRuleContext {
	var p = new(DropEncryptRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropEncryptRule
	return p
}

func (*DropEncryptRuleContext) IsDropEncryptRuleContext() {}

func NewDropEncryptRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropEncryptRuleContext {
	var p = new(DropEncryptRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropEncryptRule

	return p
}

func (s *DropEncryptRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropEncryptRuleContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropEncryptRuleContext) ENCRYPT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserENCRYPT, 0)
}

func (s *DropEncryptRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *DropEncryptRuleContext) AllTableName() []ITableNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableNameContext)(nil)).Elem())
	var tst = make([]ITableNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableNameContext)
		}
	}

	return tst
}

func (s *DropEncryptRuleContext) TableName(i int) ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DropEncryptRuleContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropEncryptRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropEncryptRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropEncryptRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropEncryptRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropEncryptRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropEncryptRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropEncryptRule() (localctx IDropEncryptRuleContext) {
	localctx = NewDropEncryptRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RDLStatementParserRULE_dropEncryptRule)
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
		p.SetState(89)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(90)
		p.Match(RDLStatementParserENCRYPT)
	}
	{
		p.SetState(91)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(92)
			p.IfExists()
		}

	}
	{
		p.SetState(95)
		p.TableName()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(96)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(97)
			p.TableName()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEncryptRuleDefinitionContext is an interface to support dynamic dispatch.
type IEncryptRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncryptRuleDefinitionContext differentiates from other interfaces.
	IsEncryptRuleDefinitionContext()
}

type EncryptRuleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncryptRuleDefinitionContext() *EncryptRuleDefinitionContext {
	var p = new(EncryptRuleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_encryptRuleDefinition
	return p
}

func (*EncryptRuleDefinitionContext) IsEncryptRuleDefinitionContext() {}

func NewEncryptRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncryptRuleDefinitionContext {
	var p = new(EncryptRuleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_encryptRuleDefinition

	return p
}

func (s *EncryptRuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EncryptRuleDefinitionContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *EncryptRuleDefinitionContext) AllLP_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserLP_)
}

func (s *EncryptRuleDefinitionContext) LP_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, i)
}

func (s *EncryptRuleDefinitionContext) COLUMNS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOLUMNS, 0)
}

func (s *EncryptRuleDefinitionContext) AllEncryptColumnDefinition() []IEncryptColumnDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEncryptColumnDefinitionContext)(nil)).Elem())
	var tst = make([]IEncryptColumnDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEncryptColumnDefinitionContext)
		}
	}

	return tst
}

func (s *EncryptRuleDefinitionContext) EncryptColumnDefinition(i int) IEncryptColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncryptColumnDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEncryptColumnDefinitionContext)
}

func (s *EncryptRuleDefinitionContext) AllRP_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserRP_)
}

func (s *EncryptRuleDefinitionContext) RP_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, i)
}

func (s *EncryptRuleDefinitionContext) ResourceDefinition() IResourceDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResourceDefinitionContext)
}

func (s *EncryptRuleDefinitionContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *EncryptRuleDefinitionContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *EncryptRuleDefinitionContext) QUERY_WITH_CIPHER_COLUMN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserQUERY_WITH_CIPHER_COLUMN, 0)
}

func (s *EncryptRuleDefinitionContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *EncryptRuleDefinitionContext) QueryWithCipherColumn() IQueryWithCipherColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryWithCipherColumnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryWithCipherColumnContext)
}

func (s *EncryptRuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncryptRuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncryptRuleDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitEncryptRuleDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) EncryptRuleDefinition() (localctx IEncryptRuleDefinitionContext) {
	localctx = NewEncryptRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RDLStatementParserRULE_encryptRuleDefinition)
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
		p.SetState(103)
		p.TableName()
	}
	{
		p.SetState(104)
		p.Match(RDLStatementParserLP_)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserRESOURCE {
		{
			p.SetState(105)
			p.ResourceDefinition()
		}
		{
			p.SetState(106)
			p.Match(RDLStatementParserCOMMA_)
		}

	}
	{
		p.SetState(110)
		p.Match(RDLStatementParserCOLUMNS)
	}
	{
		p.SetState(111)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(112)
		p.EncryptColumnDefinition()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(113)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(114)
			p.EncryptColumnDefinition()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(120)
		p.Match(RDLStatementParserRP_)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(121)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(122)
			p.Match(RDLStatementParserQUERY_WITH_CIPHER_COLUMN)
		}
		{
			p.SetState(123)
			p.Match(RDLStatementParserEQ_)
		}
		{
			p.SetState(124)
			p.QueryWithCipherColumn()
		}

	}
	{
		p.SetState(127)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IResourceDefinitionContext is an interface to support dynamic dispatch.
type IResourceDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResourceDefinitionContext differentiates from other interfaces.
	IsResourceDefinitionContext()
}

type ResourceDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceDefinitionContext() *ResourceDefinitionContext {
	var p = new(ResourceDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_resourceDefinition
	return p
}

func (*ResourceDefinitionContext) IsResourceDefinitionContext() {}

func NewResourceDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceDefinitionContext {
	var p = new(ResourceDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_resourceDefinition

	return p
}

func (s *ResourceDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceDefinitionContext) RESOURCE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRESOURCE, 0)
}

func (s *ResourceDefinitionContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *ResourceDefinitionContext) ResourceName() IResourceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResourceNameContext)
}

func (s *ResourceDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitResourceDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ResourceDefinition() (localctx IResourceDefinitionContext) {
	localctx = NewResourceDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RDLStatementParserRULE_resourceDefinition)

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
		p.Match(RDLStatementParserRESOURCE)
	}
	{
		p.SetState(130)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(131)
		p.ResourceName()
	}

	return localctx
}

// IResourceNameContext is an interface to support dynamic dispatch.
type IResourceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResourceNameContext differentiates from other interfaces.
	IsResourceNameContext()
}

type ResourceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceNameContext() *ResourceNameContext {
	var p = new(ResourceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_resourceName
	return p
}

func (*ResourceNameContext) IsResourceNameContext() {}

func NewResourceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceNameContext {
	var p = new(ResourceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_resourceName

	return p
}

func (s *ResourceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *ResourceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitResourceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ResourceName() (localctx IResourceNameContext) {
	localctx = NewResourceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RDLStatementParserRULE_resourceName)

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
		p.SetState(133)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IEncryptColumnDefinitionContext is an interface to support dynamic dispatch.
type IEncryptColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncryptColumnDefinitionContext differentiates from other interfaces.
	IsEncryptColumnDefinitionContext()
}

type EncryptColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncryptColumnDefinitionContext() *EncryptColumnDefinitionContext {
	var p = new(EncryptColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_encryptColumnDefinition
	return p
}

func (*EncryptColumnDefinitionContext) IsEncryptColumnDefinitionContext() {}

func NewEncryptColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncryptColumnDefinitionContext {
	var p = new(EncryptColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_encryptColumnDefinition

	return p
}

func (s *EncryptColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EncryptColumnDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *EncryptColumnDefinitionContext) ColumnDefinition() IColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *EncryptColumnDefinitionContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *EncryptColumnDefinitionContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *EncryptColumnDefinitionContext) CipherColumnDefinition() ICipherColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICipherColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICipherColumnDefinitionContext)
}

func (s *EncryptColumnDefinitionContext) EncryptAlgorithm() IEncryptAlgorithmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncryptAlgorithmContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEncryptAlgorithmContext)
}

func (s *EncryptColumnDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *EncryptColumnDefinitionContext) PlainColumnDefinition() IPlainColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlainColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlainColumnDefinitionContext)
}

func (s *EncryptColumnDefinitionContext) AssistedQueryColumnDefinition() IAssistedQueryColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssistedQueryColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssistedQueryColumnDefinitionContext)
}

func (s *EncryptColumnDefinitionContext) LikeQueryColumnDefinition() ILikeQueryColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeQueryColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeQueryColumnDefinitionContext)
}

func (s *EncryptColumnDefinitionContext) AssistedQueryAlgorithm() IAssistedQueryAlgorithmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssistedQueryAlgorithmContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssistedQueryAlgorithmContext)
}

func (s *EncryptColumnDefinitionContext) LikeQueryAlgorithm() ILikeQueryAlgorithmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeQueryAlgorithmContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeQueryAlgorithmContext)
}

func (s *EncryptColumnDefinitionContext) QUERY_WITH_CIPHER_COLUMN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserQUERY_WITH_CIPHER_COLUMN, 0)
}

func (s *EncryptColumnDefinitionContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *EncryptColumnDefinitionContext) QueryWithCipherColumn() IQueryWithCipherColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryWithCipherColumnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryWithCipherColumnContext)
}

func (s *EncryptColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncryptColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncryptColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitEncryptColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) EncryptColumnDefinition() (localctx IEncryptColumnDefinitionContext) {
	localctx = NewEncryptColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RDLStatementParserRULE_encryptColumnDefinition)
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
		p.SetState(135)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(136)
		p.ColumnDefinition()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(137)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(138)
			p.PlainColumnDefinition()
		}

	}
	{
		p.SetState(141)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(142)
		p.CipherColumnDefinition()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(143)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(144)
			p.AssistedQueryColumnDefinition()
		}

	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(147)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(148)
			p.LikeQueryColumnDefinition()
		}

	}
	{
		p.SetState(151)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(152)
		p.EncryptAlgorithm()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(153)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(154)
			p.AssistedQueryAlgorithm()
		}

	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(157)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(158)
			p.LikeQueryAlgorithm()
		}

	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(161)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(162)
			p.Match(RDLStatementParserQUERY_WITH_CIPHER_COLUMN)
		}
		{
			p.SetState(163)
			p.Match(RDLStatementParserEQ_)
		}
		{
			p.SetState(164)
			p.QueryWithCipherColumn()
		}

	}
	{
		p.SetState(167)
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

func (s *ColumnDefinitionContext) NAME() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserNAME, 0)
}

func (s *ColumnDefinitionContext) AllEQ_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserEQ_)
}

func (s *ColumnDefinitionContext) EQ_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, i)
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

func (s *ColumnDefinitionContext) DATA_TYPE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDATA_TYPE, 0)
}

func (s *ColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
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
	p.EnterRule(localctx, 14, RDLStatementParserRULE_columnDefinition)

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
		p.SetState(169)
		p.Match(RDLStatementParserNAME)
	}
	{
		p.SetState(170)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(171)
		p.ColumnName()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(172)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(173)
			p.Match(RDLStatementParserDATA_TYPE)
		}
		{
			p.SetState(174)
			p.Match(RDLStatementParserEQ_)
		}
		{
			p.SetState(175)
			p.DataType()
		}

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
	p.EnterRule(localctx, 16, RDLStatementParserRULE_columnName)

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
		p.SetState(178)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dataType
	return p
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDataType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RDLStatementParserRULE_dataType)

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
		p.SetState(180)
		p.Match(RDLStatementParserSTRING_)
	}

	return localctx
}

// IPlainColumnDefinitionContext is an interface to support dynamic dispatch.
type IPlainColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlainColumnDefinitionContext differentiates from other interfaces.
	IsPlainColumnDefinitionContext()
}

type PlainColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlainColumnDefinitionContext() *PlainColumnDefinitionContext {
	var p = new(PlainColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_plainColumnDefinition
	return p
}

func (*PlainColumnDefinitionContext) IsPlainColumnDefinitionContext() {}

func NewPlainColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainColumnDefinitionContext {
	var p = new(PlainColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_plainColumnDefinition

	return p
}

func (s *PlainColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainColumnDefinitionContext) PLAIN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserPLAIN, 0)
}

func (s *PlainColumnDefinitionContext) AllEQ_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserEQ_)
}

func (s *PlainColumnDefinitionContext) EQ_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, i)
}

func (s *PlainColumnDefinitionContext) PlainColumnName() IPlainColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlainColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlainColumnNameContext)
}

func (s *PlainColumnDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *PlainColumnDefinitionContext) PLAIN_DATA_TYPE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserPLAIN_DATA_TYPE, 0)
}

func (s *PlainColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *PlainColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitPlainColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) PlainColumnDefinition() (localctx IPlainColumnDefinitionContext) {
	localctx = NewPlainColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RDLStatementParserRULE_plainColumnDefinition)

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
		p.SetState(182)
		p.Match(RDLStatementParserPLAIN)
	}
	{
		p.SetState(183)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(184)
		p.PlainColumnName()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(185)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(186)
			p.Match(RDLStatementParserPLAIN_DATA_TYPE)
		}
		{
			p.SetState(187)
			p.Match(RDLStatementParserEQ_)
		}
		{
			p.SetState(188)
			p.DataType()
		}

	}

	return localctx
}

// IPlainColumnNameContext is an interface to support dynamic dispatch.
type IPlainColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlainColumnNameContext differentiates from other interfaces.
	IsPlainColumnNameContext()
}

type PlainColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlainColumnNameContext() *PlainColumnNameContext {
	var p = new(PlainColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_plainColumnName
	return p
}

func (*PlainColumnNameContext) IsPlainColumnNameContext() {}

func NewPlainColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainColumnNameContext {
	var p = new(PlainColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_plainColumnName

	return p
}

func (s *PlainColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainColumnNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *PlainColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitPlainColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) PlainColumnName() (localctx IPlainColumnNameContext) {
	localctx = NewPlainColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RDLStatementParserRULE_plainColumnName)

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
		p.SetState(191)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// ICipherColumnDefinitionContext is an interface to support dynamic dispatch.
type ICipherColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCipherColumnDefinitionContext differentiates from other interfaces.
	IsCipherColumnDefinitionContext()
}

type CipherColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCipherColumnDefinitionContext() *CipherColumnDefinitionContext {
	var p = new(CipherColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_cipherColumnDefinition
	return p
}

func (*CipherColumnDefinitionContext) IsCipherColumnDefinitionContext() {}

func NewCipherColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CipherColumnDefinitionContext {
	var p = new(CipherColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_cipherColumnDefinition

	return p
}

func (s *CipherColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *CipherColumnDefinitionContext) CIPHER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCIPHER, 0)
}

func (s *CipherColumnDefinitionContext) AllEQ_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserEQ_)
}

func (s *CipherColumnDefinitionContext) EQ_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, i)
}

func (s *CipherColumnDefinitionContext) CipherColumnName() ICipherColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICipherColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICipherColumnNameContext)
}

func (s *CipherColumnDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *CipherColumnDefinitionContext) CIPHER_DATA_TYPE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCIPHER_DATA_TYPE, 0)
}

func (s *CipherColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *CipherColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CipherColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CipherColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCipherColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CipherColumnDefinition() (localctx ICipherColumnDefinitionContext) {
	localctx = NewCipherColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RDLStatementParserRULE_cipherColumnDefinition)

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
		p.SetState(193)
		p.Match(RDLStatementParserCIPHER)
	}
	{
		p.SetState(194)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(195)
		p.CipherColumnName()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(196)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(197)
			p.Match(RDLStatementParserCIPHER_DATA_TYPE)
		}
		{
			p.SetState(198)
			p.Match(RDLStatementParserEQ_)
		}
		{
			p.SetState(199)
			p.DataType()
		}

	}

	return localctx
}

// ICipherColumnNameContext is an interface to support dynamic dispatch.
type ICipherColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCipherColumnNameContext differentiates from other interfaces.
	IsCipherColumnNameContext()
}

type CipherColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCipherColumnNameContext() *CipherColumnNameContext {
	var p = new(CipherColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_cipherColumnName
	return p
}

func (*CipherColumnNameContext) IsCipherColumnNameContext() {}

func NewCipherColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CipherColumnNameContext {
	var p = new(CipherColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_cipherColumnName

	return p
}

func (s *CipherColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CipherColumnNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *CipherColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CipherColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CipherColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCipherColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CipherColumnName() (localctx ICipherColumnNameContext) {
	localctx = NewCipherColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RDLStatementParserRULE_cipherColumnName)

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
		p.SetState(202)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IAssistedQueryColumnDefinitionContext is an interface to support dynamic dispatch.
type IAssistedQueryColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssistedQueryColumnDefinitionContext differentiates from other interfaces.
	IsAssistedQueryColumnDefinitionContext()
}

type AssistedQueryColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssistedQueryColumnDefinitionContext() *AssistedQueryColumnDefinitionContext {
	var p = new(AssistedQueryColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_assistedQueryColumnDefinition
	return p
}

func (*AssistedQueryColumnDefinitionContext) IsAssistedQueryColumnDefinitionContext() {}

func NewAssistedQueryColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssistedQueryColumnDefinitionContext {
	var p = new(AssistedQueryColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_assistedQueryColumnDefinition

	return p
}

func (s *AssistedQueryColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssistedQueryColumnDefinitionContext) ASSISTED_QUERY_COLUMN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserASSISTED_QUERY_COLUMN, 0)
}

func (s *AssistedQueryColumnDefinitionContext) AllEQ_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserEQ_)
}

func (s *AssistedQueryColumnDefinitionContext) EQ_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, i)
}

func (s *AssistedQueryColumnDefinitionContext) AssistedQueryColumnName() IAssistedQueryColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssistedQueryColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssistedQueryColumnNameContext)
}

func (s *AssistedQueryColumnDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *AssistedQueryColumnDefinitionContext) ASSISTED_QUERY_DATA_TYPE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserASSISTED_QUERY_DATA_TYPE, 0)
}

func (s *AssistedQueryColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *AssistedQueryColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssistedQueryColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssistedQueryColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAssistedQueryColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AssistedQueryColumnDefinition() (localctx IAssistedQueryColumnDefinitionContext) {
	localctx = NewAssistedQueryColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RDLStatementParserRULE_assistedQueryColumnDefinition)

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
		p.SetState(204)
		p.Match(RDLStatementParserASSISTED_QUERY_COLUMN)
	}
	{
		p.SetState(205)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(206)
		p.AssistedQueryColumnName()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(207)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(208)
			p.Match(RDLStatementParserASSISTED_QUERY_DATA_TYPE)
		}
		{
			p.SetState(209)
			p.Match(RDLStatementParserEQ_)
		}
		{
			p.SetState(210)
			p.DataType()
		}

	}

	return localctx
}

// IAssistedQueryColumnNameContext is an interface to support dynamic dispatch.
type IAssistedQueryColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssistedQueryColumnNameContext differentiates from other interfaces.
	IsAssistedQueryColumnNameContext()
}

type AssistedQueryColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssistedQueryColumnNameContext() *AssistedQueryColumnNameContext {
	var p = new(AssistedQueryColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_assistedQueryColumnName
	return p
}

func (*AssistedQueryColumnNameContext) IsAssistedQueryColumnNameContext() {}

func NewAssistedQueryColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssistedQueryColumnNameContext {
	var p = new(AssistedQueryColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_assistedQueryColumnName

	return p
}

func (s *AssistedQueryColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AssistedQueryColumnNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *AssistedQueryColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssistedQueryColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssistedQueryColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAssistedQueryColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AssistedQueryColumnName() (localctx IAssistedQueryColumnNameContext) {
	localctx = NewAssistedQueryColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RDLStatementParserRULE_assistedQueryColumnName)

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

// ILikeQueryColumnDefinitionContext is an interface to support dynamic dispatch.
type ILikeQueryColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikeQueryColumnDefinitionContext differentiates from other interfaces.
	IsLikeQueryColumnDefinitionContext()
}

type LikeQueryColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLikeQueryColumnDefinitionContext() *LikeQueryColumnDefinitionContext {
	var p = new(LikeQueryColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_likeQueryColumnDefinition
	return p
}

func (*LikeQueryColumnDefinitionContext) IsLikeQueryColumnDefinitionContext() {}

func NewLikeQueryColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeQueryColumnDefinitionContext {
	var p = new(LikeQueryColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_likeQueryColumnDefinition

	return p
}

func (s *LikeQueryColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *LikeQueryColumnDefinitionContext) LIKE_QUERY_COLUMN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLIKE_QUERY_COLUMN, 0)
}

func (s *LikeQueryColumnDefinitionContext) AllEQ_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserEQ_)
}

func (s *LikeQueryColumnDefinitionContext) EQ_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, i)
}

func (s *LikeQueryColumnDefinitionContext) LikeQueryColumnName() ILikeQueryColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeQueryColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeQueryColumnNameContext)
}

func (s *LikeQueryColumnDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *LikeQueryColumnDefinitionContext) LIKE_QUERY_DATA_TYPE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLIKE_QUERY_DATA_TYPE, 0)
}

func (s *LikeQueryColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *LikeQueryColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeQueryColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeQueryColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitLikeQueryColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) LikeQueryColumnDefinition() (localctx ILikeQueryColumnDefinitionContext) {
	localctx = NewLikeQueryColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RDLStatementParserRULE_likeQueryColumnDefinition)

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
		p.SetState(215)
		p.Match(RDLStatementParserLIKE_QUERY_COLUMN)
	}
	{
		p.SetState(216)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(217)
		p.LikeQueryColumnName()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(218)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(219)
			p.Match(RDLStatementParserLIKE_QUERY_DATA_TYPE)
		}
		{
			p.SetState(220)
			p.Match(RDLStatementParserEQ_)
		}
		{
			p.SetState(221)
			p.DataType()
		}

	}

	return localctx
}

// ILikeQueryColumnNameContext is an interface to support dynamic dispatch.
type ILikeQueryColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikeQueryColumnNameContext differentiates from other interfaces.
	IsLikeQueryColumnNameContext()
}

type LikeQueryColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLikeQueryColumnNameContext() *LikeQueryColumnNameContext {
	var p = new(LikeQueryColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_likeQueryColumnName
	return p
}

func (*LikeQueryColumnNameContext) IsLikeQueryColumnNameContext() {}

func NewLikeQueryColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeQueryColumnNameContext {
	var p = new(LikeQueryColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_likeQueryColumnName

	return p
}

func (s *LikeQueryColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *LikeQueryColumnNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *LikeQueryColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeQueryColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeQueryColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitLikeQueryColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) LikeQueryColumnName() (localctx ILikeQueryColumnNameContext) {
	localctx = NewLikeQueryColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RDLStatementParserRULE_likeQueryColumnName)

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
		p.SetState(224)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IEncryptAlgorithmContext is an interface to support dynamic dispatch.
type IEncryptAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncryptAlgorithmContext differentiates from other interfaces.
	IsEncryptAlgorithmContext()
}

type EncryptAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncryptAlgorithmContext() *EncryptAlgorithmContext {
	var p = new(EncryptAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_encryptAlgorithm
	return p
}

func (*EncryptAlgorithmContext) IsEncryptAlgorithmContext() {}

func NewEncryptAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncryptAlgorithmContext {
	var p = new(EncryptAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_encryptAlgorithm

	return p
}

func (s *EncryptAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *EncryptAlgorithmContext) ENCRYPT_ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserENCRYPT_ALGORITHM, 0)
}

func (s *EncryptAlgorithmContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *EncryptAlgorithmContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *EncryptAlgorithmContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *EncryptAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncryptAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncryptAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitEncryptAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) EncryptAlgorithm() (localctx IEncryptAlgorithmContext) {
	localctx = NewEncryptAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RDLStatementParserRULE_encryptAlgorithm)

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
		p.SetState(226)
		p.Match(RDLStatementParserENCRYPT_ALGORITHM)
	}
	{
		p.SetState(227)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(228)
		p.AlgorithmDefinition()
	}
	{
		p.SetState(229)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IAssistedQueryAlgorithmContext is an interface to support dynamic dispatch.
type IAssistedQueryAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssistedQueryAlgorithmContext differentiates from other interfaces.
	IsAssistedQueryAlgorithmContext()
}

type AssistedQueryAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssistedQueryAlgorithmContext() *AssistedQueryAlgorithmContext {
	var p = new(AssistedQueryAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_assistedQueryAlgorithm
	return p
}

func (*AssistedQueryAlgorithmContext) IsAssistedQueryAlgorithmContext() {}

func NewAssistedQueryAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssistedQueryAlgorithmContext {
	var p = new(AssistedQueryAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_assistedQueryAlgorithm

	return p
}

func (s *AssistedQueryAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *AssistedQueryAlgorithmContext) ASSISTED_QUERY_ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserASSISTED_QUERY_ALGORITHM, 0)
}

func (s *AssistedQueryAlgorithmContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *AssistedQueryAlgorithmContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *AssistedQueryAlgorithmContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *AssistedQueryAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssistedQueryAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssistedQueryAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAssistedQueryAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AssistedQueryAlgorithm() (localctx IAssistedQueryAlgorithmContext) {
	localctx = NewAssistedQueryAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RDLStatementParserRULE_assistedQueryAlgorithm)

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
		p.SetState(231)
		p.Match(RDLStatementParserASSISTED_QUERY_ALGORITHM)
	}
	{
		p.SetState(232)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(233)
		p.AlgorithmDefinition()
	}
	{
		p.SetState(234)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// ILikeQueryAlgorithmContext is an interface to support dynamic dispatch.
type ILikeQueryAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikeQueryAlgorithmContext differentiates from other interfaces.
	IsLikeQueryAlgorithmContext()
}

type LikeQueryAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLikeQueryAlgorithmContext() *LikeQueryAlgorithmContext {
	var p = new(LikeQueryAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_likeQueryAlgorithm
	return p
}

func (*LikeQueryAlgorithmContext) IsLikeQueryAlgorithmContext() {}

func NewLikeQueryAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeQueryAlgorithmContext {
	var p = new(LikeQueryAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_likeQueryAlgorithm

	return p
}

func (s *LikeQueryAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *LikeQueryAlgorithmContext) LIKE_QUERY_ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLIKE_QUERY_ALGORITHM, 0)
}

func (s *LikeQueryAlgorithmContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *LikeQueryAlgorithmContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *LikeQueryAlgorithmContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *LikeQueryAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeQueryAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeQueryAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitLikeQueryAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) LikeQueryAlgorithm() (localctx ILikeQueryAlgorithmContext) {
	localctx = NewLikeQueryAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RDLStatementParserRULE_likeQueryAlgorithm)

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
		p.SetState(236)
		p.Match(RDLStatementParserLIKE_QUERY_ALGORITHM)
	}
	{
		p.SetState(237)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(238)
		p.AlgorithmDefinition()
	}
	{
		p.SetState(239)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IQueryWithCipherColumnContext is an interface to support dynamic dispatch.
type IQueryWithCipherColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryWithCipherColumnContext differentiates from other interfaces.
	IsQueryWithCipherColumnContext()
}

type QueryWithCipherColumnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryWithCipherColumnContext() *QueryWithCipherColumnContext {
	var p = new(QueryWithCipherColumnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_queryWithCipherColumn
	return p
}

func (*QueryWithCipherColumnContext) IsQueryWithCipherColumnContext() {}

func NewQueryWithCipherColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryWithCipherColumnContext {
	var p = new(QueryWithCipherColumnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_queryWithCipherColumn

	return p
}

func (s *QueryWithCipherColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryWithCipherColumnContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTRUE, 0)
}

func (s *QueryWithCipherColumnContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserFALSE, 0)
}

func (s *QueryWithCipherColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryWithCipherColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryWithCipherColumnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitQueryWithCipherColumn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) QueryWithCipherColumn() (localctx IQueryWithCipherColumnContext) {
	localctx = NewQueryWithCipherColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RDLStatementParserRULE_queryWithCipherColumn)
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
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RDLStatementParserTRUE || _la == RDLStatementParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 44, RDLStatementParserRULE_ifExists)

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
		p.SetState(243)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(244)
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
	p.EnterRule(localctx, 46, RDLStatementParserRULE_ifNotExists)

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
		p.SetState(246)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(247)
		p.Match(RDLStatementParserNOT)
	}
	{
		p.SetState(248)
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
	p.EnterRule(localctx, 48, RDLStatementParserRULE_literal)
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

	p.SetState(257)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserMINUS_, RDLStatementParserINT_:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RDLStatementParserMINUS_ {
			{
				p.SetState(251)
				p.Match(RDLStatementParserMINUS_)
			}

		}
		{
			p.SetState(254)
			p.Match(RDLStatementParserINT_)
		}

	case RDLStatementParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.Match(RDLStatementParserTRUE)
		}

	case RDLStatementParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(256)
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
	p.EnterRule(localctx, 50, RDLStatementParserRULE_algorithmDefinition)
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
		p.SetState(259)
		p.Match(RDLStatementParserTYPE)
	}
	{
		p.SetState(260)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(261)
		p.Match(RDLStatementParserNAME)
	}
	{
		p.SetState(262)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(263)
		p.AlgorithmTypeName()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(264)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(265)
			p.PropertiesDefinition()
		}

	}
	{
		p.SetState(268)
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

func (s *AlgorithmTypeNameContext) BuildinAlgorithmTypeName() IBuildinAlgorithmTypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildinAlgorithmTypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildinAlgorithmTypeNameContext)
}

func (s *AlgorithmTypeNameContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
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
	p.EnterRule(localctx, 52, RDLStatementParserRULE_algorithmTypeName)

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

	p.SetState(272)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserMD5, RDLStatementParserAES, RDLStatementParserRC4, RDLStatementParserSM3, RDLStatementParserSM4, RDLStatementParserCHAR_DIGEST_LIKE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.BuildinAlgorithmTypeName()
		}

	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
			p.Match(RDLStatementParserSTRING_)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuildinAlgorithmTypeNameContext is an interface to support dynamic dispatch.
type IBuildinAlgorithmTypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildinAlgorithmTypeNameContext differentiates from other interfaces.
	IsBuildinAlgorithmTypeNameContext()
}

type BuildinAlgorithmTypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildinAlgorithmTypeNameContext() *BuildinAlgorithmTypeNameContext {
	var p = new(BuildinAlgorithmTypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildinAlgorithmTypeName
	return p
}

func (*BuildinAlgorithmTypeNameContext) IsBuildinAlgorithmTypeNameContext() {}

func NewBuildinAlgorithmTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildinAlgorithmTypeNameContext {
	var p = new(BuildinAlgorithmTypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildinAlgorithmTypeName

	return p
}

func (s *BuildinAlgorithmTypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildinAlgorithmTypeNameContext) MD5() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMD5, 0)
}

func (s *BuildinAlgorithmTypeNameContext) AES() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserAES, 0)
}

func (s *BuildinAlgorithmTypeNameContext) RC4() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRC4, 0)
}

func (s *BuildinAlgorithmTypeNameContext) SM3() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSM3, 0)
}

func (s *BuildinAlgorithmTypeNameContext) SM4() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSM4, 0)
}

func (s *BuildinAlgorithmTypeNameContext) CHAR_DIGEST_LIKE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCHAR_DIGEST_LIKE, 0)
}

func (s *BuildinAlgorithmTypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildinAlgorithmTypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildinAlgorithmTypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildinAlgorithmTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildinAlgorithmTypeName() (localctx IBuildinAlgorithmTypeNameContext) {
	localctx = NewBuildinAlgorithmTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RDLStatementParserRULE_buildinAlgorithmTypeName)
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
		p.SetState(274)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(RDLStatementParserMD5-77))|(1<<(RDLStatementParserAES-77))|(1<<(RDLStatementParserRC4-77))|(1<<(RDLStatementParserSM3-77))|(1<<(RDLStatementParserSM4-77))|(1<<(RDLStatementParserCHAR_DIGEST_LIKE-77)))) != 0) {
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
	p.EnterRule(localctx, 56, RDLStatementParserRULE_propertiesDefinition)
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
		p.SetState(276)
		p.Match(RDLStatementParserPROPERTIES)
	}
	{
		p.SetState(277)
		p.Match(RDLStatementParserLP_)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserSTRING_ {
		{
			p.SetState(278)
			p.Properties()
		}

	}
	{
		p.SetState(281)
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
	p.EnterRule(localctx, 58, RDLStatementParserRULE_properties)
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
		p.SetState(283)
		p.Property()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(284)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(285)
			p.Property()
		}

		p.SetState(290)
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
	p.EnterRule(localctx, 60, RDLStatementParserRULE_property)

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
		p.SetState(291)

		var _m = p.Match(RDLStatementParserSTRING_)

		localctx.(*PropertyContext).key = _m
	}
	{
		p.SetState(292)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(293)

		var _x = p.Literal()

		localctx.(*PropertyContext).value = _x
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
	p.EnterRule(localctx, 62, RDLStatementParserRULE_tableName)

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
		p.SetState(295)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}
