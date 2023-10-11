package digitallogic

import "testing"

func Test_NewBit(t *testing.T) {

	bit := NewBit(false)

	if bit.GetBit() {
		t.Error("Bit does not equal 0")
	}

}

func TestBit_ToString(t *testing.T) {

	bit := NewBit(true)

	if bit.ToString() != "1" {
		t.Error("bit.ToString() does not equal 1")
	}

}
