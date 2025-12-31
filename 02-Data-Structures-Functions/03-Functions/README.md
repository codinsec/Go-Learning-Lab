# Functions

Learn about functions in Go: single and multiple return values, named returns, variadic functions, closures, and more.

## Learning Objectives

- Understand function declaration and calling
- Learn about multiple return values
- Master named return values
- Learn about variadic functions
- Understand closures and higher-order functions
- Learn about function values and function types

## Concepts Covered

### Basic Function

```go
func functionName(param1 type1, param2 type2) returnType {
    // function body
    return value
}
```

### Multiple Return Values

Go functions can return multiple values, which is commonly used for error handling:

```go
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Usage
q, r := divide(17, 5)
```

### Named Return Values

You can name return values, which allows "naked return":

```go
func calculate(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return // Returns sum and product automatically
}
```

**Note:** Use named returns sparingly - they can make code less clear.

### Error Handling Pattern

The idiomatic Go pattern is to return an error as the last value:

```go
func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := safeDivide(10, 2)
if err != nil {
    // handle error
}
```

### Variadic Functions

Functions that accept a variable number of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
sum(1, 2, 3)        // 6
sum(1, 2, 3, 4, 5)  // 15
sum()               // 0
```

### Function as Value

Functions are first-class citizens in Go - they can be assigned to variables:

```go
fn := multiply
result := fn(5, 4)
```

### Higher-Order Functions

Functions that take other functions as parameters:

```go
func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

// Usage
result := applyOperation(10, 5, add)
```

### Closures

Functions that capture variables from their enclosing scope:

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Usage
counter := makeCounter()
counter() // 1
counter() // 2
```

### Anonymous Functions

Functions without a name:

```go
func() {
    fmt.Println("Anonymous function")
}()

// Or assign to variable
fn := func(x int) int {
    return x * 2
}
```

## Running the Example

```bash
# Navigate to this directory
cd 03-Functions

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/02-data-structures-functions/functions

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

1. **Multiple returns** are common in Go, especially for error handling
2. **Named returns** allow naked return, but use sparingly
3. **Variadic functions** use `...` syntax
4. **Functions are first-class** - can be assigned, passed, returned
5. **Closures** capture variables from enclosing scope
6. **Error handling** typically uses `(result, error)` pattern

## Common Patterns

**Error handling:**
```go
result, err := someFunction()
if err != nil {
    return err
}
// use result
```

**Variadic with slice:**
```go
numbers := []int{1, 2, 3}
sum(numbers...)  // Spread operator
```

**Function type:**
```go
type Operation func(int, int) int

func apply(a, b int, op Operation) int {
    return op(a, b)
}
```

## Best Practices

- **Return errors as last value** - idiomatic Go pattern
- **Use multiple returns** for functions that can fail
- **Avoid named returns** in most cases - explicit is better
- **Keep functions small** - single responsibility
- **Use descriptive names** - function names should describe what they do

## Next Steps

After understanding functions, proceed to:
- [04-Defer-Panic-Recover](../04-Defer-Panic-Recover/README.md) - Learn about defer, panic, and recover

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

