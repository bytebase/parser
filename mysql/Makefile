all: build test

build:
	antlr -Dlanguage=Go -package mysql -visitor -o . MySQLLexer.g4 MySQLParser.g4

test:
	go test -v -run TestMySQLDBSQLParser
