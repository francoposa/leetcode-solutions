package algorithmanalysis

// https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		// minor optimization; problem definition implies that a negative x can never
		// be a palindrome, as the negative sign will never be matched on the opposite end
		return false
	}
	// split x into its digits
	xDigits := []int{}
	for x > 0 {
		xDigits = append(xDigits, x%10)
		x /= 10
	}
	eqDigits := true

	// walk digits from start and end to middle; check for equality
	for i, j := 0, len(xDigits)-1; i < j; i, j = i+1, j-1 {
		if xDigits[i] != xDigits[j] {
			eqDigits = false
			break
		}
	}
	return eqDigits
}
