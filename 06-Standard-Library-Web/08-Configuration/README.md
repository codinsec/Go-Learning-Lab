# Configuration

Learn about configuration management in Go: environment variables, command-line flags, and config files.

## Learning Objectives

- Understand environment variables
- Learn about command-line flags
- Master config file parsing
- Understand configuration priority
- Learn best practices for configuration
- Understand secrets management

## Concepts Covered

### Environment Variables

```go
import "os"

host := os.Getenv("APP_HOST")
port := os.Getenv("APP_PORT")

// With default
database := os.Getenv("APP_DATABASE")
if database == "" {
    database = "default.db"
}
```

### Command-Line Flags

```go
import "flag"

host := flag.String("host", "localhost", "Server host")
port := flag.Int("port", 8080, "Server port")
debug := flag.Bool("debug", false, "Enable debug")

flag.Parse() // Parse command-line arguments
```

### Config File (JSON)

```go
type Config struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Database string `json:"database"`
}

// Read config
data, _ := os.ReadFile("config.json")
var config Config
json.Unmarshal(data, &config)
```

### Combined Approach

Priority order:
1. Environment variables (highest)
2. Command-line flags
3. Config file
4. Defaults (lowest)

## Running the Example

```bash
# Navigate to this directory
cd 08-Configuration

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/configuration

# Run the program
go run main.go

# With flags:
go run main.go -host=0.0.0.0 -port=3000 -debug
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

1. **Environment variables** - For secrets and deployment-specific config
2. **Command-line flags** - For runtime options
3. **Config files** - For complex configurations
4. **Priority order** - Env > Flags > File > Defaults
5. **Never commit secrets** - Use environment variables
6. **Validate config** - Check required values

## Common Patterns

**Environment variable with default:**
```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```

**Flag parsing:**
```go
flag.Parse()
if *host == "" {
    log.Fatal("Host is required")
}
```

**Config file:**
```go
config := loadConfig("config.json")
```

## Best Practices

- **Use environment variables** - For secrets and deployment config
- **Use flags** - For command-line options
- **Use config files** - For complex settings
- **Validate config** - Check required values
- **Never commit secrets** - Use .env files or environment
- **Document defaults** - Make defaults clear

## Important Notes

- **Environment variables** - Set by OS or deployment
- **Flags must be parsed** - Call flag.Parse()
- **Config file format** - JSON, YAML, TOML, etc.
- **Priority matters** - Later sources override earlier

## Configuration Priority

1. **Environment variables** - Highest priority
2. **Command-line flags** - Second priority
3. **Config file** - Third priority
4. **Defaults** - Lowest priority

## Next Steps

Congratulations! You've completed Section 6: Standard Library & Web.

Proceed to:
- [Section 7: Advanced & Ecosystem](../../07-Advanced-Ecosystem/) - Learn about advanced Go topics

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

