package smsender

import (
	"context"
	"fmt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
	"lite5gc/openapi/models"
)

func SendPduSessEstbRejectMsg(ctxt context.Context, cause nas.Sm5gCause) error {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)
	//send reject msg
	//get imsi
	imsi, ok := ctxt.Value(types.ImsiCK).(*types3gpp.Imsi)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "no sm header found")
		return fmt.Errorf("no sm header found")
	}

	//get the sm header
	header, ok := ctxt.Value(types.SmHeaderCK).(*nas.SmNasMessageHeader)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "no sm header found")
		return fmt.Errorf("no sm header found")
	}

	var rejectMsg nasmsg.PduSessionEstbRejectMsg
	rejectMsg.MsgHeader = *header
	rejectMsg.MsgHeader.MessageType = nas.PduSessEstabishReject
	rejectMsg.SMCause = cause

	//24501 8.3.3.3 Allowed SSC mode
	//This IE is included when the network rejects the PDU SESSION ESTABLISHMENT REQUEST with cause #68 "not supported SSC mode.
	if cause == nas.NotSupportedSSCMode {
		sscMode := configure.SmfConf.Service.SSCMode
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "scConfigMgr.GetSSCMode():", sscMode)
		switch sscMode {
		case nas.SSCMode1:
			rejectMsg.AllowedSSCMode.Ssc1 = true
		case nas.SSCMode2:
			rejectMsg.AllowedSSCMode.Ssc2 = true
		case nas.SSCMode3:
			rejectMsg.AllowedSSCMode.Ssc3 = true
		}
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "rejectMsg.AllowedSSCMode:", rejectMsg.AllowedSSCMode)
		rejectMsg.IeFlags.Set(nasmsg.Ieid_PduSessionEstbReject_AllowedSSCMode)
	}

	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "session reject cause:", rejectMsg.SMCause)
	rejectBytes, err := rejectMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "fail to encode session reject msg, err:(%s)", err)
		return fmt.Errorf("fail to encode session reject msg, err:(%s)", err)
	}

	//msg counter
	pm.PegCounter(statistics.PduSessEstablishRejectCounter)

	var msgData n11msg.N1N2MessageTransferReqData
	N1Container := n11msg.N1MessageContainerIE{N1MsgClass: n11msg.SM_N1Info, N1MessageContent: rejectBytes}

	// set n1 to msg data
	msgData.N1MessageContainer = &N1Container
	msgData.IeFlags.Set(n11msg.Ieid_n1MessageContainer)

	// set psi to msg data
	msgData.SessionId = header.PduSessionID
	msgData.IeFlags.Set(n11msg.Ieid_pdusessionId)

	// set n2
	N2Container := n11msg.N2InfoContainerIE{N2InforClass: n11msg.PWS_RF}
	msgData.N2InfoContainer = &N2Container
	msgData.IeFlags.Set(n11msg.Ieid_n2InfoContainer)

	requestData := &models.N1N2MessageTransferRequest{}
	requestData.JsonData = sbicmn.Trans_N11ToModels_N1N2MsgTransferReqFormat(&msgData)
	requestData.BinaryDataN1Message = msgData.N1MessageContainer.N1MessageContent
	SendN1N2MsgTransfer(ctxt, requestData, imsi.AddIMSIPrefix(), header.PduSessionID, 1)

	return nil
}
