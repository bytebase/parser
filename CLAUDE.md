# Parser Development Guide

This project contains multiple SQL parsers built with ANTLR 4 for different database dialects.

## Parser Structure

Each parser is organized in its own directory containing:
- ANTLR grammar files (`.g4`)
- Generated Go parser/lexer files
- Test files and examples
- Individual `Makefile` for build/test commands

## Generating Parsers

**IMPORTANT**: Always generate the parser after completing any grammar changes.

```bash
cd <parser-directory>/
make build
```

The `make build` command:
1. Runs ANTLR to generate Go code from `.g4` grammar files
2. Creates lexer, parser, listener, and visitor files
3. Must be run before testing after any grammar modifications

## Running Tests

**IMPORTANT**: Always run golang tests after generating the parser.

```bash
cd <parser-directory>/
make test
```

## Development Workflow

1. **Make grammar changes** to `.g4` files
2. **Generate parser** using `make build` in the specific parser directory
3. **Run tests** using `make test` to verify changes
4. **Add test cases** in the `examples/` directory for new features

## Notes

- Each parser directory has its own `Makefile` with `build` and `test` targets
- The `make all` target in each directory runs both build and test
- Always regenerate parsers after modifying grammar files
- Test files are located in each parser's `examples/` directory