package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},     // positive palindrome number
		{-121, false},   // negative numbers are not palindromes
		{10, false},     // number ending with 0 (but not 0 itself) is not palindrome
		{0, true},       // zero is a palindrome
		{1234321, true}, // large palindrome number
		{12345, false},  // non-palindrome number
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %v, but got %v.", test.input, test.expected, result)
		}
	}
}
