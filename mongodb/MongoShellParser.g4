/*
 * MongoDB Shell (mongosh) Parser Grammar
 * For use with ANTLR 4
 *
 * Milestone 1: Read Operations + Utility + Aggregation
 * - Shell commands: show dbs, show databases, show collections
 * - Utility: db.getCollectionNames(), db.getCollectionInfos()
 * - Collection info: db.collection.getIndexes()
 * - Read methods: find(), findOne(), countDocuments(), estimatedDocumentCount(), distinct()
 * - Aggregation: db.collection.aggregate()
 * - Cursor modifiers: sort(), limit(), skip(), count(), projection(), project()
 * - Object constructors: ObjectId(), ISODate(), UUID(), NumberInt(), NumberLong(), NumberDecimal()
 * - Document syntax with unquoted keys and trailing commas
 */

parser grammar MongoShellParser;

options { tokenVocab=MongoShellLexer; }

// Entry point - a program can contain multiple statements
program
    : statement* EOF
    ;

// A statement is either a shell command or a database statement
statement
    : shellCommand SEMI?
    | dbStatement SEMI?
    ;

// Shell commands: show dbs, show databases, show collections
shellCommand
    : SHOW (DBS | DATABASES)                    # showDatabases
    | SHOW COLLECTIONS                          # showCollections
    ;

// Database statements: db.collection.method(...) or db.getCollectionNames() or db.getCollectionInfos()
dbStatement
    : DB DOT GET_COLLECTION_NAMES LPAREN RPAREN methodChain?                    # getCollectionNames
    | DB DOT GET_COLLECTION_INFOS LPAREN arguments? RPAREN methodChain?         # getCollectionInfos
    | DB collectionAccess methodChain                                           # collectionOperation
    ;

// Collection access patterns
collectionAccess
    : DOT identifier                                    # dotAccess
    | LBRACKET stringLiteral RBRACKET                   # bracketAccess
    | DOT GET_COLLECTION LPAREN stringLiteral RPAREN    # getCollectionAccess
    ;

// Method chain: one or more method calls chained with dots
methodChain
    : DOT methodCall (DOT methodCall)*
    ;

// Method call: methodName(arguments?)
methodCall
    : findMethod
    | findOneMethod
    | countDocumentsMethod
    | estimatedDocumentCountMethod
    | distinctMethod
    | aggregateMethod
    | getIndexesMethod
    | sortMethod
    | limitMethod
    | skipMethod
    | countMethod
    | projectionMethod
    | genericMethod
    ;

// Specific method rules for better AST structure
findMethod
    : FIND LPAREN argument? RPAREN
    ;

findOneMethod
    : FIND_ONE LPAREN argument? RPAREN
    ;

// countDocuments(filter?, options?)
countDocumentsMethod
    : COUNT_DOCUMENTS LPAREN arguments? RPAREN
    ;

// estimatedDocumentCount(options?)
estimatedDocumentCountMethod
    : ESTIMATED_DOCUMENT_COUNT LPAREN argument? RPAREN
    ;

// distinct(field, query?, options?)
distinctMethod
    : DISTINCT LPAREN arguments RPAREN
    ;

// aggregate(pipeline, options?)
aggregateMethod
    : AGGREGATE LPAREN arguments RPAREN
    ;

// getIndexes()
getIndexesMethod
    : GET_INDEXES LPAREN RPAREN
    ;

sortMethod
    : SORT LPAREN document RPAREN
    ;

limitMethod
    : LIMIT LPAREN NUMBER RPAREN
    ;

skipMethod
    : SKIP_ LPAREN NUMBER RPAREN
    ;

// cursor.count() - returns count of documents matching the query
countMethod
    : COUNT LPAREN RPAREN
    ;

projectionMethod
    : (PROJECTION | PROJECT) LPAREN document RPAREN
    ;

// Generic method for extensibility (other methods will be caught here)
genericMethod
    : identifier LPAREN arguments? RPAREN
    ;

// Arguments: comma-separated list of values
arguments
    : argument (COMMA argument)* COMMA?
    ;

argument
    : value
    ;

// Document: { key: value, ... } with optional trailing comma
document
    : LBRACE (pair (COMMA pair)* COMMA?)? RBRACE
    ;

// Key-value pair
pair
    : key COLON value
    ;

// Key: can be unquoted identifier or quoted string
key
    : identifier      # unquotedKey
    | stringLiteral   # quotedKey
    ;

