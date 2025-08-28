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
	// Focus only on parser rules for now
	if parserRuleCtx := ctx.ParserRuleSpec(); parserRuleCtx != nil {
		v.VisitParserRuleSpec(parserRuleCtx)
	}
	// Skip lexer rules for now
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