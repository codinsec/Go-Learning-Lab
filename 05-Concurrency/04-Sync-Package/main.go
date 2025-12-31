package main

import (
	"fmt"
	"sync"
	"time"
)

// This program demonstrates the sync package in Go

var (
	counter int
	mu      sync.Mutex
)

func main() {
	fmt.Println("=== Sync Package ===")
	fmt.Println()

	// 1. WaitGroup
	fmt.Println("1. WaitGroup:")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("   Goroutine %d is working\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("   All goroutines finished!")
	fmt.Println()

	// 2. Mutex
	fmt.Println("2. Mutex:")
	counter = 0
	var mu sync.Mutex
	
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("   Counter (with mutex): %d\n", counter)
	fmt.Println()

	// 3. RWMutex
	fmt.Println("3. RWMutex:")
	var rwmu sync.RWMutex
	data := make(map[string]int)
	
	// Writers
	for i := 0; i < 3; i++ {
		go func(id int) {
			rwmu.Lock()
			data[fmt.Sprintf("key%d", id)] = id
			rwmu.Unlock()
		}(i)
	}
	
	// Readers
	for i := 0; i < 5; i++ {
		go func() {
			rwmu.RLock()
			_ = len(data)
			rwmu.RUnlock()
		}()
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("   Data map size: %d\n", len(data))
	fmt.Println()

	// 4. Once
	fmt.Println("4. Once:")
	var once sync.Once
	initialized := false
	
	initFunc := func() {
		initialized = true
		fmt.Println("   Initialized!")
	}
	
	for i := 0; i < 5; i++ {
		go once.Do(initFunc)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("   Initialized: %t\n", initialized)
	fmt.Println()

	// 5. Cond (Condition Variable)
	fmt.Println("5. Cond (Condition Variable):")
	var m sync.Mutex
	cond := sync.NewCond(&m)
	ready := false
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		m.Lock()
		ready = true
		cond.Signal() // Wake one waiting goroutine
		m.Unlock()
	}()
	
	m.Lock()
	for !ready {
		cond.Wait() // Wait for condition
	}
	fmt.Println("   Condition met!")
	m.Unlock()
	fmt.Println()

	// 6. Pool
	fmt.Println("6. Pool:")
	var pool sync.Pool
	
	pool.New = func() interface{} {
		return make([]byte, 1024)
	}
	
	// Get from pool
	buf1 := pool.Get().([]byte)
	fmt.Printf("   Got buffer from pool: len=%d\n", len(buf1))
	
	// Put back to pool
	pool.Put(buf1)
	
	// Get again (reused)
	buf2 := pool.Get().([]byte)
	fmt.Printf("   Got buffer again: len=%d\n", len(buf2))
	fmt.Println()

	// 7. Practical example: Counter with mutex
	fmt.Println("7. Practical Example - Thread-Safe Counter:")
	safeCounter := NewCounter()
	
	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			safeCounter.Increment()
		}()
	}
	
	wg2.Wait()
	fmt.Printf("   Final count: %d\n", safeCounter.Value())
	fmt.Println()

	// 8. Key differences
	fmt.Println("8. Key Differences:")
	fmt.Println("   Mutex:")
	fmt.Println("     - Exclusive lock (one at a time)")
	fmt.Println("     - Use for write operations")
	fmt.Println("   RWMutex:")
	fmt.Println("     - Multiple readers OR one writer")
	fmt.Println("     - Use when reads > writes")
	fmt.Println("   Once:")
	fmt.Println("     - Execute function exactly once")
	fmt.Println("     - Use for initialization")
	fmt.Println("   Cond:")
	fmt.Println("     - Wait for condition")
	fmt.Println("     - Use for signaling between goroutines")
	fmt.Println("   Pool:")
	fmt.Println("     - Object pool for reuse")
	fmt.Println("     - Use to reduce allocations")
}

// Thread-safe counter
type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

