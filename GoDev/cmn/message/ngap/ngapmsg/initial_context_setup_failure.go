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

// InitialContextSetupFailure struct definition
type InitContextSetupFailMsg struct {
	RanNgapId uint32
	AmfNgapId uint32

	PduSessResFailedToSetupReqList []*types3gpp.PduSessResFailedToSetupItem

	Cause types3gpp.NgapCause
	ctxt  codec.NgapOssCtxt
}

// NewInitContextSetupFailMsg create a new Message
func NewInitContextSetupFailMsg() *InitContextSetupFailMsg {
	return &InitContextSetupFailMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *InitContextSetupFailMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

//AddPduSessResFailToSetup add PduSessResFailToSetup into PDUSessionFailToSetupList
func (p *InitContextSetupFailMsg) AddPduSessResFailToSetupItem(
	pSessionResource *types3gpp.PduSessResFailedToSetupItem) {
	p.PduSessResFailedToSetupReqList = append(p.PduSessResFailedToSetupReqList, pSessionResource)
}

func (p *InitContextSetupFailMsg) String() string {
	rtStr := fmt.Sprintf("InitContextSetupFail:RanNgapId(%d), AmfNgapId(%d),Cause(%d,%d),"+
		"}",
		p.RanNgapId,
		p.AmfNgapId,
		p.Cause.Type,
		p.Cause.Value)

	for _, v := range p.PduSessResFailedToSetupReqList {
		rtStr += fmt.Sprintf("%s,", v)
	}
	return rtStr + "}}"
}

func (p *InitContextSetupFailMsg) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	failSetupCodec := codec.NewInitialContextSetupFailureCodec()
	defer codec.DeleteInitialContextSetupFailureCodec(failSetupCodec)

	failSetupCodec.SetRanUeNgapId(uint(p.RanNgapId))
	failSetupCodec.SetAmfUeNgapId(uint64(p.AmfNgapId))

	failSetupCodec.SetCauseType(uint16(p.Cause.Type))
	failSetupCodec.SetCauseValue(uint16(p.Cause.Value))

	// encode pdu session resource fail to set up list
	for _, v := range p.PduSessResFailedToSetupReqList {
		resFailToSetupItem := codec.NewPduSessResFailedSetupItem()
		defer codec.DeletePduSessResFailedSetupItem(resFailToSetupItem)

		resFailToSetupItem.SetPduSessId(v.PduSessionId)
		resFailToSetupItem.SetPduSessResSetupUnSussTransfer(v.UnsuccessTransfer)

		failSetupCodec.AddPduSessResFailToSetupList(resFailToSetupItem)
	}

	msgBuffer := failSetupCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *InitContextSetupFailMsg) Decode(msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	failSetupCodec := codec.NewInitialContextSetupFailureCodec()
	defer codec.DeleteInitialContextSetupFailureCodec(failSetupCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgBuf)))
	msgBuffer.SetValue(&msgBuf[0])

	if failSetupCodec.Decode(p.ctxt, msgBuffer) == true {
		p.RanNgapId = uint32(failSetupCodec.GetRanUeNgapId())
		p.AmfNgapId = uint32(failSetupCodec.GetAmfUeNgapId())

		p.Cause.Type = types3gpp.CauseType(failSetupCodec.GetCauseType())
		p.Cause.Value = types3gpp.CauseValue(failSetupCodec.GetCauseValue())

		//PduSessResFailToSetup
		resFailToSetupList := failSetupCodec.GetPduSessResFailToSetupList()
		for i := 0; i < int(resFailToSetupList.Size()); i++ {
			item := &types3gpp.PduSessResFailedToSetupItem{}
			item.PduSessionId = resFailToSetupList.Get(i).GetPduSessId()
			item.UnsuccessTransfer =
				resFailToSetupList.Get(i).GetPduSessResSetupUnSussTransfer()

			p.PduSessResFailedToSetupReqList = append(p.PduSessResFailedToSetupReqList, item)
		}
	} else {
		return fmt.Errorf("failed to decode msg Bufer")
	}
	return nil
}
