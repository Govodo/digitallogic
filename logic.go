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
