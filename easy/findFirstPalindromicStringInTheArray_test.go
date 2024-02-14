package easy

import "testing"

func TestFirstPalindrome(t *testing.T) {
	tc := map[string][]string{
		"ada":     {"abc", "car", "ada", "racecar", "cool"},
		"racecar": {"notapalindrome", "racecar"},
		"":        {"def", "ghi"},
	}
	for val, arg := range tc {
		if ret := firstPalindrome(arg); ret != val {
			t.Fatal(ret)
		}
	}
}
