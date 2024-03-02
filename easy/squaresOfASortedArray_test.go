package easy

import (
	"slices"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	if v := sortedSquares([]int{-4, -1, 0, 3, 10}); !slices.Equal(v, []int{0, 1, 9, 16, 100}) {
		t.Fatal(v)
	}
	if v := sortedSquares([]int{-7, -3, 2, 3, 11}); !slices.Equal(v, []int{4, 9, 9, 49, 121}) {
		t.Fatal(v)
	}

}
