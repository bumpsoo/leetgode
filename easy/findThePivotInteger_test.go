package easy

import "testing"

func TestPivotInteger(t *testing.T) {
	if v := pivotInteger(8); v != 6 {
		t.Fatal(v)
	}
	if v := pivotInteger(1); v != 1 {
		t.Fatal(v)
	}
	if v := pivotInteger(4); v != -1 {
		t.Fatal(v)
	}
}
