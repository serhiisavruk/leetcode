package main

import "testing"

func TestDoSomething(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"  -", 0},
		{"  -  ", 0},
		{"  -1", -1},
		{"  - 1", 1},
		{" 1", 1},
		{" 1", 1},
		{"-1", -1},
		{"0", 0},
		{"1", 1},
		{"42", 42},
		{"42 word", 42},
		{"42 ", 42},
		{"42   ", 42},
		{"42   word word  ", 42},
	}

	for _, test := range tests {
		result := myAtoi(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %v, but got %v.", test.input, test.expected, result)
		}
	}
}
