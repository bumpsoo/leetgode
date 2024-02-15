package medium

import (
	"slices"
)

// https://leetcode.com/problems/find-polygon-with-the-largest-perimeter
func largestPerimeter(nums []int) int64 {
	slices.Sort(nums)
	var sum int64
	for _, val := range nums {
		sum += int64(val)
	}
	for i := len(nums) - 1; i >= 2; i-- {
		v := int64(nums[i])
		if sum-v > v {
			return sum
		}
		sum -= v
	}

	return -1
}
