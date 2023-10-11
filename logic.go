package digitallogic

// OR

func (a *Bit) Or(b *Bit) Bit {
	return NewBit(a.GetBit() || b.GetBit())
}
