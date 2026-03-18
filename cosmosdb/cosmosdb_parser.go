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
		"'OR'", "'IN'", "'BETWEEN'", "'TOP'", "'VALUE'", "'ORDER'", "'BY'",
		"'GROUP'", "'OFFSET'", "'LIMIT'", "'ASC'", "'DESC'", "'EXISTS'", "'LIKE'",
		"'HAVING'", "'JOIN'", "'@'", "'{'", "'}'", "'['", "']'", "'('", "')'",
		"'''", "'\"'", "','", "'.'", "'?'", "':'", "'+'", "'-'", "'~'", "'/'",
		"'%'", "'&'", "'|'", "'||'", "'^'", "'='", "'<'", "'<='", "'>'", "'>='",
		"'<<'", "'>>'", "'>>>'", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "MULTIPLY_OPERATOR", "AS_SYMBOL", "SELECT_SYMBOL", "FROM_SYMBOL",
		"DISTINCT_SYMBOL", "UNDEFINED_SYMBOL", "NULL_SYMBOL", "FALSE_SYMBOL",
		"TRUE_SYMBOL", "NOT_SYMBOL", "UDF_SYMBOL", "WHERE_SYMBOL", "AND_SYMBOL",
		"OR_SYMBOL", "IN_SYMBOL", "BETWEEN_SYMBOL", "TOP_SYMBOL", "VALUE_SYMBOL",
		"ORDER_SYMBOL", "BY_SYMBOL", "GROUP_SYMBOL", "OFFSET_SYMBOL", "LIMIT_SYMBOL",
		"ASC_SYMBOL", "DESC_SYMBOL", "EXISTS_SYMBOL", "LIKE_SYMBOL", "HAVING_SYMBOL",
		"JOIN_SYMBOL", "AT_SYMBOL", "LC_BRACKET_SYMBOL", "RC_BRACKET_SYMBOL",
		"LS_BRACKET_SYMBOL", "RS_BRACKET_SYMBOL", "LR_BRACKET_SYMBOL", "RR_BRACKET_SYMBOL",
		"SINGLE_QUOTE_SYMBOL", "DOUBLE_QUOTE_SYMBOL", "COMMA_SYMBOL", "DOT_SYMBOL",
		"QUESTION_MARK_SYMBOL", "COLON_SYMBOL", "PLUS_SYMBOL", "MINUS_SYMBOL",
		"BIT_NOT_SYMBOL", "DIVIDE_SYMBOL", "MODULO_SYMBOL", "BIT_AND_SYMBOL",
		"BIT_OR_SYMBOL", "DOUBLE_BAR_SYMBOL", "BIT_XOR_SYMBOL", "EQUAL_SYMBOL",
		"LESS_THAN_OPERATOR", "LESS_THAN_EQUAL_OPERATOR", "GREATER_THAN_OPERATOR",
		"GREATER_THAN_EQUAL_OPERATOR", "LEFT_SHIFT_OPERATOR", "RIGHT_SHIFT_OPERATOR",
		"ZERO_FILL_RIGHT_SHIFT_OPERATOR", "NOT_EQUAL_OPERATOR", "IDENTIFIER",
		"WHITESPACE", "DECIMAL", "REAL", "FLOAT", "HEXADECIMAL", "SINGLE_QUOTE_STRING_LITERAL",
		"DOUBLE_QUOTE_STRING_LITERAL",
	}
	staticData.RuleNames = []string{
		"root", "select", "select_clause", "top_clause", "select_specification",
		"from_clause", "where_clause", "group_by_clause", "having_clause", "order_by_clause",
		"sort_expression", "offset_limit_clause", "from_specification", "from_source",
		"container_expression", "join_clause", "container_name", "object_property_list",
		"object_property", "property_alias", "scalar_expression", "create_array_expression",
		"create_object_expression", "object_field_pair", "scalar_function_expression",
		"udf_scalar_function_expression", "builtin_function_expression", "binary_operator",
		"unary_operator", "parameter_name", "constant", "undefined_constant",
		"null_constant", "boolean_constant", "number_constant", "string_constant",
		"string_literal", "decimal_literal", "hexadecimal_literal", "identifier",
		"property_name", "array_index", "input_alias",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 407, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 92, 8, 1, 1, 1, 3, 1, 95,
		8, 1, 1, 1, 3, 1, 98, 8, 1, 1, 1, 3, 1, 101, 8, 1, 1, 1, 3, 1, 104, 8,
		1, 1, 1, 3, 1, 107, 8, 1, 1, 2, 1, 2, 3, 2, 111, 8, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 120, 8, 4, 1, 4, 3, 4, 123, 8, 4, 1, 4, 3,
		4, 126, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 139, 8, 7, 10, 7, 12, 7, 142, 9, 7, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 152, 8, 9, 10, 9, 12, 9, 155, 9, 9, 1,
		10, 1, 10, 3, 10, 159, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 5, 13, 170, 8, 13, 10, 13, 12, 13, 173, 9, 13, 1,
		14, 1, 14, 3, 14, 177, 8, 14, 1, 14, 3, 14, 180, 8, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 192, 8, 17,
		10, 17, 12, 17, 195, 9, 17, 1, 18, 1, 18, 3, 18, 199, 8, 18, 1, 18, 3,
		18, 202, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 231,
		8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 245, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 3, 20, 254, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		3, 20, 272, 8, 20, 1, 20, 1, 20, 1, 20, 3, 20, 277, 8, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 284, 8, 20, 10, 20, 12, 20, 287, 9, 20, 3,
		20, 289, 8, 20, 1, 20, 5, 20, 292, 8, 20, 10, 20, 12, 20, 295, 9, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 5, 21, 301, 8, 21, 10, 21, 12, 21, 304, 9, 21,
		3, 21, 306, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 314,
		8, 22, 10, 22, 12, 22, 317, 9, 22, 3, 22, 319, 8, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 3, 23, 325, 8, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24,
		332, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 341,
		8, 25, 10, 25, 12, 25, 344, 9, 25, 3, 25, 346, 8, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 3, 26, 354, 8, 26, 1, 26, 1, 26, 5, 26, 358, 8,
		26, 10, 26, 12, 26, 361, 9, 26, 3, 26, 363, 8, 26, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 3, 30, 379, 8, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34,
		1, 34, 3, 34, 389, 8, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42,
		0, 1, 40, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 7, 1, 0, 24, 25, 3, 0, 1, 1, 43,
		44, 46, 60, 1, 0, 43, 45, 1, 0, 8, 9, 1, 0, 67, 68, 1, 0, 63, 65, 2, 0,
		15, 29, 61, 61, 424, 0, 86, 1, 0, 0, 0, 2, 89, 1, 0, 0, 0, 4, 108, 1, 0,
		0, 0, 6, 114, 1, 0, 0, 0, 8, 125, 1, 0, 0, 0, 10, 127, 1, 0, 0, 0, 12,
		130, 1, 0, 0, 0, 14, 133, 1, 0, 0, 0, 16, 143, 1, 0, 0, 0, 18, 146, 1,
		0, 0, 0, 20, 156, 1, 0, 0, 0, 22, 160, 1, 0, 0, 0, 24, 165, 1, 0, 0, 0,
		26, 167, 1, 0, 0, 0, 28, 174, 1, 0, 0, 0, 30, 181, 1, 0, 0, 0, 32, 186,
		1, 0, 0, 0, 34, 188, 1, 0, 0, 0, 36, 196, 1, 0, 0, 0, 38, 203, 1, 0, 0,
		0, 40, 230, 1, 0, 0, 0, 42, 296, 1, 0, 0, 0, 44, 309, 1, 0, 0, 0, 46, 324,
		1, 0, 0, 0, 48, 331, 1, 0, 0, 0, 50, 333, 1, 0, 0, 0, 52, 349, 1, 0, 0,
		0, 54, 366, 1, 0, 0, 0, 56, 368, 1, 0, 0, 0, 58, 370, 1, 0, 0, 0, 60, 378,
		1, 0, 0, 0, 62, 380, 1, 0, 0, 0, 64, 382, 1, 0, 0, 0, 66, 384, 1, 0, 0,
		0, 68, 388, 1, 0, 0, 0, 70, 390, 1, 0, 0, 0, 72, 392, 1, 0, 0, 0, 74, 394,
		1, 0, 0, 0, 76, 396, 1, 0, 0, 0, 78, 398, 1, 0, 0, 0, 80, 400, 1, 0, 0,
		0, 82, 402, 1, 0, 0, 0, 84, 404, 1, 0, 0, 0, 86, 87, 3, 2, 1, 0, 87, 88,
		5, 0, 0, 1, 88, 1, 1, 0, 0, 0, 89, 91, 3, 4, 2, 0, 90, 92, 3, 10, 5, 0,
		91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 95, 3,
		12, 6, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96,
		98, 3, 14, 7, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 100, 1, 0,
		0, 0, 99, 101, 3, 16, 8, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0,
		101, 103, 1, 0, 0, 0, 102, 104, 3, 18, 9, 0, 103, 102, 1, 0, 0, 0, 103,
		104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 107, 3, 22, 11, 0, 106, 105,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 3, 1, 0, 0, 0, 108, 110, 5, 3, 0,
		0, 109, 111, 3, 6, 3, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111,
		112, 1, 0, 0, 0, 112, 113, 3, 8, 4, 0, 113, 5, 1, 0, 0, 0, 114, 115, 5,
		17, 0, 0, 115, 116, 5, 63, 0, 0, 116, 7, 1, 0, 0, 0, 117, 126, 5, 1, 0,
		0, 118, 120, 5, 5, 0, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		122, 1, 0, 0, 0, 121, 123, 5, 18, 0, 0, 122, 121, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 3, 34, 17, 0, 125, 117, 1,
		0, 0, 0, 125, 119, 1, 0, 0, 0, 126, 9, 1, 0, 0, 0, 127, 128, 5, 4, 0, 0,
		128, 129, 3, 24, 12, 0, 129, 11, 1, 0, 0, 0, 130, 131, 5, 12, 0, 0, 131,
		132, 3, 40, 20, 0, 132, 13, 1, 0, 0, 0, 133, 134, 5, 21, 0, 0, 134, 135,
		5, 20, 0, 0, 135, 140, 3, 40, 20, 0, 136, 137, 5, 39, 0, 0, 137, 139, 3,
		40, 20, 0, 138, 136, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0,
		0, 0, 140, 141, 1, 0, 0, 0, 141, 15, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0,
		143, 144, 5, 28, 0, 0, 144, 145, 3, 40, 20, 0, 145, 17, 1, 0, 0, 0, 146,
		147, 5, 19, 0, 0, 147, 148, 5, 20, 0, 0, 148, 153, 3, 20, 10, 0, 149, 150,
		5, 39, 0, 0, 150, 152, 3, 20, 10, 0, 151, 149, 1, 0, 0, 0, 152, 155, 1,
		0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 19, 1, 0, 0,
		0, 155, 153, 1, 0, 0, 0, 156, 158, 3, 40, 20, 0, 157, 159, 7, 0, 0, 0,
		158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 21, 1, 0, 0, 0, 160, 161,
		5, 22, 0, 0, 161, 162, 5, 63, 0, 0, 162, 163, 5, 23, 0, 0, 163, 164, 5,
		63, 0, 0, 164, 23, 1, 0, 0, 0, 165, 166, 3, 26, 13, 0, 166, 25, 1, 0, 0,
		0, 167, 171, 3, 28, 14, 0, 168, 170, 3, 30, 15, 0, 169, 168, 1, 0, 0, 0,
		170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		27, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 179, 3, 32, 16, 0, 175, 177,
		5, 2, 0, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0,
		0, 0, 178, 180, 3, 78, 39, 0, 179, 176, 1, 0, 0, 0, 179, 180, 1, 0, 0,
		0, 180, 29, 1, 0, 0, 0, 181, 182, 5, 29, 0, 0, 182, 183, 3, 78, 39, 0,
		183, 184, 5, 15, 0, 0, 184, 185, 3, 40, 20, 0, 185, 31, 1, 0, 0, 0, 186,
		187, 3, 78, 39, 0, 187, 33, 1, 0, 0, 0, 188, 193, 3, 36, 18, 0, 189, 190,
		5, 39, 0, 0, 190, 192, 3, 36, 18, 0, 191, 189, 1, 0, 0, 0, 192, 195, 1,
		0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 35, 1, 0, 0,
		0, 195, 193, 1, 0, 0, 0, 196, 201, 3, 40, 20, 0, 197, 199, 5, 2, 0, 0,
		198, 197, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200,
		202, 3, 38, 19, 0, 201, 198, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 37,
		1, 0, 0, 0, 203, 204, 3, 78, 39, 0, 204, 39, 1, 0, 0, 0, 205, 206, 6, 20,
		-1, 0, 206, 231, 3, 60, 30, 0, 207, 231, 3, 84, 42, 0, 208, 231, 3, 58,
		29, 0, 209, 210, 3, 56, 28, 0, 210, 211, 3, 40, 20, 13, 211, 231, 1, 0,
		0, 0, 212, 213, 5, 10, 0, 0, 213, 231, 3, 40, 20, 12, 214, 215, 5, 26,
		0, 0, 215, 216, 5, 35, 0, 0, 216, 217, 3, 2, 1, 0, 217, 218, 5, 36, 0,
		0, 218, 231, 1, 0, 0, 0, 219, 231, 3, 48, 24, 0, 220, 231, 3, 44, 22, 0,
		221, 231, 3, 42, 21, 0, 222, 223, 5, 35, 0, 0, 223, 224, 3, 40, 20, 0,
		224, 225, 5, 36, 0, 0, 225, 231, 1, 0, 0, 0, 226, 227, 5, 35, 0, 0, 227,
		228, 3, 2, 1, 0, 228, 229, 5, 36, 0, 0, 229, 231, 1, 0, 0, 0, 230, 205,
		1, 0, 0, 0, 230, 207, 1, 0, 0, 0, 230, 208, 1, 0, 0, 0, 230, 209, 1, 0,
		0, 0, 230, 212, 1, 0, 0, 0, 230, 214, 1, 0, 0, 0, 230, 219, 1, 0, 0, 0,
		230, 220, 1, 0, 0, 0, 230, 221, 1, 0, 0, 0, 230, 222, 1, 0, 0, 0, 230,
		226, 1, 0, 0, 0, 231, 293, 1, 0, 0, 0, 232, 233, 10, 17, 0, 0, 233, 234,
		5, 13, 0, 0, 234, 292, 3, 40, 20, 18, 235, 236, 10, 16, 0, 0, 236, 237,
		5, 14, 0, 0, 237, 292, 3, 40, 20, 17, 238, 239, 10, 11, 0, 0, 239, 240,
		3, 54, 27, 0, 240, 241, 3, 40, 20, 12, 241, 292, 1, 0, 0, 0, 242, 244,
		10, 9, 0, 0, 243, 245, 5, 10, 0, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1,
		0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 5, 16, 0, 0, 247, 248, 3, 40,
		20, 0, 248, 249, 5, 13, 0, 0, 249, 250, 3, 40, 20, 10, 250, 292, 1, 0,
		0, 0, 251, 253, 10, 8, 0, 0, 252, 254, 5, 10, 0, 0, 253, 252, 1, 0, 0,
		0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 5, 27, 0, 0, 256,
		292, 3, 40, 20, 9, 257, 258, 10, 6, 0, 0, 258, 259, 5, 41, 0, 0, 259, 260,
		3, 40, 20, 0, 260, 261, 5, 42, 0, 0, 261, 262, 3, 40, 20, 7, 262, 292,
		1, 0, 0, 0, 263, 264, 10, 15, 0, 0, 264, 265, 5, 40, 0, 0, 265, 292, 3,
		80, 40, 0, 266, 267, 10, 14, 0, 0, 267, 271, 5, 33, 0, 0, 268, 272, 5,
		68, 0, 0, 269, 272, 5, 67, 0, 0, 270, 272, 3, 82, 41, 0, 271, 268, 1, 0,
		0, 0, 271, 269, 1, 0, 0, 0, 271, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0,
		273, 292, 5, 34, 0, 0, 274, 276, 10, 10, 0, 0, 275, 277, 5, 10, 0, 0, 276,
		275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 279,
		5, 15, 0, 0, 279, 288, 5, 35, 0, 0, 280, 285, 3, 40, 20, 0, 281, 282, 5,
		39, 0, 0, 282, 284, 3, 40, 20, 0, 283, 281, 1, 0, 0, 0, 284, 287, 1, 0,
		0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0,
		287, 285, 1, 0, 0, 0, 288, 280, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		290, 1, 0, 0, 0, 290, 292, 5, 36, 0, 0, 291, 232, 1, 0, 0, 0, 291, 235,
		1, 0, 0, 0, 291, 238, 1, 0, 0, 0, 291, 242, 1, 0, 0, 0, 291, 251, 1, 0,
		0, 0, 291, 257, 1, 0, 0, 0, 291, 263, 1, 0, 0, 0, 291, 266, 1, 0, 0, 0,
		291, 274, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 41, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 305, 5,
		33, 0, 0, 297, 302, 3, 40, 20, 0, 298, 299, 5, 39, 0, 0, 299, 301, 3, 40,
		20, 0, 300, 298, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0,
		302, 303, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305,
		297, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308,
		5, 34, 0, 0, 308, 43, 1, 0, 0, 0, 309, 318, 5, 31, 0, 0, 310, 315, 3, 46,
		23, 0, 311, 312, 5, 39, 0, 0, 312, 314, 3, 46, 23, 0, 313, 311, 1, 0, 0,
		0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316,
		319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 310, 1, 0, 0, 0, 318, 319,
		1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 5, 32, 0, 0, 321, 45, 1, 0,
		0, 0, 322, 325, 3, 72, 36, 0, 323, 325, 3, 80, 40, 0, 324, 322, 1, 0, 0,
		0, 324, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 5, 42, 0, 0, 327,
		328, 3, 40, 20, 0, 328, 47, 1, 0, 0, 0, 329, 332, 3, 50, 25, 0, 330, 332,
		3, 52, 26, 0, 331, 329, 1, 0, 0, 0, 331, 330, 1, 0, 0, 0, 332, 49, 1, 0,
		0, 0, 333, 334, 5, 11, 0, 0, 334, 335, 5, 40, 0, 0, 335, 336, 3, 78, 39,
		0, 336, 345, 5, 35, 0, 0, 337, 342, 3, 40, 20, 0, 338, 339, 5, 39, 0, 0,
		339, 341, 3, 40, 20, 0, 340, 338, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0, 342,
		340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344, 342,
		1, 0, 0, 0, 345, 337, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 1, 0,
		0, 0, 347, 348, 5, 36, 0, 0, 348, 51, 1, 0, 0, 0, 349, 350, 3, 78, 39,
		0, 350, 362, 5, 35, 0, 0, 351, 354, 5, 1, 0, 0, 352, 354, 3, 40, 20, 0,
		353, 351, 1, 0, 0, 0, 353, 352, 1, 0, 0, 0, 354, 359, 1, 0, 0, 0, 355,
		356, 5, 39, 0, 0, 356, 358, 3, 40, 20, 0, 357, 355, 1, 0, 0, 0, 358, 361,
		1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 363, 1, 0,
		0, 0, 361, 359, 1, 0, 0, 0, 362, 353, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0,
		363, 364, 1, 0, 0, 0, 364, 365, 5, 36, 0, 0, 365, 53, 1, 0, 0, 0, 366,
		367, 7, 1, 0, 0, 367, 55, 1, 0, 0, 0, 368, 369, 7, 2, 0, 0, 369, 57, 1,
		0, 0, 0, 370, 371, 5, 30, 0, 0, 371, 372, 3, 78, 39, 0, 372, 59, 1, 0,
		0, 0, 373, 379, 3, 62, 31, 0, 374, 379, 3, 64, 32, 0, 375, 379, 3, 66,
		33, 0, 376, 379, 3, 68, 34, 0, 377, 379, 3, 70, 35, 0, 378, 373, 1, 0,
		0, 0, 378, 374, 1, 0, 0, 0, 378, 375, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0,
		378, 377, 1, 0, 0, 0, 379, 61, 1, 0, 0, 0, 380, 381, 5, 6, 0, 0, 381, 63,
		1, 0, 0, 0, 382, 383, 5, 7, 0, 0, 383, 65, 1, 0, 0, 0, 384, 385, 7, 3,
		0, 0, 385, 67, 1, 0, 0, 0, 386, 389, 3, 74, 37, 0, 387, 389, 3, 76, 38,
		0, 388, 386, 1, 0, 0, 0, 388, 387, 1, 0, 0, 0, 389, 69, 1, 0, 0, 0, 390,
		391, 3, 72, 36, 0, 391, 71, 1, 0, 0, 0, 392, 393, 7, 4, 0, 0, 393, 73,
		1, 0, 0, 0, 394, 395, 7, 5, 0, 0, 395, 75, 1, 0, 0, 0, 396, 397, 5, 66,
		0, 0, 397, 77, 1, 0, 0, 0, 398, 399, 7, 6, 0, 0, 399, 79, 1, 0, 0, 0, 400,
		401, 3, 78, 39, 0, 401, 81, 1, 0, 0, 0, 402, 403, 5, 63, 0, 0, 403, 83,
		1, 0, 0, 0, 404, 405, 3, 78, 39, 0, 405, 85, 1, 0, 0, 0, 41, 91, 94, 97,
		100, 103, 106, 110, 119, 122, 125, 140, 153, 158, 171, 176, 179, 193, 198,
		201, 230, 244, 253, 271, 276, 285, 288, 291, 293, 302, 305, 315, 318, 324,
		331, 342, 345, 353, 359, 362, 378, 388,
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
	CosmosDBParserEOF                            = antlr.TokenEOF
	CosmosDBParserMULTIPLY_OPERATOR              = 1
	CosmosDBParserAS_SYMBOL                      = 2
	CosmosDBParserSELECT_SYMBOL                  = 3
	CosmosDBParserFROM_SYMBOL                    = 4
	CosmosDBParserDISTINCT_SYMBOL                = 5
	CosmosDBParserUNDEFINED_SYMBOL               = 6
	CosmosDBParserNULL_SYMBOL                    = 7
	CosmosDBParserFALSE_SYMBOL                   = 8
	CosmosDBParserTRUE_SYMBOL                    = 9
	CosmosDBParserNOT_SYMBOL                     = 10
	CosmosDBParserUDF_SYMBOL                     = 11
	CosmosDBParserWHERE_SYMBOL                   = 12
	CosmosDBParserAND_SYMBOL                     = 13
	CosmosDBParserOR_SYMBOL                      = 14
	CosmosDBParserIN_SYMBOL                      = 15
	CosmosDBParserBETWEEN_SYMBOL                 = 16
	CosmosDBParserTOP_SYMBOL                     = 17
	CosmosDBParserVALUE_SYMBOL                   = 18
	CosmosDBParserORDER_SYMBOL                   = 19
	CosmosDBParserBY_SYMBOL                      = 20
	CosmosDBParserGROUP_SYMBOL                   = 21
	CosmosDBParserOFFSET_SYMBOL                  = 22
	CosmosDBParserLIMIT_SYMBOL                   = 23
	CosmosDBParserASC_SYMBOL                     = 24
	CosmosDBParserDESC_SYMBOL                    = 25
	CosmosDBParserEXISTS_SYMBOL                  = 26
	CosmosDBParserLIKE_SYMBOL                    = 27
	CosmosDBParserHAVING_SYMBOL                  = 28
	CosmosDBParserJOIN_SYMBOL                    = 29
	CosmosDBParserAT_SYMBOL                      = 30
	CosmosDBParserLC_BRACKET_SYMBOL              = 31
	CosmosDBParserRC_BRACKET_SYMBOL              = 32
	CosmosDBParserLS_BRACKET_SYMBOL              = 33
	CosmosDBParserRS_BRACKET_SYMBOL              = 34
	CosmosDBParserLR_BRACKET_SYMBOL              = 35
	CosmosDBParserRR_BRACKET_SYMBOL              = 36
	CosmosDBParserSINGLE_QUOTE_SYMBOL            = 37
	CosmosDBParserDOUBLE_QUOTE_SYMBOL            = 38
	CosmosDBParserCOMMA_SYMBOL                   = 39
	CosmosDBParserDOT_SYMBOL                     = 40
	CosmosDBParserQUESTION_MARK_SYMBOL           = 41
	CosmosDBParserCOLON_SYMBOL                   = 42
	CosmosDBParserPLUS_SYMBOL                    = 43
	CosmosDBParserMINUS_SYMBOL                   = 44
	CosmosDBParserBIT_NOT_SYMBOL                 = 45
	CosmosDBParserDIVIDE_SYMBOL                  = 46
	CosmosDBParserMODULO_SYMBOL                  = 47
	CosmosDBParserBIT_AND_SYMBOL                 = 48
	CosmosDBParserBIT_OR_SYMBOL                  = 49
	CosmosDBParserDOUBLE_BAR_SYMBOL              = 50
	CosmosDBParserBIT_XOR_SYMBOL                 = 51
	CosmosDBParserEQUAL_SYMBOL                   = 52
	CosmosDBParserLESS_THAN_OPERATOR             = 53
	CosmosDBParserLESS_THAN_EQUAL_OPERATOR       = 54
	CosmosDBParserGREATER_THAN_OPERATOR          = 55
	CosmosDBParserGREATER_THAN_EQUAL_OPERATOR    = 56
	CosmosDBParserLEFT_SHIFT_OPERATOR            = 57
	CosmosDBParserRIGHT_SHIFT_OPERATOR           = 58
	CosmosDBParserZERO_FILL_RIGHT_SHIFT_OPERATOR = 59
	CosmosDBParserNOT_EQUAL_OPERATOR             = 60
	CosmosDBParserIDENTIFIER                     = 61
	CosmosDBParserWHITESPACE                     = 62
	CosmosDBParserDECIMAL                        = 63
	CosmosDBParserREAL                           = 64
	CosmosDBParserFLOAT                          = 65
	CosmosDBParserHEXADECIMAL                    = 66
	CosmosDBParserSINGLE_QUOTE_STRING_LITERAL    = 67
	CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL    = 68
)

