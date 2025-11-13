all: build test

# Generate keyword definitions from PostgreSQL official source
generate-keywords:
	@echo "Generating PostgreSQL keyword definitions from REL_18_STABLE..."
	cd keyword-generator && go run main.go
	@echo ""

# Build the parser (depends on keyword generation)
build: generate-keywords
	@echo "Building PostgreSQL parser..."
	antlr -Dlanguage=Go -package postgresql -visitor -o . PostgreSQLLexer.g4 PostgreSQLParser.g4

test:
	go test -v -run TestPostgreSQLParser

.PHONY: all build test generate-keywords