package ngapsender

import (
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func (p *NgapSender) SendErrorIndicationMsg(gnbInfo *types3gpp.GnbInfo,
	ctype types3gpp.CauseType,
	cvalue types3gpp.CauseValue) {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	errorIndicationMsg := ngapmsg.ErrIndicationMsg{}
	errorIndicationMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())

	errorIndicationMsg.CauseType = uint16(ctype)
	errorIndicationMsg.CauseValue = uint16(cvalue)
	errorIndicationMsg.IsCausePrst = true

	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = errorIndicationMsg.Encode() //encode ngap message
	msgBuf.MsgLen = len(msgBuf.Buffer)
	ngapInstId := gnbInfo.GnbInstId
	err := p.SendNgapMsg(ngapInstId, msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap,rlogger.ERROR,nil,types.ErrFailSendNgapMsg)
	}
}
