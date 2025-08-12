all: build test

build:
	antlr -Dlanguage=Go -package postgresql -visitor -o . PostgreSQLLexer.g4 PostgreSQLParser.g4

test: 
	go test -v -run TestPostgreSQLParser