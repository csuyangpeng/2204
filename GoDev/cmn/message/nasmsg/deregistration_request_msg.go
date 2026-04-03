package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

//24501 f10 8.2.12  De-registration request (UE originating de-registration)
type DeRegistrationRequestMsg struct {
	//Extended protocol discriminator
	//Security header type
	//Spare half octet
	//De-registration request message identity
	DeRegistrationType nasie.DeRegistrationTypeIE
	NgKSI              nasie.NasKSI
	MobIDGuti          nasie.MobileIdentity
}

func (p *DeRegistrationRequestMsg) Reset() {
	p.DeRegistrationType.Reset()
	p.NgKSI.Reset()
	p.MobIDGuti.Reset()

}

//encode a deRegistration request msg from nas octet stream
func (p *DeRegistrationRequestMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	//header
	var header nas.MmNasMessageHeader
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.DeregistrationRequestUeT
	encBuf = append(encBuf, byte(header.ExtendProtoDisc))
	encBuf = append(encBuf, byte(header.SecurityHeaderType)<<4)
	encBuf = append(encBuf, byte(header.MessageType))

	// De-registration type && ngKS 共用一个字节
	deRegistNgKSIOctet := byte(utils.BoolToByte(p.NgKSI.Tsc)) << 7
	deRegistNgKSIOctet |= p.NgKSI.Ksi << 4

	// De-registration type
	deRegistNgKSIOctet |= byte(utils.BoolToByte(p.DeRegistrationType.SwithOff)) << 3
	deRegistNgKSIOctet |= byte(utils.BoolToByte(p.DeRegistrationType.IsReRegistrationRequired)) << 2
	deRegistNgKSIOctet |= byte(p.DeRegistrationType.AccessType)

	encBuf = append(encBuf, deRegistNgKSIOctet)

	//MobileIdentity, it's format is LV
	var mobileId []byte
	mobOctetValue, err := p.MobIDGuti.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "failed to encode mobile id, err(%s)", err)
		return nil, err
	}
	// L
	lenID := make([]byte, 2)
	binary.BigEndian.PutUint16(lenID, uint16(len(mobOctetValue)))
	rlogger.Trace(types.ModuleCmnMsg, rlogger.INFO, nil, "guti length: %d", lenID)
	mobileId = append(mobileId, lenID...)
	// V
	mobileId = append(mobileId, mobOctetValue...)
	rlogger.Trace(types.ModuleCmnMsg, rlogger.INFO, nil, "guti: %d", mobileId)
	encBuf = append(encBuf, mobileId...)

	return encBuf, nil
}

// decode a registration request msg from nas octet stream
func (p *DeRegistrationRequestMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	//De-registration type && ngKSI 共用一个字节
	DeReType, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read de registration type")
	}
	accessBytes, err := utils.GetBitsValue(DeReType, 1, 2)
	if err != nil {
		return fmt.Errorf("failed to read AccessType")
	}
	p.DeRegistrationType.AccessType = nasie.AccessTypes(accessBytes)
	p.DeRegistrationType.IsReRegistrationRequired, _ = utils.GetBitValue(DeReType, 3)
	p.DeRegistrationType.SwithOff, _ = utils.GetBitValue(DeReType, 4)
	p.NgKSI.Tsc, _ = utils.GetBitValue(DeReType, 5)
	p.NgKSI.Ksi, _ = utils.GetBitsValue(DeReType, 6, 8)

	// 5GS mobile identity: chapter 9.11.3.4	5GS mobile identity
	err = p.MobIDGuti.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "failed to decode mobile identity, %s", err)
		return fmt.Errorf("failed to decode mobile identity, %s", err)
	}

	return nil
}
