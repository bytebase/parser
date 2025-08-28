package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bytebase/parser/tools/fuzzing/internal/config"
	"github.com/bytebase/parser/tools/fuzzing/internal/generator"
)

func main() {
	cfg := parseFlags()

	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	cfg.Print()

	gen := generator.New(cfg)
	if err := gen.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "Generation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Generation completed successfully!")
}

func parseFlags() *config.Config {
	cfg := &config.Config{}
	var grammarArg string

	flag.StringVar(&grammarArg, "grammar", "", "Grammar file(s): single file or comma-separated lexer,parser files")
	flag.StringVar(&cfg.StartRule, "start-rule", "", "Starting grammar rule name")
	flag.IntVar(&cfg.Count, "count", 10, "Number of queries to generate")
	flag.IntVar(&cfg.MaxDepth, "max-depth", 5, "Maximum recursion depth")
	flag.Float64Var(&cfg.OptionalProb, "optional-prob", 0.5, "Probability of including optional elements (0.0-1.0)")
	flag.IntVar(&cfg.MaxQuantifier, "max-quantifier", 5, "Maximum count for quantified rules (* and +)")
	flag.IntVar(&cfg.MinQuantifier, "min-quantifier", 0, "Minimum count for quantified rules (overrides grammar)")
	flag.IntVar(&cfg.QuantifierCount, "quantifier-count", 0, "Fixed count for all quantifiers (overrides min/max)")
	flag.StringVar(&cfg.Output, "output", "", "Output file path (default: stdout)")
	flag.Int64Var(&cfg.Seed, "seed", time.Now().UnixNano(), "Random seed for reproducible generation")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Grammar-Aware Fuzzing Tool\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Single combined grammar file\n")
		fmt.Fprintf(os.Stderr, "  %s --grammar combined.g4 --start-rule selectStmt --count 10\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Separate lexer and parser files\n")
		fmt.Fprintf(os.Stderr, "  %s --grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 10\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Control recursion and quantifiers\n")
		fmt.Fprintf(os.Stderr, "  %s --grammar cql/CqlLexer.g4,cql/CqlParser.g4 --start-rule expr --max-depth 3 --max-quantifier 8 --count 5\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Performance testing\n")
		fmt.Fprintf(os.Stderr, "  %s --grammar redshift/RedshiftLexer.g4,redshift/RedshiftParser.g4 --start-rule blockStmt --quantifier-count 100 --count 10\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Output to file\n")
		fmt.Fprintf(os.Stderr, "  %s --grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 100 --output queries.sql\n\n", os.Args[0])
	}

	flag.Parse()

	// Parse grammar files from comma-separated argument
	if grammarArg != "" {
		files := strings.Split(grammarArg, ",")
		// Trim whitespace from each file
		for i, file := range files {
			files[i] = strings.TrimSpace(file)
		}
		cfg.GrammarFiles = files
	}

	return cfg
}
