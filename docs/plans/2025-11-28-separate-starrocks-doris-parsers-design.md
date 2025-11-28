# Separate StarRocks and Doris Parsers

## Background

The current `doris/` directory actually contains the StarRocks SQL parser, not the Apache Doris parser. Now that the official Apache Doris parser is available, we should:
1. Rename the existing parser to accurately reflect it's StarRocks
2. Create a new Doris parser using the official Apache Doris grammar

## Part 1: Rename `doris/` to `starrocks/`

- Rename directory `doris/` to `starrocks/`
- Rename grammar files:
  - `DorisSQL.g4` → `StarRocksSQL.g4`
  - `DorisSQLLex.g4` → `StarRocksSQLLex.g4`
- Update package name in all Go files from `doris` to `starrocks`
- Update generated file names (e.g., `dorissql_parser.go` → `starrockssql_parser.go`)
- Update README.md references
- Regenerate parser with `make build`

## Part 2: Create new `doris/` with Apache Doris parser

Source: https://github.com/apache/doris/tree/master/fe/fe-core/src/main/antlr4/org/apache/doris/nereids

- Fetch grammar files from Apache Doris repo:
  - `DorisParser.g4` (main parser grammar)
  - `DorisLexer.g4` (lexer grammar)
  - Additional grammars if imports require them
- Create standard directory structure:
  ```
  doris/
  ├── DorisParser.g4
  ├── DorisLexer.g4
  ├── Makefile
  ├── README.md
  ├── parser_test.go
  └── examples/
      └── *.sql (one file per documentation page)
  ```
- Set up Makefile with `build` and `test` targets
- Create `parser_test.go` following repo conventions
- Generate Go parser code with `make build`
- Run tests to verify

## Part 3: Update GitHub Actions

In `.github/workflows/tests.yml`, update line 37 to add `starrocks`:

```
ALL_PARSERS="redshift postgresql cql snowflake tsql doris starrocks trino plsql googlesql mysql partiql tidb mariadb cosmosdb"
```

## Part 4: Gather test SQL from Doris documentation

- Source: https://doris.apache.org/docs/3.x/sql-manual/sql-statements/
- Scrape SQL examples from documentation pages
- Create one `.sql` file per documentation page in `doris/examples/`
- Cover statement types:
  - Data Query: SELECT, CTE, JOINs, subqueries
  - Data Manipulation: INSERT, UPDATE, DELETE
  - Data Definition: CREATE/ALTER/DROP TABLE, INDEX, VIEW
  - Other statements as documented
- Use these files as test cases in `parser_test.go`

## Notes

- External consumers of the current `doris` parser will need to migrate to `starrocks` (handled separately by user)
- Package naming follows repo convention: lowercase matching directory name
