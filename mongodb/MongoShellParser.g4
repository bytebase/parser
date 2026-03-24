/*
 * MongoDB Shell (mongosh) Parser Grammar
 * For use with ANTLR 4
 *
 * Supports all mongosh commands:
 * - Collection methods (48): find, insertOne, updateMany, aggregate, etc.
 * - Cursor methods (34): sort, limit, skip, forEach, toArray, etc.
 * - Database methods (39): createCollection, dropDatabase, stats, etc.
 * - Bulk operations (21): initializeOrderedBulkOp, Bulk.insert, etc.
 * - Connection methods (16): Mongo, connect, getDB, setReadPref, etc.
 * - Replication methods (14): rs.status, rs.initiate, rs.add, etc.
 * - Sharding methods (52): sh.status, sh.enableSharding, sh.shardCollection, etc.
 * - User management (12): createUser, dropUser, auth, etc.
 * - Role management (10): createRole, dropRole, grantRolesToUser, etc.
 * - Encryption methods (17): KeyVault, ClientEncryption, getKeyVault, etc.
 * - Native methods (16): cat, load, quit, pwd, etc.
 * - Object constructors (18): ObjectId, ISODate, UUID, BinData, etc.
 * - Query plan cache (5): getPlanCache, clear, list, etc.
 * - Atlas-specific (13): Search Index (4), Stream Processing (9)
 * - Shell commands: show dbs, show databases, show collections
 * - Document syntax with unquoted keys and trailing commas
 */

parser grammar MongoShellParser;

options { tokenVocab=MongoShellLexer; }

// Entry point - a program can contain multiple statements
program
    : statement* EOF
    ;

