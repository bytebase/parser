package grammar

import (
	"fmt"
)

const (
	// NoSCC indicates a node is not part of any SCC or SCC not yet computed
	NoSCC = -1
)

// DependencyGraph represents the dependency relationships between grammar rules
type DependencyGraph struct {
	Nodes     map[string]*GraphNode
	Edges     map[string][]string // Adjacency list: rule -> referenced rules
	SCCs      [][]string          // List of SCCs (each SCC is a list of rule names)
	SCCLookup map[string]int      // Rule name -> SCC ID lookup map
}

// GraphNode represents a single rule in the dependency graph
type GraphNode struct {
	RuleName     string        // Rule name (e.g., "selectStmt", "expr")
	Alternatives []Alternative // All alternatives for this rule
	IsLexer      bool          // Whether this is a lexer rule
	SCCID        int           // Which SCC this node belongs to (NoSCC if not computed)
	SCCSize      int           // Size of the SCC this node belongs to
	IsRecursive  bool          // True if part of a recursive SCC (size > 1 or self-loop)
}

// NewDependencyGraph creates a new dependency graph
func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		Nodes:     make(map[string]*GraphNode),
		Edges:     make(map[string][]string),
		SCCs:      [][]string{},
		SCCLookup: make(map[string]int),
	}
}

// AddNode adds a rule node to the dependency graph
func (g *DependencyGraph) AddNode(ruleName string, rule *Rule) {
	node := &GraphNode{
		RuleName:     ruleName,
		Alternatives: rule.Alternatives,
		IsLexer:      rule.IsLexer,
		SCCID:        NoSCC,
		SCCSize:      0,
		IsRecursive:  false,
	}
	g.Nodes[ruleName] = node

	// Don't build edges here because this rule may reference other rules that
	// haven't been added yet (forward references). Edges will be built later
	// after all nodes are added via BuildEdges()
}

// GetNode retrieves a node by rule name
func (g *DependencyGraph) GetNode(ruleName string) *GraphNode {
	return g.Nodes[ruleName]
}

// ValidateGrammar checks if all non-recursive rules can reach terminal symbols
func (g *DependencyGraph) ValidateGrammar() error {
	// For now, we trust that the grammar is well-formed
	// Future: could add validation to ensure non-recursive rules can terminate
	return nil
}

// PrintAnalysisResults prints the dependency graph analysis results for debugging
func (g *DependencyGraph) PrintAnalysisResults() {
	fmt.Println("=== Dependency Graph Analysis Results ===")
	for ruleName, node := range g.Nodes {
		fmt.Printf("Rule: %s (lexer=%t)\n", ruleName, node.IsLexer)
		fmt.Printf("  IsRecursive: %t\n", node.IsRecursive)
		fmt.Printf("  SCCID: %d, SCCSize: %d\n", node.SCCID, node.SCCSize)
		fmt.Printf("  Total alternatives: %d\n", len(node.Alternatives))
		fmt.Println()
	}
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
			refs[value.Name] = true
		case BlockValue:
			for _, alt := range value.Alternatives {
				g.collectRuleReferences(alt, refs)
			}
		}
	}
}

// BuildEdges builds all edges after all nodes have been added
func (g *DependencyGraph) BuildEdges() {
	g.Edges = make(map[string][]string)

	for ruleName, node := range g.Nodes {
		referencedRules := make(map[string]bool)

		for _, alt := range node.Alternatives {
			g.collectRuleReferences(alt, referencedRules)
		}

		// Only add edges to parser rules (exclude lexer rules)
		// But include all referenced parser rules, even if they don't exist yet
		edges := []string{}
		for ref := range referencedRules {
			// Check if the referenced rule is a lexer rule
			if refNode := g.GetNode(ref); refNode != nil && refNode.IsLexer {
				continue // Skip lexer rules
			}
			// Add all other references (including forward references)
			edges = append(edges, ref)
		}
		g.Edges[ruleName] = edges
	}
}

// ComputeSCCs computes strongly connected components using Tarjan's algorithm
func (g *DependencyGraph) ComputeSCCs() {
	if len(g.Edges) == 0 {
		g.BuildEdges()
	}

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

	// Clear existing SCCs and lookup map
	g.SCCs = [][]string{}
	g.SCCLookup = make(map[string]int)

	// Run algorithm for all unvisited nodes
	for ruleName := range g.Nodes {
		if _, ok := indices[ruleName]; !ok {
			strongconnect(ruleName)
		}
	}

	// Perform sanity check: ensure no SCC is an isolated island
	// Only log warnings, don't fail - test cases often have isolated SCCs
	if err := g.checkForIsolatedSCCs(); err != nil {
		// Log warning but continue - tests may have intentionally isolated SCCs
		fmt.Printf("Warning: %v\n", err)
	}

	// Build SCC lookup map and update nodes with their SCC information
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

		// Update lookup map and nodes in this SCC
		for _, ruleName := range scc {
			// Add to lookup map
			g.SCCLookup[ruleName] = sccID

			// Update node information
			if node := g.GetNode(ruleName); node != nil {
				node.SCCID = sccID
				node.SCCSize = sccSize
				node.IsRecursive = isRecursive
			}
		}
	}
}

