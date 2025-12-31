package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// This program demonstrates Go tools and their usage

func main() {
	fmt.Println("=== Go Tools ===")
	fmt.Println()

	// 1. go fmt
	fmt.Println("1. go fmt - Code Formatting:")
	fmt.Println("   Format all files:")
	fmt.Println("   go fmt ./...")
	fmt.Println()
	fmt.Println("   Format specific file:")
	fmt.Println("   go fmt main.go")
	fmt.Println()
	fmt.Println("   Auto-format on save in most IDEs")
	fmt.Println()

	// 2. go vet
	fmt.Println("2. go vet - Static Analysis:")
	fmt.Println("   Check for common errors:")
	fmt.Println("   go vet ./...")
	fmt.Println()
	fmt.Println("   Checks for:")
	fmt.Println("   - Unreachable code")
	fmt.Println("   - Unused variables")
	fmt.Println("   - Incorrect printf format strings")
	fmt.Println("   - Suspicious constructs")
	fmt.Println()

	// 3. golangci-lint
	fmt.Println("3. golangci-lint - Advanced Linting:")
	fmt.Println("   Install:")
	fmt.Println("   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest")
	fmt.Println()
	fmt.Println("   Run:")
	fmt.Println("   golangci-lint run")
	fmt.Println()
	fmt.Println("   Features:")
	fmt.Println("   - Multiple linters in one tool")
	fmt.Println("   - Fast execution")
	fmt.Println("   - Configurable")
	fmt.Println()

	// 4. goimports
	fmt.Println("4. goimports - Import Management:")
	fmt.Println("   Install:")
	fmt.Println("   go install golang.org/x/tools/cmd/goimports@latest")
	fmt.Println()
	fmt.Println("   Format and fix imports:")
	fmt.Println("   goimports -w .")
	fmt.Println()
	fmt.Println("   Features:")
	fmt.Println("   - Adds missing imports")
	fmt.Println("   - Removes unused imports")
	fmt.Println("   - Formats code")
	fmt.Println()

	// 5. gopls (Language Server)
	fmt.Println("5. gopls - Language Server Protocol:")
	fmt.Println("   Install:")
	fmt.Println("   go install golang.org/x/tools/gopls@latest")
	fmt.Println()
	fmt.Println("   Features:")
	fmt.Println("   - IDE integration (VS Code, etc.)")
	fmt.Println("   - Auto-completion")
	fmt.Println("   - Go to definition")
	fmt.Println("   - Refactoring")
	fmt.Println("   - Diagnostics")
	fmt.Println()

	// 6. go mod
	fmt.Println("6. go mod - Module Management:")
	fmt.Println("   Initialize module:")
	fmt.Println("   go mod init <module-path>")
	fmt.Println()
	fmt.Println("   Download dependencies:")
	fmt.Println("   go mod download")
	fmt.Println()
	fmt.Println("   Tidy dependencies:")
	fmt.Println("   go mod tidy")
	fmt.Println()
	fmt.Println("   Vendor dependencies:")
	fmt.Println("   go mod vendor")
	fmt.Println()

	// 7. go test
	fmt.Println("7. go test - Testing:")
	fmt.Println("   Run all tests:")
	fmt.Println("   go test ./...")
	fmt.Println()
	fmt.Println("   Run with coverage:")
	fmt.Println("   go test -cover ./...")
	fmt.Println()
	fmt.Println("   Run benchmarks:")
	fmt.Println("   go test -bench=. ./...")
	fmt.Println()
	fmt.Println("   Generate coverage report:")
	fmt.Println("   go test -coverprofile=coverage.out")
	fmt.Println("   go tool cover -html=coverage.out")
	fmt.Println()

	// 8. go build
	fmt.Println("8. go build - Building:")
	fmt.Println("   Build current package:")
	fmt.Println("   go build")
	fmt.Println()
	fmt.Println("   Build with flags:")
	fmt.Println("   go build -ldflags=\"-X main.version=1.0.0\"")
	fmt.Println()
	fmt.Println("   Build for different platform:")
	fmt.Println("   GOOS=linux GOARCH=amd64 go build")
	fmt.Println()

	// 9. go run
	fmt.Println("9. go run - Running:")
	fmt.Println("   Run Go program:")
	fmt.Println("   go run main.go")
	fmt.Println()
	fmt.Println("   Run multiple files:")
	fmt.Println("   go run *.go")
	fmt.Println()

	// 10. go install
	fmt.Println("10. go install - Installing:")
	fmt.Println("    Install binary:")
	fmt.Println("    go install ./cmd/myapp")
	fmt.Println()
	fmt.Println("    Install from remote:")
	fmt.Println("    go install github.com/user/repo/cmd/tool@latest")
	fmt.Println()

	// 11. go doc
	fmt.Println("11. go doc - Documentation:")
	fmt.Println("    View package documentation:")
	fmt.Println("    go doc <package>")
	fmt.Println()
	fmt.Println("    View function documentation:")
	fmt.Println("    go doc <package>.<function>")
	fmt.Println()
	fmt.Println("    Serve documentation server:")
	fmt.Println("    godoc -http=:6060")
	fmt.Println()

	// 12. go generate
	fmt.Println("12. go generate - Code Generation:")
	fmt.Println("    Run generators:")
	fmt.Println("    go generate ./...")
	fmt.Println()
	fmt.Println("    Use in code:")
	fmt.Println("    //go:generate stringer -type=Status")
	fmt.Println()

	// 13. go clean
	fmt.Println("13. go clean - Clean Build Artifacts:")
	fmt.Println("    Remove build artifacts:")
	fmt.Println("    go clean")
	fmt.Println()
	fmt.Println("    Remove test cache:")
	fmt.Println("    go clean -testcache")
	fmt.Println()
	fmt.Println("    Remove module cache:")
	fmt.Println("    go clean -modcache")
	fmt.Println()

	// 14. go list
	fmt.Println("14. go list - List Packages:")
	fmt.Println("    List all packages:")
	fmt.Println("    go list ./...")
	fmt.Println()
	fmt.Println("    List with details:")
	fmt.Println("    go list -json ./...")
	fmt.Println()
	fmt.Println("    List dependencies:")
	fmt.Println("    go list -m all")
	fmt.Println()

	// 15. go version
	fmt.Println("15. go version - Version Info:")
	fmt.Println("    Check Go version:")
	fmt.Println("    go version")
	fmt.Println()
	fmt.Println("    Check environment:")
	fmt.Println("    go env")
	fmt.Println()

	// 16. Key concepts
	fmt.Println("16. Key Concepts:")
	fmt.Println("    - Use go fmt for formatting")
	fmt.Println("    - Use go vet for static analysis")
	fmt.Println("    - Use golangci-lint for advanced linting")
	fmt.Println("    - Use goimports for import management")
	fmt.Println("    - Use gopls for IDE support")
	fmt.Println("    - Integrate tools in CI/CD")
	fmt.Println()

	// Demonstrate some tools
	demonstrateTools()
}

func demonstrateTools() {
	fmt.Println("=== Tool Demonstrations ===")
	fmt.Println()

	// Check Go version
	fmt.Println("Go Version:")
	if cmd, err := exec.Command("go", "version").Output(); err == nil {
		fmt.Printf("   %s", string(cmd))
	} else {
		fmt.Println("   go command not found")
	}
	fmt.Println()

	// Check Go environment
	fmt.Println("Go Environment (GOPATH, GOROOT):")
	if cmd, err := exec.Command("go", "env", "GOPATH", "GOROOT").Output(); err == nil {
		fmt.Printf("   %s", string(cmd))
	} else {
		fmt.Println("   go command not found")
	}
	fmt.Println()

	// List current directory packages
	fmt.Println("Packages in current directory:")
	if cmd, err := exec.Command("go", "list", ".").Output(); err == nil {
		fmt.Printf("   %s", string(cmd))
	} else {
		fmt.Println("   Could not list packages")
	}
	fmt.Println()

	// Show runtime info
	fmt.Println("Runtime Information:")
	fmt.Printf("   OS: %s\n", runtime.GOOS)
	fmt.Printf("   Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("   Go Version: %s\n", runtime.Version())
	fmt.Println()
}

