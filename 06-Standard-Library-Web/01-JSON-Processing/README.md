# JSON Processing

Learn about JSON encoding and decoding in Go using the `encoding/json` package.

## Learning Objectives

- Understand how to marshal Go structs to JSON
- Learn how to unmarshal JSON to Go structs
- Master JSON struct tags
- Understand custom marshaling and unmarshaling
- Learn about JSON with nested structs
- Understand error handling with JSON

## Concepts Covered

### Marshal (Go to JSON)

Convert Go values to JSON:

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

person := Person{Name: "Alice", Age: 30}
jsonData, err := json.Marshal(person)
```

### Unmarshal (JSON to Go)

Convert JSON to Go values:

```go
jsonStr := `{"name":"Bob","age":25}`
var person Person
err := json.Unmarshal([]byte(jsonStr), &person)
```

### JSON Tags

Control JSON field names and behavior:

```go
type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Email   string `json:"email,omitempty"` // Omit if empty
    Address string `json:"-"`                // Ignore field
}
```

### MarshalIndent

Pretty-print JSON:

```go
jsonData, err := json.MarshalIndent(person, "", "  ")
```

### Custom Marshaling

Implement `MarshalJSON` and `UnmarshalJSON`:

```go
func (d CustomDate) MarshalJSON() ([]byte, error) {
    return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d *CustomDate) UnmarshalJSON(data []byte) error {
    // Custom unmarshaling logic
}
```

## Running the Example

```bash
# Navigate to this directory
cd 01-JSON-Processing

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/json-processing

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

1. **Marshal** - Go struct to JSON
2. **Unmarshal** - JSON to Go struct
3. **JSON tags** - Control field names and behavior
4. **omitempty** - Omit empty fields
5. **Custom marshaling** - Implement MarshalJSON/UnmarshalJSON
6. **Error handling** - Always check errors

## Common Patterns

**Basic marshal/unmarshal:**
```go
jsonData, err := json.Marshal(value)
var result Type
err = json.Unmarshal(jsonData, &result)
```

**Pretty print:**
```go
jsonData, err := json.MarshalIndent(value, "", "  ")
```

**Unmarshal to map:**
```go
var data map[string]interface{}
json.Unmarshal(jsonData, &data)
```

## Best Practices

- **Always check errors** - JSON operations can fail
- **Use struct tags** - For proper field mapping
- **Use omitempty** - For optional fields
- **Custom marshaling** - When default behavior isn't enough
- **Validate data** - After unmarshaling

## JSON Tag Options

- **Field name**: `json:"field_name"`
- **Omit empty**: `json:"field,omitempty"`
- **Ignore**: `json:"-"`
- **String numbers**: `json:"field,string"`

## Next Steps

After understanding JSON processing, proceed to:
- [02-HTTP-Server](../02-HTTP-Server/README.md) - Learn about HTTP servers

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

