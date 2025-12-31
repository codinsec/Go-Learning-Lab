package main

import (
	"os/exec"
	"runtime"
	"testing"
)

func TestGoVersion(t *testing.T) {
	// Check if go command is available
	cmd := exec.Command("go", "version")
	if err := cmd.Run(); err != nil {
		t.Skip("go command not available")
	}
}

func TestRuntimeInfo(t *testing.T) {
	os := runtime.GOOS
	if os == "" {
		t.Error("OS should not be empty")
	}

	arch := runtime.GOARCH
	if arch == "" {
		t.Error("Architecture should not be empty")
	}

	version := runtime.Version()
	if version == "" {
		t.Error("Go version should not be empty")
	}
}

func TestGoList(t *testing.T) {
	// Try to list packages
	cmd := exec.Command("go", "list", ".")
	if err := cmd.Run(); err != nil {
		t.Skip("go list command not available or failed")
	}
}

func TestGoEnv(t *testing.T) {
	// Try to get Go environment
	cmd := exec.Command("go", "env", "GOPATH")
	if err := cmd.Run(); err != nil {
		t.Skip("go env command not available or failed")
	}
}

