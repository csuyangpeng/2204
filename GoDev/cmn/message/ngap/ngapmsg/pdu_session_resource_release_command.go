package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"unsafe"
)

// PduSessResRelCmmdMsg struct definition
type PduSessResRelCmmdMsg struct {
	RanUeNGAPId           uint32
	AmfUeNGAPId           uint64
	PduSessResRelCmmdList []*types3gpp.PduSessResRelCmmdItem

	//opt Ie
	IsNasPduPrst          bool
	NasPdu                []byte
	IsRanPagePriorityPrst bool
	RanPagePriority       uint16

	ctxt codec.NgapOssCtxt
}

//NewPduSessResRelRespMsg create a new Message
func NewPduSessResRelCmmdMsg() *PduSessResRelCmmdMsg {

	return &PduSessResRelCmmdMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResRelCmmdMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

//AddPduSessResRelCmmd add PduSessResRelCmmd into PduSessResRelCmmdList
func (p *PduSessResRelCmmdMsg) AddPduSessResRelCmmdItem(item *types3gpp.PduSessResRelCmmdItem) {
	p.PduSessResRelCmmdList = append(p.PduSessResRelCmmdList, item)
}

func (p *PduSessResRelCmmdMsg) String() string {
	rtStr := fmt.Sprintf("PduSessResRelCmmdMsg:{RanUeNGAPId(%d), AmfUeNGAPId(%d),Nas[%v NasPdu(%v)],"+
		"[%v RanPagingPriority(%d)],",
		p.RanUeNGAPId,
		p.AmfUeNGAPId,
		p.IsNasPduPrst,
		p.NasPdu,
		p.IsRanPagePriorityPrst,
		p.RanPagePriority)

	rtStr += ",PduSessResRelCmmdList["
	for _, v := range p.PduSessResRelCmmdList {
		rtStr += fmt.Sprintf("%s,", v)
	}

	return rtStr + "]}"
}

func (p *PduSessResRelCmmdMsg) Encode() []byte {
	relCmmdCodec := codec.NewPduSessResRelCmdCodec()
	defer codec.DeletePduSessResRelCmdCodec(relCmmdCodec)

	relCmmdCodec.SetRanUeNgapId(uint(p.RanUeNGAPId))
	relCmmdCodec.SetAmfUeNgapId(uint64(p.AmfUeNGAPId))

	if p.IsRanPagePriorityPrst {
		relCmmdCodec.SetRanPagingPriority(p.RanPagePriority)
	}

	if p.IsNasPduPrst {
		relCmmdCodec.SetNasPdu(string(p.NasPdu))
	}

	for _, relCmmdItem := range p.PduSessResRelCmmdList {
		relCmd := codec.NewPduSessResRelCmdItem()
		defer codec.DeletePduSessResRelCmdItem(relCmd)
		relCmd.SetPduSessId(relCmmdItem.PduSessionId)
		relCmd.SetPduSessResRelCmdTransfer(relCmmdItem.PduSessResRelCmmdTransfer)
		relCmmdCodec.AddPduSessResRelCmdList(relCmd)
	}

	msgBuffer := relCmmdCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResRelCmmdMsg) Decode(msgbuf []byte) error {

	relCmmdCodec := codec.NewPduSessResRelCmdCodec()
	defer codec.DeletePduSessResRelCmdCodec(relCmmdCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if relCmmdCodec.Decode(p.ctxt, msgBuffer) == true {
		p.AmfUeNGAPId = relCmmdCodec.GetAmfUeNgapId()
		p.RanUeNGAPId = uint32(relCmmdCodec.GetRanUeNgapId())

		//NasPdu
		if relCmmdCodec.IsNasPduPrst() {
			p.NasPdu = []byte(relCmmdCodec.GetNasPdu())
			p.IsNasPduPrst = true
		}

		//RanPagingPriority
		if relCmmdCodec.IsRanPagingPriorityPresent() {
			p.RanPagePriority = uint16(relCmmdCodec.GetRanPagingPriority())
			p.IsRanPagePriorityPrst = true
		}

		relCmmdVec := relCmmdCodec.GetPduSessResRelCmdList()
		for i := 0; i < int(relCmmdVec.Size()); i++ {
			resRelCmmdItem := &types3gpp.PduSessResRelCmmdItem{}
			resRelCmmdItem.PduSessionId = relCmmdVec.Get(i).GetPduSessId()
			resRelCmmdItem.PduSessResRelCmmdTransfer =
				relCmmdVec.Get(i).GetPduSessResRelCmdTransfer()
			p.PduSessResRelCmmdList = append(p.PduSessResRelCmmdList, resRelCmmdItem)
		}

	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
