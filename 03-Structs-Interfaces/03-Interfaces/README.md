# Interfaces

Learn about interfaces in Go: duck typing, polymorphism, and how interfaces enable flexible, decoupled code.

## Learning Objectives

- Understand what interfaces are and how they work in Go
- Learn about implicit interface implementation (duck typing)
- Master polymorphism using interfaces
- Understand empty interfaces (interface{})
- Learn about interface composition

## Concepts Covered

### What are Interfaces?

An interface is a collection of method signatures. A type implements an interface if it has all the methods the interface requires.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

### Implicit Implementation

In Go, interfaces are implemented **implicitly**. If a type has all the methods of an interface, it implements that interface automatically.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Rectangle automatically implements Shape interface!
```

### Duck Typing

Go follows the "duck typing" philosophy: "If it walks like a duck and quacks like a duck, it's a duck."

If a type has all the methods an interface requires, it implements that interface - no explicit declaration needed.

### Polymorphism

Interfaces enable polymorphism - different types can be used interchangeably through a common interface:

```go
func printArea(s Shape) {
    fmt.Println(s.Area())
}

rect := Rectangle{Width: 10, Height: 5}
circle := Circle{Radius: 5}

printArea(rect)   // Works
printArea(circle) // Also works
```

### Empty Interface

The empty interface `interface{}` can hold any value:

```go
var v interface{}
v = 42
v = "hello"
v = 3.14
```

**Note:** In Go 1.18+, `any` is an alias for `interface{}`.

### Interface Composition

Interfaces can be composed from other interfaces:

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

## Running the Example

```bash
# Navigate to this directory
cd 03-Interfaces

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/03-structs-interfaces/interfaces

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

1. **Interfaces define behavior** - What methods a type must have
2. **Implicit implementation** - No "implements" keyword needed
3. **Duck typing** - If it has the methods, it implements the interface
4. **Polymorphism** - Different types can be used through common interface
5. **Empty interface** - Can hold any value (`interface{}` or `any`)
6. **Interface composition** - Build complex interfaces from simple ones

## Common Patterns

**Accept interfaces, return structs:**
```go
func ProcessData(r Reader) error {
    // Work with Reader interface
}
```

**Interface as parameter:**
```go
func PrintArea(s Shape) {
    fmt.Println(s.Area())
}
```

**Type assertion (covered in next topic):**
```go
if str, ok := v.(string); ok {
    // v is a string
}
```

## Best Practices

- **Keep interfaces small** - Prefer many small interfaces
- **Accept interfaces** - Functions should accept interfaces, not concrete types
- **Return structs** - Functions should return concrete types
- **Compose interfaces** - Build complex interfaces from simple ones
- **Use meaningful names** - Interface names should describe behavior

## Interface vs Struct

- **Struct**: Defines data (what it is)
- **Interface**: Defines behavior (what it can do)

## Important Notes

- **Nil interface** - An interface variable can be nil
- **Nil value in interface** - An interface can hold a nil value
- **Method sets matter** - Pointer receivers affect interface implementation
- **Empty interface is powerful** - But use sparingly (loses type safety)

## Next Steps

After understanding interfaces, proceed to:
- [04-Type-Assertions-Switches](../04-Type-Assertions-Switches/README.md) - Learn about type assertions and type switches

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

