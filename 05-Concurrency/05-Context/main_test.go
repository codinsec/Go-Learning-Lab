package main

import (
	"context"
	"testing"
	"time"
)

func TestWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()
	
	select {
	case <-ctx.Done():
		if ctx.Err() != context.Canceled {
			t.Errorf("Expected Canceled, got %v", ctx.Err())
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Context should have been cancelled")
	}
}

func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	
	select {
	case <-ctx.Done():
		if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("Expected DeadlineExceeded, got %v", ctx.Err())
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Context should have timed out")
	}
}

func TestWithDeadline(t *testing.T) {
	deadline := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	select {
	case <-ctx.Done():
		if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("Expected DeadlineExceeded, got %v", ctx.Err())
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Context should have exceeded deadline")
	}
}

func TestWithValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	
	val := ctx.Value("key")
	if val != "value" {
		t.Errorf("Expected 'value', got %v", val)
	}
	
	val2 := ctx.Value("nonexistent")
	if val2 != nil {
		t.Errorf("Expected nil, got %v", val2)
	}
}

func TestDoWork(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()
	
	result := doWork(ctx)
	if result == "" {
		t.Error("Result should not be empty")
	}
}

func TestContextPropagation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()
	
	select {
	case <-ctx.Done():
		// Expected
	case <-time.After(100 * time.Millisecond):
		t.Error("Context should have been cancelled")
	}
}

