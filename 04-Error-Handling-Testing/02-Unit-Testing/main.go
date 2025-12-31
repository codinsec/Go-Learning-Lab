package main

import (
	"fmt"
	"math"
)

// This file contains functions to test

// Add adds two integers
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts b from a
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide divides a by b, returns error if b is 0
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// CalculateArea calculates the area of a circle
func CalculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// Factorial calculates factorial of n
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	fmt.Println("=== Unit Testing ===")
	fmt.Println()
	fmt.Println("This file contains functions to test.")
	fmt.Println("Run tests with: go test")
	fmt.Println("Run tests with verbose output: go test -v")
	fmt.Println("Run tests with coverage: go test -cover")
	fmt.Println("Run specific test: go test -run TestAdd")
}

