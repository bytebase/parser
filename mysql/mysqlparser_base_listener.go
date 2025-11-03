// Code generated from MySQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mysql // MySQLParser
import "github.com/antlr4-go/antlr/v4"

// BaseMySQLParserListener is a complete listener for a parse tree produced by MySQLParser.
type BaseMySQLParserListener struct{}

var _ MySQLParserListener = &BaseMySQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMySQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMySQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMySQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMySQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseMySQLParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseMySQLParserListener) ExitScript(ctx *ScriptContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseMySQLParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseMySQLParserListener) ExitQuery(ctx *QueryContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseMySQLParserListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseMySQLParserListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterAlterStatement is called when production alterStatement is entered.
func (s *BaseMySQLParserListener) EnterAlterStatement(ctx *AlterStatementContext) {}

// ExitAlterStatement is called when production alterStatement is exited.
func (s *BaseMySQLParserListener) ExitAlterStatement(ctx *AlterStatementContext) {}

// EnterAlterDatabase is called when production alterDatabase is entered.
func (s *BaseMySQLParserListener) EnterAlterDatabase(ctx *AlterDatabaseContext) {}

// ExitAlterDatabase is called when production alterDatabase is exited.
func (s *BaseMySQLParserListener) ExitAlterDatabase(ctx *AlterDatabaseContext) {}

// EnterAlterDatabaseOption is called when production alterDatabaseOption is entered.
func (s *BaseMySQLParserListener) EnterAlterDatabaseOption(ctx *AlterDatabaseOptionContext) {}

// ExitAlterDatabaseOption is called when production alterDatabaseOption is exited.
func (s *BaseMySQLParserListener) ExitAlterDatabaseOption(ctx *AlterDatabaseOptionContext) {}

// EnterAlterEvent is called when production alterEvent is entered.
func (s *BaseMySQLParserListener) EnterAlterEvent(ctx *AlterEventContext) {}

// ExitAlterEvent is called when production alterEvent is exited.
func (s *BaseMySQLParserListener) ExitAlterEvent(ctx *AlterEventContext) {}

// EnterAlterLogfileGroup is called when production alterLogfileGroup is entered.
func (s *BaseMySQLParserListener) EnterAlterLogfileGroup(ctx *AlterLogfileGroupContext) {}

// ExitAlterLogfileGroup is called when production alterLogfileGroup is exited.
func (s *BaseMySQLParserListener) ExitAlterLogfileGroup(ctx *AlterLogfileGroupContext) {}

// EnterAlterLogfileGroupOptions is called when production alterLogfileGroupOptions is entered.
func (s *BaseMySQLParserListener) EnterAlterLogfileGroupOptions(ctx *AlterLogfileGroupOptionsContext) {
}

// ExitAlterLogfileGroupOptions is called when production alterLogfileGroupOptions is exited.
func (s *BaseMySQLParserListener) ExitAlterLogfileGroupOptions(ctx *AlterLogfileGroupOptionsContext) {
}

// EnterAlterLogfileGroupOption is called when production alterLogfileGroupOption is entered.
func (s *BaseMySQLParserListener) EnterAlterLogfileGroupOption(ctx *AlterLogfileGroupOptionContext) {}

// ExitAlterLogfileGroupOption is called when production alterLogfileGroupOption is exited.
func (s *BaseMySQLParserListener) ExitAlterLogfileGroupOption(ctx *AlterLogfileGroupOptionContext) {}

// EnterAlterServer is called when production alterServer is entered.
func (s *BaseMySQLParserListener) EnterAlterServer(ctx *AlterServerContext) {}

// ExitAlterServer is called when production alterServer is exited.
func (s *BaseMySQLParserListener) ExitAlterServer(ctx *AlterServerContext) {}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseMySQLParserListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseMySQLParserListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterAlterTableActions is called when production alterTableActions is entered.
func (s *BaseMySQLParserListener) EnterAlterTableActions(ctx *AlterTableActionsContext) {}

// ExitAlterTableActions is called when production alterTableActions is exited.
func (s *BaseMySQLParserListener) ExitAlterTableActions(ctx *AlterTableActionsContext) {}

// EnterAlterCommandList is called when production alterCommandList is entered.
func (s *BaseMySQLParserListener) EnterAlterCommandList(ctx *AlterCommandListContext) {}

// ExitAlterCommandList is called when production alterCommandList is exited.
func (s *BaseMySQLParserListener) ExitAlterCommandList(ctx *AlterCommandListContext) {}

// EnterAlterCommandsModifierList is called when production alterCommandsModifierList is entered.
func (s *BaseMySQLParserListener) EnterAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) {
}

// ExitAlterCommandsModifierList is called when production alterCommandsModifierList is exited.
func (s *BaseMySQLParserListener) ExitAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) {
}

// EnterStandaloneAlterCommands is called when production standaloneAlterCommands is entered.
func (s *BaseMySQLParserListener) EnterStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) {}

// ExitStandaloneAlterCommands is called when production standaloneAlterCommands is exited.
func (s *BaseMySQLParserListener) ExitStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) {}

// EnterAlterPartition is called when production alterPartition is entered.
func (s *BaseMySQLParserListener) EnterAlterPartition(ctx *AlterPartitionContext) {}

// ExitAlterPartition is called when production alterPartition is exited.
func (s *BaseMySQLParserListener) ExitAlterPartition(ctx *AlterPartitionContext) {}

// EnterAlterList is called when production alterList is entered.
func (s *BaseMySQLParserListener) EnterAlterList(ctx *AlterListContext) {}

// ExitAlterList is called when production alterList is exited.
func (s *BaseMySQLParserListener) ExitAlterList(ctx *AlterListContext) {}

// EnterAlterCommandsModifier is called when production alterCommandsModifier is entered.
func (s *BaseMySQLParserListener) EnterAlterCommandsModifier(ctx *AlterCommandsModifierContext) {}

// ExitAlterCommandsModifier is called when production alterCommandsModifier is exited.
func (s *BaseMySQLParserListener) ExitAlterCommandsModifier(ctx *AlterCommandsModifierContext) {}

// EnterAlterListItem is called when production alterListItem is entered.
func (s *BaseMySQLParserListener) EnterAlterListItem(ctx *AlterListItemContext) {}

// ExitAlterListItem is called when production alterListItem is exited.
func (s *BaseMySQLParserListener) ExitAlterListItem(ctx *AlterListItemContext) {}

// EnterPlace is called when production place is entered.
func (s *BaseMySQLParserListener) EnterPlace(ctx *PlaceContext) {}

// ExitPlace is called when production place is exited.
func (s *BaseMySQLParserListener) ExitPlace(ctx *PlaceContext) {}

// EnterRestrict is called when production restrict is entered.
func (s *BaseMySQLParserListener) EnterRestrict(ctx *RestrictContext) {}

// ExitRestrict is called when production restrict is exited.
func (s *BaseMySQLParserListener) ExitRestrict(ctx *RestrictContext) {}

// EnterAlterOrderList is called when production alterOrderList is entered.
func (s *BaseMySQLParserListener) EnterAlterOrderList(ctx *AlterOrderListContext) {}

// ExitAlterOrderList is called when production alterOrderList is exited.
func (s *BaseMySQLParserListener) ExitAlterOrderList(ctx *AlterOrderListContext) {}

// EnterAlterAlgorithmOption is called when production alterAlgorithmOption is entered.
func (s *BaseMySQLParserListener) EnterAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) {}

// ExitAlterAlgorithmOption is called when production alterAlgorithmOption is exited.
func (s *BaseMySQLParserListener) ExitAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) {}

// EnterAlterLockOption is called when production alterLockOption is entered.
func (s *BaseMySQLParserListener) EnterAlterLockOption(ctx *AlterLockOptionContext) {}

// ExitAlterLockOption is called when production alterLockOption is exited.
func (s *BaseMySQLParserListener) ExitAlterLockOption(ctx *AlterLockOptionContext) {}

// EnterIndexLockAndAlgorithm is called when production indexLockAndAlgorithm is entered.
func (s *BaseMySQLParserListener) EnterIndexLockAndAlgorithm(ctx *IndexLockAndAlgorithmContext) {}

// ExitIndexLockAndAlgorithm is called when production indexLockAndAlgorithm is exited.
func (s *BaseMySQLParserListener) ExitIndexLockAndAlgorithm(ctx *IndexLockAndAlgorithmContext) {}

// EnterWithValidation is called when production withValidation is entered.
func (s *BaseMySQLParserListener) EnterWithValidation(ctx *WithValidationContext) {}

// ExitWithValidation is called when production withValidation is exited.
func (s *BaseMySQLParserListener) ExitWithValidation(ctx *WithValidationContext) {}

// EnterRemovePartitioning is called when production removePartitioning is entered.
func (s *BaseMySQLParserListener) EnterRemovePartitioning(ctx *RemovePartitioningContext) {}

// ExitRemovePartitioning is called when production removePartitioning is exited.
func (s *BaseMySQLParserListener) ExitRemovePartitioning(ctx *RemovePartitioningContext) {}

// EnterAllOrPartitionNameList is called when production allOrPartitionNameList is entered.
func (s *BaseMySQLParserListener) EnterAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) {}

// ExitAllOrPartitionNameList is called when production allOrPartitionNameList is exited.
func (s *BaseMySQLParserListener) ExitAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) {}

// EnterAlterTablespace is called when production alterTablespace is entered.
func (s *BaseMySQLParserListener) EnterAlterTablespace(ctx *AlterTablespaceContext) {}

// ExitAlterTablespace is called when production alterTablespace is exited.
func (s *BaseMySQLParserListener) ExitAlterTablespace(ctx *AlterTablespaceContext) {}

// EnterAlterUndoTablespace is called when production alterUndoTablespace is entered.
func (s *BaseMySQLParserListener) EnterAlterUndoTablespace(ctx *AlterUndoTablespaceContext) {}

// ExitAlterUndoTablespace is called when production alterUndoTablespace is exited.
func (s *BaseMySQLParserListener) ExitAlterUndoTablespace(ctx *AlterUndoTablespaceContext) {}

// EnterUndoTableSpaceOptions is called when production undoTableSpaceOptions is entered.
func (s *BaseMySQLParserListener) EnterUndoTableSpaceOptions(ctx *UndoTableSpaceOptionsContext) {}

// ExitUndoTableSpaceOptions is called when production undoTableSpaceOptions is exited.
func (s *BaseMySQLParserListener) ExitUndoTableSpaceOptions(ctx *UndoTableSpaceOptionsContext) {}

// EnterUndoTableSpaceOption is called when production undoTableSpaceOption is entered.
func (s *BaseMySQLParserListener) EnterUndoTableSpaceOption(ctx *UndoTableSpaceOptionContext) {}

// ExitUndoTableSpaceOption is called when production undoTableSpaceOption is exited.
func (s *BaseMySQLParserListener) ExitUndoTableSpaceOption(ctx *UndoTableSpaceOptionContext) {}

// EnterAlterTablespaceOptions is called when production alterTablespaceOptions is entered.
func (s *BaseMySQLParserListener) EnterAlterTablespaceOptions(ctx *AlterTablespaceOptionsContext) {}

// ExitAlterTablespaceOptions is called when production alterTablespaceOptions is exited.
func (s *BaseMySQLParserListener) ExitAlterTablespaceOptions(ctx *AlterTablespaceOptionsContext) {}

// EnterAlterTablespaceOption is called when production alterTablespaceOption is entered.
func (s *BaseMySQLParserListener) EnterAlterTablespaceOption(ctx *AlterTablespaceOptionContext) {}

// ExitAlterTablespaceOption is called when production alterTablespaceOption is exited.
func (s *BaseMySQLParserListener) ExitAlterTablespaceOption(ctx *AlterTablespaceOptionContext) {}

// EnterChangeTablespaceOption is called when production changeTablespaceOption is entered.
func (s *BaseMySQLParserListener) EnterChangeTablespaceOption(ctx *ChangeTablespaceOptionContext) {}

// ExitChangeTablespaceOption is called when production changeTablespaceOption is exited.
func (s *BaseMySQLParserListener) ExitChangeTablespaceOption(ctx *ChangeTablespaceOptionContext) {}

// EnterAlterView is called when production alterView is entered.
func (s *BaseMySQLParserListener) EnterAlterView(ctx *AlterViewContext) {}

// ExitAlterView is called when production alterView is exited.
func (s *BaseMySQLParserListener) ExitAlterView(ctx *AlterViewContext) {}

// EnterViewTail is called when production viewTail is entered.
func (s *BaseMySQLParserListener) EnterViewTail(ctx *ViewTailContext) {}

// ExitViewTail is called when production viewTail is exited.
func (s *BaseMySQLParserListener) ExitViewTail(ctx *ViewTailContext) {}

// EnterViewSelect is called when production viewSelect is entered.
func (s *BaseMySQLParserListener) EnterViewSelect(ctx *ViewSelectContext) {}

// ExitViewSelect is called when production viewSelect is exited.
func (s *BaseMySQLParserListener) ExitViewSelect(ctx *ViewSelectContext) {}

// EnterViewCheckOption is called when production viewCheckOption is entered.
func (s *BaseMySQLParserListener) EnterViewCheckOption(ctx *ViewCheckOptionContext) {}

// ExitViewCheckOption is called when production viewCheckOption is exited.
func (s *BaseMySQLParserListener) ExitViewCheckOption(ctx *ViewCheckOptionContext) {}

// EnterAlterProcedure is called when production alterProcedure is entered.
func (s *BaseMySQLParserListener) EnterAlterProcedure(ctx *AlterProcedureContext) {}

// ExitAlterProcedure is called when production alterProcedure is exited.
func (s *BaseMySQLParserListener) ExitAlterProcedure(ctx *AlterProcedureContext) {}

// EnterAlterFunction is called when production alterFunction is entered.
func (s *BaseMySQLParserListener) EnterAlterFunction(ctx *AlterFunctionContext) {}

// ExitAlterFunction is called when production alterFunction is exited.
func (s *BaseMySQLParserListener) ExitAlterFunction(ctx *AlterFunctionContext) {}

// EnterCreateStatement is called when production createStatement is entered.
func (s *BaseMySQLParserListener) EnterCreateStatement(ctx *CreateStatementContext) {}

// ExitCreateStatement is called when production createStatement is exited.
func (s *BaseMySQLParserListener) ExitCreateStatement(ctx *CreateStatementContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseMySQLParserListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseMySQLParserListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterCreateDatabaseOption is called when production createDatabaseOption is entered.
func (s *BaseMySQLParserListener) EnterCreateDatabaseOption(ctx *CreateDatabaseOptionContext) {}

// ExitCreateDatabaseOption is called when production createDatabaseOption is exited.
func (s *BaseMySQLParserListener) ExitCreateDatabaseOption(ctx *CreateDatabaseOptionContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseMySQLParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseMySQLParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterTableElementList is called when production tableElementList is entered.
func (s *BaseMySQLParserListener) EnterTableElementList(ctx *TableElementListContext) {}

// ExitTableElementList is called when production tableElementList is exited.
func (s *BaseMySQLParserListener) ExitTableElementList(ctx *TableElementListContext) {}

// EnterTableElement is called when production tableElement is entered.
func (s *BaseMySQLParserListener) EnterTableElement(ctx *TableElementContext) {}

// ExitTableElement is called when production tableElement is exited.
func (s *BaseMySQLParserListener) ExitTableElement(ctx *TableElementContext) {}

// EnterDuplicateAsQueryExpression is called when production duplicateAsQueryExpression is entered.
func (s *BaseMySQLParserListener) EnterDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) {
}

// ExitDuplicateAsQueryExpression is called when production duplicateAsQueryExpression is exited.
func (s *BaseMySQLParserListener) ExitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) {
}

// EnterQueryExpressionOrParens is called when production queryExpressionOrParens is entered.
func (s *BaseMySQLParserListener) EnterQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) {}

// ExitQueryExpressionOrParens is called when production queryExpressionOrParens is exited.
func (s *BaseMySQLParserListener) ExitQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) {}

// EnterCreateRoutine is called when production createRoutine is entered.
func (s *BaseMySQLParserListener) EnterCreateRoutine(ctx *CreateRoutineContext) {}

// ExitCreateRoutine is called when production createRoutine is exited.
func (s *BaseMySQLParserListener) ExitCreateRoutine(ctx *CreateRoutineContext) {}

// EnterCreateProcedure is called when production createProcedure is entered.
func (s *BaseMySQLParserListener) EnterCreateProcedure(ctx *CreateProcedureContext) {}

// ExitCreateProcedure is called when production createProcedure is exited.
func (s *BaseMySQLParserListener) ExitCreateProcedure(ctx *CreateProcedureContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseMySQLParserListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseMySQLParserListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterCreateUdf is called when production createUdf is entered.
func (s *BaseMySQLParserListener) EnterCreateUdf(ctx *CreateUdfContext) {}

// ExitCreateUdf is called when production createUdf is exited.
func (s *BaseMySQLParserListener) ExitCreateUdf(ctx *CreateUdfContext) {}

// EnterRoutineCreateOption is called when production routineCreateOption is entered.
func (s *BaseMySQLParserListener) EnterRoutineCreateOption(ctx *RoutineCreateOptionContext) {}

// ExitRoutineCreateOption is called when production routineCreateOption is exited.
func (s *BaseMySQLParserListener) ExitRoutineCreateOption(ctx *RoutineCreateOptionContext) {}

// EnterRoutineAlterOptions is called when production routineAlterOptions is entered.
func (s *BaseMySQLParserListener) EnterRoutineAlterOptions(ctx *RoutineAlterOptionsContext) {}

// ExitRoutineAlterOptions is called when production routineAlterOptions is exited.
func (s *BaseMySQLParserListener) ExitRoutineAlterOptions(ctx *RoutineAlterOptionsContext) {}

// EnterRoutineOption is called when production routineOption is entered.
func (s *BaseMySQLParserListener) EnterRoutineOption(ctx *RoutineOptionContext) {}

// ExitRoutineOption is called when production routineOption is exited.
func (s *BaseMySQLParserListener) ExitRoutineOption(ctx *RoutineOptionContext) {}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseMySQLParserListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseMySQLParserListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterIndexNameAndType is called when production indexNameAndType is entered.
func (s *BaseMySQLParserListener) EnterIndexNameAndType(ctx *IndexNameAndTypeContext) {}

// ExitIndexNameAndType is called when production indexNameAndType is exited.
func (s *BaseMySQLParserListener) ExitIndexNameAndType(ctx *IndexNameAndTypeContext) {}

// EnterCreateIndexTarget is called when production createIndexTarget is entered.
func (s *BaseMySQLParserListener) EnterCreateIndexTarget(ctx *CreateIndexTargetContext) {}

// ExitCreateIndexTarget is called when production createIndexTarget is exited.
func (s *BaseMySQLParserListener) ExitCreateIndexTarget(ctx *CreateIndexTargetContext) {}

// EnterCreateLogfileGroup is called when production createLogfileGroup is entered.
func (s *BaseMySQLParserListener) EnterCreateLogfileGroup(ctx *CreateLogfileGroupContext) {}

// ExitCreateLogfileGroup is called when production createLogfileGroup is exited.
func (s *BaseMySQLParserListener) ExitCreateLogfileGroup(ctx *CreateLogfileGroupContext) {}

// EnterLogfileGroupOptions is called when production logfileGroupOptions is entered.
func (s *BaseMySQLParserListener) EnterLogfileGroupOptions(ctx *LogfileGroupOptionsContext) {}

// ExitLogfileGroupOptions is called when production logfileGroupOptions is exited.
func (s *BaseMySQLParserListener) ExitLogfileGroupOptions(ctx *LogfileGroupOptionsContext) {}

// EnterLogfileGroupOption is called when production logfileGroupOption is entered.
func (s *BaseMySQLParserListener) EnterLogfileGroupOption(ctx *LogfileGroupOptionContext) {}

// ExitLogfileGroupOption is called when production logfileGroupOption is exited.
func (s *BaseMySQLParserListener) ExitLogfileGroupOption(ctx *LogfileGroupOptionContext) {}

// EnterCreateServer is called when production createServer is entered.
func (s *BaseMySQLParserListener) EnterCreateServer(ctx *CreateServerContext) {}

// ExitCreateServer is called when production createServer is exited.
func (s *BaseMySQLParserListener) ExitCreateServer(ctx *CreateServerContext) {}

// EnterServerOptions is called when production serverOptions is entered.
func (s *BaseMySQLParserListener) EnterServerOptions(ctx *ServerOptionsContext) {}

// ExitServerOptions is called when production serverOptions is exited.
func (s *BaseMySQLParserListener) ExitServerOptions(ctx *ServerOptionsContext) {}

// EnterServerOption is called when production serverOption is entered.
func (s *BaseMySQLParserListener) EnterServerOption(ctx *ServerOptionContext) {}

// ExitServerOption is called when production serverOption is exited.
func (s *BaseMySQLParserListener) ExitServerOption(ctx *ServerOptionContext) {}

// EnterCreateTablespace is called when production createTablespace is entered.
func (s *BaseMySQLParserListener) EnterCreateTablespace(ctx *CreateTablespaceContext) {}

// ExitCreateTablespace is called when production createTablespace is exited.
func (s *BaseMySQLParserListener) ExitCreateTablespace(ctx *CreateTablespaceContext) {}

// EnterCreateUndoTablespace is called when production createUndoTablespace is entered.
func (s *BaseMySQLParserListener) EnterCreateUndoTablespace(ctx *CreateUndoTablespaceContext) {}

// ExitCreateUndoTablespace is called when production createUndoTablespace is exited.
func (s *BaseMySQLParserListener) ExitCreateUndoTablespace(ctx *CreateUndoTablespaceContext) {}

// EnterTsDataFileName is called when production tsDataFileName is entered.
func (s *BaseMySQLParserListener) EnterTsDataFileName(ctx *TsDataFileNameContext) {}

// ExitTsDataFileName is called when production tsDataFileName is exited.
func (s *BaseMySQLParserListener) ExitTsDataFileName(ctx *TsDataFileNameContext) {}

// EnterTsDataFile is called when production tsDataFile is entered.
func (s *BaseMySQLParserListener) EnterTsDataFile(ctx *TsDataFileContext) {}

// ExitTsDataFile is called when production tsDataFile is exited.
func (s *BaseMySQLParserListener) ExitTsDataFile(ctx *TsDataFileContext) {}

// EnterTablespaceOptions is called when production tablespaceOptions is entered.
func (s *BaseMySQLParserListener) EnterTablespaceOptions(ctx *TablespaceOptionsContext) {}

// ExitTablespaceOptions is called when production tablespaceOptions is exited.
func (s *BaseMySQLParserListener) ExitTablespaceOptions(ctx *TablespaceOptionsContext) {}

// EnterTablespaceOption is called when production tablespaceOption is entered.
func (s *BaseMySQLParserListener) EnterTablespaceOption(ctx *TablespaceOptionContext) {}

// ExitTablespaceOption is called when production tablespaceOption is exited.
func (s *BaseMySQLParserListener) ExitTablespaceOption(ctx *TablespaceOptionContext) {}

// EnterTsOptionInitialSize is called when production tsOptionInitialSize is entered.
func (s *BaseMySQLParserListener) EnterTsOptionInitialSize(ctx *TsOptionInitialSizeContext) {}

// ExitTsOptionInitialSize is called when production tsOptionInitialSize is exited.
func (s *BaseMySQLParserListener) ExitTsOptionInitialSize(ctx *TsOptionInitialSizeContext) {}

// EnterTsOptionUndoRedoBufferSize is called when production tsOptionUndoRedoBufferSize is entered.
func (s *BaseMySQLParserListener) EnterTsOptionUndoRedoBufferSize(ctx *TsOptionUndoRedoBufferSizeContext) {
}

// ExitTsOptionUndoRedoBufferSize is called when production tsOptionUndoRedoBufferSize is exited.
func (s *BaseMySQLParserListener) ExitTsOptionUndoRedoBufferSize(ctx *TsOptionUndoRedoBufferSizeContext) {
}

// EnterTsOptionAutoextendSize is called when production tsOptionAutoextendSize is entered.
func (s *BaseMySQLParserListener) EnterTsOptionAutoextendSize(ctx *TsOptionAutoextendSizeContext) {}

// ExitTsOptionAutoextendSize is called when production tsOptionAutoextendSize is exited.
func (s *BaseMySQLParserListener) ExitTsOptionAutoextendSize(ctx *TsOptionAutoextendSizeContext) {}

// EnterTsOptionMaxSize is called when production tsOptionMaxSize is entered.
func (s *BaseMySQLParserListener) EnterTsOptionMaxSize(ctx *TsOptionMaxSizeContext) {}

// ExitTsOptionMaxSize is called when production tsOptionMaxSize is exited.
func (s *BaseMySQLParserListener) ExitTsOptionMaxSize(ctx *TsOptionMaxSizeContext) {}

// EnterTsOptionExtentSize is called when production tsOptionExtentSize is entered.
func (s *BaseMySQLParserListener) EnterTsOptionExtentSize(ctx *TsOptionExtentSizeContext) {}

// ExitTsOptionExtentSize is called when production tsOptionExtentSize is exited.
func (s *BaseMySQLParserListener) ExitTsOptionExtentSize(ctx *TsOptionExtentSizeContext) {}

// EnterTsOptionNodegroup is called when production tsOptionNodegroup is entered.
func (s *BaseMySQLParserListener) EnterTsOptionNodegroup(ctx *TsOptionNodegroupContext) {}

// ExitTsOptionNodegroup is called when production tsOptionNodegroup is exited.
func (s *BaseMySQLParserListener) ExitTsOptionNodegroup(ctx *TsOptionNodegroupContext) {}

// EnterTsOptionEngine is called when production tsOptionEngine is entered.
func (s *BaseMySQLParserListener) EnterTsOptionEngine(ctx *TsOptionEngineContext) {}

// ExitTsOptionEngine is called when production tsOptionEngine is exited.
func (s *BaseMySQLParserListener) ExitTsOptionEngine(ctx *TsOptionEngineContext) {}

// EnterTsOptionWait is called when production tsOptionWait is entered.
func (s *BaseMySQLParserListener) EnterTsOptionWait(ctx *TsOptionWaitContext) {}

// ExitTsOptionWait is called when production tsOptionWait is exited.
func (s *BaseMySQLParserListener) ExitTsOptionWait(ctx *TsOptionWaitContext) {}

// EnterTsOptionComment is called when production tsOptionComment is entered.
func (s *BaseMySQLParserListener) EnterTsOptionComment(ctx *TsOptionCommentContext) {}

// ExitTsOptionComment is called when production tsOptionComment is exited.
func (s *BaseMySQLParserListener) ExitTsOptionComment(ctx *TsOptionCommentContext) {}

// EnterTsOptionFileblockSize is called when production tsOptionFileblockSize is entered.
func (s *BaseMySQLParserListener) EnterTsOptionFileblockSize(ctx *TsOptionFileblockSizeContext) {}

// ExitTsOptionFileblockSize is called when production tsOptionFileblockSize is exited.
func (s *BaseMySQLParserListener) ExitTsOptionFileblockSize(ctx *TsOptionFileblockSizeContext) {}

// EnterTsOptionEncryption is called when production tsOptionEncryption is entered.
func (s *BaseMySQLParserListener) EnterTsOptionEncryption(ctx *TsOptionEncryptionContext) {}

// ExitTsOptionEncryption is called when production tsOptionEncryption is exited.
func (s *BaseMySQLParserListener) ExitTsOptionEncryption(ctx *TsOptionEncryptionContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseMySQLParserListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseMySQLParserListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterViewReplaceOrAlgorithm is called when production viewReplaceOrAlgorithm is entered.
func (s *BaseMySQLParserListener) EnterViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) {}

// ExitViewReplaceOrAlgorithm is called when production viewReplaceOrAlgorithm is exited.
func (s *BaseMySQLParserListener) ExitViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) {}

