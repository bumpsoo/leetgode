package easy

// https://leetcode.com/problems/climbing-stairs
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	one, two := 1, 2
	return func(n int) int {
		for i := 3; i < n; i++ {
			two, one = one+two, two
		}
		return one + two
	}(n)
}
