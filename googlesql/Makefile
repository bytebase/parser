all: build test

build:
	antlr -Dlanguage=Go -package googlesql -visitor -o . GoogleSQLLexer.g4 GoogleSQLParser.g4

test:
	go test -v -run TestGoogleSQLParser