// EnterViewAlgorithm is called when production viewAlgorithm is entered.
func (s *BaseMySQLParserListener) EnterViewAlgorithm(ctx *ViewAlgorithmContext) {}

// ExitViewAlgorithm is called when production viewAlgorithm is exited.
func (s *BaseMySQLParserListener) ExitViewAlgorithm(ctx *ViewAlgorithmContext) {}

// EnterViewSuid is called when production viewSuid is entered.
func (s *BaseMySQLParserListener) EnterViewSuid(ctx *ViewSuidContext) {}

// ExitViewSuid is called when production viewSuid is exited.
func (s *BaseMySQLParserListener) ExitViewSuid(ctx *ViewSuidContext) {}

// EnterCreateTrigger is called when production createTrigger is entered.
func (s *BaseMySQLParserListener) EnterCreateTrigger(ctx *CreateTriggerContext) {}

// ExitCreateTrigger is called when production createTrigger is exited.
func (s *BaseMySQLParserListener) ExitCreateTrigger(ctx *CreateTriggerContext) {}

// EnterTriggerFollowsPrecedesClause is called when production triggerFollowsPrecedesClause is entered.
func (s *BaseMySQLParserListener) EnterTriggerFollowsPrecedesClause(ctx *TriggerFollowsPrecedesClauseContext) {
}

// ExitTriggerFollowsPrecedesClause is called when production triggerFollowsPrecedesClause is exited.
func (s *BaseMySQLParserListener) ExitTriggerFollowsPrecedesClause(ctx *TriggerFollowsPrecedesClauseContext) {
}

// EnterCreateEvent is called when production createEvent is entered.
func (s *BaseMySQLParserListener) EnterCreateEvent(ctx *CreateEventContext) {}

