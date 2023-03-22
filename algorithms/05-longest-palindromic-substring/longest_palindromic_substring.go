package longest_palindromic_substring

func longestPalindrome(s string) string {
	longestPalindromeStartIndex := 0
	longestPalindromeLen := 1

	longerPalindrome := func(s string, startIndexInclusive, endIndexInclusive int) {
		for startIndexInclusive >= 0 && endIndexInclusive < len(s) {
			if s[startIndexInclusive] != s[endIndexInclusive] {
				break
			}
			candidatePalindromeLen := endIndexInclusive - startIndexInclusive + 1
			if candidatePalindromeLen > longestPalindromeLen {
				longestPalindromeStartIndex = startIndexInclusive
				longestPalindromeLen = candidatePalindromeLen
			}
			startIndexInclusive--
			endIndexInclusive++
		}
	}

	for i := 1; i < len(s); i++ {
		// Find odd-length palindrome centered at i
		// First candidate palindrome length is 1
		startIndexInclusive := i
		endIndexInclusive := i
		longerPalindrome(s, startIndexInclusive, endIndexInclusive)

		// Find odd-length palindrome "centered" at i
		// First candidate palindrome length is 2
		startIndexInclusive = i - 1
		endIndexInclusive = i
		longerPalindrome(s, startIndexInclusive, endIndexInclusive)
	}

	return s[longestPalindromeStartIndex : longestPalindromeStartIndex+longestPalindromeLen]
}
