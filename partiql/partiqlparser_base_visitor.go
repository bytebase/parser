// Code generated from PartiQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package partiql // PartiQLParser
import "github.com/antlr4-go/antlr/v4"

type BasePartiQLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePartiQLParserVisitor) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitQueryDql(ctx *QueryDqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitQueryDml(ctx *QueryDmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitQueryDdl(ctx *QueryDdlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitQueryExec(ctx *QueryExecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExplainOption(ctx *ExplainOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitAsIdent(ctx *AsIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitAtIdent(ctx *AtIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitByIdent(ctx *ByIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSymbolPrimitive(ctx *SymbolPrimitiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDql(ctx *DqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExecCommand(ctx *ExecCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDdl(ctx *DdlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCreateIndex(ctx *CreateIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDropIndex(ctx *DropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDmlBaseWrapper(ctx *DmlBaseWrapperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDmlDelete(ctx *DmlDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDmlInsertReturning(ctx *DmlInsertReturningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDmlBase(ctx *DmlBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDmlBaseCommand(ctx *DmlBaseCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathSimple(ctx *PathSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathSimpleLiteral(ctx *PathSimpleLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathSimpleSymbol(ctx *PathSimpleSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathSimpleDotSymbol(ctx *PathSimpleDotSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitReplaceCommand(ctx *ReplaceCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitUpsertCommand(ctx *UpsertCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitRemoveCommand(ctx *RemoveCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitInsertCommandReturning(ctx *InsertCommandReturningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitInsertLegacy(ctx *InsertLegacyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitInsert(ctx *InsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitOnConflictLegacy(ctx *OnConflictLegacyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitOnConflict(ctx *OnConflictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitConflictTarget(ctx *ConflictTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitConflictAction(ctx *ConflictActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDoReplace(ctx *DoReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDoUpdate(ctx *DoUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitUpdateClause(ctx *UpdateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSetCommand(ctx *SetCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSetAssignment(ctx *SetAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDeleteCommand(ctx *DeleteCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitReturningClause(ctx *ReturningClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitReturningColumn(ctx *ReturningColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitFromClauseSimpleExplicit(ctx *FromClauseSimpleExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitFromClauseSimpleImplicit(ctx *FromClauseSimpleImplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSelectAll(ctx *SelectAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSelectItems(ctx *SelectItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSelectValue(ctx *SelectValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSelectPivot(ctx *SelectPivotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitProjectionItems(ctx *ProjectionItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitProjectionItem(ctx *ProjectionItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSetQuantifierStrategy(ctx *SetQuantifierStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLetClause(ctx *LetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLetBinding(ctx *LetBindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitOrderSortSpec(ctx *OrderSortSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitGroupClause(ctx *GroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitGroupAlias(ctx *GroupAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitGroupKey(ctx *GroupKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitWindowPartitionList(ctx *WindowPartitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitWindowSortSpecList(ctx *WindowSortSpecListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitWhereClauseSelect(ctx *WhereClauseSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitOffsetByClause(ctx *OffsetByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitGpmlPattern(ctx *GpmlPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitGpmlPatternList(ctx *GpmlPatternListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitMatchPattern(ctx *MatchPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitGraphPart(ctx *GraphPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSelectorBasic(ctx *SelectorBasicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSelectorAny(ctx *SelectorAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSelectorShortest(ctx *SelectorShortestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPatternPathVariable(ctx *PatternPathVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPatternRestrictor(ctx *PatternRestrictorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitNode(ctx *NodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeWithSpec(ctx *EdgeWithSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeAbbreviated(ctx *EdgeAbbreviatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPatternQuantifier(ctx *PatternQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpecRight(ctx *EdgeSpecRightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpecUndirected(ctx *EdgeSpecUndirectedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpecLeft(ctx *EdgeSpecLeftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpecUndirectedRight(ctx *EdgeSpecUndirectedRightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpecUndirectedLeft(ctx *EdgeSpecUndirectedLeftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpecBidirectional(ctx *EdgeSpecBidirectionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpecUndirectedBidirectional(ctx *EdgeSpecUndirectedBidirectionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeSpec(ctx *EdgeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPatternPartLabel(ctx *PatternPartLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitEdgeAbbrev(ctx *EdgeAbbrevContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableWrapped(ctx *TableWrappedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableCrossJoin(ctx *TableCrossJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableQualifiedJoin(ctx *TableQualifiedJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableRefBase(ctx *TableRefBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableNonJoin(ctx *TableNonJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableBaseRefSymbol(ctx *TableBaseRefSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableBaseRefClauses(ctx *TableBaseRefClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableBaseRefMatch(ctx *TableBaseRefMatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTableUnpivot(ctx *TableUnpivotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitJoinRhsBase(ctx *JoinRhsBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitJoinRhsTableJoined(ctx *JoinRhsTableJoinedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitJoinSpec(ctx *JoinSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitJoinType(ctx *JoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitIntersect(ctx *IntersectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitQueryBase(ctx *QueryBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExcept(ctx *ExceptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitUnion(ctx *UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSfwQuery(ctx *SfwQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSfwBase(ctx *SfwBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprOrBase(ctx *ExprOrBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprAndBase(ctx *ExprAndBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprNotBase(ctx *ExprNotBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPredicateIn(ctx *PredicateInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPredicateBetween(ctx *PredicateBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPredicateBase(ctx *PredicateBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPredicateComparison(ctx *PredicateComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPredicateIs(ctx *PredicateIsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPredicateLike(ctx *PredicateLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitMathOp00(ctx *MathOp00Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitMathOp01(ctx *MathOp01Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitMathOp02(ctx *MathOp02Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitValueExpr(ctx *ValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprPrimaryPath(ctx *ExprPrimaryPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprPrimaryBase(ctx *ExprPrimaryBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprTermWrappedQuery(ctx *ExprTermWrappedQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprTermBase(ctx *ExprTermBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitNullIf(ctx *NullIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCoalesce(ctx *CoalesceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCaseExpr(ctx *CaseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitValues(ctx *ValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitValueRow(ctx *ValueRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitValueList(ctx *ValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSequenceConstructor(ctx *SequenceConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitSubstring(ctx *SubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCountAll(ctx *CountAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitAggregateBase(ctx *AggregateBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLagLeadFunction(ctx *LagLeadFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCanLosslessCast(ctx *CanLosslessCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCanCast(ctx *CanCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTrimFunction(ctx *TrimFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitDateFunction(ctx *DateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitFunctionCallReserved(ctx *FunctionCallReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitFunctionCallIdent(ctx *FunctionCallIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathStepIndexExpr(ctx *PathStepIndexExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathStepIndexAll(ctx *PathStepIndexAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathStepDotExpr(ctx *PathStepDotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPathStepDotAll(ctx *PathStepDotAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprGraphMatchMany(ctx *ExprGraphMatchManyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitExprGraphMatchOne(ctx *ExprGraphMatchOneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitVarRefExpr(ctx *VarRefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitCollection(ctx *CollectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitBag(ctx *BagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralNull(ctx *LiteralNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralMissing(ctx *LiteralMissingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralTrue(ctx *LiteralTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralFalse(ctx *LiteralFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralString(ctx *LiteralStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralInteger(ctx *LiteralIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralDecimal(ctx *LiteralDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralIon(ctx *LiteralIonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralDate(ctx *LiteralDateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitLiteralTime(ctx *LiteralTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTypeAtomic(ctx *TypeAtomicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTypeArgSingle(ctx *TypeArgSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTypeVarChar(ctx *TypeVarCharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTypeArgDouble(ctx *TypeArgDoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTypeTimeZone(ctx *TypeTimeZoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePartiQLParserVisitor) VisitTypeCustom(ctx *TypeCustomContext) interface{} {
	return v.VisitChildren(ctx)
}
