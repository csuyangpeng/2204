package types3gpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
	"strconv"
)

const (
	MinLenofRouteInd         = 1
	MaxLenofRouteInd         = 4
	MaxProtectSchemeId uint8 = 15
	MinProtectSchemeId uint8 = 0
	//MaxHomeNwPubKeyId  uint8 = 255
	//MinHomeNwPubKeyId  uint8 = 0
	MaxDigitNumofPlmn = 6
	MinDigitNumofPlmn = 5
)

const (
	NullScheme uint8 = 0x0
	ProfileA   uint8 = 0x1
	ProfileB   uint8 = 0x2
)

// Suci  : Subscription Concealed Identifier
type Suci struct {
	supiType            SupiType
	homeNetworkID       []byte
	routeIndicator      []uint8
	protectSchemeID     uint8
	homeNetworkPubKeyID uint8
	schemeOutput        []byte
	schemeOutputA       SchemeOutputA
}

// Figure 2.2B-3: Scheme Output for Elliptic Curve Integrated Encryption Scheme Profile A
type SchemeOutputA struct {
	//ECC ephemeral public key,32 bytes
	EphPublicKey []byte
	//ciphertext value, 5 bytes
	Ciphertext []byte
	//MAC tag value,8 bytes
	MacTag []byte
}

func (p *Suci) reset() {
	p.supiType = InvalidType
	p.homeNetworkID = nil
	p.routeIndicator = nil
	p.protectSchemeID = NullScheme
	p.homeNetworkPubKeyID = 0
	p.schemeOutput = nil
}

func (p *Suci) SetSupiType(t SupiType) error {
	if t != IMSIType && t != NAIType {
		return fmt.Errorf("invalid supi type(0-imsi, 1-nai), %d", t)
	}
	p.supiType = t
	return nil
}

func (p *Suci) GetSupiType() SupiType {
	return p.supiType
}

func (p *Suci) SetHomeNetworkId(val []byte, supiType SupiType) error {
	err := p.SetSupiType(supiType)
	if err != nil {
		return err
	}

	if p.supiType == IMSIType {
		length := len(val)
		if length > MaxDigitNumofPlmn ||
			length < MinDigitNumofPlmn {
			return fmt.Errorf("invalid HomeNetworkId. "+
				"Should be MCC+MNC, len(%d-%d), but length is %d",
				MinDigitNumofPlmn, MaxDigitNumofPlmn, length)
		}
		if !utils.IsDigitString(string(val)) {
			return fmt.Errorf("input value should be digits. but value is %s", string(val))
		}

		//clear the slice
		p.homeNetworkID = p.homeNetworkID[0:0]
		//set the value
		p.homeNetworkID = append(p.homeNetworkID, val...)
	} else {
		if utils.IsValidNaiString(string(val)) {
			p.homeNetworkID = p.homeNetworkID[0:0]
			p.homeNetworkID = append(p.homeNetworkID, val...)
		} else {
			return fmt.Errorf("invalid formate for the NAI")
		}
	}

	return nil
}

func (p *Suci) GetHomeNetworkId() []byte {
	return p.homeNetworkID
}

func (p *Suci) SetRouteIndicator(rind []uint8) error {
	length := uint8(len(rind))
	if length < MinLenofRouteInd ||
		length > MaxLenofRouteInd {
		return fmt.Errorf("invalid length of route "+
			"indicator[%d-%d], but input rind length %d",
			MinLenofRouteInd, MaxLenofRouteInd, length)
	}

	if !utils.IsDigitString(string(rind)) {
		return fmt.Errorf("invalid value of route ind, "+
			"should be digits.but %s", string(rind))
	}

	p.routeIndicator = p.routeIndicator[0:0]
	p.routeIndicator = append(p.routeIndicator, rind...)

	return nil
}

func (p *Suci) GetRouteIndicator() []uint8 {
	return p.routeIndicator
}

