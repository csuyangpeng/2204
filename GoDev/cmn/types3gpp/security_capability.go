package types3gpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"strings"
)

type SecurityCapability struct {
	nrEncryptAlgo    bitset.BitSet
	nrIntPrctAlgo    bitset.BitSet
	eutraEncryptAlgo bitset.BitSet
	eutraIntPrctAlgo bitset.BitSet
}

func (p *SecurityCapability) Reset() {
	p.nrEncryptAlgo.ClearAll()
	p.nrIntPrctAlgo.ClearAll()
	p.eutraEncryptAlgo.ClearAll()
	p.eutraIntPrctAlgo.ClearAll()
}

func (p *SecurityCapability) NEAReset() {
	p.nrEncryptAlgo.ClearAll()
}

func (p *SecurityCapability) NIAReset() {
	p.nrIntPrctAlgo.ClearAll()
}

func (p *SecurityCapability) IsMatched(secCap SecurityCapability) bool {
	if byteString(p.Encode()) == byteString(secCap.Encode()) {
		return true
	}
	rlogger.Trace(types.ModuleCmn3gtp, rlogger.DEBUG, nil,
		"Security Capability Mismatched.(%s) - (%s)", p.String(), secCap.String())
	return false
}

type NREncryptAlgo byte

const (
	NEA0 NREncryptAlgo = iota
	NEA1
	NEA2
	NEA3
	NEA4
	NEA5
	NEA6
	NEA7
)

func (p NREncryptAlgo) String() string {
	var rt string
	switch p {
	case NEA0:
		rt = "NEA0"
	case NEA1:
		rt = "NEA1"
	case NEA2:
		rt = "NEA2"
	case NEA3:
		rt = "NEA3"
	default:
	}
	return rt
}

type NRIntPrctAlgo byte

const (
	NIA0 NRIntPrctAlgo = iota
	NIA1
	NIA2
	NIA3
	NIA4
	NIA5
	NIA6
	NIA7
)

func (p NRIntPrctAlgo) String() string {
	var rt string
	switch p {
	case NIA0:
		rt = "NIA0"
	case NIA1:
		rt = "NIA1"
	case NIA2:
		rt = "NIA2"
	case NIA3:
		rt = "NIA3"
	default:
		rt = "unknown"
	}
	return rt
}

type EutraEncryptAlgo byte

const (
	EEA0 EutraEncryptAlgo = iota
	EEA1
	EEA2
	EEA3
	EEA4
	EEA5
	EEA6
	EEA7
)

type EutraIntPrctAlgo byte

const (
	EIA0 EutraIntPrctAlgo = iota
	EIA1
	EIA2
	EIA3
	EIA4
	EIA5
	EIA6
	EIA7
)
const (
	left_shift_bit_7 = 7
	left_shift_bit_6 = 6
	left_shift_bit_5 = 5
	left_shift_bit_4 = 4
	left_shift_bit_3 = 3
	left_shift_bit_2 = 2
	left_shift_bit_1 = 1

	left_shift_bit_15 = 15
	left_shift_bit_14 = 14
	left_shift_bit_13 = 13
	left_shift_bit_12 = 12
)

func (p *SecurityCapability) SetNrEncAlgo(ind NREncryptAlgo) {
	p.nrEncryptAlgo.Set(uint(ind))
}
func (p *SecurityCapability) SetNrEncAlgoStr(algo string) {
	switch algo {
	case "NEA0":
		p.SetNrEncAlgo(NEA0)
	case "NEA1":
		p.SetNrEncAlgo(NEA1)
	case "NEA2":
		p.SetNrEncAlgo(NEA2)
	case "NEA3":
		p.SetNrEncAlgo(NEA3)
	default:
	}
}

func (p *SecurityCapability) GetNrEncAlgo() bitset.BitSet {
	return p.nrEncryptAlgo
}
func (p *SecurityCapability) GetNrEncAlgoU16() []byte {
	var rt uint16
	if p.nrEncryptAlgo.Test(uint(NEA1)) {
		rt |= 1 << left_shift_bit_15
	}
	if p.nrEncryptAlgo.Test(uint(NEA2)) {
		rt |= 1 << left_shift_bit_14
	}
	if p.nrEncryptAlgo.Test(uint(NEA3)) {
		rt |= 1 << left_shift_bit_13
	}
	// 3gpp 38.413  9.3.1.86
	// "all bits equal to 0" ¨C UE supports no other algorithm than NEA0,
	if p.nrEncryptAlgo.Test(uint(NEA0)) {
		//rt |= 0
	}
	val := make([]byte, 2)
	binary.BigEndian.PutUint16(val, rt)
	return val
}

