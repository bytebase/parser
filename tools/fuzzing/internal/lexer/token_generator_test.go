package lexer

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
)

// TestTokenGeneratorBasic tests basic token generation functionality
func TestTokenGeneratorBasic(t *testing.T) {
	config := &TokenGeneratorConfig{
		MaxQuantifierCount:  3,
		MinQuantifierCount:  1,
		OptionalProbability: 1.0, // Always include optional elements for testing
		MaxDepth:           5,
	}
	generator := NewTokenGenerator(12345, config)

	tests := []struct {
		ruleName    string
		grammarText string
		validator   func(string) bool
		description string
	}{
		{
			ruleName:    "SELECT",
			grammarText: "SELECT: 'SELECT';",
			validator:   func(s string) bool { return s == "SELECT" },
			description: "simple string literal",
		},
		{
			ruleName:    "LETTER",
			grammarText: "LETTER: [a-z];",
			validator:   func(s string) bool { return len(s) == 1 && s[0] >= 'a' && s[0] <= 'z' },
			description: "single character range",
		},
		{
			ruleName:    "DIGIT",
			grammarText: "DIGIT: [0-9];",
			validator:   func(s string) bool { return len(s) == 1 && s[0] >= '0' && s[0] <= '9' },
			description: "digit character range",
		},
		{
			ruleName:    "IDENTIFIER",
			grammarText: "IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]*;",
			validator: func(s string) bool {
				if len(s) == 0 {
					return false
				}
				// First character must be letter or underscore
				first := s[0]
				if !((first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z') || first == '_') {
					return false
				}
				// Rest must be letters, digits, or underscore
				for _, c := range s[1:] {
					if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
						return false
					}
				}
				return true
			},
			description: "identifier with quantifier",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Create a temporary grammar file
			grammarContent := "lexer grammar Test;\n\n" + tt.grammarText
			tmpFile := createTempGrammarFile(t, grammarContent)
			defer os.Remove(tmpFile)

			// Parse the grammar
			parsedGrammar, err := grammar.ParseGrammarFile(tmpFile)
			if err != nil {
				t.Fatalf("Failed to parse grammar: %v", err)
			}

			// Get the rule
			rule := parsedGrammar.GetRule(tt.ruleName)
			if rule == nil {
				t.Fatalf("Rule %s not found", tt.ruleName)
			}

			// Generate multiple tokens to test consistency
			for i := 0; i < 10; i++ {
				token, err := generator.GenerateToken(rule)
				if err != nil {
					t.Errorf("Failed to generate token: %v", err)
					continue
				}

				if !tt.validator(token) {
					t.Errorf("Generated token '%s' does not match expected pattern for %s", token, tt.description)
				}
			}
		})
	}
}

// TestQuantifierHandling tests EBNF quantifier handling
func TestQuantifierHandling(t *testing.T) {
	config := &TokenGeneratorConfig{
		MaxQuantifierCount:  5,
		MinQuantifierCount:  2,
		OptionalProbability: 0.5,
		MaxDepth:           5,
	}
	generator := NewTokenGenerator(54321, config)

	tests := []struct {
		ruleName    string
		grammarText string
		validator   func(string) bool
		description string
	}{
		{
			ruleName:    "OPTIONAL",
			grammarText: "OPTIONAL: 'A' 'B'?;",
			validator: func(s string) bool {
				return s == "A" || s == "AB"
			},
			description: "optional element with ?",
		},
		{
			ruleName:    "ZERO_MORE",
			grammarText: "ZERO_MORE: 'X' 'Y'*;",
			validator: func(s string) bool {
				if !strings.HasPrefix(s, "X") {
					return false
				}
				rest := s[1:]
				for _, c := range rest {
					if c != 'Y' {
						return false
					}
				}
				return true
			},
			description: "zero or more with *",
		},
		{
			ruleName:    "ONE_MORE",
			grammarText: "ONE_MORE: 'Z' 'W'+;",
			validator: func(s string) bool {
				if !strings.HasPrefix(s, "Z") {
					return false
				}
				rest := s[1:]
				if len(rest) == 0 {
					return false // + requires at least one
				}
				for _, c := range rest {
					if c != 'W' {
						return false
					}
				}
				return true
			},
			description: "one or more with +",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Create a temporary grammar file
			grammarContent := "lexer grammar Test;\n\n" + tt.grammarText
			tmpFile := createTempGrammarFile(t, grammarContent)
			defer os.Remove(tmpFile)

			// Parse the grammar
			parsedGrammar, err := grammar.ParseGrammarFile(tmpFile)
			if err != nil {
				t.Fatalf("Failed to parse grammar: %v", err)
			}

			// Get the rule
			rule := parsedGrammar.GetRule(tt.ruleName)
			if rule == nil {
				t.Fatalf("Rule %s not found", tt.ruleName)
			}

			// Generate multiple tokens to test quantifier behavior
			validCount := 0
			for i := 0; i < 20; i++ {
				token, err := generator.GenerateToken(rule)
				if err != nil {
					t.Errorf("Failed to generate token: %v", err)
					continue
				}

				if tt.validator(token) {
					validCount++
				} else {
					t.Logf("Generated token '%s' for %s (validation failed but continuing)", token, tt.description)
				}
			}

			// At least 50% of generated tokens should be valid
			if validCount < 10 {
				t.Errorf("Too few valid tokens generated (%d/20) for %s", validCount, tt.description)
			}
		})
	}
}

