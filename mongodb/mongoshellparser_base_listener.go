// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

// BaseMongoShellParserListener is a complete listener for a parse tree produced by MongoShellParser.
type BaseMongoShellParserListener struct{}

var _ MongoShellParserListener = &BaseMongoShellParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMongoShellParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMongoShellParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMongoShellParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMongoShellParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMongoShellParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMongoShellParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMongoShellParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMongoShellParserListener) ExitStatement(ctx *StatementContext) {}

// EnterShowDatabases is called when production showDatabases is entered.
func (s *BaseMongoShellParserListener) EnterShowDatabases(ctx *ShowDatabasesContext) {}

// ExitShowDatabases is called when production showDatabases is exited.
func (s *BaseMongoShellParserListener) ExitShowDatabases(ctx *ShowDatabasesContext) {}

// EnterShowCollections is called when production showCollections is entered.
func (s *BaseMongoShellParserListener) EnterShowCollections(ctx *ShowCollectionsContext) {}

// ExitShowCollections is called when production showCollections is exited.
func (s *BaseMongoShellParserListener) ExitShowCollections(ctx *ShowCollectionsContext) {}

// EnterGetCollectionNames is called when production getCollectionNames is entered.
func (s *BaseMongoShellParserListener) EnterGetCollectionNames(ctx *GetCollectionNamesContext) {}

// ExitGetCollectionNames is called when production getCollectionNames is exited.
func (s *BaseMongoShellParserListener) ExitGetCollectionNames(ctx *GetCollectionNamesContext) {}

// EnterGetCollectionInfos is called when production getCollectionInfos is entered.
func (s *BaseMongoShellParserListener) EnterGetCollectionInfos(ctx *GetCollectionInfosContext) {}

// ExitGetCollectionInfos is called when production getCollectionInfos is exited.
func (s *BaseMongoShellParserListener) ExitGetCollectionInfos(ctx *GetCollectionInfosContext) {}

// EnterCreateCollection is called when production createCollection is entered.
func (s *BaseMongoShellParserListener) EnterCreateCollection(ctx *CreateCollectionContext) {}

// ExitCreateCollection is called when production createCollection is exited.
func (s *BaseMongoShellParserListener) ExitCreateCollection(ctx *CreateCollectionContext) {}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseMongoShellParserListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseMongoShellParserListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterDbStats is called when production dbStats is entered.
func (s *BaseMongoShellParserListener) EnterDbStats(ctx *DbStatsContext) {}

// ExitDbStats is called when production dbStats is exited.
func (s *BaseMongoShellParserListener) ExitDbStats(ctx *DbStatsContext) {}

// EnterServerStatus is called when production serverStatus is entered.
func (s *BaseMongoShellParserListener) EnterServerStatus(ctx *ServerStatusContext) {}

// ExitServerStatus is called when production serverStatus is exited.
func (s *BaseMongoShellParserListener) ExitServerStatus(ctx *ServerStatusContext) {}

// EnterServerBuildInfo is called when production serverBuildInfo is entered.
func (s *BaseMongoShellParserListener) EnterServerBuildInfo(ctx *ServerBuildInfoContext) {}

// ExitServerBuildInfo is called when production serverBuildInfo is exited.
func (s *BaseMongoShellParserListener) ExitServerBuildInfo(ctx *ServerBuildInfoContext) {}

// EnterDbVersion is called when production dbVersion is entered.
func (s *BaseMongoShellParserListener) EnterDbVersion(ctx *DbVersionContext) {}

// ExitDbVersion is called when production dbVersion is exited.
func (s *BaseMongoShellParserListener) ExitDbVersion(ctx *DbVersionContext) {}

// EnterHostInfo is called when production hostInfo is entered.
func (s *BaseMongoShellParserListener) EnterHostInfo(ctx *HostInfoContext) {}

// ExitHostInfo is called when production hostInfo is exited.
func (s *BaseMongoShellParserListener) ExitHostInfo(ctx *HostInfoContext) {}

// EnterListCommands is called when production listCommands is entered.
func (s *BaseMongoShellParserListener) EnterListCommands(ctx *ListCommandsContext) {}

// ExitListCommands is called when production listCommands is exited.
func (s *BaseMongoShellParserListener) ExitListCommands(ctx *ListCommandsContext) {}

// EnterRunCommand is called when production runCommand is entered.
func (s *BaseMongoShellParserListener) EnterRunCommand(ctx *RunCommandContext) {}

// ExitRunCommand is called when production runCommand is exited.
func (s *BaseMongoShellParserListener) ExitRunCommand(ctx *RunCommandContext) {}

// EnterAdminCommand is called when production adminCommand is entered.
func (s *BaseMongoShellParserListener) EnterAdminCommand(ctx *AdminCommandContext) {}

// ExitAdminCommand is called when production adminCommand is exited.
func (s *BaseMongoShellParserListener) ExitAdminCommand(ctx *AdminCommandContext) {}

// EnterGetName is called when production getName is entered.
func (s *BaseMongoShellParserListener) EnterGetName(ctx *GetNameContext) {}

// ExitGetName is called when production getName is exited.
func (s *BaseMongoShellParserListener) ExitGetName(ctx *GetNameContext) {}

// EnterGetMongo is called when production getMongo is entered.
func (s *BaseMongoShellParserListener) EnterGetMongo(ctx *GetMongoContext) {}

// ExitGetMongo is called when production getMongo is exited.
func (s *BaseMongoShellParserListener) ExitGetMongo(ctx *GetMongoContext) {}

// EnterGetSiblingDB is called when production getSiblingDB is entered.
func (s *BaseMongoShellParserListener) EnterGetSiblingDB(ctx *GetSiblingDBContext) {}

// ExitGetSiblingDB is called when production getSiblingDB is exited.
func (s *BaseMongoShellParserListener) ExitGetSiblingDB(ctx *GetSiblingDBContext) {}

// EnterDbGenericMethod is called when production dbGenericMethod is entered.
func (s *BaseMongoShellParserListener) EnterDbGenericMethod(ctx *DbGenericMethodContext) {}

// ExitDbGenericMethod is called when production dbGenericMethod is exited.
func (s *BaseMongoShellParserListener) ExitDbGenericMethod(ctx *DbGenericMethodContext) {}

// EnterCollectionOperation is called when production collectionOperation is entered.
func (s *BaseMongoShellParserListener) EnterCollectionOperation(ctx *CollectionOperationContext) {}

// ExitCollectionOperation is called when production collectionOperation is exited.
func (s *BaseMongoShellParserListener) ExitCollectionOperation(ctx *CollectionOperationContext) {}

// EnterGenericDbMethod is called when production genericDbMethod is entered.
func (s *BaseMongoShellParserListener) EnterGenericDbMethod(ctx *GenericDbMethodContext) {}

// ExitGenericDbMethod is called when production genericDbMethod is exited.
func (s *BaseMongoShellParserListener) ExitGenericDbMethod(ctx *GenericDbMethodContext) {}

// EnterBulkStatement is called when production bulkStatement is entered.
func (s *BaseMongoShellParserListener) EnterBulkStatement(ctx *BulkStatementContext) {}

// ExitBulkStatement is called when production bulkStatement is exited.
func (s *BaseMongoShellParserListener) ExitBulkStatement(ctx *BulkStatementContext) {}

// EnterBulkInitMethod is called when production bulkInitMethod is entered.
func (s *BaseMongoShellParserListener) EnterBulkInitMethod(ctx *BulkInitMethodContext) {}

// ExitBulkInitMethod is called when production bulkInitMethod is exited.
func (s *BaseMongoShellParserListener) ExitBulkInitMethod(ctx *BulkInitMethodContext) {}

// EnterBulkMethodChain is called when production bulkMethodChain is entered.
func (s *BaseMongoShellParserListener) EnterBulkMethodChain(ctx *BulkMethodChainContext) {}

// ExitBulkMethodChain is called when production bulkMethodChain is exited.
func (s *BaseMongoShellParserListener) ExitBulkMethodChain(ctx *BulkMethodChainContext) {}

// EnterBulkFind is called when production bulkFind is entered.
func (s *BaseMongoShellParserListener) EnterBulkFind(ctx *BulkFindContext) {}

