package easy

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1, l2 := []int{}, []int{}
	getLeaves(root1, &l1)
	getLeaves(root2, &l2)
	return slices.Compare(l1, l2) == 0
}

func getLeaves(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, root.Val)
	}
	getLeaves(root.Left, ret)
	getLeaves(root.Right, ret)
}
