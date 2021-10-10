class Solution:
    """https://leetcode.com/problems/longest-palindromic-substring/"""

    def longestPalindrome(self, s: str) -> str:
        longest_palindrome_start_index = 0
        longest_palindrome_len = 1

        def _longest_palindrome(
            s: str, start_index_inclusive: int, end_index_inclusive: int
        ):
            # https://stackoverflow.com/questions/12182068/python-closure-function-losing-outer-variable-access
            nonlocal longest_palindrome_start_index
            nonlocal longest_palindrome_len
            while start_index_inclusive >= 0 and end_index_inclusive < len(s):
                if s[start_index_inclusive] != s[end_index_inclusive]:
                    break

                candidate_palindrome_length = (
                    end_index_inclusive - start_index_inclusive + 1
                )
                if longest_palindrome_len < candidate_palindrome_length:
                    longest_palindrome_start_index = start_index_inclusive
                    longest_palindrome_len = candidate_palindrome_length

                start_index_inclusive -= 1
                end_index_inclusive += 1

        for i in range(1, len(s)):
            # Find odd-length palindrome centered at i
            # First candidate palindrome is length 1
            start_index_inclusive = i
            end_index_inclusive = i
            _longest_palindrome(s, start_index_inclusive, end_index_inclusive)

            # Find an even-length palindrome "centered" at i
            # First candidate palindrome is length 2
            start_index_inclusive = i - 1
            end_index_inclusive = i
            _longest_palindrome(s, start_index_inclusive, end_index_inclusive)

        return s[
            longest_palindrome_start_index : longest_palindrome_start_index
            + longest_palindrome_len
        ]
