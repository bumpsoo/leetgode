package medium

import "testing"

func TestLargestPerimeter(t *testing.T) {
	tc := map[int64][]int{
		15: {5, 5, 5},
		12: {1, 12, 1, 2, 5, 50, 3},
		-1: {5, 5, 50},
	}
	for ans, arg := range tc {
		if v := largestPerimeter(arg); v != ans {
			t.Fatal(v)
		}
	}
}
