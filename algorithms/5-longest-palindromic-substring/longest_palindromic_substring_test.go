package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var input string
	var expected string
	var actual string

	input = "babad"
	expected = "bab"
	actual = longestPalindrome(input)
	assert.Equal(t, expected, actual)

	input = "cbbd"
	expected = "bb"
	actual = longestPalindrome(input)
	assert.Equal(t, expected, actual)

	input = "a"
	expected = "a"
	actual = longestPalindrome(input)
	assert.Equal(t, expected, actual)

	input = "ac"
	expected = "a"
	actual = longestPalindrome(input)
	assert.Equal(t, expected, actual)
}