// A statement is either a shell command, a database statement, a bulk statement, a connection statement, a replica set statement, a sharding statement, an encryption statement, a plan cache statement, a stream processing statement, or a native function call
statement
    : shellCommand SEMI?
    | dbStatement SEMI?
    | bulkStatement SEMI?
    | connectionStatement SEMI?
    | rsStatement SEMI?
    | shStatement SEMI?
    | encryptionStatement SEMI?
    | planCacheStatement SEMI?
    | spStatement SEMI?
    | nativeFunctionCall SEMI?
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
    | DB DOT CREATE_COLLECTION LPAREN arguments RPAREN                          # createCollection
    | DB DOT DROP_DATABASE LPAREN RPAREN                                        # dropDatabase
    | DB DOT STATS LPAREN argument? RPAREN                                      # dbStats
    | DB DOT SERVER_STATUS LPAREN argument? RPAREN                              # serverStatus
    | DB DOT SERVER_BUILD_INFO LPAREN RPAREN                                    # serverBuildInfo
    | DB DOT VERSION LPAREN RPAREN                                              # dbVersion
    | DB DOT HOST_INFO LPAREN RPAREN                                            # hostInfo
    | DB DOT LIST_COMMANDS LPAREN RPAREN                                        # listCommands
    | DB DOT RUN_COMMAND LPAREN arguments RPAREN                                # runCommand
    | DB DOT ADMIN_COMMAND LPAREN arguments RPAREN                              # adminCommand
    | DB DOT GET_NAME LPAREN RPAREN                                             # getName
    | DB DOT GET_MONGO LPAREN RPAREN                                            # getMongo
    | DB DOT GET_SIBLING_DB LPAREN argument RPAREN                              # getSiblingDB
    | DB DOT AGGREGATE LPAREN arguments? RPAREN                                # dbAggregate
    | DB DOT AUTH LPAREN arguments? RPAREN                                     # dbAuth
    | DB DOT CHANGE_USER_PASSWORD LPAREN arguments? RPAREN                     # dbChangeUserPassword
    | DB DOT CLONE_DATABASE LPAREN arguments? RPAREN                           # dbCloneDatabase
    | DB DOT COMMAND_HELP LPAREN arguments? RPAREN                             # dbCommandHelp
    | DB DOT COPY_DATABASE LPAREN arguments? RPAREN                            # dbCopyDatabase
    | DB DOT CREATE_ROLE LPAREN arguments? RPAREN                              # dbCreateRole
    | DB DOT CREATE_USER LPAREN arguments? RPAREN                              # dbCreateUser
    | DB DOT CREATE_VIEW LPAREN arguments? RPAREN                              # dbCreateView
    | DB DOT CURRENT_OP LPAREN arguments? RPAREN                               # dbCurrentOp
    | DB DOT DROP_ALL_ROLES LPAREN arguments? RPAREN                           # dbDropAllRoles
    | DB DOT DROP_ALL_USERS LPAREN arguments? RPAREN                           # dbDropAllUsers
    | DB DOT DROP_ROLE LPAREN arguments? RPAREN                                # dbDropRole
    | DB DOT DROP_USER LPAREN arguments? RPAREN                                # dbDropUser
    | DB DOT FSYNC_LOCK LPAREN arguments? RPAREN                               # dbFsyncLock
    | DB DOT FSYNC_UNLOCK LPAREN arguments? RPAREN                             # dbFsyncUnlock
    | DB DOT GET_LOG_COMPONENTS LPAREN arguments? RPAREN                       # dbGetLogComponents
    | DB DOT GET_PROFILING_LEVEL LPAREN arguments? RPAREN                      # dbGetProfilingLevel
    | DB DOT GET_PROFILING_STATUS LPAREN arguments? RPAREN                     # dbGetProfilingStatus
    | DB DOT GET_REPLICATION_INFO LPAREN arguments? RPAREN                     # dbGetReplicationInfo
    | DB DOT GET_ROLE LPAREN arguments? RPAREN                                 # dbGetRole
    | DB DOT GET_ROLES LPAREN arguments? RPAREN                                # dbGetRoles
    | DB DOT GET_USER LPAREN arguments? RPAREN                                 # dbGetUser
    | DB DOT GET_USERS LPAREN arguments? RPAREN                                # dbGetUsers
    | DB DOT GRANT_PRIVILEGES_TO_ROLE LPAREN arguments? RPAREN                 # dbGrantPrivilegesToRole
    | DB DOT GRANT_ROLES_TO_ROLE LPAREN arguments? RPAREN                      # dbGrantRolesToRole
    | DB DOT GRANT_ROLES_TO_USER LPAREN arguments? RPAREN                      # dbGrantRolesToUser
    | DB DOT HELLO LPAREN arguments? RPAREN                                    # dbHello
    | DB DOT IS_MASTER LPAREN arguments? RPAREN                                # dbIsMaster
    | DB DOT KILL_OP LPAREN arguments? RPAREN                                  # dbKillOp
    | DB DOT LOGOUT LPAREN arguments? RPAREN                                   # dbLogout
    | DB DOT PRINT_COLLECTION_STATS LPAREN arguments? RPAREN                   # dbPrintCollectionStats
    | DB DOT PRINT_REPLICATION_INFO LPAREN arguments? RPAREN                   # dbPrintReplicationInfo
    | DB DOT PRINT_SECONDARY_REPLICATION_INFO LPAREN arguments? RPAREN         # dbPrintSecondaryReplicationInfo
    | DB DOT PRINT_SHARDING_STATUS LPAREN arguments? RPAREN                    # dbPrintShardingStatus
    | DB DOT PRINT_SLAVE_REPLICATION_INFO LPAREN arguments? RPAREN             # dbPrintSlaveReplicationInfo
    | DB DOT REVOKE_PRIVILEGES_FROM_ROLE LPAREN arguments? RPAREN              # dbRevokePrivilegesFromRole
    | DB DOT REVOKE_ROLES_FROM_ROLE LPAREN arguments? RPAREN                   # dbRevokeRolesFromRole
    | DB DOT REVOKE_ROLES_FROM_USER LPAREN arguments? RPAREN                   # dbRevokeRolesFromUser
    | DB DOT ROTATE_CERTIFICATES LPAREN arguments? RPAREN                      # dbRotateCertificates
    | DB DOT SET_LOG_LEVEL LPAREN arguments? RPAREN                            # dbSetLogLevel
    | DB DOT SET_PROFILING_LEVEL LPAREN arguments? RPAREN                      # dbSetProfilingLevel
    | DB DOT SET_SECONDARY_OK LPAREN arguments? RPAREN                         # dbSetSecondaryOk
    | DB DOT SET_WRITE_CONCERN LPAREN arguments? RPAREN                        # dbSetWriteConcern
    | DB DOT SHUTDOWN_SERVER LPAREN arguments? RPAREN                          # dbShutdownServer
    | DB DOT UPDATE_ROLE LPAREN arguments? RPAREN                              # dbUpdateRole
    | DB DOT UPDATE_USER LPAREN arguments? RPAREN                              # dbUpdateUser
    | DB DOT WATCH LPAREN arguments? RPAREN                                    # dbWatch
    | DB collectionAccess methodChain                                           # collectionOperation
    ;

