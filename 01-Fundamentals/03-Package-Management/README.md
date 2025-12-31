# Package Management

Learn about Go modules, the modern way to manage dependencies in Go.

## Learning Objectives

- Understand Go modules and `go.mod` file
- Learn how to initialize a module
- Understand module paths and versioning
- Learn essential `go mod` commands

## Concepts Covered

### Go Modules
Go modules are collections of Go packages stored in a file tree with a `go.mod` file at its root. Modules were introduced in Go 1.11 and became the standard in Go 1.13.

### go.mod File
The `go.mod` file defines:
- Module path (like `github.com/user/project`)
- Go version requirement
- Direct dependencies (in `go.sum` for checksums)

### Module Path
A module path is the canonical name for your module, typically:
- GitHub: `github.com/username/repo`
- Custom domain: `example.com/project`
- Local: `local/path/to/module`

## Essential Commands

### Initialize a Module
```bash
go mod init <module-path>
```

### Add Dependencies
```bash
go get <package-path>
go get github.com/gin-gonic/gin@v1.9.1  # Specific version
go get github.com/gin-gonic/gin@latest  # Latest version
```

### Update Dependencies
```bash
go get -u ./...              # Update all dependencies
go get -u <package-path>     # Update specific package
go mod tidy                  # Clean up dependencies
```

### Download Dependencies
```bash
go mod download              # Download to cache
go mod vendor                # Create vendor directory
```

### Verify Dependencies
```bash
go mod verify                # Verify module checksums
```

## Running the Example

```bash
# Navigate to this directory
cd 03-Package-Management

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/01-fundamentals/package-management

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v
```

## go.mod File Structure

```go
module github.com/codinsec/go-learning-lab/01-fundamentals/package-management

go 1.21

require (
    // Direct dependencies are listed here
)

replace (
    // Local replacements (for development)
)
```

## Key Takeaways

1. **Go modules** replaced GOPATH-based dependency management
2. **go.mod** defines your module and its dependencies
3. **go.sum** contains checksums for dependency verification
4. Module paths should be unique and follow conventions
5. Use `go mod tidy` regularly to keep dependencies clean

## Best Practices

- Always commit `go.mod` and `go.sum` to version control
- Use semantic versioning for your modules
- Run `go mod tidy` before committing
- Don't commit `vendor/` directory unless necessary
- Use specific versions in production (`@v1.2.3`)

## Next Steps

After understanding package management, proceed to:
- [04-Variables-Types](../04-Variables-Types/README.md) - Learn about variables and types

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

