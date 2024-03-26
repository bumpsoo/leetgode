package hard

import (
	"slices"
)

// https://leetcode.com/problems/first-missing-positive/submissions
func firstMissingPositiveSearch(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if target <= arr[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
func firstMissingPositive(nums []int) int {
	slices.Sort(nums)
	start := firstMissingPositiveSearch(nums, 1)
	if len(nums) <= start {
		return 1
	}
	i := 0
	check := 1
	for i+start < len(nums) {
		if nums[i+start] != check {
			break
		}
		for i+start < len(nums) && nums[i+start] == check {
			i++
		}
		check++
	}
	return check
}
