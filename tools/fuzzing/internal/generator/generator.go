package generator

import (
	"fmt"
	"math/rand"

	"github.com/bytebase/parser/tools/fuzzing/internal/config"
	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
	"github.com/pkg/errors"
)

// Generator handles the fuzzing logic
type Generator struct {
	config   *config.Config
	random   *rand.Rand
	grammars []*grammar.ParsedGrammar
}

// New creates a new generator with the given configuration
func New(cfg *config.Config) *Generator {
	return &Generator{
		config: cfg,
		random: rand.New(rand.NewSource(cfg.Seed)),
	}
}

// Generate produces the specified number of queries
func (g *Generator) Generate() error {
	fmt.Println("Initializing grammar parser...")
	
	// Parse all grammar files
	g.grammars = make([]*grammar.ParsedGrammar, len(g.config.GrammarFiles))
	for i, filePath := range g.config.GrammarFiles {
		parsedGrammar, err := grammar.ParseGrammarFile(filePath)
		if err != nil {
			return errors.Wrapf(err, "failed to parse grammar file %s", filePath)
		}
		g.grammars[i] = parsedGrammar
		fmt.Printf("Parsed grammar file: %s\n", filePath)
	}

	// Validate start rule exists
	if !g.hasRule(g.config.StartRule) {
		return errors.Errorf("start rule '%s' not found in any grammar file", g.config.StartRule)
	}

	fmt.Printf("Generating %d queries from rule '%s'...\n", g.config.Count, g.config.StartRule)
	
	// Generate queries
	for i := 0; i < g.config.Count; i++ {
		query := g.generateQuery(i + 1)
		fmt.Printf("Query %d: %s\n", i+1, query)
	}

	return nil
}

// hasRule checks if a rule exists in any of the parsed grammars
func (g *Generator) hasRule(ruleName string) bool {
	for _, grammar := range g.grammars {
		if grammar.GetRule(ruleName) != nil {
			return true
		}
	}
	return false
}

// getRule gets a rule from any of the parsed grammars
func (g *Generator) getRule(ruleName string) *grammar.Rule {
	for _, grammar := range g.grammars {
		if rule := grammar.GetRule(ruleName); rule != nil {
			return rule
		}
	}
	return nil
}

// generateQuery creates a single query using grammar rules
func (g *Generator) generateQuery(index int) string {
	// Start generation from the specified start rule
	result := g.generateFromRule(g.config.StartRule, 0)
	return result
}

// generateFromRule recursively generates text from a grammar rule
func (g *Generator) generateFromRule(ruleName string, currentDepth int) string {
	// Check depth limit to prevent infinite recursion
	if currentDepth >= g.config.MaxDepth {
		return g.generateTerminal(ruleName)
	}

	// Get the rule
	rule := g.getRule(ruleName)
	if rule == nil {
		// If rule not found, return placeholder
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
		elementResult := g.generateFromElement(&element, currentDepth)
		if elementResult != "" {
			result = append(result, elementResult)
		}
	}

	return fmt.Sprintf("/* %s */ %s", ruleName, joinWithSpaces(result))
}

// generateFromElement generates text from a single grammar element
func (g *Generator) generateFromElement(element *grammar.Element, currentDepth int) string {
	// Handle optional elements
	if element.IsOptional() && g.random.Float64() > g.config.OptionalProb {
		return "" // Skip optional element
	}

	// Handle quantified elements
	if element.IsQuantified() {
		return g.generateQuantified(element, currentDepth)
	}

	// Generate single element
	if element.IsRule() {
		if refValue, ok := element.Value.(grammar.ReferenceValue); ok {
			return g.generateFromRule(refValue.Name, currentDepth+1)
		} else if blockValue, ok := element.Value.(grammar.BlockValue); ok {
			return g.generateFromBlock(blockValue, currentDepth)
		}
		return g.generateFromRule(element.Value.String(), currentDepth+1)
	} else if element.IsTerminal() {
		if litValue, ok := element.Value.(grammar.LiteralValue); ok {
			return cleanLiteral(litValue.Text)
		}
		return cleanLiteral(element.Value.String())
	}

	return element.Value.String()
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
				result := g.generateFromRule(refValue.Name, currentDepth+1)
				results = append(results, result)
			} else {
				result := g.generateFromRule(element.Value.String(), currentDepth+1)
				results = append(results, result)
			}
		} else if element.IsTerminal() {
			if litValue, ok := element.Value.(grammar.LiteralValue); ok {
				results = append(results, cleanLiteral(litValue.Text))
			} else {
				results = append(results, cleanLiteral(element.Value.String()))
			}
		} else if blockValue, ok := element.Value.(grammar.BlockValue); ok {
			result := g.generateFromBlock(blockValue, currentDepth+1)
			results = append(results, result)
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

// generateTerminal generates a terminal when depth limit is reached
func (g *Generator) generateTerminal(ruleName string) string {
	// For depth-limited cases, return a simple placeholder
	return fmt.Sprintf("<%s_TERM>", ruleName)
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