package main

import (
	"fmt"
	"os"
)

// This program demonstrates Go module management
func main() {
	fmt.Println("=== Go Package Management ===")
	fmt.Println()

	// Check if go.mod exists
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("✅ go.mod file exists")
		fmt.Println("   This file defines the module path and Go version")
	} else {
		fmt.Println("❌ go.mod file not found")
		fmt.Println("   Run: go mod init <module-path>")
	}

	fmt.Println()
	fmt.Println("Module Management Commands:")
	fmt.Println("  go mod init <module-path>  - Initialize a new module")
	fmt.Println("  go mod tidy                - Add missing and remove unused modules")
	fmt.Println("  go mod download            - Download modules to local cache")
	fmt.Println("  go mod vendor              - Create vendor directory")
	fmt.Println("  go mod verify              - Verify dependencies")
	fmt.Println()
	fmt.Println("Example module path:")
	fmt.Println("  github.com/username/project-name")
}

