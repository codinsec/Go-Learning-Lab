package main

import "testing"

func TestBasicPointer(t *testing.T) {
	x := 42
	p := &x
	
	if *p != 42 {
		t.Errorf("Expected 42, got %d", *p)
	}
	
	*p = 100
	if x != 100 {
		t.Errorf("Expected 100, got %d", x)
	}
}

func TestNilPointer(t *testing.T) {
	var p *int
	if p != nil {
		t.Error("Pointer should be nil")
	}
}

func TestNewPointer(t *testing.T) {
	p := new(int)
	if p == nil {
		t.Error("new() should not return nil")
	}
	if *p != 0 {
		t.Errorf("Expected zero value 0, got %d", *p)
	}
	
	*p = 42
	if *p != 42 {
		t.Errorf("Expected 42, got %d", *p)
	}
}

func TestSwapByReference(t *testing.T) {
	a, b := 10, 20
	swapByReference(&a, &b)
	if a != 20 || b != 10 {
		t.Errorf("Expected a=20, b=10, got a=%d, b=%d", a, b)
	}
}

func TestSwapByValue(t *testing.T) {
	a, b := 10, 20
	swapByValue(a, b)
	// Values should be unchanged
	if a != 10 || b != 20 {
		t.Errorf("Values should be unchanged: a=%d, b=%d", a, b)
	}
}

func TestCounterValueReceiver(t *testing.T) {
	counter := Counter{value: 0}
	counter.Increment()
	// Value receiver works on copy, so original shouldn't change
	if counter.value != 0 {
		t.Errorf("Expected 0 (value receiver doesn't modify), got %d", counter.value)
	}
}

func TestCounterPointerReceiver(t *testing.T) {
	counter := Counter{value: 0}
	counter.IncrementByPointer()
	// Pointer receiver modifies original
	if counter.value != 1 {
		t.Errorf("Expected 1, got %d", counter.value)
	}
}

func TestCreateInt(t *testing.T) {
	p := createInt(42)
	if p == nil {
		t.Error("createInt should not return nil")
	}
	if *p != 42 {
		t.Errorf("Expected 42, got %d", *p)
	}
}

func TestStructPointer(t *testing.T) {
	person := Person{Name: "Alice", Age: 30}
	personPtr := &person
	
	if personPtr.Name != "Alice" {
		t.Errorf("Expected 'Alice', got %s", personPtr.Name)
	}
	
	personPtr.Age = 31
	if person.Age != 31 {
		t.Errorf("Expected 31, got %d", person.Age)
	}
}

func TestArrayPointer(t *testing.T) {
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	
	if arrPtr[0] != 1 {
		t.Errorf("Expected 1, got %d", arrPtr[0])
	}
	
	(*arrPtr)[1] = 99
	if arr[1] != 99 {
		t.Errorf("Expected 99, got %d", arr[1])
	}
}

