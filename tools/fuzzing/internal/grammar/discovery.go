package grammar

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// GrammarFiles represents a pair of lexer and parser grammar files
type GrammarFiles struct {
	LexerFile  string
	ParserFile string
	Directory  string
}

// DiscoverGrammarFiles finds lexer and parser files for a given grammar name
func DiscoverGrammarFiles(grammarName string) (*GrammarFiles, error) {
	// Start from fuzzing directory, go up to parser root
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current directory")
	}
	
	// Navigate to parser root (assuming we're in tools/fuzzing)
	repoRoot := filepath.Join(currentDir, "..", "..")
	
	// Try different grammar directory patterns
	grammarDirs := []string{
		filepath.Join(repoRoot, grammarName),                    // Direct: postgresql/, cql/
		filepath.Join(repoRoot, "tools", "grammar"),             // ANTLR v4 self-grammar
		filepath.Join(repoRoot, "grammars", grammarName),        // Alternative structure
	}
	
	for _, dir := range grammarDirs {
		if files, err := findGrammarFilesInDir(dir, grammarName); err == nil {
			return files, nil
		}
	}
	
	return nil, errors.Errorf("grammar '%s' not found in any of the expected locations", grammarName)
}

// findGrammarFilesInDir searches for grammar files in a specific directory
func findGrammarFilesInDir(dir, grammarName string) (*GrammarFiles, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, errors.Errorf("directory does not exist: %s", dir)
	}
	
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read directory %s", dir)
	}
	
	var lexerFile, parserFile string
	
	// Look for grammar files using different naming patterns
	patterns := []struct {
		lexerPattern  string
		parserPattern string
	}{
		// Standard patterns: PostgreSQLLexer.g4, PostgreSQLParser.g4  
		{fmt.Sprintf("%sLexer.g4", capitalize(grammarName)), fmt.Sprintf("%sParser.g4", capitalize(grammarName))},
		// Special case for PostgreSQL: postgresql -> PostgreSQL
		{fmt.Sprintf("%sLexer.g4", strings.ToUpper(grammarName)), fmt.Sprintf("%sParser.g4", strings.ToUpper(grammarName))},
		// Alternate patterns: CqlLexer.g4, CqlParser.g4  
		{fmt.Sprintf("%sLexer.g4", strings.Title(grammarName)), fmt.Sprintf("%sParser.g4", strings.Title(grammarName))},
		// Lowercase patterns: postgresql_lexer.g4, postgresql_parser.g4
		{fmt.Sprintf("%s_lexer.g4", strings.ToLower(grammarName)), fmt.Sprintf("%s_parser.g4", strings.ToLower(grammarName))},
	}
	
	// Special cases for known grammar naming conventions
	switch strings.ToLower(grammarName) {
	case "postgresql":
		patterns = append(patterns, struct {
			lexerPattern  string
			parserPattern string
		}{"PostgreSQLLexer.g4", "PostgreSQLParser.g4"})
	case "antlrv4":
		patterns = append(patterns, struct {
			lexerPattern  string
			parserPattern string
		}{"ANTLRv4Lexer.g4", "ANTLRv4Parser.g4"})
	}
	
	// Special case for ANTLR v4 self-grammar directory  
	if strings.Contains(dir, "tools/grammar") {
		patterns = append(patterns, struct {
			lexerPattern  string
			parserPattern string
		}{"ANTLRv4Lexer.g4", "ANTLRv4Parser.g4"})
	}
	
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".g4") {
			for _, pattern := range patterns {
				if entry.Name() == pattern.lexerPattern {
					lexerFile = filepath.Join(dir, entry.Name())
				}
				if entry.Name() == pattern.parserPattern {
					parserFile = filepath.Join(dir, entry.Name())
				}
			}
		}
	}
	
	// Check if we found both files
	if lexerFile == "" {
		return nil, errors.Errorf("lexer file not found in %s", dir)
	}
	if parserFile == "" {
		return nil, errors.Errorf("parser file not found in %s", dir)
	}
	
	return &GrammarFiles{
		LexerFile:  lexerFile,
		ParserFile: parserFile,
		Directory:  dir,
	}, nil
}

// ListAvailableGrammars scans for all available grammar directories
func ListAvailableGrammars() ([]string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current directory")
	}
	
	repoRoot := filepath.Join(currentDir, "..", "..")
	
	var grammars []string
	
	// Scan for grammar directories
	entries, err := os.ReadDir(repoRoot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read repository root")
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			dirPath := filepath.Join(repoRoot, entry.Name())
			if hasGrammarFiles(dirPath) {
				grammars = append(grammars, entry.Name())
			}
		}
	}
	
	// Add special case for ANTLR v4 self-grammar
	if hasGrammarFiles(filepath.Join(repoRoot, "tools", "grammar")) {
		grammars = append(grammars, "antlrv4")
	}
	
	return grammars, nil
}

// hasGrammarFiles checks if a directory contains .g4 files
func hasGrammarFiles(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	
	var hasLexer, hasParser bool
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".g4") {
			name := strings.ToLower(entry.Name())
			if strings.Contains(name, "lexer") {
				hasLexer = true
			}
			if strings.Contains(name, "parser") {
				hasParser = true
			}
		}
	}
	
	return hasLexer && hasParser
}

// capitalize capitalizes the first letter of a string, preserving the rest
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}