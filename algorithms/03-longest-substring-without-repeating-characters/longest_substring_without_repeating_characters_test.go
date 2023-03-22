package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	var input string
	var expected int
	var actual int

	input = "abcabcbb"
	expected = 3
	actual = lengthOfLongestSubstring(input)
	assert.Equal(t, expected, actual)

	input = "bbbbb"
	expected = 1
	actual = lengthOfLongestSubstring(input)
	assert.Equal(t, expected, actual)

	input = "pwwkew"
	expected = 3
	actual = lengthOfLongestSubstring(input)
	assert.Equal(t, expected, actual)

	input = ""
	expected = 0
	actual = lengthOfLongestSubstring(input)
	assert.Equal(t, expected, actual)

}
