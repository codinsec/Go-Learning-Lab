# Type Assertions & Type Switches

Learn about type assertions and type switches for runtime type checking and conversion in Go.

## Learning Objectives

- Understand what type assertions are and how to use them
- Learn the safe and unsafe forms of type assertions
- Master type switches for multiple type checks
- Understand when to use type assertions vs type switches
- Learn practical patterns for working with interfaces

## Concepts Covered

### Type Assertion

A type assertion provides access to an interface value's underlying concrete type.

**Safe form (with ok check):**
```go
var i interface{} = "hello"
str, ok := i.(string)
if ok {
    // str is a string
} else {
    // i is not a string
}
```

**Unsafe form (can panic):**
```go
var i interface{} = "hello"
str := i.(string)  // Safe if you're sure
// str := i.(int)   // Would panic!
```

### Type Assertion with Interfaces

You can assert an interface to a concrete type:

```go
var shape Shape = Rectangle{Width: 10, Height: 5}
rect, ok := shape.(Rectangle)
if ok {
    fmt.Println(rect.Width)  // Access Rectangle-specific fields
}
```

### Type Switch

A type switch is a construct that permits several type assertions in series:

```go
switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

### Multiple Cases

You can match multiple types in a single case:

```go
switch v := i.(type) {
case int, int32, int64:
    fmt.Println("Integer type")
case float32, float64:
    fmt.Println("Float type")
}
```

### Default Case

Use default for unmatched types:

```go
switch v := i.(type) {
case string:
    fmt.Println("String")
default:
    fmt.Printf("Other: %T\n", v)
}
```

## Running the Example

```bash
# Navigate to this directory
cd 04-Type-Assertions-Switches

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/03-structs-interfaces/type-assertions-switches

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

1. **Type assertion** - Access underlying concrete type from interface
2. **Safe form** - Always use `value, ok := i.(Type)` to avoid panics
3. **Type switch** - Efficient way to check multiple types
4. **Multiple cases** - Can match several types in one case
5. **Default case** - Handles unmatched types
6. **Nil handling** - Always check for nil before type assertion

## Common Patterns

**Safe type assertion:**
```go
if str, ok := v.(string); ok {
    // Use str
}
```

**Type switch with interface:**
```go
switch shape := s.(type) {
case Rectangle:
    // Handle Rectangle
case Circle:
    // Handle Circle
}
```

**Checking for nil:**
```go
if shape == nil {
    return errors.New("shape is nil")
}
rect, ok := shape.(Rectangle)
```

## Best Practices

- **Always use safe form** - Check `ok` to avoid panics
- **Use type switch for multiple types** - More efficient than multiple if-else
- **Handle nil cases** - Check for nil before type assertion
- **Use default case** - Catch unexpected types
- **Be specific** - Use concrete types, not interfaces, when possible

## When to Use

**Type Assertion:**
- When you need to check one specific type
- When you need to access type-specific fields/methods
- When working with empty interface (`interface{}`)

**Type Switch:**
- When checking multiple possible types
- When you need different handling for each type
- When processing unknown data (JSON, etc.)

## Important Notes

- **Type assertion can panic** - Always use safe form with `ok`
- **Nil interface** - Type assertion on nil interface returns false
- **Pointer types** - Use `*Type` for pointer assertions
- **Type switch variable** - Variable in switch is of the asserted type

## Practical Examples

**JSON processing:**
```go
switch v := data.(type) {
case map[string]interface{}:
    // Handle object
case []interface{}:
    // Handle array
case string:
    // Handle string
}
```

**Error handling:**
```go
if err, ok := err.(CustomError); ok {
    // Handle custom error
}
```

## Next Steps

After understanding type assertions and type switches, proceed to:
- [05-Composition](../05-Composition/README.md) - Learn about composition and embedding

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

