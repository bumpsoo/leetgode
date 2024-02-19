package easy

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	if v := isPowerOfTwo(1); !v {
		t.Fatal(v)
	}
	if v := isPowerOfTwo(16); !v {
		t.Fatal(v)
	}
	if v := isPowerOfTwo(3); v {
		t.Fatal(v)
	}
}
