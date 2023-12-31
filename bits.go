package digitallogic

type Bit struct {
	State bool
}

const BITS_IN_A_BYTE int = 8

type Byte = [BITS_IN_A_BYTE]Bit

func (bit *Bit) String() string {
	if bit.State {
		return "1"
	} else {
		return "0"
	}
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
	return &Bit{State: state}
}
