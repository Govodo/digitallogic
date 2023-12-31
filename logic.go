package digitallogic

// OR

func (a *Bit) Or(b *Bit) *Bit {
	return MakeNewBit(a.GetBit() || b.GetBit())
}

// NOR

func (a *Bit) Nor(b *Bit) *Bit {
	return MakeNewBit(!(a.GetBit() || b.GetBit()))
}

// AND

func (a *Bit) And(b *Bit) *Bit {
	return MakeNewBit(a.GetBit() && b.GetBit())
}

// NAND

func (a *Bit) Nand(b *Bit) *Bit {
	return MakeNewBit(!(a.GetBit() && b.GetBit()))
}

// NOT

func (a *Bit) Not() *Bit {
	return MakeNewBit(!a.GetBit())
}

// XOR

func (a *Bit) Xor(b *Bit) *Bit {
	c := a.And(b.Not())
	d := b.And(a.Not())
	e := c.Or(d)
	return e
}

// NXOR

func (a *Bit) Nxor(b *Bit) *Bit {
	return a.Xor(b).Not()
}
