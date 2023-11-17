package gameboygo

type WRAM[0x2E00] byte

func NewWRAM() *WRAM {
	return new(WRAM)
}

func (w *WRAM) read( address uint16 ) byte {
	return w[address & 0x2DFF]
}

func (w *WRAM) write( address uint16, value byte ) {
	w[address & 0x2DFF] = value
}