package mysql

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type MySQLBaseLexer struct {
	*antlr.BaseLexer

	pendingTokens      []antlr.Token
	reservedKeywordMap map[string]bool
}

// NextToken implements antlr.TokenSource
func (l *MySQLBaseLexer) NextToken() antlr.Token {
	// First respond with pending tokens to the next token request, if there are any.
	if len(l.pendingTokens) != 0 {
		pending := l.pendingTokens[0]
		l.pendingTokens = l.pendingTokens[1:]
		return pending
	}

	// Let the main lexer class run the next token recognition.
	// This might create additional tokens again.
	next := l.BaseLexer.NextToken()
	if len(l.pendingTokens) != 0 {
		pending := l.pendingTokens[0]
		l.pendingTokens = l.pendingTokens[1:]
		l.pendingTokens = append(l.pendingTokens, next)
		return pending
	}
	return next
}

// EmitDot puts a DOT token onto the pending token list.
func (l *MySQLBaseLexer) EmitDot() {
	dot := l.GetTokenFactory().Create(l.GetTokenSourceCharStreamPair(), MySQLLexerDOT_SYMBOL, ".", antlr.TokenDefaultChannel, l.TokenStartCharIndex, l.TokenStartCharIndex, l.TokenStartLine, l.TokenStartColumn)
	l.pendingTokens = append(l.pendingTokens, dot)
	l.TokenStartCharIndex++
}

// DetermineNumericType determines the numeric type of a given text.
func (l *MySQLBaseLexer) DetermineNumericType(text string) int {
	longStr := "2147483647"
	longLen := len(longStr)
	signedLongStr := "-2147483648"
	longLongStr := "9223372036854775807"
	longLongLen := len(longLongStr)
	signedLongLongStr := "-9223372036854775808"
	signedLongLongLen := len(signedLongLongStr) - 1 // -1 because we don't count the leading '-'
	unsignedLongLongStr := "18446744073709551615"
	unsignedLongLongLen := len(unsignedLongLongStr)

	// The original code checks for leading +/- but actually that can never happen, neither in the
	// server parser (as a digit is used to trigger processing in the lexer) nor in our parser
	// as our rules are defined without signs. But we do it anyway for maximum compatibility.
	length := len(text)
	if length < longLen {
		return MySQLLexerINT_NUMBER
	}
	negative := false

	if strings.HasPrefix(text, "+") {
		text = text[1:]
		length--
	} else if strings.HasPrefix(text, "-") {
		text = text[1:]
		length--
		negative = true
	}

	text = strings.TrimLeft(text, "0")
	length = len(text)

	if length < longLen {
		return MySQLLexerINT_NUMBER
	}

	var smaller, bigger int
	var cmp string

	if negative {
		if length == longLen {
			cmp = signedLongStr[1:]
			smaller = MySQLLexerINT_NUMBER
			bigger = MySQLLexerLONG_NUMBER
		} else if length < signedLongLongLen {
			return MySQLLexerLONG_NUMBER
		} else if length > signedLongLongLen {
			return MySQLLexerDECIMAL_NUMBER
		} else {
			cmp = signedLongLongStr[1:]
			smaller = MySQLLexerLONG_NUMBER
			bigger = MySQLLexerDECIMAL_NUMBER
		}
	} else {
		if length == longLen {
			cmp = longStr
			smaller = MySQLLexerINT_NUMBER
			bigger = MySQLLexerLONG_NUMBER
		} else if length < longLongLen {
			return MySQLLexerLONG_NUMBER
		} else if length > longLongLen {
			if length > unsignedLongLongLen {
				return MySQLLexerDECIMAL_NUMBER
			}
			cmp = unsignedLongLongStr
			smaller = MySQLLexerULONGLONG_NUMBER
			bigger = MySQLLexerDECIMAL_NUMBER
		} else {
			cmp = longLongStr
			smaller = MySQLLexerLONG_NUMBER
			bigger = MySQLLexerULONGLONG_NUMBER
		}
	}

	for i := 0; i < len(cmp); i++ {
		if cmp[i] != text[i] {
			if text[i] < cmp[i] {
				return smaller
			} else {
				return bigger
			}
		}
	}

	return smaller
}

// DetermineFunction determines the function type of a given text.
func (l *MySQLBaseLexer) DetermineFunction(proposed int) int {
	if l.GetInputStream().LA(1) == int('(') {
		return proposed
	}
	return MySQLLexerIDENTIFIER
}

// CheckCharset checks the charset of a given text.
func (l *MySQLBaseLexer) CheckCharset(test string) int {
	switch test {
	case "_utf8", "_utf8mb3", "_utf8mb4", "_ucs2", "_big5", "_latin2",
		"_ujis", "_binary", "_cp1250", "_latin1":
		return MySQLLexerUNDERSCORE_CHARSET
	default:
		return MySQLLexerIDENTIFIER
	}
}

func (l *MySQLBaseLexer) IsIdentifier(tokenType int) bool {
	if tokenType == antlr.TokenEOF {
		return false
	}

	if tokenType == MySQLLexerIDENTIFIER || tokenType == MySQLLexerBACK_TICK_QUOTED_ID {
		return true
	}

	// TODO: Double quoted text represents identifiers only if the ANSI QUOTES sql mode is active.

	symbol := l.GetSymbolicNames()[tokenType]
	symbol = strings.TrimSuffix(symbol, "_SYMBOL")
	if len(symbol) > 0 && !l.IsReservedKeyword(symbol) {
		return true
	}
	return false
}

// IsReservedKeyword checks if a given string is a reserved keyword.
// Currently, we don't check by server version, if it's a reserved keyword in 5.6, 5.7 or 8.0, return true.
func (l *MySQLBaseLexer) IsReservedKeyword(s string) bool {
	l.initReversedKeywordMap()

	_, exists := l.reservedKeywordMap[s]
	return exists
}

func (l *MySQLBaseLexer) initReversedKeywordMap() {
	if l.reservedKeywordMap != nil {
		return
	}

	l.reservedKeywordMap = make(map[string]bool)

	for _, keyword := range Keywords56 {
		if keyword.Reserved {
			l.reservedKeywordMap[keyword.Keyword] = true
		}
	}

	for _, keyword := range Keywords57 {
		if keyword.Reserved {
			l.reservedKeywordMap[keyword.Keyword] = true
		}
	}

	for _, keyword := range Keywords80 {
		if keyword.Reserved {
			l.reservedKeywordMap[keyword.Keyword] = true
		}
	}
}

// GetBuiltinFunctions returns the list of builtin functions.
func GetBuiltinFunctions() []string {
	var result []string
	result = append(result, builtinFunctions...)
	return result
}
