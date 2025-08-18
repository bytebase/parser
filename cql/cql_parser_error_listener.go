package cql

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// ParseError represents a CQL parsing error with position information
type ParseError struct {
	Line   int
	Column int
	Msg    string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d:%d %s", e.Line, e.Column, e.Msg)
}

// ErrorListener collects parsing errors
type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []ParseError
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: &antlr.DefaultErrorListener{},
		Errors:              make([]ParseError, 0),
	}
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, ParseError{
		Line:   line,
		Column: column,
		Msg:    msg,
	})
}