package medium

// https://leetcode.com/problems/rearrange-array-elements-by-sign
func rearrangeArray(nums []int) []int {
	posCnt, negCnt := 0, 0
	ret := make([]int, len(nums))
	for _, num := range nums {
		if num > 0 {
			ret[posCnt*2] = num
			posCnt++
		} else {
			ret[negCnt*2+1] = num
			negCnt++
		}
	}
	return ret
}