// Bulk operation statements
// Pattern: db.collection.initializeOrderedBulkOp().find(...).update(...).execute()
bulkStatement
    : DB collectionAccess DOT bulkInitMethod bulkMethodChain?
    ;

bulkInitMethod
    : INITIALIZE_ORDERED_BULK_OP LPAREN RPAREN
    | INITIALIZE_UNORDERED_BULK_OP LPAREN RPAREN
    ;

bulkMethodChain
    : (DOT bulkMethod)+
    ;

bulkMethod
    : FIND LPAREN argument RPAREN                    # bulkFind
    | INSERT LPAREN argument RPAREN                  # bulkInsert
    | REMOVE LPAREN RPAREN                           # bulkRemove
    | EXECUTE LPAREN argument? RPAREN                # bulkExecute
    | GET_OPERATIONS LPAREN RPAREN                   # bulkGetOperations
    | TO_STRING LPAREN RPAREN                        # bulkToString
    | identifier LPAREN arguments? RPAREN            # bulkGenericMethod
    ;

// Connection statements - top-level Mongo() constructor and connect() function
connectionStatement
    : MONGO LPAREN arguments? RPAREN connectionMethodChain?        # mongoConnection
    | CONNECT LPAREN arguments? RPAREN connectionMethodChain?      # connectCall
    | DB DOT GET_MONGO LPAREN RPAREN connectionMethodChain         # dbGetMongoChain
    ;

// Connection method chain for chaining methods on a connection
connectionMethodChain
    : (DOT connectionMethod)+
    ;

// Replication (replica set) statements - rs.method()
rsStatement
    : RS DOT identifier LPAREN arguments? RPAREN
    ;

// Sharding statements - sh.method()
shStatement
    : SH DOT identifier LPAREN arguments? RPAREN
    ;

// Encryption statements - db.getMongo().getKeyVault().xxx() or db.getMongo().getClientEncryption().xxx()
encryptionStatement
    : DB DOT GET_MONGO LPAREN RPAREN DOT GET_KEY_VAULT LPAREN RPAREN (DOT identifier LPAREN arguments? RPAREN)*     # keyVaultStatement
    | DB DOT GET_MONGO LPAREN RPAREN DOT GET_CLIENT_ENCRYPTION LPAREN RPAREN (DOT identifier LPAREN arguments? RPAREN)*  # clientEncryptionStatement
    ;

// Plan cache statements - db.collection.getPlanCache().xxx()
planCacheStatement
    : DB collectionAccess DOT GET_PLAN_CACHE LPAREN RPAREN (DOT identifier LPAREN arguments? RPAREN)*
    ;

// Stream processing statements - sp.method() or sp.processor.method()
spStatement
    : SP DOT identifier LPAREN arguments? RPAREN
    | SP DOT identifier DOT identifier LPAREN arguments? RPAREN
    ;

// Native shell function calls - top-level functions like cat(), load(), quit()
nativeFunctionCall
    : identifier LPAREN arguments? RPAREN
    ;

