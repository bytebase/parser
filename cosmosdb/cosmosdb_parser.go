// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cosmosdb // CosmosDBParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CosmosDBParser struct {
	*antlr.BaseParser
}

var CosmosDBParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cosmosdbparserParserInit() {
	staticData := &CosmosDBParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'*'", "'AS'", "'SELECT'", "'FROM'", "'DISTINCT'", "'UNDEFINED'",
		"'NULL'", "'FALSE'", "'TRUE'", "'NOT'", "'UDF'", "'WHERE'", "'AND'",
		"'OR'", "'@'", "'{'", "'}'", "'['", "']'", "'('", "')'", "'''", "'\"'",
		"','", "'.'", "'?'", "':'", "'+'", "'-'", "'~'", "'/'", "'%'", "'&'",
		"'|'", "'||'", "'^'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "MULTIPLY_OPERATOR", "AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL",
		"DISTINCT_SYMBOL", "UNDEFINED_SYMBOL", "NULL_SYMBOL", "FALSE_SYMBOL",
		"TRUE_SYMBOL", "NOT_SYMBOL", "UDF_SYMBOL", "WHERE_SYMBOL", "AND_SYMBOL",
		"OR_SYMBOL", "AT_SYMBOL", "LC_BRACKET_SYMBOL", "RC_BRACKET_SYMBOL",
		"LS_BRACKET_SYMBOL", "RS_BRACKET_SYMBOL", "LR_BRACKET_SYMBOL", "RR_BRACKET_SYMBOL",
		"SINGLE_QUOTE_SYMBOL", "DOUBLE_QUOTE_SYMBOL", "COMMA_SYMBOL", "DOT_SYMBOL",
		"QUESTION_MARK_SYMBOL", "COLON_SYMBOL", "PLUS_SYMBOL", "MINUS_SYMBOL",
		"BIT_NOT_SYMBOL", "DIVIDE_SYMBOL", "MODULO_SYMBOL", "BIT_AND_SYMBOL",
		"BIT_OR_SYMBOL", "DOUBLE_BAR_SYMBOL", "BIT_XOR_SYMBOL", "EQUAL_SYMBOL",
		"IDENTIFIER", "WHITESPACE", "DECIMAL", "REAL", "FLOAT", "HEXADECIMAL",
		"SINGLE_QUOTE_STRING_LITERAL", "DOUBLE_QUOTE_STRING_LITERAL",
	}
	staticData.RuleNames = []string{
		"root", "select", "select_clause", "select_specification", "from_clause",
		"where_clause", "from_specification", "from_source", "container_expression",
		"container_name", "object_property_list", "object_property", "property_alias",
		"scalar_expression", "scalar_expression_in_where", "create_array_expression",
		"create_object_expression", "scalar_function_expression", "udf_scalar_function_expression",
		"builtin_function_expression", "binary_operator", "unary_operator",
		"parameter_name", "constant", "object_constant", "object_constant_field_pair",
		"array_constant", "string_constant", "undefined_constant", "null_constant",
		"boolean_constant", "number_constant", "string_literal", "decimal_literal",
		"hexadecimal_literal", "property_name", "array_index", "input_alias",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 308, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 83, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 90, 8, 3, 1, 3, 3, 3, 93, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 107,
		8, 8, 1, 8, 3, 8, 110, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 117,
		8, 10, 10, 10, 12, 10, 120, 9, 10, 1, 11, 1, 11, 3, 11, 124, 8, 11, 1,
		11, 3, 11, 127, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		3, 13, 136, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3,
		13, 145, 8, 13, 1, 13, 5, 13, 148, 8, 13, 10, 13, 12, 13, 151, 9, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 167, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 192, 8,
		14, 1, 14, 5, 14, 195, 8, 14, 10, 14, 12, 14, 198, 9, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 206, 8, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 5, 18, 215, 8, 18, 10, 18, 12, 18, 218, 9, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 227, 8, 19, 10,
		19, 12, 19, 230, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 248,
		8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 254, 8, 24, 10, 24, 12, 24, 257,
		9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 266, 8,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 275, 8, 26,
		10, 26, 12, 26, 278, 9, 26, 3, 26, 280, 8, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 3, 31, 294,
		8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 0, 2, 26, 28, 38, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 0, 5, 3, 0, 1, 1, 28, 29, 31,
		37, 1, 0, 28, 30, 1, 0, 8, 9, 1, 0, 44, 45, 1, 0, 40, 42, 309, 0, 76, 1,
		0, 0, 0, 2, 79, 1, 0, 0, 0, 4, 84, 1, 0, 0, 0, 6, 92, 1, 0, 0, 0, 8, 94,
		1, 0, 0, 0, 10, 97, 1, 0, 0, 0, 12, 100, 1, 0, 0, 0, 14, 102, 1, 0, 0,
		0, 16, 104, 1, 0, 0, 0, 18, 111, 1, 0, 0, 0, 20, 113, 1, 0, 0, 0, 22, 121,
		1, 0, 0, 0, 24, 128, 1, 0, 0, 0, 26, 135, 1, 0, 0, 0, 28, 166, 1, 0, 0,
		0, 30, 199, 1, 0, 0, 0, 32, 201, 1, 0, 0, 0, 34, 205, 1, 0, 0, 0, 36, 207,
		1, 0, 0, 0, 38, 221, 1, 0, 0, 0, 40, 233, 1, 0, 0, 0, 42, 235, 1, 0, 0,
		0, 44, 237, 1, 0, 0, 0, 46, 247, 1, 0, 0, 0, 48, 249, 1, 0, 0, 0, 50, 265,
		1, 0, 0, 0, 52, 270, 1, 0, 0, 0, 54, 283, 1, 0, 0, 0, 56, 285, 1, 0, 0,
		0, 58, 287, 1, 0, 0, 0, 60, 289, 1, 0, 0, 0, 62, 293, 1, 0, 0, 0, 64, 295,
		1, 0, 0, 0, 66, 297, 1, 0, 0, 0, 68, 299, 1, 0, 0, 0, 70, 301, 1, 0, 0,
		0, 72, 303, 1, 0, 0, 0, 74, 305, 1, 0, 0, 0, 76, 77, 3, 2, 1, 0, 77, 78,
		5, 0, 0, 1, 78, 1, 1, 0, 0, 0, 79, 80, 3, 4, 2, 0, 80, 82, 3, 8, 4, 0,
		81, 83, 3, 10, 5, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 3, 1,
		0, 0, 0, 84, 85, 5, 3, 0, 0, 85, 86, 3, 6, 3, 0, 86, 5, 1, 0, 0, 0, 87,
		93, 5, 1, 0, 0, 88, 90, 5, 5, 0, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0,
		0, 90, 91, 1, 0, 0, 0, 91, 93, 3, 20, 10, 0, 92, 87, 1, 0, 0, 0, 92, 89,
		1, 0, 0, 0, 93, 7, 1, 0, 0, 0, 94, 95, 5, 4, 0, 0, 95, 96, 3, 12, 6, 0,
		96, 9, 1, 0, 0, 0, 97, 98, 5, 12, 0, 0, 98, 99, 3, 28, 14, 0, 99, 11, 1,
		0, 0, 0, 100, 101, 3, 14, 7, 0, 101, 13, 1, 0, 0, 0, 102, 103, 3, 16, 8,
		0, 103, 15, 1, 0, 0, 0, 104, 109, 3, 18, 9, 0, 105, 107, 5, 2, 0, 0, 106,
		105, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110,
		5, 38, 0, 0, 109, 106, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 17, 1, 0,
		0, 0, 111, 112, 5, 38, 0, 0, 112, 19, 1, 0, 0, 0, 113, 118, 3, 22, 11,
		0, 114, 115, 5, 24, 0, 0, 115, 117, 3, 22, 11, 0, 116, 114, 1, 0, 0, 0,
		117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		21, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 126, 3, 26, 13, 0, 122, 124,
		5, 2, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0,
		0, 0, 125, 127, 3, 24, 12, 0, 126, 123, 1, 0, 0, 0, 126, 127, 1, 0, 0,
		0, 127, 23, 1, 0, 0, 0, 128, 129, 5, 38, 0, 0, 129, 25, 1, 0, 0, 0, 130,
		131, 6, 13, -1, 0, 131, 136, 3, 74, 37, 0, 132, 133, 3, 42, 21, 0, 133,
		134, 3, 26, 13, 1, 134, 136, 1, 0, 0, 0, 135, 130, 1, 0, 0, 0, 135, 132,
		1, 0, 0, 0, 136, 149, 1, 0, 0, 0, 137, 138, 10, 3, 0, 0, 138, 139, 5, 25,
		0, 0, 139, 148, 3, 70, 35, 0, 140, 141, 10, 2, 0, 0, 141, 144, 5, 18, 0,
		0, 142, 145, 5, 45, 0, 0, 143, 145, 3, 72, 36, 0, 144, 142, 1, 0, 0, 0,
		144, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 5, 19, 0, 0, 147,
		137, 1, 0, 0, 0, 147, 140, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 27, 1, 0, 0, 0, 151, 149, 1, 0,
		0, 0, 152, 153, 6, 14, -1, 0, 153, 167, 3, 46, 23, 0, 154, 167, 3, 74,
		37, 0, 155, 167, 3, 44, 22, 0, 156, 157, 3, 42, 21, 0, 157, 158, 3, 28,
		14, 7, 158, 167, 1, 0, 0, 0, 159, 167, 3, 34, 17, 0, 160, 167, 3, 32, 16,
		0, 161, 167, 3, 30, 15, 0, 162, 163, 5, 20, 0, 0, 163, 164, 3, 28, 14,
		0, 164, 165, 5, 21, 0, 0, 165, 167, 1, 0, 0, 0, 166, 152, 1, 0, 0, 0, 166,
		154, 1, 0, 0, 0, 166, 155, 1, 0, 0, 0, 166, 156, 1, 0, 0, 0, 166, 159,
		1, 0, 0, 0, 166, 160, 1, 0, 0, 0, 166, 161, 1, 0, 0, 0, 166, 162, 1, 0,
		0, 0, 167, 196, 1, 0, 0, 0, 168, 169, 10, 11, 0, 0, 169, 170, 5, 13, 0,
		0, 170, 195, 3, 28, 14, 12, 171, 172, 10, 10, 0, 0, 172, 173, 5, 14, 0,
		0, 173, 195, 3, 28, 14, 11, 174, 175, 10, 6, 0, 0, 175, 176, 3, 40, 20,
		0, 176, 177, 3, 28, 14, 7, 177, 195, 1, 0, 0, 0, 178, 179, 10, 5, 0, 0,
		179, 180, 5, 26, 0, 0, 180, 181, 3, 28, 14, 0, 181, 182, 5, 27, 0, 0, 182,
		183, 3, 28, 14, 6, 183, 195, 1, 0, 0, 0, 184, 185, 10, 9, 0, 0, 185, 186,
		5, 25, 0, 0, 186, 195, 3, 70, 35, 0, 187, 188, 10, 8, 0, 0, 188, 191, 5,
		18, 0, 0, 189, 192, 5, 45, 0, 0, 190, 192, 3, 72, 36, 0, 191, 189, 1, 0,
		0, 0, 191, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 5, 19, 0, 0,
		194, 168, 1, 0, 0, 0, 194, 171, 1, 0, 0, 0, 194, 174, 1, 0, 0, 0, 194,
		178, 1, 0, 0, 0, 194, 184, 1, 0, 0, 0, 194, 187, 1, 0, 0, 0, 195, 198,
		1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 29, 1, 0,
		0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 3, 52, 26, 0, 200, 31, 1, 0, 0, 0,
		201, 202, 3, 48, 24, 0, 202, 33, 1, 0, 0, 0, 203, 206, 3, 36, 18, 0, 204,
		206, 3, 38, 19, 0, 205, 203, 1, 0, 0, 0, 205, 204, 1, 0, 0, 0, 206, 35,
		1, 0, 0, 0, 207, 208, 5, 11, 0, 0, 208, 209, 5, 25, 0, 0, 209, 210, 5,
		38, 0, 0, 210, 211, 5, 20, 0, 0, 211, 216, 3, 28, 14, 0, 212, 213, 5, 24,
		0, 0, 213, 215, 3, 28, 14, 0, 214, 212, 1, 0, 0, 0, 215, 218, 1, 0, 0,
		0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218,
		216, 1, 0, 0, 0, 219, 220, 5, 21, 0, 0, 220, 37, 1, 0, 0, 0, 221, 222,
		5, 38, 0, 0, 222, 223, 5, 20, 0, 0, 223, 228, 3, 28, 14, 0, 224, 225, 5,
		24, 0, 0, 225, 227, 3, 28, 14, 0, 226, 224, 1, 0, 0, 0, 227, 230, 1, 0,
		0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 231, 1, 0, 0, 0,
		230, 228, 1, 0, 0, 0, 231, 232, 5, 21, 0, 0, 232, 39, 1, 0, 0, 0, 233,
		234, 7, 0, 0, 0, 234, 41, 1, 0, 0, 0, 235, 236, 7, 1, 0, 0, 236, 43, 1,
		0, 0, 0, 237, 238, 5, 15, 0, 0, 238, 239, 5, 38, 0, 0, 239, 45, 1, 0, 0,
		0, 240, 248, 3, 56, 28, 0, 241, 248, 3, 58, 29, 0, 242, 248, 3, 60, 30,
		0, 243, 248, 3, 62, 31, 0, 244, 248, 3, 54, 27, 0, 245, 248, 3, 52, 26,
		0, 246, 248, 3, 48, 24, 0, 247, 240, 1, 0, 0, 0, 247, 241, 1, 0, 0, 0,
		247, 242, 1, 0, 0, 0, 247, 243, 1, 0, 0, 0, 247, 244, 1, 0, 0, 0, 247,
		245, 1, 0, 0, 0, 247, 246, 1, 0, 0, 0, 248, 47, 1, 0, 0, 0, 249, 250, 5,
		16, 0, 0, 250, 255, 3, 50, 25, 0, 251, 252, 5, 24, 0, 0, 252, 254, 3, 50,
		25, 0, 253, 251, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0,
		255, 256, 1, 0, 0, 0, 256, 258, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258,
		259, 5, 17, 0, 0, 259, 49, 1, 0, 0, 0, 260, 266, 3, 70, 35, 0, 261, 262,
		5, 23, 0, 0, 262, 263, 3, 70, 35, 0, 263, 264, 5, 23, 0, 0, 264, 266, 1,
		0, 0, 0, 265, 260, 1, 0, 0, 0, 265, 261, 1, 0, 0, 0, 266, 267, 1, 0, 0,
		0, 267, 268, 5, 24, 0, 0, 268, 269, 3, 46, 23, 0, 269, 51, 1, 0, 0, 0,
		270, 279, 5, 18, 0, 0, 271, 276, 3, 46, 23, 0, 272, 273, 5, 24, 0, 0, 273,
		275, 3, 46, 23, 0, 274, 272, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276, 274,
		1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276, 1, 0,
		0, 0, 279, 271, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0,
		281, 282, 5, 19, 0, 0, 282, 53, 1, 0, 0, 0, 283, 284, 3, 64, 32, 0, 284,
		55, 1, 0, 0, 0, 285, 286, 5, 6, 0, 0, 286, 57, 1, 0, 0, 0, 287, 288, 5,
		7, 0, 0, 288, 59, 1, 0, 0, 0, 289, 290, 7, 2, 0, 0, 290, 61, 1, 0, 0, 0,
		291, 294, 3, 66, 33, 0, 292, 294, 3, 68, 34, 0, 293, 291, 1, 0, 0, 0, 293,
		292, 1, 0, 0, 0, 294, 63, 1, 0, 0, 0, 295, 296, 7, 3, 0, 0, 296, 65, 1,
		0, 0, 0, 297, 298, 7, 4, 0, 0, 298, 67, 1, 0, 0, 0, 299, 300, 5, 43, 0,
		0, 300, 69, 1, 0, 0, 0, 301, 302, 5, 38, 0, 0, 302, 71, 1, 0, 0, 0, 303,
		304, 5, 40, 0, 0, 304, 73, 1, 0, 0, 0, 305, 306, 5, 38, 0, 0, 306, 75,
		1, 0, 0, 0, 25, 82, 89, 92, 106, 109, 118, 123, 126, 135, 144, 147, 149,
		166, 191, 194, 196, 205, 216, 228, 247, 255, 265, 276, 279, 293,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CosmosDBParserInit initializes any static state used to implement CosmosDBParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCosmosDBParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CosmosDBParserInit() {
	staticData := &CosmosDBParserParserStaticData
	staticData.once.Do(cosmosdbparserParserInit)
}

// NewCosmosDBParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCosmosDBParser(input antlr.TokenStream) *CosmosDBParser {
	CosmosDBParserInit()
	this := new(CosmosDBParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CosmosDBParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CosmosDBParser.g4"

	return this
}

// CosmosDBParser tokens.
const (
	CosmosDBParserEOF                         = antlr.TokenEOF
	CosmosDBParserMULTIPLY_OPERATOR           = 1
	CosmosDBParserAS_SYMBOL                   = 2
	CosmosDBParserSELECT_SYMBOL               = 3
	CosmosDBParserFROM_SYMBOL                 = 4
	CosmosDBParserDISTINCT_SYMBOL             = 5
	CosmosDBParserUNDEFINED_SYMBOL            = 6
	CosmosDBParserNULL_SYMBOL                 = 7
	CosmosDBParserFALSE_SYMBOL                = 8
	CosmosDBParserTRUE_SYMBOL                 = 9
	CosmosDBParserNOT_SYMBOL                  = 10
	CosmosDBParserUDF_SYMBOL                  = 11
	CosmosDBParserWHERE_SYMBOL                = 12
	CosmosDBParserAND_SYMBOL                  = 13
	CosmosDBParserOR_SYMBOL                   = 14
	CosmosDBParserAT_SYMBOL                   = 15
	CosmosDBParserLC_BRACKET_SYMBOL           = 16
	CosmosDBParserRC_BRACKET_SYMBOL           = 17
	CosmosDBParserLS_BRACKET_SYMBOL           = 18
	CosmosDBParserRS_BRACKET_SYMBOL           = 19
	CosmosDBParserLR_BRACKET_SYMBOL           = 20
	CosmosDBParserRR_BRACKET_SYMBOL           = 21
	CosmosDBParserSINGLE_QUOTE_SYMBOL         = 22
	CosmosDBParserDOUBLE_QUOTE_SYMBOL         = 23
	CosmosDBParserCOMMA_SYMBOL                = 24
	CosmosDBParserDOT_SYMBOL                  = 25
	CosmosDBParserQUESTION_MARK_SYMBOL        = 26
	CosmosDBParserCOLON_SYMBOL                = 27
	CosmosDBParserPLUS_SYMBOL                 = 28
	CosmosDBParserMINUS_SYMBOL                = 29
	CosmosDBParserBIT_NOT_SYMBOL              = 30
	CosmosDBParserDIVIDE_SYMBOL               = 31
	CosmosDBParserMODULO_SYMBOL               = 32
	CosmosDBParserBIT_AND_SYMBOL              = 33
	CosmosDBParserBIT_OR_SYMBOL               = 34
	CosmosDBParserDOUBLE_BAR_SYMBOL           = 35
	CosmosDBParserBIT_XOR_SYMBOL              = 36
	CosmosDBParserEQUAL_SYMBOL                = 37
	CosmosDBParserIDENTIFIER                  = 38
	CosmosDBParserWHITESPACE                  = 39
	CosmosDBParserDECIMAL                     = 40
	CosmosDBParserREAL                        = 41
	CosmosDBParserFLOAT                       = 42
	CosmosDBParserHEXADECIMAL                 = 43
	CosmosDBParserSINGLE_QUOTE_STRING_LITERAL = 44
	CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL = 45
)

// CosmosDBParser rules.
const (
	CosmosDBParserRULE_root                           = 0
	CosmosDBParserRULE_select                         = 1
	CosmosDBParserRULE_select_clause                  = 2
	CosmosDBParserRULE_select_specification           = 3
	CosmosDBParserRULE_from_clause                    = 4
	CosmosDBParserRULE_where_clause                   = 5
	CosmosDBParserRULE_from_specification             = 6
	CosmosDBParserRULE_from_source                    = 7
	CosmosDBParserRULE_container_expression           = 8
	CosmosDBParserRULE_container_name                 = 9
	CosmosDBParserRULE_object_property_list           = 10
	CosmosDBParserRULE_object_property                = 11
	CosmosDBParserRULE_property_alias                 = 12
	CosmosDBParserRULE_scalar_expression              = 13
	CosmosDBParserRULE_scalar_expression_in_where     = 14
	CosmosDBParserRULE_create_array_expression        = 15
	CosmosDBParserRULE_create_object_expression       = 16
	CosmosDBParserRULE_scalar_function_expression     = 17
	CosmosDBParserRULE_udf_scalar_function_expression = 18
	CosmosDBParserRULE_builtin_function_expression    = 19
	CosmosDBParserRULE_binary_operator                = 20
	CosmosDBParserRULE_unary_operator                 = 21
	CosmosDBParserRULE_parameter_name                 = 22
	CosmosDBParserRULE_constant                       = 23
	CosmosDBParserRULE_object_constant                = 24
	CosmosDBParserRULE_object_constant_field_pair     = 25
	CosmosDBParserRULE_array_constant                 = 26
	CosmosDBParserRULE_string_constant                = 27
	CosmosDBParserRULE_undefined_constant             = 28
	CosmosDBParserRULE_null_constant                  = 29
	CosmosDBParserRULE_boolean_constant               = 30
	CosmosDBParserRULE_number_constant                = 31
	CosmosDBParserRULE_string_literal                 = 32
	CosmosDBParserRULE_decimal_literal                = 33
	CosmosDBParserRULE_hexadecimal_literal            = 34
	CosmosDBParserRULE_property_name                  = 35
	CosmosDBParserRULE_array_index                    = 36
	CosmosDBParserRULE_input_alias                    = 37
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_() ISelectContext
	EOF() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Select_() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CosmosDBParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Select_()
	}
	{
		p.SetState(77)
		p.Match(CosmosDBParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectContext is an interface to support dynamic dispatch.
type ISelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_clause() ISelect_clauseContext
	From_clause() IFrom_clauseContext
	Where_clause() IWhere_clauseContext

	// IsSelectContext differentiates from other interfaces.
	IsSelectContext()
}

type SelectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectContext() *SelectContext {
	var p = new(SelectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select
	return p
}

func InitEmptySelectContext(p *SelectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select
}

func (*SelectContext) IsSelectContext() {}

func NewSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectContext {
	var p = new(SelectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select

	return p
}

func (s *SelectContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectContext) Select_clause() ISelect_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_clauseContext)
}

func (s *SelectContext) From_clause() IFrom_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_clauseContext)
}

