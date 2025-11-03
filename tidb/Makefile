all: build test

build:
	antlr -Dlanguage=Go -package tidb -visitor -o . TiDBLexer.g4 TiDBParser.g4

test:
	go test -v -run TestTiDBSQLParser
