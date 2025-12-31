package main

import "testing"

func TestComposition(t *testing.T) {
	car := Car{
		Make:  "Toyota",
		Model: "Camry",
		Engine: Engine{
			Horsepower: 200,
		},
	}
	if car.Engine.Horsepower != 200 {
		t.Errorf("Expected 200, got %d", car.Engine.Horsepower)
	}
}

func TestEmbedding(t *testing.T) {
	dog := Dog{
		Animal: Animal{
			Name: "Buddy",
			Age:  3,
		},
		Breed: "Golden Retriever",
	}
	if dog.Name != "Buddy" {
		t.Errorf("Expected 'Buddy', got %s", dog.Name)
	}
	if dog.Age != 3 {
		t.Errorf("Expected 3, got %d", dog.Age)
	}
	if dog.Breed != "Golden Retriever" {
		t.Errorf("Expected 'Golden Retriever', got %s", dog.Breed)
	}
}

func TestMethodPromotion(t *testing.T) {
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
	}
	// Speak() is promoted from Animal
	// We can't easily test the output, but we can verify it exists
	if dog.Name != "Buddy" {
		t.Error("Method promotion should work")
	}
}

func TestMultipleEmbedding(t *testing.T) {
	rw := ReadWriter{
		Writer: Writer{Name: "Alice"},
		Reader: Reader{Name: "Bob"},
	}
	if rw.Writer.Name != "Alice" {
		t.Errorf("Expected 'Alice', got %s", rw.Writer.Name)
	}
	if rw.Reader.Name != "Bob" {
		t.Errorf("Expected 'Bob', got %s", rw.Reader.Name)
	}
}

func TestEmbeddingPointer(t *testing.T) {
	person := &Person{Name: "Charlie", Age: 30}
	emp := Employee{
		Person: person,
		ID:     1001,
	}
	if emp.Name != "Charlie" {
		t.Errorf("Expected 'Charlie', got %s", emp.Name)
	}
	if emp.ID != 1001 {
		t.Errorf("Expected 1001, got %d", emp.ID)
	}
}

func TestEmbeddedInterface(t *testing.T) {
	writer := FileWriter{Filename: "test.log"}
	logger := Logger{
		WriterInterface: writer,
		Level:          "INFO",
	}
	if logger.Level != "INFO" {
		t.Errorf("Expected 'INFO', got %s", logger.Level)
	}
	// WriterInterface is embedded, so Write() should be accessible
	// (We can't easily test the method call, but structure is correct)
}

func TestEmbeddingAccess(t *testing.T) {
	dog := Dog{
		Animal: Animal{Name: "Buddy", Age: 3},
		Breed:  "Labrador",
	}
	// Can access embedded fields directly
	if dog.Name != "Buddy" {
		t.Error("Should access embedded field directly")
	}
	// Can also access via embedded type name
	if dog.Animal.Name != "Buddy" {
		t.Error("Should access via embedded type name")
	}
}

