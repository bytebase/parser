// Code generated from StarRocksSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package starrocks // StarRocksSQL
import "github.com/antlr4-go/antlr/v4"

type BaseStarRocksSQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseStarRocksSQLVisitor) VisitSqlStatements(ctx *SqlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUseDatabaseStatement(ctx *UseDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUseCatalogStatement(ctx *UseCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetCatalogStatement(ctx *SetCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateDbStatement(ctx *CreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropDbStatement(ctx *DropDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRecoverDbStmt(ctx *RecoverDbStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDataStmt(ctx *ShowDataStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnDesc(ctx *ColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDefaultDesc(ctx *DefaultDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIndexDesc(ctx *IndexDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitEngineDesc(ctx *EngineDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCharsetDesc(ctx *CharsetDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCollateDesc(ctx *CollateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitKeyDesc(ctx *KeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOrderByDesc(ctx *OrderByDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnNullable(ctx *ColumnNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTypeWithNullable(ctx *TypeWithNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAggStateDesc(ctx *AggStateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAggDesc(ctx *AggDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRollupDesc(ctx *RollupDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRollupItem(ctx *RollupItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDupKeys(ctx *DupKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitFromRollup(ctx *FromRollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOrReplace(ctx *OrReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropTableStatement(ctx *DropTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTableStatement(ctx *ShowTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowColumnStatement(ctx *ShowColumnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowAlterStatement(ctx *ShowAlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDescTableStatement(ctx *DescTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowIndexStatement(ctx *ShowIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRecoverTableStatement(ctx *RecoverTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropViewStatement(ctx *DropViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnNameWithComment(ctx *ColumnNameWithCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSubmitTaskStatement(ctx *SubmitTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTaskClause(ctx *TaskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropTaskStatement(ctx *DropTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTaskScheduleDesc(ctx *TaskScheduleDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMvPartitionExprs(ctx *MvPartitionExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMaterializedViewDesc(ctx *MaterializedViewDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitKillStatement(ctx *KillStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSyncStatement(ctx *SyncStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterSystemStatement(ctx *AlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterCatalogStatement(ctx *AlterCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTypeDesc(ctx *TypeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLocationsDesc(ctx *LocationsDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowFailPointStatement(ctx *ShowFailPointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropDictionaryStatement(ctx *DropDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDictionaryName(ctx *DictionaryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterClause(ctx *AlterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddFrontendClause(ctx *AddFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropFrontendClause(ctx *DropFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddBackendClause(ctx *AddBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropBackendClause(ctx *DropBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyBackendClause(ctx *ModifyBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyBrokerClause(ctx *ModifyBrokerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateImageClause(ctx *CreateImageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDisableDiskClause(ctx *DisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropIndexClause(ctx *DropIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTableRenameClause(ctx *TableRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSwapTableClause(ctx *SwapTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyCommentClause(ctx *ModifyCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOptimizeRange(ctx *OptimizeRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOptimizeClause(ctx *OptimizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddColumnClause(ctx *AddColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddColumnsClause(ctx *AddColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropColumnClause(ctx *DropColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyColumnClause(ctx *ModifyColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnRenameClause(ctx *ColumnRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitReorderColumnsClause(ctx *ReorderColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRollupRenameClause(ctx *RollupRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCompactionClause(ctx *CompactionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSubfieldName(ctx *SubfieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitNestedFieldName(ctx *NestedFieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddFieldClause(ctx *AddFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropFieldClause(ctx *DropFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropBranchClause(ctx *DropBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropTagClause(ctx *DropTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTableOperationClause(ctx *TableOperationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTagOptions(ctx *TagOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBranchOptions(ctx *BranchOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSnapshotRetention(ctx *SnapshotRetentionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRefRetain(ctx *RefRetainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSnapshotId(ctx *SnapshotIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTimeUnit(ctx *TimeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInteger_list(ctx *Integer_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddPartitionClause(ctx *AddPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropPartitionClause(ctx *DropPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitModifyPartitionClause(ctx *ModifyPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitReplacePartitionClause(ctx *ReplacePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionRenameClause(ctx *PartitionRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDataSource(ctx *DataSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLoadProperties(ctx *LoadPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColSeparatorProperty(ctx *ColSeparatorPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitImportColumns(ctx *ImportColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnProperties(ctx *ColumnPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitJobProperties(ctx *JobPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDataSourceProperties(ctx *DataSourcePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRegularColumns(ctx *RegularColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAllColumns(ctx *AllColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPredicateColumns(ctx *PredicateColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMultiColumnSet(ctx *MultiColumnSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropStatsStatement(ctx *DropStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitHistogramStatement(ctx *HistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropHistogramStatement(ctx *DropHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateResourceStatement(ctx *CreateResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterResourceStatement(ctx *AlterResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropResourceStatement(ctx *DropResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowResourceStatement(ctx *ShowResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitClassifier(ctx *ClassifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInlineFunction(ctx *InlineFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLabelName(ctx *LabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDataDescList(ctx *DataDescListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDataDesc(ctx *DataDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitFormatProps(ctx *FormatPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBrokerDesc(ctx *BrokerDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitResourceDesc(ctx *ResourceDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowLoadStatement(ctx *ShowLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelLoadStatement(ctx *CancelLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterLoadStatement(ctx *AlterLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelCompactionStatement(ctx *CancelCompactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowAuthorStatement(ctx *ShowAuthorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowBackendsStatement(ctx *ShowBackendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowBrokerStatement(ctx *ShowBrokerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCharsetStatement(ctx *ShowCharsetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCollationStatement(ctx *ShowCollationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDeleteStatement(ctx *ShowDeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowEventsStatement(ctx *ShowEventsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowEnginesStatement(ctx *ShowEnginesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowPluginsStatement(ctx *ShowPluginsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowProcedureStatement(ctx *ShowProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowProcStatement(ctx *ShowProcStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowStatusStatement(ctx *ShowStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTabletStatement(ctx *ShowTabletStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTransactionStatement(ctx *ShowTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTriggersStatement(ctx *ShowTriggersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowVariablesStatement(ctx *ShowVariablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowWarningStatement(ctx *ShowWarningStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitHelpStatement(ctx *HelpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDataSkewStatement(ctx *ShowDataSkewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDataTypesStatement(ctx *ShowDataTypesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowSyncJobStatement(ctx *ShowSyncJobStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowPolicyStatement(ctx *ShowPolicyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowLastInsertStatement(ctx *ShowLastInsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTableIdStatement(ctx *ShowTableIdStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTableStatsStatement(ctx *ShowTableStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowTrashStatement(ctx *ShowTrashStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowMigrationsStatement(ctx *ShowMigrationsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowJobTaskStatement(ctx *ShowJobTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropUserStatement(ctx *DropUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowUserStatement(ctx *ShowUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowAllAuthentication(ctx *ShowAllAuthenticationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExecuteAsStatement(ctx *ExecuteAsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowRolesStatement(ctx *ShowRolesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantRoleToUser(ctx *GrantRoleToUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantRoleToRole(ctx *GrantRoleToRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetRoleStatement(ctx *SetRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantRevokeClause(ctx *GrantRevokeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantOnUser(ctx *GrantOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantOnTableBrief(ctx *GrantOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantOnFunc(ctx *GrantOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantOnSystem(ctx *GrantOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGrantOnAll(ctx *GrantOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeOnUser(ctx *RevokeOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeOnFunc(ctx *RevokeOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeOnSystem(ctx *RevokeOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRevokeOnAll(ctx *RevokeOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowGrantsStatement(ctx *ShowGrantsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAuthWithPlugin(ctx *AuthWithPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrivObjectName(ctx *PrivObjectNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrivObjectNameList(ctx *PrivObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrivilegeTypeList(ctx *PrivilegeTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrivObjectType(ctx *PrivObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBackupStatement(ctx *BackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelBackupStatement(ctx *CancelBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowBackupStatement(ctx *ShowBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRestoreStatement(ctx *RestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelRestoreStatement(ctx *CancelRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowRestoreStatement(ctx *ShowRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropRepositoryStatement(ctx *DropRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDataCacheTarget(ctx *DataCacheTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCancelExportStatement(ctx *CancelExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowExportStatement(ctx *ShowExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInstallPluginStatement(ctx *InstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUninstallPluginStatement(ctx *UninstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateFileStatement(ctx *CreateFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropFileStatement(ctx *DropFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreatePipeStatement(ctx *CreatePipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropPipeStatement(ctx *DropPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterPipeClause(ctx *AlterPipeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterPipeStatement(ctx *AlterPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDescPipeStatement(ctx *DescPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowPipeStatement(ctx *ShowPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetStatement(ctx *SetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetNames(ctx *SetNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetPassword(ctx *SetPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetUserVar(ctx *SetUserVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetSystemVar(ctx *SetSystemVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTransaction_characteristics(ctx *Transaction_characteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTransaction_access_mode(ctx *Transaction_access_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIsolation_level(ctx *Isolation_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIsolation_types(ctx *Isolation_typesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRoleList(ctx *RoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUnsupportedStatement(ctx *UnsupportedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLock_item(ctx *Lock_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLock_type(ctx *Lock_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropWarehouseStatement(ctx *DropWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetWarehouseStatement(ctx *SetWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowClustersStatement(ctx *ShowClustersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitShowNodesStatement(ctx *ShowNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDropCNGroupStatement(ctx *DropCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBeginStatement(ctx *BeginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCommitStatement(ctx *CommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTranslateStatement(ctx *TranslateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDialect(ctx *DialectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTranslateSQL(ctx *TranslateSQLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQueryStatement(ctx *QueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQueryRelation(ctx *QueryRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQueryNoWith(ctx *QueryNoWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQueryPeriod(ctx *QueryPeriodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPeriodType(ctx *PeriodTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQueryWithParentheses(ctx *QueryWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetOperation(ctx *SetOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRowConstructor(ctx *RowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLimitConstExpr(ctx *LimitConstExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLimitElement(ctx *LimitElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitFrom(ctx *FromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDual(ctx *DualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRollup(ctx *RollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCube(ctx *CubeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGroupingSet(ctx *GroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSelectSingle(ctx *SelectSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSelectAll(ctx *SelectAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExcludeClause(ctx *ExcludeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRelations(ctx *RelationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRelationLateralView(ctx *RelationLateralViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLateralView(ctx *LateralViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGeneratorFunction(ctx *GeneratorFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTableAtom(ctx *TableAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInlineTable(ctx *InlineTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSubqueryWithAlias(ctx *SubqueryWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTableFunction(ctx *TableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitFileTableFunction(ctx *FileTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPivotClause(ctx *PivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPivotValue(ctx *PivotValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSampleClause(ctx *SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitNamedArgumentList(ctx *NamedArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitNamedArguments(ctx *NamedArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBracketHint(ctx *BracketHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitHintMap(ctx *HintMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionNames(ctx *PartitionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitKeyPartitionList(ctx *KeyPartitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTabletList(ctx *TabletListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPrepareSql(ctx *PrepareSqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDeallocateStatement(ctx *DeallocateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitReplicaList(ctx *ReplicaListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMapExpressionList(ctx *MapExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMapExpression(ctx *MapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExpressionSingleton(ctx *ExpressionSingletonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExpressionDefault(ctx *ExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIsNull(ctx *IsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitScalarSubquery(ctx *ScalarSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTupleInSubquery(ctx *TupleInSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInSubquery(ctx *InSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBetween(ctx *BetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLike(ctx *LikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMatchExpr(ctx *MatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnRef(ctx *ColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitConvert(ctx *ConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCollectionSubscript(ctx *CollectionSubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUserVariableExpression(ctx *UserVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArrowExpression(ctx *ArrowExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSystemVariableExpression(ctx *SystemVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitConcat(ctx *ConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDictionaryGetExpr(ctx *DictionaryGetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCollate(ctx *CollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMapConstructor(ctx *MapConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArraySlice(ctx *ArraySliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDateLiteral(ctx *DateLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInformationFunction(ctx *InformationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSpecialDateTime(ctx *SpecialDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAggregationFunctionCall(ctx *AggregationFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTranslateFunctionCall(ctx *TranslateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUserVariable(ctx *UserVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSystemVariable(ctx *SystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitWindowFunction(ctx *WindowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIgnoreNulls(ctx *IgnoreNullsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBoundedFrame(ctx *BoundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTableDesc(ctx *TableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExplainDesc(ctx *ExplainDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOptimizerTrace(ctx *OptimizerTraceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionExpr(ctx *PartitionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionDesc(ctx *PartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitListPartitionValues(ctx *ListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitListPartitionValue(ctx *ListPartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitLiteralExpressionList(ctx *LiteralExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSingleRangePartition(ctx *SingleRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMultiRangePartition(ctx *MultiRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionRangeDesc(ctx *PartitionRangeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionKeyDesc(ctx *PartitionKeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionValueList(ctx *PartitionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitKeyPartition(ctx *KeyPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPartitionValue(ctx *PartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDistributionClause(ctx *DistributionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDistributionDesc(ctx *DistributionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitStatusDesc(ctx *StatusDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitExtProperties(ctx *ExtPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitPropertyList(ctx *PropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUserPropertyList(ctx *UserPropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInlineProperties(ctx *InlinePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInlineProperty(ctx *InlinePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitOutfile(ctx *OutfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitFileFormat(ctx *FileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTaskInterval(ctx *TaskIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUnitIdentifier(ctx *UnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUnitBoundary(ctx *UnitBoundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSubfieldDesc(ctx *SubfieldDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitSubfieldDescs(ctx *SubfieldDescsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBaseType(ctx *BaseTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDecimalType(ctx *DecimalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitWriteBranch(ctx *WriteBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIdentifierWithAlias(ctx *IdentifierWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIdentifierOrString(ctx *IdentifierOrStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIdentifierOrStringList(ctx *IdentifierOrStringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUserWithoutHost(ctx *UserWithoutHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUserWithHost(ctx *UserWithHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDecimalValue(ctx *DecimalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitDoubleValue(ctx *DoubleValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStarRocksSQLVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}
