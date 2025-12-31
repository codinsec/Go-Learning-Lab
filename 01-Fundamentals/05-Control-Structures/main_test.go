package main

import "testing"

func TestIfStatement(t *testing.T) {
	x := 10
	if x <= 5 {
		t.Error("x should be greater than 5")
	}
}

func TestIfElse(t *testing.T) {
	age := 18
	var result string
	if age >= 18 {
		result = "adult"
	} else {
		result = "minor"
	}
	if result != "adult" {
		t.Errorf("Expected 'adult', got %s", result)
	}
}

func TestForLoop(t *testing.T) {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	expected := 0 + 1 + 2 + 3 + 4 // 10
	if sum != expected {
		t.Errorf("Expected sum %d, got %d", expected, sum)
	}
}

func TestForRangeSlice(t *testing.T) {
	numbers := []int{1, 2, 3}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestForRangeMap(t *testing.T) {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}
	count := 0
	for range colors {
		count++
	}
	if count != 2 {
		t.Errorf("Expected 2 items, got %d", count)
	}
}

func TestSwitchBasic(t *testing.T) {
	day := "Monday"
	var result string
	switch day {
	case "Monday":
		result = "weekday"
	case "Saturday", "Sunday":
		result = "weekend"
	default:
		result = "weekday"
	}
	if result != "weekday" {
		t.Errorf("Expected 'weekday', got %s", result)
	}
}

func TestSwitchNoExpression(t *testing.T) {
	value := 42
	var result string
	switch {
	case value < 0:
		result = "negative"
	case value == 0:
		result = "zero"
	default:
		result = "positive"
	}
	if result != "positive" {
		t.Errorf("Expected 'positive', got %s", result)
	}
}

func TestContinue(t *testing.T) {
	sum := 0
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	// Sum of odd numbers: 1 + 3 = 4
	if sum != 4 {
		t.Errorf("Expected sum 4, got %d", sum)
	}
}

func TestBreak(t *testing.T) {
	count := 0
	for i := 0; i < 10; i++ {
		if i >= 5 {
			break
		}
		count++
	}
	if count != 5 {
		t.Errorf("Expected count 5, got %d", count)
	}
}

