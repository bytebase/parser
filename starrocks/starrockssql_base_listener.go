// Code generated from StarRocksSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package starrocks // StarRocksSQL
import "github.com/antlr4-go/antlr/v4"

// BaseStarRocksSQLListener is a complete listener for a parse tree produced by StarRocksSQLParser.
type BaseStarRocksSQLListener struct{}

var _ StarRocksSQLListener = &BaseStarRocksSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStarRocksSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStarRocksSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStarRocksSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStarRocksSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSqlStatements is called when production sqlStatements is entered.
func (s *BaseStarRocksSQLListener) EnterSqlStatements(ctx *SqlStatementsContext) {}

// ExitSqlStatements is called when production sqlStatements is exited.
func (s *BaseStarRocksSQLListener) ExitSqlStatements(ctx *SqlStatementsContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseStarRocksSQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseStarRocksSQLListener) ExitStatement(ctx *StatementContext) {}

// EnterUseDatabaseStatement is called when production useDatabaseStatement is entered.
func (s *BaseStarRocksSQLListener) EnterUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// ExitUseDatabaseStatement is called when production useDatabaseStatement is exited.
func (s *BaseStarRocksSQLListener) ExitUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// EnterUseCatalogStatement is called when production useCatalogStatement is entered.
func (s *BaseStarRocksSQLListener) EnterUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// ExitUseCatalogStatement is called when production useCatalogStatement is exited.
func (s *BaseStarRocksSQLListener) ExitUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// EnterSetCatalogStatement is called when production setCatalogStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// ExitSetCatalogStatement is called when production setCatalogStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// EnterShowDatabasesStatement is called when production showDatabasesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {}

// ExitShowDatabasesStatement is called when production showDatabasesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {}

// EnterAlterDbQuotaStatement is called when production alterDbQuotaStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// ExitAlterDbQuotaStatement is called when production alterDbQuotaStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// EnterCreateDbStatement is called when production createDbStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateDbStatement(ctx *CreateDbStatementContext) {}

// ExitCreateDbStatement is called when production createDbStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateDbStatement(ctx *CreateDbStatementContext) {}

// EnterDropDbStatement is called when production dropDbStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropDbStatement(ctx *DropDbStatementContext) {}

// ExitDropDbStatement is called when production dropDbStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropDbStatement(ctx *DropDbStatementContext) {}

// EnterShowCreateDbStatement is called when production showCreateDbStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// ExitShowCreateDbStatement is called when production showCreateDbStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// EnterAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// ExitAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// EnterRecoverDbStmt is called when production recoverDbStmt is entered.
func (s *BaseStarRocksSQLListener) EnterRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// ExitRecoverDbStmt is called when production recoverDbStmt is exited.
func (s *BaseStarRocksSQLListener) ExitRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// EnterShowDataStmt is called when production showDataStmt is entered.
func (s *BaseStarRocksSQLListener) EnterShowDataStmt(ctx *ShowDataStmtContext) {}

// ExitShowDataStmt is called when production showDataStmt is exited.
func (s *BaseStarRocksSQLListener) ExitShowDataStmt(ctx *ShowDataStmtContext) {}

// EnterShowDataDistributionStmt is called when production showDataDistributionStmt is entered.
func (s *BaseStarRocksSQLListener) EnterShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {
}

// ExitShowDataDistributionStmt is called when production showDataDistributionStmt is exited.
func (s *BaseStarRocksSQLListener) ExitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {
}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterColumnDesc is called when production columnDesc is entered.
func (s *BaseStarRocksSQLListener) EnterColumnDesc(ctx *ColumnDescContext) {}

// ExitColumnDesc is called when production columnDesc is exited.
func (s *BaseStarRocksSQLListener) ExitColumnDesc(ctx *ColumnDescContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseStarRocksSQLListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseStarRocksSQLListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterDefaultDesc is called when production defaultDesc is entered.
func (s *BaseStarRocksSQLListener) EnterDefaultDesc(ctx *DefaultDescContext) {}

// ExitDefaultDesc is called when production defaultDesc is exited.
func (s *BaseStarRocksSQLListener) ExitDefaultDesc(ctx *DefaultDescContext) {}

// EnterGeneratedColumnDesc is called when production generatedColumnDesc is entered.
func (s *BaseStarRocksSQLListener) EnterGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// ExitGeneratedColumnDesc is called when production generatedColumnDesc is exited.
func (s *BaseStarRocksSQLListener) ExitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// EnterIndexDesc is called when production indexDesc is entered.
func (s *BaseStarRocksSQLListener) EnterIndexDesc(ctx *IndexDescContext) {}

// ExitIndexDesc is called when production indexDesc is exited.
func (s *BaseStarRocksSQLListener) ExitIndexDesc(ctx *IndexDescContext) {}

// EnterEngineDesc is called when production engineDesc is entered.
func (s *BaseStarRocksSQLListener) EnterEngineDesc(ctx *EngineDescContext) {}

// ExitEngineDesc is called when production engineDesc is exited.
func (s *BaseStarRocksSQLListener) ExitEngineDesc(ctx *EngineDescContext) {}

// EnterCharsetDesc is called when production charsetDesc is entered.
func (s *BaseStarRocksSQLListener) EnterCharsetDesc(ctx *CharsetDescContext) {}

// ExitCharsetDesc is called when production charsetDesc is exited.
func (s *BaseStarRocksSQLListener) ExitCharsetDesc(ctx *CharsetDescContext) {}

// EnterCollateDesc is called when production collateDesc is entered.
func (s *BaseStarRocksSQLListener) EnterCollateDesc(ctx *CollateDescContext) {}

// ExitCollateDesc is called when production collateDesc is exited.
func (s *BaseStarRocksSQLListener) ExitCollateDesc(ctx *CollateDescContext) {}

// EnterKeyDesc is called when production keyDesc is entered.
func (s *BaseStarRocksSQLListener) EnterKeyDesc(ctx *KeyDescContext) {}

// ExitKeyDesc is called when production keyDesc is exited.
func (s *BaseStarRocksSQLListener) ExitKeyDesc(ctx *KeyDescContext) {}

// EnterOrderByDesc is called when production orderByDesc is entered.
func (s *BaseStarRocksSQLListener) EnterOrderByDesc(ctx *OrderByDescContext) {}

// ExitOrderByDesc is called when production orderByDesc is exited.
func (s *BaseStarRocksSQLListener) ExitOrderByDesc(ctx *OrderByDescContext) {}

// EnterColumnNullable is called when production columnNullable is entered.
func (s *BaseStarRocksSQLListener) EnterColumnNullable(ctx *ColumnNullableContext) {}

// ExitColumnNullable is called when production columnNullable is exited.
func (s *BaseStarRocksSQLListener) ExitColumnNullable(ctx *ColumnNullableContext) {}

// EnterTypeWithNullable is called when production typeWithNullable is entered.
func (s *BaseStarRocksSQLListener) EnterTypeWithNullable(ctx *TypeWithNullableContext) {}

// ExitTypeWithNullable is called when production typeWithNullable is exited.
func (s *BaseStarRocksSQLListener) ExitTypeWithNullable(ctx *TypeWithNullableContext) {}

// EnterAggStateDesc is called when production aggStateDesc is entered.
func (s *BaseStarRocksSQLListener) EnterAggStateDesc(ctx *AggStateDescContext) {}

// ExitAggStateDesc is called when production aggStateDesc is exited.
func (s *BaseStarRocksSQLListener) ExitAggStateDesc(ctx *AggStateDescContext) {}

// EnterAggDesc is called when production aggDesc is entered.
func (s *BaseStarRocksSQLListener) EnterAggDesc(ctx *AggDescContext) {}

// ExitAggDesc is called when production aggDesc is exited.
func (s *BaseStarRocksSQLListener) ExitAggDesc(ctx *AggDescContext) {}

// EnterRollupDesc is called when production rollupDesc is entered.
func (s *BaseStarRocksSQLListener) EnterRollupDesc(ctx *RollupDescContext) {}

// ExitRollupDesc is called when production rollupDesc is exited.
func (s *BaseStarRocksSQLListener) ExitRollupDesc(ctx *RollupDescContext) {}

// EnterRollupItem is called when production rollupItem is entered.
func (s *BaseStarRocksSQLListener) EnterRollupItem(ctx *RollupItemContext) {}

// ExitRollupItem is called when production rollupItem is exited.
func (s *BaseStarRocksSQLListener) ExitRollupItem(ctx *RollupItemContext) {}

// EnterDupKeys is called when production dupKeys is entered.
func (s *BaseStarRocksSQLListener) EnterDupKeys(ctx *DupKeysContext) {}

// ExitDupKeys is called when production dupKeys is exited.
func (s *BaseStarRocksSQLListener) ExitDupKeys(ctx *DupKeysContext) {}

// EnterFromRollup is called when production fromRollup is entered.
func (s *BaseStarRocksSQLListener) EnterFromRollup(ctx *FromRollupContext) {}

// ExitFromRollup is called when production fromRollup is exited.
func (s *BaseStarRocksSQLListener) ExitFromRollup(ctx *FromRollupContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseStarRocksSQLListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseStarRocksSQLListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseStarRocksSQLListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseStarRocksSQLListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterCreateTableAsSelectStatement is called when production createTableAsSelectStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// ExitCreateTableAsSelectStatement is called when production createTableAsSelectStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// ExitCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterCreateIndexStatement is called when production createIndexStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// ExitCreateIndexStatement is called when production createIndexStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// EnterDropIndexStatement is called when production dropIndexStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropIndexStatement(ctx *DropIndexStatementContext) {}

// ExitDropIndexStatement is called when production dropIndexStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropIndexStatement(ctx *DropIndexStatementContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseStarRocksSQLListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseStarRocksSQLListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterShowTableStatement is called when production showTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTableStatement(ctx *ShowTableStatementContext) {}

// ExitShowTableStatement is called when production showTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTableStatement(ctx *ShowTableStatementContext) {}

// EnterShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// ExitShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// EnterShowCreateTableStatement is called when production showCreateTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {
}

// ExitShowCreateTableStatement is called when production showCreateTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {
}

// EnterShowColumnStatement is called when production showColumnStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowColumnStatement(ctx *ShowColumnStatementContext) {}

// ExitShowColumnStatement is called when production showColumnStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowColumnStatement(ctx *ShowColumnStatementContext) {}

// EnterShowTableStatusStatement is called when production showTableStatusStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {
}

// ExitShowTableStatusStatement is called when production showTableStatusStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {
}

// EnterRefreshTableStatement is called when production refreshTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// ExitRefreshTableStatement is called when production refreshTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// EnterShowAlterStatement is called when production showAlterStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowAlterStatement(ctx *ShowAlterStatementContext) {}

// ExitShowAlterStatement is called when production showAlterStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowAlterStatement(ctx *ShowAlterStatementContext) {}

// EnterDescTableStatement is called when production descTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDescTableStatement(ctx *DescTableStatementContext) {}

// ExitDescTableStatement is called when production descTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDescTableStatement(ctx *DescTableStatementContext) {}

// EnterCreateTableLikeStatement is called when production createTableLikeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {
}

// ExitCreateTableLikeStatement is called when production createTableLikeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {
}

// EnterShowIndexStatement is called when production showIndexStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowIndexStatement(ctx *ShowIndexStatementContext) {}

// ExitShowIndexStatement is called when production showIndexStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowIndexStatement(ctx *ShowIndexStatementContext) {}

// EnterRecoverTableStatement is called when production recoverTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// ExitRecoverTableStatement is called when production recoverTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// EnterTruncateTableStatement is called when production truncateTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// ExitTruncateTableStatement is called when production truncateTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// EnterCancelAlterTableStatement is called when production cancelAlterTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {
}

// ExitCancelAlterTableStatement is called when production cancelAlterTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {
}

// EnterShowPartitionsStatement is called when production showPartitionsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {
}

// ExitShowPartitionsStatement is called when production showPartitionsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {}

// EnterRecoverPartitionStatement is called when production recoverPartitionStatement is entered.
func (s *BaseStarRocksSQLListener) EnterRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {
}

// ExitRecoverPartitionStatement is called when production recoverPartitionStatement is exited.
func (s *BaseStarRocksSQLListener) ExitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {
}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateViewStatement(ctx *CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateViewStatement(ctx *CreateViewStatementContext) {}

// EnterAlterViewStatement is called when production alterViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterViewStatement(ctx *AlterViewStatementContext) {}

// ExitAlterViewStatement is called when production alterViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterViewStatement(ctx *AlterViewStatementContext) {}

// EnterDropViewStatement is called when production dropViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropViewStatement(ctx *DropViewStatementContext) {}

// ExitDropViewStatement is called when production dropViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropViewStatement(ctx *DropViewStatementContext) {}

// EnterColumnNameWithComment is called when production columnNameWithComment is entered.
func (s *BaseStarRocksSQLListener) EnterColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// ExitColumnNameWithComment is called when production columnNameWithComment is exited.
func (s *BaseStarRocksSQLListener) ExitColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// EnterSubmitTaskStatement is called when production submitTaskStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// ExitSubmitTaskStatement is called when production submitTaskStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// EnterTaskClause is called when production taskClause is entered.
func (s *BaseStarRocksSQLListener) EnterTaskClause(ctx *TaskClauseContext) {}

// ExitTaskClause is called when production taskClause is exited.
func (s *BaseStarRocksSQLListener) ExitTaskClause(ctx *TaskClauseContext) {}

// EnterDropTaskStatement is called when production dropTaskStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropTaskStatement(ctx *DropTaskStatementContext) {}

// ExitDropTaskStatement is called when production dropTaskStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropTaskStatement(ctx *DropTaskStatementContext) {}

// EnterTaskScheduleDesc is called when production taskScheduleDesc is entered.
func (s *BaseStarRocksSQLListener) EnterTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// ExitTaskScheduleDesc is called when production taskScheduleDesc is exited.
func (s *BaseStarRocksSQLListener) ExitTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// EnterMvPartitionExprs is called when production mvPartitionExprs is entered.
func (s *BaseStarRocksSQLListener) EnterMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// ExitMvPartitionExprs is called when production mvPartitionExprs is exited.
func (s *BaseStarRocksSQLListener) ExitMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// EnterMaterializedViewDesc is called when production materializedViewDesc is entered.
func (s *BaseStarRocksSQLListener) EnterMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// ExitMaterializedViewDesc is called when production materializedViewDesc is exited.
func (s *BaseStarRocksSQLListener) ExitMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// EnterShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// ExitShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// EnterAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// ExitAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// EnterRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// ExitRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// EnterCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// ExitCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// EnterAdminSetConfigStatement is called when production adminSetConfigStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {
}

// ExitAdminSetConfigStatement is called when production adminSetConfigStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {}

// EnterAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// ExitAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// EnterAdminShowConfigStatement is called when production adminShowConfigStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {
}

// ExitAdminShowConfigStatement is called when production adminShowConfigStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {
}

// EnterAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// ExitAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// EnterAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// ExitAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// EnterAdminRepairTableStatement is called when production adminRepairTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {
}

// ExitAdminRepairTableStatement is called when production adminRepairTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {
}

// EnterAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// ExitAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// EnterAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// ExitAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// EnterAdminSetPartitionVersion is called when production adminSetPartitionVersion is entered.
func (s *BaseStarRocksSQLListener) EnterAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {
}

// ExitAdminSetPartitionVersion is called when production adminSetPartitionVersion is exited.
func (s *BaseStarRocksSQLListener) ExitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {
}

// EnterKillStatement is called when production killStatement is entered.
func (s *BaseStarRocksSQLListener) EnterKillStatement(ctx *KillStatementContext) {}

// ExitKillStatement is called when production killStatement is exited.
func (s *BaseStarRocksSQLListener) ExitKillStatement(ctx *KillStatementContext) {}

// EnterSyncStatement is called when production syncStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSyncStatement(ctx *SyncStatementContext) {}

// ExitSyncStatement is called when production syncStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSyncStatement(ctx *SyncStatementContext) {}

// EnterAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// EnterAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// EnterAlterSystemStatement is called when production alterSystemStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// ExitAlterSystemStatement is called when production alterSystemStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// EnterCancelAlterSystemStatement is called when production cancelAlterSystemStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// ExitCancelAlterSystemStatement is called when production cancelAlterSystemStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// EnterShowComputeNodesStatement is called when production showComputeNodesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {
}

// ExitShowComputeNodesStatement is called when production showComputeNodesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {
}

// EnterCreateExternalCatalogStatement is called when production createExternalCatalogStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// ExitCreateExternalCatalogStatement is called when production createExternalCatalogStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// EnterShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// ExitShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// EnterDropExternalCatalogStatement is called when production dropExternalCatalogStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// ExitDropExternalCatalogStatement is called when production dropExternalCatalogStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// EnterShowCatalogsStatement is called when production showCatalogsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// ExitShowCatalogsStatement is called when production showCatalogsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// EnterAlterCatalogStatement is called when production alterCatalogStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// ExitAlterCatalogStatement is called when production alterCatalogStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// EnterCreateStorageVolumeStatement is called when production createStorageVolumeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// ExitCreateStorageVolumeStatement is called when production createStorageVolumeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// EnterTypeDesc is called when production typeDesc is entered.
func (s *BaseStarRocksSQLListener) EnterTypeDesc(ctx *TypeDescContext) {}

