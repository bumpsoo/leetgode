package easy

import "math/bits"

// https://leetcode.com/problems/power-of-two
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return bits.OnesCount(uint(n)) == 1
}
