package hard

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	if v := firstMissingPositive([]int{1, 2, 0}); v != 3 {
		t.Fatal(v)
	}
	if v := firstMissingPositive([]int{3, 4, -1, 1}); v != 2 {
		t.Fatal(v)
	}
	if v := firstMissingPositive([]int{7, 8, 9, 11, 12}); v != 1 {
		t.Fatal(v)
	}
	if v := firstMissingPositive([]int{-5}); v != 1 {
		t.Fatal(v)
	}
	if v := firstMissingPositive([]int{0, 2, 2, 1, 1}); v != 3 {
		t.Fatal(v)
	}
}