// ExitTypeDesc is called when production typeDesc is exited.
func (s *BaseStarRocksSQLListener) ExitTypeDesc(ctx *TypeDescContext) {}

// EnterLocationsDesc is called when production locationsDesc is entered.
func (s *BaseStarRocksSQLListener) EnterLocationsDesc(ctx *LocationsDescContext) {}

// ExitLocationsDesc is called when production locationsDesc is exited.
func (s *BaseStarRocksSQLListener) ExitLocationsDesc(ctx *LocationsDescContext) {}

// EnterShowStorageVolumesStatement is called when production showStorageVolumesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// ExitShowStorageVolumesStatement is called when production showStorageVolumesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// EnterDropStorageVolumeStatement is called when production dropStorageVolumeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// ExitDropStorageVolumeStatement is called when production dropStorageVolumeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// ExitAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeClause is called when production alterStorageVolumeClause is entered.
func (s *BaseStarRocksSQLListener) EnterAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {
}

// ExitAlterStorageVolumeClause is called when production alterStorageVolumeClause is exited.
func (s *BaseStarRocksSQLListener) ExitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {
}

// EnterModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// ExitModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// EnterModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// ExitModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// EnterDescStorageVolumeStatement is called when production descStorageVolumeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// ExitDescStorageVolumeStatement is called when production descStorageVolumeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// EnterSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// ExitSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// EnterUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is entered.
func (s *BaseStarRocksSQLListener) EnterUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// ExitUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is exited.
func (s *BaseStarRocksSQLListener) ExitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// EnterShowFailPointStatement is called when production showFailPointStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowFailPointStatement(ctx *ShowFailPointStatementContext) {}

// ExitShowFailPointStatement is called when production showFailPointStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowFailPointStatement(ctx *ShowFailPointStatementContext) {}

// EnterCreateDictionaryStatement is called when production createDictionaryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {
}

// ExitCreateDictionaryStatement is called when production createDictionaryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {
}

// EnterDropDictionaryStatement is called when production dropDictionaryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropDictionaryStatement(ctx *DropDictionaryStatementContext) {
}

// ExitDropDictionaryStatement is called when production dropDictionaryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropDictionaryStatement(ctx *DropDictionaryStatementContext) {}

// EnterRefreshDictionaryStatement is called when production refreshDictionaryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// ExitRefreshDictionaryStatement is called when production refreshDictionaryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// EnterShowDictionaryStatement is called when production showDictionaryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {
}

// ExitShowDictionaryStatement is called when production showDictionaryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {}

// EnterCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// ExitCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// EnterDictionaryColumnDesc is called when production dictionaryColumnDesc is entered.
func (s *BaseStarRocksSQLListener) EnterDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// ExitDictionaryColumnDesc is called when production dictionaryColumnDesc is exited.
func (s *BaseStarRocksSQLListener) ExitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// EnterDictionaryName is called when production dictionaryName is entered.
func (s *BaseStarRocksSQLListener) EnterDictionaryName(ctx *DictionaryNameContext) {}

// ExitDictionaryName is called when production dictionaryName is exited.
func (s *BaseStarRocksSQLListener) ExitDictionaryName(ctx *DictionaryNameContext) {}

// EnterAlterClause is called when production alterClause is entered.
func (s *BaseStarRocksSQLListener) EnterAlterClause(ctx *AlterClauseContext) {}

// ExitAlterClause is called when production alterClause is exited.
func (s *BaseStarRocksSQLListener) ExitAlterClause(ctx *AlterClauseContext) {}

// EnterAddFrontendClause is called when production addFrontendClause is entered.
func (s *BaseStarRocksSQLListener) EnterAddFrontendClause(ctx *AddFrontendClauseContext) {}

// ExitAddFrontendClause is called when production addFrontendClause is exited.
func (s *BaseStarRocksSQLListener) ExitAddFrontendClause(ctx *AddFrontendClauseContext) {}

// EnterDropFrontendClause is called when production dropFrontendClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropFrontendClause(ctx *DropFrontendClauseContext) {}

// ExitDropFrontendClause is called when production dropFrontendClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropFrontendClause(ctx *DropFrontendClauseContext) {}

// EnterModifyFrontendHostClause is called when production modifyFrontendHostClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {
}

// ExitModifyFrontendHostClause is called when production modifyFrontendHostClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {
}

// EnterAddBackendClause is called when production addBackendClause is entered.
func (s *BaseStarRocksSQLListener) EnterAddBackendClause(ctx *AddBackendClauseContext) {}

// ExitAddBackendClause is called when production addBackendClause is exited.
func (s *BaseStarRocksSQLListener) ExitAddBackendClause(ctx *AddBackendClauseContext) {}

// EnterDropBackendClause is called when production dropBackendClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropBackendClause(ctx *DropBackendClauseContext) {}

// ExitDropBackendClause is called when production dropBackendClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropBackendClause(ctx *DropBackendClauseContext) {}

// EnterDecommissionBackendClause is called when production decommissionBackendClause is entered.
func (s *BaseStarRocksSQLListener) EnterDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {
}

// ExitDecommissionBackendClause is called when production decommissionBackendClause is exited.
func (s *BaseStarRocksSQLListener) ExitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {
}

// EnterModifyBackendClause is called when production modifyBackendClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// ExitModifyBackendClause is called when production modifyBackendClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// EnterAddComputeNodeClause is called when production addComputeNodeClause is entered.
func (s *BaseStarRocksSQLListener) EnterAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// ExitAddComputeNodeClause is called when production addComputeNodeClause is exited.
func (s *BaseStarRocksSQLListener) ExitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// EnterDropComputeNodeClause is called when production dropComputeNodeClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// ExitDropComputeNodeClause is called when production dropComputeNodeClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// EnterModifyBrokerClause is called when production modifyBrokerClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// ExitModifyBrokerClause is called when production modifyBrokerClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// EnterAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is entered.
func (s *BaseStarRocksSQLListener) EnterAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {
}

// ExitAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is exited.
func (s *BaseStarRocksSQLListener) ExitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {}

// EnterCreateImageClause is called when production createImageClause is entered.
func (s *BaseStarRocksSQLListener) EnterCreateImageClause(ctx *CreateImageClauseContext) {}

// ExitCreateImageClause is called when production createImageClause is exited.
func (s *BaseStarRocksSQLListener) ExitCreateImageClause(ctx *CreateImageClauseContext) {}

// EnterCleanTabletSchedQClause is called when production cleanTabletSchedQClause is entered.
func (s *BaseStarRocksSQLListener) EnterCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {
}

// ExitCleanTabletSchedQClause is called when production cleanTabletSchedQClause is exited.
func (s *BaseStarRocksSQLListener) ExitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {}

// EnterDecommissionDiskClause is called when production decommissionDiskClause is entered.
func (s *BaseStarRocksSQLListener) EnterDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {}

// ExitDecommissionDiskClause is called when production decommissionDiskClause is exited.
func (s *BaseStarRocksSQLListener) ExitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {}

// EnterCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is entered.
func (s *BaseStarRocksSQLListener) EnterCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// ExitCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is exited.
func (s *BaseStarRocksSQLListener) ExitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// EnterDisableDiskClause is called when production disableDiskClause is entered.
func (s *BaseStarRocksSQLListener) EnterDisableDiskClause(ctx *DisableDiskClauseContext) {}

// ExitDisableDiskClause is called when production disableDiskClause is exited.
func (s *BaseStarRocksSQLListener) ExitDisableDiskClause(ctx *DisableDiskClauseContext) {}

// EnterCancelDisableDiskClause is called when production cancelDisableDiskClause is entered.
func (s *BaseStarRocksSQLListener) EnterCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {
}

// ExitCancelDisableDiskClause is called when production cancelDisableDiskClause is exited.
func (s *BaseStarRocksSQLListener) ExitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {}

// EnterCreateIndexClause is called when production createIndexClause is entered.
func (s *BaseStarRocksSQLListener) EnterCreateIndexClause(ctx *CreateIndexClauseContext) {}

// ExitCreateIndexClause is called when production createIndexClause is exited.
func (s *BaseStarRocksSQLListener) ExitCreateIndexClause(ctx *CreateIndexClauseContext) {}

// EnterDropIndexClause is called when production dropIndexClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropIndexClause(ctx *DropIndexClauseContext) {}

// ExitDropIndexClause is called when production dropIndexClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropIndexClause(ctx *DropIndexClauseContext) {}

// EnterTableRenameClause is called when production tableRenameClause is entered.
func (s *BaseStarRocksSQLListener) EnterTableRenameClause(ctx *TableRenameClauseContext) {}

// ExitTableRenameClause is called when production tableRenameClause is exited.
func (s *BaseStarRocksSQLListener) ExitTableRenameClause(ctx *TableRenameClauseContext) {}

// EnterSwapTableClause is called when production swapTableClause is entered.
func (s *BaseStarRocksSQLListener) EnterSwapTableClause(ctx *SwapTableClauseContext) {}

// ExitSwapTableClause is called when production swapTableClause is exited.
func (s *BaseStarRocksSQLListener) ExitSwapTableClause(ctx *SwapTableClauseContext) {}

// EnterModifyPropertiesClause is called when production modifyPropertiesClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {}

// ExitModifyPropertiesClause is called when production modifyPropertiesClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {}

// EnterModifyCommentClause is called when production modifyCommentClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// ExitModifyCommentClause is called when production modifyCommentClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// EnterOptimizeRange is called when production optimizeRange is entered.
func (s *BaseStarRocksSQLListener) EnterOptimizeRange(ctx *OptimizeRangeContext) {}

// ExitOptimizeRange is called when production optimizeRange is exited.
func (s *BaseStarRocksSQLListener) ExitOptimizeRange(ctx *OptimizeRangeContext) {}

// EnterOptimizeClause is called when production optimizeClause is entered.
func (s *BaseStarRocksSQLListener) EnterOptimizeClause(ctx *OptimizeClauseContext) {}

// ExitOptimizeClause is called when production optimizeClause is exited.
func (s *BaseStarRocksSQLListener) ExitOptimizeClause(ctx *OptimizeClauseContext) {}

// EnterAddColumnClause is called when production addColumnClause is entered.
func (s *BaseStarRocksSQLListener) EnterAddColumnClause(ctx *AddColumnClauseContext) {}

// ExitAddColumnClause is called when production addColumnClause is exited.
func (s *BaseStarRocksSQLListener) ExitAddColumnClause(ctx *AddColumnClauseContext) {}

// EnterAddColumnsClause is called when production addColumnsClause is entered.
func (s *BaseStarRocksSQLListener) EnterAddColumnsClause(ctx *AddColumnsClauseContext) {}

// ExitAddColumnsClause is called when production addColumnsClause is exited.
func (s *BaseStarRocksSQLListener) ExitAddColumnsClause(ctx *AddColumnsClauseContext) {}

// EnterDropColumnClause is called when production dropColumnClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropColumnClause(ctx *DropColumnClauseContext) {}

// ExitDropColumnClause is called when production dropColumnClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropColumnClause(ctx *DropColumnClauseContext) {}

// EnterModifyColumnClause is called when production modifyColumnClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// ExitModifyColumnClause is called when production modifyColumnClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// EnterModifyColumnCommentClause is called when production modifyColumnCommentClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {
}

// ExitModifyColumnCommentClause is called when production modifyColumnCommentClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {
}

// EnterColumnRenameClause is called when production columnRenameClause is entered.
func (s *BaseStarRocksSQLListener) EnterColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// ExitColumnRenameClause is called when production columnRenameClause is exited.
func (s *BaseStarRocksSQLListener) ExitColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// EnterReorderColumnsClause is called when production reorderColumnsClause is entered.
func (s *BaseStarRocksSQLListener) EnterReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// ExitReorderColumnsClause is called when production reorderColumnsClause is exited.
func (s *BaseStarRocksSQLListener) ExitReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// EnterRollupRenameClause is called when production rollupRenameClause is entered.
func (s *BaseStarRocksSQLListener) EnterRollupRenameClause(ctx *RollupRenameClauseContext) {}

// ExitRollupRenameClause is called when production rollupRenameClause is exited.
func (s *BaseStarRocksSQLListener) ExitRollupRenameClause(ctx *RollupRenameClauseContext) {}

// EnterCompactionClause is called when production compactionClause is entered.
func (s *BaseStarRocksSQLListener) EnterCompactionClause(ctx *CompactionClauseContext) {}

// ExitCompactionClause is called when production compactionClause is exited.
func (s *BaseStarRocksSQLListener) ExitCompactionClause(ctx *CompactionClauseContext) {}

// EnterSubfieldName is called when production subfieldName is entered.
func (s *BaseStarRocksSQLListener) EnterSubfieldName(ctx *SubfieldNameContext) {}

// ExitSubfieldName is called when production subfieldName is exited.
func (s *BaseStarRocksSQLListener) ExitSubfieldName(ctx *SubfieldNameContext) {}

// EnterNestedFieldName is called when production nestedFieldName is entered.
func (s *BaseStarRocksSQLListener) EnterNestedFieldName(ctx *NestedFieldNameContext) {}

// ExitNestedFieldName is called when production nestedFieldName is exited.
func (s *BaseStarRocksSQLListener) ExitNestedFieldName(ctx *NestedFieldNameContext) {}

// EnterAddFieldClause is called when production addFieldClause is entered.
func (s *BaseStarRocksSQLListener) EnterAddFieldClause(ctx *AddFieldClauseContext) {}

// ExitAddFieldClause is called when production addFieldClause is exited.
func (s *BaseStarRocksSQLListener) ExitAddFieldClause(ctx *AddFieldClauseContext) {}

// EnterDropFieldClause is called when production dropFieldClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropFieldClause(ctx *DropFieldClauseContext) {}

// ExitDropFieldClause is called when production dropFieldClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropFieldClause(ctx *DropFieldClauseContext) {}

// EnterCreateOrReplaceTagClause is called when production createOrReplaceTagClause is entered.
func (s *BaseStarRocksSQLListener) EnterCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {
}

// ExitCreateOrReplaceTagClause is called when production createOrReplaceTagClause is exited.
func (s *BaseStarRocksSQLListener) ExitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {
}

// EnterCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is entered.
func (s *BaseStarRocksSQLListener) EnterCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// ExitCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is exited.
func (s *BaseStarRocksSQLListener) ExitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// EnterDropBranchClause is called when production dropBranchClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropBranchClause(ctx *DropBranchClauseContext) {}

// ExitDropBranchClause is called when production dropBranchClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropBranchClause(ctx *DropBranchClauseContext) {}

// EnterDropTagClause is called when production dropTagClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropTagClause(ctx *DropTagClauseContext) {}

// ExitDropTagClause is called when production dropTagClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropTagClause(ctx *DropTagClauseContext) {}

// EnterTableOperationClause is called when production tableOperationClause is entered.
func (s *BaseStarRocksSQLListener) EnterTableOperationClause(ctx *TableOperationClauseContext) {}

// ExitTableOperationClause is called when production tableOperationClause is exited.
func (s *BaseStarRocksSQLListener) ExitTableOperationClause(ctx *TableOperationClauseContext) {}

// EnterTagOptions is called when production tagOptions is entered.
func (s *BaseStarRocksSQLListener) EnterTagOptions(ctx *TagOptionsContext) {}

// ExitTagOptions is called when production tagOptions is exited.
func (s *BaseStarRocksSQLListener) ExitTagOptions(ctx *TagOptionsContext) {}

// EnterBranchOptions is called when production branchOptions is entered.
func (s *BaseStarRocksSQLListener) EnterBranchOptions(ctx *BranchOptionsContext) {}

// ExitBranchOptions is called when production branchOptions is exited.
func (s *BaseStarRocksSQLListener) ExitBranchOptions(ctx *BranchOptionsContext) {}

// EnterSnapshotRetention is called when production snapshotRetention is entered.
func (s *BaseStarRocksSQLListener) EnterSnapshotRetention(ctx *SnapshotRetentionContext) {}

// ExitSnapshotRetention is called when production snapshotRetention is exited.
func (s *BaseStarRocksSQLListener) ExitSnapshotRetention(ctx *SnapshotRetentionContext) {}

// EnterRefRetain is called when production refRetain is entered.
func (s *BaseStarRocksSQLListener) EnterRefRetain(ctx *RefRetainContext) {}

// ExitRefRetain is called when production refRetain is exited.
func (s *BaseStarRocksSQLListener) ExitRefRetain(ctx *RefRetainContext) {}

// EnterMaxSnapshotAge is called when production maxSnapshotAge is entered.
func (s *BaseStarRocksSQLListener) EnterMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// ExitMaxSnapshotAge is called when production maxSnapshotAge is exited.
func (s *BaseStarRocksSQLListener) ExitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// EnterMinSnapshotsToKeep is called when production minSnapshotsToKeep is entered.
func (s *BaseStarRocksSQLListener) EnterMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// ExitMinSnapshotsToKeep is called when production minSnapshotsToKeep is exited.
func (s *BaseStarRocksSQLListener) ExitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// EnterSnapshotId is called when production snapshotId is entered.
func (s *BaseStarRocksSQLListener) EnterSnapshotId(ctx *SnapshotIdContext) {}

