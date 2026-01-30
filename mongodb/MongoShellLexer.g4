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
GET_COLLECTION_INFOS: 'getCollectionInfos';

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

// Additional object constructors
BIN_DATA: 'BinData';
BINARY: 'Binary';
BSON_REG_EXP: 'BSONRegExp';
HEX_DATA: 'HexData';

// Collection methods
FIND: 'find';
FIND_ONE: 'findOne';
COUNT_DOCUMENTS: 'countDocuments';
ESTIMATED_DOCUMENT_COUNT: 'estimatedDocumentCount';
DISTINCT: 'distinct';
AGGREGATE: 'aggregate';
GET_INDEXES: 'getIndexes';

// Collection write methods (M2)
INSERT_ONE: 'insertOne';
INSERT_MANY: 'insertMany';
UPDATE_ONE: 'updateOne';
UPDATE_MANY: 'updateMany';
DELETE_ONE: 'deleteOne';
DELETE_MANY: 'deleteMany';
REPLACE_ONE: 'replaceOne';
FIND_ONE_AND_UPDATE: 'findOneAndUpdate';
FIND_ONE_AND_REPLACE: 'findOneAndReplace';
FIND_ONE_AND_DELETE: 'findOneAndDelete';

// Collection index/schema methods (M3)
CREATE_INDEX: 'createIndex';
CREATE_INDEXES: 'createIndexes';
DROP_INDEX: 'dropIndex';
DROP_INDEXES: 'dropIndexes';
DROP: 'drop';
RENAME_COLLECTION: 'renameCollection';
STATS: 'stats';
STORAGE_SIZE: 'storageSize';
TOTAL_INDEX_SIZE: 'totalIndexSize';
TOTAL_SIZE: 'totalSize';
DATA_SIZE: 'dataSize';
IS_CAPPED: 'isCapped';
VALIDATE: 'validate';
LATENCY_STATS: 'latencyStats';

// Collection methods (additional)
BULK_WRITE: 'bulkWrite';
UPDATE: 'update';
MAP_REDUCE: 'mapReduce';
FIND_AND_MODIFY: 'findAndModify';
ANALYZE_SHARD_KEY: 'analyzeShardKey';
CONFIGURE_QUERY_ANALYZER: 'configureQueryAnalyzer';
COMPACT_STRUCTURED_ENCRYPTION_DATA: 'compactStructuredEncryptionData';
HIDE_INDEX: 'hideIndex';
UNHIDE_INDEX: 'unhideIndex';
RE_INDEX: 'reIndex';
GET_SHARD_DISTRIBUTION: 'getShardDistribution';
GET_SHARD_VERSION: 'getShardVersion';
// Atlas Search Index methods
CREATE_SEARCH_INDEX: 'createSearchIndex';
CREATE_SEARCH_INDEXES: 'createSearchIndexes';
DROP_SEARCH_INDEX: 'dropSearchIndex';
UPDATE_SEARCH_INDEX: 'updateSearchIndex';

// Database methods (M4)
CREATE_COLLECTION: 'createCollection';
DROP_DATABASE: 'dropDatabase';
HOST_INFO: 'hostInfo';
LIST_COMMANDS: 'listCommands';
SERVER_BUILD_INFO: 'serverBuildInfo';
SERVER_STATUS: 'serverStatus';
VERSION: 'version';
RUN_COMMAND: 'runCommand';
ADMIN_COMMAND: 'adminCommand';
GET_NAME: 'getName';
GET_MONGO: 'getMongo';
GET_SIBLING_DB: 'getSiblingDB';

// Connection methods
MONGO: 'Mongo';
CONNECT: 'connect';

// Replication methods
RS: 'rs';

// Sharding methods
SH: 'sh';

// Atlas stream processing
SP: 'sp';

GET_DB: 'getDB';
GET_READ_CONCERN: 'getReadConcern';
GET_READ_PREF: 'getReadPref';
GET_READ_PREF_MODE: 'getReadPrefMode';
GET_READ_PREF_TAG_SET: 'getReadPrefTagSet';
GET_WRITE_CONCERN: 'getWriteConcern';
SET_READ_PREF: 'setReadPref';
SET_READ_CONCERN: 'setReadConcern';
SET_WRITE_CONCERN: 'setWriteConcern';
START_SESSION: 'startSession';
WATCH: 'watch';
GET_DB_NAMES: 'getDBNames';

// Encryption methods
GET_KEY_VAULT: 'getKeyVault';
GET_CLIENT_ENCRYPTION: 'getClientEncryption';

// Plan cache methods
GET_PLAN_CACHE: 'getPlanCache';

// Cursor modifiers (methods)
SORT: 'sort';
LIMIT: 'limit';
SKIP_: 'skip';
PROJECTION: 'projection';
PROJECT: 'project';
COUNT: 'count';

// Bulk operation methods
INITIALIZE_ORDERED_BULK_OP: 'initializeOrderedBulkOp';
INITIALIZE_UNORDERED_BULK_OP: 'initializeUnorderedBulkOp';
EXECUTE: 'execute';
GET_OPERATIONS: 'getOperations';
TO_STRING: 'toString';
INSERT: 'insert';
REMOVE: 'remove';

// Additional cursor methods
BATCH_SIZE: 'batchSize';
CLOSE: 'close';
COLLATION: 'collation';
COMMENT: 'comment';
EXPLAIN: 'explain';
FOR_EACH: 'forEach';
HAS_NEXT: 'hasNext';
HINT: 'hint';
IS_CLOSED: 'isClosed';
IS_EXHAUSTED: 'isExhausted';
IT_COUNT: 'itcount';
MAP: 'map';
MAX: 'max';
MAX_AWAIT_TIME_MS: 'maxAwaitTimeMS';
MAX_TIME_MS: 'maxTimeMS';
MIN: 'min';
NEXT: 'next';
NO_CURSOR_TIMEOUT: 'noCursorTimeout';
OBJS_LEFT_IN_BATCH: 'objsLeftInBatch';
PRETTY: 'pretty';
READ_CONCERN: 'readConcern';
READ_PREF: 'readPref';
RETURN_KEY: 'returnKey';
SHOW_RECORD_ID: 'showRecordId';
SIZE: 'size';
TAILABLE: 'tailable';
TO_ARRAY: 'toArray';
TRY_NEXT: 'tryNext';
ALLOW_DISK_USE: 'allowDiskUse';
ADD_OPTION: 'addOption';

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
