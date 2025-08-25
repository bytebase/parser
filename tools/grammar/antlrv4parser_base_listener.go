// Code generated from ANTLRv4Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ANTLRv4Parser
import "github.com/antlr4-go/antlr/v4"

// BaseANTLRv4ParserListener is a complete listener for a parse tree produced by ANTLRv4Parser.
type BaseANTLRv4ParserListener struct{}

var _ ANTLRv4ParserListener = &BaseANTLRv4ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseANTLRv4ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseANTLRv4ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseANTLRv4ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseANTLRv4ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGrammarSpec is called when production grammarSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterGrammarSpec(ctx *GrammarSpecContext) {}

// ExitGrammarSpec is called when production grammarSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitGrammarSpec(ctx *GrammarSpecContext) {}

// EnterGrammarDecl is called when production grammarDecl is entered.
func (s *BaseANTLRv4ParserListener) EnterGrammarDecl(ctx *GrammarDeclContext) {}

// ExitGrammarDecl is called when production grammarDecl is exited.
func (s *BaseANTLRv4ParserListener) ExitGrammarDecl(ctx *GrammarDeclContext) {}

// EnterGrammarType is called when production grammarType is entered.
func (s *BaseANTLRv4ParserListener) EnterGrammarType(ctx *GrammarTypeContext) {}

// ExitGrammarType is called when production grammarType is exited.
func (s *BaseANTLRv4ParserListener) ExitGrammarType(ctx *GrammarTypeContext) {}

// EnterPrequelConstruct is called when production prequelConstruct is entered.
func (s *BaseANTLRv4ParserListener) EnterPrequelConstruct(ctx *PrequelConstructContext) {}

// ExitPrequelConstruct is called when production prequelConstruct is exited.
func (s *BaseANTLRv4ParserListener) ExitPrequelConstruct(ctx *PrequelConstructContext) {}

// EnterOptionsSpec is called when production optionsSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterOptionsSpec(ctx *OptionsSpecContext) {}

// ExitOptionsSpec is called when production optionsSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitOptionsSpec(ctx *OptionsSpecContext) {}

// EnterOption is called when production option is entered.
func (s *BaseANTLRv4ParserListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseANTLRv4ParserListener) ExitOption(ctx *OptionContext) {}

// EnterOptionValue is called when production optionValue is entered.
func (s *BaseANTLRv4ParserListener) EnterOptionValue(ctx *OptionValueContext) {}

// ExitOptionValue is called when production optionValue is exited.
func (s *BaseANTLRv4ParserListener) ExitOptionValue(ctx *OptionValueContext) {}

// EnterDelegateGrammars is called when production delegateGrammars is entered.
func (s *BaseANTLRv4ParserListener) EnterDelegateGrammars(ctx *DelegateGrammarsContext) {}

// ExitDelegateGrammars is called when production delegateGrammars is exited.
func (s *BaseANTLRv4ParserListener) ExitDelegateGrammars(ctx *DelegateGrammarsContext) {}

// EnterDelegateGrammar is called when production delegateGrammar is entered.
func (s *BaseANTLRv4ParserListener) EnterDelegateGrammar(ctx *DelegateGrammarContext) {}

// ExitDelegateGrammar is called when production delegateGrammar is exited.
func (s *BaseANTLRv4ParserListener) ExitDelegateGrammar(ctx *DelegateGrammarContext) {}

// EnterTokensSpec is called when production tokensSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterTokensSpec(ctx *TokensSpecContext) {}

// ExitTokensSpec is called when production tokensSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitTokensSpec(ctx *TokensSpecContext) {}

// EnterChannelsSpec is called when production channelsSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterChannelsSpec(ctx *ChannelsSpecContext) {}

// ExitChannelsSpec is called when production channelsSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitChannelsSpec(ctx *ChannelsSpecContext) {}

// EnterIdList is called when production idList is entered.
func (s *BaseANTLRv4ParserListener) EnterIdList(ctx *IdListContext) {}

// ExitIdList is called when production idList is exited.
func (s *BaseANTLRv4ParserListener) ExitIdList(ctx *IdListContext) {}

// EnterAction_ is called when production action_ is entered.
func (s *BaseANTLRv4ParserListener) EnterAction_(ctx *Action_Context) {}

// ExitAction_ is called when production action_ is exited.
func (s *BaseANTLRv4ParserListener) ExitAction_(ctx *Action_Context) {}

// EnterActionScopeName is called when production actionScopeName is entered.
func (s *BaseANTLRv4ParserListener) EnterActionScopeName(ctx *ActionScopeNameContext) {}

// ExitActionScopeName is called when production actionScopeName is exited.
func (s *BaseANTLRv4ParserListener) ExitActionScopeName(ctx *ActionScopeNameContext) {}

