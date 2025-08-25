# ANTLR v4 Grammar Parser

A Go implementation to parse ANTLR v4 grammar files (`.g4` files) in this repository.

## Source

The lexer and parser grammars come from: https://github.com/antlr/grammars-v4/blob/master/antlr/antlr4

## Why Custom NextToken()?

We added `func (l *LexerAdaptor) NextToken() antlr.Token` in `lexer_adaptor.go` because:

- ANTLR grammar parsing requires context-sensitive lexing
- Need to convert `ID` tokens to `TOKEN_REF` (uppercase) or `RULE_REF` (lowercase)
- Go ANTLR doesn't automatically call `Emit()` like Java ANTLR does
- Go tokens are immutable, so we use a `TokenTypeWrapper` to override token types

## Why Sed Command in Makefile?

We added this sed command in the Makefile:
```bash
sed -i '' 's/l\.BaseLexer = antlr\.NewBaseLexer(input)/l.LexerAdaptor = *NewLexerAdaptor(input)/' antlrv4_lexer.go
```

Because:
- ANTLR code generation creates `l.BaseLexer = antlr.NewBaseLexer(input)` 
- We need `l.LexerAdaptor = *NewLexerAdaptor(input)` to use our custom lexer
- This automatically fixes the generated constructor after each regeneration

## Usage

```bash
make build  # Generate parser and apply fixes
make test   # Test all .g4 files in repository (should show 100% success)
make all    # Build and test
```