# OpenZL Go Bindings

[![CI](https://github.com/gus3inov/openzl-go/workflows/CI/badge.svg)](https://github.com/gus3inov/openzl-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gus3inov/openzl-go)](https://goreportcard.com/report/github.com/gus3inov/openzl-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Go bindings for Facebook's [OpenZL](https://github.com/facebook/openzl) library, providing idiomatic Go APIs for compression and machine learning workloads.

## Overview

This project provides cgo-based bindings that wrap the native OpenZL C API, enabling Go applications to leverage OpenZL's compression algorithms and ML capabilities. The binding architecture follows this flow:

```
OpenZL C API → cgo layer → Go wrapper → Public Go API
```

## Project Structure

```
openzl-go/
├── README.md                 # This file
├── .gitignore              # Git ignore patterns
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── LICENSE                  # MIT license
├── third_party/openzl/      # Git submodule for OpenZL
├── cgo/                     # cgo layer and build glue
│   ├── openzl.h            # C header includes
│   ├── openzl.c            # C wrapper implementations
│   └── build.go            # Build constraints and flags
├── internal/copenzl/        # Low-level C bindings (private)
│   ├── openzl.go           # cgo declarations
│   └── types.go            # C type mappings
├── openzl/                  # Public Go API package
│   ├── compression.go      # Compression APIs
│   ├── ml.go              # Machine learning APIs
│   └── types.go           # Go type definitions
├── examples/               # Sample applications
│   └── hello/             # Basic usage example
│       └── main.go        # Hello world app
├── scripts/               # Build and automation scripts
│   ├── build-openzl.sh    # Build OpenZL shared library
│   ├── install-deps.sh    # Install prerequisites
│   └── test.sh            # Run tests
└── .github/workflows/     # GitHub Actions CI
    └── ci.yml             # Continuous integration
```

## Prerequisites

### macOS

```bash
# Install Xcode Command Line Tools
xcode-select --install

# Install dependencies via Homebrew
brew install go cmake ninja llvm

# Verify installations
go version
cmake --version
ninja --version
clang --version
```

### Linux (Ubuntu/Debian)

```bash
# Install dependencies
sudo apt update
sudo apt install -y build-essential cmake ninja-build golang-go llvm-dev libclang-dev

# Verify installations
go version
cmake --version
ninja --version
clang --version
```

## Quick Start

1. **Clone the repository and initialize submodule:**

```bash
git clone https://github.com/gus3inov/openzl-go.git
cd openzl-go
git submodule update --init --recursive
```

2. **Install prerequisites:**

```bash
# macOS
./scripts/install-deps.sh

# Linux
./scripts/install-deps.sh
```

3. **Build OpenZL shared library:**

```bash
./scripts/build-openzl.sh
```

4. **Build and run the example:**

```bash
cd examples/hello
go build -o hello
./hello
```

5. **Run tests:**

```bash
go test ./...
```

## Building OpenZL

The OpenZL  library must be built as a shared library before using the Go bindings:

```bash
# Build OpenZL (creates libopenzl.so/.dylib)
./scripts/build-openzl.sh

# Verify the shared library was created
ls -la third_party/openzl/build/libopenzl.*
```

## Usage

### Basic Compression

```go
package main

import (
    "fmt"
    "github.com/gus3inov/openzl-go/openzl"
)

func main() {
    // Initialize OpenZL context
    ctx, err := openzl.NewContext()
    if err != nil {
        panic(err)
    }
    defer ctx.Close()

    // Use compression APIs
    data := []byte("Hello, OpenZL!")
    compressed, err := ctx.Compress(data)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Compressed %d bytes to %d bytes\n",
        len(data), len(compressed))
}
```

### Context Reuse (Recommended)

**Important**: Contexts can and should be reused for better performance. Reusing a context
improves performance by approximately **27%** compared to creating a new context for each operation.

```go
ctx, err := openzl.NewContext()
if err != nil {
    panic(err)
}
defer ctx.Close()

// Reuse the same context for multiple operations
for _, data := range datasets {
    compressed, err := ctx.Compress(data)
    if err != nil {
        panic(err)
    }
    // Process compressed data...
}
```

**Note**: Contexts are not thread-safe. Each goroutine should use its own context instance.

## Project Status

### ✅ Completed (v0.1.0)
- [x] Basic project structure
- [x] Core compression API bindings
- [x] Basic example application
- [x] CI/CD pipeline
- [x] Comprehensive test suite
- [x] Performance benchmarks
- [x] Go documentation
- [x] Error handling
- [x] Memory management

### 🚧 Future Roadmap

#### Phase 2: Enhanced Features
- [ ] Streaming compression/decompression
- [ ] Memory-efficient APIs
- [ ] Progress callbacks
- [ ] Compression level configuration
- [ ] Custom compression strategies

#### Phase 3: ML Integration
- [ ] Training API bindings
- [ ] Model inference support
- [ ] Custom compression graphs

#### Phase 4: Production Ready
- [ ] Prebuilt binaries for releases
- [ ] Windows support
- [ ] Cross-platform CI/CD
- [ ] Performance optimizations
- [ ] Advanced error handling

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with race detection
go test -race ./...
```

### Building from Source

```bash
# Build all packages
go build ./...

# Build specific example
go build -o examples/hello/hello examples/hello/main.go
```

### Code Generation

If you modify C headers or need to regenerate bindings:

```bash
# Regenerate cgo bindings (if needed)
go generate ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes
4. Add tests for new functionality
5. Run the test suite: `go test ./...`
6. Commit your changes: `git commit -am 'Add feature'`
7. Push to the branch: `git push origin feature-name`
8. Submit a pull request

## Disclaimers

- **API Stability**: This project is in early development. APIs may change between versions.
- **Platform Support**: Currently supports macOS and Linux. Windows support is planned.
- **Performance**: cgo overhead may impact performance for high-frequency operations.
- **Dependencies**: Requires OpenZL to be built from source (no prebuilt binaries yet).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Facebook OpenZL](https://github.com/facebook/openzl) - The underlying compression library
- [Go cgo documentation](https://pkg.go.dev/cmd/cgo) - For cgo implementation guidance

## Links

- [OpenZL Repository](https://github.com/facebook/openzl)
- [Go Documentation](https://golang.org/doc/)
- [cgo Documentation](https://pkg.go.dev/cmd/cgo)
