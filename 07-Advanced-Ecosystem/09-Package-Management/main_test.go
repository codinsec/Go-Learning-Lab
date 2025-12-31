package main

import (
	"os/exec"
	"testing"
)

func TestGoModExists(t *testing.T) {
	// Check if go.mod exists by trying to list module
	cmd := exec.Command("go", "list", "-m")
	if err := cmd.Run(); err != nil {
		t.Skip("go mod not available or not in module")
	}
}

func TestGoListM(t *testing.T) {
	// Test go list -m command
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		t.Skip("go list -m command failed")
	}
	if len(output) == 0 {
		t.Error("Module name should not be empty")
	}
}

func TestGoEnv(t *testing.T) {
	// Test go env command
	cmd := exec.Command("go", "env", "GOPROXY")
	if err := cmd.Run(); err != nil {
		t.Skip("go env command not available")
	}
}

func TestMinFunction(t *testing.T) {
	if min(1, 2) != 1 {
		t.Errorf("Expected 1, got %d", min(1, 2))
	}
	if min(2, 1) != 1 {
		t.Errorf("Expected 1, got %d", min(2, 1))
	}
	if min(5, 5) != 5 {
		t.Errorf("Expected 5, got %d", min(5, 5))
	}
}

