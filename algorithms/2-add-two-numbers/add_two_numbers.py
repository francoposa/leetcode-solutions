from typing import Optional


class ListNode:
    """Definition for singly-linked list"""

    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    """https://leetcode.com/problems/add-two-numbers/"""

    def addTwoNumbers(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        if l1 and not l2:
            return l1
        elif not l1 and l2:
            return l2
        elif not l1 and not l2:
            return None

        sum: int = l1.val + l2.val
        # Let the rest of the recursion run, which returns the rest of the number chain,
        # not including any carrying we may need to do for two digits with sum >= 10
        next: Optional[ListNode] = self.addTwoNumbers(l1.next, l2.next)
        if sum >= 10:
            # sum can't be more than 18 (from adding two nines),
            # so we only need to modulo once to get back to a single digit
            sum %= 10
            # The next digit (which is one tens place higher) needs to be increased by 1
            # from whatever came back from the rest of the recursion.
            # If the next digit came back >= 10 (only possible case is next.Val = 9),
            # then the recursive call will hit this code path again and correct it.
            next: Optional[ListNode] = self.addTwoNumbers(
                next,
                ListNode(val=1, next=None),
            )
        return ListNode(val=sum, next=next)