// ExitBulkFind is called when production bulkFind is exited.
func (s *BaseMongoShellParserListener) ExitBulkFind(ctx *BulkFindContext) {}

// EnterBulkInsert is called when production bulkInsert is entered.
func (s *BaseMongoShellParserListener) EnterBulkInsert(ctx *BulkInsertContext) {}

// ExitBulkInsert is called when production bulkInsert is exited.
func (s *BaseMongoShellParserListener) ExitBulkInsert(ctx *BulkInsertContext) {}

// EnterBulkRemove is called when production bulkRemove is entered.
func (s *BaseMongoShellParserListener) EnterBulkRemove(ctx *BulkRemoveContext) {}

// ExitBulkRemove is called when production bulkRemove is exited.
func (s *BaseMongoShellParserListener) ExitBulkRemove(ctx *BulkRemoveContext) {}

// EnterBulkExecute is called when production bulkExecute is entered.
func (s *BaseMongoShellParserListener) EnterBulkExecute(ctx *BulkExecuteContext) {}

// ExitBulkExecute is called when production bulkExecute is exited.
func (s *BaseMongoShellParserListener) ExitBulkExecute(ctx *BulkExecuteContext) {}

// EnterBulkGetOperations is called when production bulkGetOperations is entered.
func (s *BaseMongoShellParserListener) EnterBulkGetOperations(ctx *BulkGetOperationsContext) {}

// ExitBulkGetOperations is called when production bulkGetOperations is exited.
func (s *BaseMongoShellParserListener) ExitBulkGetOperations(ctx *BulkGetOperationsContext) {}

// EnterBulkToString is called when production bulkToString is entered.
func (s *BaseMongoShellParserListener) EnterBulkToString(ctx *BulkToStringContext) {}

// ExitBulkToString is called when production bulkToString is exited.
func (s *BaseMongoShellParserListener) ExitBulkToString(ctx *BulkToStringContext) {}

// EnterBulkGenericMethod is called when production bulkGenericMethod is entered.
func (s *BaseMongoShellParserListener) EnterBulkGenericMethod(ctx *BulkGenericMethodContext) {}

// ExitBulkGenericMethod is called when production bulkGenericMethod is exited.
func (s *BaseMongoShellParserListener) ExitBulkGenericMethod(ctx *BulkGenericMethodContext) {}

// EnterMongoConnection is called when production mongoConnection is entered.
func (s *BaseMongoShellParserListener) EnterMongoConnection(ctx *MongoConnectionContext) {}

// ExitMongoConnection is called when production mongoConnection is exited.
func (s *BaseMongoShellParserListener) ExitMongoConnection(ctx *MongoConnectionContext) {}

// EnterConnectCall is called when production connectCall is entered.
func (s *BaseMongoShellParserListener) EnterConnectCall(ctx *ConnectCallContext) {}

// ExitConnectCall is called when production connectCall is exited.
func (s *BaseMongoShellParserListener) ExitConnectCall(ctx *ConnectCallContext) {}

// EnterDbGetMongoChain is called when production dbGetMongoChain is entered.
func (s *BaseMongoShellParserListener) EnterDbGetMongoChain(ctx *DbGetMongoChainContext) {}

// ExitDbGetMongoChain is called when production dbGetMongoChain is exited.
func (s *BaseMongoShellParserListener) ExitDbGetMongoChain(ctx *DbGetMongoChainContext) {}

// EnterConnectionMethodChain is called when production connectionMethodChain is entered.
func (s *BaseMongoShellParserListener) EnterConnectionMethodChain(ctx *ConnectionMethodChainContext) {
}

// ExitConnectionMethodChain is called when production connectionMethodChain is exited.
func (s *BaseMongoShellParserListener) ExitConnectionMethodChain(ctx *ConnectionMethodChainContext) {}

// EnterRsStatement is called when production rsStatement is entered.
func (s *BaseMongoShellParserListener) EnterRsStatement(ctx *RsStatementContext) {}

// ExitRsStatement is called when production rsStatement is exited.
func (s *BaseMongoShellParserListener) ExitRsStatement(ctx *RsStatementContext) {}

// EnterShStatement is called when production shStatement is entered.
func (s *BaseMongoShellParserListener) EnterShStatement(ctx *ShStatementContext) {}

// ExitShStatement is called when production shStatement is exited.
func (s *BaseMongoShellParserListener) ExitShStatement(ctx *ShStatementContext) {}

// EnterKeyVaultStatement is called when production keyVaultStatement is entered.
func (s *BaseMongoShellParserListener) EnterKeyVaultStatement(ctx *KeyVaultStatementContext) {}

// ExitKeyVaultStatement is called when production keyVaultStatement is exited.
func (s *BaseMongoShellParserListener) ExitKeyVaultStatement(ctx *KeyVaultStatementContext) {}

// EnterClientEncryptionStatement is called when production clientEncryptionStatement is entered.
func (s *BaseMongoShellParserListener) EnterClientEncryptionStatement(ctx *ClientEncryptionStatementContext) {
}

// ExitClientEncryptionStatement is called when production clientEncryptionStatement is exited.
func (s *BaseMongoShellParserListener) ExitClientEncryptionStatement(ctx *ClientEncryptionStatementContext) {
}

// EnterPlanCacheStatement is called when production planCacheStatement is entered.
func (s *BaseMongoShellParserListener) EnterPlanCacheStatement(ctx *PlanCacheStatementContext) {}

// ExitPlanCacheStatement is called when production planCacheStatement is exited.
func (s *BaseMongoShellParserListener) ExitPlanCacheStatement(ctx *PlanCacheStatementContext) {}

// EnterSpStatement is called when production spStatement is entered.
func (s *BaseMongoShellParserListener) EnterSpStatement(ctx *SpStatementContext) {}

// ExitSpStatement is called when production spStatement is exited.
func (s *BaseMongoShellParserListener) ExitSpStatement(ctx *SpStatementContext) {}

// EnterNativeFunctionCall is called when production nativeFunctionCall is entered.
func (s *BaseMongoShellParserListener) EnterNativeFunctionCall(ctx *NativeFunctionCallContext) {}

// ExitNativeFunctionCall is called when production nativeFunctionCall is exited.
func (s *BaseMongoShellParserListener) ExitNativeFunctionCall(ctx *NativeFunctionCallContext) {}

// EnterConnGetDB is called when production connGetDB is entered.
func (s *BaseMongoShellParserListener) EnterConnGetDB(ctx *ConnGetDBContext) {}

// ExitConnGetDB is called when production connGetDB is exited.
func (s *BaseMongoShellParserListener) ExitConnGetDB(ctx *ConnGetDBContext) {}

// EnterConnGetReadConcern is called when production connGetReadConcern is entered.
func (s *BaseMongoShellParserListener) EnterConnGetReadConcern(ctx *ConnGetReadConcernContext) {}

// ExitConnGetReadConcern is called when production connGetReadConcern is exited.
func (s *BaseMongoShellParserListener) ExitConnGetReadConcern(ctx *ConnGetReadConcernContext) {}

// EnterConnGetReadPref is called when production connGetReadPref is entered.
func (s *BaseMongoShellParserListener) EnterConnGetReadPref(ctx *ConnGetReadPrefContext) {}

// ExitConnGetReadPref is called when production connGetReadPref is exited.
func (s *BaseMongoShellParserListener) ExitConnGetReadPref(ctx *ConnGetReadPrefContext) {}

// EnterConnGetReadPrefMode is called when production connGetReadPrefMode is entered.
func (s *BaseMongoShellParserListener) EnterConnGetReadPrefMode(ctx *ConnGetReadPrefModeContext) {}

// ExitConnGetReadPrefMode is called when production connGetReadPrefMode is exited.
func (s *BaseMongoShellParserListener) ExitConnGetReadPrefMode(ctx *ConnGetReadPrefModeContext) {}

// EnterConnGetReadPrefTagSet is called when production connGetReadPrefTagSet is entered.
func (s *BaseMongoShellParserListener) EnterConnGetReadPrefTagSet(ctx *ConnGetReadPrefTagSetContext) {
}

