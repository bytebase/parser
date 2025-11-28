all: build test

build:
	antlr -Dlanguage=Go -package starrocks -visitor -o . StarRocksSQLLex.g4 StarRocksSQL.g4

test:
	go test -v -run TestStarRocksSQLParser
