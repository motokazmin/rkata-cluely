.PHONY: help build run clean test fmt lint

# Variables
BINARY_NAME=cluely.exe
BIN_DIR=bin
CONFIG_FILE=configs/default.toml
OUTPUT_BINARY=$(BIN_DIR)/$(BINARY_NAME)

# Help
help:
	@echo "Cluely MVP - Available targets:"
	@echo ""
	@echo "  make build       - Compile the application to bin/cluely.exe"
	@echo "  make run         - Build and run the application"
	@echo "  make clean       - Remove compiled binaries and cache"
	@echo "  make fmt         - Format Go code"
	@echo "  make lint        - Run go vet"
	@echo "  make test        - Run tests (when available)"
	@echo ""
	@echo "Example:"
	@echo "  make build       - Builds binary"
	@echo "  bin/cluely.exe   - Run the compiled binary"

# Build
build: clean
	@echo "Building Cluely MVP..."
	@go build -o $(OUTPUT_BINARY) ./cmd/cluely
	@echo "Copying configuration..."
	@go run scripts/copy_config.go
	@echo "Build complete: $(OUTPUT_BINARY)"
	@echo ""
	@echo "To run the application:"
	@echo "   $(OUTPUT_BINARY)"
	@echo ""
	@echo "Then open: http://localhost:8080"

# Run
run: build
	@echo ""
	@echo "Starting Cluely MVP..."
	@echo ""
	@$(OUTPUT_BINARY)

# Run binary directly (if already built)
run-binary:
	@$(OUTPUT_BINARY)

# Clean
clean:
	@echo "Cleaning build artifacts..."
	@go clean
	@rm -rf $(BIN_DIR)
	@echo "Clean complete"

# Format
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...
	@echo "Format complete"

# Lint
lint:
	@echo "Running linter..."
	@go vet ./...
	@echo "Lint complete"

# Test
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests complete"

# Dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies ready"

# Dev setup
dev-setup: clean deps
	@echo "Development setup complete"
	@echo "Use 'make build' to compile"
	@echo "Use 'make run' to build and run"
	@echo "Use 'make fmt' to format code"
	@echo "Use 'make lint' to check code"

# Default target
.DEFAULT_GOAL := help