func (s *SelectContext) Where_clause() IWhere_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhere_clauseContext)
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (s *SelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_() (localctx ISelectContext) {
	localctx = NewSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CosmosDBParserRULE_select)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Select_clause()
	}
	{
		p.SetState(80)
		p.From_clause()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserWHERE_SYMBOL {
		{
			p.SetState(81)
			p.Where_clause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelect_clauseContext is an interface to support dynamic dispatch.
type ISelect_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT_SYMBOL() antlr.TerminalNode
	Select_specification() ISelect_specificationContext

	// IsSelect_clauseContext differentiates from other interfaces.
	IsSelect_clauseContext()
}

type Select_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_clauseContext() *Select_clauseContext {
	var p = new(Select_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_clause
	return p
}

func InitEmptySelect_clauseContext(p *Select_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_clause
}

func (*Select_clauseContext) IsSelect_clauseContext() {}

func NewSelect_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_clauseContext {
	var p = new(Select_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select_clause

	return p
}

func (s *Select_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_clauseContext) SELECT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserSELECT_SYMBOL, 0)
}

func (s *Select_clauseContext) Select_specification() ISelect_specificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_specificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_specificationContext)
}

func (s *Select_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect_clause(s)
	}
}

func (s *Select_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect_clause(s)
	}
}

