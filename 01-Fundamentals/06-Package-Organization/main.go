package main

import (
	"fmt"
	"github.com/codinsec/go-learning-lab/01-fundamentals/package-organization/calculator"
	"github.com/codinsec/go-learning-lab/01-fundamentals/package-organization/greeter"
)

// This program demonstrates package organization and visibility
func main() {
	fmt.Println("=== Package Organization ===")
	fmt.Println()

	// Using exported functions from other packages
	fmt.Println("1. Using Exported Functions:")
	sum := calculator.Add(10, 20)
	fmt.Printf("   calculator.Add(10, 20) = %d\n", sum)

	product := calculator.Multiply(5, 4)
	fmt.Printf("   calculator.Multiply(5, 4) = %d\n", product)

	// Using exported types
	fmt.Println("\n2. Using Exported Types:")
	calc := calculator.NewCalculator()
	result := calc.Subtract(100, 30)
	fmt.Printf("   calc.Subtract(100, 30) = %d\n", result)

	// Using greeter package
	fmt.Println("\n3. Using Greeter Package:")
	greeter.SayHello("Go Developer")
	greeter.SayGoodbye("Go Developer")

	// Demonstrating visibility rules
	fmt.Println("\n4. Visibility Rules:")
	fmt.Println("   - Uppercase names are EXPORTED (public)")
	fmt.Println("   - Lowercase names are UNEXPORTED (private)")
	fmt.Println("   - Exported: Add, Multiply, NewCalculator")
	fmt.Println("   - Unexported: add, multiply, newCalculator (not accessible)")
}

