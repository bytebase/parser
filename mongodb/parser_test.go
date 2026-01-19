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
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.errors = append(l.errors, msg)
}

func (l *TestErrorListener) HasErrors() bool {
	return len(l.errors) > 0
}

func parseMongoShell(t *testing.T, input string) (antlr.Tree, *TestErrorListener, *TestErrorListener) {
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

func TestMongoShellParser(t *testing.T) {
	entries, err := os.ReadDir("examples")
	require.NoError(t, err)

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".js" {
			testFile(t, filepath.Join("examples", entry.Name()))
		}
	}
}

func TestShellCommands(t *testing.T) {
	tests := []string{
		"show dbs",
		"show databases",
		"show collections",
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestFindOperations(t *testing.T) {
	tests := []string{
		`db.users.find()`,
		`db.users.find({})`,
		`db.users.findOne()`,
		`db.users.findOne({})`,
		`db.users.find({ name: "alice" })`,
		`db.users.find({ age: { $gt: 25 } })`,
		`db.users.find({ age: { $gte: 18, $lt: 65 } })`,
		`db.users.find({ status: { $in: ["active", "pending"] } })`,
		`db.users.find({ $or: [{ name: "alice" }, { name: "bob" }] })`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestCursorModifiers(t *testing.T) {
	tests := []string{
		`db.users.find().sort({ age: -1 })`,
		`db.users.find().limit(10)`,
		`db.users.find().skip(5)`,
		`db.users.find().projection({ name: 1, age: 1 })`,
		`db.users.find().project({ name: 1, email: 1 })`,
		`db.users.find().sort({ age: -1 }).limit(10)`,
		`db.users.find().sort({ createdAt: -1 }).skip(20).limit(10)`,
		`db.users.find({ status: "active" }).sort({ name: 1 }).limit(100).skip(0)`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestCollectionAccess(t *testing.T) {
	tests := []string{
		`db.users.find()`,
		`db["users"].find()`,
		`db['users'].find()`,
		`db.getCollection("users").find()`,
		`db.getCollection('users').find()`,
		`db["user-logs"].find()`,
		`db.getCollection("my.collection").find()`,
		`db.getCollectionNames()`,
		`db.getCollectionInfos()`,
		`db.getCollectionInfos({ name: "users" })`,
		`db.getCollectionInfos({}, { nameOnly: true })`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestHelperFunctions(t *testing.T) {
	tests := []string{
		`db.users.find({ _id: ObjectId("507f1f77bcf86cd799439011") })`,
		`db.users.find({ _id: ObjectId() })`,
		`db.events.find({ createdAt: ISODate("2024-01-15T00:00:00.000Z") })`,
		`db.events.find({ createdAt: { $gt: ISODate() } })`,
		`db.events.find({ timestamp: Date() })`,
		`db.events.find({ timestamp: Date("2024-01-15") })`,
		`db.events.find({ timestamp: Date(1705276800000) })`,
		`db.sessions.find({ sessionId: UUID("550e8400-e29b-41d4-a716-446655440000") })`,
		`db.stats.find({ count: Long(9007199254740993) })`,
		`db.stats.find({ count: Long("9007199254740993") })`,
		`db.stats.find({ count: NumberLong(123456789012345) })`,
		`db.items.find({ quantity: Int32(100) })`,
		`db.items.find({ quantity: NumberInt(100) })`,
		`db.measurements.find({ value: Double(3.14159) })`,
		`db.financial.find({ amount: Decimal128("1234567890.123456789") })`,
		`db.financial.find({ amount: NumberDecimal("99.99") })`,
		`db.oplog.find({ ts: Timestamp(1627811580, 1) })`,
		`db.oplog.find({ ts: Timestamp({ t: 1627811580, i: 1 }) })`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestRegex(t *testing.T) {
	tests := []string{
		`db.users.find({ name: /alice/ })`,
		`db.users.find({ name: /^alice/i })`,
		`db.users.find({ email: /.*@example\.com$/ })`,
		`db.users.find({ name: RegExp("alice") })`,
		`db.users.find({ name: RegExp("^alice", "i") })`,
		`db.users.find({ name: RegExp("test", "gi") })`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestDocumentSyntax(t *testing.T) {
	tests := []string{
		// Unquoted keys
		`db.users.find({ name: "alice", age: 25 })`,
		// Quoted keys
		`db.users.find({ "name": "alice", "age": 25 })`,
		`db.users.find({ 'name': 'alice' })`,
		// Mixed
		`db.users.find({ name: "alice", "special-field": "value" })`,
		// Nested
		`db.users.find({ profile: { name: "test", active: true } })`,
		// Arrays
		`db.users.find({ tags: ["a", "b", "c"] })`,
		// Trailing commas
		`db.users.find({ name: "alice", age: 25, })`,
		`db.users.find({ tags: ["a", "b", "c",] })`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestLiterals(t *testing.T) {
	tests := []string{
		// Strings
		`db.users.find({ name: "alice" })`,
		`db.users.find({ name: 'alice' })`,
		// Numbers
		`db.users.find({ age: 25 })`,
		`db.users.find({ score: -10 })`,
		`db.users.find({ price: 19.99 })`,
		`db.users.find({ tiny: .001 })`,
		`db.users.find({ distance: 1.5e10 })`,
		`db.users.find({ small: 1e-6 })`,
		// Booleans
		`db.users.find({ active: true })`,
		`db.users.find({ deleted: false })`,
		// Null
		`db.users.find({ deletedAt: null })`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestComments(t *testing.T) {
	tests := []string{
		`// Line comment
db.users.find()`,
		`db.users.find() // inline comment`,
		`/* Block comment */ db.users.find()`,
		`db.users.find({ /* comment */ name: "alice" })`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
	}
}

func TestComplexQueries(t *testing.T) {
	tests := []string{
		// Complex filter with helpers
		`db.users.find({
			_id: ObjectId("507f1f77bcf86cd799439011"),
			createdAt: { $gt: ISODate("2024-01-01T00:00:00Z") },
			lastLogin: { $lt: Date() },
			sessionId: UUID("550e8400-e29b-41d4-a716-446655440000"),
			loginCount: NumberLong(1000)
		})`,
		// Multiple chained methods
		`db.users.find({ age: { $gt: 18 } }).sort({ lastName: 1, firstName: 1 }).skip(10).limit(20).projection({ firstName: 1, lastName: 1, email: 1 })`,
		// Multiple statements
		`show dbs
show collections
db.users.find()
db.users.find({ name: "alice" }).limit(10)`,
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			_, lexerErrors, parserErrors := parseMongoShell(t, tc)
			require.False(t, lexerErrors.HasErrors(), "Lexer errors: %v", lexerErrors.errors)
			require.False(t, parserErrors.HasErrors(), "Parser errors: %v", parserErrors.errors)
		})
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
			// Verify error message contains hint about 'new' keyword
			require.Contains(t, errorListener.Errors[0].Message, "new",
				"Error message should mention 'new' keyword")
		})
	}
}
