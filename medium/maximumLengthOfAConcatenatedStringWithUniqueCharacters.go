package medium

// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters
func maxLength(arr []string) int {
	isUnique := func(b []byte) bool {
		m := map[byte]struct{}{}
		for _, b := range b {
			if _, ok := m[b]; ok {
				return false
			} else {
				m[b] = struct{}{}
			}
		}
		return true
	}
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := 0
	length := len(arr)
	var loop func([]byte, int)
	loop = func(bytes []byte, idx int) {
		if !isUnique(bytes) {
			return
		}
		max = maxVal(max, len(bytes))
		for i := idx + 1; i < length; i++ {
			newBytes := append(bytes, []byte(arr[i])...)
			loop(newBytes, i)
		}
	}
	loop([]byte{}, -1)
	return max
}
