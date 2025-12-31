package main

import "testing"

func TestPersonStruct(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	if p.Name != "Alice" {
		t.Errorf("Expected Name 'Alice', got %s", p.Name)
	}
	if p.Age != 30 {
		t.Errorf("Expected Age 30, got %d", p.Age)
	}
}

func TestPersonZeroValue(t *testing.T) {
	var p Person
	if p.Name != "" {
		t.Errorf("Expected empty string, got %s", p.Name)
	}
	if p.Age != 0 {
		t.Errorf("Expected 0, got %d", p.Age)
	}
}

func TestStructComparison(t *testing.T) {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Alice", Age: 30}
	p3 := Person{Name: "Bob", Age: 30}
	
	if p1 != p2 {
		t.Error("p1 and p2 should be equal")
	}
	if p1 == p3 {
		t.Error("p1 and p3 should not be equal")
	}
}

func TestStructModification(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	p.Name = "Bob"
	p.Age = 25
	
	if p.Name != "Bob" {
		t.Errorf("Expected 'Bob', got %s", p.Name)
	}
	if p.Age != 25 {
		t.Errorf("Expected 25, got %d", p.Age)
	}
}

func TestNestedStruct(t *testing.T) {
	contact := Contact{
		Email: "test@example.com",
		Phone: "555-1234",
		Address: Address{
			City: "New York",
		},
	}
	
	if contact.Address.City != "New York" {
		t.Errorf("Expected 'New York', got %s", contact.Address.City)
	}
}

func TestPointerToStruct(t *testing.T) {
	p := &Person{Name: "Alice", Age: 30}
	if p.Name != "Alice" {
		t.Errorf("Expected 'Alice', got %s", p.Name)
	}
	
	p.Age = 31
	if p.Age != 31 {
		t.Errorf("Expected 31, got %d", p.Age)
	}
}

func TestModifyByValue(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	modifyPerson(p)
	// Original should be unchanged
	if p.Age != 30 {
		t.Errorf("Expected 30 (unchanged), got %d", p.Age)
	}
}

func TestModifyByPointer(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	modifyPersonByPointer(&p)
	// Original should be changed
	if p.Age != 99 {
		t.Errorf("Expected 99 (changed), got %d", p.Age)
	}
}

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	if area != 50 {
		t.Errorf("Expected area 50, got %d", area)
	}
}

