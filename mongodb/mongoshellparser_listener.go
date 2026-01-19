// Code generated from MongoShellParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package mongodb // MongoShellParser
import "github.com/antlr4-go/antlr/v4"

// MongoShellParserListener is a complete listener for a parse tree produced by MongoShellParser.
type MongoShellParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterShowDatabases is called when entering the showDatabases production.
	EnterShowDatabases(c *ShowDatabasesContext)

	// EnterShowCollections is called when entering the showCollections production.
	EnterShowCollections(c *ShowCollectionsContext)

	// EnterGetCollectionNames is called when entering the getCollectionNames production.
	EnterGetCollectionNames(c *GetCollectionNamesContext)

	// EnterGetCollectionInfos is called when entering the getCollectionInfos production.
	EnterGetCollectionInfos(c *GetCollectionInfosContext)

	// EnterCollectionOperation is called when entering the collectionOperation production.
	EnterCollectionOperation(c *CollectionOperationContext)

	// EnterDotAccess is called when entering the dotAccess production.
	EnterDotAccess(c *DotAccessContext)

	// EnterBracketAccess is called when entering the bracketAccess production.
	EnterBracketAccess(c *BracketAccessContext)

	// EnterGetCollectionAccess is called when entering the getCollectionAccess production.
	EnterGetCollectionAccess(c *GetCollectionAccessContext)

	// EnterMethodChain is called when entering the methodChain production.
	EnterMethodChain(c *MethodChainContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterFindMethod is called when entering the findMethod production.
	EnterFindMethod(c *FindMethodContext)

	// EnterFindOneMethod is called when entering the findOneMethod production.
	EnterFindOneMethod(c *FindOneMethodContext)

	// EnterCountDocumentsMethod is called when entering the countDocumentsMethod production.
	EnterCountDocumentsMethod(c *CountDocumentsMethodContext)

	// EnterEstimatedDocumentCountMethod is called when entering the estimatedDocumentCountMethod production.
	EnterEstimatedDocumentCountMethod(c *EstimatedDocumentCountMethodContext)

	// EnterDistinctMethod is called when entering the distinctMethod production.
	EnterDistinctMethod(c *DistinctMethodContext)

	// EnterAggregateMethod is called when entering the aggregateMethod production.
	EnterAggregateMethod(c *AggregateMethodContext)

	// EnterGetIndexesMethod is called when entering the getIndexesMethod production.
	EnterGetIndexesMethod(c *GetIndexesMethodContext)

	// EnterSortMethod is called when entering the sortMethod production.
	EnterSortMethod(c *SortMethodContext)

	// EnterLimitMethod is called when entering the limitMethod production.
	EnterLimitMethod(c *LimitMethodContext)

	// EnterSkipMethod is called when entering the skipMethod production.
	EnterSkipMethod(c *SkipMethodContext)

	// EnterCountMethod is called when entering the countMethod production.
	EnterCountMethod(c *CountMethodContext)

	// EnterProjectionMethod is called when entering the projectionMethod production.
	EnterProjectionMethod(c *ProjectionMethodContext)

	// EnterGenericMethod is called when entering the genericMethod production.
	EnterGenericMethod(c *GenericMethodContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterUnquotedKey is called when entering the unquotedKey production.
	EnterUnquotedKey(c *UnquotedKeyContext)

	// EnterQuotedKey is called when entering the quotedKey production.
	EnterQuotedKey(c *QuotedKeyContext)

	// EnterDocumentValue is called when entering the documentValue production.
	EnterDocumentValue(c *DocumentValueContext)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterHelperValue is called when entering the helperValue production.
	EnterHelperValue(c *HelperValueContext)

	// EnterRegexLiteralValue is called when entering the regexLiteralValue production.
	EnterRegexLiteralValue(c *RegexLiteralValueContext)

	// EnterRegexpConstructorValue is called when entering the regexpConstructorValue production.
	EnterRegexpConstructorValue(c *RegexpConstructorValueContext)

	// EnterLiteralValue is called when entering the literalValue production.
	EnterLiteralValue(c *LiteralValueContext)

	// EnterNewKeywordValue is called when entering the newKeywordValue production.
	EnterNewKeywordValue(c *NewKeywordValueContext)

	// EnterNewKeywordError is called when entering the newKeywordError production.
	EnterNewKeywordError(c *NewKeywordErrorContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterHelperFunction is called when entering the helperFunction production.
	EnterHelperFunction(c *HelperFunctionContext)

	// EnterObjectIdHelper is called when entering the objectIdHelper production.
	EnterObjectIdHelper(c *ObjectIdHelperContext)

	// EnterIsoDateHelper is called when entering the isoDateHelper production.
	EnterIsoDateHelper(c *IsoDateHelperContext)

	// EnterDateHelper is called when entering the dateHelper production.
	EnterDateHelper(c *DateHelperContext)

	// EnterUuidHelper is called when entering the uuidHelper production.
	EnterUuidHelper(c *UuidHelperContext)

	// EnterLongHelper is called when entering the longHelper production.
	EnterLongHelper(c *LongHelperContext)

	// EnterInt32Helper is called when entering the int32Helper production.
	EnterInt32Helper(c *Int32HelperContext)

	// EnterDoubleHelper is called when entering the doubleHelper production.
	EnterDoubleHelper(c *DoubleHelperContext)

	// EnterDecimal128Helper is called when entering the decimal128Helper production.
	EnterDecimal128Helper(c *Decimal128HelperContext)

	// EnterTimestampDocHelper is called when entering the timestampDocHelper production.
	EnterTimestampDocHelper(c *TimestampDocHelperContext)

	// EnterTimestampArgsHelper is called when entering the timestampArgsHelper production.
	EnterTimestampArgsHelper(c *TimestampArgsHelperContext)

	// EnterRegExpConstructor is called when entering the regExpConstructor production.
	EnterRegExpConstructor(c *RegExpConstructorContext)

	// EnterStringLiteralValue is called when entering the stringLiteralValue production.
	EnterStringLiteralValue(c *StringLiteralValueContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterTrueLiteral is called when entering the trueLiteral production.
	EnterTrueLiteral(c *TrueLiteralContext)

	// EnterFalseLiteral is called when entering the falseLiteral production.
	EnterFalseLiteral(c *FalseLiteralContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitShowDatabases is called when exiting the showDatabases production.
	ExitShowDatabases(c *ShowDatabasesContext)

	// ExitShowCollections is called when exiting the showCollections production.
	ExitShowCollections(c *ShowCollectionsContext)

	// ExitGetCollectionNames is called when exiting the getCollectionNames production.
	ExitGetCollectionNames(c *GetCollectionNamesContext)

	// ExitGetCollectionInfos is called when exiting the getCollectionInfos production.
	ExitGetCollectionInfos(c *GetCollectionInfosContext)

	// ExitCollectionOperation is called when exiting the collectionOperation production.
	ExitCollectionOperation(c *CollectionOperationContext)

	// ExitDotAccess is called when exiting the dotAccess production.
	ExitDotAccess(c *DotAccessContext)

	// ExitBracketAccess is called when exiting the bracketAccess production.
	ExitBracketAccess(c *BracketAccessContext)

	// ExitGetCollectionAccess is called when exiting the getCollectionAccess production.
	ExitGetCollectionAccess(c *GetCollectionAccessContext)

	// ExitMethodChain is called when exiting the methodChain production.
	ExitMethodChain(c *MethodChainContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitFindMethod is called when exiting the findMethod production.
	ExitFindMethod(c *FindMethodContext)

	// ExitFindOneMethod is called when exiting the findOneMethod production.
	ExitFindOneMethod(c *FindOneMethodContext)

	// ExitCountDocumentsMethod is called when exiting the countDocumentsMethod production.
	ExitCountDocumentsMethod(c *CountDocumentsMethodContext)

	// ExitEstimatedDocumentCountMethod is called when exiting the estimatedDocumentCountMethod production.
	ExitEstimatedDocumentCountMethod(c *EstimatedDocumentCountMethodContext)

	// ExitDistinctMethod is called when exiting the distinctMethod production.
	ExitDistinctMethod(c *DistinctMethodContext)

	// ExitAggregateMethod is called when exiting the aggregateMethod production.
	ExitAggregateMethod(c *AggregateMethodContext)

	// ExitGetIndexesMethod is called when exiting the getIndexesMethod production.
	ExitGetIndexesMethod(c *GetIndexesMethodContext)

	// ExitSortMethod is called when exiting the sortMethod production.
	ExitSortMethod(c *SortMethodContext)

	// ExitLimitMethod is called when exiting the limitMethod production.
	ExitLimitMethod(c *LimitMethodContext)

	// ExitSkipMethod is called when exiting the skipMethod production.
	ExitSkipMethod(c *SkipMethodContext)

	// ExitCountMethod is called when exiting the countMethod production.
	ExitCountMethod(c *CountMethodContext)

	// ExitProjectionMethod is called when exiting the projectionMethod production.
	ExitProjectionMethod(c *ProjectionMethodContext)

	// ExitGenericMethod is called when exiting the genericMethod production.
	ExitGenericMethod(c *GenericMethodContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitUnquotedKey is called when exiting the unquotedKey production.
	ExitUnquotedKey(c *UnquotedKeyContext)

	// ExitQuotedKey is called when exiting the quotedKey production.
	ExitQuotedKey(c *QuotedKeyContext)

	// ExitDocumentValue is called when exiting the documentValue production.
	ExitDocumentValue(c *DocumentValueContext)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitHelperValue is called when exiting the helperValue production.
	ExitHelperValue(c *HelperValueContext)

	// ExitRegexLiteralValue is called when exiting the regexLiteralValue production.
	ExitRegexLiteralValue(c *RegexLiteralValueContext)

	// ExitRegexpConstructorValue is called when exiting the regexpConstructorValue production.
	ExitRegexpConstructorValue(c *RegexpConstructorValueContext)

	// ExitLiteralValue is called when exiting the literalValue production.
	ExitLiteralValue(c *LiteralValueContext)

	// ExitNewKeywordValue is called when exiting the newKeywordValue production.
	ExitNewKeywordValue(c *NewKeywordValueContext)

	// ExitNewKeywordError is called when exiting the newKeywordError production.
	ExitNewKeywordError(c *NewKeywordErrorContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitHelperFunction is called when exiting the helperFunction production.
	ExitHelperFunction(c *HelperFunctionContext)

	// ExitObjectIdHelper is called when exiting the objectIdHelper production.
	ExitObjectIdHelper(c *ObjectIdHelperContext)

	// ExitIsoDateHelper is called when exiting the isoDateHelper production.
	ExitIsoDateHelper(c *IsoDateHelperContext)

	// ExitDateHelper is called when exiting the dateHelper production.
	ExitDateHelper(c *DateHelperContext)

	// ExitUuidHelper is called when exiting the uuidHelper production.
	ExitUuidHelper(c *UuidHelperContext)

	// ExitLongHelper is called when exiting the longHelper production.
	ExitLongHelper(c *LongHelperContext)

	// ExitInt32Helper is called when exiting the int32Helper production.
	ExitInt32Helper(c *Int32HelperContext)

	// ExitDoubleHelper is called when exiting the doubleHelper production.
	ExitDoubleHelper(c *DoubleHelperContext)

	// ExitDecimal128Helper is called when exiting the decimal128Helper production.
	ExitDecimal128Helper(c *Decimal128HelperContext)

	// ExitTimestampDocHelper is called when exiting the timestampDocHelper production.
	ExitTimestampDocHelper(c *TimestampDocHelperContext)

	// ExitTimestampArgsHelper is called when exiting the timestampArgsHelper production.
	ExitTimestampArgsHelper(c *TimestampArgsHelperContext)

	// ExitRegExpConstructor is called when exiting the regExpConstructor production.
	ExitRegExpConstructor(c *RegExpConstructorContext)

	// ExitStringLiteralValue is called when exiting the stringLiteralValue production.
	ExitStringLiteralValue(c *StringLiteralValueContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitTrueLiteral is called when exiting the trueLiteral production.
	ExitTrueLiteral(c *TrueLiteralContext)

	// ExitFalseLiteral is called when exiting the falseLiteral production.
	ExitFalseLiteral(c *FalseLiteralContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
