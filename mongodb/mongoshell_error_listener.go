package mongodb

import "github.com/antlr4-go/antlr/v4"

// MongoShellErrorListener collects parse errors.
type MongoShellErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []*MongoShellParseError
}

// NewMongoShellErrorListener creates a new error listener.
func NewMongoShellErrorListener() *MongoShellErrorListener {
	return &MongoShellErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               make([]*MongoShellParseError, 0),
	}
}

// SyntaxError is called when a syntax error is encountered.
func (l *MongoShellErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.Errors = append(l.Errors, &MongoShellParseError{
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

// HasErrors returns true if any errors were collected.
func (l *MongoShellErrorListener) HasErrors() bool {
	return len(l.Errors) > 0
}
