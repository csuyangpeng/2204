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

// PduSessResSetupReqMsg struct definition
type PduSessResSetupReqMsg struct {
	RanUeNgApId            uint32
	AmfUeNgApId            uint64
	IsRanPagePriorityPrst  bool
	RanPagePriority        uint16
	IsNasPduPrst           bool
	UeAmbr                 types3gpp.Ambr
	IsUeAmbrPrst           bool
	NasPdu                 string
	PduSessResSetupReqList []*types3gpp.PduSessResSetupReqItem

	ctxt codec.NgapOssCtxt
}

//NewPduSessResSetupReqMsg create a new Message
func NewPduSessResSetupReqMsg() *PduSessResSetupReqMsg {

	return &PduSessResSetupReqMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResSetupReqMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

//AddPduSessResSetupReq add PduSessResSetupReq into PDUSessionSetupList
func (p *PduSessResSetupReqMsg) AddPduSessResSetupReqItem(
	pSessionResource *types3gpp.PduSessResSetupReqItem) {
	p.PduSessResSetupReqList = append(p.PduSessResSetupReqList, pSessionResource)
}

func (p *PduSessResSetupReqMsg) String() string {
	rtStr := fmt.Sprintf("PduSessResSetupReqMsg{RanUeNGAPId(%d),"+
		"AmfUeNGAPId(%d),[%v NasPdu(%s)],[%v RanPagingPriority(%d)],[%v UeAmbr(%d)],",
		p.RanUeNgApId,
		p.AmfUeNgApId,
		p.IsNasPduPrst,
		p.NasPdu,
		p.IsRanPagePriorityPrst,
		p.RanPagePriority,
		p.IsUeAmbrPrst,
		p.UeAmbr)

	for _, v := range p.PduSessResSetupReqList {
		rtStr += fmt.Sprintf("%s,", v)
	}
	return rtStr + "}"
}

func (p *PduSessResSetupReqMsg) Encode() []byte {

	setupReqCodec := codec.NewPduSessResSetupReqCodec()
	defer codec.DeletePduSessResSetupReqCodec(setupReqCodec)

	setupReqCodec.SetRanUeNgapId(uint(p.RanUeNgApId))
	setupReqCodec.SetAmfUeNgapId(uint64(p.AmfUeNgApId))

	if p.IsRanPagePriorityPrst {
		setupReqCodec.SetRanPagingPriority(p.RanPagePriority)
	}

	if p.IsNasPduPrst {
		setupReqCodec.SetNasPdu(p.NasPdu)
	}

	if p.IsUeAmbrPrst {
		ambr := codec.NewUeAmbr()
		defer codec.DeleteUeAmbr(ambr)
		ambr.SetUplink(int64(p.UeAmbr.Uplink))
		ambr.SetDownlink(int64(p.UeAmbr.Downlink))
		setupReqCodec.SetUeAmbr(ambr)
	}

	for _, v := range p.PduSessResSetupReqList {
		pduSessResSetupReqItem := codec.NewPduSessResSetupReqItem()
		defer codec.DeletePduSessResSetupReqItem(pduSessResSetupReqItem)

		pduSessResSetupReqItem.SetPduSessId(v.PduSessionId)
		if v.IsNasPduPrst {
			pduSessResSetupReqItem.SetNasPduPrst(true)
			pduSessResSetupReqItem.SetNasPdu(v.NasPdu)
		}
		pduSessResSetupReqItem.SetPduSessResSetupReqTransfer(v.PduSessResSetupReqTrans)

		snssai := codec.NewSNssai()
		defer codec.DeleteSNssai(snssai)
		snssai.SetSst(&v.Snssai.Sst)
		if v.Snssai.SdPrst == true {
			var sd [types3gpp.SizeofSD]byte
			types3gpp.ConvertU32ToSd(sd[:], v.Snssai.Sd, types.BigEndian)
			snssai.SetSd(&sd[0])
			snssai.SetSdPresent(true)
		}
		pduSessResSetupReqItem.SetSnssai(snssai)

		setupReqCodec.AddPduSessResSetupReqList(pduSessResSetupReqItem)
	}

	msgBuffer := setupReqCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResSetupReqMsg) Decode(msgbuf []byte) error {

	setupReqCodec := codec.NewPduSessResSetupReqCodec()
	defer codec.DeletePduSessResSetupReqCodec(setupReqCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if setupReqCodec.Decode(p.ctxt, msgBuffer) == true {

		//RanUeNgapId
		p.RanUeNgApId = uint32(setupReqCodec.GetRanUeNgapId())

		//AmfUeNgapId
		p.AmfUeNgApId = setupReqCodec.GetAmfUeNgapId()

		//RanPagingPriority
		if setupReqCodec.IsRanPagPriPrst() {
			p.RanPagePriority = uint16(setupReqCodec.GetRanPagingPriority())
			p.IsRanPagePriorityPrst = true
		}

		//NasPdu
		if setupReqCodec.IsNasPduPrst() {
			p.NasPdu = setupReqCodec.GetNasPdu()
			p.IsNasPduPrst = true
		}

		//UE AMBR
		if setupReqCodec.IsUeAMBRPrst() {
			p.UeAmbr.Uplink = uint64(setupReqCodec.GetUeAmbr().GetUplink())
			p.UeAmbr.Downlink = uint64(setupReqCodec.GetUeAmbr().GetDownlink())
			p.IsUeAmbrPrst = true
		}

		//PduSessResSetupReq
		reqList := setupReqCodec.GetPduSessResSetupReqList()
		for i := 0; i < int(reqList.Size()); i++ {
			item := &types3gpp.PduSessResSetupReqItem{}
			item.PduSessionId = reqList.Get(i).GetPduSessId()
			item.IsNasPduPrst = reqList.Get(i).GetNasPduPrst()
			if item.IsNasPduPrst {
				item.NasPdu = reqList.Get(i).GetNasPdu()
			}
			item.PduSessResSetupReqTrans =
				reqList.Get(i).GetPduSessResSetupReqTransfer()

			snssaiMsg := &types3gpp.Snssai{}
			snnsai := reqList.Get(i).GetSnssai()
			sst := snnsai.GetSst()
			snssaiMsg.Sst = *sst
			if snnsai.GetSdPresent() == true {
				var tmp [types3gpp.SizeofSD]byte
				copy(tmp[:],
					utils.Conv2ByteSlice(snnsai.GetSd(), types3gpp.SizeofSD))

				snssaiMsg.SdPrst = true
				snssaiMsg.Sd = types3gpp.ConvertSdToU32(tmp[:], types.BigEndian)
			}
			item.Snssai = snssaiMsg

			p.PduSessResSetupReqList = append(p.PduSessResSetupReqList, item)
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
