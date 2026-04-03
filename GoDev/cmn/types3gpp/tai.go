package types3gpp

import (
	"encoding/binary"
	"fmt"
)

type TAC [SizeofTAC]byte

func (p TAC) GetByteSlice() []byte {
	var tac []byte
	for i := 0; i < SizeofTAC; i++ {
		tac = append(tac, p[i])
	}
	return tac
}

func (p *TAC) SetTac(val uint32) {
	tac := make([]byte, 4)
	binary.BigEndian.PutUint32(tac, val)

	for i := 0; i < SizeofTAC; i++ {
		p[i] = tac[i+1]
	}
}

func (p TAC) String() string {

	val := make([]byte, 4)
	for i := 0; i < SizeofTAC; i++ {
		val[i+1] = p[i]
	}

	tac := binary.BigEndian.Uint32(val)
	return fmt.Sprintf("tac(%d)", tac)
}

// TAI : Tracking Area Identity
type TAI struct {
	Plmn PlmnID
	// Tracking Area Code (TAC) is a fixed length code (of 3 octets) identifying a Tracking Area within a PLMN.
	// This part of the tracking area identification shall be coded using a full hexadecimal representation.
	Tac TAC
}

func (p *TAI) Reset() {
	p.Plmn.Reset()
	p.Tac = [SizeofTAC]byte{}
}

func (t1 *TAI) ComPare(t2 *TAI) bool {
	if t1.Plmn.GetValue(BigEndian) == t2.Plmn.GetValue(BigEndian) &&
		t1.Tac == t2.Tac {
		return true
	}
	return false
}

// SetTAI : set tai
func (t *TAI) SetTAI(p PlmnID, tac [SizeofTac]byte) {
	t.Plmn = p
	t.Tac = tac
}

func (p *TAI) Encode() []byte {
	var taiOctet []byte
	// octet 2~4
	t := p.Plmn.GetValue(BigEndian) // only for temporary conversion
	taiOctet = append(taiOctet, t[:]...)
	// octet 2~4
	taiOctet = append(taiOctet, p.Tac[:]...)
	return taiOctet
}

func (p *TAI) String() string {
	return fmt.Sprintf("tai[%s-%s]", p.Plmn, p.Tac)
}
