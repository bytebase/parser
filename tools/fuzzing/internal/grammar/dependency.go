package grammar

import (
	"fmt"
)

// DependencyGraph represents the dependency relationships between grammar rules
type DependencyGraph struct {
	Nodes map[string]*GraphNode
	Edges map[string][]string // Adjacency list: rule -> referenced rules
	SCCs  [][]string          // List of SCCs (each SCC is a list of rule names)
}

// GraphNode represents a single rule in the dependency graph
type GraphNode struct {
	RuleName                            string        // Rule name (e.g., "selectStmt", "expr")
	Alternatives                        []Alternative // All alternatives for this rule
	HasImmediatelyTerminalAlternatives  bool          // Has at least one immediately terminal alternative
	ImmediatelyTerminalAlternativeIndex []int         // Indices of alternatives that are immediately terminal
	IsLexer                             bool          // Whether this is a lexer rule
	SCCID                               int           // Which SCC this node belongs to (-1 if not computed)
	SCCSize                             int           // Size of the SCC this node belongs to
	IsRecursive                         bool          // True if part of a recursive SCC (size > 1 or self-loop)
}

// NewDependencyGraph creates a new dependency graph
func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		Nodes: make(map[string]*GraphNode),
		Edges: make(map[string][]string),
		SCCs:  [][]string{},
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
		SCCID:                               -1,
		SCCSize:                             0,
		IsRecursive:                         false,
	}
	g.Nodes[ruleName] = node
	
	// Build edges for this node
	g.buildEdgesForNode(ruleName, rule)
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
	// Phase 1: Compute SCCs to identify recursive rule groups
	g.ComputeSCCs()
	g.PrintSCCAnalysis() // Debug output
	
	// Phase 2: Mark lexer rules as immediately terminal
	g.markLexerRulesAsImmediatelyTerminal()

	// Phase 3: Analyze immediately terminal alternatives
	g.analyzeImmediatelyTerminalAlternatives()

	// Phase 4: Check for nodes without immediately terminal alternatives and report error (only if requested)
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

// buildEdgesForNode builds the edge list for a given rule node
func (g *DependencyGraph) buildEdgesForNode(ruleName string, rule *Rule) {
	referencedRules := make(map[string]bool)
	
	// Scan all alternatives for rule references
	for _, alt := range rule.Alternatives {
		g.collectRuleReferences(alt, referencedRules)
	}
	
	// Convert map to slice and store as edges
	edges := []string{}
	for ref := range referencedRules {
		edges = append(edges, ref)
	}
	g.Edges[ruleName] = edges
}

// collectRuleReferences collects all rule references in an alternative
func (g *DependencyGraph) collectRuleReferences(alt Alternative, refs map[string]bool) {
	for _, element := range alt.Elements {
		g.collectElementReferences(element, refs)
	}
}

// collectElementReferences collects rule references from a single element
func (g *DependencyGraph) collectElementReferences(element Element, refs map[string]bool) {
	if element.IsRule() {
		switch value := element.Value.(type) {
		case ReferenceValue:
			// Add all rule references (we'll filter lexer rules later if needed)
			// Don't check if node exists yet - it might not be added yet
			refs[value.Name] = true
		case BlockValue:
			// Collect references from block alternatives
			for _, alt := range value.Alternatives {
				g.collectRuleReferences(alt, refs)
			}
		}
	}
}

// RebuildEdges rebuilds all edges after all nodes have been added
func (g *DependencyGraph) RebuildEdges() {
	g.Edges = make(map[string][]string)
	
	for ruleName, node := range g.Nodes {
		referencedRules := make(map[string]bool)
		
		// Scan all alternatives for rule references
		for _, alt := range node.Alternatives {
			g.collectRuleReferences(alt, referencedRules)
		}
		
		// Filter out lexer rules and non-existent rules
		edges := []string{}
		for ref := range referencedRules {
			if refNode := g.GetNode(ref); refNode != nil && !refNode.IsLexer {
				edges = append(edges, ref)
			}
		}
		g.Edges[ruleName] = edges
	}
}