// EnterActionBlock is called when production actionBlock is entered.
func (s *BaseANTLRv4ParserListener) EnterActionBlock(ctx *ActionBlockContext) {}

// ExitActionBlock is called when production actionBlock is exited.
func (s *BaseANTLRv4ParserListener) ExitActionBlock(ctx *ActionBlockContext) {}

// EnterArgActionBlock is called when production argActionBlock is entered.
func (s *BaseANTLRv4ParserListener) EnterArgActionBlock(ctx *ArgActionBlockContext) {}

// ExitArgActionBlock is called when production argActionBlock is exited.
func (s *BaseANTLRv4ParserListener) ExitArgActionBlock(ctx *ArgActionBlockContext) {}

// EnterModeSpec is called when production modeSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterModeSpec(ctx *ModeSpecContext) {}

// ExitModeSpec is called when production modeSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitModeSpec(ctx *ModeSpecContext) {}

// EnterRules is called when production rules is entered.
func (s *BaseANTLRv4ParserListener) EnterRules(ctx *RulesContext) {}

// ExitRules is called when production rules is exited.
func (s *BaseANTLRv4ParserListener) ExitRules(ctx *RulesContext) {}

// EnterRuleSpec is called when production ruleSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleSpec(ctx *RuleSpecContext) {}

// ExitRuleSpec is called when production ruleSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleSpec(ctx *RuleSpecContext) {}

// EnterParserRuleSpec is called when production parserRuleSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterParserRuleSpec(ctx *ParserRuleSpecContext) {}

// ExitParserRuleSpec is called when production parserRuleSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitParserRuleSpec(ctx *ParserRuleSpecContext) {}

// EnterExceptionGroup is called when production exceptionGroup is entered.
func (s *BaseANTLRv4ParserListener) EnterExceptionGroup(ctx *ExceptionGroupContext) {}

// ExitExceptionGroup is called when production exceptionGroup is exited.
func (s *BaseANTLRv4ParserListener) ExitExceptionGroup(ctx *ExceptionGroupContext) {}

// EnterExceptionHandler is called when production exceptionHandler is entered.
func (s *BaseANTLRv4ParserListener) EnterExceptionHandler(ctx *ExceptionHandlerContext) {}

// ExitExceptionHandler is called when production exceptionHandler is exited.
func (s *BaseANTLRv4ParserListener) ExitExceptionHandler(ctx *ExceptionHandlerContext) {}

// EnterFinallyClause is called when production finallyClause is entered.
func (s *BaseANTLRv4ParserListener) EnterFinallyClause(ctx *FinallyClauseContext) {}

// ExitFinallyClause is called when production finallyClause is exited.
func (s *BaseANTLRv4ParserListener) ExitFinallyClause(ctx *FinallyClauseContext) {}

// EnterRulePrequel is called when production rulePrequel is entered.
func (s *BaseANTLRv4ParserListener) EnterRulePrequel(ctx *RulePrequelContext) {}

// ExitRulePrequel is called when production rulePrequel is exited.
func (s *BaseANTLRv4ParserListener) ExitRulePrequel(ctx *RulePrequelContext) {}

// EnterRuleReturns is called when production ruleReturns is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleReturns(ctx *RuleReturnsContext) {}

// ExitRuleReturns is called when production ruleReturns is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleReturns(ctx *RuleReturnsContext) {}

// EnterThrowsSpec is called when production throwsSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterThrowsSpec(ctx *ThrowsSpecContext) {}

// ExitThrowsSpec is called when production throwsSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitThrowsSpec(ctx *ThrowsSpecContext) {}

// EnterLocalsSpec is called when production localsSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterLocalsSpec(ctx *LocalsSpecContext) {}

// ExitLocalsSpec is called when production localsSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitLocalsSpec(ctx *LocalsSpecContext) {}

// EnterRuleAction is called when production ruleAction is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleAction(ctx *RuleActionContext) {}

// ExitRuleAction is called when production ruleAction is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleAction(ctx *RuleActionContext) {}

// EnterRuleModifiers is called when production ruleModifiers is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleModifiers(ctx *RuleModifiersContext) {}

// ExitRuleModifiers is called when production ruleModifiers is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleModifiers(ctx *RuleModifiersContext) {}

// EnterRuleModifier is called when production ruleModifier is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleModifier(ctx *RuleModifierContext) {}

// ExitRuleModifier is called when production ruleModifier is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleModifier(ctx *RuleModifierContext) {}

// EnterRuleBlock is called when production ruleBlock is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleBlock(ctx *RuleBlockContext) {}

// ExitRuleBlock is called when production ruleBlock is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleBlock(ctx *RuleBlockContext) {}

// EnterRuleAltList is called when production ruleAltList is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleAltList(ctx *RuleAltListContext) {}

