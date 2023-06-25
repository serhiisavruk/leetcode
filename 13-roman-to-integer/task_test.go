package main

import "testing"

func TestDoSomething(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MCDLXXVI", 1476},
	}

	for _, test := range tests {
		result := romanToInt(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %d, but got %d.", test.input, test.expected, result)
		}
	}
}
