package digitallogic

// OR

func (a *Bit) Or(b *Bit) *Bit {
	return MakeNewBit(a.GetBit() || b.GetBit())
}

// NOR

func (a *Bit) Nor(b *Bit) *Bit {
	return MakeNewBit(!(a.GetBit() || b.GetBit()))
}