func (p *Suci) SetProtectSchemeId(sid uint8) error {
	if sid > MaxProtectSchemeId ||
		sid < MinProtectSchemeId {
		return fmt.Errorf("invalid protect Scheme Id, invalid "+
			"range (%d-%d), but vallue is %d",
			MinProtectSchemeId, MaxProtectSchemeId, sid)
	}
	p.protectSchemeID = sid
	return nil
}

func (p *Suci) GetProtectSchemeId() uint8 {
	return p.protectSchemeID
}

func (p *Suci) SetHomeNwPubKeyId(pid uint8) {
	p.homeNetworkPubKeyID = pid
}

func (p *Suci) GetHomeNwPubKeyId() uint8 {
	return p.homeNetworkPubKeyID
}

func (p *Suci) SetSchemeOutput(output []byte) {
	p.schemeOutput = p.schemeOutput[0:0]
	p.schemeOutput = append(p.schemeOutput, output...)
}

func (p *Suci) GetSchemeOutput() []byte {
	return p.schemeOutput
}

func (p *Suci) GetSchemeOutputA() SchemeOutputA {
	return p.schemeOutputA
}

func (p *Suci) SetMsinHex2SchemeOutput(hexMsin []byte) error {
	// todo: 支持可变长度的数据位数
	//if len(hexMsin) != 5 {
	//	return fmt.Errorf("invalid lenght for msin,expert 10")
	//}
	p.schemeOutput = p.schemeOutput[:0]
	for i, _ := range hexMsin {
		p.schemeOutput = append(p.schemeOutput, hexMsin[i]&lowerNibbleMask)
		p.schemeOutput = append(p.schemeOutput, hexMsin[i]>>upperNibbleShift)
	}
	// Remove the fill digit （0x0f）when the odd digit number
	if p.schemeOutput[len(p.schemeOutput)-1] == 0x0f {
		p.schemeOutput = p.schemeOutput[:len(p.schemeOutput)-1]
	}
	//p.schemeOutput = append(p.schemeOutput[:0], hexMsin[0]&lowerNibbleMask)
	//p.schemeOutput = append(p.schemeOutput, hexMsin[0]>>upperNibbleShift)
	//
	//p.schemeOutput = append(p.schemeOutput, hexMsin[1]&lowerNibbleMask)
	//p.schemeOutput = append(p.schemeOutput, hexMsin[1]>>upperNibbleShift)
	//
	//p.schemeOutput = append(p.schemeOutput, hexMsin[2]&lowerNibbleMask)
	//p.schemeOutput = append(p.schemeOutput, hexMsin[2]>>upperNibbleShift)
	//
	//p.schemeOutput = append(p.schemeOutput, hexMsin[3]&lowerNibbleMask)
	//p.schemeOutput = append(p.schemeOutput, hexMsin[3]>>upperNibbleShift)
	//
	//p.schemeOutput = append(p.schemeOutput, hexMsin[4]&lowerNibbleMask)
	//p.schemeOutput = append(p.schemeOutput, hexMsin[4]>>upperNibbleShift)

	return nil
}
func byteString(p []byte) string {
	var s string
	for i := 0; i < len(p); i++ {

		s += strconv.Itoa(int(p[i]))
	}
	return s
}
func (p *Suci) GetImsi() (*Imsi, error) {
	if p.supiType == IMSIType || p.protectSchemeID == NullScheme {

		homeNwId := p.GetHomeNetworkId()
		homeStr := byteString(homeNwId)
		outStr := byteString(p.GetSchemeOutput())
		imsiStr := homeStr + outStr
		imsi := &Imsi{}
		imsi.StoreImsiString(imsiStr, uint8(len(homeNwId[3:])))
		return imsi, nil
	}
	return nil, fmt.Errorf("cannot get Imsi from suci")
}

