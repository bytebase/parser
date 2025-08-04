# Redshift Parser Development Guide

## Project Overview

This repository is a Redshift parser built with ANTLR 4, forked from github.com/bytebase/postgresql-parser. Due to incompatibility issues with PostgreSQL, this separate repository was created to support Amazon Redshift-specific syntax and features.

## Architecture

### Core Components

1. **ANTLR Grammar Files**:
   - `RedshiftLexer.g4` - Tokenization rules for Redshift SQL
   - `RedshiftParser.g4` - Parser grammar with 200+ statement types
   - Generated Go files: `redshift_parser.go`, `redshift_lexer.go`, etc.

2. **Base Classes**:
   - `redshift_parser_base.go` - Engine-aware parser with PostgreSQL/Redshift support
   - `redshift_lexer_base.go` - Base lexer implementation
   - `string_stack.go` - Utility for string stack operations

3. **Supporting Files**:
   - `keywords.go` - 600+ PostgreSQL keywords with reserved status
   - `builtin_function.go` - Built-in function definitions
   - `build.sh` - ANTLR code generation script

### Engine Support

The parser supports multiple database engines:
- `EnginePostgreSQL` - Standard PostgreSQL syntax
- `EngineRedshift` - Amazon Redshift-specific syntax extensions

## Development Guidelines

### Code Conventions

1. **Follow existing patterns**: Always examine existing code before making changes
2. **Token/Rule/Name Convention**: Maintain consistency with current ANTLR grammar naming
3. **Engine-specific features**: Use engine detection for Redshift-specific syntax
4. **Error handling**: Implement proper error listeners and recovery mechanisms

### Testing Requirements

**CRITICAL**: Every change must include a related test case.

1. **Test Structure**:
   - Add SQL test files to the `examples/` directory
   - Use Go-based tests in `parser_test.go` and `engine_specific_test.go`
   - Tests automatically parse all SQL files in `examples/`

2. **Test Content Sources**:
   - Reference https://docs.aws.amazon.com/redshift/latest/dg/c_SQL_commands.html
   - Crawl syntax examples from AWS Redshift documentation
   - Use real-world SQL examples when possible

3. **Test Categories**:
   - DDL: CREATE, ALTER, DROP statements
   - DML: SELECT, INSERT, UPDATE, DELETE
   - Redshift-specific: IDENTITY columns, DISTKEY, SORTKEY, etc.
   - Advanced: Window functions, CTEs, JSON operations

### Adding New Features

1. **Grammar Changes**:
   ```bash
   # Edit RedshiftLexer.g4 or RedshiftParser.g4
   # Run build script to regenerate Go code
   make build
   ```

2. **Engine-Specific Logic**:
   - Use `GetEngine()` method to detect Redshift vs PostgreSQL
   - Implement conditional parsing for dialect-specific features
   - See `engine_specific_test.go` for examples

3. **Testing Process**:
   - Create SQL test files in `examples/`
   - Run tests: `go test -v`
   - Verify both parsing success and error handling

### Common Tasks

#### Adding Redshift-Specific Syntax

1. Identify the syntax difference from PostgreSQL
2. Update the appropriate grammar file (lexer or parser)
3. Add engine-specific logic if needed
4. Create test cases with AWS documentation examples
5. Verify tests pass for both engines

#### Adding New Keywords

1. Add to `keywords.go` with appropriate reserved status
2. Update lexer grammar if needed
3. Test keyword recognition in various contexts

#### Adding Built-in Functions

1. Add to `builtin_function.go` in appropriate category
2. Update parser rules if function has special syntax
3. Test function parsing and recognition

## Build and Test Commands

**IMPORTANT**: Always run `./build.sh` before running tests to generate the latest Go code from ANTLR grammars.

```bash
# Generate parser code from ANTLR grammars (REQUIRED before testing)
make build

# Run all tests
go test -v

# Run specific test
go test -run TestParser -v

# Run benchmarks
go test -bench=. -v
```

## References

- [AWS Redshift SQL Commands](https://docs.aws.amazon.com/redshift/latest/dg/c_SQL_commands.html)
- [ANTLR 4 Documentation](https://github.com/antlr/antlr4/blob/master/doc/index.md)
- [PostgreSQL Grammar Reference](https://github.com/tunnelvisionlabs/antlr4-postgresql)

## Contributing

1. Always add test cases for new features
2. Follow existing code patterns and conventions
3. Test against both PostgreSQL and Redshift engines
4. Use AWS documentation for accurate syntax examples
5. Ensure all tests pass before submitting changes