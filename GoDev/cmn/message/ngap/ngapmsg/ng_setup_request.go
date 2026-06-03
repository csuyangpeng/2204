package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types"
	t3gpp "lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

// NgSetupRequestMsg struct definition
type NgSetupRequestMsg struct {
	//mandatory
	GRanNodeID      t3gpp.GlobalRanNodeID
	RanNodeNamePrst bool
	SupportTAs      []t3gpp.SupportedTA
	DefPagingDRX    t3gpp.PagingDRX
	//optional
	RanNodeName string

	ctxt codec.NgapOssCtxt
}

// NewNgSetupRequestMsg create a new Message
func NewNgSetupRequestMsg() *NgSetupRequestMsg {
	return &NgSetupRequestMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *NgSetupRequestMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// AddSupportTA add SupportTA into TAList
func (p *NgSetupRequestMsg) AddSupportTA(sta *t3gpp.SupportedTA) {
	p.SupportTAs = append(p.SupportTAs, *sta)
}

func (p *NgSetupRequestMsg) String() string {
	rtStr := fmt.Sprintf("NgSetupReqeustMsg: "+
		"GlobalRanNodeId(%s), RanNodeNamePrst(%v), "+
		"RanNodeName(%s), %s",
		&p.GRanNodeID.GNBID,
		p.RanNodeNamePrst,
		p.RanNodeName,
		p.DefPagingDRX)

	var sTAStr string
	for i, v := range p.SupportTAs {
		sTAStr += fmt.Sprintf("\n    { #%d %s}", i+1, &v)
	}

	return rtStr + "SupportTAList { " + sTAStr + " }"
}

func (p *NgSetupRequestMsg) Encode() []byte {

	setupReqCodec := codec.NewNgSetupRequestCodec()
	defer codec.DeleteNgSetupRequestCodec(setupReqCodec)

	setupReqCodec.SetPagingDrx(byte(p.DefPagingDRX))
	setupReqCodec.SetRanNodeName(p.RanNodeName)

	ggnbid := codec.NewGGnbId()
	defer codec.DeleteGGnbId(ggnbid)

	plmnid := (p.GRanNodeID.GNBID.Plmn.GetValue(t3gpp.BigEndian))
	ggnbid.SetPlmnid(&plmnid[0])

	gnbid := codec.NewGnbId()
	defer codec.DeleteGnbId(gnbid)

	gnbid.SetLen(p.GRanNodeID.GNBID.GNBLength)
	val := p.GRanNodeID.GNBID.GNBID
	gnbid.SetVal(&val[0])

	ggnbid.SetGnbId(gnbid)
	setupReqCodec.SetGgnbId(ggnbid)

	for _, ta := range p.SupportTAs {
		stai := codec.NewSupTAItem()
		defer codec.DeleteSupTAItem(stai)

		stai.SetTac(&(ta.Tac.GetByteSlice()[0]))

		bplmnvec := codec.NewBPlmnItemVector()
		defer codec.DeleteBPlmnItemVector(bplmnvec)

		for _, bp := range ta.BPlmnList {
			bplmn := codec.NewBPlmnItem()
			defer codec.DeleteBPlmnItem(bplmn)
			bplmnid := bp.Plmn.GetValue(t3gpp.BigEndian)
			bplmn.SetPlmnid(&bplmnid[0])

			ssList := codec.NewSNssaiVector()
			defer codec.DeleteSNssaiVector(ssList)
			for _, ss := range bp.SliceSupportList {
				snssai := codec.NewSNssai()
				defer codec.DeleteSNssai(snssai)
				snssai.SetSst(&ss.Sst)
				if ss.SdPrst == true {
					var sd [t3gpp.SizeofSD]byte
					t3gpp.ConvertU32ToSd(sd[:], ss.Sd, types.BigEndian)
					snssai.SetSd(&sd[0])
					snssai.SetSdPresent(true)
				}
				ssList.Add(snssai)
			}
			bplmn.SetSsList(ssList)
			bplmnvec.Add(bplmn)
		}
		stai.SetBplmnList(bplmnvec)
		setupReqCodec.AddSupTAList(stai)
	}

	msgBuffer := setupReqCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *NgSetupRequestMsg) Decode(msgbuf []byte) error {

	setupReqCodec := codec.NewNgSetupRequestCodec()
	defer codec.DeleteNgSetupRequestCodec(setupReqCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if setupReqCodec.Decode(p.ctxt, msgBuffer) == true {
		//RanNodeName
		p.RanNodeNamePrst = setupReqCodec.GetRNodeNamePrst()
		p.RanNodeName = setupReqCodec.GetRanNodeName()
		//Default Paging DRX

		p.DefPagingDRX = t3gpp.PagingDRX(setupReqCodec.GetPagingDrx())

		//GlobalRanNodeID
		plmn := t3gpp.PlmnID{}
		plmn.SetValue(
			utils.Conv2ByteSlice(
				setupReqCodec.GetGgnbId().GetPlmnid(),
				t3gpp.SizeofPlmnID),
			t3gpp.BigEndian)
		gnbID := t3gpp.GlobalGNBID{Plmn: plmn}
		gnbID.GNBLength = setupReqCodec.GetGgnbId().GetGnbId().GetLen()
		tmpGnbID := utils.Conv2ByteSlice(
			setupReqCodec.GetGgnbId().GetGnbId().GetVal(),
			t3gpp.SizeofGNBID)

		for i, v := range tmpGnbID {
			gnbID.GNBID[i] = v
		}
		p.GRanNodeID.GNBID = gnbID
		p.GRanNodeID.GNBID = gnbID

		//Supported TA List
		taList := setupReqCodec.GetSupTAList()
		for i := 0; i < int(taList.Size()); i++ {
			supTAMsg := t3gpp.SupportedTA{}

			supTAItem := taList.Get(i)
			tac := utils.Conv2ByteSlice(supTAItem.GetTac(), t3gpp.SizeofTAC)
			for i, v := range tac {
				supTAMsg.Tac[i] = v
			}

			bplmnList := supTAItem.GetBplmnList()
			for j := 0; j < int(bplmnList.Size()); j++ {
				bplmnMsg := t3gpp.BPlmn{
					Plmn:             t3gpp.PlmnID{},
					SliceSupportList: nil,
				}

				bplmnListItem := bplmnList.Get(j)
				bplmnMsg.Plmn.SetValue(utils.Conv2ByteSlice(
					bplmnListItem.GetPlmnid(),
					t3gpp.SizeofPlmnID),
					t3gpp.BigEndian)

				ssList := bplmnListItem.GetSsList()
				for k := 0; k < int(ssList.Size()); k++ {
					snssaiMsg := t3gpp.Snssai{}

					snnsai := ssList.Get(k)
					sst := snnsai.GetSst()
					snssaiMsg.Sst = *sst

					if snnsai.GetSdPresent() == true {
						var tmp [t3gpp.SizeofSD]byte
						copy(tmp[:],
							utils.Conv2ByteSlice(snnsai.GetSd(), t3gpp.SizeofSD))

						snssaiMsg.SdPrst = true
						snssaiMsg.Sd = t3gpp.ConvertSdToU32(tmp[:], types.BigEndian)
					}
					bplmnMsg.AddSnssai(&snssaiMsg)
				}
				supTAMsg.AddBPlmn(&bplmnMsg)

			}
			p.AddSupportTA(&supTAMsg)
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
