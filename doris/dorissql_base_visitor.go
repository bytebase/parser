// Code generated from DorisSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package doris // DorisSQL
import "github.com/antlr4-go/antlr/v4"

type BaseDorisSQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDorisSQLVisitor) VisitSqlStatements(ctx *SqlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUseDatabaseStatement(ctx *UseDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUseCatalogStatement(ctx *UseCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetCatalogStatement(ctx *SetCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateDbStatement(ctx *CreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropDbStatement(ctx *DropDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRecoverDbStmt(ctx *RecoverDbStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDataStmt(ctx *ShowDataStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnDesc(ctx *ColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDefaultDesc(ctx *DefaultDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIndexDesc(ctx *IndexDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitEngineDesc(ctx *EngineDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCharsetDesc(ctx *CharsetDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCollateDesc(ctx *CollateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitKeyDesc(ctx *KeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOrderByDesc(ctx *OrderByDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnNullable(ctx *ColumnNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTypeWithNullable(ctx *TypeWithNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAggStateDesc(ctx *AggStateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAggDesc(ctx *AggDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRollupDesc(ctx *RollupDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRollupItem(ctx *RollupItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDupKeys(ctx *DupKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitFromRollup(ctx *FromRollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOrReplace(ctx *OrReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropTableStatement(ctx *DropTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTableStatement(ctx *ShowTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowColumnStatement(ctx *ShowColumnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowAlterStatement(ctx *ShowAlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDescTableStatement(ctx *DescTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowIndexStatement(ctx *ShowIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRecoverTableStatement(ctx *RecoverTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropViewStatement(ctx *DropViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnNameWithComment(ctx *ColumnNameWithCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSubmitTaskStatement(ctx *SubmitTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTaskClause(ctx *TaskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropTaskStatement(ctx *DropTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTaskScheduleDesc(ctx *TaskScheduleDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMvPartitionExprs(ctx *MvPartitionExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMaterializedViewDesc(ctx *MaterializedViewDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitKillStatement(ctx *KillStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSyncStatement(ctx *SyncStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterSystemStatement(ctx *AlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterCatalogStatement(ctx *AlterCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTypeDesc(ctx *TypeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLocationsDesc(ctx *LocationsDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowFailPointStatement(ctx *ShowFailPointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropDictionaryStatement(ctx *DropDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDictionaryName(ctx *DictionaryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterClause(ctx *AlterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddFrontendClause(ctx *AddFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropFrontendClause(ctx *DropFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddBackendClause(ctx *AddBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropBackendClause(ctx *DropBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyBackendClause(ctx *ModifyBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyBrokerClause(ctx *ModifyBrokerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateImageClause(ctx *CreateImageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDisableDiskClause(ctx *DisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropIndexClause(ctx *DropIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTableRenameClause(ctx *TableRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSwapTableClause(ctx *SwapTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyCommentClause(ctx *ModifyCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOptimizeRange(ctx *OptimizeRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOptimizeClause(ctx *OptimizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddColumnClause(ctx *AddColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddColumnsClause(ctx *AddColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropColumnClause(ctx *DropColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyColumnClause(ctx *ModifyColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnRenameClause(ctx *ColumnRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitReorderColumnsClause(ctx *ReorderColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRollupRenameClause(ctx *RollupRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCompactionClause(ctx *CompactionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSubfieldName(ctx *SubfieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitNestedFieldName(ctx *NestedFieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddFieldClause(ctx *AddFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropFieldClause(ctx *DropFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropBranchClause(ctx *DropBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropTagClause(ctx *DropTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTableOperationClause(ctx *TableOperationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTagOptions(ctx *TagOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBranchOptions(ctx *BranchOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSnapshotRetention(ctx *SnapshotRetentionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRefRetain(ctx *RefRetainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSnapshotId(ctx *SnapshotIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTimeUnit(ctx *TimeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInteger_list(ctx *Integer_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddPartitionClause(ctx *AddPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropPartitionClause(ctx *DropPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitModifyPartitionClause(ctx *ModifyPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitReplacePartitionClause(ctx *ReplacePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionRenameClause(ctx *PartitionRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDataSource(ctx *DataSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLoadProperties(ctx *LoadPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColSeparatorProperty(ctx *ColSeparatorPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitImportColumns(ctx *ImportColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnProperties(ctx *ColumnPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitJobProperties(ctx *JobPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDataSourceProperties(ctx *DataSourcePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRegularColumns(ctx *RegularColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAllColumns(ctx *AllColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPredicateColumns(ctx *PredicateColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMultiColumnSet(ctx *MultiColumnSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropStatsStatement(ctx *DropStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitHistogramStatement(ctx *HistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropHistogramStatement(ctx *DropHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateResourceStatement(ctx *CreateResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterResourceStatement(ctx *AlterResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropResourceStatement(ctx *DropResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowResourceStatement(ctx *ShowResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitClassifier(ctx *ClassifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInlineFunction(ctx *InlineFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLabelName(ctx *LabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDataDescList(ctx *DataDescListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDataDesc(ctx *DataDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitFormatProps(ctx *FormatPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBrokerDesc(ctx *BrokerDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitResourceDesc(ctx *ResourceDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowLoadStatement(ctx *ShowLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelLoadStatement(ctx *CancelLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterLoadStatement(ctx *AlterLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelCompactionStatement(ctx *CancelCompactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowAuthorStatement(ctx *ShowAuthorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowBackendsStatement(ctx *ShowBackendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowBrokerStatement(ctx *ShowBrokerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCharsetStatement(ctx *ShowCharsetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCollationStatement(ctx *ShowCollationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDeleteStatement(ctx *ShowDeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowEventsStatement(ctx *ShowEventsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowEnginesStatement(ctx *ShowEnginesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowPluginsStatement(ctx *ShowPluginsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowProcedureStatement(ctx *ShowProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowProcStatement(ctx *ShowProcStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowStatusStatement(ctx *ShowStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTabletStatement(ctx *ShowTabletStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTransactionStatement(ctx *ShowTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTriggersStatement(ctx *ShowTriggersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowVariablesStatement(ctx *ShowVariablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowWarningStatement(ctx *ShowWarningStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitHelpStatement(ctx *HelpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDataSkewStatement(ctx *ShowDataSkewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDataTypesStatement(ctx *ShowDataTypesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowSyncJobStatement(ctx *ShowSyncJobStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowPolicyStatement(ctx *ShowPolicyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowLastInsertStatement(ctx *ShowLastInsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTableIdStatement(ctx *ShowTableIdStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTableStatsStatement(ctx *ShowTableStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowTrashStatement(ctx *ShowTrashStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowMigrationsStatement(ctx *ShowMigrationsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowJobTaskStatement(ctx *ShowJobTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropUserStatement(ctx *DropUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowUserStatement(ctx *ShowUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowAllAuthentication(ctx *ShowAllAuthenticationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExecuteAsStatement(ctx *ExecuteAsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowRolesStatement(ctx *ShowRolesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantRoleToUser(ctx *GrantRoleToUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantRoleToRole(ctx *GrantRoleToRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetRoleStatement(ctx *SetRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantRevokeClause(ctx *GrantRevokeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantOnUser(ctx *GrantOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantOnTableBrief(ctx *GrantOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantOnFunc(ctx *GrantOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantOnSystem(ctx *GrantOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGrantOnAll(ctx *GrantOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeOnUser(ctx *RevokeOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeOnFunc(ctx *RevokeOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeOnSystem(ctx *RevokeOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRevokeOnAll(ctx *RevokeOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowGrantsStatement(ctx *ShowGrantsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAuthWithPlugin(ctx *AuthWithPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrivObjectName(ctx *PrivObjectNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrivObjectNameList(ctx *PrivObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrivilegeTypeList(ctx *PrivilegeTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrivObjectType(ctx *PrivObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBackupStatement(ctx *BackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelBackupStatement(ctx *CancelBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowBackupStatement(ctx *ShowBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRestoreStatement(ctx *RestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelRestoreStatement(ctx *CancelRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowRestoreStatement(ctx *ShowRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropRepositoryStatement(ctx *DropRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDataCacheTarget(ctx *DataCacheTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCancelExportStatement(ctx *CancelExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowExportStatement(ctx *ShowExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInstallPluginStatement(ctx *InstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUninstallPluginStatement(ctx *UninstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateFileStatement(ctx *CreateFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropFileStatement(ctx *DropFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreatePipeStatement(ctx *CreatePipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropPipeStatement(ctx *DropPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterPipeClause(ctx *AlterPipeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterPipeStatement(ctx *AlterPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDescPipeStatement(ctx *DescPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowPipeStatement(ctx *ShowPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetStatement(ctx *SetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetNames(ctx *SetNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetPassword(ctx *SetPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetUserVar(ctx *SetUserVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetSystemVar(ctx *SetSystemVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTransaction_characteristics(ctx *Transaction_characteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTransaction_access_mode(ctx *Transaction_access_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIsolation_level(ctx *Isolation_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIsolation_types(ctx *Isolation_typesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRoleList(ctx *RoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUnsupportedStatement(ctx *UnsupportedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLock_item(ctx *Lock_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLock_type(ctx *Lock_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropWarehouseStatement(ctx *DropWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetWarehouseStatement(ctx *SetWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowClustersStatement(ctx *ShowClustersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitShowNodesStatement(ctx *ShowNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDropCNGroupStatement(ctx *DropCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBeginStatement(ctx *BeginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCommitStatement(ctx *CommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTranslateStatement(ctx *TranslateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDialect(ctx *DialectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTranslateSQL(ctx *TranslateSQLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQueryStatement(ctx *QueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQueryRelation(ctx *QueryRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQueryNoWith(ctx *QueryNoWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQueryPeriod(ctx *QueryPeriodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPeriodType(ctx *PeriodTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQueryWithParentheses(ctx *QueryWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetOperation(ctx *SetOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRowConstructor(ctx *RowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLimitConstExpr(ctx *LimitConstExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLimitElement(ctx *LimitElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitFrom(ctx *FromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDual(ctx *DualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRollup(ctx *RollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCube(ctx *CubeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGroupingSet(ctx *GroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSelectSingle(ctx *SelectSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSelectAll(ctx *SelectAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExcludeClause(ctx *ExcludeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRelations(ctx *RelationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRelationLateralView(ctx *RelationLateralViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLateralView(ctx *LateralViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGeneratorFunction(ctx *GeneratorFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTableAtom(ctx *TableAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInlineTable(ctx *InlineTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSubqueryWithAlias(ctx *SubqueryWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTableFunction(ctx *TableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitFileTableFunction(ctx *FileTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPivotClause(ctx *PivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPivotValue(ctx *PivotValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSampleClause(ctx *SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitNamedArgumentList(ctx *NamedArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitNamedArguments(ctx *NamedArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBracketHint(ctx *BracketHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitHintMap(ctx *HintMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionNames(ctx *PartitionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitKeyPartitionList(ctx *KeyPartitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTabletList(ctx *TabletListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPrepareSql(ctx *PrepareSqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDeallocateStatement(ctx *DeallocateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitReplicaList(ctx *ReplicaListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMapExpressionList(ctx *MapExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMapExpression(ctx *MapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExpressionSingleton(ctx *ExpressionSingletonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExpressionDefault(ctx *ExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIsNull(ctx *IsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitScalarSubquery(ctx *ScalarSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTupleInSubquery(ctx *TupleInSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInSubquery(ctx *InSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBetween(ctx *BetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLike(ctx *LikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMatchExpr(ctx *MatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnRef(ctx *ColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitConvert(ctx *ConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCollectionSubscript(ctx *CollectionSubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUserVariableExpression(ctx *UserVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArrowExpression(ctx *ArrowExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSystemVariableExpression(ctx *SystemVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitConcat(ctx *ConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDictionaryGetExpr(ctx *DictionaryGetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCollate(ctx *CollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMapConstructor(ctx *MapConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArraySlice(ctx *ArraySliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDateLiteral(ctx *DateLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInformationFunction(ctx *InformationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSpecialDateTime(ctx *SpecialDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAggregationFunctionCall(ctx *AggregationFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTranslateFunctionCall(ctx *TranslateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUserVariable(ctx *UserVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSystemVariable(ctx *SystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitWindowFunction(ctx *WindowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIgnoreNulls(ctx *IgnoreNullsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBoundedFrame(ctx *BoundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTableDesc(ctx *TableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExplainDesc(ctx *ExplainDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOptimizerTrace(ctx *OptimizerTraceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionExpr(ctx *PartitionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionDesc(ctx *PartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitListPartitionValues(ctx *ListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitListPartitionValue(ctx *ListPartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitLiteralExpressionList(ctx *LiteralExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSingleRangePartition(ctx *SingleRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMultiRangePartition(ctx *MultiRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionRangeDesc(ctx *PartitionRangeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionKeyDesc(ctx *PartitionKeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionValueList(ctx *PartitionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitKeyPartition(ctx *KeyPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPartitionValue(ctx *PartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDistributionClause(ctx *DistributionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDistributionDesc(ctx *DistributionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitStatusDesc(ctx *StatusDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitExtProperties(ctx *ExtPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitPropertyList(ctx *PropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUserPropertyList(ctx *UserPropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInlineProperties(ctx *InlinePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInlineProperty(ctx *InlinePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitOutfile(ctx *OutfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitFileFormat(ctx *FileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTaskInterval(ctx *TaskIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUnitIdentifier(ctx *UnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUnitBoundary(ctx *UnitBoundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSubfieldDesc(ctx *SubfieldDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitSubfieldDescs(ctx *SubfieldDescsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBaseType(ctx *BaseTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDecimalType(ctx *DecimalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitWriteBranch(ctx *WriteBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIdentifierWithAlias(ctx *IdentifierWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIdentifierOrString(ctx *IdentifierOrStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIdentifierOrStringList(ctx *IdentifierOrStringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUserWithoutHost(ctx *UserWithoutHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUserWithHost(ctx *UserWithHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDecimalValue(ctx *DecimalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitDoubleValue(ctx *DoubleValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}
