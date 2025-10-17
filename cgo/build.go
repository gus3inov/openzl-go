//go:build cgo
// +build cgo

// Package cgo contains the cgo layer and build glue for OpenZL.
package cgo

/*
#cgo CFLAGS: -I${SRCDIR}/../third_party/openzl/include
#cgo LDFLAGS: -L${SRCDIR}/../third_party/openzl/build -lopenzl -Wl,-rpath,${SRCDIR}/../third_party/openzl/build

#include "openzl.h"
*/
import "C"

// This file contains build constraints and flags for cgo.
