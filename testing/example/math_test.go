package math

import "testing"

func TestDouble(t *testing.T) {
	v := Double(21)
	if v != 42 {
		t.Errorf("Got %v, expected %v\n", v, 42)
	}
}
