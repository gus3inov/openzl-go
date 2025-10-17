#!/bin/bash

# test.sh - Run tests for OpenZL Go bindings
# This script runs the test suite with various configurations

set -e

echo "Running OpenZL Go bindings tests..."

# Check if OpenZL is built
if [ ! -d "third_party/openzl/build" ]; then
    echo "OpenZL not built. Building first..."
    ./scripts/build-openzl.sh
fi

# Run basic tests
echo "Running basic tests..."
go test -v ./...

# Run tests with race detection
echo "Running tests with race detection..."
go test -race ./...

# Run tests with coverage
echo "Running tests with coverage..."
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run benchmarks (if any)
echo "Running benchmarks..."
go test -bench=. ./...

# Build examples to verify they compile
echo "Building examples..."
cd examples/hello
go build -o hello main.go
echo "✓ Hello example built successfully"

# Run the hello example (if it doesn't crash, it's a good sign)
echo "Running hello example..."
if ./hello; then
    echo "✓ Hello example ran successfully"
else
    echo "✗ Hello example failed"
    exit 1
fi

cd ../..

echo "All tests completed successfully!"
echo ""
echo "Coverage report generated: coverage.html"
echo "To view coverage: open coverage.html"
