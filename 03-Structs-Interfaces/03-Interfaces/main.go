package main

import "fmt"

// This program demonstrates interfaces in Go

// 1. Basic interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Rectangle implements Shape
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

// 3. Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// 4. Interface with single method
type Writer interface {
	Write([]byte) (int, error)
}

type FileWriter struct {
	Filename string
}

func (f FileWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing to %s: %s\n", f.Filename, string(data))
	return len(data), nil
}

// 5. Interface composition
type Reader interface {
	Read([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

// 6. Empty interface (interface{})
func printAnything(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// 7. Interface with multiple methods
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Walking"
}

// 8. Interface satisfaction (implicit)
// No explicit "implements" keyword needed
// If a type has all methods, it implements the interface

func main() {
	fmt.Println("=== Interfaces ===")
	fmt.Println()

	// 1. Basic interface usage
	fmt.Println("1. Basic Interface Usage:")
	var shape Shape
	
	rect := Rectangle{Width: 10, Height: 5}
	shape = rect
	fmt.Printf("   Rectangle area: %.2f\n", shape.Area())
	fmt.Printf("   Rectangle perimeter: %.2f\n", shape.Perimeter())
	
	circle := Circle{Radius: 5}
	shape = circle
	fmt.Printf("   Circle area: %.2f\n", shape.Area())
	fmt.Printf("   Circle perimeter: %.2f\n", shape.Perimeter())
	fmt.Println()

	// 2. Polymorphism
	fmt.Println("2. Polymorphism:")
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 3},
	}
	
	for i, s := range shapes {
		fmt.Printf("   Shape %d: Area=%.2f, Perimeter=%.2f\n", i+1, s.Area(), s.Perimeter())
	}
	fmt.Println()

	// 3. Interface as parameter
	fmt.Println("3. Interface as Parameter:")
	printShapeInfo(rect)
	printShapeInfo(circle)
	fmt.Println()

	// 4. Empty interface
	fmt.Println("4. Empty Interface (interface{}):")
	printAnything(42)
	printAnything("Hello")
	printAnything(3.14)
	printAnything(true)
	printAnything([]int{1, 2, 3})
	fmt.Println()

	// 5. Type assertion (preview)
	fmt.Println("5. Type Assertion (Preview):")
	var i interface{} = "Hello"
	str, ok := i.(string)
	if ok {
		fmt.Printf("   Asserted to string: %s\n", str)
	}
	fmt.Println()

	// 6. Interface composition
	fmt.Println("6. Interface Composition:")
	var rw ReadWriter = &FileReadWriter{Filename: "test.txt"}
	data := make([]byte, 100)
	n, _ := rw.Read(data)
	fmt.Printf("   Read %d bytes\n", n)
	rw.Write([]byte("Hello, World!"))
	fmt.Println()

	// 7. Multiple interfaces
	fmt.Println("7. Multiple Interfaces:")
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	
	animals := []Animal{dog, cat}
	for _, animal := range animals {
		fmt.Printf("   %s: %s, %s\n", 
			fmt.Sprintf("%T", animal), 
			animal.Speak(), 
			animal.Move())
	}
	fmt.Println()

	// 8. Interface nil value
	fmt.Println("8. Interface Nil Value:")
	var shape2 Shape
	fmt.Printf("   shape2 is nil: %t\n", shape2 == nil)
	// shape2.Area() would cause runtime error (nil pointer)
	fmt.Println()

	// 9. Duck typing
	fmt.Println("9. Duck Typing:")
	fmt.Println("   'If it walks like a duck and quacks like a duck,")
	fmt.Println("    it's a duck.' - Go's interface philosophy")
	fmt.Println("   If a type has all methods of an interface,")
	fmt.Println("   it implements that interface automatically.")
}

// FileReadWriter implements ReadWriter
type FileReadWriter struct {
	Filename string
}

func (f *FileReadWriter) Read(data []byte) (int, error) {
	fmt.Printf("Reading from %s\n", f.Filename)
	copy(data, []byte("Sample data"))
	return 11, nil
}

func (f *FileReadWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing to %s: %s\n", f.Filename, string(data))
	return len(data), nil
}

// Function that takes interface
func printShapeInfo(s Shape) {
	fmt.Printf("   Shape: Area=%.2f, Perimeter=%.2f\n", s.Area(), s.Perimeter())
}

