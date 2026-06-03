package types3gpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"regexp"
	"strings"
)

const (
	MaxLengthofApnNi = 63
	MaxLengthofApn   = 100
)

type ApnNI struct {
	length uint8
	value  []uint8
}

func (p *ApnNI) SetValue(apnNi []byte) error {
	length := len(apnNi)
	if length > MaxLengthofApnNi {
		return fmt.Errorf("lengh of ApnNi(%d) exceed"+
			" max lenghth allowed(%d)", length, MaxLengthofApnNi)
	}

	if !IsApnNiValid(string(apnNi)) {
		return fmt.Errorf("format of ApnNi is invalid")
	}

	p.value = p.value[0:0]
	p.value = append(p.value, apnNi...)
	//copy(p.value[:], apnNi)
	p.length = uint8(length)

	return nil
}

func (p *ApnNI) GetValue() (out []uint8) {
	copy(out, p.value[:p.length])
	return
}

func (p ApnNI) String() string {
	return fmt.Sprintf("%s",
		string(p.value[:p.length]))
}

func IsApnNiValid(str string) bool {
	//TS 23.003 9.1.1
	//An APN Network Identifier shall not start with any of the strings
	// "rac", "lac", "sgsn" or "rnc", and it shall not end in ".gprs",
	// i.e. the last label of the APN Network Identifier shall not be "gprs".
	// Further, it shall not take the value "*".
	if strings.HasPrefix(str, "rac") ||
		strings.HasPrefix(str, "lac") ||
		strings.HasPrefix(str, "sgsn") ||
		strings.HasPrefix(str, "rnc") ||
		strings.HasSuffix(str, ".gprs") {
		return false
	}

	// "*" is the The wild card APN Ts 23.003 2.9.2
	if str != "*" &&
		strings.Contains(str, "*") {
		return false
	}

	return true
}

type ApnOI struct {
	length uint8
	value  []byte
}

func (p *ApnOI) GetValue() (out []byte) {
	copy(out, p.value[:p.length])
	return
}

func (p *ApnOI) SetValue(apnOi []byte) error {
	//TODO validation for the apn oi here

	p.length = uint8(len(apnOi))
	//copy(p.value[:], apnOi)
	p.value = p.value[0:0]
	p.value = append(p.value, apnOi...)
	return nil
}

func (p *ApnOI) DeriveFromImsi(imsi *Imsi) {

	oiBytes := []byte(fmt.Sprintf("mnc%03d.mcc%d.gprs",
		imsi.GetMnc(), imsi.GetMcc()))

	//copy(p.value[:], oiBytes)
	p.value = p.value[0:0]
	p.value = append(p.value, oiBytes...)

	p.length = uint8(len(oiBytes))
}

func (p ApnOI) String() string {
	return fmt.Sprintf("%s", string(p.value[:p.length]))
}

// REFER to 23.003 9.1	Structure of APN
type Apn struct {
	ni       ApnNI // network id
	oi       ApnOI // operator id
	isOiPrst bool  //present flag for apn oi
}

func (p *ApnNI) Reset() {
	p.length = 0
	p.value = []uint8{}
}

func (p *ApnOI) Reset() {
	p.length = 0
	p.value = []byte{}
}
func (p *Apn) Reset() {
	p.ni.Reset()
	p.oi.Reset()
	p.isOiPrst = false
}

func (p *Apn) SetNI(ni *ApnNI) {
	p.ni = *ni
}

func (p *Apn) GetNI() ApnNI {
	return p.ni
}

func (p *Apn) SetOI(oi *ApnOI) {
	p.isOiPrst = true
	p.oi = *oi
}

func (p *Apn) GetOI() ApnOI {
	return p.oi
}

func (p Apn) String() string {
	rtStr := fmt.Sprintf("%s", p.ni)
	if p.isOiPrst {
		rtStr += fmt.Sprintf(".%s", p.oi)
	}
	return rtStr
}

func (p *Apn) Decode(msgBuf *bytes.Reader) error {

	var isApnNi bool = true
	var apnNi []byte
	var apnOi []byte

	//read the length byte
	length, _ := msgBuf.ReadByte()

	//read the first lable
	length, err := msgBuf.ReadByte()
	for err == nil {
		lable := make([]byte, length)
		if binary.Read(msgBuf, binary.BigEndian, lable) != nil {
			return fmt.Errorf("failed to decode dnn")
		}
		//check the label ans set he isApnNi flag
		rt, _ := regexp.Match(`mcc[\d]{3}`, lable)
		if rt {
			isApnNi = false
		}

		//store the lable
		if isApnNi {
			//store the apn ni
			apnNi = append(apnNi, lable[:]...)
			apnNi = append(apnNi, '.')
		} else {
			//store the apn oi
			apnOi = append(apnOi, lable[:]...)
			apnOi = append(apnOi, '.')
		}

		length, err = msgBuf.ReadByte()
	}

	if len(apnNi) > 0 {
		p.ni.SetValue(apnNi[:len(apnNi)-1])
	}
	if len(apnOi) > 0 {
		p.isOiPrst = true
		p.oi.SetValue(apnOi[:len(apnOi)-1])
	}

	return nil
}

//encode Apn to nas octet stream
func (p *Apn) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)

	var encBuf []byte

	encApnNi := EncodeLables(p.ni.value)
	encBuf = append(encBuf, encApnNi[:]...)

	if p.isOiPrst {
		encApnOi := EncodeLables(p.oi.value)
		encBuf = append(encBuf, encApnOi[:]...)
	}

	return encBuf
}

func (p *Apn) StoreWithString(val string) error {
	//convert string apn(ni.oi) to APN

	apnBytes := []byte(val)

	// eg. hello5g.mnc001.mcc460.gprs
	var apnNi []byte
	var apnOi []byte

	lables := bytes.Split(apnBytes, []byte("."))
	var index int
	for i, v := range lables {
		rt, _ := regexp.Match(`mnc[\d]{3}`, v)
		if rt {
			index = i
			break
		} else {
			index = 2
		}
	}

	// store apn ni
	for i := 0; i < (index - 1); i++ {
		apnNi = append(apnNi, lables[i]...)
		apnNi = append(apnNi, '.')
	}

	apnNi = apnNi[:len(apnNi)-1]
	err := p.ni.SetValue(apnNi)
	if err != nil {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.DEBUG, nil,
			"failed to store apn ni, erro(%s)", err)
	}

	if index <= len(lables) {
		for i := index - 1; i < len(lables); i++ {
			apnOi = append(apnOi, lables[i]...)
			apnOi = append(apnOi, '.')
		}
		apnOi = apnOi[:len(apnOi)-1]
		err = p.oi.SetValue(apnOi)
		if err != nil {
			rlogger.Trace(types.ModuleCmn3gtp, rlogger.DEBUG, nil, "failed to store apn oi, erro(%s)", err)
		}
		p.isOiPrst = true
	}
	return nil
}

func EncodeLables(vals []byte) []byte {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)

	var encBuf []byte
	var length uint8

	lables := bytes.Split(vals, []byte("."))

	for _, v := range lables {
		length = uint8(len(v))
		encBuf = append(encBuf, length)
		encBuf = append(encBuf, v[:]...)
	}

	return encBuf
}
