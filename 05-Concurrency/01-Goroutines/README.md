# Goroutines

Learn about goroutines, Go's lightweight threads that make concurrency simple and efficient.

## Learning Objectives

- Understand what goroutines are and how they work
- Learn how to start goroutines with the `go` keyword
- Understand goroutine lifecycle
- Learn about goroutine synchronization
- Understand common patterns and pitfalls

## Concepts Covered

### What are Goroutines?

Goroutines are lightweight threads managed by the Go runtime. They're much cheaper than OS threads:

- **Stack size**: ~2KB (grows as needed)
- **OS threads**: ~1-2MB
- **Managed by Go runtime**: Not OS threads
- **M:N threading**: Many goroutines on few OS threads

### Starting a Goroutine

Use the `go` keyword to start a goroutine:

```go
go functionName()
go func() {
    // Anonymous function
}()
```

### Goroutine Lifecycle

1. **Creation**: `go` keyword starts goroutine
2. **Execution**: Runs concurrently with other goroutines
3. **Completion**: Finishes when function returns
4. **Main goroutine**: Program exits when main goroutine finishes

### Synchronization

**Problem**: Main goroutine may exit before other goroutines finish.

**Solutions**:
- `sync.WaitGroup` (covered in Sync Package topic)
- Channels (covered in Channels topic)
- `time.Sleep()` (not recommended for production)

## Running the Example

```bash
# Navigate to this directory
cd 01-Goroutines

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/05-concurrency/goroutines

# Run the program
go run main.go
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run benchmarks
go test -bench=.
```

## Key Takeaways

1. **Goroutines are lightweight** - Can create thousands easily
2. **Use `go` keyword** - Starts goroutine immediately
3. **Main goroutine must wait** - Or program exits early
4. **M:N threading model** - Efficient scheduling
5. **No return values directly** - Use channels instead
6. **Be careful with closures** - Capture loop variables correctly

## Common Patterns

**Basic goroutine:**
```go
go func() {
    // Do work
}()
```

**With parameters:**
```go
for i := 0; i < 10; i++ {
    go func(id int) {
        // Use id
    }(i) // Pass i as parameter
}
```

**With WaitGroup:**
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()
wg.Wait()
```

## Best Practices

- **Don't create unlimited goroutines** - Use worker pools
- **Always wait for goroutines** - Use WaitGroup or channels
- **Pass parameters explicitly** - Don't rely on closure capture
- **Handle errors** - Goroutines can fail silently
- **Use channels for communication** - Don't share memory

## Common Mistakes

**1. Not waiting for goroutines:**
```go
go doWork() // Main exits before goroutine finishes
```

**2. Capturing loop variable:**
```go
for i := 0; i < 10; i++ {
    go func() {
        fmt.Println(i) // All print same value!
    }()
}
// Fix: Pass as parameter
go func(id int) {
    fmt.Println(id)
}(i)
```

**3. Ignoring errors:**
```go
go func() {
    err := doWork()
    // Error ignored!
}()
```

## Goroutine vs Thread

| Feature | Goroutine | OS Thread |
|---------|-----------|-----------|
| Stack | ~2KB | ~1-2MB |
| Creation | Fast | Slow |
| Scheduling | Go runtime | OS kernel |
| Communication | Channels | Shared memory |

## Important Notes

- **Goroutines are not OS threads** - Managed by Go runtime
- **GOMAXPROCS** - Controls number of OS threads (default: CPU cores)
- **Goroutine scheduler** - Automatically manages goroutines
- **No guarantee of execution order** - Goroutines run concurrently

## Next Steps

After understanding goroutines, proceed to:
- [02-Channels](../02-Channels/README.md) - Learn about channels for goroutine communication

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

