package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input    []string
	expected string
}{
	{
		input:    []string{"flower", "flow", "flight"},
		expected: "fl",
	},
	{
		input:    []string{"dog", "racecar", "car"},
		expected: "",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, longestCommonPrefix(test.input))
	}
}