// ExitSnapshotId is called when production snapshotId is exited.
func (s *BaseStarRocksSQLListener) ExitSnapshotId(ctx *SnapshotIdContext) {}

// EnterTimeUnit is called when production timeUnit is entered.
func (s *BaseStarRocksSQLListener) EnterTimeUnit(ctx *TimeUnitContext) {}

// ExitTimeUnit is called when production timeUnit is exited.
func (s *BaseStarRocksSQLListener) ExitTimeUnit(ctx *TimeUnitContext) {}

// EnterInteger_list is called when production integer_list is entered.
func (s *BaseStarRocksSQLListener) EnterInteger_list(ctx *Integer_listContext) {}

// ExitInteger_list is called when production integer_list is exited.
func (s *BaseStarRocksSQLListener) ExitInteger_list(ctx *Integer_listContext) {}

// EnterDropPersistentIndexClause is called when production dropPersistentIndexClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {
}

// ExitDropPersistentIndexClause is called when production dropPersistentIndexClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {
}

// EnterAddPartitionClause is called when production addPartitionClause is entered.
func (s *BaseStarRocksSQLListener) EnterAddPartitionClause(ctx *AddPartitionClauseContext) {}

// ExitAddPartitionClause is called when production addPartitionClause is exited.
func (s *BaseStarRocksSQLListener) ExitAddPartitionClause(ctx *AddPartitionClauseContext) {}

// EnterDropPartitionClause is called when production dropPartitionClause is entered.
func (s *BaseStarRocksSQLListener) EnterDropPartitionClause(ctx *DropPartitionClauseContext) {}

// ExitDropPartitionClause is called when production dropPartitionClause is exited.
func (s *BaseStarRocksSQLListener) ExitDropPartitionClause(ctx *DropPartitionClauseContext) {}

// EnterTruncatePartitionClause is called when production truncatePartitionClause is entered.
func (s *BaseStarRocksSQLListener) EnterTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {
}

// ExitTruncatePartitionClause is called when production truncatePartitionClause is exited.
func (s *BaseStarRocksSQLListener) ExitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {}

// EnterModifyPartitionClause is called when production modifyPartitionClause is entered.
func (s *BaseStarRocksSQLListener) EnterModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// ExitModifyPartitionClause is called when production modifyPartitionClause is exited.
func (s *BaseStarRocksSQLListener) ExitModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// EnterReplacePartitionClause is called when production replacePartitionClause is entered.
func (s *BaseStarRocksSQLListener) EnterReplacePartitionClause(ctx *ReplacePartitionClauseContext) {}

// ExitReplacePartitionClause is called when production replacePartitionClause is exited.
func (s *BaseStarRocksSQLListener) ExitReplacePartitionClause(ctx *ReplacePartitionClauseContext) {}

// EnterPartitionRenameClause is called when production partitionRenameClause is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// ExitPartitionRenameClause is called when production partitionRenameClause is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseStarRocksSQLListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseStarRocksSQLListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is entered.
func (s *BaseStarRocksSQLListener) EnterInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// ExitInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is exited.
func (s *BaseStarRocksSQLListener) ExitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// EnterColumnAliasesOrByName is called when production columnAliasesOrByName is entered.
func (s *BaseStarRocksSQLListener) EnterColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// ExitColumnAliasesOrByName is called when production columnAliasesOrByName is exited.
func (s *BaseStarRocksSQLListener) ExitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseStarRocksSQLListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseStarRocksSQLListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterCreateRoutineLoadStatement is called when production createRoutineLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// ExitCreateRoutineLoadStatement is called when production createRoutineLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// EnterAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {
}

// ExitAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {
}

// EnterDataSource is called when production dataSource is entered.
func (s *BaseStarRocksSQLListener) EnterDataSource(ctx *DataSourceContext) {}

// ExitDataSource is called when production dataSource is exited.
func (s *BaseStarRocksSQLListener) ExitDataSource(ctx *DataSourceContext) {}

// EnterLoadProperties is called when production loadProperties is entered.
func (s *BaseStarRocksSQLListener) EnterLoadProperties(ctx *LoadPropertiesContext) {}

// ExitLoadProperties is called when production loadProperties is exited.
func (s *BaseStarRocksSQLListener) ExitLoadProperties(ctx *LoadPropertiesContext) {}

// EnterColSeparatorProperty is called when production colSeparatorProperty is entered.
func (s *BaseStarRocksSQLListener) EnterColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// ExitColSeparatorProperty is called when production colSeparatorProperty is exited.
func (s *BaseStarRocksSQLListener) ExitColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// EnterRowDelimiterProperty is called when production rowDelimiterProperty is entered.
func (s *BaseStarRocksSQLListener) EnterRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// ExitRowDelimiterProperty is called when production rowDelimiterProperty is exited.
func (s *BaseStarRocksSQLListener) ExitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// EnterImportColumns is called when production importColumns is entered.
func (s *BaseStarRocksSQLListener) EnterImportColumns(ctx *ImportColumnsContext) {}

// ExitImportColumns is called when production importColumns is exited.
func (s *BaseStarRocksSQLListener) ExitImportColumns(ctx *ImportColumnsContext) {}

// EnterColumnProperties is called when production columnProperties is entered.
func (s *BaseStarRocksSQLListener) EnterColumnProperties(ctx *ColumnPropertiesContext) {}

// ExitColumnProperties is called when production columnProperties is exited.
func (s *BaseStarRocksSQLListener) ExitColumnProperties(ctx *ColumnPropertiesContext) {}

// EnterJobProperties is called when production jobProperties is entered.
func (s *BaseStarRocksSQLListener) EnterJobProperties(ctx *JobPropertiesContext) {}

// ExitJobProperties is called when production jobProperties is exited.
func (s *BaseStarRocksSQLListener) ExitJobProperties(ctx *JobPropertiesContext) {}

// EnterDataSourceProperties is called when production dataSourceProperties is entered.
func (s *BaseStarRocksSQLListener) EnterDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// ExitDataSourceProperties is called when production dataSourceProperties is exited.
func (s *BaseStarRocksSQLListener) ExitDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// EnterStopRoutineLoadStatement is called when production stopRoutineLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {
}

// ExitStopRoutineLoadStatement is called when production stopRoutineLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {
}

// EnterResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// ExitResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// EnterPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {
}

// ExitPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {
}

// EnterShowRoutineLoadStatement is called when production showRoutineLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {
}

// ExitShowRoutineLoadStatement is called when production showRoutineLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {
}

// EnterShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// ExitShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// EnterShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// ExitShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// EnterShowStreamLoadStatement is called when production showStreamLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {
}

// ExitShowStreamLoadStatement is called when production showStreamLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {}

// EnterAnalyzeStatement is called when production analyzeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// ExitAnalyzeStatement is called when production analyzeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// EnterRegularColumns is called when production regularColumns is entered.
func (s *BaseStarRocksSQLListener) EnterRegularColumns(ctx *RegularColumnsContext) {}

// ExitRegularColumns is called when production regularColumns is exited.
func (s *BaseStarRocksSQLListener) ExitRegularColumns(ctx *RegularColumnsContext) {}

// EnterAllColumns is called when production allColumns is entered.
func (s *BaseStarRocksSQLListener) EnterAllColumns(ctx *AllColumnsContext) {}

// ExitAllColumns is called when production allColumns is exited.
func (s *BaseStarRocksSQLListener) ExitAllColumns(ctx *AllColumnsContext) {}

// EnterPredicateColumns is called when production predicateColumns is entered.
func (s *BaseStarRocksSQLListener) EnterPredicateColumns(ctx *PredicateColumnsContext) {}

// ExitPredicateColumns is called when production predicateColumns is exited.
func (s *BaseStarRocksSQLListener) ExitPredicateColumns(ctx *PredicateColumnsContext) {}

// EnterMultiColumnSet is called when production multiColumnSet is entered.
func (s *BaseStarRocksSQLListener) EnterMultiColumnSet(ctx *MultiColumnSetContext) {}

// ExitMultiColumnSet is called when production multiColumnSet is exited.
func (s *BaseStarRocksSQLListener) ExitMultiColumnSet(ctx *MultiColumnSetContext) {}

// EnterDropStatsStatement is called when production dropStatsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropStatsStatement(ctx *DropStatsStatementContext) {}

// ExitDropStatsStatement is called when production dropStatsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropStatsStatement(ctx *DropStatsStatementContext) {}

// EnterHistogramStatement is called when production histogramStatement is entered.
func (s *BaseStarRocksSQLListener) EnterHistogramStatement(ctx *HistogramStatementContext) {}

// ExitHistogramStatement is called when production histogramStatement is exited.
func (s *BaseStarRocksSQLListener) ExitHistogramStatement(ctx *HistogramStatementContext) {}

// EnterAnalyzeHistogramStatement is called when production analyzeHistogramStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {
}

// ExitAnalyzeHistogramStatement is called when production analyzeHistogramStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {
}

// EnterDropHistogramStatement is called when production dropHistogramStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropHistogramStatement(ctx *DropHistogramStatementContext) {}

// ExitDropHistogramStatement is called when production dropHistogramStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropHistogramStatement(ctx *DropHistogramStatementContext) {}

// EnterCreateAnalyzeStatement is called when production createAnalyzeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {}

// ExitCreateAnalyzeStatement is called when production createAnalyzeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {}

// EnterDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {
}

// ExitDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {}

// EnterShowAnalyzeStatement is called when production showAnalyzeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// ExitShowAnalyzeStatement is called when production showAnalyzeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// EnterShowStatsMetaStatement is called when production showStatsMetaStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {}

// ExitShowStatsMetaStatement is called when production showStatsMetaStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {}

// EnterShowHistogramMetaStatement is called when production showHistogramMetaStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// ExitShowHistogramMetaStatement is called when production showHistogramMetaStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// EnterKillAnalyzeStatement is called when production killAnalyzeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// ExitKillAnalyzeStatement is called when production killAnalyzeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// EnterAnalyzeProfileStatement is called when production analyzeProfileStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {
}

// ExitAnalyzeProfileStatement is called when production analyzeProfileStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {}

// EnterCreateBaselinePlanStatement is called when production createBaselinePlanStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// ExitCreateBaselinePlanStatement is called when production createBaselinePlanStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// EnterDropBaselinePlanStatement is called when production dropBaselinePlanStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {
}

// ExitDropBaselinePlanStatement is called when production dropBaselinePlanStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {
}

// EnterShowBaselinePlanStatement is called when production showBaselinePlanStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {
}

// ExitShowBaselinePlanStatement is called when production showBaselinePlanStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {
}

// EnterCreateResourceGroupStatement is called when production createResourceGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// ExitCreateResourceGroupStatement is called when production createResourceGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// EnterDropResourceGroupStatement is called when production dropResourceGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// ExitDropResourceGroupStatement is called when production dropResourceGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// EnterAlterResourceGroupStatement is called when production alterResourceGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// ExitAlterResourceGroupStatement is called when production alterResourceGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// EnterShowResourceGroupStatement is called when production showResourceGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// ExitShowResourceGroupStatement is called when production showResourceGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// EnterShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// ExitShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// EnterCreateResourceStatement is called when production createResourceStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateResourceStatement(ctx *CreateResourceStatementContext) {
}

// ExitCreateResourceStatement is called when production createResourceStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateResourceStatement(ctx *CreateResourceStatementContext) {}

// EnterAlterResourceStatement is called when production alterResourceStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterResourceStatement(ctx *AlterResourceStatementContext) {}

// ExitAlterResourceStatement is called when production alterResourceStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterResourceStatement(ctx *AlterResourceStatementContext) {}

// EnterDropResourceStatement is called when production dropResourceStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropResourceStatement(ctx *DropResourceStatementContext) {}

// ExitDropResourceStatement is called when production dropResourceStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropResourceStatement(ctx *DropResourceStatementContext) {}

// EnterShowResourceStatement is called when production showResourceStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowResourceStatement(ctx *ShowResourceStatementContext) {}

// ExitShowResourceStatement is called when production showResourceStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowResourceStatement(ctx *ShowResourceStatementContext) {}

// EnterClassifier is called when production classifier is entered.
func (s *BaseStarRocksSQLListener) EnterClassifier(ctx *ClassifierContext) {}

// ExitClassifier is called when production classifier is exited.
func (s *BaseStarRocksSQLListener) ExitClassifier(ctx *ClassifierContext) {}

// EnterShowFunctionsStatement is called when production showFunctionsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {}

// ExitShowFunctionsStatement is called when production showFunctionsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {}

// EnterDropFunctionStatement is called when production dropFunctionStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// ExitDropFunctionStatement is called when production dropFunctionStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateFunctionStatement(ctx *CreateFunctionStatementContext) {
}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateFunctionStatement(ctx *CreateFunctionStatementContext) {}

// EnterInlineFunction is called when production inlineFunction is entered.
func (s *BaseStarRocksSQLListener) EnterInlineFunction(ctx *InlineFunctionContext) {}

// ExitInlineFunction is called when production inlineFunction is exited.
func (s *BaseStarRocksSQLListener) ExitInlineFunction(ctx *InlineFunctionContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseStarRocksSQLListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseStarRocksSQLListener) ExitTypeList(ctx *TypeListContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseStarRocksSQLListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseStarRocksSQLListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterDataDescList is called when production dataDescList is entered.
func (s *BaseStarRocksSQLListener) EnterDataDescList(ctx *DataDescListContext) {}

// ExitDataDescList is called when production dataDescList is exited.
func (s *BaseStarRocksSQLListener) ExitDataDescList(ctx *DataDescListContext) {}

// EnterDataDesc is called when production dataDesc is entered.
func (s *BaseStarRocksSQLListener) EnterDataDesc(ctx *DataDescContext) {}

// ExitDataDesc is called when production dataDesc is exited.
func (s *BaseStarRocksSQLListener) ExitDataDesc(ctx *DataDescContext) {}

// EnterFormatProps is called when production formatProps is entered.
func (s *BaseStarRocksSQLListener) EnterFormatProps(ctx *FormatPropsContext) {}

// ExitFormatProps is called when production formatProps is exited.
func (s *BaseStarRocksSQLListener) ExitFormatProps(ctx *FormatPropsContext) {}

// EnterBrokerDesc is called when production brokerDesc is entered.
func (s *BaseStarRocksSQLListener) EnterBrokerDesc(ctx *BrokerDescContext) {}

// ExitBrokerDesc is called when production brokerDesc is exited.
func (s *BaseStarRocksSQLListener) ExitBrokerDesc(ctx *BrokerDescContext) {}

// EnterResourceDesc is called when production resourceDesc is entered.
func (s *BaseStarRocksSQLListener) EnterResourceDesc(ctx *ResourceDescContext) {}

// ExitResourceDesc is called when production resourceDesc is exited.
func (s *BaseStarRocksSQLListener) ExitResourceDesc(ctx *ResourceDescContext) {}

// EnterShowLoadStatement is called when production showLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowLoadStatement(ctx *ShowLoadStatementContext) {}

// ExitShowLoadStatement is called when production showLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowLoadStatement(ctx *ShowLoadStatementContext) {}

// EnterShowLoadWarningsStatement is called when production showLoadWarningsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {
}

// ExitShowLoadWarningsStatement is called when production showLoadWarningsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {
}

// EnterCancelLoadStatement is called when production cancelLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// ExitCancelLoadStatement is called when production cancelLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// EnterAlterLoadStatement is called when production alterLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// ExitAlterLoadStatement is called when production alterLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// EnterCancelCompactionStatement is called when production cancelCompactionStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelCompactionStatement(ctx *CancelCompactionStatementContext) {
}

// ExitCancelCompactionStatement is called when production cancelCompactionStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelCompactionStatement(ctx *CancelCompactionStatementContext) {
}

// EnterShowAuthorStatement is called when production showAuthorStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// ExitShowAuthorStatement is called when production showAuthorStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// EnterShowBackendsStatement is called when production showBackendsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// ExitShowBackendsStatement is called when production showBackendsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// EnterShowBrokerStatement is called when production showBrokerStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// ExitShowBrokerStatement is called when production showBrokerStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// EnterShowCharsetStatement is called when production showCharsetStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// ExitShowCharsetStatement is called when production showCharsetStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// EnterShowCollationStatement is called when production showCollationStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCollationStatement(ctx *ShowCollationStatementContext) {}

// ExitShowCollationStatement is called when production showCollationStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCollationStatement(ctx *ShowCollationStatementContext) {}

// EnterShowDeleteStatement is called when production showDeleteStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// ExitShowDeleteStatement is called when production showDeleteStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// EnterShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// ExitShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// EnterShowEventsStatement is called when production showEventsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowEventsStatement(ctx *ShowEventsStatementContext) {}

// ExitShowEventsStatement is called when production showEventsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowEventsStatement(ctx *ShowEventsStatementContext) {}

// EnterShowEnginesStatement is called when production showEnginesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// ExitShowEnginesStatement is called when production showEnginesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// EnterShowFrontendsStatement is called when production showFrontendsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {}

// ExitShowFrontendsStatement is called when production showFrontendsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {}

