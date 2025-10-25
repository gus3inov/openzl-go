//go:build cgo
// +build cgo

// C wrapper implementations for OpenZL Go bindings
#include "openzl.h"

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

    // Enable sticky parameters to allow context reuse across multiple operations
    ZL_Report result = ZL_CCtx_setParameter(ctx->cctx, ZL_CParam_stickyParameters, 1);
    if (ZL_isError(result)) {
        openzl_context_free(ctx);
        return NULL;
    }

    result = ZL_DCtx_setParameter(ctx->dctx, ZL_DParam_stickyParameters, 1);
    if (ZL_isError(result)) {
        openzl_context_free(ctx);
        return NULL;
    }

    // Set default compression parameters
    unsigned defaultVersion = ZL_getDefaultEncodingVersion();
    result = ZL_CCtx_setParameter(ctx->cctx, ZL_CParam_formatVersion, (int)defaultVersion);
    if (ZL_isError(result)) {
        openzl_context_free(ctx);
        return NULL;
    }

    result = ZL_CCtx_setParameter(ctx->cctx, ZL_CParam_compressionLevel, ZL_COMPRESSIONLEVEL_DEFAULT);
    if (ZL_isError(result)) {
        openzl_context_free(ctx);
        return NULL;
    }

    return ctx;
}

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

long long openzl_compress(openzl_context_t* ctx, 
                         void* dst, size_t dst_capacity,
                         const void* src, size_t src_size) {
    if (ctx == NULL || ctx->cctx == NULL) {
        return -1;
    }
    
    ZL_Report result = ZL_CCtx_compress(ctx->cctx, dst, dst_capacity, src, src_size);
    
    if (ZL_isError(result)) {
        return -(long long)ZL_errorCode(result);
    }
    
    return (long long)ZL_validResult(result);
}

long long openzl_decompress(openzl_context_t* ctx,
                           void* dst, size_t dst_capacity, 
                           const void* src, size_t src_size) {
    if (ctx == NULL || ctx->dctx == NULL) {
        return -1;
    }
    
    ZL_Report result = ZL_decompress(dst, dst_capacity, src, src_size);
    
    if (ZL_isError(result)) {
        return -(long long)ZL_errorCode(result);
    }
    
    return (long long)ZL_validResult(result);
}

size_t openzl_compress_bound(size_t src_size) {
    return ZL_compressBound(src_size);
}