// TestCharacterSetExpansion tests character set expansion functionality
func TestCharacterSetExpansion(t *testing.T) {
	generator := NewTokenGenerator(9999, nil)

	tests := []struct {
		charset  string
		expected []rune
	}{
		{"abc", []rune{'a', 'b', 'c'}},
		{"a-c", []rune{'a', 'b', 'c'}},
		{"0-2", []rune{'0', '1', '2'}},
		{"a-cX", []rune{'a', 'b', 'c', 'X'}},
		{"A-Z_", append(makeRange('A', 'Z'), '_')},
	}

	for _, tt := range tests {
		t.Run(tt.charset, func(t *testing.T) {
			result, err := generator.expandCharacterSet(tt.charset)
			if err != nil {
				t.Fatalf("Failed to expand character set '%s': %v", tt.charset, err)
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d characters, got %d", len(tt.expected), len(result))
				return
			}

			for i, expected := range tt.expected {
				if result[i] != expected {
					t.Errorf("At position %d: expected '%c', got '%c'", i, expected, result[i])
				}
			}
		})
	}
}

// TestComplexLexerRules tests complex lexer rules with multiple elements
func TestComplexLexerRules(t *testing.T) {
	config := &TokenGeneratorConfig{
		MaxQuantifierCount:  3,
		MinQuantifierCount:  1,
		OptionalProbability: 0.8,
		MaxDepth:           10,
	}
	generator := NewTokenGenerator(11111, config)

	grammarContent := `
lexer grammar ComplexTest;

// Complex identifier rule
IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]*;

// Number with optional decimal part
NUMBER: [0-9]+ ('.' [0-9]+)?;

// String with escaped quotes  
STRING: '"' (~'"')* '"';

// Comment line
COMMENT: '//' (~[\r\n])*;
`

	tmpFile := createTempGrammarFile(t, grammarContent)
	defer os.Remove(tmpFile)

	parsedGrammar, err := grammar.ParseGrammarFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to parse complex grammar: %v", err)
	}

	tests := []struct {
		ruleName string
		pattern  string
	}{
		{"IDENTIFIER", `^[a-zA-Z_][a-zA-Z0-9_]*$`},
		{"NUMBER", `^[0-9]+(\.[0-9]+)?$`},
		{"STRING", `^"[^"]*"$`},
		{"COMMENT", `^//.*$`},
	}

	for _, tt := range tests {
		t.Run(tt.ruleName, func(t *testing.T) {
			rule := parsedGrammar.GetRule(tt.ruleName)
			if rule == nil {
				t.Fatalf("Rule %s not found", tt.ruleName)
			}

			regex := regexp.MustCompile(tt.pattern)
			validCount := 0

			for i := 0; i < 10; i++ {
				token, err := generator.GenerateToken(rule)
				if err != nil {
					t.Errorf("Failed to generate token for %s: %v", tt.ruleName, err)
					continue
				}

				t.Logf("Generated token for %s: '%s'", tt.ruleName, token)

				if regex.MatchString(token) {
					validCount++
				}
			}

			// Expect at least some valid tokens
			if validCount == 0 {
				t.Errorf("No valid tokens generated for %s", tt.ruleName)
			}
		})
	}
}

// Helper functions

func createTempGrammarFile(t *testing.T, content string) string {
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_lexer.g4")

	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp grammar file: %v", err)
	}

	return tmpFile
}

func makeRange(start, end rune) []rune {
	var result []rune
	for c := start; c <= end; c++ {
		result = append(result, c)
	}
	return result
}