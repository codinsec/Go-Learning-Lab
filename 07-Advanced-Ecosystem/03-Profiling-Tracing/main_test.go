package main

import (
	"runtime"
	"testing"
)

func BenchmarkCPUIntensiveTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cpuIntensiveTask(10000)
	}
}

func BenchmarkMemoryIntensiveTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryIntensiveTask(1000)
	}
}

func BenchmarkConcurrentTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrentTask()
	}
}

func TestCPUIntensiveTask(t *testing.T) {
	result := cpuIntensiveTask(1000)
	if result < 0 {
		t.Error("Result should be non-negative")
	}
}

func TestMemoryIntensiveTask(t *testing.T) {
	data := memoryIntensiveTask(100)
	if len(data) != 100 {
		t.Errorf("Expected length 100, got %d", len(data))
	}
}

func TestConcurrentTask(t *testing.T) {
	// Test that concurrent task completes without deadlock
	concurrentTask()
	
	// Check that goroutines are cleaned up
	runtime.GC()
	numGoroutines := runtime.NumGoroutine()
	if numGoroutines > 10 { // Allow some overhead
		t.Errorf("Too many goroutines: %d", numGoroutines)
	}
}

func TestRuntimeStats(t *testing.T) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	
	if stats.Alloc == 0 && stats.TotalAlloc == 0 {
		t.Error("Memory stats should not be zero")
	}
}

