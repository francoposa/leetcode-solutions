package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "III",
		expected: 3,
	},
	{
		input:    "LVIII",
		expected: 58,
	},
	{
		input:    "MCMXCIV",
		expected: 1994,
	},
}

func TestRomanToInteger(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, romanToInt(test.input))
	}
}
