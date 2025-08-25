// Code generated from ANTLRv4Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // ANTLRv4Parser
import "github.com/antlr4-go/antlr/v4"

type BaseANTLRv4ParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseANTLRv4ParserVisitor) VisitGrammarSpec(ctx *GrammarSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitGrammarDecl(ctx *GrammarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitGrammarType(ctx *GrammarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitPrequelConstruct(ctx *PrequelConstructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitOptionsSpec(ctx *OptionsSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitOptionValue(ctx *OptionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitDelegateGrammars(ctx *DelegateGrammarsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitDelegateGrammar(ctx *DelegateGrammarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitTokensSpec(ctx *TokensSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitChannelsSpec(ctx *ChannelsSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitIdList(ctx *IdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitAction_(ctx *Action_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitActionScopeName(ctx *ActionScopeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitActionBlock(ctx *ActionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitArgActionBlock(ctx *ArgActionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitModeSpec(ctx *ModeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRules(ctx *RulesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleSpec(ctx *RuleSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitParserRuleSpec(ctx *ParserRuleSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitExceptionGroup(ctx *ExceptionGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitExceptionHandler(ctx *ExceptionHandlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitFinallyClause(ctx *FinallyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRulePrequel(ctx *RulePrequelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleReturns(ctx *RuleReturnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitThrowsSpec(ctx *ThrowsSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLocalsSpec(ctx *LocalsSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleAction(ctx *RuleActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleModifiers(ctx *RuleModifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleModifier(ctx *RuleModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleBlock(ctx *RuleBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleAltList(ctx *RuleAltListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLabeledAlt(ctx *LabeledAltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerRuleSpec(ctx *LexerRuleSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerRuleBlock(ctx *LexerRuleBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerAltList(ctx *LexerAltListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerAlt(ctx *LexerAltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerElements(ctx *LexerElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerElement(ctx *LexerElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerBlock(ctx *LexerBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerCommands(ctx *LexerCommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerCommand(ctx *LexerCommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerCommandName(ctx *LexerCommandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerCommandExpr(ctx *LexerCommandExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitAltList(ctx *AltListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitAlternative(ctx *AlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitPredicateOptions(ctx *PredicateOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitPredicateOption(ctx *PredicateOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLabeledElement(ctx *LabeledElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitEbnf(ctx *EbnfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitBlockSuffix(ctx *BlockSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitEbnfSuffix(ctx *EbnfSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitLexerAtom(ctx *LexerAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitWildcard(ctx *WildcardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitNotSet(ctx *NotSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitBlockSet(ctx *BlockSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitSetElement(ctx *SetElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitRuleref(ctx *RulerefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitCharacterRange(ctx *CharacterRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitTerminalDef(ctx *TerminalDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitElementOptions(ctx *ElementOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitElementOption(ctx *ElementOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseANTLRv4ParserVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
