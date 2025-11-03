package tidb

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type TiDBBaseLexer struct {
	*antlr.BaseLexer

	pendingTokens      []antlr.Token
	reservedKeywordMap map[string]bool
}

// NextToken implements antlr.TokenSource
func (l *TiDBBaseLexer) NextToken() antlr.Token {
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
func (l *TiDBBaseLexer) EmitDot() {
	dot := l.GetTokenFactory().Create(l.GetTokenSourceCharStreamPair(), TiDBLexerDOT_SYMBOL, ".", antlr.TokenDefaultChannel, l.TokenStartCharIndex, l.TokenStartCharIndex, l.TokenStartLine, l.TokenStartColumn)
	l.pendingTokens = append(l.pendingTokens, dot)
	l.TokenStartCharIndex++
}

// DetermineNumericType determines the numeric type of a given text.
func (l *TiDBBaseLexer) DetermineNumericType(text string) int {
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
		return TiDBLexerINT_NUMBER
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
		return TiDBLexerINT_NUMBER
	}

	var smaller, bigger int
	var cmp string

	if negative {
		if length == longLen {
			cmp = signedLongStr[1:]
			smaller = TiDBLexerINT_NUMBER
			bigger = TiDBLexerLONG_NUMBER
		} else if length < signedLongLongLen {
			return TiDBLexerLONG_NUMBER
		} else if length > signedLongLongLen {
			return TiDBLexerDECIMAL_NUMBER
		} else {
			cmp = signedLongLongStr[1:]
			smaller = TiDBLexerLONG_NUMBER
			bigger = TiDBLexerDECIMAL_NUMBER
		}
	} else {
		if length == longLen {
			cmp = longStr
			smaller = TiDBLexerINT_NUMBER
			bigger = TiDBLexerLONG_NUMBER
		} else if length < longLongLen {
			return TiDBLexerLONG_NUMBER
		} else if length > longLongLen {
			if length > unsignedLongLongLen {
				return TiDBLexerDECIMAL_NUMBER
			}
			cmp = unsignedLongLongStr
			smaller = TiDBLexerULONGLONG_NUMBER
			bigger = TiDBLexerDECIMAL_NUMBER
		} else {
			cmp = longLongStr
			smaller = TiDBLexerLONG_NUMBER
			bigger = TiDBLexerULONGLONG_NUMBER
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
func (l *TiDBBaseLexer) DetermineFunction(proposed int) int {
	if l.GetInputStream().LA(1) == int('(') {
		return proposed
	}
	return TiDBLexerIDENTIFIER
}

// CheckCharset checks the charset of a given text.
func (l *TiDBBaseLexer) CheckCharset(test string) int {
	switch test {
	case "_utf8", "_utf8mb3", "_utf8mb4", "_ucs2", "_big5", "_latin2",
		"_ujis", "_binary", "_cp1250", "_latin1":
		return TiDBLexerUNDERSCORE_CHARSET
	default:
		return TiDBLexerIDENTIFIER
	}
}
