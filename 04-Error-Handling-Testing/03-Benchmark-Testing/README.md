# Benchmark Testing

Learn about benchmark testing in Go to measure and compare performance of your code.

## Learning Objectives

- Understand what benchmarks are and when to use them
- Learn how to write benchmark functions
- Master the `go test -bench` command
- Understand benchmark results
- Learn about sub-benchmarks
- Understand memory profiling with benchmarks

## Concepts Covered

### Writing Benchmarks

Benchmark functions start with `Benchmark` and take `*testing.B` as parameter:

```go
func BenchmarkFunctionName(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Code to benchmark
    }
}
```

### Running Benchmarks

**Run all benchmarks:**
```bash
go test -bench=.
```

**Run specific benchmark:**
```bash
go test -bench=BenchmarkFunctionName
```

**Run with memory profiling:**
```bash
go test -bench=. -benchmem
```

**Run with CPU profiling:**
```bash
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Benchmark Results

Benchmark output shows:
- **ns/op** - Nanoseconds per operation
- **B/op** - Bytes allocated per operation
- **allocs/op** - Number of allocations per operation

Example output:
```
BenchmarkConcatenateStrings-8    1000000    1200 ns/op    512 B/op    4 allocs/op
```

### Sub-benchmarks

Use `b.Run()` to create sub-benchmarks:

```go
func BenchmarkComparison(b *testing.B) {
    b.Run("Method1", func(b *testing.B) {
        // Benchmark method 1
    })
    b.Run("Method2", func(b *testing.B) {
        // Benchmark method 2
    })
}
```

### Benchmark Setup

Use `b.ResetTimer()` to exclude setup time:

```go
func BenchmarkFunction(b *testing.B) {
    // Setup code (excluded from timing)
    data := prepareData()
    b.ResetTimer() // Start timing from here
    for i := 0; i < b.N; i++ {
        // Code to benchmark
    }
}
```

## Running the Example

```bash
# Navigate to this directory
cd 03-Benchmark-Testing

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/04-error-handling-testing/benchmark-testing

# Run all benchmarks
go test -bench=.

# Run with memory profiling
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkConcatenateStrings
```

## Running Tests

```bash
# Run all benchmarks
go test -bench=.

# Run with memory profiling
go test -bench=. -benchmem

# Run with CPU profiling
go test -bench=. -cpuprofile=cpu.prof

# Compare benchmarks (requires benchcmp tool)
go get golang.org/x/tools/cmd/benchcmp
go test -bench=. > old.txt
# Make changes
go test -bench=. > new.txt
benchcmp old.txt new.txt
```

## Key Takeaways

1. **Benchmarks measure performance** - Use to find bottlenecks
2. **b.N is iteration count** - Go adjusts automatically
3. **Use b.ResetTimer()** - Exclude setup from timing
4. **Use -benchmem** - See memory allocations
5. **Sub-benchmarks** - Compare multiple implementations
6. **Don't optimize prematurely** - Measure first, optimize second

## Common Patterns

**Basic benchmark:**
```go
func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Function()
    }
}
```

**Benchmark with setup:**
```go
func BenchmarkFunction(b *testing.B) {
    data := prepareData()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Function(data)
    }
}
```

**Comparing implementations:**
```go
func BenchmarkComparison(b *testing.B) {
    b.Run("Method1", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Method1()
        }
    })
    b.Run("Method2", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Method2()
        }
    })
}
```

## Best Practices

- **Benchmark real scenarios** - Use realistic data sizes
- **Use b.ResetTimer()** - Exclude setup time
- **Run multiple times** - Use `-count` flag for stability
- **Compare fairly** - Same input data, same conditions
- **Watch for compiler optimizations** - Results can be misleading
- **Use -benchmem** - Understand memory usage

## Benchmark Flags

- **-bench=.** - Run all benchmarks
- **-benchmem** - Show memory allocations
- **-count=N** - Run benchmarks N times
- **-timeout=D** - Set timeout duration
- **-cpuprofile=FILE** - Write CPU profile
- **-memprofile=FILE** - Write memory profile

## Important Notes

- **b.N is adjusted automatically** - Go finds stable timing
- **Compiler optimizations** - Can affect results (use `-gcflags=-N` to disable)
- **Warm-up** - First few iterations may be slower
- **Parallel benchmarks** - Use `b.RunParallel()` for concurrent benchmarks

## When to Use Benchmarks

- **Performance critical code** - Identify bottlenecks
- **Comparing algorithms** - Which is faster?
- **Optimization verification** - Did optimization help?
- **Regression testing** - Performance didn't degrade

## Next Steps

After understanding benchmark testing, proceed to:
- [04-Table-Driven-Tests](../04-Table-Driven-Tests/README.md) - Learn about table-driven tests

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

