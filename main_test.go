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
			input:    "   Hello this is some test",
			expected: []string{"hello", "this", "is", "some", "test"},
		},
		{
			input:    "Is this going to WORK    I don't know! !!!!  ",
			expected: []string{"is", "this", "going", "to", "work", "i", "don't", "know!", "!!!!"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   1 2 3 4 5 6 7 8 9 0   ",
			expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
		},
		{
			input:    "Let's see with weird characters &/%&/%34·##@#| wdweuihdfw ",
			expected: []string{"let's", "see", "with", "weird", "characters", "&/%&/%34·##@#|", "wdweuihdfw"},
		},
	}
	for i, c := range cases {
		split := cleanInput(c.input)
		if len(split) != len(c.expected) {
			t.Errorf("split slice length is incorrect for test case %d\nactual: %d \t expected: %d", i, len(split), len(c.expected))
			t.Fail()
		}
		for j := range split {
			current_word := split[j]
			expected_word := c.expected[j]
			if current_word != expected_word {
				t.Errorf("mismatch in test case %d\nword: %s \t expected: %s", i, current_word, expected_word)
				t.Fail()
			}
		}
	}
}
