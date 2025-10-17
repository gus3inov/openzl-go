#!/bin/bash

# build-openzl.sh - Build OpenZL as a shared library
# This script builds the OpenZL library from source and creates shared libraries
# that can be used by the Go bindings via cgo

set -e

echo "Building OpenZL shared library..."

# Check if submodule is initialized
if [ ! -d "third_party/openzl/.git" ]; then
    echo "OpenZL submodule not initialized. Initializing..."
    git submodule update --init --recursive
fi

# Change to OpenZL directory
cd third_party/openzl

# Check if we're in the right directory
if [ ! -f "CMakeLists.txt" ]; then
    echo "Error: CMakeLists.txt not found. Are we in the right directory?"
    exit 1
fi

echo "Building OpenZL in $(pwd)"

# Create build directory
mkdir -p build
cd build

# Configure CMake
echo "Configuring CMake..."
cmake .. \
    -DCMAKE_BUILD_TYPE=Release \
    -DBUILD_SHARED_LIBS=ON \
    -DCMAKE_INSTALL_PREFIX=../install \
    -G Ninja

# Build only the core OpenZL shared library target to avoid optional components
echo "Building OpenZL (core library only)..."

# Try to build a specific target if it exists, falling back to full build
if ninja -t targets all | grep -q '^openzl:'; then
    ninja openzl
elif ninja -t targets all | grep -q '^libopenzl:'; then
    ninja libopenzl
elif ninja -t targets all | grep -q '^openzl_shared:'; then
    ninja openzl_shared
else
    echo "Specific openzl target not found, performing minimal build (this may build extras)"
    ninja
fi

# Verify the shared library was created
echo "Verifying build..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    LIB_EXT="dylib"
else
    LIB_EXT="so"
fi

LIB_PATH=""
if [ -f "libopenzl.$LIB_EXT" ]; then
    LIB_PATH="$(pwd)/libopenzl.$LIB_EXT"
elif [ -f "libopenzl.0.0.0-dev.$LIB_EXT" ]; then
    LIB_PATH="$(pwd)/libopenzl.0.0.0-dev.$LIB_EXT"
else
    # Try to find the library in common locations inside build dir
    LIB_CANDIDATE=$(find . -maxdepth 2 -name "libopenzl*.${LIB_EXT}" | head -n 1)
    if [ -n "$LIB_CANDIDATE" ]; then
        LIB_PATH="$(pwd)/${LIB_CANDIDATE#./}"
    fi
fi

if [ -n "$LIB_PATH" ] && [ -f "$LIB_PATH" ]; then
    echo "✓ Shared library created: $LIB_PATH"
    ls -la "$LIB_PATH"
else
    echo "✗ Shared library not found!"
    exit 1
fi

# Skip installation to avoid building optional components; consumers can link from build output

# Return to project root
cd ../../..

echo "OpenZL build completed successfully!"
echo ""
echo "Build artifacts:"
echo "  - Shared library: third_party/openzl/build/libopenzl.$LIB_EXT"
echo "  - Installed library: third_party/openzl/install/lib/libopenzl.$LIB_EXT"
echo ""
echo "Next steps:"
echo "1. Run tests: go test ./..."
echo "2. Build examples: cd examples/hello && go build -o hello main.go"
