package types3gpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

const MaxNumOfTMSI = 4

// Guti5G : 5G Globally Unique Temporary UE Identity 24.501 9.11.3.4
type Guti5G struct {
	plmn  PlmnID
	amfID AmfIdentifier
	tmsi  uint32
}

func (p *Guti5G) String() string {
	return fmt.Sprintf("plmn:%s, amfId:%s, tmsi:0x%x",
		&(p.plmn), &(p.amfID), p.tmsi)
}

func CreateATestGUTI() Guti5G {
	return Guti5G{
		plmn:  PlmnID{[SizeofPlmnID]byte{0x11, 0x11, 0x11}, false},
		amfID: AmfIdentifier{0x11, [SizeofAmfSetID]uint8{0x11, 0x11}, 0x11},
		tmsi:  0x11111111,
	}
}

func (p *Guti5G) SetPlmn(plmn *PlmnID) {
	p.plmn = *plmn
}
func (p *Guti5G) GetPlmn() *PlmnID {
	return &(p.plmn)
}
func (p *Guti5G) SetAmfID(amfid *AmfIdentifier) {
	p.amfID = *amfid
}
func (p *Guti5G) GetAmfID() *AmfIdentifier {
	return &(p.amfID)
}
func (p *Guti5G) SetTmsi(tmsi uint32) {
	p.tmsi = tmsi
}
func (p *Guti5G) GetTmsi() uint32 {
	return p.tmsi
}

func (p *Guti5G) SetPlmnByte(arr []byte) error {
	return nil
}

/*
octet 3 (Bits)
8   7   6   5   4                   3       2     1
1	1	1	1	odd/even_indication	Type_of_identity
octet 3

Type_of_identity (Bits)
3	2	1
0	0	0	No identity (NOTE)
0	0	1	SUCI
1	1	0	5G-GUTI
0	1	1	IMEI
1	0	0	5G-S-TMSI
1	0	1	IMEISVI
All other values are reserved.

Odd/even_indication (Bits)
4
0			even number of identity digits and also when the GUTI or 5G-S-TMSI is used
1			odd number of identity digits
*/
func (p *Guti5G) Encode() []byte {
	// Function Encode only encode the value of Guti5G,  start from octet4
	var buf []byte

	// octet 4~6
	plmnValue := p.plmn.GetValue(BigEndian)
	buf = append(buf, plmnValue[:]...)

	// octet 7~9
	amfIdBytes := p.amfID.Encode()
	buf = append(buf, amfIdBytes...)

	// octet 10~13
	tmsiBytes := make([]byte, MaxNumOfTMSI)
	binary.BigEndian.PutUint32(tmsiBytes, p.tmsi)
	buf = append(buf, tmsiBytes...)

	return buf
}

func (p *Guti5G) GetStmsi() Stmsi5G {
	stmsi := Stmsi5G{}
	stmsi.AmfSetID = p.amfID.SetId
	stmsi.AmfPointer = p.amfID.Pointer
	stmsi.Tmsi = p.tmsi
	return stmsi
}

func (p *Guti5G) GetStmsiKey() string {
	stmsi := Stmsi5G{}
	stmsi.AmfSetID = p.amfID.SetId
	stmsi.AmfPointer = p.amfID.Pointer
	stmsi.Tmsi = p.tmsi
	return stmsi.GetKey()
}

func ByteString(p []byte) string {
	var s string
	for i := 0; i < len(p); i++ {
		s += fmt.Sprintf("%02x", p[i]) //以防key重复现象
	}
	return s
}

func (p *Guti5G) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	// octet 2 : LV : no T , Start directly from the second octet
	lenGuti := make([]byte, 2)
	err := binary.Read(msgBuf, binary.BigEndian, &lenGuti)
	if err != nil {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "fail to read len of guti")
		return fmt.Errorf("fail to read len of guti")
	}
	len := binary.BigEndian.Uint16(lenGuti)
	if len != 11 { //固定值 see 24501 9.11.3.4.1
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "incorrect length of Guti5G")
		return fmt.Errorf("incorrect length of Guti5G")
	}
	// octet 3: not used yet
	msgBuf.ReadByte()

	// octet 4~6
	err = binary.Read(msgBuf, binary.BigEndian, &(p.plmn.value))
	if err != nil {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "failed to read plmn")
		return fmt.Errorf("failed to read plmn")
	}
	p.plmn.SetValue(p.plmn.value[:], BigEndian) //注意，这句不能少，不然plmn的值对不上
	// octet 7~9
	p.amfID = *new(AmfIdentifier)
	err = p.amfID.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "failed to decode amfID")
		return fmt.Errorf("failed to decode amfID")
	}
	//octet 10~13
	tmsiOctet := make([]byte, MaxNumOfTMSI)
	err = binary.Read(msgBuf, binary.BigEndian, &tmsiOctet)
	if err != nil {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "failed to read homeNetworkID")
		return fmt.Errorf("failed to read homeNetworkID")
	}
	p.tmsi = binary.BigEndian.Uint32(tmsiOctet)
	return nil
}
