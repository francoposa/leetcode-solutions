package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbersIter(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{}
	head := current
	carry := 0
	for {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			carry = sum / 10
			sum %= 10
		} else {
			carry = 0
		}

		current.Val = sum
		if l1 != nil || l2 != nil || carry != 0 {
			current.Next = &ListNode{}
			current = current.Next
		} else {
			break
		}
	}

	return head
}

func addTwoNumbersRecur(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		// recursion has bottomed out, return back up the chain
		return nil
	}

	sum := 0
	if l1 != nil {
		sum += l1.Val
	}
	if l2 != nil {
		sum += l2.Val
	}

	if sum >= 10 {
		carry := sum / 10
		sum %= 10
		if l1.Next != nil {
			l1.Next.Val += carry
		} else {
			l1.Next = &ListNode{Val: carry, Next: nil}
		}
	}

	// prevent calling .Next on nil node
	var l1Next, l2Next *ListNode
	if l1 != nil {
		l1Next = l1.Next
	}
	if l2 != nil {
		l2Next = l2.Next
	}

	next := addTwoNumbersRecur(l1Next, l2Next)

	return &ListNode{Val: sum, Next: next}
}