// Connection methods that can be called on a Mongo connection object
connectionMethod
    : GET_DB LPAREN argument RPAREN                           # connGetDB
    | GET_READ_CONCERN LPAREN RPAREN                          # connGetReadConcern
    | GET_READ_PREF LPAREN RPAREN                             # connGetReadPref
    | GET_READ_PREF_MODE LPAREN RPAREN                        # connGetReadPrefMode
    | GET_READ_PREF_TAG_SET LPAREN RPAREN                     # connGetReadPrefTagSet
    | GET_WRITE_CONCERN LPAREN RPAREN                         # connGetWriteConcern
    | SET_READ_PREF LPAREN arguments RPAREN                   # connSetReadPref
    | SET_READ_CONCERN LPAREN argument RPAREN                 # connSetReadConcern
    | SET_WRITE_CONCERN LPAREN argument RPAREN                # connSetWriteConcern
    | START_SESSION LPAREN argument? RPAREN                   # connStartSession
    | WATCH LPAREN arguments? RPAREN                          # connWatch
    | CLOSE LPAREN RPAREN                                     # connClose
    | ADMIN_COMMAND LPAREN arguments RPAREN                   # connAdminCommand
    | GET_DB_NAMES LPAREN RPAREN                              # connGetDBNames
    | identifier LPAREN arguments? RPAREN                     # connGenericMethod
    ;

// Collection access patterns
collectionAccess
    : DOT identifier                                    # dotAccess
    | LBRACKET stringLiteral RBRACKET                   # bracketAccess
    | DOT GET_COLLECTION LPAREN stringLiteral RPAREN    # getCollectionAccess
    ;

// Method chain: first call must be a collection method, subsequent calls are cursor methods.
// Special case: explain() returns an explainable object that supports collection methods.
methodChain
    : DOT collectionExplainMethod DOT collectionMethodCall (DOT cursorMethodCall)*
    | DOT collectionMethodCall (DOT cursorMethodCall)*
    ;

// Collection method call: methods that operate on a collection directly
collectionMethodCall
    : findMethod
    | findOneMethod
    | countDocumentsMethod
    | estimatedDocumentCountMethod
    | distinctMethod
    | aggregateMethod
    | getIndexesMethod
    | insertOneMethod
    | insertManyMethod
    | updateOneMethod
    | updateManyMethod
    | deleteOneMethod
    | deleteManyMethod
    | replaceOneMethod
    | findOneAndUpdateMethod
    | findOneAndReplaceMethod
    | findOneAndDeleteMethod
    | createIndexMethod
    | createIndexesMethod
    | dropIndexMethod
    | dropIndexesMethod
    | dropMethod
    | renameCollectionMethod
    | statsMethod
    | storageSizeMethod
    | totalIndexSizeMethod
    | totalSizeMethod
    | dataSizeMethod
    | isCappedMethod
    | validateMethod
    | latencyStatsMethod
    | watchMethod
    | bulkWriteMethod
    | collectionCountMethod
    | collectionInsertMethod
    | collectionRemoveMethod
    | updateMethod
    | mapReduceMethod
    | findAndModifyMethod
    | collectionExplainMethod
    | analyzeShardKeyMethod
    | configureQueryAnalyzerMethod
    | compactStructuredEncryptionDataMethod
    | hideIndexMethod
    | unhideIndexMethod
    | reIndexMethod
    | getShardDistributionMethod
    | getShardVersionMethod
    | createSearchIndexMethod
    | createSearchIndexesMethod
    | dropSearchIndexMethod
    | updateSearchIndexMethod
    ;

// Cursor method call: methods that operate on a cursor (chainable after collection methods)
cursorMethodCall
    : sortMethod
    | limitMethod
    | skipMethod
    | countMethod
    | projectionMethod
    | batchSizeMethod
    | closeMethod
    | collationMethod
    | commentMethod
    | explainMethod
    | forEachMethod
    | hasNextMethod
    | hintMethod
    | isClosedMethod
    | isExhaustedMethod
    | itcountMethod
    | mapMethod
    | maxMethod
    | maxAwaitTimeMSMethod
    | maxTimeMSMethod
    | minMethod
    | nextMethod
    | noCursorTimeoutMethod
    | objsLeftInBatchMethod
    | prettyMethod
    | readConcernMethod
    | readPrefMethod
    | returnKeyMethod
    | showRecordIdMethod
    | sizeMethod
    | tailableMethod
    | toArrayMethod
    | tryNextMethod
    | allowDiskUseMethod
    | addOptionMethod
    ;

// Specific method rules for better AST structure
// find(filter?, projection?)
findMethod
    : FIND LPAREN arguments? RPAREN
    ;

// findOne(filter?, projection?)
findOneMethod
    : FIND_ONE LPAREN arguments? RPAREN
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

