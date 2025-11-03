all: build test

build:
	antlr -Dlanguage=Go -package plsql -visitor -o . PlSqlLexer.g4 PlSqlParser.g4

test:
	go test -v -run TestPLSQLParser
