package medium

// https://leetcode.com/problems/partition-array-for-maximum-sum
func max(v ...int) int {
	if len(v) == 0 {
		return 0
	}
	max := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > max {
			max = v[i]
		}
	}
	return max
}

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := map[int]int{}
	var f func(int) int
	length := len(arr)
	f = func(idx int) int {
		if v, ok := dp[idx]; ok {
			return v
		}
		m := 0
		for i := 1; i <= k && i+idx <= length; i++ {
			newM := (max(arr[idx:i+idx]...) * i) + f(idx+i)
			if m < newM {
				m = newM
			}
		}
		dp[idx] = m
		return m
	}
	return f(0)
}
