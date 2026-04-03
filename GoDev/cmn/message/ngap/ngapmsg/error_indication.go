package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"unsafe"
)

//38413  9.2.6.13
// ErrIndicationMsg struct definition
type ErrIndicationMsg struct {
	//optional
	AmfUeNgapId       uint32
	IsAmfUeNgapIdPrst bool

	RanUeNgapId       uint32
	IsRanUeNgapIdPrst bool

	CauseType   uint16
	CauseValue  uint16
	IsCausePrst bool

	ctxt codec.NgapOssCtxt
}

// NewErrIndicationMsg create a new Message
func NewErrIndicationMsg() *ErrIndicationMsg {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)
	return &ErrIndicationMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *ErrIndicationMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *ErrIndicationMsg) String() string {
	rtStr := fmt.Sprintf("ErrIndicationMsg{RanUeNGAPIdPrst(%v),RanUeNGAPId(%d),"+
		"AmfUeNGAPIdPrst(%v),AmfUeNGAPId(%d),CausePrst(%v),CauseT&V(%d,%d)",
		p.IsRanUeNgapIdPrst,
		p.RanUeNgapId,
		p.IsAmfUeNgapIdPrst,
		p.AmfUeNgapId,
		p.IsCausePrst,
		p.CauseType,
		p.CauseValue)
	return rtStr + "}"
}

// Encode the message into bytes
func (p *ErrIndicationMsg) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	errIndicationMsg := codec.NewErrorIndicationCodec()
	defer codec.DeleteErrorIndicationCodec(errIndicationMsg)

	// ran ue ngap id
	if p.IsRanUeNgapIdPrst {
		errIndicationMsg.SetRanUeNgapId(uint(p.RanUeNgapId))
	}

	// amf ue ngap id
	if p.IsAmfUeNgapIdPrst {
		errIndicationMsg.SetAmfUeNgapId(uint64(p.AmfUeNgapId))
	}

	// cause type and value
	if p.IsCausePrst {
		errIndicationMsg.SetCauseType(p.CauseType)
		errIndicationMsg.SetCauseValue(p.CauseValue)
	}

	//encode
	msgBuffer := errIndicationMsg.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

// Decode the bytes into message
func (p *ErrIndicationMsg) Decode(msgbuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	errIndicationMsg := codec.NewErrorIndicationCodec()
	defer codec.DeleteErrorIndicationCodec(errIndicationMsg)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if errIndicationMsg.Decode(p.ctxt, msgBuffer) == true {
		if errIndicationMsg.IsAmfUeNgapIdPrst() {
			p.IsAmfUeNgapIdPrst = true
			p.AmfUeNgapId = uint32(errIndicationMsg.GetAmfUeNgapId())
		}
		if errIndicationMsg.IsRanUeNgapIdPrst() {
			p.IsRanUeNgapIdPrst = true
			p.RanUeNgapId = uint32(errIndicationMsg.GetRanUeNgapId())
		}
		if errIndicationMsg.IsCauseTypePrst() {
			p.IsCausePrst = true
			p.CauseType = errIndicationMsg.GetCauseType()
			p.CauseValue = errIndicationMsg.GetCauseValue()
		}

	} else {
		return fmt.Errorf("failed to decode msg buffer")
	}

	return nil
}
