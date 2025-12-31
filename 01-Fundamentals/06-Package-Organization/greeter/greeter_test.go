package greeter

import "testing"

func TestSayHello(t *testing.T) {
	// This test verifies the function exists and can be called
	// In a real scenario, you might capture output
	SayHello("Test")
	t.Log("SayHello function works")
}

func TestSayGoodbye(t *testing.T) {
	SayGoodbye("Test")
	t.Log("SayGoodbye function works")
}

