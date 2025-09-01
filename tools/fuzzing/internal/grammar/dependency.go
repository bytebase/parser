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
	RuleName                 string        // Rule name (e.g., "selectStmt", "expr")
	HasTerminalAlternatives  bool          // Can reach terminal without recursion
	Alternatives            []Alternative  // All alternatives for this rule
	TerminalAlternativeIndex []int         // Indices of alternatives that can terminate
	IsLexer                 bool          // Whether this is a lexer rule
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
		RuleName:                 ruleName,
		HasTerminalAlternatives:  false,
		Alternatives:            rule.Alternatives,
		TerminalAlternativeIndex: []int{},
		IsLexer:                 rule.IsLexer,
	}
	g.Nodes[ruleName] = node
}

// GetNode retrieves a node by rule name
func (g *DependencyGraph) GetNode(ruleName string) *GraphNode {
	return g.Nodes[ruleName]
}

// AnalyzeTerminalReachability performs terminal reachability analysis on the graph
func (g *DependencyGraph) AnalyzeTerminalReachability() error {
	return g.AnalyzeTerminalReachabilityWithValidation(false)
}

// AnalyzeTerminalReachabilityWithValidation performs terminal reachability analysis with optional validation
func (g *DependencyGraph) AnalyzeTerminalReachabilityWithValidation(validateUnterminated bool) error {
	// Phase 1: Mark lexer rules as terminal
	g.markLexerRulesAsTerminal()
	
	// Phase 2: Propagate terminal reachability using fixed-point iteration
	g.propagateTerminalReachability()
	
	// Phase 3: Check for unterminated nodes and report error (only if requested)
	if validateUnterminated {
		return g.validateTerminalReachability()
	}
	
	return nil
}

// markLexerRulesAsTerminal marks all lexer rules as having terminal alternatives
func (g *DependencyGraph) markLexerRulesAsTerminal() {
	for _, node := range g.Nodes {
		if node.IsLexer {
			node.HasTerminalAlternatives = true
			// All lexer alternatives are considered terminal
			for i := range node.Alternatives {
				node.TerminalAlternativeIndex = append(node.TerminalAlternativeIndex, i)
			}
		}
	}
}

// propagateTerminalReachability uses fixed-point iteration to determine which rules can terminate
func (g *DependencyGraph) propagateTerminalReachability() {
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
			
			// Check each alternative to see if it can terminate
			for altIndex, alt := range node.Alternatives {
				// Skip if this alternative is already marked as terminal
				if g.isAlternativeAlreadyMarked(node, altIndex) {
					continue
				}
				
				if g.canAlternativeTerminate(alt) {
					if !node.HasTerminalAlternatives {
						node.HasTerminalAlternatives = true
						changed = true
					}
					node.TerminalAlternativeIndex = append(node.TerminalAlternativeIndex, altIndex)
					changed = true
				}
			}
		}
	}
	
	if iterations >= maxIterations {
		fmt.Printf("Warning: Terminal reachability analysis reached max iterations (%d)\n", maxIterations)
	}
}

// validateTerminalReachability checks for rules without terminal alternatives and reports errors
func (g *DependencyGraph) validateTerminalReachability() error {
	var unterminatedRules []string
	
	for ruleName, node := range g.Nodes {
		if !node.HasTerminalAlternatives {
			unterminatedRules = append(unterminatedRules, ruleName)
		}
	}
	
	if len(unterminatedRules) > 0 {
		return fmt.Errorf("grammar contains %d rules without terminal alternatives: %v", 
			len(unterminatedRules), unterminatedRules)
	}
	
	return nil
}

// isAlternativeAlreadyMarked checks if an alternative is already in the terminal list
func (g *DependencyGraph) isAlternativeAlreadyMarked(node *GraphNode, altIndex int) bool {
	for _, terminalIndex := range node.TerminalAlternativeIndex {
		if terminalIndex == altIndex {
			return true
		}
	}
	return false
}

// CanAlternativeTerminate checks if an alternative can terminate without recursion (exported for testing)
func (g *DependencyGraph) CanAlternativeTerminate(alt Alternative) bool {
	return g.canAlternativeTerminate(alt)
}

// CanElementTerminate checks if a single element can terminate (exported for testing)
func (g *DependencyGraph) CanElementTerminate(element Element) bool {
	return g.canElementTerminate(element)
}