// ExitCreateEvent is called when production createEvent is exited.
func (s *BaseMySQLParserListener) ExitCreateEvent(ctx *CreateEventContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseMySQLParserListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseMySQLParserListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterCreateSpatialReference is called when production createSpatialReference is entered.
func (s *BaseMySQLParserListener) EnterCreateSpatialReference(ctx *CreateSpatialReferenceContext) {}

// ExitCreateSpatialReference is called when production createSpatialReference is exited.
func (s *BaseMySQLParserListener) ExitCreateSpatialReference(ctx *CreateSpatialReferenceContext) {}

// EnterSrsAttribute is called when production srsAttribute is entered.
func (s *BaseMySQLParserListener) EnterSrsAttribute(ctx *SrsAttributeContext) {}

// ExitSrsAttribute is called when production srsAttribute is exited.
func (s *BaseMySQLParserListener) ExitSrsAttribute(ctx *SrsAttributeContext) {}

// EnterDropStatement is called when production dropStatement is entered.
func (s *BaseMySQLParserListener) EnterDropStatement(ctx *DropStatementContext) {}

// ExitDropStatement is called when production dropStatement is exited.
func (s *BaseMySQLParserListener) ExitDropStatement(ctx *DropStatementContext) {}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseMySQLParserListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseMySQLParserListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterDropEvent is called when production dropEvent is entered.
func (s *BaseMySQLParserListener) EnterDropEvent(ctx *DropEventContext) {}

// ExitDropEvent is called when production dropEvent is exited.
func (s *BaseMySQLParserListener) ExitDropEvent(ctx *DropEventContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseMySQLParserListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseMySQLParserListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterDropProcedure is called when production dropProcedure is entered.
func (s *BaseMySQLParserListener) EnterDropProcedure(ctx *DropProcedureContext) {}

// ExitDropProcedure is called when production dropProcedure is exited.
func (s *BaseMySQLParserListener) ExitDropProcedure(ctx *DropProcedureContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseMySQLParserListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseMySQLParserListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterDropLogfileGroup is called when production dropLogfileGroup is entered.
func (s *BaseMySQLParserListener) EnterDropLogfileGroup(ctx *DropLogfileGroupContext) {}

// ExitDropLogfileGroup is called when production dropLogfileGroup is exited.
func (s *BaseMySQLParserListener) ExitDropLogfileGroup(ctx *DropLogfileGroupContext) {}

// EnterDropLogfileGroupOption is called when production dropLogfileGroupOption is entered.
func (s *BaseMySQLParserListener) EnterDropLogfileGroupOption(ctx *DropLogfileGroupOptionContext) {}

// ExitDropLogfileGroupOption is called when production dropLogfileGroupOption is exited.
func (s *BaseMySQLParserListener) ExitDropLogfileGroupOption(ctx *DropLogfileGroupOptionContext) {}

// EnterDropServer is called when production dropServer is entered.
func (s *BaseMySQLParserListener) EnterDropServer(ctx *DropServerContext) {}

// ExitDropServer is called when production dropServer is exited.
func (s *BaseMySQLParserListener) ExitDropServer(ctx *DropServerContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseMySQLParserListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseMySQLParserListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropTableSpace is called when production dropTableSpace is entered.
func (s *BaseMySQLParserListener) EnterDropTableSpace(ctx *DropTableSpaceContext) {}

// ExitDropTableSpace is called when production dropTableSpace is exited.
func (s *BaseMySQLParserListener) ExitDropTableSpace(ctx *DropTableSpaceContext) {}

// EnterDropTrigger is called when production dropTrigger is entered.
func (s *BaseMySQLParserListener) EnterDropTrigger(ctx *DropTriggerContext) {}

// ExitDropTrigger is called when production dropTrigger is exited.
func (s *BaseMySQLParserListener) ExitDropTrigger(ctx *DropTriggerContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseMySQLParserListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseMySQLParserListener) ExitDropView(ctx *DropViewContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseMySQLParserListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseMySQLParserListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterDropSpatialReference is called when production dropSpatialReference is entered.
func (s *BaseMySQLParserListener) EnterDropSpatialReference(ctx *DropSpatialReferenceContext) {}

// ExitDropSpatialReference is called when production dropSpatialReference is exited.
func (s *BaseMySQLParserListener) ExitDropSpatialReference(ctx *DropSpatialReferenceContext) {}

// EnterDropUndoTablespace is called when production dropUndoTablespace is entered.
func (s *BaseMySQLParserListener) EnterDropUndoTablespace(ctx *DropUndoTablespaceContext) {}

// ExitDropUndoTablespace is called when production dropUndoTablespace is exited.
func (s *BaseMySQLParserListener) ExitDropUndoTablespace(ctx *DropUndoTablespaceContext) {}

// EnterRenameTableStatement is called when production renameTableStatement is entered.
func (s *BaseMySQLParserListener) EnterRenameTableStatement(ctx *RenameTableStatementContext) {}

// ExitRenameTableStatement is called when production renameTableStatement is exited.
func (s *BaseMySQLParserListener) ExitRenameTableStatement(ctx *RenameTableStatementContext) {}

// EnterRenamePair is called when production renamePair is entered.
func (s *BaseMySQLParserListener) EnterRenamePair(ctx *RenamePairContext) {}

// ExitRenamePair is called when production renamePair is exited.
func (s *BaseMySQLParserListener) ExitRenamePair(ctx *RenamePairContext) {}

// EnterTruncateTableStatement is called when production truncateTableStatement is entered.
func (s *BaseMySQLParserListener) EnterTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// ExitTruncateTableStatement is called when production truncateTableStatement is exited.
func (s *BaseMySQLParserListener) ExitTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseMySQLParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseMySQLParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterCallStatement is called when production callStatement is entered.
func (s *BaseMySQLParserListener) EnterCallStatement(ctx *CallStatementContext) {}

// ExitCallStatement is called when production callStatement is exited.
func (s *BaseMySQLParserListener) ExitCallStatement(ctx *CallStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseMySQLParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseMySQLParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterPartitionDelete is called when production partitionDelete is entered.
func (s *BaseMySQLParserListener) EnterPartitionDelete(ctx *PartitionDeleteContext) {}

// ExitPartitionDelete is called when production partitionDelete is exited.
func (s *BaseMySQLParserListener) ExitPartitionDelete(ctx *PartitionDeleteContext) {}

// EnterDeleteStatementOption is called when production deleteStatementOption is entered.
func (s *BaseMySQLParserListener) EnterDeleteStatementOption(ctx *DeleteStatementOptionContext) {}

// ExitDeleteStatementOption is called when production deleteStatementOption is exited.
func (s *BaseMySQLParserListener) ExitDeleteStatementOption(ctx *DeleteStatementOptionContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseMySQLParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseMySQLParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterHandlerStatement is called when production handlerStatement is entered.
func (s *BaseMySQLParserListener) EnterHandlerStatement(ctx *HandlerStatementContext) {}

// ExitHandlerStatement is called when production handlerStatement is exited.
func (s *BaseMySQLParserListener) ExitHandlerStatement(ctx *HandlerStatementContext) {}

// EnterHandlerReadOrScan is called when production handlerReadOrScan is entered.
func (s *BaseMySQLParserListener) EnterHandlerReadOrScan(ctx *HandlerReadOrScanContext) {}

// ExitHandlerReadOrScan is called when production handlerReadOrScan is exited.
func (s *BaseMySQLParserListener) ExitHandlerReadOrScan(ctx *HandlerReadOrScanContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseMySQLParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseMySQLParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterInsertLockOption is called when production insertLockOption is entered.
func (s *BaseMySQLParserListener) EnterInsertLockOption(ctx *InsertLockOptionContext) {}

// ExitInsertLockOption is called when production insertLockOption is exited.
func (s *BaseMySQLParserListener) ExitInsertLockOption(ctx *InsertLockOptionContext) {}

// EnterInsertFromConstructor is called when production insertFromConstructor is entered.
func (s *BaseMySQLParserListener) EnterInsertFromConstructor(ctx *InsertFromConstructorContext) {}

// ExitInsertFromConstructor is called when production insertFromConstructor is exited.
func (s *BaseMySQLParserListener) ExitInsertFromConstructor(ctx *InsertFromConstructorContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseMySQLParserListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseMySQLParserListener) ExitFields(ctx *FieldsContext) {}

// EnterInsertValues is called when production insertValues is entered.
func (s *BaseMySQLParserListener) EnterInsertValues(ctx *InsertValuesContext) {}

// ExitInsertValues is called when production insertValues is exited.
func (s *BaseMySQLParserListener) ExitInsertValues(ctx *InsertValuesContext) {}

// EnterInsertQueryExpression is called when production insertQueryExpression is entered.
func (s *BaseMySQLParserListener) EnterInsertQueryExpression(ctx *InsertQueryExpressionContext) {}

// ExitInsertQueryExpression is called when production insertQueryExpression is exited.
func (s *BaseMySQLParserListener) ExitInsertQueryExpression(ctx *InsertQueryExpressionContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseMySQLParserListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseMySQLParserListener) ExitValueList(ctx *ValueListContext) {}

// EnterValues is called when production values is entered.
func (s *BaseMySQLParserListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseMySQLParserListener) ExitValues(ctx *ValuesContext) {}

// EnterValuesReference is called when production valuesReference is entered.
func (s *BaseMySQLParserListener) EnterValuesReference(ctx *ValuesReferenceContext) {}

// ExitValuesReference is called when production valuesReference is exited.
func (s *BaseMySQLParserListener) ExitValuesReference(ctx *ValuesReferenceContext) {}

// EnterInsertUpdateList is called when production insertUpdateList is entered.
func (s *BaseMySQLParserListener) EnterInsertUpdateList(ctx *InsertUpdateListContext) {}

// ExitInsertUpdateList is called when production insertUpdateList is exited.
func (s *BaseMySQLParserListener) ExitInsertUpdateList(ctx *InsertUpdateListContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseMySQLParserListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseMySQLParserListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterDataOrXml is called when production dataOrXml is entered.
func (s *BaseMySQLParserListener) EnterDataOrXml(ctx *DataOrXmlContext) {}

// ExitDataOrXml is called when production dataOrXml is exited.
func (s *BaseMySQLParserListener) ExitDataOrXml(ctx *DataOrXmlContext) {}

// EnterXmlRowsIdentifiedBy is called when production xmlRowsIdentifiedBy is entered.
func (s *BaseMySQLParserListener) EnterXmlRowsIdentifiedBy(ctx *XmlRowsIdentifiedByContext) {}

// ExitXmlRowsIdentifiedBy is called when production xmlRowsIdentifiedBy is exited.
func (s *BaseMySQLParserListener) ExitXmlRowsIdentifiedBy(ctx *XmlRowsIdentifiedByContext) {}

// EnterLoadDataFileTail is called when production loadDataFileTail is entered.
func (s *BaseMySQLParserListener) EnterLoadDataFileTail(ctx *LoadDataFileTailContext) {}

// ExitLoadDataFileTail is called when production loadDataFileTail is exited.
func (s *BaseMySQLParserListener) ExitLoadDataFileTail(ctx *LoadDataFileTailContext) {}

// EnterLoadDataFileTargetList is called when production loadDataFileTargetList is entered.
func (s *BaseMySQLParserListener) EnterLoadDataFileTargetList(ctx *LoadDataFileTargetListContext) {}

// ExitLoadDataFileTargetList is called when production loadDataFileTargetList is exited.
func (s *BaseMySQLParserListener) ExitLoadDataFileTargetList(ctx *LoadDataFileTargetListContext) {}

// EnterFieldOrVariableList is called when production fieldOrVariableList is entered.
func (s *BaseMySQLParserListener) EnterFieldOrVariableList(ctx *FieldOrVariableListContext) {}

// ExitFieldOrVariableList is called when production fieldOrVariableList is exited.
func (s *BaseMySQLParserListener) ExitFieldOrVariableList(ctx *FieldOrVariableListContext) {}

// EnterReplaceStatement is called when production replaceStatement is entered.
func (s *BaseMySQLParserListener) EnterReplaceStatement(ctx *ReplaceStatementContext) {}

// ExitReplaceStatement is called when production replaceStatement is exited.
func (s *BaseMySQLParserListener) ExitReplaceStatement(ctx *ReplaceStatementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseMySQLParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseMySQLParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectStatementWithInto is called when production selectStatementWithInto is entered.
func (s *BaseMySQLParserListener) EnterSelectStatementWithInto(ctx *SelectStatementWithIntoContext) {}

// ExitSelectStatementWithInto is called when production selectStatementWithInto is exited.
func (s *BaseMySQLParserListener) ExitSelectStatementWithInto(ctx *SelectStatementWithIntoContext) {}

// EnterQueryExpression is called when production queryExpression is entered.
func (s *BaseMySQLParserListener) EnterQueryExpression(ctx *QueryExpressionContext) {}

// ExitQueryExpression is called when production queryExpression is exited.
func (s *BaseMySQLParserListener) ExitQueryExpression(ctx *QueryExpressionContext) {}

// EnterQueryExpressionBody is called when production queryExpressionBody is entered.
func (s *BaseMySQLParserListener) EnterQueryExpressionBody(ctx *QueryExpressionBodyContext) {}

// ExitQueryExpressionBody is called when production queryExpressionBody is exited.
func (s *BaseMySQLParserListener) ExitQueryExpressionBody(ctx *QueryExpressionBodyContext) {}

// EnterQueryExpressionParens is called when production queryExpressionParens is entered.
func (s *BaseMySQLParserListener) EnterQueryExpressionParens(ctx *QueryExpressionParensContext) {}

// ExitQueryExpressionParens is called when production queryExpressionParens is exited.
func (s *BaseMySQLParserListener) ExitQueryExpressionParens(ctx *QueryExpressionParensContext) {}

// EnterQueryPrimary is called when production queryPrimary is entered.
func (s *BaseMySQLParserListener) EnterQueryPrimary(ctx *QueryPrimaryContext) {}

// ExitQueryPrimary is called when production queryPrimary is exited.
func (s *BaseMySQLParserListener) ExitQueryPrimary(ctx *QueryPrimaryContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseMySQLParserListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseMySQLParserListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseMySQLParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseMySQLParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterQuerySpecOption is called when production querySpecOption is entered.
func (s *BaseMySQLParserListener) EnterQuerySpecOption(ctx *QuerySpecOptionContext) {}

// ExitQuerySpecOption is called when production querySpecOption is exited.
func (s *BaseMySQLParserListener) ExitQuerySpecOption(ctx *QuerySpecOptionContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseMySQLParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseMySQLParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterSimpleLimitClause is called when production simpleLimitClause is entered.
func (s *BaseMySQLParserListener) EnterSimpleLimitClause(ctx *SimpleLimitClauseContext) {}

// ExitSimpleLimitClause is called when production simpleLimitClause is exited.
func (s *BaseMySQLParserListener) ExitSimpleLimitClause(ctx *SimpleLimitClauseContext) {}

// EnterLimitOptions is called when production limitOptions is entered.
func (s *BaseMySQLParserListener) EnterLimitOptions(ctx *LimitOptionsContext) {}

// ExitLimitOptions is called when production limitOptions is exited.
func (s *BaseMySQLParserListener) ExitLimitOptions(ctx *LimitOptionsContext) {}

// EnterLimitOption is called when production limitOption is entered.
func (s *BaseMySQLParserListener) EnterLimitOption(ctx *LimitOptionContext) {}

// ExitLimitOption is called when production limitOption is exited.
func (s *BaseMySQLParserListener) ExitLimitOption(ctx *LimitOptionContext) {}

// EnterIntoClause is called when production intoClause is entered.
func (s *BaseMySQLParserListener) EnterIntoClause(ctx *IntoClauseContext) {}

// ExitIntoClause is called when production intoClause is exited.
func (s *BaseMySQLParserListener) ExitIntoClause(ctx *IntoClauseContext) {}

// EnterProcedureAnalyseClause is called when production procedureAnalyseClause is entered.
func (s *BaseMySQLParserListener) EnterProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) {}

// ExitProcedureAnalyseClause is called when production procedureAnalyseClause is exited.
func (s *BaseMySQLParserListener) ExitProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseMySQLParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseMySQLParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseMySQLParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseMySQLParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterWindowDefinition is called when production windowDefinition is entered.
func (s *BaseMySQLParserListener) EnterWindowDefinition(ctx *WindowDefinitionContext) {}

// ExitWindowDefinition is called when production windowDefinition is exited.
func (s *BaseMySQLParserListener) ExitWindowDefinition(ctx *WindowDefinitionContext) {}

// EnterWindowSpec is called when production windowSpec is entered.
func (s *BaseMySQLParserListener) EnterWindowSpec(ctx *WindowSpecContext) {}

// ExitWindowSpec is called when production windowSpec is exited.
func (s *BaseMySQLParserListener) ExitWindowSpec(ctx *WindowSpecContext) {}

// EnterWindowSpecDetails is called when production windowSpecDetails is entered.
func (s *BaseMySQLParserListener) EnterWindowSpecDetails(ctx *WindowSpecDetailsContext) {}

// ExitWindowSpecDetails is called when production windowSpecDetails is exited.
func (s *BaseMySQLParserListener) ExitWindowSpecDetails(ctx *WindowSpecDetailsContext) {}

// EnterWindowFrameClause is called when production windowFrameClause is entered.
func (s *BaseMySQLParserListener) EnterWindowFrameClause(ctx *WindowFrameClauseContext) {}

// ExitWindowFrameClause is called when production windowFrameClause is exited.
func (s *BaseMySQLParserListener) ExitWindowFrameClause(ctx *WindowFrameClauseContext) {}

// EnterWindowFrameUnits is called when production windowFrameUnits is entered.
func (s *BaseMySQLParserListener) EnterWindowFrameUnits(ctx *WindowFrameUnitsContext) {}

// ExitWindowFrameUnits is called when production windowFrameUnits is exited.
func (s *BaseMySQLParserListener) ExitWindowFrameUnits(ctx *WindowFrameUnitsContext) {}

// EnterWindowFrameExtent is called when production windowFrameExtent is entered.
func (s *BaseMySQLParserListener) EnterWindowFrameExtent(ctx *WindowFrameExtentContext) {}

// ExitWindowFrameExtent is called when production windowFrameExtent is exited.
func (s *BaseMySQLParserListener) ExitWindowFrameExtent(ctx *WindowFrameExtentContext) {}

// EnterWindowFrameStart is called when production windowFrameStart is entered.
func (s *BaseMySQLParserListener) EnterWindowFrameStart(ctx *WindowFrameStartContext) {}

// ExitWindowFrameStart is called when production windowFrameStart is exited.
func (s *BaseMySQLParserListener) ExitWindowFrameStart(ctx *WindowFrameStartContext) {}

// EnterWindowFrameBetween is called when production windowFrameBetween is entered.
func (s *BaseMySQLParserListener) EnterWindowFrameBetween(ctx *WindowFrameBetweenContext) {}

// ExitWindowFrameBetween is called when production windowFrameBetween is exited.
func (s *BaseMySQLParserListener) ExitWindowFrameBetween(ctx *WindowFrameBetweenContext) {}

// EnterWindowFrameBound is called when production windowFrameBound is entered.
func (s *BaseMySQLParserListener) EnterWindowFrameBound(ctx *WindowFrameBoundContext) {}

// ExitWindowFrameBound is called when production windowFrameBound is exited.
func (s *BaseMySQLParserListener) ExitWindowFrameBound(ctx *WindowFrameBoundContext) {}

// EnterWindowFrameExclusion is called when production windowFrameExclusion is entered.
func (s *BaseMySQLParserListener) EnterWindowFrameExclusion(ctx *WindowFrameExclusionContext) {}

// ExitWindowFrameExclusion is called when production windowFrameExclusion is exited.
func (s *BaseMySQLParserListener) ExitWindowFrameExclusion(ctx *WindowFrameExclusionContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseMySQLParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseMySQLParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterCommonTableExpression is called when production commonTableExpression is entered.
func (s *BaseMySQLParserListener) EnterCommonTableExpression(ctx *CommonTableExpressionContext) {}

// ExitCommonTableExpression is called when production commonTableExpression is exited.
func (s *BaseMySQLParserListener) ExitCommonTableExpression(ctx *CommonTableExpressionContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseMySQLParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseMySQLParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterOlapOption is called when production olapOption is entered.
func (s *BaseMySQLParserListener) EnterOlapOption(ctx *OlapOptionContext) {}

// ExitOlapOption is called when production olapOption is exited.
func (s *BaseMySQLParserListener) ExitOlapOption(ctx *OlapOptionContext) {}

// EnterOrderClause is called when production orderClause is entered.
func (s *BaseMySQLParserListener) EnterOrderClause(ctx *OrderClauseContext) {}

// ExitOrderClause is called when production orderClause is exited.
func (s *BaseMySQLParserListener) ExitOrderClause(ctx *OrderClauseContext) {}

// EnterDirection is called when production direction is entered.
func (s *BaseMySQLParserListener) EnterDirection(ctx *DirectionContext) {}

// ExitDirection is called when production direction is exited.
func (s *BaseMySQLParserListener) ExitDirection(ctx *DirectionContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseMySQLParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseMySQLParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableReferenceList is called when production tableReferenceList is entered.
func (s *BaseMySQLParserListener) EnterTableReferenceList(ctx *TableReferenceListContext) {}

// ExitTableReferenceList is called when production tableReferenceList is exited.
func (s *BaseMySQLParserListener) ExitTableReferenceList(ctx *TableReferenceListContext) {}

// EnterTableValueConstructor is called when production tableValueConstructor is entered.
func (s *BaseMySQLParserListener) EnterTableValueConstructor(ctx *TableValueConstructorContext) {}

// ExitTableValueConstructor is called when production tableValueConstructor is exited.
func (s *BaseMySQLParserListener) ExitTableValueConstructor(ctx *TableValueConstructorContext) {}

// EnterExplicitTable is called when production explicitTable is entered.
func (s *BaseMySQLParserListener) EnterExplicitTable(ctx *ExplicitTableContext) {}

// ExitExplicitTable is called when production explicitTable is exited.
func (s *BaseMySQLParserListener) ExitExplicitTable(ctx *ExplicitTableContext) {}

// EnterRowValueExplicit is called when production rowValueExplicit is entered.
func (s *BaseMySQLParserListener) EnterRowValueExplicit(ctx *RowValueExplicitContext) {}

// ExitRowValueExplicit is called when production rowValueExplicit is exited.
func (s *BaseMySQLParserListener) ExitRowValueExplicit(ctx *RowValueExplicitContext) {}

// EnterSelectOption is called when production selectOption is entered.
func (s *BaseMySQLParserListener) EnterSelectOption(ctx *SelectOptionContext) {}

// ExitSelectOption is called when production selectOption is exited.
func (s *BaseMySQLParserListener) ExitSelectOption(ctx *SelectOptionContext) {}

// EnterLockingClauseList is called when production lockingClauseList is entered.
func (s *BaseMySQLParserListener) EnterLockingClauseList(ctx *LockingClauseListContext) {}

// ExitLockingClauseList is called when production lockingClauseList is exited.
func (s *BaseMySQLParserListener) ExitLockingClauseList(ctx *LockingClauseListContext) {}

// EnterLockingClause is called when production lockingClause is entered.
func (s *BaseMySQLParserListener) EnterLockingClause(ctx *LockingClauseContext) {}

// ExitLockingClause is called when production lockingClause is exited.
func (s *BaseMySQLParserListener) ExitLockingClause(ctx *LockingClauseContext) {}

// EnterLockStrengh is called when production lockStrengh is entered.
func (s *BaseMySQLParserListener) EnterLockStrengh(ctx *LockStrenghContext) {}

// ExitLockStrengh is called when production lockStrengh is exited.
func (s *BaseMySQLParserListener) ExitLockStrengh(ctx *LockStrenghContext) {}

// EnterLockedRowAction is called when production lockedRowAction is entered.
func (s *BaseMySQLParserListener) EnterLockedRowAction(ctx *LockedRowActionContext) {}

// ExitLockedRowAction is called when production lockedRowAction is exited.
func (s *BaseMySQLParserListener) ExitLockedRowAction(ctx *LockedRowActionContext) {}

// EnterSelectItemList is called when production selectItemList is entered.
func (s *BaseMySQLParserListener) EnterSelectItemList(ctx *SelectItemListContext) {}

// ExitSelectItemList is called when production selectItemList is exited.
func (s *BaseMySQLParserListener) ExitSelectItemList(ctx *SelectItemListContext) {}

// EnterSelectItem is called when production selectItem is entered.
func (s *BaseMySQLParserListener) EnterSelectItem(ctx *SelectItemContext) {}

// ExitSelectItem is called when production selectItem is exited.
func (s *BaseMySQLParserListener) ExitSelectItem(ctx *SelectItemContext) {}

// EnterSelectAlias is called when production selectAlias is entered.
func (s *BaseMySQLParserListener) EnterSelectAlias(ctx *SelectAliasContext) {}

// ExitSelectAlias is called when production selectAlias is exited.
func (s *BaseMySQLParserListener) ExitSelectAlias(ctx *SelectAliasContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseMySQLParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseMySQLParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterTableReference is called when production tableReference is entered.
func (s *BaseMySQLParserListener) EnterTableReference(ctx *TableReferenceContext) {}

// ExitTableReference is called when production tableReference is exited.
func (s *BaseMySQLParserListener) ExitTableReference(ctx *TableReferenceContext) {}

// EnterEscapedTableReference is called when production escapedTableReference is entered.
func (s *BaseMySQLParserListener) EnterEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// ExitEscapedTableReference is called when production escapedTableReference is exited.
func (s *BaseMySQLParserListener) ExitEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// EnterJoinedTable is called when production joinedTable is entered.
func (s *BaseMySQLParserListener) EnterJoinedTable(ctx *JoinedTableContext) {}

// ExitJoinedTable is called when production joinedTable is exited.
func (s *BaseMySQLParserListener) ExitJoinedTable(ctx *JoinedTableContext) {}

// EnterNaturalJoinType is called when production naturalJoinType is entered.
func (s *BaseMySQLParserListener) EnterNaturalJoinType(ctx *NaturalJoinTypeContext) {}

// ExitNaturalJoinType is called when production naturalJoinType is exited.
func (s *BaseMySQLParserListener) ExitNaturalJoinType(ctx *NaturalJoinTypeContext) {}

// EnterInnerJoinType is called when production innerJoinType is entered.
func (s *BaseMySQLParserListener) EnterInnerJoinType(ctx *InnerJoinTypeContext) {}

// ExitInnerJoinType is called when production innerJoinType is exited.
func (s *BaseMySQLParserListener) ExitInnerJoinType(ctx *InnerJoinTypeContext) {}

// EnterOuterJoinType is called when production outerJoinType is entered.
func (s *BaseMySQLParserListener) EnterOuterJoinType(ctx *OuterJoinTypeContext) {}

// ExitOuterJoinType is called when production outerJoinType is exited.
func (s *BaseMySQLParserListener) ExitOuterJoinType(ctx *OuterJoinTypeContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (s *BaseMySQLParserListener) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (s *BaseMySQLParserListener) ExitTableFactor(ctx *TableFactorContext) {}

// EnterSingleTable is called when production singleTable is entered.
func (s *BaseMySQLParserListener) EnterSingleTable(ctx *SingleTableContext) {}

// ExitSingleTable is called when production singleTable is exited.
func (s *BaseMySQLParserListener) ExitSingleTable(ctx *SingleTableContext) {}

// EnterSingleTableParens is called when production singleTableParens is entered.
func (s *BaseMySQLParserListener) EnterSingleTableParens(ctx *SingleTableParensContext) {}

// ExitSingleTableParens is called when production singleTableParens is exited.
func (s *BaseMySQLParserListener) ExitSingleTableParens(ctx *SingleTableParensContext) {}

// EnterDerivedTable is called when production derivedTable is entered.
func (s *BaseMySQLParserListener) EnterDerivedTable(ctx *DerivedTableContext) {}

// ExitDerivedTable is called when production derivedTable is exited.
func (s *BaseMySQLParserListener) ExitDerivedTable(ctx *DerivedTableContext) {}

// EnterTableReferenceListParens is called when production tableReferenceListParens is entered.
func (s *BaseMySQLParserListener) EnterTableReferenceListParens(ctx *TableReferenceListParensContext) {
}

// ExitTableReferenceListParens is called when production tableReferenceListParens is exited.
func (s *BaseMySQLParserListener) ExitTableReferenceListParens(ctx *TableReferenceListParensContext) {
}

// EnterTableFunction is called when production tableFunction is entered.
func (s *BaseMySQLParserListener) EnterTableFunction(ctx *TableFunctionContext) {}

// ExitTableFunction is called when production tableFunction is exited.
func (s *BaseMySQLParserListener) ExitTableFunction(ctx *TableFunctionContext) {}

// EnterColumnsClause is called when production columnsClause is entered.
func (s *BaseMySQLParserListener) EnterColumnsClause(ctx *ColumnsClauseContext) {}

// ExitColumnsClause is called when production columnsClause is exited.
func (s *BaseMySQLParserListener) ExitColumnsClause(ctx *ColumnsClauseContext) {}

// EnterJtColumn is called when production jtColumn is entered.
func (s *BaseMySQLParserListener) EnterJtColumn(ctx *JtColumnContext) {}

// ExitJtColumn is called when production jtColumn is exited.
func (s *BaseMySQLParserListener) ExitJtColumn(ctx *JtColumnContext) {}

// EnterOnEmptyOrError is called when production onEmptyOrError is entered.
func (s *BaseMySQLParserListener) EnterOnEmptyOrError(ctx *OnEmptyOrErrorContext) {}

// ExitOnEmptyOrError is called when production onEmptyOrError is exited.
func (s *BaseMySQLParserListener) ExitOnEmptyOrError(ctx *OnEmptyOrErrorContext) {}

// EnterOnEmpty is called when production onEmpty is entered.
func (s *BaseMySQLParserListener) EnterOnEmpty(ctx *OnEmptyContext) {}

// ExitOnEmpty is called when production onEmpty is exited.
func (s *BaseMySQLParserListener) ExitOnEmpty(ctx *OnEmptyContext) {}

// EnterOnError is called when production onError is entered.
func (s *BaseMySQLParserListener) EnterOnError(ctx *OnErrorContext) {}

// ExitOnError is called when production onError is exited.
func (s *BaseMySQLParserListener) ExitOnError(ctx *OnErrorContext) {}

// EnterJtOnResponse is called when production jtOnResponse is entered.
func (s *BaseMySQLParserListener) EnterJtOnResponse(ctx *JtOnResponseContext) {}

// ExitJtOnResponse is called when production jtOnResponse is exited.
func (s *BaseMySQLParserListener) ExitJtOnResponse(ctx *JtOnResponseContext) {}

// EnterSetOprSymbol is called when production setOprSymbol is entered.
func (s *BaseMySQLParserListener) EnterSetOprSymbol(ctx *SetOprSymbolContext) {}

// ExitSetOprSymbol is called when production setOprSymbol is exited.
func (s *BaseMySQLParserListener) ExitSetOprSymbol(ctx *SetOprSymbolContext) {}

// EnterSetOprOption is called when production setOprOption is entered.
func (s *BaseMySQLParserListener) EnterSetOprOption(ctx *SetOprOptionContext) {}

// ExitSetOprOption is called when production setOprOption is exited.
func (s *BaseMySQLParserListener) ExitSetOprOption(ctx *SetOprOptionContext) {}

// EnterTableAlias is called when production tableAlias is entered.
func (s *BaseMySQLParserListener) EnterTableAlias(ctx *TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *BaseMySQLParserListener) ExitTableAlias(ctx *TableAliasContext) {}

// EnterIndexHintList is called when production indexHintList is entered.
func (s *BaseMySQLParserListener) EnterIndexHintList(ctx *IndexHintListContext) {}

// ExitIndexHintList is called when production indexHintList is exited.
func (s *BaseMySQLParserListener) ExitIndexHintList(ctx *IndexHintListContext) {}

// EnterIndexHint is called when production indexHint is entered.
func (s *BaseMySQLParserListener) EnterIndexHint(ctx *IndexHintContext) {}

// ExitIndexHint is called when production indexHint is exited.
func (s *BaseMySQLParserListener) ExitIndexHint(ctx *IndexHintContext) {}

// EnterIndexHintType is called when production indexHintType is entered.
func (s *BaseMySQLParserListener) EnterIndexHintType(ctx *IndexHintTypeContext) {}

// ExitIndexHintType is called when production indexHintType is exited.
func (s *BaseMySQLParserListener) ExitIndexHintType(ctx *IndexHintTypeContext) {}

// EnterKeyOrIndex is called when production keyOrIndex is entered.
func (s *BaseMySQLParserListener) EnterKeyOrIndex(ctx *KeyOrIndexContext) {}

// ExitKeyOrIndex is called when production keyOrIndex is exited.
func (s *BaseMySQLParserListener) ExitKeyOrIndex(ctx *KeyOrIndexContext) {}

// EnterConstraintKeyType is called when production constraintKeyType is entered.
func (s *BaseMySQLParserListener) EnterConstraintKeyType(ctx *ConstraintKeyTypeContext) {}

// ExitConstraintKeyType is called when production constraintKeyType is exited.
func (s *BaseMySQLParserListener) ExitConstraintKeyType(ctx *ConstraintKeyTypeContext) {}

// EnterIndexHintClause is called when production indexHintClause is entered.
func (s *BaseMySQLParserListener) EnterIndexHintClause(ctx *IndexHintClauseContext) {}

// ExitIndexHintClause is called when production indexHintClause is exited.
func (s *BaseMySQLParserListener) ExitIndexHintClause(ctx *IndexHintClauseContext) {}

// EnterIndexList is called when production indexList is entered.
func (s *BaseMySQLParserListener) EnterIndexList(ctx *IndexListContext) {}

// ExitIndexList is called when production indexList is exited.
func (s *BaseMySQLParserListener) ExitIndexList(ctx *IndexListContext) {}

// EnterIndexListElement is called when production indexListElement is entered.
func (s *BaseMySQLParserListener) EnterIndexListElement(ctx *IndexListElementContext) {}

// ExitIndexListElement is called when production indexListElement is exited.
func (s *BaseMySQLParserListener) ExitIndexListElement(ctx *IndexListElementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseMySQLParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseMySQLParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterTransactionOrLockingStatement is called when production transactionOrLockingStatement is entered.
func (s *BaseMySQLParserListener) EnterTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) {
}

// ExitTransactionOrLockingStatement is called when production transactionOrLockingStatement is exited.
func (s *BaseMySQLParserListener) ExitTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) {
}

// EnterTransactionStatement is called when production transactionStatement is entered.
func (s *BaseMySQLParserListener) EnterTransactionStatement(ctx *TransactionStatementContext) {}

// ExitTransactionStatement is called when production transactionStatement is exited.
func (s *BaseMySQLParserListener) ExitTransactionStatement(ctx *TransactionStatementContext) {}

// EnterBeginWork is called when production beginWork is entered.
func (s *BaseMySQLParserListener) EnterBeginWork(ctx *BeginWorkContext) {}

// ExitBeginWork is called when production beginWork is exited.
func (s *BaseMySQLParserListener) ExitBeginWork(ctx *BeginWorkContext) {}

// EnterTransactionCharacteristic is called when production transactionCharacteristic is entered.
func (s *BaseMySQLParserListener) EnterTransactionCharacteristic(ctx *TransactionCharacteristicContext) {
}

// ExitTransactionCharacteristic is called when production transactionCharacteristic is exited.
func (s *BaseMySQLParserListener) ExitTransactionCharacteristic(ctx *TransactionCharacteristicContext) {
}

// EnterSavepointStatement is called when production savepointStatement is entered.
func (s *BaseMySQLParserListener) EnterSavepointStatement(ctx *SavepointStatementContext) {}

// ExitSavepointStatement is called when production savepointStatement is exited.
func (s *BaseMySQLParserListener) ExitSavepointStatement(ctx *SavepointStatementContext) {}

// EnterLockStatement is called when production lockStatement is entered.
func (s *BaseMySQLParserListener) EnterLockStatement(ctx *LockStatementContext) {}

// ExitLockStatement is called when production lockStatement is exited.
func (s *BaseMySQLParserListener) ExitLockStatement(ctx *LockStatementContext) {}

// EnterLockItem is called when production lockItem is entered.
func (s *BaseMySQLParserListener) EnterLockItem(ctx *LockItemContext) {}

// ExitLockItem is called when production lockItem is exited.
func (s *BaseMySQLParserListener) ExitLockItem(ctx *LockItemContext) {}

// EnterLockOption is called when production lockOption is entered.
func (s *BaseMySQLParserListener) EnterLockOption(ctx *LockOptionContext) {}

// ExitLockOption is called when production lockOption is exited.
func (s *BaseMySQLParserListener) ExitLockOption(ctx *LockOptionContext) {}

// EnterXaStatement is called when production xaStatement is entered.
func (s *BaseMySQLParserListener) EnterXaStatement(ctx *XaStatementContext) {}

// ExitXaStatement is called when production xaStatement is exited.
func (s *BaseMySQLParserListener) ExitXaStatement(ctx *XaStatementContext) {}

// EnterXaConvert is called when production xaConvert is entered.
func (s *BaseMySQLParserListener) EnterXaConvert(ctx *XaConvertContext) {}

// ExitXaConvert is called when production xaConvert is exited.
func (s *BaseMySQLParserListener) ExitXaConvert(ctx *XaConvertContext) {}

// EnterXid is called when production xid is entered.
func (s *BaseMySQLParserListener) EnterXid(ctx *XidContext) {}

// ExitXid is called when production xid is exited.
func (s *BaseMySQLParserListener) ExitXid(ctx *XidContext) {}

// EnterReplicationStatement is called when production replicationStatement is entered.
func (s *BaseMySQLParserListener) EnterReplicationStatement(ctx *ReplicationStatementContext) {}

// ExitReplicationStatement is called when production replicationStatement is exited.
func (s *BaseMySQLParserListener) ExitReplicationStatement(ctx *ReplicationStatementContext) {}

// EnterResetOption is called when production resetOption is entered.
func (s *BaseMySQLParserListener) EnterResetOption(ctx *ResetOptionContext) {}

// ExitResetOption is called when production resetOption is exited.
func (s *BaseMySQLParserListener) ExitResetOption(ctx *ResetOptionContext) {}

// EnterMasterResetOptions is called when production masterResetOptions is entered.
func (s *BaseMySQLParserListener) EnterMasterResetOptions(ctx *MasterResetOptionsContext) {}

// ExitMasterResetOptions is called when production masterResetOptions is exited.
func (s *BaseMySQLParserListener) ExitMasterResetOptions(ctx *MasterResetOptionsContext) {}

// EnterReplicationLoad is called when production replicationLoad is entered.
func (s *BaseMySQLParserListener) EnterReplicationLoad(ctx *ReplicationLoadContext) {}

// ExitReplicationLoad is called when production replicationLoad is exited.
func (s *BaseMySQLParserListener) ExitReplicationLoad(ctx *ReplicationLoadContext) {}

// EnterChangeMaster is called when production changeMaster is entered.
func (s *BaseMySQLParserListener) EnterChangeMaster(ctx *ChangeMasterContext) {}

// ExitChangeMaster is called when production changeMaster is exited.
func (s *BaseMySQLParserListener) ExitChangeMaster(ctx *ChangeMasterContext) {}

// EnterChangeMasterOptions is called when production changeMasterOptions is entered.
func (s *BaseMySQLParserListener) EnterChangeMasterOptions(ctx *ChangeMasterOptionsContext) {}

// ExitChangeMasterOptions is called when production changeMasterOptions is exited.
func (s *BaseMySQLParserListener) ExitChangeMasterOptions(ctx *ChangeMasterOptionsContext) {}

// EnterMasterOption is called when production masterOption is entered.
func (s *BaseMySQLParserListener) EnterMasterOption(ctx *MasterOptionContext) {}

// ExitMasterOption is called when production masterOption is exited.
func (s *BaseMySQLParserListener) ExitMasterOption(ctx *MasterOptionContext) {}

// EnterPrivilegeCheckDef is called when production privilegeCheckDef is entered.
func (s *BaseMySQLParserListener) EnterPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) {}

// ExitPrivilegeCheckDef is called when production privilegeCheckDef is exited.
func (s *BaseMySQLParserListener) ExitPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) {}

// EnterTablePrimaryKeyCheckDef is called when production tablePrimaryKeyCheckDef is entered.
func (s *BaseMySQLParserListener) EnterTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) {}

// ExitTablePrimaryKeyCheckDef is called when production tablePrimaryKeyCheckDef is exited.
func (s *BaseMySQLParserListener) ExitTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) {}

// EnterMasterTlsCiphersuitesDef is called when production masterTlsCiphersuitesDef is entered.
func (s *BaseMySQLParserListener) EnterMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) {
}

// ExitMasterTlsCiphersuitesDef is called when production masterTlsCiphersuitesDef is exited.
func (s *BaseMySQLParserListener) ExitMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) {
}

// EnterMasterFileDef is called when production masterFileDef is entered.
func (s *BaseMySQLParserListener) EnterMasterFileDef(ctx *MasterFileDefContext) {}

// ExitMasterFileDef is called when production masterFileDef is exited.
func (s *BaseMySQLParserListener) ExitMasterFileDef(ctx *MasterFileDefContext) {}

// EnterServerIdList is called when production serverIdList is entered.
func (s *BaseMySQLParserListener) EnterServerIdList(ctx *ServerIdListContext) {}

// ExitServerIdList is called when production serverIdList is exited.
func (s *BaseMySQLParserListener) ExitServerIdList(ctx *ServerIdListContext) {}

// EnterChangeReplication is called when production changeReplication is entered.
func (s *BaseMySQLParserListener) EnterChangeReplication(ctx *ChangeReplicationContext) {}

// ExitChangeReplication is called when production changeReplication is exited.
func (s *BaseMySQLParserListener) ExitChangeReplication(ctx *ChangeReplicationContext) {}

// EnterFilterDefinition is called when production filterDefinition is entered.
func (s *BaseMySQLParserListener) EnterFilterDefinition(ctx *FilterDefinitionContext) {}

// ExitFilterDefinition is called when production filterDefinition is exited.
func (s *BaseMySQLParserListener) ExitFilterDefinition(ctx *FilterDefinitionContext) {}

// EnterFilterDbList is called when production filterDbList is entered.
func (s *BaseMySQLParserListener) EnterFilterDbList(ctx *FilterDbListContext) {}

// ExitFilterDbList is called when production filterDbList is exited.
func (s *BaseMySQLParserListener) ExitFilterDbList(ctx *FilterDbListContext) {}

// EnterFilterTableList is called when production filterTableList is entered.
func (s *BaseMySQLParserListener) EnterFilterTableList(ctx *FilterTableListContext) {}

// ExitFilterTableList is called when production filterTableList is exited.
func (s *BaseMySQLParserListener) ExitFilterTableList(ctx *FilterTableListContext) {}

// EnterFilterStringList is called when production filterStringList is entered.
func (s *BaseMySQLParserListener) EnterFilterStringList(ctx *FilterStringListContext) {}

// ExitFilterStringList is called when production filterStringList is exited.
func (s *BaseMySQLParserListener) ExitFilterStringList(ctx *FilterStringListContext) {}

// EnterFilterWildDbTableString is called when production filterWildDbTableString is entered.
func (s *BaseMySQLParserListener) EnterFilterWildDbTableString(ctx *FilterWildDbTableStringContext) {}

// ExitFilterWildDbTableString is called when production filterWildDbTableString is exited.
func (s *BaseMySQLParserListener) ExitFilterWildDbTableString(ctx *FilterWildDbTableStringContext) {}

// EnterFilterDbPairList is called when production filterDbPairList is entered.
func (s *BaseMySQLParserListener) EnterFilterDbPairList(ctx *FilterDbPairListContext) {}

// ExitFilterDbPairList is called when production filterDbPairList is exited.
func (s *BaseMySQLParserListener) ExitFilterDbPairList(ctx *FilterDbPairListContext) {}

// EnterSlave is called when production slave is entered.
func (s *BaseMySQLParserListener) EnterSlave(ctx *SlaveContext) {}

// ExitSlave is called when production slave is exited.
func (s *BaseMySQLParserListener) ExitSlave(ctx *SlaveContext) {}

// EnterSlaveUntilOptions is called when production slaveUntilOptions is entered.
func (s *BaseMySQLParserListener) EnterSlaveUntilOptions(ctx *SlaveUntilOptionsContext) {}

// ExitSlaveUntilOptions is called when production slaveUntilOptions is exited.
func (s *BaseMySQLParserListener) ExitSlaveUntilOptions(ctx *SlaveUntilOptionsContext) {}

// EnterSlaveConnectionOptions is called when production slaveConnectionOptions is entered.
func (s *BaseMySQLParserListener) EnterSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) {}

// ExitSlaveConnectionOptions is called when production slaveConnectionOptions is exited.
func (s *BaseMySQLParserListener) ExitSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) {}

// EnterSlaveThreadOptions is called when production slaveThreadOptions is entered.
func (s *BaseMySQLParserListener) EnterSlaveThreadOptions(ctx *SlaveThreadOptionsContext) {}

// ExitSlaveThreadOptions is called when production slaveThreadOptions is exited.
func (s *BaseMySQLParserListener) ExitSlaveThreadOptions(ctx *SlaveThreadOptionsContext) {}

// EnterSlaveThreadOption is called when production slaveThreadOption is entered.
func (s *BaseMySQLParserListener) EnterSlaveThreadOption(ctx *SlaveThreadOptionContext) {}

// ExitSlaveThreadOption is called when production slaveThreadOption is exited.
func (s *BaseMySQLParserListener) ExitSlaveThreadOption(ctx *SlaveThreadOptionContext) {}

// EnterGroupReplication is called when production groupReplication is entered.
func (s *BaseMySQLParserListener) EnterGroupReplication(ctx *GroupReplicationContext) {}

// ExitGroupReplication is called when production groupReplication is exited.
func (s *BaseMySQLParserListener) ExitGroupReplication(ctx *GroupReplicationContext) {}

// EnterPreparedStatement is called when production preparedStatement is entered.
func (s *BaseMySQLParserListener) EnterPreparedStatement(ctx *PreparedStatementContext) {}

// ExitPreparedStatement is called when production preparedStatement is exited.
func (s *BaseMySQLParserListener) ExitPreparedStatement(ctx *PreparedStatementContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseMySQLParserListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseMySQLParserListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterExecuteVarList is called when production executeVarList is entered.
func (s *BaseMySQLParserListener) EnterExecuteVarList(ctx *ExecuteVarListContext) {}

// ExitExecuteVarList is called when production executeVarList is exited.
func (s *BaseMySQLParserListener) ExitExecuteVarList(ctx *ExecuteVarListContext) {}

// EnterCloneStatement is called when production cloneStatement is entered.
func (s *BaseMySQLParserListener) EnterCloneStatement(ctx *CloneStatementContext) {}

// ExitCloneStatement is called when production cloneStatement is exited.
func (s *BaseMySQLParserListener) ExitCloneStatement(ctx *CloneStatementContext) {}

// EnterDataDirSSL is called when production dataDirSSL is entered.
func (s *BaseMySQLParserListener) EnterDataDirSSL(ctx *DataDirSSLContext) {}

// ExitDataDirSSL is called when production dataDirSSL is exited.
func (s *BaseMySQLParserListener) ExitDataDirSSL(ctx *DataDirSSLContext) {}

// EnterSsl is called when production ssl is entered.
func (s *BaseMySQLParserListener) EnterSsl(ctx *SslContext) {}

// ExitSsl is called when production ssl is exited.
func (s *BaseMySQLParserListener) ExitSsl(ctx *SslContext) {}

// EnterAccountManagementStatement is called when production accountManagementStatement is entered.
func (s *BaseMySQLParserListener) EnterAccountManagementStatement(ctx *AccountManagementStatementContext) {
}

// ExitAccountManagementStatement is called when production accountManagementStatement is exited.
func (s *BaseMySQLParserListener) ExitAccountManagementStatement(ctx *AccountManagementStatementContext) {
}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseMySQLParserListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseMySQLParserListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterAlterUserTail is called when production alterUserTail is entered.
func (s *BaseMySQLParserListener) EnterAlterUserTail(ctx *AlterUserTailContext) {}

// ExitAlterUserTail is called when production alterUserTail is exited.
func (s *BaseMySQLParserListener) ExitAlterUserTail(ctx *AlterUserTailContext) {}

// EnterUserFunction is called when production userFunction is entered.
func (s *BaseMySQLParserListener) EnterUserFunction(ctx *UserFunctionContext) {}

// ExitUserFunction is called when production userFunction is exited.
func (s *BaseMySQLParserListener) ExitUserFunction(ctx *UserFunctionContext) {}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseMySQLParserListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseMySQLParserListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterCreateUserTail is called when production createUserTail is entered.
func (s *BaseMySQLParserListener) EnterCreateUserTail(ctx *CreateUserTailContext) {}

// ExitCreateUserTail is called when production createUserTail is exited.
func (s *BaseMySQLParserListener) ExitCreateUserTail(ctx *CreateUserTailContext) {}

// EnterDefaultRoleClause is called when production defaultRoleClause is entered.
func (s *BaseMySQLParserListener) EnterDefaultRoleClause(ctx *DefaultRoleClauseContext) {}

// ExitDefaultRoleClause is called when production defaultRoleClause is exited.
func (s *BaseMySQLParserListener) ExitDefaultRoleClause(ctx *DefaultRoleClauseContext) {}

// EnterRequireClause is called when production requireClause is entered.
func (s *BaseMySQLParserListener) EnterRequireClause(ctx *RequireClauseContext) {}

// ExitRequireClause is called when production requireClause is exited.
func (s *BaseMySQLParserListener) ExitRequireClause(ctx *RequireClauseContext) {}

// EnterConnectOptions is called when production connectOptions is entered.
func (s *BaseMySQLParserListener) EnterConnectOptions(ctx *ConnectOptionsContext) {}

// ExitConnectOptions is called when production connectOptions is exited.
func (s *BaseMySQLParserListener) ExitConnectOptions(ctx *ConnectOptionsContext) {}

// EnterAccountLockPasswordExpireOptions is called when production accountLockPasswordExpireOptions is entered.
func (s *BaseMySQLParserListener) EnterAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) {
}

// ExitAccountLockPasswordExpireOptions is called when production accountLockPasswordExpireOptions is exited.
func (s *BaseMySQLParserListener) ExitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) {
}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseMySQLParserListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseMySQLParserListener) ExitDropUser(ctx *DropUserContext) {}

// EnterGrant is called when production grant is entered.
func (s *BaseMySQLParserListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseMySQLParserListener) ExitGrant(ctx *GrantContext) {}

// EnterGrantTargetList is called when production grantTargetList is entered.
func (s *BaseMySQLParserListener) EnterGrantTargetList(ctx *GrantTargetListContext) {}

// ExitGrantTargetList is called when production grantTargetList is exited.
func (s *BaseMySQLParserListener) ExitGrantTargetList(ctx *GrantTargetListContext) {}

// EnterGrantOptions is called when production grantOptions is entered.
func (s *BaseMySQLParserListener) EnterGrantOptions(ctx *GrantOptionsContext) {}

// ExitGrantOptions is called when production grantOptions is exited.
func (s *BaseMySQLParserListener) ExitGrantOptions(ctx *GrantOptionsContext) {}

// EnterExceptRoleList is called when production exceptRoleList is entered.
func (s *BaseMySQLParserListener) EnterExceptRoleList(ctx *ExceptRoleListContext) {}

// ExitExceptRoleList is called when production exceptRoleList is exited.
func (s *BaseMySQLParserListener) ExitExceptRoleList(ctx *ExceptRoleListContext) {}

// EnterWithRoles is called when production withRoles is entered.
func (s *BaseMySQLParserListener) EnterWithRoles(ctx *WithRolesContext) {}

// ExitWithRoles is called when production withRoles is exited.
func (s *BaseMySQLParserListener) ExitWithRoles(ctx *WithRolesContext) {}

// EnterGrantAs is called when production grantAs is entered.
func (s *BaseMySQLParserListener) EnterGrantAs(ctx *GrantAsContext) {}

// ExitGrantAs is called when production grantAs is exited.
func (s *BaseMySQLParserListener) ExitGrantAs(ctx *GrantAsContext) {}

// EnterVersionedRequireClause is called when production versionedRequireClause is entered.
func (s *BaseMySQLParserListener) EnterVersionedRequireClause(ctx *VersionedRequireClauseContext) {}

// ExitVersionedRequireClause is called when production versionedRequireClause is exited.
func (s *BaseMySQLParserListener) ExitVersionedRequireClause(ctx *VersionedRequireClauseContext) {}

// EnterRenameUser is called when production renameUser is entered.
func (s *BaseMySQLParserListener) EnterRenameUser(ctx *RenameUserContext) {}

// ExitRenameUser is called when production renameUser is exited.
func (s *BaseMySQLParserListener) ExitRenameUser(ctx *RenameUserContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseMySQLParserListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseMySQLParserListener) ExitRevoke(ctx *RevokeContext) {}

// EnterOnTypeTo is called when production onTypeTo is entered.
func (s *BaseMySQLParserListener) EnterOnTypeTo(ctx *OnTypeToContext) {}

// ExitOnTypeTo is called when production onTypeTo is exited.
func (s *BaseMySQLParserListener) ExitOnTypeTo(ctx *OnTypeToContext) {}

// EnterAclType is called when production aclType is entered.
func (s *BaseMySQLParserListener) EnterAclType(ctx *AclTypeContext) {}

// ExitAclType is called when production aclType is exited.
func (s *BaseMySQLParserListener) ExitAclType(ctx *AclTypeContext) {}

// EnterRoleOrPrivilegesList is called when production roleOrPrivilegesList is entered.
func (s *BaseMySQLParserListener) EnterRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) {}

// ExitRoleOrPrivilegesList is called when production roleOrPrivilegesList is exited.
func (s *BaseMySQLParserListener) ExitRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) {}

// EnterRoleOrPrivilege is called when production roleOrPrivilege is entered.
func (s *BaseMySQLParserListener) EnterRoleOrPrivilege(ctx *RoleOrPrivilegeContext) {}

// ExitRoleOrPrivilege is called when production roleOrPrivilege is exited.
func (s *BaseMySQLParserListener) ExitRoleOrPrivilege(ctx *RoleOrPrivilegeContext) {}

// EnterGrantIdentifier is called when production grantIdentifier is entered.
func (s *BaseMySQLParserListener) EnterGrantIdentifier(ctx *GrantIdentifierContext) {}

// ExitGrantIdentifier is called when production grantIdentifier is exited.
func (s *BaseMySQLParserListener) ExitGrantIdentifier(ctx *GrantIdentifierContext) {}

// EnterRequireList is called when production requireList is entered.
func (s *BaseMySQLParserListener) EnterRequireList(ctx *RequireListContext) {}

// ExitRequireList is called when production requireList is exited.
func (s *BaseMySQLParserListener) ExitRequireList(ctx *RequireListContext) {}

// EnterRequireListElement is called when production requireListElement is entered.
func (s *BaseMySQLParserListener) EnterRequireListElement(ctx *RequireListElementContext) {}

// ExitRequireListElement is called when production requireListElement is exited.
func (s *BaseMySQLParserListener) ExitRequireListElement(ctx *RequireListElementContext) {}

// EnterGrantOption is called when production grantOption is entered.
func (s *BaseMySQLParserListener) EnterGrantOption(ctx *GrantOptionContext) {}

// ExitGrantOption is called when production grantOption is exited.
func (s *BaseMySQLParserListener) ExitGrantOption(ctx *GrantOptionContext) {}

// EnterSetRole is called when production setRole is entered.
func (s *BaseMySQLParserListener) EnterSetRole(ctx *SetRoleContext) {}

// ExitSetRole is called when production setRole is exited.
func (s *BaseMySQLParserListener) ExitSetRole(ctx *SetRoleContext) {}

// EnterRoleList is called when production roleList is entered.
func (s *BaseMySQLParserListener) EnterRoleList(ctx *RoleListContext) {}

// ExitRoleList is called when production roleList is exited.
func (s *BaseMySQLParserListener) ExitRoleList(ctx *RoleListContext) {}

// EnterRole is called when production role is entered.
func (s *BaseMySQLParserListener) EnterRole(ctx *RoleContext) {}

// ExitRole is called when production role is exited.
func (s *BaseMySQLParserListener) ExitRole(ctx *RoleContext) {}

// EnterTableAdministrationStatement is called when production tableAdministrationStatement is entered.
func (s *BaseMySQLParserListener) EnterTableAdministrationStatement(ctx *TableAdministrationStatementContext) {
}

// ExitTableAdministrationStatement is called when production tableAdministrationStatement is exited.
func (s *BaseMySQLParserListener) ExitTableAdministrationStatement(ctx *TableAdministrationStatementContext) {
}

// EnterHistogram is called when production histogram is entered.
func (s *BaseMySQLParserListener) EnterHistogram(ctx *HistogramContext) {}

// ExitHistogram is called when production histogram is exited.
func (s *BaseMySQLParserListener) ExitHistogram(ctx *HistogramContext) {}

// EnterCheckOption is called when production checkOption is entered.
func (s *BaseMySQLParserListener) EnterCheckOption(ctx *CheckOptionContext) {}

// ExitCheckOption is called when production checkOption is exited.
func (s *BaseMySQLParserListener) ExitCheckOption(ctx *CheckOptionContext) {}

// EnterRepairType is called when production repairType is entered.
func (s *BaseMySQLParserListener) EnterRepairType(ctx *RepairTypeContext) {}

// ExitRepairType is called when production repairType is exited.
func (s *BaseMySQLParserListener) ExitRepairType(ctx *RepairTypeContext) {}

// EnterInstallUninstallStatment is called when production installUninstallStatment is entered.
func (s *BaseMySQLParserListener) EnterInstallUninstallStatment(ctx *InstallUninstallStatmentContext) {
}

// ExitInstallUninstallStatment is called when production installUninstallStatment is exited.
func (s *BaseMySQLParserListener) ExitInstallUninstallStatment(ctx *InstallUninstallStatmentContext) {
}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseMySQLParserListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseMySQLParserListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterStartOptionValueList is called when production startOptionValueList is entered.
func (s *BaseMySQLParserListener) EnterStartOptionValueList(ctx *StartOptionValueListContext) {}

// ExitStartOptionValueList is called when production startOptionValueList is exited.
func (s *BaseMySQLParserListener) ExitStartOptionValueList(ctx *StartOptionValueListContext) {}

// EnterTransactionCharacteristics is called when production transactionCharacteristics is entered.
func (s *BaseMySQLParserListener) EnterTransactionCharacteristics(ctx *TransactionCharacteristicsContext) {
}

// ExitTransactionCharacteristics is called when production transactionCharacteristics is exited.
func (s *BaseMySQLParserListener) ExitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) {
}

// EnterTransactionAccessMode is called when production transactionAccessMode is entered.
func (s *BaseMySQLParserListener) EnterTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// ExitTransactionAccessMode is called when production transactionAccessMode is exited.
func (s *BaseMySQLParserListener) ExitTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// EnterIsolationLevel is called when production isolationLevel is entered.
func (s *BaseMySQLParserListener) EnterIsolationLevel(ctx *IsolationLevelContext) {}

// ExitIsolationLevel is called when production isolationLevel is exited.
func (s *BaseMySQLParserListener) ExitIsolationLevel(ctx *IsolationLevelContext) {}

// EnterOptionValueListContinued is called when production optionValueListContinued is entered.
func (s *BaseMySQLParserListener) EnterOptionValueListContinued(ctx *OptionValueListContinuedContext) {
}

// ExitOptionValueListContinued is called when production optionValueListContinued is exited.
func (s *BaseMySQLParserListener) ExitOptionValueListContinued(ctx *OptionValueListContinuedContext) {
}

// EnterOptionValueNoOptionType is called when production optionValueNoOptionType is entered.
func (s *BaseMySQLParserListener) EnterOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) {}

// ExitOptionValueNoOptionType is called when production optionValueNoOptionType is exited.
func (s *BaseMySQLParserListener) ExitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) {}

