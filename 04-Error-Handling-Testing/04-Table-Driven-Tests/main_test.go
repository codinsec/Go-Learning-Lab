package main

import "testing"

// TestIsPalindrome demonstrates basic table-driven test
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"palindrome word", "racecar", true},
		{"palindrome phrase", "A man a plan a canal Panama", true},
		{"not palindrome", "hello", false},
		{"case insensitive", "RaceCar", true},
		{"with spaces", "race car", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestValidateEmail demonstrates table-driven test with error cases
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid email", "user@example.com", true},
		{"valid email with subdomain", "user@mail.example.com", true},
		{"missing @", "userexample.com", false},
		{"multiple @", "user@@example.com", false},
		{"missing domain", "user@", false},
		{"missing local", "@example.com", false},
		{"missing dot", "user@examplecom", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateEmail(tt.input)
			if result != tt.expected {
				t.Errorf("ValidateEmail(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCalculateGrade demonstrates table-driven test with ranges
func TestCalculateGrade(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"A grade", 95, "A"},
		{"A grade boundary", 90, "A"},
		{"B grade", 85, "B"},
		{"B grade boundary", 80, "B"},
		{"C grade", 75, "C"},
		{"C grade boundary", 70, "C"},
		{"D grade", 65, "D"},
		{"D grade boundary", 60, "D"},
		{"F grade", 50, "F"},
		{"F grade boundary", 0, "F"},
		{"negative score", -10, "F"},
		{"over 100", 150, "A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateGrade(tt.input)
			if result != tt.expected {
				t.Errorf("CalculateGrade(%d) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestReverseString demonstrates table-driven test with strings
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"simple word", "hello", "olleh"},
		{"palindrome", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "café", "éfac"},
		{"numbers", "12345", "54321"},
		{"mixed", "a1b2c3", "3c2b1a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestMultipleFunctions demonstrates testing multiple functions in one table
func TestMultipleFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func(string) bool
		input    string
		expected bool
	}{
		{"IsPalindrome - valid", IsPalindrome, "racecar", true},
		{"IsPalindrome - invalid", IsPalindrome, "hello", false},
		{"ValidateEmail - valid", ValidateEmail, "user@example.com", true},
		{"ValidateEmail - invalid", ValidateEmail, "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.input)
			if result != tt.expected {
				t.Errorf("Function(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestWithHelper demonstrates table-driven test with helper function
func TestWithHelper(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			checkPalindrome(t, tt.input, tt.expected)
		})
	}
}

// Helper function for testing
func checkPalindrome(t *testing.T, input string, expected bool) {
	t.Helper()
	result := IsPalindrome(input)
	if result != expected {
		t.Errorf("IsPalindrome(%q) = %v; expected %v", input, result, expected)
	}
}

