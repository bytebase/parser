package trino_test

import (
	"os"
	"path"
	"sort"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	trino "github.com/bytebase/parser/trino"
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

func TestTrinoParser(t *testing.T) {
	testFilePaths := scanSQLFileInDirRecursive(t, "examples")
	sort.StringSlice(testFilePaths).Sort()

	for _, fp := range testFilePaths {
		t.Run(fp, func(t *testing.T) {
			t.Parallel()
			
			input, err := antlr.NewFileStream(fp)
			require.NoError(t, err)

			lexer := trino.NewTrinoLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := trino.NewTrinoParser(stream)

			lexerErrors := &CustomErrorListener{}
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrors)

			parserErrors := &CustomErrorListener{}
			p.RemoveErrorListeners()
			p.AddErrorListener(parserErrors)

			p.BuildParseTrees = true

			_ = p.Parse()

			require.Equal(t, 0, lexerErrors.errors, "file: %s", fp)
			require.Equal(t, 0, parserErrors.errors, "file: %s", fp)
		})
	}
}

func TestBenchmarkTrinoParser(t *testing.T) {
	testFilePaths := scanSQLFileInDirRecursive(t, "examples")
	sort.StringSlice(testFilePaths).Sort()

	for _, fp := range testFilePaths {
		t.Run(fp, func(t *testing.T) {
			data, err := os.ReadFile(fp)
			require.NoError(t, err)
			
			sql := string(data)
			
			input := antlr.NewInputStream(sql)

			lexer := trino.NewTrinoLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := trino.NewTrinoParser(stream)

			lexerErrors := &CustomErrorListener{}
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrors)

			parserErrors := &CustomErrorListener{}
			p.RemoveErrorListeners()
			p.AddErrorListener(parserErrors)

			p.BuildParseTrees = false
			_ = p.Parse()

			require.Equal(t, 0, lexerErrors.errors)
			require.Equal(t, 0, parserErrors.errors)
		})
	}
}

func scanSQLFileInDirRecursive(t *testing.T, dir string) []string {
	var fps []string

	files, err := os.ReadDir(dir)
	require.NoError(t, err)

	for _, file := range files {
		if file.IsDir() {
			rfps := scanSQLFileInDirRecursive(t, path.Join(dir, file.Name()))
			fps = append(fps, rfps...)
		} else if strings.HasSuffix(strings.ToLower(file.Name()), ".sql") {
			fps = append(fps, path.Join(dir, file.Name()))
		}
	}

	return fps
}