// CosmosDBParser rules.
const (
	CosmosDBParserRULE_root                           = 0
	CosmosDBParserRULE_select                         = 1
	CosmosDBParserRULE_select_clause                  = 2
	CosmosDBParserRULE_top_clause                     = 3
	CosmosDBParserRULE_select_specification           = 4
	CosmosDBParserRULE_from_clause                    = 5
	CosmosDBParserRULE_where_clause                   = 6
	CosmosDBParserRULE_group_by_clause                = 7
	CosmosDBParserRULE_having_clause                  = 8
	CosmosDBParserRULE_order_by_clause                = 9
	CosmosDBParserRULE_sort_expression                = 10
	CosmosDBParserRULE_offset_limit_clause            = 11
	CosmosDBParserRULE_from_specification             = 12
	CosmosDBParserRULE_from_source                    = 13
	CosmosDBParserRULE_container_expression           = 14
	CosmosDBParserRULE_join_clause                    = 15
	CosmosDBParserRULE_container_name                 = 16
	CosmosDBParserRULE_object_property_list           = 17
	CosmosDBParserRULE_object_property                = 18
	CosmosDBParserRULE_property_alias                 = 19
	CosmosDBParserRULE_scalar_expression              = 20
	CosmosDBParserRULE_create_array_expression        = 21
	CosmosDBParserRULE_create_object_expression       = 22
	CosmosDBParserRULE_object_field_pair              = 23
	CosmosDBParserRULE_scalar_function_expression     = 24
	CosmosDBParserRULE_udf_scalar_function_expression = 25
	CosmosDBParserRULE_builtin_function_expression    = 26
	CosmosDBParserRULE_binary_operator                = 27
	CosmosDBParserRULE_unary_operator                 = 28
	CosmosDBParserRULE_parameter_name                 = 29
	CosmosDBParserRULE_constant                       = 30
	CosmosDBParserRULE_undefined_constant             = 31
	CosmosDBParserRULE_null_constant                  = 32
	CosmosDBParserRULE_boolean_constant               = 33
	CosmosDBParserRULE_number_constant                = 34
	CosmosDBParserRULE_string_constant                = 35
	CosmosDBParserRULE_string_literal                 = 36
	CosmosDBParserRULE_decimal_literal                = 37
	CosmosDBParserRULE_hexadecimal_literal            = 38
	CosmosDBParserRULE_identifier                     = 39
	CosmosDBParserRULE_property_name                  = 40
	CosmosDBParserRULE_array_index                    = 41
	CosmosDBParserRULE_input_alias                    = 42
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
		p.SetState(86)
		p.Select_()
	}
	{
		p.SetState(87)
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
	Group_by_clause() IGroup_by_clauseContext
	Having_clause() IHaving_clauseContext
	Order_by_clause() IOrder_by_clauseContext
	Offset_limit_clause() IOffset_limit_clauseContext

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

func (s *SelectContext) Group_by_clause() IGroup_by_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroup_by_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroup_by_clauseContext)
}

