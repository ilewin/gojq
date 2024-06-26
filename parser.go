// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package gojq

import __yyfmt__ "fmt"

//line parser.go.y:2

func reverseFuncDef(xs []*FuncDef) []*FuncDef {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs
}

func prependFuncDef(xs []*FuncDef, x *FuncDef) []*FuncDef {
	xs = append(xs, nil)
	copy(xs[1:], xs)
	xs[0] = x
	return xs
}

//line parser.go.y:19
type yySymType struct {
	yys      int
	value    any
	token    string
	operator Operator
}

const tokAltOp = 57346
const tokUpdateOp = 57347
const tokDestAltOp = 57348
const tokCompareOp = 57349
const tokOrOp = 57350
const tokAndOp = 57351
const tokModule = 57352
const tokImport = 57353
const tokInclude = 57354
const tokDef = 57355
const tokAs = 57356
const tokLabel = 57357
const tokBreak = 57358
const tokNull = 57359
const tokTrue = 57360
const tokFalse = 57361
const tokIf = 57362
const tokThen = 57363
const tokElif = 57364
const tokElse = 57365
const tokEnd = 57366
const tokTry = 57367
const tokCatch = 57368
const tokReduce = 57369
const tokForeach = 57370
const tokIdent = 57371
const tokVariable = 57372
const tokModuleIdent = 57373
const tokModuleVariable = 57374
const tokRecurse = 57375
const tokIndex = 57376
const tokNumber = 57377
const tokFormat = 57378
const tokString = 57379
const tokStringStart = 57380
const tokStringQuery = 57381
const tokStringEnd = 57382
const tokInvalid = 57383
const tokInvalidEscapeSequence = 57384
const tokUnterminatedString = 57385
const tokFuncDefPost = 57386
const tokTermPost = 57387
const tokEmptyCatch = 57388

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokAltOp",
	"tokUpdateOp",
	"tokDestAltOp",
	"tokCompareOp",
	"tokOrOp",
	"tokAndOp",
	"tokModule",
	"tokImport",
	"tokInclude",
	"tokDef",
	"tokAs",
	"tokLabel",
	"tokBreak",
	"tokNull",
	"tokTrue",
	"tokFalse",
	"tokIf",
	"tokThen",
	"tokElif",
	"tokElse",
	"tokEnd",
	"tokTry",
	"tokCatch",
	"tokReduce",
	"tokForeach",
	"tokIdent",
	"tokVariable",
	"tokModuleIdent",
	"tokModuleVariable",
	"tokRecurse",
	"tokIndex",
	"tokNumber",
	"tokFormat",
	"tokString",
	"tokStringStart",
	"tokStringQuery",
	"tokStringEnd",
	"tokInvalid",
	"tokInvalidEscapeSequence",
	"tokUnterminatedString",
	"tokFuncDefPost",
	"tokTermPost",
	"'|'",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"'?'",
	"tokEmptyCatch",
	"'['",
	"';'",
	"':'",
	"'('",
	"')'",
	"']'",
	"'{'",
	"'}'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:671

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 129,
	5, 0,
	-2, 26,
	-1, 132,
	7, 0,
	-2, 29,
	-1, 197,
	58, 113,
	-2, 48,
}

const yyPrivate = 57344

const yyLast = 1055

