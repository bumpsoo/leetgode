package easy

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	if v := intersection([]int{1, 2, 2, 1}, []int{2, 2}); !slices.Equal(v, []int{2}) {
		t.Fatal(v)
	}
	if v := intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}); !slices.Equal(v, []int{9, 4}) {
		t.Fatal(v)
	}
}