// EnterShowFrontendsDisksStatement is called when production showFrontendsDisksStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) {
}

// ExitShowFrontendsDisksStatement is called when production showFrontendsDisksStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowFrontendsDisksStatement(ctx *ShowFrontendsDisksStatementContext) {
}

// EnterShowPluginsStatement is called when production showPluginsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// ExitShowPluginsStatement is called when production showPluginsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// EnterShowRepositoriesStatement is called when production showRepositoriesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {
}

// ExitShowRepositoriesStatement is called when production showRepositoriesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {
}

// EnterShowOpenTableStatement is called when production showOpenTableStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {}

// ExitShowOpenTableStatement is called when production showOpenTableStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {}

// EnterShowPrivilegesStatement is called when production showPrivilegesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {
}

// ExitShowPrivilegesStatement is called when production showPrivilegesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {}

// EnterShowProcedureStatement is called when production showProcedureStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowProcedureStatement(ctx *ShowProcedureStatementContext) {}

// ExitShowProcedureStatement is called when production showProcedureStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowProcedureStatement(ctx *ShowProcedureStatementContext) {}

// EnterShowProcStatement is called when production showProcStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowProcStatement(ctx *ShowProcStatementContext) {}

// ExitShowProcStatement is called when production showProcStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowProcStatement(ctx *ShowProcStatementContext) {}

// EnterShowProcesslistStatement is called when production showProcesslistStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {
}

// ExitShowProcesslistStatement is called when production showProcesslistStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {
}

// EnterShowProfilelistStatement is called when production showProfilelistStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {
}

// ExitShowProfilelistStatement is called when production showProfilelistStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {
}

// EnterShowRunningQueriesStatement is called when production showRunningQueriesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// ExitShowRunningQueriesStatement is called when production showRunningQueriesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// EnterShowStatusStatement is called when production showStatusStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowStatusStatement(ctx *ShowStatusStatementContext) {}

// ExitShowStatusStatement is called when production showStatusStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowStatusStatement(ctx *ShowStatusStatementContext) {}

// EnterShowTabletStatement is called when production showTabletStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTabletStatement(ctx *ShowTabletStatementContext) {}

// ExitShowTabletStatement is called when production showTabletStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTabletStatement(ctx *ShowTabletStatementContext) {}

// EnterShowTransactionStatement is called when production showTransactionStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTransactionStatement(ctx *ShowTransactionStatementContext) {
}

// ExitShowTransactionStatement is called when production showTransactionStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTransactionStatement(ctx *ShowTransactionStatementContext) {
}

// EnterShowTriggersStatement is called when production showTriggersStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// ExitShowTriggersStatement is called when production showTriggersStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// EnterShowUserPropertyStatement is called when production showUserPropertyStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {
}

// ExitShowUserPropertyStatement is called when production showUserPropertyStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {
}

// EnterShowVariablesStatement is called when production showVariablesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowVariablesStatement(ctx *ShowVariablesStatementContext) {}

// ExitShowVariablesStatement is called when production showVariablesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowVariablesStatement(ctx *ShowVariablesStatementContext) {}

// EnterShowWarningStatement is called when production showWarningStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowWarningStatement(ctx *ShowWarningStatementContext) {}

// ExitShowWarningStatement is called when production showWarningStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowWarningStatement(ctx *ShowWarningStatementContext) {}

// EnterHelpStatement is called when production helpStatement is entered.
func (s *BaseStarRocksSQLListener) EnterHelpStatement(ctx *HelpStatementContext) {}

// ExitHelpStatement is called when production helpStatement is exited.
func (s *BaseStarRocksSQLListener) ExitHelpStatement(ctx *HelpStatementContext) {}

// EnterShowQueryProfileStatement is called when production showQueryProfileStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) {
}

// ExitShowQueryProfileStatement is called when production showQueryProfileStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowQueryProfileStatement(ctx *ShowQueryProfileStatementContext) {
}

// EnterShowQueryStatsStatement is called when production showQueryStatsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) {
}

// ExitShowQueryStatsStatement is called when production showQueryStatsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowQueryStatsStatement(ctx *ShowQueryStatsStatementContext) {}

// EnterShowLoadProfileStatement is called when production showLoadProfileStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) {
}

// ExitShowLoadProfileStatement is called when production showLoadProfileStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowLoadProfileStatement(ctx *ShowLoadProfileStatementContext) {
}

// EnterShowDataSkewStatement is called when production showDataSkewStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDataSkewStatement(ctx *ShowDataSkewStatementContext) {}

// ExitShowDataSkewStatement is called when production showDataSkewStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDataSkewStatement(ctx *ShowDataSkewStatementContext) {}

// EnterShowDataTypesStatement is called when production showDataTypesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDataTypesStatement(ctx *ShowDataTypesStatementContext) {}

// ExitShowDataTypesStatement is called when production showDataTypesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDataTypesStatement(ctx *ShowDataTypesStatementContext) {}

// EnterShowSyncJobStatement is called when production showSyncJobStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowSyncJobStatement(ctx *ShowSyncJobStatementContext) {}

// ExitShowSyncJobStatement is called when production showSyncJobStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowSyncJobStatement(ctx *ShowSyncJobStatementContext) {}

// EnterShowPolicyStatement is called when production showPolicyStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowPolicyStatement(ctx *ShowPolicyStatementContext) {}

// ExitShowPolicyStatement is called when production showPolicyStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowPolicyStatement(ctx *ShowPolicyStatementContext) {}

// EnterShowSqlBlockRuleStatement is called when production showSqlBlockRuleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) {
}

// ExitShowSqlBlockRuleStatement is called when production showSqlBlockRuleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowSqlBlockRuleStatement(ctx *ShowSqlBlockRuleStatementContext) {
}

// EnterShowEncryptKeysStatement is called when production showEncryptKeysStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) {
}

// ExitShowEncryptKeysStatement is called when production showEncryptKeysStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowEncryptKeysStatement(ctx *ShowEncryptKeysStatementContext) {
}

// EnterShowCreateLoadStatement is called when production showCreateLoadStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) {
}

// ExitShowCreateLoadStatement is called when production showCreateLoadStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateLoadStatement(ctx *ShowCreateLoadStatementContext) {}

// EnterShowCreateRepositoryStatement is called when production showCreateRepositoryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) {
}

// ExitShowCreateRepositoryStatement is called when production showCreateRepositoryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateRepositoryStatement(ctx *ShowCreateRepositoryStatementContext) {
}

// EnterShowLastInsertStatement is called when production showLastInsertStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowLastInsertStatement(ctx *ShowLastInsertStatementContext) {
}

// ExitShowLastInsertStatement is called when production showLastInsertStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowLastInsertStatement(ctx *ShowLastInsertStatementContext) {}

// EnterShowTableIdStatement is called when production showTableIdStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTableIdStatement(ctx *ShowTableIdStatementContext) {}

// ExitShowTableIdStatement is called when production showTableIdStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTableIdStatement(ctx *ShowTableIdStatementContext) {}

// EnterShowDatabaseIdStatement is called when production showDatabaseIdStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) {
}

// ExitShowDatabaseIdStatement is called when production showDatabaseIdStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDatabaseIdStatement(ctx *ShowDatabaseIdStatementContext) {}

// EnterShowPartitionIdStatement is called when production showPartitionIdStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) {
}

// ExitShowPartitionIdStatement is called when production showPartitionIdStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowPartitionIdStatement(ctx *ShowPartitionIdStatementContext) {
}

// EnterShowTableStatsStatement is called when production showTableStatsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTableStatsStatement(ctx *ShowTableStatsStatementContext) {
}

// ExitShowTableStatsStatement is called when production showTableStatsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTableStatsStatement(ctx *ShowTableStatsStatementContext) {}

// EnterShowColumnStatsStatement is called when production showColumnStatsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) {
}

// ExitShowColumnStatsStatement is called when production showColumnStatsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowColumnStatsStatement(ctx *ShowColumnStatsStatementContext) {
}

// EnterShowConvertLightSchemaChangeStatement is called when production showConvertLightSchemaChangeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) {
}

// ExitShowConvertLightSchemaChangeStatement is called when production showConvertLightSchemaChangeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowConvertLightSchemaChangeStatement(ctx *ShowConvertLightSchemaChangeStatementContext) {
}

// EnterShowCatalogRecycleBinStatement is called when production showCatalogRecycleBinStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) {
}

// ExitShowCatalogRecycleBinStatement is called when production showCatalogRecycleBinStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCatalogRecycleBinStatement(ctx *ShowCatalogRecycleBinStatementContext) {
}

// EnterShowTrashStatement is called when production showTrashStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowTrashStatement(ctx *ShowTrashStatementContext) {}

// ExitShowTrashStatement is called when production showTrashStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowTrashStatement(ctx *ShowTrashStatementContext) {}

// EnterShowMigrationsStatement is called when production showMigrationsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowMigrationsStatement(ctx *ShowMigrationsStatementContext) {
}

// ExitShowMigrationsStatement is called when production showMigrationsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowMigrationsStatement(ctx *ShowMigrationsStatementContext) {}

// EnterShowWorkloadGroupsStatement is called when production showWorkloadGroupsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) {
}

// ExitShowWorkloadGroupsStatement is called when production showWorkloadGroupsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowWorkloadGroupsStatement(ctx *ShowWorkloadGroupsStatementContext) {
}

// EnterShowJobTaskStatement is called when production showJobTaskStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowJobTaskStatement(ctx *ShowJobTaskStatementContext) {}

// ExitShowJobTaskStatement is called when production showJobTaskStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowJobTaskStatement(ctx *ShowJobTaskStatementContext) {}

// EnterCreateUserStatement is called when production createUserStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateUserStatement(ctx *CreateUserStatementContext) {}

// ExitCreateUserStatement is called when production createUserStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateUserStatement(ctx *CreateUserStatementContext) {}

// EnterDropUserStatement is called when production dropUserStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropUserStatement(ctx *DropUserStatementContext) {}

// ExitDropUserStatement is called when production dropUserStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropUserStatement(ctx *DropUserStatementContext) {}

// EnterAlterUserStatement is called when production alterUserStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterUserStatement(ctx *AlterUserStatementContext) {}

// ExitAlterUserStatement is called when production alterUserStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterUserStatement(ctx *AlterUserStatementContext) {}

// EnterShowUserStatement is called when production showUserStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowUserStatement(ctx *ShowUserStatementContext) {}

// ExitShowUserStatement is called when production showUserStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowUserStatement(ctx *ShowUserStatementContext) {}

// EnterShowAllAuthentication is called when production showAllAuthentication is entered.
func (s *BaseStarRocksSQLListener) EnterShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// ExitShowAllAuthentication is called when production showAllAuthentication is exited.
func (s *BaseStarRocksSQLListener) ExitShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// EnterShowAuthenticationForUser is called when production showAuthenticationForUser is entered.
func (s *BaseStarRocksSQLListener) EnterShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {
}

// ExitShowAuthenticationForUser is called when production showAuthenticationForUser is exited.
func (s *BaseStarRocksSQLListener) ExitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {
}

// EnterExecuteAsStatement is called when production executeAsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// ExitExecuteAsStatement is called when production executeAsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// EnterAlterRoleStatement is called when production alterRoleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// ExitAlterRoleStatement is called when production alterRoleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// EnterDropRoleStatement is called when production dropRoleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropRoleStatement(ctx *DropRoleStatementContext) {}

// ExitDropRoleStatement is called when production dropRoleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropRoleStatement(ctx *DropRoleStatementContext) {}

// EnterShowRolesStatement is called when production showRolesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowRolesStatement(ctx *ShowRolesStatementContext) {}

// ExitShowRolesStatement is called when production showRolesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowRolesStatement(ctx *ShowRolesStatementContext) {}

// EnterGrantRoleToUser is called when production grantRoleToUser is entered.
func (s *BaseStarRocksSQLListener) EnterGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// ExitGrantRoleToUser is called when production grantRoleToUser is exited.
func (s *BaseStarRocksSQLListener) ExitGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// EnterGrantRoleToRole is called when production grantRoleToRole is entered.
func (s *BaseStarRocksSQLListener) EnterGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// ExitGrantRoleToRole is called when production grantRoleToRole is exited.
func (s *BaseStarRocksSQLListener) ExitGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// EnterRevokeRoleFromUser is called when production revokeRoleFromUser is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// ExitRevokeRoleFromUser is called when production revokeRoleFromUser is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// EnterRevokeRoleFromRole is called when production revokeRoleFromRole is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// ExitRevokeRoleFromRole is called when production revokeRoleFromRole is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// EnterSetRoleStatement is called when production setRoleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSetRoleStatement(ctx *SetRoleStatementContext) {}

// ExitSetRoleStatement is called when production setRoleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSetRoleStatement(ctx *SetRoleStatementContext) {}

// EnterSetDefaultRoleStatement is called when production setDefaultRoleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {
}

// ExitSetDefaultRoleStatement is called when production setDefaultRoleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {}

// EnterGrantRevokeClause is called when production grantRevokeClause is entered.
func (s *BaseStarRocksSQLListener) EnterGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// ExitGrantRevokeClause is called when production grantRevokeClause is exited.
func (s *BaseStarRocksSQLListener) ExitGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// EnterGrantOnUser is called when production grantOnUser is entered.
func (s *BaseStarRocksSQLListener) EnterGrantOnUser(ctx *GrantOnUserContext) {}

// ExitGrantOnUser is called when production grantOnUser is exited.
func (s *BaseStarRocksSQLListener) ExitGrantOnUser(ctx *GrantOnUserContext) {}

// EnterGrantOnTableBrief is called when production grantOnTableBrief is entered.
func (s *BaseStarRocksSQLListener) EnterGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// ExitGrantOnTableBrief is called when production grantOnTableBrief is exited.
func (s *BaseStarRocksSQLListener) ExitGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// EnterGrantOnFunc is called when production grantOnFunc is entered.
func (s *BaseStarRocksSQLListener) EnterGrantOnFunc(ctx *GrantOnFuncContext) {}

// ExitGrantOnFunc is called when production grantOnFunc is exited.
func (s *BaseStarRocksSQLListener) ExitGrantOnFunc(ctx *GrantOnFuncContext) {}

// EnterGrantOnSystem is called when production grantOnSystem is entered.
func (s *BaseStarRocksSQLListener) EnterGrantOnSystem(ctx *GrantOnSystemContext) {}

// ExitGrantOnSystem is called when production grantOnSystem is exited.
func (s *BaseStarRocksSQLListener) ExitGrantOnSystem(ctx *GrantOnSystemContext) {}

// EnterGrantOnPrimaryObj is called when production grantOnPrimaryObj is entered.
func (s *BaseStarRocksSQLListener) EnterGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// ExitGrantOnPrimaryObj is called when production grantOnPrimaryObj is exited.
func (s *BaseStarRocksSQLListener) ExitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// EnterGrantOnAll is called when production grantOnAll is entered.
func (s *BaseStarRocksSQLListener) EnterGrantOnAll(ctx *GrantOnAllContext) {}

// ExitGrantOnAll is called when production grantOnAll is exited.
func (s *BaseStarRocksSQLListener) ExitGrantOnAll(ctx *GrantOnAllContext) {}

// EnterRevokeOnUser is called when production revokeOnUser is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeOnUser(ctx *RevokeOnUserContext) {}

// ExitRevokeOnUser is called when production revokeOnUser is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeOnUser(ctx *RevokeOnUserContext) {}

// EnterRevokeOnTableBrief is called when production revokeOnTableBrief is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// ExitRevokeOnTableBrief is called when production revokeOnTableBrief is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// EnterRevokeOnFunc is called when production revokeOnFunc is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// ExitRevokeOnFunc is called when production revokeOnFunc is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// EnterRevokeOnSystem is called when production revokeOnSystem is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// ExitRevokeOnSystem is called when production revokeOnSystem is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// EnterRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// ExitRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// EnterRevokeOnAll is called when production revokeOnAll is entered.
func (s *BaseStarRocksSQLListener) EnterRevokeOnAll(ctx *RevokeOnAllContext) {}

// ExitRevokeOnAll is called when production revokeOnAll is exited.
func (s *BaseStarRocksSQLListener) ExitRevokeOnAll(ctx *RevokeOnAllContext) {}

// EnterShowGrantsStatement is called when production showGrantsStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// ExitShowGrantsStatement is called when production showGrantsStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// EnterAuthWithoutPlugin is called when production authWithoutPlugin is entered.
func (s *BaseStarRocksSQLListener) EnterAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// ExitAuthWithoutPlugin is called when production authWithoutPlugin is exited.
func (s *BaseStarRocksSQLListener) ExitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// EnterAuthWithPlugin is called when production authWithPlugin is entered.
func (s *BaseStarRocksSQLListener) EnterAuthWithPlugin(ctx *AuthWithPluginContext) {}

// ExitAuthWithPlugin is called when production authWithPlugin is exited.
func (s *BaseStarRocksSQLListener) ExitAuthWithPlugin(ctx *AuthWithPluginContext) {}

// EnterPrivObjectName is called when production privObjectName is entered.
func (s *BaseStarRocksSQLListener) EnterPrivObjectName(ctx *PrivObjectNameContext) {}

