// Code generated from DorisSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package doris // DorisSQL
import "github.com/antlr4-go/antlr/v4"

// BaseDorisSQLListener is a complete listener for a parse tree produced by DorisSQLParser.
type BaseDorisSQLListener struct{}

var _ DorisSQLListener = &BaseDorisSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDorisSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDorisSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDorisSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDorisSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSqlStatements is called when production sqlStatements is entered.
func (s *BaseDorisSQLListener) EnterSqlStatements(ctx *SqlStatementsContext) {}

// ExitSqlStatements is called when production sqlStatements is exited.
func (s *BaseDorisSQLListener) ExitSqlStatements(ctx *SqlStatementsContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseDorisSQLListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseDorisSQLListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDorisSQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDorisSQLListener) ExitStatement(ctx *StatementContext) {}

// EnterUseDatabaseStatement is called when production useDatabaseStatement is entered.
func (s *BaseDorisSQLListener) EnterUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// ExitUseDatabaseStatement is called when production useDatabaseStatement is exited.
func (s *BaseDorisSQLListener) ExitUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// EnterUseCatalogStatement is called when production useCatalogStatement is entered.
func (s *BaseDorisSQLListener) EnterUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// ExitUseCatalogStatement is called when production useCatalogStatement is exited.
func (s *BaseDorisSQLListener) ExitUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// EnterSetCatalogStatement is called when production setCatalogStatement is entered.
func (s *BaseDorisSQLListener) EnterSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// ExitSetCatalogStatement is called when production setCatalogStatement is exited.
func (s *BaseDorisSQLListener) ExitSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// EnterShowDatabasesStatement is called when production showDatabasesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {}

// ExitShowDatabasesStatement is called when production showDatabasesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {}

// EnterAlterDbQuotaStatement is called when production alterDbQuotaStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// ExitAlterDbQuotaStatement is called when production alterDbQuotaStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// EnterCreateDbStatement is called when production createDbStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateDbStatement(ctx *CreateDbStatementContext) {}

// ExitCreateDbStatement is called when production createDbStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateDbStatement(ctx *CreateDbStatementContext) {}

// EnterDropDbStatement is called when production dropDbStatement is entered.
func (s *BaseDorisSQLListener) EnterDropDbStatement(ctx *DropDbStatementContext) {}

// ExitDropDbStatement is called when production dropDbStatement is exited.
func (s *BaseDorisSQLListener) ExitDropDbStatement(ctx *DropDbStatementContext) {}

// EnterShowCreateDbStatement is called when production showCreateDbStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// ExitShowCreateDbStatement is called when production showCreateDbStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// EnterAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// ExitAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// EnterRecoverDbStmt is called when production recoverDbStmt is entered.
func (s *BaseDorisSQLListener) EnterRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// ExitRecoverDbStmt is called when production recoverDbStmt is exited.
func (s *BaseDorisSQLListener) ExitRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// EnterShowDataStmt is called when production showDataStmt is entered.
func (s *BaseDorisSQLListener) EnterShowDataStmt(ctx *ShowDataStmtContext) {}

// ExitShowDataStmt is called when production showDataStmt is exited.
func (s *BaseDorisSQLListener) ExitShowDataStmt(ctx *ShowDataStmtContext) {}

// EnterShowDataDistributionStmt is called when production showDataDistributionStmt is entered.
func (s *BaseDorisSQLListener) EnterShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {}

// ExitShowDataDistributionStmt is called when production showDataDistributionStmt is exited.
func (s *BaseDorisSQLListener) ExitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterColumnDesc is called when production columnDesc is entered.
func (s *BaseDorisSQLListener) EnterColumnDesc(ctx *ColumnDescContext) {}

// ExitColumnDesc is called when production columnDesc is exited.
func (s *BaseDorisSQLListener) ExitColumnDesc(ctx *ColumnDescContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseDorisSQLListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseDorisSQLListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterDefaultDesc is called when production defaultDesc is entered.
func (s *BaseDorisSQLListener) EnterDefaultDesc(ctx *DefaultDescContext) {}

// ExitDefaultDesc is called when production defaultDesc is exited.
func (s *BaseDorisSQLListener) ExitDefaultDesc(ctx *DefaultDescContext) {}

// EnterGeneratedColumnDesc is called when production generatedColumnDesc is entered.
func (s *BaseDorisSQLListener) EnterGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// ExitGeneratedColumnDesc is called when production generatedColumnDesc is exited.
func (s *BaseDorisSQLListener) ExitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// EnterIndexDesc is called when production indexDesc is entered.
func (s *BaseDorisSQLListener) EnterIndexDesc(ctx *IndexDescContext) {}

// ExitIndexDesc is called when production indexDesc is exited.
func (s *BaseDorisSQLListener) ExitIndexDesc(ctx *IndexDescContext) {}

// EnterEngineDesc is called when production engineDesc is entered.
func (s *BaseDorisSQLListener) EnterEngineDesc(ctx *EngineDescContext) {}

// ExitEngineDesc is called when production engineDesc is exited.
func (s *BaseDorisSQLListener) ExitEngineDesc(ctx *EngineDescContext) {}

// EnterCharsetDesc is called when production charsetDesc is entered.
func (s *BaseDorisSQLListener) EnterCharsetDesc(ctx *CharsetDescContext) {}

// ExitCharsetDesc is called when production charsetDesc is exited.
func (s *BaseDorisSQLListener) ExitCharsetDesc(ctx *CharsetDescContext) {}

// EnterCollateDesc is called when production collateDesc is entered.
func (s *BaseDorisSQLListener) EnterCollateDesc(ctx *CollateDescContext) {}

// ExitCollateDesc is called when production collateDesc is exited.
func (s *BaseDorisSQLListener) ExitCollateDesc(ctx *CollateDescContext) {}

// EnterKeyDesc is called when production keyDesc is entered.
func (s *BaseDorisSQLListener) EnterKeyDesc(ctx *KeyDescContext) {}

// ExitKeyDesc is called when production keyDesc is exited.
func (s *BaseDorisSQLListener) ExitKeyDesc(ctx *KeyDescContext) {}

// EnterOrderByDesc is called when production orderByDesc is entered.
func (s *BaseDorisSQLListener) EnterOrderByDesc(ctx *OrderByDescContext) {}

// ExitOrderByDesc is called when production orderByDesc is exited.
func (s *BaseDorisSQLListener) ExitOrderByDesc(ctx *OrderByDescContext) {}

// EnterColumnNullable is called when production columnNullable is entered.
func (s *BaseDorisSQLListener) EnterColumnNullable(ctx *ColumnNullableContext) {}

// ExitColumnNullable is called when production columnNullable is exited.
func (s *BaseDorisSQLListener) ExitColumnNullable(ctx *ColumnNullableContext) {}

// EnterTypeWithNullable is called when production typeWithNullable is entered.
func (s *BaseDorisSQLListener) EnterTypeWithNullable(ctx *TypeWithNullableContext) {}

// ExitTypeWithNullable is called when production typeWithNullable is exited.
func (s *BaseDorisSQLListener) ExitTypeWithNullable(ctx *TypeWithNullableContext) {}

// EnterAggStateDesc is called when production aggStateDesc is entered.
func (s *BaseDorisSQLListener) EnterAggStateDesc(ctx *AggStateDescContext) {}

// ExitAggStateDesc is called when production aggStateDesc is exited.
func (s *BaseDorisSQLListener) ExitAggStateDesc(ctx *AggStateDescContext) {}

// EnterAggDesc is called when production aggDesc is entered.
func (s *BaseDorisSQLListener) EnterAggDesc(ctx *AggDescContext) {}

// ExitAggDesc is called when production aggDesc is exited.
func (s *BaseDorisSQLListener) ExitAggDesc(ctx *AggDescContext) {}

// EnterRollupDesc is called when production rollupDesc is entered.
func (s *BaseDorisSQLListener) EnterRollupDesc(ctx *RollupDescContext) {}

// ExitRollupDesc is called when production rollupDesc is exited.
func (s *BaseDorisSQLListener) ExitRollupDesc(ctx *RollupDescContext) {}

// EnterRollupItem is called when production rollupItem is entered.
func (s *BaseDorisSQLListener) EnterRollupItem(ctx *RollupItemContext) {}

// ExitRollupItem is called when production rollupItem is exited.
func (s *BaseDorisSQLListener) ExitRollupItem(ctx *RollupItemContext) {}

// EnterDupKeys is called when production dupKeys is entered.
func (s *BaseDorisSQLListener) EnterDupKeys(ctx *DupKeysContext) {}

// ExitDupKeys is called when production dupKeys is exited.
func (s *BaseDorisSQLListener) ExitDupKeys(ctx *DupKeysContext) {}

// EnterFromRollup is called when production fromRollup is entered.
func (s *BaseDorisSQLListener) EnterFromRollup(ctx *FromRollupContext) {}

// ExitFromRollup is called when production fromRollup is exited.
func (s *BaseDorisSQLListener) ExitFromRollup(ctx *FromRollupContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseDorisSQLListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseDorisSQLListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseDorisSQLListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseDorisSQLListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterCreateTableAsSelectStatement is called when production createTableAsSelectStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// ExitCreateTableAsSelectStatement is called when production createTableAsSelectStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseDorisSQLListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseDorisSQLListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is entered.
func (s *BaseDorisSQLListener) EnterCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// ExitCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is exited.
func (s *BaseDorisSQLListener) ExitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterCreateIndexStatement is called when production createIndexStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// ExitCreateIndexStatement is called when production createIndexStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// EnterDropIndexStatement is called when production dropIndexStatement is entered.
func (s *BaseDorisSQLListener) EnterDropIndexStatement(ctx *DropIndexStatementContext) {}

// ExitDropIndexStatement is called when production dropIndexStatement is exited.
func (s *BaseDorisSQLListener) ExitDropIndexStatement(ctx *DropIndexStatementContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseDorisSQLListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseDorisSQLListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterShowTableStatement is called when production showTableStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTableStatement(ctx *ShowTableStatementContext) {}

// ExitShowTableStatement is called when production showTableStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTableStatement(ctx *ShowTableStatementContext) {}

// EnterShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// ExitShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// EnterShowCreateTableStatement is called when production showCreateTableStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {}

// ExitShowCreateTableStatement is called when production showCreateTableStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {}

// EnterShowColumnStatement is called when production showColumnStatement is entered.
func (s *BaseDorisSQLListener) EnterShowColumnStatement(ctx *ShowColumnStatementContext) {}

// ExitShowColumnStatement is called when production showColumnStatement is exited.
func (s *BaseDorisSQLListener) ExitShowColumnStatement(ctx *ShowColumnStatementContext) {}

// EnterShowTableStatusStatement is called when production showTableStatusStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {}

// ExitShowTableStatusStatement is called when production showTableStatusStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {}

// EnterRefreshTableStatement is called when production refreshTableStatement is entered.
func (s *BaseDorisSQLListener) EnterRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// ExitRefreshTableStatement is called when production refreshTableStatement is exited.
func (s *BaseDorisSQLListener) ExitRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// EnterShowAlterStatement is called when production showAlterStatement is entered.
func (s *BaseDorisSQLListener) EnterShowAlterStatement(ctx *ShowAlterStatementContext) {}

// ExitShowAlterStatement is called when production showAlterStatement is exited.
func (s *BaseDorisSQLListener) ExitShowAlterStatement(ctx *ShowAlterStatementContext) {}

// EnterDescTableStatement is called when production descTableStatement is entered.
func (s *BaseDorisSQLListener) EnterDescTableStatement(ctx *DescTableStatementContext) {}

// ExitDescTableStatement is called when production descTableStatement is exited.
func (s *BaseDorisSQLListener) ExitDescTableStatement(ctx *DescTableStatementContext) {}

// EnterCreateTableLikeStatement is called when production createTableLikeStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {}

// ExitCreateTableLikeStatement is called when production createTableLikeStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {}

// EnterShowIndexStatement is called when production showIndexStatement is entered.
func (s *BaseDorisSQLListener) EnterShowIndexStatement(ctx *ShowIndexStatementContext) {}

// ExitShowIndexStatement is called when production showIndexStatement is exited.
func (s *BaseDorisSQLListener) ExitShowIndexStatement(ctx *ShowIndexStatementContext) {}

// EnterRecoverTableStatement is called when production recoverTableStatement is entered.
func (s *BaseDorisSQLListener) EnterRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// ExitRecoverTableStatement is called when production recoverTableStatement is exited.
func (s *BaseDorisSQLListener) ExitRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// EnterTruncateTableStatement is called when production truncateTableStatement is entered.
func (s *BaseDorisSQLListener) EnterTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// ExitTruncateTableStatement is called when production truncateTableStatement is exited.
func (s *BaseDorisSQLListener) ExitTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// EnterCancelAlterTableStatement is called when production cancelAlterTableStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {
}

// ExitCancelAlterTableStatement is called when production cancelAlterTableStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {}

// EnterShowPartitionsStatement is called when production showPartitionsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {}

// ExitShowPartitionsStatement is called when production showPartitionsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {}

// EnterRecoverPartitionStatement is called when production recoverPartitionStatement is entered.
func (s *BaseDorisSQLListener) EnterRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {
}

// ExitRecoverPartitionStatement is called when production recoverPartitionStatement is exited.
func (s *BaseDorisSQLListener) ExitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateViewStatement(ctx *CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateViewStatement(ctx *CreateViewStatementContext) {}

// EnterAlterViewStatement is called when production alterViewStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterViewStatement(ctx *AlterViewStatementContext) {}

// ExitAlterViewStatement is called when production alterViewStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterViewStatement(ctx *AlterViewStatementContext) {}

// EnterDropViewStatement is called when production dropViewStatement is entered.
func (s *BaseDorisSQLListener) EnterDropViewStatement(ctx *DropViewStatementContext) {}

// ExitDropViewStatement is called when production dropViewStatement is exited.
func (s *BaseDorisSQLListener) ExitDropViewStatement(ctx *DropViewStatementContext) {}

// EnterColumnNameWithComment is called when production columnNameWithComment is entered.
func (s *BaseDorisSQLListener) EnterColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// ExitColumnNameWithComment is called when production columnNameWithComment is exited.
func (s *BaseDorisSQLListener) ExitColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// EnterSubmitTaskStatement is called when production submitTaskStatement is entered.
func (s *BaseDorisSQLListener) EnterSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// ExitSubmitTaskStatement is called when production submitTaskStatement is exited.
func (s *BaseDorisSQLListener) ExitSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// EnterTaskClause is called when production taskClause is entered.
func (s *BaseDorisSQLListener) EnterTaskClause(ctx *TaskClauseContext) {}

// ExitTaskClause is called when production taskClause is exited.
func (s *BaseDorisSQLListener) ExitTaskClause(ctx *TaskClauseContext) {}

// EnterDropTaskStatement is called when production dropTaskStatement is entered.
func (s *BaseDorisSQLListener) EnterDropTaskStatement(ctx *DropTaskStatementContext) {}

// ExitDropTaskStatement is called when production dropTaskStatement is exited.
func (s *BaseDorisSQLListener) ExitDropTaskStatement(ctx *DropTaskStatementContext) {}

// EnterTaskScheduleDesc is called when production taskScheduleDesc is entered.
func (s *BaseDorisSQLListener) EnterTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// ExitTaskScheduleDesc is called when production taskScheduleDesc is exited.
func (s *BaseDorisSQLListener) ExitTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// EnterMvPartitionExprs is called when production mvPartitionExprs is entered.
func (s *BaseDorisSQLListener) EnterMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// ExitMvPartitionExprs is called when production mvPartitionExprs is exited.
func (s *BaseDorisSQLListener) ExitMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// EnterMaterializedViewDesc is called when production materializedViewDesc is entered.
func (s *BaseDorisSQLListener) EnterMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// ExitMaterializedViewDesc is called when production materializedViewDesc is exited.
func (s *BaseDorisSQLListener) ExitMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// EnterShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// ExitShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *BaseDorisSQLListener) EnterDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *BaseDorisSQLListener) ExitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// EnterAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// ExitAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// EnterRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is entered.
func (s *BaseDorisSQLListener) EnterRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// ExitRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is exited.
func (s *BaseDorisSQLListener) ExitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// EnterCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// ExitCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// EnterAdminSetConfigStatement is called when production adminSetConfigStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {}

// ExitAdminSetConfigStatement is called when production adminSetConfigStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {}

// EnterAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// ExitAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// EnterAdminShowConfigStatement is called when production adminShowConfigStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {}

// ExitAdminShowConfigStatement is called when production adminShowConfigStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {}

// EnterAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// ExitAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// EnterAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// ExitAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// EnterAdminRepairTableStatement is called when production adminRepairTableStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {
}

// ExitAdminRepairTableStatement is called when production adminRepairTableStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {}

// EnterAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// ExitAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// EnterAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// ExitAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// EnterAdminSetPartitionVersion is called when production adminSetPartitionVersion is entered.
func (s *BaseDorisSQLListener) EnterAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {}

// ExitAdminSetPartitionVersion is called when production adminSetPartitionVersion is exited.
func (s *BaseDorisSQLListener) ExitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {}

// EnterKillStatement is called when production killStatement is entered.
func (s *BaseDorisSQLListener) EnterKillStatement(ctx *KillStatementContext) {}

// ExitKillStatement is called when production killStatement is exited.
func (s *BaseDorisSQLListener) ExitKillStatement(ctx *KillStatementContext) {}

// EnterSyncStatement is called when production syncStatement is entered.
func (s *BaseDorisSQLListener) EnterSyncStatement(ctx *SyncStatementContext) {}

// ExitSyncStatement is called when production syncStatement is exited.
func (s *BaseDorisSQLListener) ExitSyncStatement(ctx *SyncStatementContext) {}

// EnterAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// EnterAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is entered.
func (s *BaseDorisSQLListener) EnterAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is exited.
func (s *BaseDorisSQLListener) ExitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// EnterAlterSystemStatement is called when production alterSystemStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// ExitAlterSystemStatement is called when production alterSystemStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// EnterCancelAlterSystemStatement is called when production cancelAlterSystemStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// ExitCancelAlterSystemStatement is called when production cancelAlterSystemStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// EnterShowComputeNodesStatement is called when production showComputeNodesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {
}

// ExitShowComputeNodesStatement is called when production showComputeNodesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {}

// EnterCreateExternalCatalogStatement is called when production createExternalCatalogStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// ExitCreateExternalCatalogStatement is called when production createExternalCatalogStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// EnterShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// ExitShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// EnterDropExternalCatalogStatement is called when production dropExternalCatalogStatement is entered.
func (s *BaseDorisSQLListener) EnterDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// ExitDropExternalCatalogStatement is called when production dropExternalCatalogStatement is exited.
func (s *BaseDorisSQLListener) ExitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// EnterShowCatalogsStatement is called when production showCatalogsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// ExitShowCatalogsStatement is called when production showCatalogsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// EnterAlterCatalogStatement is called when production alterCatalogStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// ExitAlterCatalogStatement is called when production alterCatalogStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// EnterCreateStorageVolumeStatement is called when production createStorageVolumeStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// ExitCreateStorageVolumeStatement is called when production createStorageVolumeStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// EnterTypeDesc is called when production typeDesc is entered.
func (s *BaseDorisSQLListener) EnterTypeDesc(ctx *TypeDescContext) {}

// ExitTypeDesc is called when production typeDesc is exited.
func (s *BaseDorisSQLListener) ExitTypeDesc(ctx *TypeDescContext) {}

// EnterLocationsDesc is called when production locationsDesc is entered.
func (s *BaseDorisSQLListener) EnterLocationsDesc(ctx *LocationsDescContext) {}

// ExitLocationsDesc is called when production locationsDesc is exited.
func (s *BaseDorisSQLListener) ExitLocationsDesc(ctx *LocationsDescContext) {}

// EnterShowStorageVolumesStatement is called when production showStorageVolumesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// ExitShowStorageVolumesStatement is called when production showStorageVolumesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// EnterDropStorageVolumeStatement is called when production dropStorageVolumeStatement is entered.
func (s *BaseDorisSQLListener) EnterDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// ExitDropStorageVolumeStatement is called when production dropStorageVolumeStatement is exited.
func (s *BaseDorisSQLListener) ExitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// ExitAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeClause is called when production alterStorageVolumeClause is entered.
func (s *BaseDorisSQLListener) EnterAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {}

// ExitAlterStorageVolumeClause is called when production alterStorageVolumeClause is exited.
func (s *BaseDorisSQLListener) ExitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {}

// EnterModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is entered.
func (s *BaseDorisSQLListener) EnterModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// ExitModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is exited.
func (s *BaseDorisSQLListener) ExitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// EnterModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is entered.
func (s *BaseDorisSQLListener) EnterModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// ExitModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is exited.
func (s *BaseDorisSQLListener) ExitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// EnterDescStorageVolumeStatement is called when production descStorageVolumeStatement is entered.
func (s *BaseDorisSQLListener) EnterDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// ExitDescStorageVolumeStatement is called when production descStorageVolumeStatement is exited.
func (s *BaseDorisSQLListener) ExitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// EnterSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is entered.
func (s *BaseDorisSQLListener) EnterSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// ExitSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is exited.
func (s *BaseDorisSQLListener) ExitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// EnterUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is entered.
func (s *BaseDorisSQLListener) EnterUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// ExitUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is exited.
func (s *BaseDorisSQLListener) ExitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// EnterShowFailPointStatement is called when production showFailPointStatement is entered.
func (s *BaseDorisSQLListener) EnterShowFailPointStatement(ctx *ShowFailPointStatementContext) {}

// ExitShowFailPointStatement is called when production showFailPointStatement is exited.
func (s *BaseDorisSQLListener) ExitShowFailPointStatement(ctx *ShowFailPointStatementContext) {}

// EnterCreateDictionaryStatement is called when production createDictionaryStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {
}

