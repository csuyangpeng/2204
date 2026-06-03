package main

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/message/ngap/ngapmsg"
	T3 "lite5gc/cmn/types3gpp"
	"time"
)

func ngSetupRequestMessageEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewNgSetupRequestMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	plmn := T3.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, T3.LittleEndian)

	gnbID := T3.GlobalGNBID{
		Plmn:      plmn,
		GNBLength: 24,
		GNBID:     [T3.SizeofGNBID]byte{0x12, 0x34, 0x56, 0x78},
	}
	msg.GRanNodeID.GNBID = gnbID

	msg.RanNodeName = "shenzhen01"
	msg.RanNodeNamePrst = true
	msg.DefPagingDRX = T3.PagingDrx_v256

	snssai := T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	bplmn := T3.BPlmn{Plmn: plmn}
	bplmn.AddSnssai(&snssai)
	bplmn.AddSnssai(&snssai)

	supportedTA := T3.SupportedTA{
		Tac: T3.TAC{0x44, 0x55, 0x66}}
	supportedTA.AddBPlmn(&bplmn)

	msg.AddSupportTA(&supportedTA)
	fmt.Println(msg)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)

}

func ngSetupRequestDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 21, 0, 61, 0, 0, 4, 0, 102, 0, 21, 0, 0, 68, 85, 102, 0, 100, 240, 16, 0, 1, 16, 136, 17, 34, 51, 16, 136, 17, 34, 51, 0, 27, 0, 8, 0, 100, 240, 16, 16, 18, 52, 86, 0, 82, 64, 12, 4, 128, 115, 104, 101, 110, 122, 104, 101, 110, 48, 49, 0, 21, 64, 1, 96}
	fmt.Println(encMsg)

	setupRequestMsg := ngapmsg.NewNgSetupRequestMsg()
	setupRequestMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := setupRequestMsg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(setupRequestMsg)
}

func ngSetupResponseMessageEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewNgSetupResponseMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	plmn := T3.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, T3.LittleEndian)

	msg.AmfName = "foxconn"
	msg.RelativeAmfCapacity = 200

	guami := T3.Guami{
		PlmnId: plmn,
	}
	guami.AmfId.SetAmfRegionID(16)
	guami.AmfId.SetAmfPointer(8)
	setID := [2]uint8{1, 2}
	guami.AmfId.SetAmfSetID(setID[:])
	msg.AddServedGuami(&guami)

	snssai := T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	bplmn := T3.BPlmn{Plmn: plmn}
	bplmn.AddSnssai(&snssai)
	bplmn.AddSnssai(&snssai)

	msg.AddPlmnSupport(&bplmn)

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func ngSetupResponseMessageDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 21, 0, 47, 0, 0, 4, 0, 80, 0, 10, 0, 100, 240, 16, 0, 1, 0, 136, 4, 64, 0, 86, 64, 1, 200, 0, 96, 0, 8, 0, 0, 100, 240, 16, 8, 1, 2, 0, 1, 0, 9, 3, 0, 102, 111, 120, 99, 111, 110, 110}
	fmt.Println(encMsg)

	setupResponseMsg := ngapmsg.NewNgSetupResponseMsg()
	setupResponseMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := setupResponseMsg.Decode(encMsg)
	fmt.Println("result", decode)
}

func initialUeMessageEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewInitialUeMessage()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNgapId = 1234
	var nasData string = "123456789987654321"
	msg.NasPdu = []byte(nasData)

	plmn := T3.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, T3.LittleEndian)

	msg.UserLocInfo.NrCgi.PlmnID = plmn
	msg.UserLocInfo.NrCgi.NrCellId = [T3.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x90}
	msg.UserLocInfo.Tai.Plmn = plmn
	msg.UserLocInfo.Tai.Tac = T3.TAC{0x44, 0x55, 0x66}
	msg.UserLocInfo.IsTimeStampPrst = false

	plmn1 := T3.PlmnID{}
	val1 := []byte{0x47, 0x0f, 0x01}
	plmn1.SetValue(val1, T3.LittleEndian)
	msg.UserLocInfo.PsCellInfo.PlmnID = plmn1
	msg.UserLocInfo.PsCellInfo.NrCellId = [T3.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x91}
	msg.UserLocInfo.IsPsCellInfoPrst = true

	msg.IsStmsi5gPrst = true
	msg.Stmsi5G.SetTmsi(12345)
	msg.Stmsi5G.SetAmfSetID([]byte{0x12, 0x34})
	msg.Stmsi5G.SetAmfPointer(15)

	msg.IsAmfSetIdPrst = true
	msg.AmfSetId = 0x1245

	//msg.IsRrcEstbCausePrst = true
	msg.RrcEstablishmentCause = T3.Rrc_estb_cause_mo_Signalling

	msg.IsUeContextRequestPrst = true
	msg.UeContextRequest = 0

	msg.IsAllowedNssaiPrst = true
	snssai1 := T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddAllowedNssai(&snssai1)
	snssai2 := T3.Snssai{
		Sst:    44,
		SdPrst: false,
	}
	msg.AddAllowedNssai(&snssai2)

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func initUeMessageDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 15, 64, 105, 0, 0, 8, 0, 85, 0, 3, 64, 4, 210, 0, 38, 0, 19, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 57, 56, 55, 54, 53, 52, 51, 50, 49, 0, 121, 0, 30, 72, 100, 240, 16, 18, 52, 86, 120, 144, 100, 240, 16, 68, 85, 102, 0, 0, 0, 149, 64, 9, 0, 116, 240, 16, 18, 52, 86, 120, 144, 0, 90, 64, 1, 24, 0, 26, 0, 7, 4, 128, 192, 0, 0, 48, 57, 0, 3, 64, 2, 18, 64, 0, 112, 64, 1, 0, 0, 0, 0, 7, 32, 68, 16, 136, 17, 34, 51}

	fmt.Println(encMsg)

	initUeMessage := ngapmsg.NewInitialUeMessage()
	initUeMessage.SetOssCodecCtxt(ossCtxtPtr)
	decode := initUeMessage.Decode(encMsg)
	fmt.Println("result", decode)
}

func downlinkNasTransportEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewDownlinkNasTransportMessage()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNgapId = 0x1234
	msg.AmfUeNgapId = 0x5678
	var nasData string = "123456789987654321"
	msg.NasPdu = []byte(nasData)

	msg.OldAmfName = "longhua"
	msg.IsOldAmfNamePrst = true

	msg.RanPagingPriority = 250
	msg.IsRanPagingPriorityPrst = true

	msg.IndexToRfsp = 0x12345678
	msg.IsIndexToRfspPrst = true

	msg.UeAmbr.Downlink = 0x123456
	msg.UeAmbr.Uplink = 0x123456
	msg.IsUeAmbrPrst = true

	msg.IsAllowedNssaiPrst = true
	snssai1 := T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddAllowedNssai(&snssai1)
	snssai2 := T3.Snssai{
		Sst:    0x44,
		SdPrst: false,
	}
	msg.AddAllowedNssai(&snssai2)

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)

}

func downlinkNasTransportDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 4, 64, 91, 0, 0, 8, 0, 0, 0, 7, 32, 68, 16, 136, 17, 34, 51, 0, 110, 64, 8, 8, 18, 52, 86, 32, 18, 52, 86, 0, 31, 64, 6, 128, 4, 18, 52, 86, 120, 0, 38, 0, 19, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 57, 56, 55, 54, 53, 52, 51, 50, 49, 0, 83, 64, 1, 249, 0, 48, 0, 9, 3, 0, 108, 111, 110, 103, 104, 117, 97, 0, 10, 0, 3, 32, 86, 120, 0, 85, 0, 3, 64, 18, 52}

	fmt.Println(encMsg)

	downlinTransportMsg := ngapmsg.NewDownlinkNasTransportMessage()
	downlinTransportMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := downlinTransportMsg.Decode(encMsg)
	fmt.Println("result", decode)
}

func uplinkNasTransportEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewUplinkNasTransportMessage()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNgapId = 0x1234
	msg.AmfUeNgapId = 0x5678
	var nasData string = "123456789987654321"
	msg.NasPdu = []byte(nasData)

	plmn := T3.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, T3.LittleEndian)
	msg.UserLocInfo.NrCgi.PlmnID = plmn
	msg.UserLocInfo.NrCgi.NrCellId = [T3.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x90}
	msg.UserLocInfo.Tai.Plmn = plmn
	msg.UserLocInfo.Tai.Tac = T3.TAC{0x44, 0x55, 0x66}
	msg.UserLocInfo.IsTimeStampPrst = false

	plmn1 := T3.PlmnID{}
	val1 := []byte{0x47, 0x0f, 0x01}
	plmn1.SetValue(val1, T3.LittleEndian)
	msg.UserLocInfo.PsCellInfo.PlmnID = plmn1
	msg.UserLocInfo.PsCellInfo.NrCellId = [T3.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x91}
	msg.UserLocInfo.IsPsCellInfoPrst = true

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func uplinkNasTransportDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 46, 64, 74, 0, 0, 4, 0, 10, 0, 3, 32, 86, 120, 0, 85, 0, 3, 64, 18, 52, 0, 38, 0, 19, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 57, 56, 55, 54, 53, 52, 51, 50, 49, 0, 121, 64, 30, 72, 100, 240, 16, 18, 52, 86, 120, 144, 100, 240, 16, 68, 85, 102, 0, 0, 0, 149, 64, 9, 0, 116, 240, 16, 18, 52, 86, 120, 144}

	fmt.Println(encMsg)

	uplinkTransportMsg := ngapmsg.NewUplinkNasTransportMessage()
	uplinkTransportMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := uplinkTransportMsg.Decode(encMsg)
	fmt.Println("result", decode)
}

func pduSessResSetupReqMsgDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 29, 0, 128, 165, 0, 0, 5, 0, 10, 0, 2, 0, 1, 0, 85, 0, 2, 0, 2, 0, 83, 64, 1, 9, 0, 38, 0, 10, 9, 49, 50, 51, 52, 53, 54, 55, 56, 57, 0, 74, 0, 127, 0, 64, 5, 9, 57, 56, 55, 54, 53, 52, 51, 50, 49, 66, 32, 17, 34, 51, 108, 0, 0, 6, 0, 130, 0, 7, 8, 7, 161, 32, 16, 78, 32, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 0, 0, 37, 55, 0, 127, 0, 1, 0, 0, 134, 0, 1, 16, 0, 138, 0, 2, 1, 0, 0, 136, 0, 60, 5, 6, 2, 243, 28, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 8, 84, 130, 184, 112, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 4, 92, 64, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18}

	//cmnNgAp := codec.NewNgApOssCodec()
	//cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))
	//
	//cmnNgAp.DecodeHeader(ossCtxtPtr)
	//
	//fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := ngapmsg.NewPduSessResSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(msg)
}

func pduSessResSetupReqMsgEncode(ossCtxtPtr codec.NgapOssCtxt) {

	msg := ngapmsg.NewPduSessResSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNgApId = 1
	msg.RanPagePriority = 10
	msg.RanUeNgApId = 2
	msg.IsRanPagePriorityPrst = true
	msg.IsNasPduPrst = true
	msg.NasPdu = "123456789"
	//msg.IsUeAmbrPrst = true
	//msg.UeAmbr.Uplink = 5000000
	//msg.UeAmbr.Downlink = 6000000

	item := &T3.PduSessResSetupReqItem{}
	item.PduSessionId = 5
	item.IsNasPduPrst = true
	item.NasPdu = "987654321"
	item.PduSessResSetupReqTrans = string([]byte{0, 0, 6, 0, 130, 0, 7, 8, 7, 161, 32, 16, 78, 32, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 0, 0, 37, 55, 0, 127, 0, 1, 0, 0, 134, 0, 1, 16, 0, 138, 0, 2, 1, 0, 0, 136, 0, 60, 5, 6, 2, 243, 28, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 8, 84, 130, 184, 112, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 4, 92, 64, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18})
	item.Snssai = &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddPduSessResSetupReqItem(item)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func pduSessResSetupRespMsgEncode(ossCtxtPtr codec.NgapOssCtxt) {
	resp := ngapmsg.NewPduSessResSetupRespMsg()
	resp.SetOssCodecCtxt(ossCtxtPtr)

	resp.RanUeNGAPId = 5
	resp.AmfUeNGAPId = 6
	setupItem := &T3.PduSessResSetupRespItem{}
	setupItem.PduSessionId = 5
	setupItem.PduSessResSetupRespTransfer = string([]byte{48, 3, 224, 10, 180, 9, 242, 0, 0, 37, 55, 12, 4, 3, 2, 1, 64, 4, 14, 144, 24, 128})
	resp.AddPduSessResSetupResp(setupItem)

	//failItem := &T3.PduSessResFailedToSetupItem{}
	//failItem.PduSessionId = 6
	//failItem.UnsuccessTransfer = string([]byte{0,128})
	//resp.AddPduSessResFailedToSetup(failItem)

	encodeMsg := resp.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func pduSessResSetupRespMsgDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 29, 0, 55, 0, 0, 4, 0, 58, 64, 6, 0, 0, 6, 2, 0, 128, 0, 75, 64, 26, 0, 0, 5, 22, 48, 3, 224, 10, 180, 9, 242, 0, 0, 37, 55, 12, 4, 3, 2, 1, 64, 4, 14, 144, 24, 128, 0, 10, 64, 2, 0, 6, 0, 85, 64, 2, 0, 5}

	fmt.Println(encMsg)

	resp := ngapmsg.NewPduSessResSetupRespMsg()
	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)

}

