package main

import "testing"

func TestDoSomething(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"  -", 0},
		{"+", 0},
		{"  -  ", 0},
		{"  -1", -1},
		{" 1", 1},
		{" 1", 1},
		{"-1", -1},
		{"0", 0},
		{"1", 1},
		{"12", 12},
		{"34 word", 34},
		{"56 ", 56},
		{"78   ", 78},
		{"90   word word  ", 90},
		{"+90   word word  ", 90},
		{"-90   word word  ", -90},
		{"word 123", 0},
		{"word word 123 ", 0},
		{"-91283472332", -2147483648},
		{"+-12", 0},
		{"-+12", 0},
		{"9223372036854775808", 2147483647},
		{"2147483648", 2147483647},
		{"-2147483649", -2147483648},
		{"-9223372036854775808", -2147483648},
		{"2147483646", 2147483646},
		{"-2147483647", -2147483647},
	}

	for _, test := range tests {
		result := myAtoi(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %v, but got %v.", test.input, test.expected, result)
		}
	}
}
