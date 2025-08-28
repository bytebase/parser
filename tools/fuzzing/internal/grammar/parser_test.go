package grammar

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestCompleteGrammarIR tests the complete intermediate representation of parsed grammar
func TestCompleteGrammarIR(t *testing.T) {
	grammarContent := `
parser grammar CompleteIRTest;

// Simple rule with literals
greeting: 'Hello' 'World';

// Rule with alternatives  
statement: selectStmt | insertStmt | 'DELETE';

// Rule with quantifiers and mixed elements
selectStmt: 'SELECT' columnList 'FROM' IDENTIFIER whereClause?;

// Rule with quantified elements
columnList: column (',' column)*;

// Rule with token reference
column: IDENTIFIER ('AS' IDENTIFIER)?;

// Rule with optional and alternatives
whereClause: 'WHERE' expr;

// Complex rule with multiple alternatives and quantifiers
expr: expr '+' expr
    | expr '*' expr  
    | '(' expr ')'
    | IDENTIFIER
    | NUMBER;
`

	tmpFile := createTempGrammarFile(t, grammarContent)
	defer os.Remove(tmpFile)

	grammar, err := ParseGrammarFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to parse grammar: %v", err)
	}

	// Basic grammar properties
	if grammar == nil {
		t.Fatal("Grammar is nil")
	}
	if grammar.FilePath != tmpFile {
		t.Errorf("Expected file path %s, got %s", tmpFile, grammar.FilePath)
	}
	if len(grammar.LexerRules) != 0 {
		t.Errorf("Expected 0 lexer rules, got %d", len(grammar.LexerRules))
	}
	if len(grammar.ParserRules) != 7 {
		t.Errorf("Expected 7 parser rules, got %d", len(grammar.ParserRules))
	}

	// Test cases for rule validation
	tests := []struct {
		ruleName     string
		alternatives int
		elements     []elementTest
	}{
		{
			ruleName:     "greeting",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "'Hello'", quantifier: NONE, elementType: "literal"},
				{altIndex: 0, elementIndex: 1, value: "'World'", quantifier: NONE, elementType: "literal"},
			},
		},
		{
			ruleName:     "statement",
			alternatives: 3,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "selectStmt", quantifier: NONE, elementType: "reference"},
				{altIndex: 1, elementIndex: 0, value: "insertStmt", quantifier: NONE, elementType: "reference"},
				{altIndex: 2, elementIndex: 0, value: "'DELETE'", quantifier: NONE, elementType: "literal"},
			},
		},
		{
			ruleName:     "selectStmt",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "'SELECT'", quantifier: NONE, elementType: "literal"},
				{altIndex: 0, elementIndex: 1, value: "columnList", quantifier: NONE, elementType: "reference"},
				{altIndex: 0, elementIndex: 2, value: "'FROM'", quantifier: NONE, elementType: "literal"},
				{altIndex: 0, elementIndex: 3, value: "IDENTIFIER", quantifier: NONE, elementType: "reference"},
				{altIndex: 0, elementIndex: 4, value: "whereClause", quantifier: OPTIONAL_Q, elementType: "reference"},
			},
		},
		{
			ruleName:     "columnList",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "column", quantifier: NONE, elementType: "reference"},
				{altIndex: 0, elementIndex: 1, value: "(',' column)", quantifier: ZERO_MORE, elementType: "block"},
			},
		},
		{
			ruleName:     "column",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "IDENTIFIER", quantifier: NONE, elementType: "reference"},
				{altIndex: 0, elementIndex: 1, value: "('AS' IDENTIFIER)", quantifier: OPTIONAL_Q, elementType: "block"},
			},
		},
		{
			ruleName:     "whereClause",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "'WHERE'", quantifier: NONE, elementType: "literal"},
				{altIndex: 0, elementIndex: 1, value: "expr", quantifier: NONE, elementType: "reference"},
			},
		},
		{
			ruleName:     "expr",
			alternatives: 5,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "expr", quantifier: NONE, elementType: "reference"},
				{altIndex: 0, elementIndex: 1, value: "'+'", quantifier: NONE, elementType: "literal"},
				{altIndex: 0, elementIndex: 2, value: "expr", quantifier: NONE, elementType: "reference"},
				{altIndex: 1, elementIndex: 1, value: "'*'", quantifier: NONE, elementType: "literal"},
				{altIndex: 2, elementIndex: 0, value: "'('", quantifier: NONE, elementType: "literal"},
				{altIndex: 2, elementIndex: 1, value: "expr", quantifier: NONE, elementType: "reference"},
				{altIndex: 2, elementIndex: 2, value: "')'", quantifier: NONE, elementType: "literal"},
				{altIndex: 3, elementIndex: 0, value: "IDENTIFIER", quantifier: NONE, elementType: "reference"},
				{altIndex: 4, elementIndex: 0, value: "NUMBER", quantifier: NONE, elementType: "reference"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.ruleName, func(t *testing.T) {
			rule := grammar.GetRule(tc.ruleName)
			if rule == nil {
				t.Fatalf("rule %s not found", tc.ruleName)
			}
			if rule.Name != tc.ruleName || rule.IsLexer {
				t.Errorf("rule %s has incorrect metadata", tc.ruleName)
			}
			if len(rule.Alternatives) != tc.alternatives {
				t.Errorf("%s: expected %d alternatives, got %d", tc.ruleName, tc.alternatives, len(rule.Alternatives))
			}

			for _, elem := range tc.elements {
				altIndex := elem.altIndex
				elementIndex := elem.elementIndex

				if altIndex >= len(rule.Alternatives) {
					t.Errorf("%s: alternative %d out of range", tc.ruleName, altIndex)
					continue
				}

				elements := rule.Alternatives[altIndex].Elements
				if elementIndex >= len(elements) {
					t.Errorf("%s alt %d: element %d out of range", tc.ruleName, altIndex, elementIndex)
					continue
				}

				element := elements[elementIndex]
				if elem.value != "" && element.Value.String() != elem.value {
					t.Errorf("%s alt %d elem %d: expected value %s, got %s", tc.ruleName, altIndex, elementIndex, elem.value, element.Value.String())
				}
				if element.Quantifier != elem.quantifier {
					t.Errorf("%s alt %d elem %d: expected quantifier %v, got %v", tc.ruleName, altIndex, elementIndex, elem.quantifier, element.Quantifier)
				}
				
				// Validate element type using type assertions
				switch elem.elementType {
				case "literal":
					if _, ok := element.Value.(LiteralValue); !ok {
						t.Errorf("%s alt %d elem %d: expected LiteralValue, got %T", tc.ruleName, altIndex, elementIndex, element.Value)
					}
				case "reference":
					if _, ok := element.Value.(ReferenceValue); !ok {
						t.Errorf("%s alt %d elem %d: expected ReferenceValue, got %T", tc.ruleName, altIndex, elementIndex, element.Value)
					}
				case "block":
					if _, ok := element.Value.(BlockValue); !ok {
						t.Errorf("%s alt %d elem %d: expected BlockValue, got %T", tc.ruleName, altIndex, elementIndex, element.Value)
					}
				}
			}
		})
	}

	// Test GetAllRules method
	allRules := grammar.GetAllRules()
	if len(allRules) != 7 {
		t.Errorf("GetAllRules: expected 7 rules, got %d", len(allRules))
	}
}

