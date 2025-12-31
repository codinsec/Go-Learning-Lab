# Sync Package

Learn about the sync package primitives for synchronization: WaitGroup, Mutex, RWMutex, Once, Cond, and Pool.

## Learning Objectives

- Understand WaitGroup for goroutine synchronization
- Learn about Mutex for mutual exclusion
- Master RWMutex for read-write locks
- Understand Once for one-time initialization
- Learn about Cond for condition variables
- Understand Pool for object reuse

## Concepts Covered

### WaitGroup

WaitGroup waits for a collection of goroutines to finish:

```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // Do work
    }()
}

wg.Wait() // Wait for all goroutines
```

### Mutex

Mutex provides mutual exclusion (lock/unlock):

```go
var mu sync.Mutex
var counter int

mu.Lock()
counter++
mu.Unlock()
```

### RWMutex

RWMutex allows multiple readers or one writer:

```go
var rwmu sync.RWMutex

// Readers
rwmu.RLock()
// Read data
rwmu.RUnlock()

// Writer
rwmu.Lock()
// Write data
rwmu.Unlock()
```

### Once

Once ensures a function executes exactly once:

```go
var once sync.Once

once.Do(func() {
    // Initialize
})
```

### Cond

Cond provides condition variables for signaling:

```go
var mu sync.Mutex
cond := sync.NewCond(&mu)

// Wait for condition
mu.Lock()
for !condition {
    cond.Wait()
}
mu.Unlock()

// Signal condition
mu.Lock()
condition = true
cond.Signal() // or cond.Broadcast()
mu.Unlock()
```

### Pool

Pool provides object pooling for reuse:

```go
var pool sync.Pool

pool.New = func() interface{} {
    return make([]byte, 1024)
}

buf := pool.Get().([]byte)
// Use buf
pool.Put(buf) // Return to pool
```

## Running the Example

```bash
# Navigate to this directory
cd 04-Sync-Package

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/05-concurrency/sync-package

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

1. **WaitGroup** - Wait for goroutines to finish
2. **Mutex** - Exclusive lock for mutual exclusion
3. **RWMutex** - Multiple readers or one writer
4. **Once** - Execute function exactly once
5. **Cond** - Condition variables for signaling
6. **Pool** - Object pooling to reduce allocations

## Common Patterns

**WaitGroup pattern:**
```go
var wg sync.WaitGroup
wg.Add(n)
go func() {
    defer wg.Done()
    // Work
}()
wg.Wait()
```

**Mutex pattern:**
```go
mu.Lock()
defer mu.Unlock()
// Critical section
```

**Once pattern:**
```go
var once sync.Once
once.Do(initFunction)
```

## Best Practices

- **Use defer** - Always unlock mutex with defer
- **Keep locks short** - Minimize time in critical section
- **Use RWMutex** - When reads > writes
- **Use Once** - For one-time initialization
- **Don't copy** - Sync primitives should not be copied

## Important Notes

- **Don't copy sync primitives** - Pass by pointer
- **Always unlock** - Use defer to ensure unlock
- **Deadlock risk** - Be careful with lock ordering
- **Pool objects** - May be nil, check before use

## Next Steps

After understanding sync package, proceed to:
- [05-Context](../05-Context/README.md) - Learn about context for cancellation and timeouts

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

