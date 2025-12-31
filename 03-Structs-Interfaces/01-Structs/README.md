# Structs

Learn about structs in Go: how to define, create, and work with custom data types.

## Learning Objectives

- Understand what structs are and when to use them
- Learn how to define and create structs
- Master struct literals and field access
- Understand nested structs
- Learn about struct comparison and zero values

## Concepts Covered

### Struct Definition

A struct is a collection of fields. It's Go's way of creating custom types:

```go
type Person struct {
    Name string
    Age  int
}
```

### Creating Struct Instances

**Zero value:**
```go
var p Person  // All fields are zero values
```

**Struct literal (named fields):**
```go
p := Person{Name: "Alice", Age: 30}
```

**Struct literal (positional):**
```go
p := Person{"Alice", 30}  // Order must match struct definition
```

**Using new():**
```go
p := new(Person)  // Returns pointer
p.Name = "Alice"
```

### Accessing Fields

```go
p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name)  // "Alice"
p.Age = 31           // Modify field
```

### Nested Structs

```go
type Address struct {
    Street string
    City   string
}

type Contact struct {
    Email   string
    Address Address  // Nested struct
}

contact := Contact{
    Email: "test@example.com",
    Address: Address{
        Street: "123 Main St",
        City:   "New York",
    },
}
fmt.Println(contact.Address.City)  // "New York"
```

### Anonymous Structs

Structs without a type name:

```go
anon := struct {
    ID   int
    Name string
}{
    ID:   1,
    Name: "Anonymous",
}
```

### Struct Comparison

Structs are comparable if all their fields are comparable:

```go
p1 := Person{Name: "Alice", Age: 30}
p2 := Person{Name: "Alice", Age: 30}
p1 == p2  // true
```

## Running the Example

```bash
# Navigate to this directory
cd 01-Structs

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/03-structs-interfaces/structs

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

1. **Structs group related data** - Like classes in OOP, but simpler
2. **Fields are accessed with dot notation** - `struct.Field`
3. **Structs are value types** - Copied when assigned (unless using pointers)
4. **Zero values** - All fields initialized to their zero values
5. **Comparable** - If all fields are comparable, structs can be compared
6. **Nested structs** - Structs can contain other structs

## Common Patterns

**Constructor function:**
```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}
```

**Struct with tags:**
```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`  // Excluded from JSON
}
```

**Multiple fields of same type:**
```go
type Point struct {
    X, Y int  // Shorthand for X int, Y int
}
```

## Best Practices

- **Use descriptive field names** - Clear, self-documenting code
- **Group related fields** - Logical organization
- **Use struct tags** - For serialization (JSON, XML, etc.)
- **Prefer named fields** - More readable than positional
- **Use pointers for large structs** - Avoid copying overhead

## Struct Tags

Struct tags are metadata attached to struct fields:

```go
type User struct {
    ID   int    `json:"id" db:"user_id"`
    Name string `json:"name" validate:"required"`
}
```

Common uses:
- JSON serialization
- Database mapping
- Validation
- Documentation

## Zero Values

Each field type has a zero value:
- `int`: `0`
- `float64`: `0.0`
- `string`: `""`
- `bool`: `false`
- Pointers, slices, maps: `nil`

## Next Steps

After understanding structs, proceed to:
- [02-Methods](../02-Methods/README.md) - Learn about methods (functions on structs)

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

