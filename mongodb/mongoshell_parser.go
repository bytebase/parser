// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
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

type MongoShellParser struct {
	*antlr.BaseParser
}

var MongoShellParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mongoshellparserParserInit() {
	staticData := &MongoShellParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'show'", "'dbs'", "'databases'", "'collections'", "'db'", "'new'",
		"'true'", "'false'", "'null'", "'getCollection'", "'getCollectionNames'",
		"'getCollectionInfos'", "'ObjectId'", "'ISODate'", "'Date'", "'UUID'",
		"'Long'", "'NumberLong'", "'Int32'", "'NumberInt'", "'Double'", "'Decimal128'",
		"'NumberDecimal'", "'Timestamp'", "'RegExp'", "'find'", "'findOne'",
		"'sort'", "'limit'", "'skip'", "'projection'", "'project'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "':'", "','", "'.'", "';'", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "SHOW", "DBS", "DATABASES", "COLLECTIONS", "DB", "NEW", "TRUE",
		"FALSE", "NULL", "GET_COLLECTION", "GET_COLLECTION_NAMES", "GET_COLLECTION_INFOS",
		"OBJECT_ID", "ISO_DATE", "DATE", "UUID", "LONG", "NUMBER_LONG", "INT32",
		"NUMBER_INT", "DOUBLE", "DECIMAL128", "NUMBER_DECIMAL", "TIMESTAMP",
		"REG_EXP", "FIND", "FIND_ONE", "SORT", "LIMIT", "SKIP_", "PROJECTION",
		"PROJECT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET",
		"COLON", "COMMA", "DOT", "SEMI", "DOLLAR", "LINE_COMMENT", "BLOCK_COMMENT",
		"REGEX_LITERAL", "NUMBER", "DOUBLE_QUOTED_STRING", "SINGLE_QUOTED_STRING",
		"IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "shellCommand", "dbStatement", "collectionAccess",
		"methodChain", "methodCall", "findMethod", "findOneMethod", "sortMethod",
		"limitMethod", "skipMethod", "projectionMethod", "genericMethod", "arguments",
		"argument", "document", "pair", "key", "value", "array", "helperFunction",
		"objectIdHelper", "isoDateHelper", "dateHelper", "uuidHelper", "longHelper",
		"int32Helper", "doubleHelper", "decimal128Helper", "timestampHelper",
		"regExpConstructor", "literal", "stringLiteral", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 51, 400, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 5, 0, 72, 8, 0, 10,
		0, 12, 0, 75, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 81, 8, 1, 1, 1, 1, 1,
		3, 1, 85, 8, 1, 3, 1, 87, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 93, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 101, 8, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 108, 8, 3, 1, 3, 1, 3, 3, 3, 112, 8, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 118, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5,
		5, 138, 8, 5, 10, 5, 12, 5, 141, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 150, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 155, 8, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 3, 8, 162, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 189,
		8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 196, 8, 14, 10, 14, 12,
		14, 199, 9, 14, 1, 14, 3, 14, 202, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 5, 16, 210, 8, 16, 10, 16, 12, 16, 213, 9, 16, 1, 16, 3, 16,
		216, 8, 16, 3, 16, 218, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 3, 18, 228, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 236, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 242, 8, 20, 10,
		20, 12, 20, 245, 9, 20, 1, 20, 3, 20, 248, 8, 20, 3, 20, 250, 8, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 263, 8, 21, 1, 22, 1, 22, 1, 22, 3, 22, 268, 8, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 3, 22, 274, 8, 22, 1, 23, 1, 23, 1, 23, 3, 23, 279, 8, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 285, 8, 23, 1, 24, 1, 24, 1, 24, 1,
		24, 3, 24, 291, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 297, 8, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 308,
		8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 338, 8,
		30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 345, 8, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 3, 31, 352, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 3, 32, 359, 8, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 398,
		8, 34, 1, 34, 0, 0, 35, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 0, 6, 1, 0, 2, 3, 1, 0, 31, 32, 1, 0, 17, 18, 1, 0, 19,
		20, 1, 0, 22, 23, 1, 0, 48, 49, 456, 0, 73, 1, 0, 0, 0, 2, 86, 1, 0, 0,
		0, 4, 92, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8, 131, 1, 0, 0, 0, 10, 133,
		1, 0, 0, 0, 12, 149, 1, 0, 0, 0, 14, 151, 1, 0, 0, 0, 16, 158, 1, 0, 0,
		0, 18, 165, 1, 0, 0, 0, 20, 170, 1, 0, 0, 0, 22, 175, 1, 0, 0, 0, 24, 180,
		1, 0, 0, 0, 26, 185, 1, 0, 0, 0, 28, 192, 1, 0, 0, 0, 30, 203, 1, 0, 0,
		0, 32, 205, 1, 0, 0, 0, 34, 221, 1, 0, 0, 0, 36, 227, 1, 0, 0, 0, 38, 235,
		1, 0, 0, 0, 40, 237, 1, 0, 0, 0, 42, 262, 1, 0, 0, 0, 44, 273, 1, 0, 0,
		0, 46, 284, 1, 0, 0, 0, 48, 296, 1, 0, 0, 0, 50, 298, 1, 0, 0, 0, 52, 303,
		1, 0, 0, 0, 54, 311, 1, 0, 0, 0, 56, 316, 1, 0, 0, 0, 58, 321, 1, 0, 0,
		0, 60, 337, 1, 0, 0, 0, 62, 351, 1, 0, 0, 0, 64, 358, 1, 0, 0, 0, 66, 360,
		1, 0, 0, 0, 68, 397, 1, 0, 0, 0, 70, 72, 3, 2, 1, 0, 71, 70, 1, 0, 0, 0,
		72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1,
		0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 0, 0, 1, 77, 1, 1, 0, 0, 0, 78,
		80, 3, 4, 2, 0, 79, 81, 5, 42, 0, 0, 80, 79, 1, 0, 0, 0, 80, 81, 1, 0,
		0, 0, 81, 87, 1, 0, 0, 0, 82, 84, 3, 6, 3, 0, 83, 85, 5, 42, 0, 0, 84,
		83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 78, 1, 0, 0,
		0, 86, 82, 1, 0, 0, 0, 87, 3, 1, 0, 0, 0, 88, 89, 5, 1, 0, 0, 89, 93, 7,
		0, 0, 0, 90, 91, 5, 1, 0, 0, 91, 93, 5, 4, 0, 0, 92, 88, 1, 0, 0, 0, 92,
		90, 1, 0, 0, 0, 93, 5, 1, 0, 0, 0, 94, 95, 5, 5, 0, 0, 95, 96, 5, 41, 0,
		0, 96, 97, 5, 11, 0, 0, 97, 98, 5, 33, 0, 0, 98, 100, 5, 34, 0, 0, 99,
		101, 3, 10, 5, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 118,
		1, 0, 0, 0, 102, 103, 5, 5, 0, 0, 103, 104, 5, 41, 0, 0, 104, 105, 5, 12,
		0, 0, 105, 107, 5, 33, 0, 0, 106, 108, 3, 28, 14, 0, 107, 106, 1, 0, 0,
		0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 111, 5, 34, 0, 0, 110,
		112, 3, 10, 5, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 118,
		1, 0, 0, 0, 113, 114, 5, 5, 0, 0, 114, 115, 3, 8, 4, 0, 115, 116, 3, 10,
		5, 0, 116, 118, 1, 0, 0, 0, 117, 94, 1, 0, 0, 0, 117, 102, 1, 0, 0, 0,
		117, 113, 1, 0, 0, 0, 118, 7, 1, 0, 0, 0, 119, 120, 5, 41, 0, 0, 120, 132,
		3, 68, 34, 0, 121, 122, 5, 37, 0, 0, 122, 123, 3, 66, 33, 0, 123, 124,
		5, 38, 0, 0, 124, 132, 1, 0, 0, 0, 125, 126, 5, 41, 0, 0, 126, 127, 5,
		10, 0, 0, 127, 128, 5, 33, 0, 0, 128, 129, 3, 66, 33, 0, 129, 130, 5, 34,
		0, 0, 130, 132, 1, 0, 0, 0, 131, 119, 1, 0, 0, 0, 131, 121, 1, 0, 0, 0,
		131, 125, 1, 0, 0, 0, 132, 9, 1, 0, 0, 0, 133, 134, 5, 41, 0, 0, 134, 139,
		3, 12, 6, 0, 135, 136, 5, 41, 0, 0, 136, 138, 3, 12, 6, 0, 137, 135, 1,
		0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0,
		0, 140, 11, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 150, 3, 14, 7, 0, 143,
		150, 3, 16, 8, 0, 144, 150, 3, 18, 9, 0, 145, 150, 3, 20, 10, 0, 146, 150,
		3, 22, 11, 0, 147, 150, 3, 24, 12, 0, 148, 150, 3, 26, 13, 0, 149, 142,
		1, 0, 0, 0, 149, 143, 1, 0, 0, 0, 149, 144, 1, 0, 0, 0, 149, 145, 1, 0,
		0, 0, 149, 146, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 148, 1, 0, 0, 0,
		150, 13, 1, 0, 0, 0, 151, 152, 5, 26, 0, 0, 152, 154, 5, 33, 0, 0, 153,
		155, 3, 30, 15, 0, 154, 153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 157, 5, 34, 0, 0, 157, 15, 1, 0, 0, 0, 158, 159, 5, 27,
		0, 0, 159, 161, 5, 33, 0, 0, 160, 162, 3, 30, 15, 0, 161, 160, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 5, 34, 0, 0, 164,
		17, 1, 0, 0, 0, 165, 166, 5, 28, 0, 0, 166, 167, 5, 33, 0, 0, 167, 168,
		3, 32, 16, 0, 168, 169, 5, 34, 0, 0, 169, 19, 1, 0, 0, 0, 170, 171, 5,
		29, 0, 0, 171, 172, 5, 33, 0, 0, 172, 173, 5, 47, 0, 0, 173, 174, 5, 34,
		0, 0, 174, 21, 1, 0, 0, 0, 175, 176, 5, 30, 0, 0, 176, 177, 5, 33, 0, 0,
		177, 178, 5, 47, 0, 0, 178, 179, 5, 34, 0, 0, 179, 23, 1, 0, 0, 0, 180,
		181, 7, 1, 0, 0, 181, 182, 5, 33, 0, 0, 182, 183, 3, 32, 16, 0, 183, 184,
		5, 34, 0, 0, 184, 25, 1, 0, 0, 0, 185, 186, 3, 68, 34, 0, 186, 188, 5,
		33, 0, 0, 187, 189, 3, 28, 14, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1, 0,
		0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 5, 34, 0, 0, 191, 27, 1, 0, 0, 0,
		192, 197, 3, 30, 15, 0, 193, 194, 5, 40, 0, 0, 194, 196, 3, 30, 15, 0,
		195, 193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 202,
		5, 40, 0, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 29, 1, 0,
		0, 0, 203, 204, 3, 38, 19, 0, 204, 31, 1, 0, 0, 0, 205, 217, 5, 35, 0,
		0, 206, 211, 3, 34, 17, 0, 207, 208, 5, 40, 0, 0, 208, 210, 3, 34, 17,
		0, 209, 207, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211,
		212, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214, 216,
		5, 40, 0, 0, 215, 214, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218, 1, 0,
		0, 0, 217, 206, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0,
		219, 220, 5, 36, 0, 0, 220, 33, 1, 0, 0, 0, 221, 222, 3, 36, 18, 0, 222,
		223, 5, 39, 0, 0, 223, 224, 3, 38, 19, 0, 224, 35, 1, 0, 0, 0, 225, 228,
		3, 68, 34, 0, 226, 228, 3, 66, 33, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1,
		0, 0, 0, 228, 37, 1, 0, 0, 0, 229, 236, 3, 32, 16, 0, 230, 236, 3, 40,
		20, 0, 231, 236, 3, 42, 21, 0, 232, 236, 5, 46, 0, 0, 233, 236, 3, 62,
		31, 0, 234, 236, 3, 64, 32, 0, 235, 229, 1, 0, 0, 0, 235, 230, 1, 0, 0,
		0, 235, 231, 1, 0, 0, 0, 235, 232, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235,
		234, 1, 0, 0, 0, 236, 39, 1, 0, 0, 0, 237, 249, 5, 37, 0, 0, 238, 243,
		3, 38, 19, 0, 239, 240, 5, 40, 0, 0, 240, 242, 3, 38, 19, 0, 241, 239,
		1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0,
		0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 248, 5, 40, 0, 0,
		247, 246, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249,
		238, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252,
		5, 38, 0, 0, 252, 41, 1, 0, 0, 0, 253, 263, 3, 44, 22, 0, 254, 263, 3,
		46, 23, 0, 255, 263, 3, 48, 24, 0, 256, 263, 3, 50, 25, 0, 257, 263, 3,
		52, 26, 0, 258, 263, 3, 54, 27, 0, 259, 263, 3, 56, 28, 0, 260, 263, 3,
		58, 29, 0, 261, 263, 3, 60, 30, 0, 262, 253, 1, 0, 0, 0, 262, 254, 1, 0,
		0, 0, 262, 255, 1, 0, 0, 0, 262, 256, 1, 0, 0, 0, 262, 257, 1, 0, 0, 0,
		262, 258, 1, 0, 0, 0, 262, 259, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262,
		261, 1, 0, 0, 0, 263, 43, 1, 0, 0, 0, 264, 265, 5, 13, 0, 0, 265, 267,
		5, 33, 0, 0, 266, 268, 3, 66, 33, 0, 267, 266, 1, 0, 0, 0, 267, 268, 1,
		0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 274, 5, 34, 0, 0, 270, 271, 5, 6, 0,
		0, 271, 272, 5, 13, 0, 0, 272, 274, 6, 22, -1, 0, 273, 264, 1, 0, 0, 0,
		273, 270, 1, 0, 0, 0, 274, 45, 1, 0, 0, 0, 275, 276, 5, 14, 0, 0, 276,
		278, 5, 33, 0, 0, 277, 279, 3, 66, 33, 0, 278, 277, 1, 0, 0, 0, 278, 279,
		1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 285, 5, 34, 0, 0, 281, 282, 5, 6,
		0, 0, 282, 283, 5, 14, 0, 0, 283, 285, 6, 23, -1, 0, 284, 275, 1, 0, 0,
		0, 284, 281, 1, 0, 0, 0, 285, 47, 1, 0, 0, 0, 286, 287, 5, 15, 0, 0, 287,
		290, 5, 33, 0, 0, 288, 291, 3, 66, 33, 0, 289, 291, 5, 47, 0, 0, 290, 288,
		1, 0, 0, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 297, 5, 34, 0, 0, 293, 294, 5, 6, 0, 0, 294, 295, 5, 15, 0,
		0, 295, 297, 6, 24, -1, 0, 296, 286, 1, 0, 0, 0, 296, 293, 1, 0, 0, 0,
		297, 49, 1, 0, 0, 0, 298, 299, 5, 16, 0, 0, 299, 300, 5, 33, 0, 0, 300,
		301, 3, 66, 33, 0, 301, 302, 5, 34, 0, 0, 302, 51, 1, 0, 0, 0, 303, 304,
		7, 2, 0, 0, 304, 307, 5, 33, 0, 0, 305, 308, 5, 47, 0, 0, 306, 308, 3,
		66, 33, 0, 307, 305, 1, 0, 0, 0, 307, 306, 1, 0, 0, 0, 308, 309, 1, 0,
		0, 0, 309, 310, 5, 34, 0, 0, 310, 53, 1, 0, 0, 0, 311, 312, 7, 3, 0, 0,
		312, 313, 5, 33, 0, 0, 313, 314, 5, 47, 0, 0, 314, 315, 5, 34, 0, 0, 315,
		55, 1, 0, 0, 0, 316, 317, 5, 21, 0, 0, 317, 318, 5, 33, 0, 0, 318, 319,
		5, 47, 0, 0, 319, 320, 5, 34, 0, 0, 320, 57, 1, 0, 0, 0, 321, 322, 7, 4,
		0, 0, 322, 323, 5, 33, 0, 0, 323, 324, 3, 66, 33, 0, 324, 325, 5, 34, 0,
		0, 325, 59, 1, 0, 0, 0, 326, 327, 5, 24, 0, 0, 327, 328, 5, 33, 0, 0, 328,
		329, 3, 32, 16, 0, 329, 330, 5, 34, 0, 0, 330, 338, 1, 0, 0, 0, 331, 332,
		5, 24, 0, 0, 332, 333, 5, 33, 0, 0, 333, 334, 5, 47, 0, 0, 334, 335, 5,
		40, 0, 0, 335, 336, 5, 47, 0, 0, 336, 338, 5, 34, 0, 0, 337, 326, 1, 0,
		0, 0, 337, 331, 1, 0, 0, 0, 338, 61, 1, 0, 0, 0, 339, 340, 5, 25, 0, 0,
		340, 341, 5, 33, 0, 0, 341, 344, 3, 66, 33, 0, 342, 343, 5, 40, 0, 0, 343,
		345, 3, 66, 33, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 346,
		1, 0, 0, 0, 346, 347, 5, 34, 0, 0, 347, 352, 1, 0, 0, 0, 348, 349, 5, 6,
		0, 0, 349, 350, 5, 25, 0, 0, 350, 352, 6, 31, -1, 0, 351, 339, 1, 0, 0,
		0, 351, 348, 1, 0, 0, 0, 352, 63, 1, 0, 0, 0, 353, 359, 3, 66, 33, 0, 354,
		359, 5, 47, 0, 0, 355, 359, 5, 7, 0, 0, 356, 359, 5, 8, 0, 0, 357, 359,
		5, 9, 0, 0, 358, 353, 1, 0, 0, 0, 358, 354, 1, 0, 0, 0, 358, 355, 1, 0,
		0, 0, 358, 356, 1, 0, 0, 0, 358, 357, 1, 0, 0, 0, 359, 65, 1, 0, 0, 0,
		360, 361, 7, 5, 0, 0, 361, 67, 1, 0, 0, 0, 362, 398, 5, 50, 0, 0, 363,
		364, 5, 43, 0, 0, 364, 398, 5, 50, 0, 0, 365, 398, 5, 1, 0, 0, 366, 398,
		5, 2, 0, 0, 367, 398, 5, 3, 0, 0, 368, 398, 5, 4, 0, 0, 369, 398, 5, 5,
		0, 0, 370, 398, 5, 6, 0, 0, 371, 398, 5, 7, 0, 0, 372, 398, 5, 8, 0, 0,
		373, 398, 5, 9, 0, 0, 374, 398, 5, 26, 0, 0, 375, 398, 5, 27, 0, 0, 376,
		398, 5, 28, 0, 0, 377, 398, 5, 29, 0, 0, 378, 398, 5, 30, 0, 0, 379, 398,
		5, 31, 0, 0, 380, 398, 5, 32, 0, 0, 381, 398, 5, 10, 0, 0, 382, 398, 5,
		11, 0, 0, 383, 398, 5, 12, 0, 0, 384, 398, 5, 13, 0, 0, 385, 398, 5, 14,
		0, 0, 386, 398, 5, 15, 0, 0, 387, 398, 5, 16, 0, 0, 388, 398, 5, 17, 0,
		0, 389, 398, 5, 18, 0, 0, 390, 398, 5, 19, 0, 0, 391, 398, 5, 20, 0, 0,
		392, 398, 5, 21, 0, 0, 393, 398, 5, 22, 0, 0, 394, 398, 5, 23, 0, 0, 395,
		398, 5, 24, 0, 0, 396, 398, 5, 25, 0, 0, 397, 362, 1, 0, 0, 0, 397, 363,
		1, 0, 0, 0, 397, 365, 1, 0, 0, 0, 397, 366, 1, 0, 0, 0, 397, 367, 1, 0,
		0, 0, 397, 368, 1, 0, 0, 0, 397, 369, 1, 0, 0, 0, 397, 370, 1, 0, 0, 0,
		397, 371, 1, 0, 0, 0, 397, 372, 1, 0, 0, 0, 397, 373, 1, 0, 0, 0, 397,
		374, 1, 0, 0, 0, 397, 375, 1, 0, 0, 0, 397, 376, 1, 0, 0, 0, 397, 377,
		1, 0, 0, 0, 397, 378, 1, 0, 0, 0, 397, 379, 1, 0, 0, 0, 397, 380, 1, 0,
		0, 0, 397, 381, 1, 0, 0, 0, 397, 382, 1, 0, 0, 0, 397, 383, 1, 0, 0, 0,
		397, 384, 1, 0, 0, 0, 397, 385, 1, 0, 0, 0, 397, 386, 1, 0, 0, 0, 397,
		387, 1, 0, 0, 0, 397, 388, 1, 0, 0, 0, 397, 389, 1, 0, 0, 0, 397, 390,
		1, 0, 0, 0, 397, 391, 1, 0, 0, 0, 397, 392, 1, 0, 0, 0, 397, 393, 1, 0,
		0, 0, 397, 394, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 397, 396, 1, 0, 0, 0,
		398, 69, 1, 0, 0, 0, 38, 73, 80, 84, 86, 92, 100, 107, 111, 117, 131, 139,
		149, 154, 161, 188, 197, 201, 211, 215, 217, 227, 235, 243, 247, 249, 262,
		267, 273, 278, 284, 290, 296, 307, 337, 344, 351, 358, 397,
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

// MongoShellParserInit initializes any static state used to implement MongoShellParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMongoShellParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MongoShellParserInit() {
	staticData := &MongoShellParserParserStaticData
	staticData.once.Do(mongoshellparserParserInit)
}

// NewMongoShellParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMongoShellParser(input antlr.TokenStream) *MongoShellParser {
	MongoShellParserInit()
	this := new(MongoShellParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MongoShellParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MongoShellParser.g4"

	return this
}

// MongoShellParser tokens.
const (
	MongoShellParserEOF                  = antlr.TokenEOF
	MongoShellParserSHOW                 = 1
	MongoShellParserDBS                  = 2
	MongoShellParserDATABASES            = 3
	MongoShellParserCOLLECTIONS          = 4
	MongoShellParserDB                   = 5
	MongoShellParserNEW                  = 6
	MongoShellParserTRUE                 = 7
	MongoShellParserFALSE                = 8
	MongoShellParserNULL                 = 9
	MongoShellParserGET_COLLECTION       = 10
	MongoShellParserGET_COLLECTION_NAMES = 11
	MongoShellParserGET_COLLECTION_INFOS = 12
	MongoShellParserOBJECT_ID            = 13
	MongoShellParserISO_DATE             = 14
	MongoShellParserDATE                 = 15
	MongoShellParserUUID                 = 16
	MongoShellParserLONG                 = 17
	MongoShellParserNUMBER_LONG          = 18
	MongoShellParserINT32                = 19
	MongoShellParserNUMBER_INT           = 20
	MongoShellParserDOUBLE               = 21
	MongoShellParserDECIMAL128           = 22
	MongoShellParserNUMBER_DECIMAL       = 23
	MongoShellParserTIMESTAMP            = 24
	MongoShellParserREG_EXP              = 25
	MongoShellParserFIND                 = 26
	MongoShellParserFIND_ONE             = 27
	MongoShellParserSORT                 = 28
	MongoShellParserLIMIT                = 29
	MongoShellParserSKIP_                = 30
	MongoShellParserPROJECTION           = 31
	MongoShellParserPROJECT              = 32
	MongoShellParserLPAREN               = 33
	MongoShellParserRPAREN               = 34
	MongoShellParserLBRACE               = 35
	MongoShellParserRBRACE               = 36
	MongoShellParserLBRACKET             = 37
	MongoShellParserRBRACKET             = 38
	MongoShellParserCOLON                = 39
	MongoShellParserCOMMA                = 40
	MongoShellParserDOT                  = 41
	MongoShellParserSEMI                 = 42
	MongoShellParserDOLLAR               = 43
	MongoShellParserLINE_COMMENT         = 44
	MongoShellParserBLOCK_COMMENT        = 45
	MongoShellParserREGEX_LITERAL        = 46
	MongoShellParserNUMBER               = 47
	MongoShellParserDOUBLE_QUOTED_STRING = 48
	MongoShellParserSINGLE_QUOTED_STRING = 49
	MongoShellParserIDENTIFIER           = 50
	MongoShellParserWS                   = 51
)

