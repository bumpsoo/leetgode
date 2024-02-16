package medium

import "testing"

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	if v := findLeastNumOfUniqueInts([]int{5, 5, 4}, 1); v != 1 {
		t.Fatal(v)
	}
	if v := findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3); v != 2 {
		t.Fatal(v)
	}
}
