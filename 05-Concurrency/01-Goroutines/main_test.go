package main

import (
	"sync"
	"testing"
	"time"
)

func TestGoroutineExecution(t *testing.T) {
	var wg sync.WaitGroup
	executed := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		executed = true
	}()

	wg.Wait()

	if !executed {
		t.Error("Goroutine should have executed")
	}
}

func TestMultipleGoroutines(t *testing.T) {
	var wg sync.WaitGroup
	count := 0
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if count != 10 {
		t.Errorf("Expected count 10, got %d", count)
	}
}

func TestGoroutineClosure(t *testing.T) {
	var wg sync.WaitGroup
	message := "test message"
	captured := ""

	wg.Add(1)
	go func() {
		defer wg.Done()
		captured = message
	}()

	wg.Wait()

	if captured != message {
		t.Errorf("Expected %s, got %s", message, captured)
	}
}

func TestGoroutineWithParameter(t *testing.T) {
	var wg sync.WaitGroup
	results := make([]int, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			results[id] = id
		}(i)
	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		if results[i] != i {
			t.Errorf("Expected %d, got %d", i, results[i])
		}
	}
}

func TestGoroutineLifecycle(t *testing.T) {
	var wg sync.WaitGroup
	started := false
	finished := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		started = true
		time.Sleep(10 * time.Millisecond)
		finished = true
	}()

	if finished {
		t.Error("Goroutine should not have finished yet")
	}

	wg.Wait()

	if !started {
		t.Error("Goroutine should have started")
	}
	if !finished {
		t.Error("Goroutine should have finished")
	}
}

func BenchmarkGoroutineCreation(b *testing.B) {
	var wg sync.WaitGroup
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}

	wg.Wait()
}

