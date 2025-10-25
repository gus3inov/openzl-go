// C header includes for OpenZL
#ifndef OPENZL_H
#define OPENZL_H

#include "openzl/openzl.h"
#include "openzl/zl_compressor.h"
#include <stdlib.h>

// C wrapper function declarations

// Context management
typedef struct {
    ZL_CCtx* cctx;
    ZL_DCtx* dctx;
} openzl_context_t;

// Initialize a new OpenZL context
openzl_context_t* openzl_context_create();

// Free an OpenZL context
void openzl_context_free(openzl_context_t* ctx);

// Compress data
// Returns the compressed size on success, or a negative error code on failure
long long openzl_compress(openzl_context_t* ctx, 
                         void* dst, size_t dst_capacity,
                         const void* src, size_t src_size);

// Decompress data  
// Returns the decompressed size on success, or a negative error code on failure
long long openzl_decompress(openzl_context_t* ctx,
                           void* dst, size_t dst_capacity, 
                           const void* src, size_t src_size);

// Get the maximum size needed for compression output
size_t openzl_compress_bound(size_t src_size);

#endif // OPENZL_H
