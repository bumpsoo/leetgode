package medium

import "sort"

// https://leetcode.com/problems/divide-array-into-arrays-with-max-difference
func divideArray(nums []int, k int) [][]int {
	l := len(nums)
	if l%3 != 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	ret := make([][]int, l/3)
	for i := 0; i < l; i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		ret[i/3] = nums[i : i+3]
	}
	return ret
}
