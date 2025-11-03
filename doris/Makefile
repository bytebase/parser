all: build test

build:
	antlr -Dlanguage=Go -package doris -visitor -o . DorisSQLLex.g4 DorisSQL.g4

test:
	go test -v -run TestDorisSQLParser
