package main

import (
	"fmt"
	"strings"
	"testing"
)

// BenchmarkConcatenateStrings benchmarks string concatenation with +
func BenchmarkConcatenateStrings(b *testing.B) {
	strs := []string{"hello", "world", "test", "benchmark"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatenateStrings(strs)
	}
}

// BenchmarkConcatenateStringsBuilder benchmarks string concatenation with Builder
func BenchmarkConcatenateStringsBuilder(b *testing.B) {
	strs := []string{"hello", "world", "test", "benchmark"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatenateStringsBuilder(strs)
	}
}

// BenchmarkSumArray benchmarks array summation
func BenchmarkSumArray(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumArray(arr)
	}
}

// BenchmarkFindMax benchmarks finding maximum
func BenchmarkFindMax(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMax(arr)
	}
}

// BenchmarkBinarySearch benchmarks binary search
func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	target := 500
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, target)
	}
}

// Sub-benchmark example
func BenchmarkConcatenate(b *testing.B) {
	b.Run("Plus", func(b *testing.B) {
		strs := []string{"hello", "world"}
		for i := 0; i < b.N; i++ {
			ConcatenateStrings(strs)
		}
	})

	b.Run("Builder", func(b *testing.B) {
		strs := []string{"hello", "world"}
		for i := 0; i < b.N; i++ {
			ConcatenateStringsBuilder(strs)
		}
	})
}

// Benchmark with different input sizes
func BenchmarkConcatenateStringsSize(b *testing.B) {
	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			strs := make([]string, size)
			for i := range strs {
				strs[i] = "test"
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ConcatenateStringsBuilder(strs)
			}
		})
	}
}

// Helper function to create test data
func createTestStrings(n int) []string {
	strs := make([]string, n)
	for i := range strs {
		strs[i] = "test"
	}
	return strs
}

// Example of comparing two implementations
func BenchmarkStringComparison(b *testing.B) {
	testData := createTestStrings(100)

	b.Run("Plus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ConcatenateStrings(testData)
		}
	})

	b.Run("Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ConcatenateStringsBuilder(testData)
		}
	})
}

