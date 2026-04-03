package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

// PduSessResMdfyRespMsg struct definition
type PduSessResMdfyRespMsg struct {
	RanUeNGAPId uint32
	AmfUeNGAPId uint64
	//opt Ie
	IsUserLocInfoPrst                bool
	UserLocInfo                      types3gpp.UserLocationInfo
	IsPduSessResModifyRespListPrst   bool
	PduSessResModifyRespList         []types3gpp.PduSessResModifyRespItem
	IsPduSessResFailedModifyListPrst bool
	PduSessResFailedModifyList       []types3gpp.PduSessResFailModifyItem

	ctxt codec.NgapOssCtxt
}

//NewPduSessResSetupRespMsg create a new Message
func NewPduSessResMdfyRespMsg() *PduSessResMdfyRespMsg {

	return &PduSessResMdfyRespMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResMdfyRespMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

//AddPduSessResMdfyResp add PduSessResMdfyResp into PduSessResMdfyRespList
func (p *PduSessResMdfyRespMsg) AddPduSessResMdfyResp(item *types3gpp.PduSessResModifyRespItem) {
	p.PduSessResModifyRespList = append(p.PduSessResModifyRespList, *item)
}

//AddPduSessResFailedModify add PduSessResFailedModify into PduSessResFailedModifyList
func (p *PduSessResMdfyRespMsg) AddPduSessResFailedModify(item *types3gpp.PduSessResFailModifyItem) {
	p.PduSessResFailedModifyList = append(p.PduSessResFailedModifyList, *item)
}

func (p *PduSessResMdfyRespMsg) String() string {
	rtStr := fmt.Sprintf("PduSessResMdfyRespMsg:{RanUeNGAPId(%d), AmfUeNGAPId(%d)",
		p.RanUeNGAPId,
		p.AmfUeNGAPId)

	rtStr += fmt.Sprintf(",%v,PduSessResModifyRespList[", p.IsPduSessResModifyRespListPrst)
	for _, v := range p.PduSessResModifyRespList {
		rtStr += fmt.Sprintf("%s,", v.String())
	}
	rtStr += fmt.Sprintf("],%v,PduSessResFailedModifyList[", p.IsPduSessResFailedModifyListPrst)
	for _, v := range p.PduSessResFailedModifyList {
		rtStr += fmt.Sprintf("%s,", v.String())
	}

	return rtStr + "]}"
}

func (p *PduSessResMdfyRespMsg) Encode() []byte {
	modifyRespCodec := codec.NewPduSessResModifyRespCodec()
	defer codec.DeletePduSessResModifyRespCodec(modifyRespCodec)

	modifyRespCodec.SetRanUeNgapId(uint(p.RanUeNGAPId))
	modifyRespCodec.SetAmfUeNgapId(uint64(p.AmfUeNGAPId))

	if p.IsPduSessResModifyRespListPrst {
		for _, mdfyItem := range p.PduSessResModifyRespList {
			mdfyRsp := codec.NewPduSessResModifyRespItem()
			defer codec.DeletePduSessResModifyRespItem(mdfyRsp)
			mdfyRsp.SetPduSessId(mdfyItem.PduSessionId)
			mdfyRsp.SetPduSessResMdfyRespTransferPrst(true)
			mdfyRsp.SetPduSessResMdfyRespTransfer(mdfyItem.PduSessResRelMdfyTransfer)
			modifyRespCodec.AddPduSessResModifyRespList(mdfyRsp)
		}
	}

	if p.IsPduSessResFailedModifyListPrst {
		for _, failItem := range p.PduSessResFailedModifyList {
			failRsp := codec.NewPduSessResFailedMdfyRespItem()
			defer codec.DeletePduSessResFailedMdfyRespItem(failRsp)
			failRsp.SetPduSessId(uint(failItem.PduSessionId))
			failRsp.SetPduSessResFailedMdfyRespTransfer(failItem.PduSessResRelFailMdfyTransfer)
			modifyRespCodec.AddPduSessResFailedModifyList(failRsp)
		}
	}

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
		modifyRespCodec.SetUserLocationInfo(userLocationInfo)
	}

	msgBuffer := modifyRespCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResMdfyRespMsg) Decode(msgbuf []byte) error {

	modifyRespCodec := codec.NewPduSessResModifyRespCodec()
	defer codec.DeletePduSessResModifyRespCodec(modifyRespCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if modifyRespCodec.Decode(p.ctxt, msgBuffer) == true {
		p.AmfUeNGAPId = modifyRespCodec.GetAmfUeNgapId()
		p.RanUeNGAPId = uint32(modifyRespCodec.GetRanUeNgapId())

		if modifyRespCodec.IsPduSessResMdfyRespListPrst() {
			mdfyRespVec := modifyRespCodec.GetPduSessResModifyRespList()
			for i := 0; i < int(mdfyRespVec.Size()); i++ {
				respItem := types3gpp.PduSessResModifyRespItem{}
				respItem.PduSessionId = mdfyRespVec.Get(i).GetPduSessId()
				respItem.PduSessResRelMdfyTransfer =
					mdfyRespVec.Get(i).GetPduSessResMdfyRespTransfer()
				p.AddPduSessResMdfyResp(&respItem)
			}
			p.IsPduSessResModifyRespListPrst = true
		}

		if modifyRespCodec.IsPduSessResFailMdfyRespListPrst() {
			failVec := modifyRespCodec.GetPduSessResFailedModifyList()
			for i := 0; i < int(failVec.Size()); i++ {
				failItem := types3gpp.PduSessResFailModifyItem{}
				failItem.PduSessionId = uint8(failVec.Get(i).GetPduSessId())
				failItem.PduSessResRelFailMdfyTransfer =
					failVec.Get(i).GetPduSessResFailedMdfyRespTransfer()
				p.AddPduSessResFailedModify(&failItem)
			}
			p.IsPduSessResFailedModifyListPrst = true
		}

		if modifyRespCodec.IsUserLocInfoPrst() {
			p.UserLocInfo.NrCgi.PlmnID.SetValue(
				utils.Conv2ByteSlice(modifyRespCodec.GetUserLocationInfo().GetNRCgi().GetPlmn(),
					types3gpp.SizeofPlmnID), types3gpp.BigEndian)

			copy(p.UserLocInfo.NrCgi.NrCellId[:],
				utils.Conv2ByteSlice(modifyRespCodec.GetUserLocationInfo().GetNRCgi().GetNrCellId(),
					types3gpp.SizeofNrCellID))

			if modifyRespCodec.GetUserLocationInfo().GetIsTimeStampPrst() {
				copy(p.UserLocInfo.TimeStmp[:],
					utils.Conv2ByteSlice(modifyRespCodec.GetUserLocationInfo().GetTimeStamp(),
						types3gpp.SizeofTimeStamp))
				p.UserLocInfo.IsTimeStampPrst = true
			}
			if modifyRespCodec.GetUserLocationInfo().GetPsCellInfoPrst() {
				p.UserLocInfo.PsCellInfo.PlmnID.SetValue(
					utils.Conv2ByteSlice(modifyRespCodec.GetUserLocationInfo().GetPsCellInfo().GetPlmn(),
						types3gpp.SizeofPlmnID),
					types3gpp.BigEndian)
				copy(p.UserLocInfo.PsCellInfo.NrCellId[:],
					utils.Conv2ByteSlice(modifyRespCodec.GetUserLocationInfo().GetPsCellInfo().GetNrCellId(),
						types3gpp.SizeofNrCellID))
				p.UserLocInfo.IsPsCellInfoPrst = true
			}

			p.IsUserLocInfoPrst = true
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
