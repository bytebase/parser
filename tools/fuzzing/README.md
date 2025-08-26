# Grammar-Aware Fuzzing Tool

A fuzzing tool that generates valid SQL inputs from ANTLR v4 grammar files for parser testing.

## Quick Start

```bash
# Build the fuzzer
make build

# List available grammars
./bin/fuzzer --list-grammars

# Single combined grammar file
./bin/fuzzer --grammar combined.g4 --start-rule selectStmt --count 10

# Separate lexer and parser files
./bin/fuzzer --grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 10

# Run with custom parameters  
./bin/fuzzer --grammar cql/CqlLexer.g4,cql/CqlParser.g4 --start-rule expr --max-depth 3 --max-quantifier 8 --count 5
```

## Project Structure

```
tools/fuzzing/
├── cmd/fuzzer/          # CLI application entry point
│   └── main.go
├── internal/            # Private application packages
│   ├── config/          # Configuration management
│   └── generator/       # Core fuzzing logic
├── bin/                 # Built binaries (created by make build)
├── Makefile            # Build and development tasks
└── go.mod              # Go module definition
```

## CLI Options

| Flag | Description | Default |
|------|-------------|---------|
| `--grammar` | Grammar file(s): single file or comma-separated lexer,parser | - |
| `--start-rule` | Starting grammar rule (required) | - |
| `--count` | Number of queries to generate | 10 |
| `--max-depth` | Maximum recursion depth | 5 |
| `--optional-prob` | Probability of optional elements (0.0-1.0) | 0.5 |
| `--max-quantifier` | Maximum count for `*` and `+` quantifiers | 5 |
| `--min-quantifier` | Minimum count override | 0 |
| `--quantifier-count` | Fixed count for all quantifiers | 0 |
| `--output` | Output file path | stdout |
| `--seed` | Random seed for reproducible results | current time |

## Examples

### Basic Usage
```bash
# Single combined grammar file
./bin/fuzzer --grammar combined.g4 --start-rule selectStmt --count 10

# Separate lexer and parser files
./bin/fuzzer --grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 10

# Generate CQL expressions with limited depth
./bin/fuzzer --grammar cql/CqlLexer.g4,cql/CqlParser.g4 --start-rule expr --max-depth 3 --count 5
```

### Performance Testing
```bash
# Generate queries with exactly 100 columns
./bin/fuzzer --grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --quantifier-count 100 --count 5

# Generate deeply nested expressions  
./bin/fuzzer --grammar cql/CqlLexer.g4,cql/CqlParser.g4 --start-rule expr --max-depth 15 --count 10
```

### Output Control
```bash
# Save to file
./bin/fuzzer --grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 100 --output queries.sql

# Reproducible generation
./bin/fuzzer --grammar cql/CqlLexer.g4,cql/CqlParser.g4 --start-rule expr --seed 42 --count 10
```

## Development

### Build Commands
```bash
# From tools/fuzzing directory
make build    # Build binary to bin/fuzzer
make test     # Run all tests  
make clean    # Clean build artifacts
make fmt      # Format code
make deps     # Install/update dependencies (runs from repo root)

# From repository root
go build -o tools/fuzzing/bin/fuzzer github.com/bytebase/parser/tools/fuzzing/cmd/fuzzer
```

### Running During Development
```bash
# From tools/fuzzing directory
make run ARGS='--grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 5'
make run ARGS='--help'

# From repository root
go run github.com/bytebase/parser/tools/fuzzing/cmd/fuzzer --grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 5
```

## Monolithic Repository Structure

This tool uses the single `go.mod` file at the repository root:
- **Module**: `github.com/bytebase/parser`
- **Import path**: `github.com/bytebase/parser/tools/fuzzing/...`
- **Dependencies**: Shared with other tools in the repository

## Integration

This tool is designed to integrate with:
- Existing ANTLR v4 grammar parser at `tools/grammar/`
- All parser implementations in the repository (postgresql, cql, redshift, etc.)
- Shared CI/CD pipeline and testing infrastructure

**TODO**: Grammar parser integration and actual query generation logic.