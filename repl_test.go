package main

import (
	"testing"
)

type test struct {
	name     string
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []test{
		{
			name:     "Double Trailing",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Hello, World!",
			input:    "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
		{
			name:     "One Big String",
			input:    "ThisIsOneBigString",
			expected: []string{"thisisonebigstring"},
		},
		{
			name:     "Right white space",
			input:    "There's only trailing spaces on the right !!!    \t\t",
			expected: []string{"there's", "only", "trailing", "spaces", "on", "the", "right", "!!!"},
		},
		{
			name:     "Left White Space",
			input:    "\t\t\t\n      Now they're on the left!!",
			expected: []string{"now", "they're", "on", "the", "left!!"},
		},
		{
			name:     "All Numbers",
			input:    "12 24 48 53,090", // ToLower should handle this no problem, but it doesn't hurt to test
			expected: []string{"12", "24", "48", "53,090"},
		},
		{
			name:     "All White Space",
			input:    "      \t\t\t\n\n\t   ",
			expected: nil,
		},
		{
			name:     "Empty",
			input:    "", // Don't want things to break when no input is given, we'll handle it properly
			expected: nil,
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLength := len(actual)
		expectedLength := len(c.expected)
		if actualLength < expectedLength {
			t.Errorf("Not enough elements in actual")
			t.Fatalf("%v: expected length %v, got length %v", c.name, len(c.expected), len(actual))
		}
		if actualLength > expectedLength {
			t.Errorf("Too many enough elements in actual")
			t.Fatalf("%v: expected length %v, got length %v", c.name, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words mismatch")
				t.Fatalf("%v: Expected: %v, Got: %v", c.name, expectedWord, word)
			}
		}
	}
}
