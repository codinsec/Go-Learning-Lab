# HTTP Server

Learn about building HTTP servers in Go using the `net/http` package.

## Learning Objectives

- Understand how to create HTTP servers
- Learn about HTTP handlers and handler functions
- Master routing and URL patterns
- Understand middleware patterns
- Learn about request and response handling
- Understand server configuration

## Concepts Covered

### Basic HTTP Server

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})

http.ListenAndServe(":8080", nil)
```

### Handler Functions

Functions that handle HTTP requests:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Handle request
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Response")
}
```

### Request Methods

Check HTTP method:

```go
if r.Method != http.MethodGet {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}
```

### JSON Responses

Send JSON responses:

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)
```

### Reading Request Body

Read JSON from request:

```go
var user User
json.NewDecoder(r.Body).Decode(&user)
```

### Middleware

Functions that wrap handlers:

```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}
```

### Server Configuration

Configure server timeouts:

```go
server := &http.Server{
    Addr:         ":8080",
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
}
server.ListenAndServe()
```

## Running the Example

```bash
# Navigate to this directory
cd 02-HTTP-Server

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/http-server

# Run the program
go run main.go

# In another terminal, test endpoints:
# curl http://localhost:8080/users
# curl http://localhost:8080/health
```

## Running Tests

```bash
# Run all tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover
```

## Key Takeaways

1. **http.HandleFunc** - Register route handlers
2. **http.ResponseWriter** - Write responses
3. **http.Request** - Access request data
4. **Middleware** - Wrap handlers for cross-cutting concerns
5. **JSON handling** - Use encoding/json for APIs
6. **Error handling** - Use http.Error for errors

## Common Patterns

**Basic handler:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Response")
}
```

**JSON API:**
```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)
```

**Middleware chain:**
```go
http.HandleFunc("/path", middleware1(middleware2(handler)))
```

## Best Practices

- **Check HTTP methods** - Validate allowed methods
- **Set content type** - For JSON responses
- **Handle errors** - Use http.Error appropriately
- **Use middleware** - For logging, auth, etc.
- **Configure timeouts** - Prevent resource exhaustion
- **Validate input** - Always validate request data

## Important Notes

- **Default router** - http.ServeMux is basic (consider gorilla/mux for complex routing)
- **Concurrent handlers** - Handlers run concurrently
- **Request body** - Can only be read once
- **Response headers** - Must be set before writing body

## Next Steps

After understanding HTTP servers, proceed to:
- [03-HTTP-Client](../03-HTTP-Client/README.md) - Learn about HTTP clients

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

