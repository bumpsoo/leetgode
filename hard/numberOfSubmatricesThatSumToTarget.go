package hard

// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target
func equal(a, b int) int {
	if a == b {
		return 1
	} else {
		return 0
	}
}

func prefixSum(matrix [][]int, target int) int {
	rowLen, colLen := len(matrix[0]), len(matrix)
	row, col := make([]int, rowLen), make([]int, colLen)
	row[0], col[0] = matrix[0][0], matrix[0][0]
	ret := 0 + equal(row[0], target)
	for i := 1; i < rowLen; i++ {
		row[i] = row[i-1] + matrix[0][i]
		ret += equal(row[i], target)
	}
	for i := 1; i < colLen; i++ {
		col[i] = col[i-1] + matrix[i][0]
		ret += equal(col[i], target)
	}
	for i := 1; i < colLen; i++ {
		old := row[0]
		row[0] = col[i]
		for j := 1; j < rowLen; j++ {
			olld := row[j]
			row[j] = row[j-1] + row[j] + matrix[i][j] - old
			old = olld
			ret += equal(row[j], target)
		}
	}
	return ret
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	ret := 0
	for i := range matrix {
		sub := matrix[i:]
		for j := range matrix[i] {
			subb := make([][]int, len(sub))
			for k, v := range sub {
				subb[k] = v[j:]
			}
			ret += prefixSum(subb, target)
		}
	}
	return ret
}
