# Reflection

Learn about reflection in Go, which allows runtime type inspection and manipulation.

## Learning Objectives

- Understand what reflection is and when to use it
- Learn to inspect types at runtime
- Master value reflection and manipulation
- Understand struct field iteration
- Learn to call methods dynamically
- Understand reflection limitations and performance

## Concepts Covered

### Type Reflection

Get type information at runtime:

```go
var x int = 42
t := reflect.TypeOf(x)
fmt.Println(t.Kind())  // int
```

### Value Reflection

Get and manipulate values:

```go
v := reflect.ValueOf(x)
fmt.Println(v.Int())  // 42
```

### Setting Values

Need a pointer to set values:

```go
var y int = 10
vp := reflect.ValueOf(&y)
elem := vp.Elem()
elem.SetInt(20)  // y is now 20
```

### Struct Reflection

Inspect struct fields:

```go
structType := reflect.TypeOf(p)
for i := 0; i < structType.NumField(); i++ {
    field := structType.Field(i)
    fmt.Println(field.Name, field.Type)
}
```

### Method Reflection

Inspect and call methods:

```go
personType := reflect.TypeOf(p)
method := personType.Method(0)
fmt.Println(method.Name)

// Call method
personValue := reflect.ValueOf(p)
methodValue := personValue.MethodByName("GetName")
results := methodValue.Call(nil)
```

### Struct Tags

Access struct tags:

```go
field := structType.Field(0)
jsonTag := field.Tag.Get("json")
validateTag := field.Tag.Get("validate")
```

## Running the Example

```bash
# Navigate to this directory
cd 02-Reflection

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/reflection

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

1. **Runtime type inspection** - Examine types at runtime
2. **Value manipulation** - Get and set values dynamically
3. **Struct iteration** - Iterate over struct fields
4. **Method calling** - Call methods dynamically
5. **Performance cost** - Reflection has overhead
6. **Use sparingly** - Prefer type assertions when possible

## Common Patterns

**Type checking:**
```go
t := reflect.TypeOf(value)
if t.Kind() == reflect.String {
    // handle string
}
```

**Struct field access:**
```go
structValue := reflect.ValueOf(s)
field := structValue.FieldByName("Name")
if field.IsValid() && field.CanSet() {
    field.SetString("New Name")
}
```

**Create new instance:**
```go
t := reflect.TypeOf(MyStruct{})
newValue := reflect.New(t).Elem()
```

## Best Practices

- **Use when necessary** - Reflection has performance overhead
- **Prefer type assertions** - Use type assertions when possible
- **Check validity** - Always check IsValid() and CanSet()
- **Handle errors** - Reflection can panic, handle carefully
- **Document usage** - Reflection code can be hard to understand

## Important Notes

- **Performance overhead** - Reflection is slower than direct access
- **Type safety** - Reflection bypasses compile-time type checking
- **Pointer requirement** - Need pointer to set values
- **Unexported fields** - Cannot access unexported fields directly
- **Use cases** - JSON encoding/decoding, ORMs, validation libraries

## Common Use Cases

1. **JSON encoding/decoding** - `encoding/json` uses reflection
2. **ORM libraries** - Database mapping uses reflection
3. **Validation** - Struct tag validation
4. **Serialization** - Converting to/from different formats
5. **Dependency injection** - Framework implementations

## Next Steps

After understanding reflection, proceed to:
- [03-Profiling-Tracing](../03-Profiling-Tracing/README.md) - Learn about profiling

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

