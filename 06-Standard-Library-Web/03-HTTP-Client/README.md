# HTTP Client

Learn about making HTTP requests in Go using the `net/http` package.

## Learning Objectives

- Understand how to make HTTP GET requests
- Learn how to make HTTP POST requests
- Master custom HTTP requests with headers
- Understand client configuration
- Learn about response handling
- Understand error handling and timeouts

## Concepts Covered

### Basic GET Request

```go
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    // Handle error
}
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
```

### Custom Client

Create a client with timeout:

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}
resp, err := client.Get("https://api.example.com/data")
```

### POST Request

Send POST request with JSON:

```go
jsonData, _ := json.Marshal(data)
resp, err := http.Post("https://api.example.com/data", 
    "application/json", bytes.NewBuffer(jsonData))
```

### Custom Request

Create custom request with headers:

```go
req, _ := http.NewRequest("GET", url, nil)
req.Header.Set("Authorization", "Bearer token")
req.Header.Set("Content-Type", "application/json")

resp, err := client.Do(req)
```

### Reading Response

Read and parse response:

```go
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
var data Data
json.Unmarshal(body, &data)
```

### Client Configuration

Configure client with timeouts and transport:

```go
client := &http.Client{
    Timeout: 5 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns: 10,
        IdleConnTimeout: 30 * time.Second,
    },
}
```

## Running the Example

```bash
# Navigate to this directory
cd 03-HTTP-Client

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/http-client

# Run the program
go run main.go
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

1. **Always close response body** - Use defer resp.Body.Close()
2. **Set timeouts** - Prevent hanging requests
3. **Check status codes** - Validate response status
4. **Handle errors** - Network requests can fail
5. **Reuse clients** - Don't create new client per request
6. **Use context** - For cancellation and timeouts

## Common Patterns

**Basic GET:**
```go
resp, err := http.Get(url)
defer resp.Body.Close()
```

**POST with JSON:**
```go
jsonData, _ := json.Marshal(data)
resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
```

**Custom request:**
```go
req, _ := http.NewRequest("GET", url, nil)
req.Header.Set("Key", "Value")
resp, err := client.Do(req)
```

## Best Practices

- **Always close body** - Use defer
- **Set timeouts** - Prevent resource exhaustion
- **Check status codes** - Validate responses
- **Handle errors** - Network can fail
- **Reuse clients** - More efficient
- **Use context** - For cancellation

## Important Notes

- **Response body must be closed** - Or resources leak
- **Timeout is important** - Default client has no timeout
- **Status codes** - Check resp.StatusCode
- **Body can only be read once** - Read and store if needed multiple times

## Next Steps

After understanding HTTP clients, proceed to:
- [04-File-Operations](../04-File-Operations/README.md) - Learn about file operations

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

