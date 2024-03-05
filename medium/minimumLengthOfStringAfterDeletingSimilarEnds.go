package medium

// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends
func minimumLength(s string) int {
	ret := []byte(s)
	for left, right := 0, len(ret)-1; len(ret) > 0; left, right = 0, len(ret)-1 {
		ch := ret[left]
		if left == right || ch != ret[right] {
			break
		}
		for left < len(ret) && ret[left] == ch {
			left++
		}
		for right > 0 && ret[right] == ch {
			right--
		}
		if left > right {
			return 0
		}
		ret = ret[left : right+1]
	}
	return len(ret)
}
