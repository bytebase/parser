package lexer

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/bytebase/parser/tools/fuzzing/internal/grammar"
)

// TokenGenerator generates tokens from lexer rules
type TokenGenerator struct {
	random *rand.Rand
	config *TokenGeneratorConfig
}

// TokenGeneratorConfig controls token generation behavior
type TokenGeneratorConfig struct {
	// MaxQuantifierCount limits how many times quantified elements repeat
	MaxQuantifierCount int
	// MinQuantifierCount sets minimum repetitions for + quantifiers
	MinQuantifierCount int
	// OptionalProbability controls likelihood of including optional elements (0.0-1.0)
	OptionalProbability float64
	// MaxDepth limits recursion depth to prevent infinite loops
	MaxDepth int
}

// NewTokenGenerator creates a new token generator
func NewTokenGenerator(seed int64, config *TokenGeneratorConfig) *TokenGenerator {
	if config == nil {
		config = &TokenGeneratorConfig{
			MaxQuantifierCount:  5,
			MinQuantifierCount:  1,
			OptionalProbability: 0.7,
			MaxDepth:           10,
		}
	}
	return &TokenGenerator{
		random: rand.New(rand.NewSource(seed)),
		config: config,
	}
}

// GenerateToken generates a token string from a lexer rule
func (g *TokenGenerator) GenerateToken(rule *grammar.Rule) (string, error) {
	if !rule.IsLexer {
		return "", fmt.Errorf("rule %s is not a lexer rule", rule.Name)
	}

	if len(rule.Alternatives) == 0 {
		return "", fmt.Errorf("rule %s has no alternatives", rule.Name)
	}

	// Select a random alternative
	altIndex := g.random.Intn(len(rule.Alternatives))
	alternative := rule.Alternatives[altIndex]

	// Generate from the selected alternative
	return g.generateFromAlternative(&alternative, 0)
}

// generateFromAlternative generates text from a lexer rule alternative
func (g *TokenGenerator) generateFromAlternative(alt *grammar.Alternative, depth int) (string, error) {
	if depth > g.config.MaxDepth {
		return "", fmt.Errorf("maximum depth exceeded")
	}

	var result strings.Builder
	for _, element := range alt.Elements {
		text, err := g.generateFromElement(&element, depth+1)
		if err != nil {
			return "", err
		}
		result.WriteString(text)
	}
	return result.String(), nil
}

// generateFromElement generates text from a single lexer element
func (g *TokenGenerator) generateFromElement(element *grammar.Element, depth int) (string, error) {
	if depth > g.config.MaxDepth {
		return "", fmt.Errorf("maximum depth exceeded")
	}

	// Handle quantifiers
	switch element.Quantifier {
	case grammar.OPTIONAL_Q: // ?
		if g.random.Float64() > g.config.OptionalProbability {
			return "", nil // Skip optional element
		}
		return g.generateElementContent(element, depth)

	case grammar.ZERO_MORE: // *
		count := g.random.Intn(g.config.MaxQuantifierCount + 1) // 0 to MaxQuantifierCount
		return g.generateRepeated(element, count, depth)

	case grammar.ONE_MORE: // +
		count := g.config.MinQuantifierCount + g.random.Intn(g.config.MaxQuantifierCount)
		return g.generateRepeated(element, count, depth)

	default: // NONE
		return g.generateElementContent(element, depth)
	}
}

// generateRepeated generates repeated content for quantified elements
func (g *TokenGenerator) generateRepeated(element *grammar.Element, count int, depth int) (string, error) {
	var result strings.Builder
	for i := 0; i < count; i++ {
		text, err := g.generateElementContent(element, depth)
		if err != nil {
			return "", err
		}
		result.WriteString(text)
	}
	return result.String(), nil
}

// generateElementContent generates the actual content for an element
func (g *TokenGenerator) generateElementContent(element *grammar.Element, depth int) (string, error) {
	switch value := element.Value.(type) {
	case grammar.LiteralValue:
		return g.generateFromLiteral(value)
	case grammar.BlockValue:
		return g.generateFromBlock(value, depth)
	case grammar.WildcardValue:
		return g.generateFromWildcard()
	case grammar.ReferenceValue:
		// For lexer rules, this typically shouldn't happen unless it's a fragment reference
		// For now, return the reference name as placeholder
		return fmt.Sprintf("<%s>", value.Name), nil
	default:
		return "", fmt.Errorf("unsupported element value type: %T", value)
	}
}

// generateFromLiteral generates text from a literal value
func (g *TokenGenerator) generateFromLiteral(literal grammar.LiteralValue) (string, error) {
	text := literal.Text

	// Handle string literals - remove quotes
	if len(text) >= 2 && text[0] == '\'' && text[len(text)-1] == '\'' {
		return text[1 : len(text)-1], nil
	}

	// Handle negated sets like ~[...] FIRST (before checking for ..)
	if strings.HasPrefix(text, "~[") && strings.HasSuffix(text, "]") {
		return g.generateFromNegatedSet(text)
	}

	// Handle character sets like [a-zA-Z]
	if len(text) >= 2 && text[0] == '[' && text[len(text)-1] == ']' {
		return g.generateFromCharacterSet(text[1 : len(text)-1])
	}

	// Handle character ranges like 'a'..'z'
	if strings.Contains(text, "..") {
		return g.generateFromCharacterRange(text)
	}

	// Default: return the literal as-is
	return text, nil
}

