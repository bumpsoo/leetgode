package medium

import "testing"

func TestMaxSubarrayLength(t *testing.T) {
	if v := maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2); v != 6 {
		t.Fatal(v)
	}
	if v := maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1); v != 2 {
		t.Fatal(v)
	}
	if v := maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4); v != 4 {
		t.Fatal(v)
	}
}
