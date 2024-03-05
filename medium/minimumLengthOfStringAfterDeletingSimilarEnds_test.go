package medium

import "testing"

func TestMinimumLength(t *testing.T) {
	if v := minimumLength("ca"); v != 2 {
		t.Fatal(v)
	}
	if v := minimumLength("cabaabac"); v != 0 {
		t.Fatal(v)
	}
	if v := minimumLength("aabccabba"); v != 3 {
		t.Fatal(v)
	}
}
