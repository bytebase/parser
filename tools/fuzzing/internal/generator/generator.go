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
	config           *config.Config
	random           *rand.Rand
	grammar          *grammar.ParsedGrammar
	dependencyGraph  *grammar.DependencyGraph
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
		config:           cfg,
		random:           rand.New(rand.NewSource(cfg.Seed)),
		grammar:          nil,
		dependencyGraph:  nil,
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
	
	// Check if start rule has terminal alternatives
	if !g.dependencyGraph.HasTerminalAlternatives(g.config.StartRule) {
		fmt.Printf("Warning: start rule '%s' has no terminal alternatives\n", g.config.StartRule)
	}

	fmt.Printf("Generating %d queries from rule '%s'...\n", g.config.Count, g.config.StartRule)
	
	// Generate queries
	for i := 0; i < g.config.Count; i++ {
		query := g.generateQuery(i + 1)
		fmt.Printf("Query %d: %s\n", i+1, query)
	}

	return nil
}

// getRule gets a rule from the merged grammar
func (g *Generator) getRule(ruleName string) *grammar.Rule {
	return g.grammar.GetRule(ruleName)
}


// generateQuery creates a single query using grammar rules
func (g *Generator) generateQuery(index int) string {
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
		return fmt.Sprintf("<%s>", ruleName)
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
	// Check if rule has terminal alternatives
	terminalAlts := g.dependencyGraph.GetTerminalAlternatives(ruleName)
	if len(terminalAlts) == 0 {
		// No terminal alternatives - use synthetic generation based on rule name
		return g.generateSyntheticTerminal(ruleName)
	}
	
	rule := g.getRule(ruleName)
	if rule == nil {
		return g.generateSyntheticTerminal(ruleName)
	}
	
	// Try to find the best terminal alternative (prefer ones with more literals)
	bestAltIndex := g.selectBestTerminalAlternative(rule, terminalAlts)
	if bestAltIndex == -1 {
		return g.generateSyntheticTerminal(ruleName)
	}
	
	alternative := rule.Alternatives[bestAltIndex]
	
	// Generate using aggressive terminal mode
	result := g.generateFromAlternativeAggressiveTerminal(&alternative, ruleName)
	
	switch g.config.OutputFormat {
	case config.CompactOutput:
		return result
	case config.VerboseOutput:
		return fmt.Sprintf("/* %s[terminal] */ %s", ruleName, result)
	default:
		return result
	}
}

// generateSyntheticTerminal generates a synthetic terminal based on common SQL patterns
func (g *Generator) generateSyntheticTerminal(ruleName string) string {
	switch ruleName {
	case "selectstmt", "select_no_parens", "select_with_parens":
		return "SELECT 1"
	case "a_expr", "b_expr", "c_expr":
		return "42"
	case "insertStmt", "insertstmt":
		return "INSERT INTO table1 VALUES (1)"
	case "updateStmt", "updatestmt":
		return "UPDATE table1 SET col1 = 1"
	case "deleteStmt", "deletestmt":
		return "DELETE FROM table1"
	case "where_clause", "whereClause":
		return "WHERE 1=1"
	case "having_clause", "havingClause":
		return "HAVING 1=1"
	case "order_by_clause", "orderByClause", "sort_clause":
		return "ORDER BY 1"
	case "group_by_clause", "groupByClause":
		return "GROUP BY 1"
	case "colid", "identifier", "name":
		return "col1"
	case "tablename", "table_name":
		return "table1"
	default:
		// Return a safe default that indicates the rule couldn't be generated
		return fmt.Sprintf("/* %s: synthetic terminal */", ruleName)
	}
}

