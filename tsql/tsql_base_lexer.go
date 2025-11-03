package tsql

import "github.com/antlr4-go/antlr/v4"

type TSqlBaseLexer struct {
	*antlr.BaseLexer
}

func (l *TSqlBaseLexer) IsID_(tokenType int) bool {
	switch tokenType {
	case TSqlLexerID:
		return true
	case TSqlLexerTEMP_ID:
		return true
	case TSqlLexerDOUBLE_QUOTE_ID:
		return true
	case TSqlLexerDOUBLE_QUOTE_BLANK:
		return true
	case TSqlLexerSQUARE_BRACKET_ID:
		return true
	case TSqlLexerRAW:
		return true
	}

	return l.IsKeyword(tokenType)
}

func (l *TSqlBaseLexer) IsKeyword(tokenType int) bool {
	if _, ok :=  keywordMap[tokenType]; ok{
		return true
	}
	return false
}
