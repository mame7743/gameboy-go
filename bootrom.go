package gameboygo

type BootROM struct {
	rom []byte
	isActive bool
}

func NewBootROM(rom byte[]) *BootROM {
	b := &BootROM{rom: rom, isActive: true}
}

func (b *BootROM) read(address uint16) byte {
	return b.rom[address]
}

func (b *BootROM) isActive() bool {
	return b.isActive
}

func (b *BootROM) write(address uint16, value byte) {
	b.isActive &= value == 0;
}