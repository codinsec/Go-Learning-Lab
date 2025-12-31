package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %.2f", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
	if err.Error() != "division by zero" {
		t.Errorf("Expected 'division by zero', got %v", err)
	}
}

func TestValidateAge(t *testing.T) {
	err := validateAge(25)
	if err != nil {
		t.Errorf("Expected no error for valid age, got %v", err)
	}
}

func TestValidateAgeNegative(t *testing.T) {
	err := validateAge(-5)
	if err == nil {
		t.Error("Expected error for negative age")
	}
	
	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Error("Expected ValidationError")
	}
	if validationErr.Field != "age" {
		t.Errorf("Expected field 'age', got %s", validationErr.Field)
	}
}

func TestValidateAgeTooLarge(t *testing.T) {
	err := validateAge(200)
	if err == nil {
		t.Error("Expected error for age > 150")
	}
	
	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Error("Expected ValidationError")
	}
}

func TestFindUser(t *testing.T) {
	user, err := findUser("123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Error("Expected user, got nil")
	}
	if user.ID != "123" {
		t.Errorf("Expected ID '123', got %s", user.ID)
	}
}

func TestFindUserEmptyID(t *testing.T) {
	_, err := findUser("")
	if err == nil {
		t.Error("Expected error for empty ID")
	}
}

func TestFindUserNotFound(t *testing.T) {
	_, err := findUser("999")
	if err == nil {
		t.Error("Expected error for not found")
	}
	
	var notFound NotFoundError
	if !errors.As(err, &notFound) {
		t.Error("Expected NotFoundError")
	}
}

func TestErrorsIs(t *testing.T) {
	err := errors.New("test error")
	if errors.Is(err, errors.New("test error")) {
		// errors.Is compares by value, not by reference
		// This might not work as expected
	}
}

func TestErrorsAs(t *testing.T) {
	err := ValidationError{Field: "age", Message: "invalid"}
	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Error("errors.As should work")
	}
	if validationErr.Field != "age" {
		t.Errorf("Expected field 'age', got %s", validationErr.Field)
	}
}

func TestErrorUnwrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := fmt.Errorf("wrapped: %w", originalErr)
	
	unwrapped := errors.Unwrap(wrappedErr)
	if unwrapped != originalErr {
		t.Error("Unwrap should return original error")
	}
}

