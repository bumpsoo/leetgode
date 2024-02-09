package medium

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

// https://leetcode.com/problems/add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	addTwoNumbersHelp(head, l1, l2, 0)
	return head.Next
}

func addTwoNumbersHelp(head *ListNode, l1 *ListNode, l2 *ListNode, carry int) {
	if l1 != nil && l2 != nil {
		v, c := addTwoNumbersTrim(l1.Val + l2.Val + carry)
		head.Next = &ListNode{Val: v}
		addTwoNumbersHelp(head.Next, l1.Next, l2.Next, c)
	} else if l1 != nil {
		v, c := addTwoNumbersTrim(l1.Val + carry)
		head.Next = &ListNode{Val: v}
		addTwoNumbersHelp(head.Next, l1.Next, l2, c)
	} else if l2 != nil {
		v, c := addTwoNumbersTrim(l2.Val + carry)
		head.Next = &ListNode{Val: v}
		addTwoNumbersHelp(head.Next, l1, l2.Next, c)
	} else if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}
}

func addTwoNumbersTrim(i int) (int, int) {
	return i % 10, i / 10
}
