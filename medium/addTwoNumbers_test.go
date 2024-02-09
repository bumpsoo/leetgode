package medium

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	tests := [][]*ListNode{
		{
			&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			&ListNode{Val: 0},
			&ListNode{Val: 0},
			&ListNode{Val: 0},
		},
		{
			&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
			&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			&ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}},
		},
	}
	for _, node := range tests {
		ret := addTwoNumbers(node[0], node[1])
		for ret != nil {
			if ret.Val != node[2].Val {
				t.Fatal(ret)
			}
			ret = ret.Next
			node[2] = node[2].Next
		}
	}
}
