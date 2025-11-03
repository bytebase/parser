all: build test

build:
	antlr -Dlanguage=Go -package bq -visitor -o . BigQueryLexer.g4 BigQueryParser.g4

test:
	go test -v -run TestBQParser
