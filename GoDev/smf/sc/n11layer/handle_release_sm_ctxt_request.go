package n11layer

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi/models"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/sc/sessmgnt/smutility"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func HandlePduSessSMContextReleaseReq(ctxt context.Context, msgData *sbicmn.SbiHandlerMessage) error {
	rlogger.FuncEntry(types.ModuleSmfN11, msgData)
	if msgData != nil {
		smfCtxtRef := msgData.HTTPRequest.Params["smContextRef"]

		imsi, psi, err := nas.GetSmfKeys(smfCtxtRef)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil,
				"failed to get imsi and psi from url(%s)", smfCtxtRef)
			return fmt.Errorf("failed to get imsi and psi from url(%s)", smfCtxtRef)
		}

		ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
		if err != nil {
			// no Ue context, create a new ue context
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil,
				"failed to get ue contexe with imsi(%s)", imsi.String())
			return types.ErrFailFindUeCtxt
		}
		ctxt = context.WithValue(ctxt, types.UeContextCK, ueCtxt)

		pduSessCtxt, ok := ueCtxt.PduSessCtxts[psi]
		if !ok {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get pdu session context with psi(%d)", psi)
			return types.ErrFailFindSessionCtxt
		}
		ctxt = context.WithValue(ctxt, types.SmfPduSessCtxtCK, pduSessCtxt)

		modleSmCtxtReleaseData := msgData.HTTPRequest.Body.(*models.ReleaseSmContextRequest).JsonData
		if modleSmCtxtReleaseData == nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(), "invalid modelSmCtxtReleasedata")
		}

		smCtxtReleaseRequestData := sbicmn.Trans_ModelsToN11_SmContextReleaseDataFormat(modleSmCtxtReleaseData)
		ctxt = context.WithValue(ctxt, types.ReleaseSMContextRequestDataCK, smCtxtReleaseRequestData)
		rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"release cause(%s)", smCtxtReleaseRequestData.NgApCause.String())

		// send session delete request to upf
		n4Ctxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(pduSessCtxt.SEID))
		if err != nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, n4Ctxt, "Failure to get N4 Context:%s", err)

			//release resource, delete  session ctxt in smf
			err = smutility.ReleasePduSessCtxt(ctxt, pduSessCtxt.PduSessionId)
			if err != nil {
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, n4Ctxt, "input para is nil")
			}
			rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, ueCtxt.GetImsiPtr(),
				"release ue resource in SMF side,imsi:%d,psi:%d for GnbSctpShutdownEvent ",
				imsi.String(), pduSessCtxt.PduSessionId)

			// there are no n4 context, send response msg to amf directly
			smsender.SendSbiRespMsgReleaseSMContextResponse(msgData)

		} else {
			// send fpcp session deletion request to upf
			prcdCtxt := prcdctxt.NewPduSessionRelReqPrcdCtxt(pduSessCtxt.PduSessionId)
			prcdCtxt.SbiMessage = msgData
			prcdCtxt.IsDeRegisterPrcd = true
			prcdCtxt.Imsi = ueCtxt.IMSI
			pduSessCtxt.SetPrcdCtxt(prcdCtxt)

			scN4MsgData := &gctxt.ScN4MsgData{}
			err = smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, scN4MsgData, pfcp.PFCP_Session_Deletion_Request)
			if err != nil {
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"fail to send msg to pfcp node")
				return fmt.Errorf("fail to send msg to pfcp node")
			}
		}
	} else {
		rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}

	return nil
}
