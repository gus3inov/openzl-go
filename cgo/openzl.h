// C header includes for OpenZL
#ifndef OPENZL_H
#define OPENZL_H

#include "openzl/openzl.h"
#include "openzl/zl_compress.h"
#include "openzl/zl_decompress.h"
#include "openzl/zl_compressor.h"
#include <stdlib.h>

typedef struct {
    ZL_CCtx* cctx;
    ZL_DCtx* dctx;
} openzl_context_t;

openzl_context_t* openzl_context_create();

void openzl_context_free(openzl_context_t* ctx);

long long openzl_compress(openzl_context_t* ctx, 
                         void* dst, size_t dst_capacity,
                         const void* src, size_t src_size);

long long openzl_decompress(openzl_context_t* ctx,
                           void* dst, size_t dst_capacity, 
                           const void* src, size_t src_size);

size_t openzl_compress_bound(size_t src_size);

#endif // OPENZL_H
