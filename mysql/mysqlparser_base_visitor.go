// Code generated from MySQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mysql // MySQLParser
import "github.com/antlr4-go/antlr/v4"

type BaseMySQLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMySQLParserVisitor) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleStatement(ctx *SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterStatement(ctx *AlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterDatabase(ctx *AlterDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterDatabaseOption(ctx *AlterDatabaseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterEvent(ctx *AlterEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterLogfileGroupOptions(ctx *AlterLogfileGroupOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterLogfileGroupOption(ctx *AlterLogfileGroupOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterServer(ctx *AlterServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterTable(ctx *AlterTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterTableActions(ctx *AlterTableActionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterCommandList(ctx *AlterCommandListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterPartition(ctx *AlterPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterList(ctx *AlterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterCommandsModifier(ctx *AlterCommandsModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterListItem(ctx *AlterListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPlace(ctx *PlaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRestrict(ctx *RestrictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterOrderList(ctx *AlterOrderListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterLockOption(ctx *AlterLockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexLockAndAlgorithm(ctx *IndexLockAndAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWithValidation(ctx *WithValidationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRemovePartitioning(ctx *RemovePartitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterTablespace(ctx *AlterTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterUndoTablespace(ctx *AlterUndoTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUndoTableSpaceOptions(ctx *UndoTableSpaceOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUndoTableSpaceOption(ctx *UndoTableSpaceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterTablespaceOptions(ctx *AlterTablespaceOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterTablespaceOption(ctx *AlterTablespaceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitChangeTablespaceOption(ctx *ChangeTablespaceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterView(ctx *AlterViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewTail(ctx *ViewTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewSelect(ctx *ViewSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewCheckOption(ctx *ViewCheckOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterProcedure(ctx *AlterProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterFunction(ctx *AlterFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateStatement(ctx *CreateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateDatabaseOption(ctx *CreateDatabaseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableElementList(ctx *TableElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableElement(ctx *TableElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateRoutine(ctx *CreateRoutineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateProcedure(ctx *CreateProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateFunction(ctx *CreateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateUdf(ctx *CreateUdfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoutineCreateOption(ctx *RoutineCreateOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoutineAlterOptions(ctx *RoutineAlterOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoutineOption(ctx *RoutineOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateIndex(ctx *CreateIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexNameAndType(ctx *IndexNameAndTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateIndexTarget(ctx *CreateIndexTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLogfileGroupOptions(ctx *LogfileGroupOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLogfileGroupOption(ctx *LogfileGroupOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateServer(ctx *CreateServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitServerOptions(ctx *ServerOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitServerOption(ctx *ServerOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateTablespace(ctx *CreateTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateUndoTablespace(ctx *CreateUndoTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsDataFileName(ctx *TsDataFileNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsDataFile(ctx *TsDataFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTablespaceOptions(ctx *TablespaceOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTablespaceOption(ctx *TablespaceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionInitialSize(ctx *TsOptionInitialSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionUndoRedoBufferSize(ctx *TsOptionUndoRedoBufferSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionAutoextendSize(ctx *TsOptionAutoextendSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionMaxSize(ctx *TsOptionMaxSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionExtentSize(ctx *TsOptionExtentSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionNodegroup(ctx *TsOptionNodegroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionEngine(ctx *TsOptionEngineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionWait(ctx *TsOptionWaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionComment(ctx *TsOptionCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionFileblockSize(ctx *TsOptionFileblockSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTsOptionEncryption(ctx *TsOptionEncryptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewAlgorithm(ctx *ViewAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewSuid(ctx *ViewSuidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateTrigger(ctx *CreateTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTriggerFollowsPrecedesClause(ctx *TriggerFollowsPrecedesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateEvent(ctx *CreateEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateRole(ctx *CreateRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateSpatialReference(ctx *CreateSpatialReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSrsAttribute(ctx *SrsAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropStatement(ctx *DropStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropDatabase(ctx *DropDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropEvent(ctx *DropEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropFunction(ctx *DropFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropProcedure(ctx *DropProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropIndex(ctx *DropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropLogfileGroup(ctx *DropLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropLogfileGroupOption(ctx *DropLogfileGroupOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropServer(ctx *DropServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropTableSpace(ctx *DropTableSpaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropTrigger(ctx *DropTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropRole(ctx *DropRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropSpatialReference(ctx *DropSpatialReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropUndoTablespace(ctx *DropUndoTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRenameTableStatement(ctx *RenameTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRenamePair(ctx *RenamePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCallStatement(ctx *CallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionDelete(ctx *PartitionDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDeleteStatementOption(ctx *DeleteStatementOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitHandlerStatement(ctx *HandlerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitHandlerReadOrScan(ctx *HandlerReadOrScanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInsertLockOption(ctx *InsertLockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInsertFromConstructor(ctx *InsertFromConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFields(ctx *FieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInsertValues(ctx *InsertValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInsertQueryExpression(ctx *InsertQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitValueList(ctx *ValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitValues(ctx *ValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitValuesReference(ctx *ValuesReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInsertUpdateList(ctx *InsertUpdateListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDataOrXml(ctx *DataOrXmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitXmlRowsIdentifiedBy(ctx *XmlRowsIdentifiedByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLoadDataFileTail(ctx *LoadDataFileTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLoadDataFileTargetList(ctx *LoadDataFileTargetListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFieldOrVariableList(ctx *FieldOrVariableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReplaceStatement(ctx *ReplaceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSelectStatementWithInto(ctx *SelectStatementWithIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQueryExpressionBody(ctx *QueryExpressionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQueryExpressionParens(ctx *QueryExpressionParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQueryPrimary(ctx *QueryPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQuerySpecOption(ctx *QuerySpecOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleLimitClause(ctx *SimpleLimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLimitOptions(ctx *LimitOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLimitOption(ctx *LimitOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIntoClause(ctx *IntoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowDefinition(ctx *WindowDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowSpec(ctx *WindowSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowSpecDetails(ctx *WindowSpecDetailsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFrameClause(ctx *WindowFrameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFrameUnits(ctx *WindowFrameUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFrameExtent(ctx *WindowFrameExtentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFrameStart(ctx *WindowFrameStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFrameBetween(ctx *WindowFrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFrameBound(ctx *WindowFrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFrameExclusion(ctx *WindowFrameExclusionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOlapOption(ctx *OlapOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOrderClause(ctx *OrderClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDirection(ctx *DirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableReferenceList(ctx *TableReferenceListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableValueConstructor(ctx *TableValueConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExplicitTable(ctx *ExplicitTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRowValueExplicit(ctx *RowValueExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSelectOption(ctx *SelectOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLockingClauseList(ctx *LockingClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLockingClause(ctx *LockingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLockStrengh(ctx *LockStrenghContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLockedRowAction(ctx *LockedRowActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSelectItemList(ctx *SelectItemListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSelectItem(ctx *SelectItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSelectAlias(ctx *SelectAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableReference(ctx *TableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJoinedTable(ctx *JoinedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNaturalJoinType(ctx *NaturalJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInnerJoinType(ctx *InnerJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableFactor(ctx *TableFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSingleTable(ctx *SingleTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSingleTableParens(ctx *SingleTableParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDerivedTable(ctx *DerivedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableReferenceListParens(ctx *TableReferenceListParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableFunction(ctx *TableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnsClause(ctx *ColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJtColumn(ctx *JtColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOnEmptyOrError(ctx *OnEmptyOrErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOnEmpty(ctx *OnEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOnError(ctx *OnErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJtOnResponse(ctx *JtOnResponseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetOprSymbol(ctx *SetOprSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetOprOption(ctx *SetOprOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableAlias(ctx *TableAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexHintList(ctx *IndexHintListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyOrIndex(ctx *KeyOrIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitConstraintKeyType(ctx *ConstraintKeyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexHintClause(ctx *IndexHintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexList(ctx *IndexListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexListElement(ctx *IndexListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTransactionStatement(ctx *TransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitBeginWork(ctx *BeginWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTransactionCharacteristic(ctx *TransactionCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSavepointStatement(ctx *SavepointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLockStatement(ctx *LockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLockItem(ctx *LockItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLockOption(ctx *LockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitXaStatement(ctx *XaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitXaConvert(ctx *XaConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitXid(ctx *XidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReplicationStatement(ctx *ReplicationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitResetOption(ctx *ResetOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitMasterResetOptions(ctx *MasterResetOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReplicationLoad(ctx *ReplicationLoadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitChangeMaster(ctx *ChangeMasterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitChangeMasterOptions(ctx *ChangeMasterOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitMasterOption(ctx *MasterOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitMasterFileDef(ctx *MasterFileDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitServerIdList(ctx *ServerIdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitChangeReplication(ctx *ChangeReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFilterDefinition(ctx *FilterDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFilterDbList(ctx *FilterDbListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFilterTableList(ctx *FilterTableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFilterStringList(ctx *FilterStringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFilterWildDbTableString(ctx *FilterWildDbTableStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFilterDbPairList(ctx *FilterDbPairListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSlave(ctx *SlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSlaveUntilOptions(ctx *SlaveUntilOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSlaveThreadOptions(ctx *SlaveThreadOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSlaveThreadOption(ctx *SlaveThreadOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGroupReplication(ctx *GroupReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPreparedStatement(ctx *PreparedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExecuteVarList(ctx *ExecuteVarListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCloneStatement(ctx *CloneStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDataDirSSL(ctx *DataDirSSLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSsl(ctx *SslContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAccountManagementStatement(ctx *AccountManagementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterUser(ctx *AlterUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterUserTail(ctx *AlterUserTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUserFunction(ctx *UserFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateUser(ctx *CreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateUserTail(ctx *CreateUserTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDefaultRoleClause(ctx *DefaultRoleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRequireClause(ctx *RequireClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitConnectOptions(ctx *ConnectOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropUser(ctx *DropUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGrant(ctx *GrantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGrantTargetList(ctx *GrantTargetListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGrantOptions(ctx *GrantOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExceptRoleList(ctx *ExceptRoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWithRoles(ctx *WithRolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGrantAs(ctx *GrantAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitVersionedRequireClause(ctx *VersionedRequireClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRenameUser(ctx *RenameUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRevoke(ctx *RevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOnTypeTo(ctx *OnTypeToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAclType(ctx *AclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleOrPrivilege(ctx *RoleOrPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGrantIdentifier(ctx *GrantIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRequireList(ctx *RequireListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRequireListElement(ctx *RequireListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGrantOption(ctx *GrantOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetRole(ctx *SetRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleList(ctx *RoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRole(ctx *RoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableAdministrationStatement(ctx *TableAdministrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitHistogram(ctx *HistogramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCheckOption(ctx *CheckOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRepairType(ctx *RepairTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInstallUninstallStatment(ctx *InstallUninstallStatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetStatement(ctx *SetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitStartOptionValueList(ctx *StartOptionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIsolationLevel(ctx *IsolationLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOptionValueListContinued(ctx *OptionValueListContinuedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOptionValue(ctx *OptionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetSystemVariable(ctx *SetSystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitShowStatement(ctx *ShowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitShowCommandType(ctx *ShowCommandTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNonBlocking(ctx *NonBlockingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFromOrIn(ctx *FromOrInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInDb(ctx *InDbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitProfileType(ctx *ProfileTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOtherAdministrativeStatement(ctx *OtherAdministrativeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyCacheListOrParts(ctx *KeyCacheListOrPartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyCacheList(ctx *KeyCacheListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAssignToKeycache(ctx *AssignToKeycacheContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAssignToKeycachePartition(ctx *AssignToKeycachePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCacheKeyList(ctx *CacheKeyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyUsageElement(ctx *KeyUsageElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyUsageList(ctx *KeyUsageListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFlushOption(ctx *FlushOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLogType(ctx *LogTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFlushTables(ctx *FlushTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFlushTablesOptions(ctx *FlushTablesOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPreloadTail(ctx *PreloadTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPreloadList(ctx *PreloadListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPreloadKeys(ctx *PreloadKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAdminPartition(ctx *AdminPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitResourceGroupManagement(ctx *ResourceGroupManagementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateResourceGroup(ctx *CreateResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitVcpuNumOrRange(ctx *VcpuNumOrRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitResourceGroupPriority(ctx *ResourceGroupPriorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterResourceGroup(ctx *AlterResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetResourceGroup(ctx *SetResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitThreadIdList(ctx *ThreadIdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDropResourceGroup(ctx *DropResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUtilityStatement(ctx *UtilityStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDescribeStatement(ctx *DescribeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExplainStatement(ctx *ExplainStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExplainableStatement(ctx *ExplainableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitHelpCommand(ctx *HelpCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUseCommand(ctx *UseCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRestartServer(ctx *RestartServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprOr(ctx *ExprOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprNot(ctx *ExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprIs(ctx *ExprIsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprAnd(ctx *ExprAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprXor(ctx *ExprXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPrimaryExprCompare(ctx *PrimaryExprCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCompOp(ctx *CompOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPredicateExprIn(ctx *PredicateExprInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPredicateExprBetween(ctx *PredicateExprBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPredicateExprLike(ctx *PredicateExprLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPredicateExprRegex(ctx *PredicateExprRegexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitBitExpr(ctx *BitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprConvert(ctx *SimpleExprConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprVariable(ctx *SimpleExprVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprCast(ctx *SimpleExprCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprUnary(ctx *SimpleExprUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprOdbc(ctx *SimpleExprOdbcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprFunction(ctx *SimpleExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprCollate(ctx *SimpleExprCollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprMatch(ctx *SimpleExprMatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprBinary(ctx *SimpleExprBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprSum(ctx *SimpleExprSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprNot(ctx *SimpleExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprValues(ctx *SimpleExprValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprDefault(ctx *SimpleExprDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprList(ctx *SimpleExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprInterval(ctx *SimpleExprIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprCase(ctx *SimpleExprCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprConcat(ctx *SimpleExprConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprLiteral(ctx *SimpleExprLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitArrayCast(ctx *ArrayCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJsonOperator(ctx *JsonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSumExpr(ctx *SumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowingClause(ctx *WindowingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLeadLagInfo(ctx *LeadLagInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNullTreatment(ctx *NullTreatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJsonFunction(ctx *JsonFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInSumExpr(ctx *InSumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentListArg(ctx *IdentListArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentList(ctx *IdentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFulltextOptions(ctx *FulltextOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGeometryFunction(ctx *GeometryFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTimeFunctionParameters(ctx *TimeFunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFractionalPrecision(ctx *FractionalPrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWeightStringLevels(ctx *WeightStringLevelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDateTimeTtype(ctx *DateTimeTtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTrimFunction(ctx *TrimFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSubstringFunction(ctx *SubstringFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSearchJsonFunction(ctx *SearchJsonFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJsonValueReturning(ctx *JsonValueReturningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitJsonValueOnError(ctx *JsonValueOnErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUdfExprList(ctx *UdfExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUdfExpr(ctx *UdfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUserVariable(ctx *UserVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSystemVariable(ctx *SystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInternalVariableName(ctx *InternalVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWhenExpression(ctx *WhenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitThenExpression(ctx *ThenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitElseExpression(ctx *ElseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCastType(ctx *CastTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCharset(ctx *CharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNotRule(ctx *NotRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNot2Rule(ctx *Not2RuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIntervalTimeStamp(ctx *IntervalTimeStampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprListWithParentheses(ctx *ExprListWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitExprWithParentheses(ctx *ExprWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOrderList(ctx *OrderListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOrderExpression(ctx *OrderExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGroupList(ctx *GroupListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGroupingExpression(ctx *GroupingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitChannel(ctx *ChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIfBody(ctx *IfBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitThenStatement(ctx *ThenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCompoundStatementList(ctx *CompoundStatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitElseStatement(ctx *ElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLabeledBlock(ctx *LabeledBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUnlabeledBlock(ctx *UnlabeledBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitBeginEndBlock(ctx *BeginEndBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLabeledControl(ctx *LabeledControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUnlabeledControl(ctx *UnlabeledControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLoopBlock(ctx *LoopBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWhileDoBlock(ctx *WhileDoBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRepeatUntilBlock(ctx *RepeatUntilBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSpDeclarations(ctx *SpDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSpDeclaration(ctx *SpDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitConditionDeclaration(ctx *ConditionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSpCondition(ctx *SpConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSqlstate(ctx *SqlstateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitHandlerDeclaration(ctx *HandlerDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitHandlerCondition(ctx *HandlerConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCursorDeclaration(ctx *CursorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIterateStatement(ctx *IterateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLeaveStatement(ctx *LeaveStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGetDiagnostics(ctx *GetDiagnosticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSignalAllowedExpr(ctx *SignalAllowedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitStatementInformationItem(ctx *StatementInformationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitConditionInformationItem(ctx *ConditionInformationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSignalInformationItemName(ctx *SignalInformationItemNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSignalStatement(ctx *SignalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitResignalStatement(ctx *ResignalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSignalInformationItem(ctx *SignalInformationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCursorOpen(ctx *CursorOpenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCursorClose(ctx *CursorCloseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCursorFetch(ctx *CursorFetchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSchedule(ctx *ScheduleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCheckOrReferences(ctx *CheckOrReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCheckConstraint(ctx *CheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitConstraintEnforcement(ctx *ConstraintEnforcementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableConstraintDef(ctx *TableConstraintDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFieldDefinition(ctx *FieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnAttribute(ctx *ColumnAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnFormat(ctx *ColumnFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitStorageMedia(ctx *StorageMediaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitGcolAttribute(ctx *GcolAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReferences(ctx *ReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDeleteOption(ctx *DeleteOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyList(ctx *KeyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyPart(ctx *KeyPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyListWithExpression(ctx *KeyListWithExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyPartOrExpression(ctx *KeyPartOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitKeyListVariants(ctx *KeyListVariantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexOption(ctx *IndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCommonIndexOption(ctx *CommonIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitVisibility(ctx *VisibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexTypeClause(ctx *IndexTypeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFulltextIndexOption(ctx *FulltextIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSpatialIndexOption(ctx *SpatialIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDataTypeDefinition(ctx *DataTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNchar(ctx *NcharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRealType(ctx *RealTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFieldLength(ctx *FieldLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFieldOptions(ctx *FieldOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAscii(ctx *AsciiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUnicode(ctx *UnicodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWsNumCodepoints(ctx *WsNumCodepointsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateTableOption(ctx *CreateTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTernaryOption(ctx *TernaryOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDefaultCollation(ctx *DefaultCollationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDefaultEncryption(ctx *DefaultEncryptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDefaultCharset(ctx *DefaultCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionClause(ctx *PartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionDefKey(ctx *PartitionDefKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionDefHash(ctx *PartitionDefHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionDefRangeList(ctx *PartitionDefRangeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSubPartitions(ctx *SubPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionValuesIn(ctx *PartitionValuesInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionOption(ctx *PartitionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPartitionValueItem(ctx *PartitionValueItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDefinerClause(ctx *DefinerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIfExists(ctx *IfExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitProcedureParameter(ctx *ProcedureParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCollate(ctx *CollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTypeWithOptCollate(ctx *TypeWithOptCollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewRefList(ctx *ViewRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUpdateList(ctx *UpdateListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUpdateElement(ctx *UpdateElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCharsetClause(ctx *CharsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFieldsClause(ctx *FieldsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFieldTerm(ctx *FieldTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLinesClause(ctx *LinesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLineTerm(ctx *LineTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUserList(ctx *UserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateUserList(ctx *CreateUserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterUserList(ctx *AlterUserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitCreateUserEntry(ctx *CreateUserEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitAlterUserEntry(ctx *AlterUserEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRetainCurrentPassword(ctx *RetainCurrentPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDiscardOldPassword(ctx *DiscardOldPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReplacePassword(ctx *ReplacePasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUser(ctx *UserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLikeClause(ctx *LikeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLikeOrWhere(ctx *LikeOrWhereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOnlineOption(ctx *OnlineOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNoWriteToBinLog(ctx *NoWriteToBinLogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUsePartition(ctx *UsePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFieldIdentifier(ctx *FieldIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnInternalRef(ctx *ColumnInternalRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnInternalRefList(ctx *ColumnInternalRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitColumnRef(ctx *ColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitInsertIdentifier(ctx *InsertIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIndexRef(ctx *IndexRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableWild(ctx *TableWildContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSchemaName(ctx *SchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSchemaRef(ctx *SchemaRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitProcedureName(ctx *ProcedureNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitProcedureRef(ctx *ProcedureRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFunctionRef(ctx *FunctionRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTriggerName(ctx *TriggerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTriggerRef(ctx *TriggerRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitViewRef(ctx *ViewRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTablespaceName(ctx *TablespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTablespaceRef(ctx *TablespaceRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLogfileGroupName(ctx *LogfileGroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLogfileGroupRef(ctx *LogfileGroupRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitEventName(ctx *EventNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitEventRef(ctx *EventRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUdfName(ctx *UdfNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitServerName(ctx *ServerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitServerRef(ctx *ServerRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitEngineRef(ctx *EngineRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFilterTableRef(ctx *FilterTableRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableRefWithWildcard(ctx *TableRefWithWildcardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableRef(ctx *TableRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableRefList(ctx *TableRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTableAliasRefList(ctx *TableAliasRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitParameterName(ctx *ParameterNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLabelIdentifier(ctx *LabelIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLabelRef(ctx *LabelRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleIdentifier(ctx *RoleIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleRef(ctx *RoleRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPluginRef(ctx *PluginRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitComponentRef(ctx *ComponentRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitResourceGroupRef(ctx *ResourceGroupRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitWindowName(ctx *WindowNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPureIdentifier(ctx *PureIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSimpleIdentifier(ctx *SimpleIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitDotIdentifier(ctx *DotIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUlong_number(ctx *Ulong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReal_ulong_number(ctx *Real_ulong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitUlonglong_number(ctx *Ulonglong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitReal_ulonglong_number(ctx *Real_ulonglong_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSignedLiteral(ctx *SignedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTextStringLiteral(ctx *TextStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTextString(ctx *TextStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTextStringHash(ctx *TextStringHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTextLiteral(ctx *TextLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTextStringLiteralList(ctx *TextStringLiteralListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNumLiteral(ctx *NumLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTemporalLiteral(ctx *TemporalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitFloatOptions(ctx *FloatOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitStandardFloatOptions(ctx *StandardFloatOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitPrecision(ctx *PrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitTextOrIdentifier(ctx *TextOrIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLValueIdentifier(ctx *LValueIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSizeNumber(ctx *SizeNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitParentheses(ctx *ParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitEqual(ctx *EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitOptionType(ctx *OptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitVarIdentType(ctx *VarIdentTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitSetVarIdentType(ctx *SetVarIdentTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierKeyword(ctx *IdentifierKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLabelKeyword(ctx *LabelKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleKeyword(ctx *RoleKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitLValueKeyword(ctx *LValueKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLParserVisitor) VisitRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
