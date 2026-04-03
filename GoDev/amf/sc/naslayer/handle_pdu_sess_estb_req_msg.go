package naslayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
)

func HandlePduSessEstabRequestMsg(ctx context.Context,
	ueCtxt *gctxt.UeContext,
	ulNasMsg *nasmsg.UplinkNasTransportMsg) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransPdusessid) == false {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "missing psi. ul nas transport:%s", ulNasMsg)
		return fmt.Errorf("missing PSI for pdu session establishment")
	}

	pduSess := &gctxt.AmfPduSessCtxt{}

	if ueCtxt.IsPduSessExist(ulNasMsg.PduSessId) {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "pdu session id (%d) is already exist", ulNasMsg.PduSessId)
		// return fmt.Errorf("psi(%d) duplicated", ulNasMsg.PduSessId)
		// let the new pdu session establish procedure continue...TODO
		// ueCtxt.DelPduSessCtxt(ulNasMsg.PduSessId)
		pduSess = ueCtxt.GetPduSessCtxt(ulNasMsg.PduSessId)
		pduSess.IsPSIExist = true //如果此时返回session reject，不需要删除psi - amf session context
	} else {
		// allocated new session context
		pduSess = gctxt.NewAmfPduSessCtxt(ulNasMsg.PduSessId)
		pduSess.Psi = ulNasMsg.PduSessId
		pduSess.Status = types3gpp.SessActived

		//add session ctxt into ueCtxt
		err := ueCtxt.AddPduSessCtxt(pduSess.Psi, pduSess)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to add pdu Session Ctxt")
			return err
		}

		if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransSnssai) {
			pduSess.SNssai = ulNasMsg.SNssai
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueCtxt, "use the snssai from uplink nas msg directly,%s", pduSess.SNssai.String())
		} else {
			// 23.502 4.3.2.2.1 Step2
			// If the NAS message does not contain an S-NSSAI,
			// the AMF determines a default S-NSSAI for the requested PDU Session
			// either according to the UE subscription, if it contains only one default S-NSSAI,
			// or based on operator policy.
			pduSess.SNssai = ueCtxt.AccMobSubsData.Nssai.DefSnssai
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueCtxt, "use the snssai from subscriber data,%s", pduSess.SNssai.String())
		}
		pduSess.Status = types3gpp.SessActived

		if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransOldpdusessid) {
			pduSess.OldPsi = ulNasMsg.OldPduSessId
			if ueCtxt.IsPduSessExist(pduSess.OldPsi) {
				//todo the old psi exist, select the same smf
				// 23.502 4.3.5.2
				//	In Step 1 of clause 4.3.2.2.1, according to the SSC mode, UE generates a new PDU Session ID and initiates the PDU Session Establishment Request using the new PDU Session ID. The new PDU Session ID is included as PDU Session ID in the NAS request message, and the Old PDU Session ID which indicates the existing PDU Session to be released is also provided to AMF in the NAS request message.
				//	In Step 2 of clause 4.3.2.2.1, if SMF reallocation was requested in Step 2 of this clause, the AMF selects a different SMF. Otherwise, the AMF sends the Nsmf_PDUSession_CreateSMContext Request to the same SMF serving the Old PDU Session ID.
				//	In Step 3 of clause 4.3.2.2.1, the AMF include both PDU Session ID and Old PDU Session ID in Nsmf_PDUSession_CreateSMContext Request. The SMF detects that the PDU Session establishment request is related to the trigger in step 2 based on the presence of an Old PDU Session ID in the Nsmf_PDUSession_CreateSMContext Request. The SMF stores the new PDU Session ID and selects a new PDU Session Anchor (i.e. UPF2) for the new PDU Session.
			}
		}

		if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransRequesttype) {
			pduSess.ReqType = ulNasMsg.RequestType
		} else {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "Request Type should be include in the message.")
		}

		if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransDnn) {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueCtxt, "Set dnn from message, %s", ulNasMsg.Dnn)
			pduSess.Dnn = ulNasMsg.Dnn
		} else {
			// When the NAS Message contains an S-NSSAI but it does not contain a DNN,
			// the AMF determines the DNN for the requested PDU Session by selecting
			// the default DNN  for this S-NSSAI if the default DNN is present in the UE's Subscription Information;
			// otherwise the serving AMF selects a locally configured DNN for this S-NSSAI.
			if ulNasMsg.OptIeBitSet.Test(nasmsg.IeidUplinknastransSnssai) {
				dnn, err := ueCtxt.SmfSelSubsData.GetDefDnn(&ulNasMsg.SNssai)
				if err == nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueCtxt, "use the dnn from subscriber data")
					pduSess.Dnn = *dnn
				} else {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "use the dnn from local confiuration")
					// set the dnn from local configuration
					err = pduSess.Dnn.StoreWithString(configure.SmfConf.DnnInfo[0].Name)
					if err != nil {
						rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail get dnn")
						return types.ErrFailParseDNN
					}
				}
			} else {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "us"+
					"e the dnn from local confiuration")
				// set the dnn from local configuration
				err = pduSess.Dnn.StoreWithString(configure.SmfConf.DnnInfo[0].Name)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail get dnn")
					return types.ErrFailParseDNN
				}
			}
		}
	}
	// create create session context and send to smf with N11 if
	createSmCtxtReq := n11msg.SmContextCreateData{}
	createSmCtxtReq.SetSupi(ueCtxt.GetSupi())
	createSmCtxtReq.SetPduSessionId(pduSess.Psi)
	createSmCtxtReq.SetOldPduSessionId(pduSess.OldPsi)
	createSmCtxtReq.SetRequestType(n11msg.RequestType(pduSess.ReqType))
	createSmCtxtReq.SetSnssai(&pduSess.SNssai)
	createSmCtxtReq.SetDnn(&pduSess.Dnn)
	createSmCtxtReq.ServingNfId = configure.AmfConf.Service.AmfName
	createSmCtxtReq.AnType = nasmsg.Access3gpp
	createSmCtxtReq.SmContextStatusUri = "todo url"
	createSmCtxtReq.SetUnauthenticatedSupi(false)
	createSmCtxtReq.SetGuami(configure.GetTypesGuami())
	createSmCtxtReq.ServingNetwork = configure.AmfConf.PlmnList.List[0] // todo
	createSmCtxtReq.SetRatType(types3gpp.RatType_NR)                    //todo should get from n2context
	createSmCtxtReq.SetN1SmMsg(ulNasMsg.PayloadContainer.PayloadContainerEntry[0].ContainerContents)

	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, ueCtxt, "createSmCtxtReq %v", createSmCtxtReq)
	mmsender.SendCreateSMCtxtSBIMsg(ctx, ueCtxt.GetImsi(), createSmCtxtReq)

	return nil
}
