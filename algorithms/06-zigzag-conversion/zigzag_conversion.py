from collections import defaultdict


class Solution:
    """https://leetcode.com/problems/zigzag-conversion/"""

    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        row_content = {row_num: "" for row_num in range(numRows)}
        row_num, step = 0, 1

        for char in s:
            row_content[row_num] += char
            if row_num == numRows - 1 and step == 1:
                # Reached the bottom row, go back up
                step = -1
            elif row_num == 0 and step == -1:
                # Returned to top row, head back down
                step = 1

            row_num += step

        converted_s = ""
        for content in row_content.values():
            converted_s += content

        return converted_s
