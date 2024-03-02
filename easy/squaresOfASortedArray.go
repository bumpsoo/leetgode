package easy

func sortedSquaresHelp(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	l := len(nums)
	left, right := 0, l-1
	retIdx := l - 1
	for i := 0; i < l; i++ {
		if sortedSquaresHelp(nums[left]) > sortedSquaresHelp(nums[right]) {
			ret[retIdx] = nums[left] * nums[left]
			left++
		} else {
			ret[retIdx] = nums[right] * nums[right]
			right--
		}
		retIdx--
	}
	return ret
}
