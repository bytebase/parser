The doris-parser is a parser for DORIS SQL. It is based on the [ANTLR4](https://github.com/antlr/antlr4).

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

## Test the parser

Run `TestDorisSQLDBSQLParser` in `parser_test.go` to test the parser.

```bash
go test -v
```

## References

- ANTLR4 Getting Started https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- ANTLR4 Go Garget https://github.com/antlr/antlr4/blob/master/doc/go-target.md