func (s *SelectContext) Having_clause() IHaving_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHaving_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHaving_clauseContext)
}

func (s *SelectContext) Order_by_clause() IOrder_by_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrder_by_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrder_by_clauseContext)
}

func (s *SelectContext) Offset_limit_clause() IOffset_limit_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffset_limit_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffset_limit_clauseContext)
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
		p.SetState(89)
		p.Select_clause()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserFROM_SYMBOL {
		{
			p.SetState(90)
			p.From_clause()
		}

	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserWHERE_SYMBOL {
		{
			p.SetState(93)
			p.Where_clause()
		}

	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserGROUP_SYMBOL {
		{
			p.SetState(96)
			p.Group_by_clause()
		}

	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserHAVING_SYMBOL {
		{
			p.SetState(99)
			p.Having_clause()
		}

	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserORDER_SYMBOL {
		{
			p.SetState(102)
			p.Order_by_clause()
		}

	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserOFFSET_SYMBOL {
		{
			p.SetState(105)
			p.Offset_limit_clause()
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
	Top_clause() ITop_clauseContext

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

func (s *Select_clauseContext) Top_clause() ITop_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITop_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITop_clauseContext)
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
		p.SetState(108)
		p.Match(CosmosDBParserSELECT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(109)
			p.Top_clause()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(112)
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

// ITop_clauseContext is an interface to support dynamic dispatch.
type ITop_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TOP_SYMBOL() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode

	// IsTop_clauseContext differentiates from other interfaces.
	IsTop_clauseContext()
}

type Top_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTop_clauseContext() *Top_clauseContext {
	var p = new(Top_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_top_clause
	return p
}

func InitEmptyTop_clauseContext(p *Top_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_top_clause
}

func (*Top_clauseContext) IsTop_clauseContext() {}

func NewTop_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_clauseContext {
	var p = new(Top_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_top_clause

	return p
}

func (s *Top_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Top_clauseContext) TOP_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserTOP_SYMBOL, 0)
}

func (s *Top_clauseContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDECIMAL, 0)
}

func (s *Top_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Top_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterTop_clause(s)
	}
}

func (s *Top_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitTop_clause(s)
	}
}

func (s *Top_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitTop_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Top_clause() (localctx ITop_clauseContext) {
	localctx = NewTop_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CosmosDBParserRULE_top_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(CosmosDBParserTOP_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
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

// ISelect_specificationContext is an interface to support dynamic dispatch.
type ISelect_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPLY_OPERATOR() antlr.TerminalNode
	Object_property_list() IObject_property_listContext
	DISTINCT_SYMBOL() antlr.TerminalNode
	VALUE_SYMBOL() antlr.TerminalNode

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

func (s *Select_specificationContext) VALUE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserVALUE_SYMBOL, 0)
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
	p.EnterRule(localctx, 8, CosmosDBParserRULE_select_specification)
	var _la int

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserMULTIPLY_OPERATOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(CosmosDBParserMULTIPLY_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CosmosDBParserDISTINCT_SYMBOL, CosmosDBParserUNDEFINED_SYMBOL, CosmosDBParserNULL_SYMBOL, CosmosDBParserFALSE_SYMBOL, CosmosDBParserTRUE_SYMBOL, CosmosDBParserNOT_SYMBOL, CosmosDBParserUDF_SYMBOL, CosmosDBParserIN_SYMBOL, CosmosDBParserBETWEEN_SYMBOL, CosmosDBParserTOP_SYMBOL, CosmosDBParserVALUE_SYMBOL, CosmosDBParserORDER_SYMBOL, CosmosDBParserBY_SYMBOL, CosmosDBParserGROUP_SYMBOL, CosmosDBParserOFFSET_SYMBOL, CosmosDBParserLIMIT_SYMBOL, CosmosDBParserASC_SYMBOL, CosmosDBParserDESC_SYMBOL, CosmosDBParserEXISTS_SYMBOL, CosmosDBParserLIKE_SYMBOL, CosmosDBParserHAVING_SYMBOL, CosmosDBParserJOIN_SYMBOL, CosmosDBParserAT_SYMBOL, CosmosDBParserLC_BRACKET_SYMBOL, CosmosDBParserLS_BRACKET_SYMBOL, CosmosDBParserLR_BRACKET_SYMBOL, CosmosDBParserPLUS_SYMBOL, CosmosDBParserMINUS_SYMBOL, CosmosDBParserBIT_NOT_SYMBOL, CosmosDBParserIDENTIFIER, CosmosDBParserDECIMAL, CosmosDBParserREAL, CosmosDBParserFLOAT, CosmosDBParserHEXADECIMAL, CosmosDBParserSINGLE_QUOTE_STRING_LITERAL, CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserDISTINCT_SYMBOL {
			{
				p.SetState(118)
				p.Match(CosmosDBParserDISTINCT_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(121)
				p.Match(CosmosDBParserVALUE_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(124)
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
	p.EnterRule(localctx, 10, CosmosDBParserRULE_from_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(CosmosDBParserFROM_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
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
	Scalar_expression() IScalar_expressionContext

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

func (s *Where_clauseContext) Scalar_expression() IScalar_expressionContext {
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
	p.EnterRule(localctx, 12, CosmosDBParserRULE_where_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(CosmosDBParserWHERE_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.scalar_expression(0)
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

// IGroup_by_clauseContext is an interface to support dynamic dispatch.
type IGroup_by_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUP_SYMBOL() antlr.TerminalNode
	BY_SYMBOL() antlr.TerminalNode
	AllScalar_expression() []IScalar_expressionContext
	Scalar_expression(i int) IScalar_expressionContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsGroup_by_clauseContext differentiates from other interfaces.
	IsGroup_by_clauseContext()
}

type Group_by_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_clauseContext() *Group_by_clauseContext {
	var p = new(Group_by_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_group_by_clause
	return p
}

func InitEmptyGroup_by_clauseContext(p *Group_by_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_group_by_clause
}

func (*Group_by_clauseContext) IsGroup_by_clauseContext() {}

func NewGroup_by_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_clauseContext {
	var p = new(Group_by_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_group_by_clause

	return p
}

func (s *Group_by_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_clauseContext) GROUP_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserGROUP_SYMBOL, 0)
}

func (s *Group_by_clauseContext) BY_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBY_SYMBOL, 0)
}

func (s *Group_by_clauseContext) AllScalar_expression() []IScalar_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expressionContext); ok {
			tst[i] = t.(IScalar_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Group_by_clauseContext) Scalar_expression(i int) IScalar_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
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

	return t.(IScalar_expressionContext)
}

func (s *Group_by_clauseContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Group_by_clauseContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Group_by_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_by_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterGroup_by_clause(s)
	}
}

func (s *Group_by_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitGroup_by_clause(s)
	}
}

func (s *Group_by_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitGroup_by_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Group_by_clause() (localctx IGroup_by_clauseContext) {
	localctx = NewGroup_by_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CosmosDBParserRULE_group_by_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(CosmosDBParserGROUP_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(CosmosDBParserBY_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.scalar_expression(0)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(136)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.scalar_expression(0)
		}

		p.SetState(142)
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

// IHaving_clauseContext is an interface to support dynamic dispatch.
type IHaving_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HAVING_SYMBOL() antlr.TerminalNode
	Scalar_expression() IScalar_expressionContext

	// IsHaving_clauseContext differentiates from other interfaces.
	IsHaving_clauseContext()
}

type Having_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHaving_clauseContext() *Having_clauseContext {
	var p = new(Having_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_having_clause
	return p
}

func InitEmptyHaving_clauseContext(p *Having_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_having_clause
}

func (*Having_clauseContext) IsHaving_clauseContext() {}

func NewHaving_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Having_clauseContext {
	var p = new(Having_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_having_clause

	return p
}

func (s *Having_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Having_clauseContext) HAVING_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserHAVING_SYMBOL, 0)
}

func (s *Having_clauseContext) Scalar_expression() IScalar_expressionContext {
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

func (s *Having_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Having_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Having_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterHaving_clause(s)
	}
}

func (s *Having_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitHaving_clause(s)
	}
}

func (s *Having_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitHaving_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Having_clause() (localctx IHaving_clauseContext) {
	localctx = NewHaving_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CosmosDBParserRULE_having_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(CosmosDBParserHAVING_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.scalar_expression(0)
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

// IOrder_by_clauseContext is an interface to support dynamic dispatch.
type IOrder_by_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER_SYMBOL() antlr.TerminalNode
	BY_SYMBOL() antlr.TerminalNode
	AllSort_expression() []ISort_expressionContext
	Sort_expression(i int) ISort_expressionContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

	// IsOrder_by_clauseContext differentiates from other interfaces.
	IsOrder_by_clauseContext()
}

type Order_by_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_clauseContext() *Order_by_clauseContext {
	var p = new(Order_by_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_order_by_clause
	return p
}

func InitEmptyOrder_by_clauseContext(p *Order_by_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_order_by_clause
}

func (*Order_by_clauseContext) IsOrder_by_clauseContext() {}

func NewOrder_by_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_clauseContext {
	var p = new(Order_by_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_order_by_clause

	return p
}

func (s *Order_by_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_clauseContext) ORDER_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserORDER_SYMBOL, 0)
}

func (s *Order_by_clauseContext) BY_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBY_SYMBOL, 0)
}

func (s *Order_by_clauseContext) AllSort_expression() []ISort_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISort_expressionContext); ok {
			len++
		}
	}

	tst := make([]ISort_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISort_expressionContext); ok {
			tst[i] = t.(ISort_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Order_by_clauseContext) Sort_expression(i int) ISort_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISort_expressionContext); ok {
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

	return t.(ISort_expressionContext)
}

func (s *Order_by_clauseContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Order_by_clauseContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
}

func (s *Order_by_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterOrder_by_clause(s)
	}
}

func (s *Order_by_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitOrder_by_clause(s)
	}
}

func (s *Order_by_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitOrder_by_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Order_by_clause() (localctx IOrder_by_clauseContext) {
	localctx = NewOrder_by_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CosmosDBParserRULE_order_by_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(CosmosDBParserORDER_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(CosmosDBParserBY_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Sort_expression()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(149)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Sort_expression()
		}

		p.SetState(155)
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

// ISort_expressionContext is an interface to support dynamic dispatch.
type ISort_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Scalar_expression() IScalar_expressionContext
	ASC_SYMBOL() antlr.TerminalNode
	DESC_SYMBOL() antlr.TerminalNode

	// IsSort_expressionContext differentiates from other interfaces.
	IsSort_expressionContext()
}

type Sort_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySort_expressionContext() *Sort_expressionContext {
	var p = new(Sort_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_sort_expression
	return p
}

func InitEmptySort_expressionContext(p *Sort_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_sort_expression
}

func (*Sort_expressionContext) IsSort_expressionContext() {}

func NewSort_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sort_expressionContext {
	var p = new(Sort_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_sort_expression

	return p
}

func (s *Sort_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Sort_expressionContext) Scalar_expression() IScalar_expressionContext {
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

func (s *Sort_expressionContext) ASC_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserASC_SYMBOL, 0)
}

func (s *Sort_expressionContext) DESC_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDESC_SYMBOL, 0)
}

func (s *Sort_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sort_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sort_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterSort_expression(s)
	}
}

func (s *Sort_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitSort_expression(s)
	}
}

func (s *Sort_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitSort_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Sort_expression() (localctx ISort_expressionContext) {
	localctx = NewSort_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CosmosDBParserRULE_sort_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.scalar_expression(0)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CosmosDBParserASC_SYMBOL || _la == CosmosDBParserDESC_SYMBOL {
		{
			p.SetState(157)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CosmosDBParserASC_SYMBOL || _la == CosmosDBParserDESC_SYMBOL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
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

// IOffset_limit_clauseContext is an interface to support dynamic dispatch.
type IOffset_limit_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET_SYMBOL() antlr.TerminalNode
	AllDECIMAL() []antlr.TerminalNode
	DECIMAL(i int) antlr.TerminalNode
	LIMIT_SYMBOL() antlr.TerminalNode

	// IsOffset_limit_clauseContext differentiates from other interfaces.
	IsOffset_limit_clauseContext()
}

type Offset_limit_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffset_limit_clauseContext() *Offset_limit_clauseContext {
	var p = new(Offset_limit_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_offset_limit_clause
	return p
}

func InitEmptyOffset_limit_clauseContext(p *Offset_limit_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_offset_limit_clause
}

func (*Offset_limit_clauseContext) IsOffset_limit_clauseContext() {}

func NewOffset_limit_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Offset_limit_clauseContext {
	var p = new(Offset_limit_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_offset_limit_clause

	return p
}

func (s *Offset_limit_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Offset_limit_clauseContext) OFFSET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserOFFSET_SYMBOL, 0)
}

func (s *Offset_limit_clauseContext) AllDECIMAL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserDECIMAL)
}

func (s *Offset_limit_clauseContext) DECIMAL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDECIMAL, i)
}

func (s *Offset_limit_clauseContext) LIMIT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLIMIT_SYMBOL, 0)
}

func (s *Offset_limit_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Offset_limit_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Offset_limit_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterOffset_limit_clause(s)
	}
}

func (s *Offset_limit_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitOffset_limit_clause(s)
	}
}