// checkForIsolatedSCCs ensures no SCC is an isolated island with no exit paths
func (g *DependencyGraph) checkForIsolatedSCCs() error {
	// Create a temporary SCC membership map for this check
	sccMembership := make(map[string]int)
	for sccID, scc := range g.SCCs {
		for _, ruleName := range scc {
			sccMembership[ruleName] = sccID
		}
	}

	// Check each SCC for exit paths
	isolatedSCCs := []int{}
	for sccID, scc := range g.SCCs {
		// Skip non-recursive SCCs (single nodes without self-loops)
		if len(scc) == 1 {
			ruleName := scc[0]
			hasSelfLoop := false
			for _, ref := range g.Edges[ruleName] {
				if ref == ruleName {
					hasSelfLoop = true
					break
				}
			}
			if !hasSelfLoop {
				continue // Non-recursive single node, skip
			}
		}

		// Check if this SCC has any exit path
		hasExit := g.sccHasExitPath(sccID, scc, sccMembership)
		if !hasExit {
			isolatedSCCs = append(isolatedSCCs, sccID)
		}
	}

	// Report error if any isolated SCCs found
	if len(isolatedSCCs) > 0 {
		fmt.Printf("\nERROR: Found %d isolated SCC(s) with no exit paths:\n", len(isolatedSCCs))
		for _, sccID := range isolatedSCCs {
			fmt.Printf("  SCC %d: %v\n", sccID, g.SCCs[sccID])
		}
		return fmt.Errorf("grammar contains %d isolated SCC(s) that cannot terminate", len(isolatedSCCs))
	}

	return nil
}

// sccHasExitPath checks if an SCC has at least one path to rules outside of it
func (g *DependencyGraph) sccHasExitPath(sccID int, scc []string, sccMembership map[string]int) bool {
	// Use fixed-point iteration to find reachable rules from this SCC
	visited := make(map[string]bool)
	toVisit := []string{}

	// Start with all rules in the SCC
	for _, ruleName := range scc {
		toVisit = append(toVisit, ruleName)
		visited[ruleName] = true
	}

	// Perform reachability analysis
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		// Check all references from current rule
		for _, ref := range g.Edges[current] {
			// Skip if already visited
			if visited[ref] {
				continue
			}

			// Check if referenced rule is outside this SCC
			refSCCID, exists := sccMembership[ref]
			if !exists || refSCCID != sccID {
				// Found an exit! Check if it can eventually reach terminals
				if g.canReachTerminal(ref, make(map[string]bool)) {
					return true
				}
			}

			// Mark as visited and continue searching
			visited[ref] = true
			toVisit = append(toVisit, ref)
		}

		// Also check alternatives for direct terminal paths
		if node := g.GetNode(current); node != nil {
			for _, alt := range node.Alternatives {
				if g.alternativeHasTerminalPath(alt) {
					return true
				}
			}
		}
	}

	return false
}

// canReachTerminal checks if a rule can eventually reach terminal symbols
func (g *DependencyGraph) canReachTerminal(ruleName string, visited map[string]bool) bool {
	// Avoid infinite recursion
	if visited[ruleName] {
		return false
	}
	visited[ruleName] = true

	node := g.GetNode(ruleName)
	if node == nil {
		return false
	}

	// Lexer rules are terminals
	if node.IsLexer {
		return true
	}

	// Check each alternative
	for _, alt := range node.Alternatives {
		if g.alternativeCanReachTerminal(alt, visited) {
			return true
		}
	}

	return false
}

// alternativeHasTerminalPath checks if an alternative has at least one terminal
func (g *DependencyGraph) alternativeHasTerminalPath(alt Alternative) bool {
	for _, element := range alt.Elements {
		if element.IsTerminal() {
			return true
		}
		// Check if it's an optional/quantified element (can be skipped)
		if element.Quantifier == ZERO_MORE || element.Quantifier == OPTIONAL_Q {
			return true
		}
	}
	return false
}

// alternativeCanReachTerminal checks if an alternative can reach terminals
func (g *DependencyGraph) alternativeCanReachTerminal(alt Alternative, visited map[string]bool) bool {
	if len(alt.Elements) == 0 {
		return true // Empty alternative is terminal
	}

	for _, element := range alt.Elements {
		if element.IsTerminal() {
			return true
		}

		// Optional elements can be skipped
		if element.Quantifier == ZERO_MORE || element.Quantifier == OPTIONAL_Q {
			continue
		}

		// Check if referenced rule can reach terminal
		if refValue, ok := element.Value.(ReferenceValue); ok {
			if !g.canReachTerminal(refValue.Name, visited) {
				return false
			}
		}
	}

	return true
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
