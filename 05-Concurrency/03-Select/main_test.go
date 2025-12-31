package main

import (
	"testing"
	"time"
)

func TestSelectBasic(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		ch1 <- 1
	}()
	
	select {
	case val := <-ch1:
		if val != 1 {
			t.Errorf("Expected 1, got %d", val)
		}
	case <-ch2:
		t.Error("Should not receive from ch2")
	}
}

func TestSelectDefault(t *testing.T) {
	ch := make(chan int)
	received := false
	
	select {
	case <-ch:
		received = true
	default:
		// Expected path
	}
	
	if received {
		t.Error("Should not have received from empty channel")
	}
}

func TestSelectTimeout(t *testing.T) {
	ch := make(chan int)
	timeout := false
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- 1
	}()
	
	select {
	case <-ch:
		// Should timeout before receiving
	case <-time.After(50 * time.Millisecond):
		timeout = true
	}
	
	if !timeout {
		t.Error("Should have timed out")
	}
}

func TestSelectMultipleCases(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		ch1 <- 1
	}()
	
	received := false
	select {
	case val := <-ch1:
		if val == 1 {
			received = true
		}
	case <-ch2:
		// Not expected
	}
	
	if !received {
		t.Error("Should have received from ch1")
	}
}

func TestSelectSending(t *testing.T) {
	ch := make(chan int, 1)
	sent := false
	
	select {
	case ch <- 42:
		sent = true
	default:
		// Not expected
	}
	
	if !sent {
		t.Error("Should have sent to buffered channel")
	}
	
	// Channel is full now
	select {
	case ch <- 43:
		t.Error("Should not be able to send to full channel")
	default:
		// Expected
	}
}

func TestSelectNilChannel(t *testing.T) {
	var ch chan int
	
	select {
	case <-ch:
		t.Error("Should not receive from nil channel")
	default:
		// Expected - nil channels are never ready
	}
}

