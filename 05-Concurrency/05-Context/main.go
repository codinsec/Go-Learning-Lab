package main

import (
	"context"
	"fmt"
	"time"
)

// This program demonstrates context package in Go

func main() {
	fmt.Println("=== Context ===")
	fmt.Println()

	// 1. Background context
	fmt.Println("1. Background Context:")
	ctx := context.Background()
	fmt.Printf("   Context: %v\n", ctx)
	fmt.Println()

	// 2. WithCancel
	fmt.Println("2. WithCancel:")
	ctx2, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("   Cancelling context...")
		cancel()
	}()
	
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("   Timeout")
	case <-ctx2.Done():
		fmt.Printf("   Context cancelled: %v\n", ctx2.Err())
	}
	fmt.Println()

	// 3. WithTimeout
	fmt.Println("3. WithTimeout:")
	ctx3, cancel3 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel3()
	
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("   Operation completed")
	case <-ctx3.Done():
		fmt.Printf("   Context timeout: %v\n", ctx3.Err())
	}
	fmt.Println()

	// 4. WithDeadline
	fmt.Println("4. WithDeadline:")
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx4, cancel4 := context.WithDeadline(context.Background(), deadline)
	defer cancel4()
	
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("   Operation completed")
	case <-ctx4.Done():
		fmt.Printf("   Deadline exceeded: %v\n", ctx4.Err())
	}
	fmt.Println()

	// 5. WithValue
	fmt.Println("5. WithValue:")
	ctx5 := context.WithValue(context.Background(), "userID", "12345")
	ctx5 = context.WithValue(ctx5, "requestID", "req-001")
	
	userID := ctx5.Value("userID")
	requestID := ctx5.Value("requestID")
	fmt.Printf("   UserID: %v\n", userID)
	fmt.Printf("   RequestID: %v\n", requestID)
	fmt.Println()

	// 6. Context in function
	fmt.Println("6. Context in Function:")
	ctx6, cancel6 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel6()
	
	result := doWork(ctx6)
	fmt.Printf("   Result: %s\n", result)
	fmt.Println()

	// 7. Context propagation
	fmt.Println("7. Context Propagation:")
	ctx7, cancel7 := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel7()
	}()
	
	processRequest(ctx7)
	fmt.Println()

	// 8. Best practices
	fmt.Println("8. Best Practices:")
	fmt.Println("   - Pass context as first parameter")
	fmt.Println("   - Use context for cancellation and timeouts")
	fmt.Println("   - Don't store context in structs")
	fmt.Println("   - Use WithValue sparingly (for request-scoped data)")
	fmt.Println("   - Always check ctx.Done() in loops")
}

// Function that accepts context
func doWork(ctx context.Context) string {
	select {
	case <-time.After(100 * time.Millisecond):
		return "Work completed"
	case <-ctx.Done():
		return "Work cancelled: " + ctx.Err().Error()
	}
}

// Function that propagates context
func processRequest(ctx context.Context) {
	select {
	case <-time.After(300 * time.Millisecond):
		fmt.Println("   Request processed")
	case <-ctx.Done():
		fmt.Printf("   Request cancelled: %v\n", ctx.Err())
	}
}