// ComputeSCCs computes strongly connected components using Tarjan's algorithm
func (g *DependencyGraph) ComputeSCCs() {
	// Only rebuild edges if they're empty (allows manual edge setup for testing)
	if len(g.Edges) == 0 {
		g.RebuildEdges()
	}
	// Initialize for Tarjan's algorithm
	index := 0
	stack := []string{}
	indices := make(map[string]int)
	lowlinks := make(map[string]int)
	onStack := make(map[string]bool)
	
	// Helper function for Tarjan's strongconnect
	var strongconnect func(v string)
	strongconnect = func(v string) {
		// Set the depth index for v to the smallest unused index
		indices[v] = index
		lowlinks[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true
		
		// Consider successors of v
		for _, w := range g.Edges[v] {
			if _, ok := indices[w]; !ok {
				// Successor w has not yet been visited; recurse on it
				strongconnect(w)
				if lowlinks[w] < lowlinks[v] {
					lowlinks[v] = lowlinks[w]
				}
			} else if onStack[w] {
				// Successor w is in stack S and hence in the current SCC
				if indices[w] < lowlinks[v] {
					lowlinks[v] = indices[w]
				}
			}
		}
		
		// If v is a root node, pop the stack and print an SCC
		if lowlinks[v] == indices[v] {
			scc := []string{}
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			g.SCCs = append(g.SCCs, scc)
		}
	}
	
	// Clear existing SCCs
	g.SCCs = [][]string{}
	
	// Run algorithm for all unvisited nodes
	for ruleName := range g.Nodes {
		if _, ok := indices[ruleName]; !ok {
			strongconnect(ruleName)
		}
	}
	
	// Update nodes with their SCC information
	for sccID, scc := range g.SCCs {
		sccSize := len(scc)
		isRecursive := sccSize > 1
		
		// Check for self-loops if single node SCC
		if sccSize == 1 {
			ruleName := scc[0]
			for _, ref := range g.Edges[ruleName] {
				if ref == ruleName {
					isRecursive = true
					break
				}
			}
		}
		
		// Update all nodes in this SCC
		for _, ruleName := range scc {
			if node := g.GetNode(ruleName); node != nil {
				node.SCCID = sccID
				node.SCCSize = sccSize
				node.IsRecursive = isRecursive
			}
		}
	}
}

// PrintSCCAnalysis prints the SCC analysis results for debugging
func (g *DependencyGraph) PrintSCCAnalysis() {
	fmt.Println("\n=== SCC Analysis Results ===")
	fmt.Printf("Total SCCs: %d\n", len(g.SCCs))
	
	recursiveSCCs := 0
	selfLoopSCCs := 0
	largestSCC := 0
	for i, scc := range g.SCCs {
		if len(scc) > 1 {
			recursiveSCCs++
			if len(scc) > largestSCC {
				largestSCC = len(scc)
			}
			// Print first 5 multi-node SCCs with more detail
			if recursiveSCCs <= 5 {
				fmt.Printf("\nSCC %d (RECURSIVE - mutual, size=%d):\n", i, len(scc))
				// Print first 20 nodes of the SCC for better visibility
				fmt.Printf("  Members: ")
				for j, node := range scc {
					if j < 20 {
						fmt.Printf("%s ", node)
						if j == 19 && len(scc) > 20 {
							fmt.Printf("\n           ... and %d more", len(scc)-20)
						}
					}
				}
				fmt.Println()
			}
		} else if len(scc) == 1 {
			// Check for self-loop
			ruleName := scc[0]
			hasSelfLoop := false
			for _, ref := range g.Edges[ruleName] {
				if ref == ruleName {
					hasSelfLoop = true
					break
				}
			}
			if hasSelfLoop {
				selfLoopSCCs++
				if selfLoopSCCs <= 10 { // Print first 10
					fmt.Printf("SCC %d (RECURSIVE - self-loop): %s\n", i, ruleName)
				}
			}
		}
	}
	
	fmt.Printf("\nMutually recursive SCCs (size > 1): %d\n", recursiveSCCs)
	if recursiveSCCs > 0 {
		fmt.Printf("Largest SCC size: %d\n", largestSCC)
	}
	fmt.Printf("Self-loop SCCs (size = 1 with self-ref): %d\n", selfLoopSCCs)
	fmt.Printf("Non-recursive SCCs: %d\n", len(g.SCCs)-recursiveSCCs-selfLoopSCCs)
	
	// Print sample of recursive rules
	fmt.Println("\nSample recursive rules:")
	count := 0
	for ruleName, node := range g.Nodes {
		if node.IsRecursive && count < 10 {
			fmt.Printf("  %s (SCC %d, size %d)\n", ruleName, node.SCCID, node.SCCSize)
			count++
		}
	}
	fmt.Println("=============================")
}
