// Code generated from TiDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tidb // TiDBParser
import "github.com/antlr4-go/antlr/v4"

// BaseTiDBParserListener is a complete listener for a parse tree produced by TiDBParser.
type BaseTiDBParserListener struct{}

var _ TiDBParserListener = &BaseTiDBParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTiDBParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTiDBParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTiDBParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTiDBParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseTiDBParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseTiDBParserListener) ExitScript(ctx *ScriptContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseTiDBParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseTiDBParserListener) ExitQuery(ctx *QueryContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseTiDBParserListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseTiDBParserListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterCreateStatement is called when production createStatement is entered.
func (s *BaseTiDBParserListener) EnterCreateStatement(ctx *CreateStatementContext) {}

// ExitCreateStatement is called when production createStatement is exited.
func (s *BaseTiDBParserListener) ExitCreateStatement(ctx *CreateStatementContext) {}

// EnterDropStatement is called when production dropStatement is entered.
func (s *BaseTiDBParserListener) EnterDropStatement(ctx *DropStatementContext) {}

// ExitDropStatement is called when production dropStatement is exited.
func (s *BaseTiDBParserListener) ExitDropStatement(ctx *DropStatementContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseTiDBParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseTiDBParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseTiDBParserListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseTiDBParserListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseTiDBParserListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseTiDBParserListener) ExitDropView(ctx *DropViewContext) {}

// EnterViewReplaceOrAlgorithm is called when production viewReplaceOrAlgorithm is entered.
func (s *BaseTiDBParserListener) EnterViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) {}

// ExitViewReplaceOrAlgorithm is called when production viewReplaceOrAlgorithm is exited.
func (s *BaseTiDBParserListener) ExitViewReplaceOrAlgorithm(ctx *ViewReplaceOrAlgorithmContext) {}

// EnterViewAlgorithm is called when production viewAlgorithm is entered.
func (s *BaseTiDBParserListener) EnterViewAlgorithm(ctx *ViewAlgorithmContext) {}

// ExitViewAlgorithm is called when production viewAlgorithm is exited.
func (s *BaseTiDBParserListener) ExitViewAlgorithm(ctx *ViewAlgorithmContext) {}

// EnterViewSuid is called when production viewSuid is entered.
func (s *BaseTiDBParserListener) EnterViewSuid(ctx *ViewSuidContext) {}

// ExitViewSuid is called when production viewSuid is exited.
func (s *BaseTiDBParserListener) ExitViewSuid(ctx *ViewSuidContext) {}

// EnterViewTail is called when production viewTail is entered.
func (s *BaseTiDBParserListener) EnterViewTail(ctx *ViewTailContext) {}

// ExitViewTail is called when production viewTail is exited.
func (s *BaseTiDBParserListener) ExitViewTail(ctx *ViewTailContext) {}

// EnterViewSelect is called when production viewSelect is entered.
func (s *BaseTiDBParserListener) EnterViewSelect(ctx *ViewSelectContext) {}

// ExitViewSelect is called when production viewSelect is exited.
func (s *BaseTiDBParserListener) ExitViewSelect(ctx *ViewSelectContext) {}

// EnterViewCheckOption is called when production viewCheckOption is entered.
func (s *BaseTiDBParserListener) EnterViewCheckOption(ctx *ViewCheckOptionContext) {}

// ExitViewCheckOption is called when production viewCheckOption is exited.
func (s *BaseTiDBParserListener) ExitViewCheckOption(ctx *ViewCheckOptionContext) {}

// EnterDuplicateAsQueryExpression is called when production duplicateAsQueryExpression is entered.
func (s *BaseTiDBParserListener) EnterDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) {
}

// ExitDuplicateAsQueryExpression is called when production duplicateAsQueryExpression is exited.
func (s *BaseTiDBParserListener) ExitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) {
}

// EnterQueryExpressionOrParens is called when production queryExpressionOrParens is entered.
func (s *BaseTiDBParserListener) EnterQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) {}

// ExitQueryExpressionOrParens is called when production queryExpressionOrParens is exited.
func (s *BaseTiDBParserListener) ExitQueryExpressionOrParens(ctx *QueryExpressionOrParensContext) {}

// EnterTableElementList is called when production tableElementList is entered.
func (s *BaseTiDBParserListener) EnterTableElementList(ctx *TableElementListContext) {}

// ExitTableElementList is called when production tableElementList is exited.
func (s *BaseTiDBParserListener) ExitTableElementList(ctx *TableElementListContext) {}

// EnterTableElement is called when production tableElement is entered.
func (s *BaseTiDBParserListener) EnterTableElement(ctx *TableElementContext) {}

// ExitTableElement is called when production tableElement is exited.
func (s *BaseTiDBParserListener) ExitTableElement(ctx *TableElementContext) {}

// EnterTableConstraintDef is called when production tableConstraintDef is entered.
func (s *BaseTiDBParserListener) EnterTableConstraintDef(ctx *TableConstraintDefContext) {}