func (s *Select_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_clause() (localctx ISelect_clauseContext) {
	localctx = NewSelect_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CosmosDBParserRULE_select_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(CosmosDBParserSELECT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Select_specification()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelect_specificationContext is an interface to support dynamic dispatch.
type ISelect_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPLY_OPERATOR() antlr.TerminalNode
	Object_property_list() IObject_property_listContext
	DISTINCT_SYMBOL() antlr.TerminalNode

	// IsSelect_specificationContext differentiates from other interfaces.
	IsSelect_specificationContext()
}

type Select_specificationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_specificationContext() *Select_specificationContext {
	var p = new(Select_specificationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_specification
	return p
}

func InitEmptySelect_specificationContext(p *Select_specificationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_select_specification
}

func (*Select_specificationContext) IsSelect_specificationContext() {}

func NewSelect_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_specificationContext {
	var p = new(Select_specificationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_select_specification

	return p
}

func (s *Select_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_specificationContext) MULTIPLY_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMULTIPLY_OPERATOR, 0)
}

func (s *Select_specificationContext) Object_property_list() IObject_property_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_property_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_property_listContext)
}

func (s *Select_specificationContext) DISTINCT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDISTINCT_SYMBOL, 0)
}

func (s *Select_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSelect_specification(s)
	}
}

func (s *Select_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSelect_specification(s)
	}
}

func (s *Select_specificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSelect_specification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Select_specification() (localctx ISelect_specificationContext) {
	localctx = NewSelect_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CosmosDBParserRULE_select_specification)
	var _la int

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserMULTIPLY_OPERATOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(CosmosDBParserMULTIPLY_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CosmosDBParserDISTINCT_SYMBOL, CosmosDBParserPLUS_SYMBOL, CosmosDBParserMINUS_SYMBOL, CosmosDBParserBIT_NOT_SYMBOL, CosmosDBParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserDISTINCT_SYMBOL {
			{
				p.SetState(88)
				p.Match(CosmosDBParserDISTINCT_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(91)
			p.Object_property_list()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFrom_clauseContext is an interface to support dynamic dispatch.
type IFrom_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM_SYMBOL() antlr.TerminalNode
	From_specification() IFrom_specificationContext

	// IsFrom_clauseContext differentiates from other interfaces.
	IsFrom_clauseContext()
}

type From_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_clauseContext() *From_clauseContext {
	var p = new(From_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_clause
	return p
}

func InitEmptyFrom_clauseContext(p *From_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_clause
}

func (*From_clauseContext) IsFrom_clauseContext() {}

func NewFrom_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_clauseContext {
	var p = new(From_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_clause

	return p
}

func (s *From_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *From_clauseContext) FROM_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserFROM_SYMBOL, 0)
}

func (s *From_clauseContext) From_specification() IFrom_specificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_specificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_specificationContext)
}

func (s *From_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_clause(s)
	}
}

func (s *From_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_clause(s)
	}
}

func (s *From_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_clause() (localctx IFrom_clauseContext) {
	localctx = NewFrom_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CosmosDBParserRULE_from_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(CosmosDBParserFROM_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.From_specification()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhere_clauseContext is an interface to support dynamic dispatch.
type IWhere_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE_SYMBOL() antlr.TerminalNode
	Scalar_expression_in_where() IScalar_expression_in_whereContext

	// IsWhere_clauseContext differentiates from other interfaces.
	IsWhere_clauseContext()
}

type Where_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhere_clauseContext() *Where_clauseContext {
	var p = new(Where_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_where_clause
	return p
}

func InitEmptyWhere_clauseContext(p *Where_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_where_clause
}

func (*Where_clauseContext) IsWhere_clauseContext() {}

func NewWhere_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Where_clauseContext {
	var p = new(Where_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_where_clause

	return p
}

func (s *Where_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Where_clauseContext) WHERE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserWHERE_SYMBOL, 0)
}

func (s *Where_clauseContext) Scalar_expression_in_where() IScalar_expression_in_whereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expression_in_whereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expression_in_whereContext)
}

func (s *Where_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Where_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Where_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterWhere_clause(s)
	}
}

func (s *Where_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitWhere_clause(s)
	}
}

func (s *Where_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitWhere_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Where_clause() (localctx IWhere_clauseContext) {
	localctx = NewWhere_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CosmosDBParserRULE_where_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(CosmosDBParserWHERE_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.scalar_expression_in_where(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFrom_specificationContext is an interface to support dynamic dispatch.
type IFrom_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	From_source() IFrom_sourceContext

	// IsFrom_specificationContext differentiates from other interfaces.
	IsFrom_specificationContext()
}

type From_specificationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_specificationContext() *From_specificationContext {
	var p = new(From_specificationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_specification
	return p
}

func InitEmptyFrom_specificationContext(p *From_specificationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_specification
}

func (*From_specificationContext) IsFrom_specificationContext() {}

func NewFrom_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_specificationContext {
	var p = new(From_specificationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_specification

	return p
}

func (s *From_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *From_specificationContext) From_source() IFrom_sourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_sourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_sourceContext)
}

func (s *From_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_specification(s)
	}
}

func (s *From_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_specification(s)
	}
}

func (s *From_specificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_specification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_specification() (localctx IFrom_specificationContext) {
	localctx = NewFrom_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CosmosDBParserRULE_from_specification)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.From_source()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFrom_sourceContext is an interface to support dynamic dispatch.
type IFrom_sourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Container_expression() IContainer_expressionContext

	// IsFrom_sourceContext differentiates from other interfaces.
	IsFrom_sourceContext()
}

type From_sourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_sourceContext() *From_sourceContext {
	var p = new(From_sourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_source
	return p
}

func InitEmptyFrom_sourceContext(p *From_sourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_from_source
}

func (*From_sourceContext) IsFrom_sourceContext() {}

func NewFrom_sourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_sourceContext {
	var p = new(From_sourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_from_source

	return p
}

func (s *From_sourceContext) GetParser() antlr.Parser { return s.parser }

func (s *From_sourceContext) Container_expression() IContainer_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_expressionContext)
}

func (s *From_sourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_sourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_sourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterFrom_source(s)
	}
}

