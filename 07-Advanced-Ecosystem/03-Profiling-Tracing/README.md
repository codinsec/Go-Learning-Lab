# Profiling & Tracing

Learn about profiling and tracing in Go to identify performance bottlenecks and optimize your applications.

## Learning Objectives

- Understand different types of profiling (CPU, memory, block, goroutine)
- Learn to use `pprof` for performance analysis
- Master trace profiling for concurrent programs
- Understand runtime statistics
- Learn to set up profiling in HTTP servers
- Understand benchmark profiling

## Concepts Covered

### CPU Profiling

Profile CPU usage to find bottlenecks:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

### Memory Profiling

Profile memory usage to find leaks:

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

### Block Profiling

Profile blocking operations:

```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

### Goroutine Profiling

Profile goroutines:

```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

### Trace Profiling

Create execution traces:

```bash
go test -trace=trace.out -bench=.
go tool trace trace.out
```

### HTTP Server Profiling

Add pprof to your HTTP server:

```go
import _ "net/http/pprof"

go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

## Running the Example

```bash
# Navigate to this directory
cd 03-Profiling-Tracing

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/profiling-tracing

# Run the program
go run main.go
```

## Running Tests and Benchmarks

```bash
# Run all tests
go test

# Run benchmarks
go test -bench=.

# Run with CPU profiling
go test -cpuprofile=cpu.prof -bench=.

# Run with memory profiling
go test -memprofile=mem.prof -bench=.

# Run with trace
go test -trace=trace.out -bench=.

# Analyze profiles
go tool pprof cpu.prof
go tool pprof mem.prof
go tool trace trace.out
```

## Key Takeaways

1. **CPU profiling** - Find performance bottlenecks
2. **Memory profiling** - Find memory leaks
3. **Block profiling** - Find blocking operations
4. **Goroutine profiling** - Find goroutine leaks
5. **Trace profiling** - Visualize execution
6. **pprof tool** - Analyze profiles interactively
7. **HTTP integration** - Easy profiling in web apps

## Common pprof Commands

- `top` - Show top functions by CPU/memory
- `list func` - Show source code for function
- `web` - Generate SVG graph (requires graphviz)
- `png` - Generate PNG graph
- `peek func` - Show callers and callees
- `help` - Show all commands

## Best Practices

- **Profile in production-like conditions** - Use realistic data
- **Profile multiple times** - Results can vary
- **Focus on hot paths** - Optimize what matters
- **Use benchmarks** - For consistent profiling
- **Compare before/after** - Measure improvements
- **Don't over-optimize** - Profile first, optimize second

## Important Notes

- **Performance overhead** - Profiling adds overhead
- **Production use** - Can be used in production carefully
- **Graphviz** - Required for `web` command
- **Trace viewer** - Opens in browser
- **Memory stats** - Use `runtime.ReadMemStats()`

## Profiling Workflow

1. **Identify problem** - Slow performance, high memory
2. **Enable profiling** - Add pprof or use flags
3. **Run workload** - Exercise the code
4. **Collect profile** - Save profile data
5. **Analyze** - Use pprof to find issues
6. **Optimize** - Fix identified problems
7. **Verify** - Profile again to confirm improvement

## Example: HTTP Server with Profiling

```go
package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
)

func main() {
    // Start pprof server
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Your HTTP server
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Then access:
- `http://localhost:6060/debug/pprof/` - Index
- `http://localhost:6060/debug/pprof/profile` - CPU profile
- `http://localhost:6060/debug/pprof/heap` - Memory profile

## Next Steps

After understanding profiling, proceed to:
- [04-Build-Tags](../04-Build-Tags/README.md) - Learn about build tags

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