type elementTest struct {
	altIndex     int
	elementIndex int
	value        string
	quantifier   Quantifier
	elementType  string // "literal", "reference", or "block"
}

// Helper functions

func createTempGrammarFile(t *testing.T, content string) string {
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_grammar.g4")

	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp grammar file: %v", err)
	}

	return tmpFile
}

func createTempGrammarFileWithName(t *testing.T, content string, filename string) string {
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, filename)

	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp grammar file: %v", err)
	}

	return tmpFile
}

// TestLexerRuleParsing tests the parsing of lexer rules
func TestLexerRuleParsing(t *testing.T) {
	grammarContent := `
lexer grammar TestLexer;

// Simple string literal
SELECT: 'SELECT';

// Character range
LETTER: [a-zA-Z];

// Complex rule with alternatives and quantifiers
IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]*;

// Rule with character set
DIGIT: [0-9];

// Rule with wildcard and quantifier
COMMENT: '//' .*? '\n';
`

	tmpFile := createTempGrammarFile(t, grammarContent)
	defer os.Remove(tmpFile)

	grammar, err := ParseGrammarFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to parse lexer grammar: %v", err)
	}

	// Basic grammar properties
	if grammar == nil {
		t.Fatal("Grammar is nil")
	}
	if len(grammar.ParserRules) != 0 {
		t.Errorf("Expected 0 parser rules, got %d", len(grammar.ParserRules))
	}
	if len(grammar.LexerRules) != 5 {
		t.Errorf("Expected 5 lexer rules, got %d", len(grammar.LexerRules))
	}

	// Test cases for lexer rule validation
	tests := []struct {
		ruleName     string
		alternatives int
		elements     []elementTest
	}{
		{
			ruleName:     "SELECT",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "'SELECT'", quantifier: NONE, elementType: "literal"},
			},
		},
		{
			ruleName:     "LETTER",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "[a-zA-Z]", quantifier: NONE, elementType: "literal"},
			},
		},
		{
			ruleName:     "IDENTIFIER",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "[a-zA-Z_]", quantifier: NONE, elementType: "literal"},
				{altIndex: 0, elementIndex: 1, value: "[a-zA-Z0-9_]", quantifier: ZERO_MORE, elementType: "literal"},
			},
		},
		{
			ruleName:     "DIGIT",
			alternatives: 1,
			elements: []elementTest{
				{altIndex: 0, elementIndex: 0, value: "[0-9]", quantifier: NONE, elementType: "literal"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.ruleName, func(t *testing.T) {
			rule := grammar.GetRule(tc.ruleName)
			if rule == nil {
				t.Fatalf("rule %s not found", tc.ruleName)
			}
			if rule.Name != tc.ruleName || !rule.IsLexer {
				t.Errorf("rule %s has incorrect metadata: IsLexer=%v", tc.ruleName, rule.IsLexer)
			}
			if len(rule.Alternatives) != tc.alternatives {
				t.Errorf("%s: expected %d alternatives, got %d", tc.ruleName, tc.alternatives, len(rule.Alternatives))
			}

			for _, elem := range tc.elements {
				altIndex := elem.altIndex
				elementIndex := elem.elementIndex

				if altIndex >= len(rule.Alternatives) {
					t.Errorf("%s: alternative %d out of range", tc.ruleName, altIndex)
					continue
				}

				elements := rule.Alternatives[altIndex].Elements
				if elementIndex >= len(elements) {
					t.Errorf("%s alt %d: element %d out of range", tc.ruleName, altIndex, elementIndex)
					continue
				}

				element := elements[elementIndex]
				if elem.value != "" && element.Value.String() != elem.value {
					t.Errorf("%s alt %d elem %d: expected value %s, got %s", tc.ruleName, altIndex, elementIndex, elem.value, element.Value.String())
				}
				if element.Quantifier != elem.quantifier {
					t.Errorf("%s alt %d elem %d: expected quantifier %v, got %v", tc.ruleName, altIndex, elementIndex, elem.quantifier, element.Quantifier)
				}
				
				// Validate element type using type assertions
				switch elem.elementType {
				case "literal":
					if _, ok := element.Value.(LiteralValue); !ok {
						t.Errorf("%s alt %d elem %d: expected LiteralValue, got %T", tc.ruleName, altIndex, elementIndex, element.Value)
					}
				case "reference":
					if _, ok := element.Value.(ReferenceValue); !ok {
						t.Errorf("%s alt %d elem %d: expected ReferenceValue, got %T", tc.ruleName, altIndex, elementIndex, element.Value)
					}
				case "block":
					if _, ok := element.Value.(BlockValue); !ok {
						t.Errorf("%s alt %d elem %d: expected BlockValue, got %T", tc.ruleName, altIndex, elementIndex, element.Value)
					}
				}
			}
		})
	}
}

