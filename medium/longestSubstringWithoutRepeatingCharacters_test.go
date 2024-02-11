package medium

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tc := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
	}
	for k, v := range tc {
		if val := lengthOfLongestSubstring(k); val != v {
			t.Fatal(k, val)
		}
	}
}
