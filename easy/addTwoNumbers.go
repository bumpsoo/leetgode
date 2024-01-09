package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	addTwoNumbersHelper(head, l1, l2, 0)
	return head.Next
}

func addTwoNumbersHelper(
	head *ListNode, l1 *ListNode, l2 *ListNode, carry int,
) {
	if l1 != nil && l2 != nil {
		v, c := trimValue(l1.Val + l2.Val + carry)
		head.Next = &ListNode{Val: v}
		addTwoNumbersHelper(head.Next, l1.Next, l2.Next, c)
	} else if l1 != nil {
		v, c := trimValue(l1.Val + carry)
		head.Next = &ListNode{Val: v}
		addTwoNumbersHelper(head.Next, l1.Next, l2, c)
	} else if l2 != nil {
		v, c := trimValue(l2.Val + carry)
		head.Next = &ListNode{Val: v}
		addTwoNumbersHelper(head.Next, l1, l2.Next, c)
	} else if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}
}

func trimValue(i int) (int, int) {
	return i % 10, i / 10
}
