package main

import (
	"sync"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	// Start workers
	for w := 1; w <= 2; w++ {
		go func() {
			for job := range jobs {
				results <- job * 2
			}
		}()
	}
	
	// Send jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	received := 0
	for r := 1; r <= 3; r++ {
		val := <-results
		if val != r*2 {
			t.Errorf("Expected %d, got %d", r*2, val)
		}
		received++
	}
	
	if received != 3 {
		t.Errorf("Expected 3 results, got %d", received)
	}
}

func TestPipeline(t *testing.T) {
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	
	squares := make(chan int)
	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()
	
	expected := []int{1, 4, 9}
	received := 0
	for sq := range squares {
		if sq != expected[received] {
			t.Errorf("Expected %d, got %d", expected[received], sq)
		}
		received++
	}
	
	if received != 3 {
		t.Errorf("Expected 3 values, got %d", received)
	}
}

func TestGenerator(t *testing.T) {
	gen := func(count int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 1; i <= count; i++ {
				ch <- i
			}
		}()
		return ch
	}
	
	count := 0
	for val := range gen(5) {
		count++
		if val != count {
			t.Errorf("Expected %d, got %d", count, val)
		}
	}
	
	if count != 5 {
		t.Errorf("Expected 5 values, got %d", count)
	}
}

func TestSemaphore(t *testing.T) {
	sem := make(chan struct{}, 2) // Allow 2 concurrent
	var wg sync.WaitGroup
	var mu sync.Mutex
	running := 0
	maxRunning := 0
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			
			mu.Lock()
			running++
			if running > maxRunning {
				maxRunning = running
			}
			mu.Unlock()
			
			time.Sleep(10 * time.Millisecond)
			
			mu.Lock()
			running--
			mu.Unlock()
		}()
	}
	
	wg.Wait()
	
	if maxRunning > 2 {
		t.Errorf("Expected max 2 concurrent, got %d", maxRunning)
	}
}

