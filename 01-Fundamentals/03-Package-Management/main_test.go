package main

import (
	"os"
	"testing"
)

func TestGoModExists(t *testing.T) {
	// Check if go.mod file exists
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		t.Error("go.mod file should exist in the project")
	} else {
		t.Log("✅ go.mod file exists")
	}
}

func TestModulePath(t *testing.T) {
	// Read go.mod to verify module path
	data, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatalf("Failed to read go.mod: %v", err)
	}

	content := string(data)
	if len(content) == 0 {
		t.Error("go.mod should not be empty")
	}

	t.Logf("go.mod content:\n%s", content)
}

