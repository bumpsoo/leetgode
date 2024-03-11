package medium

import "slices"

// https://leetcode.com/problems/custom-sort-string
func customSortString(order string, s string) string {
	b := []byte(s)
	m := make(map[byte]int, len(order))
	for idx, key := range []byte(order) {
		m[key] = idx + 1
	}
	f := func(b byte) int {
		if val, ok := m[b]; ok {
			return val
		}
		return 100
	}
	slices.SortFunc(b, func(a, b byte) int {
		return f(a) - f(b)
	})
	return string(b)
}