// ExitCreateDictionaryStatement is called when production createDictionaryStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {}

// EnterDropDictionaryStatement is called when production dropDictionaryStatement is entered.
func (s *BaseDorisSQLListener) EnterDropDictionaryStatement(ctx *DropDictionaryStatementContext) {}

// ExitDropDictionaryStatement is called when production dropDictionaryStatement is exited.
func (s *BaseDorisSQLListener) ExitDropDictionaryStatement(ctx *DropDictionaryStatementContext) {}

// EnterRefreshDictionaryStatement is called when production refreshDictionaryStatement is entered.
func (s *BaseDorisSQLListener) EnterRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// ExitRefreshDictionaryStatement is called when production refreshDictionaryStatement is exited.
func (s *BaseDorisSQLListener) ExitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// EnterShowDictionaryStatement is called when production showDictionaryStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {}

// ExitShowDictionaryStatement is called when production showDictionaryStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {}

// EnterCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// ExitCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// EnterDictionaryColumnDesc is called when production dictionaryColumnDesc is entered.
func (s *BaseDorisSQLListener) EnterDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// ExitDictionaryColumnDesc is called when production dictionaryColumnDesc is exited.
func (s *BaseDorisSQLListener) ExitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// EnterDictionaryName is called when production dictionaryName is entered.
func (s *BaseDorisSQLListener) EnterDictionaryName(ctx *DictionaryNameContext) {}

// ExitDictionaryName is called when production dictionaryName is exited.
func (s *BaseDorisSQLListener) ExitDictionaryName(ctx *DictionaryNameContext) {}

// EnterAlterClause is called when production alterClause is entered.
func (s *BaseDorisSQLListener) EnterAlterClause(ctx *AlterClauseContext) {}

// ExitAlterClause is called when production alterClause is exited.
func (s *BaseDorisSQLListener) ExitAlterClause(ctx *AlterClauseContext) {}

// EnterAddFrontendClause is called when production addFrontendClause is entered.
func (s *BaseDorisSQLListener) EnterAddFrontendClause(ctx *AddFrontendClauseContext) {}

// ExitAddFrontendClause is called when production addFrontendClause is exited.
func (s *BaseDorisSQLListener) ExitAddFrontendClause(ctx *AddFrontendClauseContext) {}

// EnterDropFrontendClause is called when production dropFrontendClause is entered.
func (s *BaseDorisSQLListener) EnterDropFrontendClause(ctx *DropFrontendClauseContext) {}

// ExitDropFrontendClause is called when production dropFrontendClause is exited.
func (s *BaseDorisSQLListener) ExitDropFrontendClause(ctx *DropFrontendClauseContext) {}

// EnterModifyFrontendHostClause is called when production modifyFrontendHostClause is entered.
func (s *BaseDorisSQLListener) EnterModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {}

// ExitModifyFrontendHostClause is called when production modifyFrontendHostClause is exited.
func (s *BaseDorisSQLListener) ExitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {}

// EnterAddBackendClause is called when production addBackendClause is entered.
func (s *BaseDorisSQLListener) EnterAddBackendClause(ctx *AddBackendClauseContext) {}

// ExitAddBackendClause is called when production addBackendClause is exited.
func (s *BaseDorisSQLListener) ExitAddBackendClause(ctx *AddBackendClauseContext) {}

// EnterDropBackendClause is called when production dropBackendClause is entered.
func (s *BaseDorisSQLListener) EnterDropBackendClause(ctx *DropBackendClauseContext) {}

// ExitDropBackendClause is called when production dropBackendClause is exited.
func (s *BaseDorisSQLListener) ExitDropBackendClause(ctx *DropBackendClauseContext) {}

// EnterDecommissionBackendClause is called when production decommissionBackendClause is entered.
func (s *BaseDorisSQLListener) EnterDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {
}

// ExitDecommissionBackendClause is called when production decommissionBackendClause is exited.
func (s *BaseDorisSQLListener) ExitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {}

// EnterModifyBackendClause is called when production modifyBackendClause is entered.
func (s *BaseDorisSQLListener) EnterModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// ExitModifyBackendClause is called when production modifyBackendClause is exited.
func (s *BaseDorisSQLListener) ExitModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// EnterAddComputeNodeClause is called when production addComputeNodeClause is entered.
func (s *BaseDorisSQLListener) EnterAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// ExitAddComputeNodeClause is called when production addComputeNodeClause is exited.
func (s *BaseDorisSQLListener) ExitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// EnterDropComputeNodeClause is called when production dropComputeNodeClause is entered.
func (s *BaseDorisSQLListener) EnterDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// ExitDropComputeNodeClause is called when production dropComputeNodeClause is exited.
func (s *BaseDorisSQLListener) ExitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// EnterModifyBrokerClause is called when production modifyBrokerClause is entered.
func (s *BaseDorisSQLListener) EnterModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// ExitModifyBrokerClause is called when production modifyBrokerClause is exited.
func (s *BaseDorisSQLListener) ExitModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// EnterAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is entered.
func (s *BaseDorisSQLListener) EnterAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {}

// ExitAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is exited.
func (s *BaseDorisSQLListener) ExitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {}

// EnterCreateImageClause is called when production createImageClause is entered.
func (s *BaseDorisSQLListener) EnterCreateImageClause(ctx *CreateImageClauseContext) {}

// ExitCreateImageClause is called when production createImageClause is exited.
func (s *BaseDorisSQLListener) ExitCreateImageClause(ctx *CreateImageClauseContext) {}

// EnterCleanTabletSchedQClause is called when production cleanTabletSchedQClause is entered.
func (s *BaseDorisSQLListener) EnterCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {}

// ExitCleanTabletSchedQClause is called when production cleanTabletSchedQClause is exited.
func (s *BaseDorisSQLListener) ExitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {}

// EnterDecommissionDiskClause is called when production decommissionDiskClause is entered.
func (s *BaseDorisSQLListener) EnterDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {}

// ExitDecommissionDiskClause is called when production decommissionDiskClause is exited.
func (s *BaseDorisSQLListener) ExitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {}

// EnterCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is entered.
func (s *BaseDorisSQLListener) EnterCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// ExitCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is exited.
func (s *BaseDorisSQLListener) ExitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// EnterDisableDiskClause is called when production disableDiskClause is entered.
func (s *BaseDorisSQLListener) EnterDisableDiskClause(ctx *DisableDiskClauseContext) {}

// ExitDisableDiskClause is called when production disableDiskClause is exited.
func (s *BaseDorisSQLListener) ExitDisableDiskClause(ctx *DisableDiskClauseContext) {}

// EnterCancelDisableDiskClause is called when production cancelDisableDiskClause is entered.
func (s *BaseDorisSQLListener) EnterCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {}

// ExitCancelDisableDiskClause is called when production cancelDisableDiskClause is exited.
func (s *BaseDorisSQLListener) ExitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {}

// EnterCreateIndexClause is called when production createIndexClause is entered.
func (s *BaseDorisSQLListener) EnterCreateIndexClause(ctx *CreateIndexClauseContext) {}

// ExitCreateIndexClause is called when production createIndexClause is exited.
func (s *BaseDorisSQLListener) ExitCreateIndexClause(ctx *CreateIndexClauseContext) {}

// EnterDropIndexClause is called when production dropIndexClause is entered.
func (s *BaseDorisSQLListener) EnterDropIndexClause(ctx *DropIndexClauseContext) {}

// ExitDropIndexClause is called when production dropIndexClause is exited.
func (s *BaseDorisSQLListener) ExitDropIndexClause(ctx *DropIndexClauseContext) {}

// EnterTableRenameClause is called when production tableRenameClause is entered.
func (s *BaseDorisSQLListener) EnterTableRenameClause(ctx *TableRenameClauseContext) {}

// ExitTableRenameClause is called when production tableRenameClause is exited.
func (s *BaseDorisSQLListener) ExitTableRenameClause(ctx *TableRenameClauseContext) {}

// EnterSwapTableClause is called when production swapTableClause is entered.
func (s *BaseDorisSQLListener) EnterSwapTableClause(ctx *SwapTableClauseContext) {}

// ExitSwapTableClause is called when production swapTableClause is exited.
func (s *BaseDorisSQLListener) ExitSwapTableClause(ctx *SwapTableClauseContext) {}

// EnterModifyPropertiesClause is called when production modifyPropertiesClause is entered.
func (s *BaseDorisSQLListener) EnterModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {}

// ExitModifyPropertiesClause is called when production modifyPropertiesClause is exited.
func (s *BaseDorisSQLListener) ExitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {}

// EnterModifyCommentClause is called when production modifyCommentClause is entered.
func (s *BaseDorisSQLListener) EnterModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// ExitModifyCommentClause is called when production modifyCommentClause is exited.
func (s *BaseDorisSQLListener) ExitModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// EnterOptimizeRange is called when production optimizeRange is entered.
func (s *BaseDorisSQLListener) EnterOptimizeRange(ctx *OptimizeRangeContext) {}

// ExitOptimizeRange is called when production optimizeRange is exited.
func (s *BaseDorisSQLListener) ExitOptimizeRange(ctx *OptimizeRangeContext) {}

// EnterOptimizeClause is called when production optimizeClause is entered.
func (s *BaseDorisSQLListener) EnterOptimizeClause(ctx *OptimizeClauseContext) {}

// ExitOptimizeClause is called when production optimizeClause is exited.
func (s *BaseDorisSQLListener) ExitOptimizeClause(ctx *OptimizeClauseContext) {}

// EnterAddColumnClause is called when production addColumnClause is entered.
func (s *BaseDorisSQLListener) EnterAddColumnClause(ctx *AddColumnClauseContext) {}

// ExitAddColumnClause is called when production addColumnClause is exited.
func (s *BaseDorisSQLListener) ExitAddColumnClause(ctx *AddColumnClauseContext) {}

// EnterAddColumnsClause is called when production addColumnsClause is entered.
func (s *BaseDorisSQLListener) EnterAddColumnsClause(ctx *AddColumnsClauseContext) {}

// ExitAddColumnsClause is called when production addColumnsClause is exited.
func (s *BaseDorisSQLListener) ExitAddColumnsClause(ctx *AddColumnsClauseContext) {}

// EnterDropColumnClause is called when production dropColumnClause is entered.
func (s *BaseDorisSQLListener) EnterDropColumnClause(ctx *DropColumnClauseContext) {}

// ExitDropColumnClause is called when production dropColumnClause is exited.
func (s *BaseDorisSQLListener) ExitDropColumnClause(ctx *DropColumnClauseContext) {}

// EnterModifyColumnClause is called when production modifyColumnClause is entered.
func (s *BaseDorisSQLListener) EnterModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// ExitModifyColumnClause is called when production modifyColumnClause is exited.
func (s *BaseDorisSQLListener) ExitModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// EnterModifyColumnCommentClause is called when production modifyColumnCommentClause is entered.
func (s *BaseDorisSQLListener) EnterModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {
}

// ExitModifyColumnCommentClause is called when production modifyColumnCommentClause is exited.
func (s *BaseDorisSQLListener) ExitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {}

// EnterColumnRenameClause is called when production columnRenameClause is entered.
func (s *BaseDorisSQLListener) EnterColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// ExitColumnRenameClause is called when production columnRenameClause is exited.
func (s *BaseDorisSQLListener) ExitColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// EnterReorderColumnsClause is called when production reorderColumnsClause is entered.
func (s *BaseDorisSQLListener) EnterReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// ExitReorderColumnsClause is called when production reorderColumnsClause is exited.
func (s *BaseDorisSQLListener) ExitReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// EnterRollupRenameClause is called when production rollupRenameClause is entered.
func (s *BaseDorisSQLListener) EnterRollupRenameClause(ctx *RollupRenameClauseContext) {}

// ExitRollupRenameClause is called when production rollupRenameClause is exited.
func (s *BaseDorisSQLListener) ExitRollupRenameClause(ctx *RollupRenameClauseContext) {}

// EnterCompactionClause is called when production compactionClause is entered.
func (s *BaseDorisSQLListener) EnterCompactionClause(ctx *CompactionClauseContext) {}

// ExitCompactionClause is called when production compactionClause is exited.
func (s *BaseDorisSQLListener) ExitCompactionClause(ctx *CompactionClauseContext) {}

// EnterSubfieldName is called when production subfieldName is entered.
func (s *BaseDorisSQLListener) EnterSubfieldName(ctx *SubfieldNameContext) {}

// ExitSubfieldName is called when production subfieldName is exited.
func (s *BaseDorisSQLListener) ExitSubfieldName(ctx *SubfieldNameContext) {}

// EnterNestedFieldName is called when production nestedFieldName is entered.
func (s *BaseDorisSQLListener) EnterNestedFieldName(ctx *NestedFieldNameContext) {}

// ExitNestedFieldName is called when production nestedFieldName is exited.
func (s *BaseDorisSQLListener) ExitNestedFieldName(ctx *NestedFieldNameContext) {}

// EnterAddFieldClause is called when production addFieldClause is entered.
func (s *BaseDorisSQLListener) EnterAddFieldClause(ctx *AddFieldClauseContext) {}

// ExitAddFieldClause is called when production addFieldClause is exited.
func (s *BaseDorisSQLListener) ExitAddFieldClause(ctx *AddFieldClauseContext) {}

// EnterDropFieldClause is called when production dropFieldClause is entered.
func (s *BaseDorisSQLListener) EnterDropFieldClause(ctx *DropFieldClauseContext) {}

// ExitDropFieldClause is called when production dropFieldClause is exited.
func (s *BaseDorisSQLListener) ExitDropFieldClause(ctx *DropFieldClauseContext) {}

// EnterCreateOrReplaceTagClause is called when production createOrReplaceTagClause is entered.
func (s *BaseDorisSQLListener) EnterCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {}

// ExitCreateOrReplaceTagClause is called when production createOrReplaceTagClause is exited.
func (s *BaseDorisSQLListener) ExitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {}

// EnterCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is entered.
func (s *BaseDorisSQLListener) EnterCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// ExitCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is exited.
func (s *BaseDorisSQLListener) ExitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// EnterDropBranchClause is called when production dropBranchClause is entered.
func (s *BaseDorisSQLListener) EnterDropBranchClause(ctx *DropBranchClauseContext) {}

// ExitDropBranchClause is called when production dropBranchClause is exited.
func (s *BaseDorisSQLListener) ExitDropBranchClause(ctx *DropBranchClauseContext) {}

// EnterDropTagClause is called when production dropTagClause is entered.
func (s *BaseDorisSQLListener) EnterDropTagClause(ctx *DropTagClauseContext) {}

// ExitDropTagClause is called when production dropTagClause is exited.
func (s *BaseDorisSQLListener) ExitDropTagClause(ctx *DropTagClauseContext) {}

// EnterTableOperationClause is called when production tableOperationClause is entered.
func (s *BaseDorisSQLListener) EnterTableOperationClause(ctx *TableOperationClauseContext) {}

// ExitTableOperationClause is called when production tableOperationClause is exited.
func (s *BaseDorisSQLListener) ExitTableOperationClause(ctx *TableOperationClauseContext) {}

// EnterTagOptions is called when production tagOptions is entered.
func (s *BaseDorisSQLListener) EnterTagOptions(ctx *TagOptionsContext) {}

// ExitTagOptions is called when production tagOptions is exited.
func (s *BaseDorisSQLListener) ExitTagOptions(ctx *TagOptionsContext) {}

// EnterBranchOptions is called when production branchOptions is entered.
func (s *BaseDorisSQLListener) EnterBranchOptions(ctx *BranchOptionsContext) {}

// ExitBranchOptions is called when production branchOptions is exited.
func (s *BaseDorisSQLListener) ExitBranchOptions(ctx *BranchOptionsContext) {}