// CanBlockValueTerminate checks if a block value can terminate (exported for testing)
func (g *DependencyGraph) CanBlockValueTerminate(block BlockValue) bool {
	return g.canBlockValueTerminate(block)
}

// canAlternativeTerminate checks if an alternative can terminate without recursion
func (g *DependencyGraph) canAlternativeTerminate(alt Alternative) bool {
	// Empty alternative (ε-transition) can always terminate
	if len(alt.Elements) == 0 {
		return true
	}
	
	// Check each element in the alternative
	for _, element := range alt.Elements {
		if !g.canElementTerminate(element) {
			return false
		}
	}
	
	return true
}

// canElementTerminate checks if a single element can terminate
func (g *DependencyGraph) canElementTerminate(element Element) bool {
	// Terminal elements (literals) can always terminate
	if element.IsTerminal() {
		return true
	}
	
	// Handle quantified elements
	if element.IsQuantified() {
		// * and ? quantifiers can generate 0 occurrences, so they can terminate
		if element.Quantifier == ZERO_MORE || element.Quantifier == OPTIONAL_Q {
			return true
		}
		// + quantifier requires at least one occurrence, so check the referenced rule
		if element.Quantifier == ONE_MORE {
			return g.canRuleReferenceTerminate(element)
		}
	}
	
	// For rule references, check if the referenced rule can terminate
	if element.IsRule() {
		return g.canRuleReferenceTerminate(element)
	}
	
	return false
}

// canRuleReferenceTerminate checks if a rule reference can terminate
func (g *DependencyGraph) canRuleReferenceTerminate(element Element) bool {
	var referencedRuleName string
	
	// Extract rule name based on element value type
	switch value := element.Value.(type) {
	case ReferenceValue:
		referencedRuleName = value.Name
	case BlockValue:
		// For block values, we need to check if any alternative in the block can terminate
		return g.canBlockValueTerminate(value)
	default:
		return false
	}
	
	// Check if the referenced rule exists and can terminate
	referencedNode := g.GetNode(referencedRuleName)
	if referencedNode == nil {
		// Handle ANTLR built-in tokens that are always terminal
		if isAntlrBuiltinToken(referencedRuleName) {
			return true
		}
		// Rule not found - could be a forward reference or external rule
		// For now, we'll be conservative and assume it cannot terminate
		return false
	}
	
	return referencedNode.HasTerminalAlternatives
}

// canBlockValueTerminate checks if a block value can terminate
func (g *DependencyGraph) canBlockValueTerminate(block BlockValue) bool {
	// A block can terminate if any of its alternatives can terminate
	for _, alt := range block.Alternatives {
		if g.canAlternativeTerminate(alt) {
			return true
		}
	}
	return false
}

// ValidateGrammar checks if all rules have at least one terminal alternative
func (g *DependencyGraph) ValidateGrammar() error {
	var invalidRules []string
	
	for ruleName, node := range g.Nodes {
		if !node.HasTerminalAlternatives {
			invalidRules = append(invalidRules, ruleName)
		}
	}
	
	if len(invalidRules) > 0 {
		return fmt.Errorf("grammar validation failed: the following rules have no terminal alternatives: %v", invalidRules)
	}
	
	return nil
}

// GetTerminalAlternatives returns the indices of terminal alternatives for a rule
func (g *DependencyGraph) GetTerminalAlternatives(ruleName string) []int {
	node := g.GetNode(ruleName)
	if node == nil {
		return nil
	}
	return node.TerminalAlternativeIndex
}

// HasTerminalAlternatives checks if a rule has terminal alternatives
func (g *DependencyGraph) HasTerminalAlternatives(ruleName string) bool {
	node := g.GetNode(ruleName)
	if node == nil {
		return false
	}
	return node.HasTerminalAlternatives
}

// PrintAnalysisResults prints the dependency graph analysis results for debugging
func (g *DependencyGraph) PrintAnalysisResults() {
	fmt.Println("=== Dependency Graph Analysis Results ===")
	for ruleName, node := range g.Nodes {
		fmt.Printf("Rule: %s (lexer=%t)\n", ruleName, node.IsLexer)
		fmt.Printf("  HasTerminalAlternatives: %t\n", node.HasTerminalAlternatives)
		fmt.Printf("  TerminalAlternativeIndex: %v\n", node.TerminalAlternativeIndex)
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