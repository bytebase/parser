all: build test

build:
	antlr -Dlanguage=Go -package mariadb -visitor -o . MariaDBLexer.g4 MariaDBParser.g4

test:
	go test -v -run TestMariaDBSQLParser
