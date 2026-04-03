package sbilayer

import (
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/sc/n11layer"
	"lite5gc/smf/sc/statemgr"
	"lite5gc/smf/smfcontext/gctxt"
	"strings"
)

func HandleSBIMsg(ctx context.Context, msg *router.DataMsg) error {
	/// TO DO: Handle msg from cmdcli module
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)

	message, ok := msg.MsgData.(*sbicmn.SbiMessage)
	if !ok {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "invalid IPC message - Sbi Message")
		return fmt.Errorf("invalid IPC message - SBI2ScMsg")
	}

	//dispatch message here
	switch message.MsgType {
	case sbicmn.GetSmDataMsgResponse:
		rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "received GetSmDataResponse")
		smData := message.MsgData.(*sbicmn.SbiGetSmDataMsg)
		//get the ueCtxt from ctxt
		imsistr := strings.TrimPrefix(smData.Supi, "imsi-")

		var imsi types3gpp.Imsi
		imsi.StoreImsiString(imsistr, types3gpp.CheckMncLen(imsistr))

		ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
			return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
		}

		amsd := sbicmn.Trans_ModelsToN11_SMDataFormat(*smData.Data)
		ueCtxt.SessMgntSubsDataMap = amsd

		// pdu session context
		pduSessCtxt, ok := ueCtxt.PduSessCtxts[smData.Psi]
		if !ok {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to get pdu sess ctxt")
			return fmt.Errorf("get pdu session context")
		}

		ctx = context.WithValue(ctx, types.SmfPduSessCtxtCK, pduSessCtxt)
		ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

		err = statemgr.TriggerSmfFsm(ctx,
			statemgr.SessionESTB,
			statetype.StatePduSessEstbWfSmGet,
			statetype.EventPduSessEstbSmGetAck)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	case sbicmn.PduSessCreateSMContextReq:
		//sc handle pdu session sm context create message
		rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "received PDUSessionSMContextCreate")

		smData := message.MsgData.(*sbicmn.SbiHandlerMessage)

		err := n11layer.HandlePduSessCreateSmCtxtRequest(ctx, smData)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to handle PDUSessionSMContextCreate, err(%s)", err)
		}
	case sbicmn.N1N2MessageTransferResp:
		// sc handle n1n2 message transfer response message
		rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "received N1N2MessageTransfer")

		msgData := message.MsgData.(*sbicmn.SbiPostN1N2MsgTransferMsg)

		err := n11layer.HandlePostN1N2MsgTransferResp(ctx, msgData)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to handle N1N2MessageTransfer, err(%s)", err)
		}
	case sbicmn.PduSessUpdateSMContextReq:
		rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "received update sm context request message")

		msg := message.MsgData.(*sbicmn.SbiHandlerMessage)

		err := n11layer.HandlePduSessSMContextUpdateReq(ctx, msg)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to handle PDUSessionSMContextUpdate, err(%s)", err)
		}
	case sbicmn.PduSessReleaseSMContextReq:
		rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "received pdu sess release request")
		msg := message.MsgData.(*sbicmn.SbiHandlerMessage)

		err := n11layer.HandlePduSessSMContextReleaseReq(ctx, msg)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to handle PDUSessionSMContextUpdate, err(%s)", err)
		}
	}
	return nil
}
