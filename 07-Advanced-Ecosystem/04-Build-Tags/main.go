package main

import (
	"fmt"
	"runtime"
)

// This program demonstrates build tags in Go

// Build tags are specified at the top of the file with:
// //go:build tag1 && tag2
// or
// // +build tag1 tag2

func main() {
	fmt.Println("=== Build Tags ===")
	fmt.Println()

	// 1. Platform detection
	fmt.Println("1. Platform Detection:")
	fmt.Printf("   OS: %s\n", runtime.GOOS)
	fmt.Printf("   Architecture: %s\n", runtime.GOARCH)
	fmt.Println()

	// 2. Build tag syntax
	fmt.Println("2. Build Tag Syntax:")
	fmt.Println("   //go:build tag1 && tag2")
	fmt.Println("   //go:build tag1 || tag2")
	fmt.Println("   //go:build !tag1")
	fmt.Println("   //go:build tag1 && !tag2")
	fmt.Println()

	// 3. Common build tags
	fmt.Println("3. Common Build Tags:")
	fmt.Println("   - linux, windows, darwin (OS)")
	fmt.Println("   - amd64, arm64, 386 (architecture)")
	fmt.Println("   - debug, production (custom)")
	fmt.Println("   - integration, unit (test tags)")
	fmt.Println()

	// 4. Example: Platform-specific code
	fmt.Println("4. Platform-Specific Code:")
	fmt.Println("   Create files with build tags:")
	fmt.Println("   - file_linux.go   //go:build linux")
	fmt.Println("   - file_windows.go //go:build windows")
	fmt.Println("   - file_darwin.go  //go:build darwin")
	fmt.Println()

	// 5. Build tag example files
	fmt.Println("5. Build Tag Example Files:")
	fmt.Println("   See:")
	fmt.Println("   - platform_linux.go")
	fmt.Println("   - platform_windows.go")
	fmt.Println("   - platform_darwin.go")
	fmt.Println()

	// 6. Using build tags
	fmt.Println("6. Using Build Tags:")
	fmt.Println("   Build for specific platform:")
	fmt.Println("   GOOS=linux GOARCH=amd64 go build")
	fmt.Println("   GOOS=windows GOARCH=amd64 go build")
	fmt.Println("   GOOS=darwin GOARCH=arm64 go build")
	fmt.Println()

	// 7. Test tags
	fmt.Println("7. Test Tags:")
	fmt.Println("   Run tests with tags:")
	fmt.Println("   go test -tags=integration")
	fmt.Println("   go test -tags=unit")
	fmt.Println()

	// 8. Custom build tags
	fmt.Println("8. Custom Build Tags:")
	fmt.Println("   Build with custom tags:")
	fmt.Println("   go build -tags=debug")
	fmt.Println("   go build -tags=production")
	fmt.Println()

	// 9. Multiple tags
	fmt.Println("9. Multiple Tags:")
	fmt.Println("   Combine tags:")
	fmt.Println("   go build -tags='debug,linux'")
	fmt.Println("   go test -tags='integration,slow'")
	fmt.Println()

	// 10. Key concepts
	fmt.Println("10. Key Concepts:")
	fmt.Println("   - Build tags enable conditional compilation")
	fmt.Println("   - Use for platform-specific code")
	fmt.Println("   - Use for feature flags")
	fmt.Println("   - Use for test categories")
	fmt.Println("   - Must be at top of file")
	fmt.Println("   - Space after // is required")
}

