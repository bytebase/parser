# PostgreSQL Keyword Generator

Automatically generates PostgreSQL keyword definitions from the official PostgreSQL source code.

## Overview

This tool fetches keyword definitions from PostgreSQL's `kwlist.h` and generates:
- **PostgreSQLKeywords.g4**: Parser rules for keyword categories (reserved, unreserved, etc.)
- **PostgreSQLKeywordsLexer.g4**: Lexer token definitions for all 494 keywords

## Automated Workflow

```bash
# In postgresql directory
make generate-keywords
```

This will:
1. Fetch keywords from PostgreSQL REL_18_STABLE branch
2. Generate PostgreSQLKeywords.g4 (parser rules)
3. Generate PostgreSQLKeywordsLexer.g4 (lexer tokens)

## One-Time Setup

### Step 1: Clean Up Duplicates (First Time Only)

Remove duplicate keyword definitions from PostgreSQLLexer.g4:

```bash
cd keyword-generator
./cleanup_duplicates.sh
```

This creates a backup (`PostgreSQLLexer.g4.before_cleanup`) and removes all keyword token definitions that are now in PostgreSQLKeywordsLexer.g4.

### Step 2: Add Import Statement

Add this line to PostgreSQLLexer.g4 after the `lexer grammar` declaration:

```antlr
lexer grammar PostgreSQLLexer;

import PostgreSQLKeywordsLexer;
```

### Step 3: Verify Build

```bash
cd ..
make build
make test
```

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
│   ├── cleanup_duplicates.sh     # One-time cleanup script
│   └── README.md                  # This file
├── PostgreSQLKeywords.g4          # Generated parser rules
├── PostgreSQLKeywordsLexer.g4     # Generated lexer tokens
├── PostgreSQLLexer.g4             # Main lexer (imports keywords)
└── PostgreSQLParser.g4            # Main parser (imports keywords)
```

## Benefits

✓ **Zero manual maintenance** - Keywords auto-synced with PostgreSQL source
✓ **Set operations** - Complete file regeneration, no patching
✓ **Version tracking** - Clear source and version in generated files
✓ **ANTLR compatibility** - Automatic handling of reserved names
✓ **Single source of truth** - All keywords defined once in PostgreSQLKeywordsLexer.g4
