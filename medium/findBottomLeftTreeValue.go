package medium

// https://leetcode.com/problems/find-bottom-left-tree-value
func findBottomLeftValueHelp(root *TreeNode, depth int, ret, retDepth *int) {
	if root == nil {
		return
	}
	depth += 1
	if depth > *retDepth {
		*ret = root.Val
		*retDepth = depth
	}
	findBottomLeftValueHelp(root.Left, depth, ret, retDepth)
	findBottomLeftValueHelp(root.Right, depth, ret, retDepth)
}

func findBottomLeftValue(root *TreeNode) int {
	var ret, retDepth int
	findBottomLeftValueHelp(root, 0, &ret, &retDepth)
	return ret
}
