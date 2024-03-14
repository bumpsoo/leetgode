package medium

import "testing"

func TestNumSubarrayWithSum(t *testing.T) {
	if v := numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2); v != 4 {
		t.Fatal(v)
	}
	if v := numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0); v != 15 {
		t.Fatal(v)
	}
}
