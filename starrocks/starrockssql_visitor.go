// Code generated from StarRocksSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package starrocks // StarRocksSQL
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by StarRocksSQLParser.
type StarRocksSQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by StarRocksSQLParser#sqlStatements.
	VisitSqlStatements(ctx *SqlStatementsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#singleStatement.
	VisitSingleStatement(ctx *SingleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#useDatabaseStatement.
	VisitUseDatabaseStatement(ctx *UseDatabaseStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#useCatalogStatement.
	VisitUseCatalogStatement(ctx *UseCatalogStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setCatalogStatement.
	VisitSetCatalogStatement(ctx *SetCatalogStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDatabasesStatement.
	VisitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterDbQuotaStatement.
	VisitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createDbStatement.
	VisitCreateDbStatement(ctx *CreateDbStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropDbStatement.
	VisitDropDbStatement(ctx *DropDbStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateDbStatement.
	VisitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterDatabaseRenameStatement.
	VisitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#recoverDbStmt.
	VisitRecoverDbStmt(ctx *RecoverDbStmtContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDataStmt.
	VisitShowDataStmt(ctx *ShowDataStmtContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDataDistributionStmt.
	VisitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnDesc.
	VisitColumnDesc(ctx *ColumnDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#defaultDesc.
	VisitDefaultDesc(ctx *DefaultDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#generatedColumnDesc.
	VisitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#indexDesc.
	VisitIndexDesc(ctx *IndexDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#engineDesc.
	VisitEngineDesc(ctx *EngineDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#charsetDesc.
	VisitCharsetDesc(ctx *CharsetDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#collateDesc.
	VisitCollateDesc(ctx *CollateDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#keyDesc.
	VisitKeyDesc(ctx *KeyDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#orderByDesc.
	VisitOrderByDesc(ctx *OrderByDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnNullable.
	VisitColumnNullable(ctx *ColumnNullableContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#typeWithNullable.
	VisitTypeWithNullable(ctx *TypeWithNullableContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#aggStateDesc.
	VisitAggStateDesc(ctx *AggStateDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#aggDesc.
	VisitAggDesc(ctx *AggDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rollupDesc.
	VisitRollupDesc(ctx *RollupDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rollupItem.
	VisitRollupItem(ctx *RollupItemContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dupKeys.
	VisitDupKeys(ctx *DupKeysContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#fromRollup.
	VisitFromRollup(ctx *FromRollupContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#orReplace.
	VisitOrReplace(ctx *OrReplaceContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createTableAsSelectStatement.
	VisitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropTableStatement.
	VisitDropTableStatement(ctx *DropTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cleanTemporaryTableStatement.
	VisitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterTableStatement.
	VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createIndexStatement.
	VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropIndexStatement.
	VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTableStatement.
	VisitShowTableStatement(ctx *ShowTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTemporaryTablesStatement.
	VisitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateTableStatement.
	VisitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showColumnStatement.
	VisitShowColumnStatement(ctx *ShowColumnStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTableStatusStatement.
	VisitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#refreshTableStatement.
	VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showAlterStatement.
	VisitShowAlterStatement(ctx *ShowAlterStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#descTableStatement.
	VisitDescTableStatement(ctx *DescTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createTableLikeStatement.
	VisitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showIndexStatement.
	VisitShowIndexStatement(ctx *ShowIndexStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#recoverTableStatement.
	VisitRecoverTableStatement(ctx *RecoverTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#truncateTableStatement.
	VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelAlterTableStatement.
	VisitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showPartitionsStatement.
	VisitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#recoverPartitionStatement.
	VisitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createViewStatement.
	VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterViewStatement.
	VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropViewStatement.
	VisitDropViewStatement(ctx *DropViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnNameWithComment.
	VisitColumnNameWithComment(ctx *ColumnNameWithCommentContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#submitTaskStatement.
	VisitSubmitTaskStatement(ctx *SubmitTaskStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#taskClause.
	VisitTaskClause(ctx *TaskClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropTaskStatement.
	VisitDropTaskStatement(ctx *DropTaskStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#taskScheduleDesc.
	VisitTaskScheduleDesc(ctx *TaskScheduleDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createMaterializedViewStatement.
	VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#mvPartitionExprs.
	VisitMvPartitionExprs(ctx *MvPartitionExprsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#materializedViewDesc.
	VisitMaterializedViewDesc(ctx *MaterializedViewDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showMaterializedViewsStatement.
	VisitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropMaterializedViewStatement.
	VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterMaterializedViewStatement.
	VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#refreshMaterializedViewStatement.
	VisitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelRefreshMaterializedViewStatement.
	VisitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminSetConfigStatement.
	VisitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminSetReplicaStatusStatement.
	VisitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminShowConfigStatement.
	VisitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminShowReplicaDistributionStatement.
	VisitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminShowReplicaStatusStatement.
	VisitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminRepairTableStatement.
	VisitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminCancelRepairTableStatement.
	VisitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminCheckTabletsStatement.
	VisitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminSetPartitionVersion.
	VisitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#killStatement.
	VisitKillStatement(ctx *KillStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#syncStatement.
	VisitSyncStatement(ctx *SyncStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminSetAutomatedSnapshotOnStatement.
	VisitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#adminSetAutomatedSnapshotOffStatement.
	VisitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterSystemStatement.
	VisitAlterSystemStatement(ctx *AlterSystemStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelAlterSystemStatement.
	VisitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showComputeNodesStatement.
	VisitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createExternalCatalogStatement.
	VisitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateExternalCatalogStatement.
	VisitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropExternalCatalogStatement.
	VisitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCatalogsStatement.
	VisitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterCatalogStatement.
	VisitAlterCatalogStatement(ctx *AlterCatalogStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createStorageVolumeStatement.
	VisitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#typeDesc.
	VisitTypeDesc(ctx *TypeDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#locationsDesc.
	VisitLocationsDesc(ctx *LocationsDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showStorageVolumesStatement.
	VisitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropStorageVolumeStatement.
	VisitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterStorageVolumeStatement.
	VisitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterStorageVolumeClause.
	VisitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyStorageVolumePropertiesClause.
	VisitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyStorageVolumeCommentClause.
	VisitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#descStorageVolumeStatement.
	VisitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setDefaultStorageVolumeStatement.
	VisitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#updateFailPointStatusStatement.
	VisitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showFailPointStatement.
	VisitShowFailPointStatement(ctx *ShowFailPointStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createDictionaryStatement.
	VisitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropDictionaryStatement.
	VisitDropDictionaryStatement(ctx *DropDictionaryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#refreshDictionaryStatement.
	VisitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDictionaryStatement.
	VisitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelRefreshDictionaryStatement.
	VisitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dictionaryColumnDesc.
	VisitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dictionaryName.
	VisitDictionaryName(ctx *DictionaryNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterClause.
	VisitAlterClause(ctx *AlterClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addFrontendClause.
	VisitAddFrontendClause(ctx *AddFrontendClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropFrontendClause.
	VisitDropFrontendClause(ctx *DropFrontendClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyFrontendHostClause.
	VisitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addBackendClause.
	VisitAddBackendClause(ctx *AddBackendClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropBackendClause.
	VisitDropBackendClause(ctx *DropBackendClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#decommissionBackendClause.
	VisitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyBackendClause.
	VisitModifyBackendClause(ctx *ModifyBackendClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addComputeNodeClause.
	VisitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropComputeNodeClause.
	VisitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyBrokerClause.
	VisitModifyBrokerClause(ctx *ModifyBrokerClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterLoadErrorUrlClause.
	VisitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createImageClause.
	VisitCreateImageClause(ctx *CreateImageClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cleanTabletSchedQClause.
	VisitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#decommissionDiskClause.
	VisitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelDecommissionDiskClause.
	VisitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#disableDiskClause.
	VisitDisableDiskClause(ctx *DisableDiskClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelDisableDiskClause.
	VisitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createIndexClause.
	VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropIndexClause.
	VisitDropIndexClause(ctx *DropIndexClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tableRenameClause.
	VisitTableRenameClause(ctx *TableRenameClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#swapTableClause.
	VisitSwapTableClause(ctx *SwapTableClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyPropertiesClause.
	VisitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyCommentClause.
	VisitModifyCommentClause(ctx *ModifyCommentClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#optimizeRange.
	VisitOptimizeRange(ctx *OptimizeRangeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#optimizeClause.
	VisitOptimizeClause(ctx *OptimizeClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addColumnClause.
	VisitAddColumnClause(ctx *AddColumnClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addColumnsClause.
	VisitAddColumnsClause(ctx *AddColumnsClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropColumnClause.
	VisitDropColumnClause(ctx *DropColumnClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyColumnClause.
	VisitModifyColumnClause(ctx *ModifyColumnClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyColumnCommentClause.
	VisitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnRenameClause.
	VisitColumnRenameClause(ctx *ColumnRenameClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#reorderColumnsClause.
	VisitReorderColumnsClause(ctx *ReorderColumnsClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rollupRenameClause.
	VisitRollupRenameClause(ctx *RollupRenameClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#compactionClause.
	VisitCompactionClause(ctx *CompactionClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#subfieldName.
	VisitSubfieldName(ctx *SubfieldNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#nestedFieldName.
	VisitNestedFieldName(ctx *NestedFieldNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addFieldClause.
	VisitAddFieldClause(ctx *AddFieldClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropFieldClause.
	VisitDropFieldClause(ctx *DropFieldClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createOrReplaceTagClause.
	VisitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createOrReplaceBranchClause.
	VisitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropBranchClause.
	VisitDropBranchClause(ctx *DropBranchClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropTagClause.
	VisitDropTagClause(ctx *DropTagClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tableOperationClause.
	VisitTableOperationClause(ctx *TableOperationClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tagOptions.
	VisitTagOptions(ctx *TagOptionsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#branchOptions.
	VisitBranchOptions(ctx *BranchOptionsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#snapshotRetention.
	VisitSnapshotRetention(ctx *SnapshotRetentionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#refRetain.
	VisitRefRetain(ctx *RefRetainContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#maxSnapshotAge.
	VisitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#minSnapshotsToKeep.
	VisitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#snapshotId.
	VisitSnapshotId(ctx *SnapshotIdContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#timeUnit.
	VisitTimeUnit(ctx *TimeUnitContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#integer_list.
	VisitInteger_list(ctx *Integer_listContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropPersistentIndexClause.
	VisitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addPartitionClause.
	VisitAddPartitionClause(ctx *AddPartitionClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropPartitionClause.
	VisitDropPartitionClause(ctx *DropPartitionClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#truncatePartitionClause.
	VisitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#modifyPartitionClause.
	VisitModifyPartitionClause(ctx *ModifyPartitionClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#replacePartitionClause.
	VisitReplacePartitionClause(ctx *ReplacePartitionClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionRenameClause.
	VisitPartitionRenameClause(ctx *PartitionRenameClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#insertLabelOrColumnAliases.
	VisitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnAliasesOrByName.
	VisitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createRoutineLoadStatement.
	VisitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterRoutineLoadStatement.
	VisitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dataSource.
	VisitDataSource(ctx *DataSourceContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#loadProperties.
	VisitLoadProperties(ctx *LoadPropertiesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#colSeparatorProperty.
	VisitColSeparatorProperty(ctx *ColSeparatorPropertyContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rowDelimiterProperty.
	VisitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#importColumns.
	VisitImportColumns(ctx *ImportColumnsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnProperties.
	VisitColumnProperties(ctx *ColumnPropertiesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#jobProperties.
	VisitJobProperties(ctx *JobPropertiesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dataSourceProperties.
	VisitDataSourceProperties(ctx *DataSourcePropertiesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#stopRoutineLoadStatement.
	VisitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#resumeRoutineLoadStatement.
	VisitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#pauseRoutineLoadStatement.
	VisitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showRoutineLoadStatement.
	VisitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showRoutineLoadTaskStatement.
	VisitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateRoutineLoadStatement.
	VisitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showStreamLoadStatement.
	VisitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#analyzeStatement.
	VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#regularColumns.
	VisitRegularColumns(ctx *RegularColumnsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#allColumns.
	VisitAllColumns(ctx *AllColumnsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#predicateColumns.
	VisitPredicateColumns(ctx *PredicateColumnsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#multiColumnSet.
	VisitMultiColumnSet(ctx *MultiColumnSetContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropStatsStatement.
	VisitDropStatsStatement(ctx *DropStatsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#histogramStatement.
	VisitHistogramStatement(ctx *HistogramStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#analyzeHistogramStatement.
	VisitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropHistogramStatement.
	VisitDropHistogramStatement(ctx *DropHistogramStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createAnalyzeStatement.
	VisitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropAnalyzeJobStatement.
	VisitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showAnalyzeStatement.
	VisitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showStatsMetaStatement.
	VisitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showHistogramMetaStatement.
	VisitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#killAnalyzeStatement.
	VisitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#analyzeProfileStatement.
	VisitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createBaselinePlanStatement.
	VisitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropBaselinePlanStatement.
	VisitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showBaselinePlanStatement.
	VisitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createResourceGroupStatement.
	VisitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropResourceGroupStatement.
	VisitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterResourceGroupStatement.
	VisitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showResourceGroupStatement.
	VisitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showResourceGroupUsageStatement.
	VisitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createResourceStatement.
	VisitCreateResourceStatement(ctx *CreateResourceStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterResourceStatement.
	VisitAlterResourceStatement(ctx *AlterResourceStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropResourceStatement.
	VisitDropResourceStatement(ctx *DropResourceStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showResourceStatement.
	VisitShowResourceStatement(ctx *ShowResourceStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#classifier.
	VisitClassifier(ctx *ClassifierContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showFunctionsStatement.
	VisitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropFunctionStatement.
	VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createFunctionStatement.
	VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#inlineFunction.
	VisitInlineFunction(ctx *InlineFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#loadStatement.
	VisitLoadStatement(ctx *LoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#labelName.
	VisitLabelName(ctx *LabelNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dataDescList.
	VisitDataDescList(ctx *DataDescListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dataDesc.
	VisitDataDesc(ctx *DataDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#formatProps.
	VisitFormatProps(ctx *FormatPropsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#brokerDesc.
	VisitBrokerDesc(ctx *BrokerDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#resourceDesc.
	VisitResourceDesc(ctx *ResourceDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showLoadStatement.
	VisitShowLoadStatement(ctx *ShowLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showLoadWarningsStatement.
	VisitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelLoadStatement.
	VisitCancelLoadStatement(ctx *CancelLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterLoadStatement.
	VisitAlterLoadStatement(ctx *AlterLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelCompactionStatement.
	VisitCancelCompactionStatement(ctx *CancelCompactionStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showAuthorStatement.
	VisitShowAuthorStatement(ctx *ShowAuthorStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showBackendsStatement.
	VisitShowBackendsStatement(ctx *ShowBackendsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showBrokerStatement.
	VisitShowBrokerStatement(ctx *ShowBrokerStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCharsetStatement.
	VisitShowCharsetStatement(ctx *ShowCharsetStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCollationStatement.
	VisitShowCollationStatement(ctx *ShowCollationStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDeleteStatement.
	VisitShowDeleteStatement(ctx *ShowDeleteStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDynamicPartitionStatement.
	VisitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showEventsStatement.
	VisitShowEventsStatement(ctx *ShowEventsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showEnginesStatement.
	VisitShowEnginesStatement(ctx *ShowEnginesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showFrontendsStatement.
	VisitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showFrontendsDisksStatement.
	VisitShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showPluginsStatement.
	VisitShowPluginsStatement(ctx *ShowPluginsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showRepositoriesStatement.
	VisitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showOpenTableStatement.
	VisitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showPrivilegesStatement.
	VisitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showProcedureStatement.
	VisitShowProcedureStatement(ctx *ShowProcedureStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showProcStatement.
	VisitShowProcStatement(ctx *ShowProcStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showProcesslistStatement.
	VisitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showProfilelistStatement.
	VisitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showRunningQueriesStatement.
	VisitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showStatusStatement.
	VisitShowStatusStatement(ctx *ShowStatusStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTabletStatement.
	VisitShowTabletStatement(ctx *ShowTabletStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTransactionStatement.
	VisitShowTransactionStatement(ctx *ShowTransactionStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTriggersStatement.
	VisitShowTriggersStatement(ctx *ShowTriggersStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showUserPropertyStatement.
	VisitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showVariablesStatement.
	VisitShowVariablesStatement(ctx *ShowVariablesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showWarningStatement.
	VisitShowWarningStatement(ctx *ShowWarningStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#helpStatement.
	VisitHelpStatement(ctx *HelpStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showQueryProfileStatement.
	VisitShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showQueryStatsStatement.
	VisitShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showLoadProfileStatement.
	VisitShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDataSkewStatement.
	VisitShowDataSkewStatement(ctx *ShowDataSkewStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDataTypesStatement.
	VisitShowDataTypesStatement(ctx *ShowDataTypesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showSyncJobStatement.
	VisitShowSyncJobStatement(ctx *ShowSyncJobStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showPolicyStatement.
	VisitShowPolicyStatement(ctx *ShowPolicyStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showSqlBlockRuleStatement.
	VisitShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showEncryptKeysStatement.
	VisitShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateLoadStatement.
	VisitShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateRepositoryStatement.
	VisitShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showLastInsertStatement.
	VisitShowLastInsertStatement(ctx *ShowLastInsertStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTableIdStatement.
	VisitShowTableIdStatement(ctx *ShowTableIdStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDatabaseIdStatement.
	VisitShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showPartitionIdStatement.
	VisitShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTableStatsStatement.
	VisitShowTableStatsStatement(ctx *ShowTableStatsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showColumnStatsStatement.
	VisitShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showConvertLightSchemaChangeStatement.
	VisitShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCatalogRecycleBinStatement.
	VisitShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showTrashStatement.
	VisitShowTrashStatement(ctx *ShowTrashStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showMigrationsStatement.
	VisitShowMigrationsStatement(ctx *ShowMigrationsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showWorkloadGroupsStatement.
	VisitShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showJobTaskStatement.
	VisitShowJobTaskStatement(ctx *ShowJobTaskStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createUserStatement.
	VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropUserStatement.
	VisitDropUserStatement(ctx *DropUserStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterUserStatement.
	VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showUserStatement.
	VisitShowUserStatement(ctx *ShowUserStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showAllAuthentication.
	VisitShowAllAuthentication(ctx *ShowAllAuthenticationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showAuthenticationForUser.
	VisitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#executeAsStatement.
	VisitExecuteAsStatement(ctx *ExecuteAsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createRoleStatement.
	VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterRoleStatement.
	VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropRoleStatement.
	VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showRolesStatement.
	VisitShowRolesStatement(ctx *ShowRolesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantRoleToUser.
	VisitGrantRoleToUser(ctx *GrantRoleToUserContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantRoleToRole.
	VisitGrantRoleToRole(ctx *GrantRoleToRoleContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeRoleFromUser.
	VisitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeRoleFromRole.
	VisitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setRoleStatement.
	VisitSetRoleStatement(ctx *SetRoleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setDefaultRoleStatement.
	VisitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantRevokeClause.
	VisitGrantRevokeClause(ctx *GrantRevokeClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantOnUser.
	VisitGrantOnUser(ctx *GrantOnUserContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantOnTableBrief.
	VisitGrantOnTableBrief(ctx *GrantOnTableBriefContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantOnFunc.
	VisitGrantOnFunc(ctx *GrantOnFuncContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantOnSystem.
	VisitGrantOnSystem(ctx *GrantOnSystemContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantOnPrimaryObj.
	VisitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#grantOnAll.
	VisitGrantOnAll(ctx *GrantOnAllContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeOnUser.
	VisitRevokeOnUser(ctx *RevokeOnUserContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeOnTableBrief.
	VisitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeOnFunc.
	VisitRevokeOnFunc(ctx *RevokeOnFuncContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeOnSystem.
	VisitRevokeOnSystem(ctx *RevokeOnSystemContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeOnPrimaryObj.
	VisitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#revokeOnAll.
	VisitRevokeOnAll(ctx *RevokeOnAllContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showGrantsStatement.
	VisitShowGrantsStatement(ctx *ShowGrantsStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#authWithoutPlugin.
	VisitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#authWithPlugin.
	VisitAuthWithPlugin(ctx *AuthWithPluginContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#privObjectName.
	VisitPrivObjectName(ctx *PrivObjectNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#privObjectNameList.
	VisitPrivObjectNameList(ctx *PrivObjectNameListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#privFunctionObjectNameList.
	VisitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#privilegeTypeList.
	VisitPrivilegeTypeList(ctx *PrivilegeTypeListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#privilegeType.
	VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#privObjectType.
	VisitPrivObjectType(ctx *PrivObjectTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#privObjectTypePlural.
	VisitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createSecurityIntegrationStatement.
	VisitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterSecurityIntegrationStatement.
	VisitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropSecurityIntegrationStatement.
	VisitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showSecurityIntegrationStatement.
	VisitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateSecurityIntegrationStatement.
	VisitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createGroupProviderStatement.
	VisitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropGroupProviderStatement.
	VisitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showGroupProvidersStatement.
	VisitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showCreateGroupProviderStatement.
	VisitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#backupStatement.
	VisitBackupStatement(ctx *BackupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelBackupStatement.
	VisitCancelBackupStatement(ctx *CancelBackupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showBackupStatement.
	VisitShowBackupStatement(ctx *ShowBackupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#restoreStatement.
	VisitRestoreStatement(ctx *RestoreStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelRestoreStatement.
	VisitCancelRestoreStatement(ctx *CancelRestoreStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showRestoreStatement.
	VisitShowRestoreStatement(ctx *ShowRestoreStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showSnapshotStatement.
	VisitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createRepositoryStatement.
	VisitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropRepositoryStatement.
	VisitDropRepositoryStatement(ctx *DropRepositoryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addSqlBlackListStatement.
	VisitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#delSqlBlackListStatement.
	VisitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showSqlBlackListStatement.
	VisitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showWhiteListStatement.
	VisitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#addBackendBlackListStatement.
	VisitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#delBackendBlackListStatement.
	VisitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showBackendBlackListStatement.
	VisitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dataCacheTarget.
	VisitDataCacheTarget(ctx *DataCacheTargetContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createDataCacheRuleStatement.
	VisitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showDataCacheRulesStatement.
	VisitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropDataCacheRuleStatement.
	VisitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#clearDataCacheRulesStatement.
	VisitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dataCacheSelectStatement.
	VisitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cancelExportStatement.
	VisitCancelExportStatement(ctx *CancelExportStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showExportStatement.
	VisitShowExportStatement(ctx *ShowExportStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#installPluginStatement.
	VisitInstallPluginStatement(ctx *InstallPluginStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#uninstallPluginStatement.
	VisitUninstallPluginStatement(ctx *UninstallPluginStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createFileStatement.
	VisitCreateFileStatement(ctx *CreateFileStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropFileStatement.
	VisitDropFileStatement(ctx *DropFileStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showSmallFilesStatement.
	VisitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createPipeStatement.
	VisitCreatePipeStatement(ctx *CreatePipeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropPipeStatement.
	VisitDropPipeStatement(ctx *DropPipeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterPipeClause.
	VisitAlterPipeClause(ctx *AlterPipeClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterPipeStatement.
	VisitAlterPipeStatement(ctx *AlterPipeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#descPipeStatement.
	VisitDescPipeStatement(ctx *DescPipeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showPipeStatement.
	VisitShowPipeStatement(ctx *ShowPipeStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setStatement.
	VisitSetStatement(ctx *SetStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setNames.
	VisitSetNames(ctx *SetNamesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setPassword.
	VisitSetPassword(ctx *SetPasswordContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setUserVar.
	VisitSetUserVar(ctx *SetUserVarContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setSystemVar.
	VisitSetSystemVar(ctx *SetSystemVarContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#transaction_characteristics.
	VisitTransaction_characteristics(ctx *Transaction_characteristicsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#transaction_access_mode.
	VisitTransaction_access_mode(ctx *Transaction_access_modeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#isolation_level.
	VisitIsolation_level(ctx *Isolation_levelContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#isolation_types.
	VisitIsolation_types(ctx *Isolation_typesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setExprOrDefault.
	VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setUserPropertyStatement.
	VisitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#roleList.
	VisitRoleList(ctx *RoleListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#executeScriptStatement.
	VisitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#unsupportedStatement.
	VisitUnsupportedStatement(ctx *UnsupportedStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#lock_item.
	VisitLock_item(ctx *Lock_itemContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#lock_type.
	VisitLock_type(ctx *Lock_typeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterPlanAdvisorAddStatement.
	VisitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#truncatePlanAdvisorStatement.
	VisitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterPlanAdvisorDropStatement.
	VisitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showPlanAdvisorStatement.
	VisitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createWarehouseStatement.
	VisitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropWarehouseStatement.
	VisitDropWarehouseStatement(ctx *DropWarehouseStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#suspendWarehouseStatement.
	VisitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#resumeWarehouseStatement.
	VisitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setWarehouseStatement.
	VisitSetWarehouseStatement(ctx *SetWarehouseStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showWarehousesStatement.
	VisitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showClustersStatement.
	VisitShowClustersStatement(ctx *ShowClustersStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#showNodesStatement.
	VisitShowNodesStatement(ctx *ShowNodesStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterWarehouseStatement.
	VisitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#createCNGroupStatement.
	VisitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dropCNGroupStatement.
	VisitDropCNGroupStatement(ctx *DropCNGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#enableCNGroupStatement.
	VisitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#disableCNGroupStatement.
	VisitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#alterCNGroupStatement.
	VisitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#beginStatement.
	VisitBeginStatement(ctx *BeginStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#commitStatement.
	VisitCommitStatement(ctx *CommitStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rollbackStatement.
	VisitRollbackStatement(ctx *RollbackStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#translateStatement.
	VisitTranslateStatement(ctx *TranslateStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dialect.
	VisitDialect(ctx *DialectContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#translateSQL.
	VisitTranslateSQL(ctx *TranslateSQLContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#queryStatement.
	VisitQueryStatement(ctx *QueryStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#queryRelation.
	VisitQueryRelation(ctx *QueryRelationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#queryNoWith.
	VisitQueryNoWith(ctx *QueryNoWithContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#queryPeriod.
	VisitQueryPeriod(ctx *QueryPeriodContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#periodType.
	VisitPeriodType(ctx *PeriodTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#queryWithParentheses.
	VisitQueryWithParentheses(ctx *QueryWithParenthesesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setOperation.
	VisitSetOperation(ctx *SetOperationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#queryPrimaryDefault.
	VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rowConstructor.
	VisitRowConstructor(ctx *RowConstructorContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#sortItem.
	VisitSortItem(ctx *SortItemContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#limitConstExpr.
	VisitLimitConstExpr(ctx *LimitConstExprContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#limitElement.
	VisitLimitElement(ctx *LimitElementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#from.
	VisitFrom(ctx *FromContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dual.
	VisitDual(ctx *DualContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rollup.
	VisitRollup(ctx *RollupContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cube.
	VisitCube(ctx *CubeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#multipleGroupingSets.
	VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#singleGroupingSet.
	VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#groupingSet.
	VisitGroupingSet(ctx *GroupingSetContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#commonTableExpression.
	VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#setQuantifier.
	VisitSetQuantifier(ctx *SetQuantifierContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#selectSingle.
	VisitSelectSingle(ctx *SelectSingleContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#selectAll.
	VisitSelectAll(ctx *SelectAllContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#excludeClause.
	VisitExcludeClause(ctx *ExcludeClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#relations.
	VisitRelations(ctx *RelationsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#relationLateralView.
	VisitRelationLateralView(ctx *RelationLateralViewContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#lateralView.
	VisitLateralView(ctx *LateralViewContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#generatorFunction.
	VisitGeneratorFunction(ctx *GeneratorFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tableAtom.
	VisitTableAtom(ctx *TableAtomContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#inlineTable.
	VisitInlineTable(ctx *InlineTableContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#subqueryWithAlias.
	VisitSubqueryWithAlias(ctx *SubqueryWithAliasContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tableFunction.
	VisitTableFunction(ctx *TableFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#normalizedTableFunction.
	VisitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#fileTableFunction.
	VisitFileTableFunction(ctx *FileTableFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#parenthesizedRelation.
	VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#pivotClause.
	VisitPivotClause(ctx *PivotClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#pivotAggregationExpression.
	VisitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#pivotValue.
	VisitPivotValue(ctx *PivotValueContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#sampleClause.
	VisitSampleClause(ctx *SampleClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#namedArgumentList.
	VisitNamedArgumentList(ctx *NamedArgumentListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#namedArguments.
	VisitNamedArguments(ctx *NamedArgumentsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#joinRelation.
	VisitJoinRelation(ctx *JoinRelationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#crossOrInnerJoinType.
	VisitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#outerAndSemiJoinType.
	VisitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#bracketHint.
	VisitBracketHint(ctx *BracketHintContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#hintMap.
	VisitHintMap(ctx *HintMapContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#joinCriteria.
	VisitJoinCriteria(ctx *JoinCriteriaContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnAliases.
	VisitColumnAliases(ctx *ColumnAliasesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnAliasesWithoutParentheses.
	VisitColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionNames.
	VisitPartitionNames(ctx *PartitionNamesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#keyPartitionList.
	VisitKeyPartitionList(ctx *KeyPartitionListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tabletList.
	VisitTabletList(ctx *TabletListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#prepareStatement.
	VisitPrepareStatement(ctx *PrepareStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#prepareSql.
	VisitPrepareSql(ctx *PrepareSqlContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#deallocateStatement.
	VisitDeallocateStatement(ctx *DeallocateStatementContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#replicaList.
	VisitReplicaList(ctx *ReplicaListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#expressionsWithDefault.
	VisitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#expressionOrDefault.
	VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#mapExpressionList.
	VisitMapExpressionList(ctx *MapExpressionListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#mapExpression.
	VisitMapExpression(ctx *MapExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#expressionSingleton.
	VisitExpressionSingleton(ctx *ExpressionSingletonContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#expressionDefault.
	VisitExpressionDefault(ctx *ExpressionDefaultContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#logicalNot.
	VisitLogicalNot(ctx *LogicalNotContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#logicalBinary.
	VisitLogicalBinary(ctx *LogicalBinaryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#booleanExpressionDefault.
	VisitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#isNull.
	VisitIsNull(ctx *IsNullContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#scalarSubquery.
	VisitScalarSubquery(ctx *ScalarSubqueryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tupleInSubquery.
	VisitTupleInSubquery(ctx *TupleInSubqueryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#inSubquery.
	VisitInSubquery(ctx *InSubqueryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#inList.
	VisitInList(ctx *InListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#between.
	VisitBetween(ctx *BetweenContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#like.
	VisitLike(ctx *LikeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#valueExpressionDefault.
	VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#arithmeticBinary.
	VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dereference.
	VisitDereference(ctx *DereferenceContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#odbcFunctionCallExpression.
	VisitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#matchExpr.
	VisitMatchExpr(ctx *MatchExprContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnRef.
	VisitColumnRef(ctx *ColumnRefContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#convert.
	VisitConvert(ctx *ConvertContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#collectionSubscript.
	VisitCollectionSubscript(ctx *CollectionSubscriptContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#cast.
	VisitCast(ctx *CastContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#userVariableExpression.
	VisitUserVariableExpression(ctx *UserVariableExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#simpleCase.
	VisitSimpleCase(ctx *SimpleCaseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#arrowExpression.
	VisitArrowExpression(ctx *ArrowExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#arrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#systemVariableExpression.
	VisitSystemVariableExpression(ctx *SystemVariableExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#concat.
	VisitConcat(ctx *ConcatContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#subqueryExpression.
	VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#lambdaFunctionExpr.
	VisitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dictionaryGetExpr.
	VisitDictionaryGetExpr(ctx *DictionaryGetExprContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#collate.
	VisitCollate(ctx *CollateContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#arrayConstructor.
	VisitArrayConstructor(ctx *ArrayConstructorContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#mapConstructor.
	VisitMapConstructor(ctx *MapConstructorContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#arraySlice.
	VisitArraySlice(ctx *ArraySliceContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#exists.
	VisitExists(ctx *ExistsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#searchedCase.
	VisitSearchedCase(ctx *SearchedCaseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#arithmeticUnary.
	VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#dateLiteral.
	VisitDateLiteral(ctx *DateLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#intervalLiteral.
	VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#unitBoundaryLiteral.
	VisitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#binaryLiteral.
	VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#Parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#extract.
	VisitExtract(ctx *ExtractContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#groupingOperation.
	VisitGroupingOperation(ctx *GroupingOperationContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#informationFunction.
	VisitInformationFunction(ctx *InformationFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#specialDateTime.
	VisitSpecialDateTime(ctx *SpecialDateTimeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#specialFunction.
	VisitSpecialFunction(ctx *SpecialFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#aggregationFunctionCall.
	VisitAggregationFunctionCall(ctx *AggregationFunctionCallContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#windowFunctionCall.
	VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#translateFunctionCall.
	VisitTranslateFunctionCall(ctx *TranslateFunctionCallContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#simpleFunctionCall.
	VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#aggregationFunction.
	VisitAggregationFunction(ctx *AggregationFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#userVariable.
	VisitUserVariable(ctx *UserVariableContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#systemVariable.
	VisitSystemVariable(ctx *SystemVariableContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#columnReference.
	VisitColumnReference(ctx *ColumnReferenceContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#informationFunctionExpression.
	VisitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#specialDateTimeExpression.
	VisitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#specialFunctionExpression.
	VisitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#windowFunction.
	VisitWindowFunction(ctx *WindowFunctionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#whenClause.
	VisitWhenClause(ctx *WhenClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#over.
	VisitOver(ctx *OverContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#ignoreNulls.
	VisitIgnoreNulls(ctx *IgnoreNullsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#windowFrame.
	VisitWindowFrame(ctx *WindowFrameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#unboundedFrame.
	VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#currentRowBound.
	VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#boundedFrame.
	VisitBoundedFrame(ctx *BoundedFrameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#backupRestoreObjectDesc.
	VisitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tableDesc.
	VisitTableDesc(ctx *TableDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#backupRestoreTableDesc.
	VisitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#explainDesc.
	VisitExplainDesc(ctx *ExplainDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#optimizerTrace.
	VisitOptimizerTrace(ctx *OptimizerTraceContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionExpr.
	VisitPartitionExpr(ctx *PartitionExprContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionDesc.
	VisitPartitionDesc(ctx *PartitionDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#listPartitionDesc.
	VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#singleItemListPartitionDesc.
	VisitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#multiItemListPartitionDesc.
	VisitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#multiListPartitionValues.
	VisitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#singleListPartitionValues.
	VisitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#listPartitionValues.
	VisitListPartitionValues(ctx *ListPartitionValuesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#listPartitionValue.
	VisitListPartitionValue(ctx *ListPartitionValueContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#literalExpressionList.
	VisitLiteralExpressionList(ctx *LiteralExpressionListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#rangePartitionDesc.
	VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#singleRangePartition.
	VisitSingleRangePartition(ctx *SingleRangePartitionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#multiRangePartition.
	VisitMultiRangePartition(ctx *MultiRangePartitionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionRangeDesc.
	VisitPartitionRangeDesc(ctx *PartitionRangeDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionKeyDesc.
	VisitPartitionKeyDesc(ctx *PartitionKeyDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionValueList.
	VisitPartitionValueList(ctx *PartitionValueListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#keyPartition.
	VisitKeyPartition(ctx *KeyPartitionContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#partitionValue.
	VisitPartitionValue(ctx *PartitionValueContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#distributionClause.
	VisitDistributionClause(ctx *DistributionClauseContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#distributionDesc.
	VisitDistributionDesc(ctx *DistributionDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#refreshSchemeDesc.
	VisitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#statusDesc.
	VisitStatusDesc(ctx *StatusDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#extProperties.
	VisitExtProperties(ctx *ExtPropertiesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#propertyList.
	VisitPropertyList(ctx *PropertyListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#userPropertyList.
	VisitUserPropertyList(ctx *UserPropertyListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#inlineProperties.
	VisitInlineProperties(ctx *InlinePropertiesContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#inlineProperty.
	VisitInlineProperty(ctx *InlinePropertyContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#varType.
	VisitVarType(ctx *VarTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#outfile.
	VisitOutfile(ctx *OutfileContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#fileFormat.
	VisitFileFormat(ctx *FileFormatContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#taskInterval.
	VisitTaskInterval(ctx *TaskIntervalContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#taskUnitIdentifier.
	VisitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#unitIdentifier.
	VisitUnitIdentifier(ctx *UnitIdentifierContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#unitBoundary.
	VisitUnitBoundary(ctx *UnitBoundaryContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#subfieldDesc.
	VisitSubfieldDesc(ctx *SubfieldDescContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#subfieldDescs.
	VisitSubfieldDescs(ctx *SubfieldDescsContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#decimalType.
	VisitDecimalType(ctx *DecimalTypeContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#writeBranch.
	VisitWriteBranch(ctx *WriteBranchContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#unquotedIdentifier.
	VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#digitIdentifier.
	VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#backQuotedIdentifier.
	VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#identifierWithAlias.
	VisitIdentifierWithAlias(ctx *IdentifierWithAliasContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#identifierWithAliasList.
	VisitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#identifierOrString.
	VisitIdentifierOrString(ctx *IdentifierOrStringContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#identifierOrStringList.
	VisitIdentifierOrStringList(ctx *IdentifierOrStringListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#identifierOrStringOrStar.
	VisitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#userWithoutHost.
	VisitUserWithoutHost(ctx *UserWithoutHostContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#userWithHost.
	VisitUserWithHost(ctx *UserWithHostContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#userWithHostAndBlanket.
	VisitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#assignmentList.
	VisitAssignmentList(ctx *AssignmentListContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#decimalValue.
	VisitDecimalValue(ctx *DecimalValueContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#doubleValue.
	VisitDoubleValue(ctx *DoubleValueContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#integerValue.
	VisitIntegerValue(ctx *IntegerValueContext) interface{}

	// Visit a parse tree produced by StarRocksSQLParser#nonReserved.
	VisitNonReserved(ctx *NonReservedContext) interface{}
}
