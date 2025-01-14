package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{"input string cannot be empty"},
		},
	}

	for _, c := range cases {
		actual, err := cleanInput(c.input)
		if err != nil {
			actual := err.Error()
			expectedError := c.expected[0]

			if actual != expectedError {
				t.Errorf("unexpected error mismatch:\n%s -> %s", actual, expectedError)
			}
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected and actual do not match:\n%s -> %s", word, expectedWord)
			}
		}
	}
}
