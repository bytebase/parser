# Redshift Parser

A comprehensive SQL parser for Amazon Redshift built with ANTLR 4, optimized for Redshift-specific syntax.

## Overview

This project is a Go-based SQL parser specifically designed for Amazon Redshift. It originated as a fork of the PostgreSQL parser but has been restructured to focus exclusively on Redshift's unique syntax requirements.

## Features

- **Complete SQL Support**: Parses 200+ SQL statement types including DDL, DML, and advanced constructs
- **Redshift-Specific Syntax**: Full support for Redshift extensions like `IDENTITY` columns, `DISTKEY`, `SORTKEY`, and more
- **Redshift-Optimized**: Parser optimized exclusively for Redshift syntax and features
- **Comprehensive Testing**: 200+ test cases covering real-world SQL scenarios
- **High Performance**: Optimized for production use with parser reuse and efficient error handling

## Installation

```bash
go get github.com/bytebase/redshift-parser
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/antlr4-go/antlr/v4"
    "github.com/bytebase/redshift-parser"
)

func main() {
    // Parse a Redshift-specific CREATE TABLE statement
    sql := `CREATE TABLE users (
        id INT IDENTITY(1,1),
        name VARCHAR(100),
        email VARCHAR(255) UNIQUE
    ) DISTKEY(id) SORTKEY(name);`
    
    // Create lexer and parser
    input := antlr.NewInputStream(sql)
    lexer := parser.NewRedshiftLexer(input)
    stream := antlr.NewCommonTokenStream(lexer, 0)
    p := parser.NewRedshiftParser(stream)
    
    // Parse the SQL
    tree := p.Root()
    
    fmt.Println("Successfully parsed Redshift SQL!")
}
```

## Supported SQL Features

### DDL (Data Definition Language)
- `CREATE TABLE` with Redshift-specific options (DISTKEY, SORTKEY, IDENTITY)
- `ALTER TABLE` with column modifications and constraints
- `CREATE INDEX` with various index types
- `CREATE VIEW` and materialized views
- `CREATE FUNCTION` and stored procedures

### DML (Data Manipulation Language)
- `SELECT` with complex joins, subqueries, and window functions
- `INSERT` with conflict resolution (`ON CONFLICT`)
- `UPDATE` with joins and CTEs
- `DELETE` with complex conditions
- `MERGE` statements

### Advanced Features
- Common Table Expressions (CTEs)
- Window functions and analytics
- JSON operations and path expressions
- Array operations
- Regular expressions
- Full-text search

## Redshift-Specific Features

The parser is optimized for Redshift's unique SQL extensions:

- **IDENTITY columns**: `CREATE TABLE t (id INT IDENTITY(1,1))`
- **Distribution keys**: `DISTKEY(column_name)`
- **Sort keys**: `SORTKEY(column_name)`
- **Redshift built-in functions**: Comprehensive support for Redshift-specific functions
- **Data types**: All Redshift-supported data types including extensions

## Development

### Prerequisites
- Go 1.21+
- ANTLR 4.13.2+

### Building from Source

1. **Clone the repository**:
```bash
git clone https://github.com/bytebase/redshift-parser.git
cd redshift-parser
```

2. **Generate parser code**:
```bash
make build
```

3. **Run tests**:
```bash
go test -v
```

### Project Structure

```
redshift-parser/
├── RedshiftLexer.g4              # ANTLR lexer grammar
├── RedshiftParser.g4             # ANTLR parser grammar
├── build.sh                      # Code generation script
├── redshift_lexer_base.go        # Base lexer implementation
├── redshift_parser_base.go       # Base parser implementation
├── keywords.go                   # 600+ SQL keywords
├── builtin_function.go           # Built-in function definitions
├── examples/                     # 200+ SQL test files
├── parser_test.go                # Main test suite
└── CLAUDE.md                     # Development guide
```

## Testing

The project includes comprehensive test coverage:

```bash
# Run all tests
go test -v

# Run specific test
go test -run TestRedshiftParser -v

# Run benchmarks
go test -bench=. -v

```

Test files are located in the `examples/` directory and cover:
- Basic SQL operations
- Complex queries with joins and subqueries
- Redshift-specific syntax
- Error handling scenarios
- Performance benchmarks

## Grammar Files

The parser is built using ANTLR 4 grammars:

- **RedshiftLexer.g4**: Tokenization rules for SQL keywords, operators, and literals
- **RedshiftParser.g4**: Grammar rules for SQL statement parsing

After modifying grammar files, regenerate the Go code:

```bash
make build
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for your changes
4. Ensure all tests pass
5. Update documentation as needed
6. Submit a pull request

### Development Guidelines

- Always run `./build.sh` before testing after grammar changes
- Add test cases for new SQL syntax support
- Follow existing code patterns and conventions
- Use AWS Redshift documentation for syntax reference
- Test against both PostgreSQL and Redshift engines

## License

This project is licensed under the MIT License. See the grammar files for additional license information from the original PostgreSQL grammar contributors.

## Acknowledgments

- Based on the PostgreSQL grammar from [Tunnel Vision Labs](https://github.com/tunnelvisionlabs/antlr4-postgresql)
- Forked from [Bytebase PostgreSQL Parser](https://github.com/bytebase/postgresql-parser)
- Built with [ANTLR 4](https://github.com/antlr/antlr4)

## Related Projects

- [Bytebase](https://github.com/bytebase/bytebase) - Database DevOps platform
- [PostgreSQL Parser](https://github.com/bytebase/postgresql-parser) - Original PostgreSQL parser
- [ANTLR 4](https://github.com/antlr/antlr4) - Parser generator toolkit

## Support

- [GitHub Issues](https://github.com/bytebase/redshift-parser/issues) - Bug reports and feature requests
- [AWS Redshift Documentation](https://docs.aws.amazon.com/redshift/latest/dg/c_SQL_commands.html) - SQL syntax reference
- [ANTLR Documentation](https://github.com/antlr/antlr4/blob/master/doc/index.md) - Grammar development guide