.PHONY: all test clean help

all: test

# Run tests
test:
	@echo "Running tests..."
	go test -v github.com/bytebase/parser/tools/fuzzing/...

# Clean build artifacts  
clean:
	@echo "Cleaning..."
	go clean

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
	@echo "  test     - Run all tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  deps     - Install/update dependencies"
	@echo "  fmt      - Format all Go code"
	@echo "  lint     - Run golangci-lint"
	@echo "  help     - Show this help message"