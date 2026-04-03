package nasie

import (
	"bytes"
	"encoding/binary"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

// 24.301 9.11.3.1

type MmCapability5G struct {
	LLP         bool
	SupS1Mode   bool
	SupHoAttach bool
}

// EPC NAS supported (S1 mode) (octet 3, bit 1)
// 0				S1 mode not supported
// 1				S1 mode supported

// ATTACH REQUEST message containing PDN CONNECTIVITY REQUEST message for handover support (HO attach) (octet 3, bit 2)
// 0				ATTACH REQUEST message containing PDN CONNECTIVITY REQUEST message
// with request type set to "handover"
// or "handover of emergency bearer services" to transfer PDU session from N1 mode to S1 mode not supported
// 1				ATTACH REQUEST message containing PDN CONNECTIVITY REQUEST message
// with request type set to "handover"
// or "handover of emergency bearer services" to transfer PDU session from N1 mode to S1 mode supported

// LTE Positioning Protocol (LPP) capability (octet 3, bit 3)
// 0				LPP in N1 mode not supported
// 1				LPP in N1 mode supported (see 3GPP TS 36.355 [26])

// All other bits in octet 3 to 15 are spare and shall be coded as zero,
// if the respective octet is included in the information element.

func (p *MmCapability5G) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)

	length, _ := msgBuf.ReadByte()
	if length < 1 || length > 13 {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "invalid length for "+
			"Ie 5gmm capability (%d)", length)
		return nas.ErrInvalidIeLength
	}

	octet, _ := msgBuf.ReadByte()
	p.SupS1Mode, _ = utils.GetBitValue(octet, 1)
	p.SupHoAttach, _ = utils.GetBitValue(octet, 2)
	p.LLP, _ = utils.GetBitValue(octet, 3)

	if (int(length) - 3) > 0 {
		leftBytes := make([]byte, length-3)
		binary.Read(msgBuf, binary.BigEndian, leftBytes)
	}
	return nil
}
func (p *MmCapability5G) Reset() {
	p.LLP = false
	p.SupS1Mode = false
	p.SupHoAttach = false
}
