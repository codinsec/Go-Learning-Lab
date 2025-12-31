# Setup & Environment

This topic covers Go environment setup, understanding GOPATH, GOROOT, and Go workspace structure.

## Learning Objectives

- Understand Go installation and environment variables
- Learn about GOROOT and GOPATH concepts
- Explore Go workspace structure
- Get familiar with Go runtime information

## Concepts Covered

### GOROOT
GOROOT is the location where Go is installed. It contains the Go standard library and compiler.

### GOPATH
GOPATH is the workspace directory where Go code lives. In modern Go (1.11+), modules have largely replaced GOPATH, but it's still used for some tools.

### Go Workspace
A Go workspace is a directory hierarchy with three directories at its root:
- `src/` - contains Go source files
- `pkg/` - contains package objects
- `bin/` - contains executable commands

**Note:** With Go modules (Go 1.11+), you can work outside of GOPATH.

## Running the Example

```bash
# Navigate to this directory
cd 01-Setup-Environment

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/01-fundamentals/setup-environment

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover
```

## Expected Output

When you run `go run main.go`, you should see:
- Go version information
- GOROOT path
- GOPATH (if set)
- Operating system and architecture
- Number of CPUs

## Key Takeaways

1. **GOROOT** points to your Go installation
2. **GOPATH** is your workspace (less important with modules)
3. **Go modules** (go.mod) are the modern way to manage dependencies
4. Use `runtime` package to get environment information at runtime

## Next Steps

After understanding the environment setup, proceed to:
- [02-Hello-World](../02-Hello-World/README.md) - Write your first Go program

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