// EnterOptionValue is called when production optionValue is entered.
func (s *BaseMySQLParserListener) EnterOptionValue(ctx *OptionValueContext) {}

// ExitOptionValue is called when production optionValue is exited.
func (s *BaseMySQLParserListener) ExitOptionValue(ctx *OptionValueContext) {}

// EnterSetSystemVariable is called when production setSystemVariable is entered.
func (s *BaseMySQLParserListener) EnterSetSystemVariable(ctx *SetSystemVariableContext) {}

// ExitSetSystemVariable is called when production setSystemVariable is exited.
func (s *BaseMySQLParserListener) ExitSetSystemVariable(ctx *SetSystemVariableContext) {}

// EnterStartOptionValueListFollowingOptionType is called when production startOptionValueListFollowingOptionType is entered.
func (s *BaseMySQLParserListener) EnterStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) {
}

// ExitStartOptionValueListFollowingOptionType is called when production startOptionValueListFollowingOptionType is exited.
func (s *BaseMySQLParserListener) ExitStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) {
}

// EnterOptionValueFollowingOptionType is called when production optionValueFollowingOptionType is entered.
func (s *BaseMySQLParserListener) EnterOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) {
}

// ExitOptionValueFollowingOptionType is called when production optionValueFollowingOptionType is exited.
func (s *BaseMySQLParserListener) ExitOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) {
}

// EnterSetExprOrDefault is called when production setExprOrDefault is entered.
func (s *BaseMySQLParserListener) EnterSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// ExitSetExprOrDefault is called when production setExprOrDefault is exited.
func (s *BaseMySQLParserListener) ExitSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// EnterShowStatement is called when production showStatement is entered.
func (s *BaseMySQLParserListener) EnterShowStatement(ctx *ShowStatementContext) {}

// ExitShowStatement is called when production showStatement is exited.
func (s *BaseMySQLParserListener) ExitShowStatement(ctx *ShowStatementContext) {}

// EnterShowCommandType is called when production showCommandType is entered.
func (s *BaseMySQLParserListener) EnterShowCommandType(ctx *ShowCommandTypeContext) {}

// ExitShowCommandType is called when production showCommandType is exited.
func (s *BaseMySQLParserListener) ExitShowCommandType(ctx *ShowCommandTypeContext) {}

// EnterNonBlocking is called when production nonBlocking is entered.
func (s *BaseMySQLParserListener) EnterNonBlocking(ctx *NonBlockingContext) {}

// ExitNonBlocking is called when production nonBlocking is exited.
func (s *BaseMySQLParserListener) ExitNonBlocking(ctx *NonBlockingContext) {}

// EnterFromOrIn is called when production fromOrIn is entered.
func (s *BaseMySQLParserListener) EnterFromOrIn(ctx *FromOrInContext) {}

// ExitFromOrIn is called when production fromOrIn is exited.
func (s *BaseMySQLParserListener) ExitFromOrIn(ctx *FromOrInContext) {}

// EnterInDb is called when production inDb is entered.
func (s *BaseMySQLParserListener) EnterInDb(ctx *InDbContext) {}

// ExitInDb is called when production inDb is exited.
func (s *BaseMySQLParserListener) ExitInDb(ctx *InDbContext) {}

// EnterProfileType is called when production profileType is entered.
func (s *BaseMySQLParserListener) EnterProfileType(ctx *ProfileTypeContext) {}

// ExitProfileType is called when production profileType is exited.
func (s *BaseMySQLParserListener) ExitProfileType(ctx *ProfileTypeContext) {}

// EnterOtherAdministrativeStatement is called when production otherAdministrativeStatement is entered.
func (s *BaseMySQLParserListener) EnterOtherAdministrativeStatement(ctx *OtherAdministrativeStatementContext) {
}

// ExitOtherAdministrativeStatement is called when production otherAdministrativeStatement is exited.
func (s *BaseMySQLParserListener) ExitOtherAdministrativeStatement(ctx *OtherAdministrativeStatementContext) {
}

// EnterKeyCacheListOrParts is called when production keyCacheListOrParts is entered.
func (s *BaseMySQLParserListener) EnterKeyCacheListOrParts(ctx *KeyCacheListOrPartsContext) {}

// ExitKeyCacheListOrParts is called when production keyCacheListOrParts is exited.
func (s *BaseMySQLParserListener) ExitKeyCacheListOrParts(ctx *KeyCacheListOrPartsContext) {}

// EnterKeyCacheList is called when production keyCacheList is entered.
func (s *BaseMySQLParserListener) EnterKeyCacheList(ctx *KeyCacheListContext) {}

// ExitKeyCacheList is called when production keyCacheList is exited.
func (s *BaseMySQLParserListener) ExitKeyCacheList(ctx *KeyCacheListContext) {}

// EnterAssignToKeycache is called when production assignToKeycache is entered.
func (s *BaseMySQLParserListener) EnterAssignToKeycache(ctx *AssignToKeycacheContext) {}

// ExitAssignToKeycache is called when production assignToKeycache is exited.
func (s *BaseMySQLParserListener) ExitAssignToKeycache(ctx *AssignToKeycacheContext) {}

// EnterAssignToKeycachePartition is called when production assignToKeycachePartition is entered.
func (s *BaseMySQLParserListener) EnterAssignToKeycachePartition(ctx *AssignToKeycachePartitionContext) {
}

// ExitAssignToKeycachePartition is called when production assignToKeycachePartition is exited.
func (s *BaseMySQLParserListener) ExitAssignToKeycachePartition(ctx *AssignToKeycachePartitionContext) {
}

// EnterCacheKeyList is called when production cacheKeyList is entered.
func (s *BaseMySQLParserListener) EnterCacheKeyList(ctx *CacheKeyListContext) {}

// ExitCacheKeyList is called when production cacheKeyList is exited.
func (s *BaseMySQLParserListener) ExitCacheKeyList(ctx *CacheKeyListContext) {}

// EnterKeyUsageElement is called when production keyUsageElement is entered.
func (s *BaseMySQLParserListener) EnterKeyUsageElement(ctx *KeyUsageElementContext) {}

// ExitKeyUsageElement is called when production keyUsageElement is exited.
func (s *BaseMySQLParserListener) ExitKeyUsageElement(ctx *KeyUsageElementContext) {}

// EnterKeyUsageList is called when production keyUsageList is entered.
func (s *BaseMySQLParserListener) EnterKeyUsageList(ctx *KeyUsageListContext) {}

// ExitKeyUsageList is called when production keyUsageList is exited.
func (s *BaseMySQLParserListener) ExitKeyUsageList(ctx *KeyUsageListContext) {}

// EnterFlushOption is called when production flushOption is entered.
func (s *BaseMySQLParserListener) EnterFlushOption(ctx *FlushOptionContext) {}

