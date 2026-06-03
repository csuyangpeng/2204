package types3gpp

import (
	"fmt"
)

// const definitions
const (
	BigEndian    = true
	LittleEndian = false
	SizeofPlmnID = 3
)

// PlmnID : Public Land Mobile Network
type PlmnID struct {
	value     [SizeofPlmnID]byte
	ByteOrder bool
}

func (p *PlmnID) Reset() {
	p.value = [SizeofPlmnID]byte{}
	p.ByteOrder = false
}

// SetValue func set the plmn with []byte, decode procedure
func (p *PlmnID) SetValue(val []byte, bo bool) error {
	if len(val) != SizeofPlmnID {
		return fmt.Errorf("invalid plmnid length(%d)", len(val))
	}

	copy(p.value[:], val)
	p.ByteOrder = bo

	return nil
}

// GetValue func return plmn with []byte, encode procedure
func (p *PlmnID) GetValue(bo bool) [SizeofPlmnID]byte {
	if p.ByteOrder == bo {
		return p.value
	}

	var plmn [SizeofPlmnID]byte
	//stored as BigEndian and retrive as LittleEndian
	for i := 0; i < SizeofPlmnID; i++ {
		plmn[i] = (p.value[i] << 4) ^ (p.value[i] >> 4)
		//High four bit and low four bit interchange
	}
	return plmn
}

// String func implement to LittleEndian string
// 24.501 Figure 9.11.3.49.2
// --------------------------
// big endian
// MCC2  MCC1      octet 1
// MNC3  MCC3      octet 2
// MNC2  MNC1      octet 3
// --------------------------
// little endian
// MCC1  MCC2      octet 1
// MCC3  MNC3      octet 2
// MNC1  MNC2      octet 3
// --------------------------
func (p PlmnID) GetMcc() []byte {
	var mcc []byte
	if p.ByteOrder == LittleEndian {
		mcc = append(mcc, p.value[0]>>4)
		mcc = append(mcc, p.value[0]&0x0F)
		mcc = append(mcc, p.value[1]>>4)
	} else {
		mcc = append(mcc, p.value[0]&0x0F)
		mcc = append(mcc, p.value[0]>>4)
		mcc = append(mcc, p.value[1]&0x0F)
	}

	return mcc
}
func (p PlmnID) GetMccString() string {
	mcc := p.GetMcc()
	var rtStr string
	for _, v := range mcc {
		rtStr += fmt.Sprintf("%d", v)
	}
	return rtStr
}

func (p PlmnID) GetMnc() []byte {
	var mnc []byte
	if p.ByteOrder == LittleEndian {
		mnc = append(mnc, p.value[2]>>4)
		mnc = append(mnc, p.value[2]&0x0F)
		mnc3 := p.value[1] & 0x0F
		if mnc3 != 0x0F {
			mnc = append(mnc, p.value[1]&0x0F)
		}
	} else {
		mnc = append(mnc, p.value[2]&0x0F)
		mnc = append(mnc, p.value[2]>>4)
		mnc3 := p.value[1] >> 4
		if mnc3 != 0x0F {
			mnc = append(mnc, p.value[1]>>4)
		}
	}
	return mnc
}
func (p PlmnID) GetMncString() string {
	mnc := p.GetMnc()
	var rtStr string
	for _, v := range mnc {
		rtStr += fmt.Sprintf("%d", v)
	}
	return rtStr
}
func (p PlmnID) String() string {
	return fmt.Sprintf("%s%s", p.GetMccString(), p.GetMncString())

	//var plmn []byte
	//if p.ByteOrder == LittleEndian {
	//	for _, v := range p.Value {
	//		b := (v & 0xF0) >> 4
	//		if b != 0x0F {
	//			plmn = append(plmn, b+'0')
	//		}
	//
	//		b = v & 0x0F
	//		if b != 0x0F {
	//			plmn = append(plmn, b+'0')
	//		}
	//	}
	//} else {
	//	for _, v := range p.Value {
	//		b := v & 0x0F
	//		if b != 0x0F {
	//			plmn = append(plmn, b+'0')
	//		}
	//
	//		b = (v & 0xF0) >> 4
	//		if b != 0x0F {
	//			plmn = append(plmn, b+'0')
	//		}
	//	}
	//}
	//
	//return string(plmn)
}

// --------------------------
// little endian
// MCC1  MCC2      octet 1
// MCC3  MNC3      octet 2
// MNC1  MNC2      octet 3
// --------------------------
// SetString func set the plmn with string
func (p *PlmnID) SetString(str string) error {

	strLen := len(str)

	if (strLen != 5) && (strLen != 6) {
		return fmt.Errorf("invalid input parameter plmn string: %s", str)
	}

	p.ByteOrder = LittleEndian

	var plmnbytes []byte
	for _, v := range str {
		plmnbytes = append(plmnbytes, byte(v)-'0')
	}

	p.value[0] = plmnbytes[0]<<4 | plmnbytes[1]&0x0F
	if strLen == 5 {
		p.value[1] = plmnbytes[2]<<4 | 0x0F
	} else {
		p.value[1] = plmnbytes[2]<<4 | plmnbytes[5]&0x0F
	}
	p.value[2] = plmnbytes[3]<<4 | plmnbytes[4]&0x0F

	//for i := 0; i < strLen; i++ {
	//	if i == 3 && strLen == 5 {
	//		plmnbytes = append(plmnbytes, 0x0f)
	//	}
	//	plmnbytes = append(plmnbytes, str[i]-'0')
	//}
	//
	//for i := 0; i < SizeofPlmnID; i++ {
	//	p.Value[i] = (plmnbytes[2*i] << 4) | (plmnbytes[2*i+1])
	//}

	return nil
}
