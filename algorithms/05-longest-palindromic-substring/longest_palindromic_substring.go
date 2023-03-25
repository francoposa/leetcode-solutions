package longest_palindromic_substring

// https://leetcode.com/problems/longest-palindromic-substring
func longestPalindromeSlow(s string) string {
	// transform to always have odd length by adding separators
	s = transformToOddLenWithSeparators(s)

	// for each char c, check all possible substrings with c at the center
	maxPalindrome := ""
	for centerIdx := 0; centerIdx < len(s); centerIdx++ {
		maxRadius := min(centerIdx, (len(s)-1)-centerIdx)
		// starting at char c as the center, expand in both directions until the palindrome breaks
		for radius := 0; radius <= maxRadius; radius++ {
			if s[centerIdx-radius] == s[centerIdx+radius] {
				// we still have a palindrome
				if ((2 * radius) + 1) > len(maxPalindrome) {
					// palindrome is longer than our previous max; update maxPalindrome
					maxPalindrome = s[centerIdx-radius : centerIdx+radius+1]
				}
			} else {
				// palindrome has been broken, we can't make a longer one by further expanding
				break
			}
		}
	}

	// remove separators
	return transformFromOddLenRemoveSeparators(maxPalindrome)
}

const separator = '|'

func transformToOddLenWithSeparators(input string) string {
	sPrime := make([]rune, 0, (len(input)*2)+1)
	for _, char := range input {
		sPrime = append(sPrime, separator)
		sPrime = append(sPrime, char)
	}
	sPrime = append(sPrime, separator)
	return string(sPrime)
}

func transformFromOddLenRemoveSeparators(maxPalindrome string) string {
	maxPalindromePrime := make([]rune, 0, (len(maxPalindrome)/2)-1)
	for _, char := range maxPalindrome {
		if char != separator {
			maxPalindromePrime = append(maxPalindromePrime, char)
		}
	}
	return string(maxPalindromePrime)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
