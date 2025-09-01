package tests

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
)

func TestDependencyGraphConstruction(t *testing.T) {
	repoRoot := getRepoRoot()
	
	// PostgreSQL grammar file paths
	lexerPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLLexer.g4")
	parserPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLParser.g4")

	// Parse grammar files
	parsedGrammar, err := grammar.ParseAndMergeGrammarFiles([]string{lexerPath, parserPath})
	if err != nil {
		t.Fatalf("Failed to parse grammar files: %v", err)
	}

	// Test dependency graph exists
	depGraph := parsedGrammar.GetDependencyGraph()
	if depGraph == nil {
		t.Fatal("Dependency graph was not created")
	}

	// Test that nodes were created for rules
	totalRules := len(parsedGrammar.GetAllRules())
	if len(depGraph.Nodes) != totalRules {
		t.Errorf("Expected %d nodes in dependency graph, got %d", totalRules, len(depGraph.Nodes))
	}

	// Test lexer rules are marked as terminal
	lexerTerminalCount := 0
	for ruleName := range parsedGrammar.LexerRules {
		if depGraph.HasTerminalAlternatives(ruleName) {
			lexerTerminalCount++
		}
	}

	fmt.Printf("Lexer rules marked as terminal: %d/%d\n", lexerTerminalCount, len(parsedGrammar.LexerRules))
	
	if lexerTerminalCount == 0 {
		t.Error("No lexer rules were marked as terminal")
	}

	// Debug: Print first 10 lexer and parser rules to see what's available
	fmt.Println("\nFirst 10 lexer rules:")
	count := 0
	for ruleName := range parsedGrammar.LexerRules {
		if count < 10 {
			fmt.Printf("  %s\n", ruleName)
			count++
		}
	}
	
	fmt.Println("\nFirst 10 parser rules:")
	count = 0
	for ruleName := range parsedGrammar.ParserRules {
		if count < 10 {
			node := depGraph.GetNode(ruleName)
			fmt.Printf("  %s (HasTerminal=%t, TerminalAlts=%v)\n", 
				ruleName, node.HasTerminalAlternatives, node.TerminalAlternativeIndex)
			count++
		}
	}
	
	// Test some specific rules that should exist
	testRules := []string{"selectstmt", "a_expr", "IDENT", "ICONST"}
	
	for _, ruleName := range testRules {
		node := depGraph.GetNode(ruleName)
		if node != nil {
			fmt.Printf("Rule %s: HasTerminalAlternatives=%t, TerminalAlts=%v\n", 
				ruleName, node.HasTerminalAlternatives, node.TerminalAlternativeIndex)
		} else {
			fmt.Printf("Rule %s: Not found in dependency graph\n", ruleName)
		}
	}

	t.Log("Dependency graph construction completed successfully")
}

func TestGrammarValidation(t *testing.T) {
	repoRoot := getRepoRoot()
	
	// PostgreSQL grammar file paths
	lexerPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLLexer.g4")
	parserPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLParser.g4")

	// Parse grammar files
	parsedGrammar, err := grammar.ParseAndMergeGrammarFiles([]string{lexerPath, parserPath})
	if err != nil {
		t.Fatalf("Failed to parse grammar files: %v", err)
	}

	// Validate grammar
	err = parsedGrammar.ValidateGrammar()
	if err != nil {
		t.Errorf("Grammar validation failed: %v", err)
		
		// Print analysis results for debugging
		fmt.Println("\n=== Grammar Analysis Results ===")
		parsedGrammar.PrintDependencyAnalysis()
	} else {
		t.Log("Grammar validation passed - all rules have terminal alternatives")
	}
}

func TestDependencyGraphSpecificRules(t *testing.T) {
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

	// Test specific known patterns
	tests := []struct {
		ruleName         string
		expectTerminal   bool
		description      string
	}{
		{"OPEN_PAREN", true, "Lexer rule should be terminal"},
		{"SELECT", true, "Lexer rule should be terminal"},
		{"selectstmt", true, "Should have at least one terminal alternative"},
		{"a_expr", true, "Expression rule should have terminal alternatives"},
	}

	for _, test := range tests {
		hasTerminal := depGraph.HasTerminalAlternatives(test.ruleName)
		if hasTerminal != test.expectTerminal {
			t.Errorf("Rule %s: expected HasTerminalAlternatives=%t, got %t (%s)", 
				test.ruleName, test.expectTerminal, hasTerminal, test.description)
		}

		if hasTerminal {
			terminalAlts := depGraph.GetTerminalAlternatives(test.ruleName)
			if len(terminalAlts) == 0 {
				t.Errorf("Rule %s: HasTerminalAlternatives=true but no terminal alternatives found", test.ruleName)
			}
			fmt.Printf("✓ Rule %s has %d terminal alternatives: %v\n", test.ruleName, len(terminalAlts), terminalAlts)
		}
	}
}