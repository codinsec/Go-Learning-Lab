package main

import (
	"reflect"
	"testing"
)

func TestTypeReflection(t *testing.T) {
	var x int = 42
	tp := reflect.TypeOf(x)
	if tp.Kind() != reflect.Int {
		t.Errorf("Expected Int kind, got %v", tp.Kind())
	}
}

func TestValueReflection(t *testing.T) {
	var x int = 42
	v := reflect.ValueOf(x)
	if v.Int() != 42 {
		t.Errorf("Expected 42, got %d", v.Int())
	}
}

func TestSetValue(t *testing.T) {
	var y int = 10
	vp := reflect.ValueOf(&y)
	elem := vp.Elem()
	elem.SetInt(20)
	if y != 20 {
		t.Errorf("Expected 20, got %d", y)
	}
}

func TestStructReflection(t *testing.T) {
	p := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	
	structType := reflect.TypeOf(p)
	if structType.NumField() != 4 { // Name, Age, Email, private
		t.Errorf("Expected 4 fields, got %d", structType.NumField())
	}
	
	// Check field names
	nameField, ok := structType.FieldByName("Name")
	if !ok {
		t.Error("Name field not found")
	}
	if nameField.Type.Kind() != reflect.String {
		t.Errorf("Expected String type for Name, got %v", nameField.Type.Kind())
	}
}

func TestGetFieldValue(t *testing.T) {
	p := Person{Name: "Bob", Age: 25}
	structValue := reflect.ValueOf(p)
	
	nameField := structValue.FieldByName("Name")
	if !nameField.IsValid() {
		t.Error("Name field is invalid")
	}
	if nameField.String() != "Bob" {
		t.Errorf("Expected 'Bob', got %s", nameField.String())
	}
}

func TestSetFieldValue(t *testing.T) {
	p := Person{Name: "Bob", Age: 25}
	pValue := reflect.ValueOf(&p).Elem()
	
	nameField := pValue.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Charlie")
	}
	
	if p.Name != "Charlie" {
		t.Errorf("Expected 'Charlie', got %s", p.Name)
	}
}

func TestMethodReflection(t *testing.T) {
	p := Person{Name: "Alice"}
	personType := reflect.TypeOf(p)
	
	// Should have at least GetName and String methods
	if personType.NumMethod() < 2 {
		t.Errorf("Expected at least 2 methods, got %d", personType.NumMethod())
	}
	
	_, ok := personType.MethodByName("GetName")
	if !ok {
		t.Error("GetName method not found")
	}
}

func TestCallMethod(t *testing.T) {
	p := Person{Name: "Alice"}
	personValue := reflect.ValueOf(p)
	
	getNameMethod := personValue.MethodByName("GetName")
	if !getNameMethod.IsValid() {
		t.Error("GetName method is invalid")
	}
	
	results := getNameMethod.Call(nil)
	if len(results) == 0 {
		t.Error("GetName should return a result")
	}
	if results[0].String() != "Alice" {
		t.Errorf("Expected 'Alice', got %s", results[0].String())
	}
}

func TestCreateInstance(t *testing.T) {
	personType := reflect.TypeOf(Person{})
	newPerson := reflect.New(personType).Elem()
	
	nameField := newPerson.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Eve")
	}
	
	personPtr := newPerson.Addr().Interface().(*Person)
	if personPtr.Name != "Eve" {
		t.Errorf("Expected 'Eve', got %s", personPtr.Name)
	}
}

func TestSliceReflection(t *testing.T) {
	slice := []int{1, 2, 3}
	sliceValue := reflect.ValueOf(slice)
	
	if sliceValue.Len() != 3 {
		t.Errorf("Expected length 3, got %d", sliceValue.Len())
	}
	
	firstElem := sliceValue.Index(0)
	if firstElem.Int() != 1 {
		t.Errorf("Expected 1, got %d", firstElem.Int())
	}
}

func TestMapReflection(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	mapValue := reflect.ValueOf(m)
	
	if mapValue.Len() != 2 {
		t.Errorf("Expected length 2, got %d", mapValue.Len())
	}
	
	key := reflect.ValueOf("a")
	value := mapValue.MapIndex(key)
	if !value.IsValid() {
		t.Error("Key 'a' not found in map")
	}
	if value.Int() != 1 {
		t.Errorf("Expected 1, got %d", value.Int())
	}
}