// EnterSnapshotRetention is called when production snapshotRetention is entered.
func (s *BaseDorisSQLListener) EnterSnapshotRetention(ctx *SnapshotRetentionContext) {}

// ExitSnapshotRetention is called when production snapshotRetention is exited.
func (s *BaseDorisSQLListener) ExitSnapshotRetention(ctx *SnapshotRetentionContext) {}

// EnterRefRetain is called when production refRetain is entered.
func (s *BaseDorisSQLListener) EnterRefRetain(ctx *RefRetainContext) {}

// ExitRefRetain is called when production refRetain is exited.
func (s *BaseDorisSQLListener) ExitRefRetain(ctx *RefRetainContext) {}

// EnterMaxSnapshotAge is called when production maxSnapshotAge is entered.
func (s *BaseDorisSQLListener) EnterMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// ExitMaxSnapshotAge is called when production maxSnapshotAge is exited.
func (s *BaseDorisSQLListener) ExitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// EnterMinSnapshotsToKeep is called when production minSnapshotsToKeep is entered.
func (s *BaseDorisSQLListener) EnterMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// ExitMinSnapshotsToKeep is called when production minSnapshotsToKeep is exited.
func (s *BaseDorisSQLListener) ExitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// EnterSnapshotId is called when production snapshotId is entered.
func (s *BaseDorisSQLListener) EnterSnapshotId(ctx *SnapshotIdContext) {}

// ExitSnapshotId is called when production snapshotId is exited.
func (s *BaseDorisSQLListener) ExitSnapshotId(ctx *SnapshotIdContext) {}

// EnterTimeUnit is called when production timeUnit is entered.
func (s *BaseDorisSQLListener) EnterTimeUnit(ctx *TimeUnitContext) {}

// ExitTimeUnit is called when production timeUnit is exited.
func (s *BaseDorisSQLListener) ExitTimeUnit(ctx *TimeUnitContext) {}

// EnterInteger_list is called when production integer_list is entered.
func (s *BaseDorisSQLListener) EnterInteger_list(ctx *Integer_listContext) {}

// ExitInteger_list is called when production integer_list is exited.
func (s *BaseDorisSQLListener) ExitInteger_list(ctx *Integer_listContext) {}

// EnterDropPersistentIndexClause is called when production dropPersistentIndexClause is entered.
func (s *BaseDorisSQLListener) EnterDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {
}

// ExitDropPersistentIndexClause is called when production dropPersistentIndexClause is exited.
func (s *BaseDorisSQLListener) ExitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {}

// EnterAddPartitionClause is called when production addPartitionClause is entered.
func (s *BaseDorisSQLListener) EnterAddPartitionClause(ctx *AddPartitionClauseContext) {}

// ExitAddPartitionClause is called when production addPartitionClause is exited.
func (s *BaseDorisSQLListener) ExitAddPartitionClause(ctx *AddPartitionClauseContext) {}

// EnterDropPartitionClause is called when production dropPartitionClause is entered.
func (s *BaseDorisSQLListener) EnterDropPartitionClause(ctx *DropPartitionClauseContext) {}

// ExitDropPartitionClause is called when production dropPartitionClause is exited.
func (s *BaseDorisSQLListener) ExitDropPartitionClause(ctx *DropPartitionClauseContext) {}

// EnterTruncatePartitionClause is called when production truncatePartitionClause is entered.
func (s *BaseDorisSQLListener) EnterTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {}

// ExitTruncatePartitionClause is called when production truncatePartitionClause is exited.
func (s *BaseDorisSQLListener) ExitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {}

// EnterModifyPartitionClause is called when production modifyPartitionClause is entered.
func (s *BaseDorisSQLListener) EnterModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// ExitModifyPartitionClause is called when production modifyPartitionClause is exited.
func (s *BaseDorisSQLListener) ExitModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// EnterReplacePartitionClause is called when production replacePartitionClause is entered.
func (s *BaseDorisSQLListener) EnterReplacePartitionClause(ctx *ReplacePartitionClauseContext) {}

// ExitReplacePartitionClause is called when production replacePartitionClause is exited.
func (s *BaseDorisSQLListener) ExitReplacePartitionClause(ctx *ReplacePartitionClauseContext) {}

// EnterPartitionRenameClause is called when production partitionRenameClause is entered.
func (s *BaseDorisSQLListener) EnterPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// ExitPartitionRenameClause is called when production partitionRenameClause is exited.
func (s *BaseDorisSQLListener) ExitPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseDorisSQLListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseDorisSQLListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is entered.
func (s *BaseDorisSQLListener) EnterInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// ExitInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is exited.
func (s *BaseDorisSQLListener) ExitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// EnterColumnAliasesOrByName is called when production columnAliasesOrByName is entered.
func (s *BaseDorisSQLListener) EnterColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// ExitColumnAliasesOrByName is called when production columnAliasesOrByName is exited.
func (s *BaseDorisSQLListener) ExitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseDorisSQLListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseDorisSQLListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseDorisSQLListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseDorisSQLListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterCreateRoutineLoadStatement is called when production createRoutineLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// ExitCreateRoutineLoadStatement is called when production createRoutineLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// EnterAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {
}

// ExitAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {}

// EnterDataSource is called when production dataSource is entered.
func (s *BaseDorisSQLListener) EnterDataSource(ctx *DataSourceContext) {}

// ExitDataSource is called when production dataSource is exited.
func (s *BaseDorisSQLListener) ExitDataSource(ctx *DataSourceContext) {}

// EnterLoadProperties is called when production loadProperties is entered.
func (s *BaseDorisSQLListener) EnterLoadProperties(ctx *LoadPropertiesContext) {}

// ExitLoadProperties is called when production loadProperties is exited.
func (s *BaseDorisSQLListener) ExitLoadProperties(ctx *LoadPropertiesContext) {}

// EnterColSeparatorProperty is called when production colSeparatorProperty is entered.
func (s *BaseDorisSQLListener) EnterColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// ExitColSeparatorProperty is called when production colSeparatorProperty is exited.
func (s *BaseDorisSQLListener) ExitColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// EnterRowDelimiterProperty is called when production rowDelimiterProperty is entered.
func (s *BaseDorisSQLListener) EnterRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// ExitRowDelimiterProperty is called when production rowDelimiterProperty is exited.
func (s *BaseDorisSQLListener) ExitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// EnterImportColumns is called when production importColumns is entered.
func (s *BaseDorisSQLListener) EnterImportColumns(ctx *ImportColumnsContext) {}

// ExitImportColumns is called when production importColumns is exited.
func (s *BaseDorisSQLListener) ExitImportColumns(ctx *ImportColumnsContext) {}

// EnterColumnProperties is called when production columnProperties is entered.
func (s *BaseDorisSQLListener) EnterColumnProperties(ctx *ColumnPropertiesContext) {}

// ExitColumnProperties is called when production columnProperties is exited.
func (s *BaseDorisSQLListener) ExitColumnProperties(ctx *ColumnPropertiesContext) {}

// EnterJobProperties is called when production jobProperties is entered.
func (s *BaseDorisSQLListener) EnterJobProperties(ctx *JobPropertiesContext) {}

// ExitJobProperties is called when production jobProperties is exited.
func (s *BaseDorisSQLListener) ExitJobProperties(ctx *JobPropertiesContext) {}

// EnterDataSourceProperties is called when production dataSourceProperties is entered.
func (s *BaseDorisSQLListener) EnterDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// ExitDataSourceProperties is called when production dataSourceProperties is exited.
func (s *BaseDorisSQLListener) ExitDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// EnterStopRoutineLoadStatement is called when production stopRoutineLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {}

// ExitStopRoutineLoadStatement is called when production stopRoutineLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {}

// EnterResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// ExitResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// EnterPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {
}

// ExitPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {}

// EnterShowRoutineLoadStatement is called when production showRoutineLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {}

// ExitShowRoutineLoadStatement is called when production showRoutineLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {}

// EnterShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is entered.
func (s *BaseDorisSQLListener) EnterShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// ExitShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is exited.
func (s *BaseDorisSQLListener) ExitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// EnterShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// ExitShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// EnterShowStreamLoadStatement is called when production showStreamLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {}

// ExitShowStreamLoadStatement is called when production showStreamLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {}

// EnterAnalyzeStatement is called when production analyzeStatement is entered.
func (s *BaseDorisSQLListener) EnterAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// ExitAnalyzeStatement is called when production analyzeStatement is exited.
func (s *BaseDorisSQLListener) ExitAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// EnterRegularColumns is called when production regularColumns is entered.
func (s *BaseDorisSQLListener) EnterRegularColumns(ctx *RegularColumnsContext) {}

// ExitRegularColumns is called when production regularColumns is exited.
func (s *BaseDorisSQLListener) ExitRegularColumns(ctx *RegularColumnsContext) {}

// EnterAllColumns is called when production allColumns is entered.
func (s *BaseDorisSQLListener) EnterAllColumns(ctx *AllColumnsContext) {}

// ExitAllColumns is called when production allColumns is exited.
func (s *BaseDorisSQLListener) ExitAllColumns(ctx *AllColumnsContext) {}

// EnterPredicateColumns is called when production predicateColumns is entered.
func (s *BaseDorisSQLListener) EnterPredicateColumns(ctx *PredicateColumnsContext) {}

// ExitPredicateColumns is called when production predicateColumns is exited.
func (s *BaseDorisSQLListener) ExitPredicateColumns(ctx *PredicateColumnsContext) {}

// EnterMultiColumnSet is called when production multiColumnSet is entered.
func (s *BaseDorisSQLListener) EnterMultiColumnSet(ctx *MultiColumnSetContext) {}

// ExitMultiColumnSet is called when production multiColumnSet is exited.
func (s *BaseDorisSQLListener) ExitMultiColumnSet(ctx *MultiColumnSetContext) {}

// EnterDropStatsStatement is called when production dropStatsStatement is entered.
func (s *BaseDorisSQLListener) EnterDropStatsStatement(ctx *DropStatsStatementContext) {}

// ExitDropStatsStatement is called when production dropStatsStatement is exited.
func (s *BaseDorisSQLListener) ExitDropStatsStatement(ctx *DropStatsStatementContext) {}

// EnterHistogramStatement is called when production histogramStatement is entered.
func (s *BaseDorisSQLListener) EnterHistogramStatement(ctx *HistogramStatementContext) {}

// ExitHistogramStatement is called when production histogramStatement is exited.
func (s *BaseDorisSQLListener) ExitHistogramStatement(ctx *HistogramStatementContext) {}

// EnterAnalyzeHistogramStatement is called when production analyzeHistogramStatement is entered.
func (s *BaseDorisSQLListener) EnterAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {
}

// ExitAnalyzeHistogramStatement is called when production analyzeHistogramStatement is exited.
func (s *BaseDorisSQLListener) ExitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {}

// EnterDropHistogramStatement is called when production dropHistogramStatement is entered.
func (s *BaseDorisSQLListener) EnterDropHistogramStatement(ctx *DropHistogramStatementContext) {}

// ExitDropHistogramStatement is called when production dropHistogramStatement is exited.
func (s *BaseDorisSQLListener) ExitDropHistogramStatement(ctx *DropHistogramStatementContext) {}

// EnterCreateAnalyzeStatement is called when production createAnalyzeStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {}

// ExitCreateAnalyzeStatement is called when production createAnalyzeStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {}

// EnterDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is entered.
func (s *BaseDorisSQLListener) EnterDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {}

// ExitDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is exited.
func (s *BaseDorisSQLListener) ExitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {}

// EnterShowAnalyzeStatement is called when production showAnalyzeStatement is entered.
func (s *BaseDorisSQLListener) EnterShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// ExitShowAnalyzeStatement is called when production showAnalyzeStatement is exited.
func (s *BaseDorisSQLListener) ExitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// EnterShowStatsMetaStatement is called when production showStatsMetaStatement is entered.
func (s *BaseDorisSQLListener) EnterShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {}

// ExitShowStatsMetaStatement is called when production showStatsMetaStatement is exited.
func (s *BaseDorisSQLListener) ExitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {}

// EnterShowHistogramMetaStatement is called when production showHistogramMetaStatement is entered.
func (s *BaseDorisSQLListener) EnterShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// ExitShowHistogramMetaStatement is called when production showHistogramMetaStatement is exited.
func (s *BaseDorisSQLListener) ExitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// EnterKillAnalyzeStatement is called when production killAnalyzeStatement is entered.
func (s *BaseDorisSQLListener) EnterKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// ExitKillAnalyzeStatement is called when production killAnalyzeStatement is exited.
func (s *BaseDorisSQLListener) ExitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// EnterAnalyzeProfileStatement is called when production analyzeProfileStatement is entered.
func (s *BaseDorisSQLListener) EnterAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {}

// ExitAnalyzeProfileStatement is called when production analyzeProfileStatement is exited.
func (s *BaseDorisSQLListener) ExitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {}

// EnterCreateBaselinePlanStatement is called when production createBaselinePlanStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// ExitCreateBaselinePlanStatement is called when production createBaselinePlanStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// EnterDropBaselinePlanStatement is called when production dropBaselinePlanStatement is entered.
func (s *BaseDorisSQLListener) EnterDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {
}

// ExitDropBaselinePlanStatement is called when production dropBaselinePlanStatement is exited.
func (s *BaseDorisSQLListener) ExitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {}

// EnterShowBaselinePlanStatement is called when production showBaselinePlanStatement is entered.
func (s *BaseDorisSQLListener) EnterShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {
}

// ExitShowBaselinePlanStatement is called when production showBaselinePlanStatement is exited.
func (s *BaseDorisSQLListener) ExitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {}

// EnterCreateResourceGroupStatement is called when production createResourceGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// ExitCreateResourceGroupStatement is called when production createResourceGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// EnterDropResourceGroupStatement is called when production dropResourceGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// ExitDropResourceGroupStatement is called when production dropResourceGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// EnterAlterResourceGroupStatement is called when production alterResourceGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// ExitAlterResourceGroupStatement is called when production alterResourceGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// EnterShowResourceGroupStatement is called when production showResourceGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// ExitShowResourceGroupStatement is called when production showResourceGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// EnterShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is entered.
func (s *BaseDorisSQLListener) EnterShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// ExitShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is exited.
func (s *BaseDorisSQLListener) ExitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// EnterCreateResourceStatement is called when production createResourceStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateResourceStatement(ctx *CreateResourceStatementContext) {}

// ExitCreateResourceStatement is called when production createResourceStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateResourceStatement(ctx *CreateResourceStatementContext) {}

// EnterAlterResourceStatement is called when production alterResourceStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterResourceStatement(ctx *AlterResourceStatementContext) {}

// ExitAlterResourceStatement is called when production alterResourceStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterResourceStatement(ctx *AlterResourceStatementContext) {}

// EnterDropResourceStatement is called when production dropResourceStatement is entered.
func (s *BaseDorisSQLListener) EnterDropResourceStatement(ctx *DropResourceStatementContext) {}

// ExitDropResourceStatement is called when production dropResourceStatement is exited.
func (s *BaseDorisSQLListener) ExitDropResourceStatement(ctx *DropResourceStatementContext) {}

// EnterShowResourceStatement is called when production showResourceStatement is entered.
func (s *BaseDorisSQLListener) EnterShowResourceStatement(ctx *ShowResourceStatementContext) {}

// ExitShowResourceStatement is called when production showResourceStatement is exited.
func (s *BaseDorisSQLListener) ExitShowResourceStatement(ctx *ShowResourceStatementContext) {}

// EnterClassifier is called when production classifier is entered.
func (s *BaseDorisSQLListener) EnterClassifier(ctx *ClassifierContext) {}

// ExitClassifier is called when production classifier is exited.
func (s *BaseDorisSQLListener) ExitClassifier(ctx *ClassifierContext) {}

// EnterShowFunctionsStatement is called when production showFunctionsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {}

// ExitShowFunctionsStatement is called when production showFunctionsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {}

// EnterDropFunctionStatement is called when production dropFunctionStatement is entered.
func (s *BaseDorisSQLListener) EnterDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// ExitDropFunctionStatement is called when production dropFunctionStatement is exited.
func (s *BaseDorisSQLListener) ExitDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateFunctionStatement(ctx *CreateFunctionStatementContext) {}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateFunctionStatement(ctx *CreateFunctionStatementContext) {}

// EnterInlineFunction is called when production inlineFunction is entered.
func (s *BaseDorisSQLListener) EnterInlineFunction(ctx *InlineFunctionContext) {}

// ExitInlineFunction is called when production inlineFunction is exited.
func (s *BaseDorisSQLListener) ExitInlineFunction(ctx *InlineFunctionContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseDorisSQLListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseDorisSQLListener) ExitTypeList(ctx *TypeListContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseDorisSQLListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseDorisSQLListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseDorisSQLListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseDorisSQLListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterDataDescList is called when production dataDescList is entered.
func (s *BaseDorisSQLListener) EnterDataDescList(ctx *DataDescListContext) {}

// ExitDataDescList is called when production dataDescList is exited.
func (s *BaseDorisSQLListener) ExitDataDescList(ctx *DataDescListContext) {}

// EnterDataDesc is called when production dataDesc is entered.
func (s *BaseDorisSQLListener) EnterDataDesc(ctx *DataDescContext) {}

// ExitDataDesc is called when production dataDesc is exited.
func (s *BaseDorisSQLListener) ExitDataDesc(ctx *DataDescContext) {}

// EnterFormatProps is called when production formatProps is entered.
func (s *BaseDorisSQLListener) EnterFormatProps(ctx *FormatPropsContext) {}

// ExitFormatProps is called when production formatProps is exited.
func (s *BaseDorisSQLListener) ExitFormatProps(ctx *FormatPropsContext) {}

// EnterBrokerDesc is called when production brokerDesc is entered.
func (s *BaseDorisSQLListener) EnterBrokerDesc(ctx *BrokerDescContext) {}

// ExitBrokerDesc is called when production brokerDesc is exited.
func (s *BaseDorisSQLListener) ExitBrokerDesc(ctx *BrokerDescContext) {}

// EnterResourceDesc is called when production resourceDesc is entered.
func (s *BaseDorisSQLListener) EnterResourceDesc(ctx *ResourceDescContext) {}

// ExitResourceDesc is called when production resourceDesc is exited.
func (s *BaseDorisSQLListener) ExitResourceDesc(ctx *ResourceDescContext) {}

// EnterShowLoadStatement is called when production showLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterShowLoadStatement(ctx *ShowLoadStatementContext) {}

// ExitShowLoadStatement is called when production showLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitShowLoadStatement(ctx *ShowLoadStatementContext) {}

// EnterShowLoadWarningsStatement is called when production showLoadWarningsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {
}

