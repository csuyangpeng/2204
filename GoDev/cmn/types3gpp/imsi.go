package types3gpp

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"strings"
)

const (
	sizeofImsiBytes       = 8
	imsiMaxLen            = 15
	imsiMinLen            = 5
	towDigitMnc     uint8 = 2
	threeDigitMnc   uint8 = 3
	upperNibbleMask uint8 = 0xf0
	lowerNibbleMask uint8 = 0x0f

	imsiConvFactor10        = 10
	imsiConvFactor100       = 100
	upperNibbleShift  uint8 = 4
)

// IMSI (International Mobile Subscription Identity)
//         Imsi not more than 15 digits
// <----------------------------------------------->
// 3 digits   2 or 3
// <------->  <-----> <----------------------------->
//    MCC       MNC            MSISN
type Imsi struct {
	imsiArray      [sizeofImsiBytes]byte
	length         uint8
	numDigitsInMnc uint8
}

func CreateATestIMSI() Imsi {
	return Imsi{imsiArray: [sizeofImsiBytes]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}, length: 10, numDigitsInMnc: 3}
}

// -----------------------------------------------
// IMSI 123456789567893  (U/L - upper/lower Nibble)
// -----------------------------------------------
// bcdArray  0   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
// -----------------------------------------------
// Position  len |U L|U L|U L|U L|U L|U L|U L|U L|
// -----------------------------------------------
// Value     15  |1 2|3 4|5 6|7 8|9 5|6 7|8 9|3 0|
// bcdArr[0] contain the length of digits for Imsi
func (p *Imsi) StoreBcdArray(bcdArr []byte, mncDigitsLen uint8) error {

	if len(bcdArr) > (sizeofImsiBytes+1) ||
		len(bcdArr) < (imsiMinLen/2+1) {
		return fmt.Errorf("invalid lenght of bcdArray: "+
			"len(bcdArry)=%d", len(bcdArr))
	}

	if mncDigitsLen != towDigitMnc &&
		mncDigitsLen != threeDigitMnc {
		return fmt.Errorf("invalid mncDigitsLen: %d", mncDigitsLen)
	}

	length := int(bcdArr[0])
	p.numDigitsInMnc = mncDigitsLen
	p.length = uint8(length)
	var arraySize int
	if length%2 == 0 { //odd number of digit
		arraySize = (length + 1) / 2
	} else {
		arraySize = length / 2
	}

	for i := 0; i < arraySize; i++ {
		p.imsiArray[i] = bcdArr[i+1]
	}

	if length%2 == 1 {
		p.imsiArray[arraySize] = bcdArr[arraySize+1] & upperNibbleMask
	}

	return nil
}
func (p *Imsi) StoreImsiString(imsiStr string, mncDigitsLen uint8) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	imsiLen := len(imsiStr)
	if imsiLen > imsiMaxLen ||
		imsiLen < imsiMinLen {
		return fmt.Errorf("invalid imsi string, len "+
			"should be [6-15], but length is %d", imsiLen)
	}

	byteLen := imsiLen / 2
	isOdd := imsiLen % 2

	for i := 0; i < byteLen; i++ {
		upper := imsiStr[2*i] - '0'
		lower := imsiStr[2*i+1] - '0'
		p.imsiArray[i] = (upper << upperNibbleShift) | lower
	}

	if isOdd == 1 {
		p.imsiArray[byteLen] = imsiStr[2*byteLen] << upperNibbleShift
	}

	p.length = uint8(imsiLen)
	p.numDigitsInMnc = mncDigitsLen
	return nil
}

func (p *Imsi) GetMcc() uint16 {
	return (uint16(p.imsiArray[0]>>upperNibbleShift) * imsiConvFactor100) +
		(uint16(p.imsiArray[0]&lowerNibbleMask) * imsiConvFactor10) +
		uint16(p.imsiArray[1]>>upperNibbleShift)
}
func (p *Imsi) GetMccBytes() []byte {
	var mcc []byte
	mcc = append(mcc, uint8(p.imsiArray[0]>>upperNibbleShift))
	mcc = append(mcc, uint8(p.imsiArray[0]&lowerNibbleMask))
	mcc = append(mcc, uint8(p.imsiArray[1]>>upperNibbleShift))
	return mcc
}
func (p *Imsi) GetPlmnStr() string {
	if p.numDigitsInMnc == threeDigitMnc {
		return p.String()[0:6]
	} else {
		return p.String()[0:5]
	}
}

