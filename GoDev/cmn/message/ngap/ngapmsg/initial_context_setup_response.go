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

// PduSessResSetupRespMsg struct definition
type InitialContextSetupRespMsg struct {
	RanUeNGAPId                 uint32
	AmfUeNGAPId                 uint64
	PduSessResSetupRespList     []*types3gpp.PduSessResSetupRespItem
	PduSessResFailedToSetupList []*types3gpp.PduSessResFailedToSetupItem

	ctxt codec.NgapOssCtxt
}

// NewPduSessResSetupRespMsg create a new Message
func NewInitialContextSetupRespMsg() *InitialContextSetupRespMsg {

	return &InitialContextSetupRespMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *InitialContextSetupRespMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// AddPduSessResSetupRes add PduSessResSetupRes into PduSessResSetupResList
func (p *InitialContextSetupRespMsg) AddPduSessResSetupResp(item *types3gpp.PduSessResSetupRespItem) {
	p.PduSessResSetupRespList = append(p.PduSessResSetupRespList, item)
}

// AddPduSessResFailedToSetup add PduSessResFailedToSetup into PduSessResSetupResList
func (p *InitialContextSetupRespMsg) AddPduSessResFailedToSetup(item *types3gpp.PduSessResFailedToSetupItem) {
	p.PduSessResFailedToSetupList = append(p.PduSessResFailedToSetupList, item)
}

func (p *InitialContextSetupRespMsg) String() string {
	rtStr := fmt.Sprintf("InitialContextSetupRespMsg:{RanUeNGAPId(%d), AmfUeNGAPId(%d)",
		p.RanUeNGAPId,
		p.AmfUeNGAPId)

	rtStr += ",PduSessResSetupRespList["
	for _, v := range p.PduSessResSetupRespList {
		rtStr += fmt.Sprintf("%s,", v)
	}
	rtStr += "],PduSessResFailedSetupList["
	for _, v := range p.PduSessResFailedToSetupList {
		rtStr += fmt.Sprintf("%s,", v)
	}

	return rtStr + "]}"
}

func (p *InitialContextSetupRespMsg) Encode() []byte {
	setupRespCodec := codec.NewInitialContextSetupRespCodec()
	defer codec.DeleteInitialContextSetupRespCodec(setupRespCodec)

	setupRespCodec.SetRanUeNgapId(uint(p.RanUeNGAPId))
	setupRespCodec.SetAmfUeNgapId(uint64(p.AmfUeNGAPId))

	for _, setupItem := range p.PduSessResSetupRespList {
		setupRsp := codec.NewPduSessResSetupRespItem()
		defer codec.DeletePduSessResSetupRespItem(setupRsp)
		setupRsp.SetPduSessId(setupItem.PduSessionId)
		setupRsp.SetPduSessResSetupRespTransfer(setupItem.PduSessResSetupRespTransfer)
		setupRespCodec.AddPduSessResSetupRespList(setupRsp)
	}

	for _, failItem := range p.PduSessResFailedToSetupList {
		failRsp := codec.NewPduSessResFailedSetupItem()
		defer codec.DeletePduSessResFailedSetupItem(failRsp)
		failRsp.SetPduSessId(failItem.PduSessionId)
		failRsp.SetPduSessResSetupUnSussTransfer(failItem.UnsuccessTransfer)
		setupRespCodec.AddPduSessResFailedSetupList(failRsp)
	}

	msgBuffer := setupRespCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *InitialContextSetupRespMsg) Decode(msgbuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	setupRespCodec := codec.NewInitialContextSetupRespCodec()
	defer codec.DeleteInitialContextSetupRespCodec(setupRespCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if setupRespCodec.Decode(p.ctxt, msgBuffer) == true {
		p.AmfUeNGAPId = uint64(setupRespCodec.GetAmfUeNgapId())
		p.RanUeNGAPId = uint32(setupRespCodec.GetRanUeNgapId())

		setupRespVec := setupRespCodec.GetPduSessResSetupRespList()
		for i := 0; i < int(setupRespVec.Size()); i++ {
			respItem := &types3gpp.PduSessResSetupRespItem{}
			respItem.PduSessionId = setupRespVec.Get(i).GetPduSessId()
			respItem.PduSessResSetupRespTransfer =
				setupRespVec.Get(i).GetPduSessResSetupRespTransfer()
			p.PduSessResSetupRespList = append(p.PduSessResSetupRespList, respItem)
		}

		failVec := setupRespCodec.GetPduSessResFailedToSetupList()
		for i := 0; i < int(failVec.Size()); i++ {
			failItem := &types3gpp.PduSessResFailedToSetupItem{}
			failItem.PduSessionId = failVec.Get(i).GetPduSessId()
			failItem.UnsuccessTransfer =
				failVec.Get(i).GetPduSessResSetupUnSussTransfer()
			p.PduSessResFailedToSetupList = append(p.PduSessResFailedToSetupList, failItem)
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	rlogger.Trace(types.ModuleCmnNgap, rlogger.DEBUG, nil, "decode for InitialContextSetupRespMsg:%v", *p)
	return nil
}