// ExitShowLoadWarningsStatement is called when production showLoadWarningsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {}

// EnterCancelLoadStatement is called when production cancelLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// ExitCancelLoadStatement is called when production cancelLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// EnterAlterLoadStatement is called when production alterLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// ExitAlterLoadStatement is called when production alterLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// EnterCancelCompactionStatement is called when production cancelCompactionStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelCompactionStatement(ctx *CancelCompactionStatementContext) {
}

// ExitCancelCompactionStatement is called when production cancelCompactionStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelCompactionStatement(ctx *CancelCompactionStatementContext) {}

// EnterShowAuthorStatement is called when production showAuthorStatement is entered.
func (s *BaseDorisSQLListener) EnterShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// ExitShowAuthorStatement is called when production showAuthorStatement is exited.
func (s *BaseDorisSQLListener) ExitShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// EnterShowBackendsStatement is called when production showBackendsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// ExitShowBackendsStatement is called when production showBackendsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// EnterShowBrokerStatement is called when production showBrokerStatement is entered.
func (s *BaseDorisSQLListener) EnterShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// ExitShowBrokerStatement is called when production showBrokerStatement is exited.
func (s *BaseDorisSQLListener) ExitShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// EnterShowCharsetStatement is called when production showCharsetStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// ExitShowCharsetStatement is called when production showCharsetStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// EnterShowCollationStatement is called when production showCollationStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCollationStatement(ctx *ShowCollationStatementContext) {}

// ExitShowCollationStatement is called when production showCollationStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCollationStatement(ctx *ShowCollationStatementContext) {}

// EnterShowDeleteStatement is called when production showDeleteStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// ExitShowDeleteStatement is called when production showDeleteStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// EnterShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// ExitShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// EnterShowEventsStatement is called when production showEventsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowEventsStatement(ctx *ShowEventsStatementContext) {}

// ExitShowEventsStatement is called when production showEventsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowEventsStatement(ctx *ShowEventsStatementContext) {}

// EnterShowEnginesStatement is called when production showEnginesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// ExitShowEnginesStatement is called when production showEnginesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// EnterShowFrontendsStatement is called when production showFrontendsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {}

// ExitShowFrontendsStatement is called when production showFrontendsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {}

// EnterShowFrontendsDisksStatement is called when production showFrontendsDisksStatement is entered.
func (s *BaseDorisSQLListener) EnterShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) {
}

// ExitShowFrontendsDisksStatement is called when production showFrontendsDisksStatement is exited.
func (s *BaseDorisSQLListener) ExitShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) {
}

// EnterShowPluginsStatement is called when production showPluginsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// ExitShowPluginsStatement is called when production showPluginsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// EnterShowRepositoriesStatement is called when production showRepositoriesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {
}

// ExitShowRepositoriesStatement is called when production showRepositoriesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {}

// EnterShowOpenTableStatement is called when production showOpenTableStatement is entered.
func (s *BaseDorisSQLListener) EnterShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {}

// ExitShowOpenTableStatement is called when production showOpenTableStatement is exited.
func (s *BaseDorisSQLListener) ExitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {}

// EnterShowPrivilegesStatement is called when production showPrivilegesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {}

// ExitShowPrivilegesStatement is called when production showPrivilegesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {}

// EnterShowProcedureStatement is called when production showProcedureStatement is entered.
func (s *BaseDorisSQLListener) EnterShowProcedureStatement(ctx *ShowProcedureStatementContext) {}

// ExitShowProcedureStatement is called when production showProcedureStatement is exited.
func (s *BaseDorisSQLListener) ExitShowProcedureStatement(ctx *ShowProcedureStatementContext) {}

// EnterShowProcStatement is called when production showProcStatement is entered.
func (s *BaseDorisSQLListener) EnterShowProcStatement(ctx *ShowProcStatementContext) {}

// ExitShowProcStatement is called when production showProcStatement is exited.
func (s *BaseDorisSQLListener) ExitShowProcStatement(ctx *ShowProcStatementContext) {}

// EnterShowProcesslistStatement is called when production showProcesslistStatement is entered.
func (s *BaseDorisSQLListener) EnterShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {}

// ExitShowProcesslistStatement is called when production showProcesslistStatement is exited.
func (s *BaseDorisSQLListener) ExitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {}

// EnterShowProfilelistStatement is called when production showProfilelistStatement is entered.
func (s *BaseDorisSQLListener) EnterShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {}

// ExitShowProfilelistStatement is called when production showProfilelistStatement is exited.
func (s *BaseDorisSQLListener) ExitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {}

// EnterShowRunningQueriesStatement is called when production showRunningQueriesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// ExitShowRunningQueriesStatement is called when production showRunningQueriesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// EnterShowStatusStatement is called when production showStatusStatement is entered.
func (s *BaseDorisSQLListener) EnterShowStatusStatement(ctx *ShowStatusStatementContext) {}

// ExitShowStatusStatement is called when production showStatusStatement is exited.
func (s *BaseDorisSQLListener) ExitShowStatusStatement(ctx *ShowStatusStatementContext) {}

// EnterShowTabletStatement is called when production showTabletStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTabletStatement(ctx *ShowTabletStatementContext) {}

// ExitShowTabletStatement is called when production showTabletStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTabletStatement(ctx *ShowTabletStatementContext) {}

// EnterShowTransactionStatement is called when production showTransactionStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTransactionStatement(ctx *ShowTransactionStatementContext) {}

// ExitShowTransactionStatement is called when production showTransactionStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTransactionStatement(ctx *ShowTransactionStatementContext) {}

// EnterShowTriggersStatement is called when production showTriggersStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// ExitShowTriggersStatement is called when production showTriggersStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// EnterShowUserPropertyStatement is called when production showUserPropertyStatement is entered.
func (s *BaseDorisSQLListener) EnterShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {
}

// ExitShowUserPropertyStatement is called when production showUserPropertyStatement is exited.
func (s *BaseDorisSQLListener) ExitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {}

// EnterShowVariablesStatement is called when production showVariablesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowVariablesStatement(ctx *ShowVariablesStatementContext) {}

// ExitShowVariablesStatement is called when production showVariablesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowVariablesStatement(ctx *ShowVariablesStatementContext) {}

// EnterShowWarningStatement is called when production showWarningStatement is entered.
func (s *BaseDorisSQLListener) EnterShowWarningStatement(ctx *ShowWarningStatementContext) {}

// ExitShowWarningStatement is called when production showWarningStatement is exited.
func (s *BaseDorisSQLListener) ExitShowWarningStatement(ctx *ShowWarningStatementContext) {}

// EnterHelpStatement is called when production helpStatement is entered.
func (s *BaseDorisSQLListener) EnterHelpStatement(ctx *HelpStatementContext) {}

// ExitHelpStatement is called when production helpStatement is exited.
func (s *BaseDorisSQLListener) ExitHelpStatement(ctx *HelpStatementContext) {}

// EnterShowQueryProfileStatement is called when production showQueryProfileStatement is entered.
func (s *BaseDorisSQLListener) EnterShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) {
}

// ExitShowQueryProfileStatement is called when production showQueryProfileStatement is exited.
func (s *BaseDorisSQLListener) ExitShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) {}

// EnterShowQueryStatsStatement is called when production showQueryStatsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) {}

// ExitShowQueryStatsStatement is called when production showQueryStatsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) {}

// EnterShowLoadProfileStatement is called when production showLoadProfileStatement is entered.
func (s *BaseDorisSQLListener) EnterShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) {}

// ExitShowLoadProfileStatement is called when production showLoadProfileStatement is exited.
func (s *BaseDorisSQLListener) ExitShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) {}

// EnterShowDataSkewStatement is called when production showDataSkewStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDataSkewStatement(ctx *ShowDataSkewStatementContext) {}

// ExitShowDataSkewStatement is called when production showDataSkewStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDataSkewStatement(ctx *ShowDataSkewStatementContext) {}

// EnterShowDataTypesStatement is called when production showDataTypesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDataTypesStatement(ctx *ShowDataTypesStatementContext) {}

// ExitShowDataTypesStatement is called when production showDataTypesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDataTypesStatement(ctx *ShowDataTypesStatementContext) {}

// EnterShowSyncJobStatement is called when production showSyncJobStatement is entered.
func (s *BaseDorisSQLListener) EnterShowSyncJobStatement(ctx *ShowSyncJobStatementContext) {}

// ExitShowSyncJobStatement is called when production showSyncJobStatement is exited.
func (s *BaseDorisSQLListener) ExitShowSyncJobStatement(ctx *ShowSyncJobStatementContext) {}

// EnterShowPolicyStatement is called when production showPolicyStatement is entered.
func (s *BaseDorisSQLListener) EnterShowPolicyStatement(ctx *ShowPolicyStatementContext) {}

// ExitShowPolicyStatement is called when production showPolicyStatement is exited.
func (s *BaseDorisSQLListener) ExitShowPolicyStatement(ctx *ShowPolicyStatementContext) {}

// EnterShowSqlBlockRuleStatement is called when production showSqlBlockRuleStatement is entered.
func (s *BaseDorisSQLListener) EnterShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) {
}

// ExitShowSqlBlockRuleStatement is called when production showSqlBlockRuleStatement is exited.
func (s *BaseDorisSQLListener) ExitShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) {}

// EnterShowEncryptKeysStatement is called when production showEncryptKeysStatement is entered.
func (s *BaseDorisSQLListener) EnterShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) {}

// ExitShowEncryptKeysStatement is called when production showEncryptKeysStatement is exited.
func (s *BaseDorisSQLListener) ExitShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) {}

// EnterShowCreateLoadStatement is called when production showCreateLoadStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) {}

// ExitShowCreateLoadStatement is called when production showCreateLoadStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) {}

// EnterShowCreateRepositoryStatement is called when production showCreateRepositoryStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) {
}

// ExitShowCreateRepositoryStatement is called when production showCreateRepositoryStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) {
}

// EnterShowLastInsertStatement is called when production showLastInsertStatement is entered.
func (s *BaseDorisSQLListener) EnterShowLastInsertStatement(ctx *ShowLastInsertStatementContext) {}

// ExitShowLastInsertStatement is called when production showLastInsertStatement is exited.
func (s *BaseDorisSQLListener) ExitShowLastInsertStatement(ctx *ShowLastInsertStatementContext) {}

// EnterShowTableIdStatement is called when production showTableIdStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTableIdStatement(ctx *ShowTableIdStatementContext) {}

// ExitShowTableIdStatement is called when production showTableIdStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTableIdStatement(ctx *ShowTableIdStatementContext) {}

// EnterShowDatabaseIdStatement is called when production showDatabaseIdStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) {}

// ExitShowDatabaseIdStatement is called when production showDatabaseIdStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) {}

// EnterShowPartitionIdStatement is called when production showPartitionIdStatement is entered.
func (s *BaseDorisSQLListener) EnterShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) {}

// ExitShowPartitionIdStatement is called when production showPartitionIdStatement is exited.
func (s *BaseDorisSQLListener) ExitShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) {}

// EnterShowTableStatsStatement is called when production showTableStatsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTableStatsStatement(ctx *ShowTableStatsStatementContext) {}

// ExitShowTableStatsStatement is called when production showTableStatsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTableStatsStatement(ctx *ShowTableStatsStatementContext) {}

// EnterShowColumnStatsStatement is called when production showColumnStatsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) {}

// ExitShowColumnStatsStatement is called when production showColumnStatsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) {}

// EnterShowConvertLightSchemaChangeStatement is called when production showConvertLightSchemaChangeStatement is entered.
func (s *BaseDorisSQLListener) EnterShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) {
}

// ExitShowConvertLightSchemaChangeStatement is called when production showConvertLightSchemaChangeStatement is exited.
func (s *BaseDorisSQLListener) ExitShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) {
}

// EnterShowCatalogRecycleBinStatement is called when production showCatalogRecycleBinStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) {
}

// ExitShowCatalogRecycleBinStatement is called when production showCatalogRecycleBinStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) {
}

// EnterShowTrashStatement is called when production showTrashStatement is entered.
func (s *BaseDorisSQLListener) EnterShowTrashStatement(ctx *ShowTrashStatementContext) {}

// ExitShowTrashStatement is called when production showTrashStatement is exited.
func (s *BaseDorisSQLListener) ExitShowTrashStatement(ctx *ShowTrashStatementContext) {}

// EnterShowMigrationsStatement is called when production showMigrationsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowMigrationsStatement(ctx *ShowMigrationsStatementContext) {}

// ExitShowMigrationsStatement is called when production showMigrationsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowMigrationsStatement(ctx *ShowMigrationsStatementContext) {}

// EnterShowWorkloadGroupsStatement is called when production showWorkloadGroupsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) {
}

// ExitShowWorkloadGroupsStatement is called when production showWorkloadGroupsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) {
}

// EnterShowJobTaskStatement is called when production showJobTaskStatement is entered.
func (s *BaseDorisSQLListener) EnterShowJobTaskStatement(ctx *ShowJobTaskStatementContext) {}

// ExitShowJobTaskStatement is called when production showJobTaskStatement is exited.
func (s *BaseDorisSQLListener) ExitShowJobTaskStatement(ctx *ShowJobTaskStatementContext) {}

// EnterCreateUserStatement is called when production createUserStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateUserStatement(ctx *CreateUserStatementContext) {}

// ExitCreateUserStatement is called when production createUserStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateUserStatement(ctx *CreateUserStatementContext) {}

// EnterDropUserStatement is called when production dropUserStatement is entered.
func (s *BaseDorisSQLListener) EnterDropUserStatement(ctx *DropUserStatementContext) {}

// ExitDropUserStatement is called when production dropUserStatement is exited.
func (s *BaseDorisSQLListener) ExitDropUserStatement(ctx *DropUserStatementContext) {}

// EnterAlterUserStatement is called when production alterUserStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterUserStatement(ctx *AlterUserStatementContext) {}

// ExitAlterUserStatement is called when production alterUserStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterUserStatement(ctx *AlterUserStatementContext) {}

// EnterShowUserStatement is called when production showUserStatement is entered.
func (s *BaseDorisSQLListener) EnterShowUserStatement(ctx *ShowUserStatementContext) {}

// ExitShowUserStatement is called when production showUserStatement is exited.
func (s *BaseDorisSQLListener) ExitShowUserStatement(ctx *ShowUserStatementContext) {}

// EnterShowAllAuthentication is called when production showAllAuthentication is entered.
func (s *BaseDorisSQLListener) EnterShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// ExitShowAllAuthentication is called when production showAllAuthentication is exited.
func (s *BaseDorisSQLListener) ExitShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// EnterShowAuthenticationForUser is called when production showAuthenticationForUser is entered.
func (s *BaseDorisSQLListener) EnterShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {
}

// ExitShowAuthenticationForUser is called when production showAuthenticationForUser is exited.
func (s *BaseDorisSQLListener) ExitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {}

// EnterExecuteAsStatement is called when production executeAsStatement is entered.
func (s *BaseDorisSQLListener) EnterExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// ExitExecuteAsStatement is called when production executeAsStatement is exited.
func (s *BaseDorisSQLListener) ExitExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// EnterAlterRoleStatement is called when production alterRoleStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// ExitAlterRoleStatement is called when production alterRoleStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// EnterDropRoleStatement is called when production dropRoleStatement is entered.
func (s *BaseDorisSQLListener) EnterDropRoleStatement(ctx *DropRoleStatementContext) {}

// ExitDropRoleStatement is called when production dropRoleStatement is exited.
func (s *BaseDorisSQLListener) ExitDropRoleStatement(ctx *DropRoleStatementContext) {}

// EnterShowRolesStatement is called when production showRolesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowRolesStatement(ctx *ShowRolesStatementContext) {}

// ExitShowRolesStatement is called when production showRolesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowRolesStatement(ctx *ShowRolesStatementContext) {}

// EnterGrantRoleToUser is called when production grantRoleToUser is entered.
func (s *BaseDorisSQLListener) EnterGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// ExitGrantRoleToUser is called when production grantRoleToUser is exited.
func (s *BaseDorisSQLListener) ExitGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// EnterGrantRoleToRole is called when production grantRoleToRole is entered.
func (s *BaseDorisSQLListener) EnterGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// ExitGrantRoleToRole is called when production grantRoleToRole is exited.
func (s *BaseDorisSQLListener) ExitGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// EnterRevokeRoleFromUser is called when production revokeRoleFromUser is entered.
func (s *BaseDorisSQLListener) EnterRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// ExitRevokeRoleFromUser is called when production revokeRoleFromUser is exited.
func (s *BaseDorisSQLListener) ExitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// EnterRevokeRoleFromRole is called when production revokeRoleFromRole is entered.
func (s *BaseDorisSQLListener) EnterRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// ExitRevokeRoleFromRole is called when production revokeRoleFromRole is exited.
func (s *BaseDorisSQLListener) ExitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// EnterSetRoleStatement is called when production setRoleStatement is entered.
func (s *BaseDorisSQLListener) EnterSetRoleStatement(ctx *SetRoleStatementContext) {}

// ExitSetRoleStatement is called when production setRoleStatement is exited.
func (s *BaseDorisSQLListener) ExitSetRoleStatement(ctx *SetRoleStatementContext) {}

// EnterSetDefaultRoleStatement is called when production setDefaultRoleStatement is entered.
func (s *BaseDorisSQLListener) EnterSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {}

// ExitSetDefaultRoleStatement is called when production setDefaultRoleStatement is exited.
func (s *BaseDorisSQLListener) ExitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {}

// EnterGrantRevokeClause is called when production grantRevokeClause is entered.
func (s *BaseDorisSQLListener) EnterGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// ExitGrantRevokeClause is called when production grantRevokeClause is exited.
func (s *BaseDorisSQLListener) ExitGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// EnterGrantOnUser is called when production grantOnUser is entered.
func (s *BaseDorisSQLListener) EnterGrantOnUser(ctx *GrantOnUserContext) {}

// ExitGrantOnUser is called when production grantOnUser is exited.
func (s *BaseDorisSQLListener) ExitGrantOnUser(ctx *GrantOnUserContext) {}

// EnterGrantOnTableBrief is called when production grantOnTableBrief is entered.
func (s *BaseDorisSQLListener) EnterGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// ExitGrantOnTableBrief is called when production grantOnTableBrief is exited.
func (s *BaseDorisSQLListener) ExitGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// EnterGrantOnFunc is called when production grantOnFunc is entered.
func (s *BaseDorisSQLListener) EnterGrantOnFunc(ctx *GrantOnFuncContext) {}

// ExitGrantOnFunc is called when production grantOnFunc is exited.
func (s *BaseDorisSQLListener) ExitGrantOnFunc(ctx *GrantOnFuncContext) {}

