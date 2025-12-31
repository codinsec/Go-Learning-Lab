package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

// This program demonstrates logging in Go

func main() {
	fmt.Println("=== Logging ===")
	fmt.Println()

	// 1. Basic logging with log package
	fmt.Println("1. Basic Logging (log package):")
	log.Println("This is a log message")
	log.Printf("Formatted log: %s", "value")
	fmt.Println()

	// 2. Log levels (using log package)
	fmt.Println("2. Log Levels (log package):")
	log.SetPrefix("[INFO] ")
	log.Println("Info message")
	
	log.SetPrefix("[ERROR] ")
	log.Println("Error message")
	
	log.SetPrefix("") // Reset
	fmt.Println()

	// 3. Log to file
	fmt.Println("3. Log to File:")
	file, err := os.CreateTemp("", "log-*.txt")
	if err == nil {
		defer os.Remove(file.Name())
		log.SetOutput(file)
		log.Println("This goes to file")
		log.SetOutput(os.Stdout) // Reset to stdout
		fmt.Printf("   Logged to file: %s\n", file.Name())
	}
	fmt.Println()

	// 4. Structured logging with slog (Go 1.21+)
	fmt.Println("4. Structured Logging (slog package):")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("User logged in", "user_id", 123, "ip", "192.168.1.1")
	logger.Error("Failed to connect", "error", "connection timeout")
	fmt.Println()

	// 5. JSON logging with slog
	fmt.Println("5. JSON Logging (slog):")
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("Application started", "version", "1.0.0", "port", 8080)
	fmt.Println()

	// 6. Log levels with slog
	fmt.Println("6. Log Levels (slog):")
	logger.Debug("Debug message (may not appear)")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
	fmt.Println()

	// 7. Custom logger
	fmt.Println("7. Custom Logger:")
	customLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		AddSource: true,
	}))
	customLogger.Info("Message with source location")
	fmt.Println()

	// 8. Logger with context
	fmt.Println("8. Logger with Context:")
	logger = logger.With("request_id", "req-123", "user", "alice")
	logger.Info("Processing request")
	logger.Info("Request completed")
	fmt.Println()

	// 9. Best practices
	fmt.Println("9. Best Practices:")
	fmt.Println("   - Use structured logging (slog) for production")
	fmt.Println("   - Include context (request ID, user ID)")
	fmt.Println("   - Use appropriate log levels")
	fmt.Println("   - Don't log sensitive information")
	fmt.Println("   - Use JSON format for log aggregation")
	fmt.Println("   - Set log levels per environment")
}

