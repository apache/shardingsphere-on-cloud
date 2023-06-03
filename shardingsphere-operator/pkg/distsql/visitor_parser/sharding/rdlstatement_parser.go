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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 120, 543,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 118, 10, 2, 3,
	2, 3, 2, 3, 2, 7, 2, 123, 10, 2, 12, 2, 14, 2, 126, 11, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 135, 10, 3, 12, 3, 14, 3, 138, 11,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 145, 10, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 150, 10, 4, 12, 4, 14, 4, 153, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 5, 5, 161, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 166, 10, 5, 12, 5, 14,
	5, 169, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 179,
	10, 6, 12, 6, 14, 6, 182, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 190, 10, 7, 3, 7, 3, 7, 3, 7, 7, 7, 195, 10, 7, 12, 7, 14, 7, 198, 11,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 205, 10, 8, 3, 8, 3, 8, 3, 8, 7,
	8, 210, 10, 8, 12, 8, 14, 8, 213, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 220, 10, 9, 3, 9, 3, 9, 3, 9, 7, 9, 225, 10, 9, 12, 9, 14, 9, 228,
	11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 234, 10, 10, 3, 10, 3, 10, 3,
	10, 7, 10, 239, 10, 10, 12, 10, 14, 10, 242, 11, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 250, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 271, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 278, 10, 14, 3, 14, 3, 14, 3, 14, 7, 14, 283, 10, 14,
	12, 14, 14, 14, 286, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 292, 10,
	15, 3, 15, 3, 15, 3, 15, 7, 15, 297, 10, 15, 12, 15, 14, 15, 300, 11, 15,
	3, 16, 3, 16, 5, 16, 304, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 315, 10, 17, 3, 17, 3, 17, 5, 17, 319,
	10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 328, 10,
	18, 3, 18, 3, 18, 5, 18, 332, 10, 18, 3, 18, 3, 18, 5, 18, 336, 10, 18,
	3, 18, 3, 18, 5, 18, 340, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	7, 22, 358, 10, 22, 12, 22, 14, 22, 361, 11, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 372, 10, 24, 12, 24, 14,
	24, 375, 11, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27,
	5, 27, 385, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 7, 29, 398, 10, 29, 12, 29, 14, 29, 401, 11, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5,
	31, 413, 10, 31, 3, 31, 3, 31, 5, 31, 417, 10, 31, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 7, 36, 450, 10, 36, 12,
	36, 14, 36, 453, 11, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 40, 7, 40, 466, 10, 40, 12, 40, 14, 40, 469, 11,
	40, 3, 40, 3, 40, 3, 41, 3, 41, 5, 41, 475, 10, 41, 3, 42, 3, 42, 3, 43,
	3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 5, 45, 488, 10,
	45, 3, 45, 3, 45, 3, 45, 5, 45, 493, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 5, 46, 502, 10, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 3, 47, 5, 47, 510, 10, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50,
	3, 51, 3, 51, 3, 51, 5, 51, 521, 10, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3,
	52, 7, 52, 528, 10, 52, 12, 52, 14, 52, 531, 11, 52, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 2, 2, 57, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
	80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 2,
	8, 4, 2, 56, 56, 71, 71, 3, 2, 118, 119, 3, 2, 92, 93, 5, 2, 69, 69, 112,
	113, 116, 116, 3, 2, 94, 106, 3, 2, 107, 111, 2, 535, 2, 112, 3, 2, 2,
	2, 4, 127, 3, 2, 2, 2, 6, 139, 3, 2, 2, 2, 8, 154, 3, 2, 2, 2, 10, 170,
	3, 2, 2, 2, 12, 183, 3, 2, 2, 2, 14, 199, 3, 2, 2, 2, 16, 214, 3, 2, 2,
	2, 18, 229, 3, 2, 2, 2, 20, 243, 3, 2, 2, 2, 22, 255, 3, 2, 2, 2, 24, 264,
	3, 2, 2, 2, 26, 272, 3, 2, 2, 2, 28, 287, 3, 2, 2, 2, 30, 303, 3, 2, 2,
	2, 32, 305, 3, 2, 2, 2, 34, 322, 3, 2, 2, 2, 36, 343, 3, 2, 2, 2, 38, 345,
	3, 2, 2, 2, 40, 350, 3, 2, 2, 2, 42, 352, 3, 2, 2, 2, 44, 364, 3, 2, 2,
	2, 46, 366, 3, 2, 2, 2, 48, 378, 3, 2, 2, 2, 50, 380, 3, 2, 2, 2, 52, 384,
	3, 2, 2, 2, 54, 386, 3, 2, 2, 2, 56, 390, 3, 2, 2, 2, 58, 402, 3, 2, 2,
	2, 60, 407, 3, 2, 2, 2, 62, 418, 3, 2, 2, 2, 64, 423, 3, 2, 2, 2, 66, 428,
	3, 2, 2, 2, 68, 437, 3, 2, 2, 2, 70, 446, 3, 2, 2, 2, 72, 454, 3, 2, 2,
	2, 74, 456, 3, 2, 2, 2, 76, 458, 3, 2, 2, 2, 78, 460, 3, 2, 2, 2, 80, 474,
	3, 2, 2, 2, 82, 476, 3, 2, 2, 2, 84, 478, 3, 2, 2, 2, 86, 481, 3, 2, 2,
	2, 88, 492, 3, 2, 2, 2, 90, 494, 3, 2, 2, 2, 92, 509, 3, 2, 2, 2, 94, 511,
	3, 2, 2, 2, 96, 513, 3, 2, 2, 2, 98, 515, 3, 2, 2, 2, 100, 517, 3, 2, 2,
	2, 102, 524, 3, 2, 2, 2, 104, 532, 3, 2, 2, 2, 106, 536, 3, 2, 2, 2, 108,
	538, 3, 2, 2, 2, 110, 540, 3, 2, 2, 2, 112, 113, 7, 46, 2, 2, 113, 114,
	7, 50, 2, 2, 114, 115, 7, 56, 2, 2, 115, 117, 7, 51, 2, 2, 116, 118, 5,
	86, 44, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 3, 2,
	2, 2, 119, 124, 5, 30, 16, 2, 120, 121, 7, 36, 2, 2, 121, 123, 5, 30, 16,
	2, 122, 120, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124,
	125, 3, 2, 2, 2, 125, 3, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 7,
	47, 2, 2, 128, 129, 7, 50, 2, 2, 129, 130, 7, 56, 2, 2, 130, 131, 7, 51,
	2, 2, 131, 136, 5, 30, 16, 2, 132, 133, 7, 36, 2, 2, 133, 135, 5, 30, 16,
	2, 134, 132, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 5, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 140, 7,
	48, 2, 2, 140, 141, 7, 50, 2, 2, 141, 142, 7, 56, 2, 2, 142, 144, 7, 51,
	2, 2, 143, 145, 5, 84, 43, 2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2,
	2, 145, 146, 3, 2, 2, 2, 146, 151, 5, 106, 54, 2, 147, 148, 7, 36, 2, 2,
	148, 150, 5, 106, 54, 2, 149, 147, 3, 2, 2, 2, 150, 153, 3, 2, 2, 2, 151,
	149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 7, 3, 2, 2, 2, 153, 151, 3,
	2, 2, 2, 154, 155, 7, 46, 2, 2, 155, 156, 7, 50, 2, 2, 156, 157, 7, 56,
	2, 2, 157, 158, 7, 63, 2, 2, 158, 160, 7, 51, 2, 2, 159, 161, 5, 86, 44,
	2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162,
	167, 5, 78, 40, 2, 163, 164, 7, 36, 2, 2, 164, 166, 5, 78, 40, 2, 165,
	163, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168,
	3, 2, 2, 2, 168, 9, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 171, 7, 47,
	2, 2, 171, 172, 7, 50, 2, 2, 172, 173, 7, 56, 2, 2, 173, 174, 7, 63, 2,
	2, 174, 175, 7, 51, 2, 2, 175, 180, 5, 78, 40, 2, 176, 177, 7, 36, 2, 2,
	177, 179, 5, 78, 40, 2, 178, 176, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180,
	178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 11, 3, 2, 2, 2, 182, 180, 3,
	2, 2, 2, 183, 184, 7, 48, 2, 2, 184, 185, 7, 50, 2, 2, 185, 186, 7, 56,
	2, 2, 186, 187, 7, 63, 2, 2, 187, 189, 7, 51, 2, 2, 188, 190, 5, 84, 43,
	2, 189, 188, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	196, 5, 110, 56, 2, 192, 193, 7, 36, 2, 2, 193, 195, 5, 110, 56, 2, 194,
	192, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197,
	3, 2, 2, 2, 197, 13, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 199, 200, 7, 46,
	2, 2, 200, 201, 7, 64, 2, 2, 201, 202, 7, 56, 2, 2, 202, 204, 7, 51, 2,
	2, 203, 205, 5, 86, 44, 2, 204, 203, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2,
	205, 206, 3, 2, 2, 2, 206, 211, 5, 106, 54, 2, 207, 208, 7, 36, 2, 2, 208,
	210, 5, 106, 54, 2, 209, 207, 3, 2, 2, 2, 210, 213, 3, 2, 2, 2, 211, 209,
	3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 15, 3, 2, 2, 2, 213, 211, 3, 2,
	2, 2, 214, 215, 7, 48, 2, 2, 215, 216, 7, 64, 2, 2, 216, 217, 7, 56, 2,
	2, 217, 219, 7, 51, 2, 2, 218, 220, 5, 84, 43, 2, 219, 218, 3, 2, 2, 2,
	219, 220, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 226, 5, 106, 54, 2, 222,
	223, 7, 36, 2, 2, 223, 225, 5, 106, 54, 2, 224, 222, 3, 2, 2, 2, 225, 228,
	3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 17, 3, 2,
	2, 2, 228, 226, 3, 2, 2, 2, 229, 230, 7, 48, 2, 2, 230, 231, 7, 50, 2,
	2, 231, 233, 7, 67, 2, 2, 232, 234, 5, 84, 43, 2, 233, 232, 3, 2, 2, 2,
	233, 234, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 240, 5, 108, 55, 2, 236,
	237, 7, 36, 2, 2, 237, 239, 5, 108, 55, 2, 238, 236, 3, 2, 2, 2, 239, 242,
	3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 19, 3, 2,
	2, 2, 242, 240, 3, 2, 2, 2, 243, 244, 7, 46, 2, 2, 244, 245, 7, 70, 2,
	2, 245, 246, 7, 50, 2, 2, 246, 247, 9, 2, 2, 2, 247, 249, 7, 73, 2, 2,
	248, 250, 5, 86, 44, 2, 249, 248, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250,
	251, 3, 2, 2, 2, 251, 252, 7, 30, 2, 2, 252, 253, 5, 60, 31, 2, 253, 254,
	7, 31, 2, 2, 254, 21, 3, 2, 2, 2, 255, 256, 7, 47, 2, 2, 256, 257, 7, 70,
	2, 2, 257, 258, 7, 50, 2, 2, 258, 259, 9, 2, 2, 2, 259, 260, 7, 73, 2,
	2, 260, 261, 7, 30, 2, 2, 261, 262, 5, 60, 31, 2, 262, 263, 7, 31, 2, 2,
	263, 23, 3, 2, 2, 2, 264, 265, 7, 48, 2, 2, 265, 266, 7, 70, 2, 2, 266,
	267, 7, 50, 2, 2, 267, 268, 9, 2, 2, 2, 268, 270, 7, 73, 2, 2, 269, 271,
	5, 84, 43, 2, 270, 269, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 25, 3, 2,
	2, 2, 272, 273, 7, 48, 2, 2, 273, 274, 7, 50, 2, 2, 274, 275, 7, 78, 2,
	2, 275, 277, 7, 79, 2, 2, 276, 278, 5, 84, 43, 2, 277, 276, 3, 2, 2, 2,
	277, 278, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 284, 5, 36, 19, 2, 280,
	281, 7, 36, 2, 2, 281, 283, 5, 36, 19, 2, 282, 280, 3, 2, 2, 2, 283, 286,
	3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 27, 3, 2,
	2, 2, 286, 284, 3, 2, 2, 2, 287, 288, 7, 48, 2, 2, 288, 289, 7, 50, 2,
	2, 289, 291, 7, 88, 2, 2, 290, 292, 5, 84, 43, 2, 291, 290, 3, 2, 2, 2,
	291, 292, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 298, 5, 40, 21, 2, 294,
	295, 7, 36, 2, 2, 295, 297, 5, 40, 21, 2, 296, 294, 3, 2, 2, 2, 297, 300,
	3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 29, 3, 2,
	2, 2, 300, 298, 3, 2, 2, 2, 301, 304, 5, 32, 17, 2, 302, 304, 5, 34, 18,
	2, 303, 301, 3, 2, 2, 2, 303, 302, 3, 2, 2, 2, 304, 31, 3, 2, 2, 2, 305,
	306, 5, 106, 54, 2, 306, 307, 7, 30, 2, 2, 307, 308, 5, 42, 22, 2, 308,
	309, 7, 36, 2, 2, 309, 310, 5, 50, 26, 2, 310, 311, 7, 36, 2, 2, 311, 314,
	5, 90, 46, 2, 312, 313, 7, 36, 2, 2, 313, 315, 5, 66, 34, 2, 314, 312,
	3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 318, 3, 2, 2, 2, 316, 317, 7, 36,
	2, 2, 317, 319, 5, 68, 35, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2,
	2, 319, 320, 3, 2, 2, 2, 320, 321, 7, 31, 2, 2, 321, 33, 3, 2, 2, 2, 322,
	323, 5, 106, 54, 2, 323, 324, 7, 30, 2, 2, 324, 327, 5, 46, 24, 2, 325,
	326, 7, 36, 2, 2, 326, 328, 5, 62, 32, 2, 327, 325, 3, 2, 2, 2, 327, 328,
	3, 2, 2, 2, 328, 331, 3, 2, 2, 2, 329, 330, 7, 36, 2, 2, 330, 332, 5, 64,
	33, 2, 331, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 335, 3, 2, 2, 2,
	333, 334, 7, 36, 2, 2, 334, 336, 5, 66, 34, 2, 335, 333, 3, 2, 2, 2, 335,
	336, 3, 2, 2, 2, 336, 339, 3, 2, 2, 2, 337, 338, 7, 36, 2, 2, 338, 340,
	5, 68, 35, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 341, 3,
	2, 2, 2, 341, 342, 7, 31, 2, 2, 342, 35, 3, 2, 2, 2, 343, 344, 7, 118,
	2, 2, 344, 37, 3, 2, 2, 2, 345, 346, 5, 40, 21, 2, 346, 347, 7, 30, 2,
	2, 347, 348, 5, 90, 46, 2, 348, 349, 7, 31, 2, 2, 349, 39, 3, 2, 2, 2,
	350, 351, 7, 118, 2, 2, 351, 41, 3, 2, 2, 2, 352, 353, 7, 53, 2, 2, 353,
	354, 7, 30, 2, 2, 354, 359, 5, 44, 23, 2, 355, 356, 7, 36, 2, 2, 356, 358,
	5, 44, 23, 2, 357, 355, 3, 2, 2, 2, 358, 361, 3, 2, 2, 2, 359, 357, 3,
	2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 362, 3, 2, 2, 2, 361, 359, 3, 2, 2,
	2, 362, 363, 7, 31, 2, 2, 363, 43, 3, 2, 2, 2, 364, 365, 9, 3, 2, 2, 365,
	45, 3, 2, 2, 2, 366, 367, 7, 74, 2, 2, 367, 368, 7, 30, 2, 2, 368, 373,
	5, 48, 25, 2, 369, 370, 7, 36, 2, 2, 370, 372, 5, 48, 25, 2, 371, 369,
	3, 2, 2, 2, 372, 375, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 374, 3, 2,
	2, 2, 374, 376, 3, 2, 2, 2, 375, 373, 3, 2, 2, 2, 376, 377, 7, 31, 2, 2,
	377, 47, 3, 2, 2, 2, 378, 379, 7, 119, 2, 2, 379, 49, 3, 2, 2, 2, 380,
	381, 5, 54, 28, 2, 381, 51, 3, 2, 2, 2, 382, 385, 5, 54, 28, 2, 383, 385,
	5, 56, 29, 2, 384, 382, 3, 2, 2, 2, 384, 383, 3, 2, 2, 2, 385, 53, 3, 2,
	2, 2, 386, 387, 7, 57, 2, 2, 387, 388, 7, 23, 2, 2, 388, 389, 5, 76, 39,
	2, 389, 55, 3, 2, 2, 2, 390, 391, 7, 58, 2, 2, 391, 392, 7, 23, 2, 2, 392,
	393, 5, 76, 39, 2, 393, 394, 7, 36, 2, 2, 394, 399, 5, 76, 39, 2, 395,
	396, 7, 36, 2, 2, 396, 398, 5, 76, 39, 2, 397, 395, 3, 2, 2, 2, 398, 401,
	3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 57, 3, 2,
	2, 2, 401, 399, 3, 2, 2, 2, 402, 403, 7, 72, 2, 2, 403, 404, 7, 30, 2,
	2, 404, 405, 5, 90, 46, 2, 405, 406, 7, 31, 2, 2, 406, 59, 3, 2, 2, 2,
	407, 408, 7, 59, 2, 2, 408, 409, 7, 23, 2, 2, 409, 416, 5, 80, 41, 2, 410,
	411, 7, 36, 2, 2, 411, 413, 5, 52, 27, 2, 412, 410, 3, 2, 2, 2, 412, 413,
	3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 415, 7, 36, 2, 2, 415, 417, 5, 58,
	30, 2, 416, 412, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 61, 3, 2, 2, 2,
	418, 419, 7, 75, 2, 2, 419, 420, 7, 30, 2, 2, 420, 421, 5, 60, 31, 2, 421,
	422, 7, 31, 2, 2, 422, 63, 3, 2, 2, 2, 423, 424, 7, 76, 2, 2, 424, 425,
	7, 30, 2, 2, 425, 426, 5, 60, 31, 2, 426, 427, 7, 31, 2, 2, 427, 65, 3,
	2, 2, 2, 428, 429, 7, 54, 2, 2, 429, 430, 7, 30, 2, 2, 430, 431, 7, 62,
	2, 2, 431, 432, 7, 23, 2, 2, 432, 433, 5, 76, 39, 2, 433, 434, 7, 36, 2,
	2, 434, 435, 5, 90, 46, 2, 435, 436, 7, 31, 2, 2, 436, 67, 3, 2, 2, 2,
	437, 438, 7, 90, 2, 2, 438, 439, 7, 30, 2, 2, 439, 440, 5, 70, 36, 2, 440,
	441, 7, 36, 2, 2, 441, 442, 7, 91, 2, 2, 442, 443, 7, 23, 2, 2, 443, 444,
	5, 74, 38, 2, 444, 445, 7, 31, 2, 2, 445, 69, 3, 2, 2, 2, 446, 451, 5,
	72, 37, 2, 447, 448, 7, 36, 2, 2, 448, 450, 5, 72, 37, 2, 449, 447, 3,
	2, 2, 2, 450, 453, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 451, 452, 3, 2, 2,
	2, 452, 71, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 454, 455, 5, 90, 46, 2, 455,
	73, 3, 2, 2, 2, 456, 457, 9, 4, 2, 2, 457, 75, 3, 2, 2, 2, 458, 459, 7,
	118, 2, 2, 459, 77, 3, 2, 2, 2, 460, 461, 5, 110, 56, 2, 461, 462, 7, 30,
	2, 2, 462, 467, 5, 106, 54, 2, 463, 464, 7, 36, 2, 2, 464, 466, 5, 106,
	54, 2, 465, 463, 3, 2, 2, 2, 466, 469, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2,
	467, 468, 3, 2, 2, 2, 468, 470, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2, 470,
	471, 7, 31, 2, 2, 471, 79, 3, 2, 2, 2, 472, 475, 7, 119, 2, 2, 473, 475,
	5, 82, 42, 2, 474, 472, 3, 2, 2, 2, 474, 473, 3, 2, 2, 2, 475, 81, 3, 2,
	2, 2, 476, 477, 9, 5, 2, 2, 477, 83, 3, 2, 2, 2, 478, 479, 7, 84, 2, 2,
	479, 480, 7, 85, 2, 2, 480, 85, 3, 2, 2, 2, 481, 482, 7, 84, 2, 2, 482,
	483, 7, 115, 2, 2, 483, 484, 7, 85, 2, 2, 484, 87, 3, 2, 2, 2, 485, 493,
	7, 119, 2, 2, 486, 488, 7, 15, 2, 2, 487, 486, 3, 2, 2, 2, 487, 488, 3,
	2, 2, 2, 488, 489, 3, 2, 2, 2, 489, 493, 7, 120, 2, 2, 490, 493, 7, 92,
	2, 2, 491, 493, 7, 93, 2, 2, 492, 485, 3, 2, 2, 2, 492, 487, 3, 2, 2, 2,
	492, 490, 3, 2, 2, 2, 492, 491, 3, 2, 2, 2, 493, 89, 3, 2, 2, 2, 494, 495,
	7, 59, 2, 2, 495, 496, 7, 30, 2, 2, 496, 497, 7, 60, 2, 2, 497, 498, 7,
	23, 2, 2, 498, 501, 5, 92, 47, 2, 499, 500, 7, 36, 2, 2, 500, 502, 5, 100,
	51, 2, 501, 499, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2,
	503, 504, 7, 31, 2, 2, 504, 91, 3, 2, 2, 2, 505, 510, 7, 119, 2, 2, 506,
	510, 5, 94, 48, 2, 507, 510, 5, 96, 49, 2, 508, 510, 5, 98, 50, 2, 509,
	505, 3, 2, 2, 2, 509, 506, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 509, 508,
	3, 2, 2, 2, 510, 93, 3, 2, 2, 2, 511, 512, 9, 6, 2, 2, 512, 95, 3, 2, 2,
	2, 513, 514, 9, 7, 2, 2, 514, 97, 3, 2, 2, 2, 515, 516, 7, 114, 2, 2, 516,
	99, 3, 2, 2, 2, 517, 518, 7, 61, 2, 2, 518, 520, 7, 30, 2, 2, 519, 521,
	5, 102, 52, 2, 520, 519, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 3,
	2, 2, 2, 522, 523, 7, 31, 2, 2, 523, 101, 3, 2, 2, 2, 524, 529, 5, 104,
	53, 2, 525, 526, 7, 36, 2, 2, 526, 528, 5, 104, 53, 2, 527, 525, 3, 2,
	2, 2, 528, 531, 3, 2, 2, 2, 529, 527, 3, 2, 2, 2, 529, 530, 3, 2, 2, 2,
	530, 103, 3, 2, 2, 2, 531, 529, 3, 2, 2, 2, 532, 533, 7, 119, 2, 2, 533,
	534, 7, 23, 2, 2, 534, 535, 5, 88, 45, 2, 535, 105, 3, 2, 2, 2, 536, 537,
	7, 118, 2, 2, 537, 107, 3, 2, 2, 2, 538, 539, 7, 118, 2, 2, 539, 109, 3,
	2, 2, 2, 540, 541, 7, 118, 2, 2, 541, 111, 3, 2, 2, 2, 46, 117, 124, 136,
	144, 151, 160, 167, 180, 189, 196, 204, 211, 219, 226, 233, 240, 249, 270,
	277, 284, 291, 298, 303, 314, 318, 327, 331, 335, 339, 359, 373, 384, 399,
	412, 416, 451, 467, 474, 487, 492, 501, 509, 520, 529,
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
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "'DO NOT MATCH ANY THING, JUST FOR GENERATOR'",
}
var symbolicNames = []string{
	"", "AND_", "OR_", "NOT_", "TILDE_", "VERTICALBAR_", "AMPERSAND_", "SIGNEDLEFTSHIFT_",
	"SIGNEDRIGHTSHIFT_", "CARET_", "MOD_", "COLON_", "PLUS_", "MINUS_", "ASTERISK_",
	"SLASH_", "BACKSLASH_", "DOT_", "DOTASTERISK_", "SAFEEQ_", "DEQ_", "EQ_",
	"NEQ_", "GT_", "GTE_", "LT_", "LTE_", "POUND_", "LP_", "RP_", "LBE_", "RBE_",
	"LBT_", "RBT_", "COMMA_", "DQ_", "SQ_", "BQ_", "QUESTION_", "AT_", "SEMI_",
	"JSONSEPARATOR_", "UL_", "WS", "CREATE", "ALTER", "DROP", "SHOW", "SHARDING",
	"RULE", "FROM", "STORAGE_UNITS", "KEY_GENERATE_STRATEGY", "DEFAULT_TABLE_STRATEGY",
	"TABLE", "SHARDING_COLUMN", "SHARDING_COLUMNS", "TYPE", "NAME", "PROPERTIES",
	"COLUMN", "REFERENCE", "BROADCAST", "RULES", "COLUMNS", "ALGORITHM", "ALGORITHMS",
	"HINT", "DEFAULT", "DATABASE", "SHARDING_ALGORITHM", "STRATEGY", "DATANODES",
	"DATABASE_STRATEGY", "TABLE_STRATEGY", "NODES", "KEY", "GENERATOR", "GENERATORS",
	"KEY_GENERATOR", "UNUSED", "USED", "IF", "EXISTS", "WITH", "COUNT", "AUDITOR",
	"AUDITORS", "AUDIT_STRATEGY", "ALLOW_HINT_DISABLE", "TRUE", "FALSE", "MOD",
	"COSID_MOD", "HASH_MOD", "VOLUME_RANGE", "BOUNDARY_RANGE", "AUTO_INTERVAL",
	"INLINE", "INTERVAL", "COSID_INTERVAL", "COSID_INTERVAL_SNOWFLAKE", "COMPLEX_INLINE",
	"HINT_INLINE", "CLASS_BASED", "SNOWFLAKE", "NANOID", "UUID", "COSID", "COSID_SNOWFLAKE",
	"STANDARD", "COMPLEX", "DML_SHARDING_CONDITIONS", "NOT", "NONE", "FOR_GENERATOR",
	"IDENTIFIER_", "STRING_", "INT_",
}