// ExitConnGetReadPrefTagSet is called when production connGetReadPrefTagSet is exited.
func (s *BaseMongoShellParserListener) ExitConnGetReadPrefTagSet(ctx *ConnGetReadPrefTagSetContext) {}

// EnterConnGetWriteConcern is called when production connGetWriteConcern is entered.
func (s *BaseMongoShellParserListener) EnterConnGetWriteConcern(ctx *ConnGetWriteConcernContext) {}

// ExitConnGetWriteConcern is called when production connGetWriteConcern is exited.
func (s *BaseMongoShellParserListener) ExitConnGetWriteConcern(ctx *ConnGetWriteConcernContext) {}

// EnterConnSetReadPref is called when production connSetReadPref is entered.
func (s *BaseMongoShellParserListener) EnterConnSetReadPref(ctx *ConnSetReadPrefContext) {}

// ExitConnSetReadPref is called when production connSetReadPref is exited.
func (s *BaseMongoShellParserListener) ExitConnSetReadPref(ctx *ConnSetReadPrefContext) {}

// EnterConnSetReadConcern is called when production connSetReadConcern is entered.
func (s *BaseMongoShellParserListener) EnterConnSetReadConcern(ctx *ConnSetReadConcernContext) {}

// ExitConnSetReadConcern is called when production connSetReadConcern is exited.
func (s *BaseMongoShellParserListener) ExitConnSetReadConcern(ctx *ConnSetReadConcernContext) {}

// EnterConnSetWriteConcern is called when production connSetWriteConcern is entered.
func (s *BaseMongoShellParserListener) EnterConnSetWriteConcern(ctx *ConnSetWriteConcernContext) {}

// ExitConnSetWriteConcern is called when production connSetWriteConcern is exited.
func (s *BaseMongoShellParserListener) ExitConnSetWriteConcern(ctx *ConnSetWriteConcernContext) {}

// EnterConnStartSession is called when production connStartSession is entered.
func (s *BaseMongoShellParserListener) EnterConnStartSession(ctx *ConnStartSessionContext) {}

// ExitConnStartSession is called when production connStartSession is exited.
func (s *BaseMongoShellParserListener) ExitConnStartSession(ctx *ConnStartSessionContext) {}

// EnterConnWatch is called when production connWatch is entered.
func (s *BaseMongoShellParserListener) EnterConnWatch(ctx *ConnWatchContext) {}

// ExitConnWatch is called when production connWatch is exited.
func (s *BaseMongoShellParserListener) ExitConnWatch(ctx *ConnWatchContext) {}

// EnterConnClose is called when production connClose is entered.
func (s *BaseMongoShellParserListener) EnterConnClose(ctx *ConnCloseContext) {}

// ExitConnClose is called when production connClose is exited.
func (s *BaseMongoShellParserListener) ExitConnClose(ctx *ConnCloseContext) {}

// EnterConnAdminCommand is called when production connAdminCommand is entered.
func (s *BaseMongoShellParserListener) EnterConnAdminCommand(ctx *ConnAdminCommandContext) {}

// ExitConnAdminCommand is called when production connAdminCommand is exited.
func (s *BaseMongoShellParserListener) ExitConnAdminCommand(ctx *ConnAdminCommandContext) {}

// EnterConnGetDBNames is called when production connGetDBNames is entered.
func (s *BaseMongoShellParserListener) EnterConnGetDBNames(ctx *ConnGetDBNamesContext) {}

// ExitConnGetDBNames is called when production connGetDBNames is exited.
func (s *BaseMongoShellParserListener) ExitConnGetDBNames(ctx *ConnGetDBNamesContext) {}

// EnterConnGenericMethod is called when production connGenericMethod is entered.
func (s *BaseMongoShellParserListener) EnterConnGenericMethod(ctx *ConnGenericMethodContext) {}

// ExitConnGenericMethod is called when production connGenericMethod is exited.
func (s *BaseMongoShellParserListener) ExitConnGenericMethod(ctx *ConnGenericMethodContext) {}

// EnterDotAccess is called when production dotAccess is entered.
func (s *BaseMongoShellParserListener) EnterDotAccess(ctx *DotAccessContext) {}

// ExitDotAccess is called when production dotAccess is exited.
func (s *BaseMongoShellParserListener) ExitDotAccess(ctx *DotAccessContext) {}

// EnterBracketAccess is called when production bracketAccess is entered.
func (s *BaseMongoShellParserListener) EnterBracketAccess(ctx *BracketAccessContext) {}

// ExitBracketAccess is called when production bracketAccess is exited.
func (s *BaseMongoShellParserListener) ExitBracketAccess(ctx *BracketAccessContext) {}

// EnterGetCollectionAccess is called when production getCollectionAccess is entered.
func (s *BaseMongoShellParserListener) EnterGetCollectionAccess(ctx *GetCollectionAccessContext) {}

// ExitGetCollectionAccess is called when production getCollectionAccess is exited.
func (s *BaseMongoShellParserListener) ExitGetCollectionAccess(ctx *GetCollectionAccessContext) {}

// EnterMethodChain is called when production methodChain is entered.
func (s *BaseMongoShellParserListener) EnterMethodChain(ctx *MethodChainContext) {}

// ExitMethodChain is called when production methodChain is exited.
func (s *BaseMongoShellParserListener) ExitMethodChain(ctx *MethodChainContext) {}

// EnterCollectionMethodCall is called when production collectionMethodCall is entered.
func (s *BaseMongoShellParserListener) EnterCollectionMethodCall(ctx *CollectionMethodCallContext) {}

// ExitCollectionMethodCall is called when production collectionMethodCall is exited.
func (s *BaseMongoShellParserListener) ExitCollectionMethodCall(ctx *CollectionMethodCallContext) {}

// EnterCursorMethodCall is called when production cursorMethodCall is entered.
func (s *BaseMongoShellParserListener) EnterCursorMethodCall(ctx *CursorMethodCallContext) {}

// ExitCursorMethodCall is called when production cursorMethodCall is exited.
func (s *BaseMongoShellParserListener) ExitCursorMethodCall(ctx *CursorMethodCallContext) {}

// EnterFindMethod is called when production findMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindMethod(ctx *FindMethodContext) {}

// ExitFindMethod is called when production findMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindMethod(ctx *FindMethodContext) {}

// EnterFindOneMethod is called when production findOneMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindOneMethod(ctx *FindOneMethodContext) {}

// ExitFindOneMethod is called when production findOneMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindOneMethod(ctx *FindOneMethodContext) {}

// EnterCountDocumentsMethod is called when production countDocumentsMethod is entered.
func (s *BaseMongoShellParserListener) EnterCountDocumentsMethod(ctx *CountDocumentsMethodContext) {}

// ExitCountDocumentsMethod is called when production countDocumentsMethod is exited.
func (s *BaseMongoShellParserListener) ExitCountDocumentsMethod(ctx *CountDocumentsMethodContext) {}

// EnterEstimatedDocumentCountMethod is called when production estimatedDocumentCountMethod is entered.
func (s *BaseMongoShellParserListener) EnterEstimatedDocumentCountMethod(ctx *EstimatedDocumentCountMethodContext) {
}

// ExitEstimatedDocumentCountMethod is called when production estimatedDocumentCountMethod is exited.
func (s *BaseMongoShellParserListener) ExitEstimatedDocumentCountMethod(ctx *EstimatedDocumentCountMethodContext) {
}

// EnterDistinctMethod is called when production distinctMethod is entered.
func (s *BaseMongoShellParserListener) EnterDistinctMethod(ctx *DistinctMethodContext) {}

// ExitDistinctMethod is called when production distinctMethod is exited.
func (s *BaseMongoShellParserListener) ExitDistinctMethod(ctx *DistinctMethodContext) {}

// EnterAggregateMethod is called when production aggregateMethod is entered.
func (s *BaseMongoShellParserListener) EnterAggregateMethod(ctx *AggregateMethodContext) {}

// ExitAggregateMethod is called when production aggregateMethod is exited.
func (s *BaseMongoShellParserListener) ExitAggregateMethod(ctx *AggregateMethodContext) {}

