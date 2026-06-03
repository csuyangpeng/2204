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

	"github.com/willf/bitset"
)

type InitialContextSetupReqMsg struct {
	RanUeNgapId            uint32
	AmfUeNgapId            uint64
	Guami                  types3gpp.Guami
	PduSessResSetupReqList []*types3gpp.PduSessResSetupReqItem
	AllowedNssai           []*types3gpp.Snssai
	UeSecCap               types3gpp.SecurityCapability
	SecKey                 types3gpp.SecurityKey
	// opt ie
	OldAmfName                  string
	UeAmbr                      types3gpp.Ambr
	NasPdu                      []byte
	UeRadioCap                  []byte
	IndexRFSelPriority          uint8
	MaskedImeiSv                [8]byte
	RedirectionVoiceEpsFallback types3gpp.RedirectionForVoiceEPSFallback //37413.9.3.1.116

	//UE Radio Capability for Paging
	UeRadioCap4PagingNR    string //Includes the RRC UERadioPagingInformation message as defined in TS 38.331 [18].
	UeRadioCap4PagingEutra string //Includes the RRC UERadioPagingInformation message as defined in TS 36.331 [21].

	RRCInactiveTransitionReportReq types3gpp.RRCInactTransReportReq //38413 9.3.1.91
	//Core Network Assistance Information

	OptFlags bitset.BitSet

	ctxt codec.NgapOssCtxt
}

const (
	ICSR_AmfName = iota
	ICSR_Ueambr
	ICSR_Nas
	ICSR_UeRadioCap
	ICSR_IndexRfSecPriority
	ICSR_MaskedImeiSv
	ICSR_UeRaidoCap4PagingNR
	ICSR_UeRaidoCap4PagingEutra
	ICSR_RedirectionVoiceEpsFallback
)

