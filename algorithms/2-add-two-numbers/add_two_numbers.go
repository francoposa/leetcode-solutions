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