// EnterGetIndexesMethod is called when production getIndexesMethod is entered.
func (s *BaseMongoShellParserListener) EnterGetIndexesMethod(ctx *GetIndexesMethodContext) {}

// ExitGetIndexesMethod is called when production getIndexesMethod is exited.
func (s *BaseMongoShellParserListener) ExitGetIndexesMethod(ctx *GetIndexesMethodContext) {}

// EnterInsertOneMethod is called when production insertOneMethod is entered.
func (s *BaseMongoShellParserListener) EnterInsertOneMethod(ctx *InsertOneMethodContext) {}

// ExitInsertOneMethod is called when production insertOneMethod is exited.
func (s *BaseMongoShellParserListener) ExitInsertOneMethod(ctx *InsertOneMethodContext) {}

// EnterInsertManyMethod is called when production insertManyMethod is entered.
func (s *BaseMongoShellParserListener) EnterInsertManyMethod(ctx *InsertManyMethodContext) {}

// ExitInsertManyMethod is called when production insertManyMethod is exited.
func (s *BaseMongoShellParserListener) ExitInsertManyMethod(ctx *InsertManyMethodContext) {}

// EnterUpdateOneMethod is called when production updateOneMethod is entered.
func (s *BaseMongoShellParserListener) EnterUpdateOneMethod(ctx *UpdateOneMethodContext) {}

// ExitUpdateOneMethod is called when production updateOneMethod is exited.
func (s *BaseMongoShellParserListener) ExitUpdateOneMethod(ctx *UpdateOneMethodContext) {}

// EnterUpdateManyMethod is called when production updateManyMethod is entered.
func (s *BaseMongoShellParserListener) EnterUpdateManyMethod(ctx *UpdateManyMethodContext) {}

// ExitUpdateManyMethod is called when production updateManyMethod is exited.
func (s *BaseMongoShellParserListener) ExitUpdateManyMethod(ctx *UpdateManyMethodContext) {}

// EnterDeleteOneMethod is called when production deleteOneMethod is entered.
func (s *BaseMongoShellParserListener) EnterDeleteOneMethod(ctx *DeleteOneMethodContext) {}

// ExitDeleteOneMethod is called when production deleteOneMethod is exited.
func (s *BaseMongoShellParserListener) ExitDeleteOneMethod(ctx *DeleteOneMethodContext) {}

// EnterDeleteManyMethod is called when production deleteManyMethod is entered.
func (s *BaseMongoShellParserListener) EnterDeleteManyMethod(ctx *DeleteManyMethodContext) {}

// ExitDeleteManyMethod is called when production deleteManyMethod is exited.
func (s *BaseMongoShellParserListener) ExitDeleteManyMethod(ctx *DeleteManyMethodContext) {}

// EnterReplaceOneMethod is called when production replaceOneMethod is entered.
func (s *BaseMongoShellParserListener) EnterReplaceOneMethod(ctx *ReplaceOneMethodContext) {}

// ExitReplaceOneMethod is called when production replaceOneMethod is exited.
func (s *BaseMongoShellParserListener) ExitReplaceOneMethod(ctx *ReplaceOneMethodContext) {}

// EnterFindOneAndUpdateMethod is called when production findOneAndUpdateMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindOneAndUpdateMethod(ctx *FindOneAndUpdateMethodContext) {
}

// ExitFindOneAndUpdateMethod is called when production findOneAndUpdateMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindOneAndUpdateMethod(ctx *FindOneAndUpdateMethodContext) {
}

// EnterFindOneAndReplaceMethod is called when production findOneAndReplaceMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindOneAndReplaceMethod(ctx *FindOneAndReplaceMethodContext) {
}

// ExitFindOneAndReplaceMethod is called when production findOneAndReplaceMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindOneAndReplaceMethod(ctx *FindOneAndReplaceMethodContext) {
}

// EnterFindOneAndDeleteMethod is called when production findOneAndDeleteMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindOneAndDeleteMethod(ctx *FindOneAndDeleteMethodContext) {
}

// ExitFindOneAndDeleteMethod is called when production findOneAndDeleteMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindOneAndDeleteMethod(ctx *FindOneAndDeleteMethodContext) {
}

// EnterCreateIndexMethod is called when production createIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterCreateIndexMethod(ctx *CreateIndexMethodContext) {}

// ExitCreateIndexMethod is called when production createIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitCreateIndexMethod(ctx *CreateIndexMethodContext) {}

// EnterCreateIndexesMethod is called when production createIndexesMethod is entered.
func (s *BaseMongoShellParserListener) EnterCreateIndexesMethod(ctx *CreateIndexesMethodContext) {}

// ExitCreateIndexesMethod is called when production createIndexesMethod is exited.
func (s *BaseMongoShellParserListener) ExitCreateIndexesMethod(ctx *CreateIndexesMethodContext) {}

// EnterDropIndexMethod is called when production dropIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterDropIndexMethod(ctx *DropIndexMethodContext) {}

// ExitDropIndexMethod is called when production dropIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitDropIndexMethod(ctx *DropIndexMethodContext) {}

// EnterDropIndexesMethod is called when production dropIndexesMethod is entered.
func (s *BaseMongoShellParserListener) EnterDropIndexesMethod(ctx *DropIndexesMethodContext) {}

// ExitDropIndexesMethod is called when production dropIndexesMethod is exited.
func (s *BaseMongoShellParserListener) ExitDropIndexesMethod(ctx *DropIndexesMethodContext) {}

// EnterDropMethod is called when production dropMethod is entered.
func (s *BaseMongoShellParserListener) EnterDropMethod(ctx *DropMethodContext) {}

// ExitDropMethod is called when production dropMethod is exited.
func (s *BaseMongoShellParserListener) ExitDropMethod(ctx *DropMethodContext) {}

// EnterRenameCollectionMethod is called when production renameCollectionMethod is entered.
func (s *BaseMongoShellParserListener) EnterRenameCollectionMethod(ctx *RenameCollectionMethodContext) {
}

// ExitRenameCollectionMethod is called when production renameCollectionMethod is exited.
func (s *BaseMongoShellParserListener) ExitRenameCollectionMethod(ctx *RenameCollectionMethodContext) {
}

// EnterStatsMethod is called when production statsMethod is entered.
func (s *BaseMongoShellParserListener) EnterStatsMethod(ctx *StatsMethodContext) {}

// ExitStatsMethod is called when production statsMethod is exited.
func (s *BaseMongoShellParserListener) ExitStatsMethod(ctx *StatsMethodContext) {}

// EnterStorageSizeMethod is called when production storageSizeMethod is entered.
func (s *BaseMongoShellParserListener) EnterStorageSizeMethod(ctx *StorageSizeMethodContext) {}

// ExitStorageSizeMethod is called when production storageSizeMethod is exited.
func (s *BaseMongoShellParserListener) ExitStorageSizeMethod(ctx *StorageSizeMethodContext) {}

// EnterTotalIndexSizeMethod is called when production totalIndexSizeMethod is entered.
func (s *BaseMongoShellParserListener) EnterTotalIndexSizeMethod(ctx *TotalIndexSizeMethodContext) {}

// ExitTotalIndexSizeMethod is called when production totalIndexSizeMethod is exited.
func (s *BaseMongoShellParserListener) ExitTotalIndexSizeMethod(ctx *TotalIndexSizeMethodContext) {}

// EnterTotalSizeMethod is called when production totalSizeMethod is entered.
func (s *BaseMongoShellParserListener) EnterTotalSizeMethod(ctx *TotalSizeMethodContext) {}

// ExitTotalSizeMethod is called when production totalSizeMethod is exited.
func (s *BaseMongoShellParserListener) ExitTotalSizeMethod(ctx *TotalSizeMethodContext) {}

// EnterDataSizeMethod is called when production dataSizeMethod is entered.
func (s *BaseMongoShellParserListener) EnterDataSizeMethod(ctx *DataSizeMethodContext) {}

// ExitDataSizeMethod is called when production dataSizeMethod is exited.
func (s *BaseMongoShellParserListener) ExitDataSizeMethod(ctx *DataSizeMethodContext) {}

