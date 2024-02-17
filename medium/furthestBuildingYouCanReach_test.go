package medium

import (
	"testing"
)

func TestFurthestBuilding(t *testing.T) {
	if v := furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1); v != 4 {
		t.Fatal(v)
	}
	if v := furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2); v != 7 {
		t.Fatal(v)
	}
	if v := furthestBuilding([]int{14, 3, 19, 3}, 17, 0); v != 3 {
		t.Fatal(v)
	}
}
