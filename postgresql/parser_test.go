package postgresql_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/antlr4-go/antlr/v4"
	pgparser "github.com/bytebase/parser/postgresql"
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

func TestPostgreSQLParser(t *testing.T) {
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	for _, file := range examples {
		filePath := path.Join("examples", file.Name())
		t.Run(filePath, func(t *testing.T) {
			t.Parallel()
			// read all the bytes from the file
			data, err := os.ReadFile(filePath)
			require.NoError(t, err)

			input := antlr.NewInputStream(string(data))

			lexer := pgparser.NewPostgreSQLLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := pgparser.NewPostgreSQLParser(stream)

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
			// require.Equal(t, string(data), stream.GetTextFromTokens(stream.Get(0), stream.Get(stream.Size()-1)))
		})
	}
}

type ParseService struct {
	lexer  *pgparser.PostgreSQLLexer
	parser *pgparser.PostgreSQLParser
	tokens antlr.TokenStream
	input  antlr.CharStream
	errors int
}

func NewParseService() *ParseService {
	s := &ParseService{
		input: antlr.NewInputStream(""),
	}
	s.lexer = pgparser.NewPostgreSQLLexer(s.input)
	s.tokens = antlr.NewCommonTokenStream(s.lexer, 0)
	s.parser = pgparser.NewPostgreSQLParser(s.tokens)
	return s
}

func (s *ParseService) Parse(script string, buildTree bool) antlr.RuleContext {
	s.errors = 0
	s.input = antlr.NewInputStream(script)
	s.lexer.Reset()
	s.lexer.SetInputStream(s.input)
	s.tokens.SetTokenSource(s.lexer)
	// set token stream and reset parser
	s.parser.SetTokenStream(s.tokens)

	lexerErrors := &CustomErrorListener{}
	s.lexer.RemoveErrorListeners()
	s.lexer.AddErrorListener(lexerErrors)

	parserErrors := &CustomErrorListener{}
	s.parser.RemoveErrorListeners()
	s.parser.AddErrorListener(parserErrors)

	// Set the parser's error strategy to bail on the first error encountered
	s.parser.SetErrorHandler(antlr.NewBailErrorStrategy())

	s.parser.BuildParseTrees = buildTree

	tree := s.parser.Root()
	s.errors += lexerErrors.errors
	s.errors += parserErrors.errors
	return tree
}

func TestBenchmarkParser(t *testing.T) {
	s := NewParseService()
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	var files []string

	for i, file := range examples {
		if i > 20 {
			break
		}
		filePath := path.Join("examples", file.Name())
		data, err := os.ReadFile(filePath)
		require.NoError(t, err)

		files = append(files, string(data))
	}

	fmt.Printf("Total files: %d\n", len(files))

	newParserEachFileStartTime := time.Now()

	for _, file := range files {
		input := antlr.NewInputStream(file)

		lexer := pgparser.NewPostgreSQLLexer(input)

		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := pgparser.NewPostgreSQLParser(stream)

		lexerErrors := &CustomErrorListener{}
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexerErrors)

		parserErrors := &CustomErrorListener{}
		p.RemoveErrorListeners()
		p.AddErrorListener(parserErrors)

		p.BuildParseTrees = true

		_ = p.Root()

		require.Equal(t, 0, lexerErrors.errors)
		require.Equal(t, 0, parserErrors.errors)

	}

	fmt.Printf("New parser each file took %s\n", time.Since(newParserEachFileStartTime).String())

	coldStartTime := time.Now()

	for _, file := range files {
		_ = s.Parse(string(file), true)
		require.Zero(t, s.errors)
	}

	fmt.Printf("Parse with cold start took %s\n", time.Since(coldStartTime).String())

	for i := 0; i < 5; i++ {
		for _, file := range files {
			_ = s.Parse(string(file), true)
			require.Zero(t, s.errors)
		}
	}
	fmt.Println("Finish warm up")

	warmStartTime := time.Now()
	for _, file := range files {
		_ = s.Parse(string(file), true)
		require.Zero(t, s.errors)
	}

	fmt.Printf("Parse with warm start took %s\n", time.Since(warmStartTime).String())

	withoutParseTreeStartTime := time.Now()

	for _, file := range files {
		_ = s.Parse(string(file), false)
		require.Zero(t, s.errors)
	}

	fmt.Printf("Parse without parse tree took %s\n", time.Since(withoutParseTreeStartTime).String())
}
