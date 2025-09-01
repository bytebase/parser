package tests

import (
	"fmt"
	"path/filepath"
	"sort"
	"testing"

	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
)

func TestPostgreSQLTerminalCompleteness(t *testing.T) {
	repoRoot := getRepoRoot()

	// PostgreSQL grammar file paths
	lexerPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLLexer.g4")
	parserPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLParser.g4")

	// Parse grammar files
	parsedGrammar, err := grammar.ParseAndMergeGrammarFiles([]string{lexerPath, parserPath})
	if err != nil {
		t.Fatalf("Failed to parse grammar files: %v", err)
	}

	depGraph := parsedGrammar.GetDependencyGraph()

	// Analyze all nodes
	fmt.Println("=== PostgreSQL Grammar Terminal Analysis ===")

	totalNodes := len(depGraph.Nodes)
	terminalNodes := 0
	nonTerminalNodes := 0

	var terminalRules []string
	var nonTerminalRules []string

	lexerTerminal := 0
	lexerNonTerminal := 0
	parserTerminal := 0
	parserNonTerminal := 0

	for ruleName, node := range depGraph.Nodes {
		if node.HasTerminalAlternatives {
			terminalNodes++
			terminalRules = append(terminalRules, ruleName)
			if node.IsLexer {
				lexerTerminal++
			} else {
				parserTerminal++
			}
		} else {
			nonTerminalNodes++
			nonTerminalRules = append(nonTerminalRules, ruleName)
			if node.IsLexer {
				lexerNonTerminal++
			} else {
				parserNonTerminal++
			}
		}
	}

	fmt.Printf("Total Nodes: %d\n", totalNodes)
	fmt.Printf("Terminal Nodes: %d (%.1f%%)\n", terminalNodes, float64(terminalNodes)/float64(totalNodes)*100)
	fmt.Printf("Non-Terminal Nodes: %d (%.1f%%)\n", nonTerminalNodes, float64(nonTerminalNodes)/float64(totalNodes)*100)
	fmt.Println()

	fmt.Printf("Lexer Rules: Terminal=%d, Non-Terminal=%d\n", lexerTerminal, lexerNonTerminal)
	fmt.Printf("Parser Rules: Terminal=%d, Non-Terminal=%d\n", parserTerminal, parserNonTerminal)
	fmt.Println()

	// Show non-terminal rules (these should ideally be zero)
	if len(nonTerminalRules) > 0 {
		sort.Strings(nonTerminalRules)
		fmt.Printf("❌ Non-Terminal Rules (%d):\n", len(nonTerminalRules))
		for i, ruleName := range nonTerminalRules {
			node := depGraph.GetNode(ruleName)
			ruleType := "PARSER"
			if node.IsLexer {
				ruleType = "LEXER"
			}
			fmt.Printf("  %d. %s (%s, %d alternatives)\n", i+1, ruleName, ruleType, len(node.Alternatives))
		}
		fmt.Println()
	} else {
		fmt.Println("✅ All rules have terminal alternatives!")
	}

	// Test: If your hypothesis is correct, this should pass
	if nonTerminalNodes == 0 {
		t.Log("✅ HYPOTHESIS CONFIRMED: All PostgreSQL rules have terminal alternatives")
	} else {
		t.Errorf("❌ HYPOTHESIS REJECTED: %d rules have no terminal alternatives", nonTerminalNodes)

		// Analyze WHY these rules don't have terminal alternatives
		fmt.Println("=== Analysis of Non-Terminal Rules ===")
		analyzeNonTerminalRules(parsedGrammar, depGraph, nonTerminalRules[:min(5, len(nonTerminalRules))])
	}
}

