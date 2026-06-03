package types3gpp

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"lite5gc/cmn/utils"
)

const (
	AmfSetIDBitLen   = 10
	AmfPointerBitLen = 6
	SizeofAmfSetID   = 2
)

// AmfIdentifier : AMF identifier
type AmfIdentifier struct {
	RegionId uint8                 //8  bit Length
	SetId    [SizeofAmfSetID]uint8 //10 bit Length
	Pointer  uint8                 //6  bit Length
}

func (p *AmfIdentifier) SetAmfRegionID(id uint8) {
	p.RegionId = id
}
func (p *AmfIdentifier) GetAmfRegionID() uint8 {
	return p.RegionId
}
func (p *AmfIdentifier) SetAmfSetID(setid []uint8) {
	copy(p.SetId[:], setid)
}
func (p *AmfIdentifier) GetAmfSetID() []uint8 {
	setID := make([]uint8, 2)
	copy(setID, p.SetId[:])
	return setID
}
func (p *AmfIdentifier) SetAmfSetIdU16(id uint16) {
	binary.BigEndian.PutUint16(p.SetId[:], id)
}
func (p *AmfIdentifier) GetAmfSetIdU16() uint16 {
	return binary.BigEndian.Uint16(p.SetId[:])
}
func (p *AmfIdentifier) SetAmfPointer(point uint8) {
	p.Pointer = point
}
func (p *AmfIdentifier) GetAmfPointer() uint8 {
	return p.Pointer
}

func (p AmfIdentifier) String() string {
	encodebytes := p.Encode()
	s := hex.EncodeToString(encodebytes)
	return s
}

func (p *AmfIdentifier) StoreWithString(s string) error {
	pbytes, err := hex.DecodeString(s)
	if err != nil {
		errStr := "failed to DecodeString:" + s
		return fmt.Errorf(errStr)
	}
	nasMsg := bytes.NewReader(pbytes)
	err = p.Decode(nasMsg)
	if err != nil {
		errStr := "failed to Decode AmfIdentifier"
		return fmt.Errorf(errStr)
	}
	return nil
}

func (p *AmfIdentifier) Encode() []byte {
	//rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "p", p.String())
	buf := make([]byte, 3) //three bytes

	octet1bit78 := p.SetId[0] << 6
	octet1bit16 := p.SetId[1] >> 2
	octet2bit78 := p.SetId[1] << 6

	buf[0] = p.RegionId
	buf[1] = octet1bit78 | octet1bit16
	buf[2] = octet2bit78 | p.Pointer
	return buf
}

func (p *AmfIdentifier) Decode(msgBuf *bytes.Reader) (err error) {
	// octet 1
	p.RegionId, err = msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read AmfIdentifier")
	}
	// octet 2
	octet2, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read AmfSetID[0]")
	}
	// octet 3
	octet3, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read AmfSetID[1] and AMF Pointer")
	}

	p.SetId[0] = octet2 >> 6
	p.SetId[1] = (octet2 << 2) | (octet3 >> 6)
	p.Pointer, err = utils.GetBitsValue(octet3, 1, 6)
	if err != nil {
		return fmt.Errorf("failed to read AMF Pointer")
	}
	return nil
}
