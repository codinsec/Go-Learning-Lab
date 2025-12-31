package main

import (
	"os"
	"runtime"
	"testing"
)

func TestGoVersion(t *testing.T) {
	version := runtime.Version()
	if version == "" {
		t.Error("Go version should not be empty")
	}
	t.Logf("Go version: %s", version)
}

func TestGOROOT(t *testing.T) {
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = runtime.GOROOT()
	}
	if goroot == "" {
		t.Error("GOROOT should be set")
	}
	t.Logf("GOROOT: %s", goroot)
}

func TestOSAndArch(t *testing.T) {
	os := runtime.GOOS
	arch := runtime.GOARCH

	if os == "" {
		t.Error("Operating system should not be empty")
	}
	if arch == "" {
		t.Error("Architecture should not be empty")
	}

	t.Logf("OS: %s, Arch: %s", os, arch)
}

func TestNumCPU(t *testing.T) {
	cpus := runtime.NumCPU()
	if cpus <= 0 {
		t.Error("Number of CPUs should be greater than 0")
	}
	t.Logf("Number of CPUs: %d", cpus)
}

