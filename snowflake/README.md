The plsql-parser is a parser for Snowflake. It is based on the [ANTLR4](https://github.com/antlr/antlr4) and use the grammar from [antlr4-grammars-plsql](https://github.com/antlr/grammars-v4/tree/master/sql/snowflake).

## Build

Before build, you need to install the ANTLR4.

requirements:
- https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- https://github.com/antlr/antlr4/blob/master/doc/go-target.md

```bash
./build.sh
```

## Update grammar

### Manually change the grammar file in this project

1. run `./build.sh` to generate the parser code.

### From antlr4-grammars-plsql

1. Clone the `SnowflakeLexer.g4` and `SnowflakeParser.g4` grammar files from https://github.com/antlr/grammars-v4/tree/master/sql/snowflake.
1. run `./build.sh` to generate the parser code.

## Test the parser

Run `TestSnowflakeSQLParser` in `parser_test.go` to test the parser.

```bash
go test -v
```

## References

- ANTLR4 Getting Started https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- ANTLR4 Go Garget https://github.com/antlr/antlr4/blob/master/doc/go-target.md