func (p *SecurityCapability) SetNrEncAlgoU16(val []byte) error {
	if val == nil {
		return fmt.Errorf("invalid input parameter(nil)")
	}
	input := binary.BigEndian.Uint16(val)
	if (input & (1 << left_shift_bit_15)) != 0 {
		p.SetNrEncAlgo(NEA0)
	}
	if (input & (1 << left_shift_bit_14)) != 0 {
		p.SetNrEncAlgo(NEA1)
	}
	if (input & (1 << left_shift_bit_13)) != 0 {
		p.SetNrEncAlgo(NEA2)
	}
	if (input & (1 << left_shift_bit_12)) != 0 {
		p.SetNrEncAlgo(NEA3)
	}
	return nil
}

func (p *SecurityCapability) GetNrEncAlgoU8() byte {
	var val byte
	if p.nrEncryptAlgo.Test(uint(NEA0)) {
		val |= 1 << left_shift_bit_7
	}
	if p.nrEncryptAlgo.Test(uint(NEA1)) {
		val |= 1 << left_shift_bit_6
	}
	if p.nrEncryptAlgo.Test(uint(NEA2)) {
		val |= 1 << left_shift_bit_5
	}
	if p.nrEncryptAlgo.Test(uint(NEA3)) {
		val |= 1 << left_shift_bit_4
	}
	if p.nrEncryptAlgo.Test(uint(NEA4)) {
		val |= 1 << left_shift_bit_3
	}
	if p.nrEncryptAlgo.Test(uint(NEA5)) {
		val |= 1 << left_shift_bit_2
	}
	if p.nrEncryptAlgo.Test(uint(NEA6)) {
		val |= 1 << left_shift_bit_1
	}
	if p.nrEncryptAlgo.Test(uint(NEA7)) {
		val |= 0x01
	}
	return val
}

func (p *SecurityCapability) SetNrEncAlgoU8(val byte) {
	if (val & (1 << left_shift_bit_7)) != 0 {
		p.SetNrEncAlgo(NEA0)
	}
	if (val & (1 << left_shift_bit_6)) != 0 {
		p.SetNrEncAlgo(NEA1)
	}
	if (val & (1 << left_shift_bit_5)) != 0 {
		p.SetNrEncAlgo(NEA2)
	}
	if (val & (1 << left_shift_bit_4)) != 0 {
		p.SetNrEncAlgo(NEA3)
	}
	if (val & (1 << left_shift_bit_3)) != 0 {
		p.SetNrEncAlgo(NEA4)
	}
	if (val & (1 << left_shift_bit_2)) != 0 {
		p.SetNrEncAlgo(NEA5)
	}
	if (val & (1 << left_shift_bit_1)) != 0 {
		p.SetNrEncAlgo(NEA6)
	}
	if (val & 0x01) != 0 {
		p.SetNrEncAlgo(NEA7)
	}
}
func (p *SecurityCapability) SetNrIntPrctAlgo(ind NRIntPrctAlgo) {
	p.nrIntPrctAlgo.Set(uint(ind))
}
func (p *SecurityCapability) SetNrIntAlgoStr(algo string) {
	switch algo {
	case "NIA0":
		p.SetNrIntPrctAlgo(NIA0)
	case "NIA1":
		p.SetNrIntPrctAlgo(NIA1)
	case "NIA2":
		p.SetNrIntPrctAlgo(NIA2)
	case "NIA3":
		p.SetNrIntPrctAlgo(NIA3)
	default:
	}
}
func (p *SecurityCapability) GetNrIntPrctAlgo() bitset.BitSet {
	return p.nrIntPrctAlgo
}
func (p *SecurityCapability) GetNrIntPrctAlgoU16() []byte {
	var rt uint16

	if p.nrIntPrctAlgo.Test(uint(NIA1)) {
		rt |= 1 << left_shift_bit_15
	}
	if p.nrIntPrctAlgo.Test(uint(NIA2)) {
		rt |= 1 << left_shift_bit_14
	}
	if p.nrIntPrctAlgo.Test(uint(NIA3)) {
		rt |= 1 << left_shift_bit_13
	}

	// 3gpp 38.413  9.3.1.86
	// "all bits equal to 0" ¨C UE supports no other algorithm than NIA0,
	if p.nrIntPrctAlgo.Test(uint(NIA0)) {
		//rt = 0
	}

	val := make([]byte, 2)
	binary.BigEndian.PutUint16(val, rt)
	return val
}
func (p *SecurityCapability) SetNrIntPrctAlgoU16(val []byte) error {
	if val == nil {
		return fmt.Errorf("invalid input parameter(nil)")
	}
	input := binary.BigEndian.Uint16(val)
	if (input & (1 << left_shift_bit_15)) != 0 {
		p.SetNrIntPrctAlgo(NIA0)
	}
	if (input & (1 << left_shift_bit_14)) != 0 {
		p.SetNrIntPrctAlgo(NIA1)
	}
	if (input & (1 << left_shift_bit_13)) != 0 {
		p.SetNrIntPrctAlgo(NIA2)
	}
	if (input & (1 << left_shift_bit_12)) != 0 {
		p.SetNrIntPrctAlgo(NIA3)
	}
	return nil
}

