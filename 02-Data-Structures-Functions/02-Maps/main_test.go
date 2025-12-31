package main

import "testing"

func TestMapDeclaration(t *testing.T) {
	var m map[string]int
	if m == nil {
		t.Log("Nil map created successfully")
	}
	
	m2 := make(map[string]int)
	if m2 == nil {
		t.Error("make() should not return nil map")
	}
}

func TestMapOperations(t *testing.T) {
	m := make(map[string]int)
	m["key1"] = 10
	m["key2"] = 20
	
	if m["key1"] != 10 {
		t.Errorf("Expected 10, got %d", m["key1"])
	}
	if m["key2"] != 20 {
		t.Errorf("Expected 20, got %d", m["key2"])
	}
}

func TestMapKeyExistence(t *testing.T) {
	m := map[string]int{
		"exists": 42,
	}
	
	value, exists := m["exists"]
	if !exists {
		t.Error("Key 'exists' should exist")
	}
	if value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}
	
	_, exists = m["notexists"]
	if exists {
		t.Error("Key 'notexists' should not exist")
	}
}

func TestMapZeroValue(t *testing.T) {
	m := make(map[string]int)
	value := m["nonexistent"]
	if value != 0 {
		t.Errorf("Expected zero value 0, got %d", value)
	}
}

func TestMapDelete(t *testing.T) {
	m := map[string]int{
		"key1": 10,
		"key2": 20,
		"key3": 30,
	}
	
	if len(m) != 3 {
		t.Errorf("Expected length 3, got %d", len(m))
	}
	
	delete(m, "key2")
	
	if len(m) != 2 {
		t.Errorf("Expected length 2 after delete, got %d", len(m))
	}
	
	if _, exists := m["key2"]; exists {
		t.Error("key2 should not exist after delete")
	}
}

func TestMapIteration(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	
	count := 0
	for range m {
		count++
	}
	
	if count != 3 {
		t.Errorf("Expected 3 iterations, got %d", count)
	}
}

func TestMapLength(t *testing.T) {
	m := make(map[string]int)
	if len(m) != 0 {
		t.Errorf("Expected length 0, got %d", len(m))
	}
	
	m["key1"] = 1
	m["key2"] = 2
	if len(m) != 2 {
		t.Errorf("Expected length 2, got %d", len(m))
	}
}

func TestMapAsSet(t *testing.T) {
	set := make(map[string]bool)
	set["item1"] = true
	set["item2"] = true
	
	if !set["item1"] {
		t.Error("item1 should be in set")
	}
	if set["item3"] {
		t.Error("item3 should not be in set")
	}
}