func (s *Offset_limit_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitOffset_limit_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Offset_limit_clause() (localctx IOffset_limit_clauseContext) {
	localctx = NewOffset_limit_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CosmosDBParserRULE_offset_limit_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(CosmosDBParserOFFSET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(CosmosDBParserDECIMAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(CosmosDBParserLIMIT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
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
	p.EnterRule(localctx, 24, CosmosDBParserRULE_from_specification)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
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
	AllJoin_clause() []IJoin_clauseContext
	Join_clause(i int) IJoin_clauseContext

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

func (s *From_sourceContext) AllJoin_clause() []IJoin_clauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoin_clauseContext); ok {
			len++
		}
	}

	tst := make([]IJoin_clauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoin_clauseContext); ok {
			tst[i] = t.(IJoin_clauseContext)
			i++
		}
	}

	return tst
}

func (s *From_sourceContext) Join_clause(i int) IJoin_clauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoin_clauseContext); ok {
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

	return t.(IJoin_clauseContext)
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
	p.EnterRule(localctx, 26, CosmosDBParserRULE_from_source)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Container_expression()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserJOIN_SYMBOL {
		{
			p.SetState(168)
			p.Join_clause()
		}

		p.SetState(173)
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

// IContainer_expressionContext is an interface to support dynamic dispatch.
type IContainer_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Container_name() IContainer_nameContext
	Identifier() IIdentifierContext
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

func (s *Container_expressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 28, CosmosDBParserRULE_container_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Container_name()
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserAS_SYMBOL {
			{
				p.SetState(175)
				p.Match(CosmosDBParserAS_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(178)
			p.Identifier()
		}

	} else if p.HasError() { // JIM
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

// IJoin_clauseContext is an interface to support dynamic dispatch.
type IJoin_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JOIN_SYMBOL() antlr.TerminalNode
	Identifier() IIdentifierContext
	IN_SYMBOL() antlr.TerminalNode
	Scalar_expression() IScalar_expressionContext

	// IsJoin_clauseContext differentiates from other interfaces.
	IsJoin_clauseContext()
}

type Join_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_clauseContext() *Join_clauseContext {
	var p = new(Join_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_join_clause
	return p
}

func InitEmptyJoin_clauseContext(p *Join_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_join_clause
}

func (*Join_clauseContext) IsJoin_clauseContext() {}

func NewJoin_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_clauseContext {
	var p = new(Join_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_join_clause

	return p
}

func (s *Join_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Join_clauseContext) JOIN_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserJOIN_SYMBOL, 0)
}

func (s *Join_clauseContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Join_clauseContext) IN_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIN_SYMBOL, 0)
}

func (s *Join_clauseContext) Scalar_expression() IScalar_expressionContext {
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

func (s *Join_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Join_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Join_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterJoin_clause(s)
	}
}

func (s *Join_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitJoin_clause(s)
	}
}

func (s *Join_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitJoin_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Join_clause() (localctx IJoin_clauseContext) {
	localctx = NewJoin_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CosmosDBParserRULE_join_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(CosmosDBParserJOIN_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Identifier()
	}
	{
		p.SetState(183)
		p.Match(CosmosDBParserIN_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.scalar_expression(0)
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
	Identifier() IIdentifierContext

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

func (s *Container_nameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 32, CosmosDBParserRULE_container_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Identifier()
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
	p.EnterRule(localctx, 34, CosmosDBParserRULE_object_property_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Object_property()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CosmosDBParserCOMMA_SYMBOL {
		{
			p.SetState(189)
			p.Match(CosmosDBParserCOMMA_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Object_property()
		}

		p.SetState(195)
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
	p.EnterRule(localctx, 36, CosmosDBParserRULE_object_property)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.scalar_expression(0)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CosmosDBParserAS_SYMBOL {
			{
				p.SetState(197)
				p.Match(CosmosDBParserAS_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(200)
			p.Property_alias()
		}

	} else if p.HasError() { // JIM
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

// IProperty_aliasContext is an interface to support dynamic dispatch.
type IProperty_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

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

func (s *Property_aliasContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 38, CosmosDBParserRULE_property_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Identifier()
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
	Constant() IConstantContext
	Input_alias() IInput_aliasContext
	Parameter_name() IParameter_nameContext
	Unary_operator() IUnary_operatorContext
	AllScalar_expression() []IScalar_expressionContext
	Scalar_expression(i int) IScalar_expressionContext
	NOT_SYMBOL() antlr.TerminalNode
	EXISTS_SYMBOL() antlr.TerminalNode
	LR_BRACKET_SYMBOL() antlr.TerminalNode
	Select_() ISelectContext
	RR_BRACKET_SYMBOL() antlr.TerminalNode
	Scalar_function_expression() IScalar_function_expressionContext
	Create_object_expression() ICreate_object_expressionContext
	Create_array_expression() ICreate_array_expressionContext
	AND_SYMBOL() antlr.TerminalNode
	OR_SYMBOL() antlr.TerminalNode
	Binary_operator() IBinary_operatorContext
	BETWEEN_SYMBOL() antlr.TerminalNode
	LIKE_SYMBOL() antlr.TerminalNode
	QUESTION_MARK_SYMBOL() antlr.TerminalNode
	COLON_SYMBOL() antlr.TerminalNode
	DOT_SYMBOL() antlr.TerminalNode
	Property_name() IProperty_nameContext
	LS_BRACKET_SYMBOL() antlr.TerminalNode
	RS_BRACKET_SYMBOL() antlr.TerminalNode
	DOUBLE_QUOTE_STRING_LITERAL() antlr.TerminalNode
	SINGLE_QUOTE_STRING_LITERAL() antlr.TerminalNode
	Array_index() IArray_indexContext
	IN_SYMBOL() antlr.TerminalNode
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

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

func (s *Scalar_expressionContext) Constant() IConstantContext {
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

func (s *Scalar_expressionContext) Parameter_name() IParameter_nameContext {
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

func (s *Scalar_expressionContext) AllScalar_expression() []IScalar_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expressionContext); ok {
			tst[i] = t.(IScalar_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Scalar_expressionContext) Scalar_expression(i int) IScalar_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
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

	return t.(IScalar_expressionContext)
}

func (s *Scalar_expressionContext) NOT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserNOT_SYMBOL, 0)
}

func (s *Scalar_expressionContext) EXISTS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserEXISTS_SYMBOL, 0)
}

func (s *Scalar_expressionContext) LR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLR_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expressionContext) Select_() ISelectContext {
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

func (s *Scalar_expressionContext) RR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRR_BRACKET_SYMBOL, 0)
}

func (s *Scalar_expressionContext) Scalar_function_expression() IScalar_function_expressionContext {
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

func (s *Scalar_expressionContext) Create_object_expression() ICreate_object_expressionContext {
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

func (s *Scalar_expressionContext) Create_array_expression() ICreate_array_expressionContext {
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

func (s *Scalar_expressionContext) AND_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserAND_SYMBOL, 0)
}

func (s *Scalar_expressionContext) OR_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserOR_SYMBOL, 0)
}

func (s *Scalar_expressionContext) Binary_operator() IBinary_operatorContext {
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

func (s *Scalar_expressionContext) BETWEEN_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBETWEEN_SYMBOL, 0)
}

func (s *Scalar_expressionContext) LIKE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLIKE_SYMBOL, 0)
}

func (s *Scalar_expressionContext) QUESTION_MARK_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserQUESTION_MARK_SYMBOL, 0)
}

func (s *Scalar_expressionContext) COLON_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOLON_SYMBOL, 0)
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

