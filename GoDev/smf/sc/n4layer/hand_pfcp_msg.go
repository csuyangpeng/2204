package n4layer

import (
	"context"
	"fmt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/oam/pm"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/sc/statemgr"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func HandleNodeSessionMsg(ctxt context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModuleSmfN11, ctxt)
	if msg != nil {
		//new a context for message handler
		type key = struct{}
		ctxt_new := context.WithValue(ctxt, key{}, "start handle message")

		msgData := msg.MsgData.(pfcpv1.ServiceMsg)

		//get pdu session context with seid
		var pduSessCtxt *gctxt.PduSessContext
		pduSessCtxt, err := gctxt.GetSessContext(gctxt.SeidKey(msgData.Msg.Cxt.SEID))
		if err != nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "no match session context or no session context, sessionCtxt == nil")
			return fmt.Errorf("no match session context")
		}

		cause := msgData.Msg.Cxt.Cause
		rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, msgData, "msgData.Msg.ID", msgData.Msg.ID)
		rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, msgData, "sessionCtxt.Cause", cause.Error())

		switch msgData.Msg.ID {
		case pfcp.PFCP_Session_Establishment_Response:
			//if sessionCtxt.GetPrcdCtxt() != nil {
			pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
			if !ok {
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ctxt, "fail to get session estb prcd cxt")
				return fmt.Errorf("fail to get session estb prcd cxt")
			}
			if cause == gctxt.Cause_Request_accepted {
				ctxt = context.WithValue(ctxt, types.PduSessionEstbPrcdCtxtCK, pCtxt)
				ctxt = context.WithValue(ctxt, types.SmfPduSessCtxtCK, pduSessCtxt)

				err = statemgr.TriggerSmfFsm(ctxt,
					statemgr.SessionESTB,
					pCtxt.GetCurrentState(),
					statetype.EventPduSessEstbN4SessEstbResp)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
					return types.ErrFailTriggerFsm
				}
			} else {
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ctxt,
					"procedure context is not PduSessionEstbPrcdCtxt,or cause invalid.")

				var rejectMsg nasmsg.PduSessionEstbRejectMsg
				rejectMsg.MsgHeader.PduSessionID = pCtxt.PduSessId
				rejectMsg.MsgHeader.PrcdTransactionID = pCtxt.Pti
				rejectMsg.MsgHeader.MessageType = nas.PduSessEstabishReject
				rejectMsg.SMCause = nas.RequestRejectedUnspecified
				rejectBytes, err := rejectMsg.Encode()
				if err != nil {
					rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, ctxt, "fail to encode session reject msg, err:(%s)", err)
					return fmt.Errorf("fail to encode session reject msg, err:(%s)", err)
				}

				//msg counter
				pm.PegCounter(statistics.PduSessEstablishRejectCounter)

				var n11MsgData n11msg.N1N2MessageTransferReqData
				N1Container := n11msg.N1MessageContainerIE{N1MsgClass: n11msg.SM_N1Info, N1MessageContent: rejectBytes}

				// set n1 to msg data
				n11MsgData.N1MessageContainer = &N1Container
				n11MsgData.IeFlags.Set(n11msg.Ieid_n1MessageContainer)

				// set psi to msg data
				n11MsgData.SessionId = pCtxt.PduSessId
				n11MsgData.IeFlags.Set(n11msg.Ieid_pdusessionId)

				// set n2
				N2Container := n11msg.N2InfoContainerIE{N2InforClass: n11msg.PWS_RF}
				n11MsgData.N2InfoContainer = &N2Container
				n11MsgData.IeFlags.Set(n11msg.Ieid_n2InfoContainer)

				//send n11 msg to amf
				n1n2MsgReq := sbicmn.Trans_N11ToModels_N1N2MsgTransferReqFormat(&n11MsgData)

				smsender.SendSbiReqMsgN1N2MsgTransferReq(n1n2MsgReq, pCtxt.SbiMessage)

				pduSessCtxt.SetPrcdCtxt(nil)
			}
		case pfcp.PFCP_Session_Modification_Response:
			rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, nil, "handle PFCP_Session_Modification_Response")
			if cause == gctxt.Cause_Request_accepted {
				rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, nil, "cause == gctxt.Cause_Request_accepted")
				if pduSessCtxt.GetPrcdCtxt() != nil {
					//get the state manager
					stateMgr, ok := ctxt.Value(types.SmfStateMgrCK).(*statemgr.SmfStateMgr)
					if !ok {
						rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "get state manager failed")
						return fmt.Errorf("get state manager failed")
					}

					switch pduSessCtxt.GetPrcdCtxt().(type) {
					case *prcdctxt.PduSessionEstbPrcdCtxt:
						sessEstbpCtxt := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
						ctxt_new = context.WithValue(ctxt_new, types.PduSessionEstbPrcdCtxtCK, sessEstbpCtxt)
						//trigger the FSM
						rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, nil, "sessEstbpCtxt.GetCurrentState()", sessEstbpCtxt.GetCurrentState())
						stateMgr.PduSessESTBFsm.Bfsm.SetState(sessEstbpCtxt.GetCurrentState())
						stateMgr.PduSessESTBFsm.Bfsm.Event(statetype.EventPduSessEstbN4SessModResp, ctxt_new)

					case *prcdctxt.PduSessionModPrcdCtxt:

						sessModpCtxt := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionModPrcdCtxt)
						sessModpCtxt.Seid = pduSessCtxt.SEID
						ctxt_new = context.WithValue(ctxt_new, types.PduSessionModPrcdCtxtCK, sessModpCtxt)
						//trigger the FSM
						rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, nil, "sessModpCtxt.GetCurrentState()", sessModpCtxt.GetCurrentState())
						switch sessModpCtxt.GetCurrentState() {
						case statetype.StatePduSessModWf1stN4SessModResp:
							stateMgr.PduSessModFsm.Bfsm.SetState(sessModpCtxt.GetCurrentState())
							stateMgr.PduSessModFsm.Bfsm.Event(statetype.EventPduSessMod1stUpdateSmCtxtResp, ctxt_new)
						default:
							smsender.SendUpdateSmCtxtResponseMsg4SessModPrcd(sessModpCtxt)
							pduSessCtxt.SetPrcdCtxt(nil)
							rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, nil, "session modification done in SMF")
						}
					case *prcdctxt.AnRelSerReqPrcdCtxt:
						prcdCtxt := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.AnRelSerReqPrcdCtxt)
						ctxt_new = context.WithValue(ctxt_new, types.PduSessionAnRelSerReqCtxtCK, prcdCtxt)
						ctxt_new = context.WithValue(ctxt_new, types.SmfPduSessCtxtCK, pduSessCtxt)

						if prcdCtxt.IsAnRelease {
							smsender.SendUpdateSmCtxtResponseMsg4AnRelPrcd(prcdCtxt)
							//an release finished.
							pduSessCtxt.SetPrcdCtxt(nil)
						} else {
							//trigger the FSM
							rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, nil, "GetCurrentState()", prcdCtxt.GetCurrentState())
							stateMgr.PduSessESTBFsm.Bfsm.SetState(prcdCtxt.GetCurrentState())

							stateMgr.SerReqFsm.Bfsm.SetState(prcdCtxt.GetCurrentState())
							if !prcdCtxt.IsWfSecN4ModifyResp {
								//first n4 modify resp
								stateMgr.SerReqFsm.Bfsm.Event(statetype.EventPduSessSerReqN4ModifyResp, ctxt_new)
							} else {
								//second n4 modify resp
								stateMgr.SerReqFsm.Bfsm.Event(statetype.EventPduSessSerReqN4ModifyRespSec, ctxt_new)
							}
						}

					default:
						rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, nil, "unsupport sessionCtxt.GetPrcdCtxt()", pduSessCtxt.GetPrcdCtxt())
					}

				} else {
					rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, nil, "sessionCtxt.GetPrcdCtxt() == nil")

				}
			} else {
				return fmt.Errorf("pfcp cause fail")
			}
		case pfcp.PFCP_Session_Deletion_Response:
			if cause == gctxt.Cause_Request_accepted {
				pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionRelReqPrcdCtxt)
				if !ok {
					rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "fail to get ue ctxt by PduSessionRelReqPrcdCtxt")
					return fmt.Errorf("fail to get ue ctxt by PduSessionRelReqPrcdCtxt")
				}

				//ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(pCtxt.Imsi.GetValue()))
				//if err != nil {
				//	rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, nil, "failed  to find ue context by imsi %s", pCtxt.Imsi.String())
				//	return fmt.Errorf("failed to get ue context")
				//}

				if pCtxt.IsDeRegisterPrcd == false {
					//session release流程
					// send pdu session release accept msg  to UE，n2 resource release request msg to RAN
					err := smsender.SendPduSessRelAcceptMsg(ctxt, pCtxt)
					if err != nil {
						rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "failed to send pdu sess rel command msg")
					}
					// set the next state
					err = pCtxt.SetNextState(statetype.StatePduSessRelWfUpdateSmCtxtReqSec)
					if err != nil {
						rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to set state")
						return types.ErrFailSetState
					}
				} else {
					// deregister release one pdu session context
					gctxt.RemoveSmfSessionContext(pCtxt.Imsi, pCtxt.PduSessId)
					// SMF to AMF
					smsender.SendSbiRespMsgReleaseSMContextResponse(pCtxt.SbiMessage)
				}
			} else {
				return fmt.Errorf("pfcp cause fail")
			}
		case pfcp.PFCP_Session_Report_Request:
			rlogger.Trace(types.ModuleSmfN11, rlogger.INFO, pduSessCtxt, "sessionCtxt", pduSessCtxt)

			//TODO should handle more scenarios here... not just paging

			err = smsender.SendN1N2MsgTransfer4Paging(ctxt, pduSessCtxt, 0)
			if err != nil {
				rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "fail to send n1n2 msg")
				return fmt.Errorf("fail to send n1n2 msg")
			}
			//send n11 msg to amf
			//TODO support later...
			//var msg n11msg.N11Msg
			//msg.AmfSmCtxtId = sessionCtxt.AmfSmCtxtId
			//msg.SmfSmCtxtId = sessionCtxt.SmfSmCtxtId
			//msg.MsgType = n11msg.N1N2MsgReq
			//msg.MsgData = msgData
			//
			//amfScId := idmgrsmf.RetrieveScId(msg.AmfSmCtxtId)
			//routeragent.SendIpcMessage(ctxt_new, router.ScGR, amfScId, msg)
			//
			//err := routeragent.SendMsg2PfcpNode(ctxt_new, sessionCtxt, pfcp.PFCP_Session_Report_Response)
			//if err != nil {
			//	rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, sessionCtxt, "fail to send msg to pfcp node")
			//	return fmt.Errorf("fail to send msg to pfcp node")
			//}

		default:
			rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, ctxt_new, "unsupported pfcpv1 message type(%d)", msgData.Msg.ID)
		}
	} else {
		rlogger.Trace(types.ModuleSmfN11, rlogger.DEBUG, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	return nil
}
