class Solution:
    """https://leetcode.com/problems/reverse-integer/"""

    def reverse(self, x: int) -> int:
        if x == 0:
            return 0

        int_str = str(x)

        if int_str[0] == "-":
            # Strip negative sign so string reverse deals only with digits
            int_str = int_str[1:]
            is_negative = True
        else:
            is_negative = False

        reversed_int_str = int_str[::-1].lstrip("0")
        reversed_int = int(reversed_int_str)

        # Reattach negative signed-ness
        if is_negative:
            reversed_int_signed = reversed_int * -1

        # Check for overflow of 32-bit int type
        if reversed_int_signed > (2 ** 31) - 1 or reversed_int_signed < (-(2 ** 31)):
            return 0

        return reversed_int_signed
