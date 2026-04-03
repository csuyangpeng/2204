package nasie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// 24.501 9.11.3.57
// bit equals 0, indicates that no uplink data are pending for the corresponding PDU session identity.
//type UplinkDataStatus struct {
//	PsiA byte //psi(1)-bit2, psi(7)-bit8
//	PsiB byte //psi(8)-bit1, psi(15)-bit8
//}

// 24.501 9.11.3.44,
// bit equals 0, indicates that the 5GSM state of the corresponding PDU session is PDU SESSION INACTIVE.
//type PduSessionStatus struct {
//	PsiA byte //psi(1)-bit2, psi(7)-bit8
//	PsiB byte //psi(8)-bit1, psi(15)-bit8
//}
// 24.501 9.11.3.13 not supported currently
// indicate to the network user-plane resources of PDU sessions associated with
// non-3GPP access that are allowed to be re-established over 3GPP access
//
//type AllowedPduSessStat struct {
//	PsiA byte //psi(1)-bit2, psi(7)-bit8
//	PsiB byte //psi(8)-bit1, psi(15)-bit8
//}

type SessionStatus struct {
	PsiA byte //psi(1)-bit2, psi(7)-bit8
	PsiB byte //psi(8)-bit1, psi(15)-bit8
}

func (p *SessionStatus) String() string {
	pa := Byte2bit(p.PsiA)
	s := fmt.Sprintf("psi 1:(%d), ", pa[6])
	s += fmt.Sprintf("psi 2:(%d), ", pa[5])
	s += fmt.Sprintf("psi 3:(%d), ", pa[4])
	s += fmt.Sprintf("psi 4:(%d), ", pa[3])
	s += fmt.Sprintf("psi 5:(%d), ", pa[2])
	s += fmt.Sprintf("psi 6:(%d), ", pa[1])
	s += fmt.Sprintf("psi 7:(%d), ", pa[0])
	pb := Byte2bit(p.PsiB)
	s += fmt.Sprintf("psi 8:(%d), ", pb[7])
	s += fmt.Sprintf("psi 9:(%d), ", pb[6])
	s += fmt.Sprintf("psi 10:(%d), ", pb[5])
	s += fmt.Sprintf("psi 11:(%d), ", pb[4])
	s += fmt.Sprintf("psi 12:(%d), ", pb[3])
	s += fmt.Sprintf("psi 13:(%d), ", pb[2])
	s += fmt.Sprintf("psi 14:(%d), ", pb[1])
	s += fmt.Sprintf("psi 15:(%d)  ", pb[0])
	return s
}

func Byte2bit(p byte) []byte {
	var dst []byte
	for i := 0; i < 8; i++ {
		move := uint(7 - i)
		dst = append(dst, byte((p>>move)&1))
	}
	return dst
}

func (p *SessionStatus) Reset() {
	p.PsiA = 0
	p.PsiB = 0
}

func (p *SessionStatus) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	encBuf = append(encBuf, p.PsiA)
	encBuf = append(encBuf, p.PsiB)
	return encBuf
}
func (p *SessionStatus) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)

	length, _ := msgBuf.ReadByte()

	if length < 2 || length > 32 {
		return fmt.Errorf("invalid length for session status ((%d))", length)
	}

	p.PsiA, _ = msgBuf.ReadByte()
	p.PsiB, _ = msgBuf.ReadByte()

	//if there is still has spare bytes, ...
	if length-2 > 0 {
		leftBytes := make([]byte, length-2)
		binary.Read(msgBuf, binary.BigEndian, leftBytes)
		//leftBytes todo
	}

	return nil
}
