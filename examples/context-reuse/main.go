// Package main demonstrates the performance benefits of context reuse in OpenZL.
//
// This example compares the performance of creating a new context for each
// compression operation versus reusing a single context. The results show
// that context reuse provides approximately 27% better performance.
//
// To build and run:
//   go build -o context-reuse main.go
//   ./context-reuse
package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gus3inov/openzl-go/openzl"
)

func main() {
	fmt.Println("OpenZL Context Reuse Performance Demonstration")
	fmt.Println("===============================================\n")

	// Prepare test data
	data := bytes.Repeat([]byte("Performance test data for OpenZL context reuse demonstration. "), 1000)
	iterations := 1000

	// Test 1: Creating new context for each operation
	fmt.Println("Test 1: Creating new context for each operation")
	start := time.Now()
	for i := 0; i < iterations; i++ {
		ctx, err := openzl.NewContext()
		if err != nil {
			panic(err)
		}
		_, err = ctx.Compress(data)
		if err != nil {
			panic(err)
		}
		ctx.Close()
	}
	timeWithNewContext := time.Since(start)
	fmt.Printf("  Time: %v\n", timeWithNewContext)
	fmt.Printf("  Avg per operation: %v\n\n", timeWithNewContext/time.Duration(iterations))

	// Test 2: Reusing single context
	fmt.Println("Test 2: Reusing single context")
	ctx, err := openzl.NewContext()
	if err != nil {
		panic(err)
	}
	defer ctx.Close()

	start = time.Now()
	for i := 0; i < iterations; i++ {
		_, err := ctx.Compress(data)
		if err != nil {
			panic(err)
		}
	}
	timeWithReuse := time.Since(start)
	fmt.Printf("  Time: %v\n", timeWithReuse)
	fmt.Printf("  Avg per operation: %v\n\n", timeWithReuse/time.Duration(iterations))

	// Calculate improvement
	improvement := float64(timeWithNewContext-timeWithReuse) / float64(timeWithNewContext) * 100
	speedup := float64(timeWithNewContext) / float64(timeWithReuse)

	fmt.Println("Results:")
	fmt.Println("--------")
	fmt.Printf("Performance improvement: %.1f%%\n", improvement)
	fmt.Printf("Speedup factor: %.2fx\n", speedup)
	fmt.Println("\nRecommendation: Always reuse contexts when performing multiple operations!")
}