func (p *SecurityCapability) GetNrIntPrctAlgoU8() byte {
	var val byte
	if p.nrIntPrctAlgo.Test(uint(NIA0)) {
		val |= 1 << left_shift_bit_7
	}
	if p.nrIntPrctAlgo.Test(uint(NIA1)) {
		val |= 1 << left_shift_bit_6
	}
	if p.nrIntPrctAlgo.Test(uint(NIA2)) {
		val |= 1 << left_shift_bit_5
	}
	if p.nrIntPrctAlgo.Test(uint(NIA3)) {
		val |= 1 << left_shift_bit_4
	}
	if p.nrIntPrctAlgo.Test(uint(NIA4)) {
		val |= 1 << left_shift_bit_3
	}
	if p.nrIntPrctAlgo.Test(uint(NIA5)) {
		val |= 1 << left_shift_bit_2
	}
	if p.nrIntPrctAlgo.Test(uint(NIA6)) {
		val |= 1 << left_shift_bit_1
	}
	if p.nrIntPrctAlgo.Test(uint(NIA7)) {
		val |= 0x01
	}
	return val
}

func (p *SecurityCapability) SetNrIntPrctAlgoU8(val byte) {
	if (val & (1 << left_shift_bit_7)) != 0 {
		p.SetNrIntPrctAlgo(NIA0)
	}
	if (val & (1 << left_shift_bit_6)) != 0 {
		p.SetNrIntPrctAlgo(NIA1)
	}
	if (val & (1 << left_shift_bit_5)) != 0 {
		p.SetNrIntPrctAlgo(NIA2)
	}
	if (val & (1 << left_shift_bit_4)) != 0 {
		p.SetNrIntPrctAlgo(NIA3)
	}
	if (val & (1 << left_shift_bit_3)) != 0 {
		p.SetNrIntPrctAlgo(NIA4)
	}
	if (val & (1 << left_shift_bit_2)) != 0 {
		p.SetNrIntPrctAlgo(NIA5)
	}
	if (val & (1 << left_shift_bit_1)) != 0 {
		p.SetNrIntPrctAlgo(NIA6)
	}
	if (val & 0x01) != 0 {
		p.SetNrIntPrctAlgo(NIA7)
	}
}

func (p *SecurityCapability) SetEutraEncAlgo(ind EutraEncryptAlgo) {
	p.eutraEncryptAlgo.Set(uint(ind))
}
func (p *SecurityCapability) GetEutraEncAlgo() bitset.BitSet {
	return p.eutraEncryptAlgo
}
func (p *SecurityCapability) GetEutraEncAlgoU16() []byte {
	var rt uint16

	if p.eutraEncryptAlgo.Test(uint(EEA1)) {
		rt |= 1 << left_shift_bit_14
	}
	if p.eutraEncryptAlgo.Test(uint(EEA2)) {
		rt |= 1 << left_shift_bit_13
	}
	if p.eutraEncryptAlgo.Test(uint(EEA3)) {
		rt |= 1 << left_shift_bit_12
	}

	// 3gpp 38.413  9.3.1.86
	// "all bits equal to 0" ¨C UE supports no other algorithm than EEA0,
	if p.eutraEncryptAlgo.Test(uint(EEA0)) {
		rt = 0
	}

	val := make([]byte, 2)
	binary.BigEndian.PutUint16(val, rt)
	return val
}
func (p *SecurityCapability) SetEutraEncAlgoU16(val []byte) error {
	if val == nil {
		return fmt.Errorf("invalid input parameter(nil)")
	}
	input := binary.BigEndian.Uint16(val)
	if (input & (1 << left_shift_bit_15)) != 0 {
		p.SetEutraEncAlgo(EEA0)
	}
	if (input & (1 << left_shift_bit_14)) != 0 {
		p.SetEutraEncAlgo(EEA1)
	}
	if (input & (1 << left_shift_bit_13)) != 0 {
		p.SetEutraEncAlgo(EEA2)
	}
	if (input & (1 << left_shift_bit_12)) != 0 {
		p.SetEutraEncAlgo(EEA3)
	}
	return nil
}

