package naslayer

import (
	"bytes"
	"context"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/sc/statemgr"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func (p *NasMgr) HandleSmPduSessEstbRequest(ctx context.Context,
	header *nas.SmNasMessageHeader, plainNasMsg *bytes.Reader) nas.Sm5gCause {

	rlogger.FuncEntry(types.ModuleSmfNas, nil)
	if plainNasMsg == nil || header == nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "input para is nil")
		return nas.OtherValue
	}
	cause := nas.SuccessNoReason

	p.PduSessEstbRequest.Reset()
	rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "header:%v",header)
	//decode the PduSessEstbRequest message
	p.PduSessEstbRequest.Psi = header.PduSessionID
	p.PduSessEstbRequest.Pti = header.PrcdTransactionID
	err := p.PduSessEstbRequest.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil,
			"failed to decode pdu session establishment message, err: ", err)
		return nas.OtherValue
	}

	//get ue context
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "failed to get ue context")
		return nas.OtherValue
	}

	// check pti
	pti := header.PrcdTransactionID
	if pti < nas.MinPTI || pti > nas.MaxPTI {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "pti is out of range")
		return nas.InvalidPTIValue
	}

	//check the pdu session id exist
	psi := header.PduSessionID
	if psi < nas.MinPSI || psi > nas.MaxPSI {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "psi value is out of range")
		return nas.InvalidPDUSessionIdentity
	}

	if ueCtxt.GetPduSessCtxt(psi) != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "session psi(%v) is already exist", psi)
		return nas.InvalidPDUSessionIdentity
	}

	if p.PduSessEstbRequest.IeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_SessionType) {
		switch p.PduSessEstbRequest.SessionType {
		case types3gpp.Ipv4, types3gpp.Ipv4v6:
		case types3gpp.Ipv6, types3gpp.Unstructured, types3gpp.Ethernet:
			cause = nas.PDUSessionTypeIPv4OnlyAllowed
		default:
			return nas.UnknownPDUSessionType
		}
	}

	//store the information from create sm context request
	smCtxtCreateData, ok := ctx.Value(types.SmfN11MsgDataCK).(*n11msg.SmContextCreateData)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to get create sm ctxtData")
		return nas.OtherValue
	}

	// create a new pdu session context
	pduSessCtxt := gctxt.NewPduSessContext(psi)
	pduSessCtxt.PduSessionId = psi

	// create a pdu session establishment prcd ctxt
	prcdCtxt := prcdctxt.NewPduSessEstbPrcdCtxt(psi)
	pduSessCtxt.SetPrcdCtxt(prcdCtxt)

	// save pdu session context in ue context
	ueCtxt.PduSessCtxts[psi] = pduSessCtxt
	// save psi in context
	ctx = context.WithValue(ctx, types.PduSessIdCK, psi)

	//store the information in procedure ctxt
	//Mandatory
	prcdCtxt.IMSI = ueCtxt.IMSI
	prcdCtxt.PduSessId = pduSessCtxt.PduSessionId
	prcdCtxt.Pti = p.PduSessEstbRequest.Pti
	sbimsg, _ := ctx.Value(types.SbiMsgCK).(*sbicmn.SbiHandlerMessage)
	prcdCtxt.SbiMessage = sbimsg

	//Optional
	if p.PduSessEstbRequest.IeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_SessionType) {
		//如果session request消息中携带了SessionType，先用
		prcdCtxt.PduSessType = p.PduSessEstbRequest.SessionType
		rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"use the session type(%s) in session request msg", prcdCtxt.PduSessType)

	} else {
		//没有携带，再读取配置中的SessionType
		prcdCtxt.PduSessType = configure.SmfConf.Service.SessType
		rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, ueCtxt.GetImsiPtr(),
			"use the session type(%s) in smf config file", prcdCtxt.PduSessType)
	}

	if p.PduSessEstbRequest.IeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_SscMode) {
		//如果session request消息中携带了ssc，先用
		prcdCtxt.SscMode = p.PduSessEstbRequest.SscMode
	} else {
		//没有携带，再读取配置中的ssc
		prcdCtxt.SscMode = configure.SmfConf.Service.SSCMode
	}

	if p.PduSessEstbRequest.IeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_SmCapability) {
		prcdCtxt.SmCap = p.PduSessEstbRequest.SmCapability
	}

	if p.PduSessEstbRequest.IeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_MaxNumberOfSPF) {
		prcdCtxt.MaxNumofSupPktFilter = p.PduSessEstbRequest.MaxNumberOfSPF
	}

	if p.PduSessEstbRequest.IeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_AlwaysOn) {
		prcdCtxt.AlwaysOnPduSessReq = p.PduSessEstbRequest.AlwaysOn
	}

	prcdCtxt.SessionReqIeFlags = p.PduSessEstbRequest.IeFlags

	if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_dnn) {
		prcdCtxt.DNN = smCtxtCreateData.Dnn
		prcdCtxt.IeFlags.Set(n11msg.Ieid_dnn)
		rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"Dnn from Message (%s)", prcdCtxt.DNN)
	}

	if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_sNssai) {
		prcdCtxt.SNSSAI = smCtxtCreateData.SNssai
		rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"SNssai (%v)", smCtxtCreateData.SNssai)
		prcdCtxt.IeFlags.Set(n11msg.Ieid_sNssai)
	}
	if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_supi) {
		prcdCtxt.Supi = smCtxtCreateData.Supi
		prcdCtxt.IeFlags.Set(n11msg.Ieid_supi)
	}
	if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_selMode) {
		prcdCtxt.SelMode = smCtxtCreateData.SelMode
	}
	if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_unauthenticatedSupi) {
		prcdCtxt.IsAuthed = smCtxtCreateData.UnauthenticatedSupi
	}
	if smCtxtCreateData.IeFlags.Test(n11msg.Ieid_servingNfId) {
		prcdCtxt.ServingNfId = smCtxtCreateData.ServingNfId
	}

	err = statemgr.TriggerSmfFsm(ctx,
		statemgr.SessionESTB,
		statetype.StatePduSessEstbStart,
		statetype.EventPduSessEstbReq)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to trigger fsm")
		return nas.OtherValue
	}
	return cause
}
