package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

// UeContextReleaseCmp struct definition
type UeContextReleaseCmpMsg struct {
	RanUeNgapId uint32
	AmfUeNgapId uint64

	UserLocInfo    types3gpp.UserLocationInfo
	IsUliPrst      bool
	RcmdPagingInfo types3gpp.RecommendPagingInfo
	IsRpiPrst      bool

	PduSessIdList []byte

	ctxt codec.NgapOssCtxt
}

// NewUeContextReleaseCmpMsg create a new Message
func NewUeContextReleaseCmpMsg() *UeContextReleaseCmpMsg {
	return &UeContextReleaseCmpMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *UeContextReleaseCmpMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *UeContextReleaseCmpMsg) String() string {
	return fmt.Sprintf("UeContxtReleaseComplete Message: "+
		"RanUeNgapId(%d),AmfUeNgapId(%d), UserLocationInfo(%v,%+v),"+
		"RecommanedPagingInfo(%v,%s),PduSessIdList(%v)",
		p.RanUeNgapId, p.AmfUeNgapId, p.IsRpiPrst, p.UserLocInfo,
		p.IsRpiPrst, p.RcmdPagingInfo, p.PduSessIdList)
}

func (p *UeContextReleaseCmpMsg) Encode() []byte {
	cmp := codec.NewUeContextReleaseCompleteCodec()
	defer codec.DeleteUeContextReleaseCompleteCodec(cmp)

	cmp.SetAmfUeNgapId(uint64(p.AmfUeNgapId))
	cmp.SetRanUeNgapId(uint(p.RanUeNgapId))

	if p.IsUliPrst {
		uli := codec.NewUserLocationInfo()
		defer codec.DeleteUserLocationInfo(uli)

		nrcgi := codec.NewNrCgi()
		defer codec.DeleteNrCgi(nrcgi)

		plmnid := p.UserLocInfo.NrCgi.PlmnID.GetValue(types3gpp.BigEndian)
		nrcgi.SetPlmn(&plmnid[0])

		cellId := p.UserLocInfo.NrCgi.NrCellId
		nrcgi.SetNrCellId(&cellId[0])
		uli.SetNRCgi(nrcgi)

		if p.UserLocInfo.IsTimeStampPrst {
			timeStmp := p.UserLocInfo.TimeStmp
			uli.SetTimeStamp(&(timeStmp[0]))
			uli.SetIsTimeStampPrst(true)
		}

		if p.UserLocInfo.IsPsCellInfoPrst {
			uli.SetPsCellInfoPrst(true)
			psCellInfo := codec.NewNrCgi()
			defer codec.DeleteNrCgi(psCellInfo)
			plmn1 := p.UserLocInfo.PsCellInfo.PlmnID.GetValue(types3gpp.BigEndian)
			psCellInfo.SetPlmn(&plmn1[0])
			nrCellId1 := p.UserLocInfo.PsCellInfo.NrCellId
			psCellInfo.SetNrCellId(&nrCellId1[0])
			uli.SetPsCellInfo(psCellInfo)
		}

		tai := codec.NewTaiType()
		defer codec.DeleteTaiType(tai)

		taiPlmnId := p.UserLocInfo.Tai.Plmn.GetValue(types3gpp.BigEndian)
		tai.SetPlmn(&taiPlmnId[0])
		tac := p.UserLocInfo.Tai.Tac.GetByteSlice()
		tai.SetTac(&tac[0])

		uli.SetTai(tai)

		cmp.SetUserLocationInfo(uli)
	}

	if p.IsRpiPrst {
		cellList := codec.NewRecommandCellItemVector()
		defer codec.DeleteRecommandCellItemVector(cellList)
		for _, val := range p.RcmdPagingInfo.CellInfo {
			//encode recommended paging info
			cellItem := codec.NewRecommandCellItem()
			defer codec.DeleteRecommandCellItem(cellItem)

			if val.IsTSPrst {
				cellItem.SetTimeStayedInCell(uint16(val.TimeStayed))
				cellItem.SetTimeStayedPrst(true)
			}

			nrcgi := codec.NewNrCgi()
			defer codec.DeleteNrCgi(nrcgi)

			plmnid := val.Cgi.PlmnID.GetValue(types3gpp.BigEndian)
			nrcgi.SetPlmn(&plmnid[0])
			nrcgi.SetNrCellId(&val.Cgi.NrCellId[0])

			cellItem.SetNrcgi(nrcgi)

			cellList.Add(cellItem)
		}
		if cellList.Size() > 0 {
			cmp.SetRecommandCellList(cellList)
		}

		ranNodeList := codec.NewRecommandRanNodeItemVector()
		defer codec.DeleteRecommandRanNodeItemVector(ranNodeList)
		for _, val := range p.RcmdPagingInfo.RanNodeInfo {
			item := codec.NewRecommandRanNodeItem()
			defer codec.DeleteRecommandRanNodeItem(item)

			item.SetTargetChoice(byte(val.PageTarget))

			switch val.PageTarget {
			case types3gpp.PT_RanNode:
				gGnbid := codec.NewGGnbId()
				defer codec.DeleteGGnbId(gGnbid)

				plmnid := val.RanNodeId.GNBID.Plmn.GetValue(types3gpp.BigEndian)
				gGnbid.SetPlmnid(&plmnid[0])

				gnbid := codec.NewGnbId()
				defer codec.DeleteGnbId(gnbid)
				gnbid.SetVal(&val.RanNodeId.GNBID.GNBID[0])
				gnbid.SetLen(val.RanNodeId.GNBID.GNBLength)

				gGnbid.SetGnbId(gnbid)
				item.SetGgnbId(gGnbid)
			case types3gpp.PT_TAI:
				tai := codec.NewTaiType()
				defer codec.DeleteTaiType(tai)

				taiPlmnId := val.Tai.Plmn.GetValue(types3gpp.BigEndian)
				tai.SetPlmn(&taiPlmnId[0])
				tac := val.Tai.Tac.GetByteSlice()
				tai.SetTac(&tac[0])

				item.SetTai(tai)
			}
			ranNodeList.Add(item)
		}
		cmp.SetRecommandRanNodeList(ranNodeList)
	}

	//if len(p.PduSessIdList) > 0 {
	//	for _, val := range p.PduSessIdList {
	//		cmp.AddPduSessResListRelCpl(uint16(val))
	//	}
	//}

	msgBuffer := cmp.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *UeContextReleaseCmpMsg) Decode(msgbuf []byte) error {
	cmp := codec.NewUeContextReleaseCompleteCodec()
	defer codec.DeleteUeContextReleaseCompleteCodec(cmp)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if cmp.Decode(p.ctxt, msgBuffer) == true {

		p.RanUeNgapId = uint32(cmp.GetRanUeNgapId())
		p.AmfUeNgapId = cmp.GetAmfUeNgapId()

		if cmp.IsUliPrst() {
			p.IsUliPrst = true

			uli := cmp.GetUserLocationInfo()
			p.UserLocInfo.NrCgi.PlmnID.SetValue(
				utils.Conv2ByteSlice(uli.GetNRCgi().GetPlmn(), types3gpp.SizeofPlmnID),
				types3gpp.BigEndian)

			nrCellId := utils.Conv2ByteSlice(
				uli.GetNRCgi().GetNrCellId(), types3gpp.SizeofNrCellID)
			for i, v := range nrCellId {
				p.UserLocInfo.NrCgi.NrCellId[i] = v
			}

			p.UserLocInfo.Tai.Plmn.SetValue(
				utils.Conv2ByteSlice(uli.GetTai().GetPlmn(), types3gpp.SizeofPlmnID),
				types3gpp.BigEndian)

			tac := utils.Conv2ByteSlice(uli.GetTai().GetTac(), types3gpp.SizeofTAC)
			for i, v := range tac {
				p.UserLocInfo.Tai.Tac[i] = v
			}

			if uli.GetIsTimeStampPrst() {
				p.UserLocInfo.IsTimeStampPrst = true
				timeStamp := utils.Conv2ByteSlice(
					uli.GetTimeStamp(), types3gpp.SizeofTimeStamp)
				for i, v := range timeStamp {
					p.UserLocInfo.TimeStmp[i] = v
				}
			}
			if uli.GetPsCellInfoPrst() {
				p.UserLocInfo.PsCellInfo.PlmnID.SetValue(
					utils.Conv2ByteSlice(uli.GetPsCellInfo().GetPlmn(),
						types3gpp.SizeofPlmnID),
					types3gpp.BigEndian)
				copy(p.UserLocInfo.PsCellInfo.NrCellId[:],
					utils.Conv2ByteSlice(uli.GetPsCellInfo().GetNrCellId(),
						types3gpp.SizeofNrCellID))
				p.UserLocInfo.IsPsCellInfoPrst = true
			}
		}

		if cmp.IsRecommandInfoPrst() {
			p.IsRpiPrst = true
			cellsVec := cmp.GetRecommandCellList()
			for i := 0; i < int(cellsVec.Size()); i++ {
				cell := cellsVec.Get(i)
				rmdCell := types3gpp.RecommendCell4Page{}

				if cell.GetTimeStayedPrst() {
					rmdCell.IsTSPrst = true
					rmdCell.TimeStayed = int(cell.GetTimeStayedInCell())
				}

				cgi := cell.GetNrcgi()

				nrCellId := utils.Conv2ByteSlice(
					cgi.GetNrCellId(), types3gpp.SizeofNrCellID)
				for i, v := range nrCellId {
					rmdCell.Cgi.NrCellId[i] = v
				}
				rmdCell.Cgi.PlmnID.SetValue(
					utils.Conv2ByteSlice(cgi.GetPlmn(), types3gpp.SizeofPlmnID),
					types3gpp.BigEndian)

				p.RcmdPagingInfo.CellInfo = append(p.RcmdPagingInfo.CellInfo, rmdCell)
			}

			nodesVec := cmp.GetRecommandRanNodeList()
			for i := 0; i < int(nodesVec.Size()); i++ {
				node := nodesVec.Get(i)
				rcmdNode := types3gpp.RecommendRanNode4Page{}

				target := node.GetTargetChoice()
				switch types3gpp.AmfPagingTarget(target) {
				case types3gpp.PT_RanNode:
					ggnbId := node.GetGgnbId()

					rcmdNode.RanNodeId.GNBID.Plmn.SetValue(
						utils.Conv2ByteSlice(ggnbId.GetPlmnid(), types3gpp.SizeofPlmnID),
						types3gpp.BigEndian)

					rcmdNode.RanNodeId.GNBID.GNBLength =
						ggnbId.GetGnbId().GetLen()

					gnbid := utils.Conv2ByteSlice(
						ggnbId.GetGnbId().GetVal(), types3gpp.SizeofGNBID)
					for i, v := range gnbid {
						rcmdNode.RanNodeId.GNBID.GNBID[i] = v
					}
				case types3gpp.PT_TAI:
					tai := node.GetTai()
					rcmdNode.Tai.Plmn.SetValue(
						utils.Conv2ByteSlice(tai.GetPlmn(), types3gpp.SizeofPlmnID),
						types3gpp.BigEndian)

					tac := utils.Conv2ByteSlice(
						tai.GetTac(), types3gpp.SizeofTac)
					for i, v := range tac {
						rcmdNode.Tai.Tac[i] = v
					}

				default:
					rlogger.Trace(types.ModuleCmnNgap, rlogger.DEBUG, p, "invalid AmfPagingTargetType(%s)",
						types3gpp.AmfPagingTarget(target))
					break
				}
				rcmdNode.PageTarget = types3gpp.AmfPagingTarget(target)
				p.RcmdPagingInfo.RanNodeInfo = append(p.RcmdPagingInfo.RanNodeInfo, rcmdNode)
			}
		}

		//psiVec := cmp.GetPduSessResListRelCpl()
		//for i := 0; i < int(psiVec.Size()); i++ {
		//	psi := psiVec.Get(i)
		//	p.PduSessIdList = append(p.PduSessIdList, byte(psi))
		//}

	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
