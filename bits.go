package digitallogic

type Bit struct {
	State bool
}

type Byte = [8]Bit

func (bit *Bit) ToString() string {
	if bit.State {
		return "1"
	} else {
		return "0"
	}
}

func (bit *Bit) String() string {
	return bit.ToString()
}

func (bit *Bit) SetBit(state bool) {
	bit.State = state
}

func (bit *Bit) GetBit() bool {
	return bit.State
}

func NewBit(state bool) Bit {
	return Bit{State: state}
}

func MakeNewBit(state bool) *Bit {
	bit := new(Bit)
	bit = &Bit{State: state}
	return bit
}
