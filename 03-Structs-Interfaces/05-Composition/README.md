# Composition

Learn about composition and embedding in Go: how to build complex types from simpler ones without inheritance.

## Learning Objectives

- Understand composition vs embedding
- Learn how to embed structs and interfaces
- Master method promotion
- Understand when to use composition vs embedding
- Learn about method overriding

## Concepts Covered

### Composition (has-a)

Composition is when a struct contains another struct as a field:

```go
type Engine struct {
    Horsepower int
}

type Car struct {
    Make   string
    Engine Engine  // Composition: Car has an Engine
}

car := Car{
    Make: "Toyota",
    Engine: Engine{Horsepower: 200},
}
fmt.Println(car.Engine.Horsepower)  // Access via field name
```

### Embedding (is-a)

Embedding is when a struct embeds another struct without a field name:

```go
type Animal struct {
    Name string
    Age  int
}

type Dog struct {
    Animal  // Embedding: Dog is an Animal
    Breed   string
}

dog := Dog{
    Animal: Animal{Name: "Buddy", Age: 3},
    Breed:  "Golden Retriever",
}
fmt.Println(dog.Name)  // Direct access (promoted)
dog.Speak()            // Methods are promoted too
```

### Method Promotion

When you embed a type, its methods are promoted to the outer type:

```go
func (a Animal) Speak() {
    fmt.Println("Animal sound")
}

dog := Dog{Animal: Animal{Name: "Buddy"}}
dog.Speak()  // Can call Animal's method directly
```

### Method Overriding

You can override embedded methods:

```go
func (a Animal) Speak() {
    fmt.Println("Animal sound")
}

func (d Dog) Speak() {
    fmt.Println("Woof!")  // Overrides Animal's Speak()
}
```

### Multiple Embedding

You can embed multiple types:

```go
type Writer struct {
    Name string
}

type Reader struct {
    Name string
}

type ReadWriter struct {
    Writer  // Embed Writer
    Reader  // Embed Reader
}

rw := ReadWriter{
    Writer: Writer{Name: "Alice"},
    Reader: Reader{Name: "Bob"},
}
rw.Write()  // From Writer
rw.Read()   // From Reader
```

### Embedding Interfaces

You can embed interfaces:

```go
type WriterInterface interface {
    Write()
}

type Logger struct {
    WriterInterface  // Embed interface
    Level           string
}
```

### Embedding Pointers

You can embed pointers:

```go
type Person struct {
    Name string
}

type Employee struct {
    *Person  // Embed pointer
    ID      int
}
```

## Running the Example

```bash
# Navigate to this directory
cd 05-Composition

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/03-structs-interfaces/composition

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

1. **Composition (has-a)** - Struct contains another struct as field
2. **Embedding (is-a)** - Struct embeds another struct, inherits methods
3. **Method promotion** - Embedded methods become part of outer type
4. **Method overriding** - Can override embedded methods
5. **Multiple embedding** - Can embed multiple types
6. **No inheritance** - Go uses composition, not inheritance

## Composition vs Embedding

### Use Composition When:
- You need explicit relationship
- You want to control access
- Relationship is "has-a" (Car has Engine)
- You need multiple instances

### Use Embedding When:
- You want method promotion
- Relationship is "is-a" (Dog is Animal)
- You want to extend behavior
- You want to implement interfaces

## Common Patterns

**Extending functionality:**
```go
type Base struct {
    Name string
}

type Extended struct {
    Base
    Extra string
}
```

**Interface implementation:**
```go
type Reader interface {
    Read()
}

type FileReader struct {
    Reader  // Embed interface
    Filename string
}
```

**Multiple capabilities:**
```go
type ReadWriter struct {
    Reader
    Writer
}
```

## Best Practices

- **Prefer composition over inheritance** - Go's philosophy
- **Use embedding for code reuse** - Promotes methods automatically
- **Avoid deep hierarchies** - Keep it simple
- **Be explicit when needed** - Use composition if clarity is important
- **Document relationships** - Make it clear what you're doing

## Important Notes

- **Ambiguous methods** - If two embedded types have same method, must specify
- **Field promotion** - Embedded fields are accessible directly
- **Method sets** - Embedded methods affect interface implementation
- **Pointer embedding** - Can embed pointers for shared state

## Go's Philosophy

Go doesn't have inheritance. Instead, it uses:
- **Composition** - Building complex types from simple ones
- **Embedding** - Promoting methods and fields
- **Interfaces** - Defining behavior contracts

This leads to:
- More flexible code
- Easier to understand
- Less coupling
- Better testability

## Next Steps

Congratulations! You've completed Section 3: Structs & Interfaces.

Proceed to:
- [Section 4: Error Handling & Testing](../../04-Error-Handling-Testing/) - Learn about error handling and testing

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

