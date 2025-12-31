# Maps

Learn about maps (key-value pairs) in Go, how to create, manipulate, and iterate over them.

## Learning Objectives

- Understand what maps are and when to use them
- Learn how to create and initialize maps
- Master map operations: add, access, delete, check existence
- Learn how to iterate over maps
- Understand map zero values and nil maps

## Concepts Covered

### Map Declaration

Maps are key-value pairs, similar to dictionaries in other languages.

```go
var m map[string]int              // nil map
m := make(map[string]int)          // Empty map
m := map[string]int{               // Map literal
    "key1": 10,
    "key2": 20,
}
```

### Map Operations

**Adding/Updating:**
```go
m["key"] = value
```

**Accessing:**
```go
value := m["key"]
```

**Checking existence:**
```go
value, exists := m["key"]
if exists {
    // key exists
}
```

**Deleting:**
```go
delete(m, "key")
```

**Length:**
```go
length := len(m)
```

### Map Characteristics

- **Reference type** - maps are passed by reference
- **Zero value is nil** - nil maps cannot be written to
- **Keys must be comparable** - can't use slices, maps, or functions as keys
- **Values can be any type** - including other maps, slices, etc.
- **Iteration order is random** - don't rely on order

## Running the Example

```bash
# Navigate to this directory
cd 02-Maps

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/02-data-structures-functions/maps

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

1. **Maps are reference types** - changes affect all references
2. **Nil maps cannot be written to** - use `make()` or literal to initialize
3. **Zero value for missing keys** - returns zero value of value type
4. **Always check existence** - use two-value assignment: `value, ok := m[key]`
5. **Iteration order is random** - use a slice if you need order
6. **Keys must be comparable** - strings, numbers, arrays, structs (with comparable fields)

## Common Patterns

**Check if key exists:**
```go
if value, ok := m[key]; ok {
    // key exists, use value
}
```

**Initialize if nil:**
```go
if m == nil {
    m = make(map[string]int)
}
```

**Map as set:**
```go
set := make(map[string]bool)
set["item"] = true
if set["item"] {
    // item is in set
}
```

**Nested maps:**
```go
nested := make(map[string]map[string]int)
nested["outer"] = make(map[string]int)
nested["outer"]["inner"] = 42
```

**Clear a map:**
```go
// Method 1: Iterate and delete
for key := range m {
    delete(m, key)
}

// Method 2: Reassign
m = make(map[string]int)
```

## Best Practices

- **Always check existence** when you're not sure a key exists
- **Use make()** to create maps when you need to write to them
- **Use map literals** for initialization with known values
- **Don't rely on iteration order** - use slices if order matters
- **Use descriptive key types** - prefer meaningful types over generic ones

## Use Cases

- **Configuration**: `map[string]string`
- **Counters**: `map[string]int`
- **Sets**: `map[T]bool`
- **Caching**: `map[string]interface{}`
- **Grouping**: `map[string][]T`

## Next Steps

After understanding maps, proceed to:
- [03-Functions](../03-Functions/README.md) - Learn about functions, multiple returns, and named returns

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

