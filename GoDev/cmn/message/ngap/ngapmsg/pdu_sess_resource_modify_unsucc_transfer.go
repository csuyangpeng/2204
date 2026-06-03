package ngapmsg

import "C"
import (
	"fmt"
	"lite5gc/cmn/types3gpp"
	"unsafe"

	codec "lite5gc/cmn/message/ngap/ngapcodec"
)

type PduSessResMdfyUnSuccTransfer struct {
	Cause types3gpp.NgapCause

	ctxt codec.NgapOssCtxt
}

func (p PduSessResMdfyUnSuccTransfer) String() string {
	return fmt.Sprintf("%s", p.Cause)
}

// NewCauseTransfer create a cause
func NewPduSessResMdfyUnSuccTransfer() *PduSessResMdfyUnSuccTransfer {
	return &PduSessResMdfyUnSuccTransfer{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResMdfyUnSuccTransfer) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *PduSessResMdfyUnSuccTransfer) Encode() []byte {
	transfer := codec.NewPduSessResMdfyUnSuccTransferCodec()
	defer codec.DeletePduSessResMdfyUnSuccTransferCodec(transfer)

	transfer.SetCause(byte(p.Cause.Type), byte(p.Cause.Value))

	msgBuffer := transfer.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResMdfyUnSuccTransfer) Decode(msgbuf []byte) error {

	transfer := codec.NewPduSessResMdfyUnSuccTransferCodec()
	defer codec.DeletePduSessResMdfyUnSuccTransferCodec(transfer)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if transfer.Decode(p.ctxt, msgBuffer) == true {
		p.Cause.Type = types3gpp.CauseType(transfer.GetCauseType())
		p.Cause.Value = types3gpp.CauseValue(transfer.GetCauseValue())
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
