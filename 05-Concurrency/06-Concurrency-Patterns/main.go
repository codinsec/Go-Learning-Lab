package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// This program demonstrates common concurrency patterns in Go

func main() {
	fmt.Println("=== Concurrency Patterns ===")
	fmt.Println()

	// 1. Worker Pool
	fmt.Println("1. Worker Pool Pattern:")
	workerPool()
	fmt.Println()

	// 2. Fan-In (Multiplexing)
	fmt.Println("2. Fan-In Pattern:")
	fanIn()
	fmt.Println()

	// 3. Fan-Out (Demultiplexing)
	fmt.Println("3. Fan-Out Pattern:")
	fanOut()
	fmt.Println()

	// 4. Pipeline
	fmt.Println("4. Pipeline Pattern:")
	pipeline()
	fmt.Println()

	// 5. Or-Channel
	fmt.Println("5. Or-Channel Pattern:")
	orChannel()
	fmt.Println()

	// 6. Generator
	fmt.Println("6. Generator Pattern:")
	generator()
	fmt.Println()

	// 7. Semaphore
	fmt.Println("7. Semaphore Pattern:")
	semaphore()
	fmt.Println()
}

// Worker Pool: Fixed number of workers processing jobs
func workerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("   Worker %d processing job %d\n", id, job)
				results <- job * 2
			}
		}(w)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= 5; r++ {
		fmt.Printf("   Result: %d\n", <-results)
	}
}

// Fan-In: Combine multiple channels into one
func fanIn() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	// Send values
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(ch1)
	}()
	
	go func() {
		for i := 4; i <= 6; i++ {
			ch2 <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(ch2)
	}()
	
	go func() {
		for i := 7; i <= 9; i++ {
			ch3 <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(ch3)
	}()
	
	// Fan-in
	out := make(chan int)
	var wg sync.WaitGroup
	
	for _, ch := range []chan int{ch1, ch2, ch3} {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	
	go func() {
		wg.Wait()
		close(out)
	}()
	
	// Receive from combined channel
	for val := range out {
		fmt.Printf("   Received: %d\n", val)
	}
}

// Fan-Out: Distribute work to multiple workers
func fanOut() {
	input := make(chan int)
	
	// Send input
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Fan-out to 3 workers
	output1 := make(chan int)
	output2 := make(chan int)
	output3 := make(chan int)
	
	go func() {
		for val := range input {
			select {
			case output1 <- val:
			case output2 <- val:
			case output3 <- val:
			}
		}
		close(output1)
		close(output2)
		close(output3)
	}()
	
	// Process outputs
	time.Sleep(200 * time.Millisecond)
	fmt.Println("   Fan-out completed")
}

// Pipeline: Chain of stages processing data
func pipeline() {
	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	
	// Stage 2: Square numbers
	squares := make(chan int)
	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()
	
	// Stage 3: Print results
	for sq := range squares {
		fmt.Printf("   Square: %d\n", sq)
	}
}

// Or-Channel: Returns when any channel receives a value
func orChannel() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	
	start := time.Now()
	<-or(
		sig(200*time.Millisecond),
		sig(300*time.Millisecond),
		sig(100*time.Millisecond),
	)
	
	fmt.Printf("   Done after %v\n", time.Since(start))
}

// Or function combines multiple channels
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

// Generator: Function that returns a channel
func generator() {
	gen := func(count int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 1; i <= count; i++ {
				ch <- i
				time.Sleep(50 * time.Millisecond)
			}
		}()
		return ch
	}
	
	for val := range gen(5) {
		fmt.Printf("   Generated: %d\n", val)
	}
}

// Semaphore: Limit concurrent operations
func semaphore() {
	sem := make(chan struct{}, 3) // Allow 3 concurrent operations
	var wg sync.WaitGroup
	
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{} // Acquire
			defer func() { <-sem }() // Release
			
			fmt.Printf("   Operation %d running\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("   Operation %d done\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("   All operations completed")
}

