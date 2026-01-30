// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

type BaseMongoShellParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMongoShellParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitShowDatabases(ctx *ShowDatabasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitShowCollections(ctx *ShowCollectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetCollectionNames(ctx *GetCollectionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetCollectionInfos(ctx *GetCollectionInfosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCreateCollection(ctx *CreateCollectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDropDatabase(ctx *DropDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDbStats(ctx *DbStatsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitServerStatus(ctx *ServerStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitServerBuildInfo(ctx *ServerBuildInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDbVersion(ctx *DbVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHostInfo(ctx *HostInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitListCommands(ctx *ListCommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRunCommand(ctx *RunCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitAdminCommand(ctx *AdminCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetName(ctx *GetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetMongo(ctx *GetMongoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetSiblingDB(ctx *GetSiblingDBContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDbGenericMethod(ctx *DbGenericMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollectionOperation(ctx *CollectionOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGenericDbMethod(ctx *GenericDbMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkStatement(ctx *BulkStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkInitMethod(ctx *BulkInitMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkMethodChain(ctx *BulkMethodChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkFind(ctx *BulkFindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkInsert(ctx *BulkInsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkRemove(ctx *BulkRemoveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkExecute(ctx *BulkExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkGetOperations(ctx *BulkGetOperationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkToString(ctx *BulkToStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkGenericMethod(ctx *BulkGenericMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMongoConnection(ctx *MongoConnectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnectCall(ctx *ConnectCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDbGetMongoChain(ctx *DbGetMongoChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnectionMethodChain(ctx *ConnectionMethodChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRsStatement(ctx *RsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitShStatement(ctx *ShStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitKeyVaultStatement(ctx *KeyVaultStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitClientEncryptionStatement(ctx *ClientEncryptionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitPlanCacheStatement(ctx *PlanCacheStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitSpStatement(ctx *SpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNativeFunctionCall(ctx *NativeFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGetDB(ctx *ConnGetDBContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGetReadConcern(ctx *ConnGetReadConcernContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGetReadPref(ctx *ConnGetReadPrefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGetReadPrefMode(ctx *ConnGetReadPrefModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGetReadPrefTagSet(ctx *ConnGetReadPrefTagSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGetWriteConcern(ctx *ConnGetWriteConcernContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnSetReadPref(ctx *ConnSetReadPrefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnSetReadConcern(ctx *ConnSetReadConcernContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnSetWriteConcern(ctx *ConnSetWriteConcernContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnStartSession(ctx *ConnStartSessionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnWatch(ctx *ConnWatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnClose(ctx *ConnCloseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnAdminCommand(ctx *ConnAdminCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGetDBNames(ctx *ConnGetDBNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConnGenericMethod(ctx *ConnGenericMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDotAccess(ctx *DotAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBracketAccess(ctx *BracketAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetCollectionAccess(ctx *GetCollectionAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMethodChain(ctx *MethodChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollectionMethodCall(ctx *CollectionMethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCursorMethodCall(ctx *CursorMethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindMethod(ctx *FindMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindOneMethod(ctx *FindOneMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCountDocumentsMethod(ctx *CountDocumentsMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitEstimatedDocumentCountMethod(ctx *EstimatedDocumentCountMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDistinctMethod(ctx *DistinctMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitAggregateMethod(ctx *AggregateMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetIndexesMethod(ctx *GetIndexesMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitInsertOneMethod(ctx *InsertOneMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitInsertManyMethod(ctx *InsertManyMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUpdateOneMethod(ctx *UpdateOneMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUpdateManyMethod(ctx *UpdateManyMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDeleteOneMethod(ctx *DeleteOneMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDeleteManyMethod(ctx *DeleteManyMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitReplaceOneMethod(ctx *ReplaceOneMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindOneAndUpdateMethod(ctx *FindOneAndUpdateMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindOneAndReplaceMethod(ctx *FindOneAndReplaceMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindOneAndDeleteMethod(ctx *FindOneAndDeleteMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCreateIndexMethod(ctx *CreateIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCreateIndexesMethod(ctx *CreateIndexesMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDropIndexMethod(ctx *DropIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDropIndexesMethod(ctx *DropIndexesMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDropMethod(ctx *DropMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRenameCollectionMethod(ctx *RenameCollectionMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStatsMethod(ctx *StatsMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStorageSizeMethod(ctx *StorageSizeMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTotalIndexSizeMethod(ctx *TotalIndexSizeMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTotalSizeMethod(ctx *TotalSizeMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDataSizeMethod(ctx *DataSizeMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitIsCappedMethod(ctx *IsCappedMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitValidateMethod(ctx *ValidateMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitLatencyStatsMethod(ctx *LatencyStatsMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitWatchMethod(ctx *WatchMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBulkWriteMethod(ctx *BulkWriteMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollectionCountMethod(ctx *CollectionCountMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollectionInsertMethod(ctx *CollectionInsertMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollectionRemoveMethod(ctx *CollectionRemoveMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUpdateMethod(ctx *UpdateMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMapReduceMethod(ctx *MapReduceMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindAndModifyMethod(ctx *FindAndModifyMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollectionExplainMethod(ctx *CollectionExplainMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitAnalyzeShardKeyMethod(ctx *AnalyzeShardKeyMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitConfigureQueryAnalyzerMethod(ctx *ConfigureQueryAnalyzerMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCompactStructuredEncryptionDataMethod(ctx *CompactStructuredEncryptionDataMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHideIndexMethod(ctx *HideIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUnhideIndexMethod(ctx *UnhideIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitReIndexMethod(ctx *ReIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetShardDistributionMethod(ctx *GetShardDistributionMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetShardVersionMethod(ctx *GetShardVersionMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCreateSearchIndexMethod(ctx *CreateSearchIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCreateSearchIndexesMethod(ctx *CreateSearchIndexesMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDropSearchIndexMethod(ctx *DropSearchIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUpdateSearchIndexMethod(ctx *UpdateSearchIndexMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitSortMethod(ctx *SortMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitLimitMethod(ctx *LimitMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitSkipMethod(ctx *SkipMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCountMethod(ctx *CountMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitProjectionMethod(ctx *ProjectionMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBatchSizeMethod(ctx *BatchSizeMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCloseMethod(ctx *CloseMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollationMethod(ctx *CollationMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCommentMethod(ctx *CommentMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitExplainMethod(ctx *ExplainMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitForEachMethod(ctx *ForEachMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHasNextMethod(ctx *HasNextMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHintMethod(ctx *HintMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitIsClosedMethod(ctx *IsClosedMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitIsExhaustedMethod(ctx *IsExhaustedMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitItcountMethod(ctx *ItcountMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMapMethod(ctx *MapMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMaxMethod(ctx *MaxMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMaxAwaitTimeMSMethod(ctx *MaxAwaitTimeMSMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMaxTimeMSMethod(ctx *MaxTimeMSMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMinMethod(ctx *MinMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNextMethod(ctx *NextMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNoCursorTimeoutMethod(ctx *NoCursorTimeoutMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitObjsLeftInBatchMethod(ctx *ObjsLeftInBatchMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitPrettyMethod(ctx *PrettyMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitReadConcernMethod(ctx *ReadConcernMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitReadPrefMethod(ctx *ReadPrefMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitReturnKeyMethod(ctx *ReturnKeyMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitShowRecordIdMethod(ctx *ShowRecordIdMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitSizeMethod(ctx *SizeMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTailableMethod(ctx *TailableMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitToArrayMethod(ctx *ToArrayMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTryNextMethod(ctx *TryNextMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitAllowDiskUseMethod(ctx *AllowDiskUseMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitAddOptionMethod(ctx *AddOptionMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUnquotedKey(ctx *UnquotedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitQuotedKey(ctx *QuotedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDocumentValue(ctx *DocumentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHelperValue(ctx *HelperValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRegexLiteralValue(ctx *RegexLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRegexpConstructorValue(ctx *RegexpConstructorValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitLiteralValue(ctx *LiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNewKeywordValue(ctx *NewKeywordValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNewKeywordError(ctx *NewKeywordErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHelperFunction(ctx *HelperFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitObjectIdHelper(ctx *ObjectIdHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitIsoDateHelper(ctx *IsoDateHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDateHelper(ctx *DateHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUuidHelper(ctx *UuidHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitLongHelper(ctx *LongHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitInt32Helper(ctx *Int32HelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDoubleHelper(ctx *DoubleHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDecimal128Helper(ctx *Decimal128HelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTimestampDocHelper(ctx *TimestampDocHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTimestampArgsHelper(ctx *TimestampArgsHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRegExpConstructor(ctx *RegExpConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBinDataHelper(ctx *BinDataHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBinaryHelper(ctx *BinaryHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBsonRegExpHelper(ctx *BsonRegExpHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHexDataHelper(ctx *HexDataHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStringLiteralValue(ctx *StringLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTrueLiteral(ctx *TrueLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFalseLiteral(ctx *FalseLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
