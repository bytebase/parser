all: build test

build:
	antlr -Dlanguage=Go -package grammar -visitor -o . ANTLRv4Lexer.g4 ANTLRv4Parser.g4
	@echo "Applying LexerAdaptor fix..."
	sed -i '' 's/l\.BaseLexer = antlr\.NewBaseLexer(input)/l.LexerAdaptor = *NewLexerAdaptor(input)/' antlrv4_lexer.go

test:
	go test -v .
