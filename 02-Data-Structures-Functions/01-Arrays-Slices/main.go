package main

import "fmt"

// This program demonstrates arrays and slices in Go
func main() {
	fmt.Println("=== Arrays & Slices ===")
	fmt.Println()

	// 1. Arrays - fixed size
	fmt.Println("1. Arrays (Fixed Size):")
	var arr1 [5]int                    // Zero-initialized array
	arr2 := [5]int{1, 2, 3, 4, 5}      // Array with values
	arr3 := [...]int{10, 20, 30}        // Array with inferred size
	fmt.Printf("   arr1: %v\n", arr1)
	fmt.Printf("   arr2: %v\n", arr2)
	fmt.Printf("   arr3: %v\n", arr3)
	fmt.Printf("   arr2 length: %d\n", len(arr2))
	fmt.Println()

	// 2. Slices - dynamic size
	fmt.Println("2. Slices (Dynamic Size):")
	var slice1 []int                    // nil slice
	slice2 := []int{1, 2, 3, 4, 5}       // Slice literal
	slice3 := make([]int, 5)             // Slice with make (length 5)
	slice4 := make([]int, 3, 10)         // Slice with make (length 3, capacity 10)
	fmt.Printf("   slice1: %v (nil: %t)\n", slice1, slice1 == nil)
	fmt.Printf("   slice2: %v\n", slice2)
	fmt.Printf("   slice3: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("   slice4: %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))
	fmt.Println()

	// 3. Slice from array
	fmt.Println("3. Slice from Array:")
	arr := [5]int{10, 20, 30, 40, 50}
	sliceFromArr := arr[1:4]            // Slice from index 1 to 3 (exclusive)
	fmt.Printf("   Original array: %v\n", arr)
	fmt.Printf("   Slice arr[1:4]: %v\n", sliceFromArr)
	fmt.Println()

	// 4. Slice operations
	fmt.Println("4. Slice Operations:")
	numbers := []int{1, 2, 3}
	fmt.Printf("   Original: %v\n", numbers)

	// Append
	numbers = append(numbers, 4, 5)
	fmt.Printf("   After append(4, 5): %v\n", numbers)

	// Append another slice
	moreNumbers := []int{6, 7, 8}
	numbers = append(numbers, moreNumbers...)
	fmt.Printf("   After append slice: %v\n", numbers)
	fmt.Println()

	// 5. Copy
	fmt.Println("5. Copy Operation:")
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(source))
	copied := copy(dest, source)
	fmt.Printf("   Source: %v\n", source)
	fmt.Printf("   Dest: %v\n", dest)
	fmt.Printf("   Elements copied: %d\n", copied)

	// Modify source to show independence
	source[0] = 99
	fmt.Printf("   After modifying source[0] = 99:\n")
	fmt.Printf("   Source: %v\n", source)
	fmt.Printf("   Dest: %v (unchanged)\n", dest)
	fmt.Println()

	// 6. Slice capacity and growth
	fmt.Println("6. Slice Capacity and Growth:")
	slice := make([]int, 0, 3) // length 0, capacity 3
	fmt.Printf("   Initial: len=%d, cap=%d, %v\n", len(slice), cap(slice), slice)
	
	slice = append(slice, 1)
	fmt.Printf("   After append(1): len=%d, cap=%d, %v\n", len(slice), cap(slice), slice)
	
	slice = append(slice, 2, 3)
	fmt.Printf("   After append(2, 3): len=%d, cap=%d, %v\n", len(slice), cap(slice), slice)
	
	slice = append(slice, 4) // Capacity will grow
	fmt.Printf("   After append(4): len=%d, cap=%d, %v\n", len(slice), cap(slice), slice)
	fmt.Println()

	// 7. Slice slicing
	fmt.Println("7. Slice Slicing:")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("   Original: %v\n", original)
	fmt.Printf("   original[2:5]: %v\n", original[2:5])   // [2, 3, 4]
	fmt.Printf("   original[:5]: %v\n", original[:5])      // [0, 1, 2, 3, 4]
	fmt.Printf("   original[5:]: %v\n", original[5:])      // [5, 6, 7, 8, 9]
	fmt.Printf("   original[:]: %v\n", original[:])        // Full slice
	fmt.Println()

	// 8. Key differences
	fmt.Println("8. Key Differences:")
	fmt.Println("   Arrays:")
	fmt.Println("     - Fixed size (size is part of type)")
	fmt.Println("     - Value type (copied when assigned)")
	fmt.Println("     - Less commonly used")
	fmt.Println("   Slices:")
	fmt.Println("     - Dynamic size")
	fmt.Println("     - Reference type (points to underlying array)")
	fmt.Println("     - Commonly used in Go")
}

