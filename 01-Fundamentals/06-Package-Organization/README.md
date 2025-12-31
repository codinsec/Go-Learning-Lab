# Package Organization

Learn about package visibility, import/export rules, and how to organize Go code into packages.

## Learning Objectives

- Understand package visibility rules (uppercase vs lowercase)
- Learn how to create and use packages
- Understand the difference between exported and unexported identifiers
- Learn package naming conventions
- Understand how to structure multi-package projects

## Concepts Covered

### Package Visibility

In Go, visibility is determined by the first letter of an identifier:

- **Uppercase** = **Exported** (public, accessible from other packages)
- **Lowercase** = **Unexported** (private, only accessible within the same package)

### Example

```go
// calculator/calculator.go
package calculator

// Exported function (can be used by other packages)
func Add(a, b int) int {
    return a + b
}

// Unexported function (only within this package)
func add(a, b int) int {
    return a + b
}

// Exported type
type Calculator struct {
    result int  // unexported field
}

// Exported method
func (c *Calculator) GetResult() int {
    return c.result
}
```

### Package Structure

```
package-organization/
├── main.go              # Main package
├── calculator/
│   ├── calculator.go    # Calculator package
│   └── calculator_test.go
├── greeter/
│   ├── greeter.go       # Greeter package
│   └── greeter_test.go
└── go.mod
```

### Importing Packages

```go
import (
    "fmt"  // Standard library
    "github.com/codinsec/go-learning-lab/01-fundamentals/package-organization/calculator"
    "github.com/codinsec/go-learning-lab/01-fundamentals/package-organization/greeter"
)
```

## Running the Example

```bash
# Navigate to this directory
cd 06-Package-Organization

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/01-fundamentals/package-organization

# Download dependencies
go mod tidy

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests (from package-organization directory)
go test ./...

# Run tests for specific package
go test ./calculator
go test ./greeter

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## Expected Output

```
=== Package Organization ===

1. Using Exported Functions:
   calculator.Add(10, 20) = 30
   calculator.Multiply(5, 4) = 20

2. Using Exported Types:
   calc.Subtract(100, 30) = 70

3. Using Greeter Package:
   Hello, Go Developer!
   Goodbye, Go Developer!

4. Visibility Rules:
   - Uppercase names are EXPORTED (public)
   - Lowercase names are UNEXPORTED (private)
   - Exported: Add, Multiply, NewCalculator
   - Unexported: add, multiply, newCalculator (not accessible)
```

## Key Takeaways

1. **First letter determines visibility**: Uppercase = exported, lowercase = unexported
2. **Package = directory**: One package per directory
3. **Package name = directory name**: Usually matches the directory
4. **Main package**: Special package for executables
5. **Import path**: Full module path + package path

## Package Naming Conventions

- Use **lowercase** package names
- Use **single word** when possible
- Avoid **underscores** and **mixedCaps**
- Package name should match directory name
- Examples: `calculator`, `greeter`, `utils`, `models`

## Best Practices

- Keep packages focused and cohesive
- Export only what's necessary
- Use descriptive package names
- One package per directory
- Group related functionality together
- Keep packages small and focused

## Common Patterns

**Constructor functions:**
```go
func NewCalculator() *Calculator {
    return &Calculator{}
}
```

**Package-level variables:**
```go
var (
    // Exported
    DefaultValue = 10
    
    // Unexported
    defaultConfig = "config.json"
)
```

## Next Steps

Congratulations! You've completed Section 1: Fundamentals. 

Proceed to:
- [Section 2: Data Structures & Functions](../../02-Data-Structures-Functions/) - Learn about arrays, slices, maps, and functions

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

