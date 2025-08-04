package redshift

import "github.com/antlr4-go/antlr/v4"

type RedshiftParserErrorListener struct {
	grammar *RedshiftParser
}

var _ antlr.ErrorListener = &RedshiftParserErrorListener{}

func (receiver RedshiftParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	receiver.grammar.parseErrors = append(receiver.grammar.parseErrors, &RedshiftParseError{
		Number:  0,
		Offset:  0,
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

func (receiver RedshiftParserErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// ignore
}

func (receiver RedshiftParserErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// ignore
}

func (receiver RedshiftParserErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	// ignore
}
