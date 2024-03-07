package easy

// https://leetcode.com/problems/middle-of-the-linked-list
func middleNode(head *ListNode) *ListNode {
	length, idx := 0, 0
	var ret *ListNode
	var f func(h *ListNode)
	f = func(h *ListNode) {
		if h == nil {
			length = idx
			return
		}
		idx++
		f(h.Next)
		idx--
		if length/2 == idx {
			ret = h
		}
	}
	f(head)
	return ret
}
