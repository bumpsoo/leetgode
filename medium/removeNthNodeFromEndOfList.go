package medium

// https://leetcode.com/problems/remove-nth-node-from-end-of-list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var length, idx int = -1, 0
	var f func(node *ListNode)
	f = func(node *ListNode) {
		if node == nil {
			length = idx
			return
		}
		idx++
		f(node.Next)
		idx--
		if length-n-1 == idx {
			node.Next = node.Next.Next
		}
	}
	f(head)
	if length <= n {
		return head.Next
	}
	return head
}
