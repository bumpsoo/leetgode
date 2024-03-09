package easy

// https://leetcode.com/problems/minimum-common-value
func getCommon(nums1 []int, nums2 []int) int {
	idx1, idx2 := 0, 0
	ret := -1
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] == nums2[idx2] {
			ret = nums1[idx1]
			break
		}
		if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			idx1++
		}
	}
	return ret
}
