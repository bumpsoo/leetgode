package medium

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree
func pseudoPalindromicPaths(root *TreeNode) int {
	ret := 0
	var search func(map[int]struct{}, *TreeNode)
	search = func(vals map[int]struct{}, head *TreeNode) {
		if _, ok := vals[head.Val]; ok {
			delete(vals, head.Val)
		}
		if head.Left == nil && head.Right == nil {
			if len(vals) <= 1 {
				ret += 1
			}
			return
		}
		vals[head.Val] = struct{}{}
		if head.Left != nil {
			search(vals, head.Left)
		}
		if head.Right != nil {
			search(vals, head.Right)
		}
	}
	search(map[int]struct{}{}, root)
	return ret
}
