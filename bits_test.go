package digitallogic

import "testing"

func Test_NewBit(t *testing.T) {

	bit := NewBit(false)

	if bit.GetBit() {
		t.Error("Bit does not equal 0")
	}

}

func Test_MakeNewBit(t *testing.T) {

	bit := MakeNewBit(false)

	if bit.GetBit() {
		t.Error("Bit does not equal 0")
	}

}

func TestBit_String(t *testing.T) {

	bit := NewBit(true)

	if bit.String() != "1" {
		t.Error("bit.String() does not equal 1")
	}

}
