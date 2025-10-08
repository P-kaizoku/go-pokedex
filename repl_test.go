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
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal")
			continue
		}
		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("%v is not equal %v", actual[i], cs.expected[i])
			}
		}
	}

}
