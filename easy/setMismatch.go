package easy

// https://leetcode.com/problems/set-mismatch
func findErrorNums(nums []int) []int {
	m := make(map[int]struct{}, len(nums))
	duplicated := -1
	factorial := 0
	sum := 0
	for i, v := range nums {
		if _, ok := m[v]; ok {
			duplicated = v
		} else {
			m[v] = struct{}{}
		}
		factorial += i + 1
		sum += v
	}
	return []int{duplicated, duplicated + factorial - sum}
}
