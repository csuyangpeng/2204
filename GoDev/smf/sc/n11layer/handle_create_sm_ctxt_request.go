package n11layer

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi/models"
	"lite5gc/smf/sc/naslayer"
	"lite5gc/smf/smfcontext/gctxt"
)

func HandlePduSessCreateSmCtxtRequest(ctxt context.Context, sbimsg *sbicmn.SbiHandlerMessage) error {
	rlogger.FuncEntry(types.ModuleSmfN11, nil)
	if sbimsg != nil{
		// get msgData payload : SmContextCreateData
		sbiSmCtxtCreateData := sbimsg.HTTPRequest.Body.(*models.PostSmContextsRequest).JsonData

		//transfer sbi type into interal data struct
		smCtxtCreateData := sbicmn.Trans_ModelsToN11_SmContextCreateDataFormat(sbiSmCtxtCreateData)

		var ueCtxt *gctxt.UeContext
		var err error

		// if OPTIONAL IE : supi is exist
		if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_supi) {
			// Check if this UE has UeContext
			imsi := smCtxtCreateData.Supi.GetImsi()
			ueCtxt, err = gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
			if err != nil {
				// no Ue context, create a new ue context
				ueCtxt = gctxt.NewUeContext(imsi)
				// add index(imsi) for UeContext
				err = gctxt.AddIndexUeContext(gctxt.ImsiKey(imsi.GetValue()), ueCtxt)
				if err != nil {
					rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
						"failed to add index(imsi:%s) for ue context,error(%s)", imsi, err)
					return fmt.Errorf("failed to add index(imsi:%s) for ue context, error(%s)", imsi, err)
				}
			}
		} else {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "no SUPI in create sm ctxt request")
			// todo
			// for emergency call, no supi in this message
			return types.ErrFailGetImsi
		}

		//save the message in ctxt
		ctxt = context.WithValue(ctxt, types.SmfN11MsgDataCK, &smCtxtCreateData)
		ctxt = context.WithValue(ctxt, types.UeContextCK, ueCtxt)
		ctxt = context.WithValue(ctxt, types.SbiMsgCK, sbimsg)
		ctxt = context.WithValue(ctxt, types.ImsiCK, smCtxtCreateData.Supi.GetImsi())

		//handle nas msg :UL NAS transport msg and PDU session establishment request msg
		if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_n1SmMsg) {
			nasData := smCtxtCreateData.N1SmMsg

			nasLayer, ok := ctxt.Value(types.SmfNasLayerCK).(*naslayer.NasMgr)
			if !ok {
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"fail to get nas layer mgr")
				return fmt.Errorf("fail to get nas layer mgr")
			}
			err = nasLayer.HandleIncomingNasMsg(ctxt, nasData)
			if err != nil {
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"nas layer failed to handle n1 sm message")
				return fmt.Errorf("nas layer failed to handle n1 sm message")
			}
		} else {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "no n1 sm message in create sm context message")
		}
	} else {
		rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	return nil
}
