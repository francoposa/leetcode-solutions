class ListNode:
    """Definition for singly-linked list"""

    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    """https://leetcode.com/problems/add-two-numbers/"""

    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        num1 = self._intFromListNode(l1)
        num2 = self._intFromListNode(l2)

        return self._listNodeFromInt(num1 + num2)

    def _intFromListNode(self, l: ListNode) -> int:
        factor = 1
        base = 10
        num = 0
        while l is not None:
            num += l.val * factor
            l = l.next
            factor *= base

        return num

    def _listNodeFromInt(self, n: int) -> ListNode:
        base = 10
        node = ListNode()
        first_node = node
        while n > 0:
            node.val = n % base
            n = n // base

            if n > 0:
                next_node = ListNode()
                node.next = next_node
                node = next_node

        return first_node
