package main

import "testing"

func TestDoSomething(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{" ", 0},
		{"   ", 0},
		{"a", 1},
		{"ab", 2},
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, test := range tests {
		result := lengthOfLastWord(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, but got %v.", test.input, test.expected, result)
		}
	}
}
