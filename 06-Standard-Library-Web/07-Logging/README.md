# Logging

Learn about logging in Go using the `log` and `log/slog` packages.

## Learning Objectives

- Understand basic logging with `log` package
- Learn about structured logging with `slog`
- Master log levels
- Understand JSON logging
- Learn about logging best practices
- Understand log configuration

## Concepts Covered

### Basic Logging (log package)

```go
import "log"

log.Println("Message")
log.Printf("Formatted: %s", value)
log.Fatal("Fatal error") // Exits program
log.Panic("Panic")       // Panics
```

### Structured Logging (slog)

```go
import "log/slog"

logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
logger.Info("User logged in", "user_id", 123, "ip", "192.168.1.1")
```

### JSON Logging

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("Event", "key", "value")
```

### Log Levels

```go
logger.Debug("Debug message")
logger.Info("Info message")
logger.Warn("Warning message")
logger.Error("Error message")
```

### Logger with Context

```go
logger = logger.With("request_id", "req-123")
logger.Info("Message") // Includes request_id
```

### Custom Logger

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
    AddSource: true,
}))
```

## Running the Example

```bash
# Navigate to this directory
cd 07-Logging

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/06-standard-library-web/logging

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

1. **log package** - Basic logging
2. **slog package** - Structured logging (Go 1.21+)
3. **Structured logging** - Key-value pairs
4. **JSON format** - For log aggregation
5. **Log levels** - Debug, Info, Warn, Error
6. **Context** - Add fields to logger

## Common Patterns

**Basic logging:**
```go
log.Println("Message")
```

**Structured logging:**
```go
logger.Info("Event", "key", "value")
```

**Logger with context:**
```go
logger = logger.With("request_id", id)
logger.Info("Processing")
```

## Best Practices

- **Use structured logging** - slog for production
- **Include context** - Request ID, user ID, etc.
- **Use appropriate levels** - Debug, Info, Warn, Error
- **Don't log sensitive data** - Passwords, tokens, etc.
- **Use JSON format** - For log aggregation tools
- **Set log levels** - Per environment (dev vs prod)

## Important Notes

- **slog is new** - Available in Go 1.21+
- **log.Fatal** - Exits program (use carefully)
- **log.Panic** - Panics (use carefully)
- **JSON format** - Better for log aggregation

## Next Steps

After understanding logging, proceed to:
- [08-Configuration](../08-Configuration/README.md) - Learn about configuration management

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

