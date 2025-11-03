all: build test

build:
	antlr -Dlanguage=Go -package tsql -visitor -o . TSqlLexer.g4 TSqlParser.g4

test:
	go test -v -run TestTSQLParser
