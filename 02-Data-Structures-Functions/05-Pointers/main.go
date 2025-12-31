package main

import "fmt"

// This program demonstrates pointers in Go
func main() {
	fmt.Println("=== Pointers ===")
	fmt.Println()

	// 1. Basic pointer operations
	fmt.Println("1. Basic Pointer Operations:")
	x := 42
	var p *int = &x  // p is a pointer to x
	fmt.Printf("   x = %d\n", x)
	fmt.Printf("   p (address) = %p\n", p)
	fmt.Printf("   *p (value) = %d\n", *p)
	fmt.Println()

	// 2. Zero value of pointer
	fmt.Println("2. Zero Value of Pointer:")
	var nilPtr *int
	fmt.Printf("   nilPtr = %v (nil: %t)\n", nilPtr, nilPtr == nil)
	// *nilPtr would cause panic - don't dereference nil pointer
	fmt.Println()

	// 3. Creating pointer with new
	fmt.Println("3. Creating Pointer with new():")
	ptr := new(int)  // Allocates memory and returns pointer
	fmt.Printf("   ptr = %p\n", ptr)
	fmt.Printf("   *ptr (zero value) = %d\n", *ptr)
	*ptr = 100
	fmt.Printf("   *ptr (after assignment) = %d\n", *ptr)
	fmt.Println()

	// 4. Pointer to pointer
	fmt.Println("4. Pointer to Pointer:")
	value := 10
	p1 := &value
	p2 := &p1
	fmt.Printf("   value = %d\n", value)
	fmt.Printf("   p1 = %p (points to value)\n", p1)
	fmt.Printf("   p2 = %p (points to p1)\n", p2)
	fmt.Printf("   *p1 = %d\n", *p1)
	fmt.Printf("   **p2 = %d\n", **p2)
	fmt.Println()

	// 5. Passing by value vs by reference
	fmt.Println("5. Passing by Value vs by Reference:")
	a, b := 10, 20
	fmt.Printf("   Before swap: a=%d, b=%d\n", a, b)
	swapByValue(a, b)
	fmt.Printf("   After swapByValue: a=%d, b=%d (unchanged)\n", a, b)
	swapByReference(&a, &b)
	fmt.Printf("   After swapByReference: a=%d, b=%d (swapped)\n", a, b)
	fmt.Println()

	// 6. Pointer receiver methods
	fmt.Println("6. Pointer Receiver Methods:")
	counter := Counter{value: 0}
	fmt.Printf("   Initial value: %d\n", counter.value)
	counter.Increment()  // Method with value receiver
	fmt.Printf("   After Increment (value receiver): %d\n", counter.value)
	counter.IncrementByPointer()  // Method with pointer receiver
	fmt.Printf("   After IncrementByPointer: %d\n", counter.value)
	fmt.Println()

	// 7. Returning pointer
	fmt.Println("7. Returning Pointer:")
	ptr1 := createInt(42)
	fmt.Printf("   Created int pointer: %p, value: %d\n", ptr1, *ptr1)
	fmt.Println()

	// 8. Pointer to struct
	fmt.Println("8. Pointer to Struct:")
	person := Person{Name: "Alice", Age: 30}
	personPtr := &person
	fmt.Printf("   person.Name = %s\n", person.Name)
	fmt.Printf("   personPtr.Name = %s\n", personPtr.Name)  // Go auto-dereferences
	fmt.Printf("   (*personPtr).Name = %s\n", (*personPtr).Name)  // Explicit dereference
	personPtr.Age = 31  // Can modify through pointer
	fmt.Printf("   After modification: person.Age = %d\n", person.Age)
	fmt.Println()

	// 9. Pointer to array/slice
	fmt.Println("9. Pointer to Array/Slice:")
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Printf("   arr = %v\n", arr)
	fmt.Printf("   arrPtr[0] = %d\n", arrPtr[0])  // Go auto-dereferences
	(*arrPtr)[1] = 99  // Explicit dereference needed for modification
	fmt.Printf("   After modification: arr = %v\n", arr)
	fmt.Println()

	// 10. When to use pointers
	fmt.Println("10. When to Use Pointers:")
	fmt.Println("   - When you need to modify a value in a function")
	fmt.Println("   - When passing large structs (avoid copying)")
	fmt.Println("   - When a value can be nil (optional values)")
	fmt.Println("   - When implementing methods that modify the receiver")
}

// Counter struct for demonstrating pointer receivers
type Counter struct {
	value int
}

// Value receiver - works on a copy
func (c Counter) Increment() {
	c.value++  // Modifies copy, not original
}

// Pointer receiver - works on original
func (c *Counter) IncrementByPointer() {
	c.value++  // Modifies original
}

// Person struct for demonstrating struct pointers
type Person struct {
	Name string
	Age  int
}

// Function that takes value (copy)
func swapByValue(x, y int) {
	x, y = y, x  // Only swaps local copies
}

// Function that takes pointers (references)
func swapByReference(x, y *int) {
	*x, *y = *y, *x  // Swaps actual values
}

// Function that returns a pointer
func createInt(value int) *int {
	return &value  // Returns pointer to local variable (safe in Go)
}

