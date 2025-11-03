package tsql_test

import (
	"os"
	"path"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	tsql "github.com/bytebase/parser/tsql"
	"github.com/stretchr/testify/require"
)

type CustomErrorListener struct {
	errors int
}

func NewCustomErrorListener() *CustomErrorListener {
	return new(CustomErrorListener)
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors += 1
	antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (l *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func TestTSQLParser(t *testing.T) {
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	for _, file := range examples {
		if file.Name() != "test.sql" {
			continue
		}
		filePath := path.Join("examples", file.Name())
		t.Run(filePath, func(t *testing.T) {
			t.Parallel()
			input, err := antlr.NewFileStream(filePath)
			require.NoError(t, err)

			lexer := tsql.NewTSqlLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := tsql.NewTSqlParser(stream)

			lexerErrors := &CustomErrorListener{}
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrors)

			parserErrors := &CustomErrorListener{}
			p.RemoveErrorListeners()
			p.AddErrorListener(parserErrors)

			p.BuildParseTrees = true

			_ = p.Tsql_file()

			require.Equal(t, 0, lexerErrors.errors)
			require.Equal(t, 0, parserErrors.errors)
		})
	}
}

func TestSpatialIndexParser(t *testing.T) {
	filePath := "examples/spatial_index.sql"
	input, err := antlr.NewFileStream(filePath)
	require.NoError(t, err)

	lexer := tsql.NewTSqlLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := tsql.NewTSqlParser(stream)

	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true

	_ = p.Tsql_file()

	require.Equal(t, 0, lexerErrors.errors)
	require.Equal(t, 0, parserErrors.errors)
}

func TestDDLIndexParser(t *testing.T) {
	filePath := "examples/ddl_index.sql"
	input, err := antlr.NewFileStream(filePath)
	require.NoError(t, err)

	lexer := tsql.NewTSqlLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := tsql.NewTSqlParser(stream)

	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true

	_ = p.Tsql_file()

	require.Equal(t, 0, lexerErrors.errors)
	require.Equal(t, 0, parserErrors.errors)
}

func TestSpatialComprehensiveParser(t *testing.T) {
	filePath := "examples/spatial_comprehensive.sql"
	input, err := antlr.NewFileStream(filePath)
	require.NoError(t, err)

	lexer := tsql.NewTSqlLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := tsql.NewTSqlParser(stream)

	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true

	_ = p.Tsql_file()

	require.Equal(t, 0, lexerErrors.errors)
	require.Equal(t, 0, parserErrors.errors)
}