// ExitTableConstraintDef is called when production tableConstraintDef is exited.
func (s *BaseTiDBParserListener) ExitTableConstraintDef(ctx *TableConstraintDefContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseTiDBParserListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseTiDBParserListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseTiDBParserListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseTiDBParserListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterIndexNameAndType is called when production indexNameAndType is entered.
func (s *BaseTiDBParserListener) EnterIndexNameAndType(ctx *IndexNameAndTypeContext) {}

// ExitIndexNameAndType is called when production indexNameAndType is exited.
func (s *BaseTiDBParserListener) ExitIndexNameAndType(ctx *IndexNameAndTypeContext) {}

// EnterTemporaryOption is called when production temporaryOption is entered.
func (s *BaseTiDBParserListener) EnterTemporaryOption(ctx *TemporaryOptionContext) {}

// ExitTemporaryOption is called when production temporaryOption is exited.
func (s *BaseTiDBParserListener) ExitTemporaryOption(ctx *TemporaryOptionContext) {}

// EnterColumnDef is called when production columnDef is entered.
func (s *BaseTiDBParserListener) EnterColumnDef(ctx *ColumnDefContext) {}

// ExitColumnDef is called when production columnDef is exited.
func (s *BaseTiDBParserListener) ExitColumnDef(ctx *ColumnDefContext) {}

// EnterColumnOptionList is called when production columnOptionList is entered.
func (s *BaseTiDBParserListener) EnterColumnOptionList(ctx *ColumnOptionListContext) {}

// ExitColumnOptionList is called when production columnOptionList is exited.
func (s *BaseTiDBParserListener) ExitColumnOptionList(ctx *ColumnOptionListContext) {}

// EnterColumnOption is called when production columnOption is entered.
func (s *BaseTiDBParserListener) EnterColumnOption(ctx *ColumnOptionContext) {}

// ExitColumnOption is called when production columnOption is exited.
func (s *BaseTiDBParserListener) ExitColumnOption(ctx *ColumnOptionContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseTiDBParserListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseTiDBParserListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterConstraintEnforcement is called when production constraintEnforcement is entered.
func (s *BaseTiDBParserListener) EnterConstraintEnforcement(ctx *ConstraintEnforcementContext) {}

// ExitConstraintEnforcement is called when production constraintEnforcement is exited.
func (s *BaseTiDBParserListener) ExitConstraintEnforcement(ctx *ConstraintEnforcementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseTiDBParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseTiDBParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectStatementWithInto is called when production selectStatementWithInto is entered.
func (s *BaseTiDBParserListener) EnterSelectStatementWithInto(ctx *SelectStatementWithIntoContext) {}

// ExitSelectStatementWithInto is called when production selectStatementWithInto is exited.
func (s *BaseTiDBParserListener) ExitSelectStatementWithInto(ctx *SelectStatementWithIntoContext) {}

// EnterQueryExpression is called when production queryExpression is entered.
func (s *BaseTiDBParserListener) EnterQueryExpression(ctx *QueryExpressionContext) {}

// ExitQueryExpression is called when production queryExpression is exited.
func (s *BaseTiDBParserListener) ExitQueryExpression(ctx *QueryExpressionContext) {}

// EnterQueryExpressionBody is called when production queryExpressionBody is entered.
func (s *BaseTiDBParserListener) EnterQueryExpressionBody(ctx *QueryExpressionBodyContext) {}

// ExitQueryExpressionBody is called when production queryExpressionBody is exited.
func (s *BaseTiDBParserListener) ExitQueryExpressionBody(ctx *QueryExpressionBodyContext) {}

// EnterQueryExpressionParens is called when production queryExpressionParens is entered.
func (s *BaseTiDBParserListener) EnterQueryExpressionParens(ctx *QueryExpressionParensContext) {}

// ExitQueryExpressionParens is called when production queryExpressionParens is exited.
func (s *BaseTiDBParserListener) ExitQueryExpressionParens(ctx *QueryExpressionParensContext) {}

// EnterQueryPrimary is called when production queryPrimary is entered.
func (s *BaseTiDBParserListener) EnterQueryPrimary(ctx *QueryPrimaryContext) {}

// ExitQueryPrimary is called when production queryPrimary is exited.
func (s *BaseTiDBParserListener) ExitQueryPrimary(ctx *QueryPrimaryContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseTiDBParserListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseTiDBParserListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseTiDBParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseTiDBParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterQuerySpecOption is called when production querySpecOption is entered.
func (s *BaseTiDBParserListener) EnterQuerySpecOption(ctx *QuerySpecOptionContext) {}

// ExitQuerySpecOption is called when production querySpecOption is exited.
func (s *BaseTiDBParserListener) ExitQuerySpecOption(ctx *QuerySpecOptionContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseTiDBParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseTiDBParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterSimpleLimitClause is called when production simpleLimitClause is entered.
func (s *BaseTiDBParserListener) EnterSimpleLimitClause(ctx *SimpleLimitClauseContext) {}

// ExitSimpleLimitClause is called when production simpleLimitClause is exited.
func (s *BaseTiDBParserListener) ExitSimpleLimitClause(ctx *SimpleLimitClauseContext) {}

// EnterLimitOptions is called when production limitOptions is entered.
func (s *BaseTiDBParserListener) EnterLimitOptions(ctx *LimitOptionsContext) {}

// ExitLimitOptions is called when production limitOptions is exited.
func (s *BaseTiDBParserListener) ExitLimitOptions(ctx *LimitOptionsContext) {}

// EnterLimitOption is called when production limitOption is entered.
func (s *BaseTiDBParserListener) EnterLimitOption(ctx *LimitOptionContext) {}

// ExitLimitOption is called when production limitOption is exited.
func (s *BaseTiDBParserListener) ExitLimitOption(ctx *LimitOptionContext) {}

// EnterIntoClause is called when production intoClause is entered.
func (s *BaseTiDBParserListener) EnterIntoClause(ctx *IntoClauseContext) {}

// ExitIntoClause is called when production intoClause is exited.
func (s *BaseTiDBParserListener) ExitIntoClause(ctx *IntoClauseContext) {}

// EnterProcedureAnalyseClause is called when production procedureAnalyseClause is entered.
func (s *BaseTiDBParserListener) EnterProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) {}

// ExitProcedureAnalyseClause is called when production procedureAnalyseClause is exited.
func (s *BaseTiDBParserListener) ExitProcedureAnalyseClause(ctx *ProcedureAnalyseClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseTiDBParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseTiDBParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseTiDBParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseTiDBParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterWindowDefinition is called when production windowDefinition is entered.
func (s *BaseTiDBParserListener) EnterWindowDefinition(ctx *WindowDefinitionContext) {}

// ExitWindowDefinition is called when production windowDefinition is exited.
func (s *BaseTiDBParserListener) ExitWindowDefinition(ctx *WindowDefinitionContext) {}

// EnterWindowSpec is called when production windowSpec is entered.
func (s *BaseTiDBParserListener) EnterWindowSpec(ctx *WindowSpecContext) {}

// ExitWindowSpec is called when production windowSpec is exited.
func (s *BaseTiDBParserListener) ExitWindowSpec(ctx *WindowSpecContext) {}

// EnterWindowSpecDetails is called when production windowSpecDetails is entered.
func (s *BaseTiDBParserListener) EnterWindowSpecDetails(ctx *WindowSpecDetailsContext) {}

// ExitWindowSpecDetails is called when production windowSpecDetails is exited.
func (s *BaseTiDBParserListener) ExitWindowSpecDetails(ctx *WindowSpecDetailsContext) {}

// EnterWindowFrameClause is called when production windowFrameClause is entered.
func (s *BaseTiDBParserListener) EnterWindowFrameClause(ctx *WindowFrameClauseContext) {}

// ExitWindowFrameClause is called when production windowFrameClause is exited.
func (s *BaseTiDBParserListener) ExitWindowFrameClause(ctx *WindowFrameClauseContext) {}

// EnterWindowFrameUnits is called when production windowFrameUnits is entered.
func (s *BaseTiDBParserListener) EnterWindowFrameUnits(ctx *WindowFrameUnitsContext) {}

// ExitWindowFrameUnits is called when production windowFrameUnits is exited.
func (s *BaseTiDBParserListener) ExitWindowFrameUnits(ctx *WindowFrameUnitsContext) {}

// EnterWindowFrameExtent is called when production windowFrameExtent is entered.
func (s *BaseTiDBParserListener) EnterWindowFrameExtent(ctx *WindowFrameExtentContext) {}

// ExitWindowFrameExtent is called when production windowFrameExtent is exited.
func (s *BaseTiDBParserListener) ExitWindowFrameExtent(ctx *WindowFrameExtentContext) {}

// EnterWindowFrameStart is called when production windowFrameStart is entered.
func (s *BaseTiDBParserListener) EnterWindowFrameStart(ctx *WindowFrameStartContext) {}

// ExitWindowFrameStart is called when production windowFrameStart is exited.
func (s *BaseTiDBParserListener) ExitWindowFrameStart(ctx *WindowFrameStartContext) {}

// EnterWindowFrameBetween is called when production windowFrameBetween is entered.
func (s *BaseTiDBParserListener) EnterWindowFrameBetween(ctx *WindowFrameBetweenContext) {}

// ExitWindowFrameBetween is called when production windowFrameBetween is exited.
func (s *BaseTiDBParserListener) ExitWindowFrameBetween(ctx *WindowFrameBetweenContext) {}

// EnterWindowFrameBound is called when production windowFrameBound is entered.
func (s *BaseTiDBParserListener) EnterWindowFrameBound(ctx *WindowFrameBoundContext) {}

// ExitWindowFrameBound is called when production windowFrameBound is exited.
func (s *BaseTiDBParserListener) ExitWindowFrameBound(ctx *WindowFrameBoundContext) {}

// EnterWindowFrameExclusion is called when production windowFrameExclusion is entered.
func (s *BaseTiDBParserListener) EnterWindowFrameExclusion(ctx *WindowFrameExclusionContext) {}

// ExitWindowFrameExclusion is called when production windowFrameExclusion is exited.
func (s *BaseTiDBParserListener) ExitWindowFrameExclusion(ctx *WindowFrameExclusionContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseTiDBParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseTiDBParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterCommonTableExpression is called when production commonTableExpression is entered.
func (s *BaseTiDBParserListener) EnterCommonTableExpression(ctx *CommonTableExpressionContext) {}

// ExitCommonTableExpression is called when production commonTableExpression is exited.
func (s *BaseTiDBParserListener) ExitCommonTableExpression(ctx *CommonTableExpressionContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseTiDBParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseTiDBParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterOlapOption is called when production olapOption is entered.
func (s *BaseTiDBParserListener) EnterOlapOption(ctx *OlapOptionContext) {}

// ExitOlapOption is called when production olapOption is exited.
func (s *BaseTiDBParserListener) ExitOlapOption(ctx *OlapOptionContext) {}

// EnterOrderClause is called when production orderClause is entered.
func (s *BaseTiDBParserListener) EnterOrderClause(ctx *OrderClauseContext) {}

// ExitOrderClause is called when production orderClause is exited.
func (s *BaseTiDBParserListener) ExitOrderClause(ctx *OrderClauseContext) {}

// EnterDirection is called when production direction is entered.
func (s *BaseTiDBParserListener) EnterDirection(ctx *DirectionContext) {}

// ExitDirection is called when production direction is exited.
func (s *BaseTiDBParserListener) ExitDirection(ctx *DirectionContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseTiDBParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseTiDBParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableReferenceList is called when production tableReferenceList is entered.
func (s *BaseTiDBParserListener) EnterTableReferenceList(ctx *TableReferenceListContext) {}

// ExitTableReferenceList is called when production tableReferenceList is exited.
func (s *BaseTiDBParserListener) ExitTableReferenceList(ctx *TableReferenceListContext) {}

// EnterTableValueConstructor is called when production tableValueConstructor is entered.
func (s *BaseTiDBParserListener) EnterTableValueConstructor(ctx *TableValueConstructorContext) {}

// ExitTableValueConstructor is called when production tableValueConstructor is exited.
func (s *BaseTiDBParserListener) ExitTableValueConstructor(ctx *TableValueConstructorContext) {}

// EnterExplicitTable is called when production explicitTable is entered.
func (s *BaseTiDBParserListener) EnterExplicitTable(ctx *ExplicitTableContext) {}

// ExitExplicitTable is called when production explicitTable is exited.
func (s *BaseTiDBParserListener) ExitExplicitTable(ctx *ExplicitTableContext) {}

// EnterRowValueExplicit is called when production rowValueExplicit is entered.
func (s *BaseTiDBParserListener) EnterRowValueExplicit(ctx *RowValueExplicitContext) {}

// ExitRowValueExplicit is called when production rowValueExplicit is exited.
func (s *BaseTiDBParserListener) ExitRowValueExplicit(ctx *RowValueExplicitContext) {}

// EnterValues is called when production values is entered.
func (s *BaseTiDBParserListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseTiDBParserListener) ExitValues(ctx *ValuesContext) {}

// EnterSelectOption is called when production selectOption is entered.
func (s *BaseTiDBParserListener) EnterSelectOption(ctx *SelectOptionContext) {}

// ExitSelectOption is called when production selectOption is exited.
func (s *BaseTiDBParserListener) ExitSelectOption(ctx *SelectOptionContext) {}

// EnterLockingClauseList is called when production lockingClauseList is entered.
func (s *BaseTiDBParserListener) EnterLockingClauseList(ctx *LockingClauseListContext) {}

// ExitLockingClauseList is called when production lockingClauseList is exited.
func (s *BaseTiDBParserListener) ExitLockingClauseList(ctx *LockingClauseListContext) {}

// EnterLockingClause is called when production lockingClause is entered.
func (s *BaseTiDBParserListener) EnterLockingClause(ctx *LockingClauseContext) {}

// ExitLockingClause is called when production lockingClause is exited.
func (s *BaseTiDBParserListener) ExitLockingClause(ctx *LockingClauseContext) {}

// EnterLockStrengh is called when production lockStrengh is entered.
func (s *BaseTiDBParserListener) EnterLockStrengh(ctx *LockStrenghContext) {}

// ExitLockStrengh is called when production lockStrengh is exited.
func (s *BaseTiDBParserListener) ExitLockStrengh(ctx *LockStrenghContext) {}

// EnterLockedRowAction is called when production lockedRowAction is entered.
func (s *BaseTiDBParserListener) EnterLockedRowAction(ctx *LockedRowActionContext) {}

// ExitLockedRowAction is called when production lockedRowAction is exited.
func (s *BaseTiDBParserListener) ExitLockedRowAction(ctx *LockedRowActionContext) {}

// EnterSelectItemList is called when production selectItemList is entered.
func (s *BaseTiDBParserListener) EnterSelectItemList(ctx *SelectItemListContext) {}

// ExitSelectItemList is called when production selectItemList is exited.
func (s *BaseTiDBParserListener) ExitSelectItemList(ctx *SelectItemListContext) {}

// EnterSelectItem is called when production selectItem is entered.
func (s *BaseTiDBParserListener) EnterSelectItem(ctx *SelectItemContext) {}

// ExitSelectItem is called when production selectItem is exited.
func (s *BaseTiDBParserListener) ExitSelectItem(ctx *SelectItemContext) {}

// EnterSelectAlias is called when production selectAlias is entered.
func (s *BaseTiDBParserListener) EnterSelectAlias(ctx *SelectAliasContext) {}

// ExitSelectAlias is called when production selectAlias is exited.
func (s *BaseTiDBParserListener) ExitSelectAlias(ctx *SelectAliasContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseTiDBParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseTiDBParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterTableReference is called when production tableReference is entered.
func (s *BaseTiDBParserListener) EnterTableReference(ctx *TableReferenceContext) {}

// ExitTableReference is called when production tableReference is exited.
func (s *BaseTiDBParserListener) ExitTableReference(ctx *TableReferenceContext) {}

// EnterEscapedTableReference is called when production escapedTableReference is entered.
func (s *BaseTiDBParserListener) EnterEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// ExitEscapedTableReference is called when production escapedTableReference is exited.
func (s *BaseTiDBParserListener) ExitEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// EnterJoinedTable is called when production joinedTable is entered.
func (s *BaseTiDBParserListener) EnterJoinedTable(ctx *JoinedTableContext) {}

// ExitJoinedTable is called when production joinedTable is exited.
func (s *BaseTiDBParserListener) ExitJoinedTable(ctx *JoinedTableContext) {}

// EnterNaturalJoinType is called when production naturalJoinType is entered.
func (s *BaseTiDBParserListener) EnterNaturalJoinType(ctx *NaturalJoinTypeContext) {}

// ExitNaturalJoinType is called when production naturalJoinType is exited.
func (s *BaseTiDBParserListener) ExitNaturalJoinType(ctx *NaturalJoinTypeContext) {}

// EnterInnerJoinType is called when production innerJoinType is entered.
func (s *BaseTiDBParserListener) EnterInnerJoinType(ctx *InnerJoinTypeContext) {}

// ExitInnerJoinType is called when production innerJoinType is exited.
func (s *BaseTiDBParserListener) ExitInnerJoinType(ctx *InnerJoinTypeContext) {}

// EnterOuterJoinType is called when production outerJoinType is entered.
func (s *BaseTiDBParserListener) EnterOuterJoinType(ctx *OuterJoinTypeContext) {}

// ExitOuterJoinType is called when production outerJoinType is exited.
func (s *BaseTiDBParserListener) ExitOuterJoinType(ctx *OuterJoinTypeContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (s *BaseTiDBParserListener) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (s *BaseTiDBParserListener) ExitTableFactor(ctx *TableFactorContext) {}

// EnterSingleTable is called when production singleTable is entered.
func (s *BaseTiDBParserListener) EnterSingleTable(ctx *SingleTableContext) {}

// ExitSingleTable is called when production singleTable is exited.
func (s *BaseTiDBParserListener) ExitSingleTable(ctx *SingleTableContext) {}

// EnterSingleTableParens is called when production singleTableParens is entered.
func (s *BaseTiDBParserListener) EnterSingleTableParens(ctx *SingleTableParensContext) {}

// ExitSingleTableParens is called when production singleTableParens is exited.
func (s *BaseTiDBParserListener) ExitSingleTableParens(ctx *SingleTableParensContext) {}

// EnterDerivedTable is called when production derivedTable is entered.
func (s *BaseTiDBParserListener) EnterDerivedTable(ctx *DerivedTableContext) {}

// ExitDerivedTable is called when production derivedTable is exited.
func (s *BaseTiDBParserListener) ExitDerivedTable(ctx *DerivedTableContext) {}

// EnterTableReferenceListParens is called when production tableReferenceListParens is entered.
func (s *BaseTiDBParserListener) EnterTableReferenceListParens(ctx *TableReferenceListParensContext) {
}

// ExitTableReferenceListParens is called when production tableReferenceListParens is exited.
func (s *BaseTiDBParserListener) ExitTableReferenceListParens(ctx *TableReferenceListParensContext) {}

// EnterTableFunction is called when production tableFunction is entered.
func (s *BaseTiDBParserListener) EnterTableFunction(ctx *TableFunctionContext) {}

// ExitTableFunction is called when production tableFunction is exited.
func (s *BaseTiDBParserListener) ExitTableFunction(ctx *TableFunctionContext) {}

// EnterColumnsClause is called when production columnsClause is entered.
func (s *BaseTiDBParserListener) EnterColumnsClause(ctx *ColumnsClauseContext) {}

// ExitColumnsClause is called when production columnsClause is exited.
func (s *BaseTiDBParserListener) ExitColumnsClause(ctx *ColumnsClauseContext) {}

// EnterJtColumn is called when production jtColumn is entered.
func (s *BaseTiDBParserListener) EnterJtColumn(ctx *JtColumnContext) {}

// ExitJtColumn is called when production jtColumn is exited.
func (s *BaseTiDBParserListener) ExitJtColumn(ctx *JtColumnContext) {}

// EnterOnEmptyOrError is called when production onEmptyOrError is entered.
func (s *BaseTiDBParserListener) EnterOnEmptyOrError(ctx *OnEmptyOrErrorContext) {}

// ExitOnEmptyOrError is called when production onEmptyOrError is exited.
func (s *BaseTiDBParserListener) ExitOnEmptyOrError(ctx *OnEmptyOrErrorContext) {}

// EnterOnEmpty is called when production onEmpty is entered.
func (s *BaseTiDBParserListener) EnterOnEmpty(ctx *OnEmptyContext) {}

// ExitOnEmpty is called when production onEmpty is exited.
func (s *BaseTiDBParserListener) ExitOnEmpty(ctx *OnEmptyContext) {}

// EnterOnError is called when production onError is entered.
func (s *BaseTiDBParserListener) EnterOnError(ctx *OnErrorContext) {}

// ExitOnError is called when production onError is exited.
func (s *BaseTiDBParserListener) ExitOnError(ctx *OnErrorContext) {}

// EnterJtOnResponse is called when production jtOnResponse is entered.
func (s *BaseTiDBParserListener) EnterJtOnResponse(ctx *JtOnResponseContext) {}

// ExitJtOnResponse is called when production jtOnResponse is exited.
func (s *BaseTiDBParserListener) ExitJtOnResponse(ctx *JtOnResponseContext) {}

// EnterSetOprSymbol is called when production setOprSymbol is entered.
func (s *BaseTiDBParserListener) EnterSetOprSymbol(ctx *SetOprSymbolContext) {}

// ExitSetOprSymbol is called when production setOprSymbol is exited.
func (s *BaseTiDBParserListener) ExitSetOprSymbol(ctx *SetOprSymbolContext) {}

// EnterSetOprOption is called when production setOprOption is entered.
func (s *BaseTiDBParserListener) EnterSetOprOption(ctx *SetOprOptionContext) {}

// ExitSetOprOption is called when production setOprOption is exited.
func (s *BaseTiDBParserListener) ExitSetOprOption(ctx *SetOprOptionContext) {}

// EnterTableAlias is called when production tableAlias is entered.
func (s *BaseTiDBParserListener) EnterTableAlias(ctx *TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *BaseTiDBParserListener) ExitTableAlias(ctx *TableAliasContext) {}

// EnterIndexHintList is called when production indexHintList is entered.
func (s *BaseTiDBParserListener) EnterIndexHintList(ctx *IndexHintListContext) {}

// ExitIndexHintList is called when production indexHintList is exited.
func (s *BaseTiDBParserListener) ExitIndexHintList(ctx *IndexHintListContext) {}

// EnterIndexHint is called when production indexHint is entered.
func (s *BaseTiDBParserListener) EnterIndexHint(ctx *IndexHintContext) {}

// ExitIndexHint is called when production indexHint is exited.
func (s *BaseTiDBParserListener) ExitIndexHint(ctx *IndexHintContext) {}

// EnterIndexHintType is called when production indexHintType is entered.
func (s *BaseTiDBParserListener) EnterIndexHintType(ctx *IndexHintTypeContext) {}

// ExitIndexHintType is called when production indexHintType is exited.
func (s *BaseTiDBParserListener) ExitIndexHintType(ctx *IndexHintTypeContext) {}

// EnterKeyOrIndex is called when production keyOrIndex is entered.
func (s *BaseTiDBParserListener) EnterKeyOrIndex(ctx *KeyOrIndexContext) {}

// ExitKeyOrIndex is called when production keyOrIndex is exited.
func (s *BaseTiDBParserListener) ExitKeyOrIndex(ctx *KeyOrIndexContext) {}

// EnterConstraintKeyType is called when production constraintKeyType is entered.
func (s *BaseTiDBParserListener) EnterConstraintKeyType(ctx *ConstraintKeyTypeContext) {}

// ExitConstraintKeyType is called when production constraintKeyType is exited.
func (s *BaseTiDBParserListener) ExitConstraintKeyType(ctx *ConstraintKeyTypeContext) {}

// EnterIndexHintClause is called when production indexHintClause is entered.
func (s *BaseTiDBParserListener) EnterIndexHintClause(ctx *IndexHintClauseContext) {}

// ExitIndexHintClause is called when production indexHintClause is exited.
func (s *BaseTiDBParserListener) ExitIndexHintClause(ctx *IndexHintClauseContext) {}

// EnterIndexList is called when production indexList is entered.
func (s *BaseTiDBParserListener) EnterIndexList(ctx *IndexListContext) {}

// ExitIndexList is called when production indexList is exited.
func (s *BaseTiDBParserListener) ExitIndexList(ctx *IndexListContext) {}

// EnterIndexListElement is called when production indexListElement is entered.
func (s *BaseTiDBParserListener) EnterIndexListElement(ctx *IndexListElementContext) {}

// ExitIndexListElement is called when production indexListElement is exited.
func (s *BaseTiDBParserListener) ExitIndexListElement(ctx *IndexListElementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseTiDBParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseTiDBParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterTransactionOrLockingStatement is called when production transactionOrLockingStatement is entered.
func (s *BaseTiDBParserListener) EnterTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) {
}

// ExitTransactionOrLockingStatement is called when production transactionOrLockingStatement is exited.
func (s *BaseTiDBParserListener) ExitTransactionOrLockingStatement(ctx *TransactionOrLockingStatementContext) {
}

// EnterTransactionStatement is called when production transactionStatement is entered.
func (s *BaseTiDBParserListener) EnterTransactionStatement(ctx *TransactionStatementContext) {}

// ExitTransactionStatement is called when production transactionStatement is exited.
func (s *BaseTiDBParserListener) ExitTransactionStatement(ctx *TransactionStatementContext) {}

// EnterBeginWork is called when production beginWork is entered.
func (s *BaseTiDBParserListener) EnterBeginWork(ctx *BeginWorkContext) {}

// ExitBeginWork is called when production beginWork is exited.
func (s *BaseTiDBParserListener) ExitBeginWork(ctx *BeginWorkContext) {}

// EnterTransactionCharacteristic is called when production transactionCharacteristic is entered.
func (s *BaseTiDBParserListener) EnterTransactionCharacteristic(ctx *TransactionCharacteristicContext) {
}

// ExitTransactionCharacteristic is called when production transactionCharacteristic is exited.
func (s *BaseTiDBParserListener) ExitTransactionCharacteristic(ctx *TransactionCharacteristicContext) {
}

// EnterSavepointStatement is called when production savepointStatement is entered.
func (s *BaseTiDBParserListener) EnterSavepointStatement(ctx *SavepointStatementContext) {}

// ExitSavepointStatement is called when production savepointStatement is exited.
func (s *BaseTiDBParserListener) ExitSavepointStatement(ctx *SavepointStatementContext) {}

// EnterLockStatement is called when production lockStatement is entered.
func (s *BaseTiDBParserListener) EnterLockStatement(ctx *LockStatementContext) {}

// ExitLockStatement is called when production lockStatement is exited.
func (s *BaseTiDBParserListener) ExitLockStatement(ctx *LockStatementContext) {}

// EnterLockItem is called when production lockItem is entered.
func (s *BaseTiDBParserListener) EnterLockItem(ctx *LockItemContext) {}

// ExitLockItem is called when production lockItem is exited.
func (s *BaseTiDBParserListener) ExitLockItem(ctx *LockItemContext) {}

// EnterLockOption is called when production lockOption is entered.
func (s *BaseTiDBParserListener) EnterLockOption(ctx *LockOptionContext) {}

// ExitLockOption is called when production lockOption is exited.
func (s *BaseTiDBParserListener) ExitLockOption(ctx *LockOptionContext) {}

// EnterXaStatement is called when production xaStatement is entered.
func (s *BaseTiDBParserListener) EnterXaStatement(ctx *XaStatementContext) {}

// ExitXaStatement is called when production xaStatement is exited.
func (s *BaseTiDBParserListener) ExitXaStatement(ctx *XaStatementContext) {}

// EnterXaConvert is called when production xaConvert is entered.
func (s *BaseTiDBParserListener) EnterXaConvert(ctx *XaConvertContext) {}

// ExitXaConvert is called when production xaConvert is exited.
func (s *BaseTiDBParserListener) ExitXaConvert(ctx *XaConvertContext) {}

// EnterXid is called when production xid is entered.
func (s *BaseTiDBParserListener) EnterXid(ctx *XidContext) {}

// ExitXid is called when production xid is exited.
func (s *BaseTiDBParserListener) ExitXid(ctx *XidContext) {}

// EnterReplicationStatement is called when production replicationStatement is entered.
func (s *BaseTiDBParserListener) EnterReplicationStatement(ctx *ReplicationStatementContext) {}

// ExitReplicationStatement is called when production replicationStatement is exited.
func (s *BaseTiDBParserListener) ExitReplicationStatement(ctx *ReplicationStatementContext) {}

// EnterResetOption is called when production resetOption is entered.
func (s *BaseTiDBParserListener) EnterResetOption(ctx *ResetOptionContext) {}

// ExitResetOption is called when production resetOption is exited.
func (s *BaseTiDBParserListener) ExitResetOption(ctx *ResetOptionContext) {}

// EnterMasterResetOptions is called when production masterResetOptions is entered.
func (s *BaseTiDBParserListener) EnterMasterResetOptions(ctx *MasterResetOptionsContext) {}

// ExitMasterResetOptions is called when production masterResetOptions is exited.
func (s *BaseTiDBParserListener) ExitMasterResetOptions(ctx *MasterResetOptionsContext) {}

// EnterReplicationLoad is called when production replicationLoad is entered.
func (s *BaseTiDBParserListener) EnterReplicationLoad(ctx *ReplicationLoadContext) {}

// ExitReplicationLoad is called when production replicationLoad is exited.
func (s *BaseTiDBParserListener) ExitReplicationLoad(ctx *ReplicationLoadContext) {}

// EnterChangeMaster is called when production changeMaster is entered.
func (s *BaseTiDBParserListener) EnterChangeMaster(ctx *ChangeMasterContext) {}

// ExitChangeMaster is called when production changeMaster is exited.
func (s *BaseTiDBParserListener) ExitChangeMaster(ctx *ChangeMasterContext) {}

// EnterChangeMasterOptions is called when production changeMasterOptions is entered.
func (s *BaseTiDBParserListener) EnterChangeMasterOptions(ctx *ChangeMasterOptionsContext) {}

// ExitChangeMasterOptions is called when production changeMasterOptions is exited.
func (s *BaseTiDBParserListener) ExitChangeMasterOptions(ctx *ChangeMasterOptionsContext) {}

// EnterMasterOption is called when production masterOption is entered.
func (s *BaseTiDBParserListener) EnterMasterOption(ctx *MasterOptionContext) {}

// ExitMasterOption is called when production masterOption is exited.
func (s *BaseTiDBParserListener) ExitMasterOption(ctx *MasterOptionContext) {}

// EnterPrivilegeCheckDef is called when production privilegeCheckDef is entered.
func (s *BaseTiDBParserListener) EnterPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) {}

// ExitPrivilegeCheckDef is called when production privilegeCheckDef is exited.
func (s *BaseTiDBParserListener) ExitPrivilegeCheckDef(ctx *PrivilegeCheckDefContext) {}

// EnterTablePrimaryKeyCheckDef is called when production tablePrimaryKeyCheckDef is entered.
func (s *BaseTiDBParserListener) EnterTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) {}

// ExitTablePrimaryKeyCheckDef is called when production tablePrimaryKeyCheckDef is exited.
func (s *BaseTiDBParserListener) ExitTablePrimaryKeyCheckDef(ctx *TablePrimaryKeyCheckDefContext) {}

// EnterMasterTlsCiphersuitesDef is called when production masterTlsCiphersuitesDef is entered.
func (s *BaseTiDBParserListener) EnterMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) {
}

// ExitMasterTlsCiphersuitesDef is called when production masterTlsCiphersuitesDef is exited.
func (s *BaseTiDBParserListener) ExitMasterTlsCiphersuitesDef(ctx *MasterTlsCiphersuitesDefContext) {}

// EnterMasterFileDef is called when production masterFileDef is entered.
func (s *BaseTiDBParserListener) EnterMasterFileDef(ctx *MasterFileDefContext) {}

// ExitMasterFileDef is called when production masterFileDef is exited.
func (s *BaseTiDBParserListener) ExitMasterFileDef(ctx *MasterFileDefContext) {}

// EnterServerIdList is called when production serverIdList is entered.
func (s *BaseTiDBParserListener) EnterServerIdList(ctx *ServerIdListContext) {}

// ExitServerIdList is called when production serverIdList is exited.
func (s *BaseTiDBParserListener) ExitServerIdList(ctx *ServerIdListContext) {}

// EnterChangeReplication is called when production changeReplication is entered.
func (s *BaseTiDBParserListener) EnterChangeReplication(ctx *ChangeReplicationContext) {}

// ExitChangeReplication is called when production changeReplication is exited.
func (s *BaseTiDBParserListener) ExitChangeReplication(ctx *ChangeReplicationContext) {}

// EnterFilterDefinition is called when production filterDefinition is entered.
func (s *BaseTiDBParserListener) EnterFilterDefinition(ctx *FilterDefinitionContext) {}

// ExitFilterDefinition is called when production filterDefinition is exited.
func (s *BaseTiDBParserListener) ExitFilterDefinition(ctx *FilterDefinitionContext) {}

// EnterFilterDbList is called when production filterDbList is entered.
func (s *BaseTiDBParserListener) EnterFilterDbList(ctx *FilterDbListContext) {}

// ExitFilterDbList is called when production filterDbList is exited.
func (s *BaseTiDBParserListener) ExitFilterDbList(ctx *FilterDbListContext) {}

// EnterFilterTableList is called when production filterTableList is entered.
func (s *BaseTiDBParserListener) EnterFilterTableList(ctx *FilterTableListContext) {}

// ExitFilterTableList is called when production filterTableList is exited.
func (s *BaseTiDBParserListener) ExitFilterTableList(ctx *FilterTableListContext) {}

// EnterFilterStringList is called when production filterStringList is entered.
func (s *BaseTiDBParserListener) EnterFilterStringList(ctx *FilterStringListContext) {}

// ExitFilterStringList is called when production filterStringList is exited.
func (s *BaseTiDBParserListener) ExitFilterStringList(ctx *FilterStringListContext) {}

// EnterFilterWildDbTableString is called when production filterWildDbTableString is entered.
func (s *BaseTiDBParserListener) EnterFilterWildDbTableString(ctx *FilterWildDbTableStringContext) {}

// ExitFilterWildDbTableString is called when production filterWildDbTableString is exited.
func (s *BaseTiDBParserListener) ExitFilterWildDbTableString(ctx *FilterWildDbTableStringContext) {}

// EnterFilterDbPairList is called when production filterDbPairList is entered.
func (s *BaseTiDBParserListener) EnterFilterDbPairList(ctx *FilterDbPairListContext) {}

// ExitFilterDbPairList is called when production filterDbPairList is exited.
func (s *BaseTiDBParserListener) ExitFilterDbPairList(ctx *FilterDbPairListContext) {}

// EnterSlave is called when production slave is entered.
func (s *BaseTiDBParserListener) EnterSlave(ctx *SlaveContext) {}

// ExitSlave is called when production slave is exited.
func (s *BaseTiDBParserListener) ExitSlave(ctx *SlaveContext) {}

// EnterSlaveUntilOptions is called when production slaveUntilOptions is entered.
func (s *BaseTiDBParserListener) EnterSlaveUntilOptions(ctx *SlaveUntilOptionsContext) {}

// ExitSlaveUntilOptions is called when production slaveUntilOptions is exited.
func (s *BaseTiDBParserListener) ExitSlaveUntilOptions(ctx *SlaveUntilOptionsContext) {}

// EnterSlaveConnectionOptions is called when production slaveConnectionOptions is entered.
func (s *BaseTiDBParserListener) EnterSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) {}

// ExitSlaveConnectionOptions is called when production slaveConnectionOptions is exited.
func (s *BaseTiDBParserListener) ExitSlaveConnectionOptions(ctx *SlaveConnectionOptionsContext) {}

// EnterSlaveThreadOptions is called when production slaveThreadOptions is entered.
func (s *BaseTiDBParserListener) EnterSlaveThreadOptions(ctx *SlaveThreadOptionsContext) {}

// ExitSlaveThreadOptions is called when production slaveThreadOptions is exited.
func (s *BaseTiDBParserListener) ExitSlaveThreadOptions(ctx *SlaveThreadOptionsContext) {}

// EnterSlaveThreadOption is called when production slaveThreadOption is entered.
func (s *BaseTiDBParserListener) EnterSlaveThreadOption(ctx *SlaveThreadOptionContext) {}

// ExitSlaveThreadOption is called when production slaveThreadOption is exited.
func (s *BaseTiDBParserListener) ExitSlaveThreadOption(ctx *SlaveThreadOptionContext) {}

// EnterGroupReplication is called when production groupReplication is entered.
func (s *BaseTiDBParserListener) EnterGroupReplication(ctx *GroupReplicationContext) {}

// ExitGroupReplication is called when production groupReplication is exited.
func (s *BaseTiDBParserListener) ExitGroupReplication(ctx *GroupReplicationContext) {}

// EnterPreparedStatement is called when production preparedStatement is entered.
func (s *BaseTiDBParserListener) EnterPreparedStatement(ctx *PreparedStatementContext) {}

// ExitPreparedStatement is called when production preparedStatement is exited.
func (s *BaseTiDBParserListener) ExitPreparedStatement(ctx *PreparedStatementContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseTiDBParserListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseTiDBParserListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterExecuteVarList is called when production executeVarList is entered.
func (s *BaseTiDBParserListener) EnterExecuteVarList(ctx *ExecuteVarListContext) {}

// ExitExecuteVarList is called when production executeVarList is exited.
func (s *BaseTiDBParserListener) ExitExecuteVarList(ctx *ExecuteVarListContext) {}

// EnterCloneStatement is called when production cloneStatement is entered.
func (s *BaseTiDBParserListener) EnterCloneStatement(ctx *CloneStatementContext) {}

// ExitCloneStatement is called when production cloneStatement is exited.
func (s *BaseTiDBParserListener) ExitCloneStatement(ctx *CloneStatementContext) {}

// EnterDataDirSSL is called when production dataDirSSL is entered.
func (s *BaseTiDBParserListener) EnterDataDirSSL(ctx *DataDirSSLContext) {}

// ExitDataDirSSL is called when production dataDirSSL is exited.
func (s *BaseTiDBParserListener) ExitDataDirSSL(ctx *DataDirSSLContext) {}

// EnterSsl is called when production ssl is entered.
func (s *BaseTiDBParserListener) EnterSsl(ctx *SslContext) {}

// ExitSsl is called when production ssl is exited.
func (s *BaseTiDBParserListener) ExitSsl(ctx *SslContext) {}

// EnterAccountManagementStatement is called when production accountManagementStatement is entered.
func (s *BaseTiDBParserListener) EnterAccountManagementStatement(ctx *AccountManagementStatementContext) {
}

// ExitAccountManagementStatement is called when production accountManagementStatement is exited.
func (s *BaseTiDBParserListener) ExitAccountManagementStatement(ctx *AccountManagementStatementContext) {
}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseTiDBParserListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseTiDBParserListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterAlterUserTail is called when production alterUserTail is entered.
func (s *BaseTiDBParserListener) EnterAlterUserTail(ctx *AlterUserTailContext) {}

// ExitAlterUserTail is called when production alterUserTail is exited.
func (s *BaseTiDBParserListener) ExitAlterUserTail(ctx *AlterUserTailContext) {}

// EnterUserFunction is called when production userFunction is entered.
func (s *BaseTiDBParserListener) EnterUserFunction(ctx *UserFunctionContext) {}

// ExitUserFunction is called when production userFunction is exited.
func (s *BaseTiDBParserListener) ExitUserFunction(ctx *UserFunctionContext) {}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseTiDBParserListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseTiDBParserListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterCreateUserTail is called when production createUserTail is entered.
func (s *BaseTiDBParserListener) EnterCreateUserTail(ctx *CreateUserTailContext) {}

// ExitCreateUserTail is called when production createUserTail is exited.
func (s *BaseTiDBParserListener) ExitCreateUserTail(ctx *CreateUserTailContext) {}

// EnterDefaultRoleClause is called when production defaultRoleClause is entered.
func (s *BaseTiDBParserListener) EnterDefaultRoleClause(ctx *DefaultRoleClauseContext) {}

// ExitDefaultRoleClause is called when production defaultRoleClause is exited.
func (s *BaseTiDBParserListener) ExitDefaultRoleClause(ctx *DefaultRoleClauseContext) {}

// EnterRequireClause is called when production requireClause is entered.
func (s *BaseTiDBParserListener) EnterRequireClause(ctx *RequireClauseContext) {}

// ExitRequireClause is called when production requireClause is exited.
func (s *BaseTiDBParserListener) ExitRequireClause(ctx *RequireClauseContext) {}

// EnterConnectOptions is called when production connectOptions is entered.
func (s *BaseTiDBParserListener) EnterConnectOptions(ctx *ConnectOptionsContext) {}

// ExitConnectOptions is called when production connectOptions is exited.
func (s *BaseTiDBParserListener) ExitConnectOptions(ctx *ConnectOptionsContext) {}

// EnterAccountLockPasswordExpireOptions is called when production accountLockPasswordExpireOptions is entered.
func (s *BaseTiDBParserListener) EnterAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) {
}

// ExitAccountLockPasswordExpireOptions is called when production accountLockPasswordExpireOptions is exited.
func (s *BaseTiDBParserListener) ExitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) {
}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseTiDBParserListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseTiDBParserListener) ExitDropUser(ctx *DropUserContext) {}

// EnterGrant is called when production grant is entered.
func (s *BaseTiDBParserListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseTiDBParserListener) ExitGrant(ctx *GrantContext) {}

// EnterGrantTargetList is called when production grantTargetList is entered.
func (s *BaseTiDBParserListener) EnterGrantTargetList(ctx *GrantTargetListContext) {}

// ExitGrantTargetList is called when production grantTargetList is exited.
func (s *BaseTiDBParserListener) ExitGrantTargetList(ctx *GrantTargetListContext) {}

// EnterGrantOptions is called when production grantOptions is entered.
func (s *BaseTiDBParserListener) EnterGrantOptions(ctx *GrantOptionsContext) {}

// ExitGrantOptions is called when production grantOptions is exited.
func (s *BaseTiDBParserListener) ExitGrantOptions(ctx *GrantOptionsContext) {}

// EnterExceptRoleList is called when production exceptRoleList is entered.
func (s *BaseTiDBParserListener) EnterExceptRoleList(ctx *ExceptRoleListContext) {}

// ExitExceptRoleList is called when production exceptRoleList is exited.
func (s *BaseTiDBParserListener) ExitExceptRoleList(ctx *ExceptRoleListContext) {}

// EnterWithRoles is called when production withRoles is entered.
func (s *BaseTiDBParserListener) EnterWithRoles(ctx *WithRolesContext) {}

// ExitWithRoles is called when production withRoles is exited.
func (s *BaseTiDBParserListener) ExitWithRoles(ctx *WithRolesContext) {}

// EnterGrantAs is called when production grantAs is entered.
func (s *BaseTiDBParserListener) EnterGrantAs(ctx *GrantAsContext) {}

// ExitGrantAs is called when production grantAs is exited.
func (s *BaseTiDBParserListener) ExitGrantAs(ctx *GrantAsContext) {}

// EnterVersionedRequireClause is called when production versionedRequireClause is entered.
func (s *BaseTiDBParserListener) EnterVersionedRequireClause(ctx *VersionedRequireClauseContext) {}

// ExitVersionedRequireClause is called when production versionedRequireClause is exited.
func (s *BaseTiDBParserListener) ExitVersionedRequireClause(ctx *VersionedRequireClauseContext) {}

// EnterRenameUser is called when production renameUser is entered.
func (s *BaseTiDBParserListener) EnterRenameUser(ctx *RenameUserContext) {}

// ExitRenameUser is called when production renameUser is exited.
func (s *BaseTiDBParserListener) ExitRenameUser(ctx *RenameUserContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseTiDBParserListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseTiDBParserListener) ExitRevoke(ctx *RevokeContext) {}

// EnterOnTypeTo is called when production onTypeTo is entered.
func (s *BaseTiDBParserListener) EnterOnTypeTo(ctx *OnTypeToContext) {}

// ExitOnTypeTo is called when production onTypeTo is exited.
func (s *BaseTiDBParserListener) ExitOnTypeTo(ctx *OnTypeToContext) {}

// EnterAclType is called when production aclType is entered.
func (s *BaseTiDBParserListener) EnterAclType(ctx *AclTypeContext) {}

// ExitAclType is called when production aclType is exited.
func (s *BaseTiDBParserListener) ExitAclType(ctx *AclTypeContext) {}

// EnterRoleOrPrivilegesList is called when production roleOrPrivilegesList is entered.
func (s *BaseTiDBParserListener) EnterRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) {}

// ExitRoleOrPrivilegesList is called when production roleOrPrivilegesList is exited.
func (s *BaseTiDBParserListener) ExitRoleOrPrivilegesList(ctx *RoleOrPrivilegesListContext) {}

// EnterRoleOrPrivilege is called when production roleOrPrivilege is entered.
func (s *BaseTiDBParserListener) EnterRoleOrPrivilege(ctx *RoleOrPrivilegeContext) {}

// ExitRoleOrPrivilege is called when production roleOrPrivilege is exited.
func (s *BaseTiDBParserListener) ExitRoleOrPrivilege(ctx *RoleOrPrivilegeContext) {}

// EnterGrantIdentifier is called when production grantIdentifier is entered.
func (s *BaseTiDBParserListener) EnterGrantIdentifier(ctx *GrantIdentifierContext) {}

// ExitGrantIdentifier is called when production grantIdentifier is exited.
func (s *BaseTiDBParserListener) ExitGrantIdentifier(ctx *GrantIdentifierContext) {}

// EnterRequireList is called when production requireList is entered.
func (s *BaseTiDBParserListener) EnterRequireList(ctx *RequireListContext) {}

// ExitRequireList is called when production requireList is exited.
func (s *BaseTiDBParserListener) ExitRequireList(ctx *RequireListContext) {}

// EnterRequireListElement is called when production requireListElement is entered.
func (s *BaseTiDBParserListener) EnterRequireListElement(ctx *RequireListElementContext) {}

// ExitRequireListElement is called when production requireListElement is exited.
func (s *BaseTiDBParserListener) ExitRequireListElement(ctx *RequireListElementContext) {}

// EnterGrantOption is called when production grantOption is entered.
func (s *BaseTiDBParserListener) EnterGrantOption(ctx *GrantOptionContext) {}

// ExitGrantOption is called when production grantOption is exited.
func (s *BaseTiDBParserListener) ExitGrantOption(ctx *GrantOptionContext) {}

// EnterSetRole is called when production setRole is entered.
func (s *BaseTiDBParserListener) EnterSetRole(ctx *SetRoleContext) {}

// ExitSetRole is called when production setRole is exited.
func (s *BaseTiDBParserListener) ExitSetRole(ctx *SetRoleContext) {}

// EnterRoleList is called when production roleList is entered.
func (s *BaseTiDBParserListener) EnterRoleList(ctx *RoleListContext) {}

// ExitRoleList is called when production roleList is exited.
func (s *BaseTiDBParserListener) ExitRoleList(ctx *RoleListContext) {}

// EnterRole is called when production role is entered.
func (s *BaseTiDBParserListener) EnterRole(ctx *RoleContext) {}

// ExitRole is called when production role is exited.
func (s *BaseTiDBParserListener) ExitRole(ctx *RoleContext) {}

// EnterTableAdministrationStatement is called when production tableAdministrationStatement is entered.
func (s *BaseTiDBParserListener) EnterTableAdministrationStatement(ctx *TableAdministrationStatementContext) {
}

// ExitTableAdministrationStatement is called when production tableAdministrationStatement is exited.
func (s *BaseTiDBParserListener) ExitTableAdministrationStatement(ctx *TableAdministrationStatementContext) {
}

// EnterHistogram is called when production histogram is entered.
func (s *BaseTiDBParserListener) EnterHistogram(ctx *HistogramContext) {}

// ExitHistogram is called when production histogram is exited.
func (s *BaseTiDBParserListener) ExitHistogram(ctx *HistogramContext) {}

// EnterCheckOption is called when production checkOption is entered.
func (s *BaseTiDBParserListener) EnterCheckOption(ctx *CheckOptionContext) {}

// ExitCheckOption is called when production checkOption is exited.
func (s *BaseTiDBParserListener) ExitCheckOption(ctx *CheckOptionContext) {}

// EnterRepairType is called when production repairType is entered.
func (s *BaseTiDBParserListener) EnterRepairType(ctx *RepairTypeContext) {}

// ExitRepairType is called when production repairType is exited.
func (s *BaseTiDBParserListener) ExitRepairType(ctx *RepairTypeContext) {}

// EnterInstallUninstallStatment is called when production installUninstallStatment is entered.
func (s *BaseTiDBParserListener) EnterInstallUninstallStatment(ctx *InstallUninstallStatmentContext) {
}

// ExitInstallUninstallStatment is called when production installUninstallStatment is exited.
func (s *BaseTiDBParserListener) ExitInstallUninstallStatment(ctx *InstallUninstallStatmentContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseTiDBParserListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseTiDBParserListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterStartOptionValueList is called when production startOptionValueList is entered.
func (s *BaseTiDBParserListener) EnterStartOptionValueList(ctx *StartOptionValueListContext) {}

// ExitStartOptionValueList is called when production startOptionValueList is exited.
func (s *BaseTiDBParserListener) ExitStartOptionValueList(ctx *StartOptionValueListContext) {}

// EnterTransactionCharacteristics is called when production transactionCharacteristics is entered.
func (s *BaseTiDBParserListener) EnterTransactionCharacteristics(ctx *TransactionCharacteristicsContext) {
}

// ExitTransactionCharacteristics is called when production transactionCharacteristics is exited.
func (s *BaseTiDBParserListener) ExitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) {
}

// EnterTransactionAccessMode is called when production transactionAccessMode is entered.
func (s *BaseTiDBParserListener) EnterTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// ExitTransactionAccessMode is called when production transactionAccessMode is exited.
func (s *BaseTiDBParserListener) ExitTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// EnterIsolationLevel is called when production isolationLevel is entered.
func (s *BaseTiDBParserListener) EnterIsolationLevel(ctx *IsolationLevelContext) {}

// ExitIsolationLevel is called when production isolationLevel is exited.
func (s *BaseTiDBParserListener) ExitIsolationLevel(ctx *IsolationLevelContext) {}

// EnterOptionValueListContinued is called when production optionValueListContinued is entered.
func (s *BaseTiDBParserListener) EnterOptionValueListContinued(ctx *OptionValueListContinuedContext) {
}

// ExitOptionValueListContinued is called when production optionValueListContinued is exited.
func (s *BaseTiDBParserListener) ExitOptionValueListContinued(ctx *OptionValueListContinuedContext) {}

// EnterOptionValueNoOptionType is called when production optionValueNoOptionType is entered.
func (s *BaseTiDBParserListener) EnterOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) {}

// ExitOptionValueNoOptionType is called when production optionValueNoOptionType is exited.
func (s *BaseTiDBParserListener) ExitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) {}

// EnterOptionValue is called when production optionValue is entered.
func (s *BaseTiDBParserListener) EnterOptionValue(ctx *OptionValueContext) {}

// ExitOptionValue is called when production optionValue is exited.
func (s *BaseTiDBParserListener) ExitOptionValue(ctx *OptionValueContext) {}

// EnterSetSystemVariable is called when production setSystemVariable is entered.
func (s *BaseTiDBParserListener) EnterSetSystemVariable(ctx *SetSystemVariableContext) {}

// ExitSetSystemVariable is called when production setSystemVariable is exited.
func (s *BaseTiDBParserListener) ExitSetSystemVariable(ctx *SetSystemVariableContext) {}

// EnterStartOptionValueListFollowingOptionType is called when production startOptionValueListFollowingOptionType is entered.
func (s *BaseTiDBParserListener) EnterStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) {
}

