# Hello World

Your first Go program! Learn how to write, compile, and run a simple Go program.

## Learning Objectives

- Write your first Go program
- Understand the `main` function and `main` package
- Learn basic Go commands: `go run`, `go build`
- Understand package imports

## Concepts Covered

### Package main
Every executable Go program must have a `package main` declaration and a `main()` function. This is the entry point of your program.

### Import Statement
Use `import` to include packages from the standard library or external modules. The `fmt` package provides formatted I/O functions.

### Go Commands

**`go run`** - Compiles and runs a Go program in one step
```bash
go run main.go
```

**`go build`** - Compiles a Go program into an executable
```bash
go build
go build -o hello.exe  # Windows
go build -o hello       # Linux/Mac
```

**`go install`** - Compiles and installs the package
```bash
go install
```

## Running the Example

```bash
# Navigate to this directory
cd 02-Hello-World

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/01-fundamentals/hello-world

# Run the program directly
go run main.go

# Or build it first, then run
go build
./hello.exe  # Windows
./hello      # Linux/Mac
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v
```

## Expected Output

```
Hello, World!
Welcome to Go Learning Lab!
Go version: 1.21+
Learning path: Fundamentals
```

## Key Takeaways

1. Every executable Go program needs `package main` and `main()` function
2. `go run` is great for quick testing
3. `go build` creates a standalone executable
4. Go programs compile to native binaries (no interpreter needed)

## Next Steps

After writing your first program, proceed to:
- [03-Package-Management](../03-Package-Management/README.md) - Learn about Go modules

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

