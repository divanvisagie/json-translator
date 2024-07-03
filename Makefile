# Define directories
CMD_DIR := cmd/json-translator
INTERNAL_DIRS := internal/editor internal/guesser internal/jsonstore internal/language internal/sourcebox internal/destinationbox internal/translation
PKG_DIRS := pkg/parser pkg/storage
DOCS_DIR := docs

# Define the output binary name
BINARY_NAME := json-translator

# Go command
GO := go

# Default target
.PHONY: all
all: build

# Build target to compile the application
build:
	$(GO) build -o $(BINARY_NAME) $(CMD_DIR)

# Test target to run all tests
.PHONY: test
test:
	$(GO) test ./...

# Format the code
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Update dependencies
.PHONY: deps
deps:
	$(GO) mod tidy

# Clean up compiled binary and other generated files
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	$(GO) clean -cache

# Run the application
.PHONY: run
run:
	$(GO) run $(CMD_DIR)

# Generate and view documentation
.PHONY: docs
docs:
	@echo "Generate and open docs in a preferred documentation generator/viewer"

# Help target to display available targets
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  all      - Build the application"
	@echo "  build    - Compile the application"
	@echo "  test     - Run all tests"
	@echo "  fmt      - Format the code"
	@echo "  deps     - Update dependencies"
	@echo "  clean    - Clean up compiled binary and generated files"
	@echo "  run      - Run the application"
	@echo "  docs     - Generate and view documentation"
	@echo "  help     - Display this help message"
