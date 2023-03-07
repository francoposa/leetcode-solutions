package valid_parentheses

var openToCloseBracketMap = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isOpenBracket(r rune) bool {
	_, ok := openToCloseBracketMap[r]
	return ok
}

var closeToOpenBracketMap = map[rune]rune{}

func init() {
	for k, v := range openToCloseBracketMap {
		closeToOpenBracketMap[v] = k
	}
}

// https://leetcode.com/problems/valid-parentheses
func isValid(s string) bool {
	openBracketStack := make([]rune, 0, len(s)/2)
	for _, bracket := range s {
		if isOpenBracket(bracket) {
			openBracketStack = append(openBracketStack, bracket)
		} else {
			if len(openBracketStack) == 0 {
				return false
			}
			expectedOpenBracket := closeToOpenBracketMap[bracket]
			lastOpenBracket := openBracketStack[len(openBracketStack)-1]
			if lastOpenBracket != expectedOpenBracket {
				// open brackets not closed in correct order
				return false
			} else {
				// ok; pop closed bracket off stack
				openBracketStack = openBracketStack[:len(openBracketStack)-1]
			}
		}
	}
	return len(openBracketStack) == 0
}
