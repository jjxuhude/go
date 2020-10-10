package simple

import "testing"

func TestSqrt(t *testing.T) {
	result := Sqrt(9)
	if result != 3 {
		t.Errorf("sqrt(9) fail. Got %d.", result)
	}
}