// ExitStartOptionValueListFollowingOptionType is called when production startOptionValueListFollowingOptionType is exited.
func (s *BaseTiDBParserListener) ExitStartOptionValueListFollowingOptionType(ctx *StartOptionValueListFollowingOptionTypeContext) {
}

// EnterOptionValueFollowingOptionType is called when production optionValueFollowingOptionType is entered.
func (s *BaseTiDBParserListener) EnterOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) {
}

// ExitOptionValueFollowingOptionType is called when production optionValueFollowingOptionType is exited.
func (s *BaseTiDBParserListener) ExitOptionValueFollowingOptionType(ctx *OptionValueFollowingOptionTypeContext) {
}

// EnterSetExprOrDefault is called when production setExprOrDefault is entered.
func (s *BaseTiDBParserListener) EnterSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// ExitSetExprOrDefault is called when production setExprOrDefault is exited.
func (s *BaseTiDBParserListener) ExitSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// EnterShowStatement is called when production showStatement is entered.
func (s *BaseTiDBParserListener) EnterShowStatement(ctx *ShowStatementContext) {}

// ExitShowStatement is called when production showStatement is exited.
func (s *BaseTiDBParserListener) ExitShowStatement(ctx *ShowStatementContext) {}

// EnterShowCommandType is called when production showCommandType is entered.
func (s *BaseTiDBParserListener) EnterShowCommandType(ctx *ShowCommandTypeContext) {}

