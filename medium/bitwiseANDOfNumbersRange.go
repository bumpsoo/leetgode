package medium

import (
	"math/bits"
)

// https://leetcode.com/problems/bitwise-and-of-numbers-range
func rangeBitwiseAnd(left int, right int) int {
	next := 64 - uint(bits.TrailingZeros(bits.Reverse(uint(left))))
	v := uint(1) << next
	if v < uint(right) {
		right = int(v)
	}
	ret := left
	left += 1
	for ; left <= right; left++ {
		ret = ret & left
	}
	return ret
}
