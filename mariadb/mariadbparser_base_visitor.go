// Code generated from MariaDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mariadb // MariaDBParser
import "github.com/antlr4-go/antlr/v4"

type BaseMariaDBParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMariaDBParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSqlStatements(ctx *SqlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSqlStatement(ctx *SqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetStatementFor(ctx *SetStatementForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDdlStatement(ctx *DdlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDmlStatement(ctx *DmlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTransactionStatement(ctx *TransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReplicationStatement(ctx *ReplicationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPreparedStatement(ctx *PreparedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAdministrationStatement(ctx *AdministrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUtilityStatement(ctx *UtilityStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateEvent(ctx *CreateEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateIndex(ctx *CreateIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateProcedure(ctx *CreateProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateFunction(ctx *CreateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateRole(ctx *CreateRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateServer(ctx *CreateServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCopyCreateTable(ctx *CopyCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitQueryCreateTable(ctx *QueryCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitColumnCreateTable(ctx *ColumnCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateTrigger(ctx *CreateTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCommonTableExpressions(ctx *CommonTableExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCteName(ctx *CteNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCteColumnName(ctx *CteColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateSequence(ctx *CreateSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSequenceSpec(ctx *SequenceSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateDatabaseOption(ctx *CreateDatabaseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCharSet(ctx *CharSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCurrentUserExpression(ctx *CurrentUserExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOwnerStatement(ctx *OwnerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPreciseSchedule(ctx *PreciseScheduleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIntervalSchedule(ctx *IntervalScheduleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTimestampValue(ctx *TimestampValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIntervalExpr(ctx *IntervalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIntervalType(ctx *IntervalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitEnableType(ctx *EnableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIndexOption(ctx *IndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitProcedureParameter(ctx *ProcedureParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoutineComment(ctx *RoutineCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoutineLanguage(ctx *RoutineLanguageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoutineBehavior(ctx *RoutineBehaviorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoutineData(ctx *RoutineDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoutineSecurity(ctx *RoutineSecurityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitServerOption(ctx *ServerOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateDefinitions(ctx *CreateDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitColumnDeclaration(ctx *ColumnDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitConstraintDeclaration(ctx *ConstraintDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIndexDeclaration(ctx *IndexDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNullColumnConstraint(ctx *NullColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDefaultColumnConstraint(ctx *DefaultColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitVisibilityColumnConstraint(ctx *VisibilityColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitInvisibilityColumnConstraint(ctx *InvisibilityColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAutoIncrementColumnConstraint(ctx *AutoIncrementColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPrimaryKeyColumnConstraint(ctx *PrimaryKeyColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUniqueKeyColumnConstraint(ctx *UniqueKeyColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCommentColumnConstraint(ctx *CommentColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFormatColumnConstraint(ctx *FormatColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStorageColumnConstraint(ctx *StorageColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReferenceColumnConstraint(ctx *ReferenceColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCollateColumnConstraint(ctx *CollateColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGeneratedColumnConstraint(ctx *GeneratedColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSerialDefaultColumnConstraint(ctx *SerialDefaultColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCheckColumnConstraint(ctx *CheckColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPrimaryKeyTableConstraint(ctx *PrimaryKeyTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUniqueKeyTableConstraint(ctx *UniqueKeyTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitForeignKeyTableConstraint(ctx *ForeignKeyTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCheckTableConstraint(ctx *CheckTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReferenceDefinition(ctx *ReferenceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReferenceAction(ctx *ReferenceActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReferenceControlType(ctx *ReferenceControlTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleIndexDeclaration(ctx *SimpleIndexDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSpecialIndexDeclaration(ctx *SpecialIndexDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionEngine(ctx *TableOptionEngineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionEngineAttribute(ctx *TableOptionEngineAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionAutoextendSize(ctx *TableOptionAutoextendSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionAutoIncrement(ctx *TableOptionAutoIncrementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionAverage(ctx *TableOptionAverageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionCharset(ctx *TableOptionCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionChecksum(ctx *TableOptionChecksumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionCollate(ctx *TableOptionCollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionComment(ctx *TableOptionCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionCompression(ctx *TableOptionCompressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionConnection(ctx *TableOptionConnectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionDataDirectory(ctx *TableOptionDataDirectoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionDelay(ctx *TableOptionDelayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionEncryption(ctx *TableOptionEncryptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionEncrypted(ctx *TableOptionEncryptedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionPageCompressed(ctx *TableOptionPageCompressedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionPageCompressionLevel(ctx *TableOptionPageCompressionLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionEncryptionKeyId(ctx *TableOptionEncryptionKeyIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionIndexDirectory(ctx *TableOptionIndexDirectoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionInsertMethod(ctx *TableOptionInsertMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionKeyBlockSize(ctx *TableOptionKeyBlockSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionMaxRows(ctx *TableOptionMaxRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionMinRows(ctx *TableOptionMinRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionPackKeys(ctx *TableOptionPackKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionPassword(ctx *TableOptionPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionRowFormat(ctx *TableOptionRowFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionStartTransaction(ctx *TableOptionStartTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionSecondaryEngineAttribute(ctx *TableOptionSecondaryEngineAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionRecalculation(ctx *TableOptionRecalculationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionPersistent(ctx *TableOptionPersistentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionSamplePage(ctx *TableOptionSamplePageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionTablespace(ctx *TableOptionTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionTableType(ctx *TableOptionTableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionTransactional(ctx *TableOptionTransactionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableOptionUnion(ctx *TableOptionUnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableType(ctx *TableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTablespaceStorage(ctx *TablespaceStorageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionFunctionHash(ctx *PartitionFunctionHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionFunctionKey(ctx *PartitionFunctionKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionFunctionRange(ctx *PartitionFunctionRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionFunctionList(ctx *PartitionFunctionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSubPartitionFunctionHash(ctx *SubPartitionFunctionHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSubPartitionFunctionKey(ctx *SubPartitionFunctionKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionComparison(ctx *PartitionComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionListAtom(ctx *PartitionListAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionListVector(ctx *PartitionListVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionSimple(ctx *PartitionSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionDefinerAtom(ctx *PartitionDefinerAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionDefinerVector(ctx *PartitionDefinerVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionEngine(ctx *PartitionOptionEngineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionComment(ctx *PartitionOptionCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionDataDirectory(ctx *PartitionOptionDataDirectoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionIndexDirectory(ctx *PartitionOptionIndexDirectoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionMaxRows(ctx *PartitionOptionMaxRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionMinRows(ctx *PartitionOptionMinRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionTablespace(ctx *PartitionOptionTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionOptionNodeGroup(ctx *PartitionOptionNodeGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterSimpleDatabase(ctx *AlterSimpleDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterUpgradeName(ctx *AlterUpgradeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterEvent(ctx *AlterEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterFunction(ctx *AlterFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterInstance(ctx *AlterInstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterProcedure(ctx *AlterProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterServer(ctx *AlterServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterTable(ctx *AlterTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterTablespace(ctx *AlterTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterView(ctx *AlterViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterSequence(ctx *AlterSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByTableOption(ctx *AlterByTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddColumn(ctx *AlterByAddColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddColumns(ctx *AlterByAddColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddIndex(ctx *AlterByAddIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddPrimaryKey(ctx *AlterByAddPrimaryKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddUniqueKey(ctx *AlterByAddUniqueKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddSpecialIndex(ctx *AlterByAddSpecialIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddForeignKey(ctx *AlterByAddForeignKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddCheckTableConstraint(ctx *AlterByAddCheckTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterBySetAlgorithm(ctx *AlterBySetAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByChangeDefault(ctx *AlterByChangeDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByChangeColumn(ctx *AlterByChangeColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByRenameColumn(ctx *AlterByRenameColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByLock(ctx *AlterByLockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByModifyColumn(ctx *AlterByModifyColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDropColumn(ctx *AlterByDropColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDropConstraintCheck(ctx *AlterByDropConstraintCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDropPrimaryKey(ctx *AlterByDropPrimaryKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDropIndex(ctx *AlterByDropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByRenameIndex(ctx *AlterByRenameIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAlterIndexVisibility(ctx *AlterByAlterIndexVisibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDropForeignKey(ctx *AlterByDropForeignKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDisableKeys(ctx *AlterByDisableKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByEnableKeys(ctx *AlterByEnableKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByRename(ctx *AlterByRenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByOrder(ctx *AlterByOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByConvertCharset(ctx *AlterByConvertCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDefaultCharset(ctx *AlterByDefaultCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDiscardTablespace(ctx *AlterByDiscardTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByImportTablespace(ctx *AlterByImportTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByForce(ctx *AlterByForceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByValidate(ctx *AlterByValidateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddPartition(ctx *AlterByAddPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDropPartition(ctx *AlterByDropPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByDiscardPartition(ctx *AlterByDiscardPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByImportPartition(ctx *AlterByImportPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByTruncatePartition(ctx *AlterByTruncatePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByCoalescePartition(ctx *AlterByCoalescePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByReorganizePartition(ctx *AlterByReorganizePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByExchangePartition(ctx *AlterByExchangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAnalyzePartition(ctx *AlterByAnalyzePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByCheckPartition(ctx *AlterByCheckPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByOptimizePartition(ctx *AlterByOptimizePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByRebuildPartition(ctx *AlterByRebuildPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByRepairPartition(ctx *AlterByRepairPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByRemovePartitioning(ctx *AlterByRemovePartitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByUpgradePartitioning(ctx *AlterByUpgradePartitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterByAddDefinitions(ctx *AlterByAddDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropDatabase(ctx *DropDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropEvent(ctx *DropEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropIndex(ctx *DropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropLogfileGroup(ctx *DropLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropProcedure(ctx *DropProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropFunction(ctx *DropFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropServer(ctx *DropServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropTablespace(ctx *DropTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropTrigger(ctx *DropTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropRole(ctx *DropRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetRole(ctx *SetRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropSequence(ctx *DropSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRenameTable(ctx *RenameTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRenameTableClause(ctx *RenameTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTruncateTable(ctx *TruncateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCallStatement(ctx *CallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerStatement(ctx *HandlerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLoadDataStatement(ctx *LoadDataStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLoadXmlStatement(ctx *LoadXmlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReplaceStatement(ctx *ReplaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleSelect(ctx *SimpleSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUnionSelect(ctx *UnionSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWithLateralStatement(ctx *WithLateralStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitValuesStatement(ctx *ValuesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitInsertStatementValue(ctx *InsertStatementValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUpdatedElement(ctx *UpdatedElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAssignmentField(ctx *AssignmentFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLockClause(ctx *LockClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSingleDeleteStatement(ctx *SingleDeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMultipleDeleteStatement(ctx *MultipleDeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerOpenStatement(ctx *HandlerOpenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerReadStatement(ctx *HandlerReadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerCloseStatement(ctx *HandlerCloseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSingleUpdateStatement(ctx *SingleUpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMultipleUpdateStatement(ctx *MultipleUpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableSources(ctx *TableSourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableSourceBase(ctx *TableSourceBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableSourceNested(ctx *TableSourceNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableJson(ctx *TableJsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStraightJoin(ctx *StraightJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOuterJoin(ctx *OuterJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNaturalJoin(ctx *NaturalJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUnionParenthesis(ctx *UnionParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUnionStatement(ctx *UnionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLateralStatement(ctx *LateralStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonTable(ctx *JsonTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonColumnList(ctx *JsonColumnListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonColumn(ctx *JsonColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonOnEmpty(ctx *JsonOnEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonOnError(ctx *JsonOnErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectSpec(ctx *SelectSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectElements(ctx *SelectElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectStarElement(ctx *SelectStarElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGroupByItem(ctx *GroupByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStartTransaction(ctx *StartTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBeginWork(ctx *BeginWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCommitWork(ctx *CommitWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRollbackWork(ctx *RollbackWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSavepointStatement(ctx *SavepointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReleaseStatement(ctx *ReleaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLockTables(ctx *LockTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUnlockTables(ctx *UnlockTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetAutocommitStatement(ctx *SetAutocommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetTransactionStatement(ctx *SetTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTransactionMode(ctx *TransactionModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLockTableElement(ctx *LockTableElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLockAction(ctx *LockActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTransactionOption(ctx *TransactionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTransactionLevel(ctx *TransactionLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitChangeMaster(ctx *ChangeMasterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitChangeReplicationFilter(ctx *ChangeReplicationFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPurgeBinaryLogs(ctx *PurgeBinaryLogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitResetMaster(ctx *ResetMasterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitResetSlave(ctx *ResetSlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStartSlave(ctx *StartSlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStopSlave(ctx *StopSlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStartGroupReplication(ctx *StartGroupReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStopGroupReplication(ctx *StopGroupReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMasterStringOption(ctx *MasterStringOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMasterDecimalOption(ctx *MasterDecimalOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMasterBoolOption(ctx *MasterBoolOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMasterRealOption(ctx *MasterRealOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMasterUidListOption(ctx *MasterUidListOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStringMasterOption(ctx *StringMasterOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDecimalMasterOption(ctx *DecimalMasterOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBoolMasterOption(ctx *BoolMasterOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitChannelOption(ctx *ChannelOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDoDbReplication(ctx *DoDbReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIgnoreDbReplication(ctx *IgnoreDbReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDoTableReplication(ctx *DoTableReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIgnoreTableReplication(ctx *IgnoreTableReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWildDoTableReplication(ctx *WildDoTableReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWildIgnoreTableReplication(ctx *WildIgnoreTableReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRewriteDbReplication(ctx *RewriteDbReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTablePair(ctx *TablePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitThreadType(ctx *ThreadTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGtidsUntilOption(ctx *GtidsUntilOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMasterLogUntilOption(ctx *MasterLogUntilOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRelayLogUntilOption(ctx *RelayLogUntilOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSqlGapsUntilOption(ctx *SqlGapsUntilOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUserConnectionOption(ctx *UserConnectionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPasswordConnectionOption(ctx *PasswordConnectionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDefaultAuthConnectionOption(ctx *DefaultAuthConnectionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPluginDirConnectionOption(ctx *PluginDirConnectionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGtuidSet(ctx *GtuidSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXaStartTransaction(ctx *XaStartTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXaEndTransaction(ctx *XaEndTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXaPrepareStatement(ctx *XaPrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXaCommitWork(ctx *XaCommitWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXaRollbackWork(ctx *XaRollbackWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXaRecoverWork(ctx *XaRecoverWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDeallocatePrepare(ctx *DeallocatePrepareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoutineBody(ctx *RoutineBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIterateStatement(ctx *IterateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLeaveStatement(ctx *LeaveStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCloseCursor(ctx *CloseCursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFetchCursor(ctx *FetchCursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOpenCursor(ctx *OpenCursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDeclareVariable(ctx *DeclareVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDeclareCondition(ctx *DeclareConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDeclareCursor(ctx *DeclareCursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDeclareHandler(ctx *DeclareHandlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerConditionCode(ctx *HandlerConditionCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerConditionState(ctx *HandlerConditionStateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerConditionName(ctx *HandlerConditionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerConditionWarning(ctx *HandlerConditionWarningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerConditionNotfound(ctx *HandlerConditionNotfoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHandlerConditionException(ctx *HandlerConditionExceptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitProcedureSqlStatement(ctx *ProcedureSqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCaseAlternative(ctx *CaseAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitElifAlternative(ctx *ElifAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterUserMysqlV56(ctx *AlterUserMysqlV56Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAlterUserMysqlV80(ctx *AlterUserMysqlV80Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateUserMysqlV56(ctx *CreateUserMysqlV56Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateUserMysqlV80(ctx *CreateUserMysqlV80Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDropUser(ctx *DropUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGrantStatement(ctx *GrantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoleOption(ctx *RoleOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGrantProxy(ctx *GrantProxyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRenameUser(ctx *RenameUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDetailRevoke(ctx *DetailRevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShortRevoke(ctx *ShortRevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoleRevoke(ctx *RoleRevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRevokeProxy(ctx *RevokeProxyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetPasswordStatement(ctx *SetPasswordStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUserSpecification(ctx *UserSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHashAuthOption(ctx *HashAuthOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStringAuthOption(ctx *StringAuthOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitModuleAuthOption(ctx *ModuleAuthOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleAuthOption(ctx *SimpleAuthOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPasswordModuleOption(ctx *PasswordModuleOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTlsOption(ctx *TlsOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUserResourceOption(ctx *UserResourceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUserPasswordOption(ctx *UserPasswordOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUserLockOption(ctx *UserLockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPrivelegeClause(ctx *PrivelegeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCurrentSchemaPriviLevel(ctx *CurrentSchemaPriviLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGlobalPrivLevel(ctx *GlobalPrivLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDefiniteSchemaPrivLevel(ctx *DefiniteSchemaPrivLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDefiniteFullTablePrivLevel(ctx *DefiniteFullTablePrivLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDefiniteFullTablePrivLevel2(ctx *DefiniteFullTablePrivLevel2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDefiniteTablePrivLevel(ctx *DefiniteTablePrivLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRenameUserClause(ctx *RenameUserClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAnalyzeTable(ctx *AnalyzeTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCheckTable(ctx *CheckTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitChecksumTable(ctx *ChecksumTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOptimizeTable(ctx *OptimizeTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRepairTable(ctx *RepairTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCheckTableOption(ctx *CheckTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCreateUdfunction(ctx *CreateUdfunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitInstallPlugin(ctx *InstallPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUninstallPlugin(ctx *UninstallPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetVariable(ctx *SetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetCharset(ctx *SetCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetNames(ctx *SetNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetPassword(ctx *SetPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetAutocommit(ctx *SetAutocommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSetNewValueInsideTrigger(ctx *SetNewValueInsideTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowMasterLogs(ctx *ShowMasterLogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowBinLogEvents(ctx *ShowBinLogEventsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowRelayLogEvents(ctx *ShowRelayLogEventsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowObjectFilter(ctx *ShowObjectFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowColumns(ctx *ShowColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowCreateDb(ctx *ShowCreateDbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowCreateFullIdObject(ctx *ShowCreateFullIdObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowCreatePackage(ctx *ShowCreatePackageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowCreateUser(ctx *ShowCreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowEngine(ctx *ShowEngineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowInnoDBStatus(ctx *ShowInnoDBStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowGlobalInfo(ctx *ShowGlobalInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowErrors(ctx *ShowErrorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowCountErrors(ctx *ShowCountErrorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowSchemaFilter(ctx *ShowSchemaFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowRoutine(ctx *ShowRoutineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowGrants(ctx *ShowGrantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowIndexes(ctx *ShowIndexesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowOpenTables(ctx *ShowOpenTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowProfile(ctx *ShowProfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowSlaveStatus(ctx *ShowSlaveStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowUserstatPlugin(ctx *ShowUserstatPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowExplain(ctx *ShowExplainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowPackageStatus(ctx *ShowPackageStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExplainForConnection(ctx *ExplainForConnectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitVariableClause(ctx *VariableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowCommonEntity(ctx *ShowCommonEntityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowFilter(ctx *ShowFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowGlobalInfoClause(ctx *ShowGlobalInfoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowSchemaEntity(ctx *ShowSchemaEntityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShowProfileType(ctx *ShowProfileTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBinlogStatement(ctx *BinlogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCacheIndexStatement(ctx *CacheIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFlushStatement(ctx *FlushStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitKillStatement(ctx *KillStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLoadIndexIntoCache(ctx *LoadIndexIntoCacheContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitResetStatement(ctx *ResetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitShutdownStatement(ctx *ShutdownStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableIndexes(ctx *TableIndexesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleFlushOption(ctx *SimpleFlushOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitChannelFlushOption(ctx *ChannelFlushOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableFlushOption(ctx *TableFlushOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFlushTableOption(ctx *FlushTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLoadedTableIndexes(ctx *LoadedTableIndexesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleDescribeStatement(ctx *SimpleDescribeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFullDescribeStatement(ctx *FullDescribeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFormatJsonStatement(ctx *FormatJsonStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHelpStatement(ctx *HelpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUseStatement(ctx *UseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSignalStatement(ctx *SignalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitResignalStatement(ctx *ResignalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSignalConditionInformation(ctx *SignalConditionInformationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDiagnosticsStatement(ctx *DiagnosticsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDiagnosticsConditionInformationName(ctx *DiagnosticsConditionInformationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDescribeStatements(ctx *DescribeStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDescribeConnection(ctx *DescribeConnectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFullId(ctx *FullIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRoleName(ctx *RoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFullColumnName(ctx *FullColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIndexColumnName(ctx *IndexColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUserName(ctx *UserNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMysqlVariable(ctx *MysqlVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitEngineName(ctx *EngineNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitEncryptedLiteral(ctx *EncryptedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUuidSet(ctx *UuidSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXid(ctx *XidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitXuidStringId(ctx *XuidStringIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAuthPlugin(ctx *AuthPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUid(ctx *UidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleId(ctx *SimpleIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDottedId(ctx *DottedIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFileSizeLiteral(ctx *FileSizeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNullNotnull(ctx *NullNotnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitStringDataType(ctx *StringDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNationalStringDataType(ctx *NationalStringDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDimensionDataType(ctx *DimensionDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleDataType(ctx *SimpleDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCollectionDataType(ctx *CollectionDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSpatialDataType(ctx *SpatialDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCollectionOptions(ctx *CollectionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitConvertedDataType(ctx *ConvertedDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUidList(ctx *UidListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTables(ctx *TablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIndexColumnNames(ctx *IndexColumnNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitConstants(ctx *ConstantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleStrings(ctx *SimpleStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUserVariables(ctx *UserVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCurrentTimestamp(ctx *CurrentTimestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIfExists(ctx *IfExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOrReplace(ctx *OrReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWaitNowaitClause(ctx *WaitNowaitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLockOption(ctx *LockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNonAggregateFunctionCall(ctx *NonAggregateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCaseExpressionFunctionCall(ctx *CaseExpressionFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCharFunctionCall(ctx *CharFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPositionFunctionCall(ctx *PositionFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTrimFunctionCall(ctx *TrimFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWeightFunctionCall(ctx *WeightFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonValueFunctionCall(ctx *JsonValueFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLevelWeightList(ctx *LevelWeightListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLevelWeightRange(ctx *LevelWeightRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNonAggregateWindowedFunction(ctx *NonAggregateWindowedFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitOverClause(ctx *OverClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWindowSpec(ctx *WindowSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitWindowName(ctx *WindowNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFrameClause(ctx *FrameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFrameUnits(ctx *FrameUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFrameExtent(ctx *FrameExtentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFrameBetween(ctx *FrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFrameRange(ctx *FrameRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPartitionClause(ctx *PartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitScalarFunctionName(ctx *ScalarFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIsExpression(ctx *IsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSubqueryComparisonPredicate(ctx *SubqueryComparisonPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitExistsExpressionAtom(ctx *ExistsExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitBitOperator(ctx *BitOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitMathOperator(ctx *MathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitJsonOperator(ctx *JsonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitCharsetNameBase(ctx *CharsetNameBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitTransactionLevelBase(ctx *TransactionLevelBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitPrivilegesBase(ctx *PrivilegesBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitIntervalTypeBase(ctx *IntervalTypeBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitDataTypeBase(ctx *DataTypeBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMariaDBParserVisitor) VisitFunctionNameBase(ctx *FunctionNameBaseContext) interface{} {
	return v.VisitChildren(ctx)
}