// ExitShowCommandType is called when production showCommandType is exited.
func (s *BaseTiDBParserListener) ExitShowCommandType(ctx *ShowCommandTypeContext) {}

// EnterNonBlocking is called when production nonBlocking is entered.
func (s *BaseTiDBParserListener) EnterNonBlocking(ctx *NonBlockingContext) {}

// ExitNonBlocking is called when production nonBlocking is exited.
func (s *BaseTiDBParserListener) ExitNonBlocking(ctx *NonBlockingContext) {}

// EnterFromOrIn is called when production fromOrIn is entered.
func (s *BaseTiDBParserListener) EnterFromOrIn(ctx *FromOrInContext) {}

// ExitFromOrIn is called when production fromOrIn is exited.
func (s *BaseTiDBParserListener) ExitFromOrIn(ctx *FromOrInContext) {}

// EnterInDb is called when production inDb is entered.
func (s *BaseTiDBParserListener) EnterInDb(ctx *InDbContext) {}

// ExitInDb is called when production inDb is exited.
func (s *BaseTiDBParserListener) ExitInDb(ctx *InDbContext) {}

// EnterProfileType is called when production profileType is entered.
func (s *BaseTiDBParserListener) EnterProfileType(ctx *ProfileTypeContext) {}

// ExitProfileType is called when production profileType is exited.
func (s *BaseTiDBParserListener) ExitProfileType(ctx *ProfileTypeContext) {}

// EnterResourceGroupManagement is called when production resourceGroupManagement is entered.
func (s *BaseTiDBParserListener) EnterResourceGroupManagement(ctx *ResourceGroupManagementContext) {}

// ExitResourceGroupManagement is called when production resourceGroupManagement is exited.
func (s *BaseTiDBParserListener) ExitResourceGroupManagement(ctx *ResourceGroupManagementContext) {}

// EnterCreateResourceGroup is called when production createResourceGroup is entered.
func (s *BaseTiDBParserListener) EnterCreateResourceGroup(ctx *CreateResourceGroupContext) {}

// ExitCreateResourceGroup is called when production createResourceGroup is exited.
func (s *BaseTiDBParserListener) ExitCreateResourceGroup(ctx *CreateResourceGroupContext) {}

// EnterResourceGroupVcpuList is called when production resourceGroupVcpuList is entered.
func (s *BaseTiDBParserListener) EnterResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) {}

// ExitResourceGroupVcpuList is called when production resourceGroupVcpuList is exited.
func (s *BaseTiDBParserListener) ExitResourceGroupVcpuList(ctx *ResourceGroupVcpuListContext) {}

// EnterVcpuNumOrRange is called when production vcpuNumOrRange is entered.
func (s *BaseTiDBParserListener) EnterVcpuNumOrRange(ctx *VcpuNumOrRangeContext) {}

// ExitVcpuNumOrRange is called when production vcpuNumOrRange is exited.
func (s *BaseTiDBParserListener) ExitVcpuNumOrRange(ctx *VcpuNumOrRangeContext) {}

// EnterResourceGroupPriority is called when production resourceGroupPriority is entered.
func (s *BaseTiDBParserListener) EnterResourceGroupPriority(ctx *ResourceGroupPriorityContext) {}

// ExitResourceGroupPriority is called when production resourceGroupPriority is exited.
func (s *BaseTiDBParserListener) ExitResourceGroupPriority(ctx *ResourceGroupPriorityContext) {}

// EnterResourceGroupEnableDisable is called when production resourceGroupEnableDisable is entered.
func (s *BaseTiDBParserListener) EnterResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) {
}

// ExitResourceGroupEnableDisable is called when production resourceGroupEnableDisable is exited.
func (s *BaseTiDBParserListener) ExitResourceGroupEnableDisable(ctx *ResourceGroupEnableDisableContext) {
}

// EnterAlterResourceGroup is called when production alterResourceGroup is entered.
func (s *BaseTiDBParserListener) EnterAlterResourceGroup(ctx *AlterResourceGroupContext) {}

// ExitAlterResourceGroup is called when production alterResourceGroup is exited.
func (s *BaseTiDBParserListener) ExitAlterResourceGroup(ctx *AlterResourceGroupContext) {}

// EnterSetResourceGroup is called when production setResourceGroup is entered.
func (s *BaseTiDBParserListener) EnterSetResourceGroup(ctx *SetResourceGroupContext) {}

// ExitSetResourceGroup is called when production setResourceGroup is exited.
func (s *BaseTiDBParserListener) ExitSetResourceGroup(ctx *SetResourceGroupContext) {}

// EnterThreadIdList is called when production threadIdList is entered.
func (s *BaseTiDBParserListener) EnterThreadIdList(ctx *ThreadIdListContext) {}

// ExitThreadIdList is called when production threadIdList is exited.
func (s *BaseTiDBParserListener) ExitThreadIdList(ctx *ThreadIdListContext) {}

// EnterDropResourceGroup is called when production dropResourceGroup is entered.
func (s *BaseTiDBParserListener) EnterDropResourceGroup(ctx *DropResourceGroupContext) {}

// ExitDropResourceGroup is called when production dropResourceGroup is exited.
func (s *BaseTiDBParserListener) ExitDropResourceGroup(ctx *DropResourceGroupContext) {}

// EnterExprOr is called when production exprOr is entered.
func (s *BaseTiDBParserListener) EnterExprOr(ctx *ExprOrContext) {}

// ExitExprOr is called when production exprOr is exited.
func (s *BaseTiDBParserListener) ExitExprOr(ctx *ExprOrContext) {}

// EnterExprNot is called when production exprNot is entered.
func (s *BaseTiDBParserListener) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production exprNot is exited.
func (s *BaseTiDBParserListener) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprIs is called when production exprIs is entered.
func (s *BaseTiDBParserListener) EnterExprIs(ctx *ExprIsContext) {}

// ExitExprIs is called when production exprIs is exited.
func (s *BaseTiDBParserListener) ExitExprIs(ctx *ExprIsContext) {}

// EnterExprAnd is called when production exprAnd is entered.
func (s *BaseTiDBParserListener) EnterExprAnd(ctx *ExprAndContext) {}

// ExitExprAnd is called when production exprAnd is exited.
func (s *BaseTiDBParserListener) ExitExprAnd(ctx *ExprAndContext) {}

// EnterExprXor is called when production exprXor is entered.
func (s *BaseTiDBParserListener) EnterExprXor(ctx *ExprXorContext) {}

// ExitExprXor is called when production exprXor is exited.
func (s *BaseTiDBParserListener) ExitExprXor(ctx *ExprXorContext) {}

// EnterPrimaryExprPredicate is called when production primaryExprPredicate is entered.
func (s *BaseTiDBParserListener) EnterPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) {}

// ExitPrimaryExprPredicate is called when production primaryExprPredicate is exited.
func (s *BaseTiDBParserListener) ExitPrimaryExprPredicate(ctx *PrimaryExprPredicateContext) {}

// EnterPrimaryExprCompare is called when production primaryExprCompare is entered.
func (s *BaseTiDBParserListener) EnterPrimaryExprCompare(ctx *PrimaryExprCompareContext) {}

// ExitPrimaryExprCompare is called when production primaryExprCompare is exited.
func (s *BaseTiDBParserListener) ExitPrimaryExprCompare(ctx *PrimaryExprCompareContext) {}

// EnterPrimaryExprAllAny is called when production primaryExprAllAny is entered.
func (s *BaseTiDBParserListener) EnterPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) {}

// ExitPrimaryExprAllAny is called when production primaryExprAllAny is exited.
func (s *BaseTiDBParserListener) ExitPrimaryExprAllAny(ctx *PrimaryExprAllAnyContext) {}

// EnterPrimaryExprIsNull is called when production primaryExprIsNull is entered.
func (s *BaseTiDBParserListener) EnterPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) {}

// ExitPrimaryExprIsNull is called when production primaryExprIsNull is exited.
func (s *BaseTiDBParserListener) ExitPrimaryExprIsNull(ctx *PrimaryExprIsNullContext) {}

// EnterCompOp is called when production compOp is entered.
func (s *BaseTiDBParserListener) EnterCompOp(ctx *CompOpContext) {}

// ExitCompOp is called when production compOp is exited.
func (s *BaseTiDBParserListener) ExitCompOp(ctx *CompOpContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseTiDBParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseTiDBParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterPredicateExprIn is called when production predicateExprIn is entered.
func (s *BaseTiDBParserListener) EnterPredicateExprIn(ctx *PredicateExprInContext) {}

// ExitPredicateExprIn is called when production predicateExprIn is exited.
func (s *BaseTiDBParserListener) ExitPredicateExprIn(ctx *PredicateExprInContext) {}

// EnterPredicateExprBetween is called when production predicateExprBetween is entered.
func (s *BaseTiDBParserListener) EnterPredicateExprBetween(ctx *PredicateExprBetweenContext) {}

// ExitPredicateExprBetween is called when production predicateExprBetween is exited.
func (s *BaseTiDBParserListener) ExitPredicateExprBetween(ctx *PredicateExprBetweenContext) {}

// EnterPredicateExprLike is called when production predicateExprLike is entered.
func (s *BaseTiDBParserListener) EnterPredicateExprLike(ctx *PredicateExprLikeContext) {}

// ExitPredicateExprLike is called when production predicateExprLike is exited.
func (s *BaseTiDBParserListener) ExitPredicateExprLike(ctx *PredicateExprLikeContext) {}

// EnterPredicateExprRegex is called when production predicateExprRegex is entered.
func (s *BaseTiDBParserListener) EnterPredicateExprRegex(ctx *PredicateExprRegexContext) {}

// ExitPredicateExprRegex is called when production predicateExprRegex is exited.
func (s *BaseTiDBParserListener) ExitPredicateExprRegex(ctx *PredicateExprRegexContext) {}

// EnterBitExpr is called when production bitExpr is entered.
func (s *BaseTiDBParserListener) EnterBitExpr(ctx *BitExprContext) {}

// ExitBitExpr is called when production bitExpr is exited.
func (s *BaseTiDBParserListener) ExitBitExpr(ctx *BitExprContext) {}

// EnterSimpleExprConvert is called when production simpleExprConvert is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprConvert(ctx *SimpleExprConvertContext) {}

// ExitSimpleExprConvert is called when production simpleExprConvert is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprConvert(ctx *SimpleExprConvertContext) {}

// EnterSimpleExprSearchJson is called when production simpleExprSearchJson is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) {}

// ExitSimpleExprSearchJson is called when production simpleExprSearchJson is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprSearchJson(ctx *SimpleExprSearchJsonContext) {}

// EnterSimpleExprVariable is called when production simpleExprVariable is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprVariable(ctx *SimpleExprVariableContext) {}

// ExitSimpleExprVariable is called when production simpleExprVariable is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprVariable(ctx *SimpleExprVariableContext) {}

// EnterSimpleExprCast is called when production simpleExprCast is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprCast(ctx *SimpleExprCastContext) {}

// ExitSimpleExprCast is called when production simpleExprCast is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprCast(ctx *SimpleExprCastContext) {}

// EnterSimpleExprUnary is called when production simpleExprUnary is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprUnary(ctx *SimpleExprUnaryContext) {}

// ExitSimpleExprUnary is called when production simpleExprUnary is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprUnary(ctx *SimpleExprUnaryContext) {}

// EnterSimpleExprOdbc is called when production simpleExprOdbc is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprOdbc(ctx *SimpleExprOdbcContext) {}

// ExitSimpleExprOdbc is called when production simpleExprOdbc is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprOdbc(ctx *SimpleExprOdbcContext) {}

// EnterSimpleExprRuntimeFunction is called when production simpleExprRuntimeFunction is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) {
}

// ExitSimpleExprRuntimeFunction is called when production simpleExprRuntimeFunction is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprRuntimeFunction(ctx *SimpleExprRuntimeFunctionContext) {
}

