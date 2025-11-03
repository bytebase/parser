// Code generated from DorisSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package doris // DorisSQL
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DorisSQLParser.
type DorisSQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DorisSQLParser#sqlStatements.
	VisitSqlStatements(ctx *SqlStatementsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#singleStatement.
	VisitSingleStatement(ctx *SingleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#useDatabaseStatement.
	VisitUseDatabaseStatement(ctx *UseDatabaseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#useCatalogStatement.
	VisitUseCatalogStatement(ctx *UseCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setCatalogStatement.
	VisitSetCatalogStatement(ctx *SetCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDatabasesStatement.
	VisitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterDbQuotaStatement.
	VisitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createDbStatement.
	VisitCreateDbStatement(ctx *CreateDbStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropDbStatement.
	VisitDropDbStatement(ctx *DropDbStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateDbStatement.
	VisitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterDatabaseRenameStatement.
	VisitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#recoverDbStmt.
	VisitRecoverDbStmt(ctx *RecoverDbStmtContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDataStmt.
	VisitShowDataStmt(ctx *ShowDataStmtContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDataDistributionStmt.
	VisitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnDesc.
	VisitColumnDesc(ctx *ColumnDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#defaultDesc.
	VisitDefaultDesc(ctx *DefaultDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#generatedColumnDesc.
	VisitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#indexDesc.
	VisitIndexDesc(ctx *IndexDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#engineDesc.
	VisitEngineDesc(ctx *EngineDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#charsetDesc.
	VisitCharsetDesc(ctx *CharsetDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#collateDesc.
	VisitCollateDesc(ctx *CollateDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#keyDesc.
	VisitKeyDesc(ctx *KeyDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#orderByDesc.
	VisitOrderByDesc(ctx *OrderByDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnNullable.
	VisitColumnNullable(ctx *ColumnNullableContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#typeWithNullable.
	VisitTypeWithNullable(ctx *TypeWithNullableContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#aggStateDesc.
	VisitAggStateDesc(ctx *AggStateDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#aggDesc.
	VisitAggDesc(ctx *AggDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rollupDesc.
	VisitRollupDesc(ctx *RollupDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rollupItem.
	VisitRollupItem(ctx *RollupItemContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dupKeys.
	VisitDupKeys(ctx *DupKeysContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#fromRollup.
	VisitFromRollup(ctx *FromRollupContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#orReplace.
	VisitOrReplace(ctx *OrReplaceContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createTableAsSelectStatement.
	VisitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropTableStatement.
	VisitDropTableStatement(ctx *DropTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cleanTemporaryTableStatement.
	VisitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterTableStatement.
	VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createIndexStatement.
	VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropIndexStatement.
	VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTableStatement.
	VisitShowTableStatement(ctx *ShowTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTemporaryTablesStatement.
	VisitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateTableStatement.
	VisitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showColumnStatement.
	VisitShowColumnStatement(ctx *ShowColumnStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTableStatusStatement.
	VisitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#refreshTableStatement.
	VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showAlterStatement.
	VisitShowAlterStatement(ctx *ShowAlterStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#descTableStatement.
	VisitDescTableStatement(ctx *DescTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createTableLikeStatement.
	VisitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showIndexStatement.
	VisitShowIndexStatement(ctx *ShowIndexStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#recoverTableStatement.
	VisitRecoverTableStatement(ctx *RecoverTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#truncateTableStatement.
	VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelAlterTableStatement.
	VisitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showPartitionsStatement.
	VisitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#recoverPartitionStatement.
	VisitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createViewStatement.
	VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterViewStatement.
	VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropViewStatement.
	VisitDropViewStatement(ctx *DropViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnNameWithComment.
	VisitColumnNameWithComment(ctx *ColumnNameWithCommentContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#submitTaskStatement.
	VisitSubmitTaskStatement(ctx *SubmitTaskStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#taskClause.
	VisitTaskClause(ctx *TaskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropTaskStatement.
	VisitDropTaskStatement(ctx *DropTaskStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#taskScheduleDesc.
	VisitTaskScheduleDesc(ctx *TaskScheduleDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createMaterializedViewStatement.
	VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#mvPartitionExprs.
	VisitMvPartitionExprs(ctx *MvPartitionExprsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#materializedViewDesc.
	VisitMaterializedViewDesc(ctx *MaterializedViewDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showMaterializedViewsStatement.
	VisitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropMaterializedViewStatement.
	VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterMaterializedViewStatement.
	VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#refreshMaterializedViewStatement.
	VisitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelRefreshMaterializedViewStatement.
	VisitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminSetConfigStatement.
	VisitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminSetReplicaStatusStatement.
	VisitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminShowConfigStatement.
	VisitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminShowReplicaDistributionStatement.
	VisitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminShowReplicaStatusStatement.
	VisitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminRepairTableStatement.
	VisitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminCancelRepairTableStatement.
	VisitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminCheckTabletsStatement.
	VisitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminSetPartitionVersion.
	VisitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#killStatement.
	VisitKillStatement(ctx *KillStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#syncStatement.
	VisitSyncStatement(ctx *SyncStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminSetAutomatedSnapshotOnStatement.
	VisitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#adminSetAutomatedSnapshotOffStatement.
	VisitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterSystemStatement.
	VisitAlterSystemStatement(ctx *AlterSystemStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelAlterSystemStatement.
	VisitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showComputeNodesStatement.
	VisitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createExternalCatalogStatement.
	VisitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateExternalCatalogStatement.
	VisitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropExternalCatalogStatement.
	VisitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCatalogsStatement.
	VisitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterCatalogStatement.
	VisitAlterCatalogStatement(ctx *AlterCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createStorageVolumeStatement.
	VisitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#typeDesc.
	VisitTypeDesc(ctx *TypeDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#locationsDesc.
	VisitLocationsDesc(ctx *LocationsDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showStorageVolumesStatement.
	VisitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropStorageVolumeStatement.
	VisitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterStorageVolumeStatement.
	VisitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterStorageVolumeClause.
	VisitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyStorageVolumePropertiesClause.
	VisitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyStorageVolumeCommentClause.
	VisitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#descStorageVolumeStatement.
	VisitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setDefaultStorageVolumeStatement.
	VisitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#updateFailPointStatusStatement.
	VisitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showFailPointStatement.
	VisitShowFailPointStatement(ctx *ShowFailPointStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createDictionaryStatement.
	VisitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropDictionaryStatement.
	VisitDropDictionaryStatement(ctx *DropDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#refreshDictionaryStatement.
	VisitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDictionaryStatement.
	VisitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelRefreshDictionaryStatement.
	VisitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dictionaryColumnDesc.
	VisitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dictionaryName.
	VisitDictionaryName(ctx *DictionaryNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterClause.
	VisitAlterClause(ctx *AlterClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addFrontendClause.
	VisitAddFrontendClause(ctx *AddFrontendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropFrontendClause.
	VisitDropFrontendClause(ctx *DropFrontendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyFrontendHostClause.
	VisitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addBackendClause.
	VisitAddBackendClause(ctx *AddBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropBackendClause.
	VisitDropBackendClause(ctx *DropBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#decommissionBackendClause.
	VisitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyBackendClause.
	VisitModifyBackendClause(ctx *ModifyBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addComputeNodeClause.
	VisitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropComputeNodeClause.
	VisitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyBrokerClause.
	VisitModifyBrokerClause(ctx *ModifyBrokerClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterLoadErrorUrlClause.
	VisitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createImageClause.
	VisitCreateImageClause(ctx *CreateImageClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cleanTabletSchedQClause.
	VisitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#decommissionDiskClause.
	VisitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelDecommissionDiskClause.
	VisitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#disableDiskClause.
	VisitDisableDiskClause(ctx *DisableDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelDisableDiskClause.
	VisitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createIndexClause.
	VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropIndexClause.
	VisitDropIndexClause(ctx *DropIndexClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tableRenameClause.
	VisitTableRenameClause(ctx *TableRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#swapTableClause.
	VisitSwapTableClause(ctx *SwapTableClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyPropertiesClause.
	VisitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyCommentClause.
	VisitModifyCommentClause(ctx *ModifyCommentClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#optimizeRange.
	VisitOptimizeRange(ctx *OptimizeRangeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#optimizeClause.
	VisitOptimizeClause(ctx *OptimizeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addColumnClause.
	VisitAddColumnClause(ctx *AddColumnClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addColumnsClause.
	VisitAddColumnsClause(ctx *AddColumnsClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropColumnClause.
	VisitDropColumnClause(ctx *DropColumnClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyColumnClause.
	VisitModifyColumnClause(ctx *ModifyColumnClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyColumnCommentClause.
	VisitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnRenameClause.
	VisitColumnRenameClause(ctx *ColumnRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#reorderColumnsClause.
	VisitReorderColumnsClause(ctx *ReorderColumnsClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rollupRenameClause.
	VisitRollupRenameClause(ctx *RollupRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#compactionClause.
	VisitCompactionClause(ctx *CompactionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#subfieldName.
	VisitSubfieldName(ctx *SubfieldNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#nestedFieldName.
	VisitNestedFieldName(ctx *NestedFieldNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addFieldClause.
	VisitAddFieldClause(ctx *AddFieldClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropFieldClause.
	VisitDropFieldClause(ctx *DropFieldClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createOrReplaceTagClause.
	VisitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createOrReplaceBranchClause.
	VisitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropBranchClause.
	VisitDropBranchClause(ctx *DropBranchClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropTagClause.
	VisitDropTagClause(ctx *DropTagClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tableOperationClause.
	VisitTableOperationClause(ctx *TableOperationClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tagOptions.
	VisitTagOptions(ctx *TagOptionsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#branchOptions.
	VisitBranchOptions(ctx *BranchOptionsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#snapshotRetention.
	VisitSnapshotRetention(ctx *SnapshotRetentionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#refRetain.
	VisitRefRetain(ctx *RefRetainContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#maxSnapshotAge.
	VisitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#minSnapshotsToKeep.
	VisitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#snapshotId.
	VisitSnapshotId(ctx *SnapshotIdContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#timeUnit.
	VisitTimeUnit(ctx *TimeUnitContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#integer_list.
	VisitInteger_list(ctx *Integer_listContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropPersistentIndexClause.
	VisitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addPartitionClause.
	VisitAddPartitionClause(ctx *AddPartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropPartitionClause.
	VisitDropPartitionClause(ctx *DropPartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#truncatePartitionClause.
	VisitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#modifyPartitionClause.
	VisitModifyPartitionClause(ctx *ModifyPartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#replacePartitionClause.
	VisitReplacePartitionClause(ctx *ReplacePartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionRenameClause.
	VisitPartitionRenameClause(ctx *PartitionRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#insertLabelOrColumnAliases.
	VisitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnAliasesOrByName.
	VisitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createRoutineLoadStatement.
	VisitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterRoutineLoadStatement.
	VisitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dataSource.
	VisitDataSource(ctx *DataSourceContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#loadProperties.
	VisitLoadProperties(ctx *LoadPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#colSeparatorProperty.
	VisitColSeparatorProperty(ctx *ColSeparatorPropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rowDelimiterProperty.
	VisitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#importColumns.
	VisitImportColumns(ctx *ImportColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnProperties.
	VisitColumnProperties(ctx *ColumnPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#jobProperties.
	VisitJobProperties(ctx *JobPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dataSourceProperties.
	VisitDataSourceProperties(ctx *DataSourcePropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#stopRoutineLoadStatement.
	VisitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#resumeRoutineLoadStatement.
	VisitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#pauseRoutineLoadStatement.
	VisitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showRoutineLoadStatement.
	VisitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showRoutineLoadTaskStatement.
	VisitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateRoutineLoadStatement.
	VisitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showStreamLoadStatement.
	VisitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#analyzeStatement.
	VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#regularColumns.
	VisitRegularColumns(ctx *RegularColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#allColumns.
	VisitAllColumns(ctx *AllColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#predicateColumns.
	VisitPredicateColumns(ctx *PredicateColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#multiColumnSet.
	VisitMultiColumnSet(ctx *MultiColumnSetContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropStatsStatement.
	VisitDropStatsStatement(ctx *DropStatsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#histogramStatement.
	VisitHistogramStatement(ctx *HistogramStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#analyzeHistogramStatement.
	VisitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropHistogramStatement.
	VisitDropHistogramStatement(ctx *DropHistogramStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createAnalyzeStatement.
	VisitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropAnalyzeJobStatement.
	VisitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showAnalyzeStatement.
	VisitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showStatsMetaStatement.
	VisitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showHistogramMetaStatement.
	VisitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#killAnalyzeStatement.
	VisitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#analyzeProfileStatement.
	VisitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createBaselinePlanStatement.
	VisitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropBaselinePlanStatement.
	VisitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showBaselinePlanStatement.
	VisitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createResourceGroupStatement.
	VisitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropResourceGroupStatement.
	VisitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterResourceGroupStatement.
	VisitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showResourceGroupStatement.
	VisitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showResourceGroupUsageStatement.
	VisitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createResourceStatement.
	VisitCreateResourceStatement(ctx *CreateResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterResourceStatement.
	VisitAlterResourceStatement(ctx *AlterResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropResourceStatement.
	VisitDropResourceStatement(ctx *DropResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showResourceStatement.
	VisitShowResourceStatement(ctx *ShowResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#classifier.
	VisitClassifier(ctx *ClassifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showFunctionsStatement.
	VisitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropFunctionStatement.
	VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createFunctionStatement.
	VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#inlineFunction.
	VisitInlineFunction(ctx *InlineFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#loadStatement.
	VisitLoadStatement(ctx *LoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#labelName.
	VisitLabelName(ctx *LabelNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dataDescList.
	VisitDataDescList(ctx *DataDescListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dataDesc.
	VisitDataDesc(ctx *DataDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#formatProps.
	VisitFormatProps(ctx *FormatPropsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#brokerDesc.
	VisitBrokerDesc(ctx *BrokerDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#resourceDesc.
	VisitResourceDesc(ctx *ResourceDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showLoadStatement.
	VisitShowLoadStatement(ctx *ShowLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showLoadWarningsStatement.
	VisitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelLoadStatement.
	VisitCancelLoadStatement(ctx *CancelLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterLoadStatement.
	VisitAlterLoadStatement(ctx *AlterLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelCompactionStatement.
	VisitCancelCompactionStatement(ctx *CancelCompactionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showAuthorStatement.
	VisitShowAuthorStatement(ctx *ShowAuthorStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showBackendsStatement.
	VisitShowBackendsStatement(ctx *ShowBackendsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showBrokerStatement.
	VisitShowBrokerStatement(ctx *ShowBrokerStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCharsetStatement.
	VisitShowCharsetStatement(ctx *ShowCharsetStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCollationStatement.
	VisitShowCollationStatement(ctx *ShowCollationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDeleteStatement.
	VisitShowDeleteStatement(ctx *ShowDeleteStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDynamicPartitionStatement.
	VisitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showEventsStatement.
	VisitShowEventsStatement(ctx *ShowEventsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showEnginesStatement.
	VisitShowEnginesStatement(ctx *ShowEnginesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showFrontendsStatement.
	VisitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showFrontendsDisksStatement.
	VisitShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showPluginsStatement.
	VisitShowPluginsStatement(ctx *ShowPluginsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showRepositoriesStatement.
	VisitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showOpenTableStatement.
	VisitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showPrivilegesStatement.
	VisitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showProcedureStatement.
	VisitShowProcedureStatement(ctx *ShowProcedureStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showProcStatement.
	VisitShowProcStatement(ctx *ShowProcStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showProcesslistStatement.
	VisitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showProfilelistStatement.
	VisitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showRunningQueriesStatement.
	VisitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showStatusStatement.
	VisitShowStatusStatement(ctx *ShowStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTabletStatement.
	VisitShowTabletStatement(ctx *ShowTabletStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTransactionStatement.
	VisitShowTransactionStatement(ctx *ShowTransactionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTriggersStatement.
	VisitShowTriggersStatement(ctx *ShowTriggersStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showUserPropertyStatement.
	VisitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showVariablesStatement.
	VisitShowVariablesStatement(ctx *ShowVariablesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showWarningStatement.
	VisitShowWarningStatement(ctx *ShowWarningStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#helpStatement.
	VisitHelpStatement(ctx *HelpStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showQueryProfileStatement.
	VisitShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showQueryStatsStatement.
	VisitShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showLoadProfileStatement.
	VisitShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDataSkewStatement.
	VisitShowDataSkewStatement(ctx *ShowDataSkewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDataTypesStatement.
	VisitShowDataTypesStatement(ctx *ShowDataTypesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showSyncJobStatement.
	VisitShowSyncJobStatement(ctx *ShowSyncJobStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showPolicyStatement.
	VisitShowPolicyStatement(ctx *ShowPolicyStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showSqlBlockRuleStatement.
	VisitShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showEncryptKeysStatement.
	VisitShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateLoadStatement.
	VisitShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateRepositoryStatement.
	VisitShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showLastInsertStatement.
	VisitShowLastInsertStatement(ctx *ShowLastInsertStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTableIdStatement.
	VisitShowTableIdStatement(ctx *ShowTableIdStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDatabaseIdStatement.
	VisitShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showPartitionIdStatement.
	VisitShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTableStatsStatement.
	VisitShowTableStatsStatement(ctx *ShowTableStatsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showColumnStatsStatement.
	VisitShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showConvertLightSchemaChangeStatement.
	VisitShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCatalogRecycleBinStatement.
	VisitShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showTrashStatement.
	VisitShowTrashStatement(ctx *ShowTrashStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showMigrationsStatement.
	VisitShowMigrationsStatement(ctx *ShowMigrationsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showWorkloadGroupsStatement.
	VisitShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showJobTaskStatement.
	VisitShowJobTaskStatement(ctx *ShowJobTaskStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createUserStatement.
	VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropUserStatement.
	VisitDropUserStatement(ctx *DropUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterUserStatement.
	VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showUserStatement.
	VisitShowUserStatement(ctx *ShowUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showAllAuthentication.
	VisitShowAllAuthentication(ctx *ShowAllAuthenticationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showAuthenticationForUser.
	VisitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#executeAsStatement.
	VisitExecuteAsStatement(ctx *ExecuteAsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createRoleStatement.
	VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterRoleStatement.
	VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropRoleStatement.
	VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showRolesStatement.
	VisitShowRolesStatement(ctx *ShowRolesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantRoleToUser.
	VisitGrantRoleToUser(ctx *GrantRoleToUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantRoleToRole.
	VisitGrantRoleToRole(ctx *GrantRoleToRoleContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeRoleFromUser.
	VisitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeRoleFromRole.
	VisitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setRoleStatement.
	VisitSetRoleStatement(ctx *SetRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setDefaultRoleStatement.
	VisitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantRevokeClause.
	VisitGrantRevokeClause(ctx *GrantRevokeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantOnUser.
	VisitGrantOnUser(ctx *GrantOnUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantOnTableBrief.
	VisitGrantOnTableBrief(ctx *GrantOnTableBriefContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantOnFunc.
	VisitGrantOnFunc(ctx *GrantOnFuncContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantOnSystem.
	VisitGrantOnSystem(ctx *GrantOnSystemContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantOnPrimaryObj.
	VisitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#grantOnAll.
	VisitGrantOnAll(ctx *GrantOnAllContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeOnUser.
	VisitRevokeOnUser(ctx *RevokeOnUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeOnTableBrief.
	VisitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeOnFunc.
	VisitRevokeOnFunc(ctx *RevokeOnFuncContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeOnSystem.
	VisitRevokeOnSystem(ctx *RevokeOnSystemContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeOnPrimaryObj.
	VisitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#revokeOnAll.
	VisitRevokeOnAll(ctx *RevokeOnAllContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showGrantsStatement.
	VisitShowGrantsStatement(ctx *ShowGrantsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#authWithoutPlugin.
	VisitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#authWithPlugin.
	VisitAuthWithPlugin(ctx *AuthWithPluginContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#privObjectName.
	VisitPrivObjectName(ctx *PrivObjectNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#privObjectNameList.
	VisitPrivObjectNameList(ctx *PrivObjectNameListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#privFunctionObjectNameList.
	VisitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#privilegeTypeList.
	VisitPrivilegeTypeList(ctx *PrivilegeTypeListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#privilegeType.
	VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#privObjectType.
	VisitPrivObjectType(ctx *PrivObjectTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#privObjectTypePlural.
	VisitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createSecurityIntegrationStatement.
	VisitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterSecurityIntegrationStatement.
	VisitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropSecurityIntegrationStatement.
	VisitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showSecurityIntegrationStatement.
	VisitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateSecurityIntegrationStatement.
	VisitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createGroupProviderStatement.
	VisitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropGroupProviderStatement.
	VisitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showGroupProvidersStatement.
	VisitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showCreateGroupProviderStatement.
	VisitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#backupStatement.
	VisitBackupStatement(ctx *BackupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelBackupStatement.
	VisitCancelBackupStatement(ctx *CancelBackupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showBackupStatement.
	VisitShowBackupStatement(ctx *ShowBackupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#restoreStatement.
	VisitRestoreStatement(ctx *RestoreStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelRestoreStatement.
	VisitCancelRestoreStatement(ctx *CancelRestoreStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showRestoreStatement.
	VisitShowRestoreStatement(ctx *ShowRestoreStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showSnapshotStatement.
	VisitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createRepositoryStatement.
	VisitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropRepositoryStatement.
	VisitDropRepositoryStatement(ctx *DropRepositoryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addSqlBlackListStatement.
	VisitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#delSqlBlackListStatement.
	VisitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showSqlBlackListStatement.
	VisitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showWhiteListStatement.
	VisitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#addBackendBlackListStatement.
	VisitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#delBackendBlackListStatement.
	VisitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showBackendBlackListStatement.
	VisitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dataCacheTarget.
	VisitDataCacheTarget(ctx *DataCacheTargetContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createDataCacheRuleStatement.
	VisitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showDataCacheRulesStatement.
	VisitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropDataCacheRuleStatement.
	VisitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#clearDataCacheRulesStatement.
	VisitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dataCacheSelectStatement.
	VisitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cancelExportStatement.
	VisitCancelExportStatement(ctx *CancelExportStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showExportStatement.
	VisitShowExportStatement(ctx *ShowExportStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#installPluginStatement.
	VisitInstallPluginStatement(ctx *InstallPluginStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#uninstallPluginStatement.
	VisitUninstallPluginStatement(ctx *UninstallPluginStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createFileStatement.
	VisitCreateFileStatement(ctx *CreateFileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropFileStatement.
	VisitDropFileStatement(ctx *DropFileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showSmallFilesStatement.
	VisitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createPipeStatement.
	VisitCreatePipeStatement(ctx *CreatePipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropPipeStatement.
	VisitDropPipeStatement(ctx *DropPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterPipeClause.
	VisitAlterPipeClause(ctx *AlterPipeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterPipeStatement.
	VisitAlterPipeStatement(ctx *AlterPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#descPipeStatement.
	VisitDescPipeStatement(ctx *DescPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showPipeStatement.
	VisitShowPipeStatement(ctx *ShowPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setStatement.
	VisitSetStatement(ctx *SetStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setNames.
	VisitSetNames(ctx *SetNamesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setPassword.
	VisitSetPassword(ctx *SetPasswordContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setUserVar.
	VisitSetUserVar(ctx *SetUserVarContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setSystemVar.
	VisitSetSystemVar(ctx *SetSystemVarContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#transaction_characteristics.
	VisitTransaction_characteristics(ctx *Transaction_characteristicsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#transaction_access_mode.
	VisitTransaction_access_mode(ctx *Transaction_access_modeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#isolation_level.
	VisitIsolation_level(ctx *Isolation_levelContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#isolation_types.
	VisitIsolation_types(ctx *Isolation_typesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setExprOrDefault.
	VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setUserPropertyStatement.
	VisitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#roleList.
	VisitRoleList(ctx *RoleListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#executeScriptStatement.
	VisitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#unsupportedStatement.
	VisitUnsupportedStatement(ctx *UnsupportedStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#lock_item.
	VisitLock_item(ctx *Lock_itemContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#lock_type.
	VisitLock_type(ctx *Lock_typeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterPlanAdvisorAddStatement.
	VisitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#truncatePlanAdvisorStatement.
	VisitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterPlanAdvisorDropStatement.
	VisitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showPlanAdvisorStatement.
	VisitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createWarehouseStatement.
	VisitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropWarehouseStatement.
	VisitDropWarehouseStatement(ctx *DropWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#suspendWarehouseStatement.
	VisitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#resumeWarehouseStatement.
	VisitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setWarehouseStatement.
	VisitSetWarehouseStatement(ctx *SetWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showWarehousesStatement.
	VisitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showClustersStatement.
	VisitShowClustersStatement(ctx *ShowClustersStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#showNodesStatement.
	VisitShowNodesStatement(ctx *ShowNodesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterWarehouseStatement.
	VisitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#createCNGroupStatement.
	VisitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dropCNGroupStatement.
	VisitDropCNGroupStatement(ctx *DropCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#enableCNGroupStatement.
	VisitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#disableCNGroupStatement.
	VisitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#alterCNGroupStatement.
	VisitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#beginStatement.
	VisitBeginStatement(ctx *BeginStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#commitStatement.
	VisitCommitStatement(ctx *CommitStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rollbackStatement.
	VisitRollbackStatement(ctx *RollbackStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#translateStatement.
	VisitTranslateStatement(ctx *TranslateStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dialect.
	VisitDialect(ctx *DialectContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#translateSQL.
	VisitTranslateSQL(ctx *TranslateSQLContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#queryStatement.
	VisitQueryStatement(ctx *QueryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#queryRelation.
	VisitQueryRelation(ctx *QueryRelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#queryNoWith.
	VisitQueryNoWith(ctx *QueryNoWithContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#queryPeriod.
	VisitQueryPeriod(ctx *QueryPeriodContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#periodType.
	VisitPeriodType(ctx *PeriodTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#queryWithParentheses.
	VisitQueryWithParentheses(ctx *QueryWithParenthesesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setOperation.
	VisitSetOperation(ctx *SetOperationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#queryPrimaryDefault.
	VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rowConstructor.
	VisitRowConstructor(ctx *RowConstructorContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#sortItem.
	VisitSortItem(ctx *SortItemContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#limitConstExpr.
	VisitLimitConstExpr(ctx *LimitConstExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#limitElement.
	VisitLimitElement(ctx *LimitElementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#from.
	VisitFrom(ctx *FromContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dual.
	VisitDual(ctx *DualContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rollup.
	VisitRollup(ctx *RollupContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cube.
	VisitCube(ctx *CubeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#multipleGroupingSets.
	VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#singleGroupingSet.
	VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#groupingSet.
	VisitGroupingSet(ctx *GroupingSetContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#commonTableExpression.
	VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#setQuantifier.
	VisitSetQuantifier(ctx *SetQuantifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#selectSingle.
	VisitSelectSingle(ctx *SelectSingleContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#selectAll.
	VisitSelectAll(ctx *SelectAllContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#excludeClause.
	VisitExcludeClause(ctx *ExcludeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#relations.
	VisitRelations(ctx *RelationsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#relationLateralView.
	VisitRelationLateralView(ctx *RelationLateralViewContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#lateralView.
	VisitLateralView(ctx *LateralViewContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#generatorFunction.
	VisitGeneratorFunction(ctx *GeneratorFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tableAtom.
	VisitTableAtom(ctx *TableAtomContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#inlineTable.
	VisitInlineTable(ctx *InlineTableContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#subqueryWithAlias.
	VisitSubqueryWithAlias(ctx *SubqueryWithAliasContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tableFunction.
	VisitTableFunction(ctx *TableFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#normalizedTableFunction.
	VisitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#fileTableFunction.
	VisitFileTableFunction(ctx *FileTableFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#parenthesizedRelation.
	VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#pivotClause.
	VisitPivotClause(ctx *PivotClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#pivotAggregationExpression.
	VisitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#pivotValue.
	VisitPivotValue(ctx *PivotValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#sampleClause.
	VisitSampleClause(ctx *SampleClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#namedArgumentList.
	VisitNamedArgumentList(ctx *NamedArgumentListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#namedArguments.
	VisitNamedArguments(ctx *NamedArgumentsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#joinRelation.
	VisitJoinRelation(ctx *JoinRelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#crossOrInnerJoinType.
	VisitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#outerAndSemiJoinType.
	VisitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#bracketHint.
	VisitBracketHint(ctx *BracketHintContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#hintMap.
	VisitHintMap(ctx *HintMapContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#joinCriteria.
	VisitJoinCriteria(ctx *JoinCriteriaContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnAliases.
	VisitColumnAliases(ctx *ColumnAliasesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnAliasesWithoutParentheses.
	VisitColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionNames.
	VisitPartitionNames(ctx *PartitionNamesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#keyPartitionList.
	VisitKeyPartitionList(ctx *KeyPartitionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tabletList.
	VisitTabletList(ctx *TabletListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#prepareStatement.
	VisitPrepareStatement(ctx *PrepareStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#prepareSql.
	VisitPrepareSql(ctx *PrepareSqlContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#deallocateStatement.
	VisitDeallocateStatement(ctx *DeallocateStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#replicaList.
	VisitReplicaList(ctx *ReplicaListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#expressionsWithDefault.
	VisitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#expressionOrDefault.
	VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#mapExpressionList.
	VisitMapExpressionList(ctx *MapExpressionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#mapExpression.
	VisitMapExpression(ctx *MapExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#expressionSingleton.
	VisitExpressionSingleton(ctx *ExpressionSingletonContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#expressionDefault.
	VisitExpressionDefault(ctx *ExpressionDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#logicalNot.
	VisitLogicalNot(ctx *LogicalNotContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#logicalBinary.
	VisitLogicalBinary(ctx *LogicalBinaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#booleanExpressionDefault.
	VisitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#isNull.
	VisitIsNull(ctx *IsNullContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#scalarSubquery.
	VisitScalarSubquery(ctx *ScalarSubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tupleInSubquery.
	VisitTupleInSubquery(ctx *TupleInSubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#inSubquery.
	VisitInSubquery(ctx *InSubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#inList.
	VisitInList(ctx *InListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#between.
	VisitBetween(ctx *BetweenContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#like.
	VisitLike(ctx *LikeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#valueExpressionDefault.
	VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#arithmeticBinary.
	VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dereference.
	VisitDereference(ctx *DereferenceContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#odbcFunctionCallExpression.
	VisitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#matchExpr.
	VisitMatchExpr(ctx *MatchExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnRef.
	VisitColumnRef(ctx *ColumnRefContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#convert.
	VisitConvert(ctx *ConvertContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#collectionSubscript.
	VisitCollectionSubscript(ctx *CollectionSubscriptContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#cast.
	VisitCast(ctx *CastContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#userVariableExpression.
	VisitUserVariableExpression(ctx *UserVariableExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#simpleCase.
	VisitSimpleCase(ctx *SimpleCaseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#arrowExpression.
	VisitArrowExpression(ctx *ArrowExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#arrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#systemVariableExpression.
	VisitSystemVariableExpression(ctx *SystemVariableExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#concat.
	VisitConcat(ctx *ConcatContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#subqueryExpression.
	VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#lambdaFunctionExpr.
	VisitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dictionaryGetExpr.
	VisitDictionaryGetExpr(ctx *DictionaryGetExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#collate.
	VisitCollate(ctx *CollateContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#arrayConstructor.
	VisitArrayConstructor(ctx *ArrayConstructorContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#mapConstructor.
	VisitMapConstructor(ctx *MapConstructorContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#arraySlice.
	VisitArraySlice(ctx *ArraySliceContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#exists.
	VisitExists(ctx *ExistsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#searchedCase.
	VisitSearchedCase(ctx *SearchedCaseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#arithmeticUnary.
	VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#dateLiteral.
	VisitDateLiteral(ctx *DateLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#intervalLiteral.
	VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#unitBoundaryLiteral.
	VisitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#binaryLiteral.
	VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#Parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#extract.
	VisitExtract(ctx *ExtractContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#groupingOperation.
	VisitGroupingOperation(ctx *GroupingOperationContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#informationFunction.
	VisitInformationFunction(ctx *InformationFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#specialDateTime.
	VisitSpecialDateTime(ctx *SpecialDateTimeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#specialFunction.
	VisitSpecialFunction(ctx *SpecialFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#aggregationFunctionCall.
	VisitAggregationFunctionCall(ctx *AggregationFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#windowFunctionCall.
	VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#translateFunctionCall.
	VisitTranslateFunctionCall(ctx *TranslateFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#simpleFunctionCall.
	VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#aggregationFunction.
	VisitAggregationFunction(ctx *AggregationFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#userVariable.
	VisitUserVariable(ctx *UserVariableContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#systemVariable.
	VisitSystemVariable(ctx *SystemVariableContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#columnReference.
	VisitColumnReference(ctx *ColumnReferenceContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#informationFunctionExpression.
	VisitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#specialDateTimeExpression.
	VisitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#specialFunctionExpression.
	VisitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#windowFunction.
	VisitWindowFunction(ctx *WindowFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#whenClause.
	VisitWhenClause(ctx *WhenClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#over.
	VisitOver(ctx *OverContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#ignoreNulls.
	VisitIgnoreNulls(ctx *IgnoreNullsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#windowFrame.
	VisitWindowFrame(ctx *WindowFrameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#unboundedFrame.
	VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#currentRowBound.
	VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#boundedFrame.
	VisitBoundedFrame(ctx *BoundedFrameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#backupRestoreObjectDesc.
	VisitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tableDesc.
	VisitTableDesc(ctx *TableDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#backupRestoreTableDesc.
	VisitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#explainDesc.
	VisitExplainDesc(ctx *ExplainDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#optimizerTrace.
	VisitOptimizerTrace(ctx *OptimizerTraceContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionExpr.
	VisitPartitionExpr(ctx *PartitionExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionDesc.
	VisitPartitionDesc(ctx *PartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#listPartitionDesc.
	VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#singleItemListPartitionDesc.
	VisitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#multiItemListPartitionDesc.
	VisitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#multiListPartitionValues.
	VisitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#singleListPartitionValues.
	VisitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#listPartitionValues.
	VisitListPartitionValues(ctx *ListPartitionValuesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#listPartitionValue.
	VisitListPartitionValue(ctx *ListPartitionValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#literalExpressionList.
	VisitLiteralExpressionList(ctx *LiteralExpressionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#rangePartitionDesc.
	VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#singleRangePartition.
	VisitSingleRangePartition(ctx *SingleRangePartitionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#multiRangePartition.
	VisitMultiRangePartition(ctx *MultiRangePartitionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionRangeDesc.
	VisitPartitionRangeDesc(ctx *PartitionRangeDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionKeyDesc.
	VisitPartitionKeyDesc(ctx *PartitionKeyDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionValueList.
	VisitPartitionValueList(ctx *PartitionValueListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#keyPartition.
	VisitKeyPartition(ctx *KeyPartitionContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#partitionValue.
	VisitPartitionValue(ctx *PartitionValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#distributionClause.
	VisitDistributionClause(ctx *DistributionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#distributionDesc.
	VisitDistributionDesc(ctx *DistributionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#refreshSchemeDesc.
	VisitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#statusDesc.
	VisitStatusDesc(ctx *StatusDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#extProperties.
	VisitExtProperties(ctx *ExtPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#propertyList.
	VisitPropertyList(ctx *PropertyListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#userPropertyList.
	VisitUserPropertyList(ctx *UserPropertyListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#inlineProperties.
	VisitInlineProperties(ctx *InlinePropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#inlineProperty.
	VisitInlineProperty(ctx *InlinePropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#varType.
	VisitVarType(ctx *VarTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#outfile.
	VisitOutfile(ctx *OutfileContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#fileFormat.
	VisitFileFormat(ctx *FileFormatContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#taskInterval.
	VisitTaskInterval(ctx *TaskIntervalContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#taskUnitIdentifier.
	VisitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#unitIdentifier.
	VisitUnitIdentifier(ctx *UnitIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#unitBoundary.
	VisitUnitBoundary(ctx *UnitBoundaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#subfieldDesc.
	VisitSubfieldDesc(ctx *SubfieldDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#subfieldDescs.
	VisitSubfieldDescs(ctx *SubfieldDescsContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#decimalType.
	VisitDecimalType(ctx *DecimalTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#writeBranch.
	VisitWriteBranch(ctx *WriteBranchContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#unquotedIdentifier.
	VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#digitIdentifier.
	VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#backQuotedIdentifier.
	VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#identifierWithAlias.
	VisitIdentifierWithAlias(ctx *IdentifierWithAliasContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#identifierWithAliasList.
	VisitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#identifierOrString.
	VisitIdentifierOrString(ctx *IdentifierOrStringContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#identifierOrStringList.
	VisitIdentifierOrStringList(ctx *IdentifierOrStringListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#identifierOrStringOrStar.
	VisitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#userWithoutHost.
	VisitUserWithoutHost(ctx *UserWithoutHostContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#userWithHost.
	VisitUserWithHost(ctx *UserWithHostContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#userWithHostAndBlanket.
	VisitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#assignmentList.
	VisitAssignmentList(ctx *AssignmentListContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#decimalValue.
	VisitDecimalValue(ctx *DecimalValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#doubleValue.
	VisitDoubleValue(ctx *DoubleValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#integerValue.
	VisitIntegerValue(ctx *IntegerValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParser#nonReserved.
	VisitNonReserved(ctx *NonReservedContext) interface{}
}
