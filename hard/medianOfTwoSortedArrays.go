package hard

// https://leetcode.com/problems/median-of-two-sorted-arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	merged := make([]int, l/2+1)
	idx1, idx2 := 0, 0
	for i := 0; i <= l/2; i++ {
		if len(nums1) == idx1 {
			merged[i] = nums2[idx2]
			idx2++
			continue
		}
		if len(nums2) == idx2 {
			merged[i] = nums1[idx1]
			idx1++
			continue
		}
		if n1, n2 := nums1[idx1], nums2[idx2]; n1 < n2 {
			merged[i] = n1
			idx1++
		} else {
			merged[i] = n2
			idx2++
		}
	}
	if l%2 == 0 {
		return (float64(merged[l/2]) + float64(merged[l/2-1])) / float64(2)
	}
	return float64(merged[l/2])
}
