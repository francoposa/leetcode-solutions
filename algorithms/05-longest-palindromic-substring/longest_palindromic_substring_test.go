package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    string
	expected string
}{
	{
		input:    "babad",
		expected: "bab",
	},
	{
		input:    "cbbd",
		expected: "bb",
	},
	{
		input:    "abb",
		expected: "bb",
	},
	{
		input:    "a",
		expected: "a",
	},
	{
		input:    "ac",
		expected: "a",
	},
}

func TestLongestPalindromeSlow(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, longestPalindromeSlow(test.input))
	}
}
