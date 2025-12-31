package main

import (
	"fmt"
	"time"
)

// This program demonstrates the select statement in Go

func main() {
	fmt.Println("=== Select Statement ===")
	fmt.Println()

	// 1. Basic select
	fmt.Println("1. Basic Select:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "from ch2"
	}()
	
	select {
	case msg1 := <-ch1:
		fmt.Printf("   Received: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("   Received: %s\n", msg2)
	}
	fmt.Println()

	// 2. Select with default (non-blocking)
	fmt.Println("2. Select with Default (Non-blocking):")
	ch := make(chan int)
	
	select {
	case val := <-ch:
		fmt.Printf("   Received: %d\n", val)
	default:
		fmt.Println("   No value ready, using default")
	}
	fmt.Println()

	// 3. Select with timeout
	fmt.Println("3. Select with Timeout:")
	ch3 := make(chan string)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch3 <- "message"
	}()
	
	select {
	case msg := <-ch3:
		fmt.Printf("   Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("   Timeout: No message received")
	}
	fmt.Println()

	// 4. Select with multiple cases
	fmt.Println("4. Select with Multiple Cases:")
	ch4 := make(chan int)
	ch5 := make(chan int)
	
	go func() {
		ch4 <- 1
	}()
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch5 <- 2
	}()
	
	select {
	case val1 := <-ch4:
		fmt.Printf("   Received from ch4: %d\n", val1)
	case val2 := <-ch5:
		fmt.Printf("   Received from ch5: %d\n", val2)
	}
	fmt.Println()

	// 5. Select in loop
	fmt.Println("5. Select in Loop:")
	ch6 := make(chan int)
	done := make(chan bool)
	
	go func() {
		for i := 1; i <= 3; i++ {
			ch6 <- i
			time.Sleep(100 * time.Millisecond)
		}
		done <- true
	}()
	
	for {
		select {
		case val := <-ch6:
			fmt.Printf("   Received: %d\n", val)
		case <-done:
			fmt.Println("   Done!")
			return
		}
	}
}

// 6. Select for sending
func selectForSending() {
	ch := make(chan int, 1)
	
	select {
	case ch <- 42:
		fmt.Println("   Sent: 42")
	default:
		fmt.Println("   Channel full, couldn't send")
	}
}

// 7. Select with nil channels
func selectWithNilChannels() {
	var ch1, ch2 chan int
	
	select {
	case <-ch1:
		fmt.Println("   Received from ch1")
	case <-ch2:
		fmt.Println("   Received from ch2")
	default:
		fmt.Println("   All channels nil, using default")
	}
}

// 8. Select for cancellation
func selectForCancellation() {
	ch := make(chan int)
	cancel := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- i:
				time.Sleep(100 * time.Millisecond)
			case <-cancel:
				fmt.Println("   Cancelled!")
				return
			}
		}
	}()
	
	// Receive a few values
	for i := 0; i < 3; i++ {
		fmt.Printf("   Received: %d\n", <-ch)
	}
	
	// Cancel
	cancel <- true
	time.Sleep(100 * time.Millisecond)
}

