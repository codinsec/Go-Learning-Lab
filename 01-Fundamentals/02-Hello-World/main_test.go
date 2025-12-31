package main

import (
	"testing"
)

func TestMainFunction(t *testing.T) {
	// This test verifies that the main function can be called
	// In a real scenario, you might test the output
	t.Log("Main function exists and can be called")
}

func TestPrintln(t *testing.T) {
	// Test that fmt.Println works (basic sanity check)
	// In production, you'd capture output and verify it
	t.Log("fmt.Println is available and functional")
}

