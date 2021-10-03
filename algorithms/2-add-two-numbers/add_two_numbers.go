package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	} // list repr of 342
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	} // list repr of 465

	lSum := addTwoNumbers(l1, l2) // will be list repr of 807

	l := lSum
	for {
		fmt.Printf("%#v\n", l)
		if l.Next == nil {
			break
		}
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		// exhausted both lists at the same time
		// recursion has bottomed out, return back up the chain
		return nil
	}
	if l1 != nil && l2 == nil {
		// exhausted l2, the rest of the sum is just equal to the rest of l1
		// no need to recur further, return back up the chain
		return l1
	}
	if l1 == nil && l2 != nil {
		// exhausted l1, the rest of the sum is just equal to the rest of l2
		// no need to recur further, return back up the chain
		return l2
	}
	sum := l1.Val + l2.Val
	next := addTwoNumbers(l1.Next, l2.Next) // Let the rest of the recursion run
	if sum >= 10 {
		// sum can't be more than 18 (from adding two nines),
		// so we only need to modulo once to get back to a single digit
		sum %= 10
		// The next digit (which is one tens place higher) needs to be increased by 1
		// from whatever came back from the rest of the recursion.
		// If the next digit came back >= 10 (only possible case is next.Val = 9),
		// then the recursive call will hit this code path again and correct it.
		next = addTwoNumbers(next, &ListNode{Val: 1, Next: nil})
	}
	return &ListNode{Val: sum, Next: next}
}
