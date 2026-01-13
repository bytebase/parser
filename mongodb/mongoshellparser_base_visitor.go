// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

type BaseMongoShellParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMongoShellParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitShowDatabases(ctx *ShowDatabasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitShowCollections(ctx *ShowCollectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetCollectionNames(ctx *GetCollectionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitCollectionOperation(ctx *CollectionOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDotAccess(ctx *DotAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitBracketAccess(ctx *BracketAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGetCollectionAccess(ctx *GetCollectionAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMethodChain(ctx *MethodChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitMethodCall(ctx *MethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindMethod(ctx *FindMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFindOneMethod(ctx *FindOneMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitSortMethod(ctx *SortMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitLimitMethod(ctx *LimitMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitSkipMethod(ctx *SkipMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitProjectionMethod(ctx *ProjectionMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitGenericMethod(ctx *GenericMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUnquotedKey(ctx *UnquotedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitQuotedKey(ctx *QuotedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDocumentValue(ctx *DocumentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHelperValue(ctx *HelperValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRegexLiteralValue(ctx *RegexLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRegexpConstructorValue(ctx *RegexpConstructorValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitLiteralValue(ctx *LiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitHelperFunction(ctx *HelperFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitObjectIdHelper(ctx *ObjectIdHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitIsoDateHelper(ctx *IsoDateHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDateHelper(ctx *DateHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitUuidHelper(ctx *UuidHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitLongHelper(ctx *LongHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitInt32Helper(ctx *Int32HelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDoubleHelper(ctx *DoubleHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitDecimal128Helper(ctx *Decimal128HelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTimestampDocHelper(ctx *TimestampDocHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTimestampArgsHelper(ctx *TimestampArgsHelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitRegExpConstructor(ctx *RegExpConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStringLiteralValue(ctx *StringLiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitTrueLiteral(ctx *TrueLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitFalseLiteral(ctx *FalseLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMongoShellParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
