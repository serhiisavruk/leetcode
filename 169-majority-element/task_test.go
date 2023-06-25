package main

import (
	"testing"
)

func TestIscontainingDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{-1}, -1},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}, 1},
		{[]int{-1, -1}, -1},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 2, 3, 3, 4}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 7, -1}, 7},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, 1},
	}

	for _, test := range tests {
		result := majorityElement(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %v, but got %v.", test.input, test.expected, result)
		}
	}
}
