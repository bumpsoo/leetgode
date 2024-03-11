package medium

import (
	"testing"
)

func TestCustomSortString(t *testing.T) {
	v := customSortString("cba", "abcd")
	if !(v == "dcba" || v == "cdba" || v == "cbda" || v == "cbad") {
		t.Fatal(v)
	}
	v = customSortString("bcafg", "abcd")
	if !(v == "bacd" || v == "bcda" || v == "bcad") {
		t.Fatal(v)
	}
}