// TestCombinedGrammarParsing tests parsing of combined grammar with both parser and lexer rules
func TestCombinedGrammarParsing(t *testing.T) {
	grammarContent := `
grammar CombinedTest;

// Parser rules
statement: selectStmt;
selectStmt: 'SELECT' IDENTIFIER;

// Lexer rules
IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
`

	tmpFile := createTempGrammarFile(t, grammarContent)
	defer os.Remove(tmpFile)

	grammar, err := ParseGrammarFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to parse combined grammar: %v", err)
	}

	// Basic grammar properties
	if grammar == nil {
		t.Fatal("Grammar is nil")
	}
	if len(grammar.ParserRules) != 2 {
		t.Errorf("Expected 2 parser rules, got %d", len(grammar.ParserRules))
	}
	if len(grammar.LexerRules) != 2 {
		t.Errorf("Expected 2 lexer rules, got %d", len(grammar.LexerRules))
	}

	// Test parser rule
	statement := grammar.GetRule("statement")
	if statement == nil {
		t.Fatal("Parser rule 'statement' not found")
	}
	if statement.IsLexer {
		t.Error("Parser rule incorrectly marked as lexer rule")
	}

	// Test lexer rule  
	identifier := grammar.GetRule("IDENTIFIER")
	if identifier == nil {
		t.Fatal("Lexer rule 'IDENTIFIER' not found")
	}
	if !identifier.IsLexer {
		t.Error("Lexer rule incorrectly marked as parser rule")
	}

	// Test that GetAllRules returns both types
	allRules := grammar.GetAllRules()
	if len(allRules) != 4 {
		t.Errorf("Expected 4 total rules, got %d", len(allRules))
	}
}

