// Code generated from ANTLRv4Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ANTLRv4Parser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ANTLRv4Parser.
type ANTLRv4ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ANTLRv4Parser#grammarSpec.
	VisitGrammarSpec(ctx *GrammarSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#grammarDecl.
	VisitGrammarDecl(ctx *GrammarDeclContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#grammarType.
	VisitGrammarType(ctx *GrammarTypeContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#prequelConstruct.
	VisitPrequelConstruct(ctx *PrequelConstructContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#optionsSpec.
	VisitOptionsSpec(ctx *OptionsSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#optionValue.
	VisitOptionValue(ctx *OptionValueContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#delegateGrammars.
	VisitDelegateGrammars(ctx *DelegateGrammarsContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#delegateGrammar.
	VisitDelegateGrammar(ctx *DelegateGrammarContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#tokensSpec.
	VisitTokensSpec(ctx *TokensSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#channelsSpec.
	VisitChannelsSpec(ctx *ChannelsSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#idList.
	VisitIdList(ctx *IdListContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#action_.
	VisitAction_(ctx *Action_Context) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#actionScopeName.
	VisitActionScopeName(ctx *ActionScopeNameContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#actionBlock.
	VisitActionBlock(ctx *ActionBlockContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#argActionBlock.
	VisitArgActionBlock(ctx *ArgActionBlockContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#modeSpec.
	VisitModeSpec(ctx *ModeSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#rules.
	VisitRules(ctx *RulesContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleSpec.
	VisitRuleSpec(ctx *RuleSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#parserRuleSpec.
	VisitParserRuleSpec(ctx *ParserRuleSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#exceptionGroup.
	VisitExceptionGroup(ctx *ExceptionGroupContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#exceptionHandler.
	VisitExceptionHandler(ctx *ExceptionHandlerContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#finallyClause.
	VisitFinallyClause(ctx *FinallyClauseContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#rulePrequel.
	VisitRulePrequel(ctx *RulePrequelContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleReturns.
	VisitRuleReturns(ctx *RuleReturnsContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#throwsSpec.
	VisitThrowsSpec(ctx *ThrowsSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#localsSpec.
	VisitLocalsSpec(ctx *LocalsSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleAction.
	VisitRuleAction(ctx *RuleActionContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleModifiers.
	VisitRuleModifiers(ctx *RuleModifiersContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleModifier.
	VisitRuleModifier(ctx *RuleModifierContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleBlock.
	VisitRuleBlock(ctx *RuleBlockContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleAltList.
	VisitRuleAltList(ctx *RuleAltListContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#labeledAlt.
	VisitLabeledAlt(ctx *LabeledAltContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerRuleSpec.
	VisitLexerRuleSpec(ctx *LexerRuleSpecContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerRuleBlock.
	VisitLexerRuleBlock(ctx *LexerRuleBlockContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerAltList.
	VisitLexerAltList(ctx *LexerAltListContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerAlt.
	VisitLexerAlt(ctx *LexerAltContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerElements.
	VisitLexerElements(ctx *LexerElementsContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerElement.
	VisitLexerElement(ctx *LexerElementContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerBlock.
	VisitLexerBlock(ctx *LexerBlockContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerCommands.
	VisitLexerCommands(ctx *LexerCommandsContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerCommand.
	VisitLexerCommand(ctx *LexerCommandContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerCommandName.
	VisitLexerCommandName(ctx *LexerCommandNameContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerCommandExpr.
	VisitLexerCommandExpr(ctx *LexerCommandExprContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#altList.
	VisitAltList(ctx *AltListContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#alternative.
	VisitAlternative(ctx *AlternativeContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#predicateOptions.
	VisitPredicateOptions(ctx *PredicateOptionsContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#predicateOption.
	VisitPredicateOption(ctx *PredicateOptionContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#labeledElement.
	VisitLabeledElement(ctx *LabeledElementContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ebnf.
	VisitEbnf(ctx *EbnfContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#blockSuffix.
	VisitBlockSuffix(ctx *BlockSuffixContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ebnfSuffix.
	VisitEbnfSuffix(ctx *EbnfSuffixContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#lexerAtom.
	VisitLexerAtom(ctx *LexerAtomContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#wildcard.
	VisitWildcard(ctx *WildcardContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#notSet.
	VisitNotSet(ctx *NotSetContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#blockSet.
	VisitBlockSet(ctx *BlockSetContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#setElement.
	VisitSetElement(ctx *SetElementContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#ruleref.
	VisitRuleref(ctx *RulerefContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#characterRange.
	VisitCharacterRange(ctx *CharacterRangeContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#terminalDef.
	VisitTerminalDef(ctx *TerminalDefContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#elementOptions.
	VisitElementOptions(ctx *ElementOptionsContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#elementOption.
	VisitElementOption(ctx *ElementOptionContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by ANTLRv4Parser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}
}
