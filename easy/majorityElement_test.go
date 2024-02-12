package easy

import "testing"

func TestMajorityElement(t *testing.T) {
	tc := map[int][]int{
		3: {3, 2, 3},
		2: {2, 2, 1, 1, 1, 2, 2},
	}
	for k, v := range tc {
		if val := majorityElement(v); val != k {
			t.Fatal(val)
		}
	}
}
