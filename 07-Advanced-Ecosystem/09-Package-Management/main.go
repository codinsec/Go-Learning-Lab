package main

import (
	"fmt"
	"os/exec"
)

// This program demonstrates Go package management

func main() {
	fmt.Println("=== Package Management ===")
	fmt.Println()

	// 1. go mod init
	fmt.Println("1. go mod init - Initialize Module:")
	fmt.Println("   Create a new module:")
	fmt.Println("   go mod init <module-path>")
	fmt.Println()
	fmt.Println("   Example:")
	fmt.Println("   go mod init github.com/user/project")
	fmt.Println()

	// 2. go get
	fmt.Println("2. go get - Add Dependencies:")
	fmt.Println("   Add latest version:")
	fmt.Println("   go get <package>")
	fmt.Println()
	fmt.Println("   Add specific version:")
	fmt.Println("   go get <package>@v1.2.3")
	fmt.Println()
	fmt.Println("   Add latest version:")
	fmt.Println("   go get <package>@latest")
	fmt.Println()
	fmt.Println("   Update to latest:")
	fmt.Println("   go get -u <package>")
	fmt.Println()
	fmt.Println("   Update all:")
	fmt.Println("   go get -u ./...")
	fmt.Println()

	// 3. go mod download
	fmt.Println("3. go mod download - Download Dependencies:")
	fmt.Println("   Download all dependencies:")
	fmt.Println("   go mod download")
	fmt.Println()
	fmt.Println("   Download specific module:")
	fmt.Println("   go mod download <module>")
	fmt.Println()

	// 4. go mod tidy
	fmt.Println("4. go mod tidy - Clean Dependencies:")
	fmt.Println("   Add missing and remove unused:")
	fmt.Println("   go mod tidy")
	fmt.Println()
	fmt.Println("   This command:")
	fmt.Println("   - Adds missing dependencies")
	fmt.Println("   - Removes unused dependencies")
	fmt.Println("   - Updates go.sum")
	fmt.Println()

	// 5. go mod vendor
	fmt.Println("5. go mod vendor - Vendor Dependencies:")
	fmt.Println("   Create vendor directory:")
	fmt.Println("   go mod vendor")
	fmt.Println()
	fmt.Println("   Use vendored dependencies:")
	fmt.Println("   go build -mod=vendor")
	fmt.Println()
	fmt.Println("   Benefits:")
	fmt.Println("   - Reproducible builds")
	fmt.Println("   - Offline builds")
	fmt.Println("   - Version control")
	fmt.Println()

	// 6. go mod verify
	fmt.Println("6. go mod verify - Verify Dependencies:")
	fmt.Println("   Verify module checksums:")
	fmt.Println("   go mod verify")
	fmt.Println()
	fmt.Println("   Ensures:")
	fmt.Println("   - Dependencies haven't changed")
	fmt.Println("   - Checksums match go.sum")
	fmt.Println()

	// 7. go mod edit
	fmt.Println("7. go mod edit - Edit go.mod:")
	fmt.Println("   Add replace directive:")
	fmt.Println("   go mod edit -replace <old>=<new>")
	fmt.Println()
	fmt.Println("   Remove replace directive:")
	fmt.Println("   go mod edit -dropreplace <module>")
	fmt.Println()
	fmt.Println("   Set go version:")
	fmt.Println("   go mod edit -go=1.21")
	fmt.Println()

	// 8. go mod graph
	fmt.Println("8. go mod graph - Dependency Graph:")
	fmt.Println("   Show dependency graph:")
	fmt.Println("   go mod graph")
	fmt.Println()
	fmt.Println("   Visualize with:")
	fmt.Println("   go mod graph | dot -Tpng -o graph.png")
	fmt.Println()

	// 9. go mod why
	fmt.Println("9. go mod why - Why Dependency:")
	fmt.Println("   Explain why a module is needed:")
	fmt.Println("   go mod why <module>")
	fmt.Println()
	fmt.Println("   Shows dependency chain")
	fmt.Println()

	// 10. Version management
	fmt.Println("10. Version Management:")
	fmt.Println("    Semantic versioning:")
	fmt.Println("    - v1.2.3 (major.minor.patch)")
	fmt.Println("    - v0.1.0 (pre-1.0)")
	fmt.Println("    - v1.2.3-beta.1 (pre-release)")
	fmt.Println()
	fmt.Println("    Pseudo-versions:")
	fmt.Println("    - v0.0.0-20230101120000-abc123def456")
	fmt.Println()

	// 11. Private repositories
	fmt.Println("11. Private Repositories:")
	fmt.Println("    Configure .netrc or .gitconfig:")
	fmt.Println("    machine github.com")
	fmt.Println("    login <token>")
	fmt.Println()
	fmt.Println("    Or use GOPRIVATE:")
	fmt.Println("    go env -w GOPRIVATE=github.com/company/*")
	fmt.Println()

	// 12. Proxy and checksum database
	fmt.Println("12. Proxy and Checksum Database:")
	fmt.Println("    GOPROXY (default):")
	fmt.Println("    https://proxy.golang.org,direct")
	fmt.Println()
	fmt.Println("    GOSUMDB (default):")
	fmt.Println("    sum.golang.org")
	fmt.Println()
	fmt.Println("    Disable for private repos:")
	fmt.Println("    go env -w GOSUMDB=off")
	fmt.Println()

	// 13. go install
	fmt.Println("13. go install - Install Binaries:")
	fmt.Println("    Install from module:")
	fmt.Println("    go install <module>/cmd/tool@latest")
	fmt.Println()
	fmt.Println("    Install from local:")
	fmt.Println("    go install ./cmd/myapp")
	fmt.Println()

	// 14. Key concepts
	fmt.Println("14. Key Concepts:")
	fmt.Println("    - go.mod tracks dependencies")
	fmt.Println("    - go.sum ensures integrity")
	fmt.Println("    - Semantic versioning")
	fmt.Println("    - Module proxy for downloads")
	fmt.Println("    - Checksum database for security")
	fmt.Println("    - Vendoring for reproducibility")
	fmt.Println()

	// Demonstrate commands
	demonstrateCommands()
}

func demonstrateCommands() {
	fmt.Println("=== Command Demonstrations ===")
	fmt.Println()

	// Check go.mod exists
	fmt.Println("Module Information:")
	if cmd, err := exec.Command("go", "list", "-m").Output(); err == nil {
		fmt.Printf("   Current module: %s", string(cmd))
	} else {
		fmt.Println("   Could not get module info")
	}
	fmt.Println()

	// List dependencies
	fmt.Println("Dependencies:")
	if cmd, err := exec.Command("go", "list", "-m", "all").Output(); err == nil {
		lines := string(cmd)
		if len(lines) > 0 {
			fmt.Printf("   %s", lines[:min(200, len(lines))])
			if len(lines) > 200 {
				fmt.Println("   ... (truncated)")
			}
		}
	} else {
		fmt.Println("   Could not list dependencies")
	}
	fmt.Println()

	// Go environment
	fmt.Println("Go Module Environment:")
	envVars := []string{"GOPROXY", "GOSUMDB", "GOPRIVATE", "GOMOD"}
	for _, env := range envVars {
		if cmd, err := exec.Command("go", "env", env).Output(); err == nil {
			fmt.Printf("   %s: %s", env, string(cmd))
		}
	}
	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