func initialContextSetupRequestEncode(ossCtxtPtr codec.NgapOssCtxt) {

	msg := ngapmsg.NewInitialContextSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	msg.AmfUeNgapId = 5
	msg.RanUeNgapId = 6
	msg.OldAmfName = "shenzhen.amf"
	msg.OptFlags.Set(ngapmsg.ICSR_AmfName)
	msg.UeAmbr.Uplink = 5000000
	msg.UeAmbr.Downlink = 6000000
	msg.OptFlags.Set(ngapmsg.ICSR_Ueambr)
	msg.Guami.PlmnId.SetString("460001")
	msg.Guami.AmfId.SetAmfRegionID(1)
	msg.Guami.AmfId.SetAmfPointer(2)
	msg.Guami.AmfId.SetAmfSetIdU16(1234)

	item := &T3.PduSessResSetupReqItem{}
	item.PduSessionId = 5
	item.IsNasPduPrst = true
	item.NasPdu = "987654321"
	item.PduSessResSetupReqTrans = string([]byte{0, 0, 7, 0, 136, 0, 72, 5, 6, 2, 243, 28, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18, 0, 129, 0, 2, 0, 11, 0, 138, 0, 2, 1, 0, 0, 134, 0, 1, 16, 0, 127, 0, 1, 0, 0, 139, 0, 22, 7, 240, 10, 180, 9, 242, 55, 37, 0, 0, 48, 244, 13, 1, 0, 0, 0, 0, 55, 37, 0, 0, 0, 130, 0, 7, 8, 7, 161, 32, 16, 78, 32})
	item.Snssai = &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddPduSessResSetupReqItem(item)

	snssai := &T3.Snssai{
		Sst:    1,
		SdPrst: true,
		Sd:     123,
	}
	msg.AddAllowedNssai(snssai)

	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA0)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA1)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA2)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA3)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA0)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA1)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA2)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA3)

	msg.SecKey = T3.SecurityKey{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	msg.UeRadioCap = []byte{1, 2, 3, 4, 5, 67}
	msg.OptFlags.Set(ngapmsg.ICSR_UeRadioCap)
	msg.IndexRFSelPriority = 1
	msg.OptFlags.Set(ngapmsg.ICSR_IndexRfSecPriority)
	msg.MaskedImeiSv = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	msg.OptFlags.Set(ngapmsg.ICSR_MaskedImeiSv)
	msg.NasPdu = []byte{1, 1, 1, 1, 1, 2, 2, 2, 22, 3, 3, 3, 3, 3, 3}
	msg.OptFlags.Set(ngapmsg.ICSR_Nas)
	msg.UeRadioCap4PagingNR = "aaaaaa"
	msg.OptFlags.Set(ngapmsg.ICSR_UeRaidoCap4PagingNR)
	msg.UeRadioCap4PagingEutra = "bbbbb"
	msg.OptFlags.Set(ngapmsg.ICSR_UeRaidoCap4PagingEutra)

	msg.OptFlags.Set(ngapmsg.ICSR_RedirectionVoiceEpsFallback)
	msg.RedirectionVoiceEpsFallback = T3.NotPossible

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func initialContextSetupRequestDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 14, 0, 129, 93, 0, 0, 15, 0, 10, 0, 2, 0, 5, 0, 85, 0, 2, 0, 6, 0, 48, 0, 14, 5, 128, 115, 104, 101, 110, 122, 104, 101, 110, 46, 97, 109, 102, 0, 110, 0, 8, 8, 91, 141, 128, 32, 76, 75, 64, 0, 28, 0, 7, 0, 100, 16, 0, 1, 4, 130, 0, 71, 0, 128, 158, 0, 64, 5, 9, 57, 56, 55, 54, 53, 52, 51, 50, 49, 66, 32, 17, 34, 51, 128, 138, 0, 0, 7, 0, 136, 0, 72, 5, 6, 2, 243, 28, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18, 0, 129, 0, 2, 0, 11, 0, 138, 0, 2, 1, 0, 0, 134, 0, 1, 16, 0, 127, 0, 1, 0, 0, 139, 0, 22, 7, 240, 10, 180, 9, 242, 55, 37, 0, 0, 48, 244, 13, 1, 0, 0, 0, 0, 55, 37, 0, 0, 0, 130, 0, 7, 8, 7, 161, 32, 16, 78, 32, 0, 0, 0, 5, 2, 1, 1, 2, 3, 0, 119, 0, 9, 28, 0, 14, 0, 0, 0, 0, 0, 0, 0, 94, 0, 32, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 117, 64, 7, 6, 1, 2, 3, 4, 5, 67, 0, 31, 64, 2, 0, 0, 0, 34, 64, 8, 1, 2, 3, 4, 5, 6, 7, 8, 0, 38, 64, 16, 15, 1, 1, 1, 1, 1, 2, 2, 2, 22, 3, 3, 3, 3, 3, 3, 0, 118, 64, 14, 96, 6, 97, 97, 97, 97, 97, 97, 5, 98, 98, 98, 98, 98, 0, 146, 64, 1, 64}

	fmt.Println(encMsg)

	resp := ngapmsg.NewInitialContextSetupReqMsg()
	resp.SetOssCodecCtxt(ossCtxtPtr)
	resp.Decode(encMsg)
	fmt.Println("decode message: ", resp)
}
func initialContextSetupResponseEncode(ossCtxtPtr codec.NgapOssCtxt) {
	resp := ngapmsg.NewInitialContextSetupRespMsg()
	resp.SetOssCodecCtxt(ossCtxtPtr)

	resp.RanUeNGAPId = 5
	resp.AmfUeNGAPId = 6
	setupItem := &T3.PduSessResSetupRespItem{}
	setupItem.PduSessionId = 5
	setupItem.PduSessResSetupRespTransfer = string([]byte{48, 15, 224, 10, 180, 9, 242, 0, 0, 37, 55, 208, 24, 2, 0, 57, 127, 0, 0, 0, 0, 37, 55, 5, 6, 0, 80, 1, 3, 164, 6, 32})
	resp.AddPduSessResSetupResp(setupItem)

	//failItem := &T3.PduSessResFailedToSetupItem{}
	//failItem.PduSessionId = 6
	//failItem.UnsuccessTransfer = string([]byte{0,128})
	//resp.AddPduSessResFailedToSetup(failItem)

	encodeMsg := resp.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func initialContextSetupResponseDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 14, 0, 55, 0, 0, 3, 0, 72, 64, 36, 0, 0, 5, 32, 48, 15, 224, 10, 180, 9, 242, 0, 0, 37, 55, 208, 24, 2, 0, 57, 127, 0, 0, 0, 0, 37, 55, 5, 6, 0, 80, 1, 3, 164, 6, 32, 0, 10, 64, 2, 0, 6, 0, 85, 64, 2, 0, 5}

	fmt.Println(encMsg)

	resp := ngapmsg.NewInitialContextSetupRespMsg()
	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
}