// EnterSimpleExprFunction is called when production simpleExprFunction is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprFunction(ctx *SimpleExprFunctionContext) {}

// ExitSimpleExprFunction is called when production simpleExprFunction is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprFunction(ctx *SimpleExprFunctionContext) {}

// EnterSimpleExprCollate is called when production simpleExprCollate is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprCollate(ctx *SimpleExprCollateContext) {}

// ExitSimpleExprCollate is called when production simpleExprCollate is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprCollate(ctx *SimpleExprCollateContext) {}

// EnterSimpleExprMatch is called when production simpleExprMatch is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprMatch(ctx *SimpleExprMatchContext) {}

// ExitSimpleExprMatch is called when production simpleExprMatch is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprMatch(ctx *SimpleExprMatchContext) {}

// EnterSimpleExprWindowingFunction is called when production simpleExprWindowingFunction is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) {
}

// ExitSimpleExprWindowingFunction is called when production simpleExprWindowingFunction is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprWindowingFunction(ctx *SimpleExprWindowingFunctionContext) {
}

// EnterSimpleExprBinary is called when production simpleExprBinary is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprBinary(ctx *SimpleExprBinaryContext) {}

// ExitSimpleExprBinary is called when production simpleExprBinary is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprBinary(ctx *SimpleExprBinaryContext) {}

// EnterSimpleExprColumnRef is called when production simpleExprColumnRef is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) {}

// ExitSimpleExprColumnRef is called when production simpleExprColumnRef is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprColumnRef(ctx *SimpleExprColumnRefContext) {}

// EnterSimpleExprParamMarker is called when production simpleExprParamMarker is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) {}

// ExitSimpleExprParamMarker is called when production simpleExprParamMarker is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprParamMarker(ctx *SimpleExprParamMarkerContext) {}

// EnterSimpleExprSum is called when production simpleExprSum is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprSum(ctx *SimpleExprSumContext) {}

// ExitSimpleExprSum is called when production simpleExprSum is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprSum(ctx *SimpleExprSumContext) {}

// EnterSimpleExprConvertUsing is called when production simpleExprConvertUsing is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) {}

// ExitSimpleExprConvertUsing is called when production simpleExprConvertUsing is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprConvertUsing(ctx *SimpleExprConvertUsingContext) {}

// EnterSimpleExprSubQuery is called when production simpleExprSubQuery is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) {}

// ExitSimpleExprSubQuery is called when production simpleExprSubQuery is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprSubQuery(ctx *SimpleExprSubQueryContext) {}

// EnterSimpleExprGroupingOperation is called when production simpleExprGroupingOperation is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) {
}

// ExitSimpleExprGroupingOperation is called when production simpleExprGroupingOperation is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprGroupingOperation(ctx *SimpleExprGroupingOperationContext) {
}

// EnterSimpleExprNot is called when production simpleExprNot is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprNot(ctx *SimpleExprNotContext) {}

// ExitSimpleExprNot is called when production simpleExprNot is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprNot(ctx *SimpleExprNotContext) {}

// EnterSimpleExprValues is called when production simpleExprValues is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprValues(ctx *SimpleExprValuesContext) {}

// ExitSimpleExprValues is called when production simpleExprValues is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprValues(ctx *SimpleExprValuesContext) {}

// EnterSimpleExprDefault is called when production simpleExprDefault is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprDefault(ctx *SimpleExprDefaultContext) {}

// ExitSimpleExprDefault is called when production simpleExprDefault is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprDefault(ctx *SimpleExprDefaultContext) {}

// EnterSimpleExprList is called when production simpleExprList is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprList(ctx *SimpleExprListContext) {}

// ExitSimpleExprList is called when production simpleExprList is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprList(ctx *SimpleExprListContext) {}

// EnterSimpleExprInterval is called when production simpleExprInterval is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprInterval(ctx *SimpleExprIntervalContext) {}

// ExitSimpleExprInterval is called when production simpleExprInterval is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprInterval(ctx *SimpleExprIntervalContext) {}

// EnterSimpleExprCase is called when production simpleExprCase is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprCase(ctx *SimpleExprCaseContext) {}

// ExitSimpleExprCase is called when production simpleExprCase is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprCase(ctx *SimpleExprCaseContext) {}

// EnterSimpleExprConcat is called when production simpleExprConcat is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprConcat(ctx *SimpleExprConcatContext) {}

// ExitSimpleExprConcat is called when production simpleExprConcat is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprConcat(ctx *SimpleExprConcatContext) {}

// EnterSimpleExprLiteral is called when production simpleExprLiteral is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprLiteral(ctx *SimpleExprLiteralContext) {}

// ExitSimpleExprLiteral is called when production simpleExprLiteral is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprLiteral(ctx *SimpleExprLiteralContext) {}

// EnterArrayCast is called when production arrayCast is entered.
func (s *BaseTiDBParserListener) EnterArrayCast(ctx *ArrayCastContext) {}

// ExitArrayCast is called when production arrayCast is exited.
func (s *BaseTiDBParserListener) ExitArrayCast(ctx *ArrayCastContext) {}

// EnterJsonOperator is called when production jsonOperator is entered.
func (s *BaseTiDBParserListener) EnterJsonOperator(ctx *JsonOperatorContext) {}

// ExitJsonOperator is called when production jsonOperator is exited.
func (s *BaseTiDBParserListener) ExitJsonOperator(ctx *JsonOperatorContext) {}

// EnterSumExpr is called when production sumExpr is entered.
func (s *BaseTiDBParserListener) EnterSumExpr(ctx *SumExprContext) {}

// ExitSumExpr is called when production sumExpr is exited.
func (s *BaseTiDBParserListener) ExitSumExpr(ctx *SumExprContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseTiDBParserListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseTiDBParserListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterWindowFunctionCall is called when production windowFunctionCall is entered.
func (s *BaseTiDBParserListener) EnterWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// ExitWindowFunctionCall is called when production windowFunctionCall is exited.
func (s *BaseTiDBParserListener) ExitWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// EnterWindowingClause is called when production windowingClause is entered.
func (s *BaseTiDBParserListener) EnterWindowingClause(ctx *WindowingClauseContext) {}

// ExitWindowingClause is called when production windowingClause is exited.
func (s *BaseTiDBParserListener) ExitWindowingClause(ctx *WindowingClauseContext) {}

// EnterLeadLagInfo is called when production leadLagInfo is entered.
func (s *BaseTiDBParserListener) EnterLeadLagInfo(ctx *LeadLagInfoContext) {}

// ExitLeadLagInfo is called when production leadLagInfo is exited.
func (s *BaseTiDBParserListener) ExitLeadLagInfo(ctx *LeadLagInfoContext) {}

// EnterNullTreatment is called when production nullTreatment is entered.
func (s *BaseTiDBParserListener) EnterNullTreatment(ctx *NullTreatmentContext) {}

// ExitNullTreatment is called when production nullTreatment is exited.
func (s *BaseTiDBParserListener) ExitNullTreatment(ctx *NullTreatmentContext) {}

// EnterJsonFunction is called when production jsonFunction is entered.
func (s *BaseTiDBParserListener) EnterJsonFunction(ctx *JsonFunctionContext) {}

// ExitJsonFunction is called when production jsonFunction is exited.
func (s *BaseTiDBParserListener) ExitJsonFunction(ctx *JsonFunctionContext) {}

// EnterInSumExpr is called when production inSumExpr is entered.
func (s *BaseTiDBParserListener) EnterInSumExpr(ctx *InSumExprContext) {}

// ExitInSumExpr is called when production inSumExpr is exited.
func (s *BaseTiDBParserListener) ExitInSumExpr(ctx *InSumExprContext) {}

// EnterIdentListArg is called when production identListArg is entered.
func (s *BaseTiDBParserListener) EnterIdentListArg(ctx *IdentListArgContext) {}

// ExitIdentListArg is called when production identListArg is exited.
func (s *BaseTiDBParserListener) ExitIdentListArg(ctx *IdentListArgContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseTiDBParserListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseTiDBParserListener) ExitIdentList(ctx *IdentListContext) {}

// EnterFulltextOptions is called when production fulltextOptions is entered.
func (s *BaseTiDBParserListener) EnterFulltextOptions(ctx *FulltextOptionsContext) {}

// ExitFulltextOptions is called when production fulltextOptions is exited.
func (s *BaseTiDBParserListener) ExitFulltextOptions(ctx *FulltextOptionsContext) {}

// EnterRuntimeFunctionCall is called when production runtimeFunctionCall is entered.
func (s *BaseTiDBParserListener) EnterRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) {}

// ExitRuntimeFunctionCall is called when production runtimeFunctionCall is exited.
func (s *BaseTiDBParserListener) ExitRuntimeFunctionCall(ctx *RuntimeFunctionCallContext) {}

// EnterGeometryFunction is called when production geometryFunction is entered.
func (s *BaseTiDBParserListener) EnterGeometryFunction(ctx *GeometryFunctionContext) {}

// ExitGeometryFunction is called when production geometryFunction is exited.
func (s *BaseTiDBParserListener) ExitGeometryFunction(ctx *GeometryFunctionContext) {}

// EnterTimeFunctionParameters is called when production timeFunctionParameters is entered.
func (s *BaseTiDBParserListener) EnterTimeFunctionParameters(ctx *TimeFunctionParametersContext) {}

// ExitTimeFunctionParameters is called when production timeFunctionParameters is exited.
func (s *BaseTiDBParserListener) ExitTimeFunctionParameters(ctx *TimeFunctionParametersContext) {}

// EnterFractionalPrecision is called when production fractionalPrecision is entered.
func (s *BaseTiDBParserListener) EnterFractionalPrecision(ctx *FractionalPrecisionContext) {}

// ExitFractionalPrecision is called when production fractionalPrecision is exited.
func (s *BaseTiDBParserListener) ExitFractionalPrecision(ctx *FractionalPrecisionContext) {}

// EnterWeightStringLevels is called when production weightStringLevels is entered.
func (s *BaseTiDBParserListener) EnterWeightStringLevels(ctx *WeightStringLevelsContext) {}

// ExitWeightStringLevels is called when production weightStringLevels is exited.
func (s *BaseTiDBParserListener) ExitWeightStringLevels(ctx *WeightStringLevelsContext) {}

// EnterWeightStringLevelListItem is called when production weightStringLevelListItem is entered.
func (s *BaseTiDBParserListener) EnterWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) {
}

// ExitWeightStringLevelListItem is called when production weightStringLevelListItem is exited.
func (s *BaseTiDBParserListener) ExitWeightStringLevelListItem(ctx *WeightStringLevelListItemContext) {
}

// EnterDateTimeTtype is called when production dateTimeTtype is entered.
func (s *BaseTiDBParserListener) EnterDateTimeTtype(ctx *DateTimeTtypeContext) {}

// ExitDateTimeTtype is called when production dateTimeTtype is exited.
func (s *BaseTiDBParserListener) ExitDateTimeTtype(ctx *DateTimeTtypeContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *BaseTiDBParserListener) EnterTrimFunction(ctx *TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *BaseTiDBParserListener) ExitTrimFunction(ctx *TrimFunctionContext) {}

// EnterSubstringFunction is called when production substringFunction is entered.
func (s *BaseTiDBParserListener) EnterSubstringFunction(ctx *SubstringFunctionContext) {}

// ExitSubstringFunction is called when production substringFunction is exited.
func (s *BaseTiDBParserListener) ExitSubstringFunction(ctx *SubstringFunctionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseTiDBParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseTiDBParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterSearchJsonFunction is called when production searchJsonFunction is entered.
func (s *BaseTiDBParserListener) EnterSearchJsonFunction(ctx *SearchJsonFunctionContext) {}

// ExitSearchJsonFunction is called when production searchJsonFunction is exited.
func (s *BaseTiDBParserListener) ExitSearchJsonFunction(ctx *SearchJsonFunctionContext) {}

// EnterJsonValueReturning is called when production jsonValueReturning is entered.
func (s *BaseTiDBParserListener) EnterJsonValueReturning(ctx *JsonValueReturningContext) {}

// ExitJsonValueReturning is called when production jsonValueReturning is exited.
func (s *BaseTiDBParserListener) ExitJsonValueReturning(ctx *JsonValueReturningContext) {}

// EnterJsonValueOnEmpty is called when production jsonValueOnEmpty is entered.
func (s *BaseTiDBParserListener) EnterJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) {}

// ExitJsonValueOnEmpty is called when production jsonValueOnEmpty is exited.
func (s *BaseTiDBParserListener) ExitJsonValueOnEmpty(ctx *JsonValueOnEmptyContext) {}

// EnterJsonValueOnError is called when production jsonValueOnError is entered.
func (s *BaseTiDBParserListener) EnterJsonValueOnError(ctx *JsonValueOnErrorContext) {}

// ExitJsonValueOnError is called when production jsonValueOnError is exited.
func (s *BaseTiDBParserListener) ExitJsonValueOnError(ctx *JsonValueOnErrorContext) {}

// EnterUdfExprList is called when production udfExprList is entered.
func (s *BaseTiDBParserListener) EnterUdfExprList(ctx *UdfExprListContext) {}

// ExitUdfExprList is called when production udfExprList is exited.
func (s *BaseTiDBParserListener) ExitUdfExprList(ctx *UdfExprListContext) {}

// EnterUdfExpr is called when production udfExpr is entered.
func (s *BaseTiDBParserListener) EnterUdfExpr(ctx *UdfExprContext) {}

// ExitUdfExpr is called when production udfExpr is exited.
func (s *BaseTiDBParserListener) ExitUdfExpr(ctx *UdfExprContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseTiDBParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseTiDBParserListener) ExitVariable(ctx *VariableContext) {}

// EnterUserVariable is called when production userVariable is entered.
func (s *BaseTiDBParserListener) EnterUserVariable(ctx *UserVariableContext) {}

// ExitUserVariable is called when production userVariable is exited.
func (s *BaseTiDBParserListener) ExitUserVariable(ctx *UserVariableContext) {}

// EnterSystemVariable is called when production systemVariable is entered.
func (s *BaseTiDBParserListener) EnterSystemVariable(ctx *SystemVariableContext) {}

// ExitSystemVariable is called when production systemVariable is exited.
func (s *BaseTiDBParserListener) ExitSystemVariable(ctx *SystemVariableContext) {}

// EnterInternalVariableName is called when production internalVariableName is entered.
func (s *BaseTiDBParserListener) EnterInternalVariableName(ctx *InternalVariableNameContext) {}

// ExitInternalVariableName is called when production internalVariableName is exited.
func (s *BaseTiDBParserListener) ExitInternalVariableName(ctx *InternalVariableNameContext) {}

// EnterWhenExpression is called when production whenExpression is entered.
func (s *BaseTiDBParserListener) EnterWhenExpression(ctx *WhenExpressionContext) {}

// ExitWhenExpression is called when production whenExpression is exited.
func (s *BaseTiDBParserListener) ExitWhenExpression(ctx *WhenExpressionContext) {}

// EnterThenExpression is called when production thenExpression is entered.
func (s *BaseTiDBParserListener) EnterThenExpression(ctx *ThenExpressionContext) {}

// ExitThenExpression is called when production thenExpression is exited.
func (s *BaseTiDBParserListener) ExitThenExpression(ctx *ThenExpressionContext) {}

// EnterElseExpression is called when production elseExpression is entered.
func (s *BaseTiDBParserListener) EnterElseExpression(ctx *ElseExpressionContext) {}

// ExitElseExpression is called when production elseExpression is exited.
func (s *BaseTiDBParserListener) ExitElseExpression(ctx *ElseExpressionContext) {}

// EnterCastType is called when production castType is entered.
func (s *BaseTiDBParserListener) EnterCastType(ctx *CastTypeContext) {}

// ExitCastType is called when production castType is exited.
func (s *BaseTiDBParserListener) ExitCastType(ctx *CastTypeContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseTiDBParserListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseTiDBParserListener) ExitExprList(ctx *ExprListContext) {}

// EnterCharset is called when production charset is entered.
func (s *BaseTiDBParserListener) EnterCharset(ctx *CharsetContext) {}

// ExitCharset is called when production charset is exited.
func (s *BaseTiDBParserListener) ExitCharset(ctx *CharsetContext) {}

// EnterNotRule is called when production notRule is entered.
func (s *BaseTiDBParserListener) EnterNotRule(ctx *NotRuleContext) {}

// ExitNotRule is called when production notRule is exited.
func (s *BaseTiDBParserListener) ExitNotRule(ctx *NotRuleContext) {}

// EnterNot2Rule is called when production not2Rule is entered.
func (s *BaseTiDBParserListener) EnterNot2Rule(ctx *Not2RuleContext) {}

// ExitNot2Rule is called when production not2Rule is exited.
func (s *BaseTiDBParserListener) ExitNot2Rule(ctx *Not2RuleContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseTiDBParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseTiDBParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterIntervalTimeStamp is called when production intervalTimeStamp is entered.
func (s *BaseTiDBParserListener) EnterIntervalTimeStamp(ctx *IntervalTimeStampContext) {}

// ExitIntervalTimeStamp is called when production intervalTimeStamp is exited.
func (s *BaseTiDBParserListener) ExitIntervalTimeStamp(ctx *IntervalTimeStampContext) {}

// EnterExprListWithParentheses is called when production exprListWithParentheses is entered.
func (s *BaseTiDBParserListener) EnterExprListWithParentheses(ctx *ExprListWithParenthesesContext) {}

// ExitExprListWithParentheses is called when production exprListWithParentheses is exited.
func (s *BaseTiDBParserListener) ExitExprListWithParentheses(ctx *ExprListWithParenthesesContext) {}

// EnterExprWithParentheses is called when production exprWithParentheses is entered.
func (s *BaseTiDBParserListener) EnterExprWithParentheses(ctx *ExprWithParenthesesContext) {}

// ExitExprWithParentheses is called when production exprWithParentheses is exited.
func (s *BaseTiDBParserListener) ExitExprWithParentheses(ctx *ExprWithParenthesesContext) {}

// EnterSimpleExprWithParentheses is called when production simpleExprWithParentheses is entered.
func (s *BaseTiDBParserListener) EnterSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) {
}

