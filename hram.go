package gameboygo

type HRAM[0x80] byte

func NewHRAM() *HRAM {
	return new(HRAM)
}

func (h *HRAM) read( address uint16 ) byte {
	return h[address & 0x7f]
}

func (h *HRAM) write( address uint16, value byte ) {
	h[address & 0x7f] = value
}