package main

import "fmt"

// This program demonstrates methods in Go

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// 1. Value receiver method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 2. Value receiver method (doesn't modify original)
func (r Rectangle) Scale(factor float64) Rectangle {
	return Rectangle{
		Width:  r.Width * factor,
		Height: r.Height * factor,
	}
}

// 3. Pointer receiver method (modifies original)
func (r *Rectangle) ScaleInPlace(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 4. Pointer receiver method (returns receiver for chaining)
func (r *Rectangle) SetDimensions(width, height float64) *Rectangle {
	r.Width = width
	r.Height = height
	return r
}

// Counter struct for demonstrating methods
type Counter struct {
	value int
}

// Value receiver - works on copy
func (c Counter) GetValue() int {
	return c.value
}

// Value receiver - doesn't modify original
func (c Counter) Increment() {
	c.value++ // Only modifies copy
}

// Pointer receiver - modifies original
func (c *Counter) IncrementByPointer() {
	c.value++ // Modifies original
}

// Pointer receiver - modifies original
func (c *Counter) Reset() {
	c.value = 0
}

// Person struct
type Person struct {
	Name string
	Age  int
}

// Value receiver method
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Pointer receiver method
func (p *Person) HaveBirthday() {
	p.Age++
}

// Method on non-struct type (type alias)
type MyInt int

func (m MyInt) IsEven() bool {
	return m%2 == 0
}

func (m MyInt) Double() MyInt {
	return m * 2
}

func main() {
	fmt.Println("=== Methods ===")
	fmt.Println()

	// 1. Value receiver methods
	fmt.Println("1. Value Receiver Methods:")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle: %+v\n", rect)
	fmt.Printf("   Area: %.2f\n", rect.Area())
	
	scaled := rect.Scale(2.0)
	fmt.Printf("   Scaled (new): %+v\n", scaled)
	fmt.Printf("   Original (unchanged): %+v\n", rect)
	fmt.Println()

	// 2. Pointer receiver methods
	fmt.Println("2. Pointer Receiver Methods:")
	rect2 := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Before ScaleInPlace: %+v\n", rect2)
	rect2.ScaleInPlace(2.0)
	fmt.Printf("   After ScaleInPlace: %+v (modified)\n", rect2)
	fmt.Println()

	// 3. Method call syntax
	fmt.Println("3. Method Call Syntax:")
	rect3 := &Rectangle{Width: 5, Height: 3}
	fmt.Printf("   Area via pointer: %.2f\n", rect3.Area()) // Go auto-dereferences
	fmt.Printf("   Area explicit: %.2f\n", (*rect3).Area())
	rect3.SetDimensions(10, 10)
	fmt.Printf("   After SetDimensions: %+v\n", *rect3)
	fmt.Println()

	// 4. Value vs pointer receiver
	fmt.Println("4. Value vs Pointer Receiver:")
	counter1 := Counter{value: 0}
	fmt.Printf("   Counter value: %d\n", counter1.GetValue())
	counter1.Increment() // Value receiver - doesn't modify
	fmt.Printf("   After Increment (value receiver): %d (unchanged)\n", counter1.GetValue())
	counter1.IncrementByPointer() // Pointer receiver - modifies
	fmt.Printf("   After IncrementByPointer: %d (changed)\n", counter1.GetValue())
	fmt.Println()

	// 5. Methods on structs
	fmt.Println("5. Methods on Structs:")
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("   %s\n", person.GetInfo())
	person.HaveBirthday()
	fmt.Printf("   After birthday: %s\n", person.GetInfo())
	fmt.Println()

	// 6. Methods on type aliases
	fmt.Println("6. Methods on Type Aliases:")
	num := MyInt(5)
	fmt.Printf("   Number: %d\n", num)
	fmt.Printf("   Is even: %t\n", num.IsEven())
	fmt.Printf("   Double: %d\n", num.Double())
	fmt.Println()

	// 7. Method sets and interfaces
	fmt.Println("7. Method Sets:")
	fmt.Println("   Value receiver: Can be called on value or pointer")
	fmt.Println("   Pointer receiver: Can be called on value or pointer")
	fmt.Println("   Go automatically handles conversion")
	fmt.Println()

	// 8. When to use value vs pointer receiver
	fmt.Println("8. When to Use Value vs Pointer Receiver:")
	fmt.Println("   Value receiver:")
	fmt.Println("     - Method doesn't modify receiver")
	fmt.Println("     - Receiver is small (no copying overhead)")
	fmt.Println("     - Immutability desired")
	fmt.Println("   Pointer receiver:")
	fmt.Println("     - Method modifies receiver")
	fmt.Println("     - Receiver is large (avoid copying)")
	fmt.Println("     - Consistency (if one method uses pointer, all should)")
	fmt.Println()

	// 9. Method chaining (with pointer receivers)
	fmt.Println("9. Method Chaining:")
	rect4 := &Rectangle{Width: 2, Height: 2}
	rect4.SetDimensions(5, 5).ScaleInPlace(2) // Chaining works with pointer receivers
	fmt.Printf("   Chained operations: %+v\n", *rect4)
}

