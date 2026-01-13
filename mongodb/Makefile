all: build test

build:
	@echo "Building MongoDB Shell parser..."
	antlr -Dlanguage=Go -package mongodb -visitor -o . MongoShellLexer.g4 MongoShellParser.g4

test:
	go test -v -run TestMongoShellParser

.PHONY: all build test
