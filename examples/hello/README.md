# Hello World Example

This is a basic example demonstrating how to use the OpenZL Go bindings.

## Description

The `hello` example shows:
- How to initialize an OpenZL context
- Basic compression and decompression operations
- Error handling patterns
- Resource cleanup

## Usage

```bash
# Build the example
go build -o hello main.go

# Run the example
./hello
```

## Expected Output

When the OpenZL library is properly linked and built, the example should output:

```
OpenZL Go Bindings - Hello World Example
✓ Successfully initialized OpenZL context
✓ Compression test passed: 13 bytes -> 8 bytes (38% compression)
✓ Decompression test passed: 8 bytes -> 13 bytes
✓ All tests completed successfully!
```

## Implementation Notes

This example demonstrates the basic API patterns that will be implemented:

1. **Context Management**: Creating and properly closing OpenZL contexts
2. **Compression**: Using the compression API with sample data
3. **Decompression**: Verifying that compressed data can be decompressed correctly
4. **Error Handling**: Proper error checking and reporting
5. **Resource Cleanup**: Using defer statements for proper cleanup

The actual implementation will be added once the cgo bindings are complete.
