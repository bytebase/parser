package grammar

import (
	"fmt"
)

// DependencyGraph represents the dependency relationships between grammar rules
type DependencyGraph struct {
	Nodes map[string]*GraphNode
}

// GraphNode represents a single rule in the dependency graph
type GraphNode struct {
	RuleName                            string        // Rule name (e.g., "selectStmt", "expr")
	Alternatives                        []Alternative // All alternatives for this rule
	HasImmediatelyTerminalAlternatives  bool          // Has at least one immediately terminal alternative
	ImmediatelyTerminalAlternativeIndex []int         // Indices of alternatives that are immediately terminal
	IsLexer                             bool          // Whether this is a lexer rule
}

// NewDependencyGraph creates a new dependency graph
func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		Nodes: make(map[string]*GraphNode),
	}
}

// AddNode adds a rule node to the dependency graph
func (g *DependencyGraph) AddNode(ruleName string, rule *Rule) {
	node := &GraphNode{
		RuleName:                            ruleName,
		Alternatives:                        rule.Alternatives,
		HasImmediatelyTerminalAlternatives:  false,
		ImmediatelyTerminalAlternativeIndex: []int{},
		IsLexer:                             rule.IsLexer,
	}
	g.Nodes[ruleName] = node
}

// GetNode retrieves a node by rule name
func (g *DependencyGraph) GetNode(ruleName string) *GraphNode {
	return g.Nodes[ruleName]
}

// AnalyzeTerminalReachability performs immediately terminal analysis on the graph
func (g *DependencyGraph) AnalyzeTerminalReachability() error {
	return g.AnalyzeTerminalReachabilityWithValidation(false)
}

// AnalyzeTerminalReachabilityWithValidation performs immediately terminal analysis with optional validation
func (g *DependencyGraph) AnalyzeTerminalReachabilityWithValidation(validateUnterminated bool) error {
	// Phase 1: Mark lexer rules as immediately terminal
	g.markLexerRulesAsImmediatelyTerminal()

	// Phase 2: Analyze immediately terminal alternatives
	g.analyzeImmediatelyTerminalAlternatives()

	// Phase 3: Check for nodes without immediately terminal alternatives and report error (only if requested)
	if validateUnterminated {
		return g.validateImmediatelyTerminalReachability()
	}

	return nil
}

// markLexerRulesAsImmediatelyTerminal marks all lexer rules as having immediately terminal alternatives
func (g *DependencyGraph) markLexerRulesAsImmediatelyTerminal() {
	for _, node := range g.Nodes {
		if node.IsLexer {
			node.HasImmediatelyTerminalAlternatives = true
			// All lexer alternatives are considered immediately terminal
			for i := range node.Alternatives {
				node.ImmediatelyTerminalAlternativeIndex = append(node.ImmediatelyTerminalAlternativeIndex, i)
			}
		}
	}
}

// analyzeImmediatelyTerminalAlternatives analyzes which alternatives are immediately terminal
func (g *DependencyGraph) analyzeImmediatelyTerminalAlternatives() {
	// Use fixed-point iteration similar to terminal propagation
	changed := true
	iterations := 0
	maxIterations := len(g.Nodes) * 2 // Prevent infinite loops

	for changed && iterations < maxIterations {
		changed = false
		iterations++

		for _, node := range g.Nodes {
			if node.IsLexer {
				continue // Already processed
			}

			// Check each alternative to see if it's immediately terminal
			for altIndex, alt := range node.Alternatives {
				// Skip if this alternative is already marked as immediately terminal
				if g.isAlternativeAlreadyMarkedImmediate(node, altIndex) {
					continue
				}

				if g.canAlternativeTerminateImmediately(alt) {
					if !node.HasImmediatelyTerminalAlternatives {
						node.HasImmediatelyTerminalAlternatives = true
						changed = true
					}
					node.ImmediatelyTerminalAlternativeIndex = append(node.ImmediatelyTerminalAlternativeIndex, altIndex)
					changed = true
				}
			}
		}
	}

	if iterations >= maxIterations {
		fmt.Printf("Warning: Immediately terminal analysis reached max iterations (%d)\\n", maxIterations)
	}
}

// validateImmediatelyTerminalReachability checks for rules without immediately terminal alternatives and reports errors
func (g *DependencyGraph) validateImmediatelyTerminalReachability() error {
	var unterminatedRules []string

	for ruleName, node := range g.Nodes {
		if !node.HasImmediatelyTerminalAlternatives {
			unterminatedRules = append(unterminatedRules, ruleName)
		}
	}

	if len(unterminatedRules) > 0 {
		return fmt.Errorf("grammar contains %d rules without immediately terminal alternatives: %v",
			len(unterminatedRules), unterminatedRules)
	}

	return nil
}

// isAlternativeAlreadyMarkedImmediate checks if an alternative is already in the immediately terminal list
func (g *DependencyGraph) isAlternativeAlreadyMarkedImmediate(node *GraphNode, altIndex int) bool {
	for _, immediateIndex := range node.ImmediatelyTerminalAlternativeIndex {
		if immediateIndex == altIndex {
			return true
		}
	}
	return false
}

