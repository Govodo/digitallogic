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

func TestBit_Xor(t *testing.T) {

	t.Run("XOR a: 1, b: 1", func(t *testing.T) {
		a := NewBit(true)
		b := NewBit(true)

		c := a.Xor(&b)

		if c.GetBit() {
			t.Error("XOR does not equal 0 when a is 1 and b is 1")
		}
	})

	t.Run("XOR a: 0, b: 0", func(t *testing.T) {
		a := NewBit(false)
		b := NewBit(false)

		c := a.Xor(&b)

		if c.GetBit() {
			t.Error("XOR does not equal 0 when a is 0 and b is 0")
		}
	})

	t.Run("XOR a: 1, b: 0", func(t *testing.T) {
		a := NewBit(true)
		b := NewBit(false)

		c := a.Xor(&b)

		if !c.GetBit() {
			t.Error("XOR does not equal 1 when a is 1 and b is 0")
		}
	})

	t.Run("XOR a: 0, b: 1", func(t *testing.T) {
		a := NewBit(false)
		b := NewBit(true)

		c := a.Xor(&b)

		if !c.GetBit() {
			t.Error("XOR does not equal 1 when a is 0 and b is 1")
		}
	})

}

func TestBit_Nxor(t *testing.T) {

	t.Run("NXOR a: 1, b: 1", func(t *testing.T) {
		a := NewBit(true)
		b := NewBit(true)

		c := a.Nxor(&b)

		if !c.GetBit() {
			t.Error("NXOR does not equal 1 when a is 1 and b is 1")
		}
	})

	t.Run("NXOR a: 0, b: 0", func(t *testing.T) {
		a := NewBit(false)
		b := NewBit(false)

		c := a.Nxor(&b)

		if !c.GetBit() {
			t.Error("NXOR does not equal 1 when a is 0 and b is 0")
		}
	})

	t.Run("NXOR a: 1, b: 0", func(t *testing.T) {
		a := NewBit(true)
		b := NewBit(false)

		c := a.Nxor(&b)

		if c.GetBit() {
			t.Error("NXOR does not equal 0 when a is 1 and b is 0")
		}
	})

	t.Run("NXOR a: 0, b: 1", func(t *testing.T) {
		a := NewBit(false)
		b := NewBit(true)

		c := a.Nxor(&b)

		if c.GetBit() {
			t.Error("NXOR does not equal 0 when a is 0 and b is 1")
		}
	})

}
