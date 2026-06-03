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

// NgSetupResponseMsg struct definition
type NgSetupResponseMsg struct {
	//mandatory
	AmfName             string
	RelativeAmfCapacity uint16
	ServedGuamiList     []types3gpp.Guami
	PlmnSupportList     []types3gpp.BPlmn

	ctxt codec.NgapOssCtxt
}

// NewNgSetupResponseMsg create a new Message
func NewNgSetupResponseMsg() *NgSetupResponseMsg {
	return &NgSetupResponseMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *NgSetupResponseMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// AddServedGuami add ServedGUAMI into List
func (p *NgSetupResponseMsg) AddServedGuami(sguami *types3gpp.Guami) {
	p.ServedGuamiList = append(p.ServedGuamiList, *sguami)
}

// AddPlmnSupport add PlmnSupport into List
func (p *NgSetupResponseMsg) AddPlmnSupport(bplmn *types3gpp.BPlmn) {
	p.PlmnSupportList = append(p.PlmnSupportList, *bplmn)
}

func (p *NgSetupResponseMsg) String() string {
	rtStr := fmt.Sprintf("NgSetupResponseMsg: \n"+
		"AmfName(%s), \n RelativeAmfCapacity(%d)\n",
		p.AmfName, p.RelativeAmfCapacity)

	var servedGuamiListStr string
	for i, v := range p.ServedGuamiList {
		servedGuamiListStr += fmt.Sprintf("\n    { #%d %s}", i+1, &v)
	}
	rtStr += "ServedGuamiList{ " + servedGuamiListStr + " }\n"

	var plmnSupportListStr string
	for i, v := range p.PlmnSupportList {
		plmnSupportListStr += fmt.Sprintf("\n    { #%d %s}", i+1, &v)
	}
	rtStr += "PlmnSupportList{ " + plmnSupportListStr + " }"

	return rtStr
}

func (p *NgSetupResponseMsg) Encode() []byte {

	setupRespCodec := codec.NewNgSetupResponseCodec()
	defer codec.DeleteNgSetupResponseCodec(setupRespCodec)

	setupRespCodec.SetAmfName(p.AmfName)
	setupRespCodec.SetRelativeAmfCapacity(p.RelativeAmfCapacity)

	for _, gumai := range p.ServedGuamiList {
		serGuamiItem := codec.NewServedGuamiItem()
		defer codec.DeleteServedGuamiItem(serGuamiItem)

		plmnid := gumai.PlmnId.GetValue(types3gpp.BigEndian)
		serGuamiItem.SetPlmnId(&plmnid[0])

		amfIdentifer := codec.NewAmfIdentifier()
		defer codec.DeleteAmfIdentifier(amfIdentifer)

		//region id
		amfIdentifer.SetRegionId(gumai.AmfId.GetAmfRegionID())
		//set id
		setid := gumai.AmfId.GetAmfSetID()
		amfIdentifer.SetSetId(&setid[0])
		//pointer
		amfIdentifer.SetPointer(gumai.AmfId.GetAmfPointer())

		serGuamiItem.SetAmfId(amfIdentifer)

		setupRespCodec.AddServedGuamiList(serGuamiItem)
	}

	for _, bp := range p.PlmnSupportList {
		bplmn := codec.NewBPlmnItem()
		defer codec.DeleteBPlmnItem(bplmn)

		bplmnid := bp.Plmn.GetValue(types3gpp.BigEndian)
		bplmn.SetPlmnid(&bplmnid[0])

		ssList := codec.NewSNssaiVector()
		defer codec.DeleteSNssaiVector(ssList)

		for _, ss := range bp.SliceSupportList {
			snssai := codec.NewSNssai()
			defer codec.DeleteSNssai(snssai)
			snssai.SetSst(&ss.Sst)
			if ss.SdPrst == true {
				var sd [types3gpp.SizeofSD]byte
				types3gpp.ConvertU32ToSd(sd[:], ss.Sd, types.BigEndian)
				snssai.SetSd(&sd[0])
				snssai.SetSdPresent(true)
			}
			ssList.Add(snssai)
		}

		bplmn.SetSsList(ssList)
		setupRespCodec.AddPlmnList(bplmn)
	}

	//msgBuffer := codec.NewMsgBuffer()
	//defer codec.DeleteMsgBuffer(msgBuffer)
	//rt := setupRespCodec.Encode(p.ctxt, msgBuffer)
	//if rt != true {
	//	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil,  "failed to encode ng setup response message.")
	//	return nil
	//}
	//bufLen := msgBuffer.GetLength()
	//bufValue := msgBuffer.GetValue()
	//rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "encode msg, len(%d),buf(%x)", bufLen,bufValue)
	//fmt.Println("encode msg, ", bufLen,bufValue)
	//
	//encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	//return encodeBuffer

	msgBuf := setupRespCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuf)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *NgSetupResponseMsg) Decode(msgbuf []byte) error {

	setupRespCodec := codec.NewNgSetupResponseCodec()
	defer codec.DeleteNgSetupResponseCodec(setupRespCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if setupRespCodec.Decode(p.ctxt, msgBuffer) == true {
		//AmfName
		p.AmfName = setupRespCodec.GetAmfName()

		//RelativeAmfCapacity
		p.RelativeAmfCapacity = setupRespCodec.GetRelativeAmfCapacity()

		//ServedGuamiList
		guamiListVec := setupRespCodec.GetServedGuamiList()
		for i := 0; i < int(guamiListVec.Size()); i++ {
			gumai := types3gpp.Guami{}

			gumai.PlmnId.SetValue(
				utils.Conv2ByteSlice(guamiListVec.Get(i).GetPlmnId(), types3gpp.SizeofPlmnID),
				types3gpp.BigEndian)
			gumai.AmfId.SetAmfRegionID(guamiListVec.Get(i).GetAmfId().GetRegionId())
			gumai.AmfId.SetAmfPointer(guamiListVec.Get(i).GetAmfId().GetPointer())
			setID := utils.Conv2ByteSlice(guamiListVec.Get(i).GetAmfId().GetSetId(), types3gpp.SizeofAmfSetID)
			gumai.AmfId.SetAmfSetID(setID)

			p.AddServedGuami(&gumai)
		}

		//Plmn Support List
		plmnSupList := setupRespCodec.GetPlmnList()
		for i := 0; i < int(plmnSupList.Size()); i++ {
			bplmnMsg := types3gpp.BPlmn{
				Plmn:             types3gpp.PlmnID{},
				SliceSupportList: nil,
			}

			bplmnListItem := plmnSupList.Get(i)
			bplmnMsg.Plmn.SetValue(utils.Conv2ByteSlice(
				bplmnListItem.GetPlmnid(),
				types3gpp.SizeofPlmnID),
				types3gpp.BigEndian)

			ssList := bplmnListItem.GetSsList()
			for j := 0; j < int(ssList.Size()); j++ {
				snssaiMsg := types3gpp.Snssai{}

				snnsai := ssList.Get(j)
				sst := snnsai.GetSst()
				snssaiMsg.Sst = *sst

				if snnsai.GetSdPresent() == true {
					var tmp [types3gpp.SizeofSD]byte
					copy(tmp[:],
						utils.Conv2ByteSlice(snnsai.GetSd(), types3gpp.SizeofSD))

					snssaiMsg.SdPrst = true
					snssaiMsg.Sd = types3gpp.ConvertSdToU32(tmp[:], types.BigEndian)
				}
				bplmnMsg.AddSnssai(&snssaiMsg)
			}
			p.AddPlmnSupport(&bplmnMsg)
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
