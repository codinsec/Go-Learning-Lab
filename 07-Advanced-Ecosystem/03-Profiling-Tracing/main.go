package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// This program demonstrates profiling and tracing in Go

// CPU-intensive function for profiling
func cpuIntensiveTask(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i * i
	}
	return sum
}

// Memory-intensive function
func memoryIntensiveTask(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(1000)
	}
	return result
}

// Function with goroutines for tracing
func concurrentTask() {
	ch := make(chan int, 10)
	
	// Spawn multiple goroutines
	for i := 0; i < 5; i++ {
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			ch <- id
		}(i)
	}
	
	// Collect results
	for i := 0; i < 5; i++ {
		<-ch
	}
}

func main() {
	fmt.Println("=== Profiling & Tracing ===")
	fmt.Println()

	// 1. CPU Profiling
	fmt.Println("1. CPU Profiling:")
	fmt.Println("   To profile CPU usage, run:")
	fmt.Println("   go tool pprof http://localhost:6060/debug/pprof/profile")
	fmt.Println("   Or use:")
	fmt.Println("   go test -cpuprofile=cpu.prof -bench=.")
	fmt.Println()
	
	// Simulate CPU work
	start := time.Now()
	result := cpuIntensiveTask(1000000)
	duration := time.Since(start)
	fmt.Printf("   CPU task completed: result=%d, duration=%v\n", result, duration)
	fmt.Println()

	// 2. Memory Profiling
	fmt.Println("2. Memory Profiling:")
	fmt.Println("   To profile memory usage, run:")
	fmt.Println("   go tool pprof http://localhost:6060/debug/pprof/heap")
	fmt.Println("   Or use:")
	fmt.Println("   go test -memprofile=mem.prof -bench=.")
	fmt.Println()
	
	// Simulate memory work
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	
	data := memoryIntensiveTask(100000)
	
	runtime.ReadMemStats(&m2)
	fmt.Printf("   Memory allocated: %d KB\n", (m2.Alloc-m1.Alloc)/1024)
	fmt.Printf("   Data length: %d\n", len(data))
	fmt.Println()

	// 3. Block Profiling
	fmt.Println("3. Block Profiling:")
	fmt.Println("   To profile blocking operations, run:")
	fmt.Println("   go tool pprof http://localhost:6060/debug/pprof/block")
	fmt.Println("   Enable with: runtime.SetBlockProfileRate(1)")
	fmt.Println()
	
	runtime.SetBlockProfileRate(1)
	concurrentTask()
	fmt.Println("   Block profiling enabled and concurrent task completed")
	fmt.Println()

	// 4. Goroutine Profiling
	fmt.Println("4. Goroutine Profiling:")
	fmt.Println("   To profile goroutines, run:")
	fmt.Println("   go tool pprof http://localhost:6060/debug/pprof/goroutine")
	fmt.Println()
	
	// Check current goroutines
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("   Current goroutines: %d\n", numGoroutines)
	fmt.Println()

	// 5. Trace Profiling
	fmt.Println("5. Trace Profiling:")
	fmt.Println("   To create a trace, run:")
	fmt.Println("   go test -trace=trace.out -bench=.")
	fmt.Println("   Then view with:")
	fmt.Println("   go tool trace trace.out")
	fmt.Println()
	
	// Simulate traceable work
	concurrentTask()
	fmt.Println("   Trace data can be collected for concurrent operations")
	fmt.Println()

	// 6. Runtime Statistics
	fmt.Println("6. Runtime Statistics:")
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	
	fmt.Printf("   Alloc: %d KB\n", stats.Alloc/1024)
	fmt.Printf("   TotalAlloc: %d KB\n", stats.TotalAlloc/1024)
	fmt.Printf("   Sys: %d KB\n", stats.Sys/1024)
	fmt.Printf("   NumGC: %d\n", stats.NumGC)
	fmt.Printf("   NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Println()

	// 7. Profiling HTTP Server Example
	fmt.Println("7. Profiling HTTP Server:")
	fmt.Println("   Add this to your HTTP server:")
	fmt.Println("   import _ \"net/http/pprof\"")
	fmt.Println("   go func() {")
	fmt.Println("       log.Println(http.ListenAndServe(\"localhost:6060\", nil))")
	fmt.Println("   }()")
	fmt.Println("   Then access:")
	fmt.Println("   http://localhost:6060/debug/pprof/")
	fmt.Println()

	// 8. Common pprof Commands
	fmt.Println("8. Common pprof Commands:")
	fmt.Println("   top        - Show top functions by CPU/memory")
	fmt.Println("   list func  - Show source code for function")
	fmt.Println("   web        - Generate SVG graph (requires graphviz)")
	fmt.Println("   png        - Generate PNG graph")
	fmt.Println("   peek func  - Show callers and callees")
	fmt.Println()

	// 9. Benchmark Profiling
	fmt.Println("9. Benchmark Profiling:")
	fmt.Println("   Run benchmarks with profiling:")
	fmt.Println("   go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof")
	fmt.Println("   Then analyze:")
	fmt.Println("   go tool pprof cpu.prof")
	fmt.Println("   go tool pprof mem.prof")
	fmt.Println()

	// 10. Key Concepts
	fmt.Println("10. Key Concepts:")
	fmt.Println("   - CPU profiling: Find performance bottlenecks")
	fmt.Println("   - Memory profiling: Find memory leaks")
	fmt.Println("   - Block profiling: Find blocking operations")
	fmt.Println("   - Goroutine profiling: Find goroutine leaks")
	fmt.Println("   - Trace: Visualize program execution")
	fmt.Println("   - Use pprof for analysis")
	fmt.Println("   - Profile in production-like conditions")
}

