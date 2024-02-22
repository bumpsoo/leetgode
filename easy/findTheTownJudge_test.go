package easy

import "testing"

func TestFindJudge(t *testing.T) {
	if v := findJudge(2, [][]int{{1, 2}}); v != 2 {
		t.Fatal(v)
	}
	if v := findJudge(3, [][]int{{1, 3}, {2, 3}}); v != 3 {
		t.Fatal(v)
	}
	if v := findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}); v != -1 {
		t.Fatal(v)
	}
}
