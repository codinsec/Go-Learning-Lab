package main

import (
	"fmt"
	"os"
)

// This program demonstrates defer, panic, and recover in Go
func main() {
	fmt.Println("=== Defer, Panic, Recover ===")
	fmt.Println()

	// 1. Basic defer
	fmt.Println("1. Basic Defer:")
	deferExample()
	fmt.Println()

	// 2. Multiple defers (LIFO - Last In First Out)
	fmt.Println("2. Multiple Defers (LIFO):")
	multipleDefers()
	fmt.Println()

	// 3. Defer with function arguments
	fmt.Println("3. Defer with Function Arguments:")
	deferArguments()
	fmt.Println()

	// 4. Defer for cleanup
	fmt.Println("4. Defer for Cleanup:")
	fileCleanup()
	fmt.Println()

	// 5. Panic
	fmt.Println("5. Panic:")
	fmt.Println("   Calling function that panics...")
	// Uncomment to see panic:
	// panicExample()
	fmt.Println("   (Panic example commented out to prevent program termination)")
	fmt.Println()

	// 6. Recover
	fmt.Println("6. Recover:")
	recoverExample()
	fmt.Println()

	// 7. Defer, Panic, Recover together
	fmt.Println("7. Defer, Panic, Recover Together:")
	safeFunction()
	fmt.Println()

	// 8. Common use cases
	fmt.Println("8. Common Use Cases:")
	commonUseCases()
}

func deferExample() {
	defer fmt.Println("   This runs last (deferred)")
	fmt.Println("   This runs first")
	fmt.Println("   This runs second")
}

func multipleDefers() {
	defer fmt.Println("   Defer 1 (runs last)")
	defer fmt.Println("   Defer 2 (runs second to last)")
	defer fmt.Println("   Defer 3 (runs third to last)")
	fmt.Println("   Normal execution (runs first)")
}

func deferArguments() {
	x := 10
	defer fmt.Printf("   Deferred print: x = %d\n", x) // x is evaluated now (10)
	x = 20
	fmt.Printf("   After change: x = %d\n", x)
	fmt.Println("   Note: Deferred function uses original value (10)")
}

func fileCleanup() {
	// Simulate file operations
	fmt.Println("   Opening file...")
	defer fmt.Println("   Closing file (deferred cleanup)")
	
	fmt.Println("   Reading file...")
	fmt.Println("   Processing data...")
	// File would be closed here automatically, even if an error occurred
}

func panicExample() {
	fmt.Println("   About to panic...")
	panic("Something went wrong!")
	fmt.Println("   This will never execute")
}

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   Recovered from panic: %v\n", r)
		}
	}()
	
	fmt.Println("   About to panic...")
	panic("Controlled panic for demonstration")
	fmt.Println("   This will never execute")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   Recovered: %v\n", r)
		}
		fmt.Println("   Cleanup: This always runs")
	}()
	
	fmt.Println("   Normal execution")
	fmt.Println("   Simulating error condition...")
	panic("Simulated error")
}

func commonUseCases() {
	// Use case 1: Resource cleanup
	fmt.Println("   Use Case 1: Resource Cleanup")
	resourceCleanup()
	
	// Use case 2: Error recovery
	fmt.Println("\n   Use Case 2: Error Recovery")
	errorRecovery()
	
	// Use case 3: Logging
	fmt.Println("\n   Use Case 3: Logging")
	loggingExample()
}

func resourceCleanup() {
	// Simulate opening a resource
	fmt.Println("     Opening resource...")
	defer func() {
		fmt.Println("     Closing resource (deferred)")
	}()
	
	fmt.Println("     Using resource...")
	// Resource will be closed even if function returns early or panics
}

func errorRecovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("     Error handled gracefully: %v\n", r)
		}
	}()
	
	// Simulate an operation that might panic
	fmt.Println("     Performing risky operation...")
	// In real code, this might be: result := riskyOperation()
	fmt.Println("     Operation completed successfully")
}

func loggingExample() {
	defer func() {
		fmt.Println("     Function completed (logged via defer)")
	}()
	
	fmt.Println("     Executing function logic...")
}

// Real-world example: File operations
func readFileExample(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Always close the file
	
	// Read file operations here
	// File will be closed automatically when function returns
	
	return nil
}

