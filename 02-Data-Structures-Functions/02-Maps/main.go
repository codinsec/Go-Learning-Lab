package main

import "fmt"

// This program demonstrates maps (key-value pairs) in Go
func main() {
	fmt.Println("=== Maps (Key-Value Pairs) ===")
	fmt.Println()

	// 1. Map declaration and initialization
	fmt.Println("1. Map Declaration:")
	var m1 map[string]int                    // nil map
	m2 := make(map[string]int)               // Empty map
	m3 := map[string]int{                    // Map literal
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	fmt.Printf("   m1 (nil): %v (nil: %t)\n", m1, m1 == nil)
	fmt.Printf("   m2 (empty): %v\n", m2)
	fmt.Printf("   m3 (literal): %v\n", m3)
	fmt.Println()

	// 2. Adding and accessing elements
	fmt.Println("2. Adding and Accessing Elements:")
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("   Scores: %v\n", scores)
	fmt.Printf("   Alice's score: %d\n", scores["Alice"])
	fmt.Println()

	// 3. Checking if key exists
	fmt.Println("3. Checking Key Existence:")
	value, exists := scores["Bob"]
	if exists {
		fmt.Printf("   Bob's score: %d\n", value)
	} else {
		fmt.Println("   Bob not found")
	}

	value, exists = scores["David"]
	if exists {
		fmt.Printf("   David's score: %d\n", value)
	} else {
		fmt.Println("   David not found")
	}
	fmt.Println()

	// 4. Zero value for missing keys
	fmt.Println("4. Zero Value for Missing Keys:")
	fmt.Printf("   Non-existent key 'Eve': %d (zero value)\n", scores["Eve"])
	fmt.Println()

	// 5. Deleting elements
	fmt.Println("5. Deleting Elements:")
	fmt.Printf("   Before delete: %v\n", scores)
	delete(scores, "Bob")
	fmt.Printf("   After delete(Bob): %v\n", scores)
	fmt.Println()

	// 6. Iterating over maps
	fmt.Println("6. Iterating Over Maps:")
	colors := map[string]string{
		"red":    "#FF0000",
		"green":  "#00FF00",
		"blue":   "#0000FF",
		"yellow": "#FFFF00",
	}
	fmt.Println("   Colors:")
	for key, value := range colors {
		fmt.Printf("     %s: %s\n", key, value)
	}
	fmt.Println()

	// 7. Map length
	fmt.Println("7. Map Length:")
	fmt.Printf("   Number of colors: %d\n", len(colors))
	fmt.Println()

	// 8. Maps with different value types
	fmt.Println("8. Maps with Different Value Types:")
	
	// Map of string to slice
	groups := map[string][]string{
		"fruits":  {"apple", "banana", "orange"},
		"colors":  {"red", "green", "blue"},
		"animals": {"dog", "cat", "bird"},
	}
	fmt.Printf("   Groups: %v\n", groups)
	
	// Map of string to map
	nested := map[string]map[string]int{
		"math":    {"Alice": 95, "Bob": 87},
		"science": {"Alice": 92, "Bob": 90},
	}
	fmt.Printf("   Nested map: %v\n", nested)
	fmt.Println()

	// 9. Map as set (using bool values)
	fmt.Println("9. Map as Set:")
	visited := make(map[string]bool)
	visited["page1"] = true
	visited["page2"] = true
	visited["page3"] = true
	
	if visited["page1"] {
		fmt.Println("   page1 has been visited")
	}
	if !visited["page4"] {
		fmt.Println("   page4 has not been visited")
	}
	fmt.Println()

	// 10. Clearing a map
	fmt.Println("10. Clearing a Map:")
	fmt.Printf("   Before clear: %v (len: %d)\n", scores, len(scores))
	
	// Clear by iterating and deleting
	for key := range scores {
		delete(scores, key)
	}
	fmt.Printf("   After clear: %v (len: %d)\n", scores, len(scores))
	
	// Or reassign to new map
	scores = make(map[string]int)
	fmt.Printf("   After reassign: %v (len: %d)\n", scores, len(scores))
}