// ExitFlushOption is called when production flushOption is exited.
func (s *BaseMySQLParserListener) ExitFlushOption(ctx *FlushOptionContext) {}

// EnterLogType is called when production logType is entered.
func (s *BaseMySQLParserListener) EnterLogType(ctx *LogTypeContext) {}

// ExitLogType is called when production logType is exited.
func (s *BaseMySQLParserListener) ExitLogType(ctx *LogTypeContext) {}

// EnterFlushTables is called when production flushTables is entered.
func (s *BaseMySQLParserListener) EnterFlushTables(ctx *FlushTablesContext) {}

// ExitFlushTables is called when production flushTables is exited.
func (s *BaseMySQLParserListener) ExitFlushTables(ctx *FlushTablesContext) {}

// EnterFlushTablesOptions is called when production flushTablesOptions is entered.
func (s *BaseMySQLParserListener) EnterFlushTablesOptions(ctx *FlushTablesOptionsContext) {}

// ExitFlushTablesOptions is called when production flushTablesOptions is exited.
func (s *BaseMySQLParserListener) ExitFlushTablesOptions(ctx *FlushTablesOptionsContext) {}

// EnterPreloadTail is called when production preloadTail is entered.
func (s *BaseMySQLParserListener) EnterPreloadTail(ctx *PreloadTailContext) {}

// ExitPreloadTail is called when production preloadTail is exited.
func (s *BaseMySQLParserListener) ExitPreloadTail(ctx *PreloadTailContext) {}

// EnterPreloadList is called when production preloadList is entered.
func (s *BaseMySQLParserListener) EnterPreloadList(ctx *PreloadListContext) {}

// ExitPreloadList is called when production preloadList is exited.
func (s *BaseMySQLParserListener) ExitPreloadList(ctx *PreloadListContext) {}

// EnterPreloadKeys is called when production preloadKeys is entered.
func (s *BaseMySQLParserListener) EnterPreloadKeys(ctx *PreloadKeysContext) {}

// ExitPreloadKeys is called when production preloadKeys is exited.
func (s *BaseMySQLParserListener) ExitPreloadKeys(ctx *PreloadKeysContext) {}

// EnterAdminPartition is called when production adminPartition is entered.
func (s *BaseMySQLParserListener) EnterAdminPartition(ctx *AdminPartitionContext) {}

// ExitAdminPartition is called when production adminPartition is exited.
func (s *BaseMySQLParserListener) ExitAdminPartition(ctx *AdminPartitionContext) {}

// EnterResourceGroupManagement is called when production resourceGroupManagement is entered.
func (s *BaseMySQLParserListener) EnterResourceGroupManagement(ctx *ResourceGroupManagementContext) {}

// ExitResourceGroupManagement is called when production resourceGroupManagement is exited.
func (s *BaseMySQLParserListener) ExitResourceGroupManagement(ctx *ResourceGroupManagementContext) {}

// EnterCreateResourceGroup is called when production createResourceGroup is entered.
func (s *BaseMySQLParserListener) EnterCreateResourceGroup(ctx *CreateResourceGroupContext) {}

// ExitCreateResourceGroup is called when production createResourceGroup is exited.
func (s *BaseMySQLParserListener) ExitCreateResourceGroup(ctx *CreateResourceGroupContext) {}

// EnterResourceGroupVcpuList is called when production resourceGroupVcpuList is entered.
func (s *BaseMySQLParserListener) EnterResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) {}

// ExitResourceGroupVcpuList is called when production resourceGroupVcpuList is exited.
func (s *BaseMySQLParserListener) ExitResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) {}

// EnterVcpuNumOrRange is called when production vcpuNumOrRange is entered.
func (s *BaseMySQLParserListener) EnterVcpuNumOrRange(ctx *VcpuNumOrRangeContext) {}

// ExitVcpuNumOrRange is called when production vcpuNumOrRange is exited.
func (s *BaseMySQLParserListener) ExitVcpuNumOrRange(ctx *VcpuNumOrRangeContext) {}

// EnterResourceGroupPriority is called when production resourceGroupPriority is entered.
func (s *BaseMySQLParserListener) EnterResourceGroupPriority(ctx *ResourceGroupPriorityContext) {}

// ExitResourceGroupPriority is called when production resourceGroupPriority is exited.
func (s *BaseMySQLParserListener) ExitResourceGroupPriority(ctx *ResourceGroupPriorityContext) {}

// EnterResourceGroupEnableDisable is called when production resourceGroupEnableDisable is entered.
func (s *BaseMySQLParserListener) EnterResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) {
}

// ExitResourceGroupEnableDisable is called when production resourceGroupEnableDisable is exited.
func (s *BaseMySQLParserListener) ExitResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) {
}

// EnterAlterResourceGroup is called when production alterResourceGroup is entered.
func (s *BaseMySQLParserListener) EnterAlterResourceGroup(ctx *AlterResourceGroupContext) {}

// ExitAlterResourceGroup is called when production alterResourceGroup is exited.
func (s *BaseMySQLParserListener) ExitAlterResourceGroup(ctx *AlterResourceGroupContext) {}

// EnterSetResourceGroup is called when production setResourceGroup is entered.
func (s *BaseMySQLParserListener) EnterSetResourceGroup(ctx *SetResourceGroupContext) {}

// ExitSetResourceGroup is called when production setResourceGroup is exited.
func (s *BaseMySQLParserListener) ExitSetResourceGroup(ctx *SetResourceGroupContext) {}

// EnterThreadIdList is called when production threadIdList is entered.
func (s *BaseMySQLParserListener) EnterThreadIdList(ctx *ThreadIdListContext) {}

// ExitThreadIdList is called when production threadIdList is exited.
func (s *BaseMySQLParserListener) ExitThreadIdList(ctx *ThreadIdListContext) {}

// EnterDropResourceGroup is called when production dropResourceGroup is entered.
func (s *BaseMySQLParserListener) EnterDropResourceGroup(ctx *DropResourceGroupContext) {}

// ExitDropResourceGroup is called when production dropResourceGroup is exited.
func (s *BaseMySQLParserListener) ExitDropResourceGroup(ctx *DropResourceGroupContext) {}

// EnterUtilityStatement is called when production utilityStatement is entered.
func (s *BaseMySQLParserListener) EnterUtilityStatement(ctx *UtilityStatementContext) {}

// ExitUtilityStatement is called when production utilityStatement is exited.
func (s *BaseMySQLParserListener) ExitUtilityStatement(ctx *UtilityStatementContext) {}

// EnterDescribeStatement is called when production describeStatement is entered.
func (s *BaseMySQLParserListener) EnterDescribeStatement(ctx *DescribeStatementContext) {}

// ExitDescribeStatement is called when production describeStatement is exited.
func (s *BaseMySQLParserListener) ExitDescribeStatement(ctx *DescribeStatementContext) {}

// EnterExplainStatement is called when production explainStatement is entered.
func (s *BaseMySQLParserListener) EnterExplainStatement(ctx *ExplainStatementContext) {}

// ExitExplainStatement is called when production explainStatement is exited.
func (s *BaseMySQLParserListener) ExitExplainStatement(ctx *ExplainStatementContext) {}

// EnterExplainableStatement is called when production explainableStatement is entered.
func (s *BaseMySQLParserListener) EnterExplainableStatement(ctx *ExplainableStatementContext) {}

// ExitExplainableStatement is called when production explainableStatement is exited.
func (s *BaseMySQLParserListener) ExitExplainableStatement(ctx *ExplainableStatementContext) {}

// EnterHelpCommand is called when production helpCommand is entered.
func (s *BaseMySQLParserListener) EnterHelpCommand(ctx *HelpCommandContext) {}

// ExitHelpCommand is called when production helpCommand is exited.
func (s *BaseMySQLParserListener) ExitHelpCommand(ctx *HelpCommandContext) {}

// EnterUseCommand is called when production useCommand is entered.
func (s *BaseMySQLParserListener) EnterUseCommand(ctx *UseCommandContext) {}

// ExitUseCommand is called when production useCommand is exited.
func (s *BaseMySQLParserListener) ExitUseCommand(ctx *UseCommandContext) {}

// EnterRestartServer is called when production restartServer is entered.
func (s *BaseMySQLParserListener) EnterRestartServer(ctx *RestartServerContext) {}

// ExitRestartServer is called when production restartServer is exited.
func (s *BaseMySQLParserListener) ExitRestartServer(ctx *RestartServerContext) {}

// EnterExprOr is called when production exprOr is entered.
func (s *BaseMySQLParserListener) EnterExprOr(ctx *ExprOrContext) {}

// ExitExprOr is called when production exprOr is exited.
func (s *BaseMySQLParserListener) ExitExprOr(ctx *ExprOrContext) {}

// EnterExprNot is called when production exprNot is entered.
func (s *BaseMySQLParserListener) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production exprNot is exited.
func (s *BaseMySQLParserListener) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprIs is called when production exprIs is entered.
func (s *BaseMySQLParserListener) EnterExprIs(ctx *ExprIsContext) {}

// ExitExprIs is called when production exprIs is exited.
func (s *BaseMySQLParserListener) ExitExprIs(ctx *ExprIsContext) {}

// EnterExprAnd is called when production exprAnd is entered.
func (s *BaseMySQLParserListener) EnterExprAnd(ctx *ExprAndContext) {}

// ExitExprAnd is called when production exprAnd is exited.
func (s *BaseMySQLParserListener) ExitExprAnd(ctx *ExprAndContext) {}

// EnterExprXor is called when production exprXor is entered.
func (s *BaseMySQLParserListener) EnterExprXor(ctx *ExprXorContext) {}

// ExitExprXor is called when production exprXor is exited.
func (s *BaseMySQLParserListener) ExitExprXor(ctx *ExprXorContext) {}

// EnterPrimaryExprPredicate is called when production primaryExprPredicate is entered.
func (s *BaseMySQLParserListener) EnterPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) {}

// ExitPrimaryExprPredicate is called when production primaryExprPredicate is exited.
func (s *BaseMySQLParserListener) ExitPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) {}

// EnterPrimaryExprCompare is called when production primaryExprCompare is entered.
func (s *BaseMySQLParserListener) EnterPrimaryExprCompare(ctx *PrimaryExprCompareContext) {}

// ExitPrimaryExprCompare is called when production primaryExprCompare is exited.
func (s *BaseMySQLParserListener) ExitPrimaryExprCompare(ctx *PrimaryExprCompareContext) {}

// EnterPrimaryExprAllAny is called when production primaryExprAllAny is entered.
func (s *BaseMySQLParserListener) EnterPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) {}

// ExitPrimaryExprAllAny is called when production primaryExprAllAny is exited.
func (s *BaseMySQLParserListener) ExitPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) {}

// EnterPrimaryExprIsNull is called when production primaryExprIsNull is entered.
func (s *BaseMySQLParserListener) EnterPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) {}

// ExitPrimaryExprIsNull is called when production primaryExprIsNull is exited.
func (s *BaseMySQLParserListener) ExitPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) {}

// EnterCompOp is called when production compOp is entered.
func (s *BaseMySQLParserListener) EnterCompOp(ctx *CompOpContext) {}

// ExitCompOp is called when production compOp is exited.
func (s *BaseMySQLParserListener) ExitCompOp(ctx *CompOpContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseMySQLParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseMySQLParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterPredicateExprIn is called when production predicateExprIn is entered.
func (s *BaseMySQLParserListener) EnterPredicateExprIn(ctx *PredicateExprInContext) {}

// ExitPredicateExprIn is called when production predicateExprIn is exited.
func (s *BaseMySQLParserListener) ExitPredicateExprIn(ctx *PredicateExprInContext) {}

// EnterPredicateExprBetween is called when production predicateExprBetween is entered.
func (s *BaseMySQLParserListener) EnterPredicateExprBetween(ctx *PredicateExprBetweenContext) {}

// ExitPredicateExprBetween is called when production predicateExprBetween is exited.
func (s *BaseMySQLParserListener) ExitPredicateExprBetween(ctx *PredicateExprBetweenContext) {}

// EnterPredicateExprLike is called when production predicateExprLike is entered.
func (s *BaseMySQLParserListener) EnterPredicateExprLike(ctx *PredicateExprLikeContext) {}

// ExitPredicateExprLike is called when production predicateExprLike is exited.
func (s *BaseMySQLParserListener) ExitPredicateExprLike(ctx *PredicateExprLikeContext) {}

// EnterPredicateExprRegex is called when production predicateExprRegex is entered.
func (s *BaseMySQLParserListener) EnterPredicateExprRegex(ctx *PredicateExprRegexContext) {}

// ExitPredicateExprRegex is called when production predicateExprRegex is exited.
func (s *BaseMySQLParserListener) ExitPredicateExprRegex(ctx *PredicateExprRegexContext) {}

// EnterBitExpr is called when production bitExpr is entered.
func (s *BaseMySQLParserListener) EnterBitExpr(ctx *BitExprContext) {}

// ExitBitExpr is called when production bitExpr is exited.
func (s *BaseMySQLParserListener) ExitBitExpr(ctx *BitExprContext) {}

// EnterSimpleExprConvert is called when production simpleExprConvert is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprConvert(ctx *SimpleExprConvertContext) {}

// ExitSimpleExprConvert is called when production simpleExprConvert is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprConvert(ctx *SimpleExprConvertContext) {}

// EnterSimpleExprSearchJson is called when production simpleExprSearchJson is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) {}

// ExitSimpleExprSearchJson is called when production simpleExprSearchJson is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) {}

// EnterSimpleExprVariable is called when production simpleExprVariable is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprVariable(ctx *SimpleExprVariableContext) {}

// ExitSimpleExprVariable is called when production simpleExprVariable is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprVariable(ctx *SimpleExprVariableContext) {}

// EnterSimpleExprCast is called when production simpleExprCast is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprCast(ctx *SimpleExprCastContext) {}

// ExitSimpleExprCast is called when production simpleExprCast is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprCast(ctx *SimpleExprCastContext) {}

// EnterSimpleExprUnary is called when production simpleExprUnary is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprUnary(ctx *SimpleExprUnaryContext) {}

// ExitSimpleExprUnary is called when production simpleExprUnary is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprUnary(ctx *SimpleExprUnaryContext) {}

// EnterSimpleExprOdbc is called when production simpleExprOdbc is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprOdbc(ctx *SimpleExprOdbcContext) {}

// ExitSimpleExprOdbc is called when production simpleExprOdbc is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprOdbc(ctx *SimpleExprOdbcContext) {}

// EnterSimpleExprRuntimeFunction is called when production simpleExprRuntimeFunction is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) {
}

// ExitSimpleExprRuntimeFunction is called when production simpleExprRuntimeFunction is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) {
}

// EnterSimpleExprFunction is called when production simpleExprFunction is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprFunction(ctx *SimpleExprFunctionContext) {}

// ExitSimpleExprFunction is called when production simpleExprFunction is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprFunction(ctx *SimpleExprFunctionContext) {}

// EnterSimpleExprCollate is called when production simpleExprCollate is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprCollate(ctx *SimpleExprCollateContext) {}

// ExitSimpleExprCollate is called when production simpleExprCollate is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprCollate(ctx *SimpleExprCollateContext) {}

// EnterSimpleExprMatch is called when production simpleExprMatch is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprMatch(ctx *SimpleExprMatchContext) {}

// ExitSimpleExprMatch is called when production simpleExprMatch is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprMatch(ctx *SimpleExprMatchContext) {}

// EnterSimpleExprWindowingFunction is called when production simpleExprWindowingFunction is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) {
}

// ExitSimpleExprWindowingFunction is called when production simpleExprWindowingFunction is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) {
}

// EnterSimpleExprBinary is called when production simpleExprBinary is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprBinary(ctx *SimpleExprBinaryContext) {}

// ExitSimpleExprBinary is called when production simpleExprBinary is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprBinary(ctx *SimpleExprBinaryContext) {}

// EnterSimpleExprColumnRef is called when production simpleExprColumnRef is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) {}

// ExitSimpleExprColumnRef is called when production simpleExprColumnRef is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) {}

// EnterSimpleExprParamMarker is called when production simpleExprParamMarker is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) {}

// ExitSimpleExprParamMarker is called when production simpleExprParamMarker is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) {}

// EnterSimpleExprSum is called when production simpleExprSum is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprSum(ctx *SimpleExprSumContext) {}

// ExitSimpleExprSum is called when production simpleExprSum is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprSum(ctx *SimpleExprSumContext) {}

// EnterSimpleExprConvertUsing is called when production simpleExprConvertUsing is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) {}

// ExitSimpleExprConvertUsing is called when production simpleExprConvertUsing is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) {}

// EnterSimpleExprSubQuery is called when production simpleExprSubQuery is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) {}

// ExitSimpleExprSubQuery is called when production simpleExprSubQuery is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) {}

// EnterSimpleExprGroupingOperation is called when production simpleExprGroupingOperation is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) {
}

// ExitSimpleExprGroupingOperation is called when production simpleExprGroupingOperation is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) {
}

// EnterSimpleExprNot is called when production simpleExprNot is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprNot(ctx *SimpleExprNotContext) {}

// ExitSimpleExprNot is called when production simpleExprNot is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprNot(ctx *SimpleExprNotContext) {}

// EnterSimpleExprValues is called when production simpleExprValues is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprValues(ctx *SimpleExprValuesContext) {}

// ExitSimpleExprValues is called when production simpleExprValues is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprValues(ctx *SimpleExprValuesContext) {}

// EnterSimpleExprDefault is called when production simpleExprDefault is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprDefault(ctx *SimpleExprDefaultContext) {}

// ExitSimpleExprDefault is called when production simpleExprDefault is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprDefault(ctx *SimpleExprDefaultContext) {}

// EnterSimpleExprList is called when production simpleExprList is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprList(ctx *SimpleExprListContext) {}

// ExitSimpleExprList is called when production simpleExprList is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprList(ctx *SimpleExprListContext) {}

// EnterSimpleExprInterval is called when production simpleExprInterval is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprInterval(ctx *SimpleExprIntervalContext) {}

// ExitSimpleExprInterval is called when production simpleExprInterval is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprInterval(ctx *SimpleExprIntervalContext) {}

// EnterSimpleExprCase is called when production simpleExprCase is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprCase(ctx *SimpleExprCaseContext) {}

// ExitSimpleExprCase is called when production simpleExprCase is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprCase(ctx *SimpleExprCaseContext) {}

// EnterSimpleExprConcat is called when production simpleExprConcat is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprConcat(ctx *SimpleExprConcatContext) {}

// ExitSimpleExprConcat is called when production simpleExprConcat is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprConcat(ctx *SimpleExprConcatContext) {}

// EnterSimpleExprLiteral is called when production simpleExprLiteral is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprLiteral(ctx *SimpleExprLiteralContext) {}

// ExitSimpleExprLiteral is called when production simpleExprLiteral is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprLiteral(ctx *SimpleExprLiteralContext) {}

// EnterArrayCast is called when production arrayCast is entered.
func (s *BaseMySQLParserListener) EnterArrayCast(ctx *ArrayCastContext) {}

// ExitArrayCast is called when production arrayCast is exited.
func (s *BaseMySQLParserListener) ExitArrayCast(ctx *ArrayCastContext) {}

// EnterJsonOperator is called when production jsonOperator is entered.
func (s *BaseMySQLParserListener) EnterJsonOperator(ctx *JsonOperatorContext) {}

// ExitJsonOperator is called when production jsonOperator is exited.
func (s *BaseMySQLParserListener) ExitJsonOperator(ctx *JsonOperatorContext) {}

// EnterSumExpr is called when production sumExpr is entered.
func (s *BaseMySQLParserListener) EnterSumExpr(ctx *SumExprContext) {}

// ExitSumExpr is called when production sumExpr is exited.
func (s *BaseMySQLParserListener) ExitSumExpr(ctx *SumExprContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseMySQLParserListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseMySQLParserListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterWindowFunctionCall is called when production windowFunctionCall is entered.
func (s *BaseMySQLParserListener) EnterWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// ExitWindowFunctionCall is called when production windowFunctionCall is exited.
func (s *BaseMySQLParserListener) ExitWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// EnterWindowingClause is called when production windowingClause is entered.
func (s *BaseMySQLParserListener) EnterWindowingClause(ctx *WindowingClauseContext) {}

// ExitWindowingClause is called when production windowingClause is exited.
func (s *BaseMySQLParserListener) ExitWindowingClause(ctx *WindowingClauseContext) {}

// EnterLeadLagInfo is called when production leadLagInfo is entered.
func (s *BaseMySQLParserListener) EnterLeadLagInfo(ctx *LeadLagInfoContext) {}

// ExitLeadLagInfo is called when production leadLagInfo is exited.
func (s *BaseMySQLParserListener) ExitLeadLagInfo(ctx *LeadLagInfoContext) {}

// EnterNullTreatment is called when production nullTreatment is entered.
func (s *BaseMySQLParserListener) EnterNullTreatment(ctx *NullTreatmentContext) {}

// ExitNullTreatment is called when production nullTreatment is exited.
func (s *BaseMySQLParserListener) ExitNullTreatment(ctx *NullTreatmentContext) {}

// EnterJsonFunction is called when production jsonFunction is entered.
func (s *BaseMySQLParserListener) EnterJsonFunction(ctx *JsonFunctionContext) {}

// ExitJsonFunction is called when production jsonFunction is exited.
func (s *BaseMySQLParserListener) ExitJsonFunction(ctx *JsonFunctionContext) {}

// EnterInSumExpr is called when production inSumExpr is entered.
func (s *BaseMySQLParserListener) EnterInSumExpr(ctx *InSumExprContext) {}

// ExitInSumExpr is called when production inSumExpr is exited.
func (s *BaseMySQLParserListener) ExitInSumExpr(ctx *InSumExprContext) {}

// EnterIdentListArg is called when production identListArg is entered.
func (s *BaseMySQLParserListener) EnterIdentListArg(ctx *IdentListArgContext) {}

// ExitIdentListArg is called when production identListArg is exited.
func (s *BaseMySQLParserListener) ExitIdentListArg(ctx *IdentListArgContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseMySQLParserListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseMySQLParserListener) ExitIdentList(ctx *IdentListContext) {}

// EnterFulltextOptions is called when production fulltextOptions is entered.
func (s *BaseMySQLParserListener) EnterFulltextOptions(ctx *FulltextOptionsContext) {}

// ExitFulltextOptions is called when production fulltextOptions is exited.
func (s *BaseMySQLParserListener) ExitFulltextOptions(ctx *FulltextOptionsContext) {}

// EnterRuntimeFunctionCall is called when production runtimeFunctionCall is entered.
func (s *BaseMySQLParserListener) EnterRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) {}

// ExitRuntimeFunctionCall is called when production runtimeFunctionCall is exited.
func (s *BaseMySQLParserListener) ExitRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) {}

// EnterGeometryFunction is called when production geometryFunction is entered.
func (s *BaseMySQLParserListener) EnterGeometryFunction(ctx *GeometryFunctionContext) {}

// ExitGeometryFunction is called when production geometryFunction is exited.
func (s *BaseMySQLParserListener) ExitGeometryFunction(ctx *GeometryFunctionContext) {}

// EnterTimeFunctionParameters is called when production timeFunctionParameters is entered.
func (s *BaseMySQLParserListener) EnterTimeFunctionParameters(ctx *TimeFunctionParametersContext) {}

// ExitTimeFunctionParameters is called when production timeFunctionParameters is exited.
func (s *BaseMySQLParserListener) ExitTimeFunctionParameters(ctx *TimeFunctionParametersContext) {}

// EnterFractionalPrecision is called when production fractionalPrecision is entered.
func (s *BaseMySQLParserListener) EnterFractionalPrecision(ctx *FractionalPrecisionContext) {}

// ExitFractionalPrecision is called when production fractionalPrecision is exited.
func (s *BaseMySQLParserListener) ExitFractionalPrecision(ctx *FractionalPrecisionContext) {}

// EnterWeightStringLevels is called when production weightStringLevels is entered.
func (s *BaseMySQLParserListener) EnterWeightStringLevels(ctx *WeightStringLevelsContext) {}

// ExitWeightStringLevels is called when production weightStringLevels is exited.
func (s *BaseMySQLParserListener) ExitWeightStringLevels(ctx *WeightStringLevelsContext) {}

