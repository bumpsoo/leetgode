package medium

// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency
func maxSubarrayLength(nums []int, k int) int {
	ret, currLen, start := 0, 0, 0
	set := map[int]int{}
	for _, num := range nums {
		if _, ok := set[num]; ok {
			set[num]++
		} else {
			set[num] = 1
		}
		for set[num] > k && start < len(nums) {
			set[nums[start]]--
			currLen--
			if nums[start] == num {
				start++
				break
			}
			start++
		}
		currLen++
		if ret < currLen {
			ret = currLen
		}
	}
	return ret
}
