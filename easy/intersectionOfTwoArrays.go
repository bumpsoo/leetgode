package easy

// https://leetcode.com/problems/intersection-of-two-arrays
func intersection(nums1 []int, nums2 []int) []int {
	s := map[int]int{}
	ret := []int{}
	for _, v := range nums1 {
		if _, ok := s[v]; !ok {
			s[v] = 1
		}
	}
	for _, v := range nums2 {
		if cnt, ok := s[v]; ok && cnt == 1 {
			s[v]++
			ret = append(ret, v)
		}
	}
	return ret
}
