package nasmsg

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 f10 8.2.13
type DeRegistrationAcceptMsg struct {
	//Extended protocol discriminator
	//Security header type
	//Spare half octet
	//De-registration request message identity
}

//encode a deRegistration request msg from nas octet stream
func (p *DeRegistrationAcceptMsg) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	//header
	var header nas.MmNasMessageHeader
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.DeregistrationAcceptUe

	encBuf = append(encBuf, byte(header.ExtendProtoDisc))
	encBuf = append(encBuf, byte(header.SecurityHeaderType)<<4)
	encBuf = append(encBuf, byte(header.MessageType))

	return encBuf
}
