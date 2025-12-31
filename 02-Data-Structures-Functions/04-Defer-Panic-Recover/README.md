# Defer, Panic, Recover

Learn about Go's defer, panic, and recover mechanisms for flow control and error handling.

## Learning Objectives

- Understand the defer statement and its execution order
- Learn when and how to use defer for cleanup
- Understand panic and when it occurs
- Learn how to recover from panics
- Master the defer-panic-recover pattern

## Concepts Covered

### Defer

The `defer` statement schedules a function call to be executed after the surrounding function returns.

```go
func example() {
    defer fmt.Println("This runs last")
    fmt.Println("This runs first")
}
```

**Key characteristics:**
- **LIFO order** - Multiple defers execute in reverse order (Last In First Out)
- **Arguments evaluated immediately** - Function arguments are evaluated when defer is called
- **Always executes** - Even if function panics or returns early
- **Common use case** - Resource cleanup (files, connections, etc.)

### Panic

`panic` stops normal execution and begins panicking. When a function panics, it stops executing and any deferred functions run.

```go
panic("error message")
```

**When panic occurs:**
- Program crashes (unless recovered)
- Deferred functions still execute
- Panic propagates up the call stack

**Common causes:**
- Array/slice index out of bounds
- Nil pointer dereference
- Division by zero (in some cases)
- Explicit `panic()` call

### Recover

`recover` stops the panicking sequence and returns the value passed to panic. It only works inside deferred functions.

```go
defer func() {
    if r := recover(); r != nil {
        // Handle the panic
        fmt.Println("Recovered:", r)
    }
}()
```

**Key points:**
- Only works inside deferred functions
- Returns `nil` if no panic occurred
- Returns the panic value if panic occurred
- Allows program to continue after panic

## Running the Example

```bash
# Navigate to this directory
cd 04-Defer-Panic-Recover

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/02-data-structures-functions/defer-panic-recover

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

1. **Defer executes in LIFO order** - Last deferred function runs first
2. **Defer arguments are evaluated immediately** - Not when function executes
3. **Defer always runs** - Even on panic or early return
4. **Panic stops execution** - But deferred functions still run
5. **Recover only works in defer** - Must be inside a deferred function
6. **Use defer for cleanup** - Files, connections, mutex unlocks, etc.

## Common Patterns

**Resource cleanup:**
```go
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close() // Always closes, even on error
```

**Mutex unlocking:**
```go
mu.Lock()
defer mu.Unlock() // Always unlocks
```

**Error recovery:**
```go
defer func() {
    if r := recover(); r != nil {
        log.Printf("Recovered from panic: %v", r)
    }
}()
```

**Multiple defers:**
```go
defer fmt.Println("Third")
defer fmt.Println("Second")
defer fmt.Println("First")
// Output: First, Second, Third
```

## Best Practices

- **Use defer for cleanup** - Files, connections, mutexes
- **Don't overuse panic** - Prefer returning errors
- **Recover at appropriate levels** - Usually at the top level of goroutines
- **Be careful with defer arguments** - They're evaluated immediately
- **Use named returns with defer** - Can modify return values

## When to Use

**Defer:**
- File operations
- Database connections
- Mutex locking/unlocking
- Any cleanup that must happen

**Panic:**
- Programming errors (nil pointer, index out of bounds)
- Unrecoverable situations
- Rarely used in application code

**Recover:**
- Top-level goroutine handlers
- Converting panics to errors
- Graceful shutdown handling

## Important Notes

- **Panic is not for normal errors** - Use error returns instead
- **Recover doesn't work across goroutines** - Each goroutine needs its own recover
- **Defer can modify named return values** - Be careful with this
- **Defer has a small performance cost** - Don't use in tight loops

## Next Steps

After understanding defer, panic, and recover, proceed to:
- [05-Pointers](../05-Pointers/README.md) - Learn about pointers and memory addresses

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

