# Context

Learn about the context package for cancellation, timeouts, and request-scoped values in Go.

## Learning Objectives

- Understand what context is and when to use it
- Learn about context cancellation
- Master context timeouts and deadlines
- Understand context values
- Learn context propagation patterns
- Understand best practices

## Concepts Covered

### What is Context?

Context carries cancellation signals, deadlines, and request-scoped values across API boundaries.

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

### Background Context

The root context:

```go
ctx := context.Background()
```

### WithCancel

Creates a context that can be cancelled:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // Always call cancel

// In goroutine
select {
case <-ctx.Done():
    return ctx.Err()
}
```

### WithTimeout

Creates a context with timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case <-ctx.Done():
    return ctx.Err()
}
```

### WithDeadline

Creates a context with deadline:

```go
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### WithValue

Stores request-scoped values:

```go
ctx := context.WithValue(context.Background(), "userID", "12345")
userID := ctx.Value("userID")
```

## Running the Example

```bash
# Navigate to this directory
cd 05-Context

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/05-concurrency/context

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

1. **Context for cancellation** - Signal goroutines to stop
2. **Context for timeouts** - Set operation timeouts
3. **Context for values** - Request-scoped data
4. **Always pass context** - As first parameter
5. **Always call cancel** - Use defer cancel()
6. **Check ctx.Done()** - In long-running operations

## Common Patterns

**Function with context:**
```go
func doWork(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(1 * time.Second):
        return nil
    }
}
```

**Context in loop:**
```go
for {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        // Do work
    }
}
```

**Context propagation:**
```go
func processRequest(ctx context.Context) error {
    return doWork(ctx) // Pass context down
}
```

## Best Practices

- **First parameter** - Context should be first parameter
- **Don't store in structs** - Pass as parameter
- **Always call cancel** - Use defer cancel()
- **Use WithValue sparingly** - Only for request-scoped data
- **Check ctx.Done()** - In loops and long operations
- **Use typed keys** - For context values

## Important Notes

- **Context is immutable** - Each With* creates new context
- **Cancel releases resources** - Always call cancel
- **Done() channel** - Closed when context is done
- **Err()** - Returns error (Canceled or DeadlineExceeded)
- **Background()** - Root context, never cancelled

## Context Errors

- **context.Canceled** - Context was cancelled
- **context.DeadlineExceeded** - Deadline was exceeded
- **nil** - Context is not done

## Next Steps

After understanding context, proceed to:
- [06-Concurrency-Patterns](../06-Concurrency-Patterns/README.md) - Learn about common concurrency patterns

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

