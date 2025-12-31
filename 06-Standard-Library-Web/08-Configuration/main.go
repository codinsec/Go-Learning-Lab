package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// This program demonstrates configuration management in Go

// Config struct
type Config struct {
	Host     string
	Port     int
	Database string
	Debug    bool
}

func main() {
	fmt.Println("=== Configuration ===")
	fmt.Println()

	// 1. Environment variables
	fmt.Println("1. Environment Variables:")
	envConfig()
	fmt.Println()

	// 2. Command-line flags
	fmt.Println("2. Command-Line Flags:")
	flagConfig()
	fmt.Println()

	// 3. Config file (JSON)
	fmt.Println("3. Config File (JSON):")
	fileConfig()
	fmt.Println()

	// 4. Combined approach
	fmt.Println("4. Combined Approach:")
	combinedConfig()
	fmt.Println()

	// 5. Best practices
	fmt.Println("5. Best Practices:")
	fmt.Println("   - Use environment variables for secrets")
	fmt.Println("   - Use flags for command-line options")
	fmt.Println("   - Use config files for complex settings")
	fmt.Println("   - Environment > Flags > Config file (priority)")
	fmt.Println("   - Never commit secrets to code")
}

func envConfig() {
	// Set environment variable (for demo)
	os.Setenv("APP_HOST", "localhost")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("APP_DEBUG", "true")
	
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	debug := os.Getenv("APP_DEBUG")
	
	fmt.Printf("   Host: %s\n", host)
	fmt.Printf("   Port: %s\n", port)
	fmt.Printf("   Debug: %s\n", debug)
	
	// Get with default
	database := os.Getenv("APP_DATABASE")
	if database == "" {
		database = "default.db" // Default value
	}
	fmt.Printf("   Database: %s (with default)\n", database)
}

func flagConfig() {
	// Define flags
	host := flag.String("host", "localhost", "Server host")
	port := flag.Int("port", 8080, "Server port")
	debug := flag.Bool("debug", false, "Enable debug mode")
	
	// Parse flags (in real app, call flag.Parse() in main)
	// For demo, we'll just show the definition
	fmt.Printf("   Host flag: %s (default: localhost)\n", *host)
	fmt.Printf("   Port flag: %d (default: 8080)\n", *port)
	fmt.Printf("   Debug flag: %t (default: false)\n", *debug)
	fmt.Println("   Usage: go run main.go -host=0.0.0.0 -port=3000 -debug")
}

func fileConfig() {
	config := Config{
		Host:     "localhost",
		Port:     8080,
		Database: "app.db",
		Debug:    false,
	}
	
	// Write config to file
	configFile := "config.json"
	jsonData, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(configFile, jsonData, 0644)
	defer os.Remove(configFile)
	
	fmt.Printf("   Config written to: %s\n", configFile)
	
	// Read config from file
	var loadedConfig Config
	data, err := os.ReadFile(configFile)
	if err == nil {
		json.Unmarshal(data, &loadedConfig)
		fmt.Printf("   Loaded config: %+v\n", loadedConfig)
	}
}

func combinedConfig() {
	// Priority: Environment > Flags > Config file > Defaults
	
	// Defaults
	config := Config{
		Host:     "localhost",
		Port:     8080,
		Database: "app.db",
		Debug:    false,
	}
	
	// Override with environment variables
	if host := os.Getenv("APP_HOST"); host != "" {
		config.Host = host
	}
	if port := os.Getenv("APP_PORT"); port != "" {
		fmt.Sscanf(port, "%d", &config.Port)
	}
	
	fmt.Printf("   Final config: %+v\n", config)
	fmt.Println("   (Environment variables override defaults)")
}