func TestSpecificNonTerminalRules(t *testing.T) {
	repoRoot := getRepoRoot()

	lexerPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLLexer.g4")
	parserPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLParser.g4")

	parsedGrammar, err := grammar.ParseAndMergeGrammarFiles([]string{lexerPath, parserPath})
	if err != nil {
		t.Fatalf("Failed to parse grammar files: %v", err)
	}

	depGraph := parsedGrammar.GetDependencyGraph()

	// Test specific rules that should be terminal based on our earlier analysis
	expectedTerminalRules := []string{
		"columnref",       // Should be terminal (colid + indirection)
		"c_expr",          // Should be terminal (has columnref alternative)
		"a_expr_typecast", // Should be terminal (depends on c_expr)
		"a_expr_collate",  // Should be terminal (depends on a_expr_typecast)
	}

	fmt.Println("=== Testing Expected Terminal Rules ===")
	for _, ruleName := range expectedTerminalRules {
		node := depGraph.GetNode(ruleName)
		if node == nil {
			t.Errorf("Rule %s not found", ruleName)
			continue
		}

		fmt.Printf("%s: HasTerminal=%t, TerminalAlts=%v\n",
			ruleName, node.HasTerminalAlternatives, node.TerminalAlternativeIndex)

		if !node.HasTerminalAlternatives {
			t.Errorf("Expected %s to be terminal, but it's not", ruleName)
		}
	}
}

func analyzeNonTerminalRules(parsedGrammar *grammar.ParsedGrammar, depGraph *grammar.DependencyGraph, ruleNames []string) {
	for _, ruleName := range ruleNames {
		rule := parsedGrammar.GetRule(ruleName)
		if rule == nil {
			continue
		}

		fmt.Printf("\n--- Analyzing %s ---\n", ruleName)
		fmt.Printf("Type: %s, Alternatives: %d\n",
			map[bool]string{true: "LEXER", false: "PARSER"}[rule.IsLexer], len(rule.Alternatives))

		for altIndex, alt := range rule.Alternatives {
			canTerminate := depGraph.CanAlternativeTerminate(alt)
			fmt.Printf("  Alt %d (%d elements): canTerminate=%t\n", altIndex, len(alt.Elements), canTerminate)

			for elemIndex, element := range alt.Elements {
				canElemTerminate := depGraph.CanElementTerminate(element)
				fmt.Printf("    Elem %d: %s", elemIndex, element.Value.String())

				if element.IsQuantified() {
					fmt.Printf("[%v]", element.Quantifier)
				}

				fmt.Printf(" → canTerminate=%t", canElemTerminate)

				// If it's a rule reference, show the referenced rule's status
				if element.IsRule() {
					if refValue, ok := element.Value.(grammar.ReferenceValue); ok {
						referencedNode := depGraph.GetNode(refValue.Name)
						if referencedNode != nil {
							fmt.Printf(" (ref: %s, hasTerminal=%t)", refValue.Name, referencedNode.HasTerminalAlternatives)
						} else {
							fmt.Printf(" (ref: %s, NOT_FOUND)", refValue.Name)
						}
					}
				}

				fmt.Println()
			}
		}
	}
}

