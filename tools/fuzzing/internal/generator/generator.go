package generator

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bytebase/parser/tools/fuzzing/internal/config"
	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
	"github.com/pkg/errors"
)

// Generator handles the fuzzing logic
type Generator struct {
	config          *config.Config
	random          *rand.Rand
	grammar         *grammar.ParsedGrammar
	dependencyGraph *grammar.DependencyGraph
}

// WorkItem represents a unit of work in the generation stack
type WorkItem struct {
	RuleName string
	Depth    int
	Result   *string // Pointer to where the result should be stored
}

// New creates a new generator with the given configuration
func New(cfg *config.Config) *Generator {
	return &Generator{
		config:          cfg,
		random:          rand.New(rand.NewSource(cfg.Seed)),
		grammar:         nil,
		dependencyGraph: nil,
	}
}

// Generate produces the specified number of queries
func (g *Generator) Generate() error {
	fmt.Println("Initializing grammar parser...")

	// Parse and merge all grammar files into a single grammar
	var err error
	g.grammar, err = grammar.ParseAndMergeGrammarFiles(g.config.GrammarFiles)
	if err != nil {
		return errors.Wrap(err, "failed to parse and merge grammar files")
	}

	fmt.Printf("Parsed and merged %d grammar files into single grammar\n", len(g.config.GrammarFiles))

	// Set up dependency graph
	g.dependencyGraph = g.grammar.GetDependencyGraph()

	// Validate grammar has terminal alternatives (non-fatal warning)
	if err := g.grammar.ValidateGrammar(); err != nil {
		fmt.Printf("Grammar validation warning: %v\n", err)
	}

	// Validate start rule exists
	if g.grammar.GetRule(g.config.StartRule) == nil {
		return errors.Errorf("start rule '%s' not found in merged grammar", g.config.StartRule)
	}

	// Check if start rule has immediately terminal alternatives
	if !g.dependencyGraph.HasImmediatelyTerminalAlternatives(g.config.StartRule) {
		fmt.Printf("Warning: start rule '%s' has no immediately terminal alternatives\n", g.config.StartRule)
	}

	fmt.Printf("Generating %d queries from rule '%s'...\n", g.config.Count, g.config.StartRule)

	// Generate queries
	for i := 0; i < g.config.Count; i++ {
		query := g.generateQuery()
		fmt.Printf("Query %d: %s\n", i+1, query)
	}

	return nil
}

// getRule gets a rule from the merged grammar
func (g *Generator) getRule(ruleName string) *grammar.Rule {
	return g.grammar.GetRule(ruleName)
}

// generateQuery creates a single query using grammar rules
func (g *Generator) generateQuery() string {
	// Start generation with fresh active rules tracking
	activeRules := make(map[string]bool)
	result := g.generateFromRuleWithRecursionTracking(g.config.StartRule, activeRules, 0)
	return result
}

// generateFromRuleWithRecursionTracking generates text from a grammar rule with recursion tracking
func (g *Generator) generateFromRuleWithRecursionTracking(ruleName string, activeRules map[string]bool, depth int) string {
	// Check if we're in recursion or hit depth limit
	if activeRules[ruleName] || depth >= g.config.MaxDepth {
		return g.forceTerminalGeneration(ruleName)
	}

	// Mark rule as active
	activeRules[ruleName] = true
	defer delete(activeRules, ruleName)

	// Get the rule
	rule := g.getRule(ruleName)
	if rule == nil {
		return fmt.Sprintf("<%s>", ruleName)
	}

	// Select a random alternative
	if len(rule.Alternatives) == 0 {
		// rule:; is valid but has no alternatives, return empty directly.
		return ""
	}

	altIndex := g.random.Intn(len(rule.Alternatives))
	alternative := rule.Alternatives[altIndex]

	// Generate from all elements in the alternative
	var result []string
	for _, element := range alternative.Elements {
		elementResult := g.generateFromElementWithRecursionTracking(&element, activeRules, depth+1)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}

	// Format output based on configuration
	switch g.config.OutputFormat {
	case config.CompactOutput:
		return joinWithSpaces(result)
	case config.VerboseOutput:
		return fmt.Sprintf("/* %s */ %s", ruleName, joinWithSpaces(result))
	default:
		return joinWithSpaces(result)
	}
}

