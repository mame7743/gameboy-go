package gameboygo

type Regsters struct {
	pc uint16
	sp uint16
	a byte
	f byte
}

func (reg *Regsters) af() uint16{
	return (uint16(reg.a) << 8) | uint16(reg.f)
}