// ExitRuleAltList is called when production ruleAltList is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleAltList(ctx *RuleAltListContext) {}

// EnterLabeledAlt is called when production labeledAlt is entered.
func (s *BaseANTLRv4ParserListener) EnterLabeledAlt(ctx *LabeledAltContext) {}

// ExitLabeledAlt is called when production labeledAlt is exited.
func (s *BaseANTLRv4ParserListener) ExitLabeledAlt(ctx *LabeledAltContext) {}

// EnterLexerRuleSpec is called when production lexerRuleSpec is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerRuleSpec(ctx *LexerRuleSpecContext) {}

// ExitLexerRuleSpec is called when production lexerRuleSpec is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerRuleSpec(ctx *LexerRuleSpecContext) {}

// EnterLexerRuleBlock is called when production lexerRuleBlock is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerRuleBlock(ctx *LexerRuleBlockContext) {}

// ExitLexerRuleBlock is called when production lexerRuleBlock is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerRuleBlock(ctx *LexerRuleBlockContext) {}

// EnterLexerAltList is called when production lexerAltList is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerAltList(ctx *LexerAltListContext) {}

// ExitLexerAltList is called when production lexerAltList is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerAltList(ctx *LexerAltListContext) {}

// EnterLexerAlt is called when production lexerAlt is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerAlt(ctx *LexerAltContext) {}

// ExitLexerAlt is called when production lexerAlt is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerAlt(ctx *LexerAltContext) {}

// EnterLexerElements is called when production lexerElements is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerElements(ctx *LexerElementsContext) {}

// ExitLexerElements is called when production lexerElements is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerElements(ctx *LexerElementsContext) {}

// EnterLexerElement is called when production lexerElement is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerElement(ctx *LexerElementContext) {}

// ExitLexerElement is called when production lexerElement is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerElement(ctx *LexerElementContext) {}

// EnterLexerBlock is called when production lexerBlock is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerBlock(ctx *LexerBlockContext) {}

// ExitLexerBlock is called when production lexerBlock is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerBlock(ctx *LexerBlockContext) {}

// EnterLexerCommands is called when production lexerCommands is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerCommands(ctx *LexerCommandsContext) {}

// ExitLexerCommands is called when production lexerCommands is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerCommands(ctx *LexerCommandsContext) {}

// EnterLexerCommand is called when production lexerCommand is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerCommand(ctx *LexerCommandContext) {}

// ExitLexerCommand is called when production lexerCommand is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerCommand(ctx *LexerCommandContext) {}

// EnterLexerCommandName is called when production lexerCommandName is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerCommandName(ctx *LexerCommandNameContext) {}

// ExitLexerCommandName is called when production lexerCommandName is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerCommandName(ctx *LexerCommandNameContext) {}

// EnterLexerCommandExpr is called when production lexerCommandExpr is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerCommandExpr(ctx *LexerCommandExprContext) {}

// ExitLexerCommandExpr is called when production lexerCommandExpr is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerCommandExpr(ctx *LexerCommandExprContext) {}

// EnterAltList is called when production altList is entered.
func (s *BaseANTLRv4ParserListener) EnterAltList(ctx *AltListContext) {}

// ExitAltList is called when production altList is exited.
func (s *BaseANTLRv4ParserListener) ExitAltList(ctx *AltListContext) {}

// EnterAlternative is called when production alternative is entered.
func (s *BaseANTLRv4ParserListener) EnterAlternative(ctx *AlternativeContext) {}

// ExitAlternative is called when production alternative is exited.
func (s *BaseANTLRv4ParserListener) ExitAlternative(ctx *AlternativeContext) {}

// EnterElement is called when production element is entered.
func (s *BaseANTLRv4ParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseANTLRv4ParserListener) ExitElement(ctx *ElementContext) {}

// EnterPredicateOptions is called when production predicateOptions is entered.
func (s *BaseANTLRv4ParserListener) EnterPredicateOptions(ctx *PredicateOptionsContext) {}

// ExitPredicateOptions is called when production predicateOptions is exited.
func (s *BaseANTLRv4ParserListener) ExitPredicateOptions(ctx *PredicateOptionsContext) {}

// EnterPredicateOption is called when production predicateOption is entered.
func (s *BaseANTLRv4ParserListener) EnterPredicateOption(ctx *PredicateOptionContext) {}

// ExitPredicateOption is called when production predicateOption is exited.
func (s *BaseANTLRv4ParserListener) ExitPredicateOption(ctx *PredicateOptionContext) {}

// EnterLabeledElement is called when production labeledElement is entered.
func (s *BaseANTLRv4ParserListener) EnterLabeledElement(ctx *LabeledElementContext) {}

