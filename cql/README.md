# CQL Parser

ANTLR4-based parser for Cassandra Query Language (CQL).

## Overview

This parser supports CQL3 syntax as used by:
- Apache Cassandra
- ScyllaDB  
- Amazon Keyspaces
- DataStax Astra

## Usage

```go
import "github.com/bytebase/parser/cql"

// Parse a CQL statement
tree, err := cql.ParseCQL("SELECT * FROM users WHERE id = 123;")
if err != nil {
    // Handle parsing error
    log.Fatal(err)
}
```

## Grammar Source

The grammar files are from the official ANTLR grammars repository:
- https://github.com/antlr/grammars-v4/tree/master/cql3

## Building

```bash
make build  # Generate Go files from ANTLR grammars
make test   # Run tests
make clean  # Clean generated files
```

## Testing

The parser is tested against 38 official CQL example files covering:
- DDL statements (CREATE, ALTER, DROP for keyspaces, tables, indexes, etc.)
- DML statements (INSERT, UPDATE, DELETE, SELECT)
- DCL statements (GRANT, REVOKE, LIST PERMISSIONS)
- Batch operations
- User-defined types, functions, and aggregates
- Materialized views

## Performance

Average parsing time: ~14 microseconds per statement