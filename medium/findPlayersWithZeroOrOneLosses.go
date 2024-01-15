package medium

import "slices"

// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
func findWinners(matches [][]int) [][]int {
	count := make(map[int]int, len(matches))
	for _, match := range matches {
		for idx, v := range match {
			if _, ok := count[v]; ok {
				count[v] += idx
			} else {
				count[v] = idx
			}
		}
	}
	ret := [][]int{{}, {}}
	for k, v := range count {
		switch v {
		case 0:
			fallthrough
		case 1:
			ret[v] = append(ret[v], k)
		}
	}
	slices.Sort(ret[0])
	slices.Sort(ret[1])
	return ret
}
