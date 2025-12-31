# Control Structures

Learn about Go's control flow: if statements, for loops, and switch statements.

## Learning Objectives

- Understand if/else statements and variations
- Master the for loop (Go's only loop)
- Learn switch statements and their flexibility
- Understand range loops for collections
- Learn break and continue statements

## Concepts Covered

### If Statements

**Basic if:**
```go
if condition {
    // code
}
```

**If-else:**
```go
if condition {
    // code
} else {
    // code
}
```

**If with initialization:**
```go
if x := getValue(); x > 0 {
    // x is only available in this block
}
```

### For Loop (Go's Only Loop)

**Traditional for:**
```go
for i := 0; i < 10; i++ {
    // code
}
```

**While-style:**
```go
for condition {
    // code
}
```

**Infinite loop:**
```go
for {
    // code
    if condition {
        break
    }
}
```

**Range loop:**
```go
for index, value := range slice {
    // code
}
```

### Switch Statements

**Basic switch:**
```go
switch value {
case 1:
    // code
case 2:
    // code
default:
    // code
}
```

**Switch with multiple values:**
```go
switch value {
case 1, 2, 3:
    // code
}
```

**Switch with initialization:**
```go
switch x := getValue(); x {
case 1:
    // code
}
```

**Switch without expression (like if-else):**
```go
switch {
case x < 0:
    // code
case x == 0:
    // code
default:
    // code
}
```

## Running the Example

```bash
# Navigate to this directory
cd 05-Control-Structures

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/01-fundamentals/control-structures

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

1. **Go has only one loop**: `for` (but it's very flexible)
2. **If with initialization** keeps variables scoped to the block
3. **Switch is powerful** - can work like if-else chains
4. **Range loops** are idiomatic for iterating collections
5. **Break and continue** work in loops (not in switch by default)

## Important Notes

- **No while loop**: Use `for condition {}` instead
- **No do-while**: Use `for { ... if condition { break } }`
- **Switch doesn't fall through**: Each case ends automatically (use `fallthrough` if needed)
- **Range returns index and value**: Use `_` to ignore either

## Best Practices

- Use `for range` for slices, maps, and strings
- Prefer switch over long if-else chains
- Use initialization in if/switch to limit variable scope
- Use meaningful variable names in loops
- Consider readability over cleverness

## Next Steps

After understanding control structures, proceed to:
- [06-Package-Organization](../06-Package-Organization/README.md) - Learn about package visibility and organization

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

