.PHONY: help deps build test clean install-deps build-openzl run-example lint

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

deps:
	@echo "Installing dependencies..."
	@chmod +x scripts/install-deps.sh
	@./scripts/install-deps.sh

build-openzl:
	@echo "Building OpenZL shared library..."
	@chmod +x scripts/build-openzl.sh
	@./scripts/build-openzl.sh

build: build-openzl
	@echo "Building Go packages..."
	@go build ./...

test: build-openzl
	@echo "Running tests..."
	@chmod +x scripts/test.sh
	@./scripts/test.sh

run-example: build-openzl
	@echo "Building and running hello example..."
	@cd examples/hello && go build -o hello main.go
	@cd examples/hello && ./hello

lint:
	@echo "Running linter..."
	@go vet ./...
	@go fmt ./...

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf third_party/openzl/build
	@rm -rf third_party/openzl/install
	@rm -f examples/hello/hello
	@rm -f coverage.out coverage.html
	@go clean ./...

dev-setup: deps
	@echo "Setting up development environment..."
	@git submodule update --init --recursive
	@make build-openzl
	@echo "Development setup complete!"

ci: deps build-openzl test lint
	@echo "CI pipeline completed successfully!"