// MongoShellParser rules.
const (
	MongoShellParserRULE_program           = 0
	MongoShellParserRULE_statement         = 1
	MongoShellParserRULE_shellCommand      = 2
	MongoShellParserRULE_dbStatement       = 3
	MongoShellParserRULE_collectionAccess  = 4
	MongoShellParserRULE_methodChain       = 5
	MongoShellParserRULE_methodCall        = 6
	MongoShellParserRULE_findMethod        = 7
	MongoShellParserRULE_findOneMethod     = 8
	MongoShellParserRULE_sortMethod        = 9
	MongoShellParserRULE_limitMethod       = 10
	MongoShellParserRULE_skipMethod        = 11
	MongoShellParserRULE_projectionMethod  = 12
	MongoShellParserRULE_genericMethod     = 13
	MongoShellParserRULE_arguments         = 14
	MongoShellParserRULE_argument          = 15
	MongoShellParserRULE_document          = 16
	MongoShellParserRULE_pair              = 17
	MongoShellParserRULE_key               = 18
	MongoShellParserRULE_value             = 19
	MongoShellParserRULE_array             = 20
	MongoShellParserRULE_helperFunction    = 21
	MongoShellParserRULE_objectIdHelper    = 22
	MongoShellParserRULE_isoDateHelper     = 23
	MongoShellParserRULE_dateHelper        = 24
	MongoShellParserRULE_uuidHelper        = 25
	MongoShellParserRULE_longHelper        = 26
	MongoShellParserRULE_int32Helper       = 27
	MongoShellParserRULE_doubleHelper      = 28
	MongoShellParserRULE_decimal128Helper  = 29
	MongoShellParserRULE_timestampHelper   = 30
	MongoShellParserRULE_regExpConstructor = 31
	MongoShellParserRULE_literal           = 32
	MongoShellParserRULE_stringLiteral     = 33
	MongoShellParserRULE_identifier        = 34
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MongoShellParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MongoShellParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MongoShellParserSHOW || _la == MongoShellParserDB {
		{
			p.SetState(70)
			p.Statement()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
		p.Match(MongoShellParserEOF)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ShellCommand() IShellCommandContext
	SEMI() antlr.TerminalNode
	DbStatement() IDbStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ShellCommand() IShellCommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShellCommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShellCommandContext)
}

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSEMI, 0)
}

func (s *StatementContext) DbStatement() IDbStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDbStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDbStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MongoShellParserRULE_statement)
	var _la int

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserSHOW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.ShellCommand()
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserSEMI {
			{
				p.SetState(79)
				p.Match(MongoShellParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MongoShellParserDB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.DbStatement()
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserSEMI {
			{
				p.SetState(83)
				p.Match(MongoShellParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

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

// IShellCommandContext is an interface to support dynamic dispatch.
type IShellCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsShellCommandContext differentiates from other interfaces.
	IsShellCommandContext()
}

type ShellCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShellCommandContext() *ShellCommandContext {
	var p = new(ShellCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_shellCommand
	return p
}

func InitEmptyShellCommandContext(p *ShellCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_shellCommand
}

func (*ShellCommandContext) IsShellCommandContext() {}

func NewShellCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShellCommandContext {
	var p = new(ShellCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_shellCommand

	return p
}

func (s *ShellCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *ShellCommandContext) CopyAll(ctx *ShellCommandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ShellCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShellCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ShowCollectionsContext struct {
	ShellCommandContext
}

func NewShowCollectionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShowCollectionsContext {
	var p = new(ShowCollectionsContext)

	InitEmptyShellCommandContext(&p.ShellCommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*ShellCommandContext))

	return p
}

func (s *ShowCollectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowCollectionsContext) SHOW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSHOW, 0)
}

func (s *ShowCollectionsContext) COLLECTIONS() antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOLLECTIONS, 0)
}

func (s *ShowCollectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterShowCollections(s)
	}
}

func (s *ShowCollectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitShowCollections(s)
	}
}

func (s *ShowCollectionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitShowCollections(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShowDatabasesContext struct {
	ShellCommandContext
}

func NewShowDatabasesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShowDatabasesContext {
	var p = new(ShowDatabasesContext)

	InitEmptyShellCommandContext(&p.ShellCommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*ShellCommandContext))

	return p
}

func (s *ShowDatabasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowDatabasesContext) SHOW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSHOW, 0)
}

func (s *ShowDatabasesContext) DBS() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDBS, 0)
}

func (s *ShowDatabasesContext) DATABASES() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDATABASES, 0)
}

func (s *ShowDatabasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterShowDatabases(s)
	}
}

func (s *ShowDatabasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitShowDatabases(s)
	}
}

func (s *ShowDatabasesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitShowDatabases(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) ShellCommand() (localctx IShellCommandContext) {
	localctx = NewShellCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MongoShellParserRULE_shellCommand)
	var _la int

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewShowDatabasesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(MongoShellParserSHOW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MongoShellParserDBS || _la == MongoShellParserDATABASES) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		localctx = NewShowCollectionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(MongoShellParserSHOW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(MongoShellParserCOLLECTIONS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IDbStatementContext is an interface to support dynamic dispatch.
type IDbStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDbStatementContext differentiates from other interfaces.
	IsDbStatementContext()
}

type DbStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDbStatementContext() *DbStatementContext {
	var p = new(DbStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_dbStatement
	return p
}

func InitEmptyDbStatementContext(p *DbStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_dbStatement
}

func (*DbStatementContext) IsDbStatementContext() {}

func NewDbStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DbStatementContext {
	var p = new(DbStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_dbStatement

	return p
}

func (s *DbStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DbStatementContext) CopyAll(ctx *DbStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DbStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DbStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CollectionOperationContext struct {
	DbStatementContext
}

func NewCollectionOperationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CollectionOperationContext {
	var p = new(CollectionOperationContext)

	InitEmptyDbStatementContext(&p.DbStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*DbStatementContext))

	return p
}

func (s *CollectionOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionOperationContext) DB() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDB, 0)
}

func (s *CollectionOperationContext) CollectionAccess() ICollectionAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionAccessContext)
}

func (s *CollectionOperationContext) MethodChain() IMethodChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodChainContext)
}

func (s *CollectionOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterCollectionOperation(s)
	}
}

func (s *CollectionOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitCollectionOperation(s)
	}
}

func (s *CollectionOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitCollectionOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetCollectionNamesContext struct {
	DbStatementContext
}

func NewGetCollectionNamesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetCollectionNamesContext {
	var p = new(GetCollectionNamesContext)

	InitEmptyDbStatementContext(&p.DbStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*DbStatementContext))

	return p
}

func (s *GetCollectionNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetCollectionNamesContext) DB() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDB, 0)
}

func (s *GetCollectionNamesContext) DOT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOT, 0)
}

func (s *GetCollectionNamesContext) GET_COLLECTION_NAMES() antlr.TerminalNode {
	return s.GetToken(MongoShellParserGET_COLLECTION_NAMES, 0)
}

func (s *GetCollectionNamesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *GetCollectionNamesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *GetCollectionNamesContext) MethodChain() IMethodChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodChainContext)
}

func (s *GetCollectionNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterGetCollectionNames(s)
	}
}

func (s *GetCollectionNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitGetCollectionNames(s)
	}
}

func (s *GetCollectionNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitGetCollectionNames(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetCollectionInfosContext struct {
	DbStatementContext
}

func NewGetCollectionInfosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetCollectionInfosContext {
	var p = new(GetCollectionInfosContext)

	InitEmptyDbStatementContext(&p.DbStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*DbStatementContext))

	return p
}

func (s *GetCollectionInfosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetCollectionInfosContext) DB() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDB, 0)
}

func (s *GetCollectionInfosContext) DOT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOT, 0)
}

func (s *GetCollectionInfosContext) GET_COLLECTION_INFOS() antlr.TerminalNode {
	return s.GetToken(MongoShellParserGET_COLLECTION_INFOS, 0)
}

func (s *GetCollectionInfosContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *GetCollectionInfosContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *GetCollectionInfosContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *GetCollectionInfosContext) MethodChain() IMethodChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodChainContext)
}

func (s *GetCollectionInfosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterGetCollectionInfos(s)
	}
}

func (s *GetCollectionInfosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitGetCollectionInfos(s)
	}
}

