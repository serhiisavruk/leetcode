package main

import "testing"

func TestDoSomething(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{""}, ""},
		{[]string{"a"}, "a"},
		{[]string{"aaa"}, "aaa"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"flower", "flower", "flower", "flower"}, "flower"},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s, but got %s.", test.input, test.expected, result)
		}
	}
}
