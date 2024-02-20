package easy

import "testing"

func TestMissingNumbers(t *testing.T) {
	if v := missingNumber([]int{3, 0, 1}); v != 2 {
		t.Fatal(v)
	}
	if v := missingNumber([]int{0, 1}); v != 2 {
		t.Fatal(v)
	}
	if v := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}); v != 8 {
		t.Fatal(v)
	}
}