func (s *GetCollectionInfosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitGetCollectionInfos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) DbStatement() (localctx IDbStatementContext) {
	localctx = NewDbStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MongoShellParserRULE_dbStatement)
	var _la int

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGetCollectionNamesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(MongoShellParserDB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(MongoShellParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(MongoShellParserGET_COLLECTION_NAMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserDOT {
			{
				p.SetState(99)
				p.MethodChain()
			}

		}

	case 2:
		localctx = NewGetCollectionInfosContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(MongoShellParserDB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(MongoShellParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(MongoShellParserGET_COLLECTION_INFOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1055703028458432) != 0 {
			{
				p.SetState(106)
				p.Arguments()
			}

		}
		{
			p.SetState(109)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserDOT {
			{
				p.SetState(110)
				p.MethodChain()
			}

		}

	case 3:
		localctx = NewCollectionOperationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Match(MongoShellParserDB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.CollectionAccess()
		}
		{
			p.SetState(115)
			p.MethodChain()
		}

	case antlr.ATNInvalidAltNumber:
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

// ICollectionAccessContext is an interface to support dynamic dispatch.
type ICollectionAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCollectionAccessContext differentiates from other interfaces.
	IsCollectionAccessContext()
}

type CollectionAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionAccessContext() *CollectionAccessContext {
	var p = new(CollectionAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_collectionAccess
	return p
}

func InitEmptyCollectionAccessContext(p *CollectionAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_collectionAccess
}

func (*CollectionAccessContext) IsCollectionAccessContext() {}

func NewCollectionAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionAccessContext {
	var p = new(CollectionAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_collectionAccess

	return p
}

func (s *CollectionAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionAccessContext) CopyAll(ctx *CollectionAccessContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CollectionAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DotAccessContext struct {
	CollectionAccessContext
}

func NewDotAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotAccessContext {
	var p = new(DotAccessContext)

	InitEmptyCollectionAccessContext(&p.CollectionAccessContext)
	p.parser = parser
	p.CopyAll(ctx.(*CollectionAccessContext))

	return p
}

func (s *DotAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOT, 0)
}

func (s *DotAccessContext) Identifier() IIdentifierContext {
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

func (s *DotAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterDotAccess(s)
	}
}

func (s *DotAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitDotAccess(s)
	}
}

func (s *DotAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitDotAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetCollectionAccessContext struct {
	CollectionAccessContext
}

func NewGetCollectionAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetCollectionAccessContext {
	var p = new(GetCollectionAccessContext)

	InitEmptyCollectionAccessContext(&p.CollectionAccessContext)
	p.parser = parser
	p.CopyAll(ctx.(*CollectionAccessContext))

	return p
}

func (s *GetCollectionAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetCollectionAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOT, 0)
}

func (s *GetCollectionAccessContext) GET_COLLECTION() antlr.TerminalNode {
	return s.GetToken(MongoShellParserGET_COLLECTION, 0)
}

func (s *GetCollectionAccessContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *GetCollectionAccessContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *GetCollectionAccessContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *GetCollectionAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterGetCollectionAccess(s)
	}
}

func (s *GetCollectionAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitGetCollectionAccess(s)
	}
}

func (s *GetCollectionAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitGetCollectionAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketAccessContext struct {
	CollectionAccessContext
}

func NewBracketAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketAccessContext {
	var p = new(BracketAccessContext)

	InitEmptyCollectionAccessContext(&p.CollectionAccessContext)
	p.parser = parser
	p.CopyAll(ctx.(*CollectionAccessContext))

	return p
}

func (s *BracketAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketAccessContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLBRACKET, 0)
}

func (s *BracketAccessContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *BracketAccessContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRBRACKET, 0)
}

func (s *BracketAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterBracketAccess(s)
	}
}

func (s *BracketAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitBracketAccess(s)
	}
}

func (s *BracketAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitBracketAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) CollectionAccess() (localctx ICollectionAccessContext) {
	localctx = NewCollectionAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MongoShellParserRULE_collectionAccess)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDotAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(MongoShellParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Identifier()
		}

	case 2:
		localctx = NewBracketAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Match(MongoShellParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.StringLiteral()
		}
		{
			p.SetState(123)
			p.Match(MongoShellParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewGetCollectionAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Match(MongoShellParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(MongoShellParserGET_COLLECTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.StringLiteral()
		}
		{
			p.SetState(129)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IMethodChainContext is an interface to support dynamic dispatch.
type IMethodChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllMethodCall() []IMethodCallContext
	MethodCall(i int) IMethodCallContext

	// IsMethodChainContext differentiates from other interfaces.
	IsMethodChainContext()
}

type MethodChainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodChainContext() *MethodChainContext {
	var p = new(MethodChainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_methodChain
	return p
}

func InitEmptyMethodChainContext(p *MethodChainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_methodChain
}

func (*MethodChainContext) IsMethodChainContext() {}

func NewMethodChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodChainContext {
	var p = new(MethodChainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_methodChain

	return p
}

func (s *MethodChainContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodChainContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(MongoShellParserDOT)
}

func (s *MethodChainContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOT, i)
}

func (s *MethodChainContext) AllMethodCall() []IMethodCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodCallContext); ok {
			len++
		}
	}

	tst := make([]IMethodCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodCallContext); ok {
			tst[i] = t.(IMethodCallContext)
			i++
		}
	}

	return tst
}

func (s *MethodChainContext) MethodCall(i int) IMethodCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodCallContext); ok {
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

	return t.(IMethodCallContext)
}

func (s *MethodChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterMethodChain(s)
	}
}

func (s *MethodChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitMethodChain(s)
	}
}

func (s *MethodChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitMethodChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) MethodChain() (localctx IMethodChainContext) {
	localctx = NewMethodChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MongoShellParserRULE_methodChain)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(MongoShellParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.MethodCall()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MongoShellParserDOT {
		{
			p.SetState(135)
			p.Match(MongoShellParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.MethodCall()
		}

		p.SetState(141)
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

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FindMethod() IFindMethodContext
	FindOneMethod() IFindOneMethodContext
	SortMethod() ISortMethodContext
	LimitMethod() ILimitMethodContext
	SkipMethod() ISkipMethodContext
	ProjectionMethod() IProjectionMethodContext
	GenericMethod() IGenericMethodContext

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_methodCall
	return p
}

func InitEmptyMethodCallContext(p *MethodCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_methodCall
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) FindMethod() IFindMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFindMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFindMethodContext)
}

func (s *MethodCallContext) FindOneMethod() IFindOneMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFindOneMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFindOneMethodContext)
}

func (s *MethodCallContext) SortMethod() ISortMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortMethodContext)
}

func (s *MethodCallContext) LimitMethod() ILimitMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitMethodContext)
}

func (s *MethodCallContext) SkipMethod() ISkipMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISkipMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISkipMethodContext)
}

func (s *MethodCallContext) ProjectionMethod() IProjectionMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionMethodContext)
}

func (s *MethodCallContext) GenericMethod() IGenericMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericMethodContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MongoShellParserRULE_methodCall)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.FindMethod()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.FindOneMethod()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.SortMethod()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(145)
			p.LimitMethod()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(146)
			p.SkipMethod()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(147)
			p.ProjectionMethod()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(148)
			p.GenericMethod()
		}

	case antlr.ATNInvalidAltNumber:
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

// IFindMethodContext is an interface to support dynamic dispatch.
type IFindMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FIND() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Argument() IArgumentContext

	// IsFindMethodContext differentiates from other interfaces.
	IsFindMethodContext()
}

type FindMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFindMethodContext() *FindMethodContext {
	var p = new(FindMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_findMethod
	return p
}

func InitEmptyFindMethodContext(p *FindMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_findMethod
}

func (*FindMethodContext) IsFindMethodContext() {}

func NewFindMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FindMethodContext {
	var p = new(FindMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_findMethod

	return p
}

func (s *FindMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *FindMethodContext) FIND() antlr.TerminalNode {
	return s.GetToken(MongoShellParserFIND, 0)
}

func (s *FindMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *FindMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *FindMethodContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *FindMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FindMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterFindMethod(s)
	}
}

func (s *FindMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitFindMethod(s)
	}
}

func (s *FindMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitFindMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) FindMethod() (localctx IFindMethodContext) {
	localctx = NewFindMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MongoShellParserRULE_findMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(MongoShellParserFIND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1055703028458432) != 0 {
		{
			p.SetState(153)
			p.Argument()
		}

	}
	{
		p.SetState(156)
		p.Match(MongoShellParserRPAREN)
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

// IFindOneMethodContext is an interface to support dynamic dispatch.
type IFindOneMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FIND_ONE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Argument() IArgumentContext

	// IsFindOneMethodContext differentiates from other interfaces.
	IsFindOneMethodContext()
}

type FindOneMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFindOneMethodContext() *FindOneMethodContext {
	var p = new(FindOneMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_findOneMethod
	return p
}

func InitEmptyFindOneMethodContext(p *FindOneMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_findOneMethod
}

func (*FindOneMethodContext) IsFindOneMethodContext() {}

func NewFindOneMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FindOneMethodContext {
	var p = new(FindOneMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_findOneMethod

	return p
}

func (s *FindOneMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *FindOneMethodContext) FIND_ONE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserFIND_ONE, 0)
}

func (s *FindOneMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *FindOneMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *FindOneMethodContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *FindOneMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindOneMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FindOneMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterFindOneMethod(s)
	}
}

func (s *FindOneMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitFindOneMethod(s)
	}
}

func (s *FindOneMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitFindOneMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) FindOneMethod() (localctx IFindOneMethodContext) {
	localctx = NewFindOneMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MongoShellParserRULE_findOneMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(MongoShellParserFIND_ONE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1055703028458432) != 0 {
		{
			p.SetState(160)
			p.Argument()
		}

	}
	{
		p.SetState(163)
		p.Match(MongoShellParserRPAREN)
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

// ISortMethodContext is an interface to support dynamic dispatch.
type ISortMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SORT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Document() IDocumentContext
	RPAREN() antlr.TerminalNode

	// IsSortMethodContext differentiates from other interfaces.
	IsSortMethodContext()
}

type SortMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortMethodContext() *SortMethodContext {
	var p = new(SortMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_sortMethod
	return p
}

func InitEmptySortMethodContext(p *SortMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_sortMethod
}

func (*SortMethodContext) IsSortMethodContext() {}

func NewSortMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortMethodContext {
	var p = new(SortMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_sortMethod

	return p
}

func (s *SortMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *SortMethodContext) SORT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSORT, 0)
}

func (s *SortMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *SortMethodContext) Document() IDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *SortMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *SortMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterSortMethod(s)
	}
}

func (s *SortMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitSortMethod(s)
	}
}

