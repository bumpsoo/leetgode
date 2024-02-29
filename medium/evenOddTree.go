package medium

// https://leetcode.com/problems/even-odd-tree/submissions
func isEvenOddTree(root *TreeNode) bool {
	q := NewQ[*TreeNode]()
	q.Push(root)
	depth := 1
	for l := len(*q); l > 0; l = len(*q) {
		var oldNode, newNode *TreeNode
		oldNode = q.Pop()
		if oldNode == nil {
			continue
		}
		if oldNode.Val%2 != depth%2 {
			return false
		}
		q.Push(oldNode.Left)
		q.Push(oldNode.Right)
		for i := 1; i < l; i++ {
			newNode = q.Pop()
			if newNode == nil {
				continue
			}
			if newNode.Val%2 != depth%2 {
				return false
			}
			if oldNode != nil {
				odd := depth%2 == 1
				if odd && oldNode.Val >= newNode.Val {
					return false
				}
				if !odd && oldNode.Val <= newNode.Val {
					return false
				}
			}
			q.Push(newNode.Left)
			q.Push(newNode.Right)
			oldNode = newNode
		}
		depth++
	}
	return true
}
