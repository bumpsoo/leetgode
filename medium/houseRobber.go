package medium_test

// https://leetcode.com/problems/house-robber
func rob(nums []int) int {
	m := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var f func(int) int
	memo := map[int]int{}
	f = func(idx int) int {
		if idx < 0 {
			return 0
		}
		if v, ok := memo[idx]; ok {
			return v
		}
		max := m(f(idx-2)+nums[idx], f(idx-1))
		memo[idx] = max
		return max
	}
	return f(len(nums) - 1)
}

