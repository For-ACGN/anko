//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:29
type yySymType struct {
	yys             int
	compstmt        []ast.Stmt
	stmt_if         ast.Stmt
	stmt_default    ast.Stmt
	stmt_case       ast.Stmt
	stmt_cases      []ast.Stmt
	stmts           []ast.Stmt
	stmt            ast.Stmt
	expr            ast.Expr
	exprs           []ast.Expr
	expr_many       []ast.Expr
	expr_lets       ast.Expr
	expr_pair       ast.Expr
	expr_pairs      []ast.Expr
	expr_idents     []string
	expr_type       string
	tok             ast.Token
	term            ast.Token
	terms           ast.Token
	opt_terms       ast.Token
	array_count     ast.ArrayCount
	expr_slice      ast.SliceExpr
	stmt_multi_case []ast.Stmt
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VARARG = 57350
const FUNC = 57351
const RETURN = 57352
const VAR = 57353
const THROW = 57354
const IF = 57355
const ELSE = 57356
const FOR = 57357
const IN = 57358
const EQEQ = 57359
const NEQ = 57360
const GE = 57361
const LE = 57362
const OROR = 57363
const ANDAND = 57364
const NEW = 57365
const TRUE = 57366
const FALSE = 57367
const NIL = 57368
const MODULE = 57369
const TRY = 57370
const CATCH = 57371
const FINALLY = 57372
const PLUSEQ = 57373
const MINUSEQ = 57374
const MULEQ = 57375
const DIVEQ = 57376
const ANDEQ = 57377
const OREQ = 57378
const BREAK = 57379
const CONTINUE = 57380
const PLUSPLUS = 57381
const MINUSMINUS = 57382
const POW = 57383
const SHIFTLEFT = 57384
const SHIFTRIGHT = 57385
const SWITCH = 57386
const CASE = 57387
const DEFAULT = 57388
const GO = 57389
const CHAN = 57390
const MAKE = 57391
const OPCHAN = 57392
const TYPE = 57393
const LEN = 57394
const DELETE = 57395
const UNARY = 57396

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"VAR",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OROR",
	"ANDAND",
	"NEW",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"TRY",
	"CATCH",
	"FINALLY",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"ANDEQ",
	"OREQ",
	"BREAK",
	"CONTINUE",
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
	"SHIFTLEFT",
	"SHIFTRIGHT",
	"SWITCH",
	"CASE",
	"DEFAULT",
	"GO",
	"CHAN",
	"MAKE",
	"OPCHAN",
	"TYPE",
	"LEN",
	"DELETE",
	"'='",
	"'?'",
	"':'",
	"','",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'{'",
	"'}'",
	"';'",
	"'['",
	"']'",
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"'('",
	"')'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:822

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 139,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	57, 54,
	-2, 1,
	-1, 10,
	57, 55,
	-2, 24,
	-1, 46,
	57, 54,
	-2, 140,
	-1, 89,
	67, 3,
	-2, 139,
	-1, 92,
	57, 55,
	-2, 51,
	-1, 93,
	16, 46,
	57, 46,
	-2, 64,
	-1, 95,
	67, 3,
	-2, 139,
	-1, 102,
	1, 69,
	8, 69,
	45, 69,
	46, 69,
	54, 69,
	56, 69,
	57, 69,
	66, 69,
	67, 69,
	68, 69,
	70, 69,
	76, 69,
	78, 69,
	-2, 64,
	-1, 104,
	1, 71,
	8, 71,
	45, 71,
	46, 71,
	54, 71,
	56, 71,
	57, 71,
	66, 71,
	67, 71,
	68, 71,
	70, 71,
	76, 71,
	78, 71,
	-2, 64,
	-1, 136,
	17, 0,
	18, 0,
	-2, 96,
	-1, 137,
	17, 0,
	18, 0,
	-2, 97,
	-1, 157,
	57, 55,
	-2, 51,
	-1, 159,
	67, 3,
	-2, 139,
	-1, 161,
	67, 3,
	-2, 139,
	-1, 163,
	67, 1,
	-2, 42,
	-1, 166,
	67, 3,
	-2, 139,
	-1, 193,
	67, 3,
	-2, 139,
	-1, 242,
	57, 56,
	-2, 52,
	-1, 243,
	1, 53,
	45, 53,
	46, 53,
	54, 53,
	56, 53,
	57, 57,
	67, 53,
	68, 53,
	78, 53,
	-2, 64,
	-1, 252,
	1, 57,
	8, 57,
	45, 57,
	46, 57,
	57, 57,
	67, 57,
	68, 57,
	70, 57,
	76, 57,
	78, 57,
	-2, 64,
	-1, 254,
	67, 3,
	-2, 139,
	-1, 256,
	67, 3,
	-2, 139,
	-1, 273,
	67, 3,
	-2, 139,
	-1, 283,
	1, 117,
	8, 117,
	45, 117,
	46, 117,
	54, 117,
	56, 117,
	57, 117,
	66, 117,
	67, 117,
	68, 117,
	70, 117,
	76, 117,
	78, 117,
	-2, 115,
	-1, 285,
	1, 121,
	8, 121,
	45, 121,
	46, 121,
	54, 121,
	56, 121,
	57, 121,
	66, 121,
	67, 121,
	68, 121,
	70, 121,
	76, 121,
	78, 121,
	-2, 119,
	-1, 300,
	67, 3,
	-2, 139,
	-1, 306,
	67, 3,
	-2, 139,
	-1, 307,
	67, 3,
	-2, 139,
	-1, 312,
	1, 116,
	8, 116,
	45, 116,
	46, 116,
	54, 116,
	56, 116,
	57, 116,
	66, 116,
	67, 116,
	68, 116,
	70, 116,
	76, 116,
	78, 116,
	-2, 114,
	-1, 313,
	1, 120,
	8, 120,
	45, 120,
	46, 120,
	54, 120,
	56, 120,
	57, 120,
	66, 120,
	67, 120,
	68, 120,
	70, 120,
	76, 120,
	78, 120,
	-2, 118,
	-1, 320,
	67, 3,
	-2, 139,
	-1, 321,
	67, 3,
	-2, 139,
	-1, 325,
	45, 3,
	46, 3,
	67, 3,
	-2, 139,
	-1, 329,
	67, 3,
	-2, 139,
	-1, 337,
	45, 3,
	46, 3,
	67, 3,
	-2, 139,
	-1, 338,
	45, 3,
	46, 3,
	67, 3,
	-2, 139,
	-1, 352,
	67, 3,
	-2, 139,
	-1, 353,
	67, 3,
	-2, 139,
}

