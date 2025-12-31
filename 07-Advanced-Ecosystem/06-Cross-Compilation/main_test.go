package main

import (
	"runtime"
	"testing"
)

func TestCurrentPlatform(t *testing.T) {
	os := runtime.GOOS
	if os == "" {
		t.Error("OS should not be empty")
	}
	
	arch := runtime.GOARCH
	if arch == "" {
		t.Error("Architecture should not be empty")
	}
}

func TestPlatformValues(t *testing.T) {
	os := runtime.GOOS
	validOS := map[string]bool{
		"linux":   true,
		"windows": true,
		"darwin":  true,
		"freebsd": true,
		"openbsd": true,
		"netbsd":  true,
	}
	
	// Check if current OS is in valid list (or allow others)
	if !validOS[os] && os != "" {
		// Allow other valid OSes
		t.Logf("OS %s is not in common list, but may be valid", os)
	}
	
	arch := runtime.GOARCH
	validArch := map[string]bool{
		"amd64": true,
		"386":   true,
		"arm":   true,
		"arm64": true,
	}
	
	if !validArch[arch] && arch != "" {
		// Allow other valid architectures
		t.Logf("Architecture %s is not in common list, but may be valid", arch)
	}
}

func TestCrossCompilationInfo(t *testing.T) {
	// Test that we can get platform information
	os := runtime.GOOS
	arch := runtime.GOARCH
	
	if os == "" || arch == "" {
		t.Error("Platform information should be available")
	}
	
	// Verify format
	if len(os) == 0 || len(arch) == 0 {
		t.Error("Platform values should have content")
	}
}