// generateFromElementWithDepthLimit generates from element with very strict depth limits
func (g *Generator) generateFromElementWithDepthLimit(element *grammar.Element, activeRules map[string]bool, depth int, maxDepth int) string {
	if depth >= maxDepth {
		return g.generateElementFallback(element)
	}
	
	// Handle optional elements - skip them more aggressively when depth limited
	if element.IsOptional() && g.random.Float64() > 0.3 { // Lower probability
		return ""
	}

	// Handle quantified elements - generate very few
	if element.IsQuantified() {
		count := 0
		if element.Quantifier == grammar.ONE_MORE {
			count = 1 // Only generate minimum required
		}
		// For ZERO_MORE, count stays 0
		
		var results []string
		for i := 0; i < count; i++ {
			result := g.generateFromElementWithDepthLimit(&grammar.Element{
				Value: element.Value,
				Quantifier: grammar.NONE,
			}, activeRules, depth+1, maxDepth)
			if result != "" {
				results = append(results, result)
			}
		}
		return joinWithSpaces(results)
	}

	// Generate single element
	if element.IsRule() {
		if refValue, ok := element.Value.(grammar.ReferenceValue); ok {
			if activeRules[refValue.Name] {
				return g.generateSyntheticTerminal(refValue.Name)
			}
			activeRules[refValue.Name] = true
			defer delete(activeRules, refValue.Name)
			
			// Check if this is a lexer rule
			if rule := g.grammar.GetRule(refValue.Name); rule != nil && rule.IsLexer {
				return g.generateConcreteToken(refValue.Name)
			}
			return g.generateSyntheticTerminal(refValue.Name)
		}
		return g.generateElementFallback(element)
	} else if element.IsTerminal() {
		if litValue, ok := element.Value.(grammar.LiteralValue); ok {
			return cleanLiteral(litValue.Text)
		}
		return cleanLiteral(element.Value.String())
	}

	return element.Value.String()
}