// EnterWeightStringLevelListItem is called when production weightStringLevelListItem is entered.
func (s *BaseMySQLParserListener) EnterWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) {
}

// ExitWeightStringLevelListItem is called when production weightStringLevelListItem is exited.
func (s *BaseMySQLParserListener) ExitWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) {
}

// EnterDateTimeTtype is called when production dateTimeTtype is entered.
func (s *BaseMySQLParserListener) EnterDateTimeTtype(ctx *DateTimeTtypeContext) {}

// ExitDateTimeTtype is called when production dateTimeTtype is exited.
func (s *BaseMySQLParserListener) ExitDateTimeTtype(ctx *DateTimeTtypeContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *BaseMySQLParserListener) EnterTrimFunction(ctx *TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *BaseMySQLParserListener) ExitTrimFunction(ctx *TrimFunctionContext) {}

// EnterSubstringFunction is called when production substringFunction is entered.
func (s *BaseMySQLParserListener) EnterSubstringFunction(ctx *SubstringFunctionContext) {}

// ExitSubstringFunction is called when production substringFunction is exited.
func (s *BaseMySQLParserListener) ExitSubstringFunction(ctx *SubstringFunctionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseMySQLParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseMySQLParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterSearchJsonFunction is called when production searchJsonFunction is entered.
func (s *BaseMySQLParserListener) EnterSearchJsonFunction(ctx *SearchJsonFunctionContext) {}

// ExitSearchJsonFunction is called when production searchJsonFunction is exited.
func (s *BaseMySQLParserListener) ExitSearchJsonFunction(ctx *SearchJsonFunctionContext) {}

// EnterJsonValueReturning is called when production jsonValueReturning is entered.
func (s *BaseMySQLParserListener) EnterJsonValueReturning(ctx *JsonValueReturningContext) {}

// ExitJsonValueReturning is called when production jsonValueReturning is exited.
func (s *BaseMySQLParserListener) ExitJsonValueReturning(ctx *JsonValueReturningContext) {}

// EnterJsonValueOnEmpty is called when production jsonValueOnEmpty is entered.
func (s *BaseMySQLParserListener) EnterJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) {}

// ExitJsonValueOnEmpty is called when production jsonValueOnEmpty is exited.
func (s *BaseMySQLParserListener) ExitJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) {}

// EnterJsonValueOnError is called when production jsonValueOnError is entered.
func (s *BaseMySQLParserListener) EnterJsonValueOnError(ctx *JsonValueOnErrorContext) {}

// ExitJsonValueOnError is called when production jsonValueOnError is exited.
func (s *BaseMySQLParserListener) ExitJsonValueOnError(ctx *JsonValueOnErrorContext) {}

// EnterUdfExprList is called when production udfExprList is entered.
func (s *BaseMySQLParserListener) EnterUdfExprList(ctx *UdfExprListContext) {}

// ExitUdfExprList is called when production udfExprList is exited.
func (s *BaseMySQLParserListener) ExitUdfExprList(ctx *UdfExprListContext) {}

// EnterUdfExpr is called when production udfExpr is entered.
func (s *BaseMySQLParserListener) EnterUdfExpr(ctx *UdfExprContext) {}

// ExitUdfExpr is called when production udfExpr is exited.
func (s *BaseMySQLParserListener) ExitUdfExpr(ctx *UdfExprContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseMySQLParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseMySQLParserListener) ExitVariable(ctx *VariableContext) {}

// EnterUserVariable is called when production userVariable is entered.
func (s *BaseMySQLParserListener) EnterUserVariable(ctx *UserVariableContext) {}

// ExitUserVariable is called when production userVariable is exited.
func (s *BaseMySQLParserListener) ExitUserVariable(ctx *UserVariableContext) {}

// EnterSystemVariable is called when production systemVariable is entered.
func (s *BaseMySQLParserListener) EnterSystemVariable(ctx *SystemVariableContext) {}

// ExitSystemVariable is called when production systemVariable is exited.
func (s *BaseMySQLParserListener) ExitSystemVariable(ctx *SystemVariableContext) {}

// EnterInternalVariableName is called when production internalVariableName is entered.
func (s *BaseMySQLParserListener) EnterInternalVariableName(ctx *InternalVariableNameContext) {}

// ExitInternalVariableName is called when production internalVariableName is exited.
func (s *BaseMySQLParserListener) ExitInternalVariableName(ctx *InternalVariableNameContext) {}

// EnterWhenExpression is called when production whenExpression is entered.
func (s *BaseMySQLParserListener) EnterWhenExpression(ctx *WhenExpressionContext) {}

// ExitWhenExpression is called when production whenExpression is exited.
func (s *BaseMySQLParserListener) ExitWhenExpression(ctx *WhenExpressionContext) {}

// EnterThenExpression is called when production thenExpression is entered.
func (s *BaseMySQLParserListener) EnterThenExpression(ctx *ThenExpressionContext) {}

// ExitThenExpression is called when production thenExpression is exited.
func (s *BaseMySQLParserListener) ExitThenExpression(ctx *ThenExpressionContext) {}

// EnterElseExpression is called when production elseExpression is entered.
func (s *BaseMySQLParserListener) EnterElseExpression(ctx *ElseExpressionContext) {}

// ExitElseExpression is called when production elseExpression is exited.
func (s *BaseMySQLParserListener) ExitElseExpression(ctx *ElseExpressionContext) {}

// EnterCastType is called when production castType is entered.
func (s *BaseMySQLParserListener) EnterCastType(ctx *CastTypeContext) {}

// ExitCastType is called when production castType is exited.
func (s *BaseMySQLParserListener) ExitCastType(ctx *CastTypeContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseMySQLParserListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseMySQLParserListener) ExitExprList(ctx *ExprListContext) {}

// EnterCharset is called when production charset is entered.
func (s *BaseMySQLParserListener) EnterCharset(ctx *CharsetContext) {}

// ExitCharset is called when production charset is exited.
func (s *BaseMySQLParserListener) ExitCharset(ctx *CharsetContext) {}

// EnterNotRule is called when production notRule is entered.
func (s *BaseMySQLParserListener) EnterNotRule(ctx *NotRuleContext) {}

// ExitNotRule is called when production notRule is exited.
func (s *BaseMySQLParserListener) ExitNotRule(ctx *NotRuleContext) {}

// EnterNot2Rule is called when production not2Rule is entered.
func (s *BaseMySQLParserListener) EnterNot2Rule(ctx *Not2RuleContext) {}

// ExitNot2Rule is called when production not2Rule is exited.
func (s *BaseMySQLParserListener) ExitNot2Rule(ctx *Not2RuleContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseMySQLParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseMySQLParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterIntervalTimeStamp is called when production intervalTimeStamp is entered.
func (s *BaseMySQLParserListener) EnterIntervalTimeStamp(ctx *IntervalTimeStampContext) {}

// ExitIntervalTimeStamp is called when production intervalTimeStamp is exited.
func (s *BaseMySQLParserListener) ExitIntervalTimeStamp(ctx *IntervalTimeStampContext) {}

// EnterExprListWithParentheses is called when production exprListWithParentheses is entered.
func (s *BaseMySQLParserListener) EnterExprListWithParentheses(ctx *ExprListWithParenthesesContext) {}

// ExitExprListWithParentheses is called when production exprListWithParentheses is exited.
func (s *BaseMySQLParserListener) ExitExprListWithParentheses(ctx *ExprListWithParenthesesContext) {}

// EnterExprWithParentheses is called when production exprWithParentheses is entered.
func (s *BaseMySQLParserListener) EnterExprWithParentheses(ctx *ExprWithParenthesesContext) {}

// ExitExprWithParentheses is called when production exprWithParentheses is exited.
func (s *BaseMySQLParserListener) ExitExprWithParentheses(ctx *ExprWithParenthesesContext) {}

// EnterSimpleExprWithParentheses is called when production simpleExprWithParentheses is entered.
func (s *BaseMySQLParserListener) EnterSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) {
}

// ExitSimpleExprWithParentheses is called when production simpleExprWithParentheses is exited.
func (s *BaseMySQLParserListener) ExitSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) {
}

// EnterOrderList is called when production orderList is entered.
func (s *BaseMySQLParserListener) EnterOrderList(ctx *OrderListContext) {}

// ExitOrderList is called when production orderList is exited.
func (s *BaseMySQLParserListener) ExitOrderList(ctx *OrderListContext) {}

// EnterOrderExpression is called when production orderExpression is entered.
func (s *BaseMySQLParserListener) EnterOrderExpression(ctx *OrderExpressionContext) {}

// ExitOrderExpression is called when production orderExpression is exited.
func (s *BaseMySQLParserListener) ExitOrderExpression(ctx *OrderExpressionContext) {}

// EnterGroupList is called when production groupList is entered.
func (s *BaseMySQLParserListener) EnterGroupList(ctx *GroupListContext) {}

// ExitGroupList is called when production groupList is exited.
func (s *BaseMySQLParserListener) ExitGroupList(ctx *GroupListContext) {}

// EnterGroupingExpression is called when production groupingExpression is entered.
func (s *BaseMySQLParserListener) EnterGroupingExpression(ctx *GroupingExpressionContext) {}

// ExitGroupingExpression is called when production groupingExpression is exited.
func (s *BaseMySQLParserListener) ExitGroupingExpression(ctx *GroupingExpressionContext) {}

// EnterChannel is called when production channel is entered.
func (s *BaseMySQLParserListener) EnterChannel(ctx *ChannelContext) {}

// ExitChannel is called when production channel is exited.
func (s *BaseMySQLParserListener) ExitChannel(ctx *ChannelContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseMySQLParserListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseMySQLParserListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseMySQLParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseMySQLParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseMySQLParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseMySQLParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfBody is called when production ifBody is entered.
func (s *BaseMySQLParserListener) EnterIfBody(ctx *IfBodyContext) {}

// ExitIfBody is called when production ifBody is exited.
func (s *BaseMySQLParserListener) ExitIfBody(ctx *IfBodyContext) {}

// EnterThenStatement is called when production thenStatement is entered.
func (s *BaseMySQLParserListener) EnterThenStatement(ctx *ThenStatementContext) {}

// ExitThenStatement is called when production thenStatement is exited.
func (s *BaseMySQLParserListener) ExitThenStatement(ctx *ThenStatementContext) {}

// EnterCompoundStatementList is called when production compoundStatementList is entered.
func (s *BaseMySQLParserListener) EnterCompoundStatementList(ctx *CompoundStatementListContext) {}

// ExitCompoundStatementList is called when production compoundStatementList is exited.
func (s *BaseMySQLParserListener) ExitCompoundStatementList(ctx *CompoundStatementListContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseMySQLParserListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseMySQLParserListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BaseMySQLParserListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BaseMySQLParserListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterLabeledBlock is called when production labeledBlock is entered.
func (s *BaseMySQLParserListener) EnterLabeledBlock(ctx *LabeledBlockContext) {}

// ExitLabeledBlock is called when production labeledBlock is exited.
func (s *BaseMySQLParserListener) ExitLabeledBlock(ctx *LabeledBlockContext) {}

// EnterUnlabeledBlock is called when production unlabeledBlock is entered.
func (s *BaseMySQLParserListener) EnterUnlabeledBlock(ctx *UnlabeledBlockContext) {}

// ExitUnlabeledBlock is called when production unlabeledBlock is exited.
func (s *BaseMySQLParserListener) ExitUnlabeledBlock(ctx *UnlabeledBlockContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseMySQLParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseMySQLParserListener) ExitLabel(ctx *LabelContext) {}

// EnterBeginEndBlock is called when production beginEndBlock is entered.
func (s *BaseMySQLParserListener) EnterBeginEndBlock(ctx *BeginEndBlockContext) {}

// ExitBeginEndBlock is called when production beginEndBlock is exited.
func (s *BaseMySQLParserListener) ExitBeginEndBlock(ctx *BeginEndBlockContext) {}

// EnterLabeledControl is called when production labeledControl is entered.
func (s *BaseMySQLParserListener) EnterLabeledControl(ctx *LabeledControlContext) {}

// ExitLabeledControl is called when production labeledControl is exited.
func (s *BaseMySQLParserListener) ExitLabeledControl(ctx *LabeledControlContext) {}

// EnterUnlabeledControl is called when production unlabeledControl is entered.
func (s *BaseMySQLParserListener) EnterUnlabeledControl(ctx *UnlabeledControlContext) {}

// ExitUnlabeledControl is called when production unlabeledControl is exited.
func (s *BaseMySQLParserListener) ExitUnlabeledControl(ctx *UnlabeledControlContext) {}

// EnterLoopBlock is called when production loopBlock is entered.
func (s *BaseMySQLParserListener) EnterLoopBlock(ctx *LoopBlockContext) {}

// ExitLoopBlock is called when production loopBlock is exited.
func (s *BaseMySQLParserListener) ExitLoopBlock(ctx *LoopBlockContext) {}

// EnterWhileDoBlock is called when production whileDoBlock is entered.
func (s *BaseMySQLParserListener) EnterWhileDoBlock(ctx *WhileDoBlockContext) {}

// ExitWhileDoBlock is called when production whileDoBlock is exited.
func (s *BaseMySQLParserListener) ExitWhileDoBlock(ctx *WhileDoBlockContext) {}

// EnterRepeatUntilBlock is called when production repeatUntilBlock is entered.
func (s *BaseMySQLParserListener) EnterRepeatUntilBlock(ctx *RepeatUntilBlockContext) {}

// ExitRepeatUntilBlock is called when production repeatUntilBlock is exited.
func (s *BaseMySQLParserListener) ExitRepeatUntilBlock(ctx *RepeatUntilBlockContext) {}

// EnterSpDeclarations is called when production spDeclarations is entered.
func (s *BaseMySQLParserListener) EnterSpDeclarations(ctx *SpDeclarationsContext) {}

// ExitSpDeclarations is called when production spDeclarations is exited.
func (s *BaseMySQLParserListener) ExitSpDeclarations(ctx *SpDeclarationsContext) {}

// EnterSpDeclaration is called when production spDeclaration is entered.
func (s *BaseMySQLParserListener) EnterSpDeclaration(ctx *SpDeclarationContext) {}

// ExitSpDeclaration is called when production spDeclaration is exited.
func (s *BaseMySQLParserListener) ExitSpDeclaration(ctx *SpDeclarationContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseMySQLParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseMySQLParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterConditionDeclaration is called when production conditionDeclaration is entered.
func (s *BaseMySQLParserListener) EnterConditionDeclaration(ctx *ConditionDeclarationContext) {}

// ExitConditionDeclaration is called when production conditionDeclaration is exited.
func (s *BaseMySQLParserListener) ExitConditionDeclaration(ctx *ConditionDeclarationContext) {}

// EnterSpCondition is called when production spCondition is entered.
func (s *BaseMySQLParserListener) EnterSpCondition(ctx *SpConditionContext) {}

// ExitSpCondition is called when production spCondition is exited.
func (s *BaseMySQLParserListener) ExitSpCondition(ctx *SpConditionContext) {}

// EnterSqlstate is called when production sqlstate is entered.
func (s *BaseMySQLParserListener) EnterSqlstate(ctx *SqlstateContext) {}

// ExitSqlstate is called when production sqlstate is exited.
func (s *BaseMySQLParserListener) ExitSqlstate(ctx *SqlstateContext) {}

// EnterHandlerDeclaration is called when production handlerDeclaration is entered.
func (s *BaseMySQLParserListener) EnterHandlerDeclaration(ctx *HandlerDeclarationContext) {}

// ExitHandlerDeclaration is called when production handlerDeclaration is exited.
func (s *BaseMySQLParserListener) ExitHandlerDeclaration(ctx *HandlerDeclarationContext) {}

// EnterHandlerCondition is called when production handlerCondition is entered.
func (s *BaseMySQLParserListener) EnterHandlerCondition(ctx *HandlerConditionContext) {}

// ExitHandlerCondition is called when production handlerCondition is exited.
func (s *BaseMySQLParserListener) ExitHandlerCondition(ctx *HandlerConditionContext) {}

// EnterCursorDeclaration is called when production cursorDeclaration is entered.
func (s *BaseMySQLParserListener) EnterCursorDeclaration(ctx *CursorDeclarationContext) {}

// ExitCursorDeclaration is called when production cursorDeclaration is exited.
func (s *BaseMySQLParserListener) ExitCursorDeclaration(ctx *CursorDeclarationContext) {}

// EnterIterateStatement is called when production iterateStatement is entered.
func (s *BaseMySQLParserListener) EnterIterateStatement(ctx *IterateStatementContext) {}

// ExitIterateStatement is called when production iterateStatement is exited.
func (s *BaseMySQLParserListener) ExitIterateStatement(ctx *IterateStatementContext) {}

// EnterLeaveStatement is called when production leaveStatement is entered.
func (s *BaseMySQLParserListener) EnterLeaveStatement(ctx *LeaveStatementContext) {}

// ExitLeaveStatement is called when production leaveStatement is exited.
func (s *BaseMySQLParserListener) ExitLeaveStatement(ctx *LeaveStatementContext) {}

// EnterGetDiagnostics is called when production getDiagnostics is entered.
func (s *BaseMySQLParserListener) EnterGetDiagnostics(ctx *GetDiagnosticsContext) {}

// ExitGetDiagnostics is called when production getDiagnostics is exited.
func (s *BaseMySQLParserListener) ExitGetDiagnostics(ctx *GetDiagnosticsContext) {}

// EnterSignalAllowedExpr is called when production signalAllowedExpr is entered.
func (s *BaseMySQLParserListener) EnterSignalAllowedExpr(ctx *SignalAllowedExprContext) {}

// ExitSignalAllowedExpr is called when production signalAllowedExpr is exited.
func (s *BaseMySQLParserListener) ExitSignalAllowedExpr(ctx *SignalAllowedExprContext) {}

// EnterStatementInformationItem is called when production statementInformationItem is entered.
func (s *BaseMySQLParserListener) EnterStatementInformationItem(ctx *StatementInformationItemContext) {
}

// ExitStatementInformationItem is called when production statementInformationItem is exited.
func (s *BaseMySQLParserListener) ExitStatementInformationItem(ctx *StatementInformationItemContext) {
}

// EnterConditionInformationItem is called when production conditionInformationItem is entered.
func (s *BaseMySQLParserListener) EnterConditionInformationItem(ctx *ConditionInformationItemContext) {
}

// ExitConditionInformationItem is called when production conditionInformationItem is exited.
func (s *BaseMySQLParserListener) ExitConditionInformationItem(ctx *ConditionInformationItemContext) {
}

// EnterSignalInformationItemName is called when production signalInformationItemName is entered.
func (s *BaseMySQLParserListener) EnterSignalInformationItemName(ctx *SignalInformationItemNameContext) {
}

// ExitSignalInformationItemName is called when production signalInformationItemName is exited.
func (s *BaseMySQLParserListener) ExitSignalInformationItemName(ctx *SignalInformationItemNameContext) {
}

// EnterSignalStatement is called when production signalStatement is entered.
func (s *BaseMySQLParserListener) EnterSignalStatement(ctx *SignalStatementContext) {}

// ExitSignalStatement is called when production signalStatement is exited.
func (s *BaseMySQLParserListener) ExitSignalStatement(ctx *SignalStatementContext) {}

// EnterResignalStatement is called when production resignalStatement is entered.
func (s *BaseMySQLParserListener) EnterResignalStatement(ctx *ResignalStatementContext) {}

// ExitResignalStatement is called when production resignalStatement is exited.
func (s *BaseMySQLParserListener) ExitResignalStatement(ctx *ResignalStatementContext) {}

// EnterSignalInformationItem is called when production signalInformationItem is entered.
func (s *BaseMySQLParserListener) EnterSignalInformationItem(ctx *SignalInformationItemContext) {}

// ExitSignalInformationItem is called when production signalInformationItem is exited.
func (s *BaseMySQLParserListener) ExitSignalInformationItem(ctx *SignalInformationItemContext) {}

// EnterCursorOpen is called when production cursorOpen is entered.
func (s *BaseMySQLParserListener) EnterCursorOpen(ctx *CursorOpenContext) {}

// ExitCursorOpen is called when production cursorOpen is exited.
func (s *BaseMySQLParserListener) ExitCursorOpen(ctx *CursorOpenContext) {}

// EnterCursorClose is called when production cursorClose is entered.
func (s *BaseMySQLParserListener) EnterCursorClose(ctx *CursorCloseContext) {}

// ExitCursorClose is called when production cursorClose is exited.
func (s *BaseMySQLParserListener) ExitCursorClose(ctx *CursorCloseContext) {}

// EnterCursorFetch is called when production cursorFetch is entered.
func (s *BaseMySQLParserListener) EnterCursorFetch(ctx *CursorFetchContext) {}

// ExitCursorFetch is called when production cursorFetch is exited.
func (s *BaseMySQLParserListener) ExitCursorFetch(ctx *CursorFetchContext) {}

// EnterSchedule is called when production schedule is entered.
func (s *BaseMySQLParserListener) EnterSchedule(ctx *ScheduleContext) {}

// ExitSchedule is called when production schedule is exited.
func (s *BaseMySQLParserListener) ExitSchedule(ctx *ScheduleContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseMySQLParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseMySQLParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterCheckOrReferences is called when production checkOrReferences is entered.
func (s *BaseMySQLParserListener) EnterCheckOrReferences(ctx *CheckOrReferencesContext) {}

// ExitCheckOrReferences is called when production checkOrReferences is exited.
func (s *BaseMySQLParserListener) ExitCheckOrReferences(ctx *CheckOrReferencesContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseMySQLParserListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseMySQLParserListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterConstraintEnforcement is called when production constraintEnforcement is entered.
func (s *BaseMySQLParserListener) EnterConstraintEnforcement(ctx *ConstraintEnforcementContext) {}

// ExitConstraintEnforcement is called when production constraintEnforcement is exited.
func (s *BaseMySQLParserListener) ExitConstraintEnforcement(ctx *ConstraintEnforcementContext) {}

// EnterTableConstraintDef is called when production tableConstraintDef is entered.
func (s *BaseMySQLParserListener) EnterTableConstraintDef(ctx *TableConstraintDefContext) {}

// ExitTableConstraintDef is called when production tableConstraintDef is exited.
func (s *BaseMySQLParserListener) ExitTableConstraintDef(ctx *TableConstraintDefContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseMySQLParserListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseMySQLParserListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterFieldDefinition is called when production fieldDefinition is entered.
func (s *BaseMySQLParserListener) EnterFieldDefinition(ctx *FieldDefinitionContext) {}

// ExitFieldDefinition is called when production fieldDefinition is exited.
func (s *BaseMySQLParserListener) ExitFieldDefinition(ctx *FieldDefinitionContext) {}

// EnterColumnAttribute is called when production columnAttribute is entered.
func (s *BaseMySQLParserListener) EnterColumnAttribute(ctx *ColumnAttributeContext) {}

// ExitColumnAttribute is called when production columnAttribute is exited.
func (s *BaseMySQLParserListener) ExitColumnAttribute(ctx *ColumnAttributeContext) {}

// EnterColumnFormat is called when production columnFormat is entered.
func (s *BaseMySQLParserListener) EnterColumnFormat(ctx *ColumnFormatContext) {}

// ExitColumnFormat is called when production columnFormat is exited.
func (s *BaseMySQLParserListener) ExitColumnFormat(ctx *ColumnFormatContext) {}

// EnterStorageMedia is called when production storageMedia is entered.
func (s *BaseMySQLParserListener) EnterStorageMedia(ctx *StorageMediaContext) {}

// ExitStorageMedia is called when production storageMedia is exited.
func (s *BaseMySQLParserListener) ExitStorageMedia(ctx *StorageMediaContext) {}

// EnterGcolAttribute is called when production gcolAttribute is entered.
func (s *BaseMySQLParserListener) EnterGcolAttribute(ctx *GcolAttributeContext) {}

// ExitGcolAttribute is called when production gcolAttribute is exited.
func (s *BaseMySQLParserListener) ExitGcolAttribute(ctx *GcolAttributeContext) {}

// EnterReferences is called when production references is entered.
func (s *BaseMySQLParserListener) EnterReferences(ctx *ReferencesContext) {}

// ExitReferences is called when production references is exited.
func (s *BaseMySQLParserListener) ExitReferences(ctx *ReferencesContext) {}

// EnterDeleteOption is called when production deleteOption is entered.
func (s *BaseMySQLParserListener) EnterDeleteOption(ctx *DeleteOptionContext) {}

// ExitDeleteOption is called when production deleteOption is exited.
func (s *BaseMySQLParserListener) ExitDeleteOption(ctx *DeleteOptionContext) {}

// EnterKeyList is called when production keyList is entered.
func (s *BaseMySQLParserListener) EnterKeyList(ctx *KeyListContext) {}

// ExitKeyList is called when production keyList is exited.
func (s *BaseMySQLParserListener) ExitKeyList(ctx *KeyListContext) {}

// EnterKeyPart is called when production keyPart is entered.
func (s *BaseMySQLParserListener) EnterKeyPart(ctx *KeyPartContext) {}

// ExitKeyPart is called when production keyPart is exited.
func (s *BaseMySQLParserListener) ExitKeyPart(ctx *KeyPartContext) {}

// EnterKeyListWithExpression is called when production keyListWithExpression is entered.
func (s *BaseMySQLParserListener) EnterKeyListWithExpression(ctx *KeyListWithExpressionContext) {}

// ExitKeyListWithExpression is called when production keyListWithExpression is exited.
func (s *BaseMySQLParserListener) ExitKeyListWithExpression(ctx *KeyListWithExpressionContext) {}

// EnterKeyPartOrExpression is called when production keyPartOrExpression is entered.
func (s *BaseMySQLParserListener) EnterKeyPartOrExpression(ctx *KeyPartOrExpressionContext) {}

// ExitKeyPartOrExpression is called when production keyPartOrExpression is exited.
func (s *BaseMySQLParserListener) ExitKeyPartOrExpression(ctx *KeyPartOrExpressionContext) {}

// EnterKeyListVariants is called when production keyListVariants is entered.
func (s *BaseMySQLParserListener) EnterKeyListVariants(ctx *KeyListVariantsContext) {}

// ExitKeyListVariants is called when production keyListVariants is exited.
func (s *BaseMySQLParserListener) ExitKeyListVariants(ctx *KeyListVariantsContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseMySQLParserListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseMySQLParserListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseMySQLParserListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseMySQLParserListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterCommonIndexOption is called when production commonIndexOption is entered.
func (s *BaseMySQLParserListener) EnterCommonIndexOption(ctx *CommonIndexOptionContext) {}

// ExitCommonIndexOption is called when production commonIndexOption is exited.
func (s *BaseMySQLParserListener) ExitCommonIndexOption(ctx *CommonIndexOptionContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseMySQLParserListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseMySQLParserListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterIndexTypeClause is called when production indexTypeClause is entered.
func (s *BaseMySQLParserListener) EnterIndexTypeClause(ctx *IndexTypeClauseContext) {}

// ExitIndexTypeClause is called when production indexTypeClause is exited.
func (s *BaseMySQLParserListener) ExitIndexTypeClause(ctx *IndexTypeClauseContext) {}

// EnterFulltextIndexOption is called when production fulltextIndexOption is entered.
func (s *BaseMySQLParserListener) EnterFulltextIndexOption(ctx *FulltextIndexOptionContext) {}

// ExitFulltextIndexOption is called when production fulltextIndexOption is exited.
func (s *BaseMySQLParserListener) ExitFulltextIndexOption(ctx *FulltextIndexOptionContext) {}

// EnterSpatialIndexOption is called when production spatialIndexOption is entered.
func (s *BaseMySQLParserListener) EnterSpatialIndexOption(ctx *SpatialIndexOptionContext) {}

// ExitSpatialIndexOption is called when production spatialIndexOption is exited.
func (s *BaseMySQLParserListener) ExitSpatialIndexOption(ctx *SpatialIndexOptionContext) {}

// EnterDataTypeDefinition is called when production dataTypeDefinition is entered.
func (s *BaseMySQLParserListener) EnterDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// ExitDataTypeDefinition is called when production dataTypeDefinition is exited.
func (s *BaseMySQLParserListener) ExitDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseMySQLParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseMySQLParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterNchar is called when production nchar is entered.
func (s *BaseMySQLParserListener) EnterNchar(ctx *NcharContext) {}

// ExitNchar is called when production nchar is exited.
func (s *BaseMySQLParserListener) ExitNchar(ctx *NcharContext) {}

// EnterRealType is called when production realType is entered.
func (s *BaseMySQLParserListener) EnterRealType(ctx *RealTypeContext) {}

// ExitRealType is called when production realType is exited.
func (s *BaseMySQLParserListener) ExitRealType(ctx *RealTypeContext) {}

// EnterFieldLength is called when production fieldLength is entered.
func (s *BaseMySQLParserListener) EnterFieldLength(ctx *FieldLengthContext) {}

// ExitFieldLength is called when production fieldLength is exited.
func (s *BaseMySQLParserListener) ExitFieldLength(ctx *FieldLengthContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseMySQLParserListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseMySQLParserListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterCharsetWithOptBinary is called when production charsetWithOptBinary is entered.
func (s *BaseMySQLParserListener) EnterCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) {}

// ExitCharsetWithOptBinary is called when production charsetWithOptBinary is exited.
func (s *BaseMySQLParserListener) ExitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) {}

// EnterAscii is called when production ascii is entered.
func (s *BaseMySQLParserListener) EnterAscii(ctx *AsciiContext) {}

// ExitAscii is called when production ascii is exited.
func (s *BaseMySQLParserListener) ExitAscii(ctx *AsciiContext) {}

// EnterUnicode is called when production unicode is entered.
func (s *BaseMySQLParserListener) EnterUnicode(ctx *UnicodeContext) {}

// ExitUnicode is called when production unicode is exited.
func (s *BaseMySQLParserListener) ExitUnicode(ctx *UnicodeContext) {}

// EnterWsNumCodepoints is called when production wsNumCodepoints is entered.
func (s *BaseMySQLParserListener) EnterWsNumCodepoints(ctx *WsNumCodepointsContext) {}

// ExitWsNumCodepoints is called when production wsNumCodepoints is exited.
func (s *BaseMySQLParserListener) ExitWsNumCodepoints(ctx *WsNumCodepointsContext) {}

// EnterTypeDatetimePrecision is called when production typeDatetimePrecision is entered.
func (s *BaseMySQLParserListener) EnterTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) {}

// ExitTypeDatetimePrecision is called when production typeDatetimePrecision is exited.
func (s *BaseMySQLParserListener) ExitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseMySQLParserListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseMySQLParserListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterCollationName is called when production collationName is entered.
func (s *BaseMySQLParserListener) EnterCollationName(ctx *CollationNameContext) {}

// ExitCollationName is called when production collationName is exited.
func (s *BaseMySQLParserListener) ExitCollationName(ctx *CollationNameContext) {}

// EnterCreateTableOptions is called when production createTableOptions is entered.
func (s *BaseMySQLParserListener) EnterCreateTableOptions(ctx *CreateTableOptionsContext) {}

// ExitCreateTableOptions is called when production createTableOptions is exited.
func (s *BaseMySQLParserListener) ExitCreateTableOptions(ctx *CreateTableOptionsContext) {}

// EnterCreateTableOptionsSpaceSeparated is called when production createTableOptionsSpaceSeparated is entered.
func (s *BaseMySQLParserListener) EnterCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) {
}

// ExitCreateTableOptionsSpaceSeparated is called when production createTableOptionsSpaceSeparated is exited.
func (s *BaseMySQLParserListener) ExitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) {
}

// EnterCreateTableOption is called when production createTableOption is entered.
func (s *BaseMySQLParserListener) EnterCreateTableOption(ctx *CreateTableOptionContext) {}

// ExitCreateTableOption is called when production createTableOption is exited.
func (s *BaseMySQLParserListener) ExitCreateTableOption(ctx *CreateTableOptionContext) {}

// EnterTernaryOption is called when production ternaryOption is entered.
func (s *BaseMySQLParserListener) EnterTernaryOption(ctx *TernaryOptionContext) {}

// ExitTernaryOption is called when production ternaryOption is exited.
func (s *BaseMySQLParserListener) ExitTernaryOption(ctx *TernaryOptionContext) {}

// EnterDefaultCollation is called when production defaultCollation is entered.
func (s *BaseMySQLParserListener) EnterDefaultCollation(ctx *DefaultCollationContext) {}

// ExitDefaultCollation is called when production defaultCollation is exited.
func (s *BaseMySQLParserListener) ExitDefaultCollation(ctx *DefaultCollationContext) {}

// EnterDefaultEncryption is called when production defaultEncryption is entered.
func (s *BaseMySQLParserListener) EnterDefaultEncryption(ctx *DefaultEncryptionContext) {}

// ExitDefaultEncryption is called when production defaultEncryption is exited.
func (s *BaseMySQLParserListener) ExitDefaultEncryption(ctx *DefaultEncryptionContext) {}

// EnterDefaultCharset is called when production defaultCharset is entered.
func (s *BaseMySQLParserListener) EnterDefaultCharset(ctx *DefaultCharsetContext) {}

// ExitDefaultCharset is called when production defaultCharset is exited.
func (s *BaseMySQLParserListener) ExitDefaultCharset(ctx *DefaultCharsetContext) {}

// EnterPartitionClause is called when production partitionClause is entered.
func (s *BaseMySQLParserListener) EnterPartitionClause(ctx *PartitionClauseContext) {}

// ExitPartitionClause is called when production partitionClause is exited.
func (s *BaseMySQLParserListener) ExitPartitionClause(ctx *PartitionClauseContext) {}

// EnterPartitionDefKey is called when production partitionDefKey is entered.
func (s *BaseMySQLParserListener) EnterPartitionDefKey(ctx *PartitionDefKeyContext) {}

// ExitPartitionDefKey is called when production partitionDefKey is exited.
func (s *BaseMySQLParserListener) ExitPartitionDefKey(ctx *PartitionDefKeyContext) {}

// EnterPartitionDefHash is called when production partitionDefHash is entered.
func (s *BaseMySQLParserListener) EnterPartitionDefHash(ctx *PartitionDefHashContext) {}

// ExitPartitionDefHash is called when production partitionDefHash is exited.
func (s *BaseMySQLParserListener) ExitPartitionDefHash(ctx *PartitionDefHashContext) {}

// EnterPartitionDefRangeList is called when production partitionDefRangeList is entered.
func (s *BaseMySQLParserListener) EnterPartitionDefRangeList(ctx *PartitionDefRangeListContext) {}

// ExitPartitionDefRangeList is called when production partitionDefRangeList is exited.
func (s *BaseMySQLParserListener) ExitPartitionDefRangeList(ctx *PartitionDefRangeListContext) {}

// EnterSubPartitions is called when production subPartitions is entered.
func (s *BaseMySQLParserListener) EnterSubPartitions(ctx *SubPartitionsContext) {}

// ExitSubPartitions is called when production subPartitions is exited.
func (s *BaseMySQLParserListener) ExitSubPartitions(ctx *SubPartitionsContext) {}

// EnterPartitionKeyAlgorithm is called when production partitionKeyAlgorithm is entered.
func (s *BaseMySQLParserListener) EnterPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) {}

// ExitPartitionKeyAlgorithm is called when production partitionKeyAlgorithm is exited.
func (s *BaseMySQLParserListener) ExitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) {}

// EnterPartitionDefinitions is called when production partitionDefinitions is entered.
func (s *BaseMySQLParserListener) EnterPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// ExitPartitionDefinitions is called when production partitionDefinitions is exited.
func (s *BaseMySQLParserListener) ExitPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// EnterPartitionDefinition is called when production partitionDefinition is entered.
func (s *BaseMySQLParserListener) EnterPartitionDefinition(ctx *PartitionDefinitionContext) {}

// ExitPartitionDefinition is called when production partitionDefinition is exited.
func (s *BaseMySQLParserListener) ExitPartitionDefinition(ctx *PartitionDefinitionContext) {}

// EnterPartitionValuesIn is called when production partitionValuesIn is entered.
func (s *BaseMySQLParserListener) EnterPartitionValuesIn(ctx *PartitionValuesInContext) {}

// ExitPartitionValuesIn is called when production partitionValuesIn is exited.
func (s *BaseMySQLParserListener) ExitPartitionValuesIn(ctx *PartitionValuesInContext) {}

// EnterPartitionOption is called when production partitionOption is entered.
func (s *BaseMySQLParserListener) EnterPartitionOption(ctx *PartitionOptionContext) {}

// ExitPartitionOption is called when production partitionOption is exited.
func (s *BaseMySQLParserListener) ExitPartitionOption(ctx *PartitionOptionContext) {}

// EnterSubpartitionDefinition is called when production subpartitionDefinition is entered.
func (s *BaseMySQLParserListener) EnterSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {}

// ExitSubpartitionDefinition is called when production subpartitionDefinition is exited.
func (s *BaseMySQLParserListener) ExitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {}

// EnterPartitionValueItemListParen is called when production partitionValueItemListParen is entered.
func (s *BaseMySQLParserListener) EnterPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) {
}

// ExitPartitionValueItemListParen is called when production partitionValueItemListParen is exited.
func (s *BaseMySQLParserListener) ExitPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) {
}

