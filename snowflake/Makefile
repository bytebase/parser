all: build test

build:
	python3 build_id_contains_non_reserved_keywords.py && antlr -Dlanguage=Go -package snowflake -visitor -o . SnowflakeLexer.g4 SnowflakeParser.g4

test:
	go test -v -run TestSnowflakeSQLParser
