package easy

import "testing"

func TestIsSameTree(t *testing.T) {
	v := isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
	)
	if !v {
		t.Fatal()
	}
	v = isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
	)
	if v {
		t.Fatal()
	}
	v = isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
	)
	if v {
		t.Fatal()
	}
}
