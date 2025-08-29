package config

import (
	"fmt"

	"github.com/pkg/errors"
)

// OutputFormat represents different output formatting options
type OutputFormat int

const (
	// CompactOutput shows cleaner, more readable output (default)
	CompactOutput OutputFormat = iota
	// VerboseOutput shows full grammar rule traversal with comments
	VerboseOutput
)

// ParseOutputFormat parses a string into an OutputFormat
func ParseOutputFormat(s string) OutputFormat {
	switch s {
	case "compact", "":
		return CompactOutput
	case "verbose":
		return VerboseOutput
	default:
		return CompactOutput
	}
}

// Config holds all configuration options for the fuzzer
type Config struct {
	GrammarFiles    []string // Can be one file (combined) or two files (lexer,parser)
	StartRule       string
	Count           int
	MaxDepth        int
	OptionalProb    float64
	MaxQuantifier   int
	MinQuantifier   int
	QuantifierCount int
	Output          string
	OutputFormat    OutputFormat // How to format the output
	Seed            int64
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if len(c.GrammarFiles) == 0 {
		return errors.New("--grammar is required")
	}

	if len(c.GrammarFiles) > 2 {
		return errors.New("--grammar accepts maximum 2 files (lexer,parser)")
	}

	if c.StartRule == "" {
		return errors.New("--start-rule is required")
	}

	if c.Count <= 0 {
		return errors.New("--count must be positive")
	}

	if c.MaxDepth <= 0 {
		return errors.New("--max-depth must be positive")
	}

	if c.OptionalProb < 0.0 || c.OptionalProb > 1.0 {
		return errors.New("--optional-prob must be between 0.0 and 1.0")
	}

	if c.MaxQuantifier <= 0 {
		return errors.New("--max-quantifier must be positive")
	}

	if c.MinQuantifier < 0 {
		return errors.New("--min-quantifier must be non-negative")
	}

	if c.MinQuantifier > c.MaxQuantifier {
		return errors.New("--min-quantifier cannot be greater than --max-quantifier")
	}

	if c.QuantifierCount < 0 {
		return errors.New("--quantifier-count must be non-negative")
	}

	return nil
}

// Print displays the configuration
func (c *Config) Print() {
	fmt.Printf("Grammar-Aware Fuzzer\n")
	if len(c.GrammarFiles) == 1 {
		fmt.Printf("Grammar File: %s\n", c.GrammarFiles[0])
	} else if len(c.GrammarFiles) == 2 {
		fmt.Printf("Lexer File: %s\n", c.GrammarFiles[0])
		fmt.Printf("Parser File: %s\n", c.GrammarFiles[1])
	}
	fmt.Printf("Start Rule: %s\n", c.StartRule)
	fmt.Printf("Count: %d\n", c.Count)
	fmt.Printf("Max Depth: %d\n", c.MaxDepth)
	fmt.Printf("Optional Probability: %.2f\n", c.OptionalProb)
	fmt.Printf("Max Quantifier: %d\n", c.MaxQuantifier)
	if c.MinQuantifier > 0 {
		fmt.Printf("Min Quantifier: %d\n", c.MinQuantifier)
	}
	if c.QuantifierCount > 0 {
		fmt.Printf("Fixed Quantifier Count: %d\n", c.QuantifierCount)
	}
	if c.Output != "" {
		fmt.Printf("Output: %s\n", c.Output)
	}
	fmt.Printf("Seed: %d\n", c.Seed)
	fmt.Println()
}