func (p *SecurityCapability) GetEutraEncAlgoU8() byte {
	var val byte
	if p.eutraEncryptAlgo.Test(uint(EEA0)) {
		val |= 1 << left_shift_bit_7
	}
	if p.eutraEncryptAlgo.Test(uint(EEA1)) {
		val |= 1 << left_shift_bit_6
	}
	if p.eutraEncryptAlgo.Test(uint(EEA2)) {
		val |= 1 << left_shift_bit_5
	}
	if p.eutraEncryptAlgo.Test(uint(EEA3)) {
		val |= 1 << left_shift_bit_4
	}
	if p.eutraEncryptAlgo.Test(uint(EEA4)) {
		val |= 1 << left_shift_bit_3
	}
	if p.eutraEncryptAlgo.Test(uint(EEA5)) {
		val |= 1 << left_shift_bit_2
	}
	if p.eutraEncryptAlgo.Test(uint(EEA6)) {
		val |= 1 << left_shift_bit_1
	}
	if p.eutraEncryptAlgo.Test(uint(EEA7)) {
		val |= 0x01
	}
	return val
}

func (p *SecurityCapability) SetEutraEncAlgoU8(val byte) {
	if (val & (1 << left_shift_bit_7)) != 0 {
		p.SetEutraEncAlgo(EEA0)
	}
	if (val & (1 << left_shift_bit_6)) != 0 {
		p.SetEutraEncAlgo(EEA1)
	}
	if (val & (1 << left_shift_bit_5)) != 0 {
		p.SetEutraEncAlgo(EEA2)
	}
	if (val & (1 << left_shift_bit_4)) != 0 {
		p.SetEutraEncAlgo(EEA3)
	}
	if (val & (1 << left_shift_bit_3)) != 0 {
		p.SetEutraEncAlgo(EEA4)
	}
	if (val & (1 << left_shift_bit_2)) != 0 {
		p.SetEutraEncAlgo(EEA5)
	}
	if (val & (1 << left_shift_bit_1)) != 0 {
		p.SetEutraEncAlgo(EEA6)
	}
	if (val & 0x01) != 0 {
		p.SetEutraEncAlgo(EEA7)
	}
}

func (p *SecurityCapability) SetEutraIntPrctAlgo(ind EutraIntPrctAlgo) {
	p.eutraIntPrctAlgo.Set(uint(ind))
}
func (p *SecurityCapability) GetEutraIntPrctAlgo() bitset.BitSet {
	return p.eutraIntPrctAlgo
}
func (p *SecurityCapability) GetEutraIntPrctAlgoU16() []byte {
	var rt uint16

	if p.eutraIntPrctAlgo.Test(uint(EIA1)) {
		rt |= 1 << left_shift_bit_14
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA2)) {
		rt |= 1 << left_shift_bit_13
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA3)) {
		rt |= 1 << left_shift_bit_12
	}

	// 3gpp 38.413  9.3.1.86
	// "all bits equal to 0" ¨C UE supports no other algorithm than EIA0
	if p.eutraIntPrctAlgo.Test(uint(EIA0)) {
		rt = 0
	}
	val := make([]byte, 2)
	binary.BigEndian.PutUint16(val, rt)
	return val
}
func (p *SecurityCapability) SetEutraIntPrctAlgoU16(val []byte) error {
	if val == nil {
		return fmt.Errorf("invalid input parameter(nil)")
	}
	input := binary.BigEndian.Uint16(val)
	if (input & (1 << left_shift_bit_15)) != 0 {
		p.SetEutraIntPrctAlgo(EIA0)
	}
	if (input & (1 << left_shift_bit_14)) != 0 {
		p.SetEutraIntPrctAlgo(EIA1)
	}
	if (input & (1 << left_shift_bit_13)) != 0 {
		p.SetEutraIntPrctAlgo(EIA2)
	}
	if (input & (1 << left_shift_bit_12)) != 0 {
		p.SetEutraIntPrctAlgo(EIA3)
	}
	return nil
}
func (p *SecurityCapability) GetEutraIntPrctAlgoU8() byte {
	var val byte
	if p.eutraIntPrctAlgo.Test(uint(EIA0)) {
		val |= 1 << left_shift_bit_7
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA1)) {
		val |= 1 << left_shift_bit_6
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA2)) {
		val |= 1 << left_shift_bit_5
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA3)) {
		val |= 1 << left_shift_bit_4
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA4)) {
		val |= 1 << left_shift_bit_3
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA5)) {
		val |= 1 << left_shift_bit_2
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA6)) {
		val |= 1 << left_shift_bit_1
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA7)) {
		val |= 0x01
	}
	return val
}