// aggregate(pipeline?, options?)
aggregateMethod
    : AGGREGATE LPAREN arguments? RPAREN
    ;

// getIndexes()
getIndexesMethod
    : GET_INDEXES LPAREN RPAREN
    ;

// insertOne(document, options?)
insertOneMethod
    : INSERT_ONE LPAREN arguments RPAREN
    ;

// insertMany(documents, options?)
insertManyMethod
    : INSERT_MANY LPAREN arguments RPAREN
    ;

// updateOne(filter, update, options?)
updateOneMethod
    : UPDATE_ONE LPAREN arguments RPAREN
    ;

// updateMany(filter, update, options?)
updateManyMethod
    : UPDATE_MANY LPAREN arguments RPAREN
    ;

// deleteOne(filter, options?)
deleteOneMethod
    : DELETE_ONE LPAREN arguments RPAREN
    ;

// deleteMany(filter, options?)
deleteManyMethod
    : DELETE_MANY LPAREN arguments RPAREN
    ;

// replaceOne(filter, replacement, options?)
replaceOneMethod
    : REPLACE_ONE LPAREN arguments RPAREN
    ;

// findOneAndUpdate(filter, update, options?)
findOneAndUpdateMethod
    : FIND_ONE_AND_UPDATE LPAREN arguments RPAREN
    ;

// findOneAndReplace(filter, replacement, options?)
findOneAndReplaceMethod
    : FIND_ONE_AND_REPLACE LPAREN arguments RPAREN
    ;

// findOneAndDelete(filter, options?)
findOneAndDeleteMethod
    : FIND_ONE_AND_DELETE LPAREN arguments RPAREN
    ;

// createIndex(keys, options?)
createIndexMethod
    : CREATE_INDEX LPAREN arguments RPAREN
    ;

// createIndexes(keyPatterns, options?)
createIndexesMethod
    : CREATE_INDEXES LPAREN arguments RPAREN
    ;

// dropIndex(index)
dropIndexMethod
    : DROP_INDEX LPAREN argument RPAREN
    ;

// dropIndexes(indexes?)
dropIndexesMethod
    : DROP_INDEXES LPAREN argument? RPAREN
    ;

// drop(options?)
dropMethod
    : DROP LPAREN argument? RPAREN
    ;

// renameCollection(newName, dropTarget?)
renameCollectionMethod
    : RENAME_COLLECTION LPAREN arguments RPAREN
    ;

// stats(options?)
statsMethod
    : STATS LPAREN argument? RPAREN
    ;

// storageSize()
storageSizeMethod
    : STORAGE_SIZE LPAREN RPAREN
    ;

// totalIndexSize()
totalIndexSizeMethod
    : TOTAL_INDEX_SIZE LPAREN RPAREN
    ;

// totalSize()
totalSizeMethod
    : TOTAL_SIZE LPAREN RPAREN
    ;

// dataSize()
dataSizeMethod
    : DATA_SIZE LPAREN RPAREN
    ;

// isCapped()
isCappedMethod
    : IS_CAPPED LPAREN RPAREN
    ;

// validate(options?)
validateMethod
    : VALIDATE LPAREN argument? RPAREN
    ;

// latencyStats(options?)
latencyStatsMethod
    : LATENCY_STATS LPAREN argument? RPAREN
    ;

// watch(pipeline?, options?)
watchMethod
    : WATCH LPAREN arguments? RPAREN
    ;

// bulkWrite(operations, options?)
bulkWriteMethod
    : BULK_WRITE LPAREN arguments RPAREN
    ;

// count(filter?, options?) - deprecated collection-level count
collectionCountMethod
    : COUNT LPAREN arguments? RPAREN
    ;

// insert(document/array, options?) - deprecated
collectionInsertMethod
    : INSERT LPAREN arguments RPAREN
    ;

// remove(filter, options?) - deprecated
collectionRemoveMethod
    : REMOVE LPAREN arguments RPAREN
    ;

// update(filter, update, options?) - deprecated
updateMethod
    : UPDATE LPAREN arguments RPAREN
    ;

