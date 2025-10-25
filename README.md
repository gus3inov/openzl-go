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
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Links

- [OpenZL Repository](https://github.com/facebook/openzl)
- [Go Documentation](https://golang.org/doc/)
- [cgo Documentation](https://pkg.go.dev/cmd/cgo)