func (p *SecurityCapability) SetEutraIntPrctAlgoU8(val byte) {
	if (val & (1 << left_shift_bit_7)) != 0 {
		p.SetEutraIntPrctAlgo(EIA0)
	}
	if (val & (1 << left_shift_bit_6)) != 0 {
		p.SetEutraIntPrctAlgo(EIA1)
	}
	if (val & (1 << left_shift_bit_5)) != 0 {
		p.SetEutraIntPrctAlgo(EIA2)
	}
	if (val & (1 << left_shift_bit_4)) != 0 {
		p.SetEutraIntPrctAlgo(EIA3)
	}
	if (val & (1 << left_shift_bit_3)) != 0 {
		p.SetEutraIntPrctAlgo(EIA4)
	}
	if (val & (1 << left_shift_bit_2)) != 0 {
		p.SetEutraIntPrctAlgo(EIA5)
	}
	if (val & (1 << left_shift_bit_1)) != 0 {
		p.SetEutraIntPrctAlgo(EIA6)
	}
	if (val & 0x01) != 0 {
		p.SetEutraIntPrctAlgo(EIA7)
	}
}

func (p SecurityCapability) String() string {
	var rt string
	rt = "NR EncryptAlgo("
	if p.nrEncryptAlgo.Test(uint(NEA0)) {
		rt += "NEA0,"
	}
	if p.nrEncryptAlgo.Test(uint(NEA1)) {
		rt += "NEA1,"
	}
	if p.nrEncryptAlgo.Test(uint(NEA2)) {
		rt += "NEA2,"
	}
	if p.nrEncryptAlgo.Test(uint(NEA3)) {
		rt += "NEA3,"
	}

	rt += ") NR Integrity Prctect Algo("
	if p.nrIntPrctAlgo.Test(uint(NIA0)) {
		rt += "NIA0,"
	}
	if p.nrIntPrctAlgo.Test(uint(NIA1)) {
		rt += "NIA1,"
	}
	if p.nrIntPrctAlgo.Test(uint(NIA2)) {
		rt += "NIA2,"
	}
	if p.nrIntPrctAlgo.Test(uint(NIA3)) {
		rt += "NIA3,"
	}

	rt += ") Eutra Encrypt Alog("
	if p.eutraEncryptAlgo.Test(uint(EEA0)) {
		rt += "EEA0,"
	}
	if p.eutraEncryptAlgo.Test(uint(EEA1)) {
		rt += "EEA1,"
	}
	if p.eutraEncryptAlgo.Test(uint(EEA2)) {
		rt += "EEA2,"
	}
	if p.eutraEncryptAlgo.Test(uint(EEA3)) {
		rt += "EEA3,"
	}

	rt += ") Eutra Integrity Protect Algo("
	if p.eutraIntPrctAlgo.Test(uint(EIA0)) {
		rt += "EIA0,"
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA1)) {
		rt += "EIA1,"
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA2)) {
		rt += "EIA2,"
	}
	if p.eutraIntPrctAlgo.Test(uint(EIA3)) {
		rt += "EIA3,"
	}
	rt += ")"
	return rt
}

// 24.501 9.11.3.54 Decode for UeSecCapability
func (p *SecurityCapability) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	length, _ := msgBuf.ReadByte()
	if length < 2 || length > 8 {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "invalid length for "+
			"Ie ue security capability (%d)", length)
		return fmt.Errorf("failed to decode message")
	}

	//NrEnc
	octet, _ := msgBuf.ReadByte()
	p.SetNrEncAlgoU8(octet)

	//NrIntPrct
	octet, _ = msgBuf.ReadByte()
	p.SetNrIntPrctAlgoU8(octet)

	//4G Enc and IntPrct
	// Octets 5 to 10 are optional. If octet 5 is included, then also octet 6 shall be included.
	// If the UE does not support any security algorithm for AS security over E-UTRA connected to 5GCN,
	// it shall not include octets 5 and 6. The UE shall not include octets 7 to 10.
	if length > 2 {
		octet, _ = msgBuf.ReadByte()
		p.SetEutraEncAlgoU8(octet)
		octet, _ = msgBuf.ReadByte()
		p.SetEutraIntPrctAlgoU8(octet)
	}

	//If the network includes octet 7, then it shall include also octet 8.
	// If the network includes octet 9, then it shall include also octet 10.
	if length > 4 {
		leftBytes := make([]byte, length-4)
		binary.Read(msgBuf, binary.BigEndian, leftBytes)
	}

	return nil
}

