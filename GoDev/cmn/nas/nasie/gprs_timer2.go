package nasie

type GprsTimer2 struct {
	Uint       UintType
	TimerValue uint8
}

// refer to TS24.008  10.5.7.4a
func (p *GprsTimer2) Encode() []byte {
	var octet []byte

	gprs2Octet := byte(p.Uint) << 5
	gprs2Octet |= byte(p.TimerValue)

	octet = append(octet, gprs2Octet)

	return octet
}