// EnterIsCappedMethod is called when production isCappedMethod is entered.
func (s *BaseMongoShellParserListener) EnterIsCappedMethod(ctx *IsCappedMethodContext) {}

// ExitIsCappedMethod is called when production isCappedMethod is exited.
func (s *BaseMongoShellParserListener) ExitIsCappedMethod(ctx *IsCappedMethodContext) {}

// EnterValidateMethod is called when production validateMethod is entered.
func (s *BaseMongoShellParserListener) EnterValidateMethod(ctx *ValidateMethodContext) {}

// ExitValidateMethod is called when production validateMethod is exited.
func (s *BaseMongoShellParserListener) ExitValidateMethod(ctx *ValidateMethodContext) {}

// EnterLatencyStatsMethod is called when production latencyStatsMethod is entered.
func (s *BaseMongoShellParserListener) EnterLatencyStatsMethod(ctx *LatencyStatsMethodContext) {}

// ExitLatencyStatsMethod is called when production latencyStatsMethod is exited.
func (s *BaseMongoShellParserListener) ExitLatencyStatsMethod(ctx *LatencyStatsMethodContext) {}

// EnterWatchMethod is called when production watchMethod is entered.
func (s *BaseMongoShellParserListener) EnterWatchMethod(ctx *WatchMethodContext) {}

// ExitWatchMethod is called when production watchMethod is exited.
func (s *BaseMongoShellParserListener) ExitWatchMethod(ctx *WatchMethodContext) {}

// EnterBulkWriteMethod is called when production bulkWriteMethod is entered.
func (s *BaseMongoShellParserListener) EnterBulkWriteMethod(ctx *BulkWriteMethodContext) {}

// ExitBulkWriteMethod is called when production bulkWriteMethod is exited.
func (s *BaseMongoShellParserListener) ExitBulkWriteMethod(ctx *BulkWriteMethodContext) {}

// EnterCollectionCountMethod is called when production collectionCountMethod is entered.
func (s *BaseMongoShellParserListener) EnterCollectionCountMethod(ctx *CollectionCountMethodContext) {
}

// ExitCollectionCountMethod is called when production collectionCountMethod is exited.
func (s *BaseMongoShellParserListener) ExitCollectionCountMethod(ctx *CollectionCountMethodContext) {}

// EnterCollectionInsertMethod is called when production collectionInsertMethod is entered.
func (s *BaseMongoShellParserListener) EnterCollectionInsertMethod(ctx *CollectionInsertMethodContext) {
}

// ExitCollectionInsertMethod is called when production collectionInsertMethod is exited.
func (s *BaseMongoShellParserListener) ExitCollectionInsertMethod(ctx *CollectionInsertMethodContext) {
}

// EnterCollectionRemoveMethod is called when production collectionRemoveMethod is entered.
func (s *BaseMongoShellParserListener) EnterCollectionRemoveMethod(ctx *CollectionRemoveMethodContext) {
}

// ExitCollectionRemoveMethod is called when production collectionRemoveMethod is exited.
func (s *BaseMongoShellParserListener) ExitCollectionRemoveMethod(ctx *CollectionRemoveMethodContext) {
}

// EnterUpdateMethod is called when production updateMethod is entered.
func (s *BaseMongoShellParserListener) EnterUpdateMethod(ctx *UpdateMethodContext) {}

// ExitUpdateMethod is called when production updateMethod is exited.
func (s *BaseMongoShellParserListener) ExitUpdateMethod(ctx *UpdateMethodContext) {}

// EnterMapReduceMethod is called when production mapReduceMethod is entered.
func (s *BaseMongoShellParserListener) EnterMapReduceMethod(ctx *MapReduceMethodContext) {}

// ExitMapReduceMethod is called when production mapReduceMethod is exited.
func (s *BaseMongoShellParserListener) ExitMapReduceMethod(ctx *MapReduceMethodContext) {}

// EnterFindAndModifyMethod is called when production findAndModifyMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindAndModifyMethod(ctx *FindAndModifyMethodContext) {}

// ExitFindAndModifyMethod is called when production findAndModifyMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindAndModifyMethod(ctx *FindAndModifyMethodContext) {}

// EnterCollectionExplainMethod is called when production collectionExplainMethod is entered.
func (s *BaseMongoShellParserListener) EnterCollectionExplainMethod(ctx *CollectionExplainMethodContext) {
}

// ExitCollectionExplainMethod is called when production collectionExplainMethod is exited.
func (s *BaseMongoShellParserListener) ExitCollectionExplainMethod(ctx *CollectionExplainMethodContext) {
}

// EnterAnalyzeShardKeyMethod is called when production analyzeShardKeyMethod is entered.
func (s *BaseMongoShellParserListener) EnterAnalyzeShardKeyMethod(ctx *AnalyzeShardKeyMethodContext) {
}

// ExitAnalyzeShardKeyMethod is called when production analyzeShardKeyMethod is exited.
func (s *BaseMongoShellParserListener) ExitAnalyzeShardKeyMethod(ctx *AnalyzeShardKeyMethodContext) {}

// EnterConfigureQueryAnalyzerMethod is called when production configureQueryAnalyzerMethod is entered.
func (s *BaseMongoShellParserListener) EnterConfigureQueryAnalyzerMethod(ctx *ConfigureQueryAnalyzerMethodContext) {
}

// ExitConfigureQueryAnalyzerMethod is called when production configureQueryAnalyzerMethod is exited.
func (s *BaseMongoShellParserListener) ExitConfigureQueryAnalyzerMethod(ctx *ConfigureQueryAnalyzerMethodContext) {
}

// EnterCompactStructuredEncryptionDataMethod is called when production compactStructuredEncryptionDataMethod is entered.
func (s *BaseMongoShellParserListener) EnterCompactStructuredEncryptionDataMethod(ctx *CompactStructuredEncryptionDataMethodContext) {
}

// ExitCompactStructuredEncryptionDataMethod is called when production compactStructuredEncryptionDataMethod is exited.
func (s *BaseMongoShellParserListener) ExitCompactStructuredEncryptionDataMethod(ctx *CompactStructuredEncryptionDataMethodContext) {
}

// EnterHideIndexMethod is called when production hideIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterHideIndexMethod(ctx *HideIndexMethodContext) {}

// ExitHideIndexMethod is called when production hideIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitHideIndexMethod(ctx *HideIndexMethodContext) {}

// EnterUnhideIndexMethod is called when production unhideIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterUnhideIndexMethod(ctx *UnhideIndexMethodContext) {}

// ExitUnhideIndexMethod is called when production unhideIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitUnhideIndexMethod(ctx *UnhideIndexMethodContext) {}

// EnterReIndexMethod is called when production reIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterReIndexMethod(ctx *ReIndexMethodContext) {}

// ExitReIndexMethod is called when production reIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitReIndexMethod(ctx *ReIndexMethodContext) {}

// EnterGetShardDistributionMethod is called when production getShardDistributionMethod is entered.
func (s *BaseMongoShellParserListener) EnterGetShardDistributionMethod(ctx *GetShardDistributionMethodContext) {
}

// ExitGetShardDistributionMethod is called when production getShardDistributionMethod is exited.
func (s *BaseMongoShellParserListener) ExitGetShardDistributionMethod(ctx *GetShardDistributionMethodContext) {
}

// EnterGetShardVersionMethod is called when production getShardVersionMethod is entered.
func (s *BaseMongoShellParserListener) EnterGetShardVersionMethod(ctx *GetShardVersionMethodContext) {
}

// ExitGetShardVersionMethod is called when production getShardVersionMethod is exited.
func (s *BaseMongoShellParserListener) ExitGetShardVersionMethod(ctx *GetShardVersionMethodContext) {}

// EnterCreateSearchIndexMethod is called when production createSearchIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterCreateSearchIndexMethod(ctx *CreateSearchIndexMethodContext) {
}

// ExitCreateSearchIndexMethod is called when production createSearchIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitCreateSearchIndexMethod(ctx *CreateSearchIndexMethodContext) {
}

// EnterCreateSearchIndexesMethod is called when production createSearchIndexesMethod is entered.
func (s *BaseMongoShellParserListener) EnterCreateSearchIndexesMethod(ctx *CreateSearchIndexesMethodContext) {
}