// ExitSimpleExprWithParentheses is called when production simpleExprWithParentheses is exited.
func (s *BaseTiDBParserListener) ExitSimpleExprWithParentheses(ctx *SimpleExprWithParenthesesContext) {
}

// EnterOrderList is called when production orderList is entered.
func (s *BaseTiDBParserListener) EnterOrderList(ctx *OrderListContext) {}

// ExitOrderList is called when production orderList is exited.
func (s *BaseTiDBParserListener) ExitOrderList(ctx *OrderListContext) {}

// EnterOrderExpression is called when production orderExpression is entered.
func (s *BaseTiDBParserListener) EnterOrderExpression(ctx *OrderExpressionContext) {}

// ExitOrderExpression is called when production orderExpression is exited.
func (s *BaseTiDBParserListener) ExitOrderExpression(ctx *OrderExpressionContext) {}

// EnterGroupList is called when production groupList is entered.
func (s *BaseTiDBParserListener) EnterGroupList(ctx *GroupListContext) {}

// ExitGroupList is called when production groupList is exited.
func (s *BaseTiDBParserListener) ExitGroupList(ctx *GroupListContext) {}

// EnterGroupingExpression is called when production groupingExpression is entered.
func (s *BaseTiDBParserListener) EnterGroupingExpression(ctx *GroupingExpressionContext) {}

// ExitGroupingExpression is called when production groupingExpression is exited.
func (s *BaseTiDBParserListener) ExitGroupingExpression(ctx *GroupingExpressionContext) {}

// EnterChannel is called when production channel is entered.
func (s *BaseTiDBParserListener) EnterChannel(ctx *ChannelContext) {}

// ExitChannel is called when production channel is exited.
func (s *BaseTiDBParserListener) ExitChannel(ctx *ChannelContext) {}

// EnterColumnFormat is called when production columnFormat is entered.
func (s *BaseTiDBParserListener) EnterColumnFormat(ctx *ColumnFormatContext) {}

// ExitColumnFormat is called when production columnFormat is exited.
func (s *BaseTiDBParserListener) ExitColumnFormat(ctx *ColumnFormatContext) {}

// EnterStorageMedia is called when production storageMedia is entered.
func (s *BaseTiDBParserListener) EnterStorageMedia(ctx *StorageMediaContext) {}

// ExitStorageMedia is called when production storageMedia is exited.
func (s *BaseTiDBParserListener) ExitStorageMedia(ctx *StorageMediaContext) {}

// EnterGcolAttribute is called when production gcolAttribute is entered.
func (s *BaseTiDBParserListener) EnterGcolAttribute(ctx *GcolAttributeContext) {}

// ExitGcolAttribute is called when production gcolAttribute is exited.
func (s *BaseTiDBParserListener) ExitGcolAttribute(ctx *GcolAttributeContext) {}

// EnterReferences is called when production references is entered.
func (s *BaseTiDBParserListener) EnterReferences(ctx *ReferencesContext) {}

// ExitReferences is called when production references is exited.
func (s *BaseTiDBParserListener) ExitReferences(ctx *ReferencesContext) {}

// EnterDeleteOption is called when production deleteOption is entered.
func (s *BaseTiDBParserListener) EnterDeleteOption(ctx *DeleteOptionContext) {}

// ExitDeleteOption is called when production deleteOption is exited.
func (s *BaseTiDBParserListener) ExitDeleteOption(ctx *DeleteOptionContext) {}

// EnterKeyList is called when production keyList is entered.
func (s *BaseTiDBParserListener) EnterKeyList(ctx *KeyListContext) {}

// ExitKeyList is called when production keyList is exited.
func (s *BaseTiDBParserListener) ExitKeyList(ctx *KeyListContext) {}

// EnterKeyPart is called when production keyPart is entered.
func (s *BaseTiDBParserListener) EnterKeyPart(ctx *KeyPartContext) {}

// ExitKeyPart is called when production keyPart is exited.
func (s *BaseTiDBParserListener) ExitKeyPart(ctx *KeyPartContext) {}

// EnterKeyListWithExpression is called when production keyListWithExpression is entered.
func (s *BaseTiDBParserListener) EnterKeyListWithExpression(ctx *KeyListWithExpressionContext) {}

// ExitKeyListWithExpression is called when production keyListWithExpression is exited.
func (s *BaseTiDBParserListener) ExitKeyListWithExpression(ctx *KeyListWithExpressionContext) {}

// EnterKeyPartOrExpression is called when production keyPartOrExpression is entered.
func (s *BaseTiDBParserListener) EnterKeyPartOrExpression(ctx *KeyPartOrExpressionContext) {}

// ExitKeyPartOrExpression is called when production keyPartOrExpression is exited.
func (s *BaseTiDBParserListener) ExitKeyPartOrExpression(ctx *KeyPartOrExpressionContext) {}

// EnterKeyListVariants is called when production keyListVariants is entered.
func (s *BaseTiDBParserListener) EnterKeyListVariants(ctx *KeyListVariantsContext) {}

// ExitKeyListVariants is called when production keyListVariants is exited.
func (s *BaseTiDBParserListener) ExitKeyListVariants(ctx *KeyListVariantsContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseTiDBParserListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseTiDBParserListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterCommonIndexOption is called when production commonIndexOption is entered.
func (s *BaseTiDBParserListener) EnterCommonIndexOption(ctx *CommonIndexOptionContext) {}

// ExitCommonIndexOption is called when production commonIndexOption is exited.
func (s *BaseTiDBParserListener) ExitCommonIndexOption(ctx *CommonIndexOptionContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseTiDBParserListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseTiDBParserListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterIndexTypeClause is called when production indexTypeClause is entered.
func (s *BaseTiDBParserListener) EnterIndexTypeClause(ctx *IndexTypeClauseContext) {}

// ExitIndexTypeClause is called when production indexTypeClause is exited.
func (s *BaseTiDBParserListener) ExitIndexTypeClause(ctx *IndexTypeClauseContext) {}

// EnterFulltextIndexOption is called when production fulltextIndexOption is entered.
func (s *BaseTiDBParserListener) EnterFulltextIndexOption(ctx *FulltextIndexOptionContext) {}

// ExitFulltextIndexOption is called when production fulltextIndexOption is exited.
func (s *BaseTiDBParserListener) ExitFulltextIndexOption(ctx *FulltextIndexOptionContext) {}

// EnterSpatialIndexOption is called when production spatialIndexOption is entered.
func (s *BaseTiDBParserListener) EnterSpatialIndexOption(ctx *SpatialIndexOptionContext) {}

// ExitSpatialIndexOption is called when production spatialIndexOption is exited.
func (s *BaseTiDBParserListener) ExitSpatialIndexOption(ctx *SpatialIndexOptionContext) {}

// EnterDataTypeDefinition is called when production dataTypeDefinition is entered.
func (s *BaseTiDBParserListener) EnterDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// ExitDataTypeDefinition is called when production dataTypeDefinition is exited.
func (s *BaseTiDBParserListener) ExitDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseTiDBParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseTiDBParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterNchar is called when production nchar is entered.
func (s *BaseTiDBParserListener) EnterNchar(ctx *NcharContext) {}

// ExitNchar is called when production nchar is exited.
func (s *BaseTiDBParserListener) ExitNchar(ctx *NcharContext) {}

// EnterRealType is called when production realType is entered.
func (s *BaseTiDBParserListener) EnterRealType(ctx *RealTypeContext) {}

// ExitRealType is called when production realType is exited.
func (s *BaseTiDBParserListener) ExitRealType(ctx *RealTypeContext) {}

// EnterAutoRandomFieldLength is called when production autoRandomFieldLength is entered.
func (s *BaseTiDBParserListener) EnterAutoRandomFieldLength(ctx *AutoRandomFieldLengthContext) {}

// ExitAutoRandomFieldLength is called when production autoRandomFieldLength is exited.
func (s *BaseTiDBParserListener) ExitAutoRandomFieldLength(ctx *AutoRandomFieldLengthContext) {}

// EnterFieldLength is called when production fieldLength is entered.
func (s *BaseTiDBParserListener) EnterFieldLength(ctx *FieldLengthContext) {}

// ExitFieldLength is called when production fieldLength is exited.
func (s *BaseTiDBParserListener) ExitFieldLength(ctx *FieldLengthContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseTiDBParserListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseTiDBParserListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterCharsetWithOptBinary is called when production charsetWithOptBinary is entered.
func (s *BaseTiDBParserListener) EnterCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) {}

// ExitCharsetWithOptBinary is called when production charsetWithOptBinary is exited.
func (s *BaseTiDBParserListener) ExitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) {}

// EnterAscii is called when production ascii is entered.
func (s *BaseTiDBParserListener) EnterAscii(ctx *AsciiContext) {}

// ExitAscii is called when production ascii is exited.
func (s *BaseTiDBParserListener) ExitAscii(ctx *AsciiContext) {}

// EnterUnicode is called when production unicode is entered.
func (s *BaseTiDBParserListener) EnterUnicode(ctx *UnicodeContext) {}

// ExitUnicode is called when production unicode is exited.
func (s *BaseTiDBParserListener) ExitUnicode(ctx *UnicodeContext) {}

// EnterWsNumCodepoints is called when production wsNumCodepoints is entered.
func (s *BaseTiDBParserListener) EnterWsNumCodepoints(ctx *WsNumCodepointsContext) {}

// ExitWsNumCodepoints is called when production wsNumCodepoints is exited.
func (s *BaseTiDBParserListener) ExitWsNumCodepoints(ctx *WsNumCodepointsContext) {}

// EnterTypeDatetimePrecision is called when production typeDatetimePrecision is entered.
func (s *BaseTiDBParserListener) EnterTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) {}

// ExitTypeDatetimePrecision is called when production typeDatetimePrecision is exited.
func (s *BaseTiDBParserListener) ExitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseTiDBParserListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseTiDBParserListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterCollationName is called when production collationName is entered.
func (s *BaseTiDBParserListener) EnterCollationName(ctx *CollationNameContext) {}

// ExitCollationName is called when production collationName is exited.
func (s *BaseTiDBParserListener) ExitCollationName(ctx *CollationNameContext) {}

// EnterCreateTableOptions is called when production createTableOptions is entered.
func (s *BaseTiDBParserListener) EnterCreateTableOptions(ctx *CreateTableOptionsContext) {}

// ExitCreateTableOptions is called when production createTableOptions is exited.
func (s *BaseTiDBParserListener) ExitCreateTableOptions(ctx *CreateTableOptionsContext) {}

// EnterCreateTableOptionsSpaceSeparated is called when production createTableOptionsSpaceSeparated is entered.
func (s *BaseTiDBParserListener) EnterCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) {
}

// ExitCreateTableOptionsSpaceSeparated is called when production createTableOptionsSpaceSeparated is exited.
func (s *BaseTiDBParserListener) ExitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) {
}

// EnterCreateTableOption is called when production createTableOption is entered.
func (s *BaseTiDBParserListener) EnterCreateTableOption(ctx *CreateTableOptionContext) {}

// ExitCreateTableOption is called when production createTableOption is exited.
func (s *BaseTiDBParserListener) ExitCreateTableOption(ctx *CreateTableOptionContext) {}

// EnterTernaryOption is called when production ternaryOption is entered.
func (s *BaseTiDBParserListener) EnterTernaryOption(ctx *TernaryOptionContext) {}

// ExitTernaryOption is called when production ternaryOption is exited.
func (s *BaseTiDBParserListener) ExitTernaryOption(ctx *TernaryOptionContext) {}

// EnterDefaultCollation is called when production defaultCollation is entered.
func (s *BaseTiDBParserListener) EnterDefaultCollation(ctx *DefaultCollationContext) {}

// ExitDefaultCollation is called when production defaultCollation is exited.
func (s *BaseTiDBParserListener) ExitDefaultCollation(ctx *DefaultCollationContext) {}

// EnterDefaultEncryption is called when production defaultEncryption is entered.
func (s *BaseTiDBParserListener) EnterDefaultEncryption(ctx *DefaultEncryptionContext) {}

// ExitDefaultEncryption is called when production defaultEncryption is exited.
func (s *BaseTiDBParserListener) ExitDefaultEncryption(ctx *DefaultEncryptionContext) {}

// EnterDefaultCharset is called when production defaultCharset is entered.
func (s *BaseTiDBParserListener) EnterDefaultCharset(ctx *DefaultCharsetContext) {}

// ExitDefaultCharset is called when production defaultCharset is exited.
func (s *BaseTiDBParserListener) ExitDefaultCharset(ctx *DefaultCharsetContext) {}

// EnterPartitionClause is called when production partitionClause is entered.
func (s *BaseTiDBParserListener) EnterPartitionClause(ctx *PartitionClauseContext) {}

// ExitPartitionClause is called when production partitionClause is exited.
func (s *BaseTiDBParserListener) ExitPartitionClause(ctx *PartitionClauseContext) {}

// EnterPartitionDefKey is called when production partitionDefKey is entered.
func (s *BaseTiDBParserListener) EnterPartitionDefKey(ctx *PartitionDefKeyContext) {}

// ExitPartitionDefKey is called when production partitionDefKey is exited.
func (s *BaseTiDBParserListener) ExitPartitionDefKey(ctx *PartitionDefKeyContext) {}

// EnterPartitionDefHash is called when production partitionDefHash is entered.
func (s *BaseTiDBParserListener) EnterPartitionDefHash(ctx *PartitionDefHashContext) {}

// ExitPartitionDefHash is called when production partitionDefHash is exited.
func (s *BaseTiDBParserListener) ExitPartitionDefHash(ctx *PartitionDefHashContext) {}

// EnterPartitionDefRangeList is called when production partitionDefRangeList is entered.
func (s *BaseTiDBParserListener) EnterPartitionDefRangeList(ctx *PartitionDefRangeListContext) {}

// ExitPartitionDefRangeList is called when production partitionDefRangeList is exited.
func (s *BaseTiDBParserListener) ExitPartitionDefRangeList(ctx *PartitionDefRangeListContext) {}

// EnterSubPartitions is called when production subPartitions is entered.
func (s *BaseTiDBParserListener) EnterSubPartitions(ctx *SubPartitionsContext) {}

// ExitSubPartitions is called when production subPartitions is exited.
func (s *BaseTiDBParserListener) ExitSubPartitions(ctx *SubPartitionsContext) {}

// EnterPartitionKeyAlgorithm is called when production partitionKeyAlgorithm is entered.
func (s *BaseTiDBParserListener) EnterPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) {}

// ExitPartitionKeyAlgorithm is called when production partitionKeyAlgorithm is exited.
func (s *BaseTiDBParserListener) ExitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) {}

// EnterPartitionDefinitions is called when production partitionDefinitions is entered.
func (s *BaseTiDBParserListener) EnterPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// ExitPartitionDefinitions is called when production partitionDefinitions is exited.
func (s *BaseTiDBParserListener) ExitPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// EnterPartitionDefinition is called when production partitionDefinition is entered.
func (s *BaseTiDBParserListener) EnterPartitionDefinition(ctx *PartitionDefinitionContext) {}

// ExitPartitionDefinition is called when production partitionDefinition is exited.
func (s *BaseTiDBParserListener) ExitPartitionDefinition(ctx *PartitionDefinitionContext) {}

// EnterPartitionValuesIn is called when production partitionValuesIn is entered.
func (s *BaseTiDBParserListener) EnterPartitionValuesIn(ctx *PartitionValuesInContext) {}

// ExitPartitionValuesIn is called when production partitionValuesIn is exited.
func (s *BaseTiDBParserListener) ExitPartitionValuesIn(ctx *PartitionValuesInContext) {}

// EnterPartitionOption is called when production partitionOption is entered.
func (s *BaseTiDBParserListener) EnterPartitionOption(ctx *PartitionOptionContext) {}

// ExitPartitionOption is called when production partitionOption is exited.
func (s *BaseTiDBParserListener) ExitPartitionOption(ctx *PartitionOptionContext) {}

// EnterSubpartitionDefinition is called when production subpartitionDefinition is entered.
func (s *BaseTiDBParserListener) EnterSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {}

// ExitSubpartitionDefinition is called when production subpartitionDefinition is exited.
func (s *BaseTiDBParserListener) ExitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {}

// EnterPartitionValueItemListParen is called when production partitionValueItemListParen is entered.
func (s *BaseTiDBParserListener) EnterPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) {
}

// ExitPartitionValueItemListParen is called when production partitionValueItemListParen is exited.
func (s *BaseTiDBParserListener) ExitPartitionValueItemListParen(ctx *PartitionValueItemListParenContext) {
}

// EnterPartitionValueItem is called when production partitionValueItem is entered.
func (s *BaseTiDBParserListener) EnterPartitionValueItem(ctx *PartitionValueItemContext) {}

// ExitPartitionValueItem is called when production partitionValueItem is exited.
func (s *BaseTiDBParserListener) ExitPartitionValueItem(ctx *PartitionValueItemContext) {}

// EnterDefinerClause is called when production definerClause is entered.
func (s *BaseTiDBParserListener) EnterDefinerClause(ctx *DefinerClauseContext) {}

