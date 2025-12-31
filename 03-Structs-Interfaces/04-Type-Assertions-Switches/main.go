package main

import "fmt"

// This program demonstrates type assertions and type switches in Go

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Triangle implements Shape
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	fmt.Println("=== Type Assertions & Type Switches ===")
	fmt.Println()

	// 1. Basic type assertion
	fmt.Println("1. Basic Type Assertion:")
	var i interface{} = "Hello, World!"
	
	str, ok := i.(string)
	if ok {
		fmt.Printf("   Asserted to string: %s\n", str)
	} else {
		fmt.Println("   Not a string")
	}
	
	num, ok := i.(int)
	if ok {
		fmt.Printf("   Asserted to int: %d\n", num)
	} else {
		fmt.Println("   Not an int")
	}
	fmt.Println()

	// 2. Type assertion without ok check (can panic)
	fmt.Println("2. Type Assertion Without OK Check:")
	var v interface{} = 42
	num2 := v.(int) // Safe because we know it's int
	fmt.Printf("   Direct assertion: %d\n", num2)
	
	// This would panic if v is not int:
	// str2 := v.(string) // panic!
	fmt.Println()

	// 3. Type assertion with interface
	fmt.Println("3. Type Assertion with Interface:")
	var shape Shape = Rectangle{Width: 10, Height: 5}
	
	rect, ok := shape.(Rectangle)
	if ok {
		fmt.Printf("   It's a Rectangle: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
	}
	
	circle, ok := shape.(Circle)
	if ok {
		fmt.Printf("   It's a Circle: Radius=%.2f\n", circle.Radius)
	} else {
		fmt.Println("   It's not a Circle")
	}
	fmt.Println()

	// 4. Type switch - basic
	fmt.Println("4. Type Switch - Basic:")
	describeType(42)
	describeType("hello")
	describeType(3.14)
	describeType(true)
	fmt.Println()

	// 5. Type switch with interface
	fmt.Println("5. Type Switch with Interface:")
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 5},
		Triangle{Base: 10, Height: 8},
	}
	
	for _, shape := range shapes {
		describeShape(shape)
	}
	fmt.Println()

	// 6. Type switch with multiple cases
	fmt.Println("6. Type Switch with Multiple Cases:")
	processValue(42)
	processValue("hello")
	processValue(3.14)
	processValue([]int{1, 2, 3})
	fmt.Println()

	// 7. Type switch with default case
	fmt.Println("7. Type Switch with Default Case:")
	handleValue(42)
	handleValue("hello")
	handleValue(3.14)
	handleValue(map[string]int{"a": 1})
	fmt.Println()

	// 8. Type assertion for pointer types
	fmt.Println("8. Type Assertion for Pointer Types:")
	var ptr interface{} = &Rectangle{Width: 5, Height: 3}
	
	rectPtr, ok := ptr.(*Rectangle)
	if ok {
		fmt.Printf("   Rectangle pointer: Width=%.2f, Height=%.2f\n", rectPtr.Width, rectPtr.Height)
	}
	fmt.Println()

	// 9. Type switch with nil check
	fmt.Println("9. Type Switch with Nil Check:")
	var nilShape Shape
	describeShape(nilShape)
	
	var nilPtr *Rectangle
	checkNil(nilPtr)
	fmt.Println()

	// 10. Practical example: JSON unmarshaling
	fmt.Println("10. Practical Example:")
	processJSONData()
}

// Function using type switch
func describeType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("   Integer: %d\n", v)
	case string:
		fmt.Printf("   String: %s\n", v)
	case float64:
		fmt.Printf("   Float: %.2f\n", v)
	case bool:
		fmt.Printf("   Boolean: %t\n", v)
	default:
		fmt.Printf("   Unknown type: %T\n", v)
	}
}

// Function using type switch with interface
func describeShape(s Shape) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Printf("   Rectangle: Width=%.2f, Height=%.2f, Area=%.2f\n", 
			v.Width, v.Height, v.Area())
	case Circle:
		fmt.Printf("   Circle: Radius=%.2f, Area=%.2f\n", 
			v.Radius, v.Area())
	case Triangle:
		fmt.Printf("   Triangle: Base=%.2f, Height=%.2f, Area=%.2f\n", 
			v.Base, v.Height, v.Area())
	case nil:
		fmt.Println("   Nil shape")
	default:
		fmt.Printf("   Unknown shape: %T\n", v)
	}
}

// Function with multiple type cases
func processValue(v interface{}) {
	switch v := v.(type) {
	case int, int32, int64:
		fmt.Printf("   Integer type: %v\n", v)
	case string:
		fmt.Printf("   String: %s\n", v)
	case float32, float64:
		fmt.Printf("   Float type: %v\n", v)
	case []int:
		fmt.Printf("   Int slice: %v\n", v)
	default:
		fmt.Printf("   Other type: %T\n", v)
	}
}

// Function with default case
func handleValue(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("   Handling as integer")
	case string:
		fmt.Println("   Handling as string")
	case float64:
		fmt.Println("   Handling as float")
	default:
		fmt.Printf("   Handling as unknown type: %T\n", v)
	}
}

// Function checking for nil
func checkNil(v interface{}) {
	if v == nil {
		fmt.Println("   Value is nil")
		return
	}
	
	switch v.(type) {
	case *Rectangle:
		fmt.Println("   It's a *Rectangle (not nil)")
	case *Circle:
		fmt.Println("   It's a *Circle (not nil)")
	default:
		fmt.Printf("   Unknown pointer type: %T\n", v)
	}
}

// Practical example: simulating JSON data processing
func processJSONData() {
	// Simulating different JSON data types
	data := []interface{}{
		map[string]interface{}{"name": "Alice", "age": 30},
		[]interface{}{1, 2, 3, 4, 5},
		"Hello, World!",
		42,
	}
	
	for _, item := range data {
		switch v := item.(type) {
		case map[string]interface{}:
			fmt.Printf("   JSON Object: %v\n", v)
		case []interface{}:
			fmt.Printf("   JSON Array: %v\n", v)
		case string:
			fmt.Printf("   JSON String: %s\n", v)
		case int:
			fmt.Printf("   JSON Number: %d\n", v)
		default:
			fmt.Printf("   Unknown JSON type: %T\n", v)
		}
	}
}

