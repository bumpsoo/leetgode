package easy

// https://leetcode.com/problems/missing-number
func missingNumber(nums []int) int {
	n := len(nums)
	expected := (n * (n + 1)) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return expected - sum
}
