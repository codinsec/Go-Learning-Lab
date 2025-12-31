package main

import "testing"

func TestValueReceiver(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	if area != 50 {
		t.Errorf("Expected area 50, got %.2f", area)
	}
}

func TestValueReceiverDoesNotModify(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	scaled := rect.Scale(2.0)
	if rect.Width != 10 {
		t.Errorf("Original should be unchanged, got %.2f", rect.Width)
	}
	if scaled.Width != 20 {
		t.Errorf("Scaled should be 20, got %.2f", scaled.Width)
	}
}

func TestPointerReceiverModifies(t *testing.T) {
	rect := &Rectangle{Width: 10, Height: 5}
	rect.ScaleInPlace(2.0)
	if rect.Width != 20 {
		t.Errorf("Expected 20, got %.2f", rect.Width)
	}
	if rect.Height != 10 {
		t.Errorf("Expected 10, got %.2f", rect.Height)
	}
}

func TestSetDimensions(t *testing.T) {
	rect := &Rectangle{Width: 10, Height: 5}
	rect.SetDimensions(20, 15)
	if rect.Width != 20 {
		t.Errorf("Expected 20, got %.2f", rect.Width)
	}
	if rect.Height != 15 {
		t.Errorf("Expected 15, got %.2f", rect.Height)
	}
}

func TestCounterValueReceiver(t *testing.T) {
	counter := Counter{value: 0}
	counter.Increment() // Value receiver doesn't modify
	if counter.GetValue() != 0 {
		t.Errorf("Expected 0 (unchanged), got %d", counter.GetValue())
	}
}

func TestCounterPointerReceiver(t *testing.T) {
	counter := Counter{value: 0}
	counter.IncrementByPointer() // Pointer receiver modifies
	if counter.GetValue() != 1 {
		t.Errorf("Expected 1 (changed), got %d", counter.GetValue())
	}
}

func TestCounterReset(t *testing.T) {
	counter := &Counter{value: 5}
	counter.Reset()
	if counter.GetValue() != 0 {
		t.Errorf("Expected 0, got %d", counter.GetValue())
	}
}

func TestPersonMethods(t *testing.T) {
	person := Person{Name: "Alice", Age: 30}
	info := person.GetInfo()
	if info == "" {
		t.Error("GetInfo should return non-empty string")
	}
	
	person.HaveBirthday()
	if person.Age != 31 {
		t.Errorf("Expected 31, got %d", person.Age)
	}
}

func TestMyIntMethods(t *testing.T) {
	num := MyInt(4)
	if !num.IsEven() {
		t.Error("4 should be even")
	}
	
	doubled := num.Double()
	if doubled != 8 {
		t.Errorf("Expected 8, got %d", doubled)
	}
}

func TestMethodChaining(t *testing.T) {
	rect := &Rectangle{Width: 2, Height: 2}
	rect.SetDimensions(5, 5)
	if rect.Width != 5 || rect.Height != 5 {
		t.Error("SetDimensions should modify rectangle")
	}
}