// 24.501 9.11.3.54 Encode for UeSecCapability
func (p *SecurityCapability) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)

	// encode LV
	var buf []byte
	eeax := p.GetEutraEncAlgoU8()
	eiax := p.GetEutraIntPrctAlgoU8()
	neax := p.GetNrEncAlgoU8()
	niax := p.GetNrIntPrctAlgoU8()

	length := 2

	// Octets 5 to 10 are optional. If octet 5 is included, then also octet 6 shall be included.
	is4gAlgoExist := false
	if eeax != 0 || eiax != 0 {
		length = length + 2
		is4gAlgoExist = true
	}

	buf = append(buf, byte(length))
	buf = append(buf, neax)
	buf = append(buf, niax)
	// 	if 4g exist
	if is4gAlgoExist {
		buf = append(buf, eeax)
		buf = append(buf, eiax)
	}

	return buf
}

// 24.501 9.11.3.54 Decode for UeSecCapability
//func (p *SecurityCapability) Encode() []byte {
//	rlogger.FuncEntry(types.ModCmn,nil)
//
//	// encode LV
//
//	var buf []byte
//	eeax := p.GetEutraEncAlgoU8()
//	eiax := p.GetEutraIntPrctAlgoU8()
//	neax := p.GetNrEncAlgoU8()
//	niax := p.GetNrIntPrctAlgoU8()
//
//	length := 2
//	if eeax != 0 {
//		length ++
//	}
//	if eiax != 0 {
//		length ++
//	}
//
//	buf = append(buf, byte(length))
//	buf = append(buf, neax)
//	buf = append(buf, niax)
//
//	if length >= 3 {
//		buf = append(buf, eeax)
//	}
//	if length >= 4 {
//		buf = append(buf, eiax)
//	}
//
//	return buf
//}

func (p *SecurityCapability) StoreNrIntAlgo(intAlgoStr string) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	algos := strings.Split(intAlgoStr, ",")
	for _, v := range algos {
		p.SetNrIntAlgoStr(v)
	}

	return nil
}

func (p *SecurityCapability) StoreNrEncAlgo(encAlgoStr string) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	algos := strings.Split(encAlgoStr, ",")
	for _, v := range algos {
		p.SetNrEncAlgoStr(v)
	}
	return nil
}

func (p *SecurityCapability) MatchNrIntAlgo(ueNrIntAlgos bitset.BitSet) NRIntPrctAlgo {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)

	rt := p.nrIntPrctAlgo.Intersection(&ueNrIntAlgos)
	if rt.Test(uint(NIA3)) {
		return NIA3
	}
	if rt.Test(uint(NIA2)) {
		return NIA2
	}
	if rt.Test(uint(NIA1)) {
		return NIA1
	}
	if rt.Test(uint(NIA0)) {
		return NIA0
	}
	return NIA0
}

func (p *SecurityCapability) MatchNrEncAlgo(ueNrEncAlgos bitset.BitSet) NREncryptAlgo {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)

	rt := p.nrEncryptAlgo.Intersection(&ueNrEncAlgos)
	if rt.Test(uint(NEA3)) {
		return NEA3
	}
	if rt.Test(uint(NEA2)) {
		return NEA2
	}
	if rt.Test(uint(NEA1)) {
		return NEA1
	}
	if rt.Test(uint(NEA0)) {
		return NEA0
	}
	return NEA0
}
func (p *SecurityCapability) HaveCommonAlgo(ueNrIntAlgos bitset.BitSet, ueNrEncAlgos bitset.BitSet) bool {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	rt1 := p.nrIntPrctAlgo.Intersection(&ueNrIntAlgos)
	rt2 := p.nrEncryptAlgo.Intersection(&ueNrEncAlgos)
	if rt1.Count() == 0 || rt2.Count() == 0 {
		return false
	}
	return true
}
