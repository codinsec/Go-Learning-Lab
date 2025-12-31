package main

import (
	"fmt"
	"runtime"
)

// This program demonstrates cross-compilation in Go

func main() {
	fmt.Println("=== Cross-Compilation ===")
	fmt.Println()

	// 1. Current platform
	fmt.Println("1. Current Platform:")
	fmt.Printf("   OS: %s\n", runtime.GOOS)
	fmt.Printf("   Architecture: %s\n", runtime.GOARCH)
	fmt.Println()

	// 2. Supported platforms
	fmt.Println("2. Supported Platforms:")
	fmt.Println("   Operating Systems:")
	fmt.Println("   - linux, windows, darwin, freebsd, openbsd")
	fmt.Println("   - netbsd, solaris, plan9, aix, dragonfly")
	fmt.Println()
	fmt.Println("   Architectures:")
	fmt.Println("   - amd64, 386, arm, arm64, mips, mips64")
	fmt.Println("   - mips64le, mipsle, ppc64, ppc64le")
	fmt.Println("   - riscv64, s390x, wasm")
	fmt.Println()

	// 3. Cross-compilation commands
	fmt.Println("3. Cross-Compilation Commands:")
	fmt.Println("   Linux (amd64):")
	fmt.Println("   GOOS=linux GOARCH=amd64 go build")
	fmt.Println()
	fmt.Println("   Windows (amd64):")
	fmt.Println("   GOOS=windows GOARCH=amd64 go build")
	fmt.Println()
	fmt.Println("   macOS (Intel):")
	fmt.Println("   GOOS=darwin GOARCH=amd64 go build")
	fmt.Println()
	fmt.Println("   macOS (Apple Silicon):")
	fmt.Println("   GOOS=darwin GOARCH=arm64 go build")
	fmt.Println()
	fmt.Println("   ARM Linux:")
	fmt.Println("   GOOS=linux GOARCH=arm64 go build")
	fmt.Println()

	// 4. Build for multiple platforms
	fmt.Println("4. Build for Multiple Platforms:")
	fmt.Println("   Create a build script:")
	fmt.Println("   #!/bin/bash")
	fmt.Println("   GOOS=linux GOARCH=amd64 go build -o app-linux-amd64")
	fmt.Println("   GOOS=windows GOARCH=amd64 go build -o app-windows-amd64.exe")
	fmt.Println("   GOOS=darwin GOARCH=amd64 go build -o app-darwin-amd64")
	fmt.Println("   GOOS=darwin GOARCH=arm64 go build -o app-darwin-arm64")
	fmt.Println()

	// 5. CGO and cross-compilation
	fmt.Println("5. CGO and Cross-Compilation:")
	fmt.Println("   CGO makes cross-compilation more complex:")
	fmt.Println("   - Need cross-compiler toolchain")
	fmt.Println("   - Set CC environment variable")
	fmt.Println("   - Example:")
	fmt.Println("     CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm go build")
	fmt.Println()

	// 6. Build tags for cross-compilation
	fmt.Println("6. Build Tags for Cross-Compilation:")
	fmt.Println("   Use build tags for platform-specific code:")
	fmt.Println("   //go:build linux")
	fmt.Println("   //go:build windows")
	fmt.Println("   //go:build darwin")
	fmt.Println()

	// 7. Testing cross-compiled binaries
	fmt.Println("7. Testing Cross-Compiled Binaries:")
	fmt.Println("   - Test on target platform")
	fmt.Println("   - Use emulation (QEMU)")
	fmt.Println("   - Use CI/CD with multiple platforms")
	fmt.Println("   - Use Docker for testing")
	fmt.Println()

	// 8. Common use cases
	fmt.Println("8. Common Use Cases:")
	fmt.Println("   - Build for multiple platforms from one machine")
	fmt.Println("   - CI/CD pipelines")
	fmt.Println("   - Release distributions")
	fmt.Println("   - Embedded systems")
	fmt.Println("   - Cloud deployments")
	fmt.Println()

	// 9. Environment variables
	fmt.Println("9. Environment Variables:")
	fmt.Println("   GOOS      - Target operating system")
	fmt.Println("   GOARCH    - Target architecture")
	fmt.Println("   CGO_ENABLED - Enable/disable CGO (0 or 1)")
	fmt.Println("   CC        - C compiler for CGO")
	fmt.Println()

	// 10. Key concepts
	fmt.Println("10. Key Concepts:")
	fmt.Println("   - Go supports easy cross-compilation")
	fmt.Println("   - Set GOOS and GOARCH environment variables")
	fmt.Println("   - CGO complicates cross-compilation")
	fmt.Println("   - Test on target platform when possible")
	fmt.Println("   - Use build tags for platform-specific code")
	fmt.Println()

	// 11. Example build script
	fmt.Println("11. Example Build Script (build.sh):")
	fmt.Println("   #!/bin/bash")
	fmt.Println("   set -e")
	fmt.Println("   VERSION=$(git describe --tags --always --dirty)")
	fmt.Println("   ")
	fmt.Println("   echo \"Building version $VERSION\"")
	fmt.Println("   ")
	fmt.Println("   GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.version=$VERSION\" -o dist/app-linux-amd64")
	fmt.Println("   GOOS=windows GOARCH=amd64 go build -ldflags \"-X main.version=$VERSION\" -o dist/app-windows-amd64.exe")
	fmt.Println("   GOOS=darwin GOARCH=amd64 go build -ldflags \"-X main.version=$VERSION\" -o dist/app-darwin-amd64")
	fmt.Println("   GOOS=darwin GOARCH=arm64 go build -ldflags \"-X main.version=$VERSION\" -o dist/app-darwin-arm64")
	fmt.Println("   ")
	fmt.Println("   echo \"Build complete!\"")
	fmt.Println()
}

