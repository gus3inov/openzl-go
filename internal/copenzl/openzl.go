// Package copenzl provides low-level C bindings for OpenZL.
//
// This package contains the cgo declarations and C type mappings
// for the OpenZL library. It should not be used directly by users
// of the openzl package.
package copenzl

/*
#cgo CFLAGS: -I${SRCDIR}/../../cgo -I${SRCDIR}/../../third_party/openzl/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/openzl/build -lopenzl -Wl,-rpath,${SRCDIR}/../../third_party/openzl/build

#include "../../cgo/openzl.h"
#include "../../cgo/openzl.c"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// OpenZLContext represents a C OpenZL context.
type OpenZLContext struct {
	ctx *C.openzl_context_t
}

// NewOpenZLContext creates a new OpenZL context.
func NewOpenZLContext() (*OpenZLContext, error) {
	ctx := C.openzl_context_create()
	if ctx == nil {
		return nil, errors.New("failed to create OpenZL context")
	}
	return &OpenZLContext{ctx: ctx}, nil
}

// Close frees the OpenZL context.
func (c *OpenZLContext) Close() {
	if c.ctx != nil {
		C.openzl_context_free(c.ctx)
		c.ctx = nil
	}
}

// OpenZLCompress compresses data using the C API.
func OpenZLCompress(ctx *OpenZLContext, data []byte) ([]byte, error) {
	if ctx == nil || ctx.ctx == nil {
		return nil, errors.New("invalid context")
	}

	if len(data) == 0 {
		return []byte{}, nil
	}

	// Calculate maximum compressed size
	maxCompressedSize := C.openzl_compress_bound(C.size_t(len(data)))
	compressed := make([]byte, maxCompressedSize)

	// Call the C function
	var srcPtr unsafe.Pointer
	var dstPtr unsafe.Pointer
	
	if len(data) > 0 {
		srcPtr = unsafe.Pointer(&data[0])
	}
	if len(compressed) > 0 {
		dstPtr = unsafe.Pointer(&compressed[0])
	}
	
	result := C.openzl_compress(
		ctx.ctx,
		dstPtr,
		C.size_t(len(compressed)),
		srcPtr,
		C.size_t(len(data)),
	)

	if result < 0 {
		return nil, fmt.Errorf("compression failed with error code %d (data size: %d, buffer size: %d)", -result, len(data), len(compressed))
	}

	// Return the actual compressed data (truncated to actual size)
	actualSize := int(result)
	return compressed[:actualSize], nil
}

// OpenZLDecompress decompresses data using the C API.
func OpenZLDecompress(ctx *OpenZLContext, data []byte) ([]byte, error) {
	if ctx == nil || ctx.ctx == nil {
		return nil, errors.New("invalid context")
	}

	if len(data) == 0 {
		return []byte{}, nil
	}

	// First, get the decompressed size
	sizeResult := C.ZL_getDecompressedSize(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if C.ZL_isError(sizeResult) != 0 {
		return nil, fmt.Errorf("failed to get decompressed size: error code %d", C.ZL_errorCode(sizeResult))
	}

	decompressedSize := int(C.ZL_validResult(sizeResult))
	decompressed := make([]byte, decompressedSize)

	// Call the simple decompression function
	result := C.ZL_decompress(
		unsafe.Pointer(&decompressed[0]),
		C.size_t(len(decompressed)),
		unsafe.Pointer(&data[0]),
		C.size_t(len(data)),
	)

	if C.ZL_isError(result) != 0 {
		return nil, fmt.Errorf("decompression failed with error code %d", C.ZL_errorCode(result))
	}

	// Return the actual decompressed data (truncated to actual size)
	actualSize := int(C.ZL_validResult(result))
	return decompressed[:actualSize], nil
}
