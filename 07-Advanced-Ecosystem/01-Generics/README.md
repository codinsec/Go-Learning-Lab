# Generics

Learn about generics in Go, introduced in Go 1.18, for type-safe code reuse.

## Learning Objectives

- Understand what generics are and when to use them
- Learn generic function syntax
- Master type constraints and type parameters
- Understand generic structs and methods
- Learn about type inference
- Understand best practices

## Concepts Covered

### Generic Functions

Functions that work with multiple types:

```go
func Print[T any](value T) {
    fmt.Println(value)
}

Print("hello")  // Works with string
Print(42)       // Works with int
```

### Type Parameters

```go
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

### Type Constraints

```go
type Number interface {
    int | int64 | float64
}

func Add[T Number](a, b T) T {
    return a + b
}
```

### Generic Structs

```go
type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{items: make([]T, 0)}
}
```

### Type Set Notation

```go
type Numeric interface {
    ~int | ~int64 | ~float64  // ~ allows underlying types
}
```

## Running the Example

```bash
# Navigate to this directory
cd 01-Generics

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/generics

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

## Key Takeaways

1. **Generics enable type-safe reuse** - Write once, use with multiple types
2. **Type parameters** - `[T any]` syntax
3. **Type constraints** - Limit which types can be used
4. **Type inference** - Go infers types when possible
5. **Available in Go 1.18+** - Requires Go 1.18 or later
6. **Use when appropriate** - Don't overuse generics

## Common Patterns

**Generic function:**
```go
func Identity[T any](value T) T {
    return value
}
```

**Generic struct:**
```go
type Queue[T any] struct {
    items []T
}
```

**Type constraint:**
```go
type Addable interface {
    int | float64
}
```

## Best Practices

- **Use when needed** - Don't overuse generics
- **Keep it simple** - Generics can make code complex
- **Use constraints** - Limit types appropriately
- **Type inference** - Let Go infer when possible
- **Document constraints** - Make them clear

## Important Notes

- **Go 1.18+** - Generics require Go 1.18 or later
- **Type inference** - Go can infer types from arguments
- **Constraints** - Use to limit acceptable types
- **Performance** - No runtime overhead (monomorphization)

## Next Steps

After understanding generics, proceed to:
- [02-Reflection](../02-Reflection/README.md) - Learn about reflection

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

