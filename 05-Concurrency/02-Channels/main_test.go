package main

import (
	"testing"
	"time"
)

func TestUnbufferedChannel(t *testing.T) {
	ch := make(chan int)
	
	go func() {
		ch <- 42
	}()
	
	value := <-ch
	if value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 3)
	
	ch <- 1
	ch <- 2
	ch <- 3
	
	if len(ch) != 3 {
		t.Errorf("Expected length 3, got %d", len(ch))
	}
	
	if cap(ch) != 3 {
		t.Errorf("Expected capacity 3, got %d", cap(ch))
	}
	
	val1 := <-ch
	val2 := <-ch
	val3 := <-ch
	
	if val1 != 1 || val2 != 2 || val3 != 3 {
		t.Error("Values should be 1, 2, 3")
	}
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)
	
	val1, ok1 := <-ch
	val2, ok2 := <-ch
	_, ok3 := <-ch
	
	if !ok1 || val1 != 1 {
		t.Error("First value should be 1")
	}
	if !ok2 || val2 != 2 {
		t.Error("Second value should be 2")
	}
	if ok3 {
		t.Error("Third receive should return false")
	}
}

func TestRangeOverChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	count := 0
	sum := 0
	for val := range ch {
		count++
		sum += val
	}
	
	if count != 5 {
		t.Errorf("Expected 5 values, got %d", count)
	}
	if sum != 15 {
		t.Errorf("Expected sum 15, got %d", sum)
	}
}

func TestChannelDirection(t *testing.T) {
	ch := make(chan int)
	
	go func(sendCh chan<- int) {
		sendCh <- 42
	}(ch)
	
	val := <-ch
	if val != 42 {
		t.Errorf("Expected 42, got %d", val)
	}
}

func TestMultipleGoroutinesOneChannel(t *testing.T) {
	ch := make(chan int)
	expected := []int{1, 2, 3}
	received := make([]int, 0, 3)
	
	for i := 1; i <= 3; i++ {
		go func(id int) {
			ch <- id
		}(i)
	}
	
	for i := 0; i < 3; i++ {
		received = append(received, <-ch)
	}
	
	if len(received) != 3 {
		t.Errorf("Expected 3 values, got %d", len(received))
	}
	
	// Check that all expected values are present
	found := make(map[int]bool)
	for _, val := range received {
		found[val] = true
	}
	for _, val := range expected {
		if !found[val] {
			t.Errorf("Expected value %d not found", val)
		}
	}
}

func TestChannelBlocking(t *testing.T) {
	ch := make(chan int)
	done := make(chan bool)
	
	go func() {
		ch <- 100
		done <- true
	}()
	
	// Give goroutine time to block
	time.Sleep(10 * time.Millisecond)
	
	select {
	case val := <-ch:
		if val != 100 {
			t.Errorf("Expected 100, got %d", val)
		}
		<-done // Wait for goroutine to finish
	case <-time.After(100 * time.Millisecond):
		t.Error("Channel should have received value")
	}
}

func TestNilChannel(t *testing.T) {
	var ch chan int
	if ch == nil {
		t.Log("Nil channel created successfully")
	}
}