func (s *From_sourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitFrom_source(s)
	}
}

func (s *From_sourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitFrom_source(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) From_source() (localctx IFrom_sourceContext) {
	localctx = NewFrom_sourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CosmosDBParserRULE_from_source)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Container_expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContainer_expressionContext is an interface to support dynamic dispatch.
type IContainer_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Container_name() IContainer_nameContext
	IDENTIFIER() antlr.TerminalNode
	AS_SYMBOL() antlr.TerminalNode

	// IsContainer_expressionContext differentiates from other interfaces.
	IsContainer_expressionContext()
}

type Container_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_expressionContext() *Container_expressionContext {
	var p = new(Container_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_expression
	return p
}

func InitEmptyContainer_expressionContext(p *Container_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_expression
}

func (*Container_expressionContext) IsContainer_expressionContext() {}

func NewContainer_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_expressionContext {
	var p = new(Container_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_container_expression

	return p
}

func (s *Container_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_expressionContext) Container_name() IContainer_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_nameContext)
}

func (s *Container_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Container_expressionContext) AS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserAS_SYMBOL, 0)
}

func (s *Container_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterContainer_expression(s)
	}
}

func (s *Container_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitContainer_expression(s)
	}
}

func (s *Container_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitContainer_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Container_expression() (localctx IContainer_expressionContext) {
	localctx = NewContainer_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CosmosDBParserRULE_container_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Container_name()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserAS_SYMBOL || _la == CosmosDBParserIDENTIFIER {
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserAS_SYMBOL {
			{
				p.SetState(105)
				p.Match(CosmosDBParserAS_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(108)
			p.Match(CosmosDBParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContainer_nameContext is an interface to support dynamic dispatch.
type IContainer_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsContainer_nameContext differentiates from other interfaces.
	IsContainer_nameContext()
}

type Container_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_nameContext() *Container_nameContext {
	var p = new(Container_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_name
	return p
}

func InitEmptyContainer_nameContext(p *Container_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_container_name
}

func (*Container_nameContext) IsContainer_nameContext() {}

func NewContainer_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_nameContext {
	var p = new(Container_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_container_name

	return p
}

func (s *Container_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Container_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterContainer_name(s)
	}
}

func (s *Container_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitContainer_name(s)
	}
}

func (s *Container_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitContainer_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Container_name() (localctx IContainer_nameContext) {
	localctx = NewContainer_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CosmosDBParserRULE_container_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_property_listContext is an interface to support dynamic dispatch.
type IObject_property_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllObject_property() []IObject_propertyContext
	Object_property(i int) IObject_propertyContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsObject_property_listContext differentiates from other interfaces.
	IsObject_property_listContext()
}

type Object_property_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_property_listContext() *Object_property_listContext {
	var p = new(Object_property_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property_list
	return p
}

func InitEmptyObject_property_listContext(p *Object_property_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property_list
}

func (*Object_property_listContext) IsObject_property_listContext() {}

func NewObject_property_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_property_listContext {
	var p = new(Object_property_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_object_property_list

	return p
}

func (s *Object_property_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_property_listContext) AllObject_property() []IObject_propertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_propertyContext); ok {
			len++
		}
	}

	tst := make([]IObject_propertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_propertyContext); ok {
			tst[i] = t.(IObject_propertyContext)
			i++
		}
	}

	return tst
}

func (s *Object_property_listContext) Object_property(i int) IObject_propertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_propertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_propertyContext)
}

func (s *Object_property_listContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Object_property_listContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Object_property_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_property_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_property_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterObject_property_list(s)
	}
}

func (s *Object_property_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitObject_property_list(s)
	}
}

func (s *Object_property_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitObject_property_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Object_property_list() (localctx IObject_property_listContext) {
	localctx = NewObject_property_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CosmosDBParserRULE_object_property_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Object_property()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(114)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Object_property()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_propertyContext is an interface to support dynamic dispatch.
type IObject_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Scalar_expression() IScalar_expressionContext
	Property_alias() IProperty_aliasContext
	AS_SYMBOL() antlr.TerminalNode

	// IsObject_propertyContext differentiates from other interfaces.
	IsObject_propertyContext()
}

type Object_propertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_propertyContext() *Object_propertyContext {
	var p = new(Object_propertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property
	return p
}

func InitEmptyObject_propertyContext(p *Object_propertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_property
}

func (*Object_propertyContext) IsObject_propertyContext() {}

func NewObject_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_propertyContext {
	var p = new(Object_propertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_object_property

	return p
}

func (s *Object_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_propertyContext) Scalar_expression() IScalar_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expressionContext)
}

func (s *Object_propertyContext) Property_alias() IProperty_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_aliasContext)
}

func (s *Object_propertyContext) AS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserAS_SYMBOL, 0)
}

func (s *Object_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterObject_property(s)
	}
}

func (s *Object_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitObject_property(s)
	}
}

func (s *Object_propertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitObject_property(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Object_property() (localctx IObject_propertyContext) {
	localctx = NewObject_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CosmosDBParserRULE_object_property)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.scalar_expression(0)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserAS_SYMBOL || _la == CosmosDBParserIDENTIFIER {
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserAS_SYMBOL {
			{
				p.SetState(122)
				p.Match(CosmosDBParserAS_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(125)
			p.Property_alias()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProperty_aliasContext is an interface to support dynamic dispatch.
type IProperty_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsProperty_aliasContext differentiates from other interfaces.
	IsProperty_aliasContext()
}

type Property_aliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_aliasContext() *Property_aliasContext {
	var p = new(Property_aliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_alias
	return p
}

func InitEmptyProperty_aliasContext(p *Property_aliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_alias
}

func (*Property_aliasContext) IsProperty_aliasContext() {}

func NewProperty_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_aliasContext {
	var p = new(Property_aliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_property_alias

	return p
}

func (s *Property_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Property_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterProperty_alias(s)
	}
}

func (s *Property_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitProperty_alias(s)
	}
}

func (s *Property_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitProperty_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Property_alias() (localctx IProperty_aliasContext) {
	localctx = NewProperty_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CosmosDBParserRULE_property_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalar_expressionContext is an interface to support dynamic dispatch.
type IScalar_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Input_alias() IInput_aliasContext
	Unary_operator() IUnary_operatorContext
	Scalar_expression() IScalar_expressionContext
	DOT_SYMBOL() antlr.TerminalNode
	Property_name() IProperty_nameContext
	LS_BRACKET_SYMBOL() antlr.TerminalNode
	RS_BRACKET_SYMBOL() antlr.TerminalNode
	DOUBLE_QUOTE_STRING_LITERAL() antlr.TerminalNode
	Array_index() IArray_indexContext

	// IsScalar_expressionContext differentiates from other interfaces.
	IsScalar_expressionContext()
}

type Scalar_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalar_expressionContext() *Scalar_expressionContext {
	var p = new(Scalar_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_expression
	return p
}

func InitEmptyScalar_expressionContext(p *Scalar_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_expression
}

func (*Scalar_expressionContext) IsScalar_expressionContext() {}

func NewScalar_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scalar_expressionContext {
	var p = new(Scalar_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_scalar_expression

	return p
}

func (s *Scalar_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Scalar_expressionContext) Input_alias() IInput_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInput_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInput_aliasContext)
}

func (s *Scalar_expressionContext) Unary_operator() IUnary_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_operatorContext)
}

func (s *Scalar_expressionContext) Scalar_expression() IScalar_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expressionContext)
}

func (s *Scalar_expressionContext) DOT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOT_SYMBOL, 0)
}

func (s *Scalar_expressionContext) Property_name() IProperty_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *Scalar_expressionContext) LS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLS_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expressionContext) RS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRS_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expressionContext) DOUBLE_QUOTE_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL, 0)
}

func (s *Scalar_expressionContext) Array_index() IArray_indexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_indexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_indexContext)
}

func (s *Scalar_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scalar_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scalar_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterScalar_expression(s)
	}
}

func (s *Scalar_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitScalar_expression(s)
	}
}

