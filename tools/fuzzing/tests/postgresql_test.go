package tests

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/bytebase/parser/tools/fuzzing/internal/config"
	"github.com/bytebase/parser/tools/fuzzing/internal/generator"
)

// getRepoRoot finds the repository root directory
func getRepoRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	// Go up from tools/fuzzing/tests to the repo root
	return filepath.Join(filepath.Dir(filename), "..", "..", "..")
}

func TestPostgreSQLSelectStmt(t *testing.T) {
	repoRoot := getRepoRoot()

	// PostgreSQL grammar file paths
	lexerPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLLexer.g4")
	parserPath := filepath.Join(repoRoot, "postgresql", "PostgreSQLParser.g4")

	tests := []struct {
		name         string
		startRule    string
		count        int
		maxDepth     int
		optionalProb float64
		seed         int64
	}{
		{
			name:         "Simple SELECT statements",
			startRule:    "selectstmt",
			count:        3,
			maxDepth:     10,
			optionalProb: 0.7,
			seed:         42,
		},
		{
			name:         "Deep SELECT statements",
			startRule:    "selectstmt",
			count:        2,
			maxDepth:     8,
			optionalProb: 0.5,
			seed:         123,
		},
		{
			name:         "Minimal SELECT statements",
			startRule:    "selectstmt",
			count:        5,
			maxDepth:     3,
			optionalProb: 0.3,
			seed:         456,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.Config{
				GrammarFiles:    []string{lexerPath, parserPath},
				StartRule:       tt.startRule,
				Count:           tt.count,
				MaxDepth:        tt.maxDepth,
				OptionalProb:    tt.optionalProb,
				MaxQuantifier:   3,
				MinQuantifier:   1,
				QuantifierCount: 0,
				OutputFormat:    config.CompactOutput,
				Seed:            tt.seed,
			}

			fmt.Printf("\n=== %s ===\n", tt.name)
			fmt.Printf("Config: MaxDepth=%d, OptionalProb=%.1f, Count=%d, Seed=%d\n",
				tt.maxDepth, tt.optionalProb, tt.count, tt.seed)
			fmt.Println()

			gen := generator.New(cfg)
			err := gen.Generate()

			if err != nil {
				t.Errorf("Failed to generate PostgreSQL %s: %v", tt.startRule, err)
			} else {
				t.Logf("Successfully generated %d PostgreSQL %s statements", tt.count, tt.startRule)
			}
		})
	}
}
