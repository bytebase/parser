package redshift_test

import (
	"os"
	"path"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/bytebase/parser/redshift"
	"github.com/stretchr/testify/require"
)

type CustomErrorListener struct {
	errors int
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

func TestRedshiftParser(t *testing.T) {
	redshiftDir := path.Join("examples", "redshift")
	// Check if redshift directory exists
	if _, err := os.Stat(redshiftDir); os.IsNotExist(err) {
		t.Skip("Redshift examples directory not found")
		return
	}

	examples, err := os.ReadDir(redshiftDir)
	if err != nil {
		t.Skip("Cannot read Redshift examples directory")
		return
	}

	for _, file := range examples {
		if file.IsDir() {
			continue
		}
		filePath := path.Join(redshiftDir, file.Name())
		t.Run(filePath, func(t *testing.T) {
			// read all the bytes from the file
			data, err := os.ReadFile(filePath)
			require.NoError(t, err)

			input := antlr.NewInputStream(string(data))

			lexer := redshift.NewRedshiftLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := redshift.NewRedshiftParser(stream)

			lexerErrors := &CustomErrorListener{}
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrors)

			parserErrors := &CustomErrorListener{}
			p.RemoveErrorListeners()
			p.AddErrorListener(parserErrors)

			p.BuildParseTrees = true

			tree := p.Root()

			require.Equal(t, 0, lexerErrors.errors)
			require.Equal(t, 0, parserErrors.errors)

			start := 0
			stop := stream.Size() - 1
			for i := 0; i < stream.Size(); i++ {
				if stream.Get(i).GetChannel() == antlr.TokenDefaultChannel {
					start = i
					break
				}
			}
			for i := stream.Size() - 1; i >= 0; i-- {
				if stream.Get(i).GetChannel() == antlr.TokenDefaultChannel && stream.Get(i).GetTokenType() != antlr.TokenEOF {
					stop = i
					break
				}
			}
			require.Equal(t, start, tree.GetStart().GetTokenIndex())
			require.Equal(t, stop, tree.GetStop().GetTokenIndex())
		})
	}
}
