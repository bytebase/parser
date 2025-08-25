package grammar

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

// CustomErrorListener captures parsing errors like redshift parser_test.go
type CustomErrorListener struct {
	errors int
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors += 1
	antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (l *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

// TestAllGrammarFiles tests parsing of all .g4 files in the repository
func TestAllGrammarFiles(t *testing.T) {
	// 1. Calculate all the g4 files, record their paths
	repoRoot := filepath.Join("..", "..")
	grammarFiles, err := findGrammarFiles(repoRoot)
	if err != nil {
		t.Fatalf("Failed to find grammar files: %v", err)
	}

	if len(grammarFiles) == 0 {
		t.Fatalf("No .g4 grammar files found in repository")
	}

	t.Logf("Found %d grammar files to test", len(grammarFiles))

	// 2. Parse them one by one using t.Run() 
	for _, grammarFile := range grammarFiles {
		relPath, _ := filepath.Rel(repoRoot, grammarFile)
		
		t.Run(relPath, func(t *testing.T) {
			// Parse the grammar file
			err := parseAndValidateGrammarFile(grammarFile)
			if err != nil {
				t.Errorf("Failed to parse grammar file: %v", err)
			}
		})
	}
}

// findGrammarFiles recursively finds all .g4 files
func findGrammarFiles(root string) ([]string, error) {
	var files []string
	
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if filepath.Ext(path) == ".g4" {
			files = append(files, path)
		}
		
		return nil
	})
	
	return files, err
}

// parseAndValidateGrammarFile parses a single grammar file and returns error if it fails
func parseAndValidateGrammarFile(filePath string) error {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(content) == 0 {
		return fmt.Errorf("grammar file is empty")
	}

	// Create input stream
	input := antlr.NewInputStream(string(content))
	
	// Create lexer
	lexer := NewANTLRv4Lexer(input)
	
	// Add custom error listener to lexer
	lexerErrorListener := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrorListener)
	
	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, 0)
	
	// Create parser
	parser := NewANTLRv4Parser(stream)
	
	// Add custom error listener to parser (using same pattern as redshift)
	parserErrorListener := &CustomErrorListener{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrorListener)
	
	// Parse the grammar
	tree := parser.GrammarSpec()
	
	// Check for lexer errors
	if lexerErrorListener.errors > 0 {
		return fmt.Errorf("lexer found %d errors", lexerErrorListener.errors)
	}
	
	// Check for parser errors
	if parserErrorListener.errors > 0 {
		return fmt.Errorf("parser found %d errors", parserErrorListener.errors)
	}
	
	// Check tree is not nil
	if tree == nil {
		return fmt.Errorf("parser returned nil tree")
	}
	
	return nil
}