func (s *Scalar_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitScalar_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Scalar_expression() (localctx IScalar_expressionContext) {
	return p.scalar_expression(0)
}

func (p *CosmosDBParser) scalar_expression(_p int) (localctx IScalar_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewScalar_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScalar_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, CosmosDBParserRULE_scalar_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserIDENTIFIER:
		{
			p.SetState(131)
			p.Input_alias()
		}

	case CosmosDBParserPLUS_SYMBOL, CosmosDBParserMINUS_SYMBOL, CosmosDBParserBIT_NOT_SYMBOL:
		{
			p.SetState(132)
			p.Unary_operator()
		}
		{
			p.SetState(133)
			p.scalar_expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(138)
					p.Match(CosmosDBParserDOT_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(139)
					p.Property_name()
				}

			case 2:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(141)
					p.Match(CosmosDBParserLS_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(144)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
					{
						p.SetState(142)
						p.Match(CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case CosmosDBParserDECIMAL:
					{
						p.SetState(143)
						p.Array_index()
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}
				{
					p.SetState(146)
					p.Match(CosmosDBParserRS_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalar_expression_in_whereContext is an interface to support dynamic dispatch.
type IScalar_expression_in_whereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant() IConstantContext
	Input_alias() IInput_aliasContext
	Parameter_name() IParameter_nameContext
	Unary_operator() IUnary_operatorContext
	AllScalar_expression_in_where() []IScalar_expression_in_whereContext
	Scalar_expression_in_where(i int) IScalar_expression_in_whereContext
	Scalar_function_expression() IScalar_function_expressionContext
	Create_object_expression() ICreate_object_expressionContext
	Create_array_expression() ICreate_array_expressionContext
	LR_BRACKET_SYMBOL() antlr.TerminalNode
	RR_BRACKET_SYMBOL() antlr.TerminalNode
	AND_SYMBOL() antlr.TerminalNode
	OR_SYMBOL() antlr.TerminalNode
	Binary_operator() IBinary_operatorContext
	QUESTION_MARK_SYMBOL() antlr.TerminalNode
	COLON_SYMBOL() antlr.TerminalNode
	DOT_SYMBOL() antlr.TerminalNode
	Property_name() IProperty_nameContext
	LS_BRACKET_SYMBOL() antlr.TerminalNode
	RS_BRACKET_SYMBOL() antlr.TerminalNode
	DOUBLE_QUOTE_STRING_LITERAL() antlr.TerminalNode
	Array_index() IArray_indexContext

	// IsScalar_expression_in_whereContext differentiates from other interfaces.
	IsScalar_expression_in_whereContext()
}

type Scalar_expression_in_whereContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalar_expression_in_whereContext() *Scalar_expression_in_whereContext {
	var p = new(Scalar_expression_in_whereContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_expression_in_where
	return p
}

func InitEmptyScalar_expression_in_whereContext(p *Scalar_expression_in_whereContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_expression_in_where
}

func (*Scalar_expression_in_whereContext) IsScalar_expression_in_whereContext() {}

func NewScalar_expression_in_whereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scalar_expression_in_whereContext {
	var p = new(Scalar_expression_in_whereContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_scalar_expression_in_where

	return p
}

func (s *Scalar_expression_in_whereContext) GetParser() antlr.Parser { return s.parser }

func (s *Scalar_expression_in_whereContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Scalar_expression_in_whereContext) Input_alias() IInput_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInput_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInput_aliasContext)
}

func (s *Scalar_expression_in_whereContext) Parameter_name() IParameter_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameter_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameter_nameContext)
}

func (s *Scalar_expression_in_whereContext) Unary_operator() IUnary_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_operatorContext)
}

func (s *Scalar_expression_in_whereContext) AllScalar_expression_in_where() []IScalar_expression_in_whereContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expression_in_whereContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expression_in_whereContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expression_in_whereContext); ok {
			tst[i] = t.(IScalar_expression_in_whereContext)
			i++
		}
	}

	return tst
}

func (s *Scalar_expression_in_whereContext) Scalar_expression_in_where(i int) IScalar_expression_in_whereContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expression_in_whereContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expression_in_whereContext)
}

func (s *Scalar_expression_in_whereContext) Scalar_function_expression() IScalar_function_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_function_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_function_expressionContext)
}

func (s *Scalar_expression_in_whereContext) Create_object_expression() ICreate_object_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreate_object_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreate_object_expressionContext)
}

func (s *Scalar_expression_in_whereContext) Create_array_expression() ICreate_array_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreate_array_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreate_array_expressionContext)
}

func (s *Scalar_expression_in_whereContext) LR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLR_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) RR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRR_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) AND_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserAND_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) OR_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserOR_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) Binary_operator() IBinary_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinary_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinary_operatorContext)
}

func (s *Scalar_expression_in_whereContext) QUESTION_MARK_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserQUESTION_MARK_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) COLON_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOLON_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) DOT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOT_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) Property_name() IProperty_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *Scalar_expression_in_whereContext) LS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLS_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) RS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRS_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expression_in_whereContext) DOUBLE_QUOTE_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL, 0)
}

func (s *Scalar_expression_in_whereContext) Array_index() IArray_indexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_indexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_indexContext)
}

func (s *Scalar_expression_in_whereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scalar_expression_in_whereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scalar_expression_in_whereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterScalar_expression_in_where(s)
	}
}

func (s *Scalar_expression_in_whereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitScalar_expression_in_where(s)
	}
}

func (s *Scalar_expression_in_whereContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitScalar_expression_in_where(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Scalar_expression_in_where() (localctx IScalar_expression_in_whereContext) {
	return p.scalar_expression_in_where(0)
}

func (p *CosmosDBParser) scalar_expression_in_where(_p int) (localctx IScalar_expression_in_whereContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewScalar_expression_in_whereContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScalar_expression_in_whereContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, CosmosDBParserRULE_scalar_expression_in_where, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(153)
			p.Constant()
		}

	case 2:
		{
			p.SetState(154)
			p.Input_alias()
		}

	case 3:
		{
			p.SetState(155)
			p.Parameter_name()
		}

	case 4:
		{
			p.SetState(156)
			p.Unary_operator()
		}
		{
			p.SetState(157)
			p.scalar_expression_in_where(7)
		}

	case 5:
		{
			p.SetState(159)
			p.Scalar_function_expression()
		}

	case 6:
		{
			p.SetState(160)
			p.Create_object_expression()
		}

	case 7:
		{
			p.SetState(161)
			p.Create_array_expression()
		}

	case 8:
		{
			p.SetState(162)
			p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.scalar_expression_in_where(0)
		}
		{
			p.SetState(164)
			p.Match(CosmosDBParserRR_BRACKET_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewScalar_expression_in_whereContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression_in_where)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(169)
					p.Match(CosmosDBParserAND_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(170)
					p.scalar_expression_in_where(12)
				}

			case 2:
				localctx = NewScalar_expression_in_whereContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression_in_where)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(172)
					p.Match(CosmosDBParserOR_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(173)
					p.scalar_expression_in_where(11)
				}

			case 3:
				localctx = NewScalar_expression_in_whereContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression_in_where)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(175)
					p.Binary_operator()
				}
				{
					p.SetState(176)
					p.scalar_expression_in_where(7)
				}

			case 4:
				localctx = NewScalar_expression_in_whereContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression_in_where)
				p.SetState(178)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(179)
					p.Match(CosmosDBParserQUESTION_MARK_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(180)
					p.scalar_expression_in_where(0)
				}
				{
					p.SetState(181)
					p.Match(CosmosDBParserCOLON_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(182)
					p.scalar_expression_in_where(6)
				}

			case 5:
				localctx = NewScalar_expression_in_whereContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression_in_where)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(185)
					p.Match(CosmosDBParserDOT_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(186)
					p.Property_name()
				}

			case 6:
				localctx = NewScalar_expression_in_whereContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression_in_where)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(188)
					p.Match(CosmosDBParserLS_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(191)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
					{
						p.SetState(189)
						p.Match(CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case CosmosDBParserDECIMAL:
					{
						p.SetState(190)
						p.Array_index()
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}
				{
					p.SetState(193)
					p.Match(CosmosDBParserRS_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreate_array_expressionContext is an interface to support dynamic dispatch.
type ICreate_array_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Array_constant() IArray_constantContext

	// IsCreate_array_expressionContext differentiates from other interfaces.
	IsCreate_array_expressionContext()
}

type Create_array_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_array_expressionContext() *Create_array_expressionContext {
	var p = new(Create_array_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_create_array_expression
	return p
}

func InitEmptyCreate_array_expressionContext(p *Create_array_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_create_array_expression
}

func (*Create_array_expressionContext) IsCreate_array_expressionContext() {}

func NewCreate_array_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_array_expressionContext {
	var p = new(Create_array_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_create_array_expression

	return p
}

func (s *Create_array_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_array_expressionContext) Array_constant() IArray_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_constantContext)
}

func (s *Create_array_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_array_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_array_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterCreate_array_expression(s)
	}
}

func (s *Create_array_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitCreate_array_expression(s)
	}
}

func (s *Create_array_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitCreate_array_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Create_array_expression() (localctx ICreate_array_expressionContext) {
	localctx = NewCreate_array_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CosmosDBParserRULE_create_array_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Array_constant()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreate_object_expressionContext is an interface to support dynamic dispatch.
type ICreate_object_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Object_constant() IObject_constantContext

	// IsCreate_object_expressionContext differentiates from other interfaces.
	IsCreate_object_expressionContext()
}

type Create_object_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_object_expressionContext() *Create_object_expressionContext {
	var p = new(Create_object_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_create_object_expression
	return p
}

func InitEmptyCreate_object_expressionContext(p *Create_object_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_create_object_expression
}

func (*Create_object_expressionContext) IsCreate_object_expressionContext() {}

func NewCreate_object_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_object_expressionContext {
	var p = new(Create_object_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_create_object_expression

	return p
}

func (s *Create_object_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_object_expressionContext) Object_constant() IObject_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_constantContext)
}

func (s *Create_object_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_object_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_object_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterCreate_object_expression(s)
	}
}

func (s *Create_object_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitCreate_object_expression(s)
	}
}

func (s *Create_object_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitCreate_object_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Create_object_expression() (localctx ICreate_object_expressionContext) {
	localctx = NewCreate_object_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CosmosDBParserRULE_create_object_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Object_constant()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalar_function_expressionContext is an interface to support dynamic dispatch.
type IScalar_function_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Udf_scalar_function_expression() IUdf_scalar_function_expressionContext
	Builtin_function_expression() IBuiltin_function_expressionContext

	// IsScalar_function_expressionContext differentiates from other interfaces.
	IsScalar_function_expressionContext()
}

type Scalar_function_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalar_function_expressionContext() *Scalar_function_expressionContext {
	var p = new(Scalar_function_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_function_expression
	return p
}

func InitEmptyScalar_function_expressionContext(p *Scalar_function_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_scalar_function_expression
}

func (*Scalar_function_expressionContext) IsScalar_function_expressionContext() {}

func NewScalar_function_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scalar_function_expressionContext {
	var p = new(Scalar_function_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_scalar_function_expression

	return p
}

func (s *Scalar_function_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Scalar_function_expressionContext) Udf_scalar_function_expression() IUdf_scalar_function_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUdf_scalar_function_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUdf_scalar_function_expressionContext)
}

func (s *Scalar_function_expressionContext) Builtin_function_expression() IBuiltin_function_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltin_function_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltin_function_expressionContext)
}

