package mongodb_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/bytebase/parser/mongodb"
	"github.com/stretchr/testify/require"
)

type TestErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func NewTestErrorListener() *TestErrorListener {
	return &TestErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		errors:               make([]string, 0),
	}
}

func (l *TestErrorListener) SyntaxError(
	_ antlr.Recognizer,
	_ any,
	_, _ int,
	msg string,
	_ antlr.RecognitionException,
) {
	l.errors = append(l.errors, msg)
}

func (l *TestErrorListener) HasErrors() bool {
	return len(l.errors) > 0
}

func parseMongoShell(_ *testing.T, input string) (antlr.Tree, *TestErrorListener, *TestErrorListener) {
	is := antlr.NewInputStream(input)
	lexer := mongodb.NewMongoShellLexer(is)

	lexerErrors := NewTestErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := mongodb.NewMongoShellParser(stream)

	parserErrors := NewTestErrorListener()
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	tree := parser.Program()

	return tree, lexerErrors, parserErrors
}

func testFile(t *testing.T, filePath string) {
	t.Run(filepath.Base(filePath), func(t *testing.T) {
		t.Parallel()
		data, err := os.ReadFile(filePath)
		require.NoError(t, err)

		_, lexerErrors, parserErrors := parseMongoShell(t, string(data))
		require.False(t, lexerErrors.HasErrors(), "Lexer errors in %s: %v", filePath, lexerErrors.errors)
		require.False(t, parserErrors.HasErrors(), "Parser errors in %s: %v", filePath, parserErrors.errors)
	})
}

// TestMongoShellParser runs all .js example files as parser tests.
// Each .js file in the examples/ directory tests a specific feature:
//   - collection-find.js: db.collection.find() with filters, operators, cursor modifiers
//   - collection-findOne.js: db.collection.findOne() operations
//   - collection-countDocuments.js: db.collection.countDocuments() operations
//   - collection-estimatedDocumentCount.js: db.collection.estimatedDocumentCount() operations
//   - collection-distinct.js: db.collection.distinct() operations
//   - collection-aggregate.js: db.collection.aggregate() pipelines
//   - collection-getIndexes.js: db.collection.getIndexes() operations
//   - shell_commands.js: show dbs, show databases, show collections
//   - helper_functions.js: ObjectId(), ISODate(), UUID(), NumberLong(), etc.
//   - document_syntax.js: Document syntax with unquoted keys, trailing commas
//   - literals.js: String, number, boolean, null literals
//   - regex.js: Regex literals and RegExp() constructor
//   - comments.js: Line and block comments
func TestMongoShellParser(t *testing.T) {
	entries, err := os.ReadDir("examples")
	require.NoError(t, err)

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".js" {
			testFile(t, filepath.Join("examples", entry.Name()))
		}
	}
}

func TestErrorPositions(t *testing.T) {
	// Test that error positions are reported correctly
	input := `db.users.find({ name: })`

	is := antlr.NewInputStream(input)
	lexer := mongodb.NewMongoShellLexer(is)

	errorListener := mongodb.NewMongoShellErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := mongodb.NewMongoShellParser(stream)

	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	parser.BuildParseTrees = true
	_ = parser.Program()

	require.True(t, errorListener.HasErrors(), "Expected parse errors for invalid input")
	require.NotEmpty(t, errorListener.Errors)
	// Verify error has position info
	require.Greater(t, errorListener.Errors[0].Line, 0)
}

func TestNewKeywordErrorMessage(t *testing.T) {
	// Test that 'new' keyword produces helpful error message
	tests := []string{
		`db.users.find({ _id: new ObjectId("507f1f77bcf86cd799439011") })`,
		`db.events.find({ timestamp: new Date() })`,
		`db.users.find({ name: new RegExp("test") })`,
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			is := antlr.NewInputStream(input)
			lexer := mongodb.NewMongoShellLexer(is)

			errorListener := mongodb.NewMongoShellErrorListener()
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(errorListener)

			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			parser := mongodb.NewMongoShellParser(stream)

			parser.RemoveErrorListeners()
			parser.AddErrorListener(errorListener)

			parser.BuildParseTrees = true
			_ = parser.Program()

			require.True(t, errorListener.HasErrors(), "Expected parse errors for 'new' keyword usage")
			require.NotEmpty(t, errorListener.Errors)
			// Verify error message provides helpful guidance about 'new' keyword
			require.Contains(t, errorListener.Errors[0].Message, "'new' keyword is not supported",
				"Error message should provide helpful guidance about 'new' keyword")
		})
	}
}