// ExitDefinerClause is called when production definerClause is exited.
func (s *BaseTiDBParserListener) ExitDefinerClause(ctx *DefinerClauseContext) {}

// EnterIfExists is called when production ifExists is entered.
func (s *BaseTiDBParserListener) EnterIfExists(ctx *IfExistsContext) {}

// ExitIfExists is called when production ifExists is exited.
func (s *BaseTiDBParserListener) ExitIfExists(ctx *IfExistsContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseTiDBParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseTiDBParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterProcedureParameter is called when production procedureParameter is entered.
func (s *BaseTiDBParserListener) EnterProcedureParameter(ctx *ProcedureParameterContext) {}

// ExitProcedureParameter is called when production procedureParameter is exited.
func (s *BaseTiDBParserListener) ExitProcedureParameter(ctx *ProcedureParameterContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BaseTiDBParserListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BaseTiDBParserListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterCollate is called when production collate is entered.
func (s *BaseTiDBParserListener) EnterCollate(ctx *CollateContext) {}

// ExitCollate is called when production collate is exited.
func (s *BaseTiDBParserListener) ExitCollate(ctx *CollateContext) {}

// EnterTypeWithOptCollate is called when production typeWithOptCollate is entered.
func (s *BaseTiDBParserListener) EnterTypeWithOptCollate(ctx *TypeWithOptCollateContext) {}

// ExitTypeWithOptCollate is called when production typeWithOptCollate is exited.
func (s *BaseTiDBParserListener) ExitTypeWithOptCollate(ctx *TypeWithOptCollateContext) {}

// EnterSchemaIdentifierPair is called when production schemaIdentifierPair is entered.
func (s *BaseTiDBParserListener) EnterSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) {}

// ExitSchemaIdentifierPair is called when production schemaIdentifierPair is exited.
func (s *BaseTiDBParserListener) ExitSchemaIdentifierPair(ctx *SchemaIdentifierPairContext) {}

// EnterViewRefList is called when production viewRefList is entered.
func (s *BaseTiDBParserListener) EnterViewRefList(ctx *ViewRefListContext) {}

// ExitViewRefList is called when production viewRefList is exited.
func (s *BaseTiDBParserListener) ExitViewRefList(ctx *ViewRefListContext) {}

// EnterUpdateList is called when production updateList is entered.
func (s *BaseTiDBParserListener) EnterUpdateList(ctx *UpdateListContext) {}

// ExitUpdateList is called when production updateList is exited.
func (s *BaseTiDBParserListener) ExitUpdateList(ctx *UpdateListContext) {}

// EnterUpdateElement is called when production updateElement is entered.
func (s *BaseTiDBParserListener) EnterUpdateElement(ctx *UpdateElementContext) {}

// ExitUpdateElement is called when production updateElement is exited.
func (s *BaseTiDBParserListener) ExitUpdateElement(ctx *UpdateElementContext) {}

// EnterCharsetClause is called when production charsetClause is entered.
func (s *BaseTiDBParserListener) EnterCharsetClause(ctx *CharsetClauseContext) {}

// ExitCharsetClause is called when production charsetClause is exited.
func (s *BaseTiDBParserListener) ExitCharsetClause(ctx *CharsetClauseContext) {}

// EnterFieldsClause is called when production fieldsClause is entered.
func (s *BaseTiDBParserListener) EnterFieldsClause(ctx *FieldsClauseContext) {}

// ExitFieldsClause is called when production fieldsClause is exited.
func (s *BaseTiDBParserListener) ExitFieldsClause(ctx *FieldsClauseContext) {}

// EnterFieldTerm is called when production fieldTerm is entered.
func (s *BaseTiDBParserListener) EnterFieldTerm(ctx *FieldTermContext) {}

// ExitFieldTerm is called when production fieldTerm is exited.
func (s *BaseTiDBParserListener) ExitFieldTerm(ctx *FieldTermContext) {}

// EnterLinesClause is called when production linesClause is entered.
func (s *BaseTiDBParserListener) EnterLinesClause(ctx *LinesClauseContext) {}

// ExitLinesClause is called when production linesClause is exited.
func (s *BaseTiDBParserListener) ExitLinesClause(ctx *LinesClauseContext) {}

// EnterLineTerm is called when production lineTerm is entered.
func (s *BaseTiDBParserListener) EnterLineTerm(ctx *LineTermContext) {}

// ExitLineTerm is called when production lineTerm is exited.
func (s *BaseTiDBParserListener) ExitLineTerm(ctx *LineTermContext) {}

// EnterUserList is called when production userList is entered.
func (s *BaseTiDBParserListener) EnterUserList(ctx *UserListContext) {}

// ExitUserList is called when production userList is exited.
func (s *BaseTiDBParserListener) ExitUserList(ctx *UserListContext) {}

// EnterCreateUserList is called when production createUserList is entered.
func (s *BaseTiDBParserListener) EnterCreateUserList(ctx *CreateUserListContext) {}

// ExitCreateUserList is called when production createUserList is exited.
func (s *BaseTiDBParserListener) ExitCreateUserList(ctx *CreateUserListContext) {}

// EnterAlterUserList is called when production alterUserList is entered.
func (s *BaseTiDBParserListener) EnterAlterUserList(ctx *AlterUserListContext) {}

// ExitAlterUserList is called when production alterUserList is exited.
func (s *BaseTiDBParserListener) ExitAlterUserList(ctx *AlterUserListContext) {}

// EnterCreateUserEntry is called when production createUserEntry is entered.
func (s *BaseTiDBParserListener) EnterCreateUserEntry(ctx *CreateUserEntryContext) {}

// ExitCreateUserEntry is called when production createUserEntry is exited.
func (s *BaseTiDBParserListener) ExitCreateUserEntry(ctx *CreateUserEntryContext) {}

// EnterAlterUserEntry is called when production alterUserEntry is entered.
func (s *BaseTiDBParserListener) EnterAlterUserEntry(ctx *AlterUserEntryContext) {}

// ExitAlterUserEntry is called when production alterUserEntry is exited.
func (s *BaseTiDBParserListener) ExitAlterUserEntry(ctx *AlterUserEntryContext) {}

// EnterRetainCurrentPassword is called when production retainCurrentPassword is entered.
func (s *BaseTiDBParserListener) EnterRetainCurrentPassword(ctx *RetainCurrentPasswordContext) {}

// ExitRetainCurrentPassword is called when production retainCurrentPassword is exited.
func (s *BaseTiDBParserListener) ExitRetainCurrentPassword(ctx *RetainCurrentPasswordContext) {}

// EnterDiscardOldPassword is called when production discardOldPassword is entered.
func (s *BaseTiDBParserListener) EnterDiscardOldPassword(ctx *DiscardOldPasswordContext) {}

// ExitDiscardOldPassword is called when production discardOldPassword is exited.
func (s *BaseTiDBParserListener) ExitDiscardOldPassword(ctx *DiscardOldPasswordContext) {}

// EnterReplacePassword is called when production replacePassword is entered.
func (s *BaseTiDBParserListener) EnterReplacePassword(ctx *ReplacePasswordContext) {}

// ExitReplacePassword is called when production replacePassword is exited.
func (s *BaseTiDBParserListener) ExitReplacePassword(ctx *ReplacePasswordContext) {}

// EnterUserIdentifierOrText is called when production userIdentifierOrText is entered.
func (s *BaseTiDBParserListener) EnterUserIdentifierOrText(ctx *UserIdentifierOrTextContext) {}

// ExitUserIdentifierOrText is called when production userIdentifierOrText is exited.
func (s *BaseTiDBParserListener) ExitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) {}

// EnterUser is called when production user is entered.
func (s *BaseTiDBParserListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseTiDBParserListener) ExitUser(ctx *UserContext) {}

// EnterLikeClause is called when production likeClause is entered.
func (s *BaseTiDBParserListener) EnterLikeClause(ctx *LikeClauseContext) {}

// ExitLikeClause is called when production likeClause is exited.
func (s *BaseTiDBParserListener) ExitLikeClause(ctx *LikeClauseContext) {}

// EnterLikeOrWhere is called when production likeOrWhere is entered.
func (s *BaseTiDBParserListener) EnterLikeOrWhere(ctx *LikeOrWhereContext) {}

// ExitLikeOrWhere is called when production likeOrWhere is exited.
func (s *BaseTiDBParserListener) ExitLikeOrWhere(ctx *LikeOrWhereContext) {}

// EnterOnlineOption is called when production onlineOption is entered.
func (s *BaseTiDBParserListener) EnterOnlineOption(ctx *OnlineOptionContext) {}

// ExitOnlineOption is called when production onlineOption is exited.
func (s *BaseTiDBParserListener) ExitOnlineOption(ctx *OnlineOptionContext) {}

// EnterNoWriteToBinLog is called when production noWriteToBinLog is entered.
func (s *BaseTiDBParserListener) EnterNoWriteToBinLog(ctx *NoWriteToBinLogContext) {}

// ExitNoWriteToBinLog is called when production noWriteToBinLog is exited.
func (s *BaseTiDBParserListener) ExitNoWriteToBinLog(ctx *NoWriteToBinLogContext) {}

// EnterUsePartition is called when production usePartition is entered.
func (s *BaseTiDBParserListener) EnterUsePartition(ctx *UsePartitionContext) {}

// ExitUsePartition is called when production usePartition is exited.
func (s *BaseTiDBParserListener) ExitUsePartition(ctx *UsePartitionContext) {}

// EnterFieldIdentifier is called when production fieldIdentifier is entered.
func (s *BaseTiDBParserListener) EnterFieldIdentifier(ctx *FieldIdentifierContext) {}

