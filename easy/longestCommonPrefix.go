package easy

// https://leetcode.com/problems/longest-common-prefix
func longestCommonPrefixHelper(a, b string) int {
	i := 0
	for ; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return i
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var ret string
	idx := longestCommonPrefixHelper(strs[0], strs[1])
	if idx >= 0 {
		if len(strs[0]) > len(strs[1]) {
			ret = strs[0][0:idx]
		} else {
			ret = strs[1][0:idx]
		}
	}
	for i := 2; i < len(strs); i++ {
		idx := longestCommonPrefixHelper(ret, strs[i])
		if idx >= 0 {
			ret = strs[i][0:idx]
		}
	}
	return ret
}
