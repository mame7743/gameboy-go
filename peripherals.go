package gameboygo


type Peripherals struct {
	brom BootROM
	wram WRAM
	hram HRAM
}

func (p *Peripherals) read(address uint16) byte {

	var value byte

	switch {
	case 0x0000 <= address && address <= 0x00FF:
		if p.brom.isActive(){
			value = p.brom.read(address)
		}else {
			value = 0xFF
		}
	case 0xC000 <= address && address <= 0xFDFF:
		value = p.wram.read(address)
	case 0xFF80 <= address && address <= 0xFFFE:
		value = p.hram.read(address)
	default:
		value = 0xFF
	}

	return value
}

func (p *Peripherals) write(address uint16, value byte)  {

	switch {
	case 0xC000 <= address && address <= 0xFDFF:
		p.wram.write(address, value)
	case 0xFF50 == address:
		p.brom.write(address, value)
	case 0xFF80 <= address && address <= 0xFFFE:
		p.hram.write(address, value)
	}

}
