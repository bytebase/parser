// Code generated from PartiQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package partiql // PartiQLParser
import "github.com/antlr4-go/antlr/v4"

// PartiQLParserListener is a complete listener for a parse tree produced by PartiQLParserParser.
type PartiQLParserListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterQueryDql is called when entering the QueryDql production.
	EnterQueryDql(c *QueryDqlContext)

	// EnterQueryDml is called when entering the QueryDml production.
	EnterQueryDml(c *QueryDmlContext)

	// EnterQueryDdl is called when entering the QueryDdl production.
	EnterQueryDdl(c *QueryDdlContext)

	// EnterQueryExec is called when entering the QueryExec production.
	EnterQueryExec(c *QueryExecContext)

	// EnterExplainOption is called when entering the explainOption production.
	EnterExplainOption(c *ExplainOptionContext)

	// EnterAsIdent is called when entering the asIdent production.
	EnterAsIdent(c *AsIdentContext)

	// EnterAtIdent is called when entering the atIdent production.
	EnterAtIdent(c *AtIdentContext)

	// EnterByIdent is called when entering the byIdent production.
	EnterByIdent(c *ByIdentContext)

	// EnterSymbolPrimitive is called when entering the symbolPrimitive production.
	EnterSymbolPrimitive(c *SymbolPrimitiveContext)

	// EnterDql is called when entering the dql production.
	EnterDql(c *DqlContext)

	// EnterExecCommand is called when entering the execCommand production.
	EnterExecCommand(c *ExecCommandContext)

	// EnterDdl is called when entering the ddl production.
	EnterDdl(c *DdlContext)

	// EnterCreateTable is called when entering the CreateTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterCreateIndex is called when entering the CreateIndex production.
	EnterCreateIndex(c *CreateIndexContext)

	// EnterDropTable is called when entering the DropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropIndex is called when entering the DropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterDmlBaseWrapper is called when entering the DmlBaseWrapper production.
	EnterDmlBaseWrapper(c *DmlBaseWrapperContext)

	// EnterDmlDelete is called when entering the DmlDelete production.
	EnterDmlDelete(c *DmlDeleteContext)

	// EnterDmlInsertReturning is called when entering the DmlInsertReturning production.
	EnterDmlInsertReturning(c *DmlInsertReturningContext)

	// EnterDmlBase is called when entering the DmlBase production.
	EnterDmlBase(c *DmlBaseContext)

	// EnterDmlBaseCommand is called when entering the dmlBaseCommand production.
	EnterDmlBaseCommand(c *DmlBaseCommandContext)

	// EnterPathSimple is called when entering the pathSimple production.
	EnterPathSimple(c *PathSimpleContext)

	// EnterPathSimpleLiteral is called when entering the PathSimpleLiteral production.
	EnterPathSimpleLiteral(c *PathSimpleLiteralContext)

	// EnterPathSimpleSymbol is called when entering the PathSimpleSymbol production.
	EnterPathSimpleSymbol(c *PathSimpleSymbolContext)

	// EnterPathSimpleDotSymbol is called when entering the PathSimpleDotSymbol production.
	EnterPathSimpleDotSymbol(c *PathSimpleDotSymbolContext)

	// EnterReplaceCommand is called when entering the replaceCommand production.
	EnterReplaceCommand(c *ReplaceCommandContext)

	// EnterUpsertCommand is called when entering the upsertCommand production.
	EnterUpsertCommand(c *UpsertCommandContext)

	// EnterRemoveCommand is called when entering the removeCommand production.
	EnterRemoveCommand(c *RemoveCommandContext)

	// EnterInsertCommandReturning is called when entering the insertCommandReturning production.
	EnterInsertCommandReturning(c *InsertCommandReturningContext)

	// EnterInsertLegacy is called when entering the InsertLegacy production.
	EnterInsertLegacy(c *InsertLegacyContext)

	// EnterInsert is called when entering the Insert production.
	EnterInsert(c *InsertContext)

	// EnterOnConflictLegacy is called when entering the OnConflictLegacy production.
	EnterOnConflictLegacy(c *OnConflictLegacyContext)

	// EnterOnConflict is called when entering the OnConflict production.
	EnterOnConflict(c *OnConflictContext)

	// EnterConflictTarget is called when entering the conflictTarget production.
	EnterConflictTarget(c *ConflictTargetContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterConflictAction is called when entering the conflictAction production.
	EnterConflictAction(c *ConflictActionContext)

	// EnterDoReplace is called when entering the doReplace production.
	EnterDoReplace(c *DoReplaceContext)

	// EnterDoUpdate is called when entering the doUpdate production.
	EnterDoUpdate(c *DoUpdateContext)

	// EnterUpdateClause is called when entering the updateClause production.
	EnterUpdateClause(c *UpdateClauseContext)

	// EnterSetCommand is called when entering the setCommand production.
	EnterSetCommand(c *SetCommandContext)

	// EnterSetAssignment is called when entering the setAssignment production.
	EnterSetAssignment(c *SetAssignmentContext)

	// EnterDeleteCommand is called when entering the deleteCommand production.
	EnterDeleteCommand(c *DeleteCommandContext)

	// EnterReturningClause is called when entering the returningClause production.
	EnterReturningClause(c *ReturningClauseContext)

	// EnterReturningColumn is called when entering the returningColumn production.
	EnterReturningColumn(c *ReturningColumnContext)

	// EnterFromClauseSimpleExplicit is called when entering the FromClauseSimpleExplicit production.
	EnterFromClauseSimpleExplicit(c *FromClauseSimpleExplicitContext)

	// EnterFromClauseSimpleImplicit is called when entering the FromClauseSimpleImplicit production.
	EnterFromClauseSimpleImplicit(c *FromClauseSimpleImplicitContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterSelectAll is called when entering the SelectAll production.
	EnterSelectAll(c *SelectAllContext)

	// EnterSelectItems is called when entering the SelectItems production.
	EnterSelectItems(c *SelectItemsContext)

	// EnterSelectValue is called when entering the SelectValue production.
	EnterSelectValue(c *SelectValueContext)

	// EnterSelectPivot is called when entering the SelectPivot production.
	EnterSelectPivot(c *SelectPivotContext)

	// EnterProjectionItems is called when entering the projectionItems production.
	EnterProjectionItems(c *ProjectionItemsContext)

	// EnterProjectionItem is called when entering the projectionItem production.
	EnterProjectionItem(c *ProjectionItemContext)

	// EnterSetQuantifierStrategy is called when entering the setQuantifierStrategy production.
	EnterSetQuantifierStrategy(c *SetQuantifierStrategyContext)

	// EnterLetClause is called when entering the letClause production.
	EnterLetClause(c *LetClauseContext)

	// EnterLetBinding is called when entering the letBinding production.
	EnterLetBinding(c *LetBindingContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderSortSpec is called when entering the orderSortSpec production.
	EnterOrderSortSpec(c *OrderSortSpecContext)

	// EnterGroupClause is called when entering the groupClause production.
	EnterGroupClause(c *GroupClauseContext)

	// EnterGroupAlias is called when entering the groupAlias production.
	EnterGroupAlias(c *GroupAliasContext)

	// EnterGroupKey is called when entering the groupKey production.
	EnterGroupKey(c *GroupKeyContext)

	// EnterOver is called when entering the over production.
	EnterOver(c *OverContext)

	// EnterWindowPartitionList is called when entering the windowPartitionList production.
	EnterWindowPartitionList(c *WindowPartitionListContext)

	// EnterWindowSortSpecList is called when entering the windowSortSpecList production.
	EnterWindowSortSpecList(c *WindowSortSpecListContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterWhereClauseSelect is called when entering the whereClauseSelect production.
	EnterWhereClauseSelect(c *WhereClauseSelectContext)

	// EnterOffsetByClause is called when entering the offsetByClause production.
	EnterOffsetByClause(c *OffsetByClauseContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterGpmlPattern is called when entering the gpmlPattern production.
	EnterGpmlPattern(c *GpmlPatternContext)

	// EnterGpmlPatternList is called when entering the gpmlPatternList production.
	EnterGpmlPatternList(c *GpmlPatternListContext)

	// EnterMatchPattern is called when entering the matchPattern production.
	EnterMatchPattern(c *MatchPatternContext)

	// EnterGraphPart is called when entering the graphPart production.
	EnterGraphPart(c *GraphPartContext)

	// EnterSelectorBasic is called when entering the SelectorBasic production.
	EnterSelectorBasic(c *SelectorBasicContext)

	// EnterSelectorAny is called when entering the SelectorAny production.
	EnterSelectorAny(c *SelectorAnyContext)

	// EnterSelectorShortest is called when entering the SelectorShortest production.
	EnterSelectorShortest(c *SelectorShortestContext)

	// EnterPatternPathVariable is called when entering the patternPathVariable production.
	EnterPatternPathVariable(c *PatternPathVariableContext)

	// EnterPatternRestrictor is called when entering the patternRestrictor production.
	EnterPatternRestrictor(c *PatternRestrictorContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterEdgeWithSpec is called when entering the EdgeWithSpec production.
	EnterEdgeWithSpec(c *EdgeWithSpecContext)

	// EnterEdgeAbbreviated is called when entering the EdgeAbbreviated production.
	EnterEdgeAbbreviated(c *EdgeAbbreviatedContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterPatternQuantifier is called when entering the patternQuantifier production.
	EnterPatternQuantifier(c *PatternQuantifierContext)

	// EnterEdgeSpecRight is called when entering the EdgeSpecRight production.
	EnterEdgeSpecRight(c *EdgeSpecRightContext)

	// EnterEdgeSpecUndirected is called when entering the EdgeSpecUndirected production.
	EnterEdgeSpecUndirected(c *EdgeSpecUndirectedContext)

	// EnterEdgeSpecLeft is called when entering the EdgeSpecLeft production.
	EnterEdgeSpecLeft(c *EdgeSpecLeftContext)

	// EnterEdgeSpecUndirectedRight is called when entering the EdgeSpecUndirectedRight production.
	EnterEdgeSpecUndirectedRight(c *EdgeSpecUndirectedRightContext)

	// EnterEdgeSpecUndirectedLeft is called when entering the EdgeSpecUndirectedLeft production.
	EnterEdgeSpecUndirectedLeft(c *EdgeSpecUndirectedLeftContext)

	// EnterEdgeSpecBidirectional is called when entering the EdgeSpecBidirectional production.
	EnterEdgeSpecBidirectional(c *EdgeSpecBidirectionalContext)

	// EnterEdgeSpecUndirectedBidirectional is called when entering the EdgeSpecUndirectedBidirectional production.
	EnterEdgeSpecUndirectedBidirectional(c *EdgeSpecUndirectedBidirectionalContext)

	// EnterEdgeSpec is called when entering the edgeSpec production.
	EnterEdgeSpec(c *EdgeSpecContext)

	// EnterPatternPartLabel is called when entering the patternPartLabel production.
	EnterPatternPartLabel(c *PatternPartLabelContext)

	// EnterEdgeAbbrev is called when entering the edgeAbbrev production.
	EnterEdgeAbbrev(c *EdgeAbbrevContext)

	// EnterTableWrapped is called when entering the TableWrapped production.
	EnterTableWrapped(c *TableWrappedContext)

	// EnterTableCrossJoin is called when entering the TableCrossJoin production.
	EnterTableCrossJoin(c *TableCrossJoinContext)

	// EnterTableQualifiedJoin is called when entering the TableQualifiedJoin production.
	EnterTableQualifiedJoin(c *TableQualifiedJoinContext)

	// EnterTableRefBase is called when entering the TableRefBase production.
	EnterTableRefBase(c *TableRefBaseContext)

	// EnterTableNonJoin is called when entering the tableNonJoin production.
	EnterTableNonJoin(c *TableNonJoinContext)

	// EnterTableBaseRefSymbol is called when entering the TableBaseRefSymbol production.
	EnterTableBaseRefSymbol(c *TableBaseRefSymbolContext)

	// EnterTableBaseRefClauses is called when entering the TableBaseRefClauses production.
	EnterTableBaseRefClauses(c *TableBaseRefClausesContext)

	// EnterTableBaseRefMatch is called when entering the TableBaseRefMatch production.
	EnterTableBaseRefMatch(c *TableBaseRefMatchContext)

	// EnterTableUnpivot is called when entering the tableUnpivot production.
	EnterTableUnpivot(c *TableUnpivotContext)

	// EnterJoinRhsBase is called when entering the JoinRhsBase production.
	EnterJoinRhsBase(c *JoinRhsBaseContext)

	// EnterJoinRhsTableJoined is called when entering the JoinRhsTableJoined production.
	EnterJoinRhsTableJoined(c *JoinRhsTableJoinedContext)

	// EnterJoinSpec is called when entering the joinSpec production.
	EnterJoinSpec(c *JoinSpecContext)

	// EnterJoinType is called when entering the joinType production.
	EnterJoinType(c *JoinTypeContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterIntersect is called when entering the Intersect production.
	EnterIntersect(c *IntersectContext)

	// EnterQueryBase is called when entering the QueryBase production.
	EnterQueryBase(c *QueryBaseContext)

	// EnterExcept is called when entering the Except production.
	EnterExcept(c *ExceptContext)

	// EnterUnion is called when entering the Union production.
	EnterUnion(c *UnionContext)

	// EnterSfwQuery is called when entering the SfwQuery production.
	EnterSfwQuery(c *SfwQueryContext)

	// EnterSfwBase is called when entering the SfwBase production.
	EnterSfwBase(c *SfwBaseContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterExprOrBase is called when entering the ExprOrBase production.
	EnterExprOrBase(c *ExprOrBaseContext)

	// EnterExprAndBase is called when entering the ExprAndBase production.
	EnterExprAndBase(c *ExprAndBaseContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterExprNotBase is called when entering the ExprNotBase production.
	EnterExprNotBase(c *ExprNotBaseContext)

	// EnterPredicateIn is called when entering the PredicateIn production.
	EnterPredicateIn(c *PredicateInContext)

	// EnterPredicateBetween is called when entering the PredicateBetween production.
	EnterPredicateBetween(c *PredicateBetweenContext)

	// EnterPredicateBase is called when entering the PredicateBase production.
	EnterPredicateBase(c *PredicateBaseContext)

	// EnterPredicateComparison is called when entering the PredicateComparison production.
	EnterPredicateComparison(c *PredicateComparisonContext)

	// EnterPredicateIs is called when entering the PredicateIs production.
	EnterPredicateIs(c *PredicateIsContext)

	// EnterPredicateLike is called when entering the PredicateLike production.
	EnterPredicateLike(c *PredicateLikeContext)

	// EnterMathOp00 is called when entering the mathOp00 production.
	EnterMathOp00(c *MathOp00Context)

	// EnterMathOp01 is called when entering the mathOp01 production.
	EnterMathOp01(c *MathOp01Context)

	// EnterMathOp02 is called when entering the mathOp02 production.
	EnterMathOp02(c *MathOp02Context)

	// EnterValueExpr is called when entering the valueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// EnterExprPrimaryPath is called when entering the ExprPrimaryPath production.
	EnterExprPrimaryPath(c *ExprPrimaryPathContext)

	// EnterExprPrimaryBase is called when entering the ExprPrimaryBase production.
	EnterExprPrimaryBase(c *ExprPrimaryBaseContext)

	// EnterExprTermWrappedQuery is called when entering the ExprTermWrappedQuery production.
	EnterExprTermWrappedQuery(c *ExprTermWrappedQueryContext)

	// EnterExprTermBase is called when entering the ExprTermBase production.
	EnterExprTermBase(c *ExprTermBaseContext)

	// EnterNullIf is called when entering the nullIf production.
	EnterNullIf(c *NullIfContext)

	// EnterCoalesce is called when entering the coalesce production.
	EnterCoalesce(c *CoalesceContext)

	// EnterCaseExpr is called when entering the caseExpr production.
	EnterCaseExpr(c *CaseExprContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterValueRow is called when entering the valueRow production.
	EnterValueRow(c *ValueRowContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterSequenceConstructor is called when entering the sequenceConstructor production.
	EnterSequenceConstructor(c *SequenceConstructorContext)

	// EnterSubstring is called when entering the substring production.
	EnterSubstring(c *SubstringContext)

	// EnterCountAll is called when entering the CountAll production.
	EnterCountAll(c *CountAllContext)

	// EnterAggregateBase is called when entering the AggregateBase production.
	EnterAggregateBase(c *AggregateBaseContext)

	// EnterLagLeadFunction is called when entering the LagLeadFunction production.
	EnterLagLeadFunction(c *LagLeadFunctionContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterCanLosslessCast is called when entering the canLosslessCast production.
	EnterCanLosslessCast(c *CanLosslessCastContext)

	// EnterCanCast is called when entering the canCast production.
	EnterCanCast(c *CanCastContext)

	// EnterExtract is called when entering the extract production.
	EnterExtract(c *ExtractContext)

	// EnterTrimFunction is called when entering the trimFunction production.
	EnterTrimFunction(c *TrimFunctionContext)

	// EnterDateFunction is called when entering the dateFunction production.
	EnterDateFunction(c *DateFunctionContext)

	// EnterFunctionCallReserved is called when entering the FunctionCallReserved production.
	EnterFunctionCallReserved(c *FunctionCallReservedContext)

	// EnterFunctionCallIdent is called when entering the FunctionCallIdent production.
	EnterFunctionCallIdent(c *FunctionCallIdentContext)

	// EnterPathStepIndexExpr is called when entering the PathStepIndexExpr production.
	EnterPathStepIndexExpr(c *PathStepIndexExprContext)

	// EnterPathStepIndexAll is called when entering the PathStepIndexAll production.
	EnterPathStepIndexAll(c *PathStepIndexAllContext)

	// EnterPathStepDotExpr is called when entering the PathStepDotExpr production.
	EnterPathStepDotExpr(c *PathStepDotExprContext)

	// EnterPathStepDotAll is called when entering the PathStepDotAll production.
	EnterPathStepDotAll(c *PathStepDotAllContext)

	// EnterExprGraphMatchMany is called when entering the exprGraphMatchMany production.
	EnterExprGraphMatchMany(c *ExprGraphMatchManyContext)

	// EnterExprGraphMatchOne is called when entering the exprGraphMatchOne production.
	EnterExprGraphMatchOne(c *ExprGraphMatchOneContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterVarRefExpr is called when entering the varRefExpr production.
	EnterVarRefExpr(c *VarRefExprContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterBag is called when entering the bag production.
	EnterBag(c *BagContext)

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterLiteralNull is called when entering the LiteralNull production.
	EnterLiteralNull(c *LiteralNullContext)

	// EnterLiteralMissing is called when entering the LiteralMissing production.
	EnterLiteralMissing(c *LiteralMissingContext)

	// EnterLiteralTrue is called when entering the LiteralTrue production.
	EnterLiteralTrue(c *LiteralTrueContext)

	// EnterLiteralFalse is called when entering the LiteralFalse production.
	EnterLiteralFalse(c *LiteralFalseContext)

	// EnterLiteralString is called when entering the LiteralString production.
	EnterLiteralString(c *LiteralStringContext)

	// EnterLiteralInteger is called when entering the LiteralInteger production.
	EnterLiteralInteger(c *LiteralIntegerContext)

	// EnterLiteralDecimal is called when entering the LiteralDecimal production.
	EnterLiteralDecimal(c *LiteralDecimalContext)

	// EnterLiteralIon is called when entering the LiteralIon production.
	EnterLiteralIon(c *LiteralIonContext)

	// EnterLiteralDate is called when entering the LiteralDate production.
	EnterLiteralDate(c *LiteralDateContext)

	// EnterLiteralTime is called when entering the LiteralTime production.
	EnterLiteralTime(c *LiteralTimeContext)

	// EnterTypeAtomic is called when entering the TypeAtomic production.
	EnterTypeAtomic(c *TypeAtomicContext)

	// EnterTypeArgSingle is called when entering the TypeArgSingle production.
	EnterTypeArgSingle(c *TypeArgSingleContext)

	// EnterTypeVarChar is called when entering the TypeVarChar production.
	EnterTypeVarChar(c *TypeVarCharContext)

	// EnterTypeArgDouble is called when entering the TypeArgDouble production.
	EnterTypeArgDouble(c *TypeArgDoubleContext)

	// EnterTypeTimeZone is called when entering the TypeTimeZone production.
	EnterTypeTimeZone(c *TypeTimeZoneContext)

	// EnterTypeCustom is called when entering the TypeCustom production.
	EnterTypeCustom(c *TypeCustomContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitQueryDql is called when exiting the QueryDql production.
	ExitQueryDql(c *QueryDqlContext)

	// ExitQueryDml is called when exiting the QueryDml production.
	ExitQueryDml(c *QueryDmlContext)

	// ExitQueryDdl is called when exiting the QueryDdl production.
	ExitQueryDdl(c *QueryDdlContext)

	// ExitQueryExec is called when exiting the QueryExec production.
	ExitQueryExec(c *QueryExecContext)

	// ExitExplainOption is called when exiting the explainOption production.
	ExitExplainOption(c *ExplainOptionContext)

	// ExitAsIdent is called when exiting the asIdent production.
	ExitAsIdent(c *AsIdentContext)

	// ExitAtIdent is called when exiting the atIdent production.
	ExitAtIdent(c *AtIdentContext)

	// ExitByIdent is called when exiting the byIdent production.
	ExitByIdent(c *ByIdentContext)

	// ExitSymbolPrimitive is called when exiting the symbolPrimitive production.
	ExitSymbolPrimitive(c *SymbolPrimitiveContext)

	// ExitDql is called when exiting the dql production.
	ExitDql(c *DqlContext)

	// ExitExecCommand is called when exiting the execCommand production.
	ExitExecCommand(c *ExecCommandContext)

	// ExitDdl is called when exiting the ddl production.
	ExitDdl(c *DdlContext)

	// ExitCreateTable is called when exiting the CreateTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitCreateIndex is called when exiting the CreateIndex production.
	ExitCreateIndex(c *CreateIndexContext)

	// ExitDropTable is called when exiting the DropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropIndex is called when exiting the DropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitDmlBaseWrapper is called when exiting the DmlBaseWrapper production.
	ExitDmlBaseWrapper(c *DmlBaseWrapperContext)

	// ExitDmlDelete is called when exiting the DmlDelete production.
	ExitDmlDelete(c *DmlDeleteContext)

	// ExitDmlInsertReturning is called when exiting the DmlInsertReturning production.
	ExitDmlInsertReturning(c *DmlInsertReturningContext)

	// ExitDmlBase is called when exiting the DmlBase production.
	ExitDmlBase(c *DmlBaseContext)

	// ExitDmlBaseCommand is called when exiting the dmlBaseCommand production.
	ExitDmlBaseCommand(c *DmlBaseCommandContext)

	// ExitPathSimple is called when exiting the pathSimple production.
	ExitPathSimple(c *PathSimpleContext)

	// ExitPathSimpleLiteral is called when exiting the PathSimpleLiteral production.
	ExitPathSimpleLiteral(c *PathSimpleLiteralContext)

	// ExitPathSimpleSymbol is called when exiting the PathSimpleSymbol production.
	ExitPathSimpleSymbol(c *PathSimpleSymbolContext)

	// ExitPathSimpleDotSymbol is called when exiting the PathSimpleDotSymbol production.
	ExitPathSimpleDotSymbol(c *PathSimpleDotSymbolContext)

	// ExitReplaceCommand is called when exiting the replaceCommand production.
	ExitReplaceCommand(c *ReplaceCommandContext)

	// ExitUpsertCommand is called when exiting the upsertCommand production.
	ExitUpsertCommand(c *UpsertCommandContext)

	// ExitRemoveCommand is called when exiting the removeCommand production.
	ExitRemoveCommand(c *RemoveCommandContext)

	// ExitInsertCommandReturning is called when exiting the insertCommandReturning production.
	ExitInsertCommandReturning(c *InsertCommandReturningContext)

	// ExitInsertLegacy is called when exiting the InsertLegacy production.
	ExitInsertLegacy(c *InsertLegacyContext)

	// ExitInsert is called when exiting the Insert production.
	ExitInsert(c *InsertContext)

	// ExitOnConflictLegacy is called when exiting the OnConflictLegacy production.
	ExitOnConflictLegacy(c *OnConflictLegacyContext)

	// ExitOnConflict is called when exiting the OnConflict production.
	ExitOnConflict(c *OnConflictContext)

	// ExitConflictTarget is called when exiting the conflictTarget production.
	ExitConflictTarget(c *ConflictTargetContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitConflictAction is called when exiting the conflictAction production.
	ExitConflictAction(c *ConflictActionContext)

	// ExitDoReplace is called when exiting the doReplace production.
	ExitDoReplace(c *DoReplaceContext)

	// ExitDoUpdate is called when exiting the doUpdate production.
	ExitDoUpdate(c *DoUpdateContext)

	// ExitUpdateClause is called when exiting the updateClause production.
	ExitUpdateClause(c *UpdateClauseContext)

	// ExitSetCommand is called when exiting the setCommand production.
	ExitSetCommand(c *SetCommandContext)

	// ExitSetAssignment is called when exiting the setAssignment production.
	ExitSetAssignment(c *SetAssignmentContext)

	// ExitDeleteCommand is called when exiting the deleteCommand production.
	ExitDeleteCommand(c *DeleteCommandContext)

	// ExitReturningClause is called when exiting the returningClause production.
	ExitReturningClause(c *ReturningClauseContext)

	// ExitReturningColumn is called when exiting the returningColumn production.
	ExitReturningColumn(c *ReturningColumnContext)

	// ExitFromClauseSimpleExplicit is called when exiting the FromClauseSimpleExplicit production.
	ExitFromClauseSimpleExplicit(c *FromClauseSimpleExplicitContext)

	// ExitFromClauseSimpleImplicit is called when exiting the FromClauseSimpleImplicit production.
	ExitFromClauseSimpleImplicit(c *FromClauseSimpleImplicitContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitSelectAll is called when exiting the SelectAll production.
	ExitSelectAll(c *SelectAllContext)

	// ExitSelectItems is called when exiting the SelectItems production.
	ExitSelectItems(c *SelectItemsContext)

	// ExitSelectValue is called when exiting the SelectValue production.
	ExitSelectValue(c *SelectValueContext)

	// ExitSelectPivot is called when exiting the SelectPivot production.
	ExitSelectPivot(c *SelectPivotContext)

	// ExitProjectionItems is called when exiting the projectionItems production.
	ExitProjectionItems(c *ProjectionItemsContext)

	// ExitProjectionItem is called when exiting the projectionItem production.
	ExitProjectionItem(c *ProjectionItemContext)

	// ExitSetQuantifierStrategy is called when exiting the setQuantifierStrategy production.
	ExitSetQuantifierStrategy(c *SetQuantifierStrategyContext)

	// ExitLetClause is called when exiting the letClause production.
	ExitLetClause(c *LetClauseContext)

	// ExitLetBinding is called when exiting the letBinding production.
	ExitLetBinding(c *LetBindingContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderSortSpec is called when exiting the orderSortSpec production.
	ExitOrderSortSpec(c *OrderSortSpecContext)

	// ExitGroupClause is called when exiting the groupClause production.
	ExitGroupClause(c *GroupClauseContext)

	// ExitGroupAlias is called when exiting the groupAlias production.
	ExitGroupAlias(c *GroupAliasContext)

	// ExitGroupKey is called when exiting the groupKey production.
	ExitGroupKey(c *GroupKeyContext)

	// ExitOver is called when exiting the over production.
	ExitOver(c *OverContext)

	// ExitWindowPartitionList is called when exiting the windowPartitionList production.
	ExitWindowPartitionList(c *WindowPartitionListContext)

	// ExitWindowSortSpecList is called when exiting the windowSortSpecList production.
	ExitWindowSortSpecList(c *WindowSortSpecListContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitWhereClauseSelect is called when exiting the whereClauseSelect production.
	ExitWhereClauseSelect(c *WhereClauseSelectContext)

	// ExitOffsetByClause is called when exiting the offsetByClause production.
	ExitOffsetByClause(c *OffsetByClauseContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitGpmlPattern is called when exiting the gpmlPattern production.
	ExitGpmlPattern(c *GpmlPatternContext)

	// ExitGpmlPatternList is called when exiting the gpmlPatternList production.
	ExitGpmlPatternList(c *GpmlPatternListContext)

	// ExitMatchPattern is called when exiting the matchPattern production.
	ExitMatchPattern(c *MatchPatternContext)

	// ExitGraphPart is called when exiting the graphPart production.
	ExitGraphPart(c *GraphPartContext)

	// ExitSelectorBasic is called when exiting the SelectorBasic production.
	ExitSelectorBasic(c *SelectorBasicContext)

	// ExitSelectorAny is called when exiting the SelectorAny production.
	ExitSelectorAny(c *SelectorAnyContext)

	// ExitSelectorShortest is called when exiting the SelectorShortest production.
	ExitSelectorShortest(c *SelectorShortestContext)

	// ExitPatternPathVariable is called when exiting the patternPathVariable production.
	ExitPatternPathVariable(c *PatternPathVariableContext)

	// ExitPatternRestrictor is called when exiting the patternRestrictor production.
	ExitPatternRestrictor(c *PatternRestrictorContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitEdgeWithSpec is called when exiting the EdgeWithSpec production.
	ExitEdgeWithSpec(c *EdgeWithSpecContext)

	// ExitEdgeAbbreviated is called when exiting the EdgeAbbreviated production.
	ExitEdgeAbbreviated(c *EdgeAbbreviatedContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitPatternQuantifier is called when exiting the patternQuantifier production.
	ExitPatternQuantifier(c *PatternQuantifierContext)

	// ExitEdgeSpecRight is called when exiting the EdgeSpecRight production.
	ExitEdgeSpecRight(c *EdgeSpecRightContext)

	// ExitEdgeSpecUndirected is called when exiting the EdgeSpecUndirected production.
	ExitEdgeSpecUndirected(c *EdgeSpecUndirectedContext)

	// ExitEdgeSpecLeft is called when exiting the EdgeSpecLeft production.
	ExitEdgeSpecLeft(c *EdgeSpecLeftContext)

	// ExitEdgeSpecUndirectedRight is called when exiting the EdgeSpecUndirectedRight production.
	ExitEdgeSpecUndirectedRight(c *EdgeSpecUndirectedRightContext)

	// ExitEdgeSpecUndirectedLeft is called when exiting the EdgeSpecUndirectedLeft production.
	ExitEdgeSpecUndirectedLeft(c *EdgeSpecUndirectedLeftContext)

	// ExitEdgeSpecBidirectional is called when exiting the EdgeSpecBidirectional production.
	ExitEdgeSpecBidirectional(c *EdgeSpecBidirectionalContext)

	// ExitEdgeSpecUndirectedBidirectional is called when exiting the EdgeSpecUndirectedBidirectional production.
	ExitEdgeSpecUndirectedBidirectional(c *EdgeSpecUndirectedBidirectionalContext)

	// ExitEdgeSpec is called when exiting the edgeSpec production.
	ExitEdgeSpec(c *EdgeSpecContext)

	// ExitPatternPartLabel is called when exiting the patternPartLabel production.
	ExitPatternPartLabel(c *PatternPartLabelContext)

	// ExitEdgeAbbrev is called when exiting the edgeAbbrev production.
	ExitEdgeAbbrev(c *EdgeAbbrevContext)

	// ExitTableWrapped is called when exiting the TableWrapped production.
	ExitTableWrapped(c *TableWrappedContext)

	// ExitTableCrossJoin is called when exiting the TableCrossJoin production.
	ExitTableCrossJoin(c *TableCrossJoinContext)

	// ExitTableQualifiedJoin is called when exiting the TableQualifiedJoin production.
	ExitTableQualifiedJoin(c *TableQualifiedJoinContext)

	// ExitTableRefBase is called when exiting the TableRefBase production.
	ExitTableRefBase(c *TableRefBaseContext)

	// ExitTableNonJoin is called when exiting the tableNonJoin production.
	ExitTableNonJoin(c *TableNonJoinContext)

	// ExitTableBaseRefSymbol is called when exiting the TableBaseRefSymbol production.
	ExitTableBaseRefSymbol(c *TableBaseRefSymbolContext)

	// ExitTableBaseRefClauses is called when exiting the TableBaseRefClauses production.
	ExitTableBaseRefClauses(c *TableBaseRefClausesContext)

	// ExitTableBaseRefMatch is called when exiting the TableBaseRefMatch production.
	ExitTableBaseRefMatch(c *TableBaseRefMatchContext)

	// ExitTableUnpivot is called when exiting the tableUnpivot production.
	ExitTableUnpivot(c *TableUnpivotContext)

	// ExitJoinRhsBase is called when exiting the JoinRhsBase production.
	ExitJoinRhsBase(c *JoinRhsBaseContext)

	// ExitJoinRhsTableJoined is called when exiting the JoinRhsTableJoined production.
	ExitJoinRhsTableJoined(c *JoinRhsTableJoinedContext)

	// ExitJoinSpec is called when exiting the joinSpec production.
	ExitJoinSpec(c *JoinSpecContext)

	// ExitJoinType is called when exiting the joinType production.
	ExitJoinType(c *JoinTypeContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitIntersect is called when exiting the Intersect production.
	ExitIntersect(c *IntersectContext)

	// ExitQueryBase is called when exiting the QueryBase production.
	ExitQueryBase(c *QueryBaseContext)

	// ExitExcept is called when exiting the Except production.
	ExitExcept(c *ExceptContext)

	// ExitUnion is called when exiting the Union production.
	ExitUnion(c *UnionContext)

	// ExitSfwQuery is called when exiting the SfwQuery production.
	ExitSfwQuery(c *SfwQueryContext)

	// ExitSfwBase is called when exiting the SfwBase production.
	ExitSfwBase(c *SfwBaseContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitExprOrBase is called when exiting the ExprOrBase production.
	ExitExprOrBase(c *ExprOrBaseContext)

	// ExitExprAndBase is called when exiting the ExprAndBase production.
	ExitExprAndBase(c *ExprAndBaseContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitExprNotBase is called when exiting the ExprNotBase production.
	ExitExprNotBase(c *ExprNotBaseContext)

	// ExitPredicateIn is called when exiting the PredicateIn production.
	ExitPredicateIn(c *PredicateInContext)

	// ExitPredicateBetween is called when exiting the PredicateBetween production.
	ExitPredicateBetween(c *PredicateBetweenContext)

	// ExitPredicateBase is called when exiting the PredicateBase production.
	ExitPredicateBase(c *PredicateBaseContext)

	// ExitPredicateComparison is called when exiting the PredicateComparison production.
	ExitPredicateComparison(c *PredicateComparisonContext)

	// ExitPredicateIs is called when exiting the PredicateIs production.
	ExitPredicateIs(c *PredicateIsContext)

	// ExitPredicateLike is called when exiting the PredicateLike production.
	ExitPredicateLike(c *PredicateLikeContext)

	// ExitMathOp00 is called when exiting the mathOp00 production.
	ExitMathOp00(c *MathOp00Context)

	// ExitMathOp01 is called when exiting the mathOp01 production.
	ExitMathOp01(c *MathOp01Context)

	// ExitMathOp02 is called when exiting the mathOp02 production.
	ExitMathOp02(c *MathOp02Context)

	// ExitValueExpr is called when exiting the valueExpr production.
	ExitValueExpr(c *ValueExprContext)

	// ExitExprPrimaryPath is called when exiting the ExprPrimaryPath production.
	ExitExprPrimaryPath(c *ExprPrimaryPathContext)

	// ExitExprPrimaryBase is called when exiting the ExprPrimaryBase production.
	ExitExprPrimaryBase(c *ExprPrimaryBaseContext)

	// ExitExprTermWrappedQuery is called when exiting the ExprTermWrappedQuery production.
	ExitExprTermWrappedQuery(c *ExprTermWrappedQueryContext)

	// ExitExprTermBase is called when exiting the ExprTermBase production.
	ExitExprTermBase(c *ExprTermBaseContext)

	// ExitNullIf is called when exiting the nullIf production.
	ExitNullIf(c *NullIfContext)

	// ExitCoalesce is called when exiting the coalesce production.
	ExitCoalesce(c *CoalesceContext)

	// ExitCaseExpr is called when exiting the caseExpr production.
	ExitCaseExpr(c *CaseExprContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitValueRow is called when exiting the valueRow production.
	ExitValueRow(c *ValueRowContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitSequenceConstructor is called when exiting the sequenceConstructor production.
	ExitSequenceConstructor(c *SequenceConstructorContext)

	// ExitSubstring is called when exiting the substring production.
	ExitSubstring(c *SubstringContext)

	// ExitCountAll is called when exiting the CountAll production.
	ExitCountAll(c *CountAllContext)

	// ExitAggregateBase is called when exiting the AggregateBase production.
	ExitAggregateBase(c *AggregateBaseContext)

	// ExitLagLeadFunction is called when exiting the LagLeadFunction production.
	ExitLagLeadFunction(c *LagLeadFunctionContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitCanLosslessCast is called when exiting the canLosslessCast production.
	ExitCanLosslessCast(c *CanLosslessCastContext)

	// ExitCanCast is called when exiting the canCast production.
	ExitCanCast(c *CanCastContext)

	// ExitExtract is called when exiting the extract production.
	ExitExtract(c *ExtractContext)

	// ExitTrimFunction is called when exiting the trimFunction production.
	ExitTrimFunction(c *TrimFunctionContext)

	// ExitDateFunction is called when exiting the dateFunction production.
	ExitDateFunction(c *DateFunctionContext)

	// ExitFunctionCallReserved is called when exiting the FunctionCallReserved production.
	ExitFunctionCallReserved(c *FunctionCallReservedContext)

	// ExitFunctionCallIdent is called when exiting the FunctionCallIdent production.
	ExitFunctionCallIdent(c *FunctionCallIdentContext)

	// ExitPathStepIndexExpr is called when exiting the PathStepIndexExpr production.
	ExitPathStepIndexExpr(c *PathStepIndexExprContext)

	// ExitPathStepIndexAll is called when exiting the PathStepIndexAll production.
	ExitPathStepIndexAll(c *PathStepIndexAllContext)

	// ExitPathStepDotExpr is called when exiting the PathStepDotExpr production.
	ExitPathStepDotExpr(c *PathStepDotExprContext)

	// ExitPathStepDotAll is called when exiting the PathStepDotAll production.
	ExitPathStepDotAll(c *PathStepDotAllContext)

	// ExitExprGraphMatchMany is called when exiting the exprGraphMatchMany production.
	ExitExprGraphMatchMany(c *ExprGraphMatchManyContext)

	// ExitExprGraphMatchOne is called when exiting the exprGraphMatchOne production.
	ExitExprGraphMatchOne(c *ExprGraphMatchOneContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitVarRefExpr is called when exiting the varRefExpr production.
	ExitVarRefExpr(c *VarRefExprContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitBag is called when exiting the bag production.
	ExitBag(c *BagContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitLiteralNull is called when exiting the LiteralNull production.
	ExitLiteralNull(c *LiteralNullContext)

	// ExitLiteralMissing is called when exiting the LiteralMissing production.
	ExitLiteralMissing(c *LiteralMissingContext)

	// ExitLiteralTrue is called when exiting the LiteralTrue production.
	ExitLiteralTrue(c *LiteralTrueContext)

	// ExitLiteralFalse is called when exiting the LiteralFalse production.
	ExitLiteralFalse(c *LiteralFalseContext)

	// ExitLiteralString is called when exiting the LiteralString production.
	ExitLiteralString(c *LiteralStringContext)

	// ExitLiteralInteger is called when exiting the LiteralInteger production.
	ExitLiteralInteger(c *LiteralIntegerContext)

	// ExitLiteralDecimal is called when exiting the LiteralDecimal production.
	ExitLiteralDecimal(c *LiteralDecimalContext)

	// ExitLiteralIon is called when exiting the LiteralIon production.
	ExitLiteralIon(c *LiteralIonContext)

	// ExitLiteralDate is called when exiting the LiteralDate production.
	ExitLiteralDate(c *LiteralDateContext)

	// ExitLiteralTime is called when exiting the LiteralTime production.
	ExitLiteralTime(c *LiteralTimeContext)

	// ExitTypeAtomic is called when exiting the TypeAtomic production.
	ExitTypeAtomic(c *TypeAtomicContext)

	// ExitTypeArgSingle is called when exiting the TypeArgSingle production.
	ExitTypeArgSingle(c *TypeArgSingleContext)

	// ExitTypeVarChar is called when exiting the TypeVarChar production.
	ExitTypeVarChar(c *TypeVarCharContext)

	// ExitTypeArgDouble is called when exiting the TypeArgDouble production.
	ExitTypeArgDouble(c *TypeArgDoubleContext)

	// ExitTypeTimeZone is called when exiting the TypeTimeZone production.
	ExitTypeTimeZone(c *TypeTimeZoneContext)

	// ExitTypeCustom is called when exiting the TypeCustom production.
	ExitTypeCustom(c *TypeCustomContext)
}