func (s *SortMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitSortMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) SortMethod() (localctx ISortMethodContext) {
	localctx = NewSortMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MongoShellParserRULE_sortMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(MongoShellParserSORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Document()
	}
	{
		p.SetState(168)
		p.Match(MongoShellParserRPAREN)
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

// ILimitMethodContext is an interface to support dynamic dispatch.
type ILimitMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsLimitMethodContext differentiates from other interfaces.
	IsLimitMethodContext()
}

type LimitMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitMethodContext() *LimitMethodContext {
	var p = new(LimitMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_limitMethod
	return p
}

func InitEmptyLimitMethodContext(p *LimitMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_limitMethod
}

func (*LimitMethodContext) IsLimitMethodContext() {}

func NewLimitMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitMethodContext {
	var p = new(LimitMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_limitMethod

	return p
}

func (s *LimitMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitMethodContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLIMIT, 0)
}

func (s *LimitMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *LimitMethodContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, 0)
}

func (s *LimitMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *LimitMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterLimitMethod(s)
	}
}

func (s *LimitMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitLimitMethod(s)
	}
}

func (s *LimitMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitLimitMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) LimitMethod() (localctx ILimitMethodContext) {
	localctx = NewLimitMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MongoShellParserRULE_limitMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(MongoShellParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(MongoShellParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(MongoShellParserRPAREN)
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

// ISkipMethodContext is an interface to support dynamic dispatch.
type ISkipMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SKIP_() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsSkipMethodContext differentiates from other interfaces.
	IsSkipMethodContext()
}

type SkipMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipMethodContext() *SkipMethodContext {
	var p = new(SkipMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_skipMethod
	return p
}

func InitEmptySkipMethodContext(p *SkipMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_skipMethod
}

func (*SkipMethodContext) IsSkipMethodContext() {}

func NewSkipMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipMethodContext {
	var p = new(SkipMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_skipMethod

	return p
}

func (s *SkipMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipMethodContext) SKIP_() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSKIP_, 0)
}

func (s *SkipMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *SkipMethodContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, 0)
}

func (s *SkipMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *SkipMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkipMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterSkipMethod(s)
	}
}

func (s *SkipMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitSkipMethod(s)
	}
}

func (s *SkipMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitSkipMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) SkipMethod() (localctx ISkipMethodContext) {
	localctx = NewSkipMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MongoShellParserRULE_skipMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(MongoShellParserSKIP_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(MongoShellParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(MongoShellParserRPAREN)
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

// IProjectionMethodContext is an interface to support dynamic dispatch.
type IProjectionMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Document() IDocumentContext
	RPAREN() antlr.TerminalNode
	PROJECTION() antlr.TerminalNode
	PROJECT() antlr.TerminalNode

	// IsProjectionMethodContext differentiates from other interfaces.
	IsProjectionMethodContext()
}

type ProjectionMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionMethodContext() *ProjectionMethodContext {
	var p = new(ProjectionMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_projectionMethod
	return p
}

func InitEmptyProjectionMethodContext(p *ProjectionMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_projectionMethod
}

func (*ProjectionMethodContext) IsProjectionMethodContext() {}

func NewProjectionMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionMethodContext {
	var p = new(ProjectionMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_projectionMethod

	return p
}

func (s *ProjectionMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *ProjectionMethodContext) Document() IDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *ProjectionMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *ProjectionMethodContext) PROJECTION() antlr.TerminalNode {
	return s.GetToken(MongoShellParserPROJECTION, 0)
}

func (s *ProjectionMethodContext) PROJECT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserPROJECT, 0)
}

func (s *ProjectionMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterProjectionMethod(s)
	}
}

func (s *ProjectionMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitProjectionMethod(s)
	}
}

func (s *ProjectionMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitProjectionMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) ProjectionMethod() (localctx IProjectionMethodContext) {
	localctx = NewProjectionMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MongoShellParserRULE_projectionMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MongoShellParserPROJECTION || _la == MongoShellParserPROJECT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(181)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Document()
	}
	{
		p.SetState(183)
		p.Match(MongoShellParserRPAREN)
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

// IGenericMethodContext is an interface to support dynamic dispatch.
type IGenericMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Arguments() IArgumentsContext

	// IsGenericMethodContext differentiates from other interfaces.
	IsGenericMethodContext()
}

type GenericMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericMethodContext() *GenericMethodContext {
	var p = new(GenericMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_genericMethod
	return p
}

func InitEmptyGenericMethodContext(p *GenericMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_genericMethod
}

func (*GenericMethodContext) IsGenericMethodContext() {}

func NewGenericMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericMethodContext {
	var p = new(GenericMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_genericMethod

	return p
}

func (s *GenericMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericMethodContext) Identifier() IIdentifierContext {
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

func (s *GenericMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *GenericMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *GenericMethodContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *GenericMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterGenericMethod(s)
	}
}

func (s *GenericMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitGenericMethod(s)
	}
}

func (s *GenericMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitGenericMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) GenericMethod() (localctx IGenericMethodContext) {
	localctx = NewGenericMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MongoShellParserRULE_genericMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Identifier()
	}
	{
		p.SetState(186)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1055703028458432) != 0 {
		{
			p.SetState(187)
			p.Arguments()
		}

	}
	{
		p.SetState(190)
		p.Match(MongoShellParserRPAREN)
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
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

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MongoShellParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MongoShellParserRULE_arguments)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Argument()
	}
	p.SetState(197)
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
			{
				p.SetState(193)
				p.Match(MongoShellParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(194)
				p.Argument()
			}

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MongoShellParserCOMMA {
		{
			p.SetState(200)
			p.Match(MongoShellParserCOMMA)
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

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MongoShellParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Value()
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

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllPair() []IPairContext
	Pair(i int) IPairContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLBRACE, 0)
}

func (s *DocumentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRBRACE, 0)
}

func (s *DocumentContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
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

	return t.(IPairContext)
}

func (s *DocumentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MongoShellParserCOMMA)
}

func (s *DocumentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOMMA, i)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (s *DocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MongoShellParserRULE_document)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(MongoShellParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1979129519931390) != 0 {
		{
			p.SetState(206)
			p.Pair()
		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(207)
					p.Match(MongoShellParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(208)
					p.Pair()
				}

			}
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserCOMMA {
			{
				p.SetState(214)
				p.Match(MongoShellParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(219)
		p.Match(MongoShellParserRBRACE)
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

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	COLON() antlr.TerminalNode
	Value() IValueContext

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_pair
	return p
}

func InitEmptyPairContext(p *PairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_pair
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOLON, 0)
}

func (s *PairContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MongoShellParserRULE_pair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Key()
	}
	{
		p.SetState(222)
		p.Match(MongoShellParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Value()
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

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) CopyAll(ctx *KeyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QuotedKeyContext struct {
	KeyContext
}

func NewQuotedKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedKeyContext {
	var p = new(QuotedKeyContext)

	InitEmptyKeyContext(&p.KeyContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeyContext))

	return p
}

func (s *QuotedKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedKeyContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *QuotedKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterQuotedKey(s)
	}
}

func (s *QuotedKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitQuotedKey(s)
	}
}

func (s *QuotedKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitQuotedKey(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnquotedKeyContext struct {
	KeyContext
}

func NewUnquotedKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnquotedKeyContext {
	var p = new(UnquotedKeyContext)

	InitEmptyKeyContext(&p.KeyContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeyContext))

	return p
}

func (s *UnquotedKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnquotedKeyContext) Identifier() IIdentifierContext {
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

func (s *UnquotedKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterUnquotedKey(s)
	}
}

func (s *UnquotedKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitUnquotedKey(s)
	}
}

func (s *UnquotedKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitUnquotedKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MongoShellParserRULE_key)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserSHOW, MongoShellParserDBS, MongoShellParserDATABASES, MongoShellParserCOLLECTIONS, MongoShellParserDB, MongoShellParserNEW, MongoShellParserTRUE, MongoShellParserFALSE, MongoShellParserNULL, MongoShellParserGET_COLLECTION, MongoShellParserGET_COLLECTION_NAMES, MongoShellParserGET_COLLECTION_INFOS, MongoShellParserOBJECT_ID, MongoShellParserISO_DATE, MongoShellParserDATE, MongoShellParserUUID, MongoShellParserLONG, MongoShellParserNUMBER_LONG, MongoShellParserINT32, MongoShellParserNUMBER_INT, MongoShellParserDOUBLE, MongoShellParserDECIMAL128, MongoShellParserNUMBER_DECIMAL, MongoShellParserTIMESTAMP, MongoShellParserREG_EXP, MongoShellParserFIND, MongoShellParserFIND_ONE, MongoShellParserSORT, MongoShellParserLIMIT, MongoShellParserSKIP_, MongoShellParserPROJECTION, MongoShellParserPROJECT, MongoShellParserDOLLAR, MongoShellParserIDENTIFIER:
		localctx = NewUnquotedKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.Identifier()
		}

	case MongoShellParserDOUBLE_QUOTED_STRING, MongoShellParserSINGLE_QUOTED_STRING:
		localctx = NewQuotedKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.StringLiteral()
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RegexpConstructorValueContext struct {
	ValueContext
}

func NewRegexpConstructorValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexpConstructorValueContext {
	var p = new(RegexpConstructorValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *RegexpConstructorValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexpConstructorValueContext) RegExpConstructor() IRegExpConstructorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegExpConstructorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegExpConstructorContext)
}

func (s *RegexpConstructorValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterRegexpConstructorValue(s)
	}
}

func (s *RegexpConstructorValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitRegexpConstructorValue(s)
	}
}

func (s *RegexpConstructorValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitRegexpConstructorValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type RegexLiteralValueContext struct {
	ValueContext
}

func NewRegexLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexLiteralValueContext {
	var p = new(RegexLiteralValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *RegexLiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexLiteralValueContext) REGEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(MongoShellParserREGEX_LITERAL, 0)
}

func (s *RegexLiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterRegexLiteralValue(s)
	}
}

func (s *RegexLiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitRegexLiteralValue(s)
	}
}

func (s *RegexLiteralValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitRegexLiteralValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayValueContext struct {
	ValueContext
}

func NewArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValueContext {
	var p = new(ArrayValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (s *ArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type DocumentValueContext struct {
	ValueContext
}

func NewDocumentValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DocumentValueContext {
	var p = new(DocumentValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *DocumentValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentValueContext) Document() IDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *DocumentValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterDocumentValue(s)
	}
}

func (s *DocumentValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitDocumentValue(s)
	}
}

func (s *DocumentValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitDocumentValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type HelperValueContext struct {
	ValueContext
}

func NewHelperValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HelperValueContext {
	var p = new(HelperValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *HelperValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelperValueContext) HelperFunction() IHelperFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHelperFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHelperFunctionContext)
}

func (s *HelperValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterHelperValue(s)
	}
}

func (s *HelperValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitHelperValue(s)
	}
}

func (s *HelperValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitHelperValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralValueContext struct {
	ValueContext
}

func NewLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralValueContext {
	var p = new(LiteralValueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterLiteralValue(s)
	}
}

func (s *LiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitLiteralValue(s)
	}
}

