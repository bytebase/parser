package grammar

import (
	"os"
	"path/filepath"
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