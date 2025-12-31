package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	count := 0
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	
	wg.Wait()
	
	if count != 10 {
		t.Errorf("Expected count 10, got %d", count)
	}
}

func TestMutex(t *testing.T) {
	var mu sync.Mutex
	counter := 0
	
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	
	wg.Wait()
	
	if counter != 100 {
		t.Errorf("Expected counter 100, got %d", counter)
	}
}

func TestRWMutex(t *testing.T) {
	var rwmu sync.RWMutex
	data := make(map[string]int)
	
	// Writers
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rwmu.Lock()
			data[fmt.Sprintf("key%d", id)] = id
			rwmu.Unlock()
		}(i)
	}
	
	wg.Wait()
	
	if len(data) != 10 {
		t.Errorf("Expected 10 entries, got %d", len(data))
	}
}

func TestOnce(t *testing.T) {
	var once sync.Once
	count := 0
	
	initFunc := func() {
		count++
	}
	
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initFunc)
		}()
	}
	
	wg.Wait()
	
	if count != 1 {
		t.Errorf("Expected count 1, got %d", count)
	}
}

func TestCounter(t *testing.T) {
	counter := NewCounter()
	
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	
	wg.Wait()
	
	if counter.Value() != 100 {
		t.Errorf("Expected 100, got %d", counter.Value())
	}
}

