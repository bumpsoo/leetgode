package medium

// https://leetcode.com/problems/longest-substring-without-repeating-characters
func lengthOfLongestSubstring(s string) int {
	ret := 0
	set := map[byte]int{}
	i := 0
	for i < len(s) {
		if j, ok := set[s[i]]; ok {
			if l := len(set); l > ret {
				ret = l
			}
			i = j + 1
			set = map[byte]int{}
		} else {
			set[s[i]] = i
			i++
		}
	}
	if l := len(set); l > ret {
		return l
	}
	return ret
}
