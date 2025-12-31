package main

import "testing"

func TestVariableDeclaration(t *testing.T) {
	var name string = "Go"
	if name != "Go" {
		t.Errorf("Expected 'Go', got %s", name)
	}
}

func TestTypeInference(t *testing.T) {
	var language = "Golang"
	if language != "Golang" {
		t.Errorf("Expected 'Golang', got %s", language)
	}
}

func TestShortDeclaration(t *testing.T) {
	author := "Google"
	if author != "Google" {
		t.Errorf("Expected 'Google', got %s", author)
	}
}

func TestMultipleVariables(t *testing.T) {
	x, y := 10, 20
	if x != 10 || y != 20 {
		t.Errorf("Expected x=10, y=20, got x=%d, y=%d", x, y)
	}
}

func TestConstants(t *testing.T) {
	const pi = 3.14159
	if pi != 3.14159 {
		t.Errorf("Expected 3.14159, got %f", pi)
	}
}

func TestZeroValues(t *testing.T) {
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	if zeroInt != 0 {
		t.Errorf("Expected 0 for int zero value, got %d", zeroInt)
	}
	if zeroFloat != 0.0 {
		t.Errorf("Expected 0.0 for float64 zero value, got %f", zeroFloat)
	}
	if zeroString != "" {
		t.Errorf("Expected empty string, got %q", zeroString)
	}
	if zeroBool != false {
		t.Errorf("Expected false for bool zero value, got %t", zeroBool)
	}
}

func TestTypeConversion(t *testing.T) {
	var integer int = 42
	var float float64 = float64(integer)
	var uintValue uint = uint(integer)

	if float != 42.0 {
		t.Errorf("Expected 42.0, got %f", float)
	}
	if uintValue != 42 {
		t.Errorf("Expected 42, got %d", uintValue)
	}
}

