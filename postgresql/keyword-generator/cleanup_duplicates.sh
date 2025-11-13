#!/bin/bash
# One-time script to remove duplicate keyword definitions from PostgreSQLLexer.g4
# based on keywords defined in PostgreSQLKeywordsLexer.g4

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
LEXER_FILE="$SCRIPT_DIR/../PostgreSQLLexer.g4"
KEYWORDS_FILE="$SCRIPT_DIR/../PostgreSQLKeywordsLexer.g4"

echo "Extracting keyword literals from PostgreSQLKeywordsLexer.g4..."

# Extract all string literals from PostgreSQLKeywordsLexer.g4 (e.g., 'SKIP', 'ABORT')
KEYWORDS=$(grep -E "^\s+:\s+'[^']+" "$KEYWORDS_FILE" | sed -E "s/.*'([^']+)'.*/\1/" | sort -u)

echo "Found $(echo "$KEYWORDS" | wc -l) unique keywords"
echo ""
echo "Removing duplicate keyword definitions from PostgreSQLLexer.g4..."

# Create backup
cp "$LEXER_FILE" "$LEXER_FILE.before_cleanup"

# For each keyword, find and remove the token definition
REMOVED=0
for keyword in $KEYWORDS; do
    # Search for token definition with this literal
    # Pattern: TOKEN_NAME\n   : 'KEYWORD'\n   ;
    if grep -q ": '$keyword'" "$LEXER_FILE"; then
        echo "  Removing: '$keyword'"
        # Use perl for multi-line matching and removal
        perl -i -0pe "s/^[A-Z_][A-Z_0-9]*\n\s+:\s+'$keyword'\n\s+;\n\n?//gm" "$LEXER_FILE"
        ((REMOVED++))
    fi
done

echo ""
echo "✓ Removed $REMOVED duplicate keyword definitions"
echo "✓ Backup saved to: $LEXER_FILE.before_cleanup"
