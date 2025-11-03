// Code generated from MySQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mysql // MySQLParser
import "github.com/antlr4-go/antlr/v4"

// MySQLParserListener is a complete listener for a parse tree produced by MySQLParser.
type MySQLParserListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterAlterStatement is called when entering the alterStatement production.
	EnterAlterStatement(c *AlterStatementContext)

	// EnterAlterDatabase is called when entering the alterDatabase production.
	EnterAlterDatabase(c *AlterDatabaseContext)

	// EnterAlterDatabaseOption is called when entering the alterDatabaseOption production.
	EnterAlterDatabaseOption(c *AlterDatabaseOptionContext)

	// EnterAlterEvent is called when entering the alterEvent production.
	EnterAlterEvent(c *AlterEventContext)

	// EnterAlterLogfileGroup is called when entering the alterLogfileGroup production.
	EnterAlterLogfileGroup(c *AlterLogfileGroupContext)

	// EnterAlterLogfileGroupOptions is called when entering the alterLogfileGroupOptions production.
	EnterAlterLogfileGroupOptions(c *AlterLogfileGroupOptionsContext)

	// EnterAlterLogfileGroupOption is called when entering the alterLogfileGroupOption production.
	EnterAlterLogfileGroupOption(c *AlterLogfileGroupOptionContext)

	// EnterAlterServer is called when entering the alterServer production.
	EnterAlterServer(c *AlterServerContext)

	// EnterAlterTable is called when entering the alterTable production.
	EnterAlterTable(c *AlterTableContext)

	// EnterAlterTableActions is called when entering the alterTableActions production.
	EnterAlterTableActions(c *AlterTableActionsContext)

	// EnterAlterCommandList is called when entering the alterCommandList production.
	EnterAlterCommandList(c *AlterCommandListContext)

	// EnterAlterCommandsModifierList is called when entering the alterCommandsModifierList production.
	EnterAlterCommandsModifierList(c *AlterCommandsModifierListContext)

	// EnterStandaloneAlterCommands is called when entering the standaloneAlterCommands production.
	EnterStandaloneAlterCommands(c *StandaloneAlterCommandsContext)

	// EnterAlterPartition is called when entering the alterPartition production.
	EnterAlterPartition(c *AlterPartitionContext)

	// EnterAlterList is called when entering the alterList production.
	EnterAlterList(c *AlterListContext)

	// EnterAlterCommandsModifier is called when entering the alterCommandsModifier production.
	EnterAlterCommandsModifier(c *AlterCommandsModifierContext)

	// EnterAlterListItem is called when entering the alterListItem production.
	EnterAlterListItem(c *AlterListItemContext)

	// EnterPlace is called when entering the place production.
	EnterPlace(c *PlaceContext)

	// EnterRestrict is called when entering the restrict production.
	EnterRestrict(c *RestrictContext)

	// EnterAlterOrderList is called when entering the alterOrderList production.
	EnterAlterOrderList(c *AlterOrderListContext)

	// EnterAlterAlgorithmOption is called when entering the alterAlgorithmOption production.
	EnterAlterAlgorithmOption(c *AlterAlgorithmOptionContext)

	// EnterAlterLockOption is called when entering the alterLockOption production.
	EnterAlterLockOption(c *AlterLockOptionContext)

	// EnterIndexLockAndAlgorithm is called when entering the indexLockAndAlgorithm production.
	EnterIndexLockAndAlgorithm(c *IndexLockAndAlgorithmContext)

	// EnterWithValidation is called when entering the withValidation production.
	EnterWithValidation(c *WithValidationContext)

	// EnterRemovePartitioning is called when entering the removePartitioning production.
	EnterRemovePartitioning(c *RemovePartitioningContext)

	// EnterAllOrPartitionNameList is called when entering the allOrPartitionNameList production.
	EnterAllOrPartitionNameList(c *AllOrPartitionNameListContext)

	// EnterAlterTablespace is called when entering the alterTablespace production.
	EnterAlterTablespace(c *AlterTablespaceContext)

	// EnterAlterUndoTablespace is called when entering the alterUndoTablespace production.
	EnterAlterUndoTablespace(c *AlterUndoTablespaceContext)

	// EnterUndoTableSpaceOptions is called when entering the undoTableSpaceOptions production.
	EnterUndoTableSpaceOptions(c *UndoTableSpaceOptionsContext)

	// EnterUndoTableSpaceOption is called when entering the undoTableSpaceOption production.
	EnterUndoTableSpaceOption(c *UndoTableSpaceOptionContext)

	// EnterAlterTablespaceOptions is called when entering the alterTablespaceOptions production.
	EnterAlterTablespaceOptions(c *AlterTablespaceOptionsContext)

	// EnterAlterTablespaceOption is called when entering the alterTablespaceOption production.
	EnterAlterTablespaceOption(c *AlterTablespaceOptionContext)

	// EnterChangeTablespaceOption is called when entering the changeTablespaceOption production.
	EnterChangeTablespaceOption(c *ChangeTablespaceOptionContext)

	// EnterAlterView is called when entering the alterView production.
	EnterAlterView(c *AlterViewContext)

	// EnterViewTail is called when entering the viewTail production.
	EnterViewTail(c *ViewTailContext)

	// EnterViewSelect is called when entering the viewSelect production.
	EnterViewSelect(c *ViewSelectContext)

	// EnterViewCheckOption is called when entering the viewCheckOption production.
	EnterViewCheckOption(c *ViewCheckOptionContext)

	// EnterAlterProcedure is called when entering the alterProcedure production.
	EnterAlterProcedure(c *AlterProcedureContext)

	// EnterAlterFunction is called when entering the alterFunction production.
	EnterAlterFunction(c *AlterFunctionContext)

	// EnterCreateStatement is called when entering the createStatement production.
	EnterCreateStatement(c *CreateStatementContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterCreateDatabaseOption is called when entering the createDatabaseOption production.
	EnterCreateDatabaseOption(c *CreateDatabaseOptionContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterTableElementList is called when entering the tableElementList production.
	EnterTableElementList(c *TableElementListContext)

	// EnterTableElement is called when entering the tableElement production.
	EnterTableElement(c *TableElementContext)

	// EnterDuplicateAsQueryExpression is called when entering the duplicateAsQueryExpression production.
	EnterDuplicateAsQueryExpression(c *DuplicateAsQueryExpressionContext)

	// EnterQueryExpressionOrParens is called when entering the queryExpressionOrParens production.
	EnterQueryExpressionOrParens(c *QueryExpressionOrParensContext)

	// EnterCreateRoutine is called when entering the createRoutine production.
	EnterCreateRoutine(c *CreateRoutineContext)

	// EnterCreateProcedure is called when entering the createProcedure production.
	EnterCreateProcedure(c *CreateProcedureContext)

	// EnterCreateFunction is called when entering the createFunction production.
	EnterCreateFunction(c *CreateFunctionContext)

	// EnterCreateUdf is called when entering the createUdf production.
	EnterCreateUdf(c *CreateUdfContext)

	// EnterRoutineCreateOption is called when entering the routineCreateOption production.
	EnterRoutineCreateOption(c *RoutineCreateOptionContext)

	// EnterRoutineAlterOptions is called when entering the routineAlterOptions production.
	EnterRoutineAlterOptions(c *RoutineAlterOptionsContext)

	// EnterRoutineOption is called when entering the routineOption production.
	EnterRoutineOption(c *RoutineOptionContext)

	// EnterCreateIndex is called when entering the createIndex production.
	EnterCreateIndex(c *CreateIndexContext)

	// EnterIndexNameAndType is called when entering the indexNameAndType production.
	EnterIndexNameAndType(c *IndexNameAndTypeContext)

	// EnterCreateIndexTarget is called when entering the createIndexTarget production.
	EnterCreateIndexTarget(c *CreateIndexTargetContext)

	// EnterCreateLogfileGroup is called when entering the createLogfileGroup production.
	EnterCreateLogfileGroup(c *CreateLogfileGroupContext)

	// EnterLogfileGroupOptions is called when entering the logfileGroupOptions production.
	EnterLogfileGroupOptions(c *LogfileGroupOptionsContext)

	// EnterLogfileGroupOption is called when entering the logfileGroupOption production.
	EnterLogfileGroupOption(c *LogfileGroupOptionContext)

	// EnterCreateServer is called when entering the createServer production.
	EnterCreateServer(c *CreateServerContext)

	// EnterServerOptions is called when entering the serverOptions production.
	EnterServerOptions(c *ServerOptionsContext)

	// EnterServerOption is called when entering the serverOption production.
	EnterServerOption(c *ServerOptionContext)

	// EnterCreateTablespace is called when entering the createTablespace production.
	EnterCreateTablespace(c *CreateTablespaceContext)

	// EnterCreateUndoTablespace is called when entering the createUndoTablespace production.
	EnterCreateUndoTablespace(c *CreateUndoTablespaceContext)

	// EnterTsDataFileName is called when entering the tsDataFileName production.
	EnterTsDataFileName(c *TsDataFileNameContext)

	// EnterTsDataFile is called when entering the tsDataFile production.
	EnterTsDataFile(c *TsDataFileContext)

	// EnterTablespaceOptions is called when entering the tablespaceOptions production.
	EnterTablespaceOptions(c *TablespaceOptionsContext)

	// EnterTablespaceOption is called when entering the tablespaceOption production.
	EnterTablespaceOption(c *TablespaceOptionContext)

	// EnterTsOptionInitialSize is called when entering the tsOptionInitialSize production.
	EnterTsOptionInitialSize(c *TsOptionInitialSizeContext)

	// EnterTsOptionUndoRedoBufferSize is called when entering the tsOptionUndoRedoBufferSize production.
	EnterTsOptionUndoRedoBufferSize(c *TsOptionUndoRedoBufferSizeContext)

	// EnterTsOptionAutoextendSize is called when entering the tsOptionAutoextendSize production.
	EnterTsOptionAutoextendSize(c *TsOptionAutoextendSizeContext)

	// EnterTsOptionMaxSize is called when entering the tsOptionMaxSize production.
	EnterTsOptionMaxSize(c *TsOptionMaxSizeContext)

	// EnterTsOptionExtentSize is called when entering the tsOptionExtentSize production.
	EnterTsOptionExtentSize(c *TsOptionExtentSizeContext)

	// EnterTsOptionNodegroup is called when entering the tsOptionNodegroup production.
	EnterTsOptionNodegroup(c *TsOptionNodegroupContext)

	// EnterTsOptionEngine is called when entering the tsOptionEngine production.
	EnterTsOptionEngine(c *TsOptionEngineContext)

	// EnterTsOptionWait is called when entering the tsOptionWait production.
	EnterTsOptionWait(c *TsOptionWaitContext)

	// EnterTsOptionComment is called when entering the tsOptionComment production.
	EnterTsOptionComment(c *TsOptionCommentContext)

	// EnterTsOptionFileblockSize is called when entering the tsOptionFileblockSize production.
	EnterTsOptionFileblockSize(c *TsOptionFileblockSizeContext)

	// EnterTsOptionEncryption is called when entering the tsOptionEncryption production.
	EnterTsOptionEncryption(c *TsOptionEncryptionContext)

	// EnterCreateView is called when entering the createView production.
	EnterCreateView(c *CreateViewContext)

	// EnterViewReplaceOrAlgorithm is called when entering the viewReplaceOrAlgorithm production.
	EnterViewReplaceOrAlgorithm(c *ViewReplaceOrAlgorithmContext)

	// EnterViewAlgorithm is called when entering the viewAlgorithm production.
	EnterViewAlgorithm(c *ViewAlgorithmContext)

	// EnterViewSuid is called when entering the viewSuid production.
	EnterViewSuid(c *ViewSuidContext)

	// EnterCreateTrigger is called when entering the createTrigger production.
	EnterCreateTrigger(c *CreateTriggerContext)

	// EnterTriggerFollowsPrecedesClause is called when entering the triggerFollowsPrecedesClause production.
	EnterTriggerFollowsPrecedesClause(c *TriggerFollowsPrecedesClauseContext)

	// EnterCreateEvent is called when entering the createEvent production.
	EnterCreateEvent(c *CreateEventContext)

	// EnterCreateRole is called when entering the createRole production.
	EnterCreateRole(c *CreateRoleContext)

	// EnterCreateSpatialReference is called when entering the createSpatialReference production.
	EnterCreateSpatialReference(c *CreateSpatialReferenceContext)

	// EnterSrsAttribute is called when entering the srsAttribute production.
	EnterSrsAttribute(c *SrsAttributeContext)

	// EnterDropStatement is called when entering the dropStatement production.
	EnterDropStatement(c *DropStatementContext)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterDropEvent is called when entering the dropEvent production.
	EnterDropEvent(c *DropEventContext)

	// EnterDropFunction is called when entering the dropFunction production.
	EnterDropFunction(c *DropFunctionContext)

	// EnterDropProcedure is called when entering the dropProcedure production.
	EnterDropProcedure(c *DropProcedureContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterDropLogfileGroup is called when entering the dropLogfileGroup production.
	EnterDropLogfileGroup(c *DropLogfileGroupContext)

	// EnterDropLogfileGroupOption is called when entering the dropLogfileGroupOption production.
	EnterDropLogfileGroupOption(c *DropLogfileGroupOptionContext)

	// EnterDropServer is called when entering the dropServer production.
	EnterDropServer(c *DropServerContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropTableSpace is called when entering the dropTableSpace production.
	EnterDropTableSpace(c *DropTableSpaceContext)

	// EnterDropTrigger is called when entering the dropTrigger production.
	EnterDropTrigger(c *DropTriggerContext)

	// EnterDropView is called when entering the dropView production.
	EnterDropView(c *DropViewContext)

	// EnterDropRole is called when entering the dropRole production.
	EnterDropRole(c *DropRoleContext)

	// EnterDropSpatialReference is called when entering the dropSpatialReference production.
	EnterDropSpatialReference(c *DropSpatialReferenceContext)

	// EnterDropUndoTablespace is called when entering the dropUndoTablespace production.
	EnterDropUndoTablespace(c *DropUndoTablespaceContext)

	// EnterRenameTableStatement is called when entering the renameTableStatement production.
	EnterRenameTableStatement(c *RenameTableStatementContext)

	// EnterRenamePair is called when entering the renamePair production.
	EnterRenamePair(c *RenamePairContext)

	// EnterTruncateTableStatement is called when entering the truncateTableStatement production.
	EnterTruncateTableStatement(c *TruncateTableStatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterCallStatement is called when entering the callStatement production.
	EnterCallStatement(c *CallStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterPartitionDelete is called when entering the partitionDelete production.
	EnterPartitionDelete(c *PartitionDeleteContext)

	// EnterDeleteStatementOption is called when entering the deleteStatementOption production.
	EnterDeleteStatementOption(c *DeleteStatementOptionContext)

	// EnterDoStatement is called when entering the doStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterHandlerStatement is called when entering the handlerStatement production.
	EnterHandlerStatement(c *HandlerStatementContext)

	// EnterHandlerReadOrScan is called when entering the handlerReadOrScan production.
	EnterHandlerReadOrScan(c *HandlerReadOrScanContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterInsertLockOption is called when entering the insertLockOption production.
	EnterInsertLockOption(c *InsertLockOptionContext)

	// EnterInsertFromConstructor is called when entering the insertFromConstructor production.
	EnterInsertFromConstructor(c *InsertFromConstructorContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterInsertValues is called when entering the insertValues production.
	EnterInsertValues(c *InsertValuesContext)

	// EnterInsertQueryExpression is called when entering the insertQueryExpression production.
	EnterInsertQueryExpression(c *InsertQueryExpressionContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterValuesReference is called when entering the valuesReference production.
	EnterValuesReference(c *ValuesReferenceContext)

	// EnterInsertUpdateList is called when entering the insertUpdateList production.
	EnterInsertUpdateList(c *InsertUpdateListContext)

	// EnterLoadStatement is called when entering the loadStatement production.
	EnterLoadStatement(c *LoadStatementContext)

	// EnterDataOrXml is called when entering the dataOrXml production.
	EnterDataOrXml(c *DataOrXmlContext)

	// EnterXmlRowsIdentifiedBy is called when entering the xmlRowsIdentifiedBy production.
	EnterXmlRowsIdentifiedBy(c *XmlRowsIdentifiedByContext)

	// EnterLoadDataFileTail is called when entering the loadDataFileTail production.
	EnterLoadDataFileTail(c *LoadDataFileTailContext)

	// EnterLoadDataFileTargetList is called when entering the loadDataFileTargetList production.
	EnterLoadDataFileTargetList(c *LoadDataFileTargetListContext)

	// EnterFieldOrVariableList is called when entering the fieldOrVariableList production.
	EnterFieldOrVariableList(c *FieldOrVariableListContext)

	// EnterReplaceStatement is called when entering the replaceStatement production.
	EnterReplaceStatement(c *ReplaceStatementContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSelectStatementWithInto is called when entering the selectStatementWithInto production.
	EnterSelectStatementWithInto(c *SelectStatementWithIntoContext)

	// EnterQueryExpression is called when entering the queryExpression production.
	EnterQueryExpression(c *QueryExpressionContext)

	// EnterQueryExpressionBody is called when entering the queryExpressionBody production.
	EnterQueryExpressionBody(c *QueryExpressionBodyContext)

	// EnterQueryExpressionParens is called when entering the queryExpressionParens production.
	EnterQueryExpressionParens(c *QueryExpressionParensContext)

	// EnterQueryPrimary is called when entering the queryPrimary production.
	EnterQueryPrimary(c *QueryPrimaryContext)

	// EnterQuerySpecification is called when entering the querySpecification production.
	EnterQuerySpecification(c *QuerySpecificationContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterQuerySpecOption is called when entering the querySpecOption production.
	EnterQuerySpecOption(c *QuerySpecOptionContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterSimpleLimitClause is called when entering the simpleLimitClause production.
	EnterSimpleLimitClause(c *SimpleLimitClauseContext)

	// EnterLimitOptions is called when entering the limitOptions production.
	EnterLimitOptions(c *LimitOptionsContext)

	// EnterLimitOption is called when entering the limitOption production.
	EnterLimitOption(c *LimitOptionContext)

	// EnterIntoClause is called when entering the intoClause production.
	EnterIntoClause(c *IntoClauseContext)

	// EnterProcedureAnalyseClause is called when entering the procedureAnalyseClause production.
	EnterProcedureAnalyseClause(c *ProcedureAnalyseClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

	// EnterWindowDefinition is called when entering the windowDefinition production.
	EnterWindowDefinition(c *WindowDefinitionContext)

	// EnterWindowSpec is called when entering the windowSpec production.
	EnterWindowSpec(c *WindowSpecContext)

	// EnterWindowSpecDetails is called when entering the windowSpecDetails production.
	EnterWindowSpecDetails(c *WindowSpecDetailsContext)

	// EnterWindowFrameClause is called when entering the windowFrameClause production.
	EnterWindowFrameClause(c *WindowFrameClauseContext)

	// EnterWindowFrameUnits is called when entering the windowFrameUnits production.
	EnterWindowFrameUnits(c *WindowFrameUnitsContext)

	// EnterWindowFrameExtent is called when entering the windowFrameExtent production.
	EnterWindowFrameExtent(c *WindowFrameExtentContext)

	// EnterWindowFrameStart is called when entering the windowFrameStart production.
	EnterWindowFrameStart(c *WindowFrameStartContext)

	// EnterWindowFrameBetween is called when entering the windowFrameBetween production.
	EnterWindowFrameBetween(c *WindowFrameBetweenContext)

	// EnterWindowFrameBound is called when entering the windowFrameBound production.
	EnterWindowFrameBound(c *WindowFrameBoundContext)

	// EnterWindowFrameExclusion is called when entering the windowFrameExclusion production.
	EnterWindowFrameExclusion(c *WindowFrameExclusionContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterCommonTableExpression is called when entering the commonTableExpression production.
	EnterCommonTableExpression(c *CommonTableExpressionContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterOlapOption is called when entering the olapOption production.
	EnterOlapOption(c *OlapOptionContext)

	// EnterOrderClause is called when entering the orderClause production.
	EnterOrderClause(c *OrderClauseContext)

	// EnterDirection is called when entering the direction production.
	EnterDirection(c *DirectionContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterTableReferenceList is called when entering the tableReferenceList production.
	EnterTableReferenceList(c *TableReferenceListContext)

	// EnterTableValueConstructor is called when entering the tableValueConstructor production.
	EnterTableValueConstructor(c *TableValueConstructorContext)

	// EnterExplicitTable is called when entering the explicitTable production.
	EnterExplicitTable(c *ExplicitTableContext)

	// EnterRowValueExplicit is called when entering the rowValueExplicit production.
	EnterRowValueExplicit(c *RowValueExplicitContext)

	// EnterSelectOption is called when entering the selectOption production.
	EnterSelectOption(c *SelectOptionContext)

	// EnterLockingClauseList is called when entering the lockingClauseList production.
	EnterLockingClauseList(c *LockingClauseListContext)

	// EnterLockingClause is called when entering the lockingClause production.
	EnterLockingClause(c *LockingClauseContext)

	// EnterLockStrengh is called when entering the lockStrengh production.
	EnterLockStrengh(c *LockStrenghContext)

	// EnterLockedRowAction is called when entering the lockedRowAction production.
	EnterLockedRowAction(c *LockedRowActionContext)

	// EnterSelectItemList is called when entering the selectItemList production.
	EnterSelectItemList(c *SelectItemListContext)

	// EnterSelectItem is called when entering the selectItem production.
	EnterSelectItem(c *SelectItemContext)

	// EnterSelectAlias is called when entering the selectAlias production.
	EnterSelectAlias(c *SelectAliasContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterTableReference is called when entering the tableReference production.
	EnterTableReference(c *TableReferenceContext)

	// EnterEscapedTableReference is called when entering the escapedTableReference production.
	EnterEscapedTableReference(c *EscapedTableReferenceContext)

	// EnterJoinedTable is called when entering the joinedTable production.
	EnterJoinedTable(c *JoinedTableContext)

	// EnterNaturalJoinType is called when entering the naturalJoinType production.
	EnterNaturalJoinType(c *NaturalJoinTypeContext)

	// EnterInnerJoinType is called when entering the innerJoinType production.
	EnterInnerJoinType(c *InnerJoinTypeContext)

	// EnterOuterJoinType is called when entering the outerJoinType production.
	EnterOuterJoinType(c *OuterJoinTypeContext)

	// EnterTableFactor is called when entering the tableFactor production.
	EnterTableFactor(c *TableFactorContext)

	// EnterSingleTable is called when entering the singleTable production.
	EnterSingleTable(c *SingleTableContext)

	// EnterSingleTableParens is called when entering the singleTableParens production.
	EnterSingleTableParens(c *SingleTableParensContext)

	// EnterDerivedTable is called when entering the derivedTable production.
	EnterDerivedTable(c *DerivedTableContext)

	// EnterTableReferenceListParens is called when entering the tableReferenceListParens production.
	EnterTableReferenceListParens(c *TableReferenceListParensContext)

	// EnterTableFunction is called when entering the tableFunction production.
	EnterTableFunction(c *TableFunctionContext)

	// EnterColumnsClause is called when entering the columnsClause production.
	EnterColumnsClause(c *ColumnsClauseContext)

	// EnterJtColumn is called when entering the jtColumn production.
	EnterJtColumn(c *JtColumnContext)

	// EnterOnEmptyOrError is called when entering the onEmptyOrError production.
	EnterOnEmptyOrError(c *OnEmptyOrErrorContext)

	// EnterOnEmpty is called when entering the onEmpty production.
	EnterOnEmpty(c *OnEmptyContext)

	// EnterOnError is called when entering the onError production.
	EnterOnError(c *OnErrorContext)

	// EnterJtOnResponse is called when entering the jtOnResponse production.
	EnterJtOnResponse(c *JtOnResponseContext)

	// EnterSetOprSymbol is called when entering the setOprSymbol production.
	EnterSetOprSymbol(c *SetOprSymbolContext)

	// EnterSetOprOption is called when entering the setOprOption production.
	EnterSetOprOption(c *SetOprOptionContext)

	// EnterTableAlias is called when entering the tableAlias production.
	EnterTableAlias(c *TableAliasContext)

	// EnterIndexHintList is called when entering the indexHintList production.
	EnterIndexHintList(c *IndexHintListContext)

	// EnterIndexHint is called when entering the indexHint production.
	EnterIndexHint(c *IndexHintContext)

	// EnterIndexHintType is called when entering the indexHintType production.
	EnterIndexHintType(c *IndexHintTypeContext)

	// EnterKeyOrIndex is called when entering the keyOrIndex production.
	EnterKeyOrIndex(c *KeyOrIndexContext)

	// EnterConstraintKeyType is called when entering the constraintKeyType production.
	EnterConstraintKeyType(c *ConstraintKeyTypeContext)

	// EnterIndexHintClause is called when entering the indexHintClause production.
	EnterIndexHintClause(c *IndexHintClauseContext)

	// EnterIndexList is called when entering the indexList production.
	EnterIndexList(c *IndexListContext)

	// EnterIndexListElement is called when entering the indexListElement production.
	EnterIndexListElement(c *IndexListElementContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterTransactionOrLockingStatement is called when entering the transactionOrLockingStatement production.
	EnterTransactionOrLockingStatement(c *TransactionOrLockingStatementContext)

	// EnterTransactionStatement is called when entering the transactionStatement production.
	EnterTransactionStatement(c *TransactionStatementContext)

	// EnterBeginWork is called when entering the beginWork production.
	EnterBeginWork(c *BeginWorkContext)

	// EnterTransactionCharacteristic is called when entering the transactionCharacteristic production.
	EnterTransactionCharacteristic(c *TransactionCharacteristicContext)

	// EnterSavepointStatement is called when entering the savepointStatement production.
	EnterSavepointStatement(c *SavepointStatementContext)

	// EnterLockStatement is called when entering the lockStatement production.
	EnterLockStatement(c *LockStatementContext)

	// EnterLockItem is called when entering the lockItem production.
	EnterLockItem(c *LockItemContext)

	// EnterLockOption is called when entering the lockOption production.
	EnterLockOption(c *LockOptionContext)

	// EnterXaStatement is called when entering the xaStatement production.
	EnterXaStatement(c *XaStatementContext)

	// EnterXaConvert is called when entering the xaConvert production.
	EnterXaConvert(c *XaConvertContext)

	// EnterXid is called when entering the xid production.
	EnterXid(c *XidContext)

	// EnterReplicationStatement is called when entering the replicationStatement production.
	EnterReplicationStatement(c *ReplicationStatementContext)

	// EnterResetOption is called when entering the resetOption production.
	EnterResetOption(c *ResetOptionContext)

	// EnterMasterResetOptions is called when entering the masterResetOptions production.
	EnterMasterResetOptions(c *MasterResetOptionsContext)

	// EnterReplicationLoad is called when entering the replicationLoad production.
	EnterReplicationLoad(c *ReplicationLoadContext)

	// EnterChangeMaster is called when entering the changeMaster production.
	EnterChangeMaster(c *ChangeMasterContext)

	// EnterChangeMasterOptions is called when entering the changeMasterOptions production.
	EnterChangeMasterOptions(c *ChangeMasterOptionsContext)

	// EnterMasterOption is called when entering the masterOption production.
	EnterMasterOption(c *MasterOptionContext)

	// EnterPrivilegeCheckDef is called when entering the privilegeCheckDef production.
	EnterPrivilegeCheckDef(c *PrivilegeCheckDefContext)

	// EnterTablePrimaryKeyCheckDef is called when entering the tablePrimaryKeyCheckDef production.
	EnterTablePrimaryKeyCheckDef(c *TablePrimaryKeyCheckDefContext)

	// EnterMasterTlsCiphersuitesDef is called when entering the masterTlsCiphersuitesDef production.
	EnterMasterTlsCiphersuitesDef(c *MasterTlsCiphersuitesDefContext)

	// EnterMasterFileDef is called when entering the masterFileDef production.
	EnterMasterFileDef(c *MasterFileDefContext)

	// EnterServerIdList is called when entering the serverIdList production.
	EnterServerIdList(c *ServerIdListContext)

	// EnterChangeReplication is called when entering the changeReplication production.
	EnterChangeReplication(c *ChangeReplicationContext)

	// EnterFilterDefinition is called when entering the filterDefinition production.
	EnterFilterDefinition(c *FilterDefinitionContext)

	// EnterFilterDbList is called when entering the filterDbList production.
	EnterFilterDbList(c *FilterDbListContext)

	// EnterFilterTableList is called when entering the filterTableList production.
	EnterFilterTableList(c *FilterTableListContext)

	// EnterFilterStringList is called when entering the filterStringList production.
	EnterFilterStringList(c *FilterStringListContext)

	// EnterFilterWildDbTableString is called when entering the filterWildDbTableString production.
	EnterFilterWildDbTableString(c *FilterWildDbTableStringContext)

	// EnterFilterDbPairList is called when entering the filterDbPairList production.
	EnterFilterDbPairList(c *FilterDbPairListContext)

	// EnterSlave is called when entering the slave production.
	EnterSlave(c *SlaveContext)

	// EnterSlaveUntilOptions is called when entering the slaveUntilOptions production.
	EnterSlaveUntilOptions(c *SlaveUntilOptionsContext)

	// EnterSlaveConnectionOptions is called when entering the slaveConnectionOptions production.
	EnterSlaveConnectionOptions(c *SlaveConnectionOptionsContext)

	// EnterSlaveThreadOptions is called when entering the slaveThreadOptions production.
	EnterSlaveThreadOptions(c *SlaveThreadOptionsContext)

	// EnterSlaveThreadOption is called when entering the slaveThreadOption production.
	EnterSlaveThreadOption(c *SlaveThreadOptionContext)

	// EnterGroupReplication is called when entering the groupReplication production.
	EnterGroupReplication(c *GroupReplicationContext)

	// EnterPreparedStatement is called when entering the preparedStatement production.
	EnterPreparedStatement(c *PreparedStatementContext)

	// EnterExecuteStatement is called when entering the executeStatement production.
	EnterExecuteStatement(c *ExecuteStatementContext)

	// EnterExecuteVarList is called when entering the executeVarList production.
	EnterExecuteVarList(c *ExecuteVarListContext)

	// EnterCloneStatement is called when entering the cloneStatement production.
	EnterCloneStatement(c *CloneStatementContext)

	// EnterDataDirSSL is called when entering the dataDirSSL production.
	EnterDataDirSSL(c *DataDirSSLContext)

	// EnterSsl is called when entering the ssl production.
	EnterSsl(c *SslContext)

	// EnterAccountManagementStatement is called when entering the accountManagementStatement production.
	EnterAccountManagementStatement(c *AccountManagementStatementContext)

	// EnterAlterUser is called when entering the alterUser production.
	EnterAlterUser(c *AlterUserContext)

	// EnterAlterUserTail is called when entering the alterUserTail production.
	EnterAlterUserTail(c *AlterUserTailContext)

	// EnterUserFunction is called when entering the userFunction production.
	EnterUserFunction(c *UserFunctionContext)

	// EnterCreateUser is called when entering the createUser production.
	EnterCreateUser(c *CreateUserContext)

	// EnterCreateUserTail is called when entering the createUserTail production.
	EnterCreateUserTail(c *CreateUserTailContext)

	// EnterDefaultRoleClause is called when entering the defaultRoleClause production.
	EnterDefaultRoleClause(c *DefaultRoleClauseContext)

	// EnterRequireClause is called when entering the requireClause production.
	EnterRequireClause(c *RequireClauseContext)

	// EnterConnectOptions is called when entering the connectOptions production.
	EnterConnectOptions(c *ConnectOptionsContext)

	// EnterAccountLockPasswordExpireOptions is called when entering the accountLockPasswordExpireOptions production.
	EnterAccountLockPasswordExpireOptions(c *AccountLockPasswordExpireOptionsContext)

	// EnterDropUser is called when entering the dropUser production.
	EnterDropUser(c *DropUserContext)

	// EnterGrant is called when entering the grant production.
	EnterGrant(c *GrantContext)

	// EnterGrantTargetList is called when entering the grantTargetList production.
	EnterGrantTargetList(c *GrantTargetListContext)

	// EnterGrantOptions is called when entering the grantOptions production.
	EnterGrantOptions(c *GrantOptionsContext)

	// EnterExceptRoleList is called when entering the exceptRoleList production.
	EnterExceptRoleList(c *ExceptRoleListContext)

	// EnterWithRoles is called when entering the withRoles production.
	EnterWithRoles(c *WithRolesContext)

	// EnterGrantAs is called when entering the grantAs production.
	EnterGrantAs(c *GrantAsContext)

	// EnterVersionedRequireClause is called when entering the versionedRequireClause production.
	EnterVersionedRequireClause(c *VersionedRequireClauseContext)

	// EnterRenameUser is called when entering the renameUser production.
	EnterRenameUser(c *RenameUserContext)

	// EnterRevoke is called when entering the revoke production.
	EnterRevoke(c *RevokeContext)

	// EnterOnTypeTo is called when entering the onTypeTo production.
	EnterOnTypeTo(c *OnTypeToContext)

	// EnterAclType is called when entering the aclType production.
	EnterAclType(c *AclTypeContext)

	// EnterRoleOrPrivilegesList is called when entering the roleOrPrivilegesList production.
	EnterRoleOrPrivilegesList(c *RoleOrPrivilegesListContext)

	// EnterRoleOrPrivilege is called when entering the roleOrPrivilege production.
	EnterRoleOrPrivilege(c *RoleOrPrivilegeContext)

	// EnterGrantIdentifier is called when entering the grantIdentifier production.
	EnterGrantIdentifier(c *GrantIdentifierContext)

	// EnterRequireList is called when entering the requireList production.
	EnterRequireList(c *RequireListContext)

	// EnterRequireListElement is called when entering the requireListElement production.
	EnterRequireListElement(c *RequireListElementContext)

	// EnterGrantOption is called when entering the grantOption production.
	EnterGrantOption(c *GrantOptionContext)

	// EnterSetRole is called when entering the setRole production.
	EnterSetRole(c *SetRoleContext)

	// EnterRoleList is called when entering the roleList production.
	EnterRoleList(c *RoleListContext)

	// EnterRole is called when entering the role production.
	EnterRole(c *RoleContext)

	// EnterTableAdministrationStatement is called when entering the tableAdministrationStatement production.
	EnterTableAdministrationStatement(c *TableAdministrationStatementContext)

	// EnterHistogram is called when entering the histogram production.
	EnterHistogram(c *HistogramContext)

	// EnterCheckOption is called when entering the checkOption production.
	EnterCheckOption(c *CheckOptionContext)

	// EnterRepairType is called when entering the repairType production.
	EnterRepairType(c *RepairTypeContext)

	// EnterInstallUninstallStatment is called when entering the installUninstallStatment production.
	EnterInstallUninstallStatment(c *InstallUninstallStatmentContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterStartOptionValueList is called when entering the startOptionValueList production.
	EnterStartOptionValueList(c *StartOptionValueListContext)

	// EnterTransactionCharacteristics is called when entering the transactionCharacteristics production.
	EnterTransactionCharacteristics(c *TransactionCharacteristicsContext)

	// EnterTransactionAccessMode is called when entering the transactionAccessMode production.
	EnterTransactionAccessMode(c *TransactionAccessModeContext)

	// EnterIsolationLevel is called when entering the isolationLevel production.
	EnterIsolationLevel(c *IsolationLevelContext)

	// EnterOptionValueListContinued is called when entering the optionValueListContinued production.
	EnterOptionValueListContinued(c *OptionValueListContinuedContext)

	// EnterOptionValueNoOptionType is called when entering the optionValueNoOptionType production.
	EnterOptionValueNoOptionType(c *OptionValueNoOptionTypeContext)

	// EnterOptionValue is called when entering the optionValue production.
	EnterOptionValue(c *OptionValueContext)

	// EnterSetSystemVariable is called when entering the setSystemVariable production.
	EnterSetSystemVariable(c *SetSystemVariableContext)

	// EnterStartOptionValueListFollowingOptionType is called when entering the startOptionValueListFollowingOptionType production.
	EnterStartOptionValueListFollowingOptionType(c *StartOptionValueListFollowingOptionTypeContext)

	// EnterOptionValueFollowingOptionType is called when entering the optionValueFollowingOptionType production.
	EnterOptionValueFollowingOptionType(c *OptionValueFollowingOptionTypeContext)

	// EnterSetExprOrDefault is called when entering the setExprOrDefault production.
	EnterSetExprOrDefault(c *SetExprOrDefaultContext)

	// EnterShowStatement is called when entering the showStatement production.
	EnterShowStatement(c *ShowStatementContext)

	// EnterShowCommandType is called when entering the showCommandType production.
	EnterShowCommandType(c *ShowCommandTypeContext)

	// EnterNonBlocking is called when entering the nonBlocking production.
	EnterNonBlocking(c *NonBlockingContext)

	// EnterFromOrIn is called when entering the fromOrIn production.
	EnterFromOrIn(c *FromOrInContext)

	// EnterInDb is called when entering the inDb production.
	EnterInDb(c *InDbContext)

	// EnterProfileType is called when entering the profileType production.
	EnterProfileType(c *ProfileTypeContext)

	// EnterOtherAdministrativeStatement is called when entering the otherAdministrativeStatement production.
	EnterOtherAdministrativeStatement(c *OtherAdministrativeStatementContext)

	// EnterKeyCacheListOrParts is called when entering the keyCacheListOrParts production.
	EnterKeyCacheListOrParts(c *KeyCacheListOrPartsContext)

	// EnterKeyCacheList is called when entering the keyCacheList production.
	EnterKeyCacheList(c *KeyCacheListContext)

	// EnterAssignToKeycache is called when entering the assignToKeycache production.
	EnterAssignToKeycache(c *AssignToKeycacheContext)

	// EnterAssignToKeycachePartition is called when entering the assignToKeycachePartition production.
	EnterAssignToKeycachePartition(c *AssignToKeycachePartitionContext)

	// EnterCacheKeyList is called when entering the cacheKeyList production.
	EnterCacheKeyList(c *CacheKeyListContext)

	// EnterKeyUsageElement is called when entering the keyUsageElement production.
	EnterKeyUsageElement(c *KeyUsageElementContext)

	// EnterKeyUsageList is called when entering the keyUsageList production.
	EnterKeyUsageList(c *KeyUsageListContext)

	// EnterFlushOption is called when entering the flushOption production.
	EnterFlushOption(c *FlushOptionContext)

	// EnterLogType is called when entering the logType production.
	EnterLogType(c *LogTypeContext)

	// EnterFlushTables is called when entering the flushTables production.
	EnterFlushTables(c *FlushTablesContext)

	// EnterFlushTablesOptions is called when entering the flushTablesOptions production.
	EnterFlushTablesOptions(c *FlushTablesOptionsContext)

	// EnterPreloadTail is called when entering the preloadTail production.
	EnterPreloadTail(c *PreloadTailContext)

	// EnterPreloadList is called when entering the preloadList production.
	EnterPreloadList(c *PreloadListContext)

	// EnterPreloadKeys is called when entering the preloadKeys production.
	EnterPreloadKeys(c *PreloadKeysContext)

	// EnterAdminPartition is called when entering the adminPartition production.
	EnterAdminPartition(c *AdminPartitionContext)

	// EnterResourceGroupManagement is called when entering the resourceGroupManagement production.
	EnterResourceGroupManagement(c *ResourceGroupManagementContext)

	// EnterCreateResourceGroup is called when entering the createResourceGroup production.
	EnterCreateResourceGroup(c *CreateResourceGroupContext)

	// EnterResourceGroupVcpuList is called when entering the resourceGroupVcpuList production.
	EnterResourceGroupVcpuList(c *ResourceGroupVcpuListContext)

	// EnterVcpuNumOrRange is called when entering the vcpuNumOrRange production.
	EnterVcpuNumOrRange(c *VcpuNumOrRangeContext)

	// EnterResourceGroupPriority is called when entering the resourceGroupPriority production.
	EnterResourceGroupPriority(c *ResourceGroupPriorityContext)

	// EnterResourceGroupEnableDisable is called when entering the resourceGroupEnableDisable production.
	EnterResourceGroupEnableDisable(c *ResourceGroupEnableDisableContext)

	// EnterAlterResourceGroup is called when entering the alterResourceGroup production.
	EnterAlterResourceGroup(c *AlterResourceGroupContext)

	// EnterSetResourceGroup is called when entering the setResourceGroup production.
	EnterSetResourceGroup(c *SetResourceGroupContext)

	// EnterThreadIdList is called when entering the threadIdList production.
	EnterThreadIdList(c *ThreadIdListContext)

	// EnterDropResourceGroup is called when entering the dropResourceGroup production.
	EnterDropResourceGroup(c *DropResourceGroupContext)

	// EnterUtilityStatement is called when entering the utilityStatement production.
	EnterUtilityStatement(c *UtilityStatementContext)

	// EnterDescribeStatement is called when entering the describeStatement production.
	EnterDescribeStatement(c *DescribeStatementContext)

	// EnterExplainStatement is called when entering the explainStatement production.
	EnterExplainStatement(c *ExplainStatementContext)

	// EnterExplainableStatement is called when entering the explainableStatement production.
	EnterExplainableStatement(c *ExplainableStatementContext)

	// EnterHelpCommand is called when entering the helpCommand production.
	EnterHelpCommand(c *HelpCommandContext)

	// EnterUseCommand is called when entering the useCommand production.
	EnterUseCommand(c *UseCommandContext)

	// EnterRestartServer is called when entering the restartServer production.
	EnterRestartServer(c *RestartServerContext)

	// EnterExprOr is called when entering the exprOr production.
	EnterExprOr(c *ExprOrContext)

	// EnterExprNot is called when entering the exprNot production.
	EnterExprNot(c *ExprNotContext)

	// EnterExprIs is called when entering the exprIs production.
	EnterExprIs(c *ExprIsContext)

	// EnterExprAnd is called when entering the exprAnd production.
	EnterExprAnd(c *ExprAndContext)

	// EnterExprXor is called when entering the exprXor production.
	EnterExprXor(c *ExprXorContext)

	// EnterPrimaryExprPredicate is called when entering the primaryExprPredicate production.
	EnterPrimaryExprPredicate(c *PrimaryExprPredicateContext)

	// EnterPrimaryExprCompare is called when entering the primaryExprCompare production.
	EnterPrimaryExprCompare(c *PrimaryExprCompareContext)

	// EnterPrimaryExprAllAny is called when entering the primaryExprAllAny production.
	EnterPrimaryExprAllAny(c *PrimaryExprAllAnyContext)

	// EnterPrimaryExprIsNull is called when entering the primaryExprIsNull production.
	EnterPrimaryExprIsNull(c *PrimaryExprIsNullContext)

	// EnterCompOp is called when entering the compOp production.
	EnterCompOp(c *CompOpContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterPredicateExprIn is called when entering the predicateExprIn production.
	EnterPredicateExprIn(c *PredicateExprInContext)

	// EnterPredicateExprBetween is called when entering the predicateExprBetween production.
	EnterPredicateExprBetween(c *PredicateExprBetweenContext)

	// EnterPredicateExprLike is called when entering the predicateExprLike production.
	EnterPredicateExprLike(c *PredicateExprLikeContext)

	// EnterPredicateExprRegex is called when entering the predicateExprRegex production.
	EnterPredicateExprRegex(c *PredicateExprRegexContext)

	// EnterBitExpr is called when entering the bitExpr production.
	EnterBitExpr(c *BitExprContext)

	// EnterSimpleExprConvert is called when entering the simpleExprConvert production.
	EnterSimpleExprConvert(c *SimpleExprConvertContext)

	// EnterSimpleExprSearchJson is called when entering the simpleExprSearchJson production.
	EnterSimpleExprSearchJson(c *SimpleExprSearchJsonContext)

	// EnterSimpleExprVariable is called when entering the simpleExprVariable production.
	EnterSimpleExprVariable(c *SimpleExprVariableContext)

	// EnterSimpleExprCast is called when entering the simpleExprCast production.
	EnterSimpleExprCast(c *SimpleExprCastContext)

	// EnterSimpleExprUnary is called when entering the simpleExprUnary production.
	EnterSimpleExprUnary(c *SimpleExprUnaryContext)

	// EnterSimpleExprOdbc is called when entering the simpleExprOdbc production.
	EnterSimpleExprOdbc(c *SimpleExprOdbcContext)

	// EnterSimpleExprRuntimeFunction is called when entering the simpleExprRuntimeFunction production.
	EnterSimpleExprRuntimeFunction(c *SimpleExprRuntimeFunctionContext)

	// EnterSimpleExprFunction is called when entering the simpleExprFunction production.
	EnterSimpleExprFunction(c *SimpleExprFunctionContext)

	// EnterSimpleExprCollate is called when entering the simpleExprCollate production.
	EnterSimpleExprCollate(c *SimpleExprCollateContext)

	// EnterSimpleExprMatch is called when entering the simpleExprMatch production.
	EnterSimpleExprMatch(c *SimpleExprMatchContext)

	// EnterSimpleExprWindowingFunction is called when entering the simpleExprWindowingFunction production.
	EnterSimpleExprWindowingFunction(c *SimpleExprWindowingFunctionContext)

	// EnterSimpleExprBinary is called when entering the simpleExprBinary production.
	EnterSimpleExprBinary(c *SimpleExprBinaryContext)

	// EnterSimpleExprColumnRef is called when entering the simpleExprColumnRef production.
	EnterSimpleExprColumnRef(c *SimpleExprColumnRefContext)

	// EnterSimpleExprParamMarker is called when entering the simpleExprParamMarker production.
	EnterSimpleExprParamMarker(c *SimpleExprParamMarkerContext)

	// EnterSimpleExprSum is called when entering the simpleExprSum production.
	EnterSimpleExprSum(c *SimpleExprSumContext)

	// EnterSimpleExprConvertUsing is called when entering the simpleExprConvertUsing production.
	EnterSimpleExprConvertUsing(c *SimpleExprConvertUsingContext)

	// EnterSimpleExprSubQuery is called when entering the simpleExprSubQuery production.
	EnterSimpleExprSubQuery(c *SimpleExprSubQueryContext)

	// EnterSimpleExprGroupingOperation is called when entering the simpleExprGroupingOperation production.
	EnterSimpleExprGroupingOperation(c *SimpleExprGroupingOperationContext)

	// EnterSimpleExprNot is called when entering the simpleExprNot production.
	EnterSimpleExprNot(c *SimpleExprNotContext)

	// EnterSimpleExprValues is called when entering the simpleExprValues production.
	EnterSimpleExprValues(c *SimpleExprValuesContext)

	// EnterSimpleExprDefault is called when entering the simpleExprDefault production.
	EnterSimpleExprDefault(c *SimpleExprDefaultContext)

	// EnterSimpleExprList is called when entering the simpleExprList production.
	EnterSimpleExprList(c *SimpleExprListContext)

	// EnterSimpleExprInterval is called when entering the simpleExprInterval production.
	EnterSimpleExprInterval(c *SimpleExprIntervalContext)

	// EnterSimpleExprCase is called when entering the simpleExprCase production.
	EnterSimpleExprCase(c *SimpleExprCaseContext)

	// EnterSimpleExprConcat is called when entering the simpleExprConcat production.
	EnterSimpleExprConcat(c *SimpleExprConcatContext)

	// EnterSimpleExprLiteral is called when entering the simpleExprLiteral production.
	EnterSimpleExprLiteral(c *SimpleExprLiteralContext)

	// EnterArrayCast is called when entering the arrayCast production.
	EnterArrayCast(c *ArrayCastContext)

	// EnterJsonOperator is called when entering the jsonOperator production.
	EnterJsonOperator(c *JsonOperatorContext)

	// EnterSumExpr is called when entering the sumExpr production.
	EnterSumExpr(c *SumExprContext)

	// EnterGroupingOperation is called when entering the groupingOperation production.
	EnterGroupingOperation(c *GroupingOperationContext)

	// EnterWindowFunctionCall is called when entering the windowFunctionCall production.
	EnterWindowFunctionCall(c *WindowFunctionCallContext)

	// EnterWindowingClause is called when entering the windowingClause production.
	EnterWindowingClause(c *WindowingClauseContext)

	// EnterLeadLagInfo is called when entering the leadLagInfo production.
	EnterLeadLagInfo(c *LeadLagInfoContext)

	// EnterNullTreatment is called when entering the nullTreatment production.
	EnterNullTreatment(c *NullTreatmentContext)

	// EnterJsonFunction is called when entering the jsonFunction production.
	EnterJsonFunction(c *JsonFunctionContext)

	// EnterInSumExpr is called when entering the inSumExpr production.
	EnterInSumExpr(c *InSumExprContext)

	// EnterIdentListArg is called when entering the identListArg production.
	EnterIdentListArg(c *IdentListArgContext)

	// EnterIdentList is called when entering the identList production.
	EnterIdentList(c *IdentListContext)

	// EnterFulltextOptions is called when entering the fulltextOptions production.
	EnterFulltextOptions(c *FulltextOptionsContext)

	// EnterRuntimeFunctionCall is called when entering the runtimeFunctionCall production.
	EnterRuntimeFunctionCall(c *RuntimeFunctionCallContext)

	// EnterGeometryFunction is called when entering the geometryFunction production.
	EnterGeometryFunction(c *GeometryFunctionContext)

	// EnterTimeFunctionParameters is called when entering the timeFunctionParameters production.
	EnterTimeFunctionParameters(c *TimeFunctionParametersContext)

	// EnterFractionalPrecision is called when entering the fractionalPrecision production.
	EnterFractionalPrecision(c *FractionalPrecisionContext)

	// EnterWeightStringLevels is called when entering the weightStringLevels production.
	EnterWeightStringLevels(c *WeightStringLevelsContext)

	// EnterWeightStringLevelListItem is called when entering the weightStringLevelListItem production.
	EnterWeightStringLevelListItem(c *WeightStringLevelListItemContext)

	// EnterDateTimeTtype is called when entering the dateTimeTtype production.
	EnterDateTimeTtype(c *DateTimeTtypeContext)

	// EnterTrimFunction is called when entering the trimFunction production.
	EnterTrimFunction(c *TrimFunctionContext)

	// EnterSubstringFunction is called when entering the substringFunction production.
	EnterSubstringFunction(c *SubstringFunctionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterSearchJsonFunction is called when entering the searchJsonFunction production.
	EnterSearchJsonFunction(c *SearchJsonFunctionContext)

	// EnterJsonValueReturning is called when entering the jsonValueReturning production.
	EnterJsonValueReturning(c *JsonValueReturningContext)

	// EnterJsonValueOnEmpty is called when entering the jsonValueOnEmpty production.
	EnterJsonValueOnEmpty(c *JsonValueOnEmptyContext)

	// EnterJsonValueOnError is called when entering the jsonValueOnError production.
	EnterJsonValueOnError(c *JsonValueOnErrorContext)

	// EnterUdfExprList is called when entering the udfExprList production.
	EnterUdfExprList(c *UdfExprListContext)

	// EnterUdfExpr is called when entering the udfExpr production.
	EnterUdfExpr(c *UdfExprContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterUserVariable is called when entering the userVariable production.
	EnterUserVariable(c *UserVariableContext)

	// EnterSystemVariable is called when entering the systemVariable production.
	EnterSystemVariable(c *SystemVariableContext)

	// EnterInternalVariableName is called when entering the internalVariableName production.
	EnterInternalVariableName(c *InternalVariableNameContext)

	// EnterWhenExpression is called when entering the whenExpression production.
	EnterWhenExpression(c *WhenExpressionContext)

	// EnterThenExpression is called when entering the thenExpression production.
	EnterThenExpression(c *ThenExpressionContext)

	// EnterElseExpression is called when entering the elseExpression production.
	EnterElseExpression(c *ElseExpressionContext)

	// EnterCastType is called when entering the castType production.
	EnterCastType(c *CastTypeContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterCharset is called when entering the charset production.
	EnterCharset(c *CharsetContext)

	// EnterNotRule is called when entering the notRule production.
	EnterNotRule(c *NotRuleContext)

	// EnterNot2Rule is called when entering the not2Rule production.
	EnterNot2Rule(c *Not2RuleContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterIntervalTimeStamp is called when entering the intervalTimeStamp production.
	EnterIntervalTimeStamp(c *IntervalTimeStampContext)

	// EnterExprListWithParentheses is called when entering the exprListWithParentheses production.
	EnterExprListWithParentheses(c *ExprListWithParenthesesContext)

	// EnterExprWithParentheses is called when entering the exprWithParentheses production.
	EnterExprWithParentheses(c *ExprWithParenthesesContext)

	// EnterSimpleExprWithParentheses is called when entering the simpleExprWithParentheses production.
	EnterSimpleExprWithParentheses(c *SimpleExprWithParenthesesContext)

	// EnterOrderList is called when entering the orderList production.
	EnterOrderList(c *OrderListContext)

	// EnterOrderExpression is called when entering the orderExpression production.
	EnterOrderExpression(c *OrderExpressionContext)

	// EnterGroupList is called when entering the groupList production.
	EnterGroupList(c *GroupListContext)

	// EnterGroupingExpression is called when entering the groupingExpression production.
	EnterGroupingExpression(c *GroupingExpressionContext)

	// EnterChannel is called when entering the channel production.
	EnterChannel(c *ChannelContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfBody is called when entering the ifBody production.
	EnterIfBody(c *IfBodyContext)

	// EnterThenStatement is called when entering the thenStatement production.
	EnterThenStatement(c *ThenStatementContext)

	// EnterCompoundStatementList is called when entering the compoundStatementList production.
	EnterCompoundStatementList(c *CompoundStatementListContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterLabeledBlock is called when entering the labeledBlock production.
	EnterLabeledBlock(c *LabeledBlockContext)

	// EnterUnlabeledBlock is called when entering the unlabeledBlock production.
	EnterUnlabeledBlock(c *UnlabeledBlockContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterBeginEndBlock is called when entering the beginEndBlock production.
	EnterBeginEndBlock(c *BeginEndBlockContext)

	// EnterLabeledControl is called when entering the labeledControl production.
	EnterLabeledControl(c *LabeledControlContext)

	// EnterUnlabeledControl is called when entering the unlabeledControl production.
	EnterUnlabeledControl(c *UnlabeledControlContext)

	// EnterLoopBlock is called when entering the loopBlock production.
	EnterLoopBlock(c *LoopBlockContext)

	// EnterWhileDoBlock is called when entering the whileDoBlock production.
	EnterWhileDoBlock(c *WhileDoBlockContext)

	// EnterRepeatUntilBlock is called when entering the repeatUntilBlock production.
	EnterRepeatUntilBlock(c *RepeatUntilBlockContext)

	// EnterSpDeclarations is called when entering the spDeclarations production.
	EnterSpDeclarations(c *SpDeclarationsContext)

	// EnterSpDeclaration is called when entering the spDeclaration production.
	EnterSpDeclaration(c *SpDeclarationContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterConditionDeclaration is called when entering the conditionDeclaration production.
	EnterConditionDeclaration(c *ConditionDeclarationContext)

	// EnterSpCondition is called when entering the spCondition production.
	EnterSpCondition(c *SpConditionContext)

	// EnterSqlstate is called when entering the sqlstate production.
	EnterSqlstate(c *SqlstateContext)

	// EnterHandlerDeclaration is called when entering the handlerDeclaration production.
	EnterHandlerDeclaration(c *HandlerDeclarationContext)

	// EnterHandlerCondition is called when entering the handlerCondition production.
	EnterHandlerCondition(c *HandlerConditionContext)

	// EnterCursorDeclaration is called when entering the cursorDeclaration production.
	EnterCursorDeclaration(c *CursorDeclarationContext)

	// EnterIterateStatement is called when entering the iterateStatement production.
	EnterIterateStatement(c *IterateStatementContext)

	// EnterLeaveStatement is called when entering the leaveStatement production.
	EnterLeaveStatement(c *LeaveStatementContext)

	// EnterGetDiagnostics is called when entering the getDiagnostics production.
	EnterGetDiagnostics(c *GetDiagnosticsContext)

	// EnterSignalAllowedExpr is called when entering the signalAllowedExpr production.
	EnterSignalAllowedExpr(c *SignalAllowedExprContext)

	// EnterStatementInformationItem is called when entering the statementInformationItem production.
	EnterStatementInformationItem(c *StatementInformationItemContext)

	// EnterConditionInformationItem is called when entering the conditionInformationItem production.
	EnterConditionInformationItem(c *ConditionInformationItemContext)

	// EnterSignalInformationItemName is called when entering the signalInformationItemName production.
	EnterSignalInformationItemName(c *SignalInformationItemNameContext)

	// EnterSignalStatement is called when entering the signalStatement production.
	EnterSignalStatement(c *SignalStatementContext)

	// EnterResignalStatement is called when entering the resignalStatement production.
	EnterResignalStatement(c *ResignalStatementContext)

	// EnterSignalInformationItem is called when entering the signalInformationItem production.
	EnterSignalInformationItem(c *SignalInformationItemContext)

	// EnterCursorOpen is called when entering the cursorOpen production.
	EnterCursorOpen(c *CursorOpenContext)

	// EnterCursorClose is called when entering the cursorClose production.
	EnterCursorClose(c *CursorCloseContext)

	// EnterCursorFetch is called when entering the cursorFetch production.
	EnterCursorFetch(c *CursorFetchContext)

	// EnterSchedule is called when entering the schedule production.
	EnterSchedule(c *ScheduleContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterCheckOrReferences is called when entering the checkOrReferences production.
	EnterCheckOrReferences(c *CheckOrReferencesContext)

	// EnterCheckConstraint is called when entering the checkConstraint production.
	EnterCheckConstraint(c *CheckConstraintContext)

	// EnterConstraintEnforcement is called when entering the constraintEnforcement production.
	EnterConstraintEnforcement(c *ConstraintEnforcementContext)

	// EnterTableConstraintDef is called when entering the tableConstraintDef production.
	EnterTableConstraintDef(c *TableConstraintDefContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterFieldDefinition is called when entering the fieldDefinition production.
	EnterFieldDefinition(c *FieldDefinitionContext)

	// EnterColumnAttribute is called when entering the columnAttribute production.
	EnterColumnAttribute(c *ColumnAttributeContext)

	// EnterColumnFormat is called when entering the columnFormat production.
	EnterColumnFormat(c *ColumnFormatContext)

	// EnterStorageMedia is called when entering the storageMedia production.
	EnterStorageMedia(c *StorageMediaContext)

	// EnterGcolAttribute is called when entering the gcolAttribute production.
	EnterGcolAttribute(c *GcolAttributeContext)

	// EnterReferences is called when entering the references production.
	EnterReferences(c *ReferencesContext)

	// EnterDeleteOption is called when entering the deleteOption production.
	EnterDeleteOption(c *DeleteOptionContext)

	// EnterKeyList is called when entering the keyList production.
	EnterKeyList(c *KeyListContext)

	// EnterKeyPart is called when entering the keyPart production.
	EnterKeyPart(c *KeyPartContext)

	// EnterKeyListWithExpression is called when entering the keyListWithExpression production.
	EnterKeyListWithExpression(c *KeyListWithExpressionContext)

	// EnterKeyPartOrExpression is called when entering the keyPartOrExpression production.
	EnterKeyPartOrExpression(c *KeyPartOrExpressionContext)

	// EnterKeyListVariants is called when entering the keyListVariants production.
	EnterKeyListVariants(c *KeyListVariantsContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterIndexOption is called when entering the indexOption production.
	EnterIndexOption(c *IndexOptionContext)

	// EnterCommonIndexOption is called when entering the commonIndexOption production.
	EnterCommonIndexOption(c *CommonIndexOptionContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// EnterIndexTypeClause is called when entering the indexTypeClause production.
	EnterIndexTypeClause(c *IndexTypeClauseContext)

	// EnterFulltextIndexOption is called when entering the fulltextIndexOption production.
	EnterFulltextIndexOption(c *FulltextIndexOptionContext)

	// EnterSpatialIndexOption is called when entering the spatialIndexOption production.
	EnterSpatialIndexOption(c *SpatialIndexOptionContext)

	// EnterDataTypeDefinition is called when entering the dataTypeDefinition production.
	EnterDataTypeDefinition(c *DataTypeDefinitionContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterNchar is called when entering the nchar production.
	EnterNchar(c *NcharContext)

	// EnterRealType is called when entering the realType production.
	EnterRealType(c *RealTypeContext)

	// EnterFieldLength is called when entering the fieldLength production.
	EnterFieldLength(c *FieldLengthContext)

	// EnterFieldOptions is called when entering the fieldOptions production.
	EnterFieldOptions(c *FieldOptionsContext)

	// EnterCharsetWithOptBinary is called when entering the charsetWithOptBinary production.
	EnterCharsetWithOptBinary(c *CharsetWithOptBinaryContext)

	// EnterAscii is called when entering the ascii production.
	EnterAscii(c *AsciiContext)

	// EnterUnicode is called when entering the unicode production.
	EnterUnicode(c *UnicodeContext)

	// EnterWsNumCodepoints is called when entering the wsNumCodepoints production.
	EnterWsNumCodepoints(c *WsNumCodepointsContext)

	// EnterTypeDatetimePrecision is called when entering the typeDatetimePrecision production.
	EnterTypeDatetimePrecision(c *TypeDatetimePrecisionContext)

	// EnterCharsetName is called when entering the charsetName production.
	EnterCharsetName(c *CharsetNameContext)

	// EnterCollationName is called when entering the collationName production.
	EnterCollationName(c *CollationNameContext)

	// EnterCreateTableOptions is called when entering the createTableOptions production.
	EnterCreateTableOptions(c *CreateTableOptionsContext)

	// EnterCreateTableOptionsSpaceSeparated is called when entering the createTableOptionsSpaceSeparated production.
	EnterCreateTableOptionsSpaceSeparated(c *CreateTableOptionsSpaceSeparatedContext)

	// EnterCreateTableOption is called when entering the createTableOption production.
	EnterCreateTableOption(c *CreateTableOptionContext)

	// EnterTernaryOption is called when entering the ternaryOption production.
	EnterTernaryOption(c *TernaryOptionContext)

	// EnterDefaultCollation is called when entering the defaultCollation production.
	EnterDefaultCollation(c *DefaultCollationContext)

	// EnterDefaultEncryption is called when entering the defaultEncryption production.
	EnterDefaultEncryption(c *DefaultEncryptionContext)

	// EnterDefaultCharset is called when entering the defaultCharset production.
	EnterDefaultCharset(c *DefaultCharsetContext)

	// EnterPartitionClause is called when entering the partitionClause production.
	EnterPartitionClause(c *PartitionClauseContext)

	// EnterPartitionDefKey is called when entering the partitionDefKey production.
	EnterPartitionDefKey(c *PartitionDefKeyContext)

	// EnterPartitionDefHash is called when entering the partitionDefHash production.
	EnterPartitionDefHash(c *PartitionDefHashContext)

	// EnterPartitionDefRangeList is called when entering the partitionDefRangeList production.
	EnterPartitionDefRangeList(c *PartitionDefRangeListContext)

	// EnterSubPartitions is called when entering the subPartitions production.
	EnterSubPartitions(c *SubPartitionsContext)

	// EnterPartitionKeyAlgorithm is called when entering the partitionKeyAlgorithm production.
	EnterPartitionKeyAlgorithm(c *PartitionKeyAlgorithmContext)

	// EnterPartitionDefinitions is called when entering the partitionDefinitions production.
	EnterPartitionDefinitions(c *PartitionDefinitionsContext)

	// EnterPartitionDefinition is called when entering the partitionDefinition production.
	EnterPartitionDefinition(c *PartitionDefinitionContext)

	// EnterPartitionValuesIn is called when entering the partitionValuesIn production.
	EnterPartitionValuesIn(c *PartitionValuesInContext)

	// EnterPartitionOption is called when entering the partitionOption production.
	EnterPartitionOption(c *PartitionOptionContext)

	// EnterSubpartitionDefinition is called when entering the subpartitionDefinition production.
	EnterSubpartitionDefinition(c *SubpartitionDefinitionContext)

	// EnterPartitionValueItemListParen is called when entering the partitionValueItemListParen production.
	EnterPartitionValueItemListParen(c *PartitionValueItemListParenContext)

	// EnterPartitionValueItem is called when entering the partitionValueItem production.
	EnterPartitionValueItem(c *PartitionValueItemContext)

	// EnterDefinerClause is called when entering the definerClause production.
	EnterDefinerClause(c *DefinerClauseContext)

	// EnterIfExists is called when entering the ifExists production.
	EnterIfExists(c *IfExistsContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// EnterProcedureParameter is called when entering the procedureParameter production.
	EnterProcedureParameter(c *ProcedureParameterContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterCollate is called when entering the collate production.
	EnterCollate(c *CollateContext)

	// EnterTypeWithOptCollate is called when entering the typeWithOptCollate production.
	EnterTypeWithOptCollate(c *TypeWithOptCollateContext)

	// EnterSchemaIdentifierPair is called when entering the schemaIdentifierPair production.
	EnterSchemaIdentifierPair(c *SchemaIdentifierPairContext)

	// EnterViewRefList is called when entering the viewRefList production.
	EnterViewRefList(c *ViewRefListContext)

	// EnterUpdateList is called when entering the updateList production.
	EnterUpdateList(c *UpdateListContext)

	// EnterUpdateElement is called when entering the updateElement production.
	EnterUpdateElement(c *UpdateElementContext)

	// EnterCharsetClause is called when entering the charsetClause production.
	EnterCharsetClause(c *CharsetClauseContext)

	// EnterFieldsClause is called when entering the fieldsClause production.
	EnterFieldsClause(c *FieldsClauseContext)

	// EnterFieldTerm is called when entering the fieldTerm production.
	EnterFieldTerm(c *FieldTermContext)

	// EnterLinesClause is called when entering the linesClause production.
	EnterLinesClause(c *LinesClauseContext)

	// EnterLineTerm is called when entering the lineTerm production.
	EnterLineTerm(c *LineTermContext)

	// EnterUserList is called when entering the userList production.
	EnterUserList(c *UserListContext)

	// EnterCreateUserList is called when entering the createUserList production.
	EnterCreateUserList(c *CreateUserListContext)

	// EnterAlterUserList is called when entering the alterUserList production.
	EnterAlterUserList(c *AlterUserListContext)

	// EnterCreateUserEntry is called when entering the createUserEntry production.
	EnterCreateUserEntry(c *CreateUserEntryContext)

	// EnterAlterUserEntry is called when entering the alterUserEntry production.
	EnterAlterUserEntry(c *AlterUserEntryContext)

	// EnterRetainCurrentPassword is called when entering the retainCurrentPassword production.
	EnterRetainCurrentPassword(c *RetainCurrentPasswordContext)

	// EnterDiscardOldPassword is called when entering the discardOldPassword production.
	EnterDiscardOldPassword(c *DiscardOldPasswordContext)

	// EnterReplacePassword is called when entering the replacePassword production.
	EnterReplacePassword(c *ReplacePasswordContext)

	// EnterUserIdentifierOrText is called when entering the userIdentifierOrText production.
	EnterUserIdentifierOrText(c *UserIdentifierOrTextContext)

	// EnterUser is called when entering the user production.
	EnterUser(c *UserContext)

	// EnterLikeClause is called when entering the likeClause production.
	EnterLikeClause(c *LikeClauseContext)

	// EnterLikeOrWhere is called when entering the likeOrWhere production.
	EnterLikeOrWhere(c *LikeOrWhereContext)

	// EnterOnlineOption is called when entering the onlineOption production.
	EnterOnlineOption(c *OnlineOptionContext)

	// EnterNoWriteToBinLog is called when entering the noWriteToBinLog production.
	EnterNoWriteToBinLog(c *NoWriteToBinLogContext)

	// EnterUsePartition is called when entering the usePartition production.
	EnterUsePartition(c *UsePartitionContext)

	// EnterFieldIdentifier is called when entering the fieldIdentifier production.
	EnterFieldIdentifier(c *FieldIdentifierContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterColumnInternalRef is called when entering the columnInternalRef production.
	EnterColumnInternalRef(c *ColumnInternalRefContext)

	// EnterColumnInternalRefList is called when entering the columnInternalRefList production.
	EnterColumnInternalRefList(c *ColumnInternalRefListContext)

	// EnterColumnRef is called when entering the columnRef production.
	EnterColumnRef(c *ColumnRefContext)

	// EnterInsertIdentifier is called when entering the insertIdentifier production.
	EnterInsertIdentifier(c *InsertIdentifierContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterIndexRef is called when entering the indexRef production.
	EnterIndexRef(c *IndexRefContext)

	// EnterTableWild is called when entering the tableWild production.
	EnterTableWild(c *TableWildContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterSchemaRef is called when entering the schemaRef production.
	EnterSchemaRef(c *SchemaRefContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterProcedureRef is called when entering the procedureRef production.
	EnterProcedureRef(c *ProcedureRefContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterFunctionRef is called when entering the functionRef production.
	EnterFunctionRef(c *FunctionRefContext)

	// EnterTriggerName is called when entering the triggerName production.
	EnterTriggerName(c *TriggerNameContext)

	// EnterTriggerRef is called when entering the triggerRef production.
	EnterTriggerRef(c *TriggerRefContext)

	// EnterViewName is called when entering the viewName production.
	EnterViewName(c *ViewNameContext)

	// EnterViewRef is called when entering the viewRef production.
	EnterViewRef(c *ViewRefContext)

	// EnterTablespaceName is called when entering the tablespaceName production.
	EnterTablespaceName(c *TablespaceNameContext)

	// EnterTablespaceRef is called when entering the tablespaceRef production.
	EnterTablespaceRef(c *TablespaceRefContext)

	// EnterLogfileGroupName is called when entering the logfileGroupName production.
	EnterLogfileGroupName(c *LogfileGroupNameContext)

	// EnterLogfileGroupRef is called when entering the logfileGroupRef production.
	EnterLogfileGroupRef(c *LogfileGroupRefContext)

	// EnterEventName is called when entering the eventName production.
	EnterEventName(c *EventNameContext)

	// EnterEventRef is called when entering the eventRef production.
	EnterEventRef(c *EventRefContext)

	// EnterUdfName is called when entering the udfName production.
	EnterUdfName(c *UdfNameContext)

	// EnterServerName is called when entering the serverName production.
	EnterServerName(c *ServerNameContext)

	// EnterServerRef is called when entering the serverRef production.
	EnterServerRef(c *ServerRefContext)

	// EnterEngineRef is called when entering the engineRef production.
	EnterEngineRef(c *EngineRefContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterFilterTableRef is called when entering the filterTableRef production.
	EnterFilterTableRef(c *FilterTableRefContext)

	// EnterTableRefWithWildcard is called when entering the tableRefWithWildcard production.
	EnterTableRefWithWildcard(c *TableRefWithWildcardContext)

	// EnterTableRef is called when entering the tableRef production.
	EnterTableRef(c *TableRefContext)

	// EnterTableRefList is called when entering the tableRefList production.
	EnterTableRefList(c *TableRefListContext)

	// EnterTableAliasRefList is called when entering the tableAliasRefList production.
	EnterTableAliasRefList(c *TableAliasRefListContext)

	// EnterParameterName is called when entering the parameterName production.
	EnterParameterName(c *ParameterNameContext)

	// EnterLabelIdentifier is called when entering the labelIdentifier production.
	EnterLabelIdentifier(c *LabelIdentifierContext)

	// EnterLabelRef is called when entering the labelRef production.
	EnterLabelRef(c *LabelRefContext)

	// EnterRoleIdentifier is called when entering the roleIdentifier production.
	EnterRoleIdentifier(c *RoleIdentifierContext)

	// EnterRoleRef is called when entering the roleRef production.
	EnterRoleRef(c *RoleRefContext)

	// EnterPluginRef is called when entering the pluginRef production.
	EnterPluginRef(c *PluginRefContext)

	// EnterComponentRef is called when entering the componentRef production.
	EnterComponentRef(c *ComponentRefContext)

	// EnterResourceGroupRef is called when entering the resourceGroupRef production.
	EnterResourceGroupRef(c *ResourceGroupRefContext)

	// EnterWindowName is called when entering the windowName production.
	EnterWindowName(c *WindowNameContext)

	// EnterPureIdentifier is called when entering the pureIdentifier production.
	EnterPureIdentifier(c *PureIdentifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterIdentifierListWithParentheses is called when entering the identifierListWithParentheses production.
	EnterIdentifierListWithParentheses(c *IdentifierListWithParenthesesContext)

	// EnterQualifiedIdentifier is called when entering the qualifiedIdentifier production.
	EnterQualifiedIdentifier(c *QualifiedIdentifierContext)

	// EnterSimpleIdentifier is called when entering the simpleIdentifier production.
	EnterSimpleIdentifier(c *SimpleIdentifierContext)

	// EnterDotIdentifier is called when entering the dotIdentifier production.
	EnterDotIdentifier(c *DotIdentifierContext)

	// EnterUlong_number is called when entering the ulong_number production.
	EnterUlong_number(c *Ulong_numberContext)

	// EnterReal_ulong_number is called when entering the real_ulong_number production.
	EnterReal_ulong_number(c *Real_ulong_numberContext)

	// EnterUlonglong_number is called when entering the ulonglong_number production.
	EnterUlonglong_number(c *Ulonglong_numberContext)

	// EnterReal_ulonglong_number is called when entering the real_ulonglong_number production.
	EnterReal_ulonglong_number(c *Real_ulonglong_numberContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterSignedLiteral is called when entering the signedLiteral production.
	EnterSignedLiteral(c *SignedLiteralContext)

	// EnterStringList is called when entering the stringList production.
	EnterStringList(c *StringListContext)

	// EnterTextStringLiteral is called when entering the textStringLiteral production.
	EnterTextStringLiteral(c *TextStringLiteralContext)

	// EnterTextString is called when entering the textString production.
	EnterTextString(c *TextStringContext)

	// EnterTextStringHash is called when entering the textStringHash production.
	EnterTextStringHash(c *TextStringHashContext)

	// EnterTextLiteral is called when entering the textLiteral production.
	EnterTextLiteral(c *TextLiteralContext)

	// EnterTextStringNoLinebreak is called when entering the textStringNoLinebreak production.
	EnterTextStringNoLinebreak(c *TextStringNoLinebreakContext)

	// EnterTextStringLiteralList is called when entering the textStringLiteralList production.
	EnterTextStringLiteralList(c *TextStringLiteralListContext)

	// EnterNumLiteral is called when entering the numLiteral production.
	EnterNumLiteral(c *NumLiteralContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterTemporalLiteral is called when entering the temporalLiteral production.
	EnterTemporalLiteral(c *TemporalLiteralContext)

	// EnterFloatOptions is called when entering the floatOptions production.
	EnterFloatOptions(c *FloatOptionsContext)

	// EnterStandardFloatOptions is called when entering the standardFloatOptions production.
	EnterStandardFloatOptions(c *StandardFloatOptionsContext)

	// EnterPrecision is called when entering the precision production.
	EnterPrecision(c *PrecisionContext)

	// EnterTextOrIdentifier is called when entering the textOrIdentifier production.
	EnterTextOrIdentifier(c *TextOrIdentifierContext)

	// EnterLValueIdentifier is called when entering the lValueIdentifier production.
	EnterLValueIdentifier(c *LValueIdentifierContext)

	// EnterRoleIdentifierOrText is called when entering the roleIdentifierOrText production.
	EnterRoleIdentifierOrText(c *RoleIdentifierOrTextContext)

	// EnterSizeNumber is called when entering the sizeNumber production.
	EnterSizeNumber(c *SizeNumberContext)

	// EnterParentheses is called when entering the parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterOptionType is called when entering the optionType production.
	EnterOptionType(c *OptionTypeContext)

	// EnterVarIdentType is called when entering the varIdentType production.
	EnterVarIdentType(c *VarIdentTypeContext)

	// EnterSetVarIdentType is called when entering the setVarIdentType production.
	EnterSetVarIdentType(c *SetVarIdentTypeContext)

	// EnterIdentifierKeyword is called when entering the identifierKeyword production.
	EnterIdentifierKeyword(c *IdentifierKeywordContext)

	// EnterIdentifierKeywordsAmbiguous1RolesAndLabels is called when entering the identifierKeywordsAmbiguous1RolesAndLabels production.
	EnterIdentifierKeywordsAmbiguous1RolesAndLabels(c *IdentifierKeywordsAmbiguous1RolesAndLabelsContext)

	// EnterIdentifierKeywordsAmbiguous2Labels is called when entering the identifierKeywordsAmbiguous2Labels production.
	EnterIdentifierKeywordsAmbiguous2Labels(c *IdentifierKeywordsAmbiguous2LabelsContext)

	// EnterLabelKeyword is called when entering the labelKeyword production.
	EnterLabelKeyword(c *LabelKeywordContext)

	// EnterIdentifierKeywordsAmbiguous3Roles is called when entering the identifierKeywordsAmbiguous3Roles production.
	EnterIdentifierKeywordsAmbiguous3Roles(c *IdentifierKeywordsAmbiguous3RolesContext)

	// EnterIdentifierKeywordsUnambiguous is called when entering the identifierKeywordsUnambiguous production.
	EnterIdentifierKeywordsUnambiguous(c *IdentifierKeywordsUnambiguousContext)

	// EnterRoleKeyword is called when entering the roleKeyword production.
	EnterRoleKeyword(c *RoleKeywordContext)

	// EnterLValueKeyword is called when entering the lValueKeyword production.
	EnterLValueKeyword(c *LValueKeywordContext)

	// EnterIdentifierKeywordsAmbiguous4SystemVariables is called when entering the identifierKeywordsAmbiguous4SystemVariables production.
	EnterIdentifierKeywordsAmbiguous4SystemVariables(c *IdentifierKeywordsAmbiguous4SystemVariablesContext)

	// EnterRoleOrIdentifierKeyword is called when entering the roleOrIdentifierKeyword production.
	EnterRoleOrIdentifierKeyword(c *RoleOrIdentifierKeywordContext)

	// EnterRoleOrLabelKeyword is called when entering the roleOrLabelKeyword production.
	EnterRoleOrLabelKeyword(c *RoleOrLabelKeywordContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitAlterStatement is called when exiting the alterStatement production.
	ExitAlterStatement(c *AlterStatementContext)

	// ExitAlterDatabase is called when exiting the alterDatabase production.
	ExitAlterDatabase(c *AlterDatabaseContext)

	// ExitAlterDatabaseOption is called when exiting the alterDatabaseOption production.
	ExitAlterDatabaseOption(c *AlterDatabaseOptionContext)

	// ExitAlterEvent is called when exiting the alterEvent production.
	ExitAlterEvent(c *AlterEventContext)

	// ExitAlterLogfileGroup is called when exiting the alterLogfileGroup production.
	ExitAlterLogfileGroup(c *AlterLogfileGroupContext)

	// ExitAlterLogfileGroupOptions is called when exiting the alterLogfileGroupOptions production.
	ExitAlterLogfileGroupOptions(c *AlterLogfileGroupOptionsContext)

	// ExitAlterLogfileGroupOption is called when exiting the alterLogfileGroupOption production.
	ExitAlterLogfileGroupOption(c *AlterLogfileGroupOptionContext)

	// ExitAlterServer is called when exiting the alterServer production.
	ExitAlterServer(c *AlterServerContext)

	// ExitAlterTable is called when exiting the alterTable production.
	ExitAlterTable(c *AlterTableContext)

	// ExitAlterTableActions is called when exiting the alterTableActions production.
	ExitAlterTableActions(c *AlterTableActionsContext)

	// ExitAlterCommandList is called when exiting the alterCommandList production.
	ExitAlterCommandList(c *AlterCommandListContext)

	// ExitAlterCommandsModifierList is called when exiting the alterCommandsModifierList production.
	ExitAlterCommandsModifierList(c *AlterCommandsModifierListContext)

	// ExitStandaloneAlterCommands is called when exiting the standaloneAlterCommands production.
	ExitStandaloneAlterCommands(c *StandaloneAlterCommandsContext)

	// ExitAlterPartition is called when exiting the alterPartition production.
	ExitAlterPartition(c *AlterPartitionContext)

	// ExitAlterList is called when exiting the alterList production.
	ExitAlterList(c *AlterListContext)

	// ExitAlterCommandsModifier is called when exiting the alterCommandsModifier production.
	ExitAlterCommandsModifier(c *AlterCommandsModifierContext)

	// ExitAlterListItem is called when exiting the alterListItem production.
	ExitAlterListItem(c *AlterListItemContext)

	// ExitPlace is called when exiting the place production.
	ExitPlace(c *PlaceContext)

	// ExitRestrict is called when exiting the restrict production.
	ExitRestrict(c *RestrictContext)

	// ExitAlterOrderList is called when exiting the alterOrderList production.
	ExitAlterOrderList(c *AlterOrderListContext)

	// ExitAlterAlgorithmOption is called when exiting the alterAlgorithmOption production.
	ExitAlterAlgorithmOption(c *AlterAlgorithmOptionContext)

	// ExitAlterLockOption is called when exiting the alterLockOption production.
	ExitAlterLockOption(c *AlterLockOptionContext)

	// ExitIndexLockAndAlgorithm is called when exiting the indexLockAndAlgorithm production.
	ExitIndexLockAndAlgorithm(c *IndexLockAndAlgorithmContext)

	// ExitWithValidation is called when exiting the withValidation production.
	ExitWithValidation(c *WithValidationContext)

	// ExitRemovePartitioning is called when exiting the removePartitioning production.
	ExitRemovePartitioning(c *RemovePartitioningContext)

	// ExitAllOrPartitionNameList is called when exiting the allOrPartitionNameList production.
	ExitAllOrPartitionNameList(c *AllOrPartitionNameListContext)

	// ExitAlterTablespace is called when exiting the alterTablespace production.
	ExitAlterTablespace(c *AlterTablespaceContext)

	// ExitAlterUndoTablespace is called when exiting the alterUndoTablespace production.
	ExitAlterUndoTablespace(c *AlterUndoTablespaceContext)

	// ExitUndoTableSpaceOptions is called when exiting the undoTableSpaceOptions production.
	ExitUndoTableSpaceOptions(c *UndoTableSpaceOptionsContext)

	// ExitUndoTableSpaceOption is called when exiting the undoTableSpaceOption production.
	ExitUndoTableSpaceOption(c *UndoTableSpaceOptionContext)

	// ExitAlterTablespaceOptions is called when exiting the alterTablespaceOptions production.
	ExitAlterTablespaceOptions(c *AlterTablespaceOptionsContext)

	// ExitAlterTablespaceOption is called when exiting the alterTablespaceOption production.
	ExitAlterTablespaceOption(c *AlterTablespaceOptionContext)

	// ExitChangeTablespaceOption is called when exiting the changeTablespaceOption production.
	ExitChangeTablespaceOption(c *ChangeTablespaceOptionContext)

	// ExitAlterView is called when exiting the alterView production.
	ExitAlterView(c *AlterViewContext)

	// ExitViewTail is called when exiting the viewTail production.
	ExitViewTail(c *ViewTailContext)

	// ExitViewSelect is called when exiting the viewSelect production.
	ExitViewSelect(c *ViewSelectContext)

	// ExitViewCheckOption is called when exiting the viewCheckOption production.
	ExitViewCheckOption(c *ViewCheckOptionContext)

	// ExitAlterProcedure is called when exiting the alterProcedure production.
	ExitAlterProcedure(c *AlterProcedureContext)

	// ExitAlterFunction is called when exiting the alterFunction production.
	ExitAlterFunction(c *AlterFunctionContext)

	// ExitCreateStatement is called when exiting the createStatement production.
	ExitCreateStatement(c *CreateStatementContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitCreateDatabaseOption is called when exiting the createDatabaseOption production.
	ExitCreateDatabaseOption(c *CreateDatabaseOptionContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitTableElementList is called when exiting the tableElementList production.
	ExitTableElementList(c *TableElementListContext)

	// ExitTableElement is called when exiting the tableElement production.
	ExitTableElement(c *TableElementContext)

	// ExitDuplicateAsQueryExpression is called when exiting the duplicateAsQueryExpression production.
	ExitDuplicateAsQueryExpression(c *DuplicateAsQueryExpressionContext)

	// ExitQueryExpressionOrParens is called when exiting the queryExpressionOrParens production.
	ExitQueryExpressionOrParens(c *QueryExpressionOrParensContext)

	// ExitCreateRoutine is called when exiting the createRoutine production.
	ExitCreateRoutine(c *CreateRoutineContext)

	// ExitCreateProcedure is called when exiting the createProcedure production.
	ExitCreateProcedure(c *CreateProcedureContext)

	// ExitCreateFunction is called when exiting the createFunction production.
	ExitCreateFunction(c *CreateFunctionContext)

	// ExitCreateUdf is called when exiting the createUdf production.
	ExitCreateUdf(c *CreateUdfContext)

	// ExitRoutineCreateOption is called when exiting the routineCreateOption production.
	ExitRoutineCreateOption(c *RoutineCreateOptionContext)

	// ExitRoutineAlterOptions is called when exiting the routineAlterOptions production.
	ExitRoutineAlterOptions(c *RoutineAlterOptionsContext)

	// ExitRoutineOption is called when exiting the routineOption production.
	ExitRoutineOption(c *RoutineOptionContext)

	// ExitCreateIndex is called when exiting the createIndex production.
	ExitCreateIndex(c *CreateIndexContext)

	// ExitIndexNameAndType is called when exiting the indexNameAndType production.
	ExitIndexNameAndType(c *IndexNameAndTypeContext)

	// ExitCreateIndexTarget is called when exiting the createIndexTarget production.
	ExitCreateIndexTarget(c *CreateIndexTargetContext)

	// ExitCreateLogfileGroup is called when exiting the createLogfileGroup production.
	ExitCreateLogfileGroup(c *CreateLogfileGroupContext)

	// ExitLogfileGroupOptions is called when exiting the logfileGroupOptions production.
	ExitLogfileGroupOptions(c *LogfileGroupOptionsContext)

	// ExitLogfileGroupOption is called when exiting the logfileGroupOption production.
	ExitLogfileGroupOption(c *LogfileGroupOptionContext)

	// ExitCreateServer is called when exiting the createServer production.
	ExitCreateServer(c *CreateServerContext)

	// ExitServerOptions is called when exiting the serverOptions production.
	ExitServerOptions(c *ServerOptionsContext)

	// ExitServerOption is called when exiting the serverOption production.
	ExitServerOption(c *ServerOptionContext)

	// ExitCreateTablespace is called when exiting the createTablespace production.
	ExitCreateTablespace(c *CreateTablespaceContext)

	// ExitCreateUndoTablespace is called when exiting the createUndoTablespace production.
	ExitCreateUndoTablespace(c *CreateUndoTablespaceContext)

	// ExitTsDataFileName is called when exiting the tsDataFileName production.
	ExitTsDataFileName(c *TsDataFileNameContext)

	// ExitTsDataFile is called when exiting the tsDataFile production.
	ExitTsDataFile(c *TsDataFileContext)

	// ExitTablespaceOptions is called when exiting the tablespaceOptions production.
	ExitTablespaceOptions(c *TablespaceOptionsContext)

	// ExitTablespaceOption is called when exiting the tablespaceOption production.
	ExitTablespaceOption(c *TablespaceOptionContext)

	// ExitTsOptionInitialSize is called when exiting the tsOptionInitialSize production.
	ExitTsOptionInitialSize(c *TsOptionInitialSizeContext)

	// ExitTsOptionUndoRedoBufferSize is called when exiting the tsOptionUndoRedoBufferSize production.
	ExitTsOptionUndoRedoBufferSize(c *TsOptionUndoRedoBufferSizeContext)

	// ExitTsOptionAutoextendSize is called when exiting the tsOptionAutoextendSize production.
	ExitTsOptionAutoextendSize(c *TsOptionAutoextendSizeContext)

	// ExitTsOptionMaxSize is called when exiting the tsOptionMaxSize production.
	ExitTsOptionMaxSize(c *TsOptionMaxSizeContext)

	// ExitTsOptionExtentSize is called when exiting the tsOptionExtentSize production.
	ExitTsOptionExtentSize(c *TsOptionExtentSizeContext)

	// ExitTsOptionNodegroup is called when exiting the tsOptionNodegroup production.
	ExitTsOptionNodegroup(c *TsOptionNodegroupContext)

	// ExitTsOptionEngine is called when exiting the tsOptionEngine production.
	ExitTsOptionEngine(c *TsOptionEngineContext)

	// ExitTsOptionWait is called when exiting the tsOptionWait production.
	ExitTsOptionWait(c *TsOptionWaitContext)

	// ExitTsOptionComment is called when exiting the tsOptionComment production.
	ExitTsOptionComment(c *TsOptionCommentContext)

	// ExitTsOptionFileblockSize is called when exiting the tsOptionFileblockSize production.
	ExitTsOptionFileblockSize(c *TsOptionFileblockSizeContext)

	// ExitTsOptionEncryption is called when exiting the tsOptionEncryption production.
	ExitTsOptionEncryption(c *TsOptionEncryptionContext)

	// ExitCreateView is called when exiting the createView production.
	ExitCreateView(c *CreateViewContext)

	// ExitViewReplaceOrAlgorithm is called when exiting the viewReplaceOrAlgorithm production.
	ExitViewReplaceOrAlgorithm(c *ViewReplaceOrAlgorithmContext)

	// ExitViewAlgorithm is called when exiting the viewAlgorithm production.
	ExitViewAlgorithm(c *ViewAlgorithmContext)

	// ExitViewSuid is called when exiting the viewSuid production.
	ExitViewSuid(c *ViewSuidContext)

	// ExitCreateTrigger is called when exiting the createTrigger production.
	ExitCreateTrigger(c *CreateTriggerContext)

	// ExitTriggerFollowsPrecedesClause is called when exiting the triggerFollowsPrecedesClause production.
	ExitTriggerFollowsPrecedesClause(c *TriggerFollowsPrecedesClauseContext)

	// ExitCreateEvent is called when exiting the createEvent production.
	ExitCreateEvent(c *CreateEventContext)

	// ExitCreateRole is called when exiting the createRole production.
	ExitCreateRole(c *CreateRoleContext)

	// ExitCreateSpatialReference is called when exiting the createSpatialReference production.
	ExitCreateSpatialReference(c *CreateSpatialReferenceContext)

	// ExitSrsAttribute is called when exiting the srsAttribute production.
	ExitSrsAttribute(c *SrsAttributeContext)

	// ExitDropStatement is called when exiting the dropStatement production.
	ExitDropStatement(c *DropStatementContext)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitDropEvent is called when exiting the dropEvent production.
	ExitDropEvent(c *DropEventContext)

	// ExitDropFunction is called when exiting the dropFunction production.
	ExitDropFunction(c *DropFunctionContext)

	// ExitDropProcedure is called when exiting the dropProcedure production.
	ExitDropProcedure(c *DropProcedureContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitDropLogfileGroup is called when exiting the dropLogfileGroup production.
	ExitDropLogfileGroup(c *DropLogfileGroupContext)

	// ExitDropLogfileGroupOption is called when exiting the dropLogfileGroupOption production.
	ExitDropLogfileGroupOption(c *DropLogfileGroupOptionContext)

	// ExitDropServer is called when exiting the dropServer production.
	ExitDropServer(c *DropServerContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropTableSpace is called when exiting the dropTableSpace production.
	ExitDropTableSpace(c *DropTableSpaceContext)

	// ExitDropTrigger is called when exiting the dropTrigger production.
	ExitDropTrigger(c *DropTriggerContext)

	// ExitDropView is called when exiting the dropView production.
	ExitDropView(c *DropViewContext)

	// ExitDropRole is called when exiting the dropRole production.
	ExitDropRole(c *DropRoleContext)

	// ExitDropSpatialReference is called when exiting the dropSpatialReference production.
	ExitDropSpatialReference(c *DropSpatialReferenceContext)

	// ExitDropUndoTablespace is called when exiting the dropUndoTablespace production.
	ExitDropUndoTablespace(c *DropUndoTablespaceContext)

	// ExitRenameTableStatement is called when exiting the renameTableStatement production.
	ExitRenameTableStatement(c *RenameTableStatementContext)

	// ExitRenamePair is called when exiting the renamePair production.
	ExitRenamePair(c *RenamePairContext)

	// ExitTruncateTableStatement is called when exiting the truncateTableStatement production.
	ExitTruncateTableStatement(c *TruncateTableStatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitCallStatement is called when exiting the callStatement production.
	ExitCallStatement(c *CallStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitPartitionDelete is called when exiting the partitionDelete production.
	ExitPartitionDelete(c *PartitionDeleteContext)

	// ExitDeleteStatementOption is called when exiting the deleteStatementOption production.
	ExitDeleteStatementOption(c *DeleteStatementOptionContext)

	// ExitDoStatement is called when exiting the doStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitHandlerStatement is called when exiting the handlerStatement production.
	ExitHandlerStatement(c *HandlerStatementContext)

	// ExitHandlerReadOrScan is called when exiting the handlerReadOrScan production.
	ExitHandlerReadOrScan(c *HandlerReadOrScanContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitInsertLockOption is called when exiting the insertLockOption production.
	ExitInsertLockOption(c *InsertLockOptionContext)

	// ExitInsertFromConstructor is called when exiting the insertFromConstructor production.
	ExitInsertFromConstructor(c *InsertFromConstructorContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitInsertValues is called when exiting the insertValues production.
	ExitInsertValues(c *InsertValuesContext)

	// ExitInsertQueryExpression is called when exiting the insertQueryExpression production.
	ExitInsertQueryExpression(c *InsertQueryExpressionContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitValuesReference is called when exiting the valuesReference production.
	ExitValuesReference(c *ValuesReferenceContext)

	// ExitInsertUpdateList is called when exiting the insertUpdateList production.
	ExitInsertUpdateList(c *InsertUpdateListContext)

	// ExitLoadStatement is called when exiting the loadStatement production.
	ExitLoadStatement(c *LoadStatementContext)

	// ExitDataOrXml is called when exiting the dataOrXml production.
	ExitDataOrXml(c *DataOrXmlContext)

	// ExitXmlRowsIdentifiedBy is called when exiting the xmlRowsIdentifiedBy production.
	ExitXmlRowsIdentifiedBy(c *XmlRowsIdentifiedByContext)

	// ExitLoadDataFileTail is called when exiting the loadDataFileTail production.
	ExitLoadDataFileTail(c *LoadDataFileTailContext)

	// ExitLoadDataFileTargetList is called when exiting the loadDataFileTargetList production.
	ExitLoadDataFileTargetList(c *LoadDataFileTargetListContext)

	// ExitFieldOrVariableList is called when exiting the fieldOrVariableList production.
	ExitFieldOrVariableList(c *FieldOrVariableListContext)

	// ExitReplaceStatement is called when exiting the replaceStatement production.
	ExitReplaceStatement(c *ReplaceStatementContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSelectStatementWithInto is called when exiting the selectStatementWithInto production.
	ExitSelectStatementWithInto(c *SelectStatementWithIntoContext)

	// ExitQueryExpression is called when exiting the queryExpression production.
	ExitQueryExpression(c *QueryExpressionContext)

	// ExitQueryExpressionBody is called when exiting the queryExpressionBody production.
	ExitQueryExpressionBody(c *QueryExpressionBodyContext)

	// ExitQueryExpressionParens is called when exiting the queryExpressionParens production.
	ExitQueryExpressionParens(c *QueryExpressionParensContext)

	// ExitQueryPrimary is called when exiting the queryPrimary production.
	ExitQueryPrimary(c *QueryPrimaryContext)

	// ExitQuerySpecification is called when exiting the querySpecification production.
	ExitQuerySpecification(c *QuerySpecificationContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitQuerySpecOption is called when exiting the querySpecOption production.
	ExitQuerySpecOption(c *QuerySpecOptionContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitSimpleLimitClause is called when exiting the simpleLimitClause production.
	ExitSimpleLimitClause(c *SimpleLimitClauseContext)

	// ExitLimitOptions is called when exiting the limitOptions production.
	ExitLimitOptions(c *LimitOptionsContext)

	// ExitLimitOption is called when exiting the limitOption production.
	ExitLimitOption(c *LimitOptionContext)

	// ExitIntoClause is called when exiting the intoClause production.
	ExitIntoClause(c *IntoClauseContext)

	// ExitProcedureAnalyseClause is called when exiting the procedureAnalyseClause production.
	ExitProcedureAnalyseClause(c *ProcedureAnalyseClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

	// ExitWindowDefinition is called when exiting the windowDefinition production.
	ExitWindowDefinition(c *WindowDefinitionContext)

	// ExitWindowSpec is called when exiting the windowSpec production.
	ExitWindowSpec(c *WindowSpecContext)

	// ExitWindowSpecDetails is called when exiting the windowSpecDetails production.
	ExitWindowSpecDetails(c *WindowSpecDetailsContext)

	// ExitWindowFrameClause is called when exiting the windowFrameClause production.
	ExitWindowFrameClause(c *WindowFrameClauseContext)

	// ExitWindowFrameUnits is called when exiting the windowFrameUnits production.
	ExitWindowFrameUnits(c *WindowFrameUnitsContext)

	// ExitWindowFrameExtent is called when exiting the windowFrameExtent production.
	ExitWindowFrameExtent(c *WindowFrameExtentContext)

	// ExitWindowFrameStart is called when exiting the windowFrameStart production.
	ExitWindowFrameStart(c *WindowFrameStartContext)

	// ExitWindowFrameBetween is called when exiting the windowFrameBetween production.
	ExitWindowFrameBetween(c *WindowFrameBetweenContext)

	// ExitWindowFrameBound is called when exiting the windowFrameBound production.
	ExitWindowFrameBound(c *WindowFrameBoundContext)

	// ExitWindowFrameExclusion is called when exiting the windowFrameExclusion production.
	ExitWindowFrameExclusion(c *WindowFrameExclusionContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitCommonTableExpression is called when exiting the commonTableExpression production.
	ExitCommonTableExpression(c *CommonTableExpressionContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitOlapOption is called when exiting the olapOption production.
	ExitOlapOption(c *OlapOptionContext)

	// ExitOrderClause is called when exiting the orderClause production.
	ExitOrderClause(c *OrderClauseContext)

	// ExitDirection is called when exiting the direction production.
	ExitDirection(c *DirectionContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitTableReferenceList is called when exiting the tableReferenceList production.
	ExitTableReferenceList(c *TableReferenceListContext)

	// ExitTableValueConstructor is called when exiting the tableValueConstructor production.
	ExitTableValueConstructor(c *TableValueConstructorContext)

	// ExitExplicitTable is called when exiting the explicitTable production.
	ExitExplicitTable(c *ExplicitTableContext)

	// ExitRowValueExplicit is called when exiting the rowValueExplicit production.
	ExitRowValueExplicit(c *RowValueExplicitContext)

	// ExitSelectOption is called when exiting the selectOption production.
	ExitSelectOption(c *SelectOptionContext)

	// ExitLockingClauseList is called when exiting the lockingClauseList production.
	ExitLockingClauseList(c *LockingClauseListContext)

	// ExitLockingClause is called when exiting the lockingClause production.
	ExitLockingClause(c *LockingClauseContext)

	// ExitLockStrengh is called when exiting the lockStrengh production.
	ExitLockStrengh(c *LockStrenghContext)

	// ExitLockedRowAction is called when exiting the lockedRowAction production.
	ExitLockedRowAction(c *LockedRowActionContext)

	// ExitSelectItemList is called when exiting the selectItemList production.
	ExitSelectItemList(c *SelectItemListContext)

	// ExitSelectItem is called when exiting the selectItem production.
	ExitSelectItem(c *SelectItemContext)

	// ExitSelectAlias is called when exiting the selectAlias production.
	ExitSelectAlias(c *SelectAliasContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitTableReference is called when exiting the tableReference production.
	ExitTableReference(c *TableReferenceContext)

	// ExitEscapedTableReference is called when exiting the escapedTableReference production.
	ExitEscapedTableReference(c *EscapedTableReferenceContext)

	// ExitJoinedTable is called when exiting the joinedTable production.
	ExitJoinedTable(c *JoinedTableContext)

	// ExitNaturalJoinType is called when exiting the naturalJoinType production.
	ExitNaturalJoinType(c *NaturalJoinTypeContext)

	// ExitInnerJoinType is called when exiting the innerJoinType production.
	ExitInnerJoinType(c *InnerJoinTypeContext)

	// ExitOuterJoinType is called when exiting the outerJoinType production.
	ExitOuterJoinType(c *OuterJoinTypeContext)

	// ExitTableFactor is called when exiting the tableFactor production.
	ExitTableFactor(c *TableFactorContext)

	// ExitSingleTable is called when exiting the singleTable production.
	ExitSingleTable(c *SingleTableContext)

	// ExitSingleTableParens is called when exiting the singleTableParens production.
	ExitSingleTableParens(c *SingleTableParensContext)

	// ExitDerivedTable is called when exiting the derivedTable production.
	ExitDerivedTable(c *DerivedTableContext)

	// ExitTableReferenceListParens is called when exiting the tableReferenceListParens production.
	ExitTableReferenceListParens(c *TableReferenceListParensContext)

	// ExitTableFunction is called when exiting the tableFunction production.
	ExitTableFunction(c *TableFunctionContext)

	// ExitColumnsClause is called when exiting the columnsClause production.
	ExitColumnsClause(c *ColumnsClauseContext)

	// ExitJtColumn is called when exiting the jtColumn production.
	ExitJtColumn(c *JtColumnContext)

	// ExitOnEmptyOrError is called when exiting the onEmptyOrError production.
	ExitOnEmptyOrError(c *OnEmptyOrErrorContext)

	// ExitOnEmpty is called when exiting the onEmpty production.
	ExitOnEmpty(c *OnEmptyContext)

	// ExitOnError is called when exiting the onError production.
	ExitOnError(c *OnErrorContext)

	// ExitJtOnResponse is called when exiting the jtOnResponse production.
	ExitJtOnResponse(c *JtOnResponseContext)

	// ExitSetOprSymbol is called when exiting the setOprSymbol production.
	ExitSetOprSymbol(c *SetOprSymbolContext)

	// ExitSetOprOption is called when exiting the setOprOption production.
	ExitSetOprOption(c *SetOprOptionContext)

	// ExitTableAlias is called when exiting the tableAlias production.
	ExitTableAlias(c *TableAliasContext)

	// ExitIndexHintList is called when exiting the indexHintList production.
	ExitIndexHintList(c *IndexHintListContext)

	// ExitIndexHint is called when exiting the indexHint production.
	ExitIndexHint(c *IndexHintContext)

	// ExitIndexHintType is called when exiting the indexHintType production.
	ExitIndexHintType(c *IndexHintTypeContext)

	// ExitKeyOrIndex is called when exiting the keyOrIndex production.
	ExitKeyOrIndex(c *KeyOrIndexContext)

	// ExitConstraintKeyType is called when exiting the constraintKeyType production.
	ExitConstraintKeyType(c *ConstraintKeyTypeContext)

	// ExitIndexHintClause is called when exiting the indexHintClause production.
	ExitIndexHintClause(c *IndexHintClauseContext)

	// ExitIndexList is called when exiting the indexList production.
	ExitIndexList(c *IndexListContext)

	// ExitIndexListElement is called when exiting the indexListElement production.
	ExitIndexListElement(c *IndexListElementContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitTransactionOrLockingStatement is called when exiting the transactionOrLockingStatement production.
	ExitTransactionOrLockingStatement(c *TransactionOrLockingStatementContext)

	// ExitTransactionStatement is called when exiting the transactionStatement production.
	ExitTransactionStatement(c *TransactionStatementContext)

	// ExitBeginWork is called when exiting the beginWork production.
	ExitBeginWork(c *BeginWorkContext)

	// ExitTransactionCharacteristic is called when exiting the transactionCharacteristic production.
	ExitTransactionCharacteristic(c *TransactionCharacteristicContext)

	// ExitSavepointStatement is called when exiting the savepointStatement production.
	ExitSavepointStatement(c *SavepointStatementContext)

	// ExitLockStatement is called when exiting the lockStatement production.
	ExitLockStatement(c *LockStatementContext)

	// ExitLockItem is called when exiting the lockItem production.
	ExitLockItem(c *LockItemContext)

	// ExitLockOption is called when exiting the lockOption production.
	ExitLockOption(c *LockOptionContext)

	// ExitXaStatement is called when exiting the xaStatement production.
	ExitXaStatement(c *XaStatementContext)

	// ExitXaConvert is called when exiting the xaConvert production.
	ExitXaConvert(c *XaConvertContext)

	// ExitXid is called when exiting the xid production.
	ExitXid(c *XidContext)

	// ExitReplicationStatement is called when exiting the replicationStatement production.
	ExitReplicationStatement(c *ReplicationStatementContext)

	// ExitResetOption is called when exiting the resetOption production.
	ExitResetOption(c *ResetOptionContext)

	// ExitMasterResetOptions is called when exiting the masterResetOptions production.
	ExitMasterResetOptions(c *MasterResetOptionsContext)

	// ExitReplicationLoad is called when exiting the replicationLoad production.
	ExitReplicationLoad(c *ReplicationLoadContext)

	// ExitChangeMaster is called when exiting the changeMaster production.
	ExitChangeMaster(c *ChangeMasterContext)

	// ExitChangeMasterOptions is called when exiting the changeMasterOptions production.
	ExitChangeMasterOptions(c *ChangeMasterOptionsContext)

	// ExitMasterOption is called when exiting the masterOption production.
	ExitMasterOption(c *MasterOptionContext)

	// ExitPrivilegeCheckDef is called when exiting the privilegeCheckDef production.
	ExitPrivilegeCheckDef(c *PrivilegeCheckDefContext)

	// ExitTablePrimaryKeyCheckDef is called when exiting the tablePrimaryKeyCheckDef production.
	ExitTablePrimaryKeyCheckDef(c *TablePrimaryKeyCheckDefContext)

	// ExitMasterTlsCiphersuitesDef is called when exiting the masterTlsCiphersuitesDef production.
	ExitMasterTlsCiphersuitesDef(c *MasterTlsCiphersuitesDefContext)

	// ExitMasterFileDef is called when exiting the masterFileDef production.
	ExitMasterFileDef(c *MasterFileDefContext)

	// ExitServerIdList is called when exiting the serverIdList production.
	ExitServerIdList(c *ServerIdListContext)

	// ExitChangeReplication is called when exiting the changeReplication production.
	ExitChangeReplication(c *ChangeReplicationContext)

	// ExitFilterDefinition is called when exiting the filterDefinition production.
	ExitFilterDefinition(c *FilterDefinitionContext)

	// ExitFilterDbList is called when exiting the filterDbList production.
	ExitFilterDbList(c *FilterDbListContext)

	// ExitFilterTableList is called when exiting the filterTableList production.
	ExitFilterTableList(c *FilterTableListContext)

	// ExitFilterStringList is called when exiting the filterStringList production.
	ExitFilterStringList(c *FilterStringListContext)

	// ExitFilterWildDbTableString is called when exiting the filterWildDbTableString production.
	ExitFilterWildDbTableString(c *FilterWildDbTableStringContext)

	// ExitFilterDbPairList is called when exiting the filterDbPairList production.
	ExitFilterDbPairList(c *FilterDbPairListContext)

	// ExitSlave is called when exiting the slave production.
	ExitSlave(c *SlaveContext)

	// ExitSlaveUntilOptions is called when exiting the slaveUntilOptions production.
	ExitSlaveUntilOptions(c *SlaveUntilOptionsContext)

	// ExitSlaveConnectionOptions is called when exiting the slaveConnectionOptions production.
	ExitSlaveConnectionOptions(c *SlaveConnectionOptionsContext)

	// ExitSlaveThreadOptions is called when exiting the slaveThreadOptions production.
	ExitSlaveThreadOptions(c *SlaveThreadOptionsContext)

	// ExitSlaveThreadOption is called when exiting the slaveThreadOption production.
	ExitSlaveThreadOption(c *SlaveThreadOptionContext)

	// ExitGroupReplication is called when exiting the groupReplication production.
	ExitGroupReplication(c *GroupReplicationContext)

	// ExitPreparedStatement is called when exiting the preparedStatement production.
	ExitPreparedStatement(c *PreparedStatementContext)

	// ExitExecuteStatement is called when exiting the executeStatement production.
	ExitExecuteStatement(c *ExecuteStatementContext)

	// ExitExecuteVarList is called when exiting the executeVarList production.
	ExitExecuteVarList(c *ExecuteVarListContext)

	// ExitCloneStatement is called when exiting the cloneStatement production.
	ExitCloneStatement(c *CloneStatementContext)

	// ExitDataDirSSL is called when exiting the dataDirSSL production.
	ExitDataDirSSL(c *DataDirSSLContext)

	// ExitSsl is called when exiting the ssl production.
	ExitSsl(c *SslContext)

	// ExitAccountManagementStatement is called when exiting the accountManagementStatement production.
	ExitAccountManagementStatement(c *AccountManagementStatementContext)

	// ExitAlterUser is called when exiting the alterUser production.
	ExitAlterUser(c *AlterUserContext)

	// ExitAlterUserTail is called when exiting the alterUserTail production.
	ExitAlterUserTail(c *AlterUserTailContext)

	// ExitUserFunction is called when exiting the userFunction production.
	ExitUserFunction(c *UserFunctionContext)

	// ExitCreateUser is called when exiting the createUser production.
	ExitCreateUser(c *CreateUserContext)

	// ExitCreateUserTail is called when exiting the createUserTail production.
	ExitCreateUserTail(c *CreateUserTailContext)

	// ExitDefaultRoleClause is called when exiting the defaultRoleClause production.
	ExitDefaultRoleClause(c *DefaultRoleClauseContext)

	// ExitRequireClause is called when exiting the requireClause production.
	ExitRequireClause(c *RequireClauseContext)

	// ExitConnectOptions is called when exiting the connectOptions production.
	ExitConnectOptions(c *ConnectOptionsContext)

	// ExitAccountLockPasswordExpireOptions is called when exiting the accountLockPasswordExpireOptions production.
	ExitAccountLockPasswordExpireOptions(c *AccountLockPasswordExpireOptionsContext)

	// ExitDropUser is called when exiting the dropUser production.
	ExitDropUser(c *DropUserContext)

	// ExitGrant is called when exiting the grant production.
	ExitGrant(c *GrantContext)

	// ExitGrantTargetList is called when exiting the grantTargetList production.
	ExitGrantTargetList(c *GrantTargetListContext)

	// ExitGrantOptions is called when exiting the grantOptions production.
	ExitGrantOptions(c *GrantOptionsContext)

	// ExitExceptRoleList is called when exiting the exceptRoleList production.
	ExitExceptRoleList(c *ExceptRoleListContext)

	// ExitWithRoles is called when exiting the withRoles production.
	ExitWithRoles(c *WithRolesContext)

	// ExitGrantAs is called when exiting the grantAs production.
	ExitGrantAs(c *GrantAsContext)

	// ExitVersionedRequireClause is called when exiting the versionedRequireClause production.
	ExitVersionedRequireClause(c *VersionedRequireClauseContext)

	// ExitRenameUser is called when exiting the renameUser production.
	ExitRenameUser(c *RenameUserContext)

	// ExitRevoke is called when exiting the revoke production.
	ExitRevoke(c *RevokeContext)

	// ExitOnTypeTo is called when exiting the onTypeTo production.
	ExitOnTypeTo(c *OnTypeToContext)

	// ExitAclType is called when exiting the aclType production.
	ExitAclType(c *AclTypeContext)

	// ExitRoleOrPrivilegesList is called when exiting the roleOrPrivilegesList production.
	ExitRoleOrPrivilegesList(c *RoleOrPrivilegesListContext)

	// ExitRoleOrPrivilege is called when exiting the roleOrPrivilege production.
	ExitRoleOrPrivilege(c *RoleOrPrivilegeContext)

	// ExitGrantIdentifier is called when exiting the grantIdentifier production.
	ExitGrantIdentifier(c *GrantIdentifierContext)

	// ExitRequireList is called when exiting the requireList production.
	ExitRequireList(c *RequireListContext)

	// ExitRequireListElement is called when exiting the requireListElement production.
	ExitRequireListElement(c *RequireListElementContext)

	// ExitGrantOption is called when exiting the grantOption production.
	ExitGrantOption(c *GrantOptionContext)

	// ExitSetRole is called when exiting the setRole production.
	ExitSetRole(c *SetRoleContext)

	// ExitRoleList is called when exiting the roleList production.
	ExitRoleList(c *RoleListContext)

	// ExitRole is called when exiting the role production.
	ExitRole(c *RoleContext)

	// ExitTableAdministrationStatement is called when exiting the tableAdministrationStatement production.
	ExitTableAdministrationStatement(c *TableAdministrationStatementContext)

	// ExitHistogram is called when exiting the histogram production.
	ExitHistogram(c *HistogramContext)

	// ExitCheckOption is called when exiting the checkOption production.
	ExitCheckOption(c *CheckOptionContext)

	// ExitRepairType is called when exiting the repairType production.
	ExitRepairType(c *RepairTypeContext)

	// ExitInstallUninstallStatment is called when exiting the installUninstallStatment production.
	ExitInstallUninstallStatment(c *InstallUninstallStatmentContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitStartOptionValueList is called when exiting the startOptionValueList production.
	ExitStartOptionValueList(c *StartOptionValueListContext)

	// ExitTransactionCharacteristics is called when exiting the transactionCharacteristics production.
	ExitTransactionCharacteristics(c *TransactionCharacteristicsContext)

	// ExitTransactionAccessMode is called when exiting the transactionAccessMode production.
	ExitTransactionAccessMode(c *TransactionAccessModeContext)

	// ExitIsolationLevel is called when exiting the isolationLevel production.
	ExitIsolationLevel(c *IsolationLevelContext)

	// ExitOptionValueListContinued is called when exiting the optionValueListContinued production.
	ExitOptionValueListContinued(c *OptionValueListContinuedContext)

	// ExitOptionValueNoOptionType is called when exiting the optionValueNoOptionType production.
	ExitOptionValueNoOptionType(c *OptionValueNoOptionTypeContext)

	// ExitOptionValue is called when exiting the optionValue production.
	ExitOptionValue(c *OptionValueContext)

	// ExitSetSystemVariable is called when exiting the setSystemVariable production.
	ExitSetSystemVariable(c *SetSystemVariableContext)

	// ExitStartOptionValueListFollowingOptionType is called when exiting the startOptionValueListFollowingOptionType production.
	ExitStartOptionValueListFollowingOptionType(c *StartOptionValueListFollowingOptionTypeContext)

	// ExitOptionValueFollowingOptionType is called when exiting the optionValueFollowingOptionType production.
	ExitOptionValueFollowingOptionType(c *OptionValueFollowingOptionTypeContext)

	// ExitSetExprOrDefault is called when exiting the setExprOrDefault production.
	ExitSetExprOrDefault(c *SetExprOrDefaultContext)

	// ExitShowStatement is called when exiting the showStatement production.
	ExitShowStatement(c *ShowStatementContext)

	// ExitShowCommandType is called when exiting the showCommandType production.
	ExitShowCommandType(c *ShowCommandTypeContext)

	// ExitNonBlocking is called when exiting the nonBlocking production.
	ExitNonBlocking(c *NonBlockingContext)

	// ExitFromOrIn is called when exiting the fromOrIn production.
	ExitFromOrIn(c *FromOrInContext)

	// ExitInDb is called when exiting the inDb production.
	ExitInDb(c *InDbContext)

	// ExitProfileType is called when exiting the profileType production.
	ExitProfileType(c *ProfileTypeContext)

	// ExitOtherAdministrativeStatement is called when exiting the otherAdministrativeStatement production.
	ExitOtherAdministrativeStatement(c *OtherAdministrativeStatementContext)

	// ExitKeyCacheListOrParts is called when exiting the keyCacheListOrParts production.
	ExitKeyCacheListOrParts(c *KeyCacheListOrPartsContext)

	// ExitKeyCacheList is called when exiting the keyCacheList production.
	ExitKeyCacheList(c *KeyCacheListContext)

	// ExitAssignToKeycache is called when exiting the assignToKeycache production.
	ExitAssignToKeycache(c *AssignToKeycacheContext)

	// ExitAssignToKeycachePartition is called when exiting the assignToKeycachePartition production.
	ExitAssignToKeycachePartition(c *AssignToKeycachePartitionContext)

	// ExitCacheKeyList is called when exiting the cacheKeyList production.
	ExitCacheKeyList(c *CacheKeyListContext)

	// ExitKeyUsageElement is called when exiting the keyUsageElement production.
	ExitKeyUsageElement(c *KeyUsageElementContext)

	// ExitKeyUsageList is called when exiting the keyUsageList production.
	ExitKeyUsageList(c *KeyUsageListContext)

	// ExitFlushOption is called when exiting the flushOption production.
	ExitFlushOption(c *FlushOptionContext)

	// ExitLogType is called when exiting the logType production.
	ExitLogType(c *LogTypeContext)

	// ExitFlushTables is called when exiting the flushTables production.
	ExitFlushTables(c *FlushTablesContext)

	// ExitFlushTablesOptions is called when exiting the flushTablesOptions production.
	ExitFlushTablesOptions(c *FlushTablesOptionsContext)

	// ExitPreloadTail is called when exiting the preloadTail production.
	ExitPreloadTail(c *PreloadTailContext)

	// ExitPreloadList is called when exiting the preloadList production.
	ExitPreloadList(c *PreloadListContext)

	// ExitPreloadKeys is called when exiting the preloadKeys production.
	ExitPreloadKeys(c *PreloadKeysContext)

	// ExitAdminPartition is called when exiting the adminPartition production.
	ExitAdminPartition(c *AdminPartitionContext)

	// ExitResourceGroupManagement is called when exiting the resourceGroupManagement production.
	ExitResourceGroupManagement(c *ResourceGroupManagementContext)

	// ExitCreateResourceGroup is called when exiting the createResourceGroup production.
	ExitCreateResourceGroup(c *CreateResourceGroupContext)

	// ExitResourceGroupVcpuList is called when exiting the resourceGroupVcpuList production.
	ExitResourceGroupVcpuList(c *ResourceGroupVcpuListContext)

	// ExitVcpuNumOrRange is called when exiting the vcpuNumOrRange production.
	ExitVcpuNumOrRange(c *VcpuNumOrRangeContext)

	// ExitResourceGroupPriority is called when exiting the resourceGroupPriority production.
	ExitResourceGroupPriority(c *ResourceGroupPriorityContext)

	// ExitResourceGroupEnableDisable is called when exiting the resourceGroupEnableDisable production.
	ExitResourceGroupEnableDisable(c *ResourceGroupEnableDisableContext)

	// ExitAlterResourceGroup is called when exiting the alterResourceGroup production.
	ExitAlterResourceGroup(c *AlterResourceGroupContext)

	// ExitSetResourceGroup is called when exiting the setResourceGroup production.
	ExitSetResourceGroup(c *SetResourceGroupContext)

	// ExitThreadIdList is called when exiting the threadIdList production.
	ExitThreadIdList(c *ThreadIdListContext)

	// ExitDropResourceGroup is called when exiting the dropResourceGroup production.
	ExitDropResourceGroup(c *DropResourceGroupContext)

	// ExitUtilityStatement is called when exiting the utilityStatement production.
	ExitUtilityStatement(c *UtilityStatementContext)

	// ExitDescribeStatement is called when exiting the describeStatement production.
	ExitDescribeStatement(c *DescribeStatementContext)

	// ExitExplainStatement is called when exiting the explainStatement production.
	ExitExplainStatement(c *ExplainStatementContext)

	// ExitExplainableStatement is called when exiting the explainableStatement production.
	ExitExplainableStatement(c *ExplainableStatementContext)

	// ExitHelpCommand is called when exiting the helpCommand production.
	ExitHelpCommand(c *HelpCommandContext)

	// ExitUseCommand is called when exiting the useCommand production.
	ExitUseCommand(c *UseCommandContext)

	// ExitRestartServer is called when exiting the restartServer production.
	ExitRestartServer(c *RestartServerContext)

	// ExitExprOr is called when exiting the exprOr production.
	ExitExprOr(c *ExprOrContext)

	// ExitExprNot is called when exiting the exprNot production.
	ExitExprNot(c *ExprNotContext)

	// ExitExprIs is called when exiting the exprIs production.
	ExitExprIs(c *ExprIsContext)

	// ExitExprAnd is called when exiting the exprAnd production.
	ExitExprAnd(c *ExprAndContext)

	// ExitExprXor is called when exiting the exprXor production.
	ExitExprXor(c *ExprXorContext)

	// ExitPrimaryExprPredicate is called when exiting the primaryExprPredicate production.
	ExitPrimaryExprPredicate(c *PrimaryExprPredicateContext)

	// ExitPrimaryExprCompare is called when exiting the primaryExprCompare production.
	ExitPrimaryExprCompare(c *PrimaryExprCompareContext)

	// ExitPrimaryExprAllAny is called when exiting the primaryExprAllAny production.
	ExitPrimaryExprAllAny(c *PrimaryExprAllAnyContext)

	// ExitPrimaryExprIsNull is called when exiting the primaryExprIsNull production.
	ExitPrimaryExprIsNull(c *PrimaryExprIsNullContext)

	// ExitCompOp is called when exiting the compOp production.
	ExitCompOp(c *CompOpContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitPredicateExprIn is called when exiting the predicateExprIn production.
	ExitPredicateExprIn(c *PredicateExprInContext)

	// ExitPredicateExprBetween is called when exiting the predicateExprBetween production.
	ExitPredicateExprBetween(c *PredicateExprBetweenContext)

	// ExitPredicateExprLike is called when exiting the predicateExprLike production.
	ExitPredicateExprLike(c *PredicateExprLikeContext)

	// ExitPredicateExprRegex is called when exiting the predicateExprRegex production.
	ExitPredicateExprRegex(c *PredicateExprRegexContext)

	// ExitBitExpr is called when exiting the bitExpr production.
	ExitBitExpr(c *BitExprContext)

	// ExitSimpleExprConvert is called when exiting the simpleExprConvert production.
	ExitSimpleExprConvert(c *SimpleExprConvertContext)

	// ExitSimpleExprSearchJson is called when exiting the simpleExprSearchJson production.
	ExitSimpleExprSearchJson(c *SimpleExprSearchJsonContext)

	// ExitSimpleExprVariable is called when exiting the simpleExprVariable production.
	ExitSimpleExprVariable(c *SimpleExprVariableContext)

	// ExitSimpleExprCast is called when exiting the simpleExprCast production.
	ExitSimpleExprCast(c *SimpleExprCastContext)

	// ExitSimpleExprUnary is called when exiting the simpleExprUnary production.
	ExitSimpleExprUnary(c *SimpleExprUnaryContext)

	// ExitSimpleExprOdbc is called when exiting the simpleExprOdbc production.
	ExitSimpleExprOdbc(c *SimpleExprOdbcContext)

	// ExitSimpleExprRuntimeFunction is called when exiting the simpleExprRuntimeFunction production.
	ExitSimpleExprRuntimeFunction(c *SimpleExprRuntimeFunctionContext)

	// ExitSimpleExprFunction is called when exiting the simpleExprFunction production.
	ExitSimpleExprFunction(c *SimpleExprFunctionContext)

	// ExitSimpleExprCollate is called when exiting the simpleExprCollate production.
	ExitSimpleExprCollate(c *SimpleExprCollateContext)

	// ExitSimpleExprMatch is called when exiting the simpleExprMatch production.
	ExitSimpleExprMatch(c *SimpleExprMatchContext)

	// ExitSimpleExprWindowingFunction is called when exiting the simpleExprWindowingFunction production.
	ExitSimpleExprWindowingFunction(c *SimpleExprWindowingFunctionContext)

	// ExitSimpleExprBinary is called when exiting the simpleExprBinary production.
	ExitSimpleExprBinary(c *SimpleExprBinaryContext)

	// ExitSimpleExprColumnRef is called when exiting the simpleExprColumnRef production.
	ExitSimpleExprColumnRef(c *SimpleExprColumnRefContext)

	// ExitSimpleExprParamMarker is called when exiting the simpleExprParamMarker production.
	ExitSimpleExprParamMarker(c *SimpleExprParamMarkerContext)

	// ExitSimpleExprSum is called when exiting the simpleExprSum production.
	ExitSimpleExprSum(c *SimpleExprSumContext)

	// ExitSimpleExprConvertUsing is called when exiting the simpleExprConvertUsing production.
	ExitSimpleExprConvertUsing(c *SimpleExprConvertUsingContext)

	// ExitSimpleExprSubQuery is called when exiting the simpleExprSubQuery production.
	ExitSimpleExprSubQuery(c *SimpleExprSubQueryContext)

	// ExitSimpleExprGroupingOperation is called when exiting the simpleExprGroupingOperation production.
	ExitSimpleExprGroupingOperation(c *SimpleExprGroupingOperationContext)

	// ExitSimpleExprNot is called when exiting the simpleExprNot production.
	ExitSimpleExprNot(c *SimpleExprNotContext)

	// ExitSimpleExprValues is called when exiting the simpleExprValues production.
	ExitSimpleExprValues(c *SimpleExprValuesContext)

	// ExitSimpleExprDefault is called when exiting the simpleExprDefault production.
	ExitSimpleExprDefault(c *SimpleExprDefaultContext)

	// ExitSimpleExprList is called when exiting the simpleExprList production.
	ExitSimpleExprList(c *SimpleExprListContext)

	// ExitSimpleExprInterval is called when exiting the simpleExprInterval production.
	ExitSimpleExprInterval(c *SimpleExprIntervalContext)

	// ExitSimpleExprCase is called when exiting the simpleExprCase production.
	ExitSimpleExprCase(c *SimpleExprCaseContext)

	// ExitSimpleExprConcat is called when exiting the simpleExprConcat production.
	ExitSimpleExprConcat(c *SimpleExprConcatContext)

	// ExitSimpleExprLiteral is called when exiting the simpleExprLiteral production.
	ExitSimpleExprLiteral(c *SimpleExprLiteralContext)

	// ExitArrayCast is called when exiting the arrayCast production.
	ExitArrayCast(c *ArrayCastContext)

	// ExitJsonOperator is called when exiting the jsonOperator production.
	ExitJsonOperator(c *JsonOperatorContext)

	// ExitSumExpr is called when exiting the sumExpr production.
	ExitSumExpr(c *SumExprContext)

	// ExitGroupingOperation is called when exiting the groupingOperation production.
	ExitGroupingOperation(c *GroupingOperationContext)

	// ExitWindowFunctionCall is called when exiting the windowFunctionCall production.
	ExitWindowFunctionCall(c *WindowFunctionCallContext)

	// ExitWindowingClause is called when exiting the windowingClause production.
	ExitWindowingClause(c *WindowingClauseContext)

	// ExitLeadLagInfo is called when exiting the leadLagInfo production.
	ExitLeadLagInfo(c *LeadLagInfoContext)

	// ExitNullTreatment is called when exiting the nullTreatment production.
	ExitNullTreatment(c *NullTreatmentContext)

	// ExitJsonFunction is called when exiting the jsonFunction production.
	ExitJsonFunction(c *JsonFunctionContext)

	// ExitInSumExpr is called when exiting the inSumExpr production.
	ExitInSumExpr(c *InSumExprContext)

	// ExitIdentListArg is called when exiting the identListArg production.
	ExitIdentListArg(c *IdentListArgContext)

	// ExitIdentList is called when exiting the identList production.
	ExitIdentList(c *IdentListContext)

	// ExitFulltextOptions is called when exiting the fulltextOptions production.
	ExitFulltextOptions(c *FulltextOptionsContext)

	// ExitRuntimeFunctionCall is called when exiting the runtimeFunctionCall production.
	ExitRuntimeFunctionCall(c *RuntimeFunctionCallContext)

	// ExitGeometryFunction is called when exiting the geometryFunction production.
	ExitGeometryFunction(c *GeometryFunctionContext)

	// ExitTimeFunctionParameters is called when exiting the timeFunctionParameters production.
	ExitTimeFunctionParameters(c *TimeFunctionParametersContext)

	// ExitFractionalPrecision is called when exiting the fractionalPrecision production.
	ExitFractionalPrecision(c *FractionalPrecisionContext)

	// ExitWeightStringLevels is called when exiting the weightStringLevels production.
	ExitWeightStringLevels(c *WeightStringLevelsContext)

	// ExitWeightStringLevelListItem is called when exiting the weightStringLevelListItem production.
	ExitWeightStringLevelListItem(c *WeightStringLevelListItemContext)

	// ExitDateTimeTtype is called when exiting the dateTimeTtype production.
	ExitDateTimeTtype(c *DateTimeTtypeContext)

	// ExitTrimFunction is called when exiting the trimFunction production.
	ExitTrimFunction(c *TrimFunctionContext)

	// ExitSubstringFunction is called when exiting the substringFunction production.
	ExitSubstringFunction(c *SubstringFunctionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitSearchJsonFunction is called when exiting the searchJsonFunction production.
	ExitSearchJsonFunction(c *SearchJsonFunctionContext)

	// ExitJsonValueReturning is called when exiting the jsonValueReturning production.
	ExitJsonValueReturning(c *JsonValueReturningContext)

	// ExitJsonValueOnEmpty is called when exiting the jsonValueOnEmpty production.
	ExitJsonValueOnEmpty(c *JsonValueOnEmptyContext)

	// ExitJsonValueOnError is called when exiting the jsonValueOnError production.
	ExitJsonValueOnError(c *JsonValueOnErrorContext)

	// ExitUdfExprList is called when exiting the udfExprList production.
	ExitUdfExprList(c *UdfExprListContext)

	// ExitUdfExpr is called when exiting the udfExpr production.
	ExitUdfExpr(c *UdfExprContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitUserVariable is called when exiting the userVariable production.
	ExitUserVariable(c *UserVariableContext)

	// ExitSystemVariable is called when exiting the systemVariable production.
	ExitSystemVariable(c *SystemVariableContext)

	// ExitInternalVariableName is called when exiting the internalVariableName production.
	ExitInternalVariableName(c *InternalVariableNameContext)

	// ExitWhenExpression is called when exiting the whenExpression production.
	ExitWhenExpression(c *WhenExpressionContext)

	// ExitThenExpression is called when exiting the thenExpression production.
	ExitThenExpression(c *ThenExpressionContext)

	// ExitElseExpression is called when exiting the elseExpression production.
	ExitElseExpression(c *ElseExpressionContext)

	// ExitCastType is called when exiting the castType production.
	ExitCastType(c *CastTypeContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitCharset is called when exiting the charset production.
	ExitCharset(c *CharsetContext)

	// ExitNotRule is called when exiting the notRule production.
	ExitNotRule(c *NotRuleContext)

	// ExitNot2Rule is called when exiting the not2Rule production.
	ExitNot2Rule(c *Not2RuleContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitIntervalTimeStamp is called when exiting the intervalTimeStamp production.
	ExitIntervalTimeStamp(c *IntervalTimeStampContext)

	// ExitExprListWithParentheses is called when exiting the exprListWithParentheses production.
	ExitExprListWithParentheses(c *ExprListWithParenthesesContext)

	// ExitExprWithParentheses is called when exiting the exprWithParentheses production.
	ExitExprWithParentheses(c *ExprWithParenthesesContext)

	// ExitSimpleExprWithParentheses is called when exiting the simpleExprWithParentheses production.
	ExitSimpleExprWithParentheses(c *SimpleExprWithParenthesesContext)

	// ExitOrderList is called when exiting the orderList production.
	ExitOrderList(c *OrderListContext)

	// ExitOrderExpression is called when exiting the orderExpression production.
	ExitOrderExpression(c *OrderExpressionContext)

	// ExitGroupList is called when exiting the groupList production.
	ExitGroupList(c *GroupListContext)

	// ExitGroupingExpression is called when exiting the groupingExpression production.
	ExitGroupingExpression(c *GroupingExpressionContext)

	// ExitChannel is called when exiting the channel production.
	ExitChannel(c *ChannelContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfBody is called when exiting the ifBody production.
	ExitIfBody(c *IfBodyContext)

	// ExitThenStatement is called when exiting the thenStatement production.
	ExitThenStatement(c *ThenStatementContext)

	// ExitCompoundStatementList is called when exiting the compoundStatementList production.
	ExitCompoundStatementList(c *CompoundStatementListContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitLabeledBlock is called when exiting the labeledBlock production.
	ExitLabeledBlock(c *LabeledBlockContext)

	// ExitUnlabeledBlock is called when exiting the unlabeledBlock production.
	ExitUnlabeledBlock(c *UnlabeledBlockContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitBeginEndBlock is called when exiting the beginEndBlock production.
	ExitBeginEndBlock(c *BeginEndBlockContext)

	// ExitLabeledControl is called when exiting the labeledControl production.
	ExitLabeledControl(c *LabeledControlContext)

	// ExitUnlabeledControl is called when exiting the unlabeledControl production.
	ExitUnlabeledControl(c *UnlabeledControlContext)

	// ExitLoopBlock is called when exiting the loopBlock production.
	ExitLoopBlock(c *LoopBlockContext)

	// ExitWhileDoBlock is called when exiting the whileDoBlock production.
	ExitWhileDoBlock(c *WhileDoBlockContext)

	// ExitRepeatUntilBlock is called when exiting the repeatUntilBlock production.
	ExitRepeatUntilBlock(c *RepeatUntilBlockContext)

	// ExitSpDeclarations is called when exiting the spDeclarations production.
	ExitSpDeclarations(c *SpDeclarationsContext)

	// ExitSpDeclaration is called when exiting the spDeclaration production.
	ExitSpDeclaration(c *SpDeclarationContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitConditionDeclaration is called when exiting the conditionDeclaration production.
	ExitConditionDeclaration(c *ConditionDeclarationContext)

	// ExitSpCondition is called when exiting the spCondition production.
	ExitSpCondition(c *SpConditionContext)

	// ExitSqlstate is called when exiting the sqlstate production.
	ExitSqlstate(c *SqlstateContext)

	// ExitHandlerDeclaration is called when exiting the handlerDeclaration production.
	ExitHandlerDeclaration(c *HandlerDeclarationContext)

	// ExitHandlerCondition is called when exiting the handlerCondition production.
	ExitHandlerCondition(c *HandlerConditionContext)

	// ExitCursorDeclaration is called when exiting the cursorDeclaration production.
	ExitCursorDeclaration(c *CursorDeclarationContext)

	// ExitIterateStatement is called when exiting the iterateStatement production.
	ExitIterateStatement(c *IterateStatementContext)

	// ExitLeaveStatement is called when exiting the leaveStatement production.
	ExitLeaveStatement(c *LeaveStatementContext)

	// ExitGetDiagnostics is called when exiting the getDiagnostics production.
	ExitGetDiagnostics(c *GetDiagnosticsContext)

	// ExitSignalAllowedExpr is called when exiting the signalAllowedExpr production.
	ExitSignalAllowedExpr(c *SignalAllowedExprContext)

	// ExitStatementInformationItem is called when exiting the statementInformationItem production.
	ExitStatementInformationItem(c *StatementInformationItemContext)

	// ExitConditionInformationItem is called when exiting the conditionInformationItem production.
	ExitConditionInformationItem(c *ConditionInformationItemContext)

	// ExitSignalInformationItemName is called when exiting the signalInformationItemName production.
	ExitSignalInformationItemName(c *SignalInformationItemNameContext)

	// ExitSignalStatement is called when exiting the signalStatement production.
	ExitSignalStatement(c *SignalStatementContext)

	// ExitResignalStatement is called when exiting the resignalStatement production.
	ExitResignalStatement(c *ResignalStatementContext)

	// ExitSignalInformationItem is called when exiting the signalInformationItem production.
	ExitSignalInformationItem(c *SignalInformationItemContext)

	// ExitCursorOpen is called when exiting the cursorOpen production.
	ExitCursorOpen(c *CursorOpenContext)

	// ExitCursorClose is called when exiting the cursorClose production.
	ExitCursorClose(c *CursorCloseContext)

	// ExitCursorFetch is called when exiting the cursorFetch production.
	ExitCursorFetch(c *CursorFetchContext)

	// ExitSchedule is called when exiting the schedule production.
	ExitSchedule(c *ScheduleContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitCheckOrReferences is called when exiting the checkOrReferences production.
	ExitCheckOrReferences(c *CheckOrReferencesContext)

	// ExitCheckConstraint is called when exiting the checkConstraint production.
	ExitCheckConstraint(c *CheckConstraintContext)

	// ExitConstraintEnforcement is called when exiting the constraintEnforcement production.
	ExitConstraintEnforcement(c *ConstraintEnforcementContext)

	// ExitTableConstraintDef is called when exiting the tableConstraintDef production.
	ExitTableConstraintDef(c *TableConstraintDefContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitFieldDefinition is called when exiting the fieldDefinition production.
	ExitFieldDefinition(c *FieldDefinitionContext)

	// ExitColumnAttribute is called when exiting the columnAttribute production.
	ExitColumnAttribute(c *ColumnAttributeContext)

	// ExitColumnFormat is called when exiting the columnFormat production.
	ExitColumnFormat(c *ColumnFormatContext)

	// ExitStorageMedia is called when exiting the storageMedia production.
	ExitStorageMedia(c *StorageMediaContext)

	// ExitGcolAttribute is called when exiting the gcolAttribute production.
	ExitGcolAttribute(c *GcolAttributeContext)

	// ExitReferences is called when exiting the references production.
	ExitReferences(c *ReferencesContext)

	// ExitDeleteOption is called when exiting the deleteOption production.
	ExitDeleteOption(c *DeleteOptionContext)

	// ExitKeyList is called when exiting the keyList production.
	ExitKeyList(c *KeyListContext)

	// ExitKeyPart is called when exiting the keyPart production.
	ExitKeyPart(c *KeyPartContext)

	// ExitKeyListWithExpression is called when exiting the keyListWithExpression production.
	ExitKeyListWithExpression(c *KeyListWithExpressionContext)

	// ExitKeyPartOrExpression is called when exiting the keyPartOrExpression production.
	ExitKeyPartOrExpression(c *KeyPartOrExpressionContext)

	// ExitKeyListVariants is called when exiting the keyListVariants production.
	ExitKeyListVariants(c *KeyListVariantsContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitIndexOption is called when exiting the indexOption production.
	ExitIndexOption(c *IndexOptionContext)

	// ExitCommonIndexOption is called when exiting the commonIndexOption production.
	ExitCommonIndexOption(c *CommonIndexOptionContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)

	// ExitIndexTypeClause is called when exiting the indexTypeClause production.
	ExitIndexTypeClause(c *IndexTypeClauseContext)

	// ExitFulltextIndexOption is called when exiting the fulltextIndexOption production.
	ExitFulltextIndexOption(c *FulltextIndexOptionContext)

	// ExitSpatialIndexOption is called when exiting the spatialIndexOption production.
	ExitSpatialIndexOption(c *SpatialIndexOptionContext)

	// ExitDataTypeDefinition is called when exiting the dataTypeDefinition production.
	ExitDataTypeDefinition(c *DataTypeDefinitionContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitNchar is called when exiting the nchar production.
	ExitNchar(c *NcharContext)

	// ExitRealType is called when exiting the realType production.
	ExitRealType(c *RealTypeContext)

	// ExitFieldLength is called when exiting the fieldLength production.
	ExitFieldLength(c *FieldLengthContext)

	// ExitFieldOptions is called when exiting the fieldOptions production.
	ExitFieldOptions(c *FieldOptionsContext)

	// ExitCharsetWithOptBinary is called when exiting the charsetWithOptBinary production.
	ExitCharsetWithOptBinary(c *CharsetWithOptBinaryContext)

	// ExitAscii is called when exiting the ascii production.
	ExitAscii(c *AsciiContext)

	// ExitUnicode is called when exiting the unicode production.
	ExitUnicode(c *UnicodeContext)

	// ExitWsNumCodepoints is called when exiting the wsNumCodepoints production.
	ExitWsNumCodepoints(c *WsNumCodepointsContext)

	// ExitTypeDatetimePrecision is called when exiting the typeDatetimePrecision production.
	ExitTypeDatetimePrecision(c *TypeDatetimePrecisionContext)

	// ExitCharsetName is called when exiting the charsetName production.
	ExitCharsetName(c *CharsetNameContext)

	// ExitCollationName is called when exiting the collationName production.
	ExitCollationName(c *CollationNameContext)

	// ExitCreateTableOptions is called when exiting the createTableOptions production.
	ExitCreateTableOptions(c *CreateTableOptionsContext)

	// ExitCreateTableOptionsSpaceSeparated is called when exiting the createTableOptionsSpaceSeparated production.
	ExitCreateTableOptionsSpaceSeparated(c *CreateTableOptionsSpaceSeparatedContext)

	// ExitCreateTableOption is called when exiting the createTableOption production.
	ExitCreateTableOption(c *CreateTableOptionContext)

	// ExitTernaryOption is called when exiting the ternaryOption production.
	ExitTernaryOption(c *TernaryOptionContext)

	// ExitDefaultCollation is called when exiting the defaultCollation production.
	ExitDefaultCollation(c *DefaultCollationContext)

	// ExitDefaultEncryption is called when exiting the defaultEncryption production.
	ExitDefaultEncryption(c *DefaultEncryptionContext)

	// ExitDefaultCharset is called when exiting the defaultCharset production.
	ExitDefaultCharset(c *DefaultCharsetContext)

	// ExitPartitionClause is called when exiting the partitionClause production.
	ExitPartitionClause(c *PartitionClauseContext)

	// ExitPartitionDefKey is called when exiting the partitionDefKey production.
	ExitPartitionDefKey(c *PartitionDefKeyContext)

	// ExitPartitionDefHash is called when exiting the partitionDefHash production.
	ExitPartitionDefHash(c *PartitionDefHashContext)

	// ExitPartitionDefRangeList is called when exiting the partitionDefRangeList production.
	ExitPartitionDefRangeList(c *PartitionDefRangeListContext)

	// ExitSubPartitions is called when exiting the subPartitions production.
	ExitSubPartitions(c *SubPartitionsContext)

	// ExitPartitionKeyAlgorithm is called when exiting the partitionKeyAlgorithm production.
	ExitPartitionKeyAlgorithm(c *PartitionKeyAlgorithmContext)

	// ExitPartitionDefinitions is called when exiting the partitionDefinitions production.
	ExitPartitionDefinitions(c *PartitionDefinitionsContext)

	// ExitPartitionDefinition is called when exiting the partitionDefinition production.
	ExitPartitionDefinition(c *PartitionDefinitionContext)

	// ExitPartitionValuesIn is called when exiting the partitionValuesIn production.
	ExitPartitionValuesIn(c *PartitionValuesInContext)

	// ExitPartitionOption is called when exiting the partitionOption production.
	ExitPartitionOption(c *PartitionOptionContext)

	// ExitSubpartitionDefinition is called when exiting the subpartitionDefinition production.
	ExitSubpartitionDefinition(c *SubpartitionDefinitionContext)

	// ExitPartitionValueItemListParen is called when exiting the partitionValueItemListParen production.
	ExitPartitionValueItemListParen(c *PartitionValueItemListParenContext)

	// ExitPartitionValueItem is called when exiting the partitionValueItem production.
	ExitPartitionValueItem(c *PartitionValueItemContext)

	// ExitDefinerClause is called when exiting the definerClause production.
	ExitDefinerClause(c *DefinerClauseContext)

	// ExitIfExists is called when exiting the ifExists production.
	ExitIfExists(c *IfExistsContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)

	// ExitProcedureParameter is called when exiting the procedureParameter production.
	ExitProcedureParameter(c *ProcedureParameterContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitCollate is called when exiting the collate production.
	ExitCollate(c *CollateContext)

	// ExitTypeWithOptCollate is called when exiting the typeWithOptCollate production.
	ExitTypeWithOptCollate(c *TypeWithOptCollateContext)

	// ExitSchemaIdentifierPair is called when exiting the schemaIdentifierPair production.
	ExitSchemaIdentifierPair(c *SchemaIdentifierPairContext)

	// ExitViewRefList is called when exiting the viewRefList production.
	ExitViewRefList(c *ViewRefListContext)

	// ExitUpdateList is called when exiting the updateList production.
	ExitUpdateList(c *UpdateListContext)

	// ExitUpdateElement is called when exiting the updateElement production.
	ExitUpdateElement(c *UpdateElementContext)

	// ExitCharsetClause is called when exiting the charsetClause production.
	ExitCharsetClause(c *CharsetClauseContext)

	// ExitFieldsClause is called when exiting the fieldsClause production.
	ExitFieldsClause(c *FieldsClauseContext)

	// ExitFieldTerm is called when exiting the fieldTerm production.
	ExitFieldTerm(c *FieldTermContext)

	// ExitLinesClause is called when exiting the linesClause production.
	ExitLinesClause(c *LinesClauseContext)

	// ExitLineTerm is called when exiting the lineTerm production.
	ExitLineTerm(c *LineTermContext)

	// ExitUserList is called when exiting the userList production.
	ExitUserList(c *UserListContext)

	// ExitCreateUserList is called when exiting the createUserList production.
	ExitCreateUserList(c *CreateUserListContext)

	// ExitAlterUserList is called when exiting the alterUserList production.
	ExitAlterUserList(c *AlterUserListContext)

	// ExitCreateUserEntry is called when exiting the createUserEntry production.
	ExitCreateUserEntry(c *CreateUserEntryContext)

	// ExitAlterUserEntry is called when exiting the alterUserEntry production.
	ExitAlterUserEntry(c *AlterUserEntryContext)

	// ExitRetainCurrentPassword is called when exiting the retainCurrentPassword production.
	ExitRetainCurrentPassword(c *RetainCurrentPasswordContext)

	// ExitDiscardOldPassword is called when exiting the discardOldPassword production.
	ExitDiscardOldPassword(c *DiscardOldPasswordContext)

	// ExitReplacePassword is called when exiting the replacePassword production.
	ExitReplacePassword(c *ReplacePasswordContext)

	// ExitUserIdentifierOrText is called when exiting the userIdentifierOrText production.
	ExitUserIdentifierOrText(c *UserIdentifierOrTextContext)

	// ExitUser is called when exiting the user production.
	ExitUser(c *UserContext)

	// ExitLikeClause is called when exiting the likeClause production.
	ExitLikeClause(c *LikeClauseContext)

	// ExitLikeOrWhere is called when exiting the likeOrWhere production.
	ExitLikeOrWhere(c *LikeOrWhereContext)

	// ExitOnlineOption is called when exiting the onlineOption production.
	ExitOnlineOption(c *OnlineOptionContext)

	// ExitNoWriteToBinLog is called when exiting the noWriteToBinLog production.
	ExitNoWriteToBinLog(c *NoWriteToBinLogContext)

	// ExitUsePartition is called when exiting the usePartition production.
	ExitUsePartition(c *UsePartitionContext)

	// ExitFieldIdentifier is called when exiting the fieldIdentifier production.
	ExitFieldIdentifier(c *FieldIdentifierContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitColumnInternalRef is called when exiting the columnInternalRef production.
	ExitColumnInternalRef(c *ColumnInternalRefContext)

	// ExitColumnInternalRefList is called when exiting the columnInternalRefList production.
	ExitColumnInternalRefList(c *ColumnInternalRefListContext)

	// ExitColumnRef is called when exiting the columnRef production.
	ExitColumnRef(c *ColumnRefContext)

	// ExitInsertIdentifier is called when exiting the insertIdentifier production.
	ExitInsertIdentifier(c *InsertIdentifierContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitIndexRef is called when exiting the indexRef production.
	ExitIndexRef(c *IndexRefContext)

	// ExitTableWild is called when exiting the tableWild production.
	ExitTableWild(c *TableWildContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitSchemaRef is called when exiting the schemaRef production.
	ExitSchemaRef(c *SchemaRefContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitProcedureRef is called when exiting the procedureRef production.
	ExitProcedureRef(c *ProcedureRefContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitFunctionRef is called when exiting the functionRef production.
	ExitFunctionRef(c *FunctionRefContext)

	// ExitTriggerName is called when exiting the triggerName production.
	ExitTriggerName(c *TriggerNameContext)

	// ExitTriggerRef is called when exiting the triggerRef production.
	ExitTriggerRef(c *TriggerRefContext)

	// ExitViewName is called when exiting the viewName production.
	ExitViewName(c *ViewNameContext)

	// ExitViewRef is called when exiting the viewRef production.
	ExitViewRef(c *ViewRefContext)

	// ExitTablespaceName is called when exiting the tablespaceName production.
	ExitTablespaceName(c *TablespaceNameContext)

	// ExitTablespaceRef is called when exiting the tablespaceRef production.
	ExitTablespaceRef(c *TablespaceRefContext)

	// ExitLogfileGroupName is called when exiting the logfileGroupName production.
	ExitLogfileGroupName(c *LogfileGroupNameContext)

	// ExitLogfileGroupRef is called when exiting the logfileGroupRef production.
	ExitLogfileGroupRef(c *LogfileGroupRefContext)

	// ExitEventName is called when exiting the eventName production.
	ExitEventName(c *EventNameContext)

	// ExitEventRef is called when exiting the eventRef production.
	ExitEventRef(c *EventRefContext)

	// ExitUdfName is called when exiting the udfName production.
	ExitUdfName(c *UdfNameContext)

	// ExitServerName is called when exiting the serverName production.
	ExitServerName(c *ServerNameContext)

	// ExitServerRef is called when exiting the serverRef production.
	ExitServerRef(c *ServerRefContext)

	// ExitEngineRef is called when exiting the engineRef production.
	ExitEngineRef(c *EngineRefContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitFilterTableRef is called when exiting the filterTableRef production.
	ExitFilterTableRef(c *FilterTableRefContext)

	// ExitTableRefWithWildcard is called when exiting the tableRefWithWildcard production.
	ExitTableRefWithWildcard(c *TableRefWithWildcardContext)

	// ExitTableRef is called when exiting the tableRef production.
	ExitTableRef(c *TableRefContext)

	// ExitTableRefList is called when exiting the tableRefList production.
	ExitTableRefList(c *TableRefListContext)

	// ExitTableAliasRefList is called when exiting the tableAliasRefList production.
	ExitTableAliasRefList(c *TableAliasRefListContext)

	// ExitParameterName is called when exiting the parameterName production.
	ExitParameterName(c *ParameterNameContext)

	// ExitLabelIdentifier is called when exiting the labelIdentifier production.
	ExitLabelIdentifier(c *LabelIdentifierContext)

	// ExitLabelRef is called when exiting the labelRef production.
	ExitLabelRef(c *LabelRefContext)

	// ExitRoleIdentifier is called when exiting the roleIdentifier production.
	ExitRoleIdentifier(c *RoleIdentifierContext)

	// ExitRoleRef is called when exiting the roleRef production.
	ExitRoleRef(c *RoleRefContext)

	// ExitPluginRef is called when exiting the pluginRef production.
	ExitPluginRef(c *PluginRefContext)

	// ExitComponentRef is called when exiting the componentRef production.
	ExitComponentRef(c *ComponentRefContext)

	// ExitResourceGroupRef is called when exiting the resourceGroupRef production.
	ExitResourceGroupRef(c *ResourceGroupRefContext)

	// ExitWindowName is called when exiting the windowName production.
	ExitWindowName(c *WindowNameContext)

	// ExitPureIdentifier is called when exiting the pureIdentifier production.
	ExitPureIdentifier(c *PureIdentifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitIdentifierListWithParentheses is called when exiting the identifierListWithParentheses production.
	ExitIdentifierListWithParentheses(c *IdentifierListWithParenthesesContext)

	// ExitQualifiedIdentifier is called when exiting the qualifiedIdentifier production.
	ExitQualifiedIdentifier(c *QualifiedIdentifierContext)

	// ExitSimpleIdentifier is called when exiting the simpleIdentifier production.
	ExitSimpleIdentifier(c *SimpleIdentifierContext)

	// ExitDotIdentifier is called when exiting the dotIdentifier production.
	ExitDotIdentifier(c *DotIdentifierContext)

	// ExitUlong_number is called when exiting the ulong_number production.
	ExitUlong_number(c *Ulong_numberContext)

	// ExitReal_ulong_number is called when exiting the real_ulong_number production.
	ExitReal_ulong_number(c *Real_ulong_numberContext)

	// ExitUlonglong_number is called when exiting the ulonglong_number production.
	ExitUlonglong_number(c *Ulonglong_numberContext)

	// ExitReal_ulonglong_number is called when exiting the real_ulonglong_number production.
	ExitReal_ulonglong_number(c *Real_ulonglong_numberContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitSignedLiteral is called when exiting the signedLiteral production.
	ExitSignedLiteral(c *SignedLiteralContext)

	// ExitStringList is called when exiting the stringList production.
	ExitStringList(c *StringListContext)

	// ExitTextStringLiteral is called when exiting the textStringLiteral production.
	ExitTextStringLiteral(c *TextStringLiteralContext)

	// ExitTextString is called when exiting the textString production.
	ExitTextString(c *TextStringContext)

	// ExitTextStringHash is called when exiting the textStringHash production.
	ExitTextStringHash(c *TextStringHashContext)

	// ExitTextLiteral is called when exiting the textLiteral production.
	ExitTextLiteral(c *TextLiteralContext)

	// ExitTextStringNoLinebreak is called when exiting the textStringNoLinebreak production.
	ExitTextStringNoLinebreak(c *TextStringNoLinebreakContext)

	// ExitTextStringLiteralList is called when exiting the textStringLiteralList production.
	ExitTextStringLiteralList(c *TextStringLiteralListContext)

	// ExitNumLiteral is called when exiting the numLiteral production.
	ExitNumLiteral(c *NumLiteralContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitTemporalLiteral is called when exiting the temporalLiteral production.
	ExitTemporalLiteral(c *TemporalLiteralContext)

	// ExitFloatOptions is called when exiting the floatOptions production.
	ExitFloatOptions(c *FloatOptionsContext)

	// ExitStandardFloatOptions is called when exiting the standardFloatOptions production.
	ExitStandardFloatOptions(c *StandardFloatOptionsContext)

	// ExitPrecision is called when exiting the precision production.
	ExitPrecision(c *PrecisionContext)

	// ExitTextOrIdentifier is called when exiting the textOrIdentifier production.
	ExitTextOrIdentifier(c *TextOrIdentifierContext)

	// ExitLValueIdentifier is called when exiting the lValueIdentifier production.
	ExitLValueIdentifier(c *LValueIdentifierContext)

	// ExitRoleIdentifierOrText is called when exiting the roleIdentifierOrText production.
	ExitRoleIdentifierOrText(c *RoleIdentifierOrTextContext)

	// ExitSizeNumber is called when exiting the sizeNumber production.
	ExitSizeNumber(c *SizeNumberContext)

	// ExitParentheses is called when exiting the parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitOptionType is called when exiting the optionType production.
	ExitOptionType(c *OptionTypeContext)

	// ExitVarIdentType is called when exiting the varIdentType production.
	ExitVarIdentType(c *VarIdentTypeContext)

	// ExitSetVarIdentType is called when exiting the setVarIdentType production.
	ExitSetVarIdentType(c *SetVarIdentTypeContext)

	// ExitIdentifierKeyword is called when exiting the identifierKeyword production.
	ExitIdentifierKeyword(c *IdentifierKeywordContext)

	// ExitIdentifierKeywordsAmbiguous1RolesAndLabels is called when exiting the identifierKeywordsAmbiguous1RolesAndLabels production.
	ExitIdentifierKeywordsAmbiguous1RolesAndLabels(c *IdentifierKeywordsAmbiguous1RolesAndLabelsContext)

	// ExitIdentifierKeywordsAmbiguous2Labels is called when exiting the identifierKeywordsAmbiguous2Labels production.
	ExitIdentifierKeywordsAmbiguous2Labels(c *IdentifierKeywordsAmbiguous2LabelsContext)

	// ExitLabelKeyword is called when exiting the labelKeyword production.
	ExitLabelKeyword(c *LabelKeywordContext)

	// ExitIdentifierKeywordsAmbiguous3Roles is called when exiting the identifierKeywordsAmbiguous3Roles production.
	ExitIdentifierKeywordsAmbiguous3Roles(c *IdentifierKeywordsAmbiguous3RolesContext)

	// ExitIdentifierKeywordsUnambiguous is called when exiting the identifierKeywordsUnambiguous production.
	ExitIdentifierKeywordsUnambiguous(c *IdentifierKeywordsUnambiguousContext)

	// ExitRoleKeyword is called when exiting the roleKeyword production.
	ExitRoleKeyword(c *RoleKeywordContext)

	// ExitLValueKeyword is called when exiting the lValueKeyword production.
	ExitLValueKeyword(c *LValueKeywordContext)

	// ExitIdentifierKeywordsAmbiguous4SystemVariables is called when exiting the identifierKeywordsAmbiguous4SystemVariables production.
	ExitIdentifierKeywordsAmbiguous4SystemVariables(c *IdentifierKeywordsAmbiguous4SystemVariablesContext)

	// ExitRoleOrIdentifierKeyword is called when exiting the roleOrIdentifierKeyword production.
	ExitRoleOrIdentifierKeyword(c *RoleOrIdentifierKeywordContext)

	// ExitRoleOrLabelKeyword is called when exiting the roleOrLabelKeyword production.
	ExitRoleOrLabelKeyword(c *RoleOrLabelKeywordContext)
}
