from collections import defaultdict


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        row_num_chars = defaultdict(list)
        row_num, step = 0, 1

        for char in s:
            row_num_chars[row_num].append(char)
            if row_num == numRows - 1 and step == 1:
                # Reached the bottom row, go back up
                step = -1
            elif row_num == 0 and step == -1:
                # Returned to top rowm head back down
                step = 1

            row_num += step

        result = ""
        for i in range(numRows):
            result += "".join(row_num_chars[i])

        return result