// ExitPrivObjectName is called when production privObjectName is exited.
func (s *BaseStarRocksSQLListener) ExitPrivObjectName(ctx *PrivObjectNameContext) {}

// EnterPrivObjectNameList is called when production privObjectNameList is entered.
func (s *BaseStarRocksSQLListener) EnterPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// ExitPrivObjectNameList is called when production privObjectNameList is exited.
func (s *BaseStarRocksSQLListener) ExitPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// EnterPrivFunctionObjectNameList is called when production privFunctionObjectNameList is entered.
func (s *BaseStarRocksSQLListener) EnterPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// ExitPrivFunctionObjectNameList is called when production privFunctionObjectNameList is exited.
func (s *BaseStarRocksSQLListener) ExitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// EnterPrivilegeTypeList is called when production privilegeTypeList is entered.
func (s *BaseStarRocksSQLListener) EnterPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// ExitPrivilegeTypeList is called when production privilegeTypeList is exited.
func (s *BaseStarRocksSQLListener) ExitPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *BaseStarRocksSQLListener) EnterPrivilegeType(ctx *PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *BaseStarRocksSQLListener) ExitPrivilegeType(ctx *PrivilegeTypeContext) {}

// EnterPrivObjectType is called when production privObjectType is entered.
func (s *BaseStarRocksSQLListener) EnterPrivObjectType(ctx *PrivObjectTypeContext) {}

// ExitPrivObjectType is called when production privObjectType is exited.
func (s *BaseStarRocksSQLListener) ExitPrivObjectType(ctx *PrivObjectTypeContext) {}

// EnterPrivObjectTypePlural is called when production privObjectTypePlural is entered.
func (s *BaseStarRocksSQLListener) EnterPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// ExitPrivObjectTypePlural is called when production privObjectTypePlural is exited.
func (s *BaseStarRocksSQLListener) ExitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// EnterCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// ExitCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// EnterAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// ExitAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// EnterDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// ExitDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// EnterShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// ExitShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// EnterShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// ExitShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// EnterCreateGroupProviderStatement is called when production createGroupProviderStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// ExitCreateGroupProviderStatement is called when production createGroupProviderStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// EnterDropGroupProviderStatement is called when production dropGroupProviderStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// ExitDropGroupProviderStatement is called when production dropGroupProviderStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// EnterShowGroupProvidersStatement is called when production showGroupProvidersStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// ExitShowGroupProvidersStatement is called when production showGroupProvidersStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// EnterShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// ExitShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// EnterBackupStatement is called when production backupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterBackupStatement(ctx *BackupStatementContext) {}

// ExitBackupStatement is called when production backupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitBackupStatement(ctx *BackupStatementContext) {}

// EnterCancelBackupStatement is called when production cancelBackupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// ExitCancelBackupStatement is called when production cancelBackupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// EnterShowBackupStatement is called when production showBackupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowBackupStatement(ctx *ShowBackupStatementContext) {}

// ExitShowBackupStatement is called when production showBackupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowBackupStatement(ctx *ShowBackupStatementContext) {}

// EnterRestoreStatement is called when production restoreStatement is entered.
func (s *BaseStarRocksSQLListener) EnterRestoreStatement(ctx *RestoreStatementContext) {}

// ExitRestoreStatement is called when production restoreStatement is exited.
func (s *BaseStarRocksSQLListener) ExitRestoreStatement(ctx *RestoreStatementContext) {}

// EnterCancelRestoreStatement is called when production cancelRestoreStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelRestoreStatement(ctx *CancelRestoreStatementContext) {}

// ExitCancelRestoreStatement is called when production cancelRestoreStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelRestoreStatement(ctx *CancelRestoreStatementContext) {}

// EnterShowRestoreStatement is called when production showRestoreStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// ExitShowRestoreStatement is called when production showRestoreStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// EnterShowSnapshotStatement is called when production showSnapshotStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// ExitShowSnapshotStatement is called when production showSnapshotStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// EnterCreateRepositoryStatement is called when production createRepositoryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {
}

// ExitCreateRepositoryStatement is called when production createRepositoryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {
}

// EnterDropRepositoryStatement is called when production dropRepositoryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropRepositoryStatement(ctx *DropRepositoryStatementContext) {
}

// ExitDropRepositoryStatement is called when production dropRepositoryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropRepositoryStatement(ctx *DropRepositoryStatementContext) {}

// EnterAddSqlBlackListStatement is called when production addSqlBlackListStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {
}

// ExitAddSqlBlackListStatement is called when production addSqlBlackListStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {
}

// EnterDelSqlBlackListStatement is called when production delSqlBlackListStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {
}

// ExitDelSqlBlackListStatement is called when production delSqlBlackListStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {
}

// EnterShowSqlBlackListStatement is called when production showSqlBlackListStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {
}

// ExitShowSqlBlackListStatement is called when production showSqlBlackListStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {
}

// EnterShowWhiteListStatement is called when production showWhiteListStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {}

// ExitShowWhiteListStatement is called when production showWhiteListStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {}

// EnterAddBackendBlackListStatement is called when production addBackendBlackListStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// ExitAddBackendBlackListStatement is called when production addBackendBlackListStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// EnterDelBackendBlackListStatement is called when production delBackendBlackListStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// ExitDelBackendBlackListStatement is called when production delBackendBlackListStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// EnterShowBackendBlackListStatement is called when production showBackendBlackListStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// ExitShowBackendBlackListStatement is called when production showBackendBlackListStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// EnterDataCacheTarget is called when production dataCacheTarget is entered.
func (s *BaseStarRocksSQLListener) EnterDataCacheTarget(ctx *DataCacheTargetContext) {}

// ExitDataCacheTarget is called when production dataCacheTarget is exited.
func (s *BaseStarRocksSQLListener) ExitDataCacheTarget(ctx *DataCacheTargetContext) {}

// EnterCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// ExitCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// EnterShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// ExitShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// EnterDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// ExitDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// EnterClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// ExitClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// EnterDataCacheSelectStatement is called when production dataCacheSelectStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {
}

// ExitDataCacheSelectStatement is called when production dataCacheSelectStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {
}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseStarRocksSQLListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseStarRocksSQLListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterCancelExportStatement is called when production cancelExportStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCancelExportStatement(ctx *CancelExportStatementContext) {}

// ExitCancelExportStatement is called when production cancelExportStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCancelExportStatement(ctx *CancelExportStatementContext) {}

// EnterShowExportStatement is called when production showExportStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowExportStatement(ctx *ShowExportStatementContext) {}

// ExitShowExportStatement is called when production showExportStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowExportStatement(ctx *ShowExportStatementContext) {}

// EnterInstallPluginStatement is called when production installPluginStatement is entered.
func (s *BaseStarRocksSQLListener) EnterInstallPluginStatement(ctx *InstallPluginStatementContext) {}

// ExitInstallPluginStatement is called when production installPluginStatement is exited.
func (s *BaseStarRocksSQLListener) ExitInstallPluginStatement(ctx *InstallPluginStatementContext) {}

// EnterUninstallPluginStatement is called when production uninstallPluginStatement is entered.
func (s *BaseStarRocksSQLListener) EnterUninstallPluginStatement(ctx *UninstallPluginStatementContext) {
}

// ExitUninstallPluginStatement is called when production uninstallPluginStatement is exited.
func (s *BaseStarRocksSQLListener) ExitUninstallPluginStatement(ctx *UninstallPluginStatementContext) {
}

// EnterCreateFileStatement is called when production createFileStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateFileStatement(ctx *CreateFileStatementContext) {}

// ExitCreateFileStatement is called when production createFileStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateFileStatement(ctx *CreateFileStatementContext) {}

// EnterDropFileStatement is called when production dropFileStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropFileStatement(ctx *DropFileStatementContext) {}

// ExitDropFileStatement is called when production dropFileStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropFileStatement(ctx *DropFileStatementContext) {}

// EnterShowSmallFilesStatement is called when production showSmallFilesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {
}

// ExitShowSmallFilesStatement is called when production showSmallFilesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {}

// EnterCreatePipeStatement is called when production createPipeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// ExitCreatePipeStatement is called when production createPipeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// EnterDropPipeStatement is called when production dropPipeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropPipeStatement(ctx *DropPipeStatementContext) {}

// ExitDropPipeStatement is called when production dropPipeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropPipeStatement(ctx *DropPipeStatementContext) {}

// EnterAlterPipeClause is called when production alterPipeClause is entered.
func (s *BaseStarRocksSQLListener) EnterAlterPipeClause(ctx *AlterPipeClauseContext) {}

// ExitAlterPipeClause is called when production alterPipeClause is exited.
func (s *BaseStarRocksSQLListener) ExitAlterPipeClause(ctx *AlterPipeClauseContext) {}

// EnterAlterPipeStatement is called when production alterPipeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// ExitAlterPipeStatement is called when production alterPipeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// EnterDescPipeStatement is called when production descPipeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDescPipeStatement(ctx *DescPipeStatementContext) {}

// ExitDescPipeStatement is called when production descPipeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDescPipeStatement(ctx *DescPipeStatementContext) {}

// EnterShowPipeStatement is called when production showPipeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowPipeStatement(ctx *ShowPipeStatementContext) {}

// ExitShowPipeStatement is called when production showPipeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowPipeStatement(ctx *ShowPipeStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterSetNames is called when production setNames is entered.
func (s *BaseStarRocksSQLListener) EnterSetNames(ctx *SetNamesContext) {}

// ExitSetNames is called when production setNames is exited.
func (s *BaseStarRocksSQLListener) ExitSetNames(ctx *SetNamesContext) {}

// EnterSetPassword is called when production setPassword is entered.
func (s *BaseStarRocksSQLListener) EnterSetPassword(ctx *SetPasswordContext) {}

// ExitSetPassword is called when production setPassword is exited.
func (s *BaseStarRocksSQLListener) ExitSetPassword(ctx *SetPasswordContext) {}

// EnterSetUserVar is called when production setUserVar is entered.
func (s *BaseStarRocksSQLListener) EnterSetUserVar(ctx *SetUserVarContext) {}

// ExitSetUserVar is called when production setUserVar is exited.
func (s *BaseStarRocksSQLListener) ExitSetUserVar(ctx *SetUserVarContext) {}

// EnterSetSystemVar is called when production setSystemVar is entered.
func (s *BaseStarRocksSQLListener) EnterSetSystemVar(ctx *SetSystemVarContext) {}

// ExitSetSystemVar is called when production setSystemVar is exited.
func (s *BaseStarRocksSQLListener) ExitSetSystemVar(ctx *SetSystemVarContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseStarRocksSQLListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseStarRocksSQLListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterTransaction_characteristics is called when production transaction_characteristics is entered.
func (s *BaseStarRocksSQLListener) EnterTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// ExitTransaction_characteristics is called when production transaction_characteristics is exited.
func (s *BaseStarRocksSQLListener) ExitTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// EnterTransaction_access_mode is called when production transaction_access_mode is entered.
func (s *BaseStarRocksSQLListener) EnterTransaction_access_mode(ctx *Transaction_access_modeContext) {
}

// ExitTransaction_access_mode is called when production transaction_access_mode is exited.
func (s *BaseStarRocksSQLListener) ExitTransaction_access_mode(ctx *Transaction_access_modeContext) {}

// EnterIsolation_level is called when production isolation_level is entered.
func (s *BaseStarRocksSQLListener) EnterIsolation_level(ctx *Isolation_levelContext) {}

// ExitIsolation_level is called when production isolation_level is exited.
func (s *BaseStarRocksSQLListener) ExitIsolation_level(ctx *Isolation_levelContext) {}

// EnterIsolation_types is called when production isolation_types is entered.
func (s *BaseStarRocksSQLListener) EnterIsolation_types(ctx *Isolation_typesContext) {}

// ExitIsolation_types is called when production isolation_types is exited.
func (s *BaseStarRocksSQLListener) ExitIsolation_types(ctx *Isolation_typesContext) {}

// EnterSetExprOrDefault is called when production setExprOrDefault is entered.
func (s *BaseStarRocksSQLListener) EnterSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// ExitSetExprOrDefault is called when production setExprOrDefault is exited.
func (s *BaseStarRocksSQLListener) ExitSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// EnterSetUserPropertyStatement is called when production setUserPropertyStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {
}

// ExitSetUserPropertyStatement is called when production setUserPropertyStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {
}

// EnterRoleList is called when production roleList is entered.
func (s *BaseStarRocksSQLListener) EnterRoleList(ctx *RoleListContext) {}

// ExitRoleList is called when production roleList is exited.
func (s *BaseStarRocksSQLListener) ExitRoleList(ctx *RoleListContext) {}

// EnterExecuteScriptStatement is called when production executeScriptStatement is entered.
func (s *BaseStarRocksSQLListener) EnterExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {}

// ExitExecuteScriptStatement is called when production executeScriptStatement is exited.
func (s *BaseStarRocksSQLListener) ExitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {}

// EnterUnsupportedStatement is called when production unsupportedStatement is entered.
func (s *BaseStarRocksSQLListener) EnterUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// ExitUnsupportedStatement is called when production unsupportedStatement is exited.
func (s *BaseStarRocksSQLListener) ExitUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// EnterLock_item is called when production lock_item is entered.
func (s *BaseStarRocksSQLListener) EnterLock_item(ctx *Lock_itemContext) {}

// ExitLock_item is called when production lock_item is exited.
func (s *BaseStarRocksSQLListener) ExitLock_item(ctx *Lock_itemContext) {}

// EnterLock_type is called when production lock_type is entered.
func (s *BaseStarRocksSQLListener) EnterLock_type(ctx *Lock_typeContext) {}

// ExitLock_type is called when production lock_type is exited.
func (s *BaseStarRocksSQLListener) ExitLock_type(ctx *Lock_typeContext) {}

// EnterAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// ExitAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// EnterTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is entered.
func (s *BaseStarRocksSQLListener) EnterTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// ExitTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is exited.
func (s *BaseStarRocksSQLListener) ExitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// EnterAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// ExitAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// EnterShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {
}

// ExitShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {
}

// EnterCreateWarehouseStatement is called when production createWarehouseStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {
}

// ExitCreateWarehouseStatement is called when production createWarehouseStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {
}

// EnterDropWarehouseStatement is called when production dropWarehouseStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropWarehouseStatement(ctx *DropWarehouseStatementContext) {}

// ExitDropWarehouseStatement is called when production dropWarehouseStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropWarehouseStatement(ctx *DropWarehouseStatementContext) {}

// EnterSuspendWarehouseStatement is called when production suspendWarehouseStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {
}

// ExitSuspendWarehouseStatement is called when production suspendWarehouseStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {
}

// EnterResumeWarehouseStatement is called when production resumeWarehouseStatement is entered.
func (s *BaseStarRocksSQLListener) EnterResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {
}

// ExitResumeWarehouseStatement is called when production resumeWarehouseStatement is exited.
func (s *BaseStarRocksSQLListener) ExitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {
}

// EnterSetWarehouseStatement is called when production setWarehouseStatement is entered.
func (s *BaseStarRocksSQLListener) EnterSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// ExitSetWarehouseStatement is called when production setWarehouseStatement is exited.
func (s *BaseStarRocksSQLListener) ExitSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// EnterShowWarehousesStatement is called when production showWarehousesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {
}

// ExitShowWarehousesStatement is called when production showWarehousesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {}

// EnterShowClustersStatement is called when production showClustersStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowClustersStatement(ctx *ShowClustersStatementContext) {}

// ExitShowClustersStatement is called when production showClustersStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowClustersStatement(ctx *ShowClustersStatementContext) {}

// EnterShowNodesStatement is called when production showNodesStatement is entered.
func (s *BaseStarRocksSQLListener) EnterShowNodesStatement(ctx *ShowNodesStatementContext) {}

// ExitShowNodesStatement is called when production showNodesStatement is exited.
func (s *BaseStarRocksSQLListener) ExitShowNodesStatement(ctx *ShowNodesStatementContext) {}

// EnterAlterWarehouseStatement is called when production alterWarehouseStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {
}

// ExitAlterWarehouseStatement is called when production alterWarehouseStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {}

// EnterCreateCNGroupStatement is called when production createCNGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {}

// ExitCreateCNGroupStatement is called when production createCNGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {}

// EnterDropCNGroupStatement is called when production dropCNGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// ExitDropCNGroupStatement is called when production dropCNGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// EnterEnableCNGroupStatement is called when production enableCNGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {}

// ExitEnableCNGroupStatement is called when production enableCNGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {}

// EnterDisableCNGroupStatement is called when production disableCNGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {
}

// ExitDisableCNGroupStatement is called when production disableCNGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {}

// EnterAlterCNGroupStatement is called when production alterCNGroupStatement is entered.
func (s *BaseStarRocksSQLListener) EnterAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// ExitAlterCNGroupStatement is called when production alterCNGroupStatement is exited.
func (s *BaseStarRocksSQLListener) ExitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// EnterBeginStatement is called when production beginStatement is entered.
func (s *BaseStarRocksSQLListener) EnterBeginStatement(ctx *BeginStatementContext) {}

// ExitBeginStatement is called when production beginStatement is exited.
func (s *BaseStarRocksSQLListener) ExitBeginStatement(ctx *BeginStatementContext) {}

