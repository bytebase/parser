// Code generated from ANTLRv4Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ANTLRv4Parser
import "github.com/antlr4-go/antlr/v4"

// ANTLRv4ParserListener is a complete listener for a parse tree produced by ANTLRv4Parser.
type ANTLRv4ParserListener interface {
	antlr.ParseTreeListener

	// EnterGrammarSpec is called when entering the grammarSpec production.
	EnterGrammarSpec(c *GrammarSpecContext)

	// EnterGrammarDecl is called when entering the grammarDecl production.
	EnterGrammarDecl(c *GrammarDeclContext)

	// EnterGrammarType is called when entering the grammarType production.
	EnterGrammarType(c *GrammarTypeContext)

	// EnterPrequelConstruct is called when entering the prequelConstruct production.
	EnterPrequelConstruct(c *PrequelConstructContext)

	// EnterOptionsSpec is called when entering the optionsSpec production.
	EnterOptionsSpec(c *OptionsSpecContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterOptionValue is called when entering the optionValue production.
	EnterOptionValue(c *OptionValueContext)

	// EnterDelegateGrammars is called when entering the delegateGrammars production.
	EnterDelegateGrammars(c *DelegateGrammarsContext)

	// EnterDelegateGrammar is called when entering the delegateGrammar production.
	EnterDelegateGrammar(c *DelegateGrammarContext)

	// EnterTokensSpec is called when entering the tokensSpec production.
	EnterTokensSpec(c *TokensSpecContext)

	// EnterChannelsSpec is called when entering the channelsSpec production.
	EnterChannelsSpec(c *ChannelsSpecContext)

	// EnterIdList is called when entering the idList production.
	EnterIdList(c *IdListContext)

	// EnterAction_ is called when entering the action_ production.
	EnterAction_(c *Action_Context)

	// EnterActionScopeName is called when entering the actionScopeName production.
	EnterActionScopeName(c *ActionScopeNameContext)

	// EnterActionBlock is called when entering the actionBlock production.
	EnterActionBlock(c *ActionBlockContext)

	// EnterArgActionBlock is called when entering the argActionBlock production.
	EnterArgActionBlock(c *ArgActionBlockContext)

	// EnterModeSpec is called when entering the modeSpec production.
	EnterModeSpec(c *ModeSpecContext)

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterRuleSpec is called when entering the ruleSpec production.
	EnterRuleSpec(c *RuleSpecContext)

	// EnterParserRuleSpec is called when entering the parserRuleSpec production.
	EnterParserRuleSpec(c *ParserRuleSpecContext)

	// EnterExceptionGroup is called when entering the exceptionGroup production.
	EnterExceptionGroup(c *ExceptionGroupContext)

	// EnterExceptionHandler is called when entering the exceptionHandler production.
	EnterExceptionHandler(c *ExceptionHandlerContext)

	// EnterFinallyClause is called when entering the finallyClause production.
	EnterFinallyClause(c *FinallyClauseContext)

	// EnterRulePrequel is called when entering the rulePrequel production.
	EnterRulePrequel(c *RulePrequelContext)

	// EnterRuleReturns is called when entering the ruleReturns production.
	EnterRuleReturns(c *RuleReturnsContext)

	// EnterThrowsSpec is called when entering the throwsSpec production.
	EnterThrowsSpec(c *ThrowsSpecContext)

	// EnterLocalsSpec is called when entering the localsSpec production.
	EnterLocalsSpec(c *LocalsSpecContext)

	// EnterRuleAction is called when entering the ruleAction production.
	EnterRuleAction(c *RuleActionContext)

	// EnterRuleModifiers is called when entering the ruleModifiers production.
	EnterRuleModifiers(c *RuleModifiersContext)

	// EnterRuleModifier is called when entering the ruleModifier production.
	EnterRuleModifier(c *RuleModifierContext)

	// EnterRuleBlock is called when entering the ruleBlock production.
	EnterRuleBlock(c *RuleBlockContext)

	// EnterRuleAltList is called when entering the ruleAltList production.
	EnterRuleAltList(c *RuleAltListContext)

	// EnterLabeledAlt is called when entering the labeledAlt production.
	EnterLabeledAlt(c *LabeledAltContext)

	// EnterLexerRuleSpec is called when entering the lexerRuleSpec production.
	EnterLexerRuleSpec(c *LexerRuleSpecContext)

	// EnterLexerRuleBlock is called when entering the lexerRuleBlock production.
	EnterLexerRuleBlock(c *LexerRuleBlockContext)

	// EnterLexerAltList is called when entering the lexerAltList production.
	EnterLexerAltList(c *LexerAltListContext)

	// EnterLexerAlt is called when entering the lexerAlt production.
	EnterLexerAlt(c *LexerAltContext)

	// EnterLexerElements is called when entering the lexerElements production.
	EnterLexerElements(c *LexerElementsContext)

	// EnterLexerElement is called when entering the lexerElement production.
	EnterLexerElement(c *LexerElementContext)

	// EnterLexerBlock is called when entering the lexerBlock production.
	EnterLexerBlock(c *LexerBlockContext)

	// EnterLexerCommands is called when entering the lexerCommands production.
	EnterLexerCommands(c *LexerCommandsContext)

	// EnterLexerCommand is called when entering the lexerCommand production.
	EnterLexerCommand(c *LexerCommandContext)

	// EnterLexerCommandName is called when entering the lexerCommandName production.
	EnterLexerCommandName(c *LexerCommandNameContext)

	// EnterLexerCommandExpr is called when entering the lexerCommandExpr production.
	EnterLexerCommandExpr(c *LexerCommandExprContext)

	// EnterAltList is called when entering the altList production.
	EnterAltList(c *AltListContext)

	// EnterAlternative is called when entering the alternative production.
	EnterAlternative(c *AlternativeContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterPredicateOptions is called when entering the predicateOptions production.
	EnterPredicateOptions(c *PredicateOptionsContext)

	// EnterPredicateOption is called when entering the predicateOption production.
	EnterPredicateOption(c *PredicateOptionContext)

	// EnterLabeledElement is called when entering the labeledElement production.
	EnterLabeledElement(c *LabeledElementContext)

	// EnterEbnf is called when entering the ebnf production.
	EnterEbnf(c *EbnfContext)

	// EnterBlockSuffix is called when entering the blockSuffix production.
	EnterBlockSuffix(c *BlockSuffixContext)

	// EnterEbnfSuffix is called when entering the ebnfSuffix production.
	EnterEbnfSuffix(c *EbnfSuffixContext)

	// EnterLexerAtom is called when entering the lexerAtom production.
	EnterLexerAtom(c *LexerAtomContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterWildcard is called when entering the wildcard production.
	EnterWildcard(c *WildcardContext)

	// EnterNotSet is called when entering the notSet production.
	EnterNotSet(c *NotSetContext)

	// EnterBlockSet is called when entering the blockSet production.
	EnterBlockSet(c *BlockSetContext)

	// EnterSetElement is called when entering the setElement production.
	EnterSetElement(c *SetElementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterRuleref is called when entering the ruleref production.
	EnterRuleref(c *RulerefContext)

	// EnterCharacterRange is called when entering the characterRange production.
	EnterCharacterRange(c *CharacterRangeContext)

	// EnterTerminalDef is called when entering the terminalDef production.
	EnterTerminalDef(c *TerminalDefContext)

	// EnterElementOptions is called when entering the elementOptions production.
	EnterElementOptions(c *ElementOptionsContext)

	// EnterElementOption is called when entering the elementOption production.
	EnterElementOption(c *ElementOptionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterQualifiedIdentifier is called when entering the qualifiedIdentifier production.
	EnterQualifiedIdentifier(c *QualifiedIdentifierContext)

	// ExitGrammarSpec is called when exiting the grammarSpec production.
	ExitGrammarSpec(c *GrammarSpecContext)

	// ExitGrammarDecl is called when exiting the grammarDecl production.
	ExitGrammarDecl(c *GrammarDeclContext)

	// ExitGrammarType is called when exiting the grammarType production.
	ExitGrammarType(c *GrammarTypeContext)

	// ExitPrequelConstruct is called when exiting the prequelConstruct production.
	ExitPrequelConstruct(c *PrequelConstructContext)

	// ExitOptionsSpec is called when exiting the optionsSpec production.
	ExitOptionsSpec(c *OptionsSpecContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitOptionValue is called when exiting the optionValue production.
	ExitOptionValue(c *OptionValueContext)

	// ExitDelegateGrammars is called when exiting the delegateGrammars production.
	ExitDelegateGrammars(c *DelegateGrammarsContext)

	// ExitDelegateGrammar is called when exiting the delegateGrammar production.
	ExitDelegateGrammar(c *DelegateGrammarContext)

	// ExitTokensSpec is called when exiting the tokensSpec production.
	ExitTokensSpec(c *TokensSpecContext)

	// ExitChannelsSpec is called when exiting the channelsSpec production.
	ExitChannelsSpec(c *ChannelsSpecContext)

	// ExitIdList is called when exiting the idList production.
	ExitIdList(c *IdListContext)

	// ExitAction_ is called when exiting the action_ production.
	ExitAction_(c *Action_Context)

	// ExitActionScopeName is called when exiting the actionScopeName production.
	ExitActionScopeName(c *ActionScopeNameContext)

	// ExitActionBlock is called when exiting the actionBlock production.
	ExitActionBlock(c *ActionBlockContext)

	// ExitArgActionBlock is called when exiting the argActionBlock production.
	ExitArgActionBlock(c *ArgActionBlockContext)

	// ExitModeSpec is called when exiting the modeSpec production.
	ExitModeSpec(c *ModeSpecContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitRuleSpec is called when exiting the ruleSpec production.
	ExitRuleSpec(c *RuleSpecContext)

	// ExitParserRuleSpec is called when exiting the parserRuleSpec production.
	ExitParserRuleSpec(c *ParserRuleSpecContext)

	// ExitExceptionGroup is called when exiting the exceptionGroup production.
	ExitExceptionGroup(c *ExceptionGroupContext)

	// ExitExceptionHandler is called when exiting the exceptionHandler production.
	ExitExceptionHandler(c *ExceptionHandlerContext)

	// ExitFinallyClause is called when exiting the finallyClause production.
	ExitFinallyClause(c *FinallyClauseContext)

	// ExitRulePrequel is called when exiting the rulePrequel production.
	ExitRulePrequel(c *RulePrequelContext)

	// ExitRuleReturns is called when exiting the ruleReturns production.
	ExitRuleReturns(c *RuleReturnsContext)

	// ExitThrowsSpec is called when exiting the throwsSpec production.
	ExitThrowsSpec(c *ThrowsSpecContext)

	// ExitLocalsSpec is called when exiting the localsSpec production.
	ExitLocalsSpec(c *LocalsSpecContext)

	// ExitRuleAction is called when exiting the ruleAction production.
	ExitRuleAction(c *RuleActionContext)

	// ExitRuleModifiers is called when exiting the ruleModifiers production.
	ExitRuleModifiers(c *RuleModifiersContext)

	// ExitRuleModifier is called when exiting the ruleModifier production.
	ExitRuleModifier(c *RuleModifierContext)

	// ExitRuleBlock is called when exiting the ruleBlock production.
	ExitRuleBlock(c *RuleBlockContext)

	// ExitRuleAltList is called when exiting the ruleAltList production.
	ExitRuleAltList(c *RuleAltListContext)

	// ExitLabeledAlt is called when exiting the labeledAlt production.
	ExitLabeledAlt(c *LabeledAltContext)

	// ExitLexerRuleSpec is called when exiting the lexerRuleSpec production.
	ExitLexerRuleSpec(c *LexerRuleSpecContext)

	// ExitLexerRuleBlock is called when exiting the lexerRuleBlock production.
	ExitLexerRuleBlock(c *LexerRuleBlockContext)

	// ExitLexerAltList is called when exiting the lexerAltList production.
	ExitLexerAltList(c *LexerAltListContext)

	// ExitLexerAlt is called when exiting the lexerAlt production.
	ExitLexerAlt(c *LexerAltContext)

	// ExitLexerElements is called when exiting the lexerElements production.
	ExitLexerElements(c *LexerElementsContext)

	// ExitLexerElement is called when exiting the lexerElement production.
	ExitLexerElement(c *LexerElementContext)

	// ExitLexerBlock is called when exiting the lexerBlock production.
	ExitLexerBlock(c *LexerBlockContext)

	// ExitLexerCommands is called when exiting the lexerCommands production.
	ExitLexerCommands(c *LexerCommandsContext)

	// ExitLexerCommand is called when exiting the lexerCommand production.
	ExitLexerCommand(c *LexerCommandContext)

	// ExitLexerCommandName is called when exiting the lexerCommandName production.
	ExitLexerCommandName(c *LexerCommandNameContext)

	// ExitLexerCommandExpr is called when exiting the lexerCommandExpr production.
	ExitLexerCommandExpr(c *LexerCommandExprContext)

	// ExitAltList is called when exiting the altList production.
	ExitAltList(c *AltListContext)

	// ExitAlternative is called when exiting the alternative production.
	ExitAlternative(c *AlternativeContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitPredicateOptions is called when exiting the predicateOptions production.
	ExitPredicateOptions(c *PredicateOptionsContext)

	// ExitPredicateOption is called when exiting the predicateOption production.
	ExitPredicateOption(c *PredicateOptionContext)

	// ExitLabeledElement is called when exiting the labeledElement production.
	ExitLabeledElement(c *LabeledElementContext)

	// ExitEbnf is called when exiting the ebnf production.
	ExitEbnf(c *EbnfContext)

	// ExitBlockSuffix is called when exiting the blockSuffix production.
	ExitBlockSuffix(c *BlockSuffixContext)

	// ExitEbnfSuffix is called when exiting the ebnfSuffix production.
	ExitEbnfSuffix(c *EbnfSuffixContext)

	// ExitLexerAtom is called when exiting the lexerAtom production.
	ExitLexerAtom(c *LexerAtomContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitWildcard is called when exiting the wildcard production.
	ExitWildcard(c *WildcardContext)

	// ExitNotSet is called when exiting the notSet production.
	ExitNotSet(c *NotSetContext)

	// ExitBlockSet is called when exiting the blockSet production.
	ExitBlockSet(c *BlockSetContext)

	// ExitSetElement is called when exiting the setElement production.
	ExitSetElement(c *SetElementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitRuleref is called when exiting the ruleref production.
	ExitRuleref(c *RulerefContext)

	// ExitCharacterRange is called when exiting the characterRange production.
	ExitCharacterRange(c *CharacterRangeContext)

	// ExitTerminalDef is called when exiting the terminalDef production.
	ExitTerminalDef(c *TerminalDefContext)

	// ExitElementOptions is called when exiting the elementOptions production.
	ExitElementOptions(c *ElementOptionsContext)

	// ExitElementOption is called when exiting the elementOption production.
	ExitElementOption(c *ElementOptionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitQualifiedIdentifier is called when exiting the qualifiedIdentifier production.
	ExitQualifiedIdentifier(c *QualifiedIdentifierContext)
}
