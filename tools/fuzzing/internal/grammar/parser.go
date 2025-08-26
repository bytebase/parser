package grammar

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"
	grammar "github.com/bytebase/parser/tools/grammar"
)

// ParsedGrammar represents a parsed grammar with extracted rules
type ParsedGrammar struct {
	LexerRules  map[string]*Rule
	ParserRules map[string]*Rule
	FilePath    string
}

// Rule represents a grammar rule with its alternatives
type Rule struct {
	Name         string
	Alternatives []Alternative
	IsLexer      bool
}

// Alternative represents one alternative of a rule
type Alternative struct {
	Elements []Element
}

// Element represents an element within an alternative
type Element struct {
	Type       ElementType
	Value      string
	Quantifier Quantifier
	Min, Max   int // for {n,m} quantifiers
}

// ElementType indicates the type of grammar element
type ElementType int

const (
	RULE_REF ElementType = iota
	TOKEN_REF
	LITERAL
	OPTIONAL
	QUANTIFIED
)

// Quantifier indicates repetition type
type Quantifier int

const (
	NONE Quantifier = iota
	OPTIONAL_Q // ?
	ZERO_MORE  // *
	ONE_MORE   // +
	RANGE      // {n,m}
)

// ParseGrammarFile parses a .g4 file and extracts rules for fuzzing
func ParseGrammarFile(filePath string) (*ParsedGrammar, error) {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read grammar file")
	}

	if len(content) == 0 {
		return nil, errors.New("grammar file is empty")
	}

	// Create input stream
	input := antlr.NewInputStream(string(content))

	// Create lexer
	lexer := grammar.NewANTLRv4Lexer(input)

	// Add error listener
	errorListener := &GrammarErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	parser := grammar.NewANTLRv4Parser(stream)

	// Add error listener to parser
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	// Parse the grammar
	tree := parser.GrammarSpec()

	// Check for parsing errors
	if errorListener.HasErrors() {
		return nil, errors.Errorf("failed to parse grammar: %v", errorListener.GetErrors())
	}

	if tree == nil {
		return nil, errors.New("parser returned nil tree")
	}

	// Extract rules from parse tree
	visitor := &GrammarExtractorVisitor{
		lexerRules:  make(map[string]*Rule),
		parserRules: make(map[string]*Rule),
	}

	visitor.Visit(tree)

	return &ParsedGrammar{
		LexerRules:  visitor.lexerRules,
		ParserRules: visitor.parserRules,
		FilePath:    filePath,
	}, nil
}

// GetRule gets a rule by name from either lexer or parser rules
func (g *ParsedGrammar) GetRule(name string) *Rule {
	if rule, ok := g.ParserRules[name]; ok {
		return rule
	}
	if rule, ok := g.LexerRules[name]; ok {
		return rule
	}
	return nil
}

// GetAllRules returns all rules (both lexer and parser)
func (g *ParsedGrammar) GetAllRules() map[string]*Rule {
	allRules := make(map[string]*Rule)
	for name, rule := range g.LexerRules {
		allRules[name] = rule
	}
	for name, rule := range g.ParserRules {
		allRules[name] = rule
	}
	return allRules
}

// IsRule checks if an element refers to another rule
func (e *Element) IsRule() bool {
	return e.Type == RULE_REF || e.Type == TOKEN_REF
}

// IsTerminal checks if an element is a terminal (literal)
func (e *Element) IsTerminal() bool {
	return e.Type == LITERAL
}

// IsOptional checks if an element has optional quantifier
func (e *Element) IsOptional() bool {
	return e.Quantifier == OPTIONAL_Q
}

// IsQuantified checks if an element has repetition quantifiers
func (e *Element) IsQuantified() bool {
	return e.Quantifier == ZERO_MORE || e.Quantifier == ONE_MORE || e.Quantifier == RANGE
}

// GrammarErrorListener collects parsing errors
type GrammarErrorListener struct {
	errors []string
}

func (l *GrammarErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func (l *GrammarErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Ignore ambiguity for fuzzing purposes
}

func (l *GrammarErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// Ignore for fuzzing purposes
}

func (l *GrammarErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	// Ignore for fuzzing purposes
}

func (l *GrammarErrorListener) HasErrors() bool {
	return len(l.errors) > 0
}

func (l *GrammarErrorListener) GetErrors() []string {
	return l.errors
}

// GrammarExtractorVisitor extracts rules from the parse tree
type GrammarExtractorVisitor struct {
	antlr.ParseTreeVisitor
	lexerRules  map[string]*Rule
	parserRules map[string]*Rule
	isLexer     bool
}

func (v *GrammarExtractorVisitor) Visit(tree antlr.ParseTree) interface{} {
	// TODO: Implement tree visiting to extract rules
	// This is a placeholder - we'll implement the actual visitor logic
	// to walk the parse tree and extract rule information
	
	// For now, let's create a simple placeholder structure
	v.extractPlaceholderRules()
	
	return nil
}

// extractPlaceholderRules creates placeholder rules for testing
func (v *GrammarExtractorVisitor) extractPlaceholderRules() {
	// Add some basic rules for testing
	v.parserRules["selectStmt"] = &Rule{
		Name:    "selectStmt",
		IsLexer: false,
		Alternatives: []Alternative{
			{
				Elements: []Element{
					{Type: LITERAL, Value: "SELECT"},
					{Type: RULE_REF, Value: "columnList"},
					{Type: LITERAL, Value: "FROM"},
					{Type: RULE_REF, Value: "tableRef"},
					{Type: RULE_REF, Value: "whereClause", Quantifier: OPTIONAL_Q},
				},
			},
		},
	}
	
	v.parserRules["columnList"] = &Rule{
		Name:    "columnList",
		IsLexer: false,
		Alternatives: []Alternative{
			{
				Elements: []Element{
					{Type: RULE_REF, Value: "column"},
					{
						Type:       RULE_REF,
						Value:      "column",
						Quantifier: ZERO_MORE,
					},
				},
			},
		},
	}

	v.lexerRules["SELECT"] = &Rule{
		Name:    "SELECT",
		IsLexer: true,
		Alternatives: []Alternative{
			{
				Elements: []Element{
					{Type: LITERAL, Value: "'SELECT'"},
				},
			},
		},
	}
}