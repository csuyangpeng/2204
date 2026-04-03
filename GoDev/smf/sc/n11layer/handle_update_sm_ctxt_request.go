package n11layer

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/openapi/models"
	"lite5gc/smf/sc/naslayer"
	"lite5gc/smf/sc/sessmgnt/procedure"
	"lite5gc/smf/sc/statemgr"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func HandlePduSessSMContextUpdateReq(ctxt context.Context, msgData *sbicmn.SbiHandlerMessage) error {
	rlogger.FuncEntry(types.ModuleSmfN11, msgData)
	if  msgData != nil{
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

		pduSessCtxt, ok := ueCtxt.PduSessCtxts[psi]
		if !ok {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get pdu session context with psi(%d)", psi)
			return types.ErrFailFindSessionCtxt
		}

		//modleSmCtxtUpdateData, ok := msgData.HTTPRequest.Body.(*models.SmContextUpdateData)
		modleUpdateSmCtxtRequest, ok := msgData.HTTPRequest.Body.(*models.UpdateSmContextRequest)
		if !ok {
			return fmt.Errorf("failed to get model data")
		}

		smCtxtUpdateData := sbicmn.Trans_ModelsToN11_SmContextUpdateDataReq(modleUpdateSmCtxtRequest)

		ctxt = context.WithValue(ctxt, types.UeContextCK, ueCtxt)
		ctxt = context.WithValue(ctxt, types.SmfSbiHandlerMsgCK, msgData)
		ctxt = context.WithValue(ctxt, types.SmfPduSessCtxtCK, pduSessCtxt)
		ctxt = context.WithValue(ctxt, types.SmfUpdateSmCtxtReqMsgCK, &smCtxtUpdateData)

		// get procedure context
		prcdCtxt := pduSessCtxt.GetPrcdCtxt()

		// for session release request message, create the procedure context
		if smCtxtUpdateData.IeFlags.Test(n11msg.Ieid_release) && smCtxtUpdateData.Release {
			if prcdCtxt == nil {
				//session rel 第一次update
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"receive 1st update sm ctxt in session release procedure")
				relPrcdCtxt := prcdctxt.NewPduSessionRelReqPrcdCtxt(pduSessCtxt.PduSessionId)
				relPrcdCtxt.SbiMessage = msgData // save the sbi handler message
				prcdCtxt = relPrcdCtxt
				pduSessCtxt.SetPrcdCtxt(prcdCtxt)
			} else {
				//session rel 第二次update
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"receive 2nd update sm ctxt in session release procedure")
				// procedure context
				pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionRelReqPrcdCtxt)
				if !ok {
					rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ueCtxt.GetImsiPtr(),
						"prcdctxt is not PduSessionRelReqPrcdCtxt, but:",
						gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
					return fmt.Errorf("prcdctxt is not PduSessionRelReqPrcdCtxt")
				}
				pCtxt.SbiMessage = msgData
			}
		}

		if smCtxtUpdateData.IeFlags.Test(n11msg.Ieid_n1SmMsg) {
			// handle nas message
			nasData := smCtxtUpdateData.N1SmMsg
			nasLayer := ctxt.Value(types.SmfNasLayerCK).(*naslayer.NasMgr)
			err = nasLayer.HandleIncomingNasMsg(ctxt, nasData)
			if err != nil {
				return fmt.Errorf("nas layer failed to handle n1 sm message")
			}
			return nil
		} else {

			rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, ueCtxt.GetImsiPtr(),
				"no n1 sm message in create sm context message")

			switch prcdCtxt.(type) {
			case *prcdctxt.PduSessionEstbPrcdCtxt:
				rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, ueCtxt.GetImsiPtr(),
					"pdu session establish procedure")
				pCtxt := prcdCtxt.(*prcdctxt.PduSessionEstbPrcdCtxt)
				//store info in procedure context
				pCtxt.SbiMessage = msgData

				err := statemgr.TriggerSmfFsm(ctxt,
					statemgr.SessionESTB,
					pCtxt.GetCurrentState(),
					statetype.EventPduSessEstbUpdateSmCtxtReq)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to trigger fsm")
					return types.ErrFailTriggerFsm
				}
			case *prcdctxt.PduSessionRelReqPrcdCtxt:
				rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, ueCtxt.GetImsiPtr(), "pdu session release procedure")
			case *prcdctxt.PduSessionModPrcdCtxt:
				rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, ueCtxt.GetImsiPtr(), "pdu session modification procedure")
			case *prcdctxt.AnRelSerReqPrcdCtxt:
				rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
					"receive second update smcontext request.")
				// second update sm context message to smf for service request
				pCtxt := prcdCtxt.(*prcdctxt.AnRelSerReqPrcdCtxt)
				//store info in procedure context
				pCtxt.SbiMessage = msgData

				err := statemgr.TriggerSmfFsm(ctxt,
					statemgr.SerReq,
					pCtxt.GetCurrentState(),
					statetype.EventPduSessSerReqUpdateSmCtxtReqSec)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to trigger fsm")
					return types.ErrFailTriggerFsm
				}
			default:
				if smCtxtUpdateData.IeFlags.Test(n11msg.Ieid_upCnxState) {
					rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
						"receive first update sm context request")

					anRelSerReqPrcdCtxt := prcdctxt.NewAnRelSerReqPrcdCtxt(pduSessCtxt.PduSessionId)
					anRelSerReqPrcdCtxt.SbiMessage = msgData // save the sbi handler message
					anRelSerReqPrcdCtxt.UpCnxState = smCtxtUpdateData.UpCnxState
					anRelSerReqPrcdCtxt.IsAnRelease = false
					prcdCtxt = anRelSerReqPrcdCtxt
					pduSessCtxt.SetPrcdCtxt(prcdCtxt)

					switch smCtxtUpdateData.UpCnxState {
					case n11msg.ACTIVATED:
						// first update sm ctxt request message for service request procedure
						//trigger the FSM
						err := statemgr.TriggerSmfFsm(ctxt,
							statemgr.SerReq,
							statetype.StatePduSessSerReqStart,
							statetype.EventPduSessSerReqStart)
						if err != nil {
							rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
							return types.ErrFailTriggerFsm
						}
					case n11msg.DEACTIVATED:
						anRelSerReqPrcdCtxt.IsAnRelease = true
						// an rlease, update sm ctxt request message
						procedure.HandlerUpdataSmCtxtRequestMsg4AnRelease(ctxt)
					default:
						rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
							"invalid UpCnxState(%s)", smCtxtUpdateData.UpCnxState.String())
					}
				}
			}
		}
	} else {
		rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	return nil
}