var yyAct = [...]int16{
	87, 142, 173, 184, 209, 10, 193, 174, 102, 139,
	14, 5, 107, 47, 101, 227, 143, 90, 245, 49,
	157, 122, 6, 178, 179, 180, 242, 109, 241, 31,
	225, 226, 244, 114, 115, 96, 156, 121, 119, 111,
	112, 176, 144, 177, 224, 116, 117, 97, 145, 152,
	153, 178, 179, 180, 236, 103, 206, 235, 263, 205,
	113, 98, 181, 253, 238, 239, 89, 218, 6, 176,
	229, 177, 228, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 92, 91, 159, 93, 158,
	181, 166, 246, 140, 149, 125, 6, 74, 75, 155,
	78, 76, 77, 124, 123, 160, 42, 43, 187, 44,
	146, 89, 151, 85, 74, 75, 84, 78, 76, 77,
	118, 169, 147, 170, 168, 93, 94, 182, 183, 265,
	92, 91, 95, 93, 189, 49, 172, 42, 43, 72,
	73, 79, 80, 81, 82, 83, 191, 81, 82, 83,
	271, 200, 201, 202, 194, 164, 165, 204, 79, 80,
	81, 82, 83, 213, 211, 214, 88, 215, 216, 210,
	210, 217, 208, 78, 188, 195, 89, 138, 74, 75,
	3, 78, 76, 77, 220, 110, 89, 103, 24, 222,
	13, 223, 140, 23, 221, 92, 91, 230, 93, 13,
	232, 185, 186, 258, 259, 92, 91, 237, 93, 79,
	80, 81, 82, 83, 79, 80, 81, 82, 83, 9,
	72, 73, 79, 80, 81, 82, 83, 247, 219, 175,
	249, 250, 199, 86, 248, 198, 194, 46, 100, 163,
	254, 257, 260, 261, 256, 240, 154, 120, 262, 255,
	210, 192, 190, 141, 264, 203, 7, 195, 8, 266,
	267, 4, 78, 76, 77, 270, 2, 1, 0, 273,
	274, 74, 75, 275, 78, 76, 77, 0, 0, 279,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 105, 106, 79, 80, 81, 82, 83, 0, 42,
	43, 0, 0, 72, 73, 79, 80, 81, 82, 83,
	0, 0, 16, 0, 15, 36, 20, 21, 22, 32,
	0, 104, 0, 0, 33, 207, 34, 35, 38, 40,
	39, 41, 18, 19, 27, 30, 42, 43, 0, 0,
	0, 0, 0, 0, 78, 0, 77, 28, 29, 0,
	0, 0, 17, 0, 0, 26, 0, 150, 37, 0,
	148, 25, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 105, 106, 79, 80, 81, 82, 83,
	0, 42, 43, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 11, 12, 16, 0, 15, 36, 20, 21,
	22, 32, 0, 104, 0, 0, 33, 99, 34, 35,
	38, 40, 39, 41, 18, 19, 27, 30, 42, 43,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 28,
	29, 0, 0, 0, 17, 0, 0, 26, 0, 0,
	37, 0, 16, 25, 15, 36, 20, 21, 22, 32,
	0, 0, 0, 0, 33, 0, 34, 35, 38, 40,
	39, 41, 18, 19, 27, 30, 42, 43, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 28, 29, 0,
	0, 0, 17, 0, 0, 26, 0, 0, 37, 0,
	231, 25, 16, 0, 15, 36, 20, 21, 22, 32,
	0, 0, 0, 0, 33, 0, 34, 35, 38, 40,
	39, 41, 18, 19, 27, 30, 42, 43, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 28, 29, 0,
	0, 0, 17, 0, 0, 26, 0, 0, 37, 0,
	108, 25, 16, 0, 15, 36, 20, 21, 22, 32,
	0, 0, 0, 0, 33, 0, 34, 35, 38, 40,
	39, 41, 18, 19, 27, 30, 42, 43, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 28, 29, 0,
	0, 0, 17, 0, 0, 26, 0, 0, 37, 0,
	0, 25, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 48, 0, 0, 0, 0, 0, 0,
	0, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 48, 0, 0, 0, 171, 0, 0,
	0, 50, 36, 20, 21, 22, 32, 0, 0, 0,
	0, 33, 0, 34, 35, 38, 40, 39, 41, 18,
	19, 27, 30, 42, 43, 0, 0, 45, 0, 0,
	0, 0, 0, 0, 28, 29, 0, 0, 0, 17,
	0, 0, 26, 0, 0, 37, 0, 0, 25, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	105, 197, 74, 75, 0, 78, 76, 77, 42, 43,
	0, 0, 0, 0, 0, 0, 0, 74, 75, 0,
	78, 76, 77, 0, 74, 75, 0, 78, 76, 77,
	196, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 73, 79, 80, 81, 82,
	83, 0, 0, 0, 0, 278, 0, 0, 277, 72,
	73, 79, 80, 81, 82, 83, 72, 73, 79, 80,
	81, 82, 83, 0, 252, 74, 75, 0, 78, 76,
	77, 233, 74, 75, 0, 78, 76, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 74, 75, 0,
	78, 76, 77, 0, 74, 75, 0, 78, 76, 77,
	0, 0, 0, 0, 0, 0, 0, 72, 73, 79,
	80, 81, 82, 83, 72, 73, 79, 80, 81, 82,
	83, 0, 161, 0, 0, 0, 0, 0, 280, 72,
	73, 79, 80, 81, 82, 83, 72, 73, 79, 80,
	81, 82, 83, 276, 74, 75, 0, 78, 76, 77,
	251, 74, 75, 0, 78, 76, 77, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 74, 75, 0, 78,
	76, 77, 0, 74, 75, 0, 78, 76, 77, 0,
	0, 0, 0, 0, 0, 0, 72, 73, 79, 80,
	81, 82, 83, 72, 73, 79, 80, 81, 82, 83,
	243, 74, 75, 0, 78, 76, 77, 212, 72, 73,
	79, 80, 81, 82, 83, 72, 73, 79, 80, 81,
	82, 83, 167, 0, 74, 75, 269, 78, 76, 77,
	0, 0, 0, 0, 0, 0, 74, 75, 0, 78,
	76, 77, 0, 72, 73, 79, 80, 81, 82, 83,
	0, 0, 0, 272, 268, 0, 74, 75, 0, 78,
	76, 77, 0, 0, 0, 0, 72, 73, 79, 80,
	81, 82, 83, 162, 0, 0, 0, 234, 72, 73,
	79, 80, 81, 82, 83, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 73,
	79, 80, 81, 82, 83,
}

