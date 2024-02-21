package medium

import (
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	if v := rangeBitwiseAnd(5, 7); v != 4 {
		t.Fatal(v)
	}
	if v := rangeBitwiseAnd(0, 0); v != 0 {
		t.Fatal(v)
	}
	if v := rangeBitwiseAnd(1, 2147483647); v != 0 {
		t.Fatal(v)
	}
	if v := rangeBitwiseAnd(100, 2147483644); v != 0 {
		t.Fatal(v)
	}
}
