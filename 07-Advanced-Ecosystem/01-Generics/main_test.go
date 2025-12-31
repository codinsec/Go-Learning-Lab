package main

import "testing"

func TestGenericPrint(t *testing.T) {
	// Test that Print works with different types
	Print("test")
	Print(42)
	Print(3.14)
	// If no panic, test passes
}

func TestMax(t *testing.T) {
	if Max(10, 20) != 20 {
		t.Errorf("Expected 20, got %d", Max(10, 20))
	}
	
	if Max(3.14, 2.71) != 3.14 {
		t.Errorf("Expected 3.14, got %f", Max(3.14, 2.71))
	}
}

func TestAdd(t *testing.T) {
	if Add(5, 3) != 8 {
		t.Errorf("Expected 8, got %d", Add(5, 3))
	}
	
	if Add(3.14, 2.71) != 5.85 {
		t.Errorf("Expected 5.85, got %f", Add(3.14, 2.71))
	}
}

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	
	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}
	
	stack.Push(1)
	stack.Push(2)
	
	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}
	
	val, ok := stack.Pop()
	if !ok {
		t.Error("Pop should return true")
	}
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	
	if stack.Size() != 1 {
		t.Errorf("Expected size 1, got %d", stack.Size())
	}
}

func TestContainer(t *testing.T) {
	container := Container[int]{value: 42}
	if container.Get() != 42 {
		t.Errorf("Expected 42, got %d", container.Get())
	}
	
	container.Set(100)
	if container.Get() != 100 {
		t.Errorf("Expected 100, got %d", container.Get())
	}
}

func TestCompare(t *testing.T) {
	if !Compare(5, 5) {
		t.Error("Expected true for Compare(5, 5)")
	}
	
	if Compare(5, 10) {
		t.Error("Expected false for Compare(5, 10)")
	}
}

func TestPair(t *testing.T) {
	first, second := Pair("Alice", 30)
	if first != "Alice" {
		t.Errorf("Expected 'Alice', got %s", first)
	}
	if second != 30 {
		t.Errorf("Expected 30, got %d", second)
	}
}

func TestMultiply(t *testing.T) {
	type MyInt int
	var a MyInt = 5
	var b MyInt = 3
	result := Multiply(a, b)
	if result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}
}

