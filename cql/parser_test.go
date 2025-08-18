package cql_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/antlr4-go/antlr/v4"
	cqlparser "github.com/bytebase/parser/cql"
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

func TestCQLParser(t *testing.T) {
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

			lexer := cqlparser.NewCqlLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := cqlparser.NewCqlParser(stream)

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
				if stream.Get(i).GetChannel() == antlr.TokenDefaultChannel {
					stop = i
					break
				}
			}

			require.NotNil(t, tree)
			require.NotEqual(t, 0, stop-start)
		})
	}
}

func TestParseCQL(t *testing.T) {
	type args struct {
		statement string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "CREATE KEYSPACE",
			args: args{
				statement: `CREATE KEYSPACE test_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3} AND durable_writes = true;`,
			},
		},
		{
			name: "CREATE TABLE",
			args: args{
				statement: `CREATE TABLE users (id UUID PRIMARY KEY, username text, email text, created_at timestamp);`,
			},
		},
		{
			name: "INSERT",
			args: args{
				statement: `INSERT INTO users (id, username, email, created_at) VALUES (uuid(), 'john_doe', 'john@example.com', toTimestamp(now()));`,
			},
		},
		{
			name: "SELECT",
			args: args{
				statement: `SELECT * FROM users WHERE id = 123e4567-e89b-12d3-a456-426614174000;`,
			},
		},
		{
			name: "UPDATE",
			args: args{
				statement: `UPDATE users SET email = 'newemail@example.com' WHERE id = 123e4567-e89b-12d3-a456-426614174000;`,
			},
		},
		{
			name: "DELETE",
			args: args{
				statement: `DELETE FROM users WHERE id = 123e4567-e89b-12d3-a456-426614174000;`,
			},
		},
		{
			name: "CREATE INDEX",
			args: args{
				statement: `CREATE INDEX ON users (email);`,
			},
		},
		{
			name: "ALTER TABLE",
			args: args{
				statement: `ALTER TABLE users ADD phone text;`,
			},
		},
		{
			name: "DROP TABLE",
			args: args{
				statement: `DROP TABLE IF EXISTS users;`,
			},
		},
		{
			name: "BATCH",
			args: args{
				statement: `BEGIN BATCH 
					INSERT INTO users (id, username) VALUES (uuid(), 'user1');
					INSERT INTO users (id, username) VALUES (uuid(), 'user2');
				APPLY BATCH;`,
			},
		},
	}

	start := time.Now()
	defer func() {
		t.Logf("Time elapsed: %v", time.Since(start))
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree, err := cqlparser.ParseCQL(tt.args.statement)
			require.NoError(t, err)
			require.NotNil(t, tree)
		})
	}
}

func TestCQLParserPerformance(t *testing.T) {
	statement := `CREATE TABLE users (id UUID PRIMARY KEY, username text, email text, created_at timestamp);`
	
	iterations := 1000
	start := time.Now()
	
	for i := 0; i < iterations; i++ {
		tree, err := cqlparser.ParseCQL(statement)
		require.NoError(t, err)
		require.NotNil(t, tree)
	}
	
	elapsed := time.Since(start)
	perIteration := elapsed / time.Duration(iterations)
	
	t.Logf("Parsed %d statements in %v", iterations, elapsed)
	t.Logf("Average time per parse: %v", perIteration)
	
	// Ensure parsing is reasonably fast (< 10ms per statement)
	require.Less(t, perIteration, 10*time.Millisecond, 
		fmt.Sprintf("Parsing too slow: %v per statement", perIteration))
}