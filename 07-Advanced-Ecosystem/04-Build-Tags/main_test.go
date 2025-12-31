package main

import (
	"runtime"
	"testing"
)

func TestPlatformInfo(t *testing.T) {
	info := getPlatformInfo()
	if info == "" {
		t.Error("Platform info should not be empty")
	}
	
	// Verify platform-specific implementation
	switch runtime.GOOS {
	case "linux":
		if info != "Running on Linux" {
			t.Errorf("Expected Linux message, got %s", info)
		}
	case "windows":
		if info != "Running on Windows" {
			t.Errorf("Expected Windows message, got %s", info)
		}
	case "darwin":
		if info != "Running on macOS (Darwin)" {
			t.Errorf("Expected macOS message, got %s", info)
		}
	}
}

func TestPlatformDetection(t *testing.T) {
	os := runtime.GOOS
	if os == "" {
		t.Error("OS should not be empty")
	}
	
	arch := runtime.GOARCH
	if arch == "" {
		t.Error("Architecture should not be empty")
	}
}