func (s *LiteralValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitLiteralValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MongoShellParserRULE_value)
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDocumentValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Document()
		}

	case 2:
		localctx = NewArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Array()
		}

	case 3:
		localctx = NewHelperValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.HelperFunction()
		}

	case 4:
		localctx = NewRegexLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
			p.Match(MongoShellParserREGEX_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewRegexpConstructorValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(233)
			p.RegExpConstructor()
		}

	case 6:
		localctx = NewLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(234)
			p.Literal()
		}

	case antlr.ATNInvalidAltNumber:
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

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLBRACKET, 0)
}

func (s *ArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRBRACKET, 0)
}

func (s *ArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MongoShellParserCOMMA)
}

func (s *ArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOMMA, i)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MongoShellParserRULE_array)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(MongoShellParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1055703028458432) != 0 {
		{
			p.SetState(238)
			p.Value()
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(239)
					p.Match(MongoShellParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(240)
					p.Value()
				}

			}
			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserCOMMA {
			{
				p.SetState(246)
				p.Match(MongoShellParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(251)
		p.Match(MongoShellParserRBRACKET)
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

// IHelperFunctionContext is an interface to support dynamic dispatch.
type IHelperFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ObjectIdHelper() IObjectIdHelperContext
	IsoDateHelper() IIsoDateHelperContext
	DateHelper() IDateHelperContext
	UuidHelper() IUuidHelperContext
	LongHelper() ILongHelperContext
	Int32Helper() IInt32HelperContext
	DoubleHelper() IDoubleHelperContext
	Decimal128Helper() IDecimal128HelperContext
	TimestampHelper() ITimestampHelperContext

	// IsHelperFunctionContext differentiates from other interfaces.
	IsHelperFunctionContext()
}

type HelperFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHelperFunctionContext() *HelperFunctionContext {
	var p = new(HelperFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_helperFunction
	return p
}

func InitEmptyHelperFunctionContext(p *HelperFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_helperFunction
}

func (*HelperFunctionContext) IsHelperFunctionContext() {}

func NewHelperFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HelperFunctionContext {
	var p = new(HelperFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_helperFunction

	return p
}

func (s *HelperFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *HelperFunctionContext) ObjectIdHelper() IObjectIdHelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectIdHelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectIdHelperContext)
}

func (s *HelperFunctionContext) IsoDateHelper() IIsoDateHelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsoDateHelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsoDateHelperContext)
}

func (s *HelperFunctionContext) DateHelper() IDateHelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateHelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateHelperContext)
}

func (s *HelperFunctionContext) UuidHelper() IUuidHelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUuidHelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUuidHelperContext)
}

func (s *HelperFunctionContext) LongHelper() ILongHelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILongHelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILongHelperContext)
}

func (s *HelperFunctionContext) Int32Helper() IInt32HelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInt32HelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInt32HelperContext)
}

func (s *HelperFunctionContext) DoubleHelper() IDoubleHelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoubleHelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoubleHelperContext)
}

func (s *HelperFunctionContext) Decimal128Helper() IDecimal128HelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimal128HelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimal128HelperContext)
}

func (s *HelperFunctionContext) TimestampHelper() ITimestampHelperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimestampHelperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimestampHelperContext)
}

func (s *HelperFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HelperFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HelperFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterHelperFunction(s)
	}
}

func (s *HelperFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitHelperFunction(s)
	}
}

func (s *HelperFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitHelperFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) HelperFunction() (localctx IHelperFunctionContext) {
	localctx = NewHelperFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MongoShellParserRULE_helperFunction)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.ObjectIdHelper()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.IsoDateHelper()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.DateHelper()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(256)
			p.UuidHelper()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(257)
			p.LongHelper()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(258)
			p.Int32Helper()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(259)
			p.DoubleHelper()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(260)
			p.Decimal128Helper()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(261)
			p.TimestampHelper()
		}

	case antlr.ATNInvalidAltNumber:
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

// IObjectIdHelperContext is an interface to support dynamic dispatch.
type IObjectIdHelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBJECT_ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	StringLiteral() IStringLiteralContext
	NEW() antlr.TerminalNode

	// IsObjectIdHelperContext differentiates from other interfaces.
	IsObjectIdHelperContext()
}

type ObjectIdHelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectIdHelperContext() *ObjectIdHelperContext {
	var p = new(ObjectIdHelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_objectIdHelper
	return p
}

func InitEmptyObjectIdHelperContext(p *ObjectIdHelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_objectIdHelper
}

func (*ObjectIdHelperContext) IsObjectIdHelperContext() {}

func NewObjectIdHelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectIdHelperContext {
	var p = new(ObjectIdHelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_objectIdHelper

	return p
}

func (s *ObjectIdHelperContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectIdHelperContext) OBJECT_ID() antlr.TerminalNode {
	return s.GetToken(MongoShellParserOBJECT_ID, 0)
}

func (s *ObjectIdHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *ObjectIdHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *ObjectIdHelperContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ObjectIdHelperContext) NEW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNEW, 0)
}

func (s *ObjectIdHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectIdHelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectIdHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterObjectIdHelper(s)
	}
}

func (s *ObjectIdHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitObjectIdHelper(s)
	}
}

func (s *ObjectIdHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitObjectIdHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) ObjectIdHelper() (localctx IObjectIdHelperContext) {
	localctx = NewObjectIdHelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MongoShellParserRULE_objectIdHelper)
	var _la int

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserOBJECT_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Match(MongoShellParserOBJECT_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserDOUBLE_QUOTED_STRING || _la == MongoShellParserSINGLE_QUOTED_STRING {
			{
				p.SetState(266)
				p.StringLiteral()
			}

		}
		{
			p.SetState(269)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNEW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Match(MongoShellParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Match(MongoShellParserOBJECT_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.NotifyErrorListeners("'new' keyword is not supported. Use ObjectId() directly", nil, nil)

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

// IIsoDateHelperContext is an interface to support dynamic dispatch.
type IIsoDateHelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISO_DATE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	StringLiteral() IStringLiteralContext
	NEW() antlr.TerminalNode

	// IsIsoDateHelperContext differentiates from other interfaces.
	IsIsoDateHelperContext()
}

type IsoDateHelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsoDateHelperContext() *IsoDateHelperContext {
	var p = new(IsoDateHelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_isoDateHelper
	return p
}

func InitEmptyIsoDateHelperContext(p *IsoDateHelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_isoDateHelper
}

func (*IsoDateHelperContext) IsIsoDateHelperContext() {}

func NewIsoDateHelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsoDateHelperContext {
	var p = new(IsoDateHelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_isoDateHelper

	return p
}

func (s *IsoDateHelperContext) GetParser() antlr.Parser { return s.parser }

func (s *IsoDateHelperContext) ISO_DATE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserISO_DATE, 0)
}

func (s *IsoDateHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *IsoDateHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *IsoDateHelperContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *IsoDateHelperContext) NEW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNEW, 0)
}

func (s *IsoDateHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsoDateHelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsoDateHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterIsoDateHelper(s)
	}
}

func (s *IsoDateHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitIsoDateHelper(s)
	}
}

func (s *IsoDateHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitIsoDateHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) IsoDateHelper() (localctx IIsoDateHelperContext) {
	localctx = NewIsoDateHelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MongoShellParserRULE_isoDateHelper)
	var _la int

	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserISO_DATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(MongoShellParserISO_DATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserDOUBLE_QUOTED_STRING || _la == MongoShellParserSINGLE_QUOTED_STRING {
			{
				p.SetState(277)
				p.StringLiteral()
			}

		}
		{
			p.SetState(280)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNEW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Match(MongoShellParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(MongoShellParserISO_DATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.NotifyErrorListeners("'new' keyword is not supported. Use ISODate() directly", nil, nil)

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

// IDateHelperContext is an interface to support dynamic dispatch.
type IDateHelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	StringLiteral() IStringLiteralContext
	NUMBER() antlr.TerminalNode
	NEW() antlr.TerminalNode

	// IsDateHelperContext differentiates from other interfaces.
	IsDateHelperContext()
}

type DateHelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateHelperContext() *DateHelperContext {
	var p = new(DateHelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_dateHelper
	return p
}

func InitEmptyDateHelperContext(p *DateHelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_dateHelper
}

func (*DateHelperContext) IsDateHelperContext() {}

func NewDateHelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateHelperContext {
	var p = new(DateHelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_dateHelper

	return p
}

func (s *DateHelperContext) GetParser() antlr.Parser { return s.parser }

func (s *DateHelperContext) DATE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDATE, 0)
}

func (s *DateHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *DateHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *DateHelperContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *DateHelperContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, 0)
}

func (s *DateHelperContext) NEW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNEW, 0)
}

func (s *DateHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateHelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterDateHelper(s)
	}
}

func (s *DateHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitDateHelper(s)
	}
}

func (s *DateHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitDateHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) DateHelper() (localctx IDateHelperContext) {
	localctx = NewDateHelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MongoShellParserRULE_dateHelper)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserDATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.Match(MongoShellParserDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case MongoShellParserDOUBLE_QUOTED_STRING, MongoShellParserSINGLE_QUOTED_STRING:
			{
				p.SetState(288)
				p.StringLiteral()
			}

		case MongoShellParserNUMBER:
			{
				p.SetState(289)
				p.Match(MongoShellParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case MongoShellParserRPAREN:

		default:
		}
		{
			p.SetState(292)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNEW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.Match(MongoShellParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Match(MongoShellParserDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.NotifyErrorListeners("'new' keyword is not supported. Use Date() directly", nil, nil)

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

// IUuidHelperContext is an interface to support dynamic dispatch.
type IUuidHelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UUID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	StringLiteral() IStringLiteralContext
	RPAREN() antlr.TerminalNode

	// IsUuidHelperContext differentiates from other interfaces.
	IsUuidHelperContext()
}

type UuidHelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUuidHelperContext() *UuidHelperContext {
	var p = new(UuidHelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_uuidHelper
	return p
}

func InitEmptyUuidHelperContext(p *UuidHelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_uuidHelper
}

func (*UuidHelperContext) IsUuidHelperContext() {}

func NewUuidHelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UuidHelperContext {
	var p = new(UuidHelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_uuidHelper

	return p
}

func (s *UuidHelperContext) GetParser() antlr.Parser { return s.parser }

func (s *UuidHelperContext) UUID() antlr.TerminalNode {
	return s.GetToken(MongoShellParserUUID, 0)
}

func (s *UuidHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *UuidHelperContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *UuidHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *UuidHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UuidHelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UuidHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterUuidHelper(s)
	}
}

func (s *UuidHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitUuidHelper(s)
	}
}

func (s *UuidHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitUuidHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) UuidHelper() (localctx IUuidHelperContext) {
	localctx = NewUuidHelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MongoShellParserRULE_uuidHelper)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(MongoShellParserUUID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.StringLiteral()
	}
	{
		p.SetState(301)
		p.Match(MongoShellParserRPAREN)
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

// ILongHelperContext is an interface to support dynamic dispatch.
type ILongHelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	LONG() antlr.TerminalNode
	NUMBER_LONG() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	StringLiteral() IStringLiteralContext

	// IsLongHelperContext differentiates from other interfaces.
	IsLongHelperContext()
}

type LongHelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLongHelperContext() *LongHelperContext {
	var p = new(LongHelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_longHelper
	return p
}

func InitEmptyLongHelperContext(p *LongHelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_longHelper
}

func (*LongHelperContext) IsLongHelperContext() {}

func NewLongHelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LongHelperContext {
	var p = new(LongHelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_longHelper

	return p
}

func (s *LongHelperContext) GetParser() antlr.Parser { return s.parser }

func (s *LongHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *LongHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *LongHelperContext) LONG() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLONG, 0)
}

func (s *LongHelperContext) NUMBER_LONG() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER_LONG, 0)
}

func (s *LongHelperContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, 0)
}

func (s *LongHelperContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LongHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongHelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LongHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterLongHelper(s)
	}
}

