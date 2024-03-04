package medium

import "testing"

func TestBagOfTokensScore(t *testing.T) {
	if v := bagOfTokensScore([]int{100}, 50); v != 0 {
		t.Fatal(v)
	}
	if v := bagOfTokensScore([]int{200, 100}, 150); v != 1 {
		t.Fatal(v)
	}
	if v := bagOfTokensScore([]int{100, 200, 300, 400}, 200); v != 2 {
		t.Fatal(v)
	}
}
