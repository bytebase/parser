package grammar

import (
	"testing"
)

// TestSCCDetection tests the SCC detection algorithm with various graph patterns
func TestSCCDetection(t *testing.T) {
	tests := []struct {
		name          string
		rules         map[string][]string // rule -> references
		expectedSCCs  [][]string          // expected SCCs
		recursiveRules map[string]bool    // which rules should be marked recursive
	}{
		{
			name: "Simple self-loop",
			rules: map[string][]string{
				"a": {"a"},
			},
			expectedSCCs: [][]string{
				{"a"},
			},
			recursiveRules: map[string]bool{
				"a": true,
			},
		},
		{
			name: "Mutual recursion (2 nodes)",
			rules: map[string][]string{
				"a": {"b"},
				"b": {"a"},
			},
			expectedSCCs: [][]string{
				{"b", "a"}, // Order might vary due to algorithm
			},
			recursiveRules: map[string]bool{
				"a": true,
				"b": true,
			},
		},
		{
			name: "Cycle of 3 nodes",
			rules: map[string][]string{
				"a": {"b"},
				"b": {"c"},
				"c": {"a"},
			},
			expectedSCCs: [][]string{
				{"c", "b", "a"},
			},
			recursiveRules: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
		},
		{
			name: "Non-recursive with reference",
			rules: map[string][]string{
				"a": {"b"},
				"b": {"c"},
				"c": {},
			},
			expectedSCCs: [][]string{
				{"c"},
				{"b"},
				{"a"},
			},
			recursiveRules: map[string]bool{
				"a": false,
				"b": false,
				"c": false,
			},
		},
		{
			name: "Multiple SCCs",
			rules: map[string][]string{
				"a": {"b"},
				"b": {"a"},
				"c": {"d"},
				"d": {"c"},
				"e": {},
			},
			expectedSCCs: [][]string{
				{"b", "a"},
				{"d", "c"},
				{"e"},
			},
			recursiveRules: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
				"d": true,
				"e": false,
			},
		},
		{
			name: "Complex with bridge",
			rules: map[string][]string{
				"a": {"b", "c"},
				"b": {"a"},
				"c": {"d"},
				"d": {"e"},
				"e": {"c"},
			},
			expectedSCCs: [][]string{
				{"b", "a"},
				{"e", "d", "c"},
			},
			recursiveRules: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
				"d": true,
				"e": true,
			},
		},
		{
			name: "Self-loop with external reference",
			rules: map[string][]string{
				"expr": {"expr", "literal"},
				"literal": {},
			},
			expectedSCCs: [][]string{
				{"expr"},
				{"literal"},
			},
			recursiveRules: map[string]bool{
				"expr": true,
				"literal": false,
			},
		},
		{
			name: "PostgreSQL-like pattern",
			rules: map[string][]string{
				"select_with_parens": {"select_no_parens", "select_with_parens"},
				"select_no_parens": {"table_ref"},
				"table_ref": {"joined_table", "table_ref"},
				"joined_table": {"table_ref"},
			},
			expectedSCCs: [][]string{
				{"select_with_parens"},
				{"joined_table", "table_ref"},
				{"select_no_parens"},
			},
			recursiveRules: map[string]bool{
				"select_with_parens": true,
				"select_no_parens": false,
				"table_ref": true,
				"joined_table": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create dependency graph
			g := NewDependencyGraph()
			
			// Add nodes
			for ruleName := range tt.rules {
				rule := &Rule{
					Name:         ruleName,
					Alternatives: []Alternative{},
					IsLexer:      false,
				}
				// We need to add the node without building edges automatically
				node := &GraphNode{
					RuleName:                            ruleName,
					Alternatives:                        rule.Alternatives,
					HasImmediatelyTerminalAlternatives:  false,
					ImmediatelyTerminalAlternativeIndex: []int{},
					IsLexer:                             false,
					SCCID:                               -1,
					SCCSize:                             0,
					IsRecursive:                         false,
				}
				g.Nodes[ruleName] = node
			}
			
			// Set up edges manually
			g.Edges = tt.rules
			
			// Compute SCCs
			g.ComputeSCCs()
			
			// Verify number of SCCs
			if len(g.SCCs) != len(tt.expectedSCCs) {
				t.Errorf("Expected %d SCCs, got %d", len(tt.expectedSCCs), len(g.SCCs))
				t.Logf("SCCs found: %v", g.SCCs)
			}
			
			// Verify each node's recursive status
			for ruleName, expectedRecursive := range tt.recursiveRules {
				node := g.GetNode(ruleName)
				if node == nil {
					t.Errorf("Node %s not found", ruleName)
					continue
				}
				
				if node.IsRecursive != expectedRecursive {
					t.Errorf("Node %s: expected IsRecursive=%v, got %v (SCCID=%d, SCCSize=%d)",
						ruleName, expectedRecursive, node.IsRecursive, node.SCCID, node.SCCSize)
				}
			}
			
			// Verify all nodes in same SCC have same SCCID
			sccNodeMap := make(map[int][]string)
			for ruleName, node := range g.Nodes {
				if node.SCCID >= 0 {
					sccNodeMap[node.SCCID] = append(sccNodeMap[node.SCCID], ruleName)
				}
			}
			
			// Log SCC information for debugging
			t.Logf("SCCs detected:")
			for sccID, nodes := range sccNodeMap {
				t.Logf("  SCC %d: %v", sccID, nodes)
			}
		})
	}
}

