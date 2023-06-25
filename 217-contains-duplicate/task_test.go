package main

import (
	"testing"
)

func TestIscontainingDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{-1}, false},
		{[]int{0}, false},
		{[]int{1}, false},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, false},
		{[]int{-1, -1}, true},
		{[]int{0, 0, 0}, true},
		{[]int{1, 2, 3, 3, 4}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 7, -1}, true},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, test := range tests {
		result := containsDuplicate(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %v, but got %v.", test.input, test.expected, result)
		}
	}
}
