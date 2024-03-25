package medium

import (
	"slices"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	if v := findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}); !slices.Equal(v, []int{2, 3}) {
		t.Fatal(v)
	}
	if v := findDuplicates([]int{1, 1, 2}); !slices.Equal(v, []int{1}) {
		t.Fatal(v)
	}
	if v := findDuplicates([]int{1}); !slices.Equal(v, []int{}) {
		t.Fatal(v)
	}
}
