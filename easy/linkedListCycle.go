package easy

// https://leetcode.com/problems/linked-list-cycle
func hasCycle(head *ListNode) bool {
	hare, turtoise := head, head
	for hare != nil && hare.Next != nil {
		hare = hare.Next.Next
		turtoise = turtoise.Next
		if hare == turtoise {
			return true
		}
	}
	return false
}
