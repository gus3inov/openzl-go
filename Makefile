# Makefile for OpenZL Go bindings
# Provides convenient targets for building, testing, and development

.PHONY: help deps build test clean install-deps build-openzl run-example lint

# Default target
help:
	@echo "OpenZL Go Bindings - Available targets:"
	@echo ""
	@echo "  deps          - Install dependencies"
	@echo "  build-openzl  - Build OpenZL shared library"
	@echo "  build         - Build Go packages"
	@echo "  test          - Run tests"
	@echo "  run-example   - Build and run hello example"
	@echo "  lint          - Run linter"
	@echo "  clean         - Clean build artifacts"
	@echo "  help          - Show this help message"
	@echo ""

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@chmod +x scripts/install-deps.sh
	@./scripts/install-deps.sh

# Build OpenZL shared library
build-openzl:
	@echo "Building OpenZL shared library..."
	@chmod +x scripts/build-openzl.sh
	@./scripts/build-openzl.sh

# Build Go packages
build: build-openzl
	@echo "Building Go packages..."
	@go build ./...

# Run tests
test: build-openzl
	@echo "Running tests..."
	@chmod +x scripts/test.sh
	@./scripts/test.sh

# Run hello example
run-example: build-openzl
	@echo "Building and running hello example..."
	@cd examples/hello && go build -o hello main.go
	@cd examples/hello && ./hello

# Run linter
lint:
	@echo "Running linter..."
	@go vet ./...
	@go fmt ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf third_party/openzl/build
	@rm -rf third_party/openzl/install
	@rm -f examples/hello/hello
	@rm -f coverage.out coverage.html
	@go clean ./...

# Development setup
dev-setup: deps
	@echo "Setting up development environment..."
	@git submodule update --init --recursive
	@make build-openzl
	@echo "Development setup complete!"

# CI target (used by GitHub Actions)
ci: deps build-openzl test lint
	@echo "CI pipeline completed successfully!"
