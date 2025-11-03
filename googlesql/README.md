# google-sql-parser

Google SQL parser based on ANTLR4.

## Usage

### Build

```shell
make build
```

### Test

```shell
make test
```

## References

- [Query Syntax](https://cloud.google.com/bigquery/docs/reference/standard-sql/query-syntax), is the reference for the query(SELECT) syntax.

- [zetasql Yacc parser](https://github.com/google/zetasql/blob/master/zetasql/parser/bison_parser.y), is the reference for the syntax detail missing in the document like expression.
