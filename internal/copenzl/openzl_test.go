package copenzl

import (
	"bytes"
	"testing"
)

func TestNewOpenZLContext(t *testing.T) {
	ctx, err := NewOpenZLContext()
	if err != nil {
		t.Fatalf("NewOpenZLContext() failed: %v", err)
	}
	defer ctx.Close()

	if ctx == nil {
		t.Fatal("NewOpenZLContext() returned nil context")
	}
	if ctx.ctx == nil {
		t.Fatal("NewOpenZLContext() returned context with nil ctx")
	}
}

func TestOpenZLContextClose(t *testing.T) {
	ctx, err := NewOpenZLContext()
	if err != nil {
		t.Fatalf("NewOpenZLContext() failed: %v", err)
	}

	ctx.Close()
	// Test double close - should not panic
	ctx.Close()
}

func TestOpenZLCompressDecompress(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{
			name: "empty data",
			data: []byte{},
		},
		{
			name: "small data",
			data: []byte("Hello, World!"),
		},
		{
			name: "medium data",
			data: bytes.Repeat([]byte("OpenZL compression test data "), 100),
		},
		{
			name: "large data",
			data: bytes.Repeat([]byte("This is a test of OpenZL compression with repeated data to test compression ratios. "), 1000),
		},
		{
			name: "binary data",
			data: []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new context for each test case to avoid state pollution
			ctx, err := NewOpenZLContext()
			if err != nil {
				t.Fatalf("NewOpenZLContext() failed: %v", err)
			}
			defer ctx.Close()

			// Compress
			compressed, err := OpenZLCompress(ctx, tc.data)
			if err != nil {
				t.Fatalf("OpenZLCompress() failed: %v", err)
			}

			// Verify compression
			if len(compressed) == 0 && len(tc.data) > 0 {
				t.Fatal("OpenZLCompress() returned empty result for non-empty input")
			}

			// Decompress
			decompressed, err := OpenZLDecompress(ctx, compressed)
			if err != nil {
				t.Fatalf("OpenZLDecompress() failed: %v", err)
			}

			// Verify data integrity
			if !bytes.Equal(tc.data, decompressed) {
				t.Fatalf("Data integrity check failed: expected %v, got %v", tc.data, decompressed)
			}
		})
	}
}

func TestOpenZLCompressWithNilContext(t *testing.T) {
	data := []byte("test data")
	_, err := OpenZLCompress(nil, data)
	if err == nil {
		t.Fatal("OpenZLCompress() with nil context should fail")
	}

	if err.Error() != "invalid context" {
		t.Fatalf("Expected 'invalid context' error, got: %v", err)
	}
}

func TestOpenZLDecompressWithNilContext(t *testing.T) {
	data := []byte("compressed data")
	_, err := OpenZLDecompress(nil, data)
	if err == nil {
		t.Fatal("OpenZLDecompress() with nil context should fail")
	}

	if err.Error() != "invalid context" {
		t.Fatalf("Expected 'invalid context' error, got: %v", err)
	}
}

func TestOpenZLCompressNilData(t *testing.T) {
	ctx, err := NewOpenZLContext()
	if err != nil {
		t.Fatalf("NewOpenZLContext() failed: %v", err)
	}
	defer ctx.Close()

	compressed, err := OpenZLCompress(ctx, nil)
	if err != nil {
		t.Fatalf("OpenZLCompress(nil) failed: %v", err)
	}

	if len(compressed) != 0 {
		t.Fatalf("OpenZLCompress(nil) should return empty result, got: %v", compressed)
	}
}

func TestOpenZLDecompressNilData(t *testing.T) {
	ctx, err := NewOpenZLContext()
	if err != nil {
		t.Fatalf("NewOpenZLContext() failed: %v", err)
	}
	defer ctx.Close()

	decompressed, err := OpenZLDecompress(ctx, nil)
	if err != nil {
		t.Fatalf("OpenZLDecompress(nil) failed: %v", err)
	}

	if len(decompressed) != 0 {
		t.Fatalf("OpenZLDecompress(nil) should return empty result, got: %v", decompressed)
	}
}

func BenchmarkOpenZLCompress(b *testing.B) {
	ctx, err := NewOpenZLContext()
	if err != nil {
		b.Fatalf("NewOpenZLContext() failed: %v", err)
	}
	defer ctx.Close()

	data := bytes.Repeat([]byte("Benchmark test data for OpenZL compression performance testing. "), 1000)
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_, err := OpenZLCompress(ctx, data)
		if err != nil {
			b.Fatalf("OpenZLCompress() failed: %v", err)
		}
	}
}

func BenchmarkOpenZLDecompress(b *testing.B) {
	ctx, err := NewOpenZLContext()
	if err != nil {
		b.Fatalf("NewOpenZLContext() failed: %v", err)
	}
	defer ctx.Close()

	data := bytes.Repeat([]byte("Benchmark test data for OpenZL decompression performance testing. "), 1000)
	compressed, err := OpenZLCompress(ctx, data)
	if err != nil {
		b.Fatalf("OpenZLCompress() failed: %v", err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_, err := OpenZLDecompress(ctx, compressed)
		if err != nil {
			b.Fatalf("OpenZLDecompress() failed: %v", err)
		}
	}
}
