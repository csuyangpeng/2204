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

// UeContextReleaseRequest struct definition
type UeContextReleaseReqMsg struct {
	RanUeNgapId uint32
	AmfUeNgapId uint64

	// 23502 4.2.6
	// If List of PDU Session ID(s) with active N3 user plane is included in step 1b (n2 ue release request),
	// the step 5 to 7 (amf -> smf -> upf :update SM context & n4 modification)
	// are performed before step 2(n2 ue release command)
	PduSessIdList []uint8

	Cause types3gpp.NgapCause
	ctxt  codec.NgapOssCtxt
}

// NewUeContextReleaseReqMsg create a new Message
func NewUeContextReleaseReqMsg() *UeContextReleaseReqMsg {
	return &UeContextReleaseReqMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *UeContextReleaseReqMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *UeContextReleaseReqMsg) String() string {
	return fmt.Sprintf("UeContextReleaseReq:RanNgapId(%d), AmfNgapId(%d), "+
		"Cause(%s), PSI(%v)",
		p.RanUeNgapId, p.AmfUeNgapId, p.Cause, p.PduSessIdList)
}

func (p *UeContextReleaseReqMsg) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	req := codec.NewUeContextReleaseRequestCodec()
	defer codec.DeleteUeContextReleaseRequestCodec(req)

	req.SetRanUeNgapId(uint(p.RanUeNgapId))
	req.SetAmfUeNgapId(p.AmfUeNgapId)

	req.SetCauseType(uint16(p.Cause.Type))
	req.SetCauseValue(uint16(p.Cause.Value))

	if len(p.PduSessIdList) > 0 {
		for _, v := range p.PduSessIdList {
			req.AddPduSessResListRelReq(uint16(v))
		}
	}

	msgBuffer := req.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *UeContextReleaseReqMsg) Decode(msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	req := codec.NewUeContextReleaseRequestCodec()
	defer codec.DeleteUeContextReleaseRequestCodec(req)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgBuf)))
	msgBuffer.SetValue(&msgBuf[0])

	if req.Decode(p.ctxt, msgBuffer) == true {
		p.RanUeNgapId = uint32(req.GetRanUeNgapId())
		p.AmfUeNgapId = req.GetAmfUeNgapId()

		p.Cause.Type = types3gpp.CauseType(req.GetCauseType())
		p.Cause.Value = types3gpp.CauseValue(req.GetCauseValue())

		if req.IsPduSessResListRelReqPrst() {
			for i := 0; i < int(req.GetPduSessResListRelReq().Size()); i++ {
				psi := req.GetPduSessResListRelReq().Get(i)
				p.PduSessIdList = append(p.PduSessIdList, uint8(psi))
			}
		}
	} else {
		return fmt.Errorf("failed to decode msg Bufer")
	}
	return nil
}
