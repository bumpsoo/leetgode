package medium

import "testing"

func TestFindBottomLeftValue(t *testing.T) {
	if v := findBottomLeftValue(
		&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
	); v != 1 {
		t.Fatal(v)
	}

	if v := findBottomLeftValue(
		&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 7}},
				Right: &TreeNode{Val: 6},
			},
		},
	); v != 7 {
		t.Fatal(v)
	}
}
