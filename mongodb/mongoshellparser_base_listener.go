// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

// BaseMongoShellParserListener is a complete listener for a parse tree produced by MongoShellParser.
type BaseMongoShellParserListener struct{}

var _ MongoShellParserListener = &BaseMongoShellParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMongoShellParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMongoShellParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMongoShellParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMongoShellParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMongoShellParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMongoShellParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMongoShellParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMongoShellParserListener) ExitStatement(ctx *StatementContext) {}

// EnterShowDatabases is called when production showDatabases is entered.
func (s *BaseMongoShellParserListener) EnterShowDatabases(ctx *ShowDatabasesContext) {}

// ExitShowDatabases is called when production showDatabases is exited.
func (s *BaseMongoShellParserListener) ExitShowDatabases(ctx *ShowDatabasesContext) {}

// EnterShowCollections is called when production showCollections is entered.
func (s *BaseMongoShellParserListener) EnterShowCollections(ctx *ShowCollectionsContext) {}

// ExitShowCollections is called when production showCollections is exited.
func (s *BaseMongoShellParserListener) ExitShowCollections(ctx *ShowCollectionsContext) {}

// EnterGetCollectionNames is called when production getCollectionNames is entered.
func (s *BaseMongoShellParserListener) EnterGetCollectionNames(ctx *GetCollectionNamesContext) {}

// ExitGetCollectionNames is called when production getCollectionNames is exited.
func (s *BaseMongoShellParserListener) ExitGetCollectionNames(ctx *GetCollectionNamesContext) {}

// EnterGetCollectionInfos is called when production getCollectionInfos is entered.
func (s *BaseMongoShellParserListener) EnterGetCollectionInfos(ctx *GetCollectionInfosContext) {}

// ExitGetCollectionInfos is called when production getCollectionInfos is exited.
func (s *BaseMongoShellParserListener) ExitGetCollectionInfos(ctx *GetCollectionInfosContext) {}

// EnterCollectionOperation is called when production collectionOperation is entered.
func (s *BaseMongoShellParserListener) EnterCollectionOperation(ctx *CollectionOperationContext) {}

// ExitCollectionOperation is called when production collectionOperation is exited.
func (s *BaseMongoShellParserListener) ExitCollectionOperation(ctx *CollectionOperationContext) {}

// EnterDotAccess is called when production dotAccess is entered.
func (s *BaseMongoShellParserListener) EnterDotAccess(ctx *DotAccessContext) {}

// ExitDotAccess is called when production dotAccess is exited.
func (s *BaseMongoShellParserListener) ExitDotAccess(ctx *DotAccessContext) {}

// EnterBracketAccess is called when production bracketAccess is entered.
func (s *BaseMongoShellParserListener) EnterBracketAccess(ctx *BracketAccessContext) {}

// ExitBracketAccess is called when production bracketAccess is exited.
func (s *BaseMongoShellParserListener) ExitBracketAccess(ctx *BracketAccessContext) {}

// EnterGetCollectionAccess is called when production getCollectionAccess is entered.
func (s *BaseMongoShellParserListener) EnterGetCollectionAccess(ctx *GetCollectionAccessContext) {}

// ExitGetCollectionAccess is called when production getCollectionAccess is exited.
func (s *BaseMongoShellParserListener) ExitGetCollectionAccess(ctx *GetCollectionAccessContext) {}

// EnterMethodChain is called when production methodChain is entered.
func (s *BaseMongoShellParserListener) EnterMethodChain(ctx *MethodChainContext) {}