func (s *Scalar_expressionContext) SINGLE_QUOTE_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserSINGLE_QUOTE_STRING_LITERAL, 0)
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

func (s *Scalar_expressionContext) IN_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIN_SYMBOL, 0)
}

func (s *Scalar_expressionContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Scalar_expressionContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
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
	_startState := 40
	p.EnterRecursionRule(localctx, 40, CosmosDBParserRULE_scalar_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(206)
			p.Constant()
		}

	case 2:
		{
			p.SetState(207)
			p.Input_alias()
		}

	case 3:
		{
			p.SetState(208)
			p.Parameter_name()
		}

	case 4:
		{
			p.SetState(209)
			p.Unary_operator()
		}
		{
			p.SetState(210)
			p.scalar_expression(13)
		}

	case 5:
		{
			p.SetState(212)
			p.Match(CosmosDBParserNOT_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.scalar_expression(12)
		}

	case 6:
		{
			p.SetState(214)
			p.Match(CosmosDBParserEXISTS_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Select_()
		}
		{
			p.SetState(217)
			p.Match(CosmosDBParserRR_BRACKET_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(219)
			p.Scalar_function_expression()
		}

	case 8:
		{
			p.SetState(220)
			p.Create_object_expression()
		}

	case 9:
		{
			p.SetState(221)
			p.Create_array_expression()
		}

	case 10:
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
			p.scalar_expression(0)
		}
		{
			p.SetState(224)
			p.Match(CosmosDBParserRR_BRACKET_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		{
			p.SetState(226)
			p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Select_()
		}
		{
			p.SetState(228)
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
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(233)
					p.Match(CosmosDBParserAND_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(234)
					p.scalar_expression(18)
				}

			case 2:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(235)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(236)
					p.Match(CosmosDBParserOR_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(237)
					p.scalar_expression(17)
				}

			case 3:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(238)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(239)
					p.Binary_operator()
				}
				{
					p.SetState(240)
					p.scalar_expression(12)
				}

			case 4:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				p.SetState(244)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == CosmosDBParserNOT_SYMBOL {
					{
						p.SetState(243)
						p.Match(CosmosDBParserNOT_SYMBOL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(246)
					p.Match(CosmosDBParserBETWEEN_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(247)
					p.scalar_expression(0)
				}
				{
					p.SetState(248)
					p.Match(CosmosDBParserAND_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(249)
					p.scalar_expression(10)
				}

			case 5:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				p.SetState(253)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == CosmosDBParserNOT_SYMBOL {
					{
						p.SetState(252)
						p.Match(CosmosDBParserNOT_SYMBOL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(255)
					p.Match(CosmosDBParserLIKE_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(256)
					p.scalar_expression(9)
				}

			case 6:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(258)
					p.Match(CosmosDBParserQUESTION_MARK_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(259)
					p.scalar_expression(0)
				}
				{
					p.SetState(260)
					p.Match(CosmosDBParserCOLON_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(261)
					p.scalar_expression(7)
				}

			case 7:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(264)
					p.Match(CosmosDBParserDOT_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(265)
					p.Property_name()
				}

			case 8:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(267)
					p.Match(CosmosDBParserLS_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(271)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
					{
						p.SetState(268)
						p.Match(CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case CosmosDBParserSINGLE_QUOTE_STRING_LITERAL:
					{
						p.SetState(269)
						p.Match(CosmosDBParserSINGLE_QUOTE_STRING_LITERAL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case CosmosDBParserDECIMAL:
					{
						p.SetState(270)
						p.Array_index()
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}
				{
					p.SetState(273)
					p.Match(CosmosDBParserRS_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 9:
				localctx = NewScalar_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, CosmosDBParserRULE_scalar_expression)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				p.SetState(276)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == CosmosDBParserNOT_SYMBOL {
					{
						p.SetState(275)
						p.Match(CosmosDBParserNOT_SYMBOL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(278)
					p.Match(CosmosDBParserIN_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(279)
					p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(288)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64((_la-6)) & ^0x3f) == 0 && ((int64(1)<<(_la-6))&9115286608608755263) != 0 {
					{
						p.SetState(280)
						p.scalar_expression(0)
					}
					p.SetState(285)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					for _la == CosmosDBParserCOMMA_SYMBOL {
						{
							p.SetState(281)
							p.Match(CosmosDBParserCOMMA_SYMBOL)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(282)
							p.scalar_expression(0)
						}

						p.SetState(287)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(290)
					p.Match(CosmosDBParserRR_BRACKET_SYMBOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
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
	LS_BRACKET_SYMBOL() antlr.TerminalNode
	RS_BRACKET_SYMBOL() antlr.TerminalNode
	AllScalar_expression() []IScalar_expressionContext
	Scalar_expression(i int) IScalar_expressionContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

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

func (s *Create_array_expressionContext) LS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLS_BRACKET_SYMBOL, 0)
}

func (s *Create_array_expressionContext) RS_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRS_BRACKET_SYMBOL, 0)
}

func (s *Create_array_expressionContext) AllScalar_expression() []IScalar_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expressionContext); ok {
			tst[i] = t.(IScalar_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Create_array_expressionContext) Scalar_expression(i int) IScalar_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
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

	return t.(IScalar_expressionContext)
}

func (s *Create_array_expressionContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Create_array_expressionContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
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
	p.EnterRule(localctx, 42, CosmosDBParserRULE_create_array_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(CosmosDBParserLS_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-6)) & ^0x3f) == 0 && ((int64(1)<<(_la-6))&9115286608608755263) != 0 {
		{
			p.SetState(297)
			p.scalar_expression(0)
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CosmosDBParserCOMMA_SYMBOL {
			{
				p.SetState(298)
				p.Match(CosmosDBParserCOMMA_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(299)
				p.scalar_expression(0)
			}

			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(307)
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

// ICreate_object_expressionContext is an interface to support dynamic dispatch.
type ICreate_object_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC_BRACKET_SYMBOL() antlr.TerminalNode
	RC_BRACKET_SYMBOL() antlr.TerminalNode
	AllObject_field_pair() []IObject_field_pairContext
	Object_field_pair(i int) IObject_field_pairContext
	AllCOMMA_SYMBOL() []antlr.TerminalNode
	COMMA_SYMBOL(i int) antlr.TerminalNode

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

func (s *Create_object_expressionContext) LC_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLC_BRACKET_SYMBOL, 0)
}

func (s *Create_object_expressionContext) RC_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRC_BRACKET_SYMBOL, 0)
}

func (s *Create_object_expressionContext) AllObject_field_pair() []IObject_field_pairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_field_pairContext); ok {
			len++
		}
	}

	tst := make([]IObject_field_pairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_field_pairContext); ok {
			tst[i] = t.(IObject_field_pairContext)
			i++
		}
	}

	return tst
}

func (s *Create_object_expressionContext) Object_field_pair(i int) IObject_field_pairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_field_pairContext); ok {
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

	return t.(IObject_field_pairContext)
}

func (s *Create_object_expressionContext) AllCOMMA_SYMBOL() []antlr.TerminalNode {
	return s.GetTokens(CosmosDBParserCOMMA_SYMBOL)
}

func (s *Create_object_expressionContext) COMMA_SYMBOL(i int) antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOMMA_SYMBOL, i)
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
	p.EnterRule(localctx, 44, CosmosDBParserRULE_create_object_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(CosmosDBParserLC_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&13581167626321919) != 0 {
		{
			p.SetState(310)
			p.Object_field_pair()
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CosmosDBParserCOMMA_SYMBOL {
			{
				p.SetState(311)
				p.Match(CosmosDBParserCOMMA_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(312)
				p.Object_field_pair()
			}

			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(320)
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

// IObject_field_pairContext is an interface to support dynamic dispatch.
type IObject_field_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON_SYMBOL() antlr.TerminalNode
	Scalar_expression() IScalar_expressionContext
	String_literal() IString_literalContext
	Property_name() IProperty_nameContext

	// IsObject_field_pairContext differentiates from other interfaces.
	IsObject_field_pairContext()
}

type Object_field_pairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_field_pairContext() *Object_field_pairContext {
	var p = new(Object_field_pairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_field_pair
	return p
}

func InitEmptyObject_field_pairContext(p *Object_field_pairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_object_field_pair
}

func (*Object_field_pairContext) IsObject_field_pairContext() {}

func NewObject_field_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_field_pairContext {
	var p = new(Object_field_pairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_object_field_pair

	return p
}

func (s *Object_field_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_field_pairContext) COLON_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserCOLON_SYMBOL, 0)
}

func (s *Object_field_pairContext) Scalar_expression() IScalar_expressionContext {
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

func (s *Object_field_pairContext) String_literal() IString_literalContext {
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

func (s *Object_field_pairContext) Property_name() IProperty_nameContext {
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

func (s *Object_field_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_field_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_field_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterObject_field_pair(s)
	}
}

func (s *Object_field_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitObject_field_pair(s)
	}
}

func (s *Object_field_pairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitObject_field_pair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Object_field_pair() (localctx IObject_field_pairContext) {
	localctx = NewObject_field_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CosmosDBParserRULE_object_field_pair)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserSINGLE_QUOTE_STRING_LITERAL, CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
		{
			p.SetState(322)
			p.String_literal()
		}

	case CosmosDBParserIN_SYMBOL, CosmosDBParserBETWEEN_SYMBOL, CosmosDBParserTOP_SYMBOL, CosmosDBParserVALUE_SYMBOL, CosmosDBParserORDER_SYMBOL, CosmosDBParserBY_SYMBOL, CosmosDBParserGROUP_SYMBOL, CosmosDBParserOFFSET_SYMBOL, CosmosDBParserLIMIT_SYMBOL, CosmosDBParserASC_SYMBOL, CosmosDBParserDESC_SYMBOL, CosmosDBParserEXISTS_SYMBOL, CosmosDBParserLIKE_SYMBOL, CosmosDBParserHAVING_SYMBOL, CosmosDBParserJOIN_SYMBOL, CosmosDBParserIDENTIFIER:
		{
			p.SetState(323)
			p.Property_name()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(326)
		p.Match(CosmosDBParserCOLON_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.scalar_expression(0)
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
	p.EnterRule(localctx, 48, CosmosDBParserRULE_scalar_function_expression)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserUDF_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.Udf_scalar_function_expression()
		}

	case CosmosDBParserIN_SYMBOL, CosmosDBParserBETWEEN_SYMBOL, CosmosDBParserTOP_SYMBOL, CosmosDBParserVALUE_SYMBOL, CosmosDBParserORDER_SYMBOL, CosmosDBParserBY_SYMBOL, CosmosDBParserGROUP_SYMBOL, CosmosDBParserOFFSET_SYMBOL, CosmosDBParserLIMIT_SYMBOL, CosmosDBParserASC_SYMBOL, CosmosDBParserDESC_SYMBOL, CosmosDBParserEXISTS_SYMBOL, CosmosDBParserLIKE_SYMBOL, CosmosDBParserHAVING_SYMBOL, CosmosDBParserJOIN_SYMBOL, CosmosDBParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
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
	Identifier() IIdentifierContext
	LR_BRACKET_SYMBOL() antlr.TerminalNode
	RR_BRACKET_SYMBOL() antlr.TerminalNode
	AllScalar_expression() []IScalar_expressionContext
	Scalar_expression(i int) IScalar_expressionContext
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

func (s *Udf_scalar_function_expressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Udf_scalar_function_expressionContext) LR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLR_BRACKET_SYMBOL, 0)
}

func (s *Udf_scalar_function_expressionContext) RR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRR_BRACKET_SYMBOL, 0)
}

func (s *Udf_scalar_function_expressionContext) AllScalar_expression() []IScalar_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expressionContext); ok {
			tst[i] = t.(IScalar_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Udf_scalar_function_expressionContext) Scalar_expression(i int) IScalar_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
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

	return t.(IScalar_expressionContext)
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
	p.EnterRule(localctx, 50, CosmosDBParserRULE_udf_scalar_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(CosmosDBParserUDF_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(334)
		p.Match(CosmosDBParserDOT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)
		p.Identifier()
	}
	{
		p.SetState(336)
		p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-6)) & ^0x3f) == 0 && ((int64(1)<<(_la-6))&9115286608608755263) != 0 {
		{
			p.SetState(337)
			p.scalar_expression(0)
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CosmosDBParserCOMMA_SYMBOL {
			{
				p.SetState(338)
				p.Match(CosmosDBParserCOMMA_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(339)
				p.scalar_expression(0)
			}

			p.SetState(344)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(347)
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
	Identifier() IIdentifierContext
	LR_BRACKET_SYMBOL() antlr.TerminalNode
	RR_BRACKET_SYMBOL() antlr.TerminalNode
	MULTIPLY_OPERATOR() antlr.TerminalNode
	AllScalar_expression() []IScalar_expressionContext
	Scalar_expression(i int) IScalar_expressionContext
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

func (s *Builtin_function_expressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Builtin_function_expressionContext) LR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLR_BRACKET_SYMBOL, 0)
}

func (s *Builtin_function_expressionContext) RR_BRACKET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRR_BRACKET_SYMBOL, 0)
}

func (s *Builtin_function_expressionContext) MULTIPLY_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserMULTIPLY_OPERATOR, 0)
}

func (s *Builtin_function_expressionContext) AllScalar_expression() []IScalar_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalar_expressionContext); ok {
			len++
		}
	}

	tst := make([]IScalar_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalar_expressionContext); ok {
			tst[i] = t.(IScalar_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Builtin_function_expressionContext) Scalar_expression(i int) IScalar_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalar_expressionContext); ok {
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

	return t.(IScalar_expressionContext)
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
	p.EnterRule(localctx, 52, CosmosDBParserRULE_builtin_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Identifier()
	}
	{
		p.SetState(350)
		p.Match(CosmosDBParserLR_BRACKET_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6917467407745314878) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&31) != 0) {
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CosmosDBParserMULTIPLY_OPERATOR:
			{
				p.SetState(351)
				p.Match(CosmosDBParserMULTIPLY_OPERATOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case CosmosDBParserUNDEFINED_SYMBOL, CosmosDBParserNULL_SYMBOL, CosmosDBParserFALSE_SYMBOL, CosmosDBParserTRUE_SYMBOL, CosmosDBParserNOT_SYMBOL, CosmosDBParserUDF_SYMBOL, CosmosDBParserIN_SYMBOL, CosmosDBParserBETWEEN_SYMBOL, CosmosDBParserTOP_SYMBOL, CosmosDBParserVALUE_SYMBOL, CosmosDBParserORDER_SYMBOL, CosmosDBParserBY_SYMBOL, CosmosDBParserGROUP_SYMBOL, CosmosDBParserOFFSET_SYMBOL, CosmosDBParserLIMIT_SYMBOL, CosmosDBParserASC_SYMBOL, CosmosDBParserDESC_SYMBOL, CosmosDBParserEXISTS_SYMBOL, CosmosDBParserLIKE_SYMBOL, CosmosDBParserHAVING_SYMBOL, CosmosDBParserJOIN_SYMBOL, CosmosDBParserAT_SYMBOL, CosmosDBParserLC_BRACKET_SYMBOL, CosmosDBParserLS_BRACKET_SYMBOL, CosmosDBParserLR_BRACKET_SYMBOL, CosmosDBParserPLUS_SYMBOL, CosmosDBParserMINUS_SYMBOL, CosmosDBParserBIT_NOT_SYMBOL, CosmosDBParserIDENTIFIER, CosmosDBParserDECIMAL, CosmosDBParserREAL, CosmosDBParserFLOAT, CosmosDBParserHEXADECIMAL, CosmosDBParserSINGLE_QUOTE_STRING_LITERAL, CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
			{
				p.SetState(352)
				p.scalar_expression(0)
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CosmosDBParserCOMMA_SYMBOL {
			{
				p.SetState(355)
				p.Match(CosmosDBParserCOMMA_SYMBOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(356)
				p.scalar_expression(0)
			}

			p.SetState(361)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(364)
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
	NOT_EQUAL_OPERATOR() antlr.TerminalNode
	LESS_THAN_OPERATOR() antlr.TerminalNode
	LESS_THAN_EQUAL_OPERATOR() antlr.TerminalNode
	GREATER_THAN_OPERATOR() antlr.TerminalNode
	GREATER_THAN_EQUAL_OPERATOR() antlr.TerminalNode
	LEFT_SHIFT_OPERATOR() antlr.TerminalNode
	RIGHT_SHIFT_OPERATOR() antlr.TerminalNode
	ZERO_FILL_RIGHT_SHIFT_OPERATOR() antlr.TerminalNode

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

func (s *Binary_operatorContext) NOT_EQUAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserNOT_EQUAL_OPERATOR, 0)
}

func (s *Binary_operatorContext) LESS_THAN_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLESS_THAN_OPERATOR, 0)
}

func (s *Binary_operatorContext) LESS_THAN_EQUAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLESS_THAN_EQUAL_OPERATOR, 0)
}

func (s *Binary_operatorContext) GREATER_THAN_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserGREATER_THAN_OPERATOR, 0)
}

func (s *Binary_operatorContext) GREATER_THAN_EQUAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserGREATER_THAN_EQUAL_OPERATOR, 0)
}

func (s *Binary_operatorContext) LEFT_SHIFT_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLEFT_SHIFT_OPERATOR, 0)
}

func (s *Binary_operatorContext) RIGHT_SHIFT_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserRIGHT_SHIFT_OPERATOR, 0)
}

func (s *Binary_operatorContext) ZERO_FILL_RIGHT_SHIFT_OPERATOR() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserZERO_FILL_RIGHT_SHIFT_OPERATOR, 0)
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
	p.EnterRule(localctx, 54, CosmosDBParserRULE_binary_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305799028748582914) != 0) {
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
	p.EnterRule(localctx, 56, CosmosDBParserRULE_unary_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&61572651155456) != 0) {
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
	Identifier() IIdentifierContext

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

func (s *Parameter_nameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 58, CosmosDBParserRULE_parameter_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(CosmosDBParserAT_SYMBOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Identifier()
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
	p.EnterRule(localctx, 60, CosmosDBParserRULE_constant)
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserUNDEFINED_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(373)
			p.Undefined_constant()
		}

	case CosmosDBParserNULL_SYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(374)
			p.Null_constant()
		}

	case CosmosDBParserFALSE_SYMBOL, CosmosDBParserTRUE_SYMBOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(375)
			p.Boolean_constant()
		}

	case CosmosDBParserDECIMAL, CosmosDBParserREAL, CosmosDBParserFLOAT, CosmosDBParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(376)
			p.Number_constant()
		}

	case CosmosDBParserSINGLE_QUOTE_STRING_LITERAL, CosmosDBParserDOUBLE_QUOTE_STRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(377)
			p.String_constant()
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
	p.EnterRule(localctx, 62, CosmosDBParserRULE_undefined_constant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
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
	p.EnterRule(localctx, 64, CosmosDBParserRULE_null_constant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
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
	p.EnterRule(localctx, 66, CosmosDBParserRULE_boolean_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
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
	p.EnterRule(localctx, 68, CosmosDBParserRULE_number_constant)
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CosmosDBParserDECIMAL, CosmosDBParserREAL, CosmosDBParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)
			p.Decimal_literal()
		}

	case CosmosDBParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)
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
	p.EnterRule(localctx, 70, CosmosDBParserRULE_string_constant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
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
	p.EnterRule(localctx, 72, CosmosDBParserRULE_string_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
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
	p.EnterRule(localctx, 74, CosmosDBParserRULE_decimal_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-63)) & ^0x3f) == 0 && ((int64(1)<<(_la-63))&7) != 0) {
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
	p.EnterRule(localctx, 76, CosmosDBParserRULE_hexadecimal_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	IN_SYMBOL() antlr.TerminalNode
	BETWEEN_SYMBOL() antlr.TerminalNode
	TOP_SYMBOL() antlr.TerminalNode
	VALUE_SYMBOL() antlr.TerminalNode
	ORDER_SYMBOL() antlr.TerminalNode
	BY_SYMBOL() antlr.TerminalNode
	GROUP_SYMBOL() antlr.TerminalNode
	OFFSET_SYMBOL() antlr.TerminalNode
	LIMIT_SYMBOL() antlr.TerminalNode
	ASC_SYMBOL() antlr.TerminalNode
	DESC_SYMBOL() antlr.TerminalNode
	EXISTS_SYMBOL() antlr.TerminalNode
	LIKE_SYMBOL() antlr.TerminalNode
	HAVING_SYMBOL() antlr.TerminalNode
	JOIN_SYMBOL() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CosmosDBParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CosmosDBParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIDENTIFIER, 0)
}