var yyPact = [...]int16{
	170, -1000, -1000, -40, 401, 52, 634, -1000, -1000, -1000,
	267, 79, 76, 549, 152, 96, 103, 69, -1000, -1000,
	-1000, -1000, -1000, 2, -1000, 364, 499, -1000, 656, 656,
	100, -1000, 549, 549, 656, 656, 90, 549, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -26, -1000, 46, 45,
	37, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 549, 549, 549, 549, 549, 549, 549, 549,
	549, 549, 549, 549, 163, -40, -1000, 267, -14, -1000,
	-1000, -1000, 69, 309, 66, -9, -1000, -1000, 549, -1000,
	-27, -1000, 31, 29, 549, -1000, -1000, -1000, -1000, 811,
	549, 32, 32, -1000, 1002, 129, 142, 77, -1000, 912,
	84, -1000, 604, 34, 34, 34, 267, 110, 110, 255,
	347, 166, 161, 97, 97, -1000, -1000, -1000, 172, 51,
	-1000, 128, -1000, -1000, -14, 711, -1000, -1000, -1000, 174,
	549, 549, 549, 172, -1, 267, -1000, 272, 656, 656,
	897, -1000, 549, -1000, 549, -14, -14, -1000, -1000, -1000,
	549, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 6, -1000, -1000, -40, -1000, -1000, -1000, 549, -14,
	-17, -1000, -32, -1000, 14, 12, 549, -1000, -1000, 449,
	760, 267, 970, -3, -1000, -1000, 549, -1000, -1000, 18,
	32, 18, 7, 267, -1000, -31, -33, 890, -1000, -29,
	-1000, 35, 267, -1000, -1000, -14, -1000, 711, -14, -14,
	840, -1000, 753, -1000, -1000, 5, 172, 267, 656, 656,
	181, 549, 549, -1000, -1000, 34, -1000, -1000, -1000, -1000,
	-1000, 0, -1000, 549, -1000, 32, 18, 105, 549, 549,
	947, 919, -1000, -14, 93, -1000, 982, 267, 549, 549,
	-1000, -1000, 549, 833, 738, 267, -1000, -1000, 549, 818,
	-1000,
}

