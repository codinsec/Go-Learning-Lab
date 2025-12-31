package main

import "testing"

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// TestSubtract tests the Subtract function
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
	}
}

// TestMultiply tests the Multiply function
func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	expected := 20
	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; expected %d", result, expected)
	}
}

// TestDivide tests the Divide function
func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Divide(10, 2) = %.2f; expected %.2f", result, expected)
	}
}

// TestDivideByZero tests division by zero
func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return an error")
	}
}

// TestIsEven tests the IsEven function
func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
		{-3, false},
	}

	for _, tt := range tests {
		result := IsEven(tt.input)
		if result != tt.expected {
			t.Errorf("IsEven(%d) = %t; expected %t", tt.input, result, tt.expected)
		}
	}
}

// TestMax tests the Max function
func TestMax(t *testing.T) {
	result := Max(10, 5)
	expected := 10
	if result != expected {
		t.Errorf("Max(10, 5) = %d; expected %d", result, expected)
	}

	result = Max(5, 10)
	if result != expected {
		t.Errorf("Max(5, 10) = %d; expected %d", result, expected)
	}
}

// TestMin tests the Min function
func TestMin(t *testing.T) {
	result := Min(10, 5)
	expected := 5
	if result != expected {
		t.Errorf("Min(10, 5) = %d; expected %d", result, expected)
	}
}

// TestCalculateArea tests the CalculateArea function
func TestCalculateArea(t *testing.T) {
	result := CalculateArea(5)
	expected := 78.53981633974483 // π * 5 * 5
	tolerance := 0.0001
	diff := result - expected
	if diff < 0 {
		diff = -diff
	}
	if diff > tolerance {
		t.Errorf("CalculateArea(5) = %.10f; expected approximately %.10f", result, expected)
	}
}

// TestFactorial tests the Factorial function
func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{3, 6},
	}

	for _, tt := range tests {
		result := Factorial(tt.input)
		if result != tt.expected {
			t.Errorf("Factorial(%d) = %d; expected %d", tt.input, result, tt.expected)
		}
	}
}

// TestHelper demonstrates helper functions
func TestHelper(t *testing.T) {
	// Helper function for testing
	checkResult := func(t *testing.T, got, want int) {
		t.Helper() // Marks this function as a test helper
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	checkResult(t, Add(2, 3), 5)
	checkResult(t, Subtract(5, 2), 3)
}

// BenchmarkAdd benchmarks the Add function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