func (s *Scalar_function_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scalar_function_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scalar_function_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterScalar_function_expression(s)
	}
}

func (s *Scalar_function_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitScalar_function_expression(s)
	}
}

func (s *Scalar_function_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitScalar_function_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Scalar_function_expression() (localctx IScalar_function_expressionContext) {
	localctx = NewScalar_function_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CosmosDBParserRULE_scalar_function_expression)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserUDF_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Udf_scalar_function_expression()
		}

	case CosmosDBParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Builtin_function_expression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUdf_scalar_function_expressionContext is an interface to support dynamic dispatch.
type IUdf_scalar_function_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UDF_SYMBOL() antlr.TerminalNode
	DOT_SYMBOL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LR_BRACKET_SYMBOL() antlr.TerminalNode
	RR_BRACKET_SYMBOL() antlr.TerminalNode
	AllScalar_expression_in_where() []IScalar_expression_in_whereContext
	Scalar_expression_in_where(i int) IScalar_expression_in_whereContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsUdf_scalar_function_expressionContext differentiates from other interfaces.
	IsUdf_scalar_function_expressionContext()
}

type Udf_scalar_function_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUdf_scalar_function_expressionContext() *Udf_scalar_function_expressionContext {
	var p = new(Udf_scalar_function_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_udf_scalar_function_expression
	return p
}

func InitEmptyUdf_scalar_function_expressionContext(p *Udf_scalar_function_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_udf_scalar_function_expression
}

func (*Udf_scalar_function_expressionContext) IsUdf_scalar_function_expressionContext() {}

func NewUdf_scalar_function_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Udf_scalar_function_expressionContext {
	var p = new(Udf_scalar_function_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_udf_scalar_function_expression

	return p
}

func (s *Udf_scalar_function_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Udf_scalar_function_expressionContext) UDF_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserUDF_SYMBOL, 0)
}

func (s *Udf_scalar_function_expressionContext) DOT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOT_SYMBOL, 0)
}

func (s *Udf_scalar_function_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Udf_scalar_function_expressionContext) LR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLR_BRACKET_SYMBOL, 0)
}

func (s *Udf_scalar_function_expressionContext) RR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRR_BRACKET_SYMBOL, 0)
}

func (s *Udf_scalar_function_expressionContext) AllScalar_expression_in_where() []IScalar_expression_in_whereContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expression_in_whereContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expression_in_whereContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expression_in_whereContext); ok {
			tst[i] = t.(IScalar_expression_in_whereContext)
			i++
		}
	}

	return tst
}

func (s *Udf_scalar_function_expressionContext) Scalar_expression_in_where(i int) IScalar_expression_in_whereContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expression_in_whereContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expression_in_whereContext)
}

func (s *Udf_scalar_function_expressionContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Udf_scalar_function_expressionContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Udf_scalar_function_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Udf_scalar_function_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Udf_scalar_function_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterUdf_scalar_function_expression(s)
	}
}

func (s *Udf_scalar_function_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitUdf_scalar_function_expression(s)
	}
}

func (s *Udf_scalar_function_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitUdf_scalar_function_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Udf_scalar_function_expression() (localctx IUdf_scalar_function_expressionContext) {
	localctx = NewUdf_scalar_function_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CosmosDBParserRULE_udf_scalar_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(CosmosDBParserUDF_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(CosmosDBParserDOT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(211)
		p.scalar_expression_in_where(0)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(212)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.scalar_expression_in_where(0)
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(219)
		p.Match(CosmosDBParserRR_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBuiltin_function_expressionContext is an interface to support dynamic dispatch.
type IBuiltin_function_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LR_BRACKET_SYMBOL() antlr.TerminalNode
	RR_BRACKET_SYMBOL() antlr.TerminalNode
	AllScalar_expression_in_where() []IScalar_expression_in_whereContext
	Scalar_expression_in_where(i int) IScalar_expression_in_whereContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsBuiltin_function_expressionContext differentiates from other interfaces.
	IsBuiltin_function_expressionContext()
}

type Builtin_function_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltin_function_expressionContext() *Builtin_function_expressionContext {
	var p = new(Builtin_function_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_builtin_function_expression
	return p
}

func InitEmptyBuiltin_function_expressionContext(p *Builtin_function_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_builtin_function_expression
}

func (*Builtin_function_expressionContext) IsBuiltin_function_expressionContext() {}

func NewBuiltin_function_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Builtin_function_expressionContext {
	var p = new(Builtin_function_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_builtin_function_expression

	return p
}

func (s *Builtin_function_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Builtin_function_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Builtin_function_expressionContext) LR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLR_BRACKET_SYMBOL, 0)
}

func (s *Builtin_function_expressionContext) RR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRR_BRACKET_SYMBOL, 0)
}

func (s *Builtin_function_expressionContext) AllScalar_expression_in_where() []IScalar_expression_in_whereContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expression_in_whereContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expression_in_whereContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expression_in_whereContext); ok {
			tst[i] = t.(IScalar_expression_in_whereContext)
			i++
		}
	}

	return tst
}

func (s *Builtin_function_expressionContext) Scalar_expression_in_where(i int) IScalar_expression_in_whereContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expression_in_whereContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalar_expression_in_whereContext)
}

func (s *Builtin_function_expressionContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Builtin_function_expressionContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Builtin_function_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Builtin_function_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Builtin_function_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterBuiltin_function_expression(s)
	}
}

func (s *Builtin_function_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitBuiltin_function_expression(s)
	}
}

func (s *Builtin_function_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitBuiltin_function_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Builtin_function_expression() (localctx IBuiltin_function_expressionContext) {
	localctx = NewBuiltin_function_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CosmosDBParserRULE_builtin_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(223)
		p.scalar_expression_in_where(0)
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(224)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.scalar_expression_in_where(0)
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(231)
		p.Match(CosmosDBParserRR_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinary_operatorContext is an interface to support dynamic dispatch.
type IBinary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPLY_OPERATOR() antlr.TerminalNode
	DIVIDE_SYMBOL() antlr.TerminalNode
	MODULO_SYMBOL() antlr.TerminalNode
	PLUS_SYMBOL() antlr.TerminalNode
	MINUS_SYMBOL() antlr.TerminalNode
	BIT_AND_SYMBOL() antlr.TerminalNode
	BIT_XOR_SYMBOL() antlr.TerminalNode
	BIT_OR_SYMBOL() antlr.TerminalNode
	DOUBLE_BAR_SYMBOL() antlr.TerminalNode
	EQUAL_SYMBOL() antlr.TerminalNode

	// IsBinary_operatorContext differentiates from other interfaces.
	IsBinary_operatorContext()
}

type Binary_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_operatorContext() *Binary_operatorContext {
	var p = new(Binary_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_binary_operator
	return p
}

func InitEmptyBinary_operatorContext(p *Binary_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_binary_operator
}

func (*Binary_operatorContext) IsBinary_operatorContext() {}

func NewBinary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_operatorContext {
	var p = new(Binary_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_binary_operator

	return p
}

func (s *Binary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_operatorContext) MULTIPLY_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMULTIPLY_OPERATOR, 0)
}

func (s *Binary_operatorContext) DIVIDE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDIVIDE_SYMBOL, 0)
}

func (s *Binary_operatorContext) MODULO_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMODULO_SYMBOL, 0)
}

func (s *Binary_operatorContext) PLUS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserPLUS_SYMBOL, 0)
}

func (s *Binary_operatorContext) MINUS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMINUS_SYMBOL, 0)
}

func (s *Binary_operatorContext) BIT_AND_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBIT_AND_SYMBOL, 0)
}

func (s *Binary_operatorContext) BIT_XOR_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBIT_XOR_SYMBOL, 0)
}

func (s *Binary_operatorContext) BIT_OR_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBIT_OR_SYMBOL, 0)
}

func (s *Binary_operatorContext) DOUBLE_BAR_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOUBLE_BAR_SYMBOL, 0)
}

func (s *Binary_operatorContext) EQUAL_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserEQUAL_SYMBOL, 0)
}

func (s *Binary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterBinary_operator(s)
	}
}

func (s *Binary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitBinary_operator(s)
	}
}

func (s *Binary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitBinary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Binary_operator() (localctx IBinary_operatorContext) {
	localctx = NewBinary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CosmosDBParserRULE_binary_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&273535729666) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnary_operatorContext is an interface to support dynamic dispatch.
type IUnary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BIT_NOT_SYMBOL() antlr.TerminalNode
	PLUS_SYMBOL() antlr.TerminalNode
	MINUS_SYMBOL() antlr.TerminalNode

	// IsUnary_operatorContext differentiates from other interfaces.
	IsUnary_operatorContext()
}

type Unary_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_operatorContext() *Unary_operatorContext {
	var p = new(Unary_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_unary_operator
	return p
}

func InitEmptyUnary_operatorContext(p *Unary_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_unary_operator
}

func (*Unary_operatorContext) IsUnary_operatorContext() {}

func NewUnary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_operatorContext {
	var p = new(Unary_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_unary_operator

	return p
}

func (s *Unary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_operatorContext) BIT_NOT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBIT_NOT_SYMBOL, 0)
}

func (s *Unary_operatorContext) PLUS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserPLUS_SYMBOL, 0)
}

func (s *Unary_operatorContext) MINUS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMINUS_SYMBOL, 0)
}

func (s *Unary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterUnary_operator(s)
	}
}

func (s *Unary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitUnary_operator(s)
	}
}

func (s *Unary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitUnary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Unary_operator() (localctx IUnary_operatorContext) {
	localctx = NewUnary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CosmosDBParserRULE_unary_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1879048192) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameter_nameContext is an interface to support dynamic dispatch.
type IParameter_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT_SYMBOL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsParameter_nameContext differentiates from other interfaces.
	IsParameter_nameContext()
}

