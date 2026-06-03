package ngaplayer

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *LayerMgr) handlePduSessResSetupResponseMsg(ctx context.Context, msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	message := ngapmsg.NewPduSessResSetupRespMsg()
	message.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := message.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to decode pdu session resource setup response message")
		return err
	}

	uectxt, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(message.AmfUeNGAPId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to find the ue context with AmfUeNgApId(%d)", message.AmfUeNGAPId)
		return err
	}

	var psi nas.PduSessID

	updateSmCtxtReq := n11msg.UpdateSMContextRequestData{}

	//success case
	if len(message.PduSessResSetupRespList) > 0 {
		// set the first item in update sm ctxt request message todo
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, []interface{}{nil, uectxt}, "add PduResModRsp N2SmInfo")
		respItem := message.PduSessResSetupRespList[0]
		updateSmCtxtReq.N2SmInfo = respItem.PduSessResSetupRespTransfer
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfo)
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, []interface{}{nil, uectxt},
			"respItem.PduSessResSetupRespTransfer(%v)", respItem.PduSessResSetupRespTransfer)

		updateSmCtxtReq.N2SmInfoType = n11msg.PduResSetupRsp
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfoType)

		updateSmCtxtReq.UpCnxState = n11msg.ACTIVATED
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_upCnxState)

		psi = nas.PduSessID(respItem.PduSessionId)
	}
	//failure case
	if len(message.PduSessResFailedToSetupList) > 0 {
		// set the first item todo
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, []interface{}{nil, uectxt},
			"add PduResModFail N2SmInfo")

		failedItem := message.PduSessResFailedToSetupList[0]
		updateSmCtxtReq.N2SmInfo = failedItem.UnsuccessTransfer
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfo)

		updateSmCtxtReq.N2SmInfoType = n11msg.PduResSetupFail
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfoType)

		updateSmCtxtReq.UpCnxState = n11msg.DEACTIVATED
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_upCnxState)

		psi = nas.PduSessID(failedItem.PduSessionId)
	}

	if len(message.PduSessResSetupRespList) > 0 || len(message.PduSessResFailedToSetupList) > 0 {
		mmsender.SendUpdateSMCtxtSBIMsg(ctx, psi, uectxt.GetImsi(), updateSmCtxtReq)
	} else {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, []interface{}{nil, uectxt},
			"nothing to send to SMF in PduSessResSetupResponseMsg")
	}

	return nil
}