// Value: document, array, helper function, regex, or literal
value
    : document                # documentValue
    | array                   # arrayValue
    | helperFunction          # helperValue
    | REGEX_LITERAL           # regexLiteralValue
    | regExpConstructor       # regexpConstructorValue
    | literal                 # literalValue
    | newKeywordError         # newKeywordValue
    ;

// Catch 'new' keyword usage and provide helpful error message
newKeywordError
    : NEW (OBJECT_ID | ISO_DATE | DATE | UUID | LONG | NUMBER_LONG | INT32 | NUMBER_INT | DOUBLE | DECIMAL128 | NUMBER_DECIMAL | TIMESTAMP | REG_EXP)
      { p.NotifyErrorListeners("'new' keyword is not supported. Use ObjectId(), ISODate(), UUID(), etc. directly without 'new'", nil, nil) }
      LPAREN arguments? RPAREN
    ;

// Array: [ value, ... ] with optional trailing comma
array
    : LBRACKET (value (COMMA value)* COMMA?)? RBRACKET
    ;

// Helper functions - each is a distinct node type for easy AST walking
// Note: 'new' keyword is not supported. Use ObjectId(), ISODate(), Date() directly.
helperFunction
    : objectIdHelper
    | isoDateHelper
    | dateHelper
    | uuidHelper
    | longHelper
    | int32Helper
    | doubleHelper
    | decimal128Helper
    | timestampHelper
    ;

// ObjectId("hex") or ObjectId()
objectIdHelper
    : OBJECT_ID LPAREN stringLiteral? RPAREN
    ;

// ISODate("iso-string") or ISODate()
isoDateHelper
    : ISO_DATE LPAREN stringLiteral? RPAREN
    ;

// Date() or Date("string") or Date(timestamp)
dateHelper
    : DATE LPAREN (stringLiteral | NUMBER)? RPAREN
    ;

// UUID("uuid-string")
uuidHelper
    : UUID LPAREN stringLiteral RPAREN
    ;

// Long(n), Long("n"), NumberLong(n), NumberLong("n")
longHelper
    : (LONG | NUMBER_LONG) LPAREN (NUMBER | stringLiteral) RPAREN
    ;

// Int32(n), NumberInt(n)
int32Helper
    : (INT32 | NUMBER_INT) LPAREN NUMBER RPAREN
    ;

// Double(n)
doubleHelper
    : DOUBLE LPAREN NUMBER RPAREN
    ;

// Decimal128("n"), NumberDecimal("n")
decimal128Helper
    : (DECIMAL128 | NUMBER_DECIMAL) LPAREN stringLiteral RPAREN
    ;

// Timestamp({t: n, i: n}) or Timestamp(t, i)
timestampHelper
    : TIMESTAMP LPAREN document RPAREN                    # timestampDocHelper
    | TIMESTAMP LPAREN NUMBER COMMA NUMBER RPAREN         # timestampArgsHelper
    ;

// RegExp("pattern", "flags") constructor
regExpConstructor
    : REG_EXP LPAREN stringLiteral (COMMA stringLiteral)? RPAREN
    ;

// Literals
literal
    : stringLiteral     # stringLiteralValue
    | NUMBER            # numberLiteral
    | TRUE              # trueLiteral
    | FALSE             # falseLiteral
    | NULL              # nullLiteral
    ;

// String literal - both single and double quoted
stringLiteral
    : DOUBLE_QUOTED_STRING
    | SINGLE_QUOTED_STRING
    ;

// Identifier - used for unquoted keys, collection names, method names
// Includes MongoDB operators like $gt, $in, etc.
identifier
    : IDENTIFIER
    | DOLLAR IDENTIFIER
    // Keywords that can also be used as identifiers
    | SHOW
    | DBS
    | DATABASES
    | COLLECTIONS
    | DB
    | NEW
    | TRUE
    | FALSE
    | NULL
    | FIND
    | FIND_ONE
    | COUNT_DOCUMENTS
    | ESTIMATED_DOCUMENT_COUNT
    | DISTINCT
    | AGGREGATE
    | GET_INDEXES
    | SORT
    | LIMIT
    | SKIP_
    | COUNT
    | PROJECTION
    | PROJECT
    | GET_COLLECTION
    | GET_COLLECTION_NAMES
    | GET_COLLECTION_INFOS
    | OBJECT_ID
    | ISO_DATE
    | DATE
    | UUID
    | LONG
    | NUMBER_LONG
    | INT32
    | NUMBER_INT
    | DOUBLE
    | DECIMAL128
    | NUMBER_DECIMAL
    | TIMESTAMP
    | REG_EXP
    ;
