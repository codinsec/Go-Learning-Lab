package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(5, 3)
	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	quotient, remainder := divide(17, 5)
	if quotient != 3 {
		t.Errorf("Expected quotient 3, got %d", quotient)
	}
	if remainder != 2 {
		t.Errorf("Expected remainder 2, got %d", remainder)
	}
}

func TestCalculate(t *testing.T) {
	sum, product := calculate(5, 6)
	if sum != 11 {
		t.Errorf("Expected sum 11, got %d", sum)
	}
	if product != 30 {
		t.Errorf("Expected product 30, got %d", product)
	}
}

func TestSafeDivide(t *testing.T) {
	result, err := safeDivide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}

	result, err = safeDivide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
	if result != 0 {
		t.Errorf("Expected 0 on error, got %d", result)
	}
}

func TestSum(t *testing.T) {
	if sum(1, 2, 3) != 6 {
		t.Errorf("Expected 6, got %d", sum(1, 2, 3))
	}
	if sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected 15, got %d", sum(1, 2, 3, 4, 5))
	}
	if sum() != 0 {
		t.Errorf("Expected 0 for empty sum, got %d", sum())
	}
}

func TestApplyOperation(t *testing.T) {
	result := applyOperation(10, 5, add)
	if result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}

	result = applyOperation(10, 5, multiply)
	if result != 50 {
		t.Errorf("Expected 50, got %d", result)
	}
}

func TestMakeCounter(t *testing.T) {
	counter := makeCounter()
	if counter() != 1 {
		t.Errorf("Expected 1, got %d", counter())
	}
	if counter() != 2 {
		t.Errorf("Expected 2, got %d", counter())
	}
	if counter() != 3 {
		t.Errorf("Expected 3, got %d", counter())
	}
}

func TestSwap(t *testing.T) {
	x, y := 10, 20
	swap(&x, &y)
	if x != 20 || y != 10 {
		t.Errorf("Expected x=20, y=10, got x=%d, y=%d", x, y)
	}
}

