package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	if v := findMedianSortedArrays([]int{1, 3}, []int{2}); v != 2.0 {
		t.Fatal(v)
	}
	if v := findMedianSortedArrays([]int{1, 2}, []int{3, 4}); v != 2.5 {
		t.Fatal(v)
	}
	if v := findMedianSortedArrays([]int{}, []int{2, 3}); v != 2.5 {
		t.Fatal(v)
	}
}
