package main

import "fmt"

// This program demonstrates functions in Go

// 1. Basic function
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 2. Function with return value
func add(a, b int) int {
	return a + b
}

// 3. Function with multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 4. Named return values
func calculate(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return // Naked return - returns named values
}

// 5. Function with error return
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 6. Variadic function (variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 7. Function as value
func multiply(a, b int) int {
	return a * b
}

// 8. Higher-order function (function that takes function as parameter)
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 9. Anonymous function (closure)
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 10. Function with pointer parameters
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("=== Functions ===")
	fmt.Println()

	// 1. Basic function call
	fmt.Println("1. Basic Function:")
	greet("Go Developer")
	fmt.Println()

	// 2. Function with return value
	fmt.Println("2. Function with Return Value:")
	result := add(10, 20)
	fmt.Printf("   add(10, 20) = %d\n", result)
	fmt.Println()

	// 3. Multiple return values
	fmt.Println("3. Multiple Return Values:")
	quotient, remainder := divide(17, 5)
	fmt.Printf("   divide(17, 5): quotient=%d, remainder=%d\n", quotient, remainder)
	fmt.Println()

	// 4. Named return values
	fmt.Println("4. Named Return Values:")
	sum, product := calculate(5, 6)
	fmt.Printf("   calculate(5, 6): sum=%d, product=%d\n", sum, product)
	fmt.Println()

	// 5. Error handling
	fmt.Println("5. Error Handling:")
	result1, err1 := safeDivide(10, 2)
	if err1 != nil {
		fmt.Printf("   Error: %v\n", err1)
	} else {
		fmt.Printf("   safeDivide(10, 2) = %d\n", result1)
	}

	result2, err2 := safeDivide(10, 0)
	if err2 != nil {
		fmt.Printf("   safeDivide(10, 0) error: %v\n", err2)
	} else {
		fmt.Printf("   safeDivide(10, 0) = %d\n", result2)
	}
	fmt.Println()

	// 6. Variadic function
	fmt.Println("6. Variadic Function:")
	fmt.Printf("   sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	fmt.Printf("   sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("   sum() = %d\n", sum())
	fmt.Println()

	// 7. Function as value
	fmt.Println("7. Function as Value:")
	fn := multiply
	fmt.Printf("   multiply(5, 4) = %d\n", fn(5, 4))
	fmt.Println()

	// 8. Higher-order function
	fmt.Println("8. Higher-Order Function:")
	result3 := applyOperation(10, 5, add)
	fmt.Printf("   applyOperation(10, 5, add) = %d\n", result3)
	result4 := applyOperation(10, 5, multiply)
	fmt.Printf("   applyOperation(10, 5, multiply) = %d\n", result4)
	fmt.Println()

	// 9. Closure
	fmt.Println("9. Closure:")
	counter1 := makeCounter()
	counter2 := makeCounter()
	fmt.Printf("   counter1: %d, %d, %d\n", counter1(), counter1(), counter1())
	fmt.Printf("   counter2: %d, %d\n", counter2(), counter2())
	fmt.Println()

	// 10. Pointer parameters
	fmt.Println("10. Pointer Parameters:")
	x, y := 10, 20
	fmt.Printf("   Before swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("   After swap: x=%d, y=%d\n", x, y)
	fmt.Println()

	// 11. Anonymous function inline
	fmt.Println("11. Anonymous Function:")
	func(name string) {
		fmt.Printf("   Anonymous function says: Hello, %s!\n", name)
	}("World")
	fmt.Println()

	// 12. Defer in functions (preview)
	fmt.Println("12. Defer in Functions (Preview):")
	deferExample()
}

func deferExample() {
	defer fmt.Println("   This is deferred (runs last)")
	fmt.Println("   This runs first")
	fmt.Println("   This runs second")
}