// ExitCreateSearchIndexesMethod is called when production createSearchIndexesMethod is exited.
func (s *BaseMongoShellParserListener) ExitCreateSearchIndexesMethod(ctx *CreateSearchIndexesMethodContext) {
}

// EnterDropSearchIndexMethod is called when production dropSearchIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterDropSearchIndexMethod(ctx *DropSearchIndexMethodContext) {
}

// ExitDropSearchIndexMethod is called when production dropSearchIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitDropSearchIndexMethod(ctx *DropSearchIndexMethodContext) {}

// EnterUpdateSearchIndexMethod is called when production updateSearchIndexMethod is entered.
func (s *BaseMongoShellParserListener) EnterUpdateSearchIndexMethod(ctx *UpdateSearchIndexMethodContext) {
}

// ExitUpdateSearchIndexMethod is called when production updateSearchIndexMethod is exited.
func (s *BaseMongoShellParserListener) ExitUpdateSearchIndexMethod(ctx *UpdateSearchIndexMethodContext) {
}

// EnterSortMethod is called when production sortMethod is entered.
func (s *BaseMongoShellParserListener) EnterSortMethod(ctx *SortMethodContext) {}

// ExitSortMethod is called when production sortMethod is exited.
func (s *BaseMongoShellParserListener) ExitSortMethod(ctx *SortMethodContext) {}

// EnterLimitMethod is called when production limitMethod is entered.
func (s *BaseMongoShellParserListener) EnterLimitMethod(ctx *LimitMethodContext) {}

// ExitLimitMethod is called when production limitMethod is exited.
func (s *BaseMongoShellParserListener) ExitLimitMethod(ctx *LimitMethodContext) {}

// EnterSkipMethod is called when production skipMethod is entered.
func (s *BaseMongoShellParserListener) EnterSkipMethod(ctx *SkipMethodContext) {}

// ExitSkipMethod is called when production skipMethod is exited.
func (s *BaseMongoShellParserListener) ExitSkipMethod(ctx *SkipMethodContext) {}

// EnterCountMethod is called when production countMethod is entered.
func (s *BaseMongoShellParserListener) EnterCountMethod(ctx *CountMethodContext) {}

// ExitCountMethod is called when production countMethod is exited.
func (s *BaseMongoShellParserListener) ExitCountMethod(ctx *CountMethodContext) {}

// EnterProjectionMethod is called when production projectionMethod is entered.
func (s *BaseMongoShellParserListener) EnterProjectionMethod(ctx *ProjectionMethodContext) {}

// ExitProjectionMethod is called when production projectionMethod is exited.
func (s *BaseMongoShellParserListener) ExitProjectionMethod(ctx *ProjectionMethodContext) {}

// EnterBatchSizeMethod is called when production batchSizeMethod is entered.
func (s *BaseMongoShellParserListener) EnterBatchSizeMethod(ctx *BatchSizeMethodContext) {}

// ExitBatchSizeMethod is called when production batchSizeMethod is exited.
func (s *BaseMongoShellParserListener) ExitBatchSizeMethod(ctx *BatchSizeMethodContext) {}

// EnterCloseMethod is called when production closeMethod is entered.
func (s *BaseMongoShellParserListener) EnterCloseMethod(ctx *CloseMethodContext) {}

// ExitCloseMethod is called when production closeMethod is exited.
func (s *BaseMongoShellParserListener) ExitCloseMethod(ctx *CloseMethodContext) {}

// EnterCollationMethod is called when production collationMethod is entered.
func (s *BaseMongoShellParserListener) EnterCollationMethod(ctx *CollationMethodContext) {}

// ExitCollationMethod is called when production collationMethod is exited.
func (s *BaseMongoShellParserListener) ExitCollationMethod(ctx *CollationMethodContext) {}

// EnterCommentMethod is called when production commentMethod is entered.
func (s *BaseMongoShellParserListener) EnterCommentMethod(ctx *CommentMethodContext) {}

// ExitCommentMethod is called when production commentMethod is exited.
func (s *BaseMongoShellParserListener) ExitCommentMethod(ctx *CommentMethodContext) {}

// EnterExplainMethod is called when production explainMethod is entered.
func (s *BaseMongoShellParserListener) EnterExplainMethod(ctx *ExplainMethodContext) {}

// ExitExplainMethod is called when production explainMethod is exited.
func (s *BaseMongoShellParserListener) ExitExplainMethod(ctx *ExplainMethodContext) {}

// EnterForEachMethod is called when production forEachMethod is entered.
func (s *BaseMongoShellParserListener) EnterForEachMethod(ctx *ForEachMethodContext) {}

// ExitForEachMethod is called when production forEachMethod is exited.
func (s *BaseMongoShellParserListener) ExitForEachMethod(ctx *ForEachMethodContext) {}

// EnterHasNextMethod is called when production hasNextMethod is entered.
func (s *BaseMongoShellParserListener) EnterHasNextMethod(ctx *HasNextMethodContext) {}

// ExitHasNextMethod is called when production hasNextMethod is exited.
func (s *BaseMongoShellParserListener) ExitHasNextMethod(ctx *HasNextMethodContext) {}

// EnterHintMethod is called when production hintMethod is entered.
func (s *BaseMongoShellParserListener) EnterHintMethod(ctx *HintMethodContext) {}

// ExitHintMethod is called when production hintMethod is exited.
func (s *BaseMongoShellParserListener) ExitHintMethod(ctx *HintMethodContext) {}

// EnterIsClosedMethod is called when production isClosedMethod is entered.
func (s *BaseMongoShellParserListener) EnterIsClosedMethod(ctx *IsClosedMethodContext) {}

// ExitIsClosedMethod is called when production isClosedMethod is exited.
func (s *BaseMongoShellParserListener) ExitIsClosedMethod(ctx *IsClosedMethodContext) {}

// EnterIsExhaustedMethod is called when production isExhaustedMethod is entered.
func (s *BaseMongoShellParserListener) EnterIsExhaustedMethod(ctx *IsExhaustedMethodContext) {}

// ExitIsExhaustedMethod is called when production isExhaustedMethod is exited.
func (s *BaseMongoShellParserListener) ExitIsExhaustedMethod(ctx *IsExhaustedMethodContext) {}

// EnterItcountMethod is called when production itcountMethod is entered.
func (s *BaseMongoShellParserListener) EnterItcountMethod(ctx *ItcountMethodContext) {}

// ExitItcountMethod is called when production itcountMethod is exited.
func (s *BaseMongoShellParserListener) ExitItcountMethod(ctx *ItcountMethodContext) {}

// EnterMapMethod is called when production mapMethod is entered.
func (s *BaseMongoShellParserListener) EnterMapMethod(ctx *MapMethodContext) {}

// ExitMapMethod is called when production mapMethod is exited.
func (s *BaseMongoShellParserListener) ExitMapMethod(ctx *MapMethodContext) {}

// EnterMaxMethod is called when production maxMethod is entered.
func (s *BaseMongoShellParserListener) EnterMaxMethod(ctx *MaxMethodContext) {}

// ExitMaxMethod is called when production maxMethod is exited.
func (s *BaseMongoShellParserListener) ExitMaxMethod(ctx *MaxMethodContext) {}

// EnterMaxAwaitTimeMSMethod is called when production maxAwaitTimeMSMethod is entered.
func (s *BaseMongoShellParserListener) EnterMaxAwaitTimeMSMethod(ctx *MaxAwaitTimeMSMethodContext) {}

// ExitMaxAwaitTimeMSMethod is called when production maxAwaitTimeMSMethod is exited.
func (s *BaseMongoShellParserListener) ExitMaxAwaitTimeMSMethod(ctx *MaxAwaitTimeMSMethodContext) {}

// EnterMaxTimeMSMethod is called when production maxTimeMSMethod is entered.
func (s *BaseMongoShellParserListener) EnterMaxTimeMSMethod(ctx *MaxTimeMSMethodContext) {}

// ExitMaxTimeMSMethod is called when production maxTimeMSMethod is exited.
func (s *BaseMongoShellParserListener) ExitMaxTimeMSMethod(ctx *MaxTimeMSMethodContext) {}

