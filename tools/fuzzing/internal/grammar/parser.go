package grammar

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"
	grammar "github.com/bytebase/parser/tools/grammar"
)

// ParsedGrammar represents a parsed grammar with extracted rules
type ParsedGrammar struct {
	LexerRules  map[string]*Rule
	ParserRules map[string]*Rule
	FilePath    string
	// BlockAltMap stores temporary block rules for debugging
	// Key: block ID (e.g., "block_1_alts"), Value: the block alternatives
	BlockAltMap map[string][]Alternative
}

// Rule represents a grammar rule with its alternatives
type Rule struct {
	Name         string
	Alternatives []Alternative
	IsLexer      bool
}

// Alternative represents one alternative of a rule
type Alternative struct {
	Elements []Element
}

// Global block ID counter for generating unique block names
var globalBlockID = 0

// ElementValue represents different types of element values
type ElementValue interface {
	// String returns a string representation for display/debugging
	String() string
}

// LiteralValue represents a literal string (e.g., 'SELECT')
type LiteralValue struct {
	Text string
}

func (l LiteralValue) String() string { return l.Text }

// ReferenceValue represents a reference to a rule or token (e.g., IDENTIFIER, selectStmt)
type ReferenceValue struct {
	Name string
}

func (r ReferenceValue) String() string { return r.Name }

// BlockValue represents a generated block (e.g., (',' column)*)
type BlockValue struct {
	ID           string        // Global unique ID like "block_1_alts"
	Alternatives []Alternative
}

func (b BlockValue) String() string {
	if len(b.Alternatives) == 0 {
		return "<empty_block>"
	}
	if len(b.Alternatives) == 1 {
		elements := []string{}
		for _, elem := range b.Alternatives[0].Elements {
			elements = append(elements, elem.Value.String())
		}
		return fmt.Sprintf("(%s)", strings.Join(elements, " "))
	}
	return b.ID
}


// WildcardValue represents a wildcard (.)
type WildcardValue struct{}

func (w WildcardValue) String() string { return "." }

// Element represents an element within an alternative
type Element struct {
	Value      ElementValue
	Quantifier Quantifier
}

// Quantifier indicates repetition type
type Quantifier int

const (
	NONE Quantifier = iota
	OPTIONAL_Q // ?
	ZERO_MORE  // *
	ONE_MORE   // +
)

// ParseGrammarFile parses a .g4 file and extracts rules for fuzzing
func ParseGrammarFile(filePath string) (*ParsedGrammar, error) {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read grammar file")
	}

	if len(content) == 0 {
		return nil, errors.New("grammar file is empty")
	}

	// Create input stream
	input := antlr.NewInputStream(string(content))

	// Create lexer
	lexer := grammar.NewANTLRv4Lexer(input)

	// Add error listener
	errorListener := &GrammarErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	parser := grammar.NewANTLRv4Parser(stream)

	// Add error listener to parser
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	// Parse the grammar
	tree := parser.GrammarSpec()

	// Check for parsing errors
	if errorListener.HasErrors() {
		return nil, errors.Errorf("failed to parse grammar: %v", errorListener.GetErrors())
	}

	if tree == nil {
		return nil, errors.New("parser returned nil tree")
	}

	// Extract rules from parse tree
	visitor := NewGrammarExtractorVisitor()
	visitor.VisitGrammarSpec(tree)



	return &ParsedGrammar{
		LexerRules:  visitor.lexerRules,
		ParserRules: visitor.parserRules,
		FilePath:    filePath,
		BlockAltMap: visitor.blockAltMap,
	}, nil
}

// GetRule gets a rule by name from either lexer or parser rules
func (g *ParsedGrammar) GetRule(name string) *Rule {
	if rule, ok := g.ParserRules[name]; ok {
		return rule
	}
	if rule, ok := g.LexerRules[name]; ok {
		return rule
	}
	return nil
}

// GetAllRules returns all rules (both lexer and parser)
func (g *ParsedGrammar) GetAllRules() map[string]*Rule {
	allRules := make(map[string]*Rule)
	for name, rule := range g.LexerRules {
		allRules[name] = rule
	}
	for name, rule := range g.ParserRules {
		allRules[name] = rule
	}
	return allRules
}

// GetBlockAlternatives returns the alternatives for a generated block ID
func (g *ParsedGrammar) GetBlockAlternatives(blockID string) ([]Alternative, bool) {
	alts, exists := g.BlockAltMap[blockID]
	return alts, exists
}

