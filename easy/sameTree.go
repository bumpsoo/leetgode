package easy

// https://leetcode.com/problems/same-tree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p != nil) != (q != nil) {
		return false
	}
	if p == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}
