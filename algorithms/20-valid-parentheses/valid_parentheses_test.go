package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    string
	expected bool
}{
	{
		input:    "()",
		expected: true,
	},
	{
		input:    "()[]{}",
		expected: true,
	},
	{
		input:    "(]",
		expected: false,
	},
}

func TestValidParentheses(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, isValid(test.input))
	}
}
