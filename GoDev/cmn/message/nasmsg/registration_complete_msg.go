package nasmsg

import (
	"bytes"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type RegistrationCompleteMsg struct {
	//Mandatory
	//MsgHeader          nas.MmNasMessageHeader

	//optional
	//SorTransContainer
}

func (p *RegistrationCompleteMsg) Reset() {
}

//encode a registration cmp msg from nas octet stream
func (p *RegistrationCompleteMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.RegistrationComplete
	encBuf = header.Encode()
	return encBuf, nil
}

func (p *RegistrationCompleteMsg) Decode(plainNasMsg *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	return nil
}
