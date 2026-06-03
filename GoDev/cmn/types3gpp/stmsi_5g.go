package types3gpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

type Stmsi5G struct {
	AmfSetID   [SizeofAmfSetID]uint8 //10 bit Length
	AmfPointer uint8                 //6  bit Length
	Tmsi       uint32
}

func (p *Stmsi5G) SetAmfSetID(setid []uint8) {
	copy(p.AmfSetID[:], setid)
}

func (p *Stmsi5G) String() string {
	return fmt.Sprintf("AmfSetID:%x%x, AmfPointer:%x, Tmsi:0x%x",
		p.AmfSetID[0], p.AmfSetID[1], p.AmfPointer, p.Tmsi)
}

func (p *Stmsi5G) GetAmfSetID() []uint8 {
	setID := make([]uint8, 2)
	copy(setID, p.AmfSetID[:])
	return setID
}

func (p *Stmsi5G) SetAmfPointer(point uint8) {
	p.AmfPointer = point
}
func (p *Stmsi5G) GetAmfPointer() uint8 {
	return p.AmfPointer
}
func (p *Stmsi5G) SetTmsi(Tmsi uint32) {
	p.Tmsi = Tmsi
}
func (p *Stmsi5G) GetTmsi() uint32 {
	return p.Tmsi
}

func (p *Stmsi5G) Encode() []byte {
	var buf []byte

	octet1bit78 := p.AmfSetID[0] << 6
	octet1bit16 := p.AmfSetID[1] >> 2
	octet2bit78 := p.AmfSetID[1] << 6

	amfSetPoint := make([]byte, 2) //two bytes
	amfSetPoint[0] = octet1bit78 | octet1bit16
	amfSetPoint[1] = octet2bit78 | p.AmfPointer

	buf = append(buf, amfSetPoint...)

	TmsiBytes := make([]byte, 4) //four bytes
	binary.BigEndian.PutUint32(TmsiBytes, p.Tmsi)
	buf = append(buf, TmsiBytes...)

	return buf
}

func (p *Stmsi5G) GetKey() string {
	key := p.Encode()
	return ByteString(key)
}

func (p *Stmsi5G) Decode(msgBuf *bytes.Reader) error {
	// octet 2 : LV : no T , Start directly from the second octet
	lenGuti := make([]byte, 2)
	err := binary.Read(msgBuf, binary.BigEndian, &lenGuti)
	if err != nil {
		rlogger.Trace(types.ModuleCmnNgap, rlogger.ERROR, nil, "fail to read len of guti")
		return fmt.Errorf("fail to read len of guti")
	}
	lenSuci := binary.BigEndian.Uint16(lenGuti)
	if lenSuci != 7 { //固定值 see 24501 Figure 9.11.3.4.5
		return fmt.Errorf("incorrect Length of STmsi5G")
	}
	// octet 4: not used yet
	msgBuf.ReadByte()

	// octet 5
	octet5, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read octet5")
	}
	// octet 6
	octet6, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read octet6")
	}

	p.AmfSetID[0] = octet5 >> 6
	p.AmfSetID[1] = (octet5 << 2) | (octet6 >> 6)
	p.AmfPointer, err = utils.GetBitsValue(octet6, 1, 6)
	if err != nil {
		return fmt.Errorf("failed to read AMF Pointer")
	}

	//octet 7-10
	TmsiOctet := make([]byte, MaxNumOfTMSI)
	err = binary.Read(msgBuf, binary.BigEndian, &TmsiOctet)
	if err != nil {
		return fmt.Errorf("failed to read homeNetworkID")
	}
	p.Tmsi = binary.BigEndian.Uint32(TmsiOctet)

	return nil
}