// ExitFieldIdentifier is called when production fieldIdentifier is exited.
func (s *BaseTiDBParserListener) ExitFieldIdentifier(ctx *FieldIdentifierContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseTiDBParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseTiDBParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterColumnInternalRef is called when production columnInternalRef is entered.
func (s *BaseTiDBParserListener) EnterColumnInternalRef(ctx *ColumnInternalRefContext) {}

// ExitColumnInternalRef is called when production columnInternalRef is exited.
func (s *BaseTiDBParserListener) ExitColumnInternalRef(ctx *ColumnInternalRefContext) {}

// EnterColumnInternalRefList is called when production columnInternalRefList is entered.
func (s *BaseTiDBParserListener) EnterColumnInternalRefList(ctx *ColumnInternalRefListContext) {}

// ExitColumnInternalRefList is called when production columnInternalRefList is exited.
func (s *BaseTiDBParserListener) ExitColumnInternalRefList(ctx *ColumnInternalRefListContext) {}

// EnterColumnRef is called when production columnRef is entered.
func (s *BaseTiDBParserListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production columnRef is exited.
func (s *BaseTiDBParserListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterInsertIdentifier is called when production insertIdentifier is entered.
func (s *BaseTiDBParserListener) EnterInsertIdentifier(ctx *InsertIdentifierContext) {}

// ExitInsertIdentifier is called when production insertIdentifier is exited.
func (s *BaseTiDBParserListener) ExitInsertIdentifier(ctx *InsertIdentifierContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseTiDBParserListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseTiDBParserListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterIndexRef is called when production indexRef is entered.
func (s *BaseTiDBParserListener) EnterIndexRef(ctx *IndexRefContext) {}

// ExitIndexRef is called when production indexRef is exited.
func (s *BaseTiDBParserListener) ExitIndexRef(ctx *IndexRefContext) {}

// EnterTableWild is called when production tableWild is entered.
func (s *BaseTiDBParserListener) EnterTableWild(ctx *TableWildContext) {}

// ExitTableWild is called when production tableWild is exited.
func (s *BaseTiDBParserListener) ExitTableWild(ctx *TableWildContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseTiDBParserListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseTiDBParserListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterSchemaRef is called when production schemaRef is entered.
func (s *BaseTiDBParserListener) EnterSchemaRef(ctx *SchemaRefContext) {}

// ExitSchemaRef is called when production schemaRef is exited.
func (s *BaseTiDBParserListener) ExitSchemaRef(ctx *SchemaRefContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseTiDBParserListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseTiDBParserListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterProcedureRef is called when production procedureRef is entered.
func (s *BaseTiDBParserListener) EnterProcedureRef(ctx *ProcedureRefContext) {}

// ExitProcedureRef is called when production procedureRef is exited.
func (s *BaseTiDBParserListener) ExitProcedureRef(ctx *ProcedureRefContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseTiDBParserListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseTiDBParserListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterFunctionRef is called when production functionRef is entered.
func (s *BaseTiDBParserListener) EnterFunctionRef(ctx *FunctionRefContext) {}

// ExitFunctionRef is called when production functionRef is exited.
func (s *BaseTiDBParserListener) ExitFunctionRef(ctx *FunctionRefContext) {}

// EnterTriggerName is called when production triggerName is entered.
func (s *BaseTiDBParserListener) EnterTriggerName(ctx *TriggerNameContext) {}

// ExitTriggerName is called when production triggerName is exited.
func (s *BaseTiDBParserListener) ExitTriggerName(ctx *TriggerNameContext) {}

// EnterTriggerRef is called when production triggerRef is entered.
func (s *BaseTiDBParserListener) EnterTriggerRef(ctx *TriggerRefContext) {}

// ExitTriggerRef is called when production triggerRef is exited.
func (s *BaseTiDBParserListener) ExitTriggerRef(ctx *TriggerRefContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseTiDBParserListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseTiDBParserListener) ExitViewName(ctx *ViewNameContext) {}

// EnterViewRef is called when production viewRef is entered.
func (s *BaseTiDBParserListener) EnterViewRef(ctx *ViewRefContext) {}

// ExitViewRef is called when production viewRef is exited.
func (s *BaseTiDBParserListener) ExitViewRef(ctx *ViewRefContext) {}

// EnterTablespaceName is called when production tablespaceName is entered.
func (s *BaseTiDBParserListener) EnterTablespaceName(ctx *TablespaceNameContext) {}

// ExitTablespaceName is called when production tablespaceName is exited.
func (s *BaseTiDBParserListener) ExitTablespaceName(ctx *TablespaceNameContext) {}

// EnterTablespaceRef is called when production tablespaceRef is entered.
func (s *BaseTiDBParserListener) EnterTablespaceRef(ctx *TablespaceRefContext) {}

// ExitTablespaceRef is called when production tablespaceRef is exited.
func (s *BaseTiDBParserListener) ExitTablespaceRef(ctx *TablespaceRefContext) {}

// EnterLogfileGroupName is called when production logfileGroupName is entered.
func (s *BaseTiDBParserListener) EnterLogfileGroupName(ctx *LogfileGroupNameContext) {}

// ExitLogfileGroupName is called when production logfileGroupName is exited.
func (s *BaseTiDBParserListener) ExitLogfileGroupName(ctx *LogfileGroupNameContext) {}

// EnterLogfileGroupRef is called when production logfileGroupRef is entered.
func (s *BaseTiDBParserListener) EnterLogfileGroupRef(ctx *LogfileGroupRefContext) {}

// ExitLogfileGroupRef is called when production logfileGroupRef is exited.
func (s *BaseTiDBParserListener) ExitLogfileGroupRef(ctx *LogfileGroupRefContext) {}

// EnterEventName is called when production eventName is entered.
func (s *BaseTiDBParserListener) EnterEventName(ctx *EventNameContext) {}

// ExitEventName is called when production eventName is exited.
func (s *BaseTiDBParserListener) ExitEventName(ctx *EventNameContext) {}

// EnterEventRef is called when production eventRef is entered.
func (s *BaseTiDBParserListener) EnterEventRef(ctx *EventRefContext) {}

// ExitEventRef is called when production eventRef is exited.
func (s *BaseTiDBParserListener) ExitEventRef(ctx *EventRefContext) {}

// EnterUdfName is called when production udfName is entered.
func (s *BaseTiDBParserListener) EnterUdfName(ctx *UdfNameContext) {}

// ExitUdfName is called when production udfName is exited.
func (s *BaseTiDBParserListener) ExitUdfName(ctx *UdfNameContext) {}

// EnterServerName is called when production serverName is entered.
func (s *BaseTiDBParserListener) EnterServerName(ctx *ServerNameContext) {}

// ExitServerName is called when production serverName is exited.
func (s *BaseTiDBParserListener) ExitServerName(ctx *ServerNameContext) {}

// EnterServerRef is called when production serverRef is entered.
func (s *BaseTiDBParserListener) EnterServerRef(ctx *ServerRefContext) {}

// ExitServerRef is called when production serverRef is exited.
func (s *BaseTiDBParserListener) ExitServerRef(ctx *ServerRefContext) {}

// EnterEngineRef is called when production engineRef is entered.
func (s *BaseTiDBParserListener) EnterEngineRef(ctx *EngineRefContext) {}

// ExitEngineRef is called when production engineRef is exited.
func (s *BaseTiDBParserListener) ExitEngineRef(ctx *EngineRefContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseTiDBParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseTiDBParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterFilterTableRef is called when production filterTableRef is entered.
func (s *BaseTiDBParserListener) EnterFilterTableRef(ctx *FilterTableRefContext) {}

// ExitFilterTableRef is called when production filterTableRef is exited.
func (s *BaseTiDBParserListener) ExitFilterTableRef(ctx *FilterTableRefContext) {}

// EnterTableRefWithWildcard is called when production tableRefWithWildcard is entered.
func (s *BaseTiDBParserListener) EnterTableRefWithWildcard(ctx *TableRefWithWildcardContext) {}

// ExitTableRefWithWildcard is called when production tableRefWithWildcard is exited.
func (s *BaseTiDBParserListener) ExitTableRefWithWildcard(ctx *TableRefWithWildcardContext) {}

// EnterTableRef is called when production tableRef is entered.
func (s *BaseTiDBParserListener) EnterTableRef(ctx *TableRefContext) {}

// ExitTableRef is called when production tableRef is exited.
func (s *BaseTiDBParserListener) ExitTableRef(ctx *TableRefContext) {}

// EnterTableRefList is called when production tableRefList is entered.
func (s *BaseTiDBParserListener) EnterTableRefList(ctx *TableRefListContext) {}

// ExitTableRefList is called when production tableRefList is exited.
func (s *BaseTiDBParserListener) ExitTableRefList(ctx *TableRefListContext) {}

// EnterTableAliasRefList is called when production tableAliasRefList is entered.
func (s *BaseTiDBParserListener) EnterTableAliasRefList(ctx *TableAliasRefListContext) {}

// ExitTableAliasRefList is called when production tableAliasRefList is exited.
func (s *BaseTiDBParserListener) ExitTableAliasRefList(ctx *TableAliasRefListContext) {}

// EnterParameterName is called when production parameterName is entered.
func (s *BaseTiDBParserListener) EnterParameterName(ctx *ParameterNameContext) {}

// ExitParameterName is called when production parameterName is exited.
func (s *BaseTiDBParserListener) ExitParameterName(ctx *ParameterNameContext) {}

// EnterLabelIdentifier is called when production labelIdentifier is entered.
func (s *BaseTiDBParserListener) EnterLabelIdentifier(ctx *LabelIdentifierContext) {}

// ExitLabelIdentifier is called when production labelIdentifier is exited.
func (s *BaseTiDBParserListener) ExitLabelIdentifier(ctx *LabelIdentifierContext) {}

// EnterLabelRef is called when production labelRef is entered.
func (s *BaseTiDBParserListener) EnterLabelRef(ctx *LabelRefContext) {}

// ExitLabelRef is called when production labelRef is exited.
func (s *BaseTiDBParserListener) ExitLabelRef(ctx *LabelRefContext) {}

// EnterRoleIdentifier is called when production roleIdentifier is entered.
func (s *BaseTiDBParserListener) EnterRoleIdentifier(ctx *RoleIdentifierContext) {}

// ExitRoleIdentifier is called when production roleIdentifier is exited.
func (s *BaseTiDBParserListener) ExitRoleIdentifier(ctx *RoleIdentifierContext) {}

// EnterRoleRef is called when production roleRef is entered.
func (s *BaseTiDBParserListener) EnterRoleRef(ctx *RoleRefContext) {}

// ExitRoleRef is called when production roleRef is exited.
func (s *BaseTiDBParserListener) ExitRoleRef(ctx *RoleRefContext) {}

// EnterPluginRef is called when production pluginRef is entered.
func (s *BaseTiDBParserListener) EnterPluginRef(ctx *PluginRefContext) {}

// ExitPluginRef is called when production pluginRef is exited.
func (s *BaseTiDBParserListener) ExitPluginRef(ctx *PluginRefContext) {}

// EnterComponentRef is called when production componentRef is entered.
func (s *BaseTiDBParserListener) EnterComponentRef(ctx *ComponentRefContext) {}

// ExitComponentRef is called when production componentRef is exited.
func (s *BaseTiDBParserListener) ExitComponentRef(ctx *ComponentRefContext) {}

// EnterResourceGroupRef is called when production resourceGroupRef is entered.
func (s *BaseTiDBParserListener) EnterResourceGroupRef(ctx *ResourceGroupRefContext) {}

// ExitResourceGroupRef is called when production resourceGroupRef is exited.
func (s *BaseTiDBParserListener) ExitResourceGroupRef(ctx *ResourceGroupRefContext) {}

// EnterWindowName is called when production windowName is entered.
func (s *BaseTiDBParserListener) EnterWindowName(ctx *WindowNameContext) {}

// ExitWindowName is called when production windowName is exited.
func (s *BaseTiDBParserListener) ExitWindowName(ctx *WindowNameContext) {}

// EnterPureIdentifier is called when production pureIdentifier is entered.
func (s *BaseTiDBParserListener) EnterPureIdentifier(ctx *PureIdentifierContext) {}

// ExitPureIdentifier is called when production pureIdentifier is exited.
func (s *BaseTiDBParserListener) ExitPureIdentifier(ctx *PureIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseTiDBParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseTiDBParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseTiDBParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseTiDBParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierListWithParentheses is called when production identifierListWithParentheses is entered.
func (s *BaseTiDBParserListener) EnterIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) {
}

// ExitIdentifierListWithParentheses is called when production identifierListWithParentheses is exited.
func (s *BaseTiDBParserListener) ExitIdentifierListWithParentheses(ctx *IdentifierListWithParenthesesContext) {
}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseTiDBParserListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseTiDBParserListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterSimpleIdentifier is called when production simpleIdentifier is entered.
func (s *BaseTiDBParserListener) EnterSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// ExitSimpleIdentifier is called when production simpleIdentifier is exited.
func (s *BaseTiDBParserListener) ExitSimpleIdentifier(ctx *SimpleIdentifierContext) {}

// EnterDotIdentifier is called when production dotIdentifier is entered.
func (s *BaseTiDBParserListener) EnterDotIdentifier(ctx *DotIdentifierContext) {}

// ExitDotIdentifier is called when production dotIdentifier is exited.
func (s *BaseTiDBParserListener) ExitDotIdentifier(ctx *DotIdentifierContext) {}

// EnterUlong_number is called when production ulong_number is entered.
func (s *BaseTiDBParserListener) EnterUlong_number(ctx *Ulong_numberContext) {}

// ExitUlong_number is called when production ulong_number is exited.
func (s *BaseTiDBParserListener) ExitUlong_number(ctx *Ulong_numberContext) {}

// EnterReal_ulong_number is called when production real_ulong_number is entered.
func (s *BaseTiDBParserListener) EnterReal_ulong_number(ctx *Real_ulong_numberContext) {}

// ExitReal_ulong_number is called when production real_ulong_number is exited.
func (s *BaseTiDBParserListener) ExitReal_ulong_number(ctx *Real_ulong_numberContext) {}

// EnterUlonglong_number is called when production ulonglong_number is entered.
func (s *BaseTiDBParserListener) EnterUlonglong_number(ctx *Ulonglong_numberContext) {}

// ExitUlonglong_number is called when production ulonglong_number is exited.
func (s *BaseTiDBParserListener) ExitUlonglong_number(ctx *Ulonglong_numberContext) {}

// EnterReal_ulonglong_number is called when production real_ulonglong_number is entered.
func (s *BaseTiDBParserListener) EnterReal_ulonglong_number(ctx *Real_ulonglong_numberContext) {}

// ExitReal_ulonglong_number is called when production real_ulonglong_number is exited.
func (s *BaseTiDBParserListener) ExitReal_ulonglong_number(ctx *Real_ulonglong_numberContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseTiDBParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseTiDBParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterSignedLiteral is called when production signedLiteral is entered.
func (s *BaseTiDBParserListener) EnterSignedLiteral(ctx *SignedLiteralContext) {}

// ExitSignedLiteral is called when production signedLiteral is exited.
func (s *BaseTiDBParserListener) ExitSignedLiteral(ctx *SignedLiteralContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseTiDBParserListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseTiDBParserListener) ExitStringList(ctx *StringListContext) {}

// EnterTextStringLiteral is called when production textStringLiteral is entered.
func (s *BaseTiDBParserListener) EnterTextStringLiteral(ctx *TextStringLiteralContext) {}

// ExitTextStringLiteral is called when production textStringLiteral is exited.
func (s *BaseTiDBParserListener) ExitTextStringLiteral(ctx *TextStringLiteralContext) {}

// EnterTextString is called when production textString is entered.
func (s *BaseTiDBParserListener) EnterTextString(ctx *TextStringContext) {}

// ExitTextString is called when production textString is exited.
func (s *BaseTiDBParserListener) ExitTextString(ctx *TextStringContext) {}

// EnterTextStringHash is called when production textStringHash is entered.
func (s *BaseTiDBParserListener) EnterTextStringHash(ctx *TextStringHashContext) {}

// ExitTextStringHash is called when production textStringHash is exited.
func (s *BaseTiDBParserListener) ExitTextStringHash(ctx *TextStringHashContext) {}

// EnterTextLiteral is called when production textLiteral is entered.
func (s *BaseTiDBParserListener) EnterTextLiteral(ctx *TextLiteralContext) {}

// ExitTextLiteral is called when production textLiteral is exited.
func (s *BaseTiDBParserListener) ExitTextLiteral(ctx *TextLiteralContext) {}

// EnterTextStringNoLinebreak is called when production textStringNoLinebreak is entered.
func (s *BaseTiDBParserListener) EnterTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) {}

// ExitTextStringNoLinebreak is called when production textStringNoLinebreak is exited.
func (s *BaseTiDBParserListener) ExitTextStringNoLinebreak(ctx *TextStringNoLinebreakContext) {}

// EnterTextStringLiteralList is called when production textStringLiteralList is entered.
func (s *BaseTiDBParserListener) EnterTextStringLiteralList(ctx *TextStringLiteralListContext) {}

// ExitTextStringLiteralList is called when production textStringLiteralList is exited.
func (s *BaseTiDBParserListener) ExitTextStringLiteralList(ctx *TextStringLiteralListContext) {}

// EnterNumLiteral is called when production numLiteral is entered.
func (s *BaseTiDBParserListener) EnterNumLiteral(ctx *NumLiteralContext) {}

// ExitNumLiteral is called when production numLiteral is exited.
func (s *BaseTiDBParserListener) ExitNumLiteral(ctx *NumLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BaseTiDBParserListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BaseTiDBParserListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseTiDBParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseTiDBParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterTemporalLiteral is called when production temporalLiteral is entered.
func (s *BaseTiDBParserListener) EnterTemporalLiteral(ctx *TemporalLiteralContext) {}

// ExitTemporalLiteral is called when production temporalLiteral is exited.
func (s *BaseTiDBParserListener) ExitTemporalLiteral(ctx *TemporalLiteralContext) {}

// EnterFloatOptions is called when production floatOptions is entered.
func (s *BaseTiDBParserListener) EnterFloatOptions(ctx *FloatOptionsContext) {}

// ExitFloatOptions is called when production floatOptions is exited.
func (s *BaseTiDBParserListener) ExitFloatOptions(ctx *FloatOptionsContext) {}

// EnterStandardFloatOptions is called when production standardFloatOptions is entered.
func (s *BaseTiDBParserListener) EnterStandardFloatOptions(ctx *StandardFloatOptionsContext) {}

// ExitStandardFloatOptions is called when production standardFloatOptions is exited.
func (s *BaseTiDBParserListener) ExitStandardFloatOptions(ctx *StandardFloatOptionsContext) {}

// EnterPrecision is called when production precision is entered.
func (s *BaseTiDBParserListener) EnterPrecision(ctx *PrecisionContext) {}

// ExitPrecision is called when production precision is exited.
func (s *BaseTiDBParserListener) ExitPrecision(ctx *PrecisionContext) {}

// EnterTextOrIdentifier is called when production textOrIdentifier is entered.
func (s *BaseTiDBParserListener) EnterTextOrIdentifier(ctx *TextOrIdentifierContext) {}

// ExitTextOrIdentifier is called when production textOrIdentifier is exited.
func (s *BaseTiDBParserListener) ExitTextOrIdentifier(ctx *TextOrIdentifierContext) {}

// EnterLValueIdentifier is called when production lValueIdentifier is entered.
func (s *BaseTiDBParserListener) EnterLValueIdentifier(ctx *LValueIdentifierContext) {}

// ExitLValueIdentifier is called when production lValueIdentifier is exited.
func (s *BaseTiDBParserListener) ExitLValueIdentifier(ctx *LValueIdentifierContext) {}

// EnterRoleIdentifierOrText is called when production roleIdentifierOrText is entered.
func (s *BaseTiDBParserListener) EnterRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) {}

// ExitRoleIdentifierOrText is called when production roleIdentifierOrText is exited.
func (s *BaseTiDBParserListener) ExitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) {}

// EnterSizeNumber is called when production sizeNumber is entered.
func (s *BaseTiDBParserListener) EnterSizeNumber(ctx *SizeNumberContext) {}

// ExitSizeNumber is called when production sizeNumber is exited.
func (s *BaseTiDBParserListener) ExitSizeNumber(ctx *SizeNumberContext) {}

// EnterParentheses is called when production parentheses is entered.
func (s *BaseTiDBParserListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production parentheses is exited.
func (s *BaseTiDBParserListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseTiDBParserListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseTiDBParserListener) ExitEqual(ctx *EqualContext) {}

// EnterOptionType is called when production optionType is entered.
func (s *BaseTiDBParserListener) EnterOptionType(ctx *OptionTypeContext) {}

// ExitOptionType is called when production optionType is exited.
func (s *BaseTiDBParserListener) ExitOptionType(ctx *OptionTypeContext) {}

// EnterVarIdentType is called when production varIdentType is entered.
func (s *BaseTiDBParserListener) EnterVarIdentType(ctx *VarIdentTypeContext) {}

// ExitVarIdentType is called when production varIdentType is exited.
func (s *BaseTiDBParserListener) ExitVarIdentType(ctx *VarIdentTypeContext) {}

// EnterSetVarIdentType is called when production setVarIdentType is entered.
func (s *BaseTiDBParserListener) EnterSetVarIdentType(ctx *SetVarIdentTypeContext) {}

// ExitSetVarIdentType is called when production setVarIdentType is exited.
func (s *BaseTiDBParserListener) ExitSetVarIdentType(ctx *SetVarIdentTypeContext) {}

// EnterIdentifierKeyword is called when production identifierKeyword is entered.
func (s *BaseTiDBParserListener) EnterIdentifierKeyword(ctx *IdentifierKeywordContext) {}

// ExitIdentifierKeyword is called when production identifierKeyword is exited.
func (s *BaseTiDBParserListener) ExitIdentifierKeyword(ctx *IdentifierKeywordContext) {}

// EnterIdentifierKeywordsAmbiguous1RolesAndLabels is called when production identifierKeywordsAmbiguous1RolesAndLabels is entered.
func (s *BaseTiDBParserListener) EnterIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) {
}

// ExitIdentifierKeywordsAmbiguous1RolesAndLabels is called when production identifierKeywordsAmbiguous1RolesAndLabels is exited.
func (s *BaseTiDBParserListener) ExitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) {
}

// EnterIdentifierKeywordsAmbiguous2Labels is called when production identifierKeywordsAmbiguous2Labels is entered.
func (s *BaseTiDBParserListener) EnterIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) {
}

// ExitIdentifierKeywordsAmbiguous2Labels is called when production identifierKeywordsAmbiguous2Labels is exited.
func (s *BaseTiDBParserListener) ExitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) {
}

// EnterLabelKeyword is called when production labelKeyword is entered.
func (s *BaseTiDBParserListener) EnterLabelKeyword(ctx *LabelKeywordContext) {}

// ExitLabelKeyword is called when production labelKeyword is exited.
func (s *BaseTiDBParserListener) ExitLabelKeyword(ctx *LabelKeywordContext) {}

// EnterIdentifierKeywordsAmbiguous3Roles is called when production identifierKeywordsAmbiguous3Roles is entered.
func (s *BaseTiDBParserListener) EnterIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) {
}

// ExitIdentifierKeywordsAmbiguous3Roles is called when production identifierKeywordsAmbiguous3Roles is exited.
func (s *BaseTiDBParserListener) ExitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) {
}

// EnterIdentifierKeywordsUnambiguous is called when production identifierKeywordsUnambiguous is entered.
func (s *BaseTiDBParserListener) EnterIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) {
}

// ExitIdentifierKeywordsUnambiguous is called when production identifierKeywordsUnambiguous is exited.
func (s *BaseTiDBParserListener) ExitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) {
}

// EnterRoleKeyword is called when production roleKeyword is entered.
func (s *BaseTiDBParserListener) EnterRoleKeyword(ctx *RoleKeywordContext) {}

// ExitRoleKeyword is called when production roleKeyword is exited.
func (s *BaseTiDBParserListener) ExitRoleKeyword(ctx *RoleKeywordContext) {}

// EnterLValueKeyword is called when production lValueKeyword is entered.
func (s *BaseTiDBParserListener) EnterLValueKeyword(ctx *LValueKeywordContext) {}

// ExitLValueKeyword is called when production lValueKeyword is exited.
func (s *BaseTiDBParserListener) ExitLValueKeyword(ctx *LValueKeywordContext) {}

// EnterIdentifierKeywordsAmbiguous4SystemVariables is called when production identifierKeywordsAmbiguous4SystemVariables is entered.
func (s *BaseTiDBParserListener) EnterIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) {
}

// ExitIdentifierKeywordsAmbiguous4SystemVariables is called when production identifierKeywordsAmbiguous4SystemVariables is exited.
func (s *BaseTiDBParserListener) ExitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) {
}

// EnterRoleOrIdentifierKeyword is called when production roleOrIdentifierKeyword is entered.
func (s *BaseTiDBParserListener) EnterRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) {}

// ExitRoleOrIdentifierKeyword is called when production roleOrIdentifierKeyword is exited.
func (s *BaseTiDBParserListener) ExitRoleOrIdentifierKeyword(ctx *RoleOrIdentifierKeywordContext) {}

// EnterRoleOrLabelKeyword is called when production roleOrLabelKeyword is entered.
func (s *BaseTiDBParserListener) EnterRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) {}

// ExitRoleOrLabelKeyword is called when production roleOrLabelKeyword is exited.
func (s *BaseTiDBParserListener) ExitRoleOrLabelKeyword(ctx *RoleOrLabelKeywordContext) {}