// EnterCommitStatement is called when production commitStatement is entered.
func (s *BaseStarRocksSQLListener) EnterCommitStatement(ctx *CommitStatementContext) {}

// ExitCommitStatement is called when production commitStatement is exited.
func (s *BaseStarRocksSQLListener) ExitCommitStatement(ctx *CommitStatementContext) {}

// EnterRollbackStatement is called when production rollbackStatement is entered.
func (s *BaseStarRocksSQLListener) EnterRollbackStatement(ctx *RollbackStatementContext) {}

// ExitRollbackStatement is called when production rollbackStatement is exited.
func (s *BaseStarRocksSQLListener) ExitRollbackStatement(ctx *RollbackStatementContext) {}

// EnterTranslateStatement is called when production translateStatement is entered.
func (s *BaseStarRocksSQLListener) EnterTranslateStatement(ctx *TranslateStatementContext) {}

// ExitTranslateStatement is called when production translateStatement is exited.
func (s *BaseStarRocksSQLListener) ExitTranslateStatement(ctx *TranslateStatementContext) {}

// EnterDialect is called when production dialect is entered.
func (s *BaseStarRocksSQLListener) EnterDialect(ctx *DialectContext) {}

// ExitDialect is called when production dialect is exited.
func (s *BaseStarRocksSQLListener) ExitDialect(ctx *DialectContext) {}

// EnterTranslateSQL is called when production translateSQL is entered.
func (s *BaseStarRocksSQLListener) EnterTranslateSQL(ctx *TranslateSQLContext) {}

// ExitTranslateSQL is called when production translateSQL is exited.
func (s *BaseStarRocksSQLListener) ExitTranslateSQL(ctx *TranslateSQLContext) {}

// EnterQueryStatement is called when production queryStatement is entered.
func (s *BaseStarRocksSQLListener) EnterQueryStatement(ctx *QueryStatementContext) {}

// ExitQueryStatement is called when production queryStatement is exited.
func (s *BaseStarRocksSQLListener) ExitQueryStatement(ctx *QueryStatementContext) {}

// EnterQueryRelation is called when production queryRelation is entered.
func (s *BaseStarRocksSQLListener) EnterQueryRelation(ctx *QueryRelationContext) {}

