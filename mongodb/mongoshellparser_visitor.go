// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MongoShellParser.
type MongoShellParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MongoShellParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MongoShellParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MongoShellParser#showDatabases.
	VisitShowDatabases(ctx *ShowDatabasesContext) interface{}

	// Visit a parse tree produced by MongoShellParser#showCollections.
	VisitShowCollections(ctx *ShowCollectionsContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getCollectionNames.
	VisitGetCollectionNames(ctx *GetCollectionNamesContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getCollectionInfos.
	VisitGetCollectionInfos(ctx *GetCollectionInfosContext) interface{}

	// Visit a parse tree produced by MongoShellParser#collectionOperation.
	VisitCollectionOperation(ctx *CollectionOperationContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dotAccess.
	VisitDotAccess(ctx *DotAccessContext) interface{}

	// Visit a parse tree produced by MongoShellParser#bracketAccess.
	VisitBracketAccess(ctx *BracketAccessContext) interface{}

	// Visit a parse tree produced by MongoShellParser#getCollectionAccess.
	VisitGetCollectionAccess(ctx *GetCollectionAccessContext) interface{}

	// Visit a parse tree produced by MongoShellParser#methodChain.
	VisitMethodChain(ctx *MethodChainContext) interface{}

	// Visit a parse tree produced by MongoShellParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}

	// Visit a parse tree produced by MongoShellParser#findMethod.
	VisitFindMethod(ctx *FindMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#findOneMethod.
	VisitFindOneMethod(ctx *FindOneMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#sortMethod.
	VisitSortMethod(ctx *SortMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#limitMethod.
	VisitLimitMethod(ctx *LimitMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#skipMethod.
	VisitSkipMethod(ctx *SkipMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#projectionMethod.
	VisitProjectionMethod(ctx *ProjectionMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#genericMethod.
	VisitGenericMethod(ctx *GenericMethodContext) interface{}

	// Visit a parse tree produced by MongoShellParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by MongoShellParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by MongoShellParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by MongoShellParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by MongoShellParser#unquotedKey.
	VisitUnquotedKey(ctx *UnquotedKeyContext) interface{}

	// Visit a parse tree produced by MongoShellParser#quotedKey.
	VisitQuotedKey(ctx *QuotedKeyContext) interface{}

	// Visit a parse tree produced by MongoShellParser#documentValue.
	VisitDocumentValue(ctx *DocumentValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#helperValue.
	VisitHelperValue(ctx *HelperValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#regexLiteralValue.
	VisitRegexLiteralValue(ctx *RegexLiteralValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#regexpConstructorValue.
	VisitRegexpConstructorValue(ctx *RegexpConstructorValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#literalValue.
	VisitLiteralValue(ctx *LiteralValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by MongoShellParser#helperFunction.
	VisitHelperFunction(ctx *HelperFunctionContext) interface{}

	// Visit a parse tree produced by MongoShellParser#objectIdHelper.
	VisitObjectIdHelper(ctx *ObjectIdHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#isoDateHelper.
	VisitIsoDateHelper(ctx *IsoDateHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#dateHelper.
	VisitDateHelper(ctx *DateHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#uuidHelper.
	VisitUuidHelper(ctx *UuidHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#longHelper.
	VisitLongHelper(ctx *LongHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#int32Helper.
	VisitInt32Helper(ctx *Int32HelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#doubleHelper.
	VisitDoubleHelper(ctx *DoubleHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#decimal128Helper.
	VisitDecimal128Helper(ctx *Decimal128HelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#timestampDocHelper.
	VisitTimestampDocHelper(ctx *TimestampDocHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#timestampArgsHelper.
	VisitTimestampArgsHelper(ctx *TimestampArgsHelperContext) interface{}

	// Visit a parse tree produced by MongoShellParser#regExpConstructor.
	VisitRegExpConstructor(ctx *RegExpConstructorContext) interface{}

	// Visit a parse tree produced by MongoShellParser#stringLiteralValue.
	VisitStringLiteralValue(ctx *StringLiteralValueContext) interface{}

	// Visit a parse tree produced by MongoShellParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#trueLiteral.
	VisitTrueLiteral(ctx *TrueLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#falseLiteral.
	VisitFalseLiteral(ctx *FalseLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by MongoShellParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
