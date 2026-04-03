package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

// PduSessResSetupModMsg struct definition
type PduSessResModifyReqMsg struct {
	RanUeNgApId             uint32
	AmfUeNgApId             uint64
	IsRanPagePriorityPrst   bool
	RanPagePriority         uint16
	PduSessResModifyReqList []types3gpp.PduSessResModifyReqItem

	ctxt codec.NgapOssCtxt
}

// NewPduSessResModifyReqMsg create a new Message
func NewPduSessResModifyReqMsg() *PduSessResModifyReqMsg {
	return &PduSessResModifyReqMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResModifyReqMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// add PduSessResModifyReqItem into PduSessResSetupReqList
func (p *PduSessResModifyReqMsg) AddPduSessResModifyReqItem(
	pSessionResource *types3gpp.PduSessResModifyReqItem) {
	p.PduSessResModifyReqList = append(p.PduSessResModifyReqList, *pSessionResource)
}

func (p *PduSessResModifyReqMsg) String() string {
	rtStr := fmt.Sprintf("PduSessResModifyReqMsg{ RanUeNGAPId(%d),"+
		"AmfUeNGAPId(%d),[%v RanPagingPriority(%d)],",
		p.RanUeNgApId,
		p.AmfUeNgApId,
		p.IsRanPagePriorityPrst,
		p.RanPagePriority)

	for _, v := range p.PduSessResModifyReqList {
		rtStr += fmt.Sprintf("%s,", v.String())
	}
	return rtStr + "}"
}

func (p *PduSessResModifyReqMsg) Encode() []byte {

	modifyReqCodec := codec.NewPduSessResModReqCodec()
	defer codec.DeletePduSessResModReqCodec(modifyReqCodec)

	modifyReqCodec.SetRanUeNgapId(uint(p.RanUeNgApId))
	modifyReqCodec.SetAmfUeNgapId(uint64(p.AmfUeNgApId))

	if p.IsRanPagePriorityPrst {
		modifyReqCodec.SetRanPagingPriority(p.RanPagePriority)
	}

	for _, v := range p.PduSessResModifyReqList {
		pduSessResModReqItem := codec.NewPduSessResModReqItem()
		defer codec.DeletePduSessResModReqItem(pduSessResModReqItem)

		pduSessResModReqItem.SetPduSessId(v.PduSessionId)
		if v.IsNasPduPrst {
			pduSessResModReqItem.SetNasPduPrst(true)
			pduSessResModReqItem.SetNasPdu(v.NasPdu)
		}
		pduSessResModReqItem.SetPduSessResModReqTransfer(v.PduSessResModReqTrans)

		if v.IsSnssaiPrst {
			pduSessResModReqItem.SetSnssaiPrst(true)
			//pduSessResModReqItem.SetSnssai(v.Snssai)
			snnsai := codec.NewSNssai()
			defer codec.DeleteSNssai(snnsai)

			snnsai.SetSst(&v.Snssai.Sst)
			if v.Snssai.SdPrst == true {
				var sd [types3gpp.SizeofSD]byte
				types3gpp.ConvertU32ToSd(sd[:], v.Snssai.Sd, types.BigEndian)
				snnsai.SetSd(&sd[0])
				snnsai.SetSdPresent(true)
			}

			pduSessResModReqItem.SetSnssai(snnsai)
		}

		modifyReqCodec.AddPduSessResModReqList(pduSessResModReqItem)
	}

	msgBuffer := modifyReqCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResModifyReqMsg) Decode(msgbuf []byte) error {

	modifyReqCodec := codec.NewPduSessResModReqCodec()
	defer codec.DeletePduSessResModReqCodec(modifyReqCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if modifyReqCodec.Decode(p.ctxt, msgBuffer) == true {

		//RanUeNgapId
		p.RanUeNgApId = uint32(modifyReqCodec.GetRanUeNgapId())

		//AmfUeNgapId
		p.AmfUeNgApId = modifyReqCodec.GetAmfUeNgapId()

		//RanPagingPriority
		if modifyReqCodec.IsRanPagingPriorityPresent() {
			p.RanPagePriority = uint16(modifyReqCodec.GetRanPagingPriority())
			p.IsRanPagePriorityPrst = true
		}

		//PduSessResModReq
		reqList := modifyReqCodec.GetPduSessResModReqList()
		for i := 0; i < int(reqList.Size()); i++ {
			item := types3gpp.PduSessResModifyReqItem{}
			item.PduSessionId = reqList.Get(i).GetPduSessId()
			item.IsNasPduPrst = reqList.Get(i).GetNasPduPrst()
			if item.IsNasPduPrst {
				item.NasPdu = reqList.Get(i).GetNasPdu()
			}
			item.PduSessResModReqTrans =
				reqList.Get(i).GetPduSessResModReqTransfer()
			//item.Snssai = reqList.Get(i).GetSnssai()
			snssai := reqList.Get(i).GetSnssai()
			sst := snssai.GetSst()
			item.Snssai.Sst = *sst

			if snssai.GetSdPresent() == true {
				var tmp [types3gpp.SizeofSD]byte
				copy(tmp[:],
					utils.Conv2ByteSlice(snssai.GetSd(), types3gpp.SizeofSD))

				item.Snssai.SdPrst = true
				item.Snssai.Sd = types3gpp.ConvertSdToU32(tmp[:], types.BigEndian)
			}

			p.AddPduSessResModifyReqItem(&item)
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
