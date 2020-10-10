package simple

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2,3) Fail,Got %d.", result)
	}
}