func initialContextSetupFailEncode(ossCtxtPtr codec.NgapOssCtxt) {
	failSetupContext := ngapmsg.NewInitContextSetupFailMsg()
	failSetupContext.SetOssCodecCtxt(ossCtxtPtr)

	failSetupContext.AmfNgapId = 5
	failSetupContext.RanNgapId = 6
	failSetupContext.Cause.Type = 1
	failSetupContext.Cause.Value = 2
	failsetupItem := &T3.PduSessResFailedToSetupItem{}
	failsetupItem.PduSessionId = 5
	failsetupItem.UnsuccessTransfer = string([]byte{0, 32})
	failSetupContext.AddPduSessResFailToSetupItem(failsetupItem)

	fmt.Println(failSetupContext)
	encodeMsg := failSetupContext.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func initialContextSetupFailDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{64, 14, 0, 31, 0, 0, 4, 0, 10, 64, 2, 0, 5, 0, 85, 64, 2, 0, 6, 0, 132, 64, 6, 0, 0, 5, 2, 0, 32, 0, 15, 64, 2, 0, 128}

	fmt.Println(encMsg)

	resp := ngapmsg.NewInitContextSetupFailMsg()
	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

func pduSessResRelRespEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewPduSessResRelRespMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNGAPId = 5
	msg.AmfUeNGAPId = 6
	relItem := &T3.PduSessResRelRespItem{}
	relItem.PduSessionId = 5
	relItem.PduSessResRelRespTransfer = string([]byte{0})
	msg.AddPduSessResRelRespItem(relItem)

	plmn := T3.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, T3.LittleEndian)

	msg.UserLocInfo.NrCgi.PlmnID = plmn
	msg.UserLocInfo.NrCgi.NrCellId = [T3.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x90}
	msg.UserLocInfo.Tai.Plmn = plmn
	msg.UserLocInfo.Tai.Tac = T3.TAC{0x44, 0x55, 0x66}
	msg.UserLocInfo.IsTimeStampPrst = false

	plmn1 := T3.PlmnID{}
	val1 := []byte{0x47, 0x0f, 0x01}
	plmn1.SetValue(val1, T3.LittleEndian)
	msg.UserLocInfo.PsCellInfo.PlmnID = plmn1
	msg.UserLocInfo.PsCellInfo.NrCellId = [T3.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x91}
	msg.UserLocInfo.IsPsCellInfoPrst = true
	msg.IsUserLocInfoPrst = true

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func pagingEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewPagingMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.UePagingId.SetTmsi(12345)
	msg.UePagingId.SetAmfSetID([]byte{0x12, 0x34})
	msg.UePagingId.SetAmfPointer(15)

	msg.PagingDrx = 3
	msg.OptFlags.Set(ngapmsg.ICSR_PagingDrx)

	msg.PagingPriority = 1
	msg.OptFlags.Set(ngapmsg.ICSR_PagingPriority)

	msg.PagingOrigin = 0
	msg.OptFlags.Set(ngapmsg.ICSR_PagingOrigin)

	msg.OptFlags.Set(ngapmsg.ICSR_UeRaidoCapForPaging)
	msg.OptFlags.Set(ngapmsg.ICSR_UeRaidoCapPagingNR)
	msg.UeRadioCap4PagingNR = "aaaaa"

	msg.OptFlags.Set(ngapmsg.ICSR_UeRaidoCapPagingEutra)
	msg.UeRadioCap4PagingEutra = "bbb"

	item := T3.TAI{}
	item.Tac.SetTac(512)
	item.Plmn.SetString("46000")
	msg.AddTaiPagingList(&item)

	fmt.Println(msg)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func pduSessResRelRespDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 28, 0, 58, 0, 0, 4, 0, 10, 64, 2, 0, 6, 0, 85, 64, 2, 0, 5, 0, 70, 64, 5, 0, 0, 5, 1, 0, 0, 121, 64, 30, 72, 100, 240, 16, 18, 52, 86, 120, 144, 100, 240, 16, 68, 85, 102, 0, 0, 0, 149, 64, 9, 0, 116, 240, 16, 18, 52, 86, 120, 144}

	fmt.Println(encMsg)

	resp := ngapmsg.NewPduSessResRelRespMsg()

	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)

}

func pagingDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 24, 64, 55, 0, 0, 6, 0, 51, 64, 1, 0, 0, 118, 64, 11, 96, 5, 97, 97, 97, 97, 97, 3, 98, 98, 98, 0, 52, 64, 1, 16, 0, 103, 64, 7, 0, 100, 240, 0, 0, 2, 0, 0, 50, 64, 1, 96, 0, 115, 64, 7, 2, 65, 224, 0, 0, 48, 57}

	fmt.Println(encMsg)

	resp := ngapmsg.NewPagingMsg()

	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

func pduSessResRelCmmdEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewPduSessResRelCmmdMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNGAPId = 5
	msg.AmfUeNGAPId = 6
	relItem := &T3.PduSessResRelCmmdItem{}
	relItem.PduSessionId = 5
	relItem.PduSessResRelCmmdTransfer = string([]byte{0, 32})
	msg.AddPduSessResRelCmmdItem(relItem)

	msg.RanPagePriority = 10
	msg.IsRanPagePriorityPrst = true
	msg.IsNasPduPrst = true
	msg.NasPdu = []byte("123456789")

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func pduSessResRelCmmdDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 28, 0, 44, 0, 0, 5, 0, 79, 0, 6, 0, 0, 5, 2, 0, 32, 0, 83, 64, 1, 9, 0, 38, 64, 10, 9, 49, 50, 51, 52, 53, 54, 55, 56, 57, 0, 10, 0, 2, 0, 6, 0, 85, 0, 2, 0, 5}

	fmt.Println(encMsg)

	resp := ngapmsg.NewPduSessResRelCmmdMsg()

	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

func ngResetEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewNgResetMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.CauseType = 1
	msg.CauseValue = 2

	msg.ResetTypeChoice.IsNgInterfacePrst = true
	msg.ResetTypeChoice.NgInterface = 0

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func ngResetDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 20, 0, 14, 0, 0, 2, 0, 15, 64, 2, 0, 128, 0, 88, 0, 1, 0}

	fmt.Println(encMsg)

	resp := ngapmsg.NewNgResetMsg()

	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

func ngResetAckEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewNgResetAckMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.OptFlags.Set(ngapmsg.ICSR_UeAssLogicalNgConn)
	item := &T3.UeAssLogicalNgConnItem{}
	item.IsAmfUeNgapIdPrst = true
	item.AmfUeNgapId = 1
	item.IsRanUeNgapIdPrst = true
	item.RanUeNgapId = 2

	msg.AddUeAssLogicalNgConnItem(item)

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func ngResetAckDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 20, 0, 12, 0, 0, 1, 0, 111, 64, 5, 1, 96, 1, 0, 2}

	fmt.Println(encMsg)

	resp := ngapmsg.NewNgResetAckMsg()

	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

func errorIndicationEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewErrIndicationMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.IsAmfUeNgapIdPrst = true
	msg.AmfUeNgapId = 1
	msg.IsRanUeNgapIdPrst = true
	msg.RanUeNgapId = 2
	msg.IsCausePrst = true
	msg.CauseType = 1
	msg.CauseValue = 2

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func errorIndicationDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 9, 64, 21, 0, 0, 3, 0, 10, 64, 2, 0, 1, 0, 85, 64, 2, 0, 2, 0, 15, 64, 2, 0, 128}

	fmt.Println(encMsg)

	resp := ngapmsg.NewErrIndicationMsg()

	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

//func snssaiPdusEncode(ossCtxtPtr codec.NgapOssCtxt) {
//	msg := ngapmsg.NewSnssaiPdus()
//	msg.SetOssCodecCtxt(ossCtxtPtr)
//
//	msg.Snssai = &T3.Snssai{
//		Sst:    T3.SST{0x11},
//		Sd:     T3.SD{0x11, 0x22, 0x33},
//		SdPrst: true,
//	}
//
//	fmt.Println(msg)
//	encodeMsg := msg.Encode()
//	fmt.Println("Encode Msg: ", encodeMsg)
//}
//
//func snssaiPdusDecode(ossCtxtPtr codec.NgapOssCtxt) {
//	encMsg := []byte{66,32,17,34,51}
//
//	fmt.Println(encMsg)
//
//	resp := ngapmsg.NewSnssaiPdus()
//
//	resp.SetOssCodecCtxt(ossCtxtPtr)
//	decode := resp.Decode(encMsg)
//	fmt.Println("result", decode)
//	fmt.Println(resp)
//}

func pduSessResMdfyReqEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewPduSessResModifyReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNgApId = 1
	msg.RanUeNgApId = 2

	mdfyReqItem := &T3.PduSessResModifyReqItem{}
	mdfyReqItem.PduSessionId = 5
	mdfyReqItem.PduSessResModReqTrans = string([]byte{0, 0, 1, 0, 166, 64, 10, 9, 99, 109, 110, 101, 116, 46, 99, 111, 109})

	mdfyReqItem.IsSnssaiPrst = true

	mdfyReqItem.Snssai.Sst = 0x33
	mdfyReqItem.Snssai.SdPrst = true
	mdfyReqItem.Snssai.Sd = 0x112233

	//mdfyReqItem.Snssai = string([]byte{66,32,17,34,51})

	msg.AddPduSessResModifyReqItem(mdfyReqItem)
	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func pduSessResMdfyReqDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 26, 0, 51, 0, 0, 3, 0, 10, 0, 2, 0, 1, 0, 85, 0, 2, 0, 2, 0, 64, 0, 32, 0, 32, 5, 17, 0, 0, 1, 0, 166, 64, 10, 9, 99, 109, 110, 101, 116, 46, 99, 111, 109, 0, 0, 0, 148, 0, 5, 70, 96, 17, 34, 51}

	fmt.Println(encMsg)

	resp := ngapmsg.NewPduSessResModifyReqMsg()
	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

func pduSessResMdfyRespEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := ngapmsg.NewPduSessResMdfyRespMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNGAPId = 1
	msg.RanUeNGAPId = 2

	mdfyRespItem := T3.PduSessResModifyRespItem{}
	mdfyRespItem.PduSessionId = 10
	mdfyRespItem.PduSessResRelMdfyTransfer = string([]byte{124, 3, 224, 10, 180, 9, 242, 0, 0, 37, 52, 1, 240, 10, 180, 9, 243, 0, 0, 37, 53, 0, 40, 0, 15, 128, 14, 15, 16, 17, 0, 0, 32, 177, 1, 6, 0, 2, 128, 0})
	msg.AddPduSessResMdfyResp(&mdfyRespItem)
	msg.IsPduSessResModifyRespListPrst = true

	failMdfyRespItem := T3.PduSessResFailModifyItem{}
	failMdfyRespItem.PduSessionId = 15
	failMdfyRespItem.PduSessResRelFailMdfyTransfer = string([]byte{0, 8})
	msg.AddPduSessResFailedModify(&failMdfyRespItem)
	msg.IsPduSessResFailedModifyListPrst = true

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func pduSessResMdfyRespDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 26, 0, 73, 0, 0, 4, 0, 10, 64, 2, 0, 1, 0, 85, 64, 2, 0, 2, 0, 65, 64, 44, 0, 0, 10, 40, 124, 3, 224, 10, 180, 9, 242, 0, 0, 37, 52, 1, 240, 10, 180, 9, 243, 0, 0, 37, 53, 0, 40, 0, 15, 128, 14, 15, 16, 17, 0, 0, 32, 177, 1, 6, 0, 2, 128, 0, 0, 54, 64, 6, 0, 0, 15, 2, 0, 8}

	fmt.Println(encMsg)

	resp := ngapmsg.NewPduSessResMdfyRespMsg()

	resp.SetOssCodecCtxt(ossCtxtPtr)
	decode := resp.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(resp)
}

