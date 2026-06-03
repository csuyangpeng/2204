package nasmsg

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

//24501 8.2.21
type IdentityRequestMsg struct {
	//Extended protocol discriminator	9.2	      M	V	1
	//Security header type	            9.3	      M	V	1/2
	//Spare half octet	                9.5	      M	V	1/2
	//Identity request message type     9.7	      M	V	1
	Header nas.MmNasMessageHeader

	//Identity type	                    9.11.3.3  M	V	1/2
	//Spare half octet	                9.5	      M	V	1/2
	IdentityType nasie.IdentityType
}

func (p *IdentityRequestMsg) Reset() {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	p.IdentityType = nasie.NoIdentity
}

func (p *IdentityRequestMsg) String() string {
	var msgStr string
	msgStr = "Identity Request Msg ( "
	msgStr += fmt.Sprintf("Header(%v) ", p.Header)
	msgStr += fmt.Sprintf("IdentityType(%v) ", p.IdentityType)
	return msgStr
}

func (p *IdentityRequestMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	idType, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read IdentityType,err(%s)", err)
	}

	itype, _ := utils.GetBitsValue(idType, 1, 4)
	p.IdentityType = nasie.IdentityType(itype)

	return nil
}

func (p *IdentityRequestMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	var encBuf []byte

	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.IdentifyRequest
	encBuf = append(encBuf, header.Encode()...)

	//IdentityType
	encBuf = append(encBuf, byte(p.IdentityType))

	return encBuf, nil
}