// mapReduce(map, reduce, options?) - deprecated
mapReduceMethod
    : MAP_REDUCE LPAREN arguments RPAREN
    ;

// findAndModify(document)
findAndModifyMethod
    : FIND_AND_MODIFY LPAREN arguments RPAREN
    ;

// explain(verbosity?) - collection-level explain
collectionExplainMethod
    : EXPLAIN LPAREN arguments? RPAREN
    ;

// analyzeShardKey(key, options?)
analyzeShardKeyMethod
    : ANALYZE_SHARD_KEY LPAREN arguments RPAREN
    ;

// configureQueryAnalyzer(options)
configureQueryAnalyzerMethod
    : CONFIGURE_QUERY_ANALYZER LPAREN arguments RPAREN
    ;

// compactStructuredEncryptionData(options?)
compactStructuredEncryptionDataMethod
    : COMPACT_STRUCTURED_ENCRYPTION_DATA LPAREN arguments? RPAREN
    ;

// hideIndex(indexName/spec)
hideIndexMethod
    : HIDE_INDEX LPAREN argument RPAREN
    ;

// unhideIndex(indexName/spec)
unhideIndexMethod
    : UNHIDE_INDEX LPAREN argument RPAREN
    ;

// reIndex()
reIndexMethod
    : RE_INDEX LPAREN RPAREN
    ;

// getShardDistribution()
getShardDistributionMethod
    : GET_SHARD_DISTRIBUTION LPAREN RPAREN
    ;

// getShardVersion()
getShardVersionMethod
    : GET_SHARD_VERSION LPAREN RPAREN
    ;

// createSearchIndex(definition)
createSearchIndexMethod
    : CREATE_SEARCH_INDEX LPAREN arguments RPAREN
    ;

// createSearchIndexes(definitions)
createSearchIndexesMethod
    : CREATE_SEARCH_INDEXES LPAREN arguments RPAREN
    ;

// dropSearchIndex(name)
dropSearchIndexMethod
    : DROP_SEARCH_INDEX LPAREN argument RPAREN
    ;

// updateSearchIndex(name, definition)
updateSearchIndexMethod
    : UPDATE_SEARCH_INDEX LPAREN arguments RPAREN
    ;

// sort(specification?) - can be called without args
sortMethod
    : SORT LPAREN document? RPAREN
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

// projection(doc?) - can be called without args
projectionMethod
    : (PROJECTION | PROJECT) LPAREN document? RPAREN
    ;

// Cursor methods
batchSizeMethod
    : BATCH_SIZE LPAREN NUMBER RPAREN
    ;

closeMethod
    : CLOSE LPAREN RPAREN
    ;

// collation(doc?) - can be called without args
collationMethod
    : COLLATION LPAREN document? RPAREN
    ;

// comment(str?) - can be called without args
commentMethod
    : COMMENT LPAREN stringLiteral? RPAREN
    ;

explainMethod
    : EXPLAIN LPAREN stringLiteral? RPAREN
    ;

forEachMethod
    : FOR_EACH LPAREN argument RPAREN
    ;

hasNextMethod
    : HAS_NEXT LPAREN RPAREN
    ;

// hint(indexSpec?) - can be called without args
hintMethod
    : HINT LPAREN argument? RPAREN
    ;

isClosedMethod
    : IS_CLOSED LPAREN RPAREN
    ;

isExhaustedMethod
    : IS_EXHAUSTED LPAREN RPAREN
    ;

itcountMethod
    : IT_COUNT LPAREN RPAREN
    ;

mapMethod
    : MAP LPAREN argument RPAREN
    ;

// max(indexBounds?) - can be called without args
maxMethod
    : MAX LPAREN document? RPAREN
    ;

maxAwaitTimeMSMethod
    : MAX_AWAIT_TIME_MS LPAREN NUMBER RPAREN
    ;

maxTimeMSMethod
    : MAX_TIME_MS LPAREN NUMBER RPAREN
    ;

// min(indexBounds?) - can be called without args
minMethod
    : MIN LPAREN document? RPAREN
    ;

nextMethod
    : NEXT LPAREN RPAREN
    ;

noCursorTimeoutMethod
    : NO_CURSOR_TIMEOUT LPAREN RPAREN
    ;

