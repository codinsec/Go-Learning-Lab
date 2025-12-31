# Go Tools

Learn about essential Go development tools for formatting, linting, testing, and more.

## Learning Objectives

- Understand core Go tools (fmt, vet, test, build)
- Learn advanced tools (golangci-lint, goimports, gopls)
- Master code formatting and linting
- Understand module management tools
- Learn testing and benchmarking tools
- Understand IDE integration tools

## Concepts Covered

### Core Tools

**go fmt** - Code formatting:
```bash
go fmt ./...
```

**go vet** - Static analysis:
```bash
go vet ./...
```

**go test** - Testing:
```bash
go test ./...
go test -cover ./...
go test -bench=. ./...
```

**go build** - Building:
```bash
go build
go build -ldflags="-X main.version=1.0.0"
```

### Advanced Tools

**golangci-lint** - Advanced linting:
```bash
golangci-lint run
```

**goimports** - Import management:
```bash
goimports -w .
```

**gopls** - Language Server Protocol:
```bash
# IDE integration for VS Code, etc.
```

## Running the Example

```bash
# Navigate to this directory
cd 08-Go-Tools

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/go-tools

# Run the program
go run main.go

# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Run tests
go test ./...
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run benchmarks
go test -bench=.
```

## Key Takeaways

1. **go fmt** - Always format code
2. **go vet** - Catch common errors
3. **golangci-lint** - Advanced linting
4. **goimports** - Manage imports automatically
5. **gopls** - IDE integration
6. **go test** - Comprehensive testing
7. **go mod** - Dependency management

## Common Tools

### Code Formatting

**go fmt:**
```bash
go fmt ./...
```

**goimports:**
```bash
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### Static Analysis

**go vet:**
```bash
go vet ./...
```

**golangci-lint:**
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run
```

### Testing

**go test:**
```bash
go test ./...
go test -cover ./...
go test -bench=. ./...
go test -race ./...  # Race detector
```

### Building

**go build:**
```bash
go build
go build -ldflags="-w -s"  # Strip debug info
go build -race  # Build with race detector
```

### Module Management

**go mod:**
```bash
go mod init <module-path>
go mod download
go mod tidy
go mod vendor
go mod verify
```

### Documentation

**go doc:**
```bash
go doc <package>
go doc <package>.<function>
godoc -http=:6060
```

## Best Practices

- **Always use go fmt** - Format code before committing
- **Run go vet** - Catch errors early
- **Use golangci-lint** - For comprehensive linting
- **Set up pre-commit hooks** - Automate tool running
- **Use gopls** - For better IDE experience
- **Run tests** - Before committing code
- **Check coverage** - Maintain test coverage

## Important Notes

- **go fmt** - Standard formatter, always use
- **go vet** - Built-in static analyzer
- **golangci-lint** - Combines multiple linters
- **goimports** - Automatically manages imports
- **gopls** - Language server for IDEs
- **go test** - Comprehensive testing framework
- **CI/CD integration** - Use tools in pipelines

## Tool Installation

### Core Tools (Built-in)
- `go fmt` - Included
- `go vet` - Included
- `go test` - Included
- `go build` - Included

### External Tools

**golangci-lint:**
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

**goimports:**
```bash
go install golang.org/x/tools/cmd/goimports@latest
```

**gopls:**
```bash
go install golang.org/x/tools/gopls@latest
```

**godoc:**
```bash
go install golang.org/x/tools/cmd/godoc@latest
```

## IDE Integration

### VS Code
- Install Go extension
- gopls is automatically used
- Format on save enabled

### GoLand
- Built-in Go support
- Integrated tools
- Advanced refactoring

### Vim/Neovim
- Use vim-go plugin
- Configure gopls
- Set up goimports

## CI/CD Integration

```yaml
# Example GitHub Actions
- name: Format check
  run: go fmt ./...

- name: Vet check
  run: go vet ./...

- name: Lint
  run: golangci-lint run

- name: Test
  run: go test -cover ./...
```

## Next Steps

After understanding Go tools, proceed to:
- [09-Package-Management](../09-Package-Management/README.md) - Learn about package management

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

