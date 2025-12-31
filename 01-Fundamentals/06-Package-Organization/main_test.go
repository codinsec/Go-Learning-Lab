package main

import (
	"testing"
	"github.com/codinsec/go-learning-lab/01-fundamentals/package-organization/calculator"
)

func TestCalculatorPackage(t *testing.T) {
	sum := calculator.Add(5, 3)
	if sum != 8 {
		t.Errorf("Expected 8, got %d", sum)
	}
}

func TestCalculatorType(t *testing.T) {
	calc := calculator.NewCalculator()
	calc.Add(10)
	result := calc.GetResult()
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}