// forceTerminalGeneration forces generation of terminal alternatives when recursion is detected
func (g *Generator) forceTerminalGeneration(ruleName string) string {
	// Get immediately terminal alternatives first (preferred)
	immediateTerminalAlts := g.dependencyGraph.GetImmediatelyTerminalAlternatives(ruleName)
	rule := g.getRule(ruleName)

	// Check if we have immediately terminal alternatives
	if len(immediateTerminalAlts) == 0 {
		// Fallback for rules that don't have immediately terminal alternatives
		// This can happen with forward references or missing rules
		fmt.Printf("Warning: Rule %s has no immediately terminal alternatives, generating simple fallback\n", ruleName)
		return generateSimpleFallback(ruleName)
	}

	randomIndex := g.random.Intn(len(immediateTerminalAlts))
	selectedAltIndex := immediateTerminalAlts[randomIndex]

	alternative := rule.Alternatives[selectedAltIndex]

	// Generate using normal generation since this is an immediately terminal alternative
	result := g.generateFromImmediatelyTerminalAlternative(&alternative)

	switch g.config.OutputFormat {
	case config.CompactOutput:
		return result
	case config.VerboseOutput:
		return fmt.Sprintf("/* %s[terminal] */ %s", ruleName, result)
	default:
		return result
	}
}

// ForceTerminalGenerationPublic exposes forceTerminalGeneration for testing
func (g *Generator) ForceTerminalGenerationPublic(ruleName string) string {
	return g.forceTerminalGeneration(ruleName)
}

// SetGrammarForTesting sets the grammar for testing purposes
func (g *Generator) SetGrammarForTesting(grammar *grammar.ParsedGrammar) {
	g.grammar = grammar
	g.dependencyGraph = grammar.GetDependencyGraph()
}

// generateFromRule generates text from a grammar rule (legacy method, kept for compatibility)
func (g *Generator) generateFromRule(ruleName string, currentDepth int) string {
	activeRules := make(map[string]bool)
	return g.generateFromRuleWithRecursionTracking(ruleName, activeRules, currentDepth)
}

// generateFromElementWithRecursionTracking generates text from a single grammar element with recursion tracking
func (g *Generator) generateFromElementWithRecursionTracking(element *grammar.Element, activeRules map[string]bool, depth int) string {
	// Handle optional elements
	if element.IsOptional() && g.random.Float64() > g.config.OptionalProb {
		return "" // Skip optional element
	}

	// Handle quantified elements
	if element.IsQuantified() {
		return g.generateQuantifiedWithRecursionTracking(element, activeRules, depth)
	}

	// Generate single element
	if element.IsRule() {
		if refValue, ok := element.Value.(grammar.ReferenceValue); ok {
			return g.generateFromRuleOrTokenWithRecursionTracking(refValue.Name, activeRules, depth)
		} else if blockValue, ok := element.Value.(grammar.BlockValue); ok {
			return g.generateFromBlockWithRecursionTracking(blockValue, activeRules, depth)
		}
		return g.generateFromRuleOrTokenWithRecursionTracking(element.Value.String(), activeRules, depth)
	} else if element.IsTerminal() {
		if litValue, ok := element.Value.(grammar.LiteralValue); ok {
			return cleanLiteral(litValue.Text)
		}
		return cleanLiteral(element.Value.String())
	}

	return element.Value.String()
}

// generateConcreteToken generates concrete tokens by expanding lexer rules
func (g *Generator) generateConcreteToken(ruleName string) string {
	// Get the lexer rule
	rule := g.grammar.GetRule(ruleName)
	if rule == nil || !rule.IsLexer {
		return fmt.Sprintf("<%s>", ruleName)
	}

	// For lexer rules, we need to expand them but generate concrete characters
	// at the terminal level (character sets, literals, etc.)
	return g.generateFromLexerRule(rule, 0)
}

// generateFromLexerRule generates content from a lexer rule
func (g *Generator) generateFromLexerRule(rule *grammar.Rule, currentDepth int) string {
	if len(rule.Alternatives) == 0 {
		return ""
	}

	// Select a random alternative
	altIndex := g.random.Intn(len(rule.Alternatives))
	alternative := rule.Alternatives[altIndex]

	// Generate from all elements in the alternative
	var result []string
	for _, element := range alternative.Elements {
		elementResult := g.generateFromLexerElement(&element, currentDepth)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}

	return strings.Join(result, "")
}

