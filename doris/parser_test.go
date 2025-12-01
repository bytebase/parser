package doris_test

import (
	"io/ioutil"
	"os"
	"path"
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
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	for _, file := range examples {
		if file.IsDir() {
			// Handle subdirectories like regression/
			subdir := path.Join("examples", file.Name())
			subFiles, err := os.ReadDir(subdir)
			require.NoError(t, err)
			for _, subFile := range subFiles {
				if subFile.IsDir() || !strings.HasSuffix(subFile.Name(), ".sql") {
					continue
				}
				filePath := path.Join(subdir, subFile.Name())
				t.Run(filePath, func(t *testing.T) {
					t.Parallel()
					runParserTest(t, filePath)
				})
			}
			continue
		}
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}
		filePath := path.Join("examples", file.Name())
		t.Run(filePath, func(t *testing.T) {
			t.Parallel()
			runParserTest(t, filePath)
		})
	}
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
