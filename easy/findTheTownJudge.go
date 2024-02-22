package easy

// https://leetcode.com/problems/find-the-town-judge
func findJudge(n int, trust [][]int) int {
	pointed, points := make([]int, n+1), make([]int, n+1)
	for _, v := range trust {
		points[v[0]]++
		pointed[v[1]]++
	}
	for i := 1; i <= n; i++ {
		if points[i] == 0 && pointed[i] == n-1 {
			return i
		}
	}
	return -1
}
