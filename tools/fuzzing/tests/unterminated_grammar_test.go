package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
)

func TestUnterminatedGrammarErrorReporting(t *testing.T) {
	// Create two grammar files that when merged create unterminated rules
	lexerContent := `
lexer grammar TestLexer;

PLUS: '+' ;
NUMBER: [0-9]+ ;
`

	parserContent := `
parser grammar TestParser;

// Import tokens from lexer
options { tokenVocab=TestLexer; }

// This creates infinite left recursion with no terminal alternatives
expr: expr PLUS expr ;
`
	
	// Write the grammars to temporary files
	tmpLexer := "/tmp/test_lexer.g4"
	tmpParser := "/tmp/test_parser.g4"
	
	err := os.WriteFile(tmpLexer, []byte(lexerContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write lexer grammar: %v", err)
	}
	defer os.Remove(tmpLexer)
	
	err = os.WriteFile(tmpParser, []byte(parserContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write parser grammar: %v", err)
	}
	defer os.Remove(tmpParser)
	
	// Try to parse and merge the grammars - this should fail with terminal reachability error
	_, err = grammar.ParseAndMergeGrammarFiles([]string{tmpLexer, tmpParser})
	
	// Verify that we get the expected error
	if err == nil {
		t.Fatal("Expected error for unterminated grammar, but got none")
	}
	
	if !strings.Contains(err.Error(), "without terminal alternatives") {
		t.Errorf("Expected error about terminal alternatives, got: %v", err)
	}
	
	if !strings.Contains(err.Error(), "expr") {
		t.Errorf("Expected error to mention 'expr' rule, got: %v", err)
	}
	
	t.Logf("✅ Correctly detected unterminated grammar: %v", err)
}

func TestValidSimpleGrammar(t *testing.T) {
	// Create a simple grammar that should work
	grammarContent := `
grammar TestGrammar;

// This has terminal alternatives
expr: expr '+' expr | NUMBER ;

// Lexer rules
PLUS: '+' ;
NUMBER: [0-9]+ ;
`
	
	// Write the grammar to a temporary file
	tmpFile := "/tmp/test_valid.g4"
	err := os.WriteFile(tmpFile, []byte(grammarContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test grammar: %v", err)
	}
	defer os.Remove(tmpFile)
	
	// Try to parse the grammar - this should succeed
	parsedGrammar, err := grammar.ParseAndMergeGrammarFiles([]string{tmpFile})
	if err != nil {
		t.Fatalf("Expected valid grammar to parse successfully, got error: %v", err)
	}
	
	// Verify the grammar was parsed correctly
	if parsedGrammar == nil {
		t.Fatal("Expected parsed grammar, got nil")
	}
	
	// Check that expr rule exists and has terminal alternatives
	depGraph := parsedGrammar.GetDependencyGraph()
	if !depGraph.HasTerminalAlternatives("expr") {
		t.Error("Expected expr rule to have terminal alternatives")
	}
	
	t.Log("✅ Valid grammar parsed successfully with terminal alternatives")
}