const yyPrivate = 57344

const yyLast = 2995

var yyAct = [...]int{

	85, 179, 262, 10, 37, 261, 263, 185, 43, 233,
	6, 11, 48, 1, 234, 313, 86, 312, 227, 92,
	7, 96, 225, 84, 99, 100, 101, 103, 105, 6,
	94, 90, 308, 6, 274, 291, 110, 112, 272, 7,
	97, 176, 116, 7, 284, 119, 98, 10, 282, 233,
	250, 123, 124, 126, 290, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 140, 141, 142,
	143, 144, 145, 146, 147, 97, 276, 148, 149, 150,
	151, 182, 153, 155, 157, 288, 223, 154, 218, 152,
	117, 107, 115, 158, 114, 156, 2, 158, 170, 233,
	45, 113, 187, 162, 287, 189, 200, 172, 289, 168,
	165, 278, 285, 239, 184, 357, 283, 177, 191, 175,
	186, 299, 157, 190, 356, 122, 198, 264, 265, 180,
	108, 109, 349, 194, 345, 122, 344, 158, 68, 69,
	70, 71, 72, 73, 275, 160, 341, 353, 59, 260,
	340, 336, 326, 319, 222, 158, 219, 81, 318, 294,
	280, 204, 106, 258, 10, 208, 209, 255, 157, 56,
	57, 58, 217, 203, 201, 205, 51, 253, 53, 211,
	210, 78, 80, 300, 76, 236, 163, 212, 206, 224,
	311, 230, 231, 242, 352, 235, 237, 246, 161, 329,
	249, 321, 307, 251, 306, 273, 159, 244, 68, 69,
	70, 71, 72, 73, 95, 192, 269, 268, 59, 195,
	266, 267, 164, 121, 292, 158, 122, 81, 118, 324,
	281, 303, 229, 167, 83, 5, 264, 265, 8, 351,
	47, 346, 293, 259, 88, 180, 51, 286, 53, 245,
	238, 78, 80, 186, 76, 202, 49, 221, 298, 220,
	127, 87, 188, 122, 178, 301, 214, 4, 296, 91,
	297, 46, 213, 17, 226, 228, 302, 3, 251, 0,
	0, 310, 47, 0, 0, 120, 0, 305, 0, 314,
	0, 0, 315, 316, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 0,
	0, 0, 0, 0, 322, 0, 0, 81, 0, 0,
	327, 328, 277, 0, 279, 0, 0, 54, 55, 56,
	57, 58, 0, 343, 334, 335, 51, 0, 53, 339,
	0, 78, 80, 342, 76, 0, 0, 0, 0, 0,
	0, 347, 348, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 354, 355, 0, 82,
	62, 63, 65, 67, 77, 79, 0, 0, 0, 0,
	0, 0, 0, 0, 68, 69, 70, 71, 72, 73,
	0, 0, 74, 75, 59, 60, 61, 0, 0, 0,
	325, 0, 0, 81, 0, 0, 0, 0, 52, 0,
	332, 64, 66, 54, 55, 56, 57, 58, 0, 0,
	337, 338, 51, 0, 53, 0, 0, 78, 80, 331,
	76, 82, 62, 63, 65, 67, 77, 79, 0, 0,
	0, 0, 0, 0, 0, 0, 68, 69, 70, 71,
	72, 73, 0, 0, 74, 75, 59, 60, 61, 0,
	0, 0, 0, 0, 0, 81, 0, 0, 0, 0,
	52, 0, 241, 64, 66, 54, 55, 56, 57, 58,
	0, 0, 0, 0, 51, 0, 53, 0, 0, 78,
	80, 240, 76, 82, 62, 63, 65, 67, 77, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 81, 0, 0,
	0, 0, 52, 215, 0, 64, 66, 54, 55, 56,
	57, 58, 0, 0, 0, 0, 51, 216, 53, 0,
	0, 78, 80, 0, 76, 82, 62, 63, 65, 67,
	77, 79, 0, 0, 0, 0, 0, 0, 0, 0,
	68, 69, 70, 71, 72, 73, 0, 0, 74, 75,
	59, 60, 61, 0, 0, 0, 0, 0, 0, 81,
	0, 0, 0, 0, 52, 196, 0, 64, 66, 54,
	55, 56, 57, 58, 0, 0, 0, 0, 51, 197,
	53, 0, 0, 78, 80, 0, 76, 21, 22, 28,
	0, 0, 32, 14, 9, 15, 44, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 39, 29, 30, 31,
	16, 19, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 13, 0, 0, 0, 0, 0, 20, 0, 0,
	36, 0, 40, 41, 0, 38, 42, 0, 0, 0,
	0, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 6, 33, 0, 0, 24, 25, 26, 35, 0,
	0, 7, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 350, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 80, 333, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	0, 53, 0, 0, 78, 80, 330, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 323, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 51, 0, 53, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	320, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 317, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 304, 53,
	0, 0, 78, 80, 0, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	295, 53, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 51, 271, 53, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 257, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 256, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 0, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 254, 0, 0, 51, 0, 53,
	0, 0, 78, 80, 0, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	248, 53, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 51, 0, 53, 0, 0, 78, 80, 232, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 52,
	199, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 193, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 0, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 80, 181, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 169, 0, 0, 51,
	0, 53, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 166, 0,
	0, 51, 0, 53, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 77, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 50, 52,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 82, 62, 63, 65, 67, 77, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 0, 0, 74, 75, 59, 60, 61,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 52, 0, 0, 64, 66, 54, 55, 56, 57,
	58, 0, 0, 0, 0, 51, 0, 53, 0, 0,
	78, 80, 0, 76, 82, 62, 63, 65, 67, 77,
	79, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 52, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 183, 0, 76, 82, 62, 63, 65,
	67, 77, 79, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 69, 70, 71, 72, 73, 0, 0, 74,
	75, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 52, 0, 0, 64, 66,
	54, 55, 56, 57, 58, 0, 0, 0, 0, 51,
	0, 174, 0, 0, 78, 80, 0, 76, 82, 62,
	63, 65, 67, 77, 79, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 0,
	0, 74, 75, 59, 60, 61, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 52, 0, 0,
	64, 66, 54, 55, 56, 57, 58, 0, 0, 0,
	0, 51, 0, 173, 0, 0, 78, 80, 0, 76,
	82, 62, 63, 65, 67, 0, 79, 0, 0, 0,
	0, 0, 0, 0, 0, 68, 69, 70, 71, 72,
	73, 0, 0, 74, 75, 59, 60, 61, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 64, 66, 54, 55, 56, 57, 58, 0,
	0, 0, 0, 51, 0, 53, 0, 0, 78, 80,
	0, 76, 21, 22, 207, 0, 0, 32, 14, 9,
	15, 44, 0, 18, 0, 0, 0, 0, 0, 0,
	0, 39, 29, 30, 31, 16, 19, 0, 0, 0,
	0, 0, 0, 0, 0, 12, 13, 0, 0, 0,
	0, 0, 20, 0, 0, 36, 0, 40, 41, 0,
	38, 42, 0, 0, 0, 0, 0, 0, 0, 23,
	27, 0, 0, 0, 34, 0, 0, 33, 0, 0,
	24, 25, 26, 35, 82, 62, 63, 65, 67, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 0, 0, 74, 75, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 64, 66, 54, 55,
	56, 57, 58, 0, 0, 0, 0, 51, 0, 53,
	0, 0, 78, 80, 0, 76, 21, 22, 28, 0,
	0, 32, 14, 9, 15, 44, 0, 18, 0, 0,
	0, 0, 0, 0, 0, 39, 29, 30, 31, 16,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 20, 0, 0, 36,
	0, 40, 41, 0, 38, 42, 0, 0, 0, 0,
	0, 0, 0, 23, 27, 0, 65, 67, 34, 0,
	0, 33, 0, 0, 24, 25, 26, 35, 68, 69,
	70, 71, 72, 73, 0, 0, 74, 75, 59, 60,
	61, 0, 0, 252, 22, 28, 0, 81, 32, 0,
	0, 0, 0, 0, 0, 64, 66, 54, 55, 56,
	57, 58, 39, 29, 30, 31, 51, 0, 53, 0,
	0, 78, 80, 0, 76, 21, 22, 28, 0, 0,
	32, 0, 0, 0, 0, 0, 36, 0, 40, 41,
	0, 38, 42, 0, 39, 29, 30, 31, 0, 0,
	23, 27, 0, 0, 0, 34, 0, 0, 33, 309,
	0, 24, 25, 26, 35, 0, 0, 0, 36, 0,
	40, 41, 0, 38, 42, 0, 0, 0, 0, 21,
	22, 28, 23, 27, 32, 0, 0, 34, 0, 0,
	33, 270, 0, 24, 25, 26, 35, 0, 39, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	21, 22, 28, 0, 0, 32, 0, 0, 0, 0,
	0, 0, 36, 0, 40, 41, 0, 38, 42, 39,
	29, 30, 31, 0, 0, 0, 23, 27, 0, 0,
	0, 34, 0, 0, 33, 247, 0, 24, 25, 26,
	35, 0, 0, 36, 0, 40, 41, 0, 38, 42,
	0, 0, 171, 0, 21, 22, 28, 23, 27, 32,
	0, 0, 34, 0, 0, 33, 0, 0, 24, 25,
	26, 35, 0, 39, 29, 30, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 21, 22,
	28, 0, 0, 32, 0, 0, 0, 36, 0, 40,
	41, 0, 38, 42, 0, 0, 125, 39, 29, 30,
	31, 23, 27, 0, 0, 0, 34, 0, 0, 33,
	0, 0, 24, 25, 26, 35, 0, 0, 0, 0,
	0, 36, 0, 40, 41, 0, 38, 42, 0, 0,
	0, 0, 252, 22, 28, 23, 27, 32, 0, 0,
	34, 0, 0, 33, 0, 0, 24, 25, 26, 35,
	0, 39, 29, 30, 31, 0, 0, 0, 0, 0,
	0, 0, 0, 243, 22, 28, 0, 0, 32, 0,
	0, 0, 0, 0, 0, 36, 0, 40, 41, 0,
	38, 42, 39, 29, 30, 31, 0, 0, 0, 23,
	27, 0, 0, 0, 34, 0, 0, 33, 0, 0,
	24, 25, 26, 35, 0, 0, 36, 0, 40, 41,
	0, 38, 42, 0, 0, 0, 0, 111, 22, 28,
	23, 27, 32, 0, 0, 34, 0, 0, 33, 0,
	0, 24, 25, 26, 35, 0, 39, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 104, 22,
	28, 0, 0, 32, 0, 0, 0, 0, 0, 0,
	36, 0, 40, 41, 0, 38, 42, 39, 29, 30,
	31, 0, 0, 0, 23, 27, 0, 0, 0, 34,
	0, 0, 33, 0, 0, 24, 25, 26, 35, 0,
	0, 36, 0, 40, 41, 0, 38, 42, 0, 0,
	0, 0, 102, 22, 28, 23, 27, 32, 0, 0,
	34, 0, 0, 33, 0, 0, 24, 25, 26, 35,
	0, 39, 29, 30, 31, 0, 0, 0, 0, 0,
	0, 0, 0, 93, 22, 28, 0, 0, 32, 0,
	0, 0, 0, 0, 0, 36, 0, 40, 41, 0,
	38, 42, 39, 29, 30, 31, 0, 0, 0, 23,
	27, 0, 0, 0, 34, 0, 0, 33, 0, 0,
	24, 25, 26, 35, 0, 0, 36, 0, 40, 41,
	0, 38, 42, 0, 0, 0, 0, 0, 0, 0,
	23, 27, 0, 0, 0, 89, 0, 0, 33, 0,
	0, 24, 25, 26, 35,
}
var yyPact = [...]int{

	-58, -1000, 2372, -58, -58, -1000, -1000, -1000, -1000, 252,
	1854, 180, -1000, -1000, 2664, 2664, 257, 230, 2919, 148,
	2664, -29, -1000, 2664, 2664, 2664, 2888, 2834, -1000, -1000,
	-1000, -1000, 87, -58, -58, 2664, 2803, -1000, 26, 19,
	17, 2664, 15, 171, 2664, -1000, 613, -1000, 169, -1000,
	2664, 2630, 2664, 256, 2664, 2664, 2664, 2664, 2664, 2664,
	2664, 2664, 2664, 2664, 2664, 2664, 2664, 2664, 2664, 2664,
	2664, 2664, 2664, 2664, -1000, -1000, 2664, 2664, 2664, 2664,
	2664, 2664, 2664, 2664, 168, 1916, 1916, 140, 132, -58,
	206, 42, 1792, -29, 179, -58, 1730, 2576, 2664, 177,
	177, 177, -29, 2102, -29, 2040, 252, -34, 2664, 239,
	1668, 6, 1978, 2664, 249, 54, 1916, 2664, -58, 1606,
	-1000, 2664, -58, 1916, 539, 2664, 1544, -1000, 107, 107,
	177, 177, 177, 1916, 267, 267, 2417, 2417, 267, 267,
	267, 267, 1916, 1916, 1916, 1916, 1916, 1916, 1916, 2164,
	1916, 2298, 98, 1916, -1000, 2298, -1000, 1916, -58, -58,
	2664, -58, 121, 2238, 2664, 2664, -58, 2664, 120, -58,
	477, 2664, 80, 255, 253, 78, 252, -35, -39, -1000,
	176, -1000, 2664, 2664, 1482, -62, -1000, 249, 116, 246,
	43, 415, 2749, -58, -1000, 245, 2545, -1000, 1420, 2664,
	-26, -1000, 2718, 110, 1358, 100, -1000, 176, 1296, 1234,
	96, -1000, 214, 82, 191, 2491, -1000, 1172, -38, -1000,
	-1000, -1000, 139, -42, 68, -58, 41, -58, 93, 2664,
	40, 36, -1000, 243, -1000, 28, 38, -22, 167, -1000,
	-1000, 2664, 1916, -29, 92, -1000, 1110, -1000, -1000, 1916,
	-1000, 1916, -29, -1000, -58, -1000, -58, 2664, -1000, 117,
	-1000, -1000, -1000, -1000, 2664, 175, -1000, -1000, -1000, 1048,
	-1000, -1000, -1000, -58, 138, 136, -44, 2459, -1000, 123,
	-1000, 1916, -59, -1000, -61, -1000, -1000, -1000, 2664, -1000,
	-1000, 2664, 2664, 986, -1000, -1000, 91, 86, 924, 135,
	-58, 862, 173, -58, -1000, 85, -58, -58, 133, -1000,
	-1000, -1000, -1000, -1000, 800, 353, 738, -1000, -1000, -1000,
	-58, -58, 84, -58, -58, -58, -1000, 83, 79, -58,
	-1000, -1000, 2664, -1000, 69, 67, 211, -58, -58, -1000,
	-1000, -1000, 65, 676, -1000, 209, 128, -1000, -1000, -1000,
	-1000, 81, -58, -58, 57, 48, -1000, -1000,
}
var yyPgo = [...]int{

	0, 13, 277, 238, 273, 6, 5, 272, 0, 8,
	11, 269, 1, 264, 12, 7, 262, 4, 2, 96,
	267, 235,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 4, 4, 7, 7,
	7, 7, 7, 7, 7, 6, 18, 5, 16, 16,
	16, 12, 13, 13, 13, 14, 14, 14, 15, 15,
	11, 10, 10, 10, 9, 9, 9, 9, 17, 17,
	17, 17, 17, 17, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 19,
	19, 20, 20, 21, 21,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 1, 7, 5, 5, 0, 2,
	2, 2, 2, 2, 2, 5, 5, 4, 0, 2,
	3, 3, 0, 1, 4, 0, 1, 4, 1, 3,
	3, 1, 4, 4, 0, 1, 4, 4, 6, 5,
	5, 6, 5, 5, 1, 1, 2, 2, 2, 2,
	4, 2, 4, 1, 1, 1, 1, 5, 3, 7,
	8, 8, 9, 5, 6, 5, 6, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	3, 3, 3, 3, 5, 4, 6, 5, 5, 4,
	6, 5, 4, 4, 1, 4, 4, 5, 7, 5,
	7, 9, 7, 3, 2, 4, 6, 3, 3, 0,
	1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -19, -2, -20, -21, 68, 78, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 4, 5, 61, 72, 73, 74, 62, 6, 24,
	25, 26, 9, 69, 66, 75, 47, -17, 52, 23,
	49, 50, 53, -9, 13, -19, -20, -21, -14, 4,
	54, 69, 55, 71, 60, 61, 62, 63, 64, 41,
	42, 43, 17, 18, 58, 19, 59, 20, 31, 32,
	33, 34, 35, 36, 39, 40, 77, 21, 74, 22,
	75, 50, 16, 54, -9, -8, -8, 4, 14, 66,
	-14, -11, -8, 4, -10, 66, -8, 69, 75, -8,
	-8, -8, 4, -8, 4, -8, 75, 4, -19, -19,
	-8, 4, -8, 75, 75, 75, -8, 75, 57, -8,
	-3, 54, 57, -8, -8, 56, -8, 4, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -9, -8, -17, -8, -10, -8, 57, 66,
	13, 66, -1, -19, 16, 68, 66, 54, -1, 66,
	-8, 56, -9, 71, 71, -14, 75, -9, -13, -12,
	6, 76, 75, 75, -8, -15, 4, 48, -16, 51,
	69, -8, -19, 66, -10, -19, 56, 70, -8, 56,
	8, 76, -19, -1, -8, -1, 67, 6, -8, -8,
	-1, -10, 67, -7, -19, 56, 70, -8, 8, 76,
	4, 4, 76, 8, -14, 57, -19, 57, -19, 56,
	-9, -9, 76, 71, 76, -15, 69, -15, 4, 70,
	76, 57, -8, 4, -1, 4, -8, 70, 70, -8,
	76, -8, 4, 67, 66, 67, 66, 68, 67, 29,
	67, -6, -18, -5, 45, 46, -6, -5, -18, -8,
	70, 70, 76, 66, 76, 76, 8, -19, 70, -19,
	67, -8, 8, 76, 8, 76, 4, 76, 57, 70,
	76, 57, 57, -8, 67, 70, -1, -1, -8, 4,
	66, -8, -10, 56, 70, -1, 66, 66, 76, 70,
	-12, 67, 76, 76, -8, -8, -8, 76, 67, 67,
	66, 66, -1, 56, 56, -19, 67, -1, -1, 66,
	76, 76, 57, 76, -1, -1, 67, -19, -19, -1,
	67, 67, -1, -8, 67, 67, 30, -1, -1, 67,
	76, 30, 66, 66, -1, -1, 67, 67,
}
var yyDef = [...]int{

	-2, -2, -2, 139, 140, 141, 143, 144, 4, 45,
	-2, 0, 9, 10, 54, 0, 0, 14, 45, 0,
	0, 64, 65, 0, 0, 0, 0, 0, 73, 74,
	75, 76, 0, 139, 139, 0, 0, 124, 0, 0,
	0, 0, 0, 0, 0, 2, -2, 142, 0, 46,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 108, 109, 0, 0, 0, 0,
	54, 0, 0, 54, 11, 55, 12, 0, 0, -2,
	0, 0, -2, -2, 0, -2, 0, 0, 54, 66,
	67, 68, -2, 0, -2, 0, 45, 0, 54, 42,
	0, 64, 0, 0, 0, 38, 134, 0, 139, 0,
	5, 54, 139, 7, 0, 0, 0, 78, 88, 89,
	90, 91, 92, 93, 94, 95, -2, -2, 98, 99,
	100, 101, 102, 103, 104, 105, 106, 107, 110, 111,
	112, 113, 0, 133, 124, 138, 8, -2, 139, -2,
	0, -2, 0, -2, 0, 0, -2, 54, 0, 28,
	0, 0, 0, 0, 0, 0, 45, 139, 139, 43,
	0, 87, 54, 54, 0, 0, 48, 0, 0, 0,
	0, 0, 0, -2, 6, 0, 0, 123, 0, 0,
	0, 119, 0, 0, 0, 0, 15, 73, 0, 0,
	0, 50, 0, 0, 0, 0, 122, 0, 0, 115,
	70, 72, 0, 0, 0, 139, 0, 139, 0, 0,
	0, 0, 125, 0, 126, 0, 0, 0, 0, 39,
	135, 0, -2, -2, 0, 47, 0, 62, 63, 77,
	118, 56, -2, 13, -2, 26, -2, 0, 18, 0,
	23, 31, 33, 34, 54, 0, 29, 30, 32, 0,
	59, 60, 114, -2, 0, 0, 0, 0, 83, 0,
	85, 41, 0, -2, 0, -2, 49, 127, 0, 40,
	129, 0, 0, 0, 27, 61, 0, 0, 0, 0,
	-2, 55, 0, 139, 58, 0, -2, -2, 0, 84,
	44, 86, -2, -2, 0, 0, 0, 136, 25, 16,
	-2, -2, 0, 139, 139, -2, 79, 0, 0, -2,
	128, 130, 0, 132, 0, 0, 22, -2, -2, 37,
	80, 81, 0, 0, 17, 21, 0, 35, 36, 82,
	131, 0, -2, -2, 0, 0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	78, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 72, 3, 3, 3, 64, 74, 3,
	75, 76, 62, 60, 57, 61, 71, 63, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 56, 68,
	59, 54, 58, 55, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 70, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 66, 77, 67,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 65,
}
var yyTok3 = [...]int{
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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
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
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:89
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:96
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:116
		{
			if len(yyDollar[1].expr_many) == 2 && len(yyDollar[3].expr_many) == 1 {
				if _, ok := yyDollar[3].expr_many[0].(*ast.ItemExpr); ok {
					yyVAL.stmt = &ast.LetMapItemStmt{Lhss: yyDollar[1].expr_many, Rhs: yyDollar[3].expr_many[0]}
				} else {
					yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
				}
			} else {
				yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
			}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:188
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:211
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:216
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyDollar[4].compstmt...)
			}
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:225
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:231
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:235
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:247
		{
			yyVAL.stmt_cases = yyDollar[2].stmt_multi_case
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_multi_case...)
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:255
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:272
		{
			var cases = []ast.Stmt{}
			for _, stmt := range yyDollar[2].expr_many {
				cases = append(cases, &ast.CaseStmt{Expr: stmt, Stmts: yyDollar[5].compstmt})
			}
			yyVAL.stmt_multi_case = cases
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:295
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:314
		{
			if len(yyDollar[1].expr_pairs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.expr_idents = []string{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:326
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:330
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.expr_type = yyDollar[1].tok.Lit
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.expr_type = yyVAL.expr_type + "." + yyDollar[3].tok.Lit
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:348
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:362
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:367
		{
			yyVAL.exprs = nil
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:371
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:375
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:382
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 58:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:395
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
		}
	case 60:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:399
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
		}
	case 61:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:403
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:417
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:422
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:427
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:432
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:437
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:442
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:447
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:452
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 80:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 81:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 82:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 84:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:517
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:522
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 86:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:531
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:540
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:545
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:550
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:555
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:560
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:565
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:570
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:575
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:580
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:585
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:590
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:595
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:600
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:605
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:610
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:615
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:620
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:625
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:630
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:635
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:640
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:645
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:650
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:655
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:660
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:665
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:670
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:675
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:680
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 116:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:685
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:690
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:695
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:700
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:705
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:710
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:715
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:720
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:725
		{
			yyVAL.expr = &yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:730
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:735
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:740
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 128:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:745
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 129:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:750
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:755
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 131:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:760
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 132:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:765
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:770
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:775
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:780
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 136:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:785
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr, KeyExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:790
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: yyDollar[3].expr_slice}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:795
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: ast.SliceExpr{Value: yyDollar[3].expr, Begin: nil, End: nil}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:808
		{
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:811
		{
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:816
		{
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:819
		{
		}
	}
	goto yystack /* stack new state and value */
}