//NewInitialContextSetupRequest create a new Message
func NewInitialContextSetupReqMsg() *InitialContextSetupReqMsg {
	return &InitialContextSetupReqMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *InitialContextSetupReqMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

//AddPduSessResSetupReq add PduSessResSetupReq into PDUSessionSetupList
func (p *InitialContextSetupReqMsg) AddPduSessResSetupReqItem(
	pSessionResource *types3gpp.PduSessResSetupReqItem) {
	p.PduSessResSetupReqList = append(p.PduSessResSetupReqList, pSessionResource)
}

// add allowed nssai
func (p *InitialContextSetupReqMsg) AddAllowedNssai(item *types3gpp.Snssai) {
	p.AllowedNssai = append(p.AllowedNssai, item)
}

func (p *InitialContextSetupReqMsg) String() string {
	rtStr := fmt.Sprintf("InitialContextSetupReqMsg{RanUeNGAPId(%d),"+
		"AmfUeNGAPId(%d),OldAmfName(%v %s),UeAMbr[%v %s],Nas[%v NasPdu(%v)],Guami(%s),"+
		"UeSecCapablility(%s), SecurityKey(%v),UeRadioCap(%v %v),IndexRFSlecPriority(%v,%d),"+
		"MaskedImeiSv(%v %v),UeRadioCap4PagingNR(%v,%s),UeRadioCap4PagingEutra(%v,%s)"+
		"RedirectionForVoiceEpsFallback(%v,%v)}",
		p.RanUeNgapId,
		p.AmfUeNgapId,
		p.OptFlags.Test(ICSR_AmfName),
		p.OldAmfName,
		p.OptFlags.Test(ICSR_Ueambr),
		p.UeAmbr,
		p.OptFlags.Test(ICSR_Nas),
		p.NasPdu,
		p.Guami,
		&(p.UeSecCap),
		p.SecKey,
		p.OptFlags.Test(ICSR_UeRadioCap),
		p.UeRadioCap,
		p.OptFlags.Test(ICSR_IndexRfSecPriority),
		p.IndexRFSelPriority,
		p.OptFlags.Test(ICSR_MaskedImeiSv),
		p.MaskedImeiSv,
		p.OptFlags.Test(ICSR_UeRaidoCap4PagingNR),
		p.UeRadioCap4PagingNR,
		p.OptFlags.Test(ICSR_UeRaidoCap4PagingEutra),
		p.UeRadioCap4PagingEutra,
		p.OptFlags.Test(ICSR_RedirectionVoiceEpsFallback),
		p.RedirectionVoiceEpsFallback)

	for _, v := range p.PduSessResSetupReqList {
		rtStr += fmt.Sprintf("%s,", v)
	}
	rtStr += "Allowed Nssai{"
	for _, v := range p.AllowedNssai {
		rtStr += fmt.Sprintf("%s,", v)
	}
	return rtStr + "}}"
}

func (p *InitialContextSetupReqMsg) Encode() []byte {

	setupReqCodec := codec.NewInitialContextSetupRequestCodec()
	defer codec.DeleteInitialContextSetupRequestCodec(setupReqCodec)

	setupReqCodec.SetRanUeNgapId(uint(p.RanUeNgapId))
	setupReqCodec.SetAmfUeNgapId(uint64(p.AmfUeNgapId))

	// encode guami
	guami := codec.NewGuami()
	defer codec.DeleteGuami(guami)

	plmnid := p.Guami.PlmnId.GetValue(types3gpp.BigEndian)
	guami.SetPlmnId(&plmnid[0])

	amfid := codec.NewAmfIdentifier()
	defer codec.DeleteAmfIdentifier(amfid)
	amfid.SetPointer(p.Guami.AmfId.GetAmfPointer())
	amfid.SetRegionId(p.Guami.AmfId.GetAmfRegionID())
	amfSetId := p.Guami.AmfId.GetAmfSetID()
	amfid.SetSetId(&amfSetId[0])
	guami.SetAmfId(amfid)
	setupReqCodec.SetGuami(guami)

	// encode pdu session resource setup request lisg
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
			var tmp [types3gpp.SizeofSD]byte
			types3gpp.ConvertU32ToSd(tmp[:], v.Snssai.Sd, types.BigEndian)
			snssai.SetSd(&tmp[0])
			snssai.SetSdPresent(true)
		}
		pduSessResSetupReqItem.SetSnssai(snssai)

		setupReqCodec.AddPduSessResReqList(pduSessResSetupReqItem)
	}
	// encode Allowed NSSAI
	for _, v := range p.AllowedNssai {
		snssai := codec.NewSNssai()
		defer codec.DeleteSNssai(snssai)
		snssai.SetSst(&v.Sst)
		if v.SdPrst == true {
			var tmp [types3gpp.SizeofSD]byte
			types3gpp.ConvertU32ToSd(tmp[:], v.Sd, types.BigEndian)
			snssai.SetSd(&tmp[0])
			snssai.SetSdPresent(true)
		}
		setupReqCodec.AddAllowedSNssai(snssai)
	}

	// encode ue security capabilities
	secCap := codec.NewUeSecurityCap()
	defer codec.DeleteUeSecurityCap(secCap)

	val := p.UeSecCap.GetNrEncAlgoU16()
	rlogger.Trace(types.ModuleCmnNgap, rlogger.DEBUG, nil, "Encry Alg(%d)", val)
	secCap.SetNrEncryptionAlg(&val[0])

	val = p.UeSecCap.GetNrIntPrctAlgoU16()
	rlogger.Trace(types.ModuleCmnNgap, rlogger.DEBUG, nil, "Int Alg(%d)", val)
	secCap.SetNrIntegrityProAlg(&val[0])

	val = p.UeSecCap.GetEutraEncAlgoU16()
	secCap.SetEutraEncryptionAlg(&val[0])

	val = p.UeSecCap.GetEutraIntPrctAlgoU16()
	secCap.SetEutraIntegrityProAlg(&val[0])

	setupReqCodec.SetUeSecurityCap(secCap)

	// encode security key
	setupReqCodec.SetSecurityKey(string(p.SecKey[:]))

	// optional IEs
	if p.OptFlags.Test(ICSR_AmfName) {
		setupReqCodec.SetOldAmfName(p.OldAmfName)
	}

	if p.OptFlags.Test(ICSR_Ueambr) {
		ambr := codec.NewUeAmbr()
		defer codec.DeleteUeAmbr(ambr)
		ambr.SetUplink(int64(p.UeAmbr.Uplink))
		ambr.SetDownlink(int64(p.UeAmbr.Downlink))
		setupReqCodec.SetUeAmbr(ambr)
	}

	if p.OptFlags.Test(ICSR_UeRadioCap) {
		setupReqCodec.SetUERadioCapability(string(p.UeRadioCap))
	}

	if p.OptFlags.Test(ICSR_IndexRfSecPriority) {
		setupReqCodec.SetIndexToRfsp(int64(p.IndexRFSelPriority))
	}

	if p.OptFlags.Test(ICSR_MaskedImeiSv) {
		setupReqCodec.SetMaskedImeiSv(string(p.MaskedImeiSv[:]))
	}

	if p.OptFlags.Test(ICSR_Nas) {
		setupReqCodec.SetNasPdu(string(p.NasPdu))
	}

	ueRadioCap := codec.NewUeRadioCapForPaging()
	defer codec.DeleteUeRadioCapForPaging(ueRadioCap)
	ueRadioCapPrst := false
	if p.OptFlags.Test(ICSR_UeRaidoCap4PagingNR) {
		ueRadioCap.SetUeRadioCapForPagingNrPrst(true)
		ueRadioCap.SetUeRadioCapForPagingNr(p.UeRadioCap4PagingNR)
		ueRadioCapPrst = true
	}
	if p.OptFlags.Test(ICSR_UeRaidoCap4PagingEutra) {
		ueRadioCap.SetUeRadioCapForPagingEutraPrst(true)
		ueRadioCap.SetUeRadioCapForPagingEutra(p.UeRadioCap4PagingEutra)
		ueRadioCapPrst = true
	}
	if ueRadioCapPrst == true {
		setupReqCodec.SetUeRadioCapForPaging(ueRadioCap)
	}

	if p.OptFlags.Test(ICSR_RedirectionVoiceEpsFallback) {
		setupReqCodec.SetRedirForVoiceEpsFallback(byte(p.RedirectionVoiceEpsFallback))
	}

	msgBuffer := setupReqCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *InitialContextSetupReqMsg) Decode(msgbuf []byte) error {

	setupReqCodec := codec.NewInitialContextSetupRequestCodec()
	defer codec.DeleteInitialContextSetupRequestCodec(setupReqCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if setupReqCodec.Decode(p.ctxt, msgBuffer) == true {

		//RanUeNgapId
		p.RanUeNgapId = uint32(setupReqCodec.GetRanUeNgapId())

		//AmfUeNgapId
		p.AmfUeNgapId = uint64(setupReqCodec.GetAmfUeNgapId())

		//Guami
		guami := setupReqCodec.GetGuami()
		p.Guami.PlmnId.SetValue(utils.Conv2ByteSlice(guami.GetPlmnId(), types3gpp.SizeofPlmnID), types3gpp.BigEndian)
		p.Guami.AmfId.SetAmfPointer(guami.GetAmfId().GetPointer())
		p.Guami.AmfId.SetAmfRegionID(guami.GetAmfId().GetRegionId())
		p.Guami.AmfId.SetAmfSetID(utils.Conv2ByteSlice(guami.GetAmfId().GetSetId(), types3gpp.SizeofAmfSetID))

		//PduSessResSetupReq
		reqList := setupReqCodec.GetPduSessResReqList()
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

		// allowed nssai
		allowedList := setupReqCodec.GetAllowedSNssaiList()
		for i := 0; i < int(allowedList.Size()); i++ {
			item := &types3gpp.Snssai{}
			snssai := allowedList.Get(i)
			item.Sst = *snssai.GetSst()
			if snssai.GetSdPresent() {
				var tmp [types3gpp.SizeofSD]byte
				copy(tmp[:],
					utils.Conv2ByteSlice(snssai.GetSd(), types3gpp.SizeofSD))

				item.SdPrst = true
				item.Sd = types3gpp.ConvertSdToU32(tmp[:], types.BigEndian)
			}
			p.AllowedNssai = append(p.AllowedNssai, item)
		}

		//ue security capbility
		val := make([]byte, 2)
		copy(val, utils.Conv2ByteSlice(setupReqCodec.GetUeSecurityCap().GetNrEncryptionAlg(), 2))
		p.UeSecCap.SetNrEncAlgoU16(val)
		copy(val, utils.Conv2ByteSlice(setupReqCodec.GetUeSecurityCap().GetNrIntegrityProAlg(), 2))
		p.UeSecCap.SetNrIntPrctAlgoU16(val)
		copy(val, utils.Conv2ByteSlice(setupReqCodec.GetUeSecurityCap().GetEutraEncryptionAlg(), 2))
		p.UeSecCap.SetEutraEncAlgoU16(val)
		copy(val, utils.Conv2ByteSlice(setupReqCodec.GetUeSecurityCap().GetEutraIntegrityProAlg(), 2))
		p.UeSecCap.SetEutraIntPrctAlgoU16(val)

		// security key
		copy(p.SecKey[:], []byte(setupReqCodec.GetSecurityKey()))

		// ue radio capability
		if setupReqCodec.IsUeRadioCapPrst() {
			valStr := setupReqCodec.GetUERadioCapability()
			p.UeRadioCap = make([]byte, len(valStr))
			copy(p.UeRadioCap, []byte(valStr))
			p.OptFlags.Set(ICSR_UeRadioCap)
		}
		// old amf
		if setupReqCodec.IsAmfNamePrst() {
			p.OldAmfName = setupReqCodec.GetOldAmfName()
			p.OptFlags.Set(ICSR_AmfName)
		}
		// ue ambr
		if setupReqCodec.IsUeAMBRPrst() {
			p.UeAmbr.Uplink = uint64(setupReqCodec.GetUeAmbr().GetUplink())
			p.UeAmbr.Downlink = uint64(setupReqCodec.GetUeAmbr().GetDownlink())
			p.OptFlags.Set(ICSR_Ueambr)
		}

		// index to RF selection priority
		if setupReqCodec.IsIndexToRFSPPrst() {
			p.IndexRFSelPriority = uint8(setupReqCodec.GetIndexToRfsp())
		}

		// maksed imeisv
		if setupReqCodec.IsMaskedImeiSvPrst() {
			copy(p.MaskedImeiSv[:], setupReqCodec.GetMaskedImeiSv())
			p.OptFlags.Set(ICSR_MaskedImeiSv)
		}
		//NasPdu
		if setupReqCodec.IsNasPduPrst() {
			nasStr := setupReqCodec.GetNasPdu()
			p.NasPdu = make([]byte, len(nasStr))
			copy(p.NasPdu, []byte(nasStr))
			p.OptFlags.Set(ICSR_Nas)
		}

		//Redirection For Voice Eps Fallback
		if setupReqCodec.IsRedirForVoiceEpsFallbackPrst() {
			p.RedirectionVoiceEpsFallback = types3gpp.RedirectionForVoiceEPSFallback(setupReqCodec.GetRedirForVoiceEpsFallback())
			p.OptFlags.Set(ICSR_RedirectionVoiceEpsFallback)
		}

		//UeRadioCapbilityforPaging
		if setupReqCodec.IsUeRadioCapForPagingPrst() {
			ueRaidoCap4Paging := setupReqCodec.GetUeRadioCapForPaging()
			if ueRaidoCap4Paging.GetUeRadioCapForPagingEutraPrst() {
				p.UeRadioCap4PagingEutra = ueRaidoCap4Paging.GetUeRadioCapForPagingEutra()
				p.OptFlags.Set(ICSR_UeRaidoCap4PagingEutra)
			}
			if ueRaidoCap4Paging.GetUeRadioCapForPagingNrPrst() {
				p.UeRadioCap4PagingNR = ueRaidoCap4Paging.GetUeRadioCapForPagingNr()
				p.OptFlags.Set(ICSR_UeRaidoCap4PagingNR)
			}
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
