package n11layer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/amf/sc/statistics"
	"lite5gc/amf/sc/utils"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
	"lite5gc/openapi/models"
	"strings"
)

func HandleN1N2MsgRequest(ctx context.Context, sbimsg *sbicmn.SbiHandlerMessage) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, ctx)

	// get msgData payload : SmContextCreateData
	modelsN1N2TransReqData := sbimsg.HTTPRequest.Body.(models.N1N2MessageTransferRequest)

	ueContextId := sbimsg.HTTPRequest.Params["ueContextId"] //imsi-460110002700001
	rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, nil, "ueContextId(%s)", ueContextId)

	msgData := sbicmn.Trans_ModelsToN11_N1N2MessageTransferReqDataFormat(modelsN1N2TransReqData)

	var ueCtxt *gctxt.UeContext
	var err error

	var imsi types3gpp.Imsi
	imsistr := strings.TrimPrefix(ueContextId, "imsi-")
	imsi.StoreImsiString(imsistr, types3gpp.CheckMncLen(imsistr))

	var psi nas.PduSessID
	if msgData.IeFlags.Test(n11msg.Ieid_pdusessionId) {
		psi = msgData.SessionId
	}

	ueCtxt, err = gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, nil, "failed to find the "+
			"ue context for sm ctxt id (%d), err: (%s)", imsi, err)
		return err
	}

	// get nas message
	var nasMsgData []byte
	if msgData.IeFlags.Test(n11msg.Ieid_n1MessageContainer) {
		switch msgData.N1MessageContainer.N1MsgClass {
		case n11msg.SM_N1Info:
			nasMsgData = append(nasMsgData, msgData.N1MessageContainer.N1MessageContent...)
		default: //todo handle the other types
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "unsupported N1MsgClass")
		}
	} else {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "no n1 message container")
	}

	//encode DL NAS Transport Msg Data
	var dlNasMsg []byte
	rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "dlNasMsg(%v)", dlNasMsg)
	if len(nasMsgData) != 0 {
		dlNasTransMsgData := &nasmsg.DownLinkNasTransportMsg{}
		if msgData.IeFlags.Test(n11msg.Ieid_pdusessionId) {
			dlNasTransMsgData.PduSessId = msgData.SessionId
			dlNasTransMsgData.OptIeBitSet.Set(nasmsg.IeidDownlinknastransPdusessid)
		} else {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "loss session id in N1N2MessageTransferReqData ")
		}
		dlNasTransMsgData.PayloadType = nasie.N1SmInformation

		//Suppose there is only one Payload Container Entry, if no, todo...
		dlNasTransMsgData.PayloadContainer.PayloadContainerEntry = make([]nasie.PayloadContainerENTRY, 1)
		dlNasTransMsgData.PayloadContainer.PayloadContainerEntry[0].ContainerContents =
			append(dlNasTransMsgData.PayloadContainer.PayloadContainerEntry[0].ContainerContents, nasMsgData...)

		dlNasData, err := dlNasTransMsgData.Encode()
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "failed to encode downlink nas transport message.")
		}

		//add sec nas header
		// set security protected header
		dlNasMsg, err = utils.EncodeSecPrctNasMsg(ueCtxt, nas.IntegrityPrtctCipher, dlNasData)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "failed to add security header")
			return err
		}

		//msg counter
		pm.PegCounter(statistics.DLNasTransportCounter)
	} else {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "no n1 message container ie exist")
	}

	//get n2 sm information
	var smInfo *n11msg.N2SmInformation
	if msgData.IeFlags.Test(n11msg.Ieid_n2InfoContainer) {
		n2InfoCont := msgData.N2InfoContainer
		switch n2InfoCont.N2InforClass {
		case n11msg.SM_N2Info:
			if n2InfoCont.IeFlags.Test(n11msg.Ieid_n2SmInfo) {
				smInfo = n2InfoCont.SmInfo
			} else {
				rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "no n2 sm information ie")
			}
		case n11msg.PWS_RF:
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "session fail in SMF, send session reject msg")
			if psi != 0 {
				sessionCtxt := ueCtxt.GetPduSessCtxt(psi)
				//delete temp info
				if !sessionCtxt.IsPSIExist {
					rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "send session reject msg, delete temp info, "+
						"psi(%v)", sessionCtxt.Psi)
					ueCtxt.DelPduSessCtxt(sessionCtxt.Psi)
				} else {
					rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "psi in invalid, fail to delete temp info")
				}
			}
		default:
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "unsupported N2InfoContainer.N2InforClass")
		}
	} else {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "no n2 info container ie")
	}

	//  get the scnglayer
	sender, ok := ctx.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer
	}

	//paging
	rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "deactive session list:(%v)", ueCtxt.GetPsiList(types3gpp.SessDeactive))
	if msgData.IeFlags.Test(n11msg.Ieid_pdusessionId) {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "psi:(%d)", msgData.SessionId)

		if ueCtxt.GetPduSessCtxt(msgData.SessionId).Status == types3gpp.SessDeactive {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, ueCtxt, "session is release, send paging msg")

			err := sender.SendPaging(ueCtxt)
			if err != nil {
				rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "failed to send pdu session res setup request message")
				return fmt.Errorf("fail to send paging msg")
			}
			// start t3513 timer for paging message
			//err = mmsender.StartNasTimer(ctx, ueCtxt, gctxt.T3513, bytes)
			//if err != nil {
			//	rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil,  "failed to start nas timer for timer(%s), gctxt.T3513, error(%s)",
			//		gctxt.T3513, err)
			//	return fmt.Errorf("failed to start nas timer for timer(%s), gctxt.T3550, error(%s)",
			//		gctxt.T3513, err)
			//}
			err = mmsender.SendSbiRespMsgN1N2MessageTransferRspData(sbimsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "failed to send N1N2MessageTransferRspData to Sbi Layer")
				return fmt.Errorf("failed to send N1N2MessageTransferRspData to Sbi Layer")
			}
			return nil
		}
	}

	if smInfo != nil {
		if smInfo.IeFlags.Test(n11msg.Ieid_n2InfoContent) {
			switch smInfo.N2InfoCont.NgapIeType {
			case n11msg.NgapPduResSetupReq:
				err = sender.SendPduSessResSetupRequest(ueCtxt, dlNasMsg, string(smInfo.N2InfoCont.NgapData),
					&smInfo.Snssai, msgData.SessionId)
				if err != nil {
					rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "failed to send pdu session res setup request message")
				}
				//increate nas counter
				nassecurity.UpdateDownlinkNasCounter(ueCtxt)

			default:
				rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "unsupported n2 info content")
			}
		} else {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "no n2 info content")
		}
	} else {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "no n2 info container")
	}

	err = mmsender.SendSbiRespMsgN1N2MessageTransferRspData(sbimsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ctx, "failed to send N1N2MessageTransferRspData to Sbi Layer")
		return fmt.Errorf("failed to send N1N2MessageTransferRspData to Sbi Layer")
	}
	return nil
}
