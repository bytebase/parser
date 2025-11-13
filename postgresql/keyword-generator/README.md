# PostgreSQL Keyword Generator

Automatically generates PostgreSQL keyword definitions from the official PostgreSQL source code.

## Overview

This tool fetches keyword definitions from PostgreSQL's `kwlist.h` and generates:
- **PostgreSQLKeywords.g4**: Parser rules for keyword categories (reserved, unreserved, etc.)
- **PostgreSQLLexer.g4**: Updates the auto-generated keyword section with all 494 keyword tokens

## Usage

```bash
# In postgresql directory
make generate-keywords
```

This will:
1. Fetch keywords from PostgreSQL REL_18_STABLE branch
2. Generate PostgreSQLKeywords.g4 (parser rules)
3. Update PostgreSQLLexer.g4 (keyword section between markers)

## ANTLR Reserved Names

Some PostgreSQL keywords conflict with ANTLR reserved names. These are automatically renamed:

| PostgreSQL | ANTLR Token | Reason |
|------------|-------------|--------|
| SKIP | SKIP_P | SKIP is ANTLR-reserved for skip channel |

## Future Updates

To update to a new PostgreSQL version:

1. Edit `keyword-generator/main.go`:
   ```go
   const PostgreSQLVersion = "REL_19_STABLE"  // Change version
   ```

2. Regenerate keywords:
   ```bash
   make generate-keywords
   ```

3. Build and test:
   ```bash
   make build
   make test
   ```

## File Structure

```
postgresql/
├── keyword-generator/
│   ├── main.go                    # Keyword generator tool
│   └── README.md                  # This file
├── PostgreSQLKeywords.g4          # Generated parser rules
├── PostgreSQLLexer.g4             # Main lexer (with auto-generated keyword section)
└── PostgreSQLParser.g4            # Main parser (imports keywords)
```

## Benefits

✓ **Zero manual maintenance** - Keywords auto-synced with PostgreSQL source
✓ **Set operations** - Complete file regeneration, no patching
✓ **Version tracking** - Clear source and version in generated files
✓ **ANTLR compatibility** - Automatic handling of reserved names
✓ **Single source of truth** - All keywords defined in auto-generated sections
