package main

import (
	"errors"
	"fmt"
	"os"
)

// This program demonstrates error handling in Go

// 1. Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// 2. Error with additional context
type NotFoundError struct {
	Resource string
	ID       string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID '%s' not found", e.Resource, e.ID)
}

// 3. Simple function that returns error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 4. Function with custom error
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Message: "age cannot be negative",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Message: "age cannot be greater than 150",
		}
	}
	return nil
}

// 5. Function with error wrapping
func findUser(id string) (*User, error) {
	if id == "" {
		return nil, fmt.Errorf("findUser: %w", errors.New("id cannot be empty"))
	}
	if id == "999" {
		return nil, fmt.Errorf("findUser: %w", NotFoundError{
			Resource: "User",
			ID:       id,
		})
	}
	return &User{ID: id, Name: "John Doe"}, nil
}

// 6. Function using errors.Is
func processUser(id string) error {
	user, err := findUser(id)
	if err != nil {
		// Check if it's a NotFoundError
		var notFound NotFoundError
		if errors.As(err, &notFound) {
			return fmt.Errorf("processUser: user not found: %w", err)
		}
		return fmt.Errorf("processUser: %w", err)
	}
	fmt.Printf("Processing user: %+v\n", user)
	return nil
}

// 7. Function using errors.Is
func checkError(err error) {
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	} else if err != nil {
		fmt.Printf("Other error: %v\n", err)
	}
}

// User struct for examples
type User struct {
	ID   string
	Name string
}

func main() {
	fmt.Println("=== Error Handling ===")
	fmt.Println()

	// 1. Basic error handling
	fmt.Println("1. Basic Error Handling (if err != nil):")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %.2f\n", result)
	}
	fmt.Println()

	// 2. Custom error types
	fmt.Println("2. Custom Error Types:")
	err = validateAge(25)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   Age is valid")
	}

	err = validateAge(-5)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		// Type assertion to access custom fields
		if ve, ok := err.(ValidationError); ok {
			fmt.Printf("   Field: %s, Message: %s\n", ve.Field, ve.Message)
		}
	}
	fmt.Println()

	// 3. Error wrapping with fmt.Errorf
	fmt.Println("3. Error Wrapping (fmt.Errorf with %w):")
	user, err := findUser("")
	if err != nil {
		fmt.Printf("   Wrapped error: %v\n", err)
		fmt.Printf("   Unwrapped: %v\n", errors.Unwrap(err))
	} else {
		fmt.Printf("   User: %+v\n", user)
	}
	fmt.Println()

	// 4. errors.Is for error comparison
	fmt.Println("4. errors.Is for Error Comparison:")
	err = findUser("999")
	if errors.Is(err, NotFoundError{}) {
		fmt.Println("   User not found (using errors.Is)")
	}
	
	// Better: use errors.As for type checking
	var notFound NotFoundError
	if errors.As(err, &notFound) {
		fmt.Printf("   Not found: Resource=%s, ID=%s\n", notFound.Resource, notFound.ID)
	}
	fmt.Println()

	// 5. errors.As for type assertion
	fmt.Println("5. errors.As for Type Assertion:")
	err = validateAge(-10)
	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("   Validation error: Field=%s, Message=%s\n", 
			validationErr.Field, validationErr.Message)
	}
	fmt.Println()

	// 6. Error chain unwrapping
	fmt.Println("6. Error Chain Unwrapping:")
	err = processUser("999")
	if err != nil {
		fmt.Printf("   Error chain: %v\n", err)
		// Unwrap the chain
		current := err
		for current != nil {
			fmt.Printf("     - %v\n", current)
			current = errors.Unwrap(current)
		}
	}
	fmt.Println()

	// 7. Multiple error handling
	fmt.Println("7. Multiple Error Handling:")
	errors := []error{
		validateAge(25),
		validateAge(-5),
		validateAge(200),
	}
	for i, err := range errors {
		if err != nil {
			fmt.Printf("   Error %d: %v\n", i+1, err)
		} else {
			fmt.Printf("   Validation %d: OK\n", i+1)
		}
	}
	fmt.Println()

	// 8. Best practices
	fmt.Println("8. Error Handling Best Practices:")
	fmt.Println("   - Always check errors (if err != nil)")
	fmt.Println("   - Don't ignore errors")
	fmt.Println("   - Return errors, don't panic (unless unrecoverable)")
	fmt.Println("   - Wrap errors with context (fmt.Errorf with %w)")
	fmt.Println("   - Use errors.Is for sentinel errors")
	fmt.Println("   - Use errors.As for custom error types")
	fmt.Println("   - Provide helpful error messages")
}

