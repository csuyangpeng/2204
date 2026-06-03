package nasie

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
)

type IdentityType byte

//Type of identity (octet 3)
//Bits
//3	2	1
//0	0	0	No identity (NOTE)
//0	0	1	SUCI
//0	1	1	IMEI
//1	0	0	5G-S-TMSI
//1	0	1	IMEISVI
//1	1	0	5G-GUTI
const (
	NoIdentity IdentityType = 0
	Suci       IdentityType = 1
	Guti5g     IdentityType = 2
	Imei       IdentityType = 3
	STmsi5g    IdentityType = 4
	ImeiSvi    IdentityType = 5
)

type MobileIdentity struct {
	IdType  IdentityType
	Suci    *types3gpp.Suci
	Guti5g  *types3gpp.Guti5G
	Stmsi5g *types3gpp.Stmsi5G
}

func (p IdentityType) String() string {
	switch p {
	case NoIdentity:
		return "No Identity"
	case Suci:
		return "Suci"
	case Guti5g:
		return "5g guti"
	case Imei:
		return "Imei"
	case STmsi5g:
		return "STmsi5g"
	case ImeiSvi:
		return "ImeiSvi"
	}
	return ""
}
func (p *MobileIdentity) Reset() {
	p.IdType = 0
	p.Suci = nil
	p.Guti5g = nil
	p.Stmsi5g = nil
}
func (p *MobileIdentity) Encode() ([]byte, error) {
	// NOTICE :
	// when calling this Encode function for MobileIdentity
	// if MobileIdentity.Format = "LV" , encBuf[0] does not change.
	// if MobileIdentity.Format = "TLV" , encBuf[0] = encBuf[0] + 1
	// 1 means the length of octet 1 (Title: 5GS mobile identity IEI)
	var encBuf []byte

	var octet3 byte

	switch p.IdType {
	case Suci:
		//3GPP TS 24.501 V15.1.0 (2018-09)
		octet3 = 0 // 24.501 9.11.3.4.1 SUPI format (octet 3, bits 5-7), 0:IMSI
		octet3 |= byte(p.IdType)
		encBuf = append(encBuf, octet3)
		suciBytes, err := p.Suci.Encode()
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to encode suci")
			return nil, fmt.Errorf("failed to encode suci")
		}
		encBuf = append(encBuf, suciBytes...)
	case Guti5g:
		octet3 = 0xF0
		octet3 |= byte(p.IdType)
		encBuf = append(encBuf, octet3)
		gutiBytes := p.Guti5g.Encode()
		encBuf = append(encBuf, gutiBytes...)
	case STmsi5g:
		octet3 = 0xF0
		octet3 |= byte(p.IdType)
		encBuf = append(encBuf, octet3)
		stmsiBytes := p.Stmsi5g.Encode()
		encBuf = append(encBuf, stmsiBytes...)
	default:
		return nil, nas.ErrInvalidMobileIdType
	}
	return encBuf, nil
}

func (p *MobileIdentity) Decode(msgBuf *bytes.Reader) error {
	// read the 2 bytes: length of mobileID
	msgBuf.ReadByte()
	msgBuf.ReadByte()

	//octet 3
	idTypeOctet, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "failed to decode mobile identity, %s", err)
		return fmt.Errorf("failed to decode mobile identity, %s", err)
	}
	idType, _ := utils.GetBitsValue(idTypeOctet, 1, 3)
	p.IdType = IdentityType(idType)

	//Withdraw 3 bytes, Let the length value pass in to the next process
	msgBuf.UnreadByte()
	msgBuf.UnreadByte()
	msgBuf.UnreadByte()

	switch p.IdType {
	case Suci:
		p.Suci = new(types3gpp.Suci)
		//err = p.Suci.Decode(msgBuf)
		// todo sidf add
		err = p.Suci.DecodeConcealed(msgBuf)
		if err != nil {
			return fmt.Errorf("failed to decode SUCI ie, err: %s", err)
		}
	case Guti5g:
		p.Guti5g = new(types3gpp.Guti5G)
		err = p.Guti5g.Decode(msgBuf)
		if err != nil {
			return fmt.Errorf("failed to decode 5G Guti, %s", err)
		}
	case STmsi5g:
		p.Stmsi5g = new(types3gpp.Stmsi5G)
		err = p.Stmsi5g.Decode(msgBuf)
		if err != nil {
			return fmt.Errorf("failed to decode STmsi5G, %s", err)
		}
	default:
		return fmt.Errorf("invalid identity type (%d)", p.IdType)
	}
	return nil
}

func (p MobileIdentity) String() string {
	return fmt.Sprintf("todo")
}
