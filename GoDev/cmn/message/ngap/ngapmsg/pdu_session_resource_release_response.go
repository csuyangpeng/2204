package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

// PduSessResSetupRespMsg struct definition
type PduSessResRelRespMsg struct {
	RanUeNGAPId           uint32
	AmfUeNGAPId           uint64
	PduSessResRelRespList []*types3gpp.PduSessResRelRespItem

	//opt Ie
	IsUserLocInfoPrst bool
	UserLocInfo       types3gpp.UserLocationInfo

	ctxt codec.NgapOssCtxt
}

//NewPduSessResRelRespMsg create a new Message
func NewPduSessResRelRespMsg() *PduSessResRelRespMsg {

	return &PduSessResRelRespMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResRelRespMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

//AddPduSessResRelResp add PduSessResRelResp into PduSessResRelRespList
func (p *PduSessResRelRespMsg) AddPduSessResRelRespItem(item *types3gpp.PduSessResRelRespItem) {
	p.PduSessResRelRespList = append(p.PduSessResRelRespList, item)
}

func (p *PduSessResRelRespMsg) String() string {
	rtStr := fmt.Sprintf("PduSessResRelRespMsg:{RanUeNGAPId(%d), AmfUeNGAPId(%d)",
		p.RanUeNGAPId,
		p.AmfUeNGAPId)

	rtStr += ",PduSessResRelRespList["
	for _, v := range p.PduSessResRelRespList {
		rtStr += fmt.Sprintf("%s,", v)
	}

	return rtStr + "]}"
}

func (p *PduSessResRelRespMsg) Encode() []byte {
	relRespCodec := codec.NewPduSessResRelRespCodec()
	defer codec.DeletePduSessResRelRespCodec(relRespCodec)

	relRespCodec.SetRanUeNgapId(uint(p.RanUeNGAPId))
	relRespCodec.SetAmfUeNgapId(uint64(p.AmfUeNGAPId))

	// user location info
	if p.IsUserLocInfoPrst {
		userLocationInfo := codec.NewUserLocationInfo()
		defer codec.DeleteUserLocationInfo(userLocationInfo)

		nrcgi := codec.NewNrCgi()
		defer codec.DeleteNrCgi(nrcgi)
		plmn := p.UserLocInfo.NrCgi.PlmnID.GetValue(types3gpp.BigEndian)
		nrcgi.SetPlmn(&plmn[0])
		nrCellId := p.UserLocInfo.NrCgi.NrCellId
		nrcgi.SetNrCellId(&nrCellId[0])
		userLocationInfo.SetNRCgi(nrcgi)

		plmn = p.UserLocInfo.Tai.Plmn.GetValue(types3gpp.BigEndian)
		tai := codec.NewTaiType()
		defer codec.DeleteTaiType(tai)
		tai.SetPlmn(&plmn[0])
		tac := p.UserLocInfo.Tai.Tac
		tai.SetTac(&tac[0])
		userLocationInfo.SetTai(tai)
		if p.UserLocInfo.IsTimeStampPrst {
			userLocationInfo.SetIsTimeStampPrst(true)
			userLocationInfo.SetTimeStamp(&(p.UserLocInfo.TimeStmp[0]))
		}
		if p.UserLocInfo.IsPsCellInfoPrst {
			userLocationInfo.SetPsCellInfoPrst(true)
			psCellInfo := codec.NewNrCgi()
			defer codec.DeleteNrCgi(psCellInfo)
			plmn1 := p.UserLocInfo.PsCellInfo.PlmnID.GetValue(types3gpp.BigEndian)
			psCellInfo.SetPlmn(&plmn1[0])
			nrCellId1 := p.UserLocInfo.PsCellInfo.NrCellId
			psCellInfo.SetNrCellId(&nrCellId1[0])
			userLocationInfo.SetPsCellInfo(psCellInfo)
		}
		relRespCodec.SetUserLocationInfo(userLocationInfo)
	}

	for _, relRespItem := range p.PduSessResRelRespList {
		relRsp := codec.NewPduSessResRelRespItem()
		defer codec.DeletePduSessResRelRespItem(relRsp)
		relRsp.SetPduSessId(relRespItem.PduSessionId)
		relRsp.SetPduSessResRelRespTransfer(relRespItem.PduSessResRelRespTransfer)
		relRespCodec.AddPduSessResRelRespList(relRsp)
	}

	msgBuffer := relRespCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResRelRespMsg) Decode(msgbuf []byte) error {

	relRespCodec := codec.NewPduSessResRelRespCodec()
	defer codec.DeletePduSessResRelRespCodec(relRespCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if relRespCodec.Decode(p.ctxt, msgBuffer) == true {
		p.AmfUeNGAPId = relRespCodec.GetAmfUeNgapId()
		p.RanUeNGAPId = uint32(relRespCodec.GetRanUeNgapId())

		if relRespCodec.IsUserLocInfoPrst() {
			p.UserLocInfo.NrCgi.PlmnID.SetValue(
				utils.Conv2ByteSlice(relRespCodec.GetUserLocationInfo().GetNRCgi().GetPlmn(),
					types3gpp.SizeofPlmnID), types3gpp.BigEndian)

			copy(p.UserLocInfo.NrCgi.NrCellId[:],
				utils.Conv2ByteSlice(relRespCodec.GetUserLocationInfo().GetNRCgi().GetNrCellId(),
					types3gpp.SizeofNrCellID))

			if relRespCodec.GetUserLocationInfo().GetIsTimeStampPrst() {
				copy(p.UserLocInfo.TimeStmp[:],
					utils.Conv2ByteSlice(relRespCodec.GetUserLocationInfo().GetTimeStamp(),
						types3gpp.SizeofTimeStamp))
				p.UserLocInfo.IsTimeStampPrst = true
			}
			if relRespCodec.GetUserLocationInfo().GetPsCellInfoPrst() {
				p.UserLocInfo.PsCellInfo.PlmnID.SetValue(
					utils.Conv2ByteSlice(relRespCodec.GetUserLocationInfo().GetPsCellInfo().GetPlmn(),
						types3gpp.SizeofPlmnID),
					types3gpp.BigEndian)
				copy(p.UserLocInfo.PsCellInfo.NrCellId[:],
					utils.Conv2ByteSlice(relRespCodec.GetUserLocationInfo().GetPsCellInfo().GetNrCellId(),
						types3gpp.SizeofNrCellID))
				p.UserLocInfo.IsPsCellInfoPrst = true
			}
			p.IsUserLocInfoPrst = true
		}

		relRespVec := relRespCodec.GetPduSessResRelRespList()
		for i := 0; i < int(relRespVec.Size()); i++ {
			resRelRespItem := &types3gpp.PduSessResRelRespItem{}
			resRelRespItem.PduSessionId = relRespVec.Get(i).GetPduSessId()
			resRelRespItem.PduSessResRelRespTransfer =
				relRespVec.Get(i).GetPduSessResRelRespTransfer()
			p.PduSessResRelRespList = append(p.PduSessResRelRespList, resRelRespItem)
		}

	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
