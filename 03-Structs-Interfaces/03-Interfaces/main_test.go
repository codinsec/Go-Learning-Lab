package main

import "testing"

func TestRectangleImplementsShape(t *testing.T) {
	var shape Shape = Rectangle{Width: 10, Height: 5}
	area := shape.Area()
	if area != 50 {
		t.Errorf("Expected area 50, got %.2f", area)
	}
}

func TestCircleImplementsShape(t *testing.T) {
	var shape Shape = Circle{Radius: 5}
	area := shape.Area()
	expected := 3.14159 * 5 * 5
	if area != expected {
		t.Errorf("Expected area %.2f, got %.2f", expected, area)
	}
}

func TestPolymorphism(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 5},
	}
	
	if len(shapes) != 2 {
		t.Errorf("Expected 2 shapes, got %d", len(shapes))
	}
	
	// Both should have Area method
	for _, shape := range shapes {
		area := shape.Area()
		if area <= 0 {
			t.Error("Area should be positive")
		}
	}
}

func TestEmptyInterface(t *testing.T) {
	// Empty interface can hold any value
	var v interface{}
	v = 42
	if v != 42 {
		t.Error("Empty interface should hold value")
	}
	
	v = "hello"
	if v != "hello" {
		t.Error("Empty interface should hold string")
	}
}

func TestTypeAssertion(t *testing.T) {
	var i interface{} = "hello"
	str, ok := i.(string)
	if !ok {
		t.Error("Type assertion should succeed")
	}
	if str != "hello" {
		t.Errorf("Expected 'hello', got %s", str)
	}
	
	_, ok = i.(int)
	if ok {
		t.Error("Type assertion to int should fail")
	}
}

func TestInterfaceComposition(t *testing.T) {
	var rw ReadWriter = &FileReadWriter{Filename: "test.txt"}
	data := make([]byte, 100)
	n, err := rw.Read(data)
	if err != nil {
		t.Errorf("Read should not error, got %v", err)
	}
	if n == 0 {
		t.Error("Read should return some bytes")
	}
}

func TestAnimalInterface(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	if dog.Speak() == "" {
		t.Error("Dog should speak")
	}
	if dog.Move() == "" {
		t.Error("Dog should move")
	}
	
	cat := Cat{Name: "Whiskers"}
	if cat.Speak() == "" {
		t.Error("Cat should speak")
	}
	if cat.Move() == "" {
		t.Error("Cat should move")
	}
}

func TestNilInterface(t *testing.T) {
	var shape Shape
	if shape != nil {
		t.Error("Uninitialized interface should be nil")
	}
}

