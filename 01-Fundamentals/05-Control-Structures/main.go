package main

import "fmt"

// This program demonstrates control structures in Go
func main() {
	fmt.Println("=== Control Structures ===")
	fmt.Println()

	// 1. If statement
	fmt.Println("1. If Statement:")
	x := 10
	if x > 5 {
		fmt.Println("  x is greater than 5")
	}

	// 2. If-else statement
	fmt.Println("\n2. If-Else Statement:")
	age := 18
	if age >= 18 {
		fmt.Println("  You are an adult")
	} else {
		fmt.Println("  You are a minor")
	}

	// 3. If-else if-else statement
	fmt.Println("\n3. If-Else If-Else Statement:")
	score := 85
	if score >= 90 {
		fmt.Println("  Grade: A")
	} else if score >= 80 {
		fmt.Println("  Grade: B")
	} else if score >= 70 {
		fmt.Println("  Grade: C")
	} else {
		fmt.Println("  Grade: F")
	}

	// 4. If with initialization statement
	fmt.Println("\n4. If with Initialization:")
	if num := 42; num%2 == 0 {
		fmt.Printf("  %d is even\n", num)
	}
	// Note: 'num' is only available in the if block

	// 5. For loop - basic
	fmt.Println("\n5. For Loop - Basic:")
	for i := 0; i < 3; i++ {
		fmt.Printf("  Iteration %d\n", i)
	}

	// 6. For loop - while style
	fmt.Println("\n6. For Loop - While Style:")
	count := 0
	for count < 3 {
		fmt.Printf("  Count: %d\n", count)
		count++
	}

	// 7. For loop - infinite (with break)
	fmt.Println("\n7. For Loop - Infinite (with break):")
	iterations := 0
	for {
		if iterations >= 3 {
			break
		}
		fmt.Printf("  Infinite loop iteration %d\n", iterations)
		iterations++
	}

	// 8. For loop - continue
	fmt.Println("\n8. For Loop - Continue:")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("  Odd number: %d\n", i)
	}

	// 9. For range - arrays/slices
	fmt.Println("\n9. For Range - Arrays/Slices:")
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Printf("  Index %d: %d\n", index, value)
	}

	// 10. For range - maps
	fmt.Println("\n10. For Range - Maps:")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for key, value := range colors {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// 11. For range - strings
	fmt.Println("\n11. For Range - Strings:")
	text := "Go"
	for index, char := range text {
		fmt.Printf("  Index %d: %c (rune: %d)\n", index, char, char)
	}

	// 12. Switch statement - basic
	fmt.Println("\n12. Switch Statement - Basic:")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("  Start of the work week")
	case "Friday":
		fmt.Println("  End of the work week")
	default:
		fmt.Println("  Mid week")
	}

	// 13. Switch statement - multiple values
	fmt.Println("\n13. Switch - Multiple Values:")
	character := "a"
	switch character {
	case "a", "e", "i", "o", "u":
		fmt.Println("  It's a vowel")
	default:
		fmt.Println("  It's a consonant")
	}

	// 14. Switch statement - with initialization
	fmt.Println("\n14. Switch - With Initialization:")
	switch num := 7; num {
	case 1, 3, 5, 7, 9:
		fmt.Println("  Odd number")
	case 2, 4, 6, 8:
		fmt.Println("  Even number")
	default:
		fmt.Println("  Zero or negative")
	}

	// 15. Switch statement - no expression (like if-else)
	fmt.Println("\n15. Switch - No Expression:")
	value := 42
	switch {
	case value < 0:
		fmt.Println("  Negative")
	case value == 0:
		fmt.Println("  Zero")
	case value > 0 && value < 100:
		fmt.Println("  Positive and less than 100")
	default:
		fmt.Println("  Greater than or equal to 100")
	}
}

