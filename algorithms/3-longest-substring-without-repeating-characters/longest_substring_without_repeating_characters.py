class Solution:
    """https://leetcode.com/problems/longest-substring-without-repeating-characters/"""

    def lengthOfLongestSubstring(self, s: str) -> int:

        len_full_string = len(s)

        i = 0
        j = 1
        longest_substring_len = 0

        while i < len_full_string and j <= len_full_string:
            current_candidate_substring = s[i:j]

            len_current_candidate_substring = j - i
            if len_current_candidate_substring > longest_substring_len:
                longest_substring_len = len_current_candidate_substring

            if j < len_full_string:
                next_char = s[j]
                next_char_index_in_current_candidate_substring = current_candidate_substring.find(
                    next_char
                )
                if next_char_index_in_current_candidate_substring >= 0:
                    # current_candidate_substring starting at i cannot
                    # be extended any further without repeating characters
                    # start a new search at a point where we're not going to run
                    # into the same repeated characters again
                    i += next_char_index_in_current_candidate_substring + 1
            j += 1

        return longest_substring_len
