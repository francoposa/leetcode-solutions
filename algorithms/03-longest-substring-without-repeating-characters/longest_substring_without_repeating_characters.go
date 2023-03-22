package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {

	startIndex := 0
	longestSubstringLen := 0
	charIndexMap := map[rune]int{}

	for i, char := range s {
		prevCharIndex, ok := charIndexMap[char]
		if ok && prevCharIndex >= startIndex {
			// the current char appeared previously in the current substring
			// so the current char cannot be added to the current substring.
			// before we start a new substring, check if the current one beats the longest
			currentSubstringLen := i - startIndex
			if currentSubstringLen > longestSubstringLen {
				longestSubstringLen = currentSubstringLen
			}
			// the earliest the new non-character-repeating substring can start
			// is the index after the previous time the character appeared
			startIndex = prevCharIndex + 1
		}
		charIndexMap[char] = i
	}

	// when we reach the end of the loop without hitting a duplicate character
	// we miss the check for whether the current substring is the longest
	currentSubstringLen := len(s) - startIndex
	if currentSubstringLen > longestSubstringLen {
		longestSubstringLen = currentSubstringLen
	}

	return longestSubstringLen
}
