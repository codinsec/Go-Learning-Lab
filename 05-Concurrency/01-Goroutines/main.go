package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// This program demonstrates goroutines in Go

// 1. Basic goroutine
func sayHello(name string) {
	fmt.Printf("Hello from goroutine: %s\n", name)
}

// 2. Goroutine with sleep
func countNumbers(name string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// 3. Anonymous function goroutine
func anonymousGoroutine() {
	go func() {
		fmt.Println("This is an anonymous goroutine")
	}()
	time.Sleep(100 * time.Millisecond) // Give goroutine time to execute
}

// 4. Multiple goroutines
func multipleGoroutines() {
	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d is running\n", id)
		}(i)
	}
	time.Sleep(500 * time.Millisecond) // Wait for goroutines
}

// 5. Goroutine with closure (capturing variables)
func goroutineClosure() {
	message := "Hello from closure"
	go func() {
		fmt.Println(message) // Captures message from outer scope
	}()
	time.Sleep(100 * time.Millisecond)
}

// 6. Goroutine lifecycle
func goroutineLifecycle() {
	fmt.Println("Main goroutine started")
	
	go func() {
		fmt.Println("Goroutine started")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Goroutine finished")
	}()
	
	fmt.Println("Main goroutine continues")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Main goroutine finished")
}

// 7. Goroutines and CPU cores
func showGoroutinesAndCores() {
	fmt.Printf("Number of CPU cores: %d\n", runtime.NumCPU())
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	
	for i := 0; i < 10; i++ {
		go func(id int) {
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	fmt.Printf("Number of goroutines after spawning: %d\n", runtime.NumGoroutine())
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Number of goroutines after sleep: %d\n", runtime.NumGoroutine())
}

// 8. Goroutine synchronization with WaitGroup (preview)
func goroutineWithWaitGroup() {
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter
		go func(id int) {
			defer wg.Done() // Decrement counter when done
			fmt.Printf("Goroutine %d is working\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}
	
	fmt.Println("Waiting for all goroutines to finish...")
	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All goroutines finished!")
}

// 9. Goroutine with return value (using channel - preview)
func goroutineWithReturn() {
	resultChan := make(chan int)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		resultChan <- 42 // Send result
	}()
	
	result := <-resultChan // Receive result
	fmt.Printf("Result from goroutine: %d\n", result)
}

// 10. Common mistakes
func commonMistakes() {
	fmt.Println("Common mistakes with goroutines:")
	fmt.Println("1. Not waiting for goroutines (use WaitGroup or channels)")
	fmt.Println("2. Capturing loop variables incorrectly")
	fmt.Println("3. Not handling goroutine errors")
	fmt.Println("4. Creating too many goroutines without limits")
}

func main() {
	fmt.Println("=== Goroutines ===")
	fmt.Println()

	// 1. Basic goroutine
	fmt.Println("1. Basic Goroutine:")
	go sayHello("Goroutine 1")
	time.Sleep(100 * time.Millisecond) // Give goroutine time to execute
	fmt.Println()

	// 2. Goroutine with sleep
	fmt.Println("2. Goroutine with Sleep:")
	go countNumbers("Counter", 3)
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// 3. Anonymous function
	fmt.Println("3. Anonymous Function Goroutine:")
	anonymousGoroutine()
	fmt.Println()

	// 4. Multiple goroutines
	fmt.Println("4. Multiple Goroutines:")
	multipleGoroutines()
	fmt.Println()

	// 5. Closure
	fmt.Println("5. Goroutine with Closure:")
	goroutineClosure()
	fmt.Println()

	// 6. Lifecycle
	fmt.Println("6. Goroutine Lifecycle:")
	goroutineLifecycle()
	fmt.Println()

	// 7. CPU cores and goroutines
	fmt.Println("7. Goroutines and CPU Cores:")
	showGoroutinesAndCores()
	fmt.Println()

	// 8. WaitGroup (preview)
	fmt.Println("8. Goroutine with WaitGroup (Preview):")
	goroutineWithWaitGroup()
	fmt.Println()

	// 9. Return value (preview)
	fmt.Println("9. Goroutine with Return Value (Preview):")
	goroutineWithReturn()
	fmt.Println()

	// 10. Common mistakes
	fmt.Println("10. Common Mistakes:")
	commonMistakes()
	fmt.Println()

	// 11. Key concepts
	fmt.Println("11. Key Concepts:")
	fmt.Println("   - Goroutines are lightweight (2KB stack)")
	fmt.Println("   - Managed by Go runtime")
	fmt.Println("   - M:N threading model (many goroutines on few OS threads)")
	fmt.Println("   - Use 'go' keyword to start: go function()")
	fmt.Println("   - Main goroutine must wait for others (or program exits)")
}