type Parameter_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_nameContext() *Parameter_nameContext {
	var p = new(Parameter_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_parameter_name
	return p
}

func InitEmptyParameter_nameContext(p *Parameter_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_parameter_name
}

func (*Parameter_nameContext) IsParameter_nameContext() {}

func NewParameter_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_nameContext {
	var p = new(Parameter_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_parameter_name

	return p
}

func (s *Parameter_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_nameContext) AT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserAT_SYMBOL, 0)
}

func (s *Parameter_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Parameter_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterParameter_name(s)
	}
}

func (s *Parameter_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitParameter_name(s)
	}
}

func (s *Parameter_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitParameter_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Parameter_name() (localctx IParameter_nameContext) {
	localctx = NewParameter_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CosmosDBParserRULE_parameter_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(CosmosDBParserAT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Undefined_constant() IUndefined_constantContext
	Null_constant() INull_constantContext
	Boolean_constant() IBoolean_constantContext
	Number_constant() INumber_constantContext
	String_constant() IString_constantContext
	Array_constant() IArray_constantContext
	Object_constant() IObject_constantContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Undefined_constant() IUndefined_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndefined_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndefined_constantContext)
}

func (s *ConstantContext) Null_constant() INull_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INull_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INull_constantContext)
}

func (s *ConstantContext) Boolean_constant() IBoolean_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_constantContext)
}

func (s *ConstantContext) Number_constant() INumber_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_constantContext)
}

func (s *ConstantContext) String_constant() IString_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_constantContext)
}

func (s *ConstantContext) Array_constant() IArray_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_constantContext)
}

func (s *ConstantContext) Object_constant() IObject_constantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_constantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_constantContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CosmosDBParserRULE_constant)
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserUNDEFINED_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.Undefined_constant()
		}

	case CosmosDBParserNULL_SYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Null_constant()
		}

	case CosmosDBParserFALSE_SYMBOL, CosmosDBParserTRUE_SYMBOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Boolean_constant()
		}

	case CosmosDBParserDECIMAL, CosmosDBParserREAL, CosmosDBParserFLOAT, CosmosDBParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(243)
			p.Number_constant()
		}

	case CosmosDBParserSINGLE_QUOTE_STRING_LITERAL, CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(244)
			p.String_constant()
		}

	case CosmosDBParserLS_BRACKET_SYMBOL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(245)
			p.Array_constant()
		}

	case CosmosDBParserLC_BRACKET_SYMBOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(246)
			p.Object_constant()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_constantContext is an interface to support dynamic dispatch.
type IObject_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC_BRACKET_SYMBOL() antlr.TerminalNode
	RC_BRACKET_SYMBOL() antlr.TerminalNode
	AllObject_constant_field_pair() []IObject_constant_field_pairContext
	Object_constant_field_pair(i int) IObject_constant_field_pairContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsObject_constantContext differentiates from other interfaces.
	IsObject_constantContext()
}

type Object_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_constantContext() *Object_constantContext {
	var p = new(Object_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_constant
	return p
}

func InitEmptyObject_constantContext(p *Object_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_constant
}

func (*Object_constantContext) IsObject_constantContext() {}

func NewObject_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_constantContext {
	var p = new(Object_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_object_constant

	return p
}

func (s *Object_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_constantContext) LC_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLC_BRACKET_SYMBOL, 0)
}

func (s *Object_constantContext) RC_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRC_BRACKET_SYMBOL, 0)
}

func (s *Object_constantContext) AllObject_constant_field_pair() []IObject_constant_field_pairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_constant_field_pairContext); ok {
			len++
		}
	}

	tst := make([]IObject_constant_field_pairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_constant_field_pairContext); ok {
			tst[i] = t.(IObject_constant_field_pairContext)
			i++
		}
	}

	return tst
}

func (s *Object_constantContext) Object_constant_field_pair(i int) IObject_constant_field_pairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_constant_field_pairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_constant_field_pairContext)
}

func (s *Object_constantContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Object_constantContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Object_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterObject_constant(s)
	}
}

func (s *Object_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitObject_constant(s)
	}
}

func (s *Object_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitObject_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Object_constant() (localctx IObject_constantContext) {
	localctx = NewObject_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CosmosDBParserRULE_object_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(CosmosDBParserLC_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(250)
		p.Object_constant_field_pair()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(251)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Object_constant_field_pair()
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(258)
		p.Match(CosmosDBParserRC_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_constant_field_pairContext is an interface to support dynamic dispatch.
type IObject_constant_field_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA_SYMBOL() antlr.TerminalNode
	Constant() IConstantContext
	Property_name() IProperty_nameContext
	AllDOUBLE_QUOTE_SYMBOL() []antlr.TerminalNode
	DOUBLE_QUOTE_SYMBOL(i int) antlr.TerminalNode

	// IsObject_constant_field_pairContext differentiates from other interfaces.
	IsObject_constant_field_pairContext()
}

type Object_constant_field_pairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_constant_field_pairContext() *Object_constant_field_pairContext {
	var p = new(Object_constant_field_pairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_constant_field_pair
	return p
}

func InitEmptyObject_constant_field_pairContext(p *Object_constant_field_pairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_constant_field_pair
}

func (*Object_constant_field_pairContext) IsObject_constant_field_pairContext() {}

func NewObject_constant_field_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_constant_field_pairContext {
	var p = new(Object_constant_field_pairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_object_constant_field_pair

	return p
}

func (s *Object_constant_field_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_constant_field_pairContext) COMMA_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, 0)
}

func (s *Object_constant_field_pairContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Object_constant_field_pairContext) Property_name() IProperty_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_nameContext)
}

func (s *Object_constant_field_pairContext) AllDOUBLE_QUOTE_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserDOUBLE_QUOTE_SYMBOL)
}

func (s *Object_constant_field_pairContext) DOUBLE_QUOTE_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOUBLE_QUOTE_SYMBOL, i)
}

func (s *Object_constant_field_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_constant_field_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_constant_field_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterObject_constant_field_pair(s)
	}
}

func (s *Object_constant_field_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitObject_constant_field_pair(s)
	}
}

func (s *Object_constant_field_pairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitObject_constant_field_pair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Object_constant_field_pair() (localctx IObject_constant_field_pairContext) {
	localctx = NewObject_constant_field_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CosmosDBParserRULE_object_constant_field_pair)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserIDENTIFIER:
		{
			p.SetState(260)
			p.Property_name()
		}

	case CosmosDBParserDOUBLE_QUOTE_SYMBOL:
		{
			p.SetState(261)
			p.Match(CosmosDBParserDOUBLE_QUOTE_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Property_name()
		}
		{
			p.SetState(263)
			p.Match(CosmosDBParserDOUBLE_QUOTE_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(267)
		p.Match(CosmosDBParserCOMMA_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Constant()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_constantContext is an interface to support dynamic dispatch.
type IArray_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LS_BRACKET_SYMBOL() antlr.TerminalNode
	RS_BRACKET_SYMBOL() antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsArray_constantContext differentiates from other interfaces.
	IsArray_constantContext()
}

type Array_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_constantContext() *Array_constantContext {
	var p = new(Array_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_array_constant
	return p
}

func InitEmptyArray_constantContext(p *Array_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_array_constant
}

func (*Array_constantContext) IsArray_constantContext() {}

func NewArray_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_constantContext {
	var p = new(Array_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_array_constant

	return p
}

func (s *Array_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_constantContext) LS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLS_BRACKET_SYMBOL, 0)
}

func (s *Array_constantContext) RS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRS_BRACKET_SYMBOL, 0)
}

func (s *Array_constantContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *Array_constantContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *Array_constantContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Array_constantContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Array_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterArray_constant(s)
	}
}

func (s *Array_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitArray_constant(s)
	}
}

