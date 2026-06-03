package ngapsender

import (
	"lite5gc/amf/context/gctxt"
)

func (p *NgapSender) SendN2ResourceReleaseCommand(ueCtxt *gctxt.UeContext, gnbInstId uint32, msg []byte) error {
	//td := []interface{}{ueCtxt}
	//rlogger.FuncEntry(types.AmfScNgapMod,td)

	//var relcmdmsgData nasmsg.PduSessionReleaseCommandMsg
	//nasMsg := bytes.NewReader(msg)
	//rlogger.Trace(types.AmfScNgapMod, rlogger.INFO, td, "data to be decode:", msg)
	//err := relcmdmsgData.Decode(nasMsg)
	//if err != nil {
	//	rlogger.Trace(types.AmfScNgapMod, rlogger.ERROR, td, "failed to decode "+
	//		"pdu session rel cmd message, err: ", err)
	//}
	//
	//dlNasTransMsgData := &amfnasmsg.DownLinkNasTransportMsg{}
	//dlNasTransMsgData.PduSessId = relcmdmsgData.MsgHeader.PduSessionID
	//dlNasTransMsgData.OptIeBitSet.Set(amfnasmsg.Ieid_DownlinkNasTrans_PduSessId)
	//dlNasTransMsgData.PayloadType = nasie.N1SmInformation
	//dlNasTransMsgData.PayloadContainer.PayloadContainerEntry = make([]nasie.PayloadContainerENTRY, 1)
	//dlNasTransMsgData.PayloadContainer.PayloadContainerEntry[0].ContainerContents =
	//	append(dlNasTransMsgData.PayloadContainer.PayloadContainerEntry[0].ContainerContents, msg...)
	//dlNasData, err := dlNasTransMsgData.Encode()
	//if err != nil {
	//	rlogger.Trace(types.AmfScNgapMod, rlogger.ERROR, td, "failed to encode downlink nas transport message.")
	//}
	//
	////add sec nas header
	//// set security protected header
	//var dlNasMsg []byte
	//dlNasMsg, err = utils.EncodeSecPrctNasMsg(ueCtxt, nas.IntegrityPrtctCipher, dlNasData)
	//if err != nil {
	//	rlogger.Trace(types.AmfScNgapMod, rlogger.ERROR, ueCtxt, "failed to add security header")
	//	return err
	//}
	//
	//sessionRelCmdMsg := ngapmsg.NewPduSessResRelCmmdMsg()
	//sessionRelCmdMsg.SetOssCodecCtxt(p.ossCtxt)
	//sessionRelCmdMsg.PduSessResRelCmmdList = make([]*types3gpp.PduSessResRelCmmdItem, 1)
	//sessionRelCmdMsg.PduSessResRelCmmdList[0] = &types3gpp.PduSessResRelCmmdItem{}
	//sessionRelCmdMsg.PduSessResRelCmmdList[0].PduSessionId = uint8(relcmdmsgData.MsgHeader.PduSessionID)
	//
	//cmdCause := ngapsmf.NewPduSessResRelCmmdTransfer()
	//cmdCause.SetOssCodecCtxt(p.ossCtxt)
	//cmdCause.Cause.Type = types3gpp.CT_Nas
	//cmdCause.Cause.Value = types3gpp.Nas_normal_release
	//cmdCauseBytes := cmdCause.Encode()
	//
	//sessionRelCmdMsg.PduSessResRelCmmdList[0].PduSessResRelCmmdTransfer = string(cmdCauseBytes)
	//sessionRelCmdMsg.AmfUeNGAPId = ueCtxt.GetAmfUeNgapId()
	//sessionRelCmdMsg.RanUeNGAPId = ueCtxt.GetRanUeNgapId()
	//sessionRelCmdMsg.NasPdu = dlNasMsg
	//sessionRelCmdMsg.IsNasPduPrst = true
	//
	//msgBuf := types.MsgBuf{}
	//msgBuf.Buffer = sessionRelCmdMsg.Encode() //encode ngap message
	//msgBuf.MsgLen = len(msgBuf.Buffer)
	//
	//p.SendNgapMsg(gnbInstId, &msgBuf)
	//
	//nassecurity.UpdateDownlinkNasCounter(ueCtxt)

	return nil
}
