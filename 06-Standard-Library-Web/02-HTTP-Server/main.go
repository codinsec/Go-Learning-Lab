package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// This program demonstrates HTTP server in Go

// User struct for API
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory user store
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

var nextID = 3

func main() {
	fmt.Println("=== HTTP Server ===")
	fmt.Println()
	fmt.Println("Starting server on :8080")
	fmt.Println("Endpoints:")
	fmt.Println("  GET  /users     - Get all users")
	fmt.Println("  GET  /users/:id - Get user by ID")
	fmt.Println("  POST /users     - Create user")
	fmt.Println("  GET  /health    - Health check")
	fmt.Println()

	// 1. Basic route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Go HTTP Server!")
	})

	// 2. GET /users - Get all users
	http.HandleFunc("/users", handleUsers)

	// 3. GET /users/:id - Get user by ID
	http.HandleFunc("/users/", handleUserByID)

	// 4. POST /users - Create user
	http.HandleFunc("/users/create", handleCreateUser)

	// 5. Health check
	http.HandleFunc("/health", handleHealth)

	// 6. Custom middleware example
	http.HandleFunc("/protected", loggingMiddleware(authMiddleware(handleProtected)))

	// 7. Start server
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Server started. Visit http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")
	
	// Uncomment to start server:
	// log.Fatal(server.ListenAndServe())
	
	// For demonstration, we'll just show the setup
	fmt.Println("\n(Server not started - uncomment ListenAndServe to run)")
}

// Handle GET /users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handle GET /users/:id
func handleUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Simple ID extraction (in production, use a router)
	id := 1 // Simplified for example
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// Handle POST /users/create
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user.ID = nextID
	nextID++
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Handle GET /health
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// Handle protected route
func handleProtected(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "This is a protected route",
	})
}

// Middleware: Logging
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Completed in %v", time.Since(start))
	}
}

// Middleware: Authentication (simplified)
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

