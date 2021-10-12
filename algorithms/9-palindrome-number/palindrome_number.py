class Solution:
    """https://leetcode.com/problems/palindrome-number/"""

    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            # minor optimization; problem definition implies that a negative x can never
            # be a palindrome, as the negative sign will never be matched on the opposite end
            return False

        # converting a n-digit base-10 integer to a string will generally be O(n^2)
        # https://stackoverflow.com/questions/44025315/what-is-the-time-complexity-of-int1010-2
        # https://stackoverflow.com/questions/60788680/what-is-the-time-complexity-of-type-casting-function-in-python
        int_str = str(x)

        left = 0
        right = len(int_str) - 1
        while right > left:
            if int_str[left] != int_str[right]:
                return False
            left += 1
            right -= 1

        return True
