# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of OpenZL Go bindings
- Basic compression and decompression functionality
- Context-based API for managing compression state
- Comprehensive test suite with unit tests and benchmarks
- Example application demonstrating basic usage
- CI/CD pipeline with GitHub Actions
- Support for macOS and Linux platforms
- Go documentation with examples and usage patterns

### Features
- **Context Management**: Create and manage OpenZL contexts for compression operations
- **Compression**: Compress data using OpenZL's default compression settings
- **Decompression**: Decompress data compressed with compatible OpenZL compressors
- **Error Handling**: Detailed error information with error codes and messages
- **Memory Management**: Automatic resource cleanup with context closing
- **Thread Safety**: Clear documentation about concurrent usage patterns

### Technical Details
- Uses cgo to interface with the native OpenZL C library
- Supports empty data compression/decompression
- Handles various data sizes from small to large datasets
- Provides compression ratio information
- Includes performance benchmarks for compression and decompression

### Dependencies
- Requires OpenZL library to be built from source
- Compatible with Go 1.21+
- Uses cgo for C library integration

### Limitations
- Currently requires manual building of OpenZL shared library
- No prebuilt binaries available yet
- Windows support planned for future releases
- API may change in future versions as OpenZL evolves

## [0.1.0] - 2024-10-17

### Added
- Initial implementation of OpenZL Go bindings
- Basic compression and decompression API
- Test suite and example application
- Documentation and build scripts
