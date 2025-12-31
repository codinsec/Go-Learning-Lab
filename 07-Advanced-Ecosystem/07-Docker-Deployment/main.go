package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// This program demonstrates Docker deployment for Go applications

func main() {
	// Simple HTTP server for demonstration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)

	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("=== Docker Deployment ===")
	fmt.Println()
	fmt.Println("1. Dockerfile Examples:")
	fmt.Println("   See Dockerfile and Dockerfile.multistage")
	fmt.Println()
	fmt.Println("2. Docker Commands:")
	fmt.Println("   docker build -t myapp .")
	fmt.Println("   docker run -p 8080:8080 myapp")
	fmt.Println("   docker run -p 8080:8080 -e PORT=8080 myapp")
	fmt.Println()
	fmt.Println("3. Docker Compose:")
	fmt.Println("   See docker-compose.yml")
	fmt.Println()
	fmt.Println("4. Best Practices:")
	fmt.Println("   - Use multi-stage builds")
	fmt.Println("   - Use .dockerignore")
	fmt.Println("   - Minimize image size")
	fmt.Println("   - Use specific base images")
	fmt.Println("   - Set proper working directory")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Dockerized Go app!\n")
	fmt.Fprintf(w, "Environment: %s\n", os.Getenv("ENV"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}

