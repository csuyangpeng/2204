package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"unsafe"
)

// NgSetupFailureMsg struct definition
type NgSetupFailureMsg struct {
	CauseType          uint16
	CauseValue         uint16
	RelativeTimeToWait uint32

	ctxt codec.NgapOssCtxt
}

// NewNgSetupFailureMsg create a new Message
func NewNgSetupFailureMsg() *NgSetupFailureMsg {
	return &NgSetupFailureMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *NgSetupFailureMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *NgSetupFailureMsg) String() string {
	rtStr := fmt.Sprintf("NgSetupFailureMsg: \n"+
		"RelativeTimeToWait(%d)\n",
		p.RelativeTimeToWait)

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

func (p *NgSetupFailureMsg) Encode() []byte {

	setupFailCodec := codec.NewNgSetupFailureCodec()
	defer codec.DeleteNgSetupFailureCodec(setupFailCodec)

	setupFailCodec.SetRelativeTimeToWait(uint(p.RelativeTimeToWait))
	setupFailCodec.SetCauseType(p.CauseType)
	setupFailCodec.SetCauseValue(p.CauseValue)

	msgBuffer := setupFailCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *NgSetupFailureMsg) Decode(msgbuf []byte) error {

	setupFailCodec := codec.NewNgSetupFailureCodec()
	defer codec.DeleteNgSetupFailureCodec(setupFailCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if setupFailCodec.Decode(p.ctxt, msgBuffer) == true {
		//RelativeTimeToWait
		p.RelativeTimeToWait = uint32(setupFailCodec.GetRelativeTimeToWait())

		//CauseType and CauseValue
		p.CauseType = setupFailCodec.GetCauseType()
		p.CauseValue = setupFailCodec.GetCauseValue()

	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