// TestGrammarMerging tests merging multiple grammar files
func TestGrammarMerging(t *testing.T) {
	// Create first grammar file (parser rules)
	parserGrammarContent := `
parser grammar ParserTest;

options {
    tokenVocab = LexerTest;
}

statement: selectStmt;
selectStmt: 'SELECT' IDENTIFIER;
`
	
	// Create second grammar file (lexer rules)
	lexerGrammarContent := `
lexer grammar LexerTest;

IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
`

	// Create temporary files with unique names
	tmpParserFile := createTempGrammarFileWithName(t, parserGrammarContent, "test_parser.g4")
	defer os.Remove(tmpParserFile)
	
	tmpLexerFile := createTempGrammarFileWithName(t, lexerGrammarContent, "test_lexer.g4")
	defer os.Remove(tmpLexerFile)

	// Test parsing and merging
	filePaths := []string{tmpParserFile, tmpLexerFile}
	mergedGrammar, err := ParseAndMergeGrammarFiles(filePaths)
	if err != nil {
		t.Fatalf("Failed to parse and merge grammar files: %v", err)
	}

	// Verify merged grammar properties
	if mergedGrammar == nil {
		t.Fatal("Merged grammar is nil")
	}
	
	if len(mergedGrammar.ParserRules) != 2 {
		t.Errorf("Expected 2 parser rules, got %d", len(mergedGrammar.ParserRules))
	}
	
	if len(mergedGrammar.LexerRules) != 2 {
		t.Errorf("Expected 2 lexer rules, got %d", len(mergedGrammar.LexerRules))
	}

	// Test that both parser and lexer rules are accessible
	statement := mergedGrammar.GetRule("statement")
	if statement == nil || statement.IsLexer {
		t.Error("Parser rule 'statement' not found or incorrectly marked")
	}

	identifier := mergedGrammar.GetRule("IDENTIFIER")
	if identifier == nil || !identifier.IsLexer {
		t.Error("Lexer rule 'IDENTIFIER' not found or incorrectly marked")
	}

	// Test that merged path is updated
	if !strings.Contains(mergedGrammar.FilePath, "+") {
		t.Errorf("Expected merged file path to contain '+', got: %s", mergedGrammar.FilePath)
	}

	// Test GetAllRules on merged grammar
	allRules := mergedGrammar.GetAllRules()
	if len(allRules) != 4 {
		t.Errorf("Expected 4 total rules in merged grammar, got %d", len(allRules))
	}
}

// TestGrammarMergingWithConflicts tests handling of duplicate rule names
func TestGrammarMergingWithConflicts(t *testing.T) {
	// Create two grammars with conflicting rule names
	grammar1Content := `
lexer grammar Test1;
IDENTIFIER: [a-z]+;
`
	
	grammar2Content := `
lexer grammar Test2;
IDENTIFIER: [A-Z]+;  // Conflict with first grammar
`

	tmpFile1 := createTempGrammarFileWithName(t, grammar1Content, "conflict1.g4")
	defer os.Remove(tmpFile1)
	
	tmpFile2 := createTempGrammarFileWithName(t, grammar2Content, "conflict2.g4")
	defer os.Remove(tmpFile2)

	// Test that merging fails with duplicate rule names
	filePaths := []string{tmpFile1, tmpFile2}
	_, err := ParseAndMergeGrammarFiles(filePaths)
	if err == nil {
		t.Error("Expected error when merging grammars with duplicate rule names")
	}

	if !strings.Contains(err.Error(), "duplicate") {
		t.Errorf("Expected error about duplicate rules, got: %v", err)
	}
}

// TestParseAndMergeGrammarFilesEdgeCases tests edge cases
func TestParseAndMergeGrammarFilesEdgeCases(t *testing.T) {
	// Test with empty file list
	_, err := ParseAndMergeGrammarFiles([]string{})
	if err == nil {
		t.Error("Expected error with empty file list")
	}

	// Test with single file
	grammarContent := `
lexer grammar SingleTest;
TOKEN: 'test';
`
	
	tmpFile := createTempGrammarFileWithName(t, grammarContent, "single.g4")
	defer os.Remove(tmpFile)

	grammar, err := ParseAndMergeGrammarFiles([]string{tmpFile})
	if err != nil {
		t.Fatalf("Failed to parse single grammar file: %v", err)
	}

	if len(grammar.LexerRules) != 1 {
		t.Errorf("Expected 1 lexer rule, got %d", len(grammar.LexerRules))
	}

	if grammar.GetRule("TOKEN") == nil {
		t.Error("TOKEN rule not found in single file grammar")
	}
}