package main

import (
	"fmt"
	"strings"
)

// Functions to benchmark

// ConcatenateStrings concatenates strings using + operator
func ConcatenateStrings(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

// ConcatenateStringsBuilder uses strings.Builder
func ConcatenateStringsBuilder(strs []string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// SumArray sums all elements in an array
func SumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// FindMax finds maximum value in array
func FindMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// BinarySearch performs binary search
func BinarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println("=== Benchmark Testing ===")
	fmt.Println()
	fmt.Println("Run benchmarks with: go test -bench=.")
	fmt.Println("Run specific benchmark: go test -bench=BenchmarkConcatenateStrings")
	fmt.Println("Run with memory profiling: go test -bench=. -benchmem")
	fmt.Println("Compare benchmarks: go test -bench=. -benchmem > old.txt")
	fmt.Println("                     go test -bench=. -benchmem > new.txt")
	fmt.Println("                     benchcmp old.txt new.txt")
}

