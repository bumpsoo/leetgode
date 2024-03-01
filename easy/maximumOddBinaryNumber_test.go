package easy

import "testing"

func TestMaximumBinaryNumber(t *testing.T) {
	if v := maximumOddBinaryNumber("010"); v != "001" {
		t.Fatal(v)
	}
	if v := maximumOddBinaryNumber("0101"); v != "1001" {
		t.Fatal(v)
	}
}
