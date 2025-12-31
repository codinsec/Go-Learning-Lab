package main

import (
	"fmt"
	"time"
)

// This program demonstrates channels in Go

func main() {
	fmt.Println("=== Channels ===")
	fmt.Println()

	// 1. Creating channels
	fmt.Println("1. Creating Channels:")
	ch1 := make(chan int)           // Unbuffered channel
	ch2 := make(chan string, 3)     // Buffered channel (capacity 3)
	fmt.Printf("   ch1: unbuffered channel\n")
	fmt.Printf("   ch2: buffered channel (capacity 3)\n")
	fmt.Println()

	// 2. Sending and receiving (unbuffered)
	fmt.Println("2. Unbuffered Channel (Synchronous):")
	go func() {
		ch1 <- 42 // Send value
		fmt.Println("   Sent: 42")
	}()
	
	value := <-ch1 // Receive value
	fmt.Printf("   Received: %d\n", value)
	fmt.Println()

	// 3. Buffered channel
	fmt.Println("3. Buffered Channel (Asynchronous):")
	ch2 <- "first"
	ch2 <- "second"
	ch2 <- "third"
	fmt.Println("   Sent 3 values to buffered channel")
	
	fmt.Printf("   Received: %s\n", <-ch2)
	fmt.Printf("   Received: %s\n", <-ch2)
	fmt.Printf("   Received: %s\n", <-ch2)
	fmt.Println()

	// 4. Channel directions
	fmt.Println("4. Channel Directions:")
	sendOnly := make(chan<- int)    // Send-only channel
	receiveOnly := make(<-chan int) // Receive-only channel
	fmt.Printf("   sendOnly: %T\n", sendOnly)
	fmt.Printf("   receiveOnly: %T\n", receiveOnly)
	fmt.Println()

	// 5. Closing channels
	fmt.Println("5. Closing Channels:")
	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	close(ch3) // Close channel
	
	// Receive remaining values
	for val := range ch3 {
		fmt.Printf("   Received: %d\n", val)
	}
	fmt.Println()

	// 6. Range over channel
	fmt.Println("6. Range Over Channel:")
	ch4 := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch4 <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(ch4)
	}()
	
	fmt.Println("   Receiving values:")
	for val := range ch4 {
		fmt.Printf("     %d\n", val)
	}
	fmt.Println()

	// 7. Channel as function parameter
	fmt.Println("7. Channel as Function Parameter:")
	ch5 := make(chan string)
	go sendMessage(ch5, "Hello from function")
	msg := <-ch5
	fmt.Printf("   Received: %s\n", msg)
	fmt.Println()

	// 8. Multiple goroutines, one channel
	fmt.Println("8. Multiple Goroutines, One Channel:")
	ch6 := make(chan int)
	
	for i := 1; i <= 3; i++ {
		go func(id int) {
			ch6 <- id
		}(i)
	}
	
	for i := 0; i < 3; i++ {
		fmt.Printf("   Received: %d\n", <-ch6)
	}
	fmt.Println()

	// 9. Channel blocking behavior
	fmt.Println("9. Channel Blocking Behavior:")
	unbuffered := make(chan int)
	
	go func() {
		fmt.Println("   Goroutine: About to send (will block until received)")
		unbuffered <- 100
		fmt.Println("   Goroutine: Sent successfully")
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("   Main: About to receive")
	val := <-unbuffered
	fmt.Printf("   Main: Received %d\n", val)
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// 10. Nil channels
	fmt.Println("10. Nil Channels:")
	var nilChan chan int
	fmt.Printf("   nilChan == nil: %t\n", nilChan == nil)
	// Sending/receiving on nil channel blocks forever
	fmt.Println("   Warning: Operations on nil channel block forever!")
	fmt.Println()

	// 11. Channel capacity and length
	fmt.Println("11. Channel Capacity and Length:")
	buffered := make(chan int, 5)
	buffered <- 1
	buffered <- 2
	fmt.Printf("   Capacity: %d\n", cap(buffered))
	fmt.Printf("   Length: %d\n", len(buffered))
	fmt.Println()

	// 12. Key differences
	fmt.Println("12. Key Differences:")
	fmt.Println("   Unbuffered:")
	fmt.Println("     - Synchronous (sender waits for receiver)")
	fmt.Println("     - Capacity: 0")
	fmt.Println("     - Use for synchronization")
	fmt.Println("   Buffered:")
	fmt.Println("     - Asynchronous (sender doesn't wait if buffer not full)")
	fmt.Println("     - Capacity: specified")
	fmt.Println("     - Use for queuing")
}

// Function that sends to channel
func sendMessage(ch chan<- string, msg string) {
	ch <- msg
}

// Function that receives from channel
func receiveMessage(ch <-chan string) string {
	return <-ch
}

