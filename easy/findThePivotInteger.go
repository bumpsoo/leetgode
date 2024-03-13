package easy

import "fmt"

// https://leetcode.com/problems/find-the-pivot-integer
func pivotInteger(n int) int {
	left := 0
	right := (n * (n + 1)) / 2
	for i := 1; i <= n; i++ {
		fmt.Println(left, right)
		left += i
		right -= i - 1
		if left == right {
			return i
		}
	}
	return -1
}
