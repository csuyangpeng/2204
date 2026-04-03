package ngapmsg

import "C"
import (
	"fmt"
	"lite5gc/cmn/types3gpp"

	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"unsafe"
)

type PduSessResSetupUnSuccTransfer struct {
	Cause types3gpp.NgapCause

	ctxt codec.NgapOssCtxt
}

func (p PduSessResSetupUnSuccTransfer) String() string {
	return fmt.Sprintf("%s", p.Cause)
}

// NewCauseTransfer create a cause
func NewPduSessResSetupUnSuccTransfer() *PduSessResSetupUnSuccTransfer {
	return &PduSessResSetupUnSuccTransfer{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResSetupUnSuccTransfer) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *PduSessResSetupUnSuccTransfer) Encode() []byte {
	transfer := codec.NewPduSessResSetupUnSuccTransferCodec()
	defer codec.DeletePduSessResSetupUnSuccTransferCodec(transfer)

	transfer.SetCause(byte(p.Cause.Type), byte(p.Cause.Value))

	msgBuffer := transfer.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResSetupUnSuccTransfer) Decode(msgbuf []byte) error {

	transfer := codec.NewPduSessResSetupUnSuccTransferCodec()
	defer codec.DeletePduSessResSetupUnSuccTransferCodec(transfer)

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
