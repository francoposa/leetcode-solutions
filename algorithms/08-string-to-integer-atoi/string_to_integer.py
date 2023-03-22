class Solution:

    NEGATIVE_SIGN = "-"
    MAX_INT32 = 2 ** 31 - 1
    MIN_INT32 = -(2 ** 31)

    def myAtoi(self, str: str) -> int:
        whitespace_stripped_str = (
            str.strip()
        )  # Removed all whitespace from left and right

        neg_stripped_str = whitespace_stripped_str.lstrip("-")
        is_negative = len(neg_stripped_str) != len(whitespace_stripped_str)

        neg_and_pos_stripped_str = neg_stripped_str.lstrip("+")

        if len(neg_and_pos_stripped_str) < len(whitespace_stripped_str) - 1:
            # We stripped a negative and positive sign off the front - this is invalid
            return 0

        int_str = ""
        # All possible valid leading values (whitespace and negative sign) have been stripped
        # Now we just process as many digits as we can and break out on the first non-digit

        for i, char in enumerate(neg_and_pos_stripped_str):
            if char.isdigit():
                int_str += char
            else:
                break

        if len(int_str) == 0:
            return 0

        parsed_int = -1 * int(int_str) if is_negative else int(int_str)

        if parsed_int > self.MAX_INT32:
            return self.MAX_INT32

        if parsed_int < self.MIN_INT32:
            return self.MIN_INT32

        return parsed_int
