all: build test

build:
	antlr -Dlanguage=Go -package partiql -visitor -o . PartiQLLexer.g4 PartiQLParser.g4

test:
	go test -v -run TestSimpleParse
