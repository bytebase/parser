/*
 * MongoDB Shell (mongosh) Lexer Grammar
 * For use with ANTLR 4
 */

lexer grammar MongoShellLexer;

// Keywords
SHOW: 'show';
DBS: 'dbs';
DATABASES: 'databases';
COLLECTIONS: 'collections';
DB: 'db';
NEW: 'new';
TRUE: 'true';
FALSE: 'false';
NULL: 'null';
GET_COLLECTION: 'getCollection';
GET_COLLECTION_NAMES: 'getCollectionNames';

// Helper function names (recognized as distinct tokens)
OBJECT_ID: 'ObjectId';
ISO_DATE: 'ISODate';
DATE: 'Date';
UUID: 'UUID';
LONG: 'Long';
NUMBER_LONG: 'NumberLong';
INT32: 'Int32';
NUMBER_INT: 'NumberInt';
DOUBLE: 'Double';
DECIMAL128: 'Decimal128';
NUMBER_DECIMAL: 'NumberDecimal';
TIMESTAMP: 'Timestamp';
REG_EXP: 'RegExp';

// Cursor modifiers (methods)
FIND: 'find';
FIND_ONE: 'findOne';
SORT: 'sort';
LIMIT: 'limit';
SKIP_: 'skip';
PROJECTION: 'projection';
PROJECT: 'project';

// Punctuation
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
COLON: ':';
COMMA: ',';
DOT: '.';
SEMI: ';';

// Operators (for query operators like $gt, $lt, etc.)
DOLLAR: '$';

// Comments - must come before REGEX_LITERAL to properly capture /* ... */
LINE_COMMENT
    : '//' ~[\r\n]* -> channel(HIDDEN)
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;

// Regex literal
REGEX_LITERAL
    : '/' REGEX_BODY '/' REGEX_FLAGS?
    ;

fragment REGEX_BODY
    : REGEX_CHAR+
    ;

fragment REGEX_CHAR
    : ~[/\r\n\\]
    | '\\' .
    ;

fragment REGEX_FLAGS
    : [gimsuy]+
    ;

// Numbers
NUMBER
    : '-'? INT ('.' [0-9]+)? EXPONENT?
    | '-'? '.' [0-9]+ EXPONENT?
    ;

fragment INT
    : '0'
    | [1-9] [0-9]*
    ;

fragment EXPONENT
    : [eE] [+-]? [0-9]+
    ;

// Strings - both single and double quoted
DOUBLE_QUOTED_STRING
    : '"' (ESC | ~["\\])* '"'
    ;

SINGLE_QUOTED_STRING
    : '\'' (ESC | ~['\\])* '\''
    ;

fragment ESC
    : '\\' (["\\/bfnrt] | UNICODE | '\'')
    ;

fragment UNICODE
    : 'u' HEX HEX HEX HEX
    ;

fragment HEX
    : [0-9a-fA-F]
    ;

// Identifiers - for unquoted keys, collection names, method names
// Allows $-prefixed identifiers for MongoDB operators like $gt, $in, etc.
IDENTIFIER
    : [$_a-zA-Z] [$_a-zA-Z0-9]*
    ;

// Whitespace
WS
    : [ \t\r\n]+ -> channel(HIDDEN)
    ;
