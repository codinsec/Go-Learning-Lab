# Build Tags

Learn about build tags in Go for conditional compilation and platform-specific code.

## Learning Objectives

- Understand what build tags are and when to use them
- Learn build tag syntax
- Master platform-specific code organization
- Understand test tags
- Learn custom build tags
- Understand build tag best practices

## Concepts Covered

### Build Tag Syntax

Build tags are specified at the top of the file:

```go
//go:build linux
package main
```

### Platform-Specific Files

Create separate files for different platforms:

- `platform_linux.go` - Linux-specific code
- `platform_windows.go` - Windows-specific code
- `platform_darwin.go` - macOS-specific code

### Build Tag Operators

```go
//go:build tag1 && tag2  // Both tags required
//go:build tag1 || tag2  // Either tag
//go:build !tag1         // Not tag1
//go:build tag1 && !tag2 // tag1 but not tag2
```

### Common Build Tags

- **OS tags**: `linux`, `windows`, `darwin`, `freebsd`
- **Architecture tags**: `amd64`, `arm64`, `386`, `arm`
- **Custom tags**: `debug`, `production`, `integration`

## Running the Example

```bash
# Navigate to this directory
cd 04-Build-Tags

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/build-tags

# Run the program
go run main.go

# Build for specific platform
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=darwin GOARCH=arm64 go build
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with tags
go test -tags=integration
go test -tags=unit

# Run tests with multiple tags
go test -tags='integration,slow'
```

## Key Takeaways

1. **Conditional compilation** - Include/exclude code at build time
2. **Platform-specific code** - Different implementations per platform
3. **Feature flags** - Enable/disable features
4. **Test organization** - Separate unit and integration tests
5. **Build control** - Fine-grained control over what gets compiled
6. **File naming** - Use descriptive names with platform suffix

## Common Patterns

**Platform-specific implementation:**
```go
// file_linux.go
//go:build linux
package main

func doSomething() {
    // Linux-specific code
}
```

**Feature flags:**
```go
// feature_debug.go
//go:build debug
package main

func debugLog(msg string) {
    log.Println("[DEBUG]", msg)
}
```

**Test tags:**
```go
// integration_test.go
//go:build integration
package main

func TestIntegration(t *testing.T) {
    // Integration test code
}
```

## Best Practices

- **Use descriptive names** - `platform_linux.go` not `p_linux.go`
- **Keep platform code separate** - One file per platform
- **Document build tags** - Explain what each tag does
- **Test all platforms** - Ensure code works on all targets
- **Use standard tags** - Prefer standard OS/arch tags
- **Avoid overuse** - Don't create too many custom tags

## Important Notes

- **File placement** - Build tag must be at top of file
- **Space required** - Must have space after `//`
- **New syntax** - `//go:build` is preferred over `// +build`
- **Default build** - Files without tags are always included
- **Tag evaluation** - Tags are evaluated at build time

## Build Tag Examples

### Example 1: Platform-Specific Code

```go
// unix.go
//go:build linux || darwin || freebsd
package main

func getHomeDir() string {
    return os.Getenv("HOME")
}
```

```go
// windows.go
//go:build windows
package main

func getHomeDir() string {
    return os.Getenv("USERPROFILE")
}
```

### Example 2: Feature Flags

```go
// debug.go
//go:build debug
package main

var debugMode = true
```

```go
// production.go
//go:build !debug
package main

var debugMode = false
```

### Example 3: Test Categories

```go
// unit_test.go
//go:build !integration
package main

func TestUnit(t *testing.T) {
    // Fast unit tests
}
```

```go
// integration_test.go
//go:build integration
package main

func TestIntegration(t *testing.T) {
    // Slow integration tests
}
```

## Cross-Compilation

Build for different platforms:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build

# Windows
GOOS=windows GOARCH=amd64 go build

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build

# ARM Linux
GOOS=linux GOARCH=arm64 go build
```

## Next Steps

After understanding build tags, proceed to:
- [05-CGO](../05-CGO/README.md) - Learn about CGO

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