var yyPgo = [...]int16{
	0, 267, 266, 261, 258, 9, 256, 219, 185, 255,
	0, 253, 1, 252, 251, 6, 10, 29, 247, 17,
	246, 245, 241, 239, 238, 14, 4, 2, 7, 237,
	13, 229, 228, 3, 193, 188, 12, 8,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 3, 3, 4, 4, 5, 5,
	6, 6, 7, 7, 8, 8, 9, 9, 33, 33,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 11, 11, 12, 12,
	12, 13, 13, 14, 14, 15, 15, 15, 15, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 17, 17, 18, 18, 18, 34, 34,
	35, 35, 19, 19, 19, 19, 19, 20, 20, 21,
	21, 22, 22, 23, 23, 24, 24, 25, 25, 25,
	25, 25, 37, 37, 37, 26, 26, 27, 27, 27,
	27, 27, 27, 27, 28, 28, 28, 29, 29, 30,
	30, 30, 31, 31, 32, 32, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36,
}

var yyR2 = [...]int8{
	0, 3, 0, 3, 0, 2, 6, 4, 0, 1,
	1, 1, 0, 2, 5, 8, 1, 3, 1, 1,
	2, 3, 5, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 1, 3, 1, 3,
	3, 1, 3, 1, 3, 3, 3, 5, 1, 1,
	1, 1, 2, 2, 1, 1, 1, 1, 4, 1,
	2, 3, 4, 2, 3, 1, 2, 2, 1, 2,
	1, 7, 3, 9, 9, 11, 2, 3, 2, 2,
	2, 3, 3, 1, 3, 0, 2, 4, 1, 1,
	1, 1, 2, 3, 4, 4, 5, 1, 3, 0,
	5, 0, 2, 0, 2, 1, 3, 3, 3, 5,
	1, 1, 1, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 4, 1, 3, 3,
	3, 3, 2, 3, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, 10, -3, -28, 62, -6, -4, -7,
	-10, 11, 12, -8, -16, 15, 13, 53, 33, 34,
	17, 18, 19, -34, -35, 62, 56, 35, 48, 49,
	36, -17, 20, 25, 27, 28, 16, 59, 29, 31,
	30, 32, 37, 38, 57, 63, -29, -30, 29, -36,
	37, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 46, 47, 4, 5, 8, 9, 7, 48,
	49, 50, 51, 52, 37, 37, -7, -10, 14, 34,
	-19, 54, 53, 56, 30, 29, -19, -17, 59, 63,
	-24, -25, -37, -17, 59, 29, 30, -36, 61, -10,
	-8, -16, -16, -17, -10, -10, -16, -16, 30, -10,
	-18, 63, 47, 58, 58, 58, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, 14, -5,
	-28, -11, -12, 30, 56, 62, -19, -17, 61, -10,
	58, 46, 58, 59, -20, -10, 63, 47, 58, 58,
	-10, 61, 21, -23, 26, 14, 14, 60, 40, 37,
	39, 63, -30, -27, -28, -31, 35, 37, 17, 18,
	19, 56, -27, -27, -33, 29, 30, 57, 46, 6,
	-13, -12, -14, -15, -37, -17, 59, 30, 61, 58,
	-10, -10, -10, -9, -33, 60, 57, 63, -25, -26,
	-16, -26, 60, -10, -10, -12, -12, -10, 61, -32,
	-27, -5, -10, -12, 61, 47, 63, 47, 58, 58,
	-10, 61, -10, 61, 57, 60, 57, -10, 46, 58,
	-21, 59, 59, 60, 61, 47, 57, -12, -15, -12,
	-12, 60, 61, 58, -33, -16, -26, -22, 22, 23,
	-10, -10, -27, 58, -10, 24, -10, -10, 57, 57,
	-12, 57, 21, -10, -10, -10, 60, 60, 57, -10,
	60,
}