// EnterGrantOnSystem is called when production grantOnSystem is entered.
func (s *BaseDorisSQLListener) EnterGrantOnSystem(ctx *GrantOnSystemContext) {}

// ExitGrantOnSystem is called when production grantOnSystem is exited.
func (s *BaseDorisSQLListener) ExitGrantOnSystem(ctx *GrantOnSystemContext) {}

// EnterGrantOnPrimaryObj is called when production grantOnPrimaryObj is entered.
func (s *BaseDorisSQLListener) EnterGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// ExitGrantOnPrimaryObj is called when production grantOnPrimaryObj is exited.
func (s *BaseDorisSQLListener) ExitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// EnterGrantOnAll is called when production grantOnAll is entered.
func (s *BaseDorisSQLListener) EnterGrantOnAll(ctx *GrantOnAllContext) {}

// ExitGrantOnAll is called when production grantOnAll is exited.
func (s *BaseDorisSQLListener) ExitGrantOnAll(ctx *GrantOnAllContext) {}

// EnterRevokeOnUser is called when production revokeOnUser is entered.
func (s *BaseDorisSQLListener) EnterRevokeOnUser(ctx *RevokeOnUserContext) {}

// ExitRevokeOnUser is called when production revokeOnUser is exited.
func (s *BaseDorisSQLListener) ExitRevokeOnUser(ctx *RevokeOnUserContext) {}

// EnterRevokeOnTableBrief is called when production revokeOnTableBrief is entered.
func (s *BaseDorisSQLListener) EnterRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// ExitRevokeOnTableBrief is called when production revokeOnTableBrief is exited.
func (s *BaseDorisSQLListener) ExitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// EnterRevokeOnFunc is called when production revokeOnFunc is entered.
func (s *BaseDorisSQLListener) EnterRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// ExitRevokeOnFunc is called when production revokeOnFunc is exited.
func (s *BaseDorisSQLListener) ExitRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// EnterRevokeOnSystem is called when production revokeOnSystem is entered.
func (s *BaseDorisSQLListener) EnterRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// ExitRevokeOnSystem is called when production revokeOnSystem is exited.
func (s *BaseDorisSQLListener) ExitRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// EnterRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is entered.
func (s *BaseDorisSQLListener) EnterRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// ExitRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is exited.
func (s *BaseDorisSQLListener) ExitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// EnterRevokeOnAll is called when production revokeOnAll is entered.
func (s *BaseDorisSQLListener) EnterRevokeOnAll(ctx *RevokeOnAllContext) {}

// ExitRevokeOnAll is called when production revokeOnAll is exited.
func (s *BaseDorisSQLListener) ExitRevokeOnAll(ctx *RevokeOnAllContext) {}

// EnterShowGrantsStatement is called when production showGrantsStatement is entered.
func (s *BaseDorisSQLListener) EnterShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// ExitShowGrantsStatement is called when production showGrantsStatement is exited.
func (s *BaseDorisSQLListener) ExitShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// EnterAuthWithoutPlugin is called when production authWithoutPlugin is entered.
func (s *BaseDorisSQLListener) EnterAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// ExitAuthWithoutPlugin is called when production authWithoutPlugin is exited.
func (s *BaseDorisSQLListener) ExitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// EnterAuthWithPlugin is called when production authWithPlugin is entered.
func (s *BaseDorisSQLListener) EnterAuthWithPlugin(ctx *AuthWithPluginContext) {}

// ExitAuthWithPlugin is called when production authWithPlugin is exited.
func (s *BaseDorisSQLListener) ExitAuthWithPlugin(ctx *AuthWithPluginContext) {}

// EnterPrivObjectName is called when production privObjectName is entered.
func (s *BaseDorisSQLListener) EnterPrivObjectName(ctx *PrivObjectNameContext) {}

// ExitPrivObjectName is called when production privObjectName is exited.
func (s *BaseDorisSQLListener) ExitPrivObjectName(ctx *PrivObjectNameContext) {}

// EnterPrivObjectNameList is called when production privObjectNameList is entered.
func (s *BaseDorisSQLListener) EnterPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// ExitPrivObjectNameList is called when production privObjectNameList is exited.
func (s *BaseDorisSQLListener) ExitPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// EnterPrivFunctionObjectNameList is called when production privFunctionObjectNameList is entered.
func (s *BaseDorisSQLListener) EnterPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// ExitPrivFunctionObjectNameList is called when production privFunctionObjectNameList is exited.
func (s *BaseDorisSQLListener) ExitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// EnterPrivilegeTypeList is called when production privilegeTypeList is entered.
func (s *BaseDorisSQLListener) EnterPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// ExitPrivilegeTypeList is called when production privilegeTypeList is exited.
func (s *BaseDorisSQLListener) ExitPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *BaseDorisSQLListener) EnterPrivilegeType(ctx *PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *BaseDorisSQLListener) ExitPrivilegeType(ctx *PrivilegeTypeContext) {}

// EnterPrivObjectType is called when production privObjectType is entered.
func (s *BaseDorisSQLListener) EnterPrivObjectType(ctx *PrivObjectTypeContext) {}

// ExitPrivObjectType is called when production privObjectType is exited.
func (s *BaseDorisSQLListener) ExitPrivObjectType(ctx *PrivObjectTypeContext) {}

// EnterPrivObjectTypePlural is called when production privObjectTypePlural is entered.
func (s *BaseDorisSQLListener) EnterPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// ExitPrivObjectTypePlural is called when production privObjectTypePlural is exited.
func (s *BaseDorisSQLListener) ExitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// EnterCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// ExitCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// EnterAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// ExitAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// EnterDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLListener) EnterDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// ExitDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLListener) ExitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// EnterShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLListener) EnterShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// ExitShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLListener) ExitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// EnterShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// ExitShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// EnterCreateGroupProviderStatement is called when production createGroupProviderStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// ExitCreateGroupProviderStatement is called when production createGroupProviderStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// EnterDropGroupProviderStatement is called when production dropGroupProviderStatement is entered.
func (s *BaseDorisSQLListener) EnterDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// ExitDropGroupProviderStatement is called when production dropGroupProviderStatement is exited.
func (s *BaseDorisSQLListener) ExitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// EnterShowGroupProvidersStatement is called when production showGroupProvidersStatement is entered.
func (s *BaseDorisSQLListener) EnterShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// ExitShowGroupProvidersStatement is called when production showGroupProvidersStatement is exited.
func (s *BaseDorisSQLListener) ExitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// EnterShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is entered.
func (s *BaseDorisSQLListener) EnterShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// ExitShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is exited.
func (s *BaseDorisSQLListener) ExitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// EnterBackupStatement is called when production backupStatement is entered.
func (s *BaseDorisSQLListener) EnterBackupStatement(ctx *BackupStatementContext) {}

// ExitBackupStatement is called when production backupStatement is exited.
func (s *BaseDorisSQLListener) ExitBackupStatement(ctx *BackupStatementContext) {}

// EnterCancelBackupStatement is called when production cancelBackupStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// ExitCancelBackupStatement is called when production cancelBackupStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// EnterShowBackupStatement is called when production showBackupStatement is entered.
func (s *BaseDorisSQLListener) EnterShowBackupStatement(ctx *ShowBackupStatementContext) {}

// ExitShowBackupStatement is called when production showBackupStatement is exited.
func (s *BaseDorisSQLListener) ExitShowBackupStatement(ctx *ShowBackupStatementContext) {}

// EnterRestoreStatement is called when production restoreStatement is entered.
func (s *BaseDorisSQLListener) EnterRestoreStatement(ctx *RestoreStatementContext) {}

// ExitRestoreStatement is called when production restoreStatement is exited.
func (s *BaseDorisSQLListener) ExitRestoreStatement(ctx *RestoreStatementContext) {}

// EnterCancelRestoreStatement is called when production cancelRestoreStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelRestoreStatement(ctx *CancelRestoreStatementContext) {}

// ExitCancelRestoreStatement is called when production cancelRestoreStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelRestoreStatement(ctx *CancelRestoreStatementContext) {}

// EnterShowRestoreStatement is called when production showRestoreStatement is entered.
func (s *BaseDorisSQLListener) EnterShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// ExitShowRestoreStatement is called when production showRestoreStatement is exited.
func (s *BaseDorisSQLListener) ExitShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// EnterShowSnapshotStatement is called when production showSnapshotStatement is entered.
func (s *BaseDorisSQLListener) EnterShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// ExitShowSnapshotStatement is called when production showSnapshotStatement is exited.
func (s *BaseDorisSQLListener) ExitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// EnterCreateRepositoryStatement is called when production createRepositoryStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {
}

// ExitCreateRepositoryStatement is called when production createRepositoryStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {}

// EnterDropRepositoryStatement is called when production dropRepositoryStatement is entered.
func (s *BaseDorisSQLListener) EnterDropRepositoryStatement(ctx *DropRepositoryStatementContext) {}

// ExitDropRepositoryStatement is called when production dropRepositoryStatement is exited.
func (s *BaseDorisSQLListener) ExitDropRepositoryStatement(ctx *DropRepositoryStatementContext) {}

// EnterAddSqlBlackListStatement is called when production addSqlBlackListStatement is entered.
func (s *BaseDorisSQLListener) EnterAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {}

// ExitAddSqlBlackListStatement is called when production addSqlBlackListStatement is exited.
func (s *BaseDorisSQLListener) ExitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {}

// EnterDelSqlBlackListStatement is called when production delSqlBlackListStatement is entered.
func (s *BaseDorisSQLListener) EnterDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {}

// ExitDelSqlBlackListStatement is called when production delSqlBlackListStatement is exited.
func (s *BaseDorisSQLListener) ExitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {}

// EnterShowSqlBlackListStatement is called when production showSqlBlackListStatement is entered.
func (s *BaseDorisSQLListener) EnterShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {
}

// ExitShowSqlBlackListStatement is called when production showSqlBlackListStatement is exited.
func (s *BaseDorisSQLListener) ExitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {}

// EnterShowWhiteListStatement is called when production showWhiteListStatement is entered.
func (s *BaseDorisSQLListener) EnterShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {}

// ExitShowWhiteListStatement is called when production showWhiteListStatement is exited.
func (s *BaseDorisSQLListener) ExitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {}

// EnterAddBackendBlackListStatement is called when production addBackendBlackListStatement is entered.
func (s *BaseDorisSQLListener) EnterAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// ExitAddBackendBlackListStatement is called when production addBackendBlackListStatement is exited.
func (s *BaseDorisSQLListener) ExitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// EnterDelBackendBlackListStatement is called when production delBackendBlackListStatement is entered.
func (s *BaseDorisSQLListener) EnterDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// ExitDelBackendBlackListStatement is called when production delBackendBlackListStatement is exited.
func (s *BaseDorisSQLListener) ExitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// EnterShowBackendBlackListStatement is called when production showBackendBlackListStatement is entered.
func (s *BaseDorisSQLListener) EnterShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// ExitShowBackendBlackListStatement is called when production showBackendBlackListStatement is exited.
func (s *BaseDorisSQLListener) ExitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// EnterDataCacheTarget is called when production dataCacheTarget is entered.
func (s *BaseDorisSQLListener) EnterDataCacheTarget(ctx *DataCacheTargetContext) {}

// ExitDataCacheTarget is called when production dataCacheTarget is exited.
func (s *BaseDorisSQLListener) ExitDataCacheTarget(ctx *DataCacheTargetContext) {}

// EnterCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// ExitCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// EnterShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// ExitShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// EnterDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is entered.
func (s *BaseDorisSQLListener) EnterDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// ExitDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is exited.
func (s *BaseDorisSQLListener) ExitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// EnterClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is entered.
func (s *BaseDorisSQLListener) EnterClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// ExitClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is exited.
func (s *BaseDorisSQLListener) ExitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// EnterDataCacheSelectStatement is called when production dataCacheSelectStatement is entered.
func (s *BaseDorisSQLListener) EnterDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {}

// ExitDataCacheSelectStatement is called when production dataCacheSelectStatement is exited.
func (s *BaseDorisSQLListener) ExitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseDorisSQLListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseDorisSQLListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterCancelExportStatement is called when production cancelExportStatement is entered.
func (s *BaseDorisSQLListener) EnterCancelExportStatement(ctx *CancelExportStatementContext) {}

// ExitCancelExportStatement is called when production cancelExportStatement is exited.
func (s *BaseDorisSQLListener) ExitCancelExportStatement(ctx *CancelExportStatementContext) {}

// EnterShowExportStatement is called when production showExportStatement is entered.
func (s *BaseDorisSQLListener) EnterShowExportStatement(ctx *ShowExportStatementContext) {}

// ExitShowExportStatement is called when production showExportStatement is exited.
func (s *BaseDorisSQLListener) ExitShowExportStatement(ctx *ShowExportStatementContext) {}

// EnterInstallPluginStatement is called when production installPluginStatement is entered.
func (s *BaseDorisSQLListener) EnterInstallPluginStatement(ctx *InstallPluginStatementContext) {}

// ExitInstallPluginStatement is called when production installPluginStatement is exited.
func (s *BaseDorisSQLListener) ExitInstallPluginStatement(ctx *InstallPluginStatementContext) {}

// EnterUninstallPluginStatement is called when production uninstallPluginStatement is entered.
func (s *BaseDorisSQLListener) EnterUninstallPluginStatement(ctx *UninstallPluginStatementContext) {}

// ExitUninstallPluginStatement is called when production uninstallPluginStatement is exited.
func (s *BaseDorisSQLListener) ExitUninstallPluginStatement(ctx *UninstallPluginStatementContext) {}

// EnterCreateFileStatement is called when production createFileStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateFileStatement(ctx *CreateFileStatementContext) {}

// ExitCreateFileStatement is called when production createFileStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateFileStatement(ctx *CreateFileStatementContext) {}

// EnterDropFileStatement is called when production dropFileStatement is entered.
func (s *BaseDorisSQLListener) EnterDropFileStatement(ctx *DropFileStatementContext) {}

// ExitDropFileStatement is called when production dropFileStatement is exited.
func (s *BaseDorisSQLListener) ExitDropFileStatement(ctx *DropFileStatementContext) {}

// EnterShowSmallFilesStatement is called when production showSmallFilesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {}

// ExitShowSmallFilesStatement is called when production showSmallFilesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {}

// EnterCreatePipeStatement is called when production createPipeStatement is entered.
func (s *BaseDorisSQLListener) EnterCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// ExitCreatePipeStatement is called when production createPipeStatement is exited.
func (s *BaseDorisSQLListener) ExitCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// EnterDropPipeStatement is called when production dropPipeStatement is entered.
func (s *BaseDorisSQLListener) EnterDropPipeStatement(ctx *DropPipeStatementContext) {}

// ExitDropPipeStatement is called when production dropPipeStatement is exited.
func (s *BaseDorisSQLListener) ExitDropPipeStatement(ctx *DropPipeStatementContext) {}

// EnterAlterPipeClause is called when production alterPipeClause is entered.
func (s *BaseDorisSQLListener) EnterAlterPipeClause(ctx *AlterPipeClauseContext) {}

// ExitAlterPipeClause is called when production alterPipeClause is exited.
func (s *BaseDorisSQLListener) ExitAlterPipeClause(ctx *AlterPipeClauseContext) {}

// EnterAlterPipeStatement is called when production alterPipeStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// ExitAlterPipeStatement is called when production alterPipeStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// EnterDescPipeStatement is called when production descPipeStatement is entered.
func (s *BaseDorisSQLListener) EnterDescPipeStatement(ctx *DescPipeStatementContext) {}

// ExitDescPipeStatement is called when production descPipeStatement is exited.
func (s *BaseDorisSQLListener) ExitDescPipeStatement(ctx *DescPipeStatementContext) {}

// EnterShowPipeStatement is called when production showPipeStatement is entered.
func (s *BaseDorisSQLListener) EnterShowPipeStatement(ctx *ShowPipeStatementContext) {}

