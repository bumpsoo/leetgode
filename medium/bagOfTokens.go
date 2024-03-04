package medium

import "slices"

// https://leetcode.com/problems/bag-of-tokens
func bagOfTokensScore(tokens []int, power int) int {
	slices.Sort(tokens)
	ret := 0
	l := len(tokens) - 1
	for i := 0; i <= l; {
		if tokens[i] <= power {
			power -= tokens[i]
			ret++
			i++
		} else {
			if i == l || ret == 0 {
				break
			}
			power += tokens[l]
			ret--
			l--
		}
	}
	return ret
}