// generateFromLexerElement generates content from a lexer element
func (g *Generator) generateFromLexerElement(element *grammar.Element, currentDepth int) string {
	// Handle optional elements
	if element.IsOptional() && g.random.Float64() > g.config.OptionalProb {
		return "" // Skip optional element
	}

	// Handle quantified elements
	if element.IsQuantified() {
		return g.generateQuantifiedLexer(element, currentDepth)
	}

	// Generate single element
	if element.IsRule() {
		if refValue, ok := element.Value.(grammar.ReferenceValue); ok {
			// Check if referenced rule is lexer or parser
			if referencedRule := g.grammar.GetRule(refValue.Name); referencedRule != nil && referencedRule.IsLexer {
				return g.generateFromLexerRule(referencedRule, currentDepth+1)
			} else {
				// Parser rule - shouldn't happen in lexer context, but handle it
				return g.generateFromRule(refValue.Name, currentDepth+1)
			}
		} else if blockValue, ok := element.Value.(grammar.BlockValue); ok {
			return g.generateFromLexerBlock(blockValue, currentDepth)
		}
		return element.Value.String()
	} else if element.IsTerminal() {
		if litValue, ok := element.Value.(grammar.LiteralValue); ok {
			return g.generateFromLiteral(litValue.Text)
		}
		return g.generateFromLiteral(element.Value.String())
	}

	return element.Value.String()
}

// generateQuantifiedLexer handles quantified lexer elements
func (g *Generator) generateQuantifiedLexer(element *grammar.Element, currentDepth int) string {
	var count int

	// Use fixed count if specified, otherwise use random count
	if g.config.QuantifierCount > 0 {
		count = g.config.QuantifierCount
	} else {
		switch element.Quantifier {
		case grammar.ZERO_MORE: // *
			count = g.random.Intn(g.config.MaxQuantifier + 1) // 0 to MaxQuantifier
		case grammar.ONE_MORE: // +
			count = 1 + g.random.Intn(g.config.MaxQuantifier) // 1 to MaxQuantifier
		default:
			count = 1
		}
	}

	var results []string
	for i := 0; i < count; i++ {
		result := g.generateFromLexerElement(&grammar.Element{
			Value:      element.Value,
			Quantifier: grammar.NONE, // Remove quantifier for individual generation
		}, currentDepth+1)
		if result != "" {
			results = append(results, result)
		}
	}

	return strings.Join(results, "")
}

// generateFromLexerBlock generates content from a lexer block
func (g *Generator) generateFromLexerBlock(blockValue grammar.BlockValue, currentDepth int) string {
	if len(blockValue.Alternatives) == 0 {
		return ""
	}

	// Select a random alternative from the block
	altIndex := g.random.Intn(len(blockValue.Alternatives))
	alternative := blockValue.Alternatives[altIndex]

	// Generate from all elements in the selected alternative
	var result []string
	for _, element := range alternative.Elements {
		elementResult := g.generateFromLexerElement(&element, currentDepth)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}

	return strings.Join(result, "")
}

// generateFromLiteral generates concrete characters from lexer literals and character sets
func (g *Generator) generateFromLiteral(literal string) string {
	// Handle character sets like ~[\u0000"] or [a-zA-Z_]
	if strings.HasPrefix(literal, "~[") && strings.HasSuffix(literal, "]") {
		return g.generateFromNegatedSet()
	} else if strings.HasPrefix(literal, "[") && strings.HasSuffix(literal, "]") {
		return g.generateFromCharacterSet(literal[1 : len(literal)-1])
	}

	// Handle string literals
	if strings.HasPrefix(literal, "'") && strings.HasSuffix(literal, "'") && len(literal) >= 2 {
		return literal[1 : len(literal)-1] // Remove quotes
	}

	// Handle special escape sequences
	switch literal {
	case "\\r":
		return "\r"
	case "\\n":
		return "\n"
	case "\\t":
		return "\t"
	case "\\\"":
		return "\""
	case "\\'":
		return "'"
	case "\\\\":
		return "\\"
	}

	// Return as-is for other cases
	return literal
}

