package openzl

import (
	"bytes"
	"testing"
)

func TestNewContext(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf("NewContext() failed: %v", err)
	}
	defer ctx.Close()

	if ctx == nil {
		t.Fatal("NewContext() returned nil context")
	}
}

func TestContextClose(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf("NewContext() failed: %v", err)
	}

	err = ctx.Close()
	if err != nil {
		t.Fatalf("Close() failed: %v", err)
	}

	// Test double close
	err = ctx.Close()
	if err != nil {
		t.Fatalf("Double Close() failed: %v", err)
	}
}

func TestCompressDecompress(t *testing.T) {
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
			ctx, err := NewContext()
			if err != nil {
				t.Fatalf("NewContext() failed: %v", err)
			}
			defer ctx.Close()

			// Compress
			compressed, err := ctx.Compress(tc.data)
			if err != nil {
				t.Fatalf("Compress() failed: %v", err)
			}

			// Verify compression
			if len(compressed) == 0 && len(tc.data) > 0 {
				t.Fatal("Compress() returned empty result for non-empty input")
			}

			// Decompress
			decompressed, err := ctx.Decompress(compressed)
			if err != nil {
				t.Fatalf("Decompress() failed: %v", err)
			}

			// Verify data integrity
			if !bytes.Equal(tc.data, decompressed) {
				t.Fatalf("Data integrity check failed: expected %v, got %v", tc.data, decompressed)
			}
		})
	}
}

func TestCompressWithClosedContext(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf("NewContext() failed: %v", err)
	}
	ctx.Close()

	data := []byte("test data")
	_, err = ctx.Compress(data)
	if err == nil {
		t.Fatal("Compress() with closed context should fail")
	}

	if err.Error() != "context is closed" {
		t.Fatalf("Expected 'context is closed' error, got: %v", err)
	}
}

func TestDecompressWithClosedContext(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf("NewContext() failed: %v", err)
	}
	ctx.Close()

	data := []byte("compressed data")
	_, err = ctx.Decompress(data)
	if err == nil {
		t.Fatal("Decompress() with closed context should fail")
	}

	if err.Error() != "context is closed" {
		t.Fatalf("Expected 'context is closed' error, got: %v", err)
	}
}

func TestCompressNilData(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf("NewContext() failed: %v", err)
	}
	defer ctx.Close()

	compressed, err := ctx.Compress(nil)
	if err != nil {
		t.Fatalf("Compress(nil) failed: %v", err)
	}

	if len(compressed) != 0 {
		t.Fatalf("Compress(nil) should return empty result, got: %v", compressed)
	}
}

func TestDecompressNilData(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf("NewContext() failed: %v", err)
	}
	defer ctx.Close()

	decompressed, err := ctx.Decompress(nil)
	if err != nil {
		t.Fatalf("Decompress(nil) failed: %v", err)
	}

	if len(decompressed) != 0 {
		t.Fatalf("Decompress(nil) should return empty result, got: %v", decompressed)
	}
}

func TestCompressionRatio(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf("NewContext() failed: %v", err)
	}
	defer ctx.Close()

	// Test with highly compressible data
	repeatedData := bytes.Repeat([]byte("This is highly repetitive data that should compress well. "), 1000)
	
	compressed, err := ctx.Compress(repeatedData)
	if err != nil {
		t.Fatalf("Compress() failed: %v", err)
	}

	compressionRatio := float64(len(compressed)) / float64(len(repeatedData))
	t.Logf("Compression ratio: %.2f%% (%d bytes -> %d bytes)", 
		compressionRatio*100, len(repeatedData), len(compressed))

	// For highly repetitive data, we should see significant compression
	if compressionRatio > 0.5 {
		t.Logf("Warning: Compression ratio is higher than expected: %.2f%%", compressionRatio*100)
	}
}

func BenchmarkCompress(b *testing.B) {
	data := bytes.Repeat([]byte("Benchmark test data for OpenZL compression performance testing. "), 1000)
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		ctx, err := NewContext()
		if err != nil {
			b.Fatalf("NewContext() failed: %v", err)
		}
		
		_, err = ctx.Compress(data)
		ctx.Close()
		
		if err != nil {
			b.Fatalf("Compress() failed: %v", err)
		}
	}
}

func BenchmarkDecompress(b *testing.B) {
	data := bytes.Repeat([]byte("Benchmark test data for OpenZL decompression performance testing. "), 1000)
	
	// Pre-compress the data once
	ctx, err := NewContext()
	if err != nil {
		b.Fatalf("NewContext() failed: %v", err)
	}
	compressed, err := ctx.Compress(data)
	ctx.Close()
	if err != nil {
		b.Fatalf("Compress() failed: %v", err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		ctx, err := NewContext()
		if err != nil {
			b.Fatalf("NewContext() failed: %v", err)
		}
		
		_, err = ctx.Decompress(compressed)
		ctx.Close()
		
		if err != nil {
			b.Fatalf("Decompress() failed: %v", err)
		}
	}
}

func BenchmarkRoundTrip(b *testing.B) {
	data := bytes.Repeat([]byte("Round trip benchmark test data for OpenZL compression and decompression. "), 1000)
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		ctx, err := NewContext()
		if err != nil {
			b.Fatalf("NewContext() failed: %v", err)
		}
		
		compressed, err := ctx.Compress(data)
		if err != nil {
			ctx.Close()
			b.Fatalf("Compress() failed: %v", err)
		}
		
		_, err = ctx.Decompress(compressed)
		ctx.Close()
		
		if err != nil {
			b.Fatalf("Decompress() failed: %v", err)
		}
	}
}
