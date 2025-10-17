// Package openzl provides Go bindings for Facebook's OpenZL library.
//
// This package offers idiomatic Go APIs for compression and machine learning
// workloads using the OpenZL library.
//
// Example usage:
//
//	ctx, err := openzl.NewContext()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer ctx.Close()
//
//	data := []byte("Hello, World!")
//	compressed, err := ctx.Compress(data)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	decompressed, err := ctx.Decompress(compressed)
//	if err != nil {
//		log.Fatal(err)
//	}
package openzl
