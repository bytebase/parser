package doris_test

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	doris "github.com/bytebase/parser/doris"
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

func TestDorisParser(t *testing.T) {
	err := filepath.WalkDir("examples", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".sql") {
			return nil
		}
		t.Run(filePath, func(t *testing.T) {
			t.Parallel()
			runParserTest(t, filePath)
		})
		return nil
	})
	require.NoError(t, err)
}

func runParserTest(t *testing.T, filePath string) {
	// read all the bytes from the file
	data, err := ioutil.ReadFile(filePath)
	require.NoError(t, err)

	dataString := strings.TrimRight(string(data), " \t\r\n;") + "\n;"

	input := antlr.NewInputStream(dataString)

	lexer := doris.NewDorisLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := doris.NewDorisParser(stream)

	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true

	tree := p.MultiStatements()

	require.Equal(t, 0, lexerErrors.errors)
	require.Equal(t, 0, parserErrors.errors)

	require.Equal(t, dataString, stream.GetTextFromRuleContext(tree))
}
