package easy

//https://leetcode.com/problems/unique-number-of-occurrences
func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	reverse := make(map[int]int, len(m))
	for k, occurrence := range m {
		if _, ok := reverse[occurrence]; !ok {
			reverse[occurrence] = k
		} else {
			return false
		}
	}
	return true
}
