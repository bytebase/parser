// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MongoShellParser.
type MongoShellParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MongoShellParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MongoShellParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#showDatabases.
	VisitShowDatabases(ctx *ShowDatabasesContext) interface{}

	// Visit a parse tree produced by MongoShellParser#showCollections.
	VisitShowCollections(ctx *ShowCollectionsContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getCollectionNames.
	VisitGetCollectionNames(ctx *GetCollectionNamesContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getCollectionInfos.
	VisitGetCollectionInfos(ctx *GetCollectionInfosContext) interface{}

	// Visit a parse tree produced by MongoShellParser#createCollection.
	VisitCreateCollection(ctx *CreateCollectionContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dbStats.
	VisitDbStats(ctx *DbStatsContext) interface{}

	// Visit a parse tree produced by MongoShellParser#serverStatus.
	VisitServerStatus(ctx *ServerStatusContext) interface{}

	// Visit a parse tree produced by MongoShellParser#serverBuildInfo.
	VisitServerBuildInfo(ctx *ServerBuildInfoContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dbVersion.
	VisitDbVersion(ctx *DbVersionContext) interface{}

	// Visit a parse tree produced by MongoShellParser#hostInfo.
	VisitHostInfo(ctx *HostInfoContext) interface{}

	// Visit a parse tree produced by MongoShellParser#listCommands.
	VisitListCommands(ctx *ListCommandsContext) interface{}

	// Visit a parse tree produced by MongoShellParser#runCommand.
	VisitRunCommand(ctx *RunCommandContext) interface{}

	// Visit a parse tree produced by MongoShellParser#adminCommand.
	VisitAdminCommand(ctx *AdminCommandContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getName.
	VisitGetName(ctx *GetNameContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getMongo.
	VisitGetMongo(ctx *GetMongoContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getSiblingDB.
	VisitGetSiblingDB(ctx *GetSiblingDBContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dbGenericMethod.
	VisitDbGenericMethod(ctx *DbGenericMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#collectionOperation.
	VisitCollectionOperation(ctx *CollectionOperationContext) interface{}

	// Visit a parse tree produced by MongoShellParser#genericDbMethod.
	VisitGenericDbMethod(ctx *GenericDbMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkStatement.
	VisitBulkStatement(ctx *BulkStatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkInitMethod.
	VisitBulkInitMethod(ctx *BulkInitMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkMethodChain.
	VisitBulkMethodChain(ctx *BulkMethodChainContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkFind.
	VisitBulkFind(ctx *BulkFindContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkInsert.
	VisitBulkInsert(ctx *BulkInsertContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkRemove.
	VisitBulkRemove(ctx *BulkRemoveContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkExecute.
	VisitBulkExecute(ctx *BulkExecuteContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkGetOperations.
	VisitBulkGetOperations(ctx *BulkGetOperationsContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkToString.
	VisitBulkToString(ctx *BulkToStringContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bulkGenericMethod.
	VisitBulkGenericMethod(ctx *BulkGenericMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#mongoConnection.
	VisitMongoConnection(ctx *MongoConnectionContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connectCall.
	VisitConnectCall(ctx *ConnectCallContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dbGetMongoChain.
	VisitDbGetMongoChain(ctx *DbGetMongoChainContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connectionMethodChain.
	VisitConnectionMethodChain(ctx *ConnectionMethodChainContext) interface{}

	// Visit a parse tree produced by MongoShellParser#rsStatement.
	VisitRsStatement(ctx *RsStatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#shStatement.
	VisitShStatement(ctx *ShStatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#keyVaultStatement.
	VisitKeyVaultStatement(ctx *KeyVaultStatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#clientEncryptionStatement.
	VisitClientEncryptionStatement(ctx *ClientEncryptionStatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#planCacheStatement.
	VisitPlanCacheStatement(ctx *PlanCacheStatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#spStatement.
	VisitSpStatement(ctx *SpStatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#nativeFunctionCall.
	VisitNativeFunctionCall(ctx *NativeFunctionCallContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGetDB.
	VisitConnGetDB(ctx *ConnGetDBContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGetReadConcern.
	VisitConnGetReadConcern(ctx *ConnGetReadConcernContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGetReadPref.
	VisitConnGetReadPref(ctx *ConnGetReadPrefContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGetReadPrefMode.
	VisitConnGetReadPrefMode(ctx *ConnGetReadPrefModeContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGetReadPrefTagSet.
	VisitConnGetReadPrefTagSet(ctx *ConnGetReadPrefTagSetContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGetWriteConcern.
	VisitConnGetWriteConcern(ctx *ConnGetWriteConcernContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connSetReadPref.
	VisitConnSetReadPref(ctx *ConnSetReadPrefContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connSetReadConcern.
	VisitConnSetReadConcern(ctx *ConnSetReadConcernContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connSetWriteConcern.
	VisitConnSetWriteConcern(ctx *ConnSetWriteConcernContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connStartSession.
	VisitConnStartSession(ctx *ConnStartSessionContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connWatch.
	VisitConnWatch(ctx *ConnWatchContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connClose.
	VisitConnClose(ctx *ConnCloseContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connAdminCommand.
	VisitConnAdminCommand(ctx *ConnAdminCommandContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGetDBNames.
	VisitConnGetDBNames(ctx *ConnGetDBNamesContext) interface{}

	// Visit a parse tree produced by MongoShellParser#connGenericMethod.
	VisitConnGenericMethod(ctx *ConnGenericMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dotAccess.
	VisitDotAccess(ctx *DotAccessContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bracketAccess.
	VisitBracketAccess(ctx *BracketAccessContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getCollectionAccess.
	VisitGetCollectionAccess(ctx *GetCollectionAccessContext) interface{}

	// Visit a parse tree produced by MongoShellParser#methodChain.
	VisitMethodChain(ctx *MethodChainContext) interface{}

	// Visit a parse tree produced by MongoShellParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by MongoShellParser#findMethod.
	VisitFindMethod(ctx *FindMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#findOneMethod.
	VisitFindOneMethod(ctx *FindOneMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#countDocumentsMethod.
	VisitCountDocumentsMethod(ctx *CountDocumentsMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#estimatedDocumentCountMethod.
	VisitEstimatedDocumentCountMethod(ctx *EstimatedDocumentCountMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#distinctMethod.
	VisitDistinctMethod(ctx *DistinctMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#aggregateMethod.
	VisitAggregateMethod(ctx *AggregateMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getIndexesMethod.
	VisitGetIndexesMethod(ctx *GetIndexesMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#insertOneMethod.
	VisitInsertOneMethod(ctx *InsertOneMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#insertManyMethod.
	VisitInsertManyMethod(ctx *InsertManyMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#updateOneMethod.
	VisitUpdateOneMethod(ctx *UpdateOneMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#updateManyMethod.
	VisitUpdateManyMethod(ctx *UpdateManyMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#deleteOneMethod.
	VisitDeleteOneMethod(ctx *DeleteOneMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#deleteManyMethod.
	VisitDeleteManyMethod(ctx *DeleteManyMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#replaceOneMethod.
	VisitReplaceOneMethod(ctx *ReplaceOneMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#findOneAndUpdateMethod.
	VisitFindOneAndUpdateMethod(ctx *FindOneAndUpdateMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#findOneAndReplaceMethod.
	VisitFindOneAndReplaceMethod(ctx *FindOneAndReplaceMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#findOneAndDeleteMethod.
	VisitFindOneAndDeleteMethod(ctx *FindOneAndDeleteMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#createIndexMethod.
	VisitCreateIndexMethod(ctx *CreateIndexMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#createIndexesMethod.
	VisitCreateIndexesMethod(ctx *CreateIndexesMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dropIndexMethod.
	VisitDropIndexMethod(ctx *DropIndexMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dropIndexesMethod.
	VisitDropIndexesMethod(ctx *DropIndexesMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dropMethod.
	VisitDropMethod(ctx *DropMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#renameCollectionMethod.
	VisitRenameCollectionMethod(ctx *RenameCollectionMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#statsMethod.
	VisitStatsMethod(ctx *StatsMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#storageSizeMethod.
	VisitStorageSizeMethod(ctx *StorageSizeMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#totalIndexSizeMethod.
	VisitTotalIndexSizeMethod(ctx *TotalIndexSizeMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#totalSizeMethod.
	VisitTotalSizeMethod(ctx *TotalSizeMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dataSizeMethod.
	VisitDataSizeMethod(ctx *DataSizeMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#isCappedMethod.
	VisitIsCappedMethod(ctx *IsCappedMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#validateMethod.
	VisitValidateMethod(ctx *ValidateMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#latencyStatsMethod.
	VisitLatencyStatsMethod(ctx *LatencyStatsMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#sortMethod.
	VisitSortMethod(ctx *SortMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#limitMethod.
	VisitLimitMethod(ctx *LimitMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#skipMethod.
	VisitSkipMethod(ctx *SkipMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#countMethod.
	VisitCountMethod(ctx *CountMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#projectionMethod.
	VisitProjectionMethod(ctx *ProjectionMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#batchSizeMethod.
	VisitBatchSizeMethod(ctx *BatchSizeMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#closeMethod.
	VisitCloseMethod(ctx *CloseMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#collationMethod.
	VisitCollationMethod(ctx *CollationMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#commentMethod.
	VisitCommentMethod(ctx *CommentMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#explainMethod.
	VisitExplainMethod(ctx *ExplainMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#forEachMethod.
	VisitForEachMethod(ctx *ForEachMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#hasNextMethod.
	VisitHasNextMethod(ctx *HasNextMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#hintMethod.
	VisitHintMethod(ctx *HintMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#isClosedMethod.
	VisitIsClosedMethod(ctx *IsClosedMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#isExhaustedMethod.
	VisitIsExhaustedMethod(ctx *IsExhaustedMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#itcountMethod.
	VisitItcountMethod(ctx *ItcountMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#mapMethod.
	VisitMapMethod(ctx *MapMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#maxMethod.
	VisitMaxMethod(ctx *MaxMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#maxAwaitTimeMSMethod.
	VisitMaxAwaitTimeMSMethod(ctx *MaxAwaitTimeMSMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#maxTimeMSMethod.
	VisitMaxTimeMSMethod(ctx *MaxTimeMSMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#minMethod.
	VisitMinMethod(ctx *MinMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#nextMethod.
	VisitNextMethod(ctx *NextMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#noCursorTimeoutMethod.
	VisitNoCursorTimeoutMethod(ctx *NoCursorTimeoutMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#objsLeftInBatchMethod.
	VisitObjsLeftInBatchMethod(ctx *ObjsLeftInBatchMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#prettyMethod.
	VisitPrettyMethod(ctx *PrettyMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#readConcernMethod.
	VisitReadConcernMethod(ctx *ReadConcernMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#readPrefMethod.
	VisitReadPrefMethod(ctx *ReadPrefMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#returnKeyMethod.
	VisitReturnKeyMethod(ctx *ReturnKeyMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#showRecordIdMethod.
	VisitShowRecordIdMethod(ctx *ShowRecordIdMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#sizeMethod.
	VisitSizeMethod(ctx *SizeMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#tailableMethod.
	VisitTailableMethod(ctx *TailableMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#toArrayMethod.
	VisitToArrayMethod(ctx *ToArrayMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#tryNextMethod.
	VisitTryNextMethod(ctx *TryNextMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#allowDiskUseMethod.
	VisitAllowDiskUseMethod(ctx *AllowDiskUseMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#addOptionMethod.
	VisitAddOptionMethod(ctx *AddOptionMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#genericMethod.
	VisitGenericMethod(ctx *GenericMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by MongoShellParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by MongoShellParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by MongoShellParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by MongoShellParser#unquotedKey.
	VisitUnquotedKey(ctx *UnquotedKeyContext) interface{}

	// Visit a parse tree produced by MongoShellParser#quotedKey.
	VisitQuotedKey(ctx *QuotedKeyContext) interface{}

	// Visit a parse tree produced by MongoShellParser#documentValue.
	VisitDocumentValue(ctx *DocumentValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#helperValue.
	VisitHelperValue(ctx *HelperValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#regexLiteralValue.
	VisitRegexLiteralValue(ctx *RegexLiteralValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#regexpConstructorValue.
	VisitRegexpConstructorValue(ctx *RegexpConstructorValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#literalValue.
	VisitLiteralValue(ctx *LiteralValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#newKeywordValue.
	VisitNewKeywordValue(ctx *NewKeywordValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#newKeywordError.
	VisitNewKeywordError(ctx *NewKeywordErrorContext) interface{}

	// Visit a parse tree produced by MongoShellParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by MongoShellParser#helperFunction.
	VisitHelperFunction(ctx *HelperFunctionContext) interface{}

	// Visit a parse tree produced by MongoShellParser#objectIdHelper.
	VisitObjectIdHelper(ctx *ObjectIdHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#isoDateHelper.
	VisitIsoDateHelper(ctx *IsoDateHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dateHelper.
	VisitDateHelper(ctx *DateHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#uuidHelper.
	VisitUuidHelper(ctx *UuidHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#longHelper.
	VisitLongHelper(ctx *LongHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#int32Helper.
	VisitInt32Helper(ctx *Int32HelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#doubleHelper.
	VisitDoubleHelper(ctx *DoubleHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#decimal128Helper.
	VisitDecimal128Helper(ctx *Decimal128HelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#timestampDocHelper.
	VisitTimestampDocHelper(ctx *TimestampDocHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#timestampArgsHelper.
	VisitTimestampArgsHelper(ctx *TimestampArgsHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#regExpConstructor.
	VisitRegExpConstructor(ctx *RegExpConstructorContext) interface{}

	// Visit a parse tree produced by MongoShellParser#binDataHelper.
	VisitBinDataHelper(ctx *BinDataHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#binaryHelper.
	VisitBinaryHelper(ctx *BinaryHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bsonRegExpHelper.
	VisitBsonRegExpHelper(ctx *BsonRegExpHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#hexDataHelper.
	VisitHexDataHelper(ctx *HexDataHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#stringLiteralValue.
	VisitStringLiteralValue(ctx *StringLiteralValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#trueLiteral.
	VisitTrueLiteral(ctx *TrueLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#falseLiteral.
	VisitFalseLiteral(ctx *FalseLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
