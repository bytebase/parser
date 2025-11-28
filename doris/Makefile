all: build test

build:
	antlr -Dlanguage=Go -package doris -visitor -o . DorisLexer.g4 DorisParser.g4

test:
	go test -v -run TestDorisParser
