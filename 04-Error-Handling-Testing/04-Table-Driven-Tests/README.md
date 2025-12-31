# Table-Driven Tests

Learn about table-driven tests, the idiomatic Go pattern for writing comprehensive tests efficiently.

## Learning Objectives

- Understand what table-driven tests are
- Learn when to use table-driven tests
- Master the table-driven test pattern
- Understand how to structure test cases
- Learn about subtests with `t.Run()`
- Understand best practices for table-driven tests

## Concepts Covered

### What are Table-Driven Tests?

Table-driven tests use a slice of test cases to test multiple scenarios with a single test function:

```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"test case 1", 1, 2},
        {"test case 2", 2, 4},
        {"test case 3", 3, 6},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Function(tt.input)
            if result != tt.expected {
                t.Errorf("Function(%d) = %d; expected %d", tt.input, result, tt.expected)
            }
        })
    }
}
```

### Basic Structure

**Test case struct:**
```go
tests := []struct {
    name     string  // Test case name
    input    Type    // Input value
    expected Type    // Expected output
    // Can add more fields as needed
}{
    // Test cases
}
```

### Using Subtests

Use `t.Run()` to create subtests for each test case:

```go
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // Test code
    })
}
```

**Benefits:**
- Each test case runs as separate subtest
- Can run specific test case: `go test -run TestFunction/TestCaseName`
- Better error messages (shows which case failed)
- Can run tests in parallel

### Common Patterns

**Basic table-driven test:**
```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"case 1", 1, 2},
        {"case 2", 2, 4},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Function(tt.input)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

**With error cases:**
```go
tests := []struct {
    name     string
    input    string
    expected bool
    wantErr  bool
}{
    {"valid input", "test", true, false},
    {"invalid input", "", false, true},
}
```

**With multiple inputs:**
```go
tests := []struct {
    name     string
    input1   int
    input2   int
    expected int
}{
    {"add", 1, 2, 3},
    {"subtract", 5, 3, 2},
}
```

## Running the Example

```bash
# Navigate to this directory
cd 04-Table-Driven-Tests

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/04-error-handling-testing/table-driven-tests

# Run all tests
go test

# Run with verbose output (shows subtests)
go test -v

# Run specific subtest
go test -run TestIsPalindrome/palindrome_word
```

## Running Tests

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run specific test
go test -run TestIsPalindrome

# Run specific subtest
go test -run TestIsPalindrome/palindrome_word
```

## Key Takeaways

1. **Table-driven tests are idiomatic** - Preferred pattern in Go
2. **Use struct for test cases** - Clear and extensible
3. **Use t.Run() for subtests** - Better organization and error messages
4. **Test multiple scenarios** - Easy to add more test cases
5. **Clear test names** - Help identify which case failed
6. **DRY principle** - Don't repeat test code

## Benefits

- **Less code duplication** - Single test function for multiple cases
- **Easy to add cases** - Just add to the slice
- **Clear structure** - All test cases in one place
- **Better error messages** - Shows which case failed
- **Selective running** - Can run specific subtests
- **Parallel execution** - Subtests can run in parallel

## Best Practices

- **Use descriptive names** - Clear test case names
- **Group related tests** - Similar cases together
- **Test edge cases** - Zero, negative, empty, nil
- **Test error cases** - Invalid inputs, error conditions
- **Use subtests** - `t.Run()` for better organization
- **Keep it simple** - Don't overcomplicate the structure

## Common Patterns

**Basic test:**
```go
tests := []struct {
    input    int
    expected int
}{
    {1, 2},
    {2, 4},
}
```

**With names:**
```go
tests := []struct {
    name     string
    input    int
    expected int
}{
    {"case 1", 1, 2},
    {"case 2", 2, 4},
}
```

**With error checking:**
```go
tests := []struct {
    name     string
    input    string
    expected string
    wantErr  bool
}{
    {"valid", "test", "result", false},
    {"invalid", "", "", true},
}
```

## When to Use

✅ **Use table-driven tests when:**
- Testing multiple similar scenarios
- Testing edge cases
- Testing with different inputs
- Want to avoid code duplication

❌ **Don't use when:**
- Test cases are very different
- Setup is complex for each case
- Tests need different logic

## Important Notes

- **Subtests run sequentially** - By default (can use `t.Parallel()`)
- **Test names must be unique** - Within the same test function
- **Can nest subtests** - But keep it simple
- **Use helpers** - For common test logic

## Next Steps

Congratulations! You've completed Section 4: Error Handling & Testing.

Proceed to:
- [Section 5: Concurrency](../../05-Concurrency/) - Learn about goroutines, channels, and concurrency

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

