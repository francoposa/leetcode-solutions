package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	header := &ListNode{}
	mergedList := header
	for {
		nextNode := &ListNode{}
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				nextNode.Val = list1.Val
				list1 = list1.Next
			} else {
				nextNode.Val = list2.Val
				list2 = list2.Next
			}
		} else if list1 != nil && list2 == nil {
			nextNode.Val = list1.Val
			list1 = list1.Next
		} else if list2 != nil {
			nextNode.Val = list2.Val
			list2 = list2.Next
		} else {
			// both nil
			break
		}
		mergedList.Next = nextNode
		mergedList = nextNode
	}
	return header.Next
}