func (p *Suci) SetFromImsi(imsi *Imsi) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)

	p.reset()
	p.supiType = IMSIType
	p.protectSchemeID = NullScheme

	p.homeNetworkID = append(p.homeNetworkID, imsi.GetMccBytes()...)
	p.homeNetworkID = append(p.homeNetworkID, imsi.GetMncBytes()...)
	p.routeIndicator = append(p.routeIndicator, 0)
	p.homeNetworkPubKeyID = 0

	p.schemeOutput = append(p.schemeOutput, imsi.GetMsIn()...)
	return nil
}

func (p *Suci) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)

	if p.supiType != IMSIType {
		return nil, fmt.Errorf("not Imsi, unsupported yet")
	}

	// only imsi supported currently
	var buf []byte
	var octet byte

	//encode home network id, mcc+mnc
	length := len(p.homeNetworkID)
	if length == 5 || length == 6 {
		octet = p.homeNetworkID[0] | (p.homeNetworkID[1] << 4)
		buf = append(buf, octet)
		if len(p.homeNetworkID) == MinDigitNumofPlmn {
			octet = p.homeNetworkID[2] | 0xF0
			buf = append(buf, octet)
		} else {
			octet = p.homeNetworkID[2] & (p.homeNetworkID[5] << 4)
			buf = append(buf, octet)
		}
		octet = p.homeNetworkID[3] | (p.homeNetworkID[4] << 4)
		buf = append(buf, octet)
	} else {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "incorrect length for home Network id (%d)", length)
		return nil, fmt.Errorf("invalid lenght for home Newwork id")
	}
	//encode routing indicator
	length = len(p.routeIndicator)
	switch length {
	case 1:
		octet = p.routeIndicator[0] | 0xF0
		buf = append(buf, octet)
		buf = append(buf, 0xFF)
	case 2:
		octet = p.routeIndicator[0] | (p.routeIndicator[1] << 4)
		buf = append(buf, octet)
		buf = append(buf, 0xFF)
	case 3:
		octet = p.routeIndicator[0] | (p.routeIndicator[1] << 4)
		buf = append(buf, octet)
		octet = p.routeIndicator[2] | 0xF0
		buf = append(buf, octet)
	case 4:
		octet = p.routeIndicator[0] | (p.routeIndicator[1] << 4)
		buf = append(buf, octet)
		octet = p.routeIndicator[2] | (p.routeIndicator[3] << 4)
		buf = append(buf, octet)
	}
	//protect scheme id
	octet = p.protectSchemeID & 0x0F
	buf = append(buf, octet)

	//home network public key identifier
	buf = append(buf, p.homeNetworkPubKeyID)

	//schemeOutput
	buf = append(buf, p.schemeOutput[1]<<upperNibbleShift|p.schemeOutput[0])
	buf = append(buf, p.schemeOutput[3]<<upperNibbleShift|p.schemeOutput[2])
	buf = append(buf, p.schemeOutput[5]<<upperNibbleShift|p.schemeOutput[4])
	buf = append(buf, p.schemeOutput[7]<<upperNibbleShift|p.schemeOutput[6])
	buf = append(buf, p.schemeOutput[9]<<upperNibbleShift|p.schemeOutput[8])
	return buf, nil
}

