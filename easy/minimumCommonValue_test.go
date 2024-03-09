package easy

import "testing"

func TestGetCommon(t *testing.T) {
	if v := getCommon([]int{1, 2, 3}, []int{2, 4}); v != 2 {
		t.Fatal(v)
	}
	if v := getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}); v != 2 {
		t.Fatal(v)
	}
}
