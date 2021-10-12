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
        # Any non-string-conversion approach would still need to get the magnitude of the integer
        # in order to know when we have reached the "middle", needed to determine isPalindrome
        # There is debate on whether using any other approach to get magnitude of the integer is
        # actually faster in real-world cases where the input is unbounded:
        # https://stackoverflow.com/questions/2189800/how-to-find-length-of-digits-in-an-integer
        int_str = str(x)

        left = 0
        right = len(int_str) - 1
        while right > left:
            if int_str[left] != int_str[right]:
                return False
            left += 1
            right -= 1

        return True
