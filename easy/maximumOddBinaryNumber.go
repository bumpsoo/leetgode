package easy

// https://leetcode.com/problems/maximum-odd-binary-number
func maximumOddBinaryNumber(s string) string {
	l := len(s)
	ret := make([]rune, l)
	ones, zeros := 0, 0
	for _, c := range s {
		if c == '1' {
			ret[ones] = '1'
			ones++
		} else {
			ret[l-zeros-1] = '0'
			zeros++
		}
	}
	ones--
	l--
	ret[ones], ret[l] = ret[l], ret[ones]
	return string(ret)
}