func (p *Imsi) GetMnc() uint16 {
	var mnc uint16
	if p.numDigitsInMnc == threeDigitMnc {
		mnc = (uint16(p.imsiArray[1]&lowerNibbleMask) * imsiConvFactor100) +
			(uint16(p.imsiArray[2]>>upperNibbleShift) * imsiConvFactor10) +
			uint16(p.imsiArray[2]&lowerNibbleMask)
	} else if p.numDigitsInMnc == towDigitMnc {
		mnc = (uint16(p.imsiArray[1]&lowerNibbleMask) * imsiConvFactor10) +
			uint16(p.imsiArray[2]>>upperNibbleShift)
	}
	return mnc
}
func (p *Imsi) GetMncBytes() []byte {
	var mnc []byte
	if p.numDigitsInMnc == threeDigitMnc {
		mnc = append(mnc, p.imsiArray[1]&lowerNibbleMask)
		mnc = append(mnc, p.imsiArray[2]>>upperNibbleShift)
		mnc = append(mnc, p.imsiArray[2]&lowerNibbleMask)
	} else if p.numDigitsInMnc == towDigitMnc {
		mnc = append(mnc, p.imsiArray[1]&lowerNibbleMask)
		mnc = append(mnc, p.imsiArray[2]>>upperNibbleShift)
	}
	return mnc
}
func (p *Imsi) GetMsIn() []byte {
	var msin []byte
	/*if p.numDigitsInMnc == threeDigitMnc {
		msin = append(msin, p.imsiArray[3]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[3]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[4]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[4]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[5]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[5]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[6]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[6]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[7]>>upperNibbleShift)
		// max len 9
		//msin = append(msin, p.imsiArray[7]&lowerNibbleMask)
	} else if p.numDigitsInMnc == towDigitMnc {
		msin = append(msin, p.imsiArray[2]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[3]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[3]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[4]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[4]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[5]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[5]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[6]>>upperNibbleShift)
		msin = append(msin, p.imsiArray[6]&lowerNibbleMask)

		msin = append(msin, p.imsiArray[7]>>upperNibbleShift)
	}*/
	var startIndex uint8
	if p.numDigitsInMnc == threeDigitMnc {
		startIndex = 6
	} else if p.numDigitsInMnc == towDigitMnc {
		startIndex = 5
	}
	//todo:
	for i := startIndex; i < p.length; i++ {
		if i%2 == 1 {
			msin = append(msin, p.imsiArray[(i-1)/2]&lowerNibbleMask)
		} else {
			msin = append(msin, p.imsiArray[i/2]>>upperNibbleShift)
		}
	}

	return msin
}
func (p *Imsi) GetMsInValue() uint64 {
	var msin uint64

	var startIndex uint8
	if p.numDigitsInMnc == threeDigitMnc {
		startIndex = 6
	} else if p.numDigitsInMnc == towDigitMnc {
		startIndex = 5
	} else {
		return msin
	}

	for i := startIndex; i < p.length; i++ {
		if i%2 == 1 {
			msin = msin*imsiConvFactor10 +
				uint64(p.imsiArray[(i-1)/2]&lowerNibbleMask)
		} else {
			msin = msin*imsiConvFactor10 +
				//uint64(p.imsiArray[i/2]&upperNibbleShift)
				uint64(p.imsiArray[i/2]>>upperNibbleShift)
		}
	}
	return msin
}

func (p *Imsi) GetLength() uint8 {
	return p.length
}

func (p *Imsi) GetSizeofMnc() uint8 {
	return p.numDigitsInMnc
}

func (p *Imsi) AddIMSIPrefix() string {
	return "imsi-" + p.String()
}

func CheckMncLen(imsi string) uint8 {
	//todo checkt the mnc length
	return towDigitMnc
}
func (p *Imsi) String() string {
	mcc := byteString(p.GetMccBytes())
	mnc := byteString(p.GetMncBytes())
	msin := byteString(p.GetMsIn())

	return fmt.Sprintf("%s%s%s", mcc, mnc, msin)
}

func (p *Imsi) GetValue() uint64 {
	var val uint64
	var i uint8
	for i = 0; i < p.length; i++ {
		if i%2 == 1 {
			val = val*imsiConvFactor10 +
				uint64(p.imsiArray[(i-1)/2]&lowerNibbleMask)
		} else {
			val = val*imsiConvFactor10 +
				uint64(p.imsiArray[i/2]>>upperNibbleShift)
		}
	}
	return val
}

// Spec TS23.003 28.7.2
// <IMSI>@nai.5gc.mnc<MNC>.mcc<MCC>.3gppnetwork.org
//
func (p *Imsi) GetNAIString() string {
	return fmt.Sprintf("%d@nai.5gc.mnc%03d.mcc%d.3gppnetwork.org",
		p.GetValue(), p.GetMnc(), p.GetMcc())
}

func (p *Imsi) StoreWithNAI(nai string, mncDigitsLen uint8) error {

	split := strings.Split(nai, "@")
	if len(split) != 2 {
		return fmt.Errorf("invalid imsi")
	}

	return p.StoreImsiString(split[0], mncDigitsLen)
}
