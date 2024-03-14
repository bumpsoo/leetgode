package medium

// https://leetcode.com/problems/binary-subarrays-with-sum
func numSubarraysWithSum(nums []int, goal int) int {
	helper := func(goal int) int {
		if goal < 0 {
			return 0
		}
		cnt := 0
		left, right := 0, 0
		sum := 0
		for right < len(nums) {
			sum += nums[right]
			for sum > goal {
				sum -= nums[left]
				left++
			}
			cnt += right - left + 1
			right++
		}
		return cnt
	}
	return helper(goal) - helper(goal-1)
}