// ExitQueryRelation is called when production queryRelation is exited.
func (s *BaseStarRocksSQLListener) ExitQueryRelation(ctx *QueryRelationContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseStarRocksSQLListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseStarRocksSQLListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterQueryNoWith is called when production queryNoWith is entered.
func (s *BaseStarRocksSQLListener) EnterQueryNoWith(ctx *QueryNoWithContext) {}

// ExitQueryNoWith is called when production queryNoWith is exited.
func (s *BaseStarRocksSQLListener) ExitQueryNoWith(ctx *QueryNoWithContext) {}

// EnterQueryPeriod is called when production queryPeriod is entered.
func (s *BaseStarRocksSQLListener) EnterQueryPeriod(ctx *QueryPeriodContext) {}

// ExitQueryPeriod is called when production queryPeriod is exited.
func (s *BaseStarRocksSQLListener) ExitQueryPeriod(ctx *QueryPeriodContext) {}

// EnterPeriodType is called when production periodType is entered.
func (s *BaseStarRocksSQLListener) EnterPeriodType(ctx *PeriodTypeContext) {}

// ExitPeriodType is called when production periodType is exited.
func (s *BaseStarRocksSQLListener) ExitPeriodType(ctx *PeriodTypeContext) {}

// EnterQueryWithParentheses is called when production queryWithParentheses is entered.
func (s *BaseStarRocksSQLListener) EnterQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// ExitQueryWithParentheses is called when production queryWithParentheses is exited.
func (s *BaseStarRocksSQLListener) ExitQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// EnterSetOperation is called when production setOperation is entered.
func (s *BaseStarRocksSQLListener) EnterSetOperation(ctx *SetOperationContext) {}

// ExitSetOperation is called when production setOperation is exited.
func (s *BaseStarRocksSQLListener) ExitSetOperation(ctx *SetOperationContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseStarRocksSQLListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseStarRocksSQLListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseStarRocksSQLListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseStarRocksSQLListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterRowConstructor is called when production rowConstructor is entered.
func (s *BaseStarRocksSQLListener) EnterRowConstructor(ctx *RowConstructorContext) {}

// ExitRowConstructor is called when production rowConstructor is exited.
func (s *BaseStarRocksSQLListener) ExitRowConstructor(ctx *RowConstructorContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseStarRocksSQLListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseStarRocksSQLListener) ExitSortItem(ctx *SortItemContext) {}

// EnterLimitConstExpr is called when production limitConstExpr is entered.
func (s *BaseStarRocksSQLListener) EnterLimitConstExpr(ctx *LimitConstExprContext) {}

// ExitLimitConstExpr is called when production limitConstExpr is exited.
func (s *BaseStarRocksSQLListener) ExitLimitConstExpr(ctx *LimitConstExprContext) {}

// EnterLimitElement is called when production limitElement is entered.
func (s *BaseStarRocksSQLListener) EnterLimitElement(ctx *LimitElementContext) {}

// ExitLimitElement is called when production limitElement is exited.
func (s *BaseStarRocksSQLListener) ExitLimitElement(ctx *LimitElementContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseStarRocksSQLListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseStarRocksSQLListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterFrom is called when production from is entered.
func (s *BaseStarRocksSQLListener) EnterFrom(ctx *FromContext) {}

// ExitFrom is called when production from is exited.
func (s *BaseStarRocksSQLListener) ExitFrom(ctx *FromContext) {}

// EnterDual is called when production dual is entered.
func (s *BaseStarRocksSQLListener) EnterDual(ctx *DualContext) {}

// ExitDual is called when production dual is exited.
func (s *BaseStarRocksSQLListener) ExitDual(ctx *DualContext) {}

// EnterRollup is called when production rollup is entered.
func (s *BaseStarRocksSQLListener) EnterRollup(ctx *RollupContext) {}

// ExitRollup is called when production rollup is exited.
func (s *BaseStarRocksSQLListener) ExitRollup(ctx *RollupContext) {}

// EnterCube is called when production cube is entered.
func (s *BaseStarRocksSQLListener) EnterCube(ctx *CubeContext) {}

// ExitCube is called when production cube is exited.
func (s *BaseStarRocksSQLListener) ExitCube(ctx *CubeContext) {}

// EnterMultipleGroupingSets is called when production multipleGroupingSets is entered.
func (s *BaseStarRocksSQLListener) EnterMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// ExitMultipleGroupingSets is called when production multipleGroupingSets is exited.
func (s *BaseStarRocksSQLListener) ExitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// EnterSingleGroupingSet is called when production singleGroupingSet is entered.
func (s *BaseStarRocksSQLListener) EnterSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// ExitSingleGroupingSet is called when production singleGroupingSet is exited.
func (s *BaseStarRocksSQLListener) ExitSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// EnterGroupingSet is called when production groupingSet is entered.
func (s *BaseStarRocksSQLListener) EnterGroupingSet(ctx *GroupingSetContext) {}

// ExitGroupingSet is called when production groupingSet is exited.
func (s *BaseStarRocksSQLListener) ExitGroupingSet(ctx *GroupingSetContext) {}

// EnterCommonTableExpression is called when production commonTableExpression is entered.
func (s *BaseStarRocksSQLListener) EnterCommonTableExpression(ctx *CommonTableExpressionContext) {}

// ExitCommonTableExpression is called when production commonTableExpression is exited.
func (s *BaseStarRocksSQLListener) ExitCommonTableExpression(ctx *CommonTableExpressionContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseStarRocksSQLListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseStarRocksSQLListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterSelectSingle is called when production selectSingle is entered.
func (s *BaseStarRocksSQLListener) EnterSelectSingle(ctx *SelectSingleContext) {}

// ExitSelectSingle is called when production selectSingle is exited.
func (s *BaseStarRocksSQLListener) ExitSelectSingle(ctx *SelectSingleContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseStarRocksSQLListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseStarRocksSQLListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterExcludeClause is called when production excludeClause is entered.
func (s *BaseStarRocksSQLListener) EnterExcludeClause(ctx *ExcludeClauseContext) {}

// ExitExcludeClause is called when production excludeClause is exited.
func (s *BaseStarRocksSQLListener) ExitExcludeClause(ctx *ExcludeClauseContext) {}

// EnterRelations is called when production relations is entered.
func (s *BaseStarRocksSQLListener) EnterRelations(ctx *RelationsContext) {}

// ExitRelations is called when production relations is exited.
func (s *BaseStarRocksSQLListener) ExitRelations(ctx *RelationsContext) {}

// EnterRelationLateralView is called when production relationLateralView is entered.
func (s *BaseStarRocksSQLListener) EnterRelationLateralView(ctx *RelationLateralViewContext) {}

// ExitRelationLateralView is called when production relationLateralView is exited.
func (s *BaseStarRocksSQLListener) ExitRelationLateralView(ctx *RelationLateralViewContext) {}

// EnterLateralView is called when production lateralView is entered.
func (s *BaseStarRocksSQLListener) EnterLateralView(ctx *LateralViewContext) {}

// ExitLateralView is called when production lateralView is exited.
func (s *BaseStarRocksSQLListener) ExitLateralView(ctx *LateralViewContext) {}

// EnterGeneratorFunction is called when production generatorFunction is entered.
func (s *BaseStarRocksSQLListener) EnterGeneratorFunction(ctx *GeneratorFunctionContext) {}

// ExitGeneratorFunction is called when production generatorFunction is exited.
func (s *BaseStarRocksSQLListener) ExitGeneratorFunction(ctx *GeneratorFunctionContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseStarRocksSQLListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseStarRocksSQLListener) ExitRelation(ctx *RelationContext) {}

// EnterTableAtom is called when production tableAtom is entered.
func (s *BaseStarRocksSQLListener) EnterTableAtom(ctx *TableAtomContext) {}

// ExitTableAtom is called when production tableAtom is exited.
func (s *BaseStarRocksSQLListener) ExitTableAtom(ctx *TableAtomContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseStarRocksSQLListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseStarRocksSQLListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterSubqueryWithAlias is called when production subqueryWithAlias is entered.
func (s *BaseStarRocksSQLListener) EnterSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// ExitSubqueryWithAlias is called when production subqueryWithAlias is exited.
func (s *BaseStarRocksSQLListener) ExitSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// EnterTableFunction is called when production tableFunction is entered.
func (s *BaseStarRocksSQLListener) EnterTableFunction(ctx *TableFunctionContext) {}

// ExitTableFunction is called when production tableFunction is exited.
func (s *BaseStarRocksSQLListener) ExitTableFunction(ctx *TableFunctionContext) {}

// EnterNormalizedTableFunction is called when production normalizedTableFunction is entered.
func (s *BaseStarRocksSQLListener) EnterNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {
}

// ExitNormalizedTableFunction is called when production normalizedTableFunction is exited.
func (s *BaseStarRocksSQLListener) ExitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {}

// EnterFileTableFunction is called when production fileTableFunction is entered.
func (s *BaseStarRocksSQLListener) EnterFileTableFunction(ctx *FileTableFunctionContext) {}

// ExitFileTableFunction is called when production fileTableFunction is exited.
func (s *BaseStarRocksSQLListener) ExitFileTableFunction(ctx *FileTableFunctionContext) {}

// EnterParenthesizedRelation is called when production parenthesizedRelation is entered.
func (s *BaseStarRocksSQLListener) EnterParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// ExitParenthesizedRelation is called when production parenthesizedRelation is exited.
func (s *BaseStarRocksSQLListener) ExitParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// EnterPivotClause is called when production pivotClause is entered.
func (s *BaseStarRocksSQLListener) EnterPivotClause(ctx *PivotClauseContext) {}

// ExitPivotClause is called when production pivotClause is exited.
func (s *BaseStarRocksSQLListener) ExitPivotClause(ctx *PivotClauseContext) {}

// EnterPivotAggregationExpression is called when production pivotAggregationExpression is entered.
func (s *BaseStarRocksSQLListener) EnterPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// ExitPivotAggregationExpression is called when production pivotAggregationExpression is exited.
func (s *BaseStarRocksSQLListener) ExitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// EnterPivotValue is called when production pivotValue is entered.
func (s *BaseStarRocksSQLListener) EnterPivotValue(ctx *PivotValueContext) {}

// ExitPivotValue is called when production pivotValue is exited.
func (s *BaseStarRocksSQLListener) ExitPivotValue(ctx *PivotValueContext) {}

// EnterSampleClause is called when production sampleClause is entered.
func (s *BaseStarRocksSQLListener) EnterSampleClause(ctx *SampleClauseContext) {}

// ExitSampleClause is called when production sampleClause is exited.
func (s *BaseStarRocksSQLListener) ExitSampleClause(ctx *SampleClauseContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseStarRocksSQLListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseStarRocksSQLListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterNamedArgumentList is called when production namedArgumentList is entered.
func (s *BaseStarRocksSQLListener) EnterNamedArgumentList(ctx *NamedArgumentListContext) {}

// ExitNamedArgumentList is called when production namedArgumentList is exited.
func (s *BaseStarRocksSQLListener) ExitNamedArgumentList(ctx *NamedArgumentListContext) {}

// EnterNamedArguments is called when production namedArguments is entered.
func (s *BaseStarRocksSQLListener) EnterNamedArguments(ctx *NamedArgumentsContext) {}

// ExitNamedArguments is called when production namedArguments is exited.
func (s *BaseStarRocksSQLListener) ExitNamedArguments(ctx *NamedArgumentsContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseStarRocksSQLListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseStarRocksSQLListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterCrossOrInnerJoinType is called when production crossOrInnerJoinType is entered.
func (s *BaseStarRocksSQLListener) EnterCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// ExitCrossOrInnerJoinType is called when production crossOrInnerJoinType is exited.
func (s *BaseStarRocksSQLListener) ExitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// EnterOuterAndSemiJoinType is called when production outerAndSemiJoinType is entered.
func (s *BaseStarRocksSQLListener) EnterOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// ExitOuterAndSemiJoinType is called when production outerAndSemiJoinType is exited.
func (s *BaseStarRocksSQLListener) ExitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// EnterBracketHint is called when production bracketHint is entered.
func (s *BaseStarRocksSQLListener) EnterBracketHint(ctx *BracketHintContext) {}

// ExitBracketHint is called when production bracketHint is exited.
func (s *BaseStarRocksSQLListener) ExitBracketHint(ctx *BracketHintContext) {}

// EnterHintMap is called when production hintMap is entered.
func (s *BaseStarRocksSQLListener) EnterHintMap(ctx *HintMapContext) {}

// ExitHintMap is called when production hintMap is exited.
func (s *BaseStarRocksSQLListener) ExitHintMap(ctx *HintMapContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseStarRocksSQLListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseStarRocksSQLListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseStarRocksSQLListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseStarRocksSQLListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterColumnAliasesWithoutParentheses is called when production columnAliasesWithoutParentheses is entered.
func (s *BaseStarRocksSQLListener) EnterColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) {
}

// ExitColumnAliasesWithoutParentheses is called when production columnAliasesWithoutParentheses is exited.
func (s *BaseStarRocksSQLListener) ExitColumnAliasesWithoutParentheses(ctx *ColumnAliasesWithoutParenthesesContext) {
}

// EnterPartitionNames is called when production partitionNames is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionNames(ctx *PartitionNamesContext) {}

// ExitPartitionNames is called when production partitionNames is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionNames(ctx *PartitionNamesContext) {}

// EnterKeyPartitionList is called when production keyPartitionList is entered.
func (s *BaseStarRocksSQLListener) EnterKeyPartitionList(ctx *KeyPartitionListContext) {}

// ExitKeyPartitionList is called when production keyPartitionList is exited.
func (s *BaseStarRocksSQLListener) ExitKeyPartitionList(ctx *KeyPartitionListContext) {}

// EnterTabletList is called when production tabletList is entered.
func (s *BaseStarRocksSQLListener) EnterTabletList(ctx *TabletListContext) {}

// ExitTabletList is called when production tabletList is exited.
func (s *BaseStarRocksSQLListener) ExitTabletList(ctx *TabletListContext) {}

// EnterPrepareStatement is called when production prepareStatement is entered.
func (s *BaseStarRocksSQLListener) EnterPrepareStatement(ctx *PrepareStatementContext) {}

// ExitPrepareStatement is called when production prepareStatement is exited.
func (s *BaseStarRocksSQLListener) ExitPrepareStatement(ctx *PrepareStatementContext) {}

// EnterPrepareSql is called when production prepareSql is entered.
func (s *BaseStarRocksSQLListener) EnterPrepareSql(ctx *PrepareSqlContext) {}

// ExitPrepareSql is called when production prepareSql is exited.
func (s *BaseStarRocksSQLListener) ExitPrepareSql(ctx *PrepareSqlContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseStarRocksSQLListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseStarRocksSQLListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterDeallocateStatement is called when production deallocateStatement is entered.
func (s *BaseStarRocksSQLListener) EnterDeallocateStatement(ctx *DeallocateStatementContext) {}

// ExitDeallocateStatement is called when production deallocateStatement is exited.
func (s *BaseStarRocksSQLListener) ExitDeallocateStatement(ctx *DeallocateStatementContext) {}

// EnterReplicaList is called when production replicaList is entered.
func (s *BaseStarRocksSQLListener) EnterReplicaList(ctx *ReplicaListContext) {}

// ExitReplicaList is called when production replicaList is exited.
func (s *BaseStarRocksSQLListener) ExitReplicaList(ctx *ReplicaListContext) {}

// EnterExpressionsWithDefault is called when production expressionsWithDefault is entered.
func (s *BaseStarRocksSQLListener) EnterExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {}

// ExitExpressionsWithDefault is called when production expressionsWithDefault is exited.
func (s *BaseStarRocksSQLListener) ExitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {}

// EnterExpressionOrDefault is called when production expressionOrDefault is entered.
func (s *BaseStarRocksSQLListener) EnterExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// ExitExpressionOrDefault is called when production expressionOrDefault is exited.
func (s *BaseStarRocksSQLListener) ExitExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// EnterMapExpressionList is called when production mapExpressionList is entered.
func (s *BaseStarRocksSQLListener) EnterMapExpressionList(ctx *MapExpressionListContext) {}

// ExitMapExpressionList is called when production mapExpressionList is exited.
func (s *BaseStarRocksSQLListener) ExitMapExpressionList(ctx *MapExpressionListContext) {}

// EnterMapExpression is called when production mapExpression is entered.
func (s *BaseStarRocksSQLListener) EnterMapExpression(ctx *MapExpressionContext) {}

// ExitMapExpression is called when production mapExpression is exited.
func (s *BaseStarRocksSQLListener) ExitMapExpression(ctx *MapExpressionContext) {}

// EnterExpressionSingleton is called when production expressionSingleton is entered.
func (s *BaseStarRocksSQLListener) EnterExpressionSingleton(ctx *ExpressionSingletonContext) {}

// ExitExpressionSingleton is called when production expressionSingleton is exited.
func (s *BaseStarRocksSQLListener) ExitExpressionSingleton(ctx *ExpressionSingletonContext) {}

// EnterExpressionDefault is called when production expressionDefault is entered.
func (s *BaseStarRocksSQLListener) EnterExpressionDefault(ctx *ExpressionDefaultContext) {}

// ExitExpressionDefault is called when production expressionDefault is exited.
func (s *BaseStarRocksSQLListener) ExitExpressionDefault(ctx *ExpressionDefaultContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseStarRocksSQLListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseStarRocksSQLListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseStarRocksSQLListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseStarRocksSQLListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseStarRocksSQLListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseStarRocksSQLListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseStarRocksSQLListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseStarRocksSQLListener) ExitComparison(ctx *ComparisonContext) {}

// EnterBooleanExpressionDefault is called when production booleanExpressionDefault is entered.
func (s *BaseStarRocksSQLListener) EnterBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {
}

// ExitBooleanExpressionDefault is called when production booleanExpressionDefault is exited.
func (s *BaseStarRocksSQLListener) ExitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {
}

// EnterIsNull is called when production isNull is entered.
func (s *BaseStarRocksSQLListener) EnterIsNull(ctx *IsNullContext) {}

// ExitIsNull is called when production isNull is exited.
func (s *BaseStarRocksSQLListener) ExitIsNull(ctx *IsNullContext) {}

// EnterScalarSubquery is called when production scalarSubquery is entered.
func (s *BaseStarRocksSQLListener) EnterScalarSubquery(ctx *ScalarSubqueryContext) {}

// ExitScalarSubquery is called when production scalarSubquery is exited.
func (s *BaseStarRocksSQLListener) ExitScalarSubquery(ctx *ScalarSubqueryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseStarRocksSQLListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseStarRocksSQLListener) ExitPredicate(ctx *PredicateContext) {}

// EnterTupleInSubquery is called when production tupleInSubquery is entered.
func (s *BaseStarRocksSQLListener) EnterTupleInSubquery(ctx *TupleInSubqueryContext) {}

// ExitTupleInSubquery is called when production tupleInSubquery is exited.
func (s *BaseStarRocksSQLListener) ExitTupleInSubquery(ctx *TupleInSubqueryContext) {}

// EnterInSubquery is called when production inSubquery is entered.
func (s *BaseStarRocksSQLListener) EnterInSubquery(ctx *InSubqueryContext) {}

// ExitInSubquery is called when production inSubquery is exited.
func (s *BaseStarRocksSQLListener) ExitInSubquery(ctx *InSubqueryContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseStarRocksSQLListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseStarRocksSQLListener) ExitInList(ctx *InListContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseStarRocksSQLListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseStarRocksSQLListener) ExitBetween(ctx *BetweenContext) {}

// EnterLike is called when production like is entered.
func (s *BaseStarRocksSQLListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseStarRocksSQLListener) ExitLike(ctx *LikeContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseStarRocksSQLListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseStarRocksSQLListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseStarRocksSQLListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseStarRocksSQLListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseStarRocksSQLListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseStarRocksSQLListener) ExitDereference(ctx *DereferenceContext) {}

// EnterOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is entered.
func (s *BaseStarRocksSQLListener) EnterOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// ExitOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is exited.
func (s *BaseStarRocksSQLListener) ExitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// EnterMatchExpr is called when production matchExpr is entered.
func (s *BaseStarRocksSQLListener) EnterMatchExpr(ctx *MatchExprContext) {}

// ExitMatchExpr is called when production matchExpr is exited.
func (s *BaseStarRocksSQLListener) ExitMatchExpr(ctx *MatchExprContext) {}

// EnterColumnRef is called when production columnRef is entered.
func (s *BaseStarRocksSQLListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production columnRef is exited.
func (s *BaseStarRocksSQLListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterConvert is called when production convert is entered.
func (s *BaseStarRocksSQLListener) EnterConvert(ctx *ConvertContext) {}

// ExitConvert is called when production convert is exited.
func (s *BaseStarRocksSQLListener) ExitConvert(ctx *ConvertContext) {}

// EnterCollectionSubscript is called when production collectionSubscript is entered.
func (s *BaseStarRocksSQLListener) EnterCollectionSubscript(ctx *CollectionSubscriptContext) {}

// ExitCollectionSubscript is called when production collectionSubscript is exited.
func (s *BaseStarRocksSQLListener) ExitCollectionSubscript(ctx *CollectionSubscriptContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseStarRocksSQLListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseStarRocksSQLListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseStarRocksSQLListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseStarRocksSQLListener) ExitCast(ctx *CastContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseStarRocksSQLListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseStarRocksSQLListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterUserVariableExpression is called when production userVariableExpression is entered.
func (s *BaseStarRocksSQLListener) EnterUserVariableExpression(ctx *UserVariableExpressionContext) {}

// ExitUserVariableExpression is called when production userVariableExpression is exited.
func (s *BaseStarRocksSQLListener) ExitUserVariableExpression(ctx *UserVariableExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseStarRocksSQLListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseStarRocksSQLListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseStarRocksSQLListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseStarRocksSQLListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterArrowExpression is called when production arrowExpression is entered.
func (s *BaseStarRocksSQLListener) EnterArrowExpression(ctx *ArrowExpressionContext) {}

// ExitArrowExpression is called when production arrowExpression is exited.
func (s *BaseStarRocksSQLListener) ExitArrowExpression(ctx *ArrowExpressionContext) {}

// EnterArrayExpr is called when production arrayExpr is entered.
func (s *BaseStarRocksSQLListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production arrayExpr is exited.
func (s *BaseStarRocksSQLListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterSystemVariableExpression is called when production systemVariableExpression is entered.
func (s *BaseStarRocksSQLListener) EnterSystemVariableExpression(ctx *SystemVariableExpressionContext) {
}

// ExitSystemVariableExpression is called when production systemVariableExpression is exited.
func (s *BaseStarRocksSQLListener) ExitSystemVariableExpression(ctx *SystemVariableExpressionContext) {
}

// EnterConcat is called when production concat is entered.
func (s *BaseStarRocksSQLListener) EnterConcat(ctx *ConcatContext) {}

// ExitConcat is called when production concat is exited.
func (s *BaseStarRocksSQLListener) ExitConcat(ctx *ConcatContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseStarRocksSQLListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseStarRocksSQLListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterLambdaFunctionExpr is called when production lambdaFunctionExpr is entered.
func (s *BaseStarRocksSQLListener) EnterLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// ExitLambdaFunctionExpr is called when production lambdaFunctionExpr is exited.
func (s *BaseStarRocksSQLListener) ExitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// EnterDictionaryGetExpr is called when production dictionaryGetExpr is entered.
func (s *BaseStarRocksSQLListener) EnterDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// ExitDictionaryGetExpr is called when production dictionaryGetExpr is exited.
func (s *BaseStarRocksSQLListener) ExitDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// EnterCollate is called when production collate is entered.
func (s *BaseStarRocksSQLListener) EnterCollate(ctx *CollateContext) {}

// ExitCollate is called when production collate is exited.
func (s *BaseStarRocksSQLListener) ExitCollate(ctx *CollateContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseStarRocksSQLListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseStarRocksSQLListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterMapConstructor is called when production mapConstructor is entered.
func (s *BaseStarRocksSQLListener) EnterMapConstructor(ctx *MapConstructorContext) {}

// ExitMapConstructor is called when production mapConstructor is exited.
func (s *BaseStarRocksSQLListener) ExitMapConstructor(ctx *MapConstructorContext) {}

// EnterArraySlice is called when production arraySlice is entered.
func (s *BaseStarRocksSQLListener) EnterArraySlice(ctx *ArraySliceContext) {}

// ExitArraySlice is called when production arraySlice is exited.
func (s *BaseStarRocksSQLListener) ExitArraySlice(ctx *ArraySliceContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseStarRocksSQLListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseStarRocksSQLListener) ExitExists(ctx *ExistsContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseStarRocksSQLListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseStarRocksSQLListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseStarRocksSQLListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseStarRocksSQLListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterUnitBoundaryLiteral is called when production unitBoundaryLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// ExitUnitBoundaryLiteral is called when production unitBoundaryLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// EnterBinaryLiteral is called when production binaryLiteral is entered.
func (s *BaseStarRocksSQLListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production binaryLiteral is exited.
func (s *BaseStarRocksSQLListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterParameter is called when production Parameter is entered.
func (s *BaseStarRocksSQLListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production Parameter is exited.
func (s *BaseStarRocksSQLListener) ExitParameter(ctx *ParameterContext) {}

// EnterExtract is called when production extract is entered.
func (s *BaseStarRocksSQLListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BaseStarRocksSQLListener) ExitExtract(ctx *ExtractContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseStarRocksSQLListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseStarRocksSQLListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterInformationFunction is called when production informationFunction is entered.
func (s *BaseStarRocksSQLListener) EnterInformationFunction(ctx *InformationFunctionContext) {}

// ExitInformationFunction is called when production informationFunction is exited.
func (s *BaseStarRocksSQLListener) ExitInformationFunction(ctx *InformationFunctionContext) {}

// EnterSpecialDateTime is called when production specialDateTime is entered.
func (s *BaseStarRocksSQLListener) EnterSpecialDateTime(ctx *SpecialDateTimeContext) {}

// ExitSpecialDateTime is called when production specialDateTime is exited.
func (s *BaseStarRocksSQLListener) ExitSpecialDateTime(ctx *SpecialDateTimeContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseStarRocksSQLListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseStarRocksSQLListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterAggregationFunctionCall is called when production aggregationFunctionCall is entered.
func (s *BaseStarRocksSQLListener) EnterAggregationFunctionCall(ctx *AggregationFunctionCallContext) {
}

// ExitAggregationFunctionCall is called when production aggregationFunctionCall is exited.
func (s *BaseStarRocksSQLListener) ExitAggregationFunctionCall(ctx *AggregationFunctionCallContext) {}

// EnterWindowFunctionCall is called when production windowFunctionCall is entered.
func (s *BaseStarRocksSQLListener) EnterWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// ExitWindowFunctionCall is called when production windowFunctionCall is exited.
func (s *BaseStarRocksSQLListener) ExitWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// EnterTranslateFunctionCall is called when production translateFunctionCall is entered.
func (s *BaseStarRocksSQLListener) EnterTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// ExitTranslateFunctionCall is called when production translateFunctionCall is exited.
func (s *BaseStarRocksSQLListener) ExitTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// EnterSimpleFunctionCall is called when production simpleFunctionCall is entered.
func (s *BaseStarRocksSQLListener) EnterSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// ExitSimpleFunctionCall is called when production simpleFunctionCall is exited.
func (s *BaseStarRocksSQLListener) ExitSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseStarRocksSQLListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseStarRocksSQLListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterUserVariable is called when production userVariable is entered.
func (s *BaseStarRocksSQLListener) EnterUserVariable(ctx *UserVariableContext) {}

// ExitUserVariable is called when production userVariable is exited.
func (s *BaseStarRocksSQLListener) ExitUserVariable(ctx *UserVariableContext) {}

// EnterSystemVariable is called when production systemVariable is entered.
func (s *BaseStarRocksSQLListener) EnterSystemVariable(ctx *SystemVariableContext) {}

// ExitSystemVariable is called when production systemVariable is exited.
func (s *BaseStarRocksSQLListener) ExitSystemVariable(ctx *SystemVariableContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseStarRocksSQLListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseStarRocksSQLListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterInformationFunctionExpression is called when production informationFunctionExpression is entered.
func (s *BaseStarRocksSQLListener) EnterInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// ExitInformationFunctionExpression is called when production informationFunctionExpression is exited.
func (s *BaseStarRocksSQLListener) ExitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// EnterSpecialDateTimeExpression is called when production specialDateTimeExpression is entered.
func (s *BaseStarRocksSQLListener) EnterSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {
}

// ExitSpecialDateTimeExpression is called when production specialDateTimeExpression is exited.
func (s *BaseStarRocksSQLListener) ExitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {
}

// EnterSpecialFunctionExpression is called when production specialFunctionExpression is entered.
func (s *BaseStarRocksSQLListener) EnterSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {
}

// ExitSpecialFunctionExpression is called when production specialFunctionExpression is exited.
func (s *BaseStarRocksSQLListener) ExitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {
}

// EnterWindowFunction is called when production windowFunction is entered.
func (s *BaseStarRocksSQLListener) EnterWindowFunction(ctx *WindowFunctionContext) {}

// ExitWindowFunction is called when production windowFunction is exited.
func (s *BaseStarRocksSQLListener) ExitWindowFunction(ctx *WindowFunctionContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseStarRocksSQLListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseStarRocksSQLListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterOver is called when production over is entered.
func (s *BaseStarRocksSQLListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BaseStarRocksSQLListener) ExitOver(ctx *OverContext) {}

// EnterIgnoreNulls is called when production ignoreNulls is entered.
func (s *BaseStarRocksSQLListener) EnterIgnoreNulls(ctx *IgnoreNullsContext) {}

// ExitIgnoreNulls is called when production ignoreNulls is exited.
func (s *BaseStarRocksSQLListener) ExitIgnoreNulls(ctx *IgnoreNullsContext) {}

// EnterWindowFrame is called when production windowFrame is entered.
func (s *BaseStarRocksSQLListener) EnterWindowFrame(ctx *WindowFrameContext) {}

// ExitWindowFrame is called when production windowFrame is exited.
func (s *BaseStarRocksSQLListener) ExitWindowFrame(ctx *WindowFrameContext) {}

// EnterUnboundedFrame is called when production unboundedFrame is entered.
func (s *BaseStarRocksSQLListener) EnterUnboundedFrame(ctx *UnboundedFrameContext) {}

// ExitUnboundedFrame is called when production unboundedFrame is exited.
func (s *BaseStarRocksSQLListener) ExitUnboundedFrame(ctx *UnboundedFrameContext) {}

// EnterCurrentRowBound is called when production currentRowBound is entered.
func (s *BaseStarRocksSQLListener) EnterCurrentRowBound(ctx *CurrentRowBoundContext) {}

// ExitCurrentRowBound is called when production currentRowBound is exited.
func (s *BaseStarRocksSQLListener) ExitCurrentRowBound(ctx *CurrentRowBoundContext) {}

// EnterBoundedFrame is called when production boundedFrame is entered.
func (s *BaseStarRocksSQLListener) EnterBoundedFrame(ctx *BoundedFrameContext) {}

// ExitBoundedFrame is called when production boundedFrame is exited.
func (s *BaseStarRocksSQLListener) ExitBoundedFrame(ctx *BoundedFrameContext) {}

// EnterBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is entered.
func (s *BaseStarRocksSQLListener) EnterBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {
}

// ExitBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is exited.
func (s *BaseStarRocksSQLListener) ExitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {}

// EnterTableDesc is called when production tableDesc is entered.
func (s *BaseStarRocksSQLListener) EnterTableDesc(ctx *TableDescContext) {}

// ExitTableDesc is called when production tableDesc is exited.
func (s *BaseStarRocksSQLListener) ExitTableDesc(ctx *TableDescContext) {}

// EnterBackupRestoreTableDesc is called when production backupRestoreTableDesc is entered.
func (s *BaseStarRocksSQLListener) EnterBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {}

// ExitBackupRestoreTableDesc is called when production backupRestoreTableDesc is exited.
func (s *BaseStarRocksSQLListener) ExitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {}

// EnterExplainDesc is called when production explainDesc is entered.
func (s *BaseStarRocksSQLListener) EnterExplainDesc(ctx *ExplainDescContext) {}

// ExitExplainDesc is called when production explainDesc is exited.
func (s *BaseStarRocksSQLListener) ExitExplainDesc(ctx *ExplainDescContext) {}

// EnterOptimizerTrace is called when production optimizerTrace is entered.
func (s *BaseStarRocksSQLListener) EnterOptimizerTrace(ctx *OptimizerTraceContext) {}

// ExitOptimizerTrace is called when production optimizerTrace is exited.
func (s *BaseStarRocksSQLListener) ExitOptimizerTrace(ctx *OptimizerTraceContext) {}

// EnterPartitionExpr is called when production partitionExpr is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionExpr(ctx *PartitionExprContext) {}

// ExitPartitionExpr is called when production partitionExpr is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionExpr(ctx *PartitionExprContext) {}

// EnterPartitionDesc is called when production partitionDesc is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionDesc(ctx *PartitionDescContext) {}

// ExitPartitionDesc is called when production partitionDesc is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionDesc(ctx *PartitionDescContext) {}

// EnterListPartitionDesc is called when production listPartitionDesc is entered.
func (s *BaseStarRocksSQLListener) EnterListPartitionDesc(ctx *ListPartitionDescContext) {}

// ExitListPartitionDesc is called when production listPartitionDesc is exited.
func (s *BaseStarRocksSQLListener) ExitListPartitionDesc(ctx *ListPartitionDescContext) {}

// EnterSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is entered.
func (s *BaseStarRocksSQLListener) EnterSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// ExitSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is exited.
func (s *BaseStarRocksSQLListener) ExitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// EnterMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is entered.
func (s *BaseStarRocksSQLListener) EnterMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// ExitMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is exited.
func (s *BaseStarRocksSQLListener) ExitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// EnterMultiListPartitionValues is called when production multiListPartitionValues is entered.
func (s *BaseStarRocksSQLListener) EnterMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {
}

// ExitMultiListPartitionValues is called when production multiListPartitionValues is exited.
func (s *BaseStarRocksSQLListener) ExitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {
}

// EnterSingleListPartitionValues is called when production singleListPartitionValues is entered.
func (s *BaseStarRocksSQLListener) EnterSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {
}

// ExitSingleListPartitionValues is called when production singleListPartitionValues is exited.
func (s *BaseStarRocksSQLListener) ExitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {
}

// EnterListPartitionValues is called when production listPartitionValues is entered.
func (s *BaseStarRocksSQLListener) EnterListPartitionValues(ctx *ListPartitionValuesContext) {}

// ExitListPartitionValues is called when production listPartitionValues is exited.
func (s *BaseStarRocksSQLListener) ExitListPartitionValues(ctx *ListPartitionValuesContext) {}

// EnterListPartitionValue is called when production listPartitionValue is entered.
func (s *BaseStarRocksSQLListener) EnterListPartitionValue(ctx *ListPartitionValueContext) {}

// ExitListPartitionValue is called when production listPartitionValue is exited.
func (s *BaseStarRocksSQLListener) ExitListPartitionValue(ctx *ListPartitionValueContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseStarRocksSQLListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseStarRocksSQLListener) ExitStringList(ctx *StringListContext) {}

// EnterLiteralExpressionList is called when production literalExpressionList is entered.
func (s *BaseStarRocksSQLListener) EnterLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// ExitLiteralExpressionList is called when production literalExpressionList is exited.
func (s *BaseStarRocksSQLListener) ExitLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// EnterRangePartitionDesc is called when production rangePartitionDesc is entered.
func (s *BaseStarRocksSQLListener) EnterRangePartitionDesc(ctx *RangePartitionDescContext) {}

// ExitRangePartitionDesc is called when production rangePartitionDesc is exited.
func (s *BaseStarRocksSQLListener) ExitRangePartitionDesc(ctx *RangePartitionDescContext) {}

// EnterSingleRangePartition is called when production singleRangePartition is entered.
func (s *BaseStarRocksSQLListener) EnterSingleRangePartition(ctx *SingleRangePartitionContext) {}

// ExitSingleRangePartition is called when production singleRangePartition is exited.
func (s *BaseStarRocksSQLListener) ExitSingleRangePartition(ctx *SingleRangePartitionContext) {}

// EnterMultiRangePartition is called when production multiRangePartition is entered.
func (s *BaseStarRocksSQLListener) EnterMultiRangePartition(ctx *MultiRangePartitionContext) {}

// ExitMultiRangePartition is called when production multiRangePartition is exited.
func (s *BaseStarRocksSQLListener) ExitMultiRangePartition(ctx *MultiRangePartitionContext) {}

// EnterPartitionRangeDesc is called when production partitionRangeDesc is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// ExitPartitionRangeDesc is called when production partitionRangeDesc is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// EnterPartitionKeyDesc is called when production partitionKeyDesc is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// ExitPartitionKeyDesc is called when production partitionKeyDesc is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// EnterPartitionValueList is called when production partitionValueList is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionValueList(ctx *PartitionValueListContext) {}

// ExitPartitionValueList is called when production partitionValueList is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionValueList(ctx *PartitionValueListContext) {}

// EnterKeyPartition is called when production keyPartition is entered.
func (s *BaseStarRocksSQLListener) EnterKeyPartition(ctx *KeyPartitionContext) {}

// ExitKeyPartition is called when production keyPartition is exited.
func (s *BaseStarRocksSQLListener) ExitKeyPartition(ctx *KeyPartitionContext) {}

// EnterPartitionValue is called when production partitionValue is entered.
func (s *BaseStarRocksSQLListener) EnterPartitionValue(ctx *PartitionValueContext) {}

// ExitPartitionValue is called when production partitionValue is exited.
func (s *BaseStarRocksSQLListener) ExitPartitionValue(ctx *PartitionValueContext) {}

// EnterDistributionClause is called when production distributionClause is entered.
func (s *BaseStarRocksSQLListener) EnterDistributionClause(ctx *DistributionClauseContext) {}

// ExitDistributionClause is called when production distributionClause is exited.
func (s *BaseStarRocksSQLListener) ExitDistributionClause(ctx *DistributionClauseContext) {}

// EnterDistributionDesc is called when production distributionDesc is entered.
func (s *BaseStarRocksSQLListener) EnterDistributionDesc(ctx *DistributionDescContext) {}

// ExitDistributionDesc is called when production distributionDesc is exited.
func (s *BaseStarRocksSQLListener) ExitDistributionDesc(ctx *DistributionDescContext) {}

// EnterRefreshSchemeDesc is called when production refreshSchemeDesc is entered.
func (s *BaseStarRocksSQLListener) EnterRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// ExitRefreshSchemeDesc is called when production refreshSchemeDesc is exited.
func (s *BaseStarRocksSQLListener) ExitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// EnterStatusDesc is called when production statusDesc is entered.
func (s *BaseStarRocksSQLListener) EnterStatusDesc(ctx *StatusDescContext) {}

// ExitStatusDesc is called when production statusDesc is exited.
func (s *BaseStarRocksSQLListener) ExitStatusDesc(ctx *StatusDescContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseStarRocksSQLListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseStarRocksSQLListener) ExitProperties(ctx *PropertiesContext) {}

// EnterExtProperties is called when production extProperties is entered.
func (s *BaseStarRocksSQLListener) EnterExtProperties(ctx *ExtPropertiesContext) {}

// ExitExtProperties is called when production extProperties is exited.
func (s *BaseStarRocksSQLListener) ExitExtProperties(ctx *ExtPropertiesContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseStarRocksSQLListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseStarRocksSQLListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterUserPropertyList is called when production userPropertyList is entered.
func (s *BaseStarRocksSQLListener) EnterUserPropertyList(ctx *UserPropertyListContext) {}

// ExitUserPropertyList is called when production userPropertyList is exited.
func (s *BaseStarRocksSQLListener) ExitUserPropertyList(ctx *UserPropertyListContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseStarRocksSQLListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseStarRocksSQLListener) ExitProperty(ctx *PropertyContext) {}

// EnterInlineProperties is called when production inlineProperties is entered.
func (s *BaseStarRocksSQLListener) EnterInlineProperties(ctx *InlinePropertiesContext) {}

// ExitInlineProperties is called when production inlineProperties is exited.
func (s *BaseStarRocksSQLListener) ExitInlineProperties(ctx *InlinePropertiesContext) {}

// EnterInlineProperty is called when production inlineProperty is entered.
func (s *BaseStarRocksSQLListener) EnterInlineProperty(ctx *InlinePropertyContext) {}

// ExitInlineProperty is called when production inlineProperty is exited.
func (s *BaseStarRocksSQLListener) ExitInlineProperty(ctx *InlinePropertyContext) {}

// EnterVarType is called when production varType is entered.
func (s *BaseStarRocksSQLListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseStarRocksSQLListener) ExitVarType(ctx *VarTypeContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseStarRocksSQLListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseStarRocksSQLListener) ExitComment(ctx *CommentContext) {}

// EnterOutfile is called when production outfile is entered.
func (s *BaseStarRocksSQLListener) EnterOutfile(ctx *OutfileContext) {}

// ExitOutfile is called when production outfile is exited.
func (s *BaseStarRocksSQLListener) ExitOutfile(ctx *OutfileContext) {}

// EnterFileFormat is called when production fileFormat is entered.
func (s *BaseStarRocksSQLListener) EnterFileFormat(ctx *FileFormatContext) {}

// ExitFileFormat is called when production fileFormat is exited.
func (s *BaseStarRocksSQLListener) ExitFileFormat(ctx *FileFormatContext) {}

// EnterString is called when production string is entered.
func (s *BaseStarRocksSQLListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseStarRocksSQLListener) ExitString(ctx *StringContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseStarRocksSQLListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseStarRocksSQLListener) ExitBinary(ctx *BinaryContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseStarRocksSQLListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseStarRocksSQLListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseStarRocksSQLListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseStarRocksSQLListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseStarRocksSQLListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseStarRocksSQLListener) ExitInterval(ctx *IntervalContext) {}

// EnterTaskInterval is called when production taskInterval is entered.
func (s *BaseStarRocksSQLListener) EnterTaskInterval(ctx *TaskIntervalContext) {}

// ExitTaskInterval is called when production taskInterval is exited.
func (s *BaseStarRocksSQLListener) ExitTaskInterval(ctx *TaskIntervalContext) {}

// EnterTaskUnitIdentifier is called when production taskUnitIdentifier is entered.
func (s *BaseStarRocksSQLListener) EnterTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// ExitTaskUnitIdentifier is called when production taskUnitIdentifier is exited.
func (s *BaseStarRocksSQLListener) ExitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// EnterUnitIdentifier is called when production unitIdentifier is entered.
func (s *BaseStarRocksSQLListener) EnterUnitIdentifier(ctx *UnitIdentifierContext) {}

// ExitUnitIdentifier is called when production unitIdentifier is exited.
func (s *BaseStarRocksSQLListener) ExitUnitIdentifier(ctx *UnitIdentifierContext) {}

// EnterUnitBoundary is called when production unitBoundary is entered.
func (s *BaseStarRocksSQLListener) EnterUnitBoundary(ctx *UnitBoundaryContext) {}

// ExitUnitBoundary is called when production unitBoundary is exited.
func (s *BaseStarRocksSQLListener) ExitUnitBoundary(ctx *UnitBoundaryContext) {}

// EnterType is called when production type is entered.
func (s *BaseStarRocksSQLListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseStarRocksSQLListener) ExitType(ctx *TypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseStarRocksSQLListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseStarRocksSQLListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseStarRocksSQLListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseStarRocksSQLListener) ExitMapType(ctx *MapTypeContext) {}

// EnterSubfieldDesc is called when production subfieldDesc is entered.
func (s *BaseStarRocksSQLListener) EnterSubfieldDesc(ctx *SubfieldDescContext) {}

// ExitSubfieldDesc is called when production subfieldDesc is exited.
func (s *BaseStarRocksSQLListener) ExitSubfieldDesc(ctx *SubfieldDescContext) {}

// EnterSubfieldDescs is called when production subfieldDescs is entered.
func (s *BaseStarRocksSQLListener) EnterSubfieldDescs(ctx *SubfieldDescsContext) {}

// ExitSubfieldDescs is called when production subfieldDescs is exited.
func (s *BaseStarRocksSQLListener) ExitSubfieldDescs(ctx *SubfieldDescsContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseStarRocksSQLListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseStarRocksSQLListener) ExitStructType(ctx *StructTypeContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseStarRocksSQLListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseStarRocksSQLListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseStarRocksSQLListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseStarRocksSQLListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterDecimalType is called when production decimalType is entered.
func (s *BaseStarRocksSQLListener) EnterDecimalType(ctx *DecimalTypeContext) {}

// ExitDecimalType is called when production decimalType is exited.
func (s *BaseStarRocksSQLListener) ExitDecimalType(ctx *DecimalTypeContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseStarRocksSQLListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseStarRocksSQLListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseStarRocksSQLListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseStarRocksSQLListener) ExitTableName(ctx *TableNameContext) {}

// EnterWriteBranch is called when production writeBranch is entered.
func (s *BaseStarRocksSQLListener) EnterWriteBranch(ctx *WriteBranchContext) {}

// ExitWriteBranch is called when production writeBranch is exited.
func (s *BaseStarRocksSQLListener) ExitWriteBranch(ctx *WriteBranchContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseStarRocksSQLListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseStarRocksSQLListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterDigitIdentifier is called when production digitIdentifier is entered.
func (s *BaseStarRocksSQLListener) EnterDigitIdentifier(ctx *DigitIdentifierContext) {}

// ExitDigitIdentifier is called when production digitIdentifier is exited.
func (s *BaseStarRocksSQLListener) ExitDigitIdentifier(ctx *DigitIdentifierContext) {}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseStarRocksSQLListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseStarRocksSQLListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterIdentifierWithAlias is called when production identifierWithAlias is entered.
func (s *BaseStarRocksSQLListener) EnterIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// ExitIdentifierWithAlias is called when production identifierWithAlias is exited.
func (s *BaseStarRocksSQLListener) ExitIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// EnterIdentifierWithAliasList is called when production identifierWithAliasList is entered.
func (s *BaseStarRocksSQLListener) EnterIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {
}

// ExitIdentifierWithAliasList is called when production identifierWithAliasList is exited.
func (s *BaseStarRocksSQLListener) ExitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseStarRocksSQLListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseStarRocksSQLListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierOrString is called when production identifierOrString is entered.
func (s *BaseStarRocksSQLListener) EnterIdentifierOrString(ctx *IdentifierOrStringContext) {}

// ExitIdentifierOrString is called when production identifierOrString is exited.
func (s *BaseStarRocksSQLListener) ExitIdentifierOrString(ctx *IdentifierOrStringContext) {}

// EnterIdentifierOrStringList is called when production identifierOrStringList is entered.
func (s *BaseStarRocksSQLListener) EnterIdentifierOrStringList(ctx *IdentifierOrStringListContext) {}

// ExitIdentifierOrStringList is called when production identifierOrStringList is exited.
func (s *BaseStarRocksSQLListener) ExitIdentifierOrStringList(ctx *IdentifierOrStringListContext) {}

// EnterIdentifierOrStringOrStar is called when production identifierOrStringOrStar is entered.
func (s *BaseStarRocksSQLListener) EnterIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {
}

// ExitIdentifierOrStringOrStar is called when production identifierOrStringOrStar is exited.
func (s *BaseStarRocksSQLListener) ExitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {
}

// EnterUserWithoutHost is called when production userWithoutHost is entered.
func (s *BaseStarRocksSQLListener) EnterUserWithoutHost(ctx *UserWithoutHostContext) {}

// ExitUserWithoutHost is called when production userWithoutHost is exited.
func (s *BaseStarRocksSQLListener) ExitUserWithoutHost(ctx *UserWithoutHostContext) {}

// EnterUserWithHost is called when production userWithHost is entered.
func (s *BaseStarRocksSQLListener) EnterUserWithHost(ctx *UserWithHostContext) {}

// ExitUserWithHost is called when production userWithHost is exited.
func (s *BaseStarRocksSQLListener) ExitUserWithHost(ctx *UserWithHostContext) {}

// EnterUserWithHostAndBlanket is called when production userWithHostAndBlanket is entered.
func (s *BaseStarRocksSQLListener) EnterUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {}

// ExitUserWithHostAndBlanket is called when production userWithHostAndBlanket is exited.
func (s *BaseStarRocksSQLListener) ExitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseStarRocksSQLListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseStarRocksSQLListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseStarRocksSQLListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseStarRocksSQLListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterDecimalValue is called when production decimalValue is entered.
func (s *BaseStarRocksSQLListener) EnterDecimalValue(ctx *DecimalValueContext) {}

// ExitDecimalValue is called when production decimalValue is exited.
func (s *BaseStarRocksSQLListener) ExitDecimalValue(ctx *DecimalValueContext) {}

// EnterDoubleValue is called when production doubleValue is entered.
func (s *BaseStarRocksSQLListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production doubleValue is exited.
func (s *BaseStarRocksSQLListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseStarRocksSQLListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseStarRocksSQLListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseStarRocksSQLListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseStarRocksSQLListener) ExitNonReserved(ctx *NonReservedContext) {}
