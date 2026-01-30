// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

// MongoShellParserListener is a complete listener for a parse tree produced by MongoShellParser.
type MongoShellParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterShowDatabases is called when entering the showDatabases production.
	EnterShowDatabases(c *ShowDatabasesContext)

	// EnterShowCollections is called when entering the showCollections production.
	EnterShowCollections(c *ShowCollectionsContext)

	// EnterGetCollectionNames is called when entering the getCollectionNames production.
	EnterGetCollectionNames(c *GetCollectionNamesContext)

	// EnterGetCollectionInfos is called when entering the getCollectionInfos production.
	EnterGetCollectionInfos(c *GetCollectionInfosContext)

	// EnterCreateCollection is called when entering the createCollection production.
	EnterCreateCollection(c *CreateCollectionContext)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterDbStats is called when entering the dbStats production.
	EnterDbStats(c *DbStatsContext)

	// EnterServerStatus is called when entering the serverStatus production.
	EnterServerStatus(c *ServerStatusContext)

	// EnterServerBuildInfo is called when entering the serverBuildInfo production.
	EnterServerBuildInfo(c *ServerBuildInfoContext)

	// EnterDbVersion is called when entering the dbVersion production.
	EnterDbVersion(c *DbVersionContext)

	// EnterHostInfo is called when entering the hostInfo production.
	EnterHostInfo(c *HostInfoContext)

	// EnterListCommands is called when entering the listCommands production.
	EnterListCommands(c *ListCommandsContext)

	// EnterRunCommand is called when entering the runCommand production.
	EnterRunCommand(c *RunCommandContext)

	// EnterAdminCommand is called when entering the adminCommand production.
	EnterAdminCommand(c *AdminCommandContext)

	// EnterGetName is called when entering the getName production.
	EnterGetName(c *GetNameContext)

	// EnterGetMongo is called when entering the getMongo production.
	EnterGetMongo(c *GetMongoContext)

	// EnterGetSiblingDB is called when entering the getSiblingDB production.
	EnterGetSiblingDB(c *GetSiblingDBContext)

	// EnterDbAggregate is called when entering the dbAggregate production.
	EnterDbAggregate(c *DbAggregateContext)

	// EnterDbAuth is called when entering the dbAuth production.
	EnterDbAuth(c *DbAuthContext)

	// EnterDbChangeUserPassword is called when entering the dbChangeUserPassword production.
	EnterDbChangeUserPassword(c *DbChangeUserPasswordContext)

	// EnterDbCloneDatabase is called when entering the dbCloneDatabase production.
	EnterDbCloneDatabase(c *DbCloneDatabaseContext)

	// EnterDbCommandHelp is called when entering the dbCommandHelp production.
	EnterDbCommandHelp(c *DbCommandHelpContext)

	// EnterDbCopyDatabase is called when entering the dbCopyDatabase production.
	EnterDbCopyDatabase(c *DbCopyDatabaseContext)

	// EnterDbCreateRole is called when entering the dbCreateRole production.
	EnterDbCreateRole(c *DbCreateRoleContext)

	// EnterDbCreateUser is called when entering the dbCreateUser production.
	EnterDbCreateUser(c *DbCreateUserContext)

	// EnterDbCreateView is called when entering the dbCreateView production.
	EnterDbCreateView(c *DbCreateViewContext)

	// EnterDbCurrentOp is called when entering the dbCurrentOp production.
	EnterDbCurrentOp(c *DbCurrentOpContext)

	// EnterDbDropAllRoles is called when entering the dbDropAllRoles production.
	EnterDbDropAllRoles(c *DbDropAllRolesContext)

	// EnterDbDropAllUsers is called when entering the dbDropAllUsers production.
	EnterDbDropAllUsers(c *DbDropAllUsersContext)

	// EnterDbDropRole is called when entering the dbDropRole production.
	EnterDbDropRole(c *DbDropRoleContext)

	// EnterDbDropUser is called when entering the dbDropUser production.
	EnterDbDropUser(c *DbDropUserContext)

	// EnterDbFsyncLock is called when entering the dbFsyncLock production.
	EnterDbFsyncLock(c *DbFsyncLockContext)

	// EnterDbFsyncUnlock is called when entering the dbFsyncUnlock production.
	EnterDbFsyncUnlock(c *DbFsyncUnlockContext)

	// EnterDbGetLogComponents is called when entering the dbGetLogComponents production.
	EnterDbGetLogComponents(c *DbGetLogComponentsContext)

	// EnterDbGetProfilingLevel is called when entering the dbGetProfilingLevel production.
	EnterDbGetProfilingLevel(c *DbGetProfilingLevelContext)

	// EnterDbGetProfilingStatus is called when entering the dbGetProfilingStatus production.
	EnterDbGetProfilingStatus(c *DbGetProfilingStatusContext)

	// EnterDbGetReplicationInfo is called when entering the dbGetReplicationInfo production.
	EnterDbGetReplicationInfo(c *DbGetReplicationInfoContext)

	// EnterDbGetRole is called when entering the dbGetRole production.
	EnterDbGetRole(c *DbGetRoleContext)

	// EnterDbGetRoles is called when entering the dbGetRoles production.
	EnterDbGetRoles(c *DbGetRolesContext)

	// EnterDbGetUser is called when entering the dbGetUser production.
	EnterDbGetUser(c *DbGetUserContext)

	// EnterDbGetUsers is called when entering the dbGetUsers production.
	EnterDbGetUsers(c *DbGetUsersContext)

	// EnterDbGrantPrivilegesToRole is called when entering the dbGrantPrivilegesToRole production.
	EnterDbGrantPrivilegesToRole(c *DbGrantPrivilegesToRoleContext)

	// EnterDbGrantRolesToRole is called when entering the dbGrantRolesToRole production.
	EnterDbGrantRolesToRole(c *DbGrantRolesToRoleContext)

	// EnterDbGrantRolesToUser is called when entering the dbGrantRolesToUser production.
	EnterDbGrantRolesToUser(c *DbGrantRolesToUserContext)

	// EnterDbHello is called when entering the dbHello production.
	EnterDbHello(c *DbHelloContext)

	// EnterDbIsMaster is called when entering the dbIsMaster production.
	EnterDbIsMaster(c *DbIsMasterContext)

	// EnterDbKillOp is called when entering the dbKillOp production.
	EnterDbKillOp(c *DbKillOpContext)

	// EnterDbLogout is called when entering the dbLogout production.
	EnterDbLogout(c *DbLogoutContext)

	// EnterDbPrintCollectionStats is called when entering the dbPrintCollectionStats production.
	EnterDbPrintCollectionStats(c *DbPrintCollectionStatsContext)

	// EnterDbPrintReplicationInfo is called when entering the dbPrintReplicationInfo production.
	EnterDbPrintReplicationInfo(c *DbPrintReplicationInfoContext)

	// EnterDbPrintSecondaryReplicationInfo is called when entering the dbPrintSecondaryReplicationInfo production.
	EnterDbPrintSecondaryReplicationInfo(c *DbPrintSecondaryReplicationInfoContext)

	// EnterDbPrintShardingStatus is called when entering the dbPrintShardingStatus production.
	EnterDbPrintShardingStatus(c *DbPrintShardingStatusContext)

	// EnterDbPrintSlaveReplicationInfo is called when entering the dbPrintSlaveReplicationInfo production.
	EnterDbPrintSlaveReplicationInfo(c *DbPrintSlaveReplicationInfoContext)

	// EnterDbRevokePrivilegesFromRole is called when entering the dbRevokePrivilegesFromRole production.
	EnterDbRevokePrivilegesFromRole(c *DbRevokePrivilegesFromRoleContext)

	// EnterDbRevokeRolesFromRole is called when entering the dbRevokeRolesFromRole production.
	EnterDbRevokeRolesFromRole(c *DbRevokeRolesFromRoleContext)

	// EnterDbRevokeRolesFromUser is called when entering the dbRevokeRolesFromUser production.
	EnterDbRevokeRolesFromUser(c *DbRevokeRolesFromUserContext)

	// EnterDbRotateCertificates is called when entering the dbRotateCertificates production.
	EnterDbRotateCertificates(c *DbRotateCertificatesContext)

	// EnterDbSetLogLevel is called when entering the dbSetLogLevel production.
	EnterDbSetLogLevel(c *DbSetLogLevelContext)

	// EnterDbSetProfilingLevel is called when entering the dbSetProfilingLevel production.
	EnterDbSetProfilingLevel(c *DbSetProfilingLevelContext)

	// EnterDbSetSecondaryOk is called when entering the dbSetSecondaryOk production.
	EnterDbSetSecondaryOk(c *DbSetSecondaryOkContext)

	// EnterDbSetWriteConcern is called when entering the dbSetWriteConcern production.
	EnterDbSetWriteConcern(c *DbSetWriteConcernContext)

	// EnterDbShutdownServer is called when entering the dbShutdownServer production.
	EnterDbShutdownServer(c *DbShutdownServerContext)

	// EnterDbUpdateRole is called when entering the dbUpdateRole production.
	EnterDbUpdateRole(c *DbUpdateRoleContext)

	// EnterDbUpdateUser is called when entering the dbUpdateUser production.
	EnterDbUpdateUser(c *DbUpdateUserContext)

	// EnterDbWatch is called when entering the dbWatch production.
	EnterDbWatch(c *DbWatchContext)

	// EnterCollectionOperation is called when entering the collectionOperation production.
	EnterCollectionOperation(c *CollectionOperationContext)

	// EnterBulkStatement is called when entering the bulkStatement production.
	EnterBulkStatement(c *BulkStatementContext)

	// EnterBulkInitMethod is called when entering the bulkInitMethod production.
	EnterBulkInitMethod(c *BulkInitMethodContext)

	// EnterBulkMethodChain is called when entering the bulkMethodChain production.
	EnterBulkMethodChain(c *BulkMethodChainContext)

	// EnterBulkFind is called when entering the bulkFind production.
	EnterBulkFind(c *BulkFindContext)

	// EnterBulkInsert is called when entering the bulkInsert production.
	EnterBulkInsert(c *BulkInsertContext)

	// EnterBulkRemove is called when entering the bulkRemove production.
	EnterBulkRemove(c *BulkRemoveContext)

	// EnterBulkExecute is called when entering the bulkExecute production.
	EnterBulkExecute(c *BulkExecuteContext)

	// EnterBulkGetOperations is called when entering the bulkGetOperations production.
	EnterBulkGetOperations(c *BulkGetOperationsContext)

	// EnterBulkToString is called when entering the bulkToString production.
	EnterBulkToString(c *BulkToStringContext)

	// EnterBulkGenericMethod is called when entering the bulkGenericMethod production.
	EnterBulkGenericMethod(c *BulkGenericMethodContext)

	// EnterMongoConnection is called when entering the mongoConnection production.
	EnterMongoConnection(c *MongoConnectionContext)

	// EnterConnectCall is called when entering the connectCall production.
	EnterConnectCall(c *ConnectCallContext)

	// EnterDbGetMongoChain is called when entering the dbGetMongoChain production.
	EnterDbGetMongoChain(c *DbGetMongoChainContext)

	// EnterConnectionMethodChain is called when entering the connectionMethodChain production.
	EnterConnectionMethodChain(c *ConnectionMethodChainContext)

	// EnterRsStatement is called when entering the rsStatement production.
	EnterRsStatement(c *RsStatementContext)

	// EnterShStatement is called when entering the shStatement production.
	EnterShStatement(c *ShStatementContext)

	// EnterKeyVaultStatement is called when entering the keyVaultStatement production.
	EnterKeyVaultStatement(c *KeyVaultStatementContext)

	// EnterClientEncryptionStatement is called when entering the clientEncryptionStatement production.
	EnterClientEncryptionStatement(c *ClientEncryptionStatementContext)

	// EnterPlanCacheStatement is called when entering the planCacheStatement production.
	EnterPlanCacheStatement(c *PlanCacheStatementContext)

	// EnterSpStatement is called when entering the spStatement production.
	EnterSpStatement(c *SpStatementContext)

	// EnterNativeFunctionCall is called when entering the nativeFunctionCall production.
	EnterNativeFunctionCall(c *NativeFunctionCallContext)

	// EnterConnGetDB is called when entering the connGetDB production.
	EnterConnGetDB(c *ConnGetDBContext)

	// EnterConnGetReadConcern is called when entering the connGetReadConcern production.
	EnterConnGetReadConcern(c *ConnGetReadConcernContext)

	// EnterConnGetReadPref is called when entering the connGetReadPref production.
	EnterConnGetReadPref(c *ConnGetReadPrefContext)

	// EnterConnGetReadPrefMode is called when entering the connGetReadPrefMode production.
	EnterConnGetReadPrefMode(c *ConnGetReadPrefModeContext)

	// EnterConnGetReadPrefTagSet is called when entering the connGetReadPrefTagSet production.
	EnterConnGetReadPrefTagSet(c *ConnGetReadPrefTagSetContext)

	// EnterConnGetWriteConcern is called when entering the connGetWriteConcern production.
	EnterConnGetWriteConcern(c *ConnGetWriteConcernContext)

	// EnterConnSetReadPref is called when entering the connSetReadPref production.
	EnterConnSetReadPref(c *ConnSetReadPrefContext)

	// EnterConnSetReadConcern is called when entering the connSetReadConcern production.
	EnterConnSetReadConcern(c *ConnSetReadConcernContext)

	// EnterConnSetWriteConcern is called when entering the connSetWriteConcern production.
	EnterConnSetWriteConcern(c *ConnSetWriteConcernContext)

	// EnterConnStartSession is called when entering the connStartSession production.
	EnterConnStartSession(c *ConnStartSessionContext)

	// EnterConnWatch is called when entering the connWatch production.
	EnterConnWatch(c *ConnWatchContext)

	// EnterConnClose is called when entering the connClose production.
	EnterConnClose(c *ConnCloseContext)

	// EnterConnAdminCommand is called when entering the connAdminCommand production.
	EnterConnAdminCommand(c *ConnAdminCommandContext)

	// EnterConnGetDBNames is called when entering the connGetDBNames production.
	EnterConnGetDBNames(c *ConnGetDBNamesContext)

	// EnterConnGenericMethod is called when entering the connGenericMethod production.
	EnterConnGenericMethod(c *ConnGenericMethodContext)

	// EnterDotAccess is called when entering the dotAccess production.
	EnterDotAccess(c *DotAccessContext)

	// EnterBracketAccess is called when entering the bracketAccess production.
	EnterBracketAccess(c *BracketAccessContext)

	// EnterGetCollectionAccess is called when entering the getCollectionAccess production.
	EnterGetCollectionAccess(c *GetCollectionAccessContext)

	// EnterMethodChain is called when entering the methodChain production.
	EnterMethodChain(c *MethodChainContext)

	// EnterCollectionMethodCall is called when entering the collectionMethodCall production.
	EnterCollectionMethodCall(c *CollectionMethodCallContext)

	// EnterCursorMethodCall is called when entering the cursorMethodCall production.
	EnterCursorMethodCall(c *CursorMethodCallContext)

	// EnterFindMethod is called when entering the findMethod production.
	EnterFindMethod(c *FindMethodContext)

	// EnterFindOneMethod is called when entering the findOneMethod production.
	EnterFindOneMethod(c *FindOneMethodContext)

	// EnterCountDocumentsMethod is called when entering the countDocumentsMethod production.
	EnterCountDocumentsMethod(c *CountDocumentsMethodContext)

	// EnterEstimatedDocumentCountMethod is called when entering the estimatedDocumentCountMethod production.
	EnterEstimatedDocumentCountMethod(c *EstimatedDocumentCountMethodContext)

	// EnterDistinctMethod is called when entering the distinctMethod production.
	EnterDistinctMethod(c *DistinctMethodContext)

	// EnterAggregateMethod is called when entering the aggregateMethod production.
	EnterAggregateMethod(c *AggregateMethodContext)

	// EnterGetIndexesMethod is called when entering the getIndexesMethod production.
	EnterGetIndexesMethod(c *GetIndexesMethodContext)

	// EnterInsertOneMethod is called when entering the insertOneMethod production.
	EnterInsertOneMethod(c *InsertOneMethodContext)

	// EnterInsertManyMethod is called when entering the insertManyMethod production.
	EnterInsertManyMethod(c *InsertManyMethodContext)

	// EnterUpdateOneMethod is called when entering the updateOneMethod production.
	EnterUpdateOneMethod(c *UpdateOneMethodContext)

	// EnterUpdateManyMethod is called when entering the updateManyMethod production.
	EnterUpdateManyMethod(c *UpdateManyMethodContext)

	// EnterDeleteOneMethod is called when entering the deleteOneMethod production.
	EnterDeleteOneMethod(c *DeleteOneMethodContext)

	// EnterDeleteManyMethod is called when entering the deleteManyMethod production.
	EnterDeleteManyMethod(c *DeleteManyMethodContext)

	// EnterReplaceOneMethod is called when entering the replaceOneMethod production.
	EnterReplaceOneMethod(c *ReplaceOneMethodContext)

	// EnterFindOneAndUpdateMethod is called when entering the findOneAndUpdateMethod production.
	EnterFindOneAndUpdateMethod(c *FindOneAndUpdateMethodContext)

	// EnterFindOneAndReplaceMethod is called when entering the findOneAndReplaceMethod production.
	EnterFindOneAndReplaceMethod(c *FindOneAndReplaceMethodContext)

	// EnterFindOneAndDeleteMethod is called when entering the findOneAndDeleteMethod production.
	EnterFindOneAndDeleteMethod(c *FindOneAndDeleteMethodContext)

	// EnterCreateIndexMethod is called when entering the createIndexMethod production.
	EnterCreateIndexMethod(c *CreateIndexMethodContext)

	// EnterCreateIndexesMethod is called when entering the createIndexesMethod production.
	EnterCreateIndexesMethod(c *CreateIndexesMethodContext)

	// EnterDropIndexMethod is called when entering the dropIndexMethod production.
	EnterDropIndexMethod(c *DropIndexMethodContext)

	// EnterDropIndexesMethod is called when entering the dropIndexesMethod production.
	EnterDropIndexesMethod(c *DropIndexesMethodContext)

	// EnterDropMethod is called when entering the dropMethod production.
	EnterDropMethod(c *DropMethodContext)

	// EnterRenameCollectionMethod is called when entering the renameCollectionMethod production.
	EnterRenameCollectionMethod(c *RenameCollectionMethodContext)

	// EnterStatsMethod is called when entering the statsMethod production.
	EnterStatsMethod(c *StatsMethodContext)

	// EnterStorageSizeMethod is called when entering the storageSizeMethod production.
	EnterStorageSizeMethod(c *StorageSizeMethodContext)

	// EnterTotalIndexSizeMethod is called when entering the totalIndexSizeMethod production.
	EnterTotalIndexSizeMethod(c *TotalIndexSizeMethodContext)

	// EnterTotalSizeMethod is called when entering the totalSizeMethod production.
	EnterTotalSizeMethod(c *TotalSizeMethodContext)

	// EnterDataSizeMethod is called when entering the dataSizeMethod production.
	EnterDataSizeMethod(c *DataSizeMethodContext)

	// EnterIsCappedMethod is called when entering the isCappedMethod production.
	EnterIsCappedMethod(c *IsCappedMethodContext)

	// EnterValidateMethod is called when entering the validateMethod production.
	EnterValidateMethod(c *ValidateMethodContext)

	// EnterLatencyStatsMethod is called when entering the latencyStatsMethod production.
	EnterLatencyStatsMethod(c *LatencyStatsMethodContext)

	// EnterWatchMethod is called when entering the watchMethod production.
	EnterWatchMethod(c *WatchMethodContext)

	// EnterBulkWriteMethod is called when entering the bulkWriteMethod production.
	EnterBulkWriteMethod(c *BulkWriteMethodContext)

	// EnterCollectionCountMethod is called when entering the collectionCountMethod production.
	EnterCollectionCountMethod(c *CollectionCountMethodContext)

	// EnterCollectionInsertMethod is called when entering the collectionInsertMethod production.
	EnterCollectionInsertMethod(c *CollectionInsertMethodContext)

	// EnterCollectionRemoveMethod is called when entering the collectionRemoveMethod production.
	EnterCollectionRemoveMethod(c *CollectionRemoveMethodContext)

	// EnterUpdateMethod is called when entering the updateMethod production.
	EnterUpdateMethod(c *UpdateMethodContext)

	// EnterMapReduceMethod is called when entering the mapReduceMethod production.
	EnterMapReduceMethod(c *MapReduceMethodContext)

	// EnterFindAndModifyMethod is called when entering the findAndModifyMethod production.
	EnterFindAndModifyMethod(c *FindAndModifyMethodContext)

	// EnterCollectionExplainMethod is called when entering the collectionExplainMethod production.
	EnterCollectionExplainMethod(c *CollectionExplainMethodContext)

	// EnterAnalyzeShardKeyMethod is called when entering the analyzeShardKeyMethod production.
	EnterAnalyzeShardKeyMethod(c *AnalyzeShardKeyMethodContext)

	// EnterConfigureQueryAnalyzerMethod is called when entering the configureQueryAnalyzerMethod production.
	EnterConfigureQueryAnalyzerMethod(c *ConfigureQueryAnalyzerMethodContext)

	// EnterCompactStructuredEncryptionDataMethod is called when entering the compactStructuredEncryptionDataMethod production.
	EnterCompactStructuredEncryptionDataMethod(c *CompactStructuredEncryptionDataMethodContext)

	// EnterHideIndexMethod is called when entering the hideIndexMethod production.
	EnterHideIndexMethod(c *HideIndexMethodContext)

	// EnterUnhideIndexMethod is called when entering the unhideIndexMethod production.
	EnterUnhideIndexMethod(c *UnhideIndexMethodContext)

	// EnterReIndexMethod is called when entering the reIndexMethod production.
	EnterReIndexMethod(c *ReIndexMethodContext)

	// EnterGetShardDistributionMethod is called when entering the getShardDistributionMethod production.
	EnterGetShardDistributionMethod(c *GetShardDistributionMethodContext)

	// EnterGetShardVersionMethod is called when entering the getShardVersionMethod production.
	EnterGetShardVersionMethod(c *GetShardVersionMethodContext)

	// EnterCreateSearchIndexMethod is called when entering the createSearchIndexMethod production.
	EnterCreateSearchIndexMethod(c *CreateSearchIndexMethodContext)

	// EnterCreateSearchIndexesMethod is called when entering the createSearchIndexesMethod production.
	EnterCreateSearchIndexesMethod(c *CreateSearchIndexesMethodContext)

	// EnterDropSearchIndexMethod is called when entering the dropSearchIndexMethod production.
	EnterDropSearchIndexMethod(c *DropSearchIndexMethodContext)

	// EnterUpdateSearchIndexMethod is called when entering the updateSearchIndexMethod production.
	EnterUpdateSearchIndexMethod(c *UpdateSearchIndexMethodContext)

	// EnterSortMethod is called when entering the sortMethod production.
	EnterSortMethod(c *SortMethodContext)

	// EnterLimitMethod is called when entering the limitMethod production.
	EnterLimitMethod(c *LimitMethodContext)

	// EnterSkipMethod is called when entering the skipMethod production.
	EnterSkipMethod(c *SkipMethodContext)

	// EnterCountMethod is called when entering the countMethod production.
	EnterCountMethod(c *CountMethodContext)

	// EnterProjectionMethod is called when entering the projectionMethod production.
	EnterProjectionMethod(c *ProjectionMethodContext)

	// EnterBatchSizeMethod is called when entering the batchSizeMethod production.
	EnterBatchSizeMethod(c *BatchSizeMethodContext)

	// EnterCloseMethod is called when entering the closeMethod production.
	EnterCloseMethod(c *CloseMethodContext)

	// EnterCollationMethod is called when entering the collationMethod production.
	EnterCollationMethod(c *CollationMethodContext)

	// EnterCommentMethod is called when entering the commentMethod production.
	EnterCommentMethod(c *CommentMethodContext)

	// EnterExplainMethod is called when entering the explainMethod production.
	EnterExplainMethod(c *ExplainMethodContext)

	// EnterForEachMethod is called when entering the forEachMethod production.
	EnterForEachMethod(c *ForEachMethodContext)

	// EnterHasNextMethod is called when entering the hasNextMethod production.
	EnterHasNextMethod(c *HasNextMethodContext)

	// EnterHintMethod is called when entering the hintMethod production.
	EnterHintMethod(c *HintMethodContext)

	// EnterIsClosedMethod is called when entering the isClosedMethod production.
	EnterIsClosedMethod(c *IsClosedMethodContext)

	// EnterIsExhaustedMethod is called when entering the isExhaustedMethod production.
	EnterIsExhaustedMethod(c *IsExhaustedMethodContext)

	// EnterItcountMethod is called when entering the itcountMethod production.
	EnterItcountMethod(c *ItcountMethodContext)

	// EnterMapMethod is called when entering the mapMethod production.
	EnterMapMethod(c *MapMethodContext)

	// EnterMaxMethod is called when entering the maxMethod production.
	EnterMaxMethod(c *MaxMethodContext)

	// EnterMaxAwaitTimeMSMethod is called when entering the maxAwaitTimeMSMethod production.
	EnterMaxAwaitTimeMSMethod(c *MaxAwaitTimeMSMethodContext)

	// EnterMaxTimeMSMethod is called when entering the maxTimeMSMethod production.
	EnterMaxTimeMSMethod(c *MaxTimeMSMethodContext)

	// EnterMinMethod is called when entering the minMethod production.
	EnterMinMethod(c *MinMethodContext)

	// EnterNextMethod is called when entering the nextMethod production.
	EnterNextMethod(c *NextMethodContext)

	// EnterNoCursorTimeoutMethod is called when entering the noCursorTimeoutMethod production.
	EnterNoCursorTimeoutMethod(c *NoCursorTimeoutMethodContext)

	// EnterObjsLeftInBatchMethod is called when entering the objsLeftInBatchMethod production.
	EnterObjsLeftInBatchMethod(c *ObjsLeftInBatchMethodContext)

	// EnterPrettyMethod is called when entering the prettyMethod production.
	EnterPrettyMethod(c *PrettyMethodContext)

	// EnterReadConcernMethod is called when entering the readConcernMethod production.
	EnterReadConcernMethod(c *ReadConcernMethodContext)

	// EnterReadPrefMethod is called when entering the readPrefMethod production.
	EnterReadPrefMethod(c *ReadPrefMethodContext)

	// EnterReturnKeyMethod is called when entering the returnKeyMethod production.
	EnterReturnKeyMethod(c *ReturnKeyMethodContext)

	// EnterShowRecordIdMethod is called when entering the showRecordIdMethod production.
	EnterShowRecordIdMethod(c *ShowRecordIdMethodContext)

	// EnterSizeMethod is called when entering the sizeMethod production.
	EnterSizeMethod(c *SizeMethodContext)

	// EnterTailableMethod is called when entering the tailableMethod production.
	EnterTailableMethod(c *TailableMethodContext)

	// EnterToArrayMethod is called when entering the toArrayMethod production.
	EnterToArrayMethod(c *ToArrayMethodContext)

	// EnterTryNextMethod is called when entering the tryNextMethod production.
	EnterTryNextMethod(c *TryNextMethodContext)

	// EnterAllowDiskUseMethod is called when entering the allowDiskUseMethod production.
	EnterAllowDiskUseMethod(c *AllowDiskUseMethodContext)

	// EnterAddOptionMethod is called when entering the addOptionMethod production.
	EnterAddOptionMethod(c *AddOptionMethodContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterUnquotedKey is called when entering the unquotedKey production.
	EnterUnquotedKey(c *UnquotedKeyContext)

	// EnterQuotedKey is called when entering the quotedKey production.
	EnterQuotedKey(c *QuotedKeyContext)

	// EnterDocumentValue is called when entering the documentValue production.
	EnterDocumentValue(c *DocumentValueContext)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterHelperValue is called when entering the helperValue production.
	EnterHelperValue(c *HelperValueContext)

	// EnterRegexLiteralValue is called when entering the regexLiteralValue production.
	EnterRegexLiteralValue(c *RegexLiteralValueContext)

	// EnterRegexpConstructorValue is called when entering the regexpConstructorValue production.
	EnterRegexpConstructorValue(c *RegexpConstructorValueContext)

	// EnterLiteralValue is called when entering the literalValue production.
	EnterLiteralValue(c *LiteralValueContext)

	// EnterNewKeywordValue is called when entering the newKeywordValue production.
	EnterNewKeywordValue(c *NewKeywordValueContext)

	// EnterNewKeywordError is called when entering the newKeywordError production.
	EnterNewKeywordError(c *NewKeywordErrorContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterHelperFunction is called when entering the helperFunction production.
	EnterHelperFunction(c *HelperFunctionContext)

	// EnterObjectIdHelper is called when entering the objectIdHelper production.
	EnterObjectIdHelper(c *ObjectIdHelperContext)

	// EnterIsoDateHelper is called when entering the isoDateHelper production.
	EnterIsoDateHelper(c *IsoDateHelperContext)

	// EnterDateHelper is called when entering the dateHelper production.
	EnterDateHelper(c *DateHelperContext)

	// EnterUuidHelper is called when entering the uuidHelper production.
	EnterUuidHelper(c *UuidHelperContext)

	// EnterLongHelper is called when entering the longHelper production.
	EnterLongHelper(c *LongHelperContext)

	// EnterInt32Helper is called when entering the int32Helper production.
	EnterInt32Helper(c *Int32HelperContext)

	// EnterDoubleHelper is called when entering the doubleHelper production.
	EnterDoubleHelper(c *DoubleHelperContext)

	// EnterDecimal128Helper is called when entering the decimal128Helper production.
	EnterDecimal128Helper(c *Decimal128HelperContext)

	// EnterTimestampDocHelper is called when entering the timestampDocHelper production.
	EnterTimestampDocHelper(c *TimestampDocHelperContext)

	// EnterTimestampArgsHelper is called when entering the timestampArgsHelper production.
	EnterTimestampArgsHelper(c *TimestampArgsHelperContext)

	// EnterRegExpConstructor is called when entering the regExpConstructor production.
	EnterRegExpConstructor(c *RegExpConstructorContext)

	// EnterBinDataHelper is called when entering the binDataHelper production.
	EnterBinDataHelper(c *BinDataHelperContext)

	// EnterBinaryHelper is called when entering the binaryHelper production.
	EnterBinaryHelper(c *BinaryHelperContext)

	// EnterBsonRegExpHelper is called when entering the bsonRegExpHelper production.
	EnterBsonRegExpHelper(c *BsonRegExpHelperContext)

	// EnterHexDataHelper is called when entering the hexDataHelper production.
	EnterHexDataHelper(c *HexDataHelperContext)

	// EnterStringLiteralValue is called when entering the stringLiteralValue production.
	EnterStringLiteralValue(c *StringLiteralValueContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterTrueLiteral is called when entering the trueLiteral production.
	EnterTrueLiteral(c *TrueLiteralContext)

	// EnterFalseLiteral is called when entering the falseLiteral production.
	EnterFalseLiteral(c *FalseLiteralContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitShowDatabases is called when exiting the showDatabases production.
	ExitShowDatabases(c *ShowDatabasesContext)

	// ExitShowCollections is called when exiting the showCollections production.
	ExitShowCollections(c *ShowCollectionsContext)

	// ExitGetCollectionNames is called when exiting the getCollectionNames production.
	ExitGetCollectionNames(c *GetCollectionNamesContext)

	// ExitGetCollectionInfos is called when exiting the getCollectionInfos production.
	ExitGetCollectionInfos(c *GetCollectionInfosContext)

	// ExitCreateCollection is called when exiting the createCollection production.
	ExitCreateCollection(c *CreateCollectionContext)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitDbStats is called when exiting the dbStats production.
	ExitDbStats(c *DbStatsContext)

	// ExitServerStatus is called when exiting the serverStatus production.
	ExitServerStatus(c *ServerStatusContext)

	// ExitServerBuildInfo is called when exiting the serverBuildInfo production.
	ExitServerBuildInfo(c *ServerBuildInfoContext)

	// ExitDbVersion is called when exiting the dbVersion production.
	ExitDbVersion(c *DbVersionContext)

	// ExitHostInfo is called when exiting the hostInfo production.
	ExitHostInfo(c *HostInfoContext)

	// ExitListCommands is called when exiting the listCommands production.
	ExitListCommands(c *ListCommandsContext)

	// ExitRunCommand is called when exiting the runCommand production.
	ExitRunCommand(c *RunCommandContext)

	// ExitAdminCommand is called when exiting the adminCommand production.
	ExitAdminCommand(c *AdminCommandContext)

	// ExitGetName is called when exiting the getName production.
	ExitGetName(c *GetNameContext)

	// ExitGetMongo is called when exiting the getMongo production.
	ExitGetMongo(c *GetMongoContext)

	// ExitGetSiblingDB is called when exiting the getSiblingDB production.
	ExitGetSiblingDB(c *GetSiblingDBContext)

	// ExitDbAggregate is called when exiting the dbAggregate production.
	ExitDbAggregate(c *DbAggregateContext)

	// ExitDbAuth is called when exiting the dbAuth production.
	ExitDbAuth(c *DbAuthContext)

	// ExitDbChangeUserPassword is called when exiting the dbChangeUserPassword production.
	ExitDbChangeUserPassword(c *DbChangeUserPasswordContext)

	// ExitDbCloneDatabase is called when exiting the dbCloneDatabase production.
	ExitDbCloneDatabase(c *DbCloneDatabaseContext)

	// ExitDbCommandHelp is called when exiting the dbCommandHelp production.
	ExitDbCommandHelp(c *DbCommandHelpContext)

	// ExitDbCopyDatabase is called when exiting the dbCopyDatabase production.
	ExitDbCopyDatabase(c *DbCopyDatabaseContext)

	// ExitDbCreateRole is called when exiting the dbCreateRole production.
	ExitDbCreateRole(c *DbCreateRoleContext)

	// ExitDbCreateUser is called when exiting the dbCreateUser production.
	ExitDbCreateUser(c *DbCreateUserContext)

	// ExitDbCreateView is called when exiting the dbCreateView production.
	ExitDbCreateView(c *DbCreateViewContext)

	// ExitDbCurrentOp is called when exiting the dbCurrentOp production.
	ExitDbCurrentOp(c *DbCurrentOpContext)

	// ExitDbDropAllRoles is called when exiting the dbDropAllRoles production.
	ExitDbDropAllRoles(c *DbDropAllRolesContext)

	// ExitDbDropAllUsers is called when exiting the dbDropAllUsers production.
	ExitDbDropAllUsers(c *DbDropAllUsersContext)

	// ExitDbDropRole is called when exiting the dbDropRole production.
	ExitDbDropRole(c *DbDropRoleContext)

	// ExitDbDropUser is called when exiting the dbDropUser production.
	ExitDbDropUser(c *DbDropUserContext)

	// ExitDbFsyncLock is called when exiting the dbFsyncLock production.
	ExitDbFsyncLock(c *DbFsyncLockContext)

	// ExitDbFsyncUnlock is called when exiting the dbFsyncUnlock production.
	ExitDbFsyncUnlock(c *DbFsyncUnlockContext)

	// ExitDbGetLogComponents is called when exiting the dbGetLogComponents production.
	ExitDbGetLogComponents(c *DbGetLogComponentsContext)

	// ExitDbGetProfilingLevel is called when exiting the dbGetProfilingLevel production.
	ExitDbGetProfilingLevel(c *DbGetProfilingLevelContext)

	// ExitDbGetProfilingStatus is called when exiting the dbGetProfilingStatus production.
	ExitDbGetProfilingStatus(c *DbGetProfilingStatusContext)

	// ExitDbGetReplicationInfo is called when exiting the dbGetReplicationInfo production.
	ExitDbGetReplicationInfo(c *DbGetReplicationInfoContext)

	// ExitDbGetRole is called when exiting the dbGetRole production.
	ExitDbGetRole(c *DbGetRoleContext)

	// ExitDbGetRoles is called when exiting the dbGetRoles production.
	ExitDbGetRoles(c *DbGetRolesContext)

	// ExitDbGetUser is called when exiting the dbGetUser production.
	ExitDbGetUser(c *DbGetUserContext)

	// ExitDbGetUsers is called when exiting the dbGetUsers production.
	ExitDbGetUsers(c *DbGetUsersContext)

	// ExitDbGrantPrivilegesToRole is called when exiting the dbGrantPrivilegesToRole production.
	ExitDbGrantPrivilegesToRole(c *DbGrantPrivilegesToRoleContext)

	// ExitDbGrantRolesToRole is called when exiting the dbGrantRolesToRole production.
	ExitDbGrantRolesToRole(c *DbGrantRolesToRoleContext)

	// ExitDbGrantRolesToUser is called when exiting the dbGrantRolesToUser production.
	ExitDbGrantRolesToUser(c *DbGrantRolesToUserContext)

	// ExitDbHello is called when exiting the dbHello production.
	ExitDbHello(c *DbHelloContext)

	// ExitDbIsMaster is called when exiting the dbIsMaster production.
	ExitDbIsMaster(c *DbIsMasterContext)

	// ExitDbKillOp is called when exiting the dbKillOp production.
	ExitDbKillOp(c *DbKillOpContext)

	// ExitDbLogout is called when exiting the dbLogout production.
	ExitDbLogout(c *DbLogoutContext)

	// ExitDbPrintCollectionStats is called when exiting the dbPrintCollectionStats production.
	ExitDbPrintCollectionStats(c *DbPrintCollectionStatsContext)

	// ExitDbPrintReplicationInfo is called when exiting the dbPrintReplicationInfo production.
	ExitDbPrintReplicationInfo(c *DbPrintReplicationInfoContext)

	// ExitDbPrintSecondaryReplicationInfo is called when exiting the dbPrintSecondaryReplicationInfo production.
	ExitDbPrintSecondaryReplicationInfo(c *DbPrintSecondaryReplicationInfoContext)

	// ExitDbPrintShardingStatus is called when exiting the dbPrintShardingStatus production.
	ExitDbPrintShardingStatus(c *DbPrintShardingStatusContext)

	// ExitDbPrintSlaveReplicationInfo is called when exiting the dbPrintSlaveReplicationInfo production.
	ExitDbPrintSlaveReplicationInfo(c *DbPrintSlaveReplicationInfoContext)

	// ExitDbRevokePrivilegesFromRole is called when exiting the dbRevokePrivilegesFromRole production.
	ExitDbRevokePrivilegesFromRole(c *DbRevokePrivilegesFromRoleContext)

	// ExitDbRevokeRolesFromRole is called when exiting the dbRevokeRolesFromRole production.
	ExitDbRevokeRolesFromRole(c *DbRevokeRolesFromRoleContext)

	// ExitDbRevokeRolesFromUser is called when exiting the dbRevokeRolesFromUser production.
	ExitDbRevokeRolesFromUser(c *DbRevokeRolesFromUserContext)

	// ExitDbRotateCertificates is called when exiting the dbRotateCertificates production.
	ExitDbRotateCertificates(c *DbRotateCertificatesContext)

	// ExitDbSetLogLevel is called when exiting the dbSetLogLevel production.
	ExitDbSetLogLevel(c *DbSetLogLevelContext)

	// ExitDbSetProfilingLevel is called when exiting the dbSetProfilingLevel production.
	ExitDbSetProfilingLevel(c *DbSetProfilingLevelContext)

	// ExitDbSetSecondaryOk is called when exiting the dbSetSecondaryOk production.
	ExitDbSetSecondaryOk(c *DbSetSecondaryOkContext)

	// ExitDbSetWriteConcern is called when exiting the dbSetWriteConcern production.
	ExitDbSetWriteConcern(c *DbSetWriteConcernContext)

	// ExitDbShutdownServer is called when exiting the dbShutdownServer production.
	ExitDbShutdownServer(c *DbShutdownServerContext)

	// ExitDbUpdateRole is called when exiting the dbUpdateRole production.
	ExitDbUpdateRole(c *DbUpdateRoleContext)

	// ExitDbUpdateUser is called when exiting the dbUpdateUser production.
	ExitDbUpdateUser(c *DbUpdateUserContext)

	// ExitDbWatch is called when exiting the dbWatch production.
	ExitDbWatch(c *DbWatchContext)

	// ExitCollectionOperation is called when exiting the collectionOperation production.
	ExitCollectionOperation(c *CollectionOperationContext)

	// ExitBulkStatement is called when exiting the bulkStatement production.
	ExitBulkStatement(c *BulkStatementContext)

	// ExitBulkInitMethod is called when exiting the bulkInitMethod production.
	ExitBulkInitMethod(c *BulkInitMethodContext)

	// ExitBulkMethodChain is called when exiting the bulkMethodChain production.
	ExitBulkMethodChain(c *BulkMethodChainContext)

	// ExitBulkFind is called when exiting the bulkFind production.
	ExitBulkFind(c *BulkFindContext)

	// ExitBulkInsert is called when exiting the bulkInsert production.
	ExitBulkInsert(c *BulkInsertContext)

	// ExitBulkRemove is called when exiting the bulkRemove production.
	ExitBulkRemove(c *BulkRemoveContext)

	// ExitBulkExecute is called when exiting the bulkExecute production.
	ExitBulkExecute(c *BulkExecuteContext)

	// ExitBulkGetOperations is called when exiting the bulkGetOperations production.
	ExitBulkGetOperations(c *BulkGetOperationsContext)

	// ExitBulkToString is called when exiting the bulkToString production.
	ExitBulkToString(c *BulkToStringContext)

	// ExitBulkGenericMethod is called when exiting the bulkGenericMethod production.
	ExitBulkGenericMethod(c *BulkGenericMethodContext)

	// ExitMongoConnection is called when exiting the mongoConnection production.
	ExitMongoConnection(c *MongoConnectionContext)

	// ExitConnectCall is called when exiting the connectCall production.
	ExitConnectCall(c *ConnectCallContext)

	// ExitDbGetMongoChain is called when exiting the dbGetMongoChain production.
	ExitDbGetMongoChain(c *DbGetMongoChainContext)

	// ExitConnectionMethodChain is called when exiting the connectionMethodChain production.
	ExitConnectionMethodChain(c *ConnectionMethodChainContext)

	// ExitRsStatement is called when exiting the rsStatement production.
	ExitRsStatement(c *RsStatementContext)

	// ExitShStatement is called when exiting the shStatement production.
	ExitShStatement(c *ShStatementContext)

	// ExitKeyVaultStatement is called when exiting the keyVaultStatement production.
	ExitKeyVaultStatement(c *KeyVaultStatementContext)

	// ExitClientEncryptionStatement is called when exiting the clientEncryptionStatement production.
	ExitClientEncryptionStatement(c *ClientEncryptionStatementContext)

	// ExitPlanCacheStatement is called when exiting the planCacheStatement production.
	ExitPlanCacheStatement(c *PlanCacheStatementContext)

	// ExitSpStatement is called when exiting the spStatement production.
	ExitSpStatement(c *SpStatementContext)

	// ExitNativeFunctionCall is called when exiting the nativeFunctionCall production.
	ExitNativeFunctionCall(c *NativeFunctionCallContext)

	// ExitConnGetDB is called when exiting the connGetDB production.
	ExitConnGetDB(c *ConnGetDBContext)

	// ExitConnGetReadConcern is called when exiting the connGetReadConcern production.
	ExitConnGetReadConcern(c *ConnGetReadConcernContext)

	// ExitConnGetReadPref is called when exiting the connGetReadPref production.
	ExitConnGetReadPref(c *ConnGetReadPrefContext)

	// ExitConnGetReadPrefMode is called when exiting the connGetReadPrefMode production.
	ExitConnGetReadPrefMode(c *ConnGetReadPrefModeContext)

	// ExitConnGetReadPrefTagSet is called when exiting the connGetReadPrefTagSet production.
	ExitConnGetReadPrefTagSet(c *ConnGetReadPrefTagSetContext)

	// ExitConnGetWriteConcern is called when exiting the connGetWriteConcern production.
	ExitConnGetWriteConcern(c *ConnGetWriteConcernContext)

	// ExitConnSetReadPref is called when exiting the connSetReadPref production.
	ExitConnSetReadPref(c *ConnSetReadPrefContext)

	// ExitConnSetReadConcern is called when exiting the connSetReadConcern production.
	ExitConnSetReadConcern(c *ConnSetReadConcernContext)

	// ExitConnSetWriteConcern is called when exiting the connSetWriteConcern production.
	ExitConnSetWriteConcern(c *ConnSetWriteConcernContext)

	// ExitConnStartSession is called when exiting the connStartSession production.
	ExitConnStartSession(c *ConnStartSessionContext)

	// ExitConnWatch is called when exiting the connWatch production.
	ExitConnWatch(c *ConnWatchContext)

	// ExitConnClose is called when exiting the connClose production.
	ExitConnClose(c *ConnCloseContext)

	// ExitConnAdminCommand is called when exiting the connAdminCommand production.
	ExitConnAdminCommand(c *ConnAdminCommandContext)

	// ExitConnGetDBNames is called when exiting the connGetDBNames production.
	ExitConnGetDBNames(c *ConnGetDBNamesContext)

	// ExitConnGenericMethod is called when exiting the connGenericMethod production.
	ExitConnGenericMethod(c *ConnGenericMethodContext)

	// ExitDotAccess is called when exiting the dotAccess production.
	ExitDotAccess(c *DotAccessContext)

	// ExitBracketAccess is called when exiting the bracketAccess production.
	ExitBracketAccess(c *BracketAccessContext)

	// ExitGetCollectionAccess is called when exiting the getCollectionAccess production.
	ExitGetCollectionAccess(c *GetCollectionAccessContext)

	// ExitMethodChain is called when exiting the methodChain production.
	ExitMethodChain(c *MethodChainContext)

	// ExitCollectionMethodCall is called when exiting the collectionMethodCall production.
	ExitCollectionMethodCall(c *CollectionMethodCallContext)

	// ExitCursorMethodCall is called when exiting the cursorMethodCall production.
	ExitCursorMethodCall(c *CursorMethodCallContext)

	// ExitFindMethod is called when exiting the findMethod production.
	ExitFindMethod(c *FindMethodContext)

	// ExitFindOneMethod is called when exiting the findOneMethod production.
	ExitFindOneMethod(c *FindOneMethodContext)

	// ExitCountDocumentsMethod is called when exiting the countDocumentsMethod production.
	ExitCountDocumentsMethod(c *CountDocumentsMethodContext)

	// ExitEstimatedDocumentCountMethod is called when exiting the estimatedDocumentCountMethod production.
	ExitEstimatedDocumentCountMethod(c *EstimatedDocumentCountMethodContext)

	// ExitDistinctMethod is called when exiting the distinctMethod production.
	ExitDistinctMethod(c *DistinctMethodContext)

	// ExitAggregateMethod is called when exiting the aggregateMethod production.
	ExitAggregateMethod(c *AggregateMethodContext)

	// ExitGetIndexesMethod is called when exiting the getIndexesMethod production.
	ExitGetIndexesMethod(c *GetIndexesMethodContext)

	// ExitInsertOneMethod is called when exiting the insertOneMethod production.
	ExitInsertOneMethod(c *InsertOneMethodContext)

	// ExitInsertManyMethod is called when exiting the insertManyMethod production.
	ExitInsertManyMethod(c *InsertManyMethodContext)

	// ExitUpdateOneMethod is called when exiting the updateOneMethod production.
	ExitUpdateOneMethod(c *UpdateOneMethodContext)

	// ExitUpdateManyMethod is called when exiting the updateManyMethod production.
	ExitUpdateManyMethod(c *UpdateManyMethodContext)

	// ExitDeleteOneMethod is called when exiting the deleteOneMethod production.
	ExitDeleteOneMethod(c *DeleteOneMethodContext)

	// ExitDeleteManyMethod is called when exiting the deleteManyMethod production.
	ExitDeleteManyMethod(c *DeleteManyMethodContext)

	// ExitReplaceOneMethod is called when exiting the replaceOneMethod production.
	ExitReplaceOneMethod(c *ReplaceOneMethodContext)

	// ExitFindOneAndUpdateMethod is called when exiting the findOneAndUpdateMethod production.
	ExitFindOneAndUpdateMethod(c *FindOneAndUpdateMethodContext)

	// ExitFindOneAndReplaceMethod is called when exiting the findOneAndReplaceMethod production.
	ExitFindOneAndReplaceMethod(c *FindOneAndReplaceMethodContext)

	// ExitFindOneAndDeleteMethod is called when exiting the findOneAndDeleteMethod production.
	ExitFindOneAndDeleteMethod(c *FindOneAndDeleteMethodContext)

	// ExitCreateIndexMethod is called when exiting the createIndexMethod production.
	ExitCreateIndexMethod(c *CreateIndexMethodContext)

	// ExitCreateIndexesMethod is called when exiting the createIndexesMethod production.
	ExitCreateIndexesMethod(c *CreateIndexesMethodContext)

	// ExitDropIndexMethod is called when exiting the dropIndexMethod production.
	ExitDropIndexMethod(c *DropIndexMethodContext)

	// ExitDropIndexesMethod is called when exiting the dropIndexesMethod production.
	ExitDropIndexesMethod(c *DropIndexesMethodContext)

	// ExitDropMethod is called when exiting the dropMethod production.
	ExitDropMethod(c *DropMethodContext)

	// ExitRenameCollectionMethod is called when exiting the renameCollectionMethod production.
	ExitRenameCollectionMethod(c *RenameCollectionMethodContext)

	// ExitStatsMethod is called when exiting the statsMethod production.
	ExitStatsMethod(c *StatsMethodContext)

	// ExitStorageSizeMethod is called when exiting the storageSizeMethod production.
	ExitStorageSizeMethod(c *StorageSizeMethodContext)

	// ExitTotalIndexSizeMethod is called when exiting the totalIndexSizeMethod production.
	ExitTotalIndexSizeMethod(c *TotalIndexSizeMethodContext)

	// ExitTotalSizeMethod is called when exiting the totalSizeMethod production.
	ExitTotalSizeMethod(c *TotalSizeMethodContext)

	// ExitDataSizeMethod is called when exiting the dataSizeMethod production.
	ExitDataSizeMethod(c *DataSizeMethodContext)

	// ExitIsCappedMethod is called when exiting the isCappedMethod production.
	ExitIsCappedMethod(c *IsCappedMethodContext)

	// ExitValidateMethod is called when exiting the validateMethod production.
	ExitValidateMethod(c *ValidateMethodContext)

	// ExitLatencyStatsMethod is called when exiting the latencyStatsMethod production.
	ExitLatencyStatsMethod(c *LatencyStatsMethodContext)

	// ExitWatchMethod is called when exiting the watchMethod production.
	ExitWatchMethod(c *WatchMethodContext)

	// ExitBulkWriteMethod is called when exiting the bulkWriteMethod production.
	ExitBulkWriteMethod(c *BulkWriteMethodContext)

	// ExitCollectionCountMethod is called when exiting the collectionCountMethod production.
	ExitCollectionCountMethod(c *CollectionCountMethodContext)

	// ExitCollectionInsertMethod is called when exiting the collectionInsertMethod production.
	ExitCollectionInsertMethod(c *CollectionInsertMethodContext)

	// ExitCollectionRemoveMethod is called when exiting the collectionRemoveMethod production.
	ExitCollectionRemoveMethod(c *CollectionRemoveMethodContext)

	// ExitUpdateMethod is called when exiting the updateMethod production.
	ExitUpdateMethod(c *UpdateMethodContext)

	// ExitMapReduceMethod is called when exiting the mapReduceMethod production.
	ExitMapReduceMethod(c *MapReduceMethodContext)

	// ExitFindAndModifyMethod is called when exiting the findAndModifyMethod production.
	ExitFindAndModifyMethod(c *FindAndModifyMethodContext)

	// ExitCollectionExplainMethod is called when exiting the collectionExplainMethod production.
	ExitCollectionExplainMethod(c *CollectionExplainMethodContext)

	// ExitAnalyzeShardKeyMethod is called when exiting the analyzeShardKeyMethod production.
	ExitAnalyzeShardKeyMethod(c *AnalyzeShardKeyMethodContext)

	// ExitConfigureQueryAnalyzerMethod is called when exiting the configureQueryAnalyzerMethod production.
	ExitConfigureQueryAnalyzerMethod(c *ConfigureQueryAnalyzerMethodContext)

	// ExitCompactStructuredEncryptionDataMethod is called when exiting the compactStructuredEncryptionDataMethod production.
	ExitCompactStructuredEncryptionDataMethod(c *CompactStructuredEncryptionDataMethodContext)

	// ExitHideIndexMethod is called when exiting the hideIndexMethod production.
	ExitHideIndexMethod(c *HideIndexMethodContext)

	// ExitUnhideIndexMethod is called when exiting the unhideIndexMethod production.
	ExitUnhideIndexMethod(c *UnhideIndexMethodContext)

	// ExitReIndexMethod is called when exiting the reIndexMethod production.
	ExitReIndexMethod(c *ReIndexMethodContext)

	// ExitGetShardDistributionMethod is called when exiting the getShardDistributionMethod production.
	ExitGetShardDistributionMethod(c *GetShardDistributionMethodContext)

	// ExitGetShardVersionMethod is called when exiting the getShardVersionMethod production.
	ExitGetShardVersionMethod(c *GetShardVersionMethodContext)

	// ExitCreateSearchIndexMethod is called when exiting the createSearchIndexMethod production.
	ExitCreateSearchIndexMethod(c *CreateSearchIndexMethodContext)

	// ExitCreateSearchIndexesMethod is called when exiting the createSearchIndexesMethod production.
	ExitCreateSearchIndexesMethod(c *CreateSearchIndexesMethodContext)

	// ExitDropSearchIndexMethod is called when exiting the dropSearchIndexMethod production.
	ExitDropSearchIndexMethod(c *DropSearchIndexMethodContext)

	// ExitUpdateSearchIndexMethod is called when exiting the updateSearchIndexMethod production.
	ExitUpdateSearchIndexMethod(c *UpdateSearchIndexMethodContext)

	// ExitSortMethod is called when exiting the sortMethod production.
	ExitSortMethod(c *SortMethodContext)

	// ExitLimitMethod is called when exiting the limitMethod production.
	ExitLimitMethod(c *LimitMethodContext)

	// ExitSkipMethod is called when exiting the skipMethod production.
	ExitSkipMethod(c *SkipMethodContext)

	// ExitCountMethod is called when exiting the countMethod production.
	ExitCountMethod(c *CountMethodContext)

	// ExitProjectionMethod is called when exiting the projectionMethod production.
	ExitProjectionMethod(c *ProjectionMethodContext)

	// ExitBatchSizeMethod is called when exiting the batchSizeMethod production.
	ExitBatchSizeMethod(c *BatchSizeMethodContext)

	// ExitCloseMethod is called when exiting the closeMethod production.
	ExitCloseMethod(c *CloseMethodContext)

	// ExitCollationMethod is called when exiting the collationMethod production.
	ExitCollationMethod(c *CollationMethodContext)

	// ExitCommentMethod is called when exiting the commentMethod production.
	ExitCommentMethod(c *CommentMethodContext)

	// ExitExplainMethod is called when exiting the explainMethod production.
	ExitExplainMethod(c *ExplainMethodContext)

	// ExitForEachMethod is called when exiting the forEachMethod production.
	ExitForEachMethod(c *ForEachMethodContext)

	// ExitHasNextMethod is called when exiting the hasNextMethod production.
	ExitHasNextMethod(c *HasNextMethodContext)

	// ExitHintMethod is called when exiting the hintMethod production.
	ExitHintMethod(c *HintMethodContext)

	// ExitIsClosedMethod is called when exiting the isClosedMethod production.
	ExitIsClosedMethod(c *IsClosedMethodContext)

	// ExitIsExhaustedMethod is called when exiting the isExhaustedMethod production.
	ExitIsExhaustedMethod(c *IsExhaustedMethodContext)

	// ExitItcountMethod is called when exiting the itcountMethod production.
	ExitItcountMethod(c *ItcountMethodContext)

	// ExitMapMethod is called when exiting the mapMethod production.
	ExitMapMethod(c *MapMethodContext)

	// ExitMaxMethod is called when exiting the maxMethod production.
	ExitMaxMethod(c *MaxMethodContext)

	// ExitMaxAwaitTimeMSMethod is called when exiting the maxAwaitTimeMSMethod production.
	ExitMaxAwaitTimeMSMethod(c *MaxAwaitTimeMSMethodContext)

	// ExitMaxTimeMSMethod is called when exiting the maxTimeMSMethod production.
	ExitMaxTimeMSMethod(c *MaxTimeMSMethodContext)

	// ExitMinMethod is called when exiting the minMethod production.
	ExitMinMethod(c *MinMethodContext)

	// ExitNextMethod is called when exiting the nextMethod production.
	ExitNextMethod(c *NextMethodContext)

	// ExitNoCursorTimeoutMethod is called when exiting the noCursorTimeoutMethod production.
	ExitNoCursorTimeoutMethod(c *NoCursorTimeoutMethodContext)

	// ExitObjsLeftInBatchMethod is called when exiting the objsLeftInBatchMethod production.
	ExitObjsLeftInBatchMethod(c *ObjsLeftInBatchMethodContext)

	// ExitPrettyMethod is called when exiting the prettyMethod production.
	ExitPrettyMethod(c *PrettyMethodContext)

	// ExitReadConcernMethod is called when exiting the readConcernMethod production.
	ExitReadConcernMethod(c *ReadConcernMethodContext)

	// ExitReadPrefMethod is called when exiting the readPrefMethod production.
	ExitReadPrefMethod(c *ReadPrefMethodContext)

	// ExitReturnKeyMethod is called when exiting the returnKeyMethod production.
	ExitReturnKeyMethod(c *ReturnKeyMethodContext)

	// ExitShowRecordIdMethod is called when exiting the showRecordIdMethod production.
	ExitShowRecordIdMethod(c *ShowRecordIdMethodContext)

	// ExitSizeMethod is called when exiting the sizeMethod production.
	ExitSizeMethod(c *SizeMethodContext)

	// ExitTailableMethod is called when exiting the tailableMethod production.
	ExitTailableMethod(c *TailableMethodContext)

	// ExitToArrayMethod is called when exiting the toArrayMethod production.
	ExitToArrayMethod(c *ToArrayMethodContext)

	// ExitTryNextMethod is called when exiting the tryNextMethod production.
	ExitTryNextMethod(c *TryNextMethodContext)

	// ExitAllowDiskUseMethod is called when exiting the allowDiskUseMethod production.
	ExitAllowDiskUseMethod(c *AllowDiskUseMethodContext)

	// ExitAddOptionMethod is called when exiting the addOptionMethod production.
	ExitAddOptionMethod(c *AddOptionMethodContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitUnquotedKey is called when exiting the unquotedKey production.
	ExitUnquotedKey(c *UnquotedKeyContext)

	// ExitQuotedKey is called when exiting the quotedKey production.
	ExitQuotedKey(c *QuotedKeyContext)

	// ExitDocumentValue is called when exiting the documentValue production.
	ExitDocumentValue(c *DocumentValueContext)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitHelperValue is called when exiting the helperValue production.
	ExitHelperValue(c *HelperValueContext)

	// ExitRegexLiteralValue is called when exiting the regexLiteralValue production.
	ExitRegexLiteralValue(c *RegexLiteralValueContext)

	// ExitRegexpConstructorValue is called when exiting the regexpConstructorValue production.
	ExitRegexpConstructorValue(c *RegexpConstructorValueContext)

	// ExitLiteralValue is called when exiting the literalValue production.
	ExitLiteralValue(c *LiteralValueContext)

	// ExitNewKeywordValue is called when exiting the newKeywordValue production.
	ExitNewKeywordValue(c *NewKeywordValueContext)

	// ExitNewKeywordError is called when exiting the newKeywordError production.
	ExitNewKeywordError(c *NewKeywordErrorContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitHelperFunction is called when exiting the helperFunction production.
	ExitHelperFunction(c *HelperFunctionContext)

	// ExitObjectIdHelper is called when exiting the objectIdHelper production.
	ExitObjectIdHelper(c *ObjectIdHelperContext)

	// ExitIsoDateHelper is called when exiting the isoDateHelper production.
	ExitIsoDateHelper(c *IsoDateHelperContext)

	// ExitDateHelper is called when exiting the dateHelper production.
	ExitDateHelper(c *DateHelperContext)

	// ExitUuidHelper is called when exiting the uuidHelper production.
	ExitUuidHelper(c *UuidHelperContext)

	// ExitLongHelper is called when exiting the longHelper production.
	ExitLongHelper(c *LongHelperContext)

	// ExitInt32Helper is called when exiting the int32Helper production.
	ExitInt32Helper(c *Int32HelperContext)

	// ExitDoubleHelper is called when exiting the doubleHelper production.
	ExitDoubleHelper(c *DoubleHelperContext)

	// ExitDecimal128Helper is called when exiting the decimal128Helper production.
	ExitDecimal128Helper(c *Decimal128HelperContext)

	// ExitTimestampDocHelper is called when exiting the timestampDocHelper production.
	ExitTimestampDocHelper(c *TimestampDocHelperContext)

	// ExitTimestampArgsHelper is called when exiting the timestampArgsHelper production.
	ExitTimestampArgsHelper(c *TimestampArgsHelperContext)

	// ExitRegExpConstructor is called when exiting the regExpConstructor production.
	ExitRegExpConstructor(c *RegExpConstructorContext)

	// ExitBinDataHelper is called when exiting the binDataHelper production.
	ExitBinDataHelper(c *BinDataHelperContext)

	// ExitBinaryHelper is called when exiting the binaryHelper production.
	ExitBinaryHelper(c *BinaryHelperContext)

	// ExitBsonRegExpHelper is called when exiting the bsonRegExpHelper production.
	ExitBsonRegExpHelper(c *BsonRegExpHelperContext)

	// ExitHexDataHelper is called when exiting the hexDataHelper production.
	ExitHexDataHelper(c *HexDataHelperContext)

	// ExitStringLiteralValue is called when exiting the stringLiteralValue production.
	ExitStringLiteralValue(c *StringLiteralValueContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitTrueLiteral is called when exiting the trueLiteral production.
	ExitTrueLiteral(c *TrueLiteralContext)

	// ExitFalseLiteral is called when exiting the falseLiteral production.
	ExitFalseLiteral(c *FalseLiteralContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