// EnterPartitionValueItem is called when production partitionValueItem is entered.
func (s *BaseMySQLParserListener) EnterPartitionValueItem(ctx *PartitionValueItemContext) {}

// ExitPartitionValueItem is called when production partitionValueItem is exited.
func (s *BaseMySQLParserListener) ExitPartitionValueItem(ctx *PartitionValueItemContext) {}

// EnterDefinerClause is called when production definerClause is entered.
func (s *BaseMySQLParserListener) EnterDefinerClause(ctx *DefinerClauseContext) {}

// ExitDefinerClause is called when production definerClause is exited.
func (s *BaseMySQLParserListener) ExitDefinerClause(ctx *DefinerClauseContext) {}

// EnterIfExists is called when production ifExists is entered.
func (s *BaseMySQLParserListener) EnterIfExists(ctx *IfExistsContext) {}

// ExitIfExists is called when production ifExists is exited.
func (s *BaseMySQLParserListener) ExitIfExists(ctx *IfExistsContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseMySQLParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseMySQLParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterProcedureParameter is called when production procedureParameter is entered.
func (s *BaseMySQLParserListener) EnterProcedureParameter(ctx *ProcedureParameterContext) {}

// ExitProcedureParameter is called when production procedureParameter is exited.
func (s *BaseMySQLParserListener) ExitProcedureParameter(ctx *ProcedureParameterContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BaseMySQLParserListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BaseMySQLParserListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterCollate is called when production collate is entered.
func (s *BaseMySQLParserListener) EnterCollate(ctx *CollateContext) {}

// ExitCollate is called when production collate is exited.
func (s *BaseMySQLParserListener) ExitCollate(ctx *CollateContext) {}

// EnterTypeWithOptCollate is called when production typeWithOptCollate is entered.
func (s *BaseMySQLParserListener) EnterTypeWithOptCollate(ctx *TypeWithOptCollateContext) {}

// ExitTypeWithOptCollate is called when production typeWithOptCollate is exited.
func (s *BaseMySQLParserListener) ExitTypeWithOptCollate(ctx *TypeWithOptCollateContext) {}

// EnterSchemaIdentifierPair is called when production schemaIdentifierPair is entered.
func (s *BaseMySQLParserListener) EnterSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) {}

// ExitSchemaIdentifierPair is called when production schemaIdentifierPair is exited.
func (s *BaseMySQLParserListener) ExitSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) {}

// EnterViewRefList is called when production viewRefList is entered.
func (s *BaseMySQLParserListener) EnterViewRefList(ctx *ViewRefListContext) {}

// ExitViewRefList is called when production viewRefList is exited.
func (s *BaseMySQLParserListener) ExitViewRefList(ctx *ViewRefListContext) {}

// EnterUpdateList is called when production updateList is entered.
func (s *BaseMySQLParserListener) EnterUpdateList(ctx *UpdateListContext) {}

// ExitUpdateList is called when production updateList is exited.
func (s *BaseMySQLParserListener) ExitUpdateList(ctx *UpdateListContext) {}

// EnterUpdateElement is called when production updateElement is entered.
func (s *BaseMySQLParserListener) EnterUpdateElement(ctx *UpdateElementContext) {}

// ExitUpdateElement is called when production updateElement is exited.
func (s *BaseMySQLParserListener) ExitUpdateElement(ctx *UpdateElementContext) {}

// EnterCharsetClause is called when production charsetClause is entered.
func (s *BaseMySQLParserListener) EnterCharsetClause(ctx *CharsetClauseContext) {}

// ExitCharsetClause is called when production charsetClause is exited.
func (s *BaseMySQLParserListener) ExitCharsetClause(ctx *CharsetClauseContext) {}

// EnterFieldsClause is called when production fieldsClause is entered.
func (s *BaseMySQLParserListener) EnterFieldsClause(ctx *FieldsClauseContext) {}

// ExitFieldsClause is called when production fieldsClause is exited.
func (s *BaseMySQLParserListener) ExitFieldsClause(ctx *FieldsClauseContext) {}

// EnterFieldTerm is called when production fieldTerm is entered.
func (s *BaseMySQLParserListener) EnterFieldTerm(ctx *FieldTermContext) {}

// ExitFieldTerm is called when production fieldTerm is exited.
func (s *BaseMySQLParserListener) ExitFieldTerm(ctx *FieldTermContext) {}

// EnterLinesClause is called when production linesClause is entered.
func (s *BaseMySQLParserListener) EnterLinesClause(ctx *LinesClauseContext) {}

// ExitLinesClause is called when production linesClause is exited.
func (s *BaseMySQLParserListener) ExitLinesClause(ctx *LinesClauseContext) {}

// EnterLineTerm is called when production lineTerm is entered.
func (s *BaseMySQLParserListener) EnterLineTerm(ctx *LineTermContext) {}

// ExitLineTerm is called when production lineTerm is exited.
func (s *BaseMySQLParserListener) ExitLineTerm(ctx *LineTermContext) {}

// EnterUserList is called when production userList is entered.
func (s *BaseMySQLParserListener) EnterUserList(ctx *UserListContext) {}

// ExitUserList is called when production userList is exited.
func (s *BaseMySQLParserListener) ExitUserList(ctx *UserListContext) {}

// EnterCreateUserList is called when production createUserList is entered.
func (s *BaseMySQLParserListener) EnterCreateUserList(ctx *CreateUserListContext) {}

// ExitCreateUserList is called when production createUserList is exited.
func (s *BaseMySQLParserListener) ExitCreateUserList(ctx *CreateUserListContext) {}

// EnterAlterUserList is called when production alterUserList is entered.
func (s *BaseMySQLParserListener) EnterAlterUserList(ctx *AlterUserListContext) {}

// ExitAlterUserList is called when production alterUserList is exited.
func (s *BaseMySQLParserListener) ExitAlterUserList(ctx *AlterUserListContext) {}

// EnterCreateUserEntry is called when production createUserEntry is entered.
func (s *BaseMySQLParserListener) EnterCreateUserEntry(ctx *CreateUserEntryContext) {}

// ExitCreateUserEntry is called when production createUserEntry is exited.
func (s *BaseMySQLParserListener) ExitCreateUserEntry(ctx *CreateUserEntryContext) {}

// EnterAlterUserEntry is called when production alterUserEntry is entered.
func (s *BaseMySQLParserListener) EnterAlterUserEntry(ctx *AlterUserEntryContext) {}

// ExitAlterUserEntry is called when production alterUserEntry is exited.
func (s *BaseMySQLParserListener) ExitAlterUserEntry(ctx *AlterUserEntryContext) {}

// EnterRetainCurrentPassword is called when production retainCurrentPassword is entered.
func (s *BaseMySQLParserListener) EnterRetainCurrentPassword(ctx *RetainCurrentPasswordContext) {}

// ExitRetainCurrentPassword is called when production retainCurrentPassword is exited.
func (s *BaseMySQLParserListener) ExitRetainCurrentPassword(ctx *RetainCurrentPasswordContext) {}

// EnterDiscardOldPassword is called when production discardOldPassword is entered.
func (s *BaseMySQLParserListener) EnterDiscardOldPassword(ctx *DiscardOldPasswordContext) {}

// ExitDiscardOldPassword is called when production discardOldPassword is exited.
func (s *BaseMySQLParserListener) ExitDiscardOldPassword(ctx *DiscardOldPasswordContext) {}

// EnterReplacePassword is called when production replacePassword is entered.
func (s *BaseMySQLParserListener) EnterReplacePassword(ctx *ReplacePasswordContext) {}

// ExitReplacePassword is called when production replacePassword is exited.
func (s *BaseMySQLParserListener) ExitReplacePassword(ctx *ReplacePasswordContext) {}

// EnterUserIdentifierOrText is called when production userIdentifierOrText is entered.
func (s *BaseMySQLParserListener) EnterUserIdentifierOrText(ctx *UserIdentifierOrTextContext) {}

// ExitUserIdentifierOrText is called when production userIdentifierOrText is exited.
func (s *BaseMySQLParserListener) ExitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) {}

// EnterUser is called when production user is entered.
func (s *BaseMySQLParserListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseMySQLParserListener) ExitUser(ctx *UserContext) {}

// EnterLikeClause is called when production likeClause is entered.
func (s *BaseMySQLParserListener) EnterLikeClause(ctx *LikeClauseContext) {}

// ExitLikeClause is called when production likeClause is exited.
func (s *BaseMySQLParserListener) ExitLikeClause(ctx *LikeClauseContext) {}

// EnterLikeOrWhere is called when production likeOrWhere is entered.
func (s *BaseMySQLParserListener) EnterLikeOrWhere(ctx *LikeOrWhereContext) {}

// ExitLikeOrWhere is called when production likeOrWhere is exited.
func (s *BaseMySQLParserListener) ExitLikeOrWhere(ctx *LikeOrWhereContext) {}

// EnterOnlineOption is called when production onlineOption is entered.
func (s *BaseMySQLParserListener) EnterOnlineOption(ctx *OnlineOptionContext) {}

// ExitOnlineOption is called when production onlineOption is exited.
func (s *BaseMySQLParserListener) ExitOnlineOption(ctx *OnlineOptionContext) {}

// EnterNoWriteToBinLog is called when production noWriteToBinLog is entered.
func (s *BaseMySQLParserListener) EnterNoWriteToBinLog(ctx *NoWriteToBinLogContext) {}

// ExitNoWriteToBinLog is called when production noWriteToBinLog is exited.
func (s *BaseMySQLParserListener) ExitNoWriteToBinLog(ctx *NoWriteToBinLogContext) {}

// EnterUsePartition is called when production usePartition is entered.
func (s *BaseMySQLParserListener) EnterUsePartition(ctx *UsePartitionContext) {}

// ExitUsePartition is called when production usePartition is exited.
func (s *BaseMySQLParserListener) ExitUsePartition(ctx *UsePartitionContext) {}

// EnterFieldIdentifier is called when production fieldIdentifier is entered.
func (s *BaseMySQLParserListener) EnterFieldIdentifier(ctx *FieldIdentifierContext) {}

