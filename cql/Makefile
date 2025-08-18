all: build test

build:
	antlr -Dlanguage=Go -package cql -visitor -o . CqlLexer.g4 CqlParser.g4

test: 
	go test -v -run TestCqlParser

clean:
	rm -f *.go *.tokens *.interp