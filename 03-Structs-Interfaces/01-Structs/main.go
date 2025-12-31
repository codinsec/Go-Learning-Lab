package main

import "fmt"

// This program demonstrates structs in Go

// 1. Basic struct definition
type Person struct {
	Name string
	Age  int
}

// 2. Struct with different field types
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
	Active   bool
}

// 3. Nested struct
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

type Contact struct {
	Email   string
	Phone   string
	Address Address // Nested struct
}

// 4. Anonymous struct fields
type Point struct {
	X, Y int // Multiple fields of same type
}

// 5. Struct with tags (for JSON, etc.)
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // Don't include in JSON
}

func main() {
	fmt.Println("=== Structs ===")
	fmt.Println()

	// 1. Creating struct instances
	fmt.Println("1. Creating Struct Instances:")
	
	// Zero value
	var p1 Person
	fmt.Printf("   Zero value: %+v\n", p1)
	
	// Struct literal
	p2 := Person{Name: "Alice", Age: 30}
	fmt.Printf("   With values: %+v\n", p2)
	
	// Struct literal (positional)
	p3 := Person{"Bob", 25}
	fmt.Printf("   Positional: %+v\n", p3)
	
	// Using new() - returns pointer
	p4 := new(Person)
	p4.Name = "Charlie"
	p4.Age = 35
	fmt.Printf("   Using new(): %+v\n", *p4)
	fmt.Println()

	// 2. Accessing and modifying fields
	fmt.Println("2. Accessing and Modifying Fields:")
	person := Person{Name: "David", Age: 28}
	fmt.Printf("   Original: %+v\n", person)
	person.Name = "David Smith"
	person.Age = 29
	fmt.Printf("   Modified: %+v\n", person)
	fmt.Println()

	// 3. Struct with different types
	fmt.Println("3. Struct with Different Types:")
	emp := Employee{
		ID:       1001,
		Name:     "Jane Doe",
		Position: "Software Engineer",
		Salary:   75000.50,
		Active:   true,
	}
	fmt.Printf("   Employee: %+v\n", emp)
	fmt.Printf("   Employee Name: %s\n", emp.Name)
	fmt.Printf("   Employee Salary: $%.2f\n", emp.Salary)
	fmt.Println()

	// 4. Nested structs
	fmt.Println("4. Nested Structs:")
	contact := Contact{
		Email: "john@example.com",
		Phone: "555-1234",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
			ZipCode: "10001",
		},
	}
	fmt.Printf("   Contact: %+v\n", contact)
	fmt.Printf("   City: %s\n", contact.Address.City)
	fmt.Println()

	// 5. Anonymous structs
	fmt.Println("5. Anonymous Structs:")
	anon := struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "Anonymous",
	}
	fmt.Printf("   Anonymous struct: %+v\n", anon)
	fmt.Println()

	// 6. Struct comparison
	fmt.Println("6. Struct Comparison:")
	p5 := Person{Name: "Alice", Age: 30}
	p6 := Person{Name: "Alice", Age: 30}
	p7 := Person{Name: "Bob", Age: 30}
	
	fmt.Printf("   p5 == p6: %t (same values)\n", p5 == p6)
	fmt.Printf("   p5 == p7: %t (different values)\n", p5 == p7)
	fmt.Println()

	// 7. Pointer to struct
	fmt.Println("7. Pointer to Struct:")
	ptr := &Person{Name: "Eve", Age: 32}
	fmt.Printf("   Pointer: %+v\n", *ptr)
	fmt.Printf("   Name via pointer: %s\n", ptr.Name) // Go auto-dereferences
	fmt.Printf("   Name explicit: %s\n", (*ptr).Name)
	
	ptr.Age = 33 // Modify through pointer
	fmt.Printf("   After modification: %+v\n", *ptr)
	fmt.Println()

	// 8. Struct as function parameter
	fmt.Println("8. Struct as Function Parameter:")
	displayPerson(p2)
	modifyPerson(p2)
	fmt.Printf("   After modifyPerson (by value): %+v (unchanged)\n", p2)
	modifyPersonByPointer(&p2)
	fmt.Printf("   After modifyPersonByPointer: %+v (changed)\n", p2)
	fmt.Println()

	// 9. Struct methods (preview)
	fmt.Println("9. Struct Methods (Preview):")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle: %+v\n", rect)
	fmt.Printf("   Area: %d\n", rect.Area())
	fmt.Println()

	// 10. Zero values
	fmt.Println("10. Zero Values:")
	var zeroPerson Person
	var zeroEmployee Employee
	var zeroContact Contact
	fmt.Printf("   Person zero value: %+v\n", zeroPerson)
	fmt.Printf("   Employee zero value: %+v\n", zeroEmployee)
	fmt.Printf("   Contact zero value: %+v\n", zeroContact)
}

// Rectangle struct for method example
type Rectangle struct {
	Width  int
	Height int
}

// Method on Rectangle (will be covered in next topic)
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Function that takes struct by value
func displayPerson(p Person) {
	fmt.Printf("   Displaying: %s, %d years old\n", p.Name, p.Age)
}

// Function that takes struct by value (modification won't affect original)
func modifyPerson(p Person) {
	p.Age = 99
}

// Function that takes struct by pointer (modification affects original)
func modifyPersonByPointer(p *Person) {
	p.Age = 99
}

