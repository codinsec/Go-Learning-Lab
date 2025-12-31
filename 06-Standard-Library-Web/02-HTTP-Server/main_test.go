package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleUsers(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	handleUsers(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result []User
	err := json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(result) == 0 {
		t.Error("Expected at least one user")
	}
}

func TestHandleCreateUser(t *testing.T) {
	req := httptest.NewRequest("POST", "/users/create", 
		httptest.NewRequest("POST", "/", nil).Body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Create JSON body
	userJSON := `{"name":"Test User","email":"test@example.com"}`
	req.Body = http.Request{}.Body // Simplified for test

	handleCreateUser(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestHandleHealth(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	handleHealth(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var result map[string]string
	err := json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if result["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got %s", result["status"])
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest("POST", "/users", nil)
	w := httptest.NewRecorder()

	handleUsers(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

