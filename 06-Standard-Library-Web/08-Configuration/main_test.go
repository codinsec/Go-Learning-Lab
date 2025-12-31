package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEnvConfig(t *testing.T) {
	os.Setenv("TEST_HOST", "testhost")
	host := os.Getenv("TEST_HOST")
	if host != "testhost" {
		t.Errorf("Expected 'testhost', got %s", host)
	}
	
	os.Unsetenv("TEST_HOST")
}

func TestFileConfig(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     8080,
		Database: "test.db",
		Debug:    true,
	}
	
	configFile := "test_config.json"
	jsonData, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	
	err = os.WriteFile(configFile, jsonData, 0644)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	defer os.Remove(configFile)
	
	var loadedConfig Config
	data, err := os.ReadFile(configFile)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	
	err = json.Unmarshal(data, &loadedConfig)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	
	if loadedConfig.Host != config.Host {
		t.Errorf("Expected host %s, got %s", config.Host, loadedConfig.Host)
	}
	if loadedConfig.Port != config.Port {
		t.Errorf("Expected port %d, got %d", config.Port, loadedConfig.Port)
	}
}

func TestCombinedConfig(t *testing.T) {
	os.Setenv("TEST_APP_HOST", "envhost")
	defer os.Unsetenv("TEST_APP_HOST")
	
	config := Config{
		Host: "default",
	}
	
	if host := os.Getenv("TEST_APP_HOST"); host != "" {
		config.Host = host
	}
	
	if config.Host != "envhost" {
		t.Errorf("Expected 'envhost', got %s", config.Host)
	}
}

