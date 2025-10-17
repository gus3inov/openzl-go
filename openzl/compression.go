// Package compression provides compression and decompression APIs.
//
// This package contains the core compression functionality of OpenZL,
// including various compression algorithms and streaming support.
package openzl

// CompressionLevel represents the compression level to use.
type CompressionLevel int

const (
	// Fastest compression with lowest compression ratio
	Fastest CompressionLevel = iota
	// Default compression level
	Default
	// Best compression with highest compression ratio
	Best
)

// CompressionOptions contains options for compression operations.
type CompressionOptions struct {
	Level CompressionLevel
	// Additional options will be added as needed
}

// Compressor handles compression operations.
type Compressor struct {
	// Implementation details will be added
}

// NewCompressor creates a new compressor with the given options.
func NewCompressor(opts CompressionOptions) *Compressor {
	// Implementation will be added
	return &Compressor{}
}

// Compress compresses the given data.
func (c *Compressor) Compress(data []byte) ([]byte, error) {
	// Implementation will be added
	return nil, nil
}

// Decompress decompresses the given data.
func (c *Compressor) Decompress(data []byte) ([]byte, error) {
	// Implementation will be added
	return nil, nil
}
