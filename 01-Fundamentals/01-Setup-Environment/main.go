package main

import (
	"fmt"
	"os"
	"runtime"
)

// This program demonstrates Go environment setup and basic environment information
func main() {
	fmt.Println("=== Go Environment Setup ===")
	fmt.Println()

	// Display Go version
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Println()

	// Display Go root directory (GOROOT)
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = runtime.GOROOT()
	}
	fmt.Printf("GOROOT: %s\n", goroot)
	fmt.Println()

	// Display Go path (GOPATH)
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = "Not set (using default)"
	}
	fmt.Printf("GOPATH: %s\n", gopath)
	fmt.Println()

	// Display operating system and architecture
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Println()

	// Display number of CPUs
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Println()

	fmt.Println("✅ Go environment is properly configured!")
}

