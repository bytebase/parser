package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
)

func TestPureLeftRecursiveGrammarNonTerminal(t *testing.T) {
	// Create a grammar with ONLY left-recursive rules (no terminal alternatives)
	tempDir := t.TempDir()
	grammarContent := `grammar PureLeftRecursive;

// Parser rules  
root: expr EOF;

// This rule has NO terminal alternatives - pure left recursion
expr: expr '+' expr
    | expr '*' expr
    ;

// Lexer rules
PLUS: '+';
MULTIPLY: '*';
WS: [ \t\r\n]+ -> skip;
EOF: '<EOF>';
`

	grammarFile := filepath.Join(tempDir, "PureLeftRecursive.g4")
	err := os.WriteFile(grammarFile, []byte(grammarContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create pure left-recursive grammar file: %v", err)
	}

	// Parse the grammar
	parsedGrammar, err := grammar.ParseAndMergeGrammarFiles([]string{grammarFile})
	if err != nil {
		t.Fatalf("Failed to parse pure left-recursive grammar: %v", err)
	}

	depGraph := parsedGrammar.GetDependencyGraph()

	// Check that the expr rule is NOT terminal
	exprNode := depGraph.GetNode("expr")
	if exprNode == nil {
		t.Fatal("expr rule not found in dependency graph")
	}

	t.Logf("=== Pure Left-Recursive Grammar Analysis ===")
	t.Logf("expr rule has %d alternatives", len(exprNode.Alternatives))
	t.Logf("expr HasTerminalAlternatives: %t", exprNode.HasTerminalAlternatives)
	t.Logf("expr TerminalAlternativeIndex: %v", exprNode.TerminalAlternativeIndex)

	// This rule should NOT have terminal alternatives because all alternatives
	// are left-recursive and there's no base case
	if exprNode.HasTerminalAlternatives {
		t.Errorf("Expected pure left-recursive expr rule to NOT have terminal alternatives, but it does")
	}

	// Validate the grammar should fail
	err = depGraph.ValidateGrammar()
	if err == nil {
		t.Errorf("Expected grammar validation to fail for pure left-recursive grammar, but it passed")
	} else {
		t.Logf("Grammar validation correctly failed: %v", err)
	}

	// Check that root is also non-terminal because it depends on non-terminal expr
	rootNode := depGraph.GetNode("root")
	if rootNode == nil {
		t.Fatal("root rule not found in dependency graph")
	}

	if rootNode.HasTerminalAlternatives {
		t.Errorf("Expected root rule to be non-terminal due to non-terminal expr dependency, but it's terminal")
	}

	t.Logf("root HasTerminalAlternatives: %t (expected false)", rootNode.HasTerminalAlternatives)
}