// EnterMinMethod is called when production minMethod is entered.
func (s *BaseMongoShellParserListener) EnterMinMethod(ctx *MinMethodContext) {}

// ExitMinMethod is called when production minMethod is exited.
func (s *BaseMongoShellParserListener) ExitMinMethod(ctx *MinMethodContext) {}

// EnterNextMethod is called when production nextMethod is entered.
func (s *BaseMongoShellParserListener) EnterNextMethod(ctx *NextMethodContext) {}

// ExitNextMethod is called when production nextMethod is exited.
func (s *BaseMongoShellParserListener) ExitNextMethod(ctx *NextMethodContext) {}

// EnterNoCursorTimeoutMethod is called when production noCursorTimeoutMethod is entered.
func (s *BaseMongoShellParserListener) EnterNoCursorTimeoutMethod(ctx *NoCursorTimeoutMethodContext) {
}

// ExitNoCursorTimeoutMethod is called when production noCursorTimeoutMethod is exited.
func (s *BaseMongoShellParserListener) ExitNoCursorTimeoutMethod(ctx *NoCursorTimeoutMethodContext) {}

// EnterObjsLeftInBatchMethod is called when production objsLeftInBatchMethod is entered.
func (s *BaseMongoShellParserListener) EnterObjsLeftInBatchMethod(ctx *ObjsLeftInBatchMethodContext) {
}

// ExitObjsLeftInBatchMethod is called when production objsLeftInBatchMethod is exited.
func (s *BaseMongoShellParserListener) ExitObjsLeftInBatchMethod(ctx *ObjsLeftInBatchMethodContext) {}

// EnterPrettyMethod is called when production prettyMethod is entered.
func (s *BaseMongoShellParserListener) EnterPrettyMethod(ctx *PrettyMethodContext) {}

// ExitPrettyMethod is called when production prettyMethod is exited.
func (s *BaseMongoShellParserListener) ExitPrettyMethod(ctx *PrettyMethodContext) {}

// EnterReadConcernMethod is called when production readConcernMethod is entered.
func (s *BaseMongoShellParserListener) EnterReadConcernMethod(ctx *ReadConcernMethodContext) {}

// ExitReadConcernMethod is called when production readConcernMethod is exited.
func (s *BaseMongoShellParserListener) ExitReadConcernMethod(ctx *ReadConcernMethodContext) {}

// EnterReadPrefMethod is called when production readPrefMethod is entered.
func (s *BaseMongoShellParserListener) EnterReadPrefMethod(ctx *ReadPrefMethodContext) {}

// ExitReadPrefMethod is called when production readPrefMethod is exited.
func (s *BaseMongoShellParserListener) ExitReadPrefMethod(ctx *ReadPrefMethodContext) {}

// EnterReturnKeyMethod is called when production returnKeyMethod is entered.
func (s *BaseMongoShellParserListener) EnterReturnKeyMethod(ctx *ReturnKeyMethodContext) {}

// ExitReturnKeyMethod is called when production returnKeyMethod is exited.
func (s *BaseMongoShellParserListener) ExitReturnKeyMethod(ctx *ReturnKeyMethodContext) {}

// EnterShowRecordIdMethod is called when production showRecordIdMethod is entered.
func (s *BaseMongoShellParserListener) EnterShowRecordIdMethod(ctx *ShowRecordIdMethodContext) {}

// ExitShowRecordIdMethod is called when production showRecordIdMethod is exited.
func (s *BaseMongoShellParserListener) ExitShowRecordIdMethod(ctx *ShowRecordIdMethodContext) {}

// EnterSizeMethod is called when production sizeMethod is entered.
func (s *BaseMongoShellParserListener) EnterSizeMethod(ctx *SizeMethodContext) {}

// ExitSizeMethod is called when production sizeMethod is exited.
func (s *BaseMongoShellParserListener) ExitSizeMethod(ctx *SizeMethodContext) {}

// EnterTailableMethod is called when production tailableMethod is entered.
func (s *BaseMongoShellParserListener) EnterTailableMethod(ctx *TailableMethodContext) {}

// ExitTailableMethod is called when production tailableMethod is exited.
func (s *BaseMongoShellParserListener) ExitTailableMethod(ctx *TailableMethodContext) {}

// EnterToArrayMethod is called when production toArrayMethod is entered.
func (s *BaseMongoShellParserListener) EnterToArrayMethod(ctx *ToArrayMethodContext) {}

// ExitToArrayMethod is called when production toArrayMethod is exited.
func (s *BaseMongoShellParserListener) ExitToArrayMethod(ctx *ToArrayMethodContext) {}

// EnterTryNextMethod is called when production tryNextMethod is entered.
func (s *BaseMongoShellParserListener) EnterTryNextMethod(ctx *TryNextMethodContext) {}

// ExitTryNextMethod is called when production tryNextMethod is exited.
func (s *BaseMongoShellParserListener) ExitTryNextMethod(ctx *TryNextMethodContext) {}

// EnterAllowDiskUseMethod is called when production allowDiskUseMethod is entered.
func (s *BaseMongoShellParserListener) EnterAllowDiskUseMethod(ctx *AllowDiskUseMethodContext) {}

// ExitAllowDiskUseMethod is called when production allowDiskUseMethod is exited.
func (s *BaseMongoShellParserListener) ExitAllowDiskUseMethod(ctx *AllowDiskUseMethodContext) {}

// EnterAddOptionMethod is called when production addOptionMethod is entered.
func (s *BaseMongoShellParserListener) EnterAddOptionMethod(ctx *AddOptionMethodContext) {}

// ExitAddOptionMethod is called when production addOptionMethod is exited.
func (s *BaseMongoShellParserListener) ExitAddOptionMethod(ctx *AddOptionMethodContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseMongoShellParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseMongoShellParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseMongoShellParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseMongoShellParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseMongoShellParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseMongoShellParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseMongoShellParserListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseMongoShellParserListener) ExitPair(ctx *PairContext) {}

// EnterUnquotedKey is called when production unquotedKey is entered.
func (s *BaseMongoShellParserListener) EnterUnquotedKey(ctx *UnquotedKeyContext) {}

// ExitUnquotedKey is called when production unquotedKey is exited.
func (s *BaseMongoShellParserListener) ExitUnquotedKey(ctx *UnquotedKeyContext) {}

// EnterQuotedKey is called when production quotedKey is entered.
func (s *BaseMongoShellParserListener) EnterQuotedKey(ctx *QuotedKeyContext) {}

// ExitQuotedKey is called when production quotedKey is exited.
func (s *BaseMongoShellParserListener) ExitQuotedKey(ctx *QuotedKeyContext) {}

// EnterDocumentValue is called when production documentValue is entered.
func (s *BaseMongoShellParserListener) EnterDocumentValue(ctx *DocumentValueContext) {}

