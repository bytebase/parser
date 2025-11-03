// Code generated from TiDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tidb // TiDBParser
import "github.com/antlr4-go/antlr/v4"

type BaseTiDBParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTiDBParserVisitor) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleStatement(ctx *SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateStatement(ctx *CreateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDropStatement(ctx *DropStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewAlgorithm(ctx *ViewAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewSuid(ctx *ViewSuidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewTail(ctx *ViewTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewSelect(ctx *ViewSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewCheckOption(ctx *ViewCheckOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableElementList(ctx *TableElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableElement(ctx *TableElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableConstraintDef(ctx *TableConstraintDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexOption(ctx *IndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCheckConstraint(ctx *CheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexNameAndType(ctx *IndexNameAndTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTemporaryOption(ctx *TemporaryOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnDef(ctx *ColumnDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnOptionList(ctx *ColumnOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnOption(ctx *ColumnOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitConstraintEnforcement(ctx *ConstraintEnforcementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSelectStatementWithInto(ctx *SelectStatementWithIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQueryExpressionBody(ctx *QueryExpressionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQueryExpressionParens(ctx *QueryExpressionParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQueryPrimary(ctx *QueryPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQuerySpecOption(ctx *QuerySpecOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleLimitClause(ctx *SimpleLimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLimitOptions(ctx *LimitOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLimitOption(ctx *LimitOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIntoClause(ctx *IntoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowDefinition(ctx *WindowDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowSpec(ctx *WindowSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowSpecDetails(ctx *WindowSpecDetailsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFrameClause(ctx *WindowFrameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFrameUnits(ctx *WindowFrameUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFrameExtent(ctx *WindowFrameExtentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFrameStart(ctx *WindowFrameStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFrameBetween(ctx *WindowFrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFrameBound(ctx *WindowFrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFrameExclusion(ctx *WindowFrameExclusionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOlapOption(ctx *OlapOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOrderClause(ctx *OrderClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDirection(ctx *DirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableReferenceList(ctx *TableReferenceListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableValueConstructor(ctx *TableValueConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExplicitTable(ctx *ExplicitTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRowValueExplicit(ctx *RowValueExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitValues(ctx *ValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSelectOption(ctx *SelectOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLockingClauseList(ctx *LockingClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLockingClause(ctx *LockingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLockStrengh(ctx *LockStrenghContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLockedRowAction(ctx *LockedRowActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSelectItemList(ctx *SelectItemListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSelectItem(ctx *SelectItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSelectAlias(ctx *SelectAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableReference(ctx *TableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJoinedTable(ctx *JoinedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNaturalJoinType(ctx *NaturalJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitInnerJoinType(ctx *InnerJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableFactor(ctx *TableFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSingleTable(ctx *SingleTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSingleTableParens(ctx *SingleTableParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDerivedTable(ctx *DerivedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableReferenceListParens(ctx *TableReferenceListParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableFunction(ctx *TableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnsClause(ctx *ColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJtColumn(ctx *JtColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOnEmptyOrError(ctx *OnEmptyOrErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOnEmpty(ctx *OnEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOnError(ctx *OnErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJtOnResponse(ctx *JtOnResponseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetOprSymbol(ctx *SetOprSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetOprOption(ctx *SetOprOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableAlias(ctx *TableAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexHintList(ctx *IndexHintListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitKeyOrIndex(ctx *KeyOrIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitConstraintKeyType(ctx *ConstraintKeyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexHintClause(ctx *IndexHintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexList(ctx *IndexListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexListElement(ctx *IndexListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTransactionStatement(ctx *TransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitBeginWork(ctx *BeginWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTransactionCharacteristic(ctx *TransactionCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSavepointStatement(ctx *SavepointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLockStatement(ctx *LockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLockItem(ctx *LockItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLockOption(ctx *LockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitXaStatement(ctx *XaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitXaConvert(ctx *XaConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitXid(ctx *XidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitReplicationStatement(ctx *ReplicationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitResetOption(ctx *ResetOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitMasterResetOptions(ctx *MasterResetOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitReplicationLoad(ctx *ReplicationLoadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitChangeMaster(ctx *ChangeMasterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitChangeMasterOptions(ctx *ChangeMasterOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitMasterOption(ctx *MasterOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitMasterFileDef(ctx *MasterFileDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitServerIdList(ctx *ServerIdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitChangeReplication(ctx *ChangeReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFilterDefinition(ctx *FilterDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFilterDbList(ctx *FilterDbListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFilterTableList(ctx *FilterTableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFilterStringList(ctx *FilterStringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFilterWildDbTableString(ctx *FilterWildDbTableStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFilterDbPairList(ctx *FilterDbPairListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSlave(ctx *SlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSlaveUntilOptions(ctx *SlaveUntilOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSlaveThreadOptions(ctx *SlaveThreadOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSlaveThreadOption(ctx *SlaveThreadOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGroupReplication(ctx *GroupReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPreparedStatement(ctx *PreparedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExecuteVarList(ctx *ExecuteVarListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCloneStatement(ctx *CloneStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDataDirSSL(ctx *DataDirSSLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSsl(ctx *SslContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAccountManagementStatement(ctx *AccountManagementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAlterUser(ctx *AlterUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAlterUserTail(ctx *AlterUserTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUserFunction(ctx *UserFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateUser(ctx *CreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateUserTail(ctx *CreateUserTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDefaultRoleClause(ctx *DefaultRoleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRequireClause(ctx *RequireClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitConnectOptions(ctx *ConnectOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDropUser(ctx *DropUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGrant(ctx *GrantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGrantTargetList(ctx *GrantTargetListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGrantOptions(ctx *GrantOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExceptRoleList(ctx *ExceptRoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWithRoles(ctx *WithRolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGrantAs(ctx *GrantAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitVersionedRequireClause(ctx *VersionedRequireClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRenameUser(ctx *RenameUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRevoke(ctx *RevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOnTypeTo(ctx *OnTypeToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAclType(ctx *AclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleOrPrivilege(ctx *RoleOrPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGrantIdentifier(ctx *GrantIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRequireList(ctx *RequireListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRequireListElement(ctx *RequireListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGrantOption(ctx *GrantOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetRole(ctx *SetRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleList(ctx *RoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRole(ctx *RoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableAdministrationStatement(ctx *TableAdministrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitHistogram(ctx *HistogramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCheckOption(ctx *CheckOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRepairType(ctx *RepairTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitInstallUninstallStatment(ctx *InstallUninstallStatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetStatement(ctx *SetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitStartOptionValueList(ctx *StartOptionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIsolationLevel(ctx *IsolationLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOptionValueListContinued(ctx *OptionValueListContinuedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOptionValue(ctx *OptionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetSystemVariable(ctx *SetSystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitShowStatement(ctx *ShowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitShowCommandType(ctx *ShowCommandTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNonBlocking(ctx *NonBlockingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFromOrIn(ctx *FromOrInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitInDb(ctx *InDbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitProfileType(ctx *ProfileTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitResourceGroupManagement(ctx *ResourceGroupManagementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateResourceGroup(ctx *CreateResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitVcpuNumOrRange(ctx *VcpuNumOrRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitResourceGroupPriority(ctx *ResourceGroupPriorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAlterResourceGroup(ctx *AlterResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetResourceGroup(ctx *SetResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitThreadIdList(ctx *ThreadIdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDropResourceGroup(ctx *DropResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprOr(ctx *ExprOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprNot(ctx *ExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprIs(ctx *ExprIsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprAnd(ctx *ExprAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprXor(ctx *ExprXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPrimaryExprCompare(ctx *PrimaryExprCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCompOp(ctx *CompOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPredicateExprIn(ctx *PredicateExprInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPredicateExprBetween(ctx *PredicateExprBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPredicateExprLike(ctx *PredicateExprLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPredicateExprRegex(ctx *PredicateExprRegexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitBitExpr(ctx *BitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprConvert(ctx *SimpleExprConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprVariable(ctx *SimpleExprVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprCast(ctx *SimpleExprCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprUnary(ctx *SimpleExprUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprOdbc(ctx *SimpleExprOdbcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprFunction(ctx *SimpleExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprCollate(ctx *SimpleExprCollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprMatch(ctx *SimpleExprMatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprBinary(ctx *SimpleExprBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprSum(ctx *SimpleExprSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprNot(ctx *SimpleExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprValues(ctx *SimpleExprValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprDefault(ctx *SimpleExprDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprList(ctx *SimpleExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprInterval(ctx *SimpleExprIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprCase(ctx *SimpleExprCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprConcat(ctx *SimpleExprConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprLiteral(ctx *SimpleExprLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitArrayCast(ctx *ArrayCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJsonOperator(ctx *JsonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSumExpr(ctx *SumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowingClause(ctx *WindowingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLeadLagInfo(ctx *LeadLagInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNullTreatment(ctx *NullTreatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJsonFunction(ctx *JsonFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitInSumExpr(ctx *InSumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentListArg(ctx *IdentListArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentList(ctx *IdentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFulltextOptions(ctx *FulltextOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGeometryFunction(ctx *GeometryFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTimeFunctionParameters(ctx *TimeFunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFractionalPrecision(ctx *FractionalPrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWeightStringLevels(ctx *WeightStringLevelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDateTimeTtype(ctx *DateTimeTtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTrimFunction(ctx *TrimFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSubstringFunction(ctx *SubstringFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSearchJsonFunction(ctx *SearchJsonFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJsonValueReturning(ctx *JsonValueReturningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitJsonValueOnError(ctx *JsonValueOnErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUdfExprList(ctx *UdfExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUdfExpr(ctx *UdfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUserVariable(ctx *UserVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSystemVariable(ctx *SystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitInternalVariableName(ctx *InternalVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWhenExpression(ctx *WhenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitThenExpression(ctx *ThenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitElseExpression(ctx *ElseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCastType(ctx *CastTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCharset(ctx *CharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNotRule(ctx *NotRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNot2Rule(ctx *Not2RuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIntervalTimeStamp(ctx *IntervalTimeStampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprListWithParentheses(ctx *ExprListWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitExprWithParentheses(ctx *ExprWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOrderList(ctx *OrderListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOrderExpression(ctx *OrderExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGroupList(ctx *GroupListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGroupingExpression(ctx *GroupingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitChannel(ctx *ChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnFormat(ctx *ColumnFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitStorageMedia(ctx *StorageMediaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitGcolAttribute(ctx *GcolAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitReferences(ctx *ReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDeleteOption(ctx *DeleteOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitKeyList(ctx *KeyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitKeyPart(ctx *KeyPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitKeyListWithExpression(ctx *KeyListWithExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitKeyPartOrExpression(ctx *KeyPartOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitKeyListVariants(ctx *KeyListVariantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCommonIndexOption(ctx *CommonIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitVisibility(ctx *VisibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexTypeClause(ctx *IndexTypeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFulltextIndexOption(ctx *FulltextIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSpatialIndexOption(ctx *SpatialIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDataTypeDefinition(ctx *DataTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNchar(ctx *NcharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRealType(ctx *RealTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAutoRandomFieldLength(ctx *AutoRandomFieldLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFieldLength(ctx *FieldLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFieldOptions(ctx *FieldOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAscii(ctx *AsciiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUnicode(ctx *UnicodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWsNumCodepoints(ctx *WsNumCodepointsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateTableOption(ctx *CreateTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTernaryOption(ctx *TernaryOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDefaultCollation(ctx *DefaultCollationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDefaultEncryption(ctx *DefaultEncryptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDefaultCharset(ctx *DefaultCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionClause(ctx *PartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionDefKey(ctx *PartitionDefKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionDefHash(ctx *PartitionDefHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionDefRangeList(ctx *PartitionDefRangeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSubPartitions(ctx *SubPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionValuesIn(ctx *PartitionValuesInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionOption(ctx *PartitionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPartitionValueItem(ctx *PartitionValueItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDefinerClause(ctx *DefinerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIfExists(ctx *IfExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitProcedureParameter(ctx *ProcedureParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCollate(ctx *CollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTypeWithOptCollate(ctx *TypeWithOptCollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewRefList(ctx *ViewRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUpdateList(ctx *UpdateListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUpdateElement(ctx *UpdateElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCharsetClause(ctx *CharsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFieldsClause(ctx *FieldsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFieldTerm(ctx *FieldTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLinesClause(ctx *LinesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLineTerm(ctx *LineTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUserList(ctx *UserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateUserList(ctx *CreateUserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAlterUserList(ctx *AlterUserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitCreateUserEntry(ctx *CreateUserEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitAlterUserEntry(ctx *AlterUserEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRetainCurrentPassword(ctx *RetainCurrentPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDiscardOldPassword(ctx *DiscardOldPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitReplacePassword(ctx *ReplacePasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUser(ctx *UserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLikeClause(ctx *LikeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLikeOrWhere(ctx *LikeOrWhereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOnlineOption(ctx *OnlineOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNoWriteToBinLog(ctx *NoWriteToBinLogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUsePartition(ctx *UsePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFieldIdentifier(ctx *FieldIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnInternalRef(ctx *ColumnInternalRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnInternalRefList(ctx *ColumnInternalRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitColumnRef(ctx *ColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitInsertIdentifier(ctx *InsertIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIndexRef(ctx *IndexRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableWild(ctx *TableWildContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSchemaName(ctx *SchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSchemaRef(ctx *SchemaRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitProcedureName(ctx *ProcedureNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitProcedureRef(ctx *ProcedureRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFunctionRef(ctx *FunctionRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTriggerName(ctx *TriggerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTriggerRef(ctx *TriggerRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitViewRef(ctx *ViewRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTablespaceName(ctx *TablespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTablespaceRef(ctx *TablespaceRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLogfileGroupName(ctx *LogfileGroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLogfileGroupRef(ctx *LogfileGroupRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitEventName(ctx *EventNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitEventRef(ctx *EventRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUdfName(ctx *UdfNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitServerName(ctx *ServerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitServerRef(ctx *ServerRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitEngineRef(ctx *EngineRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFilterTableRef(ctx *FilterTableRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableRefWithWildcard(ctx *TableRefWithWildcardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableRef(ctx *TableRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableRefList(ctx *TableRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTableAliasRefList(ctx *TableAliasRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitParameterName(ctx *ParameterNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLabelIdentifier(ctx *LabelIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLabelRef(ctx *LabelRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleIdentifier(ctx *RoleIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleRef(ctx *RoleRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPluginRef(ctx *PluginRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitComponentRef(ctx *ComponentRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitResourceGroupRef(ctx *ResourceGroupRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitWindowName(ctx *WindowNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPureIdentifier(ctx *PureIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitDotIdentifier(ctx *DotIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUlong_number(ctx *Ulong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitReal_ulong_number(ctx *Real_ulong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitUlonglong_number(ctx *Ulonglong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitReal_ulonglong_number(ctx *Real_ulonglong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSignedLiteral(ctx *SignedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTextStringLiteral(ctx *TextStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTextString(ctx *TextStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTextStringHash(ctx *TextStringHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTextLiteral(ctx *TextLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTextStringLiteralList(ctx *TextStringLiteralListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNumLiteral(ctx *NumLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTemporalLiteral(ctx *TemporalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitFloatOptions(ctx *FloatOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitStandardFloatOptions(ctx *StandardFloatOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitPrecision(ctx *PrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitTextOrIdentifier(ctx *TextOrIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLValueIdentifier(ctx *LValueIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSizeNumber(ctx *SizeNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitParentheses(ctx *ParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitEqual(ctx *EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitOptionType(ctx *OptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitVarIdentType(ctx *VarIdentTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitSetVarIdentType(ctx *SetVarIdentTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierKeyword(ctx *IdentifierKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLabelKeyword(ctx *LabelKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleKeyword(ctx *RoleKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitLValueKeyword(ctx *LValueKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTiDBParserVisitor) VisitRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
