# Makefile for HTMX Login & Registration App

# Variables
BINARY_NAME=htmx-login
GO=go
PORT=8080

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@$(GO) build -o $(BINARY_NAME) .

# Run the application
.PHONY: run
run: build
	@echo "Starting server on http://localhost:$(PORT)"
	@./$(BINARY_NAME)

# Run without building
.PHONY: dev
dev:
	@echo "Starting development server on http://localhost:$(PORT)"
	@$(GO) run main.go

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@$(GO) clean

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@$(GO) fmt ./...

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@$(GO) test -v ./...

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@$(GO) mod tidy

# Help command
.PHONY: help
help:
	@echo "HTMX Login & Registration App Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Build and run the application"
	@echo "  make dev      - Run without building (for development)"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make fmt      - Format code"
	@echo "  make test     - Run tests"
	@echo "  make deps     - Install dependencies"
	@echo "  make help     - Show this help message"
