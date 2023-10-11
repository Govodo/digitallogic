package digitallogic

import "testing"

func TestBit_Or(t *testing.T) {

	a := NewBit(true)
	b := NewBit(false)

	c := a.Or(&b)

	if !c.GetBit() {
		t.Error("Or does not equal 1 when a is 1 and b is 0")
	}

}