// ExitLabeledElement is called when production labeledElement is exited.
func (s *BaseANTLRv4ParserListener) ExitLabeledElement(ctx *LabeledElementContext) {}

// EnterEbnf is called when production ebnf is entered.
func (s *BaseANTLRv4ParserListener) EnterEbnf(ctx *EbnfContext) {}

// ExitEbnf is called when production ebnf is exited.
func (s *BaseANTLRv4ParserListener) ExitEbnf(ctx *EbnfContext) {}

// EnterBlockSuffix is called when production blockSuffix is entered.
func (s *BaseANTLRv4ParserListener) EnterBlockSuffix(ctx *BlockSuffixContext) {}

// ExitBlockSuffix is called when production blockSuffix is exited.
func (s *BaseANTLRv4ParserListener) ExitBlockSuffix(ctx *BlockSuffixContext) {}

// EnterEbnfSuffix is called when production ebnfSuffix is entered.
func (s *BaseANTLRv4ParserListener) EnterEbnfSuffix(ctx *EbnfSuffixContext) {}

// ExitEbnfSuffix is called when production ebnfSuffix is exited.
func (s *BaseANTLRv4ParserListener) ExitEbnfSuffix(ctx *EbnfSuffixContext) {}

// EnterLexerAtom is called when production lexerAtom is entered.
func (s *BaseANTLRv4ParserListener) EnterLexerAtom(ctx *LexerAtomContext) {}

// ExitLexerAtom is called when production lexerAtom is exited.
func (s *BaseANTLRv4ParserListener) ExitLexerAtom(ctx *LexerAtomContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseANTLRv4ParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseANTLRv4ParserListener) ExitAtom(ctx *AtomContext) {}

// EnterWildcard is called when production wildcard is entered.
func (s *BaseANTLRv4ParserListener) EnterWildcard(ctx *WildcardContext) {}

// ExitWildcard is called when production wildcard is exited.
func (s *BaseANTLRv4ParserListener) ExitWildcard(ctx *WildcardContext) {}

// EnterNotSet is called when production notSet is entered.
func (s *BaseANTLRv4ParserListener) EnterNotSet(ctx *NotSetContext) {}

// ExitNotSet is called when production notSet is exited.
func (s *BaseANTLRv4ParserListener) ExitNotSet(ctx *NotSetContext) {}

// EnterBlockSet is called when production blockSet is entered.
func (s *BaseANTLRv4ParserListener) EnterBlockSet(ctx *BlockSetContext) {}

// ExitBlockSet is called when production blockSet is exited.
func (s *BaseANTLRv4ParserListener) ExitBlockSet(ctx *BlockSetContext) {}

// EnterSetElement is called when production setElement is entered.
func (s *BaseANTLRv4ParserListener) EnterSetElement(ctx *SetElementContext) {}

// ExitSetElement is called when production setElement is exited.
func (s *BaseANTLRv4ParserListener) ExitSetElement(ctx *SetElementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseANTLRv4ParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseANTLRv4ParserListener) ExitBlock(ctx *BlockContext) {}

// EnterRuleref is called when production ruleref is entered.
func (s *BaseANTLRv4ParserListener) EnterRuleref(ctx *RulerefContext) {}

// ExitRuleref is called when production ruleref is exited.
func (s *BaseANTLRv4ParserListener) ExitRuleref(ctx *RulerefContext) {}

// EnterCharacterRange is called when production characterRange is entered.
func (s *BaseANTLRv4ParserListener) EnterCharacterRange(ctx *CharacterRangeContext) {}

// ExitCharacterRange is called when production characterRange is exited.
func (s *BaseANTLRv4ParserListener) ExitCharacterRange(ctx *CharacterRangeContext) {}

// EnterTerminalDef is called when production terminalDef is entered.
func (s *BaseANTLRv4ParserListener) EnterTerminalDef(ctx *TerminalDefContext) {}

// ExitTerminalDef is called when production terminalDef is exited.
func (s *BaseANTLRv4ParserListener) ExitTerminalDef(ctx *TerminalDefContext) {}

// EnterElementOptions is called when production elementOptions is entered.
func (s *BaseANTLRv4ParserListener) EnterElementOptions(ctx *ElementOptionsContext) {}

// ExitElementOptions is called when production elementOptions is exited.
func (s *BaseANTLRv4ParserListener) ExitElementOptions(ctx *ElementOptionsContext) {}

// EnterElementOption is called when production elementOption is entered.
func (s *BaseANTLRv4ParserListener) EnterElementOption(ctx *ElementOptionContext) {}

// ExitElementOption is called when production elementOption is exited.
func (s *BaseANTLRv4ParserListener) ExitElementOption(ctx *ElementOptionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseANTLRv4ParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseANTLRv4ParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseANTLRv4ParserListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseANTLRv4ParserListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}
