package medium

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	arg := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	v := removeNthFromEnd(arg, 2)
	for _, val := range []int{1, 2, 3, 5} {
		if v.Val != val {
			t.Fatal(v)
		}
		v = v.Next
	}
	if v != nil {
		t.Fatal(v)
	}

	if v := removeNthFromEnd(&ListNode{Val: 1}, 1); v != nil {
		t.Fatal(v)
	}
	arg = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	v = removeNthFromEnd(arg, 1)
	if v.Val != 1 || v.Next != nil {
		t.Fatal(v)
	}
}
