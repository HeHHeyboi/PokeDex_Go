package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check for length mismatch
		if len(actual) != len(c.expected) {
			t.Errorf(`Length mismatch:
Input: %q
Expected length: %d, Actual length: %d
Expected: %v
Actual:   %v`,
				c.input, len(c.expected), len(actual), c.expected, actual)
			continue
		}

		// Check for word mismatches
		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf(`Word mismatch:
Input: %q
Expected: %v, Actual: %v
Mismatch at index %d: expected %q, got %q`,
					c.input, c.expected, actual, i, c.expected[i], word)
				break
			}
		}
	}
}