var yyDef = [...]int16{
	2, -2, 4, 0, 12, 0, 0, 1, 5, 10,
	11, 0, 0, 12, 35, 0, 0, 49, 50, 51,
	54, 55, 56, 57, 59, 0, 0, 65, 0, 0,
	68, 70, 0, 0, 0, 0, 0, 0, 88, 89,
	90, 91, 83, 85, 3, 124, 0, 127, 0, 0,
	0, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 146, 147, 148, 149, 150, 151, 152, 153, 154,
	155, 156, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 8, 13, 20, 0, 78,
	79, 80, 0, 0, 0, 0, 52, 53, 0, 60,
	0, 105, 110, 111, 0, 112, 113, 114, 63, 0,
	0, 66, 67, 69, 0, 103, 0, 0, 76, 0,
	0, 125, 0, 0, 0, 0, 21, 24, 25, -2,
	27, 28, -2, 30, 31, 32, 33, 34, 0, 0,
	9, 0, 36, 38, 0, 0, 81, 82, 92, 0,
	0, 0, 0, 0, 0, 97, 61, 0, 0, 0,
	0, 64, 0, 72, 0, 0, 0, 77, 84, 86,
	0, 126, 128, 129, 117, 118, 119, 120, 121, 122,
	123, 0, 130, 131, 8, 18, 19, 7, 0, 0,
	0, 41, 0, 43, 0, 0, 0, -2, 93, 0,
	0, 23, 0, 0, 16, 58, 0, 62, 106, 107,
	115, 108, 0, 99, 104, 0, 0, 0, 132, 0,
	134, 0, 22, 37, 39, 0, 40, 0, 0, 0,
	0, 94, 0, 95, 14, 0, 0, 98, 0, 0,
	101, 0, 0, 87, 133, 0, 6, 42, 44, 45,
	46, 0, 96, 0, 17, 116, 109, 0, 0, 0,
	0, 0, 135, 0, 0, 71, 0, 102, 0, 0,
	47, 15, 0, 0, 0, 100, 73, 74, 0, 0,
	75,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 52, 3, 3,
	59, 60, 50, 48, 47, 49, 53, 51, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 58, 57,
	3, 3, 3, 54, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 56, 3, 61, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 62, 46, 63,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 55,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:58
		{
			query := yyDollar[3].value.(*Query)
			query.Meta = yyDollar[1].value.(*ConstObject)
			query.Imports = yyDollar[2].value.([]*Import)
			yylex.(*lexer).result = query
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:67
		{
			yyVAL.value = (*ConstObject)(nil)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:71
		{
			yyVAL.value = yyDollar[2].value
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:77
		{
			yyVAL.value = []*Import(nil)
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:81
		{
			yyVAL.value = append(yyDollar[1].value.([]*Import), yyDollar[2].value.(*Import))
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:87
		{
			yyVAL.value = &Import{ImportPath: yyDollar[2].token, ImportAlias: yyDollar[4].token, Meta: yyDollar[5].value.(*ConstObject)}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:91
		{
			yyVAL.value = &Import{IncludePath: yyDollar[2].token, Meta: yyDollar[3].value.(*ConstObject)}
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:97
		{
			yyVAL.value = (*ConstObject)(nil)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:104
		{
			yyVAL.value = &Query{FuncDefs: reverseFuncDef(yyDollar[1].value.([]*FuncDef))}
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:111
		{
			yyVAL.value = []*FuncDef(nil)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:115
		{
			yyVAL.value = append(yyDollar[2].value.([]*FuncDef), yyDollar[1].value.(*FuncDef))
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:121
		{
			yyVAL.value = &FuncDef{Name: yyDollar[2].token, Body: yyDollar[4].value.(*Query)}
		}
	case 15:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:125
		{
			yyVAL.value = &FuncDef{yyDollar[2].token, yyDollar[4].value.([]string), yyDollar[7].value.(*Query)}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:131
		{
			yyVAL.value = []string{yyDollar[1].token}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:135
		{
			yyVAL.value = append(yyDollar[1].value.([]string), yyDollar[3].token)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:145
		{
			query := yyDollar[2].value.(*Query)
			query.FuncDefs = prependFuncDef(query.FuncDefs, yyDollar[1].value.(*FuncDef))
			yyVAL.value = query
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:151
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpPipe, Right: yyDollar[3].value.(*Query)}
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:155
		{
			term := yyDollar[1].value.(*Term)
			term.SuffixList = append(term.SuffixList, &Suffix{Bind: &Bind{yyDollar[3].value.([]*Pattern), yyDollar[5].value.(*Query)}})
			yyVAL.value = &Query{Term: term}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:161
		{
			yyVAL.value = &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{yyDollar[2].token, yyDollar[4].value.(*Query)}}}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:165
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpComma, Right: yyDollar[3].value.(*Query)}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:169
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:173
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:177
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpOr, Right: yyDollar[3].value.(*Query)}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:181
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpAnd, Right: yyDollar[3].value.(*Query)}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:185
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:189
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpAdd, Right: yyDollar[3].value.(*Query)}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:193
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpSub, Right: yyDollar[3].value.(*Query)}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:197
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpMul, Right: yyDollar[3].value.(*Query)}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:201
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpDiv, Right: yyDollar[3].value.(*Query)}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:205
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpMod, Right: yyDollar[3].value.(*Query)}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:209
		{
			yyVAL.value = &Query{Term: yyDollar[1].value.(*Term)}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:215
		{
			yyVAL.value = []*Pattern{yyDollar[1].value.(*Pattern)}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:219
		{
			yyVAL.value = append(yyDollar[1].value.([]*Pattern), yyDollar[3].value.(*Pattern))
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:225
		{
			yyVAL.value = &Pattern{Name: yyDollar[1].token}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:229
		{
			yyVAL.value = &Pattern{Array: yyDollar[2].value.([]*Pattern)}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:233
		{
			yyVAL.value = &Pattern{Object: yyDollar[2].value.([]*PatternObject)}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:239
		{
			yyVAL.value = []*Pattern{yyDollar[1].value.(*Pattern)}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:243
		{
			yyVAL.value = append(yyDollar[1].value.([]*Pattern), yyDollar[3].value.(*Pattern))
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:249
		{
			yyVAL.value = []*PatternObject{yyDollar[1].value.(*PatternObject)}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:253
		{
			yyVAL.value = append(yyDollar[1].value.([]*PatternObject), yyDollar[3].value.(*PatternObject))
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:259
		{
			yyVAL.value = &PatternObject{Key: yyDollar[1].token, Val: yyDollar[3].value.(*Pattern)}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:263
		{
			yyVAL.value = &PatternObject{KeyString: yyDollar[1].value.(*String), Val: yyDollar[3].value.(*Pattern)}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:267
		{
			yyVAL.value = &PatternObject{KeyQuery: yyDollar[2].value.(*Query), Val: yyDollar[5].value.(*Pattern)}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:271
		{
			yyVAL.value = &PatternObject{Key: yyDollar[1].token}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:277
		{
			yyVAL.value = &Term{Type: TermTypeIdentity}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:281
		{
			yyVAL.value = &Term{Type: TermTypeRecurse}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:285
		{
			yyVAL.value = &Term{Type: TermTypeIndex, Index: &Index{Name: yyDollar[1].token}}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:289
		{
			suffix := yyDollar[2].value.(*Suffix)
			if suffix.Iter {
				yyVAL.value = &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{suffix}}
			} else {
				yyVAL.value = &Term{Type: TermTypeIndex, Index: suffix.Index}
			}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:298
		{
			yyVAL.value = &Term{Type: TermTypeIndex, Index: &Index{Str: yyDollar[2].value.(*String)}}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:302
		{
			yyVAL.value = &Term{Type: TermTypeNull}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:306
		{
			yyVAL.value = &Term{Type: TermTypeTrue}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:310
		{
			yyVAL.value = &Term{Type: TermTypeFalse}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:314
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token}}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:318
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token, Args: yyDollar[3].value.([]*Query)}}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:322
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token}}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:326
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{}}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:330
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{yyDollar[2].value.([]*ObjectKeyVal)}}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:334
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{yyDollar[2].value.([]*ObjectKeyVal)}}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:338
		{
			yyVAL.value = &Term{Type: TermTypeArray, Array: &Array{}}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:342
		{
			yyVAL.value = &Term{Type: TermTypeArray, Array: &Array{yyDollar[2].value.(*Query)}}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:346
		{
			yyVAL.value = &Term{Type: TermTypeNumber, Number: yyDollar[1].token}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:350
		{
			yyVAL.value = &Term{Type: TermTypeUnary, Unary: &Unary{OpAdd, yyDollar[2].value.(*Term)}}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:354
		{
			yyVAL.value = &Term{Type: TermTypeUnary, Unary: &Unary{OpSub, yyDollar[2].value.(*Term)}}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:358
		{
			yyVAL.value = &Term{Type: TermTypeFormat, Format: yyDollar[1].token}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:362
		{
			yyVAL.value = &Term{Type: TermTypeFormat, Format: yyDollar[1].token, Str: yyDollar[2].value.(*String)}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:366
		{
			yyVAL.value = &Term{Type: TermTypeString, Str: yyDollar[1].value.(*String)}
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:370
		{
			yyVAL.value = &Term{Type: TermTypeIf, If: &If{yyDollar[2].value.(*Query), yyDollar[4].value.(*Query), yyDollar[5].value.([]*IfElif), yyDollar[6].value.(*Query)}}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:374
		{
			yyVAL.value = &Term{Type: TermTypeTry, Try: &Try{yyDollar[2].value.(*Query), yyDollar[3].value.(*Query)}}
		}
	case 73:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:378
		{
			yyVAL.value = &Term{Type: TermTypeReduce, Reduce: &Reduce{yyDollar[2].value.(*Term), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query)}}
		}
	case 74:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:382
		{
			yyVAL.value = &Term{Type: TermTypeForeach, Foreach: &Foreach{yyDollar[2].value.(*Term), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query), nil}}
		}
	case 75:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:386
		{
			yyVAL.value = &Term{Type: TermTypeForeach, Foreach: &Foreach{yyDollar[2].value.(*Term), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query), yyDollar[10].value.(*Query)}}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:390
		{
			yyVAL.value = &Term{Type: TermTypeBreak, Break: yyDollar[2].token}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:394
		{
			yyVAL.value = &Term{Type: TermTypeQuery, Query: yyDollar[2].value.(*Query)}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:398
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Index: &Index{Name: yyDollar[2].token}})
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:402
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, yyDollar[2].value.(*Suffix))
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:406
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Optional: true})
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:410
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, yyDollar[3].value.(*Suffix))
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:414
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Index: &Index{Str: yyDollar[3].value.(*String)}})
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:420
		{
			yyVAL.value = &String{Str: yyDollar[1].token}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:424
		{
			yyVAL.value = &String{Queries: yyDollar[2].value.([]*Query)}
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:430
		{
			yyVAL.value = []*Query{}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:434
		{
			yyVAL.value = append(yyDollar[1].value.([]*Query), &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: yyDollar[2].token}}})
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:438
		{
			yylex.(*lexer).inString = true
			yyVAL.value = append(yyDollar[1].value.([]*Query), &Query{Term: &Term{Type: TermTypeQuery, Query: yyDollar[3].value.(*Query)}})
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:453
		{
			yyVAL.value = &Suffix{Iter: true}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:457
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query)}}
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:461
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query), IsSlice: true}}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:465
		{
			yyVAL.value = &Suffix{Index: &Index{End: yyDollar[3].value.(*Query), IsSlice: true}}
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:469
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query), End: yyDollar[4].value.(*Query), IsSlice: true}}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:475
		{
			yyVAL.value = []*Query{yyDollar[1].value.(*Query)}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:479
		{
			yyVAL.value = append(yyDollar[1].value.([]*Query), yyDollar[3].value.(*Query))
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:485
		{
			yyVAL.value = []*IfElif(nil)
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:489
		{
			yyVAL.value = append(yyDollar[1].value.([]*IfElif), &IfElif{yyDollar[3].value.(*Query), yyDollar[5].value.(*Query)})
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:495
		{
			yyVAL.value = (*Query)(nil)
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:499
		{
			yyVAL.value = yyDollar[2].value
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:505
		{
			yyVAL.value = (*Query)(nil)
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:509
		{
			yyVAL.value = yyDollar[2].value
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:515
		{
			yyVAL.value = []*ObjectKeyVal{yyDollar[1].value.(*ObjectKeyVal)}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:519
		{
			yyVAL.value = append(yyDollar[1].value.([]*ObjectKeyVal), yyDollar[3].value.(*ObjectKeyVal))
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:525
		{
			yyVAL.value = &ObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*ObjectVal)}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:529
		{
			yyVAL.value = &ObjectKeyVal{KeyString: yyDollar[1].value.(*String), Val: yyDollar[3].value.(*ObjectVal)}
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:533
		{
			yyVAL.value = &ObjectKeyVal{KeyQuery: yyDollar[2].value.(*Query), Val: yyDollar[5].value.(*ObjectVal)}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:537
		{
			yyVAL.value = &ObjectKeyVal{Key: yyDollar[1].token}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:541
		{
			yyVAL.value = &ObjectKeyVal{KeyString: yyDollar[1].value.(*String)}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:552
		{
			yyVAL.value = &ObjectVal{[]*Query{{Term: yyDollar[1].value.(*Term)}}}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:556
		{
			yyVAL.value = &ObjectVal{append(yyDollar[1].value.(*ObjectVal).Queries, &Query{Term: yyDollar[3].value.(*Term)})}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:562
		{
			yyVAL.value = &ConstTerm{Object: yyDollar[1].value.(*ConstObject)}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:566
		{
			yyVAL.value = &ConstTerm{Array: yyDollar[1].value.(*ConstArray)}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:570
		{
			yyVAL.value = &ConstTerm{Number: yyDollar[1].token}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:574
		{
			yyVAL.value = &ConstTerm{Str: yyDollar[1].token}
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:578
		{
			yyVAL.value = &ConstTerm{Null: true}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:582
		{
			yyVAL.value = &ConstTerm{True: true}
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:586
		{
			yyVAL.value = &ConstTerm{False: true}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:592
		{
			yyVAL.value = &ConstObject{}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:596
		{
			yyVAL.value = &ConstObject{yyDollar[2].value.([]*ConstObjectKeyVal)}
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:600
		{
			yyVAL.value = &ConstObject{yyDollar[2].value.([]*ConstObjectKeyVal)}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:606
		{
			yyVAL.value = []*ConstObjectKeyVal{yyDollar[1].value.(*ConstObjectKeyVal)}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:610
		{
			yyVAL.value = append(yyDollar[1].value.([]*ConstObjectKeyVal), yyDollar[3].value.(*ConstObjectKeyVal))
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:616
		{
			yyVAL.value = &ConstObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:620
		{
			yyVAL.value = &ConstObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:624
		{
			yyVAL.value = &ConstObjectKeyVal{KeyString: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:630
		{
			yyVAL.value = &ConstArray{}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:634
		{
			yyVAL.value = &ConstArray{yyDollar[2].value.([]*ConstTerm)}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:640
		{
			yyVAL.value = []*ConstTerm{yyDollar[1].value.(*ConstTerm)}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:644
		{
			yyVAL.value = append(yyDollar[1].value.([]*ConstTerm), yyDollar[3].value.(*ConstTerm))
		}
	}
	goto yystack /* stack new state and value */
}
