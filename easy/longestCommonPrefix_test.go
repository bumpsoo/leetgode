package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	answer := []string{"fl", ""}
	for idx, args := range [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
	} {
		if v := longestCommonPrefix(args); v != answer[idx] {
			t.Fatal(v)
		}
	}
}