func (s *LongHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitLongHelper(s)
	}
}

func (s *LongHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitLongHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) LongHelper() (localctx ILongHelperContext) {
	localctx = NewLongHelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MongoShellParserRULE_longHelper)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MongoShellParserLONG || _la == MongoShellParserNUMBER_LONG) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(304)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserNUMBER:
		{
			p.SetState(305)
			p.Match(MongoShellParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDOUBLE_QUOTED_STRING, MongoShellParserSINGLE_QUOTED_STRING:
		{
			p.SetState(306)
			p.StringLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(309)
		p.Match(MongoShellParserRPAREN)
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

// IInt32HelperContext is an interface to support dynamic dispatch.
type IInt32HelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	INT32() antlr.TerminalNode
	NUMBER_INT() antlr.TerminalNode

	// IsInt32HelperContext differentiates from other interfaces.
	IsInt32HelperContext()
}

type Int32HelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt32HelperContext() *Int32HelperContext {
	var p = new(Int32HelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_int32Helper
	return p
}

func InitEmptyInt32HelperContext(p *Int32HelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_int32Helper
}

func (*Int32HelperContext) IsInt32HelperContext() {}

func NewInt32HelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int32HelperContext {
	var p = new(Int32HelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_int32Helper

	return p
}

func (s *Int32HelperContext) GetParser() antlr.Parser { return s.parser }

func (s *Int32HelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *Int32HelperContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, 0)
}

func (s *Int32HelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *Int32HelperContext) INT32() antlr.TerminalNode {
	return s.GetToken(MongoShellParserINT32, 0)
}

func (s *Int32HelperContext) NUMBER_INT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER_INT, 0)
}

func (s *Int32HelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int32HelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int32HelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterInt32Helper(s)
	}
}

func (s *Int32HelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitInt32Helper(s)
	}
}

func (s *Int32HelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitInt32Helper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Int32Helper() (localctx IInt32HelperContext) {
	localctx = NewInt32HelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MongoShellParserRULE_int32Helper)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MongoShellParserINT32 || _la == MongoShellParserNUMBER_INT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(312)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(MongoShellParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(MongoShellParserRPAREN)
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

// IDoubleHelperContext is an interface to support dynamic dispatch.
type IDoubleHelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsDoubleHelperContext differentiates from other interfaces.
	IsDoubleHelperContext()
}

type DoubleHelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleHelperContext() *DoubleHelperContext {
	var p = new(DoubleHelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_doubleHelper
	return p
}

func InitEmptyDoubleHelperContext(p *DoubleHelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_doubleHelper
}

func (*DoubleHelperContext) IsDoubleHelperContext() {}

func NewDoubleHelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleHelperContext {
	var p = new(DoubleHelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_doubleHelper

	return p
}

func (s *DoubleHelperContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleHelperContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOUBLE, 0)
}

func (s *DoubleHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *DoubleHelperContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, 0)
}

func (s *DoubleHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *DoubleHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleHelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoubleHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterDoubleHelper(s)
	}
}

func (s *DoubleHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitDoubleHelper(s)
	}
}

func (s *DoubleHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitDoubleHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) DoubleHelper() (localctx IDoubleHelperContext) {
	localctx = NewDoubleHelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MongoShellParserRULE_doubleHelper)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(MongoShellParserDOUBLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Match(MongoShellParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(MongoShellParserRPAREN)
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

// IDecimal128HelperContext is an interface to support dynamic dispatch.
type IDecimal128HelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	StringLiteral() IStringLiteralContext
	RPAREN() antlr.TerminalNode
	DECIMAL128() antlr.TerminalNode
	NUMBER_DECIMAL() antlr.TerminalNode

	// IsDecimal128HelperContext differentiates from other interfaces.
	IsDecimal128HelperContext()
}

type Decimal128HelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimal128HelperContext() *Decimal128HelperContext {
	var p = new(Decimal128HelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_decimal128Helper
	return p
}

func InitEmptyDecimal128HelperContext(p *Decimal128HelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_decimal128Helper
}

func (*Decimal128HelperContext) IsDecimal128HelperContext() {}

func NewDecimal128HelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decimal128HelperContext {
	var p = new(Decimal128HelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_decimal128Helper

	return p
}

func (s *Decimal128HelperContext) GetParser() antlr.Parser { return s.parser }

func (s *Decimal128HelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *Decimal128HelperContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *Decimal128HelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *Decimal128HelperContext) DECIMAL128() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDECIMAL128, 0)
}

func (s *Decimal128HelperContext) NUMBER_DECIMAL() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER_DECIMAL, 0)
}

func (s *Decimal128HelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decimal128HelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decimal128HelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterDecimal128Helper(s)
	}
}

func (s *Decimal128HelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitDecimal128Helper(s)
	}
}

func (s *Decimal128HelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitDecimal128Helper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Decimal128Helper() (localctx IDecimal128HelperContext) {
	localctx = NewDecimal128HelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MongoShellParserRULE_decimal128Helper)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MongoShellParserDECIMAL128 || _la == MongoShellParserNUMBER_DECIMAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(322)
		p.Match(MongoShellParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.StringLiteral()
	}
	{
		p.SetState(324)
		p.Match(MongoShellParserRPAREN)
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

// ITimestampHelperContext is an interface to support dynamic dispatch.
type ITimestampHelperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTimestampHelperContext differentiates from other interfaces.
	IsTimestampHelperContext()
}

type TimestampHelperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimestampHelperContext() *TimestampHelperContext {
	var p = new(TimestampHelperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_timestampHelper
	return p
}

func InitEmptyTimestampHelperContext(p *TimestampHelperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_timestampHelper
}

func (*TimestampHelperContext) IsTimestampHelperContext() {}

func NewTimestampHelperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimestampHelperContext {
	var p = new(TimestampHelperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_timestampHelper

	return p
}

func (s *TimestampHelperContext) GetParser() antlr.Parser { return s.parser }

func (s *TimestampHelperContext) CopyAll(ctx *TimestampHelperContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TimestampHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampHelperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TimestampArgsHelperContext struct {
	TimestampHelperContext
}

func NewTimestampArgsHelperContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimestampArgsHelperContext {
	var p = new(TimestampArgsHelperContext)

	InitEmptyTimestampHelperContext(&p.TimestampHelperContext)
	p.parser = parser
	p.CopyAll(ctx.(*TimestampHelperContext))

	return p
}

func (s *TimestampArgsHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampArgsHelperContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(MongoShellParserTIMESTAMP, 0)
}

func (s *TimestampArgsHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *TimestampArgsHelperContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(MongoShellParserNUMBER)
}

func (s *TimestampArgsHelperContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, i)
}

func (s *TimestampArgsHelperContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOMMA, 0)
}

func (s *TimestampArgsHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *TimestampArgsHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterTimestampArgsHelper(s)
	}
}

func (s *TimestampArgsHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitTimestampArgsHelper(s)
	}
}

func (s *TimestampArgsHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitTimestampArgsHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

type TimestampDocHelperContext struct {
	TimestampHelperContext
}

func NewTimestampDocHelperContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimestampDocHelperContext {
	var p = new(TimestampDocHelperContext)

	InitEmptyTimestampHelperContext(&p.TimestampHelperContext)
	p.parser = parser
	p.CopyAll(ctx.(*TimestampHelperContext))

	return p
}

func (s *TimestampDocHelperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampDocHelperContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(MongoShellParserTIMESTAMP, 0)
}

func (s *TimestampDocHelperContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *TimestampDocHelperContext) Document() IDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *TimestampDocHelperContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *TimestampDocHelperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterTimestampDocHelper(s)
	}
}

func (s *TimestampDocHelperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitTimestampDocHelper(s)
	}
}

func (s *TimestampDocHelperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitTimestampDocHelper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) TimestampHelper() (localctx ITimestampHelperContext) {
	localctx = NewTimestampHelperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MongoShellParserRULE_timestampHelper)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTimestampDocHelperContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.Match(MongoShellParserTIMESTAMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Document()
		}
		{
			p.SetState(329)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTimestampArgsHelperContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(MongoShellParserTIMESTAMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(MongoShellParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(MongoShellParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(MongoShellParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IRegExpConstructorContext is an interface to support dynamic dispatch.
type IRegExpConstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REG_EXP() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllStringLiteral() []IStringLiteralContext
	StringLiteral(i int) IStringLiteralContext
	RPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	NEW() antlr.TerminalNode

	// IsRegExpConstructorContext differentiates from other interfaces.
	IsRegExpConstructorContext()
}

type RegExpConstructorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegExpConstructorContext() *RegExpConstructorContext {
	var p = new(RegExpConstructorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_regExpConstructor
	return p
}

func InitEmptyRegExpConstructorContext(p *RegExpConstructorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_regExpConstructor
}

func (*RegExpConstructorContext) IsRegExpConstructorContext() {}

func NewRegExpConstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegExpConstructorContext {
	var p = new(RegExpConstructorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_regExpConstructor

	return p
}

func (s *RegExpConstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *RegExpConstructorContext) REG_EXP() antlr.TerminalNode {
	return s.GetToken(MongoShellParserREG_EXP, 0)
}

func (s *RegExpConstructorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLPAREN, 0)
}

func (s *RegExpConstructorContext) AllStringLiteral() []IStringLiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringLiteralContext); ok {
			len++
		}
	}

	tst := make([]IStringLiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringLiteralContext); ok {
			tst[i] = t.(IStringLiteralContext)
			i++
		}
	}

	return tst
}

func (s *RegExpConstructorContext) StringLiteral(i int) IStringLiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
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

	return t.(IStringLiteralContext)
}

func (s *RegExpConstructorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MongoShellParserRPAREN, 0)
}

func (s *RegExpConstructorContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOMMA, 0)
}

func (s *RegExpConstructorContext) NEW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNEW, 0)
}

