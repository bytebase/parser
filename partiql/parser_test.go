package partiql_test

import (
	"os"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	partiql "github.com/bytebase/parser/partiql"
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

func TestSimpleParse(t *testing.T) {
	filepath := "simple.sql"
	content, err := os.ReadFile(filepath)
	require.NoError(t, err)
	dataString := strings.TrimRight(string(content), " \t\r\n;") + "\n;"

	input := antlr.NewInputStream(string(dataString))
	lexer := partiql.NewPartiQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := partiql.NewPartiQLParserParser(stream)
	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	p.BuildParseTrees = true
	tree := p.Script()
	require.Equal(t, 0, lexerErrors.errors)
	require.Equal(t, 0, parserErrors.errors)

	require.Equal(t, dataString, stream.GetTextFromRuleContext(tree))
}

// func TestParsePartiQL(t *testing.T) {
// 	type lookupRootDirectory struct {
// 		rootDirectory string
// 		success       bool
// 	}
// 	// Reference the [README](https://github.com/partiql/partiql-tests/tree/main) to figure out how to use the test data.
// 	lookupRootDirectories := []*lookupRootDirectory{
// 		{"partiql-tests/partiql-tests-data/fail/syntax", false},
// 		{"partiql-tests/partiql-tests-data/success/syntax", true},
// 	}
// 	for _, rootDirectory := range lookupRootDirectories {
// 		runTestParsePartiQLInDirectory(t, rootDirectory.rootDirectory, rootDirectory.success)
// 	}
// }

// type IonData struct {
// 	Tests []TestCase `ion:""`
// }

// type TestCase struct {
// 	Name      string `ion:"name"`
// 	Statement string `ion:"statement"`
// }

// func runTestParsePartiQLInDirectory(t *testing.T, directory string, success bool) {
// 	// Recurse through the directory and parse all the filepaths.
// 	filepaths, err := GetIonFilesInDirectory(directory)
// 	require.NoError(t, err)
// 	for _, fp := range filepaths {
// 		content, err := os.ReadFile(fp)
// 		require.NoErrorf(t, err, "Failed to read file %s", fp)
// 		ionData := IonData{}
// 		err = ion.Unmarshal(content, &ionData)
// 		require.NoErrorf(t, err, "Failed to unmarshal ion data from file %s", fp)
// 		t.Logf("Parsed %d test cases from file %s", len(ionData.Tests), fp)
// 	}

// }

// func GetIonFilesInDirectory(directory string) ([]string, error) {
// 	var files []string
// 	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
// 		if info.IsDir() {
// 			return nil
// 		}
// 		if filepath.Ext(path) != ".ion" {
// 			return nil
// 		}
// 		files = append(files, path)
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return files, nil
// }
