package nasie

// 24.501 9.11.3.54
//type UeSecurityCapability struct {
//	Ea5g byte
//	Ia5g byte
//	Ea4g byte
//	Ia4g byte
//}
//
//func (p *UeSecurityCapability) Reset() {
//	p.Ea5g = 0
//	p.Ia5g = 0
//	p.Ea4g = 0
//	p.Ia4g = 0
//}
//func (p *UeSecurityCapability) Decode(msgBuf *bytes.Reader) error {
//	length, _ := msgBuf.ReadByte()
//	if length < 2 || length > 8 {
//		rlogger.Trace(types.ERROR, "invalid length for "+
//			"Ie ue security capability (%d)", length)
//		return nas.ErrInvalidIeLength
//	}
//	octet, _ := msgBuf.ReadByte()
//	p.Ea5g = octet
//	octet, _ = msgBuf.ReadByte()
//	p.Ia5g = octet
//
//	if length > 2 {
//		octet, _ = msgBuf.ReadByte()
//		p.Ea4g = octet
//
//	}
//	if length > 3 {
//		octet, _ = msgBuf.ReadByte()
//		p.Ia4g = octet
//	}
//
//	if length > 4 {
//		leftBytes := make([]byte, length-4)
//		binary.Read(msgBuf, binary.BigEndian, leftBytes)
//	}
//	return nil
//}
