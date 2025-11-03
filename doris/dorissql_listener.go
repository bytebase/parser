// Code generated from DorisSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package doris // DorisSQL
import "github.com/antlr4-go/antlr/v4"

// DorisSQLListener is a complete listener for a parse tree produced by DorisSQLParser.
type DorisSQLListener interface {
	antlr.ParseTreeListener

	// EnterSqlStatements is called when entering the sqlStatements production.
	EnterSqlStatements(c *SqlStatementsContext)

	// EnterSingleStatement is called when entering the singleStatement production.
	EnterSingleStatement(c *SingleStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterUseDatabaseStatement is called when entering the useDatabaseStatement production.
	EnterUseDatabaseStatement(c *UseDatabaseStatementContext)

	// EnterUseCatalogStatement is called when entering the useCatalogStatement production.
	EnterUseCatalogStatement(c *UseCatalogStatementContext)

	// EnterSetCatalogStatement is called when entering the setCatalogStatement production.
	EnterSetCatalogStatement(c *SetCatalogStatementContext)

	// EnterShowDatabasesStatement is called when entering the showDatabasesStatement production.
	EnterShowDatabasesStatement(c *ShowDatabasesStatementContext)

	// EnterAlterDbQuotaStatement is called when entering the alterDbQuotaStatement production.
	EnterAlterDbQuotaStatement(c *AlterDbQuotaStatementContext)

	// EnterCreateDbStatement is called when entering the createDbStatement production.
	EnterCreateDbStatement(c *CreateDbStatementContext)

	// EnterDropDbStatement is called when entering the dropDbStatement production.
	EnterDropDbStatement(c *DropDbStatementContext)

	// EnterShowCreateDbStatement is called when entering the showCreateDbStatement production.
	EnterShowCreateDbStatement(c *ShowCreateDbStatementContext)

	// EnterAlterDatabaseRenameStatement is called when entering the alterDatabaseRenameStatement production.
	EnterAlterDatabaseRenameStatement(c *AlterDatabaseRenameStatementContext)

	// EnterRecoverDbStmt is called when entering the recoverDbStmt production.
	EnterRecoverDbStmt(c *RecoverDbStmtContext)

	// EnterShowDataStmt is called when entering the showDataStmt production.
	EnterShowDataStmt(c *ShowDataStmtContext)

	// EnterShowDataDistributionStmt is called when entering the showDataDistributionStmt production.
	EnterShowDataDistributionStmt(c *ShowDataDistributionStmtContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterColumnDesc is called when entering the columnDesc production.
	EnterColumnDesc(c *ColumnDescContext)

	// EnterCharsetName is called when entering the charsetName production.
	EnterCharsetName(c *CharsetNameContext)

	// EnterDefaultDesc is called when entering the defaultDesc production.
	EnterDefaultDesc(c *DefaultDescContext)

	// EnterGeneratedColumnDesc is called when entering the generatedColumnDesc production.
	EnterGeneratedColumnDesc(c *GeneratedColumnDescContext)

	// EnterIndexDesc is called when entering the indexDesc production.
	EnterIndexDesc(c *IndexDescContext)

	// EnterEngineDesc is called when entering the engineDesc production.
	EnterEngineDesc(c *EngineDescContext)

	// EnterCharsetDesc is called when entering the charsetDesc production.
	EnterCharsetDesc(c *CharsetDescContext)

	// EnterCollateDesc is called when entering the collateDesc production.
	EnterCollateDesc(c *CollateDescContext)

	// EnterKeyDesc is called when entering the keyDesc production.
	EnterKeyDesc(c *KeyDescContext)

	// EnterOrderByDesc is called when entering the orderByDesc production.
	EnterOrderByDesc(c *OrderByDescContext)

	// EnterColumnNullable is called when entering the columnNullable production.
	EnterColumnNullable(c *ColumnNullableContext)

	// EnterTypeWithNullable is called when entering the typeWithNullable production.
	EnterTypeWithNullable(c *TypeWithNullableContext)

	// EnterAggStateDesc is called when entering the aggStateDesc production.
	EnterAggStateDesc(c *AggStateDescContext)

	// EnterAggDesc is called when entering the aggDesc production.
	EnterAggDesc(c *AggDescContext)

	// EnterRollupDesc is called when entering the rollupDesc production.
	EnterRollupDesc(c *RollupDescContext)

	// EnterRollupItem is called when entering the rollupItem production.
	EnterRollupItem(c *RollupItemContext)

	// EnterDupKeys is called when entering the dupKeys production.
	EnterDupKeys(c *DupKeysContext)

	// EnterFromRollup is called when entering the fromRollup production.
	EnterFromRollup(c *FromRollupContext)

	// EnterOrReplace is called when entering the orReplace production.
	EnterOrReplace(c *OrReplaceContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// EnterCreateTableAsSelectStatement is called when entering the createTableAsSelectStatement production.
	EnterCreateTableAsSelectStatement(c *CreateTableAsSelectStatementContext)

	// EnterDropTableStatement is called when entering the dropTableStatement production.
	EnterDropTableStatement(c *DropTableStatementContext)

	// EnterCleanTemporaryTableStatement is called when entering the cleanTemporaryTableStatement production.
	EnterCleanTemporaryTableStatement(c *CleanTemporaryTableStatementContext)

	// EnterAlterTableStatement is called when entering the alterTableStatement production.
	EnterAlterTableStatement(c *AlterTableStatementContext)

	// EnterCreateIndexStatement is called when entering the createIndexStatement production.
	EnterCreateIndexStatement(c *CreateIndexStatementContext)

	// EnterDropIndexStatement is called when entering the dropIndexStatement production.
	EnterDropIndexStatement(c *DropIndexStatementContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterShowTableStatement is called when entering the showTableStatement production.
	EnterShowTableStatement(c *ShowTableStatementContext)

	// EnterShowTemporaryTablesStatement is called when entering the showTemporaryTablesStatement production.
	EnterShowTemporaryTablesStatement(c *ShowTemporaryTablesStatementContext)

	// EnterShowCreateTableStatement is called when entering the showCreateTableStatement production.
	EnterShowCreateTableStatement(c *ShowCreateTableStatementContext)

	// EnterShowColumnStatement is called when entering the showColumnStatement production.
	EnterShowColumnStatement(c *ShowColumnStatementContext)

	// EnterShowTableStatusStatement is called when entering the showTableStatusStatement production.
	EnterShowTableStatusStatement(c *ShowTableStatusStatementContext)

	// EnterRefreshTableStatement is called when entering the refreshTableStatement production.
	EnterRefreshTableStatement(c *RefreshTableStatementContext)

	// EnterShowAlterStatement is called when entering the showAlterStatement production.
	EnterShowAlterStatement(c *ShowAlterStatementContext)

	// EnterDescTableStatement is called when entering the descTableStatement production.
	EnterDescTableStatement(c *DescTableStatementContext)

	// EnterCreateTableLikeStatement is called when entering the createTableLikeStatement production.
	EnterCreateTableLikeStatement(c *CreateTableLikeStatementContext)

	// EnterShowIndexStatement is called when entering the showIndexStatement production.
	EnterShowIndexStatement(c *ShowIndexStatementContext)

	// EnterRecoverTableStatement is called when entering the recoverTableStatement production.
	EnterRecoverTableStatement(c *RecoverTableStatementContext)

	// EnterTruncateTableStatement is called when entering the truncateTableStatement production.
	EnterTruncateTableStatement(c *TruncateTableStatementContext)

	// EnterCancelAlterTableStatement is called when entering the cancelAlterTableStatement production.
	EnterCancelAlterTableStatement(c *CancelAlterTableStatementContext)

	// EnterShowPartitionsStatement is called when entering the showPartitionsStatement production.
	EnterShowPartitionsStatement(c *ShowPartitionsStatementContext)

	// EnterRecoverPartitionStatement is called when entering the recoverPartitionStatement production.
	EnterRecoverPartitionStatement(c *RecoverPartitionStatementContext)

	// EnterCreateViewStatement is called when entering the createViewStatement production.
	EnterCreateViewStatement(c *CreateViewStatementContext)

	// EnterAlterViewStatement is called when entering the alterViewStatement production.
	EnterAlterViewStatement(c *AlterViewStatementContext)

	// EnterDropViewStatement is called when entering the dropViewStatement production.
	EnterDropViewStatement(c *DropViewStatementContext)

	// EnterColumnNameWithComment is called when entering the columnNameWithComment production.
	EnterColumnNameWithComment(c *ColumnNameWithCommentContext)

	// EnterSubmitTaskStatement is called when entering the submitTaskStatement production.
	EnterSubmitTaskStatement(c *SubmitTaskStatementContext)

	// EnterTaskClause is called when entering the taskClause production.
	EnterTaskClause(c *TaskClauseContext)

	// EnterDropTaskStatement is called when entering the dropTaskStatement production.
	EnterDropTaskStatement(c *DropTaskStatementContext)

	// EnterTaskScheduleDesc is called when entering the taskScheduleDesc production.
	EnterTaskScheduleDesc(c *TaskScheduleDescContext)

	// EnterCreateMaterializedViewStatement is called when entering the createMaterializedViewStatement production.
	EnterCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// EnterMvPartitionExprs is called when entering the mvPartitionExprs production.
	EnterMvPartitionExprs(c *MvPartitionExprsContext)

	// EnterMaterializedViewDesc is called when entering the materializedViewDesc production.
	EnterMaterializedViewDesc(c *MaterializedViewDescContext)

	// EnterShowMaterializedViewsStatement is called when entering the showMaterializedViewsStatement production.
	EnterShowMaterializedViewsStatement(c *ShowMaterializedViewsStatementContext)

	// EnterDropMaterializedViewStatement is called when entering the dropMaterializedViewStatement production.
	EnterDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// EnterAlterMaterializedViewStatement is called when entering the alterMaterializedViewStatement production.
	EnterAlterMaterializedViewStatement(c *AlterMaterializedViewStatementContext)

	// EnterRefreshMaterializedViewStatement is called when entering the refreshMaterializedViewStatement production.
	EnterRefreshMaterializedViewStatement(c *RefreshMaterializedViewStatementContext)

	// EnterCancelRefreshMaterializedViewStatement is called when entering the cancelRefreshMaterializedViewStatement production.
	EnterCancelRefreshMaterializedViewStatement(c *CancelRefreshMaterializedViewStatementContext)

	// EnterAdminSetConfigStatement is called when entering the adminSetConfigStatement production.
	EnterAdminSetConfigStatement(c *AdminSetConfigStatementContext)

	// EnterAdminSetReplicaStatusStatement is called when entering the adminSetReplicaStatusStatement production.
	EnterAdminSetReplicaStatusStatement(c *AdminSetReplicaStatusStatementContext)

	// EnterAdminShowConfigStatement is called when entering the adminShowConfigStatement production.
	EnterAdminShowConfigStatement(c *AdminShowConfigStatementContext)

	// EnterAdminShowReplicaDistributionStatement is called when entering the adminShowReplicaDistributionStatement production.
	EnterAdminShowReplicaDistributionStatement(c *AdminShowReplicaDistributionStatementContext)

	// EnterAdminShowReplicaStatusStatement is called when entering the adminShowReplicaStatusStatement production.
	EnterAdminShowReplicaStatusStatement(c *AdminShowReplicaStatusStatementContext)

	// EnterAdminRepairTableStatement is called when entering the adminRepairTableStatement production.
	EnterAdminRepairTableStatement(c *AdminRepairTableStatementContext)

	// EnterAdminCancelRepairTableStatement is called when entering the adminCancelRepairTableStatement production.
	EnterAdminCancelRepairTableStatement(c *AdminCancelRepairTableStatementContext)

	// EnterAdminCheckTabletsStatement is called when entering the adminCheckTabletsStatement production.
	EnterAdminCheckTabletsStatement(c *AdminCheckTabletsStatementContext)

	// EnterAdminSetPartitionVersion is called when entering the adminSetPartitionVersion production.
	EnterAdminSetPartitionVersion(c *AdminSetPartitionVersionContext)

	// EnterKillStatement is called when entering the killStatement production.
	EnterKillStatement(c *KillStatementContext)

	// EnterSyncStatement is called when entering the syncStatement production.
	EnterSyncStatement(c *SyncStatementContext)

	// EnterAdminSetAutomatedSnapshotOnStatement is called when entering the adminSetAutomatedSnapshotOnStatement production.
	EnterAdminSetAutomatedSnapshotOnStatement(c *AdminSetAutomatedSnapshotOnStatementContext)

	// EnterAdminSetAutomatedSnapshotOffStatement is called when entering the adminSetAutomatedSnapshotOffStatement production.
	EnterAdminSetAutomatedSnapshotOffStatement(c *AdminSetAutomatedSnapshotOffStatementContext)

	// EnterAlterSystemStatement is called when entering the alterSystemStatement production.
	EnterAlterSystemStatement(c *AlterSystemStatementContext)

	// EnterCancelAlterSystemStatement is called when entering the cancelAlterSystemStatement production.
	EnterCancelAlterSystemStatement(c *CancelAlterSystemStatementContext)

	// EnterShowComputeNodesStatement is called when entering the showComputeNodesStatement production.
	EnterShowComputeNodesStatement(c *ShowComputeNodesStatementContext)

	// EnterCreateExternalCatalogStatement is called when entering the createExternalCatalogStatement production.
	EnterCreateExternalCatalogStatement(c *CreateExternalCatalogStatementContext)

	// EnterShowCreateExternalCatalogStatement is called when entering the showCreateExternalCatalogStatement production.
	EnterShowCreateExternalCatalogStatement(c *ShowCreateExternalCatalogStatementContext)

	// EnterDropExternalCatalogStatement is called when entering the dropExternalCatalogStatement production.
	EnterDropExternalCatalogStatement(c *DropExternalCatalogStatementContext)

	// EnterShowCatalogsStatement is called when entering the showCatalogsStatement production.
	EnterShowCatalogsStatement(c *ShowCatalogsStatementContext)

	// EnterAlterCatalogStatement is called when entering the alterCatalogStatement production.
	EnterAlterCatalogStatement(c *AlterCatalogStatementContext)

	// EnterCreateStorageVolumeStatement is called when entering the createStorageVolumeStatement production.
	EnterCreateStorageVolumeStatement(c *CreateStorageVolumeStatementContext)

	// EnterTypeDesc is called when entering the typeDesc production.
	EnterTypeDesc(c *TypeDescContext)

	// EnterLocationsDesc is called when entering the locationsDesc production.
	EnterLocationsDesc(c *LocationsDescContext)

	// EnterShowStorageVolumesStatement is called when entering the showStorageVolumesStatement production.
	EnterShowStorageVolumesStatement(c *ShowStorageVolumesStatementContext)

	// EnterDropStorageVolumeStatement is called when entering the dropStorageVolumeStatement production.
	EnterDropStorageVolumeStatement(c *DropStorageVolumeStatementContext)

	// EnterAlterStorageVolumeStatement is called when entering the alterStorageVolumeStatement production.
	EnterAlterStorageVolumeStatement(c *AlterStorageVolumeStatementContext)

	// EnterAlterStorageVolumeClause is called when entering the alterStorageVolumeClause production.
	EnterAlterStorageVolumeClause(c *AlterStorageVolumeClauseContext)

	// EnterModifyStorageVolumePropertiesClause is called when entering the modifyStorageVolumePropertiesClause production.
	EnterModifyStorageVolumePropertiesClause(c *ModifyStorageVolumePropertiesClauseContext)

	// EnterModifyStorageVolumeCommentClause is called when entering the modifyStorageVolumeCommentClause production.
	EnterModifyStorageVolumeCommentClause(c *ModifyStorageVolumeCommentClauseContext)

	// EnterDescStorageVolumeStatement is called when entering the descStorageVolumeStatement production.
	EnterDescStorageVolumeStatement(c *DescStorageVolumeStatementContext)

	// EnterSetDefaultStorageVolumeStatement is called when entering the setDefaultStorageVolumeStatement production.
	EnterSetDefaultStorageVolumeStatement(c *SetDefaultStorageVolumeStatementContext)

	// EnterUpdateFailPointStatusStatement is called when entering the updateFailPointStatusStatement production.
	EnterUpdateFailPointStatusStatement(c *UpdateFailPointStatusStatementContext)

	// EnterShowFailPointStatement is called when entering the showFailPointStatement production.
	EnterShowFailPointStatement(c *ShowFailPointStatementContext)

	// EnterCreateDictionaryStatement is called when entering the createDictionaryStatement production.
	EnterCreateDictionaryStatement(c *CreateDictionaryStatementContext)

	// EnterDropDictionaryStatement is called when entering the dropDictionaryStatement production.
	EnterDropDictionaryStatement(c *DropDictionaryStatementContext)

	// EnterRefreshDictionaryStatement is called when entering the refreshDictionaryStatement production.
	EnterRefreshDictionaryStatement(c *RefreshDictionaryStatementContext)

	// EnterShowDictionaryStatement is called when entering the showDictionaryStatement production.
	EnterShowDictionaryStatement(c *ShowDictionaryStatementContext)

	// EnterCancelRefreshDictionaryStatement is called when entering the cancelRefreshDictionaryStatement production.
	EnterCancelRefreshDictionaryStatement(c *CancelRefreshDictionaryStatementContext)

	// EnterDictionaryColumnDesc is called when entering the dictionaryColumnDesc production.
	EnterDictionaryColumnDesc(c *DictionaryColumnDescContext)

	// EnterDictionaryName is called when entering the dictionaryName production.
	EnterDictionaryName(c *DictionaryNameContext)

	// EnterAlterClause is called when entering the alterClause production.
	EnterAlterClause(c *AlterClauseContext)

	// EnterAddFrontendClause is called when entering the addFrontendClause production.
	EnterAddFrontendClause(c *AddFrontendClauseContext)

	// EnterDropFrontendClause is called when entering the dropFrontendClause production.
	EnterDropFrontendClause(c *DropFrontendClauseContext)

	// EnterModifyFrontendHostClause is called when entering the modifyFrontendHostClause production.
	EnterModifyFrontendHostClause(c *ModifyFrontendHostClauseContext)

	// EnterAddBackendClause is called when entering the addBackendClause production.
	EnterAddBackendClause(c *AddBackendClauseContext)

	// EnterDropBackendClause is called when entering the dropBackendClause production.
	EnterDropBackendClause(c *DropBackendClauseContext)

	// EnterDecommissionBackendClause is called when entering the decommissionBackendClause production.
	EnterDecommissionBackendClause(c *DecommissionBackendClauseContext)

	// EnterModifyBackendClause is called when entering the modifyBackendClause production.
	EnterModifyBackendClause(c *ModifyBackendClauseContext)

	// EnterAddComputeNodeClause is called when entering the addComputeNodeClause production.
	EnterAddComputeNodeClause(c *AddComputeNodeClauseContext)

	// EnterDropComputeNodeClause is called when entering the dropComputeNodeClause production.
	EnterDropComputeNodeClause(c *DropComputeNodeClauseContext)

	// EnterModifyBrokerClause is called when entering the modifyBrokerClause production.
	EnterModifyBrokerClause(c *ModifyBrokerClauseContext)

	// EnterAlterLoadErrorUrlClause is called when entering the alterLoadErrorUrlClause production.
	EnterAlterLoadErrorUrlClause(c *AlterLoadErrorUrlClauseContext)

	// EnterCreateImageClause is called when entering the createImageClause production.
	EnterCreateImageClause(c *CreateImageClauseContext)

	// EnterCleanTabletSchedQClause is called when entering the cleanTabletSchedQClause production.
	EnterCleanTabletSchedQClause(c *CleanTabletSchedQClauseContext)

	// EnterDecommissionDiskClause is called when entering the decommissionDiskClause production.
	EnterDecommissionDiskClause(c *DecommissionDiskClauseContext)

	// EnterCancelDecommissionDiskClause is called when entering the cancelDecommissionDiskClause production.
	EnterCancelDecommissionDiskClause(c *CancelDecommissionDiskClauseContext)

	// EnterDisableDiskClause is called when entering the disableDiskClause production.
	EnterDisableDiskClause(c *DisableDiskClauseContext)

	// EnterCancelDisableDiskClause is called when entering the cancelDisableDiskClause production.
	EnterCancelDisableDiskClause(c *CancelDisableDiskClauseContext)

	// EnterCreateIndexClause is called when entering the createIndexClause production.
	EnterCreateIndexClause(c *CreateIndexClauseContext)

	// EnterDropIndexClause is called when entering the dropIndexClause production.
	EnterDropIndexClause(c *DropIndexClauseContext)

	// EnterTableRenameClause is called when entering the tableRenameClause production.
	EnterTableRenameClause(c *TableRenameClauseContext)

	// EnterSwapTableClause is called when entering the swapTableClause production.
	EnterSwapTableClause(c *SwapTableClauseContext)

	// EnterModifyPropertiesClause is called when entering the modifyPropertiesClause production.
	EnterModifyPropertiesClause(c *ModifyPropertiesClauseContext)

	// EnterModifyCommentClause is called when entering the modifyCommentClause production.
	EnterModifyCommentClause(c *ModifyCommentClauseContext)

	// EnterOptimizeRange is called when entering the optimizeRange production.
	EnterOptimizeRange(c *OptimizeRangeContext)

	// EnterOptimizeClause is called when entering the optimizeClause production.
	EnterOptimizeClause(c *OptimizeClauseContext)

	// EnterAddColumnClause is called when entering the addColumnClause production.
	EnterAddColumnClause(c *AddColumnClauseContext)

	// EnterAddColumnsClause is called when entering the addColumnsClause production.
	EnterAddColumnsClause(c *AddColumnsClauseContext)

	// EnterDropColumnClause is called when entering the dropColumnClause production.
	EnterDropColumnClause(c *DropColumnClauseContext)

	// EnterModifyColumnClause is called when entering the modifyColumnClause production.
	EnterModifyColumnClause(c *ModifyColumnClauseContext)

	// EnterModifyColumnCommentClause is called when entering the modifyColumnCommentClause production.
	EnterModifyColumnCommentClause(c *ModifyColumnCommentClauseContext)

	// EnterColumnRenameClause is called when entering the columnRenameClause production.
	EnterColumnRenameClause(c *ColumnRenameClauseContext)

	// EnterReorderColumnsClause is called when entering the reorderColumnsClause production.
	EnterReorderColumnsClause(c *ReorderColumnsClauseContext)

	// EnterRollupRenameClause is called when entering the rollupRenameClause production.
	EnterRollupRenameClause(c *RollupRenameClauseContext)

	// EnterCompactionClause is called when entering the compactionClause production.
	EnterCompactionClause(c *CompactionClauseContext)

	// EnterSubfieldName is called when entering the subfieldName production.
	EnterSubfieldName(c *SubfieldNameContext)

	// EnterNestedFieldName is called when entering the nestedFieldName production.
	EnterNestedFieldName(c *NestedFieldNameContext)

	// EnterAddFieldClause is called when entering the addFieldClause production.
	EnterAddFieldClause(c *AddFieldClauseContext)

	// EnterDropFieldClause is called when entering the dropFieldClause production.
	EnterDropFieldClause(c *DropFieldClauseContext)

	// EnterCreateOrReplaceTagClause is called when entering the createOrReplaceTagClause production.
	EnterCreateOrReplaceTagClause(c *CreateOrReplaceTagClauseContext)

	// EnterCreateOrReplaceBranchClause is called when entering the createOrReplaceBranchClause production.
	EnterCreateOrReplaceBranchClause(c *CreateOrReplaceBranchClauseContext)

	// EnterDropBranchClause is called when entering the dropBranchClause production.
	EnterDropBranchClause(c *DropBranchClauseContext)

	// EnterDropTagClause is called when entering the dropTagClause production.
	EnterDropTagClause(c *DropTagClauseContext)

	// EnterTableOperationClause is called when entering the tableOperationClause production.
	EnterTableOperationClause(c *TableOperationClauseContext)

	// EnterTagOptions is called when entering the tagOptions production.
	EnterTagOptions(c *TagOptionsContext)

	// EnterBranchOptions is called when entering the branchOptions production.
	EnterBranchOptions(c *BranchOptionsContext)

	// EnterSnapshotRetention is called when entering the snapshotRetention production.
	EnterSnapshotRetention(c *SnapshotRetentionContext)

	// EnterRefRetain is called when entering the refRetain production.
	EnterRefRetain(c *RefRetainContext)

	// EnterMaxSnapshotAge is called when entering the maxSnapshotAge production.
	EnterMaxSnapshotAge(c *MaxSnapshotAgeContext)

	// EnterMinSnapshotsToKeep is called when entering the minSnapshotsToKeep production.
	EnterMinSnapshotsToKeep(c *MinSnapshotsToKeepContext)

	// EnterSnapshotId is called when entering the snapshotId production.
	EnterSnapshotId(c *SnapshotIdContext)

	// EnterTimeUnit is called when entering the timeUnit production.
	EnterTimeUnit(c *TimeUnitContext)

	// EnterInteger_list is called when entering the integer_list production.
	EnterInteger_list(c *Integer_listContext)

	// EnterDropPersistentIndexClause is called when entering the dropPersistentIndexClause production.
	EnterDropPersistentIndexClause(c *DropPersistentIndexClauseContext)

	// EnterAddPartitionClause is called when entering the addPartitionClause production.
	EnterAddPartitionClause(c *AddPartitionClauseContext)

	// EnterDropPartitionClause is called when entering the dropPartitionClause production.
	EnterDropPartitionClause(c *DropPartitionClauseContext)

	// EnterTruncatePartitionClause is called when entering the truncatePartitionClause production.
	EnterTruncatePartitionClause(c *TruncatePartitionClauseContext)

	// EnterModifyPartitionClause is called when entering the modifyPartitionClause production.
	EnterModifyPartitionClause(c *ModifyPartitionClauseContext)

	// EnterReplacePartitionClause is called when entering the replacePartitionClause production.
	EnterReplacePartitionClause(c *ReplacePartitionClauseContext)

	// EnterPartitionRenameClause is called when entering the partitionRenameClause production.
	EnterPartitionRenameClause(c *PartitionRenameClauseContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterInsertLabelOrColumnAliases is called when entering the insertLabelOrColumnAliases production.
	EnterInsertLabelOrColumnAliases(c *InsertLabelOrColumnAliasesContext)

	// EnterColumnAliasesOrByName is called when entering the columnAliasesOrByName production.
	EnterColumnAliasesOrByName(c *ColumnAliasesOrByNameContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterCreateRoutineLoadStatement is called when entering the createRoutineLoadStatement production.
	EnterCreateRoutineLoadStatement(c *CreateRoutineLoadStatementContext)

	// EnterAlterRoutineLoadStatement is called when entering the alterRoutineLoadStatement production.
	EnterAlterRoutineLoadStatement(c *AlterRoutineLoadStatementContext)

	// EnterDataSource is called when entering the dataSource production.
	EnterDataSource(c *DataSourceContext)

	// EnterLoadProperties is called when entering the loadProperties production.
	EnterLoadProperties(c *LoadPropertiesContext)

	// EnterColSeparatorProperty is called when entering the colSeparatorProperty production.
	EnterColSeparatorProperty(c *ColSeparatorPropertyContext)

	// EnterRowDelimiterProperty is called when entering the rowDelimiterProperty production.
	EnterRowDelimiterProperty(c *RowDelimiterPropertyContext)

	// EnterImportColumns is called when entering the importColumns production.
	EnterImportColumns(c *ImportColumnsContext)

	// EnterColumnProperties is called when entering the columnProperties production.
	EnterColumnProperties(c *ColumnPropertiesContext)

	// EnterJobProperties is called when entering the jobProperties production.
	EnterJobProperties(c *JobPropertiesContext)

	// EnterDataSourceProperties is called when entering the dataSourceProperties production.
	EnterDataSourceProperties(c *DataSourcePropertiesContext)

	// EnterStopRoutineLoadStatement is called when entering the stopRoutineLoadStatement production.
	EnterStopRoutineLoadStatement(c *StopRoutineLoadStatementContext)

	// EnterResumeRoutineLoadStatement is called when entering the resumeRoutineLoadStatement production.
	EnterResumeRoutineLoadStatement(c *ResumeRoutineLoadStatementContext)

	// EnterPauseRoutineLoadStatement is called when entering the pauseRoutineLoadStatement production.
	EnterPauseRoutineLoadStatement(c *PauseRoutineLoadStatementContext)

	// EnterShowRoutineLoadStatement is called when entering the showRoutineLoadStatement production.
	EnterShowRoutineLoadStatement(c *ShowRoutineLoadStatementContext)

	// EnterShowRoutineLoadTaskStatement is called when entering the showRoutineLoadTaskStatement production.
	EnterShowRoutineLoadTaskStatement(c *ShowRoutineLoadTaskStatementContext)

	// EnterShowCreateRoutineLoadStatement is called when entering the showCreateRoutineLoadStatement production.
	EnterShowCreateRoutineLoadStatement(c *ShowCreateRoutineLoadStatementContext)

	// EnterShowStreamLoadStatement is called when entering the showStreamLoadStatement production.
	EnterShowStreamLoadStatement(c *ShowStreamLoadStatementContext)

	// EnterAnalyzeStatement is called when entering the analyzeStatement production.
	EnterAnalyzeStatement(c *AnalyzeStatementContext)

	// EnterRegularColumns is called when entering the regularColumns production.
	EnterRegularColumns(c *RegularColumnsContext)

	// EnterAllColumns is called when entering the allColumns production.
	EnterAllColumns(c *AllColumnsContext)

	// EnterPredicateColumns is called when entering the predicateColumns production.
	EnterPredicateColumns(c *PredicateColumnsContext)

	// EnterMultiColumnSet is called when entering the multiColumnSet production.
	EnterMultiColumnSet(c *MultiColumnSetContext)

	// EnterDropStatsStatement is called when entering the dropStatsStatement production.
	EnterDropStatsStatement(c *DropStatsStatementContext)

	// EnterHistogramStatement is called when entering the histogramStatement production.
	EnterHistogramStatement(c *HistogramStatementContext)

	// EnterAnalyzeHistogramStatement is called when entering the analyzeHistogramStatement production.
	EnterAnalyzeHistogramStatement(c *AnalyzeHistogramStatementContext)

	// EnterDropHistogramStatement is called when entering the dropHistogramStatement production.
	EnterDropHistogramStatement(c *DropHistogramStatementContext)

	// EnterCreateAnalyzeStatement is called when entering the createAnalyzeStatement production.
	EnterCreateAnalyzeStatement(c *CreateAnalyzeStatementContext)

	// EnterDropAnalyzeJobStatement is called when entering the dropAnalyzeJobStatement production.
	EnterDropAnalyzeJobStatement(c *DropAnalyzeJobStatementContext)

	// EnterShowAnalyzeStatement is called when entering the showAnalyzeStatement production.
	EnterShowAnalyzeStatement(c *ShowAnalyzeStatementContext)

	// EnterShowStatsMetaStatement is called when entering the showStatsMetaStatement production.
	EnterShowStatsMetaStatement(c *ShowStatsMetaStatementContext)

	// EnterShowHistogramMetaStatement is called when entering the showHistogramMetaStatement production.
	EnterShowHistogramMetaStatement(c *ShowHistogramMetaStatementContext)

	// EnterKillAnalyzeStatement is called when entering the killAnalyzeStatement production.
	EnterKillAnalyzeStatement(c *KillAnalyzeStatementContext)

	// EnterAnalyzeProfileStatement is called when entering the analyzeProfileStatement production.
	EnterAnalyzeProfileStatement(c *AnalyzeProfileStatementContext)

	// EnterCreateBaselinePlanStatement is called when entering the createBaselinePlanStatement production.
	EnterCreateBaselinePlanStatement(c *CreateBaselinePlanStatementContext)

	// EnterDropBaselinePlanStatement is called when entering the dropBaselinePlanStatement production.
	EnterDropBaselinePlanStatement(c *DropBaselinePlanStatementContext)

	// EnterShowBaselinePlanStatement is called when entering the showBaselinePlanStatement production.
	EnterShowBaselinePlanStatement(c *ShowBaselinePlanStatementContext)

	// EnterCreateResourceGroupStatement is called when entering the createResourceGroupStatement production.
	EnterCreateResourceGroupStatement(c *CreateResourceGroupStatementContext)

	// EnterDropResourceGroupStatement is called when entering the dropResourceGroupStatement production.
	EnterDropResourceGroupStatement(c *DropResourceGroupStatementContext)

	// EnterAlterResourceGroupStatement is called when entering the alterResourceGroupStatement production.
	EnterAlterResourceGroupStatement(c *AlterResourceGroupStatementContext)

	// EnterShowResourceGroupStatement is called when entering the showResourceGroupStatement production.
	EnterShowResourceGroupStatement(c *ShowResourceGroupStatementContext)

	// EnterShowResourceGroupUsageStatement is called when entering the showResourceGroupUsageStatement production.
	EnterShowResourceGroupUsageStatement(c *ShowResourceGroupUsageStatementContext)

	// EnterCreateResourceStatement is called when entering the createResourceStatement production.
	EnterCreateResourceStatement(c *CreateResourceStatementContext)

	// EnterAlterResourceStatement is called when entering the alterResourceStatement production.
	EnterAlterResourceStatement(c *AlterResourceStatementContext)

	// EnterDropResourceStatement is called when entering the dropResourceStatement production.
	EnterDropResourceStatement(c *DropResourceStatementContext)

	// EnterShowResourceStatement is called when entering the showResourceStatement production.
	EnterShowResourceStatement(c *ShowResourceStatementContext)

	// EnterClassifier is called when entering the classifier production.
	EnterClassifier(c *ClassifierContext)

	// EnterShowFunctionsStatement is called when entering the showFunctionsStatement production.
	EnterShowFunctionsStatement(c *ShowFunctionsStatementContext)

	// EnterDropFunctionStatement is called when entering the dropFunctionStatement production.
	EnterDropFunctionStatement(c *DropFunctionStatementContext)

	// EnterCreateFunctionStatement is called when entering the createFunctionStatement production.
	EnterCreateFunctionStatement(c *CreateFunctionStatementContext)

	// EnterInlineFunction is called when entering the inlineFunction production.
	EnterInlineFunction(c *InlineFunctionContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterLoadStatement is called when entering the loadStatement production.
	EnterLoadStatement(c *LoadStatementContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterDataDescList is called when entering the dataDescList production.
	EnterDataDescList(c *DataDescListContext)

	// EnterDataDesc is called when entering the dataDesc production.
	EnterDataDesc(c *DataDescContext)

	// EnterFormatProps is called when entering the formatProps production.
	EnterFormatProps(c *FormatPropsContext)

	// EnterBrokerDesc is called when entering the brokerDesc production.
	EnterBrokerDesc(c *BrokerDescContext)

	// EnterResourceDesc is called when entering the resourceDesc production.
	EnterResourceDesc(c *ResourceDescContext)

	// EnterShowLoadStatement is called when entering the showLoadStatement production.
	EnterShowLoadStatement(c *ShowLoadStatementContext)

	// EnterShowLoadWarningsStatement is called when entering the showLoadWarningsStatement production.
	EnterShowLoadWarningsStatement(c *ShowLoadWarningsStatementContext)

	// EnterCancelLoadStatement is called when entering the cancelLoadStatement production.
	EnterCancelLoadStatement(c *CancelLoadStatementContext)

	// EnterAlterLoadStatement is called when entering the alterLoadStatement production.
	EnterAlterLoadStatement(c *AlterLoadStatementContext)

	// EnterCancelCompactionStatement is called when entering the cancelCompactionStatement production.
	EnterCancelCompactionStatement(c *CancelCompactionStatementContext)

	// EnterShowAuthorStatement is called when entering the showAuthorStatement production.
	EnterShowAuthorStatement(c *ShowAuthorStatementContext)

	// EnterShowBackendsStatement is called when entering the showBackendsStatement production.
	EnterShowBackendsStatement(c *ShowBackendsStatementContext)

	// EnterShowBrokerStatement is called when entering the showBrokerStatement production.
	EnterShowBrokerStatement(c *ShowBrokerStatementContext)

	// EnterShowCharsetStatement is called when entering the showCharsetStatement production.
	EnterShowCharsetStatement(c *ShowCharsetStatementContext)

	// EnterShowCollationStatement is called when entering the showCollationStatement production.
	EnterShowCollationStatement(c *ShowCollationStatementContext)

	// EnterShowDeleteStatement is called when entering the showDeleteStatement production.
	EnterShowDeleteStatement(c *ShowDeleteStatementContext)

	// EnterShowDynamicPartitionStatement is called when entering the showDynamicPartitionStatement production.
	EnterShowDynamicPartitionStatement(c *ShowDynamicPartitionStatementContext)

	// EnterShowEventsStatement is called when entering the showEventsStatement production.
	EnterShowEventsStatement(c *ShowEventsStatementContext)

	// EnterShowEnginesStatement is called when entering the showEnginesStatement production.
	EnterShowEnginesStatement(c *ShowEnginesStatementContext)

	// EnterShowFrontendsStatement is called when entering the showFrontendsStatement production.
	EnterShowFrontendsStatement(c *ShowFrontendsStatementContext)

	// EnterShowFrontendsDisksStatement is called when entering the showFrontendsDisksStatement production.
	EnterShowFrontendsDisksStatement(c *ShowFrontendsDisksStatementContext)

	// EnterShowPluginsStatement is called when entering the showPluginsStatement production.
	EnterShowPluginsStatement(c *ShowPluginsStatementContext)

	// EnterShowRepositoriesStatement is called when entering the showRepositoriesStatement production.
	EnterShowRepositoriesStatement(c *ShowRepositoriesStatementContext)

	// EnterShowOpenTableStatement is called when entering the showOpenTableStatement production.
	EnterShowOpenTableStatement(c *ShowOpenTableStatementContext)

	// EnterShowPrivilegesStatement is called when entering the showPrivilegesStatement production.
	EnterShowPrivilegesStatement(c *ShowPrivilegesStatementContext)

	// EnterShowProcedureStatement is called when entering the showProcedureStatement production.
	EnterShowProcedureStatement(c *ShowProcedureStatementContext)

	// EnterShowProcStatement is called when entering the showProcStatement production.
	EnterShowProcStatement(c *ShowProcStatementContext)

	// EnterShowProcesslistStatement is called when entering the showProcesslistStatement production.
	EnterShowProcesslistStatement(c *ShowProcesslistStatementContext)

	// EnterShowProfilelistStatement is called when entering the showProfilelistStatement production.
	EnterShowProfilelistStatement(c *ShowProfilelistStatementContext)

	// EnterShowRunningQueriesStatement is called when entering the showRunningQueriesStatement production.
	EnterShowRunningQueriesStatement(c *ShowRunningQueriesStatementContext)

	// EnterShowStatusStatement is called when entering the showStatusStatement production.
	EnterShowStatusStatement(c *ShowStatusStatementContext)

	// EnterShowTabletStatement is called when entering the showTabletStatement production.
	EnterShowTabletStatement(c *ShowTabletStatementContext)

	// EnterShowTransactionStatement is called when entering the showTransactionStatement production.
	EnterShowTransactionStatement(c *ShowTransactionStatementContext)

	// EnterShowTriggersStatement is called when entering the showTriggersStatement production.
	EnterShowTriggersStatement(c *ShowTriggersStatementContext)

	// EnterShowUserPropertyStatement is called when entering the showUserPropertyStatement production.
	EnterShowUserPropertyStatement(c *ShowUserPropertyStatementContext)

	// EnterShowVariablesStatement is called when entering the showVariablesStatement production.
	EnterShowVariablesStatement(c *ShowVariablesStatementContext)

	// EnterShowWarningStatement is called when entering the showWarningStatement production.
	EnterShowWarningStatement(c *ShowWarningStatementContext)

	// EnterHelpStatement is called when entering the helpStatement production.
	EnterHelpStatement(c *HelpStatementContext)

	// EnterShowQueryProfileStatement is called when entering the showQueryProfileStatement production.
	EnterShowQueryProfileStatement(c *ShowQueryProfileStatementContext)

	// EnterShowQueryStatsStatement is called when entering the showQueryStatsStatement production.
	EnterShowQueryStatsStatement(c *ShowQueryStatsStatementContext)

	// EnterShowLoadProfileStatement is called when entering the showLoadProfileStatement production.
	EnterShowLoadProfileStatement(c *ShowLoadProfileStatementContext)

	// EnterShowDataSkewStatement is called when entering the showDataSkewStatement production.
	EnterShowDataSkewStatement(c *ShowDataSkewStatementContext)

	// EnterShowDataTypesStatement is called when entering the showDataTypesStatement production.
	EnterShowDataTypesStatement(c *ShowDataTypesStatementContext)

	// EnterShowSyncJobStatement is called when entering the showSyncJobStatement production.
	EnterShowSyncJobStatement(c *ShowSyncJobStatementContext)

	// EnterShowPolicyStatement is called when entering the showPolicyStatement production.
	EnterShowPolicyStatement(c *ShowPolicyStatementContext)

	// EnterShowSqlBlockRuleStatement is called when entering the showSqlBlockRuleStatement production.
	EnterShowSqlBlockRuleStatement(c *ShowSqlBlockRuleStatementContext)

	// EnterShowEncryptKeysStatement is called when entering the showEncryptKeysStatement production.
	EnterShowEncryptKeysStatement(c *ShowEncryptKeysStatementContext)

	// EnterShowCreateLoadStatement is called when entering the showCreateLoadStatement production.
	EnterShowCreateLoadStatement(c *ShowCreateLoadStatementContext)

	// EnterShowCreateRepositoryStatement is called when entering the showCreateRepositoryStatement production.
	EnterShowCreateRepositoryStatement(c *ShowCreateRepositoryStatementContext)

	// EnterShowLastInsertStatement is called when entering the showLastInsertStatement production.
	EnterShowLastInsertStatement(c *ShowLastInsertStatementContext)

	// EnterShowTableIdStatement is called when entering the showTableIdStatement production.
	EnterShowTableIdStatement(c *ShowTableIdStatementContext)

	// EnterShowDatabaseIdStatement is called when entering the showDatabaseIdStatement production.
	EnterShowDatabaseIdStatement(c *ShowDatabaseIdStatementContext)

	// EnterShowPartitionIdStatement is called when entering the showPartitionIdStatement production.
	EnterShowPartitionIdStatement(c *ShowPartitionIdStatementContext)

	// EnterShowTableStatsStatement is called when entering the showTableStatsStatement production.
	EnterShowTableStatsStatement(c *ShowTableStatsStatementContext)

	// EnterShowColumnStatsStatement is called when entering the showColumnStatsStatement production.
	EnterShowColumnStatsStatement(c *ShowColumnStatsStatementContext)

	// EnterShowConvertLightSchemaChangeStatement is called when entering the showConvertLightSchemaChangeStatement production.
	EnterShowConvertLightSchemaChangeStatement(c *ShowConvertLightSchemaChangeStatementContext)

	// EnterShowCatalogRecycleBinStatement is called when entering the showCatalogRecycleBinStatement production.
	EnterShowCatalogRecycleBinStatement(c *ShowCatalogRecycleBinStatementContext)

	// EnterShowTrashStatement is called when entering the showTrashStatement production.
	EnterShowTrashStatement(c *ShowTrashStatementContext)

	// EnterShowMigrationsStatement is called when entering the showMigrationsStatement production.
	EnterShowMigrationsStatement(c *ShowMigrationsStatementContext)

	// EnterShowWorkloadGroupsStatement is called when entering the showWorkloadGroupsStatement production.
	EnterShowWorkloadGroupsStatement(c *ShowWorkloadGroupsStatementContext)

	// EnterShowJobTaskStatement is called when entering the showJobTaskStatement production.
	EnterShowJobTaskStatement(c *ShowJobTaskStatementContext)

	// EnterCreateUserStatement is called when entering the createUserStatement production.
	EnterCreateUserStatement(c *CreateUserStatementContext)

	// EnterDropUserStatement is called when entering the dropUserStatement production.
	EnterDropUserStatement(c *DropUserStatementContext)

	// EnterAlterUserStatement is called when entering the alterUserStatement production.
	EnterAlterUserStatement(c *AlterUserStatementContext)

	// EnterShowUserStatement is called when entering the showUserStatement production.
	EnterShowUserStatement(c *ShowUserStatementContext)

	// EnterShowAllAuthentication is called when entering the showAllAuthentication production.
	EnterShowAllAuthentication(c *ShowAllAuthenticationContext)

	// EnterShowAuthenticationForUser is called when entering the showAuthenticationForUser production.
	EnterShowAuthenticationForUser(c *ShowAuthenticationForUserContext)

	// EnterExecuteAsStatement is called when entering the executeAsStatement production.
	EnterExecuteAsStatement(c *ExecuteAsStatementContext)

	// EnterCreateRoleStatement is called when entering the createRoleStatement production.
	EnterCreateRoleStatement(c *CreateRoleStatementContext)

	// EnterAlterRoleStatement is called when entering the alterRoleStatement production.
	EnterAlterRoleStatement(c *AlterRoleStatementContext)

	// EnterDropRoleStatement is called when entering the dropRoleStatement production.
	EnterDropRoleStatement(c *DropRoleStatementContext)

	// EnterShowRolesStatement is called when entering the showRolesStatement production.
	EnterShowRolesStatement(c *ShowRolesStatementContext)

	// EnterGrantRoleToUser is called when entering the grantRoleToUser production.
	EnterGrantRoleToUser(c *GrantRoleToUserContext)

	// EnterGrantRoleToRole is called when entering the grantRoleToRole production.
	EnterGrantRoleToRole(c *GrantRoleToRoleContext)

	// EnterRevokeRoleFromUser is called when entering the revokeRoleFromUser production.
	EnterRevokeRoleFromUser(c *RevokeRoleFromUserContext)

	// EnterRevokeRoleFromRole is called when entering the revokeRoleFromRole production.
	EnterRevokeRoleFromRole(c *RevokeRoleFromRoleContext)

	// EnterSetRoleStatement is called when entering the setRoleStatement production.
	EnterSetRoleStatement(c *SetRoleStatementContext)

	// EnterSetDefaultRoleStatement is called when entering the setDefaultRoleStatement production.
	EnterSetDefaultRoleStatement(c *SetDefaultRoleStatementContext)

	// EnterGrantRevokeClause is called when entering the grantRevokeClause production.
	EnterGrantRevokeClause(c *GrantRevokeClauseContext)

	// EnterGrantOnUser is called when entering the grantOnUser production.
	EnterGrantOnUser(c *GrantOnUserContext)

	// EnterGrantOnTableBrief is called when entering the grantOnTableBrief production.
	EnterGrantOnTableBrief(c *GrantOnTableBriefContext)

	// EnterGrantOnFunc is called when entering the grantOnFunc production.
	EnterGrantOnFunc(c *GrantOnFuncContext)

	// EnterGrantOnSystem is called when entering the grantOnSystem production.
	EnterGrantOnSystem(c *GrantOnSystemContext)

	// EnterGrantOnPrimaryObj is called when entering the grantOnPrimaryObj production.
	EnterGrantOnPrimaryObj(c *GrantOnPrimaryObjContext)

	// EnterGrantOnAll is called when entering the grantOnAll production.
	EnterGrantOnAll(c *GrantOnAllContext)

	// EnterRevokeOnUser is called when entering the revokeOnUser production.
	EnterRevokeOnUser(c *RevokeOnUserContext)

	// EnterRevokeOnTableBrief is called when entering the revokeOnTableBrief production.
	EnterRevokeOnTableBrief(c *RevokeOnTableBriefContext)

	// EnterRevokeOnFunc is called when entering the revokeOnFunc production.
	EnterRevokeOnFunc(c *RevokeOnFuncContext)

	// EnterRevokeOnSystem is called when entering the revokeOnSystem production.
	EnterRevokeOnSystem(c *RevokeOnSystemContext)

	// EnterRevokeOnPrimaryObj is called when entering the revokeOnPrimaryObj production.
	EnterRevokeOnPrimaryObj(c *RevokeOnPrimaryObjContext)

	// EnterRevokeOnAll is called when entering the revokeOnAll production.
	EnterRevokeOnAll(c *RevokeOnAllContext)

	// EnterShowGrantsStatement is called when entering the showGrantsStatement production.
	EnterShowGrantsStatement(c *ShowGrantsStatementContext)

	// EnterAuthWithoutPlugin is called when entering the authWithoutPlugin production.
	EnterAuthWithoutPlugin(c *AuthWithoutPluginContext)

	// EnterAuthWithPlugin is called when entering the authWithPlugin production.
	EnterAuthWithPlugin(c *AuthWithPluginContext)

	// EnterPrivObjectName is called when entering the privObjectName production.
	EnterPrivObjectName(c *PrivObjectNameContext)

	// EnterPrivObjectNameList is called when entering the privObjectNameList production.
	EnterPrivObjectNameList(c *PrivObjectNameListContext)

	// EnterPrivFunctionObjectNameList is called when entering the privFunctionObjectNameList production.
	EnterPrivFunctionObjectNameList(c *PrivFunctionObjectNameListContext)

	// EnterPrivilegeTypeList is called when entering the privilegeTypeList production.
	EnterPrivilegeTypeList(c *PrivilegeTypeListContext)

	// EnterPrivilegeType is called when entering the privilegeType production.
	EnterPrivilegeType(c *PrivilegeTypeContext)

	// EnterPrivObjectType is called when entering the privObjectType production.
	EnterPrivObjectType(c *PrivObjectTypeContext)

	// EnterPrivObjectTypePlural is called when entering the privObjectTypePlural production.
	EnterPrivObjectTypePlural(c *PrivObjectTypePluralContext)

	// EnterCreateSecurityIntegrationStatement is called when entering the createSecurityIntegrationStatement production.
	EnterCreateSecurityIntegrationStatement(c *CreateSecurityIntegrationStatementContext)

	// EnterAlterSecurityIntegrationStatement is called when entering the alterSecurityIntegrationStatement production.
	EnterAlterSecurityIntegrationStatement(c *AlterSecurityIntegrationStatementContext)

	// EnterDropSecurityIntegrationStatement is called when entering the dropSecurityIntegrationStatement production.
	EnterDropSecurityIntegrationStatement(c *DropSecurityIntegrationStatementContext)

	// EnterShowSecurityIntegrationStatement is called when entering the showSecurityIntegrationStatement production.
	EnterShowSecurityIntegrationStatement(c *ShowSecurityIntegrationStatementContext)

	// EnterShowCreateSecurityIntegrationStatement is called when entering the showCreateSecurityIntegrationStatement production.
	EnterShowCreateSecurityIntegrationStatement(c *ShowCreateSecurityIntegrationStatementContext)

	// EnterCreateGroupProviderStatement is called when entering the createGroupProviderStatement production.
	EnterCreateGroupProviderStatement(c *CreateGroupProviderStatementContext)

	// EnterDropGroupProviderStatement is called when entering the dropGroupProviderStatement production.
	EnterDropGroupProviderStatement(c *DropGroupProviderStatementContext)

	// EnterShowGroupProvidersStatement is called when entering the showGroupProvidersStatement production.
	EnterShowGroupProvidersStatement(c *ShowGroupProvidersStatementContext)

	// EnterShowCreateGroupProviderStatement is called when entering the showCreateGroupProviderStatement production.
	EnterShowCreateGroupProviderStatement(c *ShowCreateGroupProviderStatementContext)

	// EnterBackupStatement is called when entering the backupStatement production.
	EnterBackupStatement(c *BackupStatementContext)

	// EnterCancelBackupStatement is called when entering the cancelBackupStatement production.
	EnterCancelBackupStatement(c *CancelBackupStatementContext)

	// EnterShowBackupStatement is called when entering the showBackupStatement production.
	EnterShowBackupStatement(c *ShowBackupStatementContext)

	// EnterRestoreStatement is called when entering the restoreStatement production.
	EnterRestoreStatement(c *RestoreStatementContext)

	// EnterCancelRestoreStatement is called when entering the cancelRestoreStatement production.
	EnterCancelRestoreStatement(c *CancelRestoreStatementContext)

	// EnterShowRestoreStatement is called when entering the showRestoreStatement production.
	EnterShowRestoreStatement(c *ShowRestoreStatementContext)

	// EnterShowSnapshotStatement is called when entering the showSnapshotStatement production.
	EnterShowSnapshotStatement(c *ShowSnapshotStatementContext)

	// EnterCreateRepositoryStatement is called when entering the createRepositoryStatement production.
	EnterCreateRepositoryStatement(c *CreateRepositoryStatementContext)

	// EnterDropRepositoryStatement is called when entering the dropRepositoryStatement production.
	EnterDropRepositoryStatement(c *DropRepositoryStatementContext)

	// EnterAddSqlBlackListStatement is called when entering the addSqlBlackListStatement production.
	EnterAddSqlBlackListStatement(c *AddSqlBlackListStatementContext)

	// EnterDelSqlBlackListStatement is called when entering the delSqlBlackListStatement production.
	EnterDelSqlBlackListStatement(c *DelSqlBlackListStatementContext)

	// EnterShowSqlBlackListStatement is called when entering the showSqlBlackListStatement production.
	EnterShowSqlBlackListStatement(c *ShowSqlBlackListStatementContext)

	// EnterShowWhiteListStatement is called when entering the showWhiteListStatement production.
	EnterShowWhiteListStatement(c *ShowWhiteListStatementContext)

	// EnterAddBackendBlackListStatement is called when entering the addBackendBlackListStatement production.
	EnterAddBackendBlackListStatement(c *AddBackendBlackListStatementContext)

	// EnterDelBackendBlackListStatement is called when entering the delBackendBlackListStatement production.
	EnterDelBackendBlackListStatement(c *DelBackendBlackListStatementContext)

	// EnterShowBackendBlackListStatement is called when entering the showBackendBlackListStatement production.
	EnterShowBackendBlackListStatement(c *ShowBackendBlackListStatementContext)

	// EnterDataCacheTarget is called when entering the dataCacheTarget production.
	EnterDataCacheTarget(c *DataCacheTargetContext)

	// EnterCreateDataCacheRuleStatement is called when entering the createDataCacheRuleStatement production.
	EnterCreateDataCacheRuleStatement(c *CreateDataCacheRuleStatementContext)

	// EnterShowDataCacheRulesStatement is called when entering the showDataCacheRulesStatement production.
	EnterShowDataCacheRulesStatement(c *ShowDataCacheRulesStatementContext)

	// EnterDropDataCacheRuleStatement is called when entering the dropDataCacheRuleStatement production.
	EnterDropDataCacheRuleStatement(c *DropDataCacheRuleStatementContext)

	// EnterClearDataCacheRulesStatement is called when entering the clearDataCacheRulesStatement production.
	EnterClearDataCacheRulesStatement(c *ClearDataCacheRulesStatementContext)

	// EnterDataCacheSelectStatement is called when entering the dataCacheSelectStatement production.
	EnterDataCacheSelectStatement(c *DataCacheSelectStatementContext)

	// EnterExportStatement is called when entering the exportStatement production.
	EnterExportStatement(c *ExportStatementContext)

	// EnterCancelExportStatement is called when entering the cancelExportStatement production.
	EnterCancelExportStatement(c *CancelExportStatementContext)

	// EnterShowExportStatement is called when entering the showExportStatement production.
	EnterShowExportStatement(c *ShowExportStatementContext)

	// EnterInstallPluginStatement is called when entering the installPluginStatement production.
	EnterInstallPluginStatement(c *InstallPluginStatementContext)

	// EnterUninstallPluginStatement is called when entering the uninstallPluginStatement production.
	EnterUninstallPluginStatement(c *UninstallPluginStatementContext)

	// EnterCreateFileStatement is called when entering the createFileStatement production.
	EnterCreateFileStatement(c *CreateFileStatementContext)

	// EnterDropFileStatement is called when entering the dropFileStatement production.
	EnterDropFileStatement(c *DropFileStatementContext)

	// EnterShowSmallFilesStatement is called when entering the showSmallFilesStatement production.
	EnterShowSmallFilesStatement(c *ShowSmallFilesStatementContext)

	// EnterCreatePipeStatement is called when entering the createPipeStatement production.
	EnterCreatePipeStatement(c *CreatePipeStatementContext)

	// EnterDropPipeStatement is called when entering the dropPipeStatement production.
	EnterDropPipeStatement(c *DropPipeStatementContext)

	// EnterAlterPipeClause is called when entering the alterPipeClause production.
	EnterAlterPipeClause(c *AlterPipeClauseContext)

	// EnterAlterPipeStatement is called when entering the alterPipeStatement production.
	EnterAlterPipeStatement(c *AlterPipeStatementContext)

	// EnterDescPipeStatement is called when entering the descPipeStatement production.
	EnterDescPipeStatement(c *DescPipeStatementContext)

	// EnterShowPipeStatement is called when entering the showPipeStatement production.
	EnterShowPipeStatement(c *ShowPipeStatementContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterSetNames is called when entering the setNames production.
	EnterSetNames(c *SetNamesContext)

	// EnterSetPassword is called when entering the setPassword production.
	EnterSetPassword(c *SetPasswordContext)

	// EnterSetUserVar is called when entering the setUserVar production.
	EnterSetUserVar(c *SetUserVarContext)

	// EnterSetSystemVar is called when entering the setSystemVar production.
	EnterSetSystemVar(c *SetSystemVarContext)

	// EnterSetTransaction is called when entering the setTransaction production.
	EnterSetTransaction(c *SetTransactionContext)

	// EnterTransaction_characteristics is called when entering the transaction_characteristics production.
	EnterTransaction_characteristics(c *Transaction_characteristicsContext)

	// EnterTransaction_access_mode is called when entering the transaction_access_mode production.
	EnterTransaction_access_mode(c *Transaction_access_modeContext)

	// EnterIsolation_level is called when entering the isolation_level production.
	EnterIsolation_level(c *Isolation_levelContext)

	// EnterIsolation_types is called when entering the isolation_types production.
	EnterIsolation_types(c *Isolation_typesContext)

	// EnterSetExprOrDefault is called when entering the setExprOrDefault production.
	EnterSetExprOrDefault(c *SetExprOrDefaultContext)

	// EnterSetUserPropertyStatement is called when entering the setUserPropertyStatement production.
	EnterSetUserPropertyStatement(c *SetUserPropertyStatementContext)

	// EnterRoleList is called when entering the roleList production.
	EnterRoleList(c *RoleListContext)

	// EnterExecuteScriptStatement is called when entering the executeScriptStatement production.
	EnterExecuteScriptStatement(c *ExecuteScriptStatementContext)

	// EnterUnsupportedStatement is called when entering the unsupportedStatement production.
	EnterUnsupportedStatement(c *UnsupportedStatementContext)

	// EnterLock_item is called when entering the lock_item production.
	EnterLock_item(c *Lock_itemContext)

	// EnterLock_type is called when entering the lock_type production.
	EnterLock_type(c *Lock_typeContext)

	// EnterAlterPlanAdvisorAddStatement is called when entering the alterPlanAdvisorAddStatement production.
	EnterAlterPlanAdvisorAddStatement(c *AlterPlanAdvisorAddStatementContext)

	// EnterTruncatePlanAdvisorStatement is called when entering the truncatePlanAdvisorStatement production.
	EnterTruncatePlanAdvisorStatement(c *TruncatePlanAdvisorStatementContext)

	// EnterAlterPlanAdvisorDropStatement is called when entering the alterPlanAdvisorDropStatement production.
	EnterAlterPlanAdvisorDropStatement(c *AlterPlanAdvisorDropStatementContext)

	// EnterShowPlanAdvisorStatement is called when entering the showPlanAdvisorStatement production.
	EnterShowPlanAdvisorStatement(c *ShowPlanAdvisorStatementContext)

	// EnterCreateWarehouseStatement is called when entering the createWarehouseStatement production.
	EnterCreateWarehouseStatement(c *CreateWarehouseStatementContext)

	// EnterDropWarehouseStatement is called when entering the dropWarehouseStatement production.
	EnterDropWarehouseStatement(c *DropWarehouseStatementContext)

	// EnterSuspendWarehouseStatement is called when entering the suspendWarehouseStatement production.
	EnterSuspendWarehouseStatement(c *SuspendWarehouseStatementContext)

	// EnterResumeWarehouseStatement is called when entering the resumeWarehouseStatement production.
	EnterResumeWarehouseStatement(c *ResumeWarehouseStatementContext)

	// EnterSetWarehouseStatement is called when entering the setWarehouseStatement production.
	EnterSetWarehouseStatement(c *SetWarehouseStatementContext)

	// EnterShowWarehousesStatement is called when entering the showWarehousesStatement production.
	EnterShowWarehousesStatement(c *ShowWarehousesStatementContext)

	// EnterShowClustersStatement is called when entering the showClustersStatement production.
	EnterShowClustersStatement(c *ShowClustersStatementContext)

	// EnterShowNodesStatement is called when entering the showNodesStatement production.
	EnterShowNodesStatement(c *ShowNodesStatementContext)

	// EnterAlterWarehouseStatement is called when entering the alterWarehouseStatement production.
	EnterAlterWarehouseStatement(c *AlterWarehouseStatementContext)

	// EnterCreateCNGroupStatement is called when entering the createCNGroupStatement production.
	EnterCreateCNGroupStatement(c *CreateCNGroupStatementContext)

	// EnterDropCNGroupStatement is called when entering the dropCNGroupStatement production.
	EnterDropCNGroupStatement(c *DropCNGroupStatementContext)

	// EnterEnableCNGroupStatement is called when entering the enableCNGroupStatement production.
	EnterEnableCNGroupStatement(c *EnableCNGroupStatementContext)

	// EnterDisableCNGroupStatement is called when entering the disableCNGroupStatement production.
	EnterDisableCNGroupStatement(c *DisableCNGroupStatementContext)

	// EnterAlterCNGroupStatement is called when entering the alterCNGroupStatement production.
	EnterAlterCNGroupStatement(c *AlterCNGroupStatementContext)

	// EnterBeginStatement is called when entering the beginStatement production.
	EnterBeginStatement(c *BeginStatementContext)

	// EnterCommitStatement is called when entering the commitStatement production.
	EnterCommitStatement(c *CommitStatementContext)

	// EnterRollbackStatement is called when entering the rollbackStatement production.
	EnterRollbackStatement(c *RollbackStatementContext)

	// EnterTranslateStatement is called when entering the translateStatement production.
	EnterTranslateStatement(c *TranslateStatementContext)

	// EnterDialect is called when entering the dialect production.
	EnterDialect(c *DialectContext)

	// EnterTranslateSQL is called when entering the translateSQL production.
	EnterTranslateSQL(c *TranslateSQLContext)

	// EnterQueryStatement is called when entering the queryStatement production.
	EnterQueryStatement(c *QueryStatementContext)

	// EnterQueryRelation is called when entering the queryRelation production.
	EnterQueryRelation(c *QueryRelationContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterQueryNoWith is called when entering the queryNoWith production.
	EnterQueryNoWith(c *QueryNoWithContext)

	// EnterQueryPeriod is called when entering the queryPeriod production.
	EnterQueryPeriod(c *QueryPeriodContext)

	// EnterPeriodType is called when entering the periodType production.
	EnterPeriodType(c *PeriodTypeContext)

	// EnterQueryWithParentheses is called when entering the queryWithParentheses production.
	EnterQueryWithParentheses(c *QueryWithParenthesesContext)

	// EnterSetOperation is called when entering the setOperation production.
	EnterSetOperation(c *SetOperationContext)

	// EnterQueryPrimaryDefault is called when entering the queryPrimaryDefault production.
	EnterQueryPrimaryDefault(c *QueryPrimaryDefaultContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterRowConstructor is called when entering the rowConstructor production.
	EnterRowConstructor(c *RowConstructorContext)

	// EnterSortItem is called when entering the sortItem production.
	EnterSortItem(c *SortItemContext)

	// EnterLimitConstExpr is called when entering the limitConstExpr production.
	EnterLimitConstExpr(c *LimitConstExprContext)

	// EnterLimitElement is called when entering the limitElement production.
	EnterLimitElement(c *LimitElementContext)

	// EnterQuerySpecification is called when entering the querySpecification production.
	EnterQuerySpecification(c *QuerySpecificationContext)

	// EnterFrom is called when entering the from production.
	EnterFrom(c *FromContext)

	// EnterDual is called when entering the dual production.
	EnterDual(c *DualContext)

	// EnterRollup is called when entering the rollup production.
	EnterRollup(c *RollupContext)

	// EnterCube is called when entering the cube production.
	EnterCube(c *CubeContext)

	// EnterMultipleGroupingSets is called when entering the multipleGroupingSets production.
	EnterMultipleGroupingSets(c *MultipleGroupingSetsContext)

	// EnterSingleGroupingSet is called when entering the singleGroupingSet production.
	EnterSingleGroupingSet(c *SingleGroupingSetContext)

	// EnterGroupingSet is called when entering the groupingSet production.
	EnterGroupingSet(c *GroupingSetContext)

	// EnterCommonTableExpression is called when entering the commonTableExpression production.
	EnterCommonTableExpression(c *CommonTableExpressionContext)

	// EnterSetQuantifier is called when entering the setQuantifier production.
	EnterSetQuantifier(c *SetQuantifierContext)

	// EnterSelectSingle is called when entering the selectSingle production.
	EnterSelectSingle(c *SelectSingleContext)

	// EnterSelectAll is called when entering the selectAll production.
	EnterSelectAll(c *SelectAllContext)

	// EnterExcludeClause is called when entering the excludeClause production.
	EnterExcludeClause(c *ExcludeClauseContext)

	// EnterRelations is called when entering the relations production.
	EnterRelations(c *RelationsContext)

	// EnterRelationLateralView is called when entering the relationLateralView production.
	EnterRelationLateralView(c *RelationLateralViewContext)

	// EnterLateralView is called when entering the lateralView production.
	EnterLateralView(c *LateralViewContext)

	// EnterGeneratorFunction is called when entering the generatorFunction production.
	EnterGeneratorFunction(c *GeneratorFunctionContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterTableAtom is called when entering the tableAtom production.
	EnterTableAtom(c *TableAtomContext)

	// EnterInlineTable is called when entering the inlineTable production.
	EnterInlineTable(c *InlineTableContext)

	// EnterSubqueryWithAlias is called when entering the subqueryWithAlias production.
	EnterSubqueryWithAlias(c *SubqueryWithAliasContext)

	// EnterTableFunction is called when entering the tableFunction production.
	EnterTableFunction(c *TableFunctionContext)

	// EnterNormalizedTableFunction is called when entering the normalizedTableFunction production.
	EnterNormalizedTableFunction(c *NormalizedTableFunctionContext)

	// EnterFileTableFunction is called when entering the fileTableFunction production.
	EnterFileTableFunction(c *FileTableFunctionContext)

	// EnterParenthesizedRelation is called when entering the parenthesizedRelation production.
	EnterParenthesizedRelation(c *ParenthesizedRelationContext)

	// EnterPivotClause is called when entering the pivotClause production.
	EnterPivotClause(c *PivotClauseContext)

	// EnterPivotAggregationExpression is called when entering the pivotAggregationExpression production.
	EnterPivotAggregationExpression(c *PivotAggregationExpressionContext)

	// EnterPivotValue is called when entering the pivotValue production.
	EnterPivotValue(c *PivotValueContext)

	// EnterSampleClause is called when entering the sampleClause production.
	EnterSampleClause(c *SampleClauseContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterNamedArgumentList is called when entering the namedArgumentList production.
	EnterNamedArgumentList(c *NamedArgumentListContext)

	// EnterNamedArguments is called when entering the namedArguments production.
	EnterNamedArguments(c *NamedArgumentsContext)

	// EnterJoinRelation is called when entering the joinRelation production.
	EnterJoinRelation(c *JoinRelationContext)

	// EnterCrossOrInnerJoinType is called when entering the crossOrInnerJoinType production.
	EnterCrossOrInnerJoinType(c *CrossOrInnerJoinTypeContext)

	// EnterOuterAndSemiJoinType is called when entering the outerAndSemiJoinType production.
	EnterOuterAndSemiJoinType(c *OuterAndSemiJoinTypeContext)

	// EnterBracketHint is called when entering the bracketHint production.
	EnterBracketHint(c *BracketHintContext)

	// EnterHintMap is called when entering the hintMap production.
	EnterHintMap(c *HintMapContext)

	// EnterJoinCriteria is called when entering the joinCriteria production.
	EnterJoinCriteria(c *JoinCriteriaContext)

	// EnterColumnAliases is called when entering the columnAliases production.
	EnterColumnAliases(c *ColumnAliasesContext)

	// EnterColumnAliasesWithoutParentheses is called when entering the columnAliasesWithoutParentheses production.
	EnterColumnAliasesWithoutParentheses(c *ColumnAliasesWithoutParenthesesContext)

	// EnterPartitionNames is called when entering the partitionNames production.
	EnterPartitionNames(c *PartitionNamesContext)

	// EnterKeyPartitionList is called when entering the keyPartitionList production.
	EnterKeyPartitionList(c *KeyPartitionListContext)

	// EnterTabletList is called when entering the tabletList production.
	EnterTabletList(c *TabletListContext)

	// EnterPrepareStatement is called when entering the prepareStatement production.
	EnterPrepareStatement(c *PrepareStatementContext)

	// EnterPrepareSql is called when entering the prepareSql production.
	EnterPrepareSql(c *PrepareSqlContext)

	// EnterExecuteStatement is called when entering the executeStatement production.
	EnterExecuteStatement(c *ExecuteStatementContext)

	// EnterDeallocateStatement is called when entering the deallocateStatement production.
	EnterDeallocateStatement(c *DeallocateStatementContext)

	// EnterReplicaList is called when entering the replicaList production.
	EnterReplicaList(c *ReplicaListContext)

	// EnterExpressionsWithDefault is called when entering the expressionsWithDefault production.
	EnterExpressionsWithDefault(c *ExpressionsWithDefaultContext)

	// EnterExpressionOrDefault is called when entering the expressionOrDefault production.
	EnterExpressionOrDefault(c *ExpressionOrDefaultContext)

	// EnterMapExpressionList is called when entering the mapExpressionList production.
	EnterMapExpressionList(c *MapExpressionListContext)

	// EnterMapExpression is called when entering the mapExpression production.
	EnterMapExpression(c *MapExpressionContext)

	// EnterExpressionSingleton is called when entering the expressionSingleton production.
	EnterExpressionSingleton(c *ExpressionSingletonContext)

	// EnterExpressionDefault is called when entering the expressionDefault production.
	EnterExpressionDefault(c *ExpressionDefaultContext)

	// EnterLogicalNot is called when entering the logicalNot production.
	EnterLogicalNot(c *LogicalNotContext)

	// EnterLogicalBinary is called when entering the logicalBinary production.
	EnterLogicalBinary(c *LogicalBinaryContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterBooleanExpressionDefault is called when entering the booleanExpressionDefault production.
	EnterBooleanExpressionDefault(c *BooleanExpressionDefaultContext)

	// EnterIsNull is called when entering the isNull production.
	EnterIsNull(c *IsNullContext)

	// EnterScalarSubquery is called when entering the scalarSubquery production.
	EnterScalarSubquery(c *ScalarSubqueryContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterTupleInSubquery is called when entering the tupleInSubquery production.
	EnterTupleInSubquery(c *TupleInSubqueryContext)

	// EnterInSubquery is called when entering the inSubquery production.
	EnterInSubquery(c *InSubqueryContext)

	// EnterInList is called when entering the inList production.
	EnterInList(c *InListContext)

	// EnterBetween is called when entering the between production.
	EnterBetween(c *BetweenContext)

	// EnterLike is called when entering the like production.
	EnterLike(c *LikeContext)

	// EnterValueExpressionDefault is called when entering the valueExpressionDefault production.
	EnterValueExpressionDefault(c *ValueExpressionDefaultContext)

	// EnterArithmeticBinary is called when entering the arithmeticBinary production.
	EnterArithmeticBinary(c *ArithmeticBinaryContext)

	// EnterDereference is called when entering the dereference production.
	EnterDereference(c *DereferenceContext)

	// EnterOdbcFunctionCallExpression is called when entering the odbcFunctionCallExpression production.
	EnterOdbcFunctionCallExpression(c *OdbcFunctionCallExpressionContext)

	// EnterMatchExpr is called when entering the matchExpr production.
	EnterMatchExpr(c *MatchExprContext)

	// EnterColumnRef is called when entering the columnRef production.
	EnterColumnRef(c *ColumnRefContext)

	// EnterConvert is called when entering the convert production.
	EnterConvert(c *ConvertContext)

	// EnterCollectionSubscript is called when entering the collectionSubscript production.
	EnterCollectionSubscript(c *CollectionSubscriptContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterUserVariableExpression is called when entering the userVariableExpression production.
	EnterUserVariableExpression(c *UserVariableExpressionContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterSimpleCase is called when entering the simpleCase production.
	EnterSimpleCase(c *SimpleCaseContext)

	// EnterArrowExpression is called when entering the arrowExpression production.
	EnterArrowExpression(c *ArrowExpressionContext)

	// EnterArrayExpr is called when entering the arrayExpr production.
	EnterArrayExpr(c *ArrayExprContext)

	// EnterSystemVariableExpression is called when entering the systemVariableExpression production.
	EnterSystemVariableExpression(c *SystemVariableExpressionContext)

	// EnterConcat is called when entering the concat production.
	EnterConcat(c *ConcatContext)

	// EnterSubqueryExpression is called when entering the subqueryExpression production.
	EnterSubqueryExpression(c *SubqueryExpressionContext)

	// EnterLambdaFunctionExpr is called when entering the lambdaFunctionExpr production.
	EnterLambdaFunctionExpr(c *LambdaFunctionExprContext)

	// EnterDictionaryGetExpr is called when entering the dictionaryGetExpr production.
	EnterDictionaryGetExpr(c *DictionaryGetExprContext)

	// EnterCollate is called when entering the collate production.
	EnterCollate(c *CollateContext)

	// EnterArrayConstructor is called when entering the arrayConstructor production.
	EnterArrayConstructor(c *ArrayConstructorContext)

	// EnterMapConstructor is called when entering the mapConstructor production.
	EnterMapConstructor(c *MapConstructorContext)

	// EnterArraySlice is called when entering the arraySlice production.
	EnterArraySlice(c *ArraySliceContext)

	// EnterExists is called when entering the exists production.
	EnterExists(c *ExistsContext)

	// EnterSearchedCase is called when entering the searchedCase production.
	EnterSearchedCase(c *SearchedCaseContext)

	// EnterArithmeticUnary is called when entering the arithmeticUnary production.
	EnterArithmeticUnary(c *ArithmeticUnaryContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterDateLiteral is called when entering the dateLiteral production.
	EnterDateLiteral(c *DateLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIntervalLiteral is called when entering the intervalLiteral production.
	EnterIntervalLiteral(c *IntervalLiteralContext)

	// EnterUnitBoundaryLiteral is called when entering the unitBoundaryLiteral production.
	EnterUnitBoundaryLiteral(c *UnitBoundaryLiteralContext)

	// EnterBinaryLiteral is called when entering the binaryLiteral production.
	EnterBinaryLiteral(c *BinaryLiteralContext)

	// EnterParameter is called when entering the Parameter production.
	EnterParameter(c *ParameterContext)

	// EnterExtract is called when entering the extract production.
	EnterExtract(c *ExtractContext)

	// EnterGroupingOperation is called when entering the groupingOperation production.
	EnterGroupingOperation(c *GroupingOperationContext)

	// EnterInformationFunction is called when entering the informationFunction production.
	EnterInformationFunction(c *InformationFunctionContext)

	// EnterSpecialDateTime is called when entering the specialDateTime production.
	EnterSpecialDateTime(c *SpecialDateTimeContext)

	// EnterSpecialFunction is called when entering the specialFunction production.
	EnterSpecialFunction(c *SpecialFunctionContext)

	// EnterAggregationFunctionCall is called when entering the aggregationFunctionCall production.
	EnterAggregationFunctionCall(c *AggregationFunctionCallContext)

	// EnterWindowFunctionCall is called when entering the windowFunctionCall production.
	EnterWindowFunctionCall(c *WindowFunctionCallContext)

	// EnterTranslateFunctionCall is called when entering the translateFunctionCall production.
	EnterTranslateFunctionCall(c *TranslateFunctionCallContext)

	// EnterSimpleFunctionCall is called when entering the simpleFunctionCall production.
	EnterSimpleFunctionCall(c *SimpleFunctionCallContext)

	// EnterAggregationFunction is called when entering the aggregationFunction production.
	EnterAggregationFunction(c *AggregationFunctionContext)

	// EnterUserVariable is called when entering the userVariable production.
	EnterUserVariable(c *UserVariableContext)

	// EnterSystemVariable is called when entering the systemVariable production.
	EnterSystemVariable(c *SystemVariableContext)

	// EnterColumnReference is called when entering the columnReference production.
	EnterColumnReference(c *ColumnReferenceContext)

	// EnterInformationFunctionExpression is called when entering the informationFunctionExpression production.
	EnterInformationFunctionExpression(c *InformationFunctionExpressionContext)

	// EnterSpecialDateTimeExpression is called when entering the specialDateTimeExpression production.
	EnterSpecialDateTimeExpression(c *SpecialDateTimeExpressionContext)

	// EnterSpecialFunctionExpression is called when entering the specialFunctionExpression production.
	EnterSpecialFunctionExpression(c *SpecialFunctionExpressionContext)

	// EnterWindowFunction is called when entering the windowFunction production.
	EnterWindowFunction(c *WindowFunctionContext)

	// EnterWhenClause is called when entering the whenClause production.
	EnterWhenClause(c *WhenClauseContext)

	// EnterOver is called when entering the over production.
	EnterOver(c *OverContext)

	// EnterIgnoreNulls is called when entering the ignoreNulls production.
	EnterIgnoreNulls(c *IgnoreNullsContext)

	// EnterWindowFrame is called when entering the windowFrame production.
	EnterWindowFrame(c *WindowFrameContext)

	// EnterUnboundedFrame is called when entering the unboundedFrame production.
	EnterUnboundedFrame(c *UnboundedFrameContext)

	// EnterCurrentRowBound is called when entering the currentRowBound production.
	EnterCurrentRowBound(c *CurrentRowBoundContext)

	// EnterBoundedFrame is called when entering the boundedFrame production.
	EnterBoundedFrame(c *BoundedFrameContext)

	// EnterBackupRestoreObjectDesc is called when entering the backupRestoreObjectDesc production.
	EnterBackupRestoreObjectDesc(c *BackupRestoreObjectDescContext)

	// EnterTableDesc is called when entering the tableDesc production.
	EnterTableDesc(c *TableDescContext)

	// EnterBackupRestoreTableDesc is called when entering the backupRestoreTableDesc production.
	EnterBackupRestoreTableDesc(c *BackupRestoreTableDescContext)

	// EnterExplainDesc is called when entering the explainDesc production.
	EnterExplainDesc(c *ExplainDescContext)

	// EnterOptimizerTrace is called when entering the optimizerTrace production.
	EnterOptimizerTrace(c *OptimizerTraceContext)

	// EnterPartitionExpr is called when entering the partitionExpr production.
	EnterPartitionExpr(c *PartitionExprContext)

	// EnterPartitionDesc is called when entering the partitionDesc production.
	EnterPartitionDesc(c *PartitionDescContext)

	// EnterListPartitionDesc is called when entering the listPartitionDesc production.
	EnterListPartitionDesc(c *ListPartitionDescContext)

	// EnterSingleItemListPartitionDesc is called when entering the singleItemListPartitionDesc production.
	EnterSingleItemListPartitionDesc(c *SingleItemListPartitionDescContext)

	// EnterMultiItemListPartitionDesc is called when entering the multiItemListPartitionDesc production.
	EnterMultiItemListPartitionDesc(c *MultiItemListPartitionDescContext)

	// EnterMultiListPartitionValues is called when entering the multiListPartitionValues production.
	EnterMultiListPartitionValues(c *MultiListPartitionValuesContext)

	// EnterSingleListPartitionValues is called when entering the singleListPartitionValues production.
	EnterSingleListPartitionValues(c *SingleListPartitionValuesContext)

	// EnterListPartitionValues is called when entering the listPartitionValues production.
	EnterListPartitionValues(c *ListPartitionValuesContext)

	// EnterListPartitionValue is called when entering the listPartitionValue production.
	EnterListPartitionValue(c *ListPartitionValueContext)

	// EnterStringList is called when entering the stringList production.
	EnterStringList(c *StringListContext)

	// EnterLiteralExpressionList is called when entering the literalExpressionList production.
	EnterLiteralExpressionList(c *LiteralExpressionListContext)

	// EnterRangePartitionDesc is called when entering the rangePartitionDesc production.
	EnterRangePartitionDesc(c *RangePartitionDescContext)

	// EnterSingleRangePartition is called when entering the singleRangePartition production.
	EnterSingleRangePartition(c *SingleRangePartitionContext)

	// EnterMultiRangePartition is called when entering the multiRangePartition production.
	EnterMultiRangePartition(c *MultiRangePartitionContext)

	// EnterPartitionRangeDesc is called when entering the partitionRangeDesc production.
	EnterPartitionRangeDesc(c *PartitionRangeDescContext)

	// EnterPartitionKeyDesc is called when entering the partitionKeyDesc production.
	EnterPartitionKeyDesc(c *PartitionKeyDescContext)

	// EnterPartitionValueList is called when entering the partitionValueList production.
	EnterPartitionValueList(c *PartitionValueListContext)

	// EnterKeyPartition is called when entering the keyPartition production.
	EnterKeyPartition(c *KeyPartitionContext)

	// EnterPartitionValue is called when entering the partitionValue production.
	EnterPartitionValue(c *PartitionValueContext)

	// EnterDistributionClause is called when entering the distributionClause production.
	EnterDistributionClause(c *DistributionClauseContext)

	// EnterDistributionDesc is called when entering the distributionDesc production.
	EnterDistributionDesc(c *DistributionDescContext)

	// EnterRefreshSchemeDesc is called when entering the refreshSchemeDesc production.
	EnterRefreshSchemeDesc(c *RefreshSchemeDescContext)

	// EnterStatusDesc is called when entering the statusDesc production.
	EnterStatusDesc(c *StatusDescContext)

	// EnterProperties is called when entering the properties production.
	EnterProperties(c *PropertiesContext)

	// EnterExtProperties is called when entering the extProperties production.
	EnterExtProperties(c *ExtPropertiesContext)

	// EnterPropertyList is called when entering the propertyList production.
	EnterPropertyList(c *PropertyListContext)

	// EnterUserPropertyList is called when entering the userPropertyList production.
	EnterUserPropertyList(c *UserPropertyListContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterInlineProperties is called when entering the inlineProperties production.
	EnterInlineProperties(c *InlinePropertiesContext)

	// EnterInlineProperty is called when entering the inlineProperty production.
	EnterInlineProperty(c *InlinePropertyContext)

	// EnterVarType is called when entering the varType production.
	EnterVarType(c *VarTypeContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterOutfile is called when entering the outfile production.
	EnterOutfile(c *OutfileContext)

	// EnterFileFormat is called when entering the fileFormat production.
	EnterFileFormat(c *FileFormatContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterTaskInterval is called when entering the taskInterval production.
	EnterTaskInterval(c *TaskIntervalContext)

	// EnterTaskUnitIdentifier is called when entering the taskUnitIdentifier production.
	EnterTaskUnitIdentifier(c *TaskUnitIdentifierContext)

	// EnterUnitIdentifier is called when entering the unitIdentifier production.
	EnterUnitIdentifier(c *UnitIdentifierContext)

	// EnterUnitBoundary is called when entering the unitBoundary production.
	EnterUnitBoundary(c *UnitBoundaryContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterSubfieldDesc is called when entering the subfieldDesc production.
	EnterSubfieldDesc(c *SubfieldDescContext)

	// EnterSubfieldDescs is called when entering the subfieldDescs production.
	EnterSubfieldDescs(c *SubfieldDescsContext)

	// EnterStructType is called when entering the structType production.
	EnterStructType(c *StructTypeContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterDecimalType is called when entering the decimalType production.
	EnterDecimalType(c *DecimalTypeContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterWriteBranch is called when entering the writeBranch production.
	EnterWriteBranch(c *WriteBranchContext)

	// EnterUnquotedIdentifier is called when entering the unquotedIdentifier production.
	EnterUnquotedIdentifier(c *UnquotedIdentifierContext)

	// EnterDigitIdentifier is called when entering the digitIdentifier production.
	EnterDigitIdentifier(c *DigitIdentifierContext)

	// EnterBackQuotedIdentifier is called when entering the backQuotedIdentifier production.
	EnterBackQuotedIdentifier(c *BackQuotedIdentifierContext)

	// EnterIdentifierWithAlias is called when entering the identifierWithAlias production.
	EnterIdentifierWithAlias(c *IdentifierWithAliasContext)

	// EnterIdentifierWithAliasList is called when entering the identifierWithAliasList production.
	EnterIdentifierWithAliasList(c *IdentifierWithAliasListContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterIdentifierOrString is called when entering the identifierOrString production.
	EnterIdentifierOrString(c *IdentifierOrStringContext)

	// EnterIdentifierOrStringList is called when entering the identifierOrStringList production.
	EnterIdentifierOrStringList(c *IdentifierOrStringListContext)

	// EnterIdentifierOrStringOrStar is called when entering the identifierOrStringOrStar production.
	EnterIdentifierOrStringOrStar(c *IdentifierOrStringOrStarContext)

	// EnterUserWithoutHost is called when entering the userWithoutHost production.
	EnterUserWithoutHost(c *UserWithoutHostContext)

	// EnterUserWithHost is called when entering the userWithHost production.
	EnterUserWithHost(c *UserWithHostContext)

	// EnterUserWithHostAndBlanket is called when entering the userWithHostAndBlanket production.
	EnterUserWithHostAndBlanket(c *UserWithHostAndBlanketContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAssignmentList is called when entering the assignmentList production.
	EnterAssignmentList(c *AssignmentListContext)

	// EnterDecimalValue is called when entering the decimalValue production.
	EnterDecimalValue(c *DecimalValueContext)

	// EnterDoubleValue is called when entering the doubleValue production.
	EnterDoubleValue(c *DoubleValueContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterNonReserved is called when entering the nonReserved production.
	EnterNonReserved(c *NonReservedContext)

	// ExitSqlStatements is called when exiting the sqlStatements production.
	ExitSqlStatements(c *SqlStatementsContext)

	// ExitSingleStatement is called when exiting the singleStatement production.
	ExitSingleStatement(c *SingleStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitUseDatabaseStatement is called when exiting the useDatabaseStatement production.
	ExitUseDatabaseStatement(c *UseDatabaseStatementContext)

	// ExitUseCatalogStatement is called when exiting the useCatalogStatement production.
	ExitUseCatalogStatement(c *UseCatalogStatementContext)

	// ExitSetCatalogStatement is called when exiting the setCatalogStatement production.
	ExitSetCatalogStatement(c *SetCatalogStatementContext)

	// ExitShowDatabasesStatement is called when exiting the showDatabasesStatement production.
	ExitShowDatabasesStatement(c *ShowDatabasesStatementContext)

	// ExitAlterDbQuotaStatement is called when exiting the alterDbQuotaStatement production.
	ExitAlterDbQuotaStatement(c *AlterDbQuotaStatementContext)

	// ExitCreateDbStatement is called when exiting the createDbStatement production.
	ExitCreateDbStatement(c *CreateDbStatementContext)

	// ExitDropDbStatement is called when exiting the dropDbStatement production.
	ExitDropDbStatement(c *DropDbStatementContext)

	// ExitShowCreateDbStatement is called when exiting the showCreateDbStatement production.
	ExitShowCreateDbStatement(c *ShowCreateDbStatementContext)

	// ExitAlterDatabaseRenameStatement is called when exiting the alterDatabaseRenameStatement production.
	ExitAlterDatabaseRenameStatement(c *AlterDatabaseRenameStatementContext)

	// ExitRecoverDbStmt is called when exiting the recoverDbStmt production.
	ExitRecoverDbStmt(c *RecoverDbStmtContext)

	// ExitShowDataStmt is called when exiting the showDataStmt production.
	ExitShowDataStmt(c *ShowDataStmtContext)

	// ExitShowDataDistributionStmt is called when exiting the showDataDistributionStmt production.
	ExitShowDataDistributionStmt(c *ShowDataDistributionStmtContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitColumnDesc is called when exiting the columnDesc production.
	ExitColumnDesc(c *ColumnDescContext)

	// ExitCharsetName is called when exiting the charsetName production.
	ExitCharsetName(c *CharsetNameContext)

	// ExitDefaultDesc is called when exiting the defaultDesc production.
	ExitDefaultDesc(c *DefaultDescContext)

	// ExitGeneratedColumnDesc is called when exiting the generatedColumnDesc production.
	ExitGeneratedColumnDesc(c *GeneratedColumnDescContext)

	// ExitIndexDesc is called when exiting the indexDesc production.
	ExitIndexDesc(c *IndexDescContext)

	// ExitEngineDesc is called when exiting the engineDesc production.
	ExitEngineDesc(c *EngineDescContext)

	// ExitCharsetDesc is called when exiting the charsetDesc production.
	ExitCharsetDesc(c *CharsetDescContext)

	// ExitCollateDesc is called when exiting the collateDesc production.
	ExitCollateDesc(c *CollateDescContext)

	// ExitKeyDesc is called when exiting the keyDesc production.
	ExitKeyDesc(c *KeyDescContext)

	// ExitOrderByDesc is called when exiting the orderByDesc production.
	ExitOrderByDesc(c *OrderByDescContext)

	// ExitColumnNullable is called when exiting the columnNullable production.
	ExitColumnNullable(c *ColumnNullableContext)

	// ExitTypeWithNullable is called when exiting the typeWithNullable production.
	ExitTypeWithNullable(c *TypeWithNullableContext)

	// ExitAggStateDesc is called when exiting the aggStateDesc production.
	ExitAggStateDesc(c *AggStateDescContext)

	// ExitAggDesc is called when exiting the aggDesc production.
	ExitAggDesc(c *AggDescContext)

	// ExitRollupDesc is called when exiting the rollupDesc production.
	ExitRollupDesc(c *RollupDescContext)

	// ExitRollupItem is called when exiting the rollupItem production.
	ExitRollupItem(c *RollupItemContext)

	// ExitDupKeys is called when exiting the dupKeys production.
	ExitDupKeys(c *DupKeysContext)

	// ExitFromRollup is called when exiting the fromRollup production.
	ExitFromRollup(c *FromRollupContext)

	// ExitOrReplace is called when exiting the orReplace production.
	ExitOrReplace(c *OrReplaceContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)

	// ExitCreateTableAsSelectStatement is called when exiting the createTableAsSelectStatement production.
	ExitCreateTableAsSelectStatement(c *CreateTableAsSelectStatementContext)

	// ExitDropTableStatement is called when exiting the dropTableStatement production.
	ExitDropTableStatement(c *DropTableStatementContext)

	// ExitCleanTemporaryTableStatement is called when exiting the cleanTemporaryTableStatement production.
	ExitCleanTemporaryTableStatement(c *CleanTemporaryTableStatementContext)

	// ExitAlterTableStatement is called when exiting the alterTableStatement production.
	ExitAlterTableStatement(c *AlterTableStatementContext)

	// ExitCreateIndexStatement is called when exiting the createIndexStatement production.
	ExitCreateIndexStatement(c *CreateIndexStatementContext)

	// ExitDropIndexStatement is called when exiting the dropIndexStatement production.
	ExitDropIndexStatement(c *DropIndexStatementContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitShowTableStatement is called when exiting the showTableStatement production.
	ExitShowTableStatement(c *ShowTableStatementContext)

	// ExitShowTemporaryTablesStatement is called when exiting the showTemporaryTablesStatement production.
	ExitShowTemporaryTablesStatement(c *ShowTemporaryTablesStatementContext)

	// ExitShowCreateTableStatement is called when exiting the showCreateTableStatement production.
	ExitShowCreateTableStatement(c *ShowCreateTableStatementContext)

	// ExitShowColumnStatement is called when exiting the showColumnStatement production.
	ExitShowColumnStatement(c *ShowColumnStatementContext)

	// ExitShowTableStatusStatement is called when exiting the showTableStatusStatement production.
	ExitShowTableStatusStatement(c *ShowTableStatusStatementContext)

	// ExitRefreshTableStatement is called when exiting the refreshTableStatement production.
	ExitRefreshTableStatement(c *RefreshTableStatementContext)

	// ExitShowAlterStatement is called when exiting the showAlterStatement production.
	ExitShowAlterStatement(c *ShowAlterStatementContext)

	// ExitDescTableStatement is called when exiting the descTableStatement production.
	ExitDescTableStatement(c *DescTableStatementContext)

	// ExitCreateTableLikeStatement is called when exiting the createTableLikeStatement production.
	ExitCreateTableLikeStatement(c *CreateTableLikeStatementContext)

	// ExitShowIndexStatement is called when exiting the showIndexStatement production.
	ExitShowIndexStatement(c *ShowIndexStatementContext)

	// ExitRecoverTableStatement is called when exiting the recoverTableStatement production.
	ExitRecoverTableStatement(c *RecoverTableStatementContext)

	// ExitTruncateTableStatement is called when exiting the truncateTableStatement production.
	ExitTruncateTableStatement(c *TruncateTableStatementContext)

	// ExitCancelAlterTableStatement is called when exiting the cancelAlterTableStatement production.
	ExitCancelAlterTableStatement(c *CancelAlterTableStatementContext)

	// ExitShowPartitionsStatement is called when exiting the showPartitionsStatement production.
	ExitShowPartitionsStatement(c *ShowPartitionsStatementContext)

	// ExitRecoverPartitionStatement is called when exiting the recoverPartitionStatement production.
	ExitRecoverPartitionStatement(c *RecoverPartitionStatementContext)

	// ExitCreateViewStatement is called when exiting the createViewStatement production.
	ExitCreateViewStatement(c *CreateViewStatementContext)

	// ExitAlterViewStatement is called when exiting the alterViewStatement production.
	ExitAlterViewStatement(c *AlterViewStatementContext)

	// ExitDropViewStatement is called when exiting the dropViewStatement production.
	ExitDropViewStatement(c *DropViewStatementContext)

	// ExitColumnNameWithComment is called when exiting the columnNameWithComment production.
	ExitColumnNameWithComment(c *ColumnNameWithCommentContext)

	// ExitSubmitTaskStatement is called when exiting the submitTaskStatement production.
	ExitSubmitTaskStatement(c *SubmitTaskStatementContext)

	// ExitTaskClause is called when exiting the taskClause production.
	ExitTaskClause(c *TaskClauseContext)

	// ExitDropTaskStatement is called when exiting the dropTaskStatement production.
	ExitDropTaskStatement(c *DropTaskStatementContext)

	// ExitTaskScheduleDesc is called when exiting the taskScheduleDesc production.
	ExitTaskScheduleDesc(c *TaskScheduleDescContext)

	// ExitCreateMaterializedViewStatement is called when exiting the createMaterializedViewStatement production.
	ExitCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// ExitMvPartitionExprs is called when exiting the mvPartitionExprs production.
	ExitMvPartitionExprs(c *MvPartitionExprsContext)

	// ExitMaterializedViewDesc is called when exiting the materializedViewDesc production.
	ExitMaterializedViewDesc(c *MaterializedViewDescContext)

	// ExitShowMaterializedViewsStatement is called when exiting the showMaterializedViewsStatement production.
	ExitShowMaterializedViewsStatement(c *ShowMaterializedViewsStatementContext)

	// ExitDropMaterializedViewStatement is called when exiting the dropMaterializedViewStatement production.
	ExitDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// ExitAlterMaterializedViewStatement is called when exiting the alterMaterializedViewStatement production.
	ExitAlterMaterializedViewStatement(c *AlterMaterializedViewStatementContext)

	// ExitRefreshMaterializedViewStatement is called when exiting the refreshMaterializedViewStatement production.
	ExitRefreshMaterializedViewStatement(c *RefreshMaterializedViewStatementContext)

	// ExitCancelRefreshMaterializedViewStatement is called when exiting the cancelRefreshMaterializedViewStatement production.
	ExitCancelRefreshMaterializedViewStatement(c *CancelRefreshMaterializedViewStatementContext)

	// ExitAdminSetConfigStatement is called when exiting the adminSetConfigStatement production.
	ExitAdminSetConfigStatement(c *AdminSetConfigStatementContext)

	// ExitAdminSetReplicaStatusStatement is called when exiting the adminSetReplicaStatusStatement production.
	ExitAdminSetReplicaStatusStatement(c *AdminSetReplicaStatusStatementContext)

	// ExitAdminShowConfigStatement is called when exiting the adminShowConfigStatement production.
	ExitAdminShowConfigStatement(c *AdminShowConfigStatementContext)

	// ExitAdminShowReplicaDistributionStatement is called when exiting the adminShowReplicaDistributionStatement production.
	ExitAdminShowReplicaDistributionStatement(c *AdminShowReplicaDistributionStatementContext)

	// ExitAdminShowReplicaStatusStatement is called when exiting the adminShowReplicaStatusStatement production.
	ExitAdminShowReplicaStatusStatement(c *AdminShowReplicaStatusStatementContext)

	// ExitAdminRepairTableStatement is called when exiting the adminRepairTableStatement production.
	ExitAdminRepairTableStatement(c *AdminRepairTableStatementContext)

	// ExitAdminCancelRepairTableStatement is called when exiting the adminCancelRepairTableStatement production.
	ExitAdminCancelRepairTableStatement(c *AdminCancelRepairTableStatementContext)

	// ExitAdminCheckTabletsStatement is called when exiting the adminCheckTabletsStatement production.
	ExitAdminCheckTabletsStatement(c *AdminCheckTabletsStatementContext)

	// ExitAdminSetPartitionVersion is called when exiting the adminSetPartitionVersion production.
	ExitAdminSetPartitionVersion(c *AdminSetPartitionVersionContext)

	// ExitKillStatement is called when exiting the killStatement production.
	ExitKillStatement(c *KillStatementContext)

	// ExitSyncStatement is called when exiting the syncStatement production.
	ExitSyncStatement(c *SyncStatementContext)

	// ExitAdminSetAutomatedSnapshotOnStatement is called when exiting the adminSetAutomatedSnapshotOnStatement production.
	ExitAdminSetAutomatedSnapshotOnStatement(c *AdminSetAutomatedSnapshotOnStatementContext)

	// ExitAdminSetAutomatedSnapshotOffStatement is called when exiting the adminSetAutomatedSnapshotOffStatement production.
	ExitAdminSetAutomatedSnapshotOffStatement(c *AdminSetAutomatedSnapshotOffStatementContext)

	// ExitAlterSystemStatement is called when exiting the alterSystemStatement production.
	ExitAlterSystemStatement(c *AlterSystemStatementContext)

	// ExitCancelAlterSystemStatement is called when exiting the cancelAlterSystemStatement production.
	ExitCancelAlterSystemStatement(c *CancelAlterSystemStatementContext)

	// ExitShowComputeNodesStatement is called when exiting the showComputeNodesStatement production.
	ExitShowComputeNodesStatement(c *ShowComputeNodesStatementContext)

	// ExitCreateExternalCatalogStatement is called when exiting the createExternalCatalogStatement production.
	ExitCreateExternalCatalogStatement(c *CreateExternalCatalogStatementContext)

	// ExitShowCreateExternalCatalogStatement is called when exiting the showCreateExternalCatalogStatement production.
	ExitShowCreateExternalCatalogStatement(c *ShowCreateExternalCatalogStatementContext)

	// ExitDropExternalCatalogStatement is called when exiting the dropExternalCatalogStatement production.
	ExitDropExternalCatalogStatement(c *DropExternalCatalogStatementContext)

	// ExitShowCatalogsStatement is called when exiting the showCatalogsStatement production.
	ExitShowCatalogsStatement(c *ShowCatalogsStatementContext)

	// ExitAlterCatalogStatement is called when exiting the alterCatalogStatement production.
	ExitAlterCatalogStatement(c *AlterCatalogStatementContext)

	// ExitCreateStorageVolumeStatement is called when exiting the createStorageVolumeStatement production.
	ExitCreateStorageVolumeStatement(c *CreateStorageVolumeStatementContext)

	// ExitTypeDesc is called when exiting the typeDesc production.
	ExitTypeDesc(c *TypeDescContext)

	// ExitLocationsDesc is called when exiting the locationsDesc production.
	ExitLocationsDesc(c *LocationsDescContext)

	// ExitShowStorageVolumesStatement is called when exiting the showStorageVolumesStatement production.
	ExitShowStorageVolumesStatement(c *ShowStorageVolumesStatementContext)

	// ExitDropStorageVolumeStatement is called when exiting the dropStorageVolumeStatement production.
	ExitDropStorageVolumeStatement(c *DropStorageVolumeStatementContext)

	// ExitAlterStorageVolumeStatement is called when exiting the alterStorageVolumeStatement production.
	ExitAlterStorageVolumeStatement(c *AlterStorageVolumeStatementContext)

	// ExitAlterStorageVolumeClause is called when exiting the alterStorageVolumeClause production.
	ExitAlterStorageVolumeClause(c *AlterStorageVolumeClauseContext)

	// ExitModifyStorageVolumePropertiesClause is called when exiting the modifyStorageVolumePropertiesClause production.
	ExitModifyStorageVolumePropertiesClause(c *ModifyStorageVolumePropertiesClauseContext)

	// ExitModifyStorageVolumeCommentClause is called when exiting the modifyStorageVolumeCommentClause production.
	ExitModifyStorageVolumeCommentClause(c *ModifyStorageVolumeCommentClauseContext)

	// ExitDescStorageVolumeStatement is called when exiting the descStorageVolumeStatement production.
	ExitDescStorageVolumeStatement(c *DescStorageVolumeStatementContext)

	// ExitSetDefaultStorageVolumeStatement is called when exiting the setDefaultStorageVolumeStatement production.
	ExitSetDefaultStorageVolumeStatement(c *SetDefaultStorageVolumeStatementContext)

	// ExitUpdateFailPointStatusStatement is called when exiting the updateFailPointStatusStatement production.
	ExitUpdateFailPointStatusStatement(c *UpdateFailPointStatusStatementContext)

	// ExitShowFailPointStatement is called when exiting the showFailPointStatement production.
	ExitShowFailPointStatement(c *ShowFailPointStatementContext)

	// ExitCreateDictionaryStatement is called when exiting the createDictionaryStatement production.
	ExitCreateDictionaryStatement(c *CreateDictionaryStatementContext)

	// ExitDropDictionaryStatement is called when exiting the dropDictionaryStatement production.
	ExitDropDictionaryStatement(c *DropDictionaryStatementContext)

	// ExitRefreshDictionaryStatement is called when exiting the refreshDictionaryStatement production.
	ExitRefreshDictionaryStatement(c *RefreshDictionaryStatementContext)

	// ExitShowDictionaryStatement is called when exiting the showDictionaryStatement production.
	ExitShowDictionaryStatement(c *ShowDictionaryStatementContext)

	// ExitCancelRefreshDictionaryStatement is called when exiting the cancelRefreshDictionaryStatement production.
	ExitCancelRefreshDictionaryStatement(c *CancelRefreshDictionaryStatementContext)

	// ExitDictionaryColumnDesc is called when exiting the dictionaryColumnDesc production.
	ExitDictionaryColumnDesc(c *DictionaryColumnDescContext)

	// ExitDictionaryName is called when exiting the dictionaryName production.
	ExitDictionaryName(c *DictionaryNameContext)

	// ExitAlterClause is called when exiting the alterClause production.
	ExitAlterClause(c *AlterClauseContext)

	// ExitAddFrontendClause is called when exiting the addFrontendClause production.
	ExitAddFrontendClause(c *AddFrontendClauseContext)

	// ExitDropFrontendClause is called when exiting the dropFrontendClause production.
	ExitDropFrontendClause(c *DropFrontendClauseContext)

	// ExitModifyFrontendHostClause is called when exiting the modifyFrontendHostClause production.
	ExitModifyFrontendHostClause(c *ModifyFrontendHostClauseContext)

	// ExitAddBackendClause is called when exiting the addBackendClause production.
	ExitAddBackendClause(c *AddBackendClauseContext)

	// ExitDropBackendClause is called when exiting the dropBackendClause production.
	ExitDropBackendClause(c *DropBackendClauseContext)

	// ExitDecommissionBackendClause is called when exiting the decommissionBackendClause production.
	ExitDecommissionBackendClause(c *DecommissionBackendClauseContext)

	// ExitModifyBackendClause is called when exiting the modifyBackendClause production.
	ExitModifyBackendClause(c *ModifyBackendClauseContext)

	// ExitAddComputeNodeClause is called when exiting the addComputeNodeClause production.
	ExitAddComputeNodeClause(c *AddComputeNodeClauseContext)

	// ExitDropComputeNodeClause is called when exiting the dropComputeNodeClause production.
	ExitDropComputeNodeClause(c *DropComputeNodeClauseContext)

	// ExitModifyBrokerClause is called when exiting the modifyBrokerClause production.
	ExitModifyBrokerClause(c *ModifyBrokerClauseContext)

	// ExitAlterLoadErrorUrlClause is called when exiting the alterLoadErrorUrlClause production.
	ExitAlterLoadErrorUrlClause(c *AlterLoadErrorUrlClauseContext)

	// ExitCreateImageClause is called when exiting the createImageClause production.
	ExitCreateImageClause(c *CreateImageClauseContext)

	// ExitCleanTabletSchedQClause is called when exiting the cleanTabletSchedQClause production.
	ExitCleanTabletSchedQClause(c *CleanTabletSchedQClauseContext)

	// ExitDecommissionDiskClause is called when exiting the decommissionDiskClause production.
	ExitDecommissionDiskClause(c *DecommissionDiskClauseContext)

	// ExitCancelDecommissionDiskClause is called when exiting the cancelDecommissionDiskClause production.
	ExitCancelDecommissionDiskClause(c *CancelDecommissionDiskClauseContext)

	// ExitDisableDiskClause is called when exiting the disableDiskClause production.
	ExitDisableDiskClause(c *DisableDiskClauseContext)

	// ExitCancelDisableDiskClause is called when exiting the cancelDisableDiskClause production.
	ExitCancelDisableDiskClause(c *CancelDisableDiskClauseContext)

	// ExitCreateIndexClause is called when exiting the createIndexClause production.
	ExitCreateIndexClause(c *CreateIndexClauseContext)

	// ExitDropIndexClause is called when exiting the dropIndexClause production.
	ExitDropIndexClause(c *DropIndexClauseContext)

	// ExitTableRenameClause is called when exiting the tableRenameClause production.
	ExitTableRenameClause(c *TableRenameClauseContext)

	// ExitSwapTableClause is called when exiting the swapTableClause production.
	ExitSwapTableClause(c *SwapTableClauseContext)

	// ExitModifyPropertiesClause is called when exiting the modifyPropertiesClause production.
	ExitModifyPropertiesClause(c *ModifyPropertiesClauseContext)

	// ExitModifyCommentClause is called when exiting the modifyCommentClause production.
	ExitModifyCommentClause(c *ModifyCommentClauseContext)

	// ExitOptimizeRange is called when exiting the optimizeRange production.
	ExitOptimizeRange(c *OptimizeRangeContext)

	// ExitOptimizeClause is called when exiting the optimizeClause production.
	ExitOptimizeClause(c *OptimizeClauseContext)

	// ExitAddColumnClause is called when exiting the addColumnClause production.
	ExitAddColumnClause(c *AddColumnClauseContext)

	// ExitAddColumnsClause is called when exiting the addColumnsClause production.
	ExitAddColumnsClause(c *AddColumnsClauseContext)

	// ExitDropColumnClause is called when exiting the dropColumnClause production.
	ExitDropColumnClause(c *DropColumnClauseContext)

	// ExitModifyColumnClause is called when exiting the modifyColumnClause production.
	ExitModifyColumnClause(c *ModifyColumnClauseContext)

	// ExitModifyColumnCommentClause is called when exiting the modifyColumnCommentClause production.
	ExitModifyColumnCommentClause(c *ModifyColumnCommentClauseContext)

	// ExitColumnRenameClause is called when exiting the columnRenameClause production.
	ExitColumnRenameClause(c *ColumnRenameClauseContext)

	// ExitReorderColumnsClause is called when exiting the reorderColumnsClause production.
	ExitReorderColumnsClause(c *ReorderColumnsClauseContext)

	// ExitRollupRenameClause is called when exiting the rollupRenameClause production.
	ExitRollupRenameClause(c *RollupRenameClauseContext)

	// ExitCompactionClause is called when exiting the compactionClause production.
	ExitCompactionClause(c *CompactionClauseContext)

	// ExitSubfieldName is called when exiting the subfieldName production.
	ExitSubfieldName(c *SubfieldNameContext)

	// ExitNestedFieldName is called when exiting the nestedFieldName production.
	ExitNestedFieldName(c *NestedFieldNameContext)

	// ExitAddFieldClause is called when exiting the addFieldClause production.
	ExitAddFieldClause(c *AddFieldClauseContext)

	// ExitDropFieldClause is called when exiting the dropFieldClause production.
	ExitDropFieldClause(c *DropFieldClauseContext)

	// ExitCreateOrReplaceTagClause is called when exiting the createOrReplaceTagClause production.
	ExitCreateOrReplaceTagClause(c *CreateOrReplaceTagClauseContext)

	// ExitCreateOrReplaceBranchClause is called when exiting the createOrReplaceBranchClause production.
	ExitCreateOrReplaceBranchClause(c *CreateOrReplaceBranchClauseContext)

	// ExitDropBranchClause is called when exiting the dropBranchClause production.
	ExitDropBranchClause(c *DropBranchClauseContext)

	// ExitDropTagClause is called when exiting the dropTagClause production.
	ExitDropTagClause(c *DropTagClauseContext)

	// ExitTableOperationClause is called when exiting the tableOperationClause production.
	ExitTableOperationClause(c *TableOperationClauseContext)

	// ExitTagOptions is called when exiting the tagOptions production.
	ExitTagOptions(c *TagOptionsContext)

	// ExitBranchOptions is called when exiting the branchOptions production.
	ExitBranchOptions(c *BranchOptionsContext)

	// ExitSnapshotRetention is called when exiting the snapshotRetention production.
	ExitSnapshotRetention(c *SnapshotRetentionContext)

	// ExitRefRetain is called when exiting the refRetain production.
	ExitRefRetain(c *RefRetainContext)

	// ExitMaxSnapshotAge is called when exiting the maxSnapshotAge production.
	ExitMaxSnapshotAge(c *MaxSnapshotAgeContext)

	// ExitMinSnapshotsToKeep is called when exiting the minSnapshotsToKeep production.
	ExitMinSnapshotsToKeep(c *MinSnapshotsToKeepContext)

	// ExitSnapshotId is called when exiting the snapshotId production.
	ExitSnapshotId(c *SnapshotIdContext)

	// ExitTimeUnit is called when exiting the timeUnit production.
	ExitTimeUnit(c *TimeUnitContext)

	// ExitInteger_list is called when exiting the integer_list production.
	ExitInteger_list(c *Integer_listContext)

	// ExitDropPersistentIndexClause is called when exiting the dropPersistentIndexClause production.
	ExitDropPersistentIndexClause(c *DropPersistentIndexClauseContext)

	// ExitAddPartitionClause is called when exiting the addPartitionClause production.
	ExitAddPartitionClause(c *AddPartitionClauseContext)

	// ExitDropPartitionClause is called when exiting the dropPartitionClause production.
	ExitDropPartitionClause(c *DropPartitionClauseContext)

	// ExitTruncatePartitionClause is called when exiting the truncatePartitionClause production.
	ExitTruncatePartitionClause(c *TruncatePartitionClauseContext)

	// ExitModifyPartitionClause is called when exiting the modifyPartitionClause production.
	ExitModifyPartitionClause(c *ModifyPartitionClauseContext)

	// ExitReplacePartitionClause is called when exiting the replacePartitionClause production.
	ExitReplacePartitionClause(c *ReplacePartitionClauseContext)

	// ExitPartitionRenameClause is called when exiting the partitionRenameClause production.
	ExitPartitionRenameClause(c *PartitionRenameClauseContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitInsertLabelOrColumnAliases is called when exiting the insertLabelOrColumnAliases production.
	ExitInsertLabelOrColumnAliases(c *InsertLabelOrColumnAliasesContext)

	// ExitColumnAliasesOrByName is called when exiting the columnAliasesOrByName production.
	ExitColumnAliasesOrByName(c *ColumnAliasesOrByNameContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitCreateRoutineLoadStatement is called when exiting the createRoutineLoadStatement production.
	ExitCreateRoutineLoadStatement(c *CreateRoutineLoadStatementContext)

	// ExitAlterRoutineLoadStatement is called when exiting the alterRoutineLoadStatement production.
	ExitAlterRoutineLoadStatement(c *AlterRoutineLoadStatementContext)

	// ExitDataSource is called when exiting the dataSource production.
	ExitDataSource(c *DataSourceContext)

	// ExitLoadProperties is called when exiting the loadProperties production.
	ExitLoadProperties(c *LoadPropertiesContext)

	// ExitColSeparatorProperty is called when exiting the colSeparatorProperty production.
	ExitColSeparatorProperty(c *ColSeparatorPropertyContext)

	// ExitRowDelimiterProperty is called when exiting the rowDelimiterProperty production.
	ExitRowDelimiterProperty(c *RowDelimiterPropertyContext)

	// ExitImportColumns is called when exiting the importColumns production.
	ExitImportColumns(c *ImportColumnsContext)

	// ExitColumnProperties is called when exiting the columnProperties production.
	ExitColumnProperties(c *ColumnPropertiesContext)

	// ExitJobProperties is called when exiting the jobProperties production.
	ExitJobProperties(c *JobPropertiesContext)

	// ExitDataSourceProperties is called when exiting the dataSourceProperties production.
	ExitDataSourceProperties(c *DataSourcePropertiesContext)

	// ExitStopRoutineLoadStatement is called when exiting the stopRoutineLoadStatement production.
	ExitStopRoutineLoadStatement(c *StopRoutineLoadStatementContext)

	// ExitResumeRoutineLoadStatement is called when exiting the resumeRoutineLoadStatement production.
	ExitResumeRoutineLoadStatement(c *ResumeRoutineLoadStatementContext)

	// ExitPauseRoutineLoadStatement is called when exiting the pauseRoutineLoadStatement production.
	ExitPauseRoutineLoadStatement(c *PauseRoutineLoadStatementContext)

	// ExitShowRoutineLoadStatement is called when exiting the showRoutineLoadStatement production.
	ExitShowRoutineLoadStatement(c *ShowRoutineLoadStatementContext)

	// ExitShowRoutineLoadTaskStatement is called when exiting the showRoutineLoadTaskStatement production.
	ExitShowRoutineLoadTaskStatement(c *ShowRoutineLoadTaskStatementContext)

	// ExitShowCreateRoutineLoadStatement is called when exiting the showCreateRoutineLoadStatement production.
	ExitShowCreateRoutineLoadStatement(c *ShowCreateRoutineLoadStatementContext)

	// ExitShowStreamLoadStatement is called when exiting the showStreamLoadStatement production.
	ExitShowStreamLoadStatement(c *ShowStreamLoadStatementContext)

	// ExitAnalyzeStatement is called when exiting the analyzeStatement production.
	ExitAnalyzeStatement(c *AnalyzeStatementContext)

	// ExitRegularColumns is called when exiting the regularColumns production.
	ExitRegularColumns(c *RegularColumnsContext)

	// ExitAllColumns is called when exiting the allColumns production.
	ExitAllColumns(c *AllColumnsContext)

	// ExitPredicateColumns is called when exiting the predicateColumns production.
	ExitPredicateColumns(c *PredicateColumnsContext)

	// ExitMultiColumnSet is called when exiting the multiColumnSet production.
	ExitMultiColumnSet(c *MultiColumnSetContext)

	// ExitDropStatsStatement is called when exiting the dropStatsStatement production.
	ExitDropStatsStatement(c *DropStatsStatementContext)

	// ExitHistogramStatement is called when exiting the histogramStatement production.
	ExitHistogramStatement(c *HistogramStatementContext)

	// ExitAnalyzeHistogramStatement is called when exiting the analyzeHistogramStatement production.
	ExitAnalyzeHistogramStatement(c *AnalyzeHistogramStatementContext)

	// ExitDropHistogramStatement is called when exiting the dropHistogramStatement production.
	ExitDropHistogramStatement(c *DropHistogramStatementContext)

	// ExitCreateAnalyzeStatement is called when exiting the createAnalyzeStatement production.
	ExitCreateAnalyzeStatement(c *CreateAnalyzeStatementContext)

	// ExitDropAnalyzeJobStatement is called when exiting the dropAnalyzeJobStatement production.
	ExitDropAnalyzeJobStatement(c *DropAnalyzeJobStatementContext)

	// ExitShowAnalyzeStatement is called when exiting the showAnalyzeStatement production.
	ExitShowAnalyzeStatement(c *ShowAnalyzeStatementContext)

	// ExitShowStatsMetaStatement is called when exiting the showStatsMetaStatement production.
	ExitShowStatsMetaStatement(c *ShowStatsMetaStatementContext)

	// ExitShowHistogramMetaStatement is called when exiting the showHistogramMetaStatement production.
	ExitShowHistogramMetaStatement(c *ShowHistogramMetaStatementContext)

	// ExitKillAnalyzeStatement is called when exiting the killAnalyzeStatement production.
	ExitKillAnalyzeStatement(c *KillAnalyzeStatementContext)

	// ExitAnalyzeProfileStatement is called when exiting the analyzeProfileStatement production.
	ExitAnalyzeProfileStatement(c *AnalyzeProfileStatementContext)

	// ExitCreateBaselinePlanStatement is called when exiting the createBaselinePlanStatement production.
	ExitCreateBaselinePlanStatement(c *CreateBaselinePlanStatementContext)

	// ExitDropBaselinePlanStatement is called when exiting the dropBaselinePlanStatement production.
	ExitDropBaselinePlanStatement(c *DropBaselinePlanStatementContext)

	// ExitShowBaselinePlanStatement is called when exiting the showBaselinePlanStatement production.
	ExitShowBaselinePlanStatement(c *ShowBaselinePlanStatementContext)

	// ExitCreateResourceGroupStatement is called when exiting the createResourceGroupStatement production.
	ExitCreateResourceGroupStatement(c *CreateResourceGroupStatementContext)

	// ExitDropResourceGroupStatement is called when exiting the dropResourceGroupStatement production.
	ExitDropResourceGroupStatement(c *DropResourceGroupStatementContext)

	// ExitAlterResourceGroupStatement is called when exiting the alterResourceGroupStatement production.
	ExitAlterResourceGroupStatement(c *AlterResourceGroupStatementContext)

	// ExitShowResourceGroupStatement is called when exiting the showResourceGroupStatement production.
	ExitShowResourceGroupStatement(c *ShowResourceGroupStatementContext)

	// ExitShowResourceGroupUsageStatement is called when exiting the showResourceGroupUsageStatement production.
	ExitShowResourceGroupUsageStatement(c *ShowResourceGroupUsageStatementContext)

	// ExitCreateResourceStatement is called when exiting the createResourceStatement production.
	ExitCreateResourceStatement(c *CreateResourceStatementContext)

	// ExitAlterResourceStatement is called when exiting the alterResourceStatement production.
	ExitAlterResourceStatement(c *AlterResourceStatementContext)

	// ExitDropResourceStatement is called when exiting the dropResourceStatement production.
	ExitDropResourceStatement(c *DropResourceStatementContext)

	// ExitShowResourceStatement is called when exiting the showResourceStatement production.
	ExitShowResourceStatement(c *ShowResourceStatementContext)

	// ExitClassifier is called when exiting the classifier production.
	ExitClassifier(c *ClassifierContext)

	// ExitShowFunctionsStatement is called when exiting the showFunctionsStatement production.
	ExitShowFunctionsStatement(c *ShowFunctionsStatementContext)

	// ExitDropFunctionStatement is called when exiting the dropFunctionStatement production.
	ExitDropFunctionStatement(c *DropFunctionStatementContext)

	// ExitCreateFunctionStatement is called when exiting the createFunctionStatement production.
	ExitCreateFunctionStatement(c *CreateFunctionStatementContext)

	// ExitInlineFunction is called when exiting the inlineFunction production.
	ExitInlineFunction(c *InlineFunctionContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitLoadStatement is called when exiting the loadStatement production.
	ExitLoadStatement(c *LoadStatementContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitDataDescList is called when exiting the dataDescList production.
	ExitDataDescList(c *DataDescListContext)

	// ExitDataDesc is called when exiting the dataDesc production.
	ExitDataDesc(c *DataDescContext)

	// ExitFormatProps is called when exiting the formatProps production.
	ExitFormatProps(c *FormatPropsContext)

	// ExitBrokerDesc is called when exiting the brokerDesc production.
	ExitBrokerDesc(c *BrokerDescContext)

	// ExitResourceDesc is called when exiting the resourceDesc production.
	ExitResourceDesc(c *ResourceDescContext)

	// ExitShowLoadStatement is called when exiting the showLoadStatement production.
	ExitShowLoadStatement(c *ShowLoadStatementContext)

	// ExitShowLoadWarningsStatement is called when exiting the showLoadWarningsStatement production.
	ExitShowLoadWarningsStatement(c *ShowLoadWarningsStatementContext)

	// ExitCancelLoadStatement is called when exiting the cancelLoadStatement production.
	ExitCancelLoadStatement(c *CancelLoadStatementContext)

	// ExitAlterLoadStatement is called when exiting the alterLoadStatement production.
	ExitAlterLoadStatement(c *AlterLoadStatementContext)

	// ExitCancelCompactionStatement is called when exiting the cancelCompactionStatement production.
	ExitCancelCompactionStatement(c *CancelCompactionStatementContext)

	// ExitShowAuthorStatement is called when exiting the showAuthorStatement production.
	ExitShowAuthorStatement(c *ShowAuthorStatementContext)

	// ExitShowBackendsStatement is called when exiting the showBackendsStatement production.
	ExitShowBackendsStatement(c *ShowBackendsStatementContext)

	// ExitShowBrokerStatement is called when exiting the showBrokerStatement production.
	ExitShowBrokerStatement(c *ShowBrokerStatementContext)

	// ExitShowCharsetStatement is called when exiting the showCharsetStatement production.
	ExitShowCharsetStatement(c *ShowCharsetStatementContext)

	// ExitShowCollationStatement is called when exiting the showCollationStatement production.
	ExitShowCollationStatement(c *ShowCollationStatementContext)

	// ExitShowDeleteStatement is called when exiting the showDeleteStatement production.
	ExitShowDeleteStatement(c *ShowDeleteStatementContext)

	// ExitShowDynamicPartitionStatement is called when exiting the showDynamicPartitionStatement production.
	ExitShowDynamicPartitionStatement(c *ShowDynamicPartitionStatementContext)

	// ExitShowEventsStatement is called when exiting the showEventsStatement production.
	ExitShowEventsStatement(c *ShowEventsStatementContext)

	// ExitShowEnginesStatement is called when exiting the showEnginesStatement production.
	ExitShowEnginesStatement(c *ShowEnginesStatementContext)

	// ExitShowFrontendsStatement is called when exiting the showFrontendsStatement production.
	ExitShowFrontendsStatement(c *ShowFrontendsStatementContext)

	// ExitShowFrontendsDisksStatement is called when exiting the showFrontendsDisksStatement production.
	ExitShowFrontendsDisksStatement(c *ShowFrontendsDisksStatementContext)

	// ExitShowPluginsStatement is called when exiting the showPluginsStatement production.
	ExitShowPluginsStatement(c *ShowPluginsStatementContext)

	// ExitShowRepositoriesStatement is called when exiting the showRepositoriesStatement production.
	ExitShowRepositoriesStatement(c *ShowRepositoriesStatementContext)

	// ExitShowOpenTableStatement is called when exiting the showOpenTableStatement production.
	ExitShowOpenTableStatement(c *ShowOpenTableStatementContext)

	// ExitShowPrivilegesStatement is called when exiting the showPrivilegesStatement production.
	ExitShowPrivilegesStatement(c *ShowPrivilegesStatementContext)

	// ExitShowProcedureStatement is called when exiting the showProcedureStatement production.
	ExitShowProcedureStatement(c *ShowProcedureStatementContext)

	// ExitShowProcStatement is called when exiting the showProcStatement production.
	ExitShowProcStatement(c *ShowProcStatementContext)

	// ExitShowProcesslistStatement is called when exiting the showProcesslistStatement production.
	ExitShowProcesslistStatement(c *ShowProcesslistStatementContext)

	// ExitShowProfilelistStatement is called when exiting the showProfilelistStatement production.
	ExitShowProfilelistStatement(c *ShowProfilelistStatementContext)

	// ExitShowRunningQueriesStatement is called when exiting the showRunningQueriesStatement production.
	ExitShowRunningQueriesStatement(c *ShowRunningQueriesStatementContext)

	// ExitShowStatusStatement is called when exiting the showStatusStatement production.
	ExitShowStatusStatement(c *ShowStatusStatementContext)

	// ExitShowTabletStatement is called when exiting the showTabletStatement production.
	ExitShowTabletStatement(c *ShowTabletStatementContext)

	// ExitShowTransactionStatement is called when exiting the showTransactionStatement production.
	ExitShowTransactionStatement(c *ShowTransactionStatementContext)

	// ExitShowTriggersStatement is called when exiting the showTriggersStatement production.
	ExitShowTriggersStatement(c *ShowTriggersStatementContext)

	// ExitShowUserPropertyStatement is called when exiting the showUserPropertyStatement production.
	ExitShowUserPropertyStatement(c *ShowUserPropertyStatementContext)

	// ExitShowVariablesStatement is called when exiting the showVariablesStatement production.
	ExitShowVariablesStatement(c *ShowVariablesStatementContext)

	// ExitShowWarningStatement is called when exiting the showWarningStatement production.
	ExitShowWarningStatement(c *ShowWarningStatementContext)

	// ExitHelpStatement is called when exiting the helpStatement production.
	ExitHelpStatement(c *HelpStatementContext)

	// ExitShowQueryProfileStatement is called when exiting the showQueryProfileStatement production.
	ExitShowQueryProfileStatement(c *ShowQueryProfileStatementContext)

	// ExitShowQueryStatsStatement is called when exiting the showQueryStatsStatement production.
	ExitShowQueryStatsStatement(c *ShowQueryStatsStatementContext)

	// ExitShowLoadProfileStatement is called when exiting the showLoadProfileStatement production.
	ExitShowLoadProfileStatement(c *ShowLoadProfileStatementContext)

	// ExitShowDataSkewStatement is called when exiting the showDataSkewStatement production.
	ExitShowDataSkewStatement(c *ShowDataSkewStatementContext)

	// ExitShowDataTypesStatement is called when exiting the showDataTypesStatement production.
	ExitShowDataTypesStatement(c *ShowDataTypesStatementContext)

	// ExitShowSyncJobStatement is called when exiting the showSyncJobStatement production.
	ExitShowSyncJobStatement(c *ShowSyncJobStatementContext)

	// ExitShowPolicyStatement is called when exiting the showPolicyStatement production.
	ExitShowPolicyStatement(c *ShowPolicyStatementContext)

	// ExitShowSqlBlockRuleStatement is called when exiting the showSqlBlockRuleStatement production.
	ExitShowSqlBlockRuleStatement(c *ShowSqlBlockRuleStatementContext)

	// ExitShowEncryptKeysStatement is called when exiting the showEncryptKeysStatement production.
	ExitShowEncryptKeysStatement(c *ShowEncryptKeysStatementContext)

	// ExitShowCreateLoadStatement is called when exiting the showCreateLoadStatement production.
	ExitShowCreateLoadStatement(c *ShowCreateLoadStatementContext)

	// ExitShowCreateRepositoryStatement is called when exiting the showCreateRepositoryStatement production.
	ExitShowCreateRepositoryStatement(c *ShowCreateRepositoryStatementContext)

	// ExitShowLastInsertStatement is called when exiting the showLastInsertStatement production.
	ExitShowLastInsertStatement(c *ShowLastInsertStatementContext)

	// ExitShowTableIdStatement is called when exiting the showTableIdStatement production.
	ExitShowTableIdStatement(c *ShowTableIdStatementContext)

	// ExitShowDatabaseIdStatement is called when exiting the showDatabaseIdStatement production.
	ExitShowDatabaseIdStatement(c *ShowDatabaseIdStatementContext)

	// ExitShowPartitionIdStatement is called when exiting the showPartitionIdStatement production.
	ExitShowPartitionIdStatement(c *ShowPartitionIdStatementContext)

	// ExitShowTableStatsStatement is called when exiting the showTableStatsStatement production.
	ExitShowTableStatsStatement(c *ShowTableStatsStatementContext)

	// ExitShowColumnStatsStatement is called when exiting the showColumnStatsStatement production.
	ExitShowColumnStatsStatement(c *ShowColumnStatsStatementContext)

	// ExitShowConvertLightSchemaChangeStatement is called when exiting the showConvertLightSchemaChangeStatement production.
	ExitShowConvertLightSchemaChangeStatement(c *ShowConvertLightSchemaChangeStatementContext)

	// ExitShowCatalogRecycleBinStatement is called when exiting the showCatalogRecycleBinStatement production.
	ExitShowCatalogRecycleBinStatement(c *ShowCatalogRecycleBinStatementContext)

	// ExitShowTrashStatement is called when exiting the showTrashStatement production.
	ExitShowTrashStatement(c *ShowTrashStatementContext)

	// ExitShowMigrationsStatement is called when exiting the showMigrationsStatement production.
	ExitShowMigrationsStatement(c *ShowMigrationsStatementContext)

	// ExitShowWorkloadGroupsStatement is called when exiting the showWorkloadGroupsStatement production.
	ExitShowWorkloadGroupsStatement(c *ShowWorkloadGroupsStatementContext)

	// ExitShowJobTaskStatement is called when exiting the showJobTaskStatement production.
	ExitShowJobTaskStatement(c *ShowJobTaskStatementContext)

	// ExitCreateUserStatement is called when exiting the createUserStatement production.
	ExitCreateUserStatement(c *CreateUserStatementContext)

	// ExitDropUserStatement is called when exiting the dropUserStatement production.
	ExitDropUserStatement(c *DropUserStatementContext)

	// ExitAlterUserStatement is called when exiting the alterUserStatement production.
	ExitAlterUserStatement(c *AlterUserStatementContext)

	// ExitShowUserStatement is called when exiting the showUserStatement production.
	ExitShowUserStatement(c *ShowUserStatementContext)

	// ExitShowAllAuthentication is called when exiting the showAllAuthentication production.
	ExitShowAllAuthentication(c *ShowAllAuthenticationContext)

	// ExitShowAuthenticationForUser is called when exiting the showAuthenticationForUser production.
	ExitShowAuthenticationForUser(c *ShowAuthenticationForUserContext)

	// ExitExecuteAsStatement is called when exiting the executeAsStatement production.
	ExitExecuteAsStatement(c *ExecuteAsStatementContext)

	// ExitCreateRoleStatement is called when exiting the createRoleStatement production.
	ExitCreateRoleStatement(c *CreateRoleStatementContext)

	// ExitAlterRoleStatement is called when exiting the alterRoleStatement production.
	ExitAlterRoleStatement(c *AlterRoleStatementContext)

	// ExitDropRoleStatement is called when exiting the dropRoleStatement production.
	ExitDropRoleStatement(c *DropRoleStatementContext)

	// ExitShowRolesStatement is called when exiting the showRolesStatement production.
	ExitShowRolesStatement(c *ShowRolesStatementContext)

	// ExitGrantRoleToUser is called when exiting the grantRoleToUser production.
	ExitGrantRoleToUser(c *GrantRoleToUserContext)

	// ExitGrantRoleToRole is called when exiting the grantRoleToRole production.
	ExitGrantRoleToRole(c *GrantRoleToRoleContext)

	// ExitRevokeRoleFromUser is called when exiting the revokeRoleFromUser production.
	ExitRevokeRoleFromUser(c *RevokeRoleFromUserContext)

	// ExitRevokeRoleFromRole is called when exiting the revokeRoleFromRole production.
	ExitRevokeRoleFromRole(c *RevokeRoleFromRoleContext)

	// ExitSetRoleStatement is called when exiting the setRoleStatement production.
	ExitSetRoleStatement(c *SetRoleStatementContext)

	// ExitSetDefaultRoleStatement is called when exiting the setDefaultRoleStatement production.
	ExitSetDefaultRoleStatement(c *SetDefaultRoleStatementContext)

	// ExitGrantRevokeClause is called when exiting the grantRevokeClause production.
	ExitGrantRevokeClause(c *GrantRevokeClauseContext)

	// ExitGrantOnUser is called when exiting the grantOnUser production.
	ExitGrantOnUser(c *GrantOnUserContext)

	// ExitGrantOnTableBrief is called when exiting the grantOnTableBrief production.
	ExitGrantOnTableBrief(c *GrantOnTableBriefContext)

	// ExitGrantOnFunc is called when exiting the grantOnFunc production.
	ExitGrantOnFunc(c *GrantOnFuncContext)

	// ExitGrantOnSystem is called when exiting the grantOnSystem production.
	ExitGrantOnSystem(c *GrantOnSystemContext)

	// ExitGrantOnPrimaryObj is called when exiting the grantOnPrimaryObj production.
	ExitGrantOnPrimaryObj(c *GrantOnPrimaryObjContext)

	// ExitGrantOnAll is called when exiting the grantOnAll production.
	ExitGrantOnAll(c *GrantOnAllContext)

	// ExitRevokeOnUser is called when exiting the revokeOnUser production.
	ExitRevokeOnUser(c *RevokeOnUserContext)

	// ExitRevokeOnTableBrief is called when exiting the revokeOnTableBrief production.
	ExitRevokeOnTableBrief(c *RevokeOnTableBriefContext)

	// ExitRevokeOnFunc is called when exiting the revokeOnFunc production.
	ExitRevokeOnFunc(c *RevokeOnFuncContext)

	// ExitRevokeOnSystem is called when exiting the revokeOnSystem production.
	ExitRevokeOnSystem(c *RevokeOnSystemContext)

	// ExitRevokeOnPrimaryObj is called when exiting the revokeOnPrimaryObj production.
	ExitRevokeOnPrimaryObj(c *RevokeOnPrimaryObjContext)

	// ExitRevokeOnAll is called when exiting the revokeOnAll production.
	ExitRevokeOnAll(c *RevokeOnAllContext)

	// ExitShowGrantsStatement is called when exiting the showGrantsStatement production.
	ExitShowGrantsStatement(c *ShowGrantsStatementContext)

	// ExitAuthWithoutPlugin is called when exiting the authWithoutPlugin production.
	ExitAuthWithoutPlugin(c *AuthWithoutPluginContext)

	// ExitAuthWithPlugin is called when exiting the authWithPlugin production.
	ExitAuthWithPlugin(c *AuthWithPluginContext)

	// ExitPrivObjectName is called when exiting the privObjectName production.
	ExitPrivObjectName(c *PrivObjectNameContext)

	// ExitPrivObjectNameList is called when exiting the privObjectNameList production.
	ExitPrivObjectNameList(c *PrivObjectNameListContext)

	// ExitPrivFunctionObjectNameList is called when exiting the privFunctionObjectNameList production.
	ExitPrivFunctionObjectNameList(c *PrivFunctionObjectNameListContext)

	// ExitPrivilegeTypeList is called when exiting the privilegeTypeList production.
	ExitPrivilegeTypeList(c *PrivilegeTypeListContext)

	// ExitPrivilegeType is called when exiting the privilegeType production.
	ExitPrivilegeType(c *PrivilegeTypeContext)

	// ExitPrivObjectType is called when exiting the privObjectType production.
	ExitPrivObjectType(c *PrivObjectTypeContext)

	// ExitPrivObjectTypePlural is called when exiting the privObjectTypePlural production.
	ExitPrivObjectTypePlural(c *PrivObjectTypePluralContext)

	// ExitCreateSecurityIntegrationStatement is called when exiting the createSecurityIntegrationStatement production.
	ExitCreateSecurityIntegrationStatement(c *CreateSecurityIntegrationStatementContext)

	// ExitAlterSecurityIntegrationStatement is called when exiting the alterSecurityIntegrationStatement production.
	ExitAlterSecurityIntegrationStatement(c *AlterSecurityIntegrationStatementContext)

	// ExitDropSecurityIntegrationStatement is called when exiting the dropSecurityIntegrationStatement production.
	ExitDropSecurityIntegrationStatement(c *DropSecurityIntegrationStatementContext)

	// ExitShowSecurityIntegrationStatement is called when exiting the showSecurityIntegrationStatement production.
	ExitShowSecurityIntegrationStatement(c *ShowSecurityIntegrationStatementContext)

	// ExitShowCreateSecurityIntegrationStatement is called when exiting the showCreateSecurityIntegrationStatement production.
	ExitShowCreateSecurityIntegrationStatement(c *ShowCreateSecurityIntegrationStatementContext)

	// ExitCreateGroupProviderStatement is called when exiting the createGroupProviderStatement production.
	ExitCreateGroupProviderStatement(c *CreateGroupProviderStatementContext)

	// ExitDropGroupProviderStatement is called when exiting the dropGroupProviderStatement production.
	ExitDropGroupProviderStatement(c *DropGroupProviderStatementContext)

	// ExitShowGroupProvidersStatement is called when exiting the showGroupProvidersStatement production.
	ExitShowGroupProvidersStatement(c *ShowGroupProvidersStatementContext)

	// ExitShowCreateGroupProviderStatement is called when exiting the showCreateGroupProviderStatement production.
	ExitShowCreateGroupProviderStatement(c *ShowCreateGroupProviderStatementContext)

	// ExitBackupStatement is called when exiting the backupStatement production.
	ExitBackupStatement(c *BackupStatementContext)

	// ExitCancelBackupStatement is called when exiting the cancelBackupStatement production.
	ExitCancelBackupStatement(c *CancelBackupStatementContext)

	// ExitShowBackupStatement is called when exiting the showBackupStatement production.
	ExitShowBackupStatement(c *ShowBackupStatementContext)

	// ExitRestoreStatement is called when exiting the restoreStatement production.
	ExitRestoreStatement(c *RestoreStatementContext)

	// ExitCancelRestoreStatement is called when exiting the cancelRestoreStatement production.
	ExitCancelRestoreStatement(c *CancelRestoreStatementContext)

	// ExitShowRestoreStatement is called when exiting the showRestoreStatement production.
	ExitShowRestoreStatement(c *ShowRestoreStatementContext)

	// ExitShowSnapshotStatement is called when exiting the showSnapshotStatement production.
	ExitShowSnapshotStatement(c *ShowSnapshotStatementContext)

	// ExitCreateRepositoryStatement is called when exiting the createRepositoryStatement production.
	ExitCreateRepositoryStatement(c *CreateRepositoryStatementContext)

	// ExitDropRepositoryStatement is called when exiting the dropRepositoryStatement production.
	ExitDropRepositoryStatement(c *DropRepositoryStatementContext)

	// ExitAddSqlBlackListStatement is called when exiting the addSqlBlackListStatement production.
	ExitAddSqlBlackListStatement(c *AddSqlBlackListStatementContext)

	// ExitDelSqlBlackListStatement is called when exiting the delSqlBlackListStatement production.
	ExitDelSqlBlackListStatement(c *DelSqlBlackListStatementContext)

	// ExitShowSqlBlackListStatement is called when exiting the showSqlBlackListStatement production.
	ExitShowSqlBlackListStatement(c *ShowSqlBlackListStatementContext)

	// ExitShowWhiteListStatement is called when exiting the showWhiteListStatement production.
	ExitShowWhiteListStatement(c *ShowWhiteListStatementContext)

	// ExitAddBackendBlackListStatement is called when exiting the addBackendBlackListStatement production.
	ExitAddBackendBlackListStatement(c *AddBackendBlackListStatementContext)

	// ExitDelBackendBlackListStatement is called when exiting the delBackendBlackListStatement production.
	ExitDelBackendBlackListStatement(c *DelBackendBlackListStatementContext)

	// ExitShowBackendBlackListStatement is called when exiting the showBackendBlackListStatement production.
	ExitShowBackendBlackListStatement(c *ShowBackendBlackListStatementContext)

	// ExitDataCacheTarget is called when exiting the dataCacheTarget production.
	ExitDataCacheTarget(c *DataCacheTargetContext)

	// ExitCreateDataCacheRuleStatement is called when exiting the createDataCacheRuleStatement production.
	ExitCreateDataCacheRuleStatement(c *CreateDataCacheRuleStatementContext)

	// ExitShowDataCacheRulesStatement is called when exiting the showDataCacheRulesStatement production.
	ExitShowDataCacheRulesStatement(c *ShowDataCacheRulesStatementContext)

	// ExitDropDataCacheRuleStatement is called when exiting the dropDataCacheRuleStatement production.
	ExitDropDataCacheRuleStatement(c *DropDataCacheRuleStatementContext)

	// ExitClearDataCacheRulesStatement is called when exiting the clearDataCacheRulesStatement production.
	ExitClearDataCacheRulesStatement(c *ClearDataCacheRulesStatementContext)

	// ExitDataCacheSelectStatement is called when exiting the dataCacheSelectStatement production.
	ExitDataCacheSelectStatement(c *DataCacheSelectStatementContext)

	// ExitExportStatement is called when exiting the exportStatement production.
	ExitExportStatement(c *ExportStatementContext)

	// ExitCancelExportStatement is called when exiting the cancelExportStatement production.
	ExitCancelExportStatement(c *CancelExportStatementContext)

	// ExitShowExportStatement is called when exiting the showExportStatement production.
	ExitShowExportStatement(c *ShowExportStatementContext)

	// ExitInstallPluginStatement is called when exiting the installPluginStatement production.
	ExitInstallPluginStatement(c *InstallPluginStatementContext)

	// ExitUninstallPluginStatement is called when exiting the uninstallPluginStatement production.
	ExitUninstallPluginStatement(c *UninstallPluginStatementContext)

	// ExitCreateFileStatement is called when exiting the createFileStatement production.
	ExitCreateFileStatement(c *CreateFileStatementContext)

	// ExitDropFileStatement is called when exiting the dropFileStatement production.
	ExitDropFileStatement(c *DropFileStatementContext)

	// ExitShowSmallFilesStatement is called when exiting the showSmallFilesStatement production.
	ExitShowSmallFilesStatement(c *ShowSmallFilesStatementContext)

	// ExitCreatePipeStatement is called when exiting the createPipeStatement production.
	ExitCreatePipeStatement(c *CreatePipeStatementContext)

	// ExitDropPipeStatement is called when exiting the dropPipeStatement production.
	ExitDropPipeStatement(c *DropPipeStatementContext)

	// ExitAlterPipeClause is called when exiting the alterPipeClause production.
	ExitAlterPipeClause(c *AlterPipeClauseContext)

	// ExitAlterPipeStatement is called when exiting the alterPipeStatement production.
	ExitAlterPipeStatement(c *AlterPipeStatementContext)

	// ExitDescPipeStatement is called when exiting the descPipeStatement production.
	ExitDescPipeStatement(c *DescPipeStatementContext)

	// ExitShowPipeStatement is called when exiting the showPipeStatement production.
	ExitShowPipeStatement(c *ShowPipeStatementContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitSetNames is called when exiting the setNames production.
	ExitSetNames(c *SetNamesContext)

	// ExitSetPassword is called when exiting the setPassword production.
	ExitSetPassword(c *SetPasswordContext)

	// ExitSetUserVar is called when exiting the setUserVar production.
	ExitSetUserVar(c *SetUserVarContext)

	// ExitSetSystemVar is called when exiting the setSystemVar production.
	ExitSetSystemVar(c *SetSystemVarContext)

	// ExitSetTransaction is called when exiting the setTransaction production.
	ExitSetTransaction(c *SetTransactionContext)

	// ExitTransaction_characteristics is called when exiting the transaction_characteristics production.
	ExitTransaction_characteristics(c *Transaction_characteristicsContext)

	// ExitTransaction_access_mode is called when exiting the transaction_access_mode production.
	ExitTransaction_access_mode(c *Transaction_access_modeContext)

	// ExitIsolation_level is called when exiting the isolation_level production.
	ExitIsolation_level(c *Isolation_levelContext)

	// ExitIsolation_types is called when exiting the isolation_types production.
	ExitIsolation_types(c *Isolation_typesContext)

	// ExitSetExprOrDefault is called when exiting the setExprOrDefault production.
	ExitSetExprOrDefault(c *SetExprOrDefaultContext)

	// ExitSetUserPropertyStatement is called when exiting the setUserPropertyStatement production.
	ExitSetUserPropertyStatement(c *SetUserPropertyStatementContext)

	// ExitRoleList is called when exiting the roleList production.
	ExitRoleList(c *RoleListContext)

	// ExitExecuteScriptStatement is called when exiting the executeScriptStatement production.
	ExitExecuteScriptStatement(c *ExecuteScriptStatementContext)

	// ExitUnsupportedStatement is called when exiting the unsupportedStatement production.
	ExitUnsupportedStatement(c *UnsupportedStatementContext)

	// ExitLock_item is called when exiting the lock_item production.
	ExitLock_item(c *Lock_itemContext)

	// ExitLock_type is called when exiting the lock_type production.
	ExitLock_type(c *Lock_typeContext)

	// ExitAlterPlanAdvisorAddStatement is called when exiting the alterPlanAdvisorAddStatement production.
	ExitAlterPlanAdvisorAddStatement(c *AlterPlanAdvisorAddStatementContext)

	// ExitTruncatePlanAdvisorStatement is called when exiting the truncatePlanAdvisorStatement production.
	ExitTruncatePlanAdvisorStatement(c *TruncatePlanAdvisorStatementContext)

	// ExitAlterPlanAdvisorDropStatement is called when exiting the alterPlanAdvisorDropStatement production.
	ExitAlterPlanAdvisorDropStatement(c *AlterPlanAdvisorDropStatementContext)

	// ExitShowPlanAdvisorStatement is called when exiting the showPlanAdvisorStatement production.
	ExitShowPlanAdvisorStatement(c *ShowPlanAdvisorStatementContext)

	// ExitCreateWarehouseStatement is called when exiting the createWarehouseStatement production.
	ExitCreateWarehouseStatement(c *CreateWarehouseStatementContext)

	// ExitDropWarehouseStatement is called when exiting the dropWarehouseStatement production.
	ExitDropWarehouseStatement(c *DropWarehouseStatementContext)

	// ExitSuspendWarehouseStatement is called when exiting the suspendWarehouseStatement production.
	ExitSuspendWarehouseStatement(c *SuspendWarehouseStatementContext)

	// ExitResumeWarehouseStatement is called when exiting the resumeWarehouseStatement production.
	ExitResumeWarehouseStatement(c *ResumeWarehouseStatementContext)

	// ExitSetWarehouseStatement is called when exiting the setWarehouseStatement production.
	ExitSetWarehouseStatement(c *SetWarehouseStatementContext)

	// ExitShowWarehousesStatement is called when exiting the showWarehousesStatement production.
	ExitShowWarehousesStatement(c *ShowWarehousesStatementContext)

	// ExitShowClustersStatement is called when exiting the showClustersStatement production.
	ExitShowClustersStatement(c *ShowClustersStatementContext)

	// ExitShowNodesStatement is called when exiting the showNodesStatement production.
	ExitShowNodesStatement(c *ShowNodesStatementContext)

	// ExitAlterWarehouseStatement is called when exiting the alterWarehouseStatement production.
	ExitAlterWarehouseStatement(c *AlterWarehouseStatementContext)

	// ExitCreateCNGroupStatement is called when exiting the createCNGroupStatement production.
	ExitCreateCNGroupStatement(c *CreateCNGroupStatementContext)

	// ExitDropCNGroupStatement is called when exiting the dropCNGroupStatement production.
	ExitDropCNGroupStatement(c *DropCNGroupStatementContext)

	// ExitEnableCNGroupStatement is called when exiting the enableCNGroupStatement production.
	ExitEnableCNGroupStatement(c *EnableCNGroupStatementContext)

	// ExitDisableCNGroupStatement is called when exiting the disableCNGroupStatement production.
	ExitDisableCNGroupStatement(c *DisableCNGroupStatementContext)

	// ExitAlterCNGroupStatement is called when exiting the alterCNGroupStatement production.
	ExitAlterCNGroupStatement(c *AlterCNGroupStatementContext)

	// ExitBeginStatement is called when exiting the beginStatement production.
	ExitBeginStatement(c *BeginStatementContext)

	// ExitCommitStatement is called when exiting the commitStatement production.
	ExitCommitStatement(c *CommitStatementContext)

	// ExitRollbackStatement is called when exiting the rollbackStatement production.
	ExitRollbackStatement(c *RollbackStatementContext)

	// ExitTranslateStatement is called when exiting the translateStatement production.
	ExitTranslateStatement(c *TranslateStatementContext)

	// ExitDialect is called when exiting the dialect production.
	ExitDialect(c *DialectContext)

	// ExitTranslateSQL is called when exiting the translateSQL production.
	ExitTranslateSQL(c *TranslateSQLContext)

	// ExitQueryStatement is called when exiting the queryStatement production.
	ExitQueryStatement(c *QueryStatementContext)

	// ExitQueryRelation is called when exiting the queryRelation production.
	ExitQueryRelation(c *QueryRelationContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitQueryNoWith is called when exiting the queryNoWith production.
	ExitQueryNoWith(c *QueryNoWithContext)

	// ExitQueryPeriod is called when exiting the queryPeriod production.
	ExitQueryPeriod(c *QueryPeriodContext)

	// ExitPeriodType is called when exiting the periodType production.
	ExitPeriodType(c *PeriodTypeContext)

	// ExitQueryWithParentheses is called when exiting the queryWithParentheses production.
	ExitQueryWithParentheses(c *QueryWithParenthesesContext)

	// ExitSetOperation is called when exiting the setOperation production.
	ExitSetOperation(c *SetOperationContext)

	// ExitQueryPrimaryDefault is called when exiting the queryPrimaryDefault production.
	ExitQueryPrimaryDefault(c *QueryPrimaryDefaultContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitRowConstructor is called when exiting the rowConstructor production.
	ExitRowConstructor(c *RowConstructorContext)

	// ExitSortItem is called when exiting the sortItem production.
	ExitSortItem(c *SortItemContext)

	// ExitLimitConstExpr is called when exiting the limitConstExpr production.
	ExitLimitConstExpr(c *LimitConstExprContext)

	// ExitLimitElement is called when exiting the limitElement production.
	ExitLimitElement(c *LimitElementContext)

	// ExitQuerySpecification is called when exiting the querySpecification production.
	ExitQuerySpecification(c *QuerySpecificationContext)

	// ExitFrom is called when exiting the from production.
	ExitFrom(c *FromContext)

	// ExitDual is called when exiting the dual production.
	ExitDual(c *DualContext)

	// ExitRollup is called when exiting the rollup production.
	ExitRollup(c *RollupContext)

	// ExitCube is called when exiting the cube production.
	ExitCube(c *CubeContext)

	// ExitMultipleGroupingSets is called when exiting the multipleGroupingSets production.
	ExitMultipleGroupingSets(c *MultipleGroupingSetsContext)

	// ExitSingleGroupingSet is called when exiting the singleGroupingSet production.
	ExitSingleGroupingSet(c *SingleGroupingSetContext)

	// ExitGroupingSet is called when exiting the groupingSet production.
	ExitGroupingSet(c *GroupingSetContext)

	// ExitCommonTableExpression is called when exiting the commonTableExpression production.
	ExitCommonTableExpression(c *CommonTableExpressionContext)

	// ExitSetQuantifier is called when exiting the setQuantifier production.
	ExitSetQuantifier(c *SetQuantifierContext)

	// ExitSelectSingle is called when exiting the selectSingle production.
	ExitSelectSingle(c *SelectSingleContext)

	// ExitSelectAll is called when exiting the selectAll production.
	ExitSelectAll(c *SelectAllContext)

	// ExitExcludeClause is called when exiting the excludeClause production.
	ExitExcludeClause(c *ExcludeClauseContext)

	// ExitRelations is called when exiting the relations production.
	ExitRelations(c *RelationsContext)

	// ExitRelationLateralView is called when exiting the relationLateralView production.
	ExitRelationLateralView(c *RelationLateralViewContext)

	// ExitLateralView is called when exiting the lateralView production.
	ExitLateralView(c *LateralViewContext)

	// ExitGeneratorFunction is called when exiting the generatorFunction production.
	ExitGeneratorFunction(c *GeneratorFunctionContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitTableAtom is called when exiting the tableAtom production.
	ExitTableAtom(c *TableAtomContext)

	// ExitInlineTable is called when exiting the inlineTable production.
	ExitInlineTable(c *InlineTableContext)

	// ExitSubqueryWithAlias is called when exiting the subqueryWithAlias production.
	ExitSubqueryWithAlias(c *SubqueryWithAliasContext)

	// ExitTableFunction is called when exiting the tableFunction production.
	ExitTableFunction(c *TableFunctionContext)

	// ExitNormalizedTableFunction is called when exiting the normalizedTableFunction production.
	ExitNormalizedTableFunction(c *NormalizedTableFunctionContext)

	// ExitFileTableFunction is called when exiting the fileTableFunction production.
	ExitFileTableFunction(c *FileTableFunctionContext)

	// ExitParenthesizedRelation is called when exiting the parenthesizedRelation production.
	ExitParenthesizedRelation(c *ParenthesizedRelationContext)

	// ExitPivotClause is called when exiting the pivotClause production.
	ExitPivotClause(c *PivotClauseContext)

	// ExitPivotAggregationExpression is called when exiting the pivotAggregationExpression production.
	ExitPivotAggregationExpression(c *PivotAggregationExpressionContext)

	// ExitPivotValue is called when exiting the pivotValue production.
	ExitPivotValue(c *PivotValueContext)

	// ExitSampleClause is called when exiting the sampleClause production.
	ExitSampleClause(c *SampleClauseContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitNamedArgumentList is called when exiting the namedArgumentList production.
	ExitNamedArgumentList(c *NamedArgumentListContext)

	// ExitNamedArguments is called when exiting the namedArguments production.
	ExitNamedArguments(c *NamedArgumentsContext)

	// ExitJoinRelation is called when exiting the joinRelation production.
	ExitJoinRelation(c *JoinRelationContext)

	// ExitCrossOrInnerJoinType is called when exiting the crossOrInnerJoinType production.
	ExitCrossOrInnerJoinType(c *CrossOrInnerJoinTypeContext)

	// ExitOuterAndSemiJoinType is called when exiting the outerAndSemiJoinType production.
	ExitOuterAndSemiJoinType(c *OuterAndSemiJoinTypeContext)

	// ExitBracketHint is called when exiting the bracketHint production.
	ExitBracketHint(c *BracketHintContext)

	// ExitHintMap is called when exiting the hintMap production.
	ExitHintMap(c *HintMapContext)

	// ExitJoinCriteria is called when exiting the joinCriteria production.
	ExitJoinCriteria(c *JoinCriteriaContext)

	// ExitColumnAliases is called when exiting the columnAliases production.
	ExitColumnAliases(c *ColumnAliasesContext)

	// ExitColumnAliasesWithoutParentheses is called when exiting the columnAliasesWithoutParentheses production.
	ExitColumnAliasesWithoutParentheses(c *ColumnAliasesWithoutParenthesesContext)

	// ExitPartitionNames is called when exiting the partitionNames production.
	ExitPartitionNames(c *PartitionNamesContext)

	// ExitKeyPartitionList is called when exiting the keyPartitionList production.
	ExitKeyPartitionList(c *KeyPartitionListContext)

	// ExitTabletList is called when exiting the tabletList production.
	ExitTabletList(c *TabletListContext)

	// ExitPrepareStatement is called when exiting the prepareStatement production.
	ExitPrepareStatement(c *PrepareStatementContext)

	// ExitPrepareSql is called when exiting the prepareSql production.
	ExitPrepareSql(c *PrepareSqlContext)

	// ExitExecuteStatement is called when exiting the executeStatement production.
	ExitExecuteStatement(c *ExecuteStatementContext)

	// ExitDeallocateStatement is called when exiting the deallocateStatement production.
	ExitDeallocateStatement(c *DeallocateStatementContext)

	// ExitReplicaList is called when exiting the replicaList production.
	ExitReplicaList(c *ReplicaListContext)

	// ExitExpressionsWithDefault is called when exiting the expressionsWithDefault production.
	ExitExpressionsWithDefault(c *ExpressionsWithDefaultContext)

	// ExitExpressionOrDefault is called when exiting the expressionOrDefault production.
	ExitExpressionOrDefault(c *ExpressionOrDefaultContext)

	// ExitMapExpressionList is called when exiting the mapExpressionList production.
	ExitMapExpressionList(c *MapExpressionListContext)

	// ExitMapExpression is called when exiting the mapExpression production.
	ExitMapExpression(c *MapExpressionContext)

	// ExitExpressionSingleton is called when exiting the expressionSingleton production.
	ExitExpressionSingleton(c *ExpressionSingletonContext)

	// ExitExpressionDefault is called when exiting the expressionDefault production.
	ExitExpressionDefault(c *ExpressionDefaultContext)

	// ExitLogicalNot is called when exiting the logicalNot production.
	ExitLogicalNot(c *LogicalNotContext)

	// ExitLogicalBinary is called when exiting the logicalBinary production.
	ExitLogicalBinary(c *LogicalBinaryContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitBooleanExpressionDefault is called when exiting the booleanExpressionDefault production.
	ExitBooleanExpressionDefault(c *BooleanExpressionDefaultContext)

	// ExitIsNull is called when exiting the isNull production.
	ExitIsNull(c *IsNullContext)

	// ExitScalarSubquery is called when exiting the scalarSubquery production.
	ExitScalarSubquery(c *ScalarSubqueryContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitTupleInSubquery is called when exiting the tupleInSubquery production.
	ExitTupleInSubquery(c *TupleInSubqueryContext)

	// ExitInSubquery is called when exiting the inSubquery production.
	ExitInSubquery(c *InSubqueryContext)

	// ExitInList is called when exiting the inList production.
	ExitInList(c *InListContext)

	// ExitBetween is called when exiting the between production.
	ExitBetween(c *BetweenContext)

	// ExitLike is called when exiting the like production.
	ExitLike(c *LikeContext)

	// ExitValueExpressionDefault is called when exiting the valueExpressionDefault production.
	ExitValueExpressionDefault(c *ValueExpressionDefaultContext)

	// ExitArithmeticBinary is called when exiting the arithmeticBinary production.
	ExitArithmeticBinary(c *ArithmeticBinaryContext)

	// ExitDereference is called when exiting the dereference production.
	ExitDereference(c *DereferenceContext)

	// ExitOdbcFunctionCallExpression is called when exiting the odbcFunctionCallExpression production.
	ExitOdbcFunctionCallExpression(c *OdbcFunctionCallExpressionContext)

	// ExitMatchExpr is called when exiting the matchExpr production.
	ExitMatchExpr(c *MatchExprContext)

	// ExitColumnRef is called when exiting the columnRef production.
	ExitColumnRef(c *ColumnRefContext)

	// ExitConvert is called when exiting the convert production.
	ExitConvert(c *ConvertContext)

	// ExitCollectionSubscript is called when exiting the collectionSubscript production.
	ExitCollectionSubscript(c *CollectionSubscriptContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitUserVariableExpression is called when exiting the userVariableExpression production.
	ExitUserVariableExpression(c *UserVariableExpressionContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitSimpleCase is called when exiting the simpleCase production.
	ExitSimpleCase(c *SimpleCaseContext)

	// ExitArrowExpression is called when exiting the arrowExpression production.
	ExitArrowExpression(c *ArrowExpressionContext)

	// ExitArrayExpr is called when exiting the arrayExpr production.
	ExitArrayExpr(c *ArrayExprContext)

	// ExitSystemVariableExpression is called when exiting the systemVariableExpression production.
	ExitSystemVariableExpression(c *SystemVariableExpressionContext)

	// ExitConcat is called when exiting the concat production.
	ExitConcat(c *ConcatContext)

	// ExitSubqueryExpression is called when exiting the subqueryExpression production.
	ExitSubqueryExpression(c *SubqueryExpressionContext)

	// ExitLambdaFunctionExpr is called when exiting the lambdaFunctionExpr production.
	ExitLambdaFunctionExpr(c *LambdaFunctionExprContext)

	// ExitDictionaryGetExpr is called when exiting the dictionaryGetExpr production.
	ExitDictionaryGetExpr(c *DictionaryGetExprContext)

	// ExitCollate is called when exiting the collate production.
	ExitCollate(c *CollateContext)

	// ExitArrayConstructor is called when exiting the arrayConstructor production.
	ExitArrayConstructor(c *ArrayConstructorContext)

	// ExitMapConstructor is called when exiting the mapConstructor production.
	ExitMapConstructor(c *MapConstructorContext)

	// ExitArraySlice is called when exiting the arraySlice production.
	ExitArraySlice(c *ArraySliceContext)

	// ExitExists is called when exiting the exists production.
	ExitExists(c *ExistsContext)

	// ExitSearchedCase is called when exiting the searchedCase production.
	ExitSearchedCase(c *SearchedCaseContext)

	// ExitArithmeticUnary is called when exiting the arithmeticUnary production.
	ExitArithmeticUnary(c *ArithmeticUnaryContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitDateLiteral is called when exiting the dateLiteral production.
	ExitDateLiteral(c *DateLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIntervalLiteral is called when exiting the intervalLiteral production.
	ExitIntervalLiteral(c *IntervalLiteralContext)

	// ExitUnitBoundaryLiteral is called when exiting the unitBoundaryLiteral production.
	ExitUnitBoundaryLiteral(c *UnitBoundaryLiteralContext)

	// ExitBinaryLiteral is called when exiting the binaryLiteral production.
	ExitBinaryLiteral(c *BinaryLiteralContext)

	// ExitParameter is called when exiting the Parameter production.
	ExitParameter(c *ParameterContext)

	// ExitExtract is called when exiting the extract production.
	ExitExtract(c *ExtractContext)

	// ExitGroupingOperation is called when exiting the groupingOperation production.
	ExitGroupingOperation(c *GroupingOperationContext)

	// ExitInformationFunction is called when exiting the informationFunction production.
	ExitInformationFunction(c *InformationFunctionContext)

	// ExitSpecialDateTime is called when exiting the specialDateTime production.
	ExitSpecialDateTime(c *SpecialDateTimeContext)

	// ExitSpecialFunction is called when exiting the specialFunction production.
	ExitSpecialFunction(c *SpecialFunctionContext)

	// ExitAggregationFunctionCall is called when exiting the aggregationFunctionCall production.
	ExitAggregationFunctionCall(c *AggregationFunctionCallContext)

	// ExitWindowFunctionCall is called when exiting the windowFunctionCall production.
	ExitWindowFunctionCall(c *WindowFunctionCallContext)

	// ExitTranslateFunctionCall is called when exiting the translateFunctionCall production.
	ExitTranslateFunctionCall(c *TranslateFunctionCallContext)

	// ExitSimpleFunctionCall is called when exiting the simpleFunctionCall production.
	ExitSimpleFunctionCall(c *SimpleFunctionCallContext)

	// ExitAggregationFunction is called when exiting the aggregationFunction production.
	ExitAggregationFunction(c *AggregationFunctionContext)

	// ExitUserVariable is called when exiting the userVariable production.
	ExitUserVariable(c *UserVariableContext)

	// ExitSystemVariable is called when exiting the systemVariable production.
	ExitSystemVariable(c *SystemVariableContext)

	// ExitColumnReference is called when exiting the columnReference production.
	ExitColumnReference(c *ColumnReferenceContext)

	// ExitInformationFunctionExpression is called when exiting the informationFunctionExpression production.
	ExitInformationFunctionExpression(c *InformationFunctionExpressionContext)

	// ExitSpecialDateTimeExpression is called when exiting the specialDateTimeExpression production.
	ExitSpecialDateTimeExpression(c *SpecialDateTimeExpressionContext)

	// ExitSpecialFunctionExpression is called when exiting the specialFunctionExpression production.
	ExitSpecialFunctionExpression(c *SpecialFunctionExpressionContext)

	// ExitWindowFunction is called when exiting the windowFunction production.
	ExitWindowFunction(c *WindowFunctionContext)

	// ExitWhenClause is called when exiting the whenClause production.
	ExitWhenClause(c *WhenClauseContext)

	// ExitOver is called when exiting the over production.
	ExitOver(c *OverContext)

	// ExitIgnoreNulls is called when exiting the ignoreNulls production.
	ExitIgnoreNulls(c *IgnoreNullsContext)

	// ExitWindowFrame is called when exiting the windowFrame production.
	ExitWindowFrame(c *WindowFrameContext)

	// ExitUnboundedFrame is called when exiting the unboundedFrame production.
	ExitUnboundedFrame(c *UnboundedFrameContext)

	// ExitCurrentRowBound is called when exiting the currentRowBound production.
	ExitCurrentRowBound(c *CurrentRowBoundContext)

	// ExitBoundedFrame is called when exiting the boundedFrame production.
	ExitBoundedFrame(c *BoundedFrameContext)

	// ExitBackupRestoreObjectDesc is called when exiting the backupRestoreObjectDesc production.
	ExitBackupRestoreObjectDesc(c *BackupRestoreObjectDescContext)

	// ExitTableDesc is called when exiting the tableDesc production.
	ExitTableDesc(c *TableDescContext)

	// ExitBackupRestoreTableDesc is called when exiting the backupRestoreTableDesc production.
	ExitBackupRestoreTableDesc(c *BackupRestoreTableDescContext)

	// ExitExplainDesc is called when exiting the explainDesc production.
	ExitExplainDesc(c *ExplainDescContext)

	// ExitOptimizerTrace is called when exiting the optimizerTrace production.
	ExitOptimizerTrace(c *OptimizerTraceContext)

	// ExitPartitionExpr is called when exiting the partitionExpr production.
	ExitPartitionExpr(c *PartitionExprContext)

	// ExitPartitionDesc is called when exiting the partitionDesc production.
	ExitPartitionDesc(c *PartitionDescContext)

	// ExitListPartitionDesc is called when exiting the listPartitionDesc production.
	ExitListPartitionDesc(c *ListPartitionDescContext)

	// ExitSingleItemListPartitionDesc is called when exiting the singleItemListPartitionDesc production.
	ExitSingleItemListPartitionDesc(c *SingleItemListPartitionDescContext)

	// ExitMultiItemListPartitionDesc is called when exiting the multiItemListPartitionDesc production.
	ExitMultiItemListPartitionDesc(c *MultiItemListPartitionDescContext)

	// ExitMultiListPartitionValues is called when exiting the multiListPartitionValues production.
	ExitMultiListPartitionValues(c *MultiListPartitionValuesContext)

	// ExitSingleListPartitionValues is called when exiting the singleListPartitionValues production.
	ExitSingleListPartitionValues(c *SingleListPartitionValuesContext)

	// ExitListPartitionValues is called when exiting the listPartitionValues production.
	ExitListPartitionValues(c *ListPartitionValuesContext)

	// ExitListPartitionValue is called when exiting the listPartitionValue production.
	ExitListPartitionValue(c *ListPartitionValueContext)

	// ExitStringList is called when exiting the stringList production.
	ExitStringList(c *StringListContext)

	// ExitLiteralExpressionList is called when exiting the literalExpressionList production.
	ExitLiteralExpressionList(c *LiteralExpressionListContext)

	// ExitRangePartitionDesc is called when exiting the rangePartitionDesc production.
	ExitRangePartitionDesc(c *RangePartitionDescContext)

	// ExitSingleRangePartition is called when exiting the singleRangePartition production.
	ExitSingleRangePartition(c *SingleRangePartitionContext)

	// ExitMultiRangePartition is called when exiting the multiRangePartition production.
	ExitMultiRangePartition(c *MultiRangePartitionContext)

	// ExitPartitionRangeDesc is called when exiting the partitionRangeDesc production.
	ExitPartitionRangeDesc(c *PartitionRangeDescContext)

	// ExitPartitionKeyDesc is called when exiting the partitionKeyDesc production.
	ExitPartitionKeyDesc(c *PartitionKeyDescContext)

	// ExitPartitionValueList is called when exiting the partitionValueList production.
	ExitPartitionValueList(c *PartitionValueListContext)

	// ExitKeyPartition is called when exiting the keyPartition production.
	ExitKeyPartition(c *KeyPartitionContext)

	// ExitPartitionValue is called when exiting the partitionValue production.
	ExitPartitionValue(c *PartitionValueContext)

	// ExitDistributionClause is called when exiting the distributionClause production.
	ExitDistributionClause(c *DistributionClauseContext)

	// ExitDistributionDesc is called when exiting the distributionDesc production.
	ExitDistributionDesc(c *DistributionDescContext)

	// ExitRefreshSchemeDesc is called when exiting the refreshSchemeDesc production.
	ExitRefreshSchemeDesc(c *RefreshSchemeDescContext)

	// ExitStatusDesc is called when exiting the statusDesc production.
	ExitStatusDesc(c *StatusDescContext)

	// ExitProperties is called when exiting the properties production.
	ExitProperties(c *PropertiesContext)

	// ExitExtProperties is called when exiting the extProperties production.
	ExitExtProperties(c *ExtPropertiesContext)

	// ExitPropertyList is called when exiting the propertyList production.
	ExitPropertyList(c *PropertyListContext)

	// ExitUserPropertyList is called when exiting the userPropertyList production.
	ExitUserPropertyList(c *UserPropertyListContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitInlineProperties is called when exiting the inlineProperties production.
	ExitInlineProperties(c *InlinePropertiesContext)

	// ExitInlineProperty is called when exiting the inlineProperty production.
	ExitInlineProperty(c *InlinePropertyContext)

	// ExitVarType is called when exiting the varType production.
	ExitVarType(c *VarTypeContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitOutfile is called when exiting the outfile production.
	ExitOutfile(c *OutfileContext)

	// ExitFileFormat is called when exiting the fileFormat production.
	ExitFileFormat(c *FileFormatContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitTaskInterval is called when exiting the taskInterval production.
	ExitTaskInterval(c *TaskIntervalContext)

	// ExitTaskUnitIdentifier is called when exiting the taskUnitIdentifier production.
	ExitTaskUnitIdentifier(c *TaskUnitIdentifierContext)

	// ExitUnitIdentifier is called when exiting the unitIdentifier production.
	ExitUnitIdentifier(c *UnitIdentifierContext)

	// ExitUnitBoundary is called when exiting the unitBoundary production.
	ExitUnitBoundary(c *UnitBoundaryContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitSubfieldDesc is called when exiting the subfieldDesc production.
	ExitSubfieldDesc(c *SubfieldDescContext)

	// ExitSubfieldDescs is called when exiting the subfieldDescs production.
	ExitSubfieldDescs(c *SubfieldDescsContext)

	// ExitStructType is called when exiting the structType production.
	ExitStructType(c *StructTypeContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitDecimalType is called when exiting the decimalType production.
	ExitDecimalType(c *DecimalTypeContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitWriteBranch is called when exiting the writeBranch production.
	ExitWriteBranch(c *WriteBranchContext)

	// ExitUnquotedIdentifier is called when exiting the unquotedIdentifier production.
	ExitUnquotedIdentifier(c *UnquotedIdentifierContext)

	// ExitDigitIdentifier is called when exiting the digitIdentifier production.
	ExitDigitIdentifier(c *DigitIdentifierContext)

	// ExitBackQuotedIdentifier is called when exiting the backQuotedIdentifier production.
	ExitBackQuotedIdentifier(c *BackQuotedIdentifierContext)

	// ExitIdentifierWithAlias is called when exiting the identifierWithAlias production.
	ExitIdentifierWithAlias(c *IdentifierWithAliasContext)

	// ExitIdentifierWithAliasList is called when exiting the identifierWithAliasList production.
	ExitIdentifierWithAliasList(c *IdentifierWithAliasListContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitIdentifierOrString is called when exiting the identifierOrString production.
	ExitIdentifierOrString(c *IdentifierOrStringContext)

	// ExitIdentifierOrStringList is called when exiting the identifierOrStringList production.
	ExitIdentifierOrStringList(c *IdentifierOrStringListContext)

	// ExitIdentifierOrStringOrStar is called when exiting the identifierOrStringOrStar production.
	ExitIdentifierOrStringOrStar(c *IdentifierOrStringOrStarContext)

	// ExitUserWithoutHost is called when exiting the userWithoutHost production.
	ExitUserWithoutHost(c *UserWithoutHostContext)

	// ExitUserWithHost is called when exiting the userWithHost production.
	ExitUserWithHost(c *UserWithHostContext)

	// ExitUserWithHostAndBlanket is called when exiting the userWithHostAndBlanket production.
	ExitUserWithHostAndBlanket(c *UserWithHostAndBlanketContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAssignmentList is called when exiting the assignmentList production.
	ExitAssignmentList(c *AssignmentListContext)

	// ExitDecimalValue is called when exiting the decimalValue production.
	ExitDecimalValue(c *DecimalValueContext)

	// ExitDoubleValue is called when exiting the doubleValue production.
	ExitDoubleValue(c *DoubleValueContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitNonReserved is called when exiting the nonReserved production.
	ExitNonReserved(c *NonReservedContext)
}
