// Code generated from MariaDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mariadb // MariaDBParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MariaDBParser.
type MariaDBParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MariaDBParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by MariaDBParser#sqlStatements.
	VisitSqlStatements(ctx *SqlStatementsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#sqlStatement.
	VisitSqlStatement(ctx *SqlStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setStatementFor.
	VisitSetStatementFor(ctx *SetStatementForContext) interface{}

	// Visit a parse tree produced by MariaDBParser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by MariaDBParser#ddlStatement.
	VisitDdlStatement(ctx *DdlStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dmlStatement.
	VisitDmlStatement(ctx *DmlStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#transactionStatement.
	VisitTransactionStatement(ctx *TransactionStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#replicationStatement.
	VisitReplicationStatement(ctx *ReplicationStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#preparedStatement.
	VisitPreparedStatement(ctx *PreparedStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#administrationStatement.
	VisitAdministrationStatement(ctx *AdministrationStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#utilityStatement.
	VisitUtilityStatement(ctx *UtilityStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createEvent.
	VisitCreateEvent(ctx *CreateEventContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createIndex.
	VisitCreateIndex(ctx *CreateIndexContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createLogfileGroup.
	VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createProcedure.
	VisitCreateProcedure(ctx *CreateProcedureContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createFunction.
	VisitCreateFunction(ctx *CreateFunctionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createRole.
	VisitCreateRole(ctx *CreateRoleContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createServer.
	VisitCreateServer(ctx *CreateServerContext) interface{}

	// Visit a parse tree produced by MariaDBParser#copyCreateTable.
	VisitCopyCreateTable(ctx *CopyCreateTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#queryCreateTable.
	VisitQueryCreateTable(ctx *QueryCreateTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#columnCreateTable.
	VisitColumnCreateTable(ctx *ColumnCreateTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createTablespaceInnodb.
	VisitCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createTablespaceNdb.
	VisitCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createTrigger.
	VisitCreateTrigger(ctx *CreateTriggerContext) interface{}

	// Visit a parse tree produced by MariaDBParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#commonTableExpressions.
	VisitCommonTableExpressions(ctx *CommonTableExpressionsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#cteName.
	VisitCteName(ctx *CteNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#cteColumnName.
	VisitCteColumnName(ctx *CteColumnNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createSequence.
	VisitCreateSequence(ctx *CreateSequenceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#sequenceSpec.
	VisitSequenceSpec(ctx *SequenceSpecContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createDatabaseOption.
	VisitCreateDatabaseOption(ctx *CreateDatabaseOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#charSet.
	VisitCharSet(ctx *CharSetContext) interface{}

	// Visit a parse tree produced by MariaDBParser#currentUserExpression.
	VisitCurrentUserExpression(ctx *CurrentUserExpressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#ownerStatement.
	VisitOwnerStatement(ctx *OwnerStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#preciseSchedule.
	VisitPreciseSchedule(ctx *PreciseScheduleContext) interface{}

	// Visit a parse tree produced by MariaDBParser#intervalSchedule.
	VisitIntervalSchedule(ctx *IntervalScheduleContext) interface{}

	// Visit a parse tree produced by MariaDBParser#timestampValue.
	VisitTimestampValue(ctx *TimestampValueContext) interface{}

	// Visit a parse tree produced by MariaDBParser#intervalExpr.
	VisitIntervalExpr(ctx *IntervalExprContext) interface{}

	// Visit a parse tree produced by MariaDBParser#intervalType.
	VisitIntervalType(ctx *IntervalTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#enableType.
	VisitEnableType(ctx *EnableTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#indexOption.
	VisitIndexOption(ctx *IndexOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#procedureParameter.
	VisitProcedureParameter(ctx *ProcedureParameterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#routineComment.
	VisitRoutineComment(ctx *RoutineCommentContext) interface{}

	// Visit a parse tree produced by MariaDBParser#routineLanguage.
	VisitRoutineLanguage(ctx *RoutineLanguageContext) interface{}

	// Visit a parse tree produced by MariaDBParser#routineBehavior.
	VisitRoutineBehavior(ctx *RoutineBehaviorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#routineData.
	VisitRoutineData(ctx *RoutineDataContext) interface{}

	// Visit a parse tree produced by MariaDBParser#routineSecurity.
	VisitRoutineSecurity(ctx *RoutineSecurityContext) interface{}

	// Visit a parse tree produced by MariaDBParser#serverOption.
	VisitServerOption(ctx *ServerOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createDefinitions.
	VisitCreateDefinitions(ctx *CreateDefinitionsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#columnDeclaration.
	VisitColumnDeclaration(ctx *ColumnDeclarationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#constraintDeclaration.
	VisitConstraintDeclaration(ctx *ConstraintDeclarationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#indexDeclaration.
	VisitIndexDeclaration(ctx *IndexDeclarationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nullColumnConstraint.
	VisitNullColumnConstraint(ctx *NullColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#defaultColumnConstraint.
	VisitDefaultColumnConstraint(ctx *DefaultColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#visibilityColumnConstraint.
	VisitVisibilityColumnConstraint(ctx *VisibilityColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#invisibilityColumnConstraint.
	VisitInvisibilityColumnConstraint(ctx *InvisibilityColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#autoIncrementColumnConstraint.
	VisitAutoIncrementColumnConstraint(ctx *AutoIncrementColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#primaryKeyColumnConstraint.
	VisitPrimaryKeyColumnConstraint(ctx *PrimaryKeyColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#uniqueKeyColumnConstraint.
	VisitUniqueKeyColumnConstraint(ctx *UniqueKeyColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#commentColumnConstraint.
	VisitCommentColumnConstraint(ctx *CommentColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#formatColumnConstraint.
	VisitFormatColumnConstraint(ctx *FormatColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#storageColumnConstraint.
	VisitStorageColumnConstraint(ctx *StorageColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#referenceColumnConstraint.
	VisitReferenceColumnConstraint(ctx *ReferenceColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#collateColumnConstraint.
	VisitCollateColumnConstraint(ctx *CollateColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#generatedColumnConstraint.
	VisitGeneratedColumnConstraint(ctx *GeneratedColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#serialDefaultColumnConstraint.
	VisitSerialDefaultColumnConstraint(ctx *SerialDefaultColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#checkColumnConstraint.
	VisitCheckColumnConstraint(ctx *CheckColumnConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#primaryKeyTableConstraint.
	VisitPrimaryKeyTableConstraint(ctx *PrimaryKeyTableConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#uniqueKeyTableConstraint.
	VisitUniqueKeyTableConstraint(ctx *UniqueKeyTableConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#foreignKeyTableConstraint.
	VisitForeignKeyTableConstraint(ctx *ForeignKeyTableConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#checkTableConstraint.
	VisitCheckTableConstraint(ctx *CheckTableConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#referenceDefinition.
	VisitReferenceDefinition(ctx *ReferenceDefinitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#referenceAction.
	VisitReferenceAction(ctx *ReferenceActionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#referenceControlType.
	VisitReferenceControlType(ctx *ReferenceControlTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleIndexDeclaration.
	VisitSimpleIndexDeclaration(ctx *SimpleIndexDeclarationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#specialIndexDeclaration.
	VisitSpecialIndexDeclaration(ctx *SpecialIndexDeclarationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionEngine.
	VisitTableOptionEngine(ctx *TableOptionEngineContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionEngineAttribute.
	VisitTableOptionEngineAttribute(ctx *TableOptionEngineAttributeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionAutoextendSize.
	VisitTableOptionAutoextendSize(ctx *TableOptionAutoextendSizeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionAutoIncrement.
	VisitTableOptionAutoIncrement(ctx *TableOptionAutoIncrementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionAverage.
	VisitTableOptionAverage(ctx *TableOptionAverageContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionCharset.
	VisitTableOptionCharset(ctx *TableOptionCharsetContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionChecksum.
	VisitTableOptionChecksum(ctx *TableOptionChecksumContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionCollate.
	VisitTableOptionCollate(ctx *TableOptionCollateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionComment.
	VisitTableOptionComment(ctx *TableOptionCommentContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionCompression.
	VisitTableOptionCompression(ctx *TableOptionCompressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionConnection.
	VisitTableOptionConnection(ctx *TableOptionConnectionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionDataDirectory.
	VisitTableOptionDataDirectory(ctx *TableOptionDataDirectoryContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionDelay.
	VisitTableOptionDelay(ctx *TableOptionDelayContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionEncryption.
	VisitTableOptionEncryption(ctx *TableOptionEncryptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionEncrypted.
	VisitTableOptionEncrypted(ctx *TableOptionEncryptedContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionPageCompressed.
	VisitTableOptionPageCompressed(ctx *TableOptionPageCompressedContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionPageCompressionLevel.
	VisitTableOptionPageCompressionLevel(ctx *TableOptionPageCompressionLevelContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionEncryptionKeyId.
	VisitTableOptionEncryptionKeyId(ctx *TableOptionEncryptionKeyIdContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionIndexDirectory.
	VisitTableOptionIndexDirectory(ctx *TableOptionIndexDirectoryContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionInsertMethod.
	VisitTableOptionInsertMethod(ctx *TableOptionInsertMethodContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionKeyBlockSize.
	VisitTableOptionKeyBlockSize(ctx *TableOptionKeyBlockSizeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionMaxRows.
	VisitTableOptionMaxRows(ctx *TableOptionMaxRowsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionMinRows.
	VisitTableOptionMinRows(ctx *TableOptionMinRowsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionPackKeys.
	VisitTableOptionPackKeys(ctx *TableOptionPackKeysContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionPassword.
	VisitTableOptionPassword(ctx *TableOptionPasswordContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionRowFormat.
	VisitTableOptionRowFormat(ctx *TableOptionRowFormatContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionStartTransaction.
	VisitTableOptionStartTransaction(ctx *TableOptionStartTransactionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionSecondaryEngineAttribute.
	VisitTableOptionSecondaryEngineAttribute(ctx *TableOptionSecondaryEngineAttributeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionRecalculation.
	VisitTableOptionRecalculation(ctx *TableOptionRecalculationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionPersistent.
	VisitTableOptionPersistent(ctx *TableOptionPersistentContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionSamplePage.
	VisitTableOptionSamplePage(ctx *TableOptionSamplePageContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionTablespace.
	VisitTableOptionTablespace(ctx *TableOptionTablespaceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionTableType.
	VisitTableOptionTableType(ctx *TableOptionTableTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionTransactional.
	VisitTableOptionTransactional(ctx *TableOptionTransactionalContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableOptionUnion.
	VisitTableOptionUnion(ctx *TableOptionUnionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableType.
	VisitTableType(ctx *TableTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tablespaceStorage.
	VisitTablespaceStorage(ctx *TablespaceStorageContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionDefinitions.
	VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionFunctionHash.
	VisitPartitionFunctionHash(ctx *PartitionFunctionHashContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionFunctionKey.
	VisitPartitionFunctionKey(ctx *PartitionFunctionKeyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionFunctionRange.
	VisitPartitionFunctionRange(ctx *PartitionFunctionRangeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionFunctionList.
	VisitPartitionFunctionList(ctx *PartitionFunctionListContext) interface{}

	// Visit a parse tree produced by MariaDBParser#subPartitionFunctionHash.
	VisitSubPartitionFunctionHash(ctx *SubPartitionFunctionHashContext) interface{}

	// Visit a parse tree produced by MariaDBParser#subPartitionFunctionKey.
	VisitSubPartitionFunctionKey(ctx *SubPartitionFunctionKeyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionComparison.
	VisitPartitionComparison(ctx *PartitionComparisonContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionListAtom.
	VisitPartitionListAtom(ctx *PartitionListAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionListVector.
	VisitPartitionListVector(ctx *PartitionListVectorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionSimple.
	VisitPartitionSimple(ctx *PartitionSimpleContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionDefinerAtom.
	VisitPartitionDefinerAtom(ctx *PartitionDefinerAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionDefinerVector.
	VisitPartitionDefinerVector(ctx *PartitionDefinerVectorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#subpartitionDefinition.
	VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionEngine.
	VisitPartitionOptionEngine(ctx *PartitionOptionEngineContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionComment.
	VisitPartitionOptionComment(ctx *PartitionOptionCommentContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionDataDirectory.
	VisitPartitionOptionDataDirectory(ctx *PartitionOptionDataDirectoryContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionIndexDirectory.
	VisitPartitionOptionIndexDirectory(ctx *PartitionOptionIndexDirectoryContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionMaxRows.
	VisitPartitionOptionMaxRows(ctx *PartitionOptionMaxRowsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionMinRows.
	VisitPartitionOptionMinRows(ctx *PartitionOptionMinRowsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionTablespace.
	VisitPartitionOptionTablespace(ctx *PartitionOptionTablespaceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionOptionNodeGroup.
	VisitPartitionOptionNodeGroup(ctx *PartitionOptionNodeGroupContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterSimpleDatabase.
	VisitAlterSimpleDatabase(ctx *AlterSimpleDatabaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterUpgradeName.
	VisitAlterUpgradeName(ctx *AlterUpgradeNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterEvent.
	VisitAlterEvent(ctx *AlterEventContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterFunction.
	VisitAlterFunction(ctx *AlterFunctionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterInstance.
	VisitAlterInstance(ctx *AlterInstanceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterLogfileGroup.
	VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterProcedure.
	VisitAlterProcedure(ctx *AlterProcedureContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterServer.
	VisitAlterServer(ctx *AlterServerContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterTablespace.
	VisitAlterTablespace(ctx *AlterTablespaceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterView.
	VisitAlterView(ctx *AlterViewContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterSequence.
	VisitAlterSequence(ctx *AlterSequenceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByTableOption.
	VisitAlterByTableOption(ctx *AlterByTableOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddColumn.
	VisitAlterByAddColumn(ctx *AlterByAddColumnContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddColumns.
	VisitAlterByAddColumns(ctx *AlterByAddColumnsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddIndex.
	VisitAlterByAddIndex(ctx *AlterByAddIndexContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddPrimaryKey.
	VisitAlterByAddPrimaryKey(ctx *AlterByAddPrimaryKeyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddUniqueKey.
	VisitAlterByAddUniqueKey(ctx *AlterByAddUniqueKeyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddSpecialIndex.
	VisitAlterByAddSpecialIndex(ctx *AlterByAddSpecialIndexContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddForeignKey.
	VisitAlterByAddForeignKey(ctx *AlterByAddForeignKeyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddCheckTableConstraint.
	VisitAlterByAddCheckTableConstraint(ctx *AlterByAddCheckTableConstraintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterBySetAlgorithm.
	VisitAlterBySetAlgorithm(ctx *AlterBySetAlgorithmContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByChangeDefault.
	VisitAlterByChangeDefault(ctx *AlterByChangeDefaultContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByChangeColumn.
	VisitAlterByChangeColumn(ctx *AlterByChangeColumnContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByRenameColumn.
	VisitAlterByRenameColumn(ctx *AlterByRenameColumnContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByLock.
	VisitAlterByLock(ctx *AlterByLockContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByModifyColumn.
	VisitAlterByModifyColumn(ctx *AlterByModifyColumnContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDropColumn.
	VisitAlterByDropColumn(ctx *AlterByDropColumnContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDropConstraintCheck.
	VisitAlterByDropConstraintCheck(ctx *AlterByDropConstraintCheckContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDropPrimaryKey.
	VisitAlterByDropPrimaryKey(ctx *AlterByDropPrimaryKeyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDropIndex.
	VisitAlterByDropIndex(ctx *AlterByDropIndexContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByRenameIndex.
	VisitAlterByRenameIndex(ctx *AlterByRenameIndexContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAlterIndexVisibility.
	VisitAlterByAlterIndexVisibility(ctx *AlterByAlterIndexVisibilityContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDropForeignKey.
	VisitAlterByDropForeignKey(ctx *AlterByDropForeignKeyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDisableKeys.
	VisitAlterByDisableKeys(ctx *AlterByDisableKeysContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByEnableKeys.
	VisitAlterByEnableKeys(ctx *AlterByEnableKeysContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByRename.
	VisitAlterByRename(ctx *AlterByRenameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByOrder.
	VisitAlterByOrder(ctx *AlterByOrderContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByConvertCharset.
	VisitAlterByConvertCharset(ctx *AlterByConvertCharsetContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDefaultCharset.
	VisitAlterByDefaultCharset(ctx *AlterByDefaultCharsetContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDiscardTablespace.
	VisitAlterByDiscardTablespace(ctx *AlterByDiscardTablespaceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByImportTablespace.
	VisitAlterByImportTablespace(ctx *AlterByImportTablespaceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByForce.
	VisitAlterByForce(ctx *AlterByForceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByValidate.
	VisitAlterByValidate(ctx *AlterByValidateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddPartition.
	VisitAlterByAddPartition(ctx *AlterByAddPartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDropPartition.
	VisitAlterByDropPartition(ctx *AlterByDropPartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByDiscardPartition.
	VisitAlterByDiscardPartition(ctx *AlterByDiscardPartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByImportPartition.
	VisitAlterByImportPartition(ctx *AlterByImportPartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByTruncatePartition.
	VisitAlterByTruncatePartition(ctx *AlterByTruncatePartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByCoalescePartition.
	VisitAlterByCoalescePartition(ctx *AlterByCoalescePartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByReorganizePartition.
	VisitAlterByReorganizePartition(ctx *AlterByReorganizePartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByExchangePartition.
	VisitAlterByExchangePartition(ctx *AlterByExchangePartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAnalyzePartition.
	VisitAlterByAnalyzePartition(ctx *AlterByAnalyzePartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByCheckPartition.
	VisitAlterByCheckPartition(ctx *AlterByCheckPartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByOptimizePartition.
	VisitAlterByOptimizePartition(ctx *AlterByOptimizePartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByRebuildPartition.
	VisitAlterByRebuildPartition(ctx *AlterByRebuildPartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByRepairPartition.
	VisitAlterByRepairPartition(ctx *AlterByRepairPartitionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByRemovePartitioning.
	VisitAlterByRemovePartitioning(ctx *AlterByRemovePartitioningContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByUpgradePartitioning.
	VisitAlterByUpgradePartitioning(ctx *AlterByUpgradePartitioningContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterByAddDefinitions.
	VisitAlterByAddDefinitions(ctx *AlterByAddDefinitionsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropEvent.
	VisitDropEvent(ctx *DropEventContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropIndex.
	VisitDropIndex(ctx *DropIndexContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropLogfileGroup.
	VisitDropLogfileGroup(ctx *DropLogfileGroupContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropProcedure.
	VisitDropProcedure(ctx *DropProcedureContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropFunction.
	VisitDropFunction(ctx *DropFunctionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropServer.
	VisitDropServer(ctx *DropServerContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropTablespace.
	VisitDropTablespace(ctx *DropTablespaceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropTrigger.
	VisitDropTrigger(ctx *DropTriggerContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropRole.
	VisitDropRole(ctx *DropRoleContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setRole.
	VisitSetRole(ctx *SetRoleContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dropSequence.
	VisitDropSequence(ctx *DropSequenceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#renameTable.
	VisitRenameTable(ctx *RenameTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#renameTableClause.
	VisitRenameTableClause(ctx *RenameTableClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#truncateTable.
	VisitTruncateTable(ctx *TruncateTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#callStatement.
	VisitCallStatement(ctx *CallStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#doStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerStatement.
	VisitHandlerStatement(ctx *HandlerStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#loadDataStatement.
	VisitLoadDataStatement(ctx *LoadDataStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#loadXmlStatement.
	VisitLoadXmlStatement(ctx *LoadXmlStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#replaceStatement.
	VisitReplaceStatement(ctx *ReplaceStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleSelect.
	VisitSimpleSelect(ctx *SimpleSelectContext) interface{}

	// Visit a parse tree produced by MariaDBParser#parenthesisSelect.
	VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{}

	// Visit a parse tree produced by MariaDBParser#unionSelect.
	VisitUnionSelect(ctx *UnionSelectContext) interface{}

	// Visit a parse tree produced by MariaDBParser#unionParenthesisSelect.
	VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) interface{}

	// Visit a parse tree produced by MariaDBParser#withLateralStatement.
	VisitWithLateralStatement(ctx *WithLateralStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#valuesStatement.
	VisitValuesStatement(ctx *ValuesStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#insertStatementValue.
	VisitInsertStatementValue(ctx *InsertStatementValueContext) interface{}

	// Visit a parse tree produced by MariaDBParser#updatedElement.
	VisitUpdatedElement(ctx *UpdatedElementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#assignmentField.
	VisitAssignmentField(ctx *AssignmentFieldContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lockClause.
	VisitLockClause(ctx *LockClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#singleDeleteStatement.
	VisitSingleDeleteStatement(ctx *SingleDeleteStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#multipleDeleteStatement.
	VisitMultipleDeleteStatement(ctx *MultipleDeleteStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerOpenStatement.
	VisitHandlerOpenStatement(ctx *HandlerOpenStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerReadIndexStatement.
	VisitHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerReadStatement.
	VisitHandlerReadStatement(ctx *HandlerReadStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerCloseStatement.
	VisitHandlerCloseStatement(ctx *HandlerCloseStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#singleUpdateStatement.
	VisitSingleUpdateStatement(ctx *SingleUpdateStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#multipleUpdateStatement.
	VisitMultipleUpdateStatement(ctx *MultipleUpdateStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#orderByExpression.
	VisitOrderByExpression(ctx *OrderByExpressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableSources.
	VisitTableSources(ctx *TableSourcesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableSourceBase.
	VisitTableSourceBase(ctx *TableSourceBaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableSourceNested.
	VisitTableSourceNested(ctx *TableSourceNestedContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableJson.
	VisitTableJson(ctx *TableJsonContext) interface{}

	// Visit a parse tree produced by MariaDBParser#atomTableItem.
	VisitAtomTableItem(ctx *AtomTableItemContext) interface{}

	// Visit a parse tree produced by MariaDBParser#subqueryTableItem.
	VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableSourcesItem.
	VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{}

	// Visit a parse tree produced by MariaDBParser#indexHint.
	VisitIndexHint(ctx *IndexHintContext) interface{}

	// Visit a parse tree produced by MariaDBParser#indexHintType.
	VisitIndexHintType(ctx *IndexHintTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#innerJoin.
	VisitInnerJoin(ctx *InnerJoinContext) interface{}

	// Visit a parse tree produced by MariaDBParser#straightJoin.
	VisitStraightJoin(ctx *StraightJoinContext) interface{}

	// Visit a parse tree produced by MariaDBParser#outerJoin.
	VisitOuterJoin(ctx *OuterJoinContext) interface{}

	// Visit a parse tree produced by MariaDBParser#naturalJoin.
	VisitNaturalJoin(ctx *NaturalJoinContext) interface{}

	// Visit a parse tree produced by MariaDBParser#queryExpression.
	VisitQueryExpression(ctx *QueryExpressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#queryExpressionNointo.
	VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) interface{}

	// Visit a parse tree produced by MariaDBParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#querySpecificationNointo.
	VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) interface{}

	// Visit a parse tree produced by MariaDBParser#unionParenthesis.
	VisitUnionParenthesis(ctx *UnionParenthesisContext) interface{}

	// Visit a parse tree produced by MariaDBParser#unionStatement.
	VisitUnionStatement(ctx *UnionStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lateralStatement.
	VisitLateralStatement(ctx *LateralStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonTable.
	VisitJsonTable(ctx *JsonTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonColumnList.
	VisitJsonColumnList(ctx *JsonColumnListContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonColumn.
	VisitJsonColumn(ctx *JsonColumnContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonOnEmpty.
	VisitJsonOnEmpty(ctx *JsonOnEmptyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonOnError.
	VisitJsonOnError(ctx *JsonOnErrorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectSpec.
	VisitSelectSpec(ctx *SelectSpecContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectElements.
	VisitSelectElements(ctx *SelectElementsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectStarElement.
	VisitSelectStarElement(ctx *SelectStarElementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectColumnElement.
	VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectFunctionElement.
	VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectExpressionElement.
	VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectIntoVariables.
	VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectIntoDumpFile.
	VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectIntoTextFile.
	VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectFieldsInto.
	VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{}

	// Visit a parse tree produced by MariaDBParser#selectLinesInto.
	VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{}

	// Visit a parse tree produced by MariaDBParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#groupByItem.
	VisitGroupByItem(ctx *GroupByItemContext) interface{}

	// Visit a parse tree produced by MariaDBParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#limitClauseAtom.
	VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#startTransaction.
	VisitStartTransaction(ctx *StartTransactionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#beginWork.
	VisitBeginWork(ctx *BeginWorkContext) interface{}

	// Visit a parse tree produced by MariaDBParser#commitWork.
	VisitCommitWork(ctx *CommitWorkContext) interface{}

	// Visit a parse tree produced by MariaDBParser#rollbackWork.
	VisitRollbackWork(ctx *RollbackWorkContext) interface{}

	// Visit a parse tree produced by MariaDBParser#savepointStatement.
	VisitSavepointStatement(ctx *SavepointStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#rollbackStatement.
	VisitRollbackStatement(ctx *RollbackStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#releaseStatement.
	VisitReleaseStatement(ctx *ReleaseStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lockTables.
	VisitLockTables(ctx *LockTablesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#unlockTables.
	VisitUnlockTables(ctx *UnlockTablesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setAutocommitStatement.
	VisitSetAutocommitStatement(ctx *SetAutocommitStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setTransactionStatement.
	VisitSetTransactionStatement(ctx *SetTransactionStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#transactionMode.
	VisitTransactionMode(ctx *TransactionModeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lockTableElement.
	VisitLockTableElement(ctx *LockTableElementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lockAction.
	VisitLockAction(ctx *LockActionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#transactionOption.
	VisitTransactionOption(ctx *TransactionOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#transactionLevel.
	VisitTransactionLevel(ctx *TransactionLevelContext) interface{}

	// Visit a parse tree produced by MariaDBParser#changeMaster.
	VisitChangeMaster(ctx *ChangeMasterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#changeReplicationFilter.
	VisitChangeReplicationFilter(ctx *ChangeReplicationFilterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#purgeBinaryLogs.
	VisitPurgeBinaryLogs(ctx *PurgeBinaryLogsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#resetMaster.
	VisitResetMaster(ctx *ResetMasterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#resetSlave.
	VisitResetSlave(ctx *ResetSlaveContext) interface{}

	// Visit a parse tree produced by MariaDBParser#startSlave.
	VisitStartSlave(ctx *StartSlaveContext) interface{}

	// Visit a parse tree produced by MariaDBParser#stopSlave.
	VisitStopSlave(ctx *StopSlaveContext) interface{}

	// Visit a parse tree produced by MariaDBParser#startGroupReplication.
	VisitStartGroupReplication(ctx *StartGroupReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#stopGroupReplication.
	VisitStopGroupReplication(ctx *StopGroupReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#masterStringOption.
	VisitMasterStringOption(ctx *MasterStringOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#masterDecimalOption.
	VisitMasterDecimalOption(ctx *MasterDecimalOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#masterBoolOption.
	VisitMasterBoolOption(ctx *MasterBoolOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#masterRealOption.
	VisitMasterRealOption(ctx *MasterRealOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#masterUidListOption.
	VisitMasterUidListOption(ctx *MasterUidListOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#stringMasterOption.
	VisitStringMasterOption(ctx *StringMasterOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#decimalMasterOption.
	VisitDecimalMasterOption(ctx *DecimalMasterOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#boolMasterOption.
	VisitBoolMasterOption(ctx *BoolMasterOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#channelOption.
	VisitChannelOption(ctx *ChannelOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#doDbReplication.
	VisitDoDbReplication(ctx *DoDbReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#ignoreDbReplication.
	VisitIgnoreDbReplication(ctx *IgnoreDbReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#doTableReplication.
	VisitDoTableReplication(ctx *DoTableReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#ignoreTableReplication.
	VisitIgnoreTableReplication(ctx *IgnoreTableReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#wildDoTableReplication.
	VisitWildDoTableReplication(ctx *WildDoTableReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#wildIgnoreTableReplication.
	VisitWildIgnoreTableReplication(ctx *WildIgnoreTableReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#rewriteDbReplication.
	VisitRewriteDbReplication(ctx *RewriteDbReplicationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tablePair.
	VisitTablePair(ctx *TablePairContext) interface{}

	// Visit a parse tree produced by MariaDBParser#threadType.
	VisitThreadType(ctx *ThreadTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#gtidsUntilOption.
	VisitGtidsUntilOption(ctx *GtidsUntilOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#masterLogUntilOption.
	VisitMasterLogUntilOption(ctx *MasterLogUntilOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#relayLogUntilOption.
	VisitRelayLogUntilOption(ctx *RelayLogUntilOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#sqlGapsUntilOption.
	VisitSqlGapsUntilOption(ctx *SqlGapsUntilOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#userConnectionOption.
	VisitUserConnectionOption(ctx *UserConnectionOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#passwordConnectionOption.
	VisitPasswordConnectionOption(ctx *PasswordConnectionOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#defaultAuthConnectionOption.
	VisitDefaultAuthConnectionOption(ctx *DefaultAuthConnectionOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#pluginDirConnectionOption.
	VisitPluginDirConnectionOption(ctx *PluginDirConnectionOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#gtuidSet.
	VisitGtuidSet(ctx *GtuidSetContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xaStartTransaction.
	VisitXaStartTransaction(ctx *XaStartTransactionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xaEndTransaction.
	VisitXaEndTransaction(ctx *XaEndTransactionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xaPrepareStatement.
	VisitXaPrepareStatement(ctx *XaPrepareStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xaCommitWork.
	VisitXaCommitWork(ctx *XaCommitWorkContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xaRollbackWork.
	VisitXaRollbackWork(ctx *XaRollbackWorkContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xaRecoverWork.
	VisitXaRecoverWork(ctx *XaRecoverWorkContext) interface{}

	// Visit a parse tree produced by MariaDBParser#prepareStatement.
	VisitPrepareStatement(ctx *PrepareStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#deallocatePrepare.
	VisitDeallocatePrepare(ctx *DeallocatePrepareContext) interface{}

	// Visit a parse tree produced by MariaDBParser#routineBody.
	VisitRoutineBody(ctx *RoutineBodyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#iterateStatement.
	VisitIterateStatement(ctx *IterateStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#leaveStatement.
	VisitLeaveStatement(ctx *LeaveStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#repeatStatement.
	VisitRepeatStatement(ctx *RepeatStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#CloseCursor.
	VisitCloseCursor(ctx *CloseCursorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#FetchCursor.
	VisitFetchCursor(ctx *FetchCursorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#OpenCursor.
	VisitOpenCursor(ctx *OpenCursorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#declareVariable.
	VisitDeclareVariable(ctx *DeclareVariableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#declareCondition.
	VisitDeclareCondition(ctx *DeclareConditionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#declareCursor.
	VisitDeclareCursor(ctx *DeclareCursorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#declareHandler.
	VisitDeclareHandler(ctx *DeclareHandlerContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerConditionCode.
	VisitHandlerConditionCode(ctx *HandlerConditionCodeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerConditionState.
	VisitHandlerConditionState(ctx *HandlerConditionStateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerConditionName.
	VisitHandlerConditionName(ctx *HandlerConditionNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerConditionWarning.
	VisitHandlerConditionWarning(ctx *HandlerConditionWarningContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerConditionNotfound.
	VisitHandlerConditionNotfound(ctx *HandlerConditionNotfoundContext) interface{}

	// Visit a parse tree produced by MariaDBParser#handlerConditionException.
	VisitHandlerConditionException(ctx *HandlerConditionExceptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#procedureSqlStatement.
	VisitProcedureSqlStatement(ctx *ProcedureSqlStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#caseAlternative.
	VisitCaseAlternative(ctx *CaseAlternativeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#elifAlternative.
	VisitElifAlternative(ctx *ElifAlternativeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#alterUserMysqlV56.
	VisitAlterUserMysqlV56(ctx *AlterUserMysqlV56Context) interface{}

	// Visit a parse tree produced by MariaDBParser#alterUserMysqlV80.
	VisitAlterUserMysqlV80(ctx *AlterUserMysqlV80Context) interface{}

	// Visit a parse tree produced by MariaDBParser#createUserMysqlV56.
	VisitCreateUserMysqlV56(ctx *CreateUserMysqlV56Context) interface{}

	// Visit a parse tree produced by MariaDBParser#createUserMysqlV80.
	VisitCreateUserMysqlV80(ctx *CreateUserMysqlV80Context) interface{}

	// Visit a parse tree produced by MariaDBParser#dropUser.
	VisitDropUser(ctx *DropUserContext) interface{}

	// Visit a parse tree produced by MariaDBParser#grantStatement.
	VisitGrantStatement(ctx *GrantStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#roleOption.
	VisitRoleOption(ctx *RoleOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#grantProxy.
	VisitGrantProxy(ctx *GrantProxyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#renameUser.
	VisitRenameUser(ctx *RenameUserContext) interface{}

	// Visit a parse tree produced by MariaDBParser#detailRevoke.
	VisitDetailRevoke(ctx *DetailRevokeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#shortRevoke.
	VisitShortRevoke(ctx *ShortRevokeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#roleRevoke.
	VisitRoleRevoke(ctx *RoleRevokeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#revokeProxy.
	VisitRevokeProxy(ctx *RevokeProxyContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setPasswordStatement.
	VisitSetPasswordStatement(ctx *SetPasswordStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#userSpecification.
	VisitUserSpecification(ctx *UserSpecificationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#hashAuthOption.
	VisitHashAuthOption(ctx *HashAuthOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#stringAuthOption.
	VisitStringAuthOption(ctx *StringAuthOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#moduleAuthOption.
	VisitModuleAuthOption(ctx *ModuleAuthOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleAuthOption.
	VisitSimpleAuthOption(ctx *SimpleAuthOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by MariaDBParser#passwordModuleOption.
	VisitPasswordModuleOption(ctx *PasswordModuleOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tlsOption.
	VisitTlsOption(ctx *TlsOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#userResourceOption.
	VisitUserResourceOption(ctx *UserResourceOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#userPasswordOption.
	VisitUserPasswordOption(ctx *UserPasswordOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#userLockOption.
	VisitUserLockOption(ctx *UserLockOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#privelegeClause.
	VisitPrivelegeClause(ctx *PrivelegeClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#privilege.
	VisitPrivilege(ctx *PrivilegeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#currentSchemaPriviLevel.
	VisitCurrentSchemaPriviLevel(ctx *CurrentSchemaPriviLevelContext) interface{}

	// Visit a parse tree produced by MariaDBParser#globalPrivLevel.
	VisitGlobalPrivLevel(ctx *GlobalPrivLevelContext) interface{}

	// Visit a parse tree produced by MariaDBParser#definiteSchemaPrivLevel.
	VisitDefiniteSchemaPrivLevel(ctx *DefiniteSchemaPrivLevelContext) interface{}

	// Visit a parse tree produced by MariaDBParser#definiteFullTablePrivLevel.
	VisitDefiniteFullTablePrivLevel(ctx *DefiniteFullTablePrivLevelContext) interface{}

	// Visit a parse tree produced by MariaDBParser#definiteFullTablePrivLevel2.
	VisitDefiniteFullTablePrivLevel2(ctx *DefiniteFullTablePrivLevel2Context) interface{}

	// Visit a parse tree produced by MariaDBParser#definiteTablePrivLevel.
	VisitDefiniteTablePrivLevel(ctx *DefiniteTablePrivLevelContext) interface{}

	// Visit a parse tree produced by MariaDBParser#renameUserClause.
	VisitRenameUserClause(ctx *RenameUserClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#analyzeTable.
	VisitAnalyzeTable(ctx *AnalyzeTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#checkTable.
	VisitCheckTable(ctx *CheckTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#checksumTable.
	VisitChecksumTable(ctx *ChecksumTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#optimizeTable.
	VisitOptimizeTable(ctx *OptimizeTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#repairTable.
	VisitRepairTable(ctx *RepairTableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#checkTableOption.
	VisitCheckTableOption(ctx *CheckTableOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#createUdfunction.
	VisitCreateUdfunction(ctx *CreateUdfunctionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#installPlugin.
	VisitInstallPlugin(ctx *InstallPluginContext) interface{}

	// Visit a parse tree produced by MariaDBParser#uninstallPlugin.
	VisitUninstallPlugin(ctx *UninstallPluginContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setVariable.
	VisitSetVariable(ctx *SetVariableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setCharset.
	VisitSetCharset(ctx *SetCharsetContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setNames.
	VisitSetNames(ctx *SetNamesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setPassword.
	VisitSetPassword(ctx *SetPasswordContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setAutocommit.
	VisitSetAutocommit(ctx *SetAutocommitContext) interface{}

	// Visit a parse tree produced by MariaDBParser#setNewValueInsideTrigger.
	VisitSetNewValueInsideTrigger(ctx *SetNewValueInsideTriggerContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showMasterLogs.
	VisitShowMasterLogs(ctx *ShowMasterLogsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showBinLogEvents.
	VisitShowBinLogEvents(ctx *ShowBinLogEventsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showRelayLogEvents.
	VisitShowRelayLogEvents(ctx *ShowRelayLogEventsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showObjectFilter.
	VisitShowObjectFilter(ctx *ShowObjectFilterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showColumns.
	VisitShowColumns(ctx *ShowColumnsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showCreateDb.
	VisitShowCreateDb(ctx *ShowCreateDbContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showCreateFullIdObject.
	VisitShowCreateFullIdObject(ctx *ShowCreateFullIdObjectContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showCreatePackage.
	VisitShowCreatePackage(ctx *ShowCreatePackageContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showCreateUser.
	VisitShowCreateUser(ctx *ShowCreateUserContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showEngine.
	VisitShowEngine(ctx *ShowEngineContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showInnoDBStatus.
	VisitShowInnoDBStatus(ctx *ShowInnoDBStatusContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showGlobalInfo.
	VisitShowGlobalInfo(ctx *ShowGlobalInfoContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showErrors.
	VisitShowErrors(ctx *ShowErrorsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showCountErrors.
	VisitShowCountErrors(ctx *ShowCountErrorsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showSchemaFilter.
	VisitShowSchemaFilter(ctx *ShowSchemaFilterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showRoutine.
	VisitShowRoutine(ctx *ShowRoutineContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showGrants.
	VisitShowGrants(ctx *ShowGrantsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showIndexes.
	VisitShowIndexes(ctx *ShowIndexesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showOpenTables.
	VisitShowOpenTables(ctx *ShowOpenTablesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showProfile.
	VisitShowProfile(ctx *ShowProfileContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showSlaveStatus.
	VisitShowSlaveStatus(ctx *ShowSlaveStatusContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showUserstatPlugin.
	VisitShowUserstatPlugin(ctx *ShowUserstatPluginContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showExplain.
	VisitShowExplain(ctx *ShowExplainContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showPackageStatus.
	VisitShowPackageStatus(ctx *ShowPackageStatusContext) interface{}

	// Visit a parse tree produced by MariaDBParser#explainForConnection.
	VisitExplainForConnection(ctx *ExplainForConnectionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#variableClause.
	VisitVariableClause(ctx *VariableClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showCommonEntity.
	VisitShowCommonEntity(ctx *ShowCommonEntityContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showFilter.
	VisitShowFilter(ctx *ShowFilterContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showGlobalInfoClause.
	VisitShowGlobalInfoClause(ctx *ShowGlobalInfoClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showSchemaEntity.
	VisitShowSchemaEntity(ctx *ShowSchemaEntityContext) interface{}

	// Visit a parse tree produced by MariaDBParser#showProfileType.
	VisitShowProfileType(ctx *ShowProfileTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#binlogStatement.
	VisitBinlogStatement(ctx *BinlogStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#cacheIndexStatement.
	VisitCacheIndexStatement(ctx *CacheIndexStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#flushStatement.
	VisitFlushStatement(ctx *FlushStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#killStatement.
	VisitKillStatement(ctx *KillStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#loadIndexIntoCache.
	VisitLoadIndexIntoCache(ctx *LoadIndexIntoCacheContext) interface{}

	// Visit a parse tree produced by MariaDBParser#resetStatement.
	VisitResetStatement(ctx *ResetStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#shutdownStatement.
	VisitShutdownStatement(ctx *ShutdownStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableIndexes.
	VisitTableIndexes(ctx *TableIndexesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleFlushOption.
	VisitSimpleFlushOption(ctx *SimpleFlushOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#channelFlushOption.
	VisitChannelFlushOption(ctx *ChannelFlushOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableFlushOption.
	VisitTableFlushOption(ctx *TableFlushOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#flushTableOption.
	VisitFlushTableOption(ctx *FlushTableOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#loadedTableIndexes.
	VisitLoadedTableIndexes(ctx *LoadedTableIndexesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleDescribeStatement.
	VisitSimpleDescribeStatement(ctx *SimpleDescribeStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#fullDescribeStatement.
	VisitFullDescribeStatement(ctx *FullDescribeStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#formatJsonStatement.
	VisitFormatJsonStatement(ctx *FormatJsonStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#helpStatement.
	VisitHelpStatement(ctx *HelpStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#useStatement.
	VisitUseStatement(ctx *UseStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#signalStatement.
	VisitSignalStatement(ctx *SignalStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#resignalStatement.
	VisitResignalStatement(ctx *ResignalStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#signalConditionInformation.
	VisitSignalConditionInformation(ctx *SignalConditionInformationContext) interface{}

	// Visit a parse tree produced by MariaDBParser#diagnosticsStatement.
	VisitDiagnosticsStatement(ctx *DiagnosticsStatementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#diagnosticsConditionInformationName.
	VisitDiagnosticsConditionInformationName(ctx *DiagnosticsConditionInformationNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#describeStatements.
	VisitDescribeStatements(ctx *DescribeStatementsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#describeConnection.
	VisitDescribeConnection(ctx *DescribeConnectionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#fullId.
	VisitFullId(ctx *FullIdContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#roleName.
	VisitRoleName(ctx *RoleNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#fullColumnName.
	VisitFullColumnName(ctx *FullColumnNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#indexColumnName.
	VisitIndexColumnName(ctx *IndexColumnNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#userName.
	VisitUserName(ctx *UserNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#mysqlVariable.
	VisitMysqlVariable(ctx *MysqlVariableContext) interface{}

	// Visit a parse tree produced by MariaDBParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#collationName.
	VisitCollationName(ctx *CollationNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#engineName.
	VisitEngineName(ctx *EngineNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#encryptedLiteral.
	VisitEncryptedLiteral(ctx *EncryptedLiteralContext) interface{}

	// Visit a parse tree produced by MariaDBParser#uuidSet.
	VisitUuidSet(ctx *UuidSetContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xid.
	VisitXid(ctx *XidContext) interface{}

	// Visit a parse tree produced by MariaDBParser#xuidStringId.
	VisitXuidStringId(ctx *XuidStringIdContext) interface{}

	// Visit a parse tree produced by MariaDBParser#authPlugin.
	VisitAuthPlugin(ctx *AuthPluginContext) interface{}

	// Visit a parse tree produced by MariaDBParser#uid.
	VisitUid(ctx *UidContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleId.
	VisitSimpleId(ctx *SimpleIdContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dottedId.
	VisitDottedId(ctx *DottedIdContext) interface{}

	// Visit a parse tree produced by MariaDBParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by MariaDBParser#fileSizeLiteral.
	VisitFileSizeLiteral(ctx *FileSizeLiteralContext) interface{}

	// Visit a parse tree produced by MariaDBParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by MariaDBParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by MariaDBParser#hexadecimalLiteral.
	VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nullNotnull.
	VisitNullNotnull(ctx *NullNotnullContext) interface{}

	// Visit a parse tree produced by MariaDBParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by MariaDBParser#stringDataType.
	VisitStringDataType(ctx *StringDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nationalStringDataType.
	VisitNationalStringDataType(ctx *NationalStringDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nationalVaryingStringDataType.
	VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dimensionDataType.
	VisitDimensionDataType(ctx *DimensionDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleDataType.
	VisitSimpleDataType(ctx *SimpleDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#collectionDataType.
	VisitCollectionDataType(ctx *CollectionDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#spatialDataType.
	VisitSpatialDataType(ctx *SpatialDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#longVarcharDataType.
	VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#longVarbinaryDataType.
	VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#collectionOptions.
	VisitCollectionOptions(ctx *CollectionOptionsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#convertedDataType.
	VisitConvertedDataType(ctx *ConvertedDataTypeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lengthOneDimension.
	VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lengthTwoDimension.
	VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lengthTwoOptionalDimension.
	VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#uidList.
	VisitUidList(ctx *UidListContext) interface{}

	// Visit a parse tree produced by MariaDBParser#tables.
	VisitTables(ctx *TablesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#indexColumnNames.
	VisitIndexColumnNames(ctx *IndexColumnNamesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#expressionsWithDefaults.
	VisitExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#constants.
	VisitConstants(ctx *ConstantsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleStrings.
	VisitSimpleStrings(ctx *SimpleStringsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#userVariables.
	VisitUserVariables(ctx *UserVariablesContext) interface{}

	// Visit a parse tree produced by MariaDBParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by MariaDBParser#currentTimestamp.
	VisitCurrentTimestamp(ctx *CurrentTimestampContext) interface{}

	// Visit a parse tree produced by MariaDBParser#expressionOrDefault.
	VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{}

	// Visit a parse tree produced by MariaDBParser#ifExists.
	VisitIfExists(ctx *IfExistsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#orReplace.
	VisitOrReplace(ctx *OrReplaceContext) interface{}

	// Visit a parse tree produced by MariaDBParser#waitNowaitClause.
	VisitWaitNowaitClause(ctx *WaitNowaitClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#lockOption.
	VisitLockOption(ctx *LockOptionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#specificFunctionCall.
	VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#aggregateFunctionCall.
	VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nonAggregateFunctionCall.
	VisitNonAggregateFunctionCall(ctx *NonAggregateFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#scalarFunctionCall.
	VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#udfFunctionCall.
	VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#passwordFunctionCall.
	VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#simpleFunctionCall.
	VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dataTypeFunctionCall.
	VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#valuesFunctionCall.
	VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#caseExpressionFunctionCall.
	VisitCaseExpressionFunctionCall(ctx *CaseExpressionFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#caseFunctionCall.
	VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#charFunctionCall.
	VisitCharFunctionCall(ctx *CharFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#positionFunctionCall.
	VisitPositionFunctionCall(ctx *PositionFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#substrFunctionCall.
	VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#trimFunctionCall.
	VisitTrimFunctionCall(ctx *TrimFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#weightFunctionCall.
	VisitWeightFunctionCall(ctx *WeightFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#extractFunctionCall.
	VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#getFormatFunctionCall.
	VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonValueFunctionCall.
	VisitJsonValueFunctionCall(ctx *JsonValueFunctionCallContext) interface{}

	// Visit a parse tree produced by MariaDBParser#caseFuncAlternative.
	VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#levelWeightList.
	VisitLevelWeightList(ctx *LevelWeightListContext) interface{}

	// Visit a parse tree produced by MariaDBParser#levelWeightRange.
	VisitLevelWeightRange(ctx *LevelWeightRangeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#levelInWeightListElement.
	VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{}

	// Visit a parse tree produced by MariaDBParser#aggregateWindowedFunction.
	VisitAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nonAggregateWindowedFunction.
	VisitNonAggregateWindowedFunction(ctx *NonAggregateWindowedFunctionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#overClause.
	VisitOverClause(ctx *OverClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#windowSpec.
	VisitWindowSpec(ctx *WindowSpecContext) interface{}

	// Visit a parse tree produced by MariaDBParser#windowName.
	VisitWindowName(ctx *WindowNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#frameClause.
	VisitFrameClause(ctx *FrameClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#frameUnits.
	VisitFrameUnits(ctx *FrameUnitsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#frameExtent.
	VisitFrameExtent(ctx *FrameExtentContext) interface{}

	// Visit a parse tree produced by MariaDBParser#frameBetween.
	VisitFrameBetween(ctx *FrameBetweenContext) interface{}

	// Visit a parse tree produced by MariaDBParser#frameRange.
	VisitFrameRange(ctx *FrameRangeContext) interface{}

	// Visit a parse tree produced by MariaDBParser#partitionClause.
	VisitPartitionClause(ctx *PartitionClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#scalarFunctionName.
	VisitScalarFunctionName(ctx *ScalarFunctionNameContext) interface{}

	// Visit a parse tree produced by MariaDBParser#passwordFunctionClause.
	VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by MariaDBParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) interface{}

	// Visit a parse tree produced by MariaDBParser#isExpression.
	VisitIsExpression(ctx *IsExpressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#predicateExpression.
	VisitPredicateExpression(ctx *PredicateExpressionContext) interface{}

	// Visit a parse tree produced by MariaDBParser#soundsLikePredicate.
	VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#expressionAtomPredicate.
	VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#subqueryComparisonPredicate.
	VisitSubqueryComparisonPredicate(ctx *SubqueryComparisonPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonMemberOfPredicate.
	VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#binaryComparisonPredicate.
	VisitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#inPredicate.
	VisitInPredicate(ctx *InPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#betweenPredicate.
	VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#regexpPredicate.
	VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{}

	// Visit a parse tree produced by MariaDBParser#unaryExpressionAtom.
	VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#collateExpressionAtom.
	VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#mysqlVariableExpressionAtom.
	VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nestedExpressionAtom.
	VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#nestedRowExpressionAtom.
	VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#mathExpressionAtom.
	VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#existsExpressionAtom.
	VisitExistsExpressionAtom(ctx *ExistsExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#intervalExpressionAtom.
	VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonExpressionAtom.
	VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#subqueryExpressionAtom.
	VisitSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#constantExpressionAtom.
	VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#functionCallExpressionAtom.
	VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#binaryExpressionAtom.
	VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#fullColumnNameExpressionAtom.
	VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#bitExpressionAtom.
	VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{}

	// Visit a parse tree produced by MariaDBParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#logicalOperator.
	VisitLogicalOperator(ctx *LogicalOperatorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#bitOperator.
	VisitBitOperator(ctx *BitOperatorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#mathOperator.
	VisitMathOperator(ctx *MathOperatorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#jsonOperator.
	VisitJsonOperator(ctx *JsonOperatorContext) interface{}

	// Visit a parse tree produced by MariaDBParser#charsetNameBase.
	VisitCharsetNameBase(ctx *CharsetNameBaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#transactionLevelBase.
	VisitTransactionLevelBase(ctx *TransactionLevelBaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#privilegesBase.
	VisitPrivilegesBase(ctx *PrivilegesBaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#intervalTypeBase.
	VisitIntervalTypeBase(ctx *IntervalTypeBaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#dataTypeBase.
	VisitDataTypeBase(ctx *DataTypeBaseContext) interface{}

	// Visit a parse tree produced by MariaDBParser#keywordsCanBeId.
	VisitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) interface{}

	// Visit a parse tree produced by MariaDBParser#functionNameBase.
	VisitFunctionNameBase(ctx *FunctionNameBaseContext) interface{}
}
