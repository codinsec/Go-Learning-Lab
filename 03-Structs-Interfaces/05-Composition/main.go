package main

import "fmt"

// This program demonstrates composition and embedding in Go

// 1. Basic composition (has-a relationship)
type Engine struct {
	Horsepower int
	Type       string
}

type Car struct {
	Make   string
	Model  string
	Engine Engine // Composition: Car has an Engine
}

// 2. Embedding (is-a relationship)
type Animal struct {
	Name string
	Age  int
}

func (a Animal) Speak() {
	fmt.Printf("%s makes a sound\n", a.Name)
}

type Dog struct {
	Animal // Embedding: Dog is an Animal (inherits fields and methods)
	Breed  string
}

// 3. Method overriding
func (d Dog) Speak() {
	fmt.Printf("%s barks: Woof!\n", d.Name)
}

// 4. Multiple embedding
type Writer struct {
	Name string
}

func (w Writer) Write() {
	fmt.Printf("%s is writing\n", w.Name)
}

type Reader struct {
	Name string
}

func (r Reader) Read() {
	fmt.Printf("%s is reading\n", r.Name)
}

type ReadWriter struct {
	Writer // Embed Writer
	Reader // Embed Reader
}

// 5. Embedding with same method names
type A struct {
	Name string
}

func (a A) Do() {
	fmt.Println("A.Do()")
}

type B struct {
	Name string
}

func (b B) Do() {
	fmt.Println("B.Do()")
}

type C struct {
	A // Embed A
	B // Embed B
}

// 6. Embedding interfaces
type WriterInterface interface {
	Write()
}

type FileWriter struct {
	Filename string
}

func (f FileWriter) Write() {
	fmt.Printf("Writing to %s\n", f.Filename)
}

type Logger struct {
	WriterInterface // Embed interface
	Level          string
}

// 7. Embedding pointer
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	*Person // Embed pointer
	ID      int
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("%s, %d years old", p.Name, p.Age)
}

func main() {
	fmt.Println("=== Composition & Embedding ===")
	fmt.Println()

	// 1. Basic composition
	fmt.Println("1. Basic Composition (has-a):")
	car := Car{
		Make:  "Toyota",
		Model: "Camry",
		Engine: Engine{
			Horsepower: 200,
			Type:       "V6",
		},
	}
	fmt.Printf("   Car: %s %s with %s engine (%d HP)\n", 
		car.Make, car.Model, car.Engine.Type, car.Engine.Horsepower)
	fmt.Println()

	// 2. Embedding
	fmt.Println("2. Embedding (is-a):")
	dog := Dog{
		Animal: Animal{
			Name: "Buddy",
			Age:  3,
		},
		Breed: "Golden Retriever",
	}
	fmt.Printf("   Dog: %s, %d years old, Breed: %s\n", 
		dog.Name, dog.Age, dog.Breed)
	dog.Speak() // Can call Animal method directly
	fmt.Println()

	// 3. Method overriding
	fmt.Println("3. Method Overriding:")
	animal := Animal{Name: "Generic Animal"}
	animal.Speak()
	dog.Speak() // Dog's Speak() overrides Animal's Speak()
	fmt.Println()

	// 4. Multiple embedding
	fmt.Println("4. Multiple Embedding:")
	rw := ReadWriter{
		Writer: Writer{Name: "Alice"},
		Reader: Reader{Name: "Bob"},
	}
	rw.Write() // Method from Writer
	rw.Read()   // Method from Reader
	fmt.Println()

	// 5. Embedding with same method names
	fmt.Println("5. Embedding with Same Method Names:")
	c := C{
		A: A{Name: "A"},
		B: B{Name: "B"},
	}
	c.A.Do() // Must specify which embedded type
	c.B.Do()
	// c.Do() would be ambiguous - compile error!
	fmt.Println()

	// 6. Embedding interfaces
	fmt.Println("6. Embedding Interfaces:")
	logger := Logger{
		WriterInterface: FileWriter{Filename: "app.log"},
		Level:          "INFO",
	}
	logger.Write() // Calls embedded interface method
	fmt.Printf("   Logger level: %s\n", logger.Level)
	fmt.Println()

	// 7. Embedding pointer
	fmt.Println("7. Embedding Pointer:")
	person := &Person{Name: "Charlie", Age: 30}
	emp := Employee{
		Person: person,
		ID:     1001,
	}
	fmt.Printf("   Employee: %s, ID: %d\n", emp.GetInfo(), emp.ID)
	fmt.Printf("   Direct access: %s, %d years old\n", emp.Name, emp.Age)
	fmt.Println()

	// 8. Composition vs Embedding
	fmt.Println("8. Composition vs Embedding:")
	fmt.Println("   Composition (has-a):")
	fmt.Println("     - Car has Engine (explicit: car.Engine.Horsepower)")
	fmt.Println("     - Explicit relationship")
	fmt.Println("   Embedding (is-a):")
	fmt.Println("     - Dog is Animal (implicit: dog.Name)")
	fmt.Println("     - Implicit relationship, inherits methods")
	fmt.Println()

	// 9. Benefits of embedding
	fmt.Println("9. Benefits of Embedding:")
	fmt.Println("   - Code reuse without inheritance")
	fmt.Println("   - Promotes composition over inheritance")
	fmt.Println("   - Methods are promoted to outer type")
	fmt.Println("   - Can override methods if needed")
	fmt.Println()

	// 10. Best practices
	fmt.Println("10. Best Practices:")
	fmt.Println("   - Use embedding for 'is-a' relationships")
	fmt.Println("   - Use composition for 'has-a' relationships")
	fmt.Println("   - Embed interfaces for flexibility")
	fmt.Println("   - Avoid deep embedding hierarchies")
}

