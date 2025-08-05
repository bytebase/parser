all: build test

build:
	antlr -Dlanguage=Go -package redshift -visitor -o . RedshiftLexer.g4 RedshiftParser.g4

test: 
	go test -v -run TestRedshiftParser