// canAlternativeTerminateImmediately checks if an alternative can terminate immediately (no rule references required)
func (g *DependencyGraph) canAlternativeTerminateImmediately(alt Alternative) bool {
	// Empty alternative (ε-transition) can always terminate immediately
	if len(alt.Elements) == 0 {
		return true
	}

	// Check each element in the alternative
	for _, element := range alt.Elements {
		if !g.canElementTerminateImmediately(element) {
			return false
		}
	}

	return true
}

// canElementTerminateImmediately checks if a single element can terminate immediately
func (g *DependencyGraph) canElementTerminateImmediately(element Element) bool {
	// Terminal elements (literals) can always terminate immediately
	if element.IsTerminal() {
		return true
	}

	// Handle quantified elements
	if element.IsQuantified() {
		// * and ? quantifiers can generate 0 occurrences, so they can terminate immediately
		if element.Quantifier == ZERO_MORE || element.Quantifier == OPTIONAL_Q {
			return true
		}
		// + quantifier requires at least one occurrence, so check the referenced rule
		if element.Quantifier == ONE_MORE {
			return g.canRuleReferenceTerminateImmediately(element)
		}
	}

	// For rule references, check if the referenced rule can terminate immediately
	if element.IsRule() {
		return g.canRuleReferenceTerminateImmediately(element)
	}

	return false
}

// canRuleReferenceTerminateImmediately checks if a rule reference can terminate immediately
func (g *DependencyGraph) canRuleReferenceTerminateImmediately(element Element) bool {
	var referencedRuleName string

	// Extract rule name based on element value type
	switch value := element.Value.(type) {
	case ReferenceValue:
		referencedRuleName = value.Name
	case BlockValue:
		// For block values, we need to check if any alternative in the block can terminate immediately
		return g.canBlockValueTerminateImmediately(value)
	default:
		return false
	}

	// Check if the referenced rule exists and can terminate immediately
	referencedNode := g.GetNode(referencedRuleName)
	if referencedNode == nil {
		// Handle ANTLR built-in tokens that are always immediately terminal
		if isAntlrBuiltinToken(referencedRuleName) {
			return true
		}
		// Rule not found - could be a forward reference or external rule
		// For now, we'll be conservative and assume it cannot terminate immediately
		return false
	}

	return referencedNode.HasImmediatelyTerminalAlternatives
}

// canBlockValueTerminateImmediately checks if a block value can terminate immediately
func (g *DependencyGraph) canBlockValueTerminateImmediately(block BlockValue) bool {
	// A block can terminate immediately if any of its alternatives can terminate immediately
	for _, alt := range block.Alternatives {
		if g.canAlternativeTerminateImmediately(alt) {
			return true
		}
	}
	return false
}

// CanAlternativeTerminateImmediately checks if an alternative can terminate immediately (exported for testing)
func (g *DependencyGraph) CanAlternativeTerminateImmediately(alt Alternative) bool {
	return g.canAlternativeTerminateImmediately(alt)
}

// CanElementTerminateImmediately checks if a single element can terminate immediately (exported for testing)
func (g *DependencyGraph) CanElementTerminateImmediately(element Element) bool {
	return g.canElementTerminateImmediately(element)
}

// ValidateGrammar checks if all rules have at least one immediately terminal alternative
func (g *DependencyGraph) ValidateGrammar() error {
	var invalidRules []string

	for ruleName, node := range g.Nodes {
		if !node.HasImmediatelyTerminalAlternatives {
			invalidRules = append(invalidRules, ruleName)
		}
	}

	if len(invalidRules) > 0 {
		return fmt.Errorf("grammar validation failed: the following rules have no immediately terminal alternatives: %v", invalidRules)
	}

	return nil
}

// GetImmediatelyTerminalAlternatives returns the indices of immediately terminal alternatives for a rule
func (g *DependencyGraph) GetImmediatelyTerminalAlternatives(ruleName string) []int {
	node := g.GetNode(ruleName)
	if node == nil {
		return nil
	}
	return node.ImmediatelyTerminalAlternativeIndex
}

// HasImmediatelyTerminalAlternatives checks if a rule has immediately terminal alternatives
func (g *DependencyGraph) HasImmediatelyTerminalAlternatives(ruleName string) bool {
	node := g.GetNode(ruleName)
	if node == nil {
		return false
	}
	return node.HasImmediatelyTerminalAlternatives
}

// PrintAnalysisResults prints the dependency graph analysis results for debugging
func (g *DependencyGraph) PrintAnalysisResults() {
	fmt.Println("=== Dependency Graph Analysis Results ===")
	for ruleName, node := range g.Nodes {
		fmt.Printf("Rule: %s (lexer=%t)\n", ruleName, node.IsLexer)
		fmt.Printf("  HasImmediatelyTerminalAlternatives: %t\n", node.HasImmediatelyTerminalAlternatives)
		fmt.Printf("  ImmediatelyTerminalAlternativeIndex: %v\n", node.ImmediatelyTerminalAlternativeIndex)
		fmt.Printf("  Total alternatives: %d\n", len(node.Alternatives))
		fmt.Println()
	}
}

// isAntlrBuiltinToken checks if a token name refers to an ANTLR built-in token
// that should always be considered terminal
func isAntlrBuiltinToken(tokenName string) bool {
	// ANTLR built-in tokens that are always terminal
	builtinTokens := map[string]bool{
		"EOF":   true, // End-of-file token
		"<EOF>": true, // Alternative EOF notation
	}

	return builtinTokens[tokenName]
}
