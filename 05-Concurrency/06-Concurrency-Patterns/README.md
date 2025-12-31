# Concurrency Patterns

Learn about common concurrency patterns in Go: worker pools, fan-in/fan-out, pipelines, and more.

## Learning Objectives

- Understand worker pool pattern
- Learn fan-in and fan-out patterns
- Master pipeline pattern
- Understand generator pattern
- Learn semaphore pattern
- Understand or-channel pattern

## Concepts Covered

### Worker Pool

Fixed number of workers processing jobs from a queue:

```go
jobs := make(chan int, 100)
results := make(chan int, 100)

// Start workers
for w := 1; w <= 3; w++ {
    go worker(jobs, results)
}

// Send jobs
for j := 1; j <= 10; j++ {
    jobs <- j
}
close(jobs)

// Collect results
for r := 1; r <= 10; r++ {
    <-results
}
```

### Fan-In

Combine multiple channels into one:

```go
func fanIn(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for {
            select {
            case v := <-ch1:
                out <- v
            case v := <-ch2:
                out <- v
            }
        }
    }()
    return out
}
```

### Fan-Out

Distribute work to multiple workers:

```go
func fanOut(input <-chan int, outputs []chan int) {
    for val := range input {
        // Distribute to available worker
        select {
        case outputs[0] <- val:
        case outputs[1] <- val:
        }
    }
}
```

### Pipeline

Chain of stages processing data:

```go
// Stage 1: Generate
numbers := generate()

// Stage 2: Process
squares := square(numbers)

// Stage 3: Consume
for sq := range squares {
    fmt.Println(sq)
}
```

### Generator

Function that returns a channel:

```go
func generator() <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
    return ch
}
```

### Semaphore

Limit concurrent operations:

```go
sem := make(chan struct{}, 3) // Allow 3 concurrent

for i := 0; i < 10; i++ {
    go func() {
        sem <- struct{}{} // Acquire
        defer func() { <-sem }() // Release
        
        // Do work
    }()
}
```

### Or-Channel

Returns when any channel receives:

```go
func or(channels ...<-chan interface{}) <-chan interface{} {
    switch len(channels) {
    case 0:
        return nil
    case 1:
        return channels[0]
    }
    
    orDone := make(chan interface{})
    go func() {
        defer close(orDone)
        select {
        case <-channels[0]:
        case <-channels[1]:
        case <-or(append(channels[2:], orDone)...):
        }
    }()
    return orDone
}
```

## Running the Example

```bash
# Navigate to this directory
cd 06-Concurrency-Patterns

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/05-concurrency/concurrency-patterns

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

1. **Worker Pool** - Fixed workers processing jobs
2. **Fan-In** - Combine multiple channels
3. **Fan-Out** - Distribute to multiple workers
4. **Pipeline** - Chain of processing stages
5. **Generator** - Function returning channel
6. **Semaphore** - Limit concurrency

## Common Patterns

**Worker pool:**
```go
// Fixed workers, job queue, result queue
```

**Pipeline:**
```go
// Stage1 -> Stage2 -> Stage3
```

**Generator:**
```go
// Function returns channel
```

## Best Practices

- **Use worker pools** - Limit goroutine creation
- **Close channels** - When done sending
- **Use buffered channels** - For queuing
- **Handle errors** - In each stage
- **Use context** - For cancellation

## Important Notes

- **Worker pools** - Prevent too many goroutines
- **Fan-in/out** - For load distribution
- **Pipelines** - For data processing
- **Semaphores** - For rate limiting

## Next Steps

Congratulations! You've completed Section 5: Concurrency.

Proceed to:
- [Section 6: Standard Library & Web](../../06-Standard-Library-Web/) - Learn about Go's standard library and web development

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

