package main

import (
	"fmt"
)

// This program demonstrates generics in Go (Go 1.18+)

// 1. Generic function
func Print[T any](value T) {
	fmt.Println(value)
}

// 2. Generic function with constraints
func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 3. Generic function with type constraint
type Number interface {
	int | int64 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

// 4. Generic struct
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// 5. Generic method
type Container[T any] struct {
	value T
}

func (c Container[T]) Get() T {
	return c.value
}

func (c *Container[T]) Set(value T) {
	c.value = value
}

// 6. Type constraints
type Comparable interface {
	int | string | float64
}

func Compare[T Comparable](a, b T) bool {
	return a == b
}

// 7. Multiple type parameters
func Pair[T, U any](first T, second U) (T, U) {
	return first, second
}

// 8. Generic with interface constraint
type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

// 9. Type set notation
type Numeric interface {
	~int | ~int64 | ~float64 // ~ allows underlying types
}

func Multiply[T Numeric](a, b T) T {
	return a * b
}

// Person type for interface constraint example
type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func main() {
	fmt.Println("=== Generics ===")
	fmt.Println()

	// 1. Generic function
	fmt.Println("1. Generic Function:")
	Print("Hello")
	Print(42)
	Print(3.14)
	fmt.Println()

	// 2. Generic with constraints
	fmt.Println("2. Generic with Constraints:")
	fmt.Printf("   Max(10, 20) = %d\n", Max(10, 20))
	fmt.Printf("   Max(3.14, 2.71) = %.2f\n", Max(3.14, 2.71))
	fmt.Println()

	// 3. Generic with type constraint
	fmt.Println("3. Generic with Type Constraint:")
	fmt.Printf("   Add(5, 3) = %d\n", Add(5, 3))
	fmt.Printf("   Add(3.14, 2.71) = %.2f\n", Add(3.14, 2.71))
	fmt.Println()

	// 4. Generic struct
	fmt.Println("4. Generic Struct (Stack):")
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Printf("   Stack size: %d\n", intStack.Size())
	
	val, ok := intStack.Pop()
	fmt.Printf("   Popped: %d (ok: %t)\n", val, ok)
	fmt.Printf("   Stack size: %d\n", intStack.Size())
	
	stringStack := NewStack[string]()
	stringStack.Push("first")
	stringStack.Push("second")
	fmt.Printf("   String stack size: %d\n", stringStack.Size())
	fmt.Println()

	// 5. Generic method
	fmt.Println("5. Generic Method:")
	container := Container[int]{value: 42}
	fmt.Printf("   Container value: %d\n", container.Get())
	container.Set(100)
	fmt.Printf("   Container value after Set: %d\n", container.Get())
	fmt.Println()

	// 6. Type constraints
	fmt.Println("6. Type Constraints:")
	fmt.Printf("   Compare(5, 5): %t\n", Compare(5, 5))
	fmt.Printf("   Compare(\"hello\", \"world\"): %t\n", Compare("hello", "world"))
	fmt.Println()

	// 7. Multiple type parameters
	fmt.Println("7. Multiple Type Parameters:")
	first, second := Pair("Alice", 30)
	fmt.Printf("   Pair: %s, %d\n", first, second)
	fmt.Println()

	// 8. Interface constraint
	fmt.Println("8. Interface Constraint:")
	person := Person{Name: "Bob"}
	PrintString(person)
	fmt.Println()

	// 9. Type set notation
	fmt.Println("9. Type Set Notation:")
	type MyInt int
	var a MyInt = 5
	var b MyInt = 3
	result := Multiply(a, b)
	fmt.Printf("   Multiply(MyInt(5), MyInt(3)) = %d\n", result)
	fmt.Println()

	// 10. Key concepts
	fmt.Println("10. Key Concepts:")
	fmt.Println("   - Generics enable type-safe code reuse")
	fmt.Println("   - Type parameters: [T any]")
	fmt.Println("   - Type constraints: [T Number]")
	fmt.Println("   - Type inference: Go infers types when possible")
	fmt.Println("   - Available in Go 1.18+")
}