// ExitDocumentValue is called when production documentValue is exited.
func (s *BaseMongoShellParserListener) ExitDocumentValue(ctx *DocumentValueContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseMongoShellParserListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseMongoShellParserListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterHelperValue is called when production helperValue is entered.
func (s *BaseMongoShellParserListener) EnterHelperValue(ctx *HelperValueContext) {}

// ExitHelperValue is called when production helperValue is exited.
func (s *BaseMongoShellParserListener) ExitHelperValue(ctx *HelperValueContext) {}

// EnterRegexLiteralValue is called when production regexLiteralValue is entered.
func (s *BaseMongoShellParserListener) EnterRegexLiteralValue(ctx *RegexLiteralValueContext) {}

// ExitRegexLiteralValue is called when production regexLiteralValue is exited.
func (s *BaseMongoShellParserListener) ExitRegexLiteralValue(ctx *RegexLiteralValueContext) {}

// EnterRegexpConstructorValue is called when production regexpConstructorValue is entered.
func (s *BaseMongoShellParserListener) EnterRegexpConstructorValue(ctx *RegexpConstructorValueContext) {
}

// ExitRegexpConstructorValue is called when production regexpConstructorValue is exited.
func (s *BaseMongoShellParserListener) ExitRegexpConstructorValue(ctx *RegexpConstructorValueContext) {
}

// EnterLiteralValue is called when production literalValue is entered.
func (s *BaseMongoShellParserListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *BaseMongoShellParserListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterNewKeywordValue is called when production newKeywordValue is entered.
func (s *BaseMongoShellParserListener) EnterNewKeywordValue(ctx *NewKeywordValueContext) {}

// ExitNewKeywordValue is called when production newKeywordValue is exited.
func (s *BaseMongoShellParserListener) ExitNewKeywordValue(ctx *NewKeywordValueContext) {}

// EnterNewKeywordError is called when production newKeywordError is entered.
func (s *BaseMongoShellParserListener) EnterNewKeywordError(ctx *NewKeywordErrorContext) {}

// ExitNewKeywordError is called when production newKeywordError is exited.
func (s *BaseMongoShellParserListener) ExitNewKeywordError(ctx *NewKeywordErrorContext) {}

// EnterArray is called when production array is entered.
func (s *BaseMongoShellParserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseMongoShellParserListener) ExitArray(ctx *ArrayContext) {}

// EnterHelperFunction is called when production helperFunction is entered.
func (s *BaseMongoShellParserListener) EnterHelperFunction(ctx *HelperFunctionContext) {}

// ExitHelperFunction is called when production helperFunction is exited.
func (s *BaseMongoShellParserListener) ExitHelperFunction(ctx *HelperFunctionContext) {}

// EnterObjectIdHelper is called when production objectIdHelper is entered.
func (s *BaseMongoShellParserListener) EnterObjectIdHelper(ctx *ObjectIdHelperContext) {}

// ExitObjectIdHelper is called when production objectIdHelper is exited.
func (s *BaseMongoShellParserListener) ExitObjectIdHelper(ctx *ObjectIdHelperContext) {}

// EnterIsoDateHelper is called when production isoDateHelper is entered.
func (s *BaseMongoShellParserListener) EnterIsoDateHelper(ctx *IsoDateHelperContext) {}

// ExitIsoDateHelper is called when production isoDateHelper is exited.
func (s *BaseMongoShellParserListener) ExitIsoDateHelper(ctx *IsoDateHelperContext) {}

// EnterDateHelper is called when production dateHelper is entered.
func (s *BaseMongoShellParserListener) EnterDateHelper(ctx *DateHelperContext) {}

// ExitDateHelper is called when production dateHelper is exited.
func (s *BaseMongoShellParserListener) ExitDateHelper(ctx *DateHelperContext) {}

// EnterUuidHelper is called when production uuidHelper is entered.
func (s *BaseMongoShellParserListener) EnterUuidHelper(ctx *UuidHelperContext) {}

// ExitUuidHelper is called when production uuidHelper is exited.
func (s *BaseMongoShellParserListener) ExitUuidHelper(ctx *UuidHelperContext) {}

// EnterLongHelper is called when production longHelper is entered.
func (s *BaseMongoShellParserListener) EnterLongHelper(ctx *LongHelperContext) {}

// ExitLongHelper is called when production longHelper is exited.
func (s *BaseMongoShellParserListener) ExitLongHelper(ctx *LongHelperContext) {}

// EnterInt32Helper is called when production int32Helper is entered.
func (s *BaseMongoShellParserListener) EnterInt32Helper(ctx *Int32HelperContext) {}

// ExitInt32Helper is called when production int32Helper is exited.
func (s *BaseMongoShellParserListener) ExitInt32Helper(ctx *Int32HelperContext) {}

// EnterDoubleHelper is called when production doubleHelper is entered.
func (s *BaseMongoShellParserListener) EnterDoubleHelper(ctx *DoubleHelperContext) {}

// ExitDoubleHelper is called when production doubleHelper is exited.
func (s *BaseMongoShellParserListener) ExitDoubleHelper(ctx *DoubleHelperContext) {}

// EnterDecimal128Helper is called when production decimal128Helper is entered.
func (s *BaseMongoShellParserListener) EnterDecimal128Helper(ctx *Decimal128HelperContext) {}

// ExitDecimal128Helper is called when production decimal128Helper is exited.
func (s *BaseMongoShellParserListener) ExitDecimal128Helper(ctx *Decimal128HelperContext) {}

// EnterTimestampDocHelper is called when production timestampDocHelper is entered.
func (s *BaseMongoShellParserListener) EnterTimestampDocHelper(ctx *TimestampDocHelperContext) {}

// ExitTimestampDocHelper is called when production timestampDocHelper is exited.
func (s *BaseMongoShellParserListener) ExitTimestampDocHelper(ctx *TimestampDocHelperContext) {}

// EnterTimestampArgsHelper is called when production timestampArgsHelper is entered.
func (s *BaseMongoShellParserListener) EnterTimestampArgsHelper(ctx *TimestampArgsHelperContext) {}

// ExitTimestampArgsHelper is called when production timestampArgsHelper is exited.
func (s *BaseMongoShellParserListener) ExitTimestampArgsHelper(ctx *TimestampArgsHelperContext) {}

// EnterRegExpConstructor is called when production regExpConstructor is entered.
func (s *BaseMongoShellParserListener) EnterRegExpConstructor(ctx *RegExpConstructorContext) {}

// ExitRegExpConstructor is called when production regExpConstructor is exited.
func (s *BaseMongoShellParserListener) ExitRegExpConstructor(ctx *RegExpConstructorContext) {}

// EnterBinDataHelper is called when production binDataHelper is entered.
func (s *BaseMongoShellParserListener) EnterBinDataHelper(ctx *BinDataHelperContext) {}

// ExitBinDataHelper is called when production binDataHelper is exited.
func (s *BaseMongoShellParserListener) ExitBinDataHelper(ctx *BinDataHelperContext) {}

// EnterBinaryHelper is called when production binaryHelper is entered.
func (s *BaseMongoShellParserListener) EnterBinaryHelper(ctx *BinaryHelperContext) {}

// ExitBinaryHelper is called when production binaryHelper is exited.
func (s *BaseMongoShellParserListener) ExitBinaryHelper(ctx *BinaryHelperContext) {}

// EnterBsonRegExpHelper is called when production bsonRegExpHelper is entered.
func (s *BaseMongoShellParserListener) EnterBsonRegExpHelper(ctx *BsonRegExpHelperContext) {}

// ExitBsonRegExpHelper is called when production bsonRegExpHelper is exited.
func (s *BaseMongoShellParserListener) ExitBsonRegExpHelper(ctx *BsonRegExpHelperContext) {}

// EnterHexDataHelper is called when production hexDataHelper is entered.
func (s *BaseMongoShellParserListener) EnterHexDataHelper(ctx *HexDataHelperContext) {}

// ExitHexDataHelper is called when production hexDataHelper is exited.
func (s *BaseMongoShellParserListener) ExitHexDataHelper(ctx *HexDataHelperContext) {}

// EnterStringLiteralValue is called when production stringLiteralValue is entered.
func (s *BaseMongoShellParserListener) EnterStringLiteralValue(ctx *StringLiteralValueContext) {}

// ExitStringLiteralValue is called when production stringLiteralValue is exited.
func (s *BaseMongoShellParserListener) ExitStringLiteralValue(ctx *StringLiteralValueContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseMongoShellParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseMongoShellParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterTrueLiteral is called when production trueLiteral is entered.
func (s *BaseMongoShellParserListener) EnterTrueLiteral(ctx *TrueLiteralContext) {}

// ExitTrueLiteral is called when production trueLiteral is exited.
func (s *BaseMongoShellParserListener) ExitTrueLiteral(ctx *TrueLiteralContext) {}

// EnterFalseLiteral is called when production falseLiteral is entered.
func (s *BaseMongoShellParserListener) EnterFalseLiteral(ctx *FalseLiteralContext) {}

// ExitFalseLiteral is called when production falseLiteral is exited.
func (s *BaseMongoShellParserListener) ExitFalseLiteral(ctx *FalseLiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseMongoShellParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseMongoShellParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseMongoShellParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseMongoShellParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMongoShellParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMongoShellParserListener) ExitIdentifier(ctx *IdentifierContext) {}
