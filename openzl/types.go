// Package types contains type definitions for OpenZL Go bindings.
package openzl

import (
	"github.com/gus3inov/openzl-go/internal/copenzl"
)

// Error represents an OpenZL error.
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// Context represents an OpenZL context.
type Context struct {
	ctx *copenzl.OpenZLContext
}

// NewContext creates a new OpenZL context.
func NewContext() (*Context, error) {
	ctx, err := copenzl.NewOpenZLContext()
	if err != nil {
		return nil, err
	}
	return &Context{ctx: ctx}, nil
}

// Close closes the OpenZL context.
func (c *Context) Close() error {
	if c.ctx != nil {
		c.ctx.Close()
		c.ctx = nil
	}
	return nil
}

// Compress compresses the given data using the context.
func (c *Context) Compress(data []byte) ([]byte, error) {
	if c.ctx == nil {
		return nil, &Error{Code: -1, Message: "context is closed"}
	}
	return copenzl.OpenZLCompress(c.ctx, data)
}

// Decompress decompresses the given data using the context.
func (c *Context) Decompress(data []byte) ([]byte, error) {
	if c.ctx == nil {
		return nil, &Error{Code: -1, Message: "context is closed"}
	}
	return copenzl.OpenZLDecompress(c.ctx, data)
}