func run() {
	ossCtxt := codec.NewOssCtxt()

	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	//ngSetupRequestMessageEncode(ossCtxtPtr)
	//ngSetupRequestDecode(ossCtxtPtr)
	//
	//ngSetupResponseMessageEncode(ossCtxtPtr)
	//ngSetupResponseMessageDecode(ossCtxtPtr)
	//
	//initialUeMessageEncode(ossCtxtPtr)
	//initUeMessageDecode(ossCtxtPtr)
	//
	//downlinkNasTransportEncode(ossCtxtPtr)
	//downlinkNasTransportDecode(ossCtxtPtr)
	//
	//uplinkNasTransportEncode(ossCtxtPtr)
	//uplinkNasTransportDecode(ossCtxtPtr)
	//
	//pduSessResSetupReqMsgEncode(ossCtxtPtr)
	//pduSessResSetupReqMsgDecode(ossCtxtPtr)
	//
	//pduSessResSetupRespMsgEncode(ossCtxtPtr)
	//pduSessResSetupRespMsgDecode(ossCtxtPtr)
	//
	//initialContextSetupRequestEncode(ossCtxtPtr)
	//initialContextSetupRequestDecode(ossCtxtPtr)
	//
	//initialContextSetupResponseEncode(ossCtxtPtr)
	//initialContextSetupResponseDecode(ossCtxtPtr)
	//
	//initialContextSetupFailEncode(ossCtxtPtr)
	//initialContextSetupFailDecode(ossCtxtPtr)
	//
	//pduSessResRelRespEncode(ossCtxtPtr)
	//pduSessResRelRespDecode(ossCtxtPtr)
	//
	//pduSessResRelCmmdEncode(ossCtxtPtr)
	//pduSessResRelCmmdDecode(ossCtxtPtr)
	//
	//ngResetEncode(ossCtxtPtr)
	//ngResetDecode(ossCtxtPtr)
	//
	//ngResetAckEncode(ossCtxtPtr)
	//ngResetAckDecode(ossCtxtPtr)
	//
	//errorIndicationEncode(ossCtxtPtr)
	//errorIndicationDecode(ossCtxtPtr)
	//
	//snssaiPdusEncode(ossCtxtPtr)
	//snssaiPdusDecode(ossCtxtPtr)
	//
	//pduSessResMdfyReqEncode(ossCtxtPtr)
	//pduSessResMdfyReqDecode(ossCtxtPtr)

	pduSessResMdfyRespEncode(ossCtxtPtr)
	pduSessResMdfyRespDecode(ossCtxtPtr)
}
func TestInitialContextSetupReqMsg3_Encode() {

	ossCtxt := codec.NewOssCtxt()

	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := ngapmsg.NewInitialContextSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	msg.AmfUeNgapId = 5
	msg.RanUeNgapId = 6
	msg.OldAmfName = "shenzhen.amf"
	msg.OptFlags.Set(ngapmsg.ICSR_AmfName)
	msg.UeAmbr.Uplink = 5000000
	msg.UeAmbr.Downlink = 6000000
	msg.OptFlags.Set(ngapmsg.ICSR_Ueambr)
	msg.Guami.PlmnId.SetString("460001")
	msg.Guami.AmfId.SetAmfRegionID(1)
	msg.Guami.AmfId.SetAmfPointer(2)
	msg.Guami.AmfId.SetAmfSetID([]byte{1, 1})

	item := &T3.PduSessResSetupReqItem{}
	item.PduSessionId = 5
	item.IsNasPduPrst = true
	item.NasPdu = "987654321"
	item.PduSessResSetupReqTrans = string([]byte{88, 32, 7, 161, 32, 16, 78, 32, 1, 240, 10, 180, 9, 242, 55, 37, 0, 0, 0, 8, 20, 24, 11, 204, 112, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18})
	item.Snssai = &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddPduSessResSetupReqItem(item)

	snssai := &T3.Snssai{
		Sst:    1,
		SdPrst: true,
		Sd:     123,
	}
	msg.AddAllowedNssai(snssai)

	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA0)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA1)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA2)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA3)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA0)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA1)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA2)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA3)

	msg.SecKey = T3.SecurityKey{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	msg.UeRadioCap = []byte{1, 2, 3, 4, 5, 67}
	msg.OptFlags.Set(ngapmsg.ICSR_UeRadioCap)
	msg.IndexRFSelPriority = 1
	msg.OptFlags.Set(ngapmsg.ICSR_IndexRfSecPriority)
	msg.MaskedImeiSv = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	msg.OptFlags.Set(ngapmsg.ICSR_MaskedImeiSv)
	//msg.NasPdu = []byte{126,2,2,162,51,240,1}
	msg.NasPdu = []byte{126, 2, 2, 162, 51, 240, 1, 126, 0, 78, 80, 2, 32, 0, 38, 2, 0, 0}

	//msg.NasPdu = []byte{126, 0, 78, 80, 2, 32, 0 ,38 ,2 ,0, 0}
	msg.OptFlags.Set(ngapmsg.ICSR_Nas)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func main() {

	//debug for pprof
	//go func() {
	//	ip := "0.0.0.0:6060"
	//	if err := http.ListenAndServe(ip, nil); err != nil {
	//		fmt.Printf("start pprof failed on %s\n", ip)
	//		os.Exit(1)
	//	}
	//}()
	//ossCtxt := codec.NewOssCtxt()
	//ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	//defer codec.DeleteOssCtxt(ossCtxt)
	//
	//for i := 0; i < 1000; i++ {
	//
	//	ngSetupRequestMessageEncode(ossCtxtPtr)
	//}

	for i := 0; i < 1; i++ {
		run()

		time.Sleep(time.Millisecond * 5)
		time.Sleep(time.Millisecond * 5)
	}
	//go run()
	//go run()
	//
	//time.Sleep(time.Second * 5)

	//TestInitialContextSetupReqMsg3_Encode()
}
