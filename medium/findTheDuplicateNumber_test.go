package medium

import "testing"

func TestFindDuplicate(t *testing.T) {
	if v := findDuplicate([]int{1, 3, 4, 2, 2}); v != 2 {
		t.Fatal(v)
	}
	if v := findDuplicate([]int{3, 1, 3, 4, 2}); v != 2 {
		t.Fatal(v)
	}
	if v := findDuplicate([]int{3, 3, 3, 3, 3}); v != 2 {
		t.Fatal(v)
	}
}
