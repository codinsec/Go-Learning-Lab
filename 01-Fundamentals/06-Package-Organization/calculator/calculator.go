package calculator

// Add is an exported function (starts with uppercase)
// It can be used by other packages
func Add(a, b int) int {
	return a + b
}

// Multiply is an exported function
func Multiply(a, b int) int {
	return a * b
}

// add is an unexported function (starts with lowercase)
// It can only be used within this package
func add(a, b int) int {
	return a + b
}

// Calculator is an exported type
type Calculator struct {
	// result is an unexported field
	result int
}

// NewCalculator is an exported constructor function
func NewCalculator() *Calculator {
	return &Calculator{result: 0}
}

// Add is an exported method
func (c *Calculator) Add(value int) {
	c.result = add(c.result, value) // Using unexported function
}

// Subtract is an exported method
func (c *Calculator) Subtract(value int) int {
	c.result -= value
	return c.result
}

// GetResult is an exported method
func (c *Calculator) GetResult() int {
	return c.result
}

// getResult is an unexported method (not accessible outside package)
func (c *Calculator) getResult() int {
	return c.result
}