// ExitShowPipeStatement is called when production showPipeStatement is exited.
func (s *BaseDorisSQLListener) ExitShowPipeStatement(ctx *ShowPipeStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseDorisSQLListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseDorisSQLListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterSetNames is called when production setNames is entered.
func (s *BaseDorisSQLListener) EnterSetNames(ctx *SetNamesContext) {}

// ExitSetNames is called when production setNames is exited.
func (s *BaseDorisSQLListener) ExitSetNames(ctx *SetNamesContext) {}

// EnterSetPassword is called when production setPassword is entered.
func (s *BaseDorisSQLListener) EnterSetPassword(ctx *SetPasswordContext) {}

// ExitSetPassword is called when production setPassword is exited.
func (s *BaseDorisSQLListener) ExitSetPassword(ctx *SetPasswordContext) {}

// EnterSetUserVar is called when production setUserVar is entered.
func (s *BaseDorisSQLListener) EnterSetUserVar(ctx *SetUserVarContext) {}

// ExitSetUserVar is called when production setUserVar is exited.
func (s *BaseDorisSQLListener) ExitSetUserVar(ctx *SetUserVarContext) {}

// EnterSetSystemVar is called when production setSystemVar is entered.
func (s *BaseDorisSQLListener) EnterSetSystemVar(ctx *SetSystemVarContext) {}

// ExitSetSystemVar is called when production setSystemVar is exited.
func (s *BaseDorisSQLListener) ExitSetSystemVar(ctx *SetSystemVarContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseDorisSQLListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseDorisSQLListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterTransaction_characteristics is called when production transaction_characteristics is entered.
func (s *BaseDorisSQLListener) EnterTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// ExitTransaction_characteristics is called when production transaction_characteristics is exited.
func (s *BaseDorisSQLListener) ExitTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// EnterTransaction_access_mode is called when production transaction_access_mode is entered.
func (s *BaseDorisSQLListener) EnterTransaction_access_mode(ctx *Transaction_access_modeContext) {}

// ExitTransaction_access_mode is called when production transaction_access_mode is exited.
func (s *BaseDorisSQLListener) ExitTransaction_access_mode(ctx *Transaction_access_modeContext) {}

// EnterIsolation_level is called when production isolation_level is entered.
func (s *BaseDorisSQLListener) EnterIsolation_level(ctx *Isolation_levelContext) {}

// ExitIsolation_level is called when production isolation_level is exited.
func (s *BaseDorisSQLListener) ExitIsolation_level(ctx *Isolation_levelContext) {}

// EnterIsolation_types is called when production isolation_types is entered.
func (s *BaseDorisSQLListener) EnterIsolation_types(ctx *Isolation_typesContext) {}

// ExitIsolation_types is called when production isolation_types is exited.
func (s *BaseDorisSQLListener) ExitIsolation_types(ctx *Isolation_typesContext) {}

// EnterSetExprOrDefault is called when production setExprOrDefault is entered.
func (s *BaseDorisSQLListener) EnterSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// ExitSetExprOrDefault is called when production setExprOrDefault is exited.
func (s *BaseDorisSQLListener) ExitSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// EnterSetUserPropertyStatement is called when production setUserPropertyStatement is entered.
func (s *BaseDorisSQLListener) EnterSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {}

// ExitSetUserPropertyStatement is called when production setUserPropertyStatement is exited.
func (s *BaseDorisSQLListener) ExitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {}

// EnterRoleList is called when production roleList is entered.
func (s *BaseDorisSQLListener) EnterRoleList(ctx *RoleListContext) {}

// ExitRoleList is called when production roleList is exited.
func (s *BaseDorisSQLListener) ExitRoleList(ctx *RoleListContext) {}

// EnterExecuteScriptStatement is called when production executeScriptStatement is entered.
func (s *BaseDorisSQLListener) EnterExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {}

// ExitExecuteScriptStatement is called when production executeScriptStatement is exited.
func (s *BaseDorisSQLListener) ExitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {}

// EnterUnsupportedStatement is called when production unsupportedStatement is entered.
func (s *BaseDorisSQLListener) EnterUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// ExitUnsupportedStatement is called when production unsupportedStatement is exited.
func (s *BaseDorisSQLListener) ExitUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// EnterLock_item is called when production lock_item is entered.
func (s *BaseDorisSQLListener) EnterLock_item(ctx *Lock_itemContext) {}

// ExitLock_item is called when production lock_item is exited.
func (s *BaseDorisSQLListener) ExitLock_item(ctx *Lock_itemContext) {}

// EnterLock_type is called when production lock_type is entered.
func (s *BaseDorisSQLListener) EnterLock_type(ctx *Lock_typeContext) {}

// ExitLock_type is called when production lock_type is exited.
func (s *BaseDorisSQLListener) ExitLock_type(ctx *Lock_typeContext) {}

// EnterAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// ExitAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// EnterTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is entered.
func (s *BaseDorisSQLListener) EnterTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// ExitTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is exited.
func (s *BaseDorisSQLListener) ExitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// EnterAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// ExitAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// EnterShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is entered.
func (s *BaseDorisSQLListener) EnterShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {}

// ExitShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is exited.
func (s *BaseDorisSQLListener) ExitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {}

// EnterCreateWarehouseStatement is called when production createWarehouseStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {}

// ExitCreateWarehouseStatement is called when production createWarehouseStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {}

// EnterDropWarehouseStatement is called when production dropWarehouseStatement is entered.
func (s *BaseDorisSQLListener) EnterDropWarehouseStatement(ctx *DropWarehouseStatementContext) {}

// ExitDropWarehouseStatement is called when production dropWarehouseStatement is exited.
func (s *BaseDorisSQLListener) ExitDropWarehouseStatement(ctx *DropWarehouseStatementContext) {}

// EnterSuspendWarehouseStatement is called when production suspendWarehouseStatement is entered.
func (s *BaseDorisSQLListener) EnterSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {
}

// ExitSuspendWarehouseStatement is called when production suspendWarehouseStatement is exited.
func (s *BaseDorisSQLListener) ExitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {}

// EnterResumeWarehouseStatement is called when production resumeWarehouseStatement is entered.
func (s *BaseDorisSQLListener) EnterResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {}

// ExitResumeWarehouseStatement is called when production resumeWarehouseStatement is exited.
func (s *BaseDorisSQLListener) ExitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {}

// EnterSetWarehouseStatement is called when production setWarehouseStatement is entered.
func (s *BaseDorisSQLListener) EnterSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// ExitSetWarehouseStatement is called when production setWarehouseStatement is exited.
func (s *BaseDorisSQLListener) ExitSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// EnterShowWarehousesStatement is called when production showWarehousesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {}

// ExitShowWarehousesStatement is called when production showWarehousesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {}

// EnterShowClustersStatement is called when production showClustersStatement is entered.
func (s *BaseDorisSQLListener) EnterShowClustersStatement(ctx *ShowClustersStatementContext) {}

// ExitShowClustersStatement is called when production showClustersStatement is exited.
func (s *BaseDorisSQLListener) ExitShowClustersStatement(ctx *ShowClustersStatementContext) {}

// EnterShowNodesStatement is called when production showNodesStatement is entered.
func (s *BaseDorisSQLListener) EnterShowNodesStatement(ctx *ShowNodesStatementContext) {}

// ExitShowNodesStatement is called when production showNodesStatement is exited.
func (s *BaseDorisSQLListener) ExitShowNodesStatement(ctx *ShowNodesStatementContext) {}

// EnterAlterWarehouseStatement is called when production alterWarehouseStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {}

// ExitAlterWarehouseStatement is called when production alterWarehouseStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {}

// EnterCreateCNGroupStatement is called when production createCNGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {}

// ExitCreateCNGroupStatement is called when production createCNGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {}

// EnterDropCNGroupStatement is called when production dropCNGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// ExitDropCNGroupStatement is called when production dropCNGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// EnterEnableCNGroupStatement is called when production enableCNGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {}

// ExitEnableCNGroupStatement is called when production enableCNGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {}

// EnterDisableCNGroupStatement is called when production disableCNGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {}

// ExitDisableCNGroupStatement is called when production disableCNGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {}

// EnterAlterCNGroupStatement is called when production alterCNGroupStatement is entered.
func (s *BaseDorisSQLListener) EnterAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// ExitAlterCNGroupStatement is called when production alterCNGroupStatement is exited.
func (s *BaseDorisSQLListener) ExitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// EnterBeginStatement is called when production beginStatement is entered.
func (s *BaseDorisSQLListener) EnterBeginStatement(ctx *BeginStatementContext) {}

// ExitBeginStatement is called when production beginStatement is exited.
func (s *BaseDorisSQLListener) ExitBeginStatement(ctx *BeginStatementContext) {}

// EnterCommitStatement is called when production commitStatement is entered.
func (s *BaseDorisSQLListener) EnterCommitStatement(ctx *CommitStatementContext) {}

// ExitCommitStatement is called when production commitStatement is exited.
func (s *BaseDorisSQLListener) ExitCommitStatement(ctx *CommitStatementContext) {}

// EnterRollbackStatement is called when production rollbackStatement is entered.
func (s *BaseDorisSQLListener) EnterRollbackStatement(ctx *RollbackStatementContext) {}

// ExitRollbackStatement is called when production rollbackStatement is exited.
func (s *BaseDorisSQLListener) ExitRollbackStatement(ctx *RollbackStatementContext) {}

// EnterTranslateStatement is called when production translateStatement is entered.
func (s *BaseDorisSQLListener) EnterTranslateStatement(ctx *TranslateStatementContext) {}

// ExitTranslateStatement is called when production translateStatement is exited.
func (s *BaseDorisSQLListener) ExitTranslateStatement(ctx *TranslateStatementContext) {}

// EnterDialect is called when production dialect is entered.
func (s *BaseDorisSQLListener) EnterDialect(ctx *DialectContext) {}

// ExitDialect is called when production dialect is exited.
func (s *BaseDorisSQLListener) ExitDialect(ctx *DialectContext) {}

// EnterTranslateSQL is called when production translateSQL is entered.
func (s *BaseDorisSQLListener) EnterTranslateSQL(ctx *TranslateSQLContext) {}

// ExitTranslateSQL is called when production translateSQL is exited.
func (s *BaseDorisSQLListener) ExitTranslateSQL(ctx *TranslateSQLContext) {}

// EnterQueryStatement is called when production queryStatement is entered.
func (s *BaseDorisSQLListener) EnterQueryStatement(ctx *QueryStatementContext) {}

// ExitQueryStatement is called when production queryStatement is exited.
func (s *BaseDorisSQLListener) ExitQueryStatement(ctx *QueryStatementContext) {}

// EnterQueryRelation is called when production queryRelation is entered.
func (s *BaseDorisSQLListener) EnterQueryRelation(ctx *QueryRelationContext) {}

// ExitQueryRelation is called when production queryRelation is exited.
func (s *BaseDorisSQLListener) ExitQueryRelation(ctx *QueryRelationContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseDorisSQLListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseDorisSQLListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterQueryNoWith is called when production queryNoWith is entered.
func (s *BaseDorisSQLListener) EnterQueryNoWith(ctx *QueryNoWithContext) {}

// ExitQueryNoWith is called when production queryNoWith is exited.
func (s *BaseDorisSQLListener) ExitQueryNoWith(ctx *QueryNoWithContext) {}

// EnterQueryPeriod is called when production queryPeriod is entered.
func (s *BaseDorisSQLListener) EnterQueryPeriod(ctx *QueryPeriodContext) {}

// ExitQueryPeriod is called when production queryPeriod is exited.
func (s *BaseDorisSQLListener) ExitQueryPeriod(ctx *QueryPeriodContext) {}

// EnterPeriodType is called when production periodType is entered.
func (s *BaseDorisSQLListener) EnterPeriodType(ctx *PeriodTypeContext) {}

// ExitPeriodType is called when production periodType is exited.
func (s *BaseDorisSQLListener) ExitPeriodType(ctx *PeriodTypeContext) {}

// EnterQueryWithParentheses is called when production queryWithParentheses is entered.
func (s *BaseDorisSQLListener) EnterQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// ExitQueryWithParentheses is called when production queryWithParentheses is exited.
func (s *BaseDorisSQLListener) ExitQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// EnterSetOperation is called when production setOperation is entered.
func (s *BaseDorisSQLListener) EnterSetOperation(ctx *SetOperationContext) {}

// ExitSetOperation is called when production setOperation is exited.
func (s *BaseDorisSQLListener) ExitSetOperation(ctx *SetOperationContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseDorisSQLListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseDorisSQLListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseDorisSQLListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseDorisSQLListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterRowConstructor is called when production rowConstructor is entered.
func (s *BaseDorisSQLListener) EnterRowConstructor(ctx *RowConstructorContext) {}

// ExitRowConstructor is called when production rowConstructor is exited.
func (s *BaseDorisSQLListener) ExitRowConstructor(ctx *RowConstructorContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseDorisSQLListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseDorisSQLListener) ExitSortItem(ctx *SortItemContext) {}

// EnterLimitConstExpr is called when production limitConstExpr is entered.
func (s *BaseDorisSQLListener) EnterLimitConstExpr(ctx *LimitConstExprContext) {}

// ExitLimitConstExpr is called when production limitConstExpr is exited.
func (s *BaseDorisSQLListener) ExitLimitConstExpr(ctx *LimitConstExprContext) {}

// EnterLimitElement is called when production limitElement is entered.
func (s *BaseDorisSQLListener) EnterLimitElement(ctx *LimitElementContext) {}

// ExitLimitElement is called when production limitElement is exited.
func (s *BaseDorisSQLListener) ExitLimitElement(ctx *LimitElementContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseDorisSQLListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseDorisSQLListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterFrom is called when production from is entered.
func (s *BaseDorisSQLListener) EnterFrom(ctx *FromContext) {}

// ExitFrom is called when production from is exited.
func (s *BaseDorisSQLListener) ExitFrom(ctx *FromContext) {}

// EnterDual is called when production dual is entered.
func (s *BaseDorisSQLListener) EnterDual(ctx *DualContext) {}

// ExitDual is called when production dual is exited.
func (s *BaseDorisSQLListener) ExitDual(ctx *DualContext) {}

// EnterRollup is called when production rollup is entered.
func (s *BaseDorisSQLListener) EnterRollup(ctx *RollupContext) {}

// ExitRollup is called when production rollup is exited.
func (s *BaseDorisSQLListener) ExitRollup(ctx *RollupContext) {}

// EnterCube is called when production cube is entered.
func (s *BaseDorisSQLListener) EnterCube(ctx *CubeContext) {}

// ExitCube is called when production cube is exited.
func (s *BaseDorisSQLListener) ExitCube(ctx *CubeContext) {}

// EnterMultipleGroupingSets is called when production multipleGroupingSets is entered.
func (s *BaseDorisSQLListener) EnterMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// ExitMultipleGroupingSets is called when production multipleGroupingSets is exited.
func (s *BaseDorisSQLListener) ExitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// EnterSingleGroupingSet is called when production singleGroupingSet is entered.
func (s *BaseDorisSQLListener) EnterSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// ExitSingleGroupingSet is called when production singleGroupingSet is exited.
func (s *BaseDorisSQLListener) ExitSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// EnterGroupingSet is called when production groupingSet is entered.
func (s *BaseDorisSQLListener) EnterGroupingSet(ctx *GroupingSetContext) {}

// ExitGroupingSet is called when production groupingSet is exited.
func (s *BaseDorisSQLListener) ExitGroupingSet(ctx *GroupingSetContext) {}

// EnterCommonTableExpression is called when production commonTableExpression is entered.
func (s *BaseDorisSQLListener) EnterCommonTableExpression(ctx *CommonTableExpressionContext) {}

// ExitCommonTableExpression is called when production commonTableExpression is exited.
func (s *BaseDorisSQLListener) ExitCommonTableExpression(ctx *CommonTableExpressionContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseDorisSQLListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseDorisSQLListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterSelectSingle is called when production selectSingle is entered.
func (s *BaseDorisSQLListener) EnterSelectSingle(ctx *SelectSingleContext) {}

// ExitSelectSingle is called when production selectSingle is exited.
func (s *BaseDorisSQLListener) ExitSelectSingle(ctx *SelectSingleContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseDorisSQLListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseDorisSQLListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterExcludeClause is called when production excludeClause is entered.
func (s *BaseDorisSQLListener) EnterExcludeClause(ctx *ExcludeClauseContext) {}

// ExitExcludeClause is called when production excludeClause is exited.
func (s *BaseDorisSQLListener) ExitExcludeClause(ctx *ExcludeClauseContext) {}

// EnterRelations is called when production relations is entered.
func (s *BaseDorisSQLListener) EnterRelations(ctx *RelationsContext) {}

// ExitRelations is called when production relations is exited.
func (s *BaseDorisSQLListener) ExitRelations(ctx *RelationsContext) {}

// EnterRelationLateralView is called when production relationLateralView is entered.
func (s *BaseDorisSQLListener) EnterRelationLateralView(ctx *RelationLateralViewContext) {}

// ExitRelationLateralView is called when production relationLateralView is exited.
func (s *BaseDorisSQLListener) ExitRelationLateralView(ctx *RelationLateralViewContext) {}

// EnterLateralView is called when production lateralView is entered.
func (s *BaseDorisSQLListener) EnterLateralView(ctx *LateralViewContext) {}

// ExitLateralView is called when production lateralView is exited.
func (s *BaseDorisSQLListener) ExitLateralView(ctx *LateralViewContext) {}

// EnterGeneratorFunction is called when production generatorFunction is entered.
func (s *BaseDorisSQLListener) EnterGeneratorFunction(ctx *GeneratorFunctionContext) {}

// ExitGeneratorFunction is called when production generatorFunction is exited.
func (s *BaseDorisSQLListener) ExitGeneratorFunction(ctx *GeneratorFunctionContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseDorisSQLListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseDorisSQLListener) ExitRelation(ctx *RelationContext) {}

// EnterTableAtom is called when production tableAtom is entered.
func (s *BaseDorisSQLListener) EnterTableAtom(ctx *TableAtomContext) {}

// ExitTableAtom is called when production tableAtom is exited.
func (s *BaseDorisSQLListener) ExitTableAtom(ctx *TableAtomContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseDorisSQLListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseDorisSQLListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterSubqueryWithAlias is called when production subqueryWithAlias is entered.
func (s *BaseDorisSQLListener) EnterSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// ExitSubqueryWithAlias is called when production subqueryWithAlias is exited.
func (s *BaseDorisSQLListener) ExitSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// EnterTableFunction is called when production tableFunction is entered.
func (s *BaseDorisSQLListener) EnterTableFunction(ctx *TableFunctionContext) {}

// ExitTableFunction is called when production tableFunction is exited.
func (s *BaseDorisSQLListener) ExitTableFunction(ctx *TableFunctionContext) {}

// EnterNormalizedTableFunction is called when production normalizedTableFunction is entered.
func (s *BaseDorisSQLListener) EnterNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {}

// ExitNormalizedTableFunction is called when production normalizedTableFunction is exited.
func (s *BaseDorisSQLListener) ExitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {}

// EnterFileTableFunction is called when production fileTableFunction is entered.
func (s *BaseDorisSQLListener) EnterFileTableFunction(ctx *FileTableFunctionContext) {}

// ExitFileTableFunction is called when production fileTableFunction is exited.
func (s *BaseDorisSQLListener) ExitFileTableFunction(ctx *FileTableFunctionContext) {}

// EnterParenthesizedRelation is called when production parenthesizedRelation is entered.
func (s *BaseDorisSQLListener) EnterParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// ExitParenthesizedRelation is called when production parenthesizedRelation is exited.
func (s *BaseDorisSQLListener) ExitParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// EnterPivotClause is called when production pivotClause is entered.
func (s *BaseDorisSQLListener) EnterPivotClause(ctx *PivotClauseContext) {}

// ExitPivotClause is called when production pivotClause is exited.
func (s *BaseDorisSQLListener) ExitPivotClause(ctx *PivotClauseContext) {}

// EnterPivotAggregationExpression is called when production pivotAggregationExpression is entered.
func (s *BaseDorisSQLListener) EnterPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// ExitPivotAggregationExpression is called when production pivotAggregationExpression is exited.
func (s *BaseDorisSQLListener) ExitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// EnterPivotValue is called when production pivotValue is entered.
func (s *BaseDorisSQLListener) EnterPivotValue(ctx *PivotValueContext) {}

// ExitPivotValue is called when production pivotValue is exited.
func (s *BaseDorisSQLListener) ExitPivotValue(ctx *PivotValueContext) {}

// EnterSampleClause is called when production sampleClause is entered.
func (s *BaseDorisSQLListener) EnterSampleClause(ctx *SampleClauseContext) {}

// ExitSampleClause is called when production sampleClause is exited.
func (s *BaseDorisSQLListener) ExitSampleClause(ctx *SampleClauseContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseDorisSQLListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseDorisSQLListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterNamedArgumentList is called when production namedArgumentList is entered.
func (s *BaseDorisSQLListener) EnterNamedArgumentList(ctx *NamedArgumentListContext) {}

// ExitNamedArgumentList is called when production namedArgumentList is exited.
func (s *BaseDorisSQLListener) ExitNamedArgumentList(ctx *NamedArgumentListContext) {}

// EnterNamedArguments is called when production namedArguments is entered.
func (s *BaseDorisSQLListener) EnterNamedArguments(ctx *NamedArgumentsContext) {}

// ExitNamedArguments is called when production namedArguments is exited.
func (s *BaseDorisSQLListener) ExitNamedArguments(ctx *NamedArgumentsContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseDorisSQLListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseDorisSQLListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterCrossOrInnerJoinType is called when production crossOrInnerJoinType is entered.
func (s *BaseDorisSQLListener) EnterCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// ExitCrossOrInnerJoinType is called when production crossOrInnerJoinType is exited.
func (s *BaseDorisSQLListener) ExitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// EnterOuterAndSemiJoinType is called when production outerAndSemiJoinType is entered.
func (s *BaseDorisSQLListener) EnterOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// ExitOuterAndSemiJoinType is called when production outerAndSemiJoinType is exited.
func (s *BaseDorisSQLListener) ExitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// EnterBracketHint is called when production bracketHint is entered.
func (s *BaseDorisSQLListener) EnterBracketHint(ctx *BracketHintContext) {}

// ExitBracketHint is called when production bracketHint is exited.
func (s *BaseDorisSQLListener) ExitBracketHint(ctx *BracketHintContext) {}

// EnterHintMap is called when production hintMap is entered.
func (s *BaseDorisSQLListener) EnterHintMap(ctx *HintMapContext) {}

// ExitHintMap is called when production hintMap is exited.
func (s *BaseDorisSQLListener) ExitHintMap(ctx *HintMapContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseDorisSQLListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseDorisSQLListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseDorisSQLListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseDorisSQLListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterColumnAliasesWithoutParentheses is called when production columnAliasesWithoutParentheses is entered.
func (s *BaseDorisSQLListener) EnterColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) {
}

// ExitColumnAliasesWithoutParentheses is called when production columnAliasesWithoutParentheses is exited.
func (s *BaseDorisSQLListener) ExitColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) {
}

// EnterPartitionNames is called when production partitionNames is entered.
func (s *BaseDorisSQLListener) EnterPartitionNames(ctx *PartitionNamesContext) {}

// ExitPartitionNames is called when production partitionNames is exited.
func (s *BaseDorisSQLListener) ExitPartitionNames(ctx *PartitionNamesContext) {}

// EnterKeyPartitionList is called when production keyPartitionList is entered.
func (s *BaseDorisSQLListener) EnterKeyPartitionList(ctx *KeyPartitionListContext) {}

// ExitKeyPartitionList is called when production keyPartitionList is exited.
func (s *BaseDorisSQLListener) ExitKeyPartitionList(ctx *KeyPartitionListContext) {}

// EnterTabletList is called when production tabletList is entered.
func (s *BaseDorisSQLListener) EnterTabletList(ctx *TabletListContext) {}

// ExitTabletList is called when production tabletList is exited.
func (s *BaseDorisSQLListener) ExitTabletList(ctx *TabletListContext) {}

// EnterPrepareStatement is called when production prepareStatement is entered.
func (s *BaseDorisSQLListener) EnterPrepareStatement(ctx *PrepareStatementContext) {}

// ExitPrepareStatement is called when production prepareStatement is exited.
func (s *BaseDorisSQLListener) ExitPrepareStatement(ctx *PrepareStatementContext) {}

// EnterPrepareSql is called when production prepareSql is entered.
func (s *BaseDorisSQLListener) EnterPrepareSql(ctx *PrepareSqlContext) {}

// ExitPrepareSql is called when production prepareSql is exited.
func (s *BaseDorisSQLListener) ExitPrepareSql(ctx *PrepareSqlContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseDorisSQLListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseDorisSQLListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterDeallocateStatement is called when production deallocateStatement is entered.
func (s *BaseDorisSQLListener) EnterDeallocateStatement(ctx *DeallocateStatementContext) {}

// ExitDeallocateStatement is called when production deallocateStatement is exited.
func (s *BaseDorisSQLListener) ExitDeallocateStatement(ctx *DeallocateStatementContext) {}

// EnterReplicaList is called when production replicaList is entered.
func (s *BaseDorisSQLListener) EnterReplicaList(ctx *ReplicaListContext) {}

// ExitReplicaList is called when production replicaList is exited.
func (s *BaseDorisSQLListener) ExitReplicaList(ctx *ReplicaListContext) {}

// EnterExpressionsWithDefault is called when production expressionsWithDefault is entered.
func (s *BaseDorisSQLListener) EnterExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {}

// ExitExpressionsWithDefault is called when production expressionsWithDefault is exited.
func (s *BaseDorisSQLListener) ExitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {}

// EnterExpressionOrDefault is called when production expressionOrDefault is entered.
func (s *BaseDorisSQLListener) EnterExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// ExitExpressionOrDefault is called when production expressionOrDefault is exited.
func (s *BaseDorisSQLListener) ExitExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// EnterMapExpressionList is called when production mapExpressionList is entered.
func (s *BaseDorisSQLListener) EnterMapExpressionList(ctx *MapExpressionListContext) {}

// ExitMapExpressionList is called when production mapExpressionList is exited.
func (s *BaseDorisSQLListener) ExitMapExpressionList(ctx *MapExpressionListContext) {}

// EnterMapExpression is called when production mapExpression is entered.
func (s *BaseDorisSQLListener) EnterMapExpression(ctx *MapExpressionContext) {}

// ExitMapExpression is called when production mapExpression is exited.
func (s *BaseDorisSQLListener) ExitMapExpression(ctx *MapExpressionContext) {}

// EnterExpressionSingleton is called when production expressionSingleton is entered.
func (s *BaseDorisSQLListener) EnterExpressionSingleton(ctx *ExpressionSingletonContext) {}

// ExitExpressionSingleton is called when production expressionSingleton is exited.
func (s *BaseDorisSQLListener) ExitExpressionSingleton(ctx *ExpressionSingletonContext) {}

// EnterExpressionDefault is called when production expressionDefault is entered.
func (s *BaseDorisSQLListener) EnterExpressionDefault(ctx *ExpressionDefaultContext) {}

// ExitExpressionDefault is called when production expressionDefault is exited.
func (s *BaseDorisSQLListener) ExitExpressionDefault(ctx *ExpressionDefaultContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseDorisSQLListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseDorisSQLListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseDorisSQLListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseDorisSQLListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseDorisSQLListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseDorisSQLListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseDorisSQLListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseDorisSQLListener) ExitComparison(ctx *ComparisonContext) {}

// EnterBooleanExpressionDefault is called when production booleanExpressionDefault is entered.
func (s *BaseDorisSQLListener) EnterBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {}

// ExitBooleanExpressionDefault is called when production booleanExpressionDefault is exited.
func (s *BaseDorisSQLListener) ExitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {}

// EnterIsNull is called when production isNull is entered.
func (s *BaseDorisSQLListener) EnterIsNull(ctx *IsNullContext) {}

// ExitIsNull is called when production isNull is exited.
func (s *BaseDorisSQLListener) ExitIsNull(ctx *IsNullContext) {}

// EnterScalarSubquery is called when production scalarSubquery is entered.
func (s *BaseDorisSQLListener) EnterScalarSubquery(ctx *ScalarSubqueryContext) {}

// ExitScalarSubquery is called when production scalarSubquery is exited.
func (s *BaseDorisSQLListener) ExitScalarSubquery(ctx *ScalarSubqueryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseDorisSQLListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseDorisSQLListener) ExitPredicate(ctx *PredicateContext) {}

// EnterTupleInSubquery is called when production tupleInSubquery is entered.
func (s *BaseDorisSQLListener) EnterTupleInSubquery(ctx *TupleInSubqueryContext) {}

// ExitTupleInSubquery is called when production tupleInSubquery is exited.
func (s *BaseDorisSQLListener) ExitTupleInSubquery(ctx *TupleInSubqueryContext) {}

// EnterInSubquery is called when production inSubquery is entered.
func (s *BaseDorisSQLListener) EnterInSubquery(ctx *InSubqueryContext) {}

// ExitInSubquery is called when production inSubquery is exited.
func (s *BaseDorisSQLListener) ExitInSubquery(ctx *InSubqueryContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseDorisSQLListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseDorisSQLListener) ExitInList(ctx *InListContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseDorisSQLListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseDorisSQLListener) ExitBetween(ctx *BetweenContext) {}

// EnterLike is called when production like is entered.
func (s *BaseDorisSQLListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseDorisSQLListener) ExitLike(ctx *LikeContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseDorisSQLListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseDorisSQLListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseDorisSQLListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseDorisSQLListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseDorisSQLListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseDorisSQLListener) ExitDereference(ctx *DereferenceContext) {}

// EnterOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is entered.
func (s *BaseDorisSQLListener) EnterOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// ExitOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is exited.
func (s *BaseDorisSQLListener) ExitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// EnterMatchExpr is called when production matchExpr is entered.
func (s *BaseDorisSQLListener) EnterMatchExpr(ctx *MatchExprContext) {}

// ExitMatchExpr is called when production matchExpr is exited.
func (s *BaseDorisSQLListener) ExitMatchExpr(ctx *MatchExprContext) {}

// EnterColumnRef is called when production columnRef is entered.
func (s *BaseDorisSQLListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production columnRef is exited.
func (s *BaseDorisSQLListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterConvert is called when production convert is entered.
func (s *BaseDorisSQLListener) EnterConvert(ctx *ConvertContext) {}

// ExitConvert is called when production convert is exited.
func (s *BaseDorisSQLListener) ExitConvert(ctx *ConvertContext) {}

// EnterCollectionSubscript is called when production collectionSubscript is entered.
func (s *BaseDorisSQLListener) EnterCollectionSubscript(ctx *CollectionSubscriptContext) {}

// ExitCollectionSubscript is called when production collectionSubscript is exited.
func (s *BaseDorisSQLListener) ExitCollectionSubscript(ctx *CollectionSubscriptContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseDorisSQLListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseDorisSQLListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseDorisSQLListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseDorisSQLListener) ExitCast(ctx *CastContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseDorisSQLListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseDorisSQLListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterUserVariableExpression is called when production userVariableExpression is entered.
func (s *BaseDorisSQLListener) EnterUserVariableExpression(ctx *UserVariableExpressionContext) {}

// ExitUserVariableExpression is called when production userVariableExpression is exited.
func (s *BaseDorisSQLListener) ExitUserVariableExpression(ctx *UserVariableExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseDorisSQLListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseDorisSQLListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseDorisSQLListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseDorisSQLListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterArrowExpression is called when production arrowExpression is entered.
func (s *BaseDorisSQLListener) EnterArrowExpression(ctx *ArrowExpressionContext) {}

// ExitArrowExpression is called when production arrowExpression is exited.
func (s *BaseDorisSQLListener) ExitArrowExpression(ctx *ArrowExpressionContext) {}

// EnterArrayExpr is called when production arrayExpr is entered.
func (s *BaseDorisSQLListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production arrayExpr is exited.
func (s *BaseDorisSQLListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterSystemVariableExpression is called when production systemVariableExpression is entered.
func (s *BaseDorisSQLListener) EnterSystemVariableExpression(ctx *SystemVariableExpressionContext) {}

// ExitSystemVariableExpression is called when production systemVariableExpression is exited.
func (s *BaseDorisSQLListener) ExitSystemVariableExpression(ctx *SystemVariableExpressionContext) {}

// EnterConcat is called when production concat is entered.
func (s *BaseDorisSQLListener) EnterConcat(ctx *ConcatContext) {}

// ExitConcat is called when production concat is exited.
func (s *BaseDorisSQLListener) ExitConcat(ctx *ConcatContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseDorisSQLListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseDorisSQLListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterLambdaFunctionExpr is called when production lambdaFunctionExpr is entered.
func (s *BaseDorisSQLListener) EnterLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// ExitLambdaFunctionExpr is called when production lambdaFunctionExpr is exited.
func (s *BaseDorisSQLListener) ExitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// EnterDictionaryGetExpr is called when production dictionaryGetExpr is entered.
func (s *BaseDorisSQLListener) EnterDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// ExitDictionaryGetExpr is called when production dictionaryGetExpr is exited.
func (s *BaseDorisSQLListener) ExitDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// EnterCollate is called when production collate is entered.
func (s *BaseDorisSQLListener) EnterCollate(ctx *CollateContext) {}

// ExitCollate is called when production collate is exited.
func (s *BaseDorisSQLListener) ExitCollate(ctx *CollateContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseDorisSQLListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseDorisSQLListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterMapConstructor is called when production mapConstructor is entered.
func (s *BaseDorisSQLListener) EnterMapConstructor(ctx *MapConstructorContext) {}

// ExitMapConstructor is called when production mapConstructor is exited.
func (s *BaseDorisSQLListener) ExitMapConstructor(ctx *MapConstructorContext) {}

// EnterArraySlice is called when production arraySlice is entered.
func (s *BaseDorisSQLListener) EnterArraySlice(ctx *ArraySliceContext) {}

// ExitArraySlice is called when production arraySlice is exited.
func (s *BaseDorisSQLListener) ExitArraySlice(ctx *ArraySliceContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseDorisSQLListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseDorisSQLListener) ExitExists(ctx *ExistsContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseDorisSQLListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseDorisSQLListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseDorisSQLListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseDorisSQLListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseDorisSQLListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseDorisSQLListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseDorisSQLListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseDorisSQLListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseDorisSQLListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseDorisSQLListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseDorisSQLListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseDorisSQLListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseDorisSQLListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseDorisSQLListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseDorisSQLListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseDorisSQLListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterUnitBoundaryLiteral is called when production unitBoundaryLiteral is entered.
func (s *BaseDorisSQLListener) EnterUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// ExitUnitBoundaryLiteral is called when production unitBoundaryLiteral is exited.
func (s *BaseDorisSQLListener) ExitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// EnterBinaryLiteral is called when production binaryLiteral is entered.
func (s *BaseDorisSQLListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production binaryLiteral is exited.
func (s *BaseDorisSQLListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterParameter is called when production Parameter is entered.
func (s *BaseDorisSQLListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production Parameter is exited.
func (s *BaseDorisSQLListener) ExitParameter(ctx *ParameterContext) {}

// EnterExtract is called when production extract is entered.
func (s *BaseDorisSQLListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BaseDorisSQLListener) ExitExtract(ctx *ExtractContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseDorisSQLListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseDorisSQLListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterInformationFunction is called when production informationFunction is entered.
func (s *BaseDorisSQLListener) EnterInformationFunction(ctx *InformationFunctionContext) {}

// ExitInformationFunction is called when production informationFunction is exited.
func (s *BaseDorisSQLListener) ExitInformationFunction(ctx *InformationFunctionContext) {}

// EnterSpecialDateTime is called when production specialDateTime is entered.
func (s *BaseDorisSQLListener) EnterSpecialDateTime(ctx *SpecialDateTimeContext) {}

// ExitSpecialDateTime is called when production specialDateTime is exited.
func (s *BaseDorisSQLListener) ExitSpecialDateTime(ctx *SpecialDateTimeContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseDorisSQLListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseDorisSQLListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterAggregationFunctionCall is called when production aggregationFunctionCall is entered.
func (s *BaseDorisSQLListener) EnterAggregationFunctionCall(ctx *AggregationFunctionCallContext) {}

// ExitAggregationFunctionCall is called when production aggregationFunctionCall is exited.
func (s *BaseDorisSQLListener) ExitAggregationFunctionCall(ctx *AggregationFunctionCallContext) {}

// EnterWindowFunctionCall is called when production windowFunctionCall is entered.
func (s *BaseDorisSQLListener) EnterWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// ExitWindowFunctionCall is called when production windowFunctionCall is exited.
func (s *BaseDorisSQLListener) ExitWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// EnterTranslateFunctionCall is called when production translateFunctionCall is entered.
func (s *BaseDorisSQLListener) EnterTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// ExitTranslateFunctionCall is called when production translateFunctionCall is exited.
func (s *BaseDorisSQLListener) ExitTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// EnterSimpleFunctionCall is called when production simpleFunctionCall is entered.
func (s *BaseDorisSQLListener) EnterSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// ExitSimpleFunctionCall is called when production simpleFunctionCall is exited.
func (s *BaseDorisSQLListener) ExitSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseDorisSQLListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseDorisSQLListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterUserVariable is called when production userVariable is entered.
func (s *BaseDorisSQLListener) EnterUserVariable(ctx *UserVariableContext) {}

// ExitUserVariable is called when production userVariable is exited.
func (s *BaseDorisSQLListener) ExitUserVariable(ctx *UserVariableContext) {}

// EnterSystemVariable is called when production systemVariable is entered.
func (s *BaseDorisSQLListener) EnterSystemVariable(ctx *SystemVariableContext) {}

// ExitSystemVariable is called when production systemVariable is exited.
func (s *BaseDorisSQLListener) ExitSystemVariable(ctx *SystemVariableContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseDorisSQLListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseDorisSQLListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterInformationFunctionExpression is called when production informationFunctionExpression is entered.
func (s *BaseDorisSQLListener) EnterInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// ExitInformationFunctionExpression is called when production informationFunctionExpression is exited.
func (s *BaseDorisSQLListener) ExitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// EnterSpecialDateTimeExpression is called when production specialDateTimeExpression is entered.
func (s *BaseDorisSQLListener) EnterSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {
}

// ExitSpecialDateTimeExpression is called when production specialDateTimeExpression is exited.
func (s *BaseDorisSQLListener) ExitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {}

// EnterSpecialFunctionExpression is called when production specialFunctionExpression is entered.
func (s *BaseDorisSQLListener) EnterSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {
}

// ExitSpecialFunctionExpression is called when production specialFunctionExpression is exited.
func (s *BaseDorisSQLListener) ExitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {}

// EnterWindowFunction is called when production windowFunction is entered.
func (s *BaseDorisSQLListener) EnterWindowFunction(ctx *WindowFunctionContext) {}

// ExitWindowFunction is called when production windowFunction is exited.
func (s *BaseDorisSQLListener) ExitWindowFunction(ctx *WindowFunctionContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseDorisSQLListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseDorisSQLListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterOver is called when production over is entered.
func (s *BaseDorisSQLListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BaseDorisSQLListener) ExitOver(ctx *OverContext) {}

// EnterIgnoreNulls is called when production ignoreNulls is entered.
func (s *BaseDorisSQLListener) EnterIgnoreNulls(ctx *IgnoreNullsContext) {}

// ExitIgnoreNulls is called when production ignoreNulls is exited.
func (s *BaseDorisSQLListener) ExitIgnoreNulls(ctx *IgnoreNullsContext) {}

// EnterWindowFrame is called when production windowFrame is entered.
func (s *BaseDorisSQLListener) EnterWindowFrame(ctx *WindowFrameContext) {}

// ExitWindowFrame is called when production windowFrame is exited.
func (s *BaseDorisSQLListener) ExitWindowFrame(ctx *WindowFrameContext) {}

// EnterUnboundedFrame is called when production unboundedFrame is entered.
func (s *BaseDorisSQLListener) EnterUnboundedFrame(ctx *UnboundedFrameContext) {}

// ExitUnboundedFrame is called when production unboundedFrame is exited.
func (s *BaseDorisSQLListener) ExitUnboundedFrame(ctx *UnboundedFrameContext) {}

// EnterCurrentRowBound is called when production currentRowBound is entered.
func (s *BaseDorisSQLListener) EnterCurrentRowBound(ctx *CurrentRowBoundContext) {}

// ExitCurrentRowBound is called when production currentRowBound is exited.
func (s *BaseDorisSQLListener) ExitCurrentRowBound(ctx *CurrentRowBoundContext) {}

// EnterBoundedFrame is called when production boundedFrame is entered.
func (s *BaseDorisSQLListener) EnterBoundedFrame(ctx *BoundedFrameContext) {}

// ExitBoundedFrame is called when production boundedFrame is exited.
func (s *BaseDorisSQLListener) ExitBoundedFrame(ctx *BoundedFrameContext) {}

// EnterBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is entered.
func (s *BaseDorisSQLListener) EnterBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {}

// ExitBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is exited.
func (s *BaseDorisSQLListener) ExitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {}

// EnterTableDesc is called when production tableDesc is entered.
func (s *BaseDorisSQLListener) EnterTableDesc(ctx *TableDescContext) {}

// ExitTableDesc is called when production tableDesc is exited.
func (s *BaseDorisSQLListener) ExitTableDesc(ctx *TableDescContext) {}

// EnterBackupRestoreTableDesc is called when production backupRestoreTableDesc is entered.
func (s *BaseDorisSQLListener) EnterBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {}

// ExitBackupRestoreTableDesc is called when production backupRestoreTableDesc is exited.
func (s *BaseDorisSQLListener) ExitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {}

// EnterExplainDesc is called when production explainDesc is entered.
func (s *BaseDorisSQLListener) EnterExplainDesc(ctx *ExplainDescContext) {}

// ExitExplainDesc is called when production explainDesc is exited.
func (s *BaseDorisSQLListener) ExitExplainDesc(ctx *ExplainDescContext) {}

// EnterOptimizerTrace is called when production optimizerTrace is entered.
func (s *BaseDorisSQLListener) EnterOptimizerTrace(ctx *OptimizerTraceContext) {}

// ExitOptimizerTrace is called when production optimizerTrace is exited.
func (s *BaseDorisSQLListener) ExitOptimizerTrace(ctx *OptimizerTraceContext) {}

// EnterPartitionExpr is called when production partitionExpr is entered.
func (s *BaseDorisSQLListener) EnterPartitionExpr(ctx *PartitionExprContext) {}

// ExitPartitionExpr is called when production partitionExpr is exited.
func (s *BaseDorisSQLListener) ExitPartitionExpr(ctx *PartitionExprContext) {}

// EnterPartitionDesc is called when production partitionDesc is entered.
func (s *BaseDorisSQLListener) EnterPartitionDesc(ctx *PartitionDescContext) {}

// ExitPartitionDesc is called when production partitionDesc is exited.
func (s *BaseDorisSQLListener) ExitPartitionDesc(ctx *PartitionDescContext) {}

// EnterListPartitionDesc is called when production listPartitionDesc is entered.
func (s *BaseDorisSQLListener) EnterListPartitionDesc(ctx *ListPartitionDescContext) {}

// ExitListPartitionDesc is called when production listPartitionDesc is exited.
func (s *BaseDorisSQLListener) ExitListPartitionDesc(ctx *ListPartitionDescContext) {}

// EnterSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is entered.
func (s *BaseDorisSQLListener) EnterSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// ExitSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is exited.
func (s *BaseDorisSQLListener) ExitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// EnterMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is entered.
func (s *BaseDorisSQLListener) EnterMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// ExitMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is exited.
func (s *BaseDorisSQLListener) ExitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// EnterMultiListPartitionValues is called when production multiListPartitionValues is entered.
func (s *BaseDorisSQLListener) EnterMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {}

// ExitMultiListPartitionValues is called when production multiListPartitionValues is exited.
func (s *BaseDorisSQLListener) ExitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {}

// EnterSingleListPartitionValues is called when production singleListPartitionValues is entered.
func (s *BaseDorisSQLListener) EnterSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {
}

// ExitSingleListPartitionValues is called when production singleListPartitionValues is exited.
func (s *BaseDorisSQLListener) ExitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {}

// EnterListPartitionValues is called when production listPartitionValues is entered.
func (s *BaseDorisSQLListener) EnterListPartitionValues(ctx *ListPartitionValuesContext) {}

// ExitListPartitionValues is called when production listPartitionValues is exited.
func (s *BaseDorisSQLListener) ExitListPartitionValues(ctx *ListPartitionValuesContext) {}

// EnterListPartitionValue is called when production listPartitionValue is entered.
func (s *BaseDorisSQLListener) EnterListPartitionValue(ctx *ListPartitionValueContext) {}

// ExitListPartitionValue is called when production listPartitionValue is exited.
func (s *BaseDorisSQLListener) ExitListPartitionValue(ctx *ListPartitionValueContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseDorisSQLListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseDorisSQLListener) ExitStringList(ctx *StringListContext) {}

// EnterLiteralExpressionList is called when production literalExpressionList is entered.
func (s *BaseDorisSQLListener) EnterLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// ExitLiteralExpressionList is called when production literalExpressionList is exited.
func (s *BaseDorisSQLListener) ExitLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// EnterRangePartitionDesc is called when production rangePartitionDesc is entered.
func (s *BaseDorisSQLListener) EnterRangePartitionDesc(ctx *RangePartitionDescContext) {}

// ExitRangePartitionDesc is called when production rangePartitionDesc is exited.
func (s *BaseDorisSQLListener) ExitRangePartitionDesc(ctx *RangePartitionDescContext) {}

// EnterSingleRangePartition is called when production singleRangePartition is entered.
func (s *BaseDorisSQLListener) EnterSingleRangePartition(ctx *SingleRangePartitionContext) {}

// ExitSingleRangePartition is called when production singleRangePartition is exited.
func (s *BaseDorisSQLListener) ExitSingleRangePartition(ctx *SingleRangePartitionContext) {}

// EnterMultiRangePartition is called when production multiRangePartition is entered.
func (s *BaseDorisSQLListener) EnterMultiRangePartition(ctx *MultiRangePartitionContext) {}

// ExitMultiRangePartition is called when production multiRangePartition is exited.
func (s *BaseDorisSQLListener) ExitMultiRangePartition(ctx *MultiRangePartitionContext) {}

// EnterPartitionRangeDesc is called when production partitionRangeDesc is entered.
func (s *BaseDorisSQLListener) EnterPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// ExitPartitionRangeDesc is called when production partitionRangeDesc is exited.
func (s *BaseDorisSQLListener) ExitPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// EnterPartitionKeyDesc is called when production partitionKeyDesc is entered.
func (s *BaseDorisSQLListener) EnterPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// ExitPartitionKeyDesc is called when production partitionKeyDesc is exited.
func (s *BaseDorisSQLListener) ExitPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// EnterPartitionValueList is called when production partitionValueList is entered.
func (s *BaseDorisSQLListener) EnterPartitionValueList(ctx *PartitionValueListContext) {}

// ExitPartitionValueList is called when production partitionValueList is exited.
func (s *BaseDorisSQLListener) ExitPartitionValueList(ctx *PartitionValueListContext) {}

// EnterKeyPartition is called when production keyPartition is entered.
func (s *BaseDorisSQLListener) EnterKeyPartition(ctx *KeyPartitionContext) {}

// ExitKeyPartition is called when production keyPartition is exited.
func (s *BaseDorisSQLListener) ExitKeyPartition(ctx *KeyPartitionContext) {}

// EnterPartitionValue is called when production partitionValue is entered.
func (s *BaseDorisSQLListener) EnterPartitionValue(ctx *PartitionValueContext) {}

// ExitPartitionValue is called when production partitionValue is exited.
func (s *BaseDorisSQLListener) ExitPartitionValue(ctx *PartitionValueContext) {}

// EnterDistributionClause is called when production distributionClause is entered.
func (s *BaseDorisSQLListener) EnterDistributionClause(ctx *DistributionClauseContext) {}

// ExitDistributionClause is called when production distributionClause is exited.
func (s *BaseDorisSQLListener) ExitDistributionClause(ctx *DistributionClauseContext) {}

// EnterDistributionDesc is called when production distributionDesc is entered.
func (s *BaseDorisSQLListener) EnterDistributionDesc(ctx *DistributionDescContext) {}

// ExitDistributionDesc is called when production distributionDesc is exited.
func (s *BaseDorisSQLListener) ExitDistributionDesc(ctx *DistributionDescContext) {}

// EnterRefreshSchemeDesc is called when production refreshSchemeDesc is entered.
func (s *BaseDorisSQLListener) EnterRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// ExitRefreshSchemeDesc is called when production refreshSchemeDesc is exited.
func (s *BaseDorisSQLListener) ExitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// EnterStatusDesc is called when production statusDesc is entered.
func (s *BaseDorisSQLListener) EnterStatusDesc(ctx *StatusDescContext) {}

// ExitStatusDesc is called when production statusDesc is exited.
func (s *BaseDorisSQLListener) ExitStatusDesc(ctx *StatusDescContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseDorisSQLListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseDorisSQLListener) ExitProperties(ctx *PropertiesContext) {}

// EnterExtProperties is called when production extProperties is entered.
func (s *BaseDorisSQLListener) EnterExtProperties(ctx *ExtPropertiesContext) {}

// ExitExtProperties is called when production extProperties is exited.
func (s *BaseDorisSQLListener) ExitExtProperties(ctx *ExtPropertiesContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseDorisSQLListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseDorisSQLListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterUserPropertyList is called when production userPropertyList is entered.
func (s *BaseDorisSQLListener) EnterUserPropertyList(ctx *UserPropertyListContext) {}

// ExitUserPropertyList is called when production userPropertyList is exited.
func (s *BaseDorisSQLListener) ExitUserPropertyList(ctx *UserPropertyListContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseDorisSQLListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseDorisSQLListener) ExitProperty(ctx *PropertyContext) {}

// EnterInlineProperties is called when production inlineProperties is entered.
func (s *BaseDorisSQLListener) EnterInlineProperties(ctx *InlinePropertiesContext) {}

// ExitInlineProperties is called when production inlineProperties is exited.
func (s *BaseDorisSQLListener) ExitInlineProperties(ctx *InlinePropertiesContext) {}

// EnterInlineProperty is called when production inlineProperty is entered.
func (s *BaseDorisSQLListener) EnterInlineProperty(ctx *InlinePropertyContext) {}

// ExitInlineProperty is called when production inlineProperty is exited.
func (s *BaseDorisSQLListener) ExitInlineProperty(ctx *InlinePropertyContext) {}

// EnterVarType is called when production varType is entered.
func (s *BaseDorisSQLListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseDorisSQLListener) ExitVarType(ctx *VarTypeContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseDorisSQLListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseDorisSQLListener) ExitComment(ctx *CommentContext) {}

// EnterOutfile is called when production outfile is entered.
func (s *BaseDorisSQLListener) EnterOutfile(ctx *OutfileContext) {}

// ExitOutfile is called when production outfile is exited.
func (s *BaseDorisSQLListener) ExitOutfile(ctx *OutfileContext) {}

// EnterFileFormat is called when production fileFormat is entered.
func (s *BaseDorisSQLListener) EnterFileFormat(ctx *FileFormatContext) {}

// ExitFileFormat is called when production fileFormat is exited.
func (s *BaseDorisSQLListener) ExitFileFormat(ctx *FileFormatContext) {}

// EnterString is called when production string is entered.
func (s *BaseDorisSQLListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseDorisSQLListener) ExitString(ctx *StringContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseDorisSQLListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseDorisSQLListener) ExitBinary(ctx *BinaryContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseDorisSQLListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseDorisSQLListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseDorisSQLListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseDorisSQLListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseDorisSQLListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseDorisSQLListener) ExitInterval(ctx *IntervalContext) {}

// EnterTaskInterval is called when production taskInterval is entered.
func (s *BaseDorisSQLListener) EnterTaskInterval(ctx *TaskIntervalContext) {}

// ExitTaskInterval is called when production taskInterval is exited.
func (s *BaseDorisSQLListener) ExitTaskInterval(ctx *TaskIntervalContext) {}

// EnterTaskUnitIdentifier is called when production taskUnitIdentifier is entered.
func (s *BaseDorisSQLListener) EnterTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// ExitTaskUnitIdentifier is called when production taskUnitIdentifier is exited.
func (s *BaseDorisSQLListener) ExitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// EnterUnitIdentifier is called when production unitIdentifier is entered.
func (s *BaseDorisSQLListener) EnterUnitIdentifier(ctx *UnitIdentifierContext) {}

// ExitUnitIdentifier is called when production unitIdentifier is exited.
func (s *BaseDorisSQLListener) ExitUnitIdentifier(ctx *UnitIdentifierContext) {}

// EnterUnitBoundary is called when production unitBoundary is entered.
func (s *BaseDorisSQLListener) EnterUnitBoundary(ctx *UnitBoundaryContext) {}

// ExitUnitBoundary is called when production unitBoundary is exited.
func (s *BaseDorisSQLListener) ExitUnitBoundary(ctx *UnitBoundaryContext) {}

// EnterType is called when production type is entered.
func (s *BaseDorisSQLListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseDorisSQLListener) ExitType(ctx *TypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseDorisSQLListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseDorisSQLListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseDorisSQLListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseDorisSQLListener) ExitMapType(ctx *MapTypeContext) {}

// EnterSubfieldDesc is called when production subfieldDesc is entered.
func (s *BaseDorisSQLListener) EnterSubfieldDesc(ctx *SubfieldDescContext) {}

// ExitSubfieldDesc is called when production subfieldDesc is exited.
func (s *BaseDorisSQLListener) ExitSubfieldDesc(ctx *SubfieldDescContext) {}

// EnterSubfieldDescs is called when production subfieldDescs is entered.
func (s *BaseDorisSQLListener) EnterSubfieldDescs(ctx *SubfieldDescsContext) {}

// ExitSubfieldDescs is called when production subfieldDescs is exited.
func (s *BaseDorisSQLListener) ExitSubfieldDescs(ctx *SubfieldDescsContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseDorisSQLListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseDorisSQLListener) ExitStructType(ctx *StructTypeContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseDorisSQLListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseDorisSQLListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseDorisSQLListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseDorisSQLListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterDecimalType is called when production decimalType is entered.
func (s *BaseDorisSQLListener) EnterDecimalType(ctx *DecimalTypeContext) {}

// ExitDecimalType is called when production decimalType is exited.
func (s *BaseDorisSQLListener) ExitDecimalType(ctx *DecimalTypeContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseDorisSQLListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseDorisSQLListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseDorisSQLListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseDorisSQLListener) ExitTableName(ctx *TableNameContext) {}

// EnterWriteBranch is called when production writeBranch is entered.
func (s *BaseDorisSQLListener) EnterWriteBranch(ctx *WriteBranchContext) {}

// ExitWriteBranch is called when production writeBranch is exited.
func (s *BaseDorisSQLListener) ExitWriteBranch(ctx *WriteBranchContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseDorisSQLListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseDorisSQLListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterDigitIdentifier is called when production digitIdentifier is entered.
func (s *BaseDorisSQLListener) EnterDigitIdentifier(ctx *DigitIdentifierContext) {}

// ExitDigitIdentifier is called when production digitIdentifier is exited.
func (s *BaseDorisSQLListener) ExitDigitIdentifier(ctx *DigitIdentifierContext) {}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseDorisSQLListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseDorisSQLListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterIdentifierWithAlias is called when production identifierWithAlias is entered.
func (s *BaseDorisSQLListener) EnterIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// ExitIdentifierWithAlias is called when production identifierWithAlias is exited.
func (s *BaseDorisSQLListener) ExitIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// EnterIdentifierWithAliasList is called when production identifierWithAliasList is entered.
func (s *BaseDorisSQLListener) EnterIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {}

// ExitIdentifierWithAliasList is called when production identifierWithAliasList is exited.
func (s *BaseDorisSQLListener) ExitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseDorisSQLListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseDorisSQLListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierOrString is called when production identifierOrString is entered.
func (s *BaseDorisSQLListener) EnterIdentifierOrString(ctx *IdentifierOrStringContext) {}

// ExitIdentifierOrString is called when production identifierOrString is exited.
func (s *BaseDorisSQLListener) ExitIdentifierOrString(ctx *IdentifierOrStringContext) {}

// EnterIdentifierOrStringList is called when production identifierOrStringList is entered.
func (s *BaseDorisSQLListener) EnterIdentifierOrStringList(ctx *IdentifierOrStringListContext) {}

// ExitIdentifierOrStringList is called when production identifierOrStringList is exited.
func (s *BaseDorisSQLListener) ExitIdentifierOrStringList(ctx *IdentifierOrStringListContext) {}

// EnterIdentifierOrStringOrStar is called when production identifierOrStringOrStar is entered.
func (s *BaseDorisSQLListener) EnterIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {}

// ExitIdentifierOrStringOrStar is called when production identifierOrStringOrStar is exited.
func (s *BaseDorisSQLListener) ExitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {}

// EnterUserWithoutHost is called when production userWithoutHost is entered.
func (s *BaseDorisSQLListener) EnterUserWithoutHost(ctx *UserWithoutHostContext) {}

// ExitUserWithoutHost is called when production userWithoutHost is exited.
func (s *BaseDorisSQLListener) ExitUserWithoutHost(ctx *UserWithoutHostContext) {}

// EnterUserWithHost is called when production userWithHost is entered.
func (s *BaseDorisSQLListener) EnterUserWithHost(ctx *UserWithHostContext) {}

// ExitUserWithHost is called when production userWithHost is exited.
func (s *BaseDorisSQLListener) ExitUserWithHost(ctx *UserWithHostContext) {}

// EnterUserWithHostAndBlanket is called when production userWithHostAndBlanket is entered.
func (s *BaseDorisSQLListener) EnterUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {}

// ExitUserWithHostAndBlanket is called when production userWithHostAndBlanket is exited.
func (s *BaseDorisSQLListener) ExitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseDorisSQLListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseDorisSQLListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseDorisSQLListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseDorisSQLListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterDecimalValue is called when production decimalValue is entered.
func (s *BaseDorisSQLListener) EnterDecimalValue(ctx *DecimalValueContext) {}

// ExitDecimalValue is called when production decimalValue is exited.
func (s *BaseDorisSQLListener) ExitDecimalValue(ctx *DecimalValueContext) {}

// EnterDoubleValue is called when production doubleValue is entered.
func (s *BaseDorisSQLListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production doubleValue is exited.
func (s *BaseDorisSQLListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseDorisSQLListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseDorisSQLListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseDorisSQLListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseDorisSQLListener) ExitNonReserved(ctx *NonReservedContext) {}
