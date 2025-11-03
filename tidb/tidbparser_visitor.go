// Code generated from TiDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tidb // TiDBParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TiDBParser.
type TiDBParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TiDBParser#script.
	VisitScript(ctx *ScriptContext) interface{}

	// Visit a parse tree produced by TiDBParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#createStatement.
	VisitCreateStatement(ctx *CreateStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#dropStatement.
	VisitDropStatement(ctx *DropStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by TiDBParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by TiDBParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewReplaceOrAlgorithm.
	VisitViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewAlgorithm.
	VisitViewAlgorithm(ctx *ViewAlgorithmContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewSuid.
	VisitViewSuid(ctx *ViewSuidContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewTail.
	VisitViewTail(ctx *ViewTailContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewSelect.
	VisitViewSelect(ctx *ViewSelectContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewCheckOption.
	VisitViewCheckOption(ctx *ViewCheckOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#duplicateAsQueryExpression.
	VisitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#queryExpressionOrParens.
	VisitQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableElementList.
	VisitTableElementList(ctx *TableElementListContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableElement.
	VisitTableElement(ctx *TableElementContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableConstraintDef.
	VisitTableConstraintDef(ctx *TableConstraintDefContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexOption.
	VisitIndexOption(ctx *IndexOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#checkConstraint.
	VisitCheckConstraint(ctx *CheckConstraintContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexNameAndType.
	VisitIndexNameAndType(ctx *IndexNameAndTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#temporaryOption.
	VisitTemporaryOption(ctx *TemporaryOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnDef.
	VisitColumnDef(ctx *ColumnDefContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnOptionList.
	VisitColumnOptionList(ctx *ColumnOptionListContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnOption.
	VisitColumnOption(ctx *ColumnOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#constraintEnforcement.
	VisitConstraintEnforcement(ctx *ConstraintEnforcementContext) interface{}

	// Visit a parse tree produced by TiDBParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#selectStatementWithInto.
	VisitSelectStatementWithInto(ctx *SelectStatementWithIntoContext) interface{}

	// Visit a parse tree produced by TiDBParser#queryExpression.
	VisitQueryExpression(ctx *QueryExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#queryExpressionBody.
	VisitQueryExpressionBody(ctx *QueryExpressionBodyContext) interface{}

	// Visit a parse tree produced by TiDBParser#queryExpressionParens.
	VisitQueryExpressionParens(ctx *QueryExpressionParensContext) interface{}

	// Visit a parse tree produced by TiDBParser#queryPrimary.
	VisitQueryPrimary(ctx *QueryPrimaryContext) interface{}

	// Visit a parse tree produced by TiDBParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by TiDBParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by TiDBParser#querySpecOption.
	VisitQuerySpecOption(ctx *QuerySpecOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleLimitClause.
	VisitSimpleLimitClause(ctx *SimpleLimitClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#limitOptions.
	VisitLimitOptions(ctx *LimitOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#limitOption.
	VisitLimitOption(ctx *LimitOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#intoClause.
	VisitIntoClause(ctx *IntoClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#procedureAnalyseClause.
	VisitProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowDefinition.
	VisitWindowDefinition(ctx *WindowDefinitionContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowSpec.
	VisitWindowSpec(ctx *WindowSpecContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowSpecDetails.
	VisitWindowSpecDetails(ctx *WindowSpecDetailsContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFrameClause.
	VisitWindowFrameClause(ctx *WindowFrameClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFrameUnits.
	VisitWindowFrameUnits(ctx *WindowFrameUnitsContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFrameExtent.
	VisitWindowFrameExtent(ctx *WindowFrameExtentContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFrameStart.
	VisitWindowFrameStart(ctx *WindowFrameStartContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFrameBetween.
	VisitWindowFrameBetween(ctx *WindowFrameBetweenContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFrameBound.
	VisitWindowFrameBound(ctx *WindowFrameBoundContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFrameExclusion.
	VisitWindowFrameExclusion(ctx *WindowFrameExclusionContext) interface{}

	// Visit a parse tree produced by TiDBParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#commonTableExpression.
	VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#olapOption.
	VisitOlapOption(ctx *OlapOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#orderClause.
	VisitOrderClause(ctx *OrderClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#direction.
	VisitDirection(ctx *DirectionContext) interface{}

	// Visit a parse tree produced by TiDBParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableReferenceList.
	VisitTableReferenceList(ctx *TableReferenceListContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableValueConstructor.
	VisitTableValueConstructor(ctx *TableValueConstructorContext) interface{}

	// Visit a parse tree produced by TiDBParser#explicitTable.
	VisitExplicitTable(ctx *ExplicitTableContext) interface{}

	// Visit a parse tree produced by TiDBParser#rowValueExplicit.
	VisitRowValueExplicit(ctx *RowValueExplicitContext) interface{}

	// Visit a parse tree produced by TiDBParser#values.
	VisitValues(ctx *ValuesContext) interface{}

	// Visit a parse tree produced by TiDBParser#selectOption.
	VisitSelectOption(ctx *SelectOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#lockingClauseList.
	VisitLockingClauseList(ctx *LockingClauseListContext) interface{}

	// Visit a parse tree produced by TiDBParser#lockingClause.
	VisitLockingClause(ctx *LockingClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#lockStrengh.
	VisitLockStrengh(ctx *LockStrenghContext) interface{}

	// Visit a parse tree produced by TiDBParser#lockedRowAction.
	VisitLockedRowAction(ctx *LockedRowActionContext) interface{}

	// Visit a parse tree produced by TiDBParser#selectItemList.
	VisitSelectItemList(ctx *SelectItemListContext) interface{}

	// Visit a parse tree produced by TiDBParser#selectItem.
	VisitSelectItem(ctx *SelectItemContext) interface{}

	// Visit a parse tree produced by TiDBParser#selectAlias.
	VisitSelectAlias(ctx *SelectAliasContext) interface{}

	// Visit a parse tree produced by TiDBParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableReference.
	VisitTableReference(ctx *TableReferenceContext) interface{}

	// Visit a parse tree produced by TiDBParser#escapedTableReference.
	VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{}

	// Visit a parse tree produced by TiDBParser#joinedTable.
	VisitJoinedTable(ctx *JoinedTableContext) interface{}

	// Visit a parse tree produced by TiDBParser#naturalJoinType.
	VisitNaturalJoinType(ctx *NaturalJoinTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#innerJoinType.
	VisitInnerJoinType(ctx *InnerJoinTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#outerJoinType.
	VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableFactor.
	VisitTableFactor(ctx *TableFactorContext) interface{}

	// Visit a parse tree produced by TiDBParser#singleTable.
	VisitSingleTable(ctx *SingleTableContext) interface{}

	// Visit a parse tree produced by TiDBParser#singleTableParens.
	VisitSingleTableParens(ctx *SingleTableParensContext) interface{}

	// Visit a parse tree produced by TiDBParser#derivedTable.
	VisitDerivedTable(ctx *DerivedTableContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableReferenceListParens.
	VisitTableReferenceListParens(ctx *TableReferenceListParensContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableFunction.
	VisitTableFunction(ctx *TableFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnsClause.
	VisitColumnsClause(ctx *ColumnsClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#jtColumn.
	VisitJtColumn(ctx *JtColumnContext) interface{}

	// Visit a parse tree produced by TiDBParser#onEmptyOrError.
	VisitOnEmptyOrError(ctx *OnEmptyOrErrorContext) interface{}

	// Visit a parse tree produced by TiDBParser#onEmpty.
	VisitOnEmpty(ctx *OnEmptyContext) interface{}

	// Visit a parse tree produced by TiDBParser#onError.
	VisitOnError(ctx *OnErrorContext) interface{}

	// Visit a parse tree produced by TiDBParser#jtOnResponse.
	VisitJtOnResponse(ctx *JtOnResponseContext) interface{}

	// Visit a parse tree produced by TiDBParser#setOprSymbol.
	VisitSetOprSymbol(ctx *SetOprSymbolContext) interface{}

	// Visit a parse tree produced by TiDBParser#setOprOption.
	VisitSetOprOption(ctx *SetOprOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableAlias.
	VisitTableAlias(ctx *TableAliasContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexHintList.
	VisitIndexHintList(ctx *IndexHintListContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexHint.
	VisitIndexHint(ctx *IndexHintContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexHintType.
	VisitIndexHintType(ctx *IndexHintTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#keyOrIndex.
	VisitKeyOrIndex(ctx *KeyOrIndexContext) interface{}

	// Visit a parse tree produced by TiDBParser#constraintKeyType.
	VisitConstraintKeyType(ctx *ConstraintKeyTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexHintClause.
	VisitIndexHintClause(ctx *IndexHintClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexList.
	VisitIndexList(ctx *IndexListContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexListElement.
	VisitIndexListElement(ctx *IndexListElementContext) interface{}

	// Visit a parse tree produced by TiDBParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#transactionOrLockingStatement.
	VisitTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#transactionStatement.
	VisitTransactionStatement(ctx *TransactionStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#beginWork.
	VisitBeginWork(ctx *BeginWorkContext) interface{}

	// Visit a parse tree produced by TiDBParser#transactionCharacteristic.
	VisitTransactionCharacteristic(ctx *TransactionCharacteristicContext) interface{}

	// Visit a parse tree produced by TiDBParser#savepointStatement.
	VisitSavepointStatement(ctx *SavepointStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#lockStatement.
	VisitLockStatement(ctx *LockStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#lockItem.
	VisitLockItem(ctx *LockItemContext) interface{}

	// Visit a parse tree produced by TiDBParser#lockOption.
	VisitLockOption(ctx *LockOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#xaStatement.
	VisitXaStatement(ctx *XaStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#xaConvert.
	VisitXaConvert(ctx *XaConvertContext) interface{}

	// Visit a parse tree produced by TiDBParser#xid.
	VisitXid(ctx *XidContext) interface{}

	// Visit a parse tree produced by TiDBParser#replicationStatement.
	VisitReplicationStatement(ctx *ReplicationStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#resetOption.
	VisitResetOption(ctx *ResetOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#masterResetOptions.
	VisitMasterResetOptions(ctx *MasterResetOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#replicationLoad.
	VisitReplicationLoad(ctx *ReplicationLoadContext) interface{}

	// Visit a parse tree produced by TiDBParser#changeMaster.
	VisitChangeMaster(ctx *ChangeMasterContext) interface{}

	// Visit a parse tree produced by TiDBParser#changeMasterOptions.
	VisitChangeMasterOptions(ctx *ChangeMasterOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#masterOption.
	VisitMasterOption(ctx *MasterOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#privilegeCheckDef.
	VisitPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) interface{}

	// Visit a parse tree produced by TiDBParser#tablePrimaryKeyCheckDef.
	VisitTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) interface{}

	// Visit a parse tree produced by TiDBParser#masterTlsCiphersuitesDef.
	VisitMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) interface{}

	// Visit a parse tree produced by TiDBParser#masterFileDef.
	VisitMasterFileDef(ctx *MasterFileDefContext) interface{}

	// Visit a parse tree produced by TiDBParser#serverIdList.
	VisitServerIdList(ctx *ServerIdListContext) interface{}

	// Visit a parse tree produced by TiDBParser#changeReplication.
	VisitChangeReplication(ctx *ChangeReplicationContext) interface{}

	// Visit a parse tree produced by TiDBParser#filterDefinition.
	VisitFilterDefinition(ctx *FilterDefinitionContext) interface{}

	// Visit a parse tree produced by TiDBParser#filterDbList.
	VisitFilterDbList(ctx *FilterDbListContext) interface{}

	// Visit a parse tree produced by TiDBParser#filterTableList.
	VisitFilterTableList(ctx *FilterTableListContext) interface{}

	// Visit a parse tree produced by TiDBParser#filterStringList.
	VisitFilterStringList(ctx *FilterStringListContext) interface{}

	// Visit a parse tree produced by TiDBParser#filterWildDbTableString.
	VisitFilterWildDbTableString(ctx *FilterWildDbTableStringContext) interface{}

	// Visit a parse tree produced by TiDBParser#filterDbPairList.
	VisitFilterDbPairList(ctx *FilterDbPairListContext) interface{}

	// Visit a parse tree produced by TiDBParser#slave.
	VisitSlave(ctx *SlaveContext) interface{}

	// Visit a parse tree produced by TiDBParser#slaveUntilOptions.
	VisitSlaveUntilOptions(ctx *SlaveUntilOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#slaveConnectionOptions.
	VisitSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#slaveThreadOptions.
	VisitSlaveThreadOptions(ctx *SlaveThreadOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#slaveThreadOption.
	VisitSlaveThreadOption(ctx *SlaveThreadOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#groupReplication.
	VisitGroupReplication(ctx *GroupReplicationContext) interface{}

	// Visit a parse tree produced by TiDBParser#preparedStatement.
	VisitPreparedStatement(ctx *PreparedStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#executeVarList.
	VisitExecuteVarList(ctx *ExecuteVarListContext) interface{}

	// Visit a parse tree produced by TiDBParser#cloneStatement.
	VisitCloneStatement(ctx *CloneStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#dataDirSSL.
	VisitDataDirSSL(ctx *DataDirSSLContext) interface{}

	// Visit a parse tree produced by TiDBParser#ssl.
	VisitSsl(ctx *SslContext) interface{}

	// Visit a parse tree produced by TiDBParser#accountManagementStatement.
	VisitAccountManagementStatement(ctx *AccountManagementStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#alterUser.
	VisitAlterUser(ctx *AlterUserContext) interface{}

	// Visit a parse tree produced by TiDBParser#alterUserTail.
	VisitAlterUserTail(ctx *AlterUserTailContext) interface{}

	// Visit a parse tree produced by TiDBParser#userFunction.
	VisitUserFunction(ctx *UserFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#createUser.
	VisitCreateUser(ctx *CreateUserContext) interface{}

	// Visit a parse tree produced by TiDBParser#createUserTail.
	VisitCreateUserTail(ctx *CreateUserTailContext) interface{}

	// Visit a parse tree produced by TiDBParser#defaultRoleClause.
	VisitDefaultRoleClause(ctx *DefaultRoleClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#requireClause.
	VisitRequireClause(ctx *RequireClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#connectOptions.
	VisitConnectOptions(ctx *ConnectOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#accountLockPasswordExpireOptions.
	VisitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#dropUser.
	VisitDropUser(ctx *DropUserContext) interface{}

	// Visit a parse tree produced by TiDBParser#grant.
	VisitGrant(ctx *GrantContext) interface{}

	// Visit a parse tree produced by TiDBParser#grantTargetList.
	VisitGrantTargetList(ctx *GrantTargetListContext) interface{}

	// Visit a parse tree produced by TiDBParser#grantOptions.
	VisitGrantOptions(ctx *GrantOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#exceptRoleList.
	VisitExceptRoleList(ctx *ExceptRoleListContext) interface{}

	// Visit a parse tree produced by TiDBParser#withRoles.
	VisitWithRoles(ctx *WithRolesContext) interface{}

	// Visit a parse tree produced by TiDBParser#grantAs.
	VisitGrantAs(ctx *GrantAsContext) interface{}

	// Visit a parse tree produced by TiDBParser#versionedRequireClause.
	VisitVersionedRequireClause(ctx *VersionedRequireClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#renameUser.
	VisitRenameUser(ctx *RenameUserContext) interface{}

	// Visit a parse tree produced by TiDBParser#revoke.
	VisitRevoke(ctx *RevokeContext) interface{}

	// Visit a parse tree produced by TiDBParser#onTypeTo.
	VisitOnTypeTo(ctx *OnTypeToContext) interface{}

	// Visit a parse tree produced by TiDBParser#aclType.
	VisitAclType(ctx *AclTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleOrPrivilegesList.
	VisitRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleOrPrivilege.
	VisitRoleOrPrivilege(ctx *RoleOrPrivilegeContext) interface{}

	// Visit a parse tree produced by TiDBParser#grantIdentifier.
	VisitGrantIdentifier(ctx *GrantIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#requireList.
	VisitRequireList(ctx *RequireListContext) interface{}

	// Visit a parse tree produced by TiDBParser#requireListElement.
	VisitRequireListElement(ctx *RequireListElementContext) interface{}

	// Visit a parse tree produced by TiDBParser#grantOption.
	VisitGrantOption(ctx *GrantOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#setRole.
	VisitSetRole(ctx *SetRoleContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleList.
	VisitRoleList(ctx *RoleListContext) interface{}

	// Visit a parse tree produced by TiDBParser#role.
	VisitRole(ctx *RoleContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableAdministrationStatement.
	VisitTableAdministrationStatement(ctx *TableAdministrationStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#histogram.
	VisitHistogram(ctx *HistogramContext) interface{}

	// Visit a parse tree produced by TiDBParser#checkOption.
	VisitCheckOption(ctx *CheckOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#repairType.
	VisitRepairType(ctx *RepairTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#installUninstallStatment.
	VisitInstallUninstallStatment(ctx *InstallUninstallStatmentContext) interface{}

	// Visit a parse tree produced by TiDBParser#setStatement.
	VisitSetStatement(ctx *SetStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#startOptionValueList.
	VisitStartOptionValueList(ctx *StartOptionValueListContext) interface{}

	// Visit a parse tree produced by TiDBParser#transactionCharacteristics.
	VisitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) interface{}

	// Visit a parse tree produced by TiDBParser#transactionAccessMode.
	VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{}

	// Visit a parse tree produced by TiDBParser#isolationLevel.
	VisitIsolationLevel(ctx *IsolationLevelContext) interface{}

	// Visit a parse tree produced by TiDBParser#optionValueListContinued.
	VisitOptionValueListContinued(ctx *OptionValueListContinuedContext) interface{}

	// Visit a parse tree produced by TiDBParser#optionValueNoOptionType.
	VisitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#optionValue.
	VisitOptionValue(ctx *OptionValueContext) interface{}

	// Visit a parse tree produced by TiDBParser#setSystemVariable.
	VisitSetSystemVariable(ctx *SetSystemVariableContext) interface{}

	// Visit a parse tree produced by TiDBParser#startOptionValueListFollowingOptionType.
	VisitStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#optionValueFollowingOptionType.
	VisitOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#setExprOrDefault.
	VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{}

	// Visit a parse tree produced by TiDBParser#showStatement.
	VisitShowStatement(ctx *ShowStatementContext) interface{}

	// Visit a parse tree produced by TiDBParser#showCommandType.
	VisitShowCommandType(ctx *ShowCommandTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#nonBlocking.
	VisitNonBlocking(ctx *NonBlockingContext) interface{}

	// Visit a parse tree produced by TiDBParser#fromOrIn.
	VisitFromOrIn(ctx *FromOrInContext) interface{}

	// Visit a parse tree produced by TiDBParser#inDb.
	VisitInDb(ctx *InDbContext) interface{}

	// Visit a parse tree produced by TiDBParser#profileType.
	VisitProfileType(ctx *ProfileTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#resourceGroupManagement.
	VisitResourceGroupManagement(ctx *ResourceGroupManagementContext) interface{}

	// Visit a parse tree produced by TiDBParser#createResourceGroup.
	VisitCreateResourceGroup(ctx *CreateResourceGroupContext) interface{}

	// Visit a parse tree produced by TiDBParser#resourceGroupVcpuList.
	VisitResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) interface{}

	// Visit a parse tree produced by TiDBParser#vcpuNumOrRange.
	VisitVcpuNumOrRange(ctx *VcpuNumOrRangeContext) interface{}

	// Visit a parse tree produced by TiDBParser#resourceGroupPriority.
	VisitResourceGroupPriority(ctx *ResourceGroupPriorityContext) interface{}

	// Visit a parse tree produced by TiDBParser#resourceGroupEnableDisable.
	VisitResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) interface{}

	// Visit a parse tree produced by TiDBParser#alterResourceGroup.
	VisitAlterResourceGroup(ctx *AlterResourceGroupContext) interface{}

	// Visit a parse tree produced by TiDBParser#setResourceGroup.
	VisitSetResourceGroup(ctx *SetResourceGroupContext) interface{}

	// Visit a parse tree produced by TiDBParser#threadIdList.
	VisitThreadIdList(ctx *ThreadIdListContext) interface{}

	// Visit a parse tree produced by TiDBParser#dropResourceGroup.
	VisitDropResourceGroup(ctx *DropResourceGroupContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprOr.
	VisitExprOr(ctx *ExprOrContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprNot.
	VisitExprNot(ctx *ExprNotContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprIs.
	VisitExprIs(ctx *ExprIsContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprAnd.
	VisitExprAnd(ctx *ExprAndContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprXor.
	VisitExprXor(ctx *ExprXorContext) interface{}

	// Visit a parse tree produced by TiDBParser#primaryExprPredicate.
	VisitPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) interface{}

	// Visit a parse tree produced by TiDBParser#primaryExprCompare.
	VisitPrimaryExprCompare(ctx *PrimaryExprCompareContext) interface{}

	// Visit a parse tree produced by TiDBParser#primaryExprAllAny.
	VisitPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) interface{}

	// Visit a parse tree produced by TiDBParser#primaryExprIsNull.
	VisitPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) interface{}

	// Visit a parse tree produced by TiDBParser#compOp.
	VisitCompOp(ctx *CompOpContext) interface{}

	// Visit a parse tree produced by TiDBParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by TiDBParser#predicateExprIn.
	VisitPredicateExprIn(ctx *PredicateExprInContext) interface{}

	// Visit a parse tree produced by TiDBParser#predicateExprBetween.
	VisitPredicateExprBetween(ctx *PredicateExprBetweenContext) interface{}

	// Visit a parse tree produced by TiDBParser#predicateExprLike.
	VisitPredicateExprLike(ctx *PredicateExprLikeContext) interface{}

	// Visit a parse tree produced by TiDBParser#predicateExprRegex.
	VisitPredicateExprRegex(ctx *PredicateExprRegexContext) interface{}

	// Visit a parse tree produced by TiDBParser#bitExpr.
	VisitBitExpr(ctx *BitExprContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprConvert.
	VisitSimpleExprConvert(ctx *SimpleExprConvertContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprSearchJson.
	VisitSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprVariable.
	VisitSimpleExprVariable(ctx *SimpleExprVariableContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprCast.
	VisitSimpleExprCast(ctx *SimpleExprCastContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprUnary.
	VisitSimpleExprUnary(ctx *SimpleExprUnaryContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprOdbc.
	VisitSimpleExprOdbc(ctx *SimpleExprOdbcContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprRuntimeFunction.
	VisitSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprFunction.
	VisitSimpleExprFunction(ctx *SimpleExprFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprCollate.
	VisitSimpleExprCollate(ctx *SimpleExprCollateContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprMatch.
	VisitSimpleExprMatch(ctx *SimpleExprMatchContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprWindowingFunction.
	VisitSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprBinary.
	VisitSimpleExprBinary(ctx *SimpleExprBinaryContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprColumnRef.
	VisitSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprParamMarker.
	VisitSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprSum.
	VisitSimpleExprSum(ctx *SimpleExprSumContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprConvertUsing.
	VisitSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprSubQuery.
	VisitSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprGroupingOperation.
	VisitSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprNot.
	VisitSimpleExprNot(ctx *SimpleExprNotContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprValues.
	VisitSimpleExprValues(ctx *SimpleExprValuesContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprDefault.
	VisitSimpleExprDefault(ctx *SimpleExprDefaultContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprList.
	VisitSimpleExprList(ctx *SimpleExprListContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprInterval.
	VisitSimpleExprInterval(ctx *SimpleExprIntervalContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprCase.
	VisitSimpleExprCase(ctx *SimpleExprCaseContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprConcat.
	VisitSimpleExprConcat(ctx *SimpleExprConcatContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprLiteral.
	VisitSimpleExprLiteral(ctx *SimpleExprLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#arrayCast.
	VisitArrayCast(ctx *ArrayCastContext) interface{}

	// Visit a parse tree produced by TiDBParser#jsonOperator.
	VisitJsonOperator(ctx *JsonOperatorContext) interface{}

	// Visit a parse tree produced by TiDBParser#sumExpr.
	VisitSumExpr(ctx *SumExprContext) interface{}

	// Visit a parse tree produced by TiDBParser#groupingOperation.
	VisitGroupingOperation(ctx *GroupingOperationContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowFunctionCall.
	VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowingClause.
	VisitWindowingClause(ctx *WindowingClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#leadLagInfo.
	VisitLeadLagInfo(ctx *LeadLagInfoContext) interface{}

	// Visit a parse tree produced by TiDBParser#nullTreatment.
	VisitNullTreatment(ctx *NullTreatmentContext) interface{}

	// Visit a parse tree produced by TiDBParser#jsonFunction.
	VisitJsonFunction(ctx *JsonFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#inSumExpr.
	VisitInSumExpr(ctx *InSumExprContext) interface{}

	// Visit a parse tree produced by TiDBParser#identListArg.
	VisitIdentListArg(ctx *IdentListArgContext) interface{}

	// Visit a parse tree produced by TiDBParser#identList.
	VisitIdentList(ctx *IdentListContext) interface{}

	// Visit a parse tree produced by TiDBParser#fulltextOptions.
	VisitFulltextOptions(ctx *FulltextOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#runtimeFunctionCall.
	VisitRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) interface{}

	// Visit a parse tree produced by TiDBParser#geometryFunction.
	VisitGeometryFunction(ctx *GeometryFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#timeFunctionParameters.
	VisitTimeFunctionParameters(ctx *TimeFunctionParametersContext) interface{}

	// Visit a parse tree produced by TiDBParser#fractionalPrecision.
	VisitFractionalPrecision(ctx *FractionalPrecisionContext) interface{}

	// Visit a parse tree produced by TiDBParser#weightStringLevels.
	VisitWeightStringLevels(ctx *WeightStringLevelsContext) interface{}

	// Visit a parse tree produced by TiDBParser#weightStringLevelListItem.
	VisitWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) interface{}

	// Visit a parse tree produced by TiDBParser#dateTimeTtype.
	VisitDateTimeTtype(ctx *DateTimeTtypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#trimFunction.
	VisitTrimFunction(ctx *TrimFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#substringFunction.
	VisitSubstringFunction(ctx *SubstringFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by TiDBParser#searchJsonFunction.
	VisitSearchJsonFunction(ctx *SearchJsonFunctionContext) interface{}

	// Visit a parse tree produced by TiDBParser#jsonValueReturning.
	VisitJsonValueReturning(ctx *JsonValueReturningContext) interface{}

	// Visit a parse tree produced by TiDBParser#jsonValueOnEmpty.
	VisitJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) interface{}

	// Visit a parse tree produced by TiDBParser#jsonValueOnError.
	VisitJsonValueOnError(ctx *JsonValueOnErrorContext) interface{}

	// Visit a parse tree produced by TiDBParser#udfExprList.
	VisitUdfExprList(ctx *UdfExprListContext) interface{}

	// Visit a parse tree produced by TiDBParser#udfExpr.
	VisitUdfExpr(ctx *UdfExprContext) interface{}

	// Visit a parse tree produced by TiDBParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by TiDBParser#userVariable.
	VisitUserVariable(ctx *UserVariableContext) interface{}

	// Visit a parse tree produced by TiDBParser#systemVariable.
	VisitSystemVariable(ctx *SystemVariableContext) interface{}

	// Visit a parse tree produced by TiDBParser#internalVariableName.
	VisitInternalVariableName(ctx *InternalVariableNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#whenExpression.
	VisitWhenExpression(ctx *WhenExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#thenExpression.
	VisitThenExpression(ctx *ThenExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#elseExpression.
	VisitElseExpression(ctx *ElseExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#castType.
	VisitCastType(ctx *CastTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by TiDBParser#charset.
	VisitCharset(ctx *CharsetContext) interface{}

	// Visit a parse tree produced by TiDBParser#notRule.
	VisitNotRule(ctx *NotRuleContext) interface{}

	// Visit a parse tree produced by TiDBParser#not2Rule.
	VisitNot2Rule(ctx *Not2RuleContext) interface{}

	// Visit a parse tree produced by TiDBParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by TiDBParser#intervalTimeStamp.
	VisitIntervalTimeStamp(ctx *IntervalTimeStampContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprListWithParentheses.
	VisitExprListWithParentheses(ctx *ExprListWithParenthesesContext) interface{}

	// Visit a parse tree produced by TiDBParser#exprWithParentheses.
	VisitExprWithParentheses(ctx *ExprWithParenthesesContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleExprWithParentheses.
	VisitSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) interface{}

	// Visit a parse tree produced by TiDBParser#orderList.
	VisitOrderList(ctx *OrderListContext) interface{}

	// Visit a parse tree produced by TiDBParser#orderExpression.
	VisitOrderExpression(ctx *OrderExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#groupList.
	VisitGroupList(ctx *GroupListContext) interface{}

	// Visit a parse tree produced by TiDBParser#groupingExpression.
	VisitGroupingExpression(ctx *GroupingExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#channel.
	VisitChannel(ctx *ChannelContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnFormat.
	VisitColumnFormat(ctx *ColumnFormatContext) interface{}

	// Visit a parse tree produced by TiDBParser#storageMedia.
	VisitStorageMedia(ctx *StorageMediaContext) interface{}

	// Visit a parse tree produced by TiDBParser#gcolAttribute.
	VisitGcolAttribute(ctx *GcolAttributeContext) interface{}

	// Visit a parse tree produced by TiDBParser#references.
	VisitReferences(ctx *ReferencesContext) interface{}

	// Visit a parse tree produced by TiDBParser#deleteOption.
	VisitDeleteOption(ctx *DeleteOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#keyList.
	VisitKeyList(ctx *KeyListContext) interface{}

	// Visit a parse tree produced by TiDBParser#keyPart.
	VisitKeyPart(ctx *KeyPartContext) interface{}

	// Visit a parse tree produced by TiDBParser#keyListWithExpression.
	VisitKeyListWithExpression(ctx *KeyListWithExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#keyPartOrExpression.
	VisitKeyPartOrExpression(ctx *KeyPartOrExpressionContext) interface{}

	// Visit a parse tree produced by TiDBParser#keyListVariants.
	VisitKeyListVariants(ctx *KeyListVariantsContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#commonIndexOption.
	VisitCommonIndexOption(ctx *CommonIndexOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#visibility.
	VisitVisibility(ctx *VisibilityContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexTypeClause.
	VisitIndexTypeClause(ctx *IndexTypeClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#fulltextIndexOption.
	VisitFulltextIndexOption(ctx *FulltextIndexOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#spatialIndexOption.
	VisitSpatialIndexOption(ctx *SpatialIndexOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#dataTypeDefinition.
	VisitDataTypeDefinition(ctx *DataTypeDefinitionContext) interface{}

	// Visit a parse tree produced by TiDBParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#nchar.
	VisitNchar(ctx *NcharContext) interface{}

	// Visit a parse tree produced by TiDBParser#realType.
	VisitRealType(ctx *RealTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#autoRandomFieldLength.
	VisitAutoRandomFieldLength(ctx *AutoRandomFieldLengthContext) interface{}

	// Visit a parse tree produced by TiDBParser#fieldLength.
	VisitFieldLength(ctx *FieldLengthContext) interface{}

	// Visit a parse tree produced by TiDBParser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#charsetWithOptBinary.
	VisitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) interface{}

	// Visit a parse tree produced by TiDBParser#ascii.
	VisitAscii(ctx *AsciiContext) interface{}

	// Visit a parse tree produced by TiDBParser#unicode.
	VisitUnicode(ctx *UnicodeContext) interface{}

	// Visit a parse tree produced by TiDBParser#wsNumCodepoints.
	VisitWsNumCodepoints(ctx *WsNumCodepointsContext) interface{}

	// Visit a parse tree produced by TiDBParser#typeDatetimePrecision.
	VisitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) interface{}

	// Visit a parse tree produced by TiDBParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#collationName.
	VisitCollationName(ctx *CollationNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#createTableOptions.
	VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#createTableOptionsSpaceSeparated.
	VisitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) interface{}

	// Visit a parse tree produced by TiDBParser#createTableOption.
	VisitCreateTableOption(ctx *CreateTableOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#ternaryOption.
	VisitTernaryOption(ctx *TernaryOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#defaultCollation.
	VisitDefaultCollation(ctx *DefaultCollationContext) interface{}

	// Visit a parse tree produced by TiDBParser#defaultEncryption.
	VisitDefaultEncryption(ctx *DefaultEncryptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#defaultCharset.
	VisitDefaultCharset(ctx *DefaultCharsetContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionClause.
	VisitPartitionClause(ctx *PartitionClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionDefKey.
	VisitPartitionDefKey(ctx *PartitionDefKeyContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionDefHash.
	VisitPartitionDefHash(ctx *PartitionDefHashContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionDefRangeList.
	VisitPartitionDefRangeList(ctx *PartitionDefRangeListContext) interface{}

	// Visit a parse tree produced by TiDBParser#subPartitions.
	VisitSubPartitions(ctx *SubPartitionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionKeyAlgorithm.
	VisitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionDefinitions.
	VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionDefinition.
	VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionValuesIn.
	VisitPartitionValuesIn(ctx *PartitionValuesInContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionOption.
	VisitPartitionOption(ctx *PartitionOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#subpartitionDefinition.
	VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionValueItemListParen.
	VisitPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) interface{}

	// Visit a parse tree produced by TiDBParser#partitionValueItem.
	VisitPartitionValueItem(ctx *PartitionValueItemContext) interface{}

	// Visit a parse tree produced by TiDBParser#definerClause.
	VisitDefinerClause(ctx *DefinerClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#ifExists.
	VisitIfExists(ctx *IfExistsContext) interface{}

	// Visit a parse tree produced by TiDBParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by TiDBParser#procedureParameter.
	VisitProcedureParameter(ctx *ProcedureParameterContext) interface{}

	// Visit a parse tree produced by TiDBParser#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) interface{}

	// Visit a parse tree produced by TiDBParser#collate.
	VisitCollate(ctx *CollateContext) interface{}

	// Visit a parse tree produced by TiDBParser#typeWithOptCollate.
	VisitTypeWithOptCollate(ctx *TypeWithOptCollateContext) interface{}

	// Visit a parse tree produced by TiDBParser#schemaIdentifierPair.
	VisitSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewRefList.
	VisitViewRefList(ctx *ViewRefListContext) interface{}

	// Visit a parse tree produced by TiDBParser#updateList.
	VisitUpdateList(ctx *UpdateListContext) interface{}

	// Visit a parse tree produced by TiDBParser#updateElement.
	VisitUpdateElement(ctx *UpdateElementContext) interface{}

	// Visit a parse tree produced by TiDBParser#charsetClause.
	VisitCharsetClause(ctx *CharsetClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#fieldsClause.
	VisitFieldsClause(ctx *FieldsClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#fieldTerm.
	VisitFieldTerm(ctx *FieldTermContext) interface{}

	// Visit a parse tree produced by TiDBParser#linesClause.
	VisitLinesClause(ctx *LinesClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#lineTerm.
	VisitLineTerm(ctx *LineTermContext) interface{}

	// Visit a parse tree produced by TiDBParser#userList.
	VisitUserList(ctx *UserListContext) interface{}

	// Visit a parse tree produced by TiDBParser#createUserList.
	VisitCreateUserList(ctx *CreateUserListContext) interface{}

	// Visit a parse tree produced by TiDBParser#alterUserList.
	VisitAlterUserList(ctx *AlterUserListContext) interface{}

	// Visit a parse tree produced by TiDBParser#createUserEntry.
	VisitCreateUserEntry(ctx *CreateUserEntryContext) interface{}

	// Visit a parse tree produced by TiDBParser#alterUserEntry.
	VisitAlterUserEntry(ctx *AlterUserEntryContext) interface{}

	// Visit a parse tree produced by TiDBParser#retainCurrentPassword.
	VisitRetainCurrentPassword(ctx *RetainCurrentPasswordContext) interface{}

	// Visit a parse tree produced by TiDBParser#discardOldPassword.
	VisitDiscardOldPassword(ctx *DiscardOldPasswordContext) interface{}

	// Visit a parse tree produced by TiDBParser#replacePassword.
	VisitReplacePassword(ctx *ReplacePasswordContext) interface{}

	// Visit a parse tree produced by TiDBParser#userIdentifierOrText.
	VisitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) interface{}

	// Visit a parse tree produced by TiDBParser#user.
	VisitUser(ctx *UserContext) interface{}

	// Visit a parse tree produced by TiDBParser#likeClause.
	VisitLikeClause(ctx *LikeClauseContext) interface{}

	// Visit a parse tree produced by TiDBParser#likeOrWhere.
	VisitLikeOrWhere(ctx *LikeOrWhereContext) interface{}

	// Visit a parse tree produced by TiDBParser#onlineOption.
	VisitOnlineOption(ctx *OnlineOptionContext) interface{}

	// Visit a parse tree produced by TiDBParser#noWriteToBinLog.
	VisitNoWriteToBinLog(ctx *NoWriteToBinLogContext) interface{}

	// Visit a parse tree produced by TiDBParser#usePartition.
	VisitUsePartition(ctx *UsePartitionContext) interface{}

	// Visit a parse tree produced by TiDBParser#fieldIdentifier.
	VisitFieldIdentifier(ctx *FieldIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnInternalRef.
	VisitColumnInternalRef(ctx *ColumnInternalRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnInternalRefList.
	VisitColumnInternalRefList(ctx *ColumnInternalRefListContext) interface{}

	// Visit a parse tree produced by TiDBParser#columnRef.
	VisitColumnRef(ctx *ColumnRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#insertIdentifier.
	VisitInsertIdentifier(ctx *InsertIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#indexRef.
	VisitIndexRef(ctx *IndexRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableWild.
	VisitTableWild(ctx *TableWildContext) interface{}

	// Visit a parse tree produced by TiDBParser#schemaName.
	VisitSchemaName(ctx *SchemaNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#schemaRef.
	VisitSchemaRef(ctx *SchemaRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#procedureName.
	VisitProcedureName(ctx *ProcedureNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#procedureRef.
	VisitProcedureRef(ctx *ProcedureRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#functionRef.
	VisitFunctionRef(ctx *FunctionRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#triggerName.
	VisitTriggerName(ctx *TriggerNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#triggerRef.
	VisitTriggerRef(ctx *TriggerRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#viewRef.
	VisitViewRef(ctx *ViewRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#tablespaceName.
	VisitTablespaceName(ctx *TablespaceNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#tablespaceRef.
	VisitTablespaceRef(ctx *TablespaceRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#logfileGroupName.
	VisitLogfileGroupName(ctx *LogfileGroupNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#logfileGroupRef.
	VisitLogfileGroupRef(ctx *LogfileGroupRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#eventName.
	VisitEventName(ctx *EventNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#eventRef.
	VisitEventRef(ctx *EventRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#udfName.
	VisitUdfName(ctx *UdfNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#serverName.
	VisitServerName(ctx *ServerNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#serverRef.
	VisitServerRef(ctx *ServerRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#engineRef.
	VisitEngineRef(ctx *EngineRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#filterTableRef.
	VisitFilterTableRef(ctx *FilterTableRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableRefWithWildcard.
	VisitTableRefWithWildcard(ctx *TableRefWithWildcardContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableRef.
	VisitTableRef(ctx *TableRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableRefList.
	VisitTableRefList(ctx *TableRefListContext) interface{}

	// Visit a parse tree produced by TiDBParser#tableAliasRefList.
	VisitTableAliasRefList(ctx *TableAliasRefListContext) interface{}

	// Visit a parse tree produced by TiDBParser#parameterName.
	VisitParameterName(ctx *ParameterNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#labelIdentifier.
	VisitLabelIdentifier(ctx *LabelIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#labelRef.
	VisitLabelRef(ctx *LabelRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleIdentifier.
	VisitRoleIdentifier(ctx *RoleIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleRef.
	VisitRoleRef(ctx *RoleRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#pluginRef.
	VisitPluginRef(ctx *PluginRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#componentRef.
	VisitComponentRef(ctx *ComponentRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#resourceGroupRef.
	VisitResourceGroupRef(ctx *ResourceGroupRefContext) interface{}

	// Visit a parse tree produced by TiDBParser#windowName.
	VisitWindowName(ctx *WindowNameContext) interface{}

	// Visit a parse tree produced by TiDBParser#pureIdentifier.
	VisitPureIdentifier(ctx *PureIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierListWithParentheses.
	VisitIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) interface{}

	// Visit a parse tree produced by TiDBParser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#simpleIdentifier.
	VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#dotIdentifier.
	VisitDotIdentifier(ctx *DotIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#ulong_number.
	VisitUlong_number(ctx *Ulong_numberContext) interface{}

	// Visit a parse tree produced by TiDBParser#real_ulong_number.
	VisitReal_ulong_number(ctx *Real_ulong_numberContext) interface{}

	// Visit a parse tree produced by TiDBParser#ulonglong_number.
	VisitUlonglong_number(ctx *Ulonglong_numberContext) interface{}

	// Visit a parse tree produced by TiDBParser#real_ulonglong_number.
	VisitReal_ulonglong_number(ctx *Real_ulonglong_numberContext) interface{}

	// Visit a parse tree produced by TiDBParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#signedLiteral.
	VisitSignedLiteral(ctx *SignedLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by TiDBParser#textStringLiteral.
	VisitTextStringLiteral(ctx *TextStringLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#textString.
	VisitTextString(ctx *TextStringContext) interface{}

	// Visit a parse tree produced by TiDBParser#textStringHash.
	VisitTextStringHash(ctx *TextStringHashContext) interface{}

	// Visit a parse tree produced by TiDBParser#textLiteral.
	VisitTextLiteral(ctx *TextLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#textStringNoLinebreak.
	VisitTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) interface{}

	// Visit a parse tree produced by TiDBParser#textStringLiteralList.
	VisitTextStringLiteralList(ctx *TextStringLiteralListContext) interface{}

	// Visit a parse tree produced by TiDBParser#numLiteral.
	VisitNumLiteral(ctx *NumLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#boolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#temporalLiteral.
	VisitTemporalLiteral(ctx *TemporalLiteralContext) interface{}

	// Visit a parse tree produced by TiDBParser#floatOptions.
	VisitFloatOptions(ctx *FloatOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#standardFloatOptions.
	VisitStandardFloatOptions(ctx *StandardFloatOptionsContext) interface{}

	// Visit a parse tree produced by TiDBParser#precision.
	VisitPrecision(ctx *PrecisionContext) interface{}

	// Visit a parse tree produced by TiDBParser#textOrIdentifier.
	VisitTextOrIdentifier(ctx *TextOrIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#lValueIdentifier.
	VisitLValueIdentifier(ctx *LValueIdentifierContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleIdentifierOrText.
	VisitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) interface{}

	// Visit a parse tree produced by TiDBParser#sizeNumber.
	VisitSizeNumber(ctx *SizeNumberContext) interface{}

	// Visit a parse tree produced by TiDBParser#parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by TiDBParser#equal.
	VisitEqual(ctx *EqualContext) interface{}

	// Visit a parse tree produced by TiDBParser#optionType.
	VisitOptionType(ctx *OptionTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#varIdentType.
	VisitVarIdentType(ctx *VarIdentTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#setVarIdentType.
	VisitSetVarIdentType(ctx *SetVarIdentTypeContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierKeyword.
	VisitIdentifierKeyword(ctx *IdentifierKeywordContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierKeywordsAmbiguous1RolesAndLabels.
	VisitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierKeywordsAmbiguous2Labels.
	VisitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) interface{}

	// Visit a parse tree produced by TiDBParser#labelKeyword.
	VisitLabelKeyword(ctx *LabelKeywordContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierKeywordsAmbiguous3Roles.
	VisitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierKeywordsUnambiguous.
	VisitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleKeyword.
	VisitRoleKeyword(ctx *RoleKeywordContext) interface{}

	// Visit a parse tree produced by TiDBParser#lValueKeyword.
	VisitLValueKeyword(ctx *LValueKeywordContext) interface{}

	// Visit a parse tree produced by TiDBParser#identifierKeywordsAmbiguous4SystemVariables.
	VisitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleOrIdentifierKeyword.
	VisitRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) interface{}

	// Visit a parse tree produced by TiDBParser#roleOrLabelKeyword.
	VisitRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) interface{}
}
