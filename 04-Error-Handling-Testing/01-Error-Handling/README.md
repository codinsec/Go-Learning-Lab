# Error Handling

Learn about error handling in Go: the idiomatic `if err != nil` pattern, custom error types, and error wrapping.

## Learning Objectives

- Understand Go's error handling philosophy
- Master the `if err != nil` pattern
- Learn to create custom error types
- Understand error wrapping with `fmt.Errorf` and `%w`
- Learn to use `errors.Is` and `errors.As`
- Understand error unwrapping

## Concepts Covered

### Error as a Value

In Go, errors are values, not exceptions. Functions return errors as the last return value:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### The Idiomatic Pattern

The standard Go pattern for error handling:

```go
result, err := divide(10, 2)
if err != nil {
    // Handle error
    return err
}
// Use result
```

### Custom Error Types

Create custom error types by implementing the `error` interface:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}
```

### Error Wrapping

Wrap errors to add context using `fmt.Errorf` with `%w`:

```go
if err != nil {
    return fmt.Errorf("findUser: %w", err)
}
```

### errors.Is

Check if an error is a specific sentinel error:

```go
if errors.Is(err, os.ErrNotExist) {
    // Handle file not found
}
```

### errors.As

Check if an error is of a specific type and extract it:

```go
var validationErr ValidationError
if errors.As(err, &validationErr) {
    // Use validationErr.Field, validationErr.Message
}
```

### Error Unwrapping

Unwrap errors to access the underlying error:

```go
unwrapped := errors.Unwrap(err)
```

## Running the Example

```bash
# Navigate to this directory
cd 01-Error-Handling

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/04-error-handling-testing/error-handling

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

1. **Errors are values** - Not exceptions, returned as last value
2. **Always check errors** - Never ignore them
3. **if err != nil** - The idiomatic Go pattern
4. **Wrap errors with context** - Use `fmt.Errorf` with `%w`
5. **Use errors.Is** - For sentinel error comparison
6. **Use errors.As** - For custom error type checking
7. **Don't panic** - Return errors instead (unless unrecoverable)

## Common Patterns

**Basic error handling:**
```go
result, err := someFunction()
if err != nil {
    return err
}
```

**Error wrapping:**
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

**Custom error checking:**
```go
var customErr CustomError
if errors.As(err, &customErr) {
    // Handle custom error
}
```

**Multiple error checks:**
```go
if err1 != nil {
    return err1
}
if err2 != nil {
    return err2
}
```

## Best Practices

- **Always check errors** - Don't ignore them
- **Return errors early** - Fail fast
- **Wrap errors with context** - Add helpful information
- **Use custom error types** - For structured error information
- **Don't panic** - Use errors for expected failures
- **Provide helpful messages** - Make errors actionable

## Error Handling Philosophy

Go's error handling is explicit:
- Errors are visible in function signatures
- No hidden exceptions
- Forces developers to handle errors
- Makes error handling part of the code flow

## Important Notes

- **nil means success** - If error is nil, operation succeeded
- **Error is last return value** - Standard Go convention
- **Don't ignore errors** - Use `_` only if you're certain
- **Wrap, don't replace** - Use `%w` to preserve error chain

## Next Steps

After understanding error handling, proceed to:
- [02-Unit-Testing](../02-Unit-Testing/README.md) - Learn about unit testing in Go

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