// IsGeneratedBlock checks if a name refers to a generated block
func (g *ParsedGrammar) IsGeneratedBlock(name string) bool {
	_, exists := g.BlockAltMap[name]
	return exists
}

// MergeGrammar merges another grammar into this one
func (g *ParsedGrammar) MergeGrammar(other *ParsedGrammar) error {
	// Merge lexer rules
	for name, rule := range other.LexerRules {
		if _, exists := g.LexerRules[name]; exists {
			return fmt.Errorf("duplicate lexer rule '%s' found in grammars '%s' and '%s'", name, g.FilePath, other.FilePath)
		}
		g.LexerRules[name] = rule
	}
	
	// Merge parser rules
	for name, rule := range other.ParserRules {
		if _, exists := g.ParserRules[name]; exists {
			return fmt.Errorf("duplicate parser rule '%s' found in grammars '%s' and '%s'", name, g.FilePath, other.FilePath)
		}
		g.ParserRules[name] = rule
	}
	
	// Merge block alternatives map
	for blockID, alternatives := range other.BlockAltMap {
		if _, exists := g.BlockAltMap[blockID]; exists {
			return fmt.Errorf("duplicate block ID '%s' found in grammars '%s' and '%s'", blockID, g.FilePath, other.FilePath)
		}
		g.BlockAltMap[blockID] = alternatives
	}
	
	// Update file path to indicate it's a merged grammar
	if g.FilePath != other.FilePath {
		g.FilePath = fmt.Sprintf("%s + %s", g.FilePath, other.FilePath)
	}
	
	return nil
}

// ParseAndMergeGrammarFiles parses multiple grammar files and merges them into a single ParsedGrammar
func ParseAndMergeGrammarFiles(filePaths []string) (*ParsedGrammar, error) {
	if len(filePaths) == 0 {
		return nil, errors.New("no grammar files provided")
	}
	
	// Parse the first grammar file
	mergedGrammar, err := ParseGrammarFile(filePaths[0])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse first grammar file %s", filePaths[0])
	}
	
	// Merge additional grammar files
	for i := 1; i < len(filePaths); i++ {
		filePath := filePaths[i]
		grammar, err := ParseGrammarFile(filePath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse grammar file %s", filePath)
		}
		
		if err := mergedGrammar.MergeGrammar(grammar); err != nil {
			return nil, errors.Wrapf(err, "failed to merge grammar file %s", filePath)
		}
	}
	
	return mergedGrammar, nil
}

// IsRule checks if an element refers to another rule or generated block
func (e *Element) IsRule() bool {
	_, isRef := e.Value.(ReferenceValue)
	_, isBlock := e.Value.(BlockValue)
	return isRef || isBlock
}

// IsTerminal checks if an element is a terminal (literal)
func (e *Element) IsTerminal() bool {
	_, isLit := e.Value.(LiteralValue)
	_, isWild := e.Value.(WildcardValue)
	return isLit || isWild
}

// IsOptional checks if an element has optional quantifier
func (e *Element) IsOptional() bool {
	return e.Quantifier == OPTIONAL_Q
}

// IsQuantified checks if an element has repetition quantifiers
func (e *Element) IsQuantified() bool {
	return e.Quantifier == ZERO_MORE || e.Quantifier == ONE_MORE
}

// GrammarErrorListener collects parsing errors
type GrammarErrorListener struct {
	errors []string
}

