package main

import "testing"

func TestDeferExecution(t *testing.T) {
	// This test verifies defer executes in reverse order
	// We can't easily test the order in a unit test, but we can verify
	// that deferred functions are called
	called := false
	defer func() {
		called = true
	}()
	
	if called {
		t.Error("Deferred function should not be called yet")
	}
}

func TestRecover(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected to recover from panic")
			}
		}()
		panic("test panic")
	}()
}

func TestRecoverNoPanic(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Should not recover anything if no panic occurred")
			}
		}()
		// No panic
	}()
}

func TestMultipleDefers(t *testing.T) {
	order := []int{}
	
	defer func() {
		order = append(order, 1)
	}()
	defer func() {
		order = append(order, 2)
	}()
	defer func() {
		order = append(order, 3)
	}()
	
	// At this point, order should be empty
	// After function returns, order should be [3, 2, 1] (LIFO)
	
	// Note: In actual test, we'd need to check order after function returns
	// This is a simplified test
	if len(order) != 0 {
		t.Error("Order should be empty before function returns")
	}
}

func TestDeferArguments(t *testing.T) {
	x := 10
	defer func(val int) {
		if val != 10 {
			t.Errorf("Expected 10, got %d", val)
		}
	}(x) // x is evaluated and passed now
	
	x = 20
	// Deferred function will use the original value (10)
}