objsLeftInBatchMethod
    : OBJS_LEFT_IN_BATCH LPAREN RPAREN
    ;

prettyMethod
    : PRETTY LPAREN RPAREN
    ;

// readConcern(doc?) - can be called without args
readConcernMethod
    : READ_CONCERN LPAREN document? RPAREN
    ;

readPrefMethod
    : READ_PREF LPAREN arguments RPAREN
    ;

// returnKey(bool?) - can be called without args
returnKeyMethod
    : RETURN_KEY LPAREN (TRUE | FALSE)? RPAREN
    ;

// showRecordId(bool?) - can be called without args
showRecordIdMethod
    : SHOW_RECORD_ID LPAREN (TRUE | FALSE)? RPAREN
    ;

sizeMethod
    : SIZE LPAREN RPAREN
    ;

tailableMethod
    : TAILABLE LPAREN (TRUE | FALSE)? RPAREN
    ;

toArrayMethod
    : TO_ARRAY LPAREN RPAREN
    ;

tryNextMethod
    : TRY_NEXT LPAREN RPAREN
    ;

allowDiskUseMethod
    : ALLOW_DISK_USE LPAREN (TRUE | FALSE)? RPAREN
    ;

addOptionMethod
    : ADD_OPTION LPAREN NUMBER RPAREN
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
    : NEW (OBJECT_ID | ISO_DATE | DATE | UUID | LONG | NUMBER_LONG | INT32 | NUMBER_INT | DOUBLE | DECIMAL128 | NUMBER_DECIMAL | TIMESTAMP | REG_EXP | BIN_DATA | BINARY | BSON_REG_EXP | HEX_DATA)
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
    | binDataHelper
    | binaryHelper
    | bsonRegExpHelper
    | hexDataHelper
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

// Int32(n), Int32("n"), NumberInt(n), NumberInt("n")
int32Helper
    : (INT32 | NUMBER_INT) LPAREN (NUMBER | stringLiteral) RPAREN
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

// BinData(subtype, base64) - Binary data with subtype
binDataHelper
    : BIN_DATA LPAREN NUMBER COMMA stringLiteral RPAREN
    ;

// Binary(buffer, subtype) or Binary.createFromBase64(base64, subtype)
binaryHelper
    : BINARY LPAREN arguments RPAREN
    | BINARY DOT identifier LPAREN arguments RPAREN
    ;

// BSONRegExp(pattern, flags)
bsonRegExpHelper
    : BSON_REG_EXP LPAREN arguments RPAREN
    ;

