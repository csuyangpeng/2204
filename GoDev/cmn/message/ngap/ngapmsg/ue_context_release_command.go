package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"unsafe"
)

// UeContextReleaseCmd struct definition
type UeContextReleaseCmdMsg struct {
	UeNgapIdType types3gpp.UeNgapIdType
	RanNgapId    uint32
	AmfNgapId    uint64

	Cause types3gpp.NgapCause

	ctxt codec.NgapOssCtxt
}

// NewUeContextReleaseCmdMsg create a new Message
func NewUeContextReleaseCmdMsg() *UeContextReleaseCmdMsg {
	return &UeContextReleaseCmdMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *UeContextReleaseCmdMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *UeContextReleaseCmdMsg) String() string {
	return fmt.Sprintf("UeContextReleaseCmd: %s(ran-%d,amf-%d), "+
		"Cause(%s)",
		p.UeNgapIdType, p.RanNgapId, p.AmfNgapId,
		p.Cause)
}

func (p *UeContextReleaseCmdMsg) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)
	rlogger.Trace(types.ModuleCmnNgap, rlogger.INFO, nil, "UeContextReleaseCmdMsg:%s", p.String())

	cmd := codec.NewUeContextReleaseCommandCodec()
	defer codec.DeleteUeContextReleaseCommandCodec(cmd)

	cmd.SetUeNgapIdsType(uint16(p.UeNgapIdType))
	cmd.SetRanUeNgapId(uint(p.RanNgapId))
	cmd.SetAmfUeNgapId(p.AmfNgapId)

	cmd.SetCauseType(uint16(p.Cause.Type))
	cmd.SetCauseValue(uint16(p.Cause.Value))

	msgBuffer := cmd.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer

}

func (p *UeContextReleaseCmdMsg) Decode(msgbuf []byte) error {

	cmd := codec.NewUeContextReleaseCommandCodec()
	defer codec.DeleteUeContextReleaseCommandCodec(cmd)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if cmd.Decode(p.ctxt, msgBuffer) == true {
		p.UeNgapIdType = types3gpp.UeNgapIdType(cmd.GetUeNgapIdsType())
		if p.UeNgapIdType == types3gpp.UeNgApIdPairType {
			p.RanNgapId = uint32(cmd.GetRanUeNgapId())
			p.AmfNgapId = uint64(cmd.GetAmfUeNgapId())
		} else if p.UeNgapIdType == types3gpp.AmfUeNgApIdType {
			p.AmfNgapId = uint64(cmd.GetAmfUeNgapId())
		}

		p.Cause.Type = types3gpp.CauseType(cmd.GetCauseType())
		p.Cause.Value = types3gpp.CauseValue(cmd.GetCauseValue())

	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
