# Methods

Learn about methods in Go: how to attach functions to types using receivers.

## Learning Objectives

- Understand what methods are and how they differ from functions
- Learn about value receivers vs pointer receivers
- Understand when to use each type of receiver
- Learn about method sets and how Go handles them
- Master method chaining

## Concepts Covered

### What are Methods?

Methods are functions with a special receiver argument. They allow you to define behavior on types.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) ScaleInPlace(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### Value Receiver

Value receiver methods work on a **copy** of the value:

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()  // Works on copy
```

**Characteristics:**
- Doesn't modify the original value
- Can be called on value or pointer (Go auto-converts)
- Good for small types or when immutability is desired

### Pointer Receiver

Pointer receiver methods work on the **original** value:

```go
func (r *Rectangle) ScaleInPlace(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

rect := &Rectangle{Width: 10, Height: 5}
rect.ScaleInPlace(2.0)  // Modifies original
```

**Characteristics:**
- Modifies the original value
- Can be called on value or pointer (Go auto-converts)
- Good for large types or when modification is needed

### Method Call Syntax

Go automatically handles conversion between values and pointers:

```go
rect := Rectangle{Width: 10, Height: 5}
rect.Area()        // Works on value
(&rect).Area()     // Also works (Go auto-dereferences)

ptr := &Rectangle{Width: 10, Height: 5}
ptr.Area()         // Works on pointer (Go auto-dereferences)
(*ptr).Area()      // Explicit dereference
```

### Methods on Type Aliases

You can define methods on any type, including type aliases:

```go
type MyInt int

func (m MyInt) IsEven() bool {
    return m%2 == 0
}

num := MyInt(4)
num.IsEven()  // true
```

## Running the Example

```bash
# Navigate to this directory
cd 02-Methods

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/03-structs-interfaces/methods

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

1. **Methods attach behavior to types** - Like methods in OOP
2. **Value receiver** - Works on copy, doesn't modify original
3. **Pointer receiver** - Works on original, can modify
4. **Go auto-converts** - Can call value receiver on pointer and vice versa
5. **Method sets** - Value receiver methods are in both value and pointer method sets
6. **Pointer receiver methods** - Only in pointer method set (important for interfaces)

## When to Use Value vs Pointer Receiver

### Use Value Receiver When:
- Method doesn't modify the receiver
- Receiver is small (no copying overhead)
- You want immutability
- Consistency with other methods

### Use Pointer Receiver When:
- Method modifies the receiver
- Receiver is large (avoid copying overhead)
- Consistency (if one method uses pointer, all should)
- Method needs to handle nil receiver

## Common Patterns

**Getter methods:**
```go
func (p Person) GetName() string {
    return p.Name
}
```

**Setter methods:**
```go
func (p *Person) SetName(name string) {
    p.Name = name
}
```

**Method chaining:**
```go
func (r *Rectangle) SetDimensions(w, h float64) *Rectangle {
    r.Width = w
    r.Height = h
    return r  // Return receiver for chaining
}

rect.SetDimensions(10, 5).ScaleInPlace(2)
```

## Best Practices

- **Be consistent** - If one method uses pointer receiver, all should
- **Use pointer receivers** for methods that modify
- **Use value receivers** for methods that don't modify
- **Consider size** - Large structs should use pointer receivers
- **Document behavior** - Make it clear if method modifies receiver

## Method Sets

The method set of a type determines which methods can be called:

- **Value method set**: Methods with value receivers
- **Pointer method set**: Methods with pointer receivers

**Important for interfaces:**
- Value type implements interface if it has all value receiver methods
- Pointer type implements interface if it has all methods (value or pointer)

## Next Steps

After understanding methods, proceed to:
- [03-Interfaces](../03-Interfaces/README.md) - Learn about interfaces and polymorphism

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