// HexData(subtype, hex)
hexDataHelper
    : HEX_DATA LPAREN NUMBER COMMA stringLiteral RPAREN
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
    | INSERT_ONE
    | INSERT_MANY
    | UPDATE_ONE
    | UPDATE_MANY
    | DELETE_ONE
    | DELETE_MANY
    | REPLACE_ONE
    | FIND_ONE_AND_UPDATE
    | FIND_ONE_AND_REPLACE
    | FIND_ONE_AND_DELETE
    | CREATE_INDEX
    | CREATE_INDEXES
    | DROP_INDEX
    | DROP_INDEXES
    | DROP
    | RENAME_COLLECTION
    | STATS
    | STORAGE_SIZE
    | TOTAL_INDEX_SIZE
    | TOTAL_SIZE
    | DATA_SIZE
    | IS_CAPPED
    | VALIDATE
    | LATENCY_STATS
    | SORT
    | LIMIT
    | SKIP_
    | COUNT
    | PROJECTION
    | PROJECT
    | GET_COLLECTION
    | GET_COLLECTION_NAMES
    | GET_COLLECTION_INFOS
    | CREATE_COLLECTION
    | DROP_DATABASE
    | HOST_INFO
    | LIST_COMMANDS
    | SERVER_BUILD_INFO
    | SERVER_STATUS
    | VERSION
    | RUN_COMMAND
    | ADMIN_COMMAND
    | GET_NAME
    | GET_MONGO
    | GET_SIBLING_DB
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
    | BIN_DATA
    | BINARY
    | BSON_REG_EXP
    | HEX_DATA
    // Cursor method tokens
    | BATCH_SIZE
    | CLOSE
    | COLLATION
    | COMMENT
    | EXPLAIN
    | FOR_EACH
    | HAS_NEXT
    | HINT
    | IS_CLOSED
    | IS_EXHAUSTED
    | IT_COUNT
    | MAP
    | MAX
    | MAX_AWAIT_TIME_MS
    | MAX_TIME_MS
    | MIN
    | NEXT
    | NO_CURSOR_TIMEOUT
    | OBJS_LEFT_IN_BATCH
    | PRETTY
    | READ_CONCERN
    | READ_PREF
    | RETURN_KEY
    | SHOW_RECORD_ID
    | SIZE
    | TAILABLE
    | TO_ARRAY
    | TRY_NEXT
    | ALLOW_DISK_USE
    | ADD_OPTION
    // Bulk operation tokens
    | INITIALIZE_ORDERED_BULK_OP
    | INITIALIZE_UNORDERED_BULK_OP
    | EXECUTE
    | GET_OPERATIONS
    | TO_STRING
    | INSERT
    | REMOVE
    // Connection method tokens
    | MONGO
    | CONNECT
    | GET_DB
    | GET_READ_CONCERN
    | GET_READ_PREF
    | GET_READ_PREF_MODE
    | GET_READ_PREF_TAG_SET
    | GET_WRITE_CONCERN
    | SET_READ_PREF
    | SET_READ_CONCERN
    | SET_WRITE_CONCERN
    | START_SESSION
    | WATCH
    | GET_DB_NAMES
    // Replication method tokens
    | RS
    // Sharding method tokens
    | SH
    // Stream processing token
    | SP
    // Encryption method tokens
    | GET_KEY_VAULT
    | GET_CLIENT_ENCRYPTION
    // Plan cache method tokens
    | GET_PLAN_CACHE
    // Collection method tokens (additional)
    | BULK_WRITE
    | UPDATE
    | MAP_REDUCE
    | FIND_AND_MODIFY
    | ANALYZE_SHARD_KEY
    | CONFIGURE_QUERY_ANALYZER
    | COMPACT_STRUCTURED_ENCRYPTION_DATA
    | HIDE_INDEX
    | UNHIDE_INDEX
    | RE_INDEX
    | GET_SHARD_DISTRIBUTION
    | GET_SHARD_VERSION
    | CREATE_SEARCH_INDEX
    | CREATE_SEARCH_INDEXES
    | DROP_SEARCH_INDEX
    | UPDATE_SEARCH_INDEX
    // Database method tokens (additional)
    | AUTH
    | CHANGE_USER_PASSWORD
    | CLONE_DATABASE
    | COMMAND_HELP
    | COPY_DATABASE
    | CREATE_ROLE
    | CREATE_USER
    | CREATE_VIEW
    | CURRENT_OP
    | DROP_ALL_ROLES
    | DROP_ALL_USERS
    | DROP_ROLE
    | DROP_USER
    | FSYNC_LOCK
    | FSYNC_UNLOCK
    | GET_LOG_COMPONENTS
    | GET_PROFILING_LEVEL
    | GET_PROFILING_STATUS
    | GET_REPLICATION_INFO
    | GET_ROLE
    | GET_ROLES
    | GET_USER
    | GET_USERS
    | GRANT_PRIVILEGES_TO_ROLE
    | GRANT_ROLES_TO_ROLE
    | GRANT_ROLES_TO_USER
    | HELLO
    | IS_MASTER
    | KILL_OP
    | LOGOUT
    | PRINT_COLLECTION_STATS
    | PRINT_REPLICATION_INFO
    | PRINT_SECONDARY_REPLICATION_INFO
    | PRINT_SHARDING_STATUS
    | PRINT_SLAVE_REPLICATION_INFO
    | REVOKE_PRIVILEGES_FROM_ROLE
    | REVOKE_ROLES_FROM_ROLE
    | REVOKE_ROLES_FROM_USER
    | ROTATE_CERTIFICATES
    | SET_LOG_LEVEL
    | SET_PROFILING_LEVEL
    | SET_SECONDARY_OK
    | SHUTDOWN_SERVER
    | UPDATE_ROLE
    | UPDATE_USER
    ;