// ExitMethodChain is called when production methodChain is exited.
func (s *BaseMongoShellParserListener) ExitMethodChain(ctx *MethodChainContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseMongoShellParserListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseMongoShellParserListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterFindMethod is called when production findMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindMethod(ctx *FindMethodContext) {}

// ExitFindMethod is called when production findMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindMethod(ctx *FindMethodContext) {}

// EnterFindOneMethod is called when production findOneMethod is entered.
func (s *BaseMongoShellParserListener) EnterFindOneMethod(ctx *FindOneMethodContext) {}

// ExitFindOneMethod is called when production findOneMethod is exited.
func (s *BaseMongoShellParserListener) ExitFindOneMethod(ctx *FindOneMethodContext) {}

// EnterCountDocumentsMethod is called when production countDocumentsMethod is entered.
func (s *BaseMongoShellParserListener) EnterCountDocumentsMethod(ctx *CountDocumentsMethodContext) {}

// ExitCountDocumentsMethod is called when production countDocumentsMethod is exited.
func (s *BaseMongoShellParserListener) ExitCountDocumentsMethod(ctx *CountDocumentsMethodContext) {}

// EnterEstimatedDocumentCountMethod is called when production estimatedDocumentCountMethod is entered.
func (s *BaseMongoShellParserListener) EnterEstimatedDocumentCountMethod(ctx *EstimatedDocumentCountMethodContext) {
}

// ExitEstimatedDocumentCountMethod is called when production estimatedDocumentCountMethod is exited.
func (s *BaseMongoShellParserListener) ExitEstimatedDocumentCountMethod(ctx *EstimatedDocumentCountMethodContext) {
}

// EnterDistinctMethod is called when production distinctMethod is entered.
func (s *BaseMongoShellParserListener) EnterDistinctMethod(ctx *DistinctMethodContext) {}

// ExitDistinctMethod is called when production distinctMethod is exited.
func (s *BaseMongoShellParserListener) ExitDistinctMethod(ctx *DistinctMethodContext) {}

// EnterAggregateMethod is called when production aggregateMethod is entered.
func (s *BaseMongoShellParserListener) EnterAggregateMethod(ctx *AggregateMethodContext) {}

// ExitAggregateMethod is called when production aggregateMethod is exited.
func (s *BaseMongoShellParserListener) ExitAggregateMethod(ctx *AggregateMethodContext) {}

// EnterGetIndexesMethod is called when production getIndexesMethod is entered.
func (s *BaseMongoShellParserListener) EnterGetIndexesMethod(ctx *GetIndexesMethodContext) {}

// ExitGetIndexesMethod is called when production getIndexesMethod is exited.
func (s *BaseMongoShellParserListener) ExitGetIndexesMethod(ctx *GetIndexesMethodContext) {}

// EnterSortMethod is called when production sortMethod is entered.
func (s *BaseMongoShellParserListener) EnterSortMethod(ctx *SortMethodContext) {}

// ExitSortMethod is called when production sortMethod is exited.
func (s *BaseMongoShellParserListener) ExitSortMethod(ctx *SortMethodContext) {}

// EnterLimitMethod is called when production limitMethod is entered.
func (s *BaseMongoShellParserListener) EnterLimitMethod(ctx *LimitMethodContext) {}

// ExitLimitMethod is called when production limitMethod is exited.
func (s *BaseMongoShellParserListener) ExitLimitMethod(ctx *LimitMethodContext) {}

// EnterSkipMethod is called when production skipMethod is entered.
func (s *BaseMongoShellParserListener) EnterSkipMethod(ctx *SkipMethodContext) {}

// ExitSkipMethod is called when production skipMethod is exited.
func (s *BaseMongoShellParserListener) ExitSkipMethod(ctx *SkipMethodContext) {}

// EnterCountMethod is called when production countMethod is entered.
func (s *BaseMongoShellParserListener) EnterCountMethod(ctx *CountMethodContext) {}

// ExitCountMethod is called when production countMethod is exited.
func (s *BaseMongoShellParserListener) ExitCountMethod(ctx *CountMethodContext) {}

// EnterProjectionMethod is called when production projectionMethod is entered.
func (s *BaseMongoShellParserListener) EnterProjectionMethod(ctx *ProjectionMethodContext) {}

// ExitProjectionMethod is called when production projectionMethod is exited.
func (s *BaseMongoShellParserListener) ExitProjectionMethod(ctx *ProjectionMethodContext) {}

// EnterGenericMethod is called when production genericMethod is entered.
func (s *BaseMongoShellParserListener) EnterGenericMethod(ctx *GenericMethodContext) {}

// ExitGenericMethod is called when production genericMethod is exited.
func (s *BaseMongoShellParserListener) ExitGenericMethod(ctx *GenericMethodContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseMongoShellParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseMongoShellParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseMongoShellParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseMongoShellParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseMongoShellParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseMongoShellParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseMongoShellParserListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseMongoShellParserListener) ExitPair(ctx *PairContext) {}

// EnterUnquotedKey is called when production unquotedKey is entered.
func (s *BaseMongoShellParserListener) EnterUnquotedKey(ctx *UnquotedKeyContext) {}

// ExitUnquotedKey is called when production unquotedKey is exited.
func (s *BaseMongoShellParserListener) ExitUnquotedKey(ctx *UnquotedKeyContext) {}

// EnterQuotedKey is called when production quotedKey is entered.
func (s *BaseMongoShellParserListener) EnterQuotedKey(ctx *QuotedKeyContext) {}

// ExitQuotedKey is called when production quotedKey is exited.
func (s *BaseMongoShellParserListener) ExitQuotedKey(ctx *QuotedKeyContext) {}

// EnterDocumentValue is called when production documentValue is entered.
func (s *BaseMongoShellParserListener) EnterDocumentValue(ctx *DocumentValueContext) {}

// ExitDocumentValue is called when production documentValue is exited.
func (s *BaseMongoShellParserListener) ExitDocumentValue(ctx *DocumentValueContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseMongoShellParserListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseMongoShellParserListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterHelperValue is called when production helperValue is entered.
func (s *BaseMongoShellParserListener) EnterHelperValue(ctx *HelperValueContext) {}

// ExitHelperValue is called when production helperValue is exited.
func (s *BaseMongoShellParserListener) ExitHelperValue(ctx *HelperValueContext) {}

// EnterRegexLiteralValue is called when production regexLiteralValue is entered.
func (s *BaseMongoShellParserListener) EnterRegexLiteralValue(ctx *RegexLiteralValueContext) {}

// ExitRegexLiteralValue is called when production regexLiteralValue is exited.
func (s *BaseMongoShellParserListener) ExitRegexLiteralValue(ctx *RegexLiteralValueContext) {}

// EnterRegexpConstructorValue is called when production regexpConstructorValue is entered.
func (s *BaseMongoShellParserListener) EnterRegexpConstructorValue(ctx *RegexpConstructorValueContext) {
}

// ExitRegexpConstructorValue is called when production regexpConstructorValue is exited.
func (s *BaseMongoShellParserListener) ExitRegexpConstructorValue(ctx *RegexpConstructorValueContext) {
}

// EnterLiteralValue is called when production literalValue is entered.
func (s *BaseMongoShellParserListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *BaseMongoShellParserListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterNewKeywordValue is called when production newKeywordValue is entered.
func (s *BaseMongoShellParserListener) EnterNewKeywordValue(ctx *NewKeywordValueContext) {}

// ExitNewKeywordValue is called when production newKeywordValue is exited.
func (s *BaseMongoShellParserListener) ExitNewKeywordValue(ctx *NewKeywordValueContext) {}

// EnterNewKeywordError is called when production newKeywordError is entered.
func (s *BaseMongoShellParserListener) EnterNewKeywordError(ctx *NewKeywordErrorContext) {}

// ExitNewKeywordError is called when production newKeywordError is exited.
func (s *BaseMongoShellParserListener) ExitNewKeywordError(ctx *NewKeywordErrorContext) {}

// EnterArray is called when production array is entered.
func (s *BaseMongoShellParserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseMongoShellParserListener) ExitArray(ctx *ArrayContext) {}

// EnterHelperFunction is called when production helperFunction is entered.
func (s *BaseMongoShellParserListener) EnterHelperFunction(ctx *HelperFunctionContext) {}

// ExitHelperFunction is called when production helperFunction is exited.
func (s *BaseMongoShellParserListener) ExitHelperFunction(ctx *HelperFunctionContext) {}

// EnterObjectIdHelper is called when production objectIdHelper is entered.
func (s *BaseMongoShellParserListener) EnterObjectIdHelper(ctx *ObjectIdHelperContext) {}

// ExitObjectIdHelper is called when production objectIdHelper is exited.
func (s *BaseMongoShellParserListener) ExitObjectIdHelper(ctx *ObjectIdHelperContext) {}

// EnterIsoDateHelper is called when production isoDateHelper is entered.
func (s *BaseMongoShellParserListener) EnterIsoDateHelper(ctx *IsoDateHelperContext) {}

// ExitIsoDateHelper is called when production isoDateHelper is exited.
func (s *BaseMongoShellParserListener) ExitIsoDateHelper(ctx *IsoDateHelperContext) {}

// EnterDateHelper is called when production dateHelper is entered.
func (s *BaseMongoShellParserListener) EnterDateHelper(ctx *DateHelperContext) {}

// ExitDateHelper is called when production dateHelper is exited.
func (s *BaseMongoShellParserListener) ExitDateHelper(ctx *DateHelperContext) {}

// EnterUuidHelper is called when production uuidHelper is entered.
func (s *BaseMongoShellParserListener) EnterUuidHelper(ctx *UuidHelperContext) {}

// ExitUuidHelper is called when production uuidHelper is exited.
func (s *BaseMongoShellParserListener) ExitUuidHelper(ctx *UuidHelperContext) {}

// EnterLongHelper is called when production longHelper is entered.
func (s *BaseMongoShellParserListener) EnterLongHelper(ctx *LongHelperContext) {}

// ExitLongHelper is called when production longHelper is exited.
func (s *BaseMongoShellParserListener) ExitLongHelper(ctx *LongHelperContext) {}

// EnterInt32Helper is called when production int32Helper is entered.
func (s *BaseMongoShellParserListener) EnterInt32Helper(ctx *Int32HelperContext) {}

// ExitInt32Helper is called when production int32Helper is exited.
func (s *BaseMongoShellParserListener) ExitInt32Helper(ctx *Int32HelperContext) {}

// EnterDoubleHelper is called when production doubleHelper is entered.
func (s *BaseMongoShellParserListener) EnterDoubleHelper(ctx *DoubleHelperContext) {}

// ExitDoubleHelper is called when production doubleHelper is exited.
func (s *BaseMongoShellParserListener) ExitDoubleHelper(ctx *DoubleHelperContext) {}

// EnterDecimal128Helper is called when production decimal128Helper is entered.
func (s *BaseMongoShellParserListener) EnterDecimal128Helper(ctx *Decimal128HelperContext) {}

// ExitDecimal128Helper is called when production decimal128Helper is exited.
func (s *BaseMongoShellParserListener) ExitDecimal128Helper(ctx *Decimal128HelperContext) {}

// EnterTimestampDocHelper is called when production timestampDocHelper is entered.
func (s *BaseMongoShellParserListener) EnterTimestampDocHelper(ctx *TimestampDocHelperContext) {}

// ExitTimestampDocHelper is called when production timestampDocHelper is exited.
func (s *BaseMongoShellParserListener) ExitTimestampDocHelper(ctx *TimestampDocHelperContext) {}

// EnterTimestampArgsHelper is called when production timestampArgsHelper is entered.
func (s *BaseMongoShellParserListener) EnterTimestampArgsHelper(ctx *TimestampArgsHelperContext) {}

// ExitTimestampArgsHelper is called when production timestampArgsHelper is exited.
func (s *BaseMongoShellParserListener) ExitTimestampArgsHelper(ctx *TimestampArgsHelperContext) {}

// EnterRegExpConstructor is called when production regExpConstructor is entered.
func (s *BaseMongoShellParserListener) EnterRegExpConstructor(ctx *RegExpConstructorContext) {}

// ExitRegExpConstructor is called when production regExpConstructor is exited.
func (s *BaseMongoShellParserListener) ExitRegExpConstructor(ctx *RegExpConstructorContext) {}

// EnterStringLiteralValue is called when production stringLiteralValue is entered.
func (s *BaseMongoShellParserListener) EnterStringLiteralValue(ctx *StringLiteralValueContext) {}

// ExitStringLiteralValue is called when production stringLiteralValue is exited.
func (s *BaseMongoShellParserListener) ExitStringLiteralValue(ctx *StringLiteralValueContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseMongoShellParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseMongoShellParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterTrueLiteral is called when production trueLiteral is entered.
func (s *BaseMongoShellParserListener) EnterTrueLiteral(ctx *TrueLiteralContext) {}

// ExitTrueLiteral is called when production trueLiteral is exited.
func (s *BaseMongoShellParserListener) ExitTrueLiteral(ctx *TrueLiteralContext) {}

// EnterFalseLiteral is called when production falseLiteral is entered.
func (s *BaseMongoShellParserListener) EnterFalseLiteral(ctx *FalseLiteralContext) {}

// ExitFalseLiteral is called when production falseLiteral is exited.
func (s *BaseMongoShellParserListener) ExitFalseLiteral(ctx *FalseLiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseMongoShellParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseMongoShellParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseMongoShellParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseMongoShellParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMongoShellParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMongoShellParserListener) ExitIdentifier(ctx *IdentifierContext) {}
