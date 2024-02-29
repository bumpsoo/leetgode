package medium

import "testing"

func TestIsEventOddTree(t *testing.T) {
	if v := isEvenOddTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 12,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	},
	); v != true {
		t.Fatal()
	}
	if v := isEvenOddTree(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
		},
	}); v != false {
		t.Fatal()
	}
	if v := isEvenOddTree(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 7,
			},
		},
	}); v != false {
		t.Fatal()
	}
	if v := isEvenOddTree(&TreeNode{
		Val: 11,
		Left: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 10,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 11,
			},
		},
	}); v != true {
		t.Fatal()
	}
}
