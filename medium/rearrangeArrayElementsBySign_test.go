package medium

import (
	"slices"
	"testing"
)

func TestRearrangeArray(t *testing.T) {
	tc := [][][]int{
		{
			{3, 1, -2, -5, 2, -4},
			{3, -2, 1, -5, 2, -4},
		},
		{
			{-1, 1},
			{1, -1},
		},
	}
	for _, v := range tc {
		ans := rearrangeArray(v[0])
		if !slices.Equal(ans, v[1]) {
			t.Fatal(ans)
		}
	}
}
