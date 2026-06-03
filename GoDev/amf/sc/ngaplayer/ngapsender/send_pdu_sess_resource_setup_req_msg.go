package ngapsender

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/amf/sc/utils"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"

	"lite5gc/cmn/types"
)

func (p *NgapSender) SendPduSessResSetupRequest(ueCtxt *gctxt.UeContext,
	nasmsg []byte,
	pduResReqTransfer string,
	s *nasie.SNssai,
	psi nas.PduSessID) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, ueCtxt)

	msg := ngapmsg.NewPduSessResSetupReqMsg()
	msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	msg.AmfUeNgApId = ueCtxt.GetAmfUeNgapId()
	msg.RanUeNgApId = ueCtxt.GetRanUeNgapId()
	pduSessResSetupReqItem := &types3gpp.PduSessResSetupReqItem{}
	pduSessResSetupReqItem.PduSessionId = uint8(psi)
	pduSessResSetupReqItem.Snssai = utils.Convert2NgapSNssai(s)
	pduSessResSetupReqItem.IsNasPduPrst = true
	pduSessResSetupReqItem.NasPdu = string(nasmsg)
	pduSessResSetupReqItem.PduSessResSetupReqTrans = pduResReqTransfer
	msg.AddPduSessResSetupReqItem(pduSessResSetupReqItem)

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
	//msg counter
	pm.PegCounter(statistics.PDUSessionResourceSetupResponseCounter)

	return nil
}
