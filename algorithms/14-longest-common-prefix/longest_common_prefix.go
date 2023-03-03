package longest_common_prefix

import "math"

// https://leetcode.com/problems/longest-common-prefix
func longestCommonPrefix(strs []string) string {
	minLen := math.MaxInt32
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}
	prefix := ""
	for i := 0; i < minLen; i++ {
		char := strs[0][i]
		match := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				match = false
				break
			}
		}
		if match {
			prefix += string(char)
		} else {
			break
		}
	}
	return prefix
}
