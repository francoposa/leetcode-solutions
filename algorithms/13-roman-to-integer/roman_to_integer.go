package roman_to_integer

var symbolVals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// https://leetcode.com/problems/roman-to-integer
func romanToInt(s string) int {
	sum := 0
	previousVal := 0
	for i, symbol := range s {
		val := symbolVals[symbol]
		sum += val

		if i > 0 && previousVal < val {
			sum -= 2 * previousVal
		}
		previousVal = val
	}
	return sum
}
