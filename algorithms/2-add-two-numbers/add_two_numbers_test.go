package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	l1  *ListNode
	l2  *ListNode
	sum *ListNode
}{
	{
		l1: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}, // list repr of 342
		l2: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}, // list repr of 465
		sum: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		}, // list repr of 807
	},
}

func TestAddTwoNumbersIter(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.sum, addTwoNumbersIter(test.l1, test.l2))
	}
}

func TestAddTwoNumbersRecur(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.sum, addTwoNumbersRecur(test.l1, test.l2))
	}
}
