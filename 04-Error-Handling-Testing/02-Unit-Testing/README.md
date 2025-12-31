# Unit Testing

Learn about unit testing in Go using the `testing` package and `go test` command.

## Learning Objectives

- Understand how to write tests in Go
- Learn the `testing` package
- Master the `go test` command
- Understand test coverage
- Learn about test helpers
- Understand test organization

## Concepts Covered

### Writing Tests

Tests are functions that start with `Test` and take `*testing.T` as parameter:

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
    }
}
```

### Test File Naming

Test files must end with `_test.go`:
- `main.go` → `main_test.go`
- `calculator.go` → `calculator_test.go`

### Running Tests

**Run all tests:**
```bash
go test
```

**Run with verbose output:**
```bash
go test -v
```

**Run specific test:**
```bash
go test -run TestAdd
```

**Run with coverage:**
```bash
go test -cover
```

**Generate coverage report:**
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Functions

**Basic test:**
```go
func TestFunctionName(t *testing.T) {
    // Test code
}
```

**Helper function:**
```go
func helperFunction(t *testing.T) {
    t.Helper() // Marks as helper
    // Helper code
}
```

### Test Methods

**t.Error / t.Errorf:**
- Logs error and continues test
- Test fails but continues

**t.Fatal / t.Fatalf:**
- Logs error and stops test immediately
- Test fails and stops

**t.Skip / t.Skipf:**
- Skips the test
- Useful for conditional tests

**t.Helper():**
- Marks function as test helper
- Removes from stack trace

## Running the Example

```bash
# Navigate to this directory
cd 02-Unit-Testing

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/04-error-handling-testing/unit-testing

# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run specific test
go test -run TestAdd
```

## Running Tests

```bash
# Run all tests
go test

# Run with verbose output
go test -v

# Run with coverage
go test -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Key Takeaways

1. **Test files end with `_test.go`** - Required naming convention
2. **Test functions start with `Test`** - Required naming convention
3. **Use `t.Error` for failures** - Logs error, continues test
4. **Use `t.Fatal` for critical failures** - Stops test immediately
5. **Test coverage** - Use `-cover` flag to see coverage
6. **Test helpers** - Use `t.Helper()` for helper functions

## Common Patterns

**Basic assertion:**
```go
if got != want {
    t.Errorf("got %v, want %v", got, want)
}
```

**Error checking:**
```go
if err != nil {
    t.Fatalf("unexpected error: %v", err)
}
```

**Test helper:**
```go
func checkResult(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

## Best Practices

- **Test one thing per test** - Keep tests focused
- **Use descriptive names** - `TestAdd` not `Test1`
- **Test edge cases** - Zero, negative, empty, nil
- **Test error cases** - Invalid input, error conditions
- **Keep tests simple** - Easy to read and understand
- **Use table-driven tests** - For multiple test cases (see next topic)

## Test Organization

**Same package:**
- Tests in same package can access unexported functions
- Use for unit testing internal functions

**Separate package:**
- Tests in `_test` package (e.g., `package main_test`)
- Can only access exported functions
- Use for integration testing

## Coverage

Test coverage shows which code is tested:
- **100% coverage** - All code is tested (ideal but not always practical)
- **Aim for high coverage** - But focus on important code
- **Use coverage reports** - Identify untested code

## Important Notes

- **Tests run in parallel** - By default (can disable with `-parallel 1`)
- **Test order is random** - Don't rely on execution order
- **Clean up resources** - Use `defer` for cleanup
- **Don't ignore errors** - Test error cases too

## Next Steps

After understanding unit testing, proceed to:
- [03-Benchmark-Testing](../03-Benchmark-Testing/README.md) - Learn about benchmark tests

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

