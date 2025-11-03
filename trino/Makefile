all: build test

build:
	antlr -Dlanguage=Go -package trino -visitor -o . TrinoLexer.g4 TrinoParser.g4

test:
	go test -v -run TestTrinoParser
