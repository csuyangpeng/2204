package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 8.2.21
type IdentityResponseMsg struct {
	//Extended protocol discriminator	9.2	      M	V	1
	//Security header type	            9.3	      M	V	1/2
	//Spare half octet	                9.5	      M	V	1/2
	//Identity request message type     9.7	      M	V	1
	Header nas.MmNasMessageHeader

	//Mobile identity	                9.11.3.4  M	LV-E	3-n
	MobileIdentity nasie.MobileIdentity
}

func (p *IdentityResponseMsg) Reset() {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	p.MobileIdentity.Reset()
}

func (p *IdentityResponseMsg) String() string {
	var msgStr string
	msgStr = "Identity Response Msg ( "
	msgStr += fmt.Sprintf("Header(%v) ", p.Header)
	msgStr += fmt.Sprintf("MobileIdentity(%s) ", p.MobileIdentity)
	return msgStr
}

func (p *IdentityResponseMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	// 5GS mobile identity: chapter 9.11.3.4	5GS mobile identity
	err := p.MobileIdentity.Decode(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode mobile identity,  %s", err)
	}

	return nil
}

func (p *IdentityResponseMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	var encBuf []byte

	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.IdentifyRequest
	encBuf = append(encBuf, header.Encode()...)

	//MobileIdentity, it's format is LV
	var mobileId []byte
	mobOctetValue, err := p.MobileIdentity.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "failed to encode mobile id, err(%s)", err)
		return nil, err
	}
	// L
	lenMobID := make([]byte, 2)
	binary.BigEndian.PutUint16(lenMobID, uint16(len(mobOctetValue)))
	mobileId = append(mobileId, lenMobID...)
	// V
	mobileId = append(mobileId, mobOctetValue...)
	encBuf = append(encBuf, mobileId...)

	return encBuf, nil
}