func (s *RegExpConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegExpConstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegExpConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterRegExpConstructor(s)
	}
}

func (s *RegExpConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitRegExpConstructor(s)
	}
}

func (s *RegExpConstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitRegExpConstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) RegExpConstructor() (localctx IRegExpConstructorContext) {
	localctx = NewRegExpConstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MongoShellParserRULE_regExpConstructor)
	var _la int

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserREG_EXP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.Match(MongoShellParserREG_EXP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(MongoShellParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.StringLiteral()
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MongoShellParserCOMMA {
			{
				p.SetState(342)
				p.Match(MongoShellParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(343)
				p.StringLiteral()
			}

		}
		{
			p.SetState(346)
			p.Match(MongoShellParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNEW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(348)
			p.Match(MongoShellParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Match(MongoShellParserREG_EXP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.NotifyErrorListeners("'new' keyword is not supported. Use RegExp() directly", nil, nil)

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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralValueContext struct {
	LiteralContext
}

func NewStringLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralValueContext {
	var p = new(StringLiteralValueContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralValueContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *StringLiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterStringLiteralValue(s)
	}
}

func (s *StringLiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitStringLiteralValue(s)
	}
}

func (s *StringLiteralValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitStringLiteralValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullLiteralContext struct {
	LiteralContext
}

func NewNullLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullLiteralContext {
	var p = new(NullLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NullLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullLiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNULL, 0)
}

func (s *NullLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterNullLiteral(s)
	}
}

func (s *NullLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitNullLiteral(s)
	}
}

func (s *NullLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitNullLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseLiteralContext struct {
	LiteralContext
}

func NewFalseLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseLiteralContext {
	var p = new(FalseLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FalseLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserFALSE, 0)
}

func (s *FalseLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitFalseLiteral(s)
	}
}

func (s *FalseLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitFalseLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueLiteralContext struct {
	LiteralContext
}

func NewTrueLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueLiteralContext {
	var p = new(TrueLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *TrueLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserTRUE, 0)
}

func (s *TrueLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitTrueLiteral(s)
	}
}

func (s *TrueLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitTrueLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberLiteralContext struct {
	LiteralContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER, 0)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MongoShellParserRULE_literal)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserDOUBLE_QUOTED_STRING, MongoShellParserSINGLE_QUOTED_STRING:
		localctx = NewStringLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.StringLiteral()
		}

	case MongoShellParserNUMBER:
		localctx = NewNumberLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(MongoShellParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserTRUE:
		localctx = NewTrueLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(355)
			p.Match(MongoShellParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserFALSE:
		localctx = NewFalseLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(356)
			p.Match(MongoShellParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNULL:
		localctx = NewNullLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(357)
			p.Match(MongoShellParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE_QUOTED_STRING() antlr.TerminalNode
	SINGLE_QUOTED_STRING() antlr.TerminalNode

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_stringLiteral
	return p
}

func InitEmptyStringLiteralContext(p *StringLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_stringLiteral
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DOUBLE_QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOUBLE_QUOTED_STRING, 0)
}

func (s *StringLiteralContext) SINGLE_QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSINGLE_QUOTED_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MongoShellParserRULE_stringLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MongoShellParserDOUBLE_QUOTED_STRING || _la == MongoShellParserSINGLE_QUOTED_STRING) {
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	DOLLAR() antlr.TerminalNode
	SHOW() antlr.TerminalNode
	DBS() antlr.TerminalNode
	DATABASES() antlr.TerminalNode
	COLLECTIONS() antlr.TerminalNode
	DB() antlr.TerminalNode
	NEW() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NULL() antlr.TerminalNode
	FIND() antlr.TerminalNode
	FIND_ONE() antlr.TerminalNode
	SORT() antlr.TerminalNode
	LIMIT() antlr.TerminalNode
	SKIP_() antlr.TerminalNode
	PROJECTION() antlr.TerminalNode
	PROJECT() antlr.TerminalNode
	GET_COLLECTION() antlr.TerminalNode
	GET_COLLECTION_NAMES() antlr.TerminalNode
	GET_COLLECTION_INFOS() antlr.TerminalNode
	OBJECT_ID() antlr.TerminalNode
	ISO_DATE() antlr.TerminalNode
	DATE() antlr.TerminalNode
	UUID() antlr.TerminalNode
	LONG() antlr.TerminalNode
	NUMBER_LONG() antlr.TerminalNode
	INT32() antlr.TerminalNode
	NUMBER_INT() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode
	DECIMAL128() antlr.TerminalNode
	NUMBER_DECIMAL() antlr.TerminalNode
	TIMESTAMP() antlr.TerminalNode
	REG_EXP() antlr.TerminalNode

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
	p.RuleIndex = MongoShellParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MongoShellParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoShellParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MongoShellParserIDENTIFIER, 0)
}

func (s *IdentifierContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOLLAR, 0)
}

func (s *IdentifierContext) SHOW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSHOW, 0)
}

func (s *IdentifierContext) DBS() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDBS, 0)
}

func (s *IdentifierContext) DATABASES() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDATABASES, 0)
}

func (s *IdentifierContext) COLLECTIONS() antlr.TerminalNode {
	return s.GetToken(MongoShellParserCOLLECTIONS, 0)
}

func (s *IdentifierContext) DB() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDB, 0)
}

func (s *IdentifierContext) NEW() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNEW, 0)
}

func (s *IdentifierContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserTRUE, 0)
}

func (s *IdentifierContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserFALSE, 0)
}

func (s *IdentifierContext) NULL() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNULL, 0)
}

func (s *IdentifierContext) FIND() antlr.TerminalNode {
	return s.GetToken(MongoShellParserFIND, 0)
}

func (s *IdentifierContext) FIND_ONE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserFIND_ONE, 0)
}

func (s *IdentifierContext) SORT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSORT, 0)
}

func (s *IdentifierContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLIMIT, 0)
}

func (s *IdentifierContext) SKIP_() antlr.TerminalNode {
	return s.GetToken(MongoShellParserSKIP_, 0)
}

func (s *IdentifierContext) PROJECTION() antlr.TerminalNode {
	return s.GetToken(MongoShellParserPROJECTION, 0)
}

func (s *IdentifierContext) PROJECT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserPROJECT, 0)
}

func (s *IdentifierContext) GET_COLLECTION() antlr.TerminalNode {
	return s.GetToken(MongoShellParserGET_COLLECTION, 0)
}

func (s *IdentifierContext) GET_COLLECTION_NAMES() antlr.TerminalNode {
	return s.GetToken(MongoShellParserGET_COLLECTION_NAMES, 0)
}

func (s *IdentifierContext) GET_COLLECTION_INFOS() antlr.TerminalNode {
	return s.GetToken(MongoShellParserGET_COLLECTION_INFOS, 0)
}

func (s *IdentifierContext) OBJECT_ID() antlr.TerminalNode {
	return s.GetToken(MongoShellParserOBJECT_ID, 0)
}

func (s *IdentifierContext) ISO_DATE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserISO_DATE, 0)
}

func (s *IdentifierContext) DATE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDATE, 0)
}

func (s *IdentifierContext) UUID() antlr.TerminalNode {
	return s.GetToken(MongoShellParserUUID, 0)
}

func (s *IdentifierContext) LONG() antlr.TerminalNode {
	return s.GetToken(MongoShellParserLONG, 0)
}

func (s *IdentifierContext) NUMBER_LONG() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER_LONG, 0)
}

func (s *IdentifierContext) INT32() antlr.TerminalNode {
	return s.GetToken(MongoShellParserINT32, 0)
}

func (s *IdentifierContext) NUMBER_INT() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER_INT, 0)
}

func (s *IdentifierContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDOUBLE, 0)
}

func (s *IdentifierContext) DECIMAL128() antlr.TerminalNode {
	return s.GetToken(MongoShellParserDECIMAL128, 0)
}

func (s *IdentifierContext) NUMBER_DECIMAL() antlr.TerminalNode {
	return s.GetToken(MongoShellParserNUMBER_DECIMAL, 0)
}

func (s *IdentifierContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(MongoShellParserTIMESTAMP, 0)
}

func (s *IdentifierContext) REG_EXP() antlr.TerminalNode {
	return s.GetToken(MongoShellParserREG_EXP, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoShellParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MongoShellParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MongoShellParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MongoShellParserRULE_identifier)
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MongoShellParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.Match(MongoShellParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDOLLAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.Match(MongoShellParserDOLLAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Match(MongoShellParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserSHOW:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(365)
			p.Match(MongoShellParserSHOW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDBS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(366)
			p.Match(MongoShellParserDBS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDATABASES:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(367)
			p.Match(MongoShellParserDATABASES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserCOLLECTIONS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(368)
			p.Match(MongoShellParserCOLLECTIONS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDB:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(369)
			p.Match(MongoShellParserDB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNEW:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(370)
			p.Match(MongoShellParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserTRUE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(371)
			p.Match(MongoShellParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserFALSE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(372)
			p.Match(MongoShellParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNULL:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(373)
			p.Match(MongoShellParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserFIND:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(374)
			p.Match(MongoShellParserFIND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserFIND_ONE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(375)
			p.Match(MongoShellParserFIND_ONE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserSORT:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(376)
			p.Match(MongoShellParserSORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserLIMIT:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(377)
			p.Match(MongoShellParserLIMIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserSKIP_:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(378)
			p.Match(MongoShellParserSKIP_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserPROJECTION:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(379)
			p.Match(MongoShellParserPROJECTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserPROJECT:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(380)
			p.Match(MongoShellParserPROJECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserGET_COLLECTION:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(381)
			p.Match(MongoShellParserGET_COLLECTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserGET_COLLECTION_NAMES:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(382)
			p.Match(MongoShellParserGET_COLLECTION_NAMES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserGET_COLLECTION_INFOS:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(383)
			p.Match(MongoShellParserGET_COLLECTION_INFOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserOBJECT_ID:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(384)
			p.Match(MongoShellParserOBJECT_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserISO_DATE:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(385)
			p.Match(MongoShellParserISO_DATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDATE:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(386)
			p.Match(MongoShellParserDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserUUID:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(387)
			p.Match(MongoShellParserUUID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserLONG:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(388)
			p.Match(MongoShellParserLONG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNUMBER_LONG:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(389)
			p.Match(MongoShellParserNUMBER_LONG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserINT32:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(390)
			p.Match(MongoShellParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNUMBER_INT:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(391)
			p.Match(MongoShellParserNUMBER_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDOUBLE:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(392)
			p.Match(MongoShellParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserDECIMAL128:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(393)
			p.Match(MongoShellParserDECIMAL128)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserNUMBER_DECIMAL:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(394)
			p.Match(MongoShellParserNUMBER_DECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserTIMESTAMP:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(395)
			p.Match(MongoShellParserTIMESTAMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MongoShellParserREG_EXP:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(396)
			p.Match(MongoShellParserREG_EXP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