var ruleNames = []string{
	"createShardingTableRule", "alterShardingTableRule", "dropShardingTableRule",
	"createShardingTableReferenceRule", "alterShardingTableReferenceRule",
	"dropShardingTableReferenceRule", "createBroadcastTableRule", "dropBroadcastTableRule",
	"dropShardingAlgorithm", "createDefaultShardingStrategy", "alterDefaultShardingStrategy",
	"dropDefaultShardingStrategy", "dropShardingKeyGenerator", "dropShardingAuditor",
	"shardingTableRuleDefinition", "shardingAutoTableRule", "shardingTableRule",
	"keyGeneratorName", "auditorDefinition", "auditorName", "storageUnits",
	"storageUnit", "dataNodes", "dataNode", "autoShardingColumnDefinition",
	"shardingColumnDefinition", "shardingColumn", "shardingColumns", "shardingAlgorithm",
	"shardingStrategy", "databaseStrategy", "tableStrategy", "keyGenerateDefinition",
	"auditDefinition", "multiAuditDefinition", "singleAuditDefinition", "auditAllowHintDisable",
	"columnName", "tableReferenceRuleDefinition", "strategyType", "buildInStrategyType",
	"ifExists", "ifNotExists", "literal", "algorithmDefinition", "algorithmTypeName",
	"buildInShardingAlgorithmType", "buildInKeyGenerateAlgorithmType", "buildInShardingAuditAlgorithmType",
	"propertiesDefinition", "properties", "property", "tableName", "shardingAlgorithmName",
	"ruleName",
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
	RDLStatementParserSHARDING                 = 48
	RDLStatementParserRULE                     = 49
	RDLStatementParserFROM                     = 50
	RDLStatementParserSTORAGE_UNITS            = 51
	RDLStatementParserKEY_GENERATE_STRATEGY    = 52
	RDLStatementParserDEFAULT_TABLE_STRATEGY   = 53
	RDLStatementParserTABLE                    = 54
	RDLStatementParserSHARDING_COLUMN          = 55
	RDLStatementParserSHARDING_COLUMNS         = 56
	RDLStatementParserTYPE                     = 57
	RDLStatementParserNAME                     = 58
	RDLStatementParserPROPERTIES               = 59
	RDLStatementParserCOLUMN                   = 60
	RDLStatementParserREFERENCE                = 61
	RDLStatementParserBROADCAST                = 62
	RDLStatementParserRULES                    = 63
	RDLStatementParserCOLUMNS                  = 64
	RDLStatementParserALGORITHM                = 65
	RDLStatementParserALGORITHMS               = 66
	RDLStatementParserHINT                     = 67
	RDLStatementParserDEFAULT                  = 68
	RDLStatementParserDATABASE                 = 69
	RDLStatementParserSHARDING_ALGORITHM       = 70
	RDLStatementParserSTRATEGY                 = 71
	RDLStatementParserDATANODES                = 72
	RDLStatementParserDATABASE_STRATEGY        = 73
	RDLStatementParserTABLE_STRATEGY           = 74
	RDLStatementParserNODES                    = 75
	RDLStatementParserKEY                      = 76
	RDLStatementParserGENERATOR                = 77
	RDLStatementParserGENERATORS               = 78
	RDLStatementParserKEY_GENERATOR            = 79
	RDLStatementParserUNUSED                   = 80
	RDLStatementParserUSED                     = 81
	RDLStatementParserIF                       = 82
	RDLStatementParserEXISTS                   = 83
	RDLStatementParserWITH                     = 84
	RDLStatementParserCOUNT                    = 85
	RDLStatementParserAUDITOR                  = 86
	RDLStatementParserAUDITORS                 = 87
	RDLStatementParserAUDIT_STRATEGY           = 88
	RDLStatementParserALLOW_HINT_DISABLE       = 89
	RDLStatementParserTRUE                     = 90
	RDLStatementParserFALSE                    = 91
	RDLStatementParserMOD                      = 92
	RDLStatementParserCOSID_MOD                = 93
	RDLStatementParserHASH_MOD                 = 94
	RDLStatementParserVOLUME_RANGE             = 95
	RDLStatementParserBOUNDARY_RANGE           = 96
	RDLStatementParserAUTO_INTERVAL            = 97
	RDLStatementParserINLINE                   = 98
	RDLStatementParserINTERVAL                 = 99
	RDLStatementParserCOSID_INTERVAL           = 100
	RDLStatementParserCOSID_INTERVAL_SNOWFLAKE = 101
	RDLStatementParserCOMPLEX_INLINE           = 102
	RDLStatementParserHINT_INLINE              = 103
	RDLStatementParserCLASS_BASED              = 104
	RDLStatementParserSNOWFLAKE                = 105
	RDLStatementParserNANOID                   = 106
	RDLStatementParserUUID                     = 107
	RDLStatementParserCOSID                    = 108
	RDLStatementParserCOSID_SNOWFLAKE          = 109
	RDLStatementParserSTANDARD                 = 110
	RDLStatementParserCOMPLEX                  = 111
	RDLStatementParserDML_SHARDING_CONDITIONS  = 112
	RDLStatementParserNOT                      = 113
	RDLStatementParserNONE                     = 114
	RDLStatementParserFOR_GENERATOR            = 115
	RDLStatementParserIDENTIFIER_              = 116
	RDLStatementParserSTRING_                  = 117
	RDLStatementParserINT_                     = 118
)