// ExitFieldIdentifier is called when production fieldIdentifier is exited.
func (s *BaseMySQLParserListener) ExitFieldIdentifier(ctx *FieldIdentifierContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseMySQLParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseMySQLParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterColumnInternalRef is called when production columnInternalRef is entered.
func (s *BaseMySQLParserListener) EnterColumnInternalRef(ctx *ColumnInternalRefContext) {}

// ExitColumnInternalRef is called when production columnInternalRef is exited.
func (s *BaseMySQLParserListener) ExitColumnInternalRef(ctx *ColumnInternalRefContext) {}

// EnterColumnInternalRefList is called when production columnInternalRefList is entered.
func (s *BaseMySQLParserListener) EnterColumnInternalRefList(ctx *ColumnInternalRefListContext) {}

// ExitColumnInternalRefList is called when production columnInternalRefList is exited.
func (s *BaseMySQLParserListener) ExitColumnInternalRefList(ctx *ColumnInternalRefListContext) {}

// EnterColumnRef is called when production columnRef is entered.
func (s *BaseMySQLParserListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production columnRef is exited.
func (s *BaseMySQLParserListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterInsertIdentifier is called when production insertIdentifier is entered.
func (s *BaseMySQLParserListener) EnterInsertIdentifier(ctx *InsertIdentifierContext) {}

// ExitInsertIdentifier is called when production insertIdentifier is exited.
func (s *BaseMySQLParserListener) ExitInsertIdentifier(ctx *InsertIdentifierContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseMySQLParserListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseMySQLParserListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterIndexRef is called when production indexRef is entered.
func (s *BaseMySQLParserListener) EnterIndexRef(ctx *IndexRefContext) {}

// ExitIndexRef is called when production indexRef is exited.
func (s *BaseMySQLParserListener) ExitIndexRef(ctx *IndexRefContext) {}

// EnterTableWild is called when production tableWild is entered.
func (s *BaseMySQLParserListener) EnterTableWild(ctx *TableWildContext) {}

// ExitTableWild is called when production tableWild is exited.
func (s *BaseMySQLParserListener) ExitTableWild(ctx *TableWildContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseMySQLParserListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseMySQLParserListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterSchemaRef is called when production schemaRef is entered.
func (s *BaseMySQLParserListener) EnterSchemaRef(ctx *SchemaRefContext) {}

// ExitSchemaRef is called when production schemaRef is exited.
func (s *BaseMySQLParserListener) ExitSchemaRef(ctx *SchemaRefContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseMySQLParserListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseMySQLParserListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterProcedureRef is called when production procedureRef is entered.
func (s *BaseMySQLParserListener) EnterProcedureRef(ctx *ProcedureRefContext) {}

// ExitProcedureRef is called when production procedureRef is exited.
func (s *BaseMySQLParserListener) ExitProcedureRef(ctx *ProcedureRefContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseMySQLParserListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseMySQLParserListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterFunctionRef is called when production functionRef is entered.
func (s *BaseMySQLParserListener) EnterFunctionRef(ctx *FunctionRefContext) {}

// ExitFunctionRef is called when production functionRef is exited.
func (s *BaseMySQLParserListener) ExitFunctionRef(ctx *FunctionRefContext) {}

// EnterTriggerName is called when production triggerName is entered.
func (s *BaseMySQLParserListener) EnterTriggerName(ctx *TriggerNameContext) {}

// ExitTriggerName is called when production triggerName is exited.
func (s *BaseMySQLParserListener) ExitTriggerName(ctx *TriggerNameContext) {}

// EnterTriggerRef is called when production triggerRef is entered.
func (s *BaseMySQLParserListener) EnterTriggerRef(ctx *TriggerRefContext) {}

// ExitTriggerRef is called when production triggerRef is exited.
func (s *BaseMySQLParserListener) ExitTriggerRef(ctx *TriggerRefContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseMySQLParserListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseMySQLParserListener) ExitViewName(ctx *ViewNameContext) {}

// EnterViewRef is called when production viewRef is entered.
func (s *BaseMySQLParserListener) EnterViewRef(ctx *ViewRefContext) {}

// ExitViewRef is called when production viewRef is exited.
func (s *BaseMySQLParserListener) ExitViewRef(ctx *ViewRefContext) {}

// EnterTablespaceName is called when production tablespaceName is entered.
func (s *BaseMySQLParserListener) EnterTablespaceName(ctx *TablespaceNameContext) {}

// ExitTablespaceName is called when production tablespaceName is exited.
func (s *BaseMySQLParserListener) ExitTablespaceName(ctx *TablespaceNameContext) {}

// EnterTablespaceRef is called when production tablespaceRef is entered.
func (s *BaseMySQLParserListener) EnterTablespaceRef(ctx *TablespaceRefContext) {}

// ExitTablespaceRef is called when production tablespaceRef is exited.
func (s *BaseMySQLParserListener) ExitTablespaceRef(ctx *TablespaceRefContext) {}

// EnterLogfileGroupName is called when production logfileGroupName is entered.
func (s *BaseMySQLParserListener) EnterLogfileGroupName(ctx *LogfileGroupNameContext) {}

// ExitLogfileGroupName is called when production logfileGroupName is exited.
func (s *BaseMySQLParserListener) ExitLogfileGroupName(ctx *LogfileGroupNameContext) {}

// EnterLogfileGroupRef is called when production logfileGroupRef is entered.
func (s *BaseMySQLParserListener) EnterLogfileGroupRef(ctx *LogfileGroupRefContext) {}

// ExitLogfileGroupRef is called when production logfileGroupRef is exited.
func (s *BaseMySQLParserListener) ExitLogfileGroupRef(ctx *LogfileGroupRefContext) {}

// EnterEventName is called when production eventName is entered.
func (s *BaseMySQLParserListener) EnterEventName(ctx *EventNameContext) {}

// ExitEventName is called when production eventName is exited.
func (s *BaseMySQLParserListener) ExitEventName(ctx *EventNameContext) {}

// EnterEventRef is called when production eventRef is entered.
func (s *BaseMySQLParserListener) EnterEventRef(ctx *EventRefContext) {}

// ExitEventRef is called when production eventRef is exited.
func (s *BaseMySQLParserListener) ExitEventRef(ctx *EventRefContext) {}

// EnterUdfName is called when production udfName is entered.
func (s *BaseMySQLParserListener) EnterUdfName(ctx *UdfNameContext) {}

// ExitUdfName is called when production udfName is exited.
func (s *BaseMySQLParserListener) ExitUdfName(ctx *UdfNameContext) {}

// EnterServerName is called when production serverName is entered.
func (s *BaseMySQLParserListener) EnterServerName(ctx *ServerNameContext) {}

// ExitServerName is called when production serverName is exited.
func (s *BaseMySQLParserListener) ExitServerName(ctx *ServerNameContext) {}

// EnterServerRef is called when production serverRef is entered.
func (s *BaseMySQLParserListener) EnterServerRef(ctx *ServerRefContext) {}

// ExitServerRef is called when production serverRef is exited.
func (s *BaseMySQLParserListener) ExitServerRef(ctx *ServerRefContext) {}

// EnterEngineRef is called when production engineRef is entered.
func (s *BaseMySQLParserListener) EnterEngineRef(ctx *EngineRefContext) {}

// ExitEngineRef is called when production engineRef is exited.
func (s *BaseMySQLParserListener) ExitEngineRef(ctx *EngineRefContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseMySQLParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseMySQLParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterFilterTableRef is called when production filterTableRef is entered.
func (s *BaseMySQLParserListener) EnterFilterTableRef(ctx *FilterTableRefContext) {}

// ExitFilterTableRef is called when production filterTableRef is exited.
func (s *BaseMySQLParserListener) ExitFilterTableRef(ctx *FilterTableRefContext) {}

// EnterTableRefWithWildcard is called when production tableRefWithWildcard is entered.
func (s *BaseMySQLParserListener) EnterTableRefWithWildcard(ctx *TableRefWithWildcardContext) {}

// ExitTableRefWithWildcard is called when production tableRefWithWildcard is exited.
func (s *BaseMySQLParserListener) ExitTableRefWithWildcard(ctx *TableRefWithWildcardContext) {}

// EnterTableRef is called when production tableRef is entered.
func (s *BaseMySQLParserListener) EnterTableRef(ctx *TableRefContext) {}

// ExitTableRef is called when production tableRef is exited.
func (s *BaseMySQLParserListener) ExitTableRef(ctx *TableRefContext) {}

// EnterTableRefList is called when production tableRefList is entered.
func (s *BaseMySQLParserListener) EnterTableRefList(ctx *TableRefListContext) {}

// ExitTableRefList is called when production tableRefList is exited.
func (s *BaseMySQLParserListener) ExitTableRefList(ctx *TableRefListContext) {}

// EnterTableAliasRefList is called when production tableAliasRefList is entered.
func (s *BaseMySQLParserListener) EnterTableAliasRefList(ctx *TableAliasRefListContext) {}

// ExitTableAliasRefList is called when production tableAliasRefList is exited.
func (s *BaseMySQLParserListener) ExitTableAliasRefList(ctx *TableAliasRefListContext) {}

// EnterParameterName is called when production parameterName is entered.
func (s *BaseMySQLParserListener) EnterParameterName(ctx *ParameterNameContext) {}

// ExitParameterName is called when production parameterName is exited.
func (s *BaseMySQLParserListener) ExitParameterName(ctx *ParameterNameContext) {}

// EnterLabelIdentifier is called when production labelIdentifier is entered.
func (s *BaseMySQLParserListener) EnterLabelIdentifier(ctx *LabelIdentifierContext) {}

// ExitLabelIdentifier is called when production labelIdentifier is exited.
func (s *BaseMySQLParserListener) ExitLabelIdentifier(ctx *LabelIdentifierContext) {}

// EnterLabelRef is called when production labelRef is entered.
func (s *BaseMySQLParserListener) EnterLabelRef(ctx *LabelRefContext) {}

// ExitLabelRef is called when production labelRef is exited.
func (s *BaseMySQLParserListener) ExitLabelRef(ctx *LabelRefContext) {}

// EnterRoleIdentifier is called when production roleIdentifier is entered.
func (s *BaseMySQLParserListener) EnterRoleIdentifier(ctx *RoleIdentifierContext) {}

// ExitRoleIdentifier is called when production roleIdentifier is exited.
func (s *BaseMySQLParserListener) ExitRoleIdentifier(ctx *RoleIdentifierContext) {}

// EnterRoleRef is called when production roleRef is entered.
func (s *BaseMySQLParserListener) EnterRoleRef(ctx *RoleRefContext) {}

// ExitRoleRef is called when production roleRef is exited.
func (s *BaseMySQLParserListener) ExitRoleRef(ctx *RoleRefContext) {}

// EnterPluginRef is called when production pluginRef is entered.
func (s *BaseMySQLParserListener) EnterPluginRef(ctx *PluginRefContext) {}

// ExitPluginRef is called when production pluginRef is exited.
func (s *BaseMySQLParserListener) ExitPluginRef(ctx *PluginRefContext) {}

// EnterComponentRef is called when production componentRef is entered.
func (s *BaseMySQLParserListener) EnterComponentRef(ctx *ComponentRefContext) {}

// ExitComponentRef is called when production componentRef is exited.
func (s *BaseMySQLParserListener) ExitComponentRef(ctx *ComponentRefContext) {}

// EnterResourceGroupRef is called when production resourceGroupRef is entered.
func (s *BaseMySQLParserListener) EnterResourceGroupRef(ctx *ResourceGroupRefContext) {}

// ExitResourceGroupRef is called when production resourceGroupRef is exited.
func (s *BaseMySQLParserListener) ExitResourceGroupRef(ctx *ResourceGroupRefContext) {}

// EnterWindowName is called when production windowName is entered.
func (s *BaseMySQLParserListener) EnterWindowName(ctx *WindowNameContext) {}

// ExitWindowName is called when production windowName is exited.
func (s *BaseMySQLParserListener) ExitWindowName(ctx *WindowNameContext) {}

// EnterPureIdentifier is called when production pureIdentifier is entered.
func (s *BaseMySQLParserListener) EnterPureIdentifier(ctx *PureIdentifierContext) {}

// ExitPureIdentifier is called when production pureIdentifier is exited.
func (s *BaseMySQLParserListener) ExitPureIdentifier(ctx *PureIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMySQLParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMySQLParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseMySQLParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseMySQLParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierListWithParentheses is called when production identifierListWithParentheses is entered.
func (s *BaseMySQLParserListener) EnterIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) {
}

// ExitIdentifierListWithParentheses is called when production identifierListWithParentheses is exited.
func (s *BaseMySQLParserListener) ExitIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) {
}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseMySQLParserListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseMySQLParserListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterSimpleIdentifier is called when production simpleIdentifier is entered.
func (s *BaseMySQLParserListener) EnterSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// ExitSimpleIdentifier is called when production simpleIdentifier is exited.
func (s *BaseMySQLParserListener) ExitSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// EnterDotIdentifier is called when production dotIdentifier is entered.
func (s *BaseMySQLParserListener) EnterDotIdentifier(ctx *DotIdentifierContext) {}

// ExitDotIdentifier is called when production dotIdentifier is exited.
func (s *BaseMySQLParserListener) ExitDotIdentifier(ctx *DotIdentifierContext) {}

// EnterUlong_number is called when production ulong_number is entered.
func (s *BaseMySQLParserListener) EnterUlong_number(ctx *Ulong_numberContext) {}

// ExitUlong_number is called when production ulong_number is exited.
func (s *BaseMySQLParserListener) ExitUlong_number(ctx *Ulong_numberContext) {}

// EnterReal_ulong_number is called when production real_ulong_number is entered.
func (s *BaseMySQLParserListener) EnterReal_ulong_number(ctx *Real_ulong_numberContext) {}

// ExitReal_ulong_number is called when production real_ulong_number is exited.
func (s *BaseMySQLParserListener) ExitReal_ulong_number(ctx *Real_ulong_numberContext) {}

// EnterUlonglong_number is called when production ulonglong_number is entered.
func (s *BaseMySQLParserListener) EnterUlonglong_number(ctx *Ulonglong_numberContext) {}

// ExitUlonglong_number is called when production ulonglong_number is exited.
func (s *BaseMySQLParserListener) ExitUlonglong_number(ctx *Ulonglong_numberContext) {}

// EnterReal_ulonglong_number is called when production real_ulonglong_number is entered.
func (s *BaseMySQLParserListener) EnterReal_ulonglong_number(ctx *Real_ulonglong_numberContext) {}

// ExitReal_ulonglong_number is called when production real_ulonglong_number is exited.
func (s *BaseMySQLParserListener) ExitReal_ulonglong_number(ctx *Real_ulonglong_numberContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseMySQLParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseMySQLParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterSignedLiteral is called when production signedLiteral is entered.
func (s *BaseMySQLParserListener) EnterSignedLiteral(ctx *SignedLiteralContext) {}

// ExitSignedLiteral is called when production signedLiteral is exited.
func (s *BaseMySQLParserListener) ExitSignedLiteral(ctx *SignedLiteralContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseMySQLParserListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseMySQLParserListener) ExitStringList(ctx *StringListContext) {}

// EnterTextStringLiteral is called when production textStringLiteral is entered.
func (s *BaseMySQLParserListener) EnterTextStringLiteral(ctx *TextStringLiteralContext) {}

// ExitTextStringLiteral is called when production textStringLiteral is exited.
func (s *BaseMySQLParserListener) ExitTextStringLiteral(ctx *TextStringLiteralContext) {}

// EnterTextString is called when production textString is entered.
func (s *BaseMySQLParserListener) EnterTextString(ctx *TextStringContext) {}

// ExitTextString is called when production textString is exited.
func (s *BaseMySQLParserListener) ExitTextString(ctx *TextStringContext) {}

// EnterTextStringHash is called when production textStringHash is entered.
func (s *BaseMySQLParserListener) EnterTextStringHash(ctx *TextStringHashContext) {}

// ExitTextStringHash is called when production textStringHash is exited.
func (s *BaseMySQLParserListener) ExitTextStringHash(ctx *TextStringHashContext) {}

// EnterTextLiteral is called when production textLiteral is entered.
func (s *BaseMySQLParserListener) EnterTextLiteral(ctx *TextLiteralContext) {}

// ExitTextLiteral is called when production textLiteral is exited.
func (s *BaseMySQLParserListener) ExitTextLiteral(ctx *TextLiteralContext) {}

// EnterTextStringNoLinebreak is called when production textStringNoLinebreak is entered.
func (s *BaseMySQLParserListener) EnterTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) {}

// ExitTextStringNoLinebreak is called when production textStringNoLinebreak is exited.
func (s *BaseMySQLParserListener) ExitTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) {}

// EnterTextStringLiteralList is called when production textStringLiteralList is entered.
func (s *BaseMySQLParserListener) EnterTextStringLiteralList(ctx *TextStringLiteralListContext) {}

// ExitTextStringLiteralList is called when production textStringLiteralList is exited.
func (s *BaseMySQLParserListener) ExitTextStringLiteralList(ctx *TextStringLiteralListContext) {}

// EnterNumLiteral is called when production numLiteral is entered.
func (s *BaseMySQLParserListener) EnterNumLiteral(ctx *NumLiteralContext) {}

// ExitNumLiteral is called when production numLiteral is exited.
func (s *BaseMySQLParserListener) ExitNumLiteral(ctx *NumLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BaseMySQLParserListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BaseMySQLParserListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseMySQLParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseMySQLParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterTemporalLiteral is called when production temporalLiteral is entered.
func (s *BaseMySQLParserListener) EnterTemporalLiteral(ctx *TemporalLiteralContext) {}

// ExitTemporalLiteral is called when production temporalLiteral is exited.
func (s *BaseMySQLParserListener) ExitTemporalLiteral(ctx *TemporalLiteralContext) {}

// EnterFloatOptions is called when production floatOptions is entered.
func (s *BaseMySQLParserListener) EnterFloatOptions(ctx *FloatOptionsContext) {}

// ExitFloatOptions is called when production floatOptions is exited.
func (s *BaseMySQLParserListener) ExitFloatOptions(ctx *FloatOptionsContext) {}

// EnterStandardFloatOptions is called when production standardFloatOptions is entered.
func (s *BaseMySQLParserListener) EnterStandardFloatOptions(ctx *StandardFloatOptionsContext) {}

// ExitStandardFloatOptions is called when production standardFloatOptions is exited.
func (s *BaseMySQLParserListener) ExitStandardFloatOptions(ctx *StandardFloatOptionsContext) {}

// EnterPrecision is called when production precision is entered.
func (s *BaseMySQLParserListener) EnterPrecision(ctx *PrecisionContext) {}

// ExitPrecision is called when production precision is exited.
func (s *BaseMySQLParserListener) ExitPrecision(ctx *PrecisionContext) {}

// EnterTextOrIdentifier is called when production textOrIdentifier is entered.
func (s *BaseMySQLParserListener) EnterTextOrIdentifier(ctx *TextOrIdentifierContext) {}

// ExitTextOrIdentifier is called when production textOrIdentifier is exited.
func (s *BaseMySQLParserListener) ExitTextOrIdentifier(ctx *TextOrIdentifierContext) {}

// EnterLValueIdentifier is called when production lValueIdentifier is entered.
func (s *BaseMySQLParserListener) EnterLValueIdentifier(ctx *LValueIdentifierContext) {}

// ExitLValueIdentifier is called when production lValueIdentifier is exited.
func (s *BaseMySQLParserListener) ExitLValueIdentifier(ctx *LValueIdentifierContext) {}

// EnterRoleIdentifierOrText is called when production roleIdentifierOrText is entered.
func (s *BaseMySQLParserListener) EnterRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) {}

// ExitRoleIdentifierOrText is called when production roleIdentifierOrText is exited.
func (s *BaseMySQLParserListener) ExitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) {}

// EnterSizeNumber is called when production sizeNumber is entered.
func (s *BaseMySQLParserListener) EnterSizeNumber(ctx *SizeNumberContext) {}

// ExitSizeNumber is called when production sizeNumber is exited.
func (s *BaseMySQLParserListener) ExitSizeNumber(ctx *SizeNumberContext) {}

// EnterParentheses is called when production parentheses is entered.
func (s *BaseMySQLParserListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production parentheses is exited.
func (s *BaseMySQLParserListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseMySQLParserListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseMySQLParserListener) ExitEqual(ctx *EqualContext) {}

// EnterOptionType is called when production optionType is entered.
func (s *BaseMySQLParserListener) EnterOptionType(ctx *OptionTypeContext) {}

// ExitOptionType is called when production optionType is exited.
func (s *BaseMySQLParserListener) ExitOptionType(ctx *OptionTypeContext) {}

// EnterVarIdentType is called when production varIdentType is entered.
func (s *BaseMySQLParserListener) EnterVarIdentType(ctx *VarIdentTypeContext) {}

// ExitVarIdentType is called when production varIdentType is exited.
func (s *BaseMySQLParserListener) ExitVarIdentType(ctx *VarIdentTypeContext) {}

// EnterSetVarIdentType is called when production setVarIdentType is entered.
func (s *BaseMySQLParserListener) EnterSetVarIdentType(ctx *SetVarIdentTypeContext) {}

// ExitSetVarIdentType is called when production setVarIdentType is exited.
func (s *BaseMySQLParserListener) ExitSetVarIdentType(ctx *SetVarIdentTypeContext) {}

// EnterIdentifierKeyword is called when production identifierKeyword is entered.
func (s *BaseMySQLParserListener) EnterIdentifierKeyword(ctx *IdentifierKeywordContext) {}

// ExitIdentifierKeyword is called when production identifierKeyword is exited.
func (s *BaseMySQLParserListener) ExitIdentifierKeyword(ctx *IdentifierKeywordContext) {}

// EnterIdentifierKeywordsAmbiguous1RolesAndLabels is called when production identifierKeywordsAmbiguous1RolesAndLabels is entered.
func (s *BaseMySQLParserListener) EnterIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) {
}

// ExitIdentifierKeywordsAmbiguous1RolesAndLabels is called when production identifierKeywordsAmbiguous1RolesAndLabels is exited.
func (s *BaseMySQLParserListener) ExitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) {
}

// EnterIdentifierKeywordsAmbiguous2Labels is called when production identifierKeywordsAmbiguous2Labels is entered.
func (s *BaseMySQLParserListener) EnterIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) {
}

// ExitIdentifierKeywordsAmbiguous2Labels is called when production identifierKeywordsAmbiguous2Labels is exited.
func (s *BaseMySQLParserListener) ExitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) {
}

// EnterLabelKeyword is called when production labelKeyword is entered.
func (s *BaseMySQLParserListener) EnterLabelKeyword(ctx *LabelKeywordContext) {}

// ExitLabelKeyword is called when production labelKeyword is exited.
func (s *BaseMySQLParserListener) ExitLabelKeyword(ctx *LabelKeywordContext) {}

// EnterIdentifierKeywordsAmbiguous3Roles is called when production identifierKeywordsAmbiguous3Roles is entered.
func (s *BaseMySQLParserListener) EnterIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) {
}

// ExitIdentifierKeywordsAmbiguous3Roles is called when production identifierKeywordsAmbiguous3Roles is exited.
func (s *BaseMySQLParserListener) ExitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) {
}

// EnterIdentifierKeywordsUnambiguous is called when production identifierKeywordsUnambiguous is entered.
func (s *BaseMySQLParserListener) EnterIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) {
}

// ExitIdentifierKeywordsUnambiguous is called when production identifierKeywordsUnambiguous is exited.
func (s *BaseMySQLParserListener) ExitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) {
}

// EnterRoleKeyword is called when production roleKeyword is entered.
func (s *BaseMySQLParserListener) EnterRoleKeyword(ctx *RoleKeywordContext) {}

// ExitRoleKeyword is called when production roleKeyword is exited.
func (s *BaseMySQLParserListener) ExitRoleKeyword(ctx *RoleKeywordContext) {}

// EnterLValueKeyword is called when production lValueKeyword is entered.
func (s *BaseMySQLParserListener) EnterLValueKeyword(ctx *LValueKeywordContext) {}

// ExitLValueKeyword is called when production lValueKeyword is exited.
func (s *BaseMySQLParserListener) ExitLValueKeyword(ctx *LValueKeywordContext) {}

// EnterIdentifierKeywordsAmbiguous4SystemVariables is called when production identifierKeywordsAmbiguous4SystemVariables is entered.
func (s *BaseMySQLParserListener) EnterIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) {
}

// ExitIdentifierKeywordsAmbiguous4SystemVariables is called when production identifierKeywordsAmbiguous4SystemVariables is exited.
func (s *BaseMySQLParserListener) ExitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) {
}

// EnterRoleOrIdentifierKeyword is called when production roleOrIdentifierKeyword is entered.
func (s *BaseMySQLParserListener) EnterRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) {}

// ExitRoleOrIdentifierKeyword is called when production roleOrIdentifierKeyword is exited.
func (s *BaseMySQLParserListener) ExitRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) {}

// EnterRoleOrLabelKeyword is called when production roleOrLabelKeyword is entered.
func (s *BaseMySQLParserListener) EnterRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) {}

// ExitRoleOrLabelKeyword is called when production roleOrLabelKeyword is exited.
func (s *BaseMySQLParserListener) ExitRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) {}
