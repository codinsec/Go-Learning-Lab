# Select Statement

Learn about the select statement for handling multiple channel operations in Go.

## Learning Objectives

- Understand what the select statement does
- Learn how to use select with multiple channels
- Master the default case for non-blocking operations
- Understand select with timeouts
- Learn about select in loops
- Understand select for sending and receiving

## Concepts Covered

### What is Select?

The `select` statement lets a goroutine wait on multiple channel operations. It's like a `switch` but for channels.

```go
select {
case msg1 := <-ch1:
    // Handle message from ch1
case msg2 := <-ch2:
    // Handle message from ch2
}
```

### Basic Select

Select blocks until one of its cases can proceed:

```go
select {
case val := <-ch1:
    fmt.Println("Received from ch1")
case val := <-ch2:
    fmt.Println("Received from ch2")
}
```

### Default Case

The `default` case makes select non-blocking:

```go
select {
case val := <-ch:
    fmt.Println("Received:", val)
default:
    fmt.Println("No value ready")
}
```

### Select with Timeout

Use `time.After` for timeouts:

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
}
```

### Select in Loop

Common pattern for processing multiple channels:

```go
for {
    select {
    case val := <-ch1:
        // Process val
    case val := <-ch2:
        // Process val
    case <-done:
        return
    }
}
```

### Select for Sending

Select can also handle sending:

```go
select {
case ch <- value:
    fmt.Println("Sent")
default:
    fmt.Println("Channel full")
}
```

## Running the Example

```bash
# Navigate to this directory
cd 03-Select

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/05-concurrency/select

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

1. **Select waits on multiple channels** - Like switch for channels
2. **Random selection** - If multiple cases ready, one is chosen randomly
3. **Default case** - Makes select non-blocking
4. **Timeout pattern** - Use `time.After` for timeouts
5. **Select in loops** - Common pattern for processing
6. **Nil channels** - Never ready in select

## Common Patterns

**Non-blocking receive:**
```go
select {
case val := <-ch:
    // Process val
default:
    // No value ready
}
```

**Timeout:**
```go
select {
case result := <-ch:
    // Use result
case <-time.After(5 * time.Second):
    // Handle timeout
}
```

**Cancellation:**
```go
select {
case val := <-ch:
    // Process val
case <-cancel:
    return
}
```

**Multiple channels:**
```go
for {
    select {
    case val1 := <-ch1:
        // Handle ch1
    case val2 := <-ch2:
        // Handle ch2
    case <-done:
        return
    }
}
```

## Best Practices

- **Use default** - For non-blocking operations
- **Use timeouts** - Prevent indefinite blocking
- **Handle all cases** - Don't ignore channels
- **Use done channel** - For cancellation
- **Random selection** - Don't rely on order

## Important Notes

- **Random selection** - If multiple cases ready, selection is random
- **Nil channels** - Never ready, so never selected
- **Closed channels** - Always ready, return zero value
- **Default always ready** - Selected if no other case ready

## Next Steps

After understanding select, proceed to:
- [04-Sync-Package](../04-Sync-Package/README.md) - Learn about sync package primitives

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

