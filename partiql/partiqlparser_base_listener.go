// Code generated from PartiQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package partiql // PartiQLParser
import "github.com/antlr4-go/antlr/v4"

// BasePartiQLParserListener is a complete listener for a parse tree produced by PartiQLParserParser.
type BasePartiQLParserListener struct{}

var _ PartiQLParserListener = &BasePartiQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePartiQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePartiQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePartiQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePartiQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BasePartiQLParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BasePartiQLParserListener) ExitScript(ctx *ScriptContext) {}

// EnterRoot is called when production root is entered.
func (s *BasePartiQLParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasePartiQLParserListener) ExitRoot(ctx *RootContext) {}

// EnterQueryDql is called when production QueryDql is entered.
func (s *BasePartiQLParserListener) EnterQueryDql(ctx *QueryDqlContext) {}

// ExitQueryDql is called when production QueryDql is exited.
func (s *BasePartiQLParserListener) ExitQueryDql(ctx *QueryDqlContext) {}

// EnterQueryDml is called when production QueryDml is entered.
func (s *BasePartiQLParserListener) EnterQueryDml(ctx *QueryDmlContext) {}

// ExitQueryDml is called when production QueryDml is exited.
func (s *BasePartiQLParserListener) ExitQueryDml(ctx *QueryDmlContext) {}

// EnterQueryDdl is called when production QueryDdl is entered.
func (s *BasePartiQLParserListener) EnterQueryDdl(ctx *QueryDdlContext) {}

// ExitQueryDdl is called when production QueryDdl is exited.
func (s *BasePartiQLParserListener) ExitQueryDdl(ctx *QueryDdlContext) {}

// EnterQueryExec is called when production QueryExec is entered.
func (s *BasePartiQLParserListener) EnterQueryExec(ctx *QueryExecContext) {}

// ExitQueryExec is called when production QueryExec is exited.
func (s *BasePartiQLParserListener) ExitQueryExec(ctx *QueryExecContext) {}

// EnterExplainOption is called when production explainOption is entered.
func (s *BasePartiQLParserListener) EnterExplainOption(ctx *ExplainOptionContext) {}

// ExitExplainOption is called when production explainOption is exited.
func (s *BasePartiQLParserListener) ExitExplainOption(ctx *ExplainOptionContext) {}

// EnterAsIdent is called when production asIdent is entered.
func (s *BasePartiQLParserListener) EnterAsIdent(ctx *AsIdentContext) {}

// ExitAsIdent is called when production asIdent is exited.
func (s *BasePartiQLParserListener) ExitAsIdent(ctx *AsIdentContext) {}

// EnterAtIdent is called when production atIdent is entered.
func (s *BasePartiQLParserListener) EnterAtIdent(ctx *AtIdentContext) {}

// ExitAtIdent is called when production atIdent is exited.
func (s *BasePartiQLParserListener) ExitAtIdent(ctx *AtIdentContext) {}

// EnterByIdent is called when production byIdent is entered.
func (s *BasePartiQLParserListener) EnterByIdent(ctx *ByIdentContext) {}

// ExitByIdent is called when production byIdent is exited.
func (s *BasePartiQLParserListener) ExitByIdent(ctx *ByIdentContext) {}

// EnterSymbolPrimitive is called when production symbolPrimitive is entered.
func (s *BasePartiQLParserListener) EnterSymbolPrimitive(ctx *SymbolPrimitiveContext) {}

// ExitSymbolPrimitive is called when production symbolPrimitive is exited.
func (s *BasePartiQLParserListener) ExitSymbolPrimitive(ctx *SymbolPrimitiveContext) {}

// EnterDql is called when production dql is entered.
func (s *BasePartiQLParserListener) EnterDql(ctx *DqlContext) {}

// ExitDql is called when production dql is exited.
func (s *BasePartiQLParserListener) ExitDql(ctx *DqlContext) {}

// EnterExecCommand is called when production execCommand is entered.
func (s *BasePartiQLParserListener) EnterExecCommand(ctx *ExecCommandContext) {}

// ExitExecCommand is called when production execCommand is exited.
func (s *BasePartiQLParserListener) ExitExecCommand(ctx *ExecCommandContext) {}

// EnterDdl is called when production ddl is entered.
func (s *BasePartiQLParserListener) EnterDdl(ctx *DdlContext) {}

// ExitDdl is called when production ddl is exited.
func (s *BasePartiQLParserListener) ExitDdl(ctx *DdlContext) {}