func (s *Array_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitArray_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Array_constant() (localctx IArray_constantContext) {
	localctx = NewArray_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CosmosDBParserRULE_array_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(CosmosDBParserLS_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69269232878528) != 0 {
		{
			p.SetState(271)
			p.Constant()
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CosmosDBParserCOMMA_SYMBOL {
			{
				p.SetState(272)
				p.Match(CosmosDBParserCOMMA_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(273)
				p.Constant()
			}

			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(281)
		p.Match(CosmosDBParserRS_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_constantContext is an interface to support dynamic dispatch.
type IString_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_literal() IString_literalContext

	// IsString_constantContext differentiates from other interfaces.
	IsString_constantContext()
}

type String_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_constantContext() *String_constantContext {
	var p = new(String_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_string_constant
	return p
}

func InitEmptyString_constantContext(p *String_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_string_constant
}

func (*String_constantContext) IsString_constantContext() {}

func NewString_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_constantContext {
	var p = new(String_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_string_constant

	return p
}

func (s *String_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *String_constantContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *String_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterString_constant(s)
	}
}

func (s *String_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitString_constant(s)
	}
}

func (s *String_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitString_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) String_constant() (localctx IString_constantContext) {
	localctx = NewString_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CosmosDBParserRULE_string_constant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.String_literal()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUndefined_constantContext is an interface to support dynamic dispatch.
type IUndefined_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNDEFINED_SYMBOL() antlr.TerminalNode

	// IsUndefined_constantContext differentiates from other interfaces.
	IsUndefined_constantContext()
}

type Undefined_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUndefined_constantContext() *Undefined_constantContext {
	var p = new(Undefined_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_undefined_constant
	return p
}

func InitEmptyUndefined_constantContext(p *Undefined_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_undefined_constant
}

func (*Undefined_constantContext) IsUndefined_constantContext() {}

func NewUndefined_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Undefined_constantContext {
	var p = new(Undefined_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_undefined_constant

	return p
}

func (s *Undefined_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Undefined_constantContext) UNDEFINED_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserUNDEFINED_SYMBOL, 0)
}

func (s *Undefined_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Undefined_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Undefined_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterUndefined_constant(s)
	}
}

func (s *Undefined_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitUndefined_constant(s)
	}
}

func (s *Undefined_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitUndefined_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Undefined_constant() (localctx IUndefined_constantContext) {
	localctx = NewUndefined_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CosmosDBParserRULE_undefined_constant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(CosmosDBParserUNDEFINED_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INull_constantContext is an interface to support dynamic dispatch.
type INull_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NULL_SYMBOL() antlr.TerminalNode

	// IsNull_constantContext differentiates from other interfaces.
	IsNull_constantContext()
}

type Null_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNull_constantContext() *Null_constantContext {
	var p = new(Null_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_null_constant
	return p
}

func InitEmptyNull_constantContext(p *Null_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_null_constant
}

func (*Null_constantContext) IsNull_constantContext() {}

func NewNull_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Null_constantContext {
	var p = new(Null_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_null_constant

	return p
}

func (s *Null_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Null_constantContext) NULL_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserNULL_SYMBOL, 0)
}

func (s *Null_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Null_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Null_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterNull_constant(s)
	}
}

func (s *Null_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitNull_constant(s)
	}
}

func (s *Null_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitNull_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Null_constant() (localctx INull_constantContext) {
	localctx = NewNull_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CosmosDBParserRULE_null_constant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(CosmosDBParserNULL_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_constantContext is an interface to support dynamic dispatch.
type IBoolean_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE_SYMBOL() antlr.TerminalNode
	FALSE_SYMBOL() antlr.TerminalNode

	// IsBoolean_constantContext differentiates from other interfaces.
	IsBoolean_constantContext()
}

type Boolean_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_constantContext() *Boolean_constantContext {
	var p = new(Boolean_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_boolean_constant
	return p
}

func InitEmptyBoolean_constantContext(p *Boolean_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_boolean_constant
}

func (*Boolean_constantContext) IsBoolean_constantContext() {}

func NewBoolean_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_constantContext {
	var p = new(Boolean_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_boolean_constant

	return p
}

func (s *Boolean_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_constantContext) TRUE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserTRUE_SYMBOL, 0)
}

func (s *Boolean_constantContext) FALSE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserFALSE_SYMBOL, 0)
}

func (s *Boolean_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterBoolean_constant(s)
	}
}

func (s *Boolean_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitBoolean_constant(s)
	}
}

func (s *Boolean_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitBoolean_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Boolean_constant() (localctx IBoolean_constantContext) {
	localctx = NewBoolean_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CosmosDBParserRULE_boolean_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CosmosDBParserFALSE_SYMBOL || _la == CosmosDBParserTRUE_SYMBOL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumber_constantContext is an interface to support dynamic dispatch.
type INumber_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decimal_literal() IDecimal_literalContext
	Hexadecimal_literal() IHexadecimal_literalContext

	// IsNumber_constantContext differentiates from other interfaces.
	IsNumber_constantContext()
}

type Number_constantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_constantContext() *Number_constantContext {
	var p = new(Number_constantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_number_constant
	return p
}

func InitEmptyNumber_constantContext(p *Number_constantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_number_constant
}

func (*Number_constantContext) IsNumber_constantContext() {}

func NewNumber_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_constantContext {
	var p = new(Number_constantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_number_constant

	return p
}

func (s *Number_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_constantContext) Decimal_literal() IDecimal_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimal_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimal_literalContext)
}

func (s *Number_constantContext) Hexadecimal_literal() IHexadecimal_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHexadecimal_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHexadecimal_literalContext)
}

func (s *Number_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterNumber_constant(s)
	}
}

func (s *Number_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitNumber_constant(s)
	}
}

func (s *Number_constantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitNumber_constant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Number_constant() (localctx INumber_constantContext) {
	localctx = NewNumber_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CosmosDBParserRULE_number_constant)
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserDECIMAL, CosmosDBParserREAL, CosmosDBParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.Decimal_literal()
		}

	case CosmosDBParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Hexadecimal_literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SINGLE_QUOTE_STRING_LITERAL() antlr.TerminalNode
	DOUBLE_QUOTE_STRING_LITERAL() antlr.TerminalNode

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_string_literal
	return p
}

func InitEmptyString_literalContext(p *String_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_string_literal
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) SINGLE_QUOTE_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserSINGLE_QUOTE_STRING_LITERAL, 0)
}

func (s *String_literalContext) DOUBLE_QUOTE_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (s *String_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitString_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) String_literal() (localctx IString_literalContext) {
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CosmosDBParserRULE_string_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CosmosDBParserSINGLE_QUOTE_STRING_LITERAL || _la == CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecimal_literalContext is an interface to support dynamic dispatch.
type IDecimal_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL() antlr.TerminalNode
	REAL() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsDecimal_literalContext differentiates from other interfaces.
	IsDecimal_literalContext()
}

type Decimal_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimal_literalContext() *Decimal_literalContext {
	var p = new(Decimal_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_decimal_literal
	return p
}

func InitEmptyDecimal_literalContext(p *Decimal_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_decimal_literal
}

func (*Decimal_literalContext) IsDecimal_literalContext() {}

func NewDecimal_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decimal_literalContext {
	var p = new(Decimal_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_decimal_literal

	return p
}

func (s *Decimal_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Decimal_literalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDECIMAL, 0)
}

func (s *Decimal_literalContext) REAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserREAL, 0)
}

func (s *Decimal_literalContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserFLOAT, 0)
}

func (s *Decimal_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decimal_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decimal_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterDecimal_literal(s)
	}
}

func (s *Decimal_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitDecimal_literal(s)
	}
}

func (s *Decimal_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitDecimal_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Decimal_literal() (localctx IDecimal_literalContext) {
	localctx = NewDecimal_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CosmosDBParserRULE_decimal_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7696581394432) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHexadecimal_literalContext is an interface to support dynamic dispatch.
type IHexadecimal_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HEXADECIMAL() antlr.TerminalNode

	// IsHexadecimal_literalContext differentiates from other interfaces.
	IsHexadecimal_literalContext()
}

type Hexadecimal_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexadecimal_literalContext() *Hexadecimal_literalContext {
	var p = new(Hexadecimal_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_hexadecimal_literal
	return p
}

func InitEmptyHexadecimal_literalContext(p *Hexadecimal_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_hexadecimal_literal
}

func (*Hexadecimal_literalContext) IsHexadecimal_literalContext() {}

func NewHexadecimal_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hexadecimal_literalContext {
	var p = new(Hexadecimal_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_hexadecimal_literal

	return p
}

func (s *Hexadecimal_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Hexadecimal_literalContext) HEXADECIMAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserHEXADECIMAL, 0)
}

func (s *Hexadecimal_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hexadecimal_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hexadecimal_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterHexadecimal_literal(s)
	}
}

func (s *Hexadecimal_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitHexadecimal_literal(s)
	}
}

func (s *Hexadecimal_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitHexadecimal_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Hexadecimal_literal() (localctx IHexadecimal_literalContext) {
	localctx = NewHexadecimal_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CosmosDBParserRULE_hexadecimal_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(CosmosDBParserHEXADECIMAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProperty_nameContext is an interface to support dynamic dispatch.
type IProperty_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsProperty_nameContext differentiates from other interfaces.
	IsProperty_nameContext()
}

type Property_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_nameContext() *Property_nameContext {
	var p = new(Property_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_name
	return p
}

func InitEmptyProperty_nameContext(p *Property_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_property_name
}

func (*Property_nameContext) IsProperty_nameContext() {}

func NewProperty_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_nameContext {
	var p = new(Property_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_property_name

	return p
}

func (s *Property_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Property_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterProperty_name(s)
	}
}

func (s *Property_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitProperty_name(s)
	}
}

func (s *Property_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitProperty_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Property_name() (localctx IProperty_nameContext) {
	localctx = NewProperty_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CosmosDBParserRULE_property_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_indexContext is an interface to support dynamic dispatch.
type IArray_indexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL() antlr.TerminalNode

	// IsArray_indexContext differentiates from other interfaces.
	IsArray_indexContext()
}

type Array_indexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_indexContext() *Array_indexContext {
	var p = new(Array_indexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_array_index
	return p
}

func InitEmptyArray_indexContext(p *Array_indexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_array_index
}

func (*Array_indexContext) IsArray_indexContext() {}

func NewArray_indexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_indexContext {
	var p = new(Array_indexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_array_index

	return p
}

func (s *Array_indexContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_indexContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDECIMAL, 0)
}

func (s *Array_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_indexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterArray_index(s)
	}
}

func (s *Array_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitArray_index(s)
	}
}

func (s *Array_indexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitArray_index(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Array_index() (localctx IArray_indexContext) {
	localctx = NewArray_indexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CosmosDBParserRULE_array_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(CosmosDBParserDECIMAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInput_aliasContext is an interface to support dynamic dispatch.
type IInput_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsInput_aliasContext differentiates from other interfaces.
	IsInput_aliasContext()
}

type Input_aliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInput_aliasContext() *Input_aliasContext {
	var p = new(Input_aliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_input_alias
	return p
}

func InitEmptyInput_aliasContext(p *Input_aliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_input_alias
}

func (*Input_aliasContext) IsInput_aliasContext() {}

func NewInput_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Input_aliasContext {
	var p = new(Input_aliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_input_alias

	return p
}

func (s *Input_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Input_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *Input_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Input_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Input_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterInput_alias(s)
	}
}

func (s *Input_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitInput_alias(s)
	}
}

func (s *Input_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitInput_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Input_alias() (localctx IInput_aliasContext) {
	localctx = NewInput_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CosmosDBParserRULE_input_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(CosmosDBParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CosmosDBParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *Scalar_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Scalar_expressionContext)
		}
		return p.Scalar_expression_Sempred(t, predIndex)

	case 14:
		var t *Scalar_expression_in_whereContext = nil
		if localctx != nil {
			t = localctx.(*Scalar_expression_in_whereContext)
		}
		return p.Scalar_expression_in_where_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CosmosDBParser) Scalar_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CosmosDBParser) Scalar_expression_in_where_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