func (p *Suci) Decode(msgBuf *bytes.Reader) error {
	// octet 2 : LV : no T , Start directly from the second octet
	lenSuci := make([]byte, 2)
	err := binary.Read(msgBuf, binary.BigEndian, &lenSuci)
	if err != nil {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "fail to read len of suci")
		return fmt.Errorf("fail to read len of suci")
	}
	len := binary.BigEndian.Uint16(lenSuci)

	// octet 3
	idTypeOctet, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read octet 3")
	}
	supiFormat, err := utils.GetBitsValue(idTypeOctet, 5, 8)
	if err != nil {
		return fmt.Errorf("failed to read supiFormat")
	}
	if supiFormat == 0 {
		p.supiType = IMSIType
	}
	// octet 4~6
	homeNetworkID := make([]byte, SizeofPlmnID)
	err = binary.Read(msgBuf, binary.BigEndian, &(homeNetworkID))
	if err != nil {
		return fmt.Errorf("failed to read homeNetworkID")
	}
	//mcc
	mcc1, _ := utils.GetBitsValue(homeNetworkID[0], 1, 4)
	p.homeNetworkID = append(p.homeNetworkID, mcc1)
	mcc2, _ := utils.GetBitsValue(homeNetworkID[0], 5, 8)
	p.homeNetworkID = append(p.homeNetworkID, mcc2>>4)
	mcc3, _ := utils.GetBitsValue(homeNetworkID[1], 1, 4)
	p.homeNetworkID = append(p.homeNetworkID, mcc3)
	//mnc
	mnc1, _ := utils.GetBitsValue(homeNetworkID[2], 1, 4)
	p.homeNetworkID = append(p.homeNetworkID, mnc1)
	mnc2, _ := utils.GetBitsValue(homeNetworkID[2], 5, 8)
	p.homeNetworkID = append(p.homeNetworkID, mnc2>>4)
	mnc3, _ := utils.GetBitsValue(homeNetworkID[1], 5, 8)
	if mnc3 != 0xf0 {
		p.homeNetworkID = append(p.homeNetworkID, mnc3>>4)
	}

	// octet 7
	p.routeIndicator = make([]byte, MaxLenofRouteInd)
	routeIndicatorOctet1, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read routeIndicator1")
	}
	p.routeIndicator[0], err = utils.GetBitsValue(routeIndicatorOctet1, 1, 4)
	p.routeIndicator[1], err = utils.GetBitsValue(routeIndicatorOctet1, 5, 8)
	// octet 8
	routeIndicatorOctet2, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read routeIndicator2")
	}
	p.routeIndicator[2], err = utils.GetBitsValue(routeIndicatorOctet2, 1, 4)
	p.routeIndicator[3], err = utils.GetBitsValue(routeIndicatorOctet2, 5, 8)
	// octet 9
	p.protectSchemeID, err = msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read protectSchemeID")
	}
	// octet 10
	p.homeNetworkPubKeyID, err = msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read homeNetworkPubKeyID")
	}
	// octet 11 ~ lenSuci
	schemeOutput := make([]byte, len-8)
	err = binary.Read(msgBuf, binary.BigEndian, &(schemeOutput))
	if err != nil {
		return fmt.Errorf("failed to read schemeOutput")
	}
	p.schemeOutput = append(p.schemeOutput, schemeOutput[0]&lowerNibbleMask)
	p.schemeOutput = append(p.schemeOutput, schemeOutput[0]>>upperNibbleShift)

	p.schemeOutput = append(p.schemeOutput, schemeOutput[1]&lowerNibbleMask)
	p.schemeOutput = append(p.schemeOutput, schemeOutput[1]>>upperNibbleShift)

	p.schemeOutput = append(p.schemeOutput, schemeOutput[2]&lowerNibbleMask)
	p.schemeOutput = append(p.schemeOutput, schemeOutput[2]>>upperNibbleShift)

	p.schemeOutput = append(p.schemeOutput, schemeOutput[3]&lowerNibbleMask)
	p.schemeOutput = append(p.schemeOutput, schemeOutput[3]>>upperNibbleShift)

	p.schemeOutput = append(p.schemeOutput, schemeOutput[4]&lowerNibbleMask)
	p.schemeOutput = append(p.schemeOutput, schemeOutput[4]>>upperNibbleShift)
	return nil
}

// profile A
// SUPI Type + Home Network Identifier +
//Routing Indicator +
//Protection Scheme Identifier +
//Home Network Public Key Identifier +
//Scheme Output
// Figure 9.11.3.4.3: 5GS mobile identity information element for type of identity "SUCI" and SUPI format "IMSI"
const MaxSchemeOutput int = 3000
const MinProASchemeOutputLen int = 40

