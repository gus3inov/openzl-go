// Package openzl provides Go bindings for Facebook's OpenZL library.
//
// OpenZL is a novel data compression framework that delivers high compression
// ratios while preserving high speed. This package offers idiomatic Go APIs
// for compression and machine learning workloads using the OpenZL library.
//
// Basic Usage:
//
//	ctx, err := openzl.NewContext()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer ctx.Close()
//
//	data := []byte("Hello, World!")
//	compressed, err := ctx.Compress(data)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	decompressed, err := ctx.Decompress(compressed)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Compressed %d bytes to %d bytes\n", len(data), len(compressed))
//
// Context Reuse:
//
// Contexts can and should be reused for multiple operations. This improves
// performance by avoiding context creation overhead (~27% faster):
//
//	ctx, _ := openzl.NewContext()
//	defer ctx.Close()
//
//	for _, data := range datasets {
//		compressed, _ := ctx.Compress(data) // Reuse same context
//		// Process compressed data...
//	}
//
// Error Handling:
//
// The package provides detailed error information through the Error type,
// which includes both an error code and a descriptive message.
//
// Performance:
//
// OpenZL is designed for high-performance compression workloads. The Go bindings
// use cgo to interface with the native C library, which may introduce some
// overhead for very high-frequency operations.
//
// Thread Safety:
//
// Context objects are not safe for concurrent use. Each goroutine should use
// its own Context instance.
package openzl

import (
	"github.com/gus3inov/openzl-go/internal/copenzl"
)

// Error represents an OpenZL error with detailed information.
type Error struct {
	Code    int    // Error code from the OpenZL library
	Message string // Human-readable error description
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.Message
}

// Context represents an OpenZL compression/decompression context.
//
// A Context manages the state needed for compression and decompression operations.
// Contexts can and should be reused for multiple operations to improve performance.
// Context creation incurs overhead, so reusing a context across multiple compress/
// decompress calls is recommended for better performance.
//
// Thread Safety: Contexts are not safe for concurrent use by multiple goroutines.
// Each goroutine should use its own Context instance.
//
// Performance: Reusing a context can improve performance by ~27% compared to creating
// a new context for each operation.
type Context struct {
	ctx *copenzl.OpenZLContext
}

// NewContext creates a new OpenZL context.
//
// The context must be closed when no longer needed to free associated resources.
// Returns an error if the context could not be created.
func NewContext() (*Context, error) {
	ctx, err := copenzl.NewOpenZLContext()
	if err != nil {
		return nil, err
	}
	return &Context{ctx: ctx}, nil
}

// Close closes the OpenZL context and frees associated resources.
//
// It is safe to call Close multiple times. After calling Close, the context
// should not be used for further operations.
func (c *Context) Close() error {
	if c.ctx != nil {
		c.ctx.Close()
		c.ctx = nil
	}
	return nil
}

// Compress compresses the given data using the context.
//
// The compression uses OpenZL's default compression settings. For empty input,
// returns empty output. Returns an error if compression fails.
func (c *Context) Compress(data []byte) ([]byte, error) {
	if c.ctx == nil {
		return nil, &Error{Code: -1, Message: "context is closed"}
	}
	return copenzl.OpenZLCompress(c.ctx, data)
}

// Decompress decompresses the given data using the context.
//
// The data must have been compressed with a compatible OpenZL compressor.
// Returns an error if decompression fails or if the compressed data is invalid.
func (c *Context) Decompress(data []byte) ([]byte, error) {
	if c.ctx == nil {
		return nil, &Error{Code: -1, Message: "context is closed"}
	}
	return copenzl.OpenZLDecompress(c.ctx, data)
}
