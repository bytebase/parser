// Code generated from PartiQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package partiql // PartiQLParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PartiQLParserParser.
type PartiQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PartiQLParserParser#script.
	VisitScript(ctx *ScriptContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#QueryDql.
	VisitQueryDql(ctx *QueryDqlContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#QueryDml.
	VisitQueryDml(ctx *QueryDmlContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#QueryDdl.
	VisitQueryDdl(ctx *QueryDdlContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#QueryExec.
	VisitQueryExec(ctx *QueryExecContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#explainOption.
	VisitExplainOption(ctx *ExplainOptionContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#asIdent.
	VisitAsIdent(ctx *AsIdentContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#atIdent.
	VisitAtIdent(ctx *AtIdentContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#byIdent.
	VisitByIdent(ctx *ByIdentContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#symbolPrimitive.
	VisitSymbolPrimitive(ctx *SymbolPrimitiveContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#dql.
	VisitDql(ctx *DqlContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#execCommand.
	VisitExecCommand(ctx *ExecCommandContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ddl.
	VisitDdl(ctx *DdlContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#CreateTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#CreateIndex.
	VisitCreateIndex(ctx *CreateIndexContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#DropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#DropIndex.
	VisitDropIndex(ctx *DropIndexContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#DmlBaseWrapper.
	VisitDmlBaseWrapper(ctx *DmlBaseWrapperContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#DmlDelete.
	VisitDmlDelete(ctx *DmlDeleteContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#DmlInsertReturning.
	VisitDmlInsertReturning(ctx *DmlInsertReturningContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#DmlBase.
	VisitDmlBase(ctx *DmlBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#dmlBaseCommand.
	VisitDmlBaseCommand(ctx *DmlBaseCommandContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#pathSimple.
	VisitPathSimple(ctx *PathSimpleContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PathSimpleLiteral.
	VisitPathSimpleLiteral(ctx *PathSimpleLiteralContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PathSimpleSymbol.
	VisitPathSimpleSymbol(ctx *PathSimpleSymbolContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PathSimpleDotSymbol.
	VisitPathSimpleDotSymbol(ctx *PathSimpleDotSymbolContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#replaceCommand.
	VisitReplaceCommand(ctx *ReplaceCommandContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#upsertCommand.
	VisitUpsertCommand(ctx *UpsertCommandContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#removeCommand.
	VisitRemoveCommand(ctx *RemoveCommandContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#insertCommandReturning.
	VisitInsertCommandReturning(ctx *InsertCommandReturningContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#InsertLegacy.
	VisitInsertLegacy(ctx *InsertLegacyContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#Insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#OnConflictLegacy.
	VisitOnConflictLegacy(ctx *OnConflictLegacyContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#OnConflict.
	VisitOnConflict(ctx *OnConflictContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#conflictTarget.
	VisitConflictTarget(ctx *ConflictTargetContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#conflictAction.
	VisitConflictAction(ctx *ConflictActionContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#doReplace.
	VisitDoReplace(ctx *DoReplaceContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#doUpdate.
	VisitDoUpdate(ctx *DoUpdateContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#updateClause.
	VisitUpdateClause(ctx *UpdateClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#setCommand.
	VisitSetCommand(ctx *SetCommandContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#setAssignment.
	VisitSetAssignment(ctx *SetAssignmentContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#deleteCommand.
	VisitDeleteCommand(ctx *DeleteCommandContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#returningClause.
	VisitReturningClause(ctx *ReturningClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#returningColumn.
	VisitReturningColumn(ctx *ReturningColumnContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#FromClauseSimpleExplicit.
	VisitFromClauseSimpleExplicit(ctx *FromClauseSimpleExplicitContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#FromClauseSimpleImplicit.
	VisitFromClauseSimpleImplicit(ctx *FromClauseSimpleImplicitContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SelectAll.
	VisitSelectAll(ctx *SelectAllContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SelectItems.
	VisitSelectItems(ctx *SelectItemsContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SelectValue.
	VisitSelectValue(ctx *SelectValueContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SelectPivot.
	VisitSelectPivot(ctx *SelectPivotContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#projectionItems.
	VisitProjectionItems(ctx *ProjectionItemsContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#projectionItem.
	VisitProjectionItem(ctx *ProjectionItemContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#setQuantifierStrategy.
	VisitSetQuantifierStrategy(ctx *SetQuantifierStrategyContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#letClause.
	VisitLetClause(ctx *LetClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#letBinding.
	VisitLetBinding(ctx *LetBindingContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#orderSortSpec.
	VisitOrderSortSpec(ctx *OrderSortSpecContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#groupClause.
	VisitGroupClause(ctx *GroupClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#groupAlias.
	VisitGroupAlias(ctx *GroupAliasContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#groupKey.
	VisitGroupKey(ctx *GroupKeyContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#over.
	VisitOver(ctx *OverContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#windowPartitionList.
	VisitWindowPartitionList(ctx *WindowPartitionListContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#windowSortSpecList.
	VisitWindowSortSpecList(ctx *WindowSortSpecListContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#whereClauseSelect.
	VisitWhereClauseSelect(ctx *WhereClauseSelectContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#offsetByClause.
	VisitOffsetByClause(ctx *OffsetByClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#gpmlPattern.
	VisitGpmlPattern(ctx *GpmlPatternContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#gpmlPatternList.
	VisitGpmlPatternList(ctx *GpmlPatternListContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#matchPattern.
	VisitMatchPattern(ctx *MatchPatternContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#graphPart.
	VisitGraphPart(ctx *GraphPartContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SelectorBasic.
	VisitSelectorBasic(ctx *SelectorBasicContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SelectorAny.
	VisitSelectorAny(ctx *SelectorAnyContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SelectorShortest.
	VisitSelectorShortest(ctx *SelectorShortestContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#patternPathVariable.
	VisitPatternPathVariable(ctx *PatternPathVariableContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#patternRestrictor.
	VisitPatternRestrictor(ctx *PatternRestrictorContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#node.
	VisitNode(ctx *NodeContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeWithSpec.
	VisitEdgeWithSpec(ctx *EdgeWithSpecContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeAbbreviated.
	VisitEdgeAbbreviated(ctx *EdgeAbbreviatedContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#patternQuantifier.
	VisitPatternQuantifier(ctx *PatternQuantifierContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeSpecRight.
	VisitEdgeSpecRight(ctx *EdgeSpecRightContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeSpecUndirected.
	VisitEdgeSpecUndirected(ctx *EdgeSpecUndirectedContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeSpecLeft.
	VisitEdgeSpecLeft(ctx *EdgeSpecLeftContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeSpecUndirectedRight.
	VisitEdgeSpecUndirectedRight(ctx *EdgeSpecUndirectedRightContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeSpecUndirectedLeft.
	VisitEdgeSpecUndirectedLeft(ctx *EdgeSpecUndirectedLeftContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeSpecBidirectional.
	VisitEdgeSpecBidirectional(ctx *EdgeSpecBidirectionalContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#EdgeSpecUndirectedBidirectional.
	VisitEdgeSpecUndirectedBidirectional(ctx *EdgeSpecUndirectedBidirectionalContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#edgeSpec.
	VisitEdgeSpec(ctx *EdgeSpecContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#patternPartLabel.
	VisitPatternPartLabel(ctx *PatternPartLabelContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#edgeAbbrev.
	VisitEdgeAbbrev(ctx *EdgeAbbrevContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TableWrapped.
	VisitTableWrapped(ctx *TableWrappedContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TableCrossJoin.
	VisitTableCrossJoin(ctx *TableCrossJoinContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TableQualifiedJoin.
	VisitTableQualifiedJoin(ctx *TableQualifiedJoinContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TableRefBase.
	VisitTableRefBase(ctx *TableRefBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#tableNonJoin.
	VisitTableNonJoin(ctx *TableNonJoinContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TableBaseRefSymbol.
	VisitTableBaseRefSymbol(ctx *TableBaseRefSymbolContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TableBaseRefClauses.
	VisitTableBaseRefClauses(ctx *TableBaseRefClausesContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TableBaseRefMatch.
	VisitTableBaseRefMatch(ctx *TableBaseRefMatchContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#tableUnpivot.
	VisitTableUnpivot(ctx *TableUnpivotContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#JoinRhsBase.
	VisitJoinRhsBase(ctx *JoinRhsBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#JoinRhsTableJoined.
	VisitJoinRhsTableJoined(ctx *JoinRhsTableJoinedContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#joinSpec.
	VisitJoinSpec(ctx *JoinSpecContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#joinType.
	VisitJoinType(ctx *JoinTypeContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#Intersect.
	VisitIntersect(ctx *IntersectContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#QueryBase.
	VisitQueryBase(ctx *QueryBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#Except.
	VisitExcept(ctx *ExceptContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#Union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SfwQuery.
	VisitSfwQuery(ctx *SfwQueryContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#SfwBase.
	VisitSfwBase(ctx *SfwBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ExprOrBase.
	VisitExprOrBase(ctx *ExprOrBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ExprAndBase.
	VisitExprAndBase(ctx *ExprAndBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ExprNotBase.
	VisitExprNotBase(ctx *ExprNotBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PredicateIn.
	VisitPredicateIn(ctx *PredicateInContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PredicateBetween.
	VisitPredicateBetween(ctx *PredicateBetweenContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PredicateBase.
	VisitPredicateBase(ctx *PredicateBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PredicateComparison.
	VisitPredicateComparison(ctx *PredicateComparisonContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PredicateIs.
	VisitPredicateIs(ctx *PredicateIsContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PredicateLike.
	VisitPredicateLike(ctx *PredicateLikeContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#mathOp00.
	VisitMathOp00(ctx *MathOp00Context) interface{}

	// Visit a parse tree produced by PartiQLParserParser#mathOp01.
	VisitMathOp01(ctx *MathOp01Context) interface{}

	// Visit a parse tree produced by PartiQLParserParser#mathOp02.
	VisitMathOp02(ctx *MathOp02Context) interface{}

	// Visit a parse tree produced by PartiQLParserParser#valueExpr.
	VisitValueExpr(ctx *ValueExprContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ExprPrimaryPath.
	VisitExprPrimaryPath(ctx *ExprPrimaryPathContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ExprPrimaryBase.
	VisitExprPrimaryBase(ctx *ExprPrimaryBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ExprTermWrappedQuery.
	VisitExprTermWrappedQuery(ctx *ExprTermWrappedQueryContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#ExprTermBase.
	VisitExprTermBase(ctx *ExprTermBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#nullIf.
	VisitNullIf(ctx *NullIfContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#coalesce.
	VisitCoalesce(ctx *CoalesceContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#caseExpr.
	VisitCaseExpr(ctx *CaseExprContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#values.
	VisitValues(ctx *ValuesContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#valueRow.
	VisitValueRow(ctx *ValueRowContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#sequenceConstructor.
	VisitSequenceConstructor(ctx *SequenceConstructorContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#substring.
	VisitSubstring(ctx *SubstringContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#CountAll.
	VisitCountAll(ctx *CountAllContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#AggregateBase.
	VisitAggregateBase(ctx *AggregateBaseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LagLeadFunction.
	VisitLagLeadFunction(ctx *LagLeadFunctionContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#cast.
	VisitCast(ctx *CastContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#canLosslessCast.
	VisitCanLosslessCast(ctx *CanLosslessCastContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#canCast.
	VisitCanCast(ctx *CanCastContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#extract.
	VisitExtract(ctx *ExtractContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#trimFunction.
	VisitTrimFunction(ctx *TrimFunctionContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#dateFunction.
	VisitDateFunction(ctx *DateFunctionContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#FunctionCallReserved.
	VisitFunctionCallReserved(ctx *FunctionCallReservedContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#FunctionCallIdent.
	VisitFunctionCallIdent(ctx *FunctionCallIdentContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PathStepIndexExpr.
	VisitPathStepIndexExpr(ctx *PathStepIndexExprContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PathStepIndexAll.
	VisitPathStepIndexAll(ctx *PathStepIndexAllContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PathStepDotExpr.
	VisitPathStepDotExpr(ctx *PathStepDotExprContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#PathStepDotAll.
	VisitPathStepDotAll(ctx *PathStepDotAllContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#exprGraphMatchMany.
	VisitExprGraphMatchMany(ctx *ExprGraphMatchManyContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#exprGraphMatchOne.
	VisitExprGraphMatchOne(ctx *ExprGraphMatchOneContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#varRefExpr.
	VisitVarRefExpr(ctx *VarRefExprContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#collection.
	VisitCollection(ctx *CollectionContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#bag.
	VisitBag(ctx *BagContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralNull.
	VisitLiteralNull(ctx *LiteralNullContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralMissing.
	VisitLiteralMissing(ctx *LiteralMissingContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralTrue.
	VisitLiteralTrue(ctx *LiteralTrueContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralFalse.
	VisitLiteralFalse(ctx *LiteralFalseContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralString.
	VisitLiteralString(ctx *LiteralStringContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralInteger.
	VisitLiteralInteger(ctx *LiteralIntegerContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralDecimal.
	VisitLiteralDecimal(ctx *LiteralDecimalContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralIon.
	VisitLiteralIon(ctx *LiteralIonContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralDate.
	VisitLiteralDate(ctx *LiteralDateContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#LiteralTime.
	VisitLiteralTime(ctx *LiteralTimeContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TypeAtomic.
	VisitTypeAtomic(ctx *TypeAtomicContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TypeArgSingle.
	VisitTypeArgSingle(ctx *TypeArgSingleContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TypeVarChar.
	VisitTypeVarChar(ctx *TypeVarCharContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TypeArgDouble.
	VisitTypeArgDouble(ctx *TypeArgDoubleContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TypeTimeZone.
	VisitTypeTimeZone(ctx *TypeTimeZoneContext) interface{}

	// Visit a parse tree produced by PartiQLParserParser#TypeCustom.
	VisitTypeCustom(ctx *TypeCustomContext) interface{}
}
