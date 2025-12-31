package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	
	jsonData, err := json.Marshal(person)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	if len(jsonData) == 0 {
		t.Error("JSON data should not be empty")
	}
}

func TestUnmarshal(t *testing.T) {
	jsonStr := `{"name":"Bob","age":25}`
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	
	if person.Name != "Bob" {
		t.Errorf("Expected name 'Bob', got %s", person.Name)
	}
	if person.Age != 25 {
		t.Errorf("Expected age 25, got %d", person.Age)
	}
}

func TestOmitempty(t *testing.T) {
	person := Person{
		Name:  "Charlie",
		Age:   28,
		Email: "", // Should be omitted
	}
	
	jsonData, err := json.Marshal(person)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	jsonStr := string(jsonData)
	if contains(jsonStr, "email") {
		t.Error("Email should be omitted when empty")
	}
}

func TestIgnoreField(t *testing.T) {
	person := Person{
		Name:    "David",
		Age:     30,
		Address: "123 Main St", // Should be ignored
	}
	
	jsonData, err := json.Marshal(person)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	jsonStr := string(jsonData)
	if contains(jsonStr, "Address") || contains(jsonStr, "address") {
		t.Error("Address field should be ignored")
	}
}

func TestNestedStruct(t *testing.T) {
	employee := Employee{
		ID:   1001,
		Name: "Eve",
		Address: Address{
			Street:  "456 Oak Ave",
			City:    "New York",
			Country: "USA",
		},
	}
	
	jsonData, err := json.Marshal(employee)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	var result Employee
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	
	if result.Address.City != "New York" {
		t.Errorf("Expected city 'New York', got %s", result.Address.City)
	}
}

func TestCustomMarshaling(t *testing.T) {
	event := Event{
		Title: "Conference",
		Date:  CustomDate{Time: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)},
	}
	
	jsonData, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	jsonStr := string(jsonData)
	if !contains(jsonStr, "2024-01-15") {
		t.Error("Date should be formatted as YYYY-MM-DD")
	}
}

func TestUnmarshalToMap(t *testing.T) {
	jsonStr := `{"name":"Frank","age":35}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	
	if data["name"] != "Frank" {
		t.Errorf("Expected name 'Frank', got %v", data["name"])
	}
}

func TestArray(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	
	jsonData, err := json.Marshal(people)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	var result []Person
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	
	if len(result) != 2 {
		t.Errorf("Expected 2 people, got %d", len(result))
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		indexOfSubstring(s, substr) >= 0)
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