// generateElementFallback provides fallback generation for complex elements
func (g *Generator) generateElementFallback(element *grammar.Element) string {
	if element.IsRule() {
		if refValue, ok := element.Value.(grammar.ReferenceValue); ok {
			return g.generateSyntheticTerminal(refValue.Name)
		}
	} else if element.IsTerminal() {
		if litValue, ok := element.Value.(grammar.LiteralValue); ok {
			return cleanLiteral(litValue.Text)
		}
		return cleanLiteral(element.Value.String())
	}
	return "1" // Ultimate fallback
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

// generateFromElement generates text from a single grammar element (legacy method)
func (g *Generator) generateFromElement(element *grammar.Element, currentDepth int) string {
	activeRules := make(map[string]bool)
	return g.generateFromElementWithRecursionTracking(element, activeRules, currentDepth)
}

// generateQuantified handles quantified elements (* +)
func (g *Generator) generateQuantified(element *grammar.Element, currentDepth int) string {
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
				result := g.generateFromRuleOrToken(refValue.Name, currentDepth+1)
				results = append(results, result)
			} else if blockValue, ok := element.Value.(grammar.BlockValue); ok {
				result := g.generateFromBlock(blockValue, currentDepth+1)
				results = append(results, result)
			} else {
				result := g.generateFromRuleOrToken(element.Value.String(), currentDepth+1)
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

// generateFromBlock generates content from a block value
func (g *Generator) generateFromBlock(blockValue grammar.BlockValue, currentDepth int) string {
	if len(blockValue.Alternatives) == 0 {
		return ""
	}

	// Select a random alternative from the block
	altIndex := g.random.Intn(len(blockValue.Alternatives))
	alternative := blockValue.Alternatives[altIndex]

	// Generate from all elements in the selected alternative
	var result []string
	for _, element := range alternative.Elements {
		elementResult := g.generateFromElement(&element, currentDepth)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}

	return joinWithSpaces(result)
}


// generateFromRuleOrToken generates from a rule using standard rule-based generation
func (g *Generator) generateFromRuleOrToken(ruleName string, currentDepth int) string {
	// Check if this is a lexer rule and generate concrete token
	if rule := g.grammar.GetRule(ruleName); rule != nil && rule.IsLexer {
		return g.generateConcreteToken(ruleName)
	}
	
	// Otherwise expand as parser rule
	return g.generateFromRule(ruleName, currentDepth)
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
		return g.generateFromNegatedSet(literal[2 : len(literal)-1])
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
func (g *Generator) generateFromNegatedSet(negatedSet string) string {
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

// selectBestTerminalAlternative randomly selects from terminal alternatives
func (g *Generator) selectBestTerminalAlternative(rule *grammar.Rule, terminalAlts []int) int {
	if len(terminalAlts) == 0 {
		return -1
	}
	
	// Simply choose randomly from available terminal alternatives
	randomIndex := g.random.Intn(len(terminalAlts))
	return terminalAlts[randomIndex]
}

// generateFromAlternativeAggressiveTerminal generates from an alternative using aggressive terminal mode
func (g *Generator) generateFromAlternativeAggressiveTerminal(alt *grammar.Alternative, ruleName string) string {
	var result []string
	
	for _, element := range alt.Elements {
		elementResult := g.generateFromElementAggressiveTerminal(&element, ruleName)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}
	
	if len(result) == 0 {
		// Ultimate fallback - use simple pattern based on rule name
		return g.generateSimpleFallback(ruleName)
	}
	
	return joinWithSpaces(result)
}

// generateFromElementAggressiveTerminal generates from an element using aggressive terminal mode
func (g *Generator) generateFromElementAggressiveTerminal(element *grammar.Element, contextRuleName string) string {
	// Handle quantified elements - be very conservative
	if element.IsQuantified() {
		if element.Quantifier == grammar.ZERO_MORE || element.Quantifier == grammar.OPTIONAL_Q {
			// Skip optional/zero-more elements in terminal mode
			return ""
		} else if element.Quantifier == grammar.ONE_MORE {
			// Generate exactly one for ONE_MORE
			nonQuantifiedElement := grammar.Element{
				Value:      element.Value,
				Quantifier: grammar.NONE,
			}
			return g.generateFromElementAggressiveTerminal(&nonQuantifiedElement, contextRuleName)
		}
	}
	
	// Handle different element types
	if element.IsTerminal() {
		// Direct literal - just return it
		if literal, ok := element.Value.(grammar.LiteralValue); ok {
			return strings.Trim(literal.Text, "'\"")
		}
		return element.Value.String()
	}
	
	if element.IsRule() {
		switch value := element.Value.(type) {
		case grammar.ReferenceValue:
			// Check if it's a simple lexer rule
			if g.isSimpleLexerRule(value.Name) {
				return g.generateConcreteToken(value.Name)
			}
			
			// For parser rules, generate simple fallback based on rule name
			return g.generateSimpleFallback(value.Name)
			
		case grammar.BlockValue:
			// For blocks, try the first alternative only
			if len(value.Alternatives) > 0 {
				return g.generateFromAlternativeAggressiveTerminal(&value.Alternatives[0], contextRuleName)
			}
			return ""
		}
	}
	
	return ""
}

// isSimpleLexerRule checks if a rule is a simple lexer rule that can be safely generated
func (g *Generator) isSimpleLexerRule(ruleName string) bool {
	rule := g.getRule(ruleName)
	if rule == nil || !rule.IsLexer {
		return false
	}
	
	// Consider lexer rules with simple patterns as safe
	simpleLexerRules := map[string]bool{
		"IDENTIFIER": true, "ID": true, "NAME": true,
		"INTEGER": true, "NUMBER": true, "NUMERIC": true, "INT": true,
		"STRING": true, "STRING_LITERAL": true,
		"SELECT": true, "FROM": true, "WHERE": true, "AND": true, "OR": true,
		"COMMA": true, "SEMICOLON": true, "DOT": true,
		"OPEN_PAREN": true, "CLOSE_PAREN": true,
		"PLUS": true, "MINUS": true, "STAR": true, "SLASH": true,
	}
	
	return simpleLexerRules[ruleName]
}

// generateSimpleFallback generates a simple fallback value based on rule name patterns
func (g *Generator) generateSimpleFallback(ruleName string) string {
	// Generate context-appropriate fallbacks
	ruleLower := strings.ToLower(ruleName)
	
	if strings.Contains(ruleLower, "expr") || strings.Contains(ruleLower, "expression") {
		return "1"
	} else if strings.Contains(ruleLower, "name") || strings.Contains(ruleLower, "id") {
		return "col1"
	} else if strings.Contains(ruleLower, "list") {
		return "1"
	} else if strings.Contains(ruleLower, "clause") {
		return "1"
	} else if strings.Contains(ruleLower, "stmt") || strings.Contains(ruleLower, "statement") {
		return "SELECT 1"
	} else if strings.Contains(ruleLower, "select") {
		return "SELECT 1"
	} else {
		// Generic fallback
		return "1"
	}
}