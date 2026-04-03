package main

import "C"
import (
	codec "lite5gc/cmn/message/ngap/ngapcodec"

	"encoding/binary"
	"fmt"
	"lite5gc/cmn/utils"
	"net"
	"time"
	"unsafe"
)

func NgSetupRequestDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 21, 0, 53, 0, 0, 4, 0, 27, 0, 8, 0, 52, 54, 48, 0, 65, 66, 64, 0, 82, 64, 7, 2, 0, 82, 65, 78, 48, 49, 0, 102, 0, 18, 0, 0, 49, 50, 51, 0, 52, 54, 49, 0, 1, 2, 16, 144, 64, 49, 49, 49, 0, 21, 64, 1, 32}
	fmt.Println(encMsg)

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))

	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	setupReqCodec := codec.NewNgSetupRequestCodec()
	defer codec.DeleteNgSetupRequestCodec(setupReqCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	// num := setupReqCodec.DecodePerToOss(ossCtxtPtr, msgBuffer)
	// fmt.Println("return ", num)
	rt := setupReqCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	setupReqCodec.DumpMessage()
	fmt.Println("--------------------")
}
func NgSetupRequestEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	ngSetupReqEncode := codec.NewNgSetupRequestCodec()
	defer codec.DeleteNgSetupRequestCodec(ngSetupReqEncode)

	ngSetupReqEncode.SetRanNodeName("RAN01")
	ngSetupReqEncode.SetPagingDrx(1)

	//for gloabl gnb id
	ggnbid := codec.NewGGnbId()
	defer codec.DeleteGGnbId(ggnbid)

	plmnid := []byte("460")
	ggnbid.SetPlmnid(&plmnid[0])

	gnbid := codec.NewGnbId()
	defer codec.DeleteGnbId(gnbid)

	gnbid.SetLen(22)
	val := []byte("ABC")
	gnbid.SetVal(&val[0])

	ggnbid.SetGnbId(gnbid)
	ngSetupReqEncode.SetGgnbId(ggnbid)

	//for support ta list
	snssai1 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai1)
	sst1 := byte('A')
	snssai1.SetSst(&sst1)
	sd1 := []byte("111")
	snssai1.SetSd(&sd1[0])
	snssai1.SetSdPresent(true)

	snssai2 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai2)
	sst2 := byte('B')
	snssai2.SetSst(&sst2)
	sd2 := []byte("222")
	snssai2.SetSd(&sd2[0])
	snssai2.SetSdPresent(false)

	snssai3 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai3)
	sst3 := byte('C')
	snssai3.SetSst(&sst3)
	sd3 := []byte("333")
	snssai3.SetSd(&sd3[0])
	snssai3.SetSdPresent(false)

	snssai4 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai4)
	sst4 := byte('D')
	snssai4.SetSst(&sst4)
	sd4 := []byte("444")
	snssai4.SetSd(&sd4[0])
	snssai4.SetSdPresent(false)

	snssai5 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai5)
	sst5 := byte('E')
	snssai5.SetSst(&sst5)
	sd5 := []byte("222")
	snssai5.SetSd(&sd5[0])
	snssai5.SetSdPresent(false)

	snssai6 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai6)
	sst6 := byte('F')
	snssai6.SetSst(&sst6)
	sd6 := []byte("666")
	snssai6.SetSd(&sd6[0])
	snssai6.SetSdPresent(false)

	ssList := codec.NewSNssaiVector()
	defer codec.DeleteSNssaiVector(ssList)
	ssList.Add(snssai1)
	ssList.Add(snssai2)
	//ssList.Add(snssai3)
	//ssList.Add(snssai4)
	//ssList.Add(snssai5)
	//ssList.Add(snssai6)

	bplmn := codec.NewBPlmnItem()
	defer codec.DeleteBPlmnItem(bplmn)
	bplmnid := []byte("461")
	bplmn.SetPlmnid(&bplmnid[0])
	bplmn.SetSsList(ssList)

	bplmnList := codec.NewBPlmnItemVector()
	defer codec.DeleteBPlmnItemVector(bplmnList)
	bplmnList.Add(bplmn)

	stai := codec.NewSupTAItem()
	defer codec.DeleteSupTAItem(stai)
	tac := []byte("123")
	stai.SetTac(&tac[0])
	stai.SetBplmnList(bplmnList)

	ngSetupReqEncode.AddSupTAList(stai)

	//ngSetupReqEncode.SetUeRetentionInfo(byte(0))

	ngSetupReqEncode.DumpMessage()
	msgBuffer := ngSetupReqEncode.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	fmt.Println("encode msg buffer:", msgBuffer)

	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func NgSetupRespDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 21, 0, 50, 0, 0, 4, 0, 1, 0, 7, 2, 0, 65, 77, 70, 48, 49, 0, 96, 0, 15, 0, 64, 69, 15, 1, 2, 34, 130, 2, 0, 65, 77, 70, 48, 51, 0, 86, 64, 1, 1, 0, 80, 0, 8, 0, 52, 54, 49, 0, 0, 2, 16}
	fmt.Println(encMsg)

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))

	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	setupRespCodec := codec.NewNgSetupResponseCodec()
	defer codec.DeleteNgSetupResponseCodec(setupRespCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	// num := setupReqCodec.DecodePerToOss(ossCtxtPtr, msgBuffer)
	// fmt.Println("return ", num)
	rt := setupRespCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(setupRespCodec.ToString())
	fmt.Println("--------------------")
}
func NgSetupRespEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	ngSetupRespEncode := codec.NewNgSetupResponseCodec()
	defer codec.DeleteNgSetupResponseCodec(ngSetupRespEncode)

	ngSetupRespEncode.SetAmfName("AMF01")
	ngSetupRespEncode.SetRelativeAmfCapacity(1)

	//for gloabl gnb id
	//guami1 := codec.NewServedGuamiItem()
	//plmnid1 := []byte{0x46, 0x0f, 0x01}
	//guami1.SetPlmnId(&plmnid1[0])
	//amfid1 := codec.NewAmfIdentifier()
	//amfid1.SetPointer(1)
	//amfid1.SetRegionId(1)
	//amfSetId1 := []byte{0x11, 0x22}
	//amfid1.SetSetId(&amfSetId1[0])
	//guami1.SetAmfId(amfid1)
	//guami1.SetBackupAmfName("AMF02")
	//guami1.SetBackupAmfNamePrst(false)

	guami2 := codec.NewServedGuamiItem()
	defer codec.DeleteServedGuamiItem(guami2)

	plmnid2 := []byte{0x45, 0x0f, 0x01}
	guami2.SetPlmnId(&plmnid2[0])
	amfid2 := codec.NewAmfIdentifier()
	defer codec.DeleteAmfIdentifier(amfid2)

	amfid2.SetPointer(2)
	amfid2.SetRegionId(2)
	amfSetId2 := []byte{0x22, 0x22}
	amfid2.SetSetId(&amfSetId2[0])
	guami2.SetAmfId(amfid2)
	guami2.SetBackupAmfName("AMF03")
	guami2.SetBackupAmfNamePrst(true)

	//	ngSetupRespEncode.AddServedGuamiList(guami1)
	ngSetupRespEncode.AddServedGuamiList(guami2)

	//for support plmn list
	//snssai1 := codec.NewSNssai()
	//sst1 := byte('A')
	//snssai1.SetSst(&sst1)
	//sd1 := []byte("111")
	//snssai1.SetSd(&sd1[0])
	//snssai1.SetSdPresent(true)

	snssai2 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai2)
	sst2 := byte('B')
	snssai2.SetSst(&sst2)
	sd2 := []byte("222")
	snssai2.SetSd(&sd2[0])
	snssai2.SetSdPresent(false)

	ssList := codec.NewSNssaiVector()
	defer codec.DeleteSNssaiVector(ssList)
	//ssList.Add(snssai1)
	ssList.Add(snssai2)

	bplmn := codec.NewBPlmnItem()
	defer codec.DeleteBPlmnItem(bplmn)
	bplmnid := []byte("461")
	bplmn.SetPlmnid(&bplmnid[0])
	bplmn.SetSsList(ssList)

	ngSetupRespEncode.AddPlmnList(bplmn)

	//ngSetupRespEncode.SetUeRetentionInfo(byte(0))

	fmt.Println(ngSetupRespEncode.ToString())
	msgBuffer := ngSetupRespEncode.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func NgSetupFailureDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{64, 21, 0, 9, 0, 0, 1, 0, 15, 64, 2, 0, 64}
	fmt.Println(encMsg)

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))

	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	setupFailCodec := codec.NewNgSetupFailureCodec()
	defer codec.DeleteNgSetupFailureCodec(setupFailCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	// num := setupReqCodec.DecodePerToOss(ossCtxtPtr, msgBuffer)
	// fmt.Println("return ", num)
	rt := setupFailCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(setupFailCodec.ToString())
	fmt.Println("--------------------")
}
func NgSetupFailureEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	ngSetupFailEncode := codec.NewNgSetupFailureCodec()
	defer codec.DeleteNgSetupFailureCodec(ngSetupFailEncode)
	ngSetupFailEncode.SetCauseType(1)
	ngSetupFailEncode.SetCauseValue(1)
	//ngSetupFailEncode.SetRelativeTimeToWait(1)

	fmt.Println(ngSetupFailEncode.ToString())
	msgBuf := ngSetupFailEncode.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuf)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func InitialUeMsgDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 15, 64, 81, 0, 0, 8, 0, 0, 0, 5, 2, 65, 49, 49, 49, 0, 112, 64, 1, 0, 0, 3, 64, 2, 50, 0, 0, 26, 0, 7, 12, 131, 0, 49, 50, 51, 52, 0, 90, 64, 1, 16, 0, 121, 0, 19, 80, 49, 48, 48, 49, 50, 51, 52, 48, 49, 48, 48, 49, 48, 50, 49, 50, 51, 52, 0, 38, 0, 9, 8, 110, 97, 115, 80, 100, 117, 48, 49, 0, 85, 0, 2, 0, 1}
	fmt.Println(encMsg)

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)

	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))

	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	msg := codec.NewInitialUEMessageCodec()
	defer codec.DeleteInitialUEMessageCodec(msg)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	// num := setupReqCodec.DecodePerToOss(ossCtxtPtr, msgBuffer)
	// fmt.Println("return ", num)
	rt := msg.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(msg.ToString())
	fmt.Println("--------------------")
}
func InitialUeMsgEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	msg := codec.NewInitialUEMessageCodec()
	defer codec.DeleteInitialUEMessageCodec(msg)

	msg.SetRanUeNgapId(1)
	msg.SetNasPdu("nasPdu01")

	//user Location Information
	userLocationInfo := codec.NewUserLocationInfo()
	defer codec.DeleteUserLocationInfo(userLocationInfo)

	nrcgi := codec.NewNrCgi()
	defer codec.DeleteNrCgi(nrcgi)
	plmnid := []byte("100")
	nrcgi.SetPlmn(&plmnid[0])
	nrCGI1 := []byte("12345")
	nrcgi.SetNrCellId(&nrCGI1[0])

	timeStamp := []byte("1234")
	userLocationInfo.SetTimeStamp(&timeStamp[0])
	userLocationInfo.SetIsTimeStampPrst(true)
	userLocationInfo.SetNRCgi(nrcgi)

	tai := codec.NewTaiType()
	defer codec.DeleteTaiType(tai)

	tai.SetPlmn(&plmnid[0])
	tac := []byte("102")
	tai.SetTac(&tac[0])

	userLocationInfo.SetTai(tai)

	msg.SetUserLocationInfo(userLocationInfo)
	msg.SetRrcEstablishmentCause(2)

	//5g tmsi
	fivegctmsi := codec.NewStmsi5G()
	defer codec.DeleteStmsi5G(fivegctmsi)
	tempPointer := byte(0x31)
	fivegctmsi.SetPointer(tempPointer)
	tempTmsi := []byte{0x31, 0x32, 0x33, 0x34}
	fivegctmsi.SetTmsi(&tempTmsi[0])
	tempSetid := []byte{0x32, 0x33}
	fivegctmsi.SetSetId(&tempSetid[0])
	msg.Set5gStmsi(fivegctmsi)

	msg.SetAmfSetId(&tempSetid[0])

	msg.SetUeContextRequest(0)

	//for allowed nssai list
	snssai1 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai1)
	sst1 := byte('A')
	snssai1.SetSst(&sst1)
	sd1 := []byte("111")
	snssai1.SetSd(&sd1[0])
	snssai1.SetSdPresent(true)
	msg.AddAllowedNssai(snssai1)

	fmt.Println(msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func DownlinkNasTransportDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 4, 64, 76, 0, 0, 8, 0, 0, 0, 5, 2, 65, 49, 49, 49, 0, 110, 64, 8, 8, 118, 173, 241, 32, 118, 173, 240, 0, 31, 64, 2, 0, 64, 0, 38, 0, 9, 8, 110, 97, 115, 80, 100, 117, 48, 54, 0, 83, 64, 1, 5, 0, 48, 0, 12, 4, 128, 111, 108, 100, 65, 109, 102, 78, 97, 109, 101, 0, 10, 0, 2, 0, 64, 0, 85, 0, 2, 0, 32}
	fmt.Println(encMsg)

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))

	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	msg := codec.NewDownlinkNASTransportCodec()
	defer codec.DeleteDownlinkNASTransportCodec(msg)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	// num := setupReqCodec.DecodePerToOss(ossCtxtPtr, msgBuffer)
	// fmt.Println("return ", num)
	rt := msg.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(msg.ToString())
	fmt.Println("--------------------")
}
func DownlinkNasTransportEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	msg := codec.NewDownlinkNASTransportCodec()
	defer codec.DeleteDownlinkNASTransportCodec(msg)
	msg.SetRanUeNgapId(32)
	msg.SetAmfUeNgapId(64)
	msg.SetOldAmfName("oldAmfName")
	msg.SetRanPagingPriority(6)
	msg.SetNasPdu("nasPdu06")
	msg.SetIndexToRfsp(65)

	ueAmber := codec.NewUeAmbr()
	defer codec.DeleteUeAmbr(ueAmber)
	ueAmber.SetDownlink(7777777)
	ueAmber.SetUplink(7777776)
	msg.SetUeAmbr(ueAmber)

	//for allowed nssai list
	snssai1 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai1)
	sst1 := byte('A')
	snssai1.SetSst(&sst1)
	sd1 := []byte("111")
	snssai1.SetSd(&sd1[0])
	snssai1.SetSdPresent(true)
	msg.AddAllowedNssai(snssai1)

	//fmt.Println(msg.ToString())
	//msgBuffer := codec.NewMsgBuffer()
	//defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func UplinkNasTransportDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 46, 64, 63, 0, 0, 4, 0, 10, 0, 2, 0, 64, 0, 85, 0, 2, 0, 32, 0, 38, 0, 10, 9, 110, 97, 115, 80, 100, 117, 48, 48, 55, 0, 121, 64, 30, 72, 49, 48, 48, 49, 50, 51, 52, 48, 49, 48, 48, 49, 48, 50, 0, 0, 0, 149, 64, 9, 0, 49, 48, 48, 49, 50, 51, 52, 48}
	fmt.Println(encMsg)

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))

	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	msg := codec.NewUplinkNASTransportCodec()
	defer codec.DeleteUplinkNASTransportCodec(msg)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	// num := setupReqCodec.DecodePerToOss(ossCtxtPtr, msgBuffer)
	// fmt.Println("return ", num)
	rt := msg.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(msg.ToString())
	fmt.Println("--------------------")
}
func UplinkNasTransportEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	msg := codec.NewUplinkNASTransportCodec()
	defer codec.DeleteUplinkNASTransportCodec(msg)
	msg.SetRanUeNgapId(32)
	msg.SetAmfUeNgapId(64)
	msg.SetNasPdu("nasPdu007")

	//user Location Information
	userLocationInfo := codec.NewUserLocationInfo()
	defer codec.DeleteUserLocationInfo(userLocationInfo)

	nrcgi := codec.NewNrCgi()
	defer codec.DeleteNrCgi(nrcgi)
	plmnid := []byte("100")
	nrcgi.SetPlmn(&plmnid[0])
	nrCGI1 := []byte("12345")
	nrcgi.SetNrCellId(&nrCGI1[0])

	timeStamp := []byte("1234")
	userLocationInfo.SetTimeStamp(&timeStamp[0])
	userLocationInfo.SetIsTimeStampPrst(true)
	userLocationInfo.SetNRCgi(nrcgi)

	psCellInfo := codec.NewNrCgi()
	defer codec.DeleteNrCgi(psCellInfo)
	plmnid1 := []byte("321")
	psCellInfo.SetPlmn(&plmnid1[0])
	nrCGI2 := []byte("12345")
	psCellInfo.SetNrCellId(&nrCGI2[0])

	userLocationInfo.SetPsCellInfo(psCellInfo)
	userLocationInfo.SetPsCellInfoPrst(true)

	tai := codec.NewTaiType()
	defer codec.DeleteTaiType(tai)

	tai.SetPlmn(&plmnid[0])
	tac := []byte("102")
	tai.SetTac(&tac[0])

	userLocationInfo.SetTai(tai)

	msg.SetUserLocationInfo(userLocationInfo)

	//fmt.Println(msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func PduSessionResourceSetupRequestDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 29, 0, 128, 254, 0, 0, 5, 0, 10, 0, 2, 0, 1, 0, 85, 0, 2, 0, 2, 0, 83, 64, 1, 2, 0, 38, 0, 6, 5, 97, 98, 99, 100, 101, 0, 74, 0, 128, 219, 1, 0, 1, 64, 32, 123, 0, 0, 98, 0, 0, 7, 0, 136, 0, 39, 1, 1, 96, 224, 5, 9, 0, 0, 20, 0, 1, 244, 17, 92, 64, 3, 13, 64, 32, 7, 161, 32, 32, 8, 122, 35, 32, 10, 44, 42, 0, 0, 20, 0, 0, 10, 64, 1, 20, 0, 129, 0, 6, 128, 4, 5, 76, 86, 56, 0, 138, 0, 2, 68, 160, 0, 134, 0, 1, 0, 0, 127, 0, 1, 0, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 52, 51, 50, 49, 0, 130, 0, 8, 8, 91, 141, 128, 32, 7, 161, 32, 64, 1, 5, 97, 98, 99, 100, 101, 64, 32, 123, 0, 0, 98, 0, 0, 7, 0, 136, 0, 39, 1, 1, 96, 224, 5, 9, 0, 0, 20, 0, 1, 244, 17, 92, 64, 3, 13, 64, 32, 7, 161, 32, 32, 8, 122, 35, 32, 10, 44, 42, 0, 0, 20, 0, 0, 10, 64, 1, 20, 0, 129, 0, 6, 128, 4, 5, 76, 86, 56, 0, 138, 0, 2, 68, 160, 0, 134, 0, 1, 0, 0, 127, 0, 1, 0, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 52, 51, 50, 49, 0, 130, 0, 8, 8, 91, 141, 128, 32, 7, 161, 32}

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))
	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	reqCodec := codec.NewPduSessResSetupReqCodec()
	defer codec.DeletePduSessResSetupReqCodec(reqCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := reqCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(reqCodec.ToString())
	fmt.Println("--------------------")
}
func PduSessionResourceSetupRequestEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	requestCodec := codec.NewPduSessResSetupReqCodec()
	defer codec.DeletePduSessResSetupReqCodec(requestCodec)
	//AMF UE NGAP ID
	requestCodec.SetAmfUeNgapId(1)
	//RAN UE NGAP ID
	requestCodec.SetRanUeNgapId(2)

	//NAS PDU
	nasdata := []byte{'a', 'b', 'c', 'd', 'e'}
	requestCodec.SetNasPdu(string(nasdata))

	////UE  Aggregate Max Bit Rate
	//ambr := codec.NewUeAmbr()
	//defer codec.DeleteUeAmbr(ambr)
	//ambr.SetUplink(5000000)
	//ambr.SetDownlink(60000000)
	//requestCodec.SetUeAmbr(ambr)

	//RAN Paging Priority
	requestCodec.SetRanPagingPriority(3)

	//PDUSessionResourceSetupRequestList
	setupReqListItem := codec.NewPduSessResSetupReqItem()
	defer codec.DeletePduSessResSetupReqItem(setupReqListItem)
	setupReqListItem.SetPduSessId(1)

	snssai1 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai1)
	sst3 := byte(1)
	snssai1.SetSst(&sst3)
	sd3 := []byte{123}
	snssai1.SetSd(&sd3[0])
	snssai1.SetSdPresent(true)
	setupReqListItem.SetSnssai(snssai1)

	transfer := []byte{0, 0, 7, 0, 136, 0, 39, 1, 1, 96, 224, 5, 9, 0, 0, 20, 0, 1, 244, 17, 92, 64, 3, 13, 64, 32, 7, 161, 32, 32, 8, 122, 35, 32, 10, 44, 42, 0, 0, 20, 0, 0, 10, 64, 1, 20, 0, 129, 0, 6, 128, 4, 5, 76, 86, 56, 0, 138, 0, 2, 68, 160, 0, 134, 0, 1, 0, 0, 127, 0, 1, 0, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 52, 51, 50, 49, 0, 130, 0, 8, 8, 91, 141, 128, 32, 7, 161, 32}
	setupReqListItem.SetPduSessResSetupReqTransfer(string(transfer))

	nasdata1 := []byte{'a', 'b', 'c', 'd', 'e'}
	setupReqListItem.SetNasPduPrst(true)
	setupReqListItem.SetNasPdu(string(nasdata1))

	requestCodec.AddPduSessResSetupReqList(setupReqListItem)

	setupReqListItem1 := codec.NewPduSessResSetupReqItem()
	defer codec.DeletePduSessResSetupReqItem(setupReqListItem1)
	setupReqListItem1.SetPduSessId(1)

	snssai2 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai2)
	sst32 := byte(1)
	snssai2.SetSst(&sst32)
	sd2 := []byte{123}
	snssai2.SetSd(&sd2[0])
	snssai2.SetSdPresent(true)
	setupReqListItem1.SetSnssai(snssai2)

	transfer2 := []byte{0, 0, 7, 0, 136, 0, 39, 1, 1, 96, 224, 5, 9, 0, 0, 20, 0, 1, 244, 17, 92, 64, 3, 13, 64, 32, 7, 161, 32, 32, 8, 122, 35, 32, 10, 44, 42, 0, 0, 20, 0, 0, 10, 64, 1, 20, 0, 129, 0, 6, 128, 4, 5, 76, 86, 56, 0, 138, 0, 2, 68, 160, 0, 134, 0, 1, 0, 0, 127, 0, 1, 0, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 52, 51, 50, 49, 0, 130, 0, 8, 8, 91, 141, 128, 32, 7, 161, 32}
	setupReqListItem1.SetPduSessResSetupReqTransfer(string(transfer2))

	requestCodec.AddPduSessResSetupReqList(setupReqListItem1)

	// PDUSessionResourceSetupReqEncode
	fmt.Println(requestCodec.ToString())

	msgBuffer := requestCodec.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func PduSessResSetupReqTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {

	transfer := codec.NewPduSesResSetupReqTransferCodec()
	defer codec.DeletePduSesResSetupReqTransferCodec(transfer)

	// pdu session type
	transfer.SetPduSessType(0)

	// up tunnel info
	ipu32 := utils.Ip2long(net.ParseIP("10.180.9.242"))
	ipaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(ipaddr, ipu32)
	teid := make([]byte, 4)
	binary.LittleEndian.PutUint32(teid, 0x31323334)

	gtpTunnel := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(gtpTunnel)
	gtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2
	gtpTunnel.SetTransportLayerAddr(&ipaddr[0])
	gtpTunnel.SetGtpTeid(&teid[0])

	transfer.SetUpGtpTunnel(gtpTunnel)

	//Additional UL NG-U TNL Information List
	addGtpTunnelItem := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(addGtpTunnelItem)

	addIpu32 := utils.Ip2long(net.ParseIP("10.180.9.243"))
	addIpaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(addIpaddr, addIpu32)
	addTeid := make([]byte, 4)
	binary.LittleEndian.PutUint32(addTeid, 0x36373839)

	addGtpTunnelItem.SetIpType(0) //0 ipv4 1 ipv6 2
	addGtpTunnelItem.SetTransportLayerAddr(&addIpaddr[0])
	addGtpTunnelItem.SetGtpTeid(&addTeid[0])

	transfer.AddUpTransLayerInfoList(addGtpTunnelItem)

	//add up tunnel info
	//addGtpTunnel := codec.NewGtpTunnel()
	//defer codec.DeleteGtpTunnel(addGtpTunnel)
	//
	//addGtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2
	//
	//addIpu32 := utils.Ip2long(net.ParseIP("10.180.9.243"))
	//addIpaddr := make([]byte, 4)
	//binary.BigEndian.PutUint32(addIpaddr, addIpu32)
	//addTeid := make([]byte, 4)
	//binary.LittleEndian.PutUint32(addTeid, 0x36373839)
	//
	//addGtpTunnel := codec.NewGtpTunnel()
	//addGtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2
	//addGtpTunnel.SetTransportLayerAddr(&addIpaddr[0])
	//addGtpTunnel.SetGtpTeid(&addTeid[0])
	//
	//transfer.SetAddUpGtpTunnel(addGtpTunnel)

	// qos flow list
	qosFlowItem := codec.NewQosFlowSetupReqItem()
	defer codec.DeleteQosFlowSetupReqItem(qosFlowItem)
	qosFlowItem.SetQosFlowInd(1)
	qosFlowItem.SetERABIdPresent(true)
	qosFlowItem.SetERABId(20)

	qosFlowPara := codec.NewQosFlowLevelQosPara()
	defer codec.DeleteQosFlowLevelQosPara(qosFlowPara)
	arp := codec.NewARP()
	defer codec.DeleteARP(arp)
	arp.SetEmptionCapability(1)
	arp.SetEmptionVulnerability(1)
	arp.SetPriorityLevel(5)
	qosFlowPara.SetArp(arp)

	qosCc := codec.NewQosCharacter()
	defer codec.DeleteQosCharacter(qosCc)
	qosCc.SetIsDynamic(false)
	nonDynamic5qi := codec.NewNonDynamic5QI()
	defer codec.DeleteNonDynamic5QI(nonDynamic5qi)
	nonDynamic5qi.SetFiveQI(5)
	nonDynamic5qi.SetAveragingWindow(20)
	nonDynamic5qi.SetIsAveWindowPrst(true)
	nonDynamic5qi.SetIsPriorityLevelQosPrst(true)
	nonDynamic5qi.SetPriorityLevelQos(10)
	nonDynamic5qi.SetIsMaxDataBusrtVolPrst(true)
	nonDynamic5qi.SetMaximumDataBurstVolume(500)
	qosCc.SetNonDynamic5qi(nonDynamic5qi)
	qosFlowPara.SetQosCharacter(qosCc)

	qosFlowPara.SetGBRQosInfoPresent(true)
	gbrInfo := codec.NewGBRQosInformation()
	defer codec.DeleteGBRQosInformation(gbrInfo)
	gbrInfo.SetMaxFlowBitRateDL(200000)
	gbrInfo.SetMaxFlowBitRateUL(500000)
	gbrInfo.SetGuaFlowBitRateDL(555555)
	gbrInfo.SetGuaFlowBitRateUL(666666)
	gbrInfo.SetMaxPacketLossRateDLPresent(true)
	gbrInfo.SetMaxPacketLossRateUL(10)
	gbrInfo.SetMaxPacketLossRateULPresent(true)
	gbrInfo.SetMaxPacketLossRateDL(20)
	gbrInfo.SetNotiControlPresent(true)
	gbrInfo.SetNotiControl(0)
	qosFlowPara.SetGBRQosInfo(gbrInfo)

	qosFlowPara.SetRefQosAttrPresent(true)
	qosFlowPara.SetRefQosAttr(0)
	qosFlowPara.SetAddQosFlowInfoPresent(false)
	qosFlowItem.SetQosFlowLevQosPara(qosFlowPara)

	transfer.AddQosFlowSetupReqList(qosFlowItem)

	transfer.SetSessAmbr(500000, 6000000)
	transfer.SetDataFwNotPssble(0)
	transfer.SetSecInd(1, 1)
	transfer.SetMaxPrtDataRate(1)
	transfer.SetMaxPrtDataRateDl(0)
	transfer.SetNtwkInstance(88888888)

	transfer.SetCommonNetworkInstance("hello world")

	fmt.Println(transfer.ToString())

	msgBuffer := transfer.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func PduSessResSetupReqTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 0, 9, 0, 166, 64, 12, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0, 130, 0, 8, 8, 91, 141, 128, 32, 7, 161, 32, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 52, 51, 50, 49, 0, 126, 0, 10, 0, 31, 10, 180, 9, 243, 57, 56, 55, 54, 0, 127, 0, 1, 0, 0, 134, 0, 1, 0, 0, 138, 0, 2, 68, 160, 0, 129, 0, 6, 128, 4, 5, 76, 86, 56, 0, 136, 0, 39, 1, 1, 96, 224, 5, 9, 0, 0, 20, 0, 1, 244, 17, 92, 64, 3, 13, 64, 32, 7, 161, 32, 32, 8, 122, 35, 32, 10, 44, 42, 0, 0, 20, 0, 0, 10, 64, 1, 20}

	transfer := codec.NewPduSesResSetupReqTransferCodec()
	defer codec.DeletePduSesResSetupReqTransferCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func PduSessResSetupRespTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {
	transfer := codec.NewPduSesResSetupRespTransferCodec()
	defer codec.DeletePduSesResSetupRespTransferCodec(transfer)

	gtpTunnel := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(gtpTunnel)

	gtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2

	ipu32 := utils.Ip2long(net.ParseIP("10.180.9.242"))
	ipaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(ipaddr, ipu32)
	gtpTunnel.SetTransportLayerAddr(&ipaddr[0])

	teid := make([]byte, 4)
	binary.LittleEndian.PutUint32(teid, 1234)
	gtpTunnel.SetGtpTeid(&teid[0])

	transfer.SetGtpTunnel(gtpTunnel)

	associateQosFlowItem := codec.NewAssQosFlowItem()
	defer codec.DeleteAssQosFlowItem(associateQosFlowItem)
	associateQosFlowItem.SetQosFlowInd(1)
	//associateQosFlowItem1.SetQosFlowMapInd(0x01)
	associateQosFlowItem.SetQosFlowMapIndPrst(false)

	transfer.AddQosFlowIndList(associateQosFlowItem)

	addQosFlowTnlInfo := codec.NewAddQosFlowPerTNLInfo()
	defer codec.DeleteAddQosFlowPerTNLInfo(addQosFlowTnlInfo)

	addGtpTunnel := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(addGtpTunnel)

	addGtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2

	ipu32b := utils.Ip2long(net.ParseIP("10.180.9.244"))
	addIpaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(addIpaddr, ipu32b)
	addGtpTunnel.SetTransportLayerAddr(&addIpaddr[0])

	addTeid1 := make([]byte, 4)
	binary.LittleEndian.PutUint32(addTeid1, 1236)
	addGtpTunnel.SetGtpTeid(&addTeid1[0])

	addQosFlowTnlInfo.SetUpTransportLayerInfo(addGtpTunnel)

	assQosFlowVector := codec.NewAssQosFlowItemVector()
	defer codec.DeleteAssQosFlowItemVector(assQosFlowVector)

	associateQosFlowItem1 := codec.NewAssQosFlowItem()
	defer codec.DeleteAssQosFlowItem(associateQosFlowItem1)
	associateQosFlowItem1.SetQosFlowInd(1)
	associateQosFlowItem1.SetQosFlowMapInd(0x01)
	associateQosFlowItem1.SetQosFlowMapIndPrst(true)

	assQosFlowVector.Add(associateQosFlowItem1)

	addQosFlowTnlInfo.SetAssQosFlowList(assQosFlowVector)

	transfer.AddQosFlowTNLInfoList(addQosFlowTnlInfo)

	transfer.SetSecResult(1, 1)

	qosflow := codec.NewQosFlowCodecItem()
	defer codec.DeleteQosFlowCodecItem(qosflow)
	qosflow.SetQosFlowInd(6)
	qosflow.SetCauseType(1)
	qosflow.SetCauseValue(1)
	transfer.AddQosFlowFailedSetupList(qosflow)

	qosflow2 := codec.NewQosFlowCodecItem()
	defer codec.DeleteQosFlowCodecItem(qosflow2)
	qosflow2.SetQosFlowInd(7)
	qosflow2.SetCauseType(1)
	qosflow2.SetCauseValue(2)
	transfer.AddQosFlowFailedSetupList(qosflow2)

	fmt.Println(transfer.ToString())

	msgBuffer := transfer.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PduSessResSetupRespTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{112, 3, 224, 10, 180, 9, 242, 210, 4, 0, 0, 0, 1, 0, 7, 192, 10, 180, 9, 244, 212, 4, 0, 0, 1, 1, 69, 4, 14, 1, 1, 128, 16}

	transfer := codec.NewPduSesResSetupRespTransferCodec()
	defer codec.DeletePduSesResSetupRespTransferCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}
