//go:build cgo
// +build cgo

// C wrapper implementations for OpenZL Go bindings
#include "openzl.h"

// Initialize a new OpenZL context
openzl_context_t* openzl_context_create() {
    openzl_context_t* ctx = (openzl_context_t*)malloc(sizeof(openzl_context_t));
    if (ctx == NULL) {
        return NULL;
    }
    
    ctx->cctx = ZL_CCtx_create();
    ctx->dctx = ZL_DCtx_create();
    
    if (ctx->cctx == NULL || ctx->dctx == NULL) {
        openzl_context_free(ctx);
        return NULL;
    }
    
    // Set default compression parameters
    unsigned defaultVersion = ZL_getDefaultEncodingVersion();
    ZL_Report result = ZL_CCtx_setParameter(ctx->cctx, ZL_CParam_formatVersion, (int)defaultVersion);
    if (ZL_isError(result)) {
        openzl_context_free(ctx);
        return NULL;
    }
    
    // Set default compression level
    result = ZL_CCtx_setParameter(ctx->cctx, ZL_CParam_compressionLevel, ZL_COMPRESSIONLEVEL_DEFAULT);
    if (ZL_isError(result)) {
        openzl_context_free(ctx);
        return NULL;
    }
    
    return ctx;
}

// Free an OpenZL context
void openzl_context_free(openzl_context_t* ctx) {
    if (ctx == NULL) {
        return;
    }
    
    if (ctx->cctx != NULL) {
        ZL_CCtx_free(ctx->cctx);
    }
    
    if (ctx->dctx != NULL) {
        ZL_DCtx_free(ctx->dctx);
    }
    
    free(ctx);
}

// Compress data
// Returns the compressed size on success, or a negative error code on failure
long long openzl_compress(openzl_context_t* ctx, 
                         void* dst, size_t dst_capacity,
                         const void* src, size_t src_size) {
    if (ctx == NULL || ctx->cctx == NULL) {
        return -1; // Invalid context
    }
    
    ZL_Report result = ZL_CCtx_compress(ctx->cctx, dst, dst_capacity, src, src_size);
    
    if (ZL_isError(result)) {
        return -(long long)ZL_errorCode(result); // Return negative error code
    }
    
    return (long long)ZL_validResult(result); // Return compressed size
}

// Decompress data
// Returns the decompressed size on success, or a negative error code on failure
long long openzl_decompress(openzl_context_t* ctx,
                           void* dst, size_t dst_capacity, 
                           const void* src, size_t src_size) {
    if (ctx == NULL || ctx->dctx == NULL) {
        return -1; // Invalid context
    }
    
    ZL_Report result = ZL_DCtx_decompress(ctx->dctx, dst, dst_capacity, src, src_size);
    
    if (ZL_isError(result)) {
        return -(long long)ZL_errorCode(result); // Return negative error code
    }
    
    return (long long)ZL_validResult(result); // Return decompressed size
}

// Get the maximum size needed for compression output
size_t openzl_compress_bound(size_t src_size) {
    return ZL_compressBound(src_size);
}
