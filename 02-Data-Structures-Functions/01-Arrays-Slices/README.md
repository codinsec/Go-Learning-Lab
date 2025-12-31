# Arrays & Slices

Learn about arrays and slices in Go, their differences, and how to work with them effectively.

## Learning Objectives

- Understand the difference between arrays and slices
- Learn how to create and manipulate arrays
- Master slice operations: append, copy, slicing
- Understand slice capacity and growth
- Learn when to use arrays vs slices

## Concepts Covered

### Arrays

Arrays in Go are **fixed-size** collections of elements of the same type. The size is part of the array's type.

```go
var arr [5]int                    // Zero-initialized array of 5 integers
arr := [5]int{1, 2, 3, 4, 5}     // Array with values
arr := [...]int{10, 20, 30}       // Array with inferred size
```

**Key characteristics:**
- Fixed size (cannot be resized)
- Value type (copied when assigned)
- Size is part of the type: `[5]int` and `[10]int` are different types
- Less commonly used in Go

### Slices

Slices are **dynamic** views into arrays. They are reference types and can grow/shrink.

```go
var slice []int                   // nil slice
slice := []int{1, 2, 3}          // Slice literal
slice := make([]int, 5)          // Slice with length 5
slice := make([]int, 3, 10)      // Slice with length 3, capacity 10
```

**Key characteristics:**
- Dynamic size (can grow/shrink)
- Reference type (points to underlying array)
- More commonly used than arrays
- Built on top of arrays

### Slice Operations

**Append:**
```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)              // Append single values
slice = append(slice, anotherSlice...)    // Append another slice
```

**Copy:**
```go
source := []int{1, 2, 3, 4, 5}
dest := make([]int, len(source))
copied := copy(dest, source)  // Returns number of elements copied
```

**Slicing:**
```go
slice := []int{0, 1, 2, 3, 4, 5}
slice[2:5]   // [2, 3, 4] - from index 2 to 4 (exclusive)
slice[:5]    // [0, 1, 2, 3, 4] - from start to index 4
slice[2:]    // [2, 3, 4, 5] - from index 2 to end
slice[:]     // Full slice
```

### Slice Capacity

- **Length**: Number of elements in the slice
- **Capacity**: Maximum number of elements without reallocation
- When capacity is exceeded, Go automatically allocates a larger array (typically doubles)

## Running the Example

```bash
# Navigate to this directory
cd 01-Arrays-Slices

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/02-data-structures-functions/arrays-slices

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover
```

## Key Takeaways

1. **Arrays are fixed-size**, slices are dynamic
2. **Arrays are value types**, slices are reference types
3. **Slices are built on arrays** - they're views into underlying arrays
4. **Use slices in most cases** - arrays are rarely needed
5. **Append can trigger reallocation** - capacity grows automatically
6. **Copy creates independent slices** - changes to one don't affect the other
7. **Slice operations are efficient** - slicing doesn't copy data, just creates a new view

## Common Patterns

**Initialize empty slice:**
```go
var slice []int              // nil slice
slice := []int{}             // Empty slice (not nil)
slice := make([]int, 0)      // Empty slice with capacity 0
```

**Pre-allocate with capacity:**
```go
slice := make([]int, 0, 100)  // Length 0, capacity 100
```

**Check if slice is empty:**
```go
if len(slice) == 0 { ... }
```

**Iterate over slice:**
```go
for i, v := range slice {
    // i is index, v is value
}
```

## Best Practices

- Prefer **slices** over arrays in most cases
- Use **make** with capacity when you know the approximate size
- Use **copy** when you need an independent copy
- Be careful with **slice slicing** - it shares the underlying array
- Use **append** for adding elements (it handles capacity automatically)

## Next Steps

After understanding arrays and slices, proceed to:
- [02-Maps](../02-Maps/README.md) - Learn about maps (key-value pairs)

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

