package main

import "testing"

func TestTypeAssertion(t *testing.T) {
	var i interface{} = "hello"
	str, ok := i.(string)
	if !ok {
		t.Error("Type assertion should succeed")
	}
	if str != "hello" {
		t.Errorf("Expected 'hello', got %s", str)
	}
}

func TestTypeAssertionFailure(t *testing.T) {
	var i interface{} = 42
	_, ok := i.(string)
	if ok {
		t.Error("Type assertion should fail")
	}
}

func TestTypeAssertionWithInterface(t *testing.T) {
	var shape Shape = Rectangle{Width: 10, Height: 5}
	rect, ok := shape.(Rectangle)
	if !ok {
		t.Error("Should be able to assert to Rectangle")
	}
	if rect.Width != 10 {
		t.Errorf("Expected width 10, got %.2f", rect.Width)
	}
}

func TestTypeSwitch(t *testing.T) {
	var v interface{} = 42
	switch v.(type) {
	case int:
		// Correct type
	default:
		t.Error("Should match int case")
	}
}

func TestTypeSwitchWithInterface(t *testing.T) {
	var shape Shape = Circle{Radius: 5}
	switch shape.(type) {
	case Circle:
		// Correct type
	case Rectangle:
		t.Error("Should not match Rectangle")
	default:
		t.Error("Should match Circle case")
	}
}

func TestTypeSwitchMultipleCases(t *testing.T) {
	var v interface{} = int32(42)
	matched := false
	switch v.(type) {
	case int, int32, int64:
		matched = true
	}
	if !matched {
		t.Error("Should match integer types")
	}
}

func TestTypeSwitchDefault(t *testing.T) {
	var v interface{} = []int{1, 2, 3}
	matched := false
	switch v.(type) {
	case int:
		// Not matched
	case string:
		// Not matched
	default:
		matched = true
	}
	if !matched {
		t.Error("Should match default case")
	}
}

func TestNilTypeAssertion(t *testing.T) {
	var shape Shape
	if shape != nil {
		t.Error("Nil interface should be nil")
	}
	
	_, ok := shape.(Rectangle)
	if ok {
		t.Error("Nil interface should not assert to Rectangle")
	}
}

func TestPointerTypeAssertion(t *testing.T) {
	var ptr interface{} = &Rectangle{Width: 5, Height: 3}
	rectPtr, ok := ptr.(*Rectangle)
	if !ok {
		t.Error("Should assert to *Rectangle")
	}
	if rectPtr.Width != 5 {
		t.Errorf("Expected width 5, got %.2f", rectPtr.Width)
	}
}

