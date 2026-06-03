package ngapsender

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/utils"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	t3 "lite5gc/cmn/types3gpp"
)

func (p *NgapSender) SendPduSessResModRequest(ueCtxt *gctxt.UeContext,
	nasmsg []byte,
	pduResReqTransfer string,
	psi nas.PduSessID) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	msg := ngapmsg.NewPduSessResModifyReqMsg()
	msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	msg.AmfUeNgApId = ueCtxt.GetAmfUeNgapId()
	msg.RanUeNgApId = ueCtxt.GetRanUeNgapId()
	pduSessResSetupReqItem := &t3.PduSessResModifyReqItem{}
	pduSessResSetupReqItem.PduSessionId = uint8(psi)
	pduSessResSetupReqItem.IsNasPduPrst = true
	pduSessResSetupReqItem.NasPdu = string(nasmsg)
	pduSessResSetupReqItem.PduSessResModReqTrans = pduResReqTransfer
	msg.AddPduSessResModifyReqItem(pduSessResSetupReqItem)

	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = msg.Encode() //encode ngap message
	msgBuf.MsgLen = len(msgBuf.Buffer)
	ngapInstId, err := utils.GetNgapInstId(ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt, "failed to get ngap instance id, error(%s)", err)
		return err
	}

	err = p.SendNgapMsg(ngapInstId, msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap,rlogger.ERROR,ueCtxt.GetImsiPtr(),types.ErrFailSendNgapMsg)
		return types.ErrFailSendNgapMsg
	}

	return nil
}