// Test to validate that our terminal propagation algorithm is working correctly
func TestManualTerminalPropagation(t *testing.T) {
	repoRoot := getRepoRoot()

	lexerPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLLexer.g4")
	parserPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLParser.g4")

	parsedGrammar, err := grammar.ParseAndMergeGrammarFiles([]string{lexerPath, parserPath})
	if err != nil {
		t.Fatalf("Failed to parse grammar files: %v", err)
	}

	// Create fresh dependency graph and run manual propagation
	freshGraph := grammar.NewDependencyGraph()

	// Add all nodes
	for ruleName, rule := range parsedGrammar.GetAllRules() {
		freshGraph.AddNode(ruleName, rule)
	}

	// Mark lexer rules as terminal
	initialTerminalCount := 0
	for _, node := range freshGraph.Nodes {
		if node.IsLexer {
			node.HasTerminalAlternatives = true
			for i := range node.Alternatives {
				node.TerminalAlternativeIndex = append(node.TerminalAlternativeIndex, i)
			}
			initialTerminalCount++
		}
	}

	fmt.Printf("Starting with %d lexer rules marked as terminal\n", initialTerminalCount)

	// Manual propagation with more iterations
	maxIterations := 100
	totalNewTerminals := 0

	for iteration := 0; iteration < maxIterations; iteration++ {
		changed := false
		newTerminalsThisIteration := 0

		for _, node := range freshGraph.Nodes {
			if node.IsLexer || node.HasTerminalAlternatives {
				continue
			}

			for altIndex, alt := range node.Alternatives {
				// Check if already marked
				alreadyMarked := false
				for _, termIndex := range node.TerminalAlternativeIndex {
					if termIndex == altIndex {
						alreadyMarked = true
						break
					}
				}
				if alreadyMarked {
					continue
				}

				// Check if this alternative can terminate
				if canAlternativeTerminateManual(alt, freshGraph) {
					if !node.HasTerminalAlternatives {
						node.HasTerminalAlternatives = true
						changed = true
						newTerminalsThisIteration++
						totalNewTerminals++
					}
					node.TerminalAlternativeIndex = append(node.TerminalAlternativeIndex, altIndex)
				}
			}
		}

		if newTerminalsThisIteration > 0 {
			fmt.Printf("Iteration %d: +%d new terminal rules (total: %d)\n",
				iteration+1, newTerminalsThisIteration, initialTerminalCount+totalNewTerminals)
		}

		if !changed {
			fmt.Printf("Converged after %d iterations\n", iteration+1)
			break
		}

		if iteration == maxIterations-1 {
			fmt.Printf("Reached max iterations (%d)\n", maxIterations)
		}
	}

	// Count final results
	finalTerminalCount := 0
	finalNonTerminalCount := 0

	for _, node := range freshGraph.Nodes {
		if node.HasTerminalAlternatives {
			finalTerminalCount++
		} else {
			finalNonTerminalCount++
		}
	}

	fmt.Printf("\nFinal Results:\n")
	fmt.Printf("Terminal: %d\n", finalTerminalCount)
	fmt.Printf("Non-Terminal: %d\n", finalNonTerminalCount)

	if finalNonTerminalCount == 0 {
		t.Log("✅ Manual propagation: All rules are terminal!")
	} else {
		t.Logf("❌ Manual propagation: Still %d non-terminal rules", finalNonTerminalCount)
	}
}

func canAlternativeTerminateManual(alt grammar.Alternative, graph *grammar.DependencyGraph) bool {
	// Empty alternative is always terminal
	if len(alt.Elements) == 0 {
		return true
	}

	// All elements must be able to terminate
	for _, element := range alt.Elements {
		if !canElementTerminateManual(element, graph) {
			return false
		}
	}

	return true
}

func canElementTerminateManual(element grammar.Element, graph *grammar.DependencyGraph) bool {
	// Terminal elements (literals) can always terminate
	if element.IsTerminal() {
		return true
	}

	// Handle quantified elements - THIS IS KEY!
	if element.IsQuantified() {
		// * and ? quantifiers can generate 0 occurrences, so they can terminate
		if element.Quantifier == grammar.ZERO_MORE || element.Quantifier == grammar.OPTIONAL_Q {
			return true // Can be empty, so always terminal
		}
		// + quantifier requires at least one occurrence, so check the content
	}

	// For rule references
	if element.IsRule() {
		switch value := element.Value.(type) {
		case grammar.ReferenceValue:
			referencedNode := graph.GetNode(value.Name)
			if referencedNode == nil {
				// Handle ANTLR built-in tokens like EOF
				return isBuiltinTerminal(value.Name)
			}
			return referencedNode.HasTerminalAlternatives
		case grammar.BlockValue:
			// A block can terminate if any of its alternatives can terminate
			for _, alt := range value.Alternatives {
				if canAlternativeTerminateManual(alt, graph) {
					return true
				}
			}
			return false
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// isBuiltinTerminal checks if a token name refers to an ANTLR built-in token
func isBuiltinTerminal(tokenName string) bool {
	builtinTokens := map[string]bool{
		"EOF":   true, // End-of-file token
		"<EOF>": true, // Alternative EOF notation
	}
	return builtinTokens[tokenName]
}
