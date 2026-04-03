package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	t3gpp "lite5gc/cmn/types3gpp"
	"unsafe"
)

// NgResetMsg struct definition
type NgResetMsg struct {
	CauseType       uint16
	CauseValue      uint16
	ResetTypeChoice t3gpp.ResetType

	ctxt codec.NgapOssCtxt
}

// NewNgResetMsg create a new Message
func NewNgResetMsg() *NgResetMsg {
	return &NgResetMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *NgResetMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *NgResetMsg) String() string {
	rtStr := fmt.Sprintf("NgResetMsg: \n"+
		"ResetType(%s)\n", p.ResetTypeChoice.String())

	CauseTypeStr := fmt.Sprintf("NgSetupFailureMsg: \n"+
		"CauseType(%d)\n",
		p.CauseType)
	rtStr += CauseTypeStr

	CauseValueStr := fmt.Sprintf("NgSetupFailureMsg: \n"+
		"CauseValue(%d)\n",
		p.CauseValue)
	rtStr += CauseValueStr

	return rtStr
}

func (p *NgResetMsg) Encode() []byte {

	resetCodec := codec.NewNgResetCodec()
	defer codec.DeleteNgResetCodec(resetCodec)

	//cause type and value
	resetCodec.SetCauseType(p.CauseType)
	resetCodec.SetCauseValue(p.CauseValue)

	//reset type
	resetTypeChoice := codec.NewResetTypeChoice()
	defer codec.DeleteResetTypeChoice(resetTypeChoice)
	resetTypeChoice.SetNgInterface(uint8(p.ResetTypeChoice.NgInterface))
	resetTypeChoice.SetNgInterfacePrst(p.ResetTypeChoice.IsNgInterfacePrst)
	resetCodec.SetResetType(resetTypeChoice)

	msgBuf := resetCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuf)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *NgResetMsg) Decode(msgbuf []byte) error {

	resetCodec := codec.NewNgResetCodec()
	defer codec.DeleteNgResetCodec(resetCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if resetCodec.Decode(p.ctxt, msgBuffer) == true {
		//CauseType and CauseValue
		p.CauseType = resetCodec.GetCauseType()
		p.CauseValue = resetCodec.GetCauseValue()

		//ResetType
		p.ResetTypeChoice.IsNgInterfacePrst = resetCodec.GetResetType().GetNgInterfacePrst()
		p.ResetTypeChoice.NgInterface = t3gpp.NgInterfaceReset(resetCodec.GetResetType().GetNgInterface())

	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
