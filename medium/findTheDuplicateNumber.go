package medium

// https://leetcode.com/problems/find-the-duplicate-number
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	ret := 0
	for {
		ret = nums[ret]
		slow = nums[slow]
		if ret == slow {
			break
		}
	}
	return ret
}
