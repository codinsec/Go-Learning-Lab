package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// This program demonstrates HTTP client in Go

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("=== HTTP Client ===")
	fmt.Println()

	// 1. Basic GET request
	fmt.Println("1. Basic GET Request:")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("   Status: %s\n", resp.Status)
		fmt.Printf("   Response length: %d bytes\n", len(body))
	}
	fmt.Println()

	// 2. GET with custom client
	fmt.Println("2. GET with Custom Client:")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp2, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp2.Body.Close()
		fmt.Printf("   Status: %s\n", resp2.Status)
	}
	fmt.Println()

	// 3. POST request
	fmt.Println("3. POST Request:")
	post := Post{
		Title: "Test Post",
		Body:  "This is a test post",
	}
	
	jsonData, _ := json.Marshal(post)
	resp3, err := http.Post("https://jsonplaceholder.typicode.com/posts", 
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp3.Body.Close()
		fmt.Printf("   Status: %s\n", resp3.Status)
	}
	fmt.Println()

	// 4. Custom request with headers
	fmt.Println("4. Custom Request with Headers:")
	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	req.Header.Set("User-Agent", "Go-Learning-Lab/1.0")
	req.Header.Set("Accept", "application/json")
	
	resp4, err := client.Do(req)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		defer resp4.Body.Close()
		fmt.Printf("   Status: %s\n", resp4.Status)
	}
	fmt.Println()

	// 5. Reading response body
	fmt.Println("5. Reading Response Body:")
	resp5, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err == nil {
		defer resp5.Body.Close()
		body, _ := io.ReadAll(resp5.Body)
		var post Post
		json.Unmarshal(body, &post)
		fmt.Printf("   Post ID: %d\n", post.ID)
		fmt.Printf("   Post Title: %s\n", post.Title)
	}
	fmt.Println()

	// 6. Error handling
	fmt.Println("6. Error Handling:")
	resp6, err := http.Get("https://invalid-url-that-does-not-exist.com")
	if err != nil {
		fmt.Printf("   Error (expected): %v\n", err)
	} else {
		defer resp6.Body.Close()
	}
	fmt.Println()

	// 7. Response status codes
	fmt.Println("7. Response Status Codes:")
	resp7, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err == nil {
		defer resp7.Body.Close()
		fmt.Printf("   Status Code: %d\n", resp7.StatusCode)
		fmt.Printf("   Status: %s\n", resp7.Status)
		
		if resp7.StatusCode == http.StatusOK {
			fmt.Println("   Request successful!")
		}
	}
	fmt.Println()

	// 8. Client configuration
	fmt.Println("8. Client Configuration:")
	customClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  false,
		},
	}
	fmt.Printf("   Client timeout: %v\n", customClient.Timeout)
	fmt.Println()

	// 9. Best practices
	fmt.Println("9. Best Practices:")
	fmt.Println("   - Always close response body (use defer)")
	fmt.Println("   - Set timeouts on client")
	fmt.Println("   - Check status codes")
	fmt.Println("   - Handle errors properly")
	fmt.Println("   - Use context for cancellation")
	fmt.Println("   - Reuse clients (don't create new client per request)")
}

