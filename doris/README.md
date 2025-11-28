# Apache Doris SQL Parser

An ANTLR4-based SQL parser for Apache Doris, built using the official Apache Doris grammar files.

## About

This parser is generated from the official Apache Doris ANTLR4 grammar files, providing accurate SQL parsing capabilities for Apache Doris SQL dialect.

## Grammar Source

The grammar files are sourced from the official Apache Doris repository:
- [DorisLexer.g4](https://github.com/apache/doris/blob/master/fe/fe-core/src/main/antlr4/org/apache/doris/nereids/DorisLexer.g4)
- [DorisParser.g4](https://github.com/apache/doris/blob/master/fe/fe-core/src/main/antlr4/org/apache/doris/nereids/DorisParser.g4)

## Build

Generate the parser from the ANTLR4 grammar files:

```bash
make build
```

This command uses ANTLR4 to generate the Go lexer, parser, listener, and visitor files.

## Test

Run the test suite:

```bash
make test
```

## Development

After modifying the grammar files, always regenerate the parser:

```bash
make build
make test
```

## References

- [Apache Doris](https://doris.apache.org/)
- [Apache Doris GitHub Repository](https://github.com/apache/doris)
- [ANTLR4](https://www.antlr.org/)
- [ANTLR4 Documentation](https://github.com/antlr/antlr4/blob/master/doc/index.md)
