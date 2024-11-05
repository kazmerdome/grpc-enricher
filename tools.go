//go:build tools
// +build tools

// This file is used to track development dependencies.
// This ensures `go mod` can detect and include them in the project's dependencies.

package tools

import (
	_ "github.com/vektra/mockery/v2"
)
