package digitallogic

import (
	"testing"
)

func TestBit_Or(t *testing.T) {

	a := NewBit(true)
	b := NewBit(false)

	c := a.Or(&b)

	if !c.GetBit() {
		t.Error("Or does not equal 1 when a is 1 and b is 0")
	}

}

func TestBit_Nor(t *testing.T) {

	a := NewBit(true)
	b := NewBit(false)

	c := a.Nor(&b)

	if c.GetBit() {
		t.Error("Nor does not equal 0 when a is 1 and b is 0")
	}

}

func TestBit_And(t *testing.T) {

	a := NewBit(true)
	b := NewBit(false)

	c := a.And(&b)

	if c.GetBit() {
		t.Error("And does not equal 0 when a is 1 and b is 0")
		
	}

}

func TestBit_Nand(t *testing.T) {

	a := NewBit(true)
	b := NewBit(false)

	c := a.Nand(&b)
	
	if !c.GetBit() {
		t.Error("Nand does not equal 1 when a is 1 and b is 0")
	}

}

func TestBit_Not(t *testing.T) {

	a := NewBit(true)

	b := a.Not()
	
	if b.GetBit() {
		t.Error("Not does not equal 0 when input is 1")
	}

}