// RDLStatementParser rules.
const (
	RDLStatementParserRULE_createShardingTableRule           = 0
	RDLStatementParserRULE_alterShardingTableRule            = 1
	RDLStatementParserRULE_dropShardingTableRule             = 2
	RDLStatementParserRULE_createShardingTableReferenceRule  = 3
	RDLStatementParserRULE_alterShardingTableReferenceRule   = 4
	RDLStatementParserRULE_dropShardingTableReferenceRule    = 5
	RDLStatementParserRULE_createBroadcastTableRule          = 6
	RDLStatementParserRULE_dropBroadcastTableRule            = 7
	RDLStatementParserRULE_dropShardingAlgorithm             = 8
	RDLStatementParserRULE_createDefaultShardingStrategy     = 9
	RDLStatementParserRULE_alterDefaultShardingStrategy      = 10
	RDLStatementParserRULE_dropDefaultShardingStrategy       = 11
	RDLStatementParserRULE_dropShardingKeyGenerator          = 12
	RDLStatementParserRULE_dropShardingAuditor               = 13
	RDLStatementParserRULE_shardingTableRuleDefinition       = 14
	RDLStatementParserRULE_shardingAutoTableRule             = 15
	RDLStatementParserRULE_shardingTableRule                 = 16
	RDLStatementParserRULE_keyGeneratorName                  = 17
	RDLStatementParserRULE_auditorDefinition                 = 18
	RDLStatementParserRULE_auditorName                       = 19
	RDLStatementParserRULE_storageUnits                      = 20
	RDLStatementParserRULE_storageUnit                       = 21
	RDLStatementParserRULE_dataNodes                         = 22
	RDLStatementParserRULE_dataNode                          = 23
	RDLStatementParserRULE_autoShardingColumnDefinition      = 24
	RDLStatementParserRULE_shardingColumnDefinition          = 25
	RDLStatementParserRULE_shardingColumn                    = 26
	RDLStatementParserRULE_shardingColumns                   = 27
	RDLStatementParserRULE_shardingAlgorithm                 = 28
	RDLStatementParserRULE_shardingStrategy                  = 29
	RDLStatementParserRULE_databaseStrategy                  = 30
	RDLStatementParserRULE_tableStrategy                     = 31
	RDLStatementParserRULE_keyGenerateDefinition             = 32
	RDLStatementParserRULE_auditDefinition                   = 33
	RDLStatementParserRULE_multiAuditDefinition              = 34
	RDLStatementParserRULE_singleAuditDefinition             = 35
	RDLStatementParserRULE_auditAllowHintDisable             = 36
	RDLStatementParserRULE_columnName                        = 37
	RDLStatementParserRULE_tableReferenceRuleDefinition      = 38
	RDLStatementParserRULE_strategyType                      = 39
	RDLStatementParserRULE_buildInStrategyType               = 40
	RDLStatementParserRULE_ifExists                          = 41
	RDLStatementParserRULE_ifNotExists                       = 42
	RDLStatementParserRULE_literal                           = 43
	RDLStatementParserRULE_algorithmDefinition               = 44
	RDLStatementParserRULE_algorithmTypeName                 = 45
	RDLStatementParserRULE_buildInShardingAlgorithmType      = 46
	RDLStatementParserRULE_buildInKeyGenerateAlgorithmType   = 47
	RDLStatementParserRULE_buildInShardingAuditAlgorithmType = 48
	RDLStatementParserRULE_propertiesDefinition              = 49
	RDLStatementParserRULE_properties                        = 50
	RDLStatementParserRULE_property                          = 51
	RDLStatementParserRULE_tableName                         = 52
	RDLStatementParserRULE_shardingAlgorithmName             = 53
	RDLStatementParserRULE_ruleName                          = 54
)

// ICreateShardingTableRuleContext is an interface to support dynamic dispatch.
type ICreateShardingTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateShardingTableRuleContext differentiates from other interfaces.
	IsCreateShardingTableRuleContext()
}

type CreateShardingTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateShardingTableRuleContext() *CreateShardingTableRuleContext {
	var p = new(CreateShardingTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createShardingTableRule
	return p
}

func (*CreateShardingTableRuleContext) IsCreateShardingTableRuleContext() {}

func NewCreateShardingTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateShardingTableRuleContext {
	var p = new(CreateShardingTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createShardingTableRule

	return p
}

func (s *CreateShardingTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateShardingTableRuleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateShardingTableRuleContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *CreateShardingTableRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *CreateShardingTableRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *CreateShardingTableRuleContext) AllShardingTableRuleDefinition() []IShardingTableRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShardingTableRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IShardingTableRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShardingTableRuleDefinitionContext)
		}
	}

	return tst
}

func (s *CreateShardingTableRuleContext) ShardingTableRuleDefinition(i int) IShardingTableRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingTableRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShardingTableRuleDefinitionContext)
}

func (s *CreateShardingTableRuleContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateShardingTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *CreateShardingTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *CreateShardingTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateShardingTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateShardingTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateShardingTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateShardingTableRule() (localctx ICreateShardingTableRuleContext) {
	localctx = NewCreateShardingTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RDLStatementParserRULE_createShardingTableRule)
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
		p.SetState(110)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(111)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(112)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(113)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(114)
			p.IfNotExists()
		}

	}
	{
		p.SetState(117)
		p.ShardingTableRuleDefinition()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(118)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(119)
			p.ShardingTableRuleDefinition()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlterShardingTableRuleContext is an interface to support dynamic dispatch.
type IAlterShardingTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterShardingTableRuleContext differentiates from other interfaces.
	IsAlterShardingTableRuleContext()
}

type AlterShardingTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterShardingTableRuleContext() *AlterShardingTableRuleContext {
	var p = new(AlterShardingTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterShardingTableRule
	return p
}

func (*AlterShardingTableRuleContext) IsAlterShardingTableRuleContext() {}

func NewAlterShardingTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterShardingTableRuleContext {
	var p = new(AlterShardingTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterShardingTableRule

	return p
}

func (s *AlterShardingTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterShardingTableRuleContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterShardingTableRuleContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *AlterShardingTableRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *AlterShardingTableRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *AlterShardingTableRuleContext) AllShardingTableRuleDefinition() []IShardingTableRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShardingTableRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IShardingTableRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShardingTableRuleDefinitionContext)
		}
	}

	return tst
}

func (s *AlterShardingTableRuleContext) ShardingTableRuleDefinition(i int) IShardingTableRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingTableRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShardingTableRuleDefinitionContext)
}

func (s *AlterShardingTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *AlterShardingTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *AlterShardingTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterShardingTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterShardingTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterShardingTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterShardingTableRule() (localctx IAlterShardingTableRuleContext) {
	localctx = NewAlterShardingTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RDLStatementParserRULE_alterShardingTableRule)
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
		p.SetState(125)
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(126)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(127)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(128)
		p.Match(RDLStatementParserRULE)
	}
	{
		p.SetState(129)
		p.ShardingTableRuleDefinition()
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(130)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(131)
			p.ShardingTableRuleDefinition()
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropShardingTableRuleContext is an interface to support dynamic dispatch.
type IDropShardingTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropShardingTableRuleContext differentiates from other interfaces.
	IsDropShardingTableRuleContext()
}

type DropShardingTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropShardingTableRuleContext() *DropShardingTableRuleContext {
	var p = new(DropShardingTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropShardingTableRule
	return p
}

func (*DropShardingTableRuleContext) IsDropShardingTableRuleContext() {}

func NewDropShardingTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropShardingTableRuleContext {
	var p = new(DropShardingTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropShardingTableRule

	return p
}

func (s *DropShardingTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropShardingTableRuleContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropShardingTableRuleContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *DropShardingTableRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *DropShardingTableRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *DropShardingTableRuleContext) AllTableName() []ITableNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableNameContext)(nil)).Elem())
	var tst = make([]ITableNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableNameContext)
		}
	}

	return tst
}

func (s *DropShardingTableRuleContext) TableName(i int) ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DropShardingTableRuleContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropShardingTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropShardingTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropShardingTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropShardingTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropShardingTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropShardingTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropShardingTableRule() (localctx IDropShardingTableRuleContext) {
	localctx = NewDropShardingTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RDLStatementParserRULE_dropShardingTableRule)
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
		p.SetState(137)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(138)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(139)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(140)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(141)
			p.IfExists()
		}

	}
	{
		p.SetState(144)
		p.TableName()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(145)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(146)
			p.TableName()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICreateShardingTableReferenceRuleContext is an interface to support dynamic dispatch.
type ICreateShardingTableReferenceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateShardingTableReferenceRuleContext differentiates from other interfaces.
	IsCreateShardingTableReferenceRuleContext()
}

type CreateShardingTableReferenceRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateShardingTableReferenceRuleContext() *CreateShardingTableReferenceRuleContext {
	var p = new(CreateShardingTableReferenceRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createShardingTableReferenceRule
	return p
}

func (*CreateShardingTableReferenceRuleContext) IsCreateShardingTableReferenceRuleContext() {}

func NewCreateShardingTableReferenceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateShardingTableReferenceRuleContext {
	var p = new(CreateShardingTableReferenceRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createShardingTableReferenceRule

	return p
}

func (s *CreateShardingTableReferenceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateShardingTableReferenceRuleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateShardingTableReferenceRuleContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *CreateShardingTableReferenceRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *CreateShardingTableReferenceRuleContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREFERENCE, 0)
}

func (s *CreateShardingTableReferenceRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *CreateShardingTableReferenceRuleContext) AllTableReferenceRuleDefinition() []ITableReferenceRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableReferenceRuleDefinitionContext)(nil)).Elem())
	var tst = make([]ITableReferenceRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableReferenceRuleDefinitionContext)
		}
	}

	return tst
}

func (s *CreateShardingTableReferenceRuleContext) TableReferenceRuleDefinition(i int) ITableReferenceRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableReferenceRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableReferenceRuleDefinitionContext)
}

func (s *CreateShardingTableReferenceRuleContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateShardingTableReferenceRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *CreateShardingTableReferenceRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *CreateShardingTableReferenceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateShardingTableReferenceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateShardingTableReferenceRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateShardingTableReferenceRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateShardingTableReferenceRule() (localctx ICreateShardingTableReferenceRuleContext) {
	localctx = NewCreateShardingTableReferenceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RDLStatementParserRULE_createShardingTableReferenceRule)
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
		p.SetState(152)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(153)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(154)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(155)
		p.Match(RDLStatementParserREFERENCE)
	}
	{
		p.SetState(156)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(157)
			p.IfNotExists()
		}

	}
	{
		p.SetState(160)
		p.TableReferenceRuleDefinition()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(161)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(162)
			p.TableReferenceRuleDefinition()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlterShardingTableReferenceRuleContext is an interface to support dynamic dispatch.
type IAlterShardingTableReferenceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterShardingTableReferenceRuleContext differentiates from other interfaces.
	IsAlterShardingTableReferenceRuleContext()
}

type AlterShardingTableReferenceRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterShardingTableReferenceRuleContext() *AlterShardingTableReferenceRuleContext {
	var p = new(AlterShardingTableReferenceRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterShardingTableReferenceRule
	return p
}

func (*AlterShardingTableReferenceRuleContext) IsAlterShardingTableReferenceRuleContext() {}

func NewAlterShardingTableReferenceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterShardingTableReferenceRuleContext {
	var p = new(AlterShardingTableReferenceRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterShardingTableReferenceRule

	return p
}

func (s *AlterShardingTableReferenceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterShardingTableReferenceRuleContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterShardingTableReferenceRuleContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *AlterShardingTableReferenceRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *AlterShardingTableReferenceRuleContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREFERENCE, 0)
}

func (s *AlterShardingTableReferenceRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *AlterShardingTableReferenceRuleContext) AllTableReferenceRuleDefinition() []ITableReferenceRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableReferenceRuleDefinitionContext)(nil)).Elem())
	var tst = make([]ITableReferenceRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableReferenceRuleDefinitionContext)
		}
	}

	return tst
}

func (s *AlterShardingTableReferenceRuleContext) TableReferenceRuleDefinition(i int) ITableReferenceRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableReferenceRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableReferenceRuleDefinitionContext)
}

func (s *AlterShardingTableReferenceRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *AlterShardingTableReferenceRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *AlterShardingTableReferenceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterShardingTableReferenceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterShardingTableReferenceRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterShardingTableReferenceRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterShardingTableReferenceRule() (localctx IAlterShardingTableReferenceRuleContext) {
	localctx = NewAlterShardingTableReferenceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RDLStatementParserRULE_alterShardingTableReferenceRule)
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
		p.SetState(168)
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(169)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(170)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(171)
		p.Match(RDLStatementParserREFERENCE)
	}
	{
		p.SetState(172)
		p.Match(RDLStatementParserRULE)
	}
	{
		p.SetState(173)
		p.TableReferenceRuleDefinition()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(174)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(175)
			p.TableReferenceRuleDefinition()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropShardingTableReferenceRuleContext is an interface to support dynamic dispatch.
type IDropShardingTableReferenceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropShardingTableReferenceRuleContext differentiates from other interfaces.
	IsDropShardingTableReferenceRuleContext()
}

type DropShardingTableReferenceRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropShardingTableReferenceRuleContext() *DropShardingTableReferenceRuleContext {
	var p = new(DropShardingTableReferenceRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropShardingTableReferenceRule
	return p
}

func (*DropShardingTableReferenceRuleContext) IsDropShardingTableReferenceRuleContext() {}

func NewDropShardingTableReferenceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropShardingTableReferenceRuleContext {
	var p = new(DropShardingTableReferenceRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropShardingTableReferenceRule

	return p
}

func (s *DropShardingTableReferenceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropShardingTableReferenceRuleContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropShardingTableReferenceRuleContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *DropShardingTableReferenceRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *DropShardingTableReferenceRuleContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserREFERENCE, 0)
}

func (s *DropShardingTableReferenceRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *DropShardingTableReferenceRuleContext) AllRuleName() []IRuleNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleNameContext)(nil)).Elem())
	var tst = make([]IRuleNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleNameContext)
		}
	}

	return tst
}

func (s *DropShardingTableReferenceRuleContext) RuleName(i int) IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *DropShardingTableReferenceRuleContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropShardingTableReferenceRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropShardingTableReferenceRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropShardingTableReferenceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropShardingTableReferenceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropShardingTableReferenceRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropShardingTableReferenceRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropShardingTableReferenceRule() (localctx IDropShardingTableReferenceRuleContext) {
	localctx = NewDropShardingTableReferenceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RDLStatementParserRULE_dropShardingTableReferenceRule)
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
		p.SetState(181)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(182)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(183)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(184)
		p.Match(RDLStatementParserREFERENCE)
	}
	{
		p.SetState(185)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(186)
			p.IfExists()
		}

	}
	{
		p.SetState(189)
		p.RuleName()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(190)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(191)
			p.RuleName()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICreateBroadcastTableRuleContext is an interface to support dynamic dispatch.
type ICreateBroadcastTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateBroadcastTableRuleContext differentiates from other interfaces.
	IsCreateBroadcastTableRuleContext()
}

type CreateBroadcastTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateBroadcastTableRuleContext() *CreateBroadcastTableRuleContext {
	var p = new(CreateBroadcastTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createBroadcastTableRule
	return p
}

func (*CreateBroadcastTableRuleContext) IsCreateBroadcastTableRuleContext() {}

func NewCreateBroadcastTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateBroadcastTableRuleContext {
	var p = new(CreateBroadcastTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createBroadcastTableRule

	return p
}

func (s *CreateBroadcastTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateBroadcastTableRuleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateBroadcastTableRuleContext) BROADCAST() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserBROADCAST, 0)
}

func (s *CreateBroadcastTableRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *CreateBroadcastTableRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *CreateBroadcastTableRuleContext) AllTableName() []ITableNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableNameContext)(nil)).Elem())
	var tst = make([]ITableNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableNameContext)
		}
	}

	return tst
}

func (s *CreateBroadcastTableRuleContext) TableName(i int) ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateBroadcastTableRuleContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateBroadcastTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *CreateBroadcastTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *CreateBroadcastTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateBroadcastTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateBroadcastTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateBroadcastTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateBroadcastTableRule() (localctx ICreateBroadcastTableRuleContext) {
	localctx = NewCreateBroadcastTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RDLStatementParserRULE_createBroadcastTableRule)
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
		p.SetState(197)
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(198)
		p.Match(RDLStatementParserBROADCAST)
	}
	{
		p.SetState(199)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(200)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(201)
			p.IfNotExists()
		}

	}
	{
		p.SetState(204)
		p.TableName()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(205)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(206)
			p.TableName()
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropBroadcastTableRuleContext is an interface to support dynamic dispatch.
type IDropBroadcastTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropBroadcastTableRuleContext differentiates from other interfaces.
	IsDropBroadcastTableRuleContext()
}

type DropBroadcastTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropBroadcastTableRuleContext() *DropBroadcastTableRuleContext {
	var p = new(DropBroadcastTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropBroadcastTableRule
	return p
}

func (*DropBroadcastTableRuleContext) IsDropBroadcastTableRuleContext() {}

func NewDropBroadcastTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropBroadcastTableRuleContext {
	var p = new(DropBroadcastTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropBroadcastTableRule

	return p
}

func (s *DropBroadcastTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropBroadcastTableRuleContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropBroadcastTableRuleContext) BROADCAST() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserBROADCAST, 0)
}

func (s *DropBroadcastTableRuleContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *DropBroadcastTableRuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRULE, 0)
}

func (s *DropBroadcastTableRuleContext) AllTableName() []ITableNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableNameContext)(nil)).Elem())
	var tst = make([]ITableNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableNameContext)
		}
	}

	return tst
}

func (s *DropBroadcastTableRuleContext) TableName(i int) ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DropBroadcastTableRuleContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropBroadcastTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropBroadcastTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropBroadcastTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropBroadcastTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropBroadcastTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropBroadcastTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropBroadcastTableRule() (localctx IDropBroadcastTableRuleContext) {
	localctx = NewDropBroadcastTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RDLStatementParserRULE_dropBroadcastTableRule)
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
		p.SetState(212)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(213)
		p.Match(RDLStatementParserBROADCAST)
	}
	{
		p.SetState(214)
		p.Match(RDLStatementParserTABLE)
	}
	{
		p.SetState(215)
		p.Match(RDLStatementParserRULE)
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(216)
			p.IfExists()
		}

	}
	{
		p.SetState(219)
		p.TableName()
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(220)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(221)
			p.TableName()
		}

		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropShardingAlgorithmContext is an interface to support dynamic dispatch.
type IDropShardingAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropShardingAlgorithmContext differentiates from other interfaces.
	IsDropShardingAlgorithmContext()
}

type DropShardingAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropShardingAlgorithmContext() *DropShardingAlgorithmContext {
	var p = new(DropShardingAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropShardingAlgorithm
	return p
}

func (*DropShardingAlgorithmContext) IsDropShardingAlgorithmContext() {}

func NewDropShardingAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropShardingAlgorithmContext {
	var p = new(DropShardingAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropShardingAlgorithm

	return p
}

func (s *DropShardingAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *DropShardingAlgorithmContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropShardingAlgorithmContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *DropShardingAlgorithmContext) ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALGORITHM, 0)
}

func (s *DropShardingAlgorithmContext) AllShardingAlgorithmName() []IShardingAlgorithmNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShardingAlgorithmNameContext)(nil)).Elem())
	var tst = make([]IShardingAlgorithmNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShardingAlgorithmNameContext)
		}
	}

	return tst
}

func (s *DropShardingAlgorithmContext) ShardingAlgorithmName(i int) IShardingAlgorithmNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingAlgorithmNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShardingAlgorithmNameContext)
}

func (s *DropShardingAlgorithmContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropShardingAlgorithmContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropShardingAlgorithmContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropShardingAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropShardingAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropShardingAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropShardingAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropShardingAlgorithm() (localctx IDropShardingAlgorithmContext) {
	localctx = NewDropShardingAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RDLStatementParserRULE_dropShardingAlgorithm)
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
		p.SetState(227)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(228)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(229)
		p.Match(RDLStatementParserALGORITHM)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(230)
			p.IfExists()
		}

	}
	{
		p.SetState(233)
		p.ShardingAlgorithmName()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(234)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(235)
			p.ShardingAlgorithmName()
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICreateDefaultShardingStrategyContext is an interface to support dynamic dispatch.
type ICreateDefaultShardingStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType returns the type token.
	GetType() antlr.Token

	// SetType sets the type token.
	SetType(antlr.Token)

	// IsCreateDefaultShardingStrategyContext differentiates from other interfaces.
	IsCreateDefaultShardingStrategyContext()
}

type CreateDefaultShardingStrategyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	typ    antlr.Token
}

func NewEmptyCreateDefaultShardingStrategyContext() *CreateDefaultShardingStrategyContext {
	var p = new(CreateDefaultShardingStrategyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_createDefaultShardingStrategy
	return p
}

func (*CreateDefaultShardingStrategyContext) IsCreateDefaultShardingStrategyContext() {}

func NewCreateDefaultShardingStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDefaultShardingStrategyContext {
	var p = new(CreateDefaultShardingStrategyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_createDefaultShardingStrategy

	return p
}

func (s *CreateDefaultShardingStrategyContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDefaultShardingStrategyContext) GetType() antlr.Token { return s.typ }

func (s *CreateDefaultShardingStrategyContext) SetType(v antlr.Token) { s.typ = v }

func (s *CreateDefaultShardingStrategyContext) CREATE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCREATE, 0)
}

func (s *CreateDefaultShardingStrategyContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDEFAULT, 0)
}

func (s *CreateDefaultShardingStrategyContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *CreateDefaultShardingStrategyContext) STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRATEGY, 0)
}

func (s *CreateDefaultShardingStrategyContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *CreateDefaultShardingStrategyContext) ShardingStrategy() IShardingStrategyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingStrategyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingStrategyContext)
}

func (s *CreateDefaultShardingStrategyContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *CreateDefaultShardingStrategyContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDATABASE, 0)
}

func (s *CreateDefaultShardingStrategyContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *CreateDefaultShardingStrategyContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateDefaultShardingStrategyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDefaultShardingStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDefaultShardingStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitCreateDefaultShardingStrategy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) CreateDefaultShardingStrategy() (localctx ICreateDefaultShardingStrategyContext) {
	localctx = NewCreateDefaultShardingStrategyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RDLStatementParserRULE_createDefaultShardingStrategy)
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
		p.Match(RDLStatementParserCREATE)
	}
	{
		p.SetState(242)
		p.Match(RDLStatementParserDEFAULT)
	}
	{
		p.SetState(243)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(244)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CreateDefaultShardingStrategyContext).typ = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == RDLStatementParserTABLE || _la == RDLStatementParserDATABASE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CreateDefaultShardingStrategyContext).typ = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(245)
		p.Match(RDLStatementParserSTRATEGY)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(246)
			p.IfNotExists()
		}

	}
	{
		p.SetState(249)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(250)
		p.ShardingStrategy()
	}
	{
		p.SetState(251)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IAlterDefaultShardingStrategyContext is an interface to support dynamic dispatch.
type IAlterDefaultShardingStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType returns the type token.
	GetType() antlr.Token

	// SetType sets the type token.
	SetType(antlr.Token)

	// IsAlterDefaultShardingStrategyContext differentiates from other interfaces.
	IsAlterDefaultShardingStrategyContext()
}

type AlterDefaultShardingStrategyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	typ    antlr.Token
}

func NewEmptyAlterDefaultShardingStrategyContext() *AlterDefaultShardingStrategyContext {
	var p = new(AlterDefaultShardingStrategyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_alterDefaultShardingStrategy
	return p
}

func (*AlterDefaultShardingStrategyContext) IsAlterDefaultShardingStrategyContext() {}

func NewAlterDefaultShardingStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterDefaultShardingStrategyContext {
	var p = new(AlterDefaultShardingStrategyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_alterDefaultShardingStrategy

	return p
}

func (s *AlterDefaultShardingStrategyContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterDefaultShardingStrategyContext) GetType() antlr.Token { return s.typ }

func (s *AlterDefaultShardingStrategyContext) SetType(v antlr.Token) { s.typ = v }

func (s *AlterDefaultShardingStrategyContext) ALTER() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALTER, 0)
}

func (s *AlterDefaultShardingStrategyContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDEFAULT, 0)
}

func (s *AlterDefaultShardingStrategyContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *AlterDefaultShardingStrategyContext) STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRATEGY, 0)
}

func (s *AlterDefaultShardingStrategyContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *AlterDefaultShardingStrategyContext) ShardingStrategy() IShardingStrategyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingStrategyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingStrategyContext)
}

func (s *AlterDefaultShardingStrategyContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *AlterDefaultShardingStrategyContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDATABASE, 0)
}

func (s *AlterDefaultShardingStrategyContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *AlterDefaultShardingStrategyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterDefaultShardingStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterDefaultShardingStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAlterDefaultShardingStrategy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AlterDefaultShardingStrategy() (localctx IAlterDefaultShardingStrategyContext) {
	localctx = NewAlterDefaultShardingStrategyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RDLStatementParserRULE_alterDefaultShardingStrategy)
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
		p.SetState(253)
		p.Match(RDLStatementParserALTER)
	}
	{
		p.SetState(254)
		p.Match(RDLStatementParserDEFAULT)
	}
	{
		p.SetState(255)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(256)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AlterDefaultShardingStrategyContext).typ = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == RDLStatementParserTABLE || _la == RDLStatementParserDATABASE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AlterDefaultShardingStrategyContext).typ = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(257)
		p.Match(RDLStatementParserSTRATEGY)
	}
	{
		p.SetState(258)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(259)
		p.ShardingStrategy()
	}
	{
		p.SetState(260)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IDropDefaultShardingStrategyContext is an interface to support dynamic dispatch.
type IDropDefaultShardingStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType returns the type token.
	GetType() antlr.Token

	// SetType sets the type token.
	SetType(antlr.Token)

	// IsDropDefaultShardingStrategyContext differentiates from other interfaces.
	IsDropDefaultShardingStrategyContext()
}

type DropDefaultShardingStrategyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	typ    antlr.Token
}

func NewEmptyDropDefaultShardingStrategyContext() *DropDefaultShardingStrategyContext {
	var p = new(DropDefaultShardingStrategyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropDefaultShardingStrategy
	return p
}

func (*DropDefaultShardingStrategyContext) IsDropDefaultShardingStrategyContext() {}

func NewDropDefaultShardingStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropDefaultShardingStrategyContext {
	var p = new(DropDefaultShardingStrategyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropDefaultShardingStrategy

	return p
}

func (s *DropDefaultShardingStrategyContext) GetParser() antlr.Parser { return s.parser }

func (s *DropDefaultShardingStrategyContext) GetType() antlr.Token { return s.typ }

func (s *DropDefaultShardingStrategyContext) SetType(v antlr.Token) { s.typ = v }

func (s *DropDefaultShardingStrategyContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropDefaultShardingStrategyContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDEFAULT, 0)
}

func (s *DropDefaultShardingStrategyContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *DropDefaultShardingStrategyContext) STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRATEGY, 0)
}

func (s *DropDefaultShardingStrategyContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDATABASE, 0)
}

func (s *DropDefaultShardingStrategyContext) TABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE, 0)
}

func (s *DropDefaultShardingStrategyContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropDefaultShardingStrategyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropDefaultShardingStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropDefaultShardingStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropDefaultShardingStrategy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropDefaultShardingStrategy() (localctx IDropDefaultShardingStrategyContext) {
	localctx = NewDropDefaultShardingStrategyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RDLStatementParserRULE_dropDefaultShardingStrategy)
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
		p.SetState(262)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(263)
		p.Match(RDLStatementParserDEFAULT)
	}
	{
		p.SetState(264)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(265)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*DropDefaultShardingStrategyContext).typ = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == RDLStatementParserTABLE || _la == RDLStatementParserDATABASE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*DropDefaultShardingStrategyContext).typ = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(266)
		p.Match(RDLStatementParserSTRATEGY)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(267)
			p.IfExists()
		}

	}

	return localctx
}

// IDropShardingKeyGeneratorContext is an interface to support dynamic dispatch.
type IDropShardingKeyGeneratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropShardingKeyGeneratorContext differentiates from other interfaces.
	IsDropShardingKeyGeneratorContext()
}

type DropShardingKeyGeneratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropShardingKeyGeneratorContext() *DropShardingKeyGeneratorContext {
	var p = new(DropShardingKeyGeneratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropShardingKeyGenerator
	return p
}

func (*DropShardingKeyGeneratorContext) IsDropShardingKeyGeneratorContext() {}

func NewDropShardingKeyGeneratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropShardingKeyGeneratorContext {
	var p = new(DropShardingKeyGeneratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropShardingKeyGenerator

	return p
}

func (s *DropShardingKeyGeneratorContext) GetParser() antlr.Parser { return s.parser }

func (s *DropShardingKeyGeneratorContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropShardingKeyGeneratorContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *DropShardingKeyGeneratorContext) KEY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserKEY, 0)
}

func (s *DropShardingKeyGeneratorContext) GENERATOR() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserGENERATOR, 0)
}

func (s *DropShardingKeyGeneratorContext) AllKeyGeneratorName() []IKeyGeneratorNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyGeneratorNameContext)(nil)).Elem())
	var tst = make([]IKeyGeneratorNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyGeneratorNameContext)
		}
	}

	return tst
}

func (s *DropShardingKeyGeneratorContext) KeyGeneratorName(i int) IKeyGeneratorNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyGeneratorNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyGeneratorNameContext)
}

func (s *DropShardingKeyGeneratorContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropShardingKeyGeneratorContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropShardingKeyGeneratorContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropShardingKeyGeneratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropShardingKeyGeneratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropShardingKeyGeneratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropShardingKeyGenerator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropShardingKeyGenerator() (localctx IDropShardingKeyGeneratorContext) {
	localctx = NewDropShardingKeyGeneratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RDLStatementParserRULE_dropShardingKeyGenerator)
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
		p.SetState(270)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(271)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(272)
		p.Match(RDLStatementParserKEY)
	}
	{
		p.SetState(273)
		p.Match(RDLStatementParserGENERATOR)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(274)
			p.IfExists()
		}

	}
	{
		p.SetState(277)
		p.KeyGeneratorName()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(278)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(279)
			p.KeyGeneratorName()
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDropShardingAuditorContext is an interface to support dynamic dispatch.
type IDropShardingAuditorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropShardingAuditorContext differentiates from other interfaces.
	IsDropShardingAuditorContext()
}

type DropShardingAuditorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropShardingAuditorContext() *DropShardingAuditorContext {
	var p = new(DropShardingAuditorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dropShardingAuditor
	return p
}

func (*DropShardingAuditorContext) IsDropShardingAuditorContext() {}

func NewDropShardingAuditorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropShardingAuditorContext {
	var p = new(DropShardingAuditorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dropShardingAuditor

	return p
}

func (s *DropShardingAuditorContext) GetParser() antlr.Parser { return s.parser }

func (s *DropShardingAuditorContext) DROP() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDROP, 0)
}

func (s *DropShardingAuditorContext) SHARDING() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING, 0)
}

func (s *DropShardingAuditorContext) AUDITOR() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserAUDITOR, 0)
}

func (s *DropShardingAuditorContext) AllAuditorName() []IAuditorNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAuditorNameContext)(nil)).Elem())
	var tst = make([]IAuditorNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAuditorNameContext)
		}
	}

	return tst
}

func (s *DropShardingAuditorContext) AuditorName(i int) IAuditorNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuditorNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAuditorNameContext)
}

func (s *DropShardingAuditorContext) IfExists() IIfExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExistsContext)
}

func (s *DropShardingAuditorContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DropShardingAuditorContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DropShardingAuditorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropShardingAuditorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropShardingAuditorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDropShardingAuditor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DropShardingAuditor() (localctx IDropShardingAuditorContext) {
	localctx = NewDropShardingAuditorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RDLStatementParserRULE_dropShardingAuditor)
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
		p.SetState(285)
		p.Match(RDLStatementParserDROP)
	}
	{
		p.SetState(286)
		p.Match(RDLStatementParserSHARDING)
	}
	{
		p.SetState(287)
		p.Match(RDLStatementParserAUDITOR)
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserIF {
		{
			p.SetState(288)
			p.IfExists()
		}

	}
	{
		p.SetState(291)
		p.AuditorName()
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(292)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(293)
			p.AuditorName()
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IShardingTableRuleDefinitionContext is an interface to support dynamic dispatch.
type IShardingTableRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingTableRuleDefinitionContext differentiates from other interfaces.
	IsShardingTableRuleDefinitionContext()
}

type ShardingTableRuleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingTableRuleDefinitionContext() *ShardingTableRuleDefinitionContext {
	var p = new(ShardingTableRuleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingTableRuleDefinition
	return p
}

func (*ShardingTableRuleDefinitionContext) IsShardingTableRuleDefinitionContext() {}

func NewShardingTableRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingTableRuleDefinitionContext {
	var p = new(ShardingTableRuleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingTableRuleDefinition

	return p
}

func (s *ShardingTableRuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingTableRuleDefinitionContext) ShardingAutoTableRule() IShardingAutoTableRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingAutoTableRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingAutoTableRuleContext)
}

func (s *ShardingTableRuleDefinitionContext) ShardingTableRule() IShardingTableRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingTableRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingTableRuleContext)
}

func (s *ShardingTableRuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingTableRuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingTableRuleDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingTableRuleDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingTableRuleDefinition() (localctx IShardingTableRuleDefinitionContext) {
	localctx = NewShardingTableRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RDLStatementParserRULE_shardingTableRuleDefinition)

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
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(299)
			p.ShardingAutoTableRule()
		}

	case 2:
		{
			p.SetState(300)
			p.ShardingTableRule()
		}

	}

	return localctx
}

// IShardingAutoTableRuleContext is an interface to support dynamic dispatch.
type IShardingAutoTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingAutoTableRuleContext differentiates from other interfaces.
	IsShardingAutoTableRuleContext()
}

type ShardingAutoTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingAutoTableRuleContext() *ShardingAutoTableRuleContext {
	var p = new(ShardingAutoTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingAutoTableRule
	return p
}

func (*ShardingAutoTableRuleContext) IsShardingAutoTableRuleContext() {}

func NewShardingAutoTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingAutoTableRuleContext {
	var p = new(ShardingAutoTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingAutoTableRule

	return p
}

func (s *ShardingAutoTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingAutoTableRuleContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *ShardingAutoTableRuleContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ShardingAutoTableRuleContext) StorageUnits() IStorageUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStorageUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStorageUnitsContext)
}

func (s *ShardingAutoTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ShardingAutoTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ShardingAutoTableRuleContext) AutoShardingColumnDefinition() IAutoShardingColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAutoShardingColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAutoShardingColumnDefinitionContext)
}

func (s *ShardingAutoTableRuleContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *ShardingAutoTableRuleContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ShardingAutoTableRuleContext) KeyGenerateDefinition() IKeyGenerateDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyGenerateDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyGenerateDefinitionContext)
}

func (s *ShardingAutoTableRuleContext) AuditDefinition() IAuditDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuditDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAuditDefinitionContext)
}

func (s *ShardingAutoTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingAutoTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingAutoTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingAutoTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingAutoTableRule() (localctx IShardingAutoTableRuleContext) {
	localctx = NewShardingAutoTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RDLStatementParserRULE_shardingAutoTableRule)
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
		p.SetState(303)
		p.TableName()
	}
	{
		p.SetState(304)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(305)
		p.StorageUnits()
	}
	{
		p.SetState(306)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(307)
		p.AutoShardingColumnDefinition()
	}
	{
		p.SetState(308)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(309)
		p.AlgorithmDefinition()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(310)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(311)
			p.KeyGenerateDefinition()
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(314)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(315)
			p.AuditDefinition()
		}

	}
	{
		p.SetState(318)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IShardingTableRuleContext is an interface to support dynamic dispatch.
type IShardingTableRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingTableRuleContext differentiates from other interfaces.
	IsShardingTableRuleContext()
}

type ShardingTableRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingTableRuleContext() *ShardingTableRuleContext {
	var p = new(ShardingTableRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingTableRule
	return p
}

func (*ShardingTableRuleContext) IsShardingTableRuleContext() {}

func NewShardingTableRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingTableRuleContext {
	var p = new(ShardingTableRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingTableRule

	return p
}

func (s *ShardingTableRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingTableRuleContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *ShardingTableRuleContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ShardingTableRuleContext) DataNodes() IDataNodesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNodesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataNodesContext)
}

func (s *ShardingTableRuleContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ShardingTableRuleContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ShardingTableRuleContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ShardingTableRuleContext) DatabaseStrategy() IDatabaseStrategyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatabaseStrategyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatabaseStrategyContext)
}

func (s *ShardingTableRuleContext) TableStrategy() ITableStrategyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableStrategyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableStrategyContext)
}

func (s *ShardingTableRuleContext) KeyGenerateDefinition() IKeyGenerateDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyGenerateDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyGenerateDefinitionContext)
}

func (s *ShardingTableRuleContext) AuditDefinition() IAuditDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuditDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAuditDefinitionContext)
}

func (s *ShardingTableRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingTableRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingTableRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingTableRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingTableRule() (localctx IShardingTableRuleContext) {
	localctx = NewShardingTableRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RDLStatementParserRULE_shardingTableRule)
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
		p.SetState(320)
		p.TableName()
	}
	{
		p.SetState(321)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(322)
		p.DataNodes()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(323)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(324)
			p.DatabaseStrategy()
		}

	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(327)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(328)
			p.TableStrategy()
		}

	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(331)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(332)
			p.KeyGenerateDefinition()
		}

	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(335)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(336)
			p.AuditDefinition()
		}

	}
	{
		p.SetState(339)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IKeyGeneratorNameContext is an interface to support dynamic dispatch.
type IKeyGeneratorNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyGeneratorNameContext differentiates from other interfaces.
	IsKeyGeneratorNameContext()
}

type KeyGeneratorNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyGeneratorNameContext() *KeyGeneratorNameContext {
	var p = new(KeyGeneratorNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_keyGeneratorName
	return p
}

func (*KeyGeneratorNameContext) IsKeyGeneratorNameContext() {}

func NewKeyGeneratorNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyGeneratorNameContext {
	var p = new(KeyGeneratorNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_keyGeneratorName

	return p
}

func (s *KeyGeneratorNameContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyGeneratorNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *KeyGeneratorNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyGeneratorNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyGeneratorNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitKeyGeneratorName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) KeyGeneratorName() (localctx IKeyGeneratorNameContext) {
	localctx = NewKeyGeneratorNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RDLStatementParserRULE_keyGeneratorName)

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
		p.SetState(341)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IAuditorDefinitionContext is an interface to support dynamic dispatch.
type IAuditorDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuditorDefinitionContext differentiates from other interfaces.
	IsAuditorDefinitionContext()
}

type AuditorDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditorDefinitionContext() *AuditorDefinitionContext {
	var p = new(AuditorDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_auditorDefinition
	return p
}

func (*AuditorDefinitionContext) IsAuditorDefinitionContext() {}

func NewAuditorDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditorDefinitionContext {
	var p = new(AuditorDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_auditorDefinition

	return p
}

func (s *AuditorDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditorDefinitionContext) AuditorName() IAuditorNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuditorNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAuditorNameContext)
}

func (s *AuditorDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *AuditorDefinitionContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *AuditorDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *AuditorDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditorDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuditorDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAuditorDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AuditorDefinition() (localctx IAuditorDefinitionContext) {
	localctx = NewAuditorDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RDLStatementParserRULE_auditorDefinition)

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
		p.SetState(343)
		p.AuditorName()
	}
	{
		p.SetState(344)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(345)
		p.AlgorithmDefinition()
	}
	{
		p.SetState(346)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IAuditorNameContext is an interface to support dynamic dispatch.
type IAuditorNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuditorNameContext differentiates from other interfaces.
	IsAuditorNameContext()
}

type AuditorNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditorNameContext() *AuditorNameContext {
	var p = new(AuditorNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_auditorName
	return p
}

func (*AuditorNameContext) IsAuditorNameContext() {}

func NewAuditorNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditorNameContext {
	var p = new(AuditorNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_auditorName

	return p
}

func (s *AuditorNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditorNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *AuditorNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditorNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuditorNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAuditorName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AuditorName() (localctx IAuditorNameContext) {
	localctx = NewAuditorNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RDLStatementParserRULE_auditorName)

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
		p.SetState(348)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IStorageUnitsContext is an interface to support dynamic dispatch.
type IStorageUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStorageUnitsContext differentiates from other interfaces.
	IsStorageUnitsContext()
}

type StorageUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStorageUnitsContext() *StorageUnitsContext {
	var p = new(StorageUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_storageUnits
	return p
}

func (*StorageUnitsContext) IsStorageUnitsContext() {}

func NewStorageUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StorageUnitsContext {
	var p = new(StorageUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_storageUnits

	return p
}

func (s *StorageUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *StorageUnitsContext) STORAGE_UNITS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTORAGE_UNITS, 0)
}

func (s *StorageUnitsContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *StorageUnitsContext) AllStorageUnit() []IStorageUnitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStorageUnitContext)(nil)).Elem())
	var tst = make([]IStorageUnitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStorageUnitContext)
		}
	}

	return tst
}

func (s *StorageUnitsContext) StorageUnit(i int) IStorageUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStorageUnitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStorageUnitContext)
}

func (s *StorageUnitsContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *StorageUnitsContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *StorageUnitsContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *StorageUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StorageUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StorageUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitStorageUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) StorageUnits() (localctx IStorageUnitsContext) {
	localctx = NewStorageUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RDLStatementParserRULE_storageUnits)
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
		p.SetState(350)
		p.Match(RDLStatementParserSTORAGE_UNITS)
	}
	{
		p.SetState(351)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(352)
		p.StorageUnit()
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(353)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(354)
			p.StorageUnit()
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(360)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IStorageUnitContext is an interface to support dynamic dispatch.
type IStorageUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStorageUnitContext differentiates from other interfaces.
	IsStorageUnitContext()
}

type StorageUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStorageUnitContext() *StorageUnitContext {
	var p = new(StorageUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_storageUnit
	return p
}

func (*StorageUnitContext) IsStorageUnitContext() {}

func NewStorageUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StorageUnitContext {
	var p = new(StorageUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_storageUnit

	return p
}

func (s *StorageUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *StorageUnitContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *StorageUnitContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *StorageUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StorageUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StorageUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitStorageUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) StorageUnit() (localctx IStorageUnitContext) {
	localctx = NewStorageUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RDLStatementParserRULE_storageUnit)
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
		p.SetState(362)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RDLStatementParserIDENTIFIER_ || _la == RDLStatementParserSTRING_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDataNodesContext is an interface to support dynamic dispatch.
type IDataNodesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataNodesContext differentiates from other interfaces.
	IsDataNodesContext()
}

type DataNodesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataNodesContext() *DataNodesContext {
	var p = new(DataNodesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dataNodes
	return p
}

func (*DataNodesContext) IsDataNodesContext() {}

func NewDataNodesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataNodesContext {
	var p = new(DataNodesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dataNodes

	return p
}

func (s *DataNodesContext) GetParser() antlr.Parser { return s.parser }

func (s *DataNodesContext) DATANODES() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDATANODES, 0)
}

func (s *DataNodesContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *DataNodesContext) AllDataNode() []IDataNodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataNodeContext)(nil)).Elem())
	var tst = make([]IDataNodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataNodeContext)
		}
	}

	return tst
}

func (s *DataNodesContext) DataNode(i int) IDataNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataNodeContext)
}

func (s *DataNodesContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *DataNodesContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *DataNodesContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *DataNodesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataNodesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataNodesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDataNodes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DataNodes() (localctx IDataNodesContext) {
	localctx = NewDataNodesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RDLStatementParserRULE_dataNodes)
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
		p.SetState(364)
		p.Match(RDLStatementParserDATANODES)
	}
	{
		p.SetState(365)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(366)
		p.DataNode()
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(367)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(368)
			p.DataNode()
		}

		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(374)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IDataNodeContext is an interface to support dynamic dispatch.
type IDataNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataNodeContext differentiates from other interfaces.
	IsDataNodeContext()
}

type DataNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataNodeContext() *DataNodeContext {
	var p = new(DataNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_dataNode
	return p
}

func (*DataNodeContext) IsDataNodeContext() {}

func NewDataNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataNodeContext {
	var p = new(DataNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_dataNode

	return p
}

func (s *DataNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataNodeContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *DataNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDataNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DataNode() (localctx IDataNodeContext) {
	localctx = NewDataNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RDLStatementParserRULE_dataNode)

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
		p.SetState(376)
		p.Match(RDLStatementParserSTRING_)
	}

	return localctx
}

// IAutoShardingColumnDefinitionContext is an interface to support dynamic dispatch.
type IAutoShardingColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAutoShardingColumnDefinitionContext differentiates from other interfaces.
	IsAutoShardingColumnDefinitionContext()
}

type AutoShardingColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAutoShardingColumnDefinitionContext() *AutoShardingColumnDefinitionContext {
	var p = new(AutoShardingColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_autoShardingColumnDefinition
	return p
}

func (*AutoShardingColumnDefinitionContext) IsAutoShardingColumnDefinitionContext() {}

func NewAutoShardingColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AutoShardingColumnDefinitionContext {
	var p = new(AutoShardingColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_autoShardingColumnDefinition

	return p
}

func (s *AutoShardingColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *AutoShardingColumnDefinitionContext) ShardingColumn() IShardingColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingColumnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingColumnContext)
}

func (s *AutoShardingColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AutoShardingColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AutoShardingColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAutoShardingColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AutoShardingColumnDefinition() (localctx IAutoShardingColumnDefinitionContext) {
	localctx = NewAutoShardingColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RDLStatementParserRULE_autoShardingColumnDefinition)

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
		p.SetState(378)
		p.ShardingColumn()
	}

	return localctx
}

// IShardingColumnDefinitionContext is an interface to support dynamic dispatch.
type IShardingColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingColumnDefinitionContext differentiates from other interfaces.
	IsShardingColumnDefinitionContext()
}

type ShardingColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingColumnDefinitionContext() *ShardingColumnDefinitionContext {
	var p = new(ShardingColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingColumnDefinition
	return p
}

func (*ShardingColumnDefinitionContext) IsShardingColumnDefinitionContext() {}

func NewShardingColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingColumnDefinitionContext {
	var p = new(ShardingColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingColumnDefinition

	return p
}

func (s *ShardingColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingColumnDefinitionContext) ShardingColumn() IShardingColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingColumnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingColumnContext)
}

func (s *ShardingColumnDefinitionContext) ShardingColumns() IShardingColumnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingColumnsContext)
}

func (s *ShardingColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingColumnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingColumnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingColumnDefinition() (localctx IShardingColumnDefinitionContext) {
	localctx = NewShardingColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RDLStatementParserRULE_shardingColumnDefinition)

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

	p.SetState(382)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSHARDING_COLUMN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.ShardingColumn()
		}

	case RDLStatementParserSHARDING_COLUMNS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.ShardingColumns()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IShardingColumnContext is an interface to support dynamic dispatch.
type IShardingColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingColumnContext differentiates from other interfaces.
	IsShardingColumnContext()
}

type ShardingColumnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingColumnContext() *ShardingColumnContext {
	var p = new(ShardingColumnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingColumn
	return p
}

func (*ShardingColumnContext) IsShardingColumnContext() {}

func NewShardingColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingColumnContext {
	var p = new(ShardingColumnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingColumn

	return p
}

func (s *ShardingColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingColumnContext) SHARDING_COLUMN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING_COLUMN, 0)
}

func (s *ShardingColumnContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *ShardingColumnContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ShardingColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingColumnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingColumn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingColumn() (localctx IShardingColumnContext) {
	localctx = NewShardingColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RDLStatementParserRULE_shardingColumn)

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
		p.SetState(384)
		p.Match(RDLStatementParserSHARDING_COLUMN)
	}
	{
		p.SetState(385)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(386)
		p.ColumnName()
	}

	return localctx
}

// IShardingColumnsContext is an interface to support dynamic dispatch.
type IShardingColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingColumnsContext differentiates from other interfaces.
	IsShardingColumnsContext()
}

type ShardingColumnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingColumnsContext() *ShardingColumnsContext {
	var p = new(ShardingColumnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingColumns
	return p
}

func (*ShardingColumnsContext) IsShardingColumnsContext() {}

func NewShardingColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingColumnsContext {
	var p = new(ShardingColumnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingColumns

	return p
}

func (s *ShardingColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingColumnsContext) SHARDING_COLUMNS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING_COLUMNS, 0)
}

func (s *ShardingColumnsContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *ShardingColumnsContext) AllColumnName() []IColumnNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnNameContext)(nil)).Elem())
	var tst = make([]IColumnNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnNameContext)
		}
	}

	return tst
}

func (s *ShardingColumnsContext) ColumnName(i int) IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ShardingColumnsContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ShardingColumnsContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ShardingColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingColumnsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingColumns(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingColumns() (localctx IShardingColumnsContext) {
	localctx = NewShardingColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RDLStatementParserRULE_shardingColumns)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(RDLStatementParserSHARDING_COLUMNS)
	}
	{
		p.SetState(389)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(390)
		p.ColumnName()
	}
	{
		p.SetState(391)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(392)
		p.ColumnName()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(393)
				p.Match(RDLStatementParserCOMMA_)
			}
			{
				p.SetState(394)
				p.ColumnName()
			}

		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
	}

	return localctx
}

// IShardingAlgorithmContext is an interface to support dynamic dispatch.
type IShardingAlgorithmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingAlgorithmContext differentiates from other interfaces.
	IsShardingAlgorithmContext()
}

type ShardingAlgorithmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingAlgorithmContext() *ShardingAlgorithmContext {
	var p = new(ShardingAlgorithmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingAlgorithm
	return p
}

func (*ShardingAlgorithmContext) IsShardingAlgorithmContext() {}

func NewShardingAlgorithmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingAlgorithmContext {
	var p = new(ShardingAlgorithmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingAlgorithm

	return p
}

func (s *ShardingAlgorithmContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingAlgorithmContext) SHARDING_ALGORITHM() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSHARDING_ALGORITHM, 0)
}

func (s *ShardingAlgorithmContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *ShardingAlgorithmContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *ShardingAlgorithmContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *ShardingAlgorithmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingAlgorithmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingAlgorithmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingAlgorithm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingAlgorithm() (localctx IShardingAlgorithmContext) {
	localctx = NewShardingAlgorithmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, RDLStatementParserRULE_shardingAlgorithm)

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
		p.SetState(400)
		p.Match(RDLStatementParserSHARDING_ALGORITHM)
	}
	{
		p.SetState(401)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(402)
		p.AlgorithmDefinition()
	}
	{
		p.SetState(403)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IShardingStrategyContext is an interface to support dynamic dispatch.
type IShardingStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingStrategyContext differentiates from other interfaces.
	IsShardingStrategyContext()
}

type ShardingStrategyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingStrategyContext() *ShardingStrategyContext {
	var p = new(ShardingStrategyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingStrategy
	return p
}

func (*ShardingStrategyContext) IsShardingStrategyContext() {}

func NewShardingStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingStrategyContext {
	var p = new(ShardingStrategyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingStrategy

	return p
}

func (s *ShardingStrategyContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingStrategyContext) TYPE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTYPE, 0)
}

func (s *ShardingStrategyContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *ShardingStrategyContext) StrategyType() IStrategyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrategyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrategyTypeContext)
}

func (s *ShardingStrategyContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *ShardingStrategyContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *ShardingStrategyContext) ShardingAlgorithm() IShardingAlgorithmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingAlgorithmContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingAlgorithmContext)
}

func (s *ShardingStrategyContext) ShardingColumnDefinition() IShardingColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingColumnDefinitionContext)
}

