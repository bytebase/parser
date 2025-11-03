all: build test

build:
	antlr -Dlanguage=Go -package cosmosdb -visitor -o . CosmosDBLexer.g4 CosmosDBParser.g4

test:
	go test -v -run TestCosmosDBParser