// generateFromCharacterSet generates a random character from a character set like [a-zA-Z_]
func (g *Generator) generateFromCharacterSet(charset string) string {
	chars := []rune{}

	// Simple character set expansion - handle ranges like a-z, A-Z, 0-9
	i := 0
	for i < len(charset) {
		if i+2 < len(charset) && charset[i+1] == '-' {
			// Handle range like a-z
			start := rune(charset[i])
			end := rune(charset[i+2])
			for r := start; r <= end; r++ {
				chars = append(chars, r)
			}
			i += 3
		} else {
			// Single character
			chars = append(chars, rune(charset[i]))
			i++
		}
	}

	if len(chars) == 0 {
		return "x" // Fallback
	}

	return string(chars[g.random.Intn(len(chars))])
}

// generateFromNegatedSet generates a character NOT in the specified set
func (g *Generator) generateFromNegatedSet() string {
	// For simplicity, generate common safe characters that are typically not in negated sets
	safeChars := []string{"a", "b", "c", "x", "y", "z", "_", "1", "2", "3"}

	// TODO: Implement proper negated set handling by expanding the set and excluding those characters
	// For now, just return a safe character
	return safeChars[g.random.Intn(len(safeChars))]
}

// cleanLiteral removes quotes from literal strings
func cleanLiteral(literal string) string {
	// Remove single quotes from literals like 'SELECT'
	if len(literal) >= 2 && literal[0] == '\'' && literal[len(literal)-1] == '\'' {
		return literal[1 : len(literal)-1]
	}
	return literal
}

// joinWithSpaces joins strings with spaces, skipping empty strings
func joinWithSpaces(strs []string) string {
	var nonEmpty []string
	for _, s := range strs {
		if s != "" {
			nonEmpty = append(nonEmpty, s)
		}
	}
	if len(nonEmpty) == 0 {
		return ""
	}
	return joinStrings(nonEmpty, " ")
}

// joinStrings joins strings with a separator
func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}

// generateQuantifiedWithRecursionTracking handles quantified elements with recursion tracking
func (g *Generator) generateQuantifiedWithRecursionTracking(element *grammar.Element, activeRules map[string]bool, depth int) string {
	var count int

	// Use fixed count if specified, otherwise use random count
	if g.config.QuantifierCount > 0 {
		count = g.config.QuantifierCount
	} else {
		switch element.Quantifier {
		case grammar.ZERO_MORE: // *
			count = g.random.Intn(g.config.MaxQuantifier + 1) // 0 to MaxQuantifier
		case grammar.ONE_MORE: // +
			count = 1 + g.random.Intn(g.config.MaxQuantifier) // 1 to MaxQuantifier
		default:
			count = 1
		}
	}

	var results []string
	for i := 0; i < count; i++ {
		if element.IsRule() {
			if refValue, ok := element.Value.(grammar.ReferenceValue); ok {
				result := g.generateFromRuleOrTokenWithRecursionTracking(refValue.Name, activeRules, depth)
				results = append(results, result)
			} else if blockValue, ok := element.Value.(grammar.BlockValue); ok {
				result := g.generateFromBlockWithRecursionTracking(blockValue, activeRules, depth)
				results = append(results, result)
			} else {
				result := g.generateFromRuleOrTokenWithRecursionTracking(element.Value.String(), activeRules, depth)
				results = append(results, result)
			}
		} else if element.IsTerminal() {
			if litValue, ok := element.Value.(grammar.LiteralValue); ok {
				results = append(results, cleanLiteral(litValue.Text))
			} else {
				results = append(results, cleanLiteral(element.Value.String()))
			}
		}
	}

	return joinWithSpaces(results)
}

// generateFromBlockWithRecursionTracking generates content from a block value with recursion tracking
func (g *Generator) generateFromBlockWithRecursionTracking(blockValue grammar.BlockValue, activeRules map[string]bool, depth int) string {
	if len(blockValue.Alternatives) == 0 {
		return ""
	}

	// Select a random alternative from the block
	altIndex := g.random.Intn(len(blockValue.Alternatives))
	alternative := blockValue.Alternatives[altIndex]

	// Generate from all elements in the selected alternative
	var result []string
	for _, element := range alternative.Elements {
		elementResult := g.generateFromElementWithRecursionTracking(&element, activeRules, depth)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}

	return joinWithSpaces(result)
}