func (s *ShardingStrategyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingStrategy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingStrategy() (localctx IShardingStrategyContext) {
	localctx = NewShardingStrategyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, RDLStatementParserRULE_shardingStrategy)
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
		p.SetState(405)
		p.Match(RDLStatementParserTYPE)
	}
	{
		p.SetState(406)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(407)
		p.StrategyType()
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		p.SetState(410)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(408)
				p.Match(RDLStatementParserCOMMA_)
			}
			{
				p.SetState(409)
				p.ShardingColumnDefinition()
			}

		}
		{
			p.SetState(412)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(413)
			p.ShardingAlgorithm()
		}

	}

	return localctx
}

// IDatabaseStrategyContext is an interface to support dynamic dispatch.
type IDatabaseStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatabaseStrategyContext differentiates from other interfaces.
	IsDatabaseStrategyContext()
}

type DatabaseStrategyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabaseStrategyContext() *DatabaseStrategyContext {
	var p = new(DatabaseStrategyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_databaseStrategy
	return p
}

func (*DatabaseStrategyContext) IsDatabaseStrategyContext() {}

func NewDatabaseStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabaseStrategyContext {
	var p = new(DatabaseStrategyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_databaseStrategy

	return p
}

func (s *DatabaseStrategyContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabaseStrategyContext) DATABASE_STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDATABASE_STRATEGY, 0)
}

func (s *DatabaseStrategyContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *DatabaseStrategyContext) ShardingStrategy() IShardingStrategyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingStrategyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingStrategyContext)
}

func (s *DatabaseStrategyContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *DatabaseStrategyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabaseStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabaseStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitDatabaseStrategy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) DatabaseStrategy() (localctx IDatabaseStrategyContext) {
	localctx = NewDatabaseStrategyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, RDLStatementParserRULE_databaseStrategy)

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
		p.SetState(416)
		p.Match(RDLStatementParserDATABASE_STRATEGY)
	}
	{
		p.SetState(417)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(418)
		p.ShardingStrategy()
	}
	{
		p.SetState(419)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// ITableStrategyContext is an interface to support dynamic dispatch.
type ITableStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableStrategyContext differentiates from other interfaces.
	IsTableStrategyContext()
}

type TableStrategyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableStrategyContext() *TableStrategyContext {
	var p = new(TableStrategyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_tableStrategy
	return p
}

func (*TableStrategyContext) IsTableStrategyContext() {}

func NewTableStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableStrategyContext {
	var p = new(TableStrategyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_tableStrategy

	return p
}

func (s *TableStrategyContext) GetParser() antlr.Parser { return s.parser }

func (s *TableStrategyContext) TABLE_STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTABLE_STRATEGY, 0)
}

func (s *TableStrategyContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *TableStrategyContext) ShardingStrategy() IShardingStrategyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShardingStrategyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShardingStrategyContext)
}

func (s *TableStrategyContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *TableStrategyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitTableStrategy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) TableStrategy() (localctx ITableStrategyContext) {
	localctx = NewTableStrategyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, RDLStatementParserRULE_tableStrategy)

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
		p.SetState(421)
		p.Match(RDLStatementParserTABLE_STRATEGY)
	}
	{
		p.SetState(422)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(423)
		p.ShardingStrategy()
	}
	{
		p.SetState(424)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IKeyGenerateDefinitionContext is an interface to support dynamic dispatch.
type IKeyGenerateDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyGenerateDefinitionContext differentiates from other interfaces.
	IsKeyGenerateDefinitionContext()
}

type KeyGenerateDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyGenerateDefinitionContext() *KeyGenerateDefinitionContext {
	var p = new(KeyGenerateDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_keyGenerateDefinition
	return p
}

func (*KeyGenerateDefinitionContext) IsKeyGenerateDefinitionContext() {}

func NewKeyGenerateDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyGenerateDefinitionContext {
	var p = new(KeyGenerateDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_keyGenerateDefinition

	return p
}

func (s *KeyGenerateDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyGenerateDefinitionContext) KEY_GENERATE_STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserKEY_GENERATE_STRATEGY, 0)
}

func (s *KeyGenerateDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *KeyGenerateDefinitionContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOLUMN, 0)
}

func (s *KeyGenerateDefinitionContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *KeyGenerateDefinitionContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *KeyGenerateDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *KeyGenerateDefinitionContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *KeyGenerateDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *KeyGenerateDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyGenerateDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyGenerateDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitKeyGenerateDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) KeyGenerateDefinition() (localctx IKeyGenerateDefinitionContext) {
	localctx = NewKeyGenerateDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, RDLStatementParserRULE_keyGenerateDefinition)

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
		p.SetState(426)
		p.Match(RDLStatementParserKEY_GENERATE_STRATEGY)
	}
	{
		p.SetState(427)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(428)
		p.Match(RDLStatementParserCOLUMN)
	}
	{
		p.SetState(429)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(430)
		p.ColumnName()
	}
	{
		p.SetState(431)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(432)
		p.AlgorithmDefinition()
	}
	{
		p.SetState(433)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IAuditDefinitionContext is an interface to support dynamic dispatch.
type IAuditDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuditDefinitionContext differentiates from other interfaces.
	IsAuditDefinitionContext()
}

type AuditDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditDefinitionContext() *AuditDefinitionContext {
	var p = new(AuditDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_auditDefinition
	return p
}

func (*AuditDefinitionContext) IsAuditDefinitionContext() {}

func NewAuditDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditDefinitionContext {
	var p = new(AuditDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_auditDefinition

	return p
}

func (s *AuditDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditDefinitionContext) AUDIT_STRATEGY() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserAUDIT_STRATEGY, 0)
}

func (s *AuditDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *AuditDefinitionContext) MultiAuditDefinition() IMultiAuditDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiAuditDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiAuditDefinitionContext)
}

func (s *AuditDefinitionContext) COMMA_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, 0)
}

func (s *AuditDefinitionContext) ALLOW_HINT_DISABLE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserALLOW_HINT_DISABLE, 0)
}

func (s *AuditDefinitionContext) EQ_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserEQ_, 0)
}

func (s *AuditDefinitionContext) AuditAllowHintDisable() IAuditAllowHintDisableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuditAllowHintDisableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAuditAllowHintDisableContext)
}

func (s *AuditDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *AuditDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuditDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAuditDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AuditDefinition() (localctx IAuditDefinitionContext) {
	localctx = NewAuditDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, RDLStatementParserRULE_auditDefinition)

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
		p.SetState(435)
		p.Match(RDLStatementParserAUDIT_STRATEGY)
	}
	{
		p.SetState(436)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(437)
		p.MultiAuditDefinition()
	}
	{
		p.SetState(438)
		p.Match(RDLStatementParserCOMMA_)
	}
	{
		p.SetState(439)
		p.Match(RDLStatementParserALLOW_HINT_DISABLE)
	}
	{
		p.SetState(440)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(441)
		p.AuditAllowHintDisable()
	}
	{
		p.SetState(442)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IMultiAuditDefinitionContext is an interface to support dynamic dispatch.
type IMultiAuditDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiAuditDefinitionContext differentiates from other interfaces.
	IsMultiAuditDefinitionContext()
}

type MultiAuditDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiAuditDefinitionContext() *MultiAuditDefinitionContext {
	var p = new(MultiAuditDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_multiAuditDefinition
	return p
}

func (*MultiAuditDefinitionContext) IsMultiAuditDefinitionContext() {}

func NewMultiAuditDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiAuditDefinitionContext {
	var p = new(MultiAuditDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_multiAuditDefinition

	return p
}

func (s *MultiAuditDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiAuditDefinitionContext) AllSingleAuditDefinition() []ISingleAuditDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleAuditDefinitionContext)(nil)).Elem())
	var tst = make([]ISingleAuditDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleAuditDefinitionContext)
		}
	}

	return tst
}

func (s *MultiAuditDefinitionContext) SingleAuditDefinition(i int) ISingleAuditDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleAuditDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleAuditDefinitionContext)
}

func (s *MultiAuditDefinitionContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *MultiAuditDefinitionContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *MultiAuditDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiAuditDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiAuditDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitMultiAuditDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) MultiAuditDefinition() (localctx IMultiAuditDefinitionContext) {
	localctx = NewMultiAuditDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, RDLStatementParserRULE_multiAuditDefinition)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.SingleAuditDefinition()
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(445)
				p.Match(RDLStatementParserCOMMA_)
			}
			{
				p.SetState(446)
				p.SingleAuditDefinition()
			}

		}
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// ISingleAuditDefinitionContext is an interface to support dynamic dispatch.
type ISingleAuditDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleAuditDefinitionContext differentiates from other interfaces.
	IsSingleAuditDefinitionContext()
}

type SingleAuditDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleAuditDefinitionContext() *SingleAuditDefinitionContext {
	var p = new(SingleAuditDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_singleAuditDefinition
	return p
}

func (*SingleAuditDefinitionContext) IsSingleAuditDefinitionContext() {}

func NewSingleAuditDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleAuditDefinitionContext {
	var p = new(SingleAuditDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_singleAuditDefinition

	return p
}

func (s *SingleAuditDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleAuditDefinitionContext) AlgorithmDefinition() IAlgorithmDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlgorithmDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *SingleAuditDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleAuditDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleAuditDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitSingleAuditDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) SingleAuditDefinition() (localctx ISingleAuditDefinitionContext) {
	localctx = NewSingleAuditDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, RDLStatementParserRULE_singleAuditDefinition)

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
		p.SetState(452)
		p.AlgorithmDefinition()
	}

	return localctx
}

// IAuditAllowHintDisableContext is an interface to support dynamic dispatch.
type IAuditAllowHintDisableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuditAllowHintDisableContext differentiates from other interfaces.
	IsAuditAllowHintDisableContext()
}

type AuditAllowHintDisableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditAllowHintDisableContext() *AuditAllowHintDisableContext {
	var p = new(AuditAllowHintDisableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_auditAllowHintDisable
	return p
}

func (*AuditAllowHintDisableContext) IsAuditAllowHintDisableContext() {}

func NewAuditAllowHintDisableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditAllowHintDisableContext {
	var p = new(AuditAllowHintDisableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_auditAllowHintDisable

	return p
}

func (s *AuditAllowHintDisableContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditAllowHintDisableContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserTRUE, 0)
}

func (s *AuditAllowHintDisableContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserFALSE, 0)
}

func (s *AuditAllowHintDisableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditAllowHintDisableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuditAllowHintDisableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitAuditAllowHintDisable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) AuditAllowHintDisable() (localctx IAuditAllowHintDisableContext) {
	localctx = NewAuditAllowHintDisableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, RDLStatementParserRULE_auditAllowHintDisable)
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
		p.SetState(454)
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
	p.EnterRule(localctx, 74, RDLStatementParserRULE_columnName)

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
		p.SetState(456)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// ITableReferenceRuleDefinitionContext is an interface to support dynamic dispatch.
type ITableReferenceRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableReferenceRuleDefinitionContext differentiates from other interfaces.
	IsTableReferenceRuleDefinitionContext()
}

type TableReferenceRuleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableReferenceRuleDefinitionContext() *TableReferenceRuleDefinitionContext {
	var p = new(TableReferenceRuleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_tableReferenceRuleDefinition
	return p
}

func (*TableReferenceRuleDefinitionContext) IsTableReferenceRuleDefinitionContext() {}

func NewTableReferenceRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableReferenceRuleDefinitionContext {
	var p = new(TableReferenceRuleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_tableReferenceRuleDefinition

	return p
}

func (s *TableReferenceRuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TableReferenceRuleDefinitionContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *TableReferenceRuleDefinitionContext) LP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserLP_, 0)
}

func (s *TableReferenceRuleDefinitionContext) AllTableName() []ITableNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableNameContext)(nil)).Elem())
	var tst = make([]ITableNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableNameContext)
		}
	}

	return tst
}

func (s *TableReferenceRuleDefinitionContext) TableName(i int) ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *TableReferenceRuleDefinitionContext) RP_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserRP_, 0)
}

func (s *TableReferenceRuleDefinitionContext) AllCOMMA_() []antlr.TerminalNode {
	return s.GetTokens(RDLStatementParserCOMMA_)
}

func (s *TableReferenceRuleDefinitionContext) COMMA_(i int) antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMMA_, i)
}

func (s *TableReferenceRuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableReferenceRuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableReferenceRuleDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitTableReferenceRuleDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) TableReferenceRuleDefinition() (localctx ITableReferenceRuleDefinitionContext) {
	localctx = NewTableReferenceRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, RDLStatementParserRULE_tableReferenceRuleDefinition)
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
		p.SetState(458)
		p.RuleName()
	}
	{
		p.SetState(459)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(460)
		p.TableName()
	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(461)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(462)
			p.TableName()
		}

		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(468)
		p.Match(RDLStatementParserRP_)
	}

	return localctx
}

// IStrategyTypeContext is an interface to support dynamic dispatch.
type IStrategyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrategyTypeContext differentiates from other interfaces.
	IsStrategyTypeContext()
}

type StrategyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrategyTypeContext() *StrategyTypeContext {
	var p = new(StrategyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_strategyType
	return p
}

func (*StrategyTypeContext) IsStrategyTypeContext() {}

func NewStrategyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrategyTypeContext {
	var p = new(StrategyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_strategyType

	return p
}

func (s *StrategyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StrategyTypeContext) STRING_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTRING_, 0)
}

func (s *StrategyTypeContext) BuildInStrategyType() IBuildInStrategyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildInStrategyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildInStrategyTypeContext)
}

func (s *StrategyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrategyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrategyTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitStrategyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) StrategyType() (localctx IStrategyTypeContext) {
	localctx = NewStrategyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, RDLStatementParserRULE_strategyType)

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

	p.SetState(472)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(470)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserHINT, RDLStatementParserSTANDARD, RDLStatementParserCOMPLEX, RDLStatementParserNONE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(471)
			p.BuildInStrategyType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuildInStrategyTypeContext is an interface to support dynamic dispatch.
type IBuildInStrategyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildInStrategyTypeContext differentiates from other interfaces.
	IsBuildInStrategyTypeContext()
}

type BuildInStrategyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildInStrategyTypeContext() *BuildInStrategyTypeContext {
	var p = new(BuildInStrategyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildInStrategyType
	return p
}

func (*BuildInStrategyTypeContext) IsBuildInStrategyTypeContext() {}

func NewBuildInStrategyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildInStrategyTypeContext {
	var p = new(BuildInStrategyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildInStrategyType

	return p
}

func (s *BuildInStrategyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildInStrategyTypeContext) STANDARD() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSTANDARD, 0)
}

func (s *BuildInStrategyTypeContext) COMPLEX() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMPLEX, 0)
}

func (s *BuildInStrategyTypeContext) HINT() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserHINT, 0)
}

func (s *BuildInStrategyTypeContext) NONE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserNONE, 0)
}

func (s *BuildInStrategyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInStrategyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildInStrategyTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildInStrategyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildInStrategyType() (localctx IBuildInStrategyTypeContext) {
	localctx = NewBuildInStrategyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, RDLStatementParserRULE_buildInStrategyType)
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
		p.SetState(474)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RDLStatementParserHINT || (((_la-110)&-(0x1f+1)) == 0 && ((1<<uint((_la-110)))&((1<<(RDLStatementParserSTANDARD-110))|(1<<(RDLStatementParserCOMPLEX-110))|(1<<(RDLStatementParserNONE-110)))) != 0)) {
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
	p.EnterRule(localctx, 82, RDLStatementParserRULE_ifExists)

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
		p.SetState(476)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(477)
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
	p.EnterRule(localctx, 84, RDLStatementParserRULE_ifNotExists)

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
		p.SetState(479)
		p.Match(RDLStatementParserIF)
	}
	{
		p.SetState(480)
		p.Match(RDLStatementParserNOT)
	}
	{
		p.SetState(481)
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
	p.EnterRule(localctx, 86, RDLStatementParserRULE_literal)
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

	p.SetState(490)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserMINUS_, RDLStatementParserINT_:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RDLStatementParserMINUS_ {
			{
				p.SetState(484)
				p.Match(RDLStatementParserMINUS_)
			}

		}
		{
			p.SetState(487)
			p.Match(RDLStatementParserINT_)
		}

	case RDLStatementParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(488)
			p.Match(RDLStatementParserTRUE)
		}

	case RDLStatementParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(489)
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
	p.EnterRule(localctx, 88, RDLStatementParserRULE_algorithmDefinition)
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
		p.SetState(492)
		p.Match(RDLStatementParserTYPE)
	}
	{
		p.SetState(493)
		p.Match(RDLStatementParserLP_)
	}
	{
		p.SetState(494)
		p.Match(RDLStatementParserNAME)
	}
	{
		p.SetState(495)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(496)
		p.AlgorithmTypeName()
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(497)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(498)
			p.PropertiesDefinition()
		}

	}
	{
		p.SetState(501)
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

func (s *AlgorithmTypeNameContext) BuildInShardingAlgorithmType() IBuildInShardingAlgorithmTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildInShardingAlgorithmTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildInShardingAlgorithmTypeContext)
}

func (s *AlgorithmTypeNameContext) BuildInKeyGenerateAlgorithmType() IBuildInKeyGenerateAlgorithmTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildInKeyGenerateAlgorithmTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildInKeyGenerateAlgorithmTypeContext)
}

func (s *AlgorithmTypeNameContext) BuildInShardingAuditAlgorithmType() IBuildInShardingAuditAlgorithmTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuildInShardingAuditAlgorithmTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuildInShardingAuditAlgorithmTypeContext)
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
	p.EnterRule(localctx, 90, RDLStatementParserRULE_algorithmTypeName)

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

	p.SetState(507)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RDLStatementParserSTRING_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(503)
			p.Match(RDLStatementParserSTRING_)
		}

	case RDLStatementParserMOD, RDLStatementParserCOSID_MOD, RDLStatementParserHASH_MOD, RDLStatementParserVOLUME_RANGE, RDLStatementParserBOUNDARY_RANGE, RDLStatementParserAUTO_INTERVAL, RDLStatementParserINLINE, RDLStatementParserINTERVAL, RDLStatementParserCOSID_INTERVAL, RDLStatementParserCOSID_INTERVAL_SNOWFLAKE, RDLStatementParserCOMPLEX_INLINE, RDLStatementParserHINT_INLINE, RDLStatementParserCLASS_BASED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(504)
			p.BuildInShardingAlgorithmType()
		}

	case RDLStatementParserSNOWFLAKE, RDLStatementParserNANOID, RDLStatementParserUUID, RDLStatementParserCOSID, RDLStatementParserCOSID_SNOWFLAKE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(505)
			p.BuildInKeyGenerateAlgorithmType()
		}

	case RDLStatementParserDML_SHARDING_CONDITIONS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(506)
			p.BuildInShardingAuditAlgorithmType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuildInShardingAlgorithmTypeContext is an interface to support dynamic dispatch.
type IBuildInShardingAlgorithmTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildInShardingAlgorithmTypeContext differentiates from other interfaces.
	IsBuildInShardingAlgorithmTypeContext()
}

type BuildInShardingAlgorithmTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildInShardingAlgorithmTypeContext() *BuildInShardingAlgorithmTypeContext {
	var p = new(BuildInShardingAlgorithmTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildInShardingAlgorithmType
	return p
}

func (*BuildInShardingAlgorithmTypeContext) IsBuildInShardingAlgorithmTypeContext() {}

func NewBuildInShardingAlgorithmTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildInShardingAlgorithmTypeContext {
	var p = new(BuildInShardingAlgorithmTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildInShardingAlgorithmType

	return p
}

func (s *BuildInShardingAlgorithmTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildInShardingAlgorithmTypeContext) MOD() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserMOD, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) HASH_MOD() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserHASH_MOD, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) VOLUME_RANGE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserVOLUME_RANGE, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) BOUNDARY_RANGE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserBOUNDARY_RANGE, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) AUTO_INTERVAL() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserAUTO_INTERVAL, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) INLINE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserINLINE, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) INTERVAL() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserINTERVAL, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) COSID_MOD() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOSID_MOD, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) COSID_INTERVAL() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOSID_INTERVAL, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) COSID_INTERVAL_SNOWFLAKE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOSID_INTERVAL_SNOWFLAKE, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) COMPLEX_INLINE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOMPLEX_INLINE, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) HINT_INLINE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserHINT_INLINE, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) CLASS_BASED() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCLASS_BASED, 0)
}

func (s *BuildInShardingAlgorithmTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInShardingAlgorithmTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildInShardingAlgorithmTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildInShardingAlgorithmType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildInShardingAlgorithmType() (localctx IBuildInShardingAlgorithmTypeContext) {
	localctx = NewBuildInShardingAlgorithmTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, RDLStatementParserRULE_buildInShardingAlgorithmType)
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
		p.SetState(509)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-92)&-(0x1f+1)) == 0 && ((1<<uint((_la-92)))&((1<<(RDLStatementParserMOD-92))|(1<<(RDLStatementParserCOSID_MOD-92))|(1<<(RDLStatementParserHASH_MOD-92))|(1<<(RDLStatementParserVOLUME_RANGE-92))|(1<<(RDLStatementParserBOUNDARY_RANGE-92))|(1<<(RDLStatementParserAUTO_INTERVAL-92))|(1<<(RDLStatementParserINLINE-92))|(1<<(RDLStatementParserINTERVAL-92))|(1<<(RDLStatementParserCOSID_INTERVAL-92))|(1<<(RDLStatementParserCOSID_INTERVAL_SNOWFLAKE-92))|(1<<(RDLStatementParserCOMPLEX_INLINE-92))|(1<<(RDLStatementParserHINT_INLINE-92))|(1<<(RDLStatementParserCLASS_BASED-92)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBuildInKeyGenerateAlgorithmTypeContext is an interface to support dynamic dispatch.
type IBuildInKeyGenerateAlgorithmTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildInKeyGenerateAlgorithmTypeContext differentiates from other interfaces.
	IsBuildInKeyGenerateAlgorithmTypeContext()
}

type BuildInKeyGenerateAlgorithmTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildInKeyGenerateAlgorithmTypeContext() *BuildInKeyGenerateAlgorithmTypeContext {
	var p = new(BuildInKeyGenerateAlgorithmTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildInKeyGenerateAlgorithmType
	return p
}

func (*BuildInKeyGenerateAlgorithmTypeContext) IsBuildInKeyGenerateAlgorithmTypeContext() {}

func NewBuildInKeyGenerateAlgorithmTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildInKeyGenerateAlgorithmTypeContext {
	var p = new(BuildInKeyGenerateAlgorithmTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildInKeyGenerateAlgorithmType

	return p
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildInKeyGenerateAlgorithmTypeContext) SNOWFLAKE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserSNOWFLAKE, 0)
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) NANOID() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserNANOID, 0)
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) UUID() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserUUID, 0)
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) COSID() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOSID, 0)
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) COSID_SNOWFLAKE() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserCOSID_SNOWFLAKE, 0)
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildInKeyGenerateAlgorithmTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildInKeyGenerateAlgorithmType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildInKeyGenerateAlgorithmType() (localctx IBuildInKeyGenerateAlgorithmTypeContext) {
	localctx = NewBuildInKeyGenerateAlgorithmTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, RDLStatementParserRULE_buildInKeyGenerateAlgorithmType)
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
		p.SetState(511)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(RDLStatementParserSNOWFLAKE-105))|(1<<(RDLStatementParserNANOID-105))|(1<<(RDLStatementParserUUID-105))|(1<<(RDLStatementParserCOSID-105))|(1<<(RDLStatementParserCOSID_SNOWFLAKE-105)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBuildInShardingAuditAlgorithmTypeContext is an interface to support dynamic dispatch.
type IBuildInShardingAuditAlgorithmTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuildInShardingAuditAlgorithmTypeContext differentiates from other interfaces.
	IsBuildInShardingAuditAlgorithmTypeContext()
}

type BuildInShardingAuditAlgorithmTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildInShardingAuditAlgorithmTypeContext() *BuildInShardingAuditAlgorithmTypeContext {
	var p = new(BuildInShardingAuditAlgorithmTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_buildInShardingAuditAlgorithmType
	return p
}

func (*BuildInShardingAuditAlgorithmTypeContext) IsBuildInShardingAuditAlgorithmTypeContext() {}

func NewBuildInShardingAuditAlgorithmTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildInShardingAuditAlgorithmTypeContext {
	var p = new(BuildInShardingAuditAlgorithmTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_buildInShardingAuditAlgorithmType

	return p
}

func (s *BuildInShardingAuditAlgorithmTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildInShardingAuditAlgorithmTypeContext) DML_SHARDING_CONDITIONS() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserDML_SHARDING_CONDITIONS, 0)
}

func (s *BuildInShardingAuditAlgorithmTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildInShardingAuditAlgorithmTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildInShardingAuditAlgorithmTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitBuildInShardingAuditAlgorithmType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) BuildInShardingAuditAlgorithmType() (localctx IBuildInShardingAuditAlgorithmTypeContext) {
	localctx = NewBuildInShardingAuditAlgorithmTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, RDLStatementParserRULE_buildInShardingAuditAlgorithmType)

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
		p.SetState(513)
		p.Match(RDLStatementParserDML_SHARDING_CONDITIONS)
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
	p.EnterRule(localctx, 98, RDLStatementParserRULE_propertiesDefinition)
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
		p.SetState(515)
		p.Match(RDLStatementParserPROPERTIES)
	}
	{
		p.SetState(516)
		p.Match(RDLStatementParserLP_)
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RDLStatementParserSTRING_ {
		{
			p.SetState(517)
			p.Properties()
		}

	}
	{
		p.SetState(520)
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
	p.EnterRule(localctx, 100, RDLStatementParserRULE_properties)
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
		p.SetState(522)
		p.Property()
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RDLStatementParserCOMMA_ {
		{
			p.SetState(523)
			p.Match(RDLStatementParserCOMMA_)
		}
		{
			p.SetState(524)
			p.Property()
		}

		p.SetState(529)
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
	p.EnterRule(localctx, 102, RDLStatementParserRULE_property)

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
		p.SetState(530)

		var _m = p.Match(RDLStatementParserSTRING_)

		localctx.(*PropertyContext).key = _m
	}
	{
		p.SetState(531)
		p.Match(RDLStatementParserEQ_)
	}
	{
		p.SetState(532)

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
	p.EnterRule(localctx, 104, RDLStatementParserRULE_tableName)

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
		p.SetState(534)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}

// IShardingAlgorithmNameContext is an interface to support dynamic dispatch.
type IShardingAlgorithmNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShardingAlgorithmNameContext differentiates from other interfaces.
	IsShardingAlgorithmNameContext()
}

type ShardingAlgorithmNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShardingAlgorithmNameContext() *ShardingAlgorithmNameContext {
	var p = new(ShardingAlgorithmNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RDLStatementParserRULE_shardingAlgorithmName
	return p
}

func (*ShardingAlgorithmNameContext) IsShardingAlgorithmNameContext() {}

func NewShardingAlgorithmNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShardingAlgorithmNameContext {
	var p = new(ShardingAlgorithmNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RDLStatementParserRULE_shardingAlgorithmName

	return p
}

func (s *ShardingAlgorithmNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ShardingAlgorithmNameContext) IDENTIFIER_() antlr.TerminalNode {
	return s.GetToken(RDLStatementParserIDENTIFIER_, 0)
}

func (s *ShardingAlgorithmNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShardingAlgorithmNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShardingAlgorithmNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RDLStatementVisitor:
		return t.VisitShardingAlgorithmName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RDLStatementParser) ShardingAlgorithmName() (localctx IShardingAlgorithmNameContext) {
	localctx = NewShardingAlgorithmNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, RDLStatementParserRULE_shardingAlgorithmName)

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
		p.SetState(536)
		p.Match(RDLStatementParserIDENTIFIER_)
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
	p.EnterRule(localctx, 108, RDLStatementParserRULE_ruleName)

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
		p.SetState(538)
		p.Match(RDLStatementParserIDENTIFIER_)
	}

	return localctx
}
