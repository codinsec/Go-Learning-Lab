package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 3)
	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	if result != 20 {
		t.Errorf("Expected 20, got %d", result)
	}
}

func TestNewCalculator(t *testing.T) {
	calc := NewCalculator()
	if calc == nil {
		t.Error("NewCalculator should not return nil")
	}
}

func TestCalculatorAdd(t *testing.T) {
	calc := NewCalculator()
	calc.Add(10)
	calc.Add(20)
	result := calc.GetResult()
	if result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}
}

func TestCalculatorSubtract(t *testing.T) {
	calc := NewCalculator()
	calc.Add(100)
	result := calc.Subtract(30)
	if result != 70 {
		t.Errorf("Expected 70, got %d", result)
	}
}