func CauseEncode(ossCtxtPtr codec.NgapOssCtxt) {
	transfer := codec.NewPduSessResSetupUnSuccTransferCodec()
	defer codec.DeletePduSessResSetupUnSuccTransferCodec(transfer)

	transfer.SetCause(1, 2)

	fmt.Println(transfer.ToString())

	msgBuffer := transfer.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func CauseDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 16}

	transfer := codec.NewPduSessResSetupUnSuccTransferCodec()
	defer codec.DeletePduSessResSetupUnSuccTransferCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}
func PduSessResSetupRespEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewPduSessResSetupRespCodec()
	defer codec.DeletePduSessResSetupRespCodec(msg)
	msg.SetRanUeNgapId(20)
	msg.SetAmfUeNgapId(25)
	respItem := codec.NewPduSessResSetupRespItem()
	defer codec.DeletePduSessResSetupRespItem(respItem)
	respItem.SetPduSessId(5)
	succTransfer := []byte{48, 3, 224, 10, 180, 9, 242, 52, 18, 0, 0, 1, 1, 69, 4, 14, 112, 4, 127, 255, 255, 253, 3, 0, 32}
	respItem.SetPduSessResSetupRespTransfer(string(succTransfer))
	msg.AddPduSessResSetupRespList(respItem)

	//failItem := codec.NewPduSessResFailedSetupItem()
	//failItem.SetPduSessId(6)
	//failTransfer :=[]byte{40}
	//failItem.SetPduSessResSetupUnSussTransfer(string(failTransfer))
	//msg.AddPduSessResFailedSetupList(failItem)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PduSessResSetupRespDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 29, 0, 48, 0, 0, 3, 0, 75, 64, 29, 0, 0, 5, 25, 48, 3, 224, 10, 180, 9, 242, 52, 18, 0, 0, 1, 1, 69, 4, 14, 112, 4, 127, 255, 255, 253, 3, 0, 32, 0, 10, 64, 2, 0, 25, 0, 85, 64, 2, 0, 20}
	transfer := codec.NewPduSessResSetupRespCodec()
	defer codec.DeletePduSessResSetupRespCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func InitialContextSetupRequestEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewInitialContextSetupRequestCodec()
	defer codec.DeleteInitialContextSetupRequestCodec(msg)

	msg.SetRanUeNgapId(5)
	msg.SetAmfUeNgapId(6)
	msg.SetRRCITRR(1)
	msg.SetIndexToRfsp(5)
	msg.SetNasPdu("hello world")
	msg.SetOldAmfName("amf.shenzhen")
	msg.SetSecurityKey("security.key.123456")
	msg.SetUERadioCapability("ue.radio.capbility.123456")
	msg.SetMaskedImeiSv("masked.imeisv.123456")
	ambr := codec.NewUeAmbr()
	defer codec.DeleteUeAmbr(ambr)
	ambr.SetUplink(5000000)
	ambr.SetDownlink(60000000)
	msg.SetUeAmbr(ambr)

	tempUeRadioCapForPaging := codec.NewUeRadioCapForPaging()
	defer codec.DeleteUeRadioCapForPaging(tempUeRadioCapForPaging)
	tempUeRadioCapForPaging.SetUeRadioCapForPagingNr("ue.radio.capbility.for.paging.123456 Nr")
	tempUeRadioCapForPaging.SetUeRadioCapForPagingNrPrst(true)
	tempUeRadioCapForPaging.SetUeRadioCapForPagingEutra("ue.radio.capbility.for.paging.123456 Eutra")
	tempUeRadioCapForPaging.SetUeRadioCapForPagingEutraPrst(true)
	//msg.SetUeRadioCapForPaging(tempUeRadioCapForPaging)

	guami := codec.NewGuami()
	defer codec.DeleteGuami(guami)
	plmnid := []byte{0x46, 0x0f, 0x01}
	guami.SetPlmnId(&plmnid[0])
	amfid := codec.NewAmfIdentifier()
	defer codec.DeleteAmfIdentifier(amfid)
	amfid.SetPointer(1)
	amfid.SetRegionId(1)
	amfSetId := []byte{0x11, 0x22}
	amfid.SetSetId(&amfSetId[0])
	guami.SetAmfId(amfid)
	msg.SetGuami(guami)

	ueSecCap := codec.NewUeSecurityCap()
	defer codec.DeleteUeSecurityCap(ueSecCap)
	val := []byte{0xe0, 0x00}
	ueSecCap.SetNrIntegrityProAlg(&val[0])
	ueSecCap.SetNrEncryptionAlg(&val[0])
	ueSecCap.SetEutraEncryptionAlg(&val[0])
	ueSecCap.SetEutraIntegrityProAlg(&val[0])
	msg.SetUeSecurityCap(ueSecCap)

	trace := codec.NewTraceAct()
	defer codec.DeleteTraceAct(trace)
	trace.SetInterfaces(1)
	trace.SetTraceDepth(2)
	traceId := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
	trace.SetTraceID(&traceId[0])
	traceEntyIp := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00}
	trace.SetTraceEntityIPAddr(&traceEntyIp[0])
	msg.SetTraceAct(trace)

	efbid := codec.NewEmergeFBReqInd()
	defer codec.DeleteEmergeFBReqInd(efbid)
	efbid.SetIndicator(0)
	efbid.SetServiceTargetCn(0)
	msg.SetEmergeFBInd(efbid)

	snssai := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai)
	snssai.SetSdPresent(true)
	var sst byte = 1
	snssai.SetSst(&sst)
	sd := []byte{1, 2, 3}
	snssai.SetSd(&sd[0])
	msg.AddAllowedSNssai(snssai)

	//PDUSessionResourceSetupRequestList
	setupReqListItem := codec.NewPduSessResSetupReqItem()
	defer codec.DeletePduSessResSetupReqItem(setupReqListItem)
	setupReqListItem.SetPduSessId(1)

	snssai1 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai1)
	sst3 := byte(1)
	snssai1.SetSst(&sst3)
	sd3 := []byte{123}
	snssai1.SetSd(&sd3[0])
	snssai1.SetSdPresent(true)
	setupReqListItem.SetSnssai(snssai1)

	transfer := []byte{0, 0, 7, 0, 136, 0, 39, 1, 1, 96, 224, 5, 9, 0, 0, 20, 0, 1, 244, 17, 92, 64, 3, 13, 64, 32, 7, 161, 32, 32, 8, 122, 35, 32, 10, 44, 42, 0, 0, 20, 0, 0, 10, 64, 1, 20, 0, 129, 0, 6, 128, 4, 5, 76, 86, 56, 0, 138, 0, 2, 68, 160, 0, 134, 0, 1, 0, 0, 127, 0, 1, 0, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 52, 51, 50, 49, 0, 130, 0, 8, 8, 91, 141, 128, 32, 7, 161, 32}
	setupReqListItem.SetPduSessResSetupReqTransfer(string(transfer))

	nasdata1 := []byte{'a', 'b', 'c', 'd', 'e'}
	setupReqListItem.SetNasPduPrst(true)
	setupReqListItem.SetNasPdu(string(nasdata1))

	msg.AddPduSessResReqList(setupReqListItem)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func InitialContextSetupRequestDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 14, 0, 129, 174, 0, 0, 17, 0, 118, 64, 84, 96, 39, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 102, 111, 114, 46, 112, 97, 103, 105, 110, 103, 46, 49, 50, 51, 52, 53, 54, 32, 78, 114, 42, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 102, 111, 114, 46, 112, 97, 103, 105, 110, 103, 46, 49, 50, 51, 52, 53, 54, 32, 69, 117, 116, 114, 97, 0, 91, 64, 1, 32, 0, 24, 0, 1, 0, 0, 34, 64, 8, 109, 97, 115, 107, 101, 100, 46, 105, 0, 38, 64, 12, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0, 31, 64, 2, 0, 4, 0, 117, 64, 26, 25, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 49, 50, 51, 52, 53, 54, 0, 108, 64, 32, 0, 49, 50, 51, 52, 53, 54, 55, 56, 1, 36, 248, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 94, 0, 32, 115, 101, 99, 117, 114, 105, 116, 121, 46, 107, 101, 121, 46, 49, 50, 51, 52, 53, 54, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 119, 0, 9, 28, 0, 14, 0, 7, 0, 3, 128, 0, 0, 71, 0, 113, 0, 64, 1, 5, 97, 98, 99, 100, 101, 64, 32, 123, 0, 0, 98, 0, 0, 7, 0, 136, 0, 39, 1, 1, 96, 224, 5, 9, 0, 0, 20, 0, 1, 244, 17, 92, 64, 3, 13, 64, 32, 7, 161, 32, 32, 8, 122, 35, 32, 10, 44, 42, 0, 0, 20, 0, 0, 10, 64, 1, 20, 0, 129, 0, 6, 128, 4, 5, 76, 86, 56, 0, 138, 0, 2, 68, 160, 0, 134, 0, 1, 0, 0, 127, 0, 1, 0, 0, 139, 0, 10, 1, 240, 10, 180, 9, 242, 52, 51, 50, 49, 0, 130, 0, 8, 8, 91, 141, 128, 32, 7, 161, 32, 0, 0, 0, 5, 2, 1, 1, 2, 3, 0, 28, 0, 7, 0, 70, 15, 1, 1, 17, 129, 0, 110, 0, 9, 12, 3, 147, 135, 0, 32, 76, 75, 64, 0, 48, 0, 14, 5, 128, 97, 109, 102, 46, 115, 104, 101, 110, 122, 104, 101, 110, 0, 10, 0, 2, 0, 6, 0, 85, 0, 2, 0, 5}

	transfer := codec.NewInitialContextSetupRequestCodec()
	defer codec.DeleteInitialContextSetupRequestCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func InitialContextSetupResponseEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewInitialContextSetupRespCodec()
	defer codec.DeleteInitialContextSetupRespCodec(msg)
	msg.SetAmfUeNgapId(5)
	msg.SetRanUeNgapId(6)

	respItem := codec.NewPduSessResSetupRespItem()
	defer codec.DeletePduSessResSetupRespItem(respItem)
	respItem.SetPduSessId(5)
	succTransfer := []byte{48, 3, 224, 10, 180, 9, 242, 52, 18, 0, 0, 1, 1, 69, 4, 14, 112, 4, 127, 255, 255, 253, 3, 0, 32}
	respItem.SetPduSessResSetupRespTransfer(string(succTransfer))
	msg.AddPduSessResSetupRespList(respItem)

	failItem := codec.NewPduSessResFailedSetupItem()
	defer codec.DeletePduSessResFailedSetupItem(failItem)
	failItem.SetPduSessId(6)
	failTransfer := []byte{0, 32}
	failItem.SetPduSessResSetupUnSussTransfer(string(failTransfer))
	msg.AddPduSessResFailedSetupList(failItem)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func InitialContextFailToSetupDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{64, 14, 0, 31, 0, 0, 4, 0, 10, 64, 2, 0, 15, 0, 85, 64, 2, 0, 10, 0, 132, 64, 6, 0, 0, 1, 2, 0, 32, 0, 15, 64, 2, 0, 192}
	transfer := codec.NewInitialContextSetupFailureCodec()
	defer codec.DeleteInitialContextSetupFailureCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func InitialContextFailToSetupEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewInitialContextSetupFailureCodec()
	defer codec.DeleteInitialContextSetupFailureCodec(msg)
	msg.SetRanUeNgapId(10)
	msg.SetAmfUeNgapId(15)
	msg.SetCauseType(1)
	msg.SetCauseValue(3)

	//PDUSessionResourceFailedToSetupListCxtFail
	failToSetupListItem := codec.NewPduSessResFailedSetupItem()
	defer codec.DeletePduSessResFailedSetupItem(failToSetupListItem)
	failToSetupListItem.SetPduSessId(1)

	transfer := []byte{0, 32}
	failToSetupListItem.SetPduSessResSetupUnSussTransfer(string(transfer))

	msg.AddPduSessResFailToSetupList(failToSetupListItem)

	fmt.Println(msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func InitialContextSetupResponseDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 14, 0, 58, 0, 0, 4, 0, 55, 64, 6, 0, 0, 6, 2, 0, 32, 0, 72, 64, 29, 0, 0, 5, 25, 48, 3, 224, 10, 180, 9, 242, 52, 18, 0, 0, 1, 1, 69, 4, 14, 112, 4, 127, 255, 255, 253, 3, 0, 32, 0, 10, 64, 2, 0, 5, 0, 85, 64, 2, 0, 6}
	transfer := codec.NewInitialContextSetupRespCodec()
	defer codec.DeleteInitialContextSetupRespCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func UeContextReleaseRequestEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewUeContextReleaseRequestCodec()
	defer codec.DeleteUeContextReleaseRequestCodec(msg)
	msg.SetRanUeNgapId(20)
	msg.SetAmfUeNgapId(25)
	msg.SetCauseType(1)
	msg.SetCauseValue(3)

	//PDUSessionResourceListCxtRelReq
	var pduSessId uint16 = 10
	msg.AddPduSessResListRelReq(pduSessId)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func UeContextReleaseRequestDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 42, 64, 28, 0, 0, 4, 0, 133, 0, 3, 0, 0, 10, 0, 15, 64, 2, 0, 192, 0, 10, 0, 2, 0, 25, 0, 85, 0, 2, 0, 20}
	//fmt.Println(encMsg)
	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))
	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	UeContextReleaseRequestCodec := codec.NewUeContextReleaseRequestCodec()
	defer codec.DeleteUeContextReleaseRequestCodec(UeContextReleaseRequestCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := UeContextReleaseRequestCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(UeContextReleaseRequestCodec.ToString())
	fmt.Println("--------------------")
}

func UeContextReleaseCommandEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewUeContextReleaseCommandCodec()
	defer codec.DeleteUeContextReleaseCommandCodec(msg)
	msg.SetCauseType(1)
	msg.SetCauseValue(3)
	msg.SetUeNgapIdsType(1)
	msg.SetAmfUeNgapId(8)
	msg.SetRanUeNgapId(9)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func UeContextReleaseCommandDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 41, 0, 17, 0, 0, 2, 0, 15, 64, 2, 0, 192, 0, 114, 0, 4, 0, 8, 0, 9}
	fmt.Println(encMsg)
	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))
	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	UeContextReleaseCommandCodec := codec.NewUeContextReleaseCommandCodec()
	defer codec.DeleteUeContextReleaseCommandCodec(UeContextReleaseCommandCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := UeContextReleaseCommandCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(UeContextReleaseCommandCodec.ToString())
	fmt.Println("--------------------")
}

func UeContextReleaseCommandTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 32}
	fmt.Println(encMsg)

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))

	//ret := ossCodec.DecodeHeader(ossCtxtPtr)
	//fmt.Println("deocde header restult:", ret)
	//
	//msgType := ossCodec.GetMsgType()
	//procCode := ossCodec.GetProcedureCode()
	//fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	relCmdTransferCodec := codec.NewPduSessResRelCmdTransferCodec()
	defer codec.DeletePduSessResRelCmdTransferCodec(relCmdTransferCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := relCmdTransferCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(relCmdTransferCodec.ToString())
	fmt.Println("--------------------")
}
func UeContextReleaseCommandTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	relCmdTranEncode := codec.NewPduSessResRelCmdTransferCodec()
	defer codec.DeletePduSessResRelCmdTransferCodec(relCmdTranEncode)
	relCmdTranEncode.SetCause(1, 2)

	fmt.Println(relCmdTranEncode.ToString())
	msgBuf := relCmdTranEncode.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuf)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func UeContextReleaseCompleteEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewUeContextReleaseCompleteCodec()
	defer codec.DeleteUeContextReleaseCompleteCodec(msg)

	msg.SetRanUeNgapId(20)
	msg.SetAmfUeNgapId(25)

	//user Location Information
	userLocationInfo := codec.NewUserLocationInfo()
	defer codec.DeleteUserLocationInfo(userLocationInfo)

	nrcgi := codec.NewNrCgi()
	defer codec.DeleteNrCgi(nrcgi)
	plmnid := []byte("100")
	nrcgi.SetPlmn(&plmnid[0])
	nrCGI1 := []byte("12345")
	nrcgi.SetNrCellId(&nrCGI1[0])

	timeStamp := []byte("1234")
	userLocationInfo.SetTimeStamp(&timeStamp[0])
	userLocationInfo.SetIsTimeStampPrst(true)
	userLocationInfo.SetNRCgi(nrcgi)

	psCellInfo := codec.NewNrCgi()
	defer codec.DeleteNrCgi(psCellInfo)
	plmnid1 := []byte("100")
	psCellInfo.SetPlmn(&plmnid1[0])
	nrCGI2 := []byte("12345")
	psCellInfo.SetNrCellId(&nrCGI2[0])
	userLocationInfo.SetPsCellInfo(psCellInfo)
	userLocationInfo.SetPsCellInfoPrst(true)

	tai := codec.NewTaiType()
	defer codec.DeleteTaiType(tai)

	tai.SetPlmn(&plmnid[0])
	tac := []byte("102")
	tai.SetTac(&tac[0])

	userLocationInfo.SetTai(tai)

	msg.SetUserLocationInfo(userLocationInfo)

	//InfoOnRecommendedCellsAndRANNodesForPaging
	cellItem := codec.NewRecommandCellItem()
	defer codec.DeleteRecommandCellItem(cellItem)

	cellItem.SetTimeStayedInCell(uint16(512))
	cellItem.SetTimeStayedPrst(true)
	cellItem.SetNrcgi(nrcgi)

	cellList := codec.NewRecommandCellItemVector()
	defer codec.DeleteRecommandCellItemVector(cellList)
	cellList.Add(cellItem)

	msg.SetRecommandCellList(cellList)

	ranNodeItem := codec.NewRecommandRanNodeItem()
	defer codec.DeleteRecommandRanNodeItem(ranNodeItem)
	ranNodeItem.SetTargetChoice(2)

	ggnbid := codec.NewGGnbId()
	defer codec.DeleteGGnbId(ggnbid)
	ggnbid.SetPlmnid(&plmnid[0])

	gnbid := codec.NewGnbId()
	defer codec.DeleteGnbId(gnbid)
	gnbid.SetLen(22)
	val := []byte("ABC")
	gnbid.SetVal(&val[0])

	ggnbid.SetGnbId(gnbid)
	ranNodeItem.SetGgnbId(ggnbid)

	ranNodeItem.SetTai(tai)

	ranNodeList := codec.NewRecommandRanNodeItemVector()
	defer codec.DeleteRecommandRanNodeItemVector(ranNodeList)
	ranNodeList.Add(ranNodeItem)

	msg.SetRecommandRanNodeList(ranNodeList)

	//PDUSessionResourceListCxtRelCpl
	var pduSessId byte = 5
	pduSessResListRelCplItem := codec.NewPduSessResRelCplItem()
	defer codec.DeletePduSessResRelCplItem(pduSessResListRelCplItem)

	pduSessResListRelCplItem.SetPduSessId(pduSessId)

	msg.AddPduSessResListRelCpl(pduSessResListRelCplItem)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func UeContextReleaseCompleteDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 41, 0, 80, 0, 0, 5, 0, 10, 64, 2, 0, 25, 0, 85, 64, 2, 0, 20, 0, 121, 64, 30, 72, 49, 48, 48, 49, 50, 51, 52, 48, 49, 48, 48, 49, 48, 50, 0, 0, 0, 149, 64, 9, 0, 49, 48, 48, 49, 50, 51, 52, 48, 0, 32, 64, 20, 0, 64, 49, 48, 48, 49, 50, 51, 52, 48, 2, 0, 0, 64, 49, 48, 48, 49, 48, 50, 0, 60, 0, 3, 0, 0, 5}
	fmt.Println(encMsg)
	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))
	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	ueContextRelCplCodec := codec.NewUeContextReleaseCompleteCodec()
	defer codec.DeleteUeContextReleaseCompleteCodec(ueContextRelCplCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := ueContextRelCplCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(ueContextRelCplCodec.ToString())
	fmt.Println("--------------------")
}

func UeRadioCapaInfoIndicationEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewUeRadioCapaInfoIndicationCodec()
	defer codec.DeleteUeRadioCapaInfoIndicationCodec(msg)

	msg.SetRanUeNgapId(20)
	msg.SetAmfUeNgapId(25)
	msg.SetUERadioCapability("ue.radio.capbility.123456")

	tempUeRadioCapForPaging := codec.NewUeRadioCapForPaging()
	defer codec.DeleteUeRadioCapForPaging(tempUeRadioCapForPaging)
	var cap4PagingNr string = "ue.radio.capbility.for.paging.123456 Nr"
	tempUeRadioCapForPaging.SetUeRadioCapForPagingNr(cap4PagingNr)
	tempUeRadioCapForPaging.SetUeRadioCapForPagingNrPrst(true)
	var cap4PagingEurtra string = "ue.radio.capbility.for.paging.123456 Eutra"
	tempUeRadioCapForPaging.SetUeRadioCapForPagingEutra(cap4PagingEurtra)
	tempUeRadioCapForPaging.SetUeRadioCapForPagingEutraPrst(true)

	//msg.SetUeRadioCapForPaging(tempUeRadioCapForPaging)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func UeRadioCapaInfoIndicationDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 44, 64, 128, 133, 0, 0, 4, 0, 118, 64, 84, 96, 39, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 102, 111, 114, 46, 112, 97, 103, 105, 110, 103, 46, 49, 50, 51, 52, 53, 54, 32, 78, 114, 42, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 102, 111, 114, 46, 112, 97, 103, 105, 110, 103, 46, 49, 50, 51, 52, 53, 54, 32, 69, 117, 116, 114, 97, 0, 117, 64, 26, 25, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 49, 50, 51, 52, 53, 54, 0, 10, 0, 2, 0, 25, 0, 85, 0, 2, 0, 20}
	fmt.Println(encMsg)
	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))
	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	ueRadioCapaInfoIndicationCodec := codec.NewUeRadioCapaInfoIndicationCodec()
	defer codec.DeleteUeRadioCapaInfoIndicationCodec(ueRadioCapaInfoIndicationCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := ueRadioCapaInfoIndicationCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(ueRadioCapaInfoIndicationCodec.ToString())
	fmt.Println("--------------------")
}

func PduSessResRelCmdTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {
	transfer := codec.NewPduSessResRelCmdTransferCodec()
	defer codec.DeletePduSessResRelCmdTransferCodec(transfer)

	transfer.SetCause(1, 2)

	msgBuffer := transfer.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PduSessResRelCmdTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 32}

	transfer := codec.NewPduSessResRelCmdTransferCodec()
	defer codec.DeletePduSessResRelCmdTransferCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}
func PduSessResRelCmdEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewPduSessResRelCmdCodec()
	defer codec.DeletePduSessResRelCmdCodec(msg)
	msg.SetRanUeNgapId(20)
	msg.SetAmfUeNgapId(25)
	//NAS PDU
	nasdata := []byte{'a', 'b', 'c', 'd', 'e'}
	msg.SetNasPdu(string(nasdata))

	msg.SetRanPagingPriority(6)

	relcmdItem := codec.NewPduSessResRelCmdItem()
	defer codec.DeletePduSessResRelCmdItem(relcmdItem)
	relcmdItem.SetPduSessId(5)
	relcmdTransfer := []byte{0, 32}
	relcmdItem.SetPduSessResRelCmdTransfer(string(relcmdTransfer))
	msg.AddPduSessResRelCmdList(relcmdItem)

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PduSessResRelCmdDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 28, 0, 40, 0, 0, 5, 0, 79, 0, 6, 0, 0, 5, 2, 0, 32, 0, 83, 64, 1, 5, 0, 38, 64, 6, 5, 97, 98, 99, 100, 101, 0, 10, 0, 2, 0, 25, 0, 85, 0, 2, 0, 20}
	transfer := codec.NewPduSessResRelCmdCodec()
	defer codec.DeletePduSessResRelCmdCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func ErrorIndicationEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewErrorIndicationCodec()
	defer codec.DeleteErrorIndicationCodec(msg)
	msg.SetAmfUeNgapId(5)
	msg.SetRanUeNgapId(6)

	msg.SetCauseType(1)
	msg.SetCauseValue(1)

	fmt.Println("encode status", msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func ErrorIndicationDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 9, 64, 21, 0, 0, 3, 0, 15, 64, 2, 0, 64, 0, 10, 64, 2, 0, 5, 0, 85, 64, 2, 0, 6}

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))
	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	errCodec := codec.NewErrorIndicationCodec()
	defer codec.DeleteErrorIndicationCodec(errCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := errCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(errCodec.ToString())
	fmt.Println("--------------------")
}
func PagingEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewPagingCodec()
	defer codec.DeletePagingCodec(msg)

	uePagingId := codec.NewUePagingIdentity()
	defer codec.DeleteUePagingIdentity(uePagingId)
	amfSetId := []byte{0x01, 0xb2}
	uePagingId.SetSetId(&amfSetId[0])
	tmsi := []byte("6666")
	uePagingId.SetTmsi(&tmsi[0])
	uePagingId.SetPointer(0x37)
	msg.SetUePagingIdentity(uePagingId)

	msg.SetPagingDrx(1)

	tai := codec.NewTaiType()
	defer codec.DeleteTaiType(tai)
	plmnid := []byte("100")
	tai.SetPlmn(&plmnid[0])
	tac := []byte("102")
	tai.SetTac(&tac[0])

	msg.AddTaiPagingList(tai)

	msg.SetPagingDrx(2)

	tempUeRadioCapForPaging := codec.NewUeRadioCapForPaging()
	defer codec.DeleteUeRadioCapForPaging(tempUeRadioCapForPaging)
	tempUeRadioCapForPaging.SetUeRadioCapForPagingNr("ue.radio.capbility.for.paging.123456 Nr")
	tempUeRadioCapForPaging.SetUeRadioCapForPagingNrPrst(false)
	tempUeRadioCapForPaging.SetUeRadioCapForPagingEutra("ue.radio.capbility.for.paging.123456 Eutra")
	tempUeRadioCapForPaging.SetUeRadioCapForPagingEutraPrst(true)
	msg.SetUeRadioCapForPaging(tempUeRadioCapForPaging)

	msg.SetPagingOrigin(0)

	assDataForPaging := codec.NewAssDateForPaging()
	defer codec.DeleteAssDateForPaging(assDataForPaging)

	cellItem := codec.NewRecommandCellItem()
	defer codec.DeleteRecommandCellItem(cellItem)

	cellItem.SetTimeStayedInCell(uint16(444))
	cellItem.SetTimeStayedPrst(true)

	nrcgi := codec.NewNrCgi()
	defer codec.DeleteNrCgi(nrcgi)

	nrcgi.SetPlmn(&plmnid[0])
	nrCGI1 := []byte("12345")
	nrcgi.SetNrCellId(&nrCGI1[0])

	cellItem.SetNrcgi(nrcgi)

	cellList := codec.NewRecommandCellItemVector()
	defer codec.DeleteRecommandCellItemVector(cellList)
	cellList.Add(cellItem)
	assDataForPaging.SetRecommandCellList(cellList)
	assDataForPaging.SetRecommandCellListPrst(true)

	pagingAttemptInfo := codec.NewPagingAttemptInfo()
	defer codec.DeletePagingAttemptInfo(pagingAttemptInfo)
	pagingAttemptInfo.SetCount(3)
	pagingAttemptInfo.SetIntendedNum(2)
	pagingAttemptInfo.SetNextAreaScope(1)
	pagingAttemptInfo.SetNextAreaScopePrst(true)

	assDataForPaging.SetPagingAttemptInfo(pagingAttemptInfo)
	assDataForPaging.SetPagingAttemptInfoPrst(true)

	msg.SetAssDateForPaging(assDataForPaging)

	fmt.Println("--------------------")
	fmt.Println(msg.ToString())
	fmt.Println("--------------------")

	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PagingDecode(ossCtxtPtr codec.NgapOssCtxt) {
	//encMsg := []byte{0, 24, 64, 102, 0, 0, 6, 0, 11, 64, 15, 96, 4, 0, 49, 48, 48, 49, 50, 51, 52, 48, 1, 188, 66, 10, 0, 51, 64, 1, 0, 0, 118, 64, 44, 32, 42, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 102, 111, 114, 46, 112, 97, 103, 105, 110, 103, 46, 49, 50, 51, 52, 53, 54, 32, 69, 117, 116, 114, 97, 0, 103, 64, 7, 0, 49, 48, 48, 49, 48, 50, 0, 50, 64, 1, 64, 0, 115, 64, 7, 2, 48, 96, 54, 54, 54, 54}
	encMsg := []byte{0, 24, 64, 102, 0, 0, 6, 0, 115, 64, 7, 22, 78, 224, 54, 54, 54, 54, 0, 50, 64, 1, 64, 0, 103, 64, 7, 0, 49, 48, 48, 49, 48, 50, 0, 118, 64, 44, 32, 42, 117, 101, 46, 114, 97, 100, 105, 111, 46, 99, 97, 112, 98, 105, 108, 105, 116, 121, 46, 102, 111, 114, 46, 112, 97, 103, 105, 110, 103, 46, 49, 50, 51, 52, 53, 54, 32, 69, 117, 116, 114, 97, 0, 11, 64, 15, 96, 4, 0, 49, 48, 48, 49, 50, 51, 52, 48, 1, 188, 66, 10, 0, 51, 64, 1, 0}

	ossCodec := codec.NewNgapCodec()
	defer codec.DeleteNgapCodec(ossCodec)
	ossCodec.SetEncBuffer(&encMsg[0], len(encMsg))
	ret := ossCodec.DecodeHeader(ossCtxtPtr)
	fmt.Println("deocde header restult:", ret)

	msgType := ossCodec.GetMsgType()
	procCode := ossCodec.GetProcedureCode()
	fmt.Printf("msg type: %d, procCode: %d\n", msgType, procCode)

	pagingCodec := codec.NewPagingCodec()
	defer codec.DeletePagingCodec(pagingCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := pagingCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(pagingCodec.ToString())
	fmt.Println("--------------------")
}

func PduSessResRelRespTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {
	transfer := codec.NewPduSessResRelRespTransferCodec()
	defer codec.DeletePduSessResRelRespTransferCodec(transfer)

	secondaryRATUsageInformation := codec.NewSecRatUsageInformation()
	defer codec.DeleteSecRatUsageInformation(secondaryRATUsageInformation)
	//secondaryRATUsageInformation.SetPduSessUsageReportPrst(true)
	//pduSessUsageRprt := codec.NewPduSessUsageReport()
	//defer codec.DeletePduSessUsageReport(pduSessUsageRprt)
	//pduSessUsageRprt.SetRatType(byte(1))
	//
	//timeRprtList := codec.NewVolumeTimeReportVector()
	//defer codec.DeleteVolumeTimeReportVector(timeRprtList)
	//
	//timeRprt1 := codec.NewVolumeTimeReport()
	//defer codec.DeleteVolumeTimeReport(timeRprt1)
	//startTime1 := []byte{0x31,0x31,0x31,0x31}
	//timeRprt1.SetStartTimeStamp(&startTime1[0])
	//endTime1 := []byte{0x32,0x32,0x32,0x32}
	//timeRprt1.SetEndTimeStamp(&endTime1[0])
	//timeRprt1.SetUsageCountUL(1)
	//timeRprt1.SetUsageCountDL(2)
	//
	//timeRprtList.Add(timeRprt1)
	//
	//timeRprt2 := codec.NewVolumeTimeReport()
	//defer codec.DeleteVolumeTimeReport(timeRprt2)
	//startTime2 := []byte{0x03,0x03,0x03,0x03}
	//timeRprt2.SetStartTimeStamp(&startTime2[0])
	//endTime2 := []byte{0x04,0x04,0x04,0x04}
	//timeRprt2.SetEndTimeStamp(&endTime2[0])
	//timeRprt2.SetUsageCountUL(3)
	//timeRprt2.SetUsageCountDL(4)
	//
	//timeRprtList.Add(timeRprt2)
	//
	//pduSessUsageRprt.SetVolumeTimeReportList(timeRprtList)
	//secondaryRATUsageInformation.SetPduSessUsageReport(pduSessUsageRprt)

	secondaryRATUsageInformation.SetQosFlowUsageReportListPrst(true)
	qosFlowUsageRprtList := codec.NewQosFlowUsageReportVector()
	defer codec.DeleteQosFlowUsageReportVector(qosFlowUsageRprtList)

	qosFlowUsageRprt1 := codec.NewQosFlowUsageReport()
	defer codec.DeleteQosFlowUsageReport(qosFlowUsageRprt1)
	qosFlowUsageRprt1.SetRatType(1)
	qosFlowUsageRprt1.SetQosFlowId(6)

	qosFlowTimeRprtList := codec.NewVolumeTimeReportVector()
	defer codec.DeleteVolumeTimeReportVector(qosFlowTimeRprtList)

	timeRprt3 := codec.NewVolumeTimeReport()
	defer codec.DeleteVolumeTimeReport(timeRprt3)
	startTime3 := []byte{0x31, 0x31, 0x31, 0x31}
	timeRprt3.SetStartTimeStamp(&startTime3[0])
	endTime3 := []byte{0x32, 0x32, 0x32, 0x32}
	timeRprt3.SetEndTimeStamp(&endTime3[0])
	timeRprt3.SetUsageCountUL(1)
	timeRprt3.SetUsageCountDL(2)

	qosFlowTimeRprtList.Add(timeRprt3)

	qosFlowUsageRprt1.SetQosFlowsTimeReportList(qosFlowTimeRprtList)
	qosFlowUsageRprtList.Add(qosFlowUsageRprt1)

	secondaryRATUsageInformation.SetQosFlowUsageReportList(qosFlowUsageRprtList)

	transfer.SetSecRatUsageInfo(secondaryRATUsageInformation)

	fmt.Println("encode status:", transfer.ToString())
	msgBuffer := transfer.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func PduSessResRelRespTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{64, 0, 0, 0, 144, 64, 15, 32, 0, 200, 49, 49, 49, 49, 50, 50, 50, 50, 0, 1, 0, 2}

	transfer := codec.NewPduSessResRelRespTransferCodec()
	defer codec.DeletePduSessResRelRespTransferCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func PduSessResRelRespEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewPduSessResRelRespCodec()
	defer codec.DeletePduSessResRelRespCodec(msg)
	msg.SetRanUeNgapId(20)
	msg.SetAmfUeNgapId(25)

	//user Location Information
	userLocationInfo := codec.NewUserLocationInfo()
	defer codec.DeleteUserLocationInfo(userLocationInfo)

	nrcgi := codec.NewNrCgi()
	defer codec.DeleteNrCgi(nrcgi)
	plmnid := []byte("100")
	nrcgi.SetPlmn(&plmnid[0])
	nrCGI1 := []byte("12345")
	nrcgi.SetNrCellId(&nrCGI1[0])

	timeStamp := []byte{0x31, 0x31, 0x31, 0x31}
	userLocationInfo.SetTimeStamp(&timeStamp[0])
	userLocationInfo.SetIsTimeStampPrst(true)
	userLocationInfo.SetNRCgi(nrcgi)

	tai := codec.NewTaiType()
	defer codec.DeleteTaiType(tai)

	tai.SetPlmn(&plmnid[0])
	tac := []byte("102")
	tai.SetTac(&tac[0])

	userLocationInfo.SetTai(tai)

	msg.SetUserLocationInfo(userLocationInfo)

	relRespItem := codec.NewPduSessResRelRespItem()
	defer codec.DeletePduSessResRelRespItem(relRespItem)
	relRespItem.SetPduSessId(31)
	relRespTransfer := []byte{0}
	relRespItem.SetPduSessResRelRespTransfer(string(relRespTransfer))
	msg.AddPduSessResRelRespList(relRespItem)

	fmt.Println("encode status", msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PduSessResRelRespDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 28, 0, 47, 0, 0, 4, 0, 70, 64, 5, 0, 0, 31, 1, 0, 0, 121, 64, 19, 80, 49, 48, 48, 49, 50, 51, 52, 48, 49, 48, 48, 49, 48, 50, 49, 49, 49, 49, 0, 10, 64, 2, 0, 25, 0, 85, 64, 2, 0, 20}
	transfer := codec.NewPduSessResRelRespCodec()
	defer codec.DeletePduSessResRelRespCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func PduSessResMdfyRespEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewPduSessResModifyRespCodec()
	defer codec.DeletePduSessResModifyRespCodec(msg)
	msg.SetRanUeNgapId(20)
	msg.SetAmfUeNgapId(25)

	//user Location Information
	userLocationInfo := codec.NewUserLocationInfo()
	defer codec.DeleteUserLocationInfo(userLocationInfo)

	nrcgi := codec.NewNrCgi()
	defer codec.DeleteNrCgi(nrcgi)
	plmnid := []byte("100")
	nrcgi.SetPlmn(&plmnid[0])
	nrCGI1 := []byte("12345")
	nrcgi.SetNrCellId(&nrCGI1[0])

	timeStamp := []byte{0x31, 0x31, 0x31, 0x31}
	userLocationInfo.SetTimeStamp(&timeStamp[0])
	userLocationInfo.SetIsTimeStampPrst(true)
	userLocationInfo.SetNRCgi(nrcgi)

	tai := codec.NewTaiType()
	defer codec.DeleteTaiType(tai)

	tai.SetPlmn(&plmnid[0])
	tac := []byte("102")
	tai.SetTac(&tac[0])

	userLocationInfo.SetTai(tai)

	psCellInfo := codec.NewNrCgi()
	defer codec.DeleteNrCgi(psCellInfo)
	plmnid1 := []byte("321")
	psCellInfo.SetPlmn(&plmnid1[0])
	nrCGI2 := []byte("12345")
	psCellInfo.SetNrCellId(&nrCGI2[0])
	userLocationInfo.SetPsCellInfo(psCellInfo)
	userLocationInfo.SetPsCellInfoPrst(true)

	msg.SetUserLocationInfo(userLocationInfo)

	mdfyRespItem := codec.NewPduSessResModifyRespItem()
	defer codec.DeletePduSessResModifyRespItem(mdfyRespItem)
	mdfyRespItem.SetPduSessId(35)
	mdfyRespTransfer := []byte{120, 3, 224, 10, 180, 9, 242, 210, 4, 0, 0, 1, 240, 10, 180, 9, 243, 211, 4, 0, 0, 0, 10, 0, 15, 128, 10, 180, 9, 244, 212, 4, 0, 0, 1, 1, 64}
	mdfyRespItem.SetPduSessResMdfyRespTransfer(string(mdfyRespTransfer))
	msg.AddPduSessResModifyRespList(mdfyRespItem)

	mdfyFailRespItem := codec.NewPduSessResFailedMdfyRespItem()
	defer codec.DeletePduSessResFailedMdfyRespItem(mdfyFailRespItem)
	mdfyFailRespItem.SetPduSessId(31)
	mdfyFailRespTransfer := []byte{0, 8}
	mdfyFailRespItem.SetPduSessResFailedMdfyRespTransfer(string(mdfyFailRespTransfer))
	msg.AddPduSessResFailedModifyList(mdfyFailRespItem)

	fmt.Println("encode status", msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func PduSessResMdfyRespDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 26, 0, 104, 0, 0, 5, 0, 10, 64, 2, 0, 25, 0, 85, 64, 2, 0, 20, 0, 65, 64, 41, 0, 0, 35, 37, 120, 3, 224, 10, 180, 9, 242, 210, 4, 0, 0, 1, 240, 10, 180, 9, 243, 211, 4, 0, 0, 0, 10, 0, 15, 128, 10, 180, 9, 244, 212, 4, 0, 0, 1, 1, 64, 0, 54, 64, 6, 0, 0, 31, 2, 0, 8, 0, 121, 64, 30, 72, 49, 48, 48, 49, 50, 51, 52, 48, 49, 48, 48, 49, 48, 50, 0, 0, 0, 149, 64, 9, 0, 51, 50, 49, 49, 50, 51, 52, 48}
	transfer := codec.NewPduSessResModifyRespCodec()
	defer codec.DeletePduSessResModifyRespCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func PduSessResMdfyRespTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {
	transfer := codec.NewPduSessResMdfyRespTransferCodec()
	defer codec.DeletePduSessResMdfyRespTransferCodec(transfer)

	dlGtpTunnel := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(dlGtpTunnel)

	dlGtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2

	ipu32 := utils.Ip2long(net.ParseIP("10.180.9.242"))
	dlIpaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(dlIpaddr, ipu32)
	dlGtpTunnel.SetTransportLayerAddr(&dlIpaddr[0])

	dlTeid := make([]byte, 4)
	binary.LittleEndian.PutUint32(dlTeid, 1234)
	dlGtpTunnel.SetGtpTeid(&dlTeid[0])

	transfer.SetDlGtpTunnel(dlGtpTunnel)

	ulGtpTunnel := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(ulGtpTunnel)

	ulGtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2

	ipu32a := utils.Ip2long(net.ParseIP("10.180.9.243"))
	ulIpaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(ulIpaddr, ipu32a)
	ulGtpTunnel.SetTransportLayerAddr(&ulIpaddr[0])

	ulTeid1 := make([]byte, 4)
	binary.LittleEndian.PutUint32(ulTeid1, 1235)
	ulGtpTunnel.SetGtpTeid(&ulTeid1[0])

	transfer.SetUlGtpTunnel(ulGtpTunnel)

	addQosFlowTnlInfo := codec.NewAddQosFlowPerTNLInfo()
	defer codec.DeleteAddQosFlowPerTNLInfo(addQosFlowTnlInfo)

	addGtpTunnel := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(addGtpTunnel)

	addGtpTunnel.SetIpType(0) //0 ipv4 1 ipv6 2

	ipu32b := utils.Ip2long(net.ParseIP("10.180.9.244"))
	addIpaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(addIpaddr, ipu32b)
	addGtpTunnel.SetTransportLayerAddr(&addIpaddr[0])

	addTeid1 := make([]byte, 4)
	binary.LittleEndian.PutUint32(addTeid1, 1236)
	addGtpTunnel.SetGtpTeid(&addTeid1[0])

	addQosFlowTnlInfo.SetUpTransportLayerInfo(addGtpTunnel)

	assQosFlowVector := codec.NewAssQosFlowItemVector()
	defer codec.DeleteAssQosFlowItemVector(assQosFlowVector)

	associateQosFlowItem := codec.NewAssQosFlowItem()
	defer codec.DeleteAssQosFlowItem(associateQosFlowItem)
	associateQosFlowItem.SetQosFlowInd(1)
	associateQosFlowItem.SetQosFlowMapInd(0x01)
	associateQosFlowItem.SetQosFlowMapIndPrst(true)

	assQosFlowVector.Add(associateQosFlowItem)

	addQosFlowTnlInfo.SetAssQosFlowList(assQosFlowVector)

	transfer.AddQosFlowTNLInfoList(addQosFlowTnlInfo)

	var qosFlowInd uint = 5
	transfer.AddQosFlowAddOrMdfyRespList(qosFlowInd)

	qosflow := codec.NewQosFlowCodecItem()
	defer codec.DeleteQosFlowCodecItem(qosflow)
	qosflow.SetQosFlowInd(6)
	qosflow.SetCauseType(1)
	qosflow.SetCauseValue(1)
	transfer.AddQosFlowFailedAddOrMdfyList(qosflow)

	fmt.Println(transfer.ToString())
	msgBuffer := transfer.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PduSessResMdfyRespTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{120, 3, 224, 10, 180, 9, 242, 210, 4, 0, 0, 1, 240, 10, 180, 9, 243, 211, 4, 0, 0, 0, 10, 0, 15, 128, 10, 180, 9, 244, 212, 4, 0, 0, 1, 1, 64}

	transfer := codec.NewPduSessResMdfyRespTransferCodec()
	defer codec.DeletePduSessResMdfyRespTransferCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func PduSessResMdfyFailRespTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 8}
	fmt.Println(encMsg)

	mdfyFailCodec := codec.NewPduSessResMdfyUnSuccTransferCodec()
	defer codec.DeletePduSessResMdfyUnSuccTransferCodec(mdfyFailCodec)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := mdfyFailCodec.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(mdfyFailCodec.ToString())
	fmt.Println("--------------------")
}
func PduSessResMdfyFailRespTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	pduSessResMdfyFailEncode := codec.NewPduSessResMdfyUnSuccTransferCodec()
	defer codec.DeletePduSessResMdfyUnSuccTransferCodec(pduSessResMdfyFailEncode)
	pduSessResMdfyFailEncode.SetCause(1, 1)

	fmt.Println(pduSessResMdfyFailEncode.ToString())
	msgBuf := pduSessResMdfyFailEncode.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuf)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func PduSessResMdfyReqTransferEncode(ossCtxtPtr codec.NgapOssCtxt) {
	transfer := codec.NewPduSesResModReqTransferCodec()
	defer codec.DeletePduSesResModReqTransferCodec(transfer)

	transfer.SetCmmnNtwrkInstance("cmnet.com")

	fmt.Println(transfer.ToString())
	msgBuffer := transfer.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}
func PduSessResMdfyReqTransferDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 0, 1, 0, 166, 64, 10, 9, 99, 109, 110, 101, 116, 46, 99, 111, 109}

	transfer := codec.NewPduSesResModReqTransferCodec()
	defer codec.DeletePduSesResModReqTransferCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func NgResetEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewNgResetCodec()
	defer codec.DeleteNgResetCodec(msg)
	msg.SetCauseType(1)
	msg.SetCauseValue(2)

	//Reset Type
	resetTypeInfo := codec.NewResetTypeChoice()
	defer codec.DeleteResetTypeChoice(resetTypeInfo)

	resetTypeInfo.SetNgInterfacePrst(true)
	resetTypeInfo.SetNgInterface(0)

	msg.SetResetType(resetTypeInfo)

	fmt.Println("encode status", msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func NgResetDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 20, 0, 14, 0, 0, 2, 0, 15, 64, 2, 0, 128, 0, 88, 0, 1, 0}
	transfer := codec.NewNgResetCodec()
	defer codec.DeleteNgResetCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func NgResetAcknowledgeEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewNgResetAckCodec()
	defer codec.DeleteNgResetAckCodec(msg)
	ueAssLogicalNgConnItem := codec.NewUeAssLogicalNgConn()
	defer codec.DeleteUeAssLogicalNgConn(ueAssLogicalNgConnItem)
	ueAssLogicalNgConnItem.SetAmfUeNgapIdPrst(true)
	ueAssLogicalNgConnItem.SetAmfUeNgapId(1)
	//ueAssLogicalNgConnItem.SetRanUeNgapIdPrst(true)
	//ueAssLogicalNgConnItem.SetRanUeNgapId(2)
	msg.AddUeAssLogicalNgConnList(ueAssLogicalNgConnItem)

	fmt.Println("encode status", msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func NgResetAcknowledgeDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{32, 20, 0, 12, 0, 0, 1, 0, 111, 64, 5, 1, 96, 1, 0, 2}
	transfer := codec.NewNgResetAckCodec()
	defer codec.DeleteNgResetAckCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func SnssaiEncode(ossCtxtPtr codec.NgapOssCtxt) {
	msg := codec.NewSnssaiCodec()
	defer codec.DeleteSnssaiCodec(msg)

	tempSnssai := codec.NewSNssai()
	defer codec.DeleteSNssai(tempSnssai)

	sst := byte(1)
	tempSnssai.SetSst(&sst)
	sd := []byte{123}
	tempSnssai.SetSd(&sd[0])
	tempSnssai.SetSdPresent(true)

	msg.SetSnssai(tempSnssai)

	fmt.Println("encode status", msg.ToString())
	msgBuffer := msg.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func SnssaiDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{64, 32, 123, 0, 0}
	transfer := codec.NewSnssaiCodec()
	defer codec.DeleteSnssaiCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func pduSessResMdfyReqEncode(ossCtxtPtr codec.NgapOssCtxt) {
	//for encode test
	pduSessResMdfyReqEncode := codec.NewPduSessResModReqCodec()
	defer codec.DeletePduSessResModReqCodec(pduSessResMdfyReqEncode)

	pduSessResMdfyReqEncode.SetRanUeNgapId(20)
	pduSessResMdfyReqEncode.SetAmfUeNgapId(25)
	mdfyReqItem := codec.NewPduSessResModReqItem()
	defer codec.DeletePduSessResModReqItem(mdfyReqItem)
	mdfyReqItem.SetPduSessId(5)
	mdfyReqTransfer := []byte{0, 0, 1, 0, 166, 64, 10, 9, 99, 109, 110, 101, 116, 46, 99, 111, 109}
	mdfyReqItem.SetPduSessResModReqTransfer(string(mdfyReqTransfer))

	mdfyReqItem.SetSnssaiPrst(true)

	snssai1 := codec.NewSNssai()
	defer codec.DeleteSNssai(snssai1)
	sst3 := byte(1)
	snssai1.SetSst(&sst3)
	sd3 := []byte{123}
	snssai1.SetSd(&sd3[0])
	snssai1.SetSdPresent(true)
	mdfyReqItem.SetSnssai(snssai1)

	//mdfySnssai := []byte{64,32,123,0,0}
	//mdfyReqItem.SetSnssai(string(mdfySnssai))
	pduSessResMdfyReqEncode.AddPduSessResModReqList(mdfyReqItem)

	fmt.Println(pduSessResMdfyReqEncode.ToString())
	msgBuf := pduSessResMdfyReqEncode.Encode(ossCtxtPtr)
	defer codec.DeleteMsgBuffer(msgBuf)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	fmt.Println("encode message Buffer: ", encodeBuffer)
}

func pduSessResMdfyReqDecode(ossCtxtPtr codec.NgapOssCtxt) {
	encMsg := []byte{0, 26, 0, 51, 0, 0, 3, 0, 10, 0, 2, 0, 25, 0, 85, 0, 2, 0, 20, 0, 64, 0, 32, 0, 32, 5, 17, 0, 0, 1, 0, 166, 64, 10, 9, 99, 109, 110, 101, 116, 46, 99, 111, 109, 0, 0, 0, 148, 0, 5, 64, 32, 123, 0, 0}
	transfer := codec.NewPduSessResModReqCodec()
	defer codec.DeletePduSessResModReqCodec(transfer)
	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)
	msgBuffer.SetLength(int64(len(encMsg)))
	msgBuffer.SetValue(&encMsg[0])

	rt := transfer.Decode(ossCtxtPtr, msgBuffer)
	fmt.Println("Result: ", rt)
	fmt.Println("--------------------")
	fmt.Println(transfer.ToString())
	fmt.Println("--------------------")
}

func run() {
	ossCtxt := codec.NewOssCtxt()
	defer codec.DeleteOssCtxt(ossCtxt)

	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	ossCtxt.GetOssCtxtPtr_m()

	for i := 0; i < 1; i++ {

		NgSetupRequestEncode(ossCtxtPtr)
		NgSetupRequestDecode(ossCtxtPtr)
		NgSetupRespEncode(ossCtxtPtr)
		NgSetupRespDecode(ossCtxtPtr)
		NgSetupFailureEncode(ossCtxtPtr)
		NgSetupFailureDecode(ossCtxtPtr)
		InitialUeMsgEncode(ossCtxtPtr)
		InitialUeMsgDecode(ossCtxtPtr)
		DownlinkNasTransportEncode(ossCtxtPtr)
		DownlinkNasTransportDecode(ossCtxtPtr)
		UplinkNasTransportEncode(ossCtxtPtr)
		UplinkNasTransportDecode(ossCtxtPtr)

		PduSessionResourceSetupRequestEncode(ossCtxtPtr)
		PduSessionResourceSetupRequestDecode(ossCtxtPtr)

		PduSessResSetupReqTransferEncode(ossCtxtPtr)
		PduSessResSetupReqTransferDecode(ossCtxtPtr)

		PduSessResSetupRespTransferEncode(ossCtxtPtr)
		PduSessResSetupRespTransferDecode(ossCtxtPtr)

		CauseEncode(ossCtxtPtr)
		CauseDecode(ossCtxtPtr)

		PduSessResSetupRespEncode(ossCtxtPtr)
		PduSessResSetupRespDecode(ossCtxtPtr)

		InitialContextSetupRequestEncode(ossCtxtPtr)
		InitialContextSetupRequestDecode(ossCtxtPtr)

		PduSessionResourceSetupRequestEncode(ossCtxtPtr)
		PduSessionResourceSetupRequestDecode(ossCtxtPtr)

		PduSessResSetupRespTransferEncode(ossCtxtPtr)
		PduSessResSetupRespTransferDecode(ossCtxtPtr)

		CauseEncode(ossCtxtPtr)
		CauseDecode(ossCtxtPtr)

		PduSessResSetupRespEncode(ossCtxtPtr)
		PduSessResSetupRespDecode(ossCtxtPtr)

		InitialContextSetupRequestEncode(ossCtxtPtr)
		InitialContextSetupRequestDecode(ossCtxtPtr)

		InitialContextSetupResponseEncode(ossCtxtPtr)
		InitialContextSetupResponseDecode(ossCtxtPtr)

		InitialContextFailToSetupEncode(ossCtxtPtr)
		InitialContextFailToSetupDecode(ossCtxtPtr)

		// Ue Context Release Command
		UeContextReleaseCommandEncode(ossCtxtPtr)
		UeContextReleaseCommandDecode(ossCtxtPtr)
		UeContextReleaseCommandTransferEncode(ossCtxtPtr)
		UeContextReleaseCommandTransferDecode(ossCtxtPtr)

		// Ue Context Release Complete
		UeContextReleaseCompleteEncode(ossCtxtPtr)
		UeContextReleaseCompleteDecode(ossCtxtPtr)

		//Ue Context Release Request
		UeContextReleaseRequestEncode(ossCtxtPtr)
		UeContextReleaseRequestDecode(ossCtxtPtr)

		//Ue Radio Capability Info Indication
		UeRadioCapaInfoIndicationEncode(ossCtxtPtr)
		UeRadioCapaInfoIndicationDecode(ossCtxtPtr)

		//Pdu Session Resource Release Command
		PduSessResRelCmdTransferEncode(ossCtxtPtr)
		PduSessResRelCmdTransferDecode(ossCtxtPtr)
		PduSessResRelCmdEncode(ossCtxtPtr)
		PduSessResRelCmdDecode(ossCtxtPtr)

		//Pdu Session Resource Release Response
		PduSessResRelRespEncode(ossCtxtPtr)
		PduSessResRelRespDecode(ossCtxtPtr)

		//Paging
		PagingEncode(ossCtxtPtr)
		PagingDecode(ossCtxtPtr)

		//Error Indication
		ErrorIndicationEncode(ossCtxtPtr)
		ErrorIndicationDecode(ossCtxtPtr)

		//Pdu Session Resource Release Response Transfer Encode
		PduSessResRelRespTransferEncode(ossCtxtPtr)
		PduSessResRelRespTransferDecode(ossCtxtPtr)
		PduSessResRelRespEncode(ossCtxtPtr)
		PduSessResRelRespDecode(ossCtxtPtr)

		//ngReset
		NgResetEncode(ossCtxtPtr)
		NgResetDecode(ossCtxtPtr)

		//ngResetAcknowledge
		NgResetAcknowledgeEncode(ossCtxtPtr)
		NgResetAcknowledgeDecode(ossCtxtPtr)

		InitialContextSetupResponseEncode(ossCtxtPtr)
		InitialContextSetupResponseDecode(ossCtxtPtr)

		InitialContextFailToSetupEncode(ossCtxtPtr)
		InitialContextFailToSetupDecode(ossCtxtPtr)

		// Ue Context Release Command
		UeContextReleaseCommandEncode(ossCtxtPtr)
		UeContextReleaseCommandDecode(ossCtxtPtr)
		UeContextReleaseCommandTransferEncode(ossCtxtPtr)
		UeContextReleaseCommandTransferDecode(ossCtxtPtr)

		// Ue Context Release Complete
		UeContextReleaseCompleteEncode(ossCtxtPtr)
		UeContextReleaseCompleteDecode(ossCtxtPtr)

		//Ue Context Release Request
		UeContextReleaseRequestEncode(ossCtxtPtr)
		UeContextReleaseRequestDecode(ossCtxtPtr)

		//Ue Radio Capability Info Indication
		UeRadioCapaInfoIndicationEncode(ossCtxtPtr)
		UeRadioCapaInfoIndicationDecode(ossCtxtPtr)

		//Pdu Session Resource Release Command
		PduSessResRelCmdTransferEncode(ossCtxtPtr)
		PduSessResRelCmdTransferDecode(ossCtxtPtr)
		PduSessResRelCmdEncode(ossCtxtPtr)
		PduSessResRelCmdDecode(ossCtxtPtr)

		//Pdu Session Resource Release Response
		PduSessResRelRespEncode(ossCtxtPtr)
		PduSessResRelRespDecode(ossCtxtPtr)

		//Pdu Session Resource Modify Response
		PduSessResMdfyRespEncode(ossCtxtPtr)
		PduSessResMdfyRespDecode(ossCtxtPtr)

		PduSessResMdfyRespTransferEncode(ossCtxtPtr)
		PduSessResMdfyRespTransferDecode(ossCtxtPtr)

		PduSessResMdfyFailRespTransferEncode(ossCtxtPtr)
		PduSessResMdfyFailRespTransferDecode(ossCtxtPtr)

		pduSessResMdfyReqEncode(ossCtxtPtr)
		pduSessResMdfyReqDecode(ossCtxtPtr)

		//Pdu Session Resource Modify Request
		PduSessResMdfyReqTransferEncode(ossCtxtPtr)
		PduSessResMdfyReqTransferDecode(ossCtxtPtr)

		//Paging
		PagingEncode(ossCtxtPtr)
		PagingDecode(ossCtxtPtr)

		//Error Indication
		ErrorIndicationEncode(ossCtxtPtr)
		ErrorIndicationDecode(ossCtxtPtr)

		//Pdu Session Resource Release Response Transfer Encode
		PduSessResRelRespTransferEncode(ossCtxtPtr)
		PduSessResRelRespTransferDecode(ossCtxtPtr)
		PduSessResRelRespEncode(ossCtxtPtr)
		PduSessResRelRespDecode(ossCtxtPtr)

		//ngReset
		NgResetEncode(ossCtxtPtr)
		NgResetDecode(ossCtxtPtr)

		//ngResetAcknowledge
		NgResetAcknowledgeEncode(ossCtxtPtr)
		NgResetAcknowledgeDecode(ossCtxtPtr)

		SnssaiEncode(ossCtxtPtr)
		SnssaiDecode(ossCtxtPtr)

		time.Sleep(time.Millisecond)
		break
	}

}

func main() {
	run()
	//go run()

	//time.Sleep(time.Second * 3600)
}
