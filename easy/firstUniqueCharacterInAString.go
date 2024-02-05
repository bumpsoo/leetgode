package easy

// https://leetcode.com/problems/first-unique-character-in-a-string
func firstUniqChar(s string) int {
	m := make(map[rune]int, len(s))
	a := make([]bool, len(s))
	for idx, r := range s {
		if jdx, ok := m[r]; ok {
			a[jdx] = true
			a[idx] = true
		} else {
			m[r] = idx
		}
	}
	for idx, f := range a {
		if !f {
			return idx
		}
	}
	return -1
}
