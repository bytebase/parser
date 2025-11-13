# Regression Tests

This directory contains regression tests for specific bugs that were fixed in the parser.

## Test Files

### `function_names_as_identifiers.sql`
**Issue**: After removing `builtin_function_name` tokens, mathematical and string function names should be usable as regular identifiers.

**Tests**:
- Using function names (exp, div, floor, mod, power, sqrt, log, reverse) as column names
- Using function names as table names
- SELECT statements with these identifiers

**Related PR**: #40 - Remove non-standard builtin_function_name tokens

### `json_typecast.sql`
**Issue**: Missing `jsontype` rule caused `::json` and `::jsonb` typecasts to fail with "mismatched input ';' expecting '%'"

**Tests**:
- Basic JSON/JSONB typecast syntax
- JSON with unicode escape sequences
- JSON with surrogate pairs
- JSON type in CREATE TABLE
- JSON typecast with -> operator

**Related PR**: #40 - Add jsontype rule for proper JSON type support

### `complex_joins.sql`
**Issue**: Complex multi-table JOINs with multiple LEFT JOINs and ORDER BY clauses

**Tests**:
- Multiple LEFT JOIN statements
- Column aliases with AS keyword
- WHERE clause filtering
- ORDER BY with multiple columns and mixed ASC/DESC

**Related PR**: #40 - General regression test for complex query patterns

## Running Tests

These tests are automatically included when running the full test suite:

```bash
cd postgresql
make test
```

To run only regression tests:

```bash
go test -run "TestPostgreSQLParser/examples/regression" -v
```

## Adding New Regression Tests

When fixing a parser bug:

1. Create a new `.sql` file in this directory
2. Add a comment at the top describing the issue and fix
3. Include minimal test cases that reproduce the bug
4. The test will be automatically picked up by the test runner
