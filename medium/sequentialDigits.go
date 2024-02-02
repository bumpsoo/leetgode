package medium

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/sequential-digits
func seq(s, l int) int {
	ret := 0
	for i := 0; i < l; i++ {
		ret += (s + i) * int(math.Pow10(l-i-1))
	}
	return ret
}

func sequentialDigits(low, high int) []int {
	length := 0
	first := low
	for first > 9 {
		first /= 10
		length++
	}
	length++
	seqs := []int{}
	var f func(int)
	f = func(v int) {
		if low > v || v > high {
			return
		}
		seqs = append(seqs, v)
		mod := (v % 10) + 1
		if mod > 9 {
			return
		}
		newV := (v * 10) + mod
		if newV <= high {
			f(newV)
		}
	}
	for i := 1; i < 10; i++ {
		newL := length
		if i < first {
			newL += 1
		}
		if i+newL <= 10 {
			f(seq(i, newL))
		}
	}
	sort.Ints(seqs)
	return seqs
}
