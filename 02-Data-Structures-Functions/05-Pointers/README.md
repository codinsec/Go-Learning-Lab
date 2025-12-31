# Pointers

Learn about pointers in Go: memory addresses, dereferencing, and when to use pointers.

## Learning Objectives

- Understand what pointers are and how they work
- Learn pointer operators: `&` (address) and `*` (dereference)
- Understand when to use pointers vs values
- Learn about pointer receivers in methods
- Understand nil pointers and how to handle them

## Concepts Covered

### What are Pointers?

A pointer holds the memory address of a value. Pointers allow you to:
- Pass large values efficiently (without copying)
- Modify values in functions
- Represent optional values (nil)

### Pointer Operators

**`&` (address operator):**
```go
x := 42
p := &x  // p is a pointer to x (holds address of x)
```

**`*` (dereference operator):**
```go
*p = 100  // Dereference p to access/modify the value
value := *p  // Get the value p points to
```

### Creating Pointers

**Using address operator:**
```go
x := 42
p := &x
```

**Using new():**
```go
p := new(int)  // Allocates memory, returns pointer to zero value
*p = 42
```

### Zero Value

The zero value of a pointer is `nil`:
```go
var p *int  // p is nil
// *p would panic - don't dereference nil pointer
```

### Passing by Value vs Reference

**By value (copy):**
```go
func swapByValue(x, y int) {
    x, y = y, x  // Only swaps local copies
}
```

**By reference (pointer):**
```go
func swapByReference(x, y *int) {
    *x, *y = *y, *x  // Swaps actual values
}
```

### Pointer Receivers

Methods can have pointer receivers to modify the original value:

```go
type Counter struct {
    value int
}

// Value receiver (works on copy)
func (c Counter) Increment() {
    c.value++  // Doesn't modify original
}

// Pointer receiver (works on original)
func (c *Counter) IncrementByPointer() {
    c.value++  // Modifies original
}
```

## Running the Example

```bash
# Navigate to this directory
cd 05-Pointers

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/02-data-structures-functions/pointers

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

1. **`&` gets address**, `*` dereferences pointer
2. **Zero value is nil** - always check before dereferencing
3. **Pointers allow modification** - pass by reference when needed
4. **Pointer receivers modify** - value receivers work on copies
5. **Go auto-dereferences** - `ptr.field` works, no need for `(*ptr).field` in most cases
6. **Use pointers for large structs** - avoid copying overhead

## Common Patterns

**Optional values:**
```go
func findUser(id int) (*User, error) {
    // Returns nil if not found
    if notFound {
        return nil, errors.New("user not found")
    }
    return &user, nil
}
```

**Modifying in function:**
```go
func increment(x *int) {
    *x++
}
```

**Pointer to struct:**
```go
person := &Person{Name: "Alice"}
// Go auto-dereferences:
person.Name = "Bob"  // Same as (*person).Name
```

## Best Practices

- **Use pointers when you need to modify** - Pass by reference
- **Use pointers for large structs** - Avoid copying overhead
- **Check for nil** - Always verify pointer is not nil before dereferencing
- **Use pointer receivers** - When methods need to modify the receiver
- **Prefer value receivers** - When methods don't need to modify (simpler)

## When to Use Pointers

✅ **Use pointers when:**
- You need to modify a value in a function
- Passing large structs (avoid copying)
- Value can be nil (optional values)
- Method needs to modify receiver

❌ **Don't use pointers when:**
- Value is small (int, bool, etc.)
- You don't need to modify
- Method doesn't modify receiver
- Unnecessary complexity

## Important Notes

- **Nil pointer dereference causes panic** - Always check for nil
- **Go auto-dereferences struct fields** - `ptr.field` works
- **Pointers are passed by value** - The pointer itself is copied, not the value
- **Returning pointer to local variable is safe** - Go's escape analysis handles it

## Next Steps

Congratulations! You've completed Section 2: Data Structures & Functions.

Proceed to:
- [Section 3: Structs & Interfaces](../../03-Structs-Interfaces/) - Learn about structs, methods, and interfaces

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

