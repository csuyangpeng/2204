package ngapsender

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
)

func (p *NgapSender) SendUEContextReleaseCommand(ueCtxt *gctxt.UeContext,
	gnbInstId uint32,
	relCause types3gpp.CauseValue) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, ueCtxt)

	ueCtxtRelCmdMsg := ngapmsg.NewUeContextReleaseCmdMsg()
	ueCtxtRelCmdMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())

	ueCtxtRelCmdMsg.UeNgapIdType = types3gpp.UeNgApIdPairType
	ueCtxtRelCmdMsg.AmfNgapId = ueCtxt.GetAmfUeNgapId()
	ueCtxtRelCmdMsg.RanNgapId = ueCtxt.GetRanUeNgapId()
	ueCtxtRelCmdMsg.Cause.Type = types3gpp.CT_Nas
	ueCtxtRelCmdMsg.Cause.Value = relCause

	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = ueCtxtRelCmdMsg.Encode() //encode ngap message
	msgBuf.MsgLen = len(msgBuf.Buffer)

	err := p.SendNgapMsg(gnbInstId, msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap,rlogger.ERROR,ueCtxt.GetImsiPtr(),types.ErrFailSendNgapMsg)
		return types.ErrFailSendNgapMsg
	}
	//msg counter
	pm.PegCounter(statistics.UEContextReleaseCommandCounter)

	return nil
}

func (p *NgapSender) SendUEContextReleaseCommandSim(n2Conn *gctxt.N2ConnCtxt, relCause types3gpp.CauseValue) error {

	ueCtxtRelCmdMsg := ngapmsg.NewUeContextReleaseCmdMsg()
	ueCtxtRelCmdMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())

	ueCtxtRelCmdMsg.UeNgapIdType = types3gpp.UeNgApIdPairType
	ueCtxtRelCmdMsg.AmfNgapId = uint64(n2Conn.AmfUeNgapID)
	ueCtxtRelCmdMsg.RanNgapId = n2Conn.GnbConnID
	ueCtxtRelCmdMsg.Cause.Type = types3gpp.CT_Nas
	ueCtxtRelCmdMsg.Cause.Value = relCause

	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = ueCtxtRelCmdMsg.Encode() //encode ngap message
	msgBuf.MsgLen = len(msgBuf.Buffer)

	ngapInstId := n2Conn.GnbInfo.GnbInstId
	err := p.SendNgapMsg(ngapInstId, msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap,rlogger.ERROR,nil,types.ErrFailSendNgapMsg)
		return types.ErrFailSendNgapMsg
	}
	//msg counter
	pm.PegCounter(statistics.UEContextReleaseCommandCounter)

	return nil
}