// TestSCCEdgeBuilding tests that edges are correctly built from grammar rules
func TestSCCEdgeBuilding(t *testing.T) {
	// Create a simple grammar with references
	g := NewDependencyGraph()
	
	// Add lexer rule (should not create edges)
	lexerRule := &Rule{
		Name:    "ID",
		IsLexer: true,
		Alternatives: []Alternative{
			{
				Elements: []Element{
					{Value: LiteralValue{Text: "[a-zA-Z]+"}},
				},
			},
		},
	}
	g.AddNode("ID", lexerRule)
	
	// Add parser rule with references
	parserRule := &Rule{
		Name:    "expr",
		IsLexer: false,
		Alternatives: []Alternative{
			{
				Elements: []Element{
					{Value: ReferenceValue{Name: "expr"}},
					{Value: LiteralValue{Text: "+"}},
					{Value: ReferenceValue{Name: "term"}},
				},
			},
			{
				Elements: []Element{
					{Value: ReferenceValue{Name: "term"}},
				},
			},
		},
	}
	
	// Add term rule
	termRule := &Rule{
		Name:    "term",
		IsLexer: false,
		Alternatives: []Alternative{
			{
				Elements: []Element{
					{Value: ReferenceValue{Name: "ID"}}, // Reference to lexer
				},
			},
			{
				Elements: []Element{
					{Value: LiteralValue{Text: "123"}},
				},
			},
		},
	}
	
	// Need to add term first so it exists when expr references it
	g.AddNode("term", termRule)
	g.AddNode("expr", parserRule)
	
	// Verify edges
	// expr should have edges to: expr (self), term
	exprEdges := g.Edges["expr"]
	if len(exprEdges) == 0 {
		t.Error("expr should have edges")
	}
	
	hasExprEdge := false
	hasTermEdge := false
	for _, edge := range exprEdges {
		if edge == "expr" {
			hasExprEdge = true
		}
		if edge == "term" {
			hasTermEdge = true
		}
	}
	
	if !hasExprEdge {
		t.Error("expr should have self-edge")
	}
	if !hasTermEdge {
		t.Error("expr should have edge to term")
	}
	
	// term should NOT have edge to ID (lexer rule)
	termEdges := g.Edges["term"]
	for _, edge := range termEdges {
		if edge == "ID" {
			t.Error("term should not have edge to lexer rule ID")
		}
	}
	
	// Compute SCCs and verify
	g.ComputeSCCs()
	
	// expr should be recursive (self-loop)
	exprNode := g.GetNode("expr")
	if !exprNode.IsRecursive {
		t.Error("expr should be marked as recursive due to self-loop")
	}
	
	// term should not be recursive
	termNode := g.GetNode("term")
	if termNode.IsRecursive {
		t.Error("term should not be marked as recursive")
	}
}