func (l *GrammarErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func (l *GrammarErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Ignore ambiguity for fuzzing purposes
}

func (l *GrammarErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Ignore for fuzzing purposes
}

func (l *GrammarErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	// Ignore for fuzzing purposes
}

func (l *GrammarErrorListener) HasErrors() bool {
	return len(l.errors) > 0
}

func (l *GrammarErrorListener) GetErrors() []string {
	return l.errors
}

// GrammarExtractorVisitor extracts rules from the parse tree
type GrammarExtractorVisitor struct {
	*grammar.BaseANTLRv4ParserVisitor
	lexerRules  map[string]*Rule
	parserRules map[string]*Rule
	blockAltMap map[string][]Alternative
}

// NewGrammarExtractorVisitor creates a new visitor
func NewGrammarExtractorVisitor() *GrammarExtractorVisitor {
	v := &GrammarExtractorVisitor{
		BaseANTLRv4ParserVisitor: &grammar.BaseANTLRv4ParserVisitor{},
		lexerRules:               make(map[string]*Rule),
		parserRules:              make(map[string]*Rule),
		blockAltMap:              make(map[string][]Alternative),
	}
	return v
}

// VisitGrammarSpec visits the grammar specification
func (v *GrammarExtractorVisitor) VisitGrammarSpec(ctx grammar.IGrammarSpecContext) interface{} {
	// Visit rules section
	if rulesCtx := ctx.Rules(); rulesCtx != nil {
		v.VisitRules(rulesCtx)
	}
	return nil
}

// VisitRules visits the rules section
func (v *GrammarExtractorVisitor) VisitRules(ctx grammar.IRulesContext) interface{} {
	// Visit all rule specifications
	for _, ruleSpecCtx := range ctx.AllRuleSpec() {
		v.VisitRuleSpec(ruleSpecCtx)
	}
	return nil
}

// VisitRuleSpec visits a rule specification (could be parser or lexer rule)
func (v *GrammarExtractorVisitor) VisitRuleSpec(ctx grammar.IRuleSpecContext) interface{} {
	// Handle parser rules
	if parserRuleCtx := ctx.ParserRuleSpec(); parserRuleCtx != nil {
		v.VisitParserRuleSpec(parserRuleCtx)
	}
	// Handle lexer rules
	if lexerRuleCtx := ctx.LexerRuleSpec(); lexerRuleCtx != nil {
		v.VisitLexerRuleSpec(lexerRuleCtx)
	}
	return nil
}

// VisitParserRuleSpec visits a parser rule specification
func (v *GrammarExtractorVisitor) VisitParserRuleSpec(ctx grammar.IParserRuleSpecContext) interface{} {
	// Get rule name
	ruleNameToken := ctx.RULE_REF()
	if ruleNameToken == nil {
		return nil
	}
	ruleName := ruleNameToken.GetText()

	// Get rule block (alternatives)
	ruleBlockCtx := ctx.RuleBlock()
	if ruleBlockCtx == nil {
		return nil
	}

	// Extract alternatives
	alternatives := v.extractAlternatives(ruleBlockCtx)

	// Create rule
	rule := &Rule{
		Name:         ruleName,
		IsLexer:      false,
		Alternatives: alternatives,
	}

	// Store rule
	v.parserRules[ruleName] = rule

	return nil
}

// VisitLexerRuleSpec visits a lexer rule specification
func (v *GrammarExtractorVisitor) VisitLexerRuleSpec(ctx grammar.ILexerRuleSpecContext) interface{} {
	// Get rule name
	ruleNameToken := ctx.TOKEN_REF()
	if ruleNameToken == nil {
		return nil
	}
	ruleName := ruleNameToken.GetText()

	// Get lexer rule block (alternatives)
	lexerRuleBlockCtx := ctx.LexerRuleBlock()
	if lexerRuleBlockCtx == nil {
		return nil
	}

	// Extract alternatives from lexer rule block
	alternatives := v.extractLexerAlternatives(lexerRuleBlockCtx)

	// Create rule
	rule := &Rule{
		Name:         ruleName,
		IsLexer:      true,
		Alternatives: alternatives,
	}

	// Store rule
	v.lexerRules[ruleName] = rule

	return nil
}

// extractAlternatives extracts alternatives from a rule block
func (v *GrammarExtractorVisitor) extractAlternatives(ruleBlockCtx grammar.IRuleBlockContext) []Alternative {
	var alternatives []Alternative

	// Get rule alternative list
	ruleAltListCtx := ruleBlockCtx.RuleAltList()
	if ruleAltListCtx == nil {
		return alternatives
	}

	// Process each labeled alternative
	for _, labeledAltCtx := range ruleAltListCtx.AllLabeledAlt() {
		alternative := v.extractAlternative(labeledAltCtx)
		alternatives = append(alternatives, alternative)
	}

	return alternatives
}

// extractLexerAlternatives extracts alternatives from a lexer rule block
func (v *GrammarExtractorVisitor) extractLexerAlternatives(lexerRuleBlockCtx grammar.ILexerRuleBlockContext) []Alternative {
	var alternatives []Alternative

	// Get lexer alternative list
	lexerAltListCtx := lexerRuleBlockCtx.LexerAltList()
	if lexerAltListCtx == nil {
		return alternatives
	}

	// Process each lexer alternative
	for _, lexerAltCtx := range lexerAltListCtx.AllLexerAlt() {
		alternative := v.extractLexerAlternative(lexerAltCtx)
		alternatives = append(alternatives, alternative)
	}

	return alternatives
}

// extractLexerAlternative extracts a single lexer alternative
func (v *GrammarExtractorVisitor) extractLexerAlternative(lexerAltCtx grammar.ILexerAltContext) Alternative {
	var elements []Element

	// Get lexer elements context
	lexerElementsCtx := lexerAltCtx.LexerElements()
	if lexerElementsCtx != nil {
		// Process each lexer element
		for _, lexerElementCtx := range lexerElementsCtx.AllLexerElement() {
			element := v.extractLexerElement(lexerElementCtx)
			if element != nil {
				elements = append(elements, *element)
			}
		}
	}

	return Alternative{
		Elements: elements,
	}
}

// extractAlternative extracts a single alternative
func (v *GrammarExtractorVisitor) extractAlternative(labeledAltCtx grammar.ILabeledAltContext) Alternative {
	var elements []Element

	// Get alternative context
	altCtx := labeledAltCtx.Alternative()
	if altCtx != nil {
		// Process each element in the alternative
		for _, elementCtx := range altCtx.AllElement() {
			element := v.extractElement(elementCtx)
			if element != nil {
				elements = append(elements, *element)
			}
		}
	}

	return Alternative{
		Elements: elements,
	}
}

// extractElement extracts an element from an element context
func (v *GrammarExtractorVisitor) extractElement(elementCtx grammar.IElementContext) *Element {
	// Handle labeled elements
	if labeledElementCtx := elementCtx.LabeledElement(); labeledElementCtx != nil {
		return v.extractLabeledElement(labeledElementCtx)
	}

	// Handle atoms (terminals/non-terminals)
	if atomCtx := elementCtx.Atom(); atomCtx != nil {
		element := v.extractAtom(atomCtx)
		// Check for quantifiers
		if element != nil {
			element.Quantifier = v.extractQuantifier(elementCtx.EbnfSuffix())
		}
		return element
	}

	// Handle EBNF constructs (blocks with quantifiers)
	if ebnfCtx := elementCtx.Ebnf(); ebnfCtx != nil {
		return v.extractEbnf(ebnfCtx)
	}

	return nil
}

// extractLexerElement extracts a lexer element from a lexer element context
func (v *GrammarExtractorVisitor) extractLexerElement(lexerElementCtx grammar.ILexerElementContext) *Element {
	// Handle lexer atoms (character ranges, terminals, etc.)
	if lexerAtomCtx := lexerElementCtx.LexerAtom(); lexerAtomCtx != nil {
		element := v.extractLexerAtom(lexerAtomCtx)
		// Check for quantifiers
		if element != nil {
			element.Quantifier = v.extractQuantifier(lexerElementCtx.EbnfSuffix())
		}
		return element
	}

	// Handle lexer blocks (grouped alternatives)
	if lexerBlockCtx := lexerElementCtx.LexerBlock(); lexerBlockCtx != nil {
		element := v.extractLexerBlock(lexerBlockCtx)
		// Check for quantifiers
		if element != nil {
			element.Quantifier = v.extractQuantifier(lexerElementCtx.EbnfSuffix())
		}
		return element
	}

	// Handle action blocks (for now, just return nil as they don't generate text)
	if lexerElementCtx.ActionBlock() != nil {
		// Action blocks don't contribute to generated text, so we skip them
		return nil
	}

	return nil
}

// extractLexerAtom extracts a lexer atom (character ranges, terminals, etc.)
func (v *GrammarExtractorVisitor) extractLexerAtom(lexerAtomCtx grammar.ILexerAtomContext) *Element {
	// Handle terminal definition (string literal or token reference)
	if terminalDefCtx := lexerAtomCtx.TerminalDef(); terminalDefCtx != nil {
		return v.extractTerminalDef(terminalDefCtx)
	}

	// Handle character range (e.g., [a-z])
	if characterRangeCtx := lexerAtomCtx.CharacterRange(); characterRangeCtx != nil {
		return v.extractCharacterRange(characterRangeCtx)
	}

	// Handle not set (e.g., ~[abc])
	if notSetCtx := lexerAtomCtx.NotSet(); notSetCtx != nil {
		return v.extractNotSet(notSetCtx)
	}

	// Handle lexer character set (e.g., [abc])
	if lexerCharSetToken := lexerAtomCtx.LEXER_CHAR_SET(); lexerCharSetToken != nil {
		return &Element{
			Value: LiteralValue{Text: lexerCharSetToken.GetText()},
		}
	}

	// Handle wildcard (.)
	if wildcardCtx := lexerAtomCtx.Wildcard(); wildcardCtx != nil {
		return &Element{
			Value: WildcardValue{},
		}
	}

	return nil
}

// extractLexerBlock extracts a lexer block (grouped alternatives)
func (v *GrammarExtractorVisitor) extractLexerBlock(lexerBlockCtx grammar.ILexerBlockContext) *Element {
	// Get the lexer alternative list from the block
	lexerAltListCtx := lexerBlockCtx.LexerAltList()
	if lexerAltListCtx == nil {
		globalBlockID++
		blockID := fmt.Sprintf("lexer_block_%d_alts", globalBlockID)
		emptyAlts := []Alternative{}
		v.blockAltMap[blockID] = emptyAlts
		
		return &Element{
			Value: BlockValue{ID: blockID, Alternatives: emptyAlts},
		}
	}

	// Extract all lexer alternatives from the block
	lexerAlts := lexerAltListCtx.AllLexerAlt()
	if len(lexerAlts) == 0 {
		globalBlockID++
		blockID := fmt.Sprintf("lexer_block_%d_alts", globalBlockID)
		emptyAlts := []Alternative{}
		v.blockAltMap[blockID] = emptyAlts
		
		return &Element{
			Value: BlockValue{ID: blockID, Alternatives: emptyAlts},
		}
	}

	// Extract all alternatives
	blockAlternatives := []Alternative{}
	for _, lexerAltCtx := range lexerAlts {
		elements := []Element{}
		if lexerElementsCtx := lexerAltCtx.LexerElements(); lexerElementsCtx != nil {
			for _, lexerElementCtx := range lexerElementsCtx.AllLexerElement() {
				element := v.extractLexerElement(lexerElementCtx)
				if element != nil {
					elements = append(elements, *element)
				}
			}
		}
		blockAlternatives = append(blockAlternatives, Alternative{Elements: elements})
	}
	
	// Generate global unique block ID and store mapping
	globalBlockID++
	blockID := fmt.Sprintf("lexer_block_%d_alts", globalBlockID)
	v.blockAltMap[blockID] = blockAlternatives
	
	return &Element{
		Value: BlockValue{ID: blockID, Alternatives: blockAlternatives},
	}
}

// extractCharacterRange extracts a character range (e.g., 'a'..'z')
func (v *GrammarExtractorVisitor) extractCharacterRange(characterRangeCtx grammar.ICharacterRangeContext) *Element {
	// Get the start and end of the range
	stringLiterals := characterRangeCtx.AllSTRING_LITERAL()
	if len(stringLiterals) == 2 {
		startChar := stringLiterals[0].GetText()
		endChar := stringLiterals[1].GetText()
		rangeText := fmt.Sprintf("%s..%s", startChar, endChar)
		return &Element{
			Value: LiteralValue{Text: rangeText},
		}
	}
	return nil
}

// extractNotSet extracts a not set (e.g., ~[abc])
func (v *GrammarExtractorVisitor) extractNotSet(notSetCtx grammar.INotSetContext) *Element {
	// For now, represent as a literal text
	// In a real implementation, this would need more sophisticated handling
	return &Element{
		Value: LiteralValue{Text: "~[...]"},
	}
}

// extractLabeledElement extracts a labeled element (e.g., label=atom)
func (v *GrammarExtractorVisitor) extractLabeledElement(labeledElementCtx grammar.ILabeledElementContext) *Element {
	// For now, just extract the atom part and ignore the label
	if atomCtx := labeledElementCtx.Atom(); atomCtx != nil {
		return v.extractAtom(atomCtx)
	}
	if blockCtx := labeledElementCtx.Block(); blockCtx != nil {
		return v.extractBlock(blockCtx)
	}
	return nil
}

// extractAtom extracts an atom (terminal or non-terminal)
func (v *GrammarExtractorVisitor) extractAtom(atomCtx grammar.IAtomContext) *Element {
	// Handle terminal definition (string literal or token reference)
	if terminalDefCtx := atomCtx.TerminalDef(); terminalDefCtx != nil {
		return v.extractTerminalDef(terminalDefCtx)
	}

	// Handle rule reference
	if rulerefCtx := atomCtx.Ruleref(); rulerefCtx != nil {
		return v.extractRuleRef(rulerefCtx)
	}

	// Handle wildcard (.)
	if wildcardCtx := atomCtx.Wildcard(); wildcardCtx != nil {
		return &Element{
			Value: WildcardValue{},
		}
	}

	// Handle not sets, ranges, etc. - for now just return nil
	return nil
}

// extractTerminalDef extracts a terminal definition (literal string or token reference)
func (v *GrammarExtractorVisitor) extractTerminalDef(terminalDefCtx grammar.ITerminalDefContext) *Element {
	if stringLiteralToken := terminalDefCtx.STRING_LITERAL(); stringLiteralToken != nil {
		return &Element{
			Value: LiteralValue{Text: stringLiteralToken.GetText()},
		}
	}
	if tokenRefToken := terminalDefCtx.TOKEN_REF(); tokenRefToken != nil {
		return &Element{
			Value: ReferenceValue{Name: tokenRefToken.GetText()},
		}
	}
	return nil
}


// extractRuleRef extracts a rule reference
func (v *GrammarExtractorVisitor) extractRuleRef(rulerefCtx grammar.IRulerefContext) *Element {
	if ruleRefToken := rulerefCtx.RULE_REF(); ruleRefToken != nil {
		return &Element{
			Value: ReferenceValue{Name: ruleRefToken.GetText()},
		}
	}
	return nil
}

// extractBlock extracts a block (grouped alternatives)
func (v *GrammarExtractorVisitor) extractBlock(blockCtx grammar.IBlockContext) *Element {
	// Get the alternative list from the block
	altListCtx := blockCtx.AltList()
	if altListCtx == nil {
		globalBlockID++
		blockID := fmt.Sprintf("block_%d_alts", globalBlockID)
		emptyAlts := []Alternative{}
		v.blockAltMap[blockID] = emptyAlts
		
		return &Element{
			Value: BlockValue{ID: blockID, Alternatives: emptyAlts},
		}
	}

	// Extract all alternatives from the block
	alts := altListCtx.AllAlternative()
	if len(alts) == 0 {
		globalBlockID++
		blockID := fmt.Sprintf("block_%d_alts", globalBlockID)
		emptyAlts := []Alternative{}
		v.blockAltMap[blockID] = emptyAlts
		
		return &Element{
			Value: BlockValue{ID: blockID, Alternatives: emptyAlts},
		}
	}

	// Extract all alternatives
	blockAlternatives := []Alternative{}
	for _, altCtx := range alts {
		elements := []Element{}
		for _, elementCtx := range altCtx.AllElement() {
			element := v.extractElement(elementCtx)
			if element != nil {
				elements = append(elements, *element)
			}
		}
		blockAlternatives = append(blockAlternatives, Alternative{Elements: elements})
	}

	// If it's a single element in a single alternative, we can simplify
	if len(blockAlternatives) == 1 && len(blockAlternatives[0].Elements) == 1 {
		return &blockAlternatives[0].Elements[0]
	}
	
	// Generate global unique block ID and store mapping
	globalBlockID++
	blockID := fmt.Sprintf("block_%d_alts", globalBlockID)
	v.blockAltMap[blockID] = blockAlternatives
	
	return &Element{
		Value: BlockValue{ID: blockID, Alternatives: blockAlternatives},
	}
}

// extractEbnf extracts EBNF constructs (blocks with suffixes)
func (v *GrammarExtractorVisitor) extractEbnf(ebnfCtx grammar.IEbnfContext) *Element {
	// Get the block
	blockCtx := ebnfCtx.Block()
	if blockCtx == nil {
		return nil
	}

	element := v.extractBlock(blockCtx)
	if element != nil {
		// Apply quantifier from block suffix
		if blockSuffixCtx := ebnfCtx.BlockSuffix(); blockSuffixCtx != nil {
			if ebnfSuffixCtx := blockSuffixCtx.EbnfSuffix(); ebnfSuffixCtx != nil {
				element.Quantifier = v.extractQuantifier(ebnfSuffixCtx)
			}
		}
	}

	return element
}

// extractQuantifier extracts quantifier from EBNF suffix
func (v *GrammarExtractorVisitor) extractQuantifier(ebnfSuffixCtx grammar.IEbnfSuffixContext) Quantifier {
	if ebnfSuffixCtx == nil {
		return NONE
	}

	// Check for question mark (optional)
	if ebnfSuffixCtx.QUESTION(0) != nil {
		return OPTIONAL_Q
	}

	// Check for star (zero or more)
	if ebnfSuffixCtx.STAR() != nil {
		return ZERO_MORE
	}

	// Check for plus (one or more)
	if ebnfSuffixCtx.PLUS() != nil {
		return ONE_MORE
	}

	return NONE
}