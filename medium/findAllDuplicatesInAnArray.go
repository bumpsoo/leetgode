package medium

import "slices"

func findDuplicates(nums []int) []int {
	slices.Sort(nums)
	ret := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ret = append(ret, nums[i])
			i++
		}
	}
	return ret
}
