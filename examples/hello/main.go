// Package main demonstrates basic usage of the OpenZL Go bindings.
//
// This example shows how to:
// - Initialize an OpenZL context
// - Perform basic compression operations
// - Handle errors properly
// - Clean up resources
//
// To build and run:
//   go build -o hello main.go
//   ./hello
package main

import (
	"fmt"
	"log"

	"github.com/gus3inov/openzl-go/openzl"
)

func main() {
	fmt.Println("OpenZL Go Bindings - Hello World Example")

	// Initialize OpenZL context
	ctx, err := openzl.NewContext()
	if err != nil {
		log.Fatalf("Failed to initialize OpenZL context: %v", err)
	}
	defer ctx.Close()

	fmt.Println("✓ Successfully initialized OpenZL context")

	// Test data
	testData := []byte("Hello, OpenZL!")
	fmt.Printf("Original data: %s (%d bytes)\n", string(testData), len(testData))

	// Compress the data
	compressed, err := ctx.Compress(testData)
	if err != nil {
		log.Fatalf("Compression failed: %v", err)
	}

	compressionRatio := float64(len(compressed)) / float64(len(testData)) * 100
	fmt.Printf("✓ Compression test passed: %d bytes -> %d bytes (%.0f%% compression)\n",
		len(testData), len(compressed), compressionRatio)

	// Decompress the data
	decompressed, err := ctx.Decompress(compressed)
	if err != nil {
		log.Fatalf("Decompression failed: %v", err)
	}

	fmt.Printf("✓ Decompression test passed: %d bytes -> %d bytes\n",
		len(compressed), len(decompressed))

	// Verify data integrity
	if string(decompressed) != string(testData) {
		log.Fatalf("Data integrity check failed: expected %s, got %s",
			string(testData), string(decompressed))
	}

	fmt.Println("✓ All tests completed successfully!")
}
