// Code generated from MySQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mysql // MySQLParser
import "github.com/antlr4-go/antlr/v4"

/*
 * Copyright (c) 2018, 2020, Oracle and/or its affiliates. All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License, version 2.0,
 * as published by the Free Software Foundation.
 *
 * This program is also distributed with certain software (including
 * but not limited to OpenSSL) that is licensed under separate terms, as
 * designated in a particular file or component or in included license
 * documentation. The authors of MySQL hereby grant you an additional
 * permission to link the program and your derivative works with the
 * separately licensed software that they have included with MySQL.
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
 * the GNU General Public License, version 2.0, for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin St, Fifth Floor, Boston, MA 02110-1301 USA
 */

// A complete Visitor for a parse tree produced by MySQLParser.
type MySQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MySQLParser#script.
	VisitScript(ctx *ScriptContext) interface{}

	// Visit a parse tree produced by MySQLParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterStatement.
	VisitAlterStatement(ctx *AlterStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterDatabase.
	VisitAlterDatabase(ctx *AlterDatabaseContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterDatabaseOption.
	VisitAlterDatabaseOption(ctx *AlterDatabaseOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterEvent.
	VisitAlterEvent(ctx *AlterEventContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterLogfileGroup.
	VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterLogfileGroupOptions.
	VisitAlterLogfileGroupOptions(ctx *AlterLogfileGroupOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterLogfileGroupOption.
	VisitAlterLogfileGroupOption(ctx *AlterLogfileGroupOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterServer.
	VisitAlterServer(ctx *AlterServerContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterTableActions.
	VisitAlterTableActions(ctx *AlterTableActionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterCommandList.
	VisitAlterCommandList(ctx *AlterCommandListContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterCommandsModifierList.
	VisitAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) interface{}

	// Visit a parse tree produced by MySQLParser#standaloneAlterCommands.
	VisitStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterPartition.
	VisitAlterPartition(ctx *AlterPartitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterList.
	VisitAlterList(ctx *AlterListContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterCommandsModifier.
	VisitAlterCommandsModifier(ctx *AlterCommandsModifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterListItem.
	VisitAlterListItem(ctx *AlterListItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#place.
	VisitPlace(ctx *PlaceContext) interface{}

	// Visit a parse tree produced by MySQLParser#restrict.
	VisitRestrict(ctx *RestrictContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterOrderList.
	VisitAlterOrderList(ctx *AlterOrderListContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterAlgorithmOption.
	VisitAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterLockOption.
	VisitAlterLockOption(ctx *AlterLockOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexLockAndAlgorithm.
	VisitIndexLockAndAlgorithm(ctx *IndexLockAndAlgorithmContext) interface{}

	// Visit a parse tree produced by MySQLParser#withValidation.
	VisitWithValidation(ctx *WithValidationContext) interface{}

	// Visit a parse tree produced by MySQLParser#removePartitioning.
	VisitRemovePartitioning(ctx *RemovePartitioningContext) interface{}

	// Visit a parse tree produced by MySQLParser#allOrPartitionNameList.
	VisitAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterTablespace.
	VisitAlterTablespace(ctx *AlterTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterUndoTablespace.
	VisitAlterUndoTablespace(ctx *AlterUndoTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLParser#undoTableSpaceOptions.
	VisitUndoTableSpaceOptions(ctx *UndoTableSpaceOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#undoTableSpaceOption.
	VisitUndoTableSpaceOption(ctx *UndoTableSpaceOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterTablespaceOptions.
	VisitAlterTablespaceOptions(ctx *AlterTablespaceOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterTablespaceOption.
	VisitAlterTablespaceOption(ctx *AlterTablespaceOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#changeTablespaceOption.
	VisitChangeTablespaceOption(ctx *ChangeTablespaceOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterView.
	VisitAlterView(ctx *AlterViewContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewTail.
	VisitViewTail(ctx *ViewTailContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewSelect.
	VisitViewSelect(ctx *ViewSelectContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewCheckOption.
	VisitViewCheckOption(ctx *ViewCheckOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterProcedure.
	VisitAlterProcedure(ctx *AlterProcedureContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterFunction.
	VisitAlterFunction(ctx *AlterFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createStatement.
	VisitCreateStatement(ctx *CreateStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by MySQLParser#createDatabaseOption.
	VisitCreateDatabaseOption(ctx *CreateDatabaseOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableElementList.
	VisitTableElementList(ctx *TableElementListContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableElement.
	VisitTableElement(ctx *TableElementContext) interface{}

	// Visit a parse tree produced by MySQLParser#duplicateAsQueryExpression.
	VisitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#queryExpressionOrParens.
	VisitQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) interface{}

	// Visit a parse tree produced by MySQLParser#createRoutine.
	VisitCreateRoutine(ctx *CreateRoutineContext) interface{}

	// Visit a parse tree produced by MySQLParser#createProcedure.
	VisitCreateProcedure(ctx *CreateProcedureContext) interface{}

	// Visit a parse tree produced by MySQLParser#createFunction.
	VisitCreateFunction(ctx *CreateFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createUdf.
	VisitCreateUdf(ctx *CreateUdfContext) interface{}

	// Visit a parse tree produced by MySQLParser#routineCreateOption.
	VisitRoutineCreateOption(ctx *RoutineCreateOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#routineAlterOptions.
	VisitRoutineAlterOptions(ctx *RoutineAlterOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#routineOption.
	VisitRoutineOption(ctx *RoutineOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createIndex.
	VisitCreateIndex(ctx *CreateIndexContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexNameAndType.
	VisitIndexNameAndType(ctx *IndexNameAndTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#createIndexTarget.
	VisitCreateIndexTarget(ctx *CreateIndexTargetContext) interface{}

	// Visit a parse tree produced by MySQLParser#createLogfileGroup.
	VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#logfileGroupOptions.
	VisitLogfileGroupOptions(ctx *LogfileGroupOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#logfileGroupOption.
	VisitLogfileGroupOption(ctx *LogfileGroupOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createServer.
	VisitCreateServer(ctx *CreateServerContext) interface{}

	// Visit a parse tree produced by MySQLParser#serverOptions.
	VisitServerOptions(ctx *ServerOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#serverOption.
	VisitServerOption(ctx *ServerOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createTablespace.
	VisitCreateTablespace(ctx *CreateTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLParser#createUndoTablespace.
	VisitCreateUndoTablespace(ctx *CreateUndoTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsDataFileName.
	VisitTsDataFileName(ctx *TsDataFileNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsDataFile.
	VisitTsDataFile(ctx *TsDataFileContext) interface{}

	// Visit a parse tree produced by MySQLParser#tablespaceOptions.
	VisitTablespaceOptions(ctx *TablespaceOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#tablespaceOption.
	VisitTablespaceOption(ctx *TablespaceOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionInitialSize.
	VisitTsOptionInitialSize(ctx *TsOptionInitialSizeContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionUndoRedoBufferSize.
	VisitTsOptionUndoRedoBufferSize(ctx *TsOptionUndoRedoBufferSizeContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionAutoextendSize.
	VisitTsOptionAutoextendSize(ctx *TsOptionAutoextendSizeContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionMaxSize.
	VisitTsOptionMaxSize(ctx *TsOptionMaxSizeContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionExtentSize.
	VisitTsOptionExtentSize(ctx *TsOptionExtentSizeContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionNodegroup.
	VisitTsOptionNodegroup(ctx *TsOptionNodegroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionEngine.
	VisitTsOptionEngine(ctx *TsOptionEngineContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionWait.
	VisitTsOptionWait(ctx *TsOptionWaitContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionComment.
	VisitTsOptionComment(ctx *TsOptionCommentContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionFileblockSize.
	VisitTsOptionFileblockSize(ctx *TsOptionFileblockSizeContext) interface{}

	// Visit a parse tree produced by MySQLParser#tsOptionEncryption.
	VisitTsOptionEncryption(ctx *TsOptionEncryptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewReplaceOrAlgorithm.
	VisitViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewAlgorithm.
	VisitViewAlgorithm(ctx *ViewAlgorithmContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewSuid.
	VisitViewSuid(ctx *ViewSuidContext) interface{}

	// Visit a parse tree produced by MySQLParser#createTrigger.
	VisitCreateTrigger(ctx *CreateTriggerContext) interface{}

	// Visit a parse tree produced by MySQLParser#triggerFollowsPrecedesClause.
	VisitTriggerFollowsPrecedesClause(ctx *TriggerFollowsPrecedesClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#createEvent.
	VisitCreateEvent(ctx *CreateEventContext) interface{}

	// Visit a parse tree produced by MySQLParser#createRole.
	VisitCreateRole(ctx *CreateRoleContext) interface{}

	// Visit a parse tree produced by MySQLParser#createSpatialReference.
	VisitCreateSpatialReference(ctx *CreateSpatialReferenceContext) interface{}

	// Visit a parse tree produced by MySQLParser#srsAttribute.
	VisitSrsAttribute(ctx *SrsAttributeContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropStatement.
	VisitDropStatement(ctx *DropStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropEvent.
	VisitDropEvent(ctx *DropEventContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropFunction.
	VisitDropFunction(ctx *DropFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropProcedure.
	VisitDropProcedure(ctx *DropProcedureContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropIndex.
	VisitDropIndex(ctx *DropIndexContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropLogfileGroup.
	VisitDropLogfileGroup(ctx *DropLogfileGroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropLogfileGroupOption.
	VisitDropLogfileGroupOption(ctx *DropLogfileGroupOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropServer.
	VisitDropServer(ctx *DropServerContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropTableSpace.
	VisitDropTableSpace(ctx *DropTableSpaceContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropTrigger.
	VisitDropTrigger(ctx *DropTriggerContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropRole.
	VisitDropRole(ctx *DropRoleContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropSpatialReference.
	VisitDropSpatialReference(ctx *DropSpatialReferenceContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropUndoTablespace.
	VisitDropUndoTablespace(ctx *DropUndoTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLParser#renameTableStatement.
	VisitRenameTableStatement(ctx *RenameTableStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#renamePair.
	VisitRenamePair(ctx *RenamePairContext) interface{}

	// Visit a parse tree produced by MySQLParser#truncateTableStatement.
	VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#callStatement.
	VisitCallStatement(ctx *CallStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionDelete.
	VisitPartitionDelete(ctx *PartitionDeleteContext) interface{}

	// Visit a parse tree produced by MySQLParser#deleteStatementOption.
	VisitDeleteStatementOption(ctx *DeleteStatementOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#doStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#handlerStatement.
	VisitHandlerStatement(ctx *HandlerStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#handlerReadOrScan.
	VisitHandlerReadOrScan(ctx *HandlerReadOrScanContext) interface{}

	// Visit a parse tree produced by MySQLParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#insertLockOption.
	VisitInsertLockOption(ctx *InsertLockOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#insertFromConstructor.
	VisitInsertFromConstructor(ctx *InsertFromConstructorContext) interface{}

	// Visit a parse tree produced by MySQLParser#fields.
	VisitFields(ctx *FieldsContext) interface{}

	// Visit a parse tree produced by MySQLParser#insertValues.
	VisitInsertValues(ctx *InsertValuesContext) interface{}

	// Visit a parse tree produced by MySQLParser#insertQueryExpression.
	VisitInsertQueryExpression(ctx *InsertQueryExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by MySQLParser#values.
	VisitValues(ctx *ValuesContext) interface{}

	// Visit a parse tree produced by MySQLParser#valuesReference.
	VisitValuesReference(ctx *ValuesReferenceContext) interface{}

	// Visit a parse tree produced by MySQLParser#insertUpdateList.
	VisitInsertUpdateList(ctx *InsertUpdateListContext) interface{}

	// Visit a parse tree produced by MySQLParser#loadStatement.
	VisitLoadStatement(ctx *LoadStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#dataOrXml.
	VisitDataOrXml(ctx *DataOrXmlContext) interface{}

	// Visit a parse tree produced by MySQLParser#xmlRowsIdentifiedBy.
	VisitXmlRowsIdentifiedBy(ctx *XmlRowsIdentifiedByContext) interface{}

	// Visit a parse tree produced by MySQLParser#loadDataFileTail.
	VisitLoadDataFileTail(ctx *LoadDataFileTailContext) interface{}

	// Visit a parse tree produced by MySQLParser#loadDataFileTargetList.
	VisitLoadDataFileTargetList(ctx *LoadDataFileTargetListContext) interface{}

	// Visit a parse tree produced by MySQLParser#fieldOrVariableList.
	VisitFieldOrVariableList(ctx *FieldOrVariableListContext) interface{}

	// Visit a parse tree produced by MySQLParser#replaceStatement.
	VisitReplaceStatement(ctx *ReplaceStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#selectStatementWithInto.
	VisitSelectStatementWithInto(ctx *SelectStatementWithIntoContext) interface{}

	// Visit a parse tree produced by MySQLParser#queryExpression.
	VisitQueryExpression(ctx *QueryExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#queryExpressionBody.
	VisitQueryExpressionBody(ctx *QueryExpressionBodyContext) interface{}

	// Visit a parse tree produced by MySQLParser#queryExpressionParens.
	VisitQueryExpressionParens(ctx *QueryExpressionParensContext) interface{}

	// Visit a parse tree produced by MySQLParser#queryPrimary.
	VisitQueryPrimary(ctx *QueryPrimaryContext) interface{}

	// Visit a parse tree produced by MySQLParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by MySQLParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by MySQLParser#querySpecOption.
	VisitQuerySpecOption(ctx *QuerySpecOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleLimitClause.
	VisitSimpleLimitClause(ctx *SimpleLimitClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#limitOptions.
	VisitLimitOptions(ctx *LimitOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#limitOption.
	VisitLimitOption(ctx *LimitOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#intoClause.
	VisitIntoClause(ctx *IntoClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#procedureAnalyseClause.
	VisitProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowDefinition.
	VisitWindowDefinition(ctx *WindowDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowSpec.
	VisitWindowSpec(ctx *WindowSpecContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowSpecDetails.
	VisitWindowSpecDetails(ctx *WindowSpecDetailsContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFrameClause.
	VisitWindowFrameClause(ctx *WindowFrameClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFrameUnits.
	VisitWindowFrameUnits(ctx *WindowFrameUnitsContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFrameExtent.
	VisitWindowFrameExtent(ctx *WindowFrameExtentContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFrameStart.
	VisitWindowFrameStart(ctx *WindowFrameStartContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFrameBetween.
	VisitWindowFrameBetween(ctx *WindowFrameBetweenContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFrameBound.
	VisitWindowFrameBound(ctx *WindowFrameBoundContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFrameExclusion.
	VisitWindowFrameExclusion(ctx *WindowFrameExclusionContext) interface{}

	// Visit a parse tree produced by MySQLParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#commonTableExpression.
	VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#olapOption.
	VisitOlapOption(ctx *OlapOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#orderClause.
	VisitOrderClause(ctx *OrderClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#direction.
	VisitDirection(ctx *DirectionContext) interface{}

	// Visit a parse tree produced by MySQLParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableReferenceList.
	VisitTableReferenceList(ctx *TableReferenceListContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableValueConstructor.
	VisitTableValueConstructor(ctx *TableValueConstructorContext) interface{}

	// Visit a parse tree produced by MySQLParser#explicitTable.
	VisitExplicitTable(ctx *ExplicitTableContext) interface{}

	// Visit a parse tree produced by MySQLParser#rowValueExplicit.
	VisitRowValueExplicit(ctx *RowValueExplicitContext) interface{}

	// Visit a parse tree produced by MySQLParser#selectOption.
	VisitSelectOption(ctx *SelectOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#lockingClauseList.
	VisitLockingClauseList(ctx *LockingClauseListContext) interface{}

	// Visit a parse tree produced by MySQLParser#lockingClause.
	VisitLockingClause(ctx *LockingClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#lockStrengh.
	VisitLockStrengh(ctx *LockStrenghContext) interface{}

	// Visit a parse tree produced by MySQLParser#lockedRowAction.
	VisitLockedRowAction(ctx *LockedRowActionContext) interface{}

	// Visit a parse tree produced by MySQLParser#selectItemList.
	VisitSelectItemList(ctx *SelectItemListContext) interface{}

	// Visit a parse tree produced by MySQLParser#selectItem.
	VisitSelectItem(ctx *SelectItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#selectAlias.
	VisitSelectAlias(ctx *SelectAliasContext) interface{}

	// Visit a parse tree produced by MySQLParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableReference.
	VisitTableReference(ctx *TableReferenceContext) interface{}

	// Visit a parse tree produced by MySQLParser#escapedTableReference.
	VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{}

	// Visit a parse tree produced by MySQLParser#joinedTable.
	VisitJoinedTable(ctx *JoinedTableContext) interface{}

	// Visit a parse tree produced by MySQLParser#naturalJoinType.
	VisitNaturalJoinType(ctx *NaturalJoinTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#innerJoinType.
	VisitInnerJoinType(ctx *InnerJoinTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#outerJoinType.
	VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableFactor.
	VisitTableFactor(ctx *TableFactorContext) interface{}

	// Visit a parse tree produced by MySQLParser#singleTable.
	VisitSingleTable(ctx *SingleTableContext) interface{}

	// Visit a parse tree produced by MySQLParser#singleTableParens.
	VisitSingleTableParens(ctx *SingleTableParensContext) interface{}

	// Visit a parse tree produced by MySQLParser#derivedTable.
	VisitDerivedTable(ctx *DerivedTableContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableReferenceListParens.
	VisitTableReferenceListParens(ctx *TableReferenceListParensContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableFunction.
	VisitTableFunction(ctx *TableFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnsClause.
	VisitColumnsClause(ctx *ColumnsClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#jtColumn.
	VisitJtColumn(ctx *JtColumnContext) interface{}

	// Visit a parse tree produced by MySQLParser#onEmptyOrError.
	VisitOnEmptyOrError(ctx *OnEmptyOrErrorContext) interface{}

	// Visit a parse tree produced by MySQLParser#onEmpty.
	VisitOnEmpty(ctx *OnEmptyContext) interface{}

	// Visit a parse tree produced by MySQLParser#onError.
	VisitOnError(ctx *OnErrorContext) interface{}

	// Visit a parse tree produced by MySQLParser#jtOnResponse.
	VisitJtOnResponse(ctx *JtOnResponseContext) interface{}

	// Visit a parse tree produced by MySQLParser#setOprSymbol.
	VisitSetOprSymbol(ctx *SetOprSymbolContext) interface{}

	// Visit a parse tree produced by MySQLParser#setOprOption.
	VisitSetOprOption(ctx *SetOprOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableAlias.
	VisitTableAlias(ctx *TableAliasContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexHintList.
	VisitIndexHintList(ctx *IndexHintListContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexHint.
	VisitIndexHint(ctx *IndexHintContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexHintType.
	VisitIndexHintType(ctx *IndexHintTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyOrIndex.
	VisitKeyOrIndex(ctx *KeyOrIndexContext) interface{}

	// Visit a parse tree produced by MySQLParser#constraintKeyType.
	VisitConstraintKeyType(ctx *ConstraintKeyTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexHintClause.
	VisitIndexHintClause(ctx *IndexHintClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexList.
	VisitIndexList(ctx *IndexListContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexListElement.
	VisitIndexListElement(ctx *IndexListElementContext) interface{}

	// Visit a parse tree produced by MySQLParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#transactionOrLockingStatement.
	VisitTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#transactionStatement.
	VisitTransactionStatement(ctx *TransactionStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#beginWork.
	VisitBeginWork(ctx *BeginWorkContext) interface{}

	// Visit a parse tree produced by MySQLParser#transactionCharacteristic.
	VisitTransactionCharacteristic(ctx *TransactionCharacteristicContext) interface{}

	// Visit a parse tree produced by MySQLParser#savepointStatement.
	VisitSavepointStatement(ctx *SavepointStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#lockStatement.
	VisitLockStatement(ctx *LockStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#lockItem.
	VisitLockItem(ctx *LockItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#lockOption.
	VisitLockOption(ctx *LockOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#xaStatement.
	VisitXaStatement(ctx *XaStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#xaConvert.
	VisitXaConvert(ctx *XaConvertContext) interface{}

	// Visit a parse tree produced by MySQLParser#xid.
	VisitXid(ctx *XidContext) interface{}

	// Visit a parse tree produced by MySQLParser#replicationStatement.
	VisitReplicationStatement(ctx *ReplicationStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#resetOption.
	VisitResetOption(ctx *ResetOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#masterResetOptions.
	VisitMasterResetOptions(ctx *MasterResetOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#replicationLoad.
	VisitReplicationLoad(ctx *ReplicationLoadContext) interface{}

	// Visit a parse tree produced by MySQLParser#changeMaster.
	VisitChangeMaster(ctx *ChangeMasterContext) interface{}

	// Visit a parse tree produced by MySQLParser#changeMasterOptions.
	VisitChangeMasterOptions(ctx *ChangeMasterOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#masterOption.
	VisitMasterOption(ctx *MasterOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#privilegeCheckDef.
	VisitPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) interface{}

	// Visit a parse tree produced by MySQLParser#tablePrimaryKeyCheckDef.
	VisitTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) interface{}

	// Visit a parse tree produced by MySQLParser#masterTlsCiphersuitesDef.
	VisitMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) interface{}

	// Visit a parse tree produced by MySQLParser#masterFileDef.
	VisitMasterFileDef(ctx *MasterFileDefContext) interface{}

	// Visit a parse tree produced by MySQLParser#serverIdList.
	VisitServerIdList(ctx *ServerIdListContext) interface{}

	// Visit a parse tree produced by MySQLParser#changeReplication.
	VisitChangeReplication(ctx *ChangeReplicationContext) interface{}

	// Visit a parse tree produced by MySQLParser#filterDefinition.
	VisitFilterDefinition(ctx *FilterDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#filterDbList.
	VisitFilterDbList(ctx *FilterDbListContext) interface{}

	// Visit a parse tree produced by MySQLParser#filterTableList.
	VisitFilterTableList(ctx *FilterTableListContext) interface{}

	// Visit a parse tree produced by MySQLParser#filterStringList.
	VisitFilterStringList(ctx *FilterStringListContext) interface{}

	// Visit a parse tree produced by MySQLParser#filterWildDbTableString.
	VisitFilterWildDbTableString(ctx *FilterWildDbTableStringContext) interface{}

	// Visit a parse tree produced by MySQLParser#filterDbPairList.
	VisitFilterDbPairList(ctx *FilterDbPairListContext) interface{}

	// Visit a parse tree produced by MySQLParser#slave.
	VisitSlave(ctx *SlaveContext) interface{}

	// Visit a parse tree produced by MySQLParser#slaveUntilOptions.
	VisitSlaveUntilOptions(ctx *SlaveUntilOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#slaveConnectionOptions.
	VisitSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#slaveThreadOptions.
	VisitSlaveThreadOptions(ctx *SlaveThreadOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#slaveThreadOption.
	VisitSlaveThreadOption(ctx *SlaveThreadOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#groupReplication.
	VisitGroupReplication(ctx *GroupReplicationContext) interface{}

	// Visit a parse tree produced by MySQLParser#preparedStatement.
	VisitPreparedStatement(ctx *PreparedStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#executeVarList.
	VisitExecuteVarList(ctx *ExecuteVarListContext) interface{}

	// Visit a parse tree produced by MySQLParser#cloneStatement.
	VisitCloneStatement(ctx *CloneStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#dataDirSSL.
	VisitDataDirSSL(ctx *DataDirSSLContext) interface{}

	// Visit a parse tree produced by MySQLParser#ssl.
	VisitSsl(ctx *SslContext) interface{}

	// Visit a parse tree produced by MySQLParser#accountManagementStatement.
	VisitAccountManagementStatement(ctx *AccountManagementStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterUser.
	VisitAlterUser(ctx *AlterUserContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterUserTail.
	VisitAlterUserTail(ctx *AlterUserTailContext) interface{}

	// Visit a parse tree produced by MySQLParser#userFunction.
	VisitUserFunction(ctx *UserFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#createUser.
	VisitCreateUser(ctx *CreateUserContext) interface{}

	// Visit a parse tree produced by MySQLParser#createUserTail.
	VisitCreateUserTail(ctx *CreateUserTailContext) interface{}

	// Visit a parse tree produced by MySQLParser#defaultRoleClause.
	VisitDefaultRoleClause(ctx *DefaultRoleClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#requireClause.
	VisitRequireClause(ctx *RequireClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#connectOptions.
	VisitConnectOptions(ctx *ConnectOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#accountLockPasswordExpireOptions.
	VisitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropUser.
	VisitDropUser(ctx *DropUserContext) interface{}

	// Visit a parse tree produced by MySQLParser#grant.
	VisitGrant(ctx *GrantContext) interface{}

	// Visit a parse tree produced by MySQLParser#grantTargetList.
	VisitGrantTargetList(ctx *GrantTargetListContext) interface{}

	// Visit a parse tree produced by MySQLParser#grantOptions.
	VisitGrantOptions(ctx *GrantOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#exceptRoleList.
	VisitExceptRoleList(ctx *ExceptRoleListContext) interface{}

	// Visit a parse tree produced by MySQLParser#withRoles.
	VisitWithRoles(ctx *WithRolesContext) interface{}

	// Visit a parse tree produced by MySQLParser#grantAs.
	VisitGrantAs(ctx *GrantAsContext) interface{}

	// Visit a parse tree produced by MySQLParser#versionedRequireClause.
	VisitVersionedRequireClause(ctx *VersionedRequireClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#renameUser.
	VisitRenameUser(ctx *RenameUserContext) interface{}

	// Visit a parse tree produced by MySQLParser#revoke.
	VisitRevoke(ctx *RevokeContext) interface{}

	// Visit a parse tree produced by MySQLParser#onTypeTo.
	VisitOnTypeTo(ctx *OnTypeToContext) interface{}

	// Visit a parse tree produced by MySQLParser#aclType.
	VisitAclType(ctx *AclTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleOrPrivilegesList.
	VisitRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleOrPrivilege.
	VisitRoleOrPrivilege(ctx *RoleOrPrivilegeContext) interface{}

	// Visit a parse tree produced by MySQLParser#grantIdentifier.
	VisitGrantIdentifier(ctx *GrantIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#requireList.
	VisitRequireList(ctx *RequireListContext) interface{}

	// Visit a parse tree produced by MySQLParser#requireListElement.
	VisitRequireListElement(ctx *RequireListElementContext) interface{}

	// Visit a parse tree produced by MySQLParser#grantOption.
	VisitGrantOption(ctx *GrantOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#setRole.
	VisitSetRole(ctx *SetRoleContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleList.
	VisitRoleList(ctx *RoleListContext) interface{}

	// Visit a parse tree produced by MySQLParser#role.
	VisitRole(ctx *RoleContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableAdministrationStatement.
	VisitTableAdministrationStatement(ctx *TableAdministrationStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#histogram.
	VisitHistogram(ctx *HistogramContext) interface{}

	// Visit a parse tree produced by MySQLParser#checkOption.
	VisitCheckOption(ctx *CheckOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#repairType.
	VisitRepairType(ctx *RepairTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#installUninstallStatment.
	VisitInstallUninstallStatment(ctx *InstallUninstallStatmentContext) interface{}

	// Visit a parse tree produced by MySQLParser#setStatement.
	VisitSetStatement(ctx *SetStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#startOptionValueList.
	VisitStartOptionValueList(ctx *StartOptionValueListContext) interface{}

	// Visit a parse tree produced by MySQLParser#transactionCharacteristics.
	VisitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) interface{}

	// Visit a parse tree produced by MySQLParser#transactionAccessMode.
	VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{}

	// Visit a parse tree produced by MySQLParser#isolationLevel.
	VisitIsolationLevel(ctx *IsolationLevelContext) interface{}

	// Visit a parse tree produced by MySQLParser#optionValueListContinued.
	VisitOptionValueListContinued(ctx *OptionValueListContinuedContext) interface{}

	// Visit a parse tree produced by MySQLParser#optionValueNoOptionType.
	VisitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#optionValue.
	VisitOptionValue(ctx *OptionValueContext) interface{}

	// Visit a parse tree produced by MySQLParser#setSystemVariable.
	VisitSetSystemVariable(ctx *SetSystemVariableContext) interface{}

	// Visit a parse tree produced by MySQLParser#startOptionValueListFollowingOptionType.
	VisitStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#optionValueFollowingOptionType.
	VisitOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#setExprOrDefault.
	VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{}

	// Visit a parse tree produced by MySQLParser#showStatement.
	VisitShowStatement(ctx *ShowStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#showCommandType.
	VisitShowCommandType(ctx *ShowCommandTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#nonBlocking.
	VisitNonBlocking(ctx *NonBlockingContext) interface{}

	// Visit a parse tree produced by MySQLParser#fromOrIn.
	VisitFromOrIn(ctx *FromOrInContext) interface{}

	// Visit a parse tree produced by MySQLParser#inDb.
	VisitInDb(ctx *InDbContext) interface{}

	// Visit a parse tree produced by MySQLParser#profileType.
	VisitProfileType(ctx *ProfileTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#otherAdministrativeStatement.
	VisitOtherAdministrativeStatement(ctx *OtherAdministrativeStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyCacheListOrParts.
	VisitKeyCacheListOrParts(ctx *KeyCacheListOrPartsContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyCacheList.
	VisitKeyCacheList(ctx *KeyCacheListContext) interface{}

	// Visit a parse tree produced by MySQLParser#assignToKeycache.
	VisitAssignToKeycache(ctx *AssignToKeycacheContext) interface{}

	// Visit a parse tree produced by MySQLParser#assignToKeycachePartition.
	VisitAssignToKeycachePartition(ctx *AssignToKeycachePartitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#cacheKeyList.
	VisitCacheKeyList(ctx *CacheKeyListContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyUsageElement.
	VisitKeyUsageElement(ctx *KeyUsageElementContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyUsageList.
	VisitKeyUsageList(ctx *KeyUsageListContext) interface{}

	// Visit a parse tree produced by MySQLParser#flushOption.
	VisitFlushOption(ctx *FlushOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#logType.
	VisitLogType(ctx *LogTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#flushTables.
	VisitFlushTables(ctx *FlushTablesContext) interface{}

	// Visit a parse tree produced by MySQLParser#flushTablesOptions.
	VisitFlushTablesOptions(ctx *FlushTablesOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#preloadTail.
	VisitPreloadTail(ctx *PreloadTailContext) interface{}

	// Visit a parse tree produced by MySQLParser#preloadList.
	VisitPreloadList(ctx *PreloadListContext) interface{}

	// Visit a parse tree produced by MySQLParser#preloadKeys.
	VisitPreloadKeys(ctx *PreloadKeysContext) interface{}

	// Visit a parse tree produced by MySQLParser#adminPartition.
	VisitAdminPartition(ctx *AdminPartitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#resourceGroupManagement.
	VisitResourceGroupManagement(ctx *ResourceGroupManagementContext) interface{}

	// Visit a parse tree produced by MySQLParser#createResourceGroup.
	VisitCreateResourceGroup(ctx *CreateResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#resourceGroupVcpuList.
	VisitResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) interface{}

	// Visit a parse tree produced by MySQLParser#vcpuNumOrRange.
	VisitVcpuNumOrRange(ctx *VcpuNumOrRangeContext) interface{}

	// Visit a parse tree produced by MySQLParser#resourceGroupPriority.
	VisitResourceGroupPriority(ctx *ResourceGroupPriorityContext) interface{}

	// Visit a parse tree produced by MySQLParser#resourceGroupEnableDisable.
	VisitResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterResourceGroup.
	VisitAlterResourceGroup(ctx *AlterResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#setResourceGroup.
	VisitSetResourceGroup(ctx *SetResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#threadIdList.
	VisitThreadIdList(ctx *ThreadIdListContext) interface{}

	// Visit a parse tree produced by MySQLParser#dropResourceGroup.
	VisitDropResourceGroup(ctx *DropResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLParser#utilityStatement.
	VisitUtilityStatement(ctx *UtilityStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#describeStatement.
	VisitDescribeStatement(ctx *DescribeStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#explainStatement.
	VisitExplainStatement(ctx *ExplainStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#explainableStatement.
	VisitExplainableStatement(ctx *ExplainableStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#helpCommand.
	VisitHelpCommand(ctx *HelpCommandContext) interface{}

	// Visit a parse tree produced by MySQLParser#useCommand.
	VisitUseCommand(ctx *UseCommandContext) interface{}

	// Visit a parse tree produced by MySQLParser#restartServer.
	VisitRestartServer(ctx *RestartServerContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprOr.
	VisitExprOr(ctx *ExprOrContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprNot.
	VisitExprNot(ctx *ExprNotContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprIs.
	VisitExprIs(ctx *ExprIsContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprAnd.
	VisitExprAnd(ctx *ExprAndContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprXor.
	VisitExprXor(ctx *ExprXorContext) interface{}

	// Visit a parse tree produced by MySQLParser#primaryExprPredicate.
	VisitPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) interface{}

	// Visit a parse tree produced by MySQLParser#primaryExprCompare.
	VisitPrimaryExprCompare(ctx *PrimaryExprCompareContext) interface{}

	// Visit a parse tree produced by MySQLParser#primaryExprAllAny.
	VisitPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) interface{}

	// Visit a parse tree produced by MySQLParser#primaryExprIsNull.
	VisitPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) interface{}

	// Visit a parse tree produced by MySQLParser#compOp.
	VisitCompOp(ctx *CompOpContext) interface{}

	// Visit a parse tree produced by MySQLParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by MySQLParser#predicateExprIn.
	VisitPredicateExprIn(ctx *PredicateExprInContext) interface{}

	// Visit a parse tree produced by MySQLParser#predicateExprBetween.
	VisitPredicateExprBetween(ctx *PredicateExprBetweenContext) interface{}

	// Visit a parse tree produced by MySQLParser#predicateExprLike.
	VisitPredicateExprLike(ctx *PredicateExprLikeContext) interface{}

	// Visit a parse tree produced by MySQLParser#predicateExprRegex.
	VisitPredicateExprRegex(ctx *PredicateExprRegexContext) interface{}

	// Visit a parse tree produced by MySQLParser#bitExpr.
	VisitBitExpr(ctx *BitExprContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprConvert.
	VisitSimpleExprConvert(ctx *SimpleExprConvertContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprSearchJson.
	VisitSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprVariable.
	VisitSimpleExprVariable(ctx *SimpleExprVariableContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprCast.
	VisitSimpleExprCast(ctx *SimpleExprCastContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprUnary.
	VisitSimpleExprUnary(ctx *SimpleExprUnaryContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprOdbc.
	VisitSimpleExprOdbc(ctx *SimpleExprOdbcContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprRuntimeFunction.
	VisitSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprFunction.
	VisitSimpleExprFunction(ctx *SimpleExprFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprCollate.
	VisitSimpleExprCollate(ctx *SimpleExprCollateContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprMatch.
	VisitSimpleExprMatch(ctx *SimpleExprMatchContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprWindowingFunction.
	VisitSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprBinary.
	VisitSimpleExprBinary(ctx *SimpleExprBinaryContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprColumnRef.
	VisitSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprParamMarker.
	VisitSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprSum.
	VisitSimpleExprSum(ctx *SimpleExprSumContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprConvertUsing.
	VisitSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprSubQuery.
	VisitSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprGroupingOperation.
	VisitSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprNot.
	VisitSimpleExprNot(ctx *SimpleExprNotContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprValues.
	VisitSimpleExprValues(ctx *SimpleExprValuesContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprDefault.
	VisitSimpleExprDefault(ctx *SimpleExprDefaultContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprList.
	VisitSimpleExprList(ctx *SimpleExprListContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprInterval.
	VisitSimpleExprInterval(ctx *SimpleExprIntervalContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprCase.
	VisitSimpleExprCase(ctx *SimpleExprCaseContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprConcat.
	VisitSimpleExprConcat(ctx *SimpleExprConcatContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprLiteral.
	VisitSimpleExprLiteral(ctx *SimpleExprLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#arrayCast.
	VisitArrayCast(ctx *ArrayCastContext) interface{}

	// Visit a parse tree produced by MySQLParser#jsonOperator.
	VisitJsonOperator(ctx *JsonOperatorContext) interface{}

	// Visit a parse tree produced by MySQLParser#sumExpr.
	VisitSumExpr(ctx *SumExprContext) interface{}

	// Visit a parse tree produced by MySQLParser#groupingOperation.
	VisitGroupingOperation(ctx *GroupingOperationContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowFunctionCall.
	VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowingClause.
	VisitWindowingClause(ctx *WindowingClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#leadLagInfo.
	VisitLeadLagInfo(ctx *LeadLagInfoContext) interface{}

	// Visit a parse tree produced by MySQLParser#nullTreatment.
	VisitNullTreatment(ctx *NullTreatmentContext) interface{}

	// Visit a parse tree produced by MySQLParser#jsonFunction.
	VisitJsonFunction(ctx *JsonFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#inSumExpr.
	VisitInSumExpr(ctx *InSumExprContext) interface{}

	// Visit a parse tree produced by MySQLParser#identListArg.
	VisitIdentListArg(ctx *IdentListArgContext) interface{}

	// Visit a parse tree produced by MySQLParser#identList.
	VisitIdentList(ctx *IdentListContext) interface{}

	// Visit a parse tree produced by MySQLParser#fulltextOptions.
	VisitFulltextOptions(ctx *FulltextOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#runtimeFunctionCall.
	VisitRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) interface{}

	// Visit a parse tree produced by MySQLParser#geometryFunction.
	VisitGeometryFunction(ctx *GeometryFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#timeFunctionParameters.
	VisitTimeFunctionParameters(ctx *TimeFunctionParametersContext) interface{}

	// Visit a parse tree produced by MySQLParser#fractionalPrecision.
	VisitFractionalPrecision(ctx *FractionalPrecisionContext) interface{}

	// Visit a parse tree produced by MySQLParser#weightStringLevels.
	VisitWeightStringLevels(ctx *WeightStringLevelsContext) interface{}

	// Visit a parse tree produced by MySQLParser#weightStringLevelListItem.
	VisitWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#dateTimeTtype.
	VisitDateTimeTtype(ctx *DateTimeTtypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#trimFunction.
	VisitTrimFunction(ctx *TrimFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#substringFunction.
	VisitSubstringFunction(ctx *SubstringFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by MySQLParser#searchJsonFunction.
	VisitSearchJsonFunction(ctx *SearchJsonFunctionContext) interface{}

	// Visit a parse tree produced by MySQLParser#jsonValueReturning.
	VisitJsonValueReturning(ctx *JsonValueReturningContext) interface{}

	// Visit a parse tree produced by MySQLParser#jsonValueOnEmpty.
	VisitJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) interface{}

	// Visit a parse tree produced by MySQLParser#jsonValueOnError.
	VisitJsonValueOnError(ctx *JsonValueOnErrorContext) interface{}

	// Visit a parse tree produced by MySQLParser#udfExprList.
	VisitUdfExprList(ctx *UdfExprListContext) interface{}

	// Visit a parse tree produced by MySQLParser#udfExpr.
	VisitUdfExpr(ctx *UdfExprContext) interface{}

	// Visit a parse tree produced by MySQLParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by MySQLParser#userVariable.
	VisitUserVariable(ctx *UserVariableContext) interface{}

	// Visit a parse tree produced by MySQLParser#systemVariable.
	VisitSystemVariable(ctx *SystemVariableContext) interface{}

	// Visit a parse tree produced by MySQLParser#internalVariableName.
	VisitInternalVariableName(ctx *InternalVariableNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#whenExpression.
	VisitWhenExpression(ctx *WhenExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#thenExpression.
	VisitThenExpression(ctx *ThenExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#elseExpression.
	VisitElseExpression(ctx *ElseExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#castType.
	VisitCastType(ctx *CastTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by MySQLParser#charset.
	VisitCharset(ctx *CharsetContext) interface{}

	// Visit a parse tree produced by MySQLParser#notRule.
	VisitNotRule(ctx *NotRuleContext) interface{}

	// Visit a parse tree produced by MySQLParser#not2Rule.
	VisitNot2Rule(ctx *Not2RuleContext) interface{}

	// Visit a parse tree produced by MySQLParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by MySQLParser#intervalTimeStamp.
	VisitIntervalTimeStamp(ctx *IntervalTimeStampContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprListWithParentheses.
	VisitExprListWithParentheses(ctx *ExprListWithParenthesesContext) interface{}

	// Visit a parse tree produced by MySQLParser#exprWithParentheses.
	VisitExprWithParentheses(ctx *ExprWithParenthesesContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleExprWithParentheses.
	VisitSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) interface{}

	// Visit a parse tree produced by MySQLParser#orderList.
	VisitOrderList(ctx *OrderListContext) interface{}

	// Visit a parse tree produced by MySQLParser#orderExpression.
	VisitOrderExpression(ctx *OrderExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#groupList.
	VisitGroupList(ctx *GroupListContext) interface{}

	// Visit a parse tree produced by MySQLParser#groupingExpression.
	VisitGroupingExpression(ctx *GroupingExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#channel.
	VisitChannel(ctx *ChannelContext) interface{}

	// Visit a parse tree produced by MySQLParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#ifBody.
	VisitIfBody(ctx *IfBodyContext) interface{}

	// Visit a parse tree produced by MySQLParser#thenStatement.
	VisitThenStatement(ctx *ThenStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#compoundStatementList.
	VisitCompoundStatementList(ctx *CompoundStatementListContext) interface{}

	// Visit a parse tree produced by MySQLParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#labeledBlock.
	VisitLabeledBlock(ctx *LabeledBlockContext) interface{}

	// Visit a parse tree produced by MySQLParser#unlabeledBlock.
	VisitUnlabeledBlock(ctx *UnlabeledBlockContext) interface{}

	// Visit a parse tree produced by MySQLParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by MySQLParser#beginEndBlock.
	VisitBeginEndBlock(ctx *BeginEndBlockContext) interface{}

	// Visit a parse tree produced by MySQLParser#labeledControl.
	VisitLabeledControl(ctx *LabeledControlContext) interface{}

	// Visit a parse tree produced by MySQLParser#unlabeledControl.
	VisitUnlabeledControl(ctx *UnlabeledControlContext) interface{}

	// Visit a parse tree produced by MySQLParser#loopBlock.
	VisitLoopBlock(ctx *LoopBlockContext) interface{}

	// Visit a parse tree produced by MySQLParser#whileDoBlock.
	VisitWhileDoBlock(ctx *WhileDoBlockContext) interface{}

	// Visit a parse tree produced by MySQLParser#repeatUntilBlock.
	VisitRepeatUntilBlock(ctx *RepeatUntilBlockContext) interface{}

	// Visit a parse tree produced by MySQLParser#spDeclarations.
	VisitSpDeclarations(ctx *SpDeclarationsContext) interface{}

	// Visit a parse tree produced by MySQLParser#spDeclaration.
	VisitSpDeclaration(ctx *SpDeclarationContext) interface{}

	// Visit a parse tree produced by MySQLParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by MySQLParser#conditionDeclaration.
	VisitConditionDeclaration(ctx *ConditionDeclarationContext) interface{}

	// Visit a parse tree produced by MySQLParser#spCondition.
	VisitSpCondition(ctx *SpConditionContext) interface{}

	// Visit a parse tree produced by MySQLParser#sqlstate.
	VisitSqlstate(ctx *SqlstateContext) interface{}

	// Visit a parse tree produced by MySQLParser#handlerDeclaration.
	VisitHandlerDeclaration(ctx *HandlerDeclarationContext) interface{}

	// Visit a parse tree produced by MySQLParser#handlerCondition.
	VisitHandlerCondition(ctx *HandlerConditionContext) interface{}

	// Visit a parse tree produced by MySQLParser#cursorDeclaration.
	VisitCursorDeclaration(ctx *CursorDeclarationContext) interface{}

	// Visit a parse tree produced by MySQLParser#iterateStatement.
	VisitIterateStatement(ctx *IterateStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#leaveStatement.
	VisitLeaveStatement(ctx *LeaveStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#getDiagnostics.
	VisitGetDiagnostics(ctx *GetDiagnosticsContext) interface{}

	// Visit a parse tree produced by MySQLParser#signalAllowedExpr.
	VisitSignalAllowedExpr(ctx *SignalAllowedExprContext) interface{}

	// Visit a parse tree produced by MySQLParser#statementInformationItem.
	VisitStatementInformationItem(ctx *StatementInformationItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#conditionInformationItem.
	VisitConditionInformationItem(ctx *ConditionInformationItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#signalInformationItemName.
	VisitSignalInformationItemName(ctx *SignalInformationItemNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#signalStatement.
	VisitSignalStatement(ctx *SignalStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#resignalStatement.
	VisitResignalStatement(ctx *ResignalStatementContext) interface{}

	// Visit a parse tree produced by MySQLParser#signalInformationItem.
	VisitSignalInformationItem(ctx *SignalInformationItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#cursorOpen.
	VisitCursorOpen(ctx *CursorOpenContext) interface{}

	// Visit a parse tree produced by MySQLParser#cursorClose.
	VisitCursorClose(ctx *CursorCloseContext) interface{}

	// Visit a parse tree produced by MySQLParser#cursorFetch.
	VisitCursorFetch(ctx *CursorFetchContext) interface{}

	// Visit a parse tree produced by MySQLParser#schedule.
	VisitSchedule(ctx *ScheduleContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#checkOrReferences.
	VisitCheckOrReferences(ctx *CheckOrReferencesContext) interface{}

	// Visit a parse tree produced by MySQLParser#checkConstraint.
	VisitCheckConstraint(ctx *CheckConstraintContext) interface{}

	// Visit a parse tree produced by MySQLParser#constraintEnforcement.
	VisitConstraintEnforcement(ctx *ConstraintEnforcementContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableConstraintDef.
	VisitTableConstraintDef(ctx *TableConstraintDefContext) interface{}

	// Visit a parse tree produced by MySQLParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#fieldDefinition.
	VisitFieldDefinition(ctx *FieldDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnAttribute.
	VisitColumnAttribute(ctx *ColumnAttributeContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnFormat.
	VisitColumnFormat(ctx *ColumnFormatContext) interface{}

	// Visit a parse tree produced by MySQLParser#storageMedia.
	VisitStorageMedia(ctx *StorageMediaContext) interface{}

	// Visit a parse tree produced by MySQLParser#gcolAttribute.
	VisitGcolAttribute(ctx *GcolAttributeContext) interface{}

	// Visit a parse tree produced by MySQLParser#references.
	VisitReferences(ctx *ReferencesContext) interface{}

	// Visit a parse tree produced by MySQLParser#deleteOption.
	VisitDeleteOption(ctx *DeleteOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyList.
	VisitKeyList(ctx *KeyListContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyPart.
	VisitKeyPart(ctx *KeyPartContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyListWithExpression.
	VisitKeyListWithExpression(ctx *KeyListWithExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyPartOrExpression.
	VisitKeyPartOrExpression(ctx *KeyPartOrExpressionContext) interface{}

	// Visit a parse tree produced by MySQLParser#keyListVariants.
	VisitKeyListVariants(ctx *KeyListVariantsContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexOption.
	VisitIndexOption(ctx *IndexOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#commonIndexOption.
	VisitCommonIndexOption(ctx *CommonIndexOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#visibility.
	VisitVisibility(ctx *VisibilityContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexTypeClause.
	VisitIndexTypeClause(ctx *IndexTypeClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#fulltextIndexOption.
	VisitFulltextIndexOption(ctx *FulltextIndexOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#spatialIndexOption.
	VisitSpatialIndexOption(ctx *SpatialIndexOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#dataTypeDefinition.
	VisitDataTypeDefinition(ctx *DataTypeDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#nchar.
	VisitNchar(ctx *NcharContext) interface{}

	// Visit a parse tree produced by MySQLParser#realType.
	VisitRealType(ctx *RealTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#fieldLength.
	VisitFieldLength(ctx *FieldLengthContext) interface{}

	// Visit a parse tree produced by MySQLParser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#charsetWithOptBinary.
	VisitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) interface{}

	// Visit a parse tree produced by MySQLParser#ascii.
	VisitAscii(ctx *AsciiContext) interface{}

	// Visit a parse tree produced by MySQLParser#unicode.
	VisitUnicode(ctx *UnicodeContext) interface{}

	// Visit a parse tree produced by MySQLParser#wsNumCodepoints.
	VisitWsNumCodepoints(ctx *WsNumCodepointsContext) interface{}

	// Visit a parse tree produced by MySQLParser#typeDatetimePrecision.
	VisitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) interface{}

	// Visit a parse tree produced by MySQLParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#collationName.
	VisitCollationName(ctx *CollationNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#createTableOptions.
	VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#createTableOptionsSpaceSeparated.
	VisitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) interface{}

	// Visit a parse tree produced by MySQLParser#createTableOption.
	VisitCreateTableOption(ctx *CreateTableOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#ternaryOption.
	VisitTernaryOption(ctx *TernaryOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#defaultCollation.
	VisitDefaultCollation(ctx *DefaultCollationContext) interface{}

	// Visit a parse tree produced by MySQLParser#defaultEncryption.
	VisitDefaultEncryption(ctx *DefaultEncryptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#defaultCharset.
	VisitDefaultCharset(ctx *DefaultCharsetContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionClause.
	VisitPartitionClause(ctx *PartitionClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionDefKey.
	VisitPartitionDefKey(ctx *PartitionDefKeyContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionDefHash.
	VisitPartitionDefHash(ctx *PartitionDefHashContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionDefRangeList.
	VisitPartitionDefRangeList(ctx *PartitionDefRangeListContext) interface{}

	// Visit a parse tree produced by MySQLParser#subPartitions.
	VisitSubPartitions(ctx *SubPartitionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionKeyAlgorithm.
	VisitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionDefinitions.
	VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionDefinition.
	VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionValuesIn.
	VisitPartitionValuesIn(ctx *PartitionValuesInContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionOption.
	VisitPartitionOption(ctx *PartitionOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#subpartitionDefinition.
	VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionValueItemListParen.
	VisitPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) interface{}

	// Visit a parse tree produced by MySQLParser#partitionValueItem.
	VisitPartitionValueItem(ctx *PartitionValueItemContext) interface{}

	// Visit a parse tree produced by MySQLParser#definerClause.
	VisitDefinerClause(ctx *DefinerClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#ifExists.
	VisitIfExists(ctx *IfExistsContext) interface{}

	// Visit a parse tree produced by MySQLParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by MySQLParser#procedureParameter.
	VisitProcedureParameter(ctx *ProcedureParameterContext) interface{}

	// Visit a parse tree produced by MySQLParser#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) interface{}

	// Visit a parse tree produced by MySQLParser#collate.
	VisitCollate(ctx *CollateContext) interface{}

	// Visit a parse tree produced by MySQLParser#typeWithOptCollate.
	VisitTypeWithOptCollate(ctx *TypeWithOptCollateContext) interface{}

	// Visit a parse tree produced by MySQLParser#schemaIdentifierPair.
	VisitSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewRefList.
	VisitViewRefList(ctx *ViewRefListContext) interface{}

	// Visit a parse tree produced by MySQLParser#updateList.
	VisitUpdateList(ctx *UpdateListContext) interface{}

	// Visit a parse tree produced by MySQLParser#updateElement.
	VisitUpdateElement(ctx *UpdateElementContext) interface{}

	// Visit a parse tree produced by MySQLParser#charsetClause.
	VisitCharsetClause(ctx *CharsetClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#fieldsClause.
	VisitFieldsClause(ctx *FieldsClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#fieldTerm.
	VisitFieldTerm(ctx *FieldTermContext) interface{}

	// Visit a parse tree produced by MySQLParser#linesClause.
	VisitLinesClause(ctx *LinesClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#lineTerm.
	VisitLineTerm(ctx *LineTermContext) interface{}

	// Visit a parse tree produced by MySQLParser#userList.
	VisitUserList(ctx *UserListContext) interface{}

	// Visit a parse tree produced by MySQLParser#createUserList.
	VisitCreateUserList(ctx *CreateUserListContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterUserList.
	VisitAlterUserList(ctx *AlterUserListContext) interface{}

	// Visit a parse tree produced by MySQLParser#createUserEntry.
	VisitCreateUserEntry(ctx *CreateUserEntryContext) interface{}

	// Visit a parse tree produced by MySQLParser#alterUserEntry.
	VisitAlterUserEntry(ctx *AlterUserEntryContext) interface{}

	// Visit a parse tree produced by MySQLParser#retainCurrentPassword.
	VisitRetainCurrentPassword(ctx *RetainCurrentPasswordContext) interface{}

	// Visit a parse tree produced by MySQLParser#discardOldPassword.
	VisitDiscardOldPassword(ctx *DiscardOldPasswordContext) interface{}

	// Visit a parse tree produced by MySQLParser#replacePassword.
	VisitReplacePassword(ctx *ReplacePasswordContext) interface{}

	// Visit a parse tree produced by MySQLParser#userIdentifierOrText.
	VisitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) interface{}

	// Visit a parse tree produced by MySQLParser#user.
	VisitUser(ctx *UserContext) interface{}

	// Visit a parse tree produced by MySQLParser#likeClause.
	VisitLikeClause(ctx *LikeClauseContext) interface{}

	// Visit a parse tree produced by MySQLParser#likeOrWhere.
	VisitLikeOrWhere(ctx *LikeOrWhereContext) interface{}

	// Visit a parse tree produced by MySQLParser#onlineOption.
	VisitOnlineOption(ctx *OnlineOptionContext) interface{}

	// Visit a parse tree produced by MySQLParser#noWriteToBinLog.
	VisitNoWriteToBinLog(ctx *NoWriteToBinLogContext) interface{}

	// Visit a parse tree produced by MySQLParser#usePartition.
	VisitUsePartition(ctx *UsePartitionContext) interface{}

	// Visit a parse tree produced by MySQLParser#fieldIdentifier.
	VisitFieldIdentifier(ctx *FieldIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnInternalRef.
	VisitColumnInternalRef(ctx *ColumnInternalRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnInternalRefList.
	VisitColumnInternalRefList(ctx *ColumnInternalRefListContext) interface{}

	// Visit a parse tree produced by MySQLParser#columnRef.
	VisitColumnRef(ctx *ColumnRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#insertIdentifier.
	VisitInsertIdentifier(ctx *InsertIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#indexRef.
	VisitIndexRef(ctx *IndexRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableWild.
	VisitTableWild(ctx *TableWildContext) interface{}

	// Visit a parse tree produced by MySQLParser#schemaName.
	VisitSchemaName(ctx *SchemaNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#schemaRef.
	VisitSchemaRef(ctx *SchemaRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#procedureName.
	VisitProcedureName(ctx *ProcedureNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#procedureRef.
	VisitProcedureRef(ctx *ProcedureRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#functionRef.
	VisitFunctionRef(ctx *FunctionRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#triggerName.
	VisitTriggerName(ctx *TriggerNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#triggerRef.
	VisitTriggerRef(ctx *TriggerRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#viewRef.
	VisitViewRef(ctx *ViewRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#tablespaceName.
	VisitTablespaceName(ctx *TablespaceNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#tablespaceRef.
	VisitTablespaceRef(ctx *TablespaceRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#logfileGroupName.
	VisitLogfileGroupName(ctx *LogfileGroupNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#logfileGroupRef.
	VisitLogfileGroupRef(ctx *LogfileGroupRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#eventName.
	VisitEventName(ctx *EventNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#eventRef.
	VisitEventRef(ctx *EventRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#udfName.
	VisitUdfName(ctx *UdfNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#serverName.
	VisitServerName(ctx *ServerNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#serverRef.
	VisitServerRef(ctx *ServerRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#engineRef.
	VisitEngineRef(ctx *EngineRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#filterTableRef.
	VisitFilterTableRef(ctx *FilterTableRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableRefWithWildcard.
	VisitTableRefWithWildcard(ctx *TableRefWithWildcardContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableRef.
	VisitTableRef(ctx *TableRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableRefList.
	VisitTableRefList(ctx *TableRefListContext) interface{}

	// Visit a parse tree produced by MySQLParser#tableAliasRefList.
	VisitTableAliasRefList(ctx *TableAliasRefListContext) interface{}

	// Visit a parse tree produced by MySQLParser#parameterName.
	VisitParameterName(ctx *ParameterNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#labelIdentifier.
	VisitLabelIdentifier(ctx *LabelIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#labelRef.
	VisitLabelRef(ctx *LabelRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleIdentifier.
	VisitRoleIdentifier(ctx *RoleIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleRef.
	VisitRoleRef(ctx *RoleRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#pluginRef.
	VisitPluginRef(ctx *PluginRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#componentRef.
	VisitComponentRef(ctx *ComponentRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#resourceGroupRef.
	VisitResourceGroupRef(ctx *ResourceGroupRefContext) interface{}

	// Visit a parse tree produced by MySQLParser#windowName.
	VisitWindowName(ctx *WindowNameContext) interface{}

	// Visit a parse tree produced by MySQLParser#pureIdentifier.
	VisitPureIdentifier(ctx *PureIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierListWithParentheses.
	VisitIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) interface{}

	// Visit a parse tree produced by MySQLParser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#simpleIdentifier.
	VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#dotIdentifier.
	VisitDotIdentifier(ctx *DotIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#ulong_number.
	VisitUlong_number(ctx *Ulong_numberContext) interface{}

	// Visit a parse tree produced by MySQLParser#real_ulong_number.
	VisitReal_ulong_number(ctx *Real_ulong_numberContext) interface{}

	// Visit a parse tree produced by MySQLParser#ulonglong_number.
	VisitUlonglong_number(ctx *Ulonglong_numberContext) interface{}

	// Visit a parse tree produced by MySQLParser#real_ulonglong_number.
	VisitReal_ulonglong_number(ctx *Real_ulonglong_numberContext) interface{}

	// Visit a parse tree produced by MySQLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#signedLiteral.
	VisitSignedLiteral(ctx *SignedLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by MySQLParser#textStringLiteral.
	VisitTextStringLiteral(ctx *TextStringLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#textString.
	VisitTextString(ctx *TextStringContext) interface{}

	// Visit a parse tree produced by MySQLParser#textStringHash.
	VisitTextStringHash(ctx *TextStringHashContext) interface{}

	// Visit a parse tree produced by MySQLParser#textLiteral.
	VisitTextLiteral(ctx *TextLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#textStringNoLinebreak.
	VisitTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) interface{}

	// Visit a parse tree produced by MySQLParser#textStringLiteralList.
	VisitTextStringLiteralList(ctx *TextStringLiteralListContext) interface{}

	// Visit a parse tree produced by MySQLParser#numLiteral.
	VisitNumLiteral(ctx *NumLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#boolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#temporalLiteral.
	VisitTemporalLiteral(ctx *TemporalLiteralContext) interface{}

	// Visit a parse tree produced by MySQLParser#floatOptions.
	VisitFloatOptions(ctx *FloatOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#standardFloatOptions.
	VisitStandardFloatOptions(ctx *StandardFloatOptionsContext) interface{}

	// Visit a parse tree produced by MySQLParser#precision.
	VisitPrecision(ctx *PrecisionContext) interface{}

	// Visit a parse tree produced by MySQLParser#textOrIdentifier.
	VisitTextOrIdentifier(ctx *TextOrIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#lValueIdentifier.
	VisitLValueIdentifier(ctx *LValueIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleIdentifierOrText.
	VisitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) interface{}

	// Visit a parse tree produced by MySQLParser#sizeNumber.
	VisitSizeNumber(ctx *SizeNumberContext) interface{}

	// Visit a parse tree produced by MySQLParser#parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by MySQLParser#equal.
	VisitEqual(ctx *EqualContext) interface{}

	// Visit a parse tree produced by MySQLParser#optionType.
	VisitOptionType(ctx *OptionTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#varIdentType.
	VisitVarIdentType(ctx *VarIdentTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#setVarIdentType.
	VisitSetVarIdentType(ctx *SetVarIdentTypeContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierKeyword.
	VisitIdentifierKeyword(ctx *IdentifierKeywordContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierKeywordsAmbiguous1RolesAndLabels.
	VisitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierKeywordsAmbiguous2Labels.
	VisitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) interface{}

	// Visit a parse tree produced by MySQLParser#labelKeyword.
	VisitLabelKeyword(ctx *LabelKeywordContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierKeywordsAmbiguous3Roles.
	VisitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierKeywordsUnambiguous.
	VisitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleKeyword.
	VisitRoleKeyword(ctx *RoleKeywordContext) interface{}

	// Visit a parse tree produced by MySQLParser#lValueKeyword.
	VisitLValueKeyword(ctx *LValueKeywordContext) interface{}

	// Visit a parse tree produced by MySQLParser#identifierKeywordsAmbiguous4SystemVariables.
	VisitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleOrIdentifierKeyword.
	VisitRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) interface{}

	// Visit a parse tree produced by MySQLParser#roleOrLabelKeyword.
	VisitRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) interface{}
}
