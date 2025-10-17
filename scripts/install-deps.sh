#!/bin/bash

# install-deps.sh - Install prerequisites for building OpenZL Go bindings
# Supports macOS and Linux (Ubuntu/Debian)

set -e

echo "Installing prerequisites for OpenZL Go bindings..."

# Detect operating system
if [[ "$OSTYPE" == "darwin"* ]]; then
    OS="macos"
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS="linux"
else
    echo "Unsupported operating system: $OSTYPE"
    exit 1
fi

echo "Detected OS: $OS"

# Check if running as root (not recommended)
if [[ $EUID -eq 0 ]]; then
    echo "Warning: Running as root. Consider using a non-root user."
fi

# Install dependencies based on OS
if [[ "$OS" == "macos" ]]; then
    echo "Installing dependencies for macOS..."
    
    # Check if Homebrew is installed
    if ! command -v brew &> /dev/null; then
        echo "Homebrew not found. Installing Homebrew..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    fi
    
    # Install required packages
    echo "Installing packages via Homebrew..."
    brew install go cmake ninja llvm
    
    # Set up LLVM paths for macOS
    echo "Setting up LLVM environment..."
    export PATH="/opt/homebrew/bin:$PATH"
    export LLVM_CONFIG="/opt/homebrew/bin/llvm-config"
    
elif [[ "$OS" == "linux" ]]; then
    echo "Installing dependencies for Linux..."
    
    # Update package list
    sudo apt update
    
    # Install required packages
    sudo apt install -y \
        build-essential \
        cmake \
        ninja-build \
        golang-go \
        llvm-dev \
        libclang-dev \
        pkg-config
    
    # Set up LLVM paths for Linux
    export LLVM_CONFIG="/usr/bin/llvm-config"
fi

# Verify installations
echo "Verifying installations..."

echo "Go version:"
go version

echo "CMake version:"
cmake --version

echo "Ninja version:"
ninja --version

echo "Clang version:"
clang --version

echo "LLVM config:"
if command -v llvm-config &> /dev/null; then
    llvm-config --version
else
    echo "llvm-config not found in PATH"
fi

echo "Prerequisites installation completed successfully!"
echo ""
echo "Next steps:"
echo "1. Initialize the OpenZL submodule: git submodule update --init --recursive"
echo "2. Build OpenZL: ./scripts/build-openzl.sh"
echo "3. Run tests: go test ./..."
