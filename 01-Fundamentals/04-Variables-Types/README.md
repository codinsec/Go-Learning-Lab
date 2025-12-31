# Variables & Types

Learn about variable declaration, constants, type inference, and Go's type system.

## Learning Objectives

- Understand different ways to declare variables
- Learn about constants and their usage
- Understand type inference in Go
- Learn about zero values
- Understand type conversion

## Concepts Covered

### Variable Declaration Methods

**1. Full Declaration with Type**
```go
var name string = "Go"
```

**2. Type Inference**
```go
var name = "Go"  // Go infers the type as string
```

**3. Short Variable Declaration (Most Common)**
```go
name := "Go"  // Only inside functions
```

### Constants

Constants are declared with `const` keyword and cannot be changed:
```go
const pi = 3.14159
const greeting = "Hello, Go!"
```

### Zero Values

In Go, variables are initialized to their zero value if not explicitly initialized:
- `int`: `0`
- `float64`: `0.0`
- `string`: `""` (empty string)
- `bool`: `false`
- Pointers, slices, maps, channels: `nil`

### Type Conversion

Go requires explicit type conversion:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(i)
```

## Running the Example

```bash
# Navigate to this directory
cd 04-Variables-Types

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/01-fundamentals/variables-types

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

```
=== Variables & Types ===

Language: Go, Version: 1
Language: Golang, Year: 2009
Author: Google, Open Source: true
x = 10, y = 20
a = 5, b = 15
Pi: 3.141590
Greeting: Hello, Go!
Max Int: 100, Min Float: 0.000000
Status OK: 200, Status Error: 500

Zero values:
  int: 0
  float64: 0.000000
  string: ""
  bool: false

Type conversions:
  int to float64: 42 -> 42.000000
  int to uint: 42 -> 42
```

## Key Takeaways

1. **Short declaration (`:=`)** is the most common way to declare variables
2. **Type inference** makes code cleaner - Go figures out the type
3. **Zero values** ensure variables always have a value
4. **Explicit type conversion** is required (no implicit conversion)
5. **Constants** are immutable and can be untyped or typed

## Common Types in Go

- **Integers**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Floats**: `float32`, `float64`
- **Complex**: `complex64`, `complex128`
- **Boolean**: `bool`
- **String**: `string`
- **Byte**: `byte` (alias for `uint8`)
- **Rune**: `rune` (alias for `int32`, represents a Unicode code point)

## Best Practices

- Use short declaration (`:=`) inside functions
- Use `var` for package-level variables
- Prefer type inference when the type is obvious
- Use constants for values that don't change
- Be explicit with type conversions

## Next Steps

After understanding variables and types, proceed to:
- [05-Control-Structures](../05-Control-Structures/README.md) - Learn about control flow

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

