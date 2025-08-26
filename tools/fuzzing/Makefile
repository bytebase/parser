BINARY_NAME=fuzzer
BUILD_DIR=bin
CMD_PATH=github.com/bytebase/parser/tools/fuzzing/cmd/fuzzer

.PHONY: all build test clean run help

all: build test

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)

# Run tests
test:
	@echo "Running tests..."
	go test -v github.com/bytebase/parser/tools/fuzzing/...

# Clean build artifacts  
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	go clean

# Run the fuzzer (requires arguments)
run:
	go run $(CMD_PATH) $(ARGS)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	cd ../.. && go mod tidy && go mod download

# Format code
fmt:
	@echo "Formatting code..."
	go fmt github.com/bytebase/parser/tools/fuzzing/...

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the fuzzer binary"
	@echo "  test     - Run all tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  run      - Run the fuzzer (use ARGS='--grammar postgresql --start-rule selectStmt')"
	@echo "  deps     - Install/update dependencies"
	@echo "  fmt      - Format all Go code"
	@echo "  lint     - Run golangci-lint"
	@echo "  help     - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make run ARGS='--grammar postgresql/PostgreSQLLexer.g4,postgresql/PostgreSQLParser.g4 --start-rule selectStmt --count 5'"
	@echo "  make run ARGS='--help'"