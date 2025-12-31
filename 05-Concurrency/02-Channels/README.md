# Channels

Learn about channels, Go's primary mechanism for goroutine communication and synchronization.

## Learning Objectives

- Understand what channels are and how they work
- Learn the difference between buffered and unbuffered channels
- Master sending and receiving values
- Understand channel directions
- Learn about closing channels
- Understand channel blocking behavior

## Concepts Covered

### What are Channels?

Channels are typed conduits for sending and receiving values between goroutines. They provide a safe way to communicate.

```go
ch := make(chan int)        // Unbuffered channel
ch := make(chan int, 10)    // Buffered channel (capacity 10)
```

### Unbuffered Channels

Unbuffered channels are **synchronous**:
- Sender blocks until receiver is ready
- Receiver blocks until sender is ready
- Capacity: 0
- Use for synchronization

```go
ch := make(chan int)
go func() {
    ch <- 42  // Blocks until received
}()
value := <-ch  // Blocks until sent
```

### Buffered Channels

Buffered channels are **asynchronous**:
- Sender only blocks if buffer is full
- Receiver only blocks if buffer is empty
- Capacity: specified
- Use for queuing

```go
ch := make(chan int, 3)
ch <- 1  // Doesn't block (buffer has space)
ch <- 2
ch <- 3
ch <- 4  // Blocks (buffer is full)
```

### Sending and Receiving

**Send:**
```go
ch <- value
```

**Receive:**
```go
value := <-ch
value, ok := <-ch  // ok is false if channel is closed
```

### Closing Channels

Close a channel when no more values will be sent:

```go
close(ch)
```

**Important:**
- Only sender should close
- Closing closed channel causes panic
- Receiving from closed channel returns zero value and false

### Range Over Channel

Range automatically receives until channel is closed:

```go
for value := range ch {
    // Process value
}
```

### Channel Directions

Specify channel direction in function signatures:

```go
func send(ch chan<- int) {  // Send-only
    ch <- 42
}

func receive(ch <-chan int) {  // Receive-only
    value := <-ch
}
```

## Running the Example

```bash
# Navigate to this directory
cd 02-Channels

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/05-concurrency/channels

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

1. **Channels enable safe communication** - Between goroutines
2. **Unbuffered = synchronous** - Sender and receiver must be ready
3. **Buffered = asynchronous** - Can queue values
4. **Close channels** - When done sending (only sender should close)
5. **Range over channel** - Automatically receives until closed
6. **Channel directions** - Enforce send-only or receive-only

## Common Patterns

**Basic communication:**
```go
ch := make(chan int)
go func() {
    ch <- compute()
}()
result := <-ch
```

**Multiple values:**
```go
ch := make(chan int)
go func() {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}()
for val := range ch {
    // Process val
}
```

**Return value from goroutine:**
```go
resultCh := make(chan int)
go func() {
    resultCh <- compute()
}()
result := <-resultCh
```

## Best Practices

- **Close channels** - When done sending
- **Only sender closes** - Never receiver
- **Check closed status** - Use `value, ok := <-ch`
- **Use buffered channels** - When you need queuing
- **Use unbuffered channels** - For synchronization
- **Don't send on closed channel** - Causes panic

## Channel Operations

| Operation | Unbuffered | Buffered (not full) | Buffered (full) |
|-----------|------------|---------------------|-----------------|
| Send | Blocks until received | Non-blocking | Blocks |
| Receive | Blocks until sent | Non-blocking | Blocks if empty |

## Important Notes

- **Nil channels** - Operations on nil channel block forever
- **Closed channel** - Sending causes panic, receiving returns zero value
- **Channel capacity** - Use `cap(ch)` to get capacity
- **Channel length** - Use `len(ch)` to get number of queued values

## Next Steps

After understanding channels, proceed to:
- [03-Select](../03-Select/README.md) - Learn about select statement for multiple channels

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