func (s *IdentifierContext) IN_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserIN_SYMBOL, 0)
}

func (s *IdentifierContext) BETWEEN_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBETWEEN_SYMBOL, 0)
}

func (s *IdentifierContext) TOP_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserTOP_SYMBOL, 0)
}

func (s *IdentifierContext) VALUE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserVALUE_SYMBOL, 0)
}

func (s *IdentifierContext) ORDER_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserORDER_SYMBOL, 0)
}

func (s *IdentifierContext) BY_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserBY_SYMBOL, 0)
}

func (s *IdentifierContext) GROUP_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserGROUP_SYMBOL, 0)
}

func (s *IdentifierContext) OFFSET_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserOFFSET_SYMBOL, 0)
}

func (s *IdentifierContext) LIMIT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLIMIT_SYMBOL, 0)
}

func (s *IdentifierContext) ASC_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserASC_SYMBOL, 0)
}

func (s *IdentifierContext) DESC_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserDESC_SYMBOL, 0)
}

func (s *IdentifierContext) EXISTS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserEXISTS_SYMBOL, 0)
}

func (s *IdentifierContext) LIKE_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserLIKE_SYMBOL, 0)
}

func (s *IdentifierContext) HAVING_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserHAVING_SYMBOL, 0)
}

func (s *IdentifierContext) JOIN_SYMBOL() antlr.TerminalNode {
	return s.GetToken(CosmosDBParserJOIN_SYMBOL, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CosmosDBParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CosmosDBParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CosmosDBParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CosmosDBParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305843010287403008) != 0) {
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

// IProperty_nameContext is an interface to support dynamic dispatch.
type IProperty_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

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

func (s *Property_nameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 80, CosmosDBParserRULE_property_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Identifier()
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
	p.EnterRule(localctx, 82, CosmosDBParserRULE_array_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
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
	Identifier() IIdentifierContext

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

func (s *Input_aliasContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.EnterRule(localctx, 84, CosmosDBParserRULE_input_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Identifier()
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
	case 20:
		var t *Scalar_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Scalar_expressionContext)
		}
		return p.Scalar_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CosmosDBParser) Scalar_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