// EnterCreateTable is called when production CreateTable is entered.
func (s *BasePartiQLParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production CreateTable is exited.
func (s *BasePartiQLParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateIndex is called when production CreateIndex is entered.
func (s *BasePartiQLParserListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production CreateIndex is exited.
func (s *BasePartiQLParserListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterDropTable is called when production DropTable is entered.
func (s *BasePartiQLParserListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production DropTable is exited.
func (s *BasePartiQLParserListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropIndex is called when production DropIndex is entered.
func (s *BasePartiQLParserListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production DropIndex is exited.
func (s *BasePartiQLParserListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterDmlBaseWrapper is called when production DmlBaseWrapper is entered.
func (s *BasePartiQLParserListener) EnterDmlBaseWrapper(ctx *DmlBaseWrapperContext) {}

// ExitDmlBaseWrapper is called when production DmlBaseWrapper is exited.
func (s *BasePartiQLParserListener) ExitDmlBaseWrapper(ctx *DmlBaseWrapperContext) {}

// EnterDmlDelete is called when production DmlDelete is entered.
func (s *BasePartiQLParserListener) EnterDmlDelete(ctx *DmlDeleteContext) {}

// ExitDmlDelete is called when production DmlDelete is exited.
func (s *BasePartiQLParserListener) ExitDmlDelete(ctx *DmlDeleteContext) {}

// EnterDmlInsertReturning is called when production DmlInsertReturning is entered.
func (s *BasePartiQLParserListener) EnterDmlInsertReturning(ctx *DmlInsertReturningContext) {}

// ExitDmlInsertReturning is called when production DmlInsertReturning is exited.
func (s *BasePartiQLParserListener) ExitDmlInsertReturning(ctx *DmlInsertReturningContext) {}

// EnterDmlBase is called when production DmlBase is entered.
func (s *BasePartiQLParserListener) EnterDmlBase(ctx *DmlBaseContext) {}

// ExitDmlBase is called when production DmlBase is exited.
func (s *BasePartiQLParserListener) ExitDmlBase(ctx *DmlBaseContext) {}

// EnterDmlBaseCommand is called when production dmlBaseCommand is entered.
func (s *BasePartiQLParserListener) EnterDmlBaseCommand(ctx *DmlBaseCommandContext) {}

// ExitDmlBaseCommand is called when production dmlBaseCommand is exited.
func (s *BasePartiQLParserListener) ExitDmlBaseCommand(ctx *DmlBaseCommandContext) {}

// EnterPathSimple is called when production pathSimple is entered.
func (s *BasePartiQLParserListener) EnterPathSimple(ctx *PathSimpleContext) {}

// ExitPathSimple is called when production pathSimple is exited.
func (s *BasePartiQLParserListener) ExitPathSimple(ctx *PathSimpleContext) {}

// EnterPathSimpleLiteral is called when production PathSimpleLiteral is entered.
func (s *BasePartiQLParserListener) EnterPathSimpleLiteral(ctx *PathSimpleLiteralContext) {}

// ExitPathSimpleLiteral is called when production PathSimpleLiteral is exited.
func (s *BasePartiQLParserListener) ExitPathSimpleLiteral(ctx *PathSimpleLiteralContext) {}

// EnterPathSimpleSymbol is called when production PathSimpleSymbol is entered.
func (s *BasePartiQLParserListener) EnterPathSimpleSymbol(ctx *PathSimpleSymbolContext) {}

// ExitPathSimpleSymbol is called when production PathSimpleSymbol is exited.
func (s *BasePartiQLParserListener) ExitPathSimpleSymbol(ctx *PathSimpleSymbolContext) {}

// EnterPathSimpleDotSymbol is called when production PathSimpleDotSymbol is entered.
func (s *BasePartiQLParserListener) EnterPathSimpleDotSymbol(ctx *PathSimpleDotSymbolContext) {}

// ExitPathSimpleDotSymbol is called when production PathSimpleDotSymbol is exited.
func (s *BasePartiQLParserListener) ExitPathSimpleDotSymbol(ctx *PathSimpleDotSymbolContext) {}

// EnterReplaceCommand is called when production replaceCommand is entered.
func (s *BasePartiQLParserListener) EnterReplaceCommand(ctx *ReplaceCommandContext) {}

// ExitReplaceCommand is called when production replaceCommand is exited.
func (s *BasePartiQLParserListener) ExitReplaceCommand(ctx *ReplaceCommandContext) {}

// EnterUpsertCommand is called when production upsertCommand is entered.
func (s *BasePartiQLParserListener) EnterUpsertCommand(ctx *UpsertCommandContext) {}

// ExitUpsertCommand is called when production upsertCommand is exited.
func (s *BasePartiQLParserListener) ExitUpsertCommand(ctx *UpsertCommandContext) {}

// EnterRemoveCommand is called when production removeCommand is entered.
func (s *BasePartiQLParserListener) EnterRemoveCommand(ctx *RemoveCommandContext) {}

// ExitRemoveCommand is called when production removeCommand is exited.
func (s *BasePartiQLParserListener) ExitRemoveCommand(ctx *RemoveCommandContext) {}

// EnterInsertCommandReturning is called when production insertCommandReturning is entered.
func (s *BasePartiQLParserListener) EnterInsertCommandReturning(ctx *InsertCommandReturningContext) {}

// ExitInsertCommandReturning is called when production insertCommandReturning is exited.
func (s *BasePartiQLParserListener) ExitInsertCommandReturning(ctx *InsertCommandReturningContext) {}

// EnterInsertLegacy is called when production InsertLegacy is entered.
func (s *BasePartiQLParserListener) EnterInsertLegacy(ctx *InsertLegacyContext) {}

// ExitInsertLegacy is called when production InsertLegacy is exited.
func (s *BasePartiQLParserListener) ExitInsertLegacy(ctx *InsertLegacyContext) {}

// EnterInsert is called when production Insert is entered.
func (s *BasePartiQLParserListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production Insert is exited.
func (s *BasePartiQLParserListener) ExitInsert(ctx *InsertContext) {}

// EnterOnConflictLegacy is called when production OnConflictLegacy is entered.
func (s *BasePartiQLParserListener) EnterOnConflictLegacy(ctx *OnConflictLegacyContext) {}

// ExitOnConflictLegacy is called when production OnConflictLegacy is exited.
func (s *BasePartiQLParserListener) ExitOnConflictLegacy(ctx *OnConflictLegacyContext) {}

// EnterOnConflict is called when production OnConflict is entered.
func (s *BasePartiQLParserListener) EnterOnConflict(ctx *OnConflictContext) {}

// ExitOnConflict is called when production OnConflict is exited.
func (s *BasePartiQLParserListener) ExitOnConflict(ctx *OnConflictContext) {}

// EnterConflictTarget is called when production conflictTarget is entered.
func (s *BasePartiQLParserListener) EnterConflictTarget(ctx *ConflictTargetContext) {}

// ExitConflictTarget is called when production conflictTarget is exited.
func (s *BasePartiQLParserListener) ExitConflictTarget(ctx *ConflictTargetContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BasePartiQLParserListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BasePartiQLParserListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterConflictAction is called when production conflictAction is entered.
func (s *BasePartiQLParserListener) EnterConflictAction(ctx *ConflictActionContext) {}

// ExitConflictAction is called when production conflictAction is exited.
func (s *BasePartiQLParserListener) ExitConflictAction(ctx *ConflictActionContext) {}

// EnterDoReplace is called when production doReplace is entered.
func (s *BasePartiQLParserListener) EnterDoReplace(ctx *DoReplaceContext) {}

// ExitDoReplace is called when production doReplace is exited.
func (s *BasePartiQLParserListener) ExitDoReplace(ctx *DoReplaceContext) {}

// EnterDoUpdate is called when production doUpdate is entered.
func (s *BasePartiQLParserListener) EnterDoUpdate(ctx *DoUpdateContext) {}

// ExitDoUpdate is called when production doUpdate is exited.
func (s *BasePartiQLParserListener) ExitDoUpdate(ctx *DoUpdateContext) {}

// EnterUpdateClause is called when production updateClause is entered.
func (s *BasePartiQLParserListener) EnterUpdateClause(ctx *UpdateClauseContext) {}

// ExitUpdateClause is called when production updateClause is exited.
func (s *BasePartiQLParserListener) ExitUpdateClause(ctx *UpdateClauseContext) {}

// EnterSetCommand is called when production setCommand is entered.
func (s *BasePartiQLParserListener) EnterSetCommand(ctx *SetCommandContext) {}

// ExitSetCommand is called when production setCommand is exited.
func (s *BasePartiQLParserListener) ExitSetCommand(ctx *SetCommandContext) {}

// EnterSetAssignment is called when production setAssignment is entered.
func (s *BasePartiQLParserListener) EnterSetAssignment(ctx *SetAssignmentContext) {}

// ExitSetAssignment is called when production setAssignment is exited.
func (s *BasePartiQLParserListener) ExitSetAssignment(ctx *SetAssignmentContext) {}

// EnterDeleteCommand is called when production deleteCommand is entered.
func (s *BasePartiQLParserListener) EnterDeleteCommand(ctx *DeleteCommandContext) {}

// ExitDeleteCommand is called when production deleteCommand is exited.
func (s *BasePartiQLParserListener) ExitDeleteCommand(ctx *DeleteCommandContext) {}

// EnterReturningClause is called when production returningClause is entered.
func (s *BasePartiQLParserListener) EnterReturningClause(ctx *ReturningClauseContext) {}

// ExitReturningClause is called when production returningClause is exited.
func (s *BasePartiQLParserListener) ExitReturningClause(ctx *ReturningClauseContext) {}

// EnterReturningColumn is called when production returningColumn is entered.
func (s *BasePartiQLParserListener) EnterReturningColumn(ctx *ReturningColumnContext) {}

// ExitReturningColumn is called when production returningColumn is exited.
func (s *BasePartiQLParserListener) ExitReturningColumn(ctx *ReturningColumnContext) {}

// EnterFromClauseSimpleExplicit is called when production FromClauseSimpleExplicit is entered.
func (s *BasePartiQLParserListener) EnterFromClauseSimpleExplicit(ctx *FromClauseSimpleExplicitContext) {
}

// ExitFromClauseSimpleExplicit is called when production FromClauseSimpleExplicit is exited.
func (s *BasePartiQLParserListener) ExitFromClauseSimpleExplicit(ctx *FromClauseSimpleExplicitContext) {
}

// EnterFromClauseSimpleImplicit is called when production FromClauseSimpleImplicit is entered.
func (s *BasePartiQLParserListener) EnterFromClauseSimpleImplicit(ctx *FromClauseSimpleImplicitContext) {
}

// ExitFromClauseSimpleImplicit is called when production FromClauseSimpleImplicit is exited.
func (s *BasePartiQLParserListener) ExitFromClauseSimpleImplicit(ctx *FromClauseSimpleImplicitContext) {
}

// EnterWhereClause is called when production whereClause is entered.
func (s *BasePartiQLParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BasePartiQLParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterSelectAll is called when production SelectAll is entered.
func (s *BasePartiQLParserListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production SelectAll is exited.
func (s *BasePartiQLParserListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterSelectItems is called when production SelectItems is entered.
func (s *BasePartiQLParserListener) EnterSelectItems(ctx *SelectItemsContext) {}

// ExitSelectItems is called when production SelectItems is exited.
func (s *BasePartiQLParserListener) ExitSelectItems(ctx *SelectItemsContext) {}

// EnterSelectValue is called when production SelectValue is entered.
func (s *BasePartiQLParserListener) EnterSelectValue(ctx *SelectValueContext) {}

// ExitSelectValue is called when production SelectValue is exited.
func (s *BasePartiQLParserListener) ExitSelectValue(ctx *SelectValueContext) {}

// EnterSelectPivot is called when production SelectPivot is entered.
func (s *BasePartiQLParserListener) EnterSelectPivot(ctx *SelectPivotContext) {}

// ExitSelectPivot is called when production SelectPivot is exited.
func (s *BasePartiQLParserListener) ExitSelectPivot(ctx *SelectPivotContext) {}

// EnterProjectionItems is called when production projectionItems is entered.
func (s *BasePartiQLParserListener) EnterProjectionItems(ctx *ProjectionItemsContext) {}

// ExitProjectionItems is called when production projectionItems is exited.
func (s *BasePartiQLParserListener) ExitProjectionItems(ctx *ProjectionItemsContext) {}

// EnterProjectionItem is called when production projectionItem is entered.
func (s *BasePartiQLParserListener) EnterProjectionItem(ctx *ProjectionItemContext) {}

// ExitProjectionItem is called when production projectionItem is exited.
func (s *BasePartiQLParserListener) ExitProjectionItem(ctx *ProjectionItemContext) {}

// EnterSetQuantifierStrategy is called when production setQuantifierStrategy is entered.
func (s *BasePartiQLParserListener) EnterSetQuantifierStrategy(ctx *SetQuantifierStrategyContext) {}

// ExitSetQuantifierStrategy is called when production setQuantifierStrategy is exited.
func (s *BasePartiQLParserListener) ExitSetQuantifierStrategy(ctx *SetQuantifierStrategyContext) {}

// EnterLetClause is called when production letClause is entered.
func (s *BasePartiQLParserListener) EnterLetClause(ctx *LetClauseContext) {}

// ExitLetClause is called when production letClause is exited.
func (s *BasePartiQLParserListener) ExitLetClause(ctx *LetClauseContext) {}

// EnterLetBinding is called when production letBinding is entered.
func (s *BasePartiQLParserListener) EnterLetBinding(ctx *LetBindingContext) {}

// ExitLetBinding is called when production letBinding is exited.
func (s *BasePartiQLParserListener) ExitLetBinding(ctx *LetBindingContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BasePartiQLParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BasePartiQLParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderSortSpec is called when production orderSortSpec is entered.
func (s *BasePartiQLParserListener) EnterOrderSortSpec(ctx *OrderSortSpecContext) {}

// ExitOrderSortSpec is called when production orderSortSpec is exited.
func (s *BasePartiQLParserListener) ExitOrderSortSpec(ctx *OrderSortSpecContext) {}

// EnterGroupClause is called when production groupClause is entered.
func (s *BasePartiQLParserListener) EnterGroupClause(ctx *GroupClauseContext) {}

// ExitGroupClause is called when production groupClause is exited.
func (s *BasePartiQLParserListener) ExitGroupClause(ctx *GroupClauseContext) {}

// EnterGroupAlias is called when production groupAlias is entered.
func (s *BasePartiQLParserListener) EnterGroupAlias(ctx *GroupAliasContext) {}

// ExitGroupAlias is called when production groupAlias is exited.
func (s *BasePartiQLParserListener) ExitGroupAlias(ctx *GroupAliasContext) {}

// EnterGroupKey is called when production groupKey is entered.
func (s *BasePartiQLParserListener) EnterGroupKey(ctx *GroupKeyContext) {}

// ExitGroupKey is called when production groupKey is exited.
func (s *BasePartiQLParserListener) ExitGroupKey(ctx *GroupKeyContext) {}

// EnterOver is called when production over is entered.
func (s *BasePartiQLParserListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BasePartiQLParserListener) ExitOver(ctx *OverContext) {}

// EnterWindowPartitionList is called when production windowPartitionList is entered.
func (s *BasePartiQLParserListener) EnterWindowPartitionList(ctx *WindowPartitionListContext) {}

// ExitWindowPartitionList is called when production windowPartitionList is exited.
func (s *BasePartiQLParserListener) ExitWindowPartitionList(ctx *WindowPartitionListContext) {}

// EnterWindowSortSpecList is called when production windowSortSpecList is entered.
func (s *BasePartiQLParserListener) EnterWindowSortSpecList(ctx *WindowSortSpecListContext) {}

// ExitWindowSortSpecList is called when production windowSortSpecList is exited.
func (s *BasePartiQLParserListener) ExitWindowSortSpecList(ctx *WindowSortSpecListContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BasePartiQLParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BasePartiQLParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BasePartiQLParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BasePartiQLParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterWhereClauseSelect is called when production whereClauseSelect is entered.
func (s *BasePartiQLParserListener) EnterWhereClauseSelect(ctx *WhereClauseSelectContext) {}

// ExitWhereClauseSelect is called when production whereClauseSelect is exited.
func (s *BasePartiQLParserListener) ExitWhereClauseSelect(ctx *WhereClauseSelectContext) {}

// EnterOffsetByClause is called when production offsetByClause is entered.
func (s *BasePartiQLParserListener) EnterOffsetByClause(ctx *OffsetByClauseContext) {}

// ExitOffsetByClause is called when production offsetByClause is exited.
func (s *BasePartiQLParserListener) ExitOffsetByClause(ctx *OffsetByClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BasePartiQLParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BasePartiQLParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterGpmlPattern is called when production gpmlPattern is entered.
func (s *BasePartiQLParserListener) EnterGpmlPattern(ctx *GpmlPatternContext) {}

// ExitGpmlPattern is called when production gpmlPattern is exited.
func (s *BasePartiQLParserListener) ExitGpmlPattern(ctx *GpmlPatternContext) {}

// EnterGpmlPatternList is called when production gpmlPatternList is entered.
func (s *BasePartiQLParserListener) EnterGpmlPatternList(ctx *GpmlPatternListContext) {}

// ExitGpmlPatternList is called when production gpmlPatternList is exited.
func (s *BasePartiQLParserListener) ExitGpmlPatternList(ctx *GpmlPatternListContext) {}

// EnterMatchPattern is called when production matchPattern is entered.
func (s *BasePartiQLParserListener) EnterMatchPattern(ctx *MatchPatternContext) {}

// ExitMatchPattern is called when production matchPattern is exited.
func (s *BasePartiQLParserListener) ExitMatchPattern(ctx *MatchPatternContext) {}

// EnterGraphPart is called when production graphPart is entered.
func (s *BasePartiQLParserListener) EnterGraphPart(ctx *GraphPartContext) {}

// ExitGraphPart is called when production graphPart is exited.
func (s *BasePartiQLParserListener) ExitGraphPart(ctx *GraphPartContext) {}

// EnterSelectorBasic is called when production SelectorBasic is entered.
func (s *BasePartiQLParserListener) EnterSelectorBasic(ctx *SelectorBasicContext) {}

// ExitSelectorBasic is called when production SelectorBasic is exited.
func (s *BasePartiQLParserListener) ExitSelectorBasic(ctx *SelectorBasicContext) {}

// EnterSelectorAny is called when production SelectorAny is entered.
func (s *BasePartiQLParserListener) EnterSelectorAny(ctx *SelectorAnyContext) {}

// ExitSelectorAny is called when production SelectorAny is exited.
func (s *BasePartiQLParserListener) ExitSelectorAny(ctx *SelectorAnyContext) {}

// EnterSelectorShortest is called when production SelectorShortest is entered.
func (s *BasePartiQLParserListener) EnterSelectorShortest(ctx *SelectorShortestContext) {}

// ExitSelectorShortest is called when production SelectorShortest is exited.
func (s *BasePartiQLParserListener) ExitSelectorShortest(ctx *SelectorShortestContext) {}

// EnterPatternPathVariable is called when production patternPathVariable is entered.
func (s *BasePartiQLParserListener) EnterPatternPathVariable(ctx *PatternPathVariableContext) {}

// ExitPatternPathVariable is called when production patternPathVariable is exited.
func (s *BasePartiQLParserListener) ExitPatternPathVariable(ctx *PatternPathVariableContext) {}

// EnterPatternRestrictor is called when production patternRestrictor is entered.
func (s *BasePartiQLParserListener) EnterPatternRestrictor(ctx *PatternRestrictorContext) {}

// ExitPatternRestrictor is called when production patternRestrictor is exited.
func (s *BasePartiQLParserListener) ExitPatternRestrictor(ctx *PatternRestrictorContext) {}

// EnterNode is called when production node is entered.
func (s *BasePartiQLParserListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BasePartiQLParserListener) ExitNode(ctx *NodeContext) {}

// EnterEdgeWithSpec is called when production EdgeWithSpec is entered.
func (s *BasePartiQLParserListener) EnterEdgeWithSpec(ctx *EdgeWithSpecContext) {}

// ExitEdgeWithSpec is called when production EdgeWithSpec is exited.
func (s *BasePartiQLParserListener) ExitEdgeWithSpec(ctx *EdgeWithSpecContext) {}

// EnterEdgeAbbreviated is called when production EdgeAbbreviated is entered.
func (s *BasePartiQLParserListener) EnterEdgeAbbreviated(ctx *EdgeAbbreviatedContext) {}

// ExitEdgeAbbreviated is called when production EdgeAbbreviated is exited.
func (s *BasePartiQLParserListener) ExitEdgeAbbreviated(ctx *EdgeAbbreviatedContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BasePartiQLParserListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BasePartiQLParserListener) ExitPattern(ctx *PatternContext) {}

// EnterPatternQuantifier is called when production patternQuantifier is entered.
func (s *BasePartiQLParserListener) EnterPatternQuantifier(ctx *PatternQuantifierContext) {}

// ExitPatternQuantifier is called when production patternQuantifier is exited.
func (s *BasePartiQLParserListener) ExitPatternQuantifier(ctx *PatternQuantifierContext) {}

// EnterEdgeSpecRight is called when production EdgeSpecRight is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpecRight(ctx *EdgeSpecRightContext) {}

// ExitEdgeSpecRight is called when production EdgeSpecRight is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpecRight(ctx *EdgeSpecRightContext) {}

// EnterEdgeSpecUndirected is called when production EdgeSpecUndirected is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpecUndirected(ctx *EdgeSpecUndirectedContext) {}

// ExitEdgeSpecUndirected is called when production EdgeSpecUndirected is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpecUndirected(ctx *EdgeSpecUndirectedContext) {}

// EnterEdgeSpecLeft is called when production EdgeSpecLeft is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpecLeft(ctx *EdgeSpecLeftContext) {}

// ExitEdgeSpecLeft is called when production EdgeSpecLeft is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpecLeft(ctx *EdgeSpecLeftContext) {}

// EnterEdgeSpecUndirectedRight is called when production EdgeSpecUndirectedRight is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpecUndirectedRight(ctx *EdgeSpecUndirectedRightContext) {
}

// ExitEdgeSpecUndirectedRight is called when production EdgeSpecUndirectedRight is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpecUndirectedRight(ctx *EdgeSpecUndirectedRightContext) {
}

// EnterEdgeSpecUndirectedLeft is called when production EdgeSpecUndirectedLeft is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpecUndirectedLeft(ctx *EdgeSpecUndirectedLeftContext) {}

// ExitEdgeSpecUndirectedLeft is called when production EdgeSpecUndirectedLeft is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpecUndirectedLeft(ctx *EdgeSpecUndirectedLeftContext) {}

// EnterEdgeSpecBidirectional is called when production EdgeSpecBidirectional is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpecBidirectional(ctx *EdgeSpecBidirectionalContext) {}

// ExitEdgeSpecBidirectional is called when production EdgeSpecBidirectional is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpecBidirectional(ctx *EdgeSpecBidirectionalContext) {}

// EnterEdgeSpecUndirectedBidirectional is called when production EdgeSpecUndirectedBidirectional is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpecUndirectedBidirectional(ctx *EdgeSpecUndirectedBidirectionalContext) {
}

// ExitEdgeSpecUndirectedBidirectional is called when production EdgeSpecUndirectedBidirectional is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpecUndirectedBidirectional(ctx *EdgeSpecUndirectedBidirectionalContext) {
}

// EnterEdgeSpec is called when production edgeSpec is entered.
func (s *BasePartiQLParserListener) EnterEdgeSpec(ctx *EdgeSpecContext) {}

// ExitEdgeSpec is called when production edgeSpec is exited.
func (s *BasePartiQLParserListener) ExitEdgeSpec(ctx *EdgeSpecContext) {}

// EnterPatternPartLabel is called when production patternPartLabel is entered.
func (s *BasePartiQLParserListener) EnterPatternPartLabel(ctx *PatternPartLabelContext) {}

// ExitPatternPartLabel is called when production patternPartLabel is exited.
func (s *BasePartiQLParserListener) ExitPatternPartLabel(ctx *PatternPartLabelContext) {}

// EnterEdgeAbbrev is called when production edgeAbbrev is entered.
func (s *BasePartiQLParserListener) EnterEdgeAbbrev(ctx *EdgeAbbrevContext) {}

// ExitEdgeAbbrev is called when production edgeAbbrev is exited.
func (s *BasePartiQLParserListener) ExitEdgeAbbrev(ctx *EdgeAbbrevContext) {}

// EnterTableWrapped is called when production TableWrapped is entered.
func (s *BasePartiQLParserListener) EnterTableWrapped(ctx *TableWrappedContext) {}

// ExitTableWrapped is called when production TableWrapped is exited.
func (s *BasePartiQLParserListener) ExitTableWrapped(ctx *TableWrappedContext) {}

// EnterTableCrossJoin is called when production TableCrossJoin is entered.
func (s *BasePartiQLParserListener) EnterTableCrossJoin(ctx *TableCrossJoinContext) {}

// ExitTableCrossJoin is called when production TableCrossJoin is exited.
func (s *BasePartiQLParserListener) ExitTableCrossJoin(ctx *TableCrossJoinContext) {}

// EnterTableQualifiedJoin is called when production TableQualifiedJoin is entered.
func (s *BasePartiQLParserListener) EnterTableQualifiedJoin(ctx *TableQualifiedJoinContext) {}

// ExitTableQualifiedJoin is called when production TableQualifiedJoin is exited.
func (s *BasePartiQLParserListener) ExitTableQualifiedJoin(ctx *TableQualifiedJoinContext) {}

// EnterTableRefBase is called when production TableRefBase is entered.
func (s *BasePartiQLParserListener) EnterTableRefBase(ctx *TableRefBaseContext) {}

// ExitTableRefBase is called when production TableRefBase is exited.
func (s *BasePartiQLParserListener) ExitTableRefBase(ctx *TableRefBaseContext) {}

// EnterTableNonJoin is called when production tableNonJoin is entered.
func (s *BasePartiQLParserListener) EnterTableNonJoin(ctx *TableNonJoinContext) {}

// ExitTableNonJoin is called when production tableNonJoin is exited.
func (s *BasePartiQLParserListener) ExitTableNonJoin(ctx *TableNonJoinContext) {}

// EnterTableBaseRefSymbol is called when production TableBaseRefSymbol is entered.
func (s *BasePartiQLParserListener) EnterTableBaseRefSymbol(ctx *TableBaseRefSymbolContext) {}

// ExitTableBaseRefSymbol is called when production TableBaseRefSymbol is exited.
func (s *BasePartiQLParserListener) ExitTableBaseRefSymbol(ctx *TableBaseRefSymbolContext) {}

// EnterTableBaseRefClauses is called when production TableBaseRefClauses is entered.
func (s *BasePartiQLParserListener) EnterTableBaseRefClauses(ctx *TableBaseRefClausesContext) {}

// ExitTableBaseRefClauses is called when production TableBaseRefClauses is exited.
func (s *BasePartiQLParserListener) ExitTableBaseRefClauses(ctx *TableBaseRefClausesContext) {}

// EnterTableBaseRefMatch is called when production TableBaseRefMatch is entered.
func (s *BasePartiQLParserListener) EnterTableBaseRefMatch(ctx *TableBaseRefMatchContext) {}

// ExitTableBaseRefMatch is called when production TableBaseRefMatch is exited.
func (s *BasePartiQLParserListener) ExitTableBaseRefMatch(ctx *TableBaseRefMatchContext) {}

// EnterTableUnpivot is called when production tableUnpivot is entered.
func (s *BasePartiQLParserListener) EnterTableUnpivot(ctx *TableUnpivotContext) {}

// ExitTableUnpivot is called when production tableUnpivot is exited.
func (s *BasePartiQLParserListener) ExitTableUnpivot(ctx *TableUnpivotContext) {}

// EnterJoinRhsBase is called when production JoinRhsBase is entered.
func (s *BasePartiQLParserListener) EnterJoinRhsBase(ctx *JoinRhsBaseContext) {}

// ExitJoinRhsBase is called when production JoinRhsBase is exited.
func (s *BasePartiQLParserListener) ExitJoinRhsBase(ctx *JoinRhsBaseContext) {}

// EnterJoinRhsTableJoined is called when production JoinRhsTableJoined is entered.
func (s *BasePartiQLParserListener) EnterJoinRhsTableJoined(ctx *JoinRhsTableJoinedContext) {}

// ExitJoinRhsTableJoined is called when production JoinRhsTableJoined is exited.
func (s *BasePartiQLParserListener) ExitJoinRhsTableJoined(ctx *JoinRhsTableJoinedContext) {}

// EnterJoinSpec is called when production joinSpec is entered.
func (s *BasePartiQLParserListener) EnterJoinSpec(ctx *JoinSpecContext) {}

// ExitJoinSpec is called when production joinSpec is exited.
func (s *BasePartiQLParserListener) ExitJoinSpec(ctx *JoinSpecContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BasePartiQLParserListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BasePartiQLParserListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePartiQLParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePartiQLParserListener) ExitExpr(ctx *ExprContext) {}

// EnterIntersect is called when production Intersect is entered.
func (s *BasePartiQLParserListener) EnterIntersect(ctx *IntersectContext) {}

// ExitIntersect is called when production Intersect is exited.
func (s *BasePartiQLParserListener) ExitIntersect(ctx *IntersectContext) {}

// EnterQueryBase is called when production QueryBase is entered.
func (s *BasePartiQLParserListener) EnterQueryBase(ctx *QueryBaseContext) {}

// ExitQueryBase is called when production QueryBase is exited.
func (s *BasePartiQLParserListener) ExitQueryBase(ctx *QueryBaseContext) {}

// EnterExcept is called when production Except is entered.
func (s *BasePartiQLParserListener) EnterExcept(ctx *ExceptContext) {}

// ExitExcept is called when production Except is exited.
func (s *BasePartiQLParserListener) ExitExcept(ctx *ExceptContext) {}

// EnterUnion is called when production Union is entered.
func (s *BasePartiQLParserListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production Union is exited.
func (s *BasePartiQLParserListener) ExitUnion(ctx *UnionContext) {}

// EnterSfwQuery is called when production SfwQuery is entered.
func (s *BasePartiQLParserListener) EnterSfwQuery(ctx *SfwQueryContext) {}

// ExitSfwQuery is called when production SfwQuery is exited.
func (s *BasePartiQLParserListener) ExitSfwQuery(ctx *SfwQueryContext) {}

// EnterSfwBase is called when production SfwBase is entered.
func (s *BasePartiQLParserListener) EnterSfwBase(ctx *SfwBaseContext) {}

// ExitSfwBase is called when production SfwBase is exited.
func (s *BasePartiQLParserListener) ExitSfwBase(ctx *SfwBaseContext) {}

// EnterOr is called when production Or is entered.
func (s *BasePartiQLParserListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BasePartiQLParserListener) ExitOr(ctx *OrContext) {}

// EnterExprOrBase is called when production ExprOrBase is entered.
func (s *BasePartiQLParserListener) EnterExprOrBase(ctx *ExprOrBaseContext) {}

// ExitExprOrBase is called when production ExprOrBase is exited.
func (s *BasePartiQLParserListener) ExitExprOrBase(ctx *ExprOrBaseContext) {}

// EnterExprAndBase is called when production ExprAndBase is entered.
func (s *BasePartiQLParserListener) EnterExprAndBase(ctx *ExprAndBaseContext) {}

// ExitExprAndBase is called when production ExprAndBase is exited.
func (s *BasePartiQLParserListener) ExitExprAndBase(ctx *ExprAndBaseContext) {}

// EnterAnd is called when production And is entered.
func (s *BasePartiQLParserListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BasePartiQLParserListener) ExitAnd(ctx *AndContext) {}

// EnterNot is called when production Not is entered.
func (s *BasePartiQLParserListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BasePartiQLParserListener) ExitNot(ctx *NotContext) {}

// EnterExprNotBase is called when production ExprNotBase is entered.
func (s *BasePartiQLParserListener) EnterExprNotBase(ctx *ExprNotBaseContext) {}

// ExitExprNotBase is called when production ExprNotBase is exited.
func (s *BasePartiQLParserListener) ExitExprNotBase(ctx *ExprNotBaseContext) {}

// EnterPredicateIn is called when production PredicateIn is entered.
func (s *BasePartiQLParserListener) EnterPredicateIn(ctx *PredicateInContext) {}

// ExitPredicateIn is called when production PredicateIn is exited.
func (s *BasePartiQLParserListener) ExitPredicateIn(ctx *PredicateInContext) {}

// EnterPredicateBetween is called when production PredicateBetween is entered.
func (s *BasePartiQLParserListener) EnterPredicateBetween(ctx *PredicateBetweenContext) {}

// ExitPredicateBetween is called when production PredicateBetween is exited.
func (s *BasePartiQLParserListener) ExitPredicateBetween(ctx *PredicateBetweenContext) {}

// EnterPredicateBase is called when production PredicateBase is entered.
func (s *BasePartiQLParserListener) EnterPredicateBase(ctx *PredicateBaseContext) {}

// ExitPredicateBase is called when production PredicateBase is exited.
func (s *BasePartiQLParserListener) ExitPredicateBase(ctx *PredicateBaseContext) {}

// EnterPredicateComparison is called when production PredicateComparison is entered.
func (s *BasePartiQLParserListener) EnterPredicateComparison(ctx *PredicateComparisonContext) {}

// ExitPredicateComparison is called when production PredicateComparison is exited.
func (s *BasePartiQLParserListener) ExitPredicateComparison(ctx *PredicateComparisonContext) {}

// EnterPredicateIs is called when production PredicateIs is entered.
func (s *BasePartiQLParserListener) EnterPredicateIs(ctx *PredicateIsContext) {}

// ExitPredicateIs is called when production PredicateIs is exited.
func (s *BasePartiQLParserListener) ExitPredicateIs(ctx *PredicateIsContext) {}

// EnterPredicateLike is called when production PredicateLike is entered.
func (s *BasePartiQLParserListener) EnterPredicateLike(ctx *PredicateLikeContext) {}

// ExitPredicateLike is called when production PredicateLike is exited.
func (s *BasePartiQLParserListener) ExitPredicateLike(ctx *PredicateLikeContext) {}

// EnterMathOp00 is called when production mathOp00 is entered.
func (s *BasePartiQLParserListener) EnterMathOp00(ctx *MathOp00Context) {}

// ExitMathOp00 is called when production mathOp00 is exited.
func (s *BasePartiQLParserListener) ExitMathOp00(ctx *MathOp00Context) {}

// EnterMathOp01 is called when production mathOp01 is entered.
func (s *BasePartiQLParserListener) EnterMathOp01(ctx *MathOp01Context) {}

// ExitMathOp01 is called when production mathOp01 is exited.
func (s *BasePartiQLParserListener) ExitMathOp01(ctx *MathOp01Context) {}

// EnterMathOp02 is called when production mathOp02 is entered.
func (s *BasePartiQLParserListener) EnterMathOp02(ctx *MathOp02Context) {}

// ExitMathOp02 is called when production mathOp02 is exited.
func (s *BasePartiQLParserListener) ExitMathOp02(ctx *MathOp02Context) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BasePartiQLParserListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BasePartiQLParserListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterExprPrimaryPath is called when production ExprPrimaryPath is entered.
func (s *BasePartiQLParserListener) EnterExprPrimaryPath(ctx *ExprPrimaryPathContext) {}

// ExitExprPrimaryPath is called when production ExprPrimaryPath is exited.
func (s *BasePartiQLParserListener) ExitExprPrimaryPath(ctx *ExprPrimaryPathContext) {}

// EnterExprPrimaryBase is called when production ExprPrimaryBase is entered.
func (s *BasePartiQLParserListener) EnterExprPrimaryBase(ctx *ExprPrimaryBaseContext) {}

// ExitExprPrimaryBase is called when production ExprPrimaryBase is exited.
func (s *BasePartiQLParserListener) ExitExprPrimaryBase(ctx *ExprPrimaryBaseContext) {}

// EnterExprTermWrappedQuery is called when production ExprTermWrappedQuery is entered.
func (s *BasePartiQLParserListener) EnterExprTermWrappedQuery(ctx *ExprTermWrappedQueryContext) {}

// ExitExprTermWrappedQuery is called when production ExprTermWrappedQuery is exited.
func (s *BasePartiQLParserListener) ExitExprTermWrappedQuery(ctx *ExprTermWrappedQueryContext) {}

// EnterExprTermBase is called when production ExprTermBase is entered.
func (s *BasePartiQLParserListener) EnterExprTermBase(ctx *ExprTermBaseContext) {}

// ExitExprTermBase is called when production ExprTermBase is exited.
func (s *BasePartiQLParserListener) ExitExprTermBase(ctx *ExprTermBaseContext) {}

// EnterNullIf is called when production nullIf is entered.
func (s *BasePartiQLParserListener) EnterNullIf(ctx *NullIfContext) {}

// ExitNullIf is called when production nullIf is exited.
func (s *BasePartiQLParserListener) ExitNullIf(ctx *NullIfContext) {}

// EnterCoalesce is called when production coalesce is entered.
func (s *BasePartiQLParserListener) EnterCoalesce(ctx *CoalesceContext) {}

// ExitCoalesce is called when production coalesce is exited.
func (s *BasePartiQLParserListener) ExitCoalesce(ctx *CoalesceContext) {}

// EnterCaseExpr is called when production caseExpr is entered.
func (s *BasePartiQLParserListener) EnterCaseExpr(ctx *CaseExprContext) {}

// ExitCaseExpr is called when production caseExpr is exited.
func (s *BasePartiQLParserListener) ExitCaseExpr(ctx *CaseExprContext) {}

// EnterValues is called when production values is entered.
func (s *BasePartiQLParserListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BasePartiQLParserListener) ExitValues(ctx *ValuesContext) {}

// EnterValueRow is called when production valueRow is entered.
func (s *BasePartiQLParserListener) EnterValueRow(ctx *ValueRowContext) {}

// ExitValueRow is called when production valueRow is exited.
func (s *BasePartiQLParserListener) ExitValueRow(ctx *ValueRowContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BasePartiQLParserListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BasePartiQLParserListener) ExitValueList(ctx *ValueListContext) {}

// EnterSequenceConstructor is called when production sequenceConstructor is entered.
func (s *BasePartiQLParserListener) EnterSequenceConstructor(ctx *SequenceConstructorContext) {}

// ExitSequenceConstructor is called when production sequenceConstructor is exited.
func (s *BasePartiQLParserListener) ExitSequenceConstructor(ctx *SequenceConstructorContext) {}

// EnterSubstring is called when production substring is entered.
func (s *BasePartiQLParserListener) EnterSubstring(ctx *SubstringContext) {}

// ExitSubstring is called when production substring is exited.
func (s *BasePartiQLParserListener) ExitSubstring(ctx *SubstringContext) {}

// EnterCountAll is called when production CountAll is entered.
func (s *BasePartiQLParserListener) EnterCountAll(ctx *CountAllContext) {}

// ExitCountAll is called when production CountAll is exited.
func (s *BasePartiQLParserListener) ExitCountAll(ctx *CountAllContext) {}

// EnterAggregateBase is called when production AggregateBase is entered.
func (s *BasePartiQLParserListener) EnterAggregateBase(ctx *AggregateBaseContext) {}

// ExitAggregateBase is called when production AggregateBase is exited.
func (s *BasePartiQLParserListener) ExitAggregateBase(ctx *AggregateBaseContext) {}

// EnterLagLeadFunction is called when production LagLeadFunction is entered.
func (s *BasePartiQLParserListener) EnterLagLeadFunction(ctx *LagLeadFunctionContext) {}

// ExitLagLeadFunction is called when production LagLeadFunction is exited.
func (s *BasePartiQLParserListener) ExitLagLeadFunction(ctx *LagLeadFunctionContext) {}

// EnterCast is called when production cast is entered.
func (s *BasePartiQLParserListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BasePartiQLParserListener) ExitCast(ctx *CastContext) {}

// EnterCanLosslessCast is called when production canLosslessCast is entered.
func (s *BasePartiQLParserListener) EnterCanLosslessCast(ctx *CanLosslessCastContext) {}

// ExitCanLosslessCast is called when production canLosslessCast is exited.
func (s *BasePartiQLParserListener) ExitCanLosslessCast(ctx *CanLosslessCastContext) {}

// EnterCanCast is called when production canCast is entered.
func (s *BasePartiQLParserListener) EnterCanCast(ctx *CanCastContext) {}

// ExitCanCast is called when production canCast is exited.
func (s *BasePartiQLParserListener) ExitCanCast(ctx *CanCastContext) {}

// EnterExtract is called when production extract is entered.
func (s *BasePartiQLParserListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BasePartiQLParserListener) ExitExtract(ctx *ExtractContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *BasePartiQLParserListener) EnterTrimFunction(ctx *TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *BasePartiQLParserListener) ExitTrimFunction(ctx *TrimFunctionContext) {}

// EnterDateFunction is called when production dateFunction is entered.
func (s *BasePartiQLParserListener) EnterDateFunction(ctx *DateFunctionContext) {}

// ExitDateFunction is called when production dateFunction is exited.
func (s *BasePartiQLParserListener) ExitDateFunction(ctx *DateFunctionContext) {}

// EnterFunctionCallReserved is called when production FunctionCallReserved is entered.
func (s *BasePartiQLParserListener) EnterFunctionCallReserved(ctx *FunctionCallReservedContext) {}

// ExitFunctionCallReserved is called when production FunctionCallReserved is exited.
func (s *BasePartiQLParserListener) ExitFunctionCallReserved(ctx *FunctionCallReservedContext) {}

// EnterFunctionCallIdent is called when production FunctionCallIdent is entered.
func (s *BasePartiQLParserListener) EnterFunctionCallIdent(ctx *FunctionCallIdentContext) {}

// ExitFunctionCallIdent is called when production FunctionCallIdent is exited.
func (s *BasePartiQLParserListener) ExitFunctionCallIdent(ctx *FunctionCallIdentContext) {}

// EnterPathStepIndexExpr is called when production PathStepIndexExpr is entered.
func (s *BasePartiQLParserListener) EnterPathStepIndexExpr(ctx *PathStepIndexExprContext) {}

// ExitPathStepIndexExpr is called when production PathStepIndexExpr is exited.
func (s *BasePartiQLParserListener) ExitPathStepIndexExpr(ctx *PathStepIndexExprContext) {}

// EnterPathStepIndexAll is called when production PathStepIndexAll is entered.
func (s *BasePartiQLParserListener) EnterPathStepIndexAll(ctx *PathStepIndexAllContext) {}

// ExitPathStepIndexAll is called when production PathStepIndexAll is exited.
func (s *BasePartiQLParserListener) ExitPathStepIndexAll(ctx *PathStepIndexAllContext) {}

// EnterPathStepDotExpr is called when production PathStepDotExpr is entered.
func (s *BasePartiQLParserListener) EnterPathStepDotExpr(ctx *PathStepDotExprContext) {}

// ExitPathStepDotExpr is called when production PathStepDotExpr is exited.
func (s *BasePartiQLParserListener) ExitPathStepDotExpr(ctx *PathStepDotExprContext) {}

// EnterPathStepDotAll is called when production PathStepDotAll is entered.
func (s *BasePartiQLParserListener) EnterPathStepDotAll(ctx *PathStepDotAllContext) {}

// ExitPathStepDotAll is called when production PathStepDotAll is exited.
func (s *BasePartiQLParserListener) ExitPathStepDotAll(ctx *PathStepDotAllContext) {}

// EnterExprGraphMatchMany is called when production exprGraphMatchMany is entered.
func (s *BasePartiQLParserListener) EnterExprGraphMatchMany(ctx *ExprGraphMatchManyContext) {}

// ExitExprGraphMatchMany is called when production exprGraphMatchMany is exited.
func (s *BasePartiQLParserListener) ExitExprGraphMatchMany(ctx *ExprGraphMatchManyContext) {}

// EnterExprGraphMatchOne is called when production exprGraphMatchOne is entered.
func (s *BasePartiQLParserListener) EnterExprGraphMatchOne(ctx *ExprGraphMatchOneContext) {}

// ExitExprGraphMatchOne is called when production exprGraphMatchOne is exited.
func (s *BasePartiQLParserListener) ExitExprGraphMatchOne(ctx *ExprGraphMatchOneContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasePartiQLParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePartiQLParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterVarRefExpr is called when production varRefExpr is entered.
func (s *BasePartiQLParserListener) EnterVarRefExpr(ctx *VarRefExprContext) {}

// ExitVarRefExpr is called when production varRefExpr is exited.
func (s *BasePartiQLParserListener) ExitVarRefExpr(ctx *VarRefExprContext) {}

// EnterCollection is called when production collection is entered.
func (s *BasePartiQLParserListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BasePartiQLParserListener) ExitCollection(ctx *CollectionContext) {}

// EnterArray is called when production array is entered.
func (s *BasePartiQLParserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BasePartiQLParserListener) ExitArray(ctx *ArrayContext) {}

// EnterBag is called when production bag is entered.
func (s *BasePartiQLParserListener) EnterBag(ctx *BagContext) {}

// ExitBag is called when production bag is exited.
func (s *BasePartiQLParserListener) ExitBag(ctx *BagContext) {}

// EnterTuple is called when production tuple is entered.
func (s *BasePartiQLParserListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production tuple is exited.
func (s *BasePartiQLParserListener) ExitTuple(ctx *TupleContext) {}

// EnterPair is called when production pair is entered.
func (s *BasePartiQLParserListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BasePartiQLParserListener) ExitPair(ctx *PairContext) {}

// EnterLiteralNull is called when production LiteralNull is entered.
func (s *BasePartiQLParserListener) EnterLiteralNull(ctx *LiteralNullContext) {}

// ExitLiteralNull is called when production LiteralNull is exited.
func (s *BasePartiQLParserListener) ExitLiteralNull(ctx *LiteralNullContext) {}

// EnterLiteralMissing is called when production LiteralMissing is entered.
func (s *BasePartiQLParserListener) EnterLiteralMissing(ctx *LiteralMissingContext) {}

// ExitLiteralMissing is called when production LiteralMissing is exited.
func (s *BasePartiQLParserListener) ExitLiteralMissing(ctx *LiteralMissingContext) {}

// EnterLiteralTrue is called when production LiteralTrue is entered.
func (s *BasePartiQLParserListener) EnterLiteralTrue(ctx *LiteralTrueContext) {}

// ExitLiteralTrue is called when production LiteralTrue is exited.
func (s *BasePartiQLParserListener) ExitLiteralTrue(ctx *LiteralTrueContext) {}

// EnterLiteralFalse is called when production LiteralFalse is entered.
func (s *BasePartiQLParserListener) EnterLiteralFalse(ctx *LiteralFalseContext) {}

// ExitLiteralFalse is called when production LiteralFalse is exited.
func (s *BasePartiQLParserListener) ExitLiteralFalse(ctx *LiteralFalseContext) {}

// EnterLiteralString is called when production LiteralString is entered.
func (s *BasePartiQLParserListener) EnterLiteralString(ctx *LiteralStringContext) {}

// ExitLiteralString is called when production LiteralString is exited.
func (s *BasePartiQLParserListener) ExitLiteralString(ctx *LiteralStringContext) {}

// EnterLiteralInteger is called when production LiteralInteger is entered.
func (s *BasePartiQLParserListener) EnterLiteralInteger(ctx *LiteralIntegerContext) {}

// ExitLiteralInteger is called when production LiteralInteger is exited.
func (s *BasePartiQLParserListener) ExitLiteralInteger(ctx *LiteralIntegerContext) {}

// EnterLiteralDecimal is called when production LiteralDecimal is entered.
func (s *BasePartiQLParserListener) EnterLiteralDecimal(ctx *LiteralDecimalContext) {}

// ExitLiteralDecimal is called when production LiteralDecimal is exited.
func (s *BasePartiQLParserListener) ExitLiteralDecimal(ctx *LiteralDecimalContext) {}

// EnterLiteralIon is called when production LiteralIon is entered.
func (s *BasePartiQLParserListener) EnterLiteralIon(ctx *LiteralIonContext) {}

// ExitLiteralIon is called when production LiteralIon is exited.
func (s *BasePartiQLParserListener) ExitLiteralIon(ctx *LiteralIonContext) {}

// EnterLiteralDate is called when production LiteralDate is entered.
func (s *BasePartiQLParserListener) EnterLiteralDate(ctx *LiteralDateContext) {}

// ExitLiteralDate is called when production LiteralDate is exited.
func (s *BasePartiQLParserListener) ExitLiteralDate(ctx *LiteralDateContext) {}

// EnterLiteralTime is called when production LiteralTime is entered.
func (s *BasePartiQLParserListener) EnterLiteralTime(ctx *LiteralTimeContext) {}

// ExitLiteralTime is called when production LiteralTime is exited.
func (s *BasePartiQLParserListener) ExitLiteralTime(ctx *LiteralTimeContext) {}

// EnterTypeAtomic is called when production TypeAtomic is entered.
func (s *BasePartiQLParserListener) EnterTypeAtomic(ctx *TypeAtomicContext) {}

// ExitTypeAtomic is called when production TypeAtomic is exited.
func (s *BasePartiQLParserListener) ExitTypeAtomic(ctx *TypeAtomicContext) {}

// EnterTypeArgSingle is called when production TypeArgSingle is entered.
func (s *BasePartiQLParserListener) EnterTypeArgSingle(ctx *TypeArgSingleContext) {}

// ExitTypeArgSingle is called when production TypeArgSingle is exited.
func (s *BasePartiQLParserListener) ExitTypeArgSingle(ctx *TypeArgSingleContext) {}

// EnterTypeVarChar is called when production TypeVarChar is entered.
func (s *BasePartiQLParserListener) EnterTypeVarChar(ctx *TypeVarCharContext) {}

// ExitTypeVarChar is called when production TypeVarChar is exited.
func (s *BasePartiQLParserListener) ExitTypeVarChar(ctx *TypeVarCharContext) {}

// EnterTypeArgDouble is called when production TypeArgDouble is entered.
func (s *BasePartiQLParserListener) EnterTypeArgDouble(ctx *TypeArgDoubleContext) {}

// ExitTypeArgDouble is called when production TypeArgDouble is exited.
func (s *BasePartiQLParserListener) ExitTypeArgDouble(ctx *TypeArgDoubleContext) {}

// EnterTypeTimeZone is called when production TypeTimeZone is entered.
func (s *BasePartiQLParserListener) EnterTypeTimeZone(ctx *TypeTimeZoneContext) {}

// ExitTypeTimeZone is called when production TypeTimeZone is exited.
func (s *BasePartiQLParserListener) ExitTypeTimeZone(ctx *TypeTimeZoneContext) {}

// EnterTypeCustom is called when production TypeCustom is entered.
func (s *BasePartiQLParserListener) EnterTypeCustom(ctx *TypeCustomContext) {}

// ExitTypeCustom is called when production TypeCustom is exited.
func (s *BasePartiQLParserListener) ExitTypeCustom(ctx *TypeCustomContext) {}