// generateFromRuleOrTokenWithRecursionTracking generates from a rule using recursion tracking
func (g *Generator) generateFromRuleOrTokenWithRecursionTracking(ruleName string, activeRules map[string]bool, depth int) string {
	// Check if this is a lexer rule and generate concrete token
	if rule := g.grammar.GetRule(ruleName); rule != nil && rule.IsLexer {
		return g.generateConcreteToken(ruleName)
	}

	// Otherwise expand as parser rule with recursion tracking
	return g.generateFromRuleWithRecursionTracking(ruleName, activeRules, depth)
}

// generateFromImmediatelyTerminalAlternative generates from an immediately terminal alternative using normal generation
// Since the alternative is immediately terminal, we can safely generate without recursion tracking
func (g *Generator) generateFromImmediatelyTerminalAlternative(alt *grammar.Alternative) string {
	var result []string

	for _, element := range alt.Elements {
		elementResult := g.generateFromImmediatelyTerminalElement(&element)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}

	return joinWithSpaces(result)
}

// generateFromImmediatelyTerminalElement generates from an element that's part of an immediately terminal alternative
func (g *Generator) generateFromImmediatelyTerminalElement(element *grammar.Element) string {
	// Handle quantified elements
	if element.IsQuantified() {
		switch element.Quantifier {
		case grammar.ZERO_MORE, grammar.OPTIONAL_Q:
			// Generate 0 or 1 occurrences for optional elements
			if g.random.Float32() < float32(g.config.OptionalProb) {
				nonQuantifiedElement := grammar.Element{
					Value:      element.Value,
					Quantifier: grammar.NONE,
				}
				return g.generateFromImmediatelyTerminalElement(&nonQuantifiedElement)
			}
			return ""
		case grammar.ONE_MORE:
			// Generate exactly 1 occurrence for ONE_MORE to stay minimal
			nonQuantifiedElement := grammar.Element{
				Value:      element.Value,
				Quantifier: grammar.NONE,
			}
			return g.generateFromImmediatelyTerminalElement(&nonQuantifiedElement)
		}
	}

	// Handle different element types
	if element.IsTerminal() {
		// Direct literal - use existing generation logic that handles character sets, literals, etc.
		if literal, ok := element.Value.(grammar.LiteralValue); ok {
			return g.generateFromLiteral(literal.Text)
		}
		return element.Value.String()
	}

	if element.IsRule() {
		switch value := element.Value.(type) {
		case grammar.ReferenceValue:
			// For rule references in immediately terminal alternatives, use normal generation with empty recursion tracking
			// Since we know this is immediately terminal, we can generate safely
			return g.generateFromRuleWithRecursionTracking(value.Name, make(map[string]bool), 0)

		case grammar.BlockValue:
			// For blocks, randomly select an alternative that's immediately terminal
			if len(value.Alternatives) > 0 {
				// Find immediately terminal alternatives within the block
				for _, alt := range value.Alternatives {
					// Use dependency graph to check if this alternative is immediately terminal
					if g.dependencyGraph.CanAlternativeTerminateImmediately(alt) {
						return g.generateFromImmediatelyTerminalAlternative(&alt)
					}
				}
				// Fallback to first alternative if none found (shouldn't happen)
				return g.generateFromImmediatelyTerminalAlternative(&value.Alternatives[0])
			}
			return ""
		}
	}

	return ""
}

// generateSimpleFallback generates a simple fallback value based on rule name patterns
func generateSimpleFallback(ruleName string) string {
	// Generate context-appropriate fallbacks
	ruleLower := strings.ToLower(ruleName)

	if strings.Contains(ruleLower, "string") || strings.Contains(ruleLower, "constant") {
		return "'fallback'"
	} else if strings.Contains(ruleLower, "expr") || strings.Contains(ruleLower, "expression") {
		return "1"
	} else if strings.Contains(ruleLower, "name") || strings.Contains(ruleLower, "id") {
		return "col1"
	} else if strings.Contains(ruleLower, "number") || strings.Contains(ruleLower, "numeric") {
		return "1"
	} else {
		// Generic fallback
		return "1"
	}
}

