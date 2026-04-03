package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"unsafe"
)

// NasNonDeliveryIndicationMessage struct definition
type NasNonDeliveryIndicationMessage struct {
	RanUeNgapId uint32
	AmfUeNgapId uint64
	NasPdu      []byte
	CauseType   uint16
	CauseValue  uint16
	ctxt        codec.NgapOssCtxt
}

// NewNasNonDeliveryIndicationMessage create a new Message
func NewNasNonDeliveryIndicationMessage() *NasNonDeliveryIndicationMessage {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	return &NasNonDeliveryIndicationMessage{}
}

// SetOssCodecCtxt set OSS codec context
func (p *NasNonDeliveryIndicationMessage) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// Encode the message into bytes
func (p *NasNonDeliveryIndicationMessage) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	nasNonDeliveryIndMsg := codec.NewNasNonDeliveryIndicationCodec()
	defer codec.DeleteNasNonDeliveryIndicationCodec(nasNonDeliveryIndMsg)

	// ran ue ngap id
	nasNonDeliveryIndMsg.SetRanUeNgapId(uint(p.RanUeNgapId))
	// amf ue ngap id
	nasNonDeliveryIndMsg.SetAmfUeNgapId(uint64(p.AmfUeNgapId))
	// nas message
	nasNonDeliveryIndMsg.SetNasPdu(string(p.NasPdu))

	nasNonDeliveryIndMsg.SetCauseType(p.CauseType)
	nasNonDeliveryIndMsg.SetCauseValue(p.CauseValue)

	//encode
	msgBuffer := nasNonDeliveryIndMsg.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

// Decode the bytes into message
func (p *NasNonDeliveryIndicationMessage) Decode(msgbuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	nasNonDeliveryIndMsg := codec.NewNasNonDeliveryIndicationCodec()
	defer codec.DeleteNasNonDeliveryIndicationCodec(nasNonDeliveryIndMsg)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if nasNonDeliveryIndMsg.Decode(p.ctxt, msgBuffer) == true {
		p.RanUeNgapId = uint32(nasNonDeliveryIndMsg.GetRanUeNgapId())
		p.AmfUeNgapId = uint64(nasNonDeliveryIndMsg.GetAmfUeNgapId())
		p.NasPdu = []byte(nasNonDeliveryIndMsg.GetNasPdu())
		p.CauseType = nasNonDeliveryIndMsg.GetCauseType()
		p.CauseValue = nasNonDeliveryIndMsg.GetCauseValue()

	} else {
		return fmt.Errorf("failed to decode msg buffer")
	}

	return nil
}