// generateFromCharacterSet generates a character from a character set like [a-zA-Z0-9_]
func (g *TokenGenerator) generateFromCharacterSet(charset string) (string, error) {
	chars, err := g.expandCharacterSet(charset)
	if err != nil {
		return "", err
	}
	if len(chars) == 0 {
		return "", fmt.Errorf("empty character set")
	}
	
	// Select a random character from the set
	index := g.random.Intn(len(chars))
	return string(chars[index]), nil
}

// expandCharacterSet expands a character set specification into actual characters
func (g *TokenGenerator) expandCharacterSet(charset string) ([]rune, error) {
	var chars []rune
	i := 0
	
	for i < len(charset) {
		// Handle escape sequences
		if i < len(charset) && charset[i] == '\\' && i+1 < len(charset) {
			switch charset[i+1] {
			case 'r':
				chars = append(chars, '\r')
			case 'n':
				chars = append(chars, '\n')
			case 't':
				chars = append(chars, '\t')
			case '\\':
				chars = append(chars, '\\')
			case '"':
				chars = append(chars, '"')
			case '\'':
				chars = append(chars, '\'')
			default:
				// For unknown escapes, use the escaped character literally
				chars = append(chars, rune(charset[i+1]))
			}
			i += 2
		} else if i+2 < len(charset) && charset[i+1] == '-' && charset[i+2] != '\\' {
			// Handle range like a-z (but not when second char is an escape)
			start := rune(charset[i])
			end := rune(charset[i+2])
			
			if start > end {
				return nil, fmt.Errorf("invalid character range: %c-%c", start, end)
			}
			
			for c := start; c <= end; c++ {
				chars = append(chars, c)
			}
			i += 3
		} else {
			// Handle single character
			chars = append(chars, rune(charset[i]))
			i++
		}
	}
	
	return chars, nil
}

// generateFromCharacterRange generates from a character range like 'a'..'z'
func (g *TokenGenerator) generateFromCharacterRange(rangeText string) (string, error) {
	// Extract start and end characters from 'a'..'z' format
	parts := strings.Split(rangeText, "..")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid character range format: %s", rangeText)
	}
	
	start := strings.Trim(parts[0], "'\"")
	end := strings.Trim(parts[1], "'\"")
	
	if len(start) != 1 || len(end) != 1 {
		return "", fmt.Errorf("character range must be single characters: %s", rangeText)
	}
	
	startChar := rune(start[0])
	endChar := rune(end[0])
	
	if startChar > endChar {
		return "", fmt.Errorf("invalid character range: %c > %c", startChar, endChar)
	}
	
	// Generate random character in range
	rangeSize := int(endChar - startChar + 1)
	offset := g.random.Intn(rangeSize)
	result := startChar + rune(offset)
	
	return string(result), nil
}

// generateFromBlock generates text from a block value
func (g *TokenGenerator) generateFromBlock(block grammar.BlockValue, depth int) (string, error) {
	if len(block.Alternatives) == 0 {
		return "", nil
	}
	
	// Select a random alternative from the block
	altIndex := g.random.Intn(len(block.Alternatives))
	alternative := &block.Alternatives[altIndex]
	
	return g.generateFromAlternative(alternative, depth)
}

// generateFromWildcard generates a character for wildcard (.)
func (g *TokenGenerator) generateFromWildcard() (string, error) {
	// Generate a random printable ASCII character
	// Range: 32-126 (space to tilde)
	char := rune(32 + g.random.Intn(95))
	return string(char), nil
}

// generateFromNegatedSet generates a character NOT in the specified set
func (g *TokenGenerator) generateFromNegatedSet(negatedSet string) (string, error) {
	// Extract the character set from ~[...] format
	if len(negatedSet) < 4 || !strings.HasPrefix(negatedSet, "~[") || !strings.HasSuffix(negatedSet, "]") {
		return "", fmt.Errorf("invalid negated set format: %s", negatedSet)
	}
	
	charset := negatedSet[2 : len(negatedSet)-1] // Remove ~[ and ]
	
	// Expand the excluded character set
	excludedChars, err := g.expandCharacterSet(charset)
	if err != nil {
		return "", fmt.Errorf("failed to expand excluded character set: %v", err)
	}
	
	// Create a map for quick lookup
	excluded := make(map[rune]bool)
	for _, c := range excludedChars {
		excluded[c] = true
	}
	
	// Generate a character that's not in the excluded set
	// Try common printable ASCII characters first
	candidates := []rune{}
	
	// Add letters
	for c := 'a'; c <= 'z'; c++ {
		if !excluded[c] {
			candidates = append(candidates, c)
		}
	}
	for c := 'A'; c <= 'Z'; c++ {
		if !excluded[c] {
			candidates = append(candidates, c)
		}
	}
	
	// Add digits
	for c := '0'; c <= '9'; c++ {
		if !excluded[c] {
			candidates = append(candidates, c)
		}
	}
	
	// Add some special characters
	specialChars := []rune{' ', '!', '#', '$', '%', '&', '*', '+', '/', '=', '?', '@', '^', '_', '`', '|', '~'}
	for _, c := range specialChars {
		if !excluded[c] {
			candidates = append(candidates, c)
		}
	}
	
	if len(candidates) == 0 {
		return "", fmt.Errorf("no valid characters available (all excluded)")
	}
	
	// Select a random candidate
	index := g.random.Intn(len(candidates))
	return string(candidates[index]), nil
}

// ValidateCharacterSet validates if a character set specification is valid
func ValidateCharacterSet(charset string) error {
	// Use regex to validate basic character set patterns
	validPattern := regexp.MustCompile(`^[a-zA-Z0-9_\-\[\]\\^]+$`)
	if !validPattern.MatchString(charset) {
		return fmt.Errorf("invalid characters in character set: %s", charset)
	}
	return nil
}