package easy

// https://leetcode.com/problems/majority-element
func majorityElement(nums []int) int {
	l := len(nums) / 2
	m := make(map[int]int, l)
	for _, val := range nums {
		m[val] += 1
		if m[val] > l {
			return val
		}
	}
	return -1
}
