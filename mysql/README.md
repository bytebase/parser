The mysql-parser is a parser for MySQL. It is based on the [ANTLR4](https://github.com/antlr/antlr4) and use the grammar from [antlr4-grammars-mysql](https://github.com/antlr/grammars-v4/tree/master/sql/mysql).

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

### From antlr4-grammars-mysql

1. Clone the `MySQLLexer.g4` and `MySQLParser.g4` grammar files from https://github.com/mysql/mysql-workbench/tree/8.0/library/parsers/grammars.
1. Update the lexer and parser.
1. run `./build.sh` to generate the parser code.


### Updates from MySQL Workbench.

1. Removed the version checks https://github.com/bytebase/mysql-parser/commit/addcea9f8779ffe95159e334d9dd1ea7f93ed32d.
1. Used SQL Mode off by default, and removed set types https://github.com/bytebase/mysql-parser/commit/878f24d9fd313f60b3a10f977c489c0599ae2d49, https://github.com/bytebase/mysql-parser/commit/1062be07340ebc148c7119cd16e980792eeb1653.
1. Hacked dot identifier and removed the base class of lexer and parser https://github.com/bytebase/mysql-parser/commit/1d535c32a0be05d14ffc333f08f9a17d375ce922.
1. Follow the change logs below to update the grammar.

## Test the parser

Run `TestMySQLDBSQLParser` in `parser_test.go` to test the parser.

```bash
go test -v
```

## References

- ANTLR4 Getting Started https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- ANTLR4 Go Garget https://github.com/antlr/antlr4/blob/master/doc/go-target.md

## Change Logs

- 2023-06-05: Support 8.0 CREATE USER statement, specifically the `[COMMENT 'comment_string' | ATTRIBUTE 'json_object']` clause.