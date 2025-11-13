package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

// Keyword represents a PostgreSQL keyword definition from kwlist.h
type Keyword struct {
	Name     string
	Token    string
	Category string
	Label    string
}

const (
	CategoryReserved     = "RESERVED_KEYWORD"
	CategoryUnreserved   = "UNRESERVED_KEYWORD"
	CategoryColName      = "COL_NAME_KEYWORD"
	CategoryTypeFuncName = "TYPE_FUNC_NAME_KEYWORD"

	// PostgreSQL version/branch to fetch keywords from
	PostgreSQLVersion   = "REL_18_STABLE"
	PostgreSQLKwlistURL = "https://raw.githubusercontent.com/postgres/postgres/" + PostgreSQLVersion + "/src/include/parser/kwlist.h"
)

func main() {
	outputDir := flag.String("output", "../", "Output directory for generated grammar files (relative to keyword-generator)")
	flag.Parse()

	fmt.Printf("PostgreSQL Keyword Generator\n")
	fmt.Printf("============================\n\n")
	fmt.Printf("Fetching keyword definitions from PostgreSQL %s...\n", PostgreSQLVersion)
	fmt.Printf("Source: %s\n\n", PostgreSQLKwlistURL)

	keywords, err := fetchAndParseKeywords()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Parsed %d keywords\n\n", len(keywords))

	// Categorize keywords
	categorized := categorizeKeywords(keywords)

	fmt.Printf("Keyword Statistics:\n")
	fmt.Printf("  Reserved keywords:          %3d\n", len(categorized[CategoryReserved]))
	fmt.Printf("  Unreserved keywords:        %3d\n", len(categorized[CategoryUnreserved]))
	fmt.Printf("  Column name keywords:       %3d\n", len(categorized[CategoryColName]))
	fmt.Printf("  Type/function keywords:     %3d\n", len(categorized[CategoryTypeFuncName]))
	fmt.Printf("  ───────────────────────────────\n")
	fmt.Printf("  Total:                      %3d\n\n", len(keywords))

	// Generate ANTLR parser grammar fragment
	parserOutputPath := path.Join(*outputDir, "PostgreSQLKeywords.g4")
	err = generateANTLRGrammar(categorized, parserOutputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating parser grammar: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ Generated: %s\n", parserOutputPath)

	// Update PostgreSQLLexer.g4 with keyword definitions
	lexerPath := fmt.Sprintf("%s/PostgreSQLLexer.g4", *outputDir)
	err = updateLexerWithKeywords(keywords, lexerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating lexer with keywords: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ Updated: %s\n", lexerPath)

	fmt.Printf("\n✓ Keyword generation complete!\n")
}

// fetchAndParseKeywords fetches kwlist.h from PostgreSQL repository and parses it
func fetchAndParseKeywords() ([]Keyword, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(PostgreSQLKwlistURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch kwlist.h: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch kwlist.h: HTTP %d", resp.StatusCode)
	}

	return parseKwlist(resp.Body)
}

// parseKwlist parses the kwlist.h file and extracts keyword definitions
func parseKwlist(r io.Reader) ([]Keyword, error) {
	// Regex to match PG_KEYWORD lines
	// Format: PG_KEYWORD("keyword", TOKEN_NAME, CATEGORY, LABEL)
	kwRegex := regexp.MustCompile(`PG_KEYWORD\("([^"]+)",\s*([A-Z_0-9]+),\s*([A-Z_]+),\s*([A-Z_]+)\)`)

	var keywords []Keyword
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip comments and empty lines
		if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "/*") || line == "" {
			continue
		}

		matches := kwRegex.FindStringSubmatch(line)
		if len(matches) == 5 {
			keywords = append(keywords, Keyword{
				Name:     matches[1],
				Token:    matches[2],
				Category: matches[3],
				Label:    matches[4],
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading kwlist.h: %w", err)
	}

	return keywords, nil
}

// categorizeKeywords groups keywords by their category
func categorizeKeywords(keywords []Keyword) map[string][]Keyword {
	categorized := make(map[string][]Keyword)

	for _, kw := range keywords {
		categorized[kw.Category] = append(categorized[kw.Category], kw)
	}

	return categorized
}

// applyTokenRename applies ANTLR reserved name renaming to token names
func applyTokenRename(token string) string {
	antlrReservedNames := map[string]bool{
		"SKIP": true,
	}

	// Automatically rename ANTLR-reserved tokens
	if antlrReservedNames[token] && !strings.HasSuffix(token, "_P") {
		return token + "_P"
	}
	return token
}

// generateANTLRGrammar generates ANTLR grammar file with keyword rules
func generateANTLRGrammar(categorized map[string][]Keyword, outputPath string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	// Write header
	fmt.Fprintf(w, "// ============================================================================\n")
	fmt.Fprintf(w, "// Auto-generated PostgreSQL Keyword Definitions\n")
	fmt.Fprintf(w, "// ============================================================================\n")
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// Source: PostgreSQL %s kwlist.h\n", PostgreSQLVersion)
	fmt.Fprintf(w, "// URL: %s\n", PostgreSQLKwlistURL)
	fmt.Fprintf(w, "// Generated: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// DO NOT EDIT MANUALLY - This file is generated by keyword-generator\n")
	fmt.Fprintf(w, "// To regenerate: cd postgresql && make generate-keywords\n")
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// Note: Keywords are defined using string literals (e.g., 'SELECT') rather\n")
	fmt.Fprintf(w, "// than token references to avoid token definition errors.\n")
	fmt.Fprintf(w, "// ============================================================================\n\n")

	// Write grammar fragment declaration
	fmt.Fprintf(w, "parser grammar PostgreSQLKeywords;\n\n")

	// Generate reserved_keyword rule
	if keywords := categorized[CategoryReserved]; len(keywords) > 0 {
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// Reserved Keywords (%d total)\n", len(keywords))
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// These keywords cannot be used as identifiers without quoting.\n")
		fmt.Fprintf(w, "// They are reserved because they would cause parser conflicts.\n")
		fmt.Fprintf(w, "//\n")
		fmt.Fprintf(w, "// Examples: SELECT, FROM, WHERE, CREATE, AND, OR\n")
		fmt.Fprintf(w, "// ============================================================================\n\n")
		fmt.Fprintf(w, "reserved_keyword\n")
		fmt.Fprintf(w, "   : %s\n", applyTokenRename(keywords[0].Token))
		for i := 1; i < len(keywords); i++ {
			fmt.Fprintf(w, "   | %s\n", applyTokenRename(keywords[i].Token))
		}
		fmt.Fprintf(w, "   ;\n\n")
	}

	// Generate unreserved_keyword rule
	if keywords := categorized[CategoryUnreserved]; len(keywords) > 0 {
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// Unreserved Keywords (%d total)\n", len(keywords))
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// These keywords can be used as identifiers in all contexts.\n")
		fmt.Fprintf(w, "// They do not cause parser conflicts.\n")
		fmt.Fprintf(w, "//\n")
		fmt.Fprintf(w, "// Examples: ABORT, ACCESS, ACTION, ADMIN, AFTER\n")
		fmt.Fprintf(w, "// ============================================================================\n\n")
		fmt.Fprintf(w, "unreserved_keyword\n")
		fmt.Fprintf(w, "   : %s\n", applyTokenRename(keywords[0].Token))
		for i := 1; i < len(keywords); i++ {
			fmt.Fprintf(w, "   | %s\n", applyTokenRename(keywords[i].Token))
		}
		fmt.Fprintf(w, "   ;\n\n")
	}

	// Generate col_name_keyword rule
	if keywords := categorized[CategoryColName]; len(keywords) > 0 {
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// Column Name Keywords (%d total)\n", len(keywords))
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// These keywords can be used as column names but may be restricted elsewhere.\n")
		fmt.Fprintf(w, "// Typically these are type names or function names.\n")
		fmt.Fprintf(w, "//\n")
		fmt.Fprintf(w, "// Examples: BETWEEN, BIGINT, BOOLEAN, INT, TIMESTAMP\n")
		fmt.Fprintf(w, "// ============================================================================\n\n")
		fmt.Fprintf(w, "col_name_keyword\n")
		fmt.Fprintf(w, "   : %s\n", applyTokenRename(keywords[0].Token))
		for i := 1; i < len(keywords); i++ {
			fmt.Fprintf(w, "   | %s\n", applyTokenRename(keywords[i].Token))
		}
		fmt.Fprintf(w, "   ;\n\n")
	}

	// Generate type_func_name_keyword rule
	if keywords := categorized[CategoryTypeFuncName]; len(keywords) > 0 {
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// Type and Function Name Keywords (%d total)\n", len(keywords))
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// These keywords can be used in type and function name contexts.\n")
		fmt.Fprintf(w, "// They may be restricted in other syntactic positions.\n")
		fmt.Fprintf(w, "//\n")
		fmt.Fprintf(w, "// Examples: AUTHORIZATION, BINARY, COLLATION, CROSS, JOIN\n")
		fmt.Fprintf(w, "// ============================================================================\n\n")
		fmt.Fprintf(w, "type_func_name_keyword\n")
		fmt.Fprintf(w, "   : %s\n", applyTokenRename(keywords[0].Token))
		for i := 1; i < len(keywords); i++ {
			fmt.Fprintf(w, "   | %s\n", applyTokenRename(keywords[i].Token))
		}
		fmt.Fprintf(w, "   ;\n\n")
	}

	return nil
}

// generateLexerTokens generates lexer token definitions for all keywords
func generateLexerTokens(keywords []Keyword, outputPath string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	// ANTLR reserved names that cannot be used as token names
	antlrReservedNames := map[string]bool{
		"SKIP": true,
	}

	// Write header
	fmt.Fprintf(w, "// ============================================================================\n")
	fmt.Fprintf(w, "// Auto-generated PostgreSQL Keyword Lexer Tokens\n")
	fmt.Fprintf(w, "// ============================================================================\n")
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// Source: PostgreSQL %s kwlist.h\n", PostgreSQLVersion)
	fmt.Fprintf(w, "// URL: %s\n", PostgreSQLKwlistURL)
	fmt.Fprintf(w, "// Generated: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// DO NOT EDIT MANUALLY - This file is generated by keyword-generator\n")
	fmt.Fprintf(w, "// To regenerate: cd postgresql && make generate-keywords\n")
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// This file contains all keyword token definitions from PostgreSQL's kwlist.h.\n")
	fmt.Fprintf(w, "// Import this into PostgreSQLLexer.g4 to use these tokens.\n")
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// NOTE: Some token names are automatically renamed to avoid ANTLR conflicts:\n")
	fmt.Fprintf(w, "//   - SKIP → SKIP_P (SKIP is ANTLR-reserved for skip channel)\n")
	fmt.Fprintf(w, "// ============================================================================\n\n")

	// Write lexer grammar declaration
	fmt.Fprintf(w, "lexer grammar PostgreSQLKeywordsLexer;\n\n")

	// Track renamed tokens for documentation
	var renamedTokens []struct {
		original string
		renamed  string
		keyword  string
	}

	fmt.Fprintf(w, "// ============================================================================\n")
	fmt.Fprintf(w, "// Keyword Tokens (%d total)\n", len(keywords))
	fmt.Fprintf(w, "// ============================================================================\n\n")

	// Generate token definitions for all keywords
	for _, kw := range keywords {
		tokenName := kw.Token

		// Automatically rename ANTLR-reserved tokens
		if antlrReservedNames[kw.Token] && !strings.HasSuffix(kw.Token, "_P") {
			tokenName = kw.Token + "_P"
			renamedTokens = append(renamedTokens, struct {
				original string
				renamed  string
				keyword  string
			}{kw.Token, tokenName, kw.Name})
		}

		fmt.Fprintf(w, "%s\n", tokenName)
		fmt.Fprintf(w, "   : '%s'\n", strings.ToUpper(kw.Name))
		fmt.Fprintf(w, "   ;\n\n")
	}

	// Document renamed tokens if any
	if len(renamedTokens) > 0 {
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// Automatically Renamed Tokens (ANTLR Compatibility)\n")
		fmt.Fprintf(w, "// ============================================================================\n")
		fmt.Fprintf(w, "// The following tokens were renamed to avoid ANTLR reserved name conflicts:\n")
		fmt.Fprintf(w, "//\n")
		for _, r := range renamedTokens {
			fmt.Fprintf(w, "//   %s → %s (keyword: '%s')\n", r.original, r.renamed, r.keyword)
		}
		fmt.Fprintf(w, "// ============================================================================\n")
	}

	return nil
}

// updateLexerWithKeywords updates PostgreSQLLexer.g4 with keyword definitions between markers
func updateLexerWithKeywords(keywords []Keyword, lexerPath string) error {
	// Read the existing lexer file
	content, err := os.ReadFile(lexerPath)
	if err != nil {
		return fmt.Errorf("failed to read lexer file: %w", err)
	}

	lexerContent := string(content)

	// Find the marker positions
	beginMarker := "// BEGIN AUTO-GENERATED KEYWORDS"
	endMarker := "// END AUTO-GENERATED KEYWORDS"

	beginIndex := strings.Index(lexerContent, beginMarker)
	endIndex := strings.Index(lexerContent, endMarker)

	if beginIndex == -1 || endIndex == -1 {
		return fmt.Errorf("markers not found in lexer file (BEGIN: %d, END: %d)", beginIndex, endIndex)
	}

	// Find the end of the BEGIN marker line
	beginLineEnd := strings.Index(lexerContent[beginIndex:], "\n")
	if beginLineEnd == -1 {
		return fmt.Errorf("invalid marker format")
	}
	insertStart := beginIndex + beginLineEnd + 1

	// Build the keyword definitions
	var builder strings.Builder
	builder.WriteString("//\n")
	builder.WriteString(fmt.Sprintf("// Source: PostgreSQL %s kwlist.h\n", PostgreSQLVersion))
	builder.WriteString(fmt.Sprintf("// URL: %s\n", PostgreSQLKwlistURL))
	builder.WriteString(fmt.Sprintf("// Generated: %s\n", time.Now().Format(time.RFC3339)))
	builder.WriteString(fmt.Sprintf("// Total Keywords: %d\n", len(keywords)))
	builder.WriteString("//\n")
	builder.WriteString("// NOTE: These keyword rules must appear BEFORE the Identifier rule\n")
	builder.WriteString("// to ensure keywords are matched with higher priority than identifiers.\n")
	builder.WriteString("//\n\n")

	// Track renamed tokens
	antlrReservedNames := map[string]bool{
		"SKIP": true,
	}
	var renamedTokens []struct {
		original string
		renamed  string
		keyword  string
	}

	// Generate token definitions for all keywords
	for _, kw := range keywords {
		tokenName := kw.Token

		// Automatically rename ANTLR-reserved tokens
		if antlrReservedNames[kw.Token] && !strings.HasSuffix(kw.Token, "_P") {
			tokenName = kw.Token + "_P"
			renamedTokens = append(renamedTokens, struct {
				original string
				renamed  string
				keyword  string
			}{kw.Token, tokenName, kw.Name})
		}

		builder.WriteString(fmt.Sprintf("%s\n", tokenName))
		builder.WriteString(fmt.Sprintf("   : '%s'\n", strings.ToUpper(kw.Name)))
		builder.WriteString("   ;\n\n")
	}

	// Document renamed tokens if any
	if len(renamedTokens) > 0 {
		builder.WriteString("// ============================================================================\n")
		builder.WriteString("// Automatically Renamed Tokens (ANTLR Compatibility)\n")
		builder.WriteString("// ============================================================================\n")
		builder.WriteString("// The following tokens were renamed to avoid ANTLR reserved name conflicts:\n")
		builder.WriteString("//\n")
		for _, r := range renamedTokens {
			builder.WriteString(fmt.Sprintf("//   %s → %s (keyword: '%s')\n", r.original, r.renamed, r.keyword))
		}
		builder.WriteString("// ============================================================================\n\n")
	}

	// Reconstruct the file
	newContent := lexerContent[:insertStart] + builder.String() + lexerContent[endIndex:]

	// Write the updated content back
	err = os.WriteFile(lexerPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write lexer file: %w", err)
	}

	return nil
}
