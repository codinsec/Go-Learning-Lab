package main

import "fmt"

// This program demonstrates variables, constants, and types in Go
func main() {
	fmt.Println("=== Variables & Types ===")
	fmt.Println()

	// 1. Variable declaration with var
	var name string = "Go"
	var version int = 1
	fmt.Printf("Language: %s, Version: %d\n", name, version)

	// 2. Type inference (Go infers the type)
	var language = "Golang"
	var year = 2009
	fmt.Printf("Language: %s, Year: %d\n", language, year)

	// 3. Short variable declaration (most common)
	author := "Google"
	openSource := true
	fmt.Printf("Author: %s, Open Source: %s\n", author, fmt.Sprintf("%t", openSource))

	// 4. Multiple variable declaration
	var x, y int = 10, 20
	fmt.Printf("x = %d, y = %d\n", x, y)

	// 5. Short declaration for multiple variables
	a, b := 5, 15
	fmt.Printf("a = %d, b = %d\n", a, b)

	// 6. Constants
	const pi = 3.14159
	const greeting = "Hello, Go!"
	fmt.Printf("Pi: %f\n", pi)
	fmt.Printf("Greeting: %s\n", greeting)

	// 7. Typed constants
	const maxInt int = 100
	const minFloat float64 = 0.0
	fmt.Printf("Max Int: %d, Min Float: %f\n", maxInt, minFloat)

	// 8. Multiple constants
	const (
		statusOK    = 200
		statusError = 500
	)
	fmt.Printf("Status OK: %d, Status Error: %d\n", statusOK, statusError)

	// 9. Zero values (default values)
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	fmt.Println("\nZero values:")
	fmt.Printf("  int: %d\n", zeroInt)
	fmt.Printf("  float64: %f\n", zeroFloat)
	fmt.Printf("  string: %q\n", zeroString)
	fmt.Printf("  bool: %t\n", zeroBool)

	// 10. Type conversion
	var integer int = 42
	var float float64 = float64(integer)
	var uintValue uint = uint(integer)
	fmt.Println("\nType conversions:")
	fmt.Printf("  int to float64: %d -> %f\n", integer, float)
	fmt.Printf("  int to uint: %d -> %d\n", integer, uintValue)
}