func (p *Suci) DecodeConcealed(msgBuf *bytes.Reader) error {
	// octet 2 : LV : no T , Start directly from the second octet
	lenSuci := make([]byte, 2)
	err := binary.Read(msgBuf, binary.BigEndian, &lenSuci)
	if err != nil {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "fail to read len of suci")
		return fmt.Errorf("fail to read len of suci")
	}
	length := binary.BigEndian.Uint16(lenSuci)

	// octet 4
	// Type of identity (octet 4)
	idTypeOctet, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read octet 4")
	}
	//fmt.Printf("idTypeOctet:%08b\n", idTypeOctet)
	// SUPI format (octet 4, bits 5-7)
	supiFormat, err := utils.GetBitsValue(idTypeOctet, 5, 8)
	if err != nil {
		return fmt.Errorf("failed to read supiFormat")
	}
	//fmt.Printf("supiFormat:%08b\n", supiFormat)
	if supiFormat == 0 {
		p.supiType = IMSIType
	} else {
		return fmt.Errorf("Invalid supi format Type,expect imsi type")
	}

	// octet 5~7
	// MCC + MNC
	homeNetworkID := make([]byte, SizeofPlmnID)
	err = binary.Read(msgBuf, binary.BigEndian, &(homeNetworkID))
	if err != nil {
		return fmt.Errorf("failed to read homeNetworkID")
	}
	//mcc
	mcc1, _ := utils.GetBitsValue(homeNetworkID[0], 1, 4)
	p.homeNetworkID = append(p.homeNetworkID, mcc1)
	mcc2, _ := utils.GetBitsValue(homeNetworkID[0], 5, 8)
	p.homeNetworkID = append(p.homeNetworkID, mcc2>>4)
	mcc3, _ := utils.GetBitsValue(homeNetworkID[1], 1, 4)
	p.homeNetworkID = append(p.homeNetworkID, mcc3)
	//fmt.Printf("mcc:%d\n", p.homeNetworkID)

	//mnc
	mnc1, _ := utils.GetBitsValue(homeNetworkID[2], 1, 4)
	p.homeNetworkID = append(p.homeNetworkID, mnc1)
	mnc2, _ := utils.GetBitsValue(homeNetworkID[2], 5, 8)
	p.homeNetworkID = append(p.homeNetworkID, mnc2>>4)
	mnc3, _ := utils.GetBitsValue(homeNetworkID[1], 5, 8)
	// MNC 2 digits,bits 5 to 8 of octet 6 shall be coded as "1111".
	if mnc3 != 0xf0 {
		p.homeNetworkID = append(p.homeNetworkID, mnc3>>4)
	}
	//fmt.Printf("mnc:%d\n", p.homeNetworkID)

	// octet 8-9
	// octet 8
	p.routeIndicator = make([]byte, MaxLenofRouteInd)
	routeIndicatorOctet1, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read routeIndicator1")
	}
	p.routeIndicator[0], err = utils.GetBitsValue(routeIndicatorOctet1, 1, 4)
	p.routeIndicator[1], err = utils.GetBitsValue(routeIndicatorOctet1, 5, 8)
	// octet 9
	routeIndicatorOctet2, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read routeIndicator2")
	}
	p.routeIndicator[2], err = utils.GetBitsValue(routeIndicatorOctet2, 1, 4)
	p.routeIndicator[3], err = utils.GetBitsValue(routeIndicatorOctet2, 5, 8)
	//fmt.Printf("routeIndicator:%d\n", p.routeIndicator)

	// octet 10
	p.protectSchemeID, err = msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read protectSchemeID")
	}
	//fmt.Printf("protectSchemeID:%d\n", p.protectSchemeID)

	// octet 11 // todo：1 公钥ID 存储在本地
	p.homeNetworkPubKeyID, err = msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read homeNetworkPubKeyID")
	}
	//fmt.Printf("homeNetworkPubKeyID:%d\n", p.homeNetworkPubKeyID)

	// octet 12 ~ lenSuci
	outputLen := length - 8
	if int(outputLen) > MaxSchemeOutput {
		return fmt.Errorf("invalid lenght for scheme Output(%d)", outputLen)
	}
	schemeOutput := make([]byte, outputLen)
	//fmt.Printf("schemeOutput len:%d\n", len(schemeOutput))

	err = binary.Read(msgBuf, binary.BigEndian, &(schemeOutput))
	if err != nil {
		return fmt.Errorf("failed to read schemeOutput")
	}
	switch p.protectSchemeID {
	case NullScheme:
		{
			// IMSI
			/*p.schemeOutput = append(p.schemeOutput, schemeOutput[0]&lowerNibbleMask)
			p.schemeOutput = append(p.schemeOutput, schemeOutput[0]>>upperNibbleShift)

			p.schemeOutput = append(p.schemeOutput, schemeOutput[1]&lowerNibbleMask)
			p.schemeOutput = append(p.schemeOutput, schemeOutput[1]>>upperNibbleShift)

			p.schemeOutput = append(p.schemeOutput, schemeOutput[2]&lowerNibbleMask)
			p.schemeOutput = append(p.schemeOutput, schemeOutput[2]>>upperNibbleShift)

			p.schemeOutput = append(p.schemeOutput, schemeOutput[3]&lowerNibbleMask)
			p.schemeOutput = append(p.schemeOutput, schemeOutput[3]>>upperNibbleShift)

			p.schemeOutput = append(p.schemeOutput, schemeOutput[4]&lowerNibbleMask)
			p.schemeOutput = append(p.schemeOutput, schemeOutput[4]>>upperNibbleShift)*/
			p.schemeOutput = p.schemeOutput[:0]
			for i, _ := range schemeOutput {
				p.schemeOutput = append(p.schemeOutput, schemeOutput[i]&lowerNibbleMask)
				p.schemeOutput = append(p.schemeOutput, schemeOutput[i]>>upperNibbleShift)
			}
			// Remove the fill digit （0x0f）when the odd digit number
			if p.schemeOutput[len(p.schemeOutput)-1] == 0x0f {
				p.schemeOutput = p.schemeOutput[:len(p.schemeOutput)-1]
			}
		}
	case ProfileA:
		{
			// 2 Ciphertext 的长度是可变的
			// TS 23.003 Figure 2.2B-3: Scheme Output for Elliptic Curve Integrated Encryption Scheme Profile A
			// The ECC ephemeral public key is formatted as 64 hexadecimal digits, which allows to encode 256 bits.
			// The ciphertext value is formatted as a variable length of hexadecimal digits.
			// The MAC tag value is formatted as 16 hexadecimal digits, which allows to encode 64 bits.

			if int(outputLen) < MinProASchemeOutputLen {
				return fmt.Errorf("invalid lenght for Scheme Profile A Output(%d)", outputLen)
			}
			p.schemeOutputA.EphPublicKey = schemeOutput[:32]
			//CiphertextLen := outputLen - uint16(MinProASchemeOutputLen)
			p.schemeOutputA.Ciphertext = schemeOutput[32 : outputLen-8]
			p.schemeOutputA.MacTag = schemeOutput[outputLen-8 : outputLen]
			//fmt.Printf("schemeOutput EphPublicKey:%d\n", len(p.schemeOutputA.EphPublicKey))
			//fmt.Printf("schemeOutput Ciphertext:%d\n", len(p.schemeOutputA.Ciphertext))
			//fmt.Printf("schemeOutput MacTag:%d\n", len(p.schemeOutputA.MacTag))

		}
	default:
		return fmt.Errorf("invalid lenght for protect scheme ID")
	}

	return nil
}
