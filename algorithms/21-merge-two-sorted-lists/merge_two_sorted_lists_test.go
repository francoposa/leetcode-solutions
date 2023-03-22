package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	list1    *ListNode
	list2    *ListNode
	expected *ListNode
}{
	{
		list1: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		list2: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		expected: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
		},
	},
	{
		list1:    nil,
		list2:    nil,
		expected: nil,
	},
	{
		list1: nil,
		list2: &ListNode{
			Val:  0,
			Next: nil,
		},
		expected: &ListNode{
			Val:  0,
			Next: nil,
		},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, mergeTwoLists(test.list1, test.